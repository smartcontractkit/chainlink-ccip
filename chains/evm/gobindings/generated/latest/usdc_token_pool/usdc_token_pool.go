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

type IPoolV2TokenTransferFeeConfig struct {
	DestGasOverhead                        uint32
	DestBytesOverhead                      uint32
	DefaultBlockConfirmationFeeUSDCents    uint32
	CustomBlockConfirmationFeeUSDCents     uint32
	DefaultBlockConfirmationTransferFeeBps uint16
	CustomBlockConfirmationTransferFeeBps  uint16
	IsEnabled                              bool
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

type TokenPoolCCVConfigArg struct {
	RemoteChainSelector             uint64
	OutboundCCVs                    []common.Address
	OutboundCCVsToAddAboveThreshold []common.Address
	InboundCCVs                     []common.Address
	InboundCCVsToAddAboveThreshold  []common.Address
}

type TokenPoolChainUpdate struct {
	RemoteChainSelector       uint64
	RemotePoolAddresses       [][]byte
	RemoteTokenAddress        []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
}

type TokenPoolCustomBlockConfirmationRateLimitConfigArgs struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
}

type TokenPoolTokenTransferFeeConfigArgs struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
}

type USDCTokenPoolDomain struct {
	AllowedCaller                 [32]byte
	MintRecipient                 [32]byte
	DomainIdentifier              uint32
	Enabled                       bool
	UseLegacySourcePoolDataFormat bool
}

type USDCTokenPoolDomainUpdate struct {
	AllowedCaller                 [32]byte
	MintRecipient                 [32]byte
	DomainIdentifier              uint32
	DestChainSelector             uint64
	Enabled                       bool
	UseLegacySourcePoolDataFormat bool
}

var USDCTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"internalType\":\"contract ITokenMessenger\"},{\"name\":\"cctpMessageTransmitterProxy\",\"type\":\"address\",\"internalType\":\"contract CCTPMessageTransmitterProxy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"supportedUSDCVersion\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"struct AuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCustomBlockConfirmationConfigUpdates\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CustomBlockConfirmationRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"destToUseDefaultFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAccumulatedFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getConfiguredMinBlockConfirmation\",\"inputs\":[],\"outputs\":[{\"name\":\"blockConfirmationConfigured\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentCustomBlockConfirmationRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDomain\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct USDCTokenPool.Domain\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useLegacySourcePoolDataFormat\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_localDomainIdentifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_messageTransmitterProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract CCTPMessageTransmitterProxy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_supportedUSDCVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_tokenMessenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract ITokenMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct RateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct RateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustomBlockConfirmationRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CustomBlockConfirmationRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDomains\",\"inputs\":[{\"name\":\"domains\",\"type\":\"tuple[]\",\"internalType\":\"struct USDCTokenPool.DomainUpdate[]\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useLegacySourcePoolDataFormat\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationUpdated\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DomainsSet\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"struct USDCTokenPool.DomainUpdate[]\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useLegacySourcePoolDataFormat\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeeWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidDestinationDomain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidDomain\",\"inputs\":[{\"name\":\"domain\",\"type\":\"tuple\",\"internalType\":\"struct USDCTokenPool.DomainUpdate\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useLegacySourcePoolDataFormat\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"type\":\"error\",\"name\":\"InvalidMessageLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmation\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmationConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNonce\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"got\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidPreviousPool\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourceDomain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenMessengerVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidTransmitterInProxy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenTransferFeeConfigNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnknownDomain\",\"inputs\":[{\"name\":\"domain\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"UnlockingUSDCFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610180806040523461068d5761772b803803809161001d8285610a56565b8339810160e08282031261068d5781516001600160a01b0381169283820361068d576020810151936001600160a01b0385169081860361068d576040830151956001600160a01b0387169586880361068d5760608501516001600160401b03811161068d5785019080601f8301121561068d578151916001600160401b038311610692578260051b9060208201936100b86040519586610a56565b845260208085019282010192831161068d57602001905b828210610a3e575050506100e560808601610a79565b946100fe60c06100f760a08401610a79565b9201610a8d565b95602099604051996101108c8c610a56565b60008b5260003681373315610a2d57600180546001600160a01b0319163317905580158015610a1c575b8015610a0b575b6109fa576004928c9260805260c0526040519283809263313ce56760e01b82525afa80916000916109c3575b509061099f575b50600660a052600480546001600160a01b0319166001600160a01b03929092169190911790558051151560e0819052610871575b50604051946101b78887610a56565b60008652600036813760408051979088016001600160401b03811189821017610692576040528752858888015260005b865181101561024e576001906001600160a01b03610205828a610a9e565b51168a61021182610b1f565b61021e575b5050016101e7565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a1388a610216565b5087955086519360005b85518110156102c9576001600160a01b036102738288610a9e565b51169081156102b8577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef89836102aa600195610c3b565b50604051908152a101610258565b6342bcdf7f60e11b60005260046000fd5b50869394508561010052841561086057604051632c12192160e01b81528481600481895afa9081156106f85760009161082b575b5060405163054fd4d560e41b81526001600160a01b039190911691908581600481865afa9081156106f8576000916107f6575b5063ffffffff8061010051169116908082036107df575050604051639cdbb18160e01b815285816004818a5afa9081156106f8576000916107aa575b5063ffffffff80610100511691169080820361079357505084600491604051928380926367e0ed8360e11b82525afa80156106f857829160009161074a575b506001600160a01b03160361073957600492849261012052610140526040519283809263234d8e3d60e21b82525afa9081156106f857600091610704575b506101605260805161012051604051636eb1769f60e11b81523060048201526001600160a01b03918216602482018190529492909116908381604481855afa9081156106f8576000916106cb575b5060001981018091116106b557604051908482019563095ea7b360e01b8752602483015260448201526044815261046f606482610a56565b6000806040968751936104828986610a56565b8785527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656488860152519082865af13d156106a8573d906001600160401b0382116106925786516104ef9490926104e1601f8201601f1916890185610a56565b83523d60008885013e610d08565b805180610614575b847f2e902d38f15b233cbb63711add0fca4545334d3a169d60c0a616494d7eea954485858351908152a1516169529081610dd98239608051818181610655015281816108430152818161089401528181610c2e01528181611d0a0152818161235f0152818161374301528181615f310152616474015260a051818181610987015281816123cf01528181614f1a0152614f91015260c0518181816129670152818161489e01528181614a28015281816157ca01526158c9015260e0518181816110a301528181612ae901526163b80152610100518181816107ff0152614c9101526101205181818161133b0152611d58015261014051818181610ba30152611ba101526101605181818161151401528181611e230152614cfe0152f35b8184918101031261068d5782015180159081150361068d576106375783806104f7565b50608491519062461bcd60e51b82526004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b916104ef92606091610d08565b634e487b7160e01b600052601160045260246000fd5b90508381813d83116106f1575b6106e28183610a56565b8101031261068d575185610437565b503d6106d8565b6040513d6000823e3d90fd5b90508181813d8311610732575b61071b8183610a56565b8101031261068d5761072c90610a8d565b836103e9565b503d610711565b632a32133b60e11b60005260046000fd5b9091508581813d831161078c575b6107628183610a56565b810103126107885751906001600160a01b038216820361078557508190876103ab565b80fd5b5080fd5b503d610758565b633785f8f160e01b60005260045260245260446000fd5b90508581813d83116107d8575b6107c18183610a56565b8101031261068d576107d290610a8d565b8761036c565b503d6107b7565b63960693cd60e01b60005260045260245260446000fd5b90508581813d8311610824575b61080d8183610a56565b8101031261068d5761081e90610a8d565b87610330565b503d610803565b90508481813d8311610859575b6108428183610a56565b8101031261068d5761085390610a79565b866102fd565b503d610838565b6306b7c75960e31b60005260046000fd5b60405192969295919490936108868988610a56565b60008752600036813760e0511561098e5760005b8751811015610901576001906001600160a01b036108b8828b610a9e565b51168b6108c482610c74565b6108d1575b50500161089a565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a1388b6108c9565b509193955091939560005b8651811015610980576001906001600160a01b0361092a828a610a9e565b5116801561097a578a61093c82610bfc565b61094a575b50505b0161090c565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a1388a610941565b50610944565b5091939590929450386101a8565b6335f4a7b360e01b60005260046000fd5b60ff1660068114610174576332ad3e0760e11b600052600660045260245260446000fd5b8b81813d83116109f3575b6109d88183610a56565b8101031261078857519060ff8216820361078557503861016d565b503d6109ce565b630a64406560e11b60005260046000fd5b506001600160a01b03831615610141565b506001600160a01b0384161561013a565b639b15e16f60e01b60005260046000fd5b60208091610a4b84610a79565b8152019101906100cf565b601f909101601f19168101906001600160401b0382119082101761069257604052565b51906001600160a01b038216820361068d57565b519063ffffffff8216820361068d57565b8051821015610ab25760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b8054821015610ab25760005260206000200190600090565b80548015610b09576000190190610af78282610ac8565b8154906000199060031b1b1916905555565b634e487b7160e01b600052603160045260246000fd5b6000818152601160205260409020548015610bca5760001981018181116106b5576010546000198101919082116106b557808203610b79575b505050610b656010610ae0565b600052601160205260006040812055600190565b610bb2610b8a610b9b936010610ac8565b90549060031b1c9283926010610ac8565b819391549060031b91821b91600019901b19161790565b90556000526011602052604060002055388080610b58565b5050600090565b805490680100000000000000008210156106925781610b9b916001610bf894018155610ac8565b9055565b80600052600360205260406000205415600014610c3557610c1e816002610bd1565b600254906000526003602052604060002055600190565b50600090565b80600052601160205260406000205415600014610c3557610c5d816010610bd1565b601054906000526011602052604060002055600190565b6000818152600360205260409020548015610bca5760001981018181116106b5576002546000198101919082116106b557818103610cce575b505050610cba6002610ae0565b600052600360205260006040812055600190565b610cf0610cdf610b9b936002610ac8565b90549060031b1c9283926002610ac8565b90556000526003602052604060002055388080610cad565b91929015610d6a5750815115610d1c575090565b3b15610d255790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b825190915015610d7d5750805190602001fd5b6040519062461bcd60e51b8252602060048301528181519182602483015260005b838110610dc05750508160006044809484010152601f80199101168101030190fd5b60208282018101516044878401015285935001610d9e56fe6080604052600436101561001257600080fd5b60003560e01c806301ffc9a7146103375780630574d6a614610332578063164e68de1461032d578063181f5a7714610328578063212a052e1461032357806321df0da71461031e578063240028e8146103195780632451a6271461031457806324f65ee71461030f57806337b192471461030a5780633907753714610305578063489a68f2146103005780634ac8bd5f146102fb5780634c5ef0ed146102f657806354c8a4f3146102f157806359152aad146102ec5780635df45a37146102e75780636155cda0146102e257806362ddd3c4146102dd578063698c2c66146102d85780636b716b0d146102d35780636d3d1a58146102ce5780637437ff9f146102c957806379ba5097146102c45780637d54534e146102bf5780638926f54f146102ba57806389720a62146102b55780638da5cb5b146102b057806391a2749a146102ab578063962d4020146102a65780639751f884146102a157806398db96431461029c5780639a4575b914610297578063a42a7b8b14610292578063a7cd63b71461028d578063acfecf9114610288578063b1c71c6514610283578063b79465801461027e578063bb6bb5a714610279578063c4bffe2b14610274578063cf7401f31461026f578063d966866b1461026a578063dc04fa1f14610265578063dc0bd97114610260578063ded8d9561461025b578063dfadfa3514610256578063e0351e1314610251578063e8a1da171461024c578063f2fde38b146102475763f37c77fe1461024257600080fd5b612fb3565b612ef7565b612b0e565b612ad1565b612a07565b61298b565b612947565b6128c5565b612717565b612653565b612505565b612459565b612422565b6122c4565b612181565b612127565b612043565b611c00565b611b81565b611b0b565b6118ed565b611828565b611778565b611706565b6116c7565b611643565b611592565b61155f565b611538565b6114f7565b61141b565b611398565b61131b565b6112f8565b61124c565b611071565b610fa5565b610db6565b610ce6565b610acb565b610a2e565b61096d565b610907565b610867565b610823565b6107e2565b610783565b61058c565b6104f2565b34610478576020600319360112610478576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361047857807ff208a58f000000000000000000000000000000000000000000000000000000006103cc921490811561044e575b8115610424575b81156103fa575b81156103d0575b5060405190151581529081906020820190565b0390f35b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014386103b9565b7f0e64dd2900000000000000000000000000000000000000000000000000000000811491506103b2565b7ff57e5f9400000000000000000000000000000000000000000000000000000000811491506103ab565b7faff2afbf00000000000000000000000000000000000000000000000000000000811491506103a4565b600080fd5b67ffffffffffffffff81160361047857565b359061049a8261047d565b565b6001600160a01b0381160361047857565b61ffff81160361047857565b359061049a826104ad565b9181601f840112156104785782359167ffffffffffffffff8311610478576020838186019501011161047857565b346104785760c06003193601126104785760043561050f8161047d565b61051a60243561049c565b61052560643561049c565b608435610531816104ad565b60a4359167ffffffffffffffff83116104785763ffffffff610567819361ffff936105608736906004016104c4565b505061305b565b6040805194855296909216602084015290921693810193909352166060820152608090f35b34610478576020600319360112610478576004356105a98161049c565b6105b16147bc565b6105b9613707565b90816105c157005b6001600160a01b037f738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b859599916107006040516106e560208201917fa9059cbb000000000000000000000000000000000000000000000000000000008352610652816106448a8860248401602090939291936001600160a01b0360408201951681520152565b03601f198101835282610ed8565b857f0000000000000000000000000000000000000000000000000000000000000000166000806040958651946106888887610ed8565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020870152519082855af13d15610726573d916106c983610f37565b926106d687519485610ed8565b83523d6000602085013e616881565b805180610705575b5050519283921694829190602083019252565b0390a2005b8160208061071a9361071f95010191016132c6565b615d2c565b38806106ed565b606091616881565b919082519283825260005b84811061075a575050601f19601f8460006020809697860101520116010190565b80602080928401015182828601015201610739565b90602061078092818152019061072e565b90565b34610478576000600319360112610478576103cc60408051906107a68183610ed8565b601782527f55534443546f6b656e506f6f6c20312e372e302d64657600000000000000000060208301525191829160208352602083019061072e565b3461047857600060031936011261047857602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346104785760006003193601126104785760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b346104785760206003193601126104785760206108ba6004356108898161049c565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691161490565b6040519015158152f35b602060408183019282815284518094520192019060005b8181106108e85750505090565b82516001600160a01b03168452602093840193909201916001016108db565b34610478576000600319360112610478576040516010548082526020820190601060005260206000209060005b818110610957576103cc8561094b81870382610ed8565b604051918291826108c4565b8254845260209093019260019283019201610934565b3461047857600060031936011261047857602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b908160a09103126104785790565b61049a9092919260c08060e083019563ffffffff815116845263ffffffff602082015116602085015263ffffffff604082015116604085015263ffffffff606082015116606085015261ffff6080820151166080850152610a2560a082015160a086019061ffff169052565b01511515910152565b346104785760a060031936011261047857610a4a60043561049c565b602435610a568161047d565b60443567ffffffffffffffff811161047857610a769036906004016109ab565b50610a826064356104ad565b60843567ffffffffffffffff8111610478576103cc91610aa9610ab09236906004016104c4565b505061311b565b604051918291826109b9565b90816101009103126104785790565b346104785760206003193601126104785760043567ffffffffffffffff811161047857610afc903690600401610abc565b610b046131e9565b506060810135610b1481836147fa565b610b956020610b31610b2960e08601866131fc565b81019061324d565b610b5a610b53610b4e610b4760c08901896131fc565b3691610f53565b614bd1565b8251614c7e565b8181519101519060405193849283927f57ecfd28000000000000000000000000000000000000000000000000000000008452600484016132db565b038160006001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165af1908115610ce157600091610cb2575b5015610c8857817ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc067ffffffffffffffff610c206040610c1960206103cc980161330c565b9401613316565b604080516001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081168252336020830152929092169082015260608101859052921691608090a2610c75610efb565b8190526040519081529081906020820190565b7fbf969f220000000000000000000000000000000000000000000000000000000060005260046000fd5b610cd4915060203d602011610cda575b610ccc8183610ed8565b8101906132c6565b38610bd4565b503d610cc2565b613300565b346104785760406003193601126104785760043567ffffffffffffffff811161047857610d1a6103cc913690600401610abc565b60243590610d27826104ad565b6000604051610d3581610e2b565b52610d67610d5f6060830135610d59610d54610b4760c08701876131fc565b614e71565b90614f8e565b9283836149ea565b7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc067ffffffffffffffff610c2060206040850194610da5863561049c565b013593610db18561047d565b613316565b346104785760206003193601126104785760043567ffffffffffffffff8111610478573660238201121561047857806004013567ffffffffffffffff81116104785736602460c0830284010111610478576024610e139201613320565b005b634e487b7160e01b600052604160045260246000fd5b6020810190811067ffffffffffffffff821117610e4757604052565b610e15565b6040810190811067ffffffffffffffff821117610e4757604052565b6060810190811067ffffffffffffffff821117610e4757604052565b60a0810190811067ffffffffffffffff821117610e4757604052565b60e0810190811067ffffffffffffffff821117610e4757604052565b60c0810190811067ffffffffffffffff821117610e4757604052565b90601f601f19910116810190811067ffffffffffffffff821117610e4757604052565b6040519061049a602083610ed8565b6040519061049a604083610ed8565b6040519061049a608083610ed8565b6040519061049a60a083610ed8565b67ffffffffffffffff8111610e4757601f01601f191660200190565b929192610f5f82610f37565b91610f6d6040519384610ed8565b829481845281830111610478578281602093846000960137010152565b9080601f830112156104785781602061078093359101610f53565b3461047857604060031936011261047857600435610fc28161047d565b60243567ffffffffffffffff811161047857602091610fe86108ba923690600401610f8a565b906136bb565b9181601f840112156104785782359167ffffffffffffffff8311610478576020808501948460051b01011161047857565b60406003198201126104785760043567ffffffffffffffff8111610478578161104a91600401610fee565b929092916024359067ffffffffffffffff82116104785761106d91600401610fee565b9091565b34610478576110996110a16110853661101f565b94916110929391936147bc565b36916117b7565b9236916117b7565b7f0000000000000000000000000000000000000000000000000000000000000000156111f15760005b825181101561115b57806110f06110e360019386613ac0565b516001600160a01b031690565b61111261110d6001600160a01b0383165b6001600160a01b031690565b6164f2565b61111e575b50016110ca565b6040516001600160a01b039190911681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756690602090a138611117565b5060005b8151811015610e1357806111786110e360019385613ac0565b6001600160a01b038116156111eb576111a161119c6001600160a01b038316611101565b616708565b6111ae575b505b0161115f565b6040516001600160a01b039190911681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d890602090a1836111a6565b506111a8565b7f35f4a7b30000000000000000000000000000000000000000000000000000000060005260046000fd5b9181601f840112156104785782359167ffffffffffffffff83116104785760208085019460e0850201011161047857565b3461047857604060031936011261047857600435611269816104ad565b60243567ffffffffffffffff8111610478577f303439e67d1363a21c3ecd1158164e797c51eced31b0351ec0e1f7afaf97779e916112ef61ffff6112b3602094369060040161121b565b9190936112be6147bc565b1692837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000600b541617600b55615095565b604051908152a1005b34610478576000600319360112610478576020611313613707565b604051908152f35b346104785760006003193601126104785760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b906040600319830112610478576004356113788161047d565b916024359067ffffffffffffffff82116104785761106d916004016104c4565b34610478576113a63661135f565b6113b19291926147bc565b67ffffffffffffffff82166113d3816000526007602052604060002054151590565b156113ee5750610e13926113e8913691610f53565b906153a8565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b34610478576040600319360112610478576004356114388161049c565b6024356114436147bc565b6001600160a01b0382169182156114cd577f78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f0238927fffffffffffffffffffffffff00000000000000000000000000000000000000006004541617600455816005556114c860405192839283602090939291936001600160a01b0360408201951681520152565b0390a1005b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461047857600060031936011261047857602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346104785760006003193601126104785760206001600160a01b03600a5416604051908152f35b3461047857600060031936011261047857600454600554604080516001600160a01b039093168352602083019190915290f35b34610478576000600319360112610478576000546001600160a01b0381163303611619577fffffffffffffffffffffffff0000000000000000000000000000000000000000600154913382841617600155166000556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b34610478576020600319360112610478577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d0917460206001600160a01b0360043561168b8161049c565b6116936147bc565b16807fffffffffffffffffffffffff0000000000000000000000000000000000000000600a541617600a55604051908152a1005b346104785760206003193601126104785760206108ba67ffffffffffffffff6004356116f28161047d565b166000526007602052604060002054151590565b346104785760c06003193601126104785761172260043561049c565b60243561172e8161047d565b6044359061173d6064356104ad565b60843567ffffffffffffffff81116104785761175d9036906004016104c4565b505060a4356002811015610478576103cc9261094b926137c1565b346104785760006003193601126104785760206001600160a01b0360015416604051908152f35b67ffffffffffffffff8111610e475760051b60200190565b9291906117c38161179f565b936117d16040519586610ed8565b602085838152019160051b810192831161047857905b8282106117f357505050565b6020809183356118028161049c565b8152019101906117e7565b9080601f8301121561047857816020610780933591016117b7565b346104785760206003193601126104785760043567ffffffffffffffff811161047857604060031982360301126104785760405161186581610e4c565b816004013567ffffffffffffffff811161047857611889906004369185010161180d565b8152602482013567ffffffffffffffff811161047857610e139260046118b2923692010161180d565b602082015261381e565b9181601f840112156104785782359167ffffffffffffffff8311610478576020808501946060850201011161047857565b346104785760606003193601126104785760043567ffffffffffffffff81116104785761191e903690600401610fee565b9060243567ffffffffffffffff81116104785761193f9036906004016118bc565b9060443567ffffffffffffffff8111610478576119609036906004016118bc565b611975611101600a546001600160a01b031690565b33141580611a54575b611a2257838614801590611a18575b6119ee5760005b86811061199d57005b806119e86119b66119b16001948b8b61396c565b61330c565b6119c183898961397c565b6119e26119da6119d286898b61397c565b92369061260a565b91369061260a565b916155d1565b01611994565b7f568efce20000000000000000000000000000000000000000000000000000000060005260046000fd5b508086141561198d565b7f8e4a23d6000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b6000fd5b50611a6a6111016001546001600160a01b031690565b33141561197e565b9160a061049a929493611ac7816101408101976001600160801b036080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b01906001600160801b036080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b346104785760206003193601126104785767ffffffffffffffff600435611b318161047d565b611b3961398c565b50611b4261398c565b501660005260086020526040600020611b71611b656002611b6a611b65856139b7565b61570c565b93016139b7565b906103cc60405192839283611a72565b346104785760006003193601126104785760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b610780916020611bde835160408452604084019061072e565b92015190602081840391015261072e565b906020610780928181520190611bc5565b346104785760206003193601126104785760043567ffffffffffffffff811161047857611c319036906004016109ab565b611c39613a08565b50611c438161578a565b60208101611c75611c70611c568361330c565b67ffffffffffffffff166000526012602052604060002090565b613a21565b611c89611c856060830151151590565b1590565b611f83576020611c9984806131fc565b905003611f3f5760208101518015611f2157606090935b013590611cc4604082015163ffffffff1690565b81516040517ff856ddb60000000000000000000000000000000000000000000000000000000081526004810185905263ffffffff909216602483015260448201959095527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316606482018190526084820195909552906020828060a48101038160006001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165af1908115610ce1576103cc95611e9394611e8e94600094611eb9575b50916080917ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff611e4f95611dfe611dd48c61330c565b604080516001600160a01b0390971687523360208801528601929092529116929081906060820190565b0390a2611e1c611e0c610f0a565b67ffffffffffffffff9095168552565b63ffffffff7f00000000000000000000000000000000000000000000000000000000000000001660208501520151151590565b15611eb05760408051825167ffffffffffffffff166020808301919091529092015163ffffffff1690820152611e888160608101610644565b9261330c565b613bef565b90611e9c610f0a565b918252602082015260405191829182611bef565b611e8890615a21565b611e4f93919450917ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff611f0e60809560203d602011611f1a575b611f068183610ed8565b810190613aab565b96939550505091611d95565b503d611efc565b506060611f39611f3185806131fc565b810190613a9c565b93611cb0565b611f4983806131fc565b90611f7f6040519283927fa3c8cf0900000000000000000000000000000000000000000000000000000000845260048401613a8b565b0390fd5b611a50611f8f8361330c565b7fd201c48a0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b602081016020825282518091526040820191602060408360051b8301019401926000915b838310611ff857505050505090565b9091929394602080612034837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08660019603018752895161072e565b97019301930191939290611fe9565b346104785760206003193601126104785767ffffffffffffffff6004356120698161047d565b1660005260086020526120826005604060002001615e02565b805190601f196120aa6120948461179f565b936120a26040519586610ed8565b80855261179f565b0160005b81811061211657505060005b815181101561210857806120ec6120e76120d660019486613ac0565b516000526009602052604060002090565b613b0e565b6120f68286613ac0565b526121018185613ac0565b50016120ba565b604051806103cc8582611fc5565b8060606020809387010152016120ae565b34610478576000600319360112610478576040516002548082526020820190600260005260206000209060005b81811061216b576103cc8561094b81870382610ed8565b8254845260209093019260019283019201612154565b346104785761218f3661135f565b61219a9291926147bc565b67ffffffffffffffff8216916121c0611c85846000526007602052604060002054151590565b61227157612203611c8560056121ea8467ffffffffffffffff166000526008602052604060002090565b016121f6368689610f53565b6020815191012090616625565b61223a57507f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76919261070060405192839283613a8b565b611f7f84926040519384937f74f23c7c00000000000000000000000000000000000000000000000000000000855260048501613bce565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b9291906122bf602091604086526040860190611bc5565b930152565b346104785760606003193601126104785760043567ffffffffffffffff8111610478576122f59036906004016109ab565b602435612301816104ad565b6044359067ffffffffffffffff82116104785761234560209161232b6123c8943690600401610f8a565b50612334613a08565b5061233f818661588c565b84615ab5565b9201356123518161047d565b604080516001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016815233602082015290810184905267ffffffffffffffff8216907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1090606090a2611e8e8161047d565b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016602082015260208152612403604082610ed8565b61240b610f0a565b91825260208201526103cc604051928392836122a8565b34610478576020600319360112610478576103cc612445600435611e8e8161047d565b60405191829160208352602083019061072e565b346104785760206003193601126104785760043567ffffffffffffffff81116104785761248a90369060040161121b565b6001600160a01b03600a5416331415806124ac575b611a2257610e1391615095565b506001600160a01b036001541633141561249f565b602060408183019282815284518094520192019060005b8181106124e55750505090565b825167ffffffffffffffff168452602093840193909201916001016124d8565b346104785760006003193601126104785761251e615db7565b805190601f196125306120948461179f565b0136602084013760005b815181101561256c578067ffffffffffffffff61255960019385613ac0565b51166125658286613ac0565b520161253a565b604051806103cc85826124c1565b8015150361047857565b359061049a8261257a565b6001600160801b0381160361047857565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c606091011261047857604051906125d782610e68565b816084356125e48161257a565b815260a4356125f28161258f565b6020820152604060c435916126068361258f565b0152565b91908260609103126104785760405161262281610e68565b604080829480356126328161257a565b845260208101356126428161258f565b60208501520135916126068361258f565b346104785760e0600319360112610478576004356126708161047d565b60607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc360112610478576040516126a681610e68565b6024356126b28161257a565b81526044356126c08161258f565b60208201526064356126d18161258f565b60408201526126df366125a0565b906001600160a01b03600a541633141580612702575b611a2257610e13926155d1565b506001600160a01b03600154163314156126f5565b346104785760206003193601126104785760043567ffffffffffffffff811161047857612748903690600401610fee565b6127506147bc565b60005b81811061275c57005b80837fece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d67ffffffffffffffff6127986119b16001968886613c11565b6128bc6127b36127a9878a88613c11565b6020810190613c51565b896127ce6127c48a838b969b613c11565b6040810190613c51565b6128016127f78c6127ef6127e582888b989b613c11565b6060810190613c51565b969095613c11565b6080810190613c51565b95909461281761281236838f6117b7565b615b54565b6128256128123685856117b7565b6128336128123687876117b7565b6128416128123689896117b7565b6128ae8c612859612850610f19565b918436916117b7565b81526128663686866117b7565b60208201526128763688886117b7565b6040820152612886368a8a6117b7565b60608201526128a98b67ffffffffffffffff16600052600e602052604060002090565b613d8f565b604051998a99169b89613e95565b0390a201612753565b346104785760406003193601126104785760043567ffffffffffffffff811161047857366023820112156104785780600401359067ffffffffffffffff8211610478573660248360081b83010111610478576024359067ffffffffffffffff821161047857610e139261293e6024933690600401610fee565b93909201613ee1565b346104785760006003193601126104785760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610478576020600319360112610478576004356129a88161047d565b6129b061398c565b506129b961398c565b50611b71611b656129e76129ec611b656129e78667ffffffffffffffff16600052600c602052604060002090565b6139b7565b9367ffffffffffffffff16600052600d602052604060002090565b346104785760206003193601126104785767ffffffffffffffff600435612a2d8161047d565b612a3561398c565b501660005260126020526103cc604060002060ff600260405192612a5884610e84565b8054845260018101546020850152015463ffffffff81166040840152818160201c161515606084015260281c16151560808201526040519182918291909160808060a0830194805184526020810151602085015263ffffffff604082015116604085015260608101511515606085015201511515910152565b346104785760006003193601126104785760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b3461047857612b1c3661101f565b919092612b276147bc565b6000915b808310612da35750505060009163ffffffff4216925b828110612b4a57005b612b5d612b58828585614422565b6144e1565b9060608201612b6c8151615c2a565b6080830193612b7b8551615c2a565b6040840190815151156114cd57612bb5611c85612bb0612ba3885167ffffffffffffffff1690565b67ffffffffffffffff1690565b616743565b612d5857612cb8612bee612bd4879a999a5167ffffffffffffffff1690565b67ffffffffffffffff166000526008602052604060002090565b612c8489612c7e8751612c6e612c0e60408301516001600160801b031690565b91612c5e612c30612c2960208401516001600160801b031690565b9251151590565b612c55612c3b610f28565b6001600160801b03851681529763ffffffff166020890152565b15156040870152565b6001600160801b03166060850152565b6001600160801b03166080830152565b82614578565b612cad89612ca48a51612c6e612c0e60408301516001600160801b031690565b60028301614578565b60048451910161466c565b602085019660005b88518051821015612cfb5790612cf5600192612cee83612ce88c5167ffffffffffffffff1690565b92613ac0565b51906153a8565b01612cc0565b50509796509490612d4f7f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c29392612d3c6001975167ffffffffffffffff1690565b9251935190519060405194859485614739565b0390a101612b41565b611a50612d6d865167ffffffffffffffff1690565b7f1d5ad3c50000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b909192612db46119b185848661396c565b94612dcb611c8567ffffffffffffffff881661659a565b612ebf57612df86005612df28867ffffffffffffffff166000526008602052604060002090565b01615e02565b9360005b8551811015612e4457600190612e3d6005612e2b8b67ffffffffffffffff166000526008602052604060002090565b01612e36838a613ac0565b5190616625565b5001612dfc565b509350937f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916612eb160019397612e96612e918267ffffffffffffffff166000526008602052604060002090565b614391565b60405167ffffffffffffffff90911681529081906020820190565b0390a1019190939293612b2b565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff861660045260246000fd5b34610478576020600319360112610478576001600160a01b03600435612f1c8161049c565b612f246147bc565b16338114612f8957807fffffffffffffffffffffffff000000000000000000000000000000000000000060005416176000556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461047857600060031936011261047857602061ffff600b5416604051908152f35b9061049a604051612fe581610ea0565b60c061305482955463ffffffff8116845263ffffffff8160201c16602085015263ffffffff8160401c16604085015263ffffffff808260601c1616606085015261303d61ffff8260801c16608086019061ffff169052565b61ffff609082901c1660a085015260a01c60ff1690565b1515910152565b61307f6130849193929367ffffffffffffffff16600052600f602052604060002090565b612fd5565b91613095611c8560c0850151151590565b61310e5761ffff1615801591906130fd57606083015163ffffffff16905b63ffffffff6130c6855163ffffffff1690565b946130d8602082015163ffffffff1690565b94156130ef5760a0015161ffff16925b1693929190565b6080015161ffff16926130e8565b604083015163ffffffff16906130b3565b5060009150819081908190565b67ffffffffffffffff90600060c060405161313581610ea0565b8281528260208201528260408201528260608201528260808201528260a0820152015216600052600f60205260406000206107806131e06040519261317984610ea0565b5463ffffffff8116845263ffffffff8160201c1660208501526131ae63ffffffff8260401c16604086019063ffffffff169052565b6131c9606082901c63ffffffff1663ffffffff166060860152565b61303d608082901c61ffff1661ffff166080860152565b151560c0830152565b604051906131f682610e2b565b60008252565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610478570180359067ffffffffffffffff82116104785760200191813603831361047857565b6020818303126104785780359067ffffffffffffffff82116104785701604081830312610478576040519161328183610e4c565b813567ffffffffffffffff8111610478578161329e918401610f8a565b8352602082013567ffffffffffffffff8111610478576132be9201610f8a565b602082015290565b9081602091031261047857516107808161257a565b90916132f26107809360408452604084019061072e565b91602081840391015261072e565b6040513d6000823e3d90fd5b356107808161047d565b356107808161049c565b6133286147bc565b60005b82811061336a5750907fc97f93e817584952f1c1d633f93784b8430f0633d002f9dcc4de4fe2780424d0916133656040519283928361361b565b0390a1565b61337d6133788285856134d5565b613503565b8051158015613499575b613420579061341a82613415611c566060600196519361340660208201516133fd6133b9604085015163ffffffff1690565b6133f56133c96080870151151590565b916133d760a0880151151590565b946133e0610f28565b9b8c5260208c015263ffffffff1660408b0152565b151588860152565b15156080870152565b015167ffffffffffffffff1690565b613576565b0161332b565b604080517f19d7585700000000000000000000000000000000000000000000000000000000815282516004820152602083015160248201529082015163ffffffff166044820152606082015167ffffffffffffffff16606482015260808201511515608482015260a090910151151560a482015260c490fd5b5067ffffffffffffffff6134b8606083015167ffffffffffffffff1690565b1615613387565b634e487b7160e01b600052603260045260246000fd5b91908110156134e55760c0020190565b6134bf565b63ffffffff81160361047857565b359061049a826134ea565b60c0813603126104785760a06040519161351c83610ebc565b80358352602081013560208401526040810135613538816134ea565b6040840152606081013561354b8161047d565b6060840152608081013561355e8161257a565b6080840152013561356e8161257a565b60a082015290565b6002608091835181556020840151600182015501916135ca63ffffffff604083015116849063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b6060810151835492909101517fffffffffffffffffffffffffffffffffffffffffffffffffffff0000ffffffff90921690151560201b64ff00000000161790151560281b65ff000000000016179055565b602080825281018390526040019160005b8181106136395750505090565b90919260c080600192863581526020870135602082015263ffffffff6040880135613663816134ea565b16604082015267ffffffffffffffff60608801356136808161047d565b16606082015260808701356136948161257a565b1515608082015260a08701356136a98161257a565b151560a082015201940192910161362c565b9067ffffffffffffffff61078092166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b90816020910312610478575190565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115610ce157600091613777575090565b610780915060203d602011613799575b6137918183610ed8565b8101906136f8565b503d613787565b67ffffffffffffffff61078091166000526007602052604060002054151590565b67ffffffffffffffff16600052600e6020526040600020916002811015613808576001146137f7578160016107809301906154fa565b8160026003610780940191016154fa565b634e487b7160e01b600052602160045260246000fd5b6138266147bc565b60208101519160005b83518110156138ab57806138486110e360019387613ac0565b61386261385d6001600160a01b038316611101565b6167f6565b61386e575b500161382f565b6040516001600160a01b039190911681527fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758090602090a138613867565b5091505160005b8151811015613968576138c86110e38284613ac0565b906001600160a01b0382161561393e577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef6139358361391a6139156111016001976001600160a01b031690565b616778565b506040516001600160a01b0390911681529081906020820190565b0390a1016138b2565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050565b91908110156134e55760051b0190565b91908110156134e5576060020190565b6040519061399982610e84565b60006080838281528260208201528260408201528260608201520152565b906040516139c481610e84565b60806001600160801b036001839560ff8154848116875263ffffffff81871c16602088015260a01c1615156040860152015481808216166060850152821c16910152565b60405190613a1582610e4c565b60606020838281520152565b90604051613a2e81610e84565b608060ff600283958054855260018101546020860152015463ffffffff81166040850152818160201c161515606085015260281c161515910152565b601f8260209493601f19938186528686013760008582860101520116010190565b916020610780938181520191613a6a565b90816020910312610478573590565b9081602091031261047857516107808161047d565b80518210156134e55760209160051b010190565b90600182811c92168015613b04575b6020831014613aee57565b634e487b7160e01b600052602260045260246000fd5b91607f1691613ae3565b9060405191826000825492613b2284613ad4565b8084529360018116908115613b8e5750600114613b47575b5061049a92500383610ed8565b90506000929192526020600020906000915b818310613b7257505090602061049a9282010138613b3a565b6020919350806001915483858901015201910190918492613b59565b6020935061049a9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138613b3a565b60409067ffffffffffffffff61078095931681528160208201520191613a6a565b67ffffffffffffffff1660005260086020526107806004604060002001613b0e565b91908110156134e55760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6181360301821215610478570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610478570180359067ffffffffffffffff821161047857602001918160051b3603831361047857565b634e487b7160e01b600052601160045260246000fd5b81810292918115918404141715613cce57565b613ca5565b91613ced918354906000199060031b92831b921b19161790565b9055565b818110613cfc575050565b60008155600101613cf1565b81519167ffffffffffffffff8311610e4757680100000000000000008311610e47576020908254848455808510613d72575b500190600052602060002060005b838110613d555750505050565b60019060206001600160a01b038551169401938184015501613d48565b613d89908460005285846000209182019101613cf1565b38613d3a565b90805180519067ffffffffffffffff8211610e4757680100000000000000008211610e47576020908454838655808410613e34575b500183600052602060002060005b838110613e115750505050906003606083613df7602061049a96015160018601613d08565b613e08604082015160028601613d08565b01519101613d08565b6001906020613e2785516001600160a01b031690565b9401938184015501613dd2565b613e4b908660005284846000209182019101613cf1565b38613dc4565b9160209082815201919060005b818110613e6b5750505090565b9091926020806001926001600160a01b038735613e878161049c565b168152019401929101613e5e565b969492613ed394613eb7613ec5936107809b999560808c5260808c0191613e51565b9189830360208b0152613e51565b918683036040880152613e51565b926060818503910152613e51565b90929192613eed6147bc565b60005b818110613f705750505060005b818110613f0957505050565b8067ffffffffffffffff613f236119b1600194868861396c565b6000613f438267ffffffffffffffff16600052600f602052604060002090565b55167f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8600080a201613efd565b613f7e6119b1828486614075565b613f89828486614075565b90602082019160a08101612710613fa9613fa283614085565b61ffff1690565b1015614039575060c001612710613fc2613fa283614085565b10156140395750907ffae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa104167ffffffffffffffff8361401f8461401a6001989767ffffffffffffffff16600052600f602052604060002090565b614099565b6140306040519283921694826142b0565b0390a201613ef0565b614045611a5091614085565b7f95f3517a0000000000000000000000000000000000000000000000000000000060005261ffff16600452602490565b91908110156134e55760081b0190565b35610780816104ad565b356107808161257a565b61426a60c061049a936140e181356140b0816134ea565b859063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b60208101356140ef816134ea565b67ffffffff0000000085549160201b16807fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff83161786557fffffffffffffffffffffffffffffffffffffffff0000000000000000ffffffff6bffffffff00000000000000006040850135614162816134ea565b60401b169216171784556141be606082013561417d816134ea565b85547fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff1660609190911b6fffffffff00000000000000000000000016178555565b6142106141cd60808301614085565b85547fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff1660809190911b71ffff0000000000000000000000000000000016178555565b61426461421f60a08301614085565b85547fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff1660909190911b73ffff00000000000000000000000000000000000016178555565b0161408f565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016179055565b61049a9092919260c06130548160e084019663ffffffff81356142d2816134ea565b16855263ffffffff60208201356142e8816134ea565b16602086015263ffffffff6040820135614301816134ea565b166040860152614323614316606083016134f8565b63ffffffff166060870152565b61433d614332608083016104b9565b61ffff166080870152565b61435761434c60a083016104b9565b61ffff1660a0870152565b01612584565b805490600081558161436d575050565b6000526020600020908101905b818110614385575050565b6000815560010161437a565b600561049a9160008155600060018201556000600282015560006003820155600481016143be8154613ad4565b90816143cd575b50500161435d565b81601f600093116001146143e55750555b38806143c5565b8183526020832061440091601f01861c810190600101613cf1565b808252602082209081548360011b906000198560031b1c1916179055556143de565b91908110156134e55760051b810135907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee181360301821215610478570190565b9080601f830112156104785781356144798161179f565b926144876040519485610ed8565b81845260208085019260051b820101918383116104785760208201905b8382106144b357505050505090565b813567ffffffffffffffff8111610478576020916144d687848094880101610f8a565b8152019101906144a4565b6101208136031261047857604051906144f982610e84565b6145028161048f565b8252602081013567ffffffffffffffff8111610478576145259036908301614462565b602083015260408101359067ffffffffffffffff82116104785761454f6145709236908301610f8a565b6040840152614561366060830161260a565b606084015260c036910161260a565b608082015290565b8151815460208401516040850151608091821b73ffffffff00000000000000000000000000000000166001600160801b039485167fffffffffffffffffffffff000000000000000000000000000000000000000000909416939093179290921791151560a01b74ff000000000000000000000000000000000000000016919091178355606084015193810151901b6fffffffffffffffffffffffffffffffff1916921691909117600190910155565b9190601f811161463657505050565b61049a926000526020600020906020601f840160051c83019310614662575b601f0160051c0190613cf1565b9091508190614655565b919091825167ffffffffffffffff8111610e47576146948161468e8454613ad4565b84614627565b6020601f82116001146146d0578190613ced9394956000926146c5575b50506000198260011b9260031b1c19161790565b0151905038806146b1565b601f198216906146e584600052602060002090565b9160005b81811061472157509583600195969710614708575b505050811b019055565b015160001960f88460031b161c191690553880806146fe565b9192602060018192868b0151815501940192016146e9565b61479461476861049a9597969467ffffffffffffffff60a095168452610100602085015261010084019061072e565b9660408301906001600160801b0360408092805115158552826020820151166020860152015116910152565b01906001600160801b0360408092805115158552826020820151166020860152015116910152565b6001600160a01b036001541633036147d057565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b60006080820161480f611c8561088983613316565b6149a757506020820191614892602061484661482d612ba38761330c565b60801b6fffffffffffffffffffffffffffffffff191690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081526fffffffffffffffffffffffffffffffff19909116600482015291829081906024820190565b03816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115610ce1578391614988575b50614960576148e46148df8461330c565b615e4d565b6148ed8361330c565b90614906611c8560a0830193610fe8610b4786866131fc565b61492057505061049a929161491b915061330c565b615ee5565b61492a92506131fc565b90611f7f6040519283927f24eb47e500000000000000000000000000000000000000000000000000000000845260048401613a8b565b6004827f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b6149a1915060203d602011610cda57610ccc8183610ed8565b386148ce565b6149b36149e791613316565b7f961c9a4f0000000000000000000000000000000000000000000000000000000083526001600160a01b0316600452602490565b90fd5b9160808301906149ff611c8561088984613316565b614b90576020840193614a1c602061484661482d612ba38961330c565b03816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115610ce157600091614b71575b50614b4757614a6a6148df8661330c565b614a738561330c565b90614a8c611c8560a0830193610fe8610b4786866131fc565b61492057505061ffff1615614b3a5767ffffffffffffffff90614b35614b10614b0a866119b1614afa614ae07f0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce615999a61330c565b67ffffffffffffffff16600052600d602052604060002090565b89614b0488613316565b91615f7c565b92613316565b94604051938493169583602090939291936001600160a01b0360408201951681520152565b0390a2565b5061491b61049a9261330c565b7f53ad11d80000000000000000000000000000000000000000000000000000000060005260046000fd5b614b8a915060203d602011610cda57610ccc8183610ed8565b38614a59565b611a50614b9c83613316565b7f961c9a4f000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b60405190614bde82610e4c565b6000825260208201600081526020820151917fffffffff00000000000000000000000000000000000000000000000000000000602c602483015160c01c92015160e01c93167ff3567d18000000000000000000000000000000000000000000000000000000008103614c51575083525290565b7fa176027f0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b90815160748110614e44575060048201517f000000000000000000000000000000000000000000000000000000000000000063ffffffff811663ffffffff831603614e0b5750506008820151916014600c82015191015192614ce7602084015163ffffffff1690565b63ffffffff811663ffffffff831603614dd25750507f000000000000000000000000000000000000000000000000000000000000000063ffffffff811663ffffffff831603614d995750505167ffffffffffffffff1667ffffffffffffffff811667ffffffffffffffff831603614d5c575050565b7ff917ffea0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff9081166004521660245260446000fd5b7f77e480260000000000000000000000000000000000000000000000000000000060005263ffffffff9081166004521660245260446000fd5b7fe366a1170000000000000000000000000000000000000000000000000000000060005263ffffffff9081166004521660245260446000fd5b7f960693cd0000000000000000000000000000000000000000000000000000000060005263ffffffff9081166004521660245260446000fd5b7f758b22cc0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b80518015614f1657602003614ed957614e9360208251830101602083016136f8565b9060ff8211614ea3575060ff1690565b611f7f906040519182917f953576f70000000000000000000000000000000000000000000000000000000083526004830161076f565b611f7f906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401818152019061072e565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff8211613cce57565b60ff16604d8111613cce57600a0a90565b8015614f6e576000190490565b634e487b7160e01b600052601260045260246000fd5b8115614f6e570490565b907f000000000000000000000000000000000000000000000000000000000000000060ff811660ff8316818114615074571161504957614fce8282614f3c565b91604d60ff8416118015615030575b614ff657505090614ff061078092614f50565b90613cbb565b7fa9cb113d0000000000000000000000000000000000000000000000000000000060005260ff908116600452166024525060445260646000fd5b5061504261503d84614f50565b614f61565b8411614fdd565b6150538183614f3c565b91604d60ff841611614ff65750509061506e61078092614f50565b90614f84565b5050505090565b91908110156134e55760e0020190565b356107808161258f565b919060005b8181106150a75750509050565b6150b281838661507b565b6150bb8161330c565b6150c7611c85826137a0565b612271578161515d6151bd92615163602060019796016150ef6150ea368361260a565b615c2a565b61515d6151108467ffffffffffffffff16600052600c602052604060002090565b9182546151306151278263ffffffff9060801c1690565b63ffffffff1690565b159081615362575b81615343575b81615331575b8161531c575b508061530d575b6152b8575b369061260a565b90616181565b6151786080840191614ae06150ea368561260a565b92835461518f6151278263ffffffff9060801c1690565b159081615293575b81615274575b81615262575b8161524d575b508061523e575b6151c3575b50369061260a565b0161509a565b6151d260a06151f7920161508b565b84906001600160801b03166fffffffffffffffffffffffffffffffff19825416179055565b82547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178355386151b5565b506152488261408f565b6151b0565b61525c915060a01c60ff161590565b386151a9565b6001600160801b0381161591506151a3565b90506001600160801b0361528b8987015460801c90565b16159061519d565b90506001600160801b036152b0898701546001600160801b031690565b161590615197565b6152c76151d26040890161508b565b82547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178355615156565b506153178161408f565b615151565b61532b915060a01c60ff161590565b3861514a565b6001600160801b038116159150615144565b90506001600160801b0361535a8c86015460801c90565b16159061513e565b90506001600160801b0361537f8c8601546001600160801b031690565b161590615138565b60409067ffffffffffffffff6107809493168152816020820152019061072e565b908051156114cd578051602082012067ffffffffffffffff8316928360005260086020526153dd8260056040600020016167ad565b156154315750816154257f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea93615420614b35946000526009602052604060002090565b61466c565b6040519182918261076f565b9050611f7f6040519283927f393b8ad200000000000000000000000000000000000000000000000000000000845260048401615387565b906040519182815491828252602082019060005260206000209260005b81811061549a57505061049a92500383610ed8565b84546001600160a01b0316835260019485019487945060209093019201615485565b91908201809211613cce57565b906154d38261179f565b6154e06040519182610ed8565b828152601f196154f0829461179f565b0190602036910137565b61550390615468565b9160055480151591826155c6575b505061551b575090565b61552490615468565b80518061553057505090565b6155416155469184959394516154bc565b6154c9565b9060005b8451811015615584578061557e6155666110e360019489613ac0565b6155708387613ac0565b906001600160a01b03169052565b0161554a565b509160005b81518110156155bf57806155b96155a56110e360019486613ac0565b6155706155b3848a516154bc565b87613ac0565b01615589565b5090925050565b101590503880615511565b67ffffffffffffffff1660008181526007602052604090205490929190156156c2577f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b9291615691826156266156bc94615c2a565b84600052600860205261563d816040600020616181565b61564683615c2a565b846000526008602052615660836002604060002001616181565b60405194855260208501906001600160801b0360408092805115158552826020820151166020860152015116910152565b60808301906001600160801b0360408092805115158552826020820151166020860152015116910152565b60e090a1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b906000198201918211613cce57565b91908203918211613cce57565b61571461398c565b506001600160801b036060820151166001600160801b0382511690602083019163ffffffff8351164203428111613cce5761575d906001600160801b0360808701511690613cbb565b8101809111613cce5761577a6001600160801b03929183926167e4565b161682524263ffffffff16905290565b600090608081016157a0611c8561088983613316565b61587f575060208101906157be602061484661482d612ba38661330c565b03816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115610ce1578491615860575b506158385761049a929160608261581e61581960406158339601613316565b6163b6565b61582a6148df8461330c565b0135925061330c565b61642b565b6004837f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b615879915060203d602011610cda57610ccc8183610ed8565b386157fa565b6149e76149b38492613316565b60808101906158a0611c8561088984613316565b614b905760208101906158bd602061484661482d612ba38661330c565b03816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115610ce157600091615a02575b50614b475780615911615819604060609401613316565b61591d6148df8461330c565b01359261ffff811690811515806159ea575b156159d957600b5461ffff169182116159a25750507f61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac2097932891614b35614b10614b0a846119b1614afa61598867ffffffffffffffff9861330c565b67ffffffffffffffff16600052600c602052604060002090565b7f7911d95b0000000000000000000000000000000000000000000000000000000060005261ffff9081166004521660245260446000fd5b505061049a9291506158339061330c565b506159fb613fa2600b5461ffff1690565b151561592f565b615a1b915060203d602011610cda57610ccc8183610ed8565b386158fa565b7fffffffff00000000000000000000000000000000000000000000000000000000602082519201517fffffffffffffffff000000000000000000000000000000000000000000000000604051937ff3567d1800000000000000000000000000000000000000000000000000000000602086015260c01b16602484015260e01b16602c82015260108152610780603082610ed8565b90615ae261307f615ac86020850161330c565b67ffffffffffffffff16600052600f602052604060002090565b90615af3611c8560c0840151151590565b615b4b5761ffff9190821615615b3e5760a0015161ffff165b168015615b3657615b30615b2860606107809401359283613cbb565b612710900490565b906156ff565b506060013590565b6080015161ffff16615b0c565b50506060013590565b80519060005b828110615b6657505050565b60018101808211613cce575b838110615b825750600101615b5a565b6001600160a01b03615b948385613ac0565b5116615ba66111016110e38487613ac0565b14615bb357600101615b72565b611a50615bc36110e38486613ac0565b7fa1726e40000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b61049a9092919260608101936001600160801b0360408092805115158552826020820151166020860152015116910152565b805115615caa5760408101516001600160801b03166001600160801b03615c6a615c5e60208501516001600160801b031690565b6001600160801b031690565b911611615c745750565b611f7f906040519182917f8020d12400000000000000000000000000000000000000000000000000000000835260048301615bf8565b6001600160801b03615cc660408301516001600160801b031690565b1615801590615d0d575b615cd75750565b611f7f906040519182917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301615bf8565b50615d25615c5e60208301516001600160801b031690565b1515615cd0565b15615d3357565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b604051906006548083528260208101600660005260206000209260005b818110615de957505061049a92500383610ed8565b8454835260019485019487945060209093019201615dd4565b906040519182815491828252602082019060005260206000209260005b818110615e3457505061049a92500383610ed8565b8454835260019485019487945060209093019201615e1f565b67ffffffffffffffff16615e6e816000526007602052604060002054151590565b15615eb8575033600052601160205260406000205415615e8a57565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fa9902c7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b67ffffffffffffffff7f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c91169182600052600860205280615f5960026040600020016001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391615f7c565b604080516001600160a01b03909216825260208201929092529081908101614b35565b8054939290919060ff60a086901c16158015616179575b61617257615fa96001600160801b038616615c5e565b9060018401958654615fe0615fda615127615fcd615c5e856001600160801b031690565b9460801c63ffffffff1690565b426156ff565b806160de575b50508381106160a0575082821061602e575061049a93945061600b91615c5e916156ff565b6001600160801b03166fffffffffffffffffffffffffffffffff19825416179055565b90616065611a50936160606160518461604b615c5e8c5460801c90565b936156ff565b61605a836156f0565b906154bc565b614f84565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024526001600160a01b0316604452606490565b7f1a76572a0000000000000000000000000000000000000000000000000000000060005260045260248390526001600160a01b031660445260646000fd5b828592939511616148576160f8615c5e6160ff9460801c90565b9185616399565b84547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178555913880615fe6565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050509050565b508115615f93565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19916162e56133659280546161c3615fda6151278363ffffffff9060801c1690565b90816162f1575b50506162b760016161e560208601516001600160801b031690565b9261623d616218615c5e6001600160801b0361620885546001600160801b031690565b166001600160801b0388166167e4565b82906001600160801b03166fffffffffffffffffffffffffffffffff19825416179055565b61629061624a8751151590565b82547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178255565b019182906001600160801b03166fffffffffffffffffffffffffffffffff19825416179055565b604083015181546001600160801b031660809190911b6fffffffffffffffffffffffffffffffff1916179055565b60405191829182615bf8565b615c5e616218916001600160801b0361634a616351958261634360018a0154928261633c616335616328876001600160801b031690565b996001600160801b031690565b9560801c90565b1690613cbb565b91166154bc565b91166167e4565b80547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff000000000000000000000000000000001617815538806161ca565b926163a49192613cbb565b8101809111613cce57610780916167e4565b7f00000000000000000000000000000000000000000000000000000000000000006163de5750565b6001600160a01b0316806000526003602052604060002054156163fe5750565b7fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b67ffffffffffffffff7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894491169182600052600860205280615f5960406000206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391615f7c565b80548210156134e55760005260206000200190600090565b805480156164dc5760001901906164cb828261649c565b60001982549160031b1b1916905555565b634e487b7160e01b600052603160045260246000fd5b60008181526003602052604090205490811561659357600019820190828211613cce57600254926000198401938411613cce5783836000956165529503616558575b50505061654160026164b4565b600390600052602052604060002090565b55600190565b6165416165849161657a61657061658a95600261649c565b90549060031b1c90565b928391600261649c565b90613cd3565b55388080616534565b5050600090565b60008181526007602052604090205490811561659357600019820190828211613cce57600654926000198401938411613cce57838360009561655295036165fa575b5050506165e960066164b4565b600790600052602052604060002090565b6165e96165849161661261657061661c95600661649c565b928391600661649c565b553880806165dc565b60018101918060005282602052604060002054928315156000146166c1576000198401848111613cce578354936000198501948511613cce576000958583616552976166799503616688575b5050506164b4565b90600052602052604060002090565b6166a86165849161669f6165706166b8958861649c565b9283918761649c565b8590600052602052604060002090565b55388080616671565b50505050600090565b80549068010000000000000000821015610e4757816166f1916001613ced9401815561649c565b81939154906000199060031b92831b921b19161790565b60008181526003602052604090205461673d576167268160026166ca565b600254906000526003602052604060002055600190565b50600090565b60008181526007602052604090205461673d576167618160066166ca565b600654906000526007602052604060002055600190565b60008181526011602052604090205461673d576167968160106166ca565b601054906000526011602052604060002055600190565b600082815260018201602052604090205461659357806167cf836001936166ca565b80549260005201602052604060002055600190565b90808210156167f1575090565b905090565b60008181526011602052604090205490811561659357600019820190828211613cce57601054926000198401938411613cce5783836165529460009603616856575b50505061684560106164b4565b601190600052602052604060002090565b6168456165849161686e61657061687895601061649c565b928391601061649c565b55388080616838565b919290156168fc5750815115616895575090565b3b1561689e5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b82519091501561690f5750805190602001fd5b611f7f906040519182917f08c379a00000000000000000000000000000000000000000000000000000000083526004830161076f56fea164736f6c634300081a000a",
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

func (_USDCTokenPool *USDCTokenPoolCaller) GetAccumulatedFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "getAccumulatedFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) GetAccumulatedFees() (*big.Int, error) {
	return _USDCTokenPool.Contract.GetAccumulatedFees(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) GetAccumulatedFees() (*big.Int, error) {
	return _USDCTokenPool.Contract.GetAccumulatedFees(&_USDCTokenPool.CallOpts)
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

func (_USDCTokenPool *USDCTokenPoolCaller) GetConfiguredMinBlockConfirmation(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "getConfiguredMinBlockConfirmation")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) GetConfiguredMinBlockConfirmation() (uint16, error) {
	return _USDCTokenPool.Contract.GetConfiguredMinBlockConfirmation(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) GetConfiguredMinBlockConfirmation() (uint16, error) {
	return _USDCTokenPool.Contract.GetConfiguredMinBlockConfirmation(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCaller) GetCurrentCustomBlockConfirmationRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (GetCurrentCustomBlockConfirmationRateLimiterState,

	error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "getCurrentCustomBlockConfirmationRateLimiterState", remoteChainSelector)

	outstruct := new(GetCurrentCustomBlockConfirmationRateLimiterState)
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutboundRateLimiterState = *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)
	outstruct.InboundRateLimiterState = *abi.ConvertType(out[1], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return *outstruct, err

}

func (_USDCTokenPool *USDCTokenPoolSession) GetCurrentCustomBlockConfirmationRateLimiterState(remoteChainSelector uint64) (GetCurrentCustomBlockConfirmationRateLimiterState,

	error) {
	return _USDCTokenPool.Contract.GetCurrentCustomBlockConfirmationRateLimiterState(&_USDCTokenPool.CallOpts, remoteChainSelector)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) GetCurrentCustomBlockConfirmationRateLimiterState(remoteChainSelector uint64) (GetCurrentCustomBlockConfirmationRateLimiterState,

	error) {
	return _USDCTokenPool.Contract.GetCurrentCustomBlockConfirmationRateLimiterState(&_USDCTokenPool.CallOpts, remoteChainSelector)
}

func (_USDCTokenPool *USDCTokenPoolCaller) GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (GetCurrentRateLimiterState,

	error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "getCurrentRateLimiterState", remoteChainSelector)

	outstruct := new(GetCurrentRateLimiterState)
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutboundRateLimiterState = *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)
	outstruct.InboundRateLimiterState = *abi.ConvertType(out[1], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return *outstruct, err

}

func (_USDCTokenPool *USDCTokenPoolSession) GetCurrentRateLimiterState(remoteChainSelector uint64) (GetCurrentRateLimiterState,

	error) {
	return _USDCTokenPool.Contract.GetCurrentRateLimiterState(&_USDCTokenPool.CallOpts, remoteChainSelector)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) GetCurrentRateLimiterState(remoteChainSelector uint64) (GetCurrentRateLimiterState,

	error) {
	return _USDCTokenPool.Contract.GetCurrentRateLimiterState(&_USDCTokenPool.CallOpts, remoteChainSelector)
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

func (_USDCTokenPool *USDCTokenPoolCaller) GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

	error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "getDynamicConfig")

	outstruct := new(GetDynamicConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Router = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ThresholdAmountForAdditionalCCVs = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_USDCTokenPool *USDCTokenPoolSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _USDCTokenPool.Contract.GetDynamicConfig(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _USDCTokenPool.Contract.GetDynamicConfig(&_USDCTokenPool.CallOpts)
}

func (_USDCTokenPool *USDCTokenPoolCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 common.Address, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "getFee", destChainSelector, arg1, arg2, arg3, blockConfirmationRequested, arg5)

	outstruct := new(GetFee)
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeeUSDCents = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DestGasOverhead = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.DestBytesOverhead = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.TokenFeeBps = *abi.ConvertType(out[3], new(uint16)).(*uint16)

	return *outstruct, err

}

func (_USDCTokenPool *USDCTokenPoolSession) GetFee(destChainSelector uint64, arg1 common.Address, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _USDCTokenPool.Contract.GetFee(&_USDCTokenPool.CallOpts, destChainSelector, arg1, arg2, arg3, blockConfirmationRequested, arg5)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) GetFee(destChainSelector uint64, arg1 common.Address, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _USDCTokenPool.Contract.GetFee(&_USDCTokenPool.CallOpts, destChainSelector, arg1, arg2, arg3, blockConfirmationRequested, arg5)
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

func (_USDCTokenPool *USDCTokenPoolCaller) GetRequiredCCVs(opts *bind.CallOpts, arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "getRequiredCCVs", arg0, remoteChainSelector, amount, arg3, arg4, direction)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) GetRequiredCCVs(arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	return _USDCTokenPool.Contract.GetRequiredCCVs(&_USDCTokenPool.CallOpts, arg0, remoteChainSelector, amount, arg3, arg4, direction)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) GetRequiredCCVs(arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	return _USDCTokenPool.Contract.GetRequiredCCVs(&_USDCTokenPool.CallOpts, arg0, remoteChainSelector, amount, arg3, arg4, direction)
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

func (_USDCTokenPool *USDCTokenPoolCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _USDCTokenPool.contract.Call(opts, &out, "getTokenTransferFeeConfig", arg0, destChainSelector, arg2, arg3, arg4)

	if err != nil {
		return *new(IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolV2TokenTransferFeeConfig)).(*IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_USDCTokenPool *USDCTokenPoolSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _USDCTokenPool.Contract.GetTokenTransferFeeConfig(&_USDCTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
}

func (_USDCTokenPool *USDCTokenPoolCallerSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _USDCTokenPool.Contract.GetTokenTransferFeeConfig(&_USDCTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
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

func (_USDCTokenPool *USDCTokenPoolTransactor) ApplyCCVConfigUpdates(opts *bind.TransactOpts, ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "applyCCVConfigUpdates", ccvConfigArgs)
}

func (_USDCTokenPool *USDCTokenPoolSession) ApplyCCVConfigUpdates(ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.ApplyCCVConfigUpdates(&_USDCTokenPool.TransactOpts, ccvConfigArgs)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) ApplyCCVConfigUpdates(ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.ApplyCCVConfigUpdates(&_USDCTokenPool.TransactOpts, ccvConfigArgs)
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

func (_USDCTokenPool *USDCTokenPoolTransactor) ApplyCustomBlockConfirmationConfigUpdates(opts *bind.TransactOpts, minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "applyCustomBlockConfirmationConfigUpdates", minBlockConfirmation, rateLimitConfigArgs)
}

func (_USDCTokenPool *USDCTokenPoolSession) ApplyCustomBlockConfirmationConfigUpdates(minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.ApplyCustomBlockConfirmationConfigUpdates(&_USDCTokenPool.TransactOpts, minBlockConfirmation, rateLimitConfigArgs)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) ApplyCustomBlockConfirmationConfigUpdates(minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.ApplyCustomBlockConfirmationConfigUpdates(&_USDCTokenPool.TransactOpts, minBlockConfirmation, rateLimitConfigArgs)
}

func (_USDCTokenPool *USDCTokenPoolTransactor) ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "applyTokenTransferFeeConfigUpdates", tokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs)
}

func (_USDCTokenPool *USDCTokenPoolSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_USDCTokenPool.TransactOpts, tokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_USDCTokenPool.TransactOpts, tokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs)
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

func (_USDCTokenPool *USDCTokenPoolTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, finality, arg2)
}

func (_USDCTokenPool *USDCTokenPoolSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.LockOrBurn0(&_USDCTokenPool.TransactOpts, lockOrBurnIn, finality, arg2)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.LockOrBurn0(&_USDCTokenPool.TransactOpts, lockOrBurnIn, finality, arg2)
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

func (_USDCTokenPool *USDCTokenPoolTransactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn, finality)
}

func (_USDCTokenPool *USDCTokenPoolSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.ReleaseOrMint0(&_USDCTokenPool.TransactOpts, releaseOrMintIn, finality)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.ReleaseOrMint0(&_USDCTokenPool.TransactOpts, releaseOrMintIn, finality)
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

func (_USDCTokenPool *USDCTokenPoolTransactor) SetCustomBlockConfirmationRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "setCustomBlockConfirmationRateLimitConfig", rateLimitConfigArgs)
}

func (_USDCTokenPool *USDCTokenPoolSession) SetCustomBlockConfirmationRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.SetCustomBlockConfirmationRateLimitConfig(&_USDCTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) SetCustomBlockConfirmationRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.SetCustomBlockConfirmationRateLimitConfig(&_USDCTokenPool.TransactOpts, rateLimitConfigArgs)
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

func (_USDCTokenPool *USDCTokenPoolTransactor) SetDynamicConfig(opts *bind.TransactOpts, router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "setDynamicConfig", router, thresholdAmountForAdditionalCCVs)
}

func (_USDCTokenPool *USDCTokenPoolSession) SetDynamicConfig(router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.SetDynamicConfig(&_USDCTokenPool.TransactOpts, router, thresholdAmountForAdditionalCCVs)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) SetDynamicConfig(router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.SetDynamicConfig(&_USDCTokenPool.TransactOpts, router, thresholdAmountForAdditionalCCVs)
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

func (_USDCTokenPool *USDCTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_USDCTokenPool *USDCTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.TransferOwnership(&_USDCTokenPool.TransactOpts, to)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.TransferOwnership(&_USDCTokenPool.TransactOpts, to)
}

func (_USDCTokenPool *USDCTokenPoolTransactor) WithdrawFees(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "withdrawFees", recipient)
}

func (_USDCTokenPool *USDCTokenPoolSession) WithdrawFees(recipient common.Address) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.WithdrawFees(&_USDCTokenPool.TransactOpts, recipient)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) WithdrawFees(recipient common.Address) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.WithdrawFees(&_USDCTokenPool.TransactOpts, recipient)
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

type USDCTokenPoolCCVConfigUpdatedIterator struct {
	Event *USDCTokenPoolCCVConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCVConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCVConfigUpdated)
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
		it.Event = new(USDCTokenPoolCCVConfigUpdated)
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

func (it *USDCTokenPoolCCVConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCVConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCVConfigUpdated struct {
	RemoteChainSelector             uint64
	OutboundCCVs                    []common.Address
	OutboundCCVsToAddAboveThreshold []common.Address
	InboundCCVs                     []common.Address
	InboundCCVsToAddAboveThreshold  []common.Address
	Raw                             types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCVConfigUpdatedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCVConfigUpdatedIterator{contract: _USDCTokenPool.contract, event: "CCVConfigUpdated", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCVConfigUpdated)
				if err := _USDCTokenPool.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
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

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseCCVConfigUpdated(log types.Log) (*USDCTokenPoolCCVConfigUpdated, error) {
	event := new(USDCTokenPoolCCVConfigUpdated)
	if err := _USDCTokenPool.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
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

type USDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator struct {
	Event *USDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
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
		it.Event = new(USDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
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

func (it *USDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator{contract: _USDCTokenPool.contract, event: "CustomBlockConfirmationInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
				if err := _USDCTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
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

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*USDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error) {
	event := new(USDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
	if err := _USDCTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator struct {
	Event *USDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
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
		it.Event = new(USDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
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

func (it *USDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator{contract: _USDCTokenPool.contract, event: "CustomBlockConfirmationOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
				if err := _USDCTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
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

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*USDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error) {
	event := new(USDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
	if err := _USDCTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCustomBlockConfirmationUpdatedIterator struct {
	Event *USDCTokenPoolCustomBlockConfirmationUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCustomBlockConfirmationUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCustomBlockConfirmationUpdated)
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
		it.Event = new(USDCTokenPoolCustomBlockConfirmationUpdated)
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

func (it *USDCTokenPoolCustomBlockConfirmationUpdatedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCustomBlockConfirmationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCustomBlockConfirmationUpdated struct {
	MinBlockConfirmation uint16
	Raw                  types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterCustomBlockConfirmationUpdated(opts *bind.FilterOpts) (*USDCTokenPoolCustomBlockConfirmationUpdatedIterator, error) {

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationUpdated")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCustomBlockConfirmationUpdatedIterator{contract: _USDCTokenPool.contract, event: "CustomBlockConfirmationUpdated", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchCustomBlockConfirmationUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCustomBlockConfirmationUpdated) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCustomBlockConfirmationUpdated)
				if err := _USDCTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationUpdated", log); err != nil {
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

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseCustomBlockConfirmationUpdated(log types.Log) (*USDCTokenPoolCustomBlockConfirmationUpdated, error) {
	event := new(USDCTokenPoolCustomBlockConfirmationUpdated)
	if err := _USDCTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationUpdated", log); err != nil {
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

type USDCTokenPoolDynamicConfigSetIterator struct {
	Event *USDCTokenPoolDynamicConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolDynamicConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolDynamicConfigSet)
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
		it.Event = new(USDCTokenPoolDynamicConfigSet)
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

func (it *USDCTokenPoolDynamicConfigSetIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolDynamicConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolDynamicConfigSet struct {
	Router                           common.Address
	ThresholdAmountForAdditionalCCVs *big.Int
	Raw                              types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterDynamicConfigSet(opts *bind.FilterOpts) (*USDCTokenPoolDynamicConfigSetIterator, error) {

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolDynamicConfigSetIterator{contract: _USDCTokenPool.contract, event: "DynamicConfigSet", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolDynamicConfigSet) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolDynamicConfigSet)
				if err := _USDCTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
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

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseDynamicConfigSet(log types.Log) (*USDCTokenPoolDynamicConfigSet, error) {
	event := new(USDCTokenPoolDynamicConfigSet)
	if err := _USDCTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
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

type USDCTokenPoolPoolFeeWithdrawnIterator struct {
	Event *USDCTokenPoolPoolFeeWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolPoolFeeWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolPoolFeeWithdrawn)
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
		it.Event = new(USDCTokenPoolPoolFeeWithdrawn)
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

func (it *USDCTokenPoolPoolFeeWithdrawnIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolPoolFeeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolPoolFeeWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterPoolFeeWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*USDCTokenPoolPoolFeeWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "PoolFeeWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolPoolFeeWithdrawnIterator{contract: _USDCTokenPool.contract, event: "PoolFeeWithdrawn", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchPoolFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolPoolFeeWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "PoolFeeWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolPoolFeeWithdrawn)
				if err := _USDCTokenPool.contract.UnpackLog(event, "PoolFeeWithdrawn", log); err != nil {
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

func (_USDCTokenPool *USDCTokenPoolFilterer) ParsePoolFeeWithdrawn(log types.Log) (*USDCTokenPoolPoolFeeWithdrawn, error) {
	event := new(USDCTokenPoolPoolFeeWithdrawn)
	if err := _USDCTokenPool.contract.UnpackLog(event, "PoolFeeWithdrawn", log); err != nil {
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

type USDCTokenPoolTokenTransferFeeConfigDeletedIterator struct {
	Event *USDCTokenPoolTokenTransferFeeConfigDeleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolTokenTransferFeeConfigDeletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolTokenTransferFeeConfigDeleted)
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
		it.Event = new(USDCTokenPoolTokenTransferFeeConfigDeleted)
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

func (it *USDCTokenPoolTokenTransferFeeConfigDeletedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolTokenTransferFeeConfigDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolTokenTransferFeeConfigDeleted struct {
	DestChainSelector uint64
	Raw               types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*USDCTokenPoolTokenTransferFeeConfigDeletedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolTokenTransferFeeConfigDeletedIterator{contract: _USDCTokenPool.contract, event: "TokenTransferFeeConfigDeleted", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolTokenTransferFeeConfigDeleted)
				if err := _USDCTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
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

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseTokenTransferFeeConfigDeleted(log types.Log) (*USDCTokenPoolTokenTransferFeeConfigDeleted, error) {
	event := new(USDCTokenPoolTokenTransferFeeConfigDeleted)
	if err := _USDCTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolTokenTransferFeeConfigUpdatedIterator struct {
	Event *USDCTokenPoolTokenTransferFeeConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolTokenTransferFeeConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolTokenTransferFeeConfigUpdated)
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
		it.Event = new(USDCTokenPoolTokenTransferFeeConfigUpdated)
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

func (it *USDCTokenPoolTokenTransferFeeConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolTokenTransferFeeConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolTokenTransferFeeConfigUpdated struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
	Raw                    types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*USDCTokenPoolTokenTransferFeeConfigUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolTokenTransferFeeConfigUpdatedIterator{contract: _USDCTokenPool.contract, event: "TokenTransferFeeConfigUpdated", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolTokenTransferFeeConfigUpdated)
				if err := _USDCTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
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

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseTokenTransferFeeConfigUpdated(log types.Log) (*USDCTokenPoolTokenTransferFeeConfigUpdated, error) {
	event := new(USDCTokenPoolTokenTransferFeeConfigUpdated)
	if err := _USDCTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetCurrentCustomBlockConfirmationRateLimiterState struct {
	OutboundRateLimiterState RateLimiterTokenBucket
	InboundRateLimiterState  RateLimiterTokenBucket
}
type GetCurrentRateLimiterState struct {
	OutboundRateLimiterState RateLimiterTokenBucket
	InboundRateLimiterState  RateLimiterTokenBucket
}
type GetDynamicConfig struct {
	Router                           common.Address
	ThresholdAmountForAdditionalCCVs *big.Int
}
type GetFee struct {
	FeeUSDCents       *big.Int
	DestGasOverhead   uint32
	DestBytesOverhead uint32
	TokenFeeBps       uint16
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

func (USDCTokenPoolCCVConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d")
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

func (USDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce615")
}

func (USDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac20979328")
}

func (USDCTokenPoolCustomBlockConfirmationUpdated) Topic() common.Hash {
	return common.HexToHash("0x303439e67d1363a21c3ecd1158164e797c51eced31b0351ec0e1f7afaf97779e")
}

func (USDCTokenPoolDomainsSet) Topic() common.Hash {
	return common.HexToHash("0xc97f93e817584952f1c1d633f93784b8430f0633d002f9dcc4de4fe2780424d0")
}

func (USDCTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f0238")
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

func (USDCTokenPoolPoolFeeWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b859599")
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

func (USDCTokenPoolTokenTransferFeeConfigDeleted) Topic() common.Hash {
	return common.HexToHash("0x5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8")
}

func (USDCTokenPoolTokenTransferFeeConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xfae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa1041")
}

func (_USDCTokenPool *USDCTokenPool) Address() common.Address {
	return _USDCTokenPool.address
}

type USDCTokenPoolInterface interface {
	GetAccumulatedFees(opts *bind.CallOpts) (*big.Int, error)

	GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetConfiguredMinBlockConfirmation(opts *bind.CallOpts) (uint16, error)

	GetCurrentCustomBlockConfirmationRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (GetCurrentCustomBlockConfirmationRateLimiterState,

		error)

	GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (GetCurrentRateLimiterState,

		error)

	GetDomain(opts *bind.CallOpts, chainSelector uint64) (USDCTokenPoolDomain, error)

	GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

		error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 common.Address, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

		error)

	GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRequiredCCVs(opts *bind.CallOpts, arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetSupportedChains(opts *bind.CallOpts) ([]uint64, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	GetTokenDecimals(opts *bind.CallOpts) (uint8, error)

	GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error)

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

	ApplyCCVConfigUpdates(opts *bind.TransactOpts, ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error)

	ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error)

	ApplyCustomBlockConfirmationConfigUpdates(opts *bind.TransactOpts, minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error)

	ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error)

	SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error)

	SetCustomBlockConfirmationRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error)

	SetDomains(opts *bind.TransactOpts, domains []USDCTokenPoolDomainUpdate) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error)

	SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFees(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error)

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

	FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCVConfigUpdatedIterator, error)

	WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCCVConfigUpdated(log types.Log) (*USDCTokenPoolCCVConfigUpdated, error)

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

	FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*USDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*USDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationUpdated(opts *bind.FilterOpts) (*USDCTokenPoolCustomBlockConfirmationUpdatedIterator, error)

	WatchCustomBlockConfirmationUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCustomBlockConfirmationUpdated) (event.Subscription, error)

	ParseCustomBlockConfirmationUpdated(log types.Log) (*USDCTokenPoolCustomBlockConfirmationUpdated, error)

	FilterDomainsSet(opts *bind.FilterOpts) (*USDCTokenPoolDomainsSetIterator, error)

	WatchDomainsSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolDomainsSet) (event.Subscription, error)

	ParseDomainsSet(log types.Log) (*USDCTokenPoolDomainsSet, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*USDCTokenPoolDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*USDCTokenPoolDynamicConfigSet, error)

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

	FilterPoolFeeWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*USDCTokenPoolPoolFeeWithdrawnIterator, error)

	WatchPoolFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolPoolFeeWithdrawn, recipient []common.Address) (event.Subscription, error)

	ParsePoolFeeWithdrawn(log types.Log) (*USDCTokenPoolPoolFeeWithdrawn, error)

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

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*USDCTokenPoolTokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*USDCTokenPoolTokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*USDCTokenPoolTokenTransferFeeConfigUpdatedIterator, error)

	WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigUpdated(log types.Log) (*USDCTokenPoolTokenTransferFeeConfigUpdated, error)

	Address() common.Address
}
