// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hybrid_lock_release_usdc_token_pool

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated"
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

type USDCTokenPoolDomain struct {
	AllowedCaller    [32]byte
	MintRecipient    [32]byte
	DomainIdentifier uint32
	Enabled          bool
}

type USDCTokenPoolDomainUpdate struct {
	AllowedCaller     [32]byte
	MintRecipient     [32]byte
	DomainIdentifier  uint32
	DestChainSelector uint64
	Enabled           bool
}

var HybridLockReleaseUSDCTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"internalType\":\"contractITokenMessenger\"},{\"name\":\"cctpMessageTransmitterProxy\",\"type\":\"address\",\"internalType\":\"contractCCTPMessageTransmitterProxy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"previousPool\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"SUPPORTED_USDC_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnLockedUSDC\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelExistingCCTPMigrationProposal\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"excludeTokensFromBurn\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentProposedCCTPChainMigration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDomain\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structUSDCTokenPool.Domain\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExcludedTokensByChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLiquidityProvider\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLockedTokensForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_localDomainIdentifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_messageTransmitterProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractCCTPMessageTransmitterProxy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_previousPool\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_tokenMessenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITokenMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposeCCTPMigration\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"provideLiquidity\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCircleMigratorAddress\",\"inputs\":[{\"name\":\"migrator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDomains\",\"inputs\":[{\"name\":\"domains\",\"type\":\"tuple[]\",\"internalType\":\"structUSDCTokenPool.DomainUpdate[]\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLiquidityProvider\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"liquidityProvider\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRouter\",\"inputs\":[{\"name\":\"newRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shouldUseLockRelease\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateChainSelectorMechanisms\",\"inputs\":[{\"name\":\"removes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"adds\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawLiquidity\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCTPMigrationCancelled\",\"inputs\":[{\"name\":\"existingProposalSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCTPMigrationExecuted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"USDCBurned\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCTPMigrationProposed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CircleMigratorAddressSet\",\"inputs\":[{\"name\":\"migratorAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DomainsSet\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structUSDCTokenPool.DomainUpdate[]\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityAdded\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityProviderSet\",\"inputs\":[{\"name\":\"oldProvider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newProvider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityRemoved\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockReleaseDisabled\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockReleaseEnabled\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RouterUpdated\",\"inputs\":[{\"name\":\"oldRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensExcludedFromBurn\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"burnableAmountAfterExclusion\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"ExistingMigrationProposal\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidChainSelector\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestinationDomain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidDomain\",\"inputs\":[{\"name\":\"domain\",\"type\":\"tuple\",\"internalType\":\"structUSDCTokenPool.DomainUpdate\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"type\":\"error\",\"name\":\"InvalidMessageLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidNonce\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"got\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidPreviousPool\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourceDomain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenMessengerVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidTransmitterInProxy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LanePausedForCCTPMigration\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoMigrationProposalPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenLockingNotAllowedAfterMigration\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnknownDomain\",\"inputs\":[{\"name\":\"domain\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"UnlockingUSDCFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"onlyCircle\",\"inputs\":[]}]",
	Bin: "0x6101a0806040523461063c57616533803803809161001d828561097d565b833981019060e08183031261063c5780516001600160a01b0381169081810361063c576020830151906001600160a01b03821680830361063c576040850151956001600160a01b0387169586880361063c5760608101516001600160401b03811161063c5781019180601f8401121561063c578251926001600160401b0384116105d9578360051b9060208201946100b8604051968761097d565b855260208086019282010192831161063c57602001905b828210610965575050506100e5608082016109a0565b6100fd60c06100f660a085016109a0565b93016109a0565b98331561095457600180546001600160a01b0319163317905588158015610943575b8015610932575b6109215760805260c05260405163313ce56760e01b81526020816004818b5afa80916000916108e5575b50906108c1575b50600660a052600480546001600160a01b0319166001600160a01b03929092169190911790558051151560e0819052610794575b50831561078357604051632c12192160e01b8152602081600481885afa90811561064957600091610749575b5060405163054fd4d560e41b81526001600160a01b03919091169190602081600481865afa80156106495763ffffffff9160009161072a575b5016806107165750604051639cdbb18160e01b8152602081600481895afa80156106495763ffffffff916000916106f7575b5016806106e357506020600491604051928380926367e0ed8360e11b82525afa8015610649578291600091610695575b506001600160a01b0316036106845760049260209261010052610120526040519283809263234d8e3d60e21b82525afa90811561064957600091610655575b506101405260805161010051604051636eb1769f60e11b81523060048201526001600160a01b0391821660248201819052959290911690602081604481855afa90811561064957600091610612575b5060001981018091116105fc5760405190602082019663095ea7b360e01b8852602483015260448201526044815261031960648261097d565b60008060409788519361032c8a8661097d565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082865af13d156105ef573d906001600160401b0382116105d957875161039d94909261038e601f8201601f19166020018561097d565b83523d6000602085013e610b72565b80519081610565575b50506001600160a01b0381163081146104dc5780151590816104ed575b506104dc577f2e902d38f15b233cbb63711add0fca4545334d3a169d60c0a616494d7eea954491602091610160528451908152a161018052516158f09081610c43823960805181818161051c015281816108270152818161089401528181611cfb01528181612b4f01528181612e9101528181613f3a0152818161434f015281816153cb01526155bf015260a051816108de015260c051818181612195015281816147aa0152614eb6015260e051818181610d7d01528181612382015261542a015261010051818181610f170152612f250152610120518181816117df0152613ebb01526101405181818161103801528181612ff101526149520152610160518181816109aa0152613e39015261018051816114e60152f35b632d4d3c3d60e21b60005260046000fd5b85516301ffc9a760e01b8152630e64dd2960e01b60048201529150602090829060249082905afa90811561055a5760009161052b575b5015386103c3565b61054d915060203d602011610553575b610545818361097d565b8101906109d0565b38610523565b503d61053b565b85513d6000823e3d90fd5b6020806105769383010191016109d0565b156105825738806103a6565b835162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b634e487b7160e01b600052604160045260246000fd5b9161039d92606091610b72565b634e487b7160e01b600052601160045260246000fd5b90506020813d602011610641575b8161062d6020938361097d565b8101031261063c5751386102e0565b600080fd5b3d9150610620565b6040513d6000823e3d90fd5b610677915060203d60201161067d575b61066f818361097d565b8101906109b4565b38610291565b503d610665565b632a32133b60e11b60005260046000fd5b9091506020813d6020116106db575b816106b16020938361097d565b810103126106d75751906001600160a01b03821682036106d45750819038610252565b80fd5b5080fd5b3d91506106a4565b6316ba39c560e31b60005260045260246000fd5b610710915060203d60201161067d5761066f818361097d565b38610222565b6334697c6b60e11b60005260045260246000fd5b610743915060203d60201161067d5761066f818361097d565b386101f0565b90506020813d60201161077b575b816107646020938361097d565b8101031261063c57610775906109a0565b386101b7565b3d9150610757565b6306b7c75960e31b60005260046000fd5b909194602094604051946107a8878761097d565b60008652600036813760e051156108b05760005b8651811015610823576001906001600160a01b036107da828a6109e8565b5116896107e682610a2a565b6107f3575b5050016107bc565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a138896107eb565b509193969092945060005b85518110156108a2576001906001600160a01b0361084c82896109e8565b5116801561089c578861085e82610b12565b61086c575b50505b0161082e565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a13888610863565b50610866565b50929591945092503861018b565b6335f4a7b360e01b60005260046000fd5b60ff1660068114610157576332ad3e0760e11b600052600660045260245260446000fd5b6020813d602011610919575b816108fe6020938361097d565b810103126106d757519060ff821682036106d4575038610150565b3d91506108f1565b6342bcdf7f60e11b60005260046000fd5b506001600160a01b03821615610126565b506001600160a01b0383161561011f565b639b15e16f60e01b60005260046000fd5b60208091610972846109a0565b8152019101906100cf565b601f909101601f19168101906001600160401b038211908210176105d957604052565b51906001600160a01b038216820361063c57565b9081602091031261063c575163ffffffff8116810361063c5790565b9081602091031261063c5751801515810361063c5790565b80518210156109fc5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b80548210156109fc5760005260206000200190600090565b6000818152600360205260409020548015610b0b5760001981018181116105fc576002546000198101919082116105fc57818103610aba575b5050506002548015610aa45760001901610a7e816002610a12565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b610af3610acb610adc936002610a12565b90549060031b1c9283926002610a12565b819391549060031b91821b91600019901b19161790565b90556000526003602052604060002055388080610a63565b5050600090565b80600052600360205260406000205415600014610b6c57600254680100000000000000008110156105d957610b53610adc8260018594016002556002610a12565b9055600254906000526003602052604060002055600190565b50600090565b91929015610bd45750815115610b86575090565b3b15610b8f5790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b825190915015610be75750805190602001fd5b6040519062461bcd60e51b8252602060048301528181519182602483015260005b838110610c2a5750508160006044809484010152601f80199101168101030190fd5b60208282018101516044878401015285935001610c0856fe6080604052600436101561001257600080fd5b60003560e01c806301ffc9a7146103475780631101dbd414610342578063181f5a771461033d57806321df0da714610338578063240028e81461033357806324f65ee71461032e5780632cfbb1191461032957806339077537146103245780633e591f2c1461031f5780634ad01f0b1461031a5780634c5ef0ed146103155780634c93ef841461031057806350d1a35a1461030b57806354c8a4f3146103065780636155cda01461030157806362ddd3c4146102fc5780636b716b0d146102f75780636b795423146102f25780636d3d1a58146102ed578063714bf907146102e857806379ba5097146102e35780637d54534e146102de5780638926f54f146102d95780638a5e52bb146102d45780638da5cb5b146102cf578063962d4020146102ca57806398db9643146102c55780639a4575b9146102c05780639fdf13ff146102bb578063a42a7b8b146102b6578063a7cd63b7146102b1578063acfecf91146102ac578063af58d59f146102a7578063b0f479a1146102a2578063b79465801461029d578063bb5eced314610298578063c0d7865514610293578063c4bffe2b1461028e578063c75eea9c14610289578063c781d0e314610284578063cd306a6c1461027f578063cf7401f31461027a578063dc0bd97114610275578063de814c5714610270578063dfadfa351461026b578063e0351e1314610266578063e8a1da1714610261578063e94ae6d01461025c578063f2fde38b14610257578063f65a8886146102525763fd6768551461024d57600080fd5b6128ab565b61286c565b6127b0565b612768565b6123a7565b61236a565b6122a1565b6121b9565b612175565b6120a3565b611f99565b611f3c565b611ef1565b611e7c565b611d74565b611c97565b611c5b565b611c34565b611b91565b611a5d565b6119f6565b611912565b611878565b611803565b6117bf565b611671565b611619565b6113fb565b6113bc565b611338565b611287565b611203565b6111dc565b61105c565b61101b565b610f98565b610ef7565b610d4b565b610b82565b610b42565b610af9565b6109ce565b61098a565b610941565b610902565b6108c4565b610867565b610807565b6107a4565b61043d565b34610419576020600319360112610419576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361041957807faff2afbf00000000000000000000000000000000000000000000000000000000602092149081156103ef575b81156103c5575b506040519015158152f35b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014386103ba565b7f0e64dd2900000000000000000000000000000000000000000000000000000000811491506103b3565b600080fd5b67ffffffffffffffff81160361041957565b359061043b8261041e565b565b346104195760406003193601126104195760043561045a8161041e565b6024359061048c61047f8267ffffffffffffffff166000526011602052604060002090565b546001600160a01b031690565b6001600160a01b03339116036105db5767ffffffffffffffff81166104be816000526010602052604060002054151590565b6105a357600b546104e49060a01c67ffffffffffffffff165b67ffffffffffffffff1690565b14610568576105079067ffffffffffffffff16600052600c602052604060002090565b610512828254612985565b90556105408130337f0000000000000000000000000000000000000000000000000000000000000000613b08565b337fc17cea59c2955cb181b03393209566960365771dbba9dc3d510180e7cb312088600080a3005b7fd0da86c40000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b6000fd5b7f646972460000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b7f8e4a23d6000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b600091031261041957565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6060810190811067ffffffffffffffff82111761065f57604052565b610614565b6080810190811067ffffffffffffffff82111761065f57604052565b6020810190811067ffffffffffffffff82111761065f57604052565b6040810190811067ffffffffffffffff82111761065f57604052565b60a0810190811067ffffffffffffffff82111761065f57604052565b90601f601f19910116810190811067ffffffffffffffff82111761065f57604052565b6040519061043b60a0836106d4565b6040519061043b6020836106d4565b6040519061043b6040836106d4565b6040519061043b6080836106d4565b67ffffffffffffffff811161065f57601f01601f191660200190565b919082519283825260005b84811061077b575050601f19601f8460006020809697860101520116010190565b8060208092840101518282860101520161075a565b9060206107a192818152019061074f565b90565b346104195760006003193601126104195761080360408051906107c781836106d4565b601382527f55534443546f6b656e506f6f6c20312e362e320000000000000000000000000060208301525191829160208352602083019061074f565b0390f35b346104195760006003193601126104195760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b6001600160a01b0381160361041957565b359061043b8261084b565b346104195760206003193601126104195760206108ba6004356108898161084b565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691161490565b6040519015158152f35b3461041957600060031936011261041957602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346104195760206003193601126104195767ffffffffffffffff6004356109288161041e565b16600052600c6020526020604060002054604051908152f35b346104195760206003193601126104195760043567ffffffffffffffff8111610419576101006003198236030112610419576109816020916004016129f6565b60405190518152f35b346104195760006003193601126104195760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610419576000600319360112610419576109e7614073565b600b5467ffffffffffffffff8160a01c168015610a7d577fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff7f375f1ad1194a2bec317c5efec05cc63ffa06ddd0c4b276619f6fd47298eda5189216600b556000610a658267ffffffffffffffff16600052600d602052604060002090565b5560405167ffffffffffffffff9091168152602090a1005b7fa94cb9880000000000000000000000000000000000000000000000000000000060005260046000fd5b929192610ab382610733565b91610ac160405193846106d4565b829481845281830111610419578281602093846000960137010152565b9080601f83011215610419578160206107a193359101610aa7565b3461041957604060031936011261041957600435610b168161041e565b60243567ffffffffffffffff811161041957602091610b3c6108ba923690600401610ade565b90612c29565b346104195760206003193601126104195760206108ba600435610b648161041e565b67ffffffffffffffff16600052600e60205260ff6040600020541690565b3461041957602060031936011261041957600435610b9f8161041e565b610ba7614073565b67ffffffffffffffff600b5460a01c16610c9e5767ffffffffffffffff81166000908152600e602052604090205460ff1615610c7457610c6f81610c547f20331f191af84dbff48b162aa5a5985e7891ae646297b0a2ac80487f9109ef49937fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff7bffffffffffffffff0000000000000000000000000000000000000000600b549260a01b16911617600b55565b60405167ffffffffffffffff90911681529081906020820190565b0390a1005b7f656535ce0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f692bc1310000000000000000000000000000000000000000000000000000000060005260046000fd5b9181601f840112156104195782359167ffffffffffffffff8311610419576020808501948460051b01011161041957565b60406003198201126104195760043567ffffffffffffffff81116104195781610d2491600401610cc8565b929092916024359067ffffffffffffffff821161041957610d4791600401610cc8565b9091565b3461041957610d73610d7b610d5f36610cf9565b9491610d6c939193614073565b3691612c7e565b923691612c7e565b7f000000000000000000000000000000000000000000000000000000000000000015610ecd5760005b8251811015610e355780610dca610dbd6001938661317a565b516001600160a01b031690565b610dec610de76001600160a01b0383165b6001600160a01b031690565b61507a565b610df8575b5001610da4565b6040516001600160a01b039190911681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756690602090a138610df1565b5060005b8151811015610ecb5780610e52610dbd6001938561317a565b6001600160a01b03811615610ec557610e7b610e766001600160a01b038316610ddb565b614ad7565b610e88575b505b01610e39565b6040516001600160a01b039190911681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d890602090a183610e80565b50610e82565b005b7f35f4a7b30000000000000000000000000000000000000000000000000000000060005260046000fd5b346104195760006003193601126104195760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b604060031982011261041957600435610f538161041e565b9160243567ffffffffffffffff811161041957826023820112156104195780600401359267ffffffffffffffff84116104195760248483010111610419576024019190565b3461041957610fa636610f3b565b610fb1929192614073565b67ffffffffffffffff8216610fd3816000526006602052604060002054151590565b15610fee5750610ecb92610fe8913691610aa7565b906140d2565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b3461041957600060031936011261041957602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346104195761106a36610cf9565b929091611075614073565b60005b81811061117f5750505060005b82811061108e57005b6110ba6110a76104d76110a2848787612d03565b612d18565b6000526010602052604060002054151590565b61113857806110fc6110ef6110d56110a26001958888612d03565b67ffffffffffffffff16600052600e602052604060002090565b600160ff19825416179055565b61110d6104d76110a2838787612d03565b7f5e3985e51df58346365017cae614e59d723143b71c9a2ce4a156687f1f2c3f5a600080a201611085565b6110a29061059f9361114993612d03565b7f646972460000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b806111a06111966110d56110a26001958789612d03565b60ff198154169055565b6111b16104d76110a2838688612d03565b7fddc5afbc5e53c63a556964db0eef76a1c2d9305e0811abd7410d2a6f4799490e600080a201611078565b346104195760006003193601126104195760206001600160a01b0360095416604051908152f35b34610419576020600319360112610419577f084e6f0e9791c2e56153bd49e6ec6dd63ba9a72c258d71558d74c63fc75b716860206001600160a01b0360043561124b8161084b565b611253614073565b16807fffffffffffffffffffffffff0000000000000000000000000000000000000000600b541617600b55604051908152a1005b34610419576000600319360112610419576000546001600160a01b038116330361130e577fffffffffffffffffffffffff0000000000000000000000000000000000000000600154913382841617600155166000556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b34610419576020600319360112610419577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d0917460206001600160a01b036004356113808161084b565b611388614073565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000006009541617600955604051908152a1005b346104195760206003193601126104195760206108ba67ffffffffffffffff6004356113e78161041e565b166000526006602052604060002054151590565b3461041957600060031936011261041957600b546114216001600160a01b038216610ddb565b33036115ef5760a01c67ffffffffffffffff1667ffffffffffffffff8116908115610a7d5761148d6114678267ffffffffffffffff16600052600c602052604060002090565b546114868367ffffffffffffffff16600052600d602052604060002090565b5490612d52565b9060006114ae8267ffffffffffffffff16600052600c602052604060002090565b556114dc7fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff600b5416600b55565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001692833b1561041957600060405180957f42966c6800000000000000000000000000000000000000000000000000000000825281838161154d89600483019190602083019252565b03925af19081156115ea577fdea60ddd4c7ebdab804f5694c70350cca7893ece3efeecb142312eacac5c73e4946115aa926115cf575b506115a56111968467ffffffffffffffff16600052600e602052604060002090565b614b4a565b506040805167ffffffffffffffff909216825260208201929092529081908101610c6f565b806115de60006115e4936106d4565b80610609565b38611583565b612d5f565b7f438a7a050000000000000000000000000000000000000000000000000000000060005260046000fd5b346104195760006003193601126104195760206001600160a01b0360015416604051908152f35b9181601f840112156104195782359167ffffffffffffffff8311610419576020808501946060850201011161041957565b346104195760606003193601126104195760043567ffffffffffffffff8111610419576116a2903690600401610cc8565b9060243567ffffffffffffffff8111610419576116c3903690600401611640565b9060443567ffffffffffffffff8111610419576116e4903690600401611640565b6116f9610ddb6009546001600160a01b031690565b331415806117a1575b6105db57838614801590611797575b61176d5760005b86811061172157005b806117676117356110a26001948b8b612d03565b611740838989612d6b565b61176161175961175186898b612d6b565b92369061205a565b91369061205a565b91614197565b01611718565b7f568efce20000000000000000000000000000000000000000000000000000000060005260046000fd5b5080861415611711565b506117b7610ddb6001546001600160a01b031690565b331415611702565b346104195760006003193601126104195760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b346104195760206003193601126104195760043567ffffffffffffffff81116104195760a060031982360301126104195761184361080391600401612d94565b6040519182916020835260206118648251604083870152606086019061074f565b910151601f1984830301604085015261074f565b3461041957600060031936011261041957602060405160008152f35b602081016020825282518091526040820191602060408360051b8301019401926000915b8383106118c757505050505090565b9091929394602080611903837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08660019603018752895161074f565b970193019301919392906118b8565b346104195760206003193601126104195767ffffffffffffffff6004356119388161041e565b1660005260076020526119516005604060002001614fd8565b805190601f1961197961196384612c66565b9361197160405195866106d4565b808552612c66565b0160005b8181106119e557505060005b81518110156119d757806119bb6119b66119a56001948661317a565b516000526008602052604060002090565b6131e1565b6119c5828661317a565b526119d0818561317a565b5001611989565b604051806108038582611894565b80606060208093870101520161197d565b3461041957600060031936011261041957611a0f614f42565b60405180916020820160208352815180915260206040840192019060005b818110611a3b575050500390f35b82516001600160a01b0316845285945060209384019390920191600101611a2d565b3461041957611a6b36610f3b565b611a76929192614073565b67ffffffffffffffff821691611aa0611a9c846000526006602052604060002054151590565b1590565b611b5a57611ae3611a9c6005611aca8467ffffffffffffffff166000526007602052604060002090565b01611ad6368689610aa7565b60208151910120906151a6565b611b1f57507f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d769192611b1a604051928392836132c5565b0390a2005b611b5684926040519384937f74f23c7c000000000000000000000000000000000000000000000000000000008552600485016132a4565b0390fd5b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b346104195760206003193601126104195767ffffffffffffffff600435611bb78161041e565b611bbf6132d6565b50166000526007602052610803611be4611bdf6002604060002001613301565b6143ca565b6040519182918291909160806001600160801b038160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b346104195760006003193601126104195760206001600160a01b0360045416604051908152f35b3461041957602060031936011261041957610803611c83600435611c7e8161041e565b613352565b60405191829160208352602083019061074f565b3461041957604060031936011261041957600435611cb48161041e565b60243590611cc0614073565b67ffffffffffffffff80600b5460a01c169116908114611d4757600052600c6020526040600020611cf2828254612d52565b9055611d1f81337f0000000000000000000000000000000000000000000000000000000000000000614448565b337fc2c3f06e49b9f15e7b4af9055e183b0d73362e033ad82a07dec9bf9840171719600080a3005b7fd0da86c40000000000000000000000000000000000000000000000000000000060005260045260246000fd5b34610419576020600319360112610419576001600160a01b03600435611d998161084b565b611da1614073565b168015611e0e5760407f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f16849160045490807fffffffffffffffffffffffff00000000000000000000000000000000000000008316176004556001600160a01b038351921682526020820152a1005b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b602060408183019282815284518094520192019060005b818110611e5c5750505090565b825167ffffffffffffffff16845260209384019390920191600101611e4f565b3461041957600060031936011261041957611e95614f8d565b805190601f19611ea761196384612c66565b0136602084013760005b8151811015611ee3578067ffffffffffffffff611ed06001938561317a565b5116611edc828661317a565b5201611eb1565b604051806108038582611e38565b346104195760206003193601126104195767ffffffffffffffff600435611f178161041e565b611f1f6132d6565b50166000526007602052610803611be4611bdf6040600020613301565b346104195760206003193601126104195760043567ffffffffffffffff8111610419573660238201121561041957806004013567ffffffffffffffff81116104195736602460a0830284010111610419576024610ecb9201613374565b3461041957600060031936011261041957602067ffffffffffffffff600b5460a01c16604051908152f35b8015150361041957565b35906001600160801b038216820361041957565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c6060910112610419576040519061201982610643565b8160843561202681611fc4565b815260a4356001600160801b038116810361041957602082015260c435906001600160801b03821682036104195760400152565b91908260609103126104195760405161207281610643565b604061209e818395803561208581611fc4565b855261209360208201611fce565b602086015201611fce565b910152565b346104195760e0600319360112610419576004356120c08161041e565b60607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc360112610419576040516120f681610643565b60243561210281611fc4565b81526044356001600160801b03811681036104195760208201526064356001600160801b038116810361041957604082015261213d36611fe2565b906001600160a01b036009541633141580612160575b6105db57610ecb92614197565b506001600160a01b0360015416331415612153565b346104195760006003193601126104195760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610419576040600319360112610419576004356121d68161041e565b602435906121e2614073565b67ffffffffffffffff600b5460a01c169167ffffffffffffffff8216809303610a7d5782600052600d602052604060002080549282840180941161229c577fe1e6c22ce6b566f66cdb457ec2e7910ff1f9a9e5654ed75303476fa8704682209361228492556114866122688267ffffffffffffffff16600052600c602052604060002090565b549167ffffffffffffffff16600052600d602052604060002090565b60408051928352602083019190915281908101611b1a565b612956565b346104195760206003193601126104195767ffffffffffffffff6004356122c78161041e565b600060606040516122d781610664565b828152826020820152826040820152015216600052600a602052610803604060002060ff60026040519261230a84610664565b8054845260018101546020850152015463ffffffff8116604084015260201c1615156060820152604051918291829190916060806080830194805184526020810151602085015263ffffffff604082015116604085015201511515910152565b346104195760006003193601126104195760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b34610419576123b536610cf9565b9190926123c0614073565b6000915b80831061262f5750505060009163ffffffff4216925b8281106123e357005b6123f66123f1828585613776565b613835565b906060820161240581516144cc565b608083019361241485516144cc565b604084019081515115611e0e57612441611a9c61243c6104d7885167ffffffffffffffff1690565b614ba0565b6125e45761254461247a612460879a999a5167ffffffffffffffff1690565b67ffffffffffffffff166000526007602052604060002090565b6125108961250a87516124fa61249a60408301516001600160801b031690565b916124ea6124bc6124b560208401516001600160801b031690565b9251151590565b6124e16124c76106f7565b6001600160801b03851681529763ffffffff166020890152565b15156040870152565b6001600160801b03166060850152565b6001600160801b03166080830152565b826138c4565b612539896125308a516124fa61249a60408301516001600160801b031690565b600283016138c4565b6004845191016139b8565b602085019660005b88518051821015612587579061258160019261257a836125748c5167ffffffffffffffff1690565b9261317a565b51906140d2565b0161254c565b505097965094906125db7f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c293926125c86001975167ffffffffffffffff1690565b9251935190519060405194859485613a85565b0390a1016123da565b61059f6125f9865167ffffffffffffffff1690565b7f1d5ad3c50000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b9091926126406110a2858486612d03565b94612657611a9c67ffffffffffffffff881661511b565b61273057612684600561267e8867ffffffffffffffff166000526007602052604060002090565b01614fd8565b9360005b85518110156126d0576001906126c960056126b78b67ffffffffffffffff166000526007602052604060002090565b016126c2838a61317a565b51906151a6565b5001612688565b509350937f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d85991661272260019397610c5461271d8267ffffffffffffffff166000526007602052604060002090565b6136e5565b0390a10191909392936123c4565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff861660045260246000fd5b346104195760206003193601126104195767ffffffffffffffff60043561278e8161041e565b16600052601160205260206001600160a01b0360406000205416604051908152f35b34610419576020600319360112610419576001600160a01b036004356127d58161084b565b6127dd614073565b1633811461284257807fffffffffffffffffffffffff000000000000000000000000000000000000000060005416176000556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346104195760206003193601126104195767ffffffffffffffff6004356128928161041e565b16600052600d6020526020604060002054604051908152f35b34610419576040600319360112610419576004356128c88161041e565b67ffffffffffffffff602435916128de8361084b565b6128e6614073565b16600081815260116020526040812080546001600160a01b039485167fffffffffffffffffffffffff0000000000000000000000000000000000000000821681179092559293909216907fc82aa48e67c70b1ad1494533456f52504bb4d62d11bbdafaeb98cfccd1ed817e9080a4005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9190820180921161229c57565b6040519061299f82610680565b60008252565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610419570180359067ffffffffffffffff82116104195760200191813603831361041957565b604051612a0281610680565b600090527ffa7c07de000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000612a5560c08401846129a5565b90358281169160048110612bf6575b50501603612bed57612a74612992565b50606081013590612a858282614708565b600b5460a01c67ffffffffffffffff166020820190612aa66104d783612d18565b67ffffffffffffffff821614610568575067ffffffffffffffff81612b08612aee7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc094612d18565b67ffffffffffffffff16600052600c602052604060002090565b54612bcf57612b33612b1982612d18565b67ffffffffffffffff16600052600d602052604060002090565b612b3e868254612d52565b90555b612bbe85612b8a612b8460407f00000000000000000000000000000000000000000000000000000000000000009801946110a284612b7e88613db9565b8b614448565b93613db9565b604080516001600160a01b039889168152336020820152919097169681019690965260608601529116929081906080820190565b0390a2612bc9610706565b90815290565b612bdb612aee82612d18565b612be6868254612d52565b9055612b41565b6107a190613dc3565b839250829060040360031b1b16163880612a64565b91612c25918354906000199060031b92831b921b19161790565b9055565b9067ffffffffffffffff6107a192166000526007602052600560406000200190602081519101209060019160005201602052604060002054151590565b67ffffffffffffffff811161065f5760051b60200190565b929190612c8a81612c66565b93612c9860405195866106d4565b602085838152019160051b810192831161041957905b828210612cba57505050565b602080918335612cc98161084b565b815201910190612cae565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9190811015612d135760051b0190565b612cd4565b356107a18161041e565b67ffffffffffffffff6107a191166000526006602052604060002054151590565b90600019820191821161229c57565b9190820391821161229c57565b6040513d6000823e3d90fd5b9190811015612d13576060020190565b60405190612d888261069c565b60606020838281520152565b612d9c612d7b565b5060208101612daf8135610b648161041e565b15612df157612dd76104d7612dd1600b5467ffffffffffffffff9060a01c1690565b92612d18565b67ffffffffffffffff82161461056857506107a1906142d9565b612dfc929192612d7b565b50612e0683614e79565b612e34612e2f612e1583612d18565b67ffffffffffffffff16600052600a602052604060002090565b61364b565b612e44611a9c6060830151151590565b613138576020612e5485806129a5565b9050036130f85760208101519293612f179380156130d8576060602091925b013591612e87604085015163ffffffff1690565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169687955160405198899485947ff856ddb60000000000000000000000000000000000000000000000000000000086528860048701919360809363ffffffff6001600160a01b039398979660a0860199865216602085015260408401521660608201520152565b038160006001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165af19384156115ea57600094613067575b50611c7e83613054937ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff612fce95612fc6612f9c6130469a612d18565b604080516001600160a01b0390971687523360208801528601929092529116929081906060820190565b0390a2612d18565b92612fea612fda610715565b67ffffffffffffffff9092168252565b63ffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015260405192839160208301919091602063ffffffff81604084019567ffffffffffffffff8151168552015116910152565b03601f1981018352826106d4565b61305c610715565b918252602082015290565b612fce91945061304693613054937ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff6130c3611c7e9560203d6020116130d1575b6130bb81836106d4565b8101906142c4565b989550505093509350612f56565b503d6130b1565b50602060606130f26130ea84806129a5565b8101906142b5565b92612e73565b61310284806129a5565b90611b566040519283927fa3c8cf09000000000000000000000000000000000000000000000000000000008452600484016132c5565b61059f61314483612d18565b7fd201c48a0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b8051821015612d135760209160051b010190565b90600182811c921680156131d7575b60208310146131a857565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f169161319d565b90604051918260008254926131f58461318e565b8084529360018116908115613261575060011461321a575b5061043b925003836106d4565b90506000929192526020600020906000915b81831061324557505090602061043b928201013861320d565b602091935080600191548385890101520191019091849261322c565b6020935061043b95925060ff1991501682840152151560051b8201013861320d565b601f8260209493601f19938186528686013760008582860101520116010190565b60409067ffffffffffffffff6107a195931681528160208201520191613283565b9160206107a1938181520191613283565b604051906132e3826106b8565b60006080838281528260208201528260408201528260608201520152565b9060405161330e816106b8565b60806001600160801b036001839560ff8154848116875263ffffffff81871c16602088015260a01c1615156040860152015481808216166060850152821c16910152565b67ffffffffffffffff1660005260076020526107a160046040600020016131e1565b61337c614073565b60005b8281106133be5750907fe6d14ea297366c7bc1265d289d924bfd8b9afb148eb972b481f70da41c842cf5916133b9604051928392836135c2565b0390a1565b6133d16133cc828585613543565b613564565b805115801561351d575b6134b057906134aa82613450612e1560606133ff6040600198015163ffffffff1690565b93613441602082015161343983519761341b6080860151151590565b92613424610724565b998a5260208a015263ffffffff166040890152565b151586840152565b015167ffffffffffffffff1690565b6002908251815560208301516001820155019063ffffffff6040820151167fffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000064ff0000000060608554940151151560201b16921617179055565b0161337f565b604080517fa606c63500000000000000000000000000000000000000000000000000000000815282516004820152602083015160248201529082015163ffffffff166044820152606082015167ffffffffffffffff1660648201526080909101511515608482015260a490fd5b5067ffffffffffffffff61353c606083015167ffffffffffffffff1690565b16156133db565b9190811015612d135760a0020190565b359063ffffffff8216820361041957565b60a0813603126104195760806040519161357d836106b8565b803583526020810135602084015261359760408201613553565b604084015260608101356135aa8161041e565b606084015201356135ba81611fc4565b608082015290565b602080825281018390526040019160005b8181106135e05750505090565b90919260a080600192863581526020870135602082015263ffffffff61360860408901613553565b16604082015267ffffffffffffffff60608801356136258161041e565b166060820152608087013561363981611fc4565b151560808201520194019291016135d3565b9060405161365881610664565b606060ff600283958054855260018101546020860152015463ffffffff8116604085015260201c161515910152565b818110613692575050565b60008155600101613687565b8181029291811591840414171561229c57565b80549060008155816136c1575050565b6000526020600020908101905b8181106136d9575050565b600081556001016136ce565b600561043b916000815560006001820155600060028201556000600382015560048101613712815461318e565b9081613721575b5050016136b1565b81601f600093116001146137395750555b3880613719565b8183526020832061375491601f01861c810190600101613687565b808252602082209081548360011b906000198560031b1c191617905555613732565b9190811015612d135760051b810135907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee181360301821215610419570190565b9080601f830112156104195781356137cd81612c66565b926137db60405194856106d4565b81845260208085019260051b820101918383116104195760208201905b83821061380757505050505090565b813567ffffffffffffffff81116104195760209161382a87848094880101610ade565b8152019101906137f8565b61012081360312610419576040519061384d826106b8565b61385681610430565b8252602081013567ffffffffffffffff81116104195761387990369083016137b6565b602083015260408101359067ffffffffffffffff8211610419576138a36135ba9236908301610ade565b60408401526138b5366060830161205a565b606084015260c036910161205a565b8151815460208401516040850151608091821b73ffffffff00000000000000000000000000000000166001600160801b039485167fffffffffffffffffffffff000000000000000000000000000000000000000000909416939093179290921791151560a01b74ff000000000000000000000000000000000000000016919091178355606084015193810151901b6fffffffffffffffffffffffffffffffff1916921691909117600190910155565b9190601f811161398257505050565b61043b926000526020600020906020601f840160051c830193106139ae575b601f0160051c0190613687565b90915081906139a1565b919091825167ffffffffffffffff811161065f576139e0816139da845461318e565b84613973565b6020601f8211600114613a1c578190612c25939495600092613a11575b50506000198260011b9260031b1c19161790565b0151905038806139fd565b601f19821690613a3184600052602060002090565b9160005b818110613a6d57509583600195969710613a54575b505050811b019055565b015160001960f88460031b161c19169055388080613a4a565b9192602060018192868b015181550194019201613a35565b613ae0613ab461043b9597969467ffffffffffffffff60a095168452610100602085015261010084019061074f565b9660408301906001600160801b0360408092805115158552826020820151166020860152015116910152565b01906001600160801b0360408092805115158552826020820151166020860152015116910152565b6040517f23b872dd0000000000000000000000000000000000000000000000000000000060208201526001600160a01b039283166024820152929091166044830152606482019290925261043b91613b6d82608481015b03601f1981018452836106d4565b614659565b9081604091031261041957613ba2602060405192613b8f8461069c565b8035613b9a8161041e565b845201613553565b602082015290565b6020818303126104195780359067ffffffffffffffff821161041957016040818303126104195760405191613bde8361069c565b813567ffffffffffffffff81116104195781613bfb918401610ade565b8352602082013567ffffffffffffffff811161041957613ba29201610ade565b908160209103126104195760405190613c3382610680565b51815290565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561041957016020813591019167ffffffffffffffff821161041957813603831361041957565b906107a19160208152613d6c613d61613d42613cbb613ca88680613c39565b6101006020880152610120870191613283565b613cdb613cca60208801610430565b67ffffffffffffffff166040870152565b613cfa613cea6040880161085c565b6001600160a01b03166060870152565b60608601356080860152613d23613d136080880161085c565b6001600160a01b031660a0870152565b613d3060a0870187613c39565b90601f198784030160c0880152613283565b613d4f60c0860186613c39565b90601f198684030160e0870152613283565b9260e0810190613c39565b91610100601f1982860301910152613283565b9081602091031261041957516107a181611fc4565b9091613dab6107a19360408452604084019061074f565b91602081840391015261074f565b356107a18161084b565b613dcb612992565b50606081013590613ddc8282614708565b613df4613dec60c08301836129a5565b810190613b72565b613e17613e0f613e0760e08501856129a5565b810190613baa565b9182516148fc565b613e2f610ddb6074835101516001600160a01b031690565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690818015159182614060575b5050613fe0575060208181613ead93519101519060405193849283927f57ecfd2800000000000000000000000000000000000000000000000000000000845260048401613d94565b038160006001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165af19081156115ea57600091613fb1575b5015613f87577ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc067ffffffffffffffff613f346040613f2d60208601612d18565b9401613db9565b604080517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03908116825233602083015290921690820152606081018590529216918060808101612bbe565b7fbf969f220000000000000000000000000000000000000000000000000000000060005260046000fd5b613fd3915060203d602011613fd9575b613fcb81836106d4565b810190613d7f565b38613eec565b503d613fc1565b6000935061402191506020926040519485809481937f3907753700000000000000000000000000000000000000000000000000000000835260048301613c89565b03925af19081156115ea57600091614037575090565b6107a1915060203d602011614059575b61405181836106d4565b810190613c1b565b503d614047565b6001600160a01b03161490508138613e65565b6001600160a01b0360015416330361408757565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b60409067ffffffffffffffff6107a19493168152816020820152019061074f565b90805115611e0e578051602082012067ffffffffffffffff831692836000526007602052614107826005604060002001614bf6565b1561416057508161414f7f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9361414a61415b946000526008602052604060002090565b6139b8565b60405191829182610790565b0390a2565b9050611b566040519283927f393b8ad2000000000000000000000000000000000000000000000000000000008452600484016140b1565b67ffffffffffffffff166000818152600660205260409020549092919015614287579161428460e092614259856141ee7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b976144cc565b846000526007602052614205816040600020614c52565b61420e836144cc565b846000526007602052614228836002604060002001614c52565b60405194855260208501906001600160801b0360408092805115158552826020820151166020860152015116910152565b60808301906001600160801b0360408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b90816020910312610419573590565b9081602091031261041957516107a18161041e565b611c7e614392916142e8612d7b565b506142f281614e79565b6020810190606001356143098235612aee8161041e565b614314828254612985565b90557ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff61434984612d18565b604080517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681523360208201529081019490945216918060608101612fc6565b6040516130548161304660208201907ffa7c07de00000000000000000000000000000000000000000000000000000000602083019252565b6143d26132d6565b506001600160801b036060820151166001600160801b0382511690602083019163ffffffff835116420342811161229c5761441b906001600160801b036080870151169061369e565b810180911161229c576144386001600160801b0392918392615416565b161682524263ffffffff16905290565b6040517fa9059cbb0000000000000000000000000000000000000000000000000000000060208201526001600160a01b039092166024830152604482019290925261043b91613b6d8260648101613b5f565b61043b9092919260608101936001600160801b0360408092805115158552826020820151166020860152015116910152565b80511561454c5760408101516001600160801b03166001600160801b0361450c61450060208501516001600160801b031690565b6001600160801b031690565b9116116145165750565b611b56906040519182917f8020d1240000000000000000000000000000000000000000000000000000000083526004830161449a565b6001600160801b0361456860408301516001600160801b031690565b16158015906145af575b6145795750565b611b56906040519182917fd68af9cc0000000000000000000000000000000000000000000000000000000083526004830161449a565b506145c761450060208301516001600160801b031690565b1515614572565b156145d557565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b906001600160a01b036146da921660409060008083519461467a85876106d4565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602087015260208151910182855af1903d156146ff573d6146cb6146c282610733565b945194856106d4565b83523d6000602085013e61581f565b8051806146e5575050565b816020806146fa9361043b9501019101613d7f565b6145ce565b6060925061581f565b6080810161471b611a9c61088983613db9565b6148bb5750602081019061479e60206147526147396104d786612d18565b60801b6fffffffffffffffffffffffffffffffff191690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081526fffffffffffffffffffffffffffffffff19909116600482015291829081906024820190565b03816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156115ea5760009161489c575b50614872576147f16147ec83612d18565b615268565b6147fa82612d18565b9061481a611a9c60a0830193610b3c61481386866129a5565b3691610aa7565b6148325750509061482d61043b92612d18565b61537f565b61483c92506129a5565b90611b566040519283927f24eb47e5000000000000000000000000000000000000000000000000000000008452600484016132c5565b7f53ad11d80000000000000000000000000000000000000000000000000000000060005260046000fd5b6148b5915060203d602011613fd957613fcb81836106d4565b386147db565b6148c761059f91613db9565b7f961c9a4f000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b90815160748110614a925750600482015163ffffffff8116614a5f57506008820151916014600c8201519101519261493b602084015163ffffffff1690565b63ffffffff811663ffffffff831603614a265750507f000000000000000000000000000000000000000000000000000000000000000063ffffffff811663ffffffff8316036149ed5750505167ffffffffffffffff1667ffffffffffffffff811667ffffffffffffffff8316036149b0575050565b7ff917ffea0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff9081166004521660245260446000fd5b7f77e480260000000000000000000000000000000000000000000000000000000060005263ffffffff9081166004521660245260446000fd5b7fe366a1170000000000000000000000000000000000000000000000000000000060005263ffffffff9081166004521660245260446000fd5b7f68d2f8d60000000000000000000000000000000000000000000000000000000060005263ffffffff1660045260246000fd5b7f758b22cc0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b8054821015612d135760005260206000200190600090565b600081815260036020526040902054614b44576002546801000000000000000081101561065f57614b2b614b148260018594016002556002614abf565b81939154906000199060031b92831b921b19161790565b9055600254906000526003602052604060002055600190565b50600090565b600081815260106020526040902054614b4457600f546801000000000000000081101561065f57614b87614b14826001859401600f55600f614abf565b9055600f54906000526010602052604060002055600190565b600081815260066020526040902054614b44576005546801000000000000000081101561065f57614bdd614b148260018594016005556005614abf565b9055600554906000526006602052604060002055600190565b6000828152600182016020526040902054614c4b578054906801000000000000000082101561065f5782614c34614b14846001809601855584614abf565b905580549260005201602052604060002055600190565b5050600090565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1991614dc56133b9928054614ca3614c9d614c948363ffffffff9060801c1690565b63ffffffff1690565b42612d52565b9081614dd1575b5050614d976001614cc560208601516001600160801b031690565b92614d1d614cf86145006001600160801b03614ce885546001600160801b031690565b166001600160801b038816615416565b82906001600160801b03166fffffffffffffffffffffffffffffffff19825416179055565b614d70614d2a8751151590565b82547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178255565b019182906001600160801b03166fffffffffffffffffffffffffffffffff19825416179055565b604083015181546001600160801b031660809190911b6fffffffffffffffffffffffffffffffff1916179055565b6040519182918261449a565b614500614cf8916001600160801b03614e2a614e319582614e2360018a01549282614e1c614e15614e08876001600160801b031690565b996001600160801b031690565b9560801c90565b169061369e565b9116612985565b9116615416565b80547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff00000000000000000000000000000000161781553880614caa565b60808101614e8c611a9c61088983613db9565b6148bb57506020810190614eaa60206147526147396104d786612d18565b03816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156115ea57600091614f23575b50614872576060614f1a61043b93614f09614f0460408601613db9565b615428565b6110a2614f1582612d18565b6154b2565b91013590615576565b614f3c915060203d602011613fd957613fcb81836106d4565b38614ee7565b604051906002548083528260208101600260005260206000209260005b818110614f7457505061043b925003836106d4565b8454835260019485019487945060209093019201614f5f565b604051906005548083528260208101600560005260206000209260005b818110614fbf57505061043b925003836106d4565b8454835260019485019487945060209093019201614faa565b906040519182815491828252602082019060005260206000209260005b81811061500a57505061043b925003836106d4565b8454835260019485019487945060209093019201614ff5565b8054801561504b57600019019061503a8282614abf565b60001982549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600081815260036020526040902054908115614c4b5760001982019082821161229c5760025492600019840193841161229c5783836000956150da95036150e0575b5050506150c96002615023565b600390600052602052604060002090565b55600190565b6150c961510c916151026150f8615112956002614abf565b90549060031b1c90565b9283916002614abf565b90612c0b565b553880806150bc565b600081815260066020526040902054908115614c4b5760001982019082821161229c5760055492600019840193841161229c5783836000956150da950361517b575b50505061516a6005615023565b600690600052602052604060002090565b61516a61510c916151936150f861519d956005614abf565b9283916005614abf565b5538808061515d565b600181019180600052826020526040600020549283151560001461524257600019840184811161229c57835493600019850194851161229c5760009585836150da976151fa9503615209575b505050615023565b90600052602052604060002090565b61522961510c916152206150f86152399588614abf565b92839187614abf565b8590600052602052604060002090565b553880806151f2565b50505050600090565b92615256919261369e565b810180911161229c576107a191615416565b615274611a9c82612d22565b6153485760206152e091615293610ddb6004546001600160a01b031690565b6040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90921660048301523360248301529092839190829081906044820190565b03915afa9081156115ea57600091615329575b50156152fb57565b7f728fe07b000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b615342915060203d602011613fd957613fcb81836106d4565b386152f3565b7fa9902c7e0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b67ffffffffffffffff7f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c911691826000526007602052806153f360026040600020016001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391615620565b604080516001600160a01b0390921682526020820192909252908190810161415b565b9080821015615423575090565b905090565b7f00000000000000000000000000000000000000000000000000000000000000006154505750565b6001600160a01b0316806000526003602052604060002054156154705750565b7fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b9081602091031261041957516107a18161084b565b6154be611a9c82612d22565b615348576020615522916154dd610ddb6004546001600160a01b031690565b60405180809581947fa8d87a3b0000000000000000000000000000000000000000000000000000000083526004830191909167ffffffffffffffff6020820193169052565b03915afa80156115ea576001600160a01b0391600091615547575b501633036152fb57565b615569915060203d60201161556f575b61556181836106d4565b81019061549d565b3861553d565b503d615557565b67ffffffffffffffff7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944911691826000526007602052806153f360406000206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391615620565b81156155f1570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b8054939290919060ff60a086901c16158015615817575b6158105761564d6001600160801b038616614500565b906001840195865461567e614c9d614c94615671614500856001600160801b031690565b9460801c63ffffffff1690565b8061577c575b505083811061573e57508282106156cc575061043b9394506156a99161450091612d52565b6001600160801b03166fffffffffffffffffffffffffffffffff19825416179055565b9061570361059f936156fe6156ef846156e96145008c5460801c90565b93612d52565b6156f883612d43565b90612985565b6155e7565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024526001600160a01b0316604452606490565b7f1a76572a0000000000000000000000000000000000000000000000000000000060005260045260248390526001600160a01b031660445260646000fd5b8285929395116157e65761579661450061579d9460801c90565b918561524b565b84547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178555913880615684565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050509050565b508115615637565b9192901561589a5750815115615833575090565b3b1561583c5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156158ad5750805190602001fd5b611b56906040519182917f08c379a00000000000000000000000000000000000000000000000000000000083526004830161079056fea164736f6c634300081a000a",
}

var HybridLockReleaseUSDCTokenPoolABI = HybridLockReleaseUSDCTokenPoolMetaData.ABI

var HybridLockReleaseUSDCTokenPoolBin = HybridLockReleaseUSDCTokenPoolMetaData.Bin

func DeployHybridLockReleaseUSDCTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, tokenMessenger common.Address, cctpMessageTransmitterProxy common.Address, token common.Address, allowlist []common.Address, rmnProxy common.Address, router common.Address, previousPool common.Address) (common.Address, *types.Transaction, *HybridLockReleaseUSDCTokenPool, error) {
	parsed, err := HybridLockReleaseUSDCTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HybridLockReleaseUSDCTokenPoolBin), backend, tokenMessenger, cctpMessageTransmitterProxy, token, allowlist, rmnProxy, router, previousPool)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HybridLockReleaseUSDCTokenPool{address: address, abi: *parsed, HybridLockReleaseUSDCTokenPoolCaller: HybridLockReleaseUSDCTokenPoolCaller{contract: contract}, HybridLockReleaseUSDCTokenPoolTransactor: HybridLockReleaseUSDCTokenPoolTransactor{contract: contract}, HybridLockReleaseUSDCTokenPoolFilterer: HybridLockReleaseUSDCTokenPoolFilterer{contract: contract}}, nil
}

type HybridLockReleaseUSDCTokenPool struct {
	address common.Address
	abi     abi.ABI
	HybridLockReleaseUSDCTokenPoolCaller
	HybridLockReleaseUSDCTokenPoolTransactor
	HybridLockReleaseUSDCTokenPoolFilterer
}

type HybridLockReleaseUSDCTokenPoolCaller struct {
	contract *bind.BoundContract
}

type HybridLockReleaseUSDCTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type HybridLockReleaseUSDCTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type HybridLockReleaseUSDCTokenPoolSession struct {
	Contract     *HybridLockReleaseUSDCTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type HybridLockReleaseUSDCTokenPoolCallerSession struct {
	Contract *HybridLockReleaseUSDCTokenPoolCaller
	CallOpts bind.CallOpts
}

type HybridLockReleaseUSDCTokenPoolTransactorSession struct {
	Contract     *HybridLockReleaseUSDCTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type HybridLockReleaseUSDCTokenPoolRaw struct {
	Contract *HybridLockReleaseUSDCTokenPool
}

type HybridLockReleaseUSDCTokenPoolCallerRaw struct {
	Contract *HybridLockReleaseUSDCTokenPoolCaller
}

type HybridLockReleaseUSDCTokenPoolTransactorRaw struct {
	Contract *HybridLockReleaseUSDCTokenPoolTransactor
}

func NewHybridLockReleaseUSDCTokenPool(address common.Address, backend bind.ContractBackend) (*HybridLockReleaseUSDCTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(HybridLockReleaseUSDCTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindHybridLockReleaseUSDCTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPool{address: address, abi: abi, HybridLockReleaseUSDCTokenPoolCaller: HybridLockReleaseUSDCTokenPoolCaller{contract: contract}, HybridLockReleaseUSDCTokenPoolTransactor: HybridLockReleaseUSDCTokenPoolTransactor{contract: contract}, HybridLockReleaseUSDCTokenPoolFilterer: HybridLockReleaseUSDCTokenPoolFilterer{contract: contract}}, nil
}

func NewHybridLockReleaseUSDCTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*HybridLockReleaseUSDCTokenPoolCaller, error) {
	contract, err := bindHybridLockReleaseUSDCTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolCaller{contract: contract}, nil
}

func NewHybridLockReleaseUSDCTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*HybridLockReleaseUSDCTokenPoolTransactor, error) {
	contract, err := bindHybridLockReleaseUSDCTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolTransactor{contract: contract}, nil
}

func NewHybridLockReleaseUSDCTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*HybridLockReleaseUSDCTokenPoolFilterer, error) {
	contract, err := bindHybridLockReleaseUSDCTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolFilterer{contract: contract}, nil
}

func bindHybridLockReleaseUSDCTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HybridLockReleaseUSDCTokenPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HybridLockReleaseUSDCTokenPool.Contract.HybridLockReleaseUSDCTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.HybridLockReleaseUSDCTokenPoolTransactor.contract.Transfer(opts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.HybridLockReleaseUSDCTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HybridLockReleaseUSDCTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.contract.Transfer(opts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) SUPPORTEDUSDCVERSION(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "SUPPORTED_USDC_VERSION")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) SUPPORTEDUSDCVERSION() (uint32, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.SUPPORTEDUSDCVERSION(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) SUPPORTEDUSDCVERSION() (uint32, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.SUPPORTEDUSDCVERSION(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) GetAllowList() ([]common.Address, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetAllowList(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) GetAllowList() ([]common.Address, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetAllowList(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) GetAllowListEnabled() (bool, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetAllowListEnabled(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) GetAllowListEnabled() (bool, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetAllowListEnabled(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "getCurrentInboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetCurrentInboundRateLimiterState(&_HybridLockReleaseUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetCurrentInboundRateLimiterState(&_HybridLockReleaseUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "getCurrentOutboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_HybridLockReleaseUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_HybridLockReleaseUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) GetCurrentProposedCCTPChainMigration(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "getCurrentProposedCCTPChainMigration")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) GetCurrentProposedCCTPChainMigration() (uint64, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetCurrentProposedCCTPChainMigration(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) GetCurrentProposedCCTPChainMigration() (uint64, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetCurrentProposedCCTPChainMigration(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) GetDomain(opts *bind.CallOpts, chainSelector uint64) (USDCTokenPoolDomain, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "getDomain", chainSelector)

	if err != nil {
		return *new(USDCTokenPoolDomain), err
	}

	out0 := *abi.ConvertType(out[0], new(USDCTokenPoolDomain)).(*USDCTokenPoolDomain)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) GetDomain(chainSelector uint64) (USDCTokenPoolDomain, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetDomain(&_HybridLockReleaseUSDCTokenPool.CallOpts, chainSelector)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) GetDomain(chainSelector uint64) (USDCTokenPoolDomain, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetDomain(&_HybridLockReleaseUSDCTokenPool.CallOpts, chainSelector)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) GetExcludedTokensByChain(opts *bind.CallOpts, remoteChainSelector uint64) (*big.Int, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "getExcludedTokensByChain", remoteChainSelector)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) GetExcludedTokensByChain(remoteChainSelector uint64) (*big.Int, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetExcludedTokensByChain(&_HybridLockReleaseUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) GetExcludedTokensByChain(remoteChainSelector uint64) (*big.Int, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetExcludedTokensByChain(&_HybridLockReleaseUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) GetLiquidityProvider(opts *bind.CallOpts, remoteChainSelector uint64) (common.Address, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "getLiquidityProvider", remoteChainSelector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) GetLiquidityProvider(remoteChainSelector uint64) (common.Address, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetLiquidityProvider(&_HybridLockReleaseUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) GetLiquidityProvider(remoteChainSelector uint64) (common.Address, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetLiquidityProvider(&_HybridLockReleaseUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) GetLockedTokensForChain(opts *bind.CallOpts, remoteChainSelector uint64) (*big.Int, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "getLockedTokensForChain", remoteChainSelector)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) GetLockedTokensForChain(remoteChainSelector uint64) (*big.Int, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetLockedTokensForChain(&_HybridLockReleaseUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) GetLockedTokensForChain(remoteChainSelector uint64) (*big.Int, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetLockedTokensForChain(&_HybridLockReleaseUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "getRateLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) GetRateLimitAdmin() (common.Address, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetRateLimitAdmin(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) GetRateLimitAdmin() (common.Address, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetRateLimitAdmin(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "getRemotePools", remoteChainSelector)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetRemotePools(&_HybridLockReleaseUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetRemotePools(&_HybridLockReleaseUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "getRemoteToken", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetRemoteToken(&_HybridLockReleaseUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetRemoteToken(&_HybridLockReleaseUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) GetRmnProxy() (common.Address, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetRmnProxy(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) GetRmnProxy() (common.Address, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetRmnProxy(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) GetRouter() (common.Address, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetRouter(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) GetRouter() (common.Address, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetRouter(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) GetSupportedChains() ([]uint64, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetSupportedChains(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) GetSupportedChains() ([]uint64, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetSupportedChains(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) GetToken() (common.Address, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetToken(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetToken(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) GetTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "getTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) GetTokenDecimals() (uint8, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetTokenDecimals(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) GetTokenDecimals() (uint8, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.GetTokenDecimals(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) ILocalDomainIdentifier(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "i_localDomainIdentifier")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) ILocalDomainIdentifier() (uint32, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.ILocalDomainIdentifier(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) ILocalDomainIdentifier() (uint32, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.ILocalDomainIdentifier(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) IMessageTransmitterProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "i_messageTransmitterProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) IMessageTransmitterProxy() (common.Address, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.IMessageTransmitterProxy(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) IMessageTransmitterProxy() (common.Address, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.IMessageTransmitterProxy(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) IPreviousPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "i_previousPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) IPreviousPool() (common.Address, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.IPreviousPool(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) IPreviousPool() (common.Address, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.IPreviousPool(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) ITokenMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "i_tokenMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) ITokenMessenger() (common.Address, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.ITokenMessenger(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) ITokenMessenger() (common.Address, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.ITokenMessenger(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "isRemotePool", remoteChainSelector, remotePoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.IsRemotePool(&_HybridLockReleaseUSDCTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.IsRemotePool(&_HybridLockReleaseUSDCTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.IsSupportedChain(&_HybridLockReleaseUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.IsSupportedChain(&_HybridLockReleaseUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) IsSupportedToken(token common.Address) (bool, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.IsSupportedToken(&_HybridLockReleaseUSDCTokenPool.CallOpts, token)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.IsSupportedToken(&_HybridLockReleaseUSDCTokenPool.CallOpts, token)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) Owner() (common.Address, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.Owner(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) Owner() (common.Address, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.Owner(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) ShouldUseLockRelease(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "shouldUseLockRelease", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) ShouldUseLockRelease(remoteChainSelector uint64) (bool, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.ShouldUseLockRelease(&_HybridLockReleaseUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) ShouldUseLockRelease(remoteChainSelector uint64) (bool, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.ShouldUseLockRelease(&_HybridLockReleaseUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.SupportsInterface(&_HybridLockReleaseUSDCTokenPool.CallOpts, interfaceId)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.SupportsInterface(&_HybridLockReleaseUSDCTokenPool.CallOpts, interfaceId)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) TypeAndVersion() (string, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.TypeAndVersion(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) TypeAndVersion() (string, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.TypeAndVersion(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.AcceptOwnership(&_HybridLockReleaseUSDCTokenPool.TransactOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.AcceptOwnership(&_HybridLockReleaseUSDCTokenPool.TransactOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactor) AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.contract.Transact(opts, "addRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.AddRemotePool(&_HybridLockReleaseUSDCTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactorSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.AddRemotePool(&_HybridLockReleaseUSDCTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.ApplyAllowListUpdates(&_HybridLockReleaseUSDCTokenPool.TransactOpts, removes, adds)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.ApplyAllowListUpdates(&_HybridLockReleaseUSDCTokenPool.TransactOpts, removes, adds)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactor) ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.contract.Transact(opts, "applyChainUpdates", remoteChainSelectorsToRemove, chainsToAdd)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.ApplyChainUpdates(&_HybridLockReleaseUSDCTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactorSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.ApplyChainUpdates(&_HybridLockReleaseUSDCTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactor) BurnLockedUSDC(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.contract.Transact(opts, "burnLockedUSDC")
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) BurnLockedUSDC() (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.BurnLockedUSDC(&_HybridLockReleaseUSDCTokenPool.TransactOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactorSession) BurnLockedUSDC() (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.BurnLockedUSDC(&_HybridLockReleaseUSDCTokenPool.TransactOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactor) CancelExistingCCTPMigrationProposal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.contract.Transact(opts, "cancelExistingCCTPMigrationProposal")
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) CancelExistingCCTPMigrationProposal() (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.CancelExistingCCTPMigrationProposal(&_HybridLockReleaseUSDCTokenPool.TransactOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactorSession) CancelExistingCCTPMigrationProposal() (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.CancelExistingCCTPMigrationProposal(&_HybridLockReleaseUSDCTokenPool.TransactOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactor) ExcludeTokensFromBurn(opts *bind.TransactOpts, remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.contract.Transact(opts, "excludeTokensFromBurn", remoteChainSelector, amount)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) ExcludeTokensFromBurn(remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.ExcludeTokensFromBurn(&_HybridLockReleaseUSDCTokenPool.TransactOpts, remoteChainSelector, amount)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactorSession) ExcludeTokensFromBurn(remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.ExcludeTokensFromBurn(&_HybridLockReleaseUSDCTokenPool.TransactOpts, remoteChainSelector, amount)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.LockOrBurn(&_HybridLockReleaseUSDCTokenPool.TransactOpts, lockOrBurnIn)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.LockOrBurn(&_HybridLockReleaseUSDCTokenPool.TransactOpts, lockOrBurnIn)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactor) ProposeCCTPMigration(opts *bind.TransactOpts, remoteChainSelector uint64) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.contract.Transact(opts, "proposeCCTPMigration", remoteChainSelector)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) ProposeCCTPMigration(remoteChainSelector uint64) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.ProposeCCTPMigration(&_HybridLockReleaseUSDCTokenPool.TransactOpts, remoteChainSelector)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactorSession) ProposeCCTPMigration(remoteChainSelector uint64) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.ProposeCCTPMigration(&_HybridLockReleaseUSDCTokenPool.TransactOpts, remoteChainSelector)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactor) ProvideLiquidity(opts *bind.TransactOpts, remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.contract.Transact(opts, "provideLiquidity", remoteChainSelector, amount)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) ProvideLiquidity(remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.ProvideLiquidity(&_HybridLockReleaseUSDCTokenPool.TransactOpts, remoteChainSelector, amount)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactorSession) ProvideLiquidity(remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.ProvideLiquidity(&_HybridLockReleaseUSDCTokenPool.TransactOpts, remoteChainSelector, amount)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.ReleaseOrMint(&_HybridLockReleaseUSDCTokenPool.TransactOpts, releaseOrMintIn)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.ReleaseOrMint(&_HybridLockReleaseUSDCTokenPool.TransactOpts, releaseOrMintIn)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactor) RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.contract.Transact(opts, "removeRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.RemoveRemotePool(&_HybridLockReleaseUSDCTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactorSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.RemoveRemotePool(&_HybridLockReleaseUSDCTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactor) SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.contract.Transact(opts, "setChainRateLimiterConfig", remoteChainSelector, outboundConfig, inboundConfig)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.SetChainRateLimiterConfig(&_HybridLockReleaseUSDCTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactorSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.SetChainRateLimiterConfig(&_HybridLockReleaseUSDCTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactor) SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.contract.Transact(opts, "setChainRateLimiterConfigs", remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.SetChainRateLimiterConfigs(&_HybridLockReleaseUSDCTokenPool.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactorSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.SetChainRateLimiterConfigs(&_HybridLockReleaseUSDCTokenPool.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactor) SetCircleMigratorAddress(opts *bind.TransactOpts, migrator common.Address) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.contract.Transact(opts, "setCircleMigratorAddress", migrator)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) SetCircleMigratorAddress(migrator common.Address) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.SetCircleMigratorAddress(&_HybridLockReleaseUSDCTokenPool.TransactOpts, migrator)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactorSession) SetCircleMigratorAddress(migrator common.Address) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.SetCircleMigratorAddress(&_HybridLockReleaseUSDCTokenPool.TransactOpts, migrator)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactor) SetDomains(opts *bind.TransactOpts, domains []USDCTokenPoolDomainUpdate) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.contract.Transact(opts, "setDomains", domains)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) SetDomains(domains []USDCTokenPoolDomainUpdate) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.SetDomains(&_HybridLockReleaseUSDCTokenPool.TransactOpts, domains)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactorSession) SetDomains(domains []USDCTokenPoolDomainUpdate) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.SetDomains(&_HybridLockReleaseUSDCTokenPool.TransactOpts, domains)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactor) SetLiquidityProvider(opts *bind.TransactOpts, remoteChainSelector uint64, liquidityProvider common.Address) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.contract.Transact(opts, "setLiquidityProvider", remoteChainSelector, liquidityProvider)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) SetLiquidityProvider(remoteChainSelector uint64, liquidityProvider common.Address) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.SetLiquidityProvider(&_HybridLockReleaseUSDCTokenPool.TransactOpts, remoteChainSelector, liquidityProvider)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactorSession) SetLiquidityProvider(remoteChainSelector uint64, liquidityProvider common.Address) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.SetLiquidityProvider(&_HybridLockReleaseUSDCTokenPool.TransactOpts, remoteChainSelector, liquidityProvider)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactor) SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.contract.Transact(opts, "setRateLimitAdmin", rateLimitAdmin)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.SetRateLimitAdmin(&_HybridLockReleaseUSDCTokenPool.TransactOpts, rateLimitAdmin)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactorSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.SetRateLimitAdmin(&_HybridLockReleaseUSDCTokenPool.TransactOpts, rateLimitAdmin)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactor) SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.contract.Transact(opts, "setRouter", newRouter)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.SetRouter(&_HybridLockReleaseUSDCTokenPool.TransactOpts, newRouter)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactorSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.SetRouter(&_HybridLockReleaseUSDCTokenPool.TransactOpts, newRouter)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.TransferOwnership(&_HybridLockReleaseUSDCTokenPool.TransactOpts, to)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.TransferOwnership(&_HybridLockReleaseUSDCTokenPool.TransactOpts, to)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactor) UpdateChainSelectorMechanisms(opts *bind.TransactOpts, removes []uint64, adds []uint64) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.contract.Transact(opts, "updateChainSelectorMechanisms", removes, adds)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) UpdateChainSelectorMechanisms(removes []uint64, adds []uint64) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.UpdateChainSelectorMechanisms(&_HybridLockReleaseUSDCTokenPool.TransactOpts, removes, adds)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactorSession) UpdateChainSelectorMechanisms(removes []uint64, adds []uint64) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.UpdateChainSelectorMechanisms(&_HybridLockReleaseUSDCTokenPool.TransactOpts, removes, adds)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactor) WithdrawLiquidity(opts *bind.TransactOpts, remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.contract.Transact(opts, "withdrawLiquidity", remoteChainSelector, amount)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) WithdrawLiquidity(remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.WithdrawLiquidity(&_HybridLockReleaseUSDCTokenPool.TransactOpts, remoteChainSelector, amount)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactorSession) WithdrawLiquidity(remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.WithdrawLiquidity(&_HybridLockReleaseUSDCTokenPool.TransactOpts, remoteChainSelector, amount)
}

type HybridLockReleaseUSDCTokenPoolAllowListAddIterator struct {
	Event *HybridLockReleaseUSDCTokenPoolAllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *HybridLockReleaseUSDCTokenPoolAllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridLockReleaseUSDCTokenPoolAllowListAdd)
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
		it.Event = new(HybridLockReleaseUSDCTokenPoolAllowListAdd)
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

func (it *HybridLockReleaseUSDCTokenPoolAllowListAddIterator) Error() error {
	return it.fail
}

func (it *HybridLockReleaseUSDCTokenPoolAllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type HybridLockReleaseUSDCTokenPoolAllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) FilterAllowListAdd(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolAllowListAddIterator, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolAllowListAddIterator{contract: _HybridLockReleaseUSDCTokenPool.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolAllowListAdd) (event.Subscription, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(HybridLockReleaseUSDCTokenPoolAllowListAdd)
				if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) ParseAllowListAdd(log types.Log) (*HybridLockReleaseUSDCTokenPoolAllowListAdd, error) {
	event := new(HybridLockReleaseUSDCTokenPoolAllowListAdd)
	if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type HybridLockReleaseUSDCTokenPoolAllowListRemoveIterator struct {
	Event *HybridLockReleaseUSDCTokenPoolAllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *HybridLockReleaseUSDCTokenPoolAllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridLockReleaseUSDCTokenPoolAllowListRemove)
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
		it.Event = new(HybridLockReleaseUSDCTokenPoolAllowListRemove)
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

func (it *HybridLockReleaseUSDCTokenPoolAllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *HybridLockReleaseUSDCTokenPoolAllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type HybridLockReleaseUSDCTokenPoolAllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) FilterAllowListRemove(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolAllowListRemoveIterator, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolAllowListRemoveIterator{contract: _HybridLockReleaseUSDCTokenPool.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolAllowListRemove) (event.Subscription, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(HybridLockReleaseUSDCTokenPoolAllowListRemove)
				if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) ParseAllowListRemove(log types.Log) (*HybridLockReleaseUSDCTokenPoolAllowListRemove, error) {
	event := new(HybridLockReleaseUSDCTokenPoolAllowListRemove)
	if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type HybridLockReleaseUSDCTokenPoolCCTPMigrationCancelledIterator struct {
	Event *HybridLockReleaseUSDCTokenPoolCCTPMigrationCancelled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *HybridLockReleaseUSDCTokenPoolCCTPMigrationCancelledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridLockReleaseUSDCTokenPoolCCTPMigrationCancelled)
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
		it.Event = new(HybridLockReleaseUSDCTokenPoolCCTPMigrationCancelled)
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

func (it *HybridLockReleaseUSDCTokenPoolCCTPMigrationCancelledIterator) Error() error {
	return it.fail
}

func (it *HybridLockReleaseUSDCTokenPoolCCTPMigrationCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type HybridLockReleaseUSDCTokenPoolCCTPMigrationCancelled struct {
	ExistingProposalSelector uint64
	Raw                      types.Log
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) FilterCCTPMigrationCancelled(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolCCTPMigrationCancelledIterator, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.FilterLogs(opts, "CCTPMigrationCancelled")
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolCCTPMigrationCancelledIterator{contract: _HybridLockReleaseUSDCTokenPool.contract, event: "CCTPMigrationCancelled", logs: logs, sub: sub}, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) WatchCCTPMigrationCancelled(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolCCTPMigrationCancelled) (event.Subscription, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.WatchLogs(opts, "CCTPMigrationCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(HybridLockReleaseUSDCTokenPoolCCTPMigrationCancelled)
				if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "CCTPMigrationCancelled", log); err != nil {
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) ParseCCTPMigrationCancelled(log types.Log) (*HybridLockReleaseUSDCTokenPoolCCTPMigrationCancelled, error) {
	event := new(HybridLockReleaseUSDCTokenPoolCCTPMigrationCancelled)
	if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "CCTPMigrationCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type HybridLockReleaseUSDCTokenPoolCCTPMigrationExecutedIterator struct {
	Event *HybridLockReleaseUSDCTokenPoolCCTPMigrationExecuted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *HybridLockReleaseUSDCTokenPoolCCTPMigrationExecutedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridLockReleaseUSDCTokenPoolCCTPMigrationExecuted)
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
		it.Event = new(HybridLockReleaseUSDCTokenPoolCCTPMigrationExecuted)
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

func (it *HybridLockReleaseUSDCTokenPoolCCTPMigrationExecutedIterator) Error() error {
	return it.fail
}

func (it *HybridLockReleaseUSDCTokenPoolCCTPMigrationExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type HybridLockReleaseUSDCTokenPoolCCTPMigrationExecuted struct {
	RemoteChainSelector uint64
	USDCBurned          *big.Int
	Raw                 types.Log
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) FilterCCTPMigrationExecuted(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolCCTPMigrationExecutedIterator, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.FilterLogs(opts, "CCTPMigrationExecuted")
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolCCTPMigrationExecutedIterator{contract: _HybridLockReleaseUSDCTokenPool.contract, event: "CCTPMigrationExecuted", logs: logs, sub: sub}, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) WatchCCTPMigrationExecuted(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolCCTPMigrationExecuted) (event.Subscription, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.WatchLogs(opts, "CCTPMigrationExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(HybridLockReleaseUSDCTokenPoolCCTPMigrationExecuted)
				if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "CCTPMigrationExecuted", log); err != nil {
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) ParseCCTPMigrationExecuted(log types.Log) (*HybridLockReleaseUSDCTokenPoolCCTPMigrationExecuted, error) {
	event := new(HybridLockReleaseUSDCTokenPoolCCTPMigrationExecuted)
	if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "CCTPMigrationExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type HybridLockReleaseUSDCTokenPoolCCTPMigrationProposedIterator struct {
	Event *HybridLockReleaseUSDCTokenPoolCCTPMigrationProposed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *HybridLockReleaseUSDCTokenPoolCCTPMigrationProposedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridLockReleaseUSDCTokenPoolCCTPMigrationProposed)
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
		it.Event = new(HybridLockReleaseUSDCTokenPoolCCTPMigrationProposed)
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

func (it *HybridLockReleaseUSDCTokenPoolCCTPMigrationProposedIterator) Error() error {
	return it.fail
}

func (it *HybridLockReleaseUSDCTokenPoolCCTPMigrationProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type HybridLockReleaseUSDCTokenPoolCCTPMigrationProposed struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) FilterCCTPMigrationProposed(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolCCTPMigrationProposedIterator, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.FilterLogs(opts, "CCTPMigrationProposed")
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolCCTPMigrationProposedIterator{contract: _HybridLockReleaseUSDCTokenPool.contract, event: "CCTPMigrationProposed", logs: logs, sub: sub}, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) WatchCCTPMigrationProposed(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolCCTPMigrationProposed) (event.Subscription, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.WatchLogs(opts, "CCTPMigrationProposed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(HybridLockReleaseUSDCTokenPoolCCTPMigrationProposed)
				if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "CCTPMigrationProposed", log); err != nil {
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) ParseCCTPMigrationProposed(log types.Log) (*HybridLockReleaseUSDCTokenPoolCCTPMigrationProposed, error) {
	event := new(HybridLockReleaseUSDCTokenPoolCCTPMigrationProposed)
	if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "CCTPMigrationProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type HybridLockReleaseUSDCTokenPoolChainAddedIterator struct {
	Event *HybridLockReleaseUSDCTokenPoolChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *HybridLockReleaseUSDCTokenPoolChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridLockReleaseUSDCTokenPoolChainAdded)
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
		it.Event = new(HybridLockReleaseUSDCTokenPoolChainAdded)
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

func (it *HybridLockReleaseUSDCTokenPoolChainAddedIterator) Error() error {
	return it.fail
}

func (it *HybridLockReleaseUSDCTokenPoolChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type HybridLockReleaseUSDCTokenPoolChainAdded struct {
	RemoteChainSelector       uint64
	RemoteToken               []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) FilterChainAdded(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolChainAddedIterator, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolChainAddedIterator{contract: _HybridLockReleaseUSDCTokenPool.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolChainAdded) (event.Subscription, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(HybridLockReleaseUSDCTokenPoolChainAdded)
				if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) ParseChainAdded(log types.Log) (*HybridLockReleaseUSDCTokenPoolChainAdded, error) {
	event := new(HybridLockReleaseUSDCTokenPoolChainAdded)
	if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type HybridLockReleaseUSDCTokenPoolChainConfiguredIterator struct {
	Event *HybridLockReleaseUSDCTokenPoolChainConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *HybridLockReleaseUSDCTokenPoolChainConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridLockReleaseUSDCTokenPoolChainConfigured)
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
		it.Event = new(HybridLockReleaseUSDCTokenPoolChainConfigured)
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

func (it *HybridLockReleaseUSDCTokenPoolChainConfiguredIterator) Error() error {
	return it.fail
}

func (it *HybridLockReleaseUSDCTokenPoolChainConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type HybridLockReleaseUSDCTokenPoolChainConfigured struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) FilterChainConfigured(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolChainConfiguredIterator, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.FilterLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolChainConfiguredIterator{contract: _HybridLockReleaseUSDCTokenPool.contract, event: "ChainConfigured", logs: logs, sub: sub}, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolChainConfigured) (event.Subscription, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.WatchLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(HybridLockReleaseUSDCTokenPoolChainConfigured)
				if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) ParseChainConfigured(log types.Log) (*HybridLockReleaseUSDCTokenPoolChainConfigured, error) {
	event := new(HybridLockReleaseUSDCTokenPoolChainConfigured)
	if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type HybridLockReleaseUSDCTokenPoolChainRemovedIterator struct {
	Event *HybridLockReleaseUSDCTokenPoolChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *HybridLockReleaseUSDCTokenPoolChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridLockReleaseUSDCTokenPoolChainRemoved)
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
		it.Event = new(HybridLockReleaseUSDCTokenPoolChainRemoved)
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

func (it *HybridLockReleaseUSDCTokenPoolChainRemovedIterator) Error() error {
	return it.fail
}

func (it *HybridLockReleaseUSDCTokenPoolChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type HybridLockReleaseUSDCTokenPoolChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) FilterChainRemoved(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolChainRemovedIterator, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolChainRemovedIterator{contract: _HybridLockReleaseUSDCTokenPool.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolChainRemoved) (event.Subscription, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(HybridLockReleaseUSDCTokenPoolChainRemoved)
				if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) ParseChainRemoved(log types.Log) (*HybridLockReleaseUSDCTokenPoolChainRemoved, error) {
	event := new(HybridLockReleaseUSDCTokenPoolChainRemoved)
	if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type HybridLockReleaseUSDCTokenPoolCircleMigratorAddressSetIterator struct {
	Event *HybridLockReleaseUSDCTokenPoolCircleMigratorAddressSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *HybridLockReleaseUSDCTokenPoolCircleMigratorAddressSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridLockReleaseUSDCTokenPoolCircleMigratorAddressSet)
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
		it.Event = new(HybridLockReleaseUSDCTokenPoolCircleMigratorAddressSet)
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

func (it *HybridLockReleaseUSDCTokenPoolCircleMigratorAddressSetIterator) Error() error {
	return it.fail
}

func (it *HybridLockReleaseUSDCTokenPoolCircleMigratorAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type HybridLockReleaseUSDCTokenPoolCircleMigratorAddressSet struct {
	MigratorAddress common.Address
	Raw             types.Log
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) FilterCircleMigratorAddressSet(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolCircleMigratorAddressSetIterator, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.FilterLogs(opts, "CircleMigratorAddressSet")
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolCircleMigratorAddressSetIterator{contract: _HybridLockReleaseUSDCTokenPool.contract, event: "CircleMigratorAddressSet", logs: logs, sub: sub}, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) WatchCircleMigratorAddressSet(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolCircleMigratorAddressSet) (event.Subscription, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.WatchLogs(opts, "CircleMigratorAddressSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(HybridLockReleaseUSDCTokenPoolCircleMigratorAddressSet)
				if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "CircleMigratorAddressSet", log); err != nil {
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) ParseCircleMigratorAddressSet(log types.Log) (*HybridLockReleaseUSDCTokenPoolCircleMigratorAddressSet, error) {
	event := new(HybridLockReleaseUSDCTokenPoolCircleMigratorAddressSet)
	if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "CircleMigratorAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type HybridLockReleaseUSDCTokenPoolConfigChangedIterator struct {
	Event *HybridLockReleaseUSDCTokenPoolConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *HybridLockReleaseUSDCTokenPoolConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridLockReleaseUSDCTokenPoolConfigChanged)
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
		it.Event = new(HybridLockReleaseUSDCTokenPoolConfigChanged)
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

func (it *HybridLockReleaseUSDCTokenPoolConfigChangedIterator) Error() error {
	return it.fail
}

func (it *HybridLockReleaseUSDCTokenPoolConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type HybridLockReleaseUSDCTokenPoolConfigChanged struct {
	Config RateLimiterConfig
	Raw    types.Log
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolConfigChangedIterator, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolConfigChangedIterator{contract: _HybridLockReleaseUSDCTokenPool.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolConfigChanged) (event.Subscription, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(HybridLockReleaseUSDCTokenPoolConfigChanged)
				if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) ParseConfigChanged(log types.Log) (*HybridLockReleaseUSDCTokenPoolConfigChanged, error) {
	event := new(HybridLockReleaseUSDCTokenPoolConfigChanged)
	if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type HybridLockReleaseUSDCTokenPoolConfigSetIterator struct {
	Event *HybridLockReleaseUSDCTokenPoolConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *HybridLockReleaseUSDCTokenPoolConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridLockReleaseUSDCTokenPoolConfigSet)
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
		it.Event = new(HybridLockReleaseUSDCTokenPoolConfigSet)
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

func (it *HybridLockReleaseUSDCTokenPoolConfigSetIterator) Error() error {
	return it.fail
}

func (it *HybridLockReleaseUSDCTokenPoolConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type HybridLockReleaseUSDCTokenPoolConfigSet struct {
	TokenMessenger common.Address
	Raw            types.Log
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) FilterConfigSet(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolConfigSetIterator, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolConfigSetIterator{contract: _HybridLockReleaseUSDCTokenPool.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolConfigSet) (event.Subscription, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(HybridLockReleaseUSDCTokenPoolConfigSet)
				if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) ParseConfigSet(log types.Log) (*HybridLockReleaseUSDCTokenPoolConfigSet, error) {
	event := new(HybridLockReleaseUSDCTokenPoolConfigSet)
	if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type HybridLockReleaseUSDCTokenPoolDomainsSetIterator struct {
	Event *HybridLockReleaseUSDCTokenPoolDomainsSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *HybridLockReleaseUSDCTokenPoolDomainsSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridLockReleaseUSDCTokenPoolDomainsSet)
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
		it.Event = new(HybridLockReleaseUSDCTokenPoolDomainsSet)
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

func (it *HybridLockReleaseUSDCTokenPoolDomainsSetIterator) Error() error {
	return it.fail
}

func (it *HybridLockReleaseUSDCTokenPoolDomainsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type HybridLockReleaseUSDCTokenPoolDomainsSet struct {
	Arg0 []USDCTokenPoolDomainUpdate
	Raw  types.Log
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) FilterDomainsSet(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolDomainsSetIterator, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.FilterLogs(opts, "DomainsSet")
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolDomainsSetIterator{contract: _HybridLockReleaseUSDCTokenPool.contract, event: "DomainsSet", logs: logs, sub: sub}, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) WatchDomainsSet(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolDomainsSet) (event.Subscription, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.WatchLogs(opts, "DomainsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(HybridLockReleaseUSDCTokenPoolDomainsSet)
				if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "DomainsSet", log); err != nil {
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) ParseDomainsSet(log types.Log) (*HybridLockReleaseUSDCTokenPoolDomainsSet, error) {
	event := new(HybridLockReleaseUSDCTokenPoolDomainsSet)
	if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "DomainsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type HybridLockReleaseUSDCTokenPoolInboundRateLimitConsumedIterator struct {
	Event *HybridLockReleaseUSDCTokenPoolInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *HybridLockReleaseUSDCTokenPoolInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridLockReleaseUSDCTokenPoolInboundRateLimitConsumed)
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
		it.Event = new(HybridLockReleaseUSDCTokenPoolInboundRateLimitConsumed)
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

func (it *HybridLockReleaseUSDCTokenPoolInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *HybridLockReleaseUSDCTokenPoolInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type HybridLockReleaseUSDCTokenPoolInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*HybridLockReleaseUSDCTokenPoolInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.FilterLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolInboundRateLimitConsumedIterator{contract: _HybridLockReleaseUSDCTokenPool.contract, event: "InboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.WatchLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(HybridLockReleaseUSDCTokenPoolInboundRateLimitConsumed)
				if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) ParseInboundRateLimitConsumed(log types.Log) (*HybridLockReleaseUSDCTokenPoolInboundRateLimitConsumed, error) {
	event := new(HybridLockReleaseUSDCTokenPoolInboundRateLimitConsumed)
	if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type HybridLockReleaseUSDCTokenPoolLiquidityAddedIterator struct {
	Event *HybridLockReleaseUSDCTokenPoolLiquidityAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *HybridLockReleaseUSDCTokenPoolLiquidityAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridLockReleaseUSDCTokenPoolLiquidityAdded)
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
		it.Event = new(HybridLockReleaseUSDCTokenPoolLiquidityAdded)
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

func (it *HybridLockReleaseUSDCTokenPoolLiquidityAddedIterator) Error() error {
	return it.fail
}

func (it *HybridLockReleaseUSDCTokenPoolLiquidityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type HybridLockReleaseUSDCTokenPoolLiquidityAdded struct {
	Provider common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) FilterLiquidityAdded(opts *bind.FilterOpts, provider []common.Address, amount []*big.Int) (*HybridLockReleaseUSDCTokenPoolLiquidityAddedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.FilterLogs(opts, "LiquidityAdded", providerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolLiquidityAddedIterator{contract: _HybridLockReleaseUSDCTokenPool.contract, event: "LiquidityAdded", logs: logs, sub: sub}, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolLiquidityAdded, provider []common.Address, amount []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.WatchLogs(opts, "LiquidityAdded", providerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(HybridLockReleaseUSDCTokenPoolLiquidityAdded)
				if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) ParseLiquidityAdded(log types.Log) (*HybridLockReleaseUSDCTokenPoolLiquidityAdded, error) {
	event := new(HybridLockReleaseUSDCTokenPoolLiquidityAdded)
	if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type HybridLockReleaseUSDCTokenPoolLiquidityProviderSetIterator struct {
	Event *HybridLockReleaseUSDCTokenPoolLiquidityProviderSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *HybridLockReleaseUSDCTokenPoolLiquidityProviderSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridLockReleaseUSDCTokenPoolLiquidityProviderSet)
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
		it.Event = new(HybridLockReleaseUSDCTokenPoolLiquidityProviderSet)
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

func (it *HybridLockReleaseUSDCTokenPoolLiquidityProviderSetIterator) Error() error {
	return it.fail
}

func (it *HybridLockReleaseUSDCTokenPoolLiquidityProviderSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type HybridLockReleaseUSDCTokenPoolLiquidityProviderSet struct {
	OldProvider         common.Address
	NewProvider         common.Address
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) FilterLiquidityProviderSet(opts *bind.FilterOpts, oldProvider []common.Address, newProvider []common.Address, remoteChainSelector []uint64) (*HybridLockReleaseUSDCTokenPoolLiquidityProviderSetIterator, error) {

	var oldProviderRule []interface{}
	for _, oldProviderItem := range oldProvider {
		oldProviderRule = append(oldProviderRule, oldProviderItem)
	}
	var newProviderRule []interface{}
	for _, newProviderItem := range newProvider {
		newProviderRule = append(newProviderRule, newProviderItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.FilterLogs(opts, "LiquidityProviderSet", oldProviderRule, newProviderRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolLiquidityProviderSetIterator{contract: _HybridLockReleaseUSDCTokenPool.contract, event: "LiquidityProviderSet", logs: logs, sub: sub}, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) WatchLiquidityProviderSet(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolLiquidityProviderSet, oldProvider []common.Address, newProvider []common.Address, remoteChainSelector []uint64) (event.Subscription, error) {

	var oldProviderRule []interface{}
	for _, oldProviderItem := range oldProvider {
		oldProviderRule = append(oldProviderRule, oldProviderItem)
	}
	var newProviderRule []interface{}
	for _, newProviderItem := range newProvider {
		newProviderRule = append(newProviderRule, newProviderItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.WatchLogs(opts, "LiquidityProviderSet", oldProviderRule, newProviderRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(HybridLockReleaseUSDCTokenPoolLiquidityProviderSet)
				if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "LiquidityProviderSet", log); err != nil {
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) ParseLiquidityProviderSet(log types.Log) (*HybridLockReleaseUSDCTokenPoolLiquidityProviderSet, error) {
	event := new(HybridLockReleaseUSDCTokenPoolLiquidityProviderSet)
	if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "LiquidityProviderSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type HybridLockReleaseUSDCTokenPoolLiquidityRemovedIterator struct {
	Event *HybridLockReleaseUSDCTokenPoolLiquidityRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *HybridLockReleaseUSDCTokenPoolLiquidityRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridLockReleaseUSDCTokenPoolLiquidityRemoved)
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
		it.Event = new(HybridLockReleaseUSDCTokenPoolLiquidityRemoved)
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

func (it *HybridLockReleaseUSDCTokenPoolLiquidityRemovedIterator) Error() error {
	return it.fail
}

func (it *HybridLockReleaseUSDCTokenPoolLiquidityRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type HybridLockReleaseUSDCTokenPoolLiquidityRemoved struct {
	Provider common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) FilterLiquidityRemoved(opts *bind.FilterOpts, provider []common.Address, amount []*big.Int) (*HybridLockReleaseUSDCTokenPoolLiquidityRemovedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.FilterLogs(opts, "LiquidityRemoved", providerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolLiquidityRemovedIterator{contract: _HybridLockReleaseUSDCTokenPool.contract, event: "LiquidityRemoved", logs: logs, sub: sub}, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolLiquidityRemoved, provider []common.Address, amount []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.WatchLogs(opts, "LiquidityRemoved", providerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(HybridLockReleaseUSDCTokenPoolLiquidityRemoved)
				if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) ParseLiquidityRemoved(log types.Log) (*HybridLockReleaseUSDCTokenPoolLiquidityRemoved, error) {
	event := new(HybridLockReleaseUSDCTokenPoolLiquidityRemoved)
	if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type HybridLockReleaseUSDCTokenPoolLiquidityTransferredIterator struct {
	Event *HybridLockReleaseUSDCTokenPoolLiquidityTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *HybridLockReleaseUSDCTokenPoolLiquidityTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridLockReleaseUSDCTokenPoolLiquidityTransferred)
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
		it.Event = new(HybridLockReleaseUSDCTokenPoolLiquidityTransferred)
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

func (it *HybridLockReleaseUSDCTokenPoolLiquidityTransferredIterator) Error() error {
	return it.fail
}

func (it *HybridLockReleaseUSDCTokenPoolLiquidityTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type HybridLockReleaseUSDCTokenPoolLiquidityTransferred struct {
	From                common.Address
	RemoteChainSelector uint64
	Amount              *big.Int
	Raw                 types.Log
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) FilterLiquidityTransferred(opts *bind.FilterOpts, from []common.Address, remoteChainSelector []uint64) (*HybridLockReleaseUSDCTokenPoolLiquidityTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.FilterLogs(opts, "LiquidityTransferred", fromRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolLiquidityTransferredIterator{contract: _HybridLockReleaseUSDCTokenPool.contract, event: "LiquidityTransferred", logs: logs, sub: sub}, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) WatchLiquidityTransferred(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolLiquidityTransferred, from []common.Address, remoteChainSelector []uint64) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.WatchLogs(opts, "LiquidityTransferred", fromRule, remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(HybridLockReleaseUSDCTokenPoolLiquidityTransferred)
				if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "LiquidityTransferred", log); err != nil {
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) ParseLiquidityTransferred(log types.Log) (*HybridLockReleaseUSDCTokenPoolLiquidityTransferred, error) {
	event := new(HybridLockReleaseUSDCTokenPoolLiquidityTransferred)
	if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "LiquidityTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type HybridLockReleaseUSDCTokenPoolLockReleaseDisabledIterator struct {
	Event *HybridLockReleaseUSDCTokenPoolLockReleaseDisabled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *HybridLockReleaseUSDCTokenPoolLockReleaseDisabledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridLockReleaseUSDCTokenPoolLockReleaseDisabled)
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
		it.Event = new(HybridLockReleaseUSDCTokenPoolLockReleaseDisabled)
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

func (it *HybridLockReleaseUSDCTokenPoolLockReleaseDisabledIterator) Error() error {
	return it.fail
}

func (it *HybridLockReleaseUSDCTokenPoolLockReleaseDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type HybridLockReleaseUSDCTokenPoolLockReleaseDisabled struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) FilterLockReleaseDisabled(opts *bind.FilterOpts, remoteChainSelector []uint64) (*HybridLockReleaseUSDCTokenPoolLockReleaseDisabledIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.FilterLogs(opts, "LockReleaseDisabled", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolLockReleaseDisabledIterator{contract: _HybridLockReleaseUSDCTokenPool.contract, event: "LockReleaseDisabled", logs: logs, sub: sub}, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) WatchLockReleaseDisabled(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolLockReleaseDisabled, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.WatchLogs(opts, "LockReleaseDisabled", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(HybridLockReleaseUSDCTokenPoolLockReleaseDisabled)
				if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "LockReleaseDisabled", log); err != nil {
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) ParseLockReleaseDisabled(log types.Log) (*HybridLockReleaseUSDCTokenPoolLockReleaseDisabled, error) {
	event := new(HybridLockReleaseUSDCTokenPoolLockReleaseDisabled)
	if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "LockReleaseDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type HybridLockReleaseUSDCTokenPoolLockReleaseEnabledIterator struct {
	Event *HybridLockReleaseUSDCTokenPoolLockReleaseEnabled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *HybridLockReleaseUSDCTokenPoolLockReleaseEnabledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridLockReleaseUSDCTokenPoolLockReleaseEnabled)
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
		it.Event = new(HybridLockReleaseUSDCTokenPoolLockReleaseEnabled)
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

func (it *HybridLockReleaseUSDCTokenPoolLockReleaseEnabledIterator) Error() error {
	return it.fail
}

func (it *HybridLockReleaseUSDCTokenPoolLockReleaseEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type HybridLockReleaseUSDCTokenPoolLockReleaseEnabled struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) FilterLockReleaseEnabled(opts *bind.FilterOpts, remoteChainSelector []uint64) (*HybridLockReleaseUSDCTokenPoolLockReleaseEnabledIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.FilterLogs(opts, "LockReleaseEnabled", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolLockReleaseEnabledIterator{contract: _HybridLockReleaseUSDCTokenPool.contract, event: "LockReleaseEnabled", logs: logs, sub: sub}, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) WatchLockReleaseEnabled(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolLockReleaseEnabled, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.WatchLogs(opts, "LockReleaseEnabled", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(HybridLockReleaseUSDCTokenPoolLockReleaseEnabled)
				if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "LockReleaseEnabled", log); err != nil {
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) ParseLockReleaseEnabled(log types.Log) (*HybridLockReleaseUSDCTokenPoolLockReleaseEnabled, error) {
	event := new(HybridLockReleaseUSDCTokenPoolLockReleaseEnabled)
	if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "LockReleaseEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type HybridLockReleaseUSDCTokenPoolLockedOrBurnedIterator struct {
	Event *HybridLockReleaseUSDCTokenPoolLockedOrBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *HybridLockReleaseUSDCTokenPoolLockedOrBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridLockReleaseUSDCTokenPoolLockedOrBurned)
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
		it.Event = new(HybridLockReleaseUSDCTokenPoolLockedOrBurned)
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

func (it *HybridLockReleaseUSDCTokenPoolLockedOrBurnedIterator) Error() error {
	return it.fail
}

func (it *HybridLockReleaseUSDCTokenPoolLockedOrBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type HybridLockReleaseUSDCTokenPoolLockedOrBurned struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*HybridLockReleaseUSDCTokenPoolLockedOrBurnedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.FilterLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolLockedOrBurnedIterator{contract: _HybridLockReleaseUSDCTokenPool.contract, event: "LockedOrBurned", logs: logs, sub: sub}, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.WatchLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(HybridLockReleaseUSDCTokenPoolLockedOrBurned)
				if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) ParseLockedOrBurned(log types.Log) (*HybridLockReleaseUSDCTokenPoolLockedOrBurned, error) {
	event := new(HybridLockReleaseUSDCTokenPoolLockedOrBurned)
	if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type HybridLockReleaseUSDCTokenPoolOutboundRateLimitConsumedIterator struct {
	Event *HybridLockReleaseUSDCTokenPoolOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *HybridLockReleaseUSDCTokenPoolOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridLockReleaseUSDCTokenPoolOutboundRateLimitConsumed)
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
		it.Event = new(HybridLockReleaseUSDCTokenPoolOutboundRateLimitConsumed)
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

func (it *HybridLockReleaseUSDCTokenPoolOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *HybridLockReleaseUSDCTokenPoolOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type HybridLockReleaseUSDCTokenPoolOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*HybridLockReleaseUSDCTokenPoolOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.FilterLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolOutboundRateLimitConsumedIterator{contract: _HybridLockReleaseUSDCTokenPool.contract, event: "OutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.WatchLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(HybridLockReleaseUSDCTokenPoolOutboundRateLimitConsumed)
				if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) ParseOutboundRateLimitConsumed(log types.Log) (*HybridLockReleaseUSDCTokenPoolOutboundRateLimitConsumed, error) {
	event := new(HybridLockReleaseUSDCTokenPoolOutboundRateLimitConsumed)
	if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type HybridLockReleaseUSDCTokenPoolOwnershipTransferRequestedIterator struct {
	Event *HybridLockReleaseUSDCTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *HybridLockReleaseUSDCTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridLockReleaseUSDCTokenPoolOwnershipTransferRequested)
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
		it.Event = new(HybridLockReleaseUSDCTokenPoolOwnershipTransferRequested)
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

func (it *HybridLockReleaseUSDCTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *HybridLockReleaseUSDCTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type HybridLockReleaseUSDCTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*HybridLockReleaseUSDCTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolOwnershipTransferRequestedIterator{contract: _HybridLockReleaseUSDCTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(HybridLockReleaseUSDCTokenPoolOwnershipTransferRequested)
				if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*HybridLockReleaseUSDCTokenPoolOwnershipTransferRequested, error) {
	event := new(HybridLockReleaseUSDCTokenPoolOwnershipTransferRequested)
	if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type HybridLockReleaseUSDCTokenPoolOwnershipTransferredIterator struct {
	Event *HybridLockReleaseUSDCTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *HybridLockReleaseUSDCTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridLockReleaseUSDCTokenPoolOwnershipTransferred)
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
		it.Event = new(HybridLockReleaseUSDCTokenPoolOwnershipTransferred)
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

func (it *HybridLockReleaseUSDCTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *HybridLockReleaseUSDCTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type HybridLockReleaseUSDCTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*HybridLockReleaseUSDCTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolOwnershipTransferredIterator{contract: _HybridLockReleaseUSDCTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(HybridLockReleaseUSDCTokenPoolOwnershipTransferred)
				if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*HybridLockReleaseUSDCTokenPoolOwnershipTransferred, error) {
	event := new(HybridLockReleaseUSDCTokenPoolOwnershipTransferred)
	if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type HybridLockReleaseUSDCTokenPoolRateLimitAdminSetIterator struct {
	Event *HybridLockReleaseUSDCTokenPoolRateLimitAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *HybridLockReleaseUSDCTokenPoolRateLimitAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridLockReleaseUSDCTokenPoolRateLimitAdminSet)
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
		it.Event = new(HybridLockReleaseUSDCTokenPoolRateLimitAdminSet)
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

func (it *HybridLockReleaseUSDCTokenPoolRateLimitAdminSetIterator) Error() error {
	return it.fail
}

func (it *HybridLockReleaseUSDCTokenPoolRateLimitAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type HybridLockReleaseUSDCTokenPoolRateLimitAdminSet struct {
	RateLimitAdmin common.Address
	Raw            types.Log
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) FilterRateLimitAdminSet(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolRateLimitAdminSetIterator, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.FilterLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolRateLimitAdminSetIterator{contract: _HybridLockReleaseUSDCTokenPool.contract, event: "RateLimitAdminSet", logs: logs, sub: sub}, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolRateLimitAdminSet) (event.Subscription, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.WatchLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(HybridLockReleaseUSDCTokenPoolRateLimitAdminSet)
				if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) ParseRateLimitAdminSet(log types.Log) (*HybridLockReleaseUSDCTokenPoolRateLimitAdminSet, error) {
	event := new(HybridLockReleaseUSDCTokenPoolRateLimitAdminSet)
	if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type HybridLockReleaseUSDCTokenPoolReleasedOrMintedIterator struct {
	Event *HybridLockReleaseUSDCTokenPoolReleasedOrMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *HybridLockReleaseUSDCTokenPoolReleasedOrMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridLockReleaseUSDCTokenPoolReleasedOrMinted)
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
		it.Event = new(HybridLockReleaseUSDCTokenPoolReleasedOrMinted)
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

func (it *HybridLockReleaseUSDCTokenPoolReleasedOrMintedIterator) Error() error {
	return it.fail
}

func (it *HybridLockReleaseUSDCTokenPoolReleasedOrMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type HybridLockReleaseUSDCTokenPoolReleasedOrMinted struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Recipient           common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*HybridLockReleaseUSDCTokenPoolReleasedOrMintedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.FilterLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolReleasedOrMintedIterator{contract: _HybridLockReleaseUSDCTokenPool.contract, event: "ReleasedOrMinted", logs: logs, sub: sub}, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.WatchLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(HybridLockReleaseUSDCTokenPoolReleasedOrMinted)
				if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) ParseReleasedOrMinted(log types.Log) (*HybridLockReleaseUSDCTokenPoolReleasedOrMinted, error) {
	event := new(HybridLockReleaseUSDCTokenPoolReleasedOrMinted)
	if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type HybridLockReleaseUSDCTokenPoolRemotePoolAddedIterator struct {
	Event *HybridLockReleaseUSDCTokenPoolRemotePoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *HybridLockReleaseUSDCTokenPoolRemotePoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridLockReleaseUSDCTokenPoolRemotePoolAdded)
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
		it.Event = new(HybridLockReleaseUSDCTokenPoolRemotePoolAdded)
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

func (it *HybridLockReleaseUSDCTokenPoolRemotePoolAddedIterator) Error() error {
	return it.fail
}

func (it *HybridLockReleaseUSDCTokenPoolRemotePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type HybridLockReleaseUSDCTokenPoolRemotePoolAdded struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*HybridLockReleaseUSDCTokenPoolRemotePoolAddedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.FilterLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolRemotePoolAddedIterator{contract: _HybridLockReleaseUSDCTokenPool.contract, event: "RemotePoolAdded", logs: logs, sub: sub}, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.WatchLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(HybridLockReleaseUSDCTokenPoolRemotePoolAdded)
				if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) ParseRemotePoolAdded(log types.Log) (*HybridLockReleaseUSDCTokenPoolRemotePoolAdded, error) {
	event := new(HybridLockReleaseUSDCTokenPoolRemotePoolAdded)
	if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type HybridLockReleaseUSDCTokenPoolRemotePoolRemovedIterator struct {
	Event *HybridLockReleaseUSDCTokenPoolRemotePoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *HybridLockReleaseUSDCTokenPoolRemotePoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridLockReleaseUSDCTokenPoolRemotePoolRemoved)
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
		it.Event = new(HybridLockReleaseUSDCTokenPoolRemotePoolRemoved)
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

func (it *HybridLockReleaseUSDCTokenPoolRemotePoolRemovedIterator) Error() error {
	return it.fail
}

func (it *HybridLockReleaseUSDCTokenPoolRemotePoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type HybridLockReleaseUSDCTokenPoolRemotePoolRemoved struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*HybridLockReleaseUSDCTokenPoolRemotePoolRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.FilterLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolRemotePoolRemovedIterator{contract: _HybridLockReleaseUSDCTokenPool.contract, event: "RemotePoolRemoved", logs: logs, sub: sub}, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.WatchLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(HybridLockReleaseUSDCTokenPoolRemotePoolRemoved)
				if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) ParseRemotePoolRemoved(log types.Log) (*HybridLockReleaseUSDCTokenPoolRemotePoolRemoved, error) {
	event := new(HybridLockReleaseUSDCTokenPoolRemotePoolRemoved)
	if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type HybridLockReleaseUSDCTokenPoolRouterUpdatedIterator struct {
	Event *HybridLockReleaseUSDCTokenPoolRouterUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *HybridLockReleaseUSDCTokenPoolRouterUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridLockReleaseUSDCTokenPoolRouterUpdated)
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
		it.Event = new(HybridLockReleaseUSDCTokenPoolRouterUpdated)
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

func (it *HybridLockReleaseUSDCTokenPoolRouterUpdatedIterator) Error() error {
	return it.fail
}

func (it *HybridLockReleaseUSDCTokenPoolRouterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type HybridLockReleaseUSDCTokenPoolRouterUpdated struct {
	OldRouter common.Address
	NewRouter common.Address
	Raw       types.Log
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) FilterRouterUpdated(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolRouterUpdatedIterator, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.FilterLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolRouterUpdatedIterator{contract: _HybridLockReleaseUSDCTokenPool.contract, event: "RouterUpdated", logs: logs, sub: sub}, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolRouterUpdated) (event.Subscription, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.WatchLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(HybridLockReleaseUSDCTokenPoolRouterUpdated)
				if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) ParseRouterUpdated(log types.Log) (*HybridLockReleaseUSDCTokenPoolRouterUpdated, error) {
	event := new(HybridLockReleaseUSDCTokenPoolRouterUpdated)
	if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type HybridLockReleaseUSDCTokenPoolTokensExcludedFromBurnIterator struct {
	Event *HybridLockReleaseUSDCTokenPoolTokensExcludedFromBurn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *HybridLockReleaseUSDCTokenPoolTokensExcludedFromBurnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridLockReleaseUSDCTokenPoolTokensExcludedFromBurn)
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
		it.Event = new(HybridLockReleaseUSDCTokenPoolTokensExcludedFromBurn)
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

func (it *HybridLockReleaseUSDCTokenPoolTokensExcludedFromBurnIterator) Error() error {
	return it.fail
}

func (it *HybridLockReleaseUSDCTokenPoolTokensExcludedFromBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type HybridLockReleaseUSDCTokenPoolTokensExcludedFromBurn struct {
	RemoteChainSelector          uint64
	Amount                       *big.Int
	BurnableAmountAfterExclusion *big.Int
	Raw                          types.Log
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) FilterTokensExcludedFromBurn(opts *bind.FilterOpts, remoteChainSelector []uint64) (*HybridLockReleaseUSDCTokenPoolTokensExcludedFromBurnIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.FilterLogs(opts, "TokensExcludedFromBurn", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolTokensExcludedFromBurnIterator{contract: _HybridLockReleaseUSDCTokenPool.contract, event: "TokensExcludedFromBurn", logs: logs, sub: sub}, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) WatchTokensExcludedFromBurn(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolTokensExcludedFromBurn, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.WatchLogs(opts, "TokensExcludedFromBurn", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(HybridLockReleaseUSDCTokenPoolTokensExcludedFromBurn)
				if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "TokensExcludedFromBurn", log); err != nil {
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) ParseTokensExcludedFromBurn(log types.Log) (*HybridLockReleaseUSDCTokenPoolTokensExcludedFromBurn, error) {
	event := new(HybridLockReleaseUSDCTokenPoolTokensExcludedFromBurn)
	if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "TokensExcludedFromBurn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPool) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _HybridLockReleaseUSDCTokenPool.abi.Events["AllowListAdd"].ID:
		return _HybridLockReleaseUSDCTokenPool.ParseAllowListAdd(log)
	case _HybridLockReleaseUSDCTokenPool.abi.Events["AllowListRemove"].ID:
		return _HybridLockReleaseUSDCTokenPool.ParseAllowListRemove(log)
	case _HybridLockReleaseUSDCTokenPool.abi.Events["CCTPMigrationCancelled"].ID:
		return _HybridLockReleaseUSDCTokenPool.ParseCCTPMigrationCancelled(log)
	case _HybridLockReleaseUSDCTokenPool.abi.Events["CCTPMigrationExecuted"].ID:
		return _HybridLockReleaseUSDCTokenPool.ParseCCTPMigrationExecuted(log)
	case _HybridLockReleaseUSDCTokenPool.abi.Events["CCTPMigrationProposed"].ID:
		return _HybridLockReleaseUSDCTokenPool.ParseCCTPMigrationProposed(log)
	case _HybridLockReleaseUSDCTokenPool.abi.Events["ChainAdded"].ID:
		return _HybridLockReleaseUSDCTokenPool.ParseChainAdded(log)
	case _HybridLockReleaseUSDCTokenPool.abi.Events["ChainConfigured"].ID:
		return _HybridLockReleaseUSDCTokenPool.ParseChainConfigured(log)
	case _HybridLockReleaseUSDCTokenPool.abi.Events["ChainRemoved"].ID:
		return _HybridLockReleaseUSDCTokenPool.ParseChainRemoved(log)
	case _HybridLockReleaseUSDCTokenPool.abi.Events["CircleMigratorAddressSet"].ID:
		return _HybridLockReleaseUSDCTokenPool.ParseCircleMigratorAddressSet(log)
	case _HybridLockReleaseUSDCTokenPool.abi.Events["ConfigChanged"].ID:
		return _HybridLockReleaseUSDCTokenPool.ParseConfigChanged(log)
	case _HybridLockReleaseUSDCTokenPool.abi.Events["ConfigSet"].ID:
		return _HybridLockReleaseUSDCTokenPool.ParseConfigSet(log)
	case _HybridLockReleaseUSDCTokenPool.abi.Events["DomainsSet"].ID:
		return _HybridLockReleaseUSDCTokenPool.ParseDomainsSet(log)
	case _HybridLockReleaseUSDCTokenPool.abi.Events["InboundRateLimitConsumed"].ID:
		return _HybridLockReleaseUSDCTokenPool.ParseInboundRateLimitConsumed(log)
	case _HybridLockReleaseUSDCTokenPool.abi.Events["LiquidityAdded"].ID:
		return _HybridLockReleaseUSDCTokenPool.ParseLiquidityAdded(log)
	case _HybridLockReleaseUSDCTokenPool.abi.Events["LiquidityProviderSet"].ID:
		return _HybridLockReleaseUSDCTokenPool.ParseLiquidityProviderSet(log)
	case _HybridLockReleaseUSDCTokenPool.abi.Events["LiquidityRemoved"].ID:
		return _HybridLockReleaseUSDCTokenPool.ParseLiquidityRemoved(log)
	case _HybridLockReleaseUSDCTokenPool.abi.Events["LiquidityTransferred"].ID:
		return _HybridLockReleaseUSDCTokenPool.ParseLiquidityTransferred(log)
	case _HybridLockReleaseUSDCTokenPool.abi.Events["LockReleaseDisabled"].ID:
		return _HybridLockReleaseUSDCTokenPool.ParseLockReleaseDisabled(log)
	case _HybridLockReleaseUSDCTokenPool.abi.Events["LockReleaseEnabled"].ID:
		return _HybridLockReleaseUSDCTokenPool.ParseLockReleaseEnabled(log)
	case _HybridLockReleaseUSDCTokenPool.abi.Events["LockedOrBurned"].ID:
		return _HybridLockReleaseUSDCTokenPool.ParseLockedOrBurned(log)
	case _HybridLockReleaseUSDCTokenPool.abi.Events["OutboundRateLimitConsumed"].ID:
		return _HybridLockReleaseUSDCTokenPool.ParseOutboundRateLimitConsumed(log)
	case _HybridLockReleaseUSDCTokenPool.abi.Events["OwnershipTransferRequested"].ID:
		return _HybridLockReleaseUSDCTokenPool.ParseOwnershipTransferRequested(log)
	case _HybridLockReleaseUSDCTokenPool.abi.Events["OwnershipTransferred"].ID:
		return _HybridLockReleaseUSDCTokenPool.ParseOwnershipTransferred(log)
	case _HybridLockReleaseUSDCTokenPool.abi.Events["RateLimitAdminSet"].ID:
		return _HybridLockReleaseUSDCTokenPool.ParseRateLimitAdminSet(log)
	case _HybridLockReleaseUSDCTokenPool.abi.Events["ReleasedOrMinted"].ID:
		return _HybridLockReleaseUSDCTokenPool.ParseReleasedOrMinted(log)
	case _HybridLockReleaseUSDCTokenPool.abi.Events["RemotePoolAdded"].ID:
		return _HybridLockReleaseUSDCTokenPool.ParseRemotePoolAdded(log)
	case _HybridLockReleaseUSDCTokenPool.abi.Events["RemotePoolRemoved"].ID:
		return _HybridLockReleaseUSDCTokenPool.ParseRemotePoolRemoved(log)
	case _HybridLockReleaseUSDCTokenPool.abi.Events["RouterUpdated"].ID:
		return _HybridLockReleaseUSDCTokenPool.ParseRouterUpdated(log)
	case _HybridLockReleaseUSDCTokenPool.abi.Events["TokensExcludedFromBurn"].ID:
		return _HybridLockReleaseUSDCTokenPool.ParseTokensExcludedFromBurn(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (HybridLockReleaseUSDCTokenPoolAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (HybridLockReleaseUSDCTokenPoolAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (HybridLockReleaseUSDCTokenPoolCCTPMigrationCancelled) Topic() common.Hash {
	return common.HexToHash("0x375f1ad1194a2bec317c5efec05cc63ffa06ddd0c4b276619f6fd47298eda518")
}

func (HybridLockReleaseUSDCTokenPoolCCTPMigrationExecuted) Topic() common.Hash {
	return common.HexToHash("0xdea60ddd4c7ebdab804f5694c70350cca7893ece3efeecb142312eacac5c73e4")
}

func (HybridLockReleaseUSDCTokenPoolCCTPMigrationProposed) Topic() common.Hash {
	return common.HexToHash("0x20331f191af84dbff48b162aa5a5985e7891ae646297b0a2ac80487f9109ef49")
}

func (HybridLockReleaseUSDCTokenPoolChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (HybridLockReleaseUSDCTokenPoolChainConfigured) Topic() common.Hash {
	return common.HexToHash("0x0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b")
}

func (HybridLockReleaseUSDCTokenPoolChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (HybridLockReleaseUSDCTokenPoolCircleMigratorAddressSet) Topic() common.Hash {
	return common.HexToHash("0x084e6f0e9791c2e56153bd49e6ec6dd63ba9a72c258d71558d74c63fc75b7168")
}

func (HybridLockReleaseUSDCTokenPoolConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (HybridLockReleaseUSDCTokenPoolConfigSet) Topic() common.Hash {
	return common.HexToHash("0x2e902d38f15b233cbb63711add0fca4545334d3a169d60c0a616494d7eea9544")
}

func (HybridLockReleaseUSDCTokenPoolDomainsSet) Topic() common.Hash {
	return common.HexToHash("0xe6d14ea297366c7bc1265d289d924bfd8b9afb148eb972b481f70da41c842cf5")
}

func (HybridLockReleaseUSDCTokenPoolInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (HybridLockReleaseUSDCTokenPoolLiquidityAdded) Topic() common.Hash {
	return common.HexToHash("0xc17cea59c2955cb181b03393209566960365771dbba9dc3d510180e7cb312088")
}

func (HybridLockReleaseUSDCTokenPoolLiquidityProviderSet) Topic() common.Hash {
	return common.HexToHash("0xc82aa48e67c70b1ad1494533456f52504bb4d62d11bbdafaeb98cfccd1ed817e")
}

func (HybridLockReleaseUSDCTokenPoolLiquidityRemoved) Topic() common.Hash {
	return common.HexToHash("0xc2c3f06e49b9f15e7b4af9055e183b0d73362e033ad82a07dec9bf9840171719")
}

func (HybridLockReleaseUSDCTokenPoolLiquidityTransferred) Topic() common.Hash {
	return common.HexToHash("0xf0aaad2ba9d6095d87c8b955b6c505eca51c469799b6293b841849b8e9ee3a3b")
}

func (HybridLockReleaseUSDCTokenPoolLockReleaseDisabled) Topic() common.Hash {
	return common.HexToHash("0xddc5afbc5e53c63a556964db0eef76a1c2d9305e0811abd7410d2a6f4799490e")
}

func (HybridLockReleaseUSDCTokenPoolLockReleaseEnabled) Topic() common.Hash {
	return common.HexToHash("0x5e3985e51df58346365017cae614e59d723143b71c9a2ce4a156687f1f2c3f5a")
}

func (HybridLockReleaseUSDCTokenPoolLockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
}

func (HybridLockReleaseUSDCTokenPoolOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944")
}

func (HybridLockReleaseUSDCTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (HybridLockReleaseUSDCTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (HybridLockReleaseUSDCTokenPoolRateLimitAdminSet) Topic() common.Hash {
	return common.HexToHash("0x44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174")
}

func (HybridLockReleaseUSDCTokenPoolReleasedOrMinted) Topic() common.Hash {
	return common.HexToHash("0xfc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0")
}

func (HybridLockReleaseUSDCTokenPoolRemotePoolAdded) Topic() common.Hash {
	return common.HexToHash("0x7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea")
}

func (HybridLockReleaseUSDCTokenPoolRemotePoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76")
}

func (HybridLockReleaseUSDCTokenPoolRouterUpdated) Topic() common.Hash {
	return common.HexToHash("0x02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f1684")
}

func (HybridLockReleaseUSDCTokenPoolTokensExcludedFromBurn) Topic() common.Hash {
	return common.HexToHash("0xe1e6c22ce6b566f66cdb457ec2e7910ff1f9a9e5654ed75303476fa870468220")
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPool) Address() common.Address {
	return _HybridLockReleaseUSDCTokenPool.address
}

type HybridLockReleaseUSDCTokenPoolInterface interface {
	SUPPORTEDUSDCVERSION(opts *bind.CallOpts) (uint32, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetCurrentProposedCCTPChainMigration(opts *bind.CallOpts) (uint64, error)

	GetDomain(opts *bind.CallOpts, chainSelector uint64) (USDCTokenPoolDomain, error)

	GetExcludedTokensByChain(opts *bind.CallOpts, remoteChainSelector uint64) (*big.Int, error)

	GetLiquidityProvider(opts *bind.CallOpts, remoteChainSelector uint64) (common.Address, error)

	GetLockedTokensForChain(opts *bind.CallOpts, remoteChainSelector uint64) (*big.Int, error)

	GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	GetSupportedChains(opts *bind.CallOpts) ([]uint64, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	GetTokenDecimals(opts *bind.CallOpts) (uint8, error)

	ILocalDomainIdentifier(opts *bind.CallOpts) (uint32, error)

	IMessageTransmitterProxy(opts *bind.CallOpts) (common.Address, error)

	IPreviousPool(opts *bind.CallOpts) (common.Address, error)

	ITokenMessenger(opts *bind.CallOpts) (common.Address, error)

	IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error)

	IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error)

	IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	ShouldUseLockRelease(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error)

	ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error)

	BurnLockedUSDC(opts *bind.TransactOpts) (*types.Transaction, error)

	CancelExistingCCTPMigrationProposal(opts *bind.TransactOpts) (*types.Transaction, error)

	ExcludeTokensFromBurn(opts *bind.TransactOpts, remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	ProposeCCTPMigration(opts *bind.TransactOpts, remoteChainSelector uint64) (*types.Transaction, error)

	ProvideLiquidity(opts *bind.TransactOpts, remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error)

	SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error)

	SetCircleMigratorAddress(opts *bind.TransactOpts, migrator common.Address) (*types.Transaction, error)

	SetDomains(opts *bind.TransactOpts, domains []USDCTokenPoolDomainUpdate) (*types.Transaction, error)

	SetLiquidityProvider(opts *bind.TransactOpts, remoteChainSelector uint64, liquidityProvider common.Address) (*types.Transaction, error)

	SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateChainSelectorMechanisms(opts *bind.TransactOpts, removes []uint64, adds []uint64) (*types.Transaction, error)

	WithdrawLiquidity(opts *bind.TransactOpts, remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error)

	FilterAllowListAdd(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*HybridLockReleaseUSDCTokenPoolAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*HybridLockReleaseUSDCTokenPoolAllowListRemove, error)

	FilterCCTPMigrationCancelled(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolCCTPMigrationCancelledIterator, error)

	WatchCCTPMigrationCancelled(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolCCTPMigrationCancelled) (event.Subscription, error)

	ParseCCTPMigrationCancelled(log types.Log) (*HybridLockReleaseUSDCTokenPoolCCTPMigrationCancelled, error)

	FilterCCTPMigrationExecuted(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolCCTPMigrationExecutedIterator, error)

	WatchCCTPMigrationExecuted(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolCCTPMigrationExecuted) (event.Subscription, error)

	ParseCCTPMigrationExecuted(log types.Log) (*HybridLockReleaseUSDCTokenPoolCCTPMigrationExecuted, error)

	FilterCCTPMigrationProposed(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolCCTPMigrationProposedIterator, error)

	WatchCCTPMigrationProposed(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolCCTPMigrationProposed) (event.Subscription, error)

	ParseCCTPMigrationProposed(log types.Log) (*HybridLockReleaseUSDCTokenPoolCCTPMigrationProposed, error)

	FilterChainAdded(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*HybridLockReleaseUSDCTokenPoolChainAdded, error)

	FilterChainConfigured(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolChainConfiguredIterator, error)

	WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolChainConfigured) (event.Subscription, error)

	ParseChainConfigured(log types.Log) (*HybridLockReleaseUSDCTokenPoolChainConfigured, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*HybridLockReleaseUSDCTokenPoolChainRemoved, error)

	FilterCircleMigratorAddressSet(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolCircleMigratorAddressSetIterator, error)

	WatchCircleMigratorAddressSet(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolCircleMigratorAddressSet) (event.Subscription, error)

	ParseCircleMigratorAddressSet(log types.Log) (*HybridLockReleaseUSDCTokenPoolCircleMigratorAddressSet, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*HybridLockReleaseUSDCTokenPoolConfigChanged, error)

	FilterConfigSet(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*HybridLockReleaseUSDCTokenPoolConfigSet, error)

	FilterDomainsSet(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolDomainsSetIterator, error)

	WatchDomainsSet(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolDomainsSet) (event.Subscription, error)

	ParseDomainsSet(log types.Log) (*HybridLockReleaseUSDCTokenPoolDomainsSet, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*HybridLockReleaseUSDCTokenPoolInboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*HybridLockReleaseUSDCTokenPoolInboundRateLimitConsumed, error)

	FilterLiquidityAdded(opts *bind.FilterOpts, provider []common.Address, amount []*big.Int) (*HybridLockReleaseUSDCTokenPoolLiquidityAddedIterator, error)

	WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolLiquidityAdded, provider []common.Address, amount []*big.Int) (event.Subscription, error)

	ParseLiquidityAdded(log types.Log) (*HybridLockReleaseUSDCTokenPoolLiquidityAdded, error)

	FilterLiquidityProviderSet(opts *bind.FilterOpts, oldProvider []common.Address, newProvider []common.Address, remoteChainSelector []uint64) (*HybridLockReleaseUSDCTokenPoolLiquidityProviderSetIterator, error)

	WatchLiquidityProviderSet(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolLiquidityProviderSet, oldProvider []common.Address, newProvider []common.Address, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLiquidityProviderSet(log types.Log) (*HybridLockReleaseUSDCTokenPoolLiquidityProviderSet, error)

	FilterLiquidityRemoved(opts *bind.FilterOpts, provider []common.Address, amount []*big.Int) (*HybridLockReleaseUSDCTokenPoolLiquidityRemovedIterator, error)

	WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolLiquidityRemoved, provider []common.Address, amount []*big.Int) (event.Subscription, error)

	ParseLiquidityRemoved(log types.Log) (*HybridLockReleaseUSDCTokenPoolLiquidityRemoved, error)

	FilterLiquidityTransferred(opts *bind.FilterOpts, from []common.Address, remoteChainSelector []uint64) (*HybridLockReleaseUSDCTokenPoolLiquidityTransferredIterator, error)

	WatchLiquidityTransferred(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolLiquidityTransferred, from []common.Address, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLiquidityTransferred(log types.Log) (*HybridLockReleaseUSDCTokenPoolLiquidityTransferred, error)

	FilterLockReleaseDisabled(opts *bind.FilterOpts, remoteChainSelector []uint64) (*HybridLockReleaseUSDCTokenPoolLockReleaseDisabledIterator, error)

	WatchLockReleaseDisabled(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolLockReleaseDisabled, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockReleaseDisabled(log types.Log) (*HybridLockReleaseUSDCTokenPoolLockReleaseDisabled, error)

	FilterLockReleaseEnabled(opts *bind.FilterOpts, remoteChainSelector []uint64) (*HybridLockReleaseUSDCTokenPoolLockReleaseEnabledIterator, error)

	WatchLockReleaseEnabled(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolLockReleaseEnabled, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockReleaseEnabled(log types.Log) (*HybridLockReleaseUSDCTokenPoolLockReleaseEnabled, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*HybridLockReleaseUSDCTokenPoolLockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*HybridLockReleaseUSDCTokenPoolLockedOrBurned, error)

	FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*HybridLockReleaseUSDCTokenPoolOutboundRateLimitConsumedIterator, error)

	WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseOutboundRateLimitConsumed(log types.Log) (*HybridLockReleaseUSDCTokenPoolOutboundRateLimitConsumed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*HybridLockReleaseUSDCTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*HybridLockReleaseUSDCTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*HybridLockReleaseUSDCTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*HybridLockReleaseUSDCTokenPoolOwnershipTransferred, error)

	FilterRateLimitAdminSet(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolRateLimitAdminSetIterator, error)

	WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolRateLimitAdminSet) (event.Subscription, error)

	ParseRateLimitAdminSet(log types.Log) (*HybridLockReleaseUSDCTokenPoolRateLimitAdminSet, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*HybridLockReleaseUSDCTokenPoolReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*HybridLockReleaseUSDCTokenPoolReleasedOrMinted, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*HybridLockReleaseUSDCTokenPoolRemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*HybridLockReleaseUSDCTokenPoolRemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*HybridLockReleaseUSDCTokenPoolRemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*HybridLockReleaseUSDCTokenPoolRemotePoolRemoved, error)

	FilterRouterUpdated(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolRouterUpdatedIterator, error)

	WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolRouterUpdated) (event.Subscription, error)

	ParseRouterUpdated(log types.Log) (*HybridLockReleaseUSDCTokenPoolRouterUpdated, error)

	FilterTokensExcludedFromBurn(opts *bind.FilterOpts, remoteChainSelector []uint64) (*HybridLockReleaseUSDCTokenPoolTokensExcludedFromBurnIterator, error)

	WatchTokensExcludedFromBurn(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolTokensExcludedFromBurn, remoteChainSelector []uint64) (event.Subscription, error)

	ParseTokensExcludedFromBurn(log types.Log) (*HybridLockReleaseUSDCTokenPoolTokensExcludedFromBurn, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
