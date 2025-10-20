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
	DestGasOverhead   uint32
	DestBytesOverhead uint32
	FeeUSDCents       uint32
	IsEnabled         bool
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

var USDCTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"internalType\":\"contractITokenMessenger\"},{\"name\":\"cctpMessageTransmitterProxy\",\"type\":\"address\",\"internalType\":\"contractCCTPMessageTransmitterProxy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"supportedUSDCVersion\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"structAuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyFinalityConfigUpdates\",\"inputs\":[{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"destToUseDefaultFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAccumulatedFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDomain\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structUSDCTokenPool.Domain\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useLegacySourcePoolDataFormat\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enumIPoolV2.CCVDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_localDomainIdentifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_messageTransmitterProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractCCTPMessageTransmitterProxy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_supportedUSDCVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_tokenMessenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITokenMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustomFinalityRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDomains\",\"inputs\":[{\"name\":\"domains\",\"type\":\"tuple[]\",\"internalType\":\"structUSDCTokenPool.DomainUpdate[]\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useLegacySourcePoolDataFormat\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityTransferInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DomainsSet\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structUSDCTokenPool.DomainUpdate[]\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useLegacySourcePoolDataFormat\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalityConfigUpdated\",\"inputs\":[{\"name\":\"finalityConfig\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeeWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidDestinationDomain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidDomain\",\"inputs\":[{\"name\":\"domain\",\"type\":\"tuple\",\"internalType\":\"structUSDCTokenPool.DomainUpdate\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useLegacySourcePoolDataFormat\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"type\":\"error\",\"name\":\"InvalidFinality\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidFinalityConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMessageLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidNonce\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"got\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidPreviousPool\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourceDomain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenMessengerVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidTransmitterInProxy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenTransferFeeConfigNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnknownDomain\",\"inputs\":[{\"name\":\"domain\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"UnlockingUSDCFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610180806040523461068d5761782c803803809161001d8285610a56565b8339810160e08282031261068d5781516001600160a01b0381169283820361068d576020810151936001600160a01b0385169081860361068d576040830151956001600160a01b0387169586880361068d5760608501516001600160401b03811161068d5785019080601f8301121561068d578151916001600160401b038311610692578260051b9060208201936100b86040519586610a56565b845260208085019282010192831161068d57602001905b828210610a3e575050506100e560808601610a79565b946100fe60c06100f760a08401610a79565b9201610a8d565b95602099604051996101108c8c610a56565b60008b5260003681373315610a2d57600180546001600160a01b0319163317905580158015610a1c575b8015610a0b575b6109fa576004928c9260805260c0526040519283809263313ce56760e01b82525afa80916000916109c3575b509061099f575b50600660a052600480546001600160a01b0319166001600160a01b03929092169190911790558051151560e0819052610871575b50604051946101b78887610a56565b60008652600036813760408051979088016001600160401b03811189821017610692576040528752858888015260005b865181101561024e576001906001600160a01b03610205828a610a9e565b51168a61021182610b1f565b61021e575b5050016101e7565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a1388a610216565b5087955086519360005b85518110156102c9576001600160a01b036102738288610a9e565b51169081156102b8577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef89836102aa600195610c3b565b50604051908152a101610258565b6342bcdf7f60e11b60005260046000fd5b50869394508561010052841561086057604051632c12192160e01b81528481600481895afa9081156106f85760009161082b575b5060405163054fd4d560e41b81526001600160a01b039190911691908581600481865afa9081156106f8576000916107f6575b5063ffffffff8061010051169116908082036107df575050604051639cdbb18160e01b815285816004818a5afa9081156106f8576000916107aa575b5063ffffffff80610100511691169080820361079357505084600491604051928380926367e0ed8360e11b82525afa80156106f857829160009161074a575b506001600160a01b03160361073957600492849261012052610140526040519283809263234d8e3d60e21b82525afa9081156106f857600091610704575b506101605260805161012051604051636eb1769f60e11b81523060048201526001600160a01b03918216602482018190529492909116908381604481855afa9081156106f8576000916106cb575b5060001981018091116106b557604051908482019563095ea7b360e01b8752602483015260448201526044815261046f606482610a56565b6000806040968751936104828986610a56565b8785527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656488860152519082865af13d156106a8573d906001600160401b0382116106925786516104ef9490926104e1601f8201601f1916890185610a56565b83523d60008885013e610d08565b805180610614575b847f2e902d38f15b233cbb63711add0fca4545334d3a169d60c0a616494d7eea954485858351908152a151616a539081610dd9823960805181818161055e01528181610759015281816107b701528181610dad01528181611df10152818161252b01528181613984015281816161cc015261653f015260a0518181816108b701528181615188015281816152180152615aa7015260c051818181612ab501528181614af201528181614c96015281816156e301526157ef015260e05181818161120401528181612bbb01526164690152610100518181816107080152614eff0152610120518181816114270152611e59015261014051818181610d150152611c8801526101605181818161161501528181611f310152614f6c0152f35b8184918101031261068d5782015180159081150361068d576106375783806104f7565b50608491519062461bcd60e51b82526004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b916104ef92606091610d08565b634e487b7160e01b600052601160045260246000fd5b90508381813d83116106f1575b6106e28183610a56565b8101031261068d575185610437565b503d6106d8565b6040513d6000823e3d90fd5b90508181813d8311610732575b61071b8183610a56565b8101031261068d5761072c90610a8d565b836103e9565b503d610711565b632a32133b60e11b60005260046000fd5b9091508581813d831161078c575b6107628183610a56565b810103126107885751906001600160a01b038216820361078557508190876103ab565b80fd5b5080fd5b503d610758565b633785f8f160e01b60005260045260245260446000fd5b90508581813d83116107d8575b6107c18183610a56565b8101031261068d576107d290610a8d565b8761036c565b503d6107b7565b63960693cd60e01b60005260045260245260446000fd5b90508581813d8311610824575b61080d8183610a56565b8101031261068d5761081e90610a8d565b87610330565b503d610803565b90508481813d8311610859575b6108428183610a56565b8101031261068d5761085390610a79565b866102fd565b503d610838565b6306b7c75960e31b60005260046000fd5b60405192969295919490936108868988610a56565b60008752600036813760e0511561098e5760005b8751811015610901576001906001600160a01b036108b8828b610a9e565b51168b6108c482610c74565b6108d1575b50500161089a565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a1388b6108c9565b509193955091939560005b8651811015610980576001906001600160a01b0361092a828a610a9e565b5116801561097a578a61093c82610bfc565b61094a575b50505b0161090c565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a1388a610941565b50610944565b5091939590929450386101a8565b6335f4a7b360e01b60005260046000fd5b60ff1660068114610174576332ad3e0760e11b600052600660045260245260446000fd5b8b81813d83116109f3575b6109d88183610a56565b8101031261078857519060ff8216820361078557503861016d565b503d6109ce565b630a64406560e11b60005260046000fd5b506001600160a01b03831615610141565b506001600160a01b0384161561013a565b639b15e16f60e01b60005260046000fd5b60208091610a4b84610a79565b8152019101906100cf565b601f909101601f19168101906001600160401b0382119082101761069257604052565b51906001600160a01b038216820361068d57565b519063ffffffff8216820361068d57565b8051821015610ab25760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b8054821015610ab25760005260206000200190600090565b80548015610b09576000190190610af78282610ac8565b8154906000199060031b1b1916905555565b634e487b7160e01b600052603160045260246000fd5b6000818152601160205260409020548015610bca5760001981018181116106b5576010546000198101919082116106b557808203610b79575b505050610b656010610ae0565b600052601160205260006040812055600190565b610bb2610b8a610b9b936010610ac8565b90549060031b1c9283926010610ac8565b819391549060031b91821b91600019901b19161790565b90556000526011602052604060002055388080610b58565b5050600090565b805490680100000000000000008210156106925781610b9b916001610bf894018155610ac8565b9055565b80600052600360205260406000205415600014610c3557610c1e816002610bd1565b600254906000526003602052604060002055600190565b50600090565b80600052601160205260406000205415600014610c3557610c5d816010610bd1565b601054906000526011602052604060002055600190565b6000818152600360205260409020548015610bca5760001981018181116106b5576002546000198101919082116106b557818103610cce575b505050610cba6002610ae0565b600052600360205260006040812055600190565b610cf0610cdf610b9b936002610ac8565b90549060031b1c9283926002610ac8565b90556000526003602052604060002055388080610cad565b91929015610d6a5750815115610d1c575090565b3b15610d255790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b825190915015610d7d5750805190602001fd5b6040519062461bcd60e51b8252602060048301528181519182602483015260005b838110610dc05750508160006044809484010152601f80199101168101030190fd5b60208282018101516044878401015285935001610d9e56fe6080604052600436101561001257600080fd5b60003560e01c806301ffc9a714610317578063164e68de14610312578063181f5a771461030d578063212a052e1461030857806321df0da714610303578063240028e8146102fe5780632451a627146102f957806324f65ee7146102f45780632a10097b146102ef5780632c286daf146102ea57806337b19247146102e557806339077537146102e0578063489a68f2146102db5780634ac8bd5f146102d65780634c5ef0ed146102d157806354c8a4f3146102cc5780635df45a37146102c75780636155cda0146102c257806362ddd3c4146102bd578063698c2c66146102b85780636b716b0d146102b35780636d3d1a58146102ae5780637437ff9f146102a957806379ba5097146102a45780637d54534e1461029f578063804ba5a91461029a5780638926f54f1461029557806389720a62146102905780638da5cb5b1461028b57806391a2749a14610286578063962d40201461028157806398db96431461027c5780639a4575b914610277578063a42a7b8b14610272578063a7cd63b71461026d578063acfecf9114610268578063af58d59f14610263578063b1c71c651461025e578063b794658014610259578063c4bffe2b14610254578063c75eea9c1461024f578063cf7401f31461024a578063d966866b14610245578063dc0bd97114610240578063dfadfa351461023b578063e0351e1314610236578063e8a1da17146102315763f2fde38b1461022c57600080fd5b612fff565b612be0565b612ba3565b612ad9565b612a88565b6128da565b6127fc565b6126da565b612665565b6125ea565b61247e565b6123b6565b61228f565b612235565b612151565b611ce7565b611c5b565b611aea565b611a25565b611968565b6118f8565b6118b9565b611809565b611778565b6116ad565b61166d565b611639565b6115f8565b611507565b611484565b6113fa565b6113d7565b6111d2565b611137565b610f31565b610e65565b610c30565b610b5c565b6109f4565b61090c565b61089d565b610837565b61077d565b61072c565b6106eb565b61068c565b61047b565b34610458576020600319360112610458576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361045857807ff208a58f000000000000000000000000000000000000000000000000000000006103ac921490811561042e575b8115610404575b81156103da575b81156103b0575b5060405190151581529081906020820190565b0390f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438610399565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150610392565b7f479eecb2000000000000000000000000000000000000000000000000000000008114915061038b565b7faff2afbf0000000000000000000000000000000000000000000000000000000081149150610384565b600080fd5b73ffffffffffffffffffffffffffffffffffffffff81160361045857565b34610458576020600319360112610458576004356104988161045d565b6104a0614652565b6104a861393b565b90816104b057005b73ffffffffffffffffffffffffffffffffffffffff7f738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b859599916106096040516105ee60208201917fa9059cbb00000000000000000000000000000000000000000000000000000000835261055b8161054d8a88602484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b03601f19810183528261106a565b857f000000000000000000000000000000000000000000000000000000000000000016600080604095865194610591888761106a565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020870152519082855af13d1561062f573d916105d2836110c9565b926105df8751948561106a565b83523d6000602085013e616982565b80518061060e575b5050519283921694829190602083019252565b0390a2005b81602080610623936106289501019101613538565b615d15565b38806105f6565b606091616982565b919082519283825260005b848110610663575050601f19601f8460006020809697860101520116010190565b80602080928401015182828601015201610642565b906020610689928181520190610637565b90565b34610458576000600319360112610458576103ac60408051906106af818361106a565b601782527f55534443546f6b656e506f6f6c20312e362e332d646576000000000000000000602083015251918291602083526020830190610637565b3461045857600060031936011261045857602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461045857600060031936011261045857602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346104585760206003193601126104585760206107dd60043561079f8161045d565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691161490565b6040519015158152f35b602060408183019282815284518094520192019060005b81811061080b5750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016107fe565b34610458576000600319360112610458576040516010548082526020820190601060005260206000209060005b818110610887576103ac8561087b8187038261106a565b604051918291826107e7565b8254845260209093019260019283019201610864565b3461045857600060031936011261045857602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b9181601f840112156104585782359167ffffffffffffffff8311610458576020808501948460051b01011161045857565b346104585760406003193601126104585760043567ffffffffffffffff811161045857366023820112156104585780600401359067ffffffffffffffff82116104585736602460a0840283010111610458576024359067ffffffffffffffff82116104585761098e9261098560249336906004016108db565b939092016130d5565b005b6004359061ffff8216820361045857565b6024359061ffff8216820361045857565b6064359061ffff8216820361045857565b9181601f840112156104585782359167ffffffffffffffff83116104585760208085019460e0850201011161045857565b3461045857606060031936011261045857610a0d610990565b610a156109a1565b60443567ffffffffffffffff811161045857610a359036906004016109c3565b9290610a3f614652565b61ffff8316612710811015610ad457507f52aa194b292c8bfb5aaca8ee2000a965c3a051b306ff841873b16147526a39ba93610ab591600b5463ffff00008660101b16907fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000061ffff871691161717600b556146b7565b6040805161ffff9283168152929091166020830152819081015b0390a1005b7f95f3517a0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b67ffffffffffffffff81160361045857565b3590610b1e82610b01565b565b908160a09103126104585790565b9181601f840112156104585782359167ffffffffffffffff8311610458576020838186019501011161045857565b346104585760a060031936011261045857610b7860043561045d565b602435610b8481610b01565b60443567ffffffffffffffff811161045857610ba4903690600401610b20565b50610bad6109b2565b5060843567ffffffffffffffff8111610458576103ac91610bd5610bdc923690600401610b2e565b50506133dd565b60405191829182919091606080608083019463ffffffff815116845263ffffffff602082015116602085015263ffffffff604082015116604085015201511515910152565b90816101009103126104585790565b346104585760206003193601126104585760043567ffffffffffffffff811161045857610c61903690600401610c21565b610c6961345b565b506060810135610c798183614a23565b610cfa6020610c96610c8e60e086018661346e565b8101906134bf565b610cbf610cb8610cb3610cac60c089018961346e565b36916110e5565b614e3f565b8251614eec565b8181519101519060405193849283927f57ecfd280000000000000000000000000000000000000000000000000000000084526004840161354d565b0381600073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af1908115610e6057600091610e31575b5015610e0757817ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc067ffffffffffffffff610d926040610d8b60206103ac9801613232565b940161357e565b6040805173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000081168252336020830152929092169082015260608101859052921691608090a2610df461108d565b8190526040519081529081906020820190565b7fbf969f220000000000000000000000000000000000000000000000000000000060005260046000fd5b610e53915060203d602011610e59575b610e4b818361106a565b810190613538565b38610d46565b503d610e41565b613572565b346104585760406003193601126104585760043567ffffffffffffffff811161045857610e996103ac913690600401610c21565b610ea16109a1565b906000604051610eb081610fbd565b52610ee2610eda6060830135610ed4610ecf610cac60c087018761346e565b6150df565b90615215565b928383614c4b565b7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc067ffffffffffffffff610d9260206040850194610f20863561045d565b013593610f2c85610b01565b61357e565b346104585760206003193601126104585760043567ffffffffffffffff8111610458573660238201121561045857806004013567ffffffffffffffff81116104585736602460c083028401011161045857602461098e9201613588565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff821117610fd957604052565b610f8e565b6040810190811067ffffffffffffffff821117610fd957604052565b6060810190811067ffffffffffffffff821117610fd957604052565b60a0810190811067ffffffffffffffff821117610fd957604052565b6080810190811067ffffffffffffffff821117610fd957604052565b60c0810190811067ffffffffffffffff821117610fd957604052565b90601f601f19910116810190811067ffffffffffffffff821117610fd957604052565b60405190610b1e60208361106a565b60405190610b1e60408361106a565b60405190610b1e60808361106a565b60405190610b1e60a08361106a565b67ffffffffffffffff8111610fd957601f01601f191660200190565b9291926110f1826110c9565b916110ff604051938461106a565b829481845281830111610458578281602093846000960137010152565b9080601f8301121561045857816020610689933591016110e5565b346104585760406003193601126104585760043561115481610b01565b60243567ffffffffffffffff81116104585760209161117a6107dd92369060040161111c565b906138ef565b60406003198201126104585760043567ffffffffffffffff811161045857816111ab916004016108db565b929092916024359067ffffffffffffffff8211610458576111ce916004016108db565b9091565b34610458576111fa6112026111e636611180565b94916111f3939193614652565b36916119b4565b9236916119b4565b7f0000000000000000000000000000000000000000000000000000000000000000156113ad5760005b82518110156112f0578061125e61124460019386613ccf565b5173ffffffffffffffffffffffffffffffffffffffff1690565b61129a61129573ffffffffffffffffffffffffffffffffffffffff83165b73ffffffffffffffffffffffffffffffffffffffff1690565b6165d6565b6112a6575b500161122b565b60405173ffffffffffffffffffffffffffffffffffffffff9190911681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756690602090a13861129f565b5060005b815181101561098e578061130d61124460019385613ccf565b73ffffffffffffffffffffffffffffffffffffffff8116156113a75761135061134b73ffffffffffffffffffffffffffffffffffffffff831661127c565b616809565b61135d575b505b016112f4565b60405173ffffffffffffffffffffffffffffffffffffffff9190911681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d890602090a183611355565b50611357565b7f35f4a7b30000000000000000000000000000000000000000000000000000000060005260046000fd5b346104585760006003193601126104585760206113f261393b565b604051908152f35b3461045857600060031936011261045857602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b9060406003198301126104585760043561146481610b01565b916024359067ffffffffffffffff8211610458576111ce91600401610b2e565b34610458576114923661144b565b61149d929192614652565b67ffffffffffffffff82166114bf816000526007602052604060002054151590565b156114da575061098e926114d49136916110e5565b90615323565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b34610458576040600319360112610458576004356115248161045d565b60243561152f614652565b73ffffffffffffffffffffffffffffffffffffffff82169182156115ce577f78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f0238927fffffffffffffffffffffffff0000000000000000000000000000000000000000600454161760045581600555610acf604051928392836020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461045857600060031936011261045857602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461045857600060031936011261045857602073ffffffffffffffffffffffffffffffffffffffff600a5416604051908152f35b34610458576000600319360112610458576004546005546040805173ffffffffffffffffffffffffffffffffffffffff9093168352602083019190915290f35b346104585760006003193601126104585760005473ffffffffffffffffffffffffffffffffffffffff8116330361174e577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b34610458576020600319360112610458577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174602073ffffffffffffffffffffffffffffffffffffffff6004356117cd8161045d565b6117d5614652565b16807fffffffffffffffffffffffff0000000000000000000000000000000000000000600a541617600a55604051908152a1005b346104585760206003193601126104585760043567ffffffffffffffff81116104585761183a9036906004016109c3565b73ffffffffffffffffffffffffffffffffffffffff600a541633141580611897575b6118695761098e916146b7565b7f8e4a23d6000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b5073ffffffffffffffffffffffffffffffffffffffff6001541633141561185c565b346104585760206003193601126104585760206107dd67ffffffffffffffff6004356118e481610b01565b166000526007602052604060002054151590565b346104585760c06003193601126104585761191460043561045d565b60243561192081610b01565b6044359061192c6109b2565b5060843567ffffffffffffffff81116104585761194d903690600401610b2e565b505060a4356002811015610458576103ac9261087b92613a02565b3461045857600060031936011261045857602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b67ffffffffffffffff8111610fd95760051b60200190565b9291906119c08161199c565b936119ce604051958661106a565b602085838152019160051b810192831161045857905b8282106119f057505050565b6020809183356119ff8161045d565b8152019101906119e4565b9080601f8301121561045857816020610689933591016119b4565b346104585760206003193601126104585760043567ffffffffffffffff8111610458576040600319823603011261045857604051611a6281610fde565b816004013567ffffffffffffffff811161045857611a869060043691850101611a0a565b8152602482013567ffffffffffffffff81116104585761098e926004611aaf9236920101611a0a565b6020820152613a78565b9181601f840112156104585782359167ffffffffffffffff8311610458576020808501946060850201011161045857565b346104585760606003193601126104585760043567ffffffffffffffff811161045857611b1b9036906004016108db565b9060243567ffffffffffffffff811161045857611b3c903690600401611ab9565b9060443567ffffffffffffffff811161045857611b5d903690600401611ab9565b611b7f61127c600a5473ffffffffffffffffffffffffffffffffffffffff1690565b33141580611c30575b61186957838614801590611c22575b611bf85760005b868110611ba757005b80611bf2611bc0611bbb6001948b8b6133cd565b613232565b611bcb838989613c07565b611bec611be4611bdc86898b613c07565b9236906127b3565b9136906127b3565b91615566565b01611b9e565b7f568efce20000000000000000000000000000000000000000000000000000000060005260046000fd5b5080861415611b97565b6000fd5b50611c5361127c60015473ffffffffffffffffffffffffffffffffffffffff1690565b331415611b88565b3461045857600060031936011261045857602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b610689916020611cc58351604084526040840190610637565b920151906020818403910152610637565b906020610689928181520190611cac565b346104585760206003193601126104585760043567ffffffffffffffff811161045857611d18903690600401610b20565b611d20613c17565b50611d2a81615696565b60208101611d5c611d57611d3d83613232565b67ffffffffffffffff166000526012602052604060002090565b613c30565b611d70611d6c6060830151151590565b1590565b612091576020611d80848061346e565b90500361204d576020810151801561202f57606090935b013590611dab604082015163ffffffff1690565b81516040517ff856ddb60000000000000000000000000000000000000000000000000000000081526004810185905263ffffffff909216602483015260448201959095527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16606482018190526084820195909552906020828060a481010381600073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af1908115610e60576103ac95611fa194611f9c94600094611fc7575b50916080917ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff611f5d95611f0c611ed58c613232565b6040805173ffffffffffffffffffffffffffffffffffffffff90971687523360208801528601929092529116929081906060820190565b0390a2611f2a611f1a61109c565b67ffffffffffffffff9095168552565b63ffffffff7f00000000000000000000000000000000000000000000000000000000000000001660208501520151151590565b15611fbe5760408051825167ffffffffffffffff166020808301919091529092015163ffffffff1690820152611f96816060810161054d565b92613232565b613e9c565b90611faa61109c565b918252602082015260405191829182611cd6565b611f969061594e565b611f5d93919450917ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff61201c60809560203d602011612028575b612014818361106a565b810190613cba565b96939550505091611e96565b503d61200a565b50606061204761203f858061346e565b810190613cab565b93611d97565b612057838061346e565b9061208d6040519283927fa3c8cf0900000000000000000000000000000000000000000000000000000000845260048401613c9a565b0390fd5b611c2c61209d83613232565b7fd201c48a0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b602081016020825282518091526040820191602060408360051b8301019401926000915b83831061210657505050505090565b9091929394602080612142837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc086600196030187528951610637565b970193019301919392906120f7565b346104585760206003193601126104585767ffffffffffffffff60043561217781610b01565b1660005260086020526121906005604060002001615deb565b805190601f196121b86121a28461199c565b936121b0604051958661106a565b80855261199c565b0160005b81811061222457505060005b815181101561221657806121fa6121f56121e460019486613ccf565b516000526009602052604060002090565b613d36565b6122048286613ccf565b5261220f8185613ccf565b50016121c8565b604051806103ac85826120d3565b8060606020809387010152016121bc565b34610458576000600319360112610458576040516002548082526020820190600260005260206000209060005b818110612279576103ac8561087b8187038261106a565b8254845260209093019260019283019201612262565b346104585761229d3661144b565b6122a8929192614652565b67ffffffffffffffff8216916122ce611d6c846000526007602052604060002054151590565b61237f57612311611d6c60056122f88467ffffffffffffffff166000526008602052604060002090565b016123043686896110e5565b6020815191012090616709565b61234857507f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76919261060960405192839283613c9a565b61208d84926040519384937f74f23c7c00000000000000000000000000000000000000000000000000000000855260048501613df6565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b346104585760206003193601126104585767ffffffffffffffff6004356123dc81610b01565b6123e4613e17565b501660005260086020526103ac6124096124046002604060002001613e42565b6159fe565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b929190612479602091604086526040860190611cac565b930152565b346104585760606003193601126104585760043567ffffffffffffffff8111610458576124af903690600401610b20565b6124b76109a1565b9060443567ffffffffffffffff8111610458576124d890369060040161111c565b506124e1613c17565b506124ec82826157a5565b606081013561ffff8193166125b9575b50611f9c6020612592920161251081613232565b6040805173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815233602082015290810186905267ffffffffffffffff91909116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1090606090a2613232565b61259a615aa0565b6125a261109c565b91825260208201526103ac60405192839283612462565b9091506127106125d261ffff600b5460101c1683613f81565b0481039081116125e55790611f9c6124fc565b613f52565b34610458576020600319360112610458576103ac61260d600435611f9c81610b01565b604051918291602083526020830190610637565b602060408183019282815284518094520192019060005b8181106126455750505090565b825167ffffffffffffffff16845260209384019390920191600101612638565b346104585760006003193601126104585761267e615da0565b805190601f196126906121a28461199c565b0136602084013760005b81518110156126cc578067ffffffffffffffff6126b960019385613ccf565b51166126c58286613ccf565b520161269a565b604051806103ac8582612621565b346104585760206003193601126104585767ffffffffffffffff60043561270081610b01565b612708613e17565b501660005260086020526103ac6124096124046040600020613e42565b8015150361045857565b6fffffffffffffffffffffffffffffffff81160361045857565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c6060910112610458576040519061278082610ffa565b8160843561278d81612725565b815260a43561279b8161272f565b6020820152604060c435916127af8361272f565b0152565b9190826060910312610458576040516127cb81610ffa565b604080829480356127db81612725565b845260208101356127eb8161272f565b60208501520135916127af8361272f565b346104585760e06003193601126104585760043561281981610b01565b60607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc3601126104585760405161284f81610ffa565b60243561285b81612725565b81526044356128698161272f565b602082015260643561287a8161272f565b604082015261288836612749565b9073ffffffffffffffffffffffffffffffffffffffff600a5416331415806128b8575b6118695761098e92615566565b5073ffffffffffffffffffffffffffffffffffffffff600154163314156128ab565b346104585760206003193601126104585760043567ffffffffffffffff81116104585761290b9036906004016108db565b612913614652565b60005b81811061291f57005b80837fece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d67ffffffffffffffff61295b611bbb6001968886613ebe565b612a7f61297661296c878a88613ebe565b6020810190613efe565b896129916129878a838b969b613ebe565b6040810190613efe565b6129c46129ba8c6129b26129a882888b989b613ebe565b6060810190613efe565b969095613ebe565b6080810190613efe565b9590946129da6129d536838f6119b4565b615adb565b6129e86129d53685856119b4565b6129f66129d53687876119b4565b612a046129d53689896119b4565b612a718c612a1c612a136110ab565b918436916119b4565b8152612a293686866119b4565b6020820152612a393688886119b4565b6040820152612a49368a8a6119b4565b6060820152612a6c8b67ffffffffffffffff16600052600e602052604060002090565b61405d565b604051998a99169b8961417d565b0390a201612916565b3461045857600060031936011261045857602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346104585760206003193601126104585767ffffffffffffffff600435612aff81610b01565b612b07613e17565b501660005260126020526103ac604060002060ff600260405192612b2a84611016565b8054845260018101546020850152015463ffffffff81166040840152818160201c161515606084015260281c16151560808201526040519182918291909160808060a0830194805184526020810151602085015263ffffffff604082015116604085015260608101511515606085015201511515910152565b346104585760006003193601126104585760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b3461045857612bee36611180565b919092612bf9614652565b6000915b808310612eab5750505060009163ffffffff4216925b828110612c1c57005b612c2f612c2a82858561428e565b61434d565b9060608201612c3e8151615bd4565b6080830193612c4d8551615bd4565b6040840190815151156115ce57612c87611d6c612c82612c75885167ffffffffffffffff1690565b67ffffffffffffffff1690565b616844565b612e6057612dc0612cc0612ca6879a999a5167ffffffffffffffff1690565b67ffffffffffffffff166000526008602052604060002090565b612d8389612d7d8751612d64612ce960408301516fffffffffffffffffffffffffffffffff1690565b91612d4b612d14612d0d60208401516fffffffffffffffffffffffffffffffff1690565b9251151590565b612d42612d1f6110ba565b6fffffffffffffffffffffffffffffffff851681529763ffffffff166020890152565b15156040870152565b6fffffffffffffffffffffffffffffffff166060850152565b6fffffffffffffffffffffffffffffffff166080830152565b826143e4565b612db589612dac8a51612d64612ce960408301516fffffffffffffffffffffffffffffffff1690565b600283016143e4565b6004845191016144f0565b602085019660005b88518051821015612e035790612dfd600192612df683612df08c5167ffffffffffffffff1690565b92613ccf565b5190615323565b01612dc8565b50509796509490612e577f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c29392612e446001975167ffffffffffffffff1690565b92519351905190604051948594856145bd565b0390a101612c13565b611c2c612e75865167ffffffffffffffff1690565b7f1d5ad3c50000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b909192612ebc611bbb8584866133cd565b94612ed3611d6c67ffffffffffffffff881661667e565b612fc757612f006005612efa8867ffffffffffffffff166000526008602052604060002090565b01615deb565b9360005b8551811015612f4c57600190612f456005612f338b67ffffffffffffffff166000526008602052604060002090565b01612f3e838a613ccf565b5190616709565b5001612f04565b509350937f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916612fb960019397612f9e612f998267ffffffffffffffff166000526008602052604060002090565b6141fd565b60405167ffffffffffffffff90911681529081906020820190565b0390a1019190939293612bfd565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff861660045260246000fd5b346104585760206003193601126104585773ffffffffffffffffffffffffffffffffffffffff6004356130318161045d565b613039614652565b163381146130ab57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b9291926130e0614652565b60005b8281106131635750505060005b8181106130fc57505050565b8067ffffffffffffffff613116611bbb60019486886133cd565b60006131368267ffffffffffffffff16600052600f602052604060002090565b55167f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8600080a2016130f0565b80613174611bbb600193868661321d565b7f56f77aeff2def50c8b5f5a0df3bab7183df09bf36c6feba496bb42551db77d7067ffffffffffffffff60206131ab85898961321d565b01926131d4846131cf8367ffffffffffffffff16600052600f602052604060002090565b613254565b6131e560405192839216948261336b565b0390a2016130e3565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b919081101561322d5760a0020190565b6131ee565b3561068981610b01565b63ffffffff81160361045857565b3561068981612725565b9060609061329781356132668161323c565b849063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b60208101356132a58161323c565b67ffffffff0000000084549160201b16807fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff83161785557fffffffffffffffffffffffffffffffffffffffff0000000000000000ffffffff6bffffffff000000000000000060408501356133188161323c565b60401b16921617178355013561332d81612725565b81547fffffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffff1690151560601b6cff00000000000000000000000016179055565b919091606080608083019463ffffffff81356133868161323c565b16845263ffffffff602082013561339c8161323c565b16602085015263ffffffff60408201356133b58161323c565b16604085015201356133c681612725565b1515910152565b919081101561322d5760051b0190565b67ffffffffffffffff90600060606040516133f781611032565b828152826020820152826040820152015216600052600f602052604060002060ff6040519161342583611032565b5463ffffffff8116835263ffffffff8160201c16602084015263ffffffff8160401c16604084015260601c161515606082015290565b6040519061346882610fbd565b60008252565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610458570180359067ffffffffffffffff82116104585760200191813603831361045857565b6020818303126104585780359067ffffffffffffffff8211610458570160408183031261045857604051916134f383610fde565b813567ffffffffffffffff8111610458578161351091840161111c565b8352602082013567ffffffffffffffff811161045857613530920161111c565b602082015290565b90816020910312610458575161068981612725565b909161356461068993604084526040840190610637565b916020818403910152610637565b6040513d6000823e3d90fd5b356106898161045d565b613590614652565b60005b8281106135d25750907fc97f93e817584952f1c1d633f93784b8430f0633d002f9dcc4de4fe2780424d0916135cd6040519283928361384f565b0390a1565b6135e56135e0828585613727565b613737565b8051158015613701575b61368857906136828261367d611d3d6060600196519361366e6020820151613665613621604085015163ffffffff1690565b61365d6136316080870151151590565b9161363f60a0880151151590565b946136486110ba565b9b8c5260208c015263ffffffff1660408b0152565b151588860152565b15156080870152565b015167ffffffffffffffff1690565b6137aa565b01613593565b604080517f19d7585700000000000000000000000000000000000000000000000000000000815282516004820152602083015160248201529082015163ffffffff166044820152606082015167ffffffffffffffff16606482015260808201511515608482015260a090910151151560a482015260c490fd5b5067ffffffffffffffff613720606083015167ffffffffffffffff1690565b16156135ef565b919081101561322d5760c0020190565b60c0813603126104585760a0604051916137508361104e565b8035835260208101356020840152604081013561376c8161323c565b6040840152606081013561377f81610b01565b6060840152608081013561379281612725565b608084015201356137a281612725565b60a082015290565b6002608091835181556020840151600182015501916137fe63ffffffff604083015116849063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b6060810151835492909101517fffffffffffffffffffffffffffffffffffffffffffffffffffff0000ffffffff90921690151560201b64ff00000000161790151560281b65ff000000000016179055565b602080825281018390526040019160005b81811061386d5750505090565b90919260c080600192863581526020870135602082015263ffffffff60408801356138978161323c565b16604082015267ffffffffffffffff60608801356138b481610b01565b16606082015260808701356138c881612725565b1515608082015260a08701356138dd81612725565b151560a0820152019401929101613860565b9067ffffffffffffffff61068992166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b90816020910312610458575190565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610e60576000916139b8575090565b610689915060203d6020116139da575b6139d2818361106a565b81019061392c565b503d6139c8565b67ffffffffffffffff61068991166000526007602052604060002054151590565b67ffffffffffffffff16600052600e6020526040600020916002811015613a4957600114613a3857816001610689930190615482565b816002600361068994019101615482565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b613a80614652565b60208101519160005b8351811015613b1f5780613aa261124460019387613ccf565b613ac9613ac473ffffffffffffffffffffffffffffffffffffffff831661127c565b6168f7565b613ad5575b5001613a89565b60405173ffffffffffffffffffffffffffffffffffffffff9190911681527fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758090602090a138613ace565b5091505160005b8151811015613c0357613b3c6112448284613ccf565b9073ffffffffffffffffffffffffffffffffffffffff821615613bd9577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef613bd083613ba8613ba361127c60019773ffffffffffffffffffffffffffffffffffffffff1690565b616879565b5060405173ffffffffffffffffffffffffffffffffffffffff90911681529081906020820190565b0390a101613b26565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050565b919081101561322d576060020190565b60405190613c2482610fde565b60606020838281520152565b90604051613c3d81611016565b608060ff600283958054855260018101546020860152015463ffffffff81166040850152818160201c161515606085015260281c161515910152565b601f8260209493601f19938186528686013760008582860101520116010190565b916020610689938181520191613c79565b90816020910312610458573590565b90816020910312610458575161068981610b01565b805182101561322d5760209160051b010190565b90600182811c92168015613d2c575b6020831014613cfd57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691613cf2565b9060405191826000825492613d4a84613ce3565b8084529360018116908115613db65750600114613d6f575b50610b1e9250038361106a565b90506000929192526020600020906000915b818310613d9a575050906020610b1e9282010138613d62565b6020919350806001915483858901015201910190918492613d81565b60209350610b1e9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138613d62565b60409067ffffffffffffffff61068995931681528160208201520191613c79565b60405190613e2482611016565b60006080838281528260208201528260408201528260608201520152565b90604051613e4f81611016565b60806fffffffffffffffffffffffffffffffff6001839560ff8154848116875263ffffffff81871c16602088015260a01c1615156040860152015481808216166060850152821c16910152565b67ffffffffffffffff1660005260086020526106896004604060002001613d36565b919081101561322d5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6181360301821215610458570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610458570180359067ffffffffffffffff821161045857602001918160051b3603831361045857565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818102929181159184041417156125e557565b91613fae918354906000199060031b92831b921b19161790565b9055565b818110613fbd575050565b60008155600101613fb2565b81519167ffffffffffffffff8311610fd957680100000000000000008311610fd9576020908254848455808510614040575b500190600052602060002060005b8381106140165750505050565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501614009565b614057908460005285846000209182019101613fb2565b38613ffb565b90805180519067ffffffffffffffff8211610fd957680100000000000000008211610fd957602090845483865580841061410f575b500183600052602060002060005b8381106140df57505050509060036060836140c56020610b1e96015160018601613fc9565b6140d6604082015160028601613fc9565b01519101613fc9565b6001906020614102855173ffffffffffffffffffffffffffffffffffffffff1690565b94019381840155016140a0565b614126908660005284846000209182019101613fb2565b38614092565b9160209082815201919060005b8181106141465750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff873561416f8161045d565b168152019401929101614139565b9694926141bb9461419f6141ad936106899b999560808c5260808c019161412c565b9189830360208b015261412c565b91868303604088015261412c565b92606081850391015261412c565b80549060008155816141d9575050565b6000526020600020908101905b8181106141f1575050565b600081556001016141e6565b6005610b1e91600081556000600182015560006002820155600060038201556004810161422a8154613ce3565b9081614239575b5050016141c9565b81601f600093116001146142515750555b3880614231565b8183526020832061426c91601f01861c810190600101613fb2565b808252602082209081548360011b906000198560031b1c19161790555561424a565b919081101561322d5760051b810135907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee181360301821215610458570190565b9080601f830112156104585781356142e58161199c565b926142f3604051948561106a565b81845260208085019260051b820101918383116104585760208201905b83821061431f57505050505090565b813567ffffffffffffffff8111610458576020916143428784809488010161111c565b815201910190614310565b61012081360312610458576040519061436582611016565b61436e81610b13565b8252602081013567ffffffffffffffff81116104585761439190369083016142ce565b602083015260408101359067ffffffffffffffff8211610458576143bb6143dc923690830161111c565b60408401526143cd36606083016127b3565b606084015260c03691016127b3565b608082015290565b8151815460208401516040850151608091821b73ffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffff000000000000000000000000000000000000000000909416939093179290921791151560a01b74ff000000000000000000000000000000000000000016919091178355606084015193810151901b7fffffffffffffffffffffffffffffffff0000000000000000000000000000000016921691909117600190910155565b9190601f81116144ba57505050565b610b1e926000526020600020906020601f840160051c830193106144e6575b601f0160051c0190613fb2565b90915081906144d9565b919091825167ffffffffffffffff8111610fd957614518816145128454613ce3565b846144ab565b6020601f8211600114614554578190613fae939495600092614549575b50506000198260011b9260031b1c19161790565b015190503880614535565b601f1982169061456984600052602060002090565b9160005b8181106145a55750958360019596971061458c575b505050811b019055565b015160001960f88460031b161c19169055388080614582565b9192602060018192868b01518155019401920161456d565b6146216145ec610b1e9597969467ffffffffffffffff60a0951684526101006020850152610100840190610637565b9660408301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b01906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b73ffffffffffffffffffffffffffffffffffffffff60015416330361467357565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b919081101561322d5760e0020190565b356106898161272f565b919060005b8181106146c95750509050565b6146d481838661469d565b6146dd81613232565b6146e9611d6c826139e1565b61237f578161477f6147f9926147856020600197960161471161470c36836127b3565b615bd4565b61477f6147328467ffffffffffffffff16600052600c602052604060002090565b9182546147526147498263ffffffff9060801c1690565b63ffffffff1690565b1590816149ec575b816149c4575b816149a9575b81614994575b5080614985575b614930575b36906127b3565b90615e36565b6147b4608084019161479a61470c36856127b3565b67ffffffffffffffff16600052600d602052604060002090565b9283546147cb6147498263ffffffff9060801c1690565b1590816148f9575b816148d1575b816148b6575b816148a1575b5080614892575b6147ff575b5036906127b3565b016146bc565b61480e60a061484b92016146ad565b84906fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b82547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178355386147f1565b5061489c8261324a565b6147ec565b6148b0915060a01c60ff161590565b386147e5565b6fffffffffffffffffffffffffffffffff81161591506147df565b90506fffffffffffffffffffffffffffffffff6148f18987015460801c90565b1615906147d9565b90506fffffffffffffffffffffffffffffffff614928898701546fffffffffffffffffffffffffffffffff1690565b1615906147d3565b61493f61480e604089016146ad565b82547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178355614778565b5061498f8161324a565b614773565b6149a3915060a01c60ff161590565b3861476c565b6fffffffffffffffffffffffffffffffff8116159150614766565b90506fffffffffffffffffffffffffffffffff6149e48c86015460801c90565b161590614760565b90506fffffffffffffffffffffffffffffffff614a1b8c8601546fffffffffffffffffffffffffffffffff1690565b16159061475a565b600060808201614a38611d6c61079f8361357e565b614bfb57506020820191614ad96020614a7e614a56612c7587613232565b60801b7fffffffffffffffffffffffffffffffff000000000000000000000000000000001690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116600482015291829081906024820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610e60578391614bdc575b50614bb457614b38614b3384613232565b6160db565b614b4183613232565b90614b5a611d6c60a083019361117a610cac868661346e565b614b74575050610b1e9291614b6f9150613232565b616173565b614b7e925061346e565b9061208d6040519283927f24eb47e500000000000000000000000000000000000000000000000000000000845260048401613c9a565b6004827f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b614bf5915060203d602011610e5957610e4b818361106a565b38614b22565b614c07614c489161357e565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835273ffffffffffffffffffffffffffffffffffffffff16600452602490565b90fd5b916080830190614c60611d6c61079f8461357e565b614df1576020840193614c7d6020614a7e614a56612c7589613232565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610e6057600091614dd2575b50614da857614cd8614b3386613232565b614ce185613232565b90614cfa611d6c60a083019361117a610cac868661346e565b614b7457505061ffff1615614d9b5767ffffffffffffffff90614d96614d64614d5e86611bbb614d4e61479a7f41a8aa8df7945f0fb8ac5f7d88279638d9dc2ef9a6bf4ec9a53b80681b34aff7999a613232565b89614d588861357e565b91616224565b9261357e565b946040519384931695836020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b0390a2565b50614b6f610b1e92613232565b7f53ad11d80000000000000000000000000000000000000000000000000000000060005260046000fd5b614deb915060203d602011610e5957610e4b818361106a565b38614cc7565b611c2c614dfd8361357e565b7f961c9a4f0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b60405190614e4c82610fde565b6000825260208201600081526020820151917fffffffff00000000000000000000000000000000000000000000000000000000602c602483015160c01c92015160e01c93167ff3567d18000000000000000000000000000000000000000000000000000000008103614ebf575083525290565b7fa176027f0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b908151607481106150b2575060048201517f000000000000000000000000000000000000000000000000000000000000000063ffffffff811663ffffffff8316036150795750506008820151916014600c82015191015192614f55602084015163ffffffff1690565b63ffffffff811663ffffffff8316036150405750507f000000000000000000000000000000000000000000000000000000000000000063ffffffff811663ffffffff8316036150075750505167ffffffffffffffff1667ffffffffffffffff811667ffffffffffffffff831603614fca575050565b7ff917ffea0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff9081166004521660245260446000fd5b7f77e480260000000000000000000000000000000000000000000000000000000060005263ffffffff9081166004521660245260446000fd5b7fe366a1170000000000000000000000000000000000000000000000000000000060005263ffffffff9081166004521660245260446000fd5b7f960693cd0000000000000000000000000000000000000000000000000000000060005263ffffffff9081166004521660245260446000fd5b7f758b22cc0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b805180156151845760200361514757615101602082518301016020830161392c565b9060ff8211615111575060ff1690565b61208d906040519182917f953576f700000000000000000000000000000000000000000000000000000000835260048301610678565b61208d906040519182917f953576f70000000000000000000000000000000000000000000000000000000083526020600484018181520190610637565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff82116125e557565b60ff16604d81116125e557600a0a90565b80156151dc576000190490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b81156151dc570490565b907f000000000000000000000000000000000000000000000000000000000000000060ff811660ff83168181146152fb57116152d05761525582826151aa565b91604d60ff84161180156152b7575b61527d57505090615277610689926151be565b90613f81565b7fa9cb113d0000000000000000000000000000000000000000000000000000000060005260ff908116600452166024525060445260646000fd5b506152c96152c4846151be565b6151cf565b8411615264565b6152da81836151aa565b91604d60ff84161161527d575050906152f5610689926151be565b9061520b565b5050505090565b60409067ffffffffffffffff61068994931681528160208201520190610637565b908051156115ce578051602082012067ffffffffffffffff8316928360005260086020526153588260056040600020016168ae565b156153ac5750816153a07f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9361539b614d96946000526009602052604060002090565b6144f0565b60405191829182610678565b905061208d6040519283927f393b8ad200000000000000000000000000000000000000000000000000000000845260048401615302565b906040519182815491828252602082019060005260206000209260005b818110615415575050610b1e9250038361106a565b845473ffffffffffffffffffffffffffffffffffffffff16835260019485019487945060209093019201615400565b919082018092116125e557565b9061545b8261199c565b615468604051918261106a565b828152601f19615478829461199c565b0190602036910137565b61548b906153e3565b91600554801515918261555b575b50506154a3575090565b6154ac906153e3565b8051806154b857505090565b6154c96154ce918495939451615444565b615451565b9060005b845181101561551957806155136154ee61124460019489613ccf565b6154f88387613ccf565b9073ffffffffffffffffffffffffffffffffffffffff169052565b016154d2565b509160005b8151811015615554578061554e61553a61124460019486613ccf565b6154f8615548848a51615444565b87613ccf565b0161551e565b5090925050565b101590503880615499565b67ffffffffffffffff166000818152600760205260409020549092919015615668579161566560e092615631856155bd7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b97615bd4565b8460005260086020526155d4816040600020615e36565b6155dd83615bd4565b8460005260086020526155f7836002604060002001615e36565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b600090608081016156ac611d6c61079f8361357e565b615798575060208101906156ca6020614a7e614a56612c7586613232565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610e60578491615779575b5061575157610b1e9291606082615737615732604061574c960161357e565b616467565b615743614b3384613232565b01359250613232565b6164e9565b6004837f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b615792915060203d602011610e5957610e4b818361106a565b38615713565b614c48614c07849261357e565b60808101906157b9611d6c61079f8461357e565b614df15760208101906157d66020614a7e614a56612c7586613232565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610e605760009161592f575b50614da8578061583761573260406060940161357e565b615843614b3384613232565b01359261ffff81169081151580615910575b156158ff57600b5461ffff169182116158c85750507f7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b291614d96614d64614d5e84611bbb614d4e6158ae67ffffffffffffffff98613232565b67ffffffffffffffff16600052600c602052604060002090565b7fe08f03ef0000000000000000000000000000000000000000000000000000000060005261ffff9081166004521660245260446000fd5b5050610b1e92915061574c90613232565b50615928615921600b5461ffff1690565b61ffff1690565b1515615855565b615948915060203d602011610e5957610e4b818361106a565b38615820565b7fffffffff00000000000000000000000000000000000000000000000000000000602082519201517fffffffffffffffff000000000000000000000000000000000000000000000000604051937ff3567d1800000000000000000000000000000000000000000000000000000000602086015260c01b16602484015260e01b16602c8201526010815261068960308261106a565b9060001982019182116125e557565b919082039182116125e557565b615a06613e17565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff82511690602083019163ffffffff83511642034281116125e557615a6a906fffffffffffffffffffffffffffffffff60808701511690613f81565b81018091116125e557615a906fffffffffffffffffffffffffffffffff929183926168e5565b161682524263ffffffff16905290565b60405160ff7f00000000000000000000000000000000000000000000000000000000000000001660208201526020815261068960408261106a565b80519060005b828110615aed57505050565b600181018082116125e5575b838110615b095750600101615ae1565b73ffffffffffffffffffffffffffffffffffffffff615b288385613ccf565b5116615b3a61127c6112448487613ccf565b14615b4757600101615af9565b611c2c615b576112448486613ccf565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b610b1e9092919260608101936fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b805115615c785760408101516fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff615c38615c2360208501516fffffffffffffffffffffffffffffffff1690565b6fffffffffffffffffffffffffffffffff1690565b911611615c425750565b61208d906040519182917f8020d12400000000000000000000000000000000000000000000000000000000835260048301615b99565b6fffffffffffffffffffffffffffffffff615ca660408301516fffffffffffffffffffffffffffffffff1690565b1615801590615ced575b615cb75750565b61208d906040519182917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301615b99565b50615d0e615c2360208301516fffffffffffffffffffffffffffffffff1690565b1515615cb0565b15615d1c57565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b604051906006548083528260208101600660005260206000209260005b818110615dd2575050610b1e9250038361106a565b8454835260019485019487945060209093019201615dbd565b906040519182815491828252602082019060005260206000209260005b818110615e1d575050610b1e9250038361106a565b8454835260019485019487945060209093019201615e08565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c199161600c6135cd928054615e7e615e786147498363ffffffff9060801c1690565b426159f1565b9081616018575b5050615fc66001615ea960208601516fffffffffffffffffffffffffffffffff1690565b92615f34615ef7615c236fffffffffffffffffffffffffffffffff615ede85546fffffffffffffffffffffffffffffffff1690565b166fffffffffffffffffffffffffffffffff88166168e5565b82906fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b615f87615f418751151590565b82547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178255565b019182906fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b604083015181546fffffffffffffffffffffffffffffffff1660809190911b7fffffffffffffffffffffffffffffffff0000000000000000000000000000000016179055565b60405191829182615b99565b615c23615ef7916fffffffffffffffffffffffffffffffff61608c616093958261608560018a0154928261607e616077616061876fffffffffffffffffffffffffffffffff1690565b996fffffffffffffffffffffffffffffffff1690565b9560801c90565b1690613f81565b9116615444565b91166168e5565b80547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff00000000000000000000000000000000161781553880615e85565b67ffffffffffffffff166160fc816000526007602052604060002054151590565b1561614657503360005260116020526040600020541561611857565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fa9902c7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b67ffffffffffffffff7f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c911691826000526008602052806161f4600260406000200173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391616224565b6040805173ffffffffffffffffffffffffffffffffffffffff909216825260208201929092529081908101614d96565b8054939290919060ff60a086901c1615801561645f575b6164585761625a6fffffffffffffffffffffffffffffffff8616615c23565b9060018401958654616294615e78614749616287615c23856fffffffffffffffffffffffffffffffff1690565b9460801c63ffffffff1690565b806163c4575b505083811061637957508282106162fa5750610b1e9394506162bf91615c23916159f1565b6fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b90616331611c2c9361632c61631d84616317615c238c5460801c90565b936159f1565b616326836159e2565b90615444565b61520b565b7fd0c8d23a0000000000000000000000000000000000000000000000000000000060005260045260245273ffffffffffffffffffffffffffffffffffffffff16604452606490565b7f1a76572a00000000000000000000000000000000000000000000000000000000600052600452602483905273ffffffffffffffffffffffffffffffffffffffff1660445260646000fd5b82859293951161642e576163de615c236163e59460801c90565b91856167ae565b84547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff000000000000000000000000000000001617855591388061629a565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050509050565b50811561623b565b7f000000000000000000000000000000000000000000000000000000000000000061648f5750565b73ffffffffffffffffffffffffffffffffffffffff16806000526003602052604060002054156164bc5750565b7fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b67ffffffffffffffff7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944911691826000526008602052806161f4604060002073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391616224565b805482101561322d5760005260206000200190600090565b805480156165a75760001901906165968282616567565b60001982549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600081815260036020526040902054908115616677576000198201908282116125e5576002549260001984019384116125e5578383600095616636950361663c575b505050616625600261657f565b600390600052602052604060002090565b55600190565b6166256166689161665e61665461666e956002616567565b90549060031b1c90565b9283916002616567565b90613f94565b55388080616618565b5050600090565b600081815260076020526040902054908115616677576000198201908282116125e5576006549260001984019384116125e557838360009561663695036166de575b5050506166cd600661657f565b600790600052602052604060002090565b6166cd616668916166f6616654616700956006616567565b9283916006616567565b553880806166c0565b60018101918060005282602052604060002054928315156000146167a55760001984018481116125e55783549360001985019485116125e55760009585836166369761675d950361676c575b50505061657f565b90600052602052604060002090565b61678c6166689161678361665461679c9588616567565b92839187616567565b8590600052602052604060002090565b55388080616755565b50505050600090565b926167b99192613f81565b81018091116125e557610689916168e5565b80549068010000000000000000821015610fd957816167f2916001613fae94018155616567565b81939154906000199060031b92831b921b19161790565b60008181526003602052604090205461683e576168278160026167cb565b600254906000526003602052604060002055600190565b50600090565b60008181526007602052604090205461683e576168628160066167cb565b600654906000526007602052604060002055600190565b60008181526011602052604090205461683e576168978160106167cb565b601054906000526011602052604060002055600190565b600082815260018201602052604090205461667757806168d0836001936167cb565b80549260005201602052604060002055600190565b90808210156168f2575090565b905090565b600081815260116020526040902054908115616677576000198201908282116125e5576010549260001984019384116125e55783836166369460009603616957575b505050616946601061657f565b601190600052602052604060002090565b6169466166689161696f616654616979956010616567565b9283916010616567565b55388080616939565b919290156169fd5750815115616996575090565b3b1561699f5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015616a105750805190602001fd5b61208d906040519182917f08c379a00000000000000000000000000000000000000000000000000000000083526004830161067856fea164736f6c634300081a000a",
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

func (_USDCTokenPool *USDCTokenPoolTransactor) ApplyFinalityConfigUpdates(opts *bind.TransactOpts, finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "applyFinalityConfigUpdates", finalityThreshold, customFinalityTransferFeeBps, rateLimitConfigArgs)
}

func (_USDCTokenPool *USDCTokenPoolSession) ApplyFinalityConfigUpdates(finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.ApplyFinalityConfigUpdates(&_USDCTokenPool.TransactOpts, finalityThreshold, customFinalityTransferFeeBps, rateLimitConfigArgs)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) ApplyFinalityConfigUpdates(finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.ApplyFinalityConfigUpdates(&_USDCTokenPool.TransactOpts, finalityThreshold, customFinalityTransferFeeBps, rateLimitConfigArgs)
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

func (_USDCTokenPool *USDCTokenPoolTransactor) SetCustomFinalityRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _USDCTokenPool.contract.Transact(opts, "setCustomFinalityRateLimitConfig", rateLimitConfigArgs)
}

func (_USDCTokenPool *USDCTokenPoolSession) SetCustomFinalityRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.SetCustomFinalityRateLimitConfig(&_USDCTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_USDCTokenPool *USDCTokenPoolTransactorSession) SetCustomFinalityRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _USDCTokenPool.Contract.SetCustomFinalityRateLimitConfig(&_USDCTokenPool.TransactOpts, rateLimitConfigArgs)
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

type USDCTokenPoolCustomFinalityOutboundRateLimitConsumedIterator struct {
	Event *USDCTokenPoolCustomFinalityOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCustomFinalityOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCustomFinalityOutboundRateLimitConsumed)
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
		it.Event = new(USDCTokenPoolCustomFinalityOutboundRateLimitConsumed)
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

func (it *USDCTokenPoolCustomFinalityOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCustomFinalityOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCustomFinalityOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterCustomFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCustomFinalityOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "CustomFinalityOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCustomFinalityOutboundRateLimitConsumedIterator{contract: _USDCTokenPool.contract, event: "CustomFinalityOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchCustomFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCustomFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "CustomFinalityOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCustomFinalityOutboundRateLimitConsumed)
				if err := _USDCTokenPool.contract.UnpackLog(event, "CustomFinalityOutboundRateLimitConsumed", log); err != nil {
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

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseCustomFinalityOutboundRateLimitConsumed(log types.Log) (*USDCTokenPoolCustomFinalityOutboundRateLimitConsumed, error) {
	event := new(USDCTokenPoolCustomFinalityOutboundRateLimitConsumed)
	if err := _USDCTokenPool.contract.UnpackLog(event, "CustomFinalityOutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator struct {
	Event *USDCTokenPoolCustomFinalityTransferInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
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
		it.Event = new(USDCTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
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

func (it *USDCTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCustomFinalityTransferInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterCustomFinalityTransferInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "CustomFinalityTransferInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator{contract: _USDCTokenPool.contract, event: "CustomFinalityTransferInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchCustomFinalityTransferInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCustomFinalityTransferInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "CustomFinalityTransferInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
				if err := _USDCTokenPool.contract.UnpackLog(event, "CustomFinalityTransferInboundRateLimitConsumed", log); err != nil {
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

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseCustomFinalityTransferInboundRateLimitConsumed(log types.Log) (*USDCTokenPoolCustomFinalityTransferInboundRateLimitConsumed, error) {
	event := new(USDCTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
	if err := _USDCTokenPool.contract.UnpackLog(event, "CustomFinalityTransferInboundRateLimitConsumed", log); err != nil {
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

type USDCTokenPoolFinalityConfigUpdatedIterator struct {
	Event *USDCTokenPoolFinalityConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolFinalityConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolFinalityConfigUpdated)
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
		it.Event = new(USDCTokenPoolFinalityConfigUpdated)
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

func (it *USDCTokenPoolFinalityConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolFinalityConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolFinalityConfigUpdated struct {
	FinalityConfig               uint16
	CustomFinalityTransferFeeBps uint16
	Raw                          types.Log
}

func (_USDCTokenPool *USDCTokenPoolFilterer) FilterFinalityConfigUpdated(opts *bind.FilterOpts) (*USDCTokenPoolFinalityConfigUpdatedIterator, error) {

	logs, sub, err := _USDCTokenPool.contract.FilterLogs(opts, "FinalityConfigUpdated")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolFinalityConfigUpdatedIterator{contract: _USDCTokenPool.contract, event: "FinalityConfigUpdated", logs: logs, sub: sub}, nil
}

func (_USDCTokenPool *USDCTokenPoolFilterer) WatchFinalityConfigUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolFinalityConfigUpdated) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPool.contract.WatchLogs(opts, "FinalityConfigUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolFinalityConfigUpdated)
				if err := _USDCTokenPool.contract.UnpackLog(event, "FinalityConfigUpdated", log); err != nil {
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

func (_USDCTokenPool *USDCTokenPoolFilterer) ParseFinalityConfigUpdated(log types.Log) (*USDCTokenPoolFinalityConfigUpdated, error) {
	event := new(USDCTokenPoolFinalityConfigUpdated)
	if err := _USDCTokenPool.contract.UnpackLog(event, "FinalityConfigUpdated", log); err != nil {
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

type GetDynamicConfig struct {
	Router                           common.Address
	ThresholdAmountForAdditionalCCVs *big.Int
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

func (USDCTokenPoolCustomFinalityOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b2")
}

func (USDCTokenPoolCustomFinalityTransferInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x41a8aa8df7945f0fb8ac5f7d88279638d9dc2ef9a6bf4ec9a53b80681b34aff7")
}

func (USDCTokenPoolDomainsSet) Topic() common.Hash {
	return common.HexToHash("0xc97f93e817584952f1c1d633f93784b8430f0633d002f9dcc4de4fe2780424d0")
}

func (USDCTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f0238")
}

func (USDCTokenPoolFinalityConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x52aa194b292c8bfb5aaca8ee2000a965c3a051b306ff841873b16147526a39ba")
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
	return common.HexToHash("0x56f77aeff2def50c8b5f5a0df3bab7183df09bf36c6feba496bb42551db77d70")
}

func (_USDCTokenPool *USDCTokenPool) Address() common.Address {
	return _USDCTokenPool.address
}

type USDCTokenPoolInterface interface {
	GetAccumulatedFees(opts *bind.CallOpts) (*big.Int, error)

	GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetDomain(opts *bind.CallOpts, chainSelector uint64) (USDCTokenPoolDomain, error)

	GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

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

	ApplyFinalityConfigUpdates(opts *bind.TransactOpts, finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error)

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

	FilterCustomFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCustomFinalityOutboundRateLimitConsumedIterator, error)

	WatchCustomFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCustomFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomFinalityOutboundRateLimitConsumed(log types.Log) (*USDCTokenPoolCustomFinalityOutboundRateLimitConsumed, error)

	FilterCustomFinalityTransferInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator, error)

	WatchCustomFinalityTransferInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCustomFinalityTransferInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomFinalityTransferInboundRateLimitConsumed(log types.Log) (*USDCTokenPoolCustomFinalityTransferInboundRateLimitConsumed, error)

	FilterDomainsSet(opts *bind.FilterOpts) (*USDCTokenPoolDomainsSetIterator, error)

	WatchDomainsSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolDomainsSet) (event.Subscription, error)

	ParseDomainsSet(log types.Log) (*USDCTokenPoolDomainsSet, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*USDCTokenPoolDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*USDCTokenPoolDynamicConfigSet, error)

	FilterFinalityConfigUpdated(opts *bind.FilterOpts) (*USDCTokenPoolFinalityConfigUpdatedIterator, error)

	WatchFinalityConfigUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolFinalityConfigUpdated) (event.Subscription, error)

	ParseFinalityConfigUpdated(log types.Log) (*USDCTokenPoolFinalityConfigUpdated, error)

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
