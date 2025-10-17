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

var USDCTokenPoolCCTPV2MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"internalType\":\"contractITokenMessenger\"},{\"name\":\"cctpMessageTransmitterProxy\",\"type\":\"address\",\"internalType\":\"contractCCTPMessageTransmitterProxy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"AUTHORIZED_CALLER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FINALITY_THRESHOLD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_FEE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_USDC_MESSAGE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RATE_LIMITER_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyFinalityConfigUpdates\",\"inputs\":[{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"destToUseDefaultFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beginDefaultAdminTransfer\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeDefaultAdminDelay\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"defaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelayIncreaseWait\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAccumulatedFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDomain\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structUSDCTokenPool.Domain\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useLegacySourcePoolDataFormat\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enumIPoolV2.CCVDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRateLimitAdminRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRateLimitAdminRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_localDomainIdentifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_messageTransmitterProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractCCTPMessageTransmitterProxy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_supportedUSDCVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_tokenMessenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITokenMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRateLimitAdminRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rollbackDefaultAdminDelay\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustomFinalityRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDomains\",\"inputs\":[{\"name\":\"domains\",\"type\":\"tuple[]\",\"internalType\":\"structUSDCTokenPool.DomainUpdate[]\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useLegacySourcePoolDataFormat\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityTransferInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeScheduled\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"effectSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferScheduled\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"acceptSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DomainsSet\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structUSDCTokenPool.DomainUpdate[]\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useLegacySourcePoolDataFormat\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalityConfigUpdated\",\"inputs\":[{\"name\":\"finalityConfig\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeeWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminRoleGranted\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminRoleRevoked\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlInvalidDefaultAdmin\",\"inputs\":[{\"name\":\"defaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidBurnToken\",\"inputs\":[{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"got\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDepositHash\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"got\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidDestinationDomain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidDomain\",\"inputs\":[{\"name\":\"domain\",\"type\":\"tuple\",\"internalType\":\"structUSDCTokenPool.DomainUpdate\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useLegacySourcePoolDataFormat\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"type\":\"error\",\"name\":\"InvalidExecutionFinalityThreshold\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidFinality\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidFinalityConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMessageLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMinFinalityThreshold\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidNonce\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"got\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidPreviousPool\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourceDomain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenMessengerVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidTransmitterInProxy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenTransferFeeConfigNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnknownDomain\",\"inputs\":[{\"name\":\"domain\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"UnlockingUSDCFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x610180806040523461057b57618454803803809161001d8285610937565b8339810160c08282031261057b5781516001600160a01b0381169283820361057b576020810151926001600160a01b0384169182850361057b5760408101516001600160a01b038116919082810361057b5760608201516001600160401b03811161057b5782019380601f8601121561057b578451946001600160401b038611610580578560051b9060208201966100b86040519889610937565b875260208088019282010192831161057b57602001905b82821061091f575050506100f160a06100ea6080850161095a565b930161095a565b92331561090957600180546001600160d01b031690556002546001600160a01b0381166108f8576001600160a01b03191633908117600255610132906109b4565b50801580156108e7575b80156108d6575b6108c55760049260209260805260c0526040519283809263313ce56760e01b82525afa8091600091610889575b5090610865575b50600660a052600580546001600160a01b0319166001600160a01b03929092169190911790558051151560e081905261073c575b50600161010052831561072b57604051632c12192160e01b8152602081600481885afa9081156105eb576000916106f1575b5060405163054fd4d560e41b81526001600160a01b03919091169190602081600481865afa9081156105eb576000916106d2575b5063ffffffff8061010051169116908082036106bb575050604051639cdbb18160e01b8152602081600481895afa9081156105eb5760009161069c575b5063ffffffff8061010051169116908082036106855750506020600491604051928380926367e0ed8360e11b82525afa80156105eb578291600091610637575b506001600160a01b0316036106265760049260209261012052610140526040519283809263234d8e3d60e21b82525afa9081156105eb576000916105f7575b506101605260805161012051604051636eb1769f60e11b81523060048201526001600160a01b0391821660248201819052939290911690602081604481855afa9081156105eb576000916105b9575b5060001981018091116105a35760405190602082019463095ea7b360e01b86526024830152604482015260448152610355606482610937565b6000806040958651936103688886610937565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082865af13d15610596573d906001600160401b0382116105805785516103d99490926103ca601f8201601f191660200185610937565b83523d6000602085013e610b9c565b8051806104ff575b837f2e902d38f15b233cbb63711add0fca4545334d3a169d60c0a616494d7eea95446020858351908152a1516177c79081610c6d82396080518181816108c701528181610ab501528181610b06015281816112480152818161231a01528181612a6d015281816142dc01528181616e6001526171dc015260a051818181610b9401528181615b6701528181615bf701526164f6015260c05181818161334d0152818161543d015281816155c7015281816160b901526161b8015260e0518181816116cf015281816134530152617120015261010051818181610a71015261584a01526101205181818161187a01526122d80152610140518181816111bd01526121c3015261016051818181611cf60152818161240901526158d00152f35b816020918101031261057b576020015180159081150361057b576105245738806103e1565b815162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b916103d992606091610b9c565b634e487b7160e01b600052601160045260246000fd5b90506020813d6020116105e3575b816105d460209383610937565b8101031261057b57513861031c565b3d91506105c7565b6040513d6000823e3d90fd5b610619915060203d60201161061f575b6106118183610937565b81019061096e565b386102cd565b503d610607565b632a32133b60e11b60005260046000fd5b9091506020813d60201161067d575b8161065360209383610937565b810103126106795751906001600160a01b0382168203610676575081903861028e565b80fd5b5080fd5b3d9150610646565b633785f8f160e01b60005260045260245260446000fd5b6106b5915060203d60201161061f576106118183610937565b3861024e565b63960693cd60e01b60005260045260245260446000fd5b6106eb915060203d60201161061f576106118183610937565b38610211565b90506020813d602011610723575b8161070c60209383610937565b8101031261057b5761071d9061095a565b386101dd565b3d91506106ff565b6306b7c75960e31b60005260046000fd5b919290602092604051926107508585610937565b60008452600036813760e051156108545760005b84518110156107cb576001906001600160a01b03610782828861098a565b51168761078e82610a5a565b61079b575b505001610764565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a13887610793565b50919490925060005b8351811015610848576001906001600160a01b036107f2828761098a565b51168015610842578661080482610b42565b610812575b50505b016107d4565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a13886610809565b5061080c565b509250929050386101ab565b6335f4a7b360e01b60005260046000fd5b60ff1660068114610177576332ad3e0760e11b600052600660045260245260446000fd5b6020813d6020116108bd575b816108a260209383610937565b8101031261067957519060ff82168203610676575038610170565b3d9150610895565b630a64406560e11b60005260046000fd5b506001600160a01b03831615610143565b506001600160a01b0384161561013c565b631fe1e13d60e11b60005260046000fd5b636116401160e11b600052600060045260246000fd5b6020809161092c8461095a565b8152019101906100cf565b601f909101601f19168101906001600160401b0382119082101761058057604052565b51906001600160a01b038216820361057b57565b9081602091031261057b575163ffffffff8116810361057b5790565b805182101561099e5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b6001600160a01b0381166000908152600080516020618434833981519152602052604090205460ff16610a3c576001600160a01b0316600081815260008051602061843483398151915260205260408120805460ff191660011790553391907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b50600090565b805482101561099e5760005260206000200190600090565b6000818152600460205260409020548015610b3b5760001981018181116105a3576003546000198101919082116105a357818103610aea575b5050506003548015610ad45760001901610aae816003610a42565b8154906000199060031b1b19169055600355600052600460205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b610b23610afb610b0c936003610a42565b90549060031b1c9283926003610a42565b819391549060031b91821b91600019901b19161790565b90556000526004602052604060002055388080610a93565b5050600090565b80600052600460205260406000205415600014610a3c576003546801000000000000000081101561058057610b83610b0c8260018594016003556003610a42565b9055600354906000526004602052604060002055600190565b91929015610bfe5750815115610bb0575090565b3b15610bb95790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b825190915015610c115750805190602001fd5b6040519062461bcd60e51b8252602060048301528181519182602483015260005b838110610c545750508160006044809484010152601f80199101168101030190fd5b60208282018101516044878401015285935001610c3256fe6080604052600436101561001257600080fd5b60003560e01c806301ffc9a71461042d578063022d63fb146104285780630aa6220b146104235780630bd7c46d1461041e578063164e68de14610419578063181f5a7714610414578063212a052e1461040f57806321df0da71461040a578063240028e814610405578063248a9ca31461040057806324f65ee7146103fb5780632a10097b146103f65780632c286daf146103f15780632f2ff15d146103ec57806336568abe146103e757806337b19247146103e257806339077537146103dd578063489a68f2146103d85780634ac8bd5f146103d35780634c5ef0ed146103ce57806354c8a4f3146103c95780635df45a37146103c45780636155cda0146103bf57806362ddd3c4146103ba578063634e93da146103b5578063649a5ec7146103b0578063698c2c66146103ab5780636b716b0d146103a65780637437ff9f146103a1578063791e5a101461039c578063804ba5a91461039757806384ef8ffc146103885780638926f54f1461039257806389720a621461038d5780638da5cb5b1461038857806391d1485414610383578063962d40201461037e57806398db9643146103795780639a4575b914610374578063a1eda53c1461036f578063a217fddf1461034c578063a42a7b8b1461036a578063a7cd63b714610365578063acfecf9114610360578063af58d59f1461035b578063b1c71c6514610356578063b794658014610351578063bc063e1a1461034c578063c4bffe2b14610347578063c75eea9c14610342578063c8c8fd191461033d578063cc8463c814610338578063cefc142914610333578063cf6eefb71461032e578063cf7401f314610329578063d547741f14610324578063d602b9fd1461031f578063d966866b1461031a578063da4b05e714610315578063da90a9f314610310578063dc0bd9711461030b578063dfadfa3514610306578063e0351e1314610301578063e58d80c7146102fc578063e8a1da17146102f75763f573388e146102f257600080fd5b613933565b6134e1565b613478565b61343b565b613371565b61332d565b6132ac565b61328f565b6130b2565b613037565b612fed565b612f3c565b612dac565b612cb3565b612c88565b612c6b565b612c20565b612bab565b6125df565b612b30565b6129d5565b61290d565b6127b7565b61275d565b612679565b61257c565b612222565b6121a3565b611fa6565b611f1e565b611df9565b611ea2565b611e20565b611d88565b611d4d565b611d1a565b611cd9565b611bd2565b611ada565b611989565b6118d7565b61185a565b61183f565b61166d565b6115d2565b6113cc565b611300565b6110e5565b611011565b610e86565b610e0d565b610ccf565b610be9565b610b7a565b610b45565b610ad9565b610a95565b610a54565b6109f5565b6107cf565b610711565b610633565b610615565b34610605576020600319360112610605576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361060557807ff208a58f000000000000000000000000000000000000000000000000000000006104c992149081156105db575b81156105b1575b8115610587575b811561055d575b81156104cd575b5060405190151581529081906020820190565b0390f35b7f3149878600000000000000000000000000000000000000000000000000000000811491508115610500575b50386104b6565b7f7965db0b00000000000000000000000000000000000000000000000000000000811491508115610533575b50386104f9565b7f01ffc9a7000000000000000000000000000000000000000000000000000000009150143861052c565b7f01ffc9a700000000000000000000000000000000000000000000000000000000811491506104af565b7f0e64dd2900000000000000000000000000000000000000000000000000000000811491506104a8565b7f479eecb200000000000000000000000000000000000000000000000000000000811491506104a1565b7faff2afbf000000000000000000000000000000000000000000000000000000008114915061049a565b600080fd5b600091031261060557565b34610605576000600319360112610605576020604051620697808152f35b346106055760006003193601126106055761064c614e09565b6002548060d01c8061066c575b600280546001600160a01b03169055005b005b4211156106d55765ffffffffffff6106ce9160a01c1679ffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffffffff00000000000000000000000000000000000000000000000000006001549260d01b16911617600155565b3880610659565b507f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5600080a16106ce565b6001600160a01b0381160361060557565b346106055760206003193601126106055760043561072e81610700565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff16156107a5576020816107927ff7af318a70f367e30346e2704f6ef646b378a7dcb49767beb98a1774cd11e38993614edc565b506001600160a01b0360405191168152a1005b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b34610605576020600319360112610605576004356107ec81610700565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff16156107a55761082b6142a0565b908161083357005b6001600160a01b037f738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b8595999161097260405161095760208201917fa9059cbb0000000000000000000000000000000000000000000000000000000083526108c4816108b68a8860248401602090939291936001600160a01b0360408201951681520152565b03601f198101835282611505565b857f0000000000000000000000000000000000000000000000000000000000000000166000806040958651946108fa8887611505565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020870152519082855af13d15610998573d9161093b83611564565b9261094887519485611505565b83523d6000602085013e6176f6565b805180610977575b5050519283921694829190602083019252565b0390a2005b8160208061098c936109919501019101613e00565b616a29565b388061095f565b6060916176f6565b919082519283825260005b8481106109cc575050601f19601f8460006020809697860101520116010190565b806020809284010151828286010152016109ab565b9060206109f29281815201906109a0565b90565b34610605576000600319360112610605576104c96040805190610a188183611505565b601d82527f55534443546f6b656e506f6f6c43435450563220312e372e302d6465760000006020830152519182916020835260208301906109a0565b3461060557600060031936011261060557602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346106055760006003193601126106055760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610605576020600319360112610605576020610b2c600435610afb81610700565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691161490565b6040519015158152f35b90816020910312610605573590565b34610605576020600319360112610605576020610b72600435600052600060205260016040600020015490565b604051908152f35b3461060557600060031936011261060557602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b9181601f840112156106055782359167ffffffffffffffff8311610605576020808501948460051b01011161060557565b346106055760406003193601126106055760043567ffffffffffffffff811161060557366023820112156106055780600401359067ffffffffffffffff82116106055736602460a0840283010111610605576024359067ffffffffffffffff82116106055761066a92610c626024933690600401610bb8565b9390920161396e565b6004359061ffff8216820361060557565b6024359061ffff8216820361060557565b6064359061ffff8216820361060557565b9181601f840112156106055782359167ffffffffffffffff83116106055760208085019460e0850201011161060557565b3461060557606060031936011261060557610ce8610c6b565b610cf0610c7c565b60443567ffffffffffffffff811161060557610d10903690600401610c9e565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb560205260409020549093919060ff16156107a55761ffff8316612710811015610de057507f52aa194b292c8bfb5aaca8ee2000a965c3a051b306ff841873b16147526a39ba93610dc191600b5463ffff00008660101b16907fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000061ffff871691161717600b55614fc8565b6040805161ffff9283168152929091166020830152819081015b0390a1005b7f95f3517a0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b3461060557604060031936011261060557600435602435610e2d81610700565b8115610e5c5781610e57610e5261066a94600052600060205260016040600020015490565b614e75565b614f54565b7f3fc3c27a0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461060557604060031936011261060557600435602435610ea681610700565b811580610f84575b610ebc575b61066a91615334565b60015465ffffffffffff60a082901c1692906001600160a01b031615801590610f70575b8015610f5e575b610f245761066a9250610f1d7fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff60015416600155565b9150610eb3565b7f19ca5ebb0000000000000000000000000000000000000000000000000000000060005265ffffffffffff831660045260246000fd5b6000fd5b504265ffffffffffff84161015610ee7565b5065ffffffffffff831615610ee0565b1590565b50610fa6610f9a6001600160a01b036002541690565b6001600160a01b031690565b6001600160a01b03821614610eae565b67ffffffffffffffff81160361060557565b3590610fd382610fb6565b565b908160a09103126106055790565b9181601f840112156106055782359167ffffffffffffffff8311610605576020838186019501011161060557565b346106055760a06003193601126106055761102d600435610700565b60243561103981610fb6565b60443567ffffffffffffffff811161060557611059903690600401610fd5565b50611062610c8d565b5060843567ffffffffffffffff8111610605576104c99161108a611091923690600401610fe3565b5050613ca5565b60405191829182919091606080608083019463ffffffff815116845263ffffffff602082015116602085015263ffffffff604082015116604085015201511515910152565b90816101009103126106055790565b346106055760206003193601126106055760043567ffffffffffffffff8111610605576111169036906004016110d6565b61111e613d23565b50606081013561112e818361537b565b6111af602061114b61114360e0860186613d36565b810190613d87565b61117461116d61116861116160c0890189613d36565b3691611580565b615756565b8251615837565b8181519101519060405193849283927f57ecfd2800000000000000000000000000000000000000000000000000000000845260048401613e15565b038160006001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165af19081156112fb576000916112cc575b50156112a257817ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc067ffffffffffffffff61123a604061123360206104c99801613afa565b9401613e46565b604080516001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081168252336020830152929092169082015260608101859052921691608090a261128f611528565b8190526040519081529081906020820190565b7fbf969f220000000000000000000000000000000000000000000000000000000060005260046000fd5b6112ee915060203d6020116112f4575b6112e68183611505565b810190613e00565b386111ee565b503d6112dc565b613e3a565b346106055760406003193601126106055760043567ffffffffffffffff8111610605576113346104c99136906004016110d6565b61133c610c7c565b90600060405161134b81611458565b5261137d611375606083013561136f61136a61116160c0870187613d36565b615abe565b90615bf4565b928383615589565b7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc067ffffffffffffffff61123a602060408501946113bb8635610700565b0135936113c785610fb6565b613e46565b346106055760206003193601126106055760043567ffffffffffffffff8111610605573660238201121561060557806004013567ffffffffffffffff81116106055736602460c083028401011161060557602461066a9201613e50565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff82111761147457604052565b611429565b6060810190811067ffffffffffffffff82111761147457604052565b60a0810190811067ffffffffffffffff82111761147457604052565b6080810190811067ffffffffffffffff82111761147457604052565b6040810190811067ffffffffffffffff82111761147457604052565b60c0810190811067ffffffffffffffff82111761147457604052565b90601f601f19910116810190811067ffffffffffffffff82111761147457604052565b60405190610fd3602083611505565b60405190610fd3604083611505565b60405190610fd3608083611505565b60405190610fd360a083611505565b67ffffffffffffffff811161147457601f01601f191660200190565b92919261158c82611564565b9161159a6040519384611505565b829481845281830111610605578281602093846000960137010152565b9080601f83011215610605578160206109f293359101611580565b34610605576040600319360112610605576004356115ef81610fb6565b60243567ffffffffffffffff811161060557602091611615610b2c9236906004016115b7565b906141e6565b60406003198201126106055760043567ffffffffffffffff8111610605578161164691600401610bb8565b929092916024359067ffffffffffffffff82116106055761166991600401610bb8565b9091565b346106055761167b3661161b565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb560205260409020549093929060ff16156107a5576116cd926116c591369161423b565b92369161423b565b7f0000000000000000000000000000000000000000000000000000000000000000156118155760005b825181101561177f578061171c61170f60019386614474565b516001600160a01b031690565b6117366117316001600160a01b038316610f9a565b617309565b611742575b50016116f6565b6040516001600160a01b039190911681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756690602090a13861173b565b5060005b815181101561066a578061179c61170f60019385614474565b6001600160a01b0381161561180f576117c56117c06001600160a01b038316610f9a565b6175b0565b6117d2575b505b01611783565b6040516001600160a01b039190911681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d890602090a1836117ca565b506117cc565b7f35f4a7b30000000000000000000000000000000000000000000000000000000060005260046000fd5b34610605576000600319360112610605576020610b726142a0565b346106055760006003193601126106055760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b906040600319830112610605576004356118b781610fb6565b916024359067ffffffffffffffff82116106055761166991600401610fe3565b34610605576118e53661189e565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205491929160ff16156107a55767ffffffffffffffff8216611941816000526008602052604060002054151590565b1561195c575061066a92611956913691611580565b90615d02565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b34610605576020600319360112610605576004356119a681610700565b6119ae614e09565b7f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed660206119eb6119dd426170d4565b6119e5614663565b90615dc2565b65ffffffffffff6001600160a01b03611a1a6001549065ffffffffffff6001600160a01b0383169260a01c1690565b9690501694857fffffffffffffffffffffffff00000000000000000000000000000000000000006001541617600155611a99837fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff79ffffffffffff00000000000000000000000000000000000000006001549260a01b16911617600155565b16611ab0575b65ffffffffffff60405191168152a2005b7f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109600080a1611a9f565b346106055760206003193601126106055760043565ffffffffffff81169081810361060557611b07614e09565b611b10426170d4565b9165ffffffffffff611b20614663565b1680821115611b9357507ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b9265ffffffffffff611b5f611b66936176ce565b1690615dc2565b90611b718282616824565b6040805165ffffffffffff928316815292909116602083015281908101610ddb565b0365ffffffffffff8111611bcd577ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b92611b669190615dc2565b614730565b3461060557604060031936011261060557600435611bef81610700565b6024353360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb56020526040902060ff905416156107a5576001600160a01b038216918215611caf577f78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f0238927fffffffffffffffffffffffff0000000000000000000000000000000000000000600554161760055581600655610ddb60405192839283602090939291936001600160a01b0360408201951681520152565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461060557600060031936011261060557602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461060557600060031936011261060557600554600654604080516001600160a01b039093168352602083019190915290f35b346106055760006003193601126106055760206040517f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af8152f35b346106055760206003193601126106055760043567ffffffffffffffff811161060557611db9903690600401610c9e565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff16156107a55761066a91614fc8565b346106055760006003193601126106055760206001600160a01b0360025416604051908152f35b34610605576020600319360112610605576020610b2c67ffffffffffffffff600435611e4b81610fb6565b166000526008602052604060002054151590565b602060408183019282815284518094520192019060005b818110611e835750505090565b82516001600160a01b0316845260209384019390920191600101611e76565b346106055760c060031936011261060557611ebe600435610700565b602435611eca81610fb6565b60443590611ed6610c8d565b5060843567ffffffffffffffff811161060557611ef7903690600401610fe3565b505060a4356002811015610605576104c992611f129261435a565b60405191829182611e5f565b3461060557604060031936011261060557602060ff611f69602435600435611f4582610700565b600052600084526040600020906001600160a01b0316600052602052604060002090565b54166040519015158152f35b9181601f840112156106055782359167ffffffffffffffff8311610605576020808501946060850201011161060557565b346106055760606003193601126106055760043567ffffffffffffffff811161060557611fd7903690600401610bb8565b9060243567ffffffffffffffff811161060557611ff8903690600401611f75565b9060443567ffffffffffffffff811161060557612019903690600401611f75565b7f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af600090815260205261208d610f80612086337f135a184f1ea19f46a2a5eda150dca2e4f81e2266dd73bb074c2f2a4e0c0032a05b906001600160a01b0316600052602052604060002090565b5460ff1690565b80612165575b6121375783861480159061212d575b6121035760005b8681106120b257005b806120fd6120cb6120c66001948b8b613c95565b613afa565b6120d68389896143d0565b6120f76120ef6120e786898b6143d0565b923690612ef3565b913690612ef3565b91615f49565b016120a9565b7f568efce20000000000000000000000000000000000000000000000000000000060005260046000fd5b50808614156120a2565b7f8e4a23d6000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b50600080805260205261219e610f80612086337fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb561206e565b612093565b346106055760006003193601126106055760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b6109f291602061220083516040845260408401906109a0565b9201519060208184039101526109a0565b9060206109f29281815201906121e7565b346106055760206003193601126106055760043567ffffffffffffffff811161060557612253903690600401610fd5565b61225b6143e0565b5061226581616079565b6020810161229761229261227883613afa565b67ffffffffffffffff166000526010602052604060002090565b6143f9565b916122a8610f806060850151151590565b61253a5760206122b88280613d36565b9050036124f657602083015180156124da57925b60606001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001692013592604082019261230f845163ffffffff1690565b926001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016938151833b15610605576040517fd04857b00000000000000000000000000000000000000000000000000000000081526004810189905263ffffffff929092166024830152604482018990526001600160a01b03861660648301526084820152600060a482018190526107d060c4830152909492859060e490829084905af180156112fb5761245e6124a29661243d7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1094866104c99c61249d9a67ffffffffffffffff976124bf575b506124337f0000000000000000000000000000000000000000000000000000000000000000955163ffffffff1690565b9251928d86616317565b612454612448611537565b63ffffffff9093168352565b60208201526163c2565b9661249561246b86613afa565b604080516001600160a01b0390971687523360208801528601929092529116929081906060820190565b0390a2613afa565b614641565b906124ab611537565b918252602082015260405191829182612211565b806124ce60006124d493611505565b8061060a565b38612403565b506124f06124e88280613d36565b810190610b36565b926122cc565b8061250091613d36565b906125366040519283927fa3c8cf0900000000000000000000000000000000000000000000000000000000845260048401614463565b0390fd5b610f5a61254683613afa565b7fd201c48a0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b34610605576000600319360112610605576002548060d01c90811515806125d5575b156125cb5760a01c65ffffffffffff165b6040805165ffffffffffff928316815292909116602083015290f35b50506000806125af565b504282101561259e565b3461060557600060031936011261060557602060405160008152f35b602081016020825282518091526040820191602060408360051b8301019401926000915b83831061262e57505050505090565b909192939460208061266a837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0866001960301875289516109a0565b9701930193019193929061261f565b346106055760206003193601126106055767ffffffffffffffff60043561269f81610fb6565b1660005260096020526126b8600560406000200161724f565b805190601f196126e06126ca84614223565b936126d86040519586611505565b808552614223565b0160005b81811061274c57505060005b815181101561273e578061272261271d61270c60019486614474565b51600052600a602052604060002090565b6144db565b61272c8286614474565b526127378185614474565b50016126f0565b604051806104c985826125fb565b8060606020809387010152016126e4565b34610605576000600319360112610605576040516003548082526020820190600360005260206000209060005b8181106127a1576104c985611f1281870382611505565b825484526020909301926001928301920161278a565b34610605576127c53661189e565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205491929160ff16156107a55767ffffffffffffffff821691612825610f80846000526008602052604060002054151590565b6128d657612868610f80600561284f8467ffffffffffffffff166000526009602052604060002090565b0161285b368689611580565b6020815191012090617435565b61289f57507f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76919261097260405192839283614463565b61253684926040519384937f74f23c7c0000000000000000000000000000000000000000000000000000000085526004850161459b565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b346106055760206003193601126106055767ffffffffffffffff60043561293381610fb6565b61293b6145bc565b501660005260096020526104c961296061295b60026040600020016145e7565b61644d565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b9291906129d06020916040865260408601906121e7565b930152565b346106055760606003193601126106055760043567ffffffffffffffff811161060557612a06903690600401610fd5565b612a0e610c7c565b9060443567ffffffffffffffff811161060557612a2f9036906004016115b7565b50612a386143e0565b50612a43828261617b565b606081013561ffff819316612b04575b5061249d6020612add9201612a6781613afa565b604080517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815233602082015290810186905267ffffffffffffffff91909116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10908060608101612495565b612ae56164ef565b612aed611537565b91825260208201526104c9604051928392836129b9565b909150612710612b1d61ffff600b5460101c168361475f565b048103908111611bcd579061249d612a53565b34610605576020600319360112610605576104c9612b5360043561249d81610fb6565b6040519182916020835260208301906109a0565b602060408183019282815284518094520192019060005b818110612b8b5750505090565b825167ffffffffffffffff16845260209384019390920191600101612b7e565b3461060557600060031936011261060557612bc4617204565b805190601f19612bd66126ca84614223565b0136602084013760005b8151811015612c12578067ffffffffffffffff612bff60019385614474565b5116612c0b8286614474565b5201612be0565b604051806104c98582612b67565b346106055760206003193601126106055767ffffffffffffffff600435612c4681610fb6565b612c4e6145bc565b501660005260096020526104c961296061295b60406000206145e7565b346106055760006003193601126106055760206040516101188152f35b34610605576000600319360112610605576020612ca3614663565b65ffffffffffff60405191168152f35b34610605576000600319360112610605576001546001600160a01b03163303612d7e5760015460a081901c65ffffffffffff16906001600160a01b031681158015612d74575b612d4657612d1b90612d156001600160a01b03600254166165ce565b50614f06565b50600180547fffffffffffff0000000000000000000000000000000000000000000000000000169055005b507f19ca5ebb0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b5042821015612cf9565b7fc22c8022000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b3461060557600060031936011261060557604065ffffffffffff612de66001549065ffffffffffff6001600160a01b0383169260a01c1690565b6001600160a01b03849392935193168352166020820152f35b8015150361060557565b6fffffffffffffffffffffffffffffffff81160361060557565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc60609101126106055760405190612e5a82611479565b81602435612e6781612dff565b8152604435612e7581612e09565b6020820152604060643591612e8983612e09565b0152565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c60609101126106055760405190612ec482611479565b81608435612ed181612dff565b815260a435612edf81612e09565b6020820152604060c43591612e8983612e09565b919082606091031261060557604051612f0b81611479565b60408082948035612f1b81612dff565b84526020810135612f2b81612e09565b6020850152013591612e8983612e09565b346106055760e060031936011261060557600435612f5981610fb6565b612f6236612e23565b612f6b36612e8d565b3360009081527f135a184f1ea19f46a2a5eda150dca2e4f81e2266dd73bb074c2f2a4e0c0032a0602052604090205490919060ff161580612fb4575b6121375761066a92615f49565b503360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff1615612fa7565b346106055760406003193601126106055760043560243561300d81610700565b8115610e5c5781613032610e5261066a94600052600060205260016040600020015490565b61664c565b3461060557600060031936011261060557613050614e09565b600180547fffffffffffff0000000000000000000000000000000000000000000000000000811690915560a01c65ffffffffffff1661308b57005b7f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109600080a1005b346106055760206003193601126106055760043567ffffffffffffffff8111610605576130e3903690600401610bb8565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff16156107a55760005b81811061312657005b80837fece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d67ffffffffffffffff6131626120c6600196888661469c565b61328661317d613173878a8861469c565b60208101906146dc565b8961319861318e8a838b969b61469c565b60408101906146dc565b6131cb6131c18c6131b96131af82888b989b61469c565b60608101906146dc565b96909561469c565b60808101906146dc565b9590946131e16131dc36838f61423b565b61652a565b6131ef6131dc36858561423b565b6131fd6131dc36878761423b565b61320b6131dc36898961423b565b6132788c61322361321a611546565b9184369161423b565b815261323036868661423b565b602082015261324036888861423b565b6040820152613250368a8a61423b565b60608201526132738b67ffffffffffffffff16600052600e602052604060002090565b61482e565b604051998a99169b89614934565b0390a20161311d565b346106055760006003193601126106055760206040516107d08152f35b34610605576020600319360112610605576004356132c981610700565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff16156107a5576020816107927fd63806009f622849e3b7cfd82d762420d57574c39f945f678871b2b5f1e8ce4b93616622565b346106055760006003193601126106055760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b346106055760206003193601126106055767ffffffffffffffff60043561339781610fb6565b61339f6145bc565b501660005260106020526104c9604060002060ff6002604051926133c284611495565b8054845260018101546020850152015463ffffffff81166040840152818160201c161515606084015260281c16151560808201526040519182918291909160808060a0830194805184526020810151602085015263ffffffff604082015116604085015260608101511515606085015201511515910152565b346106055760006003193601126106055760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b3461060557602060031936011261060557602060ff611f6960043561349c81610700565b7f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af600052600084526040600020906001600160a01b0316600052602052604060002090565b34610605576134ef3661161b565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb56020526040902054909391929060ff16156107a55783916000915b8083106137df5750505060009163ffffffff4216925b82811061355057005b61356361355e828585614a45565b614b04565b906060820161357281516166e3565b608083019361358185516166e3565b604084019081515115611caf576135bb610f806135b66135a9885167ffffffffffffffff1690565b67ffffffffffffffff1690565b617623565b613794576136f46135f46135da879a999a5167ffffffffffffffff1690565b67ffffffffffffffff166000526009602052604060002090565b6136b7896136b1875161369861361d60408301516fffffffffffffffffffffffffffffffff1690565b9161367f61364861364160208401516fffffffffffffffffffffffffffffffff1690565b9251151590565b613676613653611555565b6fffffffffffffffffffffffffffffffff851681529763ffffffff166020890152565b15156040870152565b6fffffffffffffffffffffffffffffffff166060850152565b6fffffffffffffffffffffffffffffffff166080830152565b82614b9b565b6136e9896136e08a5161369861361d60408301516fffffffffffffffffffffffffffffffff1690565b60028301614b9b565b600484519101614ca7565b602085019660005b88518051821015613737579061373160019261372a836137248c5167ffffffffffffffff1690565b92614474565b5190615d02565b016136fc565b5050979650949061378b7f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c293926137786001975167ffffffffffffffff1690565b9251935190519060405194859485614d74565b0390a101613547565b610f5a6137a9865167ffffffffffffffff1690565b7f1d5ad3c50000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b9091926137f06120c6858486613c95565b94613807610f8067ffffffffffffffff88166173aa565b6138fb57613834600561382e8867ffffffffffffffff166000526009602052604060002090565b0161724f565b9360005b85518110156138805760019061387960056138678b67ffffffffffffffff166000526009602052604060002090565b01613872838a614474565b5190617435565b5001613838565b509350937f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166138ed600193976138d26138cd8267ffffffffffffffff166000526009602052604060002090565b6149b4565b60405167ffffffffffffffff90911681529081906020820190565b0390a1019190939293613531565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff861660045260246000fd5b346106055760006003193601126106055760206040517ff12fb6eaf1f045883c82d7d192627f7a36a50ce00c45e305919895908135a8a88152f35b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205492939260ff16156107a55760005b828110613a2b5750505060005b8181106139c457505050565b8067ffffffffffffffff6139de6120c66001948688613c95565b60006139fe8267ffffffffffffffff16600052600f602052604060002090565b55167f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8600080a2016139b8565b80613a3c6120c66001938686613ae5565b7f56f77aeff2def50c8b5f5a0df3bab7183df09bf36c6feba496bb42551db77d7067ffffffffffffffff6020613a73858989613ae5565b0192613a9c84613a978367ffffffffffffffff16600052600f602052604060002090565b613b1c565b613aad604051928392169482613c33565b0390a2016139ab565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9190811015613af55760a0020190565b613ab6565b356109f281610fb6565b63ffffffff81160361060557565b356109f281612dff565b90606090613b5f8135613b2e81613b04565b849063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b6020810135613b6d81613b04565b67ffffffff0000000084549160201b16807fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff83161785557fffffffffffffffffffffffffffffffffffffffff0000000000000000ffffffff6bffffffff00000000000000006040850135613be081613b04565b60401b169216171783550135613bf581612dff565b81547fffffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffff1690151560601b6cff00000000000000000000000016179055565b919091606080608083019463ffffffff8135613c4e81613b04565b16845263ffffffff6020820135613c6481613b04565b16602085015263ffffffff6040820135613c7d81613b04565b1660408501520135613c8e81612dff565b1515910152565b9190811015613af55760051b0190565b67ffffffffffffffff9060006060604051613cbf816114b1565b828152826020820152826040820152015216600052600f602052604060002060ff60405191613ced836114b1565b5463ffffffff8116835263ffffffff8160201c16602084015263ffffffff8160401c16604084015260601c161515606082015290565b60405190613d3082611458565b60008252565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610605570180359067ffffffffffffffff82116106055760200191813603831361060557565b6020818303126106055780359067ffffffffffffffff821161060557016040818303126106055760405191613dbb836114cd565b813567ffffffffffffffff81116106055781613dd89184016115b7565b8352602082013567ffffffffffffffff811161060557613df892016115b7565b602082015290565b9081602091031261060557516109f281612dff565b9091613e2c6109f2936040845260408401906109a0565b9160208184039101526109a0565b6040513d6000823e3d90fd5b356109f281610700565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff16156107a55760005b828110613ec95750907fc97f93e817584952f1c1d633f93784b8430f0633d002f9dcc4de4fe2780424d091613ec460405192839283614146565b0390a1565b613edc613ed782858561401e565b61402e565b8051158015613ff8575b613f7f5790613f7982613f7461227860606001965193613f656020820151613f5c613f18604085015163ffffffff1690565b613f54613f286080870151151590565b91613f3660a0880151151590565b94613f3f611555565b9b8c5260208c015263ffffffff1660408b0152565b151588860152565b15156080870152565b015167ffffffffffffffff1690565b6140a1565b01613e8a565b604080517f19d7585700000000000000000000000000000000000000000000000000000000815282516004820152602083015160248201529082015163ffffffff166044820152606082015167ffffffffffffffff16606482015260808201511515608482015260a090910151151560a482015260c490fd5b5067ffffffffffffffff614017606083015167ffffffffffffffff1690565b1615613ee6565b9190811015613af55760c0020190565b60c0813603126106055760a060405191614047836114e9565b8035835260208101356020840152604081013561406381613b04565b6040840152606081013561407681610fb6565b6060840152608081013561408981612dff565b6080840152013561409981612dff565b60a082015290565b6002608091835181556020840151600182015501916140f563ffffffff604083015116849063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b6060810151835492909101517fffffffffffffffffffffffffffffffffffffffffffffffffffff0000ffffffff90921690151560201b64ff00000000161790151560281b65ff000000000016179055565b602080825281018390526040019160005b8181106141645750505090565b90919260c080600192863581526020870135602082015263ffffffff604088013561418e81613b04565b16604082015267ffffffffffffffff60608801356141ab81610fb6565b16606082015260808701356141bf81612dff565b1515608082015260a08701356141d481612dff565b151560a0820152019401929101614157565b9067ffffffffffffffff6109f292166000526009602052600560406000200190602081519101209060019160005201602052604060002054151590565b67ffffffffffffffff81116114745760051b60200190565b92919061424781614223565b936142556040519586611505565b602085838152019160051b810192831161060557905b82821061427757505050565b60208091833561428681610700565b81520191019061426b565b90816020910312610605575190565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156112fb57600091614310575090565b6109f2915060203d602011614332575b61432a8183611505565b810190614291565b503d614320565b67ffffffffffffffff6109f291166000526008602052604060002054151590565b67ffffffffffffffff16600052600e60205260406000209160028110156143a157600114614390578160016109f2930190615e72565b81600260036109f294019101615e72565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9190811015613af5576060020190565b604051906143ed826114cd565b60606020838281520152565b9060405161440681611495565b608060ff600283958054855260018101546020860152015463ffffffff81166040850152818160201c161515606085015260281c161515910152565b601f8260209493601f19938186528686013760008582860101520116010190565b9160206109f2938181520191614442565b8051821015613af55760209160051b010190565b90600182811c921680156144d1575b60208310146144a257565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691614497565b90604051918260008254926144ef84614488565b808452936001811690811561455b5750600114614514575b50610fd392500383611505565b90506000929192526020600020906000915b81831061453f575050906020610fd39282010138614507565b6020919350806001915483858901015201910190918492614526565b60209350610fd39592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138614507565b60409067ffffffffffffffff6109f295931681528160208201520191614442565b604051906145c982611495565b60006080838281528260208201528260408201528260608201520152565b906040516145f481611495565b60806fffffffffffffffffffffffffffffffff6001839560ff8154848116875263ffffffff81871c16602088015260a01c1615156040860152015481808216166060850152821c16910152565b67ffffffffffffffff1660005260096020526109f260046040600020016144db565b6002548060d01c8015159081614692575b50156146885760a01c65ffffffffffff1690565b5060015460d01c90565b9050421138614674565b9190811015613af55760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6181360301821215610605570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610605570180359067ffffffffffffffff821161060557602001918160051b3603831361060557565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b81810292918115918404141715611bcd57565b9161478c918354906000199060031b92831b921b19161790565b9055565b81811061479b575050565b60008155600101614790565b81519167ffffffffffffffff831161147457680100000000000000008311611474576020908254848455808510614811575b500190600052602060002060005b8381106147f45750505050565b60019060206001600160a01b0385511694019381840155016147e7565b614828908460005285846000209182019101614790565b386147d9565b90805180519067ffffffffffffffff8211611474576801000000000000000082116114745760209084548386558084106148d3575b500183600052602060002060005b8381106148b057505050509060036060836148966020610fd3960151600186016147a7565b6148a76040820151600286016147a7565b015191016147a7565b60019060206148c685516001600160a01b031690565b9401938184015501614871565b6148ea908660005284846000209182019101614790565b38614863565b9160209082815201919060005b81811061490a5750505090565b9091926020806001926001600160a01b03873561492681610700565b1681520194019291016148fd565b96949261497294614956614964936109f29b999560808c5260808c01916148f0565b9189830360208b01526148f0565b9186830360408801526148f0565b9260608185039101526148f0565b8054906000815581614990575050565b6000526020600020908101905b8181106149a8575050565b6000815560010161499d565b6005610fd39160008155600060018201556000600282015560006003820155600481016149e18154614488565b90816149f0575b505001614980565b81601f60009311600114614a085750555b38806149e8565b81835260208320614a2391601f01861c810190600101614790565b808252602082209081548360011b906000198560031b1c191617905555614a01565b9190811015613af55760051b810135907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee181360301821215610605570190565b9080601f83011215610605578135614a9c81614223565b92614aaa6040519485611505565b81845260208085019260051b820101918383116106055760208201905b838210614ad657505050505090565b813567ffffffffffffffff811161060557602091614af9878480948801016115b7565b815201910190614ac7565b610120813603126106055760405190614b1c82611495565b614b2581610fc8565b8252602081013567ffffffffffffffff811161060557614b489036908301614a85565b602083015260408101359067ffffffffffffffff821161060557614b72614b9392369083016115b7565b6040840152614b843660608301612ef3565b606084015260c0369101612ef3565b608082015290565b8151815460208401516040850151608091821b73ffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffff000000000000000000000000000000000000000000909416939093179290921791151560a01b74ff000000000000000000000000000000000000000016919091178355606084015193810151901b7fffffffffffffffffffffffffffffffff0000000000000000000000000000000016921691909117600190910155565b9190601f8111614c7157505050565b610fd3926000526020600020906020601f840160051c83019310614c9d575b601f0160051c0190614790565b9091508190614c90565b919091825167ffffffffffffffff811161147457614ccf81614cc98454614488565b84614c62565b6020601f8211600114614d0b57819061478c939495600092614d00575b50506000198260011b9260031b1c19161790565b015190503880614cec565b601f19821690614d2084600052602060002090565b9160005b818110614d5c57509583600195969710614d43575b505050811b019055565b015160001960f88460031b161c19169055388080614d39565b9192602060018192868b015181550194019201614d24565b614dd8614da3610fd39597969467ffffffffffffffff60a09516845261010060208501526101008401906109a0565b9660408301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b01906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff1615614e4257565b7fe2517d3f0000000000000000000000000000000000000000000000000000000060005233600452600060245260446000fd5b80600052600060205260ff614ea1336040600020906001600160a01b0316600052602052604060002090565b541615614eab5750565b7fe2517d3f000000000000000000000000000000000000000000000000000000006000523360045260245260446000fd5b6109f2907f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af616967565b600254906001600160a01b038216610e5c576109f2917fffffffffffffffffffffffff00000000000000000000000000000000000000006001600160a01b0383169116176002556000616967565b908115614f65575b6109f291616967565b600254916001600160a01b038316610e5c577fffffffffffffffffffffffff00000000000000000000000000000000000000009092166001600160a01b03821617600255614f5c565b9190811015613af55760e0020190565b356109f281612e09565b919060005b818110614fda5750509050565b614fe5818386614fae565b614fee81613afa565b614ffa610f8082614339565b6128d6578161509061510a926150966020600197960161502261501d3683612ef3565b6166e3565b6150906150438467ffffffffffffffff16600052600c602052604060002090565b91825461506361505a8263ffffffff9060801c1690565b63ffffffff1690565b1590816152fd575b816152d5575b816152ba575b816152a5575b5080615296575b615241575b3690612ef3565b90616ab4565b6150c560808401916150ab61501d3685612ef3565b67ffffffffffffffff16600052600d602052604060002090565b9283546150dc61505a8263ffffffff9060801c1690565b15908161520a575b816151e2575b816151c7575b816151b2575b50806151a3575b615110575b503690612ef3565b01614fcd565b61511f60a061515c9201614fbe565b84906fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b82547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff000000000000000000000000000000001617835538615102565b506151ad82613b12565b6150fd565b6151c1915060a01c60ff161590565b386150f6565b6fffffffffffffffffffffffffffffffff81161591506150f0565b90506fffffffffffffffffffffffffffffffff6152028987015460801c90565b1615906150ea565b90506fffffffffffffffffffffffffffffffff615239898701546fffffffffffffffffffffffffffffffff1690565b1615906150e4565b61525061511f60408901614fbe565b82547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178355615089565b506152a081613b12565b615084565b6152b4915060a01c60ff161590565b3861507d565b6fffffffffffffffffffffffffffffffff8116159150615077565b90506fffffffffffffffffffffffffffffffff6152f58c86015460801c90565b161590615071565b90506fffffffffffffffffffffffffffffffff61532c8c8601546fffffffffffffffffffffffffffffffff1690565b16159061506b565b90336001600160a01b038216036153515761534e9161664c565b50565b7f6697b2320000000000000000000000000000000000000000000000000000000060005260046000fd5b600060808201615390610f80610afb83613e46565b6155465750602082019161543160206153d66153ae6135a987613afa565b60801b7fffffffffffffffffffffffffffffffff000000000000000000000000000000001690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116600482015291829081906024820190565b03816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156112fb578391615527575b506154ff5761548361547e84613afa565b616d59565b61548c83613afa565b906154a5610f8060a08301936116156111618686613d36565b6154bf575050610fd392916154ba9150613afa565b616e14565b6154c99250613d36565b906125366040519283927f24eb47e500000000000000000000000000000000000000000000000000000000845260048401614463565b6004827f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b615540915060203d6020116112f4576112e68183611505565b3861546d565b61555261558691613e46565b7f961c9a4f0000000000000000000000000000000000000000000000000000000083526001600160a01b0316600452602490565b90fd5b91608083019061559e610f80610afb84613e46565b6157155760208401936155bb60206153d66153ae6135a989613afa565b03816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156112fb576000916156f6575b506156cc5761560961547e86613afa565b61561285613afa565b9061562b610f8060a08301936116156111618686613d36565b6154bf57505061ffff16156156bf5767ffffffffffffffff906156ba61569561568f866120c661567f6150ab7f41a8aa8df7945f0fb8ac5f7d88279638d9dc2ef9a6bf4ec9a53b80681b34aff7999a613afa565b8961568988613e46565b91616eab565b92613e46565b94604051938493169583602090939291936001600160a01b0360408201951681520152565b0390a2565b506154ba610fd392613afa565b7f53ad11d80000000000000000000000000000000000000000000000000000000060005260046000fd5b61570f915060203d6020116112f4576112e68183611505565b386155f8565b610f5a61572183613e46565b7f961c9a4f000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b60405190615763826114cd565b6000825260208201600081526020820151917fffffffff000000000000000000000000000000000000000000000000000000006028602483015160e01c92015193167fb148ea5f000000000000000000000000000000000000000000000000000000008114158061580d575b6157e0575063ffffffff1683525290565b7fa176027f0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b507f3047587c000000000000000000000000000000000000000000000000000000008114156157cf565b80516101188110615a91575060048101517f000000000000000000000000000000000000000000000000000000000000000063ffffffff811663ffffffff831603615a58575050600881015190600c81015191608c82015191609081015193609482015160b88301519360f860d8850151940151916158ba895163ffffffff1690565b63ffffffff811663ffffffff841603615a1e57507f000000000000000000000000000000000000000000000000000000000000000063ffffffff811663ffffffff8616036159e457506107d063ffffffff8916036159aa576107d063ffffffff8216036159715750916159339593916020979593616370565b910151808203615941575050565b7f7be225b60000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b7f0389caa2000000000000000000000000000000000000000000000000000000006000526107d060045263ffffffff1660245260446000fd5b7f22e102a0000000000000000000000000000000000000000000000000000000006000526107d060045263ffffffff881660245260446000fd5b7f77e480260000000000000000000000000000000000000000000000000000000060005263ffffffff908116600452841660245260446000fd5b7fe366a1170000000000000000000000000000000000000000000000000000000060005263ffffffff908116600452821660245260446000fd5b7f960693cd0000000000000000000000000000000000000000000000000000000060005263ffffffff9081166004521660245260446000fd5b7f758b22cc0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b80518015615b6357602003615b2657615ae06020825183010160208301614291565b9060ff8211615af0575060ff1690565b612536906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352600483016109e1565b612536906040519182917f953576f700000000000000000000000000000000000000000000000000000000835260206004840181815201906109a0565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff8211611bcd57565b60ff16604d8111611bcd57600a0a90565b8015615bbb576000190490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b8115615bbb570490565b907f000000000000000000000000000000000000000000000000000000000000000060ff811660ff8316818114615cda5711615caf57615c348282615b89565b91604d60ff8416118015615c96575b615c5c57505090615c566109f292615b9d565b9061475f565b7fa9cb113d0000000000000000000000000000000000000000000000000000000060005260ff908116600452166024525060445260646000fd5b50615ca8615ca384615b9d565b615bae565b8411615c43565b615cb98183615b89565b91604d60ff841611615c5c57505090615cd46109f292615b9d565b90615bea565b5050505090565b60409067ffffffffffffffff6109f2949316815281602082015201906109a0565b90805115611caf578051602082012067ffffffffffffffff831692836000526009602052615d37826005604060002001617679565b15615d8b575081615d7f7f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea93615d7a6156ba94600052600a602052604060002090565b614ca7565b604051918291826109e1565b90506125366040519283927f393b8ad200000000000000000000000000000000000000000000000000000000845260048401615ce1565b9065ffffffffffff8091169116019065ffffffffffff8211611bcd57565b906040519182815491828252602082019060005260206000209260005b818110615e12575050610fd392500383611505565b84546001600160a01b0316835260019485019487945060209093019201615dfd565b91908201809211611bcd57565b90615e4b82614223565b615e586040519182611505565b828152601f19615e688294614223565b0190602036910137565b615e7b90615de0565b916006548015159182615f3e575b5050615e93575090565b615e9c90615de0565b805180615ea857505090565b615eb9615ebe918495939451615e34565b615e41565b9060005b8451811015615efc5780615ef6615ede61170f60019489614474565b615ee88387614474565b906001600160a01b03169052565b01615ec2565b509160005b8151811015615f375780615f31615f1d61170f60019486614474565b615ee8615f2b848a51615e34565b87614474565b01615f01565b5090925050565b101590503880615e89565b67ffffffffffffffff16600081815260086020526040902054909291901561604b579161604860e09261601485615fa07f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b976166e3565b846000526009602052615fb7816040600020616ab4565b615fc0836166e3565b846000526009602052615fda836002604060002001616ab4565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6000906080810161608f610f80610afb83613e46565b61616e575060208101906160ad60206153d66153ae6135a986613afa565b03816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156112fb57849161614f575b5061612757610fd3929160608261610d61610860406161229601613e46565b61711e565b61611961547e84613afa565b01359250613afa565b617193565b6004837f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b616168915060203d6020116112f4576112e68183611505565b386160e9565b6155866155528492613e46565b608081019061618f610f80610afb84613e46565b6157155760208101906161ac60206153d66153ae6135a986613afa565b03816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156112fb576000916162f8575b506156cc5780616200616108604060609401613e46565b61620c61547e84613afa565b01359261ffff811690811515806162d9575b156162c857600b5461ffff169182116162915750507f7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b2916156ba61569561568f846120c661567f61627767ffffffffffffffff98613afa565b67ffffffffffffffff16600052600c602052604060002090565b7fe08f03ef0000000000000000000000000000000000000000000000000000000060005261ffff9081166004521660245260446000fd5b5050610fd392915061612290613afa565b506162f16162ea600b5461ffff1690565b61ffff1690565b151561621e565b616311915060203d6020116112f4576112e68183611505565b386161e9565b949290939163ffffffff90604051958260208801981688526040870152166060850152608084015260a083015260c0820152600060e08201526107d0610100820152610100815261636a61012082611505565b51902090565b959263ffffffff8095929693604051978260208a019a168a526040890152166060870152608086015260a085015260c0840152600060e084015216610100820152610100815261636a61012082611505565b602081519101517fffffffff00000000000000000000000000000000000000000000000000000000604051927fb148ea5f00000000000000000000000000000000000000000000000000000000602085015260e01b1660248301526028820152602881526109f2604882611505565b906000198201918211611bcd57565b91908203918211611bcd57565b6164556145bc565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff82511690602083019163ffffffff8351164203428111611bcd576164b9906fffffffffffffffffffffffffffffffff6080870151169061475f565b8101809111611bcd576164df6fffffffffffffffffffffffffffffffff929183926176e4565b161682524263ffffffff16905290565b60405160ff7f0000000000000000000000000000000000000000000000000000000000000000166020820152602081526109f2604082611505565b80519060005b82811061653c57505050565b60018101808211611bcd575b8381106165585750600101616530565b6001600160a01b0361656a8385614474565b511661657c610f9a61170f8487614474565b1461658957600101616548565b610f5a61659961170f8486614474565b7fa1726e40000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b6109f2906001600160a01b03600254166001600160a01b038216146165f5575b60006174f7565b7fffffffffffffffffffffffff0000000000000000000000000000000000000000600254166002556165ee565b6109f2907f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af6174f7565b906109f29180158061668b575b156174f7577fffffffffffffffffffffffff0000000000000000000000000000000000000000600254166002556174f7565b506001600160a01b03600254166001600160a01b03831614616659565b610fd39092919260608101936fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b8051156167875760408101516fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff61674761673260208501516fffffffffffffffffffffffffffffffff1690565b6fffffffffffffffffffffffffffffffff1690565b9116116167515750565b612536906040519182917f8020d124000000000000000000000000000000000000000000000000000000008352600483016166a8565b6fffffffffffffffffffffffffffffffff6167b560408301516fffffffffffffffffffffffffffffffff1690565b16158015906167fc575b6167c65750565b612536906040519182917fd68af9cc000000000000000000000000000000000000000000000000000000008352600483016166a8565b5061681d61673260208301516fffffffffffffffffffffffffffffffff1690565b15156167bf565b90616887610fd3926002548060d01c806168d3575b50507fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff79ffffffffffff00000000000000000000000000000000000000006002549260a01b16911617600255565b79ffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffffffff00000000000000000000000000000000000000000000000000006002549260d01b16911617600255565b42111561693c5765ffffffffffff6169359160a01c1679ffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffffffff00000000000000000000000000000000000000000000000000006001549260d01b16911617600155565b3880616839565b507f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5600080a1616935565b80600052600060205260ff616993836040600020906001600160a01b0316600052602052604060002090565b5416616a22578060005260006020526169c3826040600020906001600160a01b0316600052602052604060002090565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008254161790556001600160a01b03339216907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d600080a4600190565b5050600090565b15616a3057565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1991616c8a613ec4928054616afc616af661505a8363ffffffff9060801c1690565b42616440565b9081616c96575b5050616c446001616b2760208601516fffffffffffffffffffffffffffffffff1690565b92616bb2616b756167326fffffffffffffffffffffffffffffffff616b5c85546fffffffffffffffffffffffffffffffff1690565b166fffffffffffffffffffffffffffffffff88166176e4565b82906fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b616c05616bbf8751151590565b82547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178255565b019182906fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b604083015181546fffffffffffffffffffffffffffffffff1660809190911b7fffffffffffffffffffffffffffffffff0000000000000000000000000000000016179055565b604051918291826166a8565b616732616b75916fffffffffffffffffffffffffffffffff616d0a616d119582616d0360018a01549282616cfc616cf5616cdf876fffffffffffffffffffffffffffffffff1690565b996fffffffffffffffffffffffffffffffff1690565b9560801c90565b169061475f565b9116615e34565b91166176e4565b80547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff00000000000000000000000000000000161781553880616b03565b67ffffffffffffffff16616d7a816000526008602052604060002054151590565b15616de757503360009081527f04c57a7d2bd5d0e733fe996f5b5aecc738999f0a2f9ddc4137bc3e1665bdf893602052604090205460ff1615616db957565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fa9902c7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b67ffffffffffffffff7f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c91169182600052600960205280616e8860026040600020016001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391616eab565b604080516001600160a01b039092168252602082019290925290819081016156ba565b8054939290919060ff60a086901c161580156170cc575b6170c557616ee16fffffffffffffffffffffffffffffffff8616616732565b9060018401958654616f1b616af661505a616f0e616732856fffffffffffffffffffffffffffffffff1690565b9460801c63ffffffff1690565b80617031575b5050838110616ff35750828210616f815750610fd3939450616f469161673291616440565b6fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b90616fb8610f5a93616fb3616fa484616f9e6167328c5460801c90565b93616440565b616fad83616431565b90615e34565b615bea565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024526001600160a01b0316604452606490565b7f1a76572a0000000000000000000000000000000000000000000000000000000060005260045260248390526001600160a01b031660445260646000fd5b82859293951161709b5761704b6167326170529460801c90565b91856174da565b84547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178555913880616f21565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050509050565b508115616ec2565b65ffffffffffff81116170ec5765ffffffffffff1690565b7f6dfcc65000000000000000000000000000000000000000000000000000000000600052603060045260245260446000fd5b7f00000000000000000000000000000000000000000000000000000000000000006171465750565b6001600160a01b0316806000526004602052604060002054156171665750565b7fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b67ffffffffffffffff7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894491169182600052600960205280616e8860406000206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391616eab565b604051906007548083528260208101600760005260206000209260005b818110617236575050610fd392500383611505565b8454835260019485019487945060209093019201617221565b906040519182815491828252602082019060005260206000209260005b818110617281575050610fd392500383611505565b845483526001948501948794506020909301920161726c565b8054821015613af55760005260206000200190600090565b805480156172da5760001901906172c9828261729a565b60001982549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600081815260046020526040902054908115616a2257600019820190828211611bcd57600354926000198401938411611bcd578383600095617369950361736f575b50505061735860036172b2565b600490600052602052604060002090565b55600190565b61735861739b916173916173876173a195600361729a565b90549060031b1c90565b928391600361729a565b90614772565b5538808061734b565b600081815260086020526040902054908115616a2257600019820190828211611bcd57600754926000198401938411611bcd578383600095617369950361740a575b5050506173f960076172b2565b600890600052602052604060002090565b6173f961739b9161742261738761742c95600761729a565b928391600761729a565b553880806173ec565b60018101918060005282602052604060002054928315156000146174d1576000198401848111611bcd578354936000198501948511611bcd576000958583617369976174899503617498575b5050506172b2565b90600052602052604060002090565b6174b861739b916174af6173876174c8958861729a565b9283918761729a565b8590600052602052604060002090565b55388080617481565b50505050600090565b926174e5919261475f565b8101809111611bcd576109f2916176e4565b80600052600060205260ff617523836040600020906001600160a01b0316600052602052604060002090565b541615616a2257806000526000602052617554826040600020906001600160a01b0316600052602052604060002090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541690556001600160a01b03339216907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b600080a4600190565b60008181526004602052604090205461761d5760035468010000000000000000811015611474576176046175ed826001859401600355600361729a565b81939154906000199060031b92831b921b19161790565b9055600354906000526004602052604060002055600190565b50600090565b60008181526008602052604090205461761d5760075468010000000000000000811015611474576176606175ed826001859401600755600761729a565b9055600754906000526008602052604060002055600190565b6000828152600182016020526040902054616a22578054906801000000000000000082101561147457826176b76175ed84600180960185558461729a565b905580549260005201602052604060002055600190565b620697808110156176dc5790565b506206978090565b90808210156176f1575090565b905090565b91929015617771575081511561770a575090565b3b156177135790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156177845750805190602001fd5b612536906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352600483016109e156fea164736f6c634300081a000aad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5",
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) AUTHORIZEDCALLERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "AUTHORIZED_CALLER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) AUTHORIZEDCALLERROLE() ([32]byte, error) {
	return _USDCTokenPoolCCTPV2.Contract.AUTHORIZEDCALLERROLE(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) AUTHORIZEDCALLERROLE() ([32]byte, error) {
	return _USDCTokenPoolCCTPV2.Contract.AUTHORIZEDCALLERROLE(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _USDCTokenPoolCCTPV2.Contract.DEFAULTADMINROLE(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _USDCTokenPoolCCTPV2.Contract.DEFAULTADMINROLE(&_USDCTokenPoolCCTPV2.CallOpts)
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) RATELIMITERADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "RATE_LIMITER_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) RATELIMITERADMINROLE() ([32]byte, error) {
	return _USDCTokenPoolCCTPV2.Contract.RATELIMITERADMINROLE(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) RATELIMITERADMINROLE() ([32]byte, error) {
	return _USDCTokenPoolCCTPV2.Contract.RATELIMITERADMINROLE(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) DefaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "defaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) DefaultAdmin() (common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.DefaultAdmin(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) DefaultAdmin() (common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.DefaultAdmin(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "defaultAdminDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) DefaultAdminDelay() (*big.Int, error) {
	return _USDCTokenPoolCCTPV2.Contract.DefaultAdminDelay(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) DefaultAdminDelay() (*big.Int, error) {
	return _USDCTokenPoolCCTPV2.Contract.DefaultAdminDelay(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "defaultAdminDelayIncreaseWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _USDCTokenPoolCCTPV2.Contract.DefaultAdminDelayIncreaseWait(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _USDCTokenPoolCCTPV2.Contract.DefaultAdminDelayIncreaseWait(&_USDCTokenPoolCCTPV2.CallOpts)
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetRoleAdmin(&_USDCTokenPoolCCTPV2.CallOpts, role)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetRoleAdmin(&_USDCTokenPoolCCTPV2.CallOpts, role)
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) HasRateLimitAdminRole(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "hasRateLimitAdminRole", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) HasRateLimitAdminRole(account common.Address) (bool, error) {
	return _USDCTokenPoolCCTPV2.Contract.HasRateLimitAdminRole(&_USDCTokenPoolCCTPV2.CallOpts, account)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) HasRateLimitAdminRole(account common.Address) (bool, error) {
	return _USDCTokenPoolCCTPV2.Contract.HasRateLimitAdminRole(&_USDCTokenPoolCCTPV2.CallOpts, account)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _USDCTokenPoolCCTPV2.Contract.HasRole(&_USDCTokenPoolCCTPV2.CallOpts, role, account)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _USDCTokenPoolCCTPV2.Contract.HasRole(&_USDCTokenPoolCCTPV2.CallOpts, role, account)
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) PendingDefaultAdmin(opts *bind.CallOpts) (PendingDefaultAdmin,

	error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "pendingDefaultAdmin")

	outstruct := new(PendingDefaultAdmin)
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewAdmin = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) PendingDefaultAdmin() (PendingDefaultAdmin,

	error) {
	return _USDCTokenPoolCCTPV2.Contract.PendingDefaultAdmin(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) PendingDefaultAdmin() (PendingDefaultAdmin,

	error) {
	return _USDCTokenPoolCCTPV2.Contract.PendingDefaultAdmin(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) PendingDefaultAdminDelay(opts *bind.CallOpts) (PendingDefaultAdminDelay,

	error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "pendingDefaultAdminDelay")

	outstruct := new(PendingDefaultAdminDelay)
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewDelay = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) PendingDefaultAdminDelay() (PendingDefaultAdminDelay,

	error) {
	return _USDCTokenPoolCCTPV2.Contract.PendingDefaultAdminDelay(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) PendingDefaultAdminDelay() (PendingDefaultAdminDelay,

	error) {
	return _USDCTokenPoolCCTPV2.Contract.PendingDefaultAdminDelay(&_USDCTokenPoolCCTPV2.CallOpts)
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "acceptDefaultAdminTransfer")
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.AcceptDefaultAdminTransfer(&_USDCTokenPoolCCTPV2.TransactOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.AcceptDefaultAdminTransfer(&_USDCTokenPoolCCTPV2.TransactOpts)
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) ApplyFinalityConfigUpdates(opts *bind.TransactOpts, finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "applyFinalityConfigUpdates", finalityThreshold, customFinalityTransferFeeBps, rateLimitConfigArgs)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) ApplyFinalityConfigUpdates(finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.ApplyFinalityConfigUpdates(&_USDCTokenPoolCCTPV2.TransactOpts, finalityThreshold, customFinalityTransferFeeBps, rateLimitConfigArgs)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) ApplyFinalityConfigUpdates(finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.ApplyFinalityConfigUpdates(&_USDCTokenPoolCCTPV2.TransactOpts, finalityThreshold, customFinalityTransferFeeBps, rateLimitConfigArgs)
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "beginDefaultAdminTransfer", newAdmin)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.BeginDefaultAdminTransfer(&_USDCTokenPoolCCTPV2.TransactOpts, newAdmin)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.BeginDefaultAdminTransfer(&_USDCTokenPoolCCTPV2.TransactOpts, newAdmin)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "cancelDefaultAdminTransfer")
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.CancelDefaultAdminTransfer(&_USDCTokenPoolCCTPV2.TransactOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.CancelDefaultAdminTransfer(&_USDCTokenPoolCCTPV2.TransactOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "changeDefaultAdminDelay", newDelay)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.ChangeDefaultAdminDelay(&_USDCTokenPoolCCTPV2.TransactOpts, newDelay)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.ChangeDefaultAdminDelay(&_USDCTokenPoolCCTPV2.TransactOpts, newDelay)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) GrantRateLimitAdminRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "grantRateLimitAdminRole", account)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GrantRateLimitAdminRole(account common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.GrantRateLimitAdminRole(&_USDCTokenPoolCCTPV2.TransactOpts, account)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) GrantRateLimitAdminRole(account common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.GrantRateLimitAdminRole(&_USDCTokenPoolCCTPV2.TransactOpts, account)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "grantRole", role, account)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.GrantRole(&_USDCTokenPoolCCTPV2.TransactOpts, role, account)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.GrantRole(&_USDCTokenPoolCCTPV2.TransactOpts, role, account)
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "renounceRole", role, account)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.RenounceRole(&_USDCTokenPoolCCTPV2.TransactOpts, role, account)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.RenounceRole(&_USDCTokenPoolCCTPV2.TransactOpts, role, account)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) RevokeRateLimitAdminRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "revokeRateLimitAdminRole", account)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) RevokeRateLimitAdminRole(account common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.RevokeRateLimitAdminRole(&_USDCTokenPoolCCTPV2.TransactOpts, account)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) RevokeRateLimitAdminRole(account common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.RevokeRateLimitAdminRole(&_USDCTokenPoolCCTPV2.TransactOpts, account)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "revokeRole", role, account)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.RevokeRole(&_USDCTokenPoolCCTPV2.TransactOpts, role, account)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.RevokeRole(&_USDCTokenPoolCCTPV2.TransactOpts, role, account)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "rollbackDefaultAdminDelay")
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.RollbackDefaultAdminDelay(&_USDCTokenPoolCCTPV2.TransactOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.RollbackDefaultAdminDelay(&_USDCTokenPoolCCTPV2.TransactOpts)
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

type USDCTokenPoolCCTPV2DefaultAdminDelayChangeCanceledIterator struct {
	Event *USDCTokenPoolCCTPV2DefaultAdminDelayChangeCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2DefaultAdminDelayChangeCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2DefaultAdminDelayChangeCanceled)
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
		it.Event = new(USDCTokenPoolCCTPV2DefaultAdminDelayChangeCanceled)
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

func (it *USDCTokenPoolCCTPV2DefaultAdminDelayChangeCanceledIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2DefaultAdminDelayChangeCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2DefaultAdminDelayChangeCanceled struct {
	Raw types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2DefaultAdminDelayChangeCanceledIterator, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2DefaultAdminDelayChangeCanceledIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "DefaultAdminDelayChangeCanceled", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2DefaultAdminDelayChangeCanceled) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2DefaultAdminDelayChangeCanceled)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseDefaultAdminDelayChangeCanceled(log types.Log) (*USDCTokenPoolCCTPV2DefaultAdminDelayChangeCanceled, error) {
	event := new(USDCTokenPoolCCTPV2DefaultAdminDelayChangeCanceled)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2DefaultAdminDelayChangeScheduledIterator struct {
	Event *USDCTokenPoolCCTPV2DefaultAdminDelayChangeScheduled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2DefaultAdminDelayChangeScheduledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2DefaultAdminDelayChangeScheduled)
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
		it.Event = new(USDCTokenPoolCCTPV2DefaultAdminDelayChangeScheduled)
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

func (it *USDCTokenPoolCCTPV2DefaultAdminDelayChangeScheduledIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2DefaultAdminDelayChangeScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2DefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2DefaultAdminDelayChangeScheduledIterator, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2DefaultAdminDelayChangeScheduledIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "DefaultAdminDelayChangeScheduled", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2DefaultAdminDelayChangeScheduled) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2DefaultAdminDelayChangeScheduled)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseDefaultAdminDelayChangeScheduled(log types.Log) (*USDCTokenPoolCCTPV2DefaultAdminDelayChangeScheduled, error) {
	event := new(USDCTokenPoolCCTPV2DefaultAdminDelayChangeScheduled)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2DefaultAdminTransferCanceledIterator struct {
	Event *USDCTokenPoolCCTPV2DefaultAdminTransferCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2DefaultAdminTransferCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2DefaultAdminTransferCanceled)
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
		it.Event = new(USDCTokenPoolCCTPV2DefaultAdminTransferCanceled)
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

func (it *USDCTokenPoolCCTPV2DefaultAdminTransferCanceledIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2DefaultAdminTransferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2DefaultAdminTransferCanceled struct {
	Raw types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2DefaultAdminTransferCanceledIterator, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2DefaultAdminTransferCanceledIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "DefaultAdminTransferCanceled", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2DefaultAdminTransferCanceled) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2DefaultAdminTransferCanceled)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseDefaultAdminTransferCanceled(log types.Log) (*USDCTokenPoolCCTPV2DefaultAdminTransferCanceled, error) {
	event := new(USDCTokenPoolCCTPV2DefaultAdminTransferCanceled)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2DefaultAdminTransferScheduledIterator struct {
	Event *USDCTokenPoolCCTPV2DefaultAdminTransferScheduled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2DefaultAdminTransferScheduledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2DefaultAdminTransferScheduled)
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
		it.Event = new(USDCTokenPoolCCTPV2DefaultAdminTransferScheduled)
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

func (it *USDCTokenPoolCCTPV2DefaultAdminTransferScheduledIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2DefaultAdminTransferScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2DefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*USDCTokenPoolCCTPV2DefaultAdminTransferScheduledIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2DefaultAdminTransferScheduledIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "DefaultAdminTransferScheduled", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2DefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2DefaultAdminTransferScheduled)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseDefaultAdminTransferScheduled(log types.Log) (*USDCTokenPoolCCTPV2DefaultAdminTransferScheduled, error) {
	event := new(USDCTokenPoolCCTPV2DefaultAdminTransferScheduled)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
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

type USDCTokenPoolCCTPV2FinalityConfigUpdatedIterator struct {
	Event *USDCTokenPoolCCTPV2FinalityConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2FinalityConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2FinalityConfigUpdated)
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
		it.Event = new(USDCTokenPoolCCTPV2FinalityConfigUpdated)
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

func (it *USDCTokenPoolCCTPV2FinalityConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2FinalityConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2FinalityConfigUpdated struct {
	FinalityConfig               uint16
	CustomFinalityTransferFeeBps uint16
	Raw                          types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterFinalityConfigUpdated(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2FinalityConfigUpdatedIterator, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "FinalityConfigUpdated")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2FinalityConfigUpdatedIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "FinalityConfigUpdated", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchFinalityConfigUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2FinalityConfigUpdated) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "FinalityConfigUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2FinalityConfigUpdated)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "FinalityConfigUpdated", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseFinalityConfigUpdated(log types.Log) (*USDCTokenPoolCCTPV2FinalityConfigUpdated, error) {
	event := new(USDCTokenPoolCCTPV2FinalityConfigUpdated)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "FinalityConfigUpdated", log); err != nil {
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

type USDCTokenPoolCCTPV2RateLimitAdminRoleGrantedIterator struct {
	Event *USDCTokenPoolCCTPV2RateLimitAdminRoleGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2RateLimitAdminRoleGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2RateLimitAdminRoleGranted)
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
		it.Event = new(USDCTokenPoolCCTPV2RateLimitAdminRoleGranted)
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

func (it *USDCTokenPoolCCTPV2RateLimitAdminRoleGrantedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2RateLimitAdminRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2RateLimitAdminRoleGranted struct {
	Account common.Address
	Raw     types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterRateLimitAdminRoleGranted(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2RateLimitAdminRoleGrantedIterator, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "RateLimitAdminRoleGranted")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2RateLimitAdminRoleGrantedIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "RateLimitAdminRoleGranted", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchRateLimitAdminRoleGranted(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2RateLimitAdminRoleGranted) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "RateLimitAdminRoleGranted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2RateLimitAdminRoleGranted)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "RateLimitAdminRoleGranted", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseRateLimitAdminRoleGranted(log types.Log) (*USDCTokenPoolCCTPV2RateLimitAdminRoleGranted, error) {
	event := new(USDCTokenPoolCCTPV2RateLimitAdminRoleGranted)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "RateLimitAdminRoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2RateLimitAdminRoleRevokedIterator struct {
	Event *USDCTokenPoolCCTPV2RateLimitAdminRoleRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2RateLimitAdminRoleRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2RateLimitAdminRoleRevoked)
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
		it.Event = new(USDCTokenPoolCCTPV2RateLimitAdminRoleRevoked)
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

func (it *USDCTokenPoolCCTPV2RateLimitAdminRoleRevokedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2RateLimitAdminRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2RateLimitAdminRoleRevoked struct {
	Account common.Address
	Raw     types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterRateLimitAdminRoleRevoked(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2RateLimitAdminRoleRevokedIterator, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "RateLimitAdminRoleRevoked")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2RateLimitAdminRoleRevokedIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "RateLimitAdminRoleRevoked", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchRateLimitAdminRoleRevoked(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2RateLimitAdminRoleRevoked) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "RateLimitAdminRoleRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2RateLimitAdminRoleRevoked)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "RateLimitAdminRoleRevoked", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseRateLimitAdminRoleRevoked(log types.Log) (*USDCTokenPoolCCTPV2RateLimitAdminRoleRevoked, error) {
	event := new(USDCTokenPoolCCTPV2RateLimitAdminRoleRevoked)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "RateLimitAdminRoleRevoked", log); err != nil {
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

type USDCTokenPoolCCTPV2RoleAdminChangedIterator struct {
	Event *USDCTokenPoolCCTPV2RoleAdminChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2RoleAdminChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2RoleAdminChanged)
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
		it.Event = new(USDCTokenPoolCCTPV2RoleAdminChanged)
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

func (it *USDCTokenPoolCCTPV2RoleAdminChangedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*USDCTokenPoolCCTPV2RoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2RoleAdminChangedIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2RoleAdminChanged)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseRoleAdminChanged(log types.Log) (*USDCTokenPoolCCTPV2RoleAdminChanged, error) {
	event := new(USDCTokenPoolCCTPV2RoleAdminChanged)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2RoleGrantedIterator struct {
	Event *USDCTokenPoolCCTPV2RoleGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2RoleGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2RoleGranted)
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
		it.Event = new(USDCTokenPoolCCTPV2RoleGranted)
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

func (it *USDCTokenPoolCCTPV2RoleGrantedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*USDCTokenPoolCCTPV2RoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2RoleGrantedIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2RoleGranted)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseRoleGranted(log types.Log) (*USDCTokenPoolCCTPV2RoleGranted, error) {
	event := new(USDCTokenPoolCCTPV2RoleGranted)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2RoleRevokedIterator struct {
	Event *USDCTokenPoolCCTPV2RoleRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2RoleRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2RoleRevoked)
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
		it.Event = new(USDCTokenPoolCCTPV2RoleRevoked)
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

func (it *USDCTokenPoolCCTPV2RoleRevokedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*USDCTokenPoolCCTPV2RoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2RoleRevokedIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2RoleRevoked)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseRoleRevoked(log types.Log) (*USDCTokenPoolCCTPV2RoleRevoked, error) {
	event := new(USDCTokenPoolCCTPV2RoleRevoked)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
type PendingDefaultAdmin struct {
	NewAdmin common.Address
	Schedule *big.Int
}
type PendingDefaultAdminDelay struct {
	NewDelay *big.Int
	Schedule *big.Int
}

func (USDCTokenPoolCCTPV2AllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (USDCTokenPoolCCTPV2AllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
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

func (USDCTokenPoolCCTPV2CustomFinalityOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b2")
}

func (USDCTokenPoolCCTPV2CustomFinalityTransferInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x41a8aa8df7945f0fb8ac5f7d88279638d9dc2ef9a6bf4ec9a53b80681b34aff7")
}

func (USDCTokenPoolCCTPV2DefaultAdminDelayChangeCanceled) Topic() common.Hash {
	return common.HexToHash("0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5")
}

func (USDCTokenPoolCCTPV2DefaultAdminDelayChangeScheduled) Topic() common.Hash {
	return common.HexToHash("0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b")
}

func (USDCTokenPoolCCTPV2DefaultAdminTransferCanceled) Topic() common.Hash {
	return common.HexToHash("0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109")
}

func (USDCTokenPoolCCTPV2DefaultAdminTransferScheduled) Topic() common.Hash {
	return common.HexToHash("0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6")
}

func (USDCTokenPoolCCTPV2DomainsSet) Topic() common.Hash {
	return common.HexToHash("0xc97f93e817584952f1c1d633f93784b8430f0633d002f9dcc4de4fe2780424d0")
}

func (USDCTokenPoolCCTPV2DynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f0238")
}

func (USDCTokenPoolCCTPV2FinalityConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x52aa194b292c8bfb5aaca8ee2000a965c3a051b306ff841873b16147526a39ba")
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

func (USDCTokenPoolCCTPV2PoolFeeWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b859599")
}

func (USDCTokenPoolCCTPV2RateLimitAdminRoleGranted) Topic() common.Hash {
	return common.HexToHash("0xf7af318a70f367e30346e2704f6ef646b378a7dcb49767beb98a1774cd11e389")
}

func (USDCTokenPoolCCTPV2RateLimitAdminRoleRevoked) Topic() common.Hash {
	return common.HexToHash("0xd63806009f622849e3b7cfd82d762420d57574c39f945f678871b2b5f1e8ce4b")
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

func (USDCTokenPoolCCTPV2RoleAdminChanged) Topic() common.Hash {
	return common.HexToHash("0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff")
}

func (USDCTokenPoolCCTPV2RoleGranted) Topic() common.Hash {
	return common.HexToHash("0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d")
}

func (USDCTokenPoolCCTPV2RoleRevoked) Topic() common.Hash {
	return common.HexToHash("0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b")
}

func (USDCTokenPoolCCTPV2TokenTransferFeeConfigDeleted) Topic() common.Hash {
	return common.HexToHash("0x5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8")
}

func (USDCTokenPoolCCTPV2TokenTransferFeeConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x56f77aeff2def50c8b5f5a0df3bab7183df09bf36c6feba496bb42551db77d70")
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2) Address() common.Address {
	return _USDCTokenPoolCCTPV2.address
}

type USDCTokenPoolCCTPV2Interface interface {
	AUTHORIZEDCALLERROLE(opts *bind.CallOpts) ([32]byte, error)

	DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error)

	FINALITYTHRESHOLD(opts *bind.CallOpts) (uint32, error)

	MAXFEE(opts *bind.CallOpts) (uint32, error)

	MINUSDCMESSAGELENGTH(opts *bind.CallOpts) (*big.Int, error)

	RATELIMITERADMINROLE(opts *bind.CallOpts) ([32]byte, error)

	DefaultAdmin(opts *bind.CallOpts) (common.Address, error)

	DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error)

	DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error)

	GetAccumulatedFees(opts *bind.CallOpts) (*big.Int, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetDomain(opts *bind.CallOpts, chainSelector uint64) (USDCTokenPoolDomain, error)

	GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

		error)

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRequiredCCVs(opts *bind.CallOpts, arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error)

	GetSupportedChains(opts *bind.CallOpts) ([]uint64, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	GetTokenDecimals(opts *bind.CallOpts) (uint8, error)

	GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error)

	HasRateLimitAdminRole(opts *bind.CallOpts, account common.Address) (bool, error)

	HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error)

	ILocalDomainIdentifier(opts *bind.CallOpts) (uint32, error)

	IMessageTransmitterProxy(opts *bind.CallOpts) (common.Address, error)

	ISupportedUSDCVersion(opts *bind.CallOpts) (uint32, error)

	ITokenMessenger(opts *bind.CallOpts) (common.Address, error)

	IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error)

	IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error)

	IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	PendingDefaultAdmin(opts *bind.CallOpts) (PendingDefaultAdmin,

		error)

	PendingDefaultAdminDelay(opts *bind.CallOpts) (PendingDefaultAdminDelay,

		error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error)

	AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error)

	ApplyCCVConfigUpdates(opts *bind.TransactOpts, ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error)

	ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error)

	ApplyFinalityConfigUpdates(opts *bind.TransactOpts, finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error)

	ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error)

	BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error)

	ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error)

	GrantRateLimitAdminRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error)

	GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	RevokeRateLimitAdminRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error)

	RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error)

	SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error)

	SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error)

	SetCustomFinalityRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error)

	SetDomains(opts *bind.TransactOpts, domains []USDCTokenPoolDomainUpdate) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error)

	WithdrawFees(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error)

	FilterAllowListAdd(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2AllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2AllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*USDCTokenPoolCCTPV2AllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2AllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2AllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*USDCTokenPoolCCTPV2AllowListRemove, error)

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

	FilterCustomFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2CustomFinalityOutboundRateLimitConsumedIterator, error)

	WatchCustomFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2CustomFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomFinalityOutboundRateLimitConsumed(log types.Log) (*USDCTokenPoolCCTPV2CustomFinalityOutboundRateLimitConsumed, error)

	FilterCustomFinalityTransferInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2CustomFinalityTransferInboundRateLimitConsumedIterator, error)

	WatchCustomFinalityTransferInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2CustomFinalityTransferInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomFinalityTransferInboundRateLimitConsumed(log types.Log) (*USDCTokenPoolCCTPV2CustomFinalityTransferInboundRateLimitConsumed, error)

	FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2DefaultAdminDelayChangeCanceledIterator, error)

	WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2DefaultAdminDelayChangeCanceled) (event.Subscription, error)

	ParseDefaultAdminDelayChangeCanceled(log types.Log) (*USDCTokenPoolCCTPV2DefaultAdminDelayChangeCanceled, error)

	FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2DefaultAdminDelayChangeScheduledIterator, error)

	WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2DefaultAdminDelayChangeScheduled) (event.Subscription, error)

	ParseDefaultAdminDelayChangeScheduled(log types.Log) (*USDCTokenPoolCCTPV2DefaultAdminDelayChangeScheduled, error)

	FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2DefaultAdminTransferCanceledIterator, error)

	WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2DefaultAdminTransferCanceled) (event.Subscription, error)

	ParseDefaultAdminTransferCanceled(log types.Log) (*USDCTokenPoolCCTPV2DefaultAdminTransferCanceled, error)

	FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*USDCTokenPoolCCTPV2DefaultAdminTransferScheduledIterator, error)

	WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2DefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error)

	ParseDefaultAdminTransferScheduled(log types.Log) (*USDCTokenPoolCCTPV2DefaultAdminTransferScheduled, error)

	FilterDomainsSet(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2DomainsSetIterator, error)

	WatchDomainsSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2DomainsSet) (event.Subscription, error)

	ParseDomainsSet(log types.Log) (*USDCTokenPoolCCTPV2DomainsSet, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2DynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2DynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*USDCTokenPoolCCTPV2DynamicConfigSet, error)

	FilterFinalityConfigUpdated(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2FinalityConfigUpdatedIterator, error)

	WatchFinalityConfigUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2FinalityConfigUpdated) (event.Subscription, error)

	ParseFinalityConfigUpdated(log types.Log) (*USDCTokenPoolCCTPV2FinalityConfigUpdated, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2InboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2InboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*USDCTokenPoolCCTPV2InboundRateLimitConsumed, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2LockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2LockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*USDCTokenPoolCCTPV2LockedOrBurned, error)

	FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2OutboundRateLimitConsumedIterator, error)

	WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2OutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseOutboundRateLimitConsumed(log types.Log) (*USDCTokenPoolCCTPV2OutboundRateLimitConsumed, error)

	FilterPoolFeeWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*USDCTokenPoolCCTPV2PoolFeeWithdrawnIterator, error)

	WatchPoolFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2PoolFeeWithdrawn, recipient []common.Address) (event.Subscription, error)

	ParsePoolFeeWithdrawn(log types.Log) (*USDCTokenPoolCCTPV2PoolFeeWithdrawn, error)

	FilterRateLimitAdminRoleGranted(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2RateLimitAdminRoleGrantedIterator, error)

	WatchRateLimitAdminRoleGranted(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2RateLimitAdminRoleGranted) (event.Subscription, error)

	ParseRateLimitAdminRoleGranted(log types.Log) (*USDCTokenPoolCCTPV2RateLimitAdminRoleGranted, error)

	FilterRateLimitAdminRoleRevoked(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2RateLimitAdminRoleRevokedIterator, error)

	WatchRateLimitAdminRoleRevoked(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2RateLimitAdminRoleRevoked) (event.Subscription, error)

	ParseRateLimitAdminRoleRevoked(log types.Log) (*USDCTokenPoolCCTPV2RateLimitAdminRoleRevoked, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2ReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2ReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*USDCTokenPoolCCTPV2ReleasedOrMinted, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2RemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2RemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*USDCTokenPoolCCTPV2RemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2RemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2RemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*USDCTokenPoolCCTPV2RemotePoolRemoved, error)

	FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*USDCTokenPoolCCTPV2RoleAdminChangedIterator, error)

	WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error)

	ParseRoleAdminChanged(log types.Log) (*USDCTokenPoolCCTPV2RoleAdminChanged, error)

	FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*USDCTokenPoolCCTPV2RoleGrantedIterator, error)

	WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleGranted(log types.Log) (*USDCTokenPoolCCTPV2RoleGranted, error)

	FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*USDCTokenPoolCCTPV2RoleRevokedIterator, error)

	WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleRevoked(log types.Log) (*USDCTokenPoolCCTPV2RoleRevoked, error)

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*USDCTokenPoolCCTPV2TokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2TokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*USDCTokenPoolCCTPV2TokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*USDCTokenPoolCCTPV2TokenTransferFeeConfigUpdatedIterator, error)

	WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2TokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigUpdated(log types.Log) (*USDCTokenPoolCCTPV2TokenTransferFeeConfigUpdated, error)

	Address() common.Address
}
