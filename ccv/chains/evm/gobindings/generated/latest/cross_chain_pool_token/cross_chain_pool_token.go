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
	DestGasOverhead            uint32
	DestBytesOverhead          uint32
	FinalityFeeUSDCents        uint32
	FastFinalityFeeUSDCents    uint32
	FinalityTransferFeeBps     uint16
	FastFinalityTransferFeeBps uint16
	IsEnabled                  bool
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
	FastFinality              bool
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
}

type TokenPoolTokenTransferFeeConfigArgs struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
}

var CrossChainPoolTokenMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenParams\",\"type\":\"tuple\",\"internalType\":\"struct BaseERC20.ConstructorParams\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"maxSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"preMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"preMintRecipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"ccipAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"advancedPoolHooks\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fastFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fastFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"disableTokenTransferFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"_decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAdvancedPoolHooks\",\"inputs\":[],\"outputs\":[{\"name\":\"advancedPoolHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowedFinalityConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"allowedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCIPAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"ccipAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fastFinality\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requestedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fastFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fastFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"lockOrBurnOutV1\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"requestedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maxSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"_maxSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"requestedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAllowedFinalityConfig\",\"inputs\":[{\"name\":\"allowedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCCIPAdmin\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.RateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fastFinality\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"updateAdvancedPoolHooks\",\"inputs\":[{\"name\":\"newHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AdvancedPoolHooksUpdated\",\"inputs\":[{\"name\":\"oldHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"},{\"name\":\"newHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCIPAdminTransferred\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"feeAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FastFinalityInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FastFinalityOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalityConfigSet\",\"inputs\":[{\"name\":\"allowedFinality\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"fastFinality\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fastFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fastFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotOwnerOrFeeAdmin\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotRenounceCCIPAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRequestedFinality\",\"inputs\":[{\"name\":\"requestedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"allowedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MaxSupplyExceeded\",\"inputs\":[{\"name\":\"supplyAfterMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCCIPAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"PreMintAddressNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PreMintRecipientSetWithZeroPreMint\",\"inputs\":[{\"name\":\"preMintRecipient\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RequestedFinalityCanOnlyHaveOneMode\",\"inputs\":[{\"name\":\"encodedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x61012080604052346106c5576165a7803803809161001d82856106ca565b83398101906080818303126106c55780516001600160401b0381116106c55781019160e0838203126106c5576040519060e082016001600160401b038111838210176105c25760405283516001600160401b0381116106c557816100829186016106ed565b82526020840151906001600160401b0382116106c5576100a39185016106ed565b9182602083015260408401519060408301918252606085015193606084019485526100d06080870161075c565b936080810194855260a08701519160ff8316978884036106c55760c06100fd9160a085019586520161075c565b9760c083019889526101116020860161075c565b9161012a60606101236040890161075c565b970161075c565b93518051906001600160401b0382116105c25760035490600182811c921680156106bb575b60208310146105a25781601f84931161064b575b50602090601f83116001146105e3576000926105d8575b50508160011b916000199060031b1c1916176003555b8051906001600160401b0382116105c25760045490600182811c921680156105b8575b60208310146105a25781601f849311610532575b50602090601f83116001146104ca576000926104bf575b50508160011b916000199060031b1c1916176004555b33156104ae57600680546001600160a01b031916331790553015801561049d575b801561048c575b61047b573060805260c09490945260a093909352600880546001600160a01b039485166001600160a01b03199182161790915560078054929094169116179091555160ff1660e05251610100528151156104505780516001600160a01b03161561043f57519051906001600160a01b031680156104295760025491808301809311610413576020926002557fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef600093849284845283825260408420818154019055604051908152a361010051806103ed575b50505b516001600160a01b0316806103e85750335b601280546001600160a01b039283166001600160a01b03198216811790925560405192167f9524c9e4b0b61eb018dd58a1cd856e3e74009528328ab4a613b434fa631d7242600080a3615e36908161077182396080518181816102b50152818161220901528181612a610152818161309f015281816135f101526137a4015260a0518181816134c401528181614b3001528181614b7a015261516d015260c0518181816103430152818161132b0152818161229601528181612aef015261312d015260e05181612fed0152610100518181816117eb0152818161559c01526156de0152f35b610303565b6002548181116103fd57506102ee565b637502c12360e11b835260045260245260449150fd5b634e487b7160e01b600052601160045260246000fd5b63ec442f0560e01b600052600060045260246000fd5b634dd371db60e11b60005260046000fd5b516001600160a01b031690508061046757506102f1565b63f5c8f5a160e01b60005260045260246000fd5b630a64406560e11b60005260046000fd5b506001600160a01b0385161561021c565b506001600160a01b03831615610215565b639b15e16f60e01b60005260046000fd5b0151905038806101de565b600460009081528281209350601f198516905b81811061051a5750908460019594939210610501575b505050811b016004556101f4565b015160001960f88460031b161c191690553880806104f3565b929360206001819287860151815501950193016104dd565b60046000529091507f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b601f840160051c81019160208510610598575b90601f859493920160051c01905b81811061058957506101c7565b6000815584935060010161057c565b909150819061056e565b634e487b7160e01b600052602260045260246000fd5b91607f16916101b3565b634e487b7160e01b600052604160045260246000fd5b01519050388061017a565b600360009081528281209350601f198516905b818110610633575090846001959493921061061a575b505050811b01600355610190565b015160001960f88460031b161c1916905538808061060c565b929360206001819287860151815501950193016105f6565b60036000529091507fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b601f840160051c810191602085106106b1575b90601f859493920160051c01905b8181106106a25750610163565b60008155849350600101610695565b9091508190610687565b91607f169161014f565b600080fd5b601f909101601f19168101906001600160401b038211908210176105c257604052565b81601f820112156106c5578051906001600160401b0382116105c25760405192610721601f8401601f1916602001856106ca565b828452602083830101116106c55760005b82811061074757505060206000918301015290565b80602080928401015182828701015201610732565b51906001600160a01b03821682036106c55756fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a714613b8a5750806306b859ef14613ab257806306fdde0314613a0b578063095ea7b3146138fe57806318160ddd146138e0578063181f5a771461387f5780631826b1e7146137c857806321df0da71461378457806323b872dd1461361e578063240028e8146135c75780632422ac45146134e857806324f65ee7146134aa5780632cab0fb614613011578063313ce56714612fd357806337a3210d14612fac57806339077537146129c35780634c5ef0ed1461297c57806362ddd3c4146128f557806370a08231146128be5780637437ff9f1461287d57806379ba5097146127ce5780638926f54f146127885780638da5cb5b146127615780638fd6a6ac1461273a57806395d89b411461264d5780639a4575b91461219d578063a42a7b8b14612025578063a8fa343c14611fa4578063a9059cbb14611f72578063acfecf9114611e7a578063ae39a25714611d23578063b6cfa3b714611c68578063b794658014611c30578063bfeffd3f14611b9e578063c4bffe2b14611a73578063c7230a601461180e578063d5abeb01146117d3578063dc04fa1f1461134f578063dc0bd9711461130b578063dcbd41bc14611121578063dd62ed3e146110d1578063e8a1da1714610a31578063ea6396db146108f3578063ec6ae7a7146108b0578063f2fde38b146107f95763fbc801a71461021b57600080fd5b346107f65760606003193601126107f6576004359067ffffffffffffffff82116107f6578160040160a060031984360301126107f257610259613d82565b9060443567ffffffffffffffff8111610684579061027e61029b923690600401613e79565b92906102886148f7565b5061029385846153d3565b933691613ff3565b9260848601936102aa85614891565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116036107b557602487019677ffffffffffffffff00000000000000000000000000000000610303896148a5565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115610728578891610786575b5061075e5767ffffffffffffffff61038a896148a5565b166103a281600052600c602052604060002054151590565b156107335760206001600160a01b0360075416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa80156107285788906106e4575b6001600160a01b0391501633036106b8576064810135937fffffffff0000000000000000000000000000000000000000000000000000000061043887876141da565b94169384156106975761045d8161044e8a614891565b6104578d6148a5565b9061588c565b6001600160a01b0360085416938461057c575b50505050505090610480916141da565b9161048a846148a5565b503015610550575061051061054693610515926104a7853061553c565b7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff6104e36104dd856148a5565b93614891565b604080516001600160a01b039092168252336020830152810188905292169180606081015b0390a26148a5565b614953565b9061051e615166565b6040519261052b84613f5e565b835260208301526040519283926040845260408401906140bb565b9060208301520390f35b807f96c6fd1e000000000000000000000000000000000000000000000000000000006024925280600452fd5b843b15610693578994928b9694928692604051988997889687957fa8027c0f0000000000000000000000000000000000000000000000000000000087526004870160809052806105cb916157f6565b6084880160a090526101248801906105e2926141fb565b926105ec90613e64565b67ffffffffffffffff1660a487015260440161060790613e22565b6001600160a01b031660c48601528d8c60e487015261062590613e22565b6001600160a01b0316610104860152602485015283810360031901604485015261064e91613ea7565b90606483015203925af180156106885790859161066f575b80808080610470565b8161067991613fb2565b610684578338610666565b8380fd5b6040513d87823e3d90fd5b8980fd5b6106b3816106a48a614891565b6106ad8d6148a5565b90615846565b61045d565b6024877f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d602011610720575b816106fe60209383613fb2565b8101031261071c576107176001600160a01b03916141e7565b6103f6565b8780fd5b3d91506106f1565b6040513d8a823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008852600452602487fd5b6004877f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b6107a8915060203d6020116107ae575b6107a08183613fb2565b810190614ee2565b38610373565b503d610796565b6024866001600160a01b036107c988614891565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835216600452fd5b5080fd5b80fd5b50346107f65760206003193601126107f6576001600160a01b0361081b613de0565b610823614efa565b1633811461088857807fffffffffffffffffffffffff000000000000000000000000000000000000000060055416176005556001600160a01b03600654167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b50346107f657806003193601126107f65760207fffffffff0000000000000000000000000000000000000000000000000000000060075460401b16604051908152f35b50346107f65760806003193601126107f65761090d613de0565b50610916613e36565b61091e613db1565b5060643567ffffffffffffffff8111610a2d579167ffffffffffffffff60409261094e60e0953690600401613e79565b50508260c0855161095e81613f96565b82815282602082015282878201528260608201528260808201528260a082015201521681526010602052206040519061099682613f96565b5461ffff818163ffffffff82169485815263ffffffff60208201818560201c1681528160408401818760401c168152816060860193818960601c16855260ff60c060808901988a8c60801c168a528a60a082019c60901c168c52019b60a01c1615158b526040519b8c52511660208b0152511660408901525116606087015251166080850152511660a083015251151560c0820152f35b8280fd5b50346107f65760406003193601126107f65760043567ffffffffffffffff81116107f257610a639036906004016140e5565b9060243567ffffffffffffffff81116106845790610a86849236906004016140e5565b939091610a91614efa565b83905b828210610f125750505081927ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b81811015610f0e578060051b83013585811215610f0a57830161012081360312610f0a5760405194610af886613f7a565b610b0182613e64565b8652602082013567ffffffffffffffff81116107f25782019436601f870112156107f257853595610b3187614147565b96610b3f6040519889613fb2565b80885260208089019160051b83010190368211610f0a5760208301905b828210610ed7575050505060208701958652604083013567ffffffffffffffff8111610a2d57610b8f9036908501614058565b9160408801928352610bb9610ba73660608701614a09565b9460608a0195865260c0369101614a09565b956080890196875283515115610eaf57610bdd67ffffffffffffffff8a5116615b1b565b15610e785767ffffffffffffffff8951168252600d60205260408220610c048651826151a1565b610c128851600283016151a1565b6004855191019080519067ffffffffffffffff8211610e4b57610c358354614518565b601f8111610e10575b50602090601f8311600114610d8f57610c6e9291869183610d84575b50506000198260011b9260031b1c19161790565b90555b815b88518051821015610ca85790610ca2600192610c9b8367ffffffffffffffff8f511692614910565b5190614f38565b01610c73565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2939199975095610d7667ffffffffffffffff6001979694985116925193519151610d42610d0d60405196879687526101006020880152610100870190613ea7565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a1019392909193610ac7565b015190508e80610c5a565b83865281862091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416875b818110610df85750908460019594939210610ddf575b505050811b019055610c71565b015160001960f88460031b161c191690558d8080610dd2565b92936020600181928786015181550195019301610dbc565b610e3b9084875260208720601f850160051c81019160208610610e41575b601f0160051c0190614aa5565b8d610c3e565b9091508190610e2e565b6024857f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60248267ffffffffffffffff8b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b6004827f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b813567ffffffffffffffff8111610f0657602091610efb8392833691890101614058565b815201910190610b5c565b8680fd5b8480fd5b8380f35b9267ffffffffffffffff610f34610f2f8486889a9699979a6149dc565b6148a5565b1691610f3f8361595f565b156110a557828452600d602052610f5b600560408620016158fc565b94845b8651811015610f9457600190858752600d602052610f8d60056040892001610f86838b614910565b5190615a5f565b5001610f5e565b50939692909450949094808752600d6020526005604088208881558860018201558860028201558860038201558860048201610fd08154614518565b80611064575b5050500180549088815581611046575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020836001948a52600982528985604082208281550155808a52600a82528985604082208281550155604051908152a101909194939294610a94565b885260208820908101905b81811015610fe657888155600101611051565b601f811160011461107a5750555b888a80610fd6565b8183526020832061109591601f01861c810190600101614aa5565b8082528160208120915555611072565b602484847f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50346107f65760406003193601126107f6576001600160a01b0360406110f5613de0565b92826110ff613e0c565b9416815260016020522091166000526020526020604060002054604051908152f35b50346107f65760206003193601126107f65760043567ffffffffffffffff81116107f257611153903690600401614116565b6001600160a01b03600f5416331415806112f6575b6112ca57825b818110611179578380f35b61118481838561497f565b67ffffffffffffffff611196826148a5565b16906111af82600052600c602052604060002054151590565b1561129e57907f41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb60e08361125e611238602060019897018b6111f08261498f565b1561126557879052600960205261121760408d206112113660408801614a09565b906151a1565b868c52600a60205261123360408d206112113660a08801614a09565b61498f565b9160405192151583526112516020840160408301614a61565b60a0608084019101614a61565ba20161116e565b60026040828a6112339452600d60205261128782822061121136858c01614a09565b8a8152600d60205220016112113660a08801614a09565b602486837f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b6024837f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b506001600160a01b0360065416331415611168565b50346107f657806003193601126107f65760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346107f65760406003193601126107f65760043567ffffffffffffffff81116107f257611381903690600401614116565b60243567ffffffffffffffff8111610684576113a19036906004016140e5565b9190926113ac614efa565b845b82811061141857505050825b8181106113c5578380f35b8067ffffffffffffffff6113df610f2f60019486886149dc565b1680865260106020528560408120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee88680a2016113ba565b67ffffffffffffffff61142f610f2f83868661497f565b1661144781600052600c602052604060002054151590565b156117a85761145782858561497f565b602081019060e081019061146a8261498f565b1561177c5760a0810161271061ffff6114828361499c565b16101561176d5760c082019161271061ffff61149d8561499c565b1610156117355763ffffffff6114b2866149ab565b161561170957858c52601060205260408c206114cd866149ab565b63ffffffff169080549060408401916114e5836149ab565b60201b67ffffffff0000000016936060860194611501866149ab565b60401b6bffffffff0000000000000000169660800196611520886149ab565b60601b6fffffffff000000000000000000000000169161153f8a61499c565b60801b71ffff0000000000000000000000000000000016936115608c61499c565b60901b73ffff00000000000000000000000000000000000016957fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016177fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16177fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff16177fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff1617171781556116138761498f565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff00000000000000000000000000000000000000001617905560405196611664906149bc565b63ffffffff168752611675906149bc565b63ffffffff166020870152611689906149bc565b63ffffffff16604086015261169d906149bc565b63ffffffff1660608501526116b1906149cd565b61ffff1660808401526116c3906149cd565b61ffff1660a08301526116d590613f06565b151560c082015260e07ffae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa104191a26001016113ae565b60248c877f12332265000000000000000000000000000000000000000000000000000000008252600452fd5b60248c61ffff6117448661499c565b7f95f3517a00000000000000000000000000000000000000000000000000000000835216600452fd5b8a61ffff61174460249361499c565b60248a857f12332265000000000000000000000000000000000000000000000000000000008252600452fd5b7f1e670e4b000000000000000000000000000000000000000000000000000000008752600452602486fd5b50346107f657806003193601126107f65760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b50346107f65760406003193601126107f65760043567ffffffffffffffff81116107f2576118409036906004016140e5565b90611849613e0c565b916001600160a01b036006541633141580611a5e575b611a32576001600160a01b038316908115611a0a57845b818110611881578580f35b6001600160a01b0361189c6118978385886149dc565b614891565b166040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481855afa9081156107285788916119d7575b50806118f1575b5050600101611876565b6040517fa9059cbb0000000000000000000000000000000000000000000000000000000060208083019182526001600160a01b038a16602484015260448084018590528352918a9190611945606482613fb2565b519082865af1156119cc5787513d6119c35750813b155b6119975790847f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e602060019594604051908152a390386118e7565b602488837f5274afe7000000000000000000000000000000000000000000000000000000008252600452fd5b6001141561195c565b6040513d89823e3d90fd5b905060203d8111611a03575b6119ed8183613fb2565b602082600092810103126107f6575051386118e0565b503d6119e3565b6004857f8579befe000000000000000000000000000000000000000000000000000000008152fd5b6024847fcb1afbd700000000000000000000000000000000000000000000000000000000815233600452fd5b506001600160a01b036011541633141561185f565b50346107f657806003193601126107f65760405190600b548083528260208101600b84526020842092845b818110611b85575050611ab392500383613fb2565b8151611ad7611ac182614147565b91611acf6040519384613fb2565b808352614147565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602083019301368437805b8451811015611b36578067ffffffffffffffff611b2360019388614910565b5116611b2f8286614910565b5201611b04565b50925090604051928392602084019060208552518091526040840192915b818110611b62575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101611b54565b8454835260019485019487945060209093019201611a9e565b50346107f65760206003193601126107f6576004356001600160a01b0381168091036107f257611bcc614efa565b7fffffffffffffffffffffffff00000000000000000000000000000000000000006008547fbaff46844acf36d6ee996f489a1a288709c4542bd33cd557770afd267d581209604080516001600160a01b0384168152856020820152a1161760085580f35b50346107f65760206003193601126107f657611c64611c50610510613e4d565b604051918291602083526020830190613ea7565b0390f35b50346107f65760206003193601126107f6577f307cf716eade81675bea3ccb6917b0f91baa2160056765d9a83d76f819caf06a6020611ca5613d4e565b611cad614efa565b6007547fffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffff77ffffffff0000000000000000000000000000000000000000808460401c16169116176007557fffffffff0000000000000000000000000000000000000000000000000000000060405191168152a180f35b50346107f65760606003193601126107f657611d3d613de0565b90611d46613e0c565b604435926001600160a01b03841680850361068457611d63614efa565b6001600160a01b0382168015611e525794611e4c917f3f1036e85d016a93254a0b1415844f79b85424959d90ae5ad51ce8f4533fe70195967fffffffffffffffffffffffff000000000000000000000000000000000000000060075416176007556001600160a01b0385167fffffffffffffffffffffffff0000000000000000000000000000000000000000600f541617600f557fffffffffffffffffffffffff00000000000000000000000000000000000000006011541617601155604051938493849160409194936001600160a01b03809281606087019816865216602085015216910152565b0390a180f35b6004857f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b50346107f65767ffffffffffffffff611e9236614076565b929091611e9d614efa565b1691611eb683600052600c602052604060002054151590565b156110a557828452600d602052611ee560056040862001611ed8368486613ff3565b6020815191012090615a5f565b15611f2a57907f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d7691611f246040519283926020845260208401916141fb565b0390a280f35b82611f6e836040519384937f74f23c7c00000000000000000000000000000000000000000000000000000000855260048501526040602485015260448401916141fb565b0390fd5b50346107f65760406003193601126107f657611f99611f8f613de0565b6024359033614d4d565b602060405160018152f35b50346107f65760206003193601126107f657611fbe613de0565b611fc6614efa565b6001600160a01b0380601254921691827fffffffffffffffffffffffff0000000000000000000000000000000000000000821617601255167f9524c9e4b0b61eb018dd58a1cd856e3e74009528328ab4a613b434fa631d72428380a380f35b50346107f65760206003193601126107f65767ffffffffffffffff612048613e4d565b168152600d60205261205f600560408320016158fc565b80517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06120a461208e83614147565b9261209c6040519485613fb2565b808452614147565b01835b81811061218c575050825b825181101561210957806120c860019285614910565b518552600e6020526120e66120ed604087206040519283809261456b565b0382613fb2565b6120f78285614910565b526121028184614910565b50016120b2565b81846040519182916020830160208452825180915260408401602060408360051b870101940192905b82821061214157505050500390f35b9193602061217c827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851613ea7565b9601920192018594939192612132565b8060606020809386010152016120a7565b50346107f65760206003193601126107f65760043567ffffffffffffffff81116107f257806004019060a06003198236030112610a2d576121dc6148f7565b506040516020936121ed8583613fb2565b80825260848301916121fe83614891565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691160361263957602484019477ffffffffffffffff00000000000000000000000000000000612257876148a5565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015287816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156125be57849161261c575b506125f45767ffffffffffffffff6122dd876148a5565b166122f581600052600c602052604060002054151590565b156125c957876001600160a01b0360075416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa80156125be578490612583575b6001600160a01b039150163303612557576064850135946123758661236c87614891565b6106ad8a6148a5565b6001600160a01b03600854169182612449575b50505050612395846148a5565b503015610550575091817ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff612419956123da610510963061553c565b6105086123ef6123e9876148a5565b92614891565b604080516001600160a01b0390921682523360208301528101959095529116929081906060820190565b90612422615166565b6040519261242f84613f5e565b835281830152611c646040519282849384528301906140bb565b823b15610f0a57918791858094604051968795869485937fa8027c0f000000000000000000000000000000000000000000000000000000008552600485016080905280612495916157f6565b6084860160a090526101248601906124ac926141fb565b916124b690613e64565b67ffffffffffffffff1660a48501526044016124d190613e22565b6001600160a01b031660c48401528b60e48401526124ee8b613e22565b6001600160a01b031661010484015283602484015282810360031901604484015261251891613ea7565b8a606483015203925af1801561254c57908291612537575b8080612388565b8161254191613fb2565b6107f6578038612530565b6040513d84823e3d90fd5b6024837f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b508781813d83116125b7575b6125998183613fb2565b81010312610684576125b26001600160a01b03916141e7565b612348565b503d61258f565b6040513d86823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008452600452602483fd5b6004837f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b6126339150883d8a116107ae576107a08183613fb2565b386122c6565b506001600160a01b036107c9602493614891565b50346107f657806003193601126107f65760405190806004549061267082614518565b80855291600181169081156126f55750600114612698575b611c6484611c5081860382613fb2565b600481527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b939250905b8082106126db57509091508101602001611c5082612688565b9192600181602092548385880101520191019092916126c2565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208087019190915292151560051b85019092019250611c509150839050612688565b50346107f657806003193601126107f65760206001600160a01b0360125416604051908152f35b50346107f657806003193601126107f65760206001600160a01b0360065416604051908152f35b50346107f65760206003193601126107f65760206127c467ffffffffffffffff6127b0613e4d565b16600052600c602052604060002054151590565b6040519015158152f35b50346107f657806003193601126107f6576005546001600160a01b0381163303612855577fffffffffffffffffffffffff0000000000000000000000000000000000000000600654913382841617600655166005556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b50346107f657806003193601126107f657600754600f54601154604080516001600160a01b0394851681529284166020840152921691810191909152606090f35b50346107f65760206003193601126107f65760406020916001600160a01b036128e5613de0565b1681528083522054604051908152f35b50346107f65761290436614076565b61291093929193614efa565b67ffffffffffffffff821661293281600052600c602052604060002054151590565b15612951575061294e9293612948913691613ff3565b90614f38565b80f35b7f1e670e4b000000000000000000000000000000000000000000000000000000008452600452602483fd5b50346107f65760406003193601126107f657612996613e4d565b906024359067ffffffffffffffff82116107f65760206127c4846129bd3660048701614058565b906148ba565b50346107f65760206003193601126107f6576004359067ffffffffffffffff82116107f657816004019061010060031984360301126107f65780604051612a0981613f13565b5280604051612a1781613f13565b52606483013560c4840193612a47612a41612a3c612a358888614840565b3691613ff3565b614abc565b83614b77565b936084820195612a5687614891565b6001600160a01b03807f000000000000000000000000000000000000000000000000000000000000000016911603612f9857602483019377ffffffffffffffff00000000000000000000000000000000612aaf866148a5565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156119cc578791612f79575b50612f515767ffffffffffffffff612b36866148a5565b16612b4e81600052600c602052604060002054151590565b15612f265760206001600160a01b0360075416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa9081156119cc578791612f07575b5015612edb57612bb8856148a5565b92612bce60a48601946129bd612a358785614840565b15612e9457612bef88612be08b614891565b612be9896148a5565b9061571a565b6001600160a01b03600854169283612ce0575b505050505060440191612c1483614891565b90612c1e836148a5565b506001600160a01b03821615612cb4575067ffffffffffffffff6020956001600160a01b03612c80612c7a6104dd7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc097610f2f8b608099615685565b96614891565b816040519716875233898801521660408601528560608601521692a260405190612ca982613f13565b815260405190518152f35b807fec442f05000000000000000000000000000000000000000000000000000000006024925280600452fd5b833b1561071c57878795938195938c93604051988997889687957f6371157400000000000000000000000000000000000000000000000000000000875260048701606090528d612d3087806157f6565b60648a0161010090526101648a0190612d48926141fb565b94612d5290613e64565b67ffffffffffffffff166084890152604401612d6d90613e22565b6001600160a01b031660a488015260c4870152612d8990613e22565b6001600160a01b031660e4860152612da190846157f6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c86840301610104870152612dd692916141fb565b90612de190836157f6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c85840301610124860152612e1692916141fb565b9060e48a01612e24916157f6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c84840301610144850152612e5992916141fb565b8b602483015282604483015203925af180156125be57908491612e7f575b808080612c02565b81612e8991613fb2565b610a2d578238612e77565b83612e9e91614840565b611f6e6040519283927f24eb47e50000000000000000000000000000000000000000000000000000000084526020600485015260248401916141fb565b6024867f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b612f20915060203d6020116107ae576107a08183613fb2565b38612ba9565b7fa9902c7e000000000000000000000000000000000000000000000000000000008752600452602486fd5b6004867f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b612f92915060203d6020116107ae576107a08183613fb2565b38612b1f565b6024856001600160a01b036107c98a614891565b50346107f657806003193601126107f65760206001600160a01b0360085416604051908152f35b50346107f657806003193601126107f657602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346107f65760406003193601126107f6576004359067ffffffffffffffff82116107f657816004019061010060031984360301126107f657613052613d82565b918160405161306081613f13565b5260648401359360c481019361308561307f612a3c612a358887614840565b87614b77565b94608483019661309488614891565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691160361349657602484019477ffffffffffffffff000000000000000000000000000000006130ed876148a5565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115610728578891613477575b5061075e5767ffffffffffffffff613174876148a5565b1661318c81600052600c602052604060002054151590565b156107335760206001600160a01b0360075416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115610728578891613458575b50156106b8576131f6866148a5565b9361320c60a48701956129bd612a358886614840565b1561344e577fffffffff000000000000000000000000000000000000000000000000000000001690811561343357613256896132478c614891565b6132508a6148a5565b90615786565b6001600160a01b0360085416938461327c575b50505050505060440191612c1483614891565b843b1561342f57868995938c959387938b6040519a8b998a9889977f6371157400000000000000000000000000000000000000000000000000000000895260048901606090526132cc87806157f6565b60648b0161010090526101648b01906132e4926141fb565b946132ee90613e64565b67ffffffffffffffff1660848a015260440161330990613e22565b6001600160a01b031660a489015260c488015261332590613e22565b6001600160a01b031660e487015261333d90846157f6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c8784030161010488015261337292916141fb565b9061337d90836157f6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c868403016101248701526133b292916141fb565b9060e48b016133c0916157f6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c858403016101448601526133f592916141fb565b908c6024840152604483015203925af180156125be5761341a575b8080808080613269565b926134288160449395613fb2565b9290613410565b8880fd5b613449896134408c614891565b612be98a6148a5565b613256565b612e9e8583614840565b613471915060203d6020116107ae576107a08183613fb2565b386131e7565b613490915060203d6020116107ae576107a08183613fb2565b3861315d565b6024866001600160a01b036107c98b614891565b50346107f657806003193601126107f657602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346107f65760406003193601126107f657613502613e4d565b6024359182151583036107f6576101406135c561351f85856147bd565b61357560409392935180946fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b60a08301906fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565bf35b50346107f65760206003193601126107f6576020906135e4613de0565b90506001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b50346107f65760606003193601126107f657613638613de0565b613640613e0c565b604435916001600160a01b0381168085526001602052604085206001600160a01b03331686526020526040852054906000198210613685575b5050611f999350614d4d565b848210613750573033146137245780156136f85733156136cc576040868692611f99985260016020528181206001600160a01b033316825260205220910390553880613679565b6024867f94280d6200000000000000000000000000000000000000000000000000000000815280600452fd5b6024867fe602df0500000000000000000000000000000000000000000000000000000000815280600452fd5b6024867f94280d6200000000000000000000000000000000000000000000000000000000815233600452fd5b60648686847ffb8f41b200000000000000000000000000000000000000000000000000000000835233600452602452604452fd5b50346107f657806003193601126107f65760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346107f65760c06003193601126107f6576137e2613de0565b506137eb613e36565b6137f3613df6565b50608435917fffffffff00000000000000000000000000000000000000000000000000000000831683036107f65760a4359067ffffffffffffffff82116107f65760a063ffffffff8061ffff61385888886138513660048b01613e79565b505061460d565b95926040979194975197885216602087015216604085015216606083015215156080820152f35b50346107f657806003193601126107f65750611c646040516138a2604082613fb2565b601d81527f43726f7373436861696e506f6f6c546f6b656e20322e302e302d6465760000006020820152604051918291602083526020830190613ea7565b50346107f657806003193601126107f6576020600254604051908152f35b50346107f65760406003193601126107f657613918613de0565b6001600160a01b036024359116913083146139df5733156139b35782156139875760408291338152600160205281812085825260205220556040519081527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560203392a3602060405160018152f35b807f94280d62000000000000000000000000000000000000000000000000000000006024925280600452fd5b807fe602df05000000000000000000000000000000000000000000000000000000006024925280600452fd5b80837f94280d620000000000000000000000000000000000000000000000000000000060249352600452fd5b50346107f657806003193601126107f657604051908060035490613a2e82614518565b80855291600181169081156126f55750600114613a5557611c6484611c5081860382613fb2565b600381527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b939250905b808210613a9857509091508101602001611c5082612688565b919260018160209254838588010152019101909291613a7f565b50346107f65760c06003193601126107f657613acc613de0565b613ad4613e36565b906064357fffffffff00000000000000000000000000000000000000000000000000000000811681036106845760843567ffffffffffffffff8111610f0a57613b21903690600401613e79565b9160a435936002851015610f0657613b3c956044359161423a565b90604051918291602083016020845282518091526020604085019301915b818110613b68575050500390f35b82516001600160a01b0316845285945060209384019390920191600101613b5a565b9050346107f25760206003193601126107f2576020907fffffffff00000000000000000000000000000000000000000000000000000000613bc9613d4e565b167f36372b07000000000000000000000000000000000000000000000000000000008114908115613d24575b8115613cfa575b8115613cd0575b8115613c11575b5015158152f35b7faff2afbf00000000000000000000000000000000000000000000000000000000811491508115613ca6575b8115613c7c575b8115613c52575b5083613c0a565b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483613c4b565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150613c44565b7f940a15420000000000000000000000000000000000000000000000000000000081149150613c3d565b7f01ffc9a70000000000000000000000000000000000000000000000000000000081149150613c03565b7fa219a0250000000000000000000000000000000000000000000000000000000081149150613bfc565b7f8fd6a6ac0000000000000000000000000000000000000000000000000000000081149150613bf5565b600435907fffffffff0000000000000000000000000000000000000000000000000000000082168203613d7d57565b600080fd5b602435907fffffffff0000000000000000000000000000000000000000000000000000000082168203613d7d57565b604435907fffffffff0000000000000000000000000000000000000000000000000000000082168203613d7d57565b600435906001600160a01b0382168203613d7d57565b606435906001600160a01b0382168203613d7d57565b602435906001600160a01b0382168203613d7d57565b35906001600160a01b0382168203613d7d57565b6024359067ffffffffffffffff82168203613d7d57565b6004359067ffffffffffffffff82168203613d7d57565b359067ffffffffffffffff82168203613d7d57565b9181601f84011215613d7d5782359167ffffffffffffffff8311613d7d5760208381860195010111613d7d57565b919082519283825260005b848110613ef15750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201613eb2565b35908115158203613d7d57565b6020810190811067ffffffffffffffff821117613f2f57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff821117613f2f57604052565b60a0810190811067ffffffffffffffff821117613f2f57604052565b60e0810190811067ffffffffffffffff821117613f2f57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117613f2f57604052565b92919267ffffffffffffffff8211613f2f576040519161403b601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200184613fb2565b829481845281830111613d7d578281602093846000960137010152565b9080601f83011215613d7d5781602061407393359101613ff3565b90565b906040600319830112613d7d5760043567ffffffffffffffff81168103613d7d57916024359067ffffffffffffffff8211613d7d576140b791600401613e79565b9091565b6140739160206140d48351604084526040840190613ea7565b920151906020818403910152613ea7565b9181601f84011215613d7d5782359167ffffffffffffffff8311613d7d576020808501948460051b010111613d7d57565b9181601f84011215613d7d5782359167ffffffffffffffff8311613d7d576020808501948460081b010111613d7d57565b67ffffffffffffffff8111613f2f5760051b60200190565b8181029291811591840414171561417257565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b81156141ab570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b9190820391821161417257565b51906001600160a01b0382168203613d7d57565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b9295939091946001600160a01b03600854169586156144f657809760028710156144c7576001600160a01b0398614381957fffffffff00000000000000000000000000000000000000000000000000000000938961449d5767ffffffffffffffff82166000526010602052604060002090604051916142b883613f96565b549163ffffffff8316815263ffffffff8360201c16602082015263ffffffff8360401c16604082015263ffffffff8360601c16606082015260c061ffff8460801c169182608082015260ff60a082019561ffff8160901c16875260a01c1615159182910152614449575b50505067ffffffffffffffff905b6040519b8c997f06b859ef000000000000000000000000000000000000000000000000000000008b521660048a0152166024880152604487015216606485015260c0608485015260c48401916141fb565b928180600095869560a483015203915afa91821561443c5781926143a457505090565b9091503d8083833e6143b68183613fb2565b810190602081830312610a2d5780519067ffffffffffffffff8211610684570181601f82011215610a2d578051906143ed82614147565b936143fb6040519586613fb2565b82855260208086019360051b8301019384116107f65750602001905b8282106144245750505090565b60208091614431846141e7565b815201910190614417565b50604051903d90823e3d90fd5b92935067ffffffffffffffff9285871615614485575061271061447461ffff61447b9451168361415f565b04906141da565b915b903880614322565b6144979250614474612710918361415f565b9161447d565b67ffffffffffffffff9192506144c1906144bb612a3c36898b613ff3565b90614b77565b91614330565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b505050505050505060405161450c602082613fb2565b60008152600036813790565b90600182811c92168015614561575b602083101461453257565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691614527565b6000929181549161457b83614518565b80835292600181169081156145d1575060011461459757505050565b60009081526020812093945091925b8383106145b7575060209250010190565b6001816020929493945483858701015201910191906145a6565b905060209495507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091509291921683830152151560051b010190565b67ffffffffffffffff9092919261464b7fffffffff0000000000000000000000000000000000000000000000000000000060075460401b1685614c66565b16600052601060205260406000206040519061466682613f96565b549163ffffffff83169384835263ffffffff8460201c169384602085015263ffffffff8160401c169182604086015263ffffffff8260601c169081606087015261ffff8360801c169586608082015260ff61ffff8560901c16948560a084015260a01c16159060c08215910152614713577fffffffff000000000000000000000000000000000000000000000000000000001661470857505093929190600190565b959493509160019150565b5050505092505050600090600090600090600090600090565b6040519061473982613f7a565b60006080838281528260208201528260408201528260608201520152565b9060405161476481613f7a565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff916147cf61472c565b506147d861472c565b5061480c5716600052600d602052604060002090614073614800600261480561480086614757565b614e5d565b9401614757565b16908160005260096020526148276148006040600020614757565b91600052600a6020526140736148006040600020614757565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215613d7d570180359067ffffffffffffffff8211613d7d57602001918136038313613d7d57565b356001600160a01b0381168103613d7d5790565b3567ffffffffffffffff81168103613d7d5790565b9067ffffffffffffffff6140739216600052600d602052600560406000200190602081519101209060019160005201602052604060002054151590565b6040519061490482613f5e565b60606020838281520152565b80518210156149245760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b67ffffffffffffffff16600052600d6020526120e661407360046040600020016040519283809261456b565b91908110156149245760081b0190565b358015158103613d7d5790565b3561ffff81168103613d7d5790565b3563ffffffff81168103613d7d5790565b359063ffffffff82168203613d7d57565b359061ffff82168203613d7d57565b91908110156149245760051b0190565b35906fffffffffffffffffffffffffffffffff82168203613d7d57565b9190826060910312613d7d576040516060810181811067ffffffffffffffff821117613f2f576040526040614a5c818395614a4381613f06565b8552614a51602082016149ec565b6020860152016149ec565b910152565b6fffffffffffffffffffffffffffffffff614a9f60408093614a8281613f06565b1515865283614a93602083016149ec565b166020870152016149ec565b16910152565b818110614ab0575050565b60008155600101614aa5565b80518015614b2c57602003614aee578051602082810191830183900312613d7d57519060ff8211614aee575060ff1690565b611f6e906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190613ea7565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff821161417257565b60ff16604d811161417257600a0a90565b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff811692828414614c5f57828411614c355790614bbc91614b52565b91604d60ff8416118015614c1a575b614be457505090614bde61407392614b66565b9061415f565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b50614c2483614b66565b80156141ab57600019048411614bcb565b614c3e91614b52565b91604d60ff841611614be457505090614c5961407392614b66565b906141a1565b5050505090565b7fffffffff000000000000000000000000000000000000000000000000000000008116908115614d4857614c9981615461565b7dffff00000000000000000000000000000000000000000000000000000000601082811c9085901c1616614d485761ffff8360e01c168015918215614d37575b5050614ce3575050565b7fffffffff0000000000000000000000000000000000000000000000000000000092507fdf63778f000000000000000000000000000000000000000000000000000000006000526004521660245260446000fd5b60e01c61ffff161090503880614cd9565b505050565b6001600160a01b0316908115614e2e576001600160a01b0316918215614dff576000828152806020526040812054828110614dcc5791604082827fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef958760209652828652038282205586815280845220818154019055604051908152a3565b6064937fe450d38c0000000000000000000000000000000000000000000000000000000083949352600452602452604452fd5b7fec442f0500000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b7f96c6fd1e00000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b614e6561472c565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff8083511691614ec26020850193614ebc614eaf63ffffffff875116426141da565b856080890151169061415f565b90615454565b80821015614edb57505b16825263ffffffff4216905290565b9050614ecc565b90816020910312613d7d57518015158103613d7d5790565b6001600160a01b03600654163303614f0e57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b9080511561513c5767ffffffffffffffff8151602083012092169182600052600d602052614f6d816005604060002001615b7b565b156150f857600052600e6020526040600020815167ffffffffffffffff8111613f2f57614f9a8254614518565b601f81116150c6575b506020601f821160011461501e5791614ff8827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea959361500e95600091615013575b506000198260011b9260031b1c19161790565b9055604051918291602083526020830190613ea7565b0390a2565b905084015138614fe5565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b8181106150ae57509261500e9492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610615095575b5050811b019055611c50565b85015160001960f88460031b161c191690553880615089565b9192602060018192868a01518155019401920161504e565b6150f290836000526020600020601f840160051c81019160208510610e4157601f0160051c0190614aa5565b38614fa3565b5090611f6e6040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190613ea7565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016602082015260208152614073604082613fb2565b815191929115615325576fffffffffffffffffffffffffffffffff6040840151166fffffffffffffffffffffffffffffffff602085015116106152c2576152c091925b805182547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff0000000000000000000000000000000000000000161782556020810151825460409290920151608090811b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff9290921691821760018501557fffffffffffffffffffffffff0000000000000000000000000000000000000000909216174290911b73ffffffff0000000000000000000000000000000016179055565b565b606483615323604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff604084015116158015906153b4575b615353576152c091926151e4565b606483615323604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff6020840151161515615345565b906127109167ffffffffffffffff6153ed602083016148a5565b166000908152601060205260409020917fffffffff00000000000000000000000000000000000000000000000000000000161561543e57606061ffff61543a935460901c1691013561415f565b0490565b606061ffff61543a935460801c1691013561415f565b9190820180921161417257565b7fffffffff000000000000000000000000000000000000000000000000000000008116908115615538577dffff0000000000000000000000000000000000000000000000000000000081161561552f5760ff60015b169060f01c806154f9575b506001036154cc5750565b7fc512f96c0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60005b6010811061550a57506154c1565b6001811b821661551d575b6001016154fc565b91600181018091116141725791615515565b60ff60006154b6565b5050565b6001600160a01b0316801591821561560157907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef602083615581600095600254615454565b6002555b8060025403600255604051908152a361559a57565b7f0000000000000000000000000000000000000000000000000000000000000000806155c35750565b6002548181116155d1575050565b7fea0582460000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b8160005260006020526040600020548181106156515760208284937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef936000968752868452036040862055615585565b827fe450d38c0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60206001600160a01b036000936156bf86600254615454565b6002551693841584146157055780600254036002555b604051908152a37f0000000000000000000000000000000000000000000000000000000000000000806155c35750565b848452838252604084208181540190556156d5565b9167ffffffffffffffff7f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c92169283600052600d60205261576381836002604060002001615bd0565b604080516001600160a01b0390921682526020820192909252908190810161500e565b91909167ffffffffffffffff83169283600052600a60205260ff60406000205460a01c16156157eb5750907fc6735cd4fa2bbe7b203b1682936e6ee61bc1702464bbbd12abb6630229d9a5f99183600052600a60205261576381836040600020615bd0565b906152c0935061571a565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe182360301811215613d7d57016020813591019167ffffffffffffffff8211613d7d578136038313613d7d57565b9167ffffffffffffffff7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894492169283600052600d60205261576381836040600020615bd0565b91909167ffffffffffffffff83169283600052600960205260ff60406000205460a01c16156158f15750907f28d6c52e2b0b7587b0d195539fbe6af984b28791aca4d2cc0844244e38bce29e9183600052600960205261576381836040600020615bd0565b906152c09350615846565b906040519182815491828252602082019060005260206000209260005b81811061592e5750506152c092500383613fb2565b8454835260019485019487945060209093019201615919565b80548210156149245760005260206000200190600090565b6000818152600c60205260409020548015615a5857600019810181811161417257600b5490600019820191821161417257818103615a07575b505050600b5480156159d857600019016159b381600b615947565b60001982549160031b1b19169055600b55600052600c60205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b615a40615a18615a2993600b615947565b90549060031b1c928392600b615947565b81939154906000199060031b92831b921b19161790565b9055600052600c602052604060002055388080615998565b5050600090565b9060018201918160005282602052604060002054801515600014615b1257600019810181811161417257825490600019820191821161417257818103615adb575b505050805480156159d8576000190190615aba8282615947565b60001982549160031b1b191690555560005260205260006040812055600190565b615afb615aeb615a299386615947565b90549060031b1c92839286615947565b905560005283602052604060002055388080615aa0565b50505050600090565b80600052600c60205260406000205415600014615b7557600b5468010000000000000000811015613f2f57615b5c615a29826001859401600b55600b615947565b9055600b5490600052600c602052604060002055600190565b50600090565b6000828152600182016020526040902054615a585780549068010000000000000000821015613f2f5782615bb9615a29846001809601855584615947565b905580549260005201602052604060002055600190565b9182549060ff8260a01c16158015615e21575b615e1b576fffffffffffffffffffffffffffffffff82169160018501908154615c2863ffffffff6fffffffffffffffffffffffffffffffff83169360801c16426141da565b9081615d7d575b5050848110615d3e5750838310615c89575050615c5e6fffffffffffffffffffffffffffffffff9283926141da565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b9190915460801c928315615cfd5781615ca1916141da565b9260001981019080821161417257615cc4615cc9926001600160a01b0396615454565b6141a1565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b6001600160a01b0383837fd0c8d23a000000000000000000000000000000000000000000000000000000006000526000196004526024521660445260646000fd5b82856001600160a01b03927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611615df157615d9892614ebc9160801c9061415f565b80841015615dec5750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178655923880615c2f565b615da3565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b508215615be356fea164736f6c634300081a000a",
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

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) GetAllowedFinalityConfig(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "getAllowedFinalityConfig")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) GetAllowedFinalityConfig() ([4]byte, error) {
	return _CrossChainPoolToken.Contract.GetAllowedFinalityConfig(&_CrossChainPoolToken.CallOpts)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) GetAllowedFinalityConfig() ([4]byte, error) {
	return _CrossChainPoolToken.Contract.GetAllowedFinalityConfig(&_CrossChainPoolToken.CallOpts)
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

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64, fastFinality bool) (GetCurrentRateLimiterState,

	error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "getCurrentRateLimiterState", remoteChainSelector, fastFinality)

	outstruct := new(GetCurrentRateLimiterState)
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutboundRateLimiterState = *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)
	outstruct.InboundRateLimiterState = *abi.ConvertType(out[1], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return *outstruct, err

}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) GetCurrentRateLimiterState(remoteChainSelector uint64, fastFinality bool) (GetCurrentRateLimiterState,

	error) {
	return _CrossChainPoolToken.Contract.GetCurrentRateLimiterState(&_CrossChainPoolToken.CallOpts, remoteChainSelector, fastFinality)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) GetCurrentRateLimiterState(remoteChainSelector uint64, fastFinality bool) (GetCurrentRateLimiterState,

	error) {
	return _CrossChainPoolToken.Contract.GetCurrentRateLimiterState(&_CrossChainPoolToken.CallOpts, remoteChainSelector, fastFinality)
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

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, requestedFinalityConfig [4]byte, arg5 []byte) (GetFee,

	error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "getFee", arg0, destChainSelector, arg2, arg3, requestedFinalityConfig, arg5)

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

func (_CrossChainPoolToken *CrossChainPoolTokenSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, requestedFinalityConfig [4]byte, arg5 []byte) (GetFee,

	error) {
	return _CrossChainPoolToken.Contract.GetFee(&_CrossChainPoolToken.CallOpts, arg0, destChainSelector, arg2, arg3, requestedFinalityConfig, arg5)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, requestedFinalityConfig [4]byte, arg5 []byte) (GetFee,

	error) {
	return _CrossChainPoolToken.Contract.GetFee(&_CrossChainPoolToken.CallOpts, arg0, destChainSelector, arg2, arg3, requestedFinalityConfig, arg5)
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

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) GetRequiredCCVs(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, sourceDenominatedAmount *big.Int, requestedFinalityConfig [4]byte, extraData []byte, direction uint8) ([]common.Address, error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "getRequiredCCVs", localToken, remoteChainSelector, sourceDenominatedAmount, requestedFinalityConfig, extraData, direction)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, sourceDenominatedAmount *big.Int, requestedFinalityConfig [4]byte, extraData []byte, direction uint8) ([]common.Address, error) {
	return _CrossChainPoolToken.Contract.GetRequiredCCVs(&_CrossChainPoolToken.CallOpts, localToken, remoteChainSelector, sourceDenominatedAmount, requestedFinalityConfig, extraData, direction)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, sourceDenominatedAmount *big.Int, requestedFinalityConfig [4]byte, extraData []byte, direction uint8) ([]common.Address, error) {
	return _CrossChainPoolToken.Contract.GetRequiredCCVs(&_CrossChainPoolToken.CallOpts, localToken, remoteChainSelector, sourceDenominatedAmount, requestedFinalityConfig, extraData, direction)
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

func (_CrossChainPoolToken *CrossChainPoolTokenCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 [4]byte, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _CrossChainPoolToken.contract.Call(opts, &out, "getTokenTransferFeeConfig", arg0, destChainSelector, arg2, arg3)

	if err != nil {
		return *new(IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolV2TokenTransferFeeConfig)).(*IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 [4]byte, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _CrossChainPoolToken.Contract.GetTokenTransferFeeConfig(&_CrossChainPoolToken.CallOpts, arg0, destChainSelector, arg2, arg3)
}

func (_CrossChainPoolToken *CrossChainPoolTokenCallerSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 [4]byte, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
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

func (_CrossChainPoolToken *CrossChainPoolTokenTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, requestedFinalityConfig [4]byte, tokenArgs []byte) (*types.Transaction, error) {
	return _CrossChainPoolToken.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, requestedFinalityConfig, tokenArgs)
}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, requestedFinalityConfig [4]byte, tokenArgs []byte) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.LockOrBurn0(&_CrossChainPoolToken.TransactOpts, lockOrBurnIn, requestedFinalityConfig, tokenArgs)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, requestedFinalityConfig [4]byte, tokenArgs []byte) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.LockOrBurn0(&_CrossChainPoolToken.TransactOpts, lockOrBurnIn, requestedFinalityConfig, tokenArgs)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, requestedFinalityConfig [4]byte) (*types.Transaction, error) {
	return _CrossChainPoolToken.contract.Transact(opts, "releaseOrMint", releaseOrMintIn, requestedFinalityConfig)
}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1, requestedFinalityConfig [4]byte) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.ReleaseOrMint(&_CrossChainPoolToken.TransactOpts, releaseOrMintIn, requestedFinalityConfig)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1, requestedFinalityConfig [4]byte) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.ReleaseOrMint(&_CrossChainPoolToken.TransactOpts, releaseOrMintIn, requestedFinalityConfig)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _CrossChainPoolToken.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn)
}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.ReleaseOrMint0(&_CrossChainPoolToken.TransactOpts, releaseOrMintIn)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.ReleaseOrMint0(&_CrossChainPoolToken.TransactOpts, releaseOrMintIn)
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

func (_CrossChainPoolToken *CrossChainPoolTokenTransactor) SetAllowedFinalityConfig(opts *bind.TransactOpts, allowedFinality [4]byte) (*types.Transaction, error) {
	return _CrossChainPoolToken.contract.Transact(opts, "setAllowedFinalityConfig", allowedFinality)
}

func (_CrossChainPoolToken *CrossChainPoolTokenSession) SetAllowedFinalityConfig(allowedFinality [4]byte) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.SetAllowedFinalityConfig(&_CrossChainPoolToken.TransactOpts, allowedFinality)
}

func (_CrossChainPoolToken *CrossChainPoolTokenTransactorSession) SetAllowedFinalityConfig(allowedFinality [4]byte) (*types.Transaction, error) {
	return _CrossChainPoolToken.Contract.SetAllowedFinalityConfig(&_CrossChainPoolToken.TransactOpts, allowedFinality)
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

type CrossChainPoolTokenFastFinalityInboundRateLimitConsumedIterator struct {
	Event *CrossChainPoolTokenFastFinalityInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainPoolTokenFastFinalityInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainPoolTokenFastFinalityInboundRateLimitConsumed)
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
		it.Event = new(CrossChainPoolTokenFastFinalityInboundRateLimitConsumed)
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

func (it *CrossChainPoolTokenFastFinalityInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *CrossChainPoolTokenFastFinalityInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainPoolTokenFastFinalityInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) FilterFastFinalityInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CrossChainPoolTokenFastFinalityInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.FilterLogs(opts, "FastFinalityInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainPoolTokenFastFinalityInboundRateLimitConsumedIterator{contract: _CrossChainPoolToken.contract, event: "FastFinalityInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) WatchFastFinalityInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenFastFinalityInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.WatchLogs(opts, "FastFinalityInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainPoolTokenFastFinalityInboundRateLimitConsumed)
				if err := _CrossChainPoolToken.contract.UnpackLog(event, "FastFinalityInboundRateLimitConsumed", log); err != nil {
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

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) ParseFastFinalityInboundRateLimitConsumed(log types.Log) (*CrossChainPoolTokenFastFinalityInboundRateLimitConsumed, error) {
	event := new(CrossChainPoolTokenFastFinalityInboundRateLimitConsumed)
	if err := _CrossChainPoolToken.contract.UnpackLog(event, "FastFinalityInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CrossChainPoolTokenFastFinalityOutboundRateLimitConsumedIterator struct {
	Event *CrossChainPoolTokenFastFinalityOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainPoolTokenFastFinalityOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainPoolTokenFastFinalityOutboundRateLimitConsumed)
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
		it.Event = new(CrossChainPoolTokenFastFinalityOutboundRateLimitConsumed)
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

func (it *CrossChainPoolTokenFastFinalityOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *CrossChainPoolTokenFastFinalityOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainPoolTokenFastFinalityOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) FilterFastFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CrossChainPoolTokenFastFinalityOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.FilterLogs(opts, "FastFinalityOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainPoolTokenFastFinalityOutboundRateLimitConsumedIterator{contract: _CrossChainPoolToken.contract, event: "FastFinalityOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) WatchFastFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenFastFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CrossChainPoolToken.contract.WatchLogs(opts, "FastFinalityOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainPoolTokenFastFinalityOutboundRateLimitConsumed)
				if err := _CrossChainPoolToken.contract.UnpackLog(event, "FastFinalityOutboundRateLimitConsumed", log); err != nil {
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

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) ParseFastFinalityOutboundRateLimitConsumed(log types.Log) (*CrossChainPoolTokenFastFinalityOutboundRateLimitConsumed, error) {
	event := new(CrossChainPoolTokenFastFinalityOutboundRateLimitConsumed)
	if err := _CrossChainPoolToken.contract.UnpackLog(event, "FastFinalityOutboundRateLimitConsumed", log); err != nil {
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

type CrossChainPoolTokenFinalityConfigSetIterator struct {
	Event *CrossChainPoolTokenFinalityConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CrossChainPoolTokenFinalityConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainPoolTokenFinalityConfigSet)
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
		it.Event = new(CrossChainPoolTokenFinalityConfigSet)
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

func (it *CrossChainPoolTokenFinalityConfigSetIterator) Error() error {
	return it.fail
}

func (it *CrossChainPoolTokenFinalityConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CrossChainPoolTokenFinalityConfigSet struct {
	AllowedFinality [4]byte
	Raw             types.Log
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) FilterFinalityConfigSet(opts *bind.FilterOpts) (*CrossChainPoolTokenFinalityConfigSetIterator, error) {

	logs, sub, err := _CrossChainPoolToken.contract.FilterLogs(opts, "FinalityConfigSet")
	if err != nil {
		return nil, err
	}
	return &CrossChainPoolTokenFinalityConfigSetIterator{contract: _CrossChainPoolToken.contract, event: "FinalityConfigSet", logs: logs, sub: sub}, nil
}

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) WatchFinalityConfigSet(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenFinalityConfigSet) (event.Subscription, error) {

	logs, sub, err := _CrossChainPoolToken.contract.WatchLogs(opts, "FinalityConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CrossChainPoolTokenFinalityConfigSet)
				if err := _CrossChainPoolToken.contract.UnpackLog(event, "FinalityConfigSet", log); err != nil {
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

func (_CrossChainPoolToken *CrossChainPoolTokenFilterer) ParseFinalityConfigSet(log types.Log) (*CrossChainPoolTokenFinalityConfigSet, error) {
	event := new(CrossChainPoolTokenFinalityConfigSet)
	if err := _CrossChainPoolToken.contract.UnpackLog(event, "FinalityConfigSet", log); err != nil {
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
	FastFinality              bool
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

func (CrossChainPoolTokenDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x3f1036e85d016a93254a0b1415844f79b85424959d90ae5ad51ce8f4533fe701")
}

func (CrossChainPoolTokenFastFinalityInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xc6735cd4fa2bbe7b203b1682936e6ee61bc1702464bbbd12abb6630229d9a5f9")
}

func (CrossChainPoolTokenFastFinalityOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x28d6c52e2b0b7587b0d195539fbe6af984b28791aca4d2cc0844244e38bce29e")
}

func (CrossChainPoolTokenFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (CrossChainPoolTokenFinalityConfigSet) Topic() common.Hash {
	return common.HexToHash("0x307cf716eade81675bea3ccb6917b0f91baa2160056765d9a83d76f819caf06a")
}

func (CrossChainPoolTokenInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (CrossChainPoolTokenLockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
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

	GetAllowedFinalityConfig(opts *bind.CallOpts) ([4]byte, error)

	GetCCIPAdmin(opts *bind.CallOpts) (common.Address, error)

	GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64, fastFinality bool) (GetCurrentRateLimiterState,

		error)

	GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

		error)

	GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, requestedFinalityConfig [4]byte, arg5 []byte) (GetFee,

		error)

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRequiredCCVs(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, sourceDenominatedAmount *big.Int, requestedFinalityConfig [4]byte, extraData []byte, direction uint8) ([]common.Address, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetSupportedChains(opts *bind.CallOpts) ([]uint64, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	GetTokenDecimals(opts *bind.CallOpts) (uint8, error)

	GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 [4]byte, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error)

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

	LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, requestedFinalityConfig [4]byte, tokenArgs []byte) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, requestedFinalityConfig [4]byte) (*types.Transaction, error)

	ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetAllowedFinalityConfig(opts *bind.TransactOpts, allowedFinality [4]byte) (*types.Transaction, error)

	SetCCIPAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error)

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

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*CrossChainPoolTokenDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*CrossChainPoolTokenDynamicConfigSet, error)

	FilterFastFinalityInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CrossChainPoolTokenFastFinalityInboundRateLimitConsumedIterator, error)

	WatchFastFinalityInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenFastFinalityInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseFastFinalityInboundRateLimitConsumed(log types.Log) (*CrossChainPoolTokenFastFinalityInboundRateLimitConsumed, error)

	FilterFastFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CrossChainPoolTokenFastFinalityOutboundRateLimitConsumedIterator, error)

	WatchFastFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenFastFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseFastFinalityOutboundRateLimitConsumed(log types.Log) (*CrossChainPoolTokenFastFinalityOutboundRateLimitConsumed, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*CrossChainPoolTokenFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*CrossChainPoolTokenFeeTokenWithdrawn, error)

	FilterFinalityConfigSet(opts *bind.FilterOpts) (*CrossChainPoolTokenFinalityConfigSetIterator, error)

	WatchFinalityConfigSet(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenFinalityConfigSet) (event.Subscription, error)

	ParseFinalityConfigSet(log types.Log) (*CrossChainPoolTokenFinalityConfigSet, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CrossChainPoolTokenInboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*CrossChainPoolTokenInboundRateLimitConsumed, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CrossChainPoolTokenLockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *CrossChainPoolTokenLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*CrossChainPoolTokenLockedOrBurned, error)

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
