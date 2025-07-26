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
	Bin: "0x610220806040523461067057617bc8803803809161001d8285610af8565b8339810190610120818303126106705761003681610b1b565b61004260208301610b1b565b9061004f60408401610b1b565b6060840151936001600160a01b03851693918486036106705760808201516001600160401b0381116106705782019680601f89011215610670578751976001600160401b038911610ae2578860051b9060208201996100b16040519b8c610af8565b8a526020808b019282010192831161067057602001905b828210610aca575050506100de60a08301610b1b565b966100eb60c08401610b1b565b6101046101006100fd60e08701610b1b565b9501610b1b565b983315610ab957600180546001600160a01b0319163317905587158015610aa8575b8015610a97575b610a8657608089905260c05260405163313ce56760e01b81526020816004818b5afa8091600091610a4a575b5090610a26575b50600660a052600480546001600160a01b0319166001600160a01b03929092169190911790558051151560e08190526108f5575b506001610100526001600160a01b0383169283156108e457604051632c12192160e01b8152602081600481885afa90811561067d576000916108aa575b5060405163054fd4d560e41b81526001600160a01b039190911690602081600481855afa90811561067d5760009161088b575b5063ffffffff80610100511691169081036108775750604051639cdbb18160e01b8152602081600481895afa90811561067d57600091610858575b5063ffffffff806101005116911690810361084457506040516367e0ed8360e11b81526020816004816001600160a01b0388165afa801561067d5782916000916107f6575b506001600160a01b0316036107e55760049260209261012052610140526040519283809263234d8e3d60e21b82525afa90811561067d576000916107b6575b5061016052608051610120516102e5916001600160a01b039182169116610b63565b6001600160a01b0381169030821461073c5781151590818061074d575b61073c578115610731576040516398db964360e01b8152602081600481875afa600091816106f5575b5061034157632d4d3c3d60e21b60005260046000fd5b6001600160a01b03166101a0525b61018052600080516020617ba88339815191526020604051858152a1156106d657906020600492604051938480926398db964360e01b82525afa6000928161069a575b506103a857632d4d3c3d60e21b60005260046000fd5b6001600160a01b039091166101a052600080516020617ba883398151915291602091905b6080516103e5906001600160a01b038381169116610b63565b6101c052604051908152a16101e08190526102008390526001600160a01b03831690811561068957604051636eb1769f60e11b81523060048201526001600160a01b038516602482015290602090829060449082905afa90811561067d57600091610646575b506105db57610486916040519163095ea7b360e01b60208401526024830152600019604483015260448252610481606483610af8565b610c6a565b61020052604051616bf49081610fb48239608051818181610580015281816109f501528181610a7c015281816122be015281816136fa01528181613b5901528181613bfc01528181613e7f01528181614d7a015281816152830152818161664a015261687f015260a05181610ac6015260c0518181816127e201528181615797015261602b015260e051818181610fd101528181612a2601526166b60152610100518181816109a401526159270152610120518181816111ba0152613de4015261014051818181611c280152614cfb0152610160518181816112f301528181613c790152615992015261018051818181610bfe0152614c4501526101a051818181611f000152614edf01526101c05181818161214e0152613bb901526101e051816118700152610200518181816105c2015281816118380152818161225a01528181612fd9015281816136c1015261524b0152f35b60405162461bcd60e51b815260206004820152603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527f20746f206e6f6e2d7a65726f20616c6c6f77616e6365000000000000000000006064820152608490fd5b90506020813d602011610675575b8161066160209383610af8565b8101031261067057513861044b565b600080fd5b3d9150610654565b6040513d6000823e3d90fd5b6321a9a7af60e01b60005260046000fd5b9092506020813d6020116106ce575b816106b660209383610af8565b81010312610670576106c790610b1b565b9138610392565b3d91506106a9565b50602090600080516020617ba88339815191529260006101a0526103cc565b9091506020813d602011610729575b8161071160209383610af8565b810103126106705761072290610b1b565b903861032b565b3d9150610704565b60006101a05261034f565b632d4d3c3d60e21b60005260046000fd5b506040516301ffc9a760e01b8152630e64dd2960e01b6004820152602081602481875afa90811561067d57600091610787575b5015610302565b6107a9915060203d6020116107af575b6107a18183610af8565b810190610b4b565b38610780565b503d610797565b6107d8915060203d6020116107de575b6107d08183610af8565b810190610b2f565b386102c3565b503d6107c6565b632a32133b60e11b60005260046000fd5b9091506020813d60201161083c575b8161081260209383610af8565b810103126108385751906001600160a01b03821682036108355750819038610284565b80fd5b5080fd5b3d9150610805565b6316ba39c560e31b60005260045260246000fd5b610871915060203d6020116107de576107d08183610af8565b3861023f565b6334697c6b60e11b60005260045260246000fd5b6108a4915060203d6020116107de576107d08183610af8565b38610204565b90506020813d6020116108dc575b816108c560209383610af8565b81010312610670576108d690610b1b565b386101d1565b3d91506108b8565b6306b7c75960e31b60005260046000fd5b9195909293946020956040519561090c8888610af8565b60008752600036813760e05115610a155760005b8751811015610987576001906001600160a01b0361093e828b610c40565b51168a61094a82610d9b565b610957575b505001610920565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a1388a61094f565b509193955091939660005b8651811015610a06576001906001600160a01b036109b0828a610c40565b51168015610a0057896109c282610e83565b6109d0575b50505b01610992565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a138896109c7565b506109ca565b50929550929590935038610194565b6335f4a7b360e01b60005260046000fd5b60ff1660068114610160576332ad3e0760e11b600052600660045260245260446000fd5b6020813d602011610a7e575b81610a6360209383610af8565b8101031261083857519060ff82168203610835575038610159565b3d9150610a56565b6342bcdf7f60e11b60005260046000fd5b506001600160a01b0381161561012d565b506001600160a01b03821615610126565b639b15e16f60e01b60005260046000fd5b60208091610ad784610b1b565b8152019101906100c8565b634e487b7160e01b600052604160045260246000fd5b601f909101601f19168101906001600160401b03821190821017610ae257604052565b51906001600160a01b038216820361067057565b90816020910312610670575163ffffffff811681036106705790565b90816020910312610670575180151581036106705790565b604051636eb1769f60e11b81523060048201526001600160a01b03831660248201529190602083806044810103816001600160a01b0385165afa92831561067d57600093610c0c575b506000198301809311610bf65760405163095ea7b360e01b60208201526001600160a01b039092166024830152604480830193909352918152610bf491610481606483610af8565b565b634e487b7160e01b600052601160045260246000fd5b90926020823d602011610c38575b81610c2760209383610af8565b810103126108355750519138610bac565b3d9150610c1a565b8051821015610c545760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b60018060a01b0316604091600080845192610c858685610af8565b602084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602085015260208151910182865af13d15610d76573d906001600160401b038211610ae2578451610cfa949092610ceb601f8201601f191660200185610af8565b83523d6000602085013e610ee3565b80519081610d0757505050565b602080610d18938301019101610b4b565b15610d205750565b5162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b91610cfa92606091610ee3565b8054821015610c545760005260206000200190600090565b6000818152600360205260409020548015610e7c576000198101818111610bf657600254600019810191908211610bf657818103610e2b575b5050506002548015610e155760001901610def816002610d83565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b610e64610e3c610e4d936002610d83565b90549060031b1c9283926002610d83565b819391549060031b91821b91600019901b19161790565b90556000526003602052604060002055388080610dd4565b5050600090565b80600052600360205260406000205415600014610edd5760025468010000000000000000811015610ae257610ec4610e4d8260018594016002556002610d83565b9055600254906000526003602052604060002055600190565b50600090565b91929015610f455750815115610ef7575090565b3b15610f005790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b825190915015610f585750805190602001fd5b6040519062461bcd60e51b8252602060048301528181519182602483015260005b838110610f9b5750508160006044809484010152601f80199101168101030190fd5b60208282018101516044878401015285935001610f7956fe6080604052600436101561001257600080fd5b60003560e01c806301ffc9a7146103975780631101dbd414610392578063181f5a771461038d578063212a052e1461038857806321df0da714610383578063240028e81461037e57806324f65ee71461037957806324f8795a146103745780632cfbb1191461036f578063390775371461036a5780633e591f2c146103655780634ad01f0b146103605780634c5ef0ed1461035b5780634c93ef841461035657806350d1a35a1461035157806354c8a4f31461034c5780636155cda01461034757806362ddd3c4146103425780636b716b0d1461033d5780636b795423146103385780636d3d1a5814610333578063714bf9071461032e57806379ba5097146103295780637d54534e146103245780638926f54f1461031f5780638a5e52bb1461031a5780638da5cb5b14610315578063962d40201461031057806398db96431461030b5780639a4575b914610306578063a42a7b8b14610301578063a7cd63b7146102fc578063a7eff066146102f7578063acfecf91146102f2578063af58d59f146102ed578063b0f479a1146102e8578063b2278077146102e3578063b7946580146102de578063bb5eced3146102d9578063bc063e1a146102d4578063c0d78655146102cf578063c4bffe2b146102ca578063c75eea9c146102c5578063cd306a6c146102c0578063cf7401f3146102bb578063da4b05e7146102b6578063dc0bd971146102b1578063de814c57146102ac578063dfadfa35146102a7578063e0351e13146102a2578063e8a1da171461029d578063e94ae6d014610298578063f2fde38b14610293578063f65a88861461028e578063f831bbcd146102895763fd6768551461028457600080fd5b612ffd565b612fac565b612f6d565b612e97565b612e42565b612a4b565b612a0e565b612978565b612806565b6127b5565b612798565b61269a565b612575565b61252a565b612497565b612375565b612359565b6121c6565b61218a565b612138565b612104565b612058565b611f24565b611ed3565b611e5f565b611d5d565b611c4c565b611bfb565b611a93565b611a2e565b611733565b6116f4565b611663565b611598565b611507565b6114d3565b611317565b6112d6565b611253565b6111a4565b610f9f565b610dd6565b610d96565b610d4d565b610c22565b610bd1565b610b88565b610b49565b610aea565b610aac565b610a42565b6109c8565b610987565b610924565b61048d565b34610469576020600319360112610469576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361046957807faff2afbf000000000000000000000000000000000000000000000000000000006020921490811561043f575b8115610415575b506040519015158152f35b7f01ffc9a7000000000000000000000000000000000000000000000000000000009150143861040a565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150610403565b600080fd5b67ffffffffffffffff81160361046957565b359061048b8261046e565b565b34610469576040600319360112610469576004356104aa8161046e565b602435906104e96104cf8267ffffffffffffffff166000526011602052604060002090565b5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff3391160361071c5767ffffffffffffffff8116610528816000526010602052604060002054151590565b6106e457600b5461054e9060a01c67ffffffffffffffff165b67ffffffffffffffff1690565b146106a9576105718167ffffffffffffffff16600052600c602052604060002090565b61057c8382546130e4565b90557f00000000000000000000000000000000000000000000000000000000000000006105ab83303384614745565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690813b15610469576040517f17b1f24200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9190911660048201526024810184905267ffffffffffffffff9290921660448301526000908290606490829084905af180156106a457610689575b50337fc17cea59c2955cb181b03393209566960365771dbba9dc3d510180e7cb312088600080a3005b80610698600061069e936107f9565b8061074a565b38610660565b6130f1565b7fd0da86c40000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b6000fd5b7f646972460000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b7f8e4a23d6000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b600091031261046957565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6060810190811067ffffffffffffffff8211176107a057604052565b610755565b60a0810190811067ffffffffffffffff8211176107a057604052565b6020810190811067ffffffffffffffff8211176107a057604052565b6040810190811067ffffffffffffffff8211176107a057604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176107a057604052565b6040519061048b60a0836107f9565b6040519061048b6020836107f9565b6040519061048b610140836107f9565b6040519061048b6040836107f9565b67ffffffffffffffff81116107a057601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b8481106108fb5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b806020809284010151828286010152016108bc565b9060206109219281815201906108b1565b90565b3461046957600060031936011261046957610983604080519061094781836107f9565b601782527f55534443546f6b656e506f6f6c20312e362e322d6465760000000000000000006020830152519182916020835260208301906108b1565b0390f35b3461046957600060031936011261046957602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461046957600060031936011261046957602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b73ffffffffffffffffffffffffffffffffffffffff81160361046957565b359061048b82610a19565b34610469576020600319360112610469576020610aa2600435610a6481610a19565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691161490565b6040519015158152f35b3461046957600060031936011261046957602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346104695760206003193601126104695760043567ffffffffffffffff8111610469573660238201121561046957806004013567ffffffffffffffff81116104695736602460c0830284010111610469576024610b4792016130fd565b005b346104695760206003193601126104695767ffffffffffffffff600435610b6f8161046e565b16600052600c6020526020604060002054604051908152f35b346104695760206003193601126104695760043567ffffffffffffffff811161046957610100600319823603011261046957610bc860209160040161358a565b60405190518152f35b3461046957600060031936011261046957602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461046957600060031936011261046957610c3b614851565b600b5467ffffffffffffffff8160a01c168015610cd1577fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff7f375f1ad1194a2bec317c5efec05cc63ffa06ddd0c4b276619f6fd47298eda5189216600b556000610cb98267ffffffffffffffff16600052600d602052604060002090565b5560405167ffffffffffffffff9091168152602090a1005b7fa94cb9880000000000000000000000000000000000000000000000000000000060005260046000fd5b929192610d0782610877565b91610d1560405193846107f9565b829481845281830111610469578281602093846000960137010152565b9080601f830112156104695781602061092193359101610cfb565b3461046957604060031936011261046957600435610d6a8161046e565b60243567ffffffffffffffff811161046957602091610d90610aa2923690600401610d32565b906138d0565b34610469576020600319360112610469576020610aa2600435610db88161046e565b67ffffffffffffffff16600052600e60205260ff6040600020541690565b3461046957602060031936011261046957600435610df38161046e565b610dfb614851565b67ffffffffffffffff600b5460a01c16610ef25767ffffffffffffffff81166000908152600e602052604090205460ff1615610ec857610ec381610ea87f20331f191af84dbff48b162aa5a5985e7891ae646297b0a2ac80487f9109ef49937fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff7bffffffffffffffff0000000000000000000000000000000000000000600b549260a01b16911617600b55565b60405167ffffffffffffffff90911681529081906020820190565b0390a1005b7f656535ce0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f692bc1310000000000000000000000000000000000000000000000000000000060005260046000fd5b9181601f840112156104695782359167ffffffffffffffff8311610469576020808501948460051b01011161046957565b60406003198201126104695760043567ffffffffffffffff81116104695781610f7891600401610f1c565b929092916024359067ffffffffffffffff821161046957610f9b91600401610f1c565b9091565b3461046957610fc7610fcf610fb336610f4d565b9491610fc0939193614851565b3691613925565b923691613925565b7f00000000000000000000000000000000000000000000000000000000000000001561117a5760005b82518110156110bd578061102b61101160019386613f8d565b5173ffffffffffffffffffffffffffffffffffffffff1690565b61106761106273ffffffffffffffffffffffffffffffffffffffff83165b73ffffffffffffffffffffffffffffffffffffffff1690565b61622b565b611073575b5001610ff8565b60405173ffffffffffffffffffffffffffffffffffffffff9190911681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756690602090a13861106c565b5060005b8151811015610b4757806110da61101160019385613f8d565b73ffffffffffffffffffffffffffffffffffffffff8116156111745761111d61111873ffffffffffffffffffffffffffffffffffffffff8316611049565b615b9a565b61112a575b505b016110c1565b60405173ffffffffffffffffffffffffffffffffffffffff9190911681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d890602090a183611122565b50611124565b7f35f4a7b30000000000000000000000000000000000000000000000000000000060005260046000fd5b34610469576000600319360112610469576040517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168152602090f35b60406003198201126104695760043561120e8161046e565b9160243567ffffffffffffffff811161046957826023820112156104695780600401359267ffffffffffffffff84116104695760248483010111610469576024019190565b3461046957611261366111f6565b61126c929192614851565b67ffffffffffffffff821661128e816000526006602052604060002054151590565b156112a95750610b47926112a3913691610cfb565b90614f29565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b3461046957600060031936011261046957602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346104695761132536610f4d565b929091611330614851565b60005b8181106114585750505060005b82811061134957005b61137561136261054161135d84878761397b565b61398b565b6000526010602052604060002054151590565b61141157806113d56113aa61139061135d600195888861397b565b67ffffffffffffffff16600052600e602052604060002090565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055565b6113e661054161135d83878761397b565b7f5e3985e51df58346365017cae614e59d723143b71c9a2ce4a156687f1f2c3f5a600080a201611340565b61135d906106e0936114229361397b565b7f646972460000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b8061149761146f61139061135d600195878961397b565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008154169055565b6114a861054161135d83868861397b565b7fddc5afbc5e53c63a556964db0eef76a1c2d9305e0811abd7410d2a6f4799490e600080a201611333565b3461046957600060031936011261046957602073ffffffffffffffffffffffffffffffffffffffff60095416604051908152f35b34610469576020600319360112610469577f084e6f0e9791c2e56153bd49e6ec6dd63ba9a72c258d71558d74c63fc75b7168602073ffffffffffffffffffffffffffffffffffffffff60043561155c81610a19565b611564614851565b16807fffffffffffffffffffffffff0000000000000000000000000000000000000000600b541617600b55604051908152a1005b346104695760006003193601126104695760005473ffffffffffffffffffffffffffffffffffffffff81163303611639577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b34610469576020600319360112610469577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174602073ffffffffffffffffffffffffffffffffffffffff6004356116b881610a19565b6116c0614851565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000006009541617600955604051908152a1005b34610469576020600319360112610469576020610aa267ffffffffffffffff60043561171f8161046e565b166000526006602052604060002054151590565b3461046957600060031936011261046957600b5461176673ffffffffffffffffffffffffffffffffffffffff8216611049565b3303611a045760a01c67ffffffffffffffff1667ffffffffffffffff8116908115610cd1576117d26117ac8267ffffffffffffffff16600052600c602052604060002090565b546117cb8367ffffffffffffffff16600052600d602052604060002090565b54906139e3565b9060006117f38267ffffffffffffffff16600052600c602052604060002090565b556118217fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff600b5416600b55565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001693803b15610469576040517f233671d500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff861660048201526024810185905230604482015267ffffffffffffffff84166064820152906000908290608490829084905af180156106a4576119ef575b50833b1561046957600060405180957f42966c6800000000000000000000000000000000000000000000000000000000825281838161195889600483019190602083019252565b03925af19081156106a4577fdea60ddd4c7ebdab804f5694c70350cca7893ece3efeecb142312eacac5c73e4946119b5926119da575b506119b061146f8467ffffffffffffffff16600052600e602052604060002090565b615c2b565b506040805167ffffffffffffffff909216825260208201929092529081908101610ec3565b8061069860006119e9936107f9565b3861198e565b8061069860006119fe936107f9565b38611911565b7f438a7a050000000000000000000000000000000000000000000000000000000060005260046000fd5b3461046957600060031936011261046957602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b9181601f840112156104695782359167ffffffffffffffff8311610469576020808501946060850201011161046957565b346104695760606003193601126104695760043567ffffffffffffffff811161046957611ac4903690600401610f1c565b9060243567ffffffffffffffff811161046957611ae5903690600401611a62565b9060443567ffffffffffffffff811161046957611b06903690600401611a62565b611b2861104960095473ffffffffffffffffffffffffffffffffffffffff1690565b33141580611bd0575b61071c57838614801590611bc6575b611b9c5760005b868110611b5057005b80611b96611b6461135d6001948b8b61397b565b611b6f8389896139f0565b611b90611b88611b8086898b6139f0565b923690612651565b913690612651565b91614fee565b01611b47565b7f568efce20000000000000000000000000000000000000000000000000000000060005260046000fd5b5080861415611b40565b50611bf361104960015473ffffffffffffffffffffffffffffffffffffffff1690565b331415611b31565b3461046957600060031936011261046957602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346104695760206003193601126104695760043567ffffffffffffffff81116104695760a0600319823603011261046957611c8c61098391600401613a19565b604051918291602083526020611cad825160408387015260608601906108b1565b9101517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08483030160408501526108b1565b602081016020825282518091526040820191602060408360051b8301019401926000915b838310611d1257505050505090565b9091929394602080611d4e837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0866001960301875289516108b1565b97019301930191939290611d03565b346104695760206003193601126104695767ffffffffffffffff600435611d838161046e565b166000526007602052611d9c600560406000200161614d565b8051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611de2611dcc8461390d565b93611dda60405195866107f9565b80855261390d565b0160005b818110611e4e57505060005b8151811015611e405780611e24611e1f611e0e60019486613f8d565b516000526008602052604060002090565b613ff4565b611e2e8286613f8d565b52611e398185613f8d565b5001611df2565b604051806109838582611cdf565b806060602080938701015201611de6565b3461046957600060031936011261046957611e786160b7565b60405180916020820160208352815180915260206040840192019060005b818110611ea4575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101611e96565b3461046957600060031936011261046957602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461046957611f32366111f6565b611f3d929192614851565b67ffffffffffffffff821691611f67611f63846000526006602052604060002054151590565b1590565b61202157611faa611f636005611f918467ffffffffffffffff166000526007602052604060002090565b01611f9d368689610cfb565b60208151910120906163cf565b611fe657507f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d769192611fe160405192839283614114565b0390a2005b61201d84926040519384937f74f23c7c000000000000000000000000000000000000000000000000000000008552600485016140f3565b0390fd5b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b346104695760206003193601126104695767ffffffffffffffff60043561207e8161046e565b612086614125565b501660005260076020526109836120ab6120a66002604060002001614150565b615421565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b3461046957600060031936011261046957602073ffffffffffffffffffffffffffffffffffffffff60045416604051908152f35b34610469576000600319360112610469576040517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168152602090f35b34610469576020600319360112610469576109836121b26004356121ad8161046e565b6141aa565b6040519182916020835260208301906108b1565b34610469576040600319360112610469576004356121e38161046e565b602435906121ef614851565b600b546122089060a01c67ffffffffffffffff16610541565b67ffffffffffffffff8216146106a9576122368167ffffffffffffffff16600052600c602052604060002090565b6122418382546139e3565b905573ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803b15610469576040517f233671d500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001660048201526024810184905233604482015267ffffffffffffffff9290921660648301526000908290608490829084905af180156106a457612344575b50337fc2c3f06e49b9f15e7b4af9055e183b0d73362e033ad82a07dec9bf9840171719600080a3005b806106986000612353936107f9565b3861231b565b3461046957600060031936011261046957602060405160008152f35b346104695760206003193601126104695773ffffffffffffffffffffffffffffffffffffffff6004356123a781610a19565b6123af614851565b1680156124295760407f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f16849160045490807fffffffffffffffffffffffff000000000000000000000000000000000000000083161760045573ffffffffffffffffffffffffffffffffffffffff8351921682526020820152a1005b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b602060408183019282815284518094520192019060005b8181106124775750505090565b825167ffffffffffffffff1684526020938401939092019160010161246a565b34610469576000600319360112610469576124b0616102565b8051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06124e0611dcc8461390d565b0136602084013760005b815181101561251c578067ffffffffffffffff61250960019385613f8d565b51166125158286613f8d565b52016124ea565b604051806109838582612453565b346104695760206003193601126104695767ffffffffffffffff6004356125508161046e565b612558614125565b501660005260076020526109836120ab6120a66040600020614150565b3461046957600060031936011261046957602067ffffffffffffffff600b5460a01c16604051908152f35b8015150361046957565b35906fffffffffffffffffffffffffffffffff8216820361046957565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c606091011261046957604051906125fe82610784565b8160843561260b816125a0565b815260a4356fffffffffffffffffffffffffffffffff8116810361046957602082015260c435906fffffffffffffffffffffffffffffffff821682036104695760400152565b91908260609103126104695760405161266981610784565b6040612695818395803561267c816125a0565b855261268a602082016125aa565b6020860152016125aa565b910152565b346104695760e0600319360112610469576004356126b78161046e565b60607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc360112610469576040516126ed81610784565b6024356126f9816125a0565b81526044356fffffffffffffffffffffffffffffffff811681036104695760208201526064356fffffffffffffffffffffffffffffffff81168103610469576040820152612746366125c7565b9073ffffffffffffffffffffffffffffffffffffffff6009541633141580612776575b61071c57610b4792614fee565b5073ffffffffffffffffffffffffffffffffffffffff60015416331415612769565b346104695760006003193601126104695760206040516107d08152f35b3461046957600060031936011261046957602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610469576040600319360112610469576004356128238161046e565b6024359061282f614851565b67ffffffffffffffff600b5460a01c169167ffffffffffffffff8216809303610cd15782600052600d60205260406000208054928284018094116128e9577fe1e6c22ce6b566f66cdb457ec2e7910ff1f9a9e5654ed75303476fa870468220936128d192556117cb6128b58267ffffffffffffffff16600052600c602052604060002090565b549167ffffffffffffffff16600052600d602052604060002090565b60408051928352602083019190915281908101611fe1565b6130b5565b600311156128f857565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9060038210156128f85752565b91909160808060a0830194805184526020810151602085015263ffffffff604082015116604085015261296f60608201516060860190612927565b01511515910152565b346104695760206003193601126104695767ffffffffffffffff60043561299e8161046e565b6129a6614125565b5016600052600a602052610983604060002060ff6002604051926129c9846107a5565b8054845260018101546020850152015463ffffffff811660408401526129f7828260201c16606085016133a1565b60281c161515608082015260405191829182612934565b346104695760006003193601126104695760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b3461046957612a5936610f4d565b919092612a64614851565b6000915b808310612d095750505060009163ffffffff4216925b828110612a8757005b612a9a612a95828585614327565b6143e6565b9060608201612aa981516154fe565b6080830193612ab885516154fe565b60408401908151511561242957612ae5611f63612ae0610541885167ffffffffffffffff1690565b615c81565b612cbe57612c1e612b1e612b04879a999a5167ffffffffffffffff1690565b67ffffffffffffffff166000526007602052604060002090565b612be189612bdb8751612bc2612b4760408301516fffffffffffffffffffffffffffffffff1690565b91612ba9612b72612b6b60208401516fffffffffffffffffffffffffffffffff1690565b9251151590565b612ba0612b7d61083a565b6fffffffffffffffffffffffffffffffff851681529763ffffffff166020890152565b15156040870152565b6fffffffffffffffffffffffffffffffff166060850152565b6fffffffffffffffffffffffffffffffff166080830152565b8261447d565b612c1389612c0a8a51612bc2612b4760408301516fffffffffffffffffffffffffffffffff1690565b6002830161447d565b600484519101614589565b602085019660005b88518051821015612c615790612c5b600192612c5483612c4e8c5167ffffffffffffffff1690565b92613f8d565b5190614f29565b01612c26565b50509796509490612cb57f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c29392612ca26001975167ffffffffffffffff1690565b92519351905190604051948594856146b0565b0390a101612a7e565b6106e0612cd3865167ffffffffffffffff1690565b7f1d5ad3c50000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b909192612d1a61135d85848661397b565b94612d31611f6367ffffffffffffffff8816616308565b612e0a57612d5e6005612d588867ffffffffffffffff166000526007602052604060002090565b0161614d565b9360005b8551811015612daa57600190612da36005612d918b67ffffffffffffffff166000526007602052604060002090565b01612d9c838a613f8d565b51906163cf565b5001612d62565b509350937f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916612dfc60019397610ea8612df78267ffffffffffffffff166000526007602052604060002090565b614278565b0390a1019190939293612a68565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff861660045260246000fd5b346104695760206003193601126104695767ffffffffffffffff600435612e688161046e565b166000526011602052602073ffffffffffffffffffffffffffffffffffffffff60406000205416604051908152f35b346104695760206003193601126104695773ffffffffffffffffffffffffffffffffffffffff600435612ec981610a19565b612ed1614851565b16338114612f4357807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346104695760206003193601126104695767ffffffffffffffff600435612f938161046e565b16600052600d6020526020604060002054604051908152f35b3461046957600060031936011261046957602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346104695760406003193601126104695760043561301a8161046e565b67ffffffffffffffff6024359161303083610a19565b613038614851565b166000818152601160205260408120805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff0000000000000000000000000000000000000000821681179092559293909216907fc82aa48e67c70b1ad1494533456f52504bb4d62d11bbdafaeb98cfccd1ed817e9080a4005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b919082018092116128e957565b6040513d6000823e3d90fd5b613105614851565b60005b8281106131475750907f802f1e7749437f70753a318fa838e9cd1afc1c06a032818e15050bfc4f01b4249161314260405192839283613484565b0390a1565b61315a61315582858561329d565b6132d0565b8051158015613248575b613212579061320c826132076131ed606060019651936131de60208201516131d5613196604085015163ffffffff1690565b6131cd6080860151916131a8836128ee565b60a08701511515946131b861083a565b9b8c5260208c015263ffffffff1660408b0152565b8589016133a1565b15156080870152565b015167ffffffffffffffff1690565b67ffffffffffffffff16600052600a602052604060002090565b6133ad565b01613108565b61201d906040519182917fa7a9337d00000000000000000000000000000000000000000000000000000000835260048301613352565b5067ffffffffffffffff613267606083015167ffffffffffffffff1690565b1615613164565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b91908110156132ad5760c0020190565b61326e565b359063ffffffff8216820361046957565b3590600382101561046957565b60c081360312610469576040519060c082019082821067ffffffffffffffff8311176107a05760a0916040528035835260208101356020840152613316604082016132b2565b604084015260608101356133298161046e565b606084015261333a608082016132c3565b6080840152013561334a816125a0565b60a082015290565b91909160a08060c0830194805184526020810151602085015263ffffffff604082015116604085015267ffffffffffffffff606082015116606085015261296f60808201516080860190612927565b60038210156128f85752565b60029082518155602083015160018201550163ffffffff6040830151167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000082541617815560608201519160038310156128f85760038310156128f857608061344d9161048b947fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff64ff0000000086549260201b1691161784550151151590565b81547fffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffffff1690151560281b65ff000000000016179055565b602080825281018390526040019160005b8181106134a25750505090565b90919260c080600192863581526020870135602082015263ffffffff6134ca604089016132b2565b16604082015267ffffffffffffffff60608801356134e78161046e565b1660608201526135066134fc608089016132c3565b6080830190612927565b60a0870135613514816125a0565b151560a0820152019401929101613495565b60405190613533826107c1565b60008252565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610469570180359067ffffffffffffffff82116104695760200191813603831361046957565b604051613596816107c1565b600090527ffa7c07de000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000006135e960c0840184613539565b9035828116916004811061387f575b5050160361387657613608613526565b5060608101359061361982826156ca565b600b5460a01c67ffffffffffffffff16602082019061363a6105418361398b565b67ffffffffffffffff8216146106a957506136716136578261398b565b67ffffffffffffffff16600052600c602052604060002090565b546138585761369c6136828261398b565b67ffffffffffffffff16600052600d602052604060002090565b6136a78482546139e3565b90555b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016930161372581614bce565b61372e8461398b565b833b15610469576040517f233671d500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff87811660048301526024820189905292909216604483015267ffffffffffffffff166064820152916000908390608490829084905af180156106a457856137f46137ee7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09667ffffffffffffffff9661383295613843575b5061398b565b93614bce565b6040805173ffffffffffffffffffffffffffffffffffffffff9889168152336020820152979091169087015260608601529116929081906080820190565b0390a261383d610849565b90815290565b806106986000613852936107f9565b386137e8565b6138646136578261398b565b61386f8482546139e3565b90556136aa565b61092190614bd8565b839250829060040360031b1b161638806135f8565b916138cc918354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055565b9067ffffffffffffffff61092192166000526007602052600560406000200190602081519101209060019160005201602052604060002054151590565b67ffffffffffffffff81116107a05760051b60200190565b9291906139318161390d565b9361393f60405195866107f9565b602085838152019160051b810192831161046957905b82821061396157505050565b60208091833561397081610a19565b815201910190613955565b91908110156132ad5760051b0190565b356109218161046e565b67ffffffffffffffff61092191166000526006602052604060002054151590565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116128e957565b919082039182116128e957565b91908110156132ad576060020190565b60405190613a0d826107dd565b60606020838281520152565b613a21613a00565b5060208101613a348135610db88161046e565b15613a7657613a5c610541613a56600b5467ffffffffffffffff9060a01c1690565b9261398b565b67ffffffffffffffff8216146106a9575061092190615200565b613a7e613a00565b50613a8882615fe1565b613a9c613a976131ed8361398b565b6141cc565b91613aad611f636080850151151590565b613f4b576020613abd8280613539565b905003613f0b5760208301518015613eef57905b6000806060860160018151613ae5816128ee565b613aee816128ee565b03613da8575050506001916060820135602082613b12604089015163ffffffff1690565b88516040517ff856ddb6000000000000000000000000000000000000000000000000000000008152600481019590955263ffffffff909116602485015260448401919091527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1660648401526084830152818060a481010381600073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af19081156106a457600091613d79575b50945b613bf68561398b565b604080517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1680825233602083015260609687013592820183905295919267ffffffffffffffff16917ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1091a27f0000000000000000000000000000000000000000000000000000000000000000915194613ca3610858565b67ffffffffffffffff909816885263ffffffff83166020890152613cca90604089016133a1565b606087015263ffffffff16608086015260a085015273ffffffffffffffffffffffffffffffffffffffff1660c084015260e083015260006101008301526107d0610120830152613d199061398b565b613d22906141aa565b9060405180916020820190613d3691615142565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018252613d6690826107f9565b613d6e610868565b918252602082015290565b613d9b915060203d602011613da1575b613d9381836107f9565b81019061512d565b38613bea565b503d613d89565b600290969194929651613dba816128ee565b613dc3816128ee565b03613bed57925060029273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166060840135613e1a604084015163ffffffff1690565b90835192803b15610469576040517fd04857b0000000000000000000000000000000000000000000000000000000008152600481019290925263ffffffff9290921660248201526044810185905273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001660648201526084810192909252600060a483018190526107d060c484015290829060e490829084905af180156106a457613eda575b50613bed565b806106986000613ee9936107f9565b38613ed4565b50613f05613efd8280613539565b81019061511e565b90613ad1565b80613f1591613539565b9061201d6040519283927fa3c8cf0900000000000000000000000000000000000000000000000000000000845260048401614114565b6106e0613f578361398b565b7fd201c48a0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b80518210156132ad5760209160051b010190565b90600182811c92168015613fea575b6020831014613fbb57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691613fb0565b906040519182600082549261400884613fa1565b8084529360018116908115614074575060011461402d575b5061048b925003836107f9565b90506000929192526020600020906000915b81831061405857505090602061048b9282010138614020565b602091935080600191548385890101520191019091849261403f565b6020935061048b9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138614020565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b60409067ffffffffffffffff610921959316815281602082015201916140b4565b9160206109219381815201916140b4565b60405190614132826107a5565b60006080838281528260208201528260408201528260608201520152565b9060405161415d816107a5565b60806fffffffffffffffffffffffffffffffff6001839560ff8154848116875263ffffffff81871c16602088015260a01c1615156040860152015481808216166060850152821c16910152565b67ffffffffffffffff1660005260076020526109216004604060002001613ff4565b906040516141d9816107a5565b608060ff600283958054855260018101546020860152015463ffffffff8116604085015261420f828260201c16606086016133a1565b60281c161515910152565b818110614225575050565b6000815560010161421a565b818102929181159184041417156128e957565b8054906000815581614254575050565b6000526020600020908101905b81811061426c575050565b60008155600101614261565b600561048b9160008155600060018201556000600282015560006003820155600481016142a58154613fa1565b90816142b4575b505001614244565b81601f600093116001146142cc5750555b38806142ac565b818352602083206142e791601f01861c81019060010161421a565b808252602082209081548360011b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8560031b1c1916179055556142c5565b91908110156132ad5760051b810135907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee181360301821215610469570190565b9080601f8301121561046957813561437e8161390d565b9261438c60405194856107f9565b81845260208085019260051b820101918383116104695760208201905b8382106143b857505050505090565b813567ffffffffffffffff8111610469576020916143db87848094880101610d32565b8152019101906143a9565b6101208136031261046957604051906143fe826107a5565b61440781610480565b8252602081013567ffffffffffffffff81116104695761442a9036908301614367565b602083015260408101359067ffffffffffffffff8211610469576144546144759236908301610d32565b60408401526144663660608301612651565b606084015260c0369101612651565b608082015290565b8151815460208401516040850151608091821b73ffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffff000000000000000000000000000000000000000000909416939093179290921791151560a01b74ff000000000000000000000000000000000000000016919091178355606084015193810151901b7fffffffffffffffffffffffffffffffff0000000000000000000000000000000016921691909117600190910155565b9190601f811161455357505050565b61048b926000526020600020906020601f840160051c8301931061457f575b601f0160051c019061421a565b9091508190614572565b919091825167ffffffffffffffff81116107a0576145b1816145ab8454613fa1565b84614544565b6020601f821160011461460b5781906138cc939495600092614600575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b0151905038806145ce565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169061463e84600052602060002090565b9160005b81811061469857509583600195969710614661575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388080614657565b9192602060018192868b015181550194019201614642565b6147146146df61048b9597969467ffffffffffffffff60a09516845261010060208501526101008401906108b1565b9660408301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b01906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b6040517f23b872dd000000000000000000000000000000000000000000000000000000006020820190815273ffffffffffffffffffffffffffffffffffffffff93841660248301529383166044820152606480820195909552938452614823939291906147b36084856107f9565b166000806040938451956147c786886107f9565b602087527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020880152519082855af1903d15614848573d61481461480b82610877565b945194856107f9565b83523d6000602085013e616b23565b80518061482e575050565b816020806148439361048b9501019101614b94565b61563f565b60609250616b23565b73ffffffffffffffffffffffffffffffffffffffff60015416330361487257565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b6020818303126104695780359067ffffffffffffffff8211610469570160408183031261046957604051916148d0836107dd565b813567ffffffffffffffff811161046957816148ed918401610d32565b8352602082013567ffffffffffffffff81116104695761490d9201610d32565b602082015290565b90816020910312610469576040519061492d826107c1565b51815290565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561046957016020813591019167ffffffffffffffff821161046957813603831361046957565b906109219160208152614abc614ab1614a746149b56149a28680614933565b61010060208801526101208701916140b4565b6149d56149c460208801610480565b67ffffffffffffffff166040870152565b614a016149e460408801610a37565b73ffffffffffffffffffffffffffffffffffffffff166060870152565b60608601356080860152614a37614a1a60808801610a37565b73ffffffffffffffffffffffffffffffffffffffff1660a0870152565b614a4460a0870187614933565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08784030160c08801526140b4565b614a8160c0860186614933565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08684030160e08701526140b4565b9260e0810190614933565b916101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828603019101526140b4565b908161014091031261046957614b8b610120600092614b0a610858565b93614b1482610480565b8552614b22602083016132b2565b6020860152614b33604083016132c3565b604086015260608201356060860152614b4e608083016132b2565b608086015260a082013560a0860152614b6960c08301610a37565b60c086015260e082013560e086015250610100810135610100850152016132b2565b61012082015290565b908160209103126104695751610921816125a0565b9091614bc0610921936040845260408401906108b1565b9160208184039101526108b1565b3561092181610a19565b614be0613526565b50606081013590614bf182826156ca565b614c09614c0160e0830183613539565b81019061489c565b614c2e611049608c8351015173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016908115159081614ec5575b508015614ead575b614e2d57506020614ce091614ca5614c9e614c9660c0870187613539565b810190614aed565b8251615910565b8181519101519060405193849283927f57ecfd2800000000000000000000000000000000000000000000000000000000845260048401614ba9565b0381600073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af19081156106a457600091614dfe575b5015614dd4577ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc067ffffffffffffffff614d746040614d6d6020860161398b565b9401614bce565b604080517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff908116825233602083015290921690820152606081018590529216918060808101613832565b7fbf969f220000000000000000000000000000000000000000000000000000000060005260046000fd5b614e20915060203d602011614e26575b614e1881836107f9565b810190614b94565b38614d2c565b503d614e0e565b60009350614e6e91506020926040519485809481937f3907753700000000000000000000000000000000000000000000000000000000835260048301614983565b03925af19081156106a457600091614e84575090565b610921915060203d602011614ea6575b614e9e81836107f9565b810190614915565b503d614e94565b506040614ebd60c0850185613539565b905014614c78565b905073ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691161438614c70565b60409067ffffffffffffffff610921949316815281602082015201906108b1565b90805115612429578051602082012067ffffffffffffffff831692836000526007602052614f5e826005604060002001615cd7565b15614fb7575081614fa67f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea93614fa1614fb2946000526008602052604060002090565b614589565b60405191829182610910565b0390a2565b905061201d6040519283927f393b8ad200000000000000000000000000000000000000000000000000000000845260048401614f08565b67ffffffffffffffff1660008181526006602052604090205490929190156150f057916150ed60e0926150b9856150457f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b976154fe565b84600052600760205261505c816040600020615d33565b615065836154fe565b84600052600760205261507f836002604060002001615d33565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b90816020910312610469573590565b9081602091031261046957516109218161046e565b61048b909291926101208061014083019561516784825167ffffffffffffffff169052565b60208181015163ffffffff169085015261518960408201516040860190612927565b606081015160608501526151aa6080820151608086019063ffffffff169052565b60a081015160a08501526151db60c082015160c086019073ffffffffffffffffffffffffffffffffffffffff169052565b60e081015160e0850152610100810151610100850152015191019063ffffffff169052565b615208613a00565b5061521281615fe1565b6020810190606001356152276136578361398b565b6152328282546130e4565b905573ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166152ac8261398b565b93803b15610469576040517f17b1f24200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff831660048201526024810185905267ffffffffffffffff9590951660448601526000908590606490829084905af19384156106a45767ffffffffffffffff6121ad947ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10926153a89761540c575b506153a06153698661398b565b6040805173ffffffffffffffffffffffffffffffffffffffff90971687523360208801528601929092529116929081906060820190565b0390a261398b565b604051613d66816153e060208201907ffa7c07de00000000000000000000000000000000000000000000000000000000602083019252565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826107f9565b80610698600061541b936107f9565b3861535c565b615429614125565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff82511690602083019163ffffffff83511642034281116128e95761548d906fffffffffffffffffffffffffffffffff60808701511690614231565b81018091116128e9576154b36fffffffffffffffffffffffffffffffff929183926166a2565b161682524263ffffffff16905290565b61048b9092919260608101936fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b8051156155a25760408101516fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff61556261554d60208501516fffffffffffffffffffffffffffffffff1690565b6fffffffffffffffffffffffffffffffff1690565b91161161556c5750565b61201d906040519182917f8020d124000000000000000000000000000000000000000000000000000000008352600483016154c3565b6fffffffffffffffffffffffffffffffff6155d060408301516fffffffffffffffffffffffffffffffff1690565b1615801590615617575b6155e15750565b61201d906040519182917fd68af9cc000000000000000000000000000000000000000000000000000000008352600483016154c3565b5061563861554d60208301516fffffffffffffffffffffffffffffffff1690565b15156155da565b1561564657565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b608081016156dd611f63610a6483614bce565b6158a85750602081019061577e60206157236156fb6105418661398b565b60801b7fffffffffffffffffffffffffffffffff000000000000000000000000000000001690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116600482015291829081906024820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156106a457600091615889575b5061585f576157de6157d98361398b565b6164cd565b6157e78261398b565b90615807611f6360a0830193610d906158008686613539565b3691610cfb565b61581f5750509061581a61048b9261398b565b6165f1565b6158299250613539565b9061201d6040519283927f24eb47e500000000000000000000000000000000000000000000000000000000845260048401614114565b7f53ad11d80000000000000000000000000000000000000000000000000000000060005260046000fd5b6158a2915060203d602011614e2657614e1881836107f9565b386157c8565b6158b46106e091614bce565b7f961c9a4f0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b9061048b602460449361590b60046002612927565b612927565b805160948110615b555750600481015163ffffffff7f00000000000000000000000000000000000000000000000000000000000000001663ffffffff821603615b225750600881015191600c820151609460908401519301519361597b602084015163ffffffff1690565b63ffffffff811663ffffffff831603615ae95750507f000000000000000000000000000000000000000000000000000000000000000063ffffffff811663ffffffff831603615ab0575050604001600281516159d6816128ee565b6159df816128ee565b03615a7957506107d063ffffffff821603615a4057506107d063ffffffff821603615a075750565b7f0389caa2000000000000000000000000000000000000000000000000000000006000526107d060045263ffffffff1660245260446000fd5b7f22e102a0000000000000000000000000000000000000000000000000000000006000526107d060045263ffffffff1660245260446000fd5b6106e09051615a87816128ee565b7fc2fc5865000000000000000000000000000000000000000000000000000000006000526158f6565b7f77e480260000000000000000000000000000000000000000000000000000000060005263ffffffff9081166004521660245260446000fd5b7fe366a1170000000000000000000000000000000000000000000000000000000060005263ffffffff9081166004521660245260446000fd5b7f68d2f8d60000000000000000000000000000000000000000000000000000000060005263ffffffff1660045260246000fd5b7f758b22cc0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b80548210156132ad5760005260206000200190600090565b600081815260036020526040902054615c2557600254680100000000000000008110156107a057615c0c615bd78260018594016002556002615b82565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600254906000526003602052604060002055600190565b50600090565b600081815260106020526040902054615c2557600f54680100000000000000008110156107a057615c68615bd7826001859401600f55600f615b82565b9055600f54906000526010602052604060002055600190565b600081815260066020526040902054615c2557600554680100000000000000008110156107a057615cbe615bd78260018594016005556005615b82565b9055600554906000526006602052604060002055600190565b6000828152600182016020526040902054615d2c57805490680100000000000000008210156107a05782615d15615bd7846001809601855584615b82565b905580549260005201602052604060002055600190565b5050600090565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1991615f12613142928054615d84615d7e615d758363ffffffff9060801c1690565b63ffffffff1690565b426139e3565b9081615f1e575b5050615ecc6001615daf60208601516fffffffffffffffffffffffffffffffff1690565b92615e3a615dfd61554d6fffffffffffffffffffffffffffffffff615de485546fffffffffffffffffffffffffffffffff1690565b166fffffffffffffffffffffffffffffffff88166166a2565b82906fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b615e8d615e478751151590565b82547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178255565b019182906fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b604083015181546fffffffffffffffffffffffffffffffff1660809190911b7fffffffffffffffffffffffffffffffff0000000000000000000000000000000016179055565b604051918291826154c3565b61554d615dfd916fffffffffffffffffffffffffffffffff615f92615f999582615f8b60018a01549282615f84615f7d615f67876fffffffffffffffffffffffffffffffff1690565b996fffffffffffffffffffffffffffffffff1690565b9560801c90565b1690614231565b91166130e4565b91166166a2565b80547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff00000000000000000000000000000000161781553880615d8b565b60808101615ff4611f63610a6483614bce565b6158a85750602081019061601260206157236156fb6105418661398b565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156106a457600091616098575b5061585f57606061608f61048b9361607e61607960408601614bce565b6166b4565b61135d61608a8261398b565b61674b565b91013590616829565b6160b1915060203d602011614e2657614e1881836107f9565b3861605c565b604051906002548083528260208101600260005260206000209260005b8181106160e957505061048b925003836107f9565b84548352600194850194879450602090930192016160d4565b604051906005548083528260208101600560005260206000209260005b81811061613457505061048b925003836107f9565b845483526001948501948794506020909301920161611f565b906040519182815491828252602082019060005260206000209260005b81811061617f57505061048b925003836107f9565b845483526001948501948794506020909301920161616a565b805480156161fc577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906161cd8282615b82565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600081815260036020526040902054908115615d2c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201908282116128e957600254927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84019384116128e95783836000956162c795036162cd575b5050506162b66002616198565b600390600052602052604060002090565b55600190565b6162b66162f9916162ef6162e56162ff956002615b82565b90549060031b1c90565b9283916002615b82565b90613894565b553880806162a9565b600081815260066020526040902054908115615d2c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201908282116128e957600554927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84019384116128e95783836000956162c795036163a4575b5050506163936005616198565b600690600052602052604060002090565b6163936162f9916163bc6162e56163c6956005615b82565b9283916005615b82565b55388080616386565b60018101918060005282602052604060002054928315156000146164a7577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84018481116128e9578354937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff85019485116128e95760009585836162c79761645f950361646e575b505050616198565b90600052602052604060002090565b61648e6162f9916164856162e561649e9588615b82565b92839187615b82565b8590600052602052604060002090565b55388080616457565b50505050600090565b926164bb9192614231565b81018091116128e957610921916166a2565b6164d9611f6382613995565b6165ba5760206165529161650561104960045473ffffffffffffffffffffffffffffffffffffffff1690565b6040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90921660048301523360248301529092839190829081906044820190565b03915afa9081156106a45760009161659b575b501561656d57565b7f728fe07b000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b6165b4915060203d602011614e2657614e1881836107f9565b38616565565b7fa9902c7e0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b67ffffffffffffffff7f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c91169182600052600760205280616672600260406000200173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169283916168e0565b6040805173ffffffffffffffffffffffffffffffffffffffff909216825260208201929092529081908101614fb2565b90808210156166af575090565b905090565b7f00000000000000000000000000000000000000000000000000000000000000006166dc5750565b73ffffffffffffffffffffffffffffffffffffffff16806000526003602052604060002054156167095750565b7fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b90816020910312610469575161092181610a19565b616757611f6382613995565b6165ba5760206167c89161678361104960045473ffffffffffffffffffffffffffffffffffffffff1690565b60405180809581947fa8d87a3b0000000000000000000000000000000000000000000000000000000083526004830191909167ffffffffffffffff6020820193169052565b03915afa80156106a45773ffffffffffffffffffffffffffffffffffffffff916000916167fa575b5016330361656d57565b61681c915060203d602011616822575b61681481836107f9565b810190616736565b386167f0565b503d61680a565b67ffffffffffffffff7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894491169182600052600760205280616672604060002073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169283916168e0565b81156168b1570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b8054939290919060ff60a086901c16158015616b1b575b616b14576169166fffffffffffffffffffffffffffffffff861661554d565b9060018401958654616950615d7e615d7561694361554d856fffffffffffffffffffffffffffffffff1690565b9460801c63ffffffff1690565b80616a80575b5050838110616a3557508282106169b6575061048b93945061697b9161554d916139e3565b6fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b906169ed6106e0936169e86169d9846169d361554d8c5460801c90565b936139e3565b6169e2836139b6565b906130e4565b6168a7565b7fd0c8d23a0000000000000000000000000000000000000000000000000000000060005260045260245273ffffffffffffffffffffffffffffffffffffffff16604452606490565b7f1a76572a00000000000000000000000000000000000000000000000000000000600052600452602483905273ffffffffffffffffffffffffffffffffffffffff1660445260646000fd5b828592939511616aea57616a9a61554d616aa19460801c90565b91856164b0565b84547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178555913880616956565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050509050565b5081156168f7565b91929015616b9e5750815115616b37575090565b3b15616b405790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015616bb15750805190602001fd5b61201d906040519182917f08c379a00000000000000000000000000000000000000000000000000000000083526004830161091056fea164736f6c634300081a000a2e902d38f15b233cbb63711add0fca4545334d3a169d60c0a616494d7eea9544",
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
