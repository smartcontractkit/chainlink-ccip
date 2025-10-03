// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package burn_mint_token_pool_v2

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

type ClientEVM2AnyMessage struct {
	Receiver     []byte
	Data         []byte
	TokenAmounts []ClientEVMTokenAmount
	FeeToken     common.Address
	ExtraArgs    []byte
}

type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
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

type TokenPoolV2CCVConfigArg struct {
	RemoteChainSelector uint64
	OutboundCCVs        []common.Address
	InboundCCVs         []common.Address
}

var BurnMintTokenPoolV2MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIBurnMintERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPoolV2.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredInboundCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredOutboundCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRouter\",\"inputs\":[{\"name\":\"newRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RouterUpdated\",\"inputs\":[{\"name\":\"oldRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCV\",\"inputs\":[{\"name\":\"ccv\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x6101008060405234610363576151b6803803809161001d82856103e2565b833981019060a0818303126103635780516001600160a01b038116908190036103635761004c60208301610405565b60408301519091906001600160401b0381116103635783019380601f86011215610363578451946001600160401b0386116103cc578560051b90602082019661009860405198896103e2565b875260208088019282010192831161036357602001905b8282106103b4575050506100d160806100ca60608601610413565b9401610413565b9233156103a357600180546001600160a01b0319163317905581158015610392575b8015610381575b610370578160209160049360805260c0526040519283809263313ce56760e01b82525afa6000918161032f575b50610304575b5060a052600480546001600160a01b0319166001600160a01b03929092169190911790558051151560e08190526101e6575b604051614bee90816105c882396080518181816121f0015281816123c001528181612756015281816127ce01528181613119015261330e015260a0518181816126dd01528181613428015281816139cc0152613a4f015260c051818181610c0c0152818161228c01526131ba015260e051818181610b9c01528181611f3a01526131fd0152f35b60405160206101f581836103e2565b60008252600036813760e051156102f35760005b8251811015610270576001906001600160a01b036102278286610427565b51168361023382610469565b610240575b505001610209565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a13883610238565b50905060005b82518110156102ea576001906001600160a01b036102948286610427565b511680156102e457836102a682610567565b6102b4575b50505b01610276565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a138836102ab565b506102ae565b5050503861015f565b6335f4a7b360e01b60005260046000fd5b60ff1660ff8216818103610318575061012d565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d602011610368575b8161034b602093836103e2565b810103126103635761035c90610405565b9038610127565b600080fd5b3d915061033e565b630a64406560e11b60005260046000fd5b506001600160a01b038116156100fa565b506001600160a01b038416156100f3565b639b15e16f60e01b60005260046000fd5b602080916103c184610413565b8152019101906100af565b634e487b7160e01b600052604160045260246000fd5b601f909101601f19168101906001600160401b038211908210176103cc57604052565b519060ff8216820361036357565b51906001600160a01b038216820361036357565b805182101561043b5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b805482101561043b5760005260206000200190600090565b600081815260036020526040902054801561056057600019810181811161054a5760025460001981019190821161054a578181036104f9575b50505060025480156104e357600019016104bd816002610451565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61053261050a61051b936002610451565b90549060031b1c9283926002610451565b819391549060031b91821b91600019901b19161790565b905560005260036020526040600020553880806104a2565b634e487b7160e01b600052601160045260246000fd5b5050600090565b806000526003602052604060002054156000146105c157600254680100000000000000008110156103cc576105a861051b8260018594016002556002610451565b9055600254906000526003602052604060002055600190565b5060009056fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a71461287157508063181f5a77146127f257806321df0da714612783578063240028e81461270157806324f65ee7146126a5578063390775371461211d5780634c5ef0ed146120b85780634f71592c1461208357806354c8a4f314611f0657806362ddd3c414611e8257806366b852da14611dd95780636d3d1a5814611d8757806379ba509714611ca25780637d54534e14611bf55780638926f54f14611b915780638da5cb5b14611b3f578063962d40201461199b5780639a4575b9146119115780639f68f673146118d9578063a42a7b8b14611754578063a7cd63b7146116b5578063acfecf9114611591578063af58d59f1461152a578063b0f479a1146114d8578063b79019b51461120a578063b7946580146111b3578063b82a8aba14611106578063c0d786551461100e578063c4bffe2b14610ec5578063c75eea9c14610dff578063cf7401f314610c30578063dc0bd97114610bc1578063e0351e1314610b66578063e8a1da17146102915763f2fde38b146101a257600080fd5b3461028e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028e5773ffffffffffffffffffffffffffffffffffffffff6101ee612b03565b6101f6613b71565b1633811461026657807fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b80fd5b503461028e576102a036612d37565b939190926102ac613b71565b82915b8083106109d1575050508063ffffffff4216917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee1843603015b858210156109cd578160051b850135818112156109c957850190610120823603126109c9576040519561031a87612a2b565b823567ffffffffffffffff811681036109c4578752602083013567ffffffffffffffff81116109c05783019536601f880112156109c05786359661035d88613016565b9761036b604051998a612a63565b8089526020808a019160051b830101903682116109bc5760208301905b828210610989575050505060208801968752604084013567ffffffffffffffff8111610985576103bb9036908601612bc3565b9860408901998a526103e56103d33660608801612ea3565b9560608b0196875260c0369101612ea3565b9660808a019788526103f786516140b9565b61040188516140b9565b8a51511561095d5761041d67ffffffffffffffff8b51166148eb565b156109265767ffffffffffffffff8a5116815260076020526040812061055d87516fffffffffffffffffffffffffffffffff604082015116906105186fffffffffffffffffffffffffffffffff6020830151169151151583608060405161048381612a2b565b858152602081018c905260408101849052606081018690520152855474ff000000000000000000000000000000000000000091151560a01b919091167fffffffffffffffffffffff0000000000000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff84161773ffffffff0000000000000000000000000000000060808b901b1617178555565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176001830155565b61068389516fffffffffffffffffffffffffffffffff6040820151169061063e6fffffffffffffffffffffffffffffffff602083015116915115158360806040516105a781612a2b565b858152602081018c9052604081018490526060810186905201526002860180547fffffffffffffffffffffff000000000000000000000000000000000000000000166fffffffffffffffffffffffffffffffff85161773ffffffff0000000000000000000000000000000060808c901b161791151560a01b74ff000000000000000000000000000000000000000016919091179055565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176003830155565b60048c5191019080519067ffffffffffffffff82116108f9576106a68354613618565b601f81116108be575b50602090601f831160011461081f576106fd9291859183610814575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b805b89518051821015610738579061073260019261072b838f67ffffffffffffffff90511692613604565b5190613bbc565b01610702565b5050975097987f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c29295939661080667ffffffffffffffff600197949c51169251935191516107d261079d60405196879687526101006020880152610100870190612aa4565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a10190939492916102e8565b0151905038806106cb565b83855281852091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416865b8181106108a6575090846001959493921061086f575b505050811b019055610700565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388080610862565b9293602060018192878601518155019501930161084c565b6108e99084865260208620601f850160051c810191602086106108ef575b601f0160051c0190613920565b386106af565b90915081906108dc565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60249067ffffffffffffffff8b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b807f14c880ca0000000000000000000000000000000000000000000000000000000060049252fd5b8680fd5b813567ffffffffffffffff81116109b8576020916109ad8392833691890101612bc3565b815201910190610388565b8a80fd5b8880fd5b8580fd5b600080fd5b8380fd5b8280f35b9092919367ffffffffffffffff6109f16109ec87858861309b565b612f61565b16956109fc8761462c565b15610b3a578684526007602052610a1860056040862001614433565b94845b8651811015610a51576001908987526007602052610a4a60056040892001610a43838b613604565b5190614757565b5001610a1b565b5093945094909580855260076020526005604086208681558660018201558660028201558660038201558660048201610a8a8154613618565b80610af9575b5050500180549086815581610adb575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a10191909493946102af565b865260208620908101905b81811015610aa057868155600101610ae6565b601f8111600114610b0f5750555b863880610a90565b81835260208320610b2a91601f01861c810190600101613920565b8082528160208120915555610b07565b602484887f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b503461028e57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028e5760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b503461028e57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028e57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461028e5760e07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028e57610c68612b47565b9060607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc36011261028e57604051610c9f81612a47565b6024358015158103610dfb5781526044356fffffffffffffffffffffffffffffffff81168103610dfb5760208201526064356fffffffffffffffffffffffffffffffff81168103610dfb57604082015260607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c360112610df75760405190610d2682612a47565b60843580151581036109c957825260a4356fffffffffffffffffffffffffffffffff811681036109c957602083015260c4356fffffffffffffffffffffffffffffffff811681036109c957604083015273ffffffffffffffffffffffffffffffffffffffff6009541633141580610dd5575b610da957610da69293613e26565b80f35b6024837f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b5073ffffffffffffffffffffffffffffffffffffffff60015416331415610d98565b5080fd5b8280fd5b503461028e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028e57610e68610e636040610ec19367ffffffffffffffff610e4c612b47565b610e5461376a565b50168152600760205220613795565b613f63565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b0390f35b503461028e57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028e57604051906005548083528260208101600584526020842092845b818110610ff5575050610f2392500383612a63565b8151610f47610f3182613016565b91610f3f6040519384612a63565b808352613016565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602083019301368437805b8451811015610fa6578067ffffffffffffffff610f9360019388613604565b5116610f9f8286613604565b5201610f74565b50925090604051928392602084019060208552518091526040840192915b818110610fd2575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101610fc4565b8454835260019485019487945060209093019201610f0e565b503461028e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028e5773ffffffffffffffffffffffffffffffffffffffff61105b612b03565b611063613b71565b1680156110de5760407f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f16849160045490807fffffffffffffffffffffffff000000000000000000000000000000000000000083161760045573ffffffffffffffffffffffffffffffffffffffff8351921682526020820152a180f35b6004827f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b503461028e5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028e5761113e612b47565b5060243567ffffffffffffffff8111610df7577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc60a0913603011261028e57611185612be1565b5060643567ffffffffffffffff8111610df757906111a96020923690600401612bf2565b5050604051908152f35b503461028e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028e57610ec16111f66111f1612b47565b613937565b604051918291602083526020830190612aa4565b503461028e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028e5760043567ffffffffffffffff8111610df75761125a903690600401612d06565b611262613b71565b825b81811061126f578380f35b61127d6109ec8284866137fb565b61129561128b8385876137fb565b602081019061383b565b907fb0897119e8510f887b892cbc4c8506fc51d9849fd90afae4fd065e705f2d0f6c6112cf6112c586888a6137fb565b604081019061383b565b9190926112dc8582613ff5565b6112e68385613ff5565b604051946112f386612a0f565b6112fe36828461302e565b865261134867ffffffffffffffff61131736878961302e565b9860208901998a521695869561133a60405195869560408752604087019161388f565b91848303602086015261388f565b0390a28652600a60205260408620905180519067ffffffffffffffff82116114ab576801000000000000000082116114ab576020908354838555808410611491575b500182885260208820885b838110611467575050505060010190519081519167ffffffffffffffff831161143a5768010000000000000000831161143a576020908254848455808510611420575b500190865260208620865b8381106113f65750505050600101611264565b600190602073ffffffffffffffffffffffffffffffffffffffff85511694019381840155016113e3565b838952828920611434918101908601613920565b386113d8565b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501611395565b848a52828a206114a5918101908501613920565b3861138a565b6024887f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b503461028e57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028e57602073ffffffffffffffffffffffffffffffffffffffff60045416604051908152f35b503461028e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028e57610e68610e6360026040610ec19467ffffffffffffffff611579612b47565b61158161376a565b5016815260076020522001613795565b503461028e5767ffffffffffffffff6115a936612da3565b9290916115b4613b71565b16916115cd836000526006602052604060002054151590565b156116895782845260076020526115fc600560408620016115ef368486612b5e565b6020815191012090614757565b1561164157907f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d769161163b60405192839260208452602084019161372b565b0390a280f35b82611685836040519384937f74f23c7c000000000000000000000000000000000000000000000000000000008552600485015260406024850152604484019161372b565b0390fd5b602484847f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b503461028e57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028e5760405160028054808352908352909160208301917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace915b81811061173e57610ec18561173281870382612a63565b60405191829182612cb6565b825484526020909301926001928301920161171b565b503461028e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028e5767ffffffffffffffff611795612b47565b16815260076020526117ac60056040832001614433565b80517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06117f16117db83613016565b926117e96040519485612a63565b808452613016565b01835b8181106118c8575050825b8251811015611845578061181560019285613604565b51855260086020526118296040862061366b565b6118338285613604565b5261183e8184613604565b50016117ff565b81846040519182916020830160208452825180915260408401602060408360051b870101940192905b82821061187d57505050500390f35b919360206118b8827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851612aa4565b960192019201859493919261186e565b8060606020809386010152016117f4565b503461028e5761173260016040610ec19367ffffffffffffffff6118fc36612c20565b505050509050168152600a6020522001612fb3565b503461028e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028e576004359067ffffffffffffffff821161028e5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc833603011261028e57610ec161198f836004016130ea565b60405191829182612e02565b503461028e5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028e5760043567ffffffffffffffff8111610df7576119eb903690600401612d06565b60243567ffffffffffffffff81116109c957611a0b903690600401612e55565b60449291923567ffffffffffffffff81116109c057611a2e903690600401612e55565b91909273ffffffffffffffffffffffffffffffffffffffff6009541633141580611b1d575b611af157818114801590611ae7575b611abf57865b818110611a73578780f35b80611ab9611a876109ec600194868c61309b565b611a9283878b6130da565b611ab3611aab611aa3868b8d6130da565b923690612ea3565b913690612ea3565b91613e26565b01611a68565b6004877f568efce2000000000000000000000000000000000000000000000000000000008152fd5b5082811415611a62565b6024877f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b5073ffffffffffffffffffffffffffffffffffffffff60015416331415611a53565b503461028e57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028e57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b503461028e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028e576020611beb67ffffffffffffffff611bd7612b47565b166000526006602052604060002054151590565b6040519015158152f35b503461028e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028e577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174602073ffffffffffffffffffffffffffffffffffffffff611c65612b03565b611c6d613b71565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000006009541617600955604051908152a180f35b503461028e57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028e57805473ffffffffffffffffffffffffffffffffffffffff81163303611d5f577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516825573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b503461028e57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028e57602073ffffffffffffffffffffffffffffffffffffffff60095416604051908152f35b503461028e5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028e576004359067ffffffffffffffff821161028e5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc833603011261028e576024359067ffffffffffffffff821161028e57610ec161198f84611e6f3660048701612bf2565b5050611e79613082565b506004016130ea565b503461028e57611e9136612da3565b611e9d93929193613b71565b67ffffffffffffffff8216611ebf816000526006602052604060002054151590565b15611edb5750610da69293611ed5913691612b5e565b90613bbc565b7f1e670e4b000000000000000000000000000000000000000000000000000000008452600452602483fd5b503461028e57611f3090611f38611f1c36612d37565b9591611f29939193613b71565b369161302e565b93369161302e565b7f00000000000000000000000000000000000000000000000000000000000000001561205b57815b8351811015611fd3578073ffffffffffffffffffffffffffffffffffffffff611f8b60019387613604565b5116611f9681614496565b611fa2575b5001611f60565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a138611f9b565b5090805b8251811015612057578073ffffffffffffffffffffffffffffffffffffffff61200260019386613604565b51168015612051576120138161488b565b612020575b505b01611fd7565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a184612018565b5061201a565b5080f35b6004827f35f4a7b3000000000000000000000000000000000000000000000000000000008152fd5b503461028e576117326040610ec19267ffffffffffffffff6120a436612c20565b505050509050168152600a60205220612fb3565b503461028e5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028e576120f0612b47565b906024359067ffffffffffffffff821161028e576020611beb846121173660048701612bc3565b90612f76565b503461028e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028e5760043567ffffffffffffffff8111610df757806004016101007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8336030112610dfb578260405161219d816129c4565b526121ca6121c06121bb6121b460c4860185612eef565b3691612b5e565b613959565b6064840135613a4c565b91608481016121d881612f40565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691160361265b5750602481019177ffffffffffffffff0000000000000000000000000000000061223f84612f61565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561253c57869161263c575b506126145767ffffffffffffffff6122d384612f61565b166122eb816000526006602052604060002054151590565b156125e957602073ffffffffffffffffffffffffffffffffffffffff60045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa90811561253c5786916125ba575b501561258e5761236283612f61565b9061237860a48401926121176121b48585612eef565b15612547575050906044839267ffffffffffffffff61239684612f61565b1680875260076020526123e86002604089200173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001696879161499a565b6040805173ffffffffffffffffffffffffffffffffffffffff87168152602081018890527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a2019061243b82612f40565b85843b1561028e576040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff929092166004830152602482018690528160448183885af1801561253c579273ffffffffffffffffffffffffffffffffffffffff6124fc6124f660809560209a67ffffffffffffffff967ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09961252c575b5050612f61565b92612f40565b60405196875233898801521660408601528560608601521692a280604051612523816129c4565b52604051908152f35b8161253691612a63565b386124ef565b6040513d88823e3d90fd5b6125519250612eef565b6116856040519283927f24eb47e500000000000000000000000000000000000000000000000000000000845260206004850152602484019161372b565b6024857f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b6125dc915060203d6020116125e2575b6125d48183612a63565b810190613b59565b38612353565b503d6125ca565b7fa9902c7e000000000000000000000000000000000000000000000000000000008652600452602485fd5b6004857f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b612655915060203d6020116125e2576125d48183612a63565b386122bc565b8473ffffffffffffffffffffffffffffffffffffffff61267c602493612f40565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835216600452fd5b503461028e57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028e57602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461028e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028e5760209061273c612b03565b905073ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b503461028e57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028e57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461028e57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028e5750610ec1604051612833604082612a63565b601b81527f4275726e4d696e74546f6b656e506f6f6c563220312e372d64657600000000006020820152604051918291602083526020830190612aa4565b905034610df75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610df7576004357fffffffff000000000000000000000000000000000000000000000000000000008116809103610dfb57602092507ff208a58f00000000000000000000000000000000000000000000000000000000811490811561299a575b811561290c575b5015158152f35b7faff2afbf00000000000000000000000000000000000000000000000000000000811491508115612970575b8115612946575b5038612905565b7f01ffc9a7000000000000000000000000000000000000000000000000000000009150143861293f565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150612938565b7f0e8b773f00000000000000000000000000000000000000000000000000000000811491506128fe565b6020810190811067ffffffffffffffff8211176129e057604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff8211176129e057604052565b60a0810190811067ffffffffffffffff8211176129e057604052565b6060810190811067ffffffffffffffff8211176129e057604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176129e057604052565b919082519283825260005b848110612aee5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201612aaf565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036109c457565b359073ffffffffffffffffffffffffffffffffffffffff821682036109c457565b6004359067ffffffffffffffff821682036109c457565b92919267ffffffffffffffff82116129e05760405191612ba6601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200184612a63565b8294818452818301116109c4578281602093846000960137010152565b9080601f830112156109c457816020612bde93359101612b5e565b90565b6044359061ffff821682036109c457565b9181601f840112156109c45782359167ffffffffffffffff83116109c457602083818601950101116109c457565b60a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126109c45760043573ffffffffffffffffffffffffffffffffffffffff811681036109c4579160243567ffffffffffffffff811681036109c457916044359160643561ffff811681036109c457916084359067ffffffffffffffff82116109c457612cb291600401612bf2565b9091565b602060408183019282815284518094520192019060005b818110612cda5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101612ccd565b9181601f840112156109c45782359167ffffffffffffffff83116109c4576020808501948460051b0101116109c457565b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126109c45760043567ffffffffffffffff81116109c45781612d8091600401612d06565b929092916024359067ffffffffffffffff82116109c457612cb291600401612d06565b9060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126109c45760043567ffffffffffffffff811681036109c457916024359067ffffffffffffffff82116109c457612cb291600401612bf2565b90612bde91602081526020612e2283516040838501526060840190612aa4565b9201519060407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152612aa4565b9181601f840112156109c45782359167ffffffffffffffff83116109c457602080850194606085020101116109c457565b35906fffffffffffffffffffffffffffffffff821682036109c457565b91908260609103126109c457604051612ebb81612a47565b809280359081151582036109c4576040612eea9181938552612edf60208201612e86565b602086015201612e86565b910152565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156109c4570180359067ffffffffffffffff82116109c4576020019181360383136109c457565b3573ffffffffffffffffffffffffffffffffffffffff811681036109c45790565b3567ffffffffffffffff811681036109c45790565b9067ffffffffffffffff612bde92166000526007602052600560406000200190602081519101209060019160005201602052604060002054151590565b906040519182815491828252602082019060005260206000209260005b818110612fe7575050612fe592500383612a63565b565b845473ffffffffffffffffffffffffffffffffffffffff16835260019485019487945060209093019201612fd0565b67ffffffffffffffff81116129e05760051b60200190565b92919061303a81613016565b936130486040519586612a63565b602085838152019160051b81019283116109c457905b82821061306a57505050565b6020809161307784612b26565b81520191019061305e565b6040519061308f82612a0f565b60606020838281520152565b91908110156130ab5760051b0190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b91908110156130ab576060020190565b6130f2613082565b50608081019061310182612f40565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691161491600092156135e35750602081019177ffffffffffffffff0000000000000000000000000000000061316d84612f61565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156134835782916135c4575b5061359c576131fb60408301612f40565b7f0000000000000000000000000000000000000000000000000000000000000000613549575b5067ffffffffffffffff61323484612f61565b1661324c816000526006602052604060002054151590565b1561351d57602073ffffffffffffffffffffffffffffffffffffffff60045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa80156134835782906134ba575b73ffffffffffffffffffffffffffffffffffffffff915016330361348e579067ffffffffffffffff9160606132e185612f61565b9201359283921680825260076020526133366040832073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001694859161499a565b6040805173ffffffffffffffffffffffffffffffffffffffff85168152602081018690527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789449190a2813b1561028e576040517f42966c68000000000000000000000000000000000000000000000000000000008152836004820152818160248183875af180156134835761342195936111f195937ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1093606093613473575b505067ffffffffffffffff61340886612f61565b16936040519182523360208301526040820152a2612f61565b60405160ff7f00000000000000000000000000000000000000000000000000000000000000001660208201526020815261345c604082612a63565b6040519161346983612a0f565b8252602082015290565b8161347d91612a63565b386133f4565b6040513d84823e3d90fd5b807f728fe07b000000000000000000000000000000000000000000000000000000006024925233600452fd5b506020813d602011613515575b816134d460209383612a63565b81010312610df7575173ffffffffffffffffffffffffffffffffffffffff81168103610df75773ffffffffffffffffffffffffffffffffffffffff906132ad565b3d91506134c7565b602492507fa9902c7e000000000000000000000000000000000000000000000000000000008252600452fd5b73ffffffffffffffffffffffffffffffffffffffff168082526003602052604082205461322157602492507fd0d25976000000000000000000000000000000000000000000000000000000008252600452fd5b807f53ad11d80000000000000000000000000000000000000000000000000000000060049252fd5b6135dd915060203d6020116125e2576125d48183612a63565b386131ea565b8273ffffffffffffffffffffffffffffffffffffffff61267c602493612f40565b80518210156130ab5760209160051b010190565b90600182811c92168015613661575b602083101461363257565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691613627565b906040519182600082549261367f84613618565b80845293600181169081156136eb57506001146136a4575b50612fe592500383612a63565b90506000929192526020600020906000915b8183106136cf575050906020612fe59282010138613697565b60209193508060019154838589010152019101909184926136b6565b60209350612fe59592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138613697565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b6040519061377782612a2b565b60006080838281528260208201528260408201528260608201520152565b906040516137a281612a2b565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b91908110156130ab5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa1813603018212156109c4570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156109c4570180359067ffffffffffffffff82116109c457602001918160051b360383136109c457565b9160209082815201919060005b8181106138a95750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff6138d088612b26565b16815201940192910161389c565b818102929181159184041417156138f157565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b81811061392b575050565b60008155600101613920565b67ffffffffffffffff166000526007602052612bde600460406000200161366b565b805180156139c85760200361398a576020818051810103126109c45760208101519060ff821161398a575060ff1690565b611685906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190612aa4565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff82116138f157565b60ff16604d81116138f157600a0a90565b8115613a1d570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff811692828414613b5257828411613b285790613a91916139ee565b91604d60ff8416118015613aef575b613ab957505090613ab3612bde92613a02565b906138de565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b50613af983613a02565b8015613a1d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048411613aa0565b613b31916139ee565b91604d60ff841611613ab957505090613b4c612bde92613a02565b90613a13565b5050505090565b908160209103126109c4575180151581036109c45790565b73ffffffffffffffffffffffffffffffffffffffff600154163303613b9257565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b90805115613dfc5767ffffffffffffffff81516020830120921691826000526007602052613bf1816005604060002001614945565b15613db85760005260086020526040600020815167ffffffffffffffff81116129e057613c1e8254613618565b601f8111613d86575b506020601f8211600114613cc05791613c9a827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9593613cb095600091613cb5575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055604051918291602083526020830190612aa4565b0390a2565b905084015138613c69565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b818110613d6e575092613cb09492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610613d37575b5050811b0190556111f6565b8501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880613d2b565b9192602060018192868a015181550194019201613cf0565b613db290836000526020600020601f840160051c810191602085106108ef57601f0160051c0190613920565b38613c27565b50906116856040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190612aa4565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b67ffffffffffffffff166000818152600660205260409020549092919015613f285791613f2560e092613ef185613e7d7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b976140b9565b846000526007602052613e94816040600020614200565b613e9d836140b9565b846000526007602052613eb7836002604060002001614200565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b919082039182116138f157565b613f6b61376a565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff8083511691613fc86020850193613fc2613fb563ffffffff87511642613f56565b85608089015116906138de565b90613fe8565b80821015613fe157505b16825263ffffffff4216905290565b9050613fd2565b919082018092116138f157565b9060005b81811061400557505050565b600181018082116138f1575b8281106140215750600101613ff9565b61403461402f83858761309b565b612f40565b73ffffffffffffffffffffffffffffffffffffffff8061405861402f85888a61309b565b1691161461406857600101614011565b73ffffffffffffffffffffffffffffffffffffffff61408b61402f84868861309b565b7f0429a63b000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b805115614159576fffffffffffffffffffffffffffffffff6040820151166fffffffffffffffffffffffffffffffff602083015116106140f65750565b606490614157604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff604082015116158015906141e1575b6141805750565b606490614157604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff6020820151161515614179565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1991614339606092805461423d63ffffffff8260801c1642613f56565b9081614378575b50506fffffffffffffffffffffffffffffffff600181602086015116928281541680851060001461437057508280855b16167fffffffffffffffffffffffffffffffff000000000000000000000000000000008254161781556142ed8651151582907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b60408601517fffffffffffffffffffffffffffffffff0000000000000000000000000000000060809190911b16939092166fffffffffffffffffffffffffffffffff1692909217910155565b613f2560405180926fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b838091614274565b6fffffffffffffffffffffffffffffffff916143ad8392836143a66001880154948286169560801c906138de565b9116613fe8565b8082101561442c57505b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff9290911692909216167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116174260801b73ffffffff00000000000000000000000000000000161781553880614244565b90506143b7565b906040519182815491828252602082019060005260206000209260005b818110614465575050612fe592500383612a63565b8454835260019485019487945060209093019201614450565b80548210156130ab5760005260206000200190600090565b6000818152600360205260409020548015614625577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116138f157600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116138f1578181036145b6575b5050506002548015614587577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161454481600261447e565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b61460d6145c76145d893600261447e565b90549060031b1c928392600261447e565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600052600360205260406000205538808061450b565b5050600090565b6000818152600660205260409020548015614625577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116138f157600554907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116138f15781810361471d575b5050506005548015614587577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016146da81600561447e565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600555600052600660205260006040812055600190565b61473f61472e6145d893600561447e565b90549060031b1c928392600561447e565b905560005260066020526040600020553880806146a1565b9060018201918160005282602052604060002054801515600014614882577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116138f1578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116138f15781810361484b575b50505080548015614587577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019061480c828261447e565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b61486b61485b6145d8938661447e565b90549060031b1c9283928661447e565b9055600052836020526040600020553880806147d4565b50505050600090565b806000526003602052604060002054156000146148e557600254680100000000000000008110156129e0576148cc6145d8826001859401600255600261447e565b9055600254906000526003602052604060002055600190565b50600090565b806000526006602052604060002054156000146148e557600554680100000000000000008110156129e05761492c6145d8826001859401600555600561447e565b9055600554906000526006602052604060002055600190565b600082815260018201602052604090205461462557805490680100000000000000008210156129e057826149836145d884600180960185558461447e565b905580549260005201602052604060002055600190565b9182549060ff8260a01c16158015614bd9575b614bd3576fffffffffffffffffffffffffffffffff821691600185019081546149f263ffffffff6fffffffffffffffffffffffffffffffff83169360801c1642613f56565b9081614b35575b5050848110614ae95750838310614a53575050614a286fffffffffffffffffffffffffffffffff928392613f56565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b5460801c91614a628185613f56565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101908082116138f157614ab0614ab59273ffffffffffffffffffffffffffffffffffffffff96613fe8565b613a13565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828573ffffffffffffffffffffffffffffffffffffffff927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611614ba957614b5092613fc29160801c906138de565b80841015614ba45750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff00000000000000000000000000000000161786559238806149f9565b614b5b565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b5082156149ad56fea164736f6c634300081a000a",
}

var BurnMintTokenPoolV2ABI = BurnMintTokenPoolV2MetaData.ABI

var BurnMintTokenPoolV2Bin = BurnMintTokenPoolV2MetaData.Bin

func DeployBurnMintTokenPoolV2(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, localTokenDecimals uint8, allowlist []common.Address, rmnProxy common.Address, router common.Address) (common.Address, *types.Transaction, *BurnMintTokenPoolV2, error) {
	parsed, err := BurnMintTokenPoolV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BurnMintTokenPoolV2Bin), backend, token, localTokenDecimals, allowlist, rmnProxy, router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnMintTokenPoolV2{address: address, abi: *parsed, BurnMintTokenPoolV2Caller: BurnMintTokenPoolV2Caller{contract: contract}, BurnMintTokenPoolV2Transactor: BurnMintTokenPoolV2Transactor{contract: contract}, BurnMintTokenPoolV2Filterer: BurnMintTokenPoolV2Filterer{contract: contract}}, nil
}

type BurnMintTokenPoolV2 struct {
	address common.Address
	abi     abi.ABI
	BurnMintTokenPoolV2Caller
	BurnMintTokenPoolV2Transactor
	BurnMintTokenPoolV2Filterer
}

type BurnMintTokenPoolV2Caller struct {
	contract *bind.BoundContract
}

type BurnMintTokenPoolV2Transactor struct {
	contract *bind.BoundContract
}

type BurnMintTokenPoolV2Filterer struct {
	contract *bind.BoundContract
}

type BurnMintTokenPoolV2Session struct {
	Contract     *BurnMintTokenPoolV2
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type BurnMintTokenPoolV2CallerSession struct {
	Contract *BurnMintTokenPoolV2Caller
	CallOpts bind.CallOpts
}

type BurnMintTokenPoolV2TransactorSession struct {
	Contract     *BurnMintTokenPoolV2Transactor
	TransactOpts bind.TransactOpts
}

type BurnMintTokenPoolV2Raw struct {
	Contract *BurnMintTokenPoolV2
}

type BurnMintTokenPoolV2CallerRaw struct {
	Contract *BurnMintTokenPoolV2Caller
}

type BurnMintTokenPoolV2TransactorRaw struct {
	Contract *BurnMintTokenPoolV2Transactor
}

func NewBurnMintTokenPoolV2(address common.Address, backend bind.ContractBackend) (*BurnMintTokenPoolV2, error) {
	abi, err := abi.JSON(strings.NewReader(BurnMintTokenPoolV2ABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindBurnMintTokenPoolV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolV2{address: address, abi: abi, BurnMintTokenPoolV2Caller: BurnMintTokenPoolV2Caller{contract: contract}, BurnMintTokenPoolV2Transactor: BurnMintTokenPoolV2Transactor{contract: contract}, BurnMintTokenPoolV2Filterer: BurnMintTokenPoolV2Filterer{contract: contract}}, nil
}

func NewBurnMintTokenPoolV2Caller(address common.Address, caller bind.ContractCaller) (*BurnMintTokenPoolV2Caller, error) {
	contract, err := bindBurnMintTokenPoolV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolV2Caller{contract: contract}, nil
}

func NewBurnMintTokenPoolV2Transactor(address common.Address, transactor bind.ContractTransactor) (*BurnMintTokenPoolV2Transactor, error) {
	contract, err := bindBurnMintTokenPoolV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolV2Transactor{contract: contract}, nil
}

func NewBurnMintTokenPoolV2Filterer(address common.Address, filterer bind.ContractFilterer) (*BurnMintTokenPoolV2Filterer, error) {
	contract, err := bindBurnMintTokenPoolV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolV2Filterer{contract: contract}, nil
}

func bindBurnMintTokenPoolV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BurnMintTokenPoolV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintTokenPoolV2.Contract.BurnMintTokenPoolV2Caller.contract.Call(opts, result, method, params...)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.BurnMintTokenPoolV2Transactor.contract.Transfer(opts)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.BurnMintTokenPoolV2Transactor.contract.Transact(opts, method, params...)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintTokenPoolV2.Contract.contract.Call(opts, result, method, params...)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.contract.Transfer(opts)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.contract.Transact(opts, method, params...)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Caller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BurnMintTokenPoolV2.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) GetAllowList() ([]common.Address, error) {
	return _BurnMintTokenPoolV2.Contract.GetAllowList(&_BurnMintTokenPoolV2.CallOpts)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2CallerSession) GetAllowList() ([]common.Address, error) {
	return _BurnMintTokenPoolV2.Contract.GetAllowList(&_BurnMintTokenPoolV2.CallOpts)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Caller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BurnMintTokenPoolV2.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) GetAllowListEnabled() (bool, error) {
	return _BurnMintTokenPoolV2.Contract.GetAllowListEnabled(&_BurnMintTokenPoolV2.CallOpts)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2CallerSession) GetAllowListEnabled() (bool, error) {
	return _BurnMintTokenPoolV2.Contract.GetAllowListEnabled(&_BurnMintTokenPoolV2.CallOpts)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Caller) GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _BurnMintTokenPoolV2.contract.Call(opts, &out, "getCurrentInboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnMintTokenPoolV2.Contract.GetCurrentInboundRateLimiterState(&_BurnMintTokenPoolV2.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2CallerSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnMintTokenPoolV2.Contract.GetCurrentInboundRateLimiterState(&_BurnMintTokenPoolV2.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Caller) GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _BurnMintTokenPoolV2.contract.Call(opts, &out, "getCurrentOutboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnMintTokenPoolV2.Contract.GetCurrentOutboundRateLimiterState(&_BurnMintTokenPoolV2.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2CallerSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnMintTokenPoolV2.Contract.GetCurrentOutboundRateLimiterState(&_BurnMintTokenPoolV2.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Caller) GetFee(opts *bind.CallOpts, arg0 uint64, arg1 ClientEVM2AnyMessage, arg2 uint16, arg3 []byte) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintTokenPoolV2.contract.Call(opts, &out, "getFee", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) GetFee(arg0 uint64, arg1 ClientEVM2AnyMessage, arg2 uint16, arg3 []byte) (*big.Int, error) {
	return _BurnMintTokenPoolV2.Contract.GetFee(&_BurnMintTokenPoolV2.CallOpts, arg0, arg1, arg2, arg3)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2CallerSession) GetFee(arg0 uint64, arg1 ClientEVM2AnyMessage, arg2 uint16, arg3 []byte) (*big.Int, error) {
	return _BurnMintTokenPoolV2.Contract.GetFee(&_BurnMintTokenPoolV2.CallOpts, arg0, arg1, arg2, arg3)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Caller) GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintTokenPoolV2.contract.Call(opts, &out, "getRateLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) GetRateLimitAdmin() (common.Address, error) {
	return _BurnMintTokenPoolV2.Contract.GetRateLimitAdmin(&_BurnMintTokenPoolV2.CallOpts)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2CallerSession) GetRateLimitAdmin() (common.Address, error) {
	return _BurnMintTokenPoolV2.Contract.GetRateLimitAdmin(&_BurnMintTokenPoolV2.CallOpts)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Caller) GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error) {
	var out []interface{}
	err := _BurnMintTokenPoolV2.contract.Call(opts, &out, "getRemotePools", remoteChainSelector)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _BurnMintTokenPoolV2.Contract.GetRemotePools(&_BurnMintTokenPoolV2.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2CallerSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _BurnMintTokenPoolV2.Contract.GetRemotePools(&_BurnMintTokenPoolV2.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Caller) GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _BurnMintTokenPoolV2.contract.Call(opts, &out, "getRemoteToken", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _BurnMintTokenPoolV2.Contract.GetRemoteToken(&_BurnMintTokenPoolV2.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2CallerSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _BurnMintTokenPoolV2.Contract.GetRemoteToken(&_BurnMintTokenPoolV2.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Caller) GetRequiredInboundCCVs(opts *bind.CallOpts, arg0 common.Address, sourceChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	var out []interface{}
	err := _BurnMintTokenPoolV2.contract.Call(opts, &out, "getRequiredInboundCCVs", arg0, sourceChainSelector, arg2, arg3, arg4)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) GetRequiredInboundCCVs(arg0 common.Address, sourceChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _BurnMintTokenPoolV2.Contract.GetRequiredInboundCCVs(&_BurnMintTokenPoolV2.CallOpts, arg0, sourceChainSelector, arg2, arg3, arg4)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2CallerSession) GetRequiredInboundCCVs(arg0 common.Address, sourceChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _BurnMintTokenPoolV2.Contract.GetRequiredInboundCCVs(&_BurnMintTokenPoolV2.CallOpts, arg0, sourceChainSelector, arg2, arg3, arg4)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Caller) GetRequiredOutboundCCVs(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	var out []interface{}
	err := _BurnMintTokenPoolV2.contract.Call(opts, &out, "getRequiredOutboundCCVs", arg0, destChainSelector, arg2, arg3, arg4)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) GetRequiredOutboundCCVs(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _BurnMintTokenPoolV2.Contract.GetRequiredOutboundCCVs(&_BurnMintTokenPoolV2.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2CallerSession) GetRequiredOutboundCCVs(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _BurnMintTokenPoolV2.Contract.GetRequiredOutboundCCVs(&_BurnMintTokenPoolV2.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Caller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintTokenPoolV2.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) GetRmnProxy() (common.Address, error) {
	return _BurnMintTokenPoolV2.Contract.GetRmnProxy(&_BurnMintTokenPoolV2.CallOpts)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2CallerSession) GetRmnProxy() (common.Address, error) {
	return _BurnMintTokenPoolV2.Contract.GetRmnProxy(&_BurnMintTokenPoolV2.CallOpts)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Caller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintTokenPoolV2.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) GetRouter() (common.Address, error) {
	return _BurnMintTokenPoolV2.Contract.GetRouter(&_BurnMintTokenPoolV2.CallOpts)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2CallerSession) GetRouter() (common.Address, error) {
	return _BurnMintTokenPoolV2.Contract.GetRouter(&_BurnMintTokenPoolV2.CallOpts)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Caller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _BurnMintTokenPoolV2.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) GetSupportedChains() ([]uint64, error) {
	return _BurnMintTokenPoolV2.Contract.GetSupportedChains(&_BurnMintTokenPoolV2.CallOpts)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2CallerSession) GetSupportedChains() ([]uint64, error) {
	return _BurnMintTokenPoolV2.Contract.GetSupportedChains(&_BurnMintTokenPoolV2.CallOpts)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Caller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintTokenPoolV2.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) GetToken() (common.Address, error) {
	return _BurnMintTokenPoolV2.Contract.GetToken(&_BurnMintTokenPoolV2.CallOpts)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2CallerSession) GetToken() (common.Address, error) {
	return _BurnMintTokenPoolV2.Contract.GetToken(&_BurnMintTokenPoolV2.CallOpts)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Caller) GetTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BurnMintTokenPoolV2.contract.Call(opts, &out, "getTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) GetTokenDecimals() (uint8, error) {
	return _BurnMintTokenPoolV2.Contract.GetTokenDecimals(&_BurnMintTokenPoolV2.CallOpts)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2CallerSession) GetTokenDecimals() (uint8, error) {
	return _BurnMintTokenPoolV2.Contract.GetTokenDecimals(&_BurnMintTokenPoolV2.CallOpts)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Caller) IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	var out []interface{}
	err := _BurnMintTokenPoolV2.contract.Call(opts, &out, "isRemotePool", remoteChainSelector, remotePoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _BurnMintTokenPoolV2.Contract.IsRemotePool(&_BurnMintTokenPoolV2.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2CallerSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _BurnMintTokenPoolV2.Contract.IsRemotePool(&_BurnMintTokenPoolV2.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Caller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _BurnMintTokenPoolV2.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _BurnMintTokenPoolV2.Contract.IsSupportedChain(&_BurnMintTokenPoolV2.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2CallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _BurnMintTokenPoolV2.Contract.IsSupportedChain(&_BurnMintTokenPoolV2.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Caller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _BurnMintTokenPoolV2.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) IsSupportedToken(token common.Address) (bool, error) {
	return _BurnMintTokenPoolV2.Contract.IsSupportedToken(&_BurnMintTokenPoolV2.CallOpts, token)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2CallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _BurnMintTokenPoolV2.Contract.IsSupportedToken(&_BurnMintTokenPoolV2.CallOpts, token)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintTokenPoolV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) Owner() (common.Address, error) {
	return _BurnMintTokenPoolV2.Contract.Owner(&_BurnMintTokenPoolV2.CallOpts)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2CallerSession) Owner() (common.Address, error) {
	return _BurnMintTokenPoolV2.Contract.Owner(&_BurnMintTokenPoolV2.CallOpts)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BurnMintTokenPoolV2.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintTokenPoolV2.Contract.SupportsInterface(&_BurnMintTokenPoolV2.CallOpts, interfaceId)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintTokenPoolV2.Contract.SupportsInterface(&_BurnMintTokenPoolV2.CallOpts, interfaceId)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Caller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintTokenPoolV2.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) TypeAndVersion() (string, error) {
	return _BurnMintTokenPoolV2.Contract.TypeAndVersion(&_BurnMintTokenPoolV2.CallOpts)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2CallerSession) TypeAndVersion() (string, error) {
	return _BurnMintTokenPoolV2.Contract.TypeAndVersion(&_BurnMintTokenPoolV2.CallOpts)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.contract.Transact(opts, "acceptOwnership")
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) AcceptOwnership() (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.AcceptOwnership(&_BurnMintTokenPoolV2.TransactOpts)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.AcceptOwnership(&_BurnMintTokenPoolV2.TransactOpts)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Transactor) AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.contract.Transact(opts, "addRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.AddRemotePool(&_BurnMintTokenPoolV2.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2TransactorSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.AddRemotePool(&_BurnMintTokenPoolV2.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Transactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.ApplyAllowListUpdates(&_BurnMintTokenPoolV2.TransactOpts, removes, adds)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2TransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.ApplyAllowListUpdates(&_BurnMintTokenPoolV2.TransactOpts, removes, adds)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Transactor) ApplyCCVConfigUpdates(opts *bind.TransactOpts, ccvConfigArgs []TokenPoolV2CCVConfigArg) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.contract.Transact(opts, "applyCCVConfigUpdates", ccvConfigArgs)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) ApplyCCVConfigUpdates(ccvConfigArgs []TokenPoolV2CCVConfigArg) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.ApplyCCVConfigUpdates(&_BurnMintTokenPoolV2.TransactOpts, ccvConfigArgs)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2TransactorSession) ApplyCCVConfigUpdates(ccvConfigArgs []TokenPoolV2CCVConfigArg) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.ApplyCCVConfigUpdates(&_BurnMintTokenPoolV2.TransactOpts, ccvConfigArgs)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Transactor) ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.contract.Transact(opts, "applyChainUpdates", remoteChainSelectorsToRemove, chainsToAdd)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.ApplyChainUpdates(&_BurnMintTokenPoolV2.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2TransactorSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.ApplyChainUpdates(&_BurnMintTokenPoolV2.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Transactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, arg1 []byte) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.contract.Transact(opts, "lockOrBurn", lockOrBurnIn, arg1)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1, arg1 []byte) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.LockOrBurn(&_BurnMintTokenPoolV2.TransactOpts, lockOrBurnIn, arg1)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2TransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1, arg1 []byte) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.LockOrBurn(&_BurnMintTokenPoolV2.TransactOpts, lockOrBurnIn, arg1)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Transactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.LockOrBurn0(&_BurnMintTokenPoolV2.TransactOpts, lockOrBurnIn)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2TransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.LockOrBurn0(&_BurnMintTokenPoolV2.TransactOpts, lockOrBurnIn)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Transactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.ReleaseOrMint(&_BurnMintTokenPoolV2.TransactOpts, releaseOrMintIn)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2TransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.ReleaseOrMint(&_BurnMintTokenPoolV2.TransactOpts, releaseOrMintIn)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Transactor) RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.contract.Transact(opts, "removeRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.RemoveRemotePool(&_BurnMintTokenPoolV2.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2TransactorSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.RemoveRemotePool(&_BurnMintTokenPoolV2.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Transactor) SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.contract.Transact(opts, "setChainRateLimiterConfig", remoteChainSelector, outboundConfig, inboundConfig)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.SetChainRateLimiterConfig(&_BurnMintTokenPoolV2.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2TransactorSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.SetChainRateLimiterConfig(&_BurnMintTokenPoolV2.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Transactor) SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.contract.Transact(opts, "setChainRateLimiterConfigs", remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.SetChainRateLimiterConfigs(&_BurnMintTokenPoolV2.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2TransactorSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.SetChainRateLimiterConfigs(&_BurnMintTokenPoolV2.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Transactor) SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.contract.Transact(opts, "setRateLimitAdmin", rateLimitAdmin)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.SetRateLimitAdmin(&_BurnMintTokenPoolV2.TransactOpts, rateLimitAdmin)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2TransactorSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.SetRateLimitAdmin(&_BurnMintTokenPoolV2.TransactOpts, rateLimitAdmin)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Transactor) SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.contract.Transact(opts, "setRouter", newRouter)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.SetRouter(&_BurnMintTokenPoolV2.TransactOpts, newRouter)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2TransactorSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.SetRouter(&_BurnMintTokenPoolV2.TransactOpts, newRouter)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Transactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.contract.Transact(opts, "transferOwnership", to)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Session) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.TransferOwnership(&_BurnMintTokenPoolV2.TransactOpts, to)
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2TransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPoolV2.Contract.TransferOwnership(&_BurnMintTokenPoolV2.TransactOpts, to)
}

type BurnMintTokenPoolV2AllowListAddIterator struct {
	Event *BurnMintTokenPoolV2AllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolV2AllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolV2AllowListAdd)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolV2AllowListAdd)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolV2AllowListAddIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolV2AllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolV2AllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) FilterAllowListAdd(opts *bind.FilterOpts) (*BurnMintTokenPoolV2AllowListAddIterator, error) {

	logs, sub, err := _BurnMintTokenPoolV2.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolV2AllowListAddIterator{contract: _BurnMintTokenPoolV2.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2AllowListAdd) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPoolV2.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolV2AllowListAdd)
				if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) ParseAllowListAdd(log types.Log) (*BurnMintTokenPoolV2AllowListAdd, error) {
	event := new(BurnMintTokenPoolV2AllowListAdd)
	if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolV2AllowListRemoveIterator struct {
	Event *BurnMintTokenPoolV2AllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolV2AllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolV2AllowListRemove)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolV2AllowListRemove)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolV2AllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolV2AllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolV2AllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) FilterAllowListRemove(opts *bind.FilterOpts) (*BurnMintTokenPoolV2AllowListRemoveIterator, error) {

	logs, sub, err := _BurnMintTokenPoolV2.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolV2AllowListRemoveIterator{contract: _BurnMintTokenPoolV2.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2AllowListRemove) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPoolV2.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolV2AllowListRemove)
				if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) ParseAllowListRemove(log types.Log) (*BurnMintTokenPoolV2AllowListRemove, error) {
	event := new(BurnMintTokenPoolV2AllowListRemove)
	if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolV2CCVConfigUpdatedIterator struct {
	Event *BurnMintTokenPoolV2CCVConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolV2CCVConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolV2CCVConfigUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolV2CCVConfigUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolV2CCVConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolV2CCVConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolV2CCVConfigUpdated struct {
	RemoteChainSelector uint64
	OutboundCCVs        []common.Address
	InboundCCVs         []common.Address
	Raw                 types.Log
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolV2CCVConfigUpdatedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPoolV2.contract.FilterLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolV2CCVConfigUpdatedIterator{contract: _BurnMintTokenPoolV2.contract, event: "CCVConfigUpdated", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2CCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPoolV2.contract.WatchLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolV2CCVConfigUpdated)
				if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) ParseCCVConfigUpdated(log types.Log) (*BurnMintTokenPoolV2CCVConfigUpdated, error) {
	event := new(BurnMintTokenPoolV2CCVConfigUpdated)
	if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolV2ChainAddedIterator struct {
	Event *BurnMintTokenPoolV2ChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolV2ChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolV2ChainAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolV2ChainAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolV2ChainAddedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolV2ChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolV2ChainAdded struct {
	RemoteChainSelector       uint64
	RemoteToken               []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) FilterChainAdded(opts *bind.FilterOpts) (*BurnMintTokenPoolV2ChainAddedIterator, error) {

	logs, sub, err := _BurnMintTokenPoolV2.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolV2ChainAddedIterator{contract: _BurnMintTokenPoolV2.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2ChainAdded) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPoolV2.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolV2ChainAdded)
				if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "ChainAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) ParseChainAdded(log types.Log) (*BurnMintTokenPoolV2ChainAdded, error) {
	event := new(BurnMintTokenPoolV2ChainAdded)
	if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolV2ChainConfiguredIterator struct {
	Event *BurnMintTokenPoolV2ChainConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolV2ChainConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolV2ChainConfigured)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolV2ChainConfigured)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolV2ChainConfiguredIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolV2ChainConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolV2ChainConfigured struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) FilterChainConfigured(opts *bind.FilterOpts) (*BurnMintTokenPoolV2ChainConfiguredIterator, error) {

	logs, sub, err := _BurnMintTokenPoolV2.contract.FilterLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolV2ChainConfiguredIterator{contract: _BurnMintTokenPoolV2.contract, event: "ChainConfigured", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2ChainConfigured) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPoolV2.contract.WatchLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolV2ChainConfigured)
				if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) ParseChainConfigured(log types.Log) (*BurnMintTokenPoolV2ChainConfigured, error) {
	event := new(BurnMintTokenPoolV2ChainConfigured)
	if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolV2ChainRemovedIterator struct {
	Event *BurnMintTokenPoolV2ChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolV2ChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolV2ChainRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolV2ChainRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolV2ChainRemovedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolV2ChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolV2ChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) FilterChainRemoved(opts *bind.FilterOpts) (*BurnMintTokenPoolV2ChainRemovedIterator, error) {

	logs, sub, err := _BurnMintTokenPoolV2.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolV2ChainRemovedIterator{contract: _BurnMintTokenPoolV2.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2ChainRemoved) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPoolV2.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolV2ChainRemoved)
				if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) ParseChainRemoved(log types.Log) (*BurnMintTokenPoolV2ChainRemoved, error) {
	event := new(BurnMintTokenPoolV2ChainRemoved)
	if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolV2ConfigChangedIterator struct {
	Event *BurnMintTokenPoolV2ConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolV2ConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolV2ConfigChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolV2ConfigChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolV2ConfigChangedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolV2ConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolV2ConfigChanged struct {
	Config RateLimiterConfig
	Raw    types.Log
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) FilterConfigChanged(opts *bind.FilterOpts) (*BurnMintTokenPoolV2ConfigChangedIterator, error) {

	logs, sub, err := _BurnMintTokenPoolV2.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolV2ConfigChangedIterator{contract: _BurnMintTokenPoolV2.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2ConfigChanged) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPoolV2.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolV2ConfigChanged)
				if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) ParseConfigChanged(log types.Log) (*BurnMintTokenPoolV2ConfigChanged, error) {
	event := new(BurnMintTokenPoolV2ConfigChanged)
	if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolV2InboundRateLimitConsumedIterator struct {
	Event *BurnMintTokenPoolV2InboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolV2InboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolV2InboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolV2InboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolV2InboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolV2InboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolV2InboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolV2InboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPoolV2.contract.FilterLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolV2InboundRateLimitConsumedIterator{contract: _BurnMintTokenPoolV2.contract, event: "InboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2InboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPoolV2.contract.WatchLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolV2InboundRateLimitConsumed)
				if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) ParseInboundRateLimitConsumed(log types.Log) (*BurnMintTokenPoolV2InboundRateLimitConsumed, error) {
	event := new(BurnMintTokenPoolV2InboundRateLimitConsumed)
	if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolV2LockedOrBurnedIterator struct {
	Event *BurnMintTokenPoolV2LockedOrBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolV2LockedOrBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolV2LockedOrBurned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolV2LockedOrBurned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolV2LockedOrBurnedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolV2LockedOrBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolV2LockedOrBurned struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolV2LockedOrBurnedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPoolV2.contract.FilterLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolV2LockedOrBurnedIterator{contract: _BurnMintTokenPoolV2.contract, event: "LockedOrBurned", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2LockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPoolV2.contract.WatchLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolV2LockedOrBurned)
				if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) ParseLockedOrBurned(log types.Log) (*BurnMintTokenPoolV2LockedOrBurned, error) {
	event := new(BurnMintTokenPoolV2LockedOrBurned)
	if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolV2OutboundRateLimitConsumedIterator struct {
	Event *BurnMintTokenPoolV2OutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolV2OutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolV2OutboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolV2OutboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolV2OutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolV2OutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolV2OutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolV2OutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPoolV2.contract.FilterLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolV2OutboundRateLimitConsumedIterator{contract: _BurnMintTokenPoolV2.contract, event: "OutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2OutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPoolV2.contract.WatchLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolV2OutboundRateLimitConsumed)
				if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) ParseOutboundRateLimitConsumed(log types.Log) (*BurnMintTokenPoolV2OutboundRateLimitConsumed, error) {
	event := new(BurnMintTokenPoolV2OutboundRateLimitConsumed)
	if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolV2OwnershipTransferRequestedIterator struct {
	Event *BurnMintTokenPoolV2OwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolV2OwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolV2OwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolV2OwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolV2OwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolV2OwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolV2OwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintTokenPoolV2OwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintTokenPoolV2.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolV2OwnershipTransferRequestedIterator{contract: _BurnMintTokenPoolV2.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintTokenPoolV2.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolV2OwnershipTransferRequested)
				if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) ParseOwnershipTransferRequested(log types.Log) (*BurnMintTokenPoolV2OwnershipTransferRequested, error) {
	event := new(BurnMintTokenPoolV2OwnershipTransferRequested)
	if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolV2OwnershipTransferredIterator struct {
	Event *BurnMintTokenPoolV2OwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolV2OwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolV2OwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolV2OwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolV2OwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintTokenPoolV2OwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintTokenPoolV2.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolV2OwnershipTransferredIterator{contract: _BurnMintTokenPoolV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintTokenPoolV2.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolV2OwnershipTransferred)
				if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) ParseOwnershipTransferred(log types.Log) (*BurnMintTokenPoolV2OwnershipTransferred, error) {
	event := new(BurnMintTokenPoolV2OwnershipTransferred)
	if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolV2RateLimitAdminSetIterator struct {
	Event *BurnMintTokenPoolV2RateLimitAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolV2RateLimitAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolV2RateLimitAdminSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolV2RateLimitAdminSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolV2RateLimitAdminSetIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolV2RateLimitAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolV2RateLimitAdminSet struct {
	RateLimitAdmin common.Address
	Raw            types.Log
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) FilterRateLimitAdminSet(opts *bind.FilterOpts) (*BurnMintTokenPoolV2RateLimitAdminSetIterator, error) {

	logs, sub, err := _BurnMintTokenPoolV2.contract.FilterLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolV2RateLimitAdminSetIterator{contract: _BurnMintTokenPoolV2.contract, event: "RateLimitAdminSet", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2RateLimitAdminSet) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPoolV2.contract.WatchLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolV2RateLimitAdminSet)
				if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) ParseRateLimitAdminSet(log types.Log) (*BurnMintTokenPoolV2RateLimitAdminSet, error) {
	event := new(BurnMintTokenPoolV2RateLimitAdminSet)
	if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolV2ReleasedOrMintedIterator struct {
	Event *BurnMintTokenPoolV2ReleasedOrMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolV2ReleasedOrMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolV2ReleasedOrMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolV2ReleasedOrMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolV2ReleasedOrMintedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolV2ReleasedOrMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolV2ReleasedOrMinted struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Recipient           common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolV2ReleasedOrMintedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPoolV2.contract.FilterLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolV2ReleasedOrMintedIterator{contract: _BurnMintTokenPoolV2.contract, event: "ReleasedOrMinted", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2ReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPoolV2.contract.WatchLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolV2ReleasedOrMinted)
				if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) ParseReleasedOrMinted(log types.Log) (*BurnMintTokenPoolV2ReleasedOrMinted, error) {
	event := new(BurnMintTokenPoolV2ReleasedOrMinted)
	if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolV2RemotePoolAddedIterator struct {
	Event *BurnMintTokenPoolV2RemotePoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolV2RemotePoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolV2RemotePoolAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolV2RemotePoolAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolV2RemotePoolAddedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolV2RemotePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolV2RemotePoolAdded struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolV2RemotePoolAddedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPoolV2.contract.FilterLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolV2RemotePoolAddedIterator{contract: _BurnMintTokenPoolV2.contract, event: "RemotePoolAdded", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2RemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPoolV2.contract.WatchLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolV2RemotePoolAdded)
				if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) ParseRemotePoolAdded(log types.Log) (*BurnMintTokenPoolV2RemotePoolAdded, error) {
	event := new(BurnMintTokenPoolV2RemotePoolAdded)
	if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolV2RemotePoolRemovedIterator struct {
	Event *BurnMintTokenPoolV2RemotePoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolV2RemotePoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolV2RemotePoolRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolV2RemotePoolRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolV2RemotePoolRemovedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolV2RemotePoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolV2RemotePoolRemoved struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolV2RemotePoolRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPoolV2.contract.FilterLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolV2RemotePoolRemovedIterator{contract: _BurnMintTokenPoolV2.contract, event: "RemotePoolRemoved", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2RemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPoolV2.contract.WatchLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolV2RemotePoolRemoved)
				if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) ParseRemotePoolRemoved(log types.Log) (*BurnMintTokenPoolV2RemotePoolRemoved, error) {
	event := new(BurnMintTokenPoolV2RemotePoolRemoved)
	if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolV2RouterUpdatedIterator struct {
	Event *BurnMintTokenPoolV2RouterUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolV2RouterUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolV2RouterUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolV2RouterUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolV2RouterUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolV2RouterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolV2RouterUpdated struct {
	OldRouter common.Address
	NewRouter common.Address
	Raw       types.Log
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) FilterRouterUpdated(opts *bind.FilterOpts) (*BurnMintTokenPoolV2RouterUpdatedIterator, error) {

	logs, sub, err := _BurnMintTokenPoolV2.contract.FilterLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolV2RouterUpdatedIterator{contract: _BurnMintTokenPoolV2.contract, event: "RouterUpdated", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2RouterUpdated) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPoolV2.contract.WatchLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolV2RouterUpdated)
				if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2Filterer) ParseRouterUpdated(log types.Log) (*BurnMintTokenPoolV2RouterUpdated, error) {
	event := new(BurnMintTokenPoolV2RouterUpdated)
	if err := _BurnMintTokenPoolV2.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (BurnMintTokenPoolV2AllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (BurnMintTokenPoolV2AllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (BurnMintTokenPoolV2CCVConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xb0897119e8510f887b892cbc4c8506fc51d9849fd90afae4fd065e705f2d0f6c")
}

func (BurnMintTokenPoolV2ChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (BurnMintTokenPoolV2ChainConfigured) Topic() common.Hash {
	return common.HexToHash("0x0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b")
}

func (BurnMintTokenPoolV2ChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (BurnMintTokenPoolV2ConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (BurnMintTokenPoolV2InboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (BurnMintTokenPoolV2LockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
}

func (BurnMintTokenPoolV2OutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944")
}

func (BurnMintTokenPoolV2OwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (BurnMintTokenPoolV2OwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (BurnMintTokenPoolV2RateLimitAdminSet) Topic() common.Hash {
	return common.HexToHash("0x44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174")
}

func (BurnMintTokenPoolV2ReleasedOrMinted) Topic() common.Hash {
	return common.HexToHash("0xfc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0")
}

func (BurnMintTokenPoolV2RemotePoolAdded) Topic() common.Hash {
	return common.HexToHash("0x7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea")
}

func (BurnMintTokenPoolV2RemotePoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76")
}

func (BurnMintTokenPoolV2RouterUpdated) Topic() common.Hash {
	return common.HexToHash("0x02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f1684")
}

func (_BurnMintTokenPoolV2 *BurnMintTokenPoolV2) Address() common.Address {
	return _BurnMintTokenPoolV2.address
}

type BurnMintTokenPoolV2Interface interface {
	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetFee(opts *bind.CallOpts, arg0 uint64, arg1 ClientEVM2AnyMessage, arg2 uint16, arg3 []byte) (*big.Int, error)

	GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRequiredInboundCCVs(opts *bind.CallOpts, arg0 common.Address, sourceChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error)

	GetRequiredOutboundCCVs(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	GetSupportedChains(opts *bind.CallOpts) ([]uint64, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	GetTokenDecimals(opts *bind.CallOpts) (uint8, error)

	IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error)

	IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error)

	IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error)

	ApplyCCVConfigUpdates(opts *bind.TransactOpts, ccvConfigArgs []TokenPoolV2CCVConfigArg) (*types.Transaction, error)

	ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, arg1 []byte) (*types.Transaction, error)

	LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error)

	SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error)

	SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterAllowListAdd(opts *bind.FilterOpts) (*BurnMintTokenPoolV2AllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2AllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*BurnMintTokenPoolV2AllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*BurnMintTokenPoolV2AllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2AllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*BurnMintTokenPoolV2AllowListRemove, error)

	FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolV2CCVConfigUpdatedIterator, error)

	WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2CCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCCVConfigUpdated(log types.Log) (*BurnMintTokenPoolV2CCVConfigUpdated, error)

	FilterChainAdded(opts *bind.FilterOpts) (*BurnMintTokenPoolV2ChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2ChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*BurnMintTokenPoolV2ChainAdded, error)

	FilterChainConfigured(opts *bind.FilterOpts) (*BurnMintTokenPoolV2ChainConfiguredIterator, error)

	WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2ChainConfigured) (event.Subscription, error)

	ParseChainConfigured(log types.Log) (*BurnMintTokenPoolV2ChainConfigured, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*BurnMintTokenPoolV2ChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2ChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*BurnMintTokenPoolV2ChainRemoved, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*BurnMintTokenPoolV2ConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2ConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*BurnMintTokenPoolV2ConfigChanged, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolV2InboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2InboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*BurnMintTokenPoolV2InboundRateLimitConsumed, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolV2LockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2LockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*BurnMintTokenPoolV2LockedOrBurned, error)

	FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolV2OutboundRateLimitConsumedIterator, error)

	WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2OutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseOutboundRateLimitConsumed(log types.Log) (*BurnMintTokenPoolV2OutboundRateLimitConsumed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintTokenPoolV2OwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*BurnMintTokenPoolV2OwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintTokenPoolV2OwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*BurnMintTokenPoolV2OwnershipTransferred, error)

	FilterRateLimitAdminSet(opts *bind.FilterOpts) (*BurnMintTokenPoolV2RateLimitAdminSetIterator, error)

	WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2RateLimitAdminSet) (event.Subscription, error)

	ParseRateLimitAdminSet(log types.Log) (*BurnMintTokenPoolV2RateLimitAdminSet, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolV2ReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2ReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*BurnMintTokenPoolV2ReleasedOrMinted, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolV2RemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2RemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*BurnMintTokenPoolV2RemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolV2RemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2RemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*BurnMintTokenPoolV2RemotePoolRemoved, error)

	FilterRouterUpdated(opts *bind.FilterOpts) (*BurnMintTokenPoolV2RouterUpdatedIterator, error)

	WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolV2RouterUpdated) (event.Subscription, error)

	ParseRouterUpdated(log types.Log) (*BurnMintTokenPoolV2RouterUpdated, error)

	Address() common.Address
}
