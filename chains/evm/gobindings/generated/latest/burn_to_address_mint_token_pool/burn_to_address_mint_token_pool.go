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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIBurnMintERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"burnAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RATE_LIMITER_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyFinalityConfigUpdates\",\"inputs\":[{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"destToUseDefaultFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAccumulatedFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBurnAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"burnAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredInboundCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredOutboundCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRateLimitAdminRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRateLimitAdminRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_burnAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"lockOrBurnOutV1\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRateLimitAdminRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustomFinalityRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRouter\",\"inputs\":[{\"name\":\"newRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityTransferInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalityConfigUpdated\",\"inputs\":[{\"name\":\"finalityConfig\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeeWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminRoleGranted\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminRoleRevoked\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RouterUpdated\",\"inputs\":[{\"name\":\"oldRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidFinality\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidFinalityConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenTransferFeeConfigNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x61012080604052346103e6576166cd803803809161001d8285610465565b8339810160c0828203126103e65781516001600160a01b03811692908390036103e65761004c60208201610488565b60408201516001600160401b0381116103e65782019280601f850112156103e6578351936001600160401b03851161044f578460051b9060208201956100956040519788610465565b86526020808701928201019283116103e657602001905b828210610437575050506100c260608301610496565b936100db60a06100d460808601610496565b9401610496565b94331561042657600180546001600160a01b0319163317905581158015610415575b8015610404575b6103f3578160209160049360805260c0526040519283809263313ce56760e01b82525afa600091816103b2575b50610387575b5060a052600580546001600160a01b0319166001600160a01b0392909216919091179055610164336104aa565b50805115158060e05261026a575b5061010052604051615fda90816106d382396080518181816114d501528181611731015281816118cc01528181611eb7015281816120fc0152818161223b01528181612a8e01528181612ca801528181612e1201528181612f7f0152818161312801528181613824015281816138710152818161395d015261428e015260a0518181816137a801528181614f6001528181614fe3015261542e015260c051818181610c4c0152818161156301528181611f4501528181612b1c015261300e015260e051818181610c07015281816115a601528181611f880152612852015261010051818181611756015281816121220152613dd90152f35b60206040516102798282610465565b60008152600036813760e051156103765760005b81518110156102f4576001906001600160a01b036102ab8285610538565b5116846102b78261057a565b6102c4575b50500161028d565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a138846102bc565b505060005b825181101561036d576001906001600160a01b036103178286610538565b51168015610367578361032982610678565b610337575b50505b016102f9565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a1388361032e565b50610331565b50505038610172565b6335f4a7b360e01b60005260046000fd5b60ff1660ff821681810361039b5750610137565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d6020116103eb575b816103ce60209383610465565b810103126103e6576103df90610488565b9038610131565b600080fd5b3d91506103c1565b630a64406560e11b60005260046000fd5b506001600160a01b03811615610104565b506001600160a01b038416156100fd565b639b15e16f60e01b60005260046000fd5b6020809161044484610496565b8152019101906100ac565b634e487b7160e01b600052604160045260246000fd5b601f909101601f19168101906001600160401b0382119082101761044f57604052565b519060ff821682036103e657565b51906001600160a01b03821682036103e657565b6001600160a01b03811660009081526000805160206166ad833981519152602052604090205460ff16610532576001600160a01b031660008181526000805160206166ad83398151915260205260408120805460ff191660011790553391907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b50600090565b805182101561054c5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b805482101561054c5760005260206000200190600090565b600081815260046020526040902054801561067157600019810181811161065b5760035460001981019190821161065b5781810361060a575b50505060035480156105f457600019016105ce816003610562565b8154906000199060031b1b19169055600355600052600460205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61064361061b61062c936003610562565b90549060031b1c9283926003610562565b819391549060031b91821b91600019901b19161790565b905560005260046020526040600020553880806105b3565b634e487b7160e01b600052601160045260246000fd5b5050600090565b80600052600460205260406000205415600014610532576003546801000000000000000081101561044f576106b961062c8260018594016003556003610562565b905560035490600052600460205260406000205560019056fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a7146139de575080630bd7c46d14613990578063164e68de146138f6578063181f5a771461389557806321df0da714613851578063240028e8146137fa578063248a9ca3146137cc57806324f65ee71461378e5780632a10097b146135415780632c286daf1461343f5780632f2ff15d146133fe57806336568abe146133a057806337b192471461329957806338b39d2914610f015780633907753714612ef3578063489a68f2146129f95780634c5ef0ed146129b25780634f71592c1461297d57806354c8a4f31461281e5780635df45a37146127fb57806362ddd3c414612777578063791e5a101461273c57806379ba50971461268f578063804ba5a91461264c5780638926f54f146126065780638da5cb5b146125df57806391d1485414612592578063962d4020146123d85780639a4575b914611e635780639f68f67314611e2b578063a217fddf14611e0f578063a42a7b8b14611cc6578063a7cd63b714611c45578063acfecf9114611b21578063af58d59f14611ad8578063b0f479a114611ab1578063b1c71c6514611457578063b79019b5146111b4578063b79465801461117b578063c0d78655146110bb578063c4bffe2b14610fae578063c75eea9c14610f06578063c8de9fe014610f01578063cf7401f314610d1c578063d547741f14610cd2578063da90a9f314610c70578063dc0bd97114610c2c578063e0351e1314610bef578063e58d80c714610b84578063e8a1da17146103095763f2fde38b1461025257600080fd5b34610306576020600319360112610306576001600160a01b03610273613b75565b61027b614650565b163381146102de57807fffffffffffffffffffffffff00000000000000000000000000000000000000008354161782556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b80fd5b50346103065761031836613f00565b93919092610324614650565b82915b8083106109ef575050508063ffffffff4216917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee1843603015b858210156109eb578160051b850135818112156109e757850190610120823603126109e7576040519561039287613c38565b823567ffffffffffffffff811681036109e2578752602083013567ffffffffffffffff81116109de5783019536601f880112156109de578635966103d5886141e6565b976103e3604051998a613c70565b8089526020808a019160051b830101903682116109da5760208301905b8282106109a7575050505060208801968752604084013567ffffffffffffffff81116109a3576104339036908601613e34565b9860408901998a5261045d61044b3660608801614014565b9560608b0196875260c0369101614014565b9660808a0197885261046f86516155a4565b61047988516155a4565b8a51511561097b5761049567ffffffffffffffff8b5116615e52565b156109445767ffffffffffffffff8a511681526008602052604081206105d587516fffffffffffffffffffffffffffffffff604082015116906105906fffffffffffffffffffffffffffffffff602083015116915115158360806040516104fb81613c38565b858152602081018c905260408101849052606081018690520152855474ff000000000000000000000000000000000000000091151560a01b919091167fffffffffffffffffffffff0000000000000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff84161773ffffffff0000000000000000000000000000000060808b901b1617178555565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176001830155565b6106fb89516fffffffffffffffffffffffffffffffff604082015116906106b66fffffffffffffffffffffffffffffffff6020830151169151151583608060405161061f81613c38565b858152602081018c9052604081018490526060810186905201526002860180547fffffffffffffffffffffff000000000000000000000000000000000000000000166fffffffffffffffffffffffffffffffff85161773ffffffff0000000000000000000000000000000060808c901b161791151560a01b74ff000000000000000000000000000000000000000016919091179055565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176003830155565b60048c5191019080519067ffffffffffffffff82116109175761071e835461433a565b601f81116108dc575b50602090601f831160011461087957610757929185918361086e575b50506000198260011b9260031b1c19161790565b90555b805b89518051821015610792579061078c600192610785838f67ffffffffffffffff90511692614326565b51906150e7565b0161075c565b5050975097987f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c29295939661086067ffffffffffffffff600197949c511692519351915161082c6107f760405196879687526101006020880152610100870190613caf565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a1019093949291610360565b015190503880610743565b8385528185209190601f198416865b8181106108c457509084600195949392106108ab575b505050811b01905561075a565b015160001960f88460031b161c1916905538808061089e565b92936020600181928786015181550195019301610888565b6109079084865260208620601f850160051c8101916020861061090d575b601f0160051c0190614617565b38610727565b90915081906108fa565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60249067ffffffffffffffff8b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b807f14c880ca0000000000000000000000000000000000000000000000000000000060049252fd5b8680fd5b813567ffffffffffffffff81116109d6576020916109cb8392833691890101613e34565b815201910190610400565b8a80fd5b8880fd5b8580fd5b600080fd5b8380fd5b8280f35b9092919367ffffffffffffffff610a0f610a0a8785886140de565b61409a565b1695610a1a87615c89565b15610b58578684526008602052610a3660056040862001615b2d565b94845b8651811015610a6f576001908987526008602052610a6860056040892001610a61838b614326565b5190615d3c565b5001610a39565b5093945094909580855260086020526005604086208681558660018201558660028201558660038201558660048201610aa8815461433a565b80610b17575b5050500180549086815581610af9575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a1019190949394610327565b865260208620908101905b81811015610abe57868155600101610b04565b601f8111600114610b2d5750555b863880610aae565b81835260208320610b4891601f01861c810190600101614617565b8082528160208120915555610b25565b602484887f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b5034610306576020600319360112610306576001600160a01b036040610ba8613b75565b927f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af81526002602052209116600052602052602060ff604060002054166040519015158152f35b503461030657806003193601126103065760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b503461030657806003193601126103065760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b5034610306576020600319360112610306577fd63806009f622849e3b7cfd82d762420d57574c39f945f678871b2b5f1e8ce4b6020610cad613b75565b610cb5614650565b610cbe81614d55565b506001600160a01b0360405191168152a180f35b503461030657604060031936011261030657610d18600435610cf2613b8b565b90610d13610d0e82600052600260205260016040600020015490565b614cf5565b614e3f565b5080f35b50346103065760e060031936011261030657610d36613d74565b9060607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc36011261030657604051610d6d81613c54565b6024358015158103610efd5781526044356fffffffffffffffffffffffffffffffff81168103610efd5760208201526064356fffffffffffffffffffffffffffffffff81168103610efd57604082015260607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c360112610ef95760405190610df482613c54565b60843580151581036109e757825260a4356fffffffffffffffffffffffffffffffff811681036109e757602083015260c4356fffffffffffffffffffffffffffffffff811681036109e75760408301527f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af83526002602052604083206001600160a01b03331660005260205260ff604060002054161580610ecd575b610ea157610e9e92936152f7565b80f35b6024837f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b508280526002602052604083206001600160a01b03331660005260205260ff6040600020541615610e90565b5080fd5b8280fd5b613db9565b503461030657602060031936011261030657610f51610f4c6040610faa9367ffffffffffffffff610f35613d74565b610f3d61446e565b50168152600860205220614499565b61546f565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b0390f35b5034610306578060031936011261030657604051906006548083528260208101600684526020842092845b8181106110a2575050610fee92500383613c70565b8151611012610ffc826141e6565b9161100a6040519384613c70565b8083526141e6565b91601f19602083019301368437805b8451811015611053578067ffffffffffffffff61104060019388614326565b511661104c8286614326565b5201611021565b50925090604051928392602084019060208552518091526040840192915b81811061107f575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101611071565b8454835260019485019487945060209093019201610fd9565b5034610306576020600319360112610306576001600160a01b036110dd613b75565b6110e5614650565b1680156111535760407f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f16849160055490807fffffffffffffffffffffffff00000000000000000000000000000000000000008316176005556001600160a01b038351921682526020820152a180f35b6004827f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b503461030657602060031936011261030657610faa6111a061119b613d74565b61462e565b604051918291602083526020830190613caf565b50346103065760206003193601126103065760043567ffffffffffffffff8111610ef9576111e6903690600401613cf0565b6111ee614650565b825b8181106111fb578380f35b611209610a0a8284866144ff565b6112216112178385876144ff565b602081019061453f565b907fb0897119e8510f887b892cbc4c8506fc51d9849fd90afae4fd065e705f2d0f6c61125b61125186888a6144ff565b604081019061453f565b91909261127161126c3687846141fe565b615501565b61127f61126c3685876141fe565b6040519461128c86613c1c565b6112973682846141fe565b86526112e167ffffffffffffffff6112b03687896141fe565b9860208901998a52169586956112d3604051958695604087526040870191614593565b918483036020860152614593565b0390a28652600d60205260408620905180519067ffffffffffffffff821161142a5768010000000000000000821161142a576020908354838555808410611410575b500182885260208820885b8381106113f3575050505060010190519081519167ffffffffffffffff83116113c6576801000000000000000083116113c65760209082548484558085106113ac575b500190865260208620865b83811061138f57505050506001016111f0565b60019060206001600160a01b03855116940193818401550161137c565b8389528289206113c0918101908601614617565b38611371565b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60019060206001600160a01b03855116940193818401550161132e565b848a52828a20611424918101908501614617565b38611323565b6024887f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b50346103065760606003193601126103065760043567ffffffffffffffff8111610ef95760a06003198236030112610ef957611491613d21565b9060443567ffffffffffffffff81116109e7576114b2903690600401613e34565b506114bb61430d565b5060848101916114ca8361413f565b6001600160a01b03807f000000000000000000000000000000000000000000000000000000000000000016911603611a7457602482019177ffffffffffffffff000000000000000000000000000000006115238461409a565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611998578691611a45575b50611a1d576115a46044820161413f565b7f00000000000000000000000000000000000000000000000000000000000000006119ce575b5067ffffffffffffffff6115dd8461409a565b166115f5816000526007602052604060002054151590565b156119a35760206001600160a01b0360055416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa801561199857869061194f575b6001600160a01b03915016330361192357606461ffff91013591169384151593848095611914575b156118775761ffff600a5416808710611847575061180b95506116c36116b36116998661409a565b67ffffffffffffffff16600052600b602052604060002090565b846116bd8461413f565b9161591e565b7f7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b267ffffffffffffffff6116ff6116f98761409a565b9361413f565b604080516001600160a01b03929092168252602082018790529190931692a25b508092611815575b5061119b6117da917f000000000000000000000000000000000000000000000000000000000000000061177b857f000000000000000000000000000000000000000000000000000000000000000083614838565b7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff6117ae8461409a565b604080516001600160a01b0390951685523360208601528401889052169180606081015b0390a261409a565b906117e3615427565b604051926117f084613c1c565b83526020830152604051928392604084526040840190613fc0565b9060208301520390f35b6117da91925061183f61119b9161271061183861ffff600a5460101c16836145d5565b0490615462565b929150611727565b82604491887fe08f03ef000000000000000000000000000000000000000000000000000000008352600452602452fd5b5061180b945067ffffffffffffffff61188f8461409a565b1680825260086020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894483806118f4604086206001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001692839161591e565b604080516001600160a01b039290921682526020820192909252a261171f565b5061ffff600a54161515611671565b6024857f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d602011611990575b8161196960209383613c70565b810103126109de57516001600160a01b03811681036109de576001600160a01b0390611649565b3d915061195c565b6040513d88823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008652600452602485fd5b6001600160a01b03166119ee816000526004602052604060002054151590565b6115ca577fd0d25976000000000000000000000000000000000000000000000000000000008652600452602485fd5b6004857f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b611a67915060203d602011611a6d575b611a5f8183613c70565b8101906150cf565b38611593565b503d611a55565b6024846001600160a01b03611a888661413f565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835216600452fd5b503461030657806003193601126103065760206001600160a01b0360055416604051908152f35b503461030657602060031936011261030657610f51610f4c60026040610faa9467ffffffffffffffff611b09613d74565b611b1161446e565b5016815260086020522001614499565b50346103065767ffffffffffffffff611b3936613f4e565b929091611b44614650565b1691611b5d836000526007602052604060002054151590565b15611c19578284526008602052611b8c60056040862001611b7f368486613dfd565b6020815191012090615d3c565b15611bd157907f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d7691611bcb60405192839260208452602084019161444d565b0390a280f35b82611c15836040519384937f74f23c7c000000000000000000000000000000000000000000000000000000008552600485015260406024850152604484019161444d565b0390fd5b602484847f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b503461030657806003193601126103065760405160038054808352908352909160208301917fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b915b818110611cb057610faa85611ca481870382613c70565b60405191829182613ebd565b8254845260209093019260019283019201611c8d565b50346103065760206003193601126103065767ffffffffffffffff611ce9613d74565b1681526008602052611d0060056040832001615b2d565b8051601f19611d27611d11836141e6565b92611d1f6040519485613c70565b8084526141e6565b01835b818110611dfe575050825b8251811015611d7b5780611d4b60019285614326565b5185526009602052611d5f6040862061438d565b611d698285614326565b52611d748184614326565b5001611d35565b81846040519182916020830160208452825180915260408401602060408360051b870101940192905b828210611db357505050500390f35b91936020611dee827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851613caf565b9601920192018594939192611da4565b806060602080938601015201611d2a565b5034610306578060031936011261030657602090604051908152f35b503461030657611ca460016040610faa9367ffffffffffffffff611e4e36613e52565b505050509050168152600d6020522001614190565b5034610306576020600319360112610306576004359067ffffffffffffffff82116103065760a0600319833603011261030657611e9e61430d565b5060848201611eac8161413f565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116036123c457602483019277ffffffffffffffff00000000000000000000000000000000611f058561409a565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156122f85784916123a5575b5061237d57611f866044820161413f565b7f000000000000000000000000000000000000000000000000000000000000000061232e575b5067ffffffffffffffff611fbf8561409a565b16611fd7816000526007602052604060002054151590565b156123035760206001600160a01b0360055416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa80156122f85784906122af575b6001600160a01b03915016330361228357606401359180156121dd5761ffff600a5416806121ad5750508261217d9261119b9261207a612070611699610faa9861409a565b836116bd8461413f565b7f7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b267ffffffffffffffff6120b06116f98661409a565b604080516001600160a01b03929092168252602082018690529190931692a25b7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000092612147817f000000000000000000000000000000000000000000000000000000000000000086614838565b6117d26121538661409a565b604080516001600160a01b0390971687523360208801528601929092529116929081906060820190565b612185615427565b6040519161219283613c1c565b82526020820152604051918291602083526020830190613fc0565b604492507fe08f03ef00000000000000000000000000000000000000000000000000000000825281600452602452fd5b505061119b8261217d9267ffffffffffffffff6121fc610faa9661409a565b168060005260086020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944828061226360406000206001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001692839161591e565b604080516001600160a01b039290921682526020820192909252a26120d0565b6024837f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d6020116122f0575b816122c960209383613c70565b810103126109e757516001600160a01b03811681036109e7576001600160a01b039061202b565b3d91506122bc565b6040513d86823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008452600452602483fd5b6001600160a01b031661234e816000526004602052604060002054151590565b611fac577fd0d25976000000000000000000000000000000000000000000000000000000008452600452602483fd5b6004837f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b6123be915060203d602011611a6d57611a5f8183613c70565b38611f75565b906001600160a01b03611a8860249361413f565b50346103065760606003193601126103065760043567ffffffffffffffff8111610ef95761240a903690600401613cf0565b60243567ffffffffffffffff81116109e75761242a903690600401613f8f565b60449291923567ffffffffffffffff81116109de5761244d903690600401613f8f565b9190927f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af87526002602052604087206001600160a01b03331660005260205260ff604060002054161580612566575b61253a57818114801590612530575b61250857865b8181106124bc578780f35b806125026124d0610a0a600194868c6140de565b6124db83878b6142fd565b6124fc6124f46124ec868b8d6142fd565b923690614014565b913690614014565b916152f7565b016124b1565b6004877f568efce2000000000000000000000000000000000000000000000000000000008152fd5b50828114156124ab565b6024877f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b508680526002602052604087206001600160a01b03331660005260205260ff604060002054161561249c565b5034610306576040600319360112610306576001600160a01b0360406125b6613b8b565b9260043581526002602052209116600052602052602060ff604060002054166040519015158152f35b503461030657806003193601126103065760206001600160a01b0360015416604051908152f35b503461030657602060031936011261030657602061264267ffffffffffffffff61262e613d74565b166000526007602052604060002054151590565b6040519015158152f35b50346103065760206003193601126103065760043567ffffffffffffffff8111610ef957612681610e9e913690600401613d43565b9061268a614650565b6149cf565b503461030657806003193601126103065780546001600160a01b0381163303612714577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551682556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b503461030657806003193601126103065760206040517f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af8152f35b50346103065761278636613f4e565b61279293929193614650565b67ffffffffffffffff82166127b4816000526007602052604060002054151590565b156127d05750610e9e92936127ca913691613dfd565b906150e7565b7f1e670e4b000000000000000000000000000000000000000000000000000000008452600452602483fd5b50346103065780600319360112610306576020612816614252565b604051908152f35b5034610306576128489061285061283436613f00565b9591612841939193614650565b36916141fe565b9336916141fe565b7f00000000000000000000000000000000000000000000000000000000000000001561295557815b83518110156128de57806001600160a01b0361289660019387614326565b51166128a181615b90565b6128ad575b5001612878565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a1386128a6565b5090805b8251811015610d1857806001600160a01b0361290060019386614326565b5116801561294f5761291181615df8565b61291e575b505b016128e2565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a184612916565b50612918565b6004827f35f4a7b3000000000000000000000000000000000000000000000000000000008152fd5b503461030657611ca46040610faa9267ffffffffffffffff61299e36613e52565b505050509050168152600d60205220614190565b5034610306576040600319360112610306576129cc613d74565b906024359067ffffffffffffffff8211610306576020612642846129f33660048701613e34565b90614153565b50346103065760406003193601126103065760043567ffffffffffffffff8111610ef9578060040191610100600319833603011261030657612a39613d21565b9181604051612a4781613c00565b52612a74612a6a612a65612a5e60c48501886140ee565b3691613dfd565b614eec565b6064830135614fe0565b926084820190612a838261413f565b6001600160a01b03807f000000000000000000000000000000000000000000000000000000000000000016911603612edf57602483019577ffffffffffffffff00000000000000000000000000000000612adc8861409a565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611998578691612ec0575b50611a1d5767ffffffffffffffff612b638861409a565b16612b7b816000526007602052604060002054151590565b156119a35760206001600160a01b0360055416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115611998578691612ea1575b501561192357612be58761409a565b90612bfb60a48601926129f3612a5e85856140ee565b15612e5a5750604493929161ffff16159050612dbf5767ffffffffffffffff612c238761409a565b168452600c602052612c3c60408520866116bd8461413f565b7f41a8aa8df7945f0fb8ac5f7d88279638d9dc2ef9a6bf4ec9a53b80681b34aff767ffffffffffffffff612c726116f98961409a565b604080516001600160a01b03929092168252602082018990529190931692a25b0192612c9d8461413f565b916001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001692833b15610ef9576040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b0391909116600482015260248101859052818160448183885af18015612db457612d9f575b5050608067ffffffffffffffff6020956001600160a01b03612d6d612d677ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09661409a565b9261413f565b60405196875233898801521660408601528560608601521692a260405190612d9482613c00565b815260405190518152f35b612daa828092613c70565b6103065780612d22565b6040513d84823e3d90fd5b5067ffffffffffffffff612dd28661409a565b1680845260086020527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c8580612e3a600260408920016001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001692839161591e565b604080516001600160a01b039290921682526020820192909252a2612c92565b612e6492506140ee565b611c156040519283927f24eb47e500000000000000000000000000000000000000000000000000000000845260206004850152602484019161444d565b612eba915060203d602011611a6d57611a5f8183613c70565b38612bd6565b612ed9915060203d602011611a6d57611a5f8183613c70565b38612b4c565b6024846001600160a01b03611a888561413f565b50346103065760206003193601126103065760043567ffffffffffffffff8111610ef957806004016101006003198336030112610efd5782604051612f3781613c00565b5282604051612f4581613c00565b52612f66612f5c612a65612a5e60c48601856140ee565b6064840135614fe0565b9160848101612f748161413f565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116036132855750602481019177ffffffffffffffff00000000000000000000000000000000612fce8461409a565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611998578691613266575b50611a1d5767ffffffffffffffff6130558461409a565b1661306d816000526007602052604060002054151590565b156119a35760206001600160a01b0360055416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115611998578691613247575b5015611923576130d78361409a565b906130ed60a48401926129f3612a5e85856140ee565b15612e5a575050906044839267ffffffffffffffff61310b8461409a565b168087526008602052613150600260408920016001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001696879161591e565b604080516001600160a01b0387168152602081018890527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a201906131968261413f565b85843b15610306576040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b03929092166004830152602482018690528160448183885af1801561199857926001600160a01b03612d6d612d6760809560209a67ffffffffffffffff967ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc099613237575b505061409a565b8161324191613c70565b38613230565b613260915060203d602011611a6d57611a5f8183613c70565b386130c8565b61327f915060203d602011611a6d57611a5f8183613c70565b3861303e565b846001600160a01b03611a8860249361413f565b50346103065760a0600319360112610306576132b3613b75565b5060243567ffffffffffffffff8116809103610ef95760443567ffffffffffffffff8111610efd5760031960a09136030112610ef9576132f1613d32565b5060843567ffffffffffffffff8111610efd57916040916133186080943690600401613d8b565b5050816060845161332881613bb5565b828152826020820152828682015201528152600e602052206040519061334d82613bb5565b5463ffffffff808216928381528160208201818560201c16815260ff60606040850194848860401c168652019560601c161515855260405195865251166020850152511660408301525115156060820152f35b5034610306576040600319360112610306576133ba613b8b565b336001600160a01b038216036133d657610d1890600435614e3f565b6004827f6697b232000000000000000000000000000000000000000000000000000000008152fd5b503461030657604060031936011261030657610d1860043561341e613b8b565b9061343a610d0e82600052600260205260016040600020015490565b614780565b50346103065760606003193601126103065760043561ffff8116809103610ef957613468613d21565b60443567ffffffffffffffff81116109e757613488903690600401613d43565b90613491614650565b61ffff83169261271084101561351557849260409492613507927f52aa194b292c8bfb5aaca8ee2000a965c3a051b306ff841873b16147526a39ba977fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000063ffff0000600a549360101b1692161717600a556149cf565b82519182526020820152a180f35b602486857f95f3517a000000000000000000000000000000000000000000000000000000008252600452fd5b50346103065760406003193601126103065760043567ffffffffffffffff8111610ef95736602382011215610ef957806004013567ffffffffffffffff8111610efd5760248201916024369160a08402010111610efd5760243567ffffffffffffffff81116109e7576135b8903690600401613cf0565b9190926135c3614650565b84805b83811061363057509150505b8181106135dd578380f35b8067ffffffffffffffff6135f7610a0a60019486886140de565b16808652600e6020528560408120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee88680a2016135d2565b6001917f56f77aeff2def50c8b5f5a0df3bab7183df09bf36c6feba496bb42551db77d7060808561377e61374463ffffffff613773613737826137686136878f806136818f9283610a0a918e61405b565b9a61405b565b604067ffffffffffffffff602083019a169c8d8152600e60205220836136ac8b6140af565b169181549060408101937fffffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffff67ffffffff000000006136e9876140af565b60201b16918f6cff0000000000000000000000007fffffffffffffffffffffffffffffffffffffffff000000000000000000000000916bffffffff0000000000000000606088019d8e6140af565b60401b1696019e8f6140c0565b151560601b16951617161717179055826137606040519a6140cd565b1689526140cd565b1660208701526140cd565b166040840152613fea565b15156060820152a20185906135c6565b5034610306578060031936011261030657602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b5034610306576020600319360112610306576020612816600435600052600260205260016040600020015490565b503461030657602060031936011261030657602090613817613b75565b90506001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b503461030657806003193601126103065760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461030657806003193601126103065750610faa6040516138b8604082613c70565b602081527f4275726e546f41646472657373546f6b656e506f6f6c20312e362e332d6465766020820152604051918291602083526020830190613caf565b503461030657602060031936011261030657613910613b75565b613918614650565b613920614252565b908161392a578280f35b60206001600160a01b0382613981857f738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b859599957f0000000000000000000000000000000000000000000000000000000000000000614838565b6040519485521692a238808280f35b5034610306576020600319360112610306577ff7af318a70f367e30346e2704f6ef646b378a7dcb49767beb98a1774cd11e38960206139cd613b75565b6139d5614650565b610cbe8161468e565b905034610ef9576020600319360112610ef9576004357fffffffff000000000000000000000000000000000000000000000000000000008116809103610efd57602092507ff208a58f000000000000000000000000000000000000000000000000000000008114908115613b4b575b8115613b21575b8115613af7575b8115613acd575b8115613a70575b5015158152f35b7f7965db0b00000000000000000000000000000000000000000000000000000000811491508115613aa3575b5038613a69565b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438613a9c565b7f01ffc9a70000000000000000000000000000000000000000000000000000000081149150613a62565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150613a5b565b7f1ef5498f0000000000000000000000000000000000000000000000000000000081149150613a54565b7faff2afbf0000000000000000000000000000000000000000000000000000000081149150613a4d565b600435906001600160a01b03821682036109e257565b602435906001600160a01b03821682036109e257565b35906001600160a01b03821682036109e257565b6080810190811067ffffffffffffffff821117613bd157604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff821117613bd157604052565b6040810190811067ffffffffffffffff821117613bd157604052565b60a0810190811067ffffffffffffffff821117613bd157604052565b6060810190811067ffffffffffffffff821117613bd157604052565b90601f601f19910116810190811067ffffffffffffffff821117613bd157604052565b67ffffffffffffffff8111613bd157601f01601f191660200190565b919082519283825260005b848110613cdb575050601f19601f8460006020809697860101520116010190565b80602080928401015182828601015201613cba565b9181601f840112156109e25782359167ffffffffffffffff83116109e2576020808501948460051b0101116109e257565b6024359061ffff821682036109e257565b6064359061ffff821682036109e257565b9181601f840112156109e25782359167ffffffffffffffff83116109e25760208085019460e085020101116109e257565b6004359067ffffffffffffffff821682036109e257565b9181601f840112156109e25782359167ffffffffffffffff83116109e257602083818601950101116109e257565b346109e25760006003193601126109e25760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b929192613e0982613c93565b91613e176040519384613c70565b8294818452818301116109e2578281602093846000960137010152565b9080601f830112156109e257816020613e4f93359101613dfd565b90565b60a06003198201126109e2576004356001600160a01b03811681036109e2579160243567ffffffffffffffff811681036109e257916044359160643561ffff811681036109e257916084359067ffffffffffffffff82116109e257613eb991600401613d8b565b9091565b602060408183019282815284518094520192019060005b818110613ee15750505090565b82516001600160a01b0316845260209384019390920191600101613ed4565b60406003198201126109e25760043567ffffffffffffffff81116109e25781613f2b91600401613cf0565b929092916024359067ffffffffffffffff82116109e257613eb991600401613cf0565b9060406003198301126109e25760043567ffffffffffffffff811681036109e257916024359067ffffffffffffffff82116109e257613eb991600401613d8b565b9181601f840112156109e25782359167ffffffffffffffff83116109e257602080850194606085020101116109e257565b613e4f916020613fd98351604084526040840190613caf565b920151906020818403910152613caf565b359081151582036109e257565b35906fffffffffffffffffffffffffffffffff821682036109e257565b91908260609103126109e25760405161402c81613c54565b604061405681839561403d81613fea565b855261404b60208201613ff7565b602086015201613ff7565b910152565b919081101561406b5760a0020190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3567ffffffffffffffff811681036109e25790565b3563ffffffff811681036109e25790565b3580151581036109e25790565b359063ffffffff821682036109e257565b919081101561406b5760051b0190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156109e2570180359067ffffffffffffffff82116109e2576020019181360383136109e257565b356001600160a01b03811681036109e25790565b9067ffffffffffffffff613e4f92166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b906040519182815491828252602082019060005260206000209260005b8181106141c45750506141c292500383613c70565b565b84546001600160a01b03168352600194850194879450602090930192016141ad565b67ffffffffffffffff8111613bd15760051b60200190565b92919061420a816141e6565b936142186040519586613c70565b602085838152019160051b81019283116109e257905b82821061423a57505050565b6020809161424784613ba1565b81520191019061422e565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156142f1576000916142c2575090565b90506020813d6020116142e9575b816142dd60209383613c70565b810103126109e2575190565b3d91506142d0565b6040513d6000823e3d90fd5b919081101561406b576060020190565b6040519061431a82613c1c565b60606020838281520152565b805182101561406b5760209160051b010190565b90600182811c92168015614383575b602083101461435457565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691614349565b90604051918260008254926143a18461433a565b808452936001811690811561440d57506001146143c6575b506141c292500383613c70565b90506000929192526020600020906000915b8183106143f15750509060206141c292820101386143b9565b60209193508060019154838589010152019101909184926143d8565b602093506141c29592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b820101386143b9565b601f8260209493601f19938186528686013760008582860101520116010190565b6040519061447b82613c38565b60006080838281528260208201528260408201528260608201520152565b906040516144a681613c38565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b919081101561406b5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa1813603018212156109e2570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156109e2570180359067ffffffffffffffff82116109e257602001918160051b360383136109e257565b9160209082815201919060005b8181106145ad5750505090565b9091926020806001926001600160a01b036145c788613ba1565b1681520194019291016145a0565b818102929181159184041417156145e857565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818110614622575050565b60008155600101614617565b67ffffffffffffffff166000526008602052613e4f600460406000200161438d565b6001600160a01b0360015416330361466457565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b6001600160a01b03811660009081527fa0bc845bafb3af92358952eeb50d6b57502488fa7f04f1920acc69cb8f116b20602052604090205460ff1661477a576001600160a01b031660008181527fa0bc845bafb3af92358952eeb50d6b57502488fa7f04f1920acc69cb8f116b206020526040812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790553391907f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b50600090565b80600052600260205260406000206001600160a01b03831660005260205260ff60406000205416156000146148315780600052600260205260406000206001600160a01b038316600052602052604060002060017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008254161790556001600160a01b03339216907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d600080a4600190565b5050600090565b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000602082019081526001600160a01b039384166024830152604480830195909552938152614901929091614890606484613c70565b166000806040958651946148a48887613c70565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020870152519082855af13d156149aa573d916148e583613c93565b926148f287519485613c70565b83523d6000602085013e615f01565b8051908161490e57505050565b60208061491f9383010191016150cf565b156149275750565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606091615f01565b356fffffffffffffffffffffffffffffffff811681036109e25790565b9160005b82811015614cef5760e08102840160006149ec8261409a565b9067ffffffffffffffff821691614a10836000526007602052604060002054151590565b15614cc357614ad99260408593614a84614a7e94614a7e614a44602060019c9b0192611699614a3f3686614014565b6155a4565b91825463ffffffff8160801c16159081614ca5575b81614c96575b81614c7b575b81614c6c575b5080614c5d575b614bd2575b3690614014565b906156eb565b6080850192614a96614a3f3686614014565b8152600c6020522092835463ffffffff8160801c16159081614bb4575b81614ba5575b81614b8a575b81614b7b575b5080614b6c575b614adf575b503690614014565b016149d3565b614afc60a06fffffffffffffffffffffffffffffffff92016149b2565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff161717835538614ad1565b50614b76826140c0565b614acc565b60ff915060a01c161538614ac5565b6fffffffffffffffffffffffffffffffff8116159150614abf565b8589015460801c159150614ab9565b858901546fffffffffffffffffffffffffffffffff16159150614ab3565b6fffffffffffffffffffffffffffffffff614bee878b016149b2565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff1617178355614a77565b50614c67816140c0565b614a72565b60ff915060a01c161538614a6b565b6fffffffffffffffffffffffffffffffff8116159150614a65565b848e015460801c159150614a5f565b848e01546fffffffffffffffffffffffffffffffff16159150614a59565b506024917f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50915050565b80600052600260205260406000206001600160a01b03331660005260205260ff6040600020541615614d245750565b7fe2517d3f000000000000000000000000000000000000000000000000000000006000523360045260245260446000fd5b6001600160a01b03811660009081527fa0bc845bafb3af92358952eeb50d6b57502488fa7f04f1920acc69cb8f116b20602052604090205460ff161561477a576001600160a01b031660008181527fa0bc845bafb3af92358952eeb50d6b57502488fa7f04f1920acc69cb8f116b206020526040812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690553391907f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a4600190565b80600052600260205260406000206001600160a01b03831660005260205260ff604060002054166000146148315780600052600260205260406000206001600160a01b03831660005260205260406000207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541690556001600160a01b03339216907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b600080a4600190565b80518015614f5c57602003614f1e5780516020828101918301839003126109e257519060ff8211614f1e575060ff1690565b611c15906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190613caf565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff82116145e857565b60ff16604d81116145e857600a0a90565b8115614fb1570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff8116928284146150c85782841161509e579061502591614f82565b91604d60ff8416118015615083575b61504d57505090615047613e4f92614f96565b906145d5565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b5061508d83614f96565b8015614fb157600019048411615034565b6150a791614f82565b91604d60ff84161161504d575050906150c2613e4f92614f96565b90614fa7565b5050505090565b908160209103126109e2575180151581036109e25790565b908051156152cd5767ffffffffffffffff8151602083012092169182600052600860205261511c816005604060002001615eac565b156152895760005260096020526040600020815167ffffffffffffffff8111613bd157615149825461433a565b601f8111615257575b506020601f82116001146151cd57916151a7827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea95936151bd956000916151c2575b506000198260011b9260031b1c19161790565b9055604051918291602083526020830190613caf565b0390a2565b905084015138615194565b601f1982169083600052806000209160005b81811061523f5750926151bd9492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610615226575b5050811b0190556111a0565b85015160001960f88460031b161c19169055388061521a565b9192602060018192868a0151815501940192016151df565b61528390836000526020600020601f840160051c8101916020851061090d57601f0160051c0190614617565b38615152565b5090611c156040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190613caf565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b67ffffffffffffffff1660008181526007602052604090205490929190156153f957916153f660e0926153c28561534e7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b976155a4565b8460005260086020526153658160406000206156eb565b61536e836155a4565b8460005260086020526153888360026040600020016156eb565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016602082015260208152613e4f604082613c70565b919082039182116145e857565b61547761446e565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff80835116916154d460208501936154ce6154c163ffffffff87511642615462565b85608089015116906145d5565b906154f4565b808210156154ed57505b16825263ffffffff4216905290565b90506154de565b919082018092116145e857565b805160005b81811061551257505050565b600181018082116145e8575b82811061552e5750600101615506565b6001600160a01b036155408386614326565b51166001600160a01b036155548387614326565b5116146155635760010161551e565b6001600160a01b036155758386614326565b51167fa1726e400000000000000000000000000000000000000000000000000000000060005260045260246000fd5b805115615644576fffffffffffffffffffffffffffffffff6040820151166fffffffffffffffffffffffffffffffff602083015116106155e15750565b606490615642604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff604082015116158015906156cc575b61566b5750565b606490615642604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff6020820151161515615664565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1991615824606092805461572863ffffffff8260801c1642615462565b9081615863575b50506fffffffffffffffffffffffffffffffff600181602086015116928281541680851060001461585b57508280855b16167fffffffffffffffffffffffffffffffff000000000000000000000000000000008254161781556157d88651151582907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b60408601517fffffffffffffffffffffffffffffffff0000000000000000000000000000000060809190911b16939092166fffffffffffffffffffffffffffffffff1692909217910155565b6153f660405180926fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b83809161575f565b6fffffffffffffffffffffffffffffffff916158988392836158916001880154948286169560801c906145d5565b91166154f4565b8082101561591757505b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff9290911692909216167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116174260801b73ffffffff0000000000000000000000000000000016178155388061572f565b90506158a2565b9182549060ff8260a01c16158015615b25575b615b1f576fffffffffffffffffffffffffffffffff8216916001850190815461597663ffffffff6fffffffffffffffffffffffffffffffff83169360801c1642615462565b9081615a81575b5050848110615a4257508383106159d75750506159ac6fffffffffffffffffffffffffffffffff928392615462565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b5460801c916159e68185615462565b926000198101908082116145e857615a09615a0e926001600160a01b03966154f4565b614fa7565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b82856001600160a01b03927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611615af557615a9c926154ce9160801c906145d5565b80841015615af05750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff000000000000000000000000000000001617865592388061597d565b615aa7565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b508215615931565b906040519182815491828252602082019060005260206000209260005b818110615b5f5750506141c292500383613c70565b8454835260019485019487945060209093019201615b4a565b805482101561406b5760005260206000200190600090565b60008181526004602052604090205480156148315760001981018181116145e8576003549060001982019182116145e857818103615c38575b5050506003548015615c095760001901615be4816003615b78565b60001982549160031b1b19169055600355600052600460205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b615c71615c49615c5a936003615b78565b90549060031b1c9283926003615b78565b81939154906000199060031b92831b921b19161790565b90556000526004602052604060002055388080615bc9565b60008181526007602052604090205480156148315760001981018181116145e8576006549060001982019182116145e857818103615d02575b5050506006548015615c095760001901615cdd816006615b78565b60001982549160031b1b19169055600655600052600760205260006040812055600190565b615d24615d13615c5a936006615b78565b90549060031b1c9283926006615b78565b90556000526007602052604060002055388080615cc2565b9060018201918160005282602052604060002054801515600014615def5760001981018181116145e85782549060001982019182116145e857818103615db8575b50505080548015615c09576000190190615d978282615b78565b60001982549160031b1b191690555560005260205260006040812055600190565b615dd8615dc8615c5a9386615b78565b90549060031b1c92839286615b78565b905560005283602052604060002055388080615d7d565b50505050600090565b8060005260046020526040600020541560001461477a5760035468010000000000000000811015613bd157615e39615c5a8260018594016003556003615b78565b9055600354906000526004602052604060002055600190565b8060005260076020526040600020541560001461477a5760065468010000000000000000811015613bd157615e93615c5a8260018594016006556006615b78565b9055600654906000526007602052604060002055600190565b60008281526001820160205260409020546148315780549068010000000000000000821015613bd15782615eea615c5a846001809601855584615b78565b905580549260005201602052604060002055600190565b91929015615f7c5750815115615f15575090565b3b15615f1e5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015615f8f5750805190602001fd5b611c15906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190613caf56fea164736f6c634300081a000aac33ff75c19e70fe83507db0d683fd3465c996598dc972688b7ace676c89077b",
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BurnToAddressMintTokenPool.Contract.DEFAULTADMINROLE(&_BurnToAddressMintTokenPool.CallOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BurnToAddressMintTokenPool.Contract.DEFAULTADMINROLE(&_BurnToAddressMintTokenPool.CallOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) RATELIMITERADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "RATE_LIMITER_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) RATELIMITERADMINROLE() ([32]byte, error) {
	return _BurnToAddressMintTokenPool.Contract.RATELIMITERADMINROLE(&_BurnToAddressMintTokenPool.CallOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) RATELIMITERADMINROLE() ([32]byte, error) {
	return _BurnToAddressMintTokenPool.Contract.RATELIMITERADMINROLE(&_BurnToAddressMintTokenPool.CallOpts)
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BurnToAddressMintTokenPool.Contract.GetRoleAdmin(&_BurnToAddressMintTokenPool.CallOpts, role)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BurnToAddressMintTokenPool.Contract.GetRoleAdmin(&_BurnToAddressMintTokenPool.CallOpts, role)
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) HasRateLimitAdminRole(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "hasRateLimitAdminRole", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) HasRateLimitAdminRole(account common.Address) (bool, error) {
	return _BurnToAddressMintTokenPool.Contract.HasRateLimitAdminRole(&_BurnToAddressMintTokenPool.CallOpts, account)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) HasRateLimitAdminRole(account common.Address) (bool, error) {
	return _BurnToAddressMintTokenPool.Contract.HasRateLimitAdminRole(&_BurnToAddressMintTokenPool.CallOpts, account)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BurnToAddressMintTokenPool.Contract.HasRole(&_BurnToAddressMintTokenPool.CallOpts, role, account)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BurnToAddressMintTokenPool.Contract.HasRole(&_BurnToAddressMintTokenPool.CallOpts, role, account)
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactor) GrantRateLimitAdminRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.contract.Transact(opts, "grantRateLimitAdminRole", account)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) GrantRateLimitAdminRole(account common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.GrantRateLimitAdminRole(&_BurnToAddressMintTokenPool.TransactOpts, account)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorSession) GrantRateLimitAdminRole(account common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.GrantRateLimitAdminRole(&_BurnToAddressMintTokenPool.TransactOpts, account)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.contract.Transact(opts, "grantRole", role, account)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.GrantRole(&_BurnToAddressMintTokenPool.TransactOpts, role, account)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.GrantRole(&_BurnToAddressMintTokenPool.TransactOpts, role, account)
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.RenounceRole(&_BurnToAddressMintTokenPool.TransactOpts, role, callerConfirmation)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.RenounceRole(&_BurnToAddressMintTokenPool.TransactOpts, role, callerConfirmation)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactor) RevokeRateLimitAdminRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.contract.Transact(opts, "revokeRateLimitAdminRole", account)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) RevokeRateLimitAdminRole(account common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.RevokeRateLimitAdminRole(&_BurnToAddressMintTokenPool.TransactOpts, account)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorSession) RevokeRateLimitAdminRole(account common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.RevokeRateLimitAdminRole(&_BurnToAddressMintTokenPool.TransactOpts, account)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.contract.Transact(opts, "revokeRole", role, account)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.RevokeRole(&_BurnToAddressMintTokenPool.TransactOpts, role, account)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.RevokeRole(&_BurnToAddressMintTokenPool.TransactOpts, role, account)
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

type BurnToAddressMintTokenPoolRateLimitAdminRoleGrantedIterator struct {
	Event *BurnToAddressMintTokenPoolRateLimitAdminRoleGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolRateLimitAdminRoleGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolRateLimitAdminRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnToAddressMintTokenPoolRateLimitAdminRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnToAddressMintTokenPoolRateLimitAdminRoleGrantedIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolRateLimitAdminRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolRateLimitAdminRoleGranted struct {
	Account common.Address
	Raw     types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterRateLimitAdminRoleGranted(opts *bind.FilterOpts) (*BurnToAddressMintTokenPoolRateLimitAdminRoleGrantedIterator, error) {

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "RateLimitAdminRoleGranted")
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolRateLimitAdminRoleGrantedIterator{contract: _BurnToAddressMintTokenPool.contract, event: "RateLimitAdminRoleGranted", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchRateLimitAdminRoleGranted(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolRateLimitAdminRoleGranted) (event.Subscription, error) {

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "RateLimitAdminRoleGranted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolRateLimitAdminRoleGranted)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "RateLimitAdminRoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseRateLimitAdminRoleGranted(log types.Log) (*BurnToAddressMintTokenPoolRateLimitAdminRoleGranted, error) {
	event := new(BurnToAddressMintTokenPoolRateLimitAdminRoleGranted)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "RateLimitAdminRoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnToAddressMintTokenPoolRateLimitAdminRoleRevokedIterator struct {
	Event *BurnToAddressMintTokenPoolRateLimitAdminRoleRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolRateLimitAdminRoleRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolRateLimitAdminRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnToAddressMintTokenPoolRateLimitAdminRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnToAddressMintTokenPoolRateLimitAdminRoleRevokedIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolRateLimitAdminRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolRateLimitAdminRoleRevoked struct {
	Account common.Address
	Raw     types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterRateLimitAdminRoleRevoked(opts *bind.FilterOpts) (*BurnToAddressMintTokenPoolRateLimitAdminRoleRevokedIterator, error) {

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "RateLimitAdminRoleRevoked")
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolRateLimitAdminRoleRevokedIterator{contract: _BurnToAddressMintTokenPool.contract, event: "RateLimitAdminRoleRevoked", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchRateLimitAdminRoleRevoked(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolRateLimitAdminRoleRevoked) (event.Subscription, error) {

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "RateLimitAdminRoleRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolRateLimitAdminRoleRevoked)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "RateLimitAdminRoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseRateLimitAdminRoleRevoked(log types.Log) (*BurnToAddressMintTokenPoolRateLimitAdminRoleRevoked, error) {
	event := new(BurnToAddressMintTokenPoolRateLimitAdminRoleRevoked)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "RateLimitAdminRoleRevoked", log); err != nil {
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

type BurnToAddressMintTokenPoolRoleAdminChangedIterator struct {
	Event *BurnToAddressMintTokenPoolRoleAdminChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolRoleAdminChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnToAddressMintTokenPoolRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnToAddressMintTokenPoolRoleAdminChangedIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BurnToAddressMintTokenPoolRoleAdminChangedIterator, error) {

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

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolRoleAdminChangedIterator{contract: _BurnToAddressMintTokenPool.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolRoleAdminChanged)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseRoleAdminChanged(log types.Log) (*BurnToAddressMintTokenPoolRoleAdminChanged, error) {
	event := new(BurnToAddressMintTokenPoolRoleAdminChanged)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnToAddressMintTokenPoolRoleGrantedIterator struct {
	Event *BurnToAddressMintTokenPoolRoleGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolRoleGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnToAddressMintTokenPoolRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnToAddressMintTokenPoolRoleGrantedIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnToAddressMintTokenPoolRoleGrantedIterator, error) {

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

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolRoleGrantedIterator{contract: _BurnToAddressMintTokenPool.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolRoleGranted)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseRoleGranted(log types.Log) (*BurnToAddressMintTokenPoolRoleGranted, error) {
	event := new(BurnToAddressMintTokenPoolRoleGranted)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnToAddressMintTokenPoolRoleRevokedIterator struct {
	Event *BurnToAddressMintTokenPoolRoleRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolRoleRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnToAddressMintTokenPoolRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnToAddressMintTokenPoolRoleRevokedIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnToAddressMintTokenPoolRoleRevokedIterator, error) {

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

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolRoleRevokedIterator{contract: _BurnToAddressMintTokenPool.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolRoleRevoked)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseRoleRevoked(log types.Log) (*BurnToAddressMintTokenPoolRoleRevoked, error) {
	event := new(BurnToAddressMintTokenPoolRoleRevoked)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

func (BurnToAddressMintTokenPoolRateLimitAdminRoleGranted) Topic() common.Hash {
	return common.HexToHash("0xf7af318a70f367e30346e2704f6ef646b378a7dcb49767beb98a1774cd11e389")
}

func (BurnToAddressMintTokenPoolRateLimitAdminRoleRevoked) Topic() common.Hash {
	return common.HexToHash("0xd63806009f622849e3b7cfd82d762420d57574c39f945f678871b2b5f1e8ce4b")
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

func (BurnToAddressMintTokenPoolRoleAdminChanged) Topic() common.Hash {
	return common.HexToHash("0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff")
}

func (BurnToAddressMintTokenPoolRoleGranted) Topic() common.Hash {
	return common.HexToHash("0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d")
}

func (BurnToAddressMintTokenPoolRoleRevoked) Topic() common.Hash {
	return common.HexToHash("0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b")
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
	DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error)

	RATELIMITERADMINROLE(opts *bind.CallOpts) ([32]byte, error)

	GetAccumulatedFees(opts *bind.CallOpts) (*big.Int, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetBurnAddress(opts *bind.CallOpts) (common.Address, error)

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

	GrantRateLimitAdminRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error)

	GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error)

	RevokeRateLimitAdminRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error)

	RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error)

	SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error)

	SetCustomFinalityRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error)

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

	FilterRateLimitAdminRoleGranted(opts *bind.FilterOpts) (*BurnToAddressMintTokenPoolRateLimitAdminRoleGrantedIterator, error)

	WatchRateLimitAdminRoleGranted(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolRateLimitAdminRoleGranted) (event.Subscription, error)

	ParseRateLimitAdminRoleGranted(log types.Log) (*BurnToAddressMintTokenPoolRateLimitAdminRoleGranted, error)

	FilterRateLimitAdminRoleRevoked(opts *bind.FilterOpts) (*BurnToAddressMintTokenPoolRateLimitAdminRoleRevokedIterator, error)

	WatchRateLimitAdminRoleRevoked(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolRateLimitAdminRoleRevoked) (event.Subscription, error)

	ParseRateLimitAdminRoleRevoked(log types.Log) (*BurnToAddressMintTokenPoolRateLimitAdminRoleRevoked, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnToAddressMintTokenPoolReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*BurnToAddressMintTokenPoolReleasedOrMinted, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnToAddressMintTokenPoolRemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*BurnToAddressMintTokenPoolRemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnToAddressMintTokenPoolRemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*BurnToAddressMintTokenPoolRemotePoolRemoved, error)

	FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BurnToAddressMintTokenPoolRoleAdminChangedIterator, error)

	WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error)

	ParseRoleAdminChanged(log types.Log) (*BurnToAddressMintTokenPoolRoleAdminChanged, error)

	FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnToAddressMintTokenPoolRoleGrantedIterator, error)

	WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleGranted(log types.Log) (*BurnToAddressMintTokenPoolRoleGranted, error)

	FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnToAddressMintTokenPoolRoleRevokedIterator, error)

	WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleRevoked(log types.Log) (*BurnToAddressMintTokenPoolRoleRevoked, error)

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
