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

var BurnToAddressMintTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIBurnMintERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"burnAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyFinalityConfigUpdates\",\"inputs\":[{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"destToUseDefaultFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAccumulatedFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBurnAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"burnAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredInboundCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredOutboundCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_burnAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"lockOrBurnOutV1\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustomFinalityRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRouter\",\"inputs\":[{\"name\":\"newRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityTransferInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalityConfigUpdated\",\"inputs\":[{\"name\":\"finalityConfig\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeeWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RouterUpdated\",\"inputs\":[{\"name\":\"oldRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidFinality\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidFinalityConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenTransferFeeConfigNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x61012080604052346103dd5761662b803803809161001d828561045c565b8339810160c0828203126103dd5781516001600160a01b03811692908390036103dd5761004c6020820161047f565b60408201516001600160401b0381116103dd5782019280601f850112156103dd578351936001600160401b038511610446578460051b906020820195610095604051978861045c565b86526020808701928201019283116103dd57602001905b82821061042e575050506100c26060830161048d565b936100db60a06100d46080860161048d565b940161048d565b94331561041d57600180546001600160a01b031916331790558115801561040c575b80156103fb575b6103ea578160209160049360805260c0526040519283809263313ce56760e01b82525afa600091816103a9575b5061037e575b5060a052600480546001600160a01b0319166001600160a01b03929092169190911790558051151560e0819052610261575b5061010052604051615fe990816106428239608051818181611412015281816116a20152818161185701528181611e9f015281816121180152818161227101528181612b5e01528181612dac01528181612f3d015281816130d1015281816132a1015281816139110152818161396b01528181613a64015261438f015260a0518181816138b601528181614d0501528181614d88015261524b015260c051818181610c0a015281816114ad01528181611f3a01528181612bf9015261316d015260e051818181610bb8015281816114f001528181611f7d01526128f70152610100518181816116c70152818161213e0152613e990152f35b6020604051610270828261045c565b60008152600036813760e0511561036d5760005b81518110156102eb576001906001600160a01b036102a282856104a1565b5116846102ae826104e3565b6102bb575b505001610284565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a138846102b3565b505060005b8251811015610364576001906001600160a01b0361030e82866104a1565b5116801561035e5783610320826105e1565b61032e575b50505b016102f0565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a13883610325565b50610328565b50505038610169565b6335f4a7b360e01b60005260046000fd5b60ff1660ff82168181036103925750610137565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d6020116103e2575b816103c56020938361045c565b810103126103dd576103d69061047f565b9038610131565b600080fd5b3d91506103b8565b630a64406560e11b60005260046000fd5b506001600160a01b03811615610104565b506001600160a01b038416156100fd565b639b15e16f60e01b60005260046000fd5b6020809161043b8461048d565b8152019101906100ac565b634e487b7160e01b600052604160045260246000fd5b601f909101601f19168101906001600160401b0382119082101761044657604052565b519060ff821682036103dd57565b51906001600160a01b03821682036103dd57565b80518210156104b55760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b80548210156104b55760005260206000200190600090565b60008181526003602052604090205480156105da5760001981018181116105c4576002546000198101919082116105c457818103610573575b505050600254801561055d57600019016105378160026104cb565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b6105ac6105846105959360026104cb565b90549060031b1c92839260026104cb565b819391549060031b91821b91600019901b19161790565b9055600052600360205260406000205538808061051c565b634e487b7160e01b600052601160045260246000fd5b5050600090565b8060005260036020526040600020541560001461063b57600254680100000000000000008110156104465761062261059582600185940160025560026104cb565b9055600254906000526003602052604060002055600190565b5060009056fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a714613a9757508063164e68de146139f0578063181f5a771461398f57806321df0da71461393e578063240028e8146138da57806324f65ee71461389c5780632a10097b1461364f5780632c286daf1461354d57806337b192471461344657806338b39d2914610ddf5780633907753714613038578063489a68f214612abc5780634c5ef0ed14612a755780634f71592c14612a4057806354c8a4f3146128c35780635df45a37146128a057806362ddd3c41461281c5780636d3d1a58146127e857806379ba5097146127215780637d54534e14612692578063804ba5a91461264f5780638926f54f146126095780638da5cb5b146125d5578063962d40201461244f5780639a4575b914611e3e5780639f68f67314611e06578063a42a7b8b14611c9f578063a7cd63b714611c1e578063acfecf9114611afa578063af58d59f14611ab1578063b0f479a114611a7d578063b1c71c6514611387578063b79019b5146110ca578063b794658014611091578063c0d7865514610fb7578063c4bffe2b14610e8c578063c75eea9c14610de4578063c8de9fe014610ddf578063cf7401f314610c2e578063dc0bd97114610bdd578063e0351e1314610ba0578063e8a1da17146102cb5763f2fde38b146101fa57600080fd5b346102c85760206003193601126102c85773ffffffffffffffffffffffffffffffffffffffff610228613bca565b61023061477c565b163381146102a057807fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b80fd5b50346102c8576102da36613fda565b939190926102e661477c565b82915b808310610a0b575050508063ffffffff4216917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee1843603015b85821015610a07578160051b85013581811215610a035785019061012082360312610a03576040519561035487613c91565b823567ffffffffffffffff811681036109fe578752602083013567ffffffffffffffff81116109fa5783019536601f880112156109fa57863596610397886142da565b976103a5604051998a613cc9565b8089526020808a019160051b830101903682116109f65760208301905b8282106109c3575050505060208801968752604084013567ffffffffffffffff81116109bf576103f59036908601613ef4565b9860408901998a5261041f61040d36606088016140ee565b9560608b0196875260c03691016140ee565b9660808a0197885261043186516153e8565b61043b88516153e8565b8a5151156109975761045767ffffffffffffffff8b5116615e61565b156109605767ffffffffffffffff8a5116815260076020526040812061059787516fffffffffffffffffffffffffffffffff604082015116906105526fffffffffffffffffffffffffffffffff602083015116915115158360806040516104bd81613c91565b858152602081018c905260408101849052606081018690520152855474ff000000000000000000000000000000000000000091151560a01b919091167fffffffffffffffffffffff0000000000000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff84161773ffffffff0000000000000000000000000000000060808b901b1617178555565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176001830155565b6106bd89516fffffffffffffffffffffffffffffffff604082015116906106786fffffffffffffffffffffffffffffffff602083015116915115158360806040516105e181613c91565b858152602081018c9052604081018490526060810186905201526002860180547fffffffffffffffffffffff000000000000000000000000000000000000000000166fffffffffffffffffffffffffffffffff85161773ffffffff0000000000000000000000000000000060808c901b161791151560a01b74ff000000000000000000000000000000000000000016919091179055565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176003830155565b60048c5191019080519067ffffffffffffffff8211610933576106e0835461443b565b601f81116108f8575b50602090601f831160011461085957610737929185918361084e575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b805b89518051821015610772579061076c600192610765838f67ffffffffffffffff90511692614427565b5190614eaa565b0161073c565b5050975097987f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c29295939661084067ffffffffffffffff600197949c511692519351915161080c6107d760405196879687526101006020880152610100870190613d44565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a1019093949291610322565b015190503880610705565b83855281852091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416865b8181106108e057509084600195949392106108a9575b505050811b01905561073a565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c1916905538808061089c565b92936020600181928786015181550195019301610886565b6109239084865260208620601f850160051c81019160208610610929575b601f0160051c0190614743565b386106e9565b9091508190610916565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60249067ffffffffffffffff8b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b807f14c880ca0000000000000000000000000000000000000000000000000000000060049252fd5b8680fd5b813567ffffffffffffffff81116109f2576020916109e78392833691890101613ef4565b8152019101906103c2565b8a80fd5b8880fd5b8580fd5b600080fd5b8380fd5b8280f35b9092919367ffffffffffffffff610a2b610a268785886141b8565b614174565b1695610a3687615ba2565b15610b74578684526007602052610a52600560408620016159a9565b94845b8651811015610a8b576001908987526007602052610a8460056040892001610a7d838b614427565b5190615ccd565b5001610a55565b5093945094909580855260076020526005604086208681558660018201558660028201558660038201558660048201610ac4815461443b565b80610b33575b5050500180549086815581610b15575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a10191909493946102e9565b865260208620908101905b81811015610ada57868155600101610b20565b601f8111600114610b495750555b863880610aca565b81835260208320610b6491601f01861c810190600101614743565b8082528160208120915555610b41565b602484887f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50346102c857806003193601126102c85760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b50346102c857806003193601126102c857602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102c85760e06003193601126102c857610c48613e27565b9060607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc3601126102c857604051610c7f81613cad565b6024358015158103610ddb5781526044356fffffffffffffffffffffffffffffffff81168103610ddb5760208201526064356fffffffffffffffffffffffffffffffff81168103610ddb57604082015260607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c360112610dd75760405190610d0682613cad565b6084358015158103610a0357825260a4356fffffffffffffffffffffffffffffffff81168103610a0357602083015260c4356fffffffffffffffffffffffffffffffff81168103610a0357604083015273ffffffffffffffffffffffffffffffffffffffff6009541633141580610db5575b610d8957610d869293615114565b80f35b6024837f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b5073ffffffffffffffffffffffffffffffffffffffff60015416331415610d78565b5080fd5b8280fd5b613e6c565b50346102c85760206003193601126102c857610e2f610e2a6040610e889367ffffffffffffffff610e13613e27565b610e1b61458d565b501681526007602052206145b8565b61528c565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b0390f35b50346102c857806003193601126102c857604051906005548083528260208101600584526020842092845b818110610f9e575050610ecc92500383613cc9565b8151610ef0610eda826142da565b91610ee86040519384613cc9565b8083526142da565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602083019301368437805b8451811015610f4f578067ffffffffffffffff610f3c60019388614427565b5116610f488286614427565b5201610f1d565b50925090604051928392602084019060208552518091526040840192915b818110610f7b575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101610f6d565b8454835260019485019487945060209093019201610eb7565b50346102c85760206003193601126102c85773ffffffffffffffffffffffffffffffffffffffff610fe6613bca565b610fee61477c565b1680156110695760407f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f16849160045490807fffffffffffffffffffffffff000000000000000000000000000000000000000083161760045573ffffffffffffffffffffffffffffffffffffffff8351921682526020820152a180f35b6004827f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b50346102c85760206003193601126102c857610e886110b66110b1613e27565b61475a565b604051918291602083526020830190613d44565b50346102c85760206003193601126102c85760043567ffffffffffffffff8111610dd7576110fc903690600401613da3565b61110461477c565b825b818110611111578380f35b61111f610a2682848661461e565b61113761112d83858761461e565b602081019061465e565b907fb0897119e8510f887b892cbc4c8506fc51d9849fd90afae4fd065e705f2d0f6c61117161116786888a61461e565b604081019061465e565b9190926111876111823687846142f2565b61531e565b6111956111823685876142f2565b604051946111a286613c75565b6111ad3682846142f2565b86526111f767ffffffffffffffff6111c63687896142f2565b9860208901998a52169586956111e96040519586956040875260408701916146b2565b9184830360208601526146b2565b0390a28652600d60205260408620905180519067ffffffffffffffff821161135a5768010000000000000000821161135a576020908354838555808410611340575b500182885260208820885b838110611316575050505060010190519081519167ffffffffffffffff83116112e9576801000000000000000083116112e95760209082548484558085106112cf575b500190865260208620865b8381106112a55750505050600101611106565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501611292565b8389528289206112e3918101908601614743565b38611287565b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501611244565b848a52828a20611354918101908501614743565b38611239565b6024887f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b50346102c85760606003193601126102c85760043567ffffffffffffffff8111610dd75760a06003198236030112610dd7576113c1613dd4565b9060443567ffffffffffffffff8111610a03576113e2903690600401613ef4565b506113eb61440e565b5060848101916113fa83614219565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611a3357602482019177ffffffffffffffff0000000000000000000000000000000061146084614174565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561194a578691611a04575b506119dc576114ee60448201614219565b7f0000000000000000000000000000000000000000000000000000000000000000611980575b5067ffffffffffffffff61152784614174565b1661153f816000526006602052604060002054151590565b1561195557602073ffffffffffffffffffffffffffffffffffffffff60045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa801561194a5786906118e7575b73ffffffffffffffffffffffffffffffffffffffff91501633036118bb57606461ffff910135911693841515938480956118ac575b156117f55761ffff600a54168087106117c5575061178995506116276116176115fd86614174565b67ffffffffffffffff16600052600b602052604060002090565b8461162184614219565b91615762565b7f7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b267ffffffffffffffff61166361165d87614174565b93614219565b6040805173ffffffffffffffffffffffffffffffffffffffff929092168252602082018790529190931692a25b508092611793575b506110b1611758917f00000000000000000000000000000000000000000000000000000000000000006116ec857f0000000000000000000000000000000000000000000000000000000000000000836147c7565b7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff61171f84614174565b6040805173ffffffffffffffffffffffffffffffffffffffff90951685523360208601528401889052169180606081015b0390a2614174565b90611761615244565b6040519261176e84613c75565b8352602083015260405192839260408452604084019061409a565b9060208301520390f35b6117589192506117bd6110b1916127106117b661ffff600a5460101c1683614701565b049061527f565b929150611698565b82604491887fe08f03ef000000000000000000000000000000000000000000000000000000008352600452602452fd5b50611789945067ffffffffffffffff61180d84614174565b1680825260076020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944838061187f6040862073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615762565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a2611690565b5061ffff600a541615156115d5565b6024857f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d602011611942575b8161190160209383613cc9565b810103126109fa575173ffffffffffffffffffffffffffffffffffffffff811681036109fa5773ffffffffffffffffffffffffffffffffffffffff906115a0565b3d91506118f4565b6040513d88823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008652600452602485fd5b73ffffffffffffffffffffffffffffffffffffffff166119ad816000526003602052604060002054151590565b611514577fd0d25976000000000000000000000000000000000000000000000000000000008652600452602485fd5b6004857f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b611a26915060203d602011611a2c575b611a1e8183613cc9565b810190614e92565b386114dd565b503d611a14565b60248473ffffffffffffffffffffffffffffffffffffffff611a5486614219565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835216600452fd5b50346102c857806003193601126102c857602073ffffffffffffffffffffffffffffffffffffffff60045416604051908152f35b50346102c85760206003193601126102c857610e2f610e2a60026040610e889467ffffffffffffffff611ae2613e27565b611aea61458d565b50168152600760205220016145b8565b50346102c85767ffffffffffffffff611b1236614028565b929091611b1d61477c565b1691611b36836000526006602052604060002054151590565b15611bf2578284526007602052611b6560056040862001611b58368486613ebd565b6020815191012090615ccd565b15611baa57907f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d7691611ba460405192839260208452602084019161454e565b0390a280f35b82611bee836040519384937f74f23c7c000000000000000000000000000000000000000000000000000000008552600485015260406024850152604484019161454e565b0390fd5b602484847f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50346102c857806003193601126102c85760405160028054808352908352909160208301917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace915b818110611c8957610e8885611c7d81870382613cc9565b60405191829182613f8a565b8254845260209093019260019283019201611c66565b50346102c85760206003193601126102c85767ffffffffffffffff611cc2613e27565b1681526007602052611cd9600560408320016159a9565b80517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611d1e611d08836142da565b92611d166040519485613cc9565b8084526142da565b01835b818110611df5575050825b8251811015611d725780611d4260019285614427565b5185526008602052611d566040862061448e565b611d608285614427565b52611d6b8184614427565b5001611d2c565b81846040519182916020830160208452825180915260408401602060408360051b870101940192905b828210611daa57505050500390f35b91936020611de5827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851613d44565b9601920192018594939192611d9b565b806060602080938601015201611d21565b50346102c857611c7d60016040610e889367ffffffffffffffff611e2936613f12565b505050509050168152600d6020522001614277565b50346102c85760206003193601126102c8576004359067ffffffffffffffff82116102c85760a060031983360301126102c857611e7961440e565b5060848201611e8781614219565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691160361242e57602483019277ffffffffffffffff00000000000000000000000000000000611eed85614174565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561235557849161240f575b506123e757611f7b60448201614219565b7f000000000000000000000000000000000000000000000000000000000000000061238b575b5067ffffffffffffffff611fb485614174565b16611fcc816000526006602052604060002054151590565b1561236057602073ffffffffffffffffffffffffffffffffffffffff60045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa80156123555784906122f2575b73ffffffffffffffffffffffffffffffffffffffff91501633036122c657606401359180156122065761ffff600a5416806121d6575050826121a6926110b19261208961207f6115fd610e8898614174565b8361162184614219565b7f7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b267ffffffffffffffff6120bf61165d86614174565b6040805173ffffffffffffffffffffffffffffffffffffffff929092168252602082018690529190931692a25b7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000092612163817f0000000000000000000000000000000000000000000000000000000000000000866147c7565b61175061216f86614174565b6040805173ffffffffffffffffffffffffffffffffffffffff90971687523360208801528601929092529116929081906060820190565b6121ae615244565b604051916121bb83613c75565b8252602082015260405191829160208352602083019061409a565b604492507fe08f03ef00000000000000000000000000000000000000000000000000000000825281600452602452fd5b50506110b1826121a69267ffffffffffffffff612225610e8896614174565b168060005260076020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789448280612299604060002073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615762565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a26120ec565b6024837f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d60201161234d575b8161230c60209383613cc9565b81010312610a03575173ffffffffffffffffffffffffffffffffffffffff81168103610a035773ffffffffffffffffffffffffffffffffffffffff9061202d565b3d91506122ff565b6040513d86823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008452600452602483fd5b73ffffffffffffffffffffffffffffffffffffffff166123b8816000526003602052604060002054151590565b611fa1577fd0d25976000000000000000000000000000000000000000000000000000000008452600452602483fd5b6004837f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b612428915060203d602011611a2c57611a1e8183613cc9565b38611f6a565b9073ffffffffffffffffffffffffffffffffffffffff611a54602493614219565b50346102c85760606003193601126102c85760043567ffffffffffffffff8111610dd757612481903690600401613da3565b60243567ffffffffffffffff8111610a03576124a1903690600401614069565b60449291923567ffffffffffffffff81116109fa576124c4903690600401614069565b91909273ffffffffffffffffffffffffffffffffffffffff60095416331415806125b3575b6125875781811480159061257d575b61255557865b818110612509578780f35b8061254f61251d610a26600194868c6141b8565b61252883878b6143fe565b612549612541612539868b8d6143fe565b9236906140ee565b9136906140ee565b91615114565b016124fe565b6004877f568efce2000000000000000000000000000000000000000000000000000000008152fd5b50828114156124f8565b6024877f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b5073ffffffffffffffffffffffffffffffffffffffff600154163314156124e9565b50346102c857806003193601126102c857602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b50346102c85760206003193601126102c857602061264567ffffffffffffffff612631613e27565b166000526006602052604060002054151590565b6040519015158152f35b50346102c85760206003193601126102c85760043567ffffffffffffffff8111610dd757612684610d86913690600401613df6565b9061268d61477c565b61496b565b50346102c85760206003193601126102c8577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174602073ffffffffffffffffffffffffffffffffffffffff6126e4613bca565b6126ec61477c565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000006009541617600955604051908152a180f35b50346102c857806003193601126102c857805473ffffffffffffffffffffffffffffffffffffffff811633036127c0577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516825573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b50346102c857806003193601126102c857602073ffffffffffffffffffffffffffffffffffffffff60095416604051908152f35b50346102c85761282b36614028565b6128379392919361477c565b67ffffffffffffffff8216612859816000526006602052604060002054151590565b156128755750610d86929361286f913691613ebd565b90614eaa565b7f1e670e4b000000000000000000000000000000000000000000000000000000008452600452602483fd5b50346102c857806003193601126102c85760206128bb614346565b604051908152f35b50346102c8576128ed906128f56128d936613fda565b95916128e693919361477c565b36916142f2565b9336916142f2565b7f000000000000000000000000000000000000000000000000000000000000000015612a1857815b8351811015612990578073ffffffffffffffffffffffffffffffffffffffff61294860019387614427565b511661295381615a0c565b61295f575b500161291d565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a138612958565b5090805b8251811015612a14578073ffffffffffffffffffffffffffffffffffffffff6129bf60019386614427565b51168015612a0e576129d081615e01565b6129dd575b505b01612994565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a1846129d5565b506129d7565b5080f35b6004827f35f4a7b3000000000000000000000000000000000000000000000000000000008152fd5b50346102c857611c7d6040610e889267ffffffffffffffff612a6136613f12565b505050509050168152600d60205220614277565b50346102c85760406003193601126102c857612a8f613e27565b906024359067ffffffffffffffff82116102c857602061264584612ab63660048701613ef4565b9061423a565b50346102c85760406003193601126102c85760043567ffffffffffffffff8111610dd757806004019161010060031983360301126102c857612afc613dd4565b9181604051612b0a81613c59565b52612b37612b2d612b28612b2160c48501886141c8565b3691613ebd565b614c91565b6064830135614d85565b926084820190612b4682614219565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691160361301757602483019577ffffffffffffffff00000000000000000000000000000000612bac88614174565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561194a578691612ff8575b506119dc5767ffffffffffffffff612c4088614174565b16612c58816000526006602052604060002054151590565b1561195557602073ffffffffffffffffffffffffffffffffffffffff60045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa90811561194a578691612fd9575b50156118bb57612ccf87614174565b90612ce560a4860192612ab6612b2185856141c8565b15612f925750604493929161ffff16159050612edd5767ffffffffffffffff612d0d87614174565b168452600c602052612d26604085208661162184614219565b7f41a8aa8df7945f0fb8ac5f7d88279638d9dc2ef9a6bf4ec9a53b80681b34aff767ffffffffffffffff612d5c61165d89614174565b6040805173ffffffffffffffffffffffffffffffffffffffff929092168252602082018990529190931692a25b0192612d9484614219565b9173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692833b15610dd7576040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91909116600482015260248101859052818160448183885af18015612ed257612ebd575b5050608067ffffffffffffffff60209573ffffffffffffffffffffffffffffffffffffffff612e8b612e857ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc096614174565b92614219565b60405196875233898801521660408601528560608601521692a260405190612eb282613c59565b815260405190518152f35b612ec8828092613cc9565b6102c85780612e33565b6040513d84823e3d90fd5b5067ffffffffffffffff612ef086614174565b1680845260076020527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c8580612f656002604089200173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615762565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a2612d89565b612f9c92506141c8565b611bee6040519283927f24eb47e500000000000000000000000000000000000000000000000000000000845260206004850152602484019161454e565b612ff2915060203d602011611a2c57611a1e8183613cc9565b38612cc0565b613011915060203d602011611a2c57611a1e8183613cc9565b38612c29565b60248473ffffffffffffffffffffffffffffffffffffffff611a5485614219565b50346102c85760206003193601126102c85760043567ffffffffffffffff8111610dd757806004016101006003198336030112610ddb578260405161307c81613c59565b528260405161308a81613c59565b526130ab6130a1612b28612b2160c48601856141c8565b6064840135614d85565b91608481016130b981614219565b73ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116036134255750602481019177ffffffffffffffff0000000000000000000000000000000061312084614174565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561194a578691613406575b506119dc5767ffffffffffffffff6131b484614174565b166131cc816000526006602052604060002054151590565b1561195557602073ffffffffffffffffffffffffffffffffffffffff60045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa90811561194a5786916133e7575b50156118bb5761324383614174565b9061325960a4840192612ab6612b2185856141c8565b15612f92575050906044839267ffffffffffffffff61327784614174565b1680875260076020526132c96002604089200173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016968791615762565b6040805173ffffffffffffffffffffffffffffffffffffffff87168152602081018890527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a2019061331c82614219565b85843b156102c8576040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff929092166004830152602482018690528160448183885af1801561194a579273ffffffffffffffffffffffffffffffffffffffff612e8b612e8560809560209a67ffffffffffffffff967ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0996133d7575b5050614174565b816133e191613cc9565b386133d0565b613400915060203d602011611a2c57611a1e8183613cc9565b38613234565b61341f915060203d602011611a2c57611a1e8183613cc9565b3861319d565b8473ffffffffffffffffffffffffffffffffffffffff611a54602493614219565b50346102c85760a06003193601126102c857613460613bca565b5060243567ffffffffffffffff8116809103610dd75760443567ffffffffffffffff8111610ddb5760031960a09136030112610dd75761349e613de5565b5060843567ffffffffffffffff8111610ddb57916040916134c56080943690600401613e3e565b505081606084516134d581613c0e565b828152826020820152828682015201528152600e60205220604051906134fa82613c0e565b5463ffffffff808216928381528160208201818560201c16815260ff60606040850194848860401c168652019560601c161515855260405195865251166020850152511660408301525115156060820152f35b50346102c85760606003193601126102c85760043561ffff8116809103610dd757613576613dd4565b60443567ffffffffffffffff8111610a0357613596903690600401613df6565b9061359f61477c565b61ffff83169261271084101561362357849260409492613615927f52aa194b292c8bfb5aaca8ee2000a965c3a051b306ff841873b16147526a39ba977fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000063ffff0000600a549360101b1692161717600a5561496b565b82519182526020820152a180f35b602486857f95f3517a000000000000000000000000000000000000000000000000000000008252600452fd5b50346102c85760406003193601126102c85760043567ffffffffffffffff8111610dd75736602382011215610dd757806004013567ffffffffffffffff8111610ddb5760248201916024369160a08402010111610ddb5760243567ffffffffffffffff8111610a03576136c6903690600401613da3565b9190926136d161477c565b84805b83811061373e57509150505b8181106136eb578380f35b8067ffffffffffffffff613705610a2660019486886141b8565b16808652600e6020528560408120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee88680a2016136e0565b6001917f56f77aeff2def50c8b5f5a0df3bab7183df09bf36c6feba496bb42551db77d7060808561388c61385263ffffffff613881613845826138766137958f8061378f8f9283610a26918e614135565b9a614135565b604067ffffffffffffffff602083019a169c8d8152600e60205220836137ba8b614189565b169181549060408101937fffffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffff67ffffffff000000006137f787614189565b60201b16918f6cff0000000000000000000000007fffffffffffffffffffffffffffffffffffffffff000000000000000000000000916bffffffff0000000000000000606088019d8e614189565b60401b1696019e8f61419a565b151560601b169516171617171790558261386e6040519a6141a7565b1689526141a7565b1660208701526141a7565b1660408401526140c4565b15156060820152a20185906136d4565b50346102c857806003193601126102c857602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102c85760206003193601126102c8576020906138f7613bca565b905073ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b50346102c857806003193601126102c857602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102c857806003193601126102c85750610e886040516139b2604082613cc9565b602081527f4275726e546f41646472657373546f6b656e506f6f6c20312e362e332d6465766020820152604051918291602083526020830190613d44565b50346102c85760206003193601126102c857613a0a613bca565b613a1261477c565b613a1a614346565b9081613a24578280f35b602073ffffffffffffffffffffffffffffffffffffffff82613a88857f738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b859599957f00000000000000000000000000000000000000000000000000000000000000006147c7565b6040519485521692a238808280f35b905034610dd7576020600319360112610dd7576004357fffffffff000000000000000000000000000000000000000000000000000000008116809103610ddb57602092507ff208a58f000000000000000000000000000000000000000000000000000000008114908115613ba0575b8115613b76575b8115613b4c575b8115613b22575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438613b1b565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150613b14565b7f1ef5498f0000000000000000000000000000000000000000000000000000000081149150613b0d565b7faff2afbf0000000000000000000000000000000000000000000000000000000081149150613b06565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036109fe57565b359073ffffffffffffffffffffffffffffffffffffffff821682036109fe57565b6080810190811067ffffffffffffffff821117613c2a57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff821117613c2a57604052565b6040810190811067ffffffffffffffff821117613c2a57604052565b60a0810190811067ffffffffffffffff821117613c2a57604052565b6060810190811067ffffffffffffffff821117613c2a57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117613c2a57604052565b67ffffffffffffffff8111613c2a57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b848110613d8e5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201613d4f565b9181601f840112156109fe5782359167ffffffffffffffff83116109fe576020808501948460051b0101116109fe57565b6024359061ffff821682036109fe57565b6064359061ffff821682036109fe57565b9181601f840112156109fe5782359167ffffffffffffffff83116109fe5760208085019460e085020101116109fe57565b6004359067ffffffffffffffff821682036109fe57565b9181601f840112156109fe5782359167ffffffffffffffff83116109fe57602083818601950101116109fe57565b346109fe5760006003193601126109fe57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b929192613ec982613d0a565b91613ed76040519384613cc9565b8294818452818301116109fe578281602093846000960137010152565b9080601f830112156109fe57816020613f0f93359101613ebd565b90565b60a06003198201126109fe5760043573ffffffffffffffffffffffffffffffffffffffff811681036109fe579160243567ffffffffffffffff811681036109fe57916044359160643561ffff811681036109fe57916084359067ffffffffffffffff82116109fe57613f8691600401613e3e565b9091565b602060408183019282815284518094520192019060005b818110613fae5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101613fa1565b60406003198201126109fe5760043567ffffffffffffffff81116109fe578161400591600401613da3565b929092916024359067ffffffffffffffff82116109fe57613f8691600401613da3565b9060406003198301126109fe5760043567ffffffffffffffff811681036109fe57916024359067ffffffffffffffff82116109fe57613f8691600401613e3e565b9181601f840112156109fe5782359167ffffffffffffffff83116109fe57602080850194606085020101116109fe57565b613f0f9160206140b38351604084526040840190613d44565b920151906020818403910152613d44565b359081151582036109fe57565b35906fffffffffffffffffffffffffffffffff821682036109fe57565b91908260609103126109fe5760405161410681613cad565b6040614130818395614117816140c4565b8552614125602082016140d1565b6020860152016140d1565b910152565b91908110156141455760a0020190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3567ffffffffffffffff811681036109fe5790565b3563ffffffff811681036109fe5790565b3580151581036109fe5790565b359063ffffffff821682036109fe57565b91908110156141455760051b0190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156109fe570180359067ffffffffffffffff82116109fe576020019181360383136109fe57565b3573ffffffffffffffffffffffffffffffffffffffff811681036109fe5790565b9067ffffffffffffffff613f0f92166000526007602052600560406000200190602081519101209060019160005201602052604060002054151590565b906040519182815491828252602082019060005260206000209260005b8181106142ab5750506142a992500383613cc9565b565b845473ffffffffffffffffffffffffffffffffffffffff16835260019485019487945060209093019201614294565b67ffffffffffffffff8111613c2a5760051b60200190565b9291906142fe816142da565b9361430c6040519586613cc9565b602085838152019160051b81019283116109fe57905b82821061432e57505050565b6020809161433b84613bed565b815201910190614322565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156143f2576000916143c3575090565b90506020813d6020116143ea575b816143de60209383613cc9565b810103126109fe575190565b3d91506143d1565b6040513d6000823e3d90fd5b9190811015614145576060020190565b6040519061441b82613c75565b60606020838281520152565b80518210156141455760209160051b010190565b90600182811c92168015614484575b602083101461445557565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f169161444a565b90604051918260008254926144a28461443b565b808452936001811690811561450e57506001146144c7575b506142a992500383613cc9565b90506000929192526020600020906000915b8183106144f25750509060206142a992820101386144ba565b60209193508060019154838589010152019101909184926144d9565b602093506142a99592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b820101386144ba565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b6040519061459a82613c91565b60006080838281528260208201528260408201528260608201520152565b906040516145c581613c91565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b91908110156141455760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa1813603018212156109fe570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156109fe570180359067ffffffffffffffff82116109fe57602001918160051b360383136109fe57565b9160209082815201919060005b8181106146cc5750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff6146f388613bed565b1681520194019291016146bf565b8181029291811591840414171561471457565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b81811061474e575050565b60008155600101614743565b67ffffffffffffffff166000526007602052613f0f600460406000200161448e565b73ffffffffffffffffffffffffffffffffffffffff60015416330361479d57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b6040517fa9059cbb000000000000000000000000000000000000000000000000000000006020820190815273ffffffffffffffffffffffffffffffffffffffff938416602483015260448083019590955293815261489d92909161482c606484613cc9565b166000806040958651946148408887613cc9565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020870152519082855af13d15614946573d9161488183613d0a565b9261488e87519485613cc9565b83523d6000602085013e615f10565b805190816148aa57505050565b6020806148bb938301019101614e92565b156148c35750565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606091615f10565b356fffffffffffffffffffffffffffffffff811681036109fe5790565b9160005b82811015614c8b5760e081028401600061498882614174565b9067ffffffffffffffff8216916149ac836000526006602052604060002054151590565b15614c5f57614a759260408593614a20614a1a94614a1a6149e0602060019c9b01926115fd6149db36866140ee565b6153e8565b91825463ffffffff8160801c16159081614c41575b81614c32575b81614c17575b81614c08575b5080614bf9575b614b6e575b36906140ee565b9061552f565b6080850192614a326149db36866140ee565b8152600c6020522092835463ffffffff8160801c16159081614b50575b81614b41575b81614b26575b81614b17575b5080614b08575b614a7b575b5036906140ee565b0161496f565b614a9860a06fffffffffffffffffffffffffffffffff920161494e565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff161717835538614a6d565b50614b128261419a565b614a68565b60ff915060a01c161538614a61565b6fffffffffffffffffffffffffffffffff8116159150614a5b565b8589015460801c159150614a55565b858901546fffffffffffffffffffffffffffffffff16159150614a4f565b6fffffffffffffffffffffffffffffffff614b8a878b0161494e565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff1617178355614a13565b50614c038161419a565b614a0e565b60ff915060a01c161538614a07565b6fffffffffffffffffffffffffffffffff8116159150614a01565b848e015460801c1591506149fb565b848e01546fffffffffffffffffffffffffffffffff161591506149f5565b506024917f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50915050565b80518015614d0157602003614cc35780516020828101918301839003126109fe57519060ff8211614cc3575060ff1690565b611bee906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190613d44565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff821161471457565b60ff16604d811161471457600a0a90565b8115614d56570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff811692828414614e8b57828411614e615790614dca91614d27565b91604d60ff8416118015614e28575b614df257505090614dec613f0f92614d3b565b90614701565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b50614e3283614d3b565b8015614d56577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048411614dd9565b614e6a91614d27565b91604d60ff841611614df257505090614e85613f0f92614d3b565b90614d4c565b5050505090565b908160209103126109fe575180151581036109fe5790565b908051156150ea5767ffffffffffffffff81516020830120921691826000526007602052614edf816005604060002001615ebb565b156150a65760005260086020526040600020815167ffffffffffffffff8111613c2a57614f0c825461443b565b601f8111615074575b506020601f8211600114614fae5791614f88827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9593614f9e95600091614fa3575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055604051918291602083526020830190613d44565b0390a2565b905084015138614f57565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b81811061505c575092614f9e9492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610615025575b5050811b0190556110b6565b8501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880615019565b9192602060018192868a015181550194019201614fde565b6150a090836000526020600020601f840160051c8101916020851061092957601f0160051c0190614743565b38614f15565b5090611bee6040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190613d44565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b67ffffffffffffffff166000818152600660205260409020549092919015615216579161521360e0926151df8561516b7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b976153e8565b84600052600760205261518281604060002061552f565b61518b836153e8565b8460005260076020526151a583600260406000200161552f565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016602082015260208152613f0f604082613cc9565b9190820391821161471457565b61529461458d565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff80835116916152f160208501936152eb6152de63ffffffff8751164261527f565b8560808901511690614701565b90615311565b8082101561530a57505b16825263ffffffff4216905290565b90506152fb565b9190820180921161471457565b805160005b81811061532f57505050565b60018101808211614714575b82811061534b5750600101615323565b73ffffffffffffffffffffffffffffffffffffffff61536a8386614427565b511673ffffffffffffffffffffffffffffffffffffffff61538b8387614427565b51161461539a5760010161533b565b73ffffffffffffffffffffffffffffffffffffffff6153b98386614427565b51167fa1726e400000000000000000000000000000000000000000000000000000000060005260045260246000fd5b805115615488576fffffffffffffffffffffffffffffffff6040820151166fffffffffffffffffffffffffffffffff602083015116106154255750565b606490615486604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff60408201511615801590615510575b6154af5750565b606490615486604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff60208201511615156154a8565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1991615668606092805461556c63ffffffff8260801c164261527f565b90816156a7575b50506fffffffffffffffffffffffffffffffff600181602086015116928281541680851060001461569f57508280855b16167fffffffffffffffffffffffffffffffff0000000000000000000000000000000082541617815561561c8651151582907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b60408601517fffffffffffffffffffffffffffffffff0000000000000000000000000000000060809190911b16939092166fffffffffffffffffffffffffffffffff1692909217910155565b61521360405180926fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b8380916155a3565b6fffffffffffffffffffffffffffffffff916156dc8392836156d56001880154948286169560801c90614701565b9116615311565b8082101561575b57505b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff9290911692909216167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116174260801b73ffffffff00000000000000000000000000000000161781553880615573565b90506156e6565b9182549060ff8260a01c161580156159a1575b61599b576fffffffffffffffffffffffffffffffff821691600185019081546157ba63ffffffff6fffffffffffffffffffffffffffffffff83169360801c164261527f565b90816158fd575b50508481106158b1575083831061581b5750506157f06fffffffffffffffffffffffffffffffff92839261527f565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b5460801c9161582a818561527f565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101908082116147145761587861587d9273ffffffffffffffffffffffffffffffffffffffff96615311565b614d4c565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828573ffffffffffffffffffffffffffffffffffffffff927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b82869293961161597157615918926152eb9160801c90614701565b8084101561596c5750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff00000000000000000000000000000000161786559238806157c1565b615923565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b508215615775565b906040519182815491828252602082019060005260206000209260005b8181106159db5750506142a992500383613cc9565b84548352600194850194879450602090930192016159c6565b80548210156141455760005260206000200190600090565b6000818152600360205260409020548015615b9b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161471457600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161471457818103615b2c575b5050506002548015615afd577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01615aba8160026159f4565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b615b83615b3d615b4e9360026159f4565b90549060031b1c92839260026159f4565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526003602052604060002055388080615a81565b5050600090565b6000818152600660205260409020548015615b9b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161471457600554907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161471457818103615c93575b5050506005548015615afd577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01615c508160056159f4565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600555600052600660205260006040812055600190565b615cb5615ca4615b4e9360056159f4565b90549060031b1c92839260056159f4565b90556000526006602052604060002055388080615c17565b9060018201918160005282602052604060002054801515600014615df8577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111614714578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161471457818103615dc1575b50505080548015615afd577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190615d8282826159f4565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b615de1615dd1615b4e93866159f4565b90549060031b1c928392866159f4565b905560005283602052604060002055388080615d4a565b50505050600090565b80600052600360205260406000205415600014615e5b5760025468010000000000000000811015613c2a57615e42615b4e82600185940160025560026159f4565b9055600254906000526003602052604060002055600190565b50600090565b80600052600660205260406000205415600014615e5b5760055468010000000000000000811015613c2a57615ea2615b4e82600185940160055560056159f4565b9055600554906000526006602052604060002055600190565b6000828152600182016020526040902054615b9b5780549068010000000000000000821015613c2a5782615ef9615b4e8460018096018555846159f4565b905580549260005201602052604060002055600190565b91929015615f8b5750815115615f24575090565b3b15615f2d5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015615f9e5750805190602001fd5b611bee906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190613d4456fea164736f6c634300081a000a",
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) GetRequiredInboundCCVs(opts *bind.CallOpts, arg0 common.Address, sourceChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "getRequiredInboundCCVs", arg0, sourceChainSelector, arg2, arg3, arg4)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) GetRequiredInboundCCVs(arg0 common.Address, sourceChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _BurnToAddressMintTokenPool.Contract.GetRequiredInboundCCVs(&_BurnToAddressMintTokenPool.CallOpts, arg0, sourceChainSelector, arg2, arg3, arg4)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) GetRequiredInboundCCVs(arg0 common.Address, sourceChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _BurnToAddressMintTokenPool.Contract.GetRequiredInboundCCVs(&_BurnToAddressMintTokenPool.CallOpts, arg0, sourceChainSelector, arg2, arg3, arg4)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) GetRequiredOutboundCCVs(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "getRequiredOutboundCCVs", arg0, destChainSelector, arg2, arg3, arg4)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) GetRequiredOutboundCCVs(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _BurnToAddressMintTokenPool.Contract.GetRequiredOutboundCCVs(&_BurnToAddressMintTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) GetRequiredOutboundCCVs(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _BurnToAddressMintTokenPool.Contract.GetRequiredOutboundCCVs(&_BurnToAddressMintTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) GetRouter() (common.Address, error) {
	return _BurnToAddressMintTokenPool.Contract.GetRouter(&_BurnToAddressMintTokenPool.CallOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) GetRouter() (common.Address, error) {
	return _BurnToAddressMintTokenPool.Contract.GetRouter(&_BurnToAddressMintTokenPool.CallOpts)
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactor) SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.contract.Transact(opts, "setRateLimitAdmin", rateLimitAdmin)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.SetRateLimitAdmin(&_BurnToAddressMintTokenPool.TransactOpts, rateLimitAdmin)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.SetRateLimitAdmin(&_BurnToAddressMintTokenPool.TransactOpts, rateLimitAdmin)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactor) SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.contract.Transact(opts, "setRouter", newRouter)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.SetRouter(&_BurnToAddressMintTokenPool.TransactOpts, newRouter)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.SetRouter(&_BurnToAddressMintTokenPool.TransactOpts, newRouter)
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
	RemoteChainSelector uint64
	OutboundCCVs        []common.Address
	InboundCCVs         []common.Address
	Raw                 types.Log
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

type BurnToAddressMintTokenPoolRouterUpdatedIterator struct {
	Event *BurnToAddressMintTokenPoolRouterUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolRouterUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolRouterUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnToAddressMintTokenPoolRouterUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnToAddressMintTokenPoolRouterUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolRouterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolRouterUpdated struct {
	OldRouter common.Address
	NewRouter common.Address
	Raw       types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterRouterUpdated(opts *bind.FilterOpts) (*BurnToAddressMintTokenPoolRouterUpdatedIterator, error) {

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolRouterUpdatedIterator{contract: _BurnToAddressMintTokenPool.contract, event: "RouterUpdated", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolRouterUpdated) (event.Subscription, error) {

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolRouterUpdated)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseRouterUpdated(log types.Log) (*BurnToAddressMintTokenPoolRouterUpdated, error) {
	event := new(BurnToAddressMintTokenPoolRouterUpdated)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
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

func (BurnToAddressMintTokenPoolAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (BurnToAddressMintTokenPoolAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (BurnToAddressMintTokenPoolCCVConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xb0897119e8510f887b892cbc4c8506fc51d9849fd90afae4fd065e705f2d0f6c")
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

func (BurnToAddressMintTokenPoolRouterUpdated) Topic() common.Hash {
	return common.HexToHash("0x02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f1684")
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

	GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRequiredInboundCCVs(opts *bind.CallOpts, arg0 common.Address, sourceChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error)

	GetRequiredOutboundCCVs(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

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

	SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error)

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

	FilterRouterUpdated(opts *bind.FilterOpts) (*BurnToAddressMintTokenPoolRouterUpdatedIterator, error)

	WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolRouterUpdated) (event.Subscription, error)

	ParseRouterUpdated(log types.Log) (*BurnToAddressMintTokenPoolRouterUpdated, error)

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnToAddressMintTokenPoolTokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*BurnToAddressMintTokenPoolTokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnToAddressMintTokenPoolTokenTransferFeeConfigUpdatedIterator, error)

	WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigUpdated(log types.Log) (*BurnToAddressMintTokenPoolTokenTransferFeeConfigUpdated, error)

	Address() common.Address
}
