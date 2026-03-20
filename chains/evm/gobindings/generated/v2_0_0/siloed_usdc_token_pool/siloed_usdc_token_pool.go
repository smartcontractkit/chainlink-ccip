// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package siloed_usdc_token_pool

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

type AuthorizedCallersAuthorizedCallerArgs struct {
	AddedCallers   []common.Address
	RemovedCallers []common.Address
}

type IPoolV2TokenTransferFeeConfig struct {
	DestGasOverhead                         uint32
	DestBytesOverhead                       uint32
	DefaultBlockConfirmationsFeeUSDCents    uint32
	CustomBlockConfirmationsFeeUSDCents     uint32
	DefaultBlockConfirmationsTransferFeeBps uint16
	CustomBlockConfirmationsTransferFeeBps  uint16
	IsEnabled                               bool
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

type SiloedLockReleaseTokenPoolLockBoxConfig struct {
	RemoteChainSelector uint64
	LockBox             common.Address
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
	CustomBlockConfirmations  bool
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
}

type TokenPoolTokenTransferFeeConfigArgs struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
}

var SiloedUSDCTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"advancedPoolHooks\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"struct AuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationsFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationsFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationsTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationsTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"disableTokenTransferFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnLockedUSDC\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelExistingCCTPMigrationProposal\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureLockBoxes\",\"inputs\":[{\"name\":\"lockBoxConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct SiloedLockReleaseTokenPool.LockBoxConfig[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"lockBox\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"excludeTokensFromBurn\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAdvancedPoolHooks\",\"inputs\":[],\"outputs\":[{\"name\":\"advancedPoolHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllLockBoxConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"lockBoxConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct SiloedLockReleaseTokenPool.LockBoxConfig[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"lockBox\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentProposedCCTPChainMigration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmations\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExcludedTokensByChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationsRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLockBox\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract ILockBox\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLockedUSDCToBurn\",\"inputs\":[],\"outputs\":[{\"name\":\"lockedUSDCToBurn\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinBlockConfirmations\",\"inputs\":[],\"outputs\":[{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockConfirmationsRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationsFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationsFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationsTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationsTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"lockOrBurnOutV1\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"blockConfirmationsRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposeCCTPMigration\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"blockConfirmationsRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCircleMigratorAddress\",\"inputs\":[{\"name\":\"migrator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLockedUSDCToBurn\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"lockedUSDCToBurn\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinBlockConfirmations\",\"inputs\":[{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.RateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmations\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"updateAdvancedPoolHooks\",\"inputs\":[{\"name\":\"newHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AdvancedPoolHooksUpdated\",\"inputs\":[{\"name\":\"oldHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"},{\"name\":\"newHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCTPMigrationCancelled\",\"inputs\":[{\"name\":\"existingProposalSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCTPMigrationExecuted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"USDCBurned\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCTPMigrationProposed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CircleMigratorAddressSet\",\"inputs\":[{\"name\":\"migratorAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationsInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationsOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"feeAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedUSDCToBurnSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"lockedUSDCToBurn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinBlockConfirmationsSet\",\"inputs\":[{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmations\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationsFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationsFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationsTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationsTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensExcludedFromBurn\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"burnableAmountAfterExclusion\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotOwnerOrFeeAdmin\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainAlreadyMigrated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustomBlockConfirmationsNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"ExistingMigrationProposal\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientLiquidity\",\"inputs\":[{\"name\":\"availableLiquidity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidChainSelector\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmations\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"LockBoxCannotBeShared\",\"inputs\":[{\"name\":\"chainSelectorA\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"chainSelectorB\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"lockBox\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"LockBoxNotConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoMigrationProposalPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCircle\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60e080604052346103505760a081616292803803809161001f82856103a0565b8339810103126103505780516001600160a01b03811680820361035057610048602084016103c3565b610054604085016103d1565b9061006d6080610066606088016103d1565b96016103d1565b926020956040519561007f88886103a0565b600087526000368137331561038f57600180546001600160a01b031916331790558215801561037e575b801561036d575b61035c5760805260c052853082036102c9575b505060a052600380546001600160a01b039283166001600160a01b0319918216179091556002805493909216921691909117905560405161010483826103a0565b60008152600036813760408051929083016001600160401b038111848210176102b3576040528252808383015260005b815181101561019b576001906001600160a01b0361015282856103e5565b51168561015e82610427565b61016b575b505001610134565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a13885610163565b50505160005b8151811015610212576001600160a01b036101bc82846103e5565b5116908115610201577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef84836101f3600195610525565b50604051908152a1016101a1565b6342bcdf7f60e11b60005260046000fd5b604051615d0c9081610586823960805181818161038d01528181611b5401528181612355015281816127b001528181612eea0152818161307d01528181613372015281816134ef0152818161397f015281816139cc0152614e2c015260a0518181816137dd01528181614ca401528181614cee01526150dd015260c05181818161100601528181611be2015281816123e201528181612f7801526134000152f35b634e487b7160e01b600052604160045260246000fd5b60405163313ce56760e01b815291829060049082905afa60009181610320575b506102f5575b856100c3565b60ff1660ff821681810361030957506102ef565b6332ad3e0760e11b60005260045260245260446000fd5b9091508681813d8311610355575b61033881836103a0565b8101031261035057610349906103c3565b90386102e9565b600080fd5b503d61032e565b630a64406560e11b60005260046000fd5b506001600160a01b038216156100b0565b506001600160a01b038616156100a9565b639b15e16f60e01b60005260046000fd5b601f909101601f19168101906001600160401b038211908210176102b357604052565b519060ff8216820361035057565b51906001600160a01b038216820361035057565b80518210156103f95760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b80548210156103f95760005260206000200190600090565b600081815260116020526040902054801561051e57600019810181811161050857601054600019810191908211610508578082036104b7575b50505060105480156104a1576000190161047b81601061040f565b8154906000199060031b1b19169055601055600052601160205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b6104f06104c86104d993601061040f565b90549060031b1c928392601061040f565b819391549060031b91821b91600019901b19161790565b90556000526011602052604060002055388080610460565b634e487b7160e01b600052601160045260246000fd5b5050600090565b8060005260116020526040600020541560001461057f57601054680100000000000000008110156102b3576105666104d9826001859401601055601061040f565b9055601054906000526011602052604060002055600190565b5060009056fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a714613a5157508063181f5a77146139f057806321df0da7146139ac578063240028e8146139555780632422ac45146138765780632451a6271461380157806324f65ee7146137c35780632bc3c64d1461378e5780632c063404146136f557806337a3210d146136ce5780633907753714613300578063489a68f214612e7b5780634ad01f0b14612de55780634c5ef0ed14612d9e57806350d1a35a14612c695780636160701b14612bea57806362ddd3c414612b63578063714bf90714612ae15780637437ff9f14612aa057806379ba5097146129f35780638926f54f146129ad57806389720a621461292d5780638a5e52bb1461270d5780638da5cb5b146126e657806391a2749a146125525780639a4575b9146122f0578063a42a7b8b146121a7578063acdb8f7b1461207f578063acfecf9114611f87578063ae39a25714611e58578063b1c71c6514611aba578063b794658014611a7d578063b8d5005e14611a58578063bfeffd3f146119c6578063c4bffe2b146118df578063c7230a60146116b2578063cd306a6c14611687578063d4d6de23146115e8578063d8aa3f40146114ae578063dc04fa1f1461102a578063dc0bd97114610fe6578063dcbd41bc14610dfc578063de814c5714610cfd578063e4d4c14314610cdf578063e8a1da171461064b578063efd07eec14610321578063f2fde38b1461026c5763f65a88861461023157600080fd5b3461026957602060031936011261026957604060209167ffffffffffffffff610258613bd6565b168152601383522054604051908152f35b80fd5b5034610269576020600319360112610269576001600160a01b0361028e613b94565b6102966149e2565b163381146102f957807fffffffffffffffffffffffff00000000000000000000000000000000000000008354161782556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b5034610269576020600319360112610269576004359067ffffffffffffffff821161026957366023830112156102695781600401359167ffffffffffffffff8311610647576024810190602436918560061b01011161064757916103836149e2565b61038b6149e2565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690825b81811061050e5783600d546103cc816142c0565b6103d5826142c0565b835b8381106104d05750835b8381106103ec578480f35b6001600160a01b036103fe828461463f565b5116600182018083116104a3575b85811061041d5750506001016103e1565b6001600160a01b0361043682869496999795989961463f565b5116841461044e57600101959491939592909261040c565b67ffffffffffffffff9150610472908261046a6064988661463f565b51169361463f565b5116907f30daa72d000000000000000000000000000000000000000000000000000000008452600452602452604452fd5b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b8067ffffffffffffffff6104e5600193615111565b91166104f1838761463f565b526001600160a01b03610504838661463f565b91169052016103d7565b6001600160a01b0361052c602061052684868a6153d2565b0161426f565b16801561061f576040517f75151b63000000000000000000000000000000000000000000000000000000008152846004820152602081602481855afa9081156106145786916105e6575b50156105ba57906105b360019267ffffffffffffffff61059f61059a85888c6153d2565b61421e565b1690818852600f60205260408820556156e3565b50016103b8565b602485857f961c9a4f000000000000000000000000000000000000000000000000000000008252600452fd5b610607915060203d811161060d575b6105ff8183613d52565b810190614979565b38610576565b503d6105f5565b6040513d88823e3d90fd5b6004857f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b5080fd5b50346102695760406003193601126102695760043567ffffffffffffffff81116106475761067d903690600401613ecf565b9060243567ffffffffffffffff8111610cdb57906106a084923690600401613ecf565b9390916106ab6149e2565b83905b828210610b215750505081927ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b81811015610b1d578060051b83013585811215610b1457830161012081360312610b14576040519461071286613d36565b813567ffffffffffffffff81168103610b18578652602082013567ffffffffffffffff81116106475782019436601f870112156106475785359561075587613e1f565b966107636040519889613d52565b80885260208089019160051b83010190368211610b145760208301905b828210610ae1575050505060208701958652604083013567ffffffffffffffff8111610add576107b39036908501613dbc565b91604088019283526107dd6107cb3660608701614834565b9460608a0195865260c0369101614834565b956080890196875283515115610ab55761080167ffffffffffffffff8a51166156aa565b15610a7e5767ffffffffffffffff89511682526008602052604082206108288651826151a2565b6108368851600283016151a2565b6004855191019080519067ffffffffffffffff8211610a51576108598354614682565b601f8111610a16575b50602090601f83116001146109b35761089292918691836109a8575b50506000198260011b9260031b1c19161790565b90555b815b885180518210156108cc57906108c66001926108bf8367ffffffffffffffff8f51169261463f565b5190614a20565b01610897565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c293919997509561099a67ffffffffffffffff600197969498511692519351915161096661093160405196879687526101006020880152610100870190613b53565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a10193929091936106e1565b015190508e8061087e565b8386528186209190601f198416875b8181106109fe57509084600195949392106109e5575b505050811b019055610895565b015160001960f88460031b161c191690558d80806109d8565b929360206001819287860151815501950193016109c2565b610a419084875260208720601f850160051c81019160208610610a47575b601f0160051c01906148dd565b8d610862565b9091508190610a34565b6024857f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60248267ffffffffffffffff8b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b6004827f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b8280fd5b813567ffffffffffffffff8111610b1057602091610b058392833691890101613dbc565b815201910190610780565b8680fd5b8480fd5b600080fd5b8380f35b9267ffffffffffffffff610b3e61059a8486889a9699979a614807565b1691610b4983615867565b15610caf578284526008602052610b65600560408620016153e2565b94845b8651811015610b9e576001908587526008602052610b9760056040892001610b90838b61463f565b51906158fb565b5001610b68565b5093969290945094909480875260086020526005604088208881558860018201558860028201558860038201558860048201610bda8154614682565b80610c6e575b5050500180549088815581610c50575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020836001948a52600482528985604082208281550155808a52600582528985604082208281550155604051908152a1019091949392946106ae565b885260208820908101905b81811015610bf057888155600101610c5b565b601f8111600114610c845750555b888a80610be0565b81835260208320610c9f91601f01861c8101906001016148dd565b8082528160208120915555610c7c565b602484847f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b8380fd5b50346102695780600319360112610269576020601454604051908152f35b503461026957604060031936011261026957610d17613bd6565b67ffffffffffffffff60243591610d2c6149e2565b16908115610dd4578167ffffffffffffffff60125460a01c1603610dac5760407fe1e6c22ce6b566f66cdb457ec2e7910ff1f9a9e5654ed75303476fa870468220918385526013602052818520610d848282546148d0565b9055610d9e60145485875260136020528387205490614233565b82519182526020820152a280f35b6004837fa94cb988000000000000000000000000000000000000000000000000000000008152fd5b6004837f656535ce000000000000000000000000000000000000000000000000000000008152fd5b50346102695760206003193601126102695760043567ffffffffffffffff811161064757610e2e903690600401613f00565b6001600160a01b03600a541633141580610fd1575b610fa557825b818110610e54578380f35b610e5f8183856147b9565b67ffffffffffffffff610e718261421e565b1690610e8a826000526007602052604060002054151590565b15610f7957907f41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb60e083610f39610f13602060019897018b610ecb826147c9565b15610f40578790526004602052610ef260408d20610eec3660408801614834565b906151a2565b868c526005602052610f0e60408d20610eec3660a08801614834565b6147c9565b916040519215158352610f2c602084016040830161488c565b60a060808401910161488c565ba201610e49565b60026040828a610f0e94526008602052610f62828220610eec36858c01614834565b8a815260086020522001610eec3660a08801614834565b602486837f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b6024837f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b506001600160a01b0360015416331415610e43565b503461026957806003193601126102695760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102695760406003193601126102695760043567ffffffffffffffff81116106475761105c903690600401613f00565b60243567ffffffffffffffff8111610cdb5761107c903690600401613ecf565b9190926110876149e2565b845b8281106110f357505050825b8181106110a0578380f35b8067ffffffffffffffff6110ba61059a6001948688614807565b16808652600b6020528560408120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee88680a201611095565b67ffffffffffffffff61110a61059a8386866147b9565b16611122816000526007602052604060002054151590565b15611483576111328285856147b9565b602081019060e0810190611145826147c9565b156114575760a0810161271061ffff61115d836147d6565b1610156114485760c082019161271061ffff611178856147d6565b1610156114105763ffffffff61118d866147e5565b16156113e457858c52600b60205260408c206111a8866147e5565b63ffffffff169080549060408401916111c0836147e5565b60201b67ffffffff00000000169360608601946111dc866147e5565b60401b6bffffffff00000000000000001696608001966111fb886147e5565b60601b6fffffffff000000000000000000000000169161121a8a6147d6565b60801b71ffff00000000000000000000000000000000169361123b8c6147d6565b60901b73ffff00000000000000000000000000000000000016957fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016177fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16177fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff16177fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff1617171781556112ee876147c9565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff0000000000000000000000000000000000000000161790556040519661133f906147f6565b63ffffffff168752611350906147f6565b63ffffffff166020870152611364906147f6565b63ffffffff166040860152611378906147f6565b63ffffffff16606085015261138c90613c76565b61ffff16608084015261139e90613c76565b61ffff1660a08301526113b090613c04565b151560c082015260e07ffae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa104191a2600101611089565b60248c877f12332265000000000000000000000000000000000000000000000000000000008252600452fd5b60248c61ffff61141f866147d6565b7f95f3517a00000000000000000000000000000000000000000000000000000000835216600452fd5b8a61ffff61141f6024936147d6565b60248a857f12332265000000000000000000000000000000000000000000000000000000008252600452fd5b7f1e670e4b000000000000000000000000000000000000000000000000000000008752600452602486fd5b5034610269576080600319360112610269576114c8613b94565b506114d1613bed565b6114d9613c65565b5060643567ffffffffffffffff8111610add579167ffffffffffffffff60409261150960e0953690600401613c85565b50508260c0855161151981613d1a565b82815282602082015282878201528260608201528260808201528260a08201520152168152600b602052206040519061155182613d1a565b5461ffff818163ffffffff82169485815263ffffffff60208201818560201c1681528160408401818760401c168152816060860193818960601c16855260ff60c060808901988a8c60801c168a528a60a082019c60901c168c52019b60a01c1615158b526040519b8c52511660208b0152511660408901525116606087015251166080850152511660a083015251151560c0820152f35b50346102695760206003193601126102695760043561ffff811690818103610add577f46c9c0585a955b2702c7ea47fec541db623565d20827a0edda42864e6b859a01916020916116376149e2565b7fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff75ffff00000000000000000000000000000000000000006002549260a01b16911617600255604051908152a180f35b5034610269578060031936011261026957602067ffffffffffffffff60125460a01c16604051908152f35b50346102695760406003193601126102695760043567ffffffffffffffff8111610647576116e4903690600401613ecf565b906116ed613bc0565b916001600160a01b0360015416331415806118ca575b61189e576001600160a01b03831690811561187657845b818110611725578580f35b6001600160a01b0361174061173b838588614807565b61426f565b1690604051917f70a08231000000000000000000000000000000000000000000000000000000008352306004840152602083602481845afa801561186b5785908990611831575b600194508061179a575b5050500161171a565b6040517fa9059cbb000000000000000000000000000000000000000000000000000000006020828101919091526001600160a01b038b166024830152604482018390527f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9290916118229061181c81606481015b03601f198101835282613d52565b86615c82565b604051908152a3388481611791565b5050909160203d8111611864575b6118498183613d52565b60208260009281010312610269575090846001939251611787565b503d61183f565b6040513d8a823e3d90fd5b6004857f8579befe000000000000000000000000000000000000000000000000000000008152fd5b6024847fcb1afbd700000000000000000000000000000000000000000000000000000000815233600452fd5b506001600160a01b03600c5416331415611703565b5034610269578060031936011261026957604051906006548083528260208101600684526020842092845b8181106119ad57505061191f92500383613d52565b61192982516142c0565b815b835181101561195d578067ffffffffffffffff61194a6001938761463f565b5116611956828561463f565b520161192b565b5090604051918291602083016020845282518091526020604085019301915b81811061198a575050500390f35b825167ffffffffffffffff1684528594506020938401939092019160010161197c565b845483526001948501948794506020909301920161190a565b5034610269576020600319360112610269576004356001600160a01b038116809103610647576119f46149e2565b7fffffffffffffffffffffffff00000000000000000000000000000000000000006003547fbaff46844acf36d6ee996f489a1a288709c4542bd33cd557770afd267d581209604080516001600160a01b0384168152856020820152a1161760035580f35b5034610269578060031936011261026957602061ffff60025460a01c16604051908152f35b503461026957602060031936011261026957611ab6611aa2611a9d613bd6565b614797565b604051918291602083526020830190613b53565b0390f35b5034610269576060600319360112610269576004359067ffffffffffffffff82116102695760a0600319833603011261026957611af5613c54565b9060443567ffffffffffffffff81116106475790611b1a611b3a923690600401613c85565b9290611b24614626565b50611b32858760040161513f565b933691613d75565b506084840192611b498461426f565b6001600160a01b03807f000000000000000000000000000000000000000000000000000000000000000016911603611e1b57602485019477ffffffffffffffff00000000000000000000000000000000611ba28761421e565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611e10578491611df1575b50611dc957606490611c2b611c268861421e565b615464565b01359161ffff611c3b8585614233565b92168015611d9c5761ffff60025460a01c16908115611d7457818110611d455750505093611c8e611d0a93611a9d93611c89611d3b98611c7a8961426f565b611c838761421e565b906157a0565b614233565b93611ca185611c9c8461421e565b614dda565b7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff611cdd611cd78561421e565b9361426f565b604080516001600160a01b039092168252336020830152810188905292169180606081015b0390a261421e565b90611d136150d6565b60405192611d2084613cfe565b83526020830152604051928392604084526040840190613ea5565b9060208301520390f35b604493507f1f5b9f77000000000000000000000000000000000000000000000000000000008352600452602452fd5b6004837f98d50fd7000000000000000000000000000000000000000000000000000000008152fd5b505093611c8e611d0a93611a9d93611c89611d3b98611dba8961426f565b611dc38761421e565b9061575a565b6004837f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b611e0a915060203d60201161060d576105ff8183613d52565b38611c12565b6040513d86823e3d90fd5b6024826001600160a01b03611e2f8761426f565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835216600452fd5b503461026957606060031936011261026957611e72613b94565b90611e7b613bc0565b604435926001600160a01b038416808503610cdb57611e986149e2565b6001600160a01b038216801561061f5794611f81917f3f1036e85d016a93254a0b1415844f79b85424959d90ae5ad51ce8f4533fe70195967fffffffffffffffffffffffff000000000000000000000000000000000000000060025416176002556001600160a01b0385167fffffffffffffffffffffffff0000000000000000000000000000000000000000600a541617600a557fffffffffffffffffffffffff0000000000000000000000000000000000000000600c541617600c55604051938493849160409194936001600160a01b03809281606087019816865216602085015216910152565b0390a180f35b50346102695767ffffffffffffffff611f9f36613dda565b929091611faa6149e2565b1691611fc3836000526007602052604060002054151590565b15610caf578284526008602052611ff260056040862001611fe5368486613d75565b60208151910120906158fb565b1561203757907f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d769161203160405192839260208452602084019161433d565b0390a280f35b8261207b836040519384937f74f23c7c000000000000000000000000000000000000000000000000000000008552600485015260406024850152604484019161433d565b0390fd5b5034610269578060031936011261026957600d5461209c81613e1f565b906120aa6040519283613d52565b808252601f196120b982613e1f565b01835b818110612184575050825b818110612133578284604051918291602083016020845282518091526020604085019301915b8181106120fb575050500390f35b8251805167ffffffffffffffff1685526020908101516001600160a01b031681860152869550604090940193909201916001016120ed565b806001600160a01b03612147600193615111565b604051929167ffffffffffffffff9061215f85613cfe565b168352166020820152612172828661463f565b5261217d818561463f565b50016120c7565b60209060405161219381613cfe565b8681528683820152828287010152016120bc565b50346102695760206003193601126102695767ffffffffffffffff6121ca613bd6565b16815260086020526121e1600560408320016153e2565b8051601f196122086121f283613e1f565b926122006040519485613d52565b808452613e1f565b01835b8181106122df575050825b825181101561225c578061222c6001928561463f565b5185526009602052612240604086206146d5565b61224a828561463f565b52612255818461463f565b5001612216565b81846040519182916020830160208452825180915260408401602060408360051b870101940192905b82821061229457505050500390f35b919360206122cf827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851613b53565b9601920192018594939192612285565b80606060208093860101520161220b565b50346102695760206003193601126102695760043567ffffffffffffffff81116106475760a060031982360301126106475761232a614626565b506020918060405161233c8582613d52565b526084820161234a8161426f565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691160361253e57602483019277ffffffffffffffff000000000000000000000000000000006123a38561421e565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015285816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611e10578491612521575b50611dc957606490612426611c268661421e565b013591806124f4575091817ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff85612478611a9d9661246f6124c49961426f565b611dc38861421e565b61248584611c9c8761421e565b611d0261249a6124948761421e565b9261426f565b604080516001600160a01b0390921682523360208301528101959095529116929081906060820190565b906124cd6150d6565b604051926124da84613cfe565b835281830152611ab6604051928284938452830190613ea5565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526011600452fd5b6125389150863d881161060d576105ff8183613d52565b38612412565b906001600160a01b03611e2f60249361426f565b50346102695760206003193601126102695760043567ffffffffffffffff81116106475760406003198236030112610647576040519061259182613cfe565b806004013567ffffffffffffffff8111610cdb576125b59060043691840101613e37565b825260248101359067ffffffffffffffff8211610cdb5760046125db9236920101613e37565b602082019081526125ea6149e2565b5191805b835181101561265457806001600160a01b0361260c6001938761463f565b511661261781615bee565b612623575b50016125ee565b60207fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a13861261c565b509051815b81518110156126e2576001600160a01b03612674828461463f565b511680156126ba57907feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef6020836126ac600195615671565b50604051908152a101612659565b6004847f8579befe000000000000000000000000000000000000000000000000000000008152fd5b8280f35b503461026957806003193601126102695760206001600160a01b0360015416604051908152f35b50346102695780600319360112610269576012546001600160a01b03811633036129055767ffffffffffffffff8160a01c16908115610dac57826001600160a01b0361275884614045565b7fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff6127926014548786526013602052604086205490614233565b9416601255826014556127a485615632565b50166001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690803b15610add576040517f74fd18ac0000000000000000000000000000000000000000000000000000000081526001600160a01b038316600482015267ffffffffffffffff86166024820152604481018590523060648201529083908290608490829084905af19081156128fa5783916128e5575b5050803b15610647578180916024604051809481937f42966c680000000000000000000000000000000000000000000000000000000083528860048401525af180156128da576128c5575b507fdea60ddd4c7ebdab804f5694c70350cca7893ece3efeecb142312eacac5c73e46040848482519182526020820152a180f35b816128cf91613d52565b610add578238612891565b6040513d84823e3d90fd5b816128ef91613d52565b610647578138612846565b6040513d85823e3d90fd5b6004827f5fff6eee000000000000000000000000000000000000000000000000000000008152fd5b50346102695760c060031936011261026957612947613b94565b61294f613bed565b60643561ffff81168103610cdb5760843567ffffffffffffffff8111610b145761297d903690600401613c85565b93909260a43595600287101561026957611ab66129a188888888604435888a61435e565b60405191829182613c11565b50346102695760206003193601126102695760206129e967ffffffffffffffff6129d5613bd6565b166000526007602052604060002054151590565b6040519015158152f35b503461026957806003193601126102695780546001600160a01b0381163303612a78577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551682556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b5034610269578060031936011261026957600254600a54600c54604080516001600160a01b0394851681529284166020840152921691810191909152606090f35b5034610269576020600319360112610269577f084e6f0e9791c2e56153bd49e6ec6dd63ba9a72c258d71558d74c63fc75b716860206001600160a01b03612b26613b94565b612b2e6149e2565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000006012541617601255604051908152a180f35b503461026957612b7236613dda565b612b7e939291936149e2565b67ffffffffffffffff8216612ba0816000526007602052604060002054151590565b15612bbf5750612bbc9293612bb6913691613d75565b90614a20565b80f35b7f1e670e4b000000000000000000000000000000000000000000000000000000008452600452602483fd5b503461026957604060031936011261026957612c04613bd6565b67ffffffffffffffff60243591612c196149e2565b16908115610dd4578167ffffffffffffffff60125460a01c1603610dac576020817f9cd13f8e5084fb8da7065c7a18e763c679accd381b4e9beff9c2b716493d008f92601455604051908152a280f35b503461026957602060031936011261026957612c83613bd6565b612c8b6149e2565b67ffffffffffffffff8116908115610dd4576012549067ffffffffffffffff8260a01c16612d7657612cca836000526016602052604060002054151590565b612d4a57916020917fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff7bffffffffffffffff00000000000000000000000000000000000000007f20331f191af84dbff48b162aa5a5985e7891ae646297b0a2ac80487f9109ef49958760145560a01b16911617601255604051908152a180f35b602484847f1c49a87b000000000000000000000000000000000000000000000000000000008252600452fd5b6004847f692bc131000000000000000000000000000000000000000000000000000000008152fd5b503461026957604060031936011261026957612db8613bd6565b906024359067ffffffffffffffff82116102695760206129e984612ddf3660048701613dbc565b90614283565b5034610269578060031936011261026957612dfe6149e2565b60125467ffffffffffffffff8160a01c16908115610dac577f375f1ad1194a2bec317c5efec05cc63ffa06ddd0c4b276619f6fd47298eda518917fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff602092166012558084526013825283604081205583601455604051908152a180f35b5034610269576040600319360112610269576004359067ffffffffffffffff8211610269578160040190610100600319843603011261026957612ebc613c54565b9080604051612eca81613cb3565b526064840135916084850190612edf8261426f565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116036132ec57602486019477ffffffffffffffff00000000000000000000000000000000612f388761421e565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156132e15785916132c2575b5061329a57612fb9611c268761421e565b612fc28661421e565b90612fdf60a4890192612ddf612fd88585614991565b3691613d75565b156132535750849161ffff1615905061323357612ffe61300d9261426f565b6130078661421e565b90615568565b67ffffffffffffffff61301f8461421e565b16938482526013602052604082205494851580159061321b575b80156131fe575b61319f575b50604401935090916130568461426f565b916130608261421e565b926001600160a01b0361307285614045565b166001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001694813b15610cdb576040517f74fd18ac0000000000000000000000000000000000000000000000000000000081526001600160a01b03878116600483015267ffffffffffffffff929092166024820152604481018890529216606483015282908290608490829084905af180156128da5761318a575b5050608067ffffffffffffffff6020956001600160a01b036131586124947ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09661421e565b60405196875233898801521660408601528560608601521692a26040519061317f82613cb3565b815260405190518152f35b613195828092613d52565b6102695780613113565b8584116131ce578260449394959650526013602052604083206131c3858254614233565b905584939291613045565b60448385887fa17e11d5000000000000000000000000000000000000000000000000000000008352600452602452fd5b50613216816000526016602052604060002054151590565b613040565b5067ffffffffffffffff60125460a01c168114613039565b61323f61324e9261426f565b6132488661421e565b906154fc565b61300d565b61325d9250614991565b61207b6040519283927f24eb47e500000000000000000000000000000000000000000000000000000000845260206004850152602484019161433d565b6004847f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b6132db915060203d60201161060d576105ff8183613d52565b38612fa8565b6040513d87823e3d90fd5b6024836001600160a01b03611e2f8561426f565b5034610269576020600319360112610269576004359067ffffffffffffffff821161026957816004016101006003198436030112610647578160405161334581613cb3565b528160405161335381613cb3565b52606483013590608484016133678161426f565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116036136ba57602485019177ffffffffffffffff000000000000000000000000000000006133c08461421e565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa90811561061457869161369b575b5061367357613441611c268461421e565b61344a8361421e565b9061346060a4880192612ddf612fd88585614991565b156132535750508261347461347d9261426f565b6132488461421e565b67ffffffffffffffff61348f8261421e565b16938484526013602052604084205494851580159061365b575b801561363e575b6135de575b5060449192939450016134c78161426f565b916134d18161421e565b856001600160a01b036134e383614045565b16946001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001695803b15610add576040517f74fd18ac0000000000000000000000000000000000000000000000000000000081526001600160a01b03808916600483015267ffffffffffffffff90951660248201526044810189905293909116606484015282908183816084810103925af1801561061457926001600160a01b0361315861249460809560209a67ffffffffffffffff967ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0996135ce575b505061421e565b816135d891613d52565b386135c7565b85841161360d57846044939495965052601360205260408520613602858254614233565b9055849392916134b5565b5050506044927fa17e11d5000000000000000000000000000000000000000000000000000000008352600452602452fd5b50613656816000526016602052604060002054151590565b6134b0565b5067ffffffffffffffff60125460a01c1681146134a9565b6004857f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b6136b4915060203d60201161060d576105ff8183613d52565b38613430565b836001600160a01b03611e2f60249361426f565b503461026957806003193601126102695760206001600160a01b0360035416604051908152f35b50346102695760c06003193601126102695761370f613b94565b50613718613bed565b613720613baa565b506084359161ffff831683036102695760a4359067ffffffffffffffff82116102695760a063ffffffff8061ffff61376788886137603660048b01613c85565b5050614099565b95926040979194975197885216602087015216604085015216606083015215156080820152f35b50346102695760206003193601126102695760206137b26137ad613bd6565b614045565b6001600160a01b0360405191168152f35b5034610269578060031936011261026957602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461026957806003193601126102695760405160108054808352908352909160208301917f1b6847dc741a1b0cd08d278845f9d819d87b734759afb55fe2de5cb82a9ae672915b81811061386057611ab6856129a181870382613d52565b8254845260209093019260019283019201613849565b503461026957604060031936011261026957613890613bd6565b602435918215158303610269576101406139536138ad8585613fc2565b61390360409392935180946fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b60a08301906fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565bf35b503461026957602060031936011261026957602090613972613b94565b90506001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b503461026957806003193601126102695760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461026957806003193601126102695750611ab6604051613a13604082613d52565b601981527f53696c6f656455534443546f6b656e506f6f6c20322e302e30000000000000006020820152604051918291602083526020830190613b53565b905034610647576020600319360112610647576004357fffffffff000000000000000000000000000000000000000000000000000000008116809103610add57602092507faff2afbf000000000000000000000000000000000000000000000000000000008114908115613b29575b8115613aff575b8115613ad5575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438613ace565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150613ac7565b7f331710310000000000000000000000000000000000000000000000000000000081149150613ac0565b919082519283825260005b848110613b7f575050601f19601f8460006020809697860101520116010190565b80602080928401015182828601015201613b5e565b600435906001600160a01b0382168203610b1857565b606435906001600160a01b0382168203610b1857565b602435906001600160a01b0382168203610b1857565b6004359067ffffffffffffffff82168203610b1857565b6024359067ffffffffffffffff82168203610b1857565b35908115158203610b1857565b602060408183019282815284518094520192019060005b818110613c355750505090565b82516001600160a01b0316845260209384019390920191600101613c28565b6024359061ffff82168203610b1857565b6044359061ffff82168203610b1857565b359061ffff82168203610b1857565b9181601f84011215610b185782359167ffffffffffffffff8311610b185760208381860195010111610b1857565b6020810190811067ffffffffffffffff821117613ccf57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff821117613ccf57604052565b60e0810190811067ffffffffffffffff821117613ccf57604052565b60a0810190811067ffffffffffffffff821117613ccf57604052565b90601f601f19910116810190811067ffffffffffffffff821117613ccf57604052565b92919267ffffffffffffffff8211613ccf5760405191613d9f601f8201601f191660200184613d52565b829481845281830111610b18578281602093846000960137010152565b9080601f83011215610b1857816020613dd793359101613d75565b90565b906040600319830112610b185760043567ffffffffffffffff81168103610b1857916024359067ffffffffffffffff8211610b1857613e1b91600401613c85565b9091565b67ffffffffffffffff8111613ccf5760051b60200190565b9080601f83011215610b1857813590613e4f82613e1f565b92613e5d6040519485613d52565b82845260208085019360051b820101918211610b1857602001915b818310613e855750505090565b82356001600160a01b0381168103610b1857815260209283019201613e78565b613dd7916020613ebe8351604084526040840190613b53565b920151906020818403910152613b53565b9181601f84011215610b185782359167ffffffffffffffff8311610b18576020808501948460051b010111610b1857565b9181601f84011215610b185782359167ffffffffffffffff8311610b18576020808501948460081b010111610b1857565b60405190613f3e82613d36565b60006080838281528260208201528260408201528260608201520152565b90604051613f6981613d36565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff91613fd4613f31565b50613fdd613f31565b5061401157166000526008602052604060002090613dd7614005600261400a61400586613f5c565b6148f4565b9401613f5c565b169081600052600460205261402c6140056040600020613f5c565b916000526005602052613dd76140056040600020613f5c565b67ffffffffffffffff166140588161542d565b91901561406c57506001600160a01b031690565b7f4fe6a5880000000000000000000000000000000000000000000000000000000060005260045260246000fd5b9061ffff8060025460a01c1691169283151592838094614216575b6141ec5767ffffffffffffffff16600052600b602052604060002091604051926140dd84613d1a565b5463ffffffff81168452602084019563ffffffff8260201c168752604085019263ffffffff8360401c168452606086019163ffffffff8460601c168352608087019761ffff8560801c16895260ff60a089019561ffff8160901c16875260a01c1615801560c08a01526141d157614172575050505063ffffffff808061ffff9351169451169551169351169193929190600190565b8193975080929450106141a157505063ffffffff808061ffff9351169451169551169351169193929190600190565b7f1f5b9f770000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b50505050505092505050600090600090600090600090600090565b7f98d50fd70000000000000000000000000000000000000000000000000000000060005260046000fd5b5082156140b4565b3567ffffffffffffffff81168103610b185790565b9190820391821161424057565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b356001600160a01b0381168103610b185790565b9067ffffffffffffffff613dd792166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b906142ca82613e1f565b6142d76040519182613d52565b828152601f196142e78294613e1f565b0190602036910137565b8181029291811591840414171561424057565b811561430e570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b601f8260209493601f19938186528686013760008582860101520116010190565b929495939091956001600160a01b036003541696871561460457809660028710156145d5576001600160a01b03976144849561ffff93896145a65767ffffffffffffffff8216600052600b602052604060002090604051916143bf83613d1a565b549163ffffffff8316815263ffffffff8360201c16602082015263ffffffff8360401c16604082015263ffffffff8360601c16606082015260c0878460801c169182608082015260ff60a08201958a8160901c16875260a01c1615159182910152614554575b50505067ffffffffffffffff905b6040519a8b997f89720a62000000000000000000000000000000000000000000000000000000008b521660048a0152166024880152604487015216606485015260c0608485015260c484019161433d565b938180600096879560a483015203915afa9182156145485780926144a757505090565b9091503d8082843e6144b98184613d52565b8201916020818403126106475780519067ffffffffffffffff8211610add57019180601f84011215610647578251916144f183613e1f565b936144ff6040519586613d52565b83855260208086019460051b82010192831161064757602001925b828410614528575050505090565b83516001600160a01b0381168103610add5781526020938401930161451a565b604051903d90823e3d90fd5b92935067ffffffffffffffff928587161561458e575061271061457d87614584945116836142f1565b0490614233565b915b903880614425565b6145a0925061457d61271091836142f1565b91614586565b67ffffffffffffffff9192506145cf906145c96145c436898b613d75565b614c30565b90614ceb565b91614433565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b505050505050505060405161461a602082613d52565b60008152600036813790565b6040519061463382613cfe565b60606020838281520152565b80518210156146535760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600182811c921680156146cb575b602083101461469c57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691614691565b90604051918260008254926146e984614682565b80845293600181169081156147575750600114614710575b5061470e92500383613d52565b565b90506000929192526020600020906000915b81831061473b57505090602061470e9282010138614701565b6020919350806001915483858901015201910190918492614722565b6020935061470e9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138614701565b67ffffffffffffffff166000526008602052613dd760046040600020016146d5565b91908110156146535760081b0190565b358015158103610b185790565b3561ffff81168103610b185790565b3563ffffffff81168103610b185790565b359063ffffffff82168203610b1857565b91908110156146535760051b0190565b35906fffffffffffffffffffffffffffffffff82168203610b1857565b9190826060910312610b18576040516060810181811067ffffffffffffffff821117613ccf57604052604061488781839561486e81613c04565b855261487c60208201614817565b602086015201614817565b910152565b6fffffffffffffffffffffffffffffffff6148ca604080936148ad81613c04565b15158652836148be60208301614817565b16602087015201614817565b16910152565b9190820180921161424057565b8181106148e8575050565b600081556001016148dd565b6148fc613f31565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff8083511691614959602085019361495361494663ffffffff87511642614233565b85608089015116906142f1565b906148d0565b8082101561497257505b16825263ffffffff4216905290565b9050614963565b90816020910312610b1857518015158103610b185790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610b18570180359067ffffffffffffffff8211610b1857602001918136038313610b1857565b6001600160a01b036001541633036149f657565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b90805115614c065767ffffffffffffffff81516020830120921691826000526008602052614a5581600560406000200161571c565b15614bc25760005260096020526040600020815167ffffffffffffffff8111613ccf57614a828254614682565b601f8111614b90575b506020601f8211600114614b065791614ae0827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9593614af695600091614afb575b506000198260011b9260031b1c19161790565b9055604051918291602083526020830190613b53565b0390a2565b905084015138614acd565b601f1982169083600052806000209160005b818110614b78575092614af69492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610614b5f575b5050811b019055611aa2565b85015160001960f88460031b161c191690553880614b53565b9192602060018192868a015181550194019201614b18565b614bbc90836000526020600020601f840160051c81019160208510610a4757601f0160051c01906148dd565b38614a8b565b509061207b6040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190613b53565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b80518015614ca057602003614c62578051602082810191830183900312610b1857519060ff8211614c62575060ff1690565b61207b906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190613b53565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff821161424057565b60ff16604d811161424057600a0a90565b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff811692828414614dd357828411614da95790614d3091614cc6565b91604d60ff8416118015614d8e575b614d5857505090614d52613dd792614cda565b906142f1565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b50614d9883614cda565b801561430e57600019048411614d3f565b614db291614cc6565b91604d60ff841611614d5857505090614dcd613dd792614cda565b90614304565b5050505090565b90614de482614045565b6040517f095ea7b300000000000000000000000000000000000000000000000000000000602082019081526001600160a01b03929092166024820181905260448201849052937f00000000000000000000000000000000000000000000000000000000000000009291614e5a816064810161180e565b60206000809483519082885af183513d826150b1575b50501561505a575b506001600160a01b03831693853b15610add5767ffffffffffffffff604051927fa36a7fee0000000000000000000000000000000000000000000000000000000084528660048501521660248301526044820152818160648183895af180156128da57829061504a575b50506040517fdd62ed3e000000000000000000000000000000000000000000000000000000008152306004820152846024820152602081604481875afa9081156128da578291615018575b50614f39575b50505050565b604051926020828186017f095ea7b300000000000000000000000000000000000000000000000000000000815287602488015281604488015260448752614f81606488613d52565b86519082875af1903d83519083614ff9575b505050614f3357614ff093614feb91604051917f095ea7b30000000000000000000000000000000000000000000000000000000060208401526024830152604482015260448152614fe5606482613d52565b82615c82565b615c82565b38808080614f33565b9192509061500e57503b15155b388080614f93565b6001915014615006565b90506020813d602011615042575b8161503360209383613d52565b81010312610b18575138614f2d565b3d9150615026565b61505391613d52565b3881614ee2565b6150ab906150a56040517f095ea7b30000000000000000000000000000000000000000000000000000000060208201528860248201528560448201526044815261181c606482613d52565b84615c82565b38614e78565b9091506150ce57506001600160a01b0384163b15155b3880614e70565b6001146150c7565b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016602082015260208152613dd7604082613d52565b61511c90600d6155d8565b90549060031b1c9081600052600f6020526001600160a01b036040600020541690565b906127109167ffffffffffffffff6151596020830161421e565b166000908152600b602052604090209161ffff161561518c57606061ffff615188935460901c169101356142f1565b0490565b606061ffff615188935460801c169101356142f1565b815191929115615324576fffffffffffffffffffffffffffffffff6040840151166fffffffffffffffffffffffffffffffff602085015116106152c15761470e91925b805182547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff0000000000000000000000000000000000000000161782556020810151825460409290920151608090811b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff9290921691821760018501557fffffffffffffffffffffffff0000000000000000000000000000000000000000909216174290911b73ffffffff0000000000000000000000000000000016179055565b606483615322604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff604084015116158015906153b3575b6153525761470e91926151e5565b606483615322604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff6020840151161515615344565b91908110156146535760061b0190565b906040519182815491828252602082019060005260206000209260005b81811061541457505061470e92500383613d52565b84548352600194850194879450602090930192016153ff565b80600052600f602052604060002054801560001461545c5750600052600e602052604060002054151590600090565b600192909150565b67ffffffffffffffff16615485816000526007602052604060002054151590565b156154cf5750336000526011602052604060002054156154a157565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fa9902c7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b9167ffffffffffffffff7f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9216928360005260086020526155458183600260406000200161599b565b604080516001600160a01b03909216825260208201929092529081908101614af6565b91909167ffffffffffffffff83169283600052600560205260ff60406000205460a01c16156155cd5750907f63335ad9e238acd0e9e6c1c20f529ffbea4cda73578c329a7aa7e9d61e5cdcc5918360005260056020526155458183604060002061599b565b9061470e93506154fc565b80548210156146535760005260206000200190600090565b80549068010000000000000000821015613ccf578161561791600161562e940181556155d8565b81939154906000199060031b92831b921b19161790565b9055565b8060005260166020526040600020541560001461566b576156548160156155f0565b601554906000526016602052604060002055600190565b50600090565b8060005260116020526040600020541560001461566b576156938160106155f0565b601054906000526011602052604060002055600190565b8060005260076020526040600020541560001461566b576156cc8160066155f0565b600654906000526007602052604060002055600190565b80600052600e6020526040600020541560001461566b5761570581600d6155f0565b600d5490600052600e602052604060002055600190565b6000828152600182016020526040902054615753578061573e836001936155f0565b80549260005201602052604060002055600190565b5050600090565b9167ffffffffffffffff7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789449216928360005260086020526155458183604060002061599b565b91909167ffffffffffffffff83169283600052600460205260ff60406000205460a01c16156158055750907f996b829383cc7e136842d4c4c175083bcf4e20807c7432105c1b794ba885e776918360005260046020526155458183604060002061599b565b9061470e935061575a565b8054801561583857600019019061582782826155d8565b60001982549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600081815260076020526040902054801561575357600019810181811161424057600654906000198201918211614240578181036158c1575b5050506158ad6006615810565b600052600760205260006040812055600190565b6158e36158d26156179360066155d8565b90549060031b1c92839260066155d8565b905560005260076020526040600020553880806158a0565b906001820191816000528260205260406000205490811515600014615992576000198201918083116142405781546000198101908111614240578381615949950361595b575b505050615810565b60005260205260006040812055600190565b61597b61596b61561793866155d8565b90549060031b1c928392866155d8565b905560005284602052604060002055388080615941565b50505050600090565b9182549060ff8260a01c16158015615be6575b614f33576fffffffffffffffffffffffffffffffff821691600185019081546159f363ffffffff6fffffffffffffffffffffffffffffffff83169360801c1642614233565b9081615b48575b5050848110615b095750838310615a54575050615a296fffffffffffffffffffffffffffffffff928392614233565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b9190915460801c928315615ac85781615a6c91614233565b9260001981019080821161424057615a8f615a94926001600160a01b03966148d0565b614304565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b6001600160a01b0383837fd0c8d23a000000000000000000000000000000000000000000000000000000006000526000196004526024521660445260646000fd5b82856001600160a01b03927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611615bbc57615b63926149539160801c906142f1565b80841015615bb75750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff00000000000000000000000000000000161786559238806159fa565b615b6e565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b5082156159ae565b60008181526011602052604090205480156157535760001981018181116142405760105490600019820191821161424057808203615c48575b505050615c346010615810565b600052601160205260006040812055600190565b615c6a615c596156179360106155d8565b90549060031b1c92839260106155d8565b90556000526011602052604060002055388080615c27565b906000602091828151910182855af115615cf3576000513d615cea57506001600160a01b0381163b155b615cb35750565b6001600160a01b03907f5274afe7000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60011415615cac565b6040513d6000823e3d90fdfea164736f6c634300081a000a",
}

var SiloedUSDCTokenPoolABI = SiloedUSDCTokenPoolMetaData.ABI

var SiloedUSDCTokenPoolBin = SiloedUSDCTokenPoolMetaData.Bin

func DeploySiloedUSDCTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, localTokenDecimals uint8, advancedPoolHooks common.Address, rmnProxy common.Address, router common.Address) (common.Address, *types.Transaction, *SiloedUSDCTokenPool, error) {
	parsed, err := SiloedUSDCTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SiloedUSDCTokenPoolBin), backend, token, localTokenDecimals, advancedPoolHooks, rmnProxy, router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SiloedUSDCTokenPool{address: address, abi: *parsed, SiloedUSDCTokenPoolCaller: SiloedUSDCTokenPoolCaller{contract: contract}, SiloedUSDCTokenPoolTransactor: SiloedUSDCTokenPoolTransactor{contract: contract}, SiloedUSDCTokenPoolFilterer: SiloedUSDCTokenPoolFilterer{contract: contract}}, nil
}

type SiloedUSDCTokenPool struct {
	address common.Address
	abi     abi.ABI
	SiloedUSDCTokenPoolCaller
	SiloedUSDCTokenPoolTransactor
	SiloedUSDCTokenPoolFilterer
}

type SiloedUSDCTokenPoolCaller struct {
	contract *bind.BoundContract
}

type SiloedUSDCTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type SiloedUSDCTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type SiloedUSDCTokenPoolSession struct {
	Contract     *SiloedUSDCTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type SiloedUSDCTokenPoolCallerSession struct {
	Contract *SiloedUSDCTokenPoolCaller
	CallOpts bind.CallOpts
}

type SiloedUSDCTokenPoolTransactorSession struct {
	Contract     *SiloedUSDCTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type SiloedUSDCTokenPoolRaw struct {
	Contract *SiloedUSDCTokenPool
}

type SiloedUSDCTokenPoolCallerRaw struct {
	Contract *SiloedUSDCTokenPoolCaller
}

type SiloedUSDCTokenPoolTransactorRaw struct {
	Contract *SiloedUSDCTokenPoolTransactor
}

func NewSiloedUSDCTokenPool(address common.Address, backend bind.ContractBackend) (*SiloedUSDCTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(SiloedUSDCTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindSiloedUSDCTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPool{address: address, abi: abi, SiloedUSDCTokenPoolCaller: SiloedUSDCTokenPoolCaller{contract: contract}, SiloedUSDCTokenPoolTransactor: SiloedUSDCTokenPoolTransactor{contract: contract}, SiloedUSDCTokenPoolFilterer: SiloedUSDCTokenPoolFilterer{contract: contract}}, nil
}

func NewSiloedUSDCTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*SiloedUSDCTokenPoolCaller, error) {
	contract, err := bindSiloedUSDCTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolCaller{contract: contract}, nil
}

func NewSiloedUSDCTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*SiloedUSDCTokenPoolTransactor, error) {
	contract, err := bindSiloedUSDCTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolTransactor{contract: contract}, nil
}

func NewSiloedUSDCTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*SiloedUSDCTokenPoolFilterer, error) {
	contract, err := bindSiloedUSDCTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolFilterer{contract: contract}, nil
}

func bindSiloedUSDCTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SiloedUSDCTokenPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SiloedUSDCTokenPool.Contract.SiloedUSDCTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SiloedUSDCTokenPoolTransactor.contract.Transfer(opts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SiloedUSDCTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SiloedUSDCTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.contract.Transfer(opts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetAdvancedPoolHooks(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getAdvancedPoolHooks")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetAdvancedPoolHooks() (common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetAdvancedPoolHooks(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetAdvancedPoolHooks() (common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetAdvancedPoolHooks(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getAllAuthorizedCallers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetAllAuthorizedCallers(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetAllAuthorizedCallers(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetAllLockBoxConfigs(opts *bind.CallOpts) ([]SiloedLockReleaseTokenPoolLockBoxConfig, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getAllLockBoxConfigs")

	if err != nil {
		return *new([]SiloedLockReleaseTokenPoolLockBoxConfig), err
	}

	out0 := *abi.ConvertType(out[0], new([]SiloedLockReleaseTokenPoolLockBoxConfig)).(*[]SiloedLockReleaseTokenPoolLockBoxConfig)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetAllLockBoxConfigs() ([]SiloedLockReleaseTokenPoolLockBoxConfig, error) {
	return _SiloedUSDCTokenPool.Contract.GetAllLockBoxConfigs(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetAllLockBoxConfigs() ([]SiloedLockReleaseTokenPoolLockBoxConfig, error) {
	return _SiloedUSDCTokenPool.Contract.GetAllLockBoxConfigs(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetCurrentProposedCCTPChainMigration(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getCurrentProposedCCTPChainMigration")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetCurrentProposedCCTPChainMigration() (uint64, error) {
	return _SiloedUSDCTokenPool.Contract.GetCurrentProposedCCTPChainMigration(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetCurrentProposedCCTPChainMigration() (uint64, error) {
	return _SiloedUSDCTokenPool.Contract.GetCurrentProposedCCTPChainMigration(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64, customBlockConfirmations bool) (GetCurrentRateLimiterState,

	error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getCurrentRateLimiterState", remoteChainSelector, customBlockConfirmations)

	outstruct := new(GetCurrentRateLimiterState)
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutboundRateLimiterState = *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)
	outstruct.InboundRateLimiterState = *abi.ConvertType(out[1], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return *outstruct, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetCurrentRateLimiterState(remoteChainSelector uint64, customBlockConfirmations bool) (GetCurrentRateLimiterState,

	error) {
	return _SiloedUSDCTokenPool.Contract.GetCurrentRateLimiterState(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector, customBlockConfirmations)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetCurrentRateLimiterState(remoteChainSelector uint64, customBlockConfirmations bool) (GetCurrentRateLimiterState,

	error) {
	return _SiloedUSDCTokenPool.Contract.GetCurrentRateLimiterState(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector, customBlockConfirmations)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

	error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getDynamicConfig")

	outstruct := new(GetDynamicConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Router = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RateLimitAdmin = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.FeeAdmin = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _SiloedUSDCTokenPool.Contract.GetDynamicConfig(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _SiloedUSDCTokenPool.Contract.GetDynamicConfig(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetExcludedTokensByChain(opts *bind.CallOpts, remoteChainSelector uint64) (*big.Int, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getExcludedTokensByChain", remoteChainSelector)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetExcludedTokensByChain(remoteChainSelector uint64) (*big.Int, error) {
	return _SiloedUSDCTokenPool.Contract.GetExcludedTokensByChain(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetExcludedTokensByChain(remoteChainSelector uint64) (*big.Int, error) {
	return _SiloedUSDCTokenPool.Contract.GetExcludedTokensByChain(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationsRequested uint16, arg5 []byte) (GetFee,

	error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getFee", arg0, destChainSelector, arg2, arg3, blockConfirmationsRequested, arg5)

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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationsRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _SiloedUSDCTokenPool.Contract.GetFee(&_SiloedUSDCTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationsRequested, arg5)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationsRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _SiloedUSDCTokenPool.Contract.GetFee(&_SiloedUSDCTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationsRequested, arg5)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetLockBox(opts *bind.CallOpts, remoteChainSelector uint64) (common.Address, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getLockBox", remoteChainSelector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetLockBox(remoteChainSelector uint64) (common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetLockBox(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetLockBox(remoteChainSelector uint64) (common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetLockBox(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetLockedUSDCToBurn(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getLockedUSDCToBurn")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetLockedUSDCToBurn() (*big.Int, error) {
	return _SiloedUSDCTokenPool.Contract.GetLockedUSDCToBurn(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetLockedUSDCToBurn() (*big.Int, error) {
	return _SiloedUSDCTokenPool.Contract.GetLockedUSDCToBurn(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetMinBlockConfirmations(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getMinBlockConfirmations")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetMinBlockConfirmations() (uint16, error) {
	return _SiloedUSDCTokenPool.Contract.GetMinBlockConfirmations(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetMinBlockConfirmations() (uint16, error) {
	return _SiloedUSDCTokenPool.Contract.GetMinBlockConfirmations(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getRemotePools", remoteChainSelector)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _SiloedUSDCTokenPool.Contract.GetRemotePools(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _SiloedUSDCTokenPool.Contract.GetRemotePools(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getRemoteToken", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _SiloedUSDCTokenPool.Contract.GetRemoteToken(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _SiloedUSDCTokenPool.Contract.GetRemoteToken(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetRequiredCCVs(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, sourceDenominatedAmount *big.Int, blockConfirmationsRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getRequiredCCVs", localToken, remoteChainSelector, sourceDenominatedAmount, blockConfirmationsRequested, extraData, direction)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, sourceDenominatedAmount *big.Int, blockConfirmationsRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetRequiredCCVs(&_SiloedUSDCTokenPool.CallOpts, localToken, remoteChainSelector, sourceDenominatedAmount, blockConfirmationsRequested, extraData, direction)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, sourceDenominatedAmount *big.Int, blockConfirmationsRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetRequiredCCVs(&_SiloedUSDCTokenPool.CallOpts, localToken, remoteChainSelector, sourceDenominatedAmount, blockConfirmationsRequested, extraData, direction)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetRmnProxy() (common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetRmnProxy(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetRmnProxy() (common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetRmnProxy(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetSupportedChains() ([]uint64, error) {
	return _SiloedUSDCTokenPool.Contract.GetSupportedChains(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetSupportedChains() ([]uint64, error) {
	return _SiloedUSDCTokenPool.Contract.GetSupportedChains(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetToken() (common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetToken(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetToken(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetTokenDecimals() (uint8, error) {
	return _SiloedUSDCTokenPool.Contract.GetTokenDecimals(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetTokenDecimals() (uint8, error) {
	return _SiloedUSDCTokenPool.Contract.GetTokenDecimals(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getTokenTransferFeeConfig", arg0, destChainSelector, arg2, arg3)

	if err != nil {
		return *new(IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolV2TokenTransferFeeConfig)).(*IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _SiloedUSDCTokenPool.Contract.GetTokenTransferFeeConfig(&_SiloedUSDCTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _SiloedUSDCTokenPool.Contract.GetTokenTransferFeeConfig(&_SiloedUSDCTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "isRemotePool", remoteChainSelector, remotePoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _SiloedUSDCTokenPool.Contract.IsRemotePool(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _SiloedUSDCTokenPool.Contract.IsRemotePool(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _SiloedUSDCTokenPool.Contract.IsSupportedChain(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _SiloedUSDCTokenPool.Contract.IsSupportedChain(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) IsSupportedToken(token common.Address) (bool, error) {
	return _SiloedUSDCTokenPool.Contract.IsSupportedToken(&_SiloedUSDCTokenPool.CallOpts, token)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _SiloedUSDCTokenPool.Contract.IsSupportedToken(&_SiloedUSDCTokenPool.CallOpts, token)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) Owner() (common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.Owner(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) Owner() (common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.Owner(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SiloedUSDCTokenPool.Contract.SupportsInterface(&_SiloedUSDCTokenPool.CallOpts, interfaceId)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SiloedUSDCTokenPool.Contract.SupportsInterface(&_SiloedUSDCTokenPool.CallOpts, interfaceId)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) TypeAndVersion() (string, error) {
	return _SiloedUSDCTokenPool.Contract.TypeAndVersion(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) TypeAndVersion() (string, error) {
	return _SiloedUSDCTokenPool.Contract.TypeAndVersion(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.AcceptOwnership(&_SiloedUSDCTokenPool.TransactOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.AcceptOwnership(&_SiloedUSDCTokenPool.TransactOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "addRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.AddRemotePool(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.AddRemotePool(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "applyAuthorizedCallerUpdates", authorizedCallerArgs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ApplyAuthorizedCallerUpdates(&_SiloedUSDCTokenPool.TransactOpts, authorizedCallerArgs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ApplyAuthorizedCallerUpdates(&_SiloedUSDCTokenPool.TransactOpts, authorizedCallerArgs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "applyChainUpdates", remoteChainSelectorsToRemove, chainsToAdd)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ApplyChainUpdates(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ApplyChainUpdates(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "applyTokenTransferFeeConfigUpdates", tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_SiloedUSDCTokenPool.TransactOpts, tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_SiloedUSDCTokenPool.TransactOpts, tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) BurnLockedUSDC(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "burnLockedUSDC")
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) BurnLockedUSDC() (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.BurnLockedUSDC(&_SiloedUSDCTokenPool.TransactOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) BurnLockedUSDC() (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.BurnLockedUSDC(&_SiloedUSDCTokenPool.TransactOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) CancelExistingCCTPMigrationProposal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "cancelExistingCCTPMigrationProposal")
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) CancelExistingCCTPMigrationProposal() (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.CancelExistingCCTPMigrationProposal(&_SiloedUSDCTokenPool.TransactOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) CancelExistingCCTPMigrationProposal() (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.CancelExistingCCTPMigrationProposal(&_SiloedUSDCTokenPool.TransactOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) ConfigureLockBoxes(opts *bind.TransactOpts, lockBoxConfigs []SiloedLockReleaseTokenPoolLockBoxConfig) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "configureLockBoxes", lockBoxConfigs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) ConfigureLockBoxes(lockBoxConfigs []SiloedLockReleaseTokenPoolLockBoxConfig) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ConfigureLockBoxes(&_SiloedUSDCTokenPool.TransactOpts, lockBoxConfigs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) ConfigureLockBoxes(lockBoxConfigs []SiloedLockReleaseTokenPoolLockBoxConfig) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ConfigureLockBoxes(&_SiloedUSDCTokenPool.TransactOpts, lockBoxConfigs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) ExcludeTokensFromBurn(opts *bind.TransactOpts, remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "excludeTokensFromBurn", remoteChainSelector, amount)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) ExcludeTokensFromBurn(remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ExcludeTokensFromBurn(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector, amount)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) ExcludeTokensFromBurn(remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ExcludeTokensFromBurn(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector, amount)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.LockOrBurn(&_SiloedUSDCTokenPool.TransactOpts, lockOrBurnIn)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.LockOrBurn(&_SiloedUSDCTokenPool.TransactOpts, lockOrBurnIn)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationsRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, blockConfirmationsRequested, tokenArgs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationsRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.LockOrBurn0(&_SiloedUSDCTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationsRequested, tokenArgs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationsRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.LockOrBurn0(&_SiloedUSDCTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationsRequested, tokenArgs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) ProposeCCTPMigration(opts *bind.TransactOpts, remoteChainSelector uint64) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "proposeCCTPMigration", remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) ProposeCCTPMigration(remoteChainSelector uint64) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ProposeCCTPMigration(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) ProposeCCTPMigration(remoteChainSelector uint64) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ProposeCCTPMigration(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ReleaseOrMint(&_SiloedUSDCTokenPool.TransactOpts, releaseOrMintIn)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ReleaseOrMint(&_SiloedUSDCTokenPool.TransactOpts, releaseOrMintIn)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationsRequested uint16) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn, blockConfirmationsRequested)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationsRequested uint16) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ReleaseOrMint0(&_SiloedUSDCTokenPool.TransactOpts, releaseOrMintIn, blockConfirmationsRequested)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationsRequested uint16) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ReleaseOrMint0(&_SiloedUSDCTokenPool.TransactOpts, releaseOrMintIn, blockConfirmationsRequested)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "removeRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.RemoveRemotePool(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.RemoveRemotePool(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) SetCircleMigratorAddress(opts *bind.TransactOpts, migrator common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "setCircleMigratorAddress", migrator)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) SetCircleMigratorAddress(migrator common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetCircleMigratorAddress(&_SiloedUSDCTokenPool.TransactOpts, migrator)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) SetCircleMigratorAddress(migrator common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetCircleMigratorAddress(&_SiloedUSDCTokenPool.TransactOpts, migrator)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) SetDynamicConfig(opts *bind.TransactOpts, router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "setDynamicConfig", router, rateLimitAdmin, feeAdmin)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) SetDynamicConfig(router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetDynamicConfig(&_SiloedUSDCTokenPool.TransactOpts, router, rateLimitAdmin, feeAdmin)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) SetDynamicConfig(router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetDynamicConfig(&_SiloedUSDCTokenPool.TransactOpts, router, rateLimitAdmin, feeAdmin)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) SetLockedUSDCToBurn(opts *bind.TransactOpts, remoteChainSelector uint64, lockedUSDCToBurn *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "setLockedUSDCToBurn", remoteChainSelector, lockedUSDCToBurn)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) SetLockedUSDCToBurn(remoteChainSelector uint64, lockedUSDCToBurn *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetLockedUSDCToBurn(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector, lockedUSDCToBurn)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) SetLockedUSDCToBurn(remoteChainSelector uint64, lockedUSDCToBurn *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetLockedUSDCToBurn(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector, lockedUSDCToBurn)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) SetMinBlockConfirmations(opts *bind.TransactOpts, minBlockConfirmations uint16) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "setMinBlockConfirmations", minBlockConfirmations)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) SetMinBlockConfirmations(minBlockConfirmations uint16) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetMinBlockConfirmations(&_SiloedUSDCTokenPool.TransactOpts, minBlockConfirmations)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) SetMinBlockConfirmations(minBlockConfirmations uint16) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetMinBlockConfirmations(&_SiloedUSDCTokenPool.TransactOpts, minBlockConfirmations)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) SetRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "setRateLimitConfig", rateLimitConfigArgs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) SetRateLimitConfig(rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetRateLimitConfig(&_SiloedUSDCTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) SetRateLimitConfig(rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetRateLimitConfig(&_SiloedUSDCTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.TransferOwnership(&_SiloedUSDCTokenPool.TransactOpts, to)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.TransferOwnership(&_SiloedUSDCTokenPool.TransactOpts, to)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) UpdateAdvancedPoolHooks(opts *bind.TransactOpts, newHook common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "updateAdvancedPoolHooks", newHook)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) UpdateAdvancedPoolHooks(newHook common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.UpdateAdvancedPoolHooks(&_SiloedUSDCTokenPool.TransactOpts, newHook)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) UpdateAdvancedPoolHooks(newHook common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.UpdateAdvancedPoolHooks(&_SiloedUSDCTokenPool.TransactOpts, newHook)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "withdrawFeeTokens", feeTokens, recipient)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.WithdrawFeeTokens(&_SiloedUSDCTokenPool.TransactOpts, feeTokens, recipient)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.WithdrawFeeTokens(&_SiloedUSDCTokenPool.TransactOpts, feeTokens, recipient)
}

type SiloedUSDCTokenPoolAdvancedPoolHooksUpdatedIterator struct {
	Event *SiloedUSDCTokenPoolAdvancedPoolHooksUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolAdvancedPoolHooksUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolAdvancedPoolHooksUpdated)
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
		it.Event = new(SiloedUSDCTokenPoolAdvancedPoolHooksUpdated)
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

func (it *SiloedUSDCTokenPoolAdvancedPoolHooksUpdatedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolAdvancedPoolHooksUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolAdvancedPoolHooksUpdated struct {
	OldHook common.Address
	NewHook common.Address
	Raw     types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterAdvancedPoolHooksUpdated(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolAdvancedPoolHooksUpdatedIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "AdvancedPoolHooksUpdated")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolAdvancedPoolHooksUpdatedIterator{contract: _SiloedUSDCTokenPool.contract, event: "AdvancedPoolHooksUpdated", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchAdvancedPoolHooksUpdated(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolAdvancedPoolHooksUpdated) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "AdvancedPoolHooksUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolAdvancedPoolHooksUpdated)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "AdvancedPoolHooksUpdated", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseAdvancedPoolHooksUpdated(log types.Log) (*SiloedUSDCTokenPoolAdvancedPoolHooksUpdated, error) {
	event := new(SiloedUSDCTokenPoolAdvancedPoolHooksUpdated)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "AdvancedPoolHooksUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolAuthorizedCallerAddedIterator struct {
	Event *SiloedUSDCTokenPoolAuthorizedCallerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolAuthorizedCallerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolAuthorizedCallerAdded)
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
		it.Event = new(SiloedUSDCTokenPoolAuthorizedCallerAdded)
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

func (it *SiloedUSDCTokenPoolAuthorizedCallerAddedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolAuthorizedCallerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolAuthorizedCallerAdded struct {
	Caller common.Address
	Raw    types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolAuthorizedCallerAddedIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolAuthorizedCallerAddedIterator{contract: _SiloedUSDCTokenPool.contract, event: "AuthorizedCallerAdded", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolAuthorizedCallerAdded) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolAuthorizedCallerAdded)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseAuthorizedCallerAdded(log types.Log) (*SiloedUSDCTokenPoolAuthorizedCallerAdded, error) {
	event := new(SiloedUSDCTokenPoolAuthorizedCallerAdded)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolAuthorizedCallerRemovedIterator struct {
	Event *SiloedUSDCTokenPoolAuthorizedCallerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolAuthorizedCallerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolAuthorizedCallerRemoved)
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
		it.Event = new(SiloedUSDCTokenPoolAuthorizedCallerRemoved)
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

func (it *SiloedUSDCTokenPoolAuthorizedCallerRemovedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolAuthorizedCallerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolAuthorizedCallerRemoved struct {
	Caller common.Address
	Raw    types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolAuthorizedCallerRemovedIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolAuthorizedCallerRemovedIterator{contract: _SiloedUSDCTokenPool.contract, event: "AuthorizedCallerRemoved", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolAuthorizedCallerRemoved) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolAuthorizedCallerRemoved)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseAuthorizedCallerRemoved(log types.Log) (*SiloedUSDCTokenPoolAuthorizedCallerRemoved, error) {
	event := new(SiloedUSDCTokenPoolAuthorizedCallerRemoved)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolCCTPMigrationCancelledIterator struct {
	Event *SiloedUSDCTokenPoolCCTPMigrationCancelled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolCCTPMigrationCancelledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolCCTPMigrationCancelled)
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
		it.Event = new(SiloedUSDCTokenPoolCCTPMigrationCancelled)
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

func (it *SiloedUSDCTokenPoolCCTPMigrationCancelledIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolCCTPMigrationCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolCCTPMigrationCancelled struct {
	ExistingProposalSelector uint64
	Raw                      types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterCCTPMigrationCancelled(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolCCTPMigrationCancelledIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "CCTPMigrationCancelled")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolCCTPMigrationCancelledIterator{contract: _SiloedUSDCTokenPool.contract, event: "CCTPMigrationCancelled", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchCCTPMigrationCancelled(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCCTPMigrationCancelled) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "CCTPMigrationCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolCCTPMigrationCancelled)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CCTPMigrationCancelled", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseCCTPMigrationCancelled(log types.Log) (*SiloedUSDCTokenPoolCCTPMigrationCancelled, error) {
	event := new(SiloedUSDCTokenPoolCCTPMigrationCancelled)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CCTPMigrationCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolCCTPMigrationExecutedIterator struct {
	Event *SiloedUSDCTokenPoolCCTPMigrationExecuted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolCCTPMigrationExecutedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolCCTPMigrationExecuted)
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
		it.Event = new(SiloedUSDCTokenPoolCCTPMigrationExecuted)
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

func (it *SiloedUSDCTokenPoolCCTPMigrationExecutedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolCCTPMigrationExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolCCTPMigrationExecuted struct {
	RemoteChainSelector uint64
	USDCBurned          *big.Int
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterCCTPMigrationExecuted(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolCCTPMigrationExecutedIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "CCTPMigrationExecuted")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolCCTPMigrationExecutedIterator{contract: _SiloedUSDCTokenPool.contract, event: "CCTPMigrationExecuted", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchCCTPMigrationExecuted(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCCTPMigrationExecuted) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "CCTPMigrationExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolCCTPMigrationExecuted)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CCTPMigrationExecuted", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseCCTPMigrationExecuted(log types.Log) (*SiloedUSDCTokenPoolCCTPMigrationExecuted, error) {
	event := new(SiloedUSDCTokenPoolCCTPMigrationExecuted)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CCTPMigrationExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolCCTPMigrationProposedIterator struct {
	Event *SiloedUSDCTokenPoolCCTPMigrationProposed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolCCTPMigrationProposedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolCCTPMigrationProposed)
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
		it.Event = new(SiloedUSDCTokenPoolCCTPMigrationProposed)
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

func (it *SiloedUSDCTokenPoolCCTPMigrationProposedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolCCTPMigrationProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolCCTPMigrationProposed struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterCCTPMigrationProposed(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolCCTPMigrationProposedIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "CCTPMigrationProposed")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolCCTPMigrationProposedIterator{contract: _SiloedUSDCTokenPool.contract, event: "CCTPMigrationProposed", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchCCTPMigrationProposed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCCTPMigrationProposed) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "CCTPMigrationProposed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolCCTPMigrationProposed)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CCTPMigrationProposed", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseCCTPMigrationProposed(log types.Log) (*SiloedUSDCTokenPoolCCTPMigrationProposed, error) {
	event := new(SiloedUSDCTokenPoolCCTPMigrationProposed)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CCTPMigrationProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolChainAddedIterator struct {
	Event *SiloedUSDCTokenPoolChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolChainAdded)
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
		it.Event = new(SiloedUSDCTokenPoolChainAdded)
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

func (it *SiloedUSDCTokenPoolChainAddedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolChainAdded struct {
	RemoteChainSelector       uint64
	RemoteToken               []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterChainAdded(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolChainAddedIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolChainAddedIterator{contract: _SiloedUSDCTokenPool.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolChainAdded) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolChainAdded)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseChainAdded(log types.Log) (*SiloedUSDCTokenPoolChainAdded, error) {
	event := new(SiloedUSDCTokenPoolChainAdded)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolChainRemovedIterator struct {
	Event *SiloedUSDCTokenPoolChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolChainRemoved)
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
		it.Event = new(SiloedUSDCTokenPoolChainRemoved)
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

func (it *SiloedUSDCTokenPoolChainRemovedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterChainRemoved(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolChainRemovedIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolChainRemovedIterator{contract: _SiloedUSDCTokenPool.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolChainRemoved) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolChainRemoved)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseChainRemoved(log types.Log) (*SiloedUSDCTokenPoolChainRemoved, error) {
	event := new(SiloedUSDCTokenPoolChainRemoved)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolCircleMigratorAddressSetIterator struct {
	Event *SiloedUSDCTokenPoolCircleMigratorAddressSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolCircleMigratorAddressSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolCircleMigratorAddressSet)
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
		it.Event = new(SiloedUSDCTokenPoolCircleMigratorAddressSet)
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

func (it *SiloedUSDCTokenPoolCircleMigratorAddressSetIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolCircleMigratorAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolCircleMigratorAddressSet struct {
	MigratorAddress common.Address
	Raw             types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterCircleMigratorAddressSet(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolCircleMigratorAddressSetIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "CircleMigratorAddressSet")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolCircleMigratorAddressSetIterator{contract: _SiloedUSDCTokenPool.contract, event: "CircleMigratorAddressSet", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchCircleMigratorAddressSet(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCircleMigratorAddressSet) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "CircleMigratorAddressSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolCircleMigratorAddressSet)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CircleMigratorAddressSet", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseCircleMigratorAddressSet(log types.Log) (*SiloedUSDCTokenPoolCircleMigratorAddressSet, error) {
	event := new(SiloedUSDCTokenPoolCircleMigratorAddressSet)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CircleMigratorAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolCustomBlockConfirmationsInboundRateLimitConsumedIterator struct {
	Event *SiloedUSDCTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolCustomBlockConfirmationsInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed)
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
		it.Event = new(SiloedUSDCTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed)
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

func (it *SiloedUSDCTokenPoolCustomBlockConfirmationsInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolCustomBlockConfirmationsInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterCustomBlockConfirmationsInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolCustomBlockConfirmationsInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationsInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolCustomBlockConfirmationsInboundRateLimitConsumedIterator{contract: _SiloedUSDCTokenPool.contract, event: "CustomBlockConfirmationsInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchCustomBlockConfirmationsInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationsInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationsInboundRateLimitConsumed", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseCustomBlockConfirmationsInboundRateLimitConsumed(log types.Log) (*SiloedUSDCTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed, error) {
	event := new(SiloedUSDCTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationsInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumedIterator struct {
	Event *SiloedUSDCTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed)
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
		it.Event = new(SiloedUSDCTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed)
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

func (it *SiloedUSDCTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterCustomBlockConfirmationsOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationsOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumedIterator{contract: _SiloedUSDCTokenPool.contract, event: "CustomBlockConfirmationsOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchCustomBlockConfirmationsOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationsOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationsOutboundRateLimitConsumed", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseCustomBlockConfirmationsOutboundRateLimitConsumed(log types.Log) (*SiloedUSDCTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed, error) {
	event := new(SiloedUSDCTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationsOutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolDynamicConfigSetIterator struct {
	Event *SiloedUSDCTokenPoolDynamicConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolDynamicConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolDynamicConfigSet)
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
		it.Event = new(SiloedUSDCTokenPoolDynamicConfigSet)
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

func (it *SiloedUSDCTokenPoolDynamicConfigSetIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolDynamicConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolDynamicConfigSet struct {
	Router         common.Address
	RateLimitAdmin common.Address
	FeeAdmin       common.Address
	Raw            types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterDynamicConfigSet(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolDynamicConfigSetIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolDynamicConfigSetIterator{contract: _SiloedUSDCTokenPool.contract, event: "DynamicConfigSet", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolDynamicConfigSet) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolDynamicConfigSet)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseDynamicConfigSet(log types.Log) (*SiloedUSDCTokenPoolDynamicConfigSet, error) {
	event := new(SiloedUSDCTokenPoolDynamicConfigSet)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolFeeTokenWithdrawnIterator struct {
	Event *SiloedUSDCTokenPoolFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolFeeTokenWithdrawn)
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
		it.Event = new(SiloedUSDCTokenPoolFeeTokenWithdrawn)
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

func (it *SiloedUSDCTokenPoolFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolFeeTokenWithdrawn struct {
	Receiver common.Address
	FeeToken common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*SiloedUSDCTokenPoolFeeTokenWithdrawnIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolFeeTokenWithdrawnIterator{contract: _SiloedUSDCTokenPool.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolFeeTokenWithdrawn)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseFeeTokenWithdrawn(log types.Log) (*SiloedUSDCTokenPoolFeeTokenWithdrawn, error) {
	event := new(SiloedUSDCTokenPoolFeeTokenWithdrawn)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolInboundRateLimitConsumedIterator struct {
	Event *SiloedUSDCTokenPoolInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolInboundRateLimitConsumed)
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
		it.Event = new(SiloedUSDCTokenPoolInboundRateLimitConsumed)
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

func (it *SiloedUSDCTokenPoolInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolInboundRateLimitConsumedIterator{contract: _SiloedUSDCTokenPool.contract, event: "InboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolInboundRateLimitConsumed)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseInboundRateLimitConsumed(log types.Log) (*SiloedUSDCTokenPoolInboundRateLimitConsumed, error) {
	event := new(SiloedUSDCTokenPoolInboundRateLimitConsumed)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolLockedOrBurnedIterator struct {
	Event *SiloedUSDCTokenPoolLockedOrBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolLockedOrBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolLockedOrBurned)
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
		it.Event = new(SiloedUSDCTokenPoolLockedOrBurned)
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

func (it *SiloedUSDCTokenPoolLockedOrBurnedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolLockedOrBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolLockedOrBurned struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolLockedOrBurnedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolLockedOrBurnedIterator{contract: _SiloedUSDCTokenPool.contract, event: "LockedOrBurned", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolLockedOrBurned)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseLockedOrBurned(log types.Log) (*SiloedUSDCTokenPoolLockedOrBurned, error) {
	event := new(SiloedUSDCTokenPoolLockedOrBurned)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolLockedUSDCToBurnSetIterator struct {
	Event *SiloedUSDCTokenPoolLockedUSDCToBurnSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolLockedUSDCToBurnSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolLockedUSDCToBurnSet)
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
		it.Event = new(SiloedUSDCTokenPoolLockedUSDCToBurnSet)
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

func (it *SiloedUSDCTokenPoolLockedUSDCToBurnSetIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolLockedUSDCToBurnSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolLockedUSDCToBurnSet struct {
	RemoteChainSelector uint64
	LockedUSDCToBurn    *big.Int
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterLockedUSDCToBurnSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolLockedUSDCToBurnSetIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "LockedUSDCToBurnSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolLockedUSDCToBurnSetIterator{contract: _SiloedUSDCTokenPool.contract, event: "LockedUSDCToBurnSet", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchLockedUSDCToBurnSet(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolLockedUSDCToBurnSet, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "LockedUSDCToBurnSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolLockedUSDCToBurnSet)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "LockedUSDCToBurnSet", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseLockedUSDCToBurnSet(log types.Log) (*SiloedUSDCTokenPoolLockedUSDCToBurnSet, error) {
	event := new(SiloedUSDCTokenPoolLockedUSDCToBurnSet)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "LockedUSDCToBurnSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolMinBlockConfirmationsSetIterator struct {
	Event *SiloedUSDCTokenPoolMinBlockConfirmationsSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolMinBlockConfirmationsSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolMinBlockConfirmationsSet)
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
		it.Event = new(SiloedUSDCTokenPoolMinBlockConfirmationsSet)
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

func (it *SiloedUSDCTokenPoolMinBlockConfirmationsSetIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolMinBlockConfirmationsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolMinBlockConfirmationsSet struct {
	MinBlockConfirmations uint16
	Raw                   types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterMinBlockConfirmationsSet(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolMinBlockConfirmationsSetIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "MinBlockConfirmationsSet")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolMinBlockConfirmationsSetIterator{contract: _SiloedUSDCTokenPool.contract, event: "MinBlockConfirmationsSet", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchMinBlockConfirmationsSet(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolMinBlockConfirmationsSet) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "MinBlockConfirmationsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolMinBlockConfirmationsSet)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "MinBlockConfirmationsSet", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseMinBlockConfirmationsSet(log types.Log) (*SiloedUSDCTokenPoolMinBlockConfirmationsSet, error) {
	event := new(SiloedUSDCTokenPoolMinBlockConfirmationsSet)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "MinBlockConfirmationsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolOutboundRateLimitConsumedIterator struct {
	Event *SiloedUSDCTokenPoolOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolOutboundRateLimitConsumed)
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
		it.Event = new(SiloedUSDCTokenPoolOutboundRateLimitConsumed)
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

func (it *SiloedUSDCTokenPoolOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolOutboundRateLimitConsumedIterator{contract: _SiloedUSDCTokenPool.contract, event: "OutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolOutboundRateLimitConsumed)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseOutboundRateLimitConsumed(log types.Log) (*SiloedUSDCTokenPoolOutboundRateLimitConsumed, error) {
	event := new(SiloedUSDCTokenPoolOutboundRateLimitConsumed)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolOwnershipTransferRequestedIterator struct {
	Event *SiloedUSDCTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolOwnershipTransferRequested)
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
		it.Event = new(SiloedUSDCTokenPoolOwnershipTransferRequested)
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

func (it *SiloedUSDCTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SiloedUSDCTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolOwnershipTransferRequestedIterator{contract: _SiloedUSDCTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolOwnershipTransferRequested)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*SiloedUSDCTokenPoolOwnershipTransferRequested, error) {
	event := new(SiloedUSDCTokenPoolOwnershipTransferRequested)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolOwnershipTransferredIterator struct {
	Event *SiloedUSDCTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolOwnershipTransferred)
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
		it.Event = new(SiloedUSDCTokenPoolOwnershipTransferred)
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

func (it *SiloedUSDCTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SiloedUSDCTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolOwnershipTransferredIterator{contract: _SiloedUSDCTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolOwnershipTransferred)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*SiloedUSDCTokenPoolOwnershipTransferred, error) {
	event := new(SiloedUSDCTokenPoolOwnershipTransferred)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolRateLimitConfiguredIterator struct {
	Event *SiloedUSDCTokenPoolRateLimitConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolRateLimitConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolRateLimitConfigured)
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
		it.Event = new(SiloedUSDCTokenPoolRateLimitConfigured)
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

func (it *SiloedUSDCTokenPoolRateLimitConfiguredIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolRateLimitConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolRateLimitConfigured struct {
	RemoteChainSelector       uint64
	CustomBlockConfirmations  bool
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterRateLimitConfigured(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolRateLimitConfiguredIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "RateLimitConfigured", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolRateLimitConfiguredIterator{contract: _SiloedUSDCTokenPool.contract, event: "RateLimitConfigured", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolRateLimitConfigured, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "RateLimitConfigured", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolRateLimitConfigured)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "RateLimitConfigured", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseRateLimitConfigured(log types.Log) (*SiloedUSDCTokenPoolRateLimitConfigured, error) {
	event := new(SiloedUSDCTokenPoolRateLimitConfigured)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "RateLimitConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolReleasedOrMintedIterator struct {
	Event *SiloedUSDCTokenPoolReleasedOrMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolReleasedOrMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolReleasedOrMinted)
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
		it.Event = new(SiloedUSDCTokenPoolReleasedOrMinted)
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

func (it *SiloedUSDCTokenPoolReleasedOrMintedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolReleasedOrMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolReleasedOrMinted struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Recipient           common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolReleasedOrMintedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolReleasedOrMintedIterator{contract: _SiloedUSDCTokenPool.contract, event: "ReleasedOrMinted", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolReleasedOrMinted)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseReleasedOrMinted(log types.Log) (*SiloedUSDCTokenPoolReleasedOrMinted, error) {
	event := new(SiloedUSDCTokenPoolReleasedOrMinted)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolRemotePoolAddedIterator struct {
	Event *SiloedUSDCTokenPoolRemotePoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolRemotePoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolRemotePoolAdded)
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
		it.Event = new(SiloedUSDCTokenPoolRemotePoolAdded)
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

func (it *SiloedUSDCTokenPoolRemotePoolAddedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolRemotePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolRemotePoolAdded struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolRemotePoolAddedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolRemotePoolAddedIterator{contract: _SiloedUSDCTokenPool.contract, event: "RemotePoolAdded", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolRemotePoolAdded)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseRemotePoolAdded(log types.Log) (*SiloedUSDCTokenPoolRemotePoolAdded, error) {
	event := new(SiloedUSDCTokenPoolRemotePoolAdded)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolRemotePoolRemovedIterator struct {
	Event *SiloedUSDCTokenPoolRemotePoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolRemotePoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolRemotePoolRemoved)
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
		it.Event = new(SiloedUSDCTokenPoolRemotePoolRemoved)
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

func (it *SiloedUSDCTokenPoolRemotePoolRemovedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolRemotePoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolRemotePoolRemoved struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolRemotePoolRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolRemotePoolRemovedIterator{contract: _SiloedUSDCTokenPool.contract, event: "RemotePoolRemoved", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolRemotePoolRemoved)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseRemotePoolRemoved(log types.Log) (*SiloedUSDCTokenPoolRemotePoolRemoved, error) {
	event := new(SiloedUSDCTokenPoolRemotePoolRemoved)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolTokenTransferFeeConfigDeletedIterator struct {
	Event *SiloedUSDCTokenPoolTokenTransferFeeConfigDeleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolTokenTransferFeeConfigDeletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolTokenTransferFeeConfigDeleted)
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
		it.Event = new(SiloedUSDCTokenPoolTokenTransferFeeConfigDeleted)
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

func (it *SiloedUSDCTokenPoolTokenTransferFeeConfigDeletedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolTokenTransferFeeConfigDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolTokenTransferFeeConfigDeleted struct {
	DestChainSelector uint64
	Raw               types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*SiloedUSDCTokenPoolTokenTransferFeeConfigDeletedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolTokenTransferFeeConfigDeletedIterator{contract: _SiloedUSDCTokenPool.contract, event: "TokenTransferFeeConfigDeleted", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolTokenTransferFeeConfigDeleted)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseTokenTransferFeeConfigDeleted(log types.Log) (*SiloedUSDCTokenPoolTokenTransferFeeConfigDeleted, error) {
	event := new(SiloedUSDCTokenPoolTokenTransferFeeConfigDeleted)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolTokenTransferFeeConfigUpdatedIterator struct {
	Event *SiloedUSDCTokenPoolTokenTransferFeeConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolTokenTransferFeeConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolTokenTransferFeeConfigUpdated)
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
		it.Event = new(SiloedUSDCTokenPoolTokenTransferFeeConfigUpdated)
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

func (it *SiloedUSDCTokenPoolTokenTransferFeeConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolTokenTransferFeeConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolTokenTransferFeeConfigUpdated struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
	Raw                    types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*SiloedUSDCTokenPoolTokenTransferFeeConfigUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolTokenTransferFeeConfigUpdatedIterator{contract: _SiloedUSDCTokenPool.contract, event: "TokenTransferFeeConfigUpdated", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolTokenTransferFeeConfigUpdated)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseTokenTransferFeeConfigUpdated(log types.Log) (*SiloedUSDCTokenPoolTokenTransferFeeConfigUpdated, error) {
	event := new(SiloedUSDCTokenPoolTokenTransferFeeConfigUpdated)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolTokensExcludedFromBurnIterator struct {
	Event *SiloedUSDCTokenPoolTokensExcludedFromBurn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolTokensExcludedFromBurnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolTokensExcludedFromBurn)
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
		it.Event = new(SiloedUSDCTokenPoolTokensExcludedFromBurn)
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

func (it *SiloedUSDCTokenPoolTokensExcludedFromBurnIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolTokensExcludedFromBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolTokensExcludedFromBurn struct {
	RemoteChainSelector          uint64
	Amount                       *big.Int
	BurnableAmountAfterExclusion *big.Int
	Raw                          types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterTokensExcludedFromBurn(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolTokensExcludedFromBurnIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "TokensExcludedFromBurn", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolTokensExcludedFromBurnIterator{contract: _SiloedUSDCTokenPool.contract, event: "TokensExcludedFromBurn", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchTokensExcludedFromBurn(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolTokensExcludedFromBurn, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "TokensExcludedFromBurn", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolTokensExcludedFromBurn)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "TokensExcludedFromBurn", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseTokensExcludedFromBurn(log types.Log) (*SiloedUSDCTokenPoolTokensExcludedFromBurn, error) {
	event := new(SiloedUSDCTokenPoolTokensExcludedFromBurn)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "TokensExcludedFromBurn", log); err != nil {
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

func (SiloedUSDCTokenPoolAdvancedPoolHooksUpdated) Topic() common.Hash {
	return common.HexToHash("0xbaff46844acf36d6ee996f489a1a288709c4542bd33cd557770afd267d581209")
}

func (SiloedUSDCTokenPoolAuthorizedCallerAdded) Topic() common.Hash {
	return common.HexToHash("0xeb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef")
}

func (SiloedUSDCTokenPoolAuthorizedCallerRemoved) Topic() common.Hash {
	return common.HexToHash("0xc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda77580")
}

func (SiloedUSDCTokenPoolCCTPMigrationCancelled) Topic() common.Hash {
	return common.HexToHash("0x375f1ad1194a2bec317c5efec05cc63ffa06ddd0c4b276619f6fd47298eda518")
}

func (SiloedUSDCTokenPoolCCTPMigrationExecuted) Topic() common.Hash {
	return common.HexToHash("0xdea60ddd4c7ebdab804f5694c70350cca7893ece3efeecb142312eacac5c73e4")
}

func (SiloedUSDCTokenPoolCCTPMigrationProposed) Topic() common.Hash {
	return common.HexToHash("0x20331f191af84dbff48b162aa5a5985e7891ae646297b0a2ac80487f9109ef49")
}

func (SiloedUSDCTokenPoolChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (SiloedUSDCTokenPoolChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (SiloedUSDCTokenPoolCircleMigratorAddressSet) Topic() common.Hash {
	return common.HexToHash("0x084e6f0e9791c2e56153bd49e6ec6dd63ba9a72c258d71558d74c63fc75b7168")
}

func (SiloedUSDCTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x63335ad9e238acd0e9e6c1c20f529ffbea4cda73578c329a7aa7e9d61e5cdcc5")
}

func (SiloedUSDCTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x996b829383cc7e136842d4c4c175083bcf4e20807c7432105c1b794ba885e776")
}

func (SiloedUSDCTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x3f1036e85d016a93254a0b1415844f79b85424959d90ae5ad51ce8f4533fe701")
}

func (SiloedUSDCTokenPoolFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (SiloedUSDCTokenPoolInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (SiloedUSDCTokenPoolLockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
}

func (SiloedUSDCTokenPoolLockedUSDCToBurnSet) Topic() common.Hash {
	return common.HexToHash("0x9cd13f8e5084fb8da7065c7a18e763c679accd381b4e9beff9c2b716493d008f")
}

func (SiloedUSDCTokenPoolMinBlockConfirmationsSet) Topic() common.Hash {
	return common.HexToHash("0x46c9c0585a955b2702c7ea47fec541db623565d20827a0edda42864e6b859a01")
}

func (SiloedUSDCTokenPoolOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944")
}

func (SiloedUSDCTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (SiloedUSDCTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (SiloedUSDCTokenPoolRateLimitConfigured) Topic() common.Hash {
	return common.HexToHash("0x41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb")
}

func (SiloedUSDCTokenPoolReleasedOrMinted) Topic() common.Hash {
	return common.HexToHash("0xfc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0")
}

func (SiloedUSDCTokenPoolRemotePoolAdded) Topic() common.Hash {
	return common.HexToHash("0x7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea")
}

func (SiloedUSDCTokenPoolRemotePoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76")
}

func (SiloedUSDCTokenPoolTokenTransferFeeConfigDeleted) Topic() common.Hash {
	return common.HexToHash("0x5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8")
}

func (SiloedUSDCTokenPoolTokenTransferFeeConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xfae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa1041")
}

func (SiloedUSDCTokenPoolTokensExcludedFromBurn) Topic() common.Hash {
	return common.HexToHash("0xe1e6c22ce6b566f66cdb457ec2e7910ff1f9a9e5654ed75303476fa870468220")
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPool) Address() common.Address {
	return _SiloedUSDCTokenPool.address
}

type SiloedUSDCTokenPoolInterface interface {
	GetAdvancedPoolHooks(opts *bind.CallOpts) (common.Address, error)

	GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error)

	GetAllLockBoxConfigs(opts *bind.CallOpts) ([]SiloedLockReleaseTokenPoolLockBoxConfig, error)

	GetCurrentProposedCCTPChainMigration(opts *bind.CallOpts) (uint64, error)

	GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64, customBlockConfirmations bool) (GetCurrentRateLimiterState,

		error)

	GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

		error)

	GetExcludedTokensByChain(opts *bind.CallOpts, remoteChainSelector uint64) (*big.Int, error)

	GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationsRequested uint16, arg5 []byte) (GetFee,

		error)

	GetLockBox(opts *bind.CallOpts, remoteChainSelector uint64) (common.Address, error)

	GetLockedUSDCToBurn(opts *bind.CallOpts) (*big.Int, error)

	GetMinBlockConfirmations(opts *bind.CallOpts) (uint16, error)

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRequiredCCVs(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, sourceDenominatedAmount *big.Int, blockConfirmationsRequested uint16, extraData []byte, direction uint8) ([]common.Address, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetSupportedChains(opts *bind.CallOpts) ([]uint64, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	GetTokenDecimals(opts *bind.CallOpts) (uint8, error)

	GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error)

	IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error)

	IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error)

	IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error)

	ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error)

	ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error)

	BurnLockedUSDC(opts *bind.TransactOpts) (*types.Transaction, error)

	CancelExistingCCTPMigrationProposal(opts *bind.TransactOpts) (*types.Transaction, error)

	ConfigureLockBoxes(opts *bind.TransactOpts, lockBoxConfigs []SiloedLockReleaseTokenPoolLockBoxConfig) (*types.Transaction, error)

	ExcludeTokensFromBurn(opts *bind.TransactOpts, remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationsRequested uint16, tokenArgs []byte) (*types.Transaction, error)

	ProposeCCTPMigration(opts *bind.TransactOpts, remoteChainSelector uint64) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationsRequested uint16) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetCircleMigratorAddress(opts *bind.TransactOpts, migrator common.Address) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error)

	SetLockedUSDCToBurn(opts *bind.TransactOpts, remoteChainSelector uint64, lockedUSDCToBurn *big.Int) (*types.Transaction, error)

	SetMinBlockConfirmations(opts *bind.TransactOpts, minBlockConfirmations uint16) (*types.Transaction, error)

	SetRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateAdvancedPoolHooks(opts *bind.TransactOpts, newHook common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error)

	FilterAdvancedPoolHooksUpdated(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolAdvancedPoolHooksUpdatedIterator, error)

	WatchAdvancedPoolHooksUpdated(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolAdvancedPoolHooksUpdated) (event.Subscription, error)

	ParseAdvancedPoolHooksUpdated(log types.Log) (*SiloedUSDCTokenPoolAdvancedPoolHooksUpdated, error)

	FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolAuthorizedCallerAddedIterator, error)

	WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolAuthorizedCallerAdded) (event.Subscription, error)

	ParseAuthorizedCallerAdded(log types.Log) (*SiloedUSDCTokenPoolAuthorizedCallerAdded, error)

	FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolAuthorizedCallerRemovedIterator, error)

	WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolAuthorizedCallerRemoved) (event.Subscription, error)

	ParseAuthorizedCallerRemoved(log types.Log) (*SiloedUSDCTokenPoolAuthorizedCallerRemoved, error)

	FilterCCTPMigrationCancelled(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolCCTPMigrationCancelledIterator, error)

	WatchCCTPMigrationCancelled(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCCTPMigrationCancelled) (event.Subscription, error)

	ParseCCTPMigrationCancelled(log types.Log) (*SiloedUSDCTokenPoolCCTPMigrationCancelled, error)

	FilterCCTPMigrationExecuted(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolCCTPMigrationExecutedIterator, error)

	WatchCCTPMigrationExecuted(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCCTPMigrationExecuted) (event.Subscription, error)

	ParseCCTPMigrationExecuted(log types.Log) (*SiloedUSDCTokenPoolCCTPMigrationExecuted, error)

	FilterCCTPMigrationProposed(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolCCTPMigrationProposedIterator, error)

	WatchCCTPMigrationProposed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCCTPMigrationProposed) (event.Subscription, error)

	ParseCCTPMigrationProposed(log types.Log) (*SiloedUSDCTokenPoolCCTPMigrationProposed, error)

	FilterChainAdded(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*SiloedUSDCTokenPoolChainAdded, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*SiloedUSDCTokenPoolChainRemoved, error)

	FilterCircleMigratorAddressSet(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolCircleMigratorAddressSetIterator, error)

	WatchCircleMigratorAddressSet(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCircleMigratorAddressSet) (event.Subscription, error)

	ParseCircleMigratorAddressSet(log types.Log) (*SiloedUSDCTokenPoolCircleMigratorAddressSet, error)

	FilterCustomBlockConfirmationsInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolCustomBlockConfirmationsInboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationsInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationsInboundRateLimitConsumed(log types.Log) (*SiloedUSDCTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationsOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationsOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationsOutboundRateLimitConsumed(log types.Log) (*SiloedUSDCTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*SiloedUSDCTokenPoolDynamicConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*SiloedUSDCTokenPoolFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*SiloedUSDCTokenPoolFeeTokenWithdrawn, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolInboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*SiloedUSDCTokenPoolInboundRateLimitConsumed, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolLockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*SiloedUSDCTokenPoolLockedOrBurned, error)

	FilterLockedUSDCToBurnSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolLockedUSDCToBurnSetIterator, error)

	WatchLockedUSDCToBurnSet(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolLockedUSDCToBurnSet, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedUSDCToBurnSet(log types.Log) (*SiloedUSDCTokenPoolLockedUSDCToBurnSet, error)

	FilterMinBlockConfirmationsSet(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolMinBlockConfirmationsSetIterator, error)

	WatchMinBlockConfirmationsSet(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolMinBlockConfirmationsSet) (event.Subscription, error)

	ParseMinBlockConfirmationsSet(log types.Log) (*SiloedUSDCTokenPoolMinBlockConfirmationsSet, error)

	FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolOutboundRateLimitConsumedIterator, error)

	WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseOutboundRateLimitConsumed(log types.Log) (*SiloedUSDCTokenPoolOutboundRateLimitConsumed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SiloedUSDCTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*SiloedUSDCTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SiloedUSDCTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*SiloedUSDCTokenPoolOwnershipTransferred, error)

	FilterRateLimitConfigured(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolRateLimitConfiguredIterator, error)

	WatchRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolRateLimitConfigured, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRateLimitConfigured(log types.Log) (*SiloedUSDCTokenPoolRateLimitConfigured, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*SiloedUSDCTokenPoolReleasedOrMinted, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolRemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*SiloedUSDCTokenPoolRemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolRemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*SiloedUSDCTokenPoolRemotePoolRemoved, error)

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*SiloedUSDCTokenPoolTokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*SiloedUSDCTokenPoolTokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*SiloedUSDCTokenPoolTokenTransferFeeConfigUpdatedIterator, error)

	WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigUpdated(log types.Log) (*SiloedUSDCTokenPoolTokenTransferFeeConfigUpdated, error)

	FilterTokensExcludedFromBurn(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolTokensExcludedFromBurnIterator, error)

	WatchTokensExcludedFromBurn(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolTokensExcludedFromBurn, remoteChainSelector []uint64) (event.Subscription, error)

	ParseTokensExcludedFromBurn(log types.Log) (*SiloedUSDCTokenPoolTokensExcludedFromBurn, error)

	Address() common.Address
}
