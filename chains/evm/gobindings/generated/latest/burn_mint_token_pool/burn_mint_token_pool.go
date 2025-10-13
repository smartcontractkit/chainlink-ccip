// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package burn_mint_token_pool

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

var BurnMintTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIBurnMintERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"additionalOutboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"additionalInboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyFinalityConfigUpdates\",\"inputs\":[{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"destToUseDefaultFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAccumulatedFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredInboundCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredOutboundCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"lockOrBurnOutV1\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustomFinalityRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"additionalOutboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"additionalInboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityTransferInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalityConfigUpdated\",\"inputs\":[{\"name\":\"finalityConfig\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeeWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidFinality\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidFinalityConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenTransferFeeConfigNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x61010080604052346103b7576168ed803803809161001d8285610436565b833981019060a0818303126103b75780516001600160a01b038116908190036103b75761004c60208301610459565b60408301519091906001600160401b0381116103b75783019380601f860112156103b7578451946001600160401b038611610420578560051b9060208201966100986040519889610436565b87526020808801928201019283116103b757602001905b828210610408575050506100d160806100ca60608601610467565b9401610467565b9233156103f757600180546001600160a01b03191633179055811580156103e6575b80156103d5575b6103c4578160209160049360805260c0526040519283809263313ce56760e01b82525afa60009181610383575b50610358575b5060a052600480546001600160a01b0319166001600160a01b03929092169190911790558051151560e081905261023a575b6040516162d1908161061c8239608051818181611533015281816117d2015281816119ec0152818161200c015281816121fc01528181612c7001528181612ec60152818161305b015281816131c9015281816133a201528181613a0401528181613a7301528181613b990152614786015260a0518181816118a101528181612318015281816139c00152818161503501526150b8015260c051818181610c48015281816115ce015281816120a701528181612d0b0152613266015260e051818181610bf501528181611613015281816120ec0152612a1b0152f35b60405160206102498183610436565b60008252600036813760e051156103475760005b82518110156102c4576001906001600160a01b0361027b828661047b565b511683610287826104bd565b610294575b50500161025d565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a1388361028c565b50905060005b825181101561033e576001906001600160a01b036102e8828661047b565b5116801561033857836102fa826105bb565b610308575b50505b016102ca565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a138836102ff565b50610302565b5050503861015f565b6335f4a7b360e01b60005260046000fd5b60ff1660ff821681810361036c575061012d565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d6020116103bc575b8161039f60209383610436565b810103126103b7576103b090610459565b9038610127565b600080fd5b3d9150610392565b630a64406560e11b60005260046000fd5b506001600160a01b038116156100fa565b506001600160a01b038416156100f3565b639b15e16f60e01b60005260046000fd5b6020809161041584610467565b8152019101906100af565b634e487b7160e01b600052604160045260246000fd5b601f909101601f19168101906001600160401b0382119082101761042057604052565b519060ff821682036103b757565b51906001600160a01b03821682036103b757565b805182101561048f5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b805482101561048f5760005260206000200190600090565b60008181526003602052604090205480156105b457600019810181811161059e5760025460001981019190821161059e5781810361054d575b505050600254801561053757600019016105118160026104a5565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61058661055e61056f9360026104a5565b90549060031b1c92839260026104a5565b819391549060031b91821b91600019901b19161790565b905560005260036020526040600020553880806104f6565b634e487b7160e01b600052601160045260246000fd5b5050600090565b806000526003602052604060002054156000146106155760025468010000000000000000811015610420576105fc61056f82600185940160025560026104a5565b9055600254906000526003602052604060002055600190565b5060009056fe60a080604052600436101561001357600080fd5b60006080526080513560e01c90816301ffc9a714613d3e57508063164e68de14613af9578063181f5a7714613a9757806321df0da714613a45578063240028e8146139e457806324f65ee7146139a55780632a10097b1461374b5780632c286daf1461364157806337b192471461352d5780633907753714613135578063489a68f214612bca5780634c5ef0ed14612b855780634f71592c14612b6757806354c8a4f3146129e95780635df45a37146129c557806362ddd3c414612940578063698c2c66146128745780636d3d1a581461283f5780637437ff9f146127fe57806379ba5097146127275780637d54534e14612696578063804ba5a9146126545780638926f54f1461260f5780638da5cb5b146125da578063962d40201461247d5780639a4575b914611f8e5780639f68f67314611f70578063a42a7b8b14611dfb578063a7cd63b714611d94578063acfecf9114611c67578063af58d59f14611c1b578063b1c71c65146114aa578063b794658014611472578063c4bffe2b1461133f578063c75eea9c14611294578063cf7401f3146110e8578063d966866b14610c6c578063dc0bd97114610c1a578063e0351e1314610bdc578063e8a1da17146102c75763f2fde38b146101e857600080fd5b346102c15760206003193601126102c15773ffffffffffffffffffffffffffffffffffffffff610216613e6f565b61021e614c33565b163381146102955760805180547fffffffffffffffffffffffff0000000000000000000000000000000000000000168217815560015473ffffffffffffffffffffffffffffffffffffffff16907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789080a360805180f35b7fdad89dca00000000000000000000000000000000000000000000000000000000608051526004608051fd5b60805180fd5b346102c1576102d53661422e565b926102e1929192614c33565b608051905b828210610a1a5750505060805163ffffffff4216917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee181360301915b84811015610a14576000938160051b830135848112156102c157830190610120823603126102c1576040519561035787613f36565b82359067ffffffffffffffff82168203610a1157508652602082013567ffffffffffffffff81116102c15782019436601f870112156102c157853561039b8161456a565b966103a96040519889613f6e565b81885260208089019260051b820101903682116102c15760208101925b8284106109e2575050505060208701958652604083013567ffffffffffffffff81116102c1576103f99036908501614148565b91604088019283526104236104113660608701614342565b9460608a0195865260c0369101614342565b956080890196875261043585516156d0565b61043f87516156d0565b835151156109b65761045b67ffffffffffffffff8a5116616149565b1561097b5767ffffffffffffffff89511660805152600860205260406080512061059f86516fffffffffffffffffffffffffffffffff6040820151169061055a6fffffffffffffffffffffffffffffffff602083015116915115158360806040516104c581613f36565b858152602081018b905260408101849052606081018690520152855474ff000000000000000000000000000000000000000091151560a01b919091167fffffffffffffffffffffff0000000000000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff84161773ffffffff0000000000000000000000000000000060808a901b1617178555565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176001830155565b6106c588516fffffffffffffffffffffffffffffffff604082015116906106806fffffffffffffffffffffffffffffffff602083015116915115158360806040516105e981613f36565b858152602081018b9052604081018490526060810186905201526002860180547fffffffffffffffffffffff000000000000000000000000000000000000000000166fffffffffffffffffffffffffffffffff85161773ffffffff0000000000000000000000000000000060808b901b161791151560a01b74ff000000000000000000000000000000000000000016919091179055565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176003830155565b6004855191019080519067ffffffffffffffff821161094a576106e88354614921565b601f811161090b575b506020906001601f84111461086557918091610742936080519261085a575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b6080515b8851805182101561077e57906107786001926107718367ffffffffffffffff8f5116926145d1565b51906151da565b01610749565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c293919997509561084c67ffffffffffffffff60019796949851169251935191516108186107e360405196879687526101006020880152610100870190613fe9565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a1019392909193610322565b015190508e80610710565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083169184608051528160805120926080515b8181106108f357509084600195949392106108bc575b505050811b019055610745565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558d80806108af565b92936020600181928786015181550195019301610899565b61093a908460805152602060805120601f850160051c81019160208610610940575b601f0160051c0190614bcd565b8d6106f1565b909150819061092d565b7f4e487b71000000000000000000000000000000000000000000000000000000006080515260416004526024608051fd5b67ffffffffffffffff8951167f1d5ad3c500000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b7f14c880ca00000000000000000000000000000000000000000000000000000000608051526004608051fd5b833567ffffffffffffffff81116102c157602091610a068392833691870101614148565b8152019301926103c6565b80fd5b60805180f35b9092919367ffffffffffffffff610a3a610a3586888661440c565b6143c8565b1692610a4584615e8a565b15610bac5783608051526008602052610a65600560406080512001615c91565b926080515b8451811015610aa45760019086608051526008602052610a9d600560406080512001610a9683896145d1565b5190615fb5565b5001610a6a565b5093909491959250806080515260086020526005604060805120608051815560805160018201556080516002820155608051600382015560048101610ae98154614921565b80610b5b575b505001805490608051815581610b38575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a10190916102e6565b60805152602060805120908101905b81811015610b005760008155600101610b47565b601f8111600114610b735750600090555b8880610aef565b610b939082608051526001601f6020608051209201861c82019101614bcd565b8060805152600060206080512091556080519055610b6c565b837f1e670e4b00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b346102c1576080516003193601126102c15760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b346102c1576080516003193601126102c157602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102c15760206003193601126102c15760043567ffffffffffffffff81116102c157610c9d903690600401614048565b610ca5614c33565b608051915b818310610cb75760805180f35b610cc5610a35848484614b26565b610cdd610cd3858585614b26565b6020810190614b66565b9091610cf7610ced878787614b26565b6040810190614b66565b90610d10610d06898989614b26565b6060810190614b66565b9091610d2a610d208b8b8b614b26565b6080810190614b66565b949097610d40610d3b368a846146e9565b615606565b610d4e610d3b3684866146e9565b610d5c610d3b3686886146e9565b610d6a610d3b36888c6146e9565b604051610d7681613eb3565b610d81368a846146e9565b8152610d8e3684866146e9565b6020820152610d9e3686886146e9565b6040820152610dae36888c6146e9565b606082015267ffffffffffffffff881660805152600e602052604060805120815180519067ffffffffffffffff821161094a5768010000000000000000821161094a5760209083548385558084106110c9575b500182608051526020608051206080515b83811061109f5750505050602082015180519067ffffffffffffffff821161094a5768010000000000000000821161094a57602090600184015483600186015580841061107d575b500160018301608051526020608051206080515b8381106110535750505050604082015180519067ffffffffffffffff821161094a5768010000000000000000821161094a576020906002840154836002860155808410611031575b500160028301608051526020608051206080515b838110611007575050505060036060919e9c9d9e019101519081519167ffffffffffffffff831161094a5768010000000000000000831161094a576020908254848455808510610fe8575b500190608051526020608051206080515b838110610fbe5750505050610fa3608095610fb395610f957fece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d9a9660019e9d9c9a96610f8767ffffffffffffffff976040519d8d8f9e8f9081520191614be4565b918b830360208d0152614be4565b9188830360408a0152614be4565b9285840360608701521696614be4565b0390a2019190610caa565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501610f26565b6110019084608051528584608051209182019101614bcd565b38610f15565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501610eca565b61104d9060028601608051528484608051209182019101614bcd565b38610eb6565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501610e6e565b6110999060018601608051528484608051209182019101614bcd565b38610e5a565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501610e12565b6110e29085608051528484608051209182019101614bcd565b38610e01565b346102c15760e06003193601126102c1576111016140cc565b60607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc36011261128f5760405161113781613f52565b602435801515810361128f5781526044356fffffffffffffffffffffffffffffffff8116810361128f5760208201526064356fffffffffffffffffffffffffffffffff8116810361128f57604082015260607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c36011261128f57604051906111be82613f52565b608435801515810361128f57825260a4356fffffffffffffffffffffffffffffffff8116810361128f57602083015260c4356fffffffffffffffffffffffffffffffff8116810361128f57604083015273ffffffffffffffffffffffffffffffffffffffff600a54163314158061126d575b61123d57610a1492615444565b7f8e4a23d60000000000000000000000000000000000000000000000000000000060805152336004526024608051fd5b5073ffffffffffffffffffffffffffffffffffffffff60015416331415611230565b600080fd5b346102c15760206003193601126102c15767ffffffffffffffff6112b66140cc565b6112be614a73565b501660805152600860205261133b6112e26112dd604060805120614a9e565b615581565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b0390f35b346102c1576080516003193601126102c1576040516006548082528160208101600660005260206000209260005b81811061145957505061138292500382613f6e565b8051906113a76113918361456a565b9261139f6040519485613f6e565b80845261456a565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06020840192013683376080515b8151811015611408578067ffffffffffffffff6113f5600193856145d1565b511661140182876145d1565b52016113d6565b505090604051918291602083019060208452518091526040830191906080515b818110611436575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101611428565b845483526001948501948694506020909301920161136d565b346102c15760206003193601126102c15761133b6114966114916140cc565b614b04565b604051918291602083526020830190613fe9565b346102c15760606003193601126102c15760043567ffffffffffffffff81116102c15760a060031982360301126102c1576114e3614079565b60443567ffffffffffffffff81116102c157611503903690600401614148565b5061150c614805565b50608482019061151b8261446d565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611bcd57602483019277ffffffffffffffff00000000000000000000000000000000611581856143c8565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156119235760805191611b9e575b50611b72576116116044820161446d565b7f0000000000000000000000000000000000000000000000000000000000000000611b12575b5067ffffffffffffffff61164a856143c8565b16611662816000526007602052604060002054151590565b15611ae357602073ffffffffffffffffffffffffffffffffffffffff60045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa80156119235760805190611a80575b73ffffffffffffffffffffffffffffffffffffffff9150163303611a5057606461ffff91013591169182151592838094611a41575b1561198a5761ffff600b54169081811061195857505061174961173961171f866143c8565b67ffffffffffffffff16600052600c602052604060002090565b836117438461446d565b91615a4a565b7f7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b267ffffffffffffffff61178561177f876143c8565b9361446d565b6040805173ffffffffffffffffffffffffffffffffffffffff929092168252602082018690529190931692a25b908190611930575b5073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691823b156102c1576040517f42966c6800000000000000000000000000000000000000000000000000000000815282600482015260805181602481608051885af1801561192357611907575b50611491816118fd947ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10606067ffffffffffffffff61187f611899976143c8565b1692604051908152336020820152876040820152a26143c8565b9060405160ff7f0000000000000000000000000000000000000000000000000000000000000000166020820152602081526118d5604082613f6e565b604051926118e284613f1a565b835260208301526040519283926040845260408401906142ee565b9060208301520390f35b60805161191391613f6e565b6080516000126102c1578361183e565b6040513d608051823e3d90fd5b611952915061271061194b61ffff600b5460101c1683614bba565b0490615574565b826117ba565b7fe08f03ef00000000000000000000000000000000000000000000000000000000608051526004526024526044608051fd5b505067ffffffffffffffff61199e846143c8565b16806080515260086020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789448280611a1460406080512073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615a4a565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a26117b2565b5061ffff600b541615156116fa565b7f728fe07b0000000000000000000000000000000000000000000000000000000060805152336004526024608051fd5b506020813d602011611adb575b81611a9a60209383613f6e565b810103126102c1575173ffffffffffffffffffffffffffffffffffffffff811681036102c15773ffffffffffffffffffffffffffffffffffffffff906116c5565b3d9150611a8d565b7fa9902c7e00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b73ffffffffffffffffffffffffffffffffffffffff16611b3f816000526003602052604060002054151590565b611637577fd0d2597600000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b7f53ad11d800000000000000000000000000000000000000000000000000000000608051526004608051fd5b611bc0915060203d602011611bc6575b611bb88183613f6e565b8101906151c2565b85611600565b503d611bae565b73ffffffffffffffffffffffffffffffffffffffff611beb8361446d565b7f961c9a4f0000000000000000000000000000000000000000000000000000000060805152166004526024608051fd5b346102c15760206003193601126102c15767ffffffffffffffff611c3d6140cc565b611c45614a73565b501660805152600860205261133b6112e26112dd600260406080512001614a9e565b346102c15767ffffffffffffffff611c7e3661427c565b929091611c89614c33565b1690611ca2826000526007602052604060002054151590565b15611d645781608051526008602052611cd5600560406080512001611cc8368685614111565b6020815191012090615fb5565b15611d1d577f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d769192611d14604051928392602084526020840191614a34565b0390a260805180f35b611d60906040519384937f74f23c7c0000000000000000000000000000000000000000000000000000000085526004850152604060248501526044840191614a34565b0390fd5b507f1e670e4b00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b346102c1576080516003193601126102c1576040516002548082526020820190600260005260206000209060005b818110611de55761133b85611dd981870382613f6e565b604051918291826141de565b8254845260209093019260019283019201611dc2565b346102c15760206003193601126102c15767ffffffffffffffff611e1d6140cc565b16608051526008602052611e38600560406080512001615c91565b8051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611e7e611e688461456a565b93611e766040519586613f6e565b80855261456a565b016080515b818110611f5f5750506080515b8151811015611eda5780611ea6600192846145d1565b51608051526009602052611ebe604060805120614974565b611ec882866145d1565b52611ed381856145d1565b5001611e90565b826040518091602082016020835281518091526040830190602060408260051b860101930191608051905b828210611f1457505050500390f35b91936020611f4f827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851613fe9565b9601920192018594939192611f05565b806060602080938701015201611e83565b346102c15761133b611dd9611f8436614166565b505050915061481e565b346102c15760206003193601126102c15760043567ffffffffffffffff81116102c15760a060031982360301126102c157611fc7614805565b5060405160209190611fd98382613f6e565b6080519052611fe6614805565b5060848101611ff48161446d565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691160361245f5750602481019077ffffffffffffffff0000000000000000000000000000000061205b836143c8565b60801b16604051907f2cbc26bb0000000000000000000000000000000000000000000000000000000082526004820152838160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156119235760805191612442575b50611b72576120ea6044820161446d565b7f00000000000000000000000000000000000000000000000000000000000000006123e2575b5067ffffffffffffffff612123836143c8565b1661213b816000526007602052604060002054151590565b15611ae3578373ffffffffffffffffffffffffffffffffffffffff60045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa80156119235760805190612384575b73ffffffffffffffffffffffffffffffffffffffff9150163303611a50576064013590819067ffffffffffffffff6121d4826143c8565b168060005260088552612224604060002073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016948591615a4a565b6040805173ffffffffffffffffffffffffffffffffffffffff85168152602081018690527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789449190a2813b156102c157604051927f42966c6800000000000000000000000000000000000000000000000000000000845280600485015260805184602481608051875af1938415611923576060611491947ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae109261231097612371575b5067ffffffffffffffff6122f8866143c8565b169360405191825233898301526040820152a26143c8565b9060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168282015281815261234a604082613f6e565b6040519261235784613f1a565b83528183015261133b6040519282849384528301906142ee565b60805161237d91613f6e565b60006122e5565b508381813d83116123db575b61239a8183613f6e565b810103126102c1575173ffffffffffffffffffffffffffffffffffffffff811681036102c15773ffffffffffffffffffffffffffffffffffffffff9061219d565b503d612390565b73ffffffffffffffffffffffffffffffffffffffff1661240f816000526003602052604060002054151590565b612110577fd0d2597600000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b6124599150843d8611611bc657611bb88183613f6e565b846120d9565b611beb73ffffffffffffffffffffffffffffffffffffffff9161446d565b346102c15760606003193601126102c15760043567ffffffffffffffff81116102c1576124ae903690600401614048565b9060243567ffffffffffffffff81116102c1576124cf9036906004016142bd565b9060443567ffffffffffffffff81116102c1576124f09036906004016142bd565b73ffffffffffffffffffffffffffffffffffffffff600a5416331415806125b8575b61123d578386148015906125ae575b612582576080515b8681106125365760805180f35b8061257c61254a610a356001948b8b61440c565b6125558389896147f5565b61257661256e61256686898b6147f5565b923690614342565b913690614342565b91615444565b01612529565b7f568efce200000000000000000000000000000000000000000000000000000000608051526004608051fd5b5080861415612521565b5073ffffffffffffffffffffffffffffffffffffffff60015416331415612512565b346102c1576080516003193601126102c157602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346102c15760206003193601126102c157602061264a67ffffffffffffffff6126366140cc565b166000526007602052604060002054151590565b6040519015158152f35b346102c15760206003193601126102c15760043567ffffffffffffffff81116102c157612688610a1491369060040161409b565b90612691614c33565b614c9b565b346102c15760206003193601126102c1577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174602073ffffffffffffffffffffffffffffffffffffffff6126e7613e6f565b6126ef614c33565b16807fffffffffffffffffffffffff0000000000000000000000000000000000000000600a541617600a55604051908152a160805180f35b346102c1576080516003193601126102c1576080515473ffffffffffffffffffffffffffffffffffffffff811633036127d2577fffffffffffffffffffffffff0000000000000000000000000000000000000000600154913382841617600155166080515573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0608051608051a360805180f35b7f02b543c600000000000000000000000000000000000000000000000000000000608051526004608051fd5b346102c1576080516003193601126102c1576004546005546040805173ffffffffffffffffffffffffffffffffffffffff9093168352602083019190915290f35b346102c1576080516003193601126102c157602073ffffffffffffffffffffffffffffffffffffffff600a5416604051908152f35b346102c15760406003193601126102c15761288d613e6f565b602435612898614c33565b73ffffffffffffffffffffffffffffffffffffffff82169182156109b6577f78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f0238927fffffffffffffffffffffffff0000000000000000000000000000000000000000600454161760045581600555612937604051928392836020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b0390a160805180f35b346102c15761294e3661427c565b612959929192614c33565b67ffffffffffffffff821661297b816000526007602052604060002054151590565b156129965750610a1492612990913691614111565b906151da565b7f1e670e4b00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b346102c1576080516003193601126102c15760206129e161473d565b604051908152f35b346102c157612a11612a196129fd3661422e565b9491612a0a939193614c33565b36916146e9565b9236916146e9565b7f000000000000000000000000000000000000000000000000000000000000000015612b3b576080515b8251811015612ab6578073ffffffffffffffffffffffffffffffffffffffff612a6e600193866145d1565b5116612a7981615cf4565b612a85575b5001612a43565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a184612a7e565b506080515b8151811015610a14578073ffffffffffffffffffffffffffffffffffffffff612ae6600193856145d1565b51168015612b3557612af7816160e9565b612b04575b505b01612abb565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a183612afc565b50612afe565b7f35f4a7b300000000000000000000000000000000000000000000000000000000608051526004608051fd5b346102c15761133b611dd9612b7b36614166565b50505091506145e5565b346102c15760406003193601126102c157612b9e6140cc565b60243567ffffffffffffffff81116102c157602091612bc461264a923690600401614148565b9061448e565b346102c15760406003193601126102c15760043567ffffffffffffffff81116102c157806004019061010060031982360301126102c157612c09614079565b604051909190612c1881613efe565b6080519052612c49612c3f612c3a612c3360c485018761441c565b3691614111565b614fc1565b60648301356150b5565b916084820190612c588261446d565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611bcd57602483019477ffffffffffffffff00000000000000000000000000000000612cbe876143c8565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156119235760805191613116575b50611b725767ffffffffffffffff612d54876143c8565b16612d6c816000526007602052604060002054151590565b15611ae357602073ffffffffffffffffffffffffffffffffffffffff60045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa90811561192357608051916130f7575b5015611a5057612de5866143c8565b90612dfb60a4860192612bc4612c33858561441c565b156130b05750604493929161ffff16159050612ff75767ffffffffffffffff612e23866143c8565b1660805152600d602052612e40604060805120856117438461446d565b7f41a8aa8df7945f0fb8ac5f7d88279638d9dc2ef9a6bf4ec9a53b80681b34aff767ffffffffffffffff612e7661177f886143c8565b6040805173ffffffffffffffffffffffffffffffffffffffff929092168252602082018890529190931692a25b0191612eae8361446d565b9073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691823b156102c1576040517f40c10f1900000000000000000000000000000000000000000000000000000000815260805173ffffffffffffffffffffffffffffffffffffffff909216600482015260248101859052908180604481010381608051875af1801561192357612fde575b50608067ffffffffffffffff60209573ffffffffffffffffffffffffffffffffffffffff612fac612fa67ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0966143c8565b9261446d565b60405196875233898801521660408601528560608601521692a260405190612fd382613efe565b815260405190518152f35b608051612fea91613f6e565b6080516102c15784612f55565b5067ffffffffffffffff61300a856143c8565b16806080515260086020527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c848061308360026040608051200173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615a4a565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a2612ea3565b6130ba925061441c565b611d606040519283927f24eb47e5000000000000000000000000000000000000000000000000000000008452602060048501526024840191614a34565b613110915060203d602011611bc657611bb88183613f6e565b87612dd6565b61312f915060203d602011611bc657611bb88183613f6e565b87612d3d565b346102c15760206003193601126102c15760043567ffffffffffffffff81116102c157806004019061010060031982360301126102c15760405161317881613efe565b608051905260405161318981613efe565b60805190526131a4612c3f612c3a612c3360c485018661441c565b90608481016131b28161446d565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811691160361245f5750602481019277ffffffffffffffff00000000000000000000000000000000613219856143c8565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611923576080519161350e575b50611b725767ffffffffffffffff6132af856143c8565b166132c7816000526007602052604060002054151590565b15611ae357602073ffffffffffffffffffffffffffffffffffffffff60045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa90811561192357608051916134ef575b5015611a5057613340846143c8565b9061335660a4840192612bc4612c33858561441c565b156130b0575082916044915067ffffffffffffffff613374866143c8565b16806080515260086020526133ca60026040608051200173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016958691615a4a565b6040805173ffffffffffffffffffffffffffffffffffffffff86168152602081018790527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a2019261341d8461446d565b823b156102c1576040517f40c10f1900000000000000000000000000000000000000000000000000000000815260805173ffffffffffffffffffffffffffffffffffffffff909216600482015260248101859052908180604481010381608051875af180156119235760209573ffffffffffffffffffffffffffffffffffffffff612fac612fa67ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09660809667ffffffffffffffff966134de575b506143c8565b87516134e991613f6e565b8b6134d8565b613508915060203d602011611bc657611bb88183613f6e565b85613331565b613527915060203d602011611bc657611bb88183613f6e565b85613298565b346102c15760a06003193601126102c157613546613e6f565b5060243567ffffffffffffffff8116908190036102c15760443567ffffffffffffffff81116102c15760031960a091360301126102c15761358561408a565b5060843567ffffffffffffffff81116102c1576135a69036906004016140e3565b50506040516135b481613eb3565b608051815260805160208201526080516040820152606060805191015260805152600f60205260806040815120604051906135ee82613eb3565b5463ffffffff808216928381528160208201818560201c16815260ff60606040850194848860401c168652019560601c161515855260405195865251166020850152511660408301525115156060820152f35b346102c15760606003193601126102c15760043561ffff8116908190036102c15761366a614079565b9060443567ffffffffffffffff81116102c15761368b90369060040161409b565b90613694614c33565b61ffff84169361271085101561371b5783927f52aa194b292c8bfb5aaca8ee2000a965c3a051b306ff841873b16147526a39ba959261370a926040967fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000063ffff0000600b549360101b1692161717600b55614c9b565b82519182526020820152a160805180f35b847f95f3517a00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b346102c15760406003193601126102c15760043567ffffffffffffffff81116102c157366023820112156102c157806004013567ffffffffffffffff81116102c15760248201916024369160a084020101116102c15760243567ffffffffffffffff81116102c1576137c1903690600401614048565b9190926137cc614c33565b6080515b828110613848575050506080515b8181106137eb5760805180f35b8067ffffffffffffffff613805610a35600194868861440c565b168060805152600f602052608051604060805120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8608051608051a2016137de565b80613859610a356001938686614389565b7f56f77aeff2def50c8b5f5a0df3bab7183df09bf36c6feba496bb42551db77d706080613887848888614389565b9261399761395d63ffffffff61398c6139508261398167ffffffffffffffff60208c0198169a8b8a5152600f60205260408a5120836138c58b6143dd565b169181549060408101937fffffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffff67ffffffff00000000613902876143dd565b60201b16918f6cff0000000000000000000000007fffffffffffffffffffffffffffffffffffffffff000000000000000000000000916bffffffff0000000000000000606088019d8e6143dd565b60401b1696019e8f6143ee565b151560601b16951617161717179055826139796040519a6143fb565b1689526143fb565b1660208701526143fb565b166040840152614318565b15156060820152a2016137d0565b346102c1576080516003193601126102c157602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102c15760206003193601126102c15760206139ff613e6f565b6040517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff9081169216919091148152f35b346102c1576080516003193601126102c157602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102c1576080516003193601126102c1576040805161133b91613abb9082613f6e565b601b81527f4275726e4d696e74546f6b656e506f6f6c20312e362e332d64657600000000006020820152604051918291602083526020830190613fe9565b346102c15760206003193601126102c157613b12613e6f565b613b1a614c33565b613b2261473d565b9081613b2e5760805180f35b6040517fa9059cbb000000000000000000000000000000000000000000000000000000006020820190815273ffffffffffffffffffffffffffffffffffffffff831660248301526044808301859052825290613c4390613b8f606482613f6e565b60408051909390917f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169190613bda8685613f6e565b602084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602085015260805191608051915190608051855af13d15613d36573d91613c2583613faf565b92613c3287519485613f6e565b83526080513d90602085013e6161f8565b805180613c95575b505073ffffffffffffffffffffffffffffffffffffffff7f738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b85959992602092519485521692a28080610a14565b90602080613ca79383010191016151c2565b15613cb3578380613c4b565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b6060916161f8565b346102c15760206003193601126102c157600435907fffffffff0000000000000000000000000000000000000000000000000000000082168092036102c157817ff208a58f0000000000000000000000000000000000000000000000000000000060209314908115613e45575b8115613e1b575b8115613df1575b8115613dc7575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483613dc0565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150613db9565b7f1ef5498f0000000000000000000000000000000000000000000000000000000081149150613db2565b7faff2afbf0000000000000000000000000000000000000000000000000000000081149150613dab565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361128f57565b359073ffffffffffffffffffffffffffffffffffffffff8216820361128f57565b6080810190811067ffffffffffffffff821117613ecf57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff821117613ecf57604052565b6040810190811067ffffffffffffffff821117613ecf57604052565b60a0810190811067ffffffffffffffff821117613ecf57604052565b6060810190811067ffffffffffffffff821117613ecf57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117613ecf57604052565b67ffffffffffffffff8111613ecf57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b8481106140335750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201613ff4565b9181601f8401121561128f5782359167ffffffffffffffff831161128f576020808501948460051b01011161128f57565b6024359061ffff8216820361128f57565b6064359061ffff8216820361128f57565b9181601f8401121561128f5782359167ffffffffffffffff831161128f5760208085019460e0850201011161128f57565b6004359067ffffffffffffffff8216820361128f57565b9181601f8401121561128f5782359167ffffffffffffffff831161128f576020838186019501011161128f57565b92919261411d82613faf565b9161412b6040519384613f6e565b82948184528183011161128f578281602093846000960137010152565b9080601f8301121561128f5781602061416393359101614111565b90565b60a060031982011261128f5760043573ffffffffffffffffffffffffffffffffffffffff8116810361128f579160243567ffffffffffffffff8116810361128f57916044359160643561ffff8116810361128f57916084359067ffffffffffffffff821161128f576141da916004016140e3565b9091565b602060408183019282815284518094520192019060005b8181106142025750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016141f5565b604060031982011261128f5760043567ffffffffffffffff811161128f578161425991600401614048565b929092916024359067ffffffffffffffff821161128f576141da91600401614048565b90604060031983011261128f5760043567ffffffffffffffff8116810361128f57916024359067ffffffffffffffff821161128f576141da916004016140e3565b9181601f8401121561128f5782359167ffffffffffffffff831161128f576020808501946060850201011161128f57565b6141639160206143078351604084526040840190613fe9565b920151906020818403910152613fe9565b3590811515820361128f57565b35906fffffffffffffffffffffffffffffffff8216820361128f57565b919082606091031261128f5760405161435a81613f52565b604061438481839561436b81614318565b855261437960208201614325565b602086015201614325565b910152565b91908110156143995760a0020190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3567ffffffffffffffff8116810361128f5790565b3563ffffffff8116810361128f5790565b35801515810361128f5790565b359063ffffffff8216820361128f57565b91908110156143995760051b0190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561128f570180359067ffffffffffffffff821161128f5760200191813603831361128f57565b3573ffffffffffffffffffffffffffffffffffffffff8116810361128f5790565b9067ffffffffffffffff61416392166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b906040519182815491828252602082019060005260206000209260005b8181106144ff5750506144fd92500383613f6e565b565b845473ffffffffffffffffffffffffffffffffffffffff168352600194850194879450602090930192016144e8565b9190820180921161453b57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b67ffffffffffffffff8111613ecf5760051b60200190565b9061458c8261456a565b6145996040519182613f6e565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06145c7829461456a565b0190602036910137565b80518210156143995760209160051b010190565b67ffffffffffffffff16600052600e6020526040600020614605816144cb565b9160055480151591826146de575b505061461d575090565b600161462991016144cb565b908151806146375750905090565b61464561464a91835161452e565b614582565b9260005b825181101561468c578073ffffffffffffffffffffffffffffffffffffffff614679600193866145d1565b511661468582886145d1565b520161464e565b509160005b81518110156146d9578073ffffffffffffffffffffffffffffffffffffffff6146bc600193856145d1565b51166146d26146cc83875161452e565b886145d1565b5201614691565b505050565b101590503880614613565b9291906146f58161456a565b936147036040519586613f6e565b602085838152019160051b810192831161128f57905b82821061472557505050565b6020809161473284613e92565b815201910190614719565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156147e9576000916147ba575090565b90506020813d6020116147e1575b816147d560209383613f6e565b8101031261128f575190565b3d91506147c8565b6040513d6000823e3d90fd5b9190811015614399576060020190565b6040519061481282613f1a565b60606020838281520152565b67ffffffffffffffff16600052600e6020526040600020906003614844600284016144cb565b920190614850826144cb565b506005548015159182614916575b5050614868575090565b614871906144cb565b9081518061487f5750905090565b61464561488d91835161452e565b9260005b82518110156148cf578073ffffffffffffffffffffffffffffffffffffffff6148bc600193866145d1565b51166148c882886145d1565b5201614891565b509160005b81518110156146d9578073ffffffffffffffffffffffffffffffffffffffff6148ff600193856145d1565b511661490f6146cc83875161452e565b52016148d4565b10159050388061485e565b90600182811c9216801561496a575b602083101461493b57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691614930565b906040519182600082549261498884614921565b80845293600181169081156149f457506001146149ad575b506144fd92500383613f6e565b90506000929192526020600020906000915b8183106149d85750509060206144fd92820101386149a0565b60209193508060019154838589010152019101909184926149bf565b602093506144fd9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b820101386149a0565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b60405190614a8082613f36565b60006080838281528260208201528260408201528260608201520152565b90604051614aab81613f36565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff1660005260086020526141636004604060002001614974565b91908110156143995760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff618136030182121561128f570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561128f570180359067ffffffffffffffff821161128f57602001918160051b3603831361128f57565b8181029291811591840414171561453b57565b818110614bd8575050565b60008155600101614bcd565b9160209082815201919060005b818110614bfe5750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff614c2588613e92565b168152019401929101614bf1565b73ffffffffffffffffffffffffffffffffffffffff600154163303614c5457565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b356fffffffffffffffffffffffffffffffff8116810361128f5790565b9160005b82811015614fbb5760e0810284016000614cb8826143c8565b9067ffffffffffffffff821691614cdc836000526007602052604060002054151590565b15614f8f57614da59260408593614d50614d4a94614d4a614d10602060019c9b019261171f614d0b3686614342565b6156d0565b91825463ffffffff8160801c16159081614f71575b81614f62575b81614f47575b81614f38575b5080614f29575b614e9e575b3690614342565b90615817565b6080850192614d62614d0b3686614342565b8152600d6020522092835463ffffffff8160801c16159081614e80575b81614e71575b81614e56575b81614e47575b5080614e38575b614dab575b503690614342565b01614c9f565b614dc860a06fffffffffffffffffffffffffffffffff9201614c7e565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff161717835538614d9d565b50614e42826143ee565b614d98565b60ff915060a01c161538614d91565b6fffffffffffffffffffffffffffffffff8116159150614d8b565b8589015460801c159150614d85565b858901546fffffffffffffffffffffffffffffffff16159150614d7f565b6fffffffffffffffffffffffffffffffff614eba878b01614c7e565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff1617178355614d43565b50614f33816143ee565b614d3e565b60ff915060a01c161538614d37565b6fffffffffffffffffffffffffffffffff8116159150614d31565b848e015460801c159150614d2b565b848e01546fffffffffffffffffffffffffffffffff16159150614d25565b506024917f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50915050565b8051801561503157602003614ff357805160208281019183018390031261128f57519060ff8211614ff3575060ff1690565b611d60906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190613fe9565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff821161453b57565b60ff16604d811161453b57600a0a90565b8115615086570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff8116928284146151bb5782841161519157906150fa91615057565b91604d60ff8416118015615158575b6151225750509061511c6141639261506b565b90614bba565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b506151628361506b565b8015615086577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048411615109565b61519a91615057565b91604d60ff841611615122575050906151b56141639261506b565b9061507c565b5050505090565b9081602091031261128f5751801515810361128f5790565b9080511561541a5767ffffffffffffffff8151602083012092169182600052600860205261520f8160056040600020016161a3565b156153d65760005260096020526040600020815167ffffffffffffffff8111613ecf5761523c8254614921565b601f81116153a4575b506020601f82116001146152de57916152b8827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea95936152ce956000916152d3575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055604051918291602083526020830190613fe9565b0390a2565b905084015138615287565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b81811061538c5750926152ce9492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610615355575b5050811b019055611496565b8501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880615349565b9192602060018192868a01518155019401920161530e565b6153d090836000526020600020601f840160051c8101916020851061094057601f0160051c0190614bcd565b38615245565b5090611d606040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190613fe9565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b67ffffffffffffffff166000818152600760205260409020549092919015615546579161554360e09261550f8561549b7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b976156d0565b8460005260086020526154b2816040600020615817565b6154bb836156d0565b8460005260086020526154d5836002604060002001615817565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b9190820391821161453b57565b615589614a73565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff80835116916155e660208501936155e06155d363ffffffff87511642615574565b8560808901511690614bba565b9061452e565b808210156155ff57505b16825263ffffffff4216905290565b90506155f0565b805160005b81811061561757505050565b6001810180821161453b575b828110615633575060010161560b565b73ffffffffffffffffffffffffffffffffffffffff61565283866145d1565b511673ffffffffffffffffffffffffffffffffffffffff61567383876145d1565b51161461568257600101615623565b73ffffffffffffffffffffffffffffffffffffffff6156a183866145d1565b51167fa1726e400000000000000000000000000000000000000000000000000000000060005260045260246000fd5b805115615770576fffffffffffffffffffffffffffffffff6040820151166fffffffffffffffffffffffffffffffff6020830151161061570d5750565b60649061576e604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff604082015116158015906157f8575b6157975750565b60649061576e604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff6020820151161515615790565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1991615950606092805461585463ffffffff8260801c1642615574565b908161598f575b50506fffffffffffffffffffffffffffffffff600181602086015116928281541680851060001461598757508280855b16167fffffffffffffffffffffffffffffffff000000000000000000000000000000008254161781556159048651151582907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b60408601517fffffffffffffffffffffffffffffffff0000000000000000000000000000000060809190911b16939092166fffffffffffffffffffffffffffffffff1692909217910155565b61554360405180926fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b83809161588b565b6fffffffffffffffffffffffffffffffff916159c48392836159bd6001880154948286169560801c90614bba565b911661452e565b80821015615a4357505b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff9290911692909216167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116174260801b73ffffffff0000000000000000000000000000000016178155388061585b565b90506159ce565b9182549060ff8260a01c16158015615c89575b615c83576fffffffffffffffffffffffffffffffff82169160018501908154615aa263ffffffff6fffffffffffffffffffffffffffffffff83169360801c1642615574565b9081615be5575b5050848110615b995750838310615b03575050615ad86fffffffffffffffffffffffffffffffff928392615574565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b5460801c91615b128185615574565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81019080821161453b57615b60615b659273ffffffffffffffffffffffffffffffffffffffff9661452e565b61507c565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828573ffffffffffffffffffffffffffffffffffffffff927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611615c5957615c00926155e09160801c90614bba565b80841015615c545750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178655923880615aa9565b615c0b565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b508215615a5d565b906040519182815491828252602082019060005260206000209260005b818110615cc35750506144fd92500383613f6e565b8454835260019485019487945060209093019201615cae565b80548210156143995760005260206000200190600090565b6000818152600360205260409020548015615e83577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161453b57600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161453b57818103615e14575b5050506002548015615de5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01615da2816002615cdc565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b615e6b615e25615e36936002615cdc565b90549060031b1c9283926002615cdc565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526003602052604060002055388080615d69565b5050600090565b6000818152600760205260409020548015615e83577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161453b57600654907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161453b57818103615f7b575b5050506006548015615de5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01615f38816006615cdc565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600655600052600760205260006040812055600190565b615f9d615f8c615e36936006615cdc565b90549060031b1c9283926006615cdc565b90556000526007602052604060002055388080615eff565b90600182019181600052826020526040600020548015156000146160e0577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161453b578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161453b578181036160a9575b50505080548015615de5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019061606a8282615cdc565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b6160c96160b9615e369386615cdc565b90549060031b1c92839286615cdc565b905560005283602052604060002055388080616032565b50505050600090565b806000526003602052604060002054156000146161435760025468010000000000000000811015613ecf5761612a615e368260018594016002556002615cdc565b9055600254906000526003602052604060002055600190565b50600090565b806000526007602052604060002054156000146161435760065468010000000000000000811015613ecf5761618a615e368260018594016006556006615cdc565b9055600654906000526007602052604060002055600190565b6000828152600182016020526040902054615e835780549068010000000000000000821015613ecf57826161e1615e36846001809601855584615cdc565b905580549260005201602052604060002055600190565b91929015616273575081511561620c575090565b3b156162155790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156162865750805190602001fd5b611d60906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190613fe956fea164736f6c634300081a000a",
}

var BurnMintTokenPoolABI = BurnMintTokenPoolMetaData.ABI

var BurnMintTokenPoolBin = BurnMintTokenPoolMetaData.Bin

func DeployBurnMintTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, localTokenDecimals uint8, allowlist []common.Address, rmnProxy common.Address, router common.Address) (common.Address, *types.Transaction, *BurnMintTokenPool, error) {
	parsed, err := BurnMintTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BurnMintTokenPoolBin), backend, token, localTokenDecimals, allowlist, rmnProxy, router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnMintTokenPool{address: address, abi: *parsed, BurnMintTokenPoolCaller: BurnMintTokenPoolCaller{contract: contract}, BurnMintTokenPoolTransactor: BurnMintTokenPoolTransactor{contract: contract}, BurnMintTokenPoolFilterer: BurnMintTokenPoolFilterer{contract: contract}}, nil
}

type BurnMintTokenPool struct {
	address common.Address
	abi     abi.ABI
	BurnMintTokenPoolCaller
	BurnMintTokenPoolTransactor
	BurnMintTokenPoolFilterer
}

type BurnMintTokenPoolCaller struct {
	contract *bind.BoundContract
}

type BurnMintTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type BurnMintTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type BurnMintTokenPoolSession struct {
	Contract     *BurnMintTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type BurnMintTokenPoolCallerSession struct {
	Contract *BurnMintTokenPoolCaller
	CallOpts bind.CallOpts
}

type BurnMintTokenPoolTransactorSession struct {
	Contract     *BurnMintTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type BurnMintTokenPoolRaw struct {
	Contract *BurnMintTokenPool
}

type BurnMintTokenPoolCallerRaw struct {
	Contract *BurnMintTokenPoolCaller
}

type BurnMintTokenPoolTransactorRaw struct {
	Contract *BurnMintTokenPoolTransactor
}

func NewBurnMintTokenPool(address common.Address, backend bind.ContractBackend) (*BurnMintTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(BurnMintTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindBurnMintTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPool{address: address, abi: abi, BurnMintTokenPoolCaller: BurnMintTokenPoolCaller{contract: contract}, BurnMintTokenPoolTransactor: BurnMintTokenPoolTransactor{contract: contract}, BurnMintTokenPoolFilterer: BurnMintTokenPoolFilterer{contract: contract}}, nil
}

func NewBurnMintTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*BurnMintTokenPoolCaller, error) {
	contract, err := bindBurnMintTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolCaller{contract: contract}, nil
}

func NewBurnMintTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnMintTokenPoolTransactor, error) {
	contract, err := bindBurnMintTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolTransactor{contract: contract}, nil
}

func NewBurnMintTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnMintTokenPoolFilterer, error) {
	contract, err := bindBurnMintTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolFilterer{contract: contract}, nil
}

func bindBurnMintTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BurnMintTokenPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintTokenPool.Contract.BurnMintTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_BurnMintTokenPool *BurnMintTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.BurnMintTokenPoolTransactor.contract.Transfer(opts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.BurnMintTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.contract.Transfer(opts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetAccumulatedFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getAccumulatedFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetAccumulatedFees() (*big.Int, error) {
	return _BurnMintTokenPool.Contract.GetAccumulatedFees(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetAccumulatedFees() (*big.Int, error) {
	return _BurnMintTokenPool.Contract.GetAccumulatedFees(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetAllowList() ([]common.Address, error) {
	return _BurnMintTokenPool.Contract.GetAllowList(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetAllowList() ([]common.Address, error) {
	return _BurnMintTokenPool.Contract.GetAllowList(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetAllowListEnabled() (bool, error) {
	return _BurnMintTokenPool.Contract.GetAllowListEnabled(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetAllowListEnabled() (bool, error) {
	return _BurnMintTokenPool.Contract.GetAllowListEnabled(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getCurrentInboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnMintTokenPool.Contract.GetCurrentInboundRateLimiterState(&_BurnMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnMintTokenPool.Contract.GetCurrentInboundRateLimiterState(&_BurnMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getCurrentOutboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnMintTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_BurnMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnMintTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_BurnMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

	error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getDynamicConfig")

	outstruct := new(GetDynamicConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Router = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ThresholdAmountForAdditionalCCVs = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _BurnMintTokenPool.Contract.GetDynamicConfig(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _BurnMintTokenPool.Contract.GetDynamicConfig(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getRateLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetRateLimitAdmin() (common.Address, error) {
	return _BurnMintTokenPool.Contract.GetRateLimitAdmin(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetRateLimitAdmin() (common.Address, error) {
	return _BurnMintTokenPool.Contract.GetRateLimitAdmin(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getRemotePools", remoteChainSelector)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _BurnMintTokenPool.Contract.GetRemotePools(&_BurnMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _BurnMintTokenPool.Contract.GetRemotePools(&_BurnMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getRemoteToken", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _BurnMintTokenPool.Contract.GetRemoteToken(&_BurnMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _BurnMintTokenPool.Contract.GetRemoteToken(&_BurnMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetRequiredInboundCCVs(opts *bind.CallOpts, arg0 common.Address, sourceChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getRequiredInboundCCVs", arg0, sourceChainSelector, amount, arg3, arg4)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetRequiredInboundCCVs(arg0 common.Address, sourceChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _BurnMintTokenPool.Contract.GetRequiredInboundCCVs(&_BurnMintTokenPool.CallOpts, arg0, sourceChainSelector, amount, arg3, arg4)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetRequiredInboundCCVs(arg0 common.Address, sourceChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _BurnMintTokenPool.Contract.GetRequiredInboundCCVs(&_BurnMintTokenPool.CallOpts, arg0, sourceChainSelector, amount, arg3, arg4)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetRequiredOutboundCCVs(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getRequiredOutboundCCVs", arg0, destChainSelector, amount, arg3, arg4)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetRequiredOutboundCCVs(arg0 common.Address, destChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _BurnMintTokenPool.Contract.GetRequiredOutboundCCVs(&_BurnMintTokenPool.CallOpts, arg0, destChainSelector, amount, arg3, arg4)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetRequiredOutboundCCVs(arg0 common.Address, destChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _BurnMintTokenPool.Contract.GetRequiredOutboundCCVs(&_BurnMintTokenPool.CallOpts, arg0, destChainSelector, amount, arg3, arg4)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetRmnProxy() (common.Address, error) {
	return _BurnMintTokenPool.Contract.GetRmnProxy(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetRmnProxy() (common.Address, error) {
	return _BurnMintTokenPool.Contract.GetRmnProxy(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetSupportedChains() ([]uint64, error) {
	return _BurnMintTokenPool.Contract.GetSupportedChains(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetSupportedChains() ([]uint64, error) {
	return _BurnMintTokenPool.Contract.GetSupportedChains(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetToken() (common.Address, error) {
	return _BurnMintTokenPool.Contract.GetToken(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _BurnMintTokenPool.Contract.GetToken(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetTokenDecimals() (uint8, error) {
	return _BurnMintTokenPool.Contract.GetTokenDecimals(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetTokenDecimals() (uint8, error) {
	return _BurnMintTokenPool.Contract.GetTokenDecimals(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getTokenTransferFeeConfig", arg0, destChainSelector, arg2, arg3, arg4)

	if err != nil {
		return *new(IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolV2TokenTransferFeeConfig)).(*IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _BurnMintTokenPool.Contract.GetTokenTransferFeeConfig(&_BurnMintTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _BurnMintTokenPool.Contract.GetTokenTransferFeeConfig(&_BurnMintTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "isRemotePool", remoteChainSelector, remotePoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _BurnMintTokenPool.Contract.IsRemotePool(&_BurnMintTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _BurnMintTokenPool.Contract.IsRemotePool(&_BurnMintTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _BurnMintTokenPool.Contract.IsSupportedChain(&_BurnMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _BurnMintTokenPool.Contract.IsSupportedChain(&_BurnMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) IsSupportedToken(token common.Address) (bool, error) {
	return _BurnMintTokenPool.Contract.IsSupportedToken(&_BurnMintTokenPool.CallOpts, token)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _BurnMintTokenPool.Contract.IsSupportedToken(&_BurnMintTokenPool.CallOpts, token)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) Owner() (common.Address, error) {
	return _BurnMintTokenPool.Contract.Owner(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) Owner() (common.Address, error) {
	return _BurnMintTokenPool.Contract.Owner(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintTokenPool.Contract.SupportsInterface(&_BurnMintTokenPool.CallOpts, interfaceId)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintTokenPool.Contract.SupportsInterface(&_BurnMintTokenPool.CallOpts, interfaceId)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) TypeAndVersion() (string, error) {
	return _BurnMintTokenPool.Contract.TypeAndVersion(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) TypeAndVersion() (string, error) {
	return _BurnMintTokenPool.Contract.TypeAndVersion(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.AcceptOwnership(&_BurnMintTokenPool.TransactOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.AcceptOwnership(&_BurnMintTokenPool.TransactOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "addRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.AddRemotePool(&_BurnMintTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.AddRemotePool(&_BurnMintTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ApplyAllowListUpdates(&_BurnMintTokenPool.TransactOpts, removes, adds)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ApplyAllowListUpdates(&_BurnMintTokenPool.TransactOpts, removes, adds)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) ApplyCCVConfigUpdates(opts *bind.TransactOpts, ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "applyCCVConfigUpdates", ccvConfigArgs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) ApplyCCVConfigUpdates(ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ApplyCCVConfigUpdates(&_BurnMintTokenPool.TransactOpts, ccvConfigArgs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) ApplyCCVConfigUpdates(ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ApplyCCVConfigUpdates(&_BurnMintTokenPool.TransactOpts, ccvConfigArgs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "applyChainUpdates", remoteChainSelectorsToRemove, chainsToAdd)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ApplyChainUpdates(&_BurnMintTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ApplyChainUpdates(&_BurnMintTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) ApplyFinalityConfigUpdates(opts *bind.TransactOpts, finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "applyFinalityConfigUpdates", finalityThreshold, customFinalityTransferFeeBps, rateLimitConfigArgs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) ApplyFinalityConfigUpdates(finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ApplyFinalityConfigUpdates(&_BurnMintTokenPool.TransactOpts, finalityThreshold, customFinalityTransferFeeBps, rateLimitConfigArgs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) ApplyFinalityConfigUpdates(finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ApplyFinalityConfigUpdates(&_BurnMintTokenPool.TransactOpts, finalityThreshold, customFinalityTransferFeeBps, rateLimitConfigArgs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "applyTokenTransferFeeConfigUpdates", tokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_BurnMintTokenPool.TransactOpts, tokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_BurnMintTokenPool.TransactOpts, tokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.LockOrBurn(&_BurnMintTokenPool.TransactOpts, lockOrBurnIn)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.LockOrBurn(&_BurnMintTokenPool.TransactOpts, lockOrBurnIn)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, finality, arg2)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.LockOrBurn0(&_BurnMintTokenPool.TransactOpts, lockOrBurnIn, finality, arg2)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.LockOrBurn0(&_BurnMintTokenPool.TransactOpts, lockOrBurnIn, finality, arg2)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ReleaseOrMint(&_BurnMintTokenPool.TransactOpts, releaseOrMintIn)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ReleaseOrMint(&_BurnMintTokenPool.TransactOpts, releaseOrMintIn)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn, finality)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ReleaseOrMint0(&_BurnMintTokenPool.TransactOpts, releaseOrMintIn, finality)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ReleaseOrMint0(&_BurnMintTokenPool.TransactOpts, releaseOrMintIn, finality)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "removeRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.RemoveRemotePool(&_BurnMintTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.RemoveRemotePool(&_BurnMintTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "setChainRateLimiterConfig", remoteChainSelector, outboundConfig, inboundConfig)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.SetChainRateLimiterConfig(&_BurnMintTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.SetChainRateLimiterConfig(&_BurnMintTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "setChainRateLimiterConfigs", remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.SetChainRateLimiterConfigs(&_BurnMintTokenPool.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.SetChainRateLimiterConfigs(&_BurnMintTokenPool.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) SetCustomFinalityRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "setCustomFinalityRateLimitConfig", rateLimitConfigArgs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) SetCustomFinalityRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.SetCustomFinalityRateLimitConfig(&_BurnMintTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) SetCustomFinalityRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.SetCustomFinalityRateLimitConfig(&_BurnMintTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) SetDynamicConfig(opts *bind.TransactOpts, router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "setDynamicConfig", router, thresholdAmountForAdditionalCCVs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) SetDynamicConfig(router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.SetDynamicConfig(&_BurnMintTokenPool.TransactOpts, router, thresholdAmountForAdditionalCCVs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) SetDynamicConfig(router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.SetDynamicConfig(&_BurnMintTokenPool.TransactOpts, router, thresholdAmountForAdditionalCCVs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "setRateLimitAdmin", rateLimitAdmin)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.SetRateLimitAdmin(&_BurnMintTokenPool.TransactOpts, rateLimitAdmin)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.SetRateLimitAdmin(&_BurnMintTokenPool.TransactOpts, rateLimitAdmin)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.TransferOwnership(&_BurnMintTokenPool.TransactOpts, to)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.TransferOwnership(&_BurnMintTokenPool.TransactOpts, to)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) WithdrawFees(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "withdrawFees", recipient)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) WithdrawFees(recipient common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.WithdrawFees(&_BurnMintTokenPool.TransactOpts, recipient)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) WithdrawFees(recipient common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.WithdrawFees(&_BurnMintTokenPool.TransactOpts, recipient)
}

type BurnMintTokenPoolAllowListAddIterator struct {
	Event *BurnMintTokenPoolAllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolAllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolAllowListAdd)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolAllowListAdd)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolAllowListAddIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolAllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolAllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterAllowListAdd(opts *bind.FilterOpts) (*BurnMintTokenPoolAllowListAddIterator, error) {

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolAllowListAddIterator{contract: _BurnMintTokenPool.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAllowListAdd) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolAllowListAdd)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseAllowListAdd(log types.Log) (*BurnMintTokenPoolAllowListAdd, error) {
	event := new(BurnMintTokenPoolAllowListAdd)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolAllowListRemoveIterator struct {
	Event *BurnMintTokenPoolAllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolAllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolAllowListRemove)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolAllowListRemove)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolAllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolAllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolAllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterAllowListRemove(opts *bind.FilterOpts) (*BurnMintTokenPoolAllowListRemoveIterator, error) {

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolAllowListRemoveIterator{contract: _BurnMintTokenPool.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAllowListRemove) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolAllowListRemove)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseAllowListRemove(log types.Log) (*BurnMintTokenPoolAllowListRemove, error) {
	event := new(BurnMintTokenPoolAllowListRemove)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolCCVConfigUpdatedIterator struct {
	Event *BurnMintTokenPoolCCVConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolCCVConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolCCVConfigUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolCCVConfigUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolCCVConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolCCVConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolCCVConfigUpdated struct {
	RemoteChainSelector    uint64
	OutboundCCVs           []common.Address
	AdditionalOutboundCCVs []common.Address
	InboundCCVs            []common.Address
	AdditionalInboundCCVs  []common.Address
	Raw                    types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolCCVConfigUpdatedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolCCVConfigUpdatedIterator{contract: _BurnMintTokenPool.contract, event: "CCVConfigUpdated", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolCCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolCCVConfigUpdated)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseCCVConfigUpdated(log types.Log) (*BurnMintTokenPoolCCVConfigUpdated, error) {
	event := new(BurnMintTokenPoolCCVConfigUpdated)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolChainAddedIterator struct {
	Event *BurnMintTokenPoolChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolChainAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolChainAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolChainAddedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolChainAdded struct {
	RemoteChainSelector       uint64
	RemoteToken               []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterChainAdded(opts *bind.FilterOpts) (*BurnMintTokenPoolChainAddedIterator, error) {

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolChainAddedIterator{contract: _BurnMintTokenPool.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolChainAdded) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolChainAdded)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseChainAdded(log types.Log) (*BurnMintTokenPoolChainAdded, error) {
	event := new(BurnMintTokenPoolChainAdded)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolChainConfiguredIterator struct {
	Event *BurnMintTokenPoolChainConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolChainConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolChainConfigured)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolChainConfigured)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolChainConfiguredIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolChainConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolChainConfigured struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterChainConfigured(opts *bind.FilterOpts) (*BurnMintTokenPoolChainConfiguredIterator, error) {

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolChainConfiguredIterator{contract: _BurnMintTokenPool.contract, event: "ChainConfigured", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolChainConfigured) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolChainConfigured)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseChainConfigured(log types.Log) (*BurnMintTokenPoolChainConfigured, error) {
	event := new(BurnMintTokenPoolChainConfigured)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolChainRemovedIterator struct {
	Event *BurnMintTokenPoolChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolChainRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolChainRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolChainRemovedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterChainRemoved(opts *bind.FilterOpts) (*BurnMintTokenPoolChainRemovedIterator, error) {

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolChainRemovedIterator{contract: _BurnMintTokenPool.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolChainRemoved) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolChainRemoved)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseChainRemoved(log types.Log) (*BurnMintTokenPoolChainRemoved, error) {
	event := new(BurnMintTokenPoolChainRemoved)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolConfigChangedIterator struct {
	Event *BurnMintTokenPoolConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolConfigChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolConfigChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolConfigChangedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolConfigChanged struct {
	Config RateLimiterConfig
	Raw    types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*BurnMintTokenPoolConfigChangedIterator, error) {

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolConfigChangedIterator{contract: _BurnMintTokenPool.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolConfigChanged) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolConfigChanged)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseConfigChanged(log types.Log) (*BurnMintTokenPoolConfigChanged, error) {
	event := new(BurnMintTokenPoolConfigChanged)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumedIterator struct {
	Event *BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterCustomFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "CustomFinalityOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumedIterator{contract: _BurnMintTokenPool.contract, event: "CustomFinalityOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchCustomFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "CustomFinalityOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumed)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "CustomFinalityOutboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseCustomFinalityOutboundRateLimitConsumed(log types.Log) (*BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumed, error) {
	event := new(BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumed)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "CustomFinalityOutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator struct {
	Event *BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterCustomFinalityTransferInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "CustomFinalityTransferInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator{contract: _BurnMintTokenPool.contract, event: "CustomFinalityTransferInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchCustomFinalityTransferInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "CustomFinalityTransferInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "CustomFinalityTransferInboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseCustomFinalityTransferInboundRateLimitConsumed(log types.Log) (*BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed, error) {
	event := new(BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "CustomFinalityTransferInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolDynamicConfigSetIterator struct {
	Event *BurnMintTokenPoolDynamicConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolDynamicConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolDynamicConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolDynamicConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolDynamicConfigSetIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolDynamicConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolDynamicConfigSet struct {
	Router                           common.Address
	ThresholdAmountForAdditionalCCVs *big.Int
	Raw                              types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterDynamicConfigSet(opts *bind.FilterOpts) (*BurnMintTokenPoolDynamicConfigSetIterator, error) {

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolDynamicConfigSetIterator{contract: _BurnMintTokenPool.contract, event: "DynamicConfigSet", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolDynamicConfigSet) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolDynamicConfigSet)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseDynamicConfigSet(log types.Log) (*BurnMintTokenPoolDynamicConfigSet, error) {
	event := new(BurnMintTokenPoolDynamicConfigSet)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolFinalityConfigUpdatedIterator struct {
	Event *BurnMintTokenPoolFinalityConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolFinalityConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolFinalityConfigUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolFinalityConfigUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolFinalityConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolFinalityConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolFinalityConfigUpdated struct {
	FinalityConfig               uint16
	CustomFinalityTransferFeeBps uint16
	Raw                          types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterFinalityConfigUpdated(opts *bind.FilterOpts) (*BurnMintTokenPoolFinalityConfigUpdatedIterator, error) {

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "FinalityConfigUpdated")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolFinalityConfigUpdatedIterator{contract: _BurnMintTokenPool.contract, event: "FinalityConfigUpdated", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchFinalityConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolFinalityConfigUpdated) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "FinalityConfigUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolFinalityConfigUpdated)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "FinalityConfigUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseFinalityConfigUpdated(log types.Log) (*BurnMintTokenPoolFinalityConfigUpdated, error) {
	event := new(BurnMintTokenPoolFinalityConfigUpdated)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "FinalityConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolInboundRateLimitConsumedIterator struct {
	Event *BurnMintTokenPoolInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolInboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolInboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolInboundRateLimitConsumedIterator{contract: _BurnMintTokenPool.contract, event: "InboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolInboundRateLimitConsumed)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseInboundRateLimitConsumed(log types.Log) (*BurnMintTokenPoolInboundRateLimitConsumed, error) {
	event := new(BurnMintTokenPoolInboundRateLimitConsumed)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolLockedOrBurnedIterator struct {
	Event *BurnMintTokenPoolLockedOrBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolLockedOrBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolLockedOrBurned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolLockedOrBurned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolLockedOrBurnedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolLockedOrBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolLockedOrBurned struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolLockedOrBurnedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolLockedOrBurnedIterator{contract: _BurnMintTokenPool.contract, event: "LockedOrBurned", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolLockedOrBurned)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseLockedOrBurned(log types.Log) (*BurnMintTokenPoolLockedOrBurned, error) {
	event := new(BurnMintTokenPoolLockedOrBurned)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolOutboundRateLimitConsumedIterator struct {
	Event *BurnMintTokenPoolOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolOutboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolOutboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolOutboundRateLimitConsumedIterator{contract: _BurnMintTokenPool.contract, event: "OutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolOutboundRateLimitConsumed)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseOutboundRateLimitConsumed(log types.Log) (*BurnMintTokenPoolOutboundRateLimitConsumed, error) {
	event := new(BurnMintTokenPoolOutboundRateLimitConsumed)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolOwnershipTransferRequestedIterator struct {
	Event *BurnMintTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolOwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolOwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolOwnershipTransferRequestedIterator{contract: _BurnMintTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolOwnershipTransferRequested)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*BurnMintTokenPoolOwnershipTransferRequested, error) {
	event := new(BurnMintTokenPoolOwnershipTransferRequested)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolOwnershipTransferredIterator struct {
	Event *BurnMintTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolOwnershipTransferredIterator{contract: _BurnMintTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolOwnershipTransferred)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*BurnMintTokenPoolOwnershipTransferred, error) {
	event := new(BurnMintTokenPoolOwnershipTransferred)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolPoolFeeWithdrawnIterator struct {
	Event *BurnMintTokenPoolPoolFeeWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolPoolFeeWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolPoolFeeWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolPoolFeeWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolPoolFeeWithdrawnIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolPoolFeeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolPoolFeeWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterPoolFeeWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*BurnMintTokenPoolPoolFeeWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "PoolFeeWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolPoolFeeWithdrawnIterator{contract: _BurnMintTokenPool.contract, event: "PoolFeeWithdrawn", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchPoolFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolPoolFeeWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "PoolFeeWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolPoolFeeWithdrawn)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "PoolFeeWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParsePoolFeeWithdrawn(log types.Log) (*BurnMintTokenPoolPoolFeeWithdrawn, error) {
	event := new(BurnMintTokenPoolPoolFeeWithdrawn)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "PoolFeeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolRateLimitAdminSetIterator struct {
	Event *BurnMintTokenPoolRateLimitAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolRateLimitAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolRateLimitAdminSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolRateLimitAdminSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolRateLimitAdminSetIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolRateLimitAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolRateLimitAdminSet struct {
	RateLimitAdmin common.Address
	Raw            types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterRateLimitAdminSet(opts *bind.FilterOpts) (*BurnMintTokenPoolRateLimitAdminSetIterator, error) {

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolRateLimitAdminSetIterator{contract: _BurnMintTokenPool.contract, event: "RateLimitAdminSet", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolRateLimitAdminSet) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolRateLimitAdminSet)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseRateLimitAdminSet(log types.Log) (*BurnMintTokenPoolRateLimitAdminSet, error) {
	event := new(BurnMintTokenPoolRateLimitAdminSet)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolReleasedOrMintedIterator struct {
	Event *BurnMintTokenPoolReleasedOrMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolReleasedOrMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolReleasedOrMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolReleasedOrMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolReleasedOrMintedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolReleasedOrMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolReleasedOrMinted struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Recipient           common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolReleasedOrMintedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolReleasedOrMintedIterator{contract: _BurnMintTokenPool.contract, event: "ReleasedOrMinted", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolReleasedOrMinted)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseReleasedOrMinted(log types.Log) (*BurnMintTokenPoolReleasedOrMinted, error) {
	event := new(BurnMintTokenPoolReleasedOrMinted)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolRemotePoolAddedIterator struct {
	Event *BurnMintTokenPoolRemotePoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolRemotePoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolRemotePoolAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolRemotePoolAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolRemotePoolAddedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolRemotePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolRemotePoolAdded struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolRemotePoolAddedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolRemotePoolAddedIterator{contract: _BurnMintTokenPool.contract, event: "RemotePoolAdded", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolRemotePoolAdded)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseRemotePoolAdded(log types.Log) (*BurnMintTokenPoolRemotePoolAdded, error) {
	event := new(BurnMintTokenPoolRemotePoolAdded)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolRemotePoolRemovedIterator struct {
	Event *BurnMintTokenPoolRemotePoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolRemotePoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolRemotePoolRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolRemotePoolRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolRemotePoolRemovedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolRemotePoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolRemotePoolRemoved struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolRemotePoolRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolRemotePoolRemovedIterator{contract: _BurnMintTokenPool.contract, event: "RemotePoolRemoved", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolRemotePoolRemoved)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseRemotePoolRemoved(log types.Log) (*BurnMintTokenPoolRemotePoolRemoved, error) {
	event := new(BurnMintTokenPoolRemotePoolRemoved)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolTokenTransferFeeConfigDeletedIterator struct {
	Event *BurnMintTokenPoolTokenTransferFeeConfigDeleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolTokenTransferFeeConfigDeletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolTokenTransferFeeConfigDeleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolTokenTransferFeeConfigDeleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolTokenTransferFeeConfigDeletedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolTokenTransferFeeConfigDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolTokenTransferFeeConfigDeleted struct {
	DestChainSelector uint64
	Raw               types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnMintTokenPoolTokenTransferFeeConfigDeletedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolTokenTransferFeeConfigDeletedIterator{contract: _BurnMintTokenPool.contract, event: "TokenTransferFeeConfigDeleted", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolTokenTransferFeeConfigDeleted)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseTokenTransferFeeConfigDeleted(log types.Log) (*BurnMintTokenPoolTokenTransferFeeConfigDeleted, error) {
	event := new(BurnMintTokenPoolTokenTransferFeeConfigDeleted)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolTokenTransferFeeConfigUpdatedIterator struct {
	Event *BurnMintTokenPoolTokenTransferFeeConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolTokenTransferFeeConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolTokenTransferFeeConfigUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolTokenTransferFeeConfigUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolTokenTransferFeeConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolTokenTransferFeeConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolTokenTransferFeeConfigUpdated struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
	Raw                    types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnMintTokenPoolTokenTransferFeeConfigUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolTokenTransferFeeConfigUpdatedIterator{contract: _BurnMintTokenPool.contract, event: "TokenTransferFeeConfigUpdated", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolTokenTransferFeeConfigUpdated)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseTokenTransferFeeConfigUpdated(log types.Log) (*BurnMintTokenPoolTokenTransferFeeConfigUpdated, error) {
	event := new(BurnMintTokenPoolTokenTransferFeeConfigUpdated)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetDynamicConfig struct {
	Router                           common.Address
	ThresholdAmountForAdditionalCCVs *big.Int
}

func (BurnMintTokenPoolAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (BurnMintTokenPoolAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (BurnMintTokenPoolCCVConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d")
}

func (BurnMintTokenPoolChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (BurnMintTokenPoolChainConfigured) Topic() common.Hash {
	return common.HexToHash("0x0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b")
}

func (BurnMintTokenPoolChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (BurnMintTokenPoolConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b2")
}

func (BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x41a8aa8df7945f0fb8ac5f7d88279638d9dc2ef9a6bf4ec9a53b80681b34aff7")
}

func (BurnMintTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f0238")
}

func (BurnMintTokenPoolFinalityConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x52aa194b292c8bfb5aaca8ee2000a965c3a051b306ff841873b16147526a39ba")
}

func (BurnMintTokenPoolInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (BurnMintTokenPoolLockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
}

func (BurnMintTokenPoolOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944")
}

func (BurnMintTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (BurnMintTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (BurnMintTokenPoolPoolFeeWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b859599")
}

func (BurnMintTokenPoolRateLimitAdminSet) Topic() common.Hash {
	return common.HexToHash("0x44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174")
}

func (BurnMintTokenPoolReleasedOrMinted) Topic() common.Hash {
	return common.HexToHash("0xfc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0")
}

func (BurnMintTokenPoolRemotePoolAdded) Topic() common.Hash {
	return common.HexToHash("0x7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea")
}

func (BurnMintTokenPoolRemotePoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76")
}

func (BurnMintTokenPoolTokenTransferFeeConfigDeleted) Topic() common.Hash {
	return common.HexToHash("0x5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8")
}

func (BurnMintTokenPoolTokenTransferFeeConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x56f77aeff2def50c8b5f5a0df3bab7183df09bf36c6feba496bb42551db77d70")
}

func (_BurnMintTokenPool *BurnMintTokenPool) Address() common.Address {
	return _BurnMintTokenPool.address
}

type BurnMintTokenPoolInterface interface {
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

	FilterAllowListAdd(opts *bind.FilterOpts) (*BurnMintTokenPoolAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*BurnMintTokenPoolAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*BurnMintTokenPoolAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*BurnMintTokenPoolAllowListRemove, error)

	FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolCCVConfigUpdatedIterator, error)

	WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolCCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCCVConfigUpdated(log types.Log) (*BurnMintTokenPoolCCVConfigUpdated, error)

	FilterChainAdded(opts *bind.FilterOpts) (*BurnMintTokenPoolChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*BurnMintTokenPoolChainAdded, error)

	FilterChainConfigured(opts *bind.FilterOpts) (*BurnMintTokenPoolChainConfiguredIterator, error)

	WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolChainConfigured) (event.Subscription, error)

	ParseChainConfigured(log types.Log) (*BurnMintTokenPoolChainConfigured, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*BurnMintTokenPoolChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*BurnMintTokenPoolChainRemoved, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*BurnMintTokenPoolConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*BurnMintTokenPoolConfigChanged, error)

	FilterCustomFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumedIterator, error)

	WatchCustomFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomFinalityOutboundRateLimitConsumed(log types.Log) (*BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumed, error)

	FilterCustomFinalityTransferInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator, error)

	WatchCustomFinalityTransferInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomFinalityTransferInboundRateLimitConsumed(log types.Log) (*BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*BurnMintTokenPoolDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*BurnMintTokenPoolDynamicConfigSet, error)

	FilterFinalityConfigUpdated(opts *bind.FilterOpts) (*BurnMintTokenPoolFinalityConfigUpdatedIterator, error)

	WatchFinalityConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolFinalityConfigUpdated) (event.Subscription, error)

	ParseFinalityConfigUpdated(log types.Log) (*BurnMintTokenPoolFinalityConfigUpdated, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolInboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*BurnMintTokenPoolInboundRateLimitConsumed, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolLockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*BurnMintTokenPoolLockedOrBurned, error)

	FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolOutboundRateLimitConsumedIterator, error)

	WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseOutboundRateLimitConsumed(log types.Log) (*BurnMintTokenPoolOutboundRateLimitConsumed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*BurnMintTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*BurnMintTokenPoolOwnershipTransferred, error)

	FilterPoolFeeWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*BurnMintTokenPoolPoolFeeWithdrawnIterator, error)

	WatchPoolFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolPoolFeeWithdrawn, recipient []common.Address) (event.Subscription, error)

	ParsePoolFeeWithdrawn(log types.Log) (*BurnMintTokenPoolPoolFeeWithdrawn, error)

	FilterRateLimitAdminSet(opts *bind.FilterOpts) (*BurnMintTokenPoolRateLimitAdminSetIterator, error)

	WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolRateLimitAdminSet) (event.Subscription, error)

	ParseRateLimitAdminSet(log types.Log) (*BurnMintTokenPoolRateLimitAdminSet, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*BurnMintTokenPoolReleasedOrMinted, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolRemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*BurnMintTokenPoolRemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolRemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*BurnMintTokenPoolRemotePoolRemoved, error)

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnMintTokenPoolTokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*BurnMintTokenPoolTokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnMintTokenPoolTokenTransferFeeConfigUpdatedIterator, error)

	WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigUpdated(log types.Log) (*BurnMintTokenPoolTokenTransferFeeConfigUpdated, error)

	Address() common.Address
}
