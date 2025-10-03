// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token_pool_v2

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

var TokenPoolV2MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPoolV2.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredInboundCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredOutboundCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRouter\",\"inputs\":[{\"name\":\"newRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RouterUpdated\",\"inputs\":[{\"name\":\"oldRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCV\",\"inputs\":[{\"name\":\"ccv\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
}

var TokenPoolV2ABI = TokenPoolV2MetaData.ABI

type TokenPoolV2 struct {
	address common.Address
	abi     abi.ABI
	TokenPoolV2Caller
	TokenPoolV2Transactor
	TokenPoolV2Filterer
}

type TokenPoolV2Caller struct {
	contract *bind.BoundContract
}

type TokenPoolV2Transactor struct {
	contract *bind.BoundContract
}

type TokenPoolV2Filterer struct {
	contract *bind.BoundContract
}

type TokenPoolV2Session struct {
	Contract     *TokenPoolV2
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type TokenPoolV2CallerSession struct {
	Contract *TokenPoolV2Caller
	CallOpts bind.CallOpts
}

type TokenPoolV2TransactorSession struct {
	Contract     *TokenPoolV2Transactor
	TransactOpts bind.TransactOpts
}

type TokenPoolV2Raw struct {
	Contract *TokenPoolV2
}

type TokenPoolV2CallerRaw struct {
	Contract *TokenPoolV2Caller
}

type TokenPoolV2TransactorRaw struct {
	Contract *TokenPoolV2Transactor
}

func NewTokenPoolV2(address common.Address, backend bind.ContractBackend) (*TokenPoolV2, error) {
	abi, err := abi.JSON(strings.NewReader(TokenPoolV2ABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindTokenPoolV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenPoolV2{address: address, abi: abi, TokenPoolV2Caller: TokenPoolV2Caller{contract: contract}, TokenPoolV2Transactor: TokenPoolV2Transactor{contract: contract}, TokenPoolV2Filterer: TokenPoolV2Filterer{contract: contract}}, nil
}

func NewTokenPoolV2Caller(address common.Address, caller bind.ContractCaller) (*TokenPoolV2Caller, error) {
	contract, err := bindTokenPoolV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenPoolV2Caller{contract: contract}, nil
}

func NewTokenPoolV2Transactor(address common.Address, transactor bind.ContractTransactor) (*TokenPoolV2Transactor, error) {
	contract, err := bindTokenPoolV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenPoolV2Transactor{contract: contract}, nil
}

func NewTokenPoolV2Filterer(address common.Address, filterer bind.ContractFilterer) (*TokenPoolV2Filterer, error) {
	contract, err := bindTokenPoolV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenPoolV2Filterer{contract: contract}, nil
}

func bindTokenPoolV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenPoolV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_TokenPoolV2 *TokenPoolV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenPoolV2.Contract.TokenPoolV2Caller.contract.Call(opts, result, method, params...)
}

func (_TokenPoolV2 *TokenPoolV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPoolV2.Contract.TokenPoolV2Transactor.contract.Transfer(opts)
}

func (_TokenPoolV2 *TokenPoolV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenPoolV2.Contract.TokenPoolV2Transactor.contract.Transact(opts, method, params...)
}

func (_TokenPoolV2 *TokenPoolV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenPoolV2.Contract.contract.Call(opts, result, method, params...)
}

func (_TokenPoolV2 *TokenPoolV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPoolV2.Contract.contract.Transfer(opts)
}

func (_TokenPoolV2 *TokenPoolV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenPoolV2.Contract.contract.Transact(opts, method, params...)
}

func (_TokenPoolV2 *TokenPoolV2Caller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TokenPoolV2.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_TokenPoolV2 *TokenPoolV2Session) GetAllowList() ([]common.Address, error) {
	return _TokenPoolV2.Contract.GetAllowList(&_TokenPoolV2.CallOpts)
}

func (_TokenPoolV2 *TokenPoolV2CallerSession) GetAllowList() ([]common.Address, error) {
	return _TokenPoolV2.Contract.GetAllowList(&_TokenPoolV2.CallOpts)
}

func (_TokenPoolV2 *TokenPoolV2Caller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TokenPoolV2.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_TokenPoolV2 *TokenPoolV2Session) GetAllowListEnabled() (bool, error) {
	return _TokenPoolV2.Contract.GetAllowListEnabled(&_TokenPoolV2.CallOpts)
}

func (_TokenPoolV2 *TokenPoolV2CallerSession) GetAllowListEnabled() (bool, error) {
	return _TokenPoolV2.Contract.GetAllowListEnabled(&_TokenPoolV2.CallOpts)
}

func (_TokenPoolV2 *TokenPoolV2Caller) GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _TokenPoolV2.contract.Call(opts, &out, "getCurrentInboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_TokenPoolV2 *TokenPoolV2Session) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _TokenPoolV2.Contract.GetCurrentInboundRateLimiterState(&_TokenPoolV2.CallOpts, remoteChainSelector)
}

func (_TokenPoolV2 *TokenPoolV2CallerSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _TokenPoolV2.Contract.GetCurrentInboundRateLimiterState(&_TokenPoolV2.CallOpts, remoteChainSelector)
}

func (_TokenPoolV2 *TokenPoolV2Caller) GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _TokenPoolV2.contract.Call(opts, &out, "getCurrentOutboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_TokenPoolV2 *TokenPoolV2Session) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _TokenPoolV2.Contract.GetCurrentOutboundRateLimiterState(&_TokenPoolV2.CallOpts, remoteChainSelector)
}

func (_TokenPoolV2 *TokenPoolV2CallerSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _TokenPoolV2.Contract.GetCurrentOutboundRateLimiterState(&_TokenPoolV2.CallOpts, remoteChainSelector)
}

func (_TokenPoolV2 *TokenPoolV2Caller) GetFee(opts *bind.CallOpts, arg0 uint64, arg1 ClientEVM2AnyMessage, arg2 uint16, arg3 []byte) (*big.Int, error) {
	var out []interface{}
	err := _TokenPoolV2.contract.Call(opts, &out, "getFee", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_TokenPoolV2 *TokenPoolV2Session) GetFee(arg0 uint64, arg1 ClientEVM2AnyMessage, arg2 uint16, arg3 []byte) (*big.Int, error) {
	return _TokenPoolV2.Contract.GetFee(&_TokenPoolV2.CallOpts, arg0, arg1, arg2, arg3)
}

func (_TokenPoolV2 *TokenPoolV2CallerSession) GetFee(arg0 uint64, arg1 ClientEVM2AnyMessage, arg2 uint16, arg3 []byte) (*big.Int, error) {
	return _TokenPoolV2.Contract.GetFee(&_TokenPoolV2.CallOpts, arg0, arg1, arg2, arg3)
}

func (_TokenPoolV2 *TokenPoolV2Caller) GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenPoolV2.contract.Call(opts, &out, "getRateLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_TokenPoolV2 *TokenPoolV2Session) GetRateLimitAdmin() (common.Address, error) {
	return _TokenPoolV2.Contract.GetRateLimitAdmin(&_TokenPoolV2.CallOpts)
}

func (_TokenPoolV2 *TokenPoolV2CallerSession) GetRateLimitAdmin() (common.Address, error) {
	return _TokenPoolV2.Contract.GetRateLimitAdmin(&_TokenPoolV2.CallOpts)
}

func (_TokenPoolV2 *TokenPoolV2Caller) GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error) {
	var out []interface{}
	err := _TokenPoolV2.contract.Call(opts, &out, "getRemotePools", remoteChainSelector)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_TokenPoolV2 *TokenPoolV2Session) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _TokenPoolV2.Contract.GetRemotePools(&_TokenPoolV2.CallOpts, remoteChainSelector)
}

func (_TokenPoolV2 *TokenPoolV2CallerSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _TokenPoolV2.Contract.GetRemotePools(&_TokenPoolV2.CallOpts, remoteChainSelector)
}

func (_TokenPoolV2 *TokenPoolV2Caller) GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _TokenPoolV2.contract.Call(opts, &out, "getRemoteToken", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_TokenPoolV2 *TokenPoolV2Session) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _TokenPoolV2.Contract.GetRemoteToken(&_TokenPoolV2.CallOpts, remoteChainSelector)
}

func (_TokenPoolV2 *TokenPoolV2CallerSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _TokenPoolV2.Contract.GetRemoteToken(&_TokenPoolV2.CallOpts, remoteChainSelector)
}

func (_TokenPoolV2 *TokenPoolV2Caller) GetRequiredInboundCCVs(opts *bind.CallOpts, arg0 common.Address, sourceChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	var out []interface{}
	err := _TokenPoolV2.contract.Call(opts, &out, "getRequiredInboundCCVs", arg0, sourceChainSelector, arg2, arg3, arg4)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_TokenPoolV2 *TokenPoolV2Session) GetRequiredInboundCCVs(arg0 common.Address, sourceChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _TokenPoolV2.Contract.GetRequiredInboundCCVs(&_TokenPoolV2.CallOpts, arg0, sourceChainSelector, arg2, arg3, arg4)
}

func (_TokenPoolV2 *TokenPoolV2CallerSession) GetRequiredInboundCCVs(arg0 common.Address, sourceChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _TokenPoolV2.Contract.GetRequiredInboundCCVs(&_TokenPoolV2.CallOpts, arg0, sourceChainSelector, arg2, arg3, arg4)
}

func (_TokenPoolV2 *TokenPoolV2Caller) GetRequiredOutboundCCVs(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	var out []interface{}
	err := _TokenPoolV2.contract.Call(opts, &out, "getRequiredOutboundCCVs", arg0, destChainSelector, arg2, arg3, arg4)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_TokenPoolV2 *TokenPoolV2Session) GetRequiredOutboundCCVs(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _TokenPoolV2.Contract.GetRequiredOutboundCCVs(&_TokenPoolV2.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
}

func (_TokenPoolV2 *TokenPoolV2CallerSession) GetRequiredOutboundCCVs(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _TokenPoolV2.Contract.GetRequiredOutboundCCVs(&_TokenPoolV2.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
}

func (_TokenPoolV2 *TokenPoolV2Caller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenPoolV2.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_TokenPoolV2 *TokenPoolV2Session) GetRmnProxy() (common.Address, error) {
	return _TokenPoolV2.Contract.GetRmnProxy(&_TokenPoolV2.CallOpts)
}

func (_TokenPoolV2 *TokenPoolV2CallerSession) GetRmnProxy() (common.Address, error) {
	return _TokenPoolV2.Contract.GetRmnProxy(&_TokenPoolV2.CallOpts)
}

func (_TokenPoolV2 *TokenPoolV2Caller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenPoolV2.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_TokenPoolV2 *TokenPoolV2Session) GetRouter() (common.Address, error) {
	return _TokenPoolV2.Contract.GetRouter(&_TokenPoolV2.CallOpts)
}

func (_TokenPoolV2 *TokenPoolV2CallerSession) GetRouter() (common.Address, error) {
	return _TokenPoolV2.Contract.GetRouter(&_TokenPoolV2.CallOpts)
}

func (_TokenPoolV2 *TokenPoolV2Caller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _TokenPoolV2.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_TokenPoolV2 *TokenPoolV2Session) GetSupportedChains() ([]uint64, error) {
	return _TokenPoolV2.Contract.GetSupportedChains(&_TokenPoolV2.CallOpts)
}

func (_TokenPoolV2 *TokenPoolV2CallerSession) GetSupportedChains() ([]uint64, error) {
	return _TokenPoolV2.Contract.GetSupportedChains(&_TokenPoolV2.CallOpts)
}

func (_TokenPoolV2 *TokenPoolV2Caller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenPoolV2.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_TokenPoolV2 *TokenPoolV2Session) GetToken() (common.Address, error) {
	return _TokenPoolV2.Contract.GetToken(&_TokenPoolV2.CallOpts)
}

func (_TokenPoolV2 *TokenPoolV2CallerSession) GetToken() (common.Address, error) {
	return _TokenPoolV2.Contract.GetToken(&_TokenPoolV2.CallOpts)
}

func (_TokenPoolV2 *TokenPoolV2Caller) GetTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TokenPoolV2.contract.Call(opts, &out, "getTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_TokenPoolV2 *TokenPoolV2Session) GetTokenDecimals() (uint8, error) {
	return _TokenPoolV2.Contract.GetTokenDecimals(&_TokenPoolV2.CallOpts)
}

func (_TokenPoolV2 *TokenPoolV2CallerSession) GetTokenDecimals() (uint8, error) {
	return _TokenPoolV2.Contract.GetTokenDecimals(&_TokenPoolV2.CallOpts)
}

func (_TokenPoolV2 *TokenPoolV2Caller) IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	var out []interface{}
	err := _TokenPoolV2.contract.Call(opts, &out, "isRemotePool", remoteChainSelector, remotePoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_TokenPoolV2 *TokenPoolV2Session) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _TokenPoolV2.Contract.IsRemotePool(&_TokenPoolV2.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_TokenPoolV2 *TokenPoolV2CallerSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _TokenPoolV2.Contract.IsRemotePool(&_TokenPoolV2.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_TokenPoolV2 *TokenPoolV2Caller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _TokenPoolV2.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_TokenPoolV2 *TokenPoolV2Session) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _TokenPoolV2.Contract.IsSupportedChain(&_TokenPoolV2.CallOpts, remoteChainSelector)
}

func (_TokenPoolV2 *TokenPoolV2CallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _TokenPoolV2.Contract.IsSupportedChain(&_TokenPoolV2.CallOpts, remoteChainSelector)
}

func (_TokenPoolV2 *TokenPoolV2Caller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _TokenPoolV2.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_TokenPoolV2 *TokenPoolV2Session) IsSupportedToken(token common.Address) (bool, error) {
	return _TokenPoolV2.Contract.IsSupportedToken(&_TokenPoolV2.CallOpts, token)
}

func (_TokenPoolV2 *TokenPoolV2CallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _TokenPoolV2.Contract.IsSupportedToken(&_TokenPoolV2.CallOpts, token)
}

func (_TokenPoolV2 *TokenPoolV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenPoolV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_TokenPoolV2 *TokenPoolV2Session) Owner() (common.Address, error) {
	return _TokenPoolV2.Contract.Owner(&_TokenPoolV2.CallOpts)
}

func (_TokenPoolV2 *TokenPoolV2CallerSession) Owner() (common.Address, error) {
	return _TokenPoolV2.Contract.Owner(&_TokenPoolV2.CallOpts)
}

func (_TokenPoolV2 *TokenPoolV2Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TokenPoolV2.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_TokenPoolV2 *TokenPoolV2Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TokenPoolV2.Contract.SupportsInterface(&_TokenPoolV2.CallOpts, interfaceId)
}

func (_TokenPoolV2 *TokenPoolV2CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TokenPoolV2.Contract.SupportsInterface(&_TokenPoolV2.CallOpts, interfaceId)
}

func (_TokenPoolV2 *TokenPoolV2Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPoolV2.contract.Transact(opts, "acceptOwnership")
}

func (_TokenPoolV2 *TokenPoolV2Session) AcceptOwnership() (*types.Transaction, error) {
	return _TokenPoolV2.Contract.AcceptOwnership(&_TokenPoolV2.TransactOpts)
}

func (_TokenPoolV2 *TokenPoolV2TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _TokenPoolV2.Contract.AcceptOwnership(&_TokenPoolV2.TransactOpts)
}

func (_TokenPoolV2 *TokenPoolV2Transactor) AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _TokenPoolV2.contract.Transact(opts, "addRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_TokenPoolV2 *TokenPoolV2Session) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _TokenPoolV2.Contract.AddRemotePool(&_TokenPoolV2.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_TokenPoolV2 *TokenPoolV2TransactorSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _TokenPoolV2.Contract.AddRemotePool(&_TokenPoolV2.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_TokenPoolV2 *TokenPoolV2Transactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _TokenPoolV2.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_TokenPoolV2 *TokenPoolV2Session) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _TokenPoolV2.Contract.ApplyAllowListUpdates(&_TokenPoolV2.TransactOpts, removes, adds)
}

func (_TokenPoolV2 *TokenPoolV2TransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _TokenPoolV2.Contract.ApplyAllowListUpdates(&_TokenPoolV2.TransactOpts, removes, adds)
}

func (_TokenPoolV2 *TokenPoolV2Transactor) ApplyCCVConfigUpdates(opts *bind.TransactOpts, ccvConfigArgs []TokenPoolV2CCVConfigArg) (*types.Transaction, error) {
	return _TokenPoolV2.contract.Transact(opts, "applyCCVConfigUpdates", ccvConfigArgs)
}

func (_TokenPoolV2 *TokenPoolV2Session) ApplyCCVConfigUpdates(ccvConfigArgs []TokenPoolV2CCVConfigArg) (*types.Transaction, error) {
	return _TokenPoolV2.Contract.ApplyCCVConfigUpdates(&_TokenPoolV2.TransactOpts, ccvConfigArgs)
}

func (_TokenPoolV2 *TokenPoolV2TransactorSession) ApplyCCVConfigUpdates(ccvConfigArgs []TokenPoolV2CCVConfigArg) (*types.Transaction, error) {
	return _TokenPoolV2.Contract.ApplyCCVConfigUpdates(&_TokenPoolV2.TransactOpts, ccvConfigArgs)
}

func (_TokenPoolV2 *TokenPoolV2Transactor) ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _TokenPoolV2.contract.Transact(opts, "applyChainUpdates", remoteChainSelectorsToRemove, chainsToAdd)
}

func (_TokenPoolV2 *TokenPoolV2Session) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _TokenPoolV2.Contract.ApplyChainUpdates(&_TokenPoolV2.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_TokenPoolV2 *TokenPoolV2TransactorSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _TokenPoolV2.Contract.ApplyChainUpdates(&_TokenPoolV2.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_TokenPoolV2 *TokenPoolV2Transactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, arg1 []byte) (*types.Transaction, error) {
	return _TokenPoolV2.contract.Transact(opts, "lockOrBurn", lockOrBurnIn, arg1)
}

func (_TokenPoolV2 *TokenPoolV2Session) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1, arg1 []byte) (*types.Transaction, error) {
	return _TokenPoolV2.Contract.LockOrBurn(&_TokenPoolV2.TransactOpts, lockOrBurnIn, arg1)
}

func (_TokenPoolV2 *TokenPoolV2TransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1, arg1 []byte) (*types.Transaction, error) {
	return _TokenPoolV2.Contract.LockOrBurn(&_TokenPoolV2.TransactOpts, lockOrBurnIn, arg1)
}

func (_TokenPoolV2 *TokenPoolV2Transactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _TokenPoolV2.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn)
}

func (_TokenPoolV2 *TokenPoolV2Session) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _TokenPoolV2.Contract.LockOrBurn0(&_TokenPoolV2.TransactOpts, lockOrBurnIn)
}

func (_TokenPoolV2 *TokenPoolV2TransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _TokenPoolV2.Contract.LockOrBurn0(&_TokenPoolV2.TransactOpts, lockOrBurnIn)
}

func (_TokenPoolV2 *TokenPoolV2Transactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _TokenPoolV2.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_TokenPoolV2 *TokenPoolV2Session) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _TokenPoolV2.Contract.ReleaseOrMint(&_TokenPoolV2.TransactOpts, releaseOrMintIn)
}

func (_TokenPoolV2 *TokenPoolV2TransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _TokenPoolV2.Contract.ReleaseOrMint(&_TokenPoolV2.TransactOpts, releaseOrMintIn)
}

func (_TokenPoolV2 *TokenPoolV2Transactor) RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _TokenPoolV2.contract.Transact(opts, "removeRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_TokenPoolV2 *TokenPoolV2Session) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _TokenPoolV2.Contract.RemoveRemotePool(&_TokenPoolV2.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_TokenPoolV2 *TokenPoolV2TransactorSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _TokenPoolV2.Contract.RemoveRemotePool(&_TokenPoolV2.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_TokenPoolV2 *TokenPoolV2Transactor) SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _TokenPoolV2.contract.Transact(opts, "setChainRateLimiterConfig", remoteChainSelector, outboundConfig, inboundConfig)
}

func (_TokenPoolV2 *TokenPoolV2Session) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _TokenPoolV2.Contract.SetChainRateLimiterConfig(&_TokenPoolV2.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_TokenPoolV2 *TokenPoolV2TransactorSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _TokenPoolV2.Contract.SetChainRateLimiterConfig(&_TokenPoolV2.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_TokenPoolV2 *TokenPoolV2Transactor) SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _TokenPoolV2.contract.Transact(opts, "setChainRateLimiterConfigs", remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_TokenPoolV2 *TokenPoolV2Session) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _TokenPoolV2.Contract.SetChainRateLimiterConfigs(&_TokenPoolV2.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_TokenPoolV2 *TokenPoolV2TransactorSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _TokenPoolV2.Contract.SetChainRateLimiterConfigs(&_TokenPoolV2.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_TokenPoolV2 *TokenPoolV2Transactor) SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _TokenPoolV2.contract.Transact(opts, "setRateLimitAdmin", rateLimitAdmin)
}

func (_TokenPoolV2 *TokenPoolV2Session) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _TokenPoolV2.Contract.SetRateLimitAdmin(&_TokenPoolV2.TransactOpts, rateLimitAdmin)
}

func (_TokenPoolV2 *TokenPoolV2TransactorSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _TokenPoolV2.Contract.SetRateLimitAdmin(&_TokenPoolV2.TransactOpts, rateLimitAdmin)
}

func (_TokenPoolV2 *TokenPoolV2Transactor) SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error) {
	return _TokenPoolV2.contract.Transact(opts, "setRouter", newRouter)
}

func (_TokenPoolV2 *TokenPoolV2Session) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _TokenPoolV2.Contract.SetRouter(&_TokenPoolV2.TransactOpts, newRouter)
}

func (_TokenPoolV2 *TokenPoolV2TransactorSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _TokenPoolV2.Contract.SetRouter(&_TokenPoolV2.TransactOpts, newRouter)
}

func (_TokenPoolV2 *TokenPoolV2Transactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _TokenPoolV2.contract.Transact(opts, "transferOwnership", to)
}

func (_TokenPoolV2 *TokenPoolV2Session) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _TokenPoolV2.Contract.TransferOwnership(&_TokenPoolV2.TransactOpts, to)
}

func (_TokenPoolV2 *TokenPoolV2TransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _TokenPoolV2.Contract.TransferOwnership(&_TokenPoolV2.TransactOpts, to)
}

type TokenPoolV2AllowListAddIterator struct {
	Event *TokenPoolV2AllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenPoolV2AllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPoolV2AllowListAdd)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenPoolV2AllowListAdd)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenPoolV2AllowListAddIterator) Error() error {
	return it.fail
}

func (it *TokenPoolV2AllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenPoolV2AllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_TokenPoolV2 *TokenPoolV2Filterer) FilterAllowListAdd(opts *bind.FilterOpts) (*TokenPoolV2AllowListAddIterator, error) {

	logs, sub, err := _TokenPoolV2.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &TokenPoolV2AllowListAddIterator{contract: _TokenPoolV2.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *TokenPoolV2AllowListAdd) (event.Subscription, error) {

	logs, sub, err := _TokenPoolV2.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenPoolV2AllowListAdd)
				if err := _TokenPoolV2.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) ParseAllowListAdd(log types.Log) (*TokenPoolV2AllowListAdd, error) {
	event := new(TokenPoolV2AllowListAdd)
	if err := _TokenPoolV2.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenPoolV2AllowListRemoveIterator struct {
	Event *TokenPoolV2AllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenPoolV2AllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPoolV2AllowListRemove)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenPoolV2AllowListRemove)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenPoolV2AllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *TokenPoolV2AllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenPoolV2AllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_TokenPoolV2 *TokenPoolV2Filterer) FilterAllowListRemove(opts *bind.FilterOpts) (*TokenPoolV2AllowListRemoveIterator, error) {

	logs, sub, err := _TokenPoolV2.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &TokenPoolV2AllowListRemoveIterator{contract: _TokenPoolV2.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *TokenPoolV2AllowListRemove) (event.Subscription, error) {

	logs, sub, err := _TokenPoolV2.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenPoolV2AllowListRemove)
				if err := _TokenPoolV2.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) ParseAllowListRemove(log types.Log) (*TokenPoolV2AllowListRemove, error) {
	event := new(TokenPoolV2AllowListRemove)
	if err := _TokenPoolV2.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenPoolV2CCVConfigUpdatedIterator struct {
	Event *TokenPoolV2CCVConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenPoolV2CCVConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPoolV2CCVConfigUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenPoolV2CCVConfigUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenPoolV2CCVConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *TokenPoolV2CCVConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenPoolV2CCVConfigUpdated struct {
	RemoteChainSelector uint64
	OutboundCCVs        []common.Address
	InboundCCVs         []common.Address
	Raw                 types.Log
}

func (_TokenPoolV2 *TokenPoolV2Filterer) FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*TokenPoolV2CCVConfigUpdatedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _TokenPoolV2.contract.FilterLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &TokenPoolV2CCVConfigUpdatedIterator{contract: _TokenPoolV2.contract, event: "CCVConfigUpdated", logs: logs, sub: sub}, nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *TokenPoolV2CCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _TokenPoolV2.contract.WatchLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenPoolV2CCVConfigUpdated)
				if err := _TokenPoolV2.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) ParseCCVConfigUpdated(log types.Log) (*TokenPoolV2CCVConfigUpdated, error) {
	event := new(TokenPoolV2CCVConfigUpdated)
	if err := _TokenPoolV2.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenPoolV2ChainAddedIterator struct {
	Event *TokenPoolV2ChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenPoolV2ChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPoolV2ChainAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenPoolV2ChainAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenPoolV2ChainAddedIterator) Error() error {
	return it.fail
}

func (it *TokenPoolV2ChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenPoolV2ChainAdded struct {
	RemoteChainSelector       uint64
	RemoteToken               []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_TokenPoolV2 *TokenPoolV2Filterer) FilterChainAdded(opts *bind.FilterOpts) (*TokenPoolV2ChainAddedIterator, error) {

	logs, sub, err := _TokenPoolV2.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &TokenPoolV2ChainAddedIterator{contract: _TokenPoolV2.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *TokenPoolV2ChainAdded) (event.Subscription, error) {

	logs, sub, err := _TokenPoolV2.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenPoolV2ChainAdded)
				if err := _TokenPoolV2.contract.UnpackLog(event, "ChainAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) ParseChainAdded(log types.Log) (*TokenPoolV2ChainAdded, error) {
	event := new(TokenPoolV2ChainAdded)
	if err := _TokenPoolV2.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenPoolV2ChainConfiguredIterator struct {
	Event *TokenPoolV2ChainConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenPoolV2ChainConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPoolV2ChainConfigured)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenPoolV2ChainConfigured)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenPoolV2ChainConfiguredIterator) Error() error {
	return it.fail
}

func (it *TokenPoolV2ChainConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenPoolV2ChainConfigured struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_TokenPoolV2 *TokenPoolV2Filterer) FilterChainConfigured(opts *bind.FilterOpts) (*TokenPoolV2ChainConfiguredIterator, error) {

	logs, sub, err := _TokenPoolV2.contract.FilterLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return &TokenPoolV2ChainConfiguredIterator{contract: _TokenPoolV2.contract, event: "ChainConfigured", logs: logs, sub: sub}, nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *TokenPoolV2ChainConfigured) (event.Subscription, error) {

	logs, sub, err := _TokenPoolV2.contract.WatchLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenPoolV2ChainConfigured)
				if err := _TokenPoolV2.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) ParseChainConfigured(log types.Log) (*TokenPoolV2ChainConfigured, error) {
	event := new(TokenPoolV2ChainConfigured)
	if err := _TokenPoolV2.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenPoolV2ChainRemovedIterator struct {
	Event *TokenPoolV2ChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenPoolV2ChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPoolV2ChainRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenPoolV2ChainRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenPoolV2ChainRemovedIterator) Error() error {
	return it.fail
}

func (it *TokenPoolV2ChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenPoolV2ChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_TokenPoolV2 *TokenPoolV2Filterer) FilterChainRemoved(opts *bind.FilterOpts) (*TokenPoolV2ChainRemovedIterator, error) {

	logs, sub, err := _TokenPoolV2.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &TokenPoolV2ChainRemovedIterator{contract: _TokenPoolV2.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *TokenPoolV2ChainRemoved) (event.Subscription, error) {

	logs, sub, err := _TokenPoolV2.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenPoolV2ChainRemoved)
				if err := _TokenPoolV2.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) ParseChainRemoved(log types.Log) (*TokenPoolV2ChainRemoved, error) {
	event := new(TokenPoolV2ChainRemoved)
	if err := _TokenPoolV2.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenPoolV2ConfigChangedIterator struct {
	Event *TokenPoolV2ConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenPoolV2ConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPoolV2ConfigChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenPoolV2ConfigChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenPoolV2ConfigChangedIterator) Error() error {
	return it.fail
}

func (it *TokenPoolV2ConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenPoolV2ConfigChanged struct {
	Config RateLimiterConfig
	Raw    types.Log
}

func (_TokenPoolV2 *TokenPoolV2Filterer) FilterConfigChanged(opts *bind.FilterOpts) (*TokenPoolV2ConfigChangedIterator, error) {

	logs, sub, err := _TokenPoolV2.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &TokenPoolV2ConfigChangedIterator{contract: _TokenPoolV2.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *TokenPoolV2ConfigChanged) (event.Subscription, error) {

	logs, sub, err := _TokenPoolV2.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenPoolV2ConfigChanged)
				if err := _TokenPoolV2.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) ParseConfigChanged(log types.Log) (*TokenPoolV2ConfigChanged, error) {
	event := new(TokenPoolV2ConfigChanged)
	if err := _TokenPoolV2.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenPoolV2InboundRateLimitConsumedIterator struct {
	Event *TokenPoolV2InboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenPoolV2InboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPoolV2InboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenPoolV2InboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenPoolV2InboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *TokenPoolV2InboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenPoolV2InboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_TokenPoolV2 *TokenPoolV2Filterer) FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*TokenPoolV2InboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _TokenPoolV2.contract.FilterLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &TokenPoolV2InboundRateLimitConsumedIterator{contract: _TokenPoolV2.contract, event: "InboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *TokenPoolV2InboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _TokenPoolV2.contract.WatchLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenPoolV2InboundRateLimitConsumed)
				if err := _TokenPoolV2.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) ParseInboundRateLimitConsumed(log types.Log) (*TokenPoolV2InboundRateLimitConsumed, error) {
	event := new(TokenPoolV2InboundRateLimitConsumed)
	if err := _TokenPoolV2.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenPoolV2LockedOrBurnedIterator struct {
	Event *TokenPoolV2LockedOrBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenPoolV2LockedOrBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPoolV2LockedOrBurned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenPoolV2LockedOrBurned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenPoolV2LockedOrBurnedIterator) Error() error {
	return it.fail
}

func (it *TokenPoolV2LockedOrBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenPoolV2LockedOrBurned struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_TokenPoolV2 *TokenPoolV2Filterer) FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*TokenPoolV2LockedOrBurnedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _TokenPoolV2.contract.FilterLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &TokenPoolV2LockedOrBurnedIterator{contract: _TokenPoolV2.contract, event: "LockedOrBurned", logs: logs, sub: sub}, nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *TokenPoolV2LockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _TokenPoolV2.contract.WatchLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenPoolV2LockedOrBurned)
				if err := _TokenPoolV2.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) ParseLockedOrBurned(log types.Log) (*TokenPoolV2LockedOrBurned, error) {
	event := new(TokenPoolV2LockedOrBurned)
	if err := _TokenPoolV2.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenPoolV2OutboundRateLimitConsumedIterator struct {
	Event *TokenPoolV2OutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenPoolV2OutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPoolV2OutboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenPoolV2OutboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenPoolV2OutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *TokenPoolV2OutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenPoolV2OutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_TokenPoolV2 *TokenPoolV2Filterer) FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*TokenPoolV2OutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _TokenPoolV2.contract.FilterLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &TokenPoolV2OutboundRateLimitConsumedIterator{contract: _TokenPoolV2.contract, event: "OutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *TokenPoolV2OutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _TokenPoolV2.contract.WatchLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenPoolV2OutboundRateLimitConsumed)
				if err := _TokenPoolV2.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) ParseOutboundRateLimitConsumed(log types.Log) (*TokenPoolV2OutboundRateLimitConsumed, error) {
	event := new(TokenPoolV2OutboundRateLimitConsumed)
	if err := _TokenPoolV2.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenPoolV2OwnershipTransferRequestedIterator struct {
	Event *TokenPoolV2OwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenPoolV2OwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPoolV2OwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenPoolV2OwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenPoolV2OwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *TokenPoolV2OwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenPoolV2OwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_TokenPoolV2 *TokenPoolV2Filterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenPoolV2OwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenPoolV2.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenPoolV2OwnershipTransferRequestedIterator{contract: _TokenPoolV2.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *TokenPoolV2OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenPoolV2.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenPoolV2OwnershipTransferRequested)
				if err := _TokenPoolV2.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) ParseOwnershipTransferRequested(log types.Log) (*TokenPoolV2OwnershipTransferRequested, error) {
	event := new(TokenPoolV2OwnershipTransferRequested)
	if err := _TokenPoolV2.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenPoolV2OwnershipTransferredIterator struct {
	Event *TokenPoolV2OwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenPoolV2OwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPoolV2OwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenPoolV2OwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenPoolV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *TokenPoolV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenPoolV2OwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_TokenPoolV2 *TokenPoolV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenPoolV2OwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenPoolV2.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenPoolV2OwnershipTransferredIterator{contract: _TokenPoolV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenPoolV2OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenPoolV2.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenPoolV2OwnershipTransferred)
				if err := _TokenPoolV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) ParseOwnershipTransferred(log types.Log) (*TokenPoolV2OwnershipTransferred, error) {
	event := new(TokenPoolV2OwnershipTransferred)
	if err := _TokenPoolV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenPoolV2RateLimitAdminSetIterator struct {
	Event *TokenPoolV2RateLimitAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenPoolV2RateLimitAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPoolV2RateLimitAdminSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenPoolV2RateLimitAdminSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenPoolV2RateLimitAdminSetIterator) Error() error {
	return it.fail
}

func (it *TokenPoolV2RateLimitAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenPoolV2RateLimitAdminSet struct {
	RateLimitAdmin common.Address
	Raw            types.Log
}

func (_TokenPoolV2 *TokenPoolV2Filterer) FilterRateLimitAdminSet(opts *bind.FilterOpts) (*TokenPoolV2RateLimitAdminSetIterator, error) {

	logs, sub, err := _TokenPoolV2.contract.FilterLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return &TokenPoolV2RateLimitAdminSetIterator{contract: _TokenPoolV2.contract, event: "RateLimitAdminSet", logs: logs, sub: sub}, nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *TokenPoolV2RateLimitAdminSet) (event.Subscription, error) {

	logs, sub, err := _TokenPoolV2.contract.WatchLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenPoolV2RateLimitAdminSet)
				if err := _TokenPoolV2.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) ParseRateLimitAdminSet(log types.Log) (*TokenPoolV2RateLimitAdminSet, error) {
	event := new(TokenPoolV2RateLimitAdminSet)
	if err := _TokenPoolV2.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenPoolV2ReleasedOrMintedIterator struct {
	Event *TokenPoolV2ReleasedOrMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenPoolV2ReleasedOrMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPoolV2ReleasedOrMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenPoolV2ReleasedOrMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenPoolV2ReleasedOrMintedIterator) Error() error {
	return it.fail
}

func (it *TokenPoolV2ReleasedOrMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenPoolV2ReleasedOrMinted struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Recipient           common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_TokenPoolV2 *TokenPoolV2Filterer) FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*TokenPoolV2ReleasedOrMintedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _TokenPoolV2.contract.FilterLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &TokenPoolV2ReleasedOrMintedIterator{contract: _TokenPoolV2.contract, event: "ReleasedOrMinted", logs: logs, sub: sub}, nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *TokenPoolV2ReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _TokenPoolV2.contract.WatchLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenPoolV2ReleasedOrMinted)
				if err := _TokenPoolV2.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) ParseReleasedOrMinted(log types.Log) (*TokenPoolV2ReleasedOrMinted, error) {
	event := new(TokenPoolV2ReleasedOrMinted)
	if err := _TokenPoolV2.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenPoolV2RemotePoolAddedIterator struct {
	Event *TokenPoolV2RemotePoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenPoolV2RemotePoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPoolV2RemotePoolAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenPoolV2RemotePoolAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenPoolV2RemotePoolAddedIterator) Error() error {
	return it.fail
}

func (it *TokenPoolV2RemotePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenPoolV2RemotePoolAdded struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_TokenPoolV2 *TokenPoolV2Filterer) FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*TokenPoolV2RemotePoolAddedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _TokenPoolV2.contract.FilterLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &TokenPoolV2RemotePoolAddedIterator{contract: _TokenPoolV2.contract, event: "RemotePoolAdded", logs: logs, sub: sub}, nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *TokenPoolV2RemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _TokenPoolV2.contract.WatchLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenPoolV2RemotePoolAdded)
				if err := _TokenPoolV2.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) ParseRemotePoolAdded(log types.Log) (*TokenPoolV2RemotePoolAdded, error) {
	event := new(TokenPoolV2RemotePoolAdded)
	if err := _TokenPoolV2.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenPoolV2RemotePoolRemovedIterator struct {
	Event *TokenPoolV2RemotePoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenPoolV2RemotePoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPoolV2RemotePoolRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenPoolV2RemotePoolRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenPoolV2RemotePoolRemovedIterator) Error() error {
	return it.fail
}

func (it *TokenPoolV2RemotePoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenPoolV2RemotePoolRemoved struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_TokenPoolV2 *TokenPoolV2Filterer) FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*TokenPoolV2RemotePoolRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _TokenPoolV2.contract.FilterLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &TokenPoolV2RemotePoolRemovedIterator{contract: _TokenPoolV2.contract, event: "RemotePoolRemoved", logs: logs, sub: sub}, nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *TokenPoolV2RemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _TokenPoolV2.contract.WatchLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenPoolV2RemotePoolRemoved)
				if err := _TokenPoolV2.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) ParseRemotePoolRemoved(log types.Log) (*TokenPoolV2RemotePoolRemoved, error) {
	event := new(TokenPoolV2RemotePoolRemoved)
	if err := _TokenPoolV2.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenPoolV2RouterUpdatedIterator struct {
	Event *TokenPoolV2RouterUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenPoolV2RouterUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPoolV2RouterUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenPoolV2RouterUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenPoolV2RouterUpdatedIterator) Error() error {
	return it.fail
}

func (it *TokenPoolV2RouterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenPoolV2RouterUpdated struct {
	OldRouter common.Address
	NewRouter common.Address
	Raw       types.Log
}

func (_TokenPoolV2 *TokenPoolV2Filterer) FilterRouterUpdated(opts *bind.FilterOpts) (*TokenPoolV2RouterUpdatedIterator, error) {

	logs, sub, err := _TokenPoolV2.contract.FilterLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenPoolV2RouterUpdatedIterator{contract: _TokenPoolV2.contract, event: "RouterUpdated", logs: logs, sub: sub}, nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *TokenPoolV2RouterUpdated) (event.Subscription, error) {

	logs, sub, err := _TokenPoolV2.contract.WatchLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenPoolV2RouterUpdated)
				if err := _TokenPoolV2.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenPoolV2 *TokenPoolV2Filterer) ParseRouterUpdated(log types.Log) (*TokenPoolV2RouterUpdated, error) {
	event := new(TokenPoolV2RouterUpdated)
	if err := _TokenPoolV2.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (TokenPoolV2AllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (TokenPoolV2AllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (TokenPoolV2CCVConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xb0897119e8510f887b892cbc4c8506fc51d9849fd90afae4fd065e705f2d0f6c")
}

func (TokenPoolV2ChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (TokenPoolV2ChainConfigured) Topic() common.Hash {
	return common.HexToHash("0x0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b")
}

func (TokenPoolV2ChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (TokenPoolV2ConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (TokenPoolV2InboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (TokenPoolV2LockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
}

func (TokenPoolV2OutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944")
}

func (TokenPoolV2OwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (TokenPoolV2OwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (TokenPoolV2RateLimitAdminSet) Topic() common.Hash {
	return common.HexToHash("0x44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174")
}

func (TokenPoolV2ReleasedOrMinted) Topic() common.Hash {
	return common.HexToHash("0xfc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0")
}

func (TokenPoolV2RemotePoolAdded) Topic() common.Hash {
	return common.HexToHash("0x7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea")
}

func (TokenPoolV2RemotePoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76")
}

func (TokenPoolV2RouterUpdated) Topic() common.Hash {
	return common.HexToHash("0x02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f1684")
}

func (_TokenPoolV2 *TokenPoolV2) Address() common.Address {
	return _TokenPoolV2.address
}

type TokenPoolV2Interface interface {
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

	FilterAllowListAdd(opts *bind.FilterOpts) (*TokenPoolV2AllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *TokenPoolV2AllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*TokenPoolV2AllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*TokenPoolV2AllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *TokenPoolV2AllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*TokenPoolV2AllowListRemove, error)

	FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*TokenPoolV2CCVConfigUpdatedIterator, error)

	WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *TokenPoolV2CCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCCVConfigUpdated(log types.Log) (*TokenPoolV2CCVConfigUpdated, error)

	FilterChainAdded(opts *bind.FilterOpts) (*TokenPoolV2ChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *TokenPoolV2ChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*TokenPoolV2ChainAdded, error)

	FilterChainConfigured(opts *bind.FilterOpts) (*TokenPoolV2ChainConfiguredIterator, error)

	WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *TokenPoolV2ChainConfigured) (event.Subscription, error)

	ParseChainConfigured(log types.Log) (*TokenPoolV2ChainConfigured, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*TokenPoolV2ChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *TokenPoolV2ChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*TokenPoolV2ChainRemoved, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*TokenPoolV2ConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *TokenPoolV2ConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*TokenPoolV2ConfigChanged, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*TokenPoolV2InboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *TokenPoolV2InboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*TokenPoolV2InboundRateLimitConsumed, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*TokenPoolV2LockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *TokenPoolV2LockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*TokenPoolV2LockedOrBurned, error)

	FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*TokenPoolV2OutboundRateLimitConsumedIterator, error)

	WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *TokenPoolV2OutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseOutboundRateLimitConsumed(log types.Log) (*TokenPoolV2OutboundRateLimitConsumed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenPoolV2OwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *TokenPoolV2OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*TokenPoolV2OwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenPoolV2OwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenPoolV2OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*TokenPoolV2OwnershipTransferred, error)

	FilterRateLimitAdminSet(opts *bind.FilterOpts) (*TokenPoolV2RateLimitAdminSetIterator, error)

	WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *TokenPoolV2RateLimitAdminSet) (event.Subscription, error)

	ParseRateLimitAdminSet(log types.Log) (*TokenPoolV2RateLimitAdminSet, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*TokenPoolV2ReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *TokenPoolV2ReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*TokenPoolV2ReleasedOrMinted, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*TokenPoolV2RemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *TokenPoolV2RemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*TokenPoolV2RemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*TokenPoolV2RemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *TokenPoolV2RemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*TokenPoolV2RemotePoolRemoved, error)

	FilterRouterUpdated(opts *bind.FilterOpts) (*TokenPoolV2RouterUpdatedIterator, error)

	WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *TokenPoolV2RouterUpdated) (event.Subscription, error)

	ParseRouterUpdated(log types.Log) (*TokenPoolV2RouterUpdated, error)

	Address() common.Address
}
