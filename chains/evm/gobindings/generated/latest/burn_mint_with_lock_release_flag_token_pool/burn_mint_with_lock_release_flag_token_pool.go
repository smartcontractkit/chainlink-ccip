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

var BurnMintWithLockReleaseFlagTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIBurnMintERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"additionalOutboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"additionalInboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyFinalityConfigUpdates\",\"inputs\":[{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"destToUseDefaultFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAccumulatedFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredInboundCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredOutboundCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustomFinalityRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"additionalOutboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"additionalInboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityTransferInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalityConfigUpdated\",\"inputs\":[{\"name\":\"finalityConfig\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeeWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidFinality\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidFinalityConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenTransferFeeConfigNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x61010080604052346103c557616afa803803809161001d8285610444565b833981019060a0818303126103c55780516001600160a01b038116908190036103c55761004c60208301610467565b60408301519091906001600160401b0381116103c55783019380601f860112156103c5578451946001600160401b03861161042e578560051b9060208201966100986040519889610444565b87526020808801928201019283116103c557602001905b828210610416575050506100d160806100ca60608601610475565b9401610475565b92331561040557600180546001600160a01b03191633179055811580156103f4575b80156103e3575b6103d2578160209160049360805260c0526040519283809263313ce56760e01b82525afa60009181610391575b50610366575b5060a052600480546001600160a01b0319166001600160a01b03929092169190911790558051151560e0819052610248575b6040516164d0908161062a8239608051818181611533015281816117de015281816119990152818161209a015281816123260152818161247501528181612cee01528181612f43015281816130e4015281816132380152818161341101528181613a6401528181613ad301528181613bf9015281816147e901526156ed015260a05181818161186701528181613a200152818161508701528181615190015261531a015260c051818181610c45015281816115ce0152818161213501528181612d8901526132d5015260e051818181610bf2015281816116130152818161217a0152612a9b0152f35b60405160206102578183610444565b60008252600036813760e051156103555760005b82518110156102d2576001906001600160a01b036102898286610489565b511683610295826104cb565b6102a2575b50500161026b565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a1388361029a565b50905060005b825181101561034c576001906001600160a01b036102f68286610489565b511680156103465783610308826105c9565b610316575b50505b016102d8565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a1388361030d565b50610310565b5050503861015f565b6335f4a7b360e01b60005260046000fd5b60ff1660ff821681810361037a575061012d565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d6020116103ca575b816103ad60209383610444565b810103126103c5576103be90610467565b9038610127565b600080fd5b3d91506103a0565b630a64406560e11b60005260046000fd5b506001600160a01b038116156100fa565b506001600160a01b038416156100f3565b639b15e16f60e01b60005260046000fd5b6020809161042384610475565b8152019101906100af565b634e487b7160e01b600052604160045260246000fd5b601f909101601f19168101906001600160401b0382119082101761042e57604052565b519060ff821682036103c557565b51906001600160a01b03821682036103c557565b805182101561049d5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b805482101561049d5760005260206000200190600090565b60008181526003602052604090205480156105c25760001981018181116105ac576002546000198101919082116105ac5781810361055b575b5050506002548015610545576000190161051f8160026104b3565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61059461056c61057d9360026104b3565b90549060031b1c92839260026104b3565b819391549060031b91821b91600019901b19161790565b90556000526003602052604060002055388080610504565b634e487b7160e01b600052601160045260246000fd5b5050600090565b80600052600360205260406000205415600014610623576002546801000000000000000081101561042e5761060a61057d82600185940160025560026104b3565b9055600254906000526003602052604060002055600190565b5060009056fe60a080604052600436101561001357600080fd5b60006080526080513560e01c90816301ffc9a714613d9e57508063164e68de14613b59578063181f5a7714613af757806321df0da714613aa5578063240028e814613a4457806324f65ee714613a055780632a10097b146137ab5780632c286daf1461369e57806337b192471461358a57806339077537146131be578063489a68f214612c4a5780634c5ef0ed14612c055780634f71592c14612be757806354c8a4f314612a695780635df45a3714612a4557806362ddd3c4146129c05780636d3d1a581461298b5780637437ff9f1461294b57806379ba5097146128745780637d54534e146127e3578063804ba5a9146127a15780638926f54f1461275c5780638da5cb5b14612727578063962d4020146125ca5780639a4575b91461203b5780639f68f6731461201d578063a42a7b8b14611ea8578063a7cd63b714611e3a578063acfecf9114611d0d578063af58d59f14611cc1578063b184b94214611bd5578063b1c71c65146114a9578063b794658014611471578063c4bffe2b14611337578063c75eea9c1461128c578063cf7401f3146110e5578063d966866b14610c69578063dc0bd97114610c17578063e0351e1314610bd9578063e8a1da17146102c75763f2fde38b146101e857600080fd5b346102c15760206003193601126102c15773ffffffffffffffffffffffffffffffffffffffff610216613ecf565b61021e614c99565b163381146102955760805180547fffffffffffffffffffffffff0000000000000000000000000000000000000000168217815560015473ffffffffffffffffffffffffffffffffffffffff16907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789080a360805180f35b7fdad89dca00000000000000000000000000000000000000000000000000000000608051526004608051fd5b60805180fd5b346102c1576102d53661428e565b9190926102e0614c99565b608051905b828210610a17575050506080519163ffffffff4216917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b81811015610a1157600581901b830135858112156102c1578301610120813603126102c1576040519461035486613f96565b813567ffffffffffffffff81168103610a0c578652602082013567ffffffffffffffff81116102c15782019436601f870112156102c1578535610396816145ca565b966103a46040519889613fce565b81885260208089019260051b820101903682116102c15760208101925b8284106109dd575050505060208701958652604083013567ffffffffffffffff81116102c1576103f490369085016141a8565b916040880192835261041e61040c36606087016143a2565b9460608a0195865260c03691016143a2565b956080890196875261043085516158cf565b61043a87516158cf565b835151156109b15761045667ffffffffffffffff8a5116616348565b156109765767ffffffffffffffff89511660805152600760205260406080512061059a86516fffffffffffffffffffffffffffffffff604082015116906105556fffffffffffffffffffffffffffffffff602083015116915115158360806040516104c081613f96565b858152602081018b905260408101849052606081018690520152855474ff000000000000000000000000000000000000000091151560a01b919091167fffffffffffffffffffffff0000000000000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff84161773ffffffff0000000000000000000000000000000060808a901b1617178555565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176001830155565b6106c088516fffffffffffffffffffffffffffffffff6040820151169061067b6fffffffffffffffffffffffffffffffff602083015116915115158360806040516105e481613f96565b858152602081018b9052604081018490526060810186905201526002860180547fffffffffffffffffffffff000000000000000000000000000000000000000000166fffffffffffffffffffffffffffffffff85161773ffffffff0000000000000000000000000000000060808b901b161791151560a01b74ff000000000000000000000000000000000000000016919091179055565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176003830155565b6004855191019080519067ffffffffffffffff8211610945576106e38354614987565b601f8111610906575b506020906001601f8411146108605791809161073d9360805192610855575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b6080515b88518051821015610779579061077360019261076c8367ffffffffffffffff8f511692614631565b519061533c565b01610744565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c293919997509561084767ffffffffffffffff60019796949851169251935191516108136107de60405196879687526101006020880152610100870190614049565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a1019392909193610322565b015190508e8061070b565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083169184608051528160805120926080515b8181106108ee57509084600195949392106108b7575b505050811b019055610740565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558d80806108aa565b92936020600181928786015181550195019301610894565b610935908460805152602060805120601f850160051c8101916020861061093b575b601f0160051c0190614c33565b8d6106ec565b9091508190610928565b7f4e487b71000000000000000000000000000000000000000000000000000000006080515260416004526024608051fd5b67ffffffffffffffff8951167f1d5ad3c500000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b7f14c880ca00000000000000000000000000000000000000000000000000000000608051526004608051fd5b833567ffffffffffffffff81116102c157602091610a0183928336918701016141a8565b8152019301926103c1565b600080fd5b60805180f35b9092919367ffffffffffffffff610a37610a3286888661446c565b614428565b1692610a4284616089565b15610ba95783608051526007602052610a62600560406080512001615e90565b926080515b8451811015610aa15760019086608051526007602052610a9a600560406080512001610a938389614631565b51906161b4565b5001610a67565b5093909491959250806080515260076020526005604060805120608051815560805160018201556080516002820155608051600382015560048101610ae68154614987565b80610b59575b505001805490608051815581610b35575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a10190916102e5565b60805152602060805120908101905b81811015610afd576080518155600101610b44565b601f8111600114610b72575060805190555b8880610aec565b610b929082608051526001601f6020608051209201861c82019101614c33565b608080518290525160208120918190559055610b6b565b837f1e670e4b00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b346102c1576080516003193601126102c15760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b346102c1576080516003193601126102c157602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102c15760206003193601126102c15760043567ffffffffffffffff81116102c157610c9a9036906004016140a8565b610ca2614c99565b608051915b818310610cb45760805180f35b610cc2610a32848484614b8c565b610cda610cd0858585614b8c565b6020810190614bcc565b9091610cf4610cea878787614b8c565b6040810190614bcc565b90610d0d610d03898989614b8c565b6060810190614bcc565b9091610d27610d1d8b8b8b614b8c565b6080810190614bcc565b949097610d3d610d38368a8461474c565b615805565b610d4b610d3836848661474c565b610d59610d3836868861474c565b610d67610d3836888c61474c565b604051610d7381613f13565b610d7e368a8461474c565b8152610d8b36848661474c565b6020820152610d9b36868861474c565b6040820152610dab36888c61474c565b606082015267ffffffffffffffff881660805152600d602052604060805120815180519067ffffffffffffffff8211610945576801000000000000000082116109455760209083548385558084106110c6575b500182608051526020608051206080515b83811061109c5750505050602082015180519067ffffffffffffffff82116109455768010000000000000000821161094557602090600184015483600186015580841061107a575b500160018301608051526020608051206080515b8381106110505750505050604082015180519067ffffffffffffffff82116109455768010000000000000000821161094557602090600284015483600286015580841061102e575b500160028301608051526020608051206080515b838110611004575050505060036060919e9c9d9e019101519081519167ffffffffffffffff831161094557680100000000000000008311610945576020908254848455808510610fe5575b500190608051526020608051206080515b838110610fbb5750505050610fa0608095610fb095610f927fece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d9a9660019e9d9c9a96610f8467ffffffffffffffff976040519d8d8f9e8f9081520191614c4a565b918b830360208d0152614c4a565b9188830360408a0152614c4a565b9285840360608701521696614c4a565b0390a2019190610ca7565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501610f23565b610ffe9084608051528584608051209182019101614c33565b38610f12565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501610ec7565b61104a9060028601608051528484608051209182019101614c33565b38610eb3565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501610e6b565b6110969060018601608051528484608051209182019101614c33565b38610e57565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501610e0f565b6110df9085608051528484608051209182019101614c33565b38610dfe565b346102c15760e06003193601126102c1576110fe61412c565b6060367fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc01126102c15760405161113481613fb2565b60243580151581036102c15781526044356fffffffffffffffffffffffffffffffff811681036102c15760208201526064356fffffffffffffffffffffffffffffffff811681036102c15760408201526060367fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c01126102c157604051906111bb82613fb2565b60843580151581036102c157825260a4356fffffffffffffffffffffffffffffffff811681036102c157602083015260c4356fffffffffffffffffffffffffffffffff811681036102c157604083015273ffffffffffffffffffffffffffffffffffffffff600954163314158061126a575b61123a57610a11926155a6565b7f8e4a23d60000000000000000000000000000000000000000000000000000000060805152336004526024608051fd5b5073ffffffffffffffffffffffffffffffffffffffff6001541633141561122d565b346102c15760206003193601126102c15767ffffffffffffffff6112ae61412c565b6112b6614ad9565b50166080515260076020526113336112da6112d5604060805120614b04565b615780565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b0390f35b346102c1576080516003193601126102c157608051506040516005548082528160208101600560805152602060805120926080515b81811061145857505061138192500382613fce565b8051906113a6611390836145ca565b9261139e6040519485613fce565b8084526145ca565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06020840192013683376080515b8151811015611407578067ffffffffffffffff6113f460019385614631565b51166114008287614631565b52016113d5565b505090604051918291602083019060208452518091526040830191906080515b818110611435575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101611427565b845483526001948501948694506020909301920161136c565b346102c15760206003193601126102c15761133361149561149061412c565b614b6a565b604051918291602083526020830190614049565b346102c15760606003193601126102c15760043567ffffffffffffffff81116102c15760a060031982360301126102c1576114e26140d9565b9060443567ffffffffffffffff81116102c1576115039036906004016141a8565b5061150c614868565b50608481019061151b8261447c565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611b8757602481019077ffffffffffffffff0000000000000000000000000000000061158183614428565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611a905760805191611b58575b50611b2c576116116044820161447c565b7f0000000000000000000000000000000000000000000000000000000000000000611acc575b5067ffffffffffffffff61164a83614428565b16611662816000526006602052604060002054151590565b15611a9d57602073ffffffffffffffffffffffffffffffffffffffff60045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015611a905760805190611a2d575b73ffffffffffffffffffffffffffffffffffffffff91501633036119fd57606461ffff910135931692831515928380946119ee575b156119335761ffff600a5416948581106118ff57506118c3945061174d61173d61172385614428565b67ffffffffffffffff16600052600b602052604060002090565b836117478461447c565b91615c49565b7f7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b267ffffffffffffffff61178961178386614428565b9361447c565b6040805173ffffffffffffffffffffffffffffffffffffffff929092168252602082018690529190931692a25b9182906118cd575b5061149061185f916117cf846156d6565b6117d881614428565b604080517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16815233602082015290810186905267ffffffffffffffff91909116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae109080606081015b0390a2614428565b9060405160ff7f00000000000000000000000000000000000000000000000000000000000000001660208201526020815261189b604082613fce565b604051926118a884613f7a565b8352602083015260405192839260408452604084019061434e565b9060208301520390f35b61185f9192506118f7611490916127106118f061ffff600a5460101c1683614c20565b0490615773565b9291506117be565b85907fe08f03ef00000000000000000000000000000000000000000000000000000000608051526004526024526044608051fd5b506118c3935067ffffffffffffffff61194b83614428565b16806080515260076020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894482806119c160406080512073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615c49565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a26117b6565b5061ffff600a541615156116fa565b7f728fe07b0000000000000000000000000000000000000000000000000000000060805152336004526024608051fd5b506020813d602011611a88575b81611a4760209383613fce565b810103126102c1575173ffffffffffffffffffffffffffffffffffffffff811681036102c15773ffffffffffffffffffffffffffffffffffffffff906116c5565b3d9150611a3a565b6040513d608051823e3d90fd5b7fa9902c7e00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b73ffffffffffffffffffffffffffffffffffffffff16611af9816000526003602052604060002054151590565b611637577fd0d2597600000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b7f53ad11d800000000000000000000000000000000000000000000000000000000608051526004608051fd5b611b7a915060203d602011611b80575b611b728183613fce565b81019061528e565b85611600565b503d611b68565b73ffffffffffffffffffffffffffffffffffffffff611ba58361447c565b7f961c9a4f0000000000000000000000000000000000000000000000000000000060805152166004526024608051fd5b346102c15760406003193601126102c157611bee613ecf565b6024356bffffffffffffffffffffffff811681036102c157611c0e614c99565b73ffffffffffffffffffffffffffffffffffffffff82169182156109b1577f39b29a1c46dd0e3ee5683fed1d66e11a234938cde48e5c4e14389be508348eb6927fffffffffffffffffffffffff00000000000000000000000000000000000000008360a01b1617600455611cb8604051928392839092916bffffffffffffffffffffffff60209173ffffffffffffffffffffffffffffffffffffffff604085019616845216910152565b0390a160805180f35b346102c15760206003193601126102c15767ffffffffffffffff611ce361412c565b611ceb614ad9565b50166080515260076020526113336112da6112d5600260406080512001614b04565b346102c15767ffffffffffffffff611d24366142dc565b929091611d2f614c99565b1690611d48826000526006602052604060002054151590565b15611e0a5781608051526007602052611d7b600560406080512001611d6e368685614171565b60208151910120906161b4565b15611dc3577f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d769192611dba604051928392602084526020840191614a9a565b0390a260805180f35b611e06906040519384937f74f23c7c0000000000000000000000000000000000000000000000000000000085526004850152604060248501526044840191614a9a565b0390fd5b507f1e670e4b00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b346102c1576080516003193601126102c157608051506040516002548082526020820190600260805152602060805120906080515b818110611e925761133385611e8681870382613fce565b6040519182918261423e565b8254845260209093019260019283019201611e6f565b346102c15760206003193601126102c15767ffffffffffffffff611eca61412c565b16608051526007602052611ee5600560406080512001615e90565b8051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611f2b611f15846145ca565b93611f236040519586613fce565b8085526145ca565b016080515b81811061200c5750506080515b8151811015611f875780611f5360019284614631565b51608051526008602052611f6b6040608051206149da565b611f758286614631565b52611f808185614631565b5001611f3d565b826040518091602082016020835281518091526040830190602060408260051b860101930191608051905b828210611fc157505050500390f35b91936020611ffc827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851614049565b9601920192018594939192611fb2565b806060602080938701015201611f30565b346102c157611333611e86612031366141c6565b5050509150614881565b346102c15760206003193601126102c15760043567ffffffffffffffff81116102c15760a060031982360301126102c157612074614868565b50608481016120828161447c565b73ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116036125ac57602482019177ffffffffffffffff000000000000000000000000000000006120e884614428565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611a90576080519161258d575b50611b2c576121786044820161447c565b7f000000000000000000000000000000000000000000000000000000000000000061252d575b5067ffffffffffffffff6121b184614428565b166121c9816000526006602052604060002054151590565b15611a9d57602073ffffffffffffffffffffffffffffffffffffffff60045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015611a9057608051906124ca575b73ffffffffffffffffffffffffffffffffffffffff91501633036119fd57606401359060805160001461240b5761ffff600a5416806123d6575082612376926114909261228161173d61172361133398614428565b7f7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b267ffffffffffffffff6122b761178386614428565b6040805173ffffffffffffffffffffffffffffffffffffffff929092168252602082018690529190931692a25b6122ed816156d6565b7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff61232084614428565b604080517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681523360208201529081019490945216918060608101611857565b6040517ffa7c07de000000000000000000000000000000000000000000000000000000006020820152602081526123ae604082613fce565b604051916123bb83613f7a565b8252602082015260405191829160208352602083019061434e565b7fe08f03ef00000000000000000000000000000000000000000000000000000000608051526080516004526024526044608051fd5b50611490826123769267ffffffffffffffff61242961133396614428565b168060005260076020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944828061249d604060002073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615c49565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a26122e4565b506020813d602011612525575b816124e460209383613fce565b810103126102c1575173ffffffffffffffffffffffffffffffffffffffff811681036102c15773ffffffffffffffffffffffffffffffffffffffff9061222c565b3d91506124d7565b73ffffffffffffffffffffffffffffffffffffffff1661255a816000526003602052604060002054151590565b61219e577fd0d2597600000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b6125a6915060203d602011611b8057611b728183613fce565b84612167565b611ba573ffffffffffffffffffffffffffffffffffffffff9161447c565b346102c15760606003193601126102c15760043567ffffffffffffffff81116102c1576125fb9036906004016140a8565b9060243567ffffffffffffffff81116102c15761261c90369060040161431d565b9060443567ffffffffffffffff81116102c15761263d90369060040161431d565b73ffffffffffffffffffffffffffffffffffffffff6009541633141580612705575b61123a578386148015906126fb575b6126cf576080515b8681106126835760805180f35b806126c9612697610a326001948b8b61446c565b6126a2838989614858565b6126c36126bb6126b386898b614858565b9236906143a2565b9136906143a2565b916155a6565b01612676565b7f568efce200000000000000000000000000000000000000000000000000000000608051526004608051fd5b508086141561266e565b5073ffffffffffffffffffffffffffffffffffffffff6001541633141561265f565b346102c1576080516003193601126102c157602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346102c15760206003193601126102c157602061279767ffffffffffffffff61278361412c565b166000526006602052604060002054151590565b6040519015158152f35b346102c15760206003193601126102c15760043567ffffffffffffffff81116102c1576127d5610a119136906004016140fb565b906127de614c99565b614d01565b346102c15760206003193601126102c1577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174602073ffffffffffffffffffffffffffffffffffffffff612834613ecf565b61283c614c99565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000006009541617600955604051908152a160805180f35b346102c1576080516003193601126102c1576080515473ffffffffffffffffffffffffffffffffffffffff8116330361291f577fffffffffffffffffffffffff0000000000000000000000000000000000000000600154913382841617600155166080515573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0608051608051a360805180f35b7f02b543c600000000000000000000000000000000000000000000000000000000608051526004608051fd5b346102c1576080516003193601126102c1576004546040805173ffffffffffffffffffffffffffffffffffffffff8316815260a09290921c602083015290f35b346102c1576080516003193601126102c157602073ffffffffffffffffffffffffffffffffffffffff60095416604051908152f35b346102c1576129ce366142dc565b6129d9929192614c99565b67ffffffffffffffff82166129fb816000526006602052604060002054151590565b15612a165750610a1192612a10913691614171565b9061533c565b7f1e670e4b00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b346102c1576080516003193601126102c1576020612a616147a0565b604051908152f35b346102c157612a91612a99612a7d3661428e565b9491612a8a939193614c99565b369161474c565b92369161474c565b7f000000000000000000000000000000000000000000000000000000000000000015612bbb576080515b8251811015612b36578073ffffffffffffffffffffffffffffffffffffffff612aee60019386614631565b5116612af981615ef3565b612b05575b5001612ac3565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a184612afe565b506080515b8151811015610a11578073ffffffffffffffffffffffffffffffffffffffff612b6660019385614631565b51168015612bb557612b77816162e8565b612b84575b505b01612b3b565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a183612b7c565b50612b7e565b7f35f4a7b300000000000000000000000000000000000000000000000000000000608051526004608051fd5b346102c157611333611e86612bfb366141c6565b5050509150614645565b346102c15760406003193601126102c157612c1e61412c565b60243567ffffffffffffffff81116102c157602091612c446127979236906004016141a8565b906144ee565b346102c15760406003193601126102c15760043567ffffffffffffffff81116102c157806004019061010060031982360301126102c157612c896140d9565b90604051612c9681613f5e565b6080519052612cc7612cbd612cb8612cb160c485018761449d565b3691614171565b6152a6565b606483013561518d565b916084820190612cd68261447c565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611b8757602483019477ffffffffffffffff00000000000000000000000000000000612d3c87614428565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611a90576080519161319f575b50611b2c5767ffffffffffffffff612dd287614428565b16612dea816000526006602052604060002054151590565b15611a9d57602073ffffffffffffffffffffffffffffffffffffffff60045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115611a905760805191613180575b50156119fd57612e6386614428565b90612e7960a4860192612c44612cb1858561449d565b156131395750604493929161ffff161590506130805767ffffffffffffffff612ea186614428565b1660805152600c602052612ebe604060805120856117478461447c565b7f41a8aa8df7945f0fb8ac5f7d88279638d9dc2ef9a6bf4ec9a53b80681b34aff767ffffffffffffffff612ef461178388614428565b6040805173ffffffffffffffffffffffffffffffffffffffff929092168252602082018890529190931692a25b01612f2b8161447c565b9073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691823b156102c1576040517f40c10f1900000000000000000000000000000000000000000000000000000000815260805173ffffffffffffffffffffffffffffffffffffffff909216600482015260248101859052908180604481010381608051875af18015611a9057613067575b5067ffffffffffffffff7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09161304e85613010611783602099614428565b6040805173ffffffffffffffffffffffffffffffffffffffff9889168152336020820152979091169087015260608601529116929081906080820190565b0390a28060405161305e81613f5e565b52604051908152f35b60805161307391613fce565b6080516102c15784612fd2565b5067ffffffffffffffff61309385614428565b16806080515260076020527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c848061310c60026040608051200173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615c49565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a2612f21565b613143925061449d565b611e066040519283927f24eb47e5000000000000000000000000000000000000000000000000000000008452602060048501526024840191614a9a565b613199915060203d602011611b8057611b728183613fce565b87612e54565b6131b8915060203d602011611b8057611b728183613fce565b87612dbb565b346102c15760206003193601126102c15760043567ffffffffffffffff81116102c157806004019061010060031982360301126102c15760405161320181613f5e565b60805190526132136064820135615085565b90608481016132218161447c565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000081169116036125ac5750602481019277ffffffffffffffff0000000000000000000000000000000061328885614428565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611a90576080519161356b575b50611b2c5767ffffffffffffffff61331e85614428565b16613336816000526006602052604060002054151590565b15611a9d57602073ffffffffffffffffffffffffffffffffffffffff60045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115611a90576080519161354c575b50156119fd576133af84614428565b906133c560a4840192612c44612cb1858561449d565b15613139575082916044915067ffffffffffffffff6133e386614428565b168060805152600760205261343960026040608051200173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016958691615c49565b6040805173ffffffffffffffffffffffffffffffffffffffff86168152602081018790527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a2019261348c8461447c565b823b156102c1576040517f40c10f1900000000000000000000000000000000000000000000000000000000815260805173ffffffffffffffffffffffffffffffffffffffff909216600482015260248101859052908180604481010381608051875af1948515611a905761304e856130106117837ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09667ffffffffffffffff9660209b61353a575b50614428565b60805161354691613fce565b8b613534565b613565915060203d602011611b8057611b728183613fce565b856133a0565b613584915060203d602011611b8057611b728183613fce565b85613307565b346102c15760a06003193601126102c1576135a3613ecf565b5060243567ffffffffffffffff8116908190036102c15760443567ffffffffffffffff81116102c15760031960a091360301126102c1576135e26140ea565b5060843567ffffffffffffffff81116102c157613603903690600401614143565b505060405161361181613f13565b608051815260805160208201526080516040820152606060805191015260805152600e602052608060408151206040519061364b82613f13565b5463ffffffff808216928381528160208201818560201c16815260ff60606040850194848860401c168652019560601c161515855260405195865251166020850152511660408301525115156060820152f35b346102c15760606003193601126102c15760043561ffff8116908190036102c1576136c76140d9565b9060443567ffffffffffffffff81116102c1576136e89036906004016140fb565b906136f1614c99565b61ffff84169361271085101561377b5791613764917f52aa194b292c8bfb5aaca8ee2000a965c3a051b306ff841873b16147526a39ba9593857fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000063ffff0000600a549360101b1692161717600a55614d01565b60408051928352602083019190915290a160805180f35b847f95f3517a00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b346102c15760406003193601126102c15760043567ffffffffffffffff81116102c157366023820112156102c157806004013567ffffffffffffffff81116102c15760248201916024369160a084020101116102c15760243567ffffffffffffffff81116102c1576138219036906004016140a8565b91909261382c614c99565b6080515b8281106138a8575050506080515b81811061384b5760805180f35b8067ffffffffffffffff613865610a32600194868861446c565b168060805152600e602052608051604060805120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8608051608051a20161383e565b806138b9610a3260019386866143e9565b7f56f77aeff2def50c8b5f5a0df3bab7183df09bf36c6feba496bb42551db77d7060806138e78488886143e9565b926139f76139bd63ffffffff6139ec6139b0826139e167ffffffffffffffff60208c0198169a8b8a5152600e60205260408a5120836139258b61443d565b169181549060408101937fffffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffff67ffffffff000000006139628761443d565b60201b16918f6cff0000000000000000000000007fffffffffffffffffffffffffffffffffffffffff000000000000000000000000916bffffffff0000000000000000606088019d8e61443d565b60401b1696019e8f61444e565b151560601b16951617161717179055826139d96040519a61445b565b16895261445b565b16602087015261445b565b166040840152614378565b15156060820152a201613830565b346102c1576080516003193601126102c157602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102c15760206003193601126102c1576020613a5f613ecf565b6040517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff9081169216919091148152f35b346102c1576080516003193601126102c157602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102c1576080516003193601126102c1576040805161133391613b1b9082613fce565b601b81527f4275726e4d696e74546f6b656e506f6f6c20312e362e332d64657600000000006020820152604051918291602083526020830190614049565b346102c15760206003193601126102c157613b72613ecf565b613b7a614c99565b613b826147a0565b9081613b8e5760805180f35b6040517fa9059cbb000000000000000000000000000000000000000000000000000000006020820190815273ffffffffffffffffffffffffffffffffffffffff831660248301526044808301859052825290613ca390613bef606482613fce565b60408051909390917f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169190613c3a8685613fce565b602084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602085015260805191608051915190608051855af13d15613d96573d91613c858361400f565b92613c9287519485613fce565b83526080513d90602085013e6163f7565b805180613cf5575b505073ffffffffffffffffffffffffffffffffffffffff7f738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b85959992602092519485521692a28080610a11565b90602080613d0793830101910161528e565b15613d13578380613cab565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b6060916163f7565b346102c15760206003193601126102c157600435907fffffffff0000000000000000000000000000000000000000000000000000000082168092036102c157817ff208a58f0000000000000000000000000000000000000000000000000000000060209314908115613ea5575b8115613e7b575b8115613e51575b8115613e27575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483613e20565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150613e19565b7f1ef5498f0000000000000000000000000000000000000000000000000000000081149150613e12565b7faff2afbf0000000000000000000000000000000000000000000000000000000081149150613e0b565b6004359073ffffffffffffffffffffffffffffffffffffffff82168203610a0c57565b359073ffffffffffffffffffffffffffffffffffffffff82168203610a0c57565b6080810190811067ffffffffffffffff821117613f2f57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff821117613f2f57604052565b6040810190811067ffffffffffffffff821117613f2f57604052565b60a0810190811067ffffffffffffffff821117613f2f57604052565b6060810190811067ffffffffffffffff821117613f2f57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117613f2f57604052565b67ffffffffffffffff8111613f2f57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b8481106140935750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201614054565b9181601f84011215610a0c5782359167ffffffffffffffff8311610a0c576020808501948460051b010111610a0c57565b6024359061ffff82168203610a0c57565b6064359061ffff82168203610a0c57565b9181601f84011215610a0c5782359167ffffffffffffffff8311610a0c5760208085019460e08502010111610a0c57565b6004359067ffffffffffffffff82168203610a0c57565b9181601f84011215610a0c5782359167ffffffffffffffff8311610a0c5760208381860195010111610a0c57565b92919261417d8261400f565b9161418b6040519384613fce565b829481845281830111610a0c578281602093846000960137010152565b9080601f83011215610a0c578160206141c393359101614171565b90565b60a0600319820112610a0c5760043573ffffffffffffffffffffffffffffffffffffffff81168103610a0c579160243567ffffffffffffffff81168103610a0c57916044359160643561ffff81168103610a0c57916084359067ffffffffffffffff8211610a0c5761423a91600401614143565b9091565b602060408183019282815284518094520192019060005b8181106142625750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101614255565b6040600319820112610a0c5760043567ffffffffffffffff8111610a0c57816142b9916004016140a8565b929092916024359067ffffffffffffffff8211610a0c5761423a916004016140a8565b906040600319830112610a0c5760043567ffffffffffffffff81168103610a0c57916024359067ffffffffffffffff8211610a0c5761423a91600401614143565b9181601f84011215610a0c5782359167ffffffffffffffff8311610a0c5760208085019460608502010111610a0c57565b6141c39160206143678351604084526040840190614049565b920151906020818403910152614049565b35908115158203610a0c57565b35906fffffffffffffffffffffffffffffffff82168203610a0c57565b9190826060910312610a0c576040516143ba81613fb2565b60406143e48183956143cb81614378565b85526143d960208201614385565b602086015201614385565b910152565b91908110156143f95760a0020190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3567ffffffffffffffff81168103610a0c5790565b3563ffffffff81168103610a0c5790565b358015158103610a0c5790565b359063ffffffff82168203610a0c57565b91908110156143f95760051b0190565b3573ffffffffffffffffffffffffffffffffffffffff81168103610a0c5790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610a0c570180359067ffffffffffffffff8211610a0c57602001918136038313610a0c57565b9067ffffffffffffffff6141c392166000526007602052600560406000200190602081519101209060019160005201602052604060002054151590565b906040519182815491828252602082019060005260206000209260005b81811061455f57505061455d92500383613fce565b565b845473ffffffffffffffffffffffffffffffffffffffff16835260019485019487945060209093019201614548565b9190820180921161459b57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b67ffffffffffffffff8111613f2f5760051b60200190565b906145ec826145ca565b6145f96040519182613fce565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061462782946145ca565b0190602036910137565b80518210156143f95760209160051b010190565b67ffffffffffffffff16600052600d60205260406000206146658161452b565b9160045460a01c8015159182614741575b5050614680575090565b600161468c910161452b565b9081518061469a5750905090565b6146a86146ad91835161458e565b6145e2565b9260005b82518110156146ef578073ffffffffffffffffffffffffffffffffffffffff6146dc60019386614631565b51166146e88288614631565b52016146b1565b509160005b815181101561473c578073ffffffffffffffffffffffffffffffffffffffff61471f60019385614631565b511661473561472f83875161458e565b88614631565b52016146f4565b505050565b101590503880614676565b929190614758816145ca565b936147666040519586613fce565b602085838152019160051b8101928311610a0c57905b82821061478857505050565b6020809161479584613ef2565b81520191019061477c565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561484c5760009161481d575090565b90506020813d602011614844575b8161483860209383613fce565b81010312610a0c575190565b3d915061482b565b6040513d6000823e3d90fd5b91908110156143f9576060020190565b6040519061487582613f7a565b60606020838281520152565b67ffffffffffffffff16600052600d60205260406000209060036148a76002840161452b565b9201906148b38261452b565b5060045460a01c801515918261497c575b50506148ce575090565b6148d79061452b565b908151806148e55750905090565b6146a86148f391835161458e565b9260005b8251811015614935578073ffffffffffffffffffffffffffffffffffffffff61492260019386614631565b511661492e8288614631565b52016148f7565b509160005b815181101561473c578073ffffffffffffffffffffffffffffffffffffffff61496560019385614631565b511661497561472f83875161458e565b520161493a565b1015905038806148c4565b90600182811c921680156149d0575b60208310146149a157565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691614996565b90604051918260008254926149ee84614987565b8084529360018116908115614a5a5750600114614a13575b5061455d92500383613fce565b90506000929192526020600020906000915b818310614a3e57505090602061455d9282010138614a06565b6020919350806001915483858901015201910190918492614a25565b6020935061455d9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138614a06565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b60405190614ae682613f96565b60006080838281528260208201528260408201528260608201520152565b90604051614b1181613f96565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff1660005260076020526141c360046040600020016149da565b91908110156143f95760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6181360301821215610a0c570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610a0c570180359067ffffffffffffffff8211610a0c57602001918160051b36038313610a0c57565b8181029291811591840414171561459b57565b818110614c3e575050565b60008155600101614c33565b9160209082815201919060005b818110614c645750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff614c8b88613ef2565b168152019401929101614c57565b73ffffffffffffffffffffffffffffffffffffffff600154163303614cba57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b356fffffffffffffffffffffffffffffffff81168103610a0c5790565b9160005b828110156150215760e0810284016000614d1e82614428565b9067ffffffffffffffff821691614d42836000526006602052604060002054151590565b15614ff557614e0b9260408593614db6614db094614db0614d76602060019c9b0192611723614d7136866143a2565b6158cf565b91825463ffffffff8160801c16159081614fd7575b81614fc8575b81614fad575b81614f9e575b5080614f8f575b614f04575b36906143a2565b90615a16565b6080850192614dc8614d7136866143a2565b8152600c6020522092835463ffffffff8160801c16159081614ee6575b81614ed7575b81614ebc575b81614ead575b5080614e9e575b614e11575b5036906143a2565b01614d05565b614e2e60a06fffffffffffffffffffffffffffffffff9201614ce4565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff161717835538614e03565b50614ea88261444e565b614dfe565b60ff915060a01c161538614df7565b6fffffffffffffffffffffffffffffffff8116159150614df1565b8589015460801c159150614deb565b858901546fffffffffffffffffffffffffffffffff16159150614de5565b6fffffffffffffffffffffffffffffffff614f20878b01614ce4565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff1617178355614da9565b50614f998161444e565b614da4565b60ff915060a01c161538614d9d565b6fffffffffffffffffffffffffffffffff8116159150614d97565b848e015460801c159150614d91565b848e01546fffffffffffffffffffffffffffffffff16159150614d8b565b506024917f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50915050565b9060ff8091169116039060ff821161459b57565b60ff16604d811161459b57600a0a90565b8115615056570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f000000000000000000000000000000000000000000000000000000000000000060ff81169081600614615188578160061161515d5760066150c691615027565b90604d60ff8316118015615124575b6150ed5750906150e76141c39261503b565b90614c20565b90507fa9cb113d00000000000000000000000000000000000000000000000000000000600052600660045260245260445260646000fd5b5061512e8261503b565b8015615056577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483116150d5565b615168906006615027565b90604d60ff8316116150ed5750906151826141c39261503b565b9061504c565b505090565b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff8116928284146152875782841161526357906151d291615027565b91604d60ff841611801561522a575b6151f4575050906150e76141c39261503b565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b506152348361503b565b8015615056577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0484116151e1565b61526c91615027565b91604d60ff8416116151f4575050906151826141c39261503b565b5050505090565b90816020910312610a0c57518015158103610a0c5790565b80518015615316576020036152d8578051602082810191830183900312610a0c57519060ff82116152d8575060ff1690565b611e06906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190614049565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9080511561557c5767ffffffffffffffff815160208301209216918260005260076020526153718160056040600020016163a2565b156155385760005260086020526040600020815167ffffffffffffffff8111613f2f5761539e8254614987565b601f8111615506575b506020601f8211600114615440579161541a827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea959361543095600091615435575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055604051918291602083526020830190614049565b0390a2565b9050840151386153e9565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b8181106154ee5750926154309492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9896106154b7575b5050811b019055611495565b8501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c1916905538806154ab565b9192602060018192868a015181550194019201615470565b61553290836000526020600020601f840160051c8101916020851061093b57601f0160051c0190614c33565b386153a7565b5090611e066040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190614049565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b67ffffffffffffffff1660008181526006602052604090205490929190156156a857916156a560e092615671856155fd7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b976158cf565b846000526007602052615614816040600020615a16565b61561d836158cf565b846000526007602052615637836002604060002001615a16565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690813b15610a0c57604051907f42966c680000000000000000000000000000000000000000000000000000000082528160248160008096819560048401525af180156157685761575b575050565b8161576591613fce565b50565b6040513d84823e3d90fd5b9190820391821161459b57565b615788614ad9565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff80835116916157e560208501936157df6157d263ffffffff87511642615773565b8560808901511690614c20565b9061458e565b808210156157fe57505b16825263ffffffff4216905290565b90506157ef565b805160005b81811061581657505050565b6001810180821161459b575b828110615832575060010161580a565b73ffffffffffffffffffffffffffffffffffffffff6158518386614631565b511673ffffffffffffffffffffffffffffffffffffffff6158728387614631565b51161461588157600101615822565b73ffffffffffffffffffffffffffffffffffffffff6158a08386614631565b51167fa1726e400000000000000000000000000000000000000000000000000000000060005260045260246000fd5b80511561596f576fffffffffffffffffffffffffffffffff6040820151166fffffffffffffffffffffffffffffffff6020830151161061590c5750565b60649061596d604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff604082015116158015906159f7575b6159965750565b60649061596d604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff602082015116151561598f565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1991615b4f6060928054615a5363ffffffff8260801c1642615773565b9081615b8e575b50506fffffffffffffffffffffffffffffffff6001816020860151169282815416808510600014615b8657508280855b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416178155615b038651151582907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b60408601517fffffffffffffffffffffffffffffffff0000000000000000000000000000000060809190911b16939092166fffffffffffffffffffffffffffffffff1692909217910155565b6156a560405180926fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b838091615a8a565b6fffffffffffffffffffffffffffffffff91615bc3839283615bbc6001880154948286169560801c90614c20565b911661458e565b80821015615c4257505b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff9290911692909216167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116174260801b73ffffffff00000000000000000000000000000000161781553880615a5a565b9050615bcd565b9182549060ff8260a01c16158015615e88575b615e82576fffffffffffffffffffffffffffffffff82169160018501908154615ca163ffffffff6fffffffffffffffffffffffffffffffff83169360801c1642615773565b9081615de4575b5050848110615d985750838310615d02575050615cd76fffffffffffffffffffffffffffffffff928392615773565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b5460801c91615d118185615773565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81019080821161459b57615d5f615d649273ffffffffffffffffffffffffffffffffffffffff9661458e565b61504c565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828573ffffffffffffffffffffffffffffffffffffffff927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611615e5857615dff926157df9160801c90614c20565b80841015615e535750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178655923880615ca8565b615e0a565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b508215615c5c565b906040519182815491828252602082019060005260206000209260005b818110615ec257505061455d92500383613fce565b8454835260019485019487945060209093019201615ead565b80548210156143f95760005260206000200190600090565b6000818152600360205260409020548015616082577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161459b57600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161459b57818103616013575b5050506002548015615fe4577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01615fa1816002615edb565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b61606a616024616035936002615edb565b90549060031b1c9283926002615edb565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526003602052604060002055388080615f68565b5050600090565b6000818152600660205260409020548015616082577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161459b57600554907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161459b5781810361617a575b5050506005548015615fe4577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01616137816005615edb565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600555600052600660205260006040812055600190565b61619c61618b616035936005615edb565b90549060031b1c9283926005615edb565b905560005260066020526040600020553880806160fe565b90600182019181600052826020526040600020548015156000146162df577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161459b578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161459b578181036162a8575b50505080548015615fe4577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906162698282615edb565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b6162c86162b86160359386615edb565b90549060031b1c92839286615edb565b905560005283602052604060002055388080616231565b50505050600090565b806000526003602052604060002054156000146163425760025468010000000000000000811015613f2f576163296160358260018594016002556002615edb565b9055600254906000526003602052604060002055600190565b50600090565b806000526006602052604060002054156000146163425760055468010000000000000000811015613f2f576163896160358260018594016005556005615edb565b9055600554906000526006602052604060002055600190565b60008281526001820160205260409020546160825780549068010000000000000000821015613f2f57826163e0616035846001809601855584615edb565b905580549260005201602052604060002055600190565b91929015616472575081511561640b575090565b3b156164145790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156164855750805190602001fd5b611e06906040519182917f08c379a000000000000000000000000000000000000000000000000000000000835260206004840152602483019061404956fea164736f6c634300081a000a",
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

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

	error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getDynamicConfig")

	outstruct := new(GetDynamicConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Router = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ThresholdAmountForAdditionalCCVs = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetDynamicConfig(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetDynamicConfig(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getRateLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetRateLimitAdmin() (common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetRateLimitAdmin(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetRateLimitAdmin() (common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetRateLimitAdmin(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
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

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetRequiredInboundCCVs(opts *bind.CallOpts, arg0 common.Address, sourceChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getRequiredInboundCCVs", arg0, sourceChainSelector, amount, arg3, arg4)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetRequiredInboundCCVs(arg0 common.Address, sourceChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetRequiredInboundCCVs(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, arg0, sourceChainSelector, amount, arg3, arg4)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetRequiredInboundCCVs(arg0 common.Address, sourceChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetRequiredInboundCCVs(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, arg0, sourceChainSelector, amount, arg3, arg4)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetRequiredOutboundCCVs(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getRequiredOutboundCCVs", arg0, destChainSelector, amount, arg3, arg4)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetRequiredOutboundCCVs(arg0 common.Address, destChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetRequiredOutboundCCVs(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, arg0, destChainSelector, amount, arg3, arg4)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetRequiredOutboundCCVs(arg0 common.Address, destChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetRequiredOutboundCCVs(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, arg0, destChainSelector, amount, arg3, arg4)
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

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.AcceptOwnership(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.AcceptOwnership(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts)
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

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) SetDynamicConfig(opts *bind.TransactOpts, router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "setDynamicConfig", router, thresholdAmountForAdditionalCCVs)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) SetDynamicConfig(router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.SetDynamicConfig(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, router, thresholdAmountForAdditionalCCVs)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) SetDynamicConfig(router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.SetDynamicConfig(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, router, thresholdAmountForAdditionalCCVs)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "setRateLimitAdmin", rateLimitAdmin)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.SetRateLimitAdmin(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, rateLimitAdmin)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.SetRateLimitAdmin(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, rateLimitAdmin)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.TransferOwnership(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, to)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.TransferOwnership(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, to)
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
	RemoteChainSelector    uint64
	OutboundCCVs           []common.Address
	AdditionalOutboundCCVs []common.Address
	InboundCCVs            []common.Address
	AdditionalInboundCCVs  []common.Address
	Raw                    types.Log
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

type BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSetIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSetIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSet struct {
	Router                           common.Address
	ThresholdAmountForAdditionalCCVs *big.Int
	Raw                              types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterDynamicConfigSet(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSetIterator, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSetIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "DynamicConfigSet", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSet) (event.Subscription, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSet)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseDynamicConfigSet(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSet, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSet)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
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

type BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequestedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequestedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequested)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequested, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequested)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferredIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferredIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferred)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferred, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferred)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

type BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminSetIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminSetIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminSet struct {
	RateLimitAdmin common.Address
	Raw            types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterRateLimitAdminSet(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminSetIterator, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminSetIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "RateLimitAdminSet", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminSet) (event.Subscription, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminSet)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseRateLimitAdminSet(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminSet, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminSet)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
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

type GetDynamicConfig struct {
	Router                           common.Address
	ThresholdAmountForAdditionalCCVs *big.Int
}

func (BurnMintWithLockReleaseFlagTokenPoolAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (BurnMintWithLockReleaseFlagTokenPoolAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (BurnMintWithLockReleaseFlagTokenPoolCCVConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d")
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

func (BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x39b29a1c46dd0e3ee5683fed1d66e11a234938cde48e5c4e14389be508348eb6")
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

func (BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (BurnMintWithLockReleaseFlagTokenPoolPoolFeeWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b859599")
}

func (BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminSet) Topic() common.Hash {
	return common.HexToHash("0x44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174")
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
	GetAccumulatedFees(opts *bind.CallOpts) (*big.Int, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

		error)

	GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRequiredInboundCCVs(opts *bind.CallOpts, arg0 common.Address, sourceChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error)

	GetRequiredOutboundCCVs(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error)

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

	ApplyFinalityConfigUpdates(opts *bind.TransactOpts, finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error)

	ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error)

	SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error)

	SetCustomFinalityRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error)

	SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

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

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSet, error)

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

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferred, error)

	FilterPoolFeeWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*BurnMintWithLockReleaseFlagTokenPoolPoolFeeWithdrawnIterator, error)

	WatchPoolFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolPoolFeeWithdrawn, recipient []common.Address) (event.Subscription, error)

	ParsePoolFeeWithdrawn(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolPoolFeeWithdrawn, error)

	FilterRateLimitAdminSet(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminSetIterator, error)

	WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminSet) (event.Subscription, error)

	ParseRateLimitAdminSet(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminSet, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolReleasedOrMinted, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolRemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolRemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemoved, error)

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdatedIterator, error)

	WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigUpdated(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdated, error)

	Address() common.Address
}
