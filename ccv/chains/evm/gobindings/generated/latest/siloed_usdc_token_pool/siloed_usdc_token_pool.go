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

type SiloedLockReleaseTokenPoolLockBoxConfig struct {
	RemoteChainSelector uint64
	LockBox             common.Address
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

var SiloedUSDCTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"advancedPoolHooks\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockBox\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"struct AuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"disableTokenTransferFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnLockedUSDC\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelExistingCCTPMigrationProposal\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureLockBoxes\",\"inputs\":[{\"name\":\"lockBoxConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct SiloedLockReleaseTokenPool.LockBoxConfig[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"lockBox\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"excludeTokensFromBurn\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAdvancedPoolHooks\",\"inputs\":[],\"outputs\":[{\"name\":\"advancedPoolHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAvailableTokens\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"lockedTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChainRebalancer\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentProposedCCTPChainMigration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExcludedTokensByChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinBlockConfirmation\",\"inputs\":[],\"outputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRebalancer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUnsiloedLiquidity\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"out\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"out\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposeCCTPMigration\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"provideLiquidity\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"provideSiloedLiquidity\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCircleMigratorAddress\",\"inputs\":[{\"name\":\"migrator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinBlockConfirmation\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.RateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRebalancer\",\"inputs\":[{\"name\":\"newRebalancer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSiloRebalancer\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newRebalancer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"updateAdvancedPoolHooks\",\"inputs\":[{\"name\":\"newHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSiloDesignations\",\"inputs\":[{\"name\":\"removes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"adds\",\"type\":\"tuple[]\",\"internalType\":\"struct SiloedLockReleaseTokenPool.SiloConfigUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rebalancer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawLiquidity\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawSiloedLiquidity\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AdvancedPoolHooksUpdated\",\"inputs\":[{\"name\":\"oldHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"},{\"name\":\"newHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCTPMigrationCancelled\",\"inputs\":[{\"name\":\"existingProposalSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCTPMigrationExecuted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"USDCBurned\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCTPMigrationProposed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainSiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"rebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainUnsiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"amountUnsiloed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CircleMigratorAddressSet\",\"inputs\":[{\"name\":\"migratorAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinBlockConfirmationSet\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SiloRebalancerSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"oldRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensExcludedFromBurn\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"burnableAmountAfterExclusion\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnsiloedRebalancerSet\",\"inputs\":[{\"name\":\"oldRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainAlreadyMigrated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotSiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustomBlockConfirmationsNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"ExistingMigrationProposal\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientLiquidity\",\"inputs\":[{\"name\":\"availableLiquidity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidChainSelector\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmation\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"LiquidityAmountCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LockBoxNotConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoMigrationProposalPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCircle\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenLockingNotAllowedAfterMigration\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60e0806040523461049c5760c08161696a803803809161001f82856104db565b83398101031261049c5780516001600160a01b03811680820361049c57610048602084016104fe565b6100546040850161050c565b936100616060820161050c565b9461007a60a06100736080850161050c565b930161050c565b926020966040519661008c89896104db565b60008852600036813733156104ca57600180546001600160a01b03191633179055861580156104b9575b80156104a8575b6104305760805260c05260405163313ce56760e01b81528781600481895afa6000918161046c575b50610441575b5060a052600380546001600160a01b03199081166001600160a01b0393841617909155600280549091169282169290921790915516908115610430576040516375151b6360e01b8152600481018290528481602481865afa9081156103f357600091610413575b50156103ff578360009160446040518094819363095ea7b360e01b8352876004840152811960248401525af180156103f3576103c8575b506000808052600d845260409081902080546001600160a01b031916909217909155516101b683826104db565b60008152600036813760408051929083016001600160401b038111848210176103b2576040528252808383015260005b815181101561024d576001906001600160a01b036102048285610538565b5116856102108261057a565b61021d575b5050016101e6565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a13885610215565b50505160005b81518110156102c4576001600160a01b0361026e8284610538565b51169081156102b3577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef84836102a5600195610678565b50604051908152a101610253565b6342bcdf7f60e11b60005260046000fd5b60405161629190816106d982396080518181816104c30152818161058f015281816107ed01528181610fcd0152818161189e01528181611e4a01528181611f8e01528181611fed015281816121d60152818161254901528181612638015281816129ef01528181612eab015281816138620152818161399d015281816139f901528181613c4a01528181613d9801528181613ea401528181614175015281816144da0152818161452701526145ea015260a05181818161434d015281816154c10152615d39015260c05181818161128801528181611ec0015281816125be015281816138d80152613e0e0152f35b634e487b7160e01b600052604160045260246000fd5b6103e790843d86116103ec575b6103df81836104db565b810190610520565b610189565b503d6103d5565b6040513d6000823e3d90fd5b63961c9a4f60e01b60005260045260246000fd5b61042a9150853d87116103ec576103df81836104db565b38610152565b630a64406560e11b60005260046000fd5b60ff1660ff821681810361045557506100eb565b6332ad3e0760e11b60005260045260245260446000fd5b9091508881813d83116104a1575b61048481836104db565b8101031261049c57610495906104fe565b90386100e5565b600080fd5b503d61047a565b506001600160a01b038216156100bd565b506001600160a01b038516156100b6565b639b15e16f60e01b60005260046000fd5b601f909101601f19168101906001600160401b038211908210176103b257604052565b519060ff8216820361049c57565b51906001600160a01b038216820361049c57565b9081602091031261049c5751801515810361049c5790565b805182101561054c5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b805482101561054c5760005260206000200190600090565b600081815260106020526040902054801561067157600019810181811161065b57600f5460001981019190821161065b5780820361060a575b505050600f5480156105f457600019016105ce81600f610562565b8154906000199060031b1b19169055600f55600052601060205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61064361061b61062c93600f610562565b90549060031b1c928392600f610562565b819391549060031b91821b91600019901b19161790565b905560005260106020526040600020553880806105b3565b634e487b7160e01b600052601160045260246000fd5b5050600090565b806000526010602052604060002054156000146106d257600f54680100000000000000008110156103b2576106b961062c826001859401600f55600f610562565b9055600f54906000526010602052604060002055600190565b5060009056fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a71461476e575080630a861f2a146145ac578063181f5a771461454b57806321df0da714614507578063240028e8146144b35780632422ac45146143e65780632451a6271461437157806324f65ee7146143335780632c0634041461429b5780632d4a148f146140b557806331238ffc1461407557806337a3210d1461404e5780633907753714613d1f578063432a6ba314613cf8578063489a68f2146137ed5780634ad01f0b1461374d5780634c5ef0ed146137075780634e921c301461366857806350d1a35a1461352757806362ddd3c4146134ba5780636600f92c146133ff5780636cfd1553146133765780636d9d216c14612e43578063714bf90714612dd95780637437ff9f14612da557806379ba509714612d105780638632d5cc14612cdb5780638926f54f14612c9657806389720a6214612c175780638a5e52bb146129a45780638da5cb5b1461297d57806391a2749a146127ec5780639a4575b9146124dc578063a42a7b8b14612394578063acfecf91146122a3578063af0e58b914612282578063b1c71c6514611dab578063b794658014611d6e578063bfeffd3f14611cf4578063c4bffe2b14611be9578063c7230a6014611a76578063cd306a6c14611a4c578063ce3c75281461182f578063d8aa3f40146116f7578063dc04fa1f146112ac578063dc0bd97114611268578063dcbd41bc146110b2578063de814c5714610f39578063e8a1da171461093c578063eb521a4c14610769578063efd07eec14610553578063f1e733991461044c578063f2fde38b146103af578063f65a888614610377578063fa41d79c146103525763ff8e03f31461028957600080fd5b3461034f57604060031936011261034f576102a26148b1565b906102ab6148dd565b6102b36156fb565b6001600160a01b038316928315610340577f22a0dbb8195755fbfc99667a86ae684c568e9dfbb1eccf7f90084e616644797092936001600160a01b031960025416176002556001600160a01b0382166001600160a01b0319600a541617600a5561033a604051928392839092916001600160a01b0360209181604085019616845216910152565b0390a180f35b600483630a64406560e11b8152fd5b80fd5b503461034f578060031936011261034f57602061ffff60035460a01c16604051908152f35b503461034f57602060031936011261034f5760406020916001600160401b0361039e6148f3565b168152601283522054604051908152f35b503461034f57602060031936011261034f576001600160a01b036103d16148b1565b6103d96156fb565b1633811461042457806001600160a01b03198354161782556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b503461034f57602060031936011261034f576104666148f3565b6001600160401b038116610487816000526007602052604060002054151590565b15610541575061049e6001600160a01b039161569a565b1690604051916370a0823160e01b835260048301526020826024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa90811561053557906104fd575b602090604051908152f35b506020813d60201161052d575b8161051760209383614a4f565b8101031261052857602090516104f2565b600080fd5b3d915061050a565b604051903d90823e3d90fd5b631b3539ad60e31b8352600452602482fd5b503461034f57602060031936011261034f576004356001600160401b03811161076557610584903690600401614b48565b9061058d6156fb565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690835b8381106105c5578480f35b6001600160a01b036105e360206105dd848887614f89565b01614f13565b16908115610756576040517f75151b63000000000000000000000000000000000000000000000000000000008152846004820152602081602481865afa90811561074b57879161072d575b5015610701576001600160401b0361064f61064a838887614f89565b614edc565b168652600d602052604086206001600160a01b0383166001600160a01b0319825416179055604051917f095ea7b30000000000000000000000000000000000000000000000000000000083526004830152600019602483015260208260448189885af19182156106f6576001926106c8575b50016105ba565b6106e89060203d81116106ef575b6106e08183614a4f565b8101906153ef565b50386106c1565b503d6106d6565b6040513d88823e3d90fd5b602486857f961c9a4f000000000000000000000000000000000000000000000000000000008252600452fd5b610745915060203d81116106ef576106e08183614a4f565b3861062e565b6040513d89823e3d90fd5b600486630a64406560e11b8152fd5b5080fd5b503461034f57602060031936011261034f5760043581805260146020526040822054610910578015610901576001600160a01b036107a683614f99565b1633036108ee576040517f23b872dd0000000000000000000000000000000000000000000000000000000060208201523360248201523060448201526064810182905282907f00000000000000000000000000000000000000000000000000000000000000009061082e9061082881608481015b03601f198101835282614a4f565b82616213565b6001600160a01b0361083f8361569a565b16803b156108ea576040516351b53ff760e11b81526001600160a01b039290921660048301526001600160401b03831660248301526044820184905282908290606490829084905af180156108df576108ca575b50506040519082825260208201527f569a440e6842b5e5a7ac02286311855f5a0b81b9390909e552e82aaf02c9e9bf60403392a280f35b816108d491614a4f565b610765578138610893565b6040513d84823e3d90fd5b8280fd5b60248263472511eb60e11b815233600452fd5b60048263a90c0d1960e01b8152fd5b6024827f6469724600000000000000000000000000000000000000000000000000000000815280600452fd5b503461034f57604060031936011261034f576004356001600160401b0381116107655761096d903690600401614b18565b906024356001600160401b038111610f35579061098f84923690600401614b18565b93909161099a6156fb565b83905b828210610db95750505081927ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b81811015610db5578060051b83013585811215610db157830161012081360312610db15760405194610a0186614a34565b81356001600160401b038116810361052857865260208201356001600160401b0381116107655782019436601f8701121561076557853595610a4287614b78565b96610a506040519889614a4f565b80885260208089019160051b83010190368211610db15760208301905b828210610d7f57505050506020870195865260408301356001600160401b0381116108ea57610a9f9036908501614ab7565b9160408801928352610ac9610ab73660608701615326565b9460608a0195865260c0369101615326565b956080890196875283515115610d7057610aec6001600160401b038a5116615cbb565b15610d3a576001600160401b038951168252600860205260408220610b1286518261592e565b610b2088516002830161592e565b600485519101908051906001600160401b038211610d2657610b4283546151a7565b601f8111610ceb575b50602090601f8311600114610c8857610b7b9291869183610c7d575b50506000198260011b9260031b1c19161790565b90555b815b88518051821015610bb45790610bae600192610ba7836001600160401b038f511692615193565b5190615739565b01610b80565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2939199975095610c6f6001600160401b036001979694985116925193519151610c44610c1860405196879687526101006020880152610100870190614870565b9360408601906001600160801b0360408092805115158552826020820151166020860152015116910152565b60a08401906001600160801b0360408092805115158552826020820151166020860152015116910152565b0390a10193929091936109d0565b015190508e80610b67565b8386528186209190601f198416875b818110610cd35750908460019594939210610cba575b505050811b019055610b7e565b015160001960f88460031b161c191690558d8080610cad565b92936020600181928786015181550195019301610c97565b610d169084875260208720601f850160051c81019160208610610d1c575b601f0160051c01906153c5565b8d610b4b565b9091508190610d09565b602485634e487b7160e01b81526041600452fd5b6024826001600160401b038b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b600482630a64406560e11b8152fd5b81356001600160401b038111610dad57602091610da28392833691890101614ab7565b815201910190610a6d565b8680fd5b8480fd5b8380f35b926001600160401b03610dd561064a8486889a9699979a614f63565b1691610de083615dab565b15610f22578284526008602052610dfc60056040862001615b07565b94845b8651811015610e35576001908587526008602052610e2e60056040892001610e27838b615193565b5190615e3f565b5001610dff565b5093969290945094909480875260086020526005604088208881558860018201558860028201558860038201558860048201610e7181546151a7565b80610ee1575b5050500180549088815581610ec3575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a10190919493929461099d565b885260208820908101905b81811015610e8757888155600101610ece565b601f8111600114610ef75750555b888a80610e77565b81835260208320610f1291601f01861c8101906001016153c5565b8082528160208120915555610eef565b60248484631e670e4b60e01b8252600452fd5b8380fd5b503461034f57604060031936011261034f57610f536148f3565b602435610f5e6156fb565b6001600160401b0360115460a01c16916001600160401b0381168093036110a357610fa96001600160a01b0391848652601260205260408620610fa28582546153b8565b905561569a565b16604051906370a0823160e01b825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611098578491611046575b507fe1e6c22ce6b566f66cdb457ec2e7910ff1f9a9e5654ed75303476fa8704682209161103860409285875260126020528387205490614ef0565b82519182526020820152a280f35b90506020813d602011611090575b8161106160209383614a4f565b8101031261052857517fe1e6c22ce6b566f66cdb457ec2e7910ff1f9a9e5654ed75303476fa870468220610ffd565b3d9150611054565b6040513d86823e3d90fd5b600484631529973160e31b8152fd5b503461034f57602060031936011261034f576004356001600160401b038111610765576110e3903690600401614c27565b6001600160a01b03600a541633141580611253575b61124057825b818110611109578380f35b6111148183856152c4565b6001600160401b0361112582614edc565b169061113e826000526007602052604060002054151590565b1561122d57907f41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb60e0836111ed6111c7602060019897018b61117f826152d4565b156111f45787905260046020526111a660408d206111a03660408801615326565b9061592e565b868c5260056020526111c260408d206111a03660a08801615326565b6152d4565b9160405192151583526111e0602084016040830161537d565b60a060808401910161537d565ba2016110fe565b60026040828a6111c2945260086020526112168282206111a036858c01615326565b8a8152600860205220016111a03660a08801615326565b60248683631e670e4b60e01b8252600452fd5b60248363472511eb60e11b815233600452fd5b506001600160a01b03600154163314156110f8565b503461034f578060031936011261034f5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461034f57604060031936011261034f576004356001600160401b038111610765576112dd903690600401614c27565b6024356001600160401b038111610f35576112fc903690600401614b18565b9190926113076156fb565b845b82811061137257505050825b818110611320578380f35b806001600160401b0361133961064a6001948688614f63565b16808652600b6020528560408120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee88680a201611315565b61138061064a8285856152c4565b61138b8285856152c4565b90602082019060e083019061139f826152d4565b156116c35760a0840161271061ffff6113b7836152e1565b1610156116b45760c085019161271061ffff6113d2856152e1565b16101561167c5763ffffffff6113e7866152f0565b1615611648576001600160401b031694858c52600b60205260408c2061140c866152f0565b63ffffffff16908054906040840191611424836152f0565b60201b67ffffffff0000000016936060860194611440866152f0565b60401b6bffffffff000000000000000016966080019661145f886152f0565b60601b6fffffffff000000000000000000000000169161147e8a6152e1565b60801b71ffff00000000000000000000000000000000169361149f8c6152e1565b60901b73ffff00000000000000000000000000000000000016957fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016177fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16177fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff16177fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff161717178155611552876152d4565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016179055604051966115a390615301565b63ffffffff1687526115b490615301565b63ffffffff1660208701526115c890615301565b63ffffffff1660408601526115dc90615301565b63ffffffff1660608501526115f090614991565b61ffff16608084015261160290614991565b61ffff1660a08301526116149061491f565b151560c082015260e07ffae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa104191a2600101611309565b7f12332265000000000000000000000000000000000000000000000000000000008c526001600160401b031660045260248bfd5b60248c61ffff61168b866152e1565b7f95f3517a00000000000000000000000000000000000000000000000000000000835216600452fd5b8a61ffff61168b6024936152e1565b7f12332265000000000000000000000000000000000000000000000000000000008a526001600160401b0316600452602489fd5b503461034f57608060031936011261034f576117116148b1565b5061171a614909565b611722614980565b506064356001600160401b0381116108ea57916001600160401b0360409261175060e09536906004016149a0565b50508260c0855161176081614a19565b82815282602082015282878201528260608201528260808201528260a08201520152168152600b602052206040519061179882614a19565b5461ffff818163ffffffff82169485815263ffffffff60208201818560201c1681528160408401818760401c168152816060860193818960601c16855260ff60c060808901988a8c60801c168a528a60a082019c60901c168c52019b60a01c1615158b526040519b8c52511660208b0152511660408901525116606087015251166080850152511660a083015251151560c0820152f35b503461034f57604060031936011261034f576118496148f3565b6024356001600160401b038216808452600e60205260ff604085205460a01c16158015611a44575b611a3257508015611a23576001600160a01b0361188d83614f99565b163303611240576001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166001600160a01b036118cf8461569a565b16604051906370a0823160e01b82526004820152602081602481855afa908115611a185785916119e6575b508083116119cf575083906001600160a01b036119168561569a565b16803b156108ea57604051631d3f462b60e21b81526001600160a01b039290921660048301526001600160401b03851660248301526044820184905233606483015282908290608490829084905af180156108df576119ba575b5050604080516001600160401b039093168352602083019190915233917f58fca2457646a9f47422ab9eb9bff90cef88cd8b8725ab52b1d17baa392d784e91819081015b0390a280f35b816119c491614a4f565b6108ea578238611970565b848360449263a17e11d560e01b8352600452602452fd5b90506020813d602011611a10575b81611a0160209383614a4f565b810103126105285751386118fa565b3d91506119f4565b6040513d87823e3d90fd5b60048363a90c0d1960e01b8152fd5b6346f5f12b60e01b8452600452602483fd5b508015611871565b503461034f578060031936011261034f5760206001600160401b0360115460a01c16604051908152f35b503461034f57604060031936011261034f576004356001600160401b03811161076557611aa7903690600401614b18565b90611ab06148dd565b91611ab96156fb565b835b818110611ac6578480f35b6001600160a01b03611ae1611adc838587614f63565b614f13565b1690604051916370a0823160e01b8352306004840152602083602481845afa801561074b578790611bb1575b6001935080611b1f575b505001611abb565b6040517fa9059cbb0000000000000000000000000000000000000000000000000000000060208201526001600160a01b038816602482015260448101829052611b7590611b6f816064810161081a565b83616213565b6040519081527f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e60206001600160a01b03891692a33880611b17565b50909160203d8111611be2575b611bc88183614a4f565b6020826000928101031261034f5750906001929151611b0d565b503d611bbe565b503461034f578060031936011261034f57604051906006548083528260208101600684526020842092845b818110611cdb575050611c2992500383614a4f565b8151611c4d611c3782614b78565b91611c456040519384614a4f565b808352614b78565b91601f19602083019301368437805b8451811015611c8d57806001600160401b03611c7a60019388615193565b5116611c868286615193565b5201611c5c565b50925090604051928392602084019060208552518091526040840192915b818110611cb9575050500390f35b82516001600160401b0316845285945060209384019390920191600101611cab565b8454835260019485019487945060209093019201611c14565b503461034f57602060031936011261034f576004356001600160a01b03811680910361076557611d226156fb565b6001600160a01b03196003547fbaff46844acf36d6ee996f489a1a288709c4542bd33cd557770afd267d581209604080516001600160a01b0384168152856020820152a1161760035580f35b503461034f57602060031936011261034f57611da7611d93611d8e6148f3565b6152a3565b604051918291602083526020830190614870565b0390f35b503461034f57606060031936011261034f57600435906001600160401b03821161034f5760a0600319833603011261034f57611de561496f565b916044356001600160401b0381116108ea57611e08611e319136906004016149a0565b9190611e1261517a565b50611e1b61517a565b50611e298685600401615edf565b923691614a71565b5060848201611e3f81614f13565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116036122455750602482019177ffffffffffffffff00000000000000000000000000000000611e9984614edc565b60801b1660405190632cbc26bb60e01b825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611a18578591612226575b506121fe57606490611f09611f0485614edc565b615b52565b013561ffff611f188383614ef0565b95169485156121815761ffff60035460a01c168015612159578087106121295750839495507f61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac209793286001600160401b03611f72611fd696614edc565b1691828852600460205280611fb660408a206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391615f41565b604080516001600160a01b039290921682526020820192909252a2614ef0565b611fdf82614edc565b604080516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168082523360208301529181018490529093916001600160401b0316907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1090606090a28361205c611d8e83614edc565b93612065615d32565b60405195612072876149fe565b8652602086015261209c6001600160a01b0361209561209086614edc565b61569a565b1693614edc565b92803b156108ea576040516351b53ff760e11b81526001600160a01b0390921660048301526001600160401b039093166024820152604481018490529182908183816064810103925af1801561109857612114575b61210a8383604051928392604084526040840190614bfd565b9060208301520390f35b61211f848092614a4f565b6108ea57826120f1565b85604491887f7911d95b000000000000000000000000000000000000000000000000000000008352600452602452fd5b6004867f98d50fd7000000000000000000000000000000000000000000000000000000008152fd5b839495507fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789446001600160401b036121ba611fd696614edc565b1691828852600860205280611fb660408a206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391615f41565b6004847f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b61223f915060203d6020116106ef576106e08183614a4f565b38611ef0565b836001600160a01b03612259602493614f13565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835216600452fd5b503461034f578060031936011261034f576001600160a01b0361049e6155f9565b503461034f576001600160401b036122ba36614ad5565b9290916122c56156fb565b16916122de836000526007602052604060002054151590565b15610f2257828452600860205261230d60056040862001612300368486614a71565b6020815191012090615e3f565b1561234c57907f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76916119b4604051928392602084526020840191614fd7565b82612390836040519384937f74f23c7c0000000000000000000000000000000000000000000000000000000085526004850152604060248501526044840191614fd7565b0390fd5b503461034f57602060031936011261034f576001600160401b036123b66148f3565b16815260086020526123cd60056040832001615b07565b8051601f196123f46123de83614b78565b926123ec6040519485614a4f565b808452614b78565b01835b8181106124cb575050825b8251811015612448578061241860019285615193565b518552600960205261242c604086206151e1565b6124368285615193565b526124418184615193565b5001612402565b81846040519182916020830160208452825180915260408401602060408360051b870101940192905b82821061248057505050500390f35b919360206124bb827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851614870565b9601920192018594939192612471565b8060606020809386010152016123f7565b503461034f57602060031936011261034f576004356001600160401b0381116107655760a060031982360301126107655761251561517a565b5061251e61517a565b50602090826040516125308482614a4f565b526084810161253e81614f13565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116036122455750602481019077ffffffffffffffff0000000000000000000000000000000061259883614edc565b60801b1660405190632cbc26bb60e01b8252600482015283816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611a185785916127cf575b506121fe57606490612602611f0484614edc565b013592806127bb5750826001600160401b0361261d83614edc565b16806000526008845261266060406000206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016938491615f41565b604080516001600160a01b0384168152602081018790527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789449190a26126a482614edc565b604080516001600160a01b03841681523360208201529081018690526001600160401b0391909116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1090606090a26126ff611d8e83614edc565b93612708615d32565b60405195612715876149fe565b8652848601526127326001600160a01b0361209561209086614edc565b92803b15610528576040516351b53ff760e11b81526001600160a01b0390931660048401526001600160401b03909316602483015260448201529060009082908183816064810103925af180156127af5761279d575b50611da7604051928284938452830190614bfd565b60006127a891614a4f565b6000612788565b6040513d6000823e3d90fd5b80634e487b7160e01b602492526011600452fd5b6127e69150843d86116106ef576106e08183614a4f565b386125ee565b503461034f57602060031936011261034f576004356001600160401b0381116107655760406003198236030112610765576040519061282a826149fe565b80600401356001600160401b038111610f355761284d9060043691840101614b8f565b82526024810135906001600160401b038211610f355760046128729236920101614b8f565b602082019081526128816156fb565b5191805b83518110156128eb57806001600160a01b036128a360019387615193565b51166128ae8161617f565b6128ba575b5001612885565b60207fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a1386128b3565b509051815b8151811015612979576001600160a01b0361290b8284615193565b5116801561295157907feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef602083612943600195615c82565b50604051908152a1016128f0565b6004847f8579befe000000000000000000000000000000000000000000000000000000008152fd5b8280f35b503461034f578060031936011261034f5760206001600160a01b0360015416604051908152f35b503461034f578060031936011261034f576011546001600160a01b0381163303612bef5760a01c6001600160401b03168015612be0576129e38161569a565b826001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169216916040516370a0823160e01b8152836004820152602081602481855afa908115612b9d578391612ba8575b5084835260126020526040832054612a5391614ef0565b92803b156108ea57604051631d3f462b60e21b81526001600160a01b03831660048201526001600160401b0386166024820152604481018590523060648201529083908290608490829084905af1908115612b9d578391612b88575b50507fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff60115416601155803b15610765578180916024604051809481937f42966c680000000000000000000000000000000000000000000000000000000083528860048401525af180156108df57612b73575b505081612b4f7fdea60ddd4c7ebdab804f5694c70350cca7893ece3efeecb142312eacac5c73e493615c43565b50604080516001600160401b0390921682526020820192909252908190810161033a565b81612b7d91614a4f565b6108ea578238612b22565b81612b9291614a4f565b610765578138612aaf565b6040513d85823e3d90fd5b9250506020823d602011612bd8575b81612bc460209383614a4f565b8101031261052857612a5385925190612a3c565b3d9150612bb7565b600482631529973160e31b8152fd5b6004827f5fff6eee000000000000000000000000000000000000000000000000000000008152fd5b503461034f5760c060031936011261034f57612c316148b1565b612c39614909565b60643561ffff81168103610f35576084356001600160401b038111610db157612c669036906004016149a0565b93909260a43595600287101561034f57611da7612c8a88888888604435888a614ff8565b6040519182918261492c565b503461034f57602060031936011261034f576020612cd16001600160401b03612cbd6148f3565b166000526007602052604060002054151590565b6040519015158152f35b503461034f57602060031936011261034f576020612cff612cfa6148f3565b614f99565b6001600160a01b0360405191168152f35b503461034f578060031936011261034f5780546001600160a01b0381163303612d7d576001600160a01b03196001549133828416176001551682556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b503461034f578060031936011261034f57600254600a54604080516001600160a01b03938416815292909116602083015290f35b503461034f57602060031936011261034f577f084e6f0e9791c2e56153bd49e6ec6dd63ba9a72c258d71558d74c63fc75b716860206001600160a01b03612e1e6148b1565b612e266156fb565b16806001600160a01b03196011541617601155604051908152a180f35b503461034f57604060031936011261034f576004356001600160401b03811161076557612e74903690600401614b18565b906024356001600160401b038111610f355791612e9684933690600401614b48565b919092612ea16156fb565b612ea96155f9565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031692865b81811061314357878787825b818110612eed578380f35b6001600160401b03612f0361064a838587614f89565b16158015613114575b80156130f4575b6130cb576001600160a01b03612f2f60206105dd848688614f89565b16156130bc576001600160401b03612f4b61064a838587614f89565b168452600d6020526001600160a01b036040852054161561307a5780613016612f7c60206105dd6001958789614f89565b6001600160a01b0360405191612f91836149fe565b168152602081018481526001600160401b03612fb161064a86898b614f89565b168852600e602052604088209151825491517fffffffffffffffffffffff0000000000000000000000000000000000000000009092166001600160a01b03919091161790151560a01b74ff000000000000000000000000000000000000000016179055565b7f180c6940bd64ba8f75679203ca32f8be2f629477a3307b190656e4b14dd5ddeb604061304761064a848789614f89565b6001600160a01b0361305f60206105dd878a8c614f89565b6001600160401b03845193168352166020820152a101612ee2565b61064a613093916001600160401b039360249695614f89565b7f4fe6a58800000000000000000000000000000000000000000000000000000000835216600452fd5b600484630a64406560e11b8152fd5b61064a6130e4916001600160401b039360249695614f89565b631b3539ad60e31b835216600452fd5b5061310e6001600160401b03612cbd61064a848688614f89565b15612f13565b506001600160401b0361312b61064a838587614f89565b168452600e60205260ff604085205460a01c16612f0c565b6001600160401b0361315961064a838588614f63565b168852600e60205260ff604089205460a01c161561334d576001600160a01b0361318a61209061064a848689614f63565b169088604051926370a0823160e01b84528060048501526020846024818b5afa9384156108df578294613315575b5083613238575b50507f7b5efb3f8090c5cfd24e170b667d0e2b6fdc3db6540d75b86d5b6655ba00eb936001926001600160401b036131fb61064a85888b614f63565b168b52600e6020528a604081205561321761064a84878a614f63565b604080516001600160401b039290921682526020820192909252a101612ed6565b61324661064a84878a614f63565b90803b156108ea57604051631d3f462b60e21b81526001600160a01b038a1660048201526001600160401b039290921660248301526044820185905230606483015282908290608490829084905af180156108df57613300575b506001600160a01b038516803b15610765578180916064604051809481936351b53ff760e11b83528d60048401528160248401528960448401525af180156108df57156131bf57816132f191614a4f565b6132fc57888a6131bf565b8880fd5b8161330a91614a4f565b6132fc57888a6132a0565b9150925060203d8111613346575b61332d8183614a4f565b6020826000928101031261034f5750899051928b6131b8565b503d613323565b6001600160401b0361336661064a8a9360249588614f63565b6346f5f12b60e01b835216600452fd5b503461034f57602060031936011261034f577f66b1c1bdec8b60a3442bb25b5b6cd6fff3d0eceb6f5390be8e2f82a8ad39b2346001600160a01b036133b96148b1565b6133c16156fb565b61033a600c54918381166001600160a01b0319841617600c5560405193849316839092916001600160a01b0360209181604085019616845216910152565b503461034f57604060031936011261034f576134196148f3565b6001600160401b036134296148dd565b916134326156fb565b1690818352600e602052604083209081549160ff8360a01c16156134a75780546001600160a01b0319166001600160a01b039283169081179091556040805192909316825260208201527f01efd4cd7dd64263689551000d4359d6559c839f39b773b1df3fd19ff060cf5f91819081016119b4565b602485856346f5f12b60e01b8252600452fd5b503461034f576134c936614ad5565b6134d5939291936156fb565b6001600160401b0382166134f6816000526007602052604060002054151590565b156135155750613512929361350c913691614a71565b90615739565b80f35b631e670e4b60e01b8452600452602483fd5b503461034f57602060031936011261034f576135416148f3565b6135496156fb565b6011546001600160401b038160a01c16613640576001600160401b0382169161357f836000526014602052604060002054151590565b61361457821561360157916020917fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff7bffffffffffffffff00000000000000000000000000000000000000007f20331f191af84dbff48b162aa5a5985e7891ae646297b0a2ac80487f9109ef499560a01b16911617601155604051908152a180f35b602484631b3539ad60e31b815280600452fd5b602484847f1c49a87b000000000000000000000000000000000000000000000000000000008252600452fd5b6004837f692bc131000000000000000000000000000000000000000000000000000000008152fd5b503461034f57602060031936011261034f5760043561ffff8116908181036108ea577fa7f8dbba8cdb126ce4a0e7939ec58e0161b70d808b585dd651d68e59d27e11fb916020916136b76156fb565b7fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff75ffff00000000000000000000000000000000000000006003549260a01b16911617600355604051908152a180f35b503461034f57604060031936011261034f576137216148f3565b90602435906001600160401b03821161034f576020612cd1846137473660048701614ab7565b90614f27565b503461034f578060031936011261034f576137666156fb565b6011546001600160401b038160a01c169081156137de577f375f1ad1194a2bec317c5efec05cc63ffa06ddd0c4b276619f6fd47298eda518917fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff6020921660115580845260128252836040812055604051908152a180f35b600483631529973160e31b8152fd5b503461034f57604060031936011261034f57600435906001600160401b03821161034f578160040190610100600319843603011261034f5761382d61496f565b918160405161383b816149cd565b5261384960648501356154bf565b926084850161385781614f13565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116036122455750602485019177ffffffffffffffff000000000000000000000000000000006138b184614edc565b60801b1660405190632cbc26bb60e01b825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611a18578591613cd9575b506121fe57613919611f0484614edc565b61392283614edc565b9061393f60a488019261374761393885856155a9565b3691614a71565b15613c9257505061ffff1615613bf9576001600160401b0361396082614edc565b1680835260056020527f0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce61584806139c5604087206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391615f41565b604080516001600160a01b039290921682526020820192909252a25b6139ed61209082614edc565b916001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169316946040516370a0823160e01b8152866004820152602081602481885afa908115612b9d578391613bc7575b50808611613bb057506001600160401b03613a6084614edc565b1682526012602052604082205480613b78575b506044613a7f84614edc565b910195613a8b87614f13565b91813b15610f3557604051631d3f462b60e21b81526001600160a01b0387811660048301526001600160401b03929092166024820152604481018890529216606483015282908290608490829084905af180156108df57613b63575b505060806001600160401b036020956001600160a01b03613b31613b2b7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc096614edc565b92614f13565b60405196875233898801521660408601528560608601521692a260405190613b58826149cd565b815260405190518152f35b613b6e828092614a4f565b61034f5780613ae7565b808611613bb057506001600160401b03613b9184614edc565b168252601260205260408220613ba8868254614ef0565b905538613a73565b828660449263a17e11d560e01b8352600452602452fd5b90506020813d602011613bf1575b81613be260209383614a4f565b81010312610528575138613a46565b3d9150613bd5565b6001600160401b03613c0a82614edc565b1680835260086020527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c8480613c72600260408820016001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391615f41565b604080516001600160a01b039290921682526020820192909252a26139e1565b613c9c92506155a9565b6123906040519283927f24eb47e5000000000000000000000000000000000000000000000000000000008452602060048501526024840191614fd7565b613cf2915060203d6020116106ef576106e08183614a4f565b38613908565b503461034f578060031936011261034f5760206001600160a01b03600c5416604051908152f35b503461034f57602060031936011261034f576004356001600160401b038111610765578060040191610100600319833603011261034f5780604051613d63816149cd565b5280604051613d71816149cd565b52613d7f60648301356154bf565b9060848301613d8d81614f13565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691160361403a5750602483019077ffffffffffffffff00000000000000000000000000000000613de783614edc565b60801b1660405190632cbc26bb60e01b825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156108df57829161401b575b50613ff357613e4f611f0483614edc565b613e5882614edc565b613e6d60a4860191613747613938848a6155a9565b15613fe957508291906001600160401b03613e8783614edc565b168082526008602052613ecc600260408420016001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016958691615f41565b604080516001600160a01b0386168152602081018790527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a26001600160a01b03613f1b61209084614edc565b16946040516370a0823160e01b8152866004820152602081602481885afa908115612b9d578391613fb7575b50808611613bb057506001600160401b03613f6184614edc565b1682526012602052604082205480613f7f57506044613a7f84614edc565b808611613bb057506001600160401b03613f9884614edc565b168252601260205260408220613faf868254614ef0565b905586613a73565b90506020813d602011613fe1575b81613fd260209383614a4f565b81010312610528575187613f47565b3d9150613fc5565b613c9c90866155a9565b807f53ad11d80000000000000000000000000000000000000000000000000000000060049252fd5b614034915060203d6020116106ef576106e08183614a4f565b38613e3e565b906001600160a01b03612259602493614f13565b503461034f578060031936011261034f5760206001600160a01b0360035416604051908152f35b503461034f57602060031936011261034f5760ff60406020926001600160401b0361409e6148f3565b168152600e8452205460a01c166040519015158152f35b503461034f57604060031936011261034f576140cf6148f3565b6024356001600160401b038216808452600e60205260ff604085205460a01c16158015614293575b611a3257614112816000526014602052604060002054151590565b61426857508015611a23576001600160a01b0361412e83614f99565b163303611240576040517f23b872dd0000000000000000000000000000000000000000000000000000000060208201523360248201523060448201526064810182905283907f0000000000000000000000000000000000000000000000000000000000000000906141a690610828816084810161081a565b6001600160a01b036141b78561569a565b16803b156108ea576040516351b53ff760e11b81526001600160a01b039290921660048301526001600160401b03851660248301526044820184905282908290606490829084905af180156108df57614253575b5050604080516001600160401b039093168352602083019190915233917f569a440e6842b5e5a7ac02286311855f5a0b81b9390909e552e82aaf02c9e9bf91819081016119b4565b8161425d91614a4f565b6108ea57823861420b565b7f64697246000000000000000000000000000000000000000000000000000000008452600452602483fd5b5080156140f7565b503461034f5760c060031936011261034f576142b56148b1565b506142be614909565b6142c66148c7565b506084359161ffff8316830361034f5760a435906001600160401b03821161034f5760a063ffffffff8061ffff61430c88886143053660048b016149a0565b5050614d58565b95926040979194975197885216602087015216604085015216606083015215156080820152f35b503461034f578060031936011261034f57602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461034f578060031936011261034f57604051600f8054808352908352909160208301917f8d1108e10bcb7c27dddfc02ed9d693a074039d026cf4ea4240b40f7d581ac802915b8181106143d057611da785612c8a81870382614a4f565b82548452602090930192600192830192016143b9565b503461034f57604060031936011261034f576144006148f3565b60243591821515830361034f576101406144b161441d8585614cd6565b61446a60409392935180946001600160801b036080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b60a08301906001600160801b036080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565bf35b503461034f57602060031936011261034f5760206144cf6148b1565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b503461034f578060031936011261034f5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461034f578060031936011261034f5750611da760405161456e604082614a4f565b601d81527f53696c6f656455534443546f6b656e506f6f6c20312e372e302d6465760000006020820152604051918291602083526020830190614870565b503461034f57602060031936011261034f5760043590811561475f576001600160a01b036145d982614f99565b16330361474c576001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166001600160a01b0361461b8361569a565b16604051906370a0823160e01b82526004820152602081602481855afa908115612b9d57839161471a575b5080841161470357506001600160a01b036146608361569a565b1690813b156108ea57604051631d3f462b60e21b81526001600160a01b039190911660048201526001600160401b0383166024820152604481018490523360648201529082908290608490829084905af180156108df576146f3575b50906040519082825260208201527f58fca2457646a9f47422ab9eb9bff90cef88cd8b8725ab52b1d17baa392d784e60403392a280f35b816146fd91614a4f565b386146bc565b604493915063a17e11d560e01b8352600452602452fd5b90506020813d602011614744575b8161473560209383614a4f565b810103126108ea575138614646565b3d9150614728565b8063472511eb60e11b6024925233600452fd5b8063a90c0d1960e01b60049252fd5b905034610765576020600319360112610765576004357fffffffff0000000000000000000000000000000000000000000000000000000081168091036108ea57602092507faff2afbf000000000000000000000000000000000000000000000000000000008114908115614846575b811561481c575b81156147f2575b5015158152f35b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014386147eb565b7f0e64dd2900000000000000000000000000000000000000000000000000000000811491506147e4565b7f3317103100000000000000000000000000000000000000000000000000000000811491506147dd565b919082519283825260005b84811061489c575050601f19601f8460006020809697860101520116010190565b8060208092840101518282860101520161487b565b600435906001600160a01b038216820361052857565b606435906001600160a01b038216820361052857565b602435906001600160a01b038216820361052857565b600435906001600160401b038216820361052857565b602435906001600160401b038216820361052857565b3590811515820361052857565b602060408183019282815284518094520192019060005b8181106149505750505090565b82516001600160a01b0316845260209384019390920191600101614943565b6024359061ffff8216820361052857565b6044359061ffff8216820361052857565b359061ffff8216820361052857565b9181601f84011215610528578235916001600160401b038311610528576020838186019501011161052857565b602081019081106001600160401b038211176149e857604052565b634e487b7160e01b600052604160045260246000fd5b604081019081106001600160401b038211176149e857604052565b60e081019081106001600160401b038211176149e857604052565b60a081019081106001600160401b038211176149e857604052565b90601f601f1991011681019081106001600160401b038211176149e857604052565b9291926001600160401b0382116149e85760405191614a9a601f8201601f191660200184614a4f565b829481845281830111610528578281602093846000960137010152565b9080601f8301121561052857816020614ad293359101614a71565b90565b906040600319830112610528576004356001600160401b03811681036105285791602435906001600160401b03821161052857614b14916004016149a0565b9091565b9181601f84011215610528578235916001600160401b038311610528576020808501948460051b01011161052857565b9181601f84011215610528578235916001600160401b038311610528576020808501948460061b01011161052857565b6001600160401b0381116149e85760051b60200190565b9080601f8301121561052857813590614ba782614b78565b92614bb56040519485614a4f565b82845260208085019360051b82010191821161052857602001915b818310614bdd5750505090565b82356001600160a01b038116810361052857815260209283019201614bd0565b614ad2916020614c168351604084526040840190614870565b920151906020818403910152614870565b9181601f84011215610528578235916001600160401b038311610528576020808501948460081b01011161052857565b60405190614c6482614a34565b60006080838281528260208201528260408201528260608201520152565b90604051614c8f81614a34565b60806001829460ff81546001600160801b038116865263ffffffff81861c16602087015260a01c161515604085015201546001600160801b0381166060840152811c910152565b6001600160401b0391614ce7614c57565b50614cf0614c57565b50614d2457166000526008602052604060002090614ad2614d186002614d1d614d1886614c82565b615407565b9401614c82565b1690816000526004602052614d3f614d186040600020614c82565b916000526005602052614ad2614d186040600020614c82565b9061ffff8060035460a01c1691169283151592838094614ed4575b614eaa576001600160401b0316600052600b60205260406000209160405192614d9b84614a19565b5463ffffffff81168452602084019563ffffffff8260201c168752604085019263ffffffff8360401c168452606086019163ffffffff8460601c168352608087019761ffff8560801c16895260ff60a089019561ffff8160901c16875260a01c1615801560c08a0152614e8f57614e30575050505063ffffffff808061ffff9351169451169551169351169193929190600190565b819397508092945010614e5f57505063ffffffff808061ffff9351169451169551169351169193929190600190565b7f7911d95b0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b50505050505092505050600090600090600090600090600090565b7f98d50fd70000000000000000000000000000000000000000000000000000000060005260046000fd5b508215614d73565b356001600160401b03811681036105285790565b91908203918211614efd57565b634e487b7160e01b600052601160045260246000fd5b356001600160a01b03811681036105285790565b906001600160401b03614ad292166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b9190811015614f735760051b0190565b634e487b7160e01b600052603260045260246000fd5b9190811015614f735760061b0190565b6001600160401b0316600052600e60205260406000205460ff8160a01c16614fcb57506001600160a01b03600c541690565b6001600160a01b031690565b601f8260209493601f19938186528686013760008582860101520116010190565b9593919294906001600160a01b036003541695861561515857615078946001600160401b0361ffff936001600160a01b036040519b7f89720a62000000000000000000000000000000000000000000000000000000008d521660048c01521660248a0152604489015216606487015260c0608487015260c4860191614fd7565b916002821015615142578380600094819460a483015203915afa9081156127af576000916150a4575090565b903d8082843e6150b48184614a4f565b820191602081840312610765578051906001600160401b0382116108ea57019180601f84011215610765578251916150eb83614b78565b936150f96040519586614a4f565b83855260208086019460051b82010192831161076557602001925b828410615122575050505090565b83516001600160a01b03811681036108ea57815260209384019301615114565b634e487b7160e01b600052602160045260246000fd5b505050505050505060405161516e602082614a4f565b60008152600036813790565b60405190615187826149fe565b60606020838281520152565b8051821015614f735760209160051b010190565b90600182811c921680156151d7575b60208310146151c157565b634e487b7160e01b600052602260045260246000fd5b91607f16916151b6565b90604051918260008254926151f5846151a7565b8084529360018116908115615263575060011461521c575b5061521a92500383614a4f565b565b90506000929192526020600020906000915b81831061524757505090602061521a928201013861520d565b602091935080600191548385890101520191019091849261522e565b6020935061521a9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b8201013861520d565b6001600160401b03166000526008602052614ad260046040600020016151e1565b9190811015614f735760081b0190565b3580151581036105285790565b3561ffff811681036105285790565b3563ffffffff811681036105285790565b359063ffffffff8216820361052857565b35906001600160801b038216820361052857565b919082606091031261052857604051606081018181106001600160401b038211176149e857604052604061537881839561535f8161491f565b855261536d60208201615312565b602086015201615312565b910152565b6001600160801b036153b2604080936153958161491f565b15158652836153a660208301615312565b16602087015201615312565b16910152565b91908201809211614efd57565b8181106153d0575050565b600081556001016153c5565b81810292918115918404141715614efd57565b90816020910312610528575180151581036105285790565b61540f614c57565b506001600160801b036060820151166001600160801b03808351169161545a602085019361545461544763ffffffff87511642614ef0565b85608089015116906153dc565b906153b8565b8082101561547357505b16825263ffffffff4216905290565b9050615464565b9060ff8091169116039060ff8211614efd57565b60ff16604d8111614efd57600a0a90565b81156154a9570490565b634e487b7160e01b600052601260045260246000fd5b7f000000000000000000000000000000000000000000000000000000000000000060ff811690816006146155a457816006116155795760066155009161547a565b90604d60ff831611801561555e575b615527575090615521614ad29261548e565b906153dc565b90507fa9cb113d00000000000000000000000000000000000000000000000000000000600052600660045260245260445260646000fd5b506155688261548e565b80156154a95760001904831161550f565b61558490600661547a565b90604d60ff83161161552757509061559e614ad29261548e565b9061549f565b505090565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561052857018035906001600160401b0382116105285760200191813603831361052857565b60008052600e6020527fe710864318d4a32f37d6ce54cb3fadbef648dd12d8dbdf53973564d56b7f881c5460a01c60ff161561568b576001600160401b0360005b1680600052600d6020526001600160a01b036040600020541690811561565e575090565b7f4fe6a5880000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6001600160401b03600061563a565b6001600160401b038116600052600e60205260ff60406000205460a01c166000146156eb576001600160401b031680600052600d6020526001600160a01b036040600020541690811561565e575090565b506001600160401b03600061563a565b6001600160a01b0360015416330361570f57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b9080511561591d576001600160401b038151602083012092169182600052600860205261576d816005604060002001615cf4565b156158d9576000526009602052604060002081516001600160401b0381116149e85761579982546151a7565b601f81116158a7575b506020601f821160011461581d57916157f7827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea959361580d95600091615812575b506000198260011b9260031b1c19161790565b9055604051918291602083526020830190614870565b0390a2565b9050840151386157e4565b601f1982169083600052806000209160005b81811061588f57509261580d9492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610615876575b5050811b019055611d93565b85015160001960f88460031b161c19169055388061586a565b9192602060018192868a01518155019401920161582f565b6158d390836000526020600020601f840160051c81019160208510610d1c57601f0160051c01906153c5565b386157a2565b50906123906040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190614870565b630a64406560e11b60005260046000fd5b815191929115615a74576001600160801b036040840151166001600160801b0360208501511610615a1a5761521a91925b805182547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff0000000000000000000000000000000000000000161782556020810151825460409290920151608090811b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166001600160801b039290921691821760018501556001600160a01b0319909216174290911b73ffffffff0000000000000000000000000000000016179055565b606483615a72604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906001600160801b0360408092805115158552826020820151166020860152015116910152565bfd5b6001600160801b0360408401511615801590615af1575b615a995761521a919261595f565b606483615a72604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906001600160801b0360408092805115158552826020820151166020860152015116910152565b506001600160801b036020840151161515615a8b565b906040519182815491828252602082019060005260206000209260005b818110615b3957505061521a92500383614a4f565b8454835260019485019487945060209093019201615b24565b6001600160401b0316615b72816000526007602052604060002054151590565b15615bbc575033600052601060205260406000205415615b8e57565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fa9902c7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b8054821015614f735760005260206000200190600090565b805490680100000000000000008210156149e85781615c28916001615c3f94018155615be9565b81939154906000199060031b92831b921b19161790565b9055565b80600052601460205260406000205415600014615c7c57615c65816013615c01565b601354906000526014602052604060002055600190565b50600090565b80600052601060205260406000205415600014615c7c57615ca481600f615c01565b600f54906000526010602052604060002055600190565b80600052600760205260406000205415600014615c7c57615cdd816006615c01565b600654906000526007602052604060002055600190565b6000828152600182016020526040902054615d2b5780615d1683600193615c01565b80549260005201602052604060002055600190565b5050600090565b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016602082015260208152614ad2604082614a4f565b80548015615d95576000190190615d848282615be9565b60001982549160031b1b1916905555565b634e487b7160e01b600052603160045260246000fd5b6000818152600760205260409020548015615d2b576000198101818111614efd57600654906000198201918211614efd57818103615e05575b505050615df16006615d6d565b600052600760205260006040812055600190565b615e27615e16615c28936006615be9565b90549060031b1c9283926006615be9565b90556000526007602052604060002055388080615de4565b906001820191816000528260205260406000205490811515600014615ed657600019820191808311614efd5781546000198101908111614efd578381615e8d9503615e9f575b505050615d6d565b60005260205260006040812055600190565b615ebf615eaf615c289386615be9565b90549060031b1c92839286615be9565b905560005284602052604060002055388080615e85565b50505050600090565b90612710916001600160401b03615ef860208301614edc565b166000908152600b602052604090209161ffff1615615f2b57606061ffff615f27935460901c169101356153dc565b0490565b606061ffff615f27935460801c169101356153dc565b9182549060ff8260a01c16158015616177575b616171576001600160801b0382169160018501908154615f8763ffffffff6001600160801b0383169360801c1642614ef0565b90816160d3575b50508481106160945750838310615fdf575050615fb46001600160801b03928392614ef0565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b9190915460801c9283156160535781615ff791614ef0565b92600019810190808211614efd5761601a61601f926001600160a01b03966153b8565b61549f565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b6001600160a01b0383837fd0c8d23a000000000000000000000000000000000000000000000000000000006000526000196004526024521660445260646000fd5b82856001600160a01b03927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611616147576160ee926154549160801c906153dc565b808410156161425750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178655923880615f8e565b6160f9565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b508215615f54565b6000818152601060205260409020548015615d2b576000198101818111614efd57600f54906000198201918211614efd578082036161d9575b5050506161c5600f615d6d565b600052601060205260006040812055600190565b6161fb6161ea615c2893600f615be9565b90549060031b1c928392600f615be9565b905560005260106020526040600020553880806161b8565b906000602091828151910182855af1156127af576000513d61627b57506001600160a01b0381163b155b6162445750565b6001600160a01b03907f5274afe7000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b6001141561623d56fea164736f6c634300081a000a",
}

var SiloedUSDCTokenPoolABI = SiloedUSDCTokenPoolMetaData.ABI

var SiloedUSDCTokenPoolBin = SiloedUSDCTokenPoolMetaData.Bin

func DeploySiloedUSDCTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, localTokenDecimals uint8, advancedPoolHooks common.Address, rmnProxy common.Address, router common.Address, lockBox common.Address) (common.Address, *types.Transaction, *SiloedUSDCTokenPool, error) {
	parsed, err := SiloedUSDCTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SiloedUSDCTokenPoolBin), backend, token, localTokenDecimals, advancedPoolHooks, rmnProxy, router, lockBox)
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetAvailableTokens(opts *bind.CallOpts, remoteChainSelector uint64) (*big.Int, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getAvailableTokens", remoteChainSelector)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetAvailableTokens(remoteChainSelector uint64) (*big.Int, error) {
	return _SiloedUSDCTokenPool.Contract.GetAvailableTokens(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetAvailableTokens(remoteChainSelector uint64) (*big.Int, error) {
	return _SiloedUSDCTokenPool.Contract.GetAvailableTokens(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetChainRebalancer(opts *bind.CallOpts, remoteChainSelector uint64) (common.Address, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getChainRebalancer", remoteChainSelector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetChainRebalancer(remoteChainSelector uint64) (common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetChainRebalancer(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetChainRebalancer(remoteChainSelector uint64) (common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetChainRebalancer(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getCurrentRateLimiterState", remoteChainSelector, customBlockConfirmation)

	outstruct := new(GetCurrentRateLimiterState)
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutboundRateLimiterState = *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)
	outstruct.InboundRateLimiterState = *abi.ConvertType(out[1], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return *outstruct, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetCurrentRateLimiterState(remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	return _SiloedUSDCTokenPool.Contract.GetCurrentRateLimiterState(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector, customBlockConfirmation)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetCurrentRateLimiterState(remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	return _SiloedUSDCTokenPool.Contract.GetCurrentRateLimiterState(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector, customBlockConfirmation)
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getFee", arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)

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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _SiloedUSDCTokenPool.Contract.GetFee(&_SiloedUSDCTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _SiloedUSDCTokenPool.Contract.GetFee(&_SiloedUSDCTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetMinBlockConfirmation(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getMinBlockConfirmation")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetMinBlockConfirmation() (uint16, error) {
	return _SiloedUSDCTokenPool.Contract.GetMinBlockConfirmation(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetMinBlockConfirmation() (uint16, error) {
	return _SiloedUSDCTokenPool.Contract.GetMinBlockConfirmation(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetRebalancer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getRebalancer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetRebalancer() (common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetRebalancer(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetRebalancer() (common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetRebalancer(&_SiloedUSDCTokenPool.CallOpts)
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetRequiredCCVs(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getRequiredCCVs", localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetRequiredCCVs(&_SiloedUSDCTokenPool.CallOpts, localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetRequiredCCVs(&_SiloedUSDCTokenPool.CallOpts, localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetUnsiloedLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getUnsiloedLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetUnsiloedLiquidity() (*big.Int, error) {
	return _SiloedUSDCTokenPool.Contract.GetUnsiloedLiquidity(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetUnsiloedLiquidity() (*big.Int, error) {
	return _SiloedUSDCTokenPool.Contract.GetUnsiloedLiquidity(&_SiloedUSDCTokenPool.CallOpts)
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) IsSiloed(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "isSiloed", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) IsSiloed(remoteChainSelector uint64) (bool, error) {
	return _SiloedUSDCTokenPool.Contract.IsSiloed(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) IsSiloed(remoteChainSelector uint64) (bool, error) {
	return _SiloedUSDCTokenPool.Contract.IsSiloed(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.LockOrBurn0(&_SiloedUSDCTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.LockOrBurn0(&_SiloedUSDCTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationRequested, tokenArgs)
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) ProvideLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "provideLiquidity", amount)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) ProvideLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ProvideLiquidity(&_SiloedUSDCTokenPool.TransactOpts, amount)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) ProvideLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ProvideLiquidity(&_SiloedUSDCTokenPool.TransactOpts, amount)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) ProvideSiloedLiquidity(opts *bind.TransactOpts, remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "provideSiloedLiquidity", remoteChainSelector, amount)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) ProvideSiloedLiquidity(remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ProvideSiloedLiquidity(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector, amount)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) ProvideSiloedLiquidity(remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ProvideSiloedLiquidity(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector, amount)
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn, blockConfirmationRequested)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ReleaseOrMint0(&_SiloedUSDCTokenPool.TransactOpts, releaseOrMintIn, blockConfirmationRequested)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ReleaseOrMint0(&_SiloedUSDCTokenPool.TransactOpts, releaseOrMintIn, blockConfirmationRequested)
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) SetDynamicConfig(opts *bind.TransactOpts, router common.Address, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "setDynamicConfig", router, rateLimitAdmin)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) SetDynamicConfig(router common.Address, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetDynamicConfig(&_SiloedUSDCTokenPool.TransactOpts, router, rateLimitAdmin)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) SetDynamicConfig(router common.Address, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetDynamicConfig(&_SiloedUSDCTokenPool.TransactOpts, router, rateLimitAdmin)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) SetMinBlockConfirmation(opts *bind.TransactOpts, minBlockConfirmation uint16) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "setMinBlockConfirmation", minBlockConfirmation)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) SetMinBlockConfirmation(minBlockConfirmation uint16) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetMinBlockConfirmation(&_SiloedUSDCTokenPool.TransactOpts, minBlockConfirmation)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) SetMinBlockConfirmation(minBlockConfirmation uint16) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetMinBlockConfirmation(&_SiloedUSDCTokenPool.TransactOpts, minBlockConfirmation)
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) SetRebalancer(opts *bind.TransactOpts, newRebalancer common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "setRebalancer", newRebalancer)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) SetRebalancer(newRebalancer common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetRebalancer(&_SiloedUSDCTokenPool.TransactOpts, newRebalancer)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) SetRebalancer(newRebalancer common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetRebalancer(&_SiloedUSDCTokenPool.TransactOpts, newRebalancer)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) SetSiloRebalancer(opts *bind.TransactOpts, remoteChainSelector uint64, newRebalancer common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "setSiloRebalancer", remoteChainSelector, newRebalancer)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) SetSiloRebalancer(remoteChainSelector uint64, newRebalancer common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetSiloRebalancer(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector, newRebalancer)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) SetSiloRebalancer(remoteChainSelector uint64, newRebalancer common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetSiloRebalancer(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector, newRebalancer)
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) UpdateSiloDesignations(opts *bind.TransactOpts, removes []uint64, adds []SiloedLockReleaseTokenPoolSiloConfigUpdate) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "updateSiloDesignations", removes, adds)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) UpdateSiloDesignations(removes []uint64, adds []SiloedLockReleaseTokenPoolSiloConfigUpdate) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.UpdateSiloDesignations(&_SiloedUSDCTokenPool.TransactOpts, removes, adds)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) UpdateSiloDesignations(removes []uint64, adds []SiloedLockReleaseTokenPoolSiloConfigUpdate) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.UpdateSiloDesignations(&_SiloedUSDCTokenPool.TransactOpts, removes, adds)
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) WithdrawLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "withdrawLiquidity", amount)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) WithdrawLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.WithdrawLiquidity(&_SiloedUSDCTokenPool.TransactOpts, amount)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) WithdrawLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.WithdrawLiquidity(&_SiloedUSDCTokenPool.TransactOpts, amount)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) WithdrawSiloedLiquidity(opts *bind.TransactOpts, remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "withdrawSiloedLiquidity", remoteChainSelector, amount)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) WithdrawSiloedLiquidity(remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.WithdrawSiloedLiquidity(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector, amount)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) WithdrawSiloedLiquidity(remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.WithdrawSiloedLiquidity(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector, amount)
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

type SiloedUSDCTokenPoolChainSiloedIterator struct {
	Event *SiloedUSDCTokenPoolChainSiloed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolChainSiloedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolChainSiloed)
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
		it.Event = new(SiloedUSDCTokenPoolChainSiloed)
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

func (it *SiloedUSDCTokenPoolChainSiloedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolChainSiloedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolChainSiloed struct {
	RemoteChainSelector uint64
	Rebalancer          common.Address
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterChainSiloed(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolChainSiloedIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "ChainSiloed")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolChainSiloedIterator{contract: _SiloedUSDCTokenPool.contract, event: "ChainSiloed", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchChainSiloed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolChainSiloed) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "ChainSiloed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolChainSiloed)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "ChainSiloed", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseChainSiloed(log types.Log) (*SiloedUSDCTokenPoolChainSiloed, error) {
	event := new(SiloedUSDCTokenPoolChainSiloed)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "ChainSiloed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolChainUnsiloedIterator struct {
	Event *SiloedUSDCTokenPoolChainUnsiloed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolChainUnsiloedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolChainUnsiloed)
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
		it.Event = new(SiloedUSDCTokenPoolChainUnsiloed)
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

func (it *SiloedUSDCTokenPoolChainUnsiloedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolChainUnsiloedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolChainUnsiloed struct {
	RemoteChainSelector uint64
	AmountUnsiloed      *big.Int
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterChainUnsiloed(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolChainUnsiloedIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "ChainUnsiloed")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolChainUnsiloedIterator{contract: _SiloedUSDCTokenPool.contract, event: "ChainUnsiloed", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchChainUnsiloed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolChainUnsiloed) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "ChainUnsiloed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolChainUnsiloed)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "ChainUnsiloed", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseChainUnsiloed(log types.Log) (*SiloedUSDCTokenPoolChainUnsiloed, error) {
	event := new(SiloedUSDCTokenPoolChainUnsiloed)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "ChainUnsiloed", log); err != nil {
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

type SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator struct {
	Event *SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
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
		it.Event = new(SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
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

func (it *SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator{contract: _SiloedUSDCTokenPool.contract, event: "CustomBlockConfirmationInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error) {
	event := new(SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator struct {
	Event *SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
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
		it.Event = new(SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
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

func (it *SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator{contract: _SiloedUSDCTokenPool.contract, event: "CustomBlockConfirmationOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error) {
	event := new(SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
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

type SiloedUSDCTokenPoolLiquidityAddedIterator struct {
	Event *SiloedUSDCTokenPoolLiquidityAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolLiquidityAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolLiquidityAdded)
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
		it.Event = new(SiloedUSDCTokenPoolLiquidityAdded)
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

func (it *SiloedUSDCTokenPoolLiquidityAddedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolLiquidityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolLiquidityAdded struct {
	RemoteChainSelector uint64
	Provider            common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterLiquidityAdded(opts *bind.FilterOpts, provider []common.Address) (*SiloedUSDCTokenPoolLiquidityAddedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "LiquidityAdded", providerRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolLiquidityAddedIterator{contract: _SiloedUSDCTokenPool.contract, event: "LiquidityAdded", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolLiquidityAdded, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "LiquidityAdded", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolLiquidityAdded)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseLiquidityAdded(log types.Log) (*SiloedUSDCTokenPoolLiquidityAdded, error) {
	event := new(SiloedUSDCTokenPoolLiquidityAdded)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolLiquidityRemovedIterator struct {
	Event *SiloedUSDCTokenPoolLiquidityRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolLiquidityRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolLiquidityRemoved)
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
		it.Event = new(SiloedUSDCTokenPoolLiquidityRemoved)
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

func (it *SiloedUSDCTokenPoolLiquidityRemovedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolLiquidityRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolLiquidityRemoved struct {
	RemoteChainSelector uint64
	Remover             common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterLiquidityRemoved(opts *bind.FilterOpts, remover []common.Address) (*SiloedUSDCTokenPoolLiquidityRemovedIterator, error) {

	var removerRule []interface{}
	for _, removerItem := range remover {
		removerRule = append(removerRule, removerItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "LiquidityRemoved", removerRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolLiquidityRemovedIterator{contract: _SiloedUSDCTokenPool.contract, event: "LiquidityRemoved", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolLiquidityRemoved, remover []common.Address) (event.Subscription, error) {

	var removerRule []interface{}
	for _, removerItem := range remover {
		removerRule = append(removerRule, removerItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "LiquidityRemoved", removerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolLiquidityRemoved)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseLiquidityRemoved(log types.Log) (*SiloedUSDCTokenPoolLiquidityRemoved, error) {
	event := new(SiloedUSDCTokenPoolLiquidityRemoved)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
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

type SiloedUSDCTokenPoolMinBlockConfirmationSetIterator struct {
	Event *SiloedUSDCTokenPoolMinBlockConfirmationSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolMinBlockConfirmationSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolMinBlockConfirmationSet)
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
		it.Event = new(SiloedUSDCTokenPoolMinBlockConfirmationSet)
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

func (it *SiloedUSDCTokenPoolMinBlockConfirmationSetIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolMinBlockConfirmationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolMinBlockConfirmationSet struct {
	MinBlockConfirmation uint16
	Raw                  types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterMinBlockConfirmationSet(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolMinBlockConfirmationSetIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "MinBlockConfirmationSet")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolMinBlockConfirmationSetIterator{contract: _SiloedUSDCTokenPool.contract, event: "MinBlockConfirmationSet", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchMinBlockConfirmationSet(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolMinBlockConfirmationSet) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "MinBlockConfirmationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolMinBlockConfirmationSet)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "MinBlockConfirmationSet", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseMinBlockConfirmationSet(log types.Log) (*SiloedUSDCTokenPoolMinBlockConfirmationSet, error) {
	event := new(SiloedUSDCTokenPoolMinBlockConfirmationSet)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "MinBlockConfirmationSet", log); err != nil {
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
	CustomBlockConfirmation   bool
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

type SiloedUSDCTokenPoolSiloRebalancerSetIterator struct {
	Event *SiloedUSDCTokenPoolSiloRebalancerSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolSiloRebalancerSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolSiloRebalancerSet)
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
		it.Event = new(SiloedUSDCTokenPoolSiloRebalancerSet)
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

func (it *SiloedUSDCTokenPoolSiloRebalancerSetIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolSiloRebalancerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolSiloRebalancerSet struct {
	RemoteChainSelector uint64
	OldRebalancer       common.Address
	NewRebalancer       common.Address
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterSiloRebalancerSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolSiloRebalancerSetIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "SiloRebalancerSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolSiloRebalancerSetIterator{contract: _SiloedUSDCTokenPool.contract, event: "SiloRebalancerSet", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchSiloRebalancerSet(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolSiloRebalancerSet, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "SiloRebalancerSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolSiloRebalancerSet)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "SiloRebalancerSet", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseSiloRebalancerSet(log types.Log) (*SiloedUSDCTokenPoolSiloRebalancerSet, error) {
	event := new(SiloedUSDCTokenPoolSiloRebalancerSet)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "SiloRebalancerSet", log); err != nil {
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

type SiloedUSDCTokenPoolUnsiloedRebalancerSetIterator struct {
	Event *SiloedUSDCTokenPoolUnsiloedRebalancerSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolUnsiloedRebalancerSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolUnsiloedRebalancerSet)
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
		it.Event = new(SiloedUSDCTokenPoolUnsiloedRebalancerSet)
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

func (it *SiloedUSDCTokenPoolUnsiloedRebalancerSetIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolUnsiloedRebalancerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolUnsiloedRebalancerSet struct {
	OldRebalancer common.Address
	NewRebalancer common.Address
	Raw           types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterUnsiloedRebalancerSet(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolUnsiloedRebalancerSetIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "UnsiloedRebalancerSet")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolUnsiloedRebalancerSetIterator{contract: _SiloedUSDCTokenPool.contract, event: "UnsiloedRebalancerSet", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchUnsiloedRebalancerSet(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolUnsiloedRebalancerSet) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "UnsiloedRebalancerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolUnsiloedRebalancerSet)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "UnsiloedRebalancerSet", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseUnsiloedRebalancerSet(log types.Log) (*SiloedUSDCTokenPoolUnsiloedRebalancerSet, error) {
	event := new(SiloedUSDCTokenPoolUnsiloedRebalancerSet)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "UnsiloedRebalancerSet", log); err != nil {
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

func (SiloedUSDCTokenPoolChainSiloed) Topic() common.Hash {
	return common.HexToHash("0x180c6940bd64ba8f75679203ca32f8be2f629477a3307b190656e4b14dd5ddeb")
}

func (SiloedUSDCTokenPoolChainUnsiloed) Topic() common.Hash {
	return common.HexToHash("0x7b5efb3f8090c5cfd24e170b667d0e2b6fdc3db6540d75b86d5b6655ba00eb93")
}

func (SiloedUSDCTokenPoolCircleMigratorAddressSet) Topic() common.Hash {
	return common.HexToHash("0x084e6f0e9791c2e56153bd49e6ec6dd63ba9a72c258d71558d74c63fc75b7168")
}

func (SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce615")
}

func (SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac20979328")
}

func (SiloedUSDCTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x22a0dbb8195755fbfc99667a86ae684c568e9dfbb1eccf7f90084e6166447970")
}

func (SiloedUSDCTokenPoolFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (SiloedUSDCTokenPoolInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (SiloedUSDCTokenPoolLiquidityAdded) Topic() common.Hash {
	return common.HexToHash("0x569a440e6842b5e5a7ac02286311855f5a0b81b9390909e552e82aaf02c9e9bf")
}

func (SiloedUSDCTokenPoolLiquidityRemoved) Topic() common.Hash {
	return common.HexToHash("0x58fca2457646a9f47422ab9eb9bff90cef88cd8b8725ab52b1d17baa392d784e")
}

func (SiloedUSDCTokenPoolLockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
}

func (SiloedUSDCTokenPoolMinBlockConfirmationSet) Topic() common.Hash {
	return common.HexToHash("0xa7f8dbba8cdb126ce4a0e7939ec58e0161b70d808b585dd651d68e59d27e11fb")
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

func (SiloedUSDCTokenPoolSiloRebalancerSet) Topic() common.Hash {
	return common.HexToHash("0x01efd4cd7dd64263689551000d4359d6559c839f39b773b1df3fd19ff060cf5f")
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

func (SiloedUSDCTokenPoolUnsiloedRebalancerSet) Topic() common.Hash {
	return common.HexToHash("0x66b1c1bdec8b60a3442bb25b5b6cd6fff3d0eceb6f5390be8e2f82a8ad39b234")
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPool) Address() common.Address {
	return _SiloedUSDCTokenPool.address
}

type SiloedUSDCTokenPoolInterface interface {
	GetAdvancedPoolHooks(opts *bind.CallOpts) (common.Address, error)

	GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error)

	GetAvailableTokens(opts *bind.CallOpts, remoteChainSelector uint64) (*big.Int, error)

	GetChainRebalancer(opts *bind.CallOpts, remoteChainSelector uint64) (common.Address, error)

	GetCurrentProposedCCTPChainMigration(opts *bind.CallOpts) (uint64, error)

	GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

		error)

	GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

		error)

	GetExcludedTokensByChain(opts *bind.CallOpts, remoteChainSelector uint64) (*big.Int, error)

	GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

		error)

	GetMinBlockConfirmation(opts *bind.CallOpts) (uint16, error)

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

	ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error)

	ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error)

	ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error)

	BurnLockedUSDC(opts *bind.TransactOpts) (*types.Transaction, error)

	CancelExistingCCTPMigrationProposal(opts *bind.TransactOpts) (*types.Transaction, error)

	ConfigureLockBoxes(opts *bind.TransactOpts, lockBoxConfigs []SiloedLockReleaseTokenPoolLockBoxConfig) (*types.Transaction, error)

	ExcludeTokensFromBurn(opts *bind.TransactOpts, remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error)

	ProposeCCTPMigration(opts *bind.TransactOpts, remoteChainSelector uint64) (*types.Transaction, error)

	ProvideLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	ProvideSiloedLiquidity(opts *bind.TransactOpts, remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetCircleMigratorAddress(opts *bind.TransactOpts, migrator common.Address) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, router common.Address, rateLimitAdmin common.Address) (*types.Transaction, error)

	SetMinBlockConfirmation(opts *bind.TransactOpts, minBlockConfirmation uint16) (*types.Transaction, error)

	SetRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error)

	SetRebalancer(opts *bind.TransactOpts, newRebalancer common.Address) (*types.Transaction, error)

	SetSiloRebalancer(opts *bind.TransactOpts, remoteChainSelector uint64, newRebalancer common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateAdvancedPoolHooks(opts *bind.TransactOpts, newHook common.Address) (*types.Transaction, error)

	UpdateSiloDesignations(opts *bind.TransactOpts, removes []uint64, adds []SiloedLockReleaseTokenPoolSiloConfigUpdate) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error)

	WithdrawLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	WithdrawSiloedLiquidity(opts *bind.TransactOpts, remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error)

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

	FilterChainSiloed(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolChainSiloedIterator, error)

	WatchChainSiloed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolChainSiloed) (event.Subscription, error)

	ParseChainSiloed(log types.Log) (*SiloedUSDCTokenPoolChainSiloed, error)

	FilterChainUnsiloed(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolChainUnsiloedIterator, error)

	WatchChainUnsiloed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolChainUnsiloed) (event.Subscription, error)

	ParseChainUnsiloed(log types.Log) (*SiloedUSDCTokenPoolChainUnsiloed, error)

	FilterCircleMigratorAddressSet(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolCircleMigratorAddressSetIterator, error)

	WatchCircleMigratorAddressSet(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCircleMigratorAddressSet) (event.Subscription, error)

	ParseCircleMigratorAddressSet(log types.Log) (*SiloedUSDCTokenPoolCircleMigratorAddressSet, error)

	FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*SiloedUSDCTokenPoolDynamicConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*SiloedUSDCTokenPoolFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*SiloedUSDCTokenPoolFeeTokenWithdrawn, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolInboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*SiloedUSDCTokenPoolInboundRateLimitConsumed, error)

	FilterLiquidityAdded(opts *bind.FilterOpts, provider []common.Address) (*SiloedUSDCTokenPoolLiquidityAddedIterator, error)

	WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolLiquidityAdded, provider []common.Address) (event.Subscription, error)

	ParseLiquidityAdded(log types.Log) (*SiloedUSDCTokenPoolLiquidityAdded, error)

	FilterLiquidityRemoved(opts *bind.FilterOpts, remover []common.Address) (*SiloedUSDCTokenPoolLiquidityRemovedIterator, error)

	WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolLiquidityRemoved, remover []common.Address) (event.Subscription, error)

	ParseLiquidityRemoved(log types.Log) (*SiloedUSDCTokenPoolLiquidityRemoved, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolLockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*SiloedUSDCTokenPoolLockedOrBurned, error)

	FilterMinBlockConfirmationSet(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolMinBlockConfirmationSetIterator, error)

	WatchMinBlockConfirmationSet(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolMinBlockConfirmationSet) (event.Subscription, error)

	ParseMinBlockConfirmationSet(log types.Log) (*SiloedUSDCTokenPoolMinBlockConfirmationSet, error)

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

	FilterSiloRebalancerSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolSiloRebalancerSetIterator, error)

	WatchSiloRebalancerSet(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolSiloRebalancerSet, remoteChainSelector []uint64) (event.Subscription, error)

	ParseSiloRebalancerSet(log types.Log) (*SiloedUSDCTokenPoolSiloRebalancerSet, error)

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*SiloedUSDCTokenPoolTokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*SiloedUSDCTokenPoolTokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*SiloedUSDCTokenPoolTokenTransferFeeConfigUpdatedIterator, error)

	WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigUpdated(log types.Log) (*SiloedUSDCTokenPoolTokenTransferFeeConfigUpdated, error)

	FilterTokensExcludedFromBurn(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolTokensExcludedFromBurnIterator, error)

	WatchTokensExcludedFromBurn(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolTokensExcludedFromBurn, remoteChainSelector []uint64) (event.Subscription, error)

	ParseTokensExcludedFromBurn(log types.Log) (*SiloedUSDCTokenPoolTokensExcludedFromBurn, error)

	FilterUnsiloedRebalancerSet(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolUnsiloedRebalancerSetIterator, error)

	WatchUnsiloedRebalancerSet(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolUnsiloedRebalancerSet) (event.Subscription, error)

	ParseUnsiloedRebalancerSet(log types.Log) (*SiloedUSDCTokenPoolUnsiloedRebalancerSet, error)

	Address() common.Address
}
