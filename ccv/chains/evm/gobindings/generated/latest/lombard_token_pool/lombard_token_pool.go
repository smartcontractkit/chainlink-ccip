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
	DestGasOverhead                        uint32
	DestBytesOverhead                      uint32
	DefaultBlockConfirmationFeeUSDCents    uint32
	CustomBlockConfirmationFeeUSDCents     uint32
	DefaultBlockConfirmationTransferFeeBps uint16
	CustomBlockConfirmationTransferFeeBps  uint16
	IsEnabled                              bool
}

type LombardTokenPoolPath struct {
	AllowedCaller [32]byte
	LChainId      [32]byte
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
	CustomBlockConfirmation   bool
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
}

type TokenPoolTokenTransferFeeConfigArgs struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
}

var LombardTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20Metadata\"},{\"name\":\"verifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bridge\",\"type\":\"address\",\"internalType\":\"contract IBridgeV2\"},{\"name\":\"adapter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"advancedPoolHooks\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fallbackDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"disableTokenTransferFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAdvancedPoolHooks\",\"inputs\":[],\"outputs\":[{\"name\":\"advancedPoolHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLombardConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"verifierResolver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAdapter\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinBlockConfirmation\",\"inputs\":[],\"outputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPath\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct LombardTokenPool.Path\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"lChainId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_bridge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IBridgeV2\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"lockOrBurnOut\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"lockOrBurnOut\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removePath\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinBlockConfirmation\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPath\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"lChainId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowedCaller\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.RateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateAdvancedPoolHooks\",\"inputs\":[{\"name\":\"newHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AdvancedPoolHooksUpdated\",\"inputs\":[{\"name\":\"oldHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"},{\"name\":\"newHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"feeAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LombardConfigurationSet\",\"inputs\":[{\"name\":\"verifier\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"bridge\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenAdapter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinBlockConfirmationSet\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PathRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"lChainId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PathSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"lChainId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotOwnerOrFeeAdmin\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotSupported\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustomBlockConfirmationsNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"ExecutionError\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"HashMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAllowedCaller\",\"inputs\":[{\"name\":\"allowedCaller\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"received\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmation\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OutboundImplementationNotFoundForVerifier\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PathNotExist\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"RemoteTokenMismatch\",\"inputs\":[{\"name\":\"bridge\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"pool\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroBridge\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroLombardChainId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroVerifierNotAllowed\",\"inputs\":[]}]",
	Bin: "0x61014080604052346103205761010081616276803803809161002182856103f5565b8339810103126103205780516001600160a01b038116918282036103205761004b6020820161042e565b6040820151936001600160a01b03851692909190838603610320576100726060820161042e565b9461007f6080830161042e565b906100ae61008f60a0850161042e565b916100a860e06100a160c0880161042e565b9601610442565b90610450565b9033156103e457600180546001600160a01b03191633179055851580156103d3575b80156103c2575b6103b157608086905260c05260405163313ce56760e01b8152602081600481895afa60009181610375575b5061034a575b5060a052600380546001600160a01b039283166001600160a01b0319918216179091556002805493909216921691909117905582156103395760405163353c26b760e01b8152602081600481875afa801561032d576000906102ee575b60ff915016600181036102d557506001600160a01b0381169485156102c45760e052610100526101208390526001600160a01b038316928284156102b457506101ad916104c6565b604051927f01d5dd7f15328f4241da3a1d9c7b310ae9ac14e8ca441203a7b6f71c7da0c49d600080a4615c9790816105df82396080518181816116640152818161222b015281816123af01528181612fee015281816131ba01528181613268015281816134d10152818161364d015281816137f0015281816140c1015261410e015260a05181818161193201528181613f9401528181614fd10152615054015260c051818181610bd801528181611732015281816122b80152818161307c01526136db015260e051818181610976015281816124a401528181612c3901526138f9015261010051818181610951015261161f01526101205181818161099e015261246a0152f35b90506102bf916104c6565b6101ad565b639533e8c360e01b60005260046000fd5b63398bbe0560e11b600052600160045260245260446000fd5b506020813d602011610325575b81610308602093836103f5565b810103126103205761031b60ff91610442565b610165565b600080fd5b3d91506102fb565b6040513d6000823e3d90fd5b63361106cd60e01b60005260046000fd5b60ff1660ff821681810361035e5750610108565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d6020116103a9575b81610391602093836103f5565b81010312610320576103a290610442565b9038610102565b3d9150610384565b630a64406560e11b60005260046000fd5b506001600160a01b038116156100d7565b506001600160a01b038416156100d0565b639b15e16f60e01b60005260046000fd5b601f909101601f19168101906001600160401b0382119082101761041857604052565b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b038216820361032057565b519060ff8216820361032057565b60405163313ce56760e01b815290602090829060049082906001600160a01b03165afa6000918161048a575b50610485575090565b905090565b9091506020813d6020116104be575b816104a6602093836103f5565b81010312610320576104b790610442565b903861047c565b3d9150610499565b60405190602060008184019463095ea7b360e01b865260018060a01b03169485602486015281196044860152604485526105016064866103f5565b84519082855af16000513d8261055e575b50501561051e57505050565b61055761055c936040519063095ea7b360e01b6020830152602482015260006044820152604481526105516064826103f5565b82610583565b610583565b565b90915061057b57506001600160a01b0381163b15155b3880610512565b600114610574565b906000602091828151910182855af11561032d576000513d6105d557506001600160a01b0381163b155b6105b45750565b635274afe760e01b60009081526001600160a01b0391909116600452602490fd5b600114156105ad56fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a71461419357508063181f5a771461413257806321df0da7146140ee578063240028e8146140975780632422ac4514613fb857806324f65ee714613f7a5780632c06340414613ee157806337a3210d14613eba57806338ff8c3814613e4e57806339077537146135e8578063489a68f214612f565780634c5ef0ed14612f0f5780634e921c3014612e705780635fa1356514612ce457806362ddd3c414612c5d578063708e1f7914612c195780637437ff9f14612bd857806379ba509714612b2b5780638926f54f14612ae557806389720a6214612a2b5780638da5cb5b14612a045780639a4575b9146121c15780639c893fe9146120f4578063a42a7b8b14611f8d578063acfecf9114611e77578063ae39a25714611d20578063b1c71c6514611551578063b794658014611518578063bfeffd3f14611486578063c4bffe2b1461135b578063c7230a6014611187578063d8aa3f401461104d578063dc04fa1f14610bfc578063dc0bd97114610bb8578063dcbd41bc146109ce578063dd65bdb114610932578063e8a1da17146102aa578063f2fde38b146101f55763fa41d79c146101ce57600080fd5b346101f257806003193601126101f257602061ffff60025460a01c16604051908152f35b80fd5b50346101f25760206003193601126101f2576001600160a01b03610217614415565b61021f615140565b1633811461028257807fffffffffffffffffffffffff00000000000000000000000000000000000000008354161782556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b50346101f25760406003193601126101f25760043567ffffffffffffffff811161078b576102dc9036906004016145de565b9060243567ffffffffffffffff811161092e57906102ff849236906004016145de565b93909161030a615140565b83905b8282106107935750505081927ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b8181101561078f578060051b8301358581121561078757830161012081360312610787576040519461037186614318565b61037a82614499565b8652602082013567ffffffffffffffff811161078b5782019436601f8701121561078b578535956103aa87614a11565b966103b86040519889614334565b80885260208089019160051b830101903682116107875760208301905b828210610754575050505060208701958652604083013567ffffffffffffffff8111610750576104089036908501614551565b91604088019283526104326104203660608701614dd6565b9460608a0195865260c0369101614dd6565b9560808901968752835151156107285761045667ffffffffffffffff8a511661597c565b156106f15767ffffffffffffffff895116825260086020526040822061047d86518261546d565b61048b88516002830161546d565b6004855191019080519067ffffffffffffffff82116106c4576104ae8354614c24565b601f8111610689575b50602090601f8311600114610608576104e792918691836105fd575b50506000198260011b9260031b1c19161790565b90555b815b88518051821015610521579061051b6001926105148367ffffffffffffffff8f511692614be1565b519061517e565b016104ec565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c29391999750956105ef67ffffffffffffffff60019796949851169251935191516105bb610586604051968796875261010060208801526101008701906143d2565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a1019392909193610340565b015190508e806104d3565b83865281862091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416875b8181106106715750908460019594939210610658575b505050811b0190556104ea565b015160001960f88460031b161c191690558d808061064b565b92936020600181928786015181550195019301610635565b6106b49084875260208720601f850160051c810191602086106106ba575b601f0160051c0190614e72565b8d6104b7565b90915081906106a7565b6024857f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60248267ffffffffffffffff8b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b6004827f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b8280fd5b813567ffffffffffffffff8111610783576020916107788392833691890101614551565b8152019101906103d5565b8680fd5b8480fd5b5080fd5b8380f35b9267ffffffffffffffff6107b56107b08486889a9699979a614da9565b61494b565b16916107c08361575d565b156109025782845260086020526107dc600560408620016156fa565b94845b865181101561081557600190858752600860205261080e60056040892001610807838b614be1565b519061585d565b50016107df565b50939692909450949094808752600860205260056040882088815588600182015588600282015588600382015588600482016108518154614c24565b806108c1575b50505001805490888155816108a3575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a10190919493929461030d565b885260208820908101905b81811015610867578881556001016108ae565b601f81116001146108d75750555b888a80610857565b818352602083206108f291601f01861c810190600101614e72565b80825281602081209155556108cf565b602484847f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b8380fd5b50346101f257806003193601126101f257604080516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811682527f0000000000000000000000000000000000000000000000000000000000000000811660208301527f00000000000000000000000000000000000000000000000000000000000000001691810191909152606090f35b0390f35b50346101f25760206003193601126101f25760043567ffffffffffffffff811161078b57610a0090369060040161460f565b6001600160a01b03600a541633141580610ba3575b610b7757825b818110610a26578380f35b610a31818385614d5b565b67ffffffffffffffff610a438261494b565b1690610a5c826000526007602052604060002054151590565b15610b4b57907f41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb60e083610b0b610ae5602060019897018b610a9d82614d6b565b15610b12578790526004602052610ac460408d20610abe3660408801614dd6565b9061546d565b868c526005602052610ae060408d20610abe3660a08801614dd6565b614d6b565b916040519215158352610afe6020840160408301614e2e565b60a0608084019101614e2e565ba201610a1b565b60026040828a610ae094526008602052610b34828220610abe36858c01614dd6565b8a815260086020522001610abe3660a08801614dd6565b602486837f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b6024837f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b506001600160a01b0360015416331415610a15565b50346101f257806003193601126101f25760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346101f25760406003193601126101f25760043567ffffffffffffffff811161078b57610c2e90369060040161460f565b60243567ffffffffffffffff811161092e57610c4e9036906004016145de565b919092610c59615140565b845b828110610cc557505050825b818110610c72578380f35b8067ffffffffffffffff610c8c6107b06001948688614da9565b16808652600b6020528560408120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee88680a201610c67565b610cd36107b0828585614d5b565b610cde828585614d5b565b90602082019060e0830190610cf282614d6b565b156110185760a0840161271061ffff610d0a83614d78565b1610156110095760c085019161271061ffff610d2585614d78565b161015610fd15763ffffffff610d3a86614d87565b1615610f9c5767ffffffffffffffff1694858c52600b60205260408c20610d6086614d87565b63ffffffff16908054906040840191610d7883614d87565b60201b67ffffffff0000000016936060860194610d9486614d87565b60401b6bffffffff0000000000000000169660800196610db388614d87565b60601b6fffffffff0000000000000000000000001691610dd28a614d78565b60801b71ffff000000000000000000000000000000001693610df38c614d78565b60901b73ffff00000000000000000000000000000000000016957fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016177fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16177fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff16177fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff161717178155610ea687614d6b565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff00000000000000000000000000000000000000001617905560405196610ef790614d98565b63ffffffff168752610f0890614d98565b63ffffffff166020870152610f1c90614d98565b63ffffffff166040860152610f3090614d98565b63ffffffff166060850152610f44906144dd565b61ffff166080840152610f56906144dd565b61ffff1660a0830152610f68906144ae565b151560c082015260e07ffae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa104191a2600101610c5b565b7f12332265000000000000000000000000000000000000000000000000000000008c5267ffffffffffffffff1660045260248bfd5b60248c61ffff610fe086614d78565b7f95f3517a00000000000000000000000000000000000000000000000000000000835216600452fd5b8a61ffff610fe0602493614d78565b7f12332265000000000000000000000000000000000000000000000000000000008a5267ffffffffffffffff16600452602489fd5b50346101f25760806003193601126101f257611067614415565b50611070614482565b6110786144cc565b5060643567ffffffffffffffff8111610750579167ffffffffffffffff6040926110a860e09536906004016144ec565b50508260c085516110b8816142fc565b82815282602082015282878201528260608201528260808201528260a08201520152168152600b60205220604051906110f0826142fc565b5461ffff818163ffffffff82169485815263ffffffff60208201818560201c1681528160408401818760401c168152816060860193818960601c16855260ff60c060808901988a8c60801c168a528a60a082019c60901c168c52019b60a01c1615158b526040519b8c52511660208b0152511660408901525116606087015251166080850152511660a083015251151560c0820152f35b50346101f25760406003193601126101f25760043567ffffffffffffffff811161078b576111b99036906004016145de565b906111c2614441565b916001600160a01b036001541633141580611346575b61131a576001600160a01b0383169081156112f257845b8181106111fa578580f35b6001600160a01b03611215611210838588614da9565b614960565b1690604051917f70a08231000000000000000000000000000000000000000000000000000000008352306004840152602083602481845afa80156112e757859089906112ad575b600194508061126f575b505050016111ef565b60208161129e7f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e938c876153ac565b604051908152a3388481611266565b5050909160203d81116112e0575b6112c58183614334565b602082600092810103126101f257509084600193925161125c565b503d6112bb565b6040513d8a823e3d90fd5b6004857f8579befe000000000000000000000000000000000000000000000000000000008152fd5b6024847fcb1afbd700000000000000000000000000000000000000000000000000000000815233600452fd5b506001600160a01b03600c54163314156111d8565b50346101f257806003193601126101f257604051906006548083528260208101600684526020842092845b81811061146d57505061139b92500383614334565b81516113bf6113a982614a11565b916113b76040519384614334565b808352614a11565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602083019301368437805b845181101561141e578067ffffffffffffffff61140b60019388614be1565b51166114178286614be1565b52016113ec565b50925090604051928392602084019060208552518091526040840192915b81811061144a575050500390f35b825167ffffffffffffffff1684528594506020938401939092019160010161143c565b8454835260019485019487945060209093019201611386565b50346101f25760206003193601126101f2576004356001600160a01b03811680910361078b576114b4615140565b7fffffffffffffffffffffffff00000000000000000000000000000000000000006003547fbaff46844acf36d6ee996f489a1a288709c4542bd33cd557770afd267d581209604080516001600160a01b0384168152856020820152a1161760035580f35b50346101f25760206003193601126101f2576109ca61153d61153861446b565b614d39565b6040519182916020835260208301906143d2565b50346101f25760606003193601126101f25760043567ffffffffffffffff811161078b57806004019160a060031983360301126101f2576115906144bb565b9160443567ffffffffffffffff8111610750576115b19036906004016144ec565b9290936115bc614bc8565b5060248301956115cb8761494b565b9367ffffffffffffffff604051957f958021a7000000000000000000000000000000000000000000000000000000008752166004860152604060248601528360448601526020856064816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa948515611d15578495611cd9575b506001600160a01b03851615611cb1576116af7f000000000000000000000000000000000000000000000000000000000000000097611694606484013580988b6153ac565b61169c614bc8565b506116a78585615919565b97369161451a565b9060848101926116be84614960565b6001600160a01b03808b16911603611c745777ffffffffffffffff000000000000000000000000000000006116f28b61494b565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611bdc578791611c3a575b50611c125767ffffffffffffffff6117798b61494b565b16611791816000526007602052604060002054151590565b15611be75760206001600160a01b0360025416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015611bdc578790611b9c575b6001600160a01b039150163303611b705761ffff6118038989614ecb565b9516948515611af85761ffff60025460a01c168015611ad057808710611aa057507f61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac2097932867ffffffffffffffff6118578d61494b565b169182895260046020528061187c8d6001600160a01b0360408d209116928391615a31565b604080516001600160a01b039290921682526020820192909252a25b6001600160a01b03600354169283611998575b61198e8a61192a6115388e6118c08e8e614ecb565b936118ca8261494b565b507ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff6118fe8461494b565b604080516001600160a01b0390951685523360208601528401889052169180606081015b0390a261494b565b9060405160ff7f000000000000000000000000000000000000000000000000000000000000000016602082015260208152611966604082614334565b6040519261197384614295565b835260208301526040519283926040845260408401906145b4565b9060208301520390f35b833b1561078357869493929185918c604051988997889687957f5c3af7ca0000000000000000000000000000000000000000000000000000000087526004870160609052806119e6916156aa565b6064880160a090526101048801906119fd926149b1565b93611a0790614499565b67ffffffffffffffff166084870152604401611a2290614457565b6001600160a01b031660a48601528c60c4860152611a3f90614457565b6001600160a01b031660e48501526024840152828103600319016044840152611a67916143d2565b03925af18015611a9557611a80575b80808080806118ab565b611a8b828092614334565b6101f25780611a76565b6040513d84823e3d90fd5b87604491887f7911d95b000000000000000000000000000000000000000000000000000000008352600452602452fd5b6004887f98d50fd7000000000000000000000000000000000000000000000000000000008152fd5b7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894467ffffffffffffffff611b2b8d61494b565b1691828952600860205280611b508d6001600160a01b0360408d209116928391615a31565b604080516001600160a01b039290921682526020820192909252a2611898565b6024867f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d602011611bd4575b81611bb660209383614334565b8101031261078357611bcf6001600160a01b039161492a565b6117e5565b3d9150611ba9565b6040513d89823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008752600452602486fd5b6004867f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b90506020813d602011611c6c575b81611c5560209383614334565b8101031261078357611c669061493e565b38611762565b3d9150611c48565b6024866001600160a01b03611c8887614960565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835216600452fd5b6004847f7af97002000000000000000000000000000000000000000000000000000000008152fd5b9094506020813d602011611d0d575b81611cf560209383614334565b8101031261092e57611d069061492a565b933861164f565b3d9150611ce8565b6040513d86823e3d90fd5b50346101f25760606003193601126101f257611d3a614415565b90611d43614441565b604435926001600160a01b03841680850361092e57611d60615140565b6001600160a01b0382168015611e4f5794611e49917f3f1036e85d016a93254a0b1415844f79b85424959d90ae5ad51ce8f4533fe70195967fffffffffffffffffffffffff000000000000000000000000000000000000000060025416176002556001600160a01b0385167fffffffffffffffffffffffff0000000000000000000000000000000000000000600a541617600a557fffffffffffffffffffffffff0000000000000000000000000000000000000000600c541617600c55604051938493849160409194936001600160a01b03809281606087019816865216602085015216910152565b0390a180f35b6004857f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b50346101f257611e863661456f565b611e8e615140565b67ffffffffffffffff831692611eb1846000526007602052604060002054151590565b15611f61578385526008602052611ee060056040872001611ed336858761451a565b602081519101209061585d565b15611f265750907f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d7691611f206040519283926020845260208401916149b1565b0390a280f35b90611f5d906040519384937f74f23c7c000000000000000000000000000000000000000000000000000000008552600485016149f0565b0390fd5b602485857f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50346101f25760206003193601126101f25767ffffffffffffffff611fb061446b565b1681526008602052611fc7600560408320016156fa565b80517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061200c611ff683614a11565b926120046040519485614334565b808452614a11565b01835b8181106120e3575050825b8251811015612060578061203060019285614be1565b518552600960205261204460408620614c77565b61204e8285614be1565b526120598184614be1565b500161201a565b81846040519182916020830160208452825180915260408401602060408360051b870101940192905b82821061209857505050500390f35b919360206120d3827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0600195979984950301865288516143d2565b9601920192018594939192612089565b80606060208093860101520161200f565b50346101f25760206003193601126101f25767ffffffffffffffff61211761446b565b61211f615140565b16808252600d60205260408220906040519161213a83614295565b600181549182855201549060208401918252156121955760207f8a8e4c676433747219d2fee4ea128776522bb0177478e1e0a375e880948ed37b91838652600d8252856001604082208281550155519351604051908152a380f35b602484837fa28cbf38000000000000000000000000000000000000000000000000000000008252600452fd5b50346101f25760206003193601126101f25760043567ffffffffffffffff811161078b578060040160a06003198336030112610750576121ff614bc8565b506040516020926122108483614334565b8482526084810161222081614960565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116036129f057602482019377ffffffffffffffff000000000000000000000000000000006122798661494b565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015286816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156112e75788916129bb575b506129935767ffffffffffffffff6122ff8661494b565b16612317816000526007602052604060002054151590565b1561296857866001600160a01b0360025416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa80156112e7578890612929575b6001600160a01b0391501633036128fd5786606484013594859367ffffffffffffffff6123968961494b565b1680845260088a526123d7604085206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016998a91615a31565b604080516001600160a01b038a168152602081018890527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789449190a26001600160a01b03600354169081612805575b5050505067ffffffffffffffff61243b8661494b565b168752600d86526040872060405161245281614295565b6001825492838352015491888201928352156127c7577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b038116156127bf57905b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016926001600160a01b03815193604051947f6e48b60d000000000000000000000000000000000000000000000000000000008652600486015216928360248201528a81604481885afa9081156127b4578c91612787575b506125276115388b61494b565b8b815191818082019384920101031261270c5751908181036127595750508961255086806148d9565b90500361271157604490519601958961257261256b89614960565b96806148d9565b908092918101031261270c5760409460c4938d92359051906001600160a01b038851998a9889977f793ea55b00000000000000000000000000000000000000000000000000000000895260048901526024880152166044860152606485015288608485015260a48401525af196871561270057809761269f575b5050917ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff8593611922611538966126376126316126619a61494b565b93614960565b6040519485941696846001600160a01b036040929594938160608401971683521660208201520152565b916040519082820152818152612678604082614334565b6040519261268584614295565b8352818301526109ca6040519282849384528301906145b4565b909196506040823d6040116126f8575b816126bc60409383614334565b810103126101f25750840151947ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff6125ec565b3d91506126af565b604051903d90823e3d90fd5b600080fd5b611f5d8a61271f87806148d9565b92906040519384937fa3c8cf09000000000000000000000000000000000000000000000000000000008552600485015260248401916149b1565b7f81d8236e000000000000000000000000000000000000000000000000000000008d5260045260245260448bfd5b90508a81813d83116127ad575b61279e8183614334565b8101031261270c57513861251a565b503d612794565b6040513d8e823e3d90fd5b50859061249a565b60248967ffffffffffffffff6127dc8a61494b565b7fa28cbf3800000000000000000000000000000000000000000000000000000000835216600452fd5b813b1561092e5783918991836040518096819582947f5c3af7ca00000000000000000000000000000000000000000000000000000000845260048401606090528d6128508d806156aa565b6064870160a09052610104870190612867926149b1565b9261287190614499565b67ffffffffffffffff16608486015260440161288c90614457565b6001600160a01b031660a48501528c60c48501526128a990614457565b6001600160a01b031660e48401528360248401528281036003190160448401526128d2916143d2565b03925af18015611a95576128e8575b8080612425565b816128f291614334565b6107835786386128e1565b6024877f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b508681813d8311612961575b61293f8183614334565b8101031261295d576129586001600160a01b039161492a565b61236a565b8780fd5b503d612935565b7fa9902c7e000000000000000000000000000000000000000000000000000000008852600452602487fd5b6004877f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b90508681813d83116129e9575b6129d28183614334565b8101031261295d576129e39061493e565b386122e8565b503d6129c8565b856001600160a01b03611c88602493614960565b50346101f257806003193601126101f25760206001600160a01b0360015416604051908152f35b50346101f25760c06003193601126101f257612a45614415565b612a4d614482565b9060643561ffff8116810361092e5760843567ffffffffffffffff811161078757612a7c9036906004016144ec565b9160a43593600285101561078357612a979560443591614a29565b90604051918291602083016020845282518091526020604085019301915b818110612ac3575050500390f35b82516001600160a01b0316845285945060209384019390920191600101612ab5565b50346101f25760206003193601126101f2576020612b2167ffffffffffffffff612b0d61446b565b166000526007602052604060002054151590565b6040519015158152f35b50346101f257806003193601126101f25780546001600160a01b0381163303612bb0577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551682556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b50346101f257806003193601126101f257600254600a54600c54604080516001600160a01b0394851681529284166020840152921691810191909152606090f35b50346101f257806003193601126101f25760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346101f257612c6c3661456f565b612c7893929193615140565b67ffffffffffffffff8216612c9a816000526007602052604060002054151590565b15612cb95750612cb69293612cb091369161451a565b9061517e565b80f35b7f1e670e4b000000000000000000000000000000000000000000000000000000008452600452602483fd5b50346101f25760606003193601126101f257612cfe61446b565b6024359060443567ffffffffffffffff811161092e57612d229036906004016144ec565b612d2a615140565b67ffffffffffffffff831692612d4d846000526007602052604060002054151590565b15612e44578415612e1c57612d6c612d6636848661451a565b82614974565b15611f26575060208103612ddf57816020918101031261270c5760207f83eda38165c92f401f97217d5ead82ef163d0b716c3979eff4670361bc2dc0c99135604051612db781614295565b8181526001838201878152868952600d8552604089209251835551910155604051908152a380f35b611f5d6040519283927f5552d6310000000000000000000000000000000000000000000000000000000084526020600485015260248401916149b1565b6004867f5a39e303000000000000000000000000000000000000000000000000000000008152fd5b602486857f2e59db3a000000000000000000000000000000000000000000000000000000008252600452fd5b50346101f25760206003193601126101f25760043561ffff811690818103610750577fa7f8dbba8cdb126ce4a0e7939ec58e0161b70d808b585dd651d68e59d27e11fb91602091612ebf615140565b7fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff75ffff00000000000000000000000000000000000000006002549260a01b16911617600255604051908152a180f35b50346101f25760406003193601126101f257612f2961446b565b906024359067ffffffffffffffff82116101f2576020612b2184612f503660048701614551565b90614974565b50346101f25760406003193601126101f2576004359067ffffffffffffffff82116101f25781600401610100600319843603011261078b57612f966144bb565b9180604051612fa4816142e0565b5260648401359260c48501612fd4612fce612fc9612fc284886148d9565b369161451a565b614f5d565b86615051565b946084870190612fe382614960565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116036135d457602488019577ffffffffffffffff0000000000000000000000000000000061303c8861494b565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611bdc57879161359a575b50611c125767ffffffffffffffff6130c38861494b565b166130db816000526007602052604060002054151590565b15611be75760206001600160a01b0360025416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115611bdc578791613560575b5015611b70576131458761494b565b9461315b60a48b0196612f50612fc289866148d9565b156135195761ffff1690811561347f5767ffffffffffffffff61317d8961494b565b1680885260056020527f0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce6158a806131e260408c206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391615a31565b604080516001600160a01b039290921682526020820192909252a25b6001600160a01b036003541694856132ca575b60208a60448d017ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc067ffffffffffffffff6132626126318f61325286614960565b5061325c8161494b565b5061494b565b604080517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b039081168252336020830152909216908201526060810185905292169180608081015b0390a2806040516132c1816142e0565b52604051908152f35b853b1561295d5792889694928b8997959388946040519a8b998a9889977f1abfe46e000000000000000000000000000000000000000000000000000000008952600489016060905261331c86806156aa565b60648b0161010090526101648b0190613334926149b1565b9361333e90614499565b67ffffffffffffffff1660848a015261335960448801614457565b6001600160a01b031660a48a015260c489015261337590614457565b6001600160a01b031660e488015261338d90846156aa565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c888403016101048901526133c292916149b1565b906133cd90836156aa565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c8784030161012488015261340292916149b1565b9160e40161340f916156aa565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c8584030161014486015261344492916149b1565b908b6024840152604483015203925af18015611a955761346a575b808080808080613211565b613475828092614334565b6101f2578061345f565b67ffffffffffffffff6134918961494b565b1680885260086020527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c8a806134f9600260408d20016001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391615a31565b604080516001600160a01b039290921682526020820192909252a26131fe565b61352386836148d9565b611f5d6040519283927f24eb47e50000000000000000000000000000000000000000000000000000000084526020600485015260248401916149b1565b90506020813d602011613592575b8161357b60209383614334565b810103126107835761358c9061493e565b38613136565b3d915061356e565b90506020813d6020116135cc575b816135b560209383614334565b81010312610783576135c69061493e565b386130ac565b3d91506135a8565b6024856001600160a01b03611c8885614960565b50346101f25760206003193601126101f25760043567ffffffffffffffff811161078b57806004019161010060031983360301126101f2578060405161362d816142e0565b52606482013590608483019361364285614960565b6001600160a01b03807f000000000000000000000000000000000000000000000000000000000000000016911603613e3a57602484019177ffffffffffffffff0000000000000000000000000000000061369b8461494b565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611a95578291613e00575b50613dd85767ffffffffffffffff6137228461494b565b1661373a816000526007602052604060002054151590565b15613dac5760206001600160a01b0360025416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115611a95578291613d72575b5015613d46576137a48361494b565b6137b960a4870191612f50612fc284876148d9565b15613d3c5784959667ffffffffffffffff6137d38661494b565b168084526008602052613818600260408620016001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016988991615a31565b604080516001600160a01b0389168152602081018a90527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a26001600160a01b03600354169081613b90575b5050509061387760e48701826148d9565b8192910160408382031261092e57823567ffffffffffffffff811161078757816138a2918501614551565b9260208101359067ffffffffffffffff8211613b21576138c3929101614551565b6040517fd5438eae0000000000000000000000000000000000000000000000000000000081526020816004816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115613b8557908592918391613b30575b5061397c836001600160a01b0361398e97604051988996879586937fa62085060000000000000000000000000000000000000000000000000000000085526040600486015260448501906143d2565b906003198483030160248501526143d2565b0393165af18015613b255783928491613a92575b5015613a6a576139b860209160c48901906148d9565b908092918101031261270c573503613a42575067ffffffffffffffff6020946132b185613a116044613a0a7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09761494b565b9401614960565b604080516001600160a01b039889168152336020820152979091169087015260608601529116929081906080820190565b807f3f4d60530000000000000000000000000000000000000000000000000000000060049252fd5b6004837f2532cf45000000000000000000000000000000000000000000000000000000008152fd5b9250503d8084843e613aa48184614334565b820160608382031261092e57825190613abf6020850161493e565b9360408101519067ffffffffffffffff8211610783570181601f82011215613b2157805191613aed83614375565b90613afb6040519283614334565b838252602084840101116107835790602080613b1a94930191016143af565b91386139a2565b8580fd5b6040513d85823e3d90fd5b9193949250506020813d602011613b7d575b81613b4f60209383614334565b810103126107875791849161397c836001600160a01b03613b7361398e989761492a565b939750505061392d565b3d9150613b42565b6040513d87823e3d90fd5b813b1561092e579183918693838a8c604051978896879586947f1abfe46e0000000000000000000000000000000000000000000000000000000086528d600487016060905280613bdf916156aa565b606488016101009052610164880190613bf7926149b1565b93613c0190614499565b67ffffffffffffffff166084870152613c1c60448601614457565b6001600160a01b031660a487015260c4860152613c3890614457565b6001600160a01b031660e4850152613c50908c6156aa565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c85840301610104860152613c8592916149b1565b613c9260c483018c6156aa565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c85840301610124860152613cc792916149b1565b9060e401613cd5908b6156aa565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c84840301610144850152613d0a92916149b1565b8c602483015282604483015203925af18015611a9557613d2c575b8080613866565b81613d3691614334565b38613d25565b61352390836148d9565b807f728fe07b000000000000000000000000000000000000000000000000000000006024925233600452fd5b90506020813d602011613da4575b81613d8d60209383614334565b8101031261078b57613d9e9061493e565b38613795565b3d9150613d80565b602492507fa9902c7e000000000000000000000000000000000000000000000000000000008252600452fd5b807f53ad11d80000000000000000000000000000000000000000000000000000000060049252fd5b90506020813d602011613e32575b81613e1b60209383614334565b8101031261078b57613e2c9061493e565b3861370b565b3d9150613e0e565b6024826001600160a01b03611c8888614960565b50346101f25760206003193601126101f2576040809167ffffffffffffffff613e7561446b565b8260208551613e8381614295565b8281520152168152600d602052208151613e9c81614295565b60206001835493848452015491019081528251918252516020820152f35b50346101f257806003193601126101f25760206001600160a01b0360035416604051908152f35b50346101f25760c06003193601126101f257613efb614415565b50613f04614482565b613f0c61442b565b506084359161ffff831683036101f25760a4359067ffffffffffffffff82116101f25760a063ffffffff8061ffff613f538888613f4c3660048b016144ec565b5050614754565b95926040979194975197885216602087015216604085015216606083015215156080820152f35b50346101f257806003193601126101f257602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346101f25760406003193601126101f257613fd261446b565b6024359182151583036101f257610140614095613fef85856146d1565b61404560409392935180946fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b60a08301906fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565bf35b50346101f25760206003193601126101f2576020906140b4614415565b90506001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b50346101f257806003193601126101f25760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346101f257806003193601126101f257506109ca604051614155604082614334565b601a81527f4c6f6d62617264546f6b656e506f6f6c20322e302e302d64657600000000000060208201526040519182916020835260208301906143d2565b90503461078b57602060031936011261078b576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361075057602092507faff2afbf00000000000000000000000000000000000000000000000000000000811490811561426b575b8115614241575b8115614217575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438614210565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150614209565b7f331710310000000000000000000000000000000000000000000000000000000081149150614202565b6040810190811067ffffffffffffffff8211176142b157604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff8211176142b157604052565b60e0810190811067ffffffffffffffff8211176142b157604052565b60a0810190811067ffffffffffffffff8211176142b157604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176142b157604052565b67ffffffffffffffff81116142b157601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60005b8381106143c25750506000910152565b81810151838201526020016143b2565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f60209361440e815180928187528780880191016143af565b0116010190565b600435906001600160a01b038216820361270c57565b606435906001600160a01b038216820361270c57565b602435906001600160a01b038216820361270c57565b35906001600160a01b038216820361270c57565b6004359067ffffffffffffffff8216820361270c57565b6024359067ffffffffffffffff8216820361270c57565b359067ffffffffffffffff8216820361270c57565b3590811515820361270c57565b6024359061ffff8216820361270c57565b6044359061ffff8216820361270c57565b359061ffff8216820361270c57565b9181601f8401121561270c5782359167ffffffffffffffff831161270c576020838186019501011161270c57565b92919261452682614375565b916145346040519384614334565b82948184528183011161270c578281602093846000960137010152565b9080601f8301121561270c5781602061456c9335910161451a565b90565b90604060031983011261270c5760043567ffffffffffffffff8116810361270c57916024359067ffffffffffffffff821161270c576145b0916004016144ec565b9091565b61456c9160206145cd83516040845260408401906143d2565b9201519060208184039101526143d2565b9181601f8401121561270c5782359167ffffffffffffffff831161270c576020808501948460051b01011161270c57565b9181601f8401121561270c5782359167ffffffffffffffff831161270c576020808501948460081b01011161270c57565b6040519061464d82614318565b60006080838281528260208201528260408201528260608201520152565b9060405161467881614318565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff916146e3614640565b506146ec614640565b506147205716600052600860205260406000209061456c61471460026147196147148661466b565b614ed8565b940161466b565b169081600052600460205261473b614714604060002061466b565b91600052600560205261456c614714604060002061466b565b9061ffff8060025460a01c16911692831515928380946148d1575b6148a75767ffffffffffffffff16600052600b60205260406000209160405192614798846142fc565b5463ffffffff81168452602084019563ffffffff8260201c168752604085019263ffffffff8360401c168452606086019163ffffffff8460601c168352608087019761ffff8560801c16895260ff60a089019561ffff8160901c16875260a01c1615801560c08a015261488c5761482d575050505063ffffffff808061ffff9351169451169551169351169193929190600190565b81939750809294501061485c57505063ffffffff808061ffff9351169451169551169351169193929190600190565b7f7911d95b0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b50505050505092505050600090600090600090600090600090565b7f98d50fd70000000000000000000000000000000000000000000000000000000060005260046000fd5b50821561476f565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561270c570180359067ffffffffffffffff821161270c5760200191813603831361270c57565b51906001600160a01b038216820361270c57565b5190811515820361270c57565b3567ffffffffffffffff8116810361270c5790565b356001600160a01b038116810361270c5790565b9067ffffffffffffffff61456c92166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b60409067ffffffffffffffff61456c959316815281602082015201916149b1565b67ffffffffffffffff81116142b15760051b60200190565b9593919294906001600160a01b0360035416958615614ba657614aaa9467ffffffffffffffff61ffff936001600160a01b036040519b7f89720a62000000000000000000000000000000000000000000000000000000008d521660048c01521660248a0152604489015216606487015260c0608487015260c48601916149b1565b916002821015614b77578380600094819460a483015203915afa908115614b6b57600091614ad6575090565b3d8083833e614ae58183614334565b8101906020818303126107505780519067ffffffffffffffff821161092e570181601f8201121561075057805190614b1c82614a11565b93614b2a6040519586614334565b82855260208086019360051b8301019384116101f25750602001905b828210614b535750505090565b60208091614b608461492a565b815201910190614b46565b6040513d6000823e3d90fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b5050505050505050604051614bbc602082614334565b60008152600036813790565b60405190614bd582614295565b60606020838281520152565b8051821015614bf55760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600182811c92168015614c6d575b6020831014614c3e57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691614c33565b9060405191826000825492614c8b84614c24565b8084529360018116908115614cf95750600114614cb2575b50614cb092500383614334565b565b90506000929192526020600020906000915b818310614cdd575050906020614cb09282010138614ca3565b6020919350806001915483858901015201910190918492614cc4565b60209350614cb09592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138614ca3565b67ffffffffffffffff16600052600860205261456c6004604060002001614c77565b9190811015614bf55760081b0190565b35801515810361270c5790565b3561ffff8116810361270c5790565b3563ffffffff8116810361270c5790565b359063ffffffff8216820361270c57565b9190811015614bf55760051b0190565b35906fffffffffffffffffffffffffffffffff8216820361270c57565b919082606091031261270c576040516060810181811067ffffffffffffffff8211176142b1576040526040614e29818395614e10816144ae565b8552614e1e60208201614db9565b602086015201614db9565b910152565b6fffffffffffffffffffffffffffffffff614e6c60408093614e4f816144ae565b1515865283614e6060208301614db9565b16602087015201614db9565b16910152565b818110614e7d575050565b60008155600101614e72565b81810292918115918404141715614e9c57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b91908203918211614e9c57565b614ee0614640565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff8083511691614f3d6020850193614f37614f2a63ffffffff87511642614ecb565b8560808901511690614e89565b9061569d565b80821015614f5657505b16825263ffffffff4216905290565b9050614f47565b80518015614fcd57602003614f8f57805160208281019183018390031261270c57519060ff8211614f8f575060ff1690565b611f5d906040519182917f953576f70000000000000000000000000000000000000000000000000000000083526020600484015260248301906143d2565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff8211614e9c57565b60ff16604d8111614e9c57600a0a90565b8115615022570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff8116928284146151395782841161510f579061509691614ff3565b91604d60ff84161180156150f4575b6150be575050906150b861456c92615007565b90614e89565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b506150fe83615007565b8015615022576000190484116150a5565b61511891614ff3565b91604d60ff8416116150be5750509061513361456c92615007565b90615018565b5050505090565b6001600160a01b0360015416330361515457565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b908051156153825767ffffffffffffffff815160208301209216918260005260086020526151b38160056040600020016159dc565b1561533e5760005260096020526040600020815167ffffffffffffffff81116142b1576151e08254614c24565b601f811161530c575b506020601f8211600114615264579161523e827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea959361525495600091615259575b506000198260011b9260031b1c19161790565b90556040519182916020835260208301906143d2565b0390a2565b90508401513861522b565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b8181106152f45750926152549492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9896106152db575b5050811b01905561153d565b85015160001960f88460031b161c1916905538806152cf565b9192602060018192868a015181550194019201615294565b61533890836000526020600020601f840160051c810191602085106106ba57601f0160051c0190614e72565b386151e9565b5090611f5d6040519283927f393b8ad200000000000000000000000000000000000000000000000000000000845260048401526040602484015260448301906143d2565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b6040517fa9059cbb0000000000000000000000000000000000000000000000000000000060208083019182526001600160a01b039490941660248301526044808301959095529381529092600091615405606482614334565b519082855af115614b6b576000513d61546457506001600160a01b0381163b155b61542d5750565b6001600160a01b03907f5274afe7000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60011415615426565b8151919291156155ef576fffffffffffffffffffffffffffffffff6040840151166fffffffffffffffffffffffffffffffff6020850151161061558c57614cb091925b805182547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff0000000000000000000000000000000000000000161782556020810151825460409290920151608090811b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff9290921691821760018501557fffffffffffffffffffffffff0000000000000000000000000000000000000000909216174290911b73ffffffff0000000000000000000000000000000016179055565b6064836155ed604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff6040840151161580159061567e575b61561d57614cb091926154b0565b6064836155ed604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff602084015116151561560f565b91908201809211614e9c57565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561270c57016020813591019167ffffffffffffffff821161270c57813603831361270c57565b906040519182815491828252602082019060005260206000209260005b81811061572c575050614cb092500383614334565b8454835260019485019487945060209093019201615717565b8054821015614bf55760005260206000200190600090565b6000818152600760205260409020548015615856576000198101818111614e9c57600654906000198201918211614e9c57818103615805575b50505060065480156157d657600019016157b1816006615745565b60001982549160031b1b19169055600655600052600760205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b61583e615816615827936006615745565b90549060031b1c9283926006615745565b81939154906000199060031b92831b921b19161790565b90556000526007602052604060002055388080615796565b5050600090565b9060018201918160005282602052604060002054801515600014615910576000198101818111614e9c578254906000198201918211614e9c578181036158d9575b505050805480156157d65760001901906158b88282615745565b60001982549160031b1b191690555560005260205260006040812055600190565b6158f96158e96158279386615745565b90549060031b1c92839286615745565b90556000528360205260406000205538808061589e565b50505050600090565b906127109167ffffffffffffffff6159336020830161494b565b166000908152600b602052604090209161ffff161561596657606061ffff615962935460901c16910135614e89565b0490565b606061ffff615962935460801c16910135614e89565b806000526007602052604060002054156000146159d657600654680100000000000000008110156142b1576159bd6158278260018594016006556006615745565b9055600654906000526007602052604060002055600190565b50600090565b600082815260018201602052604090205461585657805490680100000000000000008210156142b15782615a1a615827846001809601855584615745565b905580549260005201602052604060002055600190565b9182549060ff8260a01c16158015615c82575b615c7c576fffffffffffffffffffffffffffffffff82169160018501908154615a8963ffffffff6fffffffffffffffffffffffffffffffff83169360801c1642614ecb565b9081615bde575b5050848110615b9f5750838310615aea575050615abf6fffffffffffffffffffffffffffffffff928392614ecb565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b9190915460801c928315615b5e5781615b0291614ecb565b92600019810190808211614e9c57615b25615b2a926001600160a01b039661569d565b615018565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b6001600160a01b0383837fd0c8d23a000000000000000000000000000000000000000000000000000000006000526000196004526024521660445260646000fd5b82856001600160a01b03927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611615c5257615bf992614f379160801c90614e89565b80841015615c4d5750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178655923880615a90565b615c04565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b508215615a4456fea164736f6c634300081a000a",
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

func (_LombardTokenPool *LombardTokenPoolCaller) GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getCurrentRateLimiterState", remoteChainSelector, customBlockConfirmation)

	outstruct := new(GetCurrentRateLimiterState)
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutboundRateLimiterState = *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)
	outstruct.InboundRateLimiterState = *abi.ConvertType(out[1], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return *outstruct, err

}

func (_LombardTokenPool *LombardTokenPoolSession) GetCurrentRateLimiterState(remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	return _LombardTokenPool.Contract.GetCurrentRateLimiterState(&_LombardTokenPool.CallOpts, remoteChainSelector, customBlockConfirmation)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetCurrentRateLimiterState(remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	return _LombardTokenPool.Contract.GetCurrentRateLimiterState(&_LombardTokenPool.CallOpts, remoteChainSelector, customBlockConfirmation)
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

func (_LombardTokenPool *LombardTokenPoolCaller) GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getFee", arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)

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

func (_LombardTokenPool *LombardTokenPoolSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _LombardTokenPool.Contract.GetFee(&_LombardTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _LombardTokenPool.Contract.GetFee(&_LombardTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)
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

func (_LombardTokenPool *LombardTokenPoolCaller) GetMinBlockConfirmation(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getMinBlockConfirmation")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) GetMinBlockConfirmation() (uint16, error) {
	return _LombardTokenPool.Contract.GetMinBlockConfirmation(&_LombardTokenPool.CallOpts)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetMinBlockConfirmation() (uint16, error) {
	return _LombardTokenPool.Contract.GetMinBlockConfirmation(&_LombardTokenPool.CallOpts)
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

func (_LombardTokenPool *LombardTokenPoolCaller) GetRequiredCCVs(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getRequiredCCVs", localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _LombardTokenPool.Contract.GetRequiredCCVs(&_LombardTokenPool.CallOpts, localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _LombardTokenPool.Contract.GetRequiredCCVs(&_LombardTokenPool.CallOpts, localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)
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

func (_LombardTokenPool *LombardTokenPoolCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _LombardTokenPool.contract.Call(opts, &out, "getTokenTransferFeeConfig", arg0, destChainSelector, arg2, arg3)

	if err != nil {
		return *new(IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolV2TokenTransferFeeConfig)).(*IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_LombardTokenPool *LombardTokenPoolSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _LombardTokenPool.Contract.GetTokenTransferFeeConfig(&_LombardTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3)
}

func (_LombardTokenPool *LombardTokenPoolCallerSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
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

func (_LombardTokenPool *LombardTokenPoolTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_LombardTokenPool *LombardTokenPoolSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.LockOrBurn0(&_LombardTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.LockOrBurn0(&_LombardTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_LombardTokenPool *LombardTokenPoolSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.ReleaseOrMint(&_LombardTokenPool.TransactOpts, releaseOrMintIn)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.ReleaseOrMint(&_LombardTokenPool.TransactOpts, releaseOrMintIn)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn, blockConfirmationRequested)
}

func (_LombardTokenPool *LombardTokenPoolSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.ReleaseOrMint0(&_LombardTokenPool.TransactOpts, releaseOrMintIn, blockConfirmationRequested)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.ReleaseOrMint0(&_LombardTokenPool.TransactOpts, releaseOrMintIn, blockConfirmationRequested)
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

func (_LombardTokenPool *LombardTokenPoolTransactor) SetDynamicConfig(opts *bind.TransactOpts, router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "setDynamicConfig", router, rateLimitAdmin, feeAdmin)
}

func (_LombardTokenPool *LombardTokenPoolSession) SetDynamicConfig(router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.SetDynamicConfig(&_LombardTokenPool.TransactOpts, router, rateLimitAdmin, feeAdmin)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) SetDynamicConfig(router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.SetDynamicConfig(&_LombardTokenPool.TransactOpts, router, rateLimitAdmin, feeAdmin)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) SetMinBlockConfirmation(opts *bind.TransactOpts, minBlockConfirmation uint16) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "setMinBlockConfirmation", minBlockConfirmation)
}

func (_LombardTokenPool *LombardTokenPoolSession) SetMinBlockConfirmation(minBlockConfirmation uint16) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.SetMinBlockConfirmation(&_LombardTokenPool.TransactOpts, minBlockConfirmation)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) SetMinBlockConfirmation(minBlockConfirmation uint16) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.SetMinBlockConfirmation(&_LombardTokenPool.TransactOpts, minBlockConfirmation)
}

func (_LombardTokenPool *LombardTokenPoolTransactor) SetPath(opts *bind.TransactOpts, remoteChainSelector uint64, lChainId [32]byte, allowedCaller []byte) (*types.Transaction, error) {
	return _LombardTokenPool.contract.Transact(opts, "setPath", remoteChainSelector, lChainId, allowedCaller)
}

func (_LombardTokenPool *LombardTokenPoolSession) SetPath(remoteChainSelector uint64, lChainId [32]byte, allowedCaller []byte) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.SetPath(&_LombardTokenPool.TransactOpts, remoteChainSelector, lChainId, allowedCaller)
}

func (_LombardTokenPool *LombardTokenPoolTransactorSession) SetPath(remoteChainSelector uint64, lChainId [32]byte, allowedCaller []byte) (*types.Transaction, error) {
	return _LombardTokenPool.Contract.SetPath(&_LombardTokenPool.TransactOpts, remoteChainSelector, lChainId, allowedCaller)
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

type LombardTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator struct {
	Event *LombardTokenPoolCustomBlockConfirmationInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LombardTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LombardTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolCustomBlockConfirmationInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator{contract: _LombardTokenPool.contract, event: "CustomBlockConfirmationInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
				if err := _LombardTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*LombardTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error) {
	event := new(LombardTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
	if err := _LombardTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LombardTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator struct {
	Event *LombardTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LombardTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LombardTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator{contract: _LombardTokenPool.contract, event: "CustomBlockConfirmationOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
				if err := _LombardTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*LombardTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error) {
	event := new(LombardTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
	if err := _LombardTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
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

type LombardTokenPoolMinBlockConfirmationSetIterator struct {
	Event *LombardTokenPoolMinBlockConfirmationSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LombardTokenPoolMinBlockConfirmationSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LombardTokenPoolMinBlockConfirmationSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LombardTokenPoolMinBlockConfirmationSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LombardTokenPoolMinBlockConfirmationSetIterator) Error() error {
	return it.fail
}

func (it *LombardTokenPoolMinBlockConfirmationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LombardTokenPoolMinBlockConfirmationSet struct {
	MinBlockConfirmation uint16
	Raw                  types.Log
}

func (_LombardTokenPool *LombardTokenPoolFilterer) FilterMinBlockConfirmationSet(opts *bind.FilterOpts) (*LombardTokenPoolMinBlockConfirmationSetIterator, error) {

	logs, sub, err := _LombardTokenPool.contract.FilterLogs(opts, "MinBlockConfirmationSet")
	if err != nil {
		return nil, err
	}
	return &LombardTokenPoolMinBlockConfirmationSetIterator{contract: _LombardTokenPool.contract, event: "MinBlockConfirmationSet", logs: logs, sub: sub}, nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) WatchMinBlockConfirmationSet(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolMinBlockConfirmationSet) (event.Subscription, error) {

	logs, sub, err := _LombardTokenPool.contract.WatchLogs(opts, "MinBlockConfirmationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LombardTokenPoolMinBlockConfirmationSet)
				if err := _LombardTokenPool.contract.UnpackLog(event, "MinBlockConfirmationSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LombardTokenPool *LombardTokenPoolFilterer) ParseMinBlockConfirmationSet(log types.Log) (*LombardTokenPoolMinBlockConfirmationSet, error) {
	event := new(LombardTokenPoolMinBlockConfirmationSet)
	if err := _LombardTokenPool.contract.UnpackLog(event, "MinBlockConfirmationSet", log); err != nil {
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
	CustomBlockConfirmation   bool
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

func (LombardTokenPoolCustomBlockConfirmationInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce615")
}

func (LombardTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac20979328")
}

func (LombardTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x3f1036e85d016a93254a0b1415844f79b85424959d90ae5ad51ce8f4533fe701")
}

func (LombardTokenPoolFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
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

func (LombardTokenPoolMinBlockConfirmationSet) Topic() common.Hash {
	return common.HexToHash("0xa7f8dbba8cdb126ce4a0e7939ec58e0161b70d808b585dd651d68e59d27e11fb")
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
	return common.HexToHash("0x8a8e4c676433747219d2fee4ea128776522bb0177478e1e0a375e880948ed37b")
}

func (LombardTokenPoolPathSet) Topic() common.Hash {
	return common.HexToHash("0x83eda38165c92f401f97217d5ead82ef163d0b716c3979eff4670361bc2dc0c9")
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

	GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

		error)

	GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

		error)

	GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

		error)

	GetLombardConfig(opts *bind.CallOpts) (GetLombardConfig,

		error)

	GetMinBlockConfirmation(opts *bind.CallOpts) (uint16, error)

	GetPath(opts *bind.CallOpts, remoteChainSelector uint64) (LombardTokenPoolPath, error)

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRequiredCCVs(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetSupportedChains(opts *bind.CallOpts) ([]uint64, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	GetTokenDecimals(opts *bind.CallOpts) (uint8, error)

	GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error)

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

	LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error)

	RemovePath(opts *bind.TransactOpts, remoteChainSelector uint64) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error)

	SetMinBlockConfirmation(opts *bind.TransactOpts, minBlockConfirmation uint16) (*types.Transaction, error)

	SetPath(opts *bind.TransactOpts, remoteChainSelector uint64, lChainId [32]byte, allowedCaller []byte) (*types.Transaction, error)

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

	FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*LombardTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*LombardTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*LombardTokenPoolDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*LombardTokenPoolDynamicConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*LombardTokenPoolFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*LombardTokenPoolFeeTokenWithdrawn, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardTokenPoolInboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*LombardTokenPoolInboundRateLimitConsumed, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LombardTokenPoolLockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*LombardTokenPoolLockedOrBurned, error)

	FilterLombardConfigurationSet(opts *bind.FilterOpts, verifier []common.Address, bridge []common.Address, tokenAdapter []common.Address) (*LombardTokenPoolLombardConfigurationSetIterator, error)

	WatchLombardConfigurationSet(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolLombardConfigurationSet, verifier []common.Address, bridge []common.Address, tokenAdapter []common.Address) (event.Subscription, error)

	ParseLombardConfigurationSet(log types.Log) (*LombardTokenPoolLombardConfigurationSet, error)

	FilterMinBlockConfirmationSet(opts *bind.FilterOpts) (*LombardTokenPoolMinBlockConfirmationSetIterator, error)

	WatchMinBlockConfirmationSet(opts *bind.WatchOpts, sink chan<- *LombardTokenPoolMinBlockConfirmationSet) (event.Subscription, error)

	ParseMinBlockConfirmationSet(log types.Log) (*LombardTokenPoolMinBlockConfirmationSet, error)

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
