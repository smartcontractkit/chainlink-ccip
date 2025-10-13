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

var BurnToAddressMintTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIBurnMintERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"burnAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"additionalOutboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"additionalInboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyFinalityConfigUpdates\",\"inputs\":[{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"destToUseDefaultFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAccumulatedFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBurnAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"burnAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredInboundCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredOutboundCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_burnAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"lockOrBurnOutV1\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustomFinalityRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"additionalOutboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"additionalInboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityTransferInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalityConfigUpdated\",\"inputs\":[{\"name\":\"finalityConfig\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeeWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidFinality\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidFinalityConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenTransferFeeConfigNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x61012080604052346103dd57616936803803809161001d828561045c565b8339810160c0828203126103dd5781516001600160a01b03811692908390036103dd5761004c6020820161047f565b60408201516001600160401b0381116103dd5782019280601f850112156103dd578351936001600160401b038511610446578460051b906020820195610095604051978861045c565b86526020808701928201019283116103dd57602001905b82821061042e575050506100c26060830161048d565b936100db60a06100d46080860161048d565b940161048d565b94331561041d57600180546001600160a01b031916331790558115801561040c575b80156103fb575b6103ea578160209160049360805260c0526040519283809263313ce56760e01b82525afa600091816103a9575b5061037e575b5060a052600480546001600160a01b0319166001600160a01b03929092169190911790558051151560e0819052610261575b50610100526040516162f49081610642823960805181818161154f015281816117e4015281816119cd015281816120f0015281816122f901528181612c5101528181612ea70152818161303c015281816131aa01528181613383015281816139e501528181613a5401528181613b4f015261461f015260a05181818161189b015281816123d9015281816139a10152818161505801526150db015260c051818181610c5e015281816115ea0152818161218b01528181612cec0152613247015260e051818181610c0b0152818161162f015281816121d001526129fc015261010051818181611809015281816123820152613f830152f35b6020604051610270828261045c565b60008152600036813760e0511561036d5760005b81518110156102eb576001906001600160a01b036102a282856104a1565b5116846102ae826104e3565b6102bb575b505001610284565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a138846102b3565b505060005b8251811015610364576001906001600160a01b0361030e82866104a1565b5116801561035e5783610320826105e1565b61032e575b50505b016102f0565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a13883610325565b50610328565b50505038610169565b6335f4a7b360e01b60005260046000fd5b60ff1660ff82168181036103925750610137565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d6020116103e2575b816103c56020938361045c565b810103126103dd576103d69061047f565b9038610131565b600080fd5b3d91506103b8565b630a64406560e11b60005260046000fd5b506001600160a01b03811615610104565b506001600160a01b038416156100fd565b639b15e16f60e01b60005260046000fd5b6020809161043b8461048d565b8152019101906100ac565b634e487b7160e01b600052604160045260246000fd5b601f909101601f19168101906001600160401b0382119082101761044657604052565b519060ff821682036103dd57565b51906001600160a01b03821682036103dd57565b80518210156104b55760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b80548210156104b55760005260206000200190600090565b60008181526003602052604090205480156105da5760001981018181116105c4576002546000198101919082116105c457818103610573575b505050600254801561055d57600019016105378160026104cb565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b6105ac6105846105959360026104cb565b90549060031b1c92839260026104cb565b819391549060031b91821b91600019901b19161790565b9055600052600360205260406000205538808061051c565b634e487b7160e01b600052601160045260246000fd5b5050600090565b8060005260036020526040600020541560001461063b57600254680100000000000000008110156104465761062261059582600185940160025560026104cb565b9055600254906000526003602052604060002055600190565b5060009056fe60a080604052600436101561001357600080fd5b60006080526080513560e01c90816301ffc9a714613b8357508063164e68de14613ada578063181f5a7714613a7857806321df0da714613a26578063240028e8146139c557806324f65ee7146139865780632a10097b1461372c5780632c286daf1461362257806337b192471461350e57806338b39d29146112aa5780633907753714613116578063489a68f214612bab5780634c5ef0ed14612b665780634f71592c14612b4857806354c8a4f3146129ca5780635df45a37146129a657806362ddd3c4146129215780636d3d1a58146128ec5780637437ff9f146128ac57806379ba5097146127d55780637d54534e14612744578063804ba5a9146127025780638926f54f146126bd5780638da5cb5b14612688578063962d40201461252b5780639a4575b9146120725780639f68f67314612054578063a42a7b8b14611edf578063a7cd63b714611e78578063acfecf9114611d4b578063af58d59f14611cff578063b184b94214611c09578063b1c71c65146114c5578063b79465801461148d578063c4bffe2b1461135a578063c75eea9c146112af578063c8de9fe0146112aa578063cf7401f3146110fe578063d966866b14610c82578063dc0bd97114610c30578063e0351e1314610bf2578063e8a1da17146102dd5763f2fde38b146101fe57600080fd5b346102d75760206003193601126102d75773ffffffffffffffffffffffffffffffffffffffff61022c613cb4565b610234614acf565b163381146102ab5760805180547fffffffffffffffffffffffff0000000000000000000000000000000000000000168217815560015473ffffffffffffffffffffffffffffffffffffffff16907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789080a360805180f35b7fdad89dca00000000000000000000000000000000000000000000000000000000608051526004608051fd5b60805180fd5b346102d7576102eb366140c4565b926102f7929192614acf565b608051905b828210610a305750505060805163ffffffff4216917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee181360301915b84811015610a2a576000938160051b830135848112156102d757830190610120823603126102d7576040519561036d87613d7b565b82359067ffffffffffffffff82168203610a2757508652602082013567ffffffffffffffff81116102d75782019436601f870112156102d75785356103b181614400565b966103bf6040519889613db3565b81885260208089019260051b820101903682116102d75760208101925b8284106109f8575050505060208701958652604083013567ffffffffffffffff81116102d75761040f9036908501613fde565b916040880192835261043961042736606087016141d8565b9460608a0195865260c03691016141d8565b956080890196875261044b85516156f3565b61045587516156f3565b835151156109cc5761047167ffffffffffffffff8a511661616c565b156109915767ffffffffffffffff8951166080515260076020526040608051206105b586516fffffffffffffffffffffffffffffffff604082015116906105706fffffffffffffffffffffffffffffffff602083015116915115158360806040516104db81613d7b565b858152602081018b905260408101849052606081018690520152855474ff000000000000000000000000000000000000000091151560a01b919091167fffffffffffffffffffffff0000000000000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff84161773ffffffff0000000000000000000000000000000060808a901b1617178555565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176001830155565b6106db88516fffffffffffffffffffffffffffffffff604082015116906106966fffffffffffffffffffffffffffffffff602083015116915115158360806040516105ff81613d7b565b858152602081018b9052604081018490526060810186905201526002860180547fffffffffffffffffffffff000000000000000000000000000000000000000000166fffffffffffffffffffffffffffffffff85161773ffffffff0000000000000000000000000000000060808b901b161791151560a01b74ff000000000000000000000000000000000000000016919091179055565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176003830155565b6004855191019080519067ffffffffffffffff8211610960576106fe83546147bd565b601f8111610921575b506020906001601f84111461087b579180916107589360805192610870575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b6080515b88518051821015610794579061078e6001926107878367ffffffffffffffff8f511692614467565b51906151fd565b0161075f565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c293919997509561086267ffffffffffffffff600197969498511692519351915161082e6107f960405196879687526101006020880152610100870190613e2e565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a1019392909193610338565b015190508e80610726565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083169184608051528160805120926080515b81811061090957509084600195949392106108d2575b505050811b01905561075b565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558d80806108c5565b929360206001819287860151815501950193016108af565b610950908460805152602060805120601f850160051c81019160208610610956575b601f0160051c0190614a69565b8d610707565b9091508190610943565b7f4e487b71000000000000000000000000000000000000000000000000000000006080515260416004526024608051fd5b67ffffffffffffffff8951167f1d5ad3c500000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b7f14c880ca00000000000000000000000000000000000000000000000000000000608051526004608051fd5b833567ffffffffffffffff81116102d757602091610a1c8392833691870101613fde565b8152019301926103dc565b80fd5b60805180f35b9092919367ffffffffffffffff610a50610a4b8688866142a2565b61425e565b1692610a5b84615ead565b15610bc25783608051526007602052610a7b600560406080512001615cb4565b926080515b8451811015610aba5760019086608051526007602052610ab3600560406080512001610aac8389614467565b5190615fd8565b5001610a80565b5093909491959250806080515260076020526005604060805120608051815560805160018201556080516002820155608051600382015560048101610aff81546147bd565b80610b71575b505001805490608051815581610b4e575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a10190916102fc565b60805152602060805120908101905b81811015610b165760008155600101610b5d565b601f8111600114610b895750600090555b8880610b05565b610ba99082608051526001601f6020608051209201861c82019101614a69565b8060805152600060206080512091556080519055610b82565b837f1e670e4b00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b346102d7576080516003193601126102d75760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b346102d7576080516003193601126102d757602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102d75760206003193601126102d75760043567ffffffffffffffff81116102d757610cb3903690600401613e8d565b610cbb614acf565b608051915b818310610ccd5760805180f35b610cdb610a4b8484846149c2565b610cf3610ce98585856149c2565b6020810190614a02565b9091610d0d610d038787876149c2565b6040810190614a02565b90610d26610d1c8989896149c2565b6060810190614a02565b9091610d40610d368b8b8b6149c2565b6080810190614a02565b949097610d56610d51368a84614582565b615629565b610d64610d51368486614582565b610d72610d51368688614582565b610d80610d5136888c614582565b604051610d8c81613cf8565b610d97368a84614582565b8152610da4368486614582565b6020820152610db4368688614582565b6040820152610dc436888c614582565b606082015267ffffffffffffffff881660805152600d602052604060805120815180519067ffffffffffffffff8211610960576801000000000000000082116109605760209083548385558084106110df575b500182608051526020608051206080515b8381106110b55750505050602082015180519067ffffffffffffffff821161096057680100000000000000008211610960576020906001840154836001860155808410611093575b500160018301608051526020608051206080515b8381106110695750505050604082015180519067ffffffffffffffff821161096057680100000000000000008211610960576020906002840154836002860155808410611047575b500160028301608051526020608051206080515b83811061101d575050505060036060919e9c9d9e019101519081519167ffffffffffffffff831161096057680100000000000000008311610960576020908254848455808510610ffe575b500190608051526020608051206080515b838110610fd45750505050610fb9608095610fc995610fab7fece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d9a9660019e9d9c9a96610f9d67ffffffffffffffff976040519d8d8f9e8f9081520191614a80565b918b830360208d0152614a80565b9188830360408a0152614a80565b9285840360608701521696614a80565b0390a2019190610cc0565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501610f3c565b6110179084608051528584608051209182019101614a69565b38610f2b565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501610ee0565b6110639060028601608051528484608051209182019101614a69565b38610ecc565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501610e84565b6110af9060018601608051528484608051209182019101614a69565b38610e70565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501610e28565b6110f89085608051528484608051209182019101614a69565b38610e17565b346102d75760e06003193601126102d757611117613f11565b60607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc3601126112a55760405161114d81613d97565b60243580151581036112a55781526044356fffffffffffffffffffffffffffffffff811681036112a55760208201526064356fffffffffffffffffffffffffffffffff811681036112a557604082015260607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c3601126112a557604051906111d482613d97565b60843580151581036112a557825260a4356fffffffffffffffffffffffffffffffff811681036112a557602083015260c4356fffffffffffffffffffffffffffffffff811681036112a557604083015273ffffffffffffffffffffffffffffffffffffffff6009541633141580611283575b61125357610a2a92615467565b7f8e4a23d60000000000000000000000000000000000000000000000000000000060805152336004526024608051fd5b5073ffffffffffffffffffffffffffffffffffffffff60015416331415611246565b600080fd5b613f56565b346102d75760206003193601126102d75767ffffffffffffffff6112d1613f11565b6112d961490f565b50166080515260076020526113566112fd6112f860406080512061493a565b6155a4565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b0390f35b346102d7576080516003193601126102d7576040516005548082528160208101600560005260206000209260005b81811061147457505061139d92500382613db3565b8051906113c26113ac83614400565b926113ba6040519485613db3565b808452614400565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06020840192013683376080515b8151811015611423578067ffffffffffffffff61141060019385614467565b511661141c8287614467565b52016113f1565b505090604051918291602083019060208452518091526040830191906080515b818110611451575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101611443565b8454835260019485019486945060209093019201611388565b346102d75760206003193601126102d7576113566114b16114ac613f11565b6149a0565b604051918291602083526020830190613e2e565b346102d75760606003193601126102d75760043567ffffffffffffffff81116102d75760a060031982360301126102d7576114fe613ebe565b9060443567ffffffffffffffff81116102d75761151f903690600401613fde565b5061152861469e565b50608481019061153782614303565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611bbb57602481019077ffffffffffffffff0000000000000000000000000000000061159d8361425e565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611ac45760805191611b8c575b50611b605761162d60448201614303565b7f0000000000000000000000000000000000000000000000000000000000000000611b00575b5067ffffffffffffffff6116668361425e565b1661167e816000526006602052604060002054151590565b15611ad157602073ffffffffffffffffffffffffffffffffffffffff60045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015611ac45760805190611a61575b73ffffffffffffffffffffffffffffffffffffffff9150163303611a3157606461ffff91013593169283151592838094611a22575b156119675761ffff600a54169485811061193357506118f7945061176961175961173f8561425e565b67ffffffffffffffff16600052600b602052604060002090565b8361176384614303565b91615a6d565b7f7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b267ffffffffffffffff6117a561179f8661425e565b93614303565b6040805173ffffffffffffffffffffffffffffffffffffffff929092168252602082018690529190931692a25b918290611901575b506114ac611893917f000000000000000000000000000000000000000000000000000000000000000061182e857f000000000000000000000000000000000000000000000000000000000000000083614b1a565b7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10606067ffffffffffffffff6118638561425e565b73ffffffffffffffffffffffffffffffffffffffff604051951685523360208601528860408601521692a261425e565b9060405160ff7f0000000000000000000000000000000000000000000000000000000000000000166020820152602081526118cf604082613db3565b604051926118dc84613d5f565b83526020830152604051928392604084526040840190614184565b9060208301520390f35b61189391925061192b6114ac9161271061192461ffff600a5460101c1683614a56565b0490615597565b9291506117da565b85907fe08f03ef00000000000000000000000000000000000000000000000000000000608051526004526024526044608051fd5b506118f7935067ffffffffffffffff61197f8361425e565b16806080515260076020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894482806119f560406080512073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615a6d565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a26117d2565b5061ffff600a54161515611716565b7f728fe07b0000000000000000000000000000000000000000000000000000000060805152336004526024608051fd5b506020813d602011611abc575b81611a7b60209383613db3565b810103126102d7575173ffffffffffffffffffffffffffffffffffffffff811681036102d75773ffffffffffffffffffffffffffffffffffffffff906116e1565b3d9150611a6e565b6040513d608051823e3d90fd5b7fa9902c7e00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b73ffffffffffffffffffffffffffffffffffffffff16611b2d816000526003602052604060002054151590565b611653577fd0d2597600000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b7f53ad11d800000000000000000000000000000000000000000000000000000000608051526004608051fd5b611bae915060203d602011611bb4575b611ba68183613db3565b8101906151e5565b8561161c565b503d611b9c565b73ffffffffffffffffffffffffffffffffffffffff611bd983614303565b7f961c9a4f0000000000000000000000000000000000000000000000000000000060805152166004526024608051fd5b346102d75760406003193601126102d757611c22613cb4565b6024356bffffffffffffffffffffffff811681036102d757611c42614acf565b73ffffffffffffffffffffffffffffffffffffffff82169182156109cc577f39b29a1c46dd0e3ee5683fed1d66e11a234938cde48e5c4e14389be508348eb692600454907fffffffffffffffffffffffff00000000000000000000000000000000000000008460a01b1691608051161717600455611cf6604051928392839092916bffffffffffffffffffffffff60209173ffffffffffffffffffffffffffffffffffffffff604085019616845216910152565b0390a160805180f35b346102d75760206003193601126102d75767ffffffffffffffff611d21613f11565b611d2961490f565b50166080515260076020526113566112fd6112f860026040608051200161493a565b346102d75767ffffffffffffffff611d6236614112565b929091611d6d614acf565b1690611d86826000526006602052604060002054151590565b15611e485781608051526007602052611db9600560406080512001611dac368685613fa7565b6020815191012090615fd8565b15611e01577f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d769192611df86040519283926020845260208401916148d0565b0390a260805180f35b611e44906040519384937f74f23c7c00000000000000000000000000000000000000000000000000000000855260048501526040602485015260448401916148d0565b0390fd5b507f1e670e4b00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b346102d7576080516003193601126102d7576040516002548082526020820190600260005260206000209060005b818110611ec95761135685611ebd81870382613db3565b60405191829182614074565b8254845260209093019260019283019201611ea6565b346102d75760206003193601126102d75767ffffffffffffffff611f01613f11565b16608051526007602052611f1c600560406080512001615cb4565b8051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611f62611f4c84614400565b93611f5a6040519586613db3565b808552614400565b016080515b8181106120435750506080515b8151811015611fbe5780611f8a60019284614467565b51608051526008602052611fa2604060805120614810565b611fac8286614467565b52611fb78185614467565b5001611f74565b826040518091602082016020835281518091526040830190602060408260051b860101930191608051905b828210611ff857505050500390f35b91936020612033827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851613e2e565b9601920192018594939192611fe9565b806060602080938701015201611f67565b346102d757611356611ebd61206836613ffc565b50505091506146b7565b346102d75760206003193601126102d75760043567ffffffffffffffff81116102d75760a060031982360301126102d7576120ab61469e565b50604051602091906120bd8382613db3565b60805190526120ca61469e565b50608481016120d881614303565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691160361250d5750602481019077ffffffffffffffff0000000000000000000000000000000061213f8361425e565b60801b16604051907f2cbc26bb0000000000000000000000000000000000000000000000000000000082526004820152838160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611ac457608051916124f0575b50611b60576121ce60448201614303565b7f0000000000000000000000000000000000000000000000000000000000000000612490575b5067ffffffffffffffff6122078361425e565b1661221f816000526006602052604060002054151590565b15611ad1578373ffffffffffffffffffffffffffffffffffffffff60045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015611ac45760805190612432575b73ffffffffffffffffffffffffffffffffffffffff9150163303611a31576123d1916114ac91606401357ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae106060826123a78167ffffffffffffffff6122e58861425e565b168060005260078b526040600020906123387f00000000000000000000000000000000000000000000000000000000000000009273ffffffffffffffffffffffffffffffffffffffff8416968791615a6d565b6040805173ffffffffffffffffffffffffffffffffffffffff87168152602081018590527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789449190a27f000000000000000000000000000000000000000000000000000000000000000090614b1a565b67ffffffffffffffff6123b98661425e565b169360405191825233898301526040820152a261425e565b9060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168282015281815261240b604082613db3565b6040519261241884613d5f565b835281830152611356604051928284938452830190614184565b508381813d8311612489575b6124488183613db3565b810103126102d7575173ffffffffffffffffffffffffffffffffffffffff811681036102d75773ffffffffffffffffffffffffffffffffffffffff90612281565b503d61243e565b73ffffffffffffffffffffffffffffffffffffffff166124bd816000526003602052604060002054151590565b6121f4577fd0d2597600000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b6125079150843d8611611bb457611ba68183613db3565b846121bd565b611bd973ffffffffffffffffffffffffffffffffffffffff91614303565b346102d75760606003193601126102d75760043567ffffffffffffffff81116102d75761255c903690600401613e8d565b9060243567ffffffffffffffff81116102d75761257d903690600401614153565b9060443567ffffffffffffffff81116102d75761259e903690600401614153565b73ffffffffffffffffffffffffffffffffffffffff6009541633141580612666575b6112535783861480159061265c575b612630576080515b8681106125e45760805180f35b8061262a6125f8610a4b6001948b8b6142a2565b61260383898961468e565b61262461261c61261486898b61468e565b9236906141d8565b9136906141d8565b91615467565b016125d7565b7f568efce200000000000000000000000000000000000000000000000000000000608051526004608051fd5b50808614156125cf565b5073ffffffffffffffffffffffffffffffffffffffff600154163314156125c0565b346102d7576080516003193601126102d757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346102d75760206003193601126102d75760206126f867ffffffffffffffff6126e4613f11565b166000526006602052604060002054151590565b6040519015158152f35b346102d75760206003193601126102d75760043567ffffffffffffffff81116102d757612736610a2a913690600401613ee0565b9061273f614acf565b614cbe565b346102d75760206003193601126102d7577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174602073ffffffffffffffffffffffffffffffffffffffff612795613cb4565b61279d614acf565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000006009541617600955604051908152a160805180f35b346102d7576080516003193601126102d7576080515473ffffffffffffffffffffffffffffffffffffffff81163303612880577fffffffffffffffffffffffff0000000000000000000000000000000000000000600154913382841617600155166080515573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0608051608051a360805180f35b7f02b543c600000000000000000000000000000000000000000000000000000000608051526004608051fd5b346102d7576080516003193601126102d7576004546040805173ffffffffffffffffffffffffffffffffffffffff8316815260a09290921c602083015290f35b346102d7576080516003193601126102d757602073ffffffffffffffffffffffffffffffffffffffff60095416604051908152f35b346102d75761292f36614112565b61293a929192614acf565b67ffffffffffffffff821661295c816000526006602052604060002054151590565b156129775750610a2a92612971913691613fa7565b906151fd565b7f1e670e4b00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b346102d7576080516003193601126102d75760206129c26145d6565b604051908152f35b346102d7576129f26129fa6129de366140c4565b94916129eb939193614acf565b3691614582565b923691614582565b7f000000000000000000000000000000000000000000000000000000000000000015612b1c576080515b8251811015612a97578073ffffffffffffffffffffffffffffffffffffffff612a4f60019386614467565b5116612a5a81615d17565b612a66575b5001612a24565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a184612a5f565b506080515b8151811015610a2a578073ffffffffffffffffffffffffffffffffffffffff612ac760019385614467565b51168015612b1657612ad88161610c565b612ae5575b505b01612a9c565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a183612add565b50612adf565b7f35f4a7b300000000000000000000000000000000000000000000000000000000608051526004608051fd5b346102d757611356611ebd612b5c36613ffc565b505050915061447b565b346102d75760406003193601126102d757612b7f613f11565b60243567ffffffffffffffff81116102d757602091612ba56126f8923690600401613fde565b90614324565b346102d75760406003193601126102d75760043567ffffffffffffffff81116102d757806004019061010060031982360301126102d757612bea613ebe565b604051909190612bf981613d43565b6080519052612c2a612c20612c1b612c1460c48501876142b2565b3691613fa7565b614fe4565b60648301356150d8565b916084820190612c3982614303565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611bbb57602483019477ffffffffffffffff00000000000000000000000000000000612c9f8761425e565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611ac457608051916130f7575b50611b605767ffffffffffffffff612d358761425e565b16612d4d816000526006602052604060002054151590565b15611ad157602073ffffffffffffffffffffffffffffffffffffffff60045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115611ac457608051916130d8575b5015611a3157612dc68661425e565b90612ddc60a4860192612ba5612c1485856142b2565b156130915750604493929161ffff16159050612fd85767ffffffffffffffff612e048661425e565b1660805152600c602052612e216040608051208561176384614303565b7f41a8aa8df7945f0fb8ac5f7d88279638d9dc2ef9a6bf4ec9a53b80681b34aff767ffffffffffffffff612e5761179f8861425e565b6040805173ffffffffffffffffffffffffffffffffffffffff929092168252602082018890529190931692a25b0191612e8f83614303565b9073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691823b156102d7576040517f40c10f1900000000000000000000000000000000000000000000000000000000815260805173ffffffffffffffffffffffffffffffffffffffff909216600482015260248101859052908180604481010381608051875af18015611ac457612fbf575b50608067ffffffffffffffff60209573ffffffffffffffffffffffffffffffffffffffff612f8d612f877ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09661425e565b92614303565b60405196875233898801521660408601528560608601521692a260405190612fb482613d43565b815260405190518152f35b608051612fcb91613db3565b6080516102d75784612f36565b5067ffffffffffffffff612feb8561425e565b16806080515260076020527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c848061306460026040608051200173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615a6d565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a2612e84565b61309b92506142b2565b611e446040519283927f24eb47e50000000000000000000000000000000000000000000000000000000084526020600485015260248401916148d0565b6130f1915060203d602011611bb457611ba68183613db3565b87612db7565b613110915060203d602011611bb457611ba68183613db3565b87612d1e565b346102d75760206003193601126102d75760043567ffffffffffffffff81116102d757806004019061010060031982360301126102d75760405161315981613d43565b608051905260405161316a81613d43565b6080519052613185612c20612c1b612c1460c48501866142b2565b906084810161319381614303565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811691160361250d5750602481019277ffffffffffffffff000000000000000000000000000000006131fa8561425e565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611ac457608051916134ef575b50611b605767ffffffffffffffff6132908561425e565b166132a8816000526006602052604060002054151590565b15611ad157602073ffffffffffffffffffffffffffffffffffffffff60045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115611ac457608051916134d0575b5015611a31576133218461425e565b9061333760a4840192612ba5612c1485856142b2565b15613091575082916044915067ffffffffffffffff6133558661425e565b16806080515260076020526133ab60026040608051200173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016958691615a6d565b6040805173ffffffffffffffffffffffffffffffffffffffff86168152602081018790527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a201926133fe84614303565b823b156102d7576040517f40c10f1900000000000000000000000000000000000000000000000000000000815260805173ffffffffffffffffffffffffffffffffffffffff909216600482015260248101859052908180604481010381608051875af18015611ac45760209573ffffffffffffffffffffffffffffffffffffffff612f8d612f877ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09660809667ffffffffffffffff966134bf575b5061425e565b87516134ca91613db3565b8b6134b9565b6134e9915060203d602011611bb457611ba68183613db3565b85613312565b613508915060203d602011611bb457611ba68183613db3565b85613279565b346102d75760a06003193601126102d757613527613cb4565b5060243567ffffffffffffffff8116908190036102d75760443567ffffffffffffffff81116102d75760031960a091360301126102d757613566613ecf565b5060843567ffffffffffffffff81116102d757613587903690600401613f28565b505060405161359581613cf8565b608051815260805160208201526080516040820152606060805191015260805152600e60205260806040815120604051906135cf82613cf8565b5463ffffffff808216928381528160208201818560201c16815260ff60606040850194848860401c168652019560601c161515855260405195865251166020850152511660408301525115156060820152f35b346102d75760606003193601126102d75760043561ffff8116908190036102d75761364b613ebe565b9060443567ffffffffffffffff81116102d75761366c903690600401613ee0565b90613675614acf565b61ffff8416936127108510156136fc5783927f52aa194b292c8bfb5aaca8ee2000a965c3a051b306ff841873b16147526a39ba95926136eb926040967fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000063ffff0000600a549360101b1692161717600a55614cbe565b82519182526020820152a160805180f35b847f95f3517a00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b346102d75760406003193601126102d75760043567ffffffffffffffff81116102d757366023820112156102d757806004013567ffffffffffffffff81116102d75760248201916024369160a084020101116102d75760243567ffffffffffffffff81116102d7576137a2903690600401613e8d565b9190926137ad614acf565b6080515b828110613829575050506080515b8181106137cc5760805180f35b8067ffffffffffffffff6137e6610a4b60019486886142a2565b168060805152600e602052608051604060805120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8608051608051a2016137bf565b8061383a610a4b600193868661421f565b7f56f77aeff2def50c8b5f5a0df3bab7183df09bf36c6feba496bb42551db77d70608061386884888861421f565b9261397861393e63ffffffff61396d6139318261396267ffffffffffffffff60208c0198169a8b8a5152600e60205260408a5120836138a68b614273565b169181549060408101937fffffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffff67ffffffff000000006138e387614273565b60201b16918f6cff0000000000000000000000007fffffffffffffffffffffffffffffffffffffffff000000000000000000000000916bffffffff0000000000000000606088019d8e614273565b60401b1696019e8f614284565b151560601b169516171617171790558261395a6040519a614291565b168952614291565b166020870152614291565b1660408401526141ae565b15156060820152a2016137b1565b346102d7576080516003193601126102d757602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102d75760206003193601126102d75760206139e0613cb4565b6040517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff9081169216919091148152f35b346102d7576080516003193601126102d757602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102d7576080516003193601126102d7576040805161135691613a9c9082613db3565b602081527f4275726e546f41646472657373546f6b656e506f6f6c20312e362e332d6465766020820152604051918291602083526020830190613e2e565b346102d75760206003193601126102d757613af3613cb4565b613afb614acf565b613b036145d6565b9081613b0f5760805180f35b602073ffffffffffffffffffffffffffffffffffffffff82613b73857f738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b859599957f0000000000000000000000000000000000000000000000000000000000000000614b1a565b6040519485521692a28080610a2a565b346102d75760206003193601126102d757600435907fffffffff0000000000000000000000000000000000000000000000000000000082168092036102d757817ff208a58f0000000000000000000000000000000000000000000000000000000060209314908115613c8a575b8115613c60575b8115613c36575b8115613c0c575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483613c05565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150613bfe565b7f1ef5498f0000000000000000000000000000000000000000000000000000000081149150613bf7565b7faff2afbf0000000000000000000000000000000000000000000000000000000081149150613bf0565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036112a557565b359073ffffffffffffffffffffffffffffffffffffffff821682036112a557565b6080810190811067ffffffffffffffff821117613d1457604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff821117613d1457604052565b6040810190811067ffffffffffffffff821117613d1457604052565b60a0810190811067ffffffffffffffff821117613d1457604052565b6060810190811067ffffffffffffffff821117613d1457604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117613d1457604052565b67ffffffffffffffff8111613d1457601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b848110613e785750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201613e39565b9181601f840112156112a55782359167ffffffffffffffff83116112a5576020808501948460051b0101116112a557565b6024359061ffff821682036112a557565b6064359061ffff821682036112a557565b9181601f840112156112a55782359167ffffffffffffffff83116112a55760208085019460e085020101116112a557565b6004359067ffffffffffffffff821682036112a557565b9181601f840112156112a55782359167ffffffffffffffff83116112a557602083818601950101116112a557565b346112a55760006003193601126112a557602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b929192613fb382613df4565b91613fc16040519384613db3565b8294818452818301116112a5578281602093846000960137010152565b9080601f830112156112a557816020613ff993359101613fa7565b90565b60a06003198201126112a55760043573ffffffffffffffffffffffffffffffffffffffff811681036112a5579160243567ffffffffffffffff811681036112a557916044359160643561ffff811681036112a557916084359067ffffffffffffffff82116112a55761407091600401613f28565b9091565b602060408183019282815284518094520192019060005b8181106140985750505090565b825173ffffffffffffffffffffffffffffffffffffffff1684526020938401939092019160010161408b565b60406003198201126112a55760043567ffffffffffffffff81116112a557816140ef91600401613e8d565b929092916024359067ffffffffffffffff82116112a55761407091600401613e8d565b9060406003198301126112a55760043567ffffffffffffffff811681036112a557916024359067ffffffffffffffff82116112a55761407091600401613f28565b9181601f840112156112a55782359167ffffffffffffffff83116112a557602080850194606085020101116112a557565b613ff991602061419d8351604084526040840190613e2e565b920151906020818403910152613e2e565b359081151582036112a557565b35906fffffffffffffffffffffffffffffffff821682036112a557565b91908260609103126112a5576040516141f081613d97565b604061421a818395614201816141ae565b855261420f602082016141bb565b6020860152016141bb565b910152565b919081101561422f5760a0020190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3567ffffffffffffffff811681036112a55790565b3563ffffffff811681036112a55790565b3580151581036112a55790565b359063ffffffff821682036112a557565b919081101561422f5760051b0190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156112a5570180359067ffffffffffffffff82116112a5576020019181360383136112a557565b3573ffffffffffffffffffffffffffffffffffffffff811681036112a55790565b9067ffffffffffffffff613ff992166000526007602052600560406000200190602081519101209060019160005201602052604060002054151590565b906040519182815491828252602082019060005260206000209260005b81811061439557505061439392500383613db3565b565b845473ffffffffffffffffffffffffffffffffffffffff1683526001948501948794506020909301920161437e565b919082018092116143d157565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b67ffffffffffffffff8111613d145760051b60200190565b9061442282614400565b61442f6040519182613db3565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061445d8294614400565b0190602036910137565b805182101561422f5760209160051b010190565b67ffffffffffffffff16600052600d602052604060002061449b81614361565b9160045460a01c8015159182614577575b50506144b6575090565b60016144c29101614361565b908151806144d05750905090565b6144de6144e39183516143c4565b614418565b9260005b8251811015614525578073ffffffffffffffffffffffffffffffffffffffff61451260019386614467565b511661451e8288614467565b52016144e7565b509160005b8151811015614572578073ffffffffffffffffffffffffffffffffffffffff61455560019385614467565b511661456b6145658387516143c4565b88614467565b520161452a565b505050565b1015905038806144ac565b92919061458e81614400565b9361459c6040519586613db3565b602085838152019160051b81019283116112a557905b8282106145be57505050565b602080916145cb84613cd7565b8152019101906145b2565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561468257600091614653575090565b90506020813d60201161467a575b8161466e60209383613db3565b810103126112a5575190565b3d9150614661565b6040513d6000823e3d90fd5b919081101561422f576060020190565b604051906146ab82613d5f565b60606020838281520152565b67ffffffffffffffff16600052600d60205260406000209060036146dd60028401614361565b9201906146e982614361565b5060045460a01c80151591826147b2575b5050614704575090565b61470d90614361565b9081518061471b5750905090565b6144de6147299183516143c4565b9260005b825181101561476b578073ffffffffffffffffffffffffffffffffffffffff61475860019386614467565b51166147648288614467565b520161472d565b509160005b8151811015614572578073ffffffffffffffffffffffffffffffffffffffff61479b60019385614467565b51166147ab6145658387516143c4565b5201614770565b1015905038806146fa565b90600182811c92168015614806575b60208310146147d757565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f16916147cc565b9060405191826000825492614824846147bd565b80845293600181169081156148905750600114614849575b5061439392500383613db3565b90506000929192526020600020906000915b818310614874575050906020614393928201013861483c565b602091935080600191548385890101520191019091849261485b565b602093506143939592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b8201013861483c565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b6040519061491c82613d7b565b60006080838281528260208201528260408201528260608201520152565b9060405161494781613d7b565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff166000526007602052613ff96004604060002001614810565b919081101561422f5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61813603018212156112a5570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156112a5570180359067ffffffffffffffff82116112a557602001918160051b360383136112a557565b818102929181159184041417156143d157565b818110614a74575050565b60008155600101614a69565b9160209082815201919060005b818110614a9a5750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff614ac188613cd7565b168152019401929101614a8d565b73ffffffffffffffffffffffffffffffffffffffff600154163303614af057565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b6040517fa9059cbb000000000000000000000000000000000000000000000000000000006020820190815273ffffffffffffffffffffffffffffffffffffffff9384166024830152604480830195909552938152614bf0929091614b7f606484613db3565b16600080604095865194614b938887613db3565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020870152519082855af13d15614c99573d91614bd483613df4565b92614be187519485613db3565b83523d6000602085013e61621b565b80519081614bfd57505050565b602080614c0e9383010191016151e5565b15614c165750565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b60609161621b565b356fffffffffffffffffffffffffffffffff811681036112a55790565b9160005b82811015614fde5760e0810284016000614cdb8261425e565b9067ffffffffffffffff821691614cff836000526006602052604060002054151590565b15614fb257614dc89260408593614d73614d6d94614d6d614d33602060019c9b019261173f614d2e36866141d8565b6156f3565b91825463ffffffff8160801c16159081614f94575b81614f85575b81614f6a575b81614f5b575b5080614f4c575b614ec1575b36906141d8565b9061583a565b6080850192614d85614d2e36866141d8565b8152600c6020522092835463ffffffff8160801c16159081614ea3575b81614e94575b81614e79575b81614e6a575b5080614e5b575b614dce575b5036906141d8565b01614cc2565b614deb60a06fffffffffffffffffffffffffffffffff9201614ca1565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff161717835538614dc0565b50614e6582614284565b614dbb565b60ff915060a01c161538614db4565b6fffffffffffffffffffffffffffffffff8116159150614dae565b8589015460801c159150614da8565b858901546fffffffffffffffffffffffffffffffff16159150614da2565b6fffffffffffffffffffffffffffffffff614edd878b01614ca1565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff1617178355614d66565b50614f5681614284565b614d61565b60ff915060a01c161538614d5a565b6fffffffffffffffffffffffffffffffff8116159150614d54565b848e015460801c159150614d4e565b848e01546fffffffffffffffffffffffffffffffff16159150614d48565b506024917f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50915050565b80518015615054576020036150165780516020828101918301839003126112a557519060ff8211615016575060ff1690565b611e44906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190613e2e565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff82116143d157565b60ff16604d81116143d157600a0a90565b81156150a9570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff8116928284146151de578284116151b4579061511d9161507a565b91604d60ff841611801561517b575b6151455750509061513f613ff99261508e565b90614a56565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b506151858361508e565b80156150a9577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04841161512c565b6151bd9161507a565b91604d60ff841611615145575050906151d8613ff99261508e565b9061509f565b5050505090565b908160209103126112a5575180151581036112a55790565b9080511561543d5767ffffffffffffffff815160208301209216918260005260076020526152328160056040600020016161c6565b156153f95760005260086020526040600020815167ffffffffffffffff8111613d145761525f82546147bd565b601f81116153c7575b506020601f821160011461530157916152db827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea95936152f1956000916152f6575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055604051918291602083526020830190613e2e565b0390a2565b9050840151386152aa565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b8181106153af5750926152f19492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610615378575b5050811b0190556114b1565b8501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388061536c565b9192602060018192868a015181550194019201615331565b6153f390836000526020600020601f840160051c8101916020851061095657601f0160051c0190614a69565b38615268565b5090611e446040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190613e2e565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b67ffffffffffffffff166000818152600660205260409020549092919015615569579161556660e092615532856154be7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b976156f3565b8460005260076020526154d581604060002061583a565b6154de836156f3565b8460005260076020526154f883600260406000200161583a565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b919082039182116143d157565b6155ac61490f565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff808351169161560960208501936156036155f663ffffffff87511642615597565b8560808901511690614a56565b906143c4565b8082101561562257505b16825263ffffffff4216905290565b9050615613565b805160005b81811061563a57505050565b600181018082116143d1575b828110615656575060010161562e565b73ffffffffffffffffffffffffffffffffffffffff6156758386614467565b511673ffffffffffffffffffffffffffffffffffffffff6156968387614467565b5116146156a557600101615646565b73ffffffffffffffffffffffffffffffffffffffff6156c48386614467565b51167fa1726e400000000000000000000000000000000000000000000000000000000060005260045260246000fd5b805115615793576fffffffffffffffffffffffffffffffff6040820151166fffffffffffffffffffffffffffffffff602083015116106157305750565b606490615791604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff6040820151161580159061581b575b6157ba5750565b606490615791604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff60208201511615156157b3565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1991615973606092805461587763ffffffff8260801c1642615597565b90816159b2575b50506fffffffffffffffffffffffffffffffff60018160208601511692828154168085106000146159aa57508280855b16167fffffffffffffffffffffffffffffffff000000000000000000000000000000008254161781556159278651151582907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b60408601517fffffffffffffffffffffffffffffffff0000000000000000000000000000000060809190911b16939092166fffffffffffffffffffffffffffffffff1692909217910155565b61556660405180926fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b8380916158ae565b6fffffffffffffffffffffffffffffffff916159e78392836159e06001880154948286169560801c90614a56565b91166143c4565b80821015615a6657505b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff9290911692909216167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116174260801b73ffffffff0000000000000000000000000000000016178155388061587e565b90506159f1565b9182549060ff8260a01c16158015615cac575b615ca6576fffffffffffffffffffffffffffffffff82169160018501908154615ac563ffffffff6fffffffffffffffffffffffffffffffff83169360801c1642615597565b9081615c08575b5050848110615bbc5750838310615b26575050615afb6fffffffffffffffffffffffffffffffff928392615597565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b5460801c91615b358185615597565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101908082116143d157615b83615b889273ffffffffffffffffffffffffffffffffffffffff966143c4565b61509f565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828573ffffffffffffffffffffffffffffffffffffffff927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611615c7c57615c23926156039160801c90614a56565b80841015615c775750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178655923880615acc565b615c2e565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b508215615a80565b906040519182815491828252602082019060005260206000209260005b818110615ce657505061439392500383613db3565b8454835260019485019487945060209093019201615cd1565b805482101561422f5760005260206000200190600090565b6000818152600360205260409020548015615ea6577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116143d157600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116143d157818103615e37575b5050506002548015615e08577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01615dc5816002615cff565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b615e8e615e48615e59936002615cff565b90549060031b1c9283926002615cff565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526003602052604060002055388080615d8c565b5050600090565b6000818152600660205260409020548015615ea6577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116143d157600554907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116143d157818103615f9e575b5050506005548015615e08577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01615f5b816005615cff565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600555600052600660205260006040812055600190565b615fc0615faf615e59936005615cff565b90549060031b1c9283926005615cff565b90556000526006602052604060002055388080615f22565b9060018201918160005282602052604060002054801515600014616103577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116143d1578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116143d1578181036160cc575b50505080548015615e08577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019061608d8282615cff565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b6160ec6160dc615e599386615cff565b90549060031b1c92839286615cff565b905560005283602052604060002055388080616055565b50505050600090565b806000526003602052604060002054156000146161665760025468010000000000000000811015613d145761614d615e598260018594016002556002615cff565b9055600254906000526003602052604060002055600190565b50600090565b806000526006602052604060002054156000146161665760055468010000000000000000811015613d14576161ad615e598260018594016005556005615cff565b9055600554906000526006602052604060002055600190565b6000828152600182016020526040902054615ea65780549068010000000000000000821015613d145782616204615e59846001809601855584615cff565b905580549260005201602052604060002055600190565b91929015616296575081511561622f575090565b3b156162385790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156162a95750805190602001fd5b611e44906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190613e2e56fea164736f6c634300081a000a",
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) GetRequiredInboundCCVs(opts *bind.CallOpts, arg0 common.Address, sourceChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "getRequiredInboundCCVs", arg0, sourceChainSelector, amount, arg3, arg4)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) GetRequiredInboundCCVs(arg0 common.Address, sourceChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _BurnToAddressMintTokenPool.Contract.GetRequiredInboundCCVs(&_BurnToAddressMintTokenPool.CallOpts, arg0, sourceChainSelector, amount, arg3, arg4)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) GetRequiredInboundCCVs(arg0 common.Address, sourceChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _BurnToAddressMintTokenPool.Contract.GetRequiredInboundCCVs(&_BurnToAddressMintTokenPool.CallOpts, arg0, sourceChainSelector, amount, arg3, arg4)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) GetRequiredOutboundCCVs(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "getRequiredOutboundCCVs", arg0, destChainSelector, amount, arg3, arg4)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) GetRequiredOutboundCCVs(arg0 common.Address, destChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _BurnToAddressMintTokenPool.Contract.GetRequiredOutboundCCVs(&_BurnToAddressMintTokenPool.CallOpts, arg0, destChainSelector, amount, arg3, arg4)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) GetRequiredOutboundCCVs(arg0 common.Address, destChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _BurnToAddressMintTokenPool.Contract.GetRequiredOutboundCCVs(&_BurnToAddressMintTokenPool.CallOpts, arg0, destChainSelector, amount, arg3, arg4)
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactor) ApplyFinalityConfigUpdates(opts *bind.TransactOpts, finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.contract.Transact(opts, "applyFinalityConfigUpdates", finalityThreshold, customFinalityTransferFeeBps, rateLimitConfigArgs)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) ApplyFinalityConfigUpdates(finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.ApplyFinalityConfigUpdates(&_BurnToAddressMintTokenPool.TransactOpts, finalityThreshold, customFinalityTransferFeeBps, rateLimitConfigArgs)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorSession) ApplyFinalityConfigUpdates(finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.ApplyFinalityConfigUpdates(&_BurnToAddressMintTokenPool.TransactOpts, finalityThreshold, customFinalityTransferFeeBps, rateLimitConfigArgs)
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactor) SetCustomFinalityRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.contract.Transact(opts, "setCustomFinalityRateLimitConfig", rateLimitConfigArgs)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) SetCustomFinalityRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.SetCustomFinalityRateLimitConfig(&_BurnToAddressMintTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorSession) SetCustomFinalityRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.SetCustomFinalityRateLimitConfig(&_BurnToAddressMintTokenPool.TransactOpts, rateLimitConfigArgs)
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
	RemoteChainSelector    uint64
	OutboundCCVs           []common.Address
	AdditionalOutboundCCVs []common.Address
	InboundCCVs            []common.Address
	AdditionalInboundCCVs  []common.Address
	Raw                    types.Log
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

type BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumedIterator struct {
	Event *BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterCustomFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "CustomFinalityOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumedIterator{contract: _BurnToAddressMintTokenPool.contract, event: "CustomFinalityOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchCustomFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "CustomFinalityOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumed)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "CustomFinalityOutboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseCustomFinalityOutboundRateLimitConsumed(log types.Log) (*BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumed, error) {
	event := new(BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumed)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "CustomFinalityOutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator struct {
	Event *BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterCustomFinalityTransferInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "CustomFinalityTransferInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator{contract: _BurnToAddressMintTokenPool.contract, event: "CustomFinalityTransferInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchCustomFinalityTransferInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "CustomFinalityTransferInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "CustomFinalityTransferInboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseCustomFinalityTransferInboundRateLimitConsumed(log types.Log) (*BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed, error) {
	event := new(BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "CustomFinalityTransferInboundRateLimitConsumed", log); err != nil {
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

type BurnToAddressMintTokenPoolFinalityConfigUpdatedIterator struct {
	Event *BurnToAddressMintTokenPoolFinalityConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolFinalityConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolFinalityConfigUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnToAddressMintTokenPoolFinalityConfigUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnToAddressMintTokenPoolFinalityConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolFinalityConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolFinalityConfigUpdated struct {
	FinalityConfig               uint16
	CustomFinalityTransferFeeBps uint16
	Raw                          types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterFinalityConfigUpdated(opts *bind.FilterOpts) (*BurnToAddressMintTokenPoolFinalityConfigUpdatedIterator, error) {

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "FinalityConfigUpdated")
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolFinalityConfigUpdatedIterator{contract: _BurnToAddressMintTokenPool.contract, event: "FinalityConfigUpdated", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchFinalityConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolFinalityConfigUpdated) (event.Subscription, error) {

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "FinalityConfigUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolFinalityConfigUpdated)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "FinalityConfigUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseFinalityConfigUpdated(log types.Log) (*BurnToAddressMintTokenPoolFinalityConfigUpdated, error) {
	event := new(BurnToAddressMintTokenPoolFinalityConfigUpdated)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "FinalityConfigUpdated", log); err != nil {
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

func (BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b2")
}

func (BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x41a8aa8df7945f0fb8ac5f7d88279638d9dc2ef9a6bf4ec9a53b80681b34aff7")
}

func (BurnToAddressMintTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x39b29a1c46dd0e3ee5683fed1d66e11a234938cde48e5c4e14389be508348eb6")
}

func (BurnToAddressMintTokenPoolFinalityConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x52aa194b292c8bfb5aaca8ee2000a965c3a051b306ff841873b16147526a39ba")
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
	return common.HexToHash("0x56f77aeff2def50c8b5f5a0df3bab7183df09bf36c6feba496bb42551db77d70")
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

	FilterCustomFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumedIterator, error)

	WatchCustomFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomFinalityOutboundRateLimitConsumed(log types.Log) (*BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumed, error)

	FilterCustomFinalityTransferInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator, error)

	WatchCustomFinalityTransferInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomFinalityTransferInboundRateLimitConsumed(log types.Log) (*BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*BurnToAddressMintTokenPoolDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*BurnToAddressMintTokenPoolDynamicConfigSet, error)

	FilterFinalityConfigUpdated(opts *bind.FilterOpts) (*BurnToAddressMintTokenPoolFinalityConfigUpdatedIterator, error)

	WatchFinalityConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolFinalityConfigUpdated) (event.Subscription, error)

	ParseFinalityConfigUpdated(log types.Log) (*BurnToAddressMintTokenPoolFinalityConfigUpdated, error)

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
