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
	CctpVersion      uint8
	Enabled          bool
}

type USDCTokenPoolDomainUpdate struct {
	AllowedCaller     [32]byte
	MintRecipient     [32]byte
	DomainIdentifier  uint32
	DestChainSelector uint64
	CctpVersion       uint8
	Enabled           bool
}

var HybridLockReleaseUSDCTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"legacyTokenMessenger\",\"type\":\"address\",\"internalType\":\"contractITokenMessenger\"},{\"name\":\"tokenMessenger\",\"type\":\"address\",\"internalType\":\"contractITokenMessenger\"},{\"name\":\"cctpMessageTransmitterProxy\",\"type\":\"address\",\"internalType\":\"contractCCTPMessageTransmitterProxy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"previousPool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockBox\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"FINALITY_THRESHOLD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_FEE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnLockedUSDC\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelExistingCCTPMigrationProposal\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"excludeTokensFromBurn\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentProposedCCTPChainMigration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDomain\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structUSDCTokenPool.Domain\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"cctpVersion\",\"type\":\"uint8\",\"internalType\":\"enumUSDCTokenPool.CCTPVersion\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExcludedTokensByChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLiquidityProvider\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLockedTokensForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_legacyTokenMessenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITokenMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_localDomainIdentifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_lockBox\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_messageTransmitterProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractCCTPMessageTransmitterProxy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_previousMessageTransmitterProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_previousPool\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_supportedUSDCVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_tokenMessenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITokenMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposeCCTPMigration\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"provideLiquidity\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCircleMigratorAddress\",\"inputs\":[{\"name\":\"migrator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDomains\",\"inputs\":[{\"name\":\"domains\",\"type\":\"tuple[]\",\"internalType\":\"structUSDCTokenPool.DomainUpdate[]\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"cctpVersion\",\"type\":\"uint8\",\"internalType\":\"enumUSDCTokenPool.CCTPVersion\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLiquidityProvider\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"liquidityProvider\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRouter\",\"inputs\":[{\"name\":\"newRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shouldUseLockRelease\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateChainSelectorMechanisms\",\"inputs\":[{\"name\":\"removes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"adds\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawLiquidity\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCTPMigrationCancelled\",\"inputs\":[{\"name\":\"existingProposalSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCTPMigrationExecuted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"USDCBurned\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCTPMigrationProposed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CircleMigratorAddressSet\",\"inputs\":[{\"name\":\"migratorAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DomainsSet\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structUSDCTokenPool.DomainUpdate[]\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"cctpVersion\",\"type\":\"uint8\",\"internalType\":\"enumUSDCTokenPool.CCTPVersion\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityAdded\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityProviderSet\",\"inputs\":[{\"name\":\"oldProvider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newProvider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityRemoved\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockReleaseDisabled\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockReleaseEnabled\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RouterUpdated\",\"inputs\":[{\"name\":\"oldRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensExcludedFromBurn\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"burnableAmountAfterExclusion\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"ExistingMigrationProposal\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCCTPVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"enumUSDCTokenPool.CCTPVersion\"},{\"name\":\"got\",\"type\":\"uint8\",\"internalType\":\"enumUSDCTokenPool.CCTPVersion\"}]},{\"type\":\"error\",\"name\":\"InvalidChainSelector\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestinationDomain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidDomain\",\"inputs\":[{\"name\":\"domain\",\"type\":\"tuple\",\"internalType\":\"structUSDCTokenPool.DomainUpdate\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"cctpVersion\",\"type\":\"uint8\",\"internalType\":\"enumUSDCTokenPool.CCTPVersion\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"type\":\"error\",\"name\":\"InvalidExecutionFinalityThreshold\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMinFinalityThreshold\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidNonce\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"got\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidPreviousPool\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourceDomain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenMessengerVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidTransmitterInProxy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LanePausedForCCTPMigration\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"LockBoxCannotBeZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoMigrationProposalPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenLockingNotAllowedAfterMigration\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnknownDomain\",\"inputs\":[{\"name\":\"domain\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"UnlockingUSDCFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"onlyCircle\",\"inputs\":[]}]",
	Bin: "0x610220806040523461066257617a79803803809161001d8285610aea565b8339810190610120818303126106625761003681610b0d565b61004260208301610b0d565b9061004f60408401610b0d565b6060840151936001600160a01b03851693918486036106625760808201516001600160401b0381116106625782019680601f89011215610662578751976001600160401b038911610ad4578860051b9060208201996100b16040519b8c610aea565b8a526020808b019282010192831161066257602001905b828210610abc575050506100de60a08301610b0d565b966100eb60c08401610b0d565b6101046101006100fd60e08701610b0d565b9501610b0d565b983315610aab57600180546001600160a01b0319163317905587158015610a9a575b8015610a89575b610a7857608089905260c05260405163313ce56760e01b81526020816004818b5afa8091600091610a3c575b5090610a18575b50600660a052600480546001600160a01b0319166001600160a01b03929092169190911790558051151560e08190526108e7575b506001610100526001600160a01b0383169283156108d657604051632c12192160e01b8152602081600481885afa90811561066f5760009161089c575b5060405163054fd4d560e41b81526001600160a01b039190911690602081600481855afa90811561066f5760009161087d575b5063ffffffff80610100511691169081036108695750604051639cdbb18160e01b8152602081600481895afa90811561066f5760009161084a575b5063ffffffff806101005116911690810361083657506040516367e0ed8360e11b81526020816004816001600160a01b0388165afa801561066f5782916000916107e8575b506001600160a01b0316036107d75760049260209261012052610140526040519283809263234d8e3d60e21b82525afa90811561066f576000916107a8575b5061016052608051610120516102e5916001600160a01b039182169116610b55565b6001600160a01b0381169030821461072e5781151590818061073f575b61072e578115610723576040516398db964360e01b8152602081600481875afa600091816106e7575b5061034157632d4d3c3d60e21b60005260046000fd5b6001600160a01b03166101a0525b61018052600080516020617a598339815191526020604051858152a1156106c857906020600492604051938480926398db964360e01b82525afa6000928161068c575b506103a857632d4d3c3d60e21b60005260046000fd5b6001600160a01b039091166101a052600080516020617a5983398151915291602091905b6080516103e5906001600160a01b038381169116610b55565b6101c052604051908152a16101e08190526102008390526001600160a01b03831690811561067b57604051636eb1769f60e11b81523060048201526001600160a01b038516602482015290602090829060449082905afa90811561066f57600091610638575b506105cd57610486916040519163095ea7b360e01b60208401526024830152600019604483015260448252610481606483610aea565b610c5c565b61020052604051616ab39081610fa68239608051818181610586015281816109d501528181610a5c0152818161375601528181613ab901528181613b5c01528181613ddf0152818161522d01528181616509015261673e015260a05181610aa6015260c051818181612769015281816156560152615eea015260e051818181610fb1015281816129ad015261657501526101005181818161098401526157e601526101205181818161119a0152613d44015261014051818181611beb0152614c5b0152610160518181816112d301528181613bd90152615851015261018051818181610bde0152614ba501526101a051818181611ec30152614ddf01526101c0518181816121110152613b1901526101e051816118b40152610200518181816105c1015281816118180152818161221d01528181612f600152818161364b015261514b0152f35b60405162461bcd60e51b815260206004820152603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527f20746f206e6f6e2d7a65726f20616c6c6f77616e6365000000000000000000006064820152608490fd5b90506020813d602011610667575b8161065360209383610aea565b8101031261066257513861044b565b600080fd5b3d9150610646565b6040513d6000823e3d90fd5b6321a9a7af60e01b60005260046000fd5b9092506020813d6020116106c0575b816106a860209383610aea565b81010312610662576106b990610b0d565b9138610392565b3d915061069b565b50602090600080516020617a598339815191529260006101a0526103cc565b9091506020813d60201161071b575b8161070360209383610aea565b810103126106625761071490610b0d565b903861032b565b3d91506106f6565b60006101a05261034f565b632d4d3c3d60e21b60005260046000fd5b506040516301ffc9a760e01b8152630e64dd2960e01b6004820152602081602481875afa90811561066f57600091610779575b5015610302565b61079b915060203d6020116107a1575b6107938183610aea565b810190610b3d565b38610772565b503d610789565b6107ca915060203d6020116107d0575b6107c28183610aea565b810190610b21565b386102c3565b503d6107b8565b632a32133b60e11b60005260046000fd5b9091506020813d60201161082e575b8161080460209383610aea565b8101031261082a5751906001600160a01b03821682036108275750819038610284565b80fd5b5080fd5b3d91506107f7565b6316ba39c560e31b60005260045260246000fd5b610863915060203d6020116107d0576107c28183610aea565b3861023f565b6334697c6b60e11b60005260045260246000fd5b610896915060203d6020116107d0576107c28183610aea565b38610204565b90506020813d6020116108ce575b816108b760209383610aea565b81010312610662576108c890610b0d565b386101d1565b3d91506108aa565b6306b7c75960e31b60005260046000fd5b919590929394602095604051956108fe8888610aea565b60008752600036813760e05115610a075760005b8751811015610979576001906001600160a01b03610930828b610c32565b51168a61093c82610d8d565b610949575b505001610912565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a1388a610941565b509193955091939660005b86518110156109f8576001906001600160a01b036109a2828a610c32565b511680156109f257896109b482610e75565b6109c2575b50505b01610984565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a138896109b9565b506109bc565b50929550929590935038610194565b6335f4a7b360e01b60005260046000fd5b60ff1660068114610160576332ad3e0760e11b600052600660045260245260446000fd5b6020813d602011610a70575b81610a5560209383610aea565b8101031261082a57519060ff82168203610827575038610159565b3d9150610a48565b6342bcdf7f60e11b60005260046000fd5b506001600160a01b0381161561012d565b506001600160a01b03821615610126565b639b15e16f60e01b60005260046000fd5b60208091610ac984610b0d565b8152019101906100c8565b634e487b7160e01b600052604160045260246000fd5b601f909101601f19168101906001600160401b03821190821017610ad457604052565b51906001600160a01b038216820361066257565b90816020910312610662575163ffffffff811681036106625790565b90816020910312610662575180151581036106625790565b604051636eb1769f60e11b81523060048201526001600160a01b03831660248201529190602083806044810103816001600160a01b0385165afa92831561066f57600093610bfe575b506000198301809311610be85760405163095ea7b360e01b60208201526001600160a01b039092166024830152604480830193909352918152610be691610481606483610aea565b565b634e487b7160e01b600052601160045260246000fd5b90926020823d602011610c2a575b81610c1960209383610aea565b810103126108275750519138610b9e565b3d9150610c0c565b8051821015610c465760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b60018060a01b0316604091600080845192610c778685610aea565b602084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602085015260208151910182865af13d15610d68573d906001600160401b038211610ad4578451610cec949092610cdd601f8201601f191660200185610aea565b83523d6000602085013e610ed5565b80519081610cf957505050565b602080610d0a938301019101610b3d565b15610d125750565b5162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b91610cec92606091610ed5565b8054821015610c465760005260206000200190600090565b6000818152600360205260409020548015610e6e576000198101818111610be857600254600019810191908211610be857818103610e1d575b5050506002548015610e075760001901610de1816002610d75565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b610e56610e2e610e3f936002610d75565b90549060031b1c9283926002610d75565b819391549060031b91821b91600019901b19161790565b90556000526003602052604060002055388080610dc6565b5050600090565b80600052600360205260406000205415600014610ecf5760025468010000000000000000811015610ad457610eb6610e3f8260018594016002556002610d75565b9055600254906000526003602052604060002055600190565b50600090565b91929015610f375750815115610ee9575090565b3b15610ef25790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b825190915015610f4a5750805190602001fd5b6040519062461bcd60e51b8252602060048301528181519182602483015260005b838110610f8d5750508160006044809484010152601f80199101168101030190fd5b60208282018101516044878401015285935001610f6b56fe6080604052600436101561001257600080fd5b60003560e01c806301ffc9a7146103975780631101dbd414610392578063181f5a771461038d578063212a052e1461038857806321df0da714610383578063240028e81461037e57806324f65ee71461037957806324f8795a146103745780632cfbb1191461036f578063390775371461036a5780633e591f2c146103655780634ad01f0b146103605780634c5ef0ed1461035b5780634c93ef841461035657806350d1a35a1461035157806354c8a4f31461034c5780636155cda01461034757806362ddd3c4146103425780636b716b0d1461033d5780636b795423146103385780636d3d1a5814610333578063714bf9071461032e57806379ba5097146103295780637d54534e146103245780638926f54f1461031f5780638a5e52bb1461031a5780638da5cb5b14610315578063962d40201461031057806398db96431461030b5780639a4575b914610306578063a42a7b8b14610301578063a7cd63b7146102fc578063a7eff066146102f7578063acfecf91146102f2578063af58d59f146102ed578063b0f479a1146102e8578063b2278077146102e3578063b7946580146102de578063bb5eced3146102d9578063bc063e1a146102d4578063c0d78655146102cf578063c4bffe2b146102ca578063c75eea9c146102c5578063cd306a6c146102c0578063cf7401f3146102bb578063da4b05e7146102b6578063dc0bd971146102b1578063de814c57146102ac578063dfadfa35146102a7578063e0351e13146102a2578063e8a1da171461029d578063e94ae6d014610298578063f2fde38b14610293578063f65a88861461028e578063f831bbcd146102895763fd6768551461028457600080fd5b612f84565b612f33565b612ef4565b612e1e565b612dc9565b6129d2565b612995565b6128ff565b61278d565b61273c565b61271f565b612621565b6124fc565b6124b1565b61241e565b6122fc565b6122e0565b612189565b61214d565b6120fb565b6120c7565b61201b565b611ee7565b611e96565b611e22565b611d20565b611c0f565b611bbe565b611a56565b6119f1565b611713565b6116d4565b611643565b611578565b6114e7565b6114b3565b6112f7565b6112b6565b611233565b611184565b610f7f565b610db6565b610d76565b610d2d565b610c02565b610bb1565b610b68565b610b29565b610aca565b610a8c565b610a22565b6109a8565b610967565b610904565b61048d565b34610469576020600319360112610469576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361046957807faff2afbf000000000000000000000000000000000000000000000000000000006020921490811561043f575b8115610415575b506040519015158152f35b7f01ffc9a7000000000000000000000000000000000000000000000000000000009150143861040a565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150610403565b600080fd5b67ffffffffffffffff81160361046957565b359061048b8261046e565b565b34610469576040600319360112610469576004356104aa8161046e565b602435906104e96104cf8267ffffffffffffffff166000526011602052604060002090565b5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff339116036106fc5767ffffffffffffffff8116610528816000526010602052604060002054151590565b6106c457600b5461054e9060a01c67ffffffffffffffff165b67ffffffffffffffff1690565b14610689576105718167ffffffffffffffff16600052600c602052604060002090565b61057c83825461306b565b90556105aa8230337f00000000000000000000000000000000000000000000000000000000000000006146a5565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803b15610469576040517f7d552ea60000000000000000000000000000000000000000000000000000000081526004810184905267ffffffffffffffff9290921660248301526000908290604490829084905af1801561068457610669575b50337fc17cea59c2955cb181b03393209566960365771dbba9dc3d510180e7cb312088600080a3005b80610678600061067e936107d9565b8061072a565b38610640565b613078565b7fd0da86c40000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b6000fd5b7f646972460000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b7f8e4a23d6000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b600091031261046957565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6060810190811067ffffffffffffffff82111761078057604052565b610735565b60a0810190811067ffffffffffffffff82111761078057604052565b6020810190811067ffffffffffffffff82111761078057604052565b6040810190811067ffffffffffffffff82111761078057604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761078057604052565b6040519061048b60a0836107d9565b6040519061048b6020836107d9565b6040519061048b610140836107d9565b6040519061048b6040836107d9565b67ffffffffffffffff811161078057601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b8481106108db5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b8060208092840101518282860101520161089c565b906020610901928181520190610891565b90565b3461046957600060031936011261046957610963604080519061092781836107d9565b601782527f55534443546f6b656e506f6f6c20312e362e312d646576000000000000000000602083015251918291602083526020830190610891565b0390f35b3461046957600060031936011261046957602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461046957600060031936011261046957602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b73ffffffffffffffffffffffffffffffffffffffff81160361046957565b359061048b826109f9565b34610469576020600319360112610469576020610a82600435610a44816109f9565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691161490565b6040519015158152f35b3461046957600060031936011261046957602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346104695760206003193601126104695760043567ffffffffffffffff8111610469573660238201121561046957806004013567ffffffffffffffff81116104695736602460c0830284010111610469576024610b279201613084565b005b346104695760206003193601126104695767ffffffffffffffff600435610b4f8161046e565b16600052600c6020526020604060002054604051908152f35b346104695760206003193601126104695760043567ffffffffffffffff811161046957610100600319823603011261046957610ba8602091600401613511565b60405190518152f35b3461046957600060031936011261046957602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461046957600060031936011261046957610c1b6147b1565b600b5467ffffffffffffffff8160a01c168015610cb1577fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff7f375f1ad1194a2bec317c5efec05cc63ffa06ddd0c4b276619f6fd47298eda5189216600b556000610c998267ffffffffffffffff16600052600d602052604060002090565b5560405167ffffffffffffffff9091168152602090a1005b7fa94cb9880000000000000000000000000000000000000000000000000000000060005260046000fd5b929192610ce782610857565b91610cf560405193846107d9565b829481845281830111610469578281602093846000960137010152565b9080601f830112156104695781602061090193359101610cdb565b3461046957604060031936011261046957600435610d4a8161046e565b60243567ffffffffffffffff811161046957602091610d70610a82923690600401610d12565b90613830565b34610469576020600319360112610469576020610a82600435610d988161046e565b67ffffffffffffffff16600052600e60205260ff6040600020541690565b3461046957602060031936011261046957600435610dd38161046e565b610ddb6147b1565b67ffffffffffffffff600b5460a01c16610ed25767ffffffffffffffff81166000908152600e602052604090205460ff1615610ea857610ea381610e887f20331f191af84dbff48b162aa5a5985e7891ae646297b0a2ac80487f9109ef49937fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff7bffffffffffffffff0000000000000000000000000000000000000000600b549260a01b16911617600b55565b60405167ffffffffffffffff90911681529081906020820190565b0390a1005b7f656535ce0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f692bc1310000000000000000000000000000000000000000000000000000000060005260046000fd5b9181601f840112156104695782359167ffffffffffffffff8311610469576020808501948460051b01011161046957565b60406003198201126104695760043567ffffffffffffffff81116104695781610f5891600401610efc565b929092916024359067ffffffffffffffff821161046957610f7b91600401610efc565b9091565b3461046957610fa7610faf610f9336610f2d565b9491610fa09391936147b1565b3691613885565b923691613885565b7f00000000000000000000000000000000000000000000000000000000000000001561115a5760005b825181101561109d578061100b610ff160019386613eed565b5173ffffffffffffffffffffffffffffffffffffffff1690565b61104761104273ffffffffffffffffffffffffffffffffffffffff83165b73ffffffffffffffffffffffffffffffffffffffff1690565b6160ea565b611053575b5001610fd8565b60405173ffffffffffffffffffffffffffffffffffffffff9190911681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756690602090a13861104c565b5060005b8151811015610b2757806110ba610ff160019385613eed565b73ffffffffffffffffffffffffffffffffffffffff811615611154576110fd6110f873ffffffffffffffffffffffffffffffffffffffff8316611029565b615a59565b61110a575b505b016110a1565b60405173ffffffffffffffffffffffffffffffffffffffff9190911681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d890602090a183611102565b50611104565b7f35f4a7b30000000000000000000000000000000000000000000000000000000060005260046000fd5b34610469576000600319360112610469576040517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168152602090f35b6040600319820112610469576004356111ee8161046e565b9160243567ffffffffffffffff811161046957826023820112156104695780600401359267ffffffffffffffff84116104695760248483010111610469576024019190565b3461046957611241366111d6565b61124c9291926147b1565b67ffffffffffffffff821661126e816000526006602052604060002054151590565b156112895750610b2792611283913691610cdb565b90614e29565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b3461046957600060031936011261046957602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346104695761130536610f2d565b9290916113106147b1565b60005b8181106114385750505060005b82811061132957005b61135561134261054161133d8487876138db565b6138eb565b6000526010602052604060002054151590565b6113f157806113b561138a61137061133d60019588886138db565b67ffffffffffffffff16600052600e602052604060002090565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055565b6113c661054161133d8387876138db565b7f5e3985e51df58346365017cae614e59d723143b71c9a2ce4a156687f1f2c3f5a600080a201611320565b61133d906106c093611402936138db565b7f646972460000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b8061147761144f61137061133d60019587896138db565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008154169055565b61148861054161133d8386886138db565b7fddc5afbc5e53c63a556964db0eef76a1c2d9305e0811abd7410d2a6f4799490e600080a201611313565b3461046957600060031936011261046957602073ffffffffffffffffffffffffffffffffffffffff60095416604051908152f35b34610469576020600319360112610469577f084e6f0e9791c2e56153bd49e6ec6dd63ba9a72c258d71558d74c63fc75b7168602073ffffffffffffffffffffffffffffffffffffffff60043561153c816109f9565b6115446147b1565b16807fffffffffffffffffffffffff0000000000000000000000000000000000000000600b541617600b55604051908152a1005b346104695760006003193601126104695760005473ffffffffffffffffffffffffffffffffffffffff81163303611619577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b34610469576020600319360112610469577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174602073ffffffffffffffffffffffffffffffffffffffff600435611698816109f9565b6116a06147b1565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000006009541617600955604051908152a1005b34610469576020600319360112610469576020610a8267ffffffffffffffff6004356116ff8161046e565b166000526006602052604060002054151590565b3461046957600060031936011261046957600b5461174673ffffffffffffffffffffffffffffffffffffffff8216611029565b33036119c75760a01c67ffffffffffffffff1667ffffffffffffffff8116908115610cb1576117b261178c8267ffffffffffffffff16600052600c602052604060002090565b546117ab8367ffffffffffffffff16600052600d602052604060002090565b5490613943565b9060006117d38267ffffffffffffffff16600052600c602052604060002090565b556118017fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff600b5416600b55565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803b15610469576040517f22ffa47a0000000000000000000000000000000000000000000000000000000081526004810184905230602482015267ffffffffffffffff83166044820152906000908290606490829084905af18015610684576119b2575b5073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692833b1561046957600060405180957f42966c6800000000000000000000000000000000000000000000000000000000825281838161191b89600483019190602083019252565b03925af1908115610684577fdea60ddd4c7ebdab804f5694c70350cca7893ece3efeecb142312eacac5c73e4946119789261199d575b5061197361144f8467ffffffffffffffff16600052600e602052604060002090565b615aea565b506040805167ffffffffffffffff909216825260208201929092529081908101610ea3565b8061067860006119ac936107d9565b38611951565b8061067860006119c1936107d9565b3861189c565b7f438a7a050000000000000000000000000000000000000000000000000000000060005260046000fd5b3461046957600060031936011261046957602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b9181601f840112156104695782359167ffffffffffffffff8311610469576020808501946060850201011161046957565b346104695760606003193601126104695760043567ffffffffffffffff811161046957611a87903690600401610efc565b9060243567ffffffffffffffff811161046957611aa8903690600401611a25565b9060443567ffffffffffffffff811161046957611ac9903690600401611a25565b611aeb61102960095473ffffffffffffffffffffffffffffffffffffffff1690565b33141580611b93575b6106fc57838614801590611b89575b611b5f5760005b868110611b1357005b80611b59611b2761133d6001948b8b6138db565b611b32838989613950565b611b53611b4b611b4386898b613950565b9236906125d8565b9136906125d8565b91614eee565b01611b0a565b7f568efce20000000000000000000000000000000000000000000000000000000060005260046000fd5b5080861415611b03565b50611bb661102960015473ffffffffffffffffffffffffffffffffffffffff1690565b331415611af4565b3461046957600060031936011261046957602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346104695760206003193601126104695760043567ffffffffffffffff81116104695760a0600319823603011261046957611c4f61096391600401613979565b604051918291602083526020611c7082516040838701526060860190610891565b9101517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0848303016040850152610891565b602081016020825282518091526040820191602060408360051b8301019401926000915b838310611cd557505050505090565b9091929394602080611d11837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc086600196030187528951610891565b97019301930191939290611cc6565b346104695760206003193601126104695767ffffffffffffffff600435611d468161046e565b166000526007602052611d5f600560406000200161600c565b8051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611da5611d8f8461386d565b93611d9d60405195866107d9565b80855261386d565b0160005b818110611e1157505060005b8151811015611e035780611de7611de2611dd160019486613eed565b516000526008602052604060002090565b613f54565b611df18286613eed565b52611dfc8185613eed565b5001611db5565b604051806109638582611ca2565b806060602080938701015201611da9565b3461046957600060031936011261046957611e3b615f76565b60405180916020820160208352815180915260206040840192019060005b818110611e67575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101611e59565b3461046957600060031936011261046957602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461046957611ef5366111d6565b611f009291926147b1565b67ffffffffffffffff821691611f2a611f26846000526006602052604060002054151590565b1590565b611fe457611f6d611f266005611f548467ffffffffffffffff166000526007602052604060002090565b01611f60368689610cdb565b602081519101209061628e565b611fa957507f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d769192611fa460405192839283614074565b0390a2005b611fe084926040519384937f74f23c7c00000000000000000000000000000000000000000000000000000000855260048501614053565b0390fd5b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b346104695760206003193601126104695767ffffffffffffffff6004356120418161046e565b612049614085565b5016600052600760205261096361206e61206960026040600020016140b0565b6152e0565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b3461046957600060031936011261046957602073ffffffffffffffffffffffffffffffffffffffff60045416604051908152f35b34610469576000600319360112610469576040517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168152602090f35b34610469576020600319360112610469576109636121756004356121708161046e565b61410a565b604051918291602083526020830190610891565b34610469576040600319360112610469576004356121a68161046e565b602435906121b26147b1565b600b546121cb9060a01c67ffffffffffffffff16610541565b67ffffffffffffffff821614610689576121f98167ffffffffffffffff16600052600c602052604060002090565b612204838254613943565b905573ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803b15610469576040517f22ffa47a0000000000000000000000000000000000000000000000000000000081526004810184905233602482015267ffffffffffffffff9290921660448301526000908290606490829084905af18015610684576122cb575b50337fc2c3f06e49b9f15e7b4af9055e183b0d73362e033ad82a07dec9bf9840171719600080a3005b8061067860006122da936107d9565b386122a2565b3461046957600060031936011261046957602060405160008152f35b346104695760206003193601126104695773ffffffffffffffffffffffffffffffffffffffff60043561232e816109f9565b6123366147b1565b1680156123b05760407f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f16849160045490807fffffffffffffffffffffffff000000000000000000000000000000000000000083161760045573ffffffffffffffffffffffffffffffffffffffff8351921682526020820152a1005b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b602060408183019282815284518094520192019060005b8181106123fe5750505090565b825167ffffffffffffffff168452602093840193909201916001016123f1565b3461046957600060031936011261046957612437615fc1565b8051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612467611d8f8461386d565b0136602084013760005b81518110156124a3578067ffffffffffffffff61249060019385613eed565b511661249c8286613eed565b5201612471565b6040518061096385826123da565b346104695760206003193601126104695767ffffffffffffffff6004356124d78161046e565b6124df614085565b5016600052600760205261096361206e61206960406000206140b0565b3461046957600060031936011261046957602067ffffffffffffffff600b5460a01c16604051908152f35b8015150361046957565b35906fffffffffffffffffffffffffffffffff8216820361046957565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c6060910112610469576040519061258582610764565b8160843561259281612527565b815260a4356fffffffffffffffffffffffffffffffff8116810361046957602082015260c435906fffffffffffffffffffffffffffffffff821682036104695760400152565b9190826060910312610469576040516125f081610764565b604061261c818395803561260381612527565b855261261160208201612531565b602086015201612531565b910152565b346104695760e06003193601126104695760043561263e8161046e565b60607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc3601126104695760405161267481610764565b60243561268081612527565b81526044356fffffffffffffffffffffffffffffffff811681036104695760208201526064356fffffffffffffffffffffffffffffffff811681036104695760408201526126cd3661254e565b9073ffffffffffffffffffffffffffffffffffffffff60095416331415806126fd575b6106fc57610b2792614eee565b5073ffffffffffffffffffffffffffffffffffffffff600154163314156126f0565b346104695760006003193601126104695760206040516107d08152f35b3461046957600060031936011261046957602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610469576040600319360112610469576004356127aa8161046e565b602435906127b66147b1565b67ffffffffffffffff600b5460a01c169167ffffffffffffffff8216809303610cb15782600052600d6020526040600020805492828401809411612870577fe1e6c22ce6b566f66cdb457ec2e7910ff1f9a9e5654ed75303476fa8704682209361285892556117ab61283c8267ffffffffffffffff16600052600c602052604060002090565b549167ffffffffffffffff16600052600d602052604060002090565b60408051928352602083019190915281908101611fa4565b61303c565b6003111561287f57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b90600382101561287f5752565b91909160808060a0830194805184526020810151602085015263ffffffff60408201511660408501526128f6606082015160608601906128ae565b01511515910152565b346104695760206003193601126104695767ffffffffffffffff6004356129258161046e565b61292d614085565b5016600052600a602052610963604060002060ff60026040519261295084610785565b8054845260018101546020850152015463ffffffff8116604084015261297e828260201c1660608501613328565b60281c1615156080820152604051918291826128bb565b346104695760006003193601126104695760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b34610469576129e036610f2d565b9190926129eb6147b1565b6000915b808310612c905750505060009163ffffffff4216925b828110612a0e57005b612a21612a1c828585614287565b614346565b9060608201612a3081516153bd565b6080830193612a3f85516153bd565b6040840190815151156123b057612a6c611f26612a67610541885167ffffffffffffffff1690565b615b40565b612c4557612ba5612aa5612a8b879a999a5167ffffffffffffffff1690565b67ffffffffffffffff166000526007602052604060002090565b612b6889612b628751612b49612ace60408301516fffffffffffffffffffffffffffffffff1690565b91612b30612af9612af260208401516fffffffffffffffffffffffffffffffff1690565b9251151590565b612b27612b0461081a565b6fffffffffffffffffffffffffffffffff851681529763ffffffff166020890152565b15156040870152565b6fffffffffffffffffffffffffffffffff166060850152565b6fffffffffffffffffffffffffffffffff166080830152565b826143dd565b612b9a89612b918a51612b49612ace60408301516fffffffffffffffffffffffffffffffff1690565b600283016143dd565b6004845191016144e9565b602085019660005b88518051821015612be85790612be2600192612bdb83612bd58c5167ffffffffffffffff1690565b92613eed565b5190614e29565b01612bad565b50509796509490612c3c7f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c29392612c296001975167ffffffffffffffff1690565b9251935190519060405194859485614610565b0390a101612a05565b6106c0612c5a865167ffffffffffffffff1690565b7f1d5ad3c50000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b909192612ca161133d8584866138db565b94612cb8611f2667ffffffffffffffff88166161c7565b612d9157612ce56005612cdf8867ffffffffffffffff166000526007602052604060002090565b0161600c565b9360005b8551811015612d3157600190612d2a6005612d188b67ffffffffffffffff166000526007602052604060002090565b01612d23838a613eed565b519061628e565b5001612ce9565b509350937f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916612d8360019397610e88612d7e8267ffffffffffffffff166000526007602052604060002090565b6141d8565b0390a10191909392936129ef565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff861660045260246000fd5b346104695760206003193601126104695767ffffffffffffffff600435612def8161046e565b166000526011602052602073ffffffffffffffffffffffffffffffffffffffff60406000205416604051908152f35b346104695760206003193601126104695773ffffffffffffffffffffffffffffffffffffffff600435612e50816109f9565b612e586147b1565b16338114612eca57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346104695760206003193601126104695767ffffffffffffffff600435612f1a8161046e565b16600052600d6020526020604060002054604051908152f35b3461046957600060031936011261046957602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461046957604060031936011261046957600435612fa18161046e565b67ffffffffffffffff60243591612fb7836109f9565b612fbf6147b1565b166000818152601160205260408120805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff0000000000000000000000000000000000000000821681179092559293909216907fc82aa48e67c70b1ad1494533456f52504bb4d62d11bbdafaeb98cfccd1ed817e9080a4005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9190820180921161287057565b6040513d6000823e3d90fd5b61308c6147b1565b60005b8281106130ce5750907f802f1e7749437f70753a318fa838e9cd1afc1c06a032818e15050bfc4f01b424916130c96040519283928361340b565b0390a1565b6130e16130dc828585613224565b613257565b80511580156131cf575b61319957906131938261318e61317460606001965193613165602082015161315c61311d604085015163ffffffff1690565b61315460808601519161312f83612875565b60a087015115159461313f61081a565b9b8c5260208c015263ffffffff1660408b0152565b858901613328565b15156080870152565b015167ffffffffffffffff1690565b67ffffffffffffffff16600052600a602052604060002090565b613334565b0161308f565b611fe0906040519182917fa7a9337d000000000000000000000000000000000000000000000000000000008352600483016132d9565b5067ffffffffffffffff6131ee606083015167ffffffffffffffff1690565b16156130eb565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b91908110156132345760c0020190565b6131f5565b359063ffffffff8216820361046957565b3590600382101561046957565b60c081360312610469576040519060c082019082821067ffffffffffffffff8311176107805760a091604052803583526020810135602084015261329d60408201613239565b604084015260608101356132b08161046e565b60608401526132c16080820161324a565b608084015201356132d181612527565b60a082015290565b91909160a08060c0830194805184526020810151602085015263ffffffff604082015116604085015267ffffffffffffffff60608201511660608501526128f6608082015160808601906128ae565b600382101561287f5752565b60029082518155602083015160018201550163ffffffff6040830151167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416178155606082015191600383101561287f57600383101561287f5760806133d49161048b947fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff64ff0000000086549260201b1691161784550151151590565b81547fffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffffff1690151560281b65ff000000000016179055565b602080825281018390526040019160005b8181106134295750505090565b90919260c080600192863581526020870135602082015263ffffffff61345160408901613239565b16604082015267ffffffffffffffff606088013561346e8161046e565b16606082015261348d6134836080890161324a565b60808301906128ae565b60a087013561349b81612527565b151560a082015201940192910161341c565b604051906134ba826107a1565b60008252565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610469570180359067ffffffffffffffff82116104695760200191813603831361046957565b60405161351d816107a1565b600090527ffa7c07de000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000061357060c08401846134c0565b903582811691600481106137df575b505016036137d65761358f6134ad565b506060810135906135a08282615589565b600b5460a01c67ffffffffffffffff169060208101916135c2610541846138eb565b67ffffffffffffffff82161461068957506135f96135df836138eb565b67ffffffffffffffff16600052600c602052604060002090565b546137b85761362461360a836138eb565b67ffffffffffffffff16600052600d602052604060002090565b61362f848254613943565b90555b604073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691019161367783614b2e565b613680826138eb565b833b15610469576040517f22ffa47a0000000000000000000000000000000000000000000000000000000081526004810187905273ffffffffffffffffffffffffffffffffffffffff92909216602483015267ffffffffffffffff166044820152916000908390606490829084905af19182156106845761373567ffffffffffffffff9261373b927ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0956137a3575b506138eb565b93614b2e565b6040805173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000081168252336020830152929092169082015260608101859052921691608090a261379d610829565b90815290565b8061067860006137b2936107d9565b3861372f565b6137c46135df836138eb565b6137cf848254613943565b9055613632565b61090190614b38565b839250829060040360031b1b1616388061357f565b9161382c918354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055565b9067ffffffffffffffff61090192166000526007602052600560406000200190602081519101209060019160005201602052604060002054151590565b67ffffffffffffffff81116107805760051b60200190565b9291906138918161386d565b9361389f60405195866107d9565b602085838152019160051b810192831161046957905b8282106138c157505050565b6020809183356138d0816109f9565b8152019101906138b5565b91908110156132345760051b0190565b356109018161046e565b67ffffffffffffffff61090191166000526006602052604060002054151590565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161287057565b9190820391821161287057565b9190811015613234576060020190565b6040519061396d826107bd565b60606020838281520152565b613981613960565b50602081016139948135610d988161046e565b156139d6576139bc6105416139b6600b5467ffffffffffffffff9060a01c1690565b926138eb565b67ffffffffffffffff821614610689575061090190615100565b6139de613960565b506139e882615ea0565b6139fc6139f7613174836138eb565b61412c565b91613a0d611f266080850151151590565b613eab576020613a1d82806134c0565b905003613e6b5760208301518015613e4f57905b6000806060860160018151613a4581612875565b613a4e81612875565b03613d08575050506001916060820135602082613a72604089015163ffffffff1690565b88516040517ff856ddb6000000000000000000000000000000000000000000000000000000008152600481019590955263ffffffff909116602485015260448401919091527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1660648401526084830152818060a481010381600073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af190811561068457600091613cd9575b50945b613b56856138eb565b604080517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1680825233602083015260609687013592820183905295919267ffffffffffffffff16917ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1091a27f0000000000000000000000000000000000000000000000000000000000000000915194613c03610838565b67ffffffffffffffff909816885263ffffffff83166020890152613c2a9060408901613328565b606087015263ffffffff16608086015260a085015273ffffffffffffffffffffffffffffffffffffffff1660c084015260e083015260006101008301526107d0610120830152613c79906138eb565b613c829061410a565b9060405180916020820190613c9691615042565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018252613cc690826107d9565b613cce610848565b918252602082015290565b613cfb915060203d602011613d01575b613cf381836107d9565b81019061502d565b38613b4a565b503d613ce9565b600290969194929651613d1a81612875565b613d2381612875565b03613b4d57925060029273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166060840135613d7a604084015163ffffffff1690565b90835192803b15610469576040517fd04857b0000000000000000000000000000000000000000000000000000000008152600481019290925263ffffffff9290921660248201526044810185905273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001660648201526084810192909252600060a483018190526107d060c484015290829060e490829084905af1801561068457613e3a575b50613b4d565b806106786000613e49936107d9565b38613e34565b50613e65613e5d82806134c0565b81019061501e565b90613a31565b80613e75916134c0565b90611fe06040519283927fa3c8cf0900000000000000000000000000000000000000000000000000000000845260048401614074565b6106c0613eb7836138eb565b7fd201c48a0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b80518210156132345760209160051b010190565b90600182811c92168015613f4a575b6020831014613f1b57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691613f10565b9060405191826000825492613f6884613f01565b8084529360018116908115613fd45750600114613f8d575b5061048b925003836107d9565b90506000929192526020600020906000915b818310613fb857505090602061048b9282010138613f80565b6020919350806001915483858901015201910190918492613f9f565b6020935061048b9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138613f80565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b60409067ffffffffffffffff61090195931681528160208201520191614014565b916020610901938181520191614014565b6040519061409282610785565b60006080838281528260208201528260408201528260608201520152565b906040516140bd81610785565b60806fffffffffffffffffffffffffffffffff6001839560ff8154848116875263ffffffff81871c16602088015260a01c1615156040860152015481808216166060850152821c16910152565b67ffffffffffffffff1660005260076020526109016004604060002001613f54565b9060405161413981610785565b608060ff600283958054855260018101546020860152015463ffffffff8116604085015261416f828260201c1660608601613328565b60281c161515910152565b818110614185575050565b6000815560010161417a565b8181029291811591840414171561287057565b80549060008155816141b4575050565b6000526020600020908101905b8181106141cc575050565b600081556001016141c1565b600561048b9160008155600060018201556000600282015560006003820155600481016142058154613f01565b9081614214575b5050016141a4565b81601f6000931160011461422c5750555b388061420c565b8183526020832061424791601f01861c81019060010161417a565b808252602082209081548360011b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8560031b1c191617905555614225565b91908110156132345760051b810135907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee181360301821215610469570190565b9080601f830112156104695781356142de8161386d565b926142ec60405194856107d9565b81845260208085019260051b820101918383116104695760208201905b83821061431857505050505090565b813567ffffffffffffffff81116104695760209161433b87848094880101610d12565b815201910190614309565b61012081360312610469576040519061435e82610785565b61436781610480565b8252602081013567ffffffffffffffff81116104695761438a90369083016142c7565b602083015260408101359067ffffffffffffffff8211610469576143b46143d59236908301610d12565b60408401526143c636606083016125d8565b606084015260c03691016125d8565b608082015290565b8151815460208401516040850151608091821b73ffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffff000000000000000000000000000000000000000000909416939093179290921791151560a01b74ff000000000000000000000000000000000000000016919091178355606084015193810151901b7fffffffffffffffffffffffffffffffff0000000000000000000000000000000016921691909117600190910155565b9190601f81116144b357505050565b61048b926000526020600020906020601f840160051c830193106144df575b601f0160051c019061417a565b90915081906144d2565b919091825167ffffffffffffffff8111610780576145118161450b8454613f01565b846144a4565b6020601f821160011461456b57819061382c939495600092614560575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b01519050388061452e565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169061459e84600052602060002090565b9160005b8181106145f8575095836001959697106145c1575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880806145b7565b9192602060018192868b0151815501940192016145a2565b61467461463f61048b9597969467ffffffffffffffff60a0951684526101006020850152610100840190610891565b9660408301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b01906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b6040517f23b872dd000000000000000000000000000000000000000000000000000000006020820190815273ffffffffffffffffffffffffffffffffffffffff93841660248301529383166044820152606480820195909552938452614783939291906147136084856107d9565b1660008060409384519561472786886107d9565b602087527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020880152519082855af1903d156147a8573d61477461476b82610857565b945194856107d9565b83523d6000602085013e6169e2565b80518061478e575050565b816020806147a39361048b9501019101614af4565b6154fe565b606092506169e2565b73ffffffffffffffffffffffffffffffffffffffff6001541633036147d257565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b6020818303126104695780359067ffffffffffffffff821161046957016040818303126104695760405191614830836107bd565b813567ffffffffffffffff8111610469578161484d918401610d12565b8352602082013567ffffffffffffffff81116104695761486d9201610d12565b602082015290565b90816020910312610469576040519061488d826107a1565b51815290565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561046957016020813591019167ffffffffffffffff821161046957813603831361046957565b906109019160208152614a1c614a116149d46149156149028680614893565b6101006020880152610120870191614014565b61493561492460208801610480565b67ffffffffffffffff166040870152565b61496161494460408801610a17565b73ffffffffffffffffffffffffffffffffffffffff166060870152565b6060860135608086015261499761497a60808801610a17565b73ffffffffffffffffffffffffffffffffffffffff1660a0870152565b6149a460a0870187614893565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08784030160c0880152614014565b6149e160c0860186614893565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08684030160e0870152614014565b9260e0810190614893565b916101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082860301910152614014565b908161014091031261046957614aeb610120600092614a6a610838565b93614a7482610480565b8552614a8260208301613239565b6020860152614a936040830161324a565b604086015260608201356060860152614aae60808301613239565b608086015260a082013560a0860152614ac960c08301610a17565b60c086015260e082013560e08601525061010081013561010085015201613239565b61012082015290565b90816020910312610469575161090181612527565b9091614b2061090193604084526040840190610891565b916020818403910152610891565b35610901816109f9565b614b406134ad565b50606081013590614b518282615589565b614b69614b6160e08301836134c0565b8101906147fc565b614b8e611029608c8351015173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016908115159081614dc5575b508015614dad575b614d2d57506020614c4091614c05614bfe614bf660c08701876134c0565b810190614a4d565b82516157cf565b8181519101519060405193849283927f57ecfd2800000000000000000000000000000000000000000000000000000000845260048401614b09565b0381600073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af190811561068457600091614cfe575b5015614cd4577ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc067ffffffffffffffff61373b6040614ccd602086016138eb565b9401614b2e565b7fbf969f220000000000000000000000000000000000000000000000000000000060005260046000fd5b614d20915060203d602011614d26575b614d1881836107d9565b810190614af4565b38614c8c565b503d614d0e565b60009350614d6e91506020926040519485809481937f39077537000000000000000000000000000000000000000000000000000000008352600483016148e3565b03925af190811561068457600091614d84575090565b610901915060203d602011614da6575b614d9e81836107d9565b810190614875565b503d614d94565b506040614dbd60c08501856134c0565b905014614bd8565b905073ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691161438614bd0565b60409067ffffffffffffffff61090194931681528160208201520190610891565b908051156123b0578051602082012067ffffffffffffffff831692836000526007602052614e5e826005604060002001615b96565b15614eb7575081614ea67f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea93614ea1614eb2946000526008602052604060002090565b6144e9565b604051918291826108f0565b0390a2565b9050611fe06040519283927f393b8ad200000000000000000000000000000000000000000000000000000000845260048401614e08565b67ffffffffffffffff166000818152600660205260409020549092919015614ff05791614fed60e092614fb985614f457f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b976153bd565b846000526007602052614f5c816040600020615bf2565b614f65836153bd565b846000526007602052614f7f836002604060002001615bf2565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b90816020910312610469573590565b9081602091031261046957516109018161046e565b61048b909291926101208061014083019561506784825167ffffffffffffffff169052565b60208181015163ffffffff1690850152615089604082015160408601906128ae565b606081015160608501526150aa6080820151608086019063ffffffff169052565b60a081015160a08501526150db60c082015160c086019073ffffffffffffffffffffffffffffffffffffffff169052565b60e081015160e0850152610100810151610100850152015191019063ffffffff169052565b615108613960565b5061511281615ea0565b6020810190606001356151276135df836138eb565b61513282825461306b565b905573ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691615175816138eb565b92803b15610469576040517f7d552ea60000000000000000000000000000000000000000000000000000000081526004810184905267ffffffffffffffff9490941660248501526000908490604490829084905af19182156106845761526793612170936152cb575b507ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff615212846138eb565b6040805173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152336020820152908101949094521691606090a26138eb565b604051613cc68161529f60208201907ffa7c07de00000000000000000000000000000000000000000000000000000000602083019252565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826107d9565b8061067860006152da936107d9565b386151de565b6152e8614085565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff82511690602083019163ffffffff83511642034281116128705761534c906fffffffffffffffffffffffffffffffff60808701511690614191565b8101809111612870576153726fffffffffffffffffffffffffffffffff92918392616561565b161682524263ffffffff16905290565b61048b9092919260608101936fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b8051156154615760408101516fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff61542161540c60208501516fffffffffffffffffffffffffffffffff1690565b6fffffffffffffffffffffffffffffffff1690565b91161161542b5750565b611fe0906040519182917f8020d12400000000000000000000000000000000000000000000000000000000835260048301615382565b6fffffffffffffffffffffffffffffffff61548f60408301516fffffffffffffffffffffffffffffffff1690565b16158015906154d6575b6154a05750565b611fe0906040519182917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301615382565b506154f761540c60208301516fffffffffffffffffffffffffffffffff1690565b1515615499565b1561550557565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b6080810161559c611f26610a4483614b2e565b6157675750602081019061563d60206155e26155ba610541866138eb565b60801b7fffffffffffffffffffffffffffffffff000000000000000000000000000000001690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116600482015291829081906024820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561068457600091615748575b5061571e5761569d615698836138eb565b61638c565b6156a6826138eb565b906156c6611f2660a0830193610d706156bf86866134c0565b3691610cdb565b6156de575050906156d961048b926138eb565b6164b0565b6156e892506134c0565b90611fe06040519283927f24eb47e500000000000000000000000000000000000000000000000000000000845260048401614074565b7f53ad11d80000000000000000000000000000000000000000000000000000000060005260046000fd5b615761915060203d602011614d2657614d1881836107d9565b38615687565b6157736106c091614b2e565b7f961c9a4f0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b9061048b60246044936157ca600460026128ae565b6128ae565b805160948110615a145750600481015163ffffffff7f00000000000000000000000000000000000000000000000000000000000000001663ffffffff8216036159e15750600881015191600c820151609460908401519301519361583a602084015163ffffffff1690565b63ffffffff811663ffffffff8316036159a85750507f000000000000000000000000000000000000000000000000000000000000000063ffffffff811663ffffffff83160361596f5750506040016002815161589581612875565b61589e81612875565b0361593857506107d063ffffffff8216036158ff57506107d063ffffffff8216036158c65750565b7f0389caa2000000000000000000000000000000000000000000000000000000006000526107d060045263ffffffff1660245260446000fd5b7f22e102a0000000000000000000000000000000000000000000000000000000006000526107d060045263ffffffff1660245260446000fd5b6106c0905161594681612875565b7fc2fc5865000000000000000000000000000000000000000000000000000000006000526157b5565b7f77e480260000000000000000000000000000000000000000000000000000000060005263ffffffff9081166004521660245260446000fd5b7fe366a1170000000000000000000000000000000000000000000000000000000060005263ffffffff9081166004521660245260446000fd5b7f68d2f8d60000000000000000000000000000000000000000000000000000000060005263ffffffff1660045260246000fd5b7f758b22cc0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b80548210156132345760005260206000200190600090565b600081815260036020526040902054615ae4576002546801000000000000000081101561078057615acb615a968260018594016002556002615a41565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600254906000526003602052604060002055600190565b50600090565b600081815260106020526040902054615ae457600f546801000000000000000081101561078057615b27615a96826001859401600f55600f615a41565b9055600f54906000526010602052604060002055600190565b600081815260066020526040902054615ae4576005546801000000000000000081101561078057615b7d615a968260018594016005556005615a41565b9055600554906000526006602052604060002055600190565b6000828152600182016020526040902054615beb57805490680100000000000000008210156107805782615bd4615a96846001809601855584615a41565b905580549260005201602052604060002055600190565b5050600090565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1991615dd16130c9928054615c43615c3d615c348363ffffffff9060801c1690565b63ffffffff1690565b42613943565b9081615ddd575b5050615d8b6001615c6e60208601516fffffffffffffffffffffffffffffffff1690565b92615cf9615cbc61540c6fffffffffffffffffffffffffffffffff615ca385546fffffffffffffffffffffffffffffffff1690565b166fffffffffffffffffffffffffffffffff8816616561565b82906fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b615d4c615d068751151590565b82547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178255565b019182906fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b604083015181546fffffffffffffffffffffffffffffffff1660809190911b7fffffffffffffffffffffffffffffffff0000000000000000000000000000000016179055565b60405191829182615382565b61540c615cbc916fffffffffffffffffffffffffffffffff615e51615e589582615e4a60018a01549282615e43615e3c615e26876fffffffffffffffffffffffffffffffff1690565b996fffffffffffffffffffffffffffffffff1690565b9560801c90565b1690614191565b911661306b565b9116616561565b80547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff00000000000000000000000000000000161781553880615c4a565b60808101615eb3611f26610a4483614b2e565b61576757506020810190615ed160206155e26155ba610541866138eb565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561068457600091615f57575b5061571e576060615f4e61048b93615f3d615f3860408601614b2e565b616573565b61133d615f49826138eb565b61660a565b910135906166e8565b615f70915060203d602011614d2657614d1881836107d9565b38615f1b565b604051906002548083528260208101600260005260206000209260005b818110615fa857505061048b925003836107d9565b8454835260019485019487945060209093019201615f93565b604051906005548083528260208101600560005260206000209260005b818110615ff357505061048b925003836107d9565b8454835260019485019487945060209093019201615fde565b906040519182815491828252602082019060005260206000209260005b81811061603e57505061048b925003836107d9565b8454835260019485019487945060209093019201616029565b805480156160bb577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019061608c8282615a41565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600081815260036020526040902054908115615beb577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019082821161287057600254927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8401938411612870578383600095616186950361618c575b5050506161756002616057565b600390600052602052604060002090565b55600190565b6161756161b8916161ae6161a46161be956002615a41565b90549060031b1c90565b9283916002615a41565b906137f4565b55388080616168565b600081815260066020526040902054908115615beb577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019082821161287057600554927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84019384116128705783836000956161869503616263575b5050506162526005616057565b600690600052602052604060002090565b6162526161b89161627b6161a4616285956005615a41565b9283916005615a41565b55388080616245565b6001810191806000528260205260406000205492831515600014616366577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8401848111612870578354937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff85019485116128705760009585836161869761631e950361632d575b505050616057565b90600052602052604060002090565b61634d6161b8916163446161a461635d9588615a41565b92839187615a41565b8590600052602052604060002090565b55388080616316565b50505050600090565b9261637a9192614191565b81018091116128705761090191616561565b616398611f26826138f5565b616479576020616411916163c461102960045473ffffffffffffffffffffffffffffffffffffffff1690565b6040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90921660048301523360248301529092839190829081906044820190565b03915afa9081156106845760009161645a575b501561642c57565b7f728fe07b000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b616473915060203d602011614d2657614d1881836107d9565b38616424565b7fa9902c7e0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b67ffffffffffffffff7f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c91169182600052600760205280616531600260406000200173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692839161679f565b6040805173ffffffffffffffffffffffffffffffffffffffff909216825260208201929092529081908101614eb2565b908082101561656e575090565b905090565b7f000000000000000000000000000000000000000000000000000000000000000061659b5750565b73ffffffffffffffffffffffffffffffffffffffff16806000526003602052604060002054156165c85750565b7fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b908160209103126104695751610901816109f9565b616616611f26826138f5565b6164795760206166879161664261102960045473ffffffffffffffffffffffffffffffffffffffff1690565b60405180809581947fa8d87a3b0000000000000000000000000000000000000000000000000000000083526004830191909167ffffffffffffffff6020820193169052565b03915afa80156106845773ffffffffffffffffffffffffffffffffffffffff916000916166b9575b5016330361642c57565b6166db915060203d6020116166e1575b6166d381836107d9565b8101906165f5565b386166af565b503d6166c9565b67ffffffffffffffff7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894491169182600052600760205280616531604060002073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692839161679f565b8115616770570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b8054939290919060ff60a086901c161580156169da575b6169d3576167d56fffffffffffffffffffffffffffffffff861661540c565b906001840195865461680f615c3d615c3461680261540c856fffffffffffffffffffffffffffffffff1690565b9460801c63ffffffff1690565b8061693f575b50508381106168f45750828210616875575061048b93945061683a9161540c91613943565b6fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b906168ac6106c0936168a76168988461689261540c8c5460801c90565b93613943565b6168a183613916565b9061306b565b616766565b7fd0c8d23a0000000000000000000000000000000000000000000000000000000060005260045260245273ffffffffffffffffffffffffffffffffffffffff16604452606490565b7f1a76572a00000000000000000000000000000000000000000000000000000000600052600452602483905273ffffffffffffffffffffffffffffffffffffffff1660445260646000fd5b8285929395116169a95761695961540c6169609460801c90565b918561636f565b84547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178555913880616815565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050509050565b5081156167b6565b91929015616a5d57508151156169f6575090565b3b156169ff5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015616a705750805190602001fd5b611fe0906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352600483016108f056fea164736f6c634300081a000a2e902d38f15b233cbb63711add0fca4545334d3a169d60c0a616494d7eea9544",
}

var HybridLockReleaseUSDCTokenPoolABI = HybridLockReleaseUSDCTokenPoolMetaData.ABI

var HybridLockReleaseUSDCTokenPoolBin = HybridLockReleaseUSDCTokenPoolMetaData.Bin

func DeployHybridLockReleaseUSDCTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, legacyTokenMessenger common.Address, tokenMessenger common.Address, cctpMessageTransmitterProxy common.Address, token common.Address, allowlist []common.Address, rmnProxy common.Address, router common.Address, previousPool common.Address, lockBox common.Address) (common.Address, *types.Transaction, *HybridLockReleaseUSDCTokenPool, error) {
	parsed, err := HybridLockReleaseUSDCTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HybridLockReleaseUSDCTokenPoolBin), backend, legacyTokenMessenger, tokenMessenger, cctpMessageTransmitterProxy, token, allowlist, rmnProxy, router, previousPool, lockBox)
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) FINALITYTHRESHOLD(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "FINALITY_THRESHOLD")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) FINALITYTHRESHOLD() (uint32, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.FINALITYTHRESHOLD(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) FINALITYTHRESHOLD() (uint32, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.FINALITYTHRESHOLD(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) MAXFEE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "MAX_FEE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) MAXFEE() (uint32, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.MAXFEE(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) MAXFEE() (uint32, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.MAXFEE(&_HybridLockReleaseUSDCTokenPool.CallOpts)
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) ILegacyTokenMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "i_legacyTokenMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) ILegacyTokenMessenger() (common.Address, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.ILegacyTokenMessenger(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) ILegacyTokenMessenger() (common.Address, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.ILegacyTokenMessenger(&_HybridLockReleaseUSDCTokenPool.CallOpts)
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) ILockBox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "i_lockBox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) ILockBox() (common.Address, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.ILockBox(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) ILockBox() (common.Address, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.ILockBox(&_HybridLockReleaseUSDCTokenPool.CallOpts)
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) IPreviousMessageTransmitterProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "i_previousMessageTransmitterProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) IPreviousMessageTransmitterProxy() (common.Address, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.IPreviousMessageTransmitterProxy(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) IPreviousMessageTransmitterProxy() (common.Address, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.IPreviousMessageTransmitterProxy(&_HybridLockReleaseUSDCTokenPool.CallOpts)
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) ISupportedUSDCVersion(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "i_supportedUSDCVersion")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) ISupportedUSDCVersion() (uint32, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.ISupportedUSDCVersion(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) ISupportedUSDCVersion() (uint32, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.ISupportedUSDCVersion(&_HybridLockReleaseUSDCTokenPool.CallOpts)
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
	return common.HexToHash("0x802f1e7749437f70753a318fa838e9cd1afc1c06a032818e15050bfc4f01b424")
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
	FINALITYTHRESHOLD(opts *bind.CallOpts) (uint32, error)

	MAXFEE(opts *bind.CallOpts) (uint32, error)

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

	ILegacyTokenMessenger(opts *bind.CallOpts) (common.Address, error)

	ILocalDomainIdentifier(opts *bind.CallOpts) (uint32, error)

	ILockBox(opts *bind.CallOpts) (common.Address, error)

	IMessageTransmitterProxy(opts *bind.CallOpts) (common.Address, error)

	IPreviousMessageTransmitterProxy(opts *bind.CallOpts) (common.Address, error)

	IPreviousPool(opts *bind.CallOpts) (common.Address, error)

	ISupportedUSDCVersion(opts *bind.CallOpts) (uint32, error)

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
