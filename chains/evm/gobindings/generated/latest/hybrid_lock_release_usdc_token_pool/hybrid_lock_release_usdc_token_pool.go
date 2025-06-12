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
	OriginalSender      []byte
	RemoteChainSelector uint64
	Receiver            common.Address
	Amount              *big.Int
	LocalToken          common.Address
	SourcePoolAddress   []byte
	SourcePoolData      []byte
	OffchainTokenData   []byte
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"internalType\":\"contractITokenMessenger\"},{\"name\":\"cctpMessageTransmitterProxy\",\"type\":\"address\",\"internalType\":\"contractCCTPMessageTransmitterProxy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"previousPool\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"SUPPORTED_USDC_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnLockedUSDC\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelExistingCCTPMigrationProposal\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"excludeTokensFromBurn\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentProposedCCTPChainMigration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDomain\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structUSDCTokenPool.Domain\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExcludedTokensByChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLiquidityProvider\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLockedTokensForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_localDomainIdentifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_messageTransmitterProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractCCTPMessageTransmitterProxy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_previousPool\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_tokenMessenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITokenMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposeCCTPMigration\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"provideLiquidity\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCircleMigratorAddress\",\"inputs\":[{\"name\":\"migrator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDomains\",\"inputs\":[{\"name\":\"domains\",\"type\":\"tuple[]\",\"internalType\":\"structUSDCTokenPool.DomainUpdate[]\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLiquidityProvider\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"liquidityProvider\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRouter\",\"inputs\":[{\"name\":\"newRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shouldUseLockRelease\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateChainSelectorMechanisms\",\"inputs\":[{\"name\":\"removes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"adds\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawLiquidity\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCTPMigrationCancelled\",\"inputs\":[{\"name\":\"existingProposalSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCTPMigrationExecuted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"USDCBurned\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCTPMigrationProposed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CircleMigratorAddressSet\",\"inputs\":[{\"name\":\"migratorAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DomainsSet\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structUSDCTokenPool.DomainUpdate[]\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityAdded\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityProviderSet\",\"inputs\":[{\"name\":\"oldProvider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newProvider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityRemoved\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockReleaseDisabled\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockReleaseEnabled\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RouterUpdated\",\"inputs\":[{\"name\":\"oldRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensExcludedFromBurn\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"burnableAmountAfterExclusion\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"ExistingMigrationProposal\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidChainSelector\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestinationDomain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidDomain\",\"inputs\":[{\"name\":\"domain\",\"type\":\"tuple\",\"internalType\":\"structUSDCTokenPool.DomainUpdate\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidNonce\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"got\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourceDomain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenMessengerVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidTransmitterInProxy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LanePausedForCCTPMigration\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoMigrationProposalPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenLockingNotAllowedAfterMigration\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnknownDomain\",\"inputs\":[{\"name\":\"domain\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"UnlockingUSDCFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"onlyCircle\",\"inputs\":[]}]",
	Bin: "0x6101a0806040523461053757616cac803803809161001d82856108db565b833981019060e0818303126105375780516001600160a01b03811690818103610537576020830151906001600160a01b038216808303610537576040850151956001600160a01b038716958688036105375760608101516001600160401b0381116105375781019180601f84011215610537578251926001600160401b03841161053c578360051b9060208201946100b860405196876108db565b855260208086019282010192831161053757602001905b8282106108c3575050506100e5608082016108fe565b6100fd60c06100f660a085016108fe565b93016108fe565b9833156108b257600180546001600160a01b03191633179055881580156108a1575b8015610890575b61087f5760805260c05260405163313ce56760e01b81526020816004818b5afa8091600091610843575b509061081f575b50600660a052600480546001600160a01b0319166001600160a01b03929092169190911790558051151560e08190526106f2575b5083156106e157604051632c12192160e01b8152602081600481885afa9081156105a7576000916106a7575b5060405163054fd4d560e41b81526001600160a01b03919091169190602081600481865afa80156105a75763ffffffff91600091610688575b5016806106745750604051639cdbb18160e01b8152602081600481895afa80156105a75763ffffffff91600091610655575b50168061064157506020600491604051928380926367e0ed8360e11b82525afa80156105a75782916000916105f3575b506001600160a01b0316036105e25760049260209261010052610120526040519283809263234d8e3d60e21b82525afa9081156105a7576000916105b3575b506101405260805161010051604051636eb1769f60e11b81523060048201526001600160a01b0391821660248201819052959290911690602081604481855afa9081156105a757600091610575575b50600019810180911161055f5760405190602082019663095ea7b360e01b885260248301526044820152604481526103196064826108db565b60008060409788519361032c8a866108db565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082865af13d15610552573d906001600160401b03821161053c57875161039d94909261038e601f8201601f1916602001856108db565b83523d6000602085013e610ab8565b8051806104bb575b50507f2e902d38f15b233cbb63711add0fca4545334d3a169d60c0a616494d7eea954491602091610160528451908152a161018052516161239081610b898239608051818181610536015281816108a80152818161092f01528181611eb301528181612df6015281816131db01528181614484015281816148b401528181615b790152615dae015260a05181610979015260c0518181816123d901528181614dce015261555a015260e051818181610e00015281816125c60152615be5015261010051818181611002015261328901526101205181818161193801526144020152610140518181816110c6015281816133620152614f71015261016051818181610a52015261437c0152610180518161160b0152f35b81602091810103126105375760200151801590811503610537576104e05738806103a5565b835162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b9161039d92606091610ab8565b634e487b7160e01b600052601160045260246000fd5b90506020813d60201161059f575b81610590602093836108db565b810103126105375751386102e0565b3d9150610583565b6040513d6000823e3d90fd5b6105d5915060203d6020116105db575b6105cd81836108db565b810190610912565b38610291565b503d6105c3565b632a32133b60e11b60005260046000fd5b9091506020813d602011610639575b8161060f602093836108db565b810103126106355751906001600160a01b03821682036106325750819038610252565b80fd5b5080fd5b3d9150610602565b6316ba39c560e31b60005260045260246000fd5b61066e915060203d6020116105db576105cd81836108db565b38610222565b6334697c6b60e11b60005260045260246000fd5b6106a1915060203d6020116105db576105cd81836108db565b386101f0565b90506020813d6020116106d9575b816106c2602093836108db565b81010312610537576106d3906108fe565b386101b7565b3d91506106b5565b6306b7c75960e31b60005260046000fd5b9091946020946040519461070687876108db565b60008652600036813760e0511561080e5760005b8651811015610781576001906001600160a01b03610738828a61092e565b51168961074482610970565b610751575b50500161071a565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a13889610749565b509193969092945060005b8551811015610800576001906001600160a01b036107aa828961092e565b511680156107fa57886107bc82610a58565b6107ca575b50505b0161078c565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a138886107c1565b506107c4565b50929591945092503861018b565b6335f4a7b360e01b60005260046000fd5b60ff1660068114610157576332ad3e0760e11b600052600660045260245260446000fd5b6020813d602011610877575b8161085c602093836108db565b8101031261063557519060ff82168203610632575038610150565b3d915061084f565b6342bcdf7f60e11b60005260046000fd5b506001600160a01b03821615610126565b506001600160a01b0383161561011f565b639b15e16f60e01b60005260046000fd5b602080916108d0846108fe565b8152019101906100cf565b601f909101601f19168101906001600160401b0382119082101761053c57604052565b51906001600160a01b038216820361053757565b90816020910312610537575163ffffffff811681036105375790565b80518210156109425760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b80548210156109425760005260206000200190600090565b6000818152600360205260409020548015610a5157600019810181811161055f5760025460001981019190821161055f57818103610a00575b50505060025480156109ea57600019016109c4816002610958565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b610a39610a11610a22936002610958565b90549060031b1c9283926002610958565b819391549060031b91821b91600019901b19161790565b905560005260036020526040600020553880806109a9565b5050600090565b80600052600360205260406000205415600014610ab2576002546801000000000000000081101561053c57610a99610a228260018594016002556002610958565b9055600254906000526003602052604060002055600190565b50600090565b91929015610b1a5750815115610acc575090565b3b15610ad55790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b825190915015610b2d5750805190602001fd5b6040519062461bcd60e51b8252602060048301528181519182602483015260005b838110610b705750508160006044809484010152601f80199101168101030190fd5b60208282018101516044878401015285935001610b4e56fe6080604052600436101561001257600080fd5b60003560e01c806301ffc9a7146103475780631101dbd414610342578063181f5a771461033d57806321df0da714610338578063240028e81461033357806324f65ee71461032e5780632cfbb1191461032957806339077537146103245780633e591f2c1461031f5780634ad01f0b1461031a5780634c5ef0ed146103155780634c93ef841461031057806350d1a35a1461030b57806354c8a4f3146103065780636155cda01461030157806362ddd3c4146102fc5780636b716b0d146102f75780636b795423146102f25780636d3d1a58146102ed578063714bf907146102e857806379ba5097146102e35780637d54534e146102de5780638926f54f146102d95780638a5e52bb146102d45780638da5cb5b146102cf578063962d4020146102ca57806398db9643146102c55780639a4575b9146102c05780639fdf13ff146102bb578063a42a7b8b146102b6578063a7cd63b7146102b1578063acfecf91146102ac578063af58d59f146102a7578063b0f479a1146102a2578063b79465801461029d578063bb5eced314610298578063c0d7865514610293578063c4bffe2b1461028e578063c75eea9c14610289578063c781d0e314610284578063cd306a6c1461027f578063cf7401f31461027a578063dc0bd97114610275578063de814c5714610270578063dfadfa351461026b578063e0351e1314610266578063e8a1da1714610261578063e94ae6d01461025c578063f2fde38b14610257578063f65a8886146102525763fd6768551461024d57600080fd5b612b4c565b612b0d565b612a37565b6129e2565b6125eb565b6125ae565b6124e5565b6123fd565b6123ac565b6122ae565b612189565b61212c565b6120e1565b61204e565b611f2c565b611e4f565b611e13565b611ddf565b611d33565b611bff565b611b8b565b611a89565b6119ef565b61195c565b61190b565b6117a3565b61173e565b611506565b6114c7565b611436565b61136b565b6112da565b6112a6565b6110ea565b6110a9565b611026565b610fd5565b610dce565b610c05565b610bc5565b610bac565b610a76565b610a25565b6109dc565b61099d565b61095f565b6108f5565b61087b565b610818565b61043d565b34610419576020600319360112610419576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361041957807faff2afbf00000000000000000000000000000000000000000000000000000000602092149081156103ef575b81156103c5575b506040519015158152f35b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014386103ba565b7f0e64dd2900000000000000000000000000000000000000000000000000000000811491506103b3565b600080fd5b67ffffffffffffffff81160361041957565b359061043b8261041e565b565b346104195760406003193601126104195760043561045a8161041e565b6024359061049961047f8267ffffffffffffffff166000526011602052604060002090565b5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff339116036105f55767ffffffffffffffff81166104d8816000526010602052604060002054151590565b6105bd57600b546104fe9060a01c67ffffffffffffffff165b67ffffffffffffffff1690565b14610582576105219067ffffffffffffffff16600052600c602052604060002090565b61052c828254612c33565b905561055a8130337f0000000000000000000000000000000000000000000000000000000000000000613f99565b337fc17cea59c2955cb181b03393209566960365771dbba9dc3d510180e7cb312088600080a3005b7fd0da86c40000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b6000fd5b7f646972460000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b7f8e4a23d6000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b600091031261041957565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6060810190811067ffffffffffffffff82111761067957604052565b61062e565b6080810190811067ffffffffffffffff82111761067957604052565b6020810190811067ffffffffffffffff82111761067957604052565b6040810190811067ffffffffffffffff82111761067957604052565b60a0810190811067ffffffffffffffff82111761067957604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761067957604052565b6040519061043b60a0836106ee565b6040519061043b6020836106ee565b6040519061043b6040836106ee565b6040519061043b6080836106ee565b67ffffffffffffffff811161067957601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b8481106107ef5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b806020809284010151828286010152016107b0565b9060206108159281815201906107a5565b90565b3461041957600060031936011261041957610877604080519061083b81836106ee565b601782527f55534443546f6b656e506f6f6c20312e362e312d6465760000000000000000006020830152519182916020835260208301906107a5565b0390f35b3461041957600060031936011261041957602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b73ffffffffffffffffffffffffffffffffffffffff81160361041957565b359061043b826108cc565b34610419576020600319360112610419576020610955600435610917816108cc565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691161490565b6040519015158152f35b3461041957600060031936011261041957602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346104195760206003193601126104195767ffffffffffffffff6004356109c38161041e565b16600052600c6020526020604060002054604051908152f35b346104195760206003193601126104195760043567ffffffffffffffff811161041957610100600319823603011261041957610a1c602091600401612ca4565b60405190518152f35b3461041957600060031936011261041957602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461041957600060031936011261041957610a8f6145b9565b600b5467ffffffffffffffff8160a01c168015610b25577fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff7f375f1ad1194a2bec317c5efec05cc63ffa06ddd0c4b276619f6fd47298eda5189216600b556000610b0d8267ffffffffffffffff16600052600d602052604060002090565b5560405167ffffffffffffffff9091168152602090a1005b7fa94cb9880000000000000000000000000000000000000000000000000000000060005260046000fd5b604060031982011261041957600435610b678161041e565b9160243567ffffffffffffffff811161041957826023820112156104195780600401359267ffffffffffffffff84116104195760248483010111610419576024019190565b34610419576020610955610bbf36610b4f565b91612f3d565b34610419576020600319360112610419576020610955600435610be78161041e565b67ffffffffffffffff16600052600e60205260ff6040600020541690565b3461041957602060031936011261041957600435610c228161041e565b610c2a6145b9565b67ffffffffffffffff600b5460a01c16610d215767ffffffffffffffff81166000908152600e602052604090205460ff1615610cf757610cf281610cd77f20331f191af84dbff48b162aa5a5985e7891ae646297b0a2ac80487f9109ef49937fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff7bffffffffffffffff0000000000000000000000000000000000000000600b549260a01b16911617600b55565b60405167ffffffffffffffff90911681529081906020820190565b0390a1005b7f656535ce0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f692bc1310000000000000000000000000000000000000000000000000000000060005260046000fd5b9181601f840112156104195782359167ffffffffffffffff8311610419576020808501948460051b01011161041957565b60406003198201126104195760043567ffffffffffffffff81116104195781610da791600401610d4b565b929092916024359067ffffffffffffffff821161041957610dca91600401610d4b565b9091565b3461041957610df6610dfe610de236610d7c565b9491610def9391936145b9565b3691612f9d565b923691612f9d565b7f000000000000000000000000000000000000000000000000000000000000000015610fab5760005b8251811015610eec5780610e5a610e4060019386613509565b5173ffffffffffffffffffffffffffffffffffffffff1690565b610e96610e9173ffffffffffffffffffffffffffffffffffffffff83165b73ffffffffffffffffffffffffffffffffffffffff1690565b61575a565b610ea2575b5001610e27565b60405173ffffffffffffffffffffffffffffffffffffffff9190911681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756690602090a138610e9b565b5060005b8151811015610fa95780610f09610e4060019385613509565b73ffffffffffffffffffffffffffffffffffffffff811615610fa357610f4c610f4773ffffffffffffffffffffffffffffffffffffffff8316610e78565b6150c9565b610f59575b505b01610ef0565b60405173ffffffffffffffffffffffffffffffffffffffff9190911681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d890602090a183610f51565b50610f53565b005b7f35f4a7b30000000000000000000000000000000000000000000000000000000060005260046000fd5b3461041957600060031936011261041957602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346104195761103436610b4f565b61103f9291926145b9565b67ffffffffffffffff8216611061816000526006602052604060002054151590565b1561107c5750610fa992611076913691612f06565b90614625565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b3461041957600060031936011261041957602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610419576110f836610d7c565b9290916111036145b9565b60005b81811061122b5750505060005b82811061111c57005b6111486111356104f1611130848787613022565b613037565b6000526010602052604060002054151590565b6111e457806111a861117d6111636111306001958888613022565b67ffffffffffffffff16600052600e602052604060002090565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055565b6111b96104f1611130838787613022565b7f5e3985e51df58346365017cae614e59d723143b71c9a2ce4a156687f1f2c3f5a600080a201611113565b611130906105b9936111f593613022565b7f646972460000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b8061126a6112426111636111306001958789613022565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008154169055565b61127b6104f1611130838688613022565b7fddc5afbc5e53c63a556964db0eef76a1c2d9305e0811abd7410d2a6f4799490e600080a201611106565b3461041957600060031936011261041957602073ffffffffffffffffffffffffffffffffffffffff60095416604051908152f35b34610419576020600319360112610419577f084e6f0e9791c2e56153bd49e6ec6dd63ba9a72c258d71558d74c63fc75b7168602073ffffffffffffffffffffffffffffffffffffffff60043561132f816108cc565b6113376145b9565b16807fffffffffffffffffffffffff0000000000000000000000000000000000000000600b541617600b55604051908152a1005b346104195760006003193601126104195760005473ffffffffffffffffffffffffffffffffffffffff8116330361140c577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b34610419576020600319360112610419577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174602073ffffffffffffffffffffffffffffffffffffffff60043561148b816108cc565b6114936145b9565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000006009541617600955604051908152a1005b3461041957602060031936011261041957602061095567ffffffffffffffff6004356114f28161041e565b166000526006602052604060002054151590565b3461041957600060031936011261041957600b5461153973ffffffffffffffffffffffffffffffffffffffff8216610e78565b33036117145760a01c67ffffffffffffffff1667ffffffffffffffff8116908115610b25576115a561157f8267ffffffffffffffff16600052600c602052604060002090565b5461159e8367ffffffffffffffff16600052600d602052604060002090565b549061308f565b9060006115c68267ffffffffffffffff16600052600c602052604060002090565b556115f47fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff600b5416600b55565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692833b1561041957600060405180957f42966c6800000000000000000000000000000000000000000000000000000000825281838161167289600483019190602083019252565b03925af190811561170f577fdea60ddd4c7ebdab804f5694c70350cca7893ece3efeecb142312eacac5c73e4946116cf926116f4575b506116ca6112428467ffffffffffffffff16600052600e602052604060002090565b61515a565b506040805167ffffffffffffffff909216825260208201929092529081908101610cf2565b806117036000611709936106ee565b80610623565b386116a8565b61309c565b7f438a7a050000000000000000000000000000000000000000000000000000000060005260046000fd5b3461041957600060031936011261041957602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b9181601f840112156104195782359167ffffffffffffffff8311610419576020808501946060850201011161041957565b346104195760606003193601126104195760043567ffffffffffffffff8111610419576117d4903690600401610d4b565b9060243567ffffffffffffffff8111610419576117f5903690600401611772565b9060443567ffffffffffffffff811161041957611816903690600401611772565b611838610e7860095473ffffffffffffffffffffffffffffffffffffffff1690565b331415806118e0575b6105f5578386148015906118d6575b6118ac5760005b86811061186057005b806118a66118746111306001948b8b613022565b61187f8389896130a8565b6118a061189861189086898b6130a8565b923690612265565b913690612265565b916146ea565b01611857565b7f568efce20000000000000000000000000000000000000000000000000000000060005260046000fd5b5080861415611850565b50611903610e7860015473ffffffffffffffffffffffffffffffffffffffff1690565b331415611841565b3461041957600060031936011261041957602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346104195760206003193601126104195760043567ffffffffffffffff81116104195760a060031982360301126104195761199c610877916004016130d1565b6040519182916020835260206119bd825160408387015260608601906107a5565b9101517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08483030160408501526107a5565b3461041957600060031936011261041957602060405160008152f35b602081016020825282518091526040820191602060408360051b8301019401926000915b838310611a3e57505050505090565b9091929394602080611a7a837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0866001960301875289516107a5565b97019301930191939290611a2f565b346104195760206003193601126104195767ffffffffffffffff600435611aaf8161041e565b166000526007602052611ac8600560406000200161567c565b8051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611b0e611af884612f85565b93611b0660405195866106ee565b808552612f85565b0160005b818110611b7a57505060005b8151811015611b6c5780611b50611b4b611b3a60019486613509565b516000526008602052604060002090565b613570565b611b5a8286613509565b52611b658185613509565b5001611b1e565b604051806108778582611a0b565b806060602080938701015201611b12565b3461041957600060031936011261041957611ba46155e6565b60405180916020820160208352815180915260206040840192019060005b818110611bd0575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101611bc2565b3461041957611c0d36610b4f565b611c189291926145b9565b67ffffffffffffffff821691611c42611c3e846000526006602052604060002054151590565b1590565b611cfc57611c85611c3e6005611c6c8467ffffffffffffffff166000526007602052604060002090565b01611c78368689612f06565b60208151910120906158fe565b611cc157507f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d769192611cbc60405192839283613690565b0390a2005b611cf884926040519384937f74f23c7c0000000000000000000000000000000000000000000000000000000085526004850161366f565b0390fd5b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b346104195760206003193601126104195767ffffffffffffffff600435611d598161041e565b611d616136a1565b50166000526007602052610877611d86611d8160026040600020016136cc565b61493c565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b3461041957600060031936011261041957602073ffffffffffffffffffffffffffffffffffffffff60045416604051908152f35b3461041957602060031936011261041957610877611e3b600435611e368161041e565b613726565b6040519182916020835260208301906107a5565b3461041957604060031936011261041957600435611e6c8161041e565b60243590611e786145b9565b67ffffffffffffffff80600b5460a01c169116908114611eff57600052600c6020526040600020611eaa82825461308f565b9055611ed781337f00000000000000000000000000000000000000000000000000000000000000006149de565b337fc2c3f06e49b9f15e7b4af9055e183b0d73362e033ad82a07dec9bf9840171719600080a3005b7fd0da86c40000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346104195760206003193601126104195773ffffffffffffffffffffffffffffffffffffffff600435611f5e816108cc565b611f666145b9565b168015611fe05760407f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f16849160045490807fffffffffffffffffffffffff000000000000000000000000000000000000000083161760045573ffffffffffffffffffffffffffffffffffffffff8351921682526020820152a1005b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b602060408183019282815284518094520192019060005b81811061202e5750505090565b825167ffffffffffffffff16845260209384019390920191600101612021565b3461041957600060031936011261041957612067615631565b8051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612097611af884612f85565b0136602084013760005b81518110156120d3578067ffffffffffffffff6120c060019385613509565b51166120cc8286613509565b52016120a1565b60405180610877858261200a565b346104195760206003193601126104195767ffffffffffffffff6004356121078161041e565b61210f6136a1565b50166000526007602052610877611d86611d8160406000206136cc565b346104195760206003193601126104195760043567ffffffffffffffff8111610419573660238201121561041957806004013567ffffffffffffffff81116104195736602460a0830284010111610419576024610fa99201613748565b3461041957600060031936011261041957602067ffffffffffffffff600b5460a01c16604051908152f35b8015150361041957565b35906fffffffffffffffffffffffffffffffff8216820361041957565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c606091011261041957604051906122128261065d565b8160843561221f816121b4565b815260a4356fffffffffffffffffffffffffffffffff8116810361041957602082015260c435906fffffffffffffffffffffffffffffffff821682036104195760400152565b91908260609103126104195760405161227d8161065d565b60406122a98183958035612290816121b4565b855261229e602082016121be565b6020860152016121be565b910152565b346104195760e0600319360112610419576004356122cb8161041e565b60607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc360112610419576040516123018161065d565b60243561230d816121b4565b81526044356fffffffffffffffffffffffffffffffff811681036104195760208201526064356fffffffffffffffffffffffffffffffff8116810361041957604082015261235a366121db565b9073ffffffffffffffffffffffffffffffffffffffff600954163314158061238a575b6105f557610fa9926146ea565b5073ffffffffffffffffffffffffffffffffffffffff6001541633141561237d565b3461041957600060031936011261041957602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346104195760406003193601126104195760043561241a8161041e565b602435906124266145b9565b67ffffffffffffffff600b5460a01c169167ffffffffffffffff8216809303610b255782600052600d60205260406000208054928284018094116124e0577fe1e6c22ce6b566f66cdb457ec2e7910ff1f9a9e5654ed75303476fa870468220936124c8925561159e6124ac8267ffffffffffffffff16600052600c602052604060002090565b549167ffffffffffffffff16600052600d602052604060002090565b60408051928352602083019190915281908101611cbc565b612c04565b346104195760206003193601126104195767ffffffffffffffff60043561250b8161041e565b6000606060405161251b8161067e565b828152826020820152826040820152015216600052600a602052610877604060002060ff60026040519261254e8461067e565b8054845260018101546020850152015463ffffffff8116604084015260201c1615156060820152604051918291829190916060806080830194805184526020810151602085015263ffffffff604082015116604085015201511515910152565b346104195760006003193601126104195760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b34610419576125f936610d7c565b9190926126046145b9565b6000915b8083106128a95750505060009163ffffffff4216925b82811061262757005b61263a612635828585613b68565b613c42565b90606082016126498151614a78565b60808301936126588551614a78565b604084019081515115611fe057612685611c3e6126806104f1885167ffffffffffffffff1690565b6151b0565b61285e576127be6126be6126a4879a999a5167ffffffffffffffff1690565b67ffffffffffffffff166000526007602052604060002090565b6127818961277b87516127626126e760408301516fffffffffffffffffffffffffffffffff1690565b9161274961271261270b60208401516fffffffffffffffffffffffffffffffff1690565b9251151590565b61274061271d61072f565b6fffffffffffffffffffffffffffffffff851681529763ffffffff166020890152565b15156040870152565b6fffffffffffffffffffffffffffffffff166060850152565b6fffffffffffffffffffffffffffffffff166080830152565b82613cd1565b6127b3896127aa8a516127626126e760408301516fffffffffffffffffffffffffffffffff1690565b60028301613cd1565b600484519101613ddd565b602085019660005b8851805182101561280157906127fb6001926127f4836127ee8c5167ffffffffffffffff1690565b92613509565b5190614625565b016127c6565b505097965094906128557f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c293926128426001975167ffffffffffffffff1690565b9251935190519060405194859485613f04565b0390a10161261e565b6105b9612873865167ffffffffffffffff1690565b7f1d5ad3c50000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b9091926128ba611130858486613022565b946128d1611c3e67ffffffffffffffff8816615837565b6129aa576128fe60056128f88867ffffffffffffffff166000526007602052604060002090565b0161567c565b9360005b855181101561294a5760019061294360056129318b67ffffffffffffffff166000526007602052604060002090565b0161293c838a613509565b51906158fe565b5001612902565b509350937f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d85991661299c60019397610cd76129978267ffffffffffffffff166000526007602052604060002090565b613ab9565b0390a1019190939293612608565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff861660045260246000fd5b346104195760206003193601126104195767ffffffffffffffff600435612a088161041e565b166000526011602052602073ffffffffffffffffffffffffffffffffffffffff60406000205416604051908152f35b346104195760206003193601126104195773ffffffffffffffffffffffffffffffffffffffff600435612a69816108cc565b612a716145b9565b16338114612ae357807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346104195760206003193601126104195767ffffffffffffffff600435612b338161041e565b16600052600d6020526020604060002054604051908152f35b3461041957604060031936011261041957600435612b698161041e565b67ffffffffffffffff60243591612b7f836108cc565b612b876145b9565b166000818152601160205260408120805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff0000000000000000000000000000000000000000821681179092559293909216907fc82aa48e67c70b1ad1494533456f52504bb4d62d11bbdafaeb98cfccd1ed817e9080a4005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b919082018092116124e057565b60405190612c4d8261069a565b60008252565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610419570180359067ffffffffffffffff82116104195760200191813603831361041957565b604051612cb08161069a565b600090527ffa7c07de000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000612d0360c0840184612c53565b90358281169160048110612eb5575b50501603612eac57612d22612c40565b50612d2c81614d00565b600b5460a01c67ffffffffffffffff166020820190612d4d6104f183613037565b67ffffffffffffffff8216146105825750907ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc067ffffffffffffffff83612db3612d9961113096613037565b67ffffffffffffffff16600052600c602052604060002090565b54612e91576060840135612dee612de6612dcc84613037565b67ffffffffffffffff16600052600d602052604060002090565b91825461308f565b90555b612e807f000000000000000000000000000000000000000000000000000000000000000094612e3f612e3960408301946060612e2c876142e9565b940135998a80958b6149de565b936142e9565b6040805173ffffffffffffffffffffffffffffffffffffffff9889168152336020820152919097169681019690965260608601529116929081906080820190565b0390a2612e8b61073e565b90815290565b6060840135612ea5612de6612d9984613037565b9055612df1565b610815906142f3565b839250829060040360031b1b16163880612d12565b91612f02918354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055565b929192612f128261076b565b91612f2060405193846106ee565b829481845281830111610419578281602093846000960137010152565b610815929167ffffffffffffffff612f68921660005260076020526005604060002001923691612f06565b602081519101209060019160005201602052604060002054151590565b67ffffffffffffffff81116106795760051b60200190565b929190612fa981612f85565b93612fb760405195866106ee565b602085838152019160051b810192831161041957905b828210612fd957505050565b602080918335612fe8816108cc565b815201910190612fcd565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b91908110156130325760051b0190565b612ff3565b356108158161041e565b67ffffffffffffffff61081591166000526006602052604060002054151590565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116124e057565b919082039182116124e057565b6040513d6000823e3d90fd5b9190811015613032576060020190565b604051906130c5826106b6565b60606020838281520152565b6130d96130b8565b50602081016130ec8135610be78161041e565b1561312e576131146104f161310e600b5467ffffffffffffffff9060a01c1690565b92613037565b67ffffffffffffffff82161461058257506108159061483e565b6131399291926130b8565b5061314383615510565b61317161316c61315283613037565b67ffffffffffffffff16600052600a602052604060002090565b613a1f565b613181611c3e6060830151151590565b6134c75760206131918580612c53565b905003613487576020810151929361326e938015613467576060602091925b0135916131c4604085015163ffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169687955160405198899485947ff856ddb60000000000000000000000000000000000000000000000000000000086528860048701919360809363ffffffff73ffffffffffffffffffffffffffffffffffffffff9398979660a0860199865216602085015260408401521660608201520152565b0381600073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af193841561170f576000946133f6575b50611e36836133e3937ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff61333f956133376133006133b79a613037565b6040805173ffffffffffffffffffffffffffffffffffffffff90971687523360208801528601929092529116929081906060820190565b0390a2613037565b9261335b61334b61074d565b67ffffffffffffffff9092168252565b63ffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015260405192839160208301919091602063ffffffff81604084019567ffffffffffffffff8151168552015116910152565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826106ee565b6133eb61074d565b918252602082015290565b61333f9194506133b7936133e3937ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff613452611e369560203d602011613460575b61344a81836106ee565b810190614829565b9895505050935093506132ba565b503d613440565b50602060606134816134798480612c53565b81019061481a565b926131b0565b6134918480612c53565b90611cf86040519283927fa3c8cf0900000000000000000000000000000000000000000000000000000000845260048401613690565b6105b96134d383613037565b7fd201c48a0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b80518210156130325760209160051b010190565b90600182811c92168015613566575b602083101461353757565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f169161352c565b90604051918260008254926135848461351d565b80845293600181169081156135f057506001146135a9575b5061043b925003836106ee565b90506000929192526020600020906000915b8183106135d457505090602061043b928201013861359c565b60209193508060019154838589010152019101909184926135bb565b6020935061043b9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b8201013861359c565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b60409067ffffffffffffffff61081595931681528160208201520191613630565b916020610815938181520191613630565b604051906136ae826106d2565b60006080838281528260208201528260408201528260608201520152565b906040516136d9816106d2565b60806fffffffffffffffffffffffffffffffff6001839560ff8154848116875263ffffffff81871c16602088015260a01c1615156040860152015481808216166060850152821c16910152565b67ffffffffffffffff1660005260076020526108156004604060002001613570565b6137506145b9565b60005b8281106137925750907fe6d14ea297366c7bc1265d289d924bfd8b9afb148eb972b481f70da41c842cf59161378d60405192839283613996565b0390a1565b6137a56137a0828585613917565b613938565b80511580156138f1575b613884579061387e8261382461315260606137d36040600198015163ffffffff1690565b93613815602082015161380d8351976137ef6080860151151590565b926137f861075c565b998a5260208a015263ffffffff166040890152565b151586840152565b015167ffffffffffffffff1690565b6002908251815560208301516001820155019063ffffffff6040820151167fffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000064ff0000000060608554940151151560201b16921617179055565b01613753565b604080517fa606c63500000000000000000000000000000000000000000000000000000000815282516004820152602083015160248201529082015163ffffffff166044820152606082015167ffffffffffffffff1660648201526080909101511515608482015260a490fd5b5067ffffffffffffffff613910606083015167ffffffffffffffff1690565b16156137af565b91908110156130325760a0020190565b359063ffffffff8216820361041957565b60a08136031261041957608060405191613951836106d2565b803583526020810135602084015261396b60408201613927565b6040840152606081013561397e8161041e565b6060840152013561398e816121b4565b608082015290565b602080825281018390526040019160005b8181106139b45750505090565b90919260a080600192863581526020870135602082015263ffffffff6139dc60408901613927565b16604082015267ffffffffffffffff60608801356139f98161041e565b1660608201526080870135613a0d816121b4565b151560808201520194019291016139a7565b90604051613a2c8161067e565b606060ff600283958054855260018101546020860152015463ffffffff8116604085015260201c161515910152565b818110613a66575050565b60008155600101613a5b565b818102929181159184041417156124e057565b8054906000815581613a95575050565b6000526020600020908101905b818110613aad575050565b60008155600101613aa2565b600561043b916000815560006001820155600060028201556000600382015560048101613ae6815461351d565b9081613af5575b505001613a85565b81601f60009311600114613b0d5750555b3880613aed565b81835260208320613b2891601f01861c810190600101613a5b565b808252602082209081548360011b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8560031b1c191617905555613b06565b91908110156130325760051b810135907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee181360301821215610419570190565b9080601f830112156104195781602061081593359101612f06565b9080601f83011215610419578135613bda81612f85565b92613be860405194856106ee565b81845260208085019260051b820101918383116104195760208201905b838210613c1457505050505090565b813567ffffffffffffffff811161041957602091613c3787848094880101613ba8565b815201910190613c05565b610120813603126104195760405190613c5a826106d2565b613c6381610430565b8252602081013567ffffffffffffffff811161041957613c869036908301613bc3565b602083015260408101359067ffffffffffffffff821161041957613cb061398e9236908301613ba8565b6040840152613cc23660608301612265565b606084015260c0369101612265565b8151815460208401516040850151608091821b73ffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffff000000000000000000000000000000000000000000909416939093179290921791151560a01b74ff000000000000000000000000000000000000000016919091178355606084015193810151901b7fffffffffffffffffffffffffffffffff0000000000000000000000000000000016921691909117600190910155565b9190601f8111613da757505050565b61043b926000526020600020906020601f840160051c83019310613dd3575b601f0160051c0190613a5b565b9091508190613dc6565b919091825167ffffffffffffffff811161067957613e0581613dff845461351d565b84613d98565b6020601f8211600114613e5f578190612f02939495600092613e54575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b015190503880613e22565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0821690613e9284600052602060002090565b9160005b818110613eec57509583600195969710613eb5575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388080613eab565b9192602060018192868b015181550194019201613e96565b613f68613f3361043b9597969467ffffffffffffffff60a09516845261010060208501526101008401906107a5565b9660408301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b01906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff9283166024820152929091166044830152606482019290925261043b9161402982608481015b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018452836106ee565b614c44565b908160409103126104195761405e60206040519261404b846106b6565b80356140568161041e565b845201613927565b602082015290565b6020818303126104195780359067ffffffffffffffff82116104195701604081830312610419576040519161409a836106b6565b813567ffffffffffffffff811161041957816140b7918401613ba8565b8352602082013567ffffffffffffffff81116104195761405e9201613ba8565b9081602091031261041957604051906140ef8261069a565b51815290565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561041957016020813591019167ffffffffffffffff821161041957813603831361041957565b90610815916020815261427e61427361423661417761416486806140f5565b6101006020880152610120870191613630565b61419761418660208801610430565b67ffffffffffffffff166040870152565b6141c36141a6604088016108ea565b73ffffffffffffffffffffffffffffffffffffffff166060870152565b606086013560808601526141f96141dc608088016108ea565b73ffffffffffffffffffffffffffffffffffffffff1660a0870152565b61420660a08701876140f5565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08784030160c0880152613630565b61424360c08601866140f5565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08684030160e0870152613630565b9260e08101906140f5565b916101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082860301910152613630565b908160209103126104195751610815816121b4565b90916142db610815936040845260408401906107a5565b9160208184039101526107a5565b35610815816108cc565b6142fb612c40565b5061430581614d00565b61431d61431560c0830183612c53565b81019061402e565b61434061433861433060e0850185612c53565b810190614066565b918251614f26565b73ffffffffffffffffffffffffffffffffffffffff61437a610e7860748451015173ffffffffffffffffffffffffffffffffffffffff1690565b7f000000000000000000000000000000000000000000000000000000000000000082169116811461453a5750602081816143e793519101519060405193849283927f57ecfd28000000000000000000000000000000000000000000000000000000008452600484016142c4565b0381600073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af190811561170f5760009161450b575b50156144e15761444560208201613037565b907ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc067ffffffffffffffff606061447e604085016142e9565b604080517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff908116825233602083015290921690820152930135606084018190529316918060808101612e80565b7fbf969f220000000000000000000000000000000000000000000000000000000060005260046000fd5b61452d915060203d602011614533575b61452581836106ee565b8101906142af565b38614433565b503d61451b565b6020915061457a906000936040519485809481937f3907753700000000000000000000000000000000000000000000000000000000835260048301614145565b03925af190811561170f57600091614590575090565b610815915060203d6020116145b2575b6145aa81836106ee565b8101906140d7565b503d6145a0565b73ffffffffffffffffffffffffffffffffffffffff6001541633036145da57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b60409067ffffffffffffffff610815949316815281602082015201906107a5565b90805115611fe0578051602082012067ffffffffffffffff83169283600052600760205261465a826005604060002001615206565b156146b35750816146a27f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9361469d6146ae946000526008602052604060002090565b613ddd565b60405191829182610804565b0390a2565b9050611cf86040519283927f393b8ad200000000000000000000000000000000000000000000000000000000845260048401614604565b67ffffffffffffffff1660008181526006602052604090205490929190156147ec57916147e960e0926147b5856147417f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b97614a78565b846000526007602052614758816040600020615262565b61476183614a78565b84600052600760205261477b836002604060002001615262565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b90816020910312610419573590565b9081602091031261041957516108158161041e565b611e366149049161484d6130b8565b5061485781615510565b60208101906060013561486e8235612d998161041e565b614879828254612c33565b90557ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff6148ae84613037565b604080517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681523360208201529081019490945216918060608101613337565b6040516133e3816133b760208201907ffa7c07de00000000000000000000000000000000000000000000000000000000602083019252565b6149446136a1565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff82511690602083019163ffffffff83511642034281116124e0576149a8906fffffffffffffffffffffffffffffffff60808701511690613a72565b81018091116124e0576149ce6fffffffffffffffffffffffffffffffff92918392615bd1565b161682524263ffffffff16905290565b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff9092166024830152604482019290925261043b916140298260648101613ffd565b61043b9092919260608101936fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b805115614b1c5760408101516fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff614adc614ac760208501516fffffffffffffffffffffffffffffffff1690565b6fffffffffffffffffffffffffffffffff1690565b911611614ae65750565b611cf8906040519182917f8020d12400000000000000000000000000000000000000000000000000000000835260048301614a3d565b6fffffffffffffffffffffffffffffffff614b4a60408301516fffffffffffffffffffffffffffffffff1690565b1615801590614b91575b614b5b5750565b611cf8906040519182917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301614a3d565b50614bb2614ac760208301516fffffffffffffffffffffffffffffffff1690565b1515614b54565b15614bc057565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b9073ffffffffffffffffffffffffffffffffffffffff614cd29216604090600080835194614c7285876106ee565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602087015260208151910182855af1903d15614cf7573d614cc3614cba8261076b565b945194856106ee565b83523d6000602085013e616052565b805180614cdd575050565b81602080614cf29361043b95010191016142af565b614bb9565b60609250616052565b9060808201614d14611c3e610917836142e9565b614ed857506020820191614db56020614d5a614d326104f187613037565b60801b7fffffffffffffffffffffffffffffffff000000000000000000000000000000001690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116600482015291829081906024820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561170f57600091614eb9575b50614e8f57614e15614e1084613037565b6159fc565b614e1e83613037565b614e33611c3e60a0840192610bbf8486612c53565b614e5057506060614e4761043b9394613037565b91013590615b20565b614e5991612c53565b90611cf86040519283927f24eb47e500000000000000000000000000000000000000000000000000000000845260048401613690565b7f53ad11d80000000000000000000000000000000000000000000000000000000060005260046000fd5b614ed2915060203d6020116145335761452581836106ee565b38614dff565b614ee46105b9916142e9565b7f961c9a4f0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b90600482015163ffffffff811661507e57506008820151916014600c82015191015192614f5a602084015163ffffffff1690565b63ffffffff811663ffffffff8316036150455750507f000000000000000000000000000000000000000000000000000000000000000063ffffffff811663ffffffff83160361500c5750505167ffffffffffffffff1667ffffffffffffffff811667ffffffffffffffff831603614fcf575050565b7ff917ffea0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff9081166004521660245260446000fd5b7f77e480260000000000000000000000000000000000000000000000000000000060005263ffffffff9081166004521660245260446000fd5b7fe366a1170000000000000000000000000000000000000000000000000000000060005263ffffffff9081166004521660245260446000fd5b7f68d2f8d60000000000000000000000000000000000000000000000000000000060005263ffffffff1660045260246000fd5b80548210156130325760005260206000200190600090565b60008181526003602052604090205461515457600254680100000000000000008110156106795761513b61510682600185940160025560026150b1565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600254906000526003602052604060002055600190565b50600090565b60008181526010602052604090205461515457600f546801000000000000000081101561067957615197615106826001859401600f55600f6150b1565b9055600f54906000526010602052604060002055600190565b6000818152600660205260409020546151545760055468010000000000000000811015610679576151ed61510682600185940160055560056150b1565b9055600554906000526006602052604060002055600190565b600082815260018201602052604090205461525b578054906801000000000000000082101561067957826152446151068460018096018555846150b1565b905580549260005201602052604060002055600190565b5050600090565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c199161544161378d9280546152b36152ad6152a48363ffffffff9060801c1690565b63ffffffff1690565b4261308f565b908161544d575b50506153fb60016152de60208601516fffffffffffffffffffffffffffffffff1690565b9261536961532c614ac76fffffffffffffffffffffffffffffffff61531385546fffffffffffffffffffffffffffffffff1690565b166fffffffffffffffffffffffffffffffff8816615bd1565b82906fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b6153bc6153768751151590565b82547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178255565b019182906fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b604083015181546fffffffffffffffffffffffffffffffff1660809190911b7fffffffffffffffffffffffffffffffff0000000000000000000000000000000016179055565b60405191829182614a3d565b614ac761532c916fffffffffffffffffffffffffffffffff6154c16154c895826154ba60018a015492826154b36154ac615496876fffffffffffffffffffffffffffffffff1690565b996fffffffffffffffffffffffffffffffff1690565b9560801c90565b1690613a72565b9116612c33565b9116615bd1565b80547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff000000000000000000000000000000001617815538806152ba565b60808101615523611c3e610917836142e9565b614ed8575060208101906155416020614d5a614d326104f186613037565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561170f576000916155c7575b50614e8f5760606155be61043b936155ad6155a8604086016142e9565b615be3565b6111306155b982613037565b615c7a565b91013590615d58565b6155e0915060203d6020116145335761452581836106ee565b3861558b565b604051906002548083528260208101600260005260206000209260005b81811061561857505061043b925003836106ee565b8454835260019485019487945060209093019201615603565b604051906005548083528260208101600560005260206000209260005b81811061566357505061043b925003836106ee565b845483526001948501948794506020909301920161564e565b906040519182815491828252602082019060005260206000209260005b8181106156ae57505061043b925003836106ee565b8454835260019485019487945060209093019201615699565b8054801561572b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906156fc82826150b1565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60008181526003602052604090205490811561525b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201908282116124e057600254927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84019384116124e05783836000956157f695036157fc575b5050506157e560026156c7565b600390600052602052604060002090565b55600190565b6157e56158289161581e61581461582e9560026150b1565b90549060031b1c90565b92839160026150b1565b90612eca565b553880806157d8565b60008181526006602052604090205490811561525b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201908282116124e057600554927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84019384116124e05783836000956157f695036158d3575b5050506158c260056156c7565b600690600052602052604060002090565b6158c2615828916158eb6158146158f59560056150b1565b92839160056150b1565b553880806158b5565b60018101918060005282602052604060002054928315156000146159d6577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84018481116124e0578354937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff85019485116124e05760009585836157f69761598e950361599d575b5050506156c7565b90600052602052604060002090565b6159bd615828916159b46158146159cd95886150b1565b928391876150b1565b8590600052602052604060002090565b55388080615986565b50505050600090565b926159ea9192613a72565b81018091116124e05761081591615bd1565b615a08611c3e82613041565b615ae9576020615a8191615a34610e7860045473ffffffffffffffffffffffffffffffffffffffff1690565b6040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90921660048301523360248301529092839190829081906044820190565b03915afa90811561170f57600091615aca575b5015615a9c57565b7f728fe07b000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b615ae3915060203d6020116145335761452581836106ee565b38615a94565b7fa9902c7e0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b67ffffffffffffffff7f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c91169182600052600760205280615ba1600260406000200173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615e0f565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252602082019290925290819081016146ae565b9080821015615bde575090565b905090565b7f0000000000000000000000000000000000000000000000000000000000000000615c0b5750565b73ffffffffffffffffffffffffffffffffffffffff1680600052600360205260406000205415615c385750565b7fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b908160209103126104195751610815816108cc565b615c86611c3e82613041565b615ae9576020615cf791615cb2610e7860045473ffffffffffffffffffffffffffffffffffffffff1690565b60405180809581947fa8d87a3b0000000000000000000000000000000000000000000000000000000083526004830191909167ffffffffffffffff6020820193169052565b03915afa801561170f5773ffffffffffffffffffffffffffffffffffffffff91600091615d29575b50163303615a9c57565b615d4b915060203d602011615d51575b615d4381836106ee565b810190615c65565b38615d1f565b503d615d39565b67ffffffffffffffff7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894491169182600052600760205280615ba1604060002073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615e0f565b8115615de0570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b8054939290919060ff60a086901c1615801561604a575b61604357615e456fffffffffffffffffffffffffffffffff8616614ac7565b9060018401958654615e7f6152ad6152a4615e72614ac7856fffffffffffffffffffffffffffffffff1690565b9460801c63ffffffff1690565b80615faf575b5050838110615f645750828210615ee5575061043b939450615eaa91614ac79161308f565b6fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b90615f1c6105b993615f17615f0884615f02614ac78c5460801c90565b9361308f565b615f1183613062565b90612c33565b615dd6565b7fd0c8d23a0000000000000000000000000000000000000000000000000000000060005260045260245273ffffffffffffffffffffffffffffffffffffffff16604452606490565b7f1a76572a00000000000000000000000000000000000000000000000000000000600052600452602483905273ffffffffffffffffffffffffffffffffffffffff1660445260646000fd5b82859293951161601957615fc9614ac7615fd09460801c90565b91856159df565b84547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178555913880615e85565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050509050565b508115615e26565b919290156160cd5750815115616066575090565b3b1561606f5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156160e05750805190602001fd5b611cf8906040519182917f08c379a00000000000000000000000000000000000000000000000000000000083526004830161080456fea164736f6c634300081a000a",
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
