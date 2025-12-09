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

type IPoolV2TokenTransferFeeConfig struct {
	DestGasOverhead                        uint32
	DestBytesOverhead                      uint32
	DefaultBlockConfirmationFeeUSDCents    uint32
	CustomBlockConfirmationFeeUSDCents     uint32
	DefaultBlockConfirmationTransferFeeBps uint16
	CustomBlockConfirmationTransferFeeBps  uint16
	IsEnabled                              bool
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

var SiloedLockReleaseTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"advancedPoolHooks\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockBox\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"disableTokenTransferFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAvailableTokens\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"lockedTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChainRebalancer\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRebalancer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUnsiloedLiquidity\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"out\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provideLiquidity\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"provideSiloedLiquidity\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.RateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRebalancer\",\"inputs\":[{\"name\":\"newRebalancer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSiloRebalancer\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newRebalancer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"updateSiloDesignations\",\"inputs\":[{\"name\":\"removes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"adds\",\"type\":\"tuple[]\",\"internalType\":\"struct SiloedLockReleaseTokenPool.SiloConfigUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rebalancer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawLiquidity\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawSiloedLiquidity\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainSiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"rebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainUnsiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"amountUnsiloed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SiloRebalancerSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"oldRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnsiloedRebalancerSet\",\"inputs\":[{\"name\":\"oldRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotSiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InsufficientLiquidity\",\"inputs\":[{\"name\":\"availableLiquidity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidChainSelector\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmation\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"LiquidityAmountCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x610120806040523461037f5760c0816162c3803803809161002082856104fb565b83398101031261037f5780516001600160a01b0381169182820361037f5761004a6020820161051e565b6100566040830161052c565b906100636060840161052c565b61007b60a06100746080870161052c565b950161052c565b9433156104ea57600180546001600160a01b03191633179055861580156104d9575b80156104c8575b6104505760805260c05260405163313ce56760e01b8152602081600481895afa6000918161048c575b50610461575b5060a0526001600160a01b0390811660e052600280546001600160a01b0319169282169290921790915516801561045057604051636eb1769f60e11b815230600482015260248101829052602081604481865afa90811561044457600091610412575b506103a75760405191602083019263095ea7b360e01b845282602482015260001960448201526044815261016b6064826104fb565b60008060409586519361017e88866104fb565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082865af13d1561039a573d906001600160401b0382116103845785516101ef9490926101e0601f8201601f1916602001856104fb565b83523d6000602085013e610540565b805180610303575b50506101005251615cb2908161061182396080518181816104b80152818161157201528181611a9401528181611c4101528181611d4a01528181611e240152818161226a015281816123f401528181612feb015281816131fc0152818161325d015281816133090152818161351a015281816136cf01528181613aaa01528181613dba01528181613e070152613f82015260a051818181613c9001528181614eda01528181614f5d015261531c015260c051818181610ebf01528181611b23015281816122f80152818161307a01526135a9015260e051816149350152610100518181816105030152818161151b015281816125470152818161379501528181613ae50152613f2b0152f35b816020918101031261037f576020015180159081150361037f576103285738806101f7565b815162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b916101ef92606091610540565b60405162461bcd60e51b815260206004820152603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527f20746f206e6f6e2d7a65726f20616c6c6f77616e6365000000000000000000006064820152608490fd5b90506020813d60201161043c575b8161042d602093836104fb565b8101031261037f575138610136565b3d9150610420565b6040513d6000823e3d90fd5b630a64406560e11b60005260046000fd5b60ff1660ff821681810361047557506100d3565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d6020116104c0575b816104a8602093836104fb565b8101031261037f576104b99061051e565b90386100cd565b3d915061049b565b506001600160a01b038216156100a4565b506001600160a01b0385161561009d565b639b15e16f60e01b60005260046000fd5b601f909101601f19168101906001600160401b0382119082101761038457604052565b519060ff8216820361037f57565b51906001600160a01b038216820361037f57565b919290156105a25750815115610554575090565b3b1561055d5790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b8251909150156105b55750805190602001fd5b6040519062461bcd60e51b8252602060048301528181519182602483015260005b8381106105f85750508160006044809484010152601f80199101168101030190fd5b602082820181015160448784010152859350016105d656fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a7146140b5575080630a861f2a14613eb2578063181f5a7714613e2b57806321df0da714613de7578063240028e814613d935780632422ac4514613cb457806324f65ee714613c765780632c06340414613bdd5780632d4a148f146139d957806331238ffc14613995578063390775371461349b578063432a6ba314613474578063489a68f214612f575780634c5ef0ed14612f1057806362ddd3c414612e895780636600f92c14612d995780636cfd155314612cf85780636d9d216c146128df5780637437ff9f1461289b57806379ba5097146127ee5780638632d5cc146127b95780638926f54f1461277357806389720a62146126b95780638da5cb5b146126925780639a4575b9146121fb578063a42a7b8b146120b2578063acfecf9114611fc0578063af0e58b914611fa2578063b1c71c65146119fd578063b7946580146119c0578063c4bffe2b146118b3578063c7230a60146116f1578063ce3c75281461146e578063d8aa3f4014611334578063dc04fa1f14610ee3578063dc0bd97114610e9f578063dcbd41bc14610cb5578063e8a1da1714610648578063eb521a4c1461041b578063f1e73399146103ef578063f2fde38b1461033a5763fdf16875146101ef57600080fd5b34610337576060600319360112610337576102086141f8565b90610211614275565b604435926001600160a01b0384168085036103335761022e615061565b6001600160a01b038216801561030b5794610305917fba9213054b14c2e884f779120bb196f0735cef27140498a9d26117eeab77a1179596600254907fffffffffffffffffffff0000000000000000000000000000000000000000000075ffff00000000000000000000000000000000000000008860a01b16921617176002557fffffffffffffffffffffffff000000000000000000000000000000000000000060095416176009556040519384938491604091949361ffff6001600160a01b039283606087019816865216602085015216910152565b0390a180f35b6004857f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b8380fd5b80fd5b5034610337576020600319360112610337576001600160a01b0361035c6141f8565b610364615061565b163381146103c757807fffffffffffffffffffffffff00000000000000000000000000000000000000008354161782556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b503461033757602060031936011261033757602061041361040e61423a565b614d5a565b604051908152f35b5034610337576020600319360112610337576004358015610620576001600160a01b0361044783614896565b1633036105f457818052600d60205260408220600181015460a01c60ff16156105df57610475828254614879565b90555b6040517f23b872dd0000000000000000000000000000000000000000000000000000000060208201523360248201523060448201526064810182905282907f0000000000000000000000000000000000000000000000000000000000000000906104f9906104f381608481015b03601f19810183528261438f565b82615798565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016803b156105db576040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b039290921660048301526024820184905282908290604490829084905af180156105d0576105b7575b50506040519082825260208201527f569a440e6842b5e5a7ac02286311855f5a0b81b9390909e552e82aaf02c9e9bf60403392a280f35b816105c19161438f565b6105cc578138610580565b5080fd5b6040513d84823e3d90fd5b8280fd5b506105ec81600b54614879565b600b55610478565b6024827f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b6004827fa90c0d19000000000000000000000000000000000000000000000000000000008152fd5b50346103375760406003193601126103375760043567ffffffffffffffff81116105cc5761067a903690600401614468565b9060243567ffffffffffffffff8111610333579061069d84923690600401614468565b9390916106a8615061565b83905b828210610b1a5750505081927ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b81811015610b16578060051b83013585811215610b0d57830161012081360312610b0d576040519461070f86614373565b813567ffffffffffffffff81168103610b11578652602082013567ffffffffffffffff81116105cc5782019436601f870112156105cc57853595610752876148d8565b96610760604051988961438f565b80885260208089019160051b83010190368211610b0d5760208301905b828210610ada575050505060208701958652604083013567ffffffffffffffff81116105db576107b09036908501614405565b91604088019283526107da6107c83660608701614ca5565b9460608a0195865260c0369101614ca5565b956080890196875283515115610ab2576107fe67ffffffffffffffff8a51166158cb565b15610a7b5767ffffffffffffffff8951168252600760205260408220610825865182615350565b610833885160028301615350565b6004855191019080519067ffffffffffffffff8211610a4e576108568354614b03565b601f8111610a13575b50602090601f83116001146109b05761088f92918691836109a5575b50506000198260011b9260031b1c19161790565b90555b815b885180518210156108c957906108c36001926108bc8367ffffffffffffffff8f511692614aef565b519061509f565b01610894565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c293919997509561099767ffffffffffffffff600197969498511692519351915161096361092e604051968796875261010060208801526101008701906141b7565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a10193929091936106de565b015190508e8061087b565b8386528186209190601f198416875b8181106109fb57509084600195949392106109e2575b505050811b019055610892565b015160001960f88460031b161c191690558d80806109d5565b929360206001819287860151815501950193016109bf565b610a3e9084875260208720601f850160051c81019160208610610a44575b601f0160051c0190614d30565b8d61085f565b9091508190610a31565b6024857f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60248267ffffffffffffffff8b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b6004827f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b813567ffffffffffffffff8111610b0957602091610afe8392833691890101614405565b81520191019061077d565b8680fd5b8480fd5b600080fd5b8380f35b9267ffffffffffffffff610b3c610b378486889a9699979a61483a565b614798565b1691610b47836155e3565b15610c89578284526007602052610b6360056040862001615580565b94845b8651811015610b9c576001908587526007602052610b9560056040892001610b8e838b614aef565b51906156dc565b5001610b66565b5093969290945094909480875260076020526005604088208881558860018201558860028201558860038201558860048201610bd88154614b03565b80610c48575b5050500180549088815581610c2a575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a1019091949392946106ab565b885260208820908101905b81811015610bee57888155600101610c35565b601f8111600114610c5e5750555b888a80610bde565b81835260208320610c7991601f01861c810190600101614d30565b8082528160208120915555610c56565b602484847f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50346103375760206003193601126103375760043567ffffffffffffffff81116105cc57610ce79036906004016144c3565b6001600160a01b036009541633141580610e8a575b610e5e57825b818110610d0d578380f35b610d18818385614c3a565b67ffffffffffffffff610d2a82614798565b1690610d43826000526006602052604060002054151590565b15610e3257907f41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb60e083610df2610dcc602060019897018b610d8482614c4a565b15610df9578790526003602052610dab60408d20610da53660408801614ca5565b90615350565b868c526004602052610dc760408d20610da53660a08801614ca5565b614c4a565b916040519215158352610de56020840160408301614cec565b60a0608084019101614cec565ba201610d02565b60026040828a610dc794526007602052610e1b828220610da536858c01614ca5565b8a815260076020522001610da53660a08801614ca5565b602486837f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b6024837f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b506001600160a01b0360015416331415610cfc565b503461033757806003193601126103375760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346103375760406003193601126103375760043567ffffffffffffffff81116105cc57610f159036906004016144c3565b60243567ffffffffffffffff811161033357610f35903690600401614468565b919092610f40615061565b845b828110610fac57505050825b818110610f59578380f35b8067ffffffffffffffff610f73610b37600194868861483a565b16808652600a6020528560408120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee88680a201610f4e565b610fba610b37828585614c3a565b610fc5828585614c3a565b90602082019060e0830190610fd982614c4a565b156112ff5760a0840161271061ffff610ff183614c57565b1610156112f05760c085019161271061ffff61100c85614c57565b1610156112b85763ffffffff61102186614c66565b16156112835767ffffffffffffffff1694858c52600a60205260408c2061104786614c66565b63ffffffff1690805490604084019161105f83614c66565b60201b67ffffffff000000001693606086019461107b86614c66565b60401b6bffffffff000000000000000016966080019661109a88614c66565b60601b6fffffffff00000000000000000000000016916110b98a614c57565b60801b71ffff0000000000000000000000000000000016936110da8c614c57565b60901b73ffff00000000000000000000000000000000000016957fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016177fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16177fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff16177fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff16171717815561118d87614c4a565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016179055604051966111de90614c77565b63ffffffff1687526111ef90614c77565b63ffffffff16602087015261120390614c77565b63ffffffff16604086015261121790614c77565b63ffffffff16606085015261122b90614297565b61ffff16608084015261123d90614297565b61ffff1660a083015261124f90614268565b151560c082015260e07ffae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa104191a2600101610f42565b7f12332265000000000000000000000000000000000000000000000000000000008c5267ffffffffffffffff1660045260248bfd5b60248c61ffff6112c786614c57565b7f95f3517a00000000000000000000000000000000000000000000000000000000835216600452fd5b8a61ffff6112c7602493614c57565b7f12332265000000000000000000000000000000000000000000000000000000008a5267ffffffffffffffff16600452602489fd5b50346103375760806003193601126103375761134e6141f8565b50611357614251565b61135f614286565b5060643567ffffffffffffffff81116105db579167ffffffffffffffff60409261138f60e09536906004016142a6565b50508260c0855161139f81614357565b82815282602082015282878201528260608201528260808201528260a08201520152168152600a60205220604051906113d782614357565b5461ffff818163ffffffff82169485815263ffffffff60208201818560201c1681528160408401818760401c168152816060860193818960601c16855260ff60c060808901988a8c60801c168a528a60a082019c60901c168c52019b60a01c1615158b526040519b8c52511660208b0152511660408901525116606087015251166080850152511660a083015251151560c0820152f35b50346103375760406003193601126103375761148861423a565b60243567ffffffffffffffff8216808452600d60205260ff600160408620015460a01c161580156116e9575b6116be578115611696576001600160a01b036114cf84614896565b16330361166a578352600d60205260408320600181015460a01c60ff1680156116625781545b80841161163257501561161d5761150d8282546147ad565b90555b826001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016803b156105cc576040517f69328dec0000000000000000000000000000000000000000000000000000000081526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166004820152602481018490523360448201529082908290606490829084905af180156105d057611608575b50506040805167ffffffffffffffff9093168352602083019190915233917f58fca2457646a9f47422ab9eb9bff90cef88cd8b8725ab52b1d17baa392d784e91819081015b0390a280f35b816116129161438f565b6105db5782386115bd565b5061162a81600b546147ad565b600b55611510565b85846044927fa17e11d5000000000000000000000000000000000000000000000000000000008352600452602452fd5b600b546114f5565b6024847f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b6004847fa90c0d19000000000000000000000000000000000000000000000000000000008152fd5b7f46f5f12b000000000000000000000000000000000000000000000000000000008452600452602483fd5b5080156114b4565b50346103375760406003193601126103375760043567ffffffffffffffff81116105cc57611723903690600401614468565b9061172c614224565b611734615061565b835b838110611741578480f35b8060206001600160a01b0361176161175c602495898961483a565b6147e9565b16604051938480927f70a082310000000000000000000000000000000000000000000000000000000082523060048301525afa80156118a8578690611872575b60019250806117b2575b5001611736565b61181f6001600160a01b036117cb61175c858a8a61483a565b6040517fa9059cbb0000000000000000000000000000000000000000000000000000000060208201526001600160a01b038816602482015260448082018690528152911661181a60648361438f565b615798565b6001600160a01b0361183561175c84898961483a565b60405192835216907f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e60206001600160a01b03871692a3386117ab565b509060203d81116118a1575b611888818361438f565b60208260009281010312610337575090600191516117a1565b503d61187e565b6040513d88823e3d90fd5b5034610337578060031936011261033757604051906005548083528260208101600584526020842092845b8181106119a75750506118f39250038361438f565b8151611917611901826148d8565b9161190f604051938461438f565b8083526148d8565b91601f19602083019301368437805b8451811015611958578067ffffffffffffffff61194560019388614aef565b51166119518286614aef565b5201611926565b50925090604051928392602084019060208552518091526040840192915b818110611984575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101611976565b84548352600194850194879450602090930192016118de565b5034610337576020600319360112610337576119f96119e56119e061423a565b614c18565b6040519182916020835260208301906141b7565b0390f35b50346103375760606003193601126103375760043567ffffffffffffffff81116105cc5760a060031982360301126105cc57611a37614275565b60443567ffffffffffffffff81116103335792611a5b611a7b9436906004016142a6565b9490611a65614ad6565b50611a7384866004016152af565b9536916143ce565b5060848301611a89816147e9565b6001600160a01b03807f000000000000000000000000000000000000000000000000000000000000000016911603611f655750602483019277ffffffffffffffff00000000000000000000000000000000611ae385614798565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611ed8578391611f36575b50611f0e5767ffffffffffffffff611b6a85614798565b16611b82816000526006602052604060002054151590565b15611ee35760206001600160a01b0360025416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015611ed8578390611e98575b6001600160a01b039150163303611e6c57606401359161ffff611bf986856147ad565b91168015611dc55761ffff60025460a01c169081611ce3575b505050506119e0611c29611cd994611ca8936147ad565b92611c3381614798565b604080516001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016815233602082015290810186905267ffffffffffffffff91909116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1090606090a2614798565b90611cb1615315565b60405192611cbe8461433b565b83526020830152604051928392604084526040840190614499565b9060208301520390f35b818110611d96575050611cd994611ca8936119e0937f61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac2097932884611c2995611d72604067ffffffffffffffff611d358d614798565b16958681526003602052206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391615980565b604080516001600160a01b039290921682526020820192909252a293819650611c12565b7f7911d95b00000000000000000000000000000000000000000000000000000000845260045260245250604490fd5b50611cd994611ca8936119e0937fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894484611c2995611e4c604067ffffffffffffffff611e0f8d614798565b16958681526007602052206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391615980565b604080516001600160a01b039290921682526020820192909252a26147ad565b6024827f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d602011611ed0575b81611eb26020938361438f565b810103126105db57611ecb6001600160a01b03916148f0565b611bd6565b3d9150611ea5565b6040513d85823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008352600452602482fd5b6004827f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b611f58915060203d602011611f5e575b611f50818361438f565b810190615049565b38611b53565b503d611f46565b906001600160a01b03611f796024936147e9565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835216600452fd5b50346103375780600319360112610337576020600b54604051908152f35b50346103375767ffffffffffffffff611fd836614423565b929091611fe3615061565b1691611ffc836000526006602052604060002054151590565b15610c8957828452600760205261202b6005604086200161201e3684866143ce565b60208151910120906156dc565b1561206a57907f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d7691611602604051928392602084526020840191614904565b826120ae836040519384937f74f23c7c0000000000000000000000000000000000000000000000000000000085526004850152604060248501526044840191614904565b0390fd5b50346103375760206003193601126103375767ffffffffffffffff6120d561423a565b16815260076020526120ec60056040832001615580565b8051601f196121136120fd836148d8565b9261210b604051948561438f565b8084526148d8565b01835b8181106121ea575050825b8251811015612167578061213760019285614aef565b518552600860205261214b60408620614b56565b6121558285614aef565b526121608184614aef565b5001612121565b81846040519182916020830160208452825180915260408401602060408360051b870101940192905b82821061219f57505050500390f35b919360206121da827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0600195979984950301865288516141b7565b9601920192018594939192612190565b806060602080938601015201612116565b5034610337576020600319360112610337576004359067ffffffffffffffff82116103375760a0600319833603011261033757612236614ad6565b5061223f614ad6565b5060209080604051612251848261438f565b526084830161225f816147e9565b6001600160a01b03807f000000000000000000000000000000000000000000000000000000000000000016911603611f655750602483019277ffffffffffffffff000000000000000000000000000000006122b985614798565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015283816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611ed8578391612675575b50611f0e5767ffffffffffffffff61233f85614798565b16612357816000526006602052604060002054151590565b15611ee357836001600160a01b0360025416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015611ed857839061263a575b6001600160a01b039150163303611e6c5760640135908061260d57508067ffffffffffffffff6123d985614798565b16806000526007845261241c60406000206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016938491615980565b604080516001600160a01b0384168152602081018590527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789449190a267ffffffffffffffff61246985614798565b604080516001600160a01b03851681523360208201529081018590529116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1090606090a26124ba6119e085614798565b936124c3615315565b604051956124d08761433b565b86528486015267ffffffffffffffff6124e882614798565b16600052600d845260ff60016040600020015460a01c166000146125f8578067ffffffffffffffff61251c61253b93614798565b16600052600d85526040600020612534858254614879565b9055614798565b505b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691823b15610b11576040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b039290921660048301526024820152906000908290604490829084905af180156125ec576125da575b506119f9604051928284938452830190614499565b60006125e59161438f565b60006125c5565b6040513d6000823e3d90fd5b5061260582600b54614879565b600b5561253d565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526011600452fd5b508381813d831161266e575b612650818361438f565b810103126105db576126696001600160a01b03916148f0565b6123aa565b503d612646565b61268c9150843d8611611f5e57611f50818361438f565b38612328565b503461033757806003193601126103375760206001600160a01b0360015416604051908152f35b50346103375760c0600319360112610337576126d36141f8565b6126db614251565b9060643561ffff811681036103335760843567ffffffffffffffff8111610b0d5761270a9036906004016142a6565b9160a435936002851015610b09576127259560443591614925565b90604051918291602083016020845282518091526020604085019301915b818110612751575050500390f35b82516001600160a01b0316845285945060209384019390920191600101612743565b50346103375760206003193601126103375760206127af67ffffffffffffffff61279b61423a565b166000526006602052604060002054151590565b6040519015158152f35b50346103375760206003193601126103375760206127dd6127d861423a565b614896565b6001600160a01b0360405191168152f35b503461033757806003193601126103375780546001600160a01b0381163303612873577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551682556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b5034610337578060031936011261033757600254600954604080516001600160a01b03808516825260a09490941c61ffff1660208201529290911690820152606090f35b50346103375760406003193601126103375760043567ffffffffffffffff81116105cc57612911903690600401614468565b906024359067ffffffffffffffff821161033357366023830112156103335781600401359267ffffffffffffffff8411610b0d576024830192602436918660061b010111610b0d57612961615061565b845b818110612bd257505050825b82811061297a578380f35b67ffffffffffffffff612991610b37838686614886565b16158015612b9f575b8015612b7e575b612b3b576001600160a01b036129c360206129bd848787614886565b016147e9565b1615612b135780612aa884846001600160a01b038060018a67ffffffffffffffff612a22610b378a6129fc60206129bd889f8d8d614886565b9960405195612a0a8761431f565b865287602087019b168b526040860199878b52614886565b168c52600d60205260408c209051815501935116167fffffffffffffffffffffffff00000000000000000000000000000000000000008354161782555115157fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b7f180c6940bd64ba8f75679203ca32f8be2f629477a3307b190656e4b14dd5ddeb612ad7610b37838787614886565b612ae760206129bd858989614886565b6040805167ffffffffffffffff9390931683526001600160a01b0391909116602083015290a10161296f565b6004847f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b610b37612b55916024959467ffffffffffffffff94614886565b7fd9a9cd6800000000000000000000000000000000000000000000000000000000835216600452fd5b50612b9967ffffffffffffffff61279b610b37848787614886565b156129a1565b5067ffffffffffffffff612bb7610b37838686614886565b168452600d60205260ff600160408620015460a01c1661299a565b67ffffffffffffffff612be9610b3783858761483a565b168652600d60205260ff600160408820015460a01c1615612cb5578067ffffffffffffffff612c1e610b37600194868861483a565b168752600d6020527f7b5efb3f8090c5cfd24e170b667d0e2b6fdc3db6540d75b86d5b6655ba00eb936040882054612c5881600b54614879565b600b5567ffffffffffffffff612c72610b3785888a61483a565b168952600d6020528884604082208281550155612c93610b3784878961483a565b6040805167ffffffffffffffff9290921682526020820192909252a101612963565b612ccf610b3767ffffffffffffffff92889460249661483a565b7f46f5f12b00000000000000000000000000000000000000000000000000000000835216600452fd5b5034610337576020600319360112610337577f66b1c1bdec8b60a3442bb25b5b6cd6fff3d0eceb6f5390be8e2f82a8ad39b2346001600160a01b03612d3b6141f8565b612d43615061565b610305600c54918381167fffffffffffffffffffffffff0000000000000000000000000000000000000000841617600c5560405193849316839092916001600160a01b0360209181604085019616845216910152565b503461033757604060031936011261033757612db361423a565b67ffffffffffffffff612dc4614224565b91612dcd615061565b1690818352600d602052600160408420019081549160ff8360a01c1615612e5d5780547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283169081179091556040805192909316825260208201527f01efd4cd7dd64263689551000d4359d6559c839f39b773b1df3fd19ff060cf5f9181908101611602565b602485857f46f5f12b000000000000000000000000000000000000000000000000000000008252600452fd5b503461033757612e9836614423565b612ea493929193615061565b67ffffffffffffffff8216612ec6816000526006602052604060002054151590565b15612ee55750612ee29293612edc9136916143ce565b9061509f565b80f35b7f1e670e4b000000000000000000000000000000000000000000000000000000008452600452602483fd5b503461033757604060031936011261033757612f2a61423a565b906024359067ffffffffffffffff82116103375760206127af84612f513660048701614405565b906147fd565b5034610337576040600319360112610337576004359067ffffffffffffffff8211610337578160040161010060031984360301126105cc57612f97614275565b9180604051612fa5816142d4565b52612fd2612fc8612fc3612fbc60c4880186614747565b36916143ce565b614e66565b6064860135614f5a565b9260848501612fe0816147e9565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116036134605750602485019277ffffffffffffffff0000000000000000000000000000000061303a85614798565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156133e3578491613441575b50613419576130b884614798565b67ffffffffffffffff81166130da816000526006602052604060002054151590565b156133ee57506002546040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff929092166004830152336024830152602090829060449082906001600160a01b03165afa9081156133e35784916133c4575b50156133985761315384614798565b9061316960a4880192612f51612fbc8585614747565b1561335157505061325761325160446020977ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc095878961ffff67ffffffffffffffff981615156000146132bf578061322460408a6131e77f0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce61596614798565b16958681528f60049052206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391615980565b604080516001600160a01b039290921682526020820192909252a25b019461324b866147e9565b50614798565b936147e9565b604080517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b039081168252336020830152909216908201526060810185905292169180608081015b0390a2806040516132b6816142d4565b52604051908152f35b80613331600260408b6132f27f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c97614798565b169687815260206007905220016001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391615980565b604080516001600160a01b039290921682526020820192909252a2613240565b61335b9250614747565b6120ae6040519283927f24eb47e5000000000000000000000000000000000000000000000000000000008452602060048501526024840191614904565b6024837f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b6133dd915060203d602011611f5e57611f50818361438f565b38613144565b6040513d86823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008552600452602484fd5b6004837f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b61345a915060203d602011611f5e57611f50818361438f565b386130aa565b826001600160a01b03611f796024936147e9565b503461033757806003193601126103375760206001600160a01b03600c5416604051908152f35b50346103375760206003193601126103375760043567ffffffffffffffff81116105cc578060040191610100600319833603011261033757806040516134e0816142d4565b526135016134f7612fc3612fbc60c4860187614747565b6064840135614f5a565b906084830161350f816147e9565b6001600160a01b03807f000000000000000000000000000000000000000000000000000000000000000016911603611f655750602483019077ffffffffffffffff0000000000000000000000000000000061356983614798565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156105d0578291613976575b5061394e576135e782614798565b67ffffffffffffffff8116613609816000526006602052604060002054151590565b15611ee357506002546040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff929092166004830152336024830152602090829060449082906001600160a01b03165afa9081156105d057829161392f575b50156139035761368282614798565b61369760a4860191612f51612fbc848a614747565b156138f9575082919067ffffffffffffffff6136b283614798565b1680825260076020526136f7600260408420016001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016958691615980565b604080516001600160a01b0386168152602081018790527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a267ffffffffffffffff61374483614798565b168152600d60205260408120600181015460a01c60ff1680156138f15781545b8087116138c15750156138ac5761377c8582546147ad565b905561378782614798565b505b60446001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169501946137c1866147e9565b813b156105db576040517f69328dec0000000000000000000000000000000000000000000000000000000081526001600160a01b038681166004830152602482018890529190911660448201529082908290606490829084905af180156105d057613897575b505067ffffffffffffffff6020946132a6856138666132517ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc096614798565b604080516001600160a01b039889168152336020820152979091169087015260608601529116929081906080820190565b6138a282809261438f565b6103375780613827565b506138b984600b546147ad565b600b55613789565b83876044927fa17e11d5000000000000000000000000000000000000000000000000000000008352600452602452fd5b600b54613764565b61335b9086614747565b807f728fe07b000000000000000000000000000000000000000000000000000000006024925233600452fd5b613948915060203d602011611f5e57611f50818361438f565b38613673565b807f53ad11d80000000000000000000000000000000000000000000000000000000060049252fd5b61398f915060203d602011611f5e57611f50818361438f565b386135d9565b50346103375760206003193601126103375760ff6001604060209367ffffffffffffffff6139c161423a565b168152600d855220015460a01c166040519015158152f35b5034610337576040600319360112610337576139f361423a565b60243567ffffffffffffffff8216808452600d60205260ff600160408620015460a01c16158015613bd5575b6116be578115611696576001600160a01b03613a3a84614896565b16330361166a578352600d60205260408320600181015460a01c60ff1615613bc057613a67828254614879565b90555b6040517f23b872dd0000000000000000000000000000000000000000000000000000000060208201523360248201523060448201526064810182905283907f000000000000000000000000000000000000000000000000000000000000000090613adb906104f381608481016104e5565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016803b156105db576040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b039290921660048301526024820184905282908290604490829084905af180156105d057613bab575b50506040805167ffffffffffffffff9093168352602083019190915233917f569a440e6842b5e5a7ac02286311855f5a0b81b9390909e552e82aaf02c9e9bf9181908101611602565b81613bb59161438f565b6105db578238613b62565b50613bcd81600b54614879565b600b55613a6a565b508015613a1f565b50346103375760c060031936011261033757613bf76141f8565b50613c00614251565b613c0861420e565b506084359161ffff831683036103375760a4359067ffffffffffffffff82116103375760a063ffffffff8061ffff613c4f8888613c483660048b016142a6565b5050614608565b95926040979194975197885216602087015216604085015216606083015215156080820152f35b5034610337578060031936011261033757602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461033757604060031936011261033757613cce61423a565b60243591821515830361033757610140613d91613ceb8585614585565b613d4160409392935180946fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b60a08301906fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565bf35b5034610337576020600319360112610337576020613daf6141f8565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b503461033757806003193601126103375760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b5034610337578060031936011261033757506119f9604051613e4e60608261438f565b602481527f53696c6f65644c6f636b52656c65617365546f6b656e506f6f6c20312e372e3060208201527f2d6465760000000000000000000000000000000000000000000000000000000060408201526040519182916020835260208301906141b7565b50346103375760206003193601126103375760043590811561408d576001600160a01b03613edf82614896565b16330361406157808052600d60205260408120600181015460a01c60ff1680156140595781545b80851161402957501561401457613f1e8382546147ad565b90555b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016803b156105cc576040517f69328dec0000000000000000000000000000000000000000000000000000000081526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166004820152602481018490523360448201529082908290606490829084905af180156105d057614004575b50906040519082825260208201527f58fca2457646a9f47422ab9eb9bff90cef88cd8b8725ab52b1d17baa392d784e60403392a280f35b8161400e9161438f565b38613fcd565b5061402182600b546147ad565b600b55613f21565b83856044927fa17e11d5000000000000000000000000000000000000000000000000000000008352600452602452fd5b600b54613f06565b807f8e4a23d6000000000000000000000000000000000000000000000000000000006024925233600452fd5b807fa90c0d190000000000000000000000000000000000000000000000000000000060049252fd5b9050346105cc5760206003193601126105cc576004357fffffffff0000000000000000000000000000000000000000000000000000000081168091036105db57602092507faff2afbf00000000000000000000000000000000000000000000000000000000811490811561418d575b8115614163575b8115614139575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438614132565b7f0e64dd29000000000000000000000000000000000000000000000000000000008114915061412b565b7f331710310000000000000000000000000000000000000000000000000000000081149150614124565b919082519283825260005b8481106141e3575050601f19601f8460006020809697860101520116010190565b806020809284010151828286010152016141c2565b600435906001600160a01b0382168203610b1157565b606435906001600160a01b0382168203610b1157565b602435906001600160a01b0382168203610b1157565b6004359067ffffffffffffffff82168203610b1157565b6024359067ffffffffffffffff82168203610b1157565b35908115158203610b1157565b6024359061ffff82168203610b1157565b6044359061ffff82168203610b1157565b359061ffff82168203610b1157565b9181601f84011215610b115782359167ffffffffffffffff8311610b115760208381860195010111610b1157565b6020810190811067ffffffffffffffff8211176142f057604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6060810190811067ffffffffffffffff8211176142f057604052565b6040810190811067ffffffffffffffff8211176142f057604052565b60e0810190811067ffffffffffffffff8211176142f057604052565b60a0810190811067ffffffffffffffff8211176142f057604052565b90601f601f19910116810190811067ffffffffffffffff8211176142f057604052565b67ffffffffffffffff81116142f057601f01601f191660200190565b9291926143da826143b2565b916143e8604051938461438f565b829481845281830111610b11578281602093846000960137010152565b9080601f83011215610b1157816020614420933591016143ce565b90565b906040600319830112610b115760043567ffffffffffffffff81168103610b1157916024359067ffffffffffffffff8211610b1157614464916004016142a6565b9091565b9181601f84011215610b115782359167ffffffffffffffff8311610b11576020808501948460051b010111610b1157565b6144209160206144b283516040845260408401906141b7565b9201519060208184039101526141b7565b9181601f84011215610b115782359167ffffffffffffffff8311610b11576020808501948460081b010111610b1157565b6040519061450182614373565b60006080838281528260208201528260408201528260608201520152565b9060405161452c81614373565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff916145976144f4565b506145a06144f4565b506145d4571660005260076020526040600020906144206145c860026145cd6145c88661451f565b614de1565b940161451f565b16908160005260036020526145ef6145c8604060002061451f565b9160005260046020526144206145c8604060002061451f565b67ffffffffffffffff16600052600a60205260406000206040519061462c82614357565b549263ffffffff84168252602082019363ffffffff8160201c168552604083019063ffffffff8160401c1682526060840163ffffffff8260601c168152608085019561ffff8360801c16875260ff60a087019361ffff8160901c16855260a01c1615801560c088015261472e5761ffff16806146c55750505063ffffffff808061ffff9351169451169551169351169193929190600190565b919550915061ffff60025460a01c16908181106146fe57505063ffffffff808061ffff9351169451169551169351169193929190600190565b7f7911d95b0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b5050505092505050600090600090600090600090600090565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610b11570180359067ffffffffffffffff8211610b1157602001918136038313610b1157565b3567ffffffffffffffff81168103610b115790565b919082039182116147ba57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b356001600160a01b0381168103610b115790565b9067ffffffffffffffff61442092166000526007602052600560406000200190602081519101209060019160005201602052604060002054151590565b919081101561484a5760051b0190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b919082018092116147ba57565b919081101561484a5760061b0190565b67ffffffffffffffff16600052600d60205260016040600020015460ff8160a01c166148cc57506001600160a01b03600c541690565b6001600160a01b031690565b67ffffffffffffffff81116142f05760051b60200190565b51906001600160a01b0382168203610b1157565b601f8260209493601f19938186528686013760008582860101520116010190565b9593919294906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016958615614ab4576149c49467ffffffffffffffff61ffff936001600160a01b036040519b7f89720a62000000000000000000000000000000000000000000000000000000008d521660048c01521660248a0152604489015216606487015260c0608487015260c4860191614904565b916002821015614a85578380600094819460a483015203915afa9081156125ec576000916149f0575090565b3d8083833e6149ff818361438f565b8101906020818303126105db5780519067ffffffffffffffff8211610333570181601f820112156105db57805190614a36826148d8565b93614a44604051958661438f565b82855260208086019360051b8301019384116103375750602001905b828210614a6d5750505090565b60208091614a7a846148f0565b815201910190614a60565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b5050505050505050604051614aca60208261438f565b60008152600036813790565b60405190614ae38261433b565b60606020838281520152565b805182101561484a5760209160051b010190565b90600182811c92168015614b4c575b6020831014614b1d57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691614b12565b9060405191826000825492614b6a84614b03565b8084529360018116908115614bd85750600114614b91575b50614b8f9250038361438f565b565b90506000929192526020600020906000915b818310614bbc575050906020614b8f9282010138614b82565b6020919350806001915483858901015201910190918492614ba3565b60209350614b8f9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138614b82565b67ffffffffffffffff1660005260076020526144206004604060002001614b56565b919081101561484a5760081b0190565b358015158103610b115790565b3561ffff81168103610b115790565b3563ffffffff81168103610b115790565b359063ffffffff82168203610b1157565b35906fffffffffffffffffffffffffffffffff82168203610b1157565b9190826060910312610b1157604051614cbd8161431f565b6040614ce7818395614cce81614268565b8552614cdc60208201614c88565b602086015201614c88565b910152565b6fffffffffffffffffffffffffffffffff614d2a60408093614d0d81614268565b1515865283614d1e60208301614c88565b16602087015201614c88565b16910152565b818110614d3b575050565b60008155600101614d30565b818102929181159184041417156147ba57565b67ffffffffffffffff16614d7b816000526006602052604060002054151590565b15614db45780600052600d60205260ff60016040600020015460a01c16614da35750600b5490565b600052600d60205260406000205490565b7fd9a9cd680000000000000000000000000000000000000000000000000000000060005260045260246000fd5b614de96144f4565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff8083511691614e466020850193614e40614e3363ffffffff875116426147ad565b8560808901511690614d47565b90614879565b80821015614e5f57505b16825263ffffffff4216905290565b9050614e50565b80518015614ed657602003614e98578051602082810191830183900312610b1157519060ff8211614e98575060ff1690565b6120ae906040519182917f953576f70000000000000000000000000000000000000000000000000000000083526020600484015260248301906141b7565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff82116147ba57565b60ff16604d81116147ba57600a0a90565b8115614f2b570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff811692828414615042578284116150185790614f9f91614efc565b91604d60ff8416118015614ffd575b614fc757505090614fc161442092614f10565b90614d47565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b5061500783614f10565b8015614f2b57600019048411614fae565b61502191614efc565b91604d60ff841611614fc75750509061503c61442092614f10565b90614f21565b5050505090565b90816020910312610b1157518015158103610b115790565b6001600160a01b0360015416330361507557565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b908051156152855767ffffffffffffffff815160208301209216918260005260076020526150d481600560406000200161592b565b156152415760005260086020526040600020815167ffffffffffffffff81116142f0576151018254614b03565b601f811161520f575b506020601f8211600114615185579161515f827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea95936151759560009161517a575b506000198260011b9260031b1c19161790565b90556040519182916020835260208301906141b7565b0390a2565b90508401513861514c565b601f1982169083600052806000209160005b8181106151f75750926151759492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9896106151de575b5050811b0190556119e5565b85015160001960f88460031b161c1916905538806151d2565b9192602060018192868a015181550194019201615197565b61523b90836000526020600020601f840160051c81019160208510610a4457601f0160051c0190614d30565b3861510a565b50906120ae6040519283927f393b8ad200000000000000000000000000000000000000000000000000000000845260048401526040602484015260448301906141b7565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b9061ffff809167ffffffffffffffff6152ca60208601614798565b166000908152600a602052604090209082161561530b575460901c165b1690811561530457612710916060615300920135614d47565b0490565b5050600090565b5460801c166152e7565b60405160ff7f00000000000000000000000000000000000000000000000000000000000000001660208201526020815261442060408261438f565b8151919291156154d2576fffffffffffffffffffffffffffffffff6040840151166fffffffffffffffffffffffffffffffff6020850151161061546f57614b8f91925b805182547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff0000000000000000000000000000000000000000161782556020810151825460409290920151608090811b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff9290921691821760018501557fffffffffffffffffffffffff0000000000000000000000000000000000000000909216174290911b73ffffffff0000000000000000000000000000000016179055565b6064836154d0604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff60408401511615801590615561575b61550057614b8f9192615393565b6064836154d0604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff60208401511615156154f2565b906040519182815491828252602082019060005260206000209260005b8181106155b2575050614b8f9250038361438f565b845483526001948501948794506020909301920161559d565b805482101561484a5760005260206000200190600090565b60008181526006602052604090205480156153045760001981018181116147ba576005549060001982019182116147ba5781810361568b575b505050600554801561565c57600019016156378160056155cb565b60001982549160031b1b19169055600555600052600660205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6156c461569c6156ad9360056155cb565b90549060031b1c92839260056155cb565b81939154906000199060031b92831b921b19161790565b9055600052600660205260406000205538808061561c565b906001820191816000528260205260406000205480151560001461578f5760001981018181116147ba5782549060001982019182116147ba57818103615758575b5050508054801561565c57600019019061573782826155cb565b60001982549160031b1b191690555560005260205260006040812055600190565b6157786157686156ad93866155cb565b90549060031b1c928392866155cb565b90556000528360205260406000205538808061571d565b50505050600090565b6001600160a01b0361581a9116916040926000808551936157b9878661438f565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602086015260208151910182855af13d156158c3573d916157fe836143b2565b9261580b8751948561438f565b83523d6000602085013e615bd9565b8051908161582757505050565b602080615838938301019101615049565b156158405750565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606091615bd9565b8060005260066020526040600020541560001461592557600554680100000000000000008110156142f05761590c6156ad82600185940160055560056155cb565b9055600554906000526006602052604060002055600190565b50600090565b600082815260018201602052604090205461530457805490680100000000000000008210156142f057826159696156ad8460018096018555846155cb565b905580549260005201602052604060002055600190565b9182549060ff8260a01c16158015615bd1575b615bcb576fffffffffffffffffffffffffffffffff821691600185019081546159d863ffffffff6fffffffffffffffffffffffffffffffff83169360801c16426147ad565b9081615b2d575b5050848110615aee5750838310615a39575050615a0e6fffffffffffffffffffffffffffffffff9283926147ad565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b9190915460801c928315615aad5781615a51916147ad565b926000198101908082116147ba57615a74615a79926001600160a01b0396614879565b614f21565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b6001600160a01b0383837fd0c8d23a000000000000000000000000000000000000000000000000000000006000526000196004526024521660445260646000fd5b82856001600160a01b03927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611615ba157615b4892614e409160801c90614d47565b80841015615b9c5750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff00000000000000000000000000000000161786559238806159df565b615b53565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b508215615993565b91929015615c545750815115615bed575090565b3b15615bf65790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015615c675750805190602001fd5b6120ae906040519182917f08c379a00000000000000000000000000000000000000000000000000000000083526020600484015260248301906141b756fea164736f6c634300081a000a",
}

var SiloedLockReleaseTokenPoolABI = SiloedLockReleaseTokenPoolMetaData.ABI

var SiloedLockReleaseTokenPoolBin = SiloedLockReleaseTokenPoolMetaData.Bin

func DeploySiloedLockReleaseTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, localTokenDecimals uint8, advancedPoolHooks common.Address, rmnProxy common.Address, router common.Address, lockBox common.Address) (common.Address, *types.Transaction, *SiloedLockReleaseTokenPool, error) {
	parsed, err := SiloedLockReleaseTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SiloedLockReleaseTokenPoolBin), backend, token, localTokenDecimals, advancedPoolHooks, rmnProxy, router, lockBox)
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getCurrentRateLimiterState", remoteChainSelector, customBlockConfirmation)

	outstruct := new(GetCurrentRateLimiterState)
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutboundRateLimiterState = *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)
	outstruct.InboundRateLimiterState = *abi.ConvertType(out[1], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return *outstruct, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetCurrentRateLimiterState(remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	return _SiloedLockReleaseTokenPool.Contract.GetCurrentRateLimiterState(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector, customBlockConfirmation)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetCurrentRateLimiterState(remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	return _SiloedLockReleaseTokenPool.Contract.GetCurrentRateLimiterState(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector, customBlockConfirmation)
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
	outstruct.MinBlockConfirmations = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.RateLimitAdmin = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getFee", arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)

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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _SiloedLockReleaseTokenPool.Contract.GetFee(&_SiloedLockReleaseTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _SiloedLockReleaseTokenPool.Contract.GetFee(&_SiloedLockReleaseTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetRequiredCCVs(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getRequiredCCVs", localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetRequiredCCVs(&_SiloedLockReleaseTokenPool.CallOpts, localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetRequiredCCVs(&_SiloedLockReleaseTokenPool.CallOpts, localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getTokenTransferFeeConfig", arg0, destChainSelector, arg2, arg3)

	if err != nil {
		return *new(IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolV2TokenTransferFeeConfig)).(*IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetTokenTransferFeeConfig(&_SiloedLockReleaseTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetTokenTransferFeeConfig(&_SiloedLockReleaseTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3)
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "applyChainUpdates", remoteChainSelectorsToRemove, chainsToAdd)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ApplyChainUpdates(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ApplyChainUpdates(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "applyTokenTransferFeeConfigUpdates", tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_SiloedLockReleaseTokenPool.TransactOpts, tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_SiloedLockReleaseTokenPool.TransactOpts, tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.LockOrBurn0(&_SiloedLockReleaseTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.LockOrBurn0(&_SiloedLockReleaseTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationRequested, tokenArgs)
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn, blockConfirmationRequested)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ReleaseOrMint0(&_SiloedLockReleaseTokenPool.TransactOpts, releaseOrMintIn, blockConfirmationRequested)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ReleaseOrMint0(&_SiloedLockReleaseTokenPool.TransactOpts, releaseOrMintIn, blockConfirmationRequested)
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) SetDynamicConfig(opts *bind.TransactOpts, router common.Address, minBlockConfirmations uint16, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "setDynamicConfig", router, minBlockConfirmations, rateLimitAdmin)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) SetDynamicConfig(router common.Address, minBlockConfirmations uint16, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetDynamicConfig(&_SiloedLockReleaseTokenPool.TransactOpts, router, minBlockConfirmations, rateLimitAdmin)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) SetDynamicConfig(router common.Address, minBlockConfirmations uint16, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetDynamicConfig(&_SiloedLockReleaseTokenPool.TransactOpts, router, minBlockConfirmations, rateLimitAdmin)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) SetRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "setRateLimitConfig", rateLimitConfigArgs)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) SetRateLimitConfig(rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetRateLimitConfig(&_SiloedLockReleaseTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) SetRateLimitConfig(rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetRateLimitConfig(&_SiloedLockReleaseTokenPool.TransactOpts, rateLimitConfigArgs)
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "withdrawFeeTokens", feeTokens, recipient)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.WithdrawFeeTokens(&_SiloedLockReleaseTokenPool.TransactOpts, feeTokens, recipient)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.WithdrawFeeTokens(&_SiloedLockReleaseTokenPool.TransactOpts, feeTokens, recipient)
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

type SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator struct {
	Event *SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
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
		it.Event = new(SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
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

func (it *SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "CustomBlockConfirmationInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error) {
	event := new(SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator struct {
	Event *SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
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
		it.Event = new(SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
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

func (it *SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "CustomBlockConfirmationOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error) {
	event := new(SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
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
	Router                common.Address
	MinBlockConfirmations uint16
	RateLimitAdmin        common.Address
	Raw                   types.Log
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

type SiloedLockReleaseTokenPoolFeeTokenWithdrawnIterator struct {
	Event *SiloedLockReleaseTokenPoolFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolFeeTokenWithdrawn)
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
		it.Event = new(SiloedLockReleaseTokenPoolFeeTokenWithdrawn)
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

func (it *SiloedLockReleaseTokenPoolFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolFeeTokenWithdrawn struct {
	Recipient common.Address
	FeeToken  common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, recipient []common.Address, feeToken []common.Address) (*SiloedLockReleaseTokenPoolFeeTokenWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "FeeTokenWithdrawn", recipientRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolFeeTokenWithdrawnIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolFeeTokenWithdrawn, recipient []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "FeeTokenWithdrawn", recipientRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolFeeTokenWithdrawn)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseFeeTokenWithdrawn(log types.Log) (*SiloedLockReleaseTokenPoolFeeTokenWithdrawn, error) {
	event := new(SiloedLockReleaseTokenPoolFeeTokenWithdrawn)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

type SiloedLockReleaseTokenPoolRateLimitConfiguredIterator struct {
	Event *SiloedLockReleaseTokenPoolRateLimitConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolRateLimitConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolRateLimitConfigured)
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
		it.Event = new(SiloedLockReleaseTokenPoolRateLimitConfigured)
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

func (it *SiloedLockReleaseTokenPoolRateLimitConfiguredIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolRateLimitConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolRateLimitConfigured struct {
	RemoteChainSelector       uint64
	CustomBlockConfirmation   bool
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterRateLimitConfigured(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolRateLimitConfiguredIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "RateLimitConfigured", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolRateLimitConfiguredIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "RateLimitConfigured", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolRateLimitConfigured, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "RateLimitConfigured", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolRateLimitConfigured)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "RateLimitConfigured", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseRateLimitConfigured(log types.Log) (*SiloedLockReleaseTokenPoolRateLimitConfigured, error) {
	event := new(SiloedLockReleaseTokenPoolRateLimitConfigured)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "RateLimitConfigured", log); err != nil {
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

type GetCurrentRateLimiterState struct {
	OutboundRateLimiterState RateLimiterTokenBucket
	InboundRateLimiterState  RateLimiterTokenBucket
}
type GetDynamicConfig struct {
	Router                common.Address
	MinBlockConfirmations uint16
	RateLimitAdmin        common.Address
}
type GetFee struct {
	FeeUSDCents       *big.Int
	DestGasOverhead   uint32
	DestBytesOverhead uint32
	TokenFeeBps       uint16
	IsEnabled         bool
}

func (SiloedLockReleaseTokenPoolChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
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

func (SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce615")
}

func (SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac20979328")
}

func (SiloedLockReleaseTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0xba9213054b14c2e884f779120bb196f0735cef27140498a9d26117eeab77a117")
}

func (SiloedLockReleaseTokenPoolFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
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

func (SiloedLockReleaseTokenPoolRateLimitConfigured) Topic() common.Hash {
	return common.HexToHash("0x41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb")
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
	return common.HexToHash("0xfae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa1041")
}

func (SiloedLockReleaseTokenPoolUnsiloedRebalancerSet) Topic() common.Hash {
	return common.HexToHash("0x66b1c1bdec8b60a3442bb25b5b6cd6fff3d0eceb6f5390be8e2f82a8ad39b234")
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPool) Address() common.Address {
	return _SiloedLockReleaseTokenPool.address
}

type SiloedLockReleaseTokenPoolInterface interface {
	GetAvailableTokens(opts *bind.CallOpts, remoteChainSelector uint64) (*big.Int, error)

	GetChainRebalancer(opts *bind.CallOpts, remoteChainSelector uint64) (common.Address, error)

	GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

		error)

	GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

		error)

	GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

		error)

	GetRebalancer(opts *bind.CallOpts) (common.Address, error)

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRequiredCCVs(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetSupportedChains(opts *bind.CallOpts) ([]uint64, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	GetTokenDecimals(opts *bind.CallOpts) (uint8, error)

	GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error)

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

	ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error)

	ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error)

	ProvideLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	ProvideSiloedLiquidity(opts *bind.TransactOpts, remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, router common.Address, minBlockConfirmations uint16, rateLimitAdmin common.Address) (*types.Transaction, error)

	SetRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error)

	SetRebalancer(opts *bind.TransactOpts, newRebalancer common.Address) (*types.Transaction, error)

	SetSiloRebalancer(opts *bind.TransactOpts, remoteChainSelector uint64, newRebalancer common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateSiloDesignations(opts *bind.TransactOpts, removes []uint64, adds []SiloedLockReleaseTokenPoolSiloConfigUpdate) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error)

	WithdrawLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	WithdrawSiloedLiquidity(opts *bind.TransactOpts, remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error)

	FilterChainAdded(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*SiloedLockReleaseTokenPoolChainAdded, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*SiloedLockReleaseTokenPoolChainRemoved, error)

	FilterChainSiloed(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolChainSiloedIterator, error)

	WatchChainSiloed(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolChainSiloed) (event.Subscription, error)

	ParseChainSiloed(log types.Log) (*SiloedLockReleaseTokenPoolChainSiloed, error)

	FilterChainUnsiloed(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolChainUnsiloedIterator, error)

	WatchChainUnsiloed(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolChainUnsiloed) (event.Subscription, error)

	ParseChainUnsiloed(log types.Log) (*SiloedLockReleaseTokenPoolChainUnsiloed, error)

	FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*SiloedLockReleaseTokenPoolDynamicConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, recipient []common.Address, feeToken []common.Address) (*SiloedLockReleaseTokenPoolFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolFeeTokenWithdrawn, recipient []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*SiloedLockReleaseTokenPoolFeeTokenWithdrawn, error)

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

	FilterRateLimitConfigured(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolRateLimitConfiguredIterator, error)

	WatchRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolRateLimitConfigured, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRateLimitConfigured(log types.Log) (*SiloedLockReleaseTokenPoolRateLimitConfigured, error)

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
