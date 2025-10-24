// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package usdc_token_pool_cctp_v2

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
	DestGasOverhead               uint32
	DestBytesOverhead             uint32
	DefaultFinalityFeeUSDCents    uint32
	CustomFinalityFeeUSDCents     uint32
	DefaultFinalityTransferFeeBps uint16
	CustomFinalityTransferFeeBps  uint16
	IsEnabled                     bool
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

type TokenPoolCustomFinalityRateLimitConfigArgs struct {
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

var USDCTokenPoolCCTPV2MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"internalType\":\"contract ITokenMessenger\"},{\"name\":\"cctpMessageTransmitterProxy\",\"type\":\"address\",\"internalType\":\"contract CCTPMessageTransmitterProxy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"FINALITY_THRESHOLD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_FEE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_USDC_MESSAGE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"struct AuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyFinalityConfigUpdates\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"destToUseDefaultFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAccumulatedFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDomain\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct USDCTokenPool.Domain\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useLegacySourcePoolDataFormat\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.CCVDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_localDomainIdentifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_messageTransmitterProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract CCTPMessageTransmitterProxy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_supportedUSDCVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_tokenMessenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract ITokenMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct RateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct RateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustomFinalityRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDomains\",\"inputs\":[{\"name\":\"domains\",\"type\":\"tuple[]\",\"internalType\":\"struct USDCTokenPool.DomainUpdate[]\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useLegacySourcePoolDataFormat\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityMinimumBlockConfirmationUpdated\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityTransferInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DomainsSet\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"struct USDCTokenPool.DomainUpdate[]\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useLegacySourcePoolDataFormat\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeeWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidBurnToken\",\"inputs\":[{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"got\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDepositHash\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"got\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidDestinationDomain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidDomain\",\"inputs\":[{\"name\":\"domain\",\"type\":\"tuple\",\"internalType\":\"struct USDCTokenPool.DomainUpdate\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useLegacySourcePoolDataFormat\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"type\":\"error\",\"name\":\"InvalidExecutionFinalityThreshold\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidFinalityConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMessageLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmation\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidMinFinalityThreshold\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidNonce\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"got\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidPreviousPool\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourceDomain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenMessengerVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidTransmitterInProxy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenTransferFeeConfigNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnknownDomain\",\"inputs\":[{\"name\":\"domain\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"UnlockingUSDCFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610180806040523461068457617d62803803809161001d8285610a0d565b8339810160c082820312610684578151906001600160a01b038216808303610684576020840151926001600160a01b038416908185036106845760408601516001600160a01b038116969094908786036106845760608101516001600160401b0381116106845781019180601f84011215610684578251926001600160401b038411610689578360051b9060208201946100ba6040519687610a0d565b855260208086019282010192831161068457602001905b8282106109f5575050506100f360a06100ec60808401610a30565b9201610a30565b9060209660405199610105898c610a0d565b60008b52600036813733156109e457600180546001600160a01b03191633179055801580156109d3575b80156109c2575b6109b157600492899260805260c0526040519283809263313ce56760e01b82525afa809160009161097a575b5090610956575b50600660a052600480546001600160a01b0319166001600160a01b03929092169190911790558051151560e0819052610830575b50604051926101ac8585610a0d565b60008452600036813760408051979088016001600160401b03811189821017610689576040528752838588015260005b8451811015610243576001906001600160a01b036101fa8288610a60565b51168761020682610ae1565b610213575b5050016101dc565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a1388761020b565b508493508587519260005b84518110156102bf576001600160a01b036102698287610a60565b51169081156102ae577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef88836102a0600195610bfd565b50604051908152a10161024e565b6342bcdf7f60e11b60005260046000fd5b5090859185600161010052841561081f57604051632c12192160e01b81528481600481895afa9081156106ef576000916107ea575b5060405163054fd4d560e41b81526001600160a01b039190911691908581600481865afa9081156106ef576000916107cd575b5063ffffffff8061010051169116908082036107b6575050604051639cdbb18160e01b815285816004818a5afa9081156106ef57600091610799575b5063ffffffff80610100511691169080820361078257505084600491604051928380926367e0ed8360e11b82525afa80156106ef578291600091610739575b506001600160a01b03160361072857600492849261012052610140526040519283809263234d8e3d60e21b82525afa9081156106ef576000916106fb575b506101605260805161012051604051636eb1769f60e11b81523060048201526001600160a01b03918216602482018190529492909116908381604481855afa9081156106ef576000916106c2575b5060001981018091116106ac57604051908482019563095ea7b360e01b87526024830152604482015260448152610466606482610a0d565b6000806040968751936104798986610a0d565b8785527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656488860152519082865af13d1561069f573d906001600160401b0382116106895786516104e69490926104d8601f8201601f1916890185610a0d565b83523d60008885013e610cca565b80518061060b575b847f2e902d38f15b233cbb63711add0fca4545334d3a169d60c0a616494d7eea954485858351908152a151616fc79081610d9b823960805181818161059e015281816107a40152818161080201528181610d1601528181611ec201528181612591015281816138110152818161649b0152616ab3015260a051818181610902015281816126010152818161561901526156a9015260c051818181612c0201528181614b3501528181614cd901528181615b740152615c80015260e05181818161127d01528181612d0801526169dd0152610100518181816107530152614f900152610120518181816114a00152611e73015261014051818181610c7e0152611d4d01526101605181818161169301528181611fbe01526150160152f35b81849181010312610684578201518015908115036106845761062e5783806104ee565b50608491519062461bcd60e51b82526004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b916104e692606091610cca565b634e487b7160e01b600052601160045260246000fd5b90508381813d83116106e8575b6106d98183610a0d565b8101031261068457518561042e565b503d6106cf565b6040513d6000823e3d90fd5b61071b9150823d8411610721575b6107138183610a0d565b810190610a44565b836103e0565b503d610709565b632a32133b60e11b60005260046000fd5b9091508581813d831161077b575b6107518183610a0d565b810103126107775751906001600160a01b038216820361077457508190876103a2565b80fd5b5080fd5b503d610747565b633785f8f160e01b60005260045260245260446000fd5b6107b09150863d8811610721576107138183610a0d565b87610363565b63960693cd60e01b60005260045260245260446000fd5b6107e49150863d8811610721576107138183610a0d565b87610327565b90508481813d8311610818575b6108018183610a0d565b810103126106845761081290610a30565b866102f4565b503d6107f7565b6306b7c75960e31b60005260046000fd5b909194604051936108418686610a0d565b60008552600036813760e051156109455760005b85518110156108bc576001906001600160a01b036108738289610a60565b51168861087f82610c36565b61088c575b505001610855565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a13888610884565b50919350919460005b8451811015610939576001906001600160a01b036108e38288610a60565b5116801561093357876108f582610bbe565b610903575b50505b016108c5565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a138876108fa565b506108fd565b5091949092503861019d565b6335f4a7b360e01b60005260046000fd5b60ff1660068114610169576332ad3e0760e11b600052600660045260245260446000fd5b8881813d83116109aa575b61098f8183610a0d565b8101031261077757519060ff82168203610774575038610162565b503d610985565b630a64406560e11b60005260046000fd5b506001600160a01b03831615610136565b506001600160a01b0384161561012f565b639b15e16f60e01b60005260046000fd5b60208091610a0284610a30565b8152019101906100d1565b601f909101601f19168101906001600160401b0382119082101761068957604052565b51906001600160a01b038216820361068457565b90816020910312610684575163ffffffff811681036106845790565b8051821015610a745760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b8054821015610a745760005260206000200190600090565b80548015610acb576000190190610ab98282610a8a565b8154906000199060031b1b1916905555565b634e487b7160e01b600052603160045260246000fd5b6000818152601160205260409020548015610b8c5760001981018181116106ac576010546000198101919082116106ac57808203610b3b575b505050610b276010610aa2565b600052601160205260006040812055600190565b610b74610b4c610b5d936010610a8a565b90549060031b1c9283926010610a8a565b819391549060031b91821b91600019901b19161790565b90556000526011602052604060002055388080610b1a565b5050600090565b805490680100000000000000008210156106895781610b5d916001610bba94018155610a8a565b9055565b80600052600360205260406000205415600014610bf757610be0816002610b93565b600254906000526003602052604060002055600190565b50600090565b80600052601160205260406000205415600014610bf757610c1f816010610b93565b601054906000526011602052604060002055600190565b6000818152600360205260409020548015610b8c5760001981018181116106ac576002546000198101919082116106ac57818103610c90575b505050610c7c6002610aa2565b600052600360205260006040812055600190565b610cb2610ca1610b5d936002610a8a565b90549060031b1c9283926002610a8a565b90556000526003602052604060002055388080610c6f565b91929015610d2c5750815115610cde575090565b3b15610ce75790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b825190915015610d3f5750805190602001fd5b6040519062461bcd60e51b8252602060048301528181519182602483015260005b838110610d825750508160006044809484010152601f80199101168101030190fd5b60208282018101516044878401015285935001610d6056fe6080604052600436101561001257600080fd5b60003560e01c806301ffc9a714610357578063164e68de14610352578063181f5a771461034d578063212a052e1461034857806321df0da714610343578063240028e81461033e5780632451a6271461033957806324f65ee71461033457806337b192471461032f578063390775371461032a5780633ce85d4114610325578063489a68f2146103205780634ac8bd5f1461031b5780634c5ef0ed1461031657806354c8a4f3146103115780635df45a371461030c5780636155cda01461030757806362ddd3c414610302578063698c2c66146102fd5780636b716b0d146102f85780636d3d1a58146102f35780637437ff9f146102ee57806379ba5097146102e95780637d54534e146102e4578063804ba5a9146102df5780638926f54f146102da57806389720a62146102d55780638a86fe14146102d05780638da5cb5b146102cb57806391a2749a146102c6578063962d4020146102c157806398db9643146102bc5780639a4575b9146102b7578063a42a7b8b146102b2578063a7cd63b7146102ad578063acfecf91146102a8578063af58d59f146102a3578063b1c71c651461029e578063b794658014610299578063bc063e1a14610294578063c4bffe2b1461028f578063c75eea9c1461028a578063c8c8fd1914610285578063cf7401f314610280578063d966866b1461027b578063da4b05e714610276578063dc04fa1f14610271578063dc0bd9711461026c578063dfadfa3514610267578063e0351e1314610262578063e8a1da171461025d5763f2fde38b1461025857600080fd5b61314c565b612d2d565b612cf0565b612c26565b612bd5565b612b53565b612b36565b612988565b6128aa565b6127ab565b612760565b6126eb565b61268b565b612654565b6124e9565b612421565b6122fa565b6122a0565b6121bc565b611dac565b611d20565b611baf565b611aea565b611a2d565b6119e8565b611976565b611937565b611887565b6117f6565b61172b565b6116eb565b6116b7565b611676565b611580565b6114fd565b611473565b611450565b61124b565b611183565b610f7b565b610eab565b610dff565b610b99565b610a83565b6108e8565b610882565b6107c8565b610777565b610736565b6106d7565b6104bb565b34610498576020600319360112610498576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361049857807ff208a58f000000000000000000000000000000000000000000000000000000006103ec921490811561046e575b8115610444575b811561041a575b81156103f0575b5060405190151581529081906020820190565b0390f35b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014386103d9565b7f0e64dd2900000000000000000000000000000000000000000000000000000000811491506103d2565b7f7a8c772600000000000000000000000000000000000000000000000000000000811491506103cb565b7faff2afbf00000000000000000000000000000000000000000000000000000000811491506103c4565b600080fd5b73ffffffffffffffffffffffffffffffffffffffff81160361049857565b34610498576020600319360112610498576004356104d88161049d565b6104e0614a1b565b6104e86137c8565b90816104f057005b73ffffffffffffffffffffffffffffffffffffffff7f738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b8595999161064960405161062e60208201917fa9059cbb00000000000000000000000000000000000000000000000000000000835261059b8161058d8a88602484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b03601f1981018352826110b6565b857f0000000000000000000000000000000000000000000000000000000000000000166000806040958651946105d188876110b6565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020870152519082855af13d1561066f573d9161061283611115565b9261061f875194856110b6565b83523d6000602085013e616ef6565b80518061064e575b5050519283921694829190602083019252565b0390a2005b8160208061066393610668950101910161336e565b616289565b3880610636565b606091616ef6565b600091031261049857565b919082519283825260005b8481106106ae575050601f19601f8460006020809697860101520116010190565b8060208092840101518282860101520161068d565b9060206106d4928181520190610682565b90565b34610498576000600319360112610498576103ec60408051906106fa81836110b6565b601d82527f55534443546f6b656e506f6f6c43435450563220312e372e302d646576000000602083015251918291602083526020830190610682565b3461049857600060031936011261049857602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461049857600060031936011261049857602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346104985760206003193601126104985760206108286004356107ea8161049d565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691161490565b6040519015158152f35b602060408183019282815284518094520192019060005b8181106108565750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101610849565b34610498576000600319360112610498576040516010548082526020820190601060005260206000209060005b8181106108d2576103ec856108c6818703826110b6565b60405191829182610832565b82548452602090930192600192830192016108af565b3461049857600060031936011261049857602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b67ffffffffffffffff81160361049857565b359061094382610926565b565b908160a09103126104985790565b61ffff81160361049857565b359061094382610953565b9181601f840112156104985782359167ffffffffffffffff8311610498576020838186019501011161049857565b60a0600319820112610498576004356109b08161049d565b916024356109bd81610926565b9160443567ffffffffffffffff811161049857816109dd91600401610945565b916064356109ea81610953565b916084359067ffffffffffffffff821161049857610a0a9160040161096a565b9091565b6109439092919260c08060e083019563ffffffff815116845263ffffffff602082015116602085015263ffffffff604082015116604085015263ffffffff606082015116606085015261ffff6080820151166080850152610a7a60a082015160a086019061ffff169052565b01511515910152565b346104985767ffffffffffffffff610a9a36610998565b505050509050600060c0604051610ab081611009565b8281528260208201528260408201528260608201528260808201528260a0820152015216600052600f6020526103ec6040600020610b7e610b7560405192610af784611009565b5463ffffffff8116845263ffffffff8160201c166020850152610b2c63ffffffff8260401c16604086019063ffffffff169052565b610b47606082901c63ffffffff1663ffffffff166060860152565b610b5e608082901c61ffff1661ffff166080860152565b61ffff609082901c1660a085015260a01c60ff1690565b151560c0830152565b60405191829182610a0e565b90816101009103126104985790565b346104985760206003193601126104985760043567ffffffffffffffff811161049857610bca903690600401610b8a565b610bd2613291565b506060810135610be28183614a66565b610c636020610bff610bf760e08601866132a4565b8101906132f5565b610c28610c21610c1c610c1560c08901896132a4565b3691611131565b614e9c565b8251614f7d565b8181519101519060405193849283927f57ecfd2800000000000000000000000000000000000000000000000000000000845260048401613383565b0381600073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af1908115610dc957600091610d9a575b5015610d7057817ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc067ffffffffffffffff610cfb6040610cf460206103ec98016133b4565b94016133be565b6040805173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000081168252336020830152929092169082015260608101859052921691608090a2610d5d6110d9565b8190526040519081529081906020820190565b7fbf969f220000000000000000000000000000000000000000000000000000000060005260046000fd5b610dbc915060203d602011610dc2575b610db481836110b6565b81019061336e565b38610caf565b503d610daa565b6133a8565b9181601f840112156104985782359167ffffffffffffffff83116104985760208085019460e0850201011161049857565b3461049857604060031936011261049857600435610e1c81610953565b60243567ffffffffffffffff8111610498577f401d6d5b7a131da25eadbfe5bc7b2c80d7a0094d8d026d2d2c452f81e9b2af7e91610ea261ffff610e666020943690600401610dce565b919093610e71614a1b565b1692837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000600b541617600b5561521e565b604051908152a1005b346104985760406003193601126104985760043567ffffffffffffffff811161049857610edf6103ec913690600401610b8a565b60243590610eec82610953565b6000604051610efa8161102a565b52610f2c610f246060830135610f1e610f19610c1560c08701876132a4565b615570565b906156a6565b928383614c8e565b7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc067ffffffffffffffff610cfb60206040850194610f6a863561049d565b013593610f7685610926565b6133be565b346104985760206003193601126104985760043567ffffffffffffffff8111610498573660238201121561049857806004013567ffffffffffffffff81116104985736602460c0830284010111610498576024610fd892016133c8565b005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60e0810190811067ffffffffffffffff82111761102557604052565b610fda565b6020810190811067ffffffffffffffff82111761102557604052565b6040810190811067ffffffffffffffff82111761102557604052565b6060810190811067ffffffffffffffff82111761102557604052565b60a0810190811067ffffffffffffffff82111761102557604052565b60c0810190811067ffffffffffffffff82111761102557604052565b90601f601f19910116810190811067ffffffffffffffff82111761102557604052565b604051906109436020836110b6565b604051906109436040836110b6565b604051906109436080836110b6565b6040519061094360a0836110b6565b67ffffffffffffffff811161102557601f01601f191660200190565b92919261113d82611115565b9161114b60405193846110b6565b829481845281830111610498578281602093846000960137010152565b9080601f83011215610498578160206106d493359101611131565b34610498576040600319360112610498576004356111a081610926565b60243567ffffffffffffffff8111610498576020916111c6610828923690600401611168565b9061377c565b9181601f840112156104985782359167ffffffffffffffff8311610498576020808501948460051b01011161049857565b60406003198201126104985760043567ffffffffffffffff81116104985781611228916004016111cc565b929092916024359067ffffffffffffffff821161049857610a0a916004016111cc565b346104985761127361127b61125f366111fd565b949161126c939193614a1b565b3691611a79565b923691611a79565b7f0000000000000000000000000000000000000000000000000000000000000000156114265760005b825181101561136957806112d76112bd60019386613c17565b5173ffffffffffffffffffffffffffffffffffffffff1690565b61131361130e73ffffffffffffffffffffffffffffffffffffffff83165b73ffffffffffffffffffffffffffffffffffffffff1690565b616b4a565b61131f575b50016112a4565b60405173ffffffffffffffffffffffffffffffffffffffff9190911681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756690602090a138611318565b5060005b8151811015610fd857806113866112bd60019385613c17565b73ffffffffffffffffffffffffffffffffffffffff811615611420576113c96113c473ffffffffffffffffffffffffffffffffffffffff83166112f5565b616d7d565b6113d6575b505b0161136d565b60405173ffffffffffffffffffffffffffffffffffffffff9190911681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d890602090a1836113ce565b506113d0565b7f35f4a7b30000000000000000000000000000000000000000000000000000000060005260046000fd5b3461049857600060031936011261049857602061146b6137c8565b604051908152f35b3461049857600060031936011261049857602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b906040600319830112610498576004356114dd81610926565b916024359067ffffffffffffffff821161049857610a0a9160040161096a565b346104985761150b366114c4565b611516929192614a1b565b67ffffffffffffffff8216611538816000526007602052604060002054151590565b156115535750610fd89261154d913691611131565b906157b4565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346104985760406003193601126104985760043561159d8161049d565b6024356115a8614a1b565b73ffffffffffffffffffffffffffffffffffffffff821691821561164c577f78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f0238927fffffffffffffffffffffffff0000000000000000000000000000000000000000600454161760045581600555611647604051928392836020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b0390a1005b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461049857600060031936011261049857602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461049857600060031936011261049857602073ffffffffffffffffffffffffffffffffffffffff600a5416604051908152f35b34610498576000600319360112610498576004546005546040805173ffffffffffffffffffffffffffffffffffffffff9093168352602083019190915290f35b346104985760006003193601126104985760005473ffffffffffffffffffffffffffffffffffffffff811633036117cc577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b34610498576020600319360112610498577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174602073ffffffffffffffffffffffffffffffffffffffff60043561184b8161049d565b611853614a1b565b16807fffffffffffffffffffffffff0000000000000000000000000000000000000000600a541617600a55604051908152a1005b346104985760206003193601126104985760043567ffffffffffffffff8111610498576118b8903690600401610dce565b73ffffffffffffffffffffffffffffffffffffffff600a541633141580611915575b6118e757610fd89161521e565b7f8e4a23d6000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b5073ffffffffffffffffffffffffffffffffffffffff600154163314156118da565b3461049857602060031936011261049857602061082867ffffffffffffffff60043561196281610926565b166000526007602052604060002054151590565b346104985760c06003193601126104985761199260043561049d565b60243561199e81610926565b604435906119ad606435610953565b60843567ffffffffffffffff8111610498576119cd90369060040161096a565b505060a4356002811015610498576103ec926108c69261388f565b3461049857608063ffffffff61ffff81611a0e611a0436610998565b5050925050613905565b9392959091604051968752166020860152166040840152166060820152f35b3461049857600060031936011261049857602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b67ffffffffffffffff81116110255760051b60200190565b929190611a8581611a61565b93611a9360405195866110b6565b602085838152019160051b810192831161049857905b828210611ab557505050565b602080918335611ac48161049d565b815201910190611aa9565b9080601f83011215610498578160206106d493359101611a79565b346104985760206003193601126104985760043567ffffffffffffffff8111610498576040600319823603011261049857604051611b2781611046565b816004013567ffffffffffffffff811161049857611b4b9060043691850101611acf565b8152602482013567ffffffffffffffff811161049857610fd8926004611b749236920101611acf565b60208201526139c5565b9181601f840112156104985782359167ffffffffffffffff8311610498576020808501946060850201011161049857565b346104985760606003193601126104985760043567ffffffffffffffff811161049857611be09036906004016111cc565b9060243567ffffffffffffffff811161049857611c01903690600401611b7e565b9060443567ffffffffffffffff811161049857611c22903690600401611b7e565b611c446112f5600a5473ffffffffffffffffffffffffffffffffffffffff1690565b33141580611cf5575b6118e757838614801590611ce7575b611cbd5760005b868110611c6c57005b80611cb7611c85611c806001948b8b613b54565b6133b4565b611c90838989613b64565b611cb1611ca9611ca186898b613b64565b923690612861565b913690612861565b916159f7565b01611c63565b7f568efce20000000000000000000000000000000000000000000000000000000060005260046000fd5b5080861415611c5c565b6000fd5b50611d186112f560015473ffffffffffffffffffffffffffffffffffffffff1690565b331415611c4d565b3461049857600060031936011261049857602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b6106d4916020611d8a8351604084526040840190610682565b920151906020818403910152610682565b9060206106d4928181520190611d71565b346104985760206003193601126104985760043567ffffffffffffffff811161049857611ddd903690600401610945565b611de5613b74565b50611def81615b27565b60208101611e21611e1c611e02836133b4565b67ffffffffffffffff166000526012602052604060002090565b613b8d565b91611e36611e326060850151151590565b1590565b6120fc576020611e4682806132a4565b9050036120b8576020830151801561209c57925b606073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016920135926040820192611eaa845163ffffffff1690565b9273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016938151833b15610498576040517fd04857b00000000000000000000000000000000000000000000000000000000081526004810189905263ffffffff9290921660248301526044820189905273ffffffffffffffffffffffffffffffffffffffff861660648301526084820152600060a482018190526107d060c4830152909492859060e490829084905af18015610dc95761201361206496611ff27ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1094866103ec9c61205f9a67ffffffffffffffff97612081575b50611fe87f0000000000000000000000000000000000000000000000000000000000000000955163ffffffff1690565b9251928d86615dd8565b612009611ffd6110e8565b63ffffffff9093168352565b6020820152615e83565b96612057612020866133b4565b6040805173ffffffffffffffffffffffffffffffffffffffff90971687523360208801528601929092529116929081906060820190565b0390a26133b4565b613de4565b9061206d6110e8565b918252602082015260405191829182611d9b565b806120906000612096936110b6565b80610677565b38611fb8565b506120b26120aa82806132a4565b810190613c08565b92611e5a565b806120c2916132a4565b906120f86040519283927fa3c8cf0900000000000000000000000000000000000000000000000000000000845260048401613bf7565b0390fd5b611cf1612108836133b4565b7fd201c48a0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b602081016020825282518091526040820191602060408360051b8301019401926000915b83831061217157505050505090565b90919293946020806121ad837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc086600196030187528951610682565b97019301930191939290612162565b346104985760206003193601126104985767ffffffffffffffff6004356121e281610926565b1660005260086020526121fb600560406000200161635f565b805190601f1961222361220d84611a61565b9361221b60405195866110b6565b808552611a61565b0160005b81811061228f57505060005b8151811015612281578061226561226061224f60019486613c17565b516000526009602052604060002090565b613c7e565b61226f8286613c17565b5261227a8185613c17565b5001612233565b604051806103ec858261213e565b806060602080938701015201612227565b34610498576000600319360112610498576040516002548082526020820190600260005260206000209060005b8181106122e4576103ec856108c6818703826110b6565b82548452602090930192600192830192016122cd565b3461049857612308366114c4565b612313929192614a1b565b67ffffffffffffffff821691612339611e32846000526007602052604060002054151590565b6123ea5761237c611e3260056123638467ffffffffffffffff166000526008602052604060002090565b0161236f368689611131565b6020815191012090616c7d565b6123b357507f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76919261064960405192839283613bf7565b6120f884926040519384937f74f23c7c00000000000000000000000000000000000000000000000000000000855260048501613d3e565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b346104985760206003193601126104985767ffffffffffffffff60043561244781610926565b61244f613d5f565b501660005260086020526103ec61247461246f6002604060002001613d8a565b615f0e565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b9291906124e4602091604086526040860190611d71565b930152565b346104985760606003193601126104985760043567ffffffffffffffff81116104985761251a903690600401610945565b60243561252681610953565b6044359067ffffffffffffffff82116104985761256a6020916125506125fa943690600401611168565b50612559613b74565b506125648186615c36565b84615fb0565b92013561257681610926565b6040805173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815233602082015290810184905267ffffffffffffffff8216907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1090606090a261205f81610926565b60405160ff7f0000000000000000000000000000000000000000000000000000000000000000166020820152602081526126356040826110b6565b61263d6110e8565b91825260208201526103ec604051928392836124cd565b34610498576020600319360112610498576103ec61267760043561205f81610926565b604051918291602083526020830190610682565b3461049857600060031936011261049857602060405160008152f35b602060408183019282815284518094520192019060005b8181106126cb5750505090565b825167ffffffffffffffff168452602093840193909201916001016126be565b3461049857600060031936011261049857612704616314565b805190601f1961271661220d84611a61565b0136602084013760005b8151811015612752578067ffffffffffffffff61273f60019385613c17565b511661274b8286613c17565b5201612720565b604051806103ec85826126a7565b346104985760206003193601126104985767ffffffffffffffff60043561278681610926565b61278e613d5f565b501660005260086020526103ec61247461246f6040600020613d8a565b346104985760006003193601126104985760206040516101188152f35b8015150361049857565b3590610943826127c8565b6fffffffffffffffffffffffffffffffff81160361049857565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c6060910112610498576040519061282e82611062565b8160843561283b816127c8565b815260a435612849816127dd565b6020820152604060c4359161285d836127dd565b0152565b91908260609103126104985760405161287981611062565b60408082948035612889816127c8565b84526020810135612899816127dd565b602085015201359161285d836127dd565b346104985760e0600319360112610498576004356128c781610926565b60607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc360112610498576040516128fd81611062565b602435612909816127c8565b8152604435612917816127dd565b6020820152606435612928816127dd565b6040820152612936366127f7565b9073ffffffffffffffffffffffffffffffffffffffff600a541633141580612966575b6118e757610fd8926159f7565b5073ffffffffffffffffffffffffffffffffffffffff60015416331415612959565b346104985760206003193601126104985760043567ffffffffffffffff8111610498576129b99036906004016111cc565b6129c1614a1b565b60005b8181106129cd57005b80837fece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d67ffffffffffffffff612a09611c806001968886613e06565b612b2d612a24612a1a878a88613e06565b6020810190613e46565b89612a3f612a358a838b969b613e06565b6040810190613e46565b612a72612a688c612a60612a5682888b989b613e06565b6060810190613e46565b969095613e06565b6080810190613e46565b959094612a88612a8336838f611a79565b61604f565b612a96612a83368585611a79565b612aa4612a83368787611a79565b612ab2612a83368989611a79565b612b1f8c612aca612ac16110f7565b91843691611a79565b8152612ad7368686611a79565b6020820152612ae7368888611a79565b6040820152612af7368a8a611a79565b6060820152612b1a8b67ffffffffffffffff16600052600e602052604060002090565b613faa565b604051998a99169b896140ca565b0390a2016129c4565b346104985760006003193601126104985760206040516107d08152f35b346104985760406003193601126104985760043567ffffffffffffffff811161049857366023820112156104985780600401359067ffffffffffffffff8211610498573660248360081b83010111610498576024359067ffffffffffffffff821161049857610fd892612bcc60249336906004016111cc565b93909201614116565b3461049857600060031936011261049857602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346104985760206003193601126104985767ffffffffffffffff600435612c4c81610926565b612c54613d5f565b501660005260126020526103ec604060002060ff600260405192612c778461107e565b8054845260018101546020850152015463ffffffff81166040840152818160201c161515606084015260281c16151560808201526040519182918291909160808060a0830194805184526020810151602085015263ffffffff604082015116604085015260608101511515606085015201511515910152565b346104985760006003193601126104985760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b3461049857612d3b366111fd565b919092612d46614a1b565b6000915b808310612ff85750505060009163ffffffff4216925b828110612d6957005b612d7c612d77828585614657565b614716565b9060608201612d8b8151616148565b6080830193612d9a8551616148565b60408401908151511561164c57612dd4611e32612dcf612dc2885167ffffffffffffffff1690565b67ffffffffffffffff1690565b616db8565b612fad57612f0d612e0d612df3879a999a5167ffffffffffffffff1690565b67ffffffffffffffff166000526008602052604060002090565b612ed089612eca8751612eb1612e3660408301516fffffffffffffffffffffffffffffffff1690565b91612e98612e61612e5a60208401516fffffffffffffffffffffffffffffffff1690565b9251151590565b612e8f612e6c611106565b6fffffffffffffffffffffffffffffffff851681529763ffffffff166020890152565b15156040870152565b6fffffffffffffffffffffffffffffffff166060850152565b6fffffffffffffffffffffffffffffffff166080830152565b826147ad565b612f0289612ef98a51612eb1612e3660408301516fffffffffffffffffffffffffffffffff1690565b600283016147ad565b6004845191016148b9565b602085019660005b88518051821015612f505790612f4a600192612f4383612f3d8c5167ffffffffffffffff1690565b92613c17565b51906157b4565b01612f15565b50509796509490612fa47f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c29392612f916001975167ffffffffffffffff1690565b9251935190519060405194859485614986565b0390a101612d60565b611cf1612fc2865167ffffffffffffffff1690565b7f1d5ad3c50000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b909192613009611c80858486613b54565b94613020611e3267ffffffffffffffff8816616bf2565b6131145761304d60056130478867ffffffffffffffff166000526008602052604060002090565b0161635f565b9360005b85518110156130995760019061309260056130808b67ffffffffffffffff166000526008602052604060002090565b0161308b838a613c17565b5190616c7d565b5001613051565b509350937f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916613106600193976130eb6130e68267ffffffffffffffff166000526008602052604060002090565b6145c6565b60405167ffffffffffffffff90911681529081906020820190565b0390a1019190939293612d4a565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff861660045260246000fd5b346104985760206003193601126104985773ffffffffffffffffffffffffffffffffffffffff60043561317e8161049d565b613186614a1b565b163381146131f857807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b9061094360405161323281611009565b60c061328a82955463ffffffff8116845263ffffffff8160201c16602085015263ffffffff8160401c16604085015263ffffffff808260601c16166060850152610b5e61ffff8260801c16608086019061ffff169052565b1515910152565b6040519061329e8261102a565b60008252565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610498570180359067ffffffffffffffff82116104985760200191813603831361049857565b6020818303126104985780359067ffffffffffffffff82116104985701604081830312610498576040519161332983611046565b813567ffffffffffffffff81116104985781613346918401611168565b8352602082013567ffffffffffffffff8111610498576133669201611168565b602082015290565b9081602091031261049857516106d4816127c8565b909161339a6106d493604084526040840190610682565b916020818403910152610682565b6040513d6000823e3d90fd5b356106d481610926565b356106d48161049d565b6133d0614a1b565b60005b8281106134125750907fc97f93e817584952f1c1d633f93784b8430f0633d002f9dcc4de4fe2780424d09161340d604051928392836136dc565b0390a1565b613425613420828585613596565b6135c4565b8051158015613541575b6134c857906134c2826134bd611e02606060019651936134ae60208201516134a5613461604085015163ffffffff1690565b61349d6134716080870151151590565b9161347f60a0880151151590565b94613488611106565b9b8c5260208c015263ffffffff1660408b0152565b151588860152565b15156080870152565b015167ffffffffffffffff1690565b613637565b016133d3565b604080517f19d7585700000000000000000000000000000000000000000000000000000000815282516004820152602083015160248201529082015163ffffffff166044820152606082015167ffffffffffffffff16606482015260808201511515608482015260a090910151151560a482015260c490fd5b5067ffffffffffffffff613560606083015167ffffffffffffffff1690565b161561342f565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b91908110156135a65760c0020190565b613567565b63ffffffff81160361049857565b3590610943826135ab565b60c0813603126104985760a0604051916135dd8361109a565b803583526020810135602084015260408101356135f9816135ab565b6040840152606081013561360c81610926565b6060840152608081013561361f816127c8565b6080840152013561362f816127c8565b60a082015290565b60026080918351815560208401516001820155019161368b63ffffffff604083015116849063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b6060810151835492909101517fffffffffffffffffffffffffffffffffffffffffffffffffffff0000ffffffff90921690151560201b64ff00000000161790151560281b65ff000000000016179055565b602080825281018390526040019160005b8181106136fa5750505090565b90919260c080600192863581526020870135602082015263ffffffff6040880135613724816135ab565b16604082015267ffffffffffffffff606088013561374181610926565b1660608201526080870135613755816127c8565b1515608082015260a087013561376a816127c8565b151560a08201520194019291016136ed565b9067ffffffffffffffff6106d492166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b90816020910312610498575190565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610dc957600091613845575090565b6106d4915060203d602011613867575b61385f81836110b6565b8101906137b9565b503d613855565b67ffffffffffffffff6106d491166000526007602052604060002054151590565b67ffffffffffffffff16600052600e60205260406000209160028110156138d6576001146138c5578160016106d4930190615913565b81600260036106d494019101615913565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b61392961392e9193929367ffffffffffffffff16600052600f602052604060002090565b613222565b9161393f611e3260c0850151151590565b6139b85761ffff1615801591906139a757606083015163ffffffff16905b63ffffffff613970855163ffffffff1690565b94613982602082015163ffffffff1690565b94156139995760a0015161ffff16925b1693929190565b6080015161ffff1692613992565b604083015163ffffffff169061395d565b5060009150819081908190565b6139cd614a1b565b60208101519160005b8351811015613a6c57806139ef6112bd60019387613c17565b613a16613a1173ffffffffffffffffffffffffffffffffffffffff83166112f5565b616e6b565b613a22575b50016139d6565b60405173ffffffffffffffffffffffffffffffffffffffff9190911681527fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758090602090a138613a1b565b5091505160005b8151811015613b5057613a896112bd8284613c17565b9073ffffffffffffffffffffffffffffffffffffffff821615613b26577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef613b1d83613af5613af06112f560019773ffffffffffffffffffffffffffffffffffffffff1690565b616ded565b5060405173ffffffffffffffffffffffffffffffffffffffff90911681529081906020820190565b0390a101613a73565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050565b91908110156135a65760051b0190565b91908110156135a6576060020190565b60405190613b8182611046565b60606020838281520152565b90604051613b9a8161107e565b608060ff600283958054855260018101546020860152015463ffffffff81166040850152818160201c161515606085015260281c161515910152565b601f8260209493601f19938186528686013760008582860101520116010190565b9160206106d4938181520191613bd6565b90816020910312610498573590565b80518210156135a65760209160051b010190565b90600182811c92168015613c74575b6020831014613c4557565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691613c3a565b9060405191826000825492613c9284613c2b565b8084529360018116908115613cfe5750600114613cb7575b50610943925003836110b6565b90506000929192526020600020906000915b818310613ce25750509060206109439282010138613caa565b6020919350806001915483858901015201910190918492613cc9565b602093506109439592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138613caa565b60409067ffffffffffffffff6106d495931681528160208201520191613bd6565b60405190613d6c8261107e565b60006080838281528260208201528260408201528260608201520152565b90604051613d978161107e565b60806fffffffffffffffffffffffffffffffff6001839560ff8154848116875263ffffffff81871c16602088015260a01c1615156040860152015481808216166060850152821c16910152565b67ffffffffffffffff1660005260086020526106d46004604060002001613c7e565b91908110156135a65760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6181360301821215610498570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610498570180359067ffffffffffffffff821161049857602001918160051b3603831361049857565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b81810292918115918404141715613edc57565b613e9a565b91613efb918354906000199060031b92831b921b19161790565b9055565b818110613f0a575050565b60008155600101613eff565b81519167ffffffffffffffff831161102557680100000000000000008311611025576020908254848455808510613f8d575b500190600052602060002060005b838110613f635750505050565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501613f56565b613fa4908460005285846000209182019101613eff565b38613f48565b90805180519067ffffffffffffffff82116110255768010000000000000000821161102557602090845483865580841061405c575b500183600052602060002060005b83811061402c5750505050906003606083614012602061094396015160018601613f16565b614023604082015160028601613f16565b01519101613f16565b600190602061404f855173ffffffffffffffffffffffffffffffffffffffff1690565b9401938184015501613fed565b614073908660005284846000209182019101613eff565b38613fdf565b9160209082815201919060005b8181106140935750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff87356140bc8161049d565b168152019401929101614086565b969492614108946140ec6140fa936106d49b999560808c5260808c0191614079565b9189830360208b0152614079565b918683036040880152614079565b926060818503910152614079565b90929192614122614a1b565b60005b8181106141a55750505060005b81811061413e57505050565b8067ffffffffffffffff614158611c806001948688613b54565b60006141788267ffffffffffffffff16600052600f602052604060002090565b55167f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8600080a201614132565b6141b3611c808284866142aa565b6141be8284866142aa565b90602082019160a081016127106141de6141d7836142ba565b61ffff1690565b101561426e575060c0016127106141f76141d7836142ba565b101561426e5750907ffae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa104167ffffffffffffffff836142548461424f6001989767ffffffffffffffff16600052600f602052604060002090565b6142ce565b6142656040519283921694826144e5565b0390a201614125565b61427a611cf1916142ba565b7f95f3517a0000000000000000000000000000000000000000000000000000000060005261ffff16600452602490565b91908110156135a65760081b0190565b356106d481610953565b356106d4816127c8565b61449f60c06109439361431681356142e5816135ab565b859063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b6020810135614324816135ab565b67ffffffff0000000085549160201b16807fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff83161786557fffffffffffffffffffffffffffffffffffffffff0000000000000000ffffffff6bffffffff00000000000000006040850135614397816135ab565b60401b169216171784556143f360608201356143b2816135ab565b85547fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff1660609190911b6fffffffff00000000000000000000000016178555565b614445614402608083016142ba565b85547fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff1660809190911b71ffff0000000000000000000000000000000016178555565b61449961445460a083016142ba565b85547fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff1660909190911b73ffff00000000000000000000000000000000000016178555565b016142c4565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016179055565b6109439092919260c061328a8160e084019663ffffffff8135614507816135ab565b16855263ffffffff602082013561451d816135ab565b16602086015263ffffffff6040820135614536816135ab565b16604086015261455861454b606083016135b9565b63ffffffff166060870152565b6145726145676080830161095f565b61ffff166080870152565b61458c61458160a0830161095f565b61ffff1660a0870152565b016127d2565b80549060008155816145a2575050565b6000526020600020908101905b8181106145ba575050565b600081556001016145af565b60056109439160008155600060018201556000600282015560006003820155600481016145f38154613c2b565b9081614602575b505001614592565b81601f6000931160011461461a5750555b38806145fa565b8183526020832061463591601f01861c810190600101613eff565b808252602082209081548360011b906000198560031b1c191617905555614613565b91908110156135a65760051b810135907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee181360301821215610498570190565b9080601f830112156104985781356146ae81611a61565b926146bc60405194856110b6565b81845260208085019260051b820101918383116104985760208201905b8382106146e857505050505090565b813567ffffffffffffffff81116104985760209161470b87848094880101611168565b8152019101906146d9565b61012081360312610498576040519061472e8261107e565b61473781610938565b8252602081013567ffffffffffffffff81116104985761475a9036908301614697565b602083015260408101359067ffffffffffffffff8211610498576147846147a59236908301611168565b60408401526147963660608301612861565b606084015260c0369101612861565b608082015290565b8151815460208401516040850151608091821b73ffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffff000000000000000000000000000000000000000000909416939093179290921791151560a01b74ff000000000000000000000000000000000000000016919091178355606084015193810151901b7fffffffffffffffffffffffffffffffff0000000000000000000000000000000016921691909117600190910155565b9190601f811161488357505050565b610943926000526020600020906020601f840160051c830193106148af575b601f0160051c0190613eff565b90915081906148a2565b919091825167ffffffffffffffff8111611025576148e1816148db8454613c2b565b84614874565b6020601f821160011461491d578190613efb939495600092614912575b50506000198260011b9260031b1c19161790565b0151905038806148fe565b601f1982169061493284600052602060002090565b9160005b81811061496e57509583600195969710614955575b505050811b019055565b015160001960f88460031b161c1916905538808061494b565b9192602060018192868b015181550194019201614936565b6149ea6149b56109439597969467ffffffffffffffff60a0951684526101006020850152610100840190610682565b9660408301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b01906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b73ffffffffffffffffffffffffffffffffffffffff600154163303614a3c57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b600060808201614a7b611e326107ea836133be565b614c3e57506020820191614b1c6020614ac1614a99612dc2876133b4565b60801b7fffffffffffffffffffffffffffffffff000000000000000000000000000000001690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116600482015291829081906024820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610dc9578391614c1f575b50614bf757614b7b614b76846133b4565b6163aa565b614b84836133b4565b90614b9d611e3260a08301936111c6610c1586866132a4565b614bb75750506109439291614bb291506133b4565b616442565b614bc192506132a4565b906120f86040519283927f24eb47e500000000000000000000000000000000000000000000000000000000845260048401613bf7565b6004827f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b614c38915060203d602011610dc257610db481836110b6565b38614b65565b614c4a614c8b916133be565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835273ffffffffffffffffffffffffffffffffffffffff16600452602490565b90fd5b916080830190614ca3611e326107ea846133be565b614e4e576020840193614cc06020614ac1614a99612dc2896133b4565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610dc957600091614e2f575b50614e0557614d1b614b76866133b4565b614d24856133b4565b90614d3d611e3260a08301936111c6610c1586866132a4565b614bb757505061ffff1615614df85767ffffffffffffffff90614df3614dc1614dbb86611c80614dab614d917f41a8aa8df7945f0fb8ac5f7d88279638d9dc2ef9a6bf4ec9a53b80681b34aff7999a6133b4565b67ffffffffffffffff16600052600d602052604060002090565b89614db5886133be565b916164f3565b926133be565b946040519384931695836020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b0390a2565b50614bb2610943926133b4565b7f53ad11d80000000000000000000000000000000000000000000000000000000060005260046000fd5b614e48915060203d602011610dc257610db481836110b6565b38614d0a565b611cf1614e5a836133be565b7f961c9a4f0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b60405190614ea982611046565b6000825260208201600081526020820151917fffffffff000000000000000000000000000000000000000000000000000000006028602483015160e01c92015193167fb148ea5f0000000000000000000000000000000000000000000000000000000081141580614f53575b614f26575063ffffffff1683525290565b7fa176027f0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b507f3047587c00000000000000000000000000000000000000000000000000000000811415614f15565b805161011881106151d7575060048101517f000000000000000000000000000000000000000000000000000000000000000063ffffffff811663ffffffff83160361519e575050600881015190600c81015191608c82015191609081015193609482015160b88301519360f860d885015194015191615000895163ffffffff1690565b63ffffffff811663ffffffff84160361516457507f000000000000000000000000000000000000000000000000000000000000000063ffffffff811663ffffffff86160361512a57506107d063ffffffff8916036150f0576107d063ffffffff8216036150b75750916150799593916020979593615e31565b910151808203615087575050565b7f7be225b60000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b7f0389caa2000000000000000000000000000000000000000000000000000000006000526107d060045263ffffffff1660245260446000fd5b7f22e102a0000000000000000000000000000000000000000000000000000000006000526107d060045263ffffffff881660245260446000fd5b7f77e480260000000000000000000000000000000000000000000000000000000060005263ffffffff908116600452841660245260446000fd5b7fe366a1170000000000000000000000000000000000000000000000000000000060005263ffffffff908116600452821660245260446000fd5b7f960693cd0000000000000000000000000000000000000000000000000000000060005263ffffffff9081166004521660245260446000fd5b7f758b22cc0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b91908110156135a65760e0020190565b356106d4816127dd565b919060005b8181106152305750509050565b61523b818386615204565b615244816133b4565b615250611e328261386e565b6123ea57816152e6615346926152ec602060019796016152786152733683612861565b616148565b6152e66152998467ffffffffffffffff16600052600c602052604060002090565b9182546152b96152b08263ffffffff9060801c1690565b63ffffffff1690565b159081615539575b81615511575b816154f6575b816154e1575b50806154d2575b61547d575b3690612861565b9061673c565b6153016080840191614d916152733685612861565b9283546153186152b08263ffffffff9060801c1690565b159081615446575b8161541e575b81615403575b816153ee575b50806153df575b61534c575b503690612861565b01615223565b61535b60a06153989201615214565b84906fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b82547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff00000000000000000000000000000000161783553861533e565b506153e9826142c4565b615339565b6153fd915060a01c60ff161590565b38615332565b6fffffffffffffffffffffffffffffffff811615915061532c565b90506fffffffffffffffffffffffffffffffff61543e8987015460801c90565b161590615326565b90506fffffffffffffffffffffffffffffffff615475898701546fffffffffffffffffffffffffffffffff1690565b161590615320565b61548c61535b60408901615214565b82547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff00000000000000000000000000000000161783556152df565b506154dc816142c4565b6152da565b6154f0915060a01c60ff161590565b386152d3565b6fffffffffffffffffffffffffffffffff81161591506152cd565b90506fffffffffffffffffffffffffffffffff6155318c86015460801c90565b1615906152c7565b90506fffffffffffffffffffffffffffffffff6155688c8601546fffffffffffffffffffffffffffffffff1690565b1615906152c1565b80518015615615576020036155d85761559260208251830101602083016137b9565b9060ff82116155a2575060ff1690565b6120f8906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352600483016106c3565b6120f8906040519182917f953576f70000000000000000000000000000000000000000000000000000000083526020600484018181520190610682565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff8211613edc57565b60ff16604d8111613edc57600a0a90565b801561566d576000190490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b811561566d570490565b907f000000000000000000000000000000000000000000000000000000000000000060ff811660ff831681811461578c5711615761576156e6828261563b565b91604d60ff8416118015615748575b61570e575050906157086106d49261564f565b90613ec9565b7fa9cb113d0000000000000000000000000000000000000000000000000000000060005260ff908116600452166024525060445260646000fd5b5061575a6157558461564f565b615660565b84116156f5565b61576b818361563b565b91604d60ff84161161570e575050906157866106d49261564f565b9061569c565b5050505090565b60409067ffffffffffffffff6106d494931681528160208201520190610682565b9080511561164c578051602082012067ffffffffffffffff8316928360005260086020526157e9826005604060002001616e22565b1561583d5750816158317f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9361582c614df3946000526009602052604060002090565b6148b9565b604051918291826106c3565b90506120f86040519283927f393b8ad200000000000000000000000000000000000000000000000000000000845260048401615793565b906040519182815491828252602082019060005260206000209260005b8181106158a6575050610943925003836110b6565b845473ffffffffffffffffffffffffffffffffffffffff16835260019485019487945060209093019201615891565b91908201809211613edc57565b906158ec82611a61565b6158f960405191826110b6565b828152601f196159098294611a61565b0190602036910137565b61591c90615874565b9160055480151591826159ec575b5050615934575090565b61593d90615874565b80518061594957505090565b61595a61595f9184959394516158d5565b6158e2565b9060005b84518110156159aa57806159a461597f6112bd60019489613c17565b6159898387613c17565b9073ffffffffffffffffffffffffffffffffffffffff169052565b01615963565b509160005b81518110156159e557806159df6159cb6112bd60019486613c17565b6159896159d9848a516158d5565b87613c17565b016159af565b5090925050565b10159050388061592a565b67ffffffffffffffff166000818152600760205260409020549092919015615af95791615af660e092615ac285615a4e7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b97616148565b846000526008602052615a6581604060002061673c565b615a6e83616148565b846000526008602052615a8883600260406000200161673c565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60009060808101615b3d611e326107ea836133be565b615c2957506020810190615b5b6020614ac1614a99612dc2866133b4565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610dc9578491615c0a575b50615be2576109439291606082615bc8615bc36040615bdd96016133be565b6169db565b615bd4614b76846133b4565b013592506133b4565b616a5d565b6004837f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b615c23915060203d602011610dc257610db481836110b6565b38615ba4565b614c8b614c4a84926133be565b6080810190615c4a611e326107ea846133be565b614e4e576020810190615c676020614ac1614a99612dc2866133b4565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610dc957600091615db9575b50614e055780615cc8615bc36040606094016133be565b615cd4614b76846133b4565b01359261ffff81169081151580615da1575b15615d9057600b5461ffff16918211615d595750507f7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b291614df3614dc1614dbb84611c80614dab615d3f67ffffffffffffffff986133b4565b67ffffffffffffffff16600052600c602052604060002090565b7f7911d95b0000000000000000000000000000000000000000000000000000000060005261ffff9081166004521660245260446000fd5b5050610943929150615bdd906133b4565b50615db26141d7600b5461ffff1690565b1515615ce6565b615dd2915060203d602011610dc257610db481836110b6565b38615cb1565b949290939163ffffffff90604051958260208801981688526040870152166060850152608084015260a083015260c0820152600060e08201526107d06101008201526101008152615e2b610120826110b6565b51902090565b959263ffffffff8095929693604051978260208a019a168a526040890152166060870152608086015260a085015260c0840152600060e0840152166101008201526101008152615e2b610120826110b6565b602081519101517fffffffff00000000000000000000000000000000000000000000000000000000604051927fb148ea5f00000000000000000000000000000000000000000000000000000000602085015260e01b1660248301526028820152602881526106d46048826110b6565b906000198201918211613edc57565b91908203918211613edc57565b615f16613d5f565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff82511690602083019163ffffffff8351164203428111613edc57615f7a906fffffffffffffffffffffffffffffffff60808701511690613ec9565b8101809111613edc57615fa06fffffffffffffffffffffffffffffffff92918392616e59565b161682524263ffffffff16905290565b90615fdd613929615fc3602085016133b4565b67ffffffffffffffff16600052600f602052604060002090565b90615fee611e3260c0840151151590565b6160465761ffff91908216156160395760a0015161ffff165b1680156160315761602b61602360606106d49401359283613ec9565b612710900490565b90615f01565b506060013590565b6080015161ffff16616007565b50506060013590565b80519060005b82811061606157505050565b60018101808211613edc575b83811061607d5750600101616055565b73ffffffffffffffffffffffffffffffffffffffff61609c8385613c17565b51166160ae6112f56112bd8487613c17565b146160bb5760010161606d565b611cf16160cb6112bd8486613c17565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b6109439092919260608101936fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b8051156161ec5760408101516fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff6161ac61619760208501516fffffffffffffffffffffffffffffffff1690565b6fffffffffffffffffffffffffffffffff1690565b9116116161b65750565b6120f8906040519182917f8020d1240000000000000000000000000000000000000000000000000000000083526004830161610d565b6fffffffffffffffffffffffffffffffff61621a60408301516fffffffffffffffffffffffffffffffff1690565b1615801590616261575b61622b5750565b6120f8906040519182917fd68af9cc0000000000000000000000000000000000000000000000000000000083526004830161610d565b5061628261619760208301516fffffffffffffffffffffffffffffffff1690565b1515616224565b1561629057565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b604051906006548083528260208101600660005260206000209260005b818110616346575050610943925003836110b6565b8454835260019485019487945060209093019201616331565b906040519182815491828252602082019060005260206000209260005b818110616391575050610943925003836110b6565b845483526001948501948794506020909301920161637c565b67ffffffffffffffff166163cb816000526007602052604060002054151590565b156164155750336000526011602052604060002054156163e757565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fa9902c7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b67ffffffffffffffff7f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c911691826000526008602052806164c3600260406000200173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169283916164f3565b6040805173ffffffffffffffffffffffffffffffffffffffff909216825260208201929092529081908101614df3565b8054939290919060ff60a086901c16158015616734575b61672d576165296fffffffffffffffffffffffffffffffff8616616197565b90600184019586546165696165636152b0616556616197856fffffffffffffffffffffffffffffffff1690565b9460801c63ffffffff1690565b42615f01565b80616699575b505083811061664e57508282106165cf57506109439394506165949161619791615f01565b6fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b90616606611cf1936166016165f2846165ec6161978c5460801c90565b93615f01565b6165fb83615ef2565b906158d5565b61569c565b7fd0c8d23a0000000000000000000000000000000000000000000000000000000060005260045260245273ffffffffffffffffffffffffffffffffffffffff16604452606490565b7f1a76572a00000000000000000000000000000000000000000000000000000000600052600452602483905273ffffffffffffffffffffffffffffffffffffffff1660445260646000fd5b828592939511616703576166b36161976166ba9460801c90565b9185616d22565b84547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff000000000000000000000000000000001617855591388061656f565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050509050565b50811561650a565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c199161690c61340d92805461677e6165636152b08363ffffffff9060801c1690565b9081616918575b50506168c660016167a960208601516fffffffffffffffffffffffffffffffff1690565b926168346167f76161976fffffffffffffffffffffffffffffffff6167de85546fffffffffffffffffffffffffffffffff1690565b166fffffffffffffffffffffffffffffffff8816616e59565b82906fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b6168876168418751151590565b82547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178255565b019182906fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b604083015181546fffffffffffffffffffffffffffffffff1660809190911b7fffffffffffffffffffffffffffffffff0000000000000000000000000000000016179055565b6040519182918261610d565b6161976167f7916fffffffffffffffffffffffffffffffff61698c616993958261698560018a0154928261697e616977616961876fffffffffffffffffffffffffffffffff1690565b996fffffffffffffffffffffffffffffffff1690565b9560801c90565b1690613ec9565b91166158d5565b9116616e59565b80547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff00000000000000000000000000000000161781553880616785565b7f0000000000000000000000000000000000000000000000000000000000000000616a035750565b73ffffffffffffffffffffffffffffffffffffffff1680600052600360205260406000205415616a305750565b7fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b67ffffffffffffffff7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944911691826000526008602052806164c3604060002073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169283916164f3565b80548210156135a65760005260206000200190600090565b80548015616b1b576000190190616b0a8282616adb565b60001982549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600081815260036020526040902054908115616beb57600019820190828211613edc57600254926000198401938411613edc578383600095616baa9503616bb0575b505050616b996002616af3565b600390600052602052604060002090565b55600190565b616b99616bdc91616bd2616bc8616be2956002616adb565b90549060031b1c90565b9283916002616adb565b90613ee1565b55388080616b8c565b5050600090565b600081815260076020526040902054908115616beb57600019820190828211613edc57600654926000198401938411613edc578383600095616baa9503616c52575b505050616c416006616af3565b600790600052602052604060002090565b616c41616bdc91616c6a616bc8616c74956006616adb565b9283916006616adb565b55388080616c34565b6001810191806000528260205260406000205492831515600014616d19576000198401848111613edc578354936000198501948511613edc576000958583616baa97616cd19503616ce0575b505050616af3565b90600052602052604060002090565b616d00616bdc91616cf7616bc8616d109588616adb565b92839187616adb565b8590600052602052604060002090565b55388080616cc9565b50505050600090565b92616d2d9192613ec9565b8101809111613edc576106d491616e59565b805490680100000000000000008210156110255781616d66916001613efb94018155616adb565b81939154906000199060031b92831b921b19161790565b600081815260036020526040902054616db257616d9b816002616d3f565b600254906000526003602052604060002055600190565b50600090565b600081815260076020526040902054616db257616dd6816006616d3f565b600654906000526007602052604060002055600190565b600081815260116020526040902054616db257616e0b816010616d3f565b601054906000526011602052604060002055600190565b6000828152600182016020526040902054616beb5780616e4483600193616d3f565b80549260005201602052604060002055600190565b9080821015616e66575090565b905090565b600081815260116020526040902054908115616beb57600019820190828211613edc57601054926000198401938411613edc578383616baa9460009603616ecb575b505050616eba6010616af3565b601190600052602052604060002090565b616eba616bdc91616ee3616bc8616eed956010616adb565b9283916010616adb565b55388080616ead565b91929015616f715750815115616f0a575090565b3b15616f135790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015616f845750805190602001fd5b6120f8906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352600483016106c356fea164736f6c634300081a000a",
}

var USDCTokenPoolCCTPV2ABI = USDCTokenPoolCCTPV2MetaData.ABI

var USDCTokenPoolCCTPV2Bin = USDCTokenPoolCCTPV2MetaData.Bin

func DeployUSDCTokenPoolCCTPV2(auth *bind.TransactOpts, backend bind.ContractBackend, tokenMessenger common.Address, cctpMessageTransmitterProxy common.Address, token common.Address, allowlist []common.Address, rmnProxy common.Address, router common.Address) (common.Address, *types.Transaction, *USDCTokenPoolCCTPV2, error) {
	parsed, err := USDCTokenPoolCCTPV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(USDCTokenPoolCCTPV2Bin), backend, tokenMessenger, cctpMessageTransmitterProxy, token, allowlist, rmnProxy, router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &USDCTokenPoolCCTPV2{address: address, abi: *parsed, USDCTokenPoolCCTPV2Caller: USDCTokenPoolCCTPV2Caller{contract: contract}, USDCTokenPoolCCTPV2Transactor: USDCTokenPoolCCTPV2Transactor{contract: contract}, USDCTokenPoolCCTPV2Filterer: USDCTokenPoolCCTPV2Filterer{contract: contract}}, nil
}

type USDCTokenPoolCCTPV2 struct {
	address common.Address
	abi     abi.ABI
	USDCTokenPoolCCTPV2Caller
	USDCTokenPoolCCTPV2Transactor
	USDCTokenPoolCCTPV2Filterer
}

type USDCTokenPoolCCTPV2Caller struct {
	contract *bind.BoundContract
}

type USDCTokenPoolCCTPV2Transactor struct {
	contract *bind.BoundContract
}

type USDCTokenPoolCCTPV2Filterer struct {
	contract *bind.BoundContract
}

type USDCTokenPoolCCTPV2Session struct {
	Contract     *USDCTokenPoolCCTPV2
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type USDCTokenPoolCCTPV2CallerSession struct {
	Contract *USDCTokenPoolCCTPV2Caller
	CallOpts bind.CallOpts
}

type USDCTokenPoolCCTPV2TransactorSession struct {
	Contract     *USDCTokenPoolCCTPV2Transactor
	TransactOpts bind.TransactOpts
}

type USDCTokenPoolCCTPV2Raw struct {
	Contract *USDCTokenPoolCCTPV2
}

type USDCTokenPoolCCTPV2CallerRaw struct {
	Contract *USDCTokenPoolCCTPV2Caller
}

type USDCTokenPoolCCTPV2TransactorRaw struct {
	Contract *USDCTokenPoolCCTPV2Transactor
}

func NewUSDCTokenPoolCCTPV2(address common.Address, backend bind.ContractBackend) (*USDCTokenPoolCCTPV2, error) {
	abi, err := abi.JSON(strings.NewReader(USDCTokenPoolCCTPV2ABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindUSDCTokenPoolCCTPV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2{address: address, abi: abi, USDCTokenPoolCCTPV2Caller: USDCTokenPoolCCTPV2Caller{contract: contract}, USDCTokenPoolCCTPV2Transactor: USDCTokenPoolCCTPV2Transactor{contract: contract}, USDCTokenPoolCCTPV2Filterer: USDCTokenPoolCCTPV2Filterer{contract: contract}}, nil
}

func NewUSDCTokenPoolCCTPV2Caller(address common.Address, caller bind.ContractCaller) (*USDCTokenPoolCCTPV2Caller, error) {
	contract, err := bindUSDCTokenPoolCCTPV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2Caller{contract: contract}, nil
}

func NewUSDCTokenPoolCCTPV2Transactor(address common.Address, transactor bind.ContractTransactor) (*USDCTokenPoolCCTPV2Transactor, error) {
	contract, err := bindUSDCTokenPoolCCTPV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2Transactor{contract: contract}, nil
}

func NewUSDCTokenPoolCCTPV2Filterer(address common.Address, filterer bind.ContractFilterer) (*USDCTokenPoolCCTPV2Filterer, error) {
	contract, err := bindUSDCTokenPoolCCTPV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2Filterer{contract: contract}, nil
}

func bindUSDCTokenPoolCCTPV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := USDCTokenPoolCCTPV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDCTokenPoolCCTPV2.Contract.USDCTokenPoolCCTPV2Caller.contract.Call(opts, result, method, params...)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.USDCTokenPoolCCTPV2Transactor.contract.Transfer(opts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.USDCTokenPoolCCTPV2Transactor.contract.Transact(opts, method, params...)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDCTokenPoolCCTPV2.Contract.contract.Call(opts, result, method, params...)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.contract.Transfer(opts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.contract.Transact(opts, method, params...)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) FINALITYTHRESHOLD(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "FINALITY_THRESHOLD")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) FINALITYTHRESHOLD() (uint32, error) {
	return _USDCTokenPoolCCTPV2.Contract.FINALITYTHRESHOLD(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) FINALITYTHRESHOLD() (uint32, error) {
	return _USDCTokenPoolCCTPV2.Contract.FINALITYTHRESHOLD(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) MAXFEE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "MAX_FEE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) MAXFEE() (uint32, error) {
	return _USDCTokenPoolCCTPV2.Contract.MAXFEE(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) MAXFEE() (uint32, error) {
	return _USDCTokenPoolCCTPV2.Contract.MAXFEE(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) MINUSDCMESSAGELENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "MIN_USDC_MESSAGE_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) MINUSDCMESSAGELENGTH() (*big.Int, error) {
	return _USDCTokenPoolCCTPV2.Contract.MINUSDCMESSAGELENGTH(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) MINUSDCMESSAGELENGTH() (*big.Int, error) {
	return _USDCTokenPoolCCTPV2.Contract.MINUSDCMESSAGELENGTH(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) GetAccumulatedFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "getAccumulatedFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GetAccumulatedFees() (*big.Int, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetAccumulatedFees(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) GetAccumulatedFees() (*big.Int, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetAccumulatedFees(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "getAllAuthorizedCallers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetAllAuthorizedCallers(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetAllAuthorizedCallers(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GetAllowList() ([]common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetAllowList(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) GetAllowList() ([]common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetAllowList(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GetAllowListEnabled() (bool, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetAllowListEnabled(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) GetAllowListEnabled() (bool, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetAllowListEnabled(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "getCurrentInboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetCurrentInboundRateLimiterState(&_USDCTokenPoolCCTPV2.CallOpts, remoteChainSelector)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetCurrentInboundRateLimiterState(&_USDCTokenPoolCCTPV2.CallOpts, remoteChainSelector)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "getCurrentOutboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetCurrentOutboundRateLimiterState(&_USDCTokenPoolCCTPV2.CallOpts, remoteChainSelector)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetCurrentOutboundRateLimiterState(&_USDCTokenPoolCCTPV2.CallOpts, remoteChainSelector)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) GetDomain(opts *bind.CallOpts, chainSelector uint64) (USDCTokenPoolDomain, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "getDomain", chainSelector)

	if err != nil {
		return *new(USDCTokenPoolDomain), err
	}

	out0 := *abi.ConvertType(out[0], new(USDCTokenPoolDomain)).(*USDCTokenPoolDomain)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GetDomain(chainSelector uint64) (USDCTokenPoolDomain, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetDomain(&_USDCTokenPoolCCTPV2.CallOpts, chainSelector)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) GetDomain(chainSelector uint64) (USDCTokenPoolDomain, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetDomain(&_USDCTokenPoolCCTPV2.CallOpts, chainSelector)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

	error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "getDynamicConfig")

	outstruct := new(GetDynamicConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Router = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ThresholdAmountForAdditionalCCVs = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _USDCTokenPoolCCTPV2.Contract.GetDynamicConfig(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _USDCTokenPoolCCTPV2.Contract.GetDynamicConfig(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, finality uint16, arg4 []byte) (GetFee,

	error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "getFee", arg0, destChainSelector, arg2, finality, arg4)

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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GetFee(arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, finality uint16, arg4 []byte) (GetFee,

	error) {
	return _USDCTokenPoolCCTPV2.Contract.GetFee(&_USDCTokenPoolCCTPV2.CallOpts, arg0, destChainSelector, arg2, finality, arg4)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, finality uint16, arg4 []byte) (GetFee,

	error) {
	return _USDCTokenPoolCCTPV2.Contract.GetFee(&_USDCTokenPoolCCTPV2.CallOpts, arg0, destChainSelector, arg2, finality, arg4)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "getRateLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GetRateLimitAdmin() (common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetRateLimitAdmin(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) GetRateLimitAdmin() (common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetRateLimitAdmin(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "getRemotePools", remoteChainSelector)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetRemotePools(&_USDCTokenPoolCCTPV2.CallOpts, remoteChainSelector)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetRemotePools(&_USDCTokenPoolCCTPV2.CallOpts, remoteChainSelector)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "getRemoteToken", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetRemoteToken(&_USDCTokenPoolCCTPV2.CallOpts, remoteChainSelector)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetRemoteToken(&_USDCTokenPoolCCTPV2.CallOpts, remoteChainSelector)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) GetRequiredCCVs(opts *bind.CallOpts, arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "getRequiredCCVs", arg0, remoteChainSelector, amount, arg3, arg4, direction)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GetRequiredCCVs(arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetRequiredCCVs(&_USDCTokenPoolCCTPV2.CallOpts, arg0, remoteChainSelector, amount, arg3, arg4, direction)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) GetRequiredCCVs(arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetRequiredCCVs(&_USDCTokenPoolCCTPV2.CallOpts, arg0, remoteChainSelector, amount, arg3, arg4, direction)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GetRmnProxy() (common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetRmnProxy(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) GetRmnProxy() (common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetRmnProxy(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GetSupportedChains() ([]uint64, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetSupportedChains(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) GetSupportedChains() ([]uint64, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetSupportedChains(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GetToken() (common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetToken(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) GetToken() (common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetToken(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) GetTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "getTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GetTokenDecimals() (uint8, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetTokenDecimals(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) GetTokenDecimals() (uint8, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetTokenDecimals(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "getTokenTransferFeeConfig", arg0, destChainSelector, arg2, arg3, arg4)

	if err != nil {
		return *new(IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolV2TokenTransferFeeConfig)).(*IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetTokenTransferFeeConfig(&_USDCTokenPoolCCTPV2.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetTokenTransferFeeConfig(&_USDCTokenPoolCCTPV2.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) ILocalDomainIdentifier(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "i_localDomainIdentifier")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) ILocalDomainIdentifier() (uint32, error) {
	return _USDCTokenPoolCCTPV2.Contract.ILocalDomainIdentifier(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) ILocalDomainIdentifier() (uint32, error) {
	return _USDCTokenPoolCCTPV2.Contract.ILocalDomainIdentifier(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) IMessageTransmitterProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "i_messageTransmitterProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) IMessageTransmitterProxy() (common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.IMessageTransmitterProxy(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) IMessageTransmitterProxy() (common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.IMessageTransmitterProxy(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) ISupportedUSDCVersion(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "i_supportedUSDCVersion")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) ISupportedUSDCVersion() (uint32, error) {
	return _USDCTokenPoolCCTPV2.Contract.ISupportedUSDCVersion(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) ISupportedUSDCVersion() (uint32, error) {
	return _USDCTokenPoolCCTPV2.Contract.ISupportedUSDCVersion(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) ITokenMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "i_tokenMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) ITokenMessenger() (common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.ITokenMessenger(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) ITokenMessenger() (common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.ITokenMessenger(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "isRemotePool", remoteChainSelector, remotePoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _USDCTokenPoolCCTPV2.Contract.IsRemotePool(&_USDCTokenPoolCCTPV2.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _USDCTokenPoolCCTPV2.Contract.IsRemotePool(&_USDCTokenPoolCCTPV2.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _USDCTokenPoolCCTPV2.Contract.IsSupportedChain(&_USDCTokenPoolCCTPV2.CallOpts, remoteChainSelector)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _USDCTokenPoolCCTPV2.Contract.IsSupportedChain(&_USDCTokenPoolCCTPV2.CallOpts, remoteChainSelector)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) IsSupportedToken(token common.Address) (bool, error) {
	return _USDCTokenPoolCCTPV2.Contract.IsSupportedToken(&_USDCTokenPoolCCTPV2.CallOpts, token)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _USDCTokenPoolCCTPV2.Contract.IsSupportedToken(&_USDCTokenPoolCCTPV2.CallOpts, token)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) Owner() (common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.Owner(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) Owner() (common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.Owner(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _USDCTokenPoolCCTPV2.Contract.SupportsInterface(&_USDCTokenPoolCCTPV2.CallOpts, interfaceId)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _USDCTokenPoolCCTPV2.Contract.SupportsInterface(&_USDCTokenPoolCCTPV2.CallOpts, interfaceId)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) TypeAndVersion() (string, error) {
	return _USDCTokenPoolCCTPV2.Contract.TypeAndVersion(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) TypeAndVersion() (string, error) {
	return _USDCTokenPoolCCTPV2.Contract.TypeAndVersion(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "acceptOwnership")
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) AcceptOwnership() (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.AcceptOwnership(&_USDCTokenPoolCCTPV2.TransactOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.AcceptOwnership(&_USDCTokenPoolCCTPV2.TransactOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "addRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.AddRemotePool(&_USDCTokenPoolCCTPV2.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.AddRemotePool(&_USDCTokenPoolCCTPV2.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.ApplyAllowListUpdates(&_USDCTokenPoolCCTPV2.TransactOpts, removes, adds)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.ApplyAllowListUpdates(&_USDCTokenPoolCCTPV2.TransactOpts, removes, adds)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "applyAuthorizedCallerUpdates", authorizedCallerArgs)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.ApplyAuthorizedCallerUpdates(&_USDCTokenPoolCCTPV2.TransactOpts, authorizedCallerArgs)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.ApplyAuthorizedCallerUpdates(&_USDCTokenPoolCCTPV2.TransactOpts, authorizedCallerArgs)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) ApplyCCVConfigUpdates(opts *bind.TransactOpts, ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "applyCCVConfigUpdates", ccvConfigArgs)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) ApplyCCVConfigUpdates(ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.ApplyCCVConfigUpdates(&_USDCTokenPoolCCTPV2.TransactOpts, ccvConfigArgs)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) ApplyCCVConfigUpdates(ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.ApplyCCVConfigUpdates(&_USDCTokenPoolCCTPV2.TransactOpts, ccvConfigArgs)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "applyChainUpdates", remoteChainSelectorsToRemove, chainsToAdd)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.ApplyChainUpdates(&_USDCTokenPoolCCTPV2.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.ApplyChainUpdates(&_USDCTokenPoolCCTPV2.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) ApplyFinalityConfigUpdates(opts *bind.TransactOpts, minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "applyFinalityConfigUpdates", minBlockConfirmation, rateLimitConfigArgs)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) ApplyFinalityConfigUpdates(minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.ApplyFinalityConfigUpdates(&_USDCTokenPoolCCTPV2.TransactOpts, minBlockConfirmation, rateLimitConfigArgs)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) ApplyFinalityConfigUpdates(minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.ApplyFinalityConfigUpdates(&_USDCTokenPoolCCTPV2.TransactOpts, minBlockConfirmation, rateLimitConfigArgs)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "applyTokenTransferFeeConfigUpdates", tokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.ApplyTokenTransferFeeConfigUpdates(&_USDCTokenPoolCCTPV2.TransactOpts, tokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.ApplyTokenTransferFeeConfigUpdates(&_USDCTokenPoolCCTPV2.TransactOpts, tokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.LockOrBurn(&_USDCTokenPoolCCTPV2.TransactOpts, lockOrBurnIn)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.LockOrBurn(&_USDCTokenPoolCCTPV2.TransactOpts, lockOrBurnIn)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, finality, arg2)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.LockOrBurn0(&_USDCTokenPoolCCTPV2.TransactOpts, lockOrBurnIn, finality, arg2)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.LockOrBurn0(&_USDCTokenPoolCCTPV2.TransactOpts, lockOrBurnIn, finality, arg2)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.ReleaseOrMint(&_USDCTokenPoolCCTPV2.TransactOpts, releaseOrMintIn)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.ReleaseOrMint(&_USDCTokenPoolCCTPV2.TransactOpts, releaseOrMintIn)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn, finality)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.ReleaseOrMint0(&_USDCTokenPoolCCTPV2.TransactOpts, releaseOrMintIn, finality)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.ReleaseOrMint0(&_USDCTokenPoolCCTPV2.TransactOpts, releaseOrMintIn, finality)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "removeRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.RemoveRemotePool(&_USDCTokenPoolCCTPV2.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.RemoveRemotePool(&_USDCTokenPoolCCTPV2.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "setChainRateLimiterConfig", remoteChainSelector, outboundConfig, inboundConfig)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.SetChainRateLimiterConfig(&_USDCTokenPoolCCTPV2.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.SetChainRateLimiterConfig(&_USDCTokenPoolCCTPV2.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "setChainRateLimiterConfigs", remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.SetChainRateLimiterConfigs(&_USDCTokenPoolCCTPV2.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.SetChainRateLimiterConfigs(&_USDCTokenPoolCCTPV2.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) SetCustomFinalityRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "setCustomFinalityRateLimitConfig", rateLimitConfigArgs)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) SetCustomFinalityRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.SetCustomFinalityRateLimitConfig(&_USDCTokenPoolCCTPV2.TransactOpts, rateLimitConfigArgs)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) SetCustomFinalityRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.SetCustomFinalityRateLimitConfig(&_USDCTokenPoolCCTPV2.TransactOpts, rateLimitConfigArgs)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) SetDomains(opts *bind.TransactOpts, domains []USDCTokenPoolDomainUpdate) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "setDomains", domains)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) SetDomains(domains []USDCTokenPoolDomainUpdate) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.SetDomains(&_USDCTokenPoolCCTPV2.TransactOpts, domains)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) SetDomains(domains []USDCTokenPoolDomainUpdate) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.SetDomains(&_USDCTokenPoolCCTPV2.TransactOpts, domains)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) SetDynamicConfig(opts *bind.TransactOpts, router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "setDynamicConfig", router, thresholdAmountForAdditionalCCVs)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) SetDynamicConfig(router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.SetDynamicConfig(&_USDCTokenPoolCCTPV2.TransactOpts, router, thresholdAmountForAdditionalCCVs)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) SetDynamicConfig(router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.SetDynamicConfig(&_USDCTokenPoolCCTPV2.TransactOpts, router, thresholdAmountForAdditionalCCVs)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "setRateLimitAdmin", rateLimitAdmin)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.SetRateLimitAdmin(&_USDCTokenPoolCCTPV2.TransactOpts, rateLimitAdmin)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.SetRateLimitAdmin(&_USDCTokenPoolCCTPV2.TransactOpts, rateLimitAdmin)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "transferOwnership", to)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.TransferOwnership(&_USDCTokenPoolCCTPV2.TransactOpts, to)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.TransferOwnership(&_USDCTokenPoolCCTPV2.TransactOpts, to)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) WithdrawFees(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "withdrawFees", recipient)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) WithdrawFees(recipient common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.WithdrawFees(&_USDCTokenPoolCCTPV2.TransactOpts, recipient)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) WithdrawFees(recipient common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.WithdrawFees(&_USDCTokenPoolCCTPV2.TransactOpts, recipient)
}

type USDCTokenPoolCCTPV2AllowListAddIterator struct {
	Event *USDCTokenPoolCCTPV2AllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2AllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2AllowListAdd)
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
		it.Event = new(USDCTokenPoolCCTPV2AllowListAdd)
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

func (it *USDCTokenPoolCCTPV2AllowListAddIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2AllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2AllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterAllowListAdd(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2AllowListAddIterator, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2AllowListAddIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2AllowListAdd) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2AllowListAdd)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseAllowListAdd(log types.Log) (*USDCTokenPoolCCTPV2AllowListAdd, error) {
	event := new(USDCTokenPoolCCTPV2AllowListAdd)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2AllowListRemoveIterator struct {
	Event *USDCTokenPoolCCTPV2AllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2AllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2AllowListRemove)
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
		it.Event = new(USDCTokenPoolCCTPV2AllowListRemove)
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

func (it *USDCTokenPoolCCTPV2AllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2AllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2AllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterAllowListRemove(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2AllowListRemoveIterator, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2AllowListRemoveIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2AllowListRemove) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2AllowListRemove)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseAllowListRemove(log types.Log) (*USDCTokenPoolCCTPV2AllowListRemove, error) {
	event := new(USDCTokenPoolCCTPV2AllowListRemove)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2AuthorizedCallerAddedIterator struct {
	Event *USDCTokenPoolCCTPV2AuthorizedCallerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2AuthorizedCallerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2AuthorizedCallerAdded)
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
		it.Event = new(USDCTokenPoolCCTPV2AuthorizedCallerAdded)
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

func (it *USDCTokenPoolCCTPV2AuthorizedCallerAddedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2AuthorizedCallerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2AuthorizedCallerAdded struct {
	Caller common.Address
	Raw    types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2AuthorizedCallerAddedIterator, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2AuthorizedCallerAddedIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "AuthorizedCallerAdded", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2AuthorizedCallerAdded) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2AuthorizedCallerAdded)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseAuthorizedCallerAdded(log types.Log) (*USDCTokenPoolCCTPV2AuthorizedCallerAdded, error) {
	event := new(USDCTokenPoolCCTPV2AuthorizedCallerAdded)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2AuthorizedCallerRemovedIterator struct {
	Event *USDCTokenPoolCCTPV2AuthorizedCallerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2AuthorizedCallerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2AuthorizedCallerRemoved)
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
		it.Event = new(USDCTokenPoolCCTPV2AuthorizedCallerRemoved)
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

func (it *USDCTokenPoolCCTPV2AuthorizedCallerRemovedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2AuthorizedCallerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2AuthorizedCallerRemoved struct {
	Caller common.Address
	Raw    types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2AuthorizedCallerRemovedIterator, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2AuthorizedCallerRemovedIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "AuthorizedCallerRemoved", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2AuthorizedCallerRemoved) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2AuthorizedCallerRemoved)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseAuthorizedCallerRemoved(log types.Log) (*USDCTokenPoolCCTPV2AuthorizedCallerRemoved, error) {
	event := new(USDCTokenPoolCCTPV2AuthorizedCallerRemoved)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2CCVConfigUpdatedIterator struct {
	Event *USDCTokenPoolCCTPV2CCVConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2CCVConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2CCVConfigUpdated)
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
		it.Event = new(USDCTokenPoolCCTPV2CCVConfigUpdated)
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

func (it *USDCTokenPoolCCTPV2CCVConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2CCVConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2CCVConfigUpdated struct {
	RemoteChainSelector             uint64
	OutboundCCVs                    []common.Address
	OutboundCCVsToAddAboveThreshold []common.Address
	InboundCCVs                     []common.Address
	InboundCCVsToAddAboveThreshold  []common.Address
	Raw                             types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2CCVConfigUpdatedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2CCVConfigUpdatedIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "CCVConfigUpdated", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2CCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2CCVConfigUpdated)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseCCVConfigUpdated(log types.Log) (*USDCTokenPoolCCTPV2CCVConfigUpdated, error) {
	event := new(USDCTokenPoolCCTPV2CCVConfigUpdated)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2ChainAddedIterator struct {
	Event *USDCTokenPoolCCTPV2ChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2ChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2ChainAdded)
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
		it.Event = new(USDCTokenPoolCCTPV2ChainAdded)
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

func (it *USDCTokenPoolCCTPV2ChainAddedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2ChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2ChainAdded struct {
	RemoteChainSelector       uint64
	RemoteToken               []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterChainAdded(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2ChainAddedIterator, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2ChainAddedIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2ChainAdded) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2ChainAdded)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "ChainAdded", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseChainAdded(log types.Log) (*USDCTokenPoolCCTPV2ChainAdded, error) {
	event := new(USDCTokenPoolCCTPV2ChainAdded)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2ChainConfiguredIterator struct {
	Event *USDCTokenPoolCCTPV2ChainConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2ChainConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2ChainConfigured)
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
		it.Event = new(USDCTokenPoolCCTPV2ChainConfigured)
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

func (it *USDCTokenPoolCCTPV2ChainConfiguredIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2ChainConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2ChainConfigured struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterChainConfigured(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2ChainConfiguredIterator, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2ChainConfiguredIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "ChainConfigured", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2ChainConfigured) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2ChainConfigured)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseChainConfigured(log types.Log) (*USDCTokenPoolCCTPV2ChainConfigured, error) {
	event := new(USDCTokenPoolCCTPV2ChainConfigured)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2ChainRemovedIterator struct {
	Event *USDCTokenPoolCCTPV2ChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2ChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2ChainRemoved)
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
		it.Event = new(USDCTokenPoolCCTPV2ChainRemoved)
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

func (it *USDCTokenPoolCCTPV2ChainRemovedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2ChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2ChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterChainRemoved(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2ChainRemovedIterator, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2ChainRemovedIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2ChainRemoved) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2ChainRemoved)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseChainRemoved(log types.Log) (*USDCTokenPoolCCTPV2ChainRemoved, error) {
	event := new(USDCTokenPoolCCTPV2ChainRemoved)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2ConfigChangedIterator struct {
	Event *USDCTokenPoolCCTPV2ConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2ConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2ConfigChanged)
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
		it.Event = new(USDCTokenPoolCCTPV2ConfigChanged)
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

func (it *USDCTokenPoolCCTPV2ConfigChangedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2ConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2ConfigChanged struct {
	Config RateLimiterConfig
	Raw    types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterConfigChanged(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2ConfigChangedIterator, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2ConfigChangedIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2ConfigChanged) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2ConfigChanged)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseConfigChanged(log types.Log) (*USDCTokenPoolCCTPV2ConfigChanged, error) {
	event := new(USDCTokenPoolCCTPV2ConfigChanged)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2ConfigSetIterator struct {
	Event *USDCTokenPoolCCTPV2ConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2ConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2ConfigSet)
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
		it.Event = new(USDCTokenPoolCCTPV2ConfigSet)
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

func (it *USDCTokenPoolCCTPV2ConfigSetIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2ConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2ConfigSet struct {
	TokenMessenger common.Address
	Raw            types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterConfigSet(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2ConfigSetIterator, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2ConfigSetIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2ConfigSet) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2ConfigSet)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseConfigSet(log types.Log) (*USDCTokenPoolCCTPV2ConfigSet, error) {
	event := new(USDCTokenPoolCCTPV2ConfigSet)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2CustomFinalityMinimumBlockConfirmationUpdatedIterator struct {
	Event *USDCTokenPoolCCTPV2CustomFinalityMinimumBlockConfirmationUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2CustomFinalityMinimumBlockConfirmationUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2CustomFinalityMinimumBlockConfirmationUpdated)
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
		it.Event = new(USDCTokenPoolCCTPV2CustomFinalityMinimumBlockConfirmationUpdated)
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

func (it *USDCTokenPoolCCTPV2CustomFinalityMinimumBlockConfirmationUpdatedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2CustomFinalityMinimumBlockConfirmationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2CustomFinalityMinimumBlockConfirmationUpdated struct {
	MinBlockConfirmation uint16
	Raw                  types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterCustomFinalityMinimumBlockConfirmationUpdated(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2CustomFinalityMinimumBlockConfirmationUpdatedIterator, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "CustomFinalityMinimumBlockConfirmationUpdated")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2CustomFinalityMinimumBlockConfirmationUpdatedIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "CustomFinalityMinimumBlockConfirmationUpdated", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchCustomFinalityMinimumBlockConfirmationUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2CustomFinalityMinimumBlockConfirmationUpdated) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "CustomFinalityMinimumBlockConfirmationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2CustomFinalityMinimumBlockConfirmationUpdated)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "CustomFinalityMinimumBlockConfirmationUpdated", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseCustomFinalityMinimumBlockConfirmationUpdated(log types.Log) (*USDCTokenPoolCCTPV2CustomFinalityMinimumBlockConfirmationUpdated, error) {
	event := new(USDCTokenPoolCCTPV2CustomFinalityMinimumBlockConfirmationUpdated)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "CustomFinalityMinimumBlockConfirmationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2CustomFinalityOutboundRateLimitConsumedIterator struct {
	Event *USDCTokenPoolCCTPV2CustomFinalityOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2CustomFinalityOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2CustomFinalityOutboundRateLimitConsumed)
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
		it.Event = new(USDCTokenPoolCCTPV2CustomFinalityOutboundRateLimitConsumed)
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

func (it *USDCTokenPoolCCTPV2CustomFinalityOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2CustomFinalityOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2CustomFinalityOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterCustomFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2CustomFinalityOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "CustomFinalityOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2CustomFinalityOutboundRateLimitConsumedIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "CustomFinalityOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchCustomFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2CustomFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "CustomFinalityOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2CustomFinalityOutboundRateLimitConsumed)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "CustomFinalityOutboundRateLimitConsumed", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseCustomFinalityOutboundRateLimitConsumed(log types.Log) (*USDCTokenPoolCCTPV2CustomFinalityOutboundRateLimitConsumed, error) {
	event := new(USDCTokenPoolCCTPV2CustomFinalityOutboundRateLimitConsumed)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "CustomFinalityOutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2CustomFinalityTransferInboundRateLimitConsumedIterator struct {
	Event *USDCTokenPoolCCTPV2CustomFinalityTransferInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2CustomFinalityTransferInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2CustomFinalityTransferInboundRateLimitConsumed)
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
		it.Event = new(USDCTokenPoolCCTPV2CustomFinalityTransferInboundRateLimitConsumed)
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

func (it *USDCTokenPoolCCTPV2CustomFinalityTransferInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2CustomFinalityTransferInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2CustomFinalityTransferInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterCustomFinalityTransferInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2CustomFinalityTransferInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "CustomFinalityTransferInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2CustomFinalityTransferInboundRateLimitConsumedIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "CustomFinalityTransferInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchCustomFinalityTransferInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2CustomFinalityTransferInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "CustomFinalityTransferInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2CustomFinalityTransferInboundRateLimitConsumed)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "CustomFinalityTransferInboundRateLimitConsumed", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseCustomFinalityTransferInboundRateLimitConsumed(log types.Log) (*USDCTokenPoolCCTPV2CustomFinalityTransferInboundRateLimitConsumed, error) {
	event := new(USDCTokenPoolCCTPV2CustomFinalityTransferInboundRateLimitConsumed)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "CustomFinalityTransferInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2DomainsSetIterator struct {
	Event *USDCTokenPoolCCTPV2DomainsSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2DomainsSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2DomainsSet)
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
		it.Event = new(USDCTokenPoolCCTPV2DomainsSet)
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

func (it *USDCTokenPoolCCTPV2DomainsSetIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2DomainsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2DomainsSet struct {
	Arg0 []USDCTokenPoolDomainUpdate
	Raw  types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterDomainsSet(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2DomainsSetIterator, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "DomainsSet")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2DomainsSetIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "DomainsSet", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchDomainsSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2DomainsSet) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "DomainsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2DomainsSet)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "DomainsSet", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseDomainsSet(log types.Log) (*USDCTokenPoolCCTPV2DomainsSet, error) {
	event := new(USDCTokenPoolCCTPV2DomainsSet)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "DomainsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2DynamicConfigSetIterator struct {
	Event *USDCTokenPoolCCTPV2DynamicConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2DynamicConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2DynamicConfigSet)
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
		it.Event = new(USDCTokenPoolCCTPV2DynamicConfigSet)
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

func (it *USDCTokenPoolCCTPV2DynamicConfigSetIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2DynamicConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2DynamicConfigSet struct {
	Router                           common.Address
	ThresholdAmountForAdditionalCCVs *big.Int
	Raw                              types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterDynamicConfigSet(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2DynamicConfigSetIterator, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2DynamicConfigSetIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "DynamicConfigSet", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2DynamicConfigSet) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2DynamicConfigSet)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseDynamicConfigSet(log types.Log) (*USDCTokenPoolCCTPV2DynamicConfigSet, error) {
	event := new(USDCTokenPoolCCTPV2DynamicConfigSet)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2InboundRateLimitConsumedIterator struct {
	Event *USDCTokenPoolCCTPV2InboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2InboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2InboundRateLimitConsumed)
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
		it.Event = new(USDCTokenPoolCCTPV2InboundRateLimitConsumed)
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

func (it *USDCTokenPoolCCTPV2InboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2InboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2InboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2InboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2InboundRateLimitConsumedIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "InboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2InboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2InboundRateLimitConsumed)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseInboundRateLimitConsumed(log types.Log) (*USDCTokenPoolCCTPV2InboundRateLimitConsumed, error) {
	event := new(USDCTokenPoolCCTPV2InboundRateLimitConsumed)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2LockedOrBurnedIterator struct {
	Event *USDCTokenPoolCCTPV2LockedOrBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2LockedOrBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2LockedOrBurned)
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
		it.Event = new(USDCTokenPoolCCTPV2LockedOrBurned)
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

func (it *USDCTokenPoolCCTPV2LockedOrBurnedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2LockedOrBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2LockedOrBurned struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2LockedOrBurnedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2LockedOrBurnedIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "LockedOrBurned", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2LockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2LockedOrBurned)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseLockedOrBurned(log types.Log) (*USDCTokenPoolCCTPV2LockedOrBurned, error) {
	event := new(USDCTokenPoolCCTPV2LockedOrBurned)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2OutboundRateLimitConsumedIterator struct {
	Event *USDCTokenPoolCCTPV2OutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2OutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2OutboundRateLimitConsumed)
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
		it.Event = new(USDCTokenPoolCCTPV2OutboundRateLimitConsumed)
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

func (it *USDCTokenPoolCCTPV2OutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2OutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2OutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2OutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2OutboundRateLimitConsumedIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "OutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2OutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2OutboundRateLimitConsumed)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseOutboundRateLimitConsumed(log types.Log) (*USDCTokenPoolCCTPV2OutboundRateLimitConsumed, error) {
	event := new(USDCTokenPoolCCTPV2OutboundRateLimitConsumed)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2OwnershipTransferRequestedIterator struct {
	Event *USDCTokenPoolCCTPV2OwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2OwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2OwnershipTransferRequested)
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
		it.Event = new(USDCTokenPoolCCTPV2OwnershipTransferRequested)
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

func (it *USDCTokenPoolCCTPV2OwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2OwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2OwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*USDCTokenPoolCCTPV2OwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2OwnershipTransferRequestedIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2OwnershipTransferRequested)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseOwnershipTransferRequested(log types.Log) (*USDCTokenPoolCCTPV2OwnershipTransferRequested, error) {
	event := new(USDCTokenPoolCCTPV2OwnershipTransferRequested)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2OwnershipTransferredIterator struct {
	Event *USDCTokenPoolCCTPV2OwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2OwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2OwnershipTransferred)
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
		it.Event = new(USDCTokenPoolCCTPV2OwnershipTransferred)
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

func (it *USDCTokenPoolCCTPV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2OwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*USDCTokenPoolCCTPV2OwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2OwnershipTransferredIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2OwnershipTransferred)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseOwnershipTransferred(log types.Log) (*USDCTokenPoolCCTPV2OwnershipTransferred, error) {
	event := new(USDCTokenPoolCCTPV2OwnershipTransferred)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2PoolFeeWithdrawnIterator struct {
	Event *USDCTokenPoolCCTPV2PoolFeeWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2PoolFeeWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2PoolFeeWithdrawn)
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
		it.Event = new(USDCTokenPoolCCTPV2PoolFeeWithdrawn)
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

func (it *USDCTokenPoolCCTPV2PoolFeeWithdrawnIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2PoolFeeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2PoolFeeWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterPoolFeeWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*USDCTokenPoolCCTPV2PoolFeeWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "PoolFeeWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2PoolFeeWithdrawnIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "PoolFeeWithdrawn", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchPoolFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2PoolFeeWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "PoolFeeWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2PoolFeeWithdrawn)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "PoolFeeWithdrawn", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParsePoolFeeWithdrawn(log types.Log) (*USDCTokenPoolCCTPV2PoolFeeWithdrawn, error) {
	event := new(USDCTokenPoolCCTPV2PoolFeeWithdrawn)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "PoolFeeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2RateLimitAdminSetIterator struct {
	Event *USDCTokenPoolCCTPV2RateLimitAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2RateLimitAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2RateLimitAdminSet)
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
		it.Event = new(USDCTokenPoolCCTPV2RateLimitAdminSet)
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

func (it *USDCTokenPoolCCTPV2RateLimitAdminSetIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2RateLimitAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2RateLimitAdminSet struct {
	RateLimitAdmin common.Address
	Raw            types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterRateLimitAdminSet(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2RateLimitAdminSetIterator, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2RateLimitAdminSetIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "RateLimitAdminSet", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2RateLimitAdminSet) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2RateLimitAdminSet)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseRateLimitAdminSet(log types.Log) (*USDCTokenPoolCCTPV2RateLimitAdminSet, error) {
	event := new(USDCTokenPoolCCTPV2RateLimitAdminSet)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2ReleasedOrMintedIterator struct {
	Event *USDCTokenPoolCCTPV2ReleasedOrMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2ReleasedOrMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2ReleasedOrMinted)
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
		it.Event = new(USDCTokenPoolCCTPV2ReleasedOrMinted)
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

func (it *USDCTokenPoolCCTPV2ReleasedOrMintedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2ReleasedOrMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2ReleasedOrMinted struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Recipient           common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2ReleasedOrMintedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2ReleasedOrMintedIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "ReleasedOrMinted", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2ReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2ReleasedOrMinted)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseReleasedOrMinted(log types.Log) (*USDCTokenPoolCCTPV2ReleasedOrMinted, error) {
	event := new(USDCTokenPoolCCTPV2ReleasedOrMinted)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2RemotePoolAddedIterator struct {
	Event *USDCTokenPoolCCTPV2RemotePoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2RemotePoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2RemotePoolAdded)
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
		it.Event = new(USDCTokenPoolCCTPV2RemotePoolAdded)
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

func (it *USDCTokenPoolCCTPV2RemotePoolAddedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2RemotePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2RemotePoolAdded struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2RemotePoolAddedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2RemotePoolAddedIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "RemotePoolAdded", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2RemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2RemotePoolAdded)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseRemotePoolAdded(log types.Log) (*USDCTokenPoolCCTPV2RemotePoolAdded, error) {
	event := new(USDCTokenPoolCCTPV2RemotePoolAdded)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2RemotePoolRemovedIterator struct {
	Event *USDCTokenPoolCCTPV2RemotePoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2RemotePoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2RemotePoolRemoved)
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
		it.Event = new(USDCTokenPoolCCTPV2RemotePoolRemoved)
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

func (it *USDCTokenPoolCCTPV2RemotePoolRemovedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2RemotePoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2RemotePoolRemoved struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2RemotePoolRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2RemotePoolRemovedIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "RemotePoolRemoved", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2RemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2RemotePoolRemoved)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseRemotePoolRemoved(log types.Log) (*USDCTokenPoolCCTPV2RemotePoolRemoved, error) {
	event := new(USDCTokenPoolCCTPV2RemotePoolRemoved)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2TokenTransferFeeConfigDeletedIterator struct {
	Event *USDCTokenPoolCCTPV2TokenTransferFeeConfigDeleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2TokenTransferFeeConfigDeletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2TokenTransferFeeConfigDeleted)
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
		it.Event = new(USDCTokenPoolCCTPV2TokenTransferFeeConfigDeleted)
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

func (it *USDCTokenPoolCCTPV2TokenTransferFeeConfigDeletedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2TokenTransferFeeConfigDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2TokenTransferFeeConfigDeleted struct {
	DestChainSelector uint64
	Raw               types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*USDCTokenPoolCCTPV2TokenTransferFeeConfigDeletedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2TokenTransferFeeConfigDeletedIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "TokenTransferFeeConfigDeleted", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2TokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2TokenTransferFeeConfigDeleted)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseTokenTransferFeeConfigDeleted(log types.Log) (*USDCTokenPoolCCTPV2TokenTransferFeeConfigDeleted, error) {
	event := new(USDCTokenPoolCCTPV2TokenTransferFeeConfigDeleted)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2TokenTransferFeeConfigUpdatedIterator struct {
	Event *USDCTokenPoolCCTPV2TokenTransferFeeConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2TokenTransferFeeConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2TokenTransferFeeConfigUpdated)
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
		it.Event = new(USDCTokenPoolCCTPV2TokenTransferFeeConfigUpdated)
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

func (it *USDCTokenPoolCCTPV2TokenTransferFeeConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2TokenTransferFeeConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2TokenTransferFeeConfigUpdated struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
	Raw                    types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*USDCTokenPoolCCTPV2TokenTransferFeeConfigUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2TokenTransferFeeConfigUpdatedIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "TokenTransferFeeConfigUpdated", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2TokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2TokenTransferFeeConfigUpdated)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseTokenTransferFeeConfigUpdated(log types.Log) (*USDCTokenPoolCCTPV2TokenTransferFeeConfigUpdated, error) {
	event := new(USDCTokenPoolCCTPV2TokenTransferFeeConfigUpdated)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

func (USDCTokenPoolCCTPV2AllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (USDCTokenPoolCCTPV2AllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (USDCTokenPoolCCTPV2AuthorizedCallerAdded) Topic() common.Hash {
	return common.HexToHash("0xeb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef")
}

func (USDCTokenPoolCCTPV2AuthorizedCallerRemoved) Topic() common.Hash {
	return common.HexToHash("0xc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda77580")
}

func (USDCTokenPoolCCTPV2CCVConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d")
}

func (USDCTokenPoolCCTPV2ChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (USDCTokenPoolCCTPV2ChainConfigured) Topic() common.Hash {
	return common.HexToHash("0x0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b")
}

func (USDCTokenPoolCCTPV2ChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (USDCTokenPoolCCTPV2ConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (USDCTokenPoolCCTPV2ConfigSet) Topic() common.Hash {
	return common.HexToHash("0x2e902d38f15b233cbb63711add0fca4545334d3a169d60c0a616494d7eea9544")
}

func (USDCTokenPoolCCTPV2CustomFinalityMinimumBlockConfirmationUpdated) Topic() common.Hash {
	return common.HexToHash("0x401d6d5b7a131da25eadbfe5bc7b2c80d7a0094d8d026d2d2c452f81e9b2af7e")
}

func (USDCTokenPoolCCTPV2CustomFinalityOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b2")
}

func (USDCTokenPoolCCTPV2CustomFinalityTransferInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x41a8aa8df7945f0fb8ac5f7d88279638d9dc2ef9a6bf4ec9a53b80681b34aff7")
}

func (USDCTokenPoolCCTPV2DomainsSet) Topic() common.Hash {
	return common.HexToHash("0xc97f93e817584952f1c1d633f93784b8430f0633d002f9dcc4de4fe2780424d0")
}

func (USDCTokenPoolCCTPV2DynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f0238")
}

func (USDCTokenPoolCCTPV2InboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (USDCTokenPoolCCTPV2LockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
}

func (USDCTokenPoolCCTPV2OutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944")
}

func (USDCTokenPoolCCTPV2OwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (USDCTokenPoolCCTPV2OwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (USDCTokenPoolCCTPV2PoolFeeWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b859599")
}

func (USDCTokenPoolCCTPV2RateLimitAdminSet) Topic() common.Hash {
	return common.HexToHash("0x44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174")
}

func (USDCTokenPoolCCTPV2ReleasedOrMinted) Topic() common.Hash {
	return common.HexToHash("0xfc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0")
}

func (USDCTokenPoolCCTPV2RemotePoolAdded) Topic() common.Hash {
	return common.HexToHash("0x7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea")
}

func (USDCTokenPoolCCTPV2RemotePoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76")
}

func (USDCTokenPoolCCTPV2TokenTransferFeeConfigDeleted) Topic() common.Hash {
	return common.HexToHash("0x5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8")
}

func (USDCTokenPoolCCTPV2TokenTransferFeeConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xfae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa1041")
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2) Address() common.Address {
	return _USDCTokenPoolCCTPV2.address
}

type USDCTokenPoolCCTPV2Interface interface {
	FINALITYTHRESHOLD(opts *bind.CallOpts) (uint32, error)

	MAXFEE(opts *bind.CallOpts) (uint32, error)

	MINUSDCMESSAGELENGTH(opts *bind.CallOpts) (*big.Int, error)

	GetAccumulatedFees(opts *bind.CallOpts) (*big.Int, error)

	GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetDomain(opts *bind.CallOpts, chainSelector uint64) (USDCTokenPoolDomain, error)

	GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

		error)

	GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, finality uint16, arg4 []byte) (GetFee,

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

	ApplyFinalityConfigUpdates(opts *bind.TransactOpts, minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error)

	ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error)

	SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error)

	SetCustomFinalityRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error)

	SetDomains(opts *bind.TransactOpts, domains []USDCTokenPoolDomainUpdate) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error)

	SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFees(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error)

	FilterAllowListAdd(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2AllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2AllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*USDCTokenPoolCCTPV2AllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2AllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2AllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*USDCTokenPoolCCTPV2AllowListRemove, error)

	FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2AuthorizedCallerAddedIterator, error)

	WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2AuthorizedCallerAdded) (event.Subscription, error)

	ParseAuthorizedCallerAdded(log types.Log) (*USDCTokenPoolCCTPV2AuthorizedCallerAdded, error)

	FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2AuthorizedCallerRemovedIterator, error)

	WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2AuthorizedCallerRemoved) (event.Subscription, error)

	ParseAuthorizedCallerRemoved(log types.Log) (*USDCTokenPoolCCTPV2AuthorizedCallerRemoved, error)

	FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2CCVConfigUpdatedIterator, error)

	WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2CCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCCVConfigUpdated(log types.Log) (*USDCTokenPoolCCTPV2CCVConfigUpdated, error)

	FilterChainAdded(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2ChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2ChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*USDCTokenPoolCCTPV2ChainAdded, error)

	FilterChainConfigured(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2ChainConfiguredIterator, error)

	WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2ChainConfigured) (event.Subscription, error)

	ParseChainConfigured(log types.Log) (*USDCTokenPoolCCTPV2ChainConfigured, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2ChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2ChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*USDCTokenPoolCCTPV2ChainRemoved, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2ConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2ConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*USDCTokenPoolCCTPV2ConfigChanged, error)

	FilterConfigSet(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2ConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2ConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*USDCTokenPoolCCTPV2ConfigSet, error)

	FilterCustomFinalityMinimumBlockConfirmationUpdated(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2CustomFinalityMinimumBlockConfirmationUpdatedIterator, error)

	WatchCustomFinalityMinimumBlockConfirmationUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2CustomFinalityMinimumBlockConfirmationUpdated) (event.Subscription, error)

	ParseCustomFinalityMinimumBlockConfirmationUpdated(log types.Log) (*USDCTokenPoolCCTPV2CustomFinalityMinimumBlockConfirmationUpdated, error)

	FilterCustomFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2CustomFinalityOutboundRateLimitConsumedIterator, error)

	WatchCustomFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2CustomFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomFinalityOutboundRateLimitConsumed(log types.Log) (*USDCTokenPoolCCTPV2CustomFinalityOutboundRateLimitConsumed, error)

	FilterCustomFinalityTransferInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2CustomFinalityTransferInboundRateLimitConsumedIterator, error)

	WatchCustomFinalityTransferInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2CustomFinalityTransferInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomFinalityTransferInboundRateLimitConsumed(log types.Log) (*USDCTokenPoolCCTPV2CustomFinalityTransferInboundRateLimitConsumed, error)

	FilterDomainsSet(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2DomainsSetIterator, error)

	WatchDomainsSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2DomainsSet) (event.Subscription, error)

	ParseDomainsSet(log types.Log) (*USDCTokenPoolCCTPV2DomainsSet, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2DynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2DynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*USDCTokenPoolCCTPV2DynamicConfigSet, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2InboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2InboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*USDCTokenPoolCCTPV2InboundRateLimitConsumed, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2LockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2LockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*USDCTokenPoolCCTPV2LockedOrBurned, error)

	FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2OutboundRateLimitConsumedIterator, error)

	WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2OutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseOutboundRateLimitConsumed(log types.Log) (*USDCTokenPoolCCTPV2OutboundRateLimitConsumed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*USDCTokenPoolCCTPV2OwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*USDCTokenPoolCCTPV2OwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*USDCTokenPoolCCTPV2OwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*USDCTokenPoolCCTPV2OwnershipTransferred, error)

	FilterPoolFeeWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*USDCTokenPoolCCTPV2PoolFeeWithdrawnIterator, error)

	WatchPoolFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2PoolFeeWithdrawn, recipient []common.Address) (event.Subscription, error)

	ParsePoolFeeWithdrawn(log types.Log) (*USDCTokenPoolCCTPV2PoolFeeWithdrawn, error)

	FilterRateLimitAdminSet(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2RateLimitAdminSetIterator, error)

	WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2RateLimitAdminSet) (event.Subscription, error)

	ParseRateLimitAdminSet(log types.Log) (*USDCTokenPoolCCTPV2RateLimitAdminSet, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2ReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2ReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*USDCTokenPoolCCTPV2ReleasedOrMinted, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2RemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2RemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*USDCTokenPoolCCTPV2RemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2RemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2RemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*USDCTokenPoolCCTPV2RemotePoolRemoved, error)

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*USDCTokenPoolCCTPV2TokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2TokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*USDCTokenPoolCCTPV2TokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*USDCTokenPoolCCTPV2TokenTransferFeeConfigUpdatedIterator, error)

	WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2TokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigUpdated(log types.Log) (*USDCTokenPoolCCTPV2TokenTransferFeeConfigUpdated, error)

	Address() common.Address
}
