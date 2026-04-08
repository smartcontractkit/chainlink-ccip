// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cross_chain_pool_token

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

type BaseERC20ConstructorParams struct {
	Name             string
	Symbol           string
	MaxSupply        *big.Int
	PreMint          *big.Int
	PreMintRecipient common.Address
	Decimals         uint8
	CcipAdmin        common.Address
}

type IPoolV2TokenTransferFeeConfig struct {
	DestGasOverhead                         uint32
	DestBytesOverhead                       uint32
	DefaultBlockConfirmationsFeeUSDCents    uint32
	CustomBlockConfirmationsFeeUSDCents     uint32
	DefaultBlockConfirmationsTransferFeeBps uint16
	CustomBlockConfirmationsTransferFeeBps  uint16
	IsEnabled                               bool
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

type TokenPoolRateLimitConfigArgs struct {
	RemoteChainSelector       uint64
	CustomBlockConfirmations  bool
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
}

type TokenPoolTokenTransferFeeConfigArgs struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
}

var CrossChainPoolTokenMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenParams\",\"type\":\"tuple\",\"internalType\":\"struct BaseERC20.ConstructorParams\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"maxSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"preMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"preMintRecipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"ccipAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"advancedPoolHooks\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationsFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationsFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationsTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationsTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"disableTokenTransferFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"_decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAdvancedPoolHooks\",\"inputs\":[],\"outputs\":[{\"name\":\"advancedPoolHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCIPAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"ccipAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmations\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationsRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinBlockConfirmations\",\"inputs\":[],\"outputs\":[{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockConfirmationsRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationsFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationsFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationsTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationsTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"lockOrBurnOutV1\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"blockConfirmationsRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maxSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"_maxSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"blockConfirmationsRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCCIPAdmin\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinBlockConfirmations\",\"inputs\":[{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.RateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmations\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"updateAdvancedPoolHooks\",\"inputs\":[{\"name\":\"newHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AdvancedPoolHooksUpdated\",\"inputs\":[{\"name\":\"oldHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"},{\"name\":\"newHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCIPAdminTransferred\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationsInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationsOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"feeAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinBlockConfirmationsSet\",\"inputs\":[{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmations\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationsFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationsFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationsTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationsTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotOwnerOrFeeAdmin\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotRenounceCCIPAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustomBlockConfirmationsNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmations\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MaxSupplyExceeded\",\"inputs\":[{\"name\":\"supplyAfterMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCCIPAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"PreMintAddressNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PreMintRecipientSetWithZeroPreMint\",\"inputs\":[{\"name\":\"preMintRecipient\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x61012080604052346106c557616360803803809161001d82856106ca565b83398101906080818303126106c55780516001600160401b0381116106c55781019160e0838203126106c5576040519060e082016001600160401b038111838210176105c25760405283516001600160401b0381116106c557816100829186016106ed565b82526020840151906001600160401b0382116106c5576100a39185016106ed565b9182602083015260408401519060408301918252606085015193606084019485526100d06080870161075c565b936080810194855260a08701519160ff8316978884036106c55760c06100fd9160a085019586520161075c565b9760c083019889526101116020860161075c565b9161012a60606101236040890161075c565b970161075c565b93518051906001600160401b0382116105c25760035490600182811c921680156106bb575b60208310146105a25781601f84931161064b575b50602090601f83116001146105e3576000926105d8575b50508160011b916000199060031b1c1916176003555b8051906001600160401b0382116105c25760045490600182811c921680156105b8575b60208310146105a25781601f849311610532575b50602090601f83116001146104ca576000926104bf575b50508160011b916000199060031b1c1916176004555b33156104ae57600680546001600160a01b031916331790553015801561049d575b801561048c575b61047b573060805260c09490945260a093909352600880546001600160a01b039485166001600160a01b03199182161790915560078054929094169116179091555160ff1660e05251610100528151156104505780516001600160a01b03161561043f57519051906001600160a01b031680156104295760025491808301809311610413576020926002557fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef600093849284845283825260408420818154019055604051908152a361010051806103ed575b50505b516001600160a01b0316806103e85750335b601280546001600160a01b039283166001600160a01b03198216811790925560405192167f9524c9e4b0b61eb018dd58a1cd856e3e74009528328ab4a613b434fa631d7242600080a3615bef908161077182396080518181816117c00152818161222101528181612b2f015281816130bf0152818161373e01526138f1015260a05181818161361101528181614c5e01528181614ca8015261501f015260c051818181610bda0152818161184e015281816122ae01528181612bbd015261314d015260e0518161353a0152610100518181816111d401528181615348015261548a0152f35b610303565b6002548181116103fd57506102ee565b637502c12360e11b835260045260245260449150fd5b634e487b7160e01b600052601160045260246000fd5b63ec442f0560e01b600052600060045260246000fd5b634dd371db60e11b60005260046000fd5b516001600160a01b031690508061046757506102f1565b63f5c8f5a160e01b60005260045260246000fd5b630a64406560e11b60005260046000fd5b506001600160a01b0385161561021c565b506001600160a01b03831615610215565b639b15e16f60e01b60005260046000fd5b0151905038806101de565b600460009081528281209350601f198516905b81811061051a5750908460019594939210610501575b505050811b016004556101f4565b015160001960f88460031b161c191690553880806104f3565b929360206001819287860151815501950193016104dd565b60046000529091507f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b601f840160051c81019160208510610598575b90601f859493920160051c01905b81811061058957506101c7565b6000815584935060010161057c565b909150819061056e565b634e487b7160e01b600052602260045260246000fd5b91607f16916101b3565b634e487b7160e01b600052604160045260246000fd5b01519050388061017a565b600360009081528281209350601f198516905b818110610633575090846001959493921061061a575b505050811b01600355610190565b015160001960f88460031b161c1916905538808061060c565b929360206001819287860151815501950193016105f6565b60036000529091507fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b601f840160051c810191602085106106b1575b90601f859493920160051c01905b8181106106a25750610163565b60008155849350600101610695565b9091508190610687565b91607f169161014f565b600080fd5b601f909101601f19168101906001600160401b038211908210176105c257604052565b81601f820112156106c5578051906001600160401b0382116105c25760405192610721601f8401601f1916602001856106ca565b828452602083830101116106c55760005b82811061074757505060206000918301015290565b80602080928401015182828701015201610732565b51906001600160a01b03821682036106c55756fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a714613b485750806306fdde0314613aa1578063095ea7b31461399457806318160ddd14613976578063181f5a771461391557806321df0da7146138d157806323b872dd1461376b578063240028e8146137145780632422ac451461363557806324f65ee7146135f75780632c0634041461355e578063313ce5671461352057806337a3210d146134f9578063390775371461302d578063489a68f214612a955780634c5ef0ed14612a4e57806362ddd3c4146129c757806370a08231146129905780637437ff9f1461294f57806379ba5097146128a05780638926f54f1461285a57806389720a62146127a05780638da5cb5b146127795780638fd6a6ac1461275257806395d89b41146126655780639a4575b9146121b5578063a42a7b8b1461203d578063a8fa343c14611fbc578063a9059cbb14611f8a578063acfecf9114611e92578063ae39a25714611d3b578063b1c71c6514611725578063b7946580146116e8578063b8d5005e146116c3578063bfeffd3f14611631578063c4bffe2b14611506578063c7230a6014611296578063d4d6de23146111f7578063d5abeb01146111bc578063d8aa3f4014611082578063dc04fa1f14610bfe578063dc0bd97114610bba578063dcbd41bc146109d0578063dd62ed3e14610980578063e8a1da17146102d45763f2fde38b1461021b57600080fd5b346102d15760206003193601126102d1576001600160a01b0361023c613d6f565b610244614dac565b163381146102a957807fffffffffffffffffffffffff000000000000000000000000000000000000000060055416176005556001600160a01b03600654167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b80fd5b50346102d15760406003193601126102d15760043567ffffffffffffffff81116107b55761030690369060040161404b565b9060243567ffffffffffffffff811161097c57906103298492369060040161404b565b939091610334614dac565b83905b8282106107bd5750505081927ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b818110156107b9578060051b830135858112156107b1578301610120813603126107b1576040519461039b86613efc565b6103a482613df8565b8652602082013567ffffffffffffffff81116107b55782019436601f870112156107b5578535956103d4876144f2565b966103e26040519889613f18565b80885260208089019160051b830101903682116107b15760208301905b82821061077e575050505060208701958652604083013567ffffffffffffffff811161077a576104329036908501613fbe565b916040880192835261045c61044a36606087016149a2565b9460608a0195865260c03691016149a2565b9560808901968752835151156107525761048067ffffffffffffffff8a51166158d4565b1561071b5767ffffffffffffffff8951168252600d602052604082206104a78651826150b6565b6104b58851600283016150b6565b6004855191019080519067ffffffffffffffff82116106ee576104d883546140ad565b601f81116106b3575b50602090601f8311600114610632576105119291869183610627575b50506000198260011b9260031b1c19161790565b90555b815b8851805182101561054b579061054560019261053e8367ffffffffffffffff8f5116926148b8565b5190614dea565b01610516565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c293919997509561061967ffffffffffffffff60019796949851169251935191516105e56105b060405196879687526101006020880152610100870190613d10565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a101939290919361036a565b015190508e806104fd565b83865281862091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416875b81811061069b5750908460019594939210610682575b505050811b019055610514565b015160001960f88460031b161c191690558d8080610675565b9293602060018192878601518155019501930161065f565b6106de9084875260208720601f850160051c810191602086106106e4575b601f0160051c0190614a3e565b8d6104e1565b90915081906106d1565b6024857f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60248267ffffffffffffffff8b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b6004827f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b8280fd5b813567ffffffffffffffff81116107ad576020916107a28392833691890101613fbe565b8152019101906103ff565b8680fd5b8480fd5b5080fd5b8380f35b9267ffffffffffffffff6107df6107da8486889a9699979a614975565b6144a0565b16916107ea83615718565b1561095057828452600d602052610806600560408620016156b5565b94845b865181101561083f57600190858752600d60205261083860056040892001610831838b6148b8565b5190615818565b5001610809565b50939692909450949094808752600d602052600560408820888155886001820155886002820155886003820155886004820161087b81546140ad565b8061090f575b50505001805490888155816108f1575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020836001948a52600982528985604082208281550155808a52600a82528985604082208281550155604051908152a101909194939294610337565b885260208820908101905b81811015610891578881556001016108fc565b601f81116001146109255750555b888a80610881565b8183526020832061094091601f01861c810190600101614a3e565b808252816020812091555561091d565b602484847f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b8380fd5b50346102d15760406003193601126102d1576001600160a01b0360406109a4613d6f565b92826109ae613d8a565b9416815260016020522091166000526020526020604060002054604051908152f35b50346102d15760206003193601126102d15760043567ffffffffffffffff81116107b557610a0290369060040161407c565b6001600160a01b03600f541633141580610ba5575b610b7957825b818110610a28578380f35b610a33818385614927565b67ffffffffffffffff610a45826144a0565b1690610a5e82600052600c602052604060002054151590565b15610b4d57907f41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb60e083610b0d610ae7602060019897018b610a9f82614937565b15610b14578790526009602052610ac660408d20610ac036604088016149a2565b906150b6565b868c52600a602052610ae260408d20610ac03660a088016149a2565b614937565b916040519215158352610b0060208401604083016149fa565b60a06080840191016149fa565ba201610a1d565b60026040828a610ae29452600d602052610b36828220610ac036858c016149a2565b8a8152600d6020522001610ac03660a088016149a2565b602486837f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b6024837f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b506001600160a01b0360065416331415610a17565b50346102d157806003193601126102d15760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102d15760406003193601126102d15760043567ffffffffffffffff81116107b557610c3090369060040161407c565b60243567ffffffffffffffff811161097c57610c5090369060040161404b565b919092610c5b614dac565b845b828110610cc757505050825b818110610c74578380f35b8067ffffffffffffffff610c8e6107da6001948688614975565b1680865260106020528560408120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee88680a201610c69565b67ffffffffffffffff610cde6107da838686614927565b16610cf681600052600c602052604060002054151590565b1561105757610d06828585614927565b602081019060e0810190610d1982614937565b1561102b5760a0810161271061ffff610d3183614944565b16101561101c5760c082019161271061ffff610d4c85614944565b161015610fe45763ffffffff610d6186614953565b1615610fb857858c52601060205260408c20610d7c86614953565b63ffffffff16908054906040840191610d9483614953565b60201b67ffffffff0000000016936060860194610db086614953565b60401b6bffffffff0000000000000000169660800196610dcf88614953565b60601b6fffffffff0000000000000000000000001691610dee8a614944565b60801b71ffff000000000000000000000000000000001693610e0f8c614944565b60901b73ffff00000000000000000000000000000000000016957fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016177fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16177fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff16177fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff161717178155610ec287614937565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff00000000000000000000000000000000000000001617905560405196610f1390614964565b63ffffffff168752610f2490614964565b63ffffffff166020870152610f3890614964565b63ffffffff166040860152610f4c90614964565b63ffffffff166060850152610f6090613e3c565b61ffff166080840152610f7290613e3c565b61ffff1660a0830152610f8490613e0d565b151560c082015260e07ffae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa104191a2600101610c5d565b60248c877f12332265000000000000000000000000000000000000000000000000000000008252600452fd5b60248c61ffff610ff386614944565b7f95f3517a00000000000000000000000000000000000000000000000000000000835216600452fd5b8a61ffff610ff3602493614944565b60248a857f12332265000000000000000000000000000000000000000000000000000000008252600452fd5b7f1e670e4b000000000000000000000000000000000000000000000000000000008752600452602486fd5b50346102d15760806003193601126102d15761109c613d6f565b506110a5613de1565b6110ad613e2b565b5060643567ffffffffffffffff811161077a579167ffffffffffffffff6040926110dd60e0953690600401613e4b565b50508260c085516110ed81613ee0565b82815282602082015282878201528260608201528260808201528260a082015201521681526010602052206040519061112582613ee0565b5461ffff818163ffffffff82169485815263ffffffff60208201818560201c1681528160408401818760401c168152816060860193818960601c16855260ff60c060808901988a8c60801c168a528a60a082019c60901c168c52019b60a01c1615158b526040519b8c52511660208b0152511660408901525116606087015251166080850152511660a083015251151560c0820152f35b50346102d157806003193601126102d15760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b50346102d15760206003193601126102d15760043561ffff81169081810361077a577f46c9c0585a955b2702c7ea47fec541db623565d20827a0edda42864e6b859a0191602091611246614dac565b7fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff75ffff00000000000000000000000000000000000000006007549260a01b16911617600755604051908152a180f35b50346102d15760406003193601126102d15760043567ffffffffffffffff81116107b5576112c890369060040161404b565b906112d1613d8a565b916001600160a01b0360065416331415806114f1575b6114c5576001600160a01b03831690811561149d57845b818110611309578580f35b6001600160a01b0361132461131f838588614975565b61448c565b166040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481855afa90811561149257889161145f575b5080611379575b50506001016112fe565b6040517fa9059cbb0000000000000000000000000000000000000000000000000000000060208083019182526001600160a01b038a16602484015260448084018590528352918a91906113cd606482613f18565b519082865af1156114545787513d61144b5750813b155b61141f5790847f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e602060019594604051908152a3903861136f565b602488837f5274afe7000000000000000000000000000000000000000000000000000000008252600452fd5b600114156113e4565b6040513d89823e3d90fd5b905060203d811161148b575b6114758183613f18565b602082600092810103126102d157505138611368565b503d61146b565b6040513d8a823e3d90fd5b6004857f8579befe000000000000000000000000000000000000000000000000000000008152fd5b6024847fcb1afbd700000000000000000000000000000000000000000000000000000000815233600452fd5b506001600160a01b03601154163314156112e7565b50346102d157806003193601126102d15760405190600b548083528260208101600b84526020842092845b81811061161857505061154692500383613f18565b815161156a611554826144f2565b916115626040519384613f18565b8083526144f2565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602083019301368437805b84518110156115c9578067ffffffffffffffff6115b6600193886148b8565b51166115c282866148b8565b5201611597565b50925090604051928392602084019060208552518091526040840192915b8181106115f5575050500390f35b825167ffffffffffffffff168452859450602093840193909201916001016115e7565b8454835260019485019487945060209093019201611531565b50346102d15760206003193601126102d1576004356001600160a01b0381168091036107b55761165f614dac565b7fffffffffffffffffffffffff00000000000000000000000000000000000000006008547fbaff46844acf36d6ee996f489a1a288709c4542bd33cd557770afd267d581209604080516001600160a01b0384168152856020820152a1161760085580f35b50346102d157806003193601126102d157602061ffff60075460a01c16604051908152f35b50346102d15760206003193601126102d15761172161170d611708613dca565b6148fb565b604051918291602083526020830190613d10565b0390f35b50346102d15760606003193601126102d1576004359067ffffffffffffffff82116102d1578160040160a060031984360301126107b557611764613e1a565b9060443567ffffffffffffffff811161097c57906117896117a6923690600401613e4b565b929061179361489f565b5061179e8584615053565b933691613f59565b9260848601936117b58561448c565b6001600160a01b03807f000000000000000000000000000000000000000000000000000000000000000016911603611cfe57602487019677ffffffffffffffff0000000000000000000000000000000061180e896144a0565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611492578891611ccf575b50611ca75767ffffffffffffffff611895896144a0565b166118ad81600052600c602052604060002054151590565b15611c7c5760206001600160a01b0360075416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015611492578890611c38575b6001600160a01b039150163303611c0c5760648101359361ffff6119258787614585565b9416938415611beb5761ffff60075460a01c168015611bc357808610611b935750611962816119538a61448c565b61195c8d6144a0565b90615645565b6001600160a01b03600854169384611a7c575b5050505050509061198591614585565b9161198f846144a0565b503015611a505750611708611a4693611a15926119ac85306152e8565b7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff6119e86119e2856144a0565b9361448c565b604080516001600160a01b039092168252336020830152810188905292169180606081015b0390a26144a0565b90611a1e615018565b60405192611a2b84613ec4565b83526020830152604051928392604084526040840190614021565b9060208301520390f35b807f96c6fd1e000000000000000000000000000000000000000000000000000000006024925280600452fd5b843b15611b8f578994928b9694928692604051988997889687957f1ff7703e000000000000000000000000000000000000000000000000000000008752600487016080905280611acb916155af565b6084880160a09052610124880190611ae2926145a6565b92611aec90613df8565b67ffffffffffffffff1660a4870152604401611b0790613db6565b6001600160a01b031660c48601528d8c60e4870152611b2590613db6565b6001600160a01b03166101048601526024850152838103600319016044850152611b4e91613d10565b90606483015203925af18015611b8457908591611b6f575b80808080611975565b81611b7991613f18565b61097c578338611b66565b6040513d87823e3d90fd5b8980fd5b89604491877f1f5b9f77000000000000000000000000000000000000000000000000000000008352600452602452fd5b60048a7f98d50fd7000000000000000000000000000000000000000000000000000000008152fd5b611c0781611bf88a61448c565b611c018d6144a0565b906155ff565b611962565b6024877f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d602011611c74575b81611c5260209383613f18565b81010312611c7057611c6b6001600160a01b0391614592565b611901565b8780fd5b3d9150611c45565b7fa9902c7e000000000000000000000000000000000000000000000000000000008852600452602487fd5b6004877f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b611cf1915060203d602011611cf7575b611ce98183613f18565b810190614d94565b3861187e565b503d611cdf565b6024866001600160a01b03611d128861448c565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835216600452fd5b50346102d15760606003193601126102d157611d55613d6f565b90611d5e613d8a565b604435926001600160a01b03841680850361097c57611d7b614dac565b6001600160a01b0382168015611e6a5794611e64917f3f1036e85d016a93254a0b1415844f79b85424959d90ae5ad51ce8f4533fe70195967fffffffffffffffffffffffff000000000000000000000000000000000000000060075416176007556001600160a01b0385167fffffffffffffffffffffffff0000000000000000000000000000000000000000600f541617600f557fffffffffffffffffffffffff00000000000000000000000000000000000000006011541617601155604051938493849160409194936001600160a01b03809281606087019816865216602085015216910152565b0390a180f35b6004857f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b50346102d15767ffffffffffffffff611eaa36613fdc565b929091611eb5614dac565b1691611ece83600052600c602052604060002054151590565b1561095057828452600d602052611efd60056040862001611ef0368486613f59565b6020815191012090615818565b15611f4257907f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d7691611f3c6040519283926020845260208401916145a6565b0390a280f35b82611f86836040519384937f74f23c7c00000000000000000000000000000000000000000000000000000000855260048501526040602485015260448401916145a6565b0390fd5b50346102d15760406003193601126102d157611fb1611fa7613d6f565b6024359033614a55565b602060405160018152f35b50346102d15760206003193601126102d157611fd6613d6f565b611fde614dac565b6001600160a01b0380601254921691827fffffffffffffffffffffffff0000000000000000000000000000000000000000821617601255167f9524c9e4b0b61eb018dd58a1cd856e3e74009528328ab4a613b434fa631d72428380a380f35b50346102d15760206003193601126102d15767ffffffffffffffff612060613dca565b168152600d602052612077600560408320016156b5565b80517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06120bc6120a6836144f2565b926120b46040519485613f18565b8084526144f2565b01835b8181106121a4575050825b825181101561212157806120e0600192856148b8565b518552600e6020526120fe6121056040872060405192838092614100565b0382613f18565b61210f82856148b8565b5261211a81846148b8565b50016120ca565b81846040519182916020830160208452825180915260408401602060408360051b870101940192905b82821061215957505050500390f35b91936020612194827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851613d10565b960192019201859493919261214a565b8060606020809386010152016120bf565b50346102d15760206003193601126102d15760043567ffffffffffffffff81116107b557806004019060a0600319823603011261077a576121f461489f565b506040516020936122058583613f18565b80825260848301916122168361448c565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691160361265157602484019477ffffffffffffffff0000000000000000000000000000000061226f876144a0565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015287816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156125d6578491612634575b5061260c5767ffffffffffffffff6122f5876144a0565b1661230d81600052600c602052604060002054151590565b156125e157876001600160a01b0360075416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa80156125d657849061259b575b6001600160a01b03915016330361256f5760648501359461238d866123848761448c565b611c018a6144a0565b6001600160a01b03600854169182612461575b505050506123ad846144a0565b503015611a50575091817ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff612431956123f261170896306152e8565b611a0d612407612401876144a0565b9261448c565b604080516001600160a01b0390921682523360208301528101959095529116929081906060820190565b9061243a615018565b6040519261244784613ec4565b835281830152611721604051928284938452830190614021565b823b156107b157918791858094604051968795869485937f1ff7703e0000000000000000000000000000000000000000000000000000000085526004850160809052806124ad916155af565b6084860160a090526101248601906124c4926145a6565b916124ce90613df8565b67ffffffffffffffff1660a48501526044016124e990613db6565b6001600160a01b031660c48401528b60e48401526125068b613db6565b6001600160a01b031661010484015283602484015282810360031901604484015261253091613d10565b8a606483015203925af180156125645790829161254f575b80806123a0565b8161255991613f18565b6102d1578038612548565b6040513d84823e3d90fd5b6024837f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b508781813d83116125cf575b6125b18183613f18565b8101031261097c576125ca6001600160a01b0391614592565b612360565b503d6125a7565b6040513d86823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008452600452602483fd5b6004837f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b61264b9150883d8a11611cf757611ce98183613f18565b386122de565b506001600160a01b03611d1260249361448c565b50346102d157806003193601126102d157604051908060045490612688826140ad565b808552916001811690811561270d57506001146126b0575b6117218461170d81860382613f18565b600481527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b939250905b8082106126f35750909150810160200161170d826126a0565b9192600181602092548385880101520191019092916126da565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208087019190915292151560051b8501909201925061170d91508390506126a0565b50346102d157806003193601126102d15760206001600160a01b0360125416604051908152f35b50346102d157806003193601126102d15760206001600160a01b0360065416604051908152f35b50346102d15760c06003193601126102d1576127ba613d6f565b6127c2613de1565b9060643561ffff8116810361097c5760843567ffffffffffffffff81116107b1576127f1903690600401613e4b565b9160a4359360028510156107ad5761280c95604435916145e5565b90604051918291602083016020845282518091526020604085019301915b818110612838575050500390f35b82516001600160a01b031684528594506020938401939092019160010161282a565b50346102d15760206003193601126102d157602061289667ffffffffffffffff612882613dca565b16600052600c602052604060002054151590565b6040519015158152f35b50346102d157806003193601126102d1576005546001600160a01b0381163303612927577fffffffffffffffffffffffff0000000000000000000000000000000000000000600654913382841617600655166005556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b50346102d157806003193601126102d157600754600f54601154604080516001600160a01b0394851681529284166020840152921691810191909152606090f35b50346102d15760206003193601126102d15760406020916001600160a01b036129b7613d6f565b1681528083522054604051908152f35b50346102d1576129d636613fdc565b6129e293929193614dac565b67ffffffffffffffff8216612a0481600052600c602052604060002054151590565b15612a235750612a209293612a1a913691613f59565b90614dea565b80f35b7f1e670e4b000000000000000000000000000000000000000000000000000000008452600452602483fd5b50346102d15760406003193601126102d157612a68613dca565b906024359067ffffffffffffffff82116102d157602061289684612a8f3660048701613fbe565b906144b5565b50346102d15760406003193601126102d1576004359067ffffffffffffffff82116102d157816004019061010060031984360301126102d157612ad6613e1a565b9181604051612ae481613e79565b5260648401359360c4810193612b15612b0f612b0a612b03888761443b565b3691613f59565b614bea565b87614ca5565b946084830196612b248861448c565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691160361301957602484019477ffffffffffffffff00000000000000000000000000000000612b7d876144a0565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611492578891612ffa575b50611ca75767ffffffffffffffff612c04876144a0565b16612c1c81600052600c602052604060002054151590565b15611c7c5760206001600160a01b0360075416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115611492578891612fdb575b5015611c0c57612c86866144a0565b93612c9c60a4870195612a8f612b03888661443b565b15612f945761ffff16908115612f7357612cc889612cb98c61448c565b612cc28a6144a0565b9061553f565b6001600160a01b03600854169384612dba575b50505050505060440191612cee8361448c565b90612cf8836144a0565b506001600160a01b03821615612d8e575067ffffffffffffffff6020956001600160a01b03612d5a612d546119e27ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0976107da8b608099615431565b9661448c565b816040519716875233898801521660408601528560608601521692a260405190612d8382613e79565b815260405190518152f35b807fec442f05000000000000000000000000000000000000000000000000000000006024925280600452fd5b843b15612f6f57868995938c959387938b6040519a8b998a9889977f1abfe46e0000000000000000000000000000000000000000000000000000000089526004890160609052612e0a87806155af565b60648b0161010090526101648b0190612e22926145a6565b94612e2c90613df8565b67ffffffffffffffff1660848a0152604401612e4790613db6565b6001600160a01b031660a489015260c4880152612e6390613db6565b6001600160a01b031660e4870152612e7b90846155af565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c87840301610104880152612eb092916145a6565b90612ebb90836155af565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c86840301610124870152612ef092916145a6565b9060e48b01612efe916155af565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c85840301610144860152612f3392916145a6565b908c6024840152604483015203925af180156125d657908491612f5a575b80808080612cdb565b81612f6491613f18565b61077a578238612f51565b8880fd5b612f8f89612f808c61448c565b612f898a6144a0565b906154d3565b612cc8565b612f9e858361443b565b611f866040519283927f24eb47e50000000000000000000000000000000000000000000000000000000084526020600485015260248401916145a6565b612ff4915060203d602011611cf757611ce98183613f18565b38612c77565b613013915060203d602011611cf757611ce98183613f18565b38612bed565b6024866001600160a01b03611d128b61448c565b50346102d15760206003193601126102d1576004359067ffffffffffffffff82116102d157816004019061010060031984360301126102d1578060405161307381613e79565b528060405161308181613e79565b52606483013560c48401936130a561309f612b0a612b03888861443b565b83614ca5565b9360848201956130b48761448c565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116036134e557602483019377ffffffffffffffff0000000000000000000000000000000061310d866144a0565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156114545787916134c6575b5061349e5767ffffffffffffffff613194866144a0565b166131ac81600052600c602052604060002054151590565b156134735760206001600160a01b0360075416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115611454578791613454575b501561342857613216856144a0565b9261322c60a4860194612a8f612b03878561443b565b1561341e576132478861323e8b61448c565b612f89896144a0565b6001600160a01b0360085416928361326c575b505050505060440191612cee8361448c565b833b15611c7057878795938195938c93604051988997889687957f1abfe46e00000000000000000000000000000000000000000000000000000000875260048701606090528d6132bc87806155af565b60648a0161010090526101648a01906132d4926145a6565b946132de90613df8565b67ffffffffffffffff1660848901526044016132f990613db6565b6001600160a01b031660a488015260c487015261331590613db6565b6001600160a01b031660e486015261332d90846155af565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c8684030161010487015261336292916145a6565b9061336d90836155af565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c858403016101248601526133a292916145a6565b9060e48a016133b0916155af565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c848403016101448501526133e592916145a6565b8b602483015282604483015203925af180156125d657613409575b8080808061325a565b926134178160449395613f18565b9290613400565b83612f9e9161443b565b6024867f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b61346d915060203d602011611cf757611ce98183613f18565b38613207565b7fa9902c7e000000000000000000000000000000000000000000000000000000008752600452602486fd5b6004867f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b6134df915060203d602011611cf757611ce98183613f18565b3861317d565b6024856001600160a01b03611d128a61448c565b50346102d157806003193601126102d15760206001600160a01b0360085416604051908152f35b50346102d157806003193601126102d157602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102d15760c06003193601126102d157613578613d6f565b50613581613de1565b613589613da0565b506084359161ffff831683036102d15760a4359067ffffffffffffffff82116102d15760a063ffffffff8061ffff6135d088886135c93660048b01613e4b565b50506142b6565b95926040979194975197885216602087015216604085015216606083015215156080820152f35b50346102d157806003193601126102d157602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102d15760406003193601126102d15761364f613dca565b6024359182151583036102d15761014061371261366c8585614233565b6136c260409392935180946fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b60a08301906fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565bf35b50346102d15760206003193601126102d157602090613731613d6f565b90506001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b50346102d15760606003193601126102d157613785613d6f565b61378d613d8a565b604435916001600160a01b0381168085526001602052604085206001600160a01b033316865260205260408520549060001982106137d2575b5050611fb19350614a55565b84821061389d57303314613871578015613845573315613819576040868692611fb1985260016020528181206001600160a01b0333168252602052209103905538806137c6565b6024867f94280d6200000000000000000000000000000000000000000000000000000000815280600452fd5b6024867fe602df0500000000000000000000000000000000000000000000000000000000815280600452fd5b6024867f94280d6200000000000000000000000000000000000000000000000000000000815233600452fd5b60648686847ffb8f41b200000000000000000000000000000000000000000000000000000000835233600452602452604452fd5b50346102d157806003193601126102d15760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102d157806003193601126102d15750611721604051613938604082613f18565b601d81527f43726f7373436861696e506f6f6c546f6b656e20322e302e302d6465760000006020820152604051918291602083526020830190613d10565b50346102d157806003193601126102d1576020600254604051908152f35b50346102d15760406003193601126102d1576139ae613d6f565b6001600160a01b03602435911691308314613a75573315613a49578215613a1d5760408291338152600160205281812085825260205220556040519081527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560203392a3602060405160018152f35b807f94280d62000000000000000000000000000000000000000000000000000000006024925280600452fd5b807fe602df05000000000000000000000000000000000000000000000000000000006024925280600452fd5b80837f94280d620000000000000000000000000000000000000000000000000000000060249352600452fd5b50346102d157806003193601126102d157604051908060035490613ac4826140ad565b808552916001811690811561270d5750600114613aeb576117218461170d81860382613f18565b600381527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b939250905b808210613b2e5750909150810160200161170d826126a0565b919260018160209254838588010152019101909291613b15565b9050346107b55760206003193601126107b5576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361077a57602092507f36372b07000000000000000000000000000000000000000000000000000000008114908115613ce6575b8115613cbc575b8115613c92575b8115613bd3575b5015158152f35b7faff2afbf00000000000000000000000000000000000000000000000000000000811491508115613c68575b8115613c3e575b8115613c14575b5038613bcc565b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438613c0d565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150613c06565b7f331710310000000000000000000000000000000000000000000000000000000081149150613bff565b7f01ffc9a70000000000000000000000000000000000000000000000000000000081149150613bc5565b7fa219a0250000000000000000000000000000000000000000000000000000000081149150613bbe565b7f8fd6a6ac0000000000000000000000000000000000000000000000000000000081149150613bb7565b919082519283825260005b848110613d5a5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201613d1b565b600435906001600160a01b0382168203613d8557565b600080fd5b602435906001600160a01b0382168203613d8557565b606435906001600160a01b0382168203613d8557565b35906001600160a01b0382168203613d8557565b6004359067ffffffffffffffff82168203613d8557565b6024359067ffffffffffffffff82168203613d8557565b359067ffffffffffffffff82168203613d8557565b35908115158203613d8557565b6024359061ffff82168203613d8557565b6044359061ffff82168203613d8557565b359061ffff82168203613d8557565b9181601f84011215613d855782359167ffffffffffffffff8311613d855760208381860195010111613d8557565b6020810190811067ffffffffffffffff821117613e9557604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff821117613e9557604052565b60e0810190811067ffffffffffffffff821117613e9557604052565b60a0810190811067ffffffffffffffff821117613e9557604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117613e9557604052565b92919267ffffffffffffffff8211613e955760405191613fa1601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200184613f18565b829481845281830111613d85578281602093846000960137010152565b9080601f83011215613d8557816020613fd993359101613f59565b90565b906040600319830112613d855760043567ffffffffffffffff81168103613d8557916024359067ffffffffffffffff8211613d855761401d91600401613e4b565b9091565b613fd991602061403a8351604084526040840190613d10565b920151906020818403910152613d10565b9181601f84011215613d855782359167ffffffffffffffff8311613d85576020808501948460051b010111613d8557565b9181601f84011215613d855782359167ffffffffffffffff8311613d85576020808501948460081b010111613d8557565b90600182811c921680156140f6575b60208310146140c757565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f16916140bc565b60009291815491614110836140ad565b8083529260018116908115614166575060011461412c57505050565b60009081526020812093945091925b83831061414c575060209250010190565b60018160209294939454838587010152019101919061413b565b905060209495507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091509291921683830152151560051b010190565b604051906141af82613efc565b60006080838281528260208201528260408201528260608201520152565b906040516141da81613efc565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff916142456141a2565b5061424e6141a2565b506142825716600052600d602052604060002090613fd9614276600261427b614276866141cd565b614b65565b94016141cd565b169081600052600960205261429d61427660406000206141cd565b91600052600a602052613fd961427660406000206141cd565b9061ffff8060075460a01c1691169283151592838094614433575b6144095767ffffffffffffffff166000526010602052604060002091604051926142fa84613ee0565b5463ffffffff81168452602084019563ffffffff8260201c168752604085019263ffffffff8360401c168452606086019163ffffffff8460601c168352608087019761ffff8560801c16895260ff60a089019561ffff8160901c16875260a01c1615801560c08a01526143ee5761438f575050505063ffffffff808061ffff9351169451169551169351169193929190600190565b8193975080929450106143be57505063ffffffff808061ffff9351169451169551169351169193929190600190565b7f1f5b9f770000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b50505050505092505050600090600090600090600090600090565b7f98d50fd70000000000000000000000000000000000000000000000000000000060005260046000fd5b5082156142d1565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215613d85570180359067ffffffffffffffff8211613d8557602001918136038313613d8557565b356001600160a01b0381168103613d855790565b3567ffffffffffffffff81168103613d855790565b9067ffffffffffffffff613fd99216600052600d602052600560406000200190602081519101209060019160005201602052604060002054151590565b67ffffffffffffffff8111613e955760051b60200190565b8181029291811591840414171561451d57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8115614556570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b9190820391821161451d57565b51906001600160a01b0382168203613d8557565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b9295939091946001600160a01b036008541695861561487d578097600287101561484e576001600160a01b039861470a9561ffff93896148245767ffffffffffffffff821660005260106020526040600020906040519161464583613ee0565b549163ffffffff8316815263ffffffff8360201c16602082015263ffffffff8360401c16604082015263ffffffff8360601c16606082015260c0878460801c169182608082015260ff60a08201958a8160901c16875260a01c16151591829101526147d2575b50505067ffffffffffffffff905b6040519b8c997f89720a62000000000000000000000000000000000000000000000000000000008b521660048a0152166024880152604487015216606485015260c0608485015260c48401916145a6565b928180600095869560a483015203915afa9182156147c557819261472d57505090565b9091503d8083833e61473f8183613f18565b81019060208183031261077a5780519067ffffffffffffffff821161097c570181601f8201121561077a57805190614776826144f2565b936147846040519586613f18565b82855260208086019360051b8301019384116102d15750602001905b8282106147ad5750505090565b602080916147ba84614592565b8152019101906147a0565b50604051903d90823e3d90fd5b92935067ffffffffffffffff928587161561480c57506127106147fb876148029451168361450a565b0490614585565b915b9038806146ab565b61481e92506147fb612710918361450a565b91614804565b67ffffffffffffffff91925061484890614842612b0a36898b613f59565b90614ca5565b916146b9565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b5050505050505050604051614893602082613f18565b60008152600036813790565b604051906148ac82613ec4565b60606020838281520152565b80518210156148cc5760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b67ffffffffffffffff16600052600d6020526120fe613fd9600460406000200160405192838092614100565b91908110156148cc5760081b0190565b358015158103613d855790565b3561ffff81168103613d855790565b3563ffffffff81168103613d855790565b359063ffffffff82168203613d8557565b91908110156148cc5760051b0190565b35906fffffffffffffffffffffffffffffffff82168203613d8557565b9190826060910312613d85576040516060810181811067ffffffffffffffff821117613e955760405260406149f58183956149dc81613e0d565b85526149ea60208201614985565b602086015201614985565b910152565b6fffffffffffffffffffffffffffffffff614a3860408093614a1b81613e0d565b1515865283614a2c60208301614985565b16602087015201614985565b16910152565b818110614a49575050565b60008155600101614a3e565b6001600160a01b0316908115614b36576001600160a01b0316918215614b07576000828152806020526040812054828110614ad45791604082827fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef958760209652828652038282205586815280845220818154019055604051908152a3565b6064937fe450d38c0000000000000000000000000000000000000000000000000000000083949352600452602452604452fd5b7fec442f0500000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b7f96c6fd1e00000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b614b6d6141a2565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff8083511691614bca6020850193614bc4614bb763ffffffff87511642614585565b856080890151169061450a565b906154c6565b80821015614be357505b16825263ffffffff4216905290565b9050614bd4565b80518015614c5a57602003614c1c578051602082810191830183900312613d8557519060ff8211614c1c575060ff1690565b611f86906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190613d10565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff821161451d57565b60ff16604d811161451d57600a0a90565b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff811692828414614d8d57828411614d635790614cea91614c80565b91604d60ff8416118015614d48575b614d1257505090614d0c613fd992614c94565b9061450a565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b50614d5283614c94565b801561455657600019048411614cf9565b614d6c91614c80565b91604d60ff841611614d1257505090614d87613fd992614c94565b9061454c565b5050505090565b90816020910312613d8557518015158103613d855790565b6001600160a01b03600654163303614dc057565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b90805115614fee5767ffffffffffffffff8151602083012092169182600052600d602052614e1f816005604060002001615934565b15614faa57600052600e6020526040600020815167ffffffffffffffff8111613e9557614e4c82546140ad565b601f8111614f78575b506020601f8211600114614ed05791614eaa827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9593614ec095600091614ec5575b506000198260011b9260031b1c19161790565b9055604051918291602083526020830190613d10565b0390a2565b905084015138614e97565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b818110614f60575092614ec09492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610614f47575b5050811b01905561170d565b85015160001960f88460031b161c191690553880614f3b565b9192602060018192868a015181550194019201614f00565b614fa490836000526020600020601f840160051c810191602085106106e457601f0160051c0190614a3e565b38614e55565b5090611f866040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190613d10565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016602082015260208152613fd9604082613f18565b906127109167ffffffffffffffff61506d602083016144a0565b1660009081526010602052604090209161ffff16156150a057606061ffff61509c935460901c1691013561450a565b0490565b606061ffff61509c935460801c1691013561450a565b81519192911561523a576fffffffffffffffffffffffffffffffff6040840151166fffffffffffffffffffffffffffffffff602085015116106151d7576151d591925b805182547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff0000000000000000000000000000000000000000161782556020810151825460409290920151608090811b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff9290921691821760018501557fffffffffffffffffffffffff0000000000000000000000000000000000000000909216174290911b73ffffffff0000000000000000000000000000000016179055565b565b606483615238604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff604084015116158015906152c9575b615268576151d591926150f9565b606483615238604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff602084015116151561525a565b6001600160a01b031680159182156153ad57907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60208361532d6000956002546154c6565b6002555b8060025403600255604051908152a361534657565b7f00000000000000000000000000000000000000000000000000000000000000008061536f5750565b60025481811161537d575050565b7fea0582460000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b8160005260006020526040600020548181106153fd5760208284937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef936000968752868452036040862055615331565b827fe450d38c0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60206001600160a01b0360009361546b866002546154c6565b6002551693841584146154b15780600254036002555b604051908152a37f00000000000000000000000000000000000000000000000000000000000000008061536f5750565b84845283825260408420818154019055615481565b9190820180921161451d57565b9167ffffffffffffffff7f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c92169283600052600d60205261551c81836002604060002001615989565b604080516001600160a01b03909216825260208201929092529081908101614ec0565b91909167ffffffffffffffff83169283600052600a60205260ff60406000205460a01c16156155a45750907f63335ad9e238acd0e9e6c1c20f529ffbea4cda73578c329a7aa7e9d61e5cdcc59183600052600a60205261551c81836040600020615989565b906151d593506154d3565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe182360301811215613d8557016020813591019167ffffffffffffffff8211613d85578136038313613d8557565b9167ffffffffffffffff7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894492169283600052600d60205261551c81836040600020615989565b91909167ffffffffffffffff83169283600052600960205260ff60406000205460a01c16156156aa5750907f996b829383cc7e136842d4c4c175083bcf4e20807c7432105c1b794ba885e7769183600052600960205261551c81836040600020615989565b906151d593506155ff565b906040519182815491828252602082019060005260206000209260005b8181106156e75750506151d592500383613f18565b84548352600194850194879450602090930192016156d2565b80548210156148cc5760005260206000200190600090565b6000818152600c6020526040902054801561581157600019810181811161451d57600b5490600019820191821161451d578181036157c0575b505050600b548015615791576000190161576c81600b615700565b60001982549160031b1b19169055600b55600052600c60205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6157f96157d16157e293600b615700565b90549060031b1c928392600b615700565b81939154906000199060031b92831b921b19161790565b9055600052600c602052604060002055388080615751565b5050600090565b90600182019181600052826020526040600020548015156000146158cb57600019810181811161451d57825490600019820191821161451d57818103615894575b505050805480156157915760001901906158738282615700565b60001982549160031b1b191690555560005260205260006040812055600190565b6158b46158a46157e29386615700565b90549060031b1c92839286615700565b905560005283602052604060002055388080615859565b50505050600090565b80600052600c6020526040600020541560001461592e57600b5468010000000000000000811015613e95576159156157e2826001859401600b55600b615700565b9055600b5490600052600c602052604060002055600190565b50600090565b60008281526001820160205260409020546158115780549068010000000000000000821015613e9557826159726157e2846001809601855584615700565b905580549260005201602052604060002055600190565b9182549060ff8260a01c16158015615bda575b615bd4576fffffffffffffffffffffffffffffffff821691600185019081546159e163ffffffff6fffffffffffffffffffffffffffffffff83169360801c1642614585565b9081615b36575b5050848110615af75750838310615a42575050615a176fffffffffffffffffffffffffffffffff928392614585565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b9190915460801c928315615ab65781615a5a91614585565b9260001981019080821161451d57615a7d615a82926001600160a01b03966154c6565b61454c565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b6001600160a01b0383837fd0c8d23a000000000000000000000000000000000000000000000000000000006000526000196004526024521660445260646000fd5b82856001600160a01b03927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611615baa57615b5192614bc49160801c9061450a565b80841015615ba55750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff00000000000000000000000000000000161786559238806159e8565b615b5c565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b50821561599c56fea164736f6c634300081a000a",
}

var CrossChainPoolTokenABI = CrossChainPoolTokenMetaData.ABI

var CrossChainPoolTokenBin = CrossChainPoolTokenMetaData.Bin

func DeployCrossChainPoolToken(auth *bind.TransactOpts, backend bind.ContractBackend, tokenParams BaseERC20ConstructorParams, advancedPoolHooks common.Address, rmnProxy common.Address, router common.Address) (common.Address, *types.Transaction, *CrossChainPoolToken, error) {
	parsed, err := CrossChainPoolTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrossChainPoolTokenBin), backend, tokenParams, advancedPoolHooks, rmnProxy, router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrossChainPoolToken{address: address, abi: *parsed, CrossChainPoolTokenCaller: CrossChainPoolTokenCaller{contract: contract}, CrossChainPoolTokenTransactor: CrossChainPoolTokenTransactor{contract: contract}, CrossChainPoolTokenFilterer: CrossChainPoolTokenFilterer{contract: contract}}, nil
}

type CrossChainPoolToken struct {
	address common.Address
	abi     abi.ABI
	CrossChainPoolTokenCaller
	CrossChainPoolTokenTransactor
	CrossChainPoolTokenFilterer
}

type CrossChainPoolTokenCaller struct {
	contract *bind.BoundContract
}

type CrossChainPoolTokenTransactor struct {
	contract *bind.BoundContract
}

type CrossChainPoolTokenFilterer struct {
	contract *bind.BoundContract
}

type CrossChainPoolTokenSession struct {
	Contract     *CrossChainPoolToken
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CrossChainPoolTokenCallerSession struct {
	Contract *CrossChainPoolTokenCaller
	CallOpts bind.CallOpts
}

type CrossChainPoolTokenTransactorSession struct {
	Contract     *CrossChainPoolTokenTransactor
	TransactOpts bind.TransactOpts
}

type CrossChainPoolTokenRaw struct {
	Contract *CrossChainPoolToken
}

type CrossChainPoolTokenCallerRaw struct {
	Contract *CrossChainPoolTokenCaller
}

type CrossChainPoolTokenTransactorRaw struct {
	Contract *CrossChainPoolTokenTransactor
}

func NewCrossChainPoolToken(address common.Address, backend bind.ContractBackend) (*CrossChainPoolToken, error) {
	abi, err := abi.JSON(strings.NewReader(CrossChainPoolTokenABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCrossChainPoolToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossChainPoolToken{address: address, abi: abi, CrossChainPoolTokenCaller: CrossChainPoolTokenCaller{contract: contract}, CrossChainPoolTokenTransactor: CrossChainPoolTokenTransactor{contract: contract}, CrossChainPoolTokenFilterer: CrossChainPoolTokenFilterer{contract: contract}}, nil
}

func NewCrossChainPoolTokenCaller(address common.Address, caller bind.ContractCaller) (*CrossChainPoolTokenCaller, error) {
	contract, err := bindCrossChainPoolToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossChainPoolTokenCaller{contract: contract}, nil
}

func NewCrossChainPoolTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossChainPoolTokenTransactor, error) {
	contract, err := bindCrossChainPoolToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossChainPoolTokenTransactor{contract: contract}, nil
}

func NewCrossChainPoolTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossChainPoolTokenFilterer, error) {
	contract, err := bindCrossChainPoolToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossChainPoolTokenFilterer{contract: contract}, nil
}

func bindCrossChainPoolToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrossChainPoolTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_CrossChainPoolToken *CrossChainPoolTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossChainPoolToken.Contract.CrossChainPoolTokenCaller.contract.Call(opts, result, method, params...)
}

func (_CrossChainPoolToken *CrossChainPoolTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.CrossChainPoolTokenTransactor.contract.Transfer(opts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.CrossChainPoolTokenTransactor.contract.Transact(opts, method, params...)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossChainPoolToken.Contract.contract.Call(opts, result, method, params...)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.contract.Transfer(opts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.contract.Transact(opts, method, params...)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CrossChainPoolToken.Contract.Allowance(&_CrossChainPoolToken.CallOpts, owner, spender)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CrossChainPoolToken.Contract.Allowance(&_CrossChainPoolToken.CallOpts, owner, spender)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CrossChainPoolToken.Contract.BalanceOf(&_CrossChainPoolToken.CallOpts, account)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CrossChainPoolToken.Contract.BalanceOf(&_CrossChainPoolToken.CallOpts, account)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) Decimals() (uint8, error) {
	return _CrossChainPoolToken.Contract.Decimals(&_CrossChainPoolToken.CallOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) Decimals() (uint8, error) {
	return _CrossChainPoolToken.Contract.Decimals(&_CrossChainPoolToken.CallOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) GetAdvancedPoolHooks(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "getAdvancedPoolHooks")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) GetAdvancedPoolHooks() (common.Address, error) {
	return _CrossChainPoolToken.Contract.GetAdvancedPoolHooks(&_CrossChainPoolToken.CallOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) GetAdvancedPoolHooks() (common.Address, error) {
	return _CrossChainPoolToken.Contract.GetAdvancedPoolHooks(&_CrossChainPoolToken.CallOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) GetCCIPAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "getCCIPAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) GetCCIPAdmin() (common.Address, error) {
	return _CrossChainPoolToken.Contract.GetCCIPAdmin(&_CrossChainPoolToken.CallOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) GetCCIPAdmin() (common.Address, error) {
	return _CrossChainPoolToken.Contract.GetCCIPAdmin(&_CrossChainPoolToken.CallOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64, customBlockConfirmations bool) (GetCurrentRateLimiterState,

	error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "getCurrentRateLimiterState", remoteChainSelector, customBlockConfirmations)

	outstruct := new(GetCurrentRateLimiterState)
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutboundRateLimiterState = *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)
	outstruct.InboundRateLimiterState = *abi.ConvertType(out[1], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return *outstruct, err

}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) GetCurrentRateLimiterState(remoteChainSelector uint64, customBlockConfirmations bool) (GetCurrentRateLimiterState,

	error) {
	return _CrossChainPoolToken.Contract.GetCurrentRateLimiterState(&_CrossChainPoolToken.CallOpts, remoteChainSelector, customBlockConfirmations)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) GetCurrentRateLimiterState(remoteChainSelector uint64, customBlockConfirmations bool) (GetCurrentRateLimiterState,

	error) {
	return _CrossChainPoolToken.Contract.GetCurrentRateLimiterState(&_CrossChainPoolToken.CallOpts, remoteChainSelector, customBlockConfirmations)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

	error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "getDynamicConfig")

	outstruct := new(GetDynamicConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Router = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RateLimitAdmin = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.FeeAdmin = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _CrossChainPoolToken.Contract.GetDynamicConfig(&_CrossChainPoolToken.CallOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _CrossChainPoolToken.Contract.GetDynamicConfig(&_CrossChainPoolToken.CallOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationsRequested uint16, arg5 []byte) (GetFee,

	error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "getFee", arg0, destChainSelector, arg2, arg3, blockConfirmationsRequested, arg5)

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

func (_CrossChainPoolToken *CrossChainPoolTokenSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationsRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _CrossChainPoolToken.Contract.GetFee(&_CrossChainPoolToken.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationsRequested, arg5)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationsRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _CrossChainPoolToken.Contract.GetFee(&_CrossChainPoolToken.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationsRequested, arg5)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) GetMinBlockConfirmations(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "getMinBlockConfirmations")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) GetMinBlockConfirmations() (uint16, error) {
	return _CrossChainPoolToken.Contract.GetMinBlockConfirmations(&_CrossChainPoolToken.CallOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) GetMinBlockConfirmations() (uint16, error) {
	return _CrossChainPoolToken.Contract.GetMinBlockConfirmations(&_CrossChainPoolToken.CallOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "getRemotePools", remoteChainSelector)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _CrossChainPoolToken.Contract.GetRemotePools(&_CrossChainPoolToken.CallOpts, remoteChainSelector)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _CrossChainPoolToken.Contract.GetRemotePools(&_CrossChainPoolToken.CallOpts, remoteChainSelector)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "getRemoteToken", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _CrossChainPoolToken.Contract.GetRemoteToken(&_CrossChainPoolToken.CallOpts, remoteChainSelector)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _CrossChainPoolToken.Contract.GetRemoteToken(&_CrossChainPoolToken.CallOpts, remoteChainSelector)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) GetRequiredCCVs(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, sourceDenominatedAmount *big.Int, blockConfirmationsRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "getRequiredCCVs", localToken, remoteChainSelector, sourceDenominatedAmount, blockConfirmationsRequested, extraData, direction)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, sourceDenominatedAmount *big.Int, blockConfirmationsRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _CrossChainPoolToken.Contract.GetRequiredCCVs(&_CrossChainPoolToken.CallOpts, localToken, remoteChainSelector, sourceDenominatedAmount, blockConfirmationsRequested, extraData, direction)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, sourceDenominatedAmount *big.Int, blockConfirmationsRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _CrossChainPoolToken.Contract.GetRequiredCCVs(&_CrossChainPoolToken.CallOpts, localToken, remoteChainSelector, sourceDenominatedAmount, blockConfirmationsRequested, extraData, direction)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) GetRmnProxy() (common.Address, error) {
	return _CrossChainPoolToken.Contract.GetRmnProxy(&_CrossChainPoolToken.CallOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) GetRmnProxy() (common.Address, error) {
	return _CrossChainPoolToken.Contract.GetRmnProxy(&_CrossChainPoolToken.CallOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) GetSupportedChains() ([]uint64, error) {
	return _CrossChainPoolToken.Contract.GetSupportedChains(&_CrossChainPoolToken.CallOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) GetSupportedChains() ([]uint64, error) {
	return _CrossChainPoolToken.Contract.GetSupportedChains(&_CrossChainPoolToken.CallOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) GetToken() (common.Address, error) {
	return _CrossChainPoolToken.Contract.GetToken(&_CrossChainPoolToken.CallOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) GetToken() (common.Address, error) {
	return _CrossChainPoolToken.Contract.GetToken(&_CrossChainPoolToken.CallOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) GetTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "getTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) GetTokenDecimals() (uint8, error) {
	return _CrossChainPoolToken.Contract.GetTokenDecimals(&_CrossChainPoolToken.CallOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) GetTokenDecimals() (uint8, error) {
	return _CrossChainPoolToken.Contract.GetTokenDecimals(&_CrossChainPoolToken.CallOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "getTokenTransferFeeConfig", arg0, destChainSelector, arg2, arg3)

	if err != nil {
		return *new(IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolV2TokenTransferFeeConfig)).(*IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _CrossChainPoolToken.Contract.GetTokenTransferFeeConfig(&_CrossChainPoolToken.CallOpts, arg0, destChainSelector, arg2, arg3)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _CrossChainPoolToken.Contract.GetTokenTransferFeeConfig(&_CrossChainPoolToken.CallOpts, arg0, destChainSelector, arg2, arg3)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "isRemotePool", remoteChainSelector, remotePoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _CrossChainPoolToken.Contract.IsRemotePool(&_CrossChainPoolToken.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _CrossChainPoolToken.Contract.IsRemotePool(&_CrossChainPoolToken.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _CrossChainPoolToken.Contract.IsSupportedChain(&_CrossChainPoolToken.CallOpts, remoteChainSelector)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _CrossChainPoolToken.Contract.IsSupportedChain(&_CrossChainPoolToken.CallOpts, remoteChainSelector)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) IsSupportedToken(token common.Address) (bool, error) {
	return _CrossChainPoolToken.Contract.IsSupportedToken(&_CrossChainPoolToken.CallOpts, token)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _CrossChainPoolToken.Contract.IsSupportedToken(&_CrossChainPoolToken.CallOpts, token)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) MaxSupply() (*big.Int, error) {
	return _CrossChainPoolToken.Contract.MaxSupply(&_CrossChainPoolToken.CallOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) MaxSupply() (*big.Int, error) {
	return _CrossChainPoolToken.Contract.MaxSupply(&_CrossChainPoolToken.CallOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) Name() (string, error) {
	return _CrossChainPoolToken.Contract.Name(&_CrossChainPoolToken.CallOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) Name() (string, error) {
	return _CrossChainPoolToken.Contract.Name(&_CrossChainPoolToken.CallOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) Owner() (common.Address, error) {
	return _CrossChainPoolToken.Contract.Owner(&_CrossChainPoolToken.CallOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) Owner() (common.Address, error) {
	return _CrossChainPoolToken.Contract.Owner(&_CrossChainPoolToken.CallOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CrossChainPoolToken.Contract.SupportsInterface(&_CrossChainPoolToken.CallOpts, interfaceId)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CrossChainPoolToken.Contract.SupportsInterface(&_CrossChainPoolToken.CallOpts, interfaceId)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) Symbol() (string, error) {
	return _CrossChainPoolToken.Contract.Symbol(&_CrossChainPoolToken.CallOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) Symbol() (string, error) {
	return _CrossChainPoolToken.Contract.Symbol(&_CrossChainPoolToken.CallOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) TotalSupply() (*big.Int, error) {
	return _CrossChainPoolToken.Contract.TotalSupply(&_CrossChainPoolToken.CallOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _CrossChainPoolToken.Contract.TotalSupply(&_CrossChainPoolToken.CallOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) TypeAndVersion() (string, error) {
	return _CrossChainPoolToken.Contract.TypeAndVersion(&_CrossChainPoolToken.CallOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) TypeAndVersion() (string, error) {
	return _CrossChainPoolToken.Contract.TypeAndVersion(&_CrossChainPoolToken.CallOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossChainPoolToken.contract.Transact(opts, "acceptOwnership")
}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) AcceptOwnership() (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.AcceptOwnership(&_CrossChainPoolToken.TransactOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.AcceptOwnership(&_CrossChainPoolToken.TransactOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactor) AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _CrossChainPoolToken.contract.Transact(opts, "addRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.AddRemotePool(&_CrossChainPoolToken.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactorSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.AddRemotePool(&_CrossChainPoolToken.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactor) ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _CrossChainPoolToken.contract.Transact(opts, "applyChainUpdates", remoteChainSelectorsToRemove, chainsToAdd)
}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.ApplyChainUpdates(&_CrossChainPoolToken.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactorSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.ApplyChainUpdates(&_CrossChainPoolToken.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactor) ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _CrossChainPoolToken.contract.Transact(opts, "applyTokenTransferFeeConfigUpdates", tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.ApplyTokenTransferFeeConfigUpdates(&_CrossChainPoolToken.TransactOpts, tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactorSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.ApplyTokenTransferFeeConfigUpdates(&_CrossChainPoolToken.TransactOpts, tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossChainPoolToken.contract.Transact(opts, "approve", spender, value)
}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.Approve(&_CrossChainPoolToken.TransactOpts, spender, value)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.Approve(&_CrossChainPoolToken.TransactOpts, spender, value)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _CrossChainPoolToken.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.LockOrBurn(&_CrossChainPoolToken.TransactOpts, lockOrBurnIn)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.LockOrBurn(&_CrossChainPoolToken.TransactOpts, lockOrBurnIn)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationsRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _CrossChainPoolToken.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, blockConfirmationsRequested, tokenArgs)
}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationsRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.LockOrBurn0(&_CrossChainPoolToken.TransactOpts, lockOrBurnIn, blockConfirmationsRequested, tokenArgs)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationsRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.LockOrBurn0(&_CrossChainPoolToken.TransactOpts, lockOrBurnIn, blockConfirmationsRequested, tokenArgs)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _CrossChainPoolToken.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.ReleaseOrMint(&_CrossChainPoolToken.TransactOpts, releaseOrMintIn)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.ReleaseOrMint(&_CrossChainPoolToken.TransactOpts, releaseOrMintIn)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationsRequested uint16) (*types.Transaction, error) {
	return _CrossChainPoolToken.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn, blockConfirmationsRequested)
}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationsRequested uint16) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.ReleaseOrMint0(&_CrossChainPoolToken.TransactOpts, releaseOrMintIn, blockConfirmationsRequested)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationsRequested uint16) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.ReleaseOrMint0(&_CrossChainPoolToken.TransactOpts, releaseOrMintIn, blockConfirmationsRequested)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactor) RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _CrossChainPoolToken.contract.Transact(opts, "removeRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.RemoveRemotePool(&_CrossChainPoolToken.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactorSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.RemoveRemotePool(&_CrossChainPoolToken.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactor) SetCCIPAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _CrossChainPoolToken.contract.Transact(opts, "setCCIPAdmin", newAdmin)
}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) SetCCIPAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.SetCCIPAdmin(&_CrossChainPoolToken.TransactOpts, newAdmin)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactorSession) SetCCIPAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.SetCCIPAdmin(&_CrossChainPoolToken.TransactOpts, newAdmin)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactor) SetDynamicConfig(opts *bind.TransactOpts, router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error) {
	return _CrossChainPoolToken.contract.Transact(opts, "setDynamicConfig", router, rateLimitAdmin, feeAdmin)
}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) SetDynamicConfig(router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.SetDynamicConfig(&_CrossChainPoolToken.TransactOpts, router, rateLimitAdmin, feeAdmin)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactorSession) SetDynamicConfig(router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.SetDynamicConfig(&_CrossChainPoolToken.TransactOpts, router, rateLimitAdmin, feeAdmin)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactor) SetMinBlockConfirmations(opts *bind.TransactOpts, minBlockConfirmations uint16) (*types.Transaction, error) {
	return _CrossChainPoolToken.contract.Transact(opts, "setMinBlockConfirmations", minBlockConfirmations)
}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) SetMinBlockConfirmations(minBlockConfirmations uint16) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.SetMinBlockConfirmations(&_CrossChainPoolToken.TransactOpts, minBlockConfirmations)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactorSession) SetMinBlockConfirmations(minBlockConfirmations uint16) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.SetMinBlockConfirmations(&_CrossChainPoolToken.TransactOpts, minBlockConfirmations)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactor) SetRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _CrossChainPoolToken.contract.Transact(opts, "setRateLimitConfig", rateLimitConfigArgs)
}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) SetRateLimitConfig(rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.SetRateLimitConfig(&_CrossChainPoolToken.TransactOpts, rateLimitConfigArgs)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactorSession) SetRateLimitConfig(rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.SetRateLimitConfig(&_CrossChainPoolToken.TransactOpts, rateLimitConfigArgs)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossChainPoolToken.contract.Transact(opts, "transfer", to, value)
}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.Transfer(&_CrossChainPoolToken.TransactOpts, to, value)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.Transfer(&_CrossChainPoolToken.TransactOpts, to, value)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossChainPoolToken.contract.Transact(opts, "transferFrom", from, to, value)
}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.TransferFrom(&_CrossChainPoolToken.TransactOpts, from, to, value)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.TransferFrom(&_CrossChainPoolToken.TransactOpts, from, to, value)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CrossChainPoolToken.contract.Transact(opts, "transferOwnership", to)
}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.TransferOwnership(&_CrossChainPoolToken.TransactOpts, to)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.TransferOwnership(&_CrossChainPoolToken.TransactOpts, to)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactor) UpdateAdvancedPoolHooks(opts *bind.TransactOpts, newHook common.Address) (*types.Transaction, error) {
	return _CrossChainPoolToken.contract.Transact(opts, "updateAdvancedPoolHooks", newHook)
}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) UpdateAdvancedPoolHooks(newHook common.Address) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.UpdateAdvancedPoolHooks(&_CrossChainPoolToken.TransactOpts, newHook)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactorSession) UpdateAdvancedPoolHooks(newHook common.Address) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.UpdateAdvancedPoolHooks(&_CrossChainPoolToken.TransactOpts, newHook)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _CrossChainPoolToken.contract.Transact(opts, "withdrawFeeTokens", feeTokens, recipient)
}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.WithdrawFeeTokens(&_CrossChainPoolToken.TransactOpts, feeTokens, recipient)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactorSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.WithdrawFeeTokens(&_CrossChainPoolToken.TransactOpts, feeTokens, recipient)
}

type CrossChainPoolTokenAdvancedPoolHooksUpdatedIterator struct {
	Event *CrossChainPoolTokenAdvancedPoolHooksUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainPoolTokenAdvancedPoolHooksUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainPoolTokenAdvancedPoolHooksUpdated)
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
		it.Event = new(CrossChainPoolTokenAdvancedPoolHooksUpdated)
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

func (it *CrossChainPoolTokenAdvancedPoolHooksUpdatedIterator) Error() error {
	return it.fail
}

func (it *CrossChainPoolTokenAdvancedPoolHooksUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainPoolTokenAdvancedPoolHooksUpdated struct {
	OldHook common.Address
	NewHook common.Address
	Raw     types.Log
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) FilterAdvancedPoolHooksUpdated(opts *bind.FilterOpts) (*CrossChainPoolTokenAdvancedPoolHooksUpdatedIterator, error) {

	logs, sub, err := _CrossChainPoolToken.contract.FilterLogs(opts, "AdvancedPoolHooksUpdated")
	if err != nil {
		return nil, err
	}
	return &CrossChainPoolTokenAdvancedPoolHooksUpdatedIterator{contract: _CrossChainPoolToken.contract, event: "AdvancedPoolHooksUpdated", logs: logs, sub: sub}, nil
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) WatchAdvancedPoolHooksUpdated(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenAdvancedPoolHooksUpdated) (event.Subscription, error) {

	logs, sub, err := _CrossChainPoolToken.contract.WatchLogs(opts, "AdvancedPoolHooksUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainPoolTokenAdvancedPoolHooksUpdated)
				if err := _CrossChainPoolToken.contract.UnpackLog(event, "AdvancedPoolHooksUpdated", log); err != nil {
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

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) ParseAdvancedPoolHooksUpdated(log types.Log) (*CrossChainPoolTokenAdvancedPoolHooksUpdated, error) {
	event := new(CrossChainPoolTokenAdvancedPoolHooksUpdated)
	if err := _CrossChainPoolToken.contract.UnpackLog(event, "AdvancedPoolHooksUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossChainPoolTokenApprovalIterator struct {
	Event *CrossChainPoolTokenApproval

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainPoolTokenApprovalIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainPoolTokenApproval)
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
		it.Event = new(CrossChainPoolTokenApproval)
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

func (it *CrossChainPoolTokenApprovalIterator) Error() error {
	return it.fail
}

func (it *CrossChainPoolTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainPoolTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CrossChainPoolTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainPoolTokenApprovalIterator{contract: _CrossChainPoolToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainPoolTokenApproval)
				if err := _CrossChainPoolToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) ParseApproval(log types.Log) (*CrossChainPoolTokenApproval, error) {
	event := new(CrossChainPoolTokenApproval)
	if err := _CrossChainPoolToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossChainPoolTokenCCIPAdminTransferredIterator struct {
	Event *CrossChainPoolTokenCCIPAdminTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainPoolTokenCCIPAdminTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainPoolTokenCCIPAdminTransferred)
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
		it.Event = new(CrossChainPoolTokenCCIPAdminTransferred)
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

func (it *CrossChainPoolTokenCCIPAdminTransferredIterator) Error() error {
	return it.fail
}

func (it *CrossChainPoolTokenCCIPAdminTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainPoolTokenCCIPAdminTransferred struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) FilterCCIPAdminTransferred(opts *bind.FilterOpts, previousAdmin []common.Address, newAdmin []common.Address) (*CrossChainPoolTokenCCIPAdminTransferredIterator, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.FilterLogs(opts, "CCIPAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainPoolTokenCCIPAdminTransferredIterator{contract: _CrossChainPoolToken.contract, event: "CCIPAdminTransferred", logs: logs, sub: sub}, nil
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) WatchCCIPAdminTransferred(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenCCIPAdminTransferred, previousAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.WatchLogs(opts, "CCIPAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainPoolTokenCCIPAdminTransferred)
				if err := _CrossChainPoolToken.contract.UnpackLog(event, "CCIPAdminTransferred", log); err != nil {
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

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) ParseCCIPAdminTransferred(log types.Log) (*CrossChainPoolTokenCCIPAdminTransferred, error) {
	event := new(CrossChainPoolTokenCCIPAdminTransferred)
	if err := _CrossChainPoolToken.contract.UnpackLog(event, "CCIPAdminTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossChainPoolTokenChainAddedIterator struct {
	Event *CrossChainPoolTokenChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainPoolTokenChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainPoolTokenChainAdded)
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
		it.Event = new(CrossChainPoolTokenChainAdded)
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

func (it *CrossChainPoolTokenChainAddedIterator) Error() error {
	return it.fail
}

func (it *CrossChainPoolTokenChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainPoolTokenChainAdded struct {
	RemoteChainSelector       uint64
	RemoteToken               []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) FilterChainAdded(opts *bind.FilterOpts) (*CrossChainPoolTokenChainAddedIterator, error) {

	logs, sub, err := _CrossChainPoolToken.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &CrossChainPoolTokenChainAddedIterator{contract: _CrossChainPoolToken.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenChainAdded) (event.Subscription, error) {

	logs, sub, err := _CrossChainPoolToken.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainPoolTokenChainAdded)
				if err := _CrossChainPoolToken.contract.UnpackLog(event, "ChainAdded", log); err != nil {
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

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) ParseChainAdded(log types.Log) (*CrossChainPoolTokenChainAdded, error) {
	event := new(CrossChainPoolTokenChainAdded)
	if err := _CrossChainPoolToken.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossChainPoolTokenChainRemovedIterator struct {
	Event *CrossChainPoolTokenChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainPoolTokenChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainPoolTokenChainRemoved)
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
		it.Event = new(CrossChainPoolTokenChainRemoved)
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

func (it *CrossChainPoolTokenChainRemovedIterator) Error() error {
	return it.fail
}

func (it *CrossChainPoolTokenChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainPoolTokenChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) FilterChainRemoved(opts *bind.FilterOpts) (*CrossChainPoolTokenChainRemovedIterator, error) {

	logs, sub, err := _CrossChainPoolToken.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &CrossChainPoolTokenChainRemovedIterator{contract: _CrossChainPoolToken.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenChainRemoved) (event.Subscription, error) {

	logs, sub, err := _CrossChainPoolToken.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainPoolTokenChainRemoved)
				if err := _CrossChainPoolToken.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
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

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) ParseChainRemoved(log types.Log) (*CrossChainPoolTokenChainRemoved, error) {
	event := new(CrossChainPoolTokenChainRemoved)
	if err := _CrossChainPoolToken.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossChainPoolTokenCustomBlockConfirmationsInboundRateLimitConsumedIterator struct {
	Event *CrossChainPoolTokenCustomBlockConfirmationsInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainPoolTokenCustomBlockConfirmationsInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainPoolTokenCustomBlockConfirmationsInboundRateLimitConsumed)
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
		it.Event = new(CrossChainPoolTokenCustomBlockConfirmationsInboundRateLimitConsumed)
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

func (it *CrossChainPoolTokenCustomBlockConfirmationsInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *CrossChainPoolTokenCustomBlockConfirmationsInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainPoolTokenCustomBlockConfirmationsInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) FilterCustomBlockConfirmationsInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CrossChainPoolTokenCustomBlockConfirmationsInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.FilterLogs(opts, "CustomBlockConfirmationsInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainPoolTokenCustomBlockConfirmationsInboundRateLimitConsumedIterator{contract: _CrossChainPoolToken.contract, event: "CustomBlockConfirmationsInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) WatchCustomBlockConfirmationsInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenCustomBlockConfirmationsInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.WatchLogs(opts, "CustomBlockConfirmationsInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainPoolTokenCustomBlockConfirmationsInboundRateLimitConsumed)
				if err := _CrossChainPoolToken.contract.UnpackLog(event, "CustomBlockConfirmationsInboundRateLimitConsumed", log); err != nil {
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

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) ParseCustomBlockConfirmationsInboundRateLimitConsumed(log types.Log) (*CrossChainPoolTokenCustomBlockConfirmationsInboundRateLimitConsumed, error) {
	event := new(CrossChainPoolTokenCustomBlockConfirmationsInboundRateLimitConsumed)
	if err := _CrossChainPoolToken.contract.UnpackLog(event, "CustomBlockConfirmationsInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossChainPoolTokenCustomBlockConfirmationsOutboundRateLimitConsumedIterator struct {
	Event *CrossChainPoolTokenCustomBlockConfirmationsOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainPoolTokenCustomBlockConfirmationsOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainPoolTokenCustomBlockConfirmationsOutboundRateLimitConsumed)
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
		it.Event = new(CrossChainPoolTokenCustomBlockConfirmationsOutboundRateLimitConsumed)
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

func (it *CrossChainPoolTokenCustomBlockConfirmationsOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *CrossChainPoolTokenCustomBlockConfirmationsOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainPoolTokenCustomBlockConfirmationsOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) FilterCustomBlockConfirmationsOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CrossChainPoolTokenCustomBlockConfirmationsOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.FilterLogs(opts, "CustomBlockConfirmationsOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainPoolTokenCustomBlockConfirmationsOutboundRateLimitConsumedIterator{contract: _CrossChainPoolToken.contract, event: "CustomBlockConfirmationsOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) WatchCustomBlockConfirmationsOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenCustomBlockConfirmationsOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.WatchLogs(opts, "CustomBlockConfirmationsOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainPoolTokenCustomBlockConfirmationsOutboundRateLimitConsumed)
				if err := _CrossChainPoolToken.contract.UnpackLog(event, "CustomBlockConfirmationsOutboundRateLimitConsumed", log); err != nil {
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

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) ParseCustomBlockConfirmationsOutboundRateLimitConsumed(log types.Log) (*CrossChainPoolTokenCustomBlockConfirmationsOutboundRateLimitConsumed, error) {
	event := new(CrossChainPoolTokenCustomBlockConfirmationsOutboundRateLimitConsumed)
	if err := _CrossChainPoolToken.contract.UnpackLog(event, "CustomBlockConfirmationsOutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossChainPoolTokenDynamicConfigSetIterator struct {
	Event *CrossChainPoolTokenDynamicConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainPoolTokenDynamicConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainPoolTokenDynamicConfigSet)
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
		it.Event = new(CrossChainPoolTokenDynamicConfigSet)
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

func (it *CrossChainPoolTokenDynamicConfigSetIterator) Error() error {
	return it.fail
}

func (it *CrossChainPoolTokenDynamicConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainPoolTokenDynamicConfigSet struct {
	Router         common.Address
	RateLimitAdmin common.Address
	FeeAdmin       common.Address
	Raw            types.Log
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) FilterDynamicConfigSet(opts *bind.FilterOpts) (*CrossChainPoolTokenDynamicConfigSetIterator, error) {

	logs, sub, err := _CrossChainPoolToken.contract.FilterLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return &CrossChainPoolTokenDynamicConfigSetIterator{contract: _CrossChainPoolToken.contract, event: "DynamicConfigSet", logs: logs, sub: sub}, nil
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenDynamicConfigSet) (event.Subscription, error) {

	logs, sub, err := _CrossChainPoolToken.contract.WatchLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainPoolTokenDynamicConfigSet)
				if err := _CrossChainPoolToken.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
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

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) ParseDynamicConfigSet(log types.Log) (*CrossChainPoolTokenDynamicConfigSet, error) {
	event := new(CrossChainPoolTokenDynamicConfigSet)
	if err := _CrossChainPoolToken.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossChainPoolTokenFeeTokenWithdrawnIterator struct {
	Event *CrossChainPoolTokenFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainPoolTokenFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainPoolTokenFeeTokenWithdrawn)
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
		it.Event = new(CrossChainPoolTokenFeeTokenWithdrawn)
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

func (it *CrossChainPoolTokenFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *CrossChainPoolTokenFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainPoolTokenFeeTokenWithdrawn struct {
	Receiver common.Address
	FeeToken common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*CrossChainPoolTokenFeeTokenWithdrawnIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.FilterLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainPoolTokenFeeTokenWithdrawnIterator{contract: _CrossChainPoolToken.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.WatchLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainPoolTokenFeeTokenWithdrawn)
				if err := _CrossChainPoolToken.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) ParseFeeTokenWithdrawn(log types.Log) (*CrossChainPoolTokenFeeTokenWithdrawn, error) {
	event := new(CrossChainPoolTokenFeeTokenWithdrawn)
	if err := _CrossChainPoolToken.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossChainPoolTokenInboundRateLimitConsumedIterator struct {
	Event *CrossChainPoolTokenInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainPoolTokenInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainPoolTokenInboundRateLimitConsumed)
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
		it.Event = new(CrossChainPoolTokenInboundRateLimitConsumed)
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

func (it *CrossChainPoolTokenInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *CrossChainPoolTokenInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainPoolTokenInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CrossChainPoolTokenInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.FilterLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainPoolTokenInboundRateLimitConsumedIterator{contract: _CrossChainPoolToken.contract, event: "InboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.WatchLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainPoolTokenInboundRateLimitConsumed)
				if err := _CrossChainPoolToken.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
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

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) ParseInboundRateLimitConsumed(log types.Log) (*CrossChainPoolTokenInboundRateLimitConsumed, error) {
	event := new(CrossChainPoolTokenInboundRateLimitConsumed)
	if err := _CrossChainPoolToken.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossChainPoolTokenLockedOrBurnedIterator struct {
	Event *CrossChainPoolTokenLockedOrBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainPoolTokenLockedOrBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainPoolTokenLockedOrBurned)
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
		it.Event = new(CrossChainPoolTokenLockedOrBurned)
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

func (it *CrossChainPoolTokenLockedOrBurnedIterator) Error() error {
	return it.fail
}

func (it *CrossChainPoolTokenLockedOrBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainPoolTokenLockedOrBurned struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CrossChainPoolTokenLockedOrBurnedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.FilterLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainPoolTokenLockedOrBurnedIterator{contract: _CrossChainPoolToken.contract, event: "LockedOrBurned", logs: logs, sub: sub}, nil
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.WatchLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainPoolTokenLockedOrBurned)
				if err := _CrossChainPoolToken.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
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

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) ParseLockedOrBurned(log types.Log) (*CrossChainPoolTokenLockedOrBurned, error) {
	event := new(CrossChainPoolTokenLockedOrBurned)
	if err := _CrossChainPoolToken.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossChainPoolTokenMinBlockConfirmationsSetIterator struct {
	Event *CrossChainPoolTokenMinBlockConfirmationsSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainPoolTokenMinBlockConfirmationsSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainPoolTokenMinBlockConfirmationsSet)
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
		it.Event = new(CrossChainPoolTokenMinBlockConfirmationsSet)
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

func (it *CrossChainPoolTokenMinBlockConfirmationsSetIterator) Error() error {
	return it.fail
}

func (it *CrossChainPoolTokenMinBlockConfirmationsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainPoolTokenMinBlockConfirmationsSet struct {
	MinBlockConfirmations uint16
	Raw                   types.Log
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) FilterMinBlockConfirmationsSet(opts *bind.FilterOpts) (*CrossChainPoolTokenMinBlockConfirmationsSetIterator, error) {

	logs, sub, err := _CrossChainPoolToken.contract.FilterLogs(opts, "MinBlockConfirmationsSet")
	if err != nil {
		return nil, err
	}
	return &CrossChainPoolTokenMinBlockConfirmationsSetIterator{contract: _CrossChainPoolToken.contract, event: "MinBlockConfirmationsSet", logs: logs, sub: sub}, nil
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) WatchMinBlockConfirmationsSet(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenMinBlockConfirmationsSet) (event.Subscription, error) {

	logs, sub, err := _CrossChainPoolToken.contract.WatchLogs(opts, "MinBlockConfirmationsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainPoolTokenMinBlockConfirmationsSet)
				if err := _CrossChainPoolToken.contract.UnpackLog(event, "MinBlockConfirmationsSet", log); err != nil {
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

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) ParseMinBlockConfirmationsSet(log types.Log) (*CrossChainPoolTokenMinBlockConfirmationsSet, error) {
	event := new(CrossChainPoolTokenMinBlockConfirmationsSet)
	if err := _CrossChainPoolToken.contract.UnpackLog(event, "MinBlockConfirmationsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossChainPoolTokenOutboundRateLimitConsumedIterator struct {
	Event *CrossChainPoolTokenOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainPoolTokenOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainPoolTokenOutboundRateLimitConsumed)
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
		it.Event = new(CrossChainPoolTokenOutboundRateLimitConsumed)
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

func (it *CrossChainPoolTokenOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *CrossChainPoolTokenOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainPoolTokenOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CrossChainPoolTokenOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.FilterLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainPoolTokenOutboundRateLimitConsumedIterator{contract: _CrossChainPoolToken.contract, event: "OutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.WatchLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainPoolTokenOutboundRateLimitConsumed)
				if err := _CrossChainPoolToken.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
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

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) ParseOutboundRateLimitConsumed(log types.Log) (*CrossChainPoolTokenOutboundRateLimitConsumed, error) {
	event := new(CrossChainPoolTokenOutboundRateLimitConsumed)
	if err := _CrossChainPoolToken.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossChainPoolTokenOwnershipTransferRequestedIterator struct {
	Event *CrossChainPoolTokenOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainPoolTokenOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainPoolTokenOwnershipTransferRequested)
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
		it.Event = new(CrossChainPoolTokenOwnershipTransferRequested)
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

func (it *CrossChainPoolTokenOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *CrossChainPoolTokenOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainPoolTokenOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CrossChainPoolTokenOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainPoolTokenOwnershipTransferRequestedIterator{contract: _CrossChainPoolToken.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainPoolTokenOwnershipTransferRequested)
				if err := _CrossChainPoolToken.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) ParseOwnershipTransferRequested(log types.Log) (*CrossChainPoolTokenOwnershipTransferRequested, error) {
	event := new(CrossChainPoolTokenOwnershipTransferRequested)
	if err := _CrossChainPoolToken.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossChainPoolTokenOwnershipTransferredIterator struct {
	Event *CrossChainPoolTokenOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainPoolTokenOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainPoolTokenOwnershipTransferred)
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
		it.Event = new(CrossChainPoolTokenOwnershipTransferred)
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

func (it *CrossChainPoolTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *CrossChainPoolTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainPoolTokenOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CrossChainPoolTokenOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainPoolTokenOwnershipTransferredIterator{contract: _CrossChainPoolToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainPoolTokenOwnershipTransferred)
				if err := _CrossChainPoolToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) ParseOwnershipTransferred(log types.Log) (*CrossChainPoolTokenOwnershipTransferred, error) {
	event := new(CrossChainPoolTokenOwnershipTransferred)
	if err := _CrossChainPoolToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossChainPoolTokenRateLimitConfiguredIterator struct {
	Event *CrossChainPoolTokenRateLimitConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainPoolTokenRateLimitConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainPoolTokenRateLimitConfigured)
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
		it.Event = new(CrossChainPoolTokenRateLimitConfigured)
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

func (it *CrossChainPoolTokenRateLimitConfiguredIterator) Error() error {
	return it.fail
}

func (it *CrossChainPoolTokenRateLimitConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainPoolTokenRateLimitConfigured struct {
	RemoteChainSelector       uint64
	CustomBlockConfirmations  bool
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) FilterRateLimitConfigured(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CrossChainPoolTokenRateLimitConfiguredIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.FilterLogs(opts, "RateLimitConfigured", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainPoolTokenRateLimitConfiguredIterator{contract: _CrossChainPoolToken.contract, event: "RateLimitConfigured", logs: logs, sub: sub}, nil
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) WatchRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenRateLimitConfigured, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.WatchLogs(opts, "RateLimitConfigured", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainPoolTokenRateLimitConfigured)
				if err := _CrossChainPoolToken.contract.UnpackLog(event, "RateLimitConfigured", log); err != nil {
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

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) ParseRateLimitConfigured(log types.Log) (*CrossChainPoolTokenRateLimitConfigured, error) {
	event := new(CrossChainPoolTokenRateLimitConfigured)
	if err := _CrossChainPoolToken.contract.UnpackLog(event, "RateLimitConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossChainPoolTokenReleasedOrMintedIterator struct {
	Event *CrossChainPoolTokenReleasedOrMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainPoolTokenReleasedOrMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainPoolTokenReleasedOrMinted)
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
		it.Event = new(CrossChainPoolTokenReleasedOrMinted)
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

func (it *CrossChainPoolTokenReleasedOrMintedIterator) Error() error {
	return it.fail
}

func (it *CrossChainPoolTokenReleasedOrMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainPoolTokenReleasedOrMinted struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Recipient           common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CrossChainPoolTokenReleasedOrMintedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.FilterLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainPoolTokenReleasedOrMintedIterator{contract: _CrossChainPoolToken.contract, event: "ReleasedOrMinted", logs: logs, sub: sub}, nil
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.WatchLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainPoolTokenReleasedOrMinted)
				if err := _CrossChainPoolToken.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
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

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) ParseReleasedOrMinted(log types.Log) (*CrossChainPoolTokenReleasedOrMinted, error) {
	event := new(CrossChainPoolTokenReleasedOrMinted)
	if err := _CrossChainPoolToken.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossChainPoolTokenRemotePoolAddedIterator struct {
	Event *CrossChainPoolTokenRemotePoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainPoolTokenRemotePoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainPoolTokenRemotePoolAdded)
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
		it.Event = new(CrossChainPoolTokenRemotePoolAdded)
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

func (it *CrossChainPoolTokenRemotePoolAddedIterator) Error() error {
	return it.fail
}

func (it *CrossChainPoolTokenRemotePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainPoolTokenRemotePoolAdded struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CrossChainPoolTokenRemotePoolAddedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.FilterLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainPoolTokenRemotePoolAddedIterator{contract: _CrossChainPoolToken.contract, event: "RemotePoolAdded", logs: logs, sub: sub}, nil
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.WatchLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainPoolTokenRemotePoolAdded)
				if err := _CrossChainPoolToken.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
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

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) ParseRemotePoolAdded(log types.Log) (*CrossChainPoolTokenRemotePoolAdded, error) {
	event := new(CrossChainPoolTokenRemotePoolAdded)
	if err := _CrossChainPoolToken.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossChainPoolTokenRemotePoolRemovedIterator struct {
	Event *CrossChainPoolTokenRemotePoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainPoolTokenRemotePoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainPoolTokenRemotePoolRemoved)
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
		it.Event = new(CrossChainPoolTokenRemotePoolRemoved)
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

func (it *CrossChainPoolTokenRemotePoolRemovedIterator) Error() error {
	return it.fail
}

func (it *CrossChainPoolTokenRemotePoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainPoolTokenRemotePoolRemoved struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CrossChainPoolTokenRemotePoolRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.FilterLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainPoolTokenRemotePoolRemovedIterator{contract: _CrossChainPoolToken.contract, event: "RemotePoolRemoved", logs: logs, sub: sub}, nil
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.WatchLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainPoolTokenRemotePoolRemoved)
				if err := _CrossChainPoolToken.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
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

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) ParseRemotePoolRemoved(log types.Log) (*CrossChainPoolTokenRemotePoolRemoved, error) {
	event := new(CrossChainPoolTokenRemotePoolRemoved)
	if err := _CrossChainPoolToken.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossChainPoolTokenTokenTransferFeeConfigDeletedIterator struct {
	Event *CrossChainPoolTokenTokenTransferFeeConfigDeleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainPoolTokenTokenTransferFeeConfigDeletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainPoolTokenTokenTransferFeeConfigDeleted)
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
		it.Event = new(CrossChainPoolTokenTokenTransferFeeConfigDeleted)
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

func (it *CrossChainPoolTokenTokenTransferFeeConfigDeletedIterator) Error() error {
	return it.fail
}

func (it *CrossChainPoolTokenTokenTransferFeeConfigDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainPoolTokenTokenTransferFeeConfigDeleted struct {
	DestChainSelector uint64
	Raw               types.Log
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*CrossChainPoolTokenTokenTransferFeeConfigDeletedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.FilterLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainPoolTokenTokenTransferFeeConfigDeletedIterator{contract: _CrossChainPoolToken.contract, event: "TokenTransferFeeConfigDeleted", logs: logs, sub: sub}, nil
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.WatchLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainPoolTokenTokenTransferFeeConfigDeleted)
				if err := _CrossChainPoolToken.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
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

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) ParseTokenTransferFeeConfigDeleted(log types.Log) (*CrossChainPoolTokenTokenTransferFeeConfigDeleted, error) {
	event := new(CrossChainPoolTokenTokenTransferFeeConfigDeleted)
	if err := _CrossChainPoolToken.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossChainPoolTokenTokenTransferFeeConfigUpdatedIterator struct {
	Event *CrossChainPoolTokenTokenTransferFeeConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainPoolTokenTokenTransferFeeConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainPoolTokenTokenTransferFeeConfigUpdated)
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
		it.Event = new(CrossChainPoolTokenTokenTransferFeeConfigUpdated)
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

func (it *CrossChainPoolTokenTokenTransferFeeConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *CrossChainPoolTokenTokenTransferFeeConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainPoolTokenTokenTransferFeeConfigUpdated struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
	Raw                    types.Log
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*CrossChainPoolTokenTokenTransferFeeConfigUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.FilterLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainPoolTokenTokenTransferFeeConfigUpdatedIterator{contract: _CrossChainPoolToken.contract, event: "TokenTransferFeeConfigUpdated", logs: logs, sub: sub}, nil
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.WatchLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainPoolTokenTokenTransferFeeConfigUpdated)
				if err := _CrossChainPoolToken.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
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

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) ParseTokenTransferFeeConfigUpdated(log types.Log) (*CrossChainPoolTokenTokenTransferFeeConfigUpdated, error) {
	event := new(CrossChainPoolTokenTokenTransferFeeConfigUpdated)
	if err := _CrossChainPoolToken.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossChainPoolTokenTransferIterator struct {
	Event *CrossChainPoolTokenTransfer

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainPoolTokenTransferIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainPoolTokenTransfer)
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
		it.Event = new(CrossChainPoolTokenTransfer)
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

func (it *CrossChainPoolTokenTransferIterator) Error() error {
	return it.fail
}

func (it *CrossChainPoolTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainPoolTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CrossChainPoolTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainPoolTokenTransferIterator{contract: _CrossChainPoolToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainPoolTokenTransfer)
				if err := _CrossChainPoolToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) ParseTransfer(log types.Log) (*CrossChainPoolTokenTransfer, error) {
	event := new(CrossChainPoolTokenTransfer)
	if err := _CrossChainPoolToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
	FeeAdmin       common.Address
}
type GetFee struct {
	FeeUSDCents       *big.Int
	DestGasOverhead   uint32
	DestBytesOverhead uint32
	TokenFeeBps       uint16
	IsEnabled         bool
}

func (CrossChainPoolTokenAdvancedPoolHooksUpdated) Topic() common.Hash {
	return common.HexToHash("0xbaff46844acf36d6ee996f489a1a288709c4542bd33cd557770afd267d581209")
}

func (CrossChainPoolTokenApproval) Topic() common.Hash {
	return common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
}

func (CrossChainPoolTokenCCIPAdminTransferred) Topic() common.Hash {
	return common.HexToHash("0x9524c9e4b0b61eb018dd58a1cd856e3e74009528328ab4a613b434fa631d7242")
}

func (CrossChainPoolTokenChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (CrossChainPoolTokenChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (CrossChainPoolTokenCustomBlockConfirmationsInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x63335ad9e238acd0e9e6c1c20f529ffbea4cda73578c329a7aa7e9d61e5cdcc5")
}

func (CrossChainPoolTokenCustomBlockConfirmationsOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x996b829383cc7e136842d4c4c175083bcf4e20807c7432105c1b794ba885e776")
}

func (CrossChainPoolTokenDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x3f1036e85d016a93254a0b1415844f79b85424959d90ae5ad51ce8f4533fe701")
}

func (CrossChainPoolTokenFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (CrossChainPoolTokenInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (CrossChainPoolTokenLockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
}

func (CrossChainPoolTokenMinBlockConfirmationsSet) Topic() common.Hash {
	return common.HexToHash("0x46c9c0585a955b2702c7ea47fec541db623565d20827a0edda42864e6b859a01")
}

func (CrossChainPoolTokenOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944")
}

func (CrossChainPoolTokenOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (CrossChainPoolTokenOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (CrossChainPoolTokenRateLimitConfigured) Topic() common.Hash {
	return common.HexToHash("0x41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb")
}

func (CrossChainPoolTokenReleasedOrMinted) Topic() common.Hash {
	return common.HexToHash("0xfc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0")
}

func (CrossChainPoolTokenRemotePoolAdded) Topic() common.Hash {
	return common.HexToHash("0x7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea")
}

func (CrossChainPoolTokenRemotePoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76")
}

func (CrossChainPoolTokenTokenTransferFeeConfigDeleted) Topic() common.Hash {
	return common.HexToHash("0x5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8")
}

func (CrossChainPoolTokenTokenTransferFeeConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xfae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa1041")
}

func (CrossChainPoolTokenTransfer) Topic() common.Hash {
	return common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
}

func (_CrossChainPoolToken *CrossChainPoolToken) Address() common.Address {
	return _CrossChainPoolToken.address
}

type CrossChainPoolTokenInterface interface {
	Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error)

	BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error)

	Decimals(opts *bind.CallOpts) (uint8, error)

	GetAdvancedPoolHooks(opts *bind.CallOpts) (common.Address, error)

	GetCCIPAdmin(opts *bind.CallOpts) (common.Address, error)

	GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64, customBlockConfirmations bool) (GetCurrentRateLimiterState,

		error)

	GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

		error)

	GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationsRequested uint16, arg5 []byte) (GetFee,

		error)

	GetMinBlockConfirmations(opts *bind.CallOpts) (uint16, error)

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRequiredCCVs(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, sourceDenominatedAmount *big.Int, blockConfirmationsRequested uint16, extraData []byte, direction uint8) ([]common.Address, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetSupportedChains(opts *bind.CallOpts) ([]uint64, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	GetTokenDecimals(opts *bind.CallOpts) (uint8, error)

	GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error)

	IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error)

	IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error)

	IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error)

	MaxSupply(opts *bind.CallOpts) (*big.Int, error)

	Name(opts *bind.CallOpts) (string, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	Symbol(opts *bind.CallOpts) (string, error)

	TotalSupply(opts *bind.CallOpts) (*big.Int, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error)

	ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error)

	Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationsRequested uint16, tokenArgs []byte) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationsRequested uint16) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetCCIPAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error)

	SetMinBlockConfirmations(opts *bind.TransactOpts, minBlockConfirmations uint16) (*types.Transaction, error)

	SetRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error)

	Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error)

	TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateAdvancedPoolHooks(opts *bind.TransactOpts, newHook common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error)

	FilterAdvancedPoolHooksUpdated(opts *bind.FilterOpts) (*CrossChainPoolTokenAdvancedPoolHooksUpdatedIterator, error)

	WatchAdvancedPoolHooksUpdated(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenAdvancedPoolHooksUpdated) (event.Subscription, error)

	ParseAdvancedPoolHooksUpdated(log types.Log) (*CrossChainPoolTokenAdvancedPoolHooksUpdated, error)

	FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CrossChainPoolTokenApprovalIterator, error)

	WatchApproval(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error)

	ParseApproval(log types.Log) (*CrossChainPoolTokenApproval, error)

	FilterCCIPAdminTransferred(opts *bind.FilterOpts, previousAdmin []common.Address, newAdmin []common.Address) (*CrossChainPoolTokenCCIPAdminTransferredIterator, error)

	WatchCCIPAdminTransferred(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenCCIPAdminTransferred, previousAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error)

	ParseCCIPAdminTransferred(log types.Log) (*CrossChainPoolTokenCCIPAdminTransferred, error)

	FilterChainAdded(opts *bind.FilterOpts) (*CrossChainPoolTokenChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*CrossChainPoolTokenChainAdded, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*CrossChainPoolTokenChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*CrossChainPoolTokenChainRemoved, error)

	FilterCustomBlockConfirmationsInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CrossChainPoolTokenCustomBlockConfirmationsInboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationsInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenCustomBlockConfirmationsInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationsInboundRateLimitConsumed(log types.Log) (*CrossChainPoolTokenCustomBlockConfirmationsInboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationsOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CrossChainPoolTokenCustomBlockConfirmationsOutboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationsOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenCustomBlockConfirmationsOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationsOutboundRateLimitConsumed(log types.Log) (*CrossChainPoolTokenCustomBlockConfirmationsOutboundRateLimitConsumed, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*CrossChainPoolTokenDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*CrossChainPoolTokenDynamicConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*CrossChainPoolTokenFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*CrossChainPoolTokenFeeTokenWithdrawn, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CrossChainPoolTokenInboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*CrossChainPoolTokenInboundRateLimitConsumed, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CrossChainPoolTokenLockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*CrossChainPoolTokenLockedOrBurned, error)

	FilterMinBlockConfirmationsSet(opts *bind.FilterOpts) (*CrossChainPoolTokenMinBlockConfirmationsSetIterator, error)

	WatchMinBlockConfirmationsSet(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenMinBlockConfirmationsSet) (event.Subscription, error)

	ParseMinBlockConfirmationsSet(log types.Log) (*CrossChainPoolTokenMinBlockConfirmationsSet, error)

	FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CrossChainPoolTokenOutboundRateLimitConsumedIterator, error)

	WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseOutboundRateLimitConsumed(log types.Log) (*CrossChainPoolTokenOutboundRateLimitConsumed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CrossChainPoolTokenOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CrossChainPoolTokenOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CrossChainPoolTokenOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CrossChainPoolTokenOwnershipTransferred, error)

	FilterRateLimitConfigured(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CrossChainPoolTokenRateLimitConfiguredIterator, error)

	WatchRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenRateLimitConfigured, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRateLimitConfigured(log types.Log) (*CrossChainPoolTokenRateLimitConfigured, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CrossChainPoolTokenReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*CrossChainPoolTokenReleasedOrMinted, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CrossChainPoolTokenRemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*CrossChainPoolTokenRemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CrossChainPoolTokenRemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*CrossChainPoolTokenRemotePoolRemoved, error)

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*CrossChainPoolTokenTokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*CrossChainPoolTokenTokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*CrossChainPoolTokenTokenTransferFeeConfigUpdatedIterator, error)

	WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigUpdated(log types.Log) (*CrossChainPoolTokenTokenTransferFeeConfigUpdated, error)

	FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CrossChainPoolTokenTransferIterator, error)

	WatchTransfer(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseTransfer(log types.Log) (*CrossChainPoolTokenTransfer, error)

	Address() common.Address
}
