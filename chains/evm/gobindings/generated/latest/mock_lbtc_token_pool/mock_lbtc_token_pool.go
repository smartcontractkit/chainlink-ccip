// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mock_lbtc_token_pool

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

var MockE2ELBTCTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIBurnMintERC20\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_destPoolData\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRouter\",\"inputs\":[{\"name\":\"newRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RouterUpdated\",\"inputs\":[{\"name\":\"oldRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenDataMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610100806040523461058c57614d3d803803809161001d82856105a9565b833981019060a08183031261058c5780516001600160a01b038116919082900361058c5760208101516001600160401b03811161058c5781019183601f8401121561058c578251926001600160401b038411610399578360051b602081019461008960405196876105a9565b85526020808601918301019186831161058c57602001905b828210610591575050506100b7604083016105cc565b6100c3606084016105cc565b608084015190936001600160401b03821161058c570185601f8201121561058c578051906001600160401b038211610399576040519661010d601f8401601f1916602001896105a9565b8288526020838301011161058c5760005b828110610577575050602060009187010152331561056657600180546001600160a01b0319163317905581158015610555575b8015610544575b610533578160209160049360805260c0526040519283809263313ce56760e01b82525afa80916000916104f0575b50906104cc575b50600860a052600480546001600160a01b0319166001600160a01b03929092169190911790558051151560e08190526103af575b5080516001600160401b03811161039957600a54600181811c9116801561038f575b602082101461037957601f8111610314575b50602091601f82116001146102b0579181926000926102a5575b50508160011b916000199060031b1c191617600a555b6040516145bc90816107818239608051818181611607015281816117f7015281816122c801528181612495015281816129090152612981015260a0518181816128900152818161345e01526134e1015260c051818181610be0015281816116a30152612364015260e051818181610b70015281816116e601526120440152f35b01519050388061020f565b601f19821692600a600052806000209160005b8581106102fc575083600195106102e3575b505050811b01600a55610225565b015160001960f88460031b161c191690553880806102d5565b919260206001819286850151815501940192016102c3565b600a6000527fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a8601f830160051c8101916020841061036f575b601f0160051c01905b81811061036357506101f5565b60008155600101610356565b909150819061034d565b634e487b7160e01b600052602260045260246000fd5b90607f16906101e3565b634e487b7160e01b600052604160045260246000fd5b60206040516103be82826105a9565b60008152600036813760e051156104bb5760005b8151811015610439576001906001600160a01b036103f082856105e0565b5116846103fc82610622565b610409575b5050016103d2565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a13884610401565b505060005b82518110156104b2576001906001600160a01b0361045c82866105e0565b511680156104ac578361046e82610720565b61047c575b50505b0161043e565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a13883610473565b50610476565b505050386101c1565b6335f4a7b360e01b60005260046000fd5b60ff166008811461018d576332ad3e0760e11b600052600860045260245260446000fd5b6020813d60201161052b575b81610509602093836105a9565b8101031261052757519060ff82168203610524575038610186565b80fd5b5080fd5b3d91506104fc565b6342bcdf7f60e11b60005260046000fd5b506001600160a01b03811615610158565b506001600160a01b03831615610151565b639b15e16f60e01b60005260046000fd5b80602080928401015182828b0101520161011e565b600080fd5b6020809161059e846105cc565b8152019101906100a1565b601f909101601f19168101906001600160401b0382119082101761039957604052565b51906001600160a01b038216820361058c57565b80518210156105f45760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b80548210156105f45760005260206000200190600090565b600081815260036020526040902054801561071957600019810181811161070357600254600019810191908211610703578181036106b2575b505050600254801561069c576000190161067681600261060a565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b6106eb6106c36106d493600261060a565b90549060031b1c928392600261060a565b819391549060031b91821b91600019901b19161790565b9055600052600360205260406000205538808061065b565b634e487b7160e01b600052601160045260246000fd5b5050600090565b8060005260036020526040600020541560001461077a5760025468010000000000000000811015610399576107616106d4826001859401600255600261060a565b9055600254906000526003602052604060002055600190565b5060009056fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a714612a2457508063181f5a77146129a557806321df0da714612936578063240028e8146128b457806324f65ee71461285857806332a7a8221461281e57806339077537146121f25780634c5ef0ed1461218d57806354c8a4f31461201057806362ddd3c414611f8c5780636d3d1a5814611f3a57806379ba509714611e555780637d54534e14611da85780638926f54f14611d445780638da5cb5b14611cf2578063962d402014611b4e5780639a4575b91461155f578063a42a7b8b146113da578063a7cd63b71461130e578063acfecf91146111ea578063af58d59f14611183578063b0f479a114611131578063b7946580146110da578063c0d7865514610fe2578063c4bffe2b14610e99578063c75eea9c14610dd3578063cf7401f314610c04578063dc0bd97114610b95578063e0351e1314610b3a578063e8a1da17146102655763f2fde38b1461017657600080fd5b346102625760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102625773ffffffffffffffffffffffffffffffffffffffff6101c2612c59565b6101ca613603565b1633811461023a57807fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b80fd5b50346102625761027436612f31565b93919092610280613603565b82915b8083106109a5575050508063ffffffff4216917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee1843603015b858210156109a1578160051b8501358181121561099d578501906101208236031261099d57604051956102ee87612b7a565b823567ffffffffffffffff81168103610998578752602083013567ffffffffffffffff81116109945783019536601f8801121561099457863596610331886131bd565b9761033f604051998a612bb2565b8089526020808a019160051b830101903682116109905760208301905b82821061095d575050505060208801968752604084013567ffffffffffffffff81116109595761038f9036908601612ee2565b9860408901998a526103b96103a73660608801613072565b9560608b0196875260c0369101613072565b9660808a019788526103cb8651613a7a565b6103d58851613a7a565b8a515115610931576103f167ffffffffffffffff8b51166142b9565b156108fa5767ffffffffffffffff8a5116815260076020526040812061053187516fffffffffffffffffffffffffffffffff604082015116906104ec6fffffffffffffffffffffffffffffffff6020830151169151151583608060405161045781612b7a565b858152602081018c905260408101849052606081018690520152855474ff000000000000000000000000000000000000000091151560a01b919091167fffffffffffffffffffffff0000000000000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff84161773ffffffff0000000000000000000000000000000060808b901b1617178555565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176001830155565b61065789516fffffffffffffffffffffffffffffffff604082015116906106126fffffffffffffffffffffffffffffffff6020830151169151151583608060405161057b81612b7a565b858152602081018c9052604081018490526060810186905201526002860180547fffffffffffffffffffffff000000000000000000000000000000000000000000166fffffffffffffffffffffffffffffffff85161773ffffffff0000000000000000000000000000000060808c901b161791151560a01b74ff000000000000000000000000000000000000000016919091179055565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176003830155565b60048c5191019080519067ffffffffffffffff82116108cd5761067a8354612c7c565b601f8111610892575b50602090601f83116001146107f3576106d192918591836107e8575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b805b8951805182101561070c57906107066001926106ff838f67ffffffffffffffff9051169261328c565b519061364e565b016106d6565b5050975097987f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2929593966107da67ffffffffffffffff600197949c51169251935191516107a661077160405196879687526101006020880152610100870190612c16565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a10190939492916102bc565b01519050388061069f565b83855281852091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416865b81811061087a5750908460019594939210610843575b505050811b0190556106d4565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388080610836565b92936020600181928786015181550195019301610820565b6108bd9084865260208620601f850160051c810191602086106108c3575b601f0160051c0190613392565b38610683565b90915081906108b0565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60249067ffffffffffffffff8b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b807f8579befe0000000000000000000000000000000000000000000000000000000060049252fd5b8680fd5b813567ffffffffffffffff811161098c576020916109818392833691890101612ee2565b81520191019061035c565b8a80fd5b8880fd5b8580fd5b600080fd5b8380fd5b8280f35b9092919367ffffffffffffffff6109c56109c087858861323d565b61316b565b16956109d087613fed565b15610b0e5786845260076020526109ec60056040862001613df4565b94845b8651811015610a25576001908987526007602052610a1e60056040892001610a17838b61328c565b5190614118565b50016109ef565b5093945094909580855260076020526005604086208681558660018201558660028201558660038201558660048201610a5e8154612c7c565b80610acd575b5050500180549086815581610aaf575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a1019190949394610283565b865260208620908101905b81811015610a7457868155600101610aba565b601f8111600114610ae35750555b863880610a64565b81835260208320610afe91601f01861c810190600101613392565b8082528160208120915555610adb565b602484887f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b503461026257807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102625760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b503461026257807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026257602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102625760e07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026257610c3c612e66565b9060607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc36011261026257604051610c7381612b96565b6024358015158103610dcf5781526044356fffffffffffffffffffffffffffffffff81168103610dcf5760208201526064356fffffffffffffffffffffffffffffffff81168103610dcf57604082015260607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c360112610dcb5760405190610cfa82612b96565b608435801515810361099d57825260a4356fffffffffffffffffffffffffffffffff8116810361099d57602083015260c4356fffffffffffffffffffffffffffffffff8116810361099d57604083015273ffffffffffffffffffffffffffffffffffffffff6009541633141580610da9575b610d7d57610d7a92936138b8565b80f35b6024837f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b5073ffffffffffffffffffffffffffffffffffffffff60015416331415610d6c565b5080fd5b8280fd5b50346102625760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026257610e3c610e376040610e959367ffffffffffffffff610e20612e66565b610e286132df565b5016815260076020522061330a565b6139f5565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b0390f35b503461026257807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026257604051906005548083528260208101600584526020842092845b818110610fc9575050610ef792500383612bb2565b8151610f1b610f05826131bd565b91610f136040519384612bb2565b8083526131bd565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602083019301368437805b8451811015610f7a578067ffffffffffffffff610f676001938861328c565b5116610f73828661328c565b5201610f48565b50925090604051928392602084019060208552518091526040840192915b818110610fa6575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101610f98565b8454835260019485019487945060209093019201610ee2565b50346102625760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102625773ffffffffffffffffffffffffffffffffffffffff61102f612c59565b611037613603565b1680156110b25760407f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f16849160045490807fffffffffffffffffffffffff000000000000000000000000000000000000000083161760045573ffffffffffffffffffffffffffffffffffffffff8351921682526020820152a180f35b6004827f8579befe000000000000000000000000000000000000000000000000000000008152fd5b50346102625760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026257610e9561111d611118612e66565b613370565b604051918291602083526020830190612c16565b503461026257807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026257602073ffffffffffffffffffffffffffffffffffffffff60045416604051908152f35b50346102625760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026257610e3c610e3760026040610e959467ffffffffffffffff6111d2612e66565b6111da6132df565b501681526007602052200161330a565b50346102625767ffffffffffffffff61120236612fa1565b92909161120d613603565b1691611226836000526006602052604060002054151590565b156112e257828452600760205261125560056040862001611248368486612e7d565b6020815191012090614118565b1561129a57907f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76916112946040519283926020845260208401916132a0565b0390a280f35b826112de836040519384937f74f23c7c00000000000000000000000000000000000000000000000000000000855260048501526040602485015260448401916132a0565b0390fd5b602484847f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b503461026257807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026257604051600254808252602082018091600285526020852090855b8181106113c4575050508261136d910383612bb2565b604051928392602084019060208552518091526040840192915b818110611395575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101611387565b8254845260209093019260019283019201611357565b50346102625760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102625767ffffffffffffffff61141b612e66565b168152600760205261143260056040832001613df4565b80517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611477611461836131bd565b9261146f6040519485612bb2565b8084526131bd565b01835b81811061154e575050825b82518110156114cb578061149b6001928561328c565b51855260086020526114af60408620612da6565b6114b9828561328c565b526114c4818461328c565b5001611485565b81846040519182916020830160208452825180915260408401602060408360051b870101940192905b82821061150357505050500390f35b9193602061153e827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851612c16565b96019201920185949391926114f4565b80606060208093860101520161147a565b50346102625760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102625760043567ffffffffffffffff8111610dcb5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8236030112610dcb57606060206040516115dd81612b5e565b8281520152608481016115ef8161314a565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611b045750602481019077ffffffffffffffff000000000000000000000000000000006116568361316b565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611a25578491611ad5575b50611aad576116e46044820161314a565b7f0000000000000000000000000000000000000000000000000000000000000000611a5b575b5067ffffffffffffffff61171d8361316b565b16611735816000526006602052604060002054151590565b15611a3057602073ffffffffffffffffffffffffffffffffffffffff60045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015611a255784906119c2575b73ffffffffffffffffffffffffffffffffffffffff915016330361199657819260646117ca67ffffffffffffffff9461316b565b92013592839216808252600760205261181f6040832073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016948591614368565b6040805173ffffffffffffffffffffffffffffffffffffffff85168152602081018690527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789449190a2813b15610262576040517f42966c68000000000000000000000000000000000000000000000000000000008152836004820152818160248183875af1801561198b57611976575b61194561190861111887877ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1060608967ffffffffffffffff6118ef8661316b565b16936040519182523360208301526040820152a261316b565b6040519061191582612b5e565b8152610e95611922612ccf565b602083019081526040519384936020855251604060208601526060850190612c16565b90517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0848303016040850152612c16565b611981828092612bb2565b61026257806118ae565b6040513d84823e3d90fd5b6024837f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d602011611a1d575b816119dc60209383612bb2565b8101031261099d575173ffffffffffffffffffffffffffffffffffffffff8116810361099d5773ffffffffffffffffffffffffffffffffffffffff90611796565b3d91506119cf565b6040513d86823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008452600452602483fd5b73ffffffffffffffffffffffffffffffffffffffff168084526003602052604084205461170a577fd0d25976000000000000000000000000000000000000000000000000000000008452600452602483fd5b6004837f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b611af7915060203d602011611afd575b611aef8183612bb2565b8101906135eb565b386116d3565b503d611ae5565b8273ffffffffffffffffffffffffffffffffffffffff611b2560249361314a565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835216600452fd5b50346102625760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102625760043567ffffffffffffffff8111610dcb57611b9e903690600401612f00565b60243567ffffffffffffffff811161099d57611bbe903690600401613024565b60449291923567ffffffffffffffff811161099457611be1903690600401613024565b91909273ffffffffffffffffffffffffffffffffffffffff6009541633141580611cd0575b611ca457818114801590611c9a575b611c7257865b818110611c26578780f35b80611c6c611c3a6109c0600194868c61323d565b611c4583878b61327c565b611c66611c5e611c56868b8d61327c565b923690613072565b913690613072565b916138b8565b01611c1b565b6004877f568efce2000000000000000000000000000000000000000000000000000000008152fd5b5082811415611c15565b6024877f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b5073ffffffffffffffffffffffffffffffffffffffff60015416331415611c06565b503461026257807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026257602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b50346102625760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610262576020611d9e67ffffffffffffffff611d8a612e66565b166000526006602052604060002054151590565b6040519015158152f35b50346102625760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610262577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174602073ffffffffffffffffffffffffffffffffffffffff611e18612c59565b611e20613603565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000006009541617600955604051908152a180f35b503461026257807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026257805473ffffffffffffffffffffffffffffffffffffffff81163303611f12577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516825573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b503461026257807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026257602073ffffffffffffffffffffffffffffffffffffffff60095416604051908152f35b503461026257611f9b36612fa1565b611fa793929193613603565b67ffffffffffffffff8216611fc9816000526006602052604060002054151590565b15611fe55750610d7a9293611fdf913691612e7d565b9061364e565b7f1e670e4b000000000000000000000000000000000000000000000000000000008452600452602483fd5b50346102625761203a9061204261202636612f31565b9591612033939193613603565b36916131d5565b9336916131d5565b7f00000000000000000000000000000000000000000000000000000000000000001561216557815b83518110156120dd578073ffffffffffffffffffffffffffffffffffffffff6120956001938761328c565b51166120a081613e57565b6120ac575b500161206a565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a1386120a5565b5090805b8251811015612161578073ffffffffffffffffffffffffffffffffffffffff61210c6001938661328c565b5116801561215b5761211d81614259565b61212a575b505b016120e1565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a184612122565b50612124565b5080f35b6004827f35f4a7b3000000000000000000000000000000000000000000000000000000008152fd5b50346102625760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610262576121c5612e66565b906024359067ffffffffffffffff8211610262576020611d9e846121ec3660048701612ee2565b90613180565b50346102625760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610262576004359067ffffffffffffffff82116102625781600401916101007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8236030112610dcb578160405161227481612b13565b5260c481016122a261229861229361228c84886130be565b3691612e7d565b6133eb565b60648401356134de565b91608481016122b08161314a565b73ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116036127fd5750602481019177ffffffffffffffff000000000000000000000000000000006123178461316b565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156126b25786916127de575b506127b65767ffffffffffffffff6123ab8461316b565b166123c3816000526006602052604060002054151590565b1561278b57602073ffffffffffffffffffffffffffffffffffffffff60045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa9081156126b257869161276c575b50156127405761243a8361316b565b61244f60a48401916121ec61228c848b6130be565b156126f957509483949567ffffffffffffffff61246b8561316b565b1680885260076020526124bd600260408a200173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016978891614368565b6040805173ffffffffffffffffffffffffffffffffffffffff88168152602081018990527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a28661251360e48501846130be565b81019190604081840312610dcb57803567ffffffffffffffff8111610dcf578361253e918301612ee2565b60208201359167ffffffffffffffff831161099d57602094612578936125649201612ee2565b508360405192828480945193849201612bf3565b8101039060025afa156126b25786519161259b61259583836130be565b9061310f565b83036126bd57505050604401906125b18261314a565b85843b15610262576040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff929092166004830152602482018690528160448183885af180156126b2579273ffffffffffffffffffffffffffffffffffffffff61267261266c60809560209a67ffffffffffffffff967ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0996126a2575b505061316b565b9261314a565b60405196875233898801521660408601528560608601521692a28060405161269981612b13565b52604051908152f35b816126ac91612bb2565b38612665565b6040513d88823e3d90fd5b916126ce61259589936044956130be565b7f7f249311000000000000000000000000000000000000000000000000000000008352600452602452fd5b61270390876130be565b6112de6040519283927f24eb47e50000000000000000000000000000000000000000000000000000000084526020600485015260248401916132a0565b6024857f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b612785915060203d602011611afd57611aef8183612bb2565b3861242b565b7fa9902c7e000000000000000000000000000000000000000000000000000000008652600452602485fd5b6004857f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b6127f7915060203d602011611afd57611aef8183612bb2565b38612394565b8473ffffffffffffffffffffffffffffffffffffffff611b2560249361314a565b503461026257807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026257610e9561111d612ccf565b503461026257807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026257602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102625760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610262576020906128ef612c59565b905073ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b503461026257807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261026257602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461026257807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102625750610e956040516129e6604082612bb2565b601a81527f4d6f636b4532454c425443546f6b656e506f6f6c20312e352e310000000000006020820152604051918291602083526020830190612c16565b905034610dcb5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610dcb576004357fffffffff000000000000000000000000000000000000000000000000000000008116809103610dcf57602092507faff2afbf000000000000000000000000000000000000000000000000000000008114908115612ae9575b8115612abf575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438612ab8565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150612ab1565b6020810190811067ffffffffffffffff821117612b2f57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff821117612b2f57604052565b60a0810190811067ffffffffffffffff821117612b2f57604052565b6060810190811067ffffffffffffffff821117612b2f57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117612b2f57604052565b60005b838110612c065750506000910152565b8181015183820152602001612bf6565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602093612c5281518092818752878088019101612bf3565b0116010190565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361099857565b90600182811c92168015612cc5575b6020831014612c9657565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691612c8b565b60405190600082600a5491612ce383612c7c565b8083529260018116908115612d695750600114612d09575b612d0792500383612bb2565b565b50600a600090815290917fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a85b818310612d4d575050906020612d0792820101612cfb565b6020919350806001915483858901015201910190918492612d35565b60209250612d079491507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001682840152151560051b820101612cfb565b9060405191826000825492612dba84612c7c565b8084529360018116908115612e265750600114612ddf575b50612d0792500383612bb2565b90506000929192526020600020906000915b818310612e0a575050906020612d079282010138612dd2565b6020919350806001915483858901015201910190918492612df1565b60209350612d079592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138612dd2565b6004359067ffffffffffffffff8216820361099857565b92919267ffffffffffffffff8211612b2f5760405191612ec5601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200184612bb2565b829481845281830111610998578281602093846000960137010152565b9080601f8301121561099857816020612efd93359101612e7d565b90565b9181601f840112156109985782359167ffffffffffffffff8311610998576020808501948460051b01011161099857565b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126109985760043567ffffffffffffffff81116109985781612f7a91600401612f00565b929092916024359067ffffffffffffffff821161099857612f9d91600401612f00565b9091565b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126109985760043567ffffffffffffffff81168103610998579160243567ffffffffffffffff811161099857826023820112156109985780600401359267ffffffffffffffff84116109985760248483010111610998576024019190565b9181601f840112156109985782359167ffffffffffffffff8311610998576020808501946060850201011161099857565b35906fffffffffffffffffffffffffffffffff8216820361099857565b91908260609103126109985760405161308a81612b96565b809280359081151582036109985760406130b991819385526130ae60208201613055565b602086015201613055565b910152565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610998570180359067ffffffffffffffff82116109985760200191813603831361099857565b35906020811061311d575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b3573ffffffffffffffffffffffffffffffffffffffff811681036109985790565b3567ffffffffffffffff811681036109985790565b9067ffffffffffffffff612efd92166000526007602052600560406000200190602081519101209060019160005201602052604060002054151590565b67ffffffffffffffff8111612b2f5760051b60200190565b92916131e0826131bd565b936131ee6040519586612bb2565b602085848152019260051b810191821161099857915b81831061321057505050565b823573ffffffffffffffffffffffffffffffffffffffff8116810361099857815260209283019201613204565b919081101561324d5760051b0190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b919081101561324d576060020190565b805182101561324d5760209160051b010190565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b604051906132ec82612b7a565b60006080838281528260208201528260408201528260608201520152565b9060405161331781612b7a565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff166000526007602052612efd6004604060002001612da6565b81811061339d575050565b60008155600101613392565b818102929181159184041417156133bc57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8051801561345a5760200361341c576020818051810103126109985760208101519060ff821161341c575060ff1690565b6112de906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190612c16565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff82116133bc57565b60ff16604d81116133bc57600a0a90565b81156134af570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff8116928284146135e4578284116135ba579061352391613480565b91604d60ff8416118015613581575b61354b57505090613545612efd92613494565b906133a9565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b5061358b83613494565b80156134af577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048411613532565b6135c391613480565b91604d60ff84161161354b575050906135de612efd92613494565b906134a5565b5050505090565b90816020910312610998575180151581036109985790565b73ffffffffffffffffffffffffffffffffffffffff60015416330361362457565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b9080511561388e5767ffffffffffffffff81516020830120921691826000526007602052613683816005604060002001614313565b1561384a5760005260086020526040600020815167ffffffffffffffff8111612b2f576136b08254612c7c565b601f8111613818575b506020601f8211600114613752579161372c827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea959361374295600091613747575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055604051918291602083526020830190612c16565b0390a2565b9050840151386136fb565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b8181106138005750926137429492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9896106137c9575b5050811b01905561111d565b8501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c1916905538806137bd565b9192602060018192868a015181550194019201613782565b61384490836000526020600020601f840160051c810191602085106108c357601f0160051c0190613392565b386136b9565b50906112de6040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190612c16565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b67ffffffffffffffff1660008181526006602052604090205490929190156139ba57916139b760e0926139838561390f7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b97613a7a565b846000526007602052613926816040600020613bc1565b61392f83613a7a565b846000526007602052613949836002604060002001613bc1565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b919082039182116133bc57565b6139fd6132df565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff8083511691613a5a6020850193613a54613a4763ffffffff875116426139e8565b85608089015116906133a9565b9061424c565b80821015613a7357505b16825263ffffffff4216905290565b9050613a64565b805115613b1a576fffffffffffffffffffffffffffffffff6040820151166fffffffffffffffffffffffffffffffff60208301511610613ab75750565b606490613b18604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff60408201511615801590613ba2575b613b415750565b606490613b18604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff6020820151161515613b3a565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1991613cfa6060928054613bfe63ffffffff8260801c16426139e8565b9081613d39575b50506fffffffffffffffffffffffffffffffff6001816020860151169282815416808510600014613d3157508280855b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416178155613cae8651151582907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b60408601517fffffffffffffffffffffffffffffffff0000000000000000000000000000000060809190911b16939092166fffffffffffffffffffffffffffffffff1692909217910155565b6139b760405180926fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b838091613c35565b6fffffffffffffffffffffffffffffffff91613d6e839283613d676001880154948286169560801c906133a9565b911661424c565b80821015613ded57505b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff9290911692909216167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116174260801b73ffffffff00000000000000000000000000000000161781553880613c05565b9050613d78565b906040519182815491828252602082019060005260206000209260005b818110613e26575050612d0792500383612bb2565b8454835260019485019487945060209093019201613e11565b805482101561324d5760005260206000200190600090565b6000818152600360205260409020548015613fe6577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116133bc57600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116133bc57818103613f77575b5050506002548015613f48577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01613f05816002613e3f565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b613fce613f88613f99936002613e3f565b90549060031b1c9283926002613e3f565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526003602052604060002055388080613ecc565b5050600090565b6000818152600660205260409020548015613fe6577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116133bc57600554907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116133bc578181036140de575b5050506005548015613f48577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161409b816005613e3f565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600555600052600660205260006040812055600190565b6141006140ef613f99936005613e3f565b90549060031b1c9283926005613e3f565b90556000526006602052604060002055388080614062565b9060018201918160005282602052604060002054801515600014614243577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116133bc578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116133bc5781810361420c575b50505080548015613f48577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906141cd8282613e3f565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b61422c61421c613f999386613e3f565b90549060031b1c92839286613e3f565b905560005283602052604060002055388080614195565b50505050600090565b919082018092116133bc57565b806000526003602052604060002054156000146142b35760025468010000000000000000811015612b2f5761429a613f998260018594016002556002613e3f565b9055600254906000526003602052604060002055600190565b50600090565b806000526006602052604060002054156000146142b35760055468010000000000000000811015612b2f576142fa613f998260018594016005556005613e3f565b9055600554906000526006602052604060002055600190565b6000828152600182016020526040902054613fe65780549068010000000000000000821015612b2f5782614351613f99846001809601855584613e3f565b905580549260005201602052604060002055600190565b9182549060ff8260a01c161580156145a7575b6145a1576fffffffffffffffffffffffffffffffff821691600185019081546143c063ffffffff6fffffffffffffffffffffffffffffffff83169360801c16426139e8565b9081614503575b50508481106144b757508383106144215750506143f66fffffffffffffffffffffffffffffffff9283926139e8565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b5460801c9161443081856139e8565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101908082116133bc5761447e6144839273ffffffffffffffffffffffffffffffffffffffff9661424c565b6134a5565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828573ffffffffffffffffffffffffffffffffffffffff927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b8286929396116145775761451e92613a549160801c906133a9565b808410156145725750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff00000000000000000000000000000000161786559238806143c7565b614529565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b50821561437b56fea164736f6c634300081a000a",
}

var MockE2ELBTCTokenPoolABI = MockE2ELBTCTokenPoolMetaData.ABI

var MockE2ELBTCTokenPoolBin = MockE2ELBTCTokenPoolMetaData.Bin

func DeployMockE2ELBTCTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, allowlist []common.Address, rmnProxy common.Address, router common.Address, destPoolData []byte) (common.Address, *types.Transaction, *MockE2ELBTCTokenPool, error) {
	parsed, err := MockE2ELBTCTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockE2ELBTCTokenPoolBin), backend, token, allowlist, rmnProxy, router, destPoolData)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockE2ELBTCTokenPool{address: address, abi: *parsed, MockE2ELBTCTokenPoolCaller: MockE2ELBTCTokenPoolCaller{contract: contract}, MockE2ELBTCTokenPoolTransactor: MockE2ELBTCTokenPoolTransactor{contract: contract}, MockE2ELBTCTokenPoolFilterer: MockE2ELBTCTokenPoolFilterer{contract: contract}}, nil
}

type MockE2ELBTCTokenPool struct {
	address common.Address
	abi     abi.ABI
	MockE2ELBTCTokenPoolCaller
	MockE2ELBTCTokenPoolTransactor
	MockE2ELBTCTokenPoolFilterer
}

type MockE2ELBTCTokenPoolCaller struct {
	contract *bind.BoundContract
}

type MockE2ELBTCTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type MockE2ELBTCTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type MockE2ELBTCTokenPoolSession struct {
	Contract     *MockE2ELBTCTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MockE2ELBTCTokenPoolCallerSession struct {
	Contract *MockE2ELBTCTokenPoolCaller
	CallOpts bind.CallOpts
}

type MockE2ELBTCTokenPoolTransactorSession struct {
	Contract     *MockE2ELBTCTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type MockE2ELBTCTokenPoolRaw struct {
	Contract *MockE2ELBTCTokenPool
}

type MockE2ELBTCTokenPoolCallerRaw struct {
	Contract *MockE2ELBTCTokenPoolCaller
}

type MockE2ELBTCTokenPoolTransactorRaw struct {
	Contract *MockE2ELBTCTokenPoolTransactor
}

func NewMockE2ELBTCTokenPool(address common.Address, backend bind.ContractBackend) (*MockE2ELBTCTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(MockE2ELBTCTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMockE2ELBTCTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPool{address: address, abi: abi, MockE2ELBTCTokenPoolCaller: MockE2ELBTCTokenPoolCaller{contract: contract}, MockE2ELBTCTokenPoolTransactor: MockE2ELBTCTokenPoolTransactor{contract: contract}, MockE2ELBTCTokenPoolFilterer: MockE2ELBTCTokenPoolFilterer{contract: contract}}, nil
}

func NewMockE2ELBTCTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*MockE2ELBTCTokenPoolCaller, error) {
	contract, err := bindMockE2ELBTCTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolCaller{contract: contract}, nil
}

func NewMockE2ELBTCTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*MockE2ELBTCTokenPoolTransactor, error) {
	contract, err := bindMockE2ELBTCTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolTransactor{contract: contract}, nil
}

func NewMockE2ELBTCTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*MockE2ELBTCTokenPoolFilterer, error) {
	contract, err := bindMockE2ELBTCTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolFilterer{contract: contract}, nil
}

func bindMockE2ELBTCTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockE2ELBTCTokenPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockE2ELBTCTokenPool.Contract.MockE2ELBTCTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.MockE2ELBTCTokenPoolTransactor.contract.Transfer(opts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.MockE2ELBTCTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockE2ELBTCTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.contract.Transfer(opts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetAllowList() ([]common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.GetAllowList(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetAllowList() ([]common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.GetAllowList(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetAllowListEnabled() (bool, error) {
	return _MockE2ELBTCTokenPool.Contract.GetAllowListEnabled(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetAllowListEnabled() (bool, error) {
	return _MockE2ELBTCTokenPool.Contract.GetAllowListEnabled(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getCurrentInboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _MockE2ELBTCTokenPool.Contract.GetCurrentInboundRateLimiterState(&_MockE2ELBTCTokenPool.CallOpts, remoteChainSelector)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _MockE2ELBTCTokenPool.Contract.GetCurrentInboundRateLimiterState(&_MockE2ELBTCTokenPool.CallOpts, remoteChainSelector)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getCurrentOutboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _MockE2ELBTCTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_MockE2ELBTCTokenPool.CallOpts, remoteChainSelector)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _MockE2ELBTCTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_MockE2ELBTCTokenPool.CallOpts, remoteChainSelector)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getRateLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetRateLimitAdmin() (common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.GetRateLimitAdmin(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetRateLimitAdmin() (common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.GetRateLimitAdmin(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getRemotePools", remoteChainSelector)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _MockE2ELBTCTokenPool.Contract.GetRemotePools(&_MockE2ELBTCTokenPool.CallOpts, remoteChainSelector)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _MockE2ELBTCTokenPool.Contract.GetRemotePools(&_MockE2ELBTCTokenPool.CallOpts, remoteChainSelector)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getRemoteToken", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _MockE2ELBTCTokenPool.Contract.GetRemoteToken(&_MockE2ELBTCTokenPool.CallOpts, remoteChainSelector)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _MockE2ELBTCTokenPool.Contract.GetRemoteToken(&_MockE2ELBTCTokenPool.CallOpts, remoteChainSelector)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetRmnProxy() (common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.GetRmnProxy(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetRmnProxy() (common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.GetRmnProxy(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetRouter() (common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.GetRouter(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetRouter() (common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.GetRouter(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetSupportedChains() ([]uint64, error) {
	return _MockE2ELBTCTokenPool.Contract.GetSupportedChains(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetSupportedChains() ([]uint64, error) {
	return _MockE2ELBTCTokenPool.Contract.GetSupportedChains(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetToken() (common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.GetToken(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.GetToken(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetTokenDecimals() (uint8, error) {
	return _MockE2ELBTCTokenPool.Contract.GetTokenDecimals(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetTokenDecimals() (uint8, error) {
	return _MockE2ELBTCTokenPool.Contract.GetTokenDecimals(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) IDestPoolData(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "i_destPoolData")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) IDestPoolData() ([]byte, error) {
	return _MockE2ELBTCTokenPool.Contract.IDestPoolData(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) IDestPoolData() ([]byte, error) {
	return _MockE2ELBTCTokenPool.Contract.IDestPoolData(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "isRemotePool", remoteChainSelector, remotePoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _MockE2ELBTCTokenPool.Contract.IsRemotePool(&_MockE2ELBTCTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _MockE2ELBTCTokenPool.Contract.IsRemotePool(&_MockE2ELBTCTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _MockE2ELBTCTokenPool.Contract.IsSupportedChain(&_MockE2ELBTCTokenPool.CallOpts, remoteChainSelector)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _MockE2ELBTCTokenPool.Contract.IsSupportedChain(&_MockE2ELBTCTokenPool.CallOpts, remoteChainSelector)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) IsSupportedToken(token common.Address) (bool, error) {
	return _MockE2ELBTCTokenPool.Contract.IsSupportedToken(&_MockE2ELBTCTokenPool.CallOpts, token)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _MockE2ELBTCTokenPool.Contract.IsSupportedToken(&_MockE2ELBTCTokenPool.CallOpts, token)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) Owner() (common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.Owner(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) Owner() (common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.Owner(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MockE2ELBTCTokenPool.Contract.SupportsInterface(&_MockE2ELBTCTokenPool.CallOpts, interfaceId)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MockE2ELBTCTokenPool.Contract.SupportsInterface(&_MockE2ELBTCTokenPool.CallOpts, interfaceId)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) TypeAndVersion() (string, error) {
	return _MockE2ELBTCTokenPool.Contract.TypeAndVersion(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) TypeAndVersion() (string, error) {
	return _MockE2ELBTCTokenPool.Contract.TypeAndVersion(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.AcceptOwnership(&_MockE2ELBTCTokenPool.TransactOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.AcceptOwnership(&_MockE2ELBTCTokenPool.TransactOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "addRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.AddRemotePool(&_MockE2ELBTCTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.AddRemotePool(&_MockE2ELBTCTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.ApplyAllowListUpdates(&_MockE2ELBTCTokenPool.TransactOpts, removes, adds)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.ApplyAllowListUpdates(&_MockE2ELBTCTokenPool.TransactOpts, removes, adds)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "applyChainUpdates", remoteChainSelectorsToRemove, chainsToAdd)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.ApplyChainUpdates(&_MockE2ELBTCTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.ApplyChainUpdates(&_MockE2ELBTCTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.LockOrBurn(&_MockE2ELBTCTokenPool.TransactOpts, lockOrBurnIn)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.LockOrBurn(&_MockE2ELBTCTokenPool.TransactOpts, lockOrBurnIn)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.ReleaseOrMint(&_MockE2ELBTCTokenPool.TransactOpts, releaseOrMintIn)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.ReleaseOrMint(&_MockE2ELBTCTokenPool.TransactOpts, releaseOrMintIn)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "removeRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.RemoveRemotePool(&_MockE2ELBTCTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.RemoveRemotePool(&_MockE2ELBTCTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "setChainRateLimiterConfig", remoteChainSelector, outboundConfig, inboundConfig)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.SetChainRateLimiterConfig(&_MockE2ELBTCTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.SetChainRateLimiterConfig(&_MockE2ELBTCTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "setChainRateLimiterConfigs", remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.SetChainRateLimiterConfigs(&_MockE2ELBTCTokenPool.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.SetChainRateLimiterConfigs(&_MockE2ELBTCTokenPool.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "setRateLimitAdmin", rateLimitAdmin)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.SetRateLimitAdmin(&_MockE2ELBTCTokenPool.TransactOpts, rateLimitAdmin)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.SetRateLimitAdmin(&_MockE2ELBTCTokenPool.TransactOpts, rateLimitAdmin)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "setRouter", newRouter)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.SetRouter(&_MockE2ELBTCTokenPool.TransactOpts, newRouter)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.SetRouter(&_MockE2ELBTCTokenPool.TransactOpts, newRouter)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.TransferOwnership(&_MockE2ELBTCTokenPool.TransactOpts, to)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.TransferOwnership(&_MockE2ELBTCTokenPool.TransactOpts, to)
}

type MockE2ELBTCTokenPoolAllowListAddIterator struct {
	Event *MockE2ELBTCTokenPoolAllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolAllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolAllowListAdd)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(MockE2ELBTCTokenPoolAllowListAdd)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *MockE2ELBTCTokenPoolAllowListAddIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolAllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolAllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterAllowListAdd(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolAllowListAddIterator, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolAllowListAddIterator{contract: _MockE2ELBTCTokenPool.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolAllowListAdd) (event.Subscription, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolAllowListAdd)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseAllowListAdd(log types.Log) (*MockE2ELBTCTokenPoolAllowListAdd, error) {
	event := new(MockE2ELBTCTokenPoolAllowListAdd)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolAllowListRemoveIterator struct {
	Event *MockE2ELBTCTokenPoolAllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolAllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolAllowListRemove)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(MockE2ELBTCTokenPoolAllowListRemove)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *MockE2ELBTCTokenPoolAllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolAllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolAllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterAllowListRemove(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolAllowListRemoveIterator, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolAllowListRemoveIterator{contract: _MockE2ELBTCTokenPool.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolAllowListRemove) (event.Subscription, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolAllowListRemove)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseAllowListRemove(log types.Log) (*MockE2ELBTCTokenPoolAllowListRemove, error) {
	event := new(MockE2ELBTCTokenPoolAllowListRemove)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolChainAddedIterator struct {
	Event *MockE2ELBTCTokenPoolChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolChainAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(MockE2ELBTCTokenPoolChainAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *MockE2ELBTCTokenPoolChainAddedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolChainAdded struct {
	RemoteChainSelector       uint64
	RemoteToken               []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterChainAdded(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolChainAddedIterator, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolChainAddedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolChainAdded) (event.Subscription, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolChainAdded)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseChainAdded(log types.Log) (*MockE2ELBTCTokenPoolChainAdded, error) {
	event := new(MockE2ELBTCTokenPoolChainAdded)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolChainConfiguredIterator struct {
	Event *MockE2ELBTCTokenPoolChainConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolChainConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolChainConfigured)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(MockE2ELBTCTokenPoolChainConfigured)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *MockE2ELBTCTokenPoolChainConfiguredIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolChainConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolChainConfigured struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterChainConfigured(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolChainConfiguredIterator, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolChainConfiguredIterator{contract: _MockE2ELBTCTokenPool.contract, event: "ChainConfigured", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolChainConfigured) (event.Subscription, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolChainConfigured)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseChainConfigured(log types.Log) (*MockE2ELBTCTokenPoolChainConfigured, error) {
	event := new(MockE2ELBTCTokenPoolChainConfigured)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolChainRemovedIterator struct {
	Event *MockE2ELBTCTokenPoolChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolChainRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(MockE2ELBTCTokenPoolChainRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *MockE2ELBTCTokenPoolChainRemovedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterChainRemoved(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolChainRemovedIterator, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolChainRemovedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolChainRemoved) (event.Subscription, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolChainRemoved)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseChainRemoved(log types.Log) (*MockE2ELBTCTokenPoolChainRemoved, error) {
	event := new(MockE2ELBTCTokenPoolChainRemoved)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolConfigChangedIterator struct {
	Event *MockE2ELBTCTokenPoolConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolConfigChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(MockE2ELBTCTokenPoolConfigChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *MockE2ELBTCTokenPoolConfigChangedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolConfigChanged struct {
	Config RateLimiterConfig
	Raw    types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolConfigChangedIterator, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolConfigChangedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolConfigChanged) (event.Subscription, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolConfigChanged)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseConfigChanged(log types.Log) (*MockE2ELBTCTokenPoolConfigChanged, error) {
	event := new(MockE2ELBTCTokenPoolConfigChanged)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolInboundRateLimitConsumedIterator struct {
	Event *MockE2ELBTCTokenPoolInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolInboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(MockE2ELBTCTokenPoolInboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *MockE2ELBTCTokenPoolInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolInboundRateLimitConsumedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "InboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolInboundRateLimitConsumed)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseInboundRateLimitConsumed(log types.Log) (*MockE2ELBTCTokenPoolInboundRateLimitConsumed, error) {
	event := new(MockE2ELBTCTokenPoolInboundRateLimitConsumed)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolLockedOrBurnedIterator struct {
	Event *MockE2ELBTCTokenPoolLockedOrBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolLockedOrBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolLockedOrBurned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(MockE2ELBTCTokenPoolLockedOrBurned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *MockE2ELBTCTokenPoolLockedOrBurnedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolLockedOrBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolLockedOrBurned struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolLockedOrBurnedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolLockedOrBurnedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "LockedOrBurned", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolLockedOrBurned)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseLockedOrBurned(log types.Log) (*MockE2ELBTCTokenPoolLockedOrBurned, error) {
	event := new(MockE2ELBTCTokenPoolLockedOrBurned)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolOutboundRateLimitConsumedIterator struct {
	Event *MockE2ELBTCTokenPoolOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolOutboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(MockE2ELBTCTokenPoolOutboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *MockE2ELBTCTokenPoolOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolOutboundRateLimitConsumedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "OutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolOutboundRateLimitConsumed)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseOutboundRateLimitConsumed(log types.Log) (*MockE2ELBTCTokenPoolOutboundRateLimitConsumed, error) {
	event := new(MockE2ELBTCTokenPoolOutboundRateLimitConsumed)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolOwnershipTransferRequestedIterator struct {
	Event *MockE2ELBTCTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolOwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(MockE2ELBTCTokenPoolOwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *MockE2ELBTCTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MockE2ELBTCTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolOwnershipTransferRequestedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolOwnershipTransferRequested)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*MockE2ELBTCTokenPoolOwnershipTransferRequested, error) {
	event := new(MockE2ELBTCTokenPoolOwnershipTransferRequested)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolOwnershipTransferredIterator struct {
	Event *MockE2ELBTCTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(MockE2ELBTCTokenPoolOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *MockE2ELBTCTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MockE2ELBTCTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolOwnershipTransferredIterator{contract: _MockE2ELBTCTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolOwnershipTransferred)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*MockE2ELBTCTokenPoolOwnershipTransferred, error) {
	event := new(MockE2ELBTCTokenPoolOwnershipTransferred)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolRateLimitAdminSetIterator struct {
	Event *MockE2ELBTCTokenPoolRateLimitAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolRateLimitAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolRateLimitAdminSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(MockE2ELBTCTokenPoolRateLimitAdminSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *MockE2ELBTCTokenPoolRateLimitAdminSetIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolRateLimitAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolRateLimitAdminSet struct {
	RateLimitAdmin common.Address
	Raw            types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterRateLimitAdminSet(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolRateLimitAdminSetIterator, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolRateLimitAdminSetIterator{contract: _MockE2ELBTCTokenPool.contract, event: "RateLimitAdminSet", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolRateLimitAdminSet) (event.Subscription, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolRateLimitAdminSet)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseRateLimitAdminSet(log types.Log) (*MockE2ELBTCTokenPoolRateLimitAdminSet, error) {
	event := new(MockE2ELBTCTokenPoolRateLimitAdminSet)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolReleasedOrMintedIterator struct {
	Event *MockE2ELBTCTokenPoolReleasedOrMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolReleasedOrMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolReleasedOrMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(MockE2ELBTCTokenPoolReleasedOrMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *MockE2ELBTCTokenPoolReleasedOrMintedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolReleasedOrMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolReleasedOrMinted struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Recipient           common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolReleasedOrMintedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolReleasedOrMintedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "ReleasedOrMinted", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolReleasedOrMinted)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseReleasedOrMinted(log types.Log) (*MockE2ELBTCTokenPoolReleasedOrMinted, error) {
	event := new(MockE2ELBTCTokenPoolReleasedOrMinted)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolRemotePoolAddedIterator struct {
	Event *MockE2ELBTCTokenPoolRemotePoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolRemotePoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolRemotePoolAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(MockE2ELBTCTokenPoolRemotePoolAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *MockE2ELBTCTokenPoolRemotePoolAddedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolRemotePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolRemotePoolAdded struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolRemotePoolAddedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolRemotePoolAddedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "RemotePoolAdded", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolRemotePoolAdded)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseRemotePoolAdded(log types.Log) (*MockE2ELBTCTokenPoolRemotePoolAdded, error) {
	event := new(MockE2ELBTCTokenPoolRemotePoolAdded)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolRemotePoolRemovedIterator struct {
	Event *MockE2ELBTCTokenPoolRemotePoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolRemotePoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolRemotePoolRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(MockE2ELBTCTokenPoolRemotePoolRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *MockE2ELBTCTokenPoolRemotePoolRemovedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolRemotePoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolRemotePoolRemoved struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolRemotePoolRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolRemotePoolRemovedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "RemotePoolRemoved", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolRemotePoolRemoved)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseRemotePoolRemoved(log types.Log) (*MockE2ELBTCTokenPoolRemotePoolRemoved, error) {
	event := new(MockE2ELBTCTokenPoolRemotePoolRemoved)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolRouterUpdatedIterator struct {
	Event *MockE2ELBTCTokenPoolRouterUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolRouterUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolRouterUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(MockE2ELBTCTokenPoolRouterUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *MockE2ELBTCTokenPoolRouterUpdatedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolRouterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolRouterUpdated struct {
	OldRouter common.Address
	NewRouter common.Address
	Raw       types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterRouterUpdated(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolRouterUpdatedIterator, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolRouterUpdatedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "RouterUpdated", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolRouterUpdated) (event.Subscription, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolRouterUpdated)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseRouterUpdated(log types.Log) (*MockE2ELBTCTokenPoolRouterUpdated, error) {
	event := new(MockE2ELBTCTokenPoolRouterUpdated)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPool) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _MockE2ELBTCTokenPool.abi.Events["AllowListAdd"].ID:
		return _MockE2ELBTCTokenPool.ParseAllowListAdd(log)
	case _MockE2ELBTCTokenPool.abi.Events["AllowListRemove"].ID:
		return _MockE2ELBTCTokenPool.ParseAllowListRemove(log)
	case _MockE2ELBTCTokenPool.abi.Events["ChainAdded"].ID:
		return _MockE2ELBTCTokenPool.ParseChainAdded(log)
	case _MockE2ELBTCTokenPool.abi.Events["ChainConfigured"].ID:
		return _MockE2ELBTCTokenPool.ParseChainConfigured(log)
	case _MockE2ELBTCTokenPool.abi.Events["ChainRemoved"].ID:
		return _MockE2ELBTCTokenPool.ParseChainRemoved(log)
	case _MockE2ELBTCTokenPool.abi.Events["ConfigChanged"].ID:
		return _MockE2ELBTCTokenPool.ParseConfigChanged(log)
	case _MockE2ELBTCTokenPool.abi.Events["InboundRateLimitConsumed"].ID:
		return _MockE2ELBTCTokenPool.ParseInboundRateLimitConsumed(log)
	case _MockE2ELBTCTokenPool.abi.Events["LockedOrBurned"].ID:
		return _MockE2ELBTCTokenPool.ParseLockedOrBurned(log)
	case _MockE2ELBTCTokenPool.abi.Events["OutboundRateLimitConsumed"].ID:
		return _MockE2ELBTCTokenPool.ParseOutboundRateLimitConsumed(log)
	case _MockE2ELBTCTokenPool.abi.Events["OwnershipTransferRequested"].ID:
		return _MockE2ELBTCTokenPool.ParseOwnershipTransferRequested(log)
	case _MockE2ELBTCTokenPool.abi.Events["OwnershipTransferred"].ID:
		return _MockE2ELBTCTokenPool.ParseOwnershipTransferred(log)
	case _MockE2ELBTCTokenPool.abi.Events["RateLimitAdminSet"].ID:
		return _MockE2ELBTCTokenPool.ParseRateLimitAdminSet(log)
	case _MockE2ELBTCTokenPool.abi.Events["ReleasedOrMinted"].ID:
		return _MockE2ELBTCTokenPool.ParseReleasedOrMinted(log)
	case _MockE2ELBTCTokenPool.abi.Events["RemotePoolAdded"].ID:
		return _MockE2ELBTCTokenPool.ParseRemotePoolAdded(log)
	case _MockE2ELBTCTokenPool.abi.Events["RemotePoolRemoved"].ID:
		return _MockE2ELBTCTokenPool.ParseRemotePoolRemoved(log)
	case _MockE2ELBTCTokenPool.abi.Events["RouterUpdated"].ID:
		return _MockE2ELBTCTokenPool.ParseRouterUpdated(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (MockE2ELBTCTokenPoolAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (MockE2ELBTCTokenPoolAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (MockE2ELBTCTokenPoolChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (MockE2ELBTCTokenPoolChainConfigured) Topic() common.Hash {
	return common.HexToHash("0x0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b")
}

func (MockE2ELBTCTokenPoolChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (MockE2ELBTCTokenPoolConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (MockE2ELBTCTokenPoolInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (MockE2ELBTCTokenPoolLockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
}

func (MockE2ELBTCTokenPoolOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944")
}

func (MockE2ELBTCTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (MockE2ELBTCTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (MockE2ELBTCTokenPoolRateLimitAdminSet) Topic() common.Hash {
	return common.HexToHash("0x44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174")
}

func (MockE2ELBTCTokenPoolReleasedOrMinted) Topic() common.Hash {
	return common.HexToHash("0xfc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0")
}

func (MockE2ELBTCTokenPoolRemotePoolAdded) Topic() common.Hash {
	return common.HexToHash("0x7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea")
}

func (MockE2ELBTCTokenPoolRemotePoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76")
}

func (MockE2ELBTCTokenPoolRouterUpdated) Topic() common.Hash {
	return common.HexToHash("0x02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f1684")
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPool) Address() common.Address {
	return _MockE2ELBTCTokenPool.address
}

type MockE2ELBTCTokenPoolInterface interface {
	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	GetSupportedChains(opts *bind.CallOpts) ([]uint64, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	GetTokenDecimals(opts *bind.CallOpts) (uint8, error)

	IDestPoolData(opts *bind.CallOpts) ([]byte, error)

	IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error)

	IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error)

	IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error)

	ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error)

	SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error)

	SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterAllowListAdd(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*MockE2ELBTCTokenPoolAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*MockE2ELBTCTokenPoolAllowListRemove, error)

	FilterChainAdded(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*MockE2ELBTCTokenPoolChainAdded, error)

	FilterChainConfigured(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolChainConfiguredIterator, error)

	WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolChainConfigured) (event.Subscription, error)

	ParseChainConfigured(log types.Log) (*MockE2ELBTCTokenPoolChainConfigured, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*MockE2ELBTCTokenPoolChainRemoved, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*MockE2ELBTCTokenPoolConfigChanged, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolInboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*MockE2ELBTCTokenPoolInboundRateLimitConsumed, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolLockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*MockE2ELBTCTokenPoolLockedOrBurned, error)

	FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolOutboundRateLimitConsumedIterator, error)

	WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseOutboundRateLimitConsumed(log types.Log) (*MockE2ELBTCTokenPoolOutboundRateLimitConsumed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MockE2ELBTCTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*MockE2ELBTCTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MockE2ELBTCTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*MockE2ELBTCTokenPoolOwnershipTransferred, error)

	FilterRateLimitAdminSet(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolRateLimitAdminSetIterator, error)

	WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolRateLimitAdminSet) (event.Subscription, error)

	ParseRateLimitAdminSet(log types.Log) (*MockE2ELBTCTokenPoolRateLimitAdminSet, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*MockE2ELBTCTokenPoolReleasedOrMinted, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolRemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*MockE2ELBTCTokenPoolRemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolRemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*MockE2ELBTCTokenPoolRemotePoolRemoved, error)

	FilterRouterUpdated(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolRouterUpdatedIterator, error)

	WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolRouterUpdated) (event.Subscription, error)

	ParseRouterUpdated(log types.Log) (*MockE2ELBTCTokenPoolRouterUpdated, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
