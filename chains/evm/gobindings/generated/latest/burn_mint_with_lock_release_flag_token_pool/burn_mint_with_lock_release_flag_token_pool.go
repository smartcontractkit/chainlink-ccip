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

type TokenPoolFastFinalityRateLimitConfigArgs struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
}

type TokenPoolTokenTransferFeeConfigArgs struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
}

var BurnMintWithLockReleaseFlagTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIBurnMintERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyFinalityConfigUpdates\",\"inputs\":[{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fastTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxAmountPerRequest\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.FastFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"destToUseDefaultFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAccumulatedFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredInboundCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredOutboundCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFastFinalityRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.FastFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRouter\",\"inputs\":[{\"name\":\"newRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FastTransferInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FastTransferOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalityConfigUpdated\",\"inputs\":[{\"name\":\"finalityConfig\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"fastTransferFeeBps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"maxAmountPerRequest\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeeWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RouterUpdated\",\"inputs\":[{\"name\":\"oldRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AmountExceedsMaxPerRequest\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidFastTransferFeeBps\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFinality\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidFinalityConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenTransferFeeConfigNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x61010080604052346103c557616802803803809161001d8285610444565b833981019060a0818303126103c55780516001600160a01b038116908190036103c55761004c60208301610467565b60408301519091906001600160401b0381116103c55783019380601f860112156103c5578451946001600160401b03861161042e578560051b9060208201966100986040519889610444565b87526020808801928201019283116103c557602001905b828210610416575050506100d160806100ca60608601610475565b9401610475565b92331561040557600180546001600160a01b03191633179055811580156103f4575b80156103e3575b6103d2578160209160049360805260c0526040519283809263313ce56760e01b82525afa60009181610391575b50610366575b5060a052600480546001600160a01b0319166001600160a01b03929092169190911790558051151560e0819052610248575b6040516161d8908161062a82396080518181816113f8015281816116aa0152818161188a01528181611ed601528181612170015281816122eb01528181612ca201528181612ef00152818161308d0152818161328a0152818161345a015281816139f801528181613a5201528181613b83015281816145b201526153f3015260a0518181816117330152818161399d01528181614d8d01528181614e960152615020015260c051818181610bf40152818161149301528181611f7101528181612d3d0152613326015260e051818181610ba2015281816114d601528181611fb40152612a3b0152f35b60405160206102578183610444565b60008252600036813760e051156103555760005b82518110156102d2576001906001600160a01b036102898286610489565b511683610295826104cb565b6102a2575b50500161026b565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a1388361029a565b50905060005b825181101561034c576001906001600160a01b036102f68286610489565b511680156103465783610308826105c9565b610316575b50505b016102d8565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a1388361030d565b50610310565b5050503861015f565b6335f4a7b360e01b60005260046000fd5b60ff1660ff821681810361037a575061012d565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d6020116103ca575b816103ad60209383610444565b810103126103c5576103be90610467565b9038610127565b600080fd5b3d91506103a0565b630a64406560e11b60005260046000fd5b506001600160a01b038116156100fa565b506001600160a01b038416156100f3565b639b15e16f60e01b60005260046000fd5b6020809161042384610475565b8152019101906100af565b634e487b7160e01b600052604160045260246000fd5b601f909101601f19168101906001600160401b0382119082101761042e57604052565b519060ff821682036103c557565b51906001600160a01b03821682036103c557565b805182101561049d5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b805482101561049d5760005260206000200190600090565b60008181526003602052604090205480156105c25760001981018181116105ac576002546000198101919082116105ac5781810361055b575b5050506002548015610545576000190161051f8160026104b3565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61059461056c61057d9360026104b3565b90549060031b1c92839260026104b3565b819391549060031b91821b91600019901b19161790565b90556000526003602052604060002055388080610504565b634e487b7160e01b600052601160045260246000fd5b5050600090565b80600052600360205260406000205415600014610623576002546801000000000000000081101561042e5761060a61057d82600185940160025560026104b3565b9055600254906000526003602052604060002055600190565b5060009056fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a714613d0b57508063164e68de14613ad7578063181f5a7714613a7657806321df0da714613a25578063240028e8146139c157806324f65ee7146139835780632a10097b14613736578063364b7376146136f357806337b19247146135ec5780633907753714613212578063489a68f214612c005780634c5ef0ed14612bb95780634f71592c14612b8457806354c8a4f314612a075780635df45a37146129e457806362ddd3c4146129605780636d3d1a581461292c57806379ba5097146128655780637d54534e146127d65780638926f54f146127905780638da5cb5b1461275c5780638f4622311461264f578063962d4020146124c95780639a4575b914611e755780639f68f67314611e3d578063a42a7b8b14611cd6578063a7cd63b714611c55578063acfecf9114611b31578063af58d59f14611ae8578063b0f479a114611ab4578063b1c71c651461136c578063b79019b5146110af578063b794658014611076578063c0d7865514610f9c578063c4bffe2b14610e71578063c75eea9c14610dc9578063cf7401f314610c18578063dc0bd97114610bc7578063e0351e1314610b8a578063e8a1da17146102b55763f2fde38b146101e457600080fd5b346102b25760206003193601126102b25773ffffffffffffffffffffffffffffffffffffffff610212613e3e565b61021a61499f565b1633811461028a57807fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b80fd5b50346102b2576102c4366141fd565b939190926102d061499f565b82915b8083106109f5575050508063ffffffff4216917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee1843603015b858210156109f1578160051b850135818112156109ed57850190610120823603126109ed576040519561033e87613f05565b823567ffffffffffffffff811681036109e8578752602083013567ffffffffffffffff81116109e45783019536601f880112156109e457863596610381886144fd565b9761038f604051998a613f3d565b8089526020808a019160051b830101903682116109e05760208301905b8282106109ad575050505060208801968752604084013567ffffffffffffffff81116109a9576103df9036908601614117565b9860408901998a526104096103f73660608801614311565b9560608b0196875260c0369101614311565b9660808a0197885261041b86516155d7565b61042588516155d7565b8a5151156109815761044167ffffffffffffffff8b5116616050565b1561094a5767ffffffffffffffff8a5116815260076020526040812061058187516fffffffffffffffffffffffffffffffff6040820151169061053c6fffffffffffffffffffffffffffffffff602083015116915115158360806040516104a781613f05565b858152602081018c905260408101849052606081018690520152855474ff000000000000000000000000000000000000000091151560a01b919091167fffffffffffffffffffffff0000000000000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff84161773ffffffff0000000000000000000000000000000060808b901b1617178555565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176001830155565b6106a789516fffffffffffffffffffffffffffffffff604082015116906106626fffffffffffffffffffffffffffffffff602083015116915115158360806040516105cb81613f05565b858152602081018c9052604081018490526060810186905201526002860180547fffffffffffffffffffffff000000000000000000000000000000000000000000166fffffffffffffffffffffffffffffffff85161773ffffffff0000000000000000000000000000000060808c901b161791151560a01b74ff000000000000000000000000000000000000000016919091179055565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176003830155565b60048c5191019080519067ffffffffffffffff821161091d576106ca835461465e565b601f81116108e2575b50602090601f8311600114610843576107219291859183610838575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b805b8951805182101561075c579061075660019261074f838f67ffffffffffffffff9051169261464a565b5190615042565b01610726565b5050975097987f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c29295939661082a67ffffffffffffffff600197949c51169251935191516107f66107c160405196879687526101006020880152610100870190613fb8565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a101909394929161030c565b0151905038806106ef565b83855281852091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416865b8181106108ca5750908460019594939210610893575b505050811b019055610724565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388080610886565b92936020600181928786015181550195019301610870565b61090d9084865260208620601f850160051c81019160208610610913575b601f0160051c0190614966565b386106d3565b9091508190610900565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60249067ffffffffffffffff8b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b807f14c880ca0000000000000000000000000000000000000000000000000000000060049252fd5b8680fd5b813567ffffffffffffffff81116109dc576020916109d18392833691890101614117565b8152019101906103ac565b8a80fd5b8880fd5b8580fd5b600080fd5b8380fd5b8280f35b9092919367ffffffffffffffff610a15610a108785886143db565b614397565b1695610a2087615d91565b15610b5e578684526007602052610a3c60056040862001615b98565b94845b8651811015610a75576001908987526007602052610a6e60056040892001610a67838b61464a565b5190615ebc565b5001610a3f565b5093945094909580855260076020526005604086208681558660018201558660028201558660038201558660048201610aae815461465e565b80610b1d575b5050500180549086815581610aff575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a10191909493946102d3565b865260208620908101905b81811015610ac457868155600101610b0a565b601f8111600114610b335750555b863880610ab4565b81835260208320610b4e91601f01861c810190600101614966565b8082528160208120915555610b2b565b602484887f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50346102b257806003193601126102b25760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b50346102b257806003193601126102b257602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102b25760e06003193601126102b257610c32614079565b9060607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc3601126102b257604051610c6981613f21565b6024358015158103610dc55781526044356fffffffffffffffffffffffffffffffff81168103610dc55760208201526064356fffffffffffffffffffffffffffffffff81168103610dc557604082015260607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c360112610dc15760405190610cf082613f21565b60843580151581036109ed57825260a4356fffffffffffffffffffffffffffffffff811681036109ed57602083015260c4356fffffffffffffffffffffffffffffffff811681036109ed57604083015273ffffffffffffffffffffffffffffffffffffffff6009541633141580610d9f575b610d7357610d7092936152ac565b80f35b6024837f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b5073ffffffffffffffffffffffffffffffffffffffff60015416331415610d62565b5080fd5b8280fd5b50346102b25760206003193601126102b257610e14610e0f6040610e6d9367ffffffffffffffff610df8614079565b610e006147b0565b501681526007602052206147db565b61547b565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b0390f35b50346102b257806003193601126102b257604051906005548083528260208101600584526020842092845b818110610f83575050610eb192500383613f3d565b8151610ed5610ebf826144fd565b91610ecd6040519384613f3d565b8083526144fd565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602083019301368437805b8451811015610f34578067ffffffffffffffff610f216001938861464a565b5116610f2d828661464a565b5201610f02565b50925090604051928392602084019060208552518091526040840192915b818110610f60575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101610f52565b8454835260019485019487945060209093019201610e9c565b50346102b25760206003193601126102b25773ffffffffffffffffffffffffffffffffffffffff610fcb613e3e565b610fd361499f565b16801561104e5760407f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f16849160045490807fffffffffffffffffffffffff000000000000000000000000000000000000000083161760045573ffffffffffffffffffffffffffffffffffffffff8351921682526020820152a180f35b6004827f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b50346102b25760206003193601126102b257610e6d61109b611096614079565b61497d565b604051918291602083526020830190613fb8565b50346102b25760206003193601126102b25760043567ffffffffffffffff8111610dc1576110e1903690600401614017565b6110e961499f565b825b8181106110f6578380f35b611104610a10828486614841565b61111c611112838587614841565b6020810190614881565b907fb0897119e8510f887b892cbc4c8506fc51d9849fd90afae4fd065e705f2d0f6c61115661114c86888a614841565b6040810190614881565b91909261116c611167368784614515565b61550d565b61117a611167368587614515565b6040519461118786613ee9565b611192368284614515565b86526111dc67ffffffffffffffff6111ab368789614515565b9860208901998a52169586956111ce6040519586956040875260408701916148d5565b9184830360208601526148d5565b0390a28652600e60205260408620905180519067ffffffffffffffff821161133f5768010000000000000000821161133f576020908354838555808410611325575b500182885260208820885b8381106112fb575050505060010190519081519167ffffffffffffffff83116112ce576801000000000000000083116112ce5760209082548484558085106112b4575b500190865260208620865b83811061128a57505050506001016110eb565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501611277565b8389528289206112c8918101908601614966565b3861126c565b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501611229565b848a52828a20611339918101908501614966565b3861121e565b6024887f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b50346102b25760606003193601126102b2576004359067ffffffffffffffff82116102b25760a060031983360301126102b2576113a76140a1565b9160443567ffffffffffffffff8111610dc5576113c8903690600401614117565b506113d1614631565b5060848101926113e0846143eb565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611a6a57602482019177ffffffffffffffff0000000000000000000000000000000061144684614397565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611981578591611a3b575b50611a13576114d4604482016143eb565b7f00000000000000000000000000000000000000000000000000000000000000006119b7575b5067ffffffffffffffff61150d84614397565b16611525816000526006602052604060002054151590565b1561198c57602073ffffffffffffffffffffffffffffffffffffffff60045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa801561198157859061191a575b73ffffffffffffffffffffffffffffffffffffffff91501633036118ee57606461ffff910135911692831515938480956118df575b156118285761ffff600a5416908181106117f9575050600b548083116117cb575061178f946116196116096115ef86614397565b67ffffffffffffffff16600052600c602052604060002090565b84611613846143eb565b91615951565b7f251a578f8b9ffb7cb748680c76e7895fe65b3fcf44c2bd03e4f163dc38b2ed0b67ffffffffffffffff61165561164f87614397565b936143eb565b6040805173ffffffffffffffffffffffffffffffffffffffff929092168252602082018790529190931692a25b508092611799575b5061109661172b9161169b846153dc565b6116a481614397565b604080517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16815233602082015290810186905267ffffffffffffffff91909116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae109080606081015b0390a2614397565b9060405160ff7f000000000000000000000000000000000000000000000000000000000000000016602082015260208152611767604082613f3d565b6040519261177484613ee9565b835260208301526040519283926040845260408401906142bd565b9060208301520390f35b61172b9192506117c3611096916127106117bc61ffff600a5460101c1683614924565b049061546e565b92915061168a565b6044927f4cfafb1f000000000000000000000000000000000000000000000000000000008352600452602452fd5b604493507fe08f03ef000000000000000000000000000000000000000000000000000000008352600452602452fd5b5061178f945067ffffffffffffffff61184084614397565b1680825260076020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894483806118b26040862073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615951565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a2611682565b5061ffff600a541615156115bb565b6024847f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d602011611979575b8161193460209383613f3d565b81010312611975575173ffffffffffffffffffffffffffffffffffffffff811681036119755773ffffffffffffffffffffffffffffffffffffffff90611586565b8480fd5b3d9150611927565b6040513d87823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008552600452602484fd5b73ffffffffffffffffffffffffffffffffffffffff166119e4816000526003602052604060002054151590565b6114fa577fd0d25976000000000000000000000000000000000000000000000000000000008552600452602484fd5b6004847f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b611a5d915060203d602011611a63575b611a558183613f3d565b810190614f94565b386114c3565b503d611a4b565b60248373ffffffffffffffffffffffffffffffffffffffff611a8b876143eb565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835216600452fd5b50346102b257806003193601126102b257602073ffffffffffffffffffffffffffffffffffffffff60045416604051908152f35b50346102b25760206003193601126102b257610e14610e0f60026040610e6d9467ffffffffffffffff611b19614079565b611b216147b0565b50168152600760205220016147db565b50346102b25767ffffffffffffffff611b493661424b565b929091611b5461499f565b1691611b6d836000526006602052604060002054151590565b15611c29578284526007602052611b9c60056040862001611b8f3684866140e0565b6020815191012090615ebc565b15611be157907f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d7691611bdb604051928392602084526020840191614771565b0390a280f35b82611c25836040519384937f74f23c7c0000000000000000000000000000000000000000000000000000000085526004850152604060248501526044840191614771565b0390fd5b602484847f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50346102b257806003193601126102b25760405160028054808352908352909160208301917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace915b818110611cc057610e6d85611cb481870382613f3d565b604051918291826141ad565b8254845260209093019260019283019201611c9d565b50346102b25760206003193601126102b25767ffffffffffffffff611cf9614079565b1681526007602052611d1060056040832001615b98565b80517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611d55611d3f836144fd565b92611d4d6040519485613f3d565b8084526144fd565b01835b818110611e2c575050825b8251811015611da95780611d796001928561464a565b5185526008602052611d8d604086206146b1565b611d97828561464a565b52611da2818461464a565b5001611d63565b81846040519182916020830160208452825180915260408401602060408360051b870101940192905b828210611de157505050500390f35b91936020611e1c827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851613fb8565b9601920192018594939192611dd2565b806060602080938601015201611d58565b50346102b257611cb460016040610e6d9367ffffffffffffffff611e6036614135565b505050509050168152600e602052200161449a565b50346102b25760206003193601126102b2576004359067ffffffffffffffff82116102b25760a060031983360301126102b257611eb0614631565b5060848201611ebe816143eb565b73ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116036124a857602483019277ffffffffffffffff00000000000000000000000000000000611f2485614397565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156123cf578491612489575b5061246157611fb2604482016143eb565b7f0000000000000000000000000000000000000000000000000000000000000000612405575b5067ffffffffffffffff611feb85614397565b16612003816000526006602052604060002054151590565b156123da57602073ffffffffffffffffffffffffffffffffffffffff60045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa80156123cf57849061236c575b73ffffffffffffffffffffffffffffffffffffffff915016330361234057606401359180156122805761ffff600a5416806122505750600b54808411612220575050826121c092611096926120cb6120c16115ef610e6d98614397565b83611613846143eb565b7f251a578f8b9ffb7cb748680c76e7895fe65b3fcf44c2bd03e4f163dc38b2ed0b67ffffffffffffffff61210161164f86614397565b6040805173ffffffffffffffffffffffffffffffffffffffff929092168252602082018690529190931692a25b612137816153dc565b7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff61216a84614397565b604080517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681523360208201529081019490945216918060608101611723565b6040517ffa7c07de000000000000000000000000000000000000000000000000000000006020820152602081526121f8604082613f3d565b6040519161220583613ee9565b825260208201526040519182916020835260208301906142bd565b60449250837f4cfafb1f000000000000000000000000000000000000000000000000000000008352600452602452fd5b604492507fe08f03ef00000000000000000000000000000000000000000000000000000000825281600452602452fd5b5050611096826121c09267ffffffffffffffff61229f610e6d96614397565b168060005260076020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789448280612313604060002073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615951565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a261212e565b6024837f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d6020116123c7575b8161238660209383613f3d565b810103126109ed575173ffffffffffffffffffffffffffffffffffffffff811681036109ed5773ffffffffffffffffffffffffffffffffffffffff90612064565b3d9150612379565b6040513d86823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008452600452602483fd5b73ffffffffffffffffffffffffffffffffffffffff16612432816000526003602052604060002054151590565b611fd8577fd0d25976000000000000000000000000000000000000000000000000000000008452600452602483fd5b6004837f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b6124a2915060203d602011611a6357611a558183613f3d565b38611fa1565b9073ffffffffffffffffffffffffffffffffffffffff611a8b6024936143eb565b50346102b25760606003193601126102b25760043567ffffffffffffffff8111610dc1576124fb903690600401614017565b60243567ffffffffffffffff81116109ed5761251b90369060040161428c565b60449291923567ffffffffffffffff81116109e45761253e90369060040161428c565b91909273ffffffffffffffffffffffffffffffffffffffff600954163314158061262d575b612601578181148015906125f7575b6125cf57865b818110612583578780f35b806125c9612597610a10600194868c6143db565b6125a283878b614621565b6125c36125bb6125b3868b8d614621565b923690614311565b913690614311565b916152ac565b01612578565b6004877f568efce2000000000000000000000000000000000000000000000000000000008152fd5b5082811415612572565b6024877f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b5073ffffffffffffffffffffffffffffffffffffffff60015416331415612563565b50346102b25760806003193601126102b25760043561ffff8116809103610dc1576126786140a1565b6044359060643567ffffffffffffffff81116119755761269c903690600401614048565b906126a561499f565b600a5461ffff8416936127108510156127345792869492612720927ffc57ba710bb816f16b18f38549598ef4552c4acd0e0ddb7b40bc5906d25bdc1d987fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000063ffff000060609a9860101b1692161717600a5584600b55614a07565b60405192835260208301526040820152a180f35b6004887fbcd2e49d000000000000000000000000000000000000000000000000000000008152fd5b50346102b257806003193601126102b257602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b50346102b25760206003193601126102b25760206127cc67ffffffffffffffff6127b8614079565b166000526006602052604060002054151590565b6040519015158152f35b50346102b25760206003193601126102b2577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174602073ffffffffffffffffffffffffffffffffffffffff612828613e3e565b61283061499f565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000006009541617600955604051908152a180f35b50346102b257806003193601126102b257805473ffffffffffffffffffffffffffffffffffffffff81163303612904577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516825573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b50346102b257806003193601126102b257602073ffffffffffffffffffffffffffffffffffffffff60095416604051908152f35b50346102b25761296f3661424b565b61297b9392919361499f565b67ffffffffffffffff821661299d816000526006602052604060002054151590565b156129b95750610d7092936129b39136916140e0565b90615042565b7f1e670e4b000000000000000000000000000000000000000000000000000000008452600452602483fd5b50346102b257806003193601126102b25760206129ff614569565b604051908152f35b50346102b257612a3190612a39612a1d366141fd565b9591612a2a93919361499f565b3691614515565b933691614515565b7f000000000000000000000000000000000000000000000000000000000000000015612b5c57815b8351811015612ad4578073ffffffffffffffffffffffffffffffffffffffff612a8c6001938761464a565b5116612a9781615bfb565b612aa3575b5001612a61565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a138612a9c565b5090805b8251811015612b58578073ffffffffffffffffffffffffffffffffffffffff612b036001938661464a565b51168015612b5257612b1481615ff0565b612b21575b505b01612ad8565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a184612b19565b50612b1b565b5080f35b6004827f35f4a7b3000000000000000000000000000000000000000000000000000000008152fd5b50346102b257611cb46040610e6d9267ffffffffffffffff612ba536614135565b505050509050168152600e6020522061449a565b50346102b25760406003193601126102b257612bd3614079565b906024359067ffffffffffffffff82116102b25760206127cc84612bfa3660048701614117565b9061445d565b50346102b25760406003193601126102b25760043567ffffffffffffffff8111610dc157806004019161010060031983360301126102b257612c406140a1565b9181604051612c4e81613ecd565b52612c7b612c71612c6c612c6560c485018861440c565b36916140e0565b614fac565b6064830135614e93565b926084820190612c8a826143eb565b73ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116036131f157602483019577ffffffffffffffff00000000000000000000000000000000612cf088614397565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156131745786916131d2575b506131aa5767ffffffffffffffff612d8488614397565b16612d9c816000526006602052604060002054151590565b1561317f57602073ffffffffffffffffffffffffffffffffffffffff60045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115613174578691613155575b501561312957612e1387614397565b90612e2960a4860192612bfa612c65858561440c565b156130e25750604493929161ffff1615905061302d5767ffffffffffffffff612e5187614397565b168452600d602052612e6a6040852086611613846143eb565b7fec8efbbe1188357e808975e80f998488a9c2747d618ffc39270836058a44428e67ffffffffffffffff612ea061164f89614397565b6040805173ffffffffffffffffffffffffffffffffffffffff929092168252602082018990529190931692a25b0192612ed8846143eb565b9173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692833b15610dc1576040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91909116600482015260248101859052818160448183885af180156130225761300d575b505067ffffffffffffffff602094612ff485612fb661164f7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc096614397565b6040805173ffffffffffffffffffffffffffffffffffffffff9889168152336020820152979091169087015260608601529116929081906080820190565b0390a28060405161300481613ecd565b52604051908152f35b613018828092613f3d565b6102b25780612f77565b6040513d84823e3d90fd5b5067ffffffffffffffff61304086614397565b1680845260076020527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c85806130b56002604089200173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615951565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a2612ecd565b6130ec925061440c565b611c256040519283927f24eb47e5000000000000000000000000000000000000000000000000000000008452602060048501526024840191614771565b6024857f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b61316e915060203d602011611a6357611a558183613f3d565b38612e04565b6040513d88823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008652600452602485fd5b6004857f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b6131eb915060203d602011611a6357611a558183613f3d565b38612d6d565b60248473ffffffffffffffffffffffffffffffffffffffff611a8b856143eb565b50346102b25760206003193601126102b25760043567ffffffffffffffff8111610dc157806004016101006003198336030112610dc5578260405161325681613ecd565b526132646064830135614d8b565b9160848101613272816143eb565b73ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116036135cb5750602481019177ffffffffffffffff000000000000000000000000000000006132d984614397565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156131745786916135ac575b506131aa5767ffffffffffffffff61336d84614397565b16613385816000526006602052604060002054151590565b1561317f57602073ffffffffffffffffffffffffffffffffffffffff60045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa90811561317457869161358d575b5015613129576133fc83614397565b9061341260a4840192612bfa612c65858561440c565b156130e2575050906044839267ffffffffffffffff61343084614397565b1680875260076020526134826002604089200173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016968791615951565b6040805173ffffffffffffffffffffffffffffffffffffffff87168152602081018890527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a201906134d5826143eb565b85843b156102b2576040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff929092166004830152602482018690528160448183885af180156131745785612fb661164f67ffffffffffffffff9560209a7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09896612ff49661357d575b5050614397565b8161358791613f3d565b38613576565b6135a6915060203d602011611a6357611a558183613f3d565b386133ed565b6135c5915060203d602011611a6357611a558183613f3d565b38613356565b8473ffffffffffffffffffffffffffffffffffffffff611a8b6024936143eb565b50346102b25760a06003193601126102b257613606613e3e565b5060243567ffffffffffffffff8116809103610dc15760443567ffffffffffffffff8111610dc55760031960a09136030112610dc157613644614090565b5060843567ffffffffffffffff8111610dc5579160409161366b60809436906004016140b2565b5050816060845161367b81613e82565b828152826020820152828682015201528152600f60205220604051906136a082613e82565b5463ffffffff808216928381528160208201818560201c16815260ff60606040850194848860401c168652019560601c161515855260405195865251166020850152511660408301525115156060820152f35b50346102b25760206003193601126102b25760043567ffffffffffffffff8111610dc157613728610d70913690600401614048565b9061373161499f565b614a07565b50346102b25760406003193601126102b25760043567ffffffffffffffff8111610dc15736602382011215610dc157806004013567ffffffffffffffff8111610dc55760248201916024369160a08402010111610dc55760243567ffffffffffffffff81116109ed576137ad903690600401614017565b9190926137b861499f565b84805b83811061382557509150505b8181106137d2578380f35b8067ffffffffffffffff6137ec610a1060019486886143db565b16808652600f6020528560408120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee88680a2016137c7565b6001917f56f77aeff2def50c8b5f5a0df3bab7183df09bf36c6feba496bb42551db77d7060808561397361393963ffffffff61396861392c8261395d61387c8f806138768f9283610a10918e614358565b9a614358565b604067ffffffffffffffff602083019a169c8d8152600f60205220836138a18b6143ac565b169181549060408101937fffffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffff67ffffffff000000006138de876143ac565b60201b16918f6cff0000000000000000000000007fffffffffffffffffffffffffffffffffffffffff000000000000000000000000916bffffffff0000000000000000606088019d8e6143ac565b60401b1696019e8f6143bd565b151560601b16951617161717179055826139556040519a6143ca565b1689526143ca565b1660208701526143ca565b1660408401526142e7565b15156060820152a20185906137bb565b50346102b257806003193601126102b257602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102b25760206003193601126102b2576020906139de613e3e565b905073ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b50346102b257806003193601126102b257602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102b257806003193601126102b25750610e6d604051613a99604082613f3d565b601b81527f4275726e4d696e74546f6b656e506f6f6c20312e362e332d64657600000000006020820152604051918291602083526020830190613fb8565b50346102b25760206003193601126102b257613af1613e3e565b613af961499f565b613b01614569565b9081613b0b578280f35b6040517fa9059cbb000000000000000000000000000000000000000000000000000000006020820190815273ffffffffffffffffffffffffffffffffffffffff831660248301526044808301859052825290613c1190613b6c606482613f3d565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168680604095865194613bb58887613f3d565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020870152519082855af13d15613d03573d91613bf683613f7e565b92613c0387519485613f3d565b83523d89602085013e6160ff565b805180613c62575b505073ffffffffffffffffffffffffffffffffffffffff7f738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b85959992602092519485521692a238808280f35b90602080613c74938301019101614f94565b15613c80573880613c19565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b6060916160ff565b905034610dc1576020600319360112610dc1576004357fffffffff000000000000000000000000000000000000000000000000000000008116809103610dc557602092507ff208a58f000000000000000000000000000000000000000000000000000000008114908115613e14575b8115613dea575b8115613dc0575b8115613d96575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438613d8f565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150613d88565b7f1ef5498f0000000000000000000000000000000000000000000000000000000081149150613d81565b7faff2afbf0000000000000000000000000000000000000000000000000000000081149150613d7a565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036109e857565b359073ffffffffffffffffffffffffffffffffffffffff821682036109e857565b6080810190811067ffffffffffffffff821117613e9e57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff821117613e9e57604052565b6040810190811067ffffffffffffffff821117613e9e57604052565b60a0810190811067ffffffffffffffff821117613e9e57604052565b6060810190811067ffffffffffffffff821117613e9e57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117613e9e57604052565b67ffffffffffffffff8111613e9e57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b8481106140025750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201613fc3565b9181601f840112156109e85782359167ffffffffffffffff83116109e8576020808501948460051b0101116109e857565b9181601f840112156109e85782359167ffffffffffffffff83116109e85760208085019460e085020101116109e857565b6004359067ffffffffffffffff821682036109e857565b6064359061ffff821682036109e857565b6024359061ffff821682036109e857565b9181601f840112156109e85782359167ffffffffffffffff83116109e857602083818601950101116109e857565b9291926140ec82613f7e565b916140fa6040519384613f3d565b8294818452818301116109e8578281602093846000960137010152565b9080601f830112156109e857816020614132933591016140e0565b90565b60a06003198201126109e85760043573ffffffffffffffffffffffffffffffffffffffff811681036109e8579160243567ffffffffffffffff811681036109e857916044359160643561ffff811681036109e857916084359067ffffffffffffffff82116109e8576141a9916004016140b2565b9091565b602060408183019282815284518094520192019060005b8181106141d15750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016141c4565b60406003198201126109e85760043567ffffffffffffffff81116109e8578161422891600401614017565b929092916024359067ffffffffffffffff82116109e8576141a991600401614017565b9060406003198301126109e85760043567ffffffffffffffff811681036109e857916024359067ffffffffffffffff82116109e8576141a9916004016140b2565b9181601f840112156109e85782359167ffffffffffffffff83116109e857602080850194606085020101116109e857565b6141329160206142d68351604084526040840190613fb8565b920151906020818403910152613fb8565b359081151582036109e857565b35906fffffffffffffffffffffffffffffffff821682036109e857565b91908260609103126109e85760405161432981613f21565b604061435381839561433a816142e7565b8552614348602082016142f4565b6020860152016142f4565b910152565b91908110156143685760a0020190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3567ffffffffffffffff811681036109e85790565b3563ffffffff811681036109e85790565b3580151581036109e85790565b359063ffffffff821682036109e857565b91908110156143685760051b0190565b3573ffffffffffffffffffffffffffffffffffffffff811681036109e85790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156109e8570180359067ffffffffffffffff82116109e8576020019181360383136109e857565b9067ffffffffffffffff61413292166000526007602052600560406000200190602081519101209060019160005201602052604060002054151590565b906040519182815491828252602082019060005260206000209260005b8181106144ce5750506144cc92500383613f3d565b565b845473ffffffffffffffffffffffffffffffffffffffff168352600194850194879450602090930192016144b7565b67ffffffffffffffff8111613e9e5760051b60200190565b929190614521816144fd565b9361452f6040519586613f3d565b602085838152019160051b81019283116109e857905b82821061455157505050565b6020809161455e84613e61565b815201910190614545565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115614615576000916145e6575090565b90506020813d60201161460d575b8161460160209383613f3d565b810103126109e8575190565b3d91506145f4565b6040513d6000823e3d90fd5b9190811015614368576060020190565b6040519061463e82613ee9565b60606020838281520152565b80518210156143685760209160051b010190565b90600182811c921680156146a7575b602083101461467857565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f169161466d565b90604051918260008254926146c58461465e565b808452936001811690811561473157506001146146ea575b506144cc92500383613f3d565b90506000929192526020600020906000915b8183106147155750509060206144cc92820101386146dd565b60209193508060019154838589010152019101909184926146fc565b602093506144cc9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b820101386146dd565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b604051906147bd82613f05565b60006080838281528260208201528260408201528260608201520152565b906040516147e881613f05565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b91908110156143685760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa1813603018212156109e8570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156109e8570180359067ffffffffffffffff82116109e857602001918160051b360383136109e857565b9160209082815201919060005b8181106148ef5750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff61491688613e61565b1681520194019291016148e2565b8181029291811591840414171561493757565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818110614971575050565b60008155600101614966565b67ffffffffffffffff16600052600760205261413260046040600020016146b1565b73ffffffffffffffffffffffffffffffffffffffff6001541633036149c057565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b356fffffffffffffffffffffffffffffffff811681036109e85790565b9160005b82811015614d275760e0810284016000614a2482614397565b9067ffffffffffffffff821691614a48836000526006602052604060002054151590565b15614cfb57614b119260408593614abc614ab694614ab6614a7c602060019c9b01926115ef614a773686614311565b6155d7565b91825463ffffffff8160801c16159081614cdd575b81614cce575b81614cb3575b81614ca4575b5080614c95575b614c0a575b3690614311565b9061571e565b6080850192614ace614a773686614311565b8152600d6020522092835463ffffffff8160801c16159081614bec575b81614bdd575b81614bc2575b81614bb3575b5080614ba4575b614b17575b503690614311565b01614a0b565b614b3460a06fffffffffffffffffffffffffffffffff92016149ea565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff161717835538614b09565b50614bae826143bd565b614b04565b60ff915060a01c161538614afd565b6fffffffffffffffffffffffffffffffff8116159150614af7565b8589015460801c159150614af1565b858901546fffffffffffffffffffffffffffffffff16159150614aeb565b6fffffffffffffffffffffffffffffffff614c26878b016149ea565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff1617178355614aaf565b50614c9f816143bd565b614aaa565b60ff915060a01c161538614aa3565b6fffffffffffffffffffffffffffffffff8116159150614a9d565b848e015460801c159150614a97565b848e01546fffffffffffffffffffffffffffffffff16159150614a91565b506024917f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50915050565b9060ff8091169116039060ff821161493757565b60ff16604d811161493757600a0a90565b8115614d5c570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f000000000000000000000000000000000000000000000000000000000000000060ff81169081600614614e8e5781600611614e63576006614dcc91614d2d565b90604d60ff8316118015614e2a575b614df3575090614ded61413292614d41565b90614924565b90507fa9cb113d00000000000000000000000000000000000000000000000000000000600052600660045260245260445260646000fd5b50614e3482614d41565b8015614d5c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311614ddb565b614e6e906006614d2d565b90604d60ff831611614df3575090614e8861413292614d41565b90614d52565b505090565b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff811692828414614f8d57828411614f695790614ed891614d2d565b91604d60ff8416118015614f30575b614efa57505090614ded61413292614d41565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b50614f3a83614d41565b8015614d5c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048411614ee7565b614f7291614d2d565b91604d60ff841611614efa57505090614e8861413292614d41565b5050505090565b908160209103126109e8575180151581036109e85790565b8051801561501c57602003614fde5780516020828101918301839003126109e857519060ff8211614fde575060ff1690565b611c25906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190613fb8565b50507f000000000000000000000000000000000000000000000000000000000000000090565b908051156152825767ffffffffffffffff815160208301209216918260005260076020526150778160056040600020016160aa565b1561523e5760005260086020526040600020815167ffffffffffffffff8111613e9e576150a4825461465e565b601f811161520c575b506020601f82116001146151465791615120827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea95936151369560009161513b575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055604051918291602083526020830190613fb8565b0390a2565b9050840151386150ef565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b8181106151f45750926151369492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9896106151bd575b5050811b01905561109b565b8501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c1916905538806151b1565b9192602060018192868a015181550194019201615176565b61523890836000526020600020601f840160051c8101916020851061091357601f0160051c0190614966565b386150ad565b5090611c256040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190613fb8565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b67ffffffffffffffff1660008181526006602052604090205490929190156153ae57916153ab60e092615377856153037f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b976155d7565b84600052600760205261531a81604060002061571e565b615323836155d7565b84600052600760205261533d83600260406000200161571e565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690813b156109e857604051907f42966c680000000000000000000000000000000000000000000000000000000082528160248160008096819560048401525af1801561302257615461575050565b8161546b91613f3d565b50565b9190820391821161493757565b6154836147b0565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff80835116916154e060208501936154da6154cd63ffffffff8751164261546e565b8560808901511690614924565b90615500565b808210156154f957505b16825263ffffffff4216905290565b90506154ea565b9190820180921161493757565b805160005b81811061551e57505050565b60018101808211614937575b82811061553a5750600101615512565b73ffffffffffffffffffffffffffffffffffffffff615559838661464a565b511673ffffffffffffffffffffffffffffffffffffffff61557a838761464a565b5116146155895760010161552a565b73ffffffffffffffffffffffffffffffffffffffff6155a8838661464a565b51167fa1726e400000000000000000000000000000000000000000000000000000000060005260045260246000fd5b805115615677576fffffffffffffffffffffffffffffffff6040820151166fffffffffffffffffffffffffffffffff602083015116106156145750565b606490615675604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff604082015116158015906156ff575b61569e5750565b606490615675604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff6020820151161515615697565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1991615857606092805461575b63ffffffff8260801c164261546e565b9081615896575b50506fffffffffffffffffffffffffffffffff600181602086015116928281541680851060001461588e57508280855b16167fffffffffffffffffffffffffffffffff0000000000000000000000000000000082541617815561580b8651151582907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b60408601517fffffffffffffffffffffffffffffffff0000000000000000000000000000000060809190911b16939092166fffffffffffffffffffffffffffffffff1692909217910155565b6153ab60405180926fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b838091615792565b6fffffffffffffffffffffffffffffffff916158cb8392836158c46001880154948286169560801c90614924565b9116615500565b8082101561594a57505b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff9290911692909216167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116174260801b73ffffffff00000000000000000000000000000000161781553880615762565b90506158d5565b9182549060ff8260a01c16158015615b90575b615b8a576fffffffffffffffffffffffffffffffff821691600185019081546159a963ffffffff6fffffffffffffffffffffffffffffffff83169360801c164261546e565b9081615aec575b5050848110615aa05750838310615a0a5750506159df6fffffffffffffffffffffffffffffffff92839261546e565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b5460801c91615a19818561546e565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81019080821161493757615a67615a6c9273ffffffffffffffffffffffffffffffffffffffff96615500565b614d52565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828573ffffffffffffffffffffffffffffffffffffffff927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611615b6057615b07926154da9160801c90614924565b80841015615b5b5750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff00000000000000000000000000000000161786559238806159b0565b615b12565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b508215615964565b906040519182815491828252602082019060005260206000209260005b818110615bca5750506144cc92500383613f3d565b8454835260019485019487945060209093019201615bb5565b80548210156143685760005260206000200190600090565b6000818152600360205260409020548015615d8a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161493757600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161493757818103615d1b575b5050506002548015615cec577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01615ca9816002615be3565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b615d72615d2c615d3d936002615be3565b90549060031b1c9283926002615be3565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526003602052604060002055388080615c70565b5050600090565b6000818152600660205260409020548015615d8a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161493757600554907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161493757818103615e82575b5050506005548015615cec577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01615e3f816005615be3565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600555600052600660205260006040812055600190565b615ea4615e93615d3d936005615be3565b90549060031b1c9283926005615be3565b90556000526006602052604060002055388080615e06565b9060018201918160005282602052604060002054801515600014615fe7577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111614937578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161493757818103615fb0575b50505080548015615cec577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190615f718282615be3565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b615fd0615fc0615d3d9386615be3565b90549060031b1c92839286615be3565b905560005283602052604060002055388080615f39565b50505050600090565b8060005260036020526040600020541560001461604a5760025468010000000000000000811015613e9e57616031615d3d8260018594016002556002615be3565b9055600254906000526003602052604060002055600190565b50600090565b8060005260066020526040600020541560001461604a5760055468010000000000000000811015613e9e57616091615d3d8260018594016005556005615be3565b9055600554906000526006602052604060002055600190565b6000828152600182016020526040902054615d8a5780549068010000000000000000821015613e9e57826160e8615d3d846001809601855584615be3565b905580549260005201602052604060002055600190565b9192901561617a5750815115616113575090565b3b1561611c5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b82519091501561618d5750805190602001fd5b611c25906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190613fb856fea164736f6c634300081a000a",
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

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) ApplyFinalityConfigUpdates(opts *bind.TransactOpts, finalityThreshold uint16, fastTransferFeeBps uint16, maxAmountPerRequest *big.Int, rateLimitConfigArgs []TokenPoolFastFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "applyFinalityConfigUpdates", finalityThreshold, fastTransferFeeBps, maxAmountPerRequest, rateLimitConfigArgs)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) ApplyFinalityConfigUpdates(finalityThreshold uint16, fastTransferFeeBps uint16, maxAmountPerRequest *big.Int, rateLimitConfigArgs []TokenPoolFastFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.ApplyFinalityConfigUpdates(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, finalityThreshold, fastTransferFeeBps, maxAmountPerRequest, rateLimitConfigArgs)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) ApplyFinalityConfigUpdates(finalityThreshold uint16, fastTransferFeeBps uint16, maxAmountPerRequest *big.Int, rateLimitConfigArgs []TokenPoolFastFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.ApplyFinalityConfigUpdates(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, finalityThreshold, fastTransferFeeBps, maxAmountPerRequest, rateLimitConfigArgs)
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

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) SetFastFinalityRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolFastFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "setFastFinalityRateLimitConfig", rateLimitConfigArgs)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) SetFastFinalityRateLimitConfig(rateLimitConfigArgs []TokenPoolFastFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.SetFastFinalityRateLimitConfig(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) SetFastFinalityRateLimitConfig(rateLimitConfigArgs []TokenPoolFastFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.SetFastFinalityRateLimitConfig(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, rateLimitConfigArgs)
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

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "setRouter", newRouter)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.SetRouter(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, newRouter)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.SetRouter(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, newRouter)
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

type BurnMintWithLockReleaseFlagTokenPoolFastTransferInboundRateLimitConsumedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolFastTransferInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolFastTransferInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolFastTransferInboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolFastTransferInboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolFastTransferInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolFastTransferInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolFastTransferInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterFastTransferInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolFastTransferInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "FastTransferInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolFastTransferInboundRateLimitConsumedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "FastTransferInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchFastTransferInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolFastTransferInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "FastTransferInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolFastTransferInboundRateLimitConsumed)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "FastTransferInboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseFastTransferInboundRateLimitConsumed(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolFastTransferInboundRateLimitConsumed, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolFastTransferInboundRateLimitConsumed)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "FastTransferInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolFastTransferOutboundRateLimitConsumedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolFastTransferOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolFastTransferOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolFastTransferOutboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolFastTransferOutboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolFastTransferOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolFastTransferOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolFastTransferOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterFastTransferOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolFastTransferOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "FastTransferOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolFastTransferOutboundRateLimitConsumedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "FastTransferOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchFastTransferOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolFastTransferOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "FastTransferOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolFastTransferOutboundRateLimitConsumed)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "FastTransferOutboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseFastTransferOutboundRateLimitConsumed(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolFastTransferOutboundRateLimitConsumed, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolFastTransferOutboundRateLimitConsumed)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "FastTransferOutboundRateLimitConsumed", log); err != nil {
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
	FinalityConfig      uint16
	FastTransferFeeBps  uint16
	MaxAmountPerRequest *big.Int
	Raw                 types.Log
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

func (BurnMintWithLockReleaseFlagTokenPoolFastTransferInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xec8efbbe1188357e808975e80f998488a9c2747d618ffc39270836058a44428e")
}

func (BurnMintWithLockReleaseFlagTokenPoolFastTransferOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x251a578f8b9ffb7cb748680c76e7895fe65b3fcf44c2bd03e4f163dc38b2ed0b")
}

func (BurnMintWithLockReleaseFlagTokenPoolFinalityConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xfc57ba710bb816f16b18f38549598ef4552c4acd0e0ddb7b40bc5906d25bdc1d")
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
	GetAccumulatedFees(opts *bind.CallOpts) (*big.Int, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

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

	ApplyFinalityConfigUpdates(opts *bind.TransactOpts, finalityThreshold uint16, fastTransferFeeBps uint16, maxAmountPerRequest *big.Int, rateLimitConfigArgs []TokenPoolFastFinalityRateLimitConfigArgs) (*types.Transaction, error)

	ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error)

	SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error)

	SetFastFinalityRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolFastFinalityRateLimitConfigArgs) (*types.Transaction, error)

	SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error)

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

	FilterFastTransferInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolFastTransferInboundRateLimitConsumedIterator, error)

	WatchFastTransferInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolFastTransferInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseFastTransferInboundRateLimitConsumed(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolFastTransferInboundRateLimitConsumed, error)

	FilterFastTransferOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolFastTransferOutboundRateLimitConsumedIterator, error)

	WatchFastTransferOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolFastTransferOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseFastTransferOutboundRateLimitConsumed(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolFastTransferOutboundRateLimitConsumed, error)

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
