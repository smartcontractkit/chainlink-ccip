// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lombard_token_pool

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

type IPoolV2TokenTransferFeeConfig struct {
	DestGasOverhead            uint32
	DestBytesOverhead          uint32
	FinalityFeeUSDCents        uint32
	FastFinalityFeeUSDCents    uint32
	FinalityTransferFeeBps     uint16
	FastFinalityTransferFeeBps uint16
	IsEnabled                  bool
}

type LombardTokenPoolPath struct {
	AllowedCaller [32]byte
	LChainId      [32]byte
	RemoteAdapter [32]byte
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

type TokenPoolChainUpdate struct {
	RemoteChainSelector       uint64
	RemotePoolAddresses       [][]byte
	RemoteTokenAddress        []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
}

type TokenPoolRateLimitConfigArgs struct {
	RemoteChainSelector       uint64
	FastFinality              bool
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
}

type TokenPoolTokenTransferFeeConfigArgs struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
}

var LombardTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20Metadata\"},{\"name\":\"verifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bridge\",\"type\":\"address\",\"internalType\":\"contract IBridgeV2\"},{\"name\":\"adapter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"advancedPoolHooks\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fallbackDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fastFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fastFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"disableTokenTransferFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAdvancedPoolHooks\",\"inputs\":[],\"outputs\":[{\"name\":\"advancedPoolHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowedFinalityConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"allowedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fastFinality\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requestedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLombardConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"verifierResolver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAdapter\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPath\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct LombardTokenPool.Path\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"lChainId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteAdapter\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fastFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fastFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_bridge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IBridgeV2\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"lockOrBurnOut\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"requestedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"lockOrBurnOut\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"requestedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removePath\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAllowedFinalityConfig\",\"inputs\":[{\"name\":\"allowedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPath\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"lChainId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCaller\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteAdapter\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.RateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fastFinality\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateAdvancedPoolHooks\",\"inputs\":[{\"name\":\"newHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AdvancedPoolHooksUpdated\",\"inputs\":[{\"name\":\"oldHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"},{\"name\":\"newHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"feeAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FastFinalityInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FastFinalityOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalityConfigSet\",\"inputs\":[{\"name\":\"allowedFinality\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LombardConfigurationSet\",\"inputs\":[{\"name\":\"verifier\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"bridge\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenAdapter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PathRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"lChainId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"remoteAdapter\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PathSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"lChainId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"remoteAdapter\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"fastFinality\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fastFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fastFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotOwnerOrFeeAdmin\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotSupported\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"ExecutionError\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"HashMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Invalid32ByteAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"received\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRequestedFinality\",\"inputs\":[{\"name\":\"requestedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"allowedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OutboundImplementationNotFoundForVerifier\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PathNotExist\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"RemoteTokenOrAdapterMismatch\",\"inputs\":[{\"name\":\"bridgeToken\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteToken\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"remoteAdapter\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"RequestedFinalityCanOnlyHaveOneMode\",\"inputs\":[{\"name\":\"encodedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroBridge\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroLombardChainId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroVerifierNotAllowed\",\"inputs\":[]}]",
	Bin: "0x6101408060405234610301576101008161646180380380916100218285610401565b8339810103126103015780516001600160a01b038116918282036103015761004b6020820161043a565b6040820151936001600160a01b03851692909190838603610301576100726060820161043a565b9461007f6080830161043a565b906100ae61008f60a0850161043a565b916100a860e06100a160c0880161043a565b960161044e565b9061045c565b9033156103f057600180546001600160a01b03191633179055851580156103df575b80156103ce575b6103bd578560805260c05230850361032b575b60a052600380546001600160a01b039283166001600160a01b03199182161790915560028054939092169216919091179055821561031a5760405163353c26b760e01b8152602081600481875afa801561030e576000906102cf575b60ff915016600281036102b657506001600160a01b0381169485156102a55760e052610100526101208390526001600160a01b03831692828415610295575061018e916104d2565b604051927f01d5dd7f15328f4241da3a1d9c7b310ae9ac14e8ca441203a7b6f71c7da0c49d600080a4615e7690816105eb82396080518181816102eb015281816103440152818161227d0152818161262b015281816127cf01528181612c7701528181612fcc0152818161357f01528181613c7b0152613cc8015260a05181818161058801528181613b4e01528181614f570152614fda015260c0518181816103cd0152818161142e0152818161230a01528181612d05015261360d015260e0518181816111cc0152818161248a01528181612b1f0152612eb001526101005181818161029b015281816111a70152613e8b0152610120518181816111f401526124500152f35b90506102a0916104d2565b61018e565b639533e8c360e01b60005260046000fd5b63398bbe0560e11b600052600260045260245260446000fd5b506020813d602011610306575b816102e960209383610401565b81010312610301576102fc60ff9161044e565b610146565b600080fd5b3d91506102dc565b6040513d6000823e3d90fd5b63361106cd60e01b60005260046000fd5b60405163313ce56760e01b8152602081600481895afa60009181610381575b50610356575b506100ea565b60ff1660ff821681810361036a5750610350565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d6020116103b5575b8161039d60209383610401565b81010312610301576103ae9061044e565b903861034a565b3d9150610390565b630a64406560e11b60005260046000fd5b506001600160a01b038116156100d7565b506001600160a01b038416156100d0565b639b15e16f60e01b60005260046000fd5b601f909101601f19168101906001600160401b0382119082101761042457604052565b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b038216820361030157565b519060ff8216820361030157565b60405163313ce56760e01b815290602090829060049082906001600160a01b03165afa60009181610496575b50610491575090565b905090565b9091506020813d6020116104ca575b816104b260209383610401565b81010312610301576104c39061044e565b9038610488565b3d91506104a5565b60405190602060008184019463095ea7b360e01b865260018060a01b031694856024860152811960448601526044855261050d606486610401565b84519082855af16000513d8261056a575b50501561052a57505050565b610563610568936040519063095ea7b360e01b60208301526024820152600060448201526044815261055d606482610401565b8261058f565b61058f565b565b90915061058757506001600160a01b0381163b15155b388061051e565b600114610580565b906000602091828151910182855af11561030e576000513d6105e157506001600160a01b0381163b155b6105c05750565b635274afe760e01b60009081526001600160a01b0391909116600452602490fd5b600114156105b956fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a71461414d575080630606699c14613f3557806306b859ef14613e04578063181f5a7714613da35780631826b1e714613cec57806321df0da714613ca8578063240028e814613c515780632422ac4514613b7257806324f65ee714613b345780632cab0fb6146134ef57806337a3210d146134c857806338ff8c38146134415780633907753714612c115780634c5ef0ed14612bca57806362ddd3c414612b43578063708e1f7914612aff5780637437ff9f14612abe57806379ba509714612a115780638926f54f146129cb5780638da5cb5b146129a45780639a4575b9146122115780639c893fe91461212a578063a42a7b8b14611fc3578063acfecf9114611ead578063ae39a25714611d56578063b6cfa3b714611c9b578063b794658014611c67578063bfeffd3f14611bd5578063c4bffe2b14611aaa578063c7230a60146118d6578063dc04fa1f14611452578063dc0bd9711461140e578063dcbd41bc14611224578063dd65bdb114611188578063e8a1da1714610ae4578063ea6396db146109aa578063ec6ae7a714610967578063f2fde38b146108b25763fbc801a7146101ce57600080fd5b346106f35760606003193601126106f3576004359067ffffffffffffffff82116106f357816004019060a060031984360301126106f35761020d6142a9565b60443567ffffffffffffffff81116108ae5761022d90369060040161434a565b929091610238614a6e565b506024860195610247876149f0565b9367ffffffffffffffff604051957f958021a7000000000000000000000000000000000000000000000000000000008752166004860152604060248601528360448601526020856064816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9485156108a3578495610882575b506001600160a01b0385161561085a5761032a9061030f606484013580977f0000000000000000000000000000000000000000000000000000000000000000615524565b610317614a6e565b506103228489615b8f565b963691614577565b906084810196610339886149dc565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691160361081d5777ffffffffffffffff0000000000000000000000000000000061038d8a6149f0565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156107815786916107df575b506107b75767ffffffffffffffff6104148a6149f0565b1661042c816000526007602052604060002054151590565b1561078c5760206001600160a01b0360025416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa90811561078157906001600160a01b03918791610752575b50163303610726577fffffffff000000000000000000000000000000000000000000000000000000006104be8888614d6a565b9416938415610705576104e3816104d48b6149dc565b6104dd8d6149f0565b9061584b565b6001600160a01b036003541693846105ee575b6105e48a61058061057b8e61050b8e8e614d6a565b93610515826149f0565b507ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff61055261054c856149f0565b936149dc565b604080516001600160a01b039290921682523360208301528101889052921691606090a26149f0565b614b9c565b9060405160ff7f0000000000000000000000000000000000000000000000000000000000000000166020820152602081526105bc604082614489565b604051926105c984614435565b83526020830152604051928392604084526040840190614611565b9060208301520390f35b843b15610701578694928a949286928d604051998a98899788967fa8027c0f00000000000000000000000000000000000000000000000000000000885260048801608090528061063d916157b5565b6084890160a090526101248901906106549261469d565b9361065e90614335565b67ffffffffffffffff1660a4880152604401610679906143ba565b6001600160a01b031660c48701528d60e4870152610696906143ba565b6001600160a01b031661010486015260248501528381036003190160448501526106bf91614527565b90606483015203925af180156106f6576106de575b80808080806104f6565b6106e9828092614489565b6106f357806106d4565b80fd5b6040513d84823e3d90fd5b8680fd5b610721816107128b6149dc565b61071b8d6149f0565b90615805565b6104e3565b6024857f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b610774915060203d60201161077a575b61076c8183614489565b810190614a05565b3861048b565b503d610762565b6040513d88823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008652600452602485fd5b6004857f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b90506020813d602011610815575b816107fa60209383614489565b810103126108115761080b90614a24565b386103fd565b8580fd5b3d91506107ed565b6024856001600160a01b036108318b6149dc565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835216600452fd5b6004847f7af97002000000000000000000000000000000000000000000000000000000008152fd5b61089c91955060203d60201161077a5761076c8183614489565b93386102cb565b6040513d86823e3d90fd5b8280fd5b50346106f35760206003193601126106f3576001600160a01b036108d4614378565b6108dc614d2c565b1633811461093f57807fffffffffffffffffffffffff00000000000000000000000000000000000000008354161782556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b50346106f357806003193601126106f35760207fffffffff0000000000000000000000000000000000000000000000000000000060025460401b16604051908152f35b50346106f35760806003193601126106f3576109c4614378565b506109cd61431e565b6109d56142d8565b5060643567ffffffffffffffff81116108ae579167ffffffffffffffff604092610a0560e095369060040161434a565b50508260c08551610a158161446d565b82815282602082015282878201528260608201528260808201528260a08201520152168152600b6020522060405190610a4d8261446d565b5461ffff818163ffffffff82169485815263ffffffff60208201818560201c1681528160408401818760401c168152816060860193818960601c16855260ff60c060808901988a8c60801c168a528a60a082019c60901c168c52019b60a01c1615158b526040519b8c52511660208b0152511660408901525116606087015251166080850152511660a083015251151560c0820152f35b50346106f35760406003193601126106f35760043567ffffffffffffffff8111610fbd57610b1690369060040161463b565b9060243567ffffffffffffffff81116111845790610b398492369060040161463b565b939091610b44614d2c565b83905b828210610fc55750505081927ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b81811015610fc1578060051b83013585811215610fb957830161012081360312610fb95760405194610bab86614451565b610bb482614335565b8652602082013567ffffffffffffffff8111610fbd5782019436601f87011215610fbd57853595610be4876146fd565b96610bf26040519889614489565b80885260208089019160051b83010190368211610fb95760208301905b828210610f8a575050505060208701958652604083013567ffffffffffffffff81116108ae57610c4290369085016145ae565b9160408801928352610c6c610c5a3660608701614c48565b9460608a0195865260c0369101614c48565b956080890196875283515115610f6257610c9067ffffffffffffffff8a5116615ada565b15610f2b5767ffffffffffffffff8951168252600860205260408220610cb78651826152f4565b610cc58851600283016152f4565b6004855191019080519067ffffffffffffffff8211610efe57610ce88354614a87565b601f8111610ec3575b50602090601f8311600114610e4257610d219291869183610e37575b50506000198260011b9260031b1c19161790565b90555b815b88518051821015610d5b5790610d55600192610d4e8367ffffffffffffffff8f511692614715565b51906150c6565b01610d26565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2939199975095610e2967ffffffffffffffff6001979694985116925193519151610df5610dc060405196879687526101006020880152610100870190614527565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a1019392909193610b7a565b015190508e80610d0d565b83865281862091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416875b818110610eab5750908460019594939210610e92575b505050811b019055610d24565b015160001960f88460031b161c191690558d8080610e85565b92936020600181928786015181550195019301610e6f565b610eee9084875260208720601f850160051c81019160208610610ef4575b601f0160051c0190614cd3565b8d610cf1565b9091508190610ee1565b6024857f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60248267ffffffffffffffff8b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b6004827f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b813567ffffffffffffffff811161070157602091610fae83928336918901016145ae565b815201910190610c0f565b8480fd5b5080fd5b8380f35b9267ffffffffffffffff610fe7610fe28486889a9699979a614c1b565b6149f0565b1691610ff28361591e565b1561115857828452600860205261100e600560408620016158bb565b94845b865181101561104757600190858752600860205261104060056040892001611039838b614715565b5190615a1e565b5001611011565b50939692909450949094808752600860205260056040882088815588600182015588600282015588600382015588600482016110838154614a87565b80611117575b50505001805490888155816110f9575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020836001948a52600482528985604082208281550155808a52600582528985604082208281550155604051908152a101909194939294610b47565b885260208820908101905b8181101561109957888155600101611104565b601f811160011461112d5750555b888a80611089565b8183526020832061114891601f01861c810190600101614cd3565b8082528160208120915555611125565b602484847f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b8380fd5b50346106f357806003193601126106f357604080516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811682527f0000000000000000000000000000000000000000000000000000000000000000811660208301527f00000000000000000000000000000000000000000000000000000000000000001691810191909152606090f35b0390f35b50346106f35760206003193601126106f35760043567ffffffffffffffff8111610fbd5761125690369060040161466c565b6001600160a01b03600a5416331415806113f9575b6113cd57825b81811061127c578380f35b611287818385614bbe565b67ffffffffffffffff611299826149f0565b16906112b2826000526007602052604060002054151590565b156113a157907f41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb60e08361136161133b602060019897018b6112f382614bce565b1561136857879052600460205261131a60408d206113143660408801614c48565b906152f4565b868c52600560205261133660408d206113143660a08801614c48565b614bce565b9160405192151583526113546020840160408301614c8f565b60a0608084019101614c8f565ba201611271565b60026040828a6113369452600860205261138a82822061131436858c01614c48565b8a8152600860205220016113143660a08801614c48565b602486837f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b6024837f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b506001600160a01b036001541633141561126b565b50346106f357806003193601126106f35760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346106f35760406003193601126106f35760043567ffffffffffffffff8111610fbd5761148490369060040161466c565b60243567ffffffffffffffff8111611184576114a490369060040161463b565b9190926114af614d2c565b845b82811061151b57505050825b8181106114c8578380f35b8067ffffffffffffffff6114e2610fe26001948688614c1b565b16808652600b6020528560408120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee88680a2016114bd565b67ffffffffffffffff611532610fe2838686614bbe565b1661154a816000526007602052604060002054151590565b156118ab5761155a828585614bbe565b602081019060e081019061156d82614bce565b1561187f5760a0810161271061ffff61158583614bdb565b1610156118705760c082019161271061ffff6115a085614bdb565b1610156118385763ffffffff6115b586614bea565b161561180c57858c52600b60205260408c206115d086614bea565b63ffffffff169080549060408401916115e883614bea565b60201b67ffffffff000000001693606086019461160486614bea565b60401b6bffffffff000000000000000016966080019661162388614bea565b60601b6fffffffff00000000000000000000000016916116428a614bdb565b60801b71ffff0000000000000000000000000000000016936116638c614bdb565b60901b73ffff00000000000000000000000000000000000016957fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016177fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16177fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff16177fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff16171717815561171687614bce565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff0000000000000000000000000000000000000000161790556040519661176790614bfb565b63ffffffff16875261177890614bfb565b63ffffffff16602087015261178c90614bfb565b63ffffffff1660408601526117a090614bfb565b63ffffffff1660608501526117b490614c0c565b61ffff1660808401526117c690614c0c565b61ffff1660a08301526117d89061456a565b151560c082015260e07ffae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa104191a26001016114b1565b60248c877f12332265000000000000000000000000000000000000000000000000000000008252600452fd5b60248c61ffff61184786614bdb565b7f95f3517a00000000000000000000000000000000000000000000000000000000835216600452fd5b8a61ffff611847602493614bdb565b60248a857f12332265000000000000000000000000000000000000000000000000000000008252600452fd5b7f1e670e4b000000000000000000000000000000000000000000000000000000008752600452602486fd5b50346106f35760406003193601126106f35760043567ffffffffffffffff8111610fbd5761190890369060040161463b565b906119116143a4565b916001600160a01b036001541633141580611a95575b611a69576001600160a01b038316908115611a4157845b818110611949578580f35b6001600160a01b0361196461195f838588614c1b565b6149dc565b1690604051917f70a08231000000000000000000000000000000000000000000000000000000008352306004840152602083602481845afa8015611a3657859089906119fc575b60019450806119be575b5050500161193e565b6020816119ed7f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e938c87615524565b604051908152a33884816119b5565b5050909160203d8111611a2f575b611a148183614489565b602082600092810103126106f35750908460019392516119ab565b503d611a0a565b6040513d8a823e3d90fd5b6004857f8579befe000000000000000000000000000000000000000000000000000000008152fd5b6024847fcb1afbd700000000000000000000000000000000000000000000000000000000815233600452fd5b506001600160a01b03600c5416331415611927565b50346106f357806003193601126106f357604051906006548083528260208101600684526020842092845b818110611bbc575050611aea92500383614489565b8151611b0e611af8826146fd565b91611b066040519384614489565b8083526146fd565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602083019301368437805b8451811015611b6d578067ffffffffffffffff611b5a60019388614715565b5116611b668286614715565b5201611b3b565b50925090604051928392602084019060208552518091526040840192915b818110611b99575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101611b8b565b8454835260019485019487945060209093019201611ad5565b50346106f35760206003193601126106f3576004356001600160a01b038116809103610fbd57611c03614d2c565b7fffffffffffffffffffffffff00000000000000000000000000000000000000006003547fbaff46844acf36d6ee996f489a1a288709c4542bd33cd557770afd267d581209604080516001600160a01b0384168152856020820152a1161760035580f35b50346106f35760206003193601126106f357611220611c8761057b614307565b604051918291602083526020830190614527565b50346106f35760206003193601126106f3577f307cf716eade81675bea3ccb6917b0f91baa2160056765d9a83d76f819caf06a6020611cd861424b565b611ce0614d2c565b6002547fffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffff77ffffffff0000000000000000000000000000000000000000808460401c16169116176002557fffffffff0000000000000000000000000000000000000000000000000000000060405191168152a180f35b50346106f35760606003193601126106f357611d70614378565b90611d796143a4565b604435926001600160a01b03841680850361118457611d96614d2c565b6001600160a01b0382168015611e855794611e7f917f3f1036e85d016a93254a0b1415844f79b85424959d90ae5ad51ce8f4533fe70195967fffffffffffffffffffffffff000000000000000000000000000000000000000060025416176002556001600160a01b0385167fffffffffffffffffffffffff0000000000000000000000000000000000000000600a541617600a557fffffffffffffffffffffffff0000000000000000000000000000000000000000600c541617600c55604051938493849160409194936001600160a01b03809281606087019816865216602085015216910152565b0390a180f35b6004857f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b50346106f357611ebc366145cc565b611ec4614d2c565b67ffffffffffffffff831692611ee7846000526007602052604060002054151590565b15611f97578385526008602052611f1660056040872001611f09368587614577565b6020815191012090615a1e565b15611f5c5750907f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d7691611f5660405192839260208452602084019161469d565b0390a280f35b90611f93906040519384937f74f23c7c000000000000000000000000000000000000000000000000000000008552600485016146dc565b0390fd5b602485857f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50346106f35760206003193601126106f35767ffffffffffffffff611fe6614307565b1681526008602052611ffd600560408320016158bb565b80517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061204261202c836146fd565b9261203a6040519485614489565b8084526146fd565b01835b818110612119575050825b8251811015612096578061206660019285614715565b518552600960205261207a60408620614ada565b6120848285614715565b5261208f8184614715565b5001612050565b81846040519182916020830160208452825180915260408401602060408360051b870101940192905b8282106120ce57505050500390f35b91936020612109827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851614527565b96019201920185949391926120bf565b806060602080938601015201612045565b50346106f35760206003193601126106f35767ffffffffffffffff61214d614307565b612155614d2c565b16808252600d602052604082209060405191612170836143ce565b80549081845260026001820154916020860192835201549160408501928352156121e5577f465d9b27e0af9978f975c48406a226aab254b237e8027798ee924ef96ee9bb0491604091848752600d6020528660028482208281558260018201550155519451905182519182526020820152a380f35b602485847fa28cbf38000000000000000000000000000000000000000000000000000000008252600452fd5b50346106f35760206003193601126106f35760043567ffffffffffffffff8111610fbd57806004019160a060031983360301126106f357612250614a6e565b506040516020926122618483614489565b8282526084810191612272836149dc565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691160361299057602482019077ffffffffffffffff000000000000000000000000000000006122cb836149f0565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015286816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa90811561078157869161295b575b506107b75767ffffffffffffffff612351836149f0565b16612369816000526007602052604060002054151590565b1561078c57866001600160a01b0360025416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa90811561078157906001600160a01b0391879161293e575b50163303610726576064830135936123eb856123e2836149dc565b61071b866149f0565b6001600160a01b036003541680612833575b50505067ffffffffffffffff612412826149f0565b168452600d855260408420956040519661242b886143ce565b805497888152600260018301549289830193845201549860408201998a52156127f5577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b038116156127c457905b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016926001600160a01b03815193604051947f6e48b60d000000000000000000000000000000000000000000000000000000008652600486015216928360248201528a81604481885afa9081156127b9578a9161278c575b5061250d61057b886149f0565b9b8c8c81519181808201938492010103126126f4575190818303612741575b5050508961253a868061498b565b9050036126f957604490519601958961255c612555896149dc565b968061498b565b90809291810103126126f45760409460c4938b92359051906001600160a01b038851998a9889977f793ea55b0000000000000000000000000000000000000000000000000000000089526004890152602488015216604486015260648501528a608485015260a48401525af19485156126e85780956126ad575b50507ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae109161266d61261861261267ffffffffffffffff946149f0565b926149dc565b9460405193849316956001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016846001600160a01b036040929594938160608401971683521660208201520152565b0390a26040519082820152818152612686604082614489565b6040519261269384614435565b835281830152611220604051928284938452830190614611565b909194506040823d6040116126e0575b816126ca60409383614489565b810103126106f35750840151928161266d6125d6565b3d91506126bd565b604051903d90823e3d90fd5b600080fd5b611f938a612707878061498b565b92906040519384937fa3c8cf090000000000000000000000000000000000000000000000000000000085526004850152602484019161469d565b519182158015612782575b1561252c577fbce7b6cd000000000000000000000000000000000000000000000000000000008c52600452602452604452606489fd5b508281141561274c565b90508a81813d83116127b2575b6127a38183614489565b810103126126f4575138612500565b503d612799565b6040513d8c823e3d90fd5b506001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690612480565b60248767ffffffffffffffff61280a876149f0565b7fa28cbf3800000000000000000000000000000000000000000000000000000000835216600452fd5b803b156107015790869189836040518096819582947fa8027c0f00000000000000000000000000000000000000000000000000000000845260048401608090528061287d916157b5565b6084850160a090526101248501906128949261469d565b9061289e8b614335565b67ffffffffffffffff1660a48501526128b960448d016143ba565b6001600160a01b031660c48501528c60e48501526128d6906143ba565b6001600160a01b031661010484015283602484015282810360031901604484015261290091614527565b8a606483015203925af180156129335790859161291e575b806123fd565b8161292891614489565b611184578338612918565b6040513d87823e3d90fd5b6129559150883d8a1161077a5761076c8183614489565b386123c7565b90508681813d8311612989575b6129728183614489565b810103126108115761298390614a24565b3861233a565b503d612968565b6024846001600160a01b03610831866149dc565b50346106f357806003193601126106f35760206001600160a01b0360015416604051908152f35b50346106f35760206003193601126106f3576020612a0767ffffffffffffffff6129f3614307565b166000526007602052604060002054151590565b6040519015158152f35b50346106f357806003193601126106f35780546001600160a01b0381163303612a96577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551682556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b50346106f357806003193601126106f357600254600a54600c54604080516001600160a01b0394851681529284166020840152921691810191909152606090f35b50346106f357806003193601126106f35760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346106f357612b52366145cc565b612b5e93929193614d2c565b67ffffffffffffffff8216612b80816000526007602052604060002054151590565b15612b9f5750612b9c9293612b96913691614577565b906150c6565b80f35b7f1e670e4b000000000000000000000000000000000000000000000000000000008452600452602483fd5b50346106f35760406003193601126106f357612be4614307565b906024359067ffffffffffffffff82116106f3576020612a0784612c0b36600487016145ae565b90614a31565b50346106f35760206003193601126106f3576004359067ffffffffffffffff82116106f35781600401916101006003198236030112610fbd5781604051612c5781614419565b526064810135916084820193612c6c856149dc565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691160361342d57602483019477ffffffffffffffff00000000000000000000000000000000612cc5876149f0565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156108a35784916133f3575b506133cb5767ffffffffffffffff612d4c876149f0565b16612d64816000526007602052604060002054151590565b156133a05760206001600160a01b0360025416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa9081156108a3578491613366575b501561333a57612dce866149f0565b90612deb60a4860192612c0b612de4858761498b565b3691614577565b156132f357612e0c86612dfd836149dc565b612e068a6149f0565b906156d9565b6001600160a01b03600354169081613144575b505050612e2f60e484018261498b565b810160408282031261118457813567ffffffffffffffff8111610fb95781612e589184016145ae565b9160208101359067ffffffffffffffff821161081157612e799291016145ae565b906040517fd5438eae0000000000000000000000000000000000000000000000000000000081526020816004816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156129335791612f489493918680948192613118575b506001600160a01b03612f369293604051988996879586937fa6208506000000000000000000000000000000000000000000000000000000008552604060048601526044850190614527565b90600319848303016024850152614527565b0393165af1801561310d578392849161307e575b501561305657612f7260209160c486019061498b565b90809291810103126126f457350361302e57507ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc067ffffffffffffffff612fc66044612fbf6020976149f0565b94016149dc565b604080517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b039081168252336020830152909216908201526060810185905292169180608081015b0390a28060405161302581614419565b52604051908152f35b807f3f4d60530000000000000000000000000000000000000000000000000000000060049252fd5b6004837f2532cf45000000000000000000000000000000000000000000000000000000008152fd5b9250503d8084843e6130908184614489565b8201606083820312611184578251906130ab60208501614a24565b9360408101519067ffffffffffffffff8211610701570181601f82011215610811578051916130d9836144ca565b906130e76040519283614489565b8382526020848401011161070157906020806131069493019101614504565b9138612f5c565b6040513d85823e3d90fd5b612f36925061313d6001600160a01b039160203d60201161077a5761076c8183614489565b9250612eea565b813b15610fb9579184918893836040518096819582947f63711574000000000000000000000000000000000000000000000000000000008452600484016060905261318f8b806157b5565b6064860161010090526101648601906131a79261469d565b926131b190614335565b67ffffffffffffffff1660848501526131cc60448e016143ba565b6001600160a01b031660a48501528d60c48501526131e9906143ba565b6001600160a01b031660e4840152613201908a6157b5565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c84840301610104850152613236929161469d565b61324360c48c018a6157b5565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c84840301610124850152613278929161469d565b61328560e48c018a6157b5565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c848403016101448501526132ba929161469d565b8b602483015282604483015203925af1801561310d579083916132de575b80612e1f565b816132e891614489565b610fbd5781386132d8565b506132fd9161498b565b611f936040519283927f24eb47e500000000000000000000000000000000000000000000000000000000845260206004850152602484019161469d565b6024837f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b90506020813d602011613398575b8161338160209383614489565b810103126111845761339290614a24565b38612dbf565b3d9150613374565b7fa9902c7e000000000000000000000000000000000000000000000000000000008452600452602483fd5b6004837f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b90506020813d602011613425575b8161340e60209383614489565b810103126111845761341f90614a24565b38612d35565b3d9150613401565b6024826001600160a01b03610831886149dc565b50346106f35760206003193601126106f357604060609167ffffffffffffffff613469614307565b82848051613476816143ce565b8281528260208201520152168152600d60205220604051613496816143ce565b815491828252604060026001830154926020850193845201549201918252604051928352516020830152516040820152f35b50346106f357806003193601126106f35760206001600160a01b0360035416604051908152f35b50346106f35760406003193601126106f35760043567ffffffffffffffff8111610fbd578060040161010060031983360301126108ae5761352e6142a9565b8360405161353b81614419565b5260648301359260c481019461356561355f61355a612de4898861498b565b614ee3565b86614fd7565b946084830193613574856149dc565b6001600160a01b03807f000000000000000000000000000000000000000000000000000000000000000016911603613b2057602484019577ffffffffffffffff000000000000000000000000000000006135cd886149f0565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115612933578591613ae6575b50613abe5767ffffffffffffffff613654886149f0565b1661366c816000526007602052604060002054151590565b15613a935760206001600160a01b0360025416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115612933578591613a59575b5015613a2d576136d6876149f0565b926136ec60a4870194612c0b612de4878661498b565b15613a23577fffffffff000000000000000000000000000000000000000000000000000000008316928315613a075761374d907fffffffff0000000000000000000000000000000000000000000000000000000060025460401b1690614d77565b6137698961375a896149dc565b6137638b6149f0565b90615745565b6001600160a01b03600354169182613814575b60208a8a7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc067ffffffffffffffff8c613015856137e38f6137d760446137dd9201986137c78a6149dc565b506137d1816149f0565b506149f0565b946149dc565b966149dc565b604080516001600160a01b039889168152336020820152979091169087015260608601529116929081906080820190565b859a839695963b15610811578894869289928c6040519a8b998a9889977f63711574000000000000000000000000000000000000000000000000000000008952600489016060905261386687806157b5565b60648b0161010090526101648b019061387e9261469d565b9461388890614335565b67ffffffffffffffff1660848a01526044016138a3906143ba565b6001600160a01b031660a489015260c48801526138bf906143ba565b6001600160a01b031660e48701526138d790846157b5565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c8784030161010488015261390c929161469d565b9061391790836157b5565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c8684030161012487015261394c929161469d565b9060e48b0161395a916157b5565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c8584030161014486015261398f929161469d565b908c6024840152604483015203925af1801561078157856137e36137dd6137d7604467ffffffffffffffff9760209c7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09a98613015986139f7575b505097505085975061377c565b81613a0191614489565b386139ea565b50613a1e89613a15896149dc565b612e068b6149f0565b613769565b6132fd848361498b565b6024847f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b90506020813d602011613a8b575b81613a7460209383614489565b81010312610fb957613a8590614a24565b386136c7565b3d9150613a67565b7fa9902c7e000000000000000000000000000000000000000000000000000000008552600452602484fd5b6004847f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b90506020813d602011613b18575b81613b0160209383614489565b81010312610fb957613b1290614a24565b3861363d565b3d9150613af4565b6024836001600160a01b03610831886149dc565b50346106f357806003193601126106f357602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346106f35760406003193601126106f357613b8c614307565b6024359182151583036106f357610140613c4f613ba98585614908565b613bff60409392935180946fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b60a08301906fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565bf35b50346106f35760206003193601126106f357602090613c6e614378565b90506001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b50346106f357806003193601126106f35760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346106f35760c06003193601126106f357613d06614378565b50613d0f61431e565b613d1761438e565b50608435917fffffffff00000000000000000000000000000000000000000000000000000000831683036106f35760a4359067ffffffffffffffff82116106f35760a063ffffffff8061ffff613d7c8888613d753660048b0161434a565b5050614758565b95926040979194975197885216602087015216604085015216606083015215156080820152f35b50346106f357806003193601126106f35750611220604051613dc6604082614489565b601681527f4c6f6d62617264546f6b656e506f6f6c20322e302e30000000000000000000006020820152604051918291602083526020830190614527565b50346106f35760c06003193601126106f357613e1e614378565b50613e2761431e565b50613e3061427a565b5060843567ffffffffffffffff8111610fbd57613e5190369060040161434a565b5050600260a43510156106f357604051613e6c606082614489565b6002815260208101906040368337805115613f08576001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168252805160011015613f085790826040830152604051928392602084019060208552518091526040840192915b818110613ee6575050500390f35b82516001600160a01b0316845285945060209384019390920191600101613ed8565b6024837f4e487b710000000000000000000000000000000000000000000000000000000081526032600452fd5b50346106f35760806003193601126106f357613f4f614307565b6024359060443567ffffffffffffffff811161118457613f7390369060040161434a565b9060643591613f80614d2c565b67ffffffffffffffff841693613fa3856000526007602052604060002054151590565b156141215785156140f957613fc2613fbc368486614577565b82614a31565b15611f5c575090613fd4913691614577565b60208151116140bb576020815191015190602081106140a8575b8060031b908082046008149015171561407b5761010003610100811161407b57916040917fbe237be2cca72f95760d5f21feb5f0cf6579119971f023d6ccc49c749ddc9263931c908251614041816143ce565b82815260026020820188815285830190848252888b52600d602052868b20935184555160018401555191015582519182526020820152a380f35b6024867f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b906000198260200360031b1b1690613fee565b611f93906040519182917fe0d7fb02000000000000000000000000000000000000000000000000000000008352602060048401526024830190614527565b6004877f5a39e303000000000000000000000000000000000000000000000000000000008152fd5b602487867f2e59db3a000000000000000000000000000000000000000000000000000000008252600452fd5b905034610fbd576020600319360112610fbd576020907fffffffff0000000000000000000000000000000000000000000000000000000061418c61424b565b167faff2afbf000000000000000000000000000000000000000000000000000000008114908115614221575b81156141f7575b81156141cd575b5015158152f35b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014836141c6565b7f0e64dd2900000000000000000000000000000000000000000000000000000000811491506141bf565b7f940a154200000000000000000000000000000000000000000000000000000000811491506141b8565b600435907fffffffff00000000000000000000000000000000000000000000000000000000821682036126f457565b606435907fffffffff00000000000000000000000000000000000000000000000000000000821682036126f457565b602435907fffffffff00000000000000000000000000000000000000000000000000000000821682036126f457565b604435907fffffffff00000000000000000000000000000000000000000000000000000000821682036126f457565b6004359067ffffffffffffffff821682036126f457565b6024359067ffffffffffffffff821682036126f457565b359067ffffffffffffffff821682036126f457565b9181601f840112156126f45782359167ffffffffffffffff83116126f457602083818601950101116126f457565b600435906001600160a01b03821682036126f457565b606435906001600160a01b03821682036126f457565b602435906001600160a01b03821682036126f457565b35906001600160a01b03821682036126f457565b6060810190811067ffffffffffffffff8211176143ea57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff8211176143ea57604052565b6040810190811067ffffffffffffffff8211176143ea57604052565b60a0810190811067ffffffffffffffff8211176143ea57604052565b60e0810190811067ffffffffffffffff8211176143ea57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176143ea57604052565b67ffffffffffffffff81116143ea57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60005b8381106145175750506000910152565b8181015183820152602001614507565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f60209361456381518092818752878088019101614504565b0116010190565b359081151582036126f457565b929192614583826144ca565b916145916040519384614489565b8294818452818301116126f4578281602093846000960137010152565b9080601f830112156126f4578160206145c993359101614577565b90565b9060406003198301126126f45760043567ffffffffffffffff811681036126f457916024359067ffffffffffffffff82116126f45761460d9160040161434a565b9091565b6145c991602061462a8351604084526040840190614527565b920151906020818403910152614527565b9181601f840112156126f45782359167ffffffffffffffff83116126f4576020808501948460051b0101116126f457565b9181601f840112156126f45782359167ffffffffffffffff83116126f4576020808501948460081b0101116126f457565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b60409067ffffffffffffffff6145c99593168152816020820152019161469d565b67ffffffffffffffff81116143ea5760051b60200190565b80518210156147295760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b67ffffffffffffffff909291926147967fffffffff0000000000000000000000000000000000000000000000000000000060025460401b1685614d77565b16600052600b6020526040600020604051906147b18261446d565b549163ffffffff83169384835263ffffffff8460201c169384602085015263ffffffff8160401c169182604086015263ffffffff8260601c169081606087015261ffff8360801c169586608082015260ff61ffff8560901c16948560a084015260a01c16159060c0821591015261485e577fffffffff000000000000000000000000000000000000000000000000000000001661485357505093929190600190565b959493509160019150565b5050505092505050600090600090600090600090600090565b6040519061488482614451565b60006080838281528260208201528260408201528260608201520152565b906040516148af81614451565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff9161491a614877565b50614923614877565b50614957571660005260086020526040600020906145c961494b600261495061494b866148a2565b614e5e565b94016148a2565b169081600052600460205261497261494b60406000206148a2565b9160005260056020526145c961494b60406000206148a2565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156126f4570180359067ffffffffffffffff82116126f4576020019181360383136126f457565b356001600160a01b03811681036126f45790565b3567ffffffffffffffff811681036126f45790565b908160209103126126f457516001600160a01b03811681036126f45790565b519081151582036126f457565b9067ffffffffffffffff6145c992166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b60405190614a7b82614435565b60606020838281520152565b90600182811c92168015614ad0575b6020831014614aa157565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691614a96565b9060405191826000825492614aee84614a87565b8084529360018116908115614b5c5750600114614b15575b50614b1392500383614489565b565b90506000929192526020600020906000915b818310614b40575050906020614b139282010138614b06565b6020919350806001915483858901015201910190918492614b27565b60209350614b139592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138614b06565b67ffffffffffffffff1660005260086020526145c96004604060002001614ada565b91908110156147295760081b0190565b3580151581036126f45790565b3561ffff811681036126f45790565b3563ffffffff811681036126f45790565b359063ffffffff821682036126f457565b359061ffff821682036126f457565b91908110156147295760051b0190565b35906fffffffffffffffffffffffffffffffff821682036126f457565b91908260609103126126f457604051614c60816143ce565b6040614c8a818395614c718161456a565b8552614c7f60208201614c2b565b602086015201614c2b565b910152565b6fffffffffffffffffffffffffffffffff614ccd60408093614cb08161456a565b1515865283614cc160208301614c2b565b16602087015201614c2b565b16910152565b818110614cde575050565b60008155600101614cd3565b81810292918115918404141715614cfd57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6001600160a01b03600154163303614d4057565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b91908203918211614cfd57565b7fffffffff000000000000000000000000000000000000000000000000000000008116908115614e5957614daa816155fe565b7dffff00000000000000000000000000000000000000000000000000000000601082811c9085901c1616614e595761ffff8360e01c168015918215614e48575b5050614df4575050565b7fffffffff0000000000000000000000000000000000000000000000000000000092507fdf63778f000000000000000000000000000000000000000000000000000000006000526004521660245260446000fd5b60e01c61ffff161090503880614dea565b505050565b614e66614877565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff8083511691614ec36020850193614ebd614eb063ffffffff87511642614d6a565b8560808901511690614cea565b906155f1565b80821015614edc57505b16825263ffffffff4216905290565b9050614ecd565b80518015614f5357602003614f155780516020828101918301839003126126f457519060ff8211614f15575060ff1690565b611f93906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190614527565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff8211614cfd57565b60ff16604d8111614cfd57600a0a90565b8115614fa8570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff8116928284146150bf57828411615095579061501c91614f79565b91604d60ff841611801561507a575b6150445750509061503e6145c992614f8d565b90614cea565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b5061508483614f8d565b8015614fa85760001904841161502b565b61509e91614f79565b91604d60ff841611615044575050906150b96145c992614f8d565b90614f9e565b5050505090565b908051156152ca5767ffffffffffffffff815160208301209216918260005260086020526150fb816005604060002001615b3a565b156152865760005260096020526040600020815167ffffffffffffffff81116143ea576151288254614a87565b601f8111615254575b506020601f82116001146151ac5791615186827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea959361519c956000916151a1575b506000198260011b9260031b1c19161790565b9055604051918291602083526020830190614527565b0390a2565b905084015138615173565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b81811061523c57509261519c9492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610615223575b5050811b019055611c87565b85015160001960f88460031b161c191690553880615217565b9192602060018192868a0151815501940192016151dc565b61528090836000526020600020601f840160051c81019160208510610ef457601f0160051c0190614cd3565b38615131565b5090611f936040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190614527565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b815191929115615476576fffffffffffffffffffffffffffffffff6040840151166fffffffffffffffffffffffffffffffff6020850151161061541357614b1391925b805182547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff0000000000000000000000000000000000000000161782556020810151825460409290920151608090811b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff9290921691821760018501557fffffffffffffffffffffffff0000000000000000000000000000000000000000909216174290911b73ffffffff0000000000000000000000000000000016179055565b606483615474604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff60408401511615801590615505575b6154a457614b139192615337565b606483615474604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff6020840151161515615496565b6040517fa9059cbb0000000000000000000000000000000000000000000000000000000060208083019182526001600160a01b03949094166024830152604480830195909552938152909260009161557d606482614489565b519082855af1156155e5576000513d6155dc57506001600160a01b0381163b155b6155a55750565b6001600160a01b03907f5274afe7000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b6001141561559e565b6040513d6000823e3d90fd5b91908201809211614cfd57565b7fffffffff0000000000000000000000000000000000000000000000000000000081169081156156d5577dffff000000000000000000000000000000000000000000000000000000008116156156cc5760ff60015b169060f01c80615696575b506001036156695750565b7fc512f96c0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60005b601081106156a7575061565e565b6001811b82166156ba575b600101615699565b9160018101809111614cfd57916156b2565b60ff6000615653565b5050565b9167ffffffffffffffff7f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c92169283600052600860205261572281836002604060002001615c10565b604080516001600160a01b0390921682526020820192909252908190810161519c565b91909167ffffffffffffffff83169283600052600560205260ff60406000205460a01c16156157aa5750907fc6735cd4fa2bbe7b203b1682936e6ee61bc1702464bbbd12abb6630229d9a5f99183600052600560205261572281836040600020615c10565b90614b1393506156d9565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156126f457016020813591019167ffffffffffffffff82116126f45781360383136126f457565b9167ffffffffffffffff7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894492169283600052600860205261572281836040600020615c10565b91909167ffffffffffffffff83169283600052600460205260ff60406000205460a01c16156158b05750907f28d6c52e2b0b7587b0d195539fbe6af984b28791aca4d2cc0844244e38bce29e9183600052600460205261572281836040600020615c10565b90614b139350615805565b906040519182815491828252602082019060005260206000209260005b8181106158ed575050614b1392500383614489565b84548352600194850194879450602090930192016158d8565b80548210156147295760005260206000200190600090565b6000818152600760205260409020548015615a17576000198101818111614cfd57600654906000198201918211614cfd578181036159c6575b50505060065480156159975760001901615972816006615906565b60001982549160031b1b19169055600655600052600760205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6159ff6159d76159e8936006615906565b90549060031b1c9283926006615906565b81939154906000199060031b92831b921b19161790565b90556000526007602052604060002055388080615957565b5050600090565b9060018201918160005282602052604060002054801515600014615ad1576000198101818111614cfd578254906000198201918211614cfd57818103615a9a575b50505080548015615997576000190190615a798282615906565b60001982549160031b1b191690555560005260205260006040812055600190565b615aba615aaa6159e89386615906565b90549060031b1c92839286615906565b905560005283602052604060002055388080615a5f565b50505050600090565b80600052600760205260406000205415600014615b3457600654680100000000000000008110156143ea57615b1b6159e88260018594016006556006615906565b9055600654906000526007602052604060002055600190565b50600090565b6000828152600182016020526040902054615a1757805490680100000000000000008210156143ea5782615b786159e8846001809601855584615906565b905580549260005201602052604060002055600190565b906127109167ffffffffffffffff615ba9602083016149f0565b166000908152600b60205260409020917fffffffff000000000000000000000000000000000000000000000000000000001615615bfa57606061ffff615bf6935460901c16910135614cea565b0490565b606061ffff615bf6935460801c16910135614cea565b9182549060ff8260a01c16158015615e61575b615e5b576fffffffffffffffffffffffffffffffff82169160018501908154615c6863ffffffff6fffffffffffffffffffffffffffffffff83169360801c1642614d6a565b9081615dbd575b5050848110615d7e5750838310615cc9575050615c9e6fffffffffffffffffffffffffffffffff928392614d6a565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b9190915460801c928315615d3d5781615ce191614d6a565b92600019810190808211614cfd57615d04615d09926001600160a01b03966155f1565b614f9e565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b6001600160a01b0383837fd0c8d23a000000000000000000000000000000000000000000000000000000006000526000196004526024521660445260646000fd5b82856001600160a01b03927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611615e3157615dd892614ebd9160801c90614cea565b80841015615e2c5750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178655923880615c6f565b615de3565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b508215615c2356fea164736f6c634300081a000a",
}

var LombardTokenPoolABI = LombardTokenPoolMetaData.ABI

var LombardTokenPoolBin = LombardTokenPoolMetaData.Bin

func DeployLombardTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, verifier common.Address, bridge common.Address, adapter common.Address, advancedPoolHooks common.Address, rmnProxy common.Address, router common.Address, fallbackDecimals uint8) (common.Address, *types.Transaction, *LombardTokenPool, error) {
	parsed, err := LombardTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LombardTokenPoolBin), backend, token, verifier, bridge, adapter, advancedPoolHooks, rmnProxy, router, fallbackDecimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LombardTokenPool{address: address, abi: *parsed, LombardTokenPoolCaller: LombardTokenPoolCaller{contract: contract}, LombardTokenPoolTransactor: LombardTokenPoolTransactor{contract: contract}, LombardTokenPoolFilterer: LombardTokenPoolFilterer{contract: contract}}, nil
}

type LombardTokenPool struct {
	address common.Address
	abi     abi.ABI
	LombardTokenPoolCaller
	LombardTokenPoolTransactor
	LombardTokenPoolFilterer
}

type LombardTokenPoolCaller struct {
	contract *bind.BoundContract
}

type LombardTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type LombardTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type LombardTokenPoolSession struct {
	Contract     *LombardTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type LombardTokenPoolCallerSession struct {
	Contract *LombardTokenPoolCaller
	CallOpts bind.CallOpts
}

type LombardTokenPoolTransactorSession struct {
	Contract     *LombardTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type LombardTokenPoolRaw struct {
	Contract *LombardTokenPool
}

type LombardTokenPoolCallerRaw struct {
	Contract *LombardTokenPoolCaller
}

type LombardTokenPoolTransactorRaw struct {
	Contract *LombardTokenPoolTransactor
}

func NewLombardTokenPool(address common.Address, backend bind.ContractBackend) (*LombardTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(LombardTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindLombardTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPool{address: address, abi: abi, LombardTokenPoolCaller: LombardTokenPoolCaller{contract: contract}, LombardTokenPoolTransactor: LombardTokenPoolTransactor{contract: contract}, LombardTokenPoolFilterer: LombardTokenPoolFilterer{contract: contract}}, nil
}

func NewLombardTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*LombardTokenPoolCaller, error) {
	contract, err := bindLombardTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolCaller{contract: contract}, nil
}

func NewLombardTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*LombardTokenPoolTransactor, error) {
	contract, err := bindLombardTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolTransactor{contract: contract}, nil
}

func NewLombardTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*LombardTokenPoolFilterer, error) {
	contract, err := bindLombardTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolFilterer{contract: contract}, nil
}

func bindLombardTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LombardTokenPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_LombardTokenPool *LombardTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LombardTokenPool.Contract.LombardTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_LombardTokenPool *LombardTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.LombardTokenPoolTransactor.contract.Transfer(opts)
}

func (_LombardTokenPool *LombardTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.LombardTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_LombardTokenPool *LombardTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LombardTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_LombardTokenPool *LombardTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.contract.Transfer(opts)
}

func (_LombardTokenPool *LombardTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_LombardTokenPool *LombardTokenPoolCaller) GetAdvancedPoolHooks(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getAdvancedPoolHooks")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) GetAdvancedPoolHooks() (common.Address, error) {
	return _LombardTokenPool.Contract.GetAdvancedPoolHooks(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetAdvancedPoolHooks() (common.Address, error) {
	return _LombardTokenPool.Contract.GetAdvancedPoolHooks(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCaller) GetAllowedFinalityConfig(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getAllowedFinalityConfig")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) GetAllowedFinalityConfig() ([4]byte, error) {
	return _LombardTokenPool.Contract.GetAllowedFinalityConfig(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetAllowedFinalityConfig() ([4]byte, error) {
	return _LombardTokenPool.Contract.GetAllowedFinalityConfig(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCaller) GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64, fastFinality bool) (GetCurrentRateLimiterState,

	error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getCurrentRateLimiterState", remoteChainSelector, fastFinality)

	outstruct := new(GetCurrentRateLimiterState)
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutboundRateLimiterState = *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)
	outstruct.InboundRateLimiterState = *abi.ConvertType(out[1], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return *outstruct, err

}

func (_LombardTokenPool *LombardTokenPoolSession) GetCurrentRateLimiterState(remoteChainSelector uint64, fastFinality bool) (GetCurrentRateLimiterState,

	error) {
	return _LombardTokenPool.Contract.GetCurrentRateLimiterState(&_LombardTokenPool.CallOpts, remoteChainSelector, fastFinality)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetCurrentRateLimiterState(remoteChainSelector uint64, fastFinality bool) (GetCurrentRateLimiterState,

	error) {
	return _LombardTokenPool.Contract.GetCurrentRateLimiterState(&_LombardTokenPool.CallOpts, remoteChainSelector, fastFinality)
}

func (_LombardTokenPool *LombardTokenPoolCaller) GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

	error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getDynamicConfig")

	outstruct := new(GetDynamicConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Router = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RateLimitAdmin = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.FeeAdmin = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

func (_LombardTokenPool *LombardTokenPoolSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _LombardTokenPool.Contract.GetDynamicConfig(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _LombardTokenPool.Contract.GetDynamicConfig(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCaller) GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, requestedFinalityConfig [4]byte, arg5 []byte) (GetFee,

	error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getFee", arg0, destChainSelector, arg2, arg3, requestedFinalityConfig, arg5)

	outstruct := new(GetFee)
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeeUSDCents = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DestGasOverhead = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.DestBytesOverhead = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.TokenFeeBps = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.IsEnabled = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

func (_LombardTokenPool *LombardTokenPoolSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, requestedFinalityConfig [4]byte, arg5 []byte) (GetFee,

	error) {
	return _LombardTokenPool.Contract.GetFee(&_LombardTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, requestedFinalityConfig, arg5)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, requestedFinalityConfig [4]byte, arg5 []byte) (GetFee,

	error) {
	return _LombardTokenPool.Contract.GetFee(&_LombardTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, requestedFinalityConfig, arg5)
}

func (_LombardTokenPool *LombardTokenPoolCaller) GetLombardConfig(opts *bind.CallOpts) (GetLombardConfig,

	error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getLombardConfig")

	outstruct := new(GetLombardConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.VerifierResolver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Bridge = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenAdapter = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

func (_LombardTokenPool *LombardTokenPoolSession) GetLombardConfig() (GetLombardConfig,

	error) {
	return _LombardTokenPool.Contract.GetLombardConfig(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetLombardConfig() (GetLombardConfig,

	error) {
	return _LombardTokenPool.Contract.GetLombardConfig(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCaller) GetPath(opts *bind.CallOpts, remoteChainSelector uint64) (LombardTokenPoolPath, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getPath", remoteChainSelector)

	if err != nil {
		return *new(LombardTokenPoolPath), err
	}

	out0 := *abi.ConvertType(out[0], new(LombardTokenPoolPath)).(*LombardTokenPoolPath)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) GetPath(remoteChainSelector uint64) (LombardTokenPoolPath, error) {
	return _LombardTokenPool.Contract.GetPath(&_LombardTokenPool.CallOpts, remoteChainSelector)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetPath(remoteChainSelector uint64) (LombardTokenPoolPath, error) {
	return _LombardTokenPool.Contract.GetPath(&_LombardTokenPool.CallOpts, remoteChainSelector)
}

func (_LombardTokenPool *LombardTokenPoolCaller) GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getRemotePools", remoteChainSelector)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _LombardTokenPool.Contract.GetRemotePools(&_LombardTokenPool.CallOpts, remoteChainSelector)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _LombardTokenPool.Contract.GetRemotePools(&_LombardTokenPool.CallOpts, remoteChainSelector)
}

func (_LombardTokenPool *LombardTokenPoolCaller) GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getRemoteToken", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _LombardTokenPool.Contract.GetRemoteToken(&_LombardTokenPool.CallOpts, remoteChainSelector)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _LombardTokenPool.Contract.GetRemoteToken(&_LombardTokenPool.CallOpts, remoteChainSelector)
}

func (_LombardTokenPool *LombardTokenPoolCaller) GetRequiredCCVs(opts *bind.CallOpts, arg0 common.Address, arg1 uint64, arg2 *big.Int, arg3 [4]byte, arg4 []byte, arg5 uint8) ([]common.Address, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getRequiredCCVs", arg0, arg1, arg2, arg3, arg4, arg5)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) GetRequiredCCVs(arg0 common.Address, arg1 uint64, arg2 *big.Int, arg3 [4]byte, arg4 []byte, arg5 uint8) ([]common.Address, error) {
	return _LombardTokenPool.Contract.GetRequiredCCVs(&_LombardTokenPool.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetRequiredCCVs(arg0 common.Address, arg1 uint64, arg2 *big.Int, arg3 [4]byte, arg4 []byte, arg5 uint8) ([]common.Address, error) {
	return _LombardTokenPool.Contract.GetRequiredCCVs(&_LombardTokenPool.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

func (_LombardTokenPool *LombardTokenPoolCaller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) GetRmnProxy() (common.Address, error) {
	return _LombardTokenPool.Contract.GetRmnProxy(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetRmnProxy() (common.Address, error) {
	return _LombardTokenPool.Contract.GetRmnProxy(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) GetSupportedChains() ([]uint64, error) {
	return _LombardTokenPool.Contract.GetSupportedChains(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetSupportedChains() ([]uint64, error) {
	return _LombardTokenPool.Contract.GetSupportedChains(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) GetToken() (common.Address, error) {
	return _LombardTokenPool.Contract.GetToken(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _LombardTokenPool.Contract.GetToken(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCaller) GetTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) GetTokenDecimals() (uint8, error) {
	return _LombardTokenPool.Contract.GetTokenDecimals(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetTokenDecimals() (uint8, error) {
	return _LombardTokenPool.Contract.GetTokenDecimals(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 [4]byte, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getTokenTransferFeeConfig", arg0, destChainSelector, arg2, arg3)

	if err != nil {
		return *new(IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolV2TokenTransferFeeConfig)).(*IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 [4]byte, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _LombardTokenPool.Contract.GetTokenTransferFeeConfig(&_LombardTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 [4]byte, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _LombardTokenPool.Contract.GetTokenTransferFeeConfig(&_LombardTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3)
}

func (_LombardTokenPool *LombardTokenPoolCaller) IBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "i_bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) IBridge() (common.Address, error) {
	return _LombardTokenPool.Contract.IBridge(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) IBridge() (common.Address, error) {
	return _LombardTokenPool.Contract.IBridge(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCaller) IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "isRemotePool", remoteChainSelector, remotePoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _LombardTokenPool.Contract.IsRemotePool(&_LombardTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _LombardTokenPool.Contract.IsRemotePool(&_LombardTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_LombardTokenPool *LombardTokenPoolCaller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _LombardTokenPool.Contract.IsSupportedChain(&_LombardTokenPool.CallOpts, remoteChainSelector)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _LombardTokenPool.Contract.IsSupportedChain(&_LombardTokenPool.CallOpts, remoteChainSelector)
}

func (_LombardTokenPool *LombardTokenPoolCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) IsSupportedToken(token common.Address) (bool, error) {
	return _LombardTokenPool.Contract.IsSupportedToken(&_LombardTokenPool.CallOpts, token)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _LombardTokenPool.Contract.IsSupportedToken(&_LombardTokenPool.CallOpts, token)
}

func (_LombardTokenPool *LombardTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) Owner() (common.Address, error) {
	return _LombardTokenPool.Contract.Owner(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) Owner() (common.Address, error) {
	return _LombardTokenPool.Contract.Owner(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LombardTokenPool.Contract.SupportsInterface(&_LombardTokenPool.CallOpts, interfaceId)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LombardTokenPool.Contract.SupportsInterface(&_LombardTokenPool.CallOpts, interfaceId)
}

func (_LombardTokenPool *LombardTokenPoolCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) TypeAndVersion() (string, error) {
	return _LombardTokenPool.Contract.TypeAndVersion(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) TypeAndVersion() (string, error) {
	return _LombardTokenPool.Contract.TypeAndVersion(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_LombardTokenPool *LombardTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _LombardTokenPool.Contract.AcceptOwnership(&_LombardTokenPool.TransactOpts)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _LombardTokenPool.Contract.AcceptOwnership(&_LombardTokenPool.TransactOpts)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "addRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_LombardTokenPool *LombardTokenPoolSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.AddRemotePool(&_LombardTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.AddRemotePool(&_LombardTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "applyChainUpdates", remoteChainSelectorsToRemove, chainsToAdd)
}

func (_LombardTokenPool *LombardTokenPoolSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.ApplyChainUpdates(&_LombardTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.ApplyChainUpdates(&_LombardTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "applyTokenTransferFeeConfigUpdates", tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_LombardTokenPool *LombardTokenPoolSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_LombardTokenPool.TransactOpts, tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_LombardTokenPool.TransactOpts, tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_LombardTokenPool *LombardTokenPoolSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.LockOrBurn(&_LombardTokenPool.TransactOpts, lockOrBurnIn)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.LockOrBurn(&_LombardTokenPool.TransactOpts, lockOrBurnIn)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, requestedFinalityConfig [4]byte, tokenArgs []byte) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, requestedFinalityConfig, tokenArgs)
}

func (_LombardTokenPool *LombardTokenPoolSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, requestedFinalityConfig [4]byte, tokenArgs []byte) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.LockOrBurn0(&_LombardTokenPool.TransactOpts, lockOrBurnIn, requestedFinalityConfig, tokenArgs)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, requestedFinalityConfig [4]byte, tokenArgs []byte) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.LockOrBurn0(&_LombardTokenPool.TransactOpts, lockOrBurnIn, requestedFinalityConfig, tokenArgs)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, requestedFinalityConfig [4]byte) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "releaseOrMint", releaseOrMintIn, requestedFinalityConfig)
}

func (_LombardTokenPool *LombardTokenPoolSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1, requestedFinalityConfig [4]byte) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.ReleaseOrMint(&_LombardTokenPool.TransactOpts, releaseOrMintIn, requestedFinalityConfig)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1, requestedFinalityConfig [4]byte) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.ReleaseOrMint(&_LombardTokenPool.TransactOpts, releaseOrMintIn, requestedFinalityConfig)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn)
}

func (_LombardTokenPool *LombardTokenPoolSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.ReleaseOrMint0(&_LombardTokenPool.TransactOpts, releaseOrMintIn)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.ReleaseOrMint0(&_LombardTokenPool.TransactOpts, releaseOrMintIn)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) RemovePath(opts *bind.TransactOpts, remoteChainSelector uint64) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "removePath", remoteChainSelector)
}

func (_LombardTokenPool *LombardTokenPoolSession) RemovePath(remoteChainSelector uint64) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.RemovePath(&_LombardTokenPool.TransactOpts, remoteChainSelector)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) RemovePath(remoteChainSelector uint64) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.RemovePath(&_LombardTokenPool.TransactOpts, remoteChainSelector)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "removeRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_LombardTokenPool *LombardTokenPoolSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.RemoveRemotePool(&_LombardTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.RemoveRemotePool(&_LombardTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) SetAllowedFinalityConfig(opts *bind.TransactOpts, allowedFinality [4]byte) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "setAllowedFinalityConfig", allowedFinality)
}

func (_LombardTokenPool *LombardTokenPoolSession) SetAllowedFinalityConfig(allowedFinality [4]byte) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.SetAllowedFinalityConfig(&_LombardTokenPool.TransactOpts, allowedFinality)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) SetAllowedFinalityConfig(allowedFinality [4]byte) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.SetAllowedFinalityConfig(&_LombardTokenPool.TransactOpts, allowedFinality)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) SetDynamicConfig(opts *bind.TransactOpts, router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "setDynamicConfig", router, rateLimitAdmin, feeAdmin)
}

func (_LombardTokenPool *LombardTokenPoolSession) SetDynamicConfig(router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.SetDynamicConfig(&_LombardTokenPool.TransactOpts, router, rateLimitAdmin, feeAdmin)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) SetDynamicConfig(router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.SetDynamicConfig(&_LombardTokenPool.TransactOpts, router, rateLimitAdmin, feeAdmin)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) SetPath(opts *bind.TransactOpts, remoteChainSelector uint64, lChainId [32]byte, allowedCaller []byte, remoteAdapter [32]byte) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "setPath", remoteChainSelector, lChainId, allowedCaller, remoteAdapter)
}

func (_LombardTokenPool *LombardTokenPoolSession) SetPath(remoteChainSelector uint64, lChainId [32]byte, allowedCaller []byte, remoteAdapter [32]byte) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.SetPath(&_LombardTokenPool.TransactOpts, remoteChainSelector, lChainId, allowedCaller, remoteAdapter)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) SetPath(remoteChainSelector uint64, lChainId [32]byte, allowedCaller []byte, remoteAdapter [32]byte) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.SetPath(&_LombardTokenPool.TransactOpts, remoteChainSelector, lChainId, allowedCaller, remoteAdapter)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) SetRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "setRateLimitConfig", rateLimitConfigArgs)
}

func (_LombardTokenPool *LombardTokenPoolSession) SetRateLimitConfig(rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.SetRateLimitConfig(&_LombardTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) SetRateLimitConfig(rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.SetRateLimitConfig(&_LombardTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_LombardTokenPool *LombardTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.TransferOwnership(&_LombardTokenPool.TransactOpts, to)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.TransferOwnership(&_LombardTokenPool.TransactOpts, to)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) UpdateAdvancedPoolHooks(opts *bind.TransactOpts, newHook common.Address) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "updateAdvancedPoolHooks", newHook)
}

func (_LombardTokenPool *LombardTokenPoolSession) UpdateAdvancedPoolHooks(newHook common.Address) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.UpdateAdvancedPoolHooks(&_LombardTokenPool.TransactOpts, newHook)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) UpdateAdvancedPoolHooks(newHook common.Address) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.UpdateAdvancedPoolHooks(&_LombardTokenPool.TransactOpts, newHook)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "withdrawFeeTokens", feeTokens, recipient)
}

func (_LombardTokenPool *LombardTokenPoolSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.WithdrawFeeTokens(&_LombardTokenPool.TransactOpts, feeTokens, recipient)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.WithdrawFeeTokens(&_LombardTokenPool.TransactOpts, feeTokens, recipient)
}

type LombardTokenPoolAdvancedPoolHooksUpdatedIterator struct {
	Event *LombardTokenPoolAdvancedPoolHooksUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolAdvancedPoolHooksUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolAdvancedPoolHooksUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LombardTokenPoolAdvancedPoolHooksUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LombardTokenPoolAdvancedPoolHooksUpdatedIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolAdvancedPoolHooksUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolAdvancedPoolHooksUpdated struct {
	OldHook common.Address
	NewHook common.Address
	Raw     types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterAdvancedPoolHooksUpdated(opts *bind.FilterOpts) (*LombardTokenPoolAdvancedPoolHooksUpdatedIterator, error) {

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "AdvancedPoolHooksUpdated")
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolAdvancedPoolHooksUpdatedIterator{contract: _LombardTokenPool.contract, event: "AdvancedPoolHooksUpdated", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchAdvancedPoolHooksUpdated(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolAdvancedPoolHooksUpdated) (event.Subscription, error) {

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "AdvancedPoolHooksUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolAdvancedPoolHooksUpdated)
				if err := _LombardTokenPool.contract.UnpackLog(event, "AdvancedPoolHooksUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseAdvancedPoolHooksUpdated(log types.Log) (*LombardTokenPoolAdvancedPoolHooksUpdated, error) {
	event := new(LombardTokenPoolAdvancedPoolHooksUpdated)
	if err := _LombardTokenPool.contract.UnpackLog(event, "AdvancedPoolHooksUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolChainAddedIterator struct {
	Event *LombardTokenPoolChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolChainAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LombardTokenPoolChainAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LombardTokenPoolChainAddedIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolChainAdded struct {
	RemoteChainSelector       uint64
	RemoteToken               []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterChainAdded(opts *bind.FilterOpts) (*LombardTokenPoolChainAddedIterator, error) {

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolChainAddedIterator{contract: _LombardTokenPool.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolChainAdded) (event.Subscription, error) {

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolChainAdded)
				if err := _LombardTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseChainAdded(log types.Log) (*LombardTokenPoolChainAdded, error) {
	event := new(LombardTokenPoolChainAdded)
	if err := _LombardTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolChainRemovedIterator struct {
	Event *LombardTokenPoolChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolChainRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LombardTokenPoolChainRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LombardTokenPoolChainRemovedIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterChainRemoved(opts *bind.FilterOpts) (*LombardTokenPoolChainRemovedIterator, error) {

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolChainRemovedIterator{contract: _LombardTokenPool.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolChainRemoved) (event.Subscription, error) {

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolChainRemoved)
				if err := _LombardTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseChainRemoved(log types.Log) (*LombardTokenPoolChainRemoved, error) {
	event := new(LombardTokenPoolChainRemoved)
	if err := _LombardTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolDynamicConfigSetIterator struct {
	Event *LombardTokenPoolDynamicConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolDynamicConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolDynamicConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LombardTokenPoolDynamicConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LombardTokenPoolDynamicConfigSetIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolDynamicConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolDynamicConfigSet struct {
	Router         common.Address
	RateLimitAdmin common.Address
	FeeAdmin       common.Address
	Raw            types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterDynamicConfigSet(opts *bind.FilterOpts) (*LombardTokenPoolDynamicConfigSetIterator, error) {

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolDynamicConfigSetIterator{contract: _LombardTokenPool.contract, event: "DynamicConfigSet", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolDynamicConfigSet) (event.Subscription, error) {

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolDynamicConfigSet)
				if err := _LombardTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseDynamicConfigSet(log types.Log) (*LombardTokenPoolDynamicConfigSet, error) {
	event := new(LombardTokenPoolDynamicConfigSet)
	if err := _LombardTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolFastFinalityInboundRateLimitConsumedIterator struct {
	Event *LombardTokenPoolFastFinalityInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolFastFinalityInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolFastFinalityInboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LombardTokenPoolFastFinalityInboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LombardTokenPoolFastFinalityInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolFastFinalityInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolFastFinalityInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterFastFinalityInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardTokenPoolFastFinalityInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "FastFinalityInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolFastFinalityInboundRateLimitConsumedIterator{contract: _LombardTokenPool.contract, event: "FastFinalityInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchFastFinalityInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolFastFinalityInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "FastFinalityInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolFastFinalityInboundRateLimitConsumed)
				if err := _LombardTokenPool.contract.UnpackLog(event, "FastFinalityInboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseFastFinalityInboundRateLimitConsumed(log types.Log) (*LombardTokenPoolFastFinalityInboundRateLimitConsumed, error) {
	event := new(LombardTokenPoolFastFinalityInboundRateLimitConsumed)
	if err := _LombardTokenPool.contract.UnpackLog(event, "FastFinalityInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolFastFinalityOutboundRateLimitConsumedIterator struct {
	Event *LombardTokenPoolFastFinalityOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolFastFinalityOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolFastFinalityOutboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LombardTokenPoolFastFinalityOutboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LombardTokenPoolFastFinalityOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolFastFinalityOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolFastFinalityOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterFastFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardTokenPoolFastFinalityOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "FastFinalityOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolFastFinalityOutboundRateLimitConsumedIterator{contract: _LombardTokenPool.contract, event: "FastFinalityOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchFastFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolFastFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "FastFinalityOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolFastFinalityOutboundRateLimitConsumed)
				if err := _LombardTokenPool.contract.UnpackLog(event, "FastFinalityOutboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseFastFinalityOutboundRateLimitConsumed(log types.Log) (*LombardTokenPoolFastFinalityOutboundRateLimitConsumed, error) {
	event := new(LombardTokenPoolFastFinalityOutboundRateLimitConsumed)
	if err := _LombardTokenPool.contract.UnpackLog(event, "FastFinalityOutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolFeeTokenWithdrawnIterator struct {
	Event *LombardTokenPoolFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolFeeTokenWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LombardTokenPoolFeeTokenWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LombardTokenPoolFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolFeeTokenWithdrawn struct {
	Receiver common.Address
	FeeToken common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*LombardTokenPoolFeeTokenWithdrawnIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolFeeTokenWithdrawnIterator{contract: _LombardTokenPool.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolFeeTokenWithdrawn)
				if err := _LombardTokenPool.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseFeeTokenWithdrawn(log types.Log) (*LombardTokenPoolFeeTokenWithdrawn, error) {
	event := new(LombardTokenPoolFeeTokenWithdrawn)
	if err := _LombardTokenPool.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolFinalityConfigSetIterator struct {
	Event *LombardTokenPoolFinalityConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolFinalityConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolFinalityConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LombardTokenPoolFinalityConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LombardTokenPoolFinalityConfigSetIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolFinalityConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolFinalityConfigSet struct {
	AllowedFinality [4]byte
	Raw             types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterFinalityConfigSet(opts *bind.FilterOpts) (*LombardTokenPoolFinalityConfigSetIterator, error) {

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "FinalityConfigSet")
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolFinalityConfigSetIterator{contract: _LombardTokenPool.contract, event: "FinalityConfigSet", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchFinalityConfigSet(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolFinalityConfigSet) (event.Subscription, error) {

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "FinalityConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolFinalityConfigSet)
				if err := _LombardTokenPool.contract.UnpackLog(event, "FinalityConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseFinalityConfigSet(log types.Log) (*LombardTokenPoolFinalityConfigSet, error) {
	event := new(LombardTokenPoolFinalityConfigSet)
	if err := _LombardTokenPool.contract.UnpackLog(event, "FinalityConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolInboundRateLimitConsumedIterator struct {
	Event *LombardTokenPoolInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolInboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LombardTokenPoolInboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LombardTokenPoolInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardTokenPoolInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolInboundRateLimitConsumedIterator{contract: _LombardTokenPool.contract, event: "InboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolInboundRateLimitConsumed)
				if err := _LombardTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseInboundRateLimitConsumed(log types.Log) (*LombardTokenPoolInboundRateLimitConsumed, error) {
	event := new(LombardTokenPoolInboundRateLimitConsumed)
	if err := _LombardTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolLockedOrBurnedIterator struct {
	Event *LombardTokenPoolLockedOrBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolLockedOrBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolLockedOrBurned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LombardTokenPoolLockedOrBurned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LombardTokenPoolLockedOrBurnedIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolLockedOrBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolLockedOrBurned struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardTokenPoolLockedOrBurnedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolLockedOrBurnedIterator{contract: _LombardTokenPool.contract, event: "LockedOrBurned", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolLockedOrBurned)
				if err := _LombardTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseLockedOrBurned(log types.Log) (*LombardTokenPoolLockedOrBurned, error) {
	event := new(LombardTokenPoolLockedOrBurned)
	if err := _LombardTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolLombardConfigurationSetIterator struct {
	Event *LombardTokenPoolLombardConfigurationSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolLombardConfigurationSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolLombardConfigurationSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LombardTokenPoolLombardConfigurationSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LombardTokenPoolLombardConfigurationSetIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolLombardConfigurationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolLombardConfigurationSet struct {
	Verifier     common.Address
	Bridge       common.Address
	TokenAdapter common.Address
	Raw          types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterLombardConfigurationSet(opts *bind.FilterOpts, verifier []common.Address, bridge []common.Address, tokenAdapter []common.Address) (*LombardTokenPoolLombardConfigurationSetIterator, error) {

	var verifierRule []interface{}
	for _, verifierItem := range verifier {
		verifierRule = append(verifierRule, verifierItem)
	}
	var bridgeRule []interface{}
	for _, bridgeItem := range bridge {
		bridgeRule = append(bridgeRule, bridgeItem)
	}
	var tokenAdapterRule []interface{}
	for _, tokenAdapterItem := range tokenAdapter {
		tokenAdapterRule = append(tokenAdapterRule, tokenAdapterItem)
	}

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "LombardConfigurationSet", verifierRule, bridgeRule, tokenAdapterRule)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolLombardConfigurationSetIterator{contract: _LombardTokenPool.contract, event: "LombardConfigurationSet", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchLombardConfigurationSet(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolLombardConfigurationSet, verifier []common.Address, bridge []common.Address, tokenAdapter []common.Address) (event.Subscription, error) {

	var verifierRule []interface{}
	for _, verifierItem := range verifier {
		verifierRule = append(verifierRule, verifierItem)
	}
	var bridgeRule []interface{}
	for _, bridgeItem := range bridge {
		bridgeRule = append(bridgeRule, bridgeItem)
	}
	var tokenAdapterRule []interface{}
	for _, tokenAdapterItem := range tokenAdapter {
		tokenAdapterRule = append(tokenAdapterRule, tokenAdapterItem)
	}

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "LombardConfigurationSet", verifierRule, bridgeRule, tokenAdapterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolLombardConfigurationSet)
				if err := _LombardTokenPool.contract.UnpackLog(event, "LombardConfigurationSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseLombardConfigurationSet(log types.Log) (*LombardTokenPoolLombardConfigurationSet, error) {
	event := new(LombardTokenPoolLombardConfigurationSet)
	if err := _LombardTokenPool.contract.UnpackLog(event, "LombardConfigurationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolOutboundRateLimitConsumedIterator struct {
	Event *LombardTokenPoolOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolOutboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LombardTokenPoolOutboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LombardTokenPoolOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardTokenPoolOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolOutboundRateLimitConsumedIterator{contract: _LombardTokenPool.contract, event: "OutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolOutboundRateLimitConsumed)
				if err := _LombardTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseOutboundRateLimitConsumed(log types.Log) (*LombardTokenPoolOutboundRateLimitConsumed, error) {
	event := new(LombardTokenPoolOutboundRateLimitConsumed)
	if err := _LombardTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolOwnershipTransferRequestedIterator struct {
	Event *LombardTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolOwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LombardTokenPoolOwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LombardTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LombardTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolOwnershipTransferRequestedIterator{contract: _LombardTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolOwnershipTransferRequested)
				if err := _LombardTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*LombardTokenPoolOwnershipTransferRequested, error) {
	event := new(LombardTokenPoolOwnershipTransferRequested)
	if err := _LombardTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolOwnershipTransferredIterator struct {
	Event *LombardTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LombardTokenPoolOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LombardTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LombardTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolOwnershipTransferredIterator{contract: _LombardTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolOwnershipTransferred)
				if err := _LombardTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*LombardTokenPoolOwnershipTransferred, error) {
	event := new(LombardTokenPoolOwnershipTransferred)
	if err := _LombardTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolPathRemovedIterator struct {
	Event *LombardTokenPoolPathRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolPathRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolPathRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LombardTokenPoolPathRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LombardTokenPoolPathRemovedIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolPathRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolPathRemoved struct {
	RemoteChainSelector uint64
	LChainId            [32]byte
	AllowedCaller       [32]byte
	RemoteAdapter       [32]byte
	Raw                 types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterPathRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64, lChainId [][32]byte) (*LombardTokenPoolPathRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}
	var lChainIdRule []interface{}
	for _, lChainIdItem := range lChainId {
		lChainIdRule = append(lChainIdRule, lChainIdItem)
	}

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "PathRemoved", remoteChainSelectorRule, lChainIdRule)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolPathRemovedIterator{contract: _LombardTokenPool.contract, event: "PathRemoved", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchPathRemoved(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolPathRemoved, remoteChainSelector []uint64, lChainId [][32]byte) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}
	var lChainIdRule []interface{}
	for _, lChainIdItem := range lChainId {
		lChainIdRule = append(lChainIdRule, lChainIdItem)
	}

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "PathRemoved", remoteChainSelectorRule, lChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolPathRemoved)
				if err := _LombardTokenPool.contract.UnpackLog(event, "PathRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) ParsePathRemoved(log types.Log) (*LombardTokenPoolPathRemoved, error) {
	event := new(LombardTokenPoolPathRemoved)
	if err := _LombardTokenPool.contract.UnpackLog(event, "PathRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolPathSetIterator struct {
	Event *LombardTokenPoolPathSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolPathSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolPathSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LombardTokenPoolPathSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LombardTokenPoolPathSetIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolPathSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolPathSet struct {
	RemoteChainSelector uint64
	LChainId            [32]byte
	AllowedCaller       [32]byte
	RemoteAdapter       [32]byte
	Raw                 types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterPathSet(opts *bind.FilterOpts, remoteChainSelector []uint64, lChainId [][32]byte) (*LombardTokenPoolPathSetIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}
	var lChainIdRule []interface{}
	for _, lChainIdItem := range lChainId {
		lChainIdRule = append(lChainIdRule, lChainIdItem)
	}

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "PathSet", remoteChainSelectorRule, lChainIdRule)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolPathSetIterator{contract: _LombardTokenPool.contract, event: "PathSet", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchPathSet(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolPathSet, remoteChainSelector []uint64, lChainId [][32]byte) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}
	var lChainIdRule []interface{}
	for _, lChainIdItem := range lChainId {
		lChainIdRule = append(lChainIdRule, lChainIdItem)
	}

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "PathSet", remoteChainSelectorRule, lChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolPathSet)
				if err := _LombardTokenPool.contract.UnpackLog(event, "PathSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) ParsePathSet(log types.Log) (*LombardTokenPoolPathSet, error) {
	event := new(LombardTokenPoolPathSet)
	if err := _LombardTokenPool.contract.UnpackLog(event, "PathSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolRateLimitConfiguredIterator struct {
	Event *LombardTokenPoolRateLimitConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolRateLimitConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolRateLimitConfigured)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LombardTokenPoolRateLimitConfigured)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LombardTokenPoolRateLimitConfiguredIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolRateLimitConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolRateLimitConfigured struct {
	RemoteChainSelector       uint64
	FastFinality              bool
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterRateLimitConfigured(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardTokenPoolRateLimitConfiguredIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "RateLimitConfigured", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolRateLimitConfiguredIterator{contract: _LombardTokenPool.contract, event: "RateLimitConfigured", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolRateLimitConfigured, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "RateLimitConfigured", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolRateLimitConfigured)
				if err := _LombardTokenPool.contract.UnpackLog(event, "RateLimitConfigured", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseRateLimitConfigured(log types.Log) (*LombardTokenPoolRateLimitConfigured, error) {
	event := new(LombardTokenPoolRateLimitConfigured)
	if err := _LombardTokenPool.contract.UnpackLog(event, "RateLimitConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolReleasedOrMintedIterator struct {
	Event *LombardTokenPoolReleasedOrMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolReleasedOrMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolReleasedOrMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LombardTokenPoolReleasedOrMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LombardTokenPoolReleasedOrMintedIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolReleasedOrMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolReleasedOrMinted struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Recipient           common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardTokenPoolReleasedOrMintedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolReleasedOrMintedIterator{contract: _LombardTokenPool.contract, event: "ReleasedOrMinted", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolReleasedOrMinted)
				if err := _LombardTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseReleasedOrMinted(log types.Log) (*LombardTokenPoolReleasedOrMinted, error) {
	event := new(LombardTokenPoolReleasedOrMinted)
	if err := _LombardTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolRemotePoolAddedIterator struct {
	Event *LombardTokenPoolRemotePoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolRemotePoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolRemotePoolAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LombardTokenPoolRemotePoolAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LombardTokenPoolRemotePoolAddedIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolRemotePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolRemotePoolAdded struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardTokenPoolRemotePoolAddedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolRemotePoolAddedIterator{contract: _LombardTokenPool.contract, event: "RemotePoolAdded", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolRemotePoolAdded)
				if err := _LombardTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseRemotePoolAdded(log types.Log) (*LombardTokenPoolRemotePoolAdded, error) {
	event := new(LombardTokenPoolRemotePoolAdded)
	if err := _LombardTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolRemotePoolRemovedIterator struct {
	Event *LombardTokenPoolRemotePoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolRemotePoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolRemotePoolRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LombardTokenPoolRemotePoolRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LombardTokenPoolRemotePoolRemovedIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolRemotePoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolRemotePoolRemoved struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardTokenPoolRemotePoolRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolRemotePoolRemovedIterator{contract: _LombardTokenPool.contract, event: "RemotePoolRemoved", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolRemotePoolRemoved)
				if err := _LombardTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseRemotePoolRemoved(log types.Log) (*LombardTokenPoolRemotePoolRemoved, error) {
	event := new(LombardTokenPoolRemotePoolRemoved)
	if err := _LombardTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolTokenTransferFeeConfigDeletedIterator struct {
	Event *LombardTokenPoolTokenTransferFeeConfigDeleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolTokenTransferFeeConfigDeletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolTokenTransferFeeConfigDeleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LombardTokenPoolTokenTransferFeeConfigDeleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LombardTokenPoolTokenTransferFeeConfigDeletedIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolTokenTransferFeeConfigDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolTokenTransferFeeConfigDeleted struct {
	DestChainSelector uint64
	Raw               types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*LombardTokenPoolTokenTransferFeeConfigDeletedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolTokenTransferFeeConfigDeletedIterator{contract: _LombardTokenPool.contract, event: "TokenTransferFeeConfigDeleted", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolTokenTransferFeeConfigDeleted)
				if err := _LombardTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseTokenTransferFeeConfigDeleted(log types.Log) (*LombardTokenPoolTokenTransferFeeConfigDeleted, error) {
	event := new(LombardTokenPoolTokenTransferFeeConfigDeleted)
	if err := _LombardTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolTokenTransferFeeConfigUpdatedIterator struct {
	Event *LombardTokenPoolTokenTransferFeeConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolTokenTransferFeeConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolTokenTransferFeeConfigUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LombardTokenPoolTokenTransferFeeConfigUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LombardTokenPoolTokenTransferFeeConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolTokenTransferFeeConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolTokenTransferFeeConfigUpdated struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
	Raw                    types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*LombardTokenPoolTokenTransferFeeConfigUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolTokenTransferFeeConfigUpdatedIterator{contract: _LombardTokenPool.contract, event: "TokenTransferFeeConfigUpdated", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolTokenTransferFeeConfigUpdated)
				if err := _LombardTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseTokenTransferFeeConfigUpdated(log types.Log) (*LombardTokenPoolTokenTransferFeeConfigUpdated, error) {
	event := new(LombardTokenPoolTokenTransferFeeConfigUpdated)
	if err := _LombardTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetCurrentRateLimiterState struct {
	OutboundRateLimiterState RateLimiterTokenBucket
	InboundRateLimiterState  RateLimiterTokenBucket
}
type GetDynamicConfig struct {
	Router         common.Address
	RateLimitAdmin common.Address
	FeeAdmin       common.Address
}
type GetFee struct {
	FeeUSDCents       *big.Int
	DestGasOverhead   uint32
	DestBytesOverhead uint32
	TokenFeeBps       uint16
	IsEnabled         bool
}
type GetLombardConfig struct {
	VerifierResolver common.Address
	Bridge           common.Address
	TokenAdapter     common.Address
}

func (LombardTokenPoolAdvancedPoolHooksUpdated) Topic() common.Hash {
	return common.HexToHash("0xbaff46844acf36d6ee996f489a1a288709c4542bd33cd557770afd267d581209")
}

func (LombardTokenPoolChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (LombardTokenPoolChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (LombardTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x3f1036e85d016a93254a0b1415844f79b85424959d90ae5ad51ce8f4533fe701")
}

func (LombardTokenPoolFastFinalityInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xc6735cd4fa2bbe7b203b1682936e6ee61bc1702464bbbd12abb6630229d9a5f9")
}

func (LombardTokenPoolFastFinalityOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x28d6c52e2b0b7587b0d195539fbe6af984b28791aca4d2cc0844244e38bce29e")
}

func (LombardTokenPoolFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (LombardTokenPoolFinalityConfigSet) Topic() common.Hash {
	return common.HexToHash("0x307cf716eade81675bea3ccb6917b0f91baa2160056765d9a83d76f819caf06a")
}

func (LombardTokenPoolInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (LombardTokenPoolLockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
}

func (LombardTokenPoolLombardConfigurationSet) Topic() common.Hash {
	return common.HexToHash("0x01d5dd7f15328f4241da3a1d9c7b310ae9ac14e8ca441203a7b6f71c7da0c49d")
}

func (LombardTokenPoolOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944")
}

func (LombardTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (LombardTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (LombardTokenPoolPathRemoved) Topic() common.Hash {
	return common.HexToHash("0x465d9b27e0af9978f975c48406a226aab254b237e8027798ee924ef96ee9bb04")
}

func (LombardTokenPoolPathSet) Topic() common.Hash {
	return common.HexToHash("0xbe237be2cca72f95760d5f21feb5f0cf6579119971f023d6ccc49c749ddc9263")
}

func (LombardTokenPoolRateLimitConfigured) Topic() common.Hash {
	return common.HexToHash("0x41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb")
}

func (LombardTokenPoolReleasedOrMinted) Topic() common.Hash {
	return common.HexToHash("0xfc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0")
}

func (LombardTokenPoolRemotePoolAdded) Topic() common.Hash {
	return common.HexToHash("0x7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea")
}

func (LombardTokenPoolRemotePoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76")
}

func (LombardTokenPoolTokenTransferFeeConfigDeleted) Topic() common.Hash {
	return common.HexToHash("0x5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8")
}

func (LombardTokenPoolTokenTransferFeeConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xfae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa1041")
}

func (_LombardTokenPool *LombardTokenPool) Address() common.Address {
	return _LombardTokenPool.address
}

type LombardTokenPoolInterface interface {
	GetAdvancedPoolHooks(opts *bind.CallOpts) (common.Address, error)

	GetAllowedFinalityConfig(opts *bind.CallOpts) ([4]byte, error)

	GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64, fastFinality bool) (GetCurrentRateLimiterState,

		error)

	GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

		error)

	GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, requestedFinalityConfig [4]byte, arg5 []byte) (GetFee,

		error)

	GetLombardConfig(opts *bind.CallOpts) (GetLombardConfig,

		error)

	GetPath(opts *bind.CallOpts, remoteChainSelector uint64) (LombardTokenPoolPath, error)

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRequiredCCVs(opts *bind.CallOpts, arg0 common.Address, arg1 uint64, arg2 *big.Int, arg3 [4]byte, arg4 []byte, arg5 uint8) ([]common.Address, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetSupportedChains(opts *bind.CallOpts) ([]uint64, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	GetTokenDecimals(opts *bind.CallOpts) (uint8, error)

	GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 [4]byte, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error)

	IBridge(opts *bind.CallOpts) (common.Address, error)

	IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error)

	IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error)

	IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error)

	ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, requestedFinalityConfig [4]byte, tokenArgs []byte) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, requestedFinalityConfig [4]byte) (*types.Transaction, error)

	ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	RemovePath(opts *bind.TransactOpts, remoteChainSelector uint64) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetAllowedFinalityConfig(opts *bind.TransactOpts, allowedFinality [4]byte) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error)

	SetPath(opts *bind.TransactOpts, remoteChainSelector uint64, lChainId [32]byte, allowedCaller []byte, remoteAdapter [32]byte) (*types.Transaction, error)

	SetRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateAdvancedPoolHooks(opts *bind.TransactOpts, newHook common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error)

	FilterAdvancedPoolHooksUpdated(opts *bind.FilterOpts) (*LombardTokenPoolAdvancedPoolHooksUpdatedIterator, error)

	WatchAdvancedPoolHooksUpdated(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolAdvancedPoolHooksUpdated) (event.Subscription, error)

	ParseAdvancedPoolHooksUpdated(log types.Log) (*LombardTokenPoolAdvancedPoolHooksUpdated, error)

	FilterChainAdded(opts *bind.FilterOpts) (*LombardTokenPoolChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*LombardTokenPoolChainAdded, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*LombardTokenPoolChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*LombardTokenPoolChainRemoved, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*LombardTokenPoolDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*LombardTokenPoolDynamicConfigSet, error)

	FilterFastFinalityInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardTokenPoolFastFinalityInboundRateLimitConsumedIterator, error)

	WatchFastFinalityInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolFastFinalityInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseFastFinalityInboundRateLimitConsumed(log types.Log) (*LombardTokenPoolFastFinalityInboundRateLimitConsumed, error)

	FilterFastFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardTokenPoolFastFinalityOutboundRateLimitConsumedIterator, error)

	WatchFastFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolFastFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseFastFinalityOutboundRateLimitConsumed(log types.Log) (*LombardTokenPoolFastFinalityOutboundRateLimitConsumed, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*LombardTokenPoolFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*LombardTokenPoolFeeTokenWithdrawn, error)

	FilterFinalityConfigSet(opts *bind.FilterOpts) (*LombardTokenPoolFinalityConfigSetIterator, error)

	WatchFinalityConfigSet(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolFinalityConfigSet) (event.Subscription, error)

	ParseFinalityConfigSet(log types.Log) (*LombardTokenPoolFinalityConfigSet, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardTokenPoolInboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*LombardTokenPoolInboundRateLimitConsumed, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardTokenPoolLockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*LombardTokenPoolLockedOrBurned, error)

	FilterLombardConfigurationSet(opts *bind.FilterOpts, verifier []common.Address, bridge []common.Address, tokenAdapter []common.Address) (*LombardTokenPoolLombardConfigurationSetIterator, error)

	WatchLombardConfigurationSet(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolLombardConfigurationSet, verifier []common.Address, bridge []common.Address, tokenAdapter []common.Address) (event.Subscription, error)

	ParseLombardConfigurationSet(log types.Log) (*LombardTokenPoolLombardConfigurationSet, error)

	FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardTokenPoolOutboundRateLimitConsumedIterator, error)

	WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseOutboundRateLimitConsumed(log types.Log) (*LombardTokenPoolOutboundRateLimitConsumed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LombardTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*LombardTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LombardTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*LombardTokenPoolOwnershipTransferred, error)

	FilterPathRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64, lChainId [][32]byte) (*LombardTokenPoolPathRemovedIterator, error)

	WatchPathRemoved(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolPathRemoved, remoteChainSelector []uint64, lChainId [][32]byte) (event.Subscription, error)

	ParsePathRemoved(log types.Log) (*LombardTokenPoolPathRemoved, error)

	FilterPathSet(opts *bind.FilterOpts, remoteChainSelector []uint64, lChainId [][32]byte) (*LombardTokenPoolPathSetIterator, error)

	WatchPathSet(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolPathSet, remoteChainSelector []uint64, lChainId [][32]byte) (event.Subscription, error)

	ParsePathSet(log types.Log) (*LombardTokenPoolPathSet, error)

	FilterRateLimitConfigured(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardTokenPoolRateLimitConfiguredIterator, error)

	WatchRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolRateLimitConfigured, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRateLimitConfigured(log types.Log) (*LombardTokenPoolRateLimitConfigured, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardTokenPoolReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*LombardTokenPoolReleasedOrMinted, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardTokenPoolRemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*LombardTokenPoolRemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardTokenPoolRemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*LombardTokenPoolRemotePoolRemoved, error)

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*LombardTokenPoolTokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*LombardTokenPoolTokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*LombardTokenPoolTokenTransferFeeConfigUpdatedIterator, error)

	WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigUpdated(log types.Log) (*LombardTokenPoolTokenTransferFeeConfigUpdated, error)

	Address() common.Address
}
