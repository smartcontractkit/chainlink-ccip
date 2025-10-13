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
	RemoteChainSelector    uint64
	OutboundCCVs           []common.Address
	AdditionalOutboundCCVs []common.Address
	InboundCCVs            []common.Address
	AdditionalInboundCCVs  []common.Address
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

var SiloedLockReleaseTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockBox\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"additionalOutboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"additionalInboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyFinalityConfigUpdates\",\"inputs\":[{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"destToUseDefaultFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAccumulatedFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAvailableTokens\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"lockedTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChainRebalancer\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRebalancer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredInboundCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredOutboundCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUnsiloedLiquidity\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"out\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provideLiquidity\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"provideSiloedLiquidity\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustomFinalityRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRebalancer\",\"inputs\":[{\"name\":\"newRebalancer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSiloRebalancer\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newRebalancer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"updateSiloDesignations\",\"inputs\":[{\"name\":\"removes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"adds\",\"type\":\"tuple[]\",\"internalType\":\"structSiloedLockReleaseTokenPool.SiloConfigUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rebalancer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawLiquidity\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawSiloedLiquidity\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"additionalOutboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"additionalInboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainSiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"rebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainUnsiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"amountUnsiloed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityTransferInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalityConfigUpdated\",\"inputs\":[{\"name\":\"finalityConfig\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeeWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SiloRebalancerSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"oldRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnsiloedRebalancerSet\",\"inputs\":[{\"name\":\"oldRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotSiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InsufficientLiquidity\",\"inputs\":[{\"name\":\"availableLiquidity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidChainSelector\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidFinality\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidFinalityConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"LiquidityAmountCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenTransferFeeConfigNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x610120806040523461040d57617fd1803803809161001d82856106c3565b8339810160c08282031261040d5781516001600160a01b0381169290919083830361040d5761004e602082016106e6565b60408201516001600160401b03811161040d5782019280601f8501121561040d578351936001600160401b038511610412578460051b90602082019561009760405197886106c3565b865260208087019282010192831161040d57602001905b8282106106ab575050506100c4606083016106f4565b6100dc60a06100d5608086016106f4565b94016106f4565b94331561069a57600180546001600160a01b0319163317905586158015610689575b8015610678575b6104de5760805260c05260405163313ce56760e01b8152602081600481895afa6000918161063c575b50610611575b5060a052600480546001600160a01b0319166001600160a01b03929092169190911790558051151560e08190526104ef575b506001600160a01b031680156104de57604051636eb1769f60e11b815230600482015260248101829052602081604481865afa9081156104d2576000916104a0575b506104355760405191602083019263095ea7b360e01b84528260248201526000196044820152604481526101dd6064826106c3565b6000806040958651936101f088866106c3565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082865af13d15610428573d906001600160401b038211610412578551610261949092610252601f8201601f1916602001856106c3565b83523d6000602085013e6108a8565b805180610391575b505061010052516176589081610979823960805181818161041f015281816116df01528181611b0e01528181611dc501528181611f66015281816126a7015281816128950152818161397d01528181613bfe01528181613cc801528181613e650152818161404e0152818161453d01528181614a9401528181614aef01528181614c6b01528181614d9e01526158e4015260a051818181611e340152818161297d01528181614a3d01528181616289015261630c015260c051818181610f6301528181611ba90152818161274201528181613a180152613f01015260e051818181610f1001528181611bee01528181612787015261372b0152610100518181816104950152818161167b01528181612a4401528181614124015281816145850152614d360152f35b816020918101031261040d576020015180159081150361040d576103b6573880610269565b815162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b91610261926060916108a8565b60405162461bcd60e51b815260206004820152603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527f20746f206e6f6e2d7a65726f20616c6c6f77616e6365000000000000000000006064820152608490fd5b90506020813d6020116104ca575b816104bb602093836106c3565b8101031261040d5751386101a8565b3d91506104ae565b6040513d6000823e3d90fd5b630a64406560e11b60005260046000fd5b906020906040519061050183836106c3565b60008252600036813760e051156106005760005b825181101561057c576001906001600160a01b036105338286610708565b51168561053f8261074a565b61054c575b505001610515565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a13885610544565b5092905060005b81518110156105f7576001906001600160a01b036105a18285610708565b511680156105f157846105b382610848565b6105c1575b50505b01610583565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a138846105b8565b506105bb565b50505038610166565b6335f4a7b360e01b60005260046000fd5b60ff1660ff82168181036106255750610134565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d602011610670575b81610658602093836106c3565b8101031261040d57610669906106e6565b903861012e565b3d915061064b565b506001600160a01b03821615610105565b506001600160a01b038416156100fe565b639b15e16f60e01b60005260046000fd5b602080916106b8846106f4565b8152019101906100ae565b601f909101601f19168101906001600160401b0382119082101761041257604052565b519060ff8216820361040d57565b51906001600160a01b038216820361040d57565b805182101561071c5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b805482101561071c5760005260206000200190600090565b600081815260036020526040902054801561084157600019810181811161082b5760025460001981019190821161082b578181036107da575b50505060025480156107c4576000190161079e816002610732565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b6108136107eb6107fc936002610732565b90549060031b1c9283926002610732565b819391549060031b91821b91600019901b19161790565b90556000526003602052604060002055388080610783565b634e487b7160e01b600052601160045260246000fd5b5050600090565b806000526003602052604060002054156000146108a25760025468010000000000000000811015610412576108896107fc8260018594016002556002610732565b9055600254906000526003602052604060002055600190565b50600090565b9192901561090a57508151156108bc575090565b3b156108c55790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b82519091501561091d5750805190602001fd5b6040519062461bcd60e51b8252602060048301528181519182602483015260005b8381106109605750508160006044809484010152601f80199101168101030190fd5b6020828201810151604487840101528593500161093e56fe60a080604052600436101561001357600080fd5b60006080526080513560e01c90816301ffc9a714614e8c575080630a861f2a14614c9f578063164e68de14614b9b578063181f5a7714614b1357806321df0da714614ac1578063240028e814614a6157806324f65ee714614a225780632a10097b146147c85780632c286daf146146be5780632d4a148f1461445857806331238ffc1461441057806337b19247146142fc5780633907753714613dd7578063432a6ba314613da2578063489a68f2146138da5780634c5ef0ed146138955780634f71592c1461387757806354c8a4f3146136f95780635df45a37146136dd57806362ddd3c4146136585780636600f92c146135385780636cfd15531461347e5780636d3d1a58146134495780636d9d216c1461300e5780637437ff9f14612fce57806379ba509714612ef75780637d54534e14612e66578063804ba5a914612e245780638632d5cc14612de35780638926f54f14612d9e5780638da5cb5b14612d69578063962d402014612c0c5780639a4575b9146126215780639f68f67314612603578063a42a7b8b1461248e578063a7cd63b714612427578063acfecf9114612303578063af0e58b9146122e4578063af58d59f14612298578063b184b942146121a2578063b1c71c6514611a84578063b794658014611a4c578063c4bffe2b14611919578063c75eea9c1461186e578063ce3c7528146115af578063cf7401f314611403578063d966866b14610f87578063dc0bd97114610f35578063e0351e1314610ef7578063e8a1da17146105e2578063eb521a4c14610376578063f1e733991461034b5763f2fde38b1461026c57600080fd5b346103455760206003193601126103455773ffffffffffffffffffffffffffffffffffffffff61029a614fbd565b6102a2615e87565b163381146103195760805180547fffffffffffffffffffffffff0000000000000000000000000000000000000000168217815560015473ffffffffffffffffffffffffffffffffffffffff16907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789080a360805180f35b7fdad89dca00000000000000000000000000000000000000000000000000000000608051526004608051fd5b60805180fd5b3461034557602060031936011261034557602061036e6103696150e4565b615e00565b604051908152f35b3461034557602060031936011261034557600435600081156105ba5773ffffffffffffffffffffffffffffffffffffffff6103b082615963565b16330361058e57808052601160205260408120600181015460a01c60ff1615610579576103de8382546156b8565b90555b6040517f23b872dd000000000000000000000000000000000000000000000000000000006020820152336024820152306044820152606481018390527f00000000000000000000000000000000000000000000000000000000000000009061047e9061047881608481015b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826151e4565b82616a5e565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690813b15610575576040517f47e7ef2400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff919091166004820152602481018490529082908290604490829084905af1801561056a5761055a575b509060405191825260208201527f569a440e6842b5e5a7ac02286311855f5a0b81b9390909e552e82aaf02c9e9bf60403392a260805180f35b81610564916151e4565b82610521565b6040513d84823e3d90fd5b8280fd5b5061058682600f546156b8565b600f556103e1565b807f8e4a23d6000000000000000000000000000000000000000000000000000000006024925233600452fd5b807fa90c0d190000000000000000000000000000000000000000000000000000000060049252fd5b34610345576105f03661537c565b926105fc929192615e87565b608051905b828210610d355750505060805163ffffffff4216917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee181360301915b84811015610d2f576000938160051b8301358481121561034557830190610120823603126103455760405195610672876151c8565b82359067ffffffffffffffff82168203610d2c57508652602082013567ffffffffffffffff81116103455782019436601f870112156103455785356106b6816156c5565b966106c460405198896151e4565b81885260208089019260051b820101903682116103455760208101925b828410610cfd575050505060208701958652604083013567ffffffffffffffff8111610345576107149036908501615296565b916040880192835261073e61072c3660608701615490565b9460608a0195865260c0369101615490565b95608089019687526107508551616917565b61075a8751616917565b83515115610cd15761077667ffffffffffffffff8a51166174d0565b15610c965767ffffffffffffffff8951166080515260076020526040608051206108ba86516fffffffffffffffffffffffffffffffff604082015116906108756fffffffffffffffffffffffffffffffff602083015116915115158360806040516107e0816151c8565b858152602081018b905260408101849052606081018690520152855474ff000000000000000000000000000000000000000091151560a01b919091167fffffffffffffffffffffff0000000000000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff84161773ffffffff0000000000000000000000000000000060808a901b1617178555565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176001830155565b6109e088516fffffffffffffffffffffffffffffffff6040820151169061099b6fffffffffffffffffffffffffffffffff60208301511691511515836080604051610904816151c8565b858152602081018b9052604081018490526060810186905201526002860180547fffffffffffffffffffffff000000000000000000000000000000000000000000166fffffffffffffffffffffffffffffffff85161773ffffffff0000000000000000000000000000000060808b901b161791151560a01b74ff000000000000000000000000000000000000000016919091179055565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176003830155565b6004855191019080519067ffffffffffffffff8211610c6557610a038354615aee565b601f8111610c26575b506020906001601f841114610b8057918091610a5d9360805192610b75575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b6080515b88518051821015610a995790610a93600192610a8c8367ffffffffffffffff8f51169261572c565b519061642e565b01610a64565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2939199975095610b6767ffffffffffffffff6001979694985116925193519151610b33610afe60405196879687526101006020880152610100870190615001565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a101939290919361063d565b015190508e80610a2b565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083169184608051528160805120926080515b818110610c0e5750908460019594939210610bd7575b505050811b019055610a60565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558d8080610bca565b92936020600181928786015181550195019301610bb4565b610c55908460805152602060805120601f850160051c81019160208610610c5b575b601f0160051c0190615d9a565b8d610a0c565b9091508190610c48565b7f4e487b71000000000000000000000000000000000000000000000000000000006080515260416004526024608051fd5b67ffffffffffffffff8951167f1d5ad3c500000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b7f14c880ca00000000000000000000000000000000000000000000000000000000608051526004608051fd5b833567ffffffffffffffff811161034557602091610d218392833691870101615296565b8152019301926106e1565b80fd5b60805180f35b9092919367ffffffffffffffff610d55610d5086888661555a565b615516565b1692610d6084617211565b15610ec75783608051526007602052610d80600560406080512001617018565b926080515b8451811015610dbf5760019086608051526007602052610db8600560406080512001610db1838961572c565b519061733c565b5001610d85565b5093909491959250806080515260076020526005604060805120608051815560805160018201556080516002820155608051600382015560048101610e048154615aee565b80610e76575b505001805490608051815581610e53575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a1019091610601565b60805152602060805120908101905b81811015610e1b5760008155600101610e62565b601f8111600114610e8e5750600090555b8880610e0a565b610eae9082608051526001601f6020608051209201861c82019101615d9a565b8060805152600060206080512091556080519055610e87565b837f1e670e4b00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b34610345576080516003193601126103455760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b346103455760805160031936011261034557602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346103455760206003193601126103455760043567ffffffffffffffff811161034557610fb8903690600401615060565b610fc0615e87565b608051915b818310610fd25760805180f35b610fe0610d50848484615cf3565b610ff8610fee858585615cf3565b6020810190615d33565b9091611012611008878787615cf3565b6040810190615d33565b9061102b611021898989615cf3565b6060810190615d33565b909161104561103b8b8b8b615cf3565b6080810190615d33565b94909761105b611056368a84615847565b61684d565b611069611056368486615847565b611077611056368688615847565b61108561105636888c615847565b60405161109181615129565b61109c368a84615847565b81526110a9368486615847565b60208201526110b9368688615847565b60408201526110c936888c615847565b606082015267ffffffffffffffff881660805152600d602052604060805120815180519067ffffffffffffffff8211610c6557680100000000000000008211610c655760209083548385558084106113e4575b500182608051526020608051206080515b8381106113ba5750505050602082015180519067ffffffffffffffff8211610c6557680100000000000000008211610c65576020906001840154836001860155808410611398575b500160018301608051526020608051206080515b83811061136e5750505050604082015180519067ffffffffffffffff8211610c6557680100000000000000008211610c6557602090600284015483600286015580841061134c575b500160028301608051526020608051206080515b838110611322575050505060036060919e9c9d9e019101519081519167ffffffffffffffff8311610c6557680100000000000000008311610c65576020908254848455808510611303575b500190608051526020608051206080515b8381106112d957505050506112be6080956112ce956112b07fece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d9a9660019e9d9c9a966112a267ffffffffffffffff976040519d8d8f9e8f9081520191615db1565b918b830360208d0152615db1565b9188830360408a0152615db1565b9285840360608701521696615db1565b0390a2019190610fc5565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501611241565b61131c9084608051528584608051209182019101615d9a565b38611230565b600190602073ffffffffffffffffffffffffffffffffffffffff85511694019381840155016111e5565b6113689060028601608051528484608051209182019101615d9a565b386111d1565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501611189565b6113b49060018601608051528484608051209182019101615d9a565b38611175565b600190602073ffffffffffffffffffffffffffffffffffffffff855116940193818401550161112d565b6113fd9085608051528484608051209182019101615d9a565b3861111c565b346103455760e06003193601126103455761141c6150e4565b60607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc3601126115aa5760405161145281615190565b60243580151581036115aa5781526044356fffffffffffffffffffffffffffffffff811681036115aa5760208201526064356fffffffffffffffffffffffffffffffff811681036115aa57604082015260607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c3601126115aa57604051906114d982615190565b60843580151581036115aa57825260a4356fffffffffffffffffffffffffffffffff811681036115aa57602083015260c4356fffffffffffffffffffffffffffffffff811681036115aa57604083015273ffffffffffffffffffffffffffffffffffffffff6009541633141580611588575b61155857610d2f92616698565b7f8e4a23d60000000000000000000000000000000000000000000000000000000060805152336004526024608051fd5b5073ffffffffffffffffffffffffffffffffffffffff6001541633141561154b565b600080fd5b34610345576040600319360112610345576115c86150e4565b60243567ffffffffffffffff82168060805152601160205260ff6001604060805120015460a01c16158015611866575b61183757600090821561180f5773ffffffffffffffffffffffffffffffffffffffff61162385615963565b1633036117e3578152601160205260408120600181015460a01c60ff1680156117db5781545b8085116117ab575015611796576116618382546155bb565b90555b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803b15611792576040517f69328dec00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166004820152602481018490523360448201529082908290606490829084905af1801561056a57917f58fca2457646a9f47422ab9eb9bff90cef88cd8b8725ab52b1d17baa392d784e939161177993611782575b50506040519182913395836020909392919367ffffffffffffffff60408201951681520152565b0390a260805180f35b8161178c916151e4565b85611752565b5080fd5b506117a382600f546155bb565b600f55611664565b83856044927fa17e11d5000000000000000000000000000000000000000000000000000000008352600452602452fd5b600f54611649565b6024827f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b6004827fa90c0d19000000000000000000000000000000000000000000000000000000008152fd5b7f46f5f12b00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b5080156115f8565b346103455760206003193601126103455767ffffffffffffffff6118906150e4565b611898615c40565b50166080515260076020526119156118bc6118b7604060805120615c6b565b6167c8565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b0390f35b3461034557608051600319360112610345576040516005548082528160208101600560005260206000209260005b818110611a3357505061195c925003826151e4565b80519061198161196b836156c5565b9261197960405194856151e4565b8084526156c5565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06020840192013683376080515b81518110156119e2578067ffffffffffffffff6119cf6001938561572c565b51166119db828761572c565b52016119b0565b505090604051918291602083019060208452518091526040830191906080515b818110611a10575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101611a02565b8454835260019485019486945060209093019201611947565b3461034557602060031936011261034557611915611a70611a6b6150e4565b615cd1565b604051918291602083526020830190615001565b346103455760606003193601126103455760043567ffffffffffffffff81116103455760a0600319823603011261034557611abd615091565b9060443567ffffffffffffffff811161034557611ade903690600401615296565b50611ae76159cf565b506084810190611af6826155f7565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691160361215457602481019077ffffffffffffffff00000000000000000000000000000000611b5c83615516565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561205d5760805191612125575b506120f957611bec604482016155f7565b7f0000000000000000000000000000000000000000000000000000000000000000612099575b5067ffffffffffffffff611c2583615516565b16611c3d816000526006602052604060002054151590565b1561206a57602073ffffffffffffffffffffffffffffffffffffffff60045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa801561205d5760805190611ffa575b73ffffffffffffffffffffffffffffffffffffffff9150163303611fca57606461ffff91013593169283151592838094611fbb575b15611f005761ffff600a541694858110611ecc5750611e909450611d28611d18611cfe85615516565b67ffffffffffffffff16600052600b602052604060002090565b83611d22846155f7565b91616dd1565b7f7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b267ffffffffffffffff611d64611d5e86615516565b936155f7565b6040805173ffffffffffffffffffffffffffffffffffffffff929092168252602082018690529190931692a25b918290611e9a575b50611a6b81611daa611e2c93615516565b6040805173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815233602082015290810186905267ffffffffffffffff91909116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1090606090a2615516565b9060405160ff7f000000000000000000000000000000000000000000000000000000000000000016602082015260208152611e686040826151e4565b60405192611e75846151ac565b8352602083015260405192839260408452604084019061543c565b9060208301520390f35b611e2c919250611ec4611a6b91612710611ebd61ffff600a5460101c1683615d87565b04906155bb565b929150611d99565b85907fe08f03ef00000000000000000000000000000000000000000000000000000000608051526004526024526044608051fd5b50611e90935067ffffffffffffffff611f1883615516565b16806080515260076020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789448280611f8e60406080512073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391616dd1565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a2611d91565b5061ffff600a54161515611cd5565b7f728fe07b0000000000000000000000000000000000000000000000000000000060805152336004526024608051fd5b506020813d602011612055575b81612014602093836151e4565b81010312610345575173ffffffffffffffffffffffffffffffffffffffff811681036103455773ffffffffffffffffffffffffffffffffffffffff90611ca0565b3d9150612007565b6040513d608051823e3d90fd5b7fa9902c7e00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b73ffffffffffffffffffffffffffffffffffffffff166120c6816000526003602052604060002054151590565b611c12577fd0d2597600000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b7f53ad11d800000000000000000000000000000000000000000000000000000000608051526004608051fd5b612147915060203d60201161214d575b61213f81836151e4565b810190616416565b85611bdb565b503d612135565b73ffffffffffffffffffffffffffffffffffffffff612172836155f7565b7f961c9a4f0000000000000000000000000000000000000000000000000000000060805152166004526024608051fd5b34610345576040600319360112610345576121bb614fbd565b6024356bffffffffffffffffffffffff81168103610345576121db615e87565b73ffffffffffffffffffffffffffffffffffffffff8216918215610cd1577f39b29a1c46dd0e3ee5683fed1d66e11a234938cde48e5c4e14389be508348eb692600454907fffffffffffffffffffffffff00000000000000000000000000000000000000008460a01b169160805116171760045561228f604051928392839092916bffffffffffffffffffffffff60209173ffffffffffffffffffffffffffffffffffffffff604085019616845216910152565b0390a160805180f35b346103455760206003193601126103455767ffffffffffffffff6122ba6150e4565b6122c2615c40565b50166080515260076020526119156118bc6118b7600260406080512001615c6b565b3461034557608051600319360112610345576020600f54604051908152f35b346103455767ffffffffffffffff61231a366153ca565b929091612325615e87565b169061233e826000526006602052604060002054151590565b156123f7578160805152600760205261237160056040608051200161236436868561525f565b602081519101209061733c565b156123b0577f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d769192611779604051928392602084526020840191615c01565b6123f3906040519384937f74f23c7c0000000000000000000000000000000000000000000000000000000085526004850152604060248501526044840191615c01565b0390fd5b507f1e670e4b00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b3461034557608051600319360112610345576040516002548082526020820190600260005260206000209060005b818110612478576119158561246c818703826151e4565b6040519182918261532c565b8254845260209093019260019283019201612455565b346103455760206003193601126103455767ffffffffffffffff6124b06150e4565b166080515260076020526124cb600560406080512001617018565b8051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06125116124fb846156c5565b9361250960405195866151e4565b8085526156c5565b016080515b8181106125f25750506080515b815181101561256d57806125396001928461572c565b51608051526008602052612551604060805120615b41565b61255b828661572c565b52612566818561572c565b5001612523565b826040518091602082016020835281518091526040830190602060408260051b860101930191608051905b8282106125a757505050500390f35b919360206125e2827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851615001565b9601920192018594939192612598565b806060602080938701015201612516565b346103455761191561246c612617366152b4565b50505091506159e8565b346103455760206003193601126103455760043567ffffffffffffffff81116103455760a060031982360301126103455761265a6159cf565b506126636159cf565b5060405160209061267482826151e4565b60805190526126816159cf565b506084820161268f816155f7565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603612bee5750602482019177ffffffffffffffff000000000000000000000000000000006126f684615516565b60801b16604051907f2cbc26bb0000000000000000000000000000000000000000000000000000000082526004820152828160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561205d5760805191612bd1575b506120f957612785604482016155f7565b7f0000000000000000000000000000000000000000000000000000000000000000612b71575b5067ffffffffffffffff6127be84615516565b166127d6816000526006602052604060002054151590565b1561206a578273ffffffffffffffffffffffffffffffffffffffff60045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa801561205d5760805190612b13575b73ffffffffffffffffffffffffffffffffffffffff9150163303611fca57606401358067ffffffffffffffff61286d85615516565b1680600052600784526128bd604060002073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016938491616dd1565b6040805173ffffffffffffffffffffffffffffffffffffffff84168152602081018590527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789449190a267ffffffffffffffff61291785615516565b6040805173ffffffffffffffffffffffffffffffffffffffff851681523360208201529081018590529116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1090606090a2612975611a6b85615516565b9360405160ff7f000000000000000000000000000000000000000000000000000000000000000016858201528481526129af6040826151e4565b604051956129bc876151ac565b86528486015267ffffffffffffffff6129d482615516565b16608051526011845260ff6001604060805120015460a01c16600014612afe578067ffffffffffffffff612a0a612a2b93615516565b166080515260118552604060805120612a248582546156b8565b9055615516565b505b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691823b15610345576040517f47e7ef2400000000000000000000000000000000000000000000000000000000815260805173ffffffffffffffffffffffffffffffffffffffff909316600482015260248101919091529182908180604481010391608051905af1801561205d57612aeb575b5061191560405192828493845283019061543c565b608051612af7916151e4565b6000612ad6565b50612b0b82600f546156b8565b600f55612a2d565b508281813d8311612b6a575b612b2981836151e4565b81010312610345575173ffffffffffffffffffffffffffffffffffffffff811681036103455773ffffffffffffffffffffffffffffffffffffffff90612838565b503d612b1f565b73ffffffffffffffffffffffffffffffffffffffff16612b9e816000526003602052604060002054151590565b6127ab577fd0d2597600000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b612be89150833d851161214d5761213f81836151e4565b84612774565b61217273ffffffffffffffffffffffffffffffffffffffff916155f7565b346103455760606003193601126103455760043567ffffffffffffffff811161034557612c3d903690600401615060565b9060243567ffffffffffffffff811161034557612c5e90369060040161540b565b9060443567ffffffffffffffff811161034557612c7f90369060040161540b565b73ffffffffffffffffffffffffffffffffffffffff6009541633141580612d47575b61155857838614801590612d3d575b612d11576080515b868110612cc55760805180f35b80612d0b612cd9610d506001948b8b61555a565b612ce48389896159bf565b612d05612cfd612cf586898b6159bf565b923690615490565b913690615490565b91616698565b01612cb8565b7f568efce200000000000000000000000000000000000000000000000000000000608051526004608051fd5b5080861415612cb0565b5073ffffffffffffffffffffffffffffffffffffffff60015416331415612ca1565b346103455760805160031936011261034557602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b34610345576020600319360112610345576020612dd967ffffffffffffffff612dc56150e4565b166000526006602052604060002054151590565b6040519015158152f35b34610345576020600319360112610345576020612e06612e016150e4565b615963565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b346103455760206003193601126103455760043567ffffffffffffffff811161034557612e58610d2f9136906004016150b3565b90612e61615e87565b615eef565b34610345576020600319360112610345577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174602073ffffffffffffffffffffffffffffffffffffffff612eb7614fbd565b612ebf615e87565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000006009541617600955604051908152a160805180f35b3461034557608051600319360112610345576080515473ffffffffffffffffffffffffffffffffffffffff81163303612fa2577fffffffffffffffffffffffff0000000000000000000000000000000000000000600154913382841617600155166080515573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0608051608051a360805180f35b7f02b543c600000000000000000000000000000000000000000000000000000000608051526004608051fd5b3461034557608051600319360112610345576004546040805173ffffffffffffffffffffffffffffffffffffffff8316815260a09290921c602083015290f35b346103455760406003193601126103455760043567ffffffffffffffff81116103455761303f903690600401615060565b6024359167ffffffffffffffff831161034557366023840112156103455782600401359167ffffffffffffffff8311610345576024840193602436918560061b0101116103455761308e615e87565b6080515b818110613310575050506080515b8181106130ad5760805180f35b67ffffffffffffffff6130c4610d50838587615953565b161580156132d9575b80156132b8575b6132715773ffffffffffffffffffffffffffffffffffffffff61310360206130fd848688615953565b016155f7565b1615610cd157806131f961311f60206130fd6001958789615953565b8573ffffffffffffffffffffffffffffffffffffffff808660405161314381615190565b608051815282602082019616865267ffffffffffffffff61316f610d508a8d6040860199878b52615953565b166080515260116020526040608051209051815501935116167fffffffffffffffffffffffff00000000000000000000000000000000000000008354161782555115157fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b7f180c6940bd64ba8f75679203ca32f8be2f629477a3307b190656e4b14dd5ddeb613228610d50838688615953565b61323860206130fd85888a615953565b6040805167ffffffffffffffff93909316835273ffffffffffffffffffffffffffffffffffffffff91909116602083015290a1016130a0565b610d50906132889267ffffffffffffffff94615953565b7fd9a9cd680000000000000000000000000000000000000000000000000000000060805152166004526024608051fd5b506132d367ffffffffffffffff612dc5610d50848688615953565b156130d4565b5067ffffffffffffffff6132f1610d50838587615953565b1660805152601160205260ff6001604060805120015460a01c166130cd565b67ffffffffffffffff613327610d5083858761555a565b1660805152601160205260ff6001604060805120015460a01c1615613402578067ffffffffffffffff613360610d50600194868861555a565b166080515260116020527f7b5efb3f8090c5cfd24e170b667d0e2b6fdc3db6540d75b86d5b6655ba00eb936040608051205461339e81600f546156b8565b600f5567ffffffffffffffff6133b8610d5085888a61555a565b60808051919092169052601160205251604081208181558501556133e0610d5084878961555a565b6040805167ffffffffffffffff9290921682526020820192909252a101613092565b610d50906134199267ffffffffffffffff9461555a565b7f46f5f12b0000000000000000000000000000000000000000000000000000000060805152166004526024608051fd5b346103455760805160031936011261034557602073ffffffffffffffffffffffffffffffffffffffff60095416604051908152f35b34610345576020600319360112610345577f66b1c1bdec8b60a3442bb25b5b6cd6fff3d0eceb6f5390be8e2f82a8ad39b23473ffffffffffffffffffffffffffffffffffffffff6134cd614fbd565b6134d5615e87565b61228f601054918381167fffffffffffffffffffffffff0000000000000000000000000000000000000000841617601055604051938493168390929173ffffffffffffffffffffffffffffffffffffffff60209181604085019616845216910152565b34610345576040600319360112610345576135516150e4565b6024359073ffffffffffffffffffffffffffffffffffffffff821682036103455767ffffffffffffffff90613584615e87565b16908160805152601160205260016040608051200190815460ff8160a01c16156136285782547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9283169081179093556040805191909216815260208101929092527f01efd4cd7dd64263689551000d4359d6559c839f39b773b1df3fd19ff060cf5f919081908101611779565b837f46f5f12b00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b3461034557613666366153ca565b613671929192615e87565b67ffffffffffffffff8216613693816000526006602052604060002054151590565b156136ae5750610d2f926136a891369161525f565b9061642e565b7f1e670e4b00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b346103455760805160031936011261034557602061036e61589b565b346103455761372161372961370d3661537c565b949161371a939193615e87565b3691615847565b923691615847565b7f00000000000000000000000000000000000000000000000000000000000000001561384b576080515b82518110156137c6578073ffffffffffffffffffffffffffffffffffffffff61377e6001938661572c565b51166137898161707b565b613795575b5001613753565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a18461378e565b506080515b8151811015610d2f578073ffffffffffffffffffffffffffffffffffffffff6137f66001938561572c565b511680156138455761380781617470565b613814575b505b016137cb565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a18361380c565b5061380e565b7f35f4a7b300000000000000000000000000000000000000000000000000000000608051526004608051fd5b346103455761191561246c61388b366152b4565b5050509150615740565b34610345576040600319360112610345576138ae6150e4565b60243567ffffffffffffffff8111610345576020916138d4612dd9923690600401615296565b90615618565b346103455760406003193601126103455760043567ffffffffffffffff81116103455780600401610100600319833603011261034557613918615091565b9060405161392581615174565b608051905261395661394c61394761394060c487018561556a565b369161525f565b616215565b6064850135616309565b916084840190613965826155f7565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691160361215457602485019277ffffffffffffffff000000000000000000000000000000006139cb85615516565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561205d5760805191613d83575b506120f957613a5884615516565b67ffffffffffffffff8116613a7a816000526006602052604060002054151590565b1561206a5750600480546040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff93909316918301919091523360248301526020908290604490829073ffffffffffffffffffffffffffffffffffffffff165afa90811561205d5760805191613d64575b5015611fca57613b0584615516565b90613b1b60a48801926138d4613940858561556a565b15613d1d575050613bf8611d5e60446020977ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09561ffff67ffffffffffffffff96161515600014613c6d5785613b7089615516565b1660805152600c8a52613b8c6040608051208a611d22846155f7565b7f41a8aa8df7945f0fb8ac5f7d88279638d9dc2ef9a6bf4ec9a53b80681b34aff786613bba611d5e8b615516565b6040805173ffffffffffffffffffffffffffffffffffffffff929092168252602082018d90529190931692a25b0194613bf2866155f7565b50615516565b604080517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff9081168252336020830152909216908201526060810185905292169180608081015b0390a280604051613c6481615174565b52604051908152f35b5084613c7888615516565b16806080515260078a527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c8980613cf060026040608051200173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391616dd1565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a2613be7565b613d27925061556a565b6123f36040519283927f24eb47e5000000000000000000000000000000000000000000000000000000008452602060048501526024840191615c01565b613d7d915060203d60201161214d5761213f81836151e4565b87613af6565b613d9c915060203d60201161214d5761213f81836151e4565b87613a4a565b346103455760805160031936011261034557602073ffffffffffffffffffffffffffffffffffffffff60105416604051908152f35b346103455760206003193601126103455760043567ffffffffffffffff8111610345578060040190610100600319823603011261034557604051613e1a81615174565b6080519052613e3f613e3561394761394060c485018661556a565b6064830135616309565b9060848101613e4d816155f7565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603612bee5750602481019277ffffffffffffffff00000000000000000000000000000000613eb485615516565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561205d57608051916142dd575b506120f957613f4184615516565b67ffffffffffffffff8116613f63816000526006602052604060002054151590565b1561206a5750600480546040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff93909316918301919091523360248301526020908290604490829073ffffffffffffffffffffffffffffffffffffffff165afa90811561205d57608051916142be575b5015611fca57613fee84615516565b9061400460a48401926138d4613940858561556a565b15613d1d57508291905067ffffffffffffffff61402085615516565b168060805152600760205261407660026040608051200173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016948591616dd1565b6040805173ffffffffffffffffffffffffffffffffffffffff85168152602081018690527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a267ffffffffffffffff6140d085615516565b1660805152601160205260406080512060ff600182015460a01c1660805150806000146142b65781545b80861161428257501561426d576141128482546155bb565b905561411d84615516565b505b6044017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16614162826155f7565b813b15610345576040517f69328dec0000000000000000000000000000000000000000000000000000000081526080805173ffffffffffffffffffffffffffffffffffffffff87811660048501526024840189905293909316604483015251909283916064918391905af1801561205d57614254575b5067ffffffffffffffff7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc091613c5485614216611d5e602099615516565b6040805173ffffffffffffffffffffffffffffffffffffffff9889168152336020820152979091169087015260608601529116929081906080820190565b608051614260916151e4565b60805161034557846141d8565b5061427a83600f546155bb565b600f5561411f565b85907fa17e11d500000000000000000000000000000000000000000000000000000000608051526004526024526044608051fd5b600f546140fa565b6142d7915060203d60201161214d5761213f81836151e4565b85613fdf565b6142f6915060203d60201161214d5761213f81836151e4565b85613f33565b346103455760a060031936011261034557614315614fbd565b5060243567ffffffffffffffff8116908190036103455760443567ffffffffffffffff81116103455760031960a09136030112610345576143546150a2565b5060843567ffffffffffffffff8111610345576143759036906004016150fb565b505060405161438381615129565b608051815260805160208201526080516040820152606060805191015260805152600e60205260806040815120604051906143bd82615129565b5463ffffffff808216928381528160208201818560201c16815260ff60606040850194848860401c168652019560601c161515855260405195865251166020850152511660408301525115156060820152f35b346103455760206003193601126103455767ffffffffffffffff6144326150e4565b16608051526011602052602060ff6001604060805120015460a01c166040519015158152f35b34610345576040600319360112610345576144716150e4565b60243567ffffffffffffffff82168060805152601160205260ff6001604060805120015460a01c161580156146b6575b61183757811561468a5773ffffffffffffffffffffffffffffffffffffffff6144c984615963565b1633036115585760805152601160205260406080512060ff600182015460a01c16600014614675576144fc8282546156b8565b90555b6040517f23b872dd000000000000000000000000000000000000000000000000000000006020820152336024820152306044820152606481018290527f00000000000000000000000000000000000000000000000000000000000000009061456e90610478816084810161044c565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690813b15610345576040517f47e7ef240000000000000000000000000000000000000000000000000000000081526080805173ffffffffffffffffffffffffffffffffffffffff9390931660048301526024820185905251909283916044918391905af1801561205d5761465c575b506040805167ffffffffffffffff9093168352602083019190915233917f569a440e6842b5e5a7ac02286311855f5a0b81b9390909e552e82aaf02c9e9bf9181908101611779565b608051614668916151e4565b6080516103455782614614565b5061468281600f546156b8565b600f556144ff565b7fa90c0d1900000000000000000000000000000000000000000000000000000000608051526004608051fd5b5080156144a1565b346103455760606003193601126103455760043561ffff811690819003610345576146e7615091565b9060443567ffffffffffffffff8111610345576147089036906004016150b3565b90614711615e87565b61ffff8416936127108510156147985783927f52aa194b292c8bfb5aaca8ee2000a965c3a051b306ff841873b16147526a39ba9592614787926040967fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000063ffff0000600a549360101b1692161717600a55615eef565b82519182526020820152a160805180f35b847f95f3517a00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b346103455760406003193601126103455760043567ffffffffffffffff8111610345573660238201121561034557806004013567ffffffffffffffff81116103455760248201916024369160a084020101116103455760243567ffffffffffffffff81116103455761483e903690600401615060565b919092614849615e87565b6080515b8281106148c5575050506080515b8181106148685760805180f35b8067ffffffffffffffff614882610d50600194868861555a565b168060805152600e602052608051604060805120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8608051608051a20161485b565b806148d6610d5060019386866154d7565b7f56f77aeff2def50c8b5f5a0df3bab7183df09bf36c6feba496bb42551db77d7060806149048488886154d7565b92614a146149da63ffffffff614a096149cd826149fe67ffffffffffffffff60208c0198169a8b8a5152600e60205260408a5120836149428b61552b565b169181549060408101937fffffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffff67ffffffff0000000061497f8761552b565b60201b16918f6cff0000000000000000000000007fffffffffffffffffffffffffffffffffffffffff000000000000000000000000916bffffffff0000000000000000606088019d8e61552b565b60401b1696019e8f61553c565b151560601b16951617161717179055826149f66040519a615549565b168952615549565b166020870152615549565b166040840152615466565b15156060820152a20161484d565b346103455760805160031936011261034557602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610345576020600319360112610345576020614a7c614fbd565b73ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b346103455760805160031936011261034557602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610345576080516003193601126103455760405161191590614b376060826151e4565b602481527f53696c6f65644c6f636b52656c65617365546f6b656e506f6f6c20312e362e3360208201527f2d646576000000000000000000000000000000000000000000000000000000006040820152604051918291602083526020830190615001565b3461034557602060031936011261034557614bb4614fbd565b614bbc615e87565b614bc461589b565b9081614bd05760805180f35b602073ffffffffffffffffffffffffffffffffffffffff7f738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b85959992614c8f6040517fa9059cbb0000000000000000000000000000000000000000000000000000000085820152614c698161044c8986602484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b7f0000000000000000000000000000000000000000000000000000000000000000616a5e565b6040519485521692a28080610d2f565b3461034557602060031936011261034557600435801561468a5773ffffffffffffffffffffffffffffffffffffffff614cd9608051615963565b1633036115585760808051805260116020525160409020600181015460a01c60ff168015614e845781545b808411614e50575015614e3b57614d1c8282546155bb565b90555b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803b15610345576040517f69328dec0000000000000000000000000000000000000000000000000000000081526080805173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016600484015260248301859052336044840152905191929091839160649183915af1801561205d57614e29575b506040519060805150608051825260208201527f58fca2457646a9f47422ab9eb9bff90cef88cd8b8725ab52b1d17baa392d784e60403392a260805180f35b608051614e35916151e4565b81614dea565b50614e4881600f546155bb565b600f55614d1f565b83907fa17e11d500000000000000000000000000000000000000000000000000000000608051526004526024526044608051fd5b600f54614d04565b3461034557602060031936011261034557600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361034557817ff208a58f0000000000000000000000000000000000000000000000000000000060209314908115614f93575b8115614f69575b8115614f3f575b8115614f15575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483614f0e565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150614f07565b7f1ef5498f0000000000000000000000000000000000000000000000000000000081149150614f00565b7faff2afbf0000000000000000000000000000000000000000000000000000000081149150614ef9565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036115aa57565b359073ffffffffffffffffffffffffffffffffffffffff821682036115aa57565b919082519283825260005b84811061504b5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b8060208092840101518282860101520161500c565b9181601f840112156115aa5782359167ffffffffffffffff83116115aa576020808501948460051b0101116115aa57565b6024359061ffff821682036115aa57565b6064359061ffff821682036115aa57565b9181601f840112156115aa5782359167ffffffffffffffff83116115aa5760208085019460e085020101116115aa57565b6004359067ffffffffffffffff821682036115aa57565b9181601f840112156115aa5782359167ffffffffffffffff83116115aa57602083818601950101116115aa57565b6080810190811067ffffffffffffffff82111761514557604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff82111761514557604052565b6060810190811067ffffffffffffffff82111761514557604052565b6040810190811067ffffffffffffffff82111761514557604052565b60a0810190811067ffffffffffffffff82111761514557604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761514557604052565b67ffffffffffffffff811161514557601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b92919261526b82615225565b9161527960405193846151e4565b8294818452818301116115aa578281602093846000960137010152565b9080601f830112156115aa578160206152b19335910161525f565b90565b60a06003198201126115aa5760043573ffffffffffffffffffffffffffffffffffffffff811681036115aa579160243567ffffffffffffffff811681036115aa57916044359160643561ffff811681036115aa57916084359067ffffffffffffffff82116115aa57615328916004016150fb565b9091565b602060408183019282815284518094520192019060005b8181106153505750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101615343565b60406003198201126115aa5760043567ffffffffffffffff81116115aa57816153a791600401615060565b929092916024359067ffffffffffffffff82116115aa5761532891600401615060565b9060406003198301126115aa5760043567ffffffffffffffff811681036115aa57916024359067ffffffffffffffff82116115aa57615328916004016150fb565b9181601f840112156115aa5782359167ffffffffffffffff83116115aa57602080850194606085020101116115aa57565b6152b19160206154558351604084526040840190615001565b920151906020818403910152615001565b359081151582036115aa57565b35906fffffffffffffffffffffffffffffffff821682036115aa57565b91908260609103126115aa576040516154a881615190565b60406154d28183956154b981615466565b85526154c760208201615473565b602086015201615473565b910152565b91908110156154e75760a0020190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3567ffffffffffffffff811681036115aa5790565b3563ffffffff811681036115aa5790565b3580151581036115aa5790565b359063ffffffff821682036115aa57565b91908110156154e75760051b0190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156115aa570180359067ffffffffffffffff82116115aa576020019181360383136115aa57565b919082039182116155c857565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b3573ffffffffffffffffffffffffffffffffffffffff811681036115aa5790565b9067ffffffffffffffff6152b192166000526007602052600560406000200190602081519101209060019160005201602052604060002054151590565b906040519182815491828252602082019060005260206000209260005b818110615689575050615687925003836151e4565b565b845473ffffffffffffffffffffffffffffffffffffffff16835260019485019487945060209093019201615672565b919082018092116155c857565b67ffffffffffffffff81116151455760051b60200190565b906156e7826156c5565b6156f460405191826151e4565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061572282946156c5565b0190602036910137565b80518210156154e75760209160051b010190565b67ffffffffffffffff16600052600d602052604060002061576081615655565b9160045460a01c801515918261583c575b505061577b575090565b60016157879101615655565b908151806157955750905090565b6157a36157a89183516156b8565b6156dd565b9260005b82518110156157ea578073ffffffffffffffffffffffffffffffffffffffff6157d76001938661572c565b51166157e3828861572c565b52016157ac565b509160005b8151811015615837578073ffffffffffffffffffffffffffffffffffffffff61581a6001938561572c565b511661583061582a8387516156b8565b8861572c565b52016157ef565b505050565b101590503880615771565b929190615853816156c5565b9361586160405195866151e4565b602085838152019160051b81019283116115aa57905b82821061588357505050565b6020809161589084614fe0565b815201910190615877565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561594757600091615918575090565b90506020813d60201161593f575b81615933602093836151e4565b810103126115aa575190565b3d9150615926565b6040513d6000823e3d90fd5b91908110156154e75760061b0190565b67ffffffffffffffff16600052601160205260016040600020015460ff8160a01c166159a6575073ffffffffffffffffffffffffffffffffffffffff6010541690565b73ffffffffffffffffffffffffffffffffffffffff1690565b91908110156154e7576060020190565b604051906159dc826151ac565b60606020838281520152565b67ffffffffffffffff16600052600d6020526040600020906003615a0e60028401615655565b920190615a1a82615655565b5060045460a01c8015159182615ae3575b5050615a35575090565b615a3e90615655565b90815180615a4c5750905090565b6157a3615a5a9183516156b8565b9260005b8251811015615a9c578073ffffffffffffffffffffffffffffffffffffffff615a896001938661572c565b5116615a95828861572c565b5201615a5e565b509160005b8151811015615837578073ffffffffffffffffffffffffffffffffffffffff615acc6001938561572c565b5116615adc61582a8387516156b8565b5201615aa1565b101590503880615a2b565b90600182811c92168015615b37575b6020831014615b0857565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691615afd565b9060405191826000825492615b5584615aee565b8084529360018116908115615bc15750600114615b7a575b50615687925003836151e4565b90506000929192526020600020906000915b818310615ba55750509060206156879282010138615b6d565b6020919350806001915483858901015201910190918492615b8c565b602093506156879592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138615b6d565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b60405190615c4d826151c8565b60006080838281528260208201528260408201528260608201520152565b90604051615c78816151c8565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff1660005260076020526152b16004604060002001615b41565b91908110156154e75760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61813603018212156115aa570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156115aa570180359067ffffffffffffffff82116115aa57602001918160051b360383136115aa57565b818102929181159184041417156155c857565b818110615da5575050565b60008155600101615d9a565b9160209082815201919060005b818110615dcb5750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff615df288614fe0565b168152019401929101615dbe565b67ffffffffffffffff16615e21816000526006602052604060002054151590565b15615e5a5780600052601160205260ff60016040600020015460a01c16615e495750600f5490565b600052601160205260406000205490565b7fd9a9cd680000000000000000000000000000000000000000000000000000000060005260045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff600154163303615ea857565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b356fffffffffffffffffffffffffffffffff811681036115aa5790565b9160005b8281101561620f5760e0810284016000615f0c82615516565b9067ffffffffffffffff821691615f30836000526006602052604060002054151590565b156161e357615ff99260408593615fa4615f9e94615f9e615f64602060019c9b0192611cfe615f5f3686615490565b616917565b91825463ffffffff8160801c161590816161c5575b816161b6575b8161619b575b8161618c575b508061617d575b6160f2575b3690615490565b90616b9e565b6080850192615fb6615f5f3686615490565b8152600c6020522092835463ffffffff8160801c161590816160d4575b816160c5575b816160aa575b8161609b575b508061608c575b615fff575b503690615490565b01615ef3565b61601c60a06fffffffffffffffffffffffffffffffff9201615ed2565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff161717835538615ff1565b506160968261553c565b615fec565b60ff915060a01c161538615fe5565b6fffffffffffffffffffffffffffffffff8116159150615fdf565b8589015460801c159150615fd9565b858901546fffffffffffffffffffffffffffffffff16159150615fd3565b6fffffffffffffffffffffffffffffffff61610e878b01615ed2565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff1617178355615f97565b506161878161553c565b615f92565b60ff915060a01c161538615f8b565b6fffffffffffffffffffffffffffffffff8116159150615f85565b848e015460801c159150615f7f565b848e01546fffffffffffffffffffffffffffffffff16159150615f79565b506024917f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50915050565b80518015616285576020036162475780516020828101918301839003126115aa57519060ff8211616247575060ff1690565b6123f3906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190615001565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff82116155c857565b60ff16604d81116155c857600a0a90565b81156162da570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff81169282841461640f578284116163e5579061634e916162ab565b91604d60ff84161180156163ac575b616376575050906163706152b1926162bf565b90615d87565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b506163b6836162bf565b80156162da577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04841161635d565b6163ee916162ab565b91604d60ff841611616376575050906164096152b1926162bf565b906162d0565b5050505090565b908160209103126115aa575180151581036115aa5790565b9080511561666e5767ffffffffffffffff8151602083012092169182600052600760205261646381600560406000200161752a565b1561662a5760005260086020526040600020815167ffffffffffffffff8111615145576164908254615aee565b601f81116165f8575b506020601f8211600114616532579161650c827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea959361652295600091616527575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055604051918291602083526020830190615001565b0390a2565b9050840151386164db565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b8181106165e05750926165229492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9896106165a9575b5050811b019055611a70565b8501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388061659d565b9192602060018192868a015181550194019201616562565b61662490836000526020600020601f840160051c81019160208510610c5b57601f0160051c0190615d9a565b38616499565b50906123f36040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190615001565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b67ffffffffffffffff16600081815260066020526040902054909291901561679a579161679760e092616763856166ef7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b97616917565b846000526007602052616706816040600020616b9e565b61670f83616917565b846000526007602052616729836002604060002001616b9e565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6167d0615c40565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff808351169161682d602085019361682761681a63ffffffff875116426155bb565b8560808901511690615d87565b906156b8565b8082101561684657505b16825263ffffffff4216905290565b9050616837565b805160005b81811061685e57505050565b600181018082116155c8575b82811061687a5750600101616852565b73ffffffffffffffffffffffffffffffffffffffff616899838661572c565b511673ffffffffffffffffffffffffffffffffffffffff6168ba838761572c565b5116146168c95760010161686a565b73ffffffffffffffffffffffffffffffffffffffff6168e8838661572c565b51167fa1726e400000000000000000000000000000000000000000000000000000000060005260045260246000fd5b8051156169b7576fffffffffffffffffffffffffffffffff6040820151166fffffffffffffffffffffffffffffffff602083015116106169545750565b6064906169b5604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff60408201511615801590616a3f575b6169de5750565b6064906169b5604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff60208201511615156169d7565b73ffffffffffffffffffffffffffffffffffffffff616aed911691604092600080855193616a8c87866151e4565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602086015260208151910182855af13d15616b96573d91616ad183615225565b92616ade875194856151e4565b83523d6000602085013e61757f565b80519081616afa57505050565b602080616b0b938301019101616416565b15616b135750565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b60609161757f565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1991616cd76060928054616bdb63ffffffff8260801c16426155bb565b9081616d16575b50506fffffffffffffffffffffffffffffffff6001816020860151169282815416808510600014616d0e57508280855b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416178155616c8b8651151582907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b60408601517fffffffffffffffffffffffffffffffff0000000000000000000000000000000060809190911b16939092166fffffffffffffffffffffffffffffffff1692909217910155565b61679760405180926fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b838091616c12565b6fffffffffffffffffffffffffffffffff91616d4b839283616d446001880154948286169560801c90615d87565b91166156b8565b80821015616dca57505b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff9290911692909216167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116174260801b73ffffffff00000000000000000000000000000000161781553880616be2565b9050616d55565b9182549060ff8260a01c16158015617010575b61700a576fffffffffffffffffffffffffffffffff82169160018501908154616e2963ffffffff6fffffffffffffffffffffffffffffffff83169360801c16426155bb565b9081616f6c575b5050848110616f205750838310616e8a575050616e5f6fffffffffffffffffffffffffffffffff9283926155bb565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b5460801c91616e9981856155bb565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101908082116155c857616ee7616eec9273ffffffffffffffffffffffffffffffffffffffff966156b8565b6162d0565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828573ffffffffffffffffffffffffffffffffffffffff927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611616fe057616f87926168279160801c90615d87565b80841015616fdb5750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178655923880616e30565b616f92565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b508215616de4565b906040519182815491828252602082019060005260206000209260005b81811061704a575050615687925003836151e4565b8454835260019485019487945060209093019201617035565b80548210156154e75760005260206000200190600090565b600081815260036020526040902054801561720a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116155c857600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116155c85781810361719b575b505050600254801561716c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01617129816002617063565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6171f26171ac6171bd936002617063565b90549060031b1c9283926002617063565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905560005260036020526040600020553880806170f0565b5050600090565b600081815260066020526040902054801561720a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116155c857600554907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116155c857818103617302575b505050600554801561716c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016172bf816005617063565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600555600052600660205260006040812055600190565b6173246173136171bd936005617063565b90549060031b1c9283926005617063565b90556000526006602052604060002055388080617286565b9060018201918160005282602052604060002054801515600014617467577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116155c8578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116155c857818103617430575b5050508054801561716c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906173f18282617063565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b6174506174406171bd9386617063565b90549060031b1c92839286617063565b9055600052836020526040600020553880806173b9565b50505050600090565b806000526003602052604060002054156000146174ca5760025468010000000000000000811015615145576174b16171bd8260018594016002556002617063565b9055600254906000526003602052604060002055600190565b50600090565b806000526006602052604060002054156000146174ca5760055468010000000000000000811015615145576175116171bd8260018594016005556005617063565b9055600554906000526006602052604060002055600190565b600082815260018201602052604090205461720a578054906801000000000000000082101561514557826175686171bd846001809601855584617063565b905580549260005201602052604060002055600190565b919290156175fa5750815115617593575090565b3b1561759c5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b82519091501561760d5750805190602001fd5b6123f3906040519182917f08c379a000000000000000000000000000000000000000000000000000000000835260206004840152602483019061500156fea164736f6c634300081a000a",
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetRequiredInboundCCVs(opts *bind.CallOpts, arg0 common.Address, sourceChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getRequiredInboundCCVs", arg0, sourceChainSelector, amount, arg3, arg4)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetRequiredInboundCCVs(arg0 common.Address, sourceChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetRequiredInboundCCVs(&_SiloedLockReleaseTokenPool.CallOpts, arg0, sourceChainSelector, amount, arg3, arg4)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetRequiredInboundCCVs(arg0 common.Address, sourceChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetRequiredInboundCCVs(&_SiloedLockReleaseTokenPool.CallOpts, arg0, sourceChainSelector, amount, arg3, arg4)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetRequiredOutboundCCVs(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getRequiredOutboundCCVs", arg0, destChainSelector, amount, arg3, arg4)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetRequiredOutboundCCVs(arg0 common.Address, destChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetRequiredOutboundCCVs(&_SiloedLockReleaseTokenPool.CallOpts, arg0, destChainSelector, amount, arg3, arg4)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetRequiredOutboundCCVs(arg0 common.Address, destChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetRequiredOutboundCCVs(&_SiloedLockReleaseTokenPool.CallOpts, arg0, destChainSelector, amount, arg3, arg4)
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) ApplyFinalityConfigUpdates(opts *bind.TransactOpts, finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "applyFinalityConfigUpdates", finalityThreshold, customFinalityTransferFeeBps, rateLimitConfigArgs)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) ApplyFinalityConfigUpdates(finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ApplyFinalityConfigUpdates(&_SiloedLockReleaseTokenPool.TransactOpts, finalityThreshold, customFinalityTransferFeeBps, rateLimitConfigArgs)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) ApplyFinalityConfigUpdates(finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ApplyFinalityConfigUpdates(&_SiloedLockReleaseTokenPool.TransactOpts, finalityThreshold, customFinalityTransferFeeBps, rateLimitConfigArgs)
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) SetCustomFinalityRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "setCustomFinalityRateLimitConfig", rateLimitConfigArgs)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) SetCustomFinalityRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetCustomFinalityRateLimitConfig(&_SiloedLockReleaseTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) SetCustomFinalityRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetCustomFinalityRateLimitConfig(&_SiloedLockReleaseTokenPool.TransactOpts, rateLimitConfigArgs)
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
	RemoteChainSelector    uint64
	OutboundCCVs           []common.Address
	AdditionalOutboundCCVs []common.Address
	InboundCCVs            []common.Address
	AdditionalInboundCCVs  []common.Address
	Raw                    types.Log
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

type SiloedLockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumedIterator struct {
	Event *SiloedLockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(SiloedLockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *SiloedLockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterCustomFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "CustomFinalityOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "CustomFinalityOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchCustomFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "CustomFinalityOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumed)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "CustomFinalityOutboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseCustomFinalityOutboundRateLimitConsumed(log types.Log) (*SiloedLockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumed, error) {
	event := new(SiloedLockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumed)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "CustomFinalityOutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator struct {
	Event *SiloedLockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(SiloedLockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *SiloedLockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterCustomFinalityTransferInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "CustomFinalityTransferInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "CustomFinalityTransferInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchCustomFinalityTransferInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "CustomFinalityTransferInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "CustomFinalityTransferInboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseCustomFinalityTransferInboundRateLimitConsumed(log types.Log) (*SiloedLockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumed, error) {
	event := new(SiloedLockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "CustomFinalityTransferInboundRateLimitConsumed", log); err != nil {
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

type SiloedLockReleaseTokenPoolFinalityConfigUpdatedIterator struct {
	Event *SiloedLockReleaseTokenPoolFinalityConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolFinalityConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolFinalityConfigUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(SiloedLockReleaseTokenPoolFinalityConfigUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *SiloedLockReleaseTokenPoolFinalityConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolFinalityConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolFinalityConfigUpdated struct {
	FinalityConfig               uint16
	CustomFinalityTransferFeeBps uint16
	Raw                          types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterFinalityConfigUpdated(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolFinalityConfigUpdatedIterator, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "FinalityConfigUpdated")
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolFinalityConfigUpdatedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "FinalityConfigUpdated", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchFinalityConfigUpdated(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolFinalityConfigUpdated) (event.Subscription, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "FinalityConfigUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolFinalityConfigUpdated)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "FinalityConfigUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseFinalityConfigUpdated(log types.Log) (*SiloedLockReleaseTokenPoolFinalityConfigUpdated, error) {
	event := new(SiloedLockReleaseTokenPoolFinalityConfigUpdated)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "FinalityConfigUpdated", log); err != nil {
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

func (SiloedLockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b2")
}

func (SiloedLockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x41a8aa8df7945f0fb8ac5f7d88279638d9dc2ef9a6bf4ec9a53b80681b34aff7")
}

func (SiloedLockReleaseTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x39b29a1c46dd0e3ee5683fed1d66e11a234938cde48e5c4e14389be508348eb6")
}

func (SiloedLockReleaseTokenPoolFinalityConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x52aa194b292c8bfb5aaca8ee2000a965c3a051b306ff841873b16147526a39ba")
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
	return common.HexToHash("0x56f77aeff2def50c8b5f5a0df3bab7183df09bf36c6feba496bb42551db77d70")
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

	GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	GetRebalancer(opts *bind.CallOpts) (common.Address, error)

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRequiredInboundCCVs(opts *bind.CallOpts, arg0 common.Address, sourceChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error)

	GetRequiredOutboundCCVs(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error)

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

	ApplyFinalityConfigUpdates(opts *bind.TransactOpts, finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error)

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

	SetCustomFinalityRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error)

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

	FilterCustomFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumedIterator, error)

	WatchCustomFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomFinalityOutboundRateLimitConsumed(log types.Log) (*SiloedLockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumed, error)

	FilterCustomFinalityTransferInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator, error)

	WatchCustomFinalityTransferInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomFinalityTransferInboundRateLimitConsumed(log types.Log) (*SiloedLockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumed, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*SiloedLockReleaseTokenPoolDynamicConfigSet, error)

	FilterFinalityConfigUpdated(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolFinalityConfigUpdatedIterator, error)

	WatchFinalityConfigUpdated(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolFinalityConfigUpdated) (event.Subscription, error)

	ParseFinalityConfigUpdated(log types.Log) (*SiloedLockReleaseTokenPoolFinalityConfigUpdated, error)

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
