// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package usdc_token_pool

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

var USDCTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"internalType\":\"contractITokenMessenger\"},{\"name\":\"cctpMessageTransmitterProxy\",\"type\":\"address\",\"internalType\":\"contractCCTPMessageTransmitterProxy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"supportedUSDCVersion\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"structAuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDomain\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structUSDCTokenPool.Domain\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_localDomainIdentifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_messageTransmitterProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractCCTPMessageTransmitterProxy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_supportedUSDCVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_tokenMessenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITokenMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDomains\",\"inputs\":[{\"name\":\"domains\",\"type\":\"tuple[]\",\"internalType\":\"structUSDCTokenPool.DomainUpdate[]\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRouter\",\"inputs\":[{\"name\":\"newRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DomainsSet\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structUSDCTokenPool.DomainUpdate[]\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RouterUpdated\",\"inputs\":[{\"name\":\"oldRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestinationDomain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidDomain\",\"inputs\":[{\"name\":\"domain\",\"type\":\"tuple\",\"internalType\":\"structUSDCTokenPool.DomainUpdate\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"type\":\"error\",\"name\":\"InvalidMessageLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidNonce\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"got\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidPreviousPool\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourceDomain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenMessengerVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidTransmitterInProxy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnknownDomain\",\"inputs\":[{\"name\":\"domain\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"UnlockingUSDCFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610180806040523461065557616216803803809161001d8285610a1e565b8339810160e0828203126106555781516001600160a01b03811692838203610655576020810151936001600160a01b03851690818603610655576040830151956001600160a01b038716958688036106555760608501516001600160401b0381116106555785019080601f83011215610655578151916001600160401b03831161065a578260051b9060208201936100b86040519586610a1e565b845260208085019282010192831161065557602001905b828210610a06575050506100e560808601610a41565b946100fe60c06100f760a08401610a41565b9201610a55565b95602099604051996101108c8c610a1e565b60008b52600036813733156109f557600180546001600160a01b03191633179055801580156109e4575b80156109d3575b6109c2576004928c9260805260c0526040519283809263313ce56760e01b82525afa809160009161098b575b5090610967575b50600660a052600480546001600160a01b0319166001600160a01b03929092169190911790558051151560e0819052610839575b50604051946101b78887610a1e565b60008652600036813760408051979088016001600160401b0381118982101761065a576040528752858888015260005b865181101561024e576001906001600160a01b03610205828a610a66565b51168a61021182610ae7565b61021e575b5050016101e7565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a1388a610216565b5087955086519360005b85518110156102c9576001600160a01b036102738288610a66565b51169081156102b8577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef89836102aa600195610c03565b50604051908152a101610258565b6342bcdf7f60e11b60005260046000fd5b50869394508561010052841561082857604051632c12192160e01b81528481600481895afa9081156106c0576000916107f3575b5060405163054fd4d560e41b81526001600160a01b039190911691908581600481865afa9081156106c0576000916107be575b5063ffffffff8061010051169116908082036107a7575050604051639cdbb18160e01b815285816004818a5afa9081156106c057600091610772575b5063ffffffff80610100511691169080820361075b57505084600491604051928380926367e0ed8360e11b82525afa80156106c0578291600091610712575b506001600160a01b03160361070157600492849261012052610140526040519283809263234d8e3d60e21b82525afa9081156106c0576000916106cc575b506101605260805161012051604051636eb1769f60e11b81523060048201526001600160a01b03918216602482018190529492909116908381604481855afa9081156106c057600091610693575b50600019810180911161067d57604051908482019563095ea7b360e01b8752602483015260448201526044815261046f606482610a1e565b6000806040968751936104828986610a1e565b8785527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656488860152519082865af13d15610670573d906001600160401b03821161065a5786516104ef9490926104e1601f8201601f1916890185610a1e565b83523d60008885013e610cd0565b8051806105dc575b847f2e902d38f15b233cbb63711add0fca4545334d3a169d60c0a616494d7eea954485858351908152a1516154759081610da182396080518181816104f601528181610590015281816108b0015281816117d9015281816148810152614c5f015260a051816106cc015260c0518181816124a801528181613a7101526143a6015260e051818181610c71015281816125e90152614b890152610100518181816104870152613f1b015261012051818181610e910152611887015261014051818181610818015261163d015261016051818181610fee0152818161195c0152613f880152f35b8184918101031261065557820151801590811503610655576105ff5783806104f7565b50608491519062461bcd60e51b82526004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b916104ef92606091610cd0565b634e487b7160e01b600052601160045260246000fd5b90508381813d83116106b9575b6106aa8183610a1e565b81010312610655575185610437565b503d6106a0565b6040513d6000823e3d90fd5b90508181813d83116106fa575b6106e38183610a1e565b81010312610655576106f490610a55565b836103e9565b503d6106d9565b632a32133b60e11b60005260046000fd5b9091508581813d8311610754575b61072a8183610a1e565b810103126107505751906001600160a01b038216820361074d57508190876103ab565b80fd5b5080fd5b503d610720565b633785f8f160e01b60005260045260245260446000fd5b90508581813d83116107a0575b6107898183610a1e565b810103126106555761079a90610a55565b8761036c565b503d61077f565b63960693cd60e01b60005260045260245260446000fd5b90508581813d83116107ec575b6107d58183610a1e565b81010312610655576107e690610a55565b87610330565b503d6107cb565b90508481813d8311610821575b61080a8183610a1e565b810103126106555761081b90610a41565b866102fd565b503d610800565b6306b7c75960e31b60005260046000fd5b604051929692959194909361084e8988610a1e565b60008752600036813760e051156109565760005b87518110156108c9576001906001600160a01b03610880828b610a66565b51168b61088c82610c3c565b610899575b505001610862565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a1388b610891565b509193955091939560005b8651811015610948576001906001600160a01b036108f2828a610a66565b51168015610942578a61090482610bc4565b610912575b50505b016108d4565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a1388a610909565b5061090c565b5091939590929450386101a8565b6335f4a7b360e01b60005260046000fd5b60ff1660068114610174576332ad3e0760e11b600052600660045260245260446000fd5b8b81813d83116109bb575b6109a08183610a1e565b8101031261075057519060ff8216820361074d57503861016d565b503d610996565b630a64406560e11b60005260046000fd5b506001600160a01b03831615610141565b506001600160a01b0384161561013a565b639b15e16f60e01b60005260046000fd5b60208091610a1384610a41565b8152019101906100cf565b601f909101601f19168101906001600160401b0382119082101761065a57604052565b51906001600160a01b038216820361065557565b519063ffffffff8216820361065557565b8051821015610a7a5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b8054821015610a7a5760005260206000200190600090565b80548015610ad1576000190190610abf8282610a90565b8154906000199060031b1b1916905555565b634e487b7160e01b600052603160045260246000fd5b6000818152600b60205260409020548015610b9257600019810181811161067d57600a5460001981019190821161067d57808203610b41575b505050610b2d600a610aa8565b600052600b60205260006040812055600190565b610b7a610b52610b6393600a610a90565b90549060031b1c928392600a610a90565b819391549060031b91821b91600019901b19161790565b9055600052600b602052604060002055388080610b20565b5050600090565b8054906801000000000000000082101561065a5781610b63916001610bc094018155610a90565b9055565b80600052600360205260406000205415600014610bfd57610be6816002610b99565b600254906000526003602052604060002055600190565b50600090565b80600052600b60205260406000205415600014610bfd57610c2581600a610b99565b600a5490600052600b602052604060002055600190565b6000818152600360205260409020548015610b9257600019810181811161067d5760025460001981019190821161067d57818103610c96575b505050610c826002610aa8565b600052600360205260006040812055600190565b610cb8610ca7610b63936002610a90565b90549060031b1c9283926002610a90565b90556000526003602052604060002055388080610c75565b91929015610d325750815115610ce4575090565b3b15610ced5790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b825190915015610d455750805190602001fd5b6040519062461bcd60e51b8252602060048301528181519182602483015260005b838110610d885750508160006044809484010152601f80199101168101030190fd5b60208282018101516044878401015285935001610d6656fe6080604052600436101561001257600080fd5b60003560e01c806301ffc9a714610277578063181f5a7714610272578063212a052e1461026d57806321df0da714610268578063240028e8146102635780632451a6271461025e57806324f65ee71461025957806339077537146102545780634c5ef0ed1461024f57806354c8a4f31461024a5780636155cda01461024557806362ddd3c4146102405780636b716b0d1461023b5780636d3d1a581461023657806379ba5097146102315780637d54534e1461022c5780638926f54f146102275780638da5cb5b1461022257806391a2749a1461021d578063962d40201461021857806398db9643146102135780639a4575b91461020e578063a42a7b8b14610209578063a7cd63b714610204578063acfecf91146101ff578063af58d59f146101fa578063b0f479a1146101f5578063b7946580146101f0578063c0d78655146101eb578063c4bffe2b146101e6578063c75eea9c146101e1578063c781d0e3146101dc578063cf7401f3146101d7578063dc0bd971146101d2578063dfadfa35146101cd578063e0351e13146101c8578063e8a1da17146101c35763f2fde38b146101be57600080fd5b612a2d565b61260e565b6125b3565b6124cc565b61245d565b612341565b6121cc565b612163565b6120b2565b611f72565b611f1d565b611eba565b611df0565b611cc4565b611c4c565b611b2c565b6116b4565b6115f2565b611435565b611334565b611259565b6111fc565b61114d565b611064565b611012565b610fb3565b610f30565b610e46565b610c3f565b610b37565b6106f0565b610694565b610610565b610538565b6104ab565b61044c565b6103cb565b346103675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610367576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361036757807faff2afbf000000000000000000000000000000000000000000000000000000006020921490811561033d575b8115610313575b506040519015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438610308565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150610301565b600080fd5b919082519283825260005b8481106103b65750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201610377565b346103675760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261036757610448604080519061040c8183610a2b565b601782527f55534443546f6b656e506f6f6c20312e362e332d64657600000000000000000060208301525191829160208352602083019061036c565b0390f35b346103675760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261036757602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346103675760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261036757602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b73ffffffffffffffffffffffffffffffffffffffff81160361036757565b346103675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103675760206105b66004356105788161051a565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691161490565b6040519015158152f35b602060408183019282815284518094520192019060005b8181106105e45750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016105d7565b346103675760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261036757604051600a548082526020820190600a60005260206000209060005b81811061067e576104488561067281870382610a2b565b604051918291826105c0565b825484526020909301926001928301920161065b565b346103675760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261036757602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346103675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103675760043567ffffffffffffffff811161036757806004016101007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc83360301126103675761076a612b21565b506107fd602060648401359261078084826139a4565b6107c26107bb6107b66107af6107a461079c60e48b0187612b45565b810190612b96565b9460c48a0190612b45565b3691610ae2565b613d70565b8251613f08565b8181519101519060405193849283927f57ecfd2800000000000000000000000000000000000000000000000000000000845260048401612c24565b0381600073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af190811561096357600091610934575b501561090a57817ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc067ffffffffffffffff610895604461088e60246104489801612c55565b9401612c5f565b6040805173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000081168252336020830152929092169082015260608101859052921691608090a26108f7610a6c565b8190526040519081529081906020820190565b7fbf969f220000000000000000000000000000000000000000000000000000000060005260046000fd5b610956915060203d60201161095c575b61094e8183610a2b565b810190612c0f565b38610849565b503d610944565b612c49565b67ffffffffffffffff81160361036757565b359061098582610968565b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff8211176109d257604052565b610987565b6060810190811067ffffffffffffffff8211176109d257604052565b6080810190811067ffffffffffffffff8211176109d257604052565b60a0810190811067ffffffffffffffff8211176109d257604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176109d257604052565b60405190610985602083610a2b565b60405190610985604083610a2b565b6040519061098560a083610a2b565b60405190610985608083610a2b565b67ffffffffffffffff81116109d257601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b929192610aee82610aa8565b91610afc6040519384610a2b565b829481845281830111610367578281602093846000960137010152565b9080601f8301121561036757816020610b3493359101610ae2565b90565b346103675760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261036757600435610b7281610968565b60243567ffffffffffffffff811161036757602091610b986105b6923690600401610b19565b90612c69565b9181601f840112156103675782359167ffffffffffffffff8311610367576020808501948460051b01011161036757565b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126103675760043567ffffffffffffffff81116103675781610c1891600401610b9e565b929092916024359067ffffffffffffffff821161036757610c3b91600401610b9e565b9091565b3461036757610c67610c6f610c5336610bcf565b9491610c609391936140fb565b36916112c3565b9236916112c3565b7f000000000000000000000000000000000000000000000000000000000000000015610e1c5760005b8251811015610d5d5780610ccb610cb160019386612f52565b5173ffffffffffffffffffffffffffffffffffffffff1690565b610d07610d0273ffffffffffffffffffffffffffffffffffffffff83165b73ffffffffffffffffffffffffffffffffffffffff1690565b614d32565b610d13575b5001610c98565b60405173ffffffffffffffffffffffffffffffffffffffff9190911681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756690602090a138610d0c565b5060005b8151811015610e1a5780610d7a610cb160019385612f52565b73ffffffffffffffffffffffffffffffffffffffff811615610e1457610dbd610db873ffffffffffffffffffffffffffffffffffffffff8316610ce9565b615037565b610dca575b505b01610d61565b60405173ffffffffffffffffffffffffffffffffffffffff9190911681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d890602090a183610dc2565b50610dc4565b005b7f35f4a7b30000000000000000000000000000000000000000000000000000000060005260046000fd5b346103675760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261036757602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82011261036757600435610eeb81610968565b9160243567ffffffffffffffff811161036757826023820112156103675780600401359267ffffffffffffffff84116103675760248483010111610367576024019190565b3461036757610f3e36610eb5565b610f499291926140fb565b67ffffffffffffffff8216610f6b816000526006602052604060002054151590565b15610f865750610e1a92610f80913691610ae2565b90614167565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346103675760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261036757602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346103675760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261036757602073ffffffffffffffffffffffffffffffffffffffff60095416604051908152f35b346103675760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103675760005473ffffffffffffffffffffffffffffffffffffffff81163303611123577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346103675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610367577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174602073ffffffffffffffffffffffffffffffffffffffff6004356111c08161051a565b6111c86140fb565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000006009541617600955604051908152a1005b346103675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103675760206105b667ffffffffffffffff60043561124581610968565b166000526006602052604060002054151590565b346103675760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261036757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b67ffffffffffffffff81116109d25760051b60200190565b9291906112cf816112ab565b936112dd6040519586610a2b565b602085838152019160051b810192831161036757905b8282106112ff57505050565b60208091833561130e8161051a565b8152019101906112f3565b9080601f8301121561036757816020610b34933591016112c3565b346103675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103675760043567ffffffffffffffff81116103675760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8236030112610367576040516113ad816109b6565b816004013567ffffffffffffffff8111610367576113d19060043691850101611319565b8152602482013567ffffffffffffffff811161036757610e1a9260046113fa9236920101611319565b6020820152612ca6565b9181601f840112156103675782359167ffffffffffffffff8311610367576020808501946060850201011161036757565b346103675760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103675760043567ffffffffffffffff811161036757611484903690600401610b9e565b9060243567ffffffffffffffff8111610367576114a5903690600401611404565b9060443567ffffffffffffffff8111610367576114c6903690600401611404565b6114e8610ce960095473ffffffffffffffffffffffffffffffffffffffff1690565b331415806115c7575b6115955783861480159061158b575b6115615760005b86811061151057005b8061155b6115296115246001948b8b612e64565b612c55565b611534838989612e79565b61155561154d61154586898b612e79565b9236906122f8565b9136906122f8565b9161422c565b01611507565b7f568efce20000000000000000000000000000000000000000000000000000000060005260046000fd5b5080861415611500565b7f8e4a23d6000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b6000fd5b506115ea610ce960015473ffffffffffffffffffffffffffffffffffffffff1690565b3314156114f1565b346103675760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261036757602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b90610b34916020815260206116818351604083850152606084019061036c565b9201519060407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08285030191015261036c565b346103675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103675760043567ffffffffffffffff811161036757806004019060a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126103675761172e612e89565b506117388261435c565b602481019061176b61176661174c84612c55565b67ffffffffffffffff16600052600c602052604060002090565b612ea2565b9061178061177c6060840151151590565b1590565b611a6c5760206117908580612b45565b905003611a285760208281015161186c95606491908115611a0b5750925b0135916117c2604085015163ffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169687955160405198899485947ff856ddb60000000000000000000000000000000000000000000000000000000086528860048701919360809363ffffffff73ffffffffffffffffffffffffffffffffffffffff9398979660a0860199865216602085015260408401521660608201520152565b0381600073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af193841561096357610448946119919361198c93611986936000936119ae575b5067ffffffffffffffff7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10916119376119008a612c55565b6040805173ffffffffffffffffffffffffffffffffffffffff90971687523360208801528601929092529116929081906060820190565b0390a2611955611945610a7b565b67ffffffffffffffff9092168252565b63ffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015261442d565b92612c55565b61311f565b9061199a610a7b565b918252602082015260405191829182611661565b7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae109193506119fc67ffffffffffffffff9160203d602011611a04575b6119f48183610a2b565b810190612f3d565b9391506118c8565b503d6119ea565b611a22915080611a1a91612b45565b810190612f2e565b926117ae565b611a328480612b45565b90611a686040519283927fa3c8cf0900000000000000000000000000000000000000000000000000000000845260048401612f1d565b0390fd5b6115c3611a7884612c55565b7fd201c48a0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b602081016020825282518091526040820191602060408360051b8301019401926000915b838310611ae157505050505090565b9091929394602080611b1d837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08660019603018752895161036c565b97019301930191939290611ad2565b346103675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103675767ffffffffffffffff600435611b7081610968565b166000526007602052611b896005604060002001614745565b8051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611bcf611bb9846112ab565b93611bc76040519586610a2b565b8085526112ab565b0160005b818110611c3b57505060005b8151811015611c2d5780611c11611c0c611bfb60019486612f52565b516000526008602052604060002090565b612fb9565b611c1b8286612f52565b52611c268185612f52565b5001611bdf565b604051806104488582611aae565b806060602080938701015201611bd3565b346103675760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610367576040516002548082526020820190600260005260206000209060005b818110611cae576104488561067281870382610a2b565b8254845260209093019260019283019201611c97565b3461036757611cd236610eb5565b611cdd9291926140fb565b67ffffffffffffffff821691611d0361177c846000526006602052604060002054151590565b611db957611d4661177c6005611d2d8467ffffffffffffffff166000526007602052604060002090565b01611d39368689610ae2565b6020815191012090614edd565b611d8257507f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d769192611d7d60405192839283612f1d565b0390a2005b611a6884926040519384937f74f23c7c00000000000000000000000000000000000000000000000000000000855260048501613079565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b346103675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103675767ffffffffffffffff600435611e3481610968565b611e3c61309a565b50166000526007602052610448611e61611e5c60026040600020016130c5565b6144dc565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b346103675760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261036757602073ffffffffffffffffffffffffffffffffffffffff60045416604051908152f35b906020610b3492818152019061036c565b346103675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261036757610448611f5e60043561198c81610968565b60405191829160208352602083019061036c565b346103675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103675773ffffffffffffffffffffffffffffffffffffffff600435611fc28161051a565b611fca6140fb565b1680156120445760407f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f16849160045490807fffffffffffffffffffffffff000000000000000000000000000000000000000083161760045573ffffffffffffffffffffffffffffffffffffffff8351921682526020820152a1005b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b602060408183019282815284518094520192019060005b8181106120925750505090565b825167ffffffffffffffff16845260209384019390920191600101612085565b346103675760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610367576120e96146fa565b8051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612119611bb9846112ab565b0136602084013760005b8151811015612155578067ffffffffffffffff61214260019385612f52565b511661214e8286612f52565b5201612123565b60405180610448858261206e565b346103675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103675767ffffffffffffffff6004356121a781610968565b6121af61309a565b50166000526007602052610448611e61611e5c60406000206130c5565b346103675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103675760043567ffffffffffffffff8111610367573660238201121561036757806004013567ffffffffffffffff81116103675736602460a0830284010111610367576024610e1a9201613141565b8015150361036757565b35906fffffffffffffffffffffffffffffffff8216820361036757565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c606091011261036757604051906122a5826109d7565b816084356122b281612247565b815260a4356fffffffffffffffffffffffffffffffff8116810361036757602082015260c435906fffffffffffffffffffffffffffffffff821682036103675760400152565b919082606091031261036757604051612310816109d7565b604061233c818395803561232381612247565b855261233160208201612251565b602086015201612251565b910152565b346103675760e07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103675760043561237c81610968565b60607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc360112610367576040516123b2816109d7565b6024356123be81612247565b81526044356fffffffffffffffffffffffffffffffff811681036103675760208201526064356fffffffffffffffffffffffffffffffff8116810361036757604082015261240b3661226e565b9073ffffffffffffffffffffffffffffffffffffffff600954163314158061243b575b61159557610e1a9261422c565b5073ffffffffffffffffffffffffffffffffffffffff6001541633141561242e565b346103675760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261036757602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346103675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103675767ffffffffffffffff60043561251081610968565b60006060604051612520816109f3565b828152826020820152826040820152015216600052600c602052610448604060002060ff600260405192612553846109f3565b8054845260018101546020850152015463ffffffff8116604084015260201c1615156060820152604051918291829190916060806080830194805184526020810151602085015263ffffffff604082015116604085015201511515910152565b346103675760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103675760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b346103675761261c36610bcf565b9190926126276140fb565b6000915b8083106128d95750505060009163ffffffff4216925b82811061264a57005b61265d61265882858561358e565b61364d565b906060820161266c81516145b9565b608083019361267b85516145b9565b604084019081515115612044576126b561177c6126b06126a3885167ffffffffffffffff1690565b67ffffffffffffffff1690565b615072565b61288e576127ee6126ee6126d4879a999a5167ffffffffffffffff1690565b67ffffffffffffffff166000526007602052604060002090565b6127b1896127ab875161279261271760408301516fffffffffffffffffffffffffffffffff1690565b9161277961274261273b60208401516fffffffffffffffffffffffffffffffff1690565b9251151590565b61277061274d610a8a565b6fffffffffffffffffffffffffffffffff851681529763ffffffff166020890152565b15156040870152565b6fffffffffffffffffffffffffffffffff166060850152565b6fffffffffffffffffffffffffffffffff166080830152565b826136dc565b6127e3896127da8a5161279261271760408301516fffffffffffffffffffffffffffffffff1690565b600283016136dc565b6004845191016137e8565b602085019660005b88518051821015612831579061282b6001926128248361281e8c5167ffffffffffffffff1690565b92612f52565b5190614167565b016127f6565b505097965094906128857f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c293926128726001975167ffffffffffffffff1690565b925193519051906040519485948561390f565b0390a101612641565b6115c36128a3865167ffffffffffffffff1690565b7f1d5ad3c50000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b9091926128ea611524858486612e64565b9461290161177c67ffffffffffffffff8816614e16565b6129f55761292e60056129288867ffffffffffffffff166000526007602052604060002090565b01614745565b9360005b855181101561297a5760019061297360056129618b67ffffffffffffffff166000526007602052604060002090565b0161296c838a612f52565b5190614edd565b5001612932565b509350937f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166129e7600193976129cc6129c78267ffffffffffffffff166000526007602052604060002090565b6134df565b60405167ffffffffffffffff90911681529081906020820190565b0390a101919093929361262b565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff861660045260246000fd5b346103675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103675773ffffffffffffffffffffffffffffffffffffffff600435612a7d8161051a565b612a856140fb565b16338114612af757807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b604051906020820182811067ffffffffffffffff8211176109d25760405260008252565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610367570180359067ffffffffffffffff82116103675760200191813603831361036757565b6020818303126103675780359067ffffffffffffffff821161036757016040818303126103675760405191612bca836109b6565b813567ffffffffffffffff81116103675781612be7918401610b19565b8352602082013567ffffffffffffffff811161036757612c079201610b19565b602082015290565b908160209103126103675751610b3481612247565b9091612c3b610b349360408452604084019061036c565b91602081840391015261036c565b6040513d6000823e3d90fd5b35610b3481610968565b35610b348161051a565b9067ffffffffffffffff610b3492166000526007602052600560406000200190602081519101209060019160005201602052604060002054151590565b612cae6140fb565b60208101519160005b8351811015612d4d5780612cd0610cb160019387612f52565b612cf7612cf273ffffffffffffffffffffffffffffffffffffffff8316610ce9565b61538f565b612d03575b5001612cb7565b60405173ffffffffffffffffffffffffffffffffffffffff9190911681527fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758090602090a138612cfc565b5091505160005b8151811015612e3157612d6a610cb18284612f52565b9073ffffffffffffffffffffffffffffffffffffffff821615612e07577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef612dfe83612dd6612dd1610ce960019773ffffffffffffffffffffffffffffffffffffffff1690565b6150a7565b5060405173ffffffffffffffffffffffffffffffffffffffff90911681529081906020820190565b0390a101612d54565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9190811015612e745760051b0190565b612e35565b9190811015612e74576060020190565b60405190612e96826109b6565b60606020838281520152565b90604051612eaf816109f3565b606060ff600283958054855260018101546020860152015463ffffffff8116604085015260201c161515910152565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b916020610b34938181520191612ede565b90816020910312610367573590565b908160209103126103675751610b3481610968565b8051821015612e745760209160051b010190565b90600182811c92168015612faf575b6020831014612f8057565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691612f75565b9060405191826000825492612fcd84612f66565b80845293600181169081156130395750600114612ff2575b5061098592500383610a2b565b90506000929192526020600020906000915b81831061301d5750509060206109859282010138612fe5565b6020919350806001915483858901015201910190918492613004565b602093506109859592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138612fe5565b60409067ffffffffffffffff610b3495931681528160208201520191612ede565b604051906130a782610a0f565b60006080838281528260208201528260408201528260608201520152565b906040516130d281610a0f565b60806fffffffffffffffffffffffffffffffff6001839560ff8154848116875263ffffffff81871c16602088015260a01c1615156040860152015481808216166060850152821c16910152565b67ffffffffffffffff166000526007602052610b346004604060002001612fb9565b6131496140fb565b60005b82811061318b5750907fe6d14ea297366c7bc1265d289d924bfd8b9afb148eb972b481f70da41c842cf59161318660405192839283613388565b0390a1565b61319e613199828585613309565b61332a565b80511580156132e3575b61327657906132708261321661174c6060600196519361320760208201516131ff6131da604085015163ffffffff1690565b60808501511515926131ea610a99565b998a5260208a015263ffffffff166040890152565b151586840152565b015167ffffffffffffffff1690565b6002908251815560208301516001820155019063ffffffff6040820151167fffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000064ff0000000060608554940151151560201b16921617179055565b0161314c565b604080517fa606c63500000000000000000000000000000000000000000000000000000000815282516004820152602083015160248201529082015163ffffffff166044820152606082015167ffffffffffffffff1660648201526080909101511515608482015260a490fd5b5067ffffffffffffffff613302606083015167ffffffffffffffff1690565b16156131a8565b9190811015612e745760a0020190565b359063ffffffff8216820361036757565b60a0813603126103675760806040519161334383610a0f565b803583526020810135602084015261335d60408201613319565b6040840152606081013561337081610968565b6060840152013561338081612247565b608082015290565b602080825281018390526040019160005b8181106133a65750505090565b90919260a080600192863581526020870135602082015263ffffffff6133ce60408901613319565b16604082015267ffffffffffffffff60608801356133eb81610968565b16606082015260808701356133ff81612247565b15156080820152019401929101613399565b91613449918354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055565b818110613458575050565b6000815560010161344d565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818102929181159184041417156134a657565b613464565b80549060008155816134bb575050565b6000526020600020908101905b8181106134d3575050565b600081556001016134c8565b600561098591600081556000600182015560006002820155600060038201556004810161350c8154612f66565b908161351b575b5050016134ab565b81601f600093116001146135335750555b3880613513565b8183526020832061354e91601f01861c81019060010161344d565b808252602082209081548360011b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8560031b1c19161790555561352c565b9190811015612e745760051b810135907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee181360301821215610367570190565b9080601f830112156103675781356135e5816112ab565b926135f36040519485610a2b565b81845260208085019260051b820101918383116103675760208201905b83821061361f57505050505090565b813567ffffffffffffffff81116103675760209161364287848094880101610b19565b815201910190613610565b61012081360312610367576040519061366582610a0f565b61366e8161097a565b8252602081013567ffffffffffffffff81116103675761369190369083016135ce565b602083015260408101359067ffffffffffffffff8211610367576136bb6133809236908301610b19565b60408401526136cd36606083016122f8565b606084015260c03691016122f8565b8151815460208401516040850151608091821b73ffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffff000000000000000000000000000000000000000000909416939093179290921791151560a01b74ff000000000000000000000000000000000000000016919091178355606084015193810151901b7fffffffffffffffffffffffffffffffff0000000000000000000000000000000016921691909117600190910155565b9190601f81116137b257505050565b610985926000526020600020906020601f840160051c830193106137de575b601f0160051c019061344d565b90915081906137d1565b919091825167ffffffffffffffff81116109d2576138108161380a8454612f66565b846137a3565b6020601f821160011461386a57819061344993949560009261385f575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b01519050388061382d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169061389d84600052602060002090565b9160005b8181106138f7575095836001959697106138c0575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880806138b6565b9192602060018192868b0151815501940192016138a1565b61397361393e6109859597969467ffffffffffffffff60a095168452610100602085015261010084019061036c565b9660408301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b01906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b608081016139b761177c61057883612c5f565b613b7b57506020810190613a5860206139fd6139d56126a386612c55565b60801b7fffffffffffffffffffffffffffffffff000000000000000000000000000000001690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116600482015291829081906024820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561096357600091613b5c575b50613b3257613ab8613ab383612c55565b614790565b613ac182612c55565b90613ada61177c60a0830193610b986107af8686612b45565b613af257505090613aed61098592612c55565b614828565b613afc9250612b45565b90611a686040519283927f24eb47e500000000000000000000000000000000000000000000000000000000845260048401612f1d565b7f53ad11d80000000000000000000000000000000000000000000000000000000060005260046000fd5b613b75915060203d60201161095c5761094e8183610a2b565b38613aa2565b613b876115c391612c5f565b7f961c9a4f0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b60405190613bd6826109b6565b60006020838281520152565b60408051909190613bf38382610a2b565b60048152917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001366020840137565b60408051909190613c338382610a2b565b60088152917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001366020840137565b90600882018092116134a657565b90600482018092116134a657565b919082018092116134a657565b908151811015612e74570160200190565b90602082519201517fffffffff0000000000000000000000000000000000000000000000000000000081169260048110613cd4575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b90602082519201517fffffffffffffffff00000000000000000000000000000000000000000000000081169260088110613d3e575050565b7fffffffffffffffff000000000000000000000000000000000000000000000000929350829060080360031b1b161690565b613d78613bc9565b90613d81613be2565b60005b60048110613ee35750613d99613d9f91613c9c565b60e01c90565b63ffffffff8116613eb057506000906004613db8613c22565b60005b60088110613e845750613dd9613dd3613de792613d06565b60c01c90565b67ffffffffffffffff168552565b613df16004613c62565b92613dfa613be2565b9260005b838110613e32575050505090613e29613e1c613d99613e2e94613c9c565b63ffffffff166020850152565b613c70565b5090565b80613e71613e4b613e456001948a613c7e565b85613c8b565b517fff000000000000000000000000000000000000000000000000000000000000001690565b841a613e7d8288613c8b565b5301613dfe565b80613e9d613e4b613e9760019487613c7e565b87613c8b565b861a613ea98285613c8b565b5301613dbb565b7f70fdcbf80000000000000000000000000000000000000000000000000000000060005263ffffffff1660045260246000fd5b6001906000613ef5613e4b8387613c8b565b901a613f018285613c8b565b5301613d84565b908151607481106140ce575060048201517f000000000000000000000000000000000000000000000000000000000000000063ffffffff811663ffffffff8316036140955750506008820151916014600c82015191015192613f71602084015163ffffffff1690565b63ffffffff811663ffffffff83160361405c5750507f000000000000000000000000000000000000000000000000000000000000000063ffffffff811663ffffffff8316036140235750505167ffffffffffffffff1667ffffffffffffffff811667ffffffffffffffff831603613fe6575050565b7ff917ffea0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff9081166004521660245260446000fd5b7f77e480260000000000000000000000000000000000000000000000000000000060005263ffffffff9081166004521660245260446000fd5b7fe366a1170000000000000000000000000000000000000000000000000000000060005263ffffffff9081166004521660245260446000fd5b7f960693cd0000000000000000000000000000000000000000000000000000000060005263ffffffff9081166004521660245260446000fd5b7f758b22cc0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff60015416330361411c57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b60409067ffffffffffffffff610b349493168152816020820152019061036c565b90805115612044578051602082012067ffffffffffffffff83169283600052600760205261419c8260056040600020016150dc565b156141f55750816141e47f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea936141df6141f0946000526008602052604060002090565b6137e8565b60405191829182611f0c565b0390a2565b9050611a686040519283927f393b8ad200000000000000000000000000000000000000000000000000000000845260048401614146565b67ffffffffffffffff16600081815260066020526040902054909291901561432e579161432b60e0926142f7856142837f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b976145b9565b84600052600760205261429a8160406000206148d9565b6142a3836145b9565b8460005260076020526142bd8360026040600020016148d9565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6080810161436f61177c61057883612c5f565b613b7b5750602081019061438d60206139fd6139d56126a386612c55565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156109635760009161440e575b50613b32576060614405610985936143f96143f460408601612c5f565b614b87565b611524613ab382612c55565b91013590614c09565b614427915060203d60201161095c5761094e8183610a2b565b386143d7565b7fffffffff00000000000000000000000000000000000000000000000000000000602082519201517fffffffffffffffff000000000000000000000000000000000000000000000000604051936000602086015260c01b16602484015260e01b16602c82015260108152610b34603082610a2b565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116134a657565b919082039182116134a657565b6144e461309a565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff82511690602083019163ffffffff83511642034281116134a657614548906fffffffffffffffffffffffffffffffff60808701511690613493565b81018091116134a65761456e6fffffffffffffffffffffffffffffffff92918392615456565b161682524263ffffffff16905290565b6109859092919260608101936fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b80511561465d5760408101516fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff61461d61460860208501516fffffffffffffffffffffffffffffffff1690565b6fffffffffffffffffffffffffffffffff1690565b9116116146275750565b611a68906040519182917f8020d1240000000000000000000000000000000000000000000000000000000083526004830161457e565b6fffffffffffffffffffffffffffffffff61468b60408301516fffffffffffffffffffffffffffffffff1690565b16158015906146d2575b61469c5750565b611a68906040519182917fd68af9cc0000000000000000000000000000000000000000000000000000000083526004830161457e565b506146f361460860208301516fffffffffffffffffffffffffffffffff1690565b1515614695565b604051906005548083528260208101600560005260206000209260005b81811061472c57505061098592500383610a2b565b8454835260019485019487945060209093019201614717565b906040519182815491828252602082019060005260206000209260005b81811061477757505061098592500383610a2b565b8454835260019485019487945060209093019201614762565b67ffffffffffffffff166147b1816000526006602052604060002054151590565b156147fb575033600052600b602052604060002054156147cd57565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fa9902c7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b67ffffffffffffffff7f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c911691826000526007602052806148a9600260406000200173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692839161514c565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252602082019290925290819081016141f0565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1991614ab861318692805461492a61492461491b8363ffffffff9060801c1690565b63ffffffff1690565b426144cf565b9081614ac4575b5050614a72600161495560208601516fffffffffffffffffffffffffffffffff1690565b926149e06149a36146086fffffffffffffffffffffffffffffffff61498a85546fffffffffffffffffffffffffffffffff1690565b166fffffffffffffffffffffffffffffffff8816615456565b82906fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b614a336149ed8751151590565b82547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178255565b019182906fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b604083015181546fffffffffffffffffffffffffffffffff1660809190911b7fffffffffffffffffffffffffffffffff0000000000000000000000000000000016179055565b6040519182918261457e565b6146086149a3916fffffffffffffffffffffffffffffffff614b38614b3f9582614b3160018a01549282614b2a614b23614b0d876fffffffffffffffffffffffffffffffff1690565b996fffffffffffffffffffffffffffffffff1690565b9560801c90565b1690613493565b9116613c7e565b9116615456565b80547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff00000000000000000000000000000000161781553880614931565b7f0000000000000000000000000000000000000000000000000000000000000000614baf5750565b73ffffffffffffffffffffffffffffffffffffffff1680600052600360205260406000205415614bdc5750565b7fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b67ffffffffffffffff7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944911691826000526007602052806148a9604060002073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692839161514c565b8054821015612e745760005260206000200190600090565b80548015614d03577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190614cd48282614c87565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600081815260036020526040902054908115614e0f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201908282116134a657600254927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84019384116134a6578383600095614dce9503614dd4575b505050614dbd6002614c9f565b600390600052602052604060002090565b55600190565b614dbd614e0091614df6614dec614e06956002614c87565b90549060031b1c90565b9283916002614c87565b90613411565b55388080614db0565b5050600090565b600081815260066020526040902054908115614e0f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201908282116134a657600554927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84019384116134a6578383600095614dce9503614eb2575b505050614ea16005614c9f565b600690600052602052604060002090565b614ea1614e0091614eca614dec614ed4956005614c87565b9283916005614c87565b55388080614e94565b6001810191806000528260205260406000205492831515600014614fb5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84018481116134a6578354937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff85019485116134a6576000958583614dce97614f6d9503614f7c575b505050614c9f565b90600052602052604060002090565b614f9c614e0091614f93614dec614fac9588614c87565b92839187614c87565b8590600052602052604060002090565b55388080614f65565b50505050600090565b92614fc99192613493565b81018091116134a657610b3491615456565b805490680100000000000000008210156109d2578161500291600161344994018155614c87565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b60008181526003602052604090205461506c57615055816002614fdb565b600254906000526003602052604060002055600190565b50600090565b60008181526006602052604090205461506c57615090816005614fdb565b600554906000526006602052604060002055600190565b6000818152600b602052604090205461506c576150c581600a614fdb565b600a5490600052600b602052604060002055600190565b6000828152600182016020526040902054614e0f57806150fe83600193614fdb565b80549260005201602052604060002055600190565b811561511d570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b8054939290919060ff60a086901c16158015615387575b615380576151826fffffffffffffffffffffffffffffffff8616614608565b90600184019586546151bc61492461491b6151af614608856fffffffffffffffffffffffffffffffff1690565b9460801c63ffffffff1690565b806152ec575b50508381106152a1575082821061522257506109859394506151e791614608916144cf565b6fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b906152596115c3936152546152458461523f6146088c5460801c90565b936144cf565b61524e836144a2565b90613c7e565b615113565b7fd0c8d23a0000000000000000000000000000000000000000000000000000000060005260045260245273ffffffffffffffffffffffffffffffffffffffff16604452606490565b7f1a76572a00000000000000000000000000000000000000000000000000000000600052600452602483905273ffffffffffffffffffffffffffffffffffffffff1660445260646000fd5b8285929395116153565761530661460861530d9460801c90565b9185614fbe565b84547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff00000000000000000000000000000000161785559138806151c2565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050509050565b508115615163565b6000818152600b6020526040902054908115614e0f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201908282116134a657600a54927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84019384116134a6578383614dce946000960361542b575b50505061541a600a614c9f565b600b90600052602052604060002090565b61541a614e0091615443614dec61544d95600a614c87565b928391600a614c87565b5538808061540d565b9080821015615463575090565b90509056fea164736f6c634300081a000a",
}

var USDCTokenPoolABI = USDCTokenPoolMetaData.ABI

var USDCTokenPoolBin = USDCTokenPoolMetaData.Bin

func DeployUSDCTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, tokenMessenger common.Address, cctpMessageTransmitterProxy common.Address, token common.Address, allowlist []common.Address, rmnProxy common.Address, router common.Address, supportedUSDCVersion uint32) (common.Address, *types.Transaction, *USDCTokenPool, error) {
	parsed, err := USDCTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(USDCTokenPoolBin), backend, tokenMessenger, cctpMessageTransmitterProxy, token, allowlist, rmnProxy, router, supportedUSDCVersion)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &USDCTokenPool{address: address, abi: *parsed, USDCTokenPoolCaller: USDCTokenPoolCaller{contract: contract}, USDCTokenPoolTransactor: USDCTokenPoolTransactor{contract: contract}, USDCTokenPoolFilterer: USDCTokenPoolFilterer{contract: contract}}, nil
}

type USDCTokenPool struct {
	address common.Address
	abi     abi.ABI
	USDCTokenPoolCaller
	USDCTokenPoolTransactor
	USDCTokenPoolFilterer
}

type USDCTokenPoolCaller struct {
	contract *bind.BoundContract
}

type USDCTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type USDCTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type USDCTokenPoolSession struct {
	Contract     *USDCTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type USDCTokenPoolCallerSession struct {
	Contract *USDCTokenPoolCaller
	CallOpts bind.CallOpts
}

type USDCTokenPoolTransactorSession struct {
	Contract     *USDCTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type USDCTokenPoolRaw struct {
	Contract *USDCTokenPool
}

type USDCTokenPoolCallerRaw struct {
	Contract *USDCTokenPoolCaller
}

type USDCTokenPoolTransactorRaw struct {
	Contract *USDCTokenPoolTransactor
}

func NewUSDCTokenPool(address common.Address, backend bind.ContractBackend) (*USDCTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(USDCTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindUSDCTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPool{address: address, abi: abi, USDCTokenPoolCaller: USDCTokenPoolCaller{contract: contract}, USDCTokenPoolTransactor: USDCTokenPoolTransactor{contract: contract}, USDCTokenPoolFilterer: USDCTokenPoolFilterer{contract: contract}}, nil
}

func NewUSDCTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*USDCTokenPoolCaller, error) {
	contract, err := bindUSDCTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCaller{contract: contract}, nil
}

func NewUSDCTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*USDCTokenPoolTransactor, error) {
	contract, err := bindUSDCTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolTransactor{contract: contract}, nil
}

func NewUSDCTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*USDCTokenPoolFilterer, error) {
	contract, err := bindUSDCTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolFilterer{contract: contract}, nil
}

func bindUSDCTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := USDCTokenPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_USDCTokenPool *USDCTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDCTokenPool.Contract.USDCTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_USDCTokenPool *USDCTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.USDCTokenPoolTransactor.contract.Transfer(opts)
}

func (_USDCTokenPool *USDCTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.USDCTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_USDCTokenPool *USDCTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDCTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_USDCTokenPool *USDCTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.contract.Transfer(opts)
}

func (_USDCTokenPool *USDCTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_USDCTokenPool *USDCTokenPoolCaller) GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "getAllAuthorizedCallers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _USDCTokenPool.Contract.GetAllAuthorizedCallers(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _USDCTokenPool.Contract.GetAllAuthorizedCallers(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) GetAllowList() ([]common.Address, error) {
	return _USDCTokenPool.Contract.GetAllowList(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) GetAllowList() ([]common.Address, error) {
	return _USDCTokenPool.Contract.GetAllowList(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCaller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) GetAllowListEnabled() (bool, error) {
	return _USDCTokenPool.Contract.GetAllowListEnabled(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) GetAllowListEnabled() (bool, error) {
	return _USDCTokenPool.Contract.GetAllowListEnabled(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCaller) GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "getCurrentInboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _USDCTokenPool.Contract.GetCurrentInboundRateLimiterState(&_USDCTokenPool.CallOpts, remoteChainSelector)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _USDCTokenPool.Contract.GetCurrentInboundRateLimiterState(&_USDCTokenPool.CallOpts, remoteChainSelector)
}

func (_USDCTokenPool *USDCTokenPoolCaller) GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "getCurrentOutboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _USDCTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_USDCTokenPool.CallOpts, remoteChainSelector)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _USDCTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_USDCTokenPool.CallOpts, remoteChainSelector)
}

func (_USDCTokenPool *USDCTokenPoolCaller) GetDomain(opts *bind.CallOpts, chainSelector uint64) (USDCTokenPoolDomain, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "getDomain", chainSelector)

	if err != nil {
		return *new(USDCTokenPoolDomain), err
	}

	out0 := *abi.ConvertType(out[0], new(USDCTokenPoolDomain)).(*USDCTokenPoolDomain)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) GetDomain(chainSelector uint64) (USDCTokenPoolDomain, error) {
	return _USDCTokenPool.Contract.GetDomain(&_USDCTokenPool.CallOpts, chainSelector)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) GetDomain(chainSelector uint64) (USDCTokenPoolDomain, error) {
	return _USDCTokenPool.Contract.GetDomain(&_USDCTokenPool.CallOpts, chainSelector)
}

func (_USDCTokenPool *USDCTokenPoolCaller) GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "getRateLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) GetRateLimitAdmin() (common.Address, error) {
	return _USDCTokenPool.Contract.GetRateLimitAdmin(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) GetRateLimitAdmin() (common.Address, error) {
	return _USDCTokenPool.Contract.GetRateLimitAdmin(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCaller) GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "getRemotePools", remoteChainSelector)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _USDCTokenPool.Contract.GetRemotePools(&_USDCTokenPool.CallOpts, remoteChainSelector)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _USDCTokenPool.Contract.GetRemotePools(&_USDCTokenPool.CallOpts, remoteChainSelector)
}

func (_USDCTokenPool *USDCTokenPoolCaller) GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "getRemoteToken", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _USDCTokenPool.Contract.GetRemoteToken(&_USDCTokenPool.CallOpts, remoteChainSelector)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _USDCTokenPool.Contract.GetRemoteToken(&_USDCTokenPool.CallOpts, remoteChainSelector)
}

func (_USDCTokenPool *USDCTokenPoolCaller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) GetRmnProxy() (common.Address, error) {
	return _USDCTokenPool.Contract.GetRmnProxy(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) GetRmnProxy() (common.Address, error) {
	return _USDCTokenPool.Contract.GetRmnProxy(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) GetRouter() (common.Address, error) {
	return _USDCTokenPool.Contract.GetRouter(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) GetRouter() (common.Address, error) {
	return _USDCTokenPool.Contract.GetRouter(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) GetSupportedChains() ([]uint64, error) {
	return _USDCTokenPool.Contract.GetSupportedChains(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) GetSupportedChains() ([]uint64, error) {
	return _USDCTokenPool.Contract.GetSupportedChains(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) GetToken() (common.Address, error) {
	return _USDCTokenPool.Contract.GetToken(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _USDCTokenPool.Contract.GetToken(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCaller) GetTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "getTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) GetTokenDecimals() (uint8, error) {
	return _USDCTokenPool.Contract.GetTokenDecimals(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) GetTokenDecimals() (uint8, error) {
	return _USDCTokenPool.Contract.GetTokenDecimals(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCaller) ILocalDomainIdentifier(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "i_localDomainIdentifier")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) ILocalDomainIdentifier() (uint32, error) {
	return _USDCTokenPool.Contract.ILocalDomainIdentifier(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) ILocalDomainIdentifier() (uint32, error) {
	return _USDCTokenPool.Contract.ILocalDomainIdentifier(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCaller) IMessageTransmitterProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "i_messageTransmitterProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) IMessageTransmitterProxy() (common.Address, error) {
	return _USDCTokenPool.Contract.IMessageTransmitterProxy(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) IMessageTransmitterProxy() (common.Address, error) {
	return _USDCTokenPool.Contract.IMessageTransmitterProxy(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCaller) ISupportedUSDCVersion(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "i_supportedUSDCVersion")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) ISupportedUSDCVersion() (uint32, error) {
	return _USDCTokenPool.Contract.ISupportedUSDCVersion(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) ISupportedUSDCVersion() (uint32, error) {
	return _USDCTokenPool.Contract.ISupportedUSDCVersion(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCaller) ITokenMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "i_tokenMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) ITokenMessenger() (common.Address, error) {
	return _USDCTokenPool.Contract.ITokenMessenger(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) ITokenMessenger() (common.Address, error) {
	return _USDCTokenPool.Contract.ITokenMessenger(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCaller) IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "isRemotePool", remoteChainSelector, remotePoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _USDCTokenPool.Contract.IsRemotePool(&_USDCTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _USDCTokenPool.Contract.IsRemotePool(&_USDCTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_USDCTokenPool *USDCTokenPoolCaller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _USDCTokenPool.Contract.IsSupportedChain(&_USDCTokenPool.CallOpts, remoteChainSelector)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _USDCTokenPool.Contract.IsSupportedChain(&_USDCTokenPool.CallOpts, remoteChainSelector)
}

func (_USDCTokenPool *USDCTokenPoolCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) IsSupportedToken(token common.Address) (bool, error) {
	return _USDCTokenPool.Contract.IsSupportedToken(&_USDCTokenPool.CallOpts, token)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _USDCTokenPool.Contract.IsSupportedToken(&_USDCTokenPool.CallOpts, token)
}

func (_USDCTokenPool *USDCTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) Owner() (common.Address, error) {
	return _USDCTokenPool.Contract.Owner(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) Owner() (common.Address, error) {
	return _USDCTokenPool.Contract.Owner(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _USDCTokenPool.Contract.SupportsInterface(&_USDCTokenPool.CallOpts, interfaceId)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _USDCTokenPool.Contract.SupportsInterface(&_USDCTokenPool.CallOpts, interfaceId)
}

func (_USDCTokenPool *USDCTokenPoolCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) TypeAndVersion() (string, error) {
	return _USDCTokenPool.Contract.TypeAndVersion(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) TypeAndVersion() (string, error) {
	return _USDCTokenPool.Contract.TypeAndVersion(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_USDCTokenPool *USDCTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _USDCTokenPool.Contract.AcceptOwnership(&_USDCTokenPool.TransactOpts)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _USDCTokenPool.Contract.AcceptOwnership(&_USDCTokenPool.TransactOpts)
}

func (_USDCTokenPool *USDCTokenPoolTransactor) AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "addRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_USDCTokenPool *USDCTokenPoolSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.AddRemotePool(&_USDCTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.AddRemotePool(&_USDCTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_USDCTokenPool *USDCTokenPoolTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_USDCTokenPool *USDCTokenPoolSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.ApplyAllowListUpdates(&_USDCTokenPool.TransactOpts, removes, adds)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.ApplyAllowListUpdates(&_USDCTokenPool.TransactOpts, removes, adds)
}

func (_USDCTokenPool *USDCTokenPoolTransactor) ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "applyAuthorizedCallerUpdates", authorizedCallerArgs)
}

func (_USDCTokenPool *USDCTokenPoolSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.ApplyAuthorizedCallerUpdates(&_USDCTokenPool.TransactOpts, authorizedCallerArgs)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.ApplyAuthorizedCallerUpdates(&_USDCTokenPool.TransactOpts, authorizedCallerArgs)
}

func (_USDCTokenPool *USDCTokenPoolTransactor) ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "applyChainUpdates", remoteChainSelectorsToRemove, chainsToAdd)
}

func (_USDCTokenPool *USDCTokenPoolSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.ApplyChainUpdates(&_USDCTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.ApplyChainUpdates(&_USDCTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_USDCTokenPool *USDCTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_USDCTokenPool *USDCTokenPoolSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.LockOrBurn(&_USDCTokenPool.TransactOpts, lockOrBurnIn)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.LockOrBurn(&_USDCTokenPool.TransactOpts, lockOrBurnIn)
}

func (_USDCTokenPool *USDCTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_USDCTokenPool *USDCTokenPoolSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.ReleaseOrMint(&_USDCTokenPool.TransactOpts, releaseOrMintIn)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.ReleaseOrMint(&_USDCTokenPool.TransactOpts, releaseOrMintIn)
}

func (_USDCTokenPool *USDCTokenPoolTransactor) RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "removeRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_USDCTokenPool *USDCTokenPoolSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.RemoveRemotePool(&_USDCTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.RemoveRemotePool(&_USDCTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_USDCTokenPool *USDCTokenPoolTransactor) SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "setChainRateLimiterConfig", remoteChainSelector, outboundConfig, inboundConfig)
}

func (_USDCTokenPool *USDCTokenPoolSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.SetChainRateLimiterConfig(&_USDCTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.SetChainRateLimiterConfig(&_USDCTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_USDCTokenPool *USDCTokenPoolTransactor) SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "setChainRateLimiterConfigs", remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_USDCTokenPool *USDCTokenPoolSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.SetChainRateLimiterConfigs(&_USDCTokenPool.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.SetChainRateLimiterConfigs(&_USDCTokenPool.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_USDCTokenPool *USDCTokenPoolTransactor) SetDomains(opts *bind.TransactOpts, domains []USDCTokenPoolDomainUpdate) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "setDomains", domains)
}

func (_USDCTokenPool *USDCTokenPoolSession) SetDomains(domains []USDCTokenPoolDomainUpdate) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.SetDomains(&_USDCTokenPool.TransactOpts, domains)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) SetDomains(domains []USDCTokenPoolDomainUpdate) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.SetDomains(&_USDCTokenPool.TransactOpts, domains)
}

func (_USDCTokenPool *USDCTokenPoolTransactor) SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "setRateLimitAdmin", rateLimitAdmin)
}

func (_USDCTokenPool *USDCTokenPoolSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.SetRateLimitAdmin(&_USDCTokenPool.TransactOpts, rateLimitAdmin)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.SetRateLimitAdmin(&_USDCTokenPool.TransactOpts, rateLimitAdmin)
}

func (_USDCTokenPool *USDCTokenPoolTransactor) SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "setRouter", newRouter)
}

func (_USDCTokenPool *USDCTokenPoolSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.SetRouter(&_USDCTokenPool.TransactOpts, newRouter)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.SetRouter(&_USDCTokenPool.TransactOpts, newRouter)
}

func (_USDCTokenPool *USDCTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_USDCTokenPool *USDCTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.TransferOwnership(&_USDCTokenPool.TransactOpts, to)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.TransferOwnership(&_USDCTokenPool.TransactOpts, to)
}

type USDCTokenPoolAllowListAddIterator struct {
	Event *USDCTokenPoolAllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolAllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolAllowListAdd)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(USDCTokenPoolAllowListAdd)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *USDCTokenPoolAllowListAddIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolAllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolAllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterAllowListAdd(opts *bind.FilterOpts) (*USDCTokenPoolAllowListAddIterator, error) {

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolAllowListAddIterator{contract: _USDCTokenPool.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolAllowListAdd) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolAllowListAdd)
				if err := _USDCTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseAllowListAdd(log types.Log) (*USDCTokenPoolAllowListAdd, error) {
	event := new(USDCTokenPoolAllowListAdd)
	if err := _USDCTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolAllowListRemoveIterator struct {
	Event *USDCTokenPoolAllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolAllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolAllowListRemove)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(USDCTokenPoolAllowListRemove)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *USDCTokenPoolAllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolAllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolAllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterAllowListRemove(opts *bind.FilterOpts) (*USDCTokenPoolAllowListRemoveIterator, error) {

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolAllowListRemoveIterator{contract: _USDCTokenPool.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolAllowListRemove) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolAllowListRemove)
				if err := _USDCTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseAllowListRemove(log types.Log) (*USDCTokenPoolAllowListRemove, error) {
	event := new(USDCTokenPoolAllowListRemove)
	if err := _USDCTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolAuthorizedCallerAddedIterator struct {
	Event *USDCTokenPoolAuthorizedCallerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolAuthorizedCallerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolAuthorizedCallerAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(USDCTokenPoolAuthorizedCallerAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *USDCTokenPoolAuthorizedCallerAddedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolAuthorizedCallerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolAuthorizedCallerAdded struct {
	Caller common.Address
	Raw    types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*USDCTokenPoolAuthorizedCallerAddedIterator, error) {

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolAuthorizedCallerAddedIterator{contract: _USDCTokenPool.contract, event: "AuthorizedCallerAdded", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolAuthorizedCallerAdded) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolAuthorizedCallerAdded)
				if err := _USDCTokenPool.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseAuthorizedCallerAdded(log types.Log) (*USDCTokenPoolAuthorizedCallerAdded, error) {
	event := new(USDCTokenPoolAuthorizedCallerAdded)
	if err := _USDCTokenPool.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolAuthorizedCallerRemovedIterator struct {
	Event *USDCTokenPoolAuthorizedCallerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolAuthorizedCallerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolAuthorizedCallerRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(USDCTokenPoolAuthorizedCallerRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *USDCTokenPoolAuthorizedCallerRemovedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolAuthorizedCallerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolAuthorizedCallerRemoved struct {
	Caller common.Address
	Raw    types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*USDCTokenPoolAuthorizedCallerRemovedIterator, error) {

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolAuthorizedCallerRemovedIterator{contract: _USDCTokenPool.contract, event: "AuthorizedCallerRemoved", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolAuthorizedCallerRemoved) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolAuthorizedCallerRemoved)
				if err := _USDCTokenPool.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseAuthorizedCallerRemoved(log types.Log) (*USDCTokenPoolAuthorizedCallerRemoved, error) {
	event := new(USDCTokenPoolAuthorizedCallerRemoved)
	if err := _USDCTokenPool.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolChainAddedIterator struct {
	Event *USDCTokenPoolChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolChainAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(USDCTokenPoolChainAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *USDCTokenPoolChainAddedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolChainAdded struct {
	RemoteChainSelector       uint64
	RemoteToken               []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterChainAdded(opts *bind.FilterOpts) (*USDCTokenPoolChainAddedIterator, error) {

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolChainAddedIterator{contract: _USDCTokenPool.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolChainAdded) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolChainAdded)
				if err := _USDCTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseChainAdded(log types.Log) (*USDCTokenPoolChainAdded, error) {
	event := new(USDCTokenPoolChainAdded)
	if err := _USDCTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolChainConfiguredIterator struct {
	Event *USDCTokenPoolChainConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolChainConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolChainConfigured)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(USDCTokenPoolChainConfigured)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *USDCTokenPoolChainConfiguredIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolChainConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolChainConfigured struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterChainConfigured(opts *bind.FilterOpts) (*USDCTokenPoolChainConfiguredIterator, error) {

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolChainConfiguredIterator{contract: _USDCTokenPool.contract, event: "ChainConfigured", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolChainConfigured) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolChainConfigured)
				if err := _USDCTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseChainConfigured(log types.Log) (*USDCTokenPoolChainConfigured, error) {
	event := new(USDCTokenPoolChainConfigured)
	if err := _USDCTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolChainRemovedIterator struct {
	Event *USDCTokenPoolChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolChainRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(USDCTokenPoolChainRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *USDCTokenPoolChainRemovedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterChainRemoved(opts *bind.FilterOpts) (*USDCTokenPoolChainRemovedIterator, error) {

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolChainRemovedIterator{contract: _USDCTokenPool.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolChainRemoved) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolChainRemoved)
				if err := _USDCTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseChainRemoved(log types.Log) (*USDCTokenPoolChainRemoved, error) {
	event := new(USDCTokenPoolChainRemoved)
	if err := _USDCTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolConfigChangedIterator struct {
	Event *USDCTokenPoolConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolConfigChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(USDCTokenPoolConfigChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *USDCTokenPoolConfigChangedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolConfigChanged struct {
	Config RateLimiterConfig
	Raw    types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*USDCTokenPoolConfigChangedIterator, error) {

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolConfigChangedIterator{contract: _USDCTokenPool.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolConfigChanged) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolConfigChanged)
				if err := _USDCTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseConfigChanged(log types.Log) (*USDCTokenPoolConfigChanged, error) {
	event := new(USDCTokenPoolConfigChanged)
	if err := _USDCTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolConfigSetIterator struct {
	Event *USDCTokenPoolConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(USDCTokenPoolConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *USDCTokenPoolConfigSetIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolConfigSet struct {
	TokenMessenger common.Address
	Raw            types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterConfigSet(opts *bind.FilterOpts) (*USDCTokenPoolConfigSetIterator, error) {

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolConfigSetIterator{contract: _USDCTokenPool.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolConfigSet) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolConfigSet)
				if err := _USDCTokenPool.contract.UnpackLog(event, "ConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseConfigSet(log types.Log) (*USDCTokenPoolConfigSet, error) {
	event := new(USDCTokenPoolConfigSet)
	if err := _USDCTokenPool.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolDomainsSetIterator struct {
	Event *USDCTokenPoolDomainsSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolDomainsSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolDomainsSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(USDCTokenPoolDomainsSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *USDCTokenPoolDomainsSetIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolDomainsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolDomainsSet struct {
	Arg0 []USDCTokenPoolDomainUpdate
	Raw  types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterDomainsSet(opts *bind.FilterOpts) (*USDCTokenPoolDomainsSetIterator, error) {

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "DomainsSet")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolDomainsSetIterator{contract: _USDCTokenPool.contract, event: "DomainsSet", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchDomainsSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolDomainsSet) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "DomainsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolDomainsSet)
				if err := _USDCTokenPool.contract.UnpackLog(event, "DomainsSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseDomainsSet(log types.Log) (*USDCTokenPoolDomainsSet, error) {
	event := new(USDCTokenPoolDomainsSet)
	if err := _USDCTokenPool.contract.UnpackLog(event, "DomainsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolInboundRateLimitConsumedIterator struct {
	Event *USDCTokenPoolInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolInboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(USDCTokenPoolInboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *USDCTokenPoolInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolInboundRateLimitConsumedIterator{contract: _USDCTokenPool.contract, event: "InboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolInboundRateLimitConsumed)
				if err := _USDCTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseInboundRateLimitConsumed(log types.Log) (*USDCTokenPoolInboundRateLimitConsumed, error) {
	event := new(USDCTokenPoolInboundRateLimitConsumed)
	if err := _USDCTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolLockedOrBurnedIterator struct {
	Event *USDCTokenPoolLockedOrBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolLockedOrBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolLockedOrBurned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(USDCTokenPoolLockedOrBurned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *USDCTokenPoolLockedOrBurnedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolLockedOrBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolLockedOrBurned struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolLockedOrBurnedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolLockedOrBurnedIterator{contract: _USDCTokenPool.contract, event: "LockedOrBurned", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolLockedOrBurned)
				if err := _USDCTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseLockedOrBurned(log types.Log) (*USDCTokenPoolLockedOrBurned, error) {
	event := new(USDCTokenPoolLockedOrBurned)
	if err := _USDCTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolOutboundRateLimitConsumedIterator struct {
	Event *USDCTokenPoolOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolOutboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(USDCTokenPoolOutboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *USDCTokenPoolOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolOutboundRateLimitConsumedIterator{contract: _USDCTokenPool.contract, event: "OutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolOutboundRateLimitConsumed)
				if err := _USDCTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseOutboundRateLimitConsumed(log types.Log) (*USDCTokenPoolOutboundRateLimitConsumed, error) {
	event := new(USDCTokenPoolOutboundRateLimitConsumed)
	if err := _USDCTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolOwnershipTransferRequestedIterator struct {
	Event *USDCTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolOwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(USDCTokenPoolOwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *USDCTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*USDCTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolOwnershipTransferRequestedIterator{contract: _USDCTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolOwnershipTransferRequested)
				if err := _USDCTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*USDCTokenPoolOwnershipTransferRequested, error) {
	event := new(USDCTokenPoolOwnershipTransferRequested)
	if err := _USDCTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolOwnershipTransferredIterator struct {
	Event *USDCTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(USDCTokenPoolOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *USDCTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*USDCTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolOwnershipTransferredIterator{contract: _USDCTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolOwnershipTransferred)
				if err := _USDCTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*USDCTokenPoolOwnershipTransferred, error) {
	event := new(USDCTokenPoolOwnershipTransferred)
	if err := _USDCTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolRateLimitAdminSetIterator struct {
	Event *USDCTokenPoolRateLimitAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolRateLimitAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolRateLimitAdminSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(USDCTokenPoolRateLimitAdminSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *USDCTokenPoolRateLimitAdminSetIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolRateLimitAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolRateLimitAdminSet struct {
	RateLimitAdmin common.Address
	Raw            types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterRateLimitAdminSet(opts *bind.FilterOpts) (*USDCTokenPoolRateLimitAdminSetIterator, error) {

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolRateLimitAdminSetIterator{contract: _USDCTokenPool.contract, event: "RateLimitAdminSet", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolRateLimitAdminSet) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolRateLimitAdminSet)
				if err := _USDCTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseRateLimitAdminSet(log types.Log) (*USDCTokenPoolRateLimitAdminSet, error) {
	event := new(USDCTokenPoolRateLimitAdminSet)
	if err := _USDCTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolReleasedOrMintedIterator struct {
	Event *USDCTokenPoolReleasedOrMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolReleasedOrMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolReleasedOrMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(USDCTokenPoolReleasedOrMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *USDCTokenPoolReleasedOrMintedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolReleasedOrMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolReleasedOrMinted struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Recipient           common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolReleasedOrMintedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolReleasedOrMintedIterator{contract: _USDCTokenPool.contract, event: "ReleasedOrMinted", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolReleasedOrMinted)
				if err := _USDCTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseReleasedOrMinted(log types.Log) (*USDCTokenPoolReleasedOrMinted, error) {
	event := new(USDCTokenPoolReleasedOrMinted)
	if err := _USDCTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolRemotePoolAddedIterator struct {
	Event *USDCTokenPoolRemotePoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolRemotePoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolRemotePoolAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(USDCTokenPoolRemotePoolAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *USDCTokenPoolRemotePoolAddedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolRemotePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolRemotePoolAdded struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolRemotePoolAddedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolRemotePoolAddedIterator{contract: _USDCTokenPool.contract, event: "RemotePoolAdded", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolRemotePoolAdded)
				if err := _USDCTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseRemotePoolAdded(log types.Log) (*USDCTokenPoolRemotePoolAdded, error) {
	event := new(USDCTokenPoolRemotePoolAdded)
	if err := _USDCTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolRemotePoolRemovedIterator struct {
	Event *USDCTokenPoolRemotePoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolRemotePoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolRemotePoolRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(USDCTokenPoolRemotePoolRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *USDCTokenPoolRemotePoolRemovedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolRemotePoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolRemotePoolRemoved struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolRemotePoolRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolRemotePoolRemovedIterator{contract: _USDCTokenPool.contract, event: "RemotePoolRemoved", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolRemotePoolRemoved)
				if err := _USDCTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseRemotePoolRemoved(log types.Log) (*USDCTokenPoolRemotePoolRemoved, error) {
	event := new(USDCTokenPoolRemotePoolRemoved)
	if err := _USDCTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolRouterUpdatedIterator struct {
	Event *USDCTokenPoolRouterUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolRouterUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolRouterUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(USDCTokenPoolRouterUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *USDCTokenPoolRouterUpdatedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolRouterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolRouterUpdated struct {
	OldRouter common.Address
	NewRouter common.Address
	Raw       types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterRouterUpdated(opts *bind.FilterOpts) (*USDCTokenPoolRouterUpdatedIterator, error) {

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolRouterUpdatedIterator{contract: _USDCTokenPool.contract, event: "RouterUpdated", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolRouterUpdated) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolRouterUpdated)
				if err := _USDCTokenPool.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseRouterUpdated(log types.Log) (*USDCTokenPoolRouterUpdated, error) {
	event := new(USDCTokenPoolRouterUpdated)
	if err := _USDCTokenPool.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (USDCTokenPoolAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (USDCTokenPoolAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (USDCTokenPoolAuthorizedCallerAdded) Topic() common.Hash {
	return common.HexToHash("0xeb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef")
}

func (USDCTokenPoolAuthorizedCallerRemoved) Topic() common.Hash {
	return common.HexToHash("0xc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda77580")
}

func (USDCTokenPoolChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (USDCTokenPoolChainConfigured) Topic() common.Hash {
	return common.HexToHash("0x0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b")
}

func (USDCTokenPoolChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (USDCTokenPoolConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (USDCTokenPoolConfigSet) Topic() common.Hash {
	return common.HexToHash("0x2e902d38f15b233cbb63711add0fca4545334d3a169d60c0a616494d7eea9544")
}

func (USDCTokenPoolDomainsSet) Topic() common.Hash {
	return common.HexToHash("0xe6d14ea297366c7bc1265d289d924bfd8b9afb148eb972b481f70da41c842cf5")
}

func (USDCTokenPoolInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (USDCTokenPoolLockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
}

func (USDCTokenPoolOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944")
}

func (USDCTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (USDCTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (USDCTokenPoolRateLimitAdminSet) Topic() common.Hash {
	return common.HexToHash("0x44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174")
}

func (USDCTokenPoolReleasedOrMinted) Topic() common.Hash {
	return common.HexToHash("0xfc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0")
}

func (USDCTokenPoolRemotePoolAdded) Topic() common.Hash {
	return common.HexToHash("0x7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea")
}

func (USDCTokenPoolRemotePoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76")
}

func (USDCTokenPoolRouterUpdated) Topic() common.Hash {
	return common.HexToHash("0x02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f1684")
}

func (_USDCTokenPool *USDCTokenPool) Address() common.Address {
	return _USDCTokenPool.address
}

type USDCTokenPoolInterface interface {
	GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetDomain(opts *bind.CallOpts, chainSelector uint64) (USDCTokenPoolDomain, error)

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

	ISupportedUSDCVersion(opts *bind.CallOpts) (uint32, error)

	ITokenMessenger(opts *bind.CallOpts) (common.Address, error)

	IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error)

	IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error)

	IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error)

	ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error)

	ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error)

	SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error)

	SetDomains(opts *bind.TransactOpts, domains []USDCTokenPoolDomainUpdate) (*types.Transaction, error)

	SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterAllowListAdd(opts *bind.FilterOpts) (*USDCTokenPoolAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*USDCTokenPoolAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*USDCTokenPoolAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*USDCTokenPoolAllowListRemove, error)

	FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*USDCTokenPoolAuthorizedCallerAddedIterator, error)

	WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolAuthorizedCallerAdded) (event.Subscription, error)

	ParseAuthorizedCallerAdded(log types.Log) (*USDCTokenPoolAuthorizedCallerAdded, error)

	FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*USDCTokenPoolAuthorizedCallerRemovedIterator, error)

	WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolAuthorizedCallerRemoved) (event.Subscription, error)

	ParseAuthorizedCallerRemoved(log types.Log) (*USDCTokenPoolAuthorizedCallerRemoved, error)

	FilterChainAdded(opts *bind.FilterOpts) (*USDCTokenPoolChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*USDCTokenPoolChainAdded, error)

	FilterChainConfigured(opts *bind.FilterOpts) (*USDCTokenPoolChainConfiguredIterator, error)

	WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolChainConfigured) (event.Subscription, error)

	ParseChainConfigured(log types.Log) (*USDCTokenPoolChainConfigured, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*USDCTokenPoolChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*USDCTokenPoolChainRemoved, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*USDCTokenPoolConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*USDCTokenPoolConfigChanged, error)

	FilterConfigSet(opts *bind.FilterOpts) (*USDCTokenPoolConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*USDCTokenPoolConfigSet, error)

	FilterDomainsSet(opts *bind.FilterOpts) (*USDCTokenPoolDomainsSetIterator, error)

	WatchDomainsSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolDomainsSet) (event.Subscription, error)

	ParseDomainsSet(log types.Log) (*USDCTokenPoolDomainsSet, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolInboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*USDCTokenPoolInboundRateLimitConsumed, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolLockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*USDCTokenPoolLockedOrBurned, error)

	FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolOutboundRateLimitConsumedIterator, error)

	WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseOutboundRateLimitConsumed(log types.Log) (*USDCTokenPoolOutboundRateLimitConsumed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*USDCTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*USDCTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*USDCTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*USDCTokenPoolOwnershipTransferred, error)

	FilterRateLimitAdminSet(opts *bind.FilterOpts) (*USDCTokenPoolRateLimitAdminSetIterator, error)

	WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolRateLimitAdminSet) (event.Subscription, error)

	ParseRateLimitAdminSet(log types.Log) (*USDCTokenPoolRateLimitAdminSet, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*USDCTokenPoolReleasedOrMinted, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolRemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*USDCTokenPoolRemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolRemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*USDCTokenPoolRemotePoolRemoved, error)

	FilterRouterUpdated(opts *bind.FilterOpts) (*USDCTokenPoolRouterUpdatedIterator, error)

	WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolRouterUpdated) (event.Subscription, error)

	ParseRouterUpdated(log types.Log) (*USDCTokenPoolRouterUpdated, error)

	Address() common.Address
}
