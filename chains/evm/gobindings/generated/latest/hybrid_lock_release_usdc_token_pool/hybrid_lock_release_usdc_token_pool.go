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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"internalType\":\"contractITokenMessenger\"},{\"name\":\"tokenMessengerV2\",\"type\":\"address\",\"internalType\":\"contractITokenMessenger\"},{\"name\":\"cctpMessageTransmitterProxy\",\"type\":\"address\",\"internalType\":\"contractCCTPMessageTransmitterProxy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"previousPool\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"FINALITY_THRESHOLD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_FEE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnLockedUSDC\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelExistingCCTPMigrationProposal\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"excludeTokensFromBurn\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentProposedCCTPChainMigration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDomain\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structUSDCTokenPool.Domain\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExcludedTokensByChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLiquidityProvider\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLockedTokensForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_localDomainIdentifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_messageTransmitterProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractCCTPMessageTransmitterProxy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_previousPool\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_tokenMessenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITokenMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_tokenMessengerCCTPV2\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITokenMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_usdcVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposeCCTPMigration\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"provideLiquidity\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCircleMigratorAddress\",\"inputs\":[{\"name\":\"migrator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDomains\",\"inputs\":[{\"name\":\"domains\",\"type\":\"tuple[]\",\"internalType\":\"structUSDCTokenPool.DomainUpdate[]\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLiquidityProvider\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"liquidityProvider\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRouter\",\"inputs\":[{\"name\":\"newRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shouldUseLockRelease\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateCCTPVersion\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"versions\",\"type\":\"uint8[]\",\"internalType\":\"enumUSDCTokenPool.CCTPVersion[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateChainSelectorMechanisms\",\"inputs\":[{\"name\":\"removes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"adds\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawLiquidity\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCTPMigrationCancelled\",\"inputs\":[{\"name\":\"existingProposalSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCTPMigrationExecuted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"USDCBurned\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCTPMigrationProposed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCTPVersionSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumUSDCTokenPool.CCTPVersion\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CircleMigratorAddressSet\",\"inputs\":[{\"name\":\"migratorAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DomainsSet\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structUSDCTokenPool.DomainUpdate[]\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityAdded\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityProviderSet\",\"inputs\":[{\"name\":\"oldProvider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newProvider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityRemoved\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockReleaseDisabled\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockReleaseEnabled\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MintRecipientOverrideSet\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RouterUpdated\",\"inputs\":[{\"name\":\"oldRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensExcludedFromBurn\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"burnableAmountAfterExclusion\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotSupportedByCCTP\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"ExistingMigrationProposal\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCCTPVersion\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"version\",\"type\":\"uint8\",\"internalType\":\"enumUSDCTokenPool.CCTPVersion\"}]},{\"type\":\"error\",\"name\":\"InvalidChainSelector\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestinationDomain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidDomain\",\"inputs\":[{\"name\":\"domain\",\"type\":\"tuple\",\"internalType\":\"structUSDCTokenPool.DomainUpdate\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"type\":\"error\",\"name\":\"InvalidExecutionFinalityThreshold\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"actual\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMinFinalityThreshold\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"actual\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidNonce\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"got\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourceDomain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenMessengerVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidTransmitterInProxy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LanePausedForCCTPMigration\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoMigrationProposalPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenLockingNotAllowedAfterMigration\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnknownDomain\",\"inputs\":[{\"name\":\"domain\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"UnlockingUSDCFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"onlyCircle\",\"inputs\":[]}]",
	Bin: "0x6101e0806040523461061957617763803803809161001d8285610904565b8339810190610100818303126106195761003681610927565b9161004360208301610927565b60408301516001600160a01b0381169390848103610619576060820151936001600160a01b038516928386036106195760808101516001600160401b0381116106195781019180601f84011215610619578251926001600160401b0384116108ee578360051b9060208201946100bc6040519687610904565b855260208086019282010192831161061957602001905b8282106108d6575050506100e960a08201610927565b61010160e06100fa60c08501610927565b9301610927565b9633156108c557600180546001600160a01b03191633179055851580156108b4575b80156108a3575b6108925760805260c05260405163313ce56760e01b8152602081600481885afa809160009161084f575b509061082b575b50600660a052600480546001600160a01b0319166001600160a01b03929092169190911790558051151560e08190526106fe575b506001600160a01b03861695861561062657604051632c12192160e01b81526020816004818b5afa90811561055d576000916106c4575b5060405163054fd4d560e41b81526001600160a01b039190911690602081600481855afa801561055d5763ffffffff916000916106a5575b50168061057d5750604051639cdbb18160e01b81526020816004818c5afa801561055d5763ffffffff91600091610686575b5016806105695750604051634a48569760e01b81526020816004818b5afa801561055d578291600091610667575b506001600160a01b03160361051d5760049260209261010052610120526040519283809263234d8e3d60e21b82525afa90811561055d57600091610648575b5061014052600061016052608051610100516102c5916001600160a01b039182169116610976565b6000805160206177438339815191526020604051878152a1610180526001600160a01b0381169283158015610637575b61062657604051632c12192160e01b8152602081600481885afa90811561055d576000916105e7575b5060405163054fd4d560e41b81526001600160a01b03919091169190602081600481865afa90811561055d576000916105c8575b50604051639cdbb18160e01b8152906020826004818a5afa91821561055d57600092610591575b5063ffffffff166001810361057d575063ffffffff1660018103610569575060206004916040519283809263025ed2dd60e11b82525afa90811561055d5760009161052e575b506001600160a01b03160361051d57600080516020617743833981519152602061040395604051908152a16101a0526101c0526080516001600160a01b0316610976565b60405161699a9081610da9823960805181818161057601528181610932015281816109b9015281816120930152818161441b015281816145fa01528181614a0801528181614d2d01528181614fbb015281816163f00152616625015260a05181610a03015260c0518181816125f2015281816154d50152615dd1015260e051818181610fd3015281816127df015261645c0152610100518181816111be0152614ddb015261012051818181611ae2015261439901526101405181818161129a01528181614b5a01528181614e750152818161567e015261585a01526101605181818161068601526157f3015261018051816117df01526101a051818181611baf01526149bc01526101c051818181610adc01526132eb0152f35b632a32133b60e11b60005260046000fd5b610550915060203d602011610556575b6105488183610904565b810190610957565b386103bf565b503d61053e565b6040513d6000823e3d90fd5b6316ba39c560e31b60005260045260246000fd5b6334697c6b60e11b60005260045260246000fd5b63ffffffff9192506105ba9060203d6020116105c1575b6105b28183610904565b81019061093b565b9190610379565b503d6105a8565b6105e1915060203d6020116105c1576105b28183610904565b38610352565b90506020813d60201161061e575b8161060260209383610904565b810103126106195761061390610927565b3861031e565b600080fd5b3d91506105f5565b6306b7c75960e31b60005260046000fd5b506001600160a01b038316156102f5565b610661915060203d6020116105c1576105b28183610904565b3861029d565b610680915060203d602011610556576105488183610904565b3861025e565b61069f915060203d6020116105c1576105b28183610904565b38610230565b6106be915060203d6020116105c1576105b28183610904565b386101fe565b90506020813d6020116106f6575b816106df60209383610904565b81010312610619576106f090610927565b386101c6565b3d91506106d2565b909194602094604051946107128787610904565b60008652600036813760e0511561081a5760005b865181101561078d576001906001600160a01b03610744828a610b4e565b51168961075082610b90565b61075d575b505001610726565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a13889610755565b509193969092945060005b855181101561080c576001906001600160a01b036107b68289610b4e565b5116801561080657886107c882610c78565b6107d6575b50505b01610798565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a138886107cd565b506107d0565b50929591945092503861018f565b6335f4a7b360e01b60005260046000fd5b60ff166006811461015b576332ad3e0760e11b600052600660045260245260446000fd5b6020813d60201161088a575b8161086860209383610904565b8101031261088657519060ff82168203610883575038610154565b80fd5b5080fd5b3d915061085b565b6342bcdf7f60e11b60005260046000fd5b506001600160a01b0382161561012a565b506001600160a01b03831615610123565b639b15e16f60e01b60005260046000fd5b602080916108e384610927565b8152019101906100d3565b634e487b7160e01b600052604160045260246000fd5b601f909101601f19168101906001600160401b038211908210176108ee57604052565b51906001600160a01b038216820361061957565b90816020910312610619575163ffffffff811681036106195790565b9081602091031261061957516001600160a01b03811681036106195790565b604051636eb1769f60e11b81523060048201526001600160a01b0392831660248201819052929190911690602081604481855afa90811561055d57600091610b1c575b506000198101809111610b065760405190602082019363095ea7b360e01b855260248301526044820152604481526109f2606482610904565b600080604094855193610a058786610904565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082865af13d15610af9573d906001600160401b0382116108ee578451610a76949092610a67601f8201601f191660200185610904565b83523d6000602085013e610cd8565b805180610a8257505050565b8160209181010312610619576020015180159081150361061957610aa35750565b5162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b91610a7692606091610cd8565b634e487b7160e01b600052601160045260246000fd5b90506020813d602011610b46575b81610b3760209383610904565b810103126106195751386109b9565b3d9150610b2a565b8051821015610b625760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b8054821015610b625760005260206000200190600090565b6000818152600360205260409020548015610c71576000198101818111610b0657600254600019810191908211610b0657818103610c20575b5050506002548015610c0a5760001901610be4816002610b78565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b610c59610c31610c42936002610b78565b90549060031b1c9283926002610b78565b819391549060031b91821b91600019901b19161790565b90556000526003602052604060002055388080610bc9565b5050600090565b80600052600360205260406000205415600014610cd257600254680100000000000000008110156108ee57610cb9610c428260018594016002556002610b78565b9055600254906000526003602052604060002055600190565b50600090565b91929015610d3a5750815115610cec575090565b3b15610cf55790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b825190915015610d4d5750805190602001fd5b6040519062461bcd60e51b8252602060048301528181519182602483015260005b838110610d905750508160006044809484010152601f80199101168101030190fd5b60208282018101516044878401015285935001610d6e56fe6080604052600436101561001257600080fd5b60003560e01c806301ffc9a7146103875780631101dbd4146103825780631604ceea1461037d578063181f5a771461037857806321df0da714610373578063240028e81461036e57806324f65ee7146103695780632cfbb11914610364578063390775371461035f5780633e591f2c1461035a5780633eb68ab7146103555780634ad01f0b146103505780634c5ef0ed1461034b5780634c93ef841461034657806350d1a35a1461034157806354c8a4f31461033c5780636155cda01461033757806362ddd3c4146103325780636b716b0d1461032d5780636b795423146103285780636d3d1a5814610323578063714bf9071461031e57806379ba5097146103195780637d54534e146103145780638926f54f1461030f5780638a5e52bb1461030a5780638da5cb5b14610305578063962d40201461030057806398db9643146102fb5780639a4575b9146102f6578063a1596fb5146102f1578063a42a7b8b146102ec578063a7cd63b7146102e7578063acfecf91146102e2578063af58d59f146102dd578063b0f479a1146102d8578063b7946580146102d3578063bb5eced3146102ce578063bc063e1a146102c9578063c0d78655146102c4578063c4bffe2b146102bf578063c75eea9c146102ba578063c781d0e3146102b5578063cd306a6c146102b0578063cf7401f3146102ab578063da4b05e7146102a6578063dc0bd971146102a1578063de814c571461029c578063dfadfa3514610297578063e0351e1314610292578063e8a1da171461028d578063e94ae6d014610288578063f2fde38b14610283578063f65a88861461027e5763fd6768551461027957600080fd5b612d65565b612d26565b612c50565b612bfb565b612804565b6127c7565b6126fe565b612616565b6125c5565b6125a8565b6124aa565b612385565b612328565b6122dd565b61224a565b612128565b61210c565b61202f565b611ff3565b611fbf565b611f13565b611ddf565b611d6b565b611c69565b611b99565b611b06565b611ab5565b611977565b611912565b6116da565b61169b565b61160a565b61153f565b6114ae565b61147a565b6112be565b61127d565b6111fa565b6111a8565b610fa1565b610e5b565b610e1b565b610e02565b610ccc565b610b83565b610aaf565b610a66565b610a27565b6109e9565b61097f565b610905565b6108a2565b61066e565b61047d565b34610459576020600319360112610459576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361045957807faff2afbf000000000000000000000000000000000000000000000000000000006020921490811561042f575b8115610405575b506040519015158152f35b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014386103fa565b7f0e64dd2900000000000000000000000000000000000000000000000000000000811491506103f3565b600080fd5b67ffffffffffffffff81160361045957565b359061047b8261045e565b565b346104595760406003193601126104595760043561049a8161045e565b602435906104d96104bf8267ffffffffffffffff166000526013602052604060002090565b5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff339116036106355767ffffffffffffffff8116610518816000526012602052604060002054151590565b6105fd57600d5461053e9060a01c67ffffffffffffffff165b67ffffffffffffffff1690565b146105c2576105619067ffffffffffffffff16600052600e602052604060002090565b61056c828254612e4c565b905561059a8130337f0000000000000000000000000000000000000000000000000000000000000000614189565b337fc17cea59c2955cb181b03393209566960365771dbba9dc3d510180e7cb312088600080a3005b7fd0da86c40000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b6000fd5b7f646972460000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b7f8e4a23d6000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b600091031261045957565b346104595760006003193601126104595760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6060810190811067ffffffffffffffff8211176106f457604052565b6106a9565b6080810190811067ffffffffffffffff8211176106f457604052565b6020810190811067ffffffffffffffff8211176106f457604052565b6040810190811067ffffffffffffffff8211176106f457604052565b60a0810190811067ffffffffffffffff8211176106f457604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176106f457604052565b6040519061047b60a083610769565b6040519061047b608083610769565b6040519061047b602083610769565b6040519061047b606083610769565b6040519061047b604083610769565b67ffffffffffffffff81116106f457601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b8481106108795750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b8060208092840101518282860101520161083a565b90602061089f92818152019061082f565b90565b346104595760006003193601126104595761090160408051906108c58183610769565b601782527f55534443546f6b656e506f6f6c20312e362e312d64657600000000000000000060208301525191829160208352602083019061082f565b0390f35b3461045957600060031936011261045957602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b73ffffffffffffffffffffffffffffffffffffffff81160361045957565b359061047b82610956565b346104595760206003193601126104595760206109df6004356109a181610956565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691161490565b6040519015158152f35b3461045957600060031936011261045957602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346104595760206003193601126104595767ffffffffffffffff600435610a4d8161045e565b16600052600e6020526020604060002054604051908152f35b346104595760206003193601126104595760043567ffffffffffffffff811161045957610100600319823603011261045957610aa660209160040161320a565b60405190518152f35b3461045957600060031936011261045957602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b9181601f840112156104595782359167ffffffffffffffff8311610459576020808501948460051b01011161045957565b60406003198201126104595760043567ffffffffffffffff81116104595781610b5c91600401610b00565b929092916024359067ffffffffffffffff821161045957610b7f91600401610b00565b9091565b3461045957610b9136610b31565b919092610b9c61469f565b828203610ca25760005b828110610baf57005b610bba81848461344d565b35610bc48161045e565b610bcf82868861344d565b3590610bda82612f34565b610c02610bfb8267ffffffffffffffff166000526010602052604060002090565b5460ff1690565b610c6b578181610c5360019594610c4e7f74cf5df65e6643e8523827033a1a33c9c370c63e2898443e7f578129ac616da69567ffffffffffffffff166000526014602052604060002090565b61346c565b610c62604051928392836134a3565b0390a101610ba6565b7f0ff1be2d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b7f568efce20000000000000000000000000000000000000000000000000000000060005260046000fd5b3461045957600060031936011261045957610ce561469f565b600d5467ffffffffffffffff8160a01c168015610d7b577fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff7f375f1ad1194a2bec317c5efec05cc63ffa06ddd0c4b276619f6fd47298eda5189216600d556000610d638267ffffffffffffffff16600052600f602052604060002090565b5560405167ffffffffffffffff9091168152602090a1005b7fa94cb9880000000000000000000000000000000000000000000000000000000060005260046000fd5b604060031982011261045957600435610dbd8161045e565b9160243567ffffffffffffffff811161045957826023820112156104595780600401359267ffffffffffffffff84116104595760248483010111610459576024019190565b346104595760206109df610e1536610da5565b91613537565b346104595760206003193601126104595760206109df600435610e3d8161045e565b67ffffffffffffffff16600052601060205260ff6040600020541690565b3461045957602060031936011261045957600435610e788161045e565b610e8061469f565b67ffffffffffffffff600d5460a01c16610f775767ffffffffffffffff811660009081526010602052604090205460ff1615610f4d57610f4881610f2d7f20331f191af84dbff48b162aa5a5985e7891ae646297b0a2ac80487f9109ef49937fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff7bffffffffffffffff0000000000000000000000000000000000000000600d549260a01b16911617600d55565b60405167ffffffffffffffff90911681529081906020820190565b0390a1005b7f656535ce0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f692bc1310000000000000000000000000000000000000000000000000000000060005260046000fd5b3461045957610fc9610fd1610fb536610b31565b9491610fc293919361469f565b3691613597565b923691613597565b7f00000000000000000000000000000000000000000000000000000000000000001561117e5760005b82518110156110bf578061102d6110136001938661372f565b5173ffffffffffffffffffffffffffffffffffffffff1690565b61106961106473ffffffffffffffffffffffffffffffffffffffff83165b73ffffffffffffffffffffffffffffffffffffffff1690565b615fd1565b611075575b5001610ffa565b60405173ffffffffffffffffffffffffffffffffffffffff9190911681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756690602090a13861106e565b5060005b815181101561117c57806110dc6110136001938561372f565b73ffffffffffffffffffffffffffffffffffffffff8116156111765761111f61111a73ffffffffffffffffffffffffffffffffffffffff831661104b565b615940565b61112c575b505b016110c3565b60405173ffffffffffffffffffffffffffffffffffffffff9190911681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d890602090a183611124565b50611126565b005b7f35f4a7b30000000000000000000000000000000000000000000000000000000060005260046000fd5b34610459576000600319360112610459576040517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168152602090f35b346104595761120836610da5565b61121392919261469f565b67ffffffffffffffff8216611235816000526006602052604060002054151590565b15611250575061117c9261124a913691613500565b9061470b565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b3461045957600060031936011261045957602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610459576112cc36610b31565b9290916112d761469f565b60005b8181106113ff5750505060005b8281106112f057005b61131c61130961053161130484878761344d565b613462565b6000526012602052604060002054151590565b6113b8578061137c611351611337611304600195888861344d565b67ffffffffffffffff166000526010602052604060002090565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055565b61138d61053161130483878761344d565b7f5e3985e51df58346365017cae614e59d723143b71c9a2ce4a156687f1f2c3f5a600080a2016112e7565b611304906105f9936113c99361344d565b7f646972460000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b8061143e611416611337611304600195878961344d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008154169055565b61144f61053161130483868861344d565b7fddc5afbc5e53c63a556964db0eef76a1c2d9305e0811abd7410d2a6f4799490e600080a2016112da565b3461045957600060031936011261045957602073ffffffffffffffffffffffffffffffffffffffff60095416604051908152f35b34610459576020600319360112610459577f084e6f0e9791c2e56153bd49e6ec6dd63ba9a72c258d71558d74c63fc75b7168602073ffffffffffffffffffffffffffffffffffffffff60043561150381610956565b61150b61469f565b16807fffffffffffffffffffffffff0000000000000000000000000000000000000000600d541617600d55604051908152a1005b346104595760006003193601126104595760005473ffffffffffffffffffffffffffffffffffffffff811633036115e0577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b34610459576020600319360112610459577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174602073ffffffffffffffffffffffffffffffffffffffff60043561165f81610956565b61166761469f565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000006009541617600955604051908152a1005b346104595760206003193601126104595760206109df67ffffffffffffffff6004356116c68161045e565b166000526006602052604060002054151590565b3461045957600060031936011261045957600d5461170d73ffffffffffffffffffffffffffffffffffffffff821661104b565b33036118e85760a01c67ffffffffffffffff1667ffffffffffffffff8116908115610d7b576117796117538267ffffffffffffffff16600052600e602052604060002090565b546117728367ffffffffffffffff16600052600f602052604060002090565b549061363b565b90600061179a8267ffffffffffffffff16600052600e602052604060002090565b556117c87fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff600d5416600d55565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692833b1561045957600060405180957f42966c6800000000000000000000000000000000000000000000000000000000825281838161184689600483019190602083019252565b03925af19081156118e3577fdea60ddd4c7ebdab804f5694c70350cca7893ece3efeecb142312eacac5c73e4946118a3926118c8575b5061189e6114168467ffffffffffffffff166000526010602052604060002090565b6159d1565b506040805167ffffffffffffffff909216825260208201929092529081908101610f48565b806118d760006118dd93610769565b80610663565b3861187c565b6131fe565b7f438a7a050000000000000000000000000000000000000000000000000000000060005260046000fd5b3461045957600060031936011261045957602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b9181601f840112156104595782359167ffffffffffffffff8311610459576020808501946060850201011161045957565b346104595760606003193601126104595760043567ffffffffffffffff8111610459576119a8903690600401610b00565b9060243567ffffffffffffffff8111610459576119c9903690600401611946565b9060443567ffffffffffffffff8111610459576119ea903690600401611946565b611a0c61104b60095473ffffffffffffffffffffffffffffffffffffffff1690565b33141580611a8a575b61063557838614801590611a80575b610ca25760005b868110611a3457005b80611a7a611a486113046001948b8b61344d565b611a53838989613648565b611a74611a6c611a6486898b613648565b923690612461565b913690612461565b916147d0565b01611a2b565b5080861415611a24565b50611aad61104b60015473ffffffffffffffffffffffffffffffffffffffff1690565b331415611a15565b3461045957600060031936011261045957602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346104595760206003193601126104595760043567ffffffffffffffff81116104595760a0600319823603011261045957611b4661090191600401613671565b604051918291602083526020611b678251604083870152606086019061082f565b9101517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe084830301604085015261082f565b34610459576000600319360112610459576040517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168152602090f35b602081016020825282518091526040820191602060408360051b8301019401926000915b838310611c1e57505050505090565b9091929394602080611c5a837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08660019603018752895161082f565b97019301930191939290611c0f565b346104595760206003193601126104595767ffffffffffffffff600435611c8f8161045e565b166000526007602052611ca86005604060002001615ef3565b8051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611cee611cd88461357f565b93611ce66040519586610769565b80855261357f565b0160005b818110611d5a57505060005b8151811015611d4c5780611d30611d2b611d1a6001948661372f565b516000526008602052604060002090565b613796565b611d3a828661372f565b52611d45818561372f565b5001611cfe565b604051806109018582611beb565b806060602080938701015201611cf2565b3461045957600060031936011261045957611d84615e5d565b60405180916020820160208352815180915260206040840192019060005b818110611db0575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101611da2565b3461045957611ded36610da5565b611df892919261469f565b67ffffffffffffffff821691611e22611e1e846000526006602052604060002054151590565b1590565b611edc57611e65611e1e6005611e4c8467ffffffffffffffff166000526007602052604060002090565b01611e58368689613500565b6020815191012090616175565b611ea157507f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d769192611e9c60405192839283613877565b0390a2005b611ed884926040519384937f74f23c7c00000000000000000000000000000000000000000000000000000000855260048501613856565b0390fd5b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b346104595760206003193601126104595767ffffffffffffffff600435611f398161045e565b611f41613888565b50166000526007602052610901611f66611f6160026040600020016138b3565b615043565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b3461045957600060031936011261045957602073ffffffffffffffffffffffffffffffffffffffff60045416604051908152f35b346104595760206003193601126104595761090161201b6004356120168161045e565b61390d565b60405191829160208352602083019061082f565b346104595760406003193601126104595760043561204c8161045e565b6024359061205861469f565b67ffffffffffffffff80600d5460a01c1691169081146120df57600052600e602052604060002061208a82825461363b565b90556120b781337f00000000000000000000000000000000000000000000000000000000000000006150e5565b337fc2c3f06e49b9f15e7b4af9055e183b0d73362e033ad82a07dec9bf9840171719600080a3005b7fd0da86c40000000000000000000000000000000000000000000000000000000060005260045260246000fd5b3461045957600060031936011261045957602060405160008152f35b346104595760206003193601126104595773ffffffffffffffffffffffffffffffffffffffff60043561215a81610956565b61216261469f565b1680156121dc5760407f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f16849160045490807fffffffffffffffffffffffff000000000000000000000000000000000000000083161760045573ffffffffffffffffffffffffffffffffffffffff8351921682526020820152a1005b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b602060408183019282815284518094520192019060005b81811061222a5750505090565b825167ffffffffffffffff1684526020938401939092019160010161221d565b3461045957600060031936011261045957612263615ea8565b8051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612293611cd88461357f565b0136602084013760005b81518110156122cf578067ffffffffffffffff6122bc6001938561372f565b51166122c8828661372f565b520161229d565b604051806109018582612206565b346104595760206003193601126104595767ffffffffffffffff6004356123038161045e565b61230b613888565b50166000526007602052610901611f66611f6160406000206138b3565b346104595760206003193601126104595760043567ffffffffffffffff8111610459573660238201121561045957806004013567ffffffffffffffff81116104595736602460a083028401011161045957602461117c920161392f565b3461045957600060031936011261045957602067ffffffffffffffff600d5460a01c16604051908152f35b8015150361045957565b35906fffffffffffffffffffffffffffffffff8216820361045957565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c6060910112610459576040519061240e826106d8565b8160843561241b816123b0565b815260a4356fffffffffffffffffffffffffffffffff8116810361045957602082015260c435906fffffffffffffffffffffffffffffffff821682036104595760400152565b919082606091031261045957604051612479816106d8565b60406124a5818395803561248c816123b0565b855261249a602082016123ba565b6020860152016123ba565b910152565b346104595760e0600319360112610459576004356124c78161045e565b60607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc360112610459576040516124fd816106d8565b602435612509816123b0565b81526044356fffffffffffffffffffffffffffffffff811681036104595760208201526064356fffffffffffffffffffffffffffffffff81168103610459576040820152612556366123d7565b9073ffffffffffffffffffffffffffffffffffffffff6009541633141580612586575b6106355761117c926147d0565b5073ffffffffffffffffffffffffffffffffffffffff60015416331415612579565b346104595760006003193601126104595760206040516107d08152f35b3461045957600060031936011261045957602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610459576040600319360112610459576004356126338161045e565b6024359061263f61469f565b67ffffffffffffffff600d5460a01c169167ffffffffffffffff8216809303610d7b5782600052600f60205260406000208054928284018094116126f9577fe1e6c22ce6b566f66cdb457ec2e7910ff1f9a9e5654ed75303476fa870468220936126e192556117726126c58267ffffffffffffffff16600052600e602052604060002090565b549167ffffffffffffffff16600052600f602052604060002090565b60408051928352602083019190915281908101611e9c565b612e1d565b346104595760206003193601126104595767ffffffffffffffff6004356127248161045e565b60006060604051612734816106f9565b828152826020820152826040820152015216600052600a602052610901604060002060ff600260405192612767846106f9565b8054845260018101546020850152015463ffffffff8116604084015260201c1615156060820152604051918291829190916060806080830194805184526020810151602085015263ffffffff604082015116604085015201511515910152565b346104595760006003193601126104595760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b346104595761281236610b31565b91909261281d61469f565b6000915b808310612ac25750505060009163ffffffff4216925b82811061284057005b61285361284e828585613d58565b613e32565b9060608201612862815161517f565b6080830193612871855161517f565b6040840190815151156121dc5761289e611e1e612899610531885167ffffffffffffffff1690565b615a27565b612a77576129d76128d76128bd879a999a5167ffffffffffffffff1690565b67ffffffffffffffff166000526007602052604060002090565b61299a89612994875161297b61290060408301516fffffffffffffffffffffffffffffffff1690565b9161296261292b61292460208401516fffffffffffffffffffffffffffffffff1690565b9251151590565b6129596129366107aa565b6fffffffffffffffffffffffffffffffff851681529763ffffffff166020890152565b15156040870152565b6fffffffffffffffffffffffffffffffff166060850152565b6fffffffffffffffffffffffffffffffff166080830152565b82613ec1565b6129cc896129c38a5161297b61290060408301516fffffffffffffffffffffffffffffffff1690565b60028301613ec1565b600484519101613fcd565b602085019660005b88518051821015612a1a5790612a14600192612a0d83612a078c5167ffffffffffffffff1690565b9261372f565b519061470b565b016129df565b50509796509490612a6e7f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c29392612a5b6001975167ffffffffffffffff1690565b92519351905190604051948594856140f4565b0390a101612837565b6105f9612a8c865167ffffffffffffffff1690565b7f1d5ad3c50000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b909192612ad361130485848661344d565b94612aea611e1e67ffffffffffffffff88166160ae565b612bc357612b176005612b118867ffffffffffffffff166000526007602052604060002090565b01615ef3565b9360005b8551811015612b6357600190612b5c6005612b4a8b67ffffffffffffffff166000526007602052604060002090565b01612b55838a61372f565b5190616175565b5001612b1b565b509350937f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916612bb560019397610f2d612bb08267ffffffffffffffff166000526007602052604060002090565b613ca9565b0390a1019190939293612821565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff861660045260246000fd5b346104595760206003193601126104595767ffffffffffffffff600435612c218161045e565b166000526013602052602073ffffffffffffffffffffffffffffffffffffffff60406000205416604051908152f35b346104595760206003193601126104595773ffffffffffffffffffffffffffffffffffffffff600435612c8281610956565b612c8a61469f565b16338114612cfc57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346104595760206003193601126104595767ffffffffffffffff600435612d4c8161045e565b16600052600f6020526020604060002054604051908152f35b3461045957604060031936011261045957600435612d828161045e565b67ffffffffffffffff60243591612d9883610956565b612da061469f565b166000818152601360205260408120805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff0000000000000000000000000000000000000000821681179092559293909216907fc82aa48e67c70b1ad1494533456f52504bb4d62d11bbdafaeb98cfccd1ed817e9080a4005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b919082018092116126f957565b60405190612e6682610715565b60008252565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610459570180359067ffffffffffffffff82116104595760200191813603831361045957565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110612ef1575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b359063ffffffff8216820361045957565b6003111561045957565b90816060910312610459576040805191612f57836106d8565b8035612f628161045e565b8352612f7060208201612f23565b60208401520135612f8081612f34565b604082015290565b60031115612f9257565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b906003821015612f925752565b9190602461047b9163ffffffff60449516600452612fc1565b908160209103126104595760405190612fff82610715565b51815290565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561045957016020813591019167ffffffffffffffff821161045957813603831361045957565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b9061089f91602081526131cd6131c26131856130c66130b38680613005565b6101006020880152610120870191613055565b6130e66130d560208801610470565b67ffffffffffffffff166040870152565b6131126130f560408801610974565b73ffffffffffffffffffffffffffffffffffffffff166060870152565b6060860135608086015261314861312b60808801610974565b73ffffffffffffffffffffffffffffffffffffffff1660a0870152565b61315560a0870187613005565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08784030160c0880152613055565b61319260c0860186613005565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08684030160e0870152613055565b9260e0810190613005565b916101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082860301910152613055565b6040513d6000823e3d90fd5b90613213612e59565b5060c082017ffa7c07de000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000061326d6132678487612e6c565b90612ebd565b160361327f575b509061089f9061451e565b604061328b8285612e6c565b90501460001461334857509060206132d091604051809381927f3907753700000000000000000000000000000000000000000000000000000000835260048301613094565b0381600073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af19081156118e35760009161331f575090565b61089f915060203d602011613341575b6133398183610769565b810190612fe7565b503d61332f565b61335561335d9184612e6c565b810190612f3e565b91604083016001815161336f81612f88565b61337881612f88565b03613389575061089f9192506144de565b6002815161339681612f88565b61339f81612f88565b036133b0575061089f9192506142ea565b80516133bb81612f88565b6133c481612f88565b156133d45750915061089f613274565b6105f9906133e9602086015163ffffffff1690565b9051906133f582612f88565b7f4f30cd0c00000000000000000000000000000000000000000000000000000000600052612fce565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b919081101561345d5760051b0190565b61341e565b3561089f8161045e565b906003811015612f925760ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008354169116179055565b91602061047b92949367ffffffffffffffff60408201961681520190612fc1565b916134fc918354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055565b92919261350c826107f5565b9161351a6040519384610769565b829481845281830111610459578281602093846000960137010152565b61089f929167ffffffffffffffff613562921660005260076020526005604060002001923691613500565b602081519101209060019160005201602052604060002054151590565b67ffffffffffffffff81116106f45760051b60200190565b9291906135a38161357f565b936135b16040519586610769565b602085838152019160051b810192831161045957905b8282106135d357505050565b6020809183356135e281610956565b8152019101906135c7565b67ffffffffffffffff61089f91166000526006602052604060002054151590565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116126f957565b919082039182116126f957565b919081101561345d576060020190565b6040519061366582610731565b60606020838281520152565b613679613658565b506020810161368d611e1e610e3d83613462565b6136cf575b6136b56105316136af600d5467ffffffffffffffff9060a01c1690565b92613462565b67ffffffffffffffff8216146105c2575061089f90614f45565b6136f8610bfb6136de83613462565b67ffffffffffffffff166000526014602052604060002090565b61370181612f88565b6001810361371457505061089f90614c9d565b80613720600292612f88565b03613692575061089f90614944565b805182101561345d5760209160051b010190565b90600182811c9216801561378c575b602083101461375d57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691613752565b90604051918260008254926137aa84613743565b808452936001811690811561381657506001146137cf575b5061047b92500383610769565b90506000929192526020600020906000915b8183106137fa57505090602061047b92820101386137c2565b60209193508060019154838589010152019101909184926137e1565b6020935061047b9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b820101386137c2565b60409067ffffffffffffffff61089f95931681528160208201520191613055565b91602061089f938181520191613055565b604051906138958261074d565b60006080838281528260208201528260408201528260608201520152565b906040516138c08161074d565b60806fffffffffffffffffffffffffffffffff6001839560ff8154848116875263ffffffff81871c16602088015260a01c1615156040860152015481808216166060850152821c16910152565b67ffffffffffffffff16600052600760205261089f6004604060002001613796565b61393761469f565b60005b8281106139795750907fe6d14ea297366c7bc1265d289d924bfd8b9afb148eb972b481f70da41c842cf59161397460405192839283613b86565b0390a1565b61398c613987828585613b18565b613b28565b8051158015613af2575b613a855790613a7f82613a25613a0b60606139ba6040600198015163ffffffff1690565b936139fc60208201516139f48351976139d66080860151151590565b926139df6107b9565b998a5260208a015263ffffffff166040890152565b151586840152565b015167ffffffffffffffff1690565b67ffffffffffffffff16600052600a602052604060002090565b6002908251815560208301516001820155019063ffffffff6040820151167fffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000064ff0000000060608554940151151560201b16921617179055565b0161393a565b604080517fa606c63500000000000000000000000000000000000000000000000000000000815282516004820152602083015160248201529082015163ffffffff166044820152606082015167ffffffffffffffff1660648201526080909101511515608482015260a490fd5b5067ffffffffffffffff613b11606083015167ffffffffffffffff1690565b1615613996565b919081101561345d5760a0020190565b60a08136031261045957608060405191613b418361074d565b8035835260208101356020840152613b5b60408201612f23565b60408401526060810135613b6e8161045e565b60608401520135613b7e816123b0565b608082015290565b602080825281018390526040019160005b818110613ba45750505090565b90919260a080600192863581526020870135602082015263ffffffff613bcc60408901612f23565b16604082015267ffffffffffffffff6060880135613be98161045e565b1660608201526080870135613bfd816123b0565b15156080820152019401929101613b97565b90604051613c1c816106f9565b606060ff600283958054855260018101546020860152015463ffffffff8116604085015260201c161515910152565b818110613c56575050565b60008155600101613c4b565b818102929181159184041417156126f957565b8054906000815581613c85575050565b6000526020600020908101905b818110613c9d575050565b60008155600101613c92565b600561047b916000815560006001820155600060028201556000600382015560048101613cd68154613743565b9081613ce5575b505001613c75565b81601f60009311600114613cfd5750555b3880613cdd565b81835260208320613d1891601f01861c810190600101613c4b565b808252602082209081548360011b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8560031b1c191617905555613cf6565b919081101561345d5760051b810135907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee181360301821215610459570190565b9080601f830112156104595781602061089f93359101613500565b9080601f83011215610459578135613dca8161357f565b92613dd86040519485610769565b81845260208085019260051b820101918383116104595760208201905b838210613e0457505050505090565b813567ffffffffffffffff811161045957602091613e2787848094880101613d98565b815201910190613df5565b610120813603126104595760405190613e4a8261074d565b613e5381610470565b8252602081013567ffffffffffffffff811161045957613e769036908301613db3565b602083015260408101359067ffffffffffffffff821161045957613ea0613b7e9236908301613d98565b6040840152613eb23660608301612461565b606084015260c0369101612461565b8151815460208401516040850151608091821b73ffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffff000000000000000000000000000000000000000000909416939093179290921791151560a01b74ff000000000000000000000000000000000000000016919091178355606084015193810151901b7fffffffffffffffffffffffffffffffff0000000000000000000000000000000016921691909117600190910155565b9190601f8111613f9757505050565b61047b926000526020600020906020601f840160051c83019310613fc3575b601f0160051c0190613c4b565b9091508190613fb6565b919091825167ffffffffffffffff81116106f457613ff581613fef8454613743565b84613f88565b6020601f821160011461404f5781906134fc939495600092614044575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b015190503880614012565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169061408284600052602060002090565b9160005b8181106140dc575095836001959697106140a5575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c1916905538808061409b565b9192602060018192868b015181550194019201614086565b61415861412361047b9597969467ffffffffffffffff60a095168452610100602085015261010084019061082f565b9660408301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b01906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff9283166024820152929091166044830152606482019290925261047b9161421982608481015b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101845283610769565b61534b565b6020818303126104595780359067ffffffffffffffff82116104595701604081830312610459576040519161425283610731565b813567ffffffffffffffff8111610459578161426f918401613d98565b8352602082013567ffffffffffffffff81116104595761428f9201613d98565b602082015290565b90816020910312610459575161089f816123b0565b9392906142d861047b936142ca60409360608952606089019061082f565b90878203602089015261082f565b940190612fc1565b3561089f81610956565b6142f2612e59565b506142fc81615407565b61430c61335560c0830183612e6c565b602061432661431e60e0850185612e6c565b81019061421e565b9161433281845161562d565b60408284519401519101519261434784612f88565b61437e60405194859384937f0e30e01a000000000000000000000000000000000000000000000000000000008552600485016142ac565b0381600073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af19081156118e3576000916144af575b5015614485576143dc60208201613462565b907ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc067ffffffffffffffff6060614415604085016142e0565b604080517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff9081168252336020830152909216908201529301356060840181905293169180608081015b0390a261447f6107c8565b90815290565b7fbf969f220000000000000000000000000000000000000000000000000000000060005260046000fd5b6144d1915060203d6020116144d7575b6144c98183610769565b810190614297565b386143ca565b503d6144bf565b6144e6612e59565b506144f081615407565b61450061335560c0830183612e6c565b602061451261431e60e0850185612e6c565b916143328184516157eb565b614526612e59565b5061453081615407565b600d5460a01c67ffffffffffffffff16602082019061455161053183613462565b67ffffffffffffffff8216146105c25750907ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc067ffffffffffffffff836145b761459d61130496613462565b67ffffffffffffffff16600052600e602052604060002090565b546146845760608401356145f26145ea6145d084613462565b67ffffffffffffffff16600052600f602052604060002090565b91825461363b565b90555b6144747f00000000000000000000000000000000000000000000000000000000000000009461464361463d60408301946060614630876142e0565b940135998a80958b6150e5565b936142e0565b6040805173ffffffffffffffffffffffffffffffffffffffff9889168152336020820152919097169681019690965260608601529116929081906080820190565b60608401356146986145ea61459d84613462565b90556145f5565b73ffffffffffffffffffffffffffffffffffffffff6001541633036146c057565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b60409067ffffffffffffffff61089f9493168152816020820152019061082f565b908051156121dc578051602082012067ffffffffffffffff831692836000526007602052614740826005604060002001615a7d565b156147995750816147887f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea93614783614794946000526008602052604060002090565b613fcd565b6040519182918261088e565b0390a2565b9050611ed86040519283927f393b8ad2000000000000000000000000000000000000000000000000000000008452600484016146ea565b67ffffffffffffffff1660008181526006602052604090205490929190156148d257916148cf60e09261489b856148277f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b9761517f565b84600052600760205261483e816040600020615ad9565b6148478361517f565b846000526007602052614861836002604060002001615ad9565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b90816020910312610459573590565b815167ffffffffffffffff16815260208083015163ffffffff169082015260409182015160608201939261047b920190612fc1565b61494c613658565b5061495681615d87565b6020810161496e614969613a0b83613462565b613c0f565b9061497f611e1e6060840151151590565b614c4657602061498f8480612e6c565b905003614c065760208201518015614bea57915b606073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016940135926149f1604083015163ffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016925195803b15610459576040517fd04857b00000000000000000000000000000000000000000000000000000000081526004810187905263ffffffff929092166024830152604482019290925273ffffffffffffffffffffffffffffffffffffffff831660648201526084810195909552600060a486018190526107d060c487015290859060e490829084905af19384156118e35767ffffffffffffffff612016947ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1092614b4197614bd5575b50614b39614b0286613462565b6040805173ffffffffffffffffffffffffffffffffffffffff90971687523360208801528601929092529116929081906060820190565b0390a2613462565b614b96614bc2614b4f6107d7565b6000815263ffffffff7f0000000000000000000000000000000000000000000000000000000000000000166020820152600260408201525b6040519283916020830161490f565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610769565b614bca6107e6565b918252602082015290565b806118d76000614be493610769565b38614af5565b50614c00614bf88480612e6c565b810190614900565b916149a3565b614c108380612e6c565b90611ed86040519283927fa3c8cf0900000000000000000000000000000000000000000000000000000000845260048401613877565b614c526105f991613462565b7fd201c48a0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b90816020910312610459575161089f8161045e565b90614ca6613658565b50614cb082615d87565b60208201614cc3614969613a0b83613462565b614cd3611e1e6060830151151590565b614f39576020614ce38580612e6c565b905003614f2f5760208101519293614dc0938015614f17576060602091925b013591614d16604085015163ffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169687955160405198899485947ff856ddb60000000000000000000000000000000000000000000000000000000086528860048701919360809363ffffffff73ffffffffffffffffffffffffffffffffffffffff9398979660a0860199865216602085015260408401521660608201520152565b0381600073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af19384156118e357600094614ea6575b5061201683614bc2937ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff614e5295614b39614b02614b969a613462565b92614e6e614e5e6107d7565b67ffffffffffffffff9092168252565b63ffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015260016040820152614b87565b614e52919450614b9693614bc2937ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff614f026120169560203d602011614f10575b614efa8183610769565b810190614c88565b989550505093509350614e0c565b503d614ef0565b5060206060614f29614bf88480612e6c565b92614d02565b614c108480612e6c565b6105f9614c5283613462565b61201661500b91614f54613658565b50614f5e81615d87565b602081019060600135614f75823561459d8161045e565b614f80828254612e4c565b90557ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff614fb584613462565b604080517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681523360208201529081019490945216918060608101614b39565b604051614bc281614b9660208201907ffa7c07de00000000000000000000000000000000000000000000000000000000602083019252565b61504b613888565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff82511690602083019163ffffffff83511642034281116126f9576150af906fffffffffffffffffffffffffffffffff60808701511690613c62565b81018091116126f9576150d56fffffffffffffffffffffffffffffffff92918392616448565b161682524263ffffffff16905290565b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff9092166024830152604482019290925261047b9161421982606481016141ed565b61047b9092919260608101936fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b8051156152235760408101516fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff6151e36151ce60208501516fffffffffffffffffffffffffffffffff1690565b6fffffffffffffffffffffffffffffffff1690565b9116116151ed5750565b611ed8906040519182917f8020d12400000000000000000000000000000000000000000000000000000000835260048301615144565b6fffffffffffffffffffffffffffffffff61525160408301516fffffffffffffffffffffffffffffffff1690565b1615801590615298575b6152625750565b611ed8906040519182917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301615144565b506152b96151ce60208301516fffffffffffffffffffffffffffffffff1690565b151561525b565b156152c757565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b9073ffffffffffffffffffffffffffffffffffffffff6153d992166040906000808351946153798587610769565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602087015260208151910182855af1903d156153fe573d6153ca6153c1826107f5565b94519485610769565b83523d6000602085013e6168c9565b8051806153e4575050565b816020806153f99361047b9501019101614297565b6152c0565b606092506168c9565b906080820161541b611e1e6109a1836142e0565b6155df575060208201916154bc602061546161543961053187613462565b60801b7fffffffffffffffffffffffffffffffff000000000000000000000000000000001690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116600482015291829081906024820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156118e3576000916155c0575b506155965761551c61551784613462565b616273565b61552583613462565b61553a611e1e60a0840192610e158486612e6c565b6155575750606061554e61047b9394613462565b91013590616397565b61556091612e6c565b90611ed86040519283927f24eb47e500000000000000000000000000000000000000000000000000000000845260048401613877565b7f53ad11d80000000000000000000000000000000000000000000000000000000060005260046000fd5b6155d9915060203d6020116144d7576144c98183610769565b38615506565b6155eb6105f9916142e0565b7f961c9a4f0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b6004810151600163ffffffff8216036157b857506008810151600c8201519061566760206094609086015195015195015163ffffffff1690565b63ffffffff811663ffffffff83160361577f5750507f000000000000000000000000000000000000000000000000000000000000000063ffffffff811663ffffffff8316036157465750506107d063ffffffff82160361570d57506107d063ffffffff8216036156d45750565b7f0389caa2000000000000000000000000000000000000000000000000000000006000526107d060045263ffffffff1660245260446000fd5b7f22e102a0000000000000000000000000000000000000000000000000000000006000526107d060045263ffffffff1660245260446000fd5b7f77e480260000000000000000000000000000000000000000000000000000000060005263ffffffff9081166004521660245260446000fd5b7fe366a1170000000000000000000000000000000000000000000000000000000060005263ffffffff9081166004521660245260446000fd5b7f68d2f8d60000000000000000000000000000000000000000000000000000000060005263ffffffff1660045260246000fd5b9060048201517f000000000000000000000000000000000000000000000000000000000000000063ffffffff8216036157b857506008820151916014600c820151910151906020830193615843855163ffffffff1690565b63ffffffff811663ffffffff83160361577f5750507f000000000000000000000000000000000000000000000000000000000000000063ffffffff811663ffffffff831603615746575050815167ffffffffffffffff1667ffffffffffffffff811667ffffffffffffffff8316036158eb57505060400190600182516158c881612f88565b6158d181612f88565b036158da575050565b516105f9919063ffffffff166133e9565b7ff917ffea0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff9081166004521660245260446000fd5b805482101561345d5760005260206000200190600090565b6000818152600360205260409020546159cb57600254680100000000000000008110156106f4576159b261597d8260018594016002556002615928565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600254906000526003602052604060002055600190565b50600090565b6000818152601260205260409020546159cb57601154680100000000000000008110156106f457615a0e61597d8260018594016011556011615928565b9055601154906000526012602052604060002055600190565b6000818152600660205260409020546159cb57600554680100000000000000008110156106f457615a6461597d8260018594016005556005615928565b9055600554906000526006602052604060002055600190565b6000828152600182016020526040902054615ad257805490680100000000000000008210156106f45782615abb61597d846001809601855584615928565b905580549260005201602052604060002055600190565b5050600090565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1991615cb8613974928054615b2a615b24615b1b8363ffffffff9060801c1690565b63ffffffff1690565b4261363b565b9081615cc4575b5050615c726001615b5560208601516fffffffffffffffffffffffffffffffff1690565b92615be0615ba36151ce6fffffffffffffffffffffffffffffffff615b8a85546fffffffffffffffffffffffffffffffff1690565b166fffffffffffffffffffffffffffffffff8816616448565b82906fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b615c33615bed8751151590565b82547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178255565b019182906fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b604083015181546fffffffffffffffffffffffffffffffff1660809190911b7fffffffffffffffffffffffffffffffff0000000000000000000000000000000016179055565b60405191829182615144565b6151ce615ba3916fffffffffffffffffffffffffffffffff615d38615d3f9582615d3160018a01549282615d2a615d23615d0d876fffffffffffffffffffffffffffffffff1690565b996fffffffffffffffffffffffffffffffff1690565b9560801c90565b1690613c62565b9116612e4c565b9116616448565b80547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff00000000000000000000000000000000161781553880615b31565b60808101615d9a611e1e6109a1836142e0565b6155df57506020810190615db8602061546161543961053186613462565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156118e357600091615e3e575b50615596576060615e3561047b93615e24615e1f604086016142e0565b61645a565b611304615e3082613462565b6164f1565b910135906165cf565b615e57915060203d6020116144d7576144c98183610769565b38615e02565b604051906002548083528260208101600260005260206000209260005b818110615e8f57505061047b92500383610769565b8454835260019485019487945060209093019201615e7a565b604051906005548083528260208101600560005260206000209260005b818110615eda57505061047b92500383610769565b8454835260019485019487945060209093019201615ec5565b906040519182815491828252602082019060005260206000209260005b818110615f2557505061047b92500383610769565b8454835260019485019487945060209093019201615f10565b80548015615fa2577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190615f738282615928565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600081815260036020526040902054908115615ad2577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201908282116126f957600254927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84019384116126f957838360009561606d9503616073575b50505061605c6002615f3e565b600390600052602052604060002090565b55600190565b61605c61609f9161609561608b6160a5956002615928565b90549060031b1c90565b9283916002615928565b906134c4565b5538808061604f565b600081815260066020526040902054908115615ad2577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201908282116126f957600554927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84019384116126f957838360009561606d950361614a575b5050506161396005615f3e565b600690600052602052604060002090565b61613961609f9161616261608b61616c956005615928565b9283916005615928565b5538808061612c565b600181019180600052826020526040600020549283151560001461624d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84018481116126f9578354937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff85019485116126f957600095858361606d976162059503616214575b505050615f3e565b90600052602052604060002090565b61623461609f9161622b61608b6162449588615928565b92839187615928565b8590600052602052604060002090565b553880806161fd565b50505050600090565b926162619192613c62565b81018091116126f95761089f91616448565b61627f611e1e826135ed565b6163605760206162f8916162ab61104b60045473ffffffffffffffffffffffffffffffffffffffff1690565b6040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90921660048301523360248301529092839190829081906044820190565b03915afa9081156118e357600091616341575b501561631357565b7f728fe07b000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b61635a915060203d6020116144d7576144c98183610769565b3861630b565b7fa9902c7e0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b67ffffffffffffffff7f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c91169182600052600760205280616418600260406000200173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391616686565b6040805173ffffffffffffffffffffffffffffffffffffffff909216825260208201929092529081908101614794565b9080821015616455575090565b905090565b7f00000000000000000000000000000000000000000000000000000000000000006164825750565b73ffffffffffffffffffffffffffffffffffffffff16806000526003602052604060002054156164af5750565b7fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b90816020910312610459575161089f81610956565b6164fd611e1e826135ed565b61636057602061656e9161652961104b60045473ffffffffffffffffffffffffffffffffffffffff1690565b60405180809581947fa8d87a3b0000000000000000000000000000000000000000000000000000000083526004830191909167ffffffffffffffff6020820193169052565b03915afa80156118e35773ffffffffffffffffffffffffffffffffffffffff916000916165a0575b5016330361631357565b6165c2915060203d6020116165c8575b6165ba8183610769565b8101906164dc565b38616596565b503d6165b0565b67ffffffffffffffff7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894491169182600052600760205280616418604060002073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391616686565b8115616657570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b8054939290919060ff60a086901c161580156168c1575b6168ba576166bc6fffffffffffffffffffffffffffffffff86166151ce565b90600184019586546166f6615b24615b1b6166e96151ce856fffffffffffffffffffffffffffffffff1690565b9460801c63ffffffff1690565b80616826575b50508381106167db575082821061675c575061047b939450616721916151ce9161363b565b6fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b906167936105f99361678e61677f846167796151ce8c5460801c90565b9361363b565b6167888361360e565b90612e4c565b61664d565b7fd0c8d23a0000000000000000000000000000000000000000000000000000000060005260045260245273ffffffffffffffffffffffffffffffffffffffff16604452606490565b7f1a76572a00000000000000000000000000000000000000000000000000000000600052600452602483905273ffffffffffffffffffffffffffffffffffffffff1660445260646000fd5b828592939511616890576168406151ce6168479460801c90565b9185616256565b84547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff00000000000000000000000000000000161785559138806166fc565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050509050565b50811561669d565b9192901561694457508151156168dd575090565b3b156168e65790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156169575750805190602001fd5b611ed8906040519182917f08c379a00000000000000000000000000000000000000000000000000000000083526004830161088e56fea164736f6c634300081a000a2e902d38f15b233cbb63711add0fca4545334d3a169d60c0a616494d7eea9544",
}

var HybridLockReleaseUSDCTokenPoolABI = HybridLockReleaseUSDCTokenPoolMetaData.ABI

var HybridLockReleaseUSDCTokenPoolBin = HybridLockReleaseUSDCTokenPoolMetaData.Bin

func DeployHybridLockReleaseUSDCTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, tokenMessenger common.Address, tokenMessengerV2 common.Address, cctpMessageTransmitterProxy common.Address, token common.Address, allowlist []common.Address, rmnProxy common.Address, router common.Address, previousPool common.Address) (common.Address, *types.Transaction, *HybridLockReleaseUSDCTokenPool, error) {
	parsed, err := HybridLockReleaseUSDCTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HybridLockReleaseUSDCTokenPoolBin), backend, tokenMessenger, tokenMessengerV2, cctpMessageTransmitterProxy, token, allowlist, rmnProxy, router, previousPool)
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) ITokenMessengerCCTPV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "i_tokenMessengerCCTPV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) ITokenMessengerCCTPV2() (common.Address, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.ITokenMessengerCCTPV2(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) ITokenMessengerCCTPV2() (common.Address, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.ITokenMessengerCCTPV2(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCaller) IUsdcVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HybridLockReleaseUSDCTokenPool.contract.Call(opts, &out, "i_usdcVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) IUsdcVersion() (*big.Int, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.IUsdcVersion(&_HybridLockReleaseUSDCTokenPool.CallOpts)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolCallerSession) IUsdcVersion() (*big.Int, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.IUsdcVersion(&_HybridLockReleaseUSDCTokenPool.CallOpts)
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

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactor) UpdateCCTPVersion(opts *bind.TransactOpts, remoteChainSelectors []uint64, versions []uint8) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.contract.Transact(opts, "updateCCTPVersion", remoteChainSelectors, versions)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolSession) UpdateCCTPVersion(remoteChainSelectors []uint64, versions []uint8) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.UpdateCCTPVersion(&_HybridLockReleaseUSDCTokenPool.TransactOpts, remoteChainSelectors, versions)
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolTransactorSession) UpdateCCTPVersion(remoteChainSelectors []uint64, versions []uint8) (*types.Transaction, error) {
	return _HybridLockReleaseUSDCTokenPool.Contract.UpdateCCTPVersion(&_HybridLockReleaseUSDCTokenPool.TransactOpts, remoteChainSelectors, versions)
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

type HybridLockReleaseUSDCTokenPoolCCTPVersionSetIterator struct {
	Event *HybridLockReleaseUSDCTokenPoolCCTPVersionSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *HybridLockReleaseUSDCTokenPoolCCTPVersionSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridLockReleaseUSDCTokenPoolCCTPVersionSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(HybridLockReleaseUSDCTokenPoolCCTPVersionSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *HybridLockReleaseUSDCTokenPoolCCTPVersionSetIterator) Error() error {
	return it.fail
}

func (it *HybridLockReleaseUSDCTokenPoolCCTPVersionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type HybridLockReleaseUSDCTokenPoolCCTPVersionSet struct {
	RemoteChainSelector uint64
	Version             uint8
	Raw                 types.Log
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) FilterCCTPVersionSet(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolCCTPVersionSetIterator, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.FilterLogs(opts, "CCTPVersionSet")
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolCCTPVersionSetIterator{contract: _HybridLockReleaseUSDCTokenPool.contract, event: "CCTPVersionSet", logs: logs, sub: sub}, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) WatchCCTPVersionSet(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolCCTPVersionSet) (event.Subscription, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.WatchLogs(opts, "CCTPVersionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(HybridLockReleaseUSDCTokenPoolCCTPVersionSet)
				if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "CCTPVersionSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) ParseCCTPVersionSet(log types.Log) (*HybridLockReleaseUSDCTokenPoolCCTPVersionSet, error) {
	event := new(HybridLockReleaseUSDCTokenPoolCCTPVersionSet)
	if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "CCTPVersionSet", log); err != nil {
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

type HybridLockReleaseUSDCTokenPoolMintRecipientOverrideSetIterator struct {
	Event *HybridLockReleaseUSDCTokenPoolMintRecipientOverrideSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *HybridLockReleaseUSDCTokenPoolMintRecipientOverrideSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HybridLockReleaseUSDCTokenPoolMintRecipientOverrideSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(HybridLockReleaseUSDCTokenPoolMintRecipientOverrideSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *HybridLockReleaseUSDCTokenPoolMintRecipientOverrideSetIterator) Error() error {
	return it.fail
}

func (it *HybridLockReleaseUSDCTokenPoolMintRecipientOverrideSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type HybridLockReleaseUSDCTokenPoolMintRecipientOverrideSet struct {
	ChainSelector uint64
	MintRecipient [32]byte
	Raw           types.Log
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) FilterMintRecipientOverrideSet(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolMintRecipientOverrideSetIterator, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.FilterLogs(opts, "MintRecipientOverrideSet")
	if err != nil {
		return nil, err
	}
	return &HybridLockReleaseUSDCTokenPoolMintRecipientOverrideSetIterator{contract: _HybridLockReleaseUSDCTokenPool.contract, event: "MintRecipientOverrideSet", logs: logs, sub: sub}, nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) WatchMintRecipientOverrideSet(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolMintRecipientOverrideSet) (event.Subscription, error) {

	logs, sub, err := _HybridLockReleaseUSDCTokenPool.contract.WatchLogs(opts, "MintRecipientOverrideSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(HybridLockReleaseUSDCTokenPoolMintRecipientOverrideSet)
				if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "MintRecipientOverrideSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_HybridLockReleaseUSDCTokenPool *HybridLockReleaseUSDCTokenPoolFilterer) ParseMintRecipientOverrideSet(log types.Log) (*HybridLockReleaseUSDCTokenPoolMintRecipientOverrideSet, error) {
	event := new(HybridLockReleaseUSDCTokenPoolMintRecipientOverrideSet)
	if err := _HybridLockReleaseUSDCTokenPool.contract.UnpackLog(event, "MintRecipientOverrideSet", log); err != nil {
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
	case _HybridLockReleaseUSDCTokenPool.abi.Events["CCTPVersionSet"].ID:
		return _HybridLockReleaseUSDCTokenPool.ParseCCTPVersionSet(log)
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
	case _HybridLockReleaseUSDCTokenPool.abi.Events["MintRecipientOverrideSet"].ID:
		return _HybridLockReleaseUSDCTokenPool.ParseMintRecipientOverrideSet(log)
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

func (HybridLockReleaseUSDCTokenPoolCCTPVersionSet) Topic() common.Hash {
	return common.HexToHash("0x74cf5df65e6643e8523827033a1a33c9c370c63e2898443e7f578129ac616da6")
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

func (HybridLockReleaseUSDCTokenPoolMintRecipientOverrideSet) Topic() common.Hash {
	return common.HexToHash("0x4635bb8cae87f240e5697c20956067e3d85ca85cf51ae94584b71a1b98d9f5bb")
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

	ILocalDomainIdentifier(opts *bind.CallOpts) (uint32, error)

	IMessageTransmitterProxy(opts *bind.CallOpts) (common.Address, error)

	IPreviousPool(opts *bind.CallOpts) (common.Address, error)

	ITokenMessenger(opts *bind.CallOpts) (common.Address, error)

	ITokenMessengerCCTPV2(opts *bind.CallOpts) (common.Address, error)

	IUsdcVersion(opts *bind.CallOpts) (*big.Int, error)

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

	UpdateCCTPVersion(opts *bind.TransactOpts, remoteChainSelectors []uint64, versions []uint8) (*types.Transaction, error)

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

	FilterCCTPVersionSet(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolCCTPVersionSetIterator, error)

	WatchCCTPVersionSet(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolCCTPVersionSet) (event.Subscription, error)

	ParseCCTPVersionSet(log types.Log) (*HybridLockReleaseUSDCTokenPoolCCTPVersionSet, error)

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

	FilterMintRecipientOverrideSet(opts *bind.FilterOpts) (*HybridLockReleaseUSDCTokenPoolMintRecipientOverrideSetIterator, error)

	WatchMintRecipientOverrideSet(opts *bind.WatchOpts, sink chan<- *HybridLockReleaseUSDCTokenPoolMintRecipientOverrideSet) (event.Subscription, error)

	ParseMintRecipientOverrideSet(log types.Log) (*HybridLockReleaseUSDCTokenPoolMintRecipientOverrideSet, error)

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
