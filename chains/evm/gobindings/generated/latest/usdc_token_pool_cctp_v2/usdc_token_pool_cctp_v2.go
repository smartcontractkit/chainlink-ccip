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
	RemoteChainSelector uint64
	OutboundCCVs        []common.Address
	InboundCCVs         []common.Address
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"internalType\":\"contractITokenMessenger\"},{\"name\":\"cctpMessageTransmitterProxy\",\"type\":\"address\",\"internalType\":\"contractCCTPMessageTransmitterProxy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"AUTHORIZED_CALLER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FINALITY_THRESHOLD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_FEE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_USDC_MESSAGE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RATE_LIMITER_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyFinalityConfigUpdates\",\"inputs\":[{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"destToUseDefaultFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beginDefaultAdminTransfer\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeDefaultAdminDelay\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"defaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelayIncreaseWait\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAccumulatedFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDomain\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structUSDCTokenPool.Domain\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useLegacySourcePoolDataFormat\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredInboundCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredOutboundCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRateLimitAdminRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRateLimitAdminRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_localDomainIdentifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_messageTransmitterProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractCCTPMessageTransmitterProxy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_supportedUSDCVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_tokenMessenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITokenMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRateLimitAdminRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rollbackDefaultAdminDelay\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustomFinalityRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDomains\",\"inputs\":[{\"name\":\"domains\",\"type\":\"tuple[]\",\"internalType\":\"structUSDCTokenPool.DomainUpdate[]\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useLegacySourcePoolDataFormat\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRouter\",\"inputs\":[{\"name\":\"newRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityTransferInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeScheduled\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"effectSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferScheduled\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"acceptSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DomainsSet\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structUSDCTokenPool.DomainUpdate[]\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useLegacySourcePoolDataFormat\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalityConfigUpdated\",\"inputs\":[{\"name\":\"finalityConfig\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeeWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminRoleGranted\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminRoleRevoked\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RouterUpdated\",\"inputs\":[{\"name\":\"oldRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlInvalidDefaultAdmin\",\"inputs\":[{\"name\":\"defaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidBurnToken\",\"inputs\":[{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"got\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDepositHash\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"got\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidDestinationDomain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidDomain\",\"inputs\":[{\"name\":\"domain\",\"type\":\"tuple\",\"internalType\":\"structUSDCTokenPool.DomainUpdate\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useLegacySourcePoolDataFormat\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"type\":\"error\",\"name\":\"InvalidExecutionFinalityThreshold\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidFinality\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidFinalityConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMessageLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMinFinalityThreshold\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidNonce\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"got\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidPreviousPool\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourceDomain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenMessengerVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidTransmitterInProxy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenTransferFeeConfigNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnknownDomain\",\"inputs\":[{\"name\":\"domain\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"UnlockingUSDCFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x610180806040523461057b5761825d803803809161001d8285610937565b8339810160c08282031261057b5781516001600160a01b0381169283820361057b576020810151926001600160a01b0384169182850361057b5760408101516001600160a01b038116919082810361057b5760608201516001600160401b03811161057b5782019380601f8601121561057b578451946001600160401b038611610580578560051b9060208201966100b86040519889610937565b875260208088019282010192831161057b57602001905b82821061091f575050506100f160a06100ea6080850161095a565b930161095a565b92331561090957600180546001600160d01b031690556002546001600160a01b0381166108f8576001600160a01b03191633908117600255610132906109b4565b50801580156108e7575b80156108d6575b6108c55760049260209260805260c0526040519283809263313ce56760e01b82525afa8091600091610889575b5090610865575b50600660a052600580546001600160a01b0319166001600160a01b03929092169190911790558051151560e081905261073c575b50600161010052831561072b57604051632c12192160e01b8152602081600481885afa9081156105eb576000916106f1575b5060405163054fd4d560e41b81526001600160a01b03919091169190602081600481865afa9081156105eb576000916106d2575b5063ffffffff8061010051169116908082036106bb575050604051639cdbb18160e01b8152602081600481895afa9081156105eb5760009161069c575b5063ffffffff8061010051169116908082036106855750506020600491604051928380926367e0ed8360e11b82525afa80156105eb578291600091610637575b506001600160a01b0316036106265760049260209261012052610140526040519283809263234d8e3d60e21b82525afa9081156105eb576000916105f7575b506101605260805161012051604051636eb1769f60e11b81523060048201526001600160a01b0391821660248201819052939290911690602081604481855afa9081156105eb576000916105b9575b5060001981018091116105a35760405190602082019463095ea7b360e01b86526024830152604482015260448152610355606482610937565b6000806040958651936103688886610937565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082865af13d15610596573d906001600160401b0382116105805785516103d99490926103ca601f8201601f191660200185610937565b83523d6000602085013e610b9c565b8051806104ff575b837f2e902d38f15b233cbb63711add0fca4545334d3a169d60c0a616494d7eea95446020858351908152a1516175d09081610c6d82396080518181816108d701528181610ac501528181610b1601528181611258015281816121ff015281816129b2015281816142fe01528181616c690152616fe5015260a051818181610ba401528181615acc01528181615b5c01526162f2015260c05181818161331b015281816153a20152818161552c01528181615eb50152615fb4015260e0518181816117ad015281816134210152616f29015261010051818181610a8101526157af01526101205181818161195801526121bd0152610140518181816111cd01526120a8015261016051818181611ccd015281816122ee01526158350152f35b816020918101031261057b576020015180159081150361057b576105245738806103e1565b815162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b916103d992606091610b9c565b634e487b7160e01b600052601160045260246000fd5b90506020813d6020116105e3575b816105d460209383610937565b8101031261057b57513861031c565b3d91506105c7565b6040513d6000823e3d90fd5b610619915060203d60201161061f575b6106118183610937565b81019061096e565b386102cd565b503d610607565b632a32133b60e11b60005260046000fd5b9091506020813d60201161067d575b8161065360209383610937565b810103126106795751906001600160a01b0382168203610676575081903861028e565b80fd5b5080fd5b3d9150610646565b633785f8f160e01b60005260045260245260446000fd5b6106b5915060203d60201161061f576106118183610937565b3861024e565b63960693cd60e01b60005260045260245260446000fd5b6106eb915060203d60201161061f576106118183610937565b38610211565b90506020813d602011610723575b8161070c60209383610937565b8101031261057b5761071d9061095a565b386101dd565b3d91506106ff565b6306b7c75960e31b60005260046000fd5b919290602092604051926107508585610937565b60008452600036813760e051156108545760005b84518110156107cb576001906001600160a01b03610782828861098a565b51168761078e82610a5a565b61079b575b505001610764565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a13887610793565b50919490925060005b8351811015610848576001906001600160a01b036107f2828761098a565b51168015610842578661080482610b42565b610812575b50505b016107d4565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a13886610809565b5061080c565b509250929050386101ab565b6335f4a7b360e01b60005260046000fd5b60ff1660068114610177576332ad3e0760e11b600052600660045260245260446000fd5b6020813d6020116108bd575b816108a260209383610937565b8101031261067957519060ff82168203610676575038610170565b3d9150610895565b630a64406560e11b60005260046000fd5b506001600160a01b03831615610143565b506001600160a01b0384161561013c565b631fe1e13d60e11b60005260046000fd5b636116401160e11b600052600060045260246000fd5b6020809161092c8461095a565b8152019101906100cf565b601f909101601f19168101906001600160401b0382119082101761058057604052565b51906001600160a01b038216820361057b57565b9081602091031261057b575163ffffffff8116810361057b5790565b805182101561099e5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b6001600160a01b038116600090815260008051602061823d833981519152602052604090205460ff16610a3c576001600160a01b0316600081815260008051602061823d83398151915260205260408120805460ff191660011790553391907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b50600090565b805482101561099e5760005260206000200190600090565b6000818152600460205260409020548015610b3b5760001981018181116105a3576003546000198101919082116105a357818103610aea575b5050506003548015610ad45760001901610aae816003610a42565b8154906000199060031b1b19169055600355600052600460205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b610b23610afb610b0c936003610a42565b90549060031b1c9283926003610a42565b819391549060031b91821b91600019901b19161790565b90556000526004602052604060002055388080610a93565b5050600090565b80600052600460205260406000205415600014610a3c576003546801000000000000000081101561058057610b83610b0c8260018594016003556003610a42565b9055600354906000526004602052604060002055600190565b91929015610bfe5750815115610bb0575090565b3b15610bb95790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b825190915015610c115750805190602001fd5b6040519062461bcd60e51b8252602060048301528181519182602483015260005b838110610c545750508160006044809484010152601f80199101168101030190fd5b60208282018101516044878401015285935001610c3256fe6080604052600436101561001257600080fd5b60003560e01c806301ffc9a71461043d578063022d63fb146104385780630aa6220b146104335780630bd7c46d1461042e578063164e68de14610429578063181f5a7714610424578063212a052e1461041f57806321df0da71461041a578063240028e814610415578063248a9ca31461041057806324f65ee71461040b5780632a10097b146104065780632c286daf146104015780632f2ff15d146103fc57806336568abe146103f757806337b19247146103f257806339077537146103ed578063489a68f2146103e85780634ac8bd5f146103e35780634c5ef0ed146103de5780634f71592c146103d957806354c8a4f3146103d45780635df45a37146103cf5780636155cda0146103ca57806362ddd3c4146103c5578063634e93da146103c0578063649a5ec7146103bb5780636b716b0d146103b6578063791e5a10146103b1578063804ba5a9146103ac57806384ef8ffc146103a25780638926f54f146103a75780638da5cb5b146103a257806391d148541461039d578063962d40201461039857806398db9643146103935780639a4575b91461038e5780639f68f67314610389578063a1eda53c14610384578063a217fddf14610357578063a42a7b8b1461037f578063a7cd63b71461037a578063acfecf9114610375578063af58d59f14610370578063b0f479a11461036b578063b1c71c6514610366578063b79019b514610361578063b79465801461035c578063bc063e1a14610357578063c0d7865514610352578063c4bffe2b1461034d578063c75eea9c14610348578063c8c8fd1914610343578063cc8463c81461033e578063cefc142914610339578063cf6eefb714610334578063cf7401f31461032f578063d547741f1461032a578063d602b9fd14610325578063da4b05e714610320578063da90a9f31461031b578063dc0bd97114610316578063dfadfa3514610311578063e0351e131461030c578063e58d80c714610307578063e8a1da17146103025763f573388e146102fd57600080fd5b613901565b6134af565b613446565b613409565b61333f565b6132fb565b61327a565b61325d565b6131e2565b613198565b6130e7565b612f57565b612e5e565b612e33565b612e16565b612dcb565b612d56565b612c1f565b6124fd565b612be8565b612a75565b61291a565b6128d7565b61282b565b6126d5565b61267b565b612597565b61249a565b612461565b612107565b612088565b611e8b565b611e03565b611d9d565b611dc4565b611d2c565b611cf1565b611cb0565b611bb8565b611a67565b6119b5565b611938565b61191d565b61174b565b6116bb565b6115d3565b6113dc565b611310565b6110f5565b611021565b610e96565b610e1d565b610cdf565b610bf9565b610b8a565b610b55565b610ae9565b610aa5565b610a64565b610a05565b6107df565b610721565b610643565b610625565b34610615576020600319360112610615576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361061557807ff208a58f000000000000000000000000000000000000000000000000000000006104d992149081156105eb575b81156105c1575b8115610597575b811561056d575b81156104dd575b5060405190151581529081906020820190565b0390f35b7f3149878600000000000000000000000000000000000000000000000000000000811491508115610510575b50386104c6565b7f7965db0b00000000000000000000000000000000000000000000000000000000811491508115610543575b5038610509565b7f01ffc9a7000000000000000000000000000000000000000000000000000000009150143861053c565b7f01ffc9a700000000000000000000000000000000000000000000000000000000811491506104bf565b7f0e64dd2900000000000000000000000000000000000000000000000000000000811491506104b8565b7f1ef5498f00000000000000000000000000000000000000000000000000000000811491506104b1565b7faff2afbf00000000000000000000000000000000000000000000000000000000811491506104aa565b600080fd5b600091031261061557565b34610615576000600319360112610615576020604051620697808152f35b346106155760006003193601126106155761065c614d6e565b6002548060d01c8061067c575b600280546001600160a01b03169055005b005b4211156106e55765ffffffffffff6106de9160a01c1679ffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffffffff00000000000000000000000000000000000000000000000000006001549260d01b16911617600155565b3880610669565b507f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5600080a16106de565b6001600160a01b0381160361061557565b346106155760206003193601126106155760043561073e81610710565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff16156107b5576020816107a27ff7af318a70f367e30346e2704f6ef646b378a7dcb49767beb98a1774cd11e38993614e41565b506001600160a01b0360405191168152a1005b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b34610615576020600319360112610615576004356107fc81610710565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff16156107b55761083b6142c2565b908161084357005b6001600160a01b037f738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b8595999161098260405161096760208201917fa9059cbb0000000000000000000000000000000000000000000000000000000083526108d4816108c68a8860248401602090939291936001600160a01b0360408201951681520152565b03601f198101835282611515565b857f00000000000000000000000000000000000000000000000000000000000000001660008060409586519461090a8887611515565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020870152519082855af13d156109a8573d9161094b83611565565b9261095887519485611515565b83523d6000602085013e6174ff565b805180610987575b5050519283921694829190602083019252565b0390a2005b8160208061099c936109a19501019101613dce565b616832565b388061096f565b6060916174ff565b919082519283825260005b8481106109dc575050601f19601f8460006020809697860101520116010190565b806020809284010151828286010152016109bb565b906020610a029281815201906109b0565b90565b34610615576000600319360112610615576104d96040805190610a288183611515565b601d82527f55534443546f6b656e506f6f6c43435450563220312e372e302d6465760000006020830152519182916020835260208301906109b0565b3461061557600060031936011261061557602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346106155760006003193601126106155760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610615576020600319360112610615576020610b3c600435610b0b81610710565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691161490565b6040519015158152f35b90816020910312610615573590565b34610615576020600319360112610615576020610b82600435600052600060205260016040600020015490565b604051908152f35b3461061557600060031936011261061557602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b9181601f840112156106155782359167ffffffffffffffff8311610615576020808501948460051b01011161061557565b346106155760406003193601126106155760043567ffffffffffffffff811161061557366023820112156106155780600401359067ffffffffffffffff82116106155736602460a0840283010111610615576024359067ffffffffffffffff82116106155761067a92610c726024933690600401610bc8565b9390920161393c565b6004359061ffff8216820361061557565b6024359061ffff8216820361061557565b6064359061ffff8216820361061557565b9181601f840112156106155782359167ffffffffffffffff83116106155760208085019460e0850201011161061557565b3461061557606060031936011261061557610cf8610c7b565b610d00610c8c565b60443567ffffffffffffffff811161061557610d20903690600401610cae565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb560205260409020549093919060ff16156107b55761ffff8316612710811015610df057507f52aa194b292c8bfb5aaca8ee2000a965c3a051b306ff841873b16147526a39ba93610dd191600a5463ffff00008660101b16907fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000061ffff871691161717600a55614f2d565b6040805161ffff9283168152929091166020830152819081015b0390a1005b7f95f3517a0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b3461061557604060031936011261061557600435602435610e3d81610710565b8115610e6c5781610e67610e6261067a94600052600060205260016040600020015490565b614dda565b614eb9565b7f3fc3c27a0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461061557604060031936011261061557600435602435610eb681610710565b811580610f94575b610ecc575b61067a91615299565b60015465ffffffffffff60a082901c1692906001600160a01b031615801590610f80575b8015610f6e575b610f345761067a9250610f2d7fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff60015416600155565b9150610ec3565b7f19ca5ebb0000000000000000000000000000000000000000000000000000000060005265ffffffffffff831660045260246000fd5b6000fd5b504265ffffffffffff84161015610ef7565b5065ffffffffffff831615610ef0565b1590565b50610fb6610faa6001600160a01b036002541690565b6001600160a01b031690565b6001600160a01b03821614610ebe565b67ffffffffffffffff81160361061557565b3590610fe382610fc6565b565b908160a09103126106155790565b9181601f840112156106155782359167ffffffffffffffff8311610615576020838186019501011161061557565b346106155760a06003193601126106155761103d600435610710565b60243561104981610fc6565b60443567ffffffffffffffff811161061557611069903690600401610fe5565b50611072610c9d565b5060843567ffffffffffffffff8111610615576104d99161109a6110a1923690600401610ff3565b5050613c73565b60405191829182919091606080608083019463ffffffff815116845263ffffffff602082015116602085015263ffffffff604082015116604085015201511515910152565b90816101009103126106155790565b346106155760206003193601126106155760043567ffffffffffffffff8111610615576111269036906004016110e6565b61112e613cf1565b50606081013561113e81836152e0565b6111bf602061115b61115360e0860186613d04565b810190613d55565b61118461117d61117861117160c0890189613d04565b3691611581565b6156bb565b825161579c565b8181519101519060405193849283927f57ecfd2800000000000000000000000000000000000000000000000000000000845260048401613de3565b038160006001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165af190811561130b576000916112dc575b50156112b257817ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc067ffffffffffffffff61124a604061124360206104d99801613ac8565b9401613e14565b604080516001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081168252336020830152929092169082015260608101859052921691608090a261129f611538565b8190526040519081529081906020820190565b7fbf969f220000000000000000000000000000000000000000000000000000000060005260046000fd5b6112fe915060203d602011611304575b6112f68183611515565b810190613dce565b386111fe565b503d6112ec565b613e08565b346106155760406003193601126106155760043567ffffffffffffffff8111610615576113446104d99136906004016110e6565b61134c610c8c565b90600060405161135b81611468565b5261138d611385606083013561137f61137a61117160c0870187613d04565b615a23565b90615b59565b9283836154ee565b7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc067ffffffffffffffff61124a602060408501946113cb8635610710565b0135936113d785610fc6565b613e14565b346106155760206003193601126106155760043567ffffffffffffffff8111610615573660238201121561061557806004013567ffffffffffffffff81116106155736602460c083028401011161061557602461067a9201613e1e565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff82111761148457604052565b611439565b6060810190811067ffffffffffffffff82111761148457604052565b60a0810190811067ffffffffffffffff82111761148457604052565b6080810190811067ffffffffffffffff82111761148457604052565b6040810190811067ffffffffffffffff82111761148457604052565b60c0810190811067ffffffffffffffff82111761148457604052565b90601f601f19910116810190811067ffffffffffffffff82111761148457604052565b60405190610fe3602083611515565b60405190610fe3604083611515565b60405190610fe360a083611515565b67ffffffffffffffff811161148457601f01601f191660200190565b92919261158d82611565565b9161159b6040519384611515565b829481845281830111610615578281602093846000960137010152565b9080601f8301121561061557816020610a0293359101611581565b34610615576040600319360112610615576004356115f081610fc6565b60243567ffffffffffffffff811161061557602091611616610b3c9236906004016115b8565b906141b4565b60a06003198201126106155760043561163481610710565b9160243561164181610fc6565b916044359160643561ffff8116810361061557916084359067ffffffffffffffff82116106155761167491600401610ff3565b9091565b602060408183019282815284518094520192019060005b81811061169c5750505090565b82516001600160a01b031684526020938401939092019160010161168f565b346106155767ffffffffffffffff6116d23661161c565b50505050905016600052600d6020526104d96116f160406000206141f1565b60405191829182611678565b60406003198201126106155760043567ffffffffffffffff8111610615578161172891600401610bc8565b929092916024359067ffffffffffffffff82116106155761167491600401610bc8565b3461061557611759366116fd565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb560205260409020549093929060ff16156107b5576117ab926117a391369161425d565b92369161425d565b7f0000000000000000000000000000000000000000000000000000000000000000156118f35760005b825181101561185d57806117fa6117ed60019386614420565b516001600160a01b031690565b61181461180f6001600160a01b038316610faa565b617112565b611820575b50016117d4565b6040516001600160a01b039190911681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756690602090a138611819565b5060005b815181101561067a578061187a6117ed60019385614420565b6001600160a01b038116156118ed576118a361189e6001600160a01b038316610faa565b6173b9565b6118b0575b505b01611861565b6040516001600160a01b039190911681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d890602090a1836118a8565b506118aa565b7f35f4a7b30000000000000000000000000000000000000000000000000000000060005260046000fd5b34610615576000600319360112610615576020610b826142c2565b346106155760006003193601126106155760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b9060406003198301126106155760043561199581610fc6565b916024359067ffffffffffffffff82116106155761167491600401610ff3565b34610615576119c33661197c565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205491929160ff16156107b55767ffffffffffffffff8216611a1f816000526007602052604060002054151590565b15611a3a575061067a92611a34913691611581565b90615c67565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b3461061557602060031936011261061557600435611a8481610710565b611a8c614d6e565b7f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed66020611ac9611abb42616edd565b611ac36148ac565b90615d27565b65ffffffffffff6001600160a01b03611af86001549065ffffffffffff6001600160a01b0383169260a01c1690565b9690501694857fffffffffffffffffffffffff00000000000000000000000000000000000000006001541617600155611b77837fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff79ffffffffffff00000000000000000000000000000000000000006001549260a01b16911617600155565b16611b8e575b65ffffffffffff60405191168152a2005b7f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109600080a1611b7d565b346106155760206003193601126106155760043565ffffffffffff81169081810361061557611be5614d6e565b611bee42616edd565b9165ffffffffffff611bfe6148ac565b1680821115611c7157507ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b9265ffffffffffff611c3d611c44936174d7565b1690615d27565b90611c4f828261662d565b6040805165ffffffffffff928316815292909116602083015281908101610deb565b0365ffffffffffff8111611cab577ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b92611c449190615d27565b6146ec565b3461061557600060031936011261061557602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346106155760006003193601126106155760206040517f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af8152f35b346106155760206003193601126106155760043567ffffffffffffffff811161061557611d5d903690600401610cae565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff16156107b55761067a91614f2d565b346106155760006003193601126106155760206001600160a01b0360025416604051908152f35b34610615576020600319360112610615576020610b3c67ffffffffffffffff600435611def81610fc6565b166000526007602052604060002054151590565b3461061557604060031936011261061557602060ff611e4e602435600435611e2a82610710565b600052600084526040600020906001600160a01b0316600052602052604060002090565b54166040519015158152f35b9181601f840112156106155782359167ffffffffffffffff8311610615576020808501946060850201011161061557565b346106155760606003193601126106155760043567ffffffffffffffff811161061557611ebc903690600401610bc8565b9060243567ffffffffffffffff811161061557611edd903690600401611e5a565b9060443567ffffffffffffffff811161061557611efe903690600401611e5a565b7f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af6000908152602052611f72610f90611f6b337f135a184f1ea19f46a2a5eda150dca2e4f81e2266dd73bb074c2f2a4e0c0032a05b906001600160a01b0316600052602052604060002090565b5460ff1690565b8061204a575b61201c57838614801590612012575b611fe85760005b868110611f9757005b80611fe2611fb0611fab6001948b8b613c63565b613ac8565b611fbb83898961437c565b611fdc611fd4611fcc86898b61437c565b92369061309e565b91369061309e565b91615d45565b01611f8e565b7f568efce20000000000000000000000000000000000000000000000000000000060005260046000fd5b5080861415611f87565b7f8e4a23d6000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b506000808052602052612083610f90611f6b337fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5611f53565b611f78565b346106155760006003193601126106155760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b610a029160206120e583516040845260408401906109b0565b9201519060208184039101526109b0565b906020610a029281815201906120cc565b346106155760206003193601126106155760043567ffffffffffffffff811161061557612138903690600401610fe5565b61214061438c565b5061214a81615e75565b6020810161217c61217761215d83613ac8565b67ffffffffffffffff16600052600f602052604060002090565b6143a5565b9161218d610f906060850151151590565b61241f57602061219d8280613d04565b9050036123db57602083015180156123bf57925b60606001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169201359260408201926121f4845163ffffffff1690565b926001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016938151833b15610615576040517fd04857b00000000000000000000000000000000000000000000000000000000081526004810189905263ffffffff929092166024830152604482018990526001600160a01b03861660648301526084820152600060a482018190526107d060c4830152909492859060e490829084905af1801561130b57612343612387966123227ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1094866104d99c6123829a67ffffffffffffffff976123a4575b506123187f0000000000000000000000000000000000000000000000000000000000000000955163ffffffff1690565b9251928d86616113565b61233961232d611547565b63ffffffff9093168352565b60208201526161be565b9661237a61235086613ac8565b604080516001600160a01b0390971687523360208801528601929092529116929081906060820190565b0390a2613ac8565b61488a565b90612390611547565b9182526020820152604051918291826120f6565b806123b360006123b993611515565b8061061a565b386122e8565b506123d56123cd8280613d04565b810190610b46565b926121b1565b806123e591613d04565b9061241b6040519283927fa3c8cf090000000000000000000000000000000000000000000000000000000084526004840161440f565b0390fd5b610f6a61242b83613ac8565b7fd201c48a0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b346106155767ffffffffffffffff6124783661161c565b50505050905016600052600d6020526104d96116f160016040600020016141f1565b34610615576000600319360112610615576002548060d01c90811515806124f3575b156124e95760a01c65ffffffffffff165b6040805165ffffffffffff928316815292909116602083015290f35b50506000806124cd565b50428210156124bc565b3461061557600060031936011261061557602060405160008152f35b602081016020825282518091526040820191602060408360051b8301019401926000915b83831061254c57505050505090565b9091929394602080612588837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0866001960301875289516109b0565b9701930193019193929061253d565b346106155760206003193601126106155767ffffffffffffffff6004356125bd81610fc6565b1660005260086020526125d66005604060002001617058565b805190601f196125fe6125e884614245565b936125f66040519586611515565b808552614245565b0160005b81811061266a57505060005b815181101561265c578061264061263b61262a60019486614420565b516000526009602052604060002090565b614487565b61264a8286614420565b526126558185614420565b500161260e565b604051806104d98582612519565b806060602080938701015201612602565b34610615576000600319360112610615576040516003548082526020820190600360005260206000209060005b8181106126bf576104d9856116f181870382611515565b82548452602090930192600192830192016126a8565b34610615576126e33661197c565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205491929160ff16156107b55767ffffffffffffffff821691612743610f90846000526007602052604060002054151590565b6127f457612786610f90600561276d8467ffffffffffffffff166000526008602052604060002090565b01612779368689611581565b602081519101209061723e565b6127bd57507f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d7691926109826040519283928361440f565b61241b84926040519384937f74f23c7c00000000000000000000000000000000000000000000000000000000855260048501614547565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b346106155760206003193601126106155767ffffffffffffffff60043561285181610fc6565b612859614568565b501660005260086020526104d961287e6128796002604060002001614593565b616249565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b346106155760006003193601126106155760206001600160a01b0360055416604051908152f35b9291906129156020916040865260408601906120cc565b930152565b346106155760606003193601126106155760043567ffffffffffffffff81116106155761294b903690600401610fe5565b612953610c8c565b9060443567ffffffffffffffff8111610615576129749036906004016115b8565b5061297d61438c565b506129888282615f77565b606081013561ffff819316612a49575b506123826020612a2292016129ac81613ac8565b604080517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815233602082015290810186905267ffffffffffffffff91909116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1090806060810161237a565b612a2a6162eb565b612a32611547565b91825260208201526104d9604051928392836128fe565b909150612710612a6261ffff600a5460101c168361471b565b048103908111611cab5790612382612998565b346106155760206003193601126106155760043567ffffffffffffffff811161061557612aa6903690600401610bc8565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205490919060ff16156107b55760005b828110612aec57005b80612be2612afd60019386866145ed565b35612b0781610fc6565b612bdd612b22612b188589896145ed565b602081019061462d565b92907fb0897119e8510f887b892cbc4c8506fc51d9849fd90afae4fd065e705f2d0f6c612b5d612b53888c8c6145ed565b604081019061462d565b95612b71612b6c36838761425d565b616333565b612b7f612b6c36898561425d565b612bc0612b8a611547565b97612b9636848861425d565b8952612ba336828661425d565b60208a015260405193849367ffffffffffffffff891697856146c5565b0390a267ffffffffffffffff16600052600d602052604060002090565b6147ea565b01612ae3565b34610615576020600319360112610615576104d9612c0b60043561238281610fc6565b6040519182916020835260208301906109b0565b3461061557602060031936011261061557600435612c3c81610710565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff16156107b5576001600160a01b03168015612ce85760407f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f16849160055490807fffffffffffffffffffffffff00000000000000000000000000000000000000008316176005556001600160a01b038351921682526020820152a1005b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b602060408183019282815284518094520192019060005b818110612d365750505090565b825167ffffffffffffffff16845260209384019390920191600101612d29565b3461061557600060031936011261061557612d6f61700d565b805190601f19612d816125e884614245565b0136602084013760005b8151811015612dbd578067ffffffffffffffff612daa60019385614420565b5116612db68286614420565b5201612d8b565b604051806104d98582612d12565b346106155760206003193601126106155767ffffffffffffffff600435612df181610fc6565b612df9614568565b501660005260086020526104d961287e6128796040600020614593565b346106155760006003193601126106155760206040516101188152f35b34610615576000600319360112610615576020612e4e6148ac565b65ffffffffffff60405191168152f35b34610615576000600319360112610615576001546001600160a01b03163303612f295760015460a081901c65ffffffffffff16906001600160a01b031681158015612f1f575b612ef157612ec690612ec06001600160a01b03600254166163d7565b50614e6b565b50600180547fffffffffffff0000000000000000000000000000000000000000000000000000169055005b507f19ca5ebb0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b5042821015612ea4565b7fc22c8022000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b3461061557600060031936011261061557604065ffffffffffff612f916001549065ffffffffffff6001600160a01b0383169260a01c1690565b6001600160a01b03849392935193168352166020820152f35b8015150361061557565b6fffffffffffffffffffffffffffffffff81160361061557565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc6060910112610615576040519061300582611489565b8160243561301281612faa565b815260443561302081612fb4565b602082015260406064359161303483612fb4565b0152565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c6060910112610615576040519061306f82611489565b8160843561307c81612faa565b815260a43561308a81612fb4565b6020820152604060c4359161303483612fb4565b9190826060910312610615576040516130b681611489565b604080829480356130c681612faa565b845260208101356130d681612fb4565b602085015201359161303483612fb4565b346106155760e06003193601126106155760043561310481610fc6565b61310d36612fce565b61311636613038565b3360009081527f135a184f1ea19f46a2a5eda150dca2e4f81e2266dd73bb074c2f2a4e0c0032a0602052604090205490919060ff16158061315f575b61201c5761067a92615d45565b503360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff1615613152565b34610615576040600319360112610615576004356024356131b881610710565b8115610e6c57816131dd610e6261067a94600052600060205260016040600020015490565b616455565b34610615576000600319360112610615576131fb614d6e565b600180547fffffffffffff0000000000000000000000000000000000000000000000000000811690915560a01c65ffffffffffff1661323657005b7f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109600080a1005b346106155760006003193601126106155760206040516107d08152f35b346106155760206003193601126106155760043561329781610710565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff16156107b5576020816107a27fd63806009f622849e3b7cfd82d762420d57574c39f945f678871b2b5f1e8ce4b9361642b565b346106155760006003193601126106155760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b346106155760206003193601126106155767ffffffffffffffff60043561336581610fc6565b61336d614568565b5016600052600f6020526104d9604060002060ff600260405192613390846114a5565b8054845260018101546020850152015463ffffffff81166040840152818160201c161515606084015260281c16151560808201526040519182918291909160808060a0830194805184526020810151602085015263ffffffff604082015116604085015260608101511515606085015201511515910152565b346106155760006003193601126106155760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b3461061557602060031936011261061557602060ff611e4e60043561346a81610710565b7f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af600052600084526040600020906001600160a01b0316600052602052604060002090565b34610615576134bd366116fd565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb56020526040902054909391929060ff16156107b55783916000915b8083106137ad5750505060009163ffffffff4216925b82811061351e57005b61353161352c8285856149aa565b614a69565b906060820161354081516164ec565b608083019361354f85516164ec565b604084019081515115612ce857613589610f90613584613577885167ffffffffffffffff1690565b67ffffffffffffffff1690565b61742c565b613762576136c26135c26135a8879a999a5167ffffffffffffffff1690565b67ffffffffffffffff166000526008602052604060002090565b6136858961367f87516136666135eb60408301516fffffffffffffffffffffffffffffffff1690565b9161364d61361661360f60208401516fffffffffffffffffffffffffffffffff1690565b9251151590565b613644613621611556565b6fffffffffffffffffffffffffffffffff851681529763ffffffff166020890152565b15156040870152565b6fffffffffffffffffffffffffffffffff166060850152565b6fffffffffffffffffffffffffffffffff166080830152565b82614b00565b6136b7896136ae8a516136666135eb60408301516fffffffffffffffffffffffffffffffff1690565b60028301614b00565b600484519101614c0c565b602085019660005b8851805182101561370557906136ff6001926136f8836136f28c5167ffffffffffffffff1690565b92614420565b5190615c67565b016136ca565b505097965094906137597f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c293926137466001975167ffffffffffffffff1690565b9251935190519060405194859485614cd9565b0390a101613515565b610f6a613777865167ffffffffffffffff1690565b7f1d5ad3c50000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b9091926137be611fab858486613c63565b946137d5610f9067ffffffffffffffff88166171b3565b6138c95761380260056137fc8867ffffffffffffffff166000526008602052604060002090565b01617058565b9360005b855181101561384e5760019061384760056138358b67ffffffffffffffff166000526008602052604060002090565b01613840838a614420565b519061723e565b5001613806565b509350937f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166138bb600193976138a061389b8267ffffffffffffffff166000526008602052604060002090565b614919565b60405167ffffffffffffffff90911681529081906020820190565b0390a10191909392936134ff565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff861660045260246000fd5b346106155760006003193601126106155760206040517ff12fb6eaf1f045883c82d7d192627f7a36a50ce00c45e305919895908135a8a88152f35b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205492939260ff16156107b55760005b8281106139f95750505060005b81811061399257505050565b8067ffffffffffffffff6139ac611fab6001948688613c63565b60006139cc8267ffffffffffffffff16600052600e602052604060002090565b55167f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8600080a201613986565b80613a0a611fab6001938686613ab3565b7f56f77aeff2def50c8b5f5a0df3bab7183df09bf36c6feba496bb42551db77d7067ffffffffffffffff6020613a41858989613ab3565b0192613a6a84613a658367ffffffffffffffff16600052600e602052604060002090565b613aea565b613a7b604051928392169482613c01565b0390a201613979565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9190811015613ac35760a0020190565b613a84565b35610a0281610fc6565b63ffffffff81160361061557565b35610a0281612faa565b90606090613b2d8135613afc81613ad2565b849063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b6020810135613b3b81613ad2565b67ffffffff0000000084549160201b16807fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff83161785557fffffffffffffffffffffffffffffffffffffffff0000000000000000ffffffff6bffffffff00000000000000006040850135613bae81613ad2565b60401b169216171783550135613bc381612faa565b81547fffffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffff1690151560601b6cff00000000000000000000000016179055565b919091606080608083019463ffffffff8135613c1c81613ad2565b16845263ffffffff6020820135613c3281613ad2565b16602085015263ffffffff6040820135613c4b81613ad2565b1660408501520135613c5c81612faa565b1515910152565b9190811015613ac35760051b0190565b67ffffffffffffffff9060006060604051613c8d816114c1565b828152826020820152826040820152015216600052600e602052604060002060ff60405191613cbb836114c1565b5463ffffffff8116835263ffffffff8160201c16602084015263ffffffff8160401c16604084015260601c161515606082015290565b60405190613cfe82611468565b60008252565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610615570180359067ffffffffffffffff82116106155760200191813603831361061557565b6020818303126106155780359067ffffffffffffffff821161061557016040818303126106155760405191613d89836114dd565b813567ffffffffffffffff81116106155781613da69184016115b8565b8352602082013567ffffffffffffffff811161061557613dc692016115b8565b602082015290565b908160209103126106155751610a0281612faa565b9091613dfa610a02936040845260408401906109b0565b9160208184039101526109b0565b6040513d6000823e3d90fd5b35610a0281610710565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff16156107b55760005b828110613e975750907fc97f93e817584952f1c1d633f93784b8430f0633d002f9dcc4de4fe2780424d091613e9260405192839283614114565b0390a1565b613eaa613ea5828585613fec565b613ffc565b8051158015613fc6575b613f4d5790613f4782613f4261215d60606001965193613f336020820151613f2a613ee6604085015163ffffffff1690565b613f22613ef66080870151151590565b91613f0460a0880151151590565b94613f0d611556565b9b8c5260208c015263ffffffff1660408b0152565b151588860152565b15156080870152565b015167ffffffffffffffff1690565b61406f565b01613e58565b604080517f19d7585700000000000000000000000000000000000000000000000000000000815282516004820152602083015160248201529082015163ffffffff166044820152606082015167ffffffffffffffff16606482015260808201511515608482015260a090910151151560a482015260c490fd5b5067ffffffffffffffff613fe5606083015167ffffffffffffffff1690565b1615613eb4565b9190811015613ac35760c0020190565b60c0813603126106155760a060405191614015836114f9565b8035835260208101356020840152604081013561403181613ad2565b6040840152606081013561404481610fc6565b6060840152608081013561405781612faa565b6080840152013561406781612faa565b60a082015290565b6002608091835181556020840151600182015501916140c363ffffffff604083015116849063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b6060810151835492909101517fffffffffffffffffffffffffffffffffffffffffffffffffffff0000ffffffff90921690151560201b64ff00000000161790151560281b65ff000000000016179055565b602080825281018390526040019160005b8181106141325750505090565b90919260c080600192863581526020870135602082015263ffffffff604088013561415c81613ad2565b16604082015267ffffffffffffffff606088013561417981610fc6565b166060820152608087013561418d81612faa565b1515608082015260a08701356141a281612faa565b151560a0820152019401929101614125565b9067ffffffffffffffff610a0292166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b906040519182815491828252602082019060005260206000209260005b818110614223575050610fe392500383611515565b84546001600160a01b031683526001948501948794506020909301920161420e565b67ffffffffffffffff81116114845760051b60200190565b92919061426981614245565b936142776040519586611515565b602085838152019160051b810192831161061557905b82821061429957505050565b6020809183356142a881610710565b81520191019061428d565b90816020910312610615575190565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa90811561130b57600091614332575090565b610a02915060203d602011614354575b61434c8183611515565b8101906142b3565b503d614342565b67ffffffffffffffff610a0291166000526007602052604060002054151590565b9190811015613ac3576060020190565b60405190614399826114dd565b60606020838281520152565b906040516143b2816114a5565b608060ff600283958054855260018101546020860152015463ffffffff81166040850152818160201c161515606085015260281c161515910152565b601f8260209493601f19938186528686013760008582860101520116010190565b916020610a029381815201916143ee565b8051821015613ac35760209160051b010190565b90600182811c9216801561447d575b602083101461444e57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691614443565b906040519182600082549261449b84614434565b808452936001811690811561450757506001146144c0575b50610fe392500383611515565b90506000929192526020600020906000915b8183106144eb575050906020610fe392820101386144b3565b60209193508060019154838589010152019101909184926144d2565b60209350610fe39592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b820101386144b3565b60409067ffffffffffffffff610a02959316815281602082015201916143ee565b60405190614575826114a5565b60006080838281528260208201528260408201528260608201520152565b906040516145a0816114a5565b60806fffffffffffffffffffffffffffffffff6001839560ff8154848116875263ffffffff81871c16602088015260a01c1615156040860152015481808216166060850152821c16910152565b9190811015613ac35760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa181360301821215610615570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610615570180359067ffffffffffffffff821161061557602001918160051b3603831361061557565b9160209082815201919060005b81811061469b5750505090565b9091926020806001926001600160a01b0387356146b781610710565b16815201940192910161468e565b92906146de90610a029593604086526040860191614681565b926020818503910152614681565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b81810292918115918404141715611cab57565b91614748918354906000199060031b92831b921b19161790565b9055565b818110614757575050565b6000815560010161474c565b81519167ffffffffffffffff8311611484576801000000000000000083116114845760209082548484558085106147cd575b500190600052602060002060005b8381106147b05750505050565b60019060206001600160a01b0385511694019381840155016147a3565b6147e490846000528584600020918201910161474c565b38614795565b90805180519067ffffffffffffffff82116114845768010000000000000000821161148457602090845483865580841061486d575b500183600052602060002060005b83811061484a57505050509060016020610fe39301519101614763565b600190602061486085516001600160a01b031690565b940193818401550161482d565b61488490866000528484600020918201910161474c565b3861481f565b67ffffffffffffffff166000526008602052610a026004604060002001614487565b6002548060d01c80151590816148db575b50156148d15760a01c65ffffffffffff1690565b5060015460d01c90565b90504211386148bd565b80549060008155816148f5575050565b6000526020600020908101905b81811061490d575050565b60008155600101614902565b6005610fe39160008155600060018201556000600282015560006003820155600481016149468154614434565b9081614955575b5050016148e5565b81601f6000931160011461496d5750555b388061494d565b8183526020832061498891601f01861c81019060010161474c565b808252602082209081548360011b906000198560031b1c191617905555614966565b9190811015613ac35760051b810135907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee181360301821215610615570190565b9080601f83011215610615578135614a0181614245565b92614a0f6040519485611515565b81845260208085019260051b820101918383116106155760208201905b838210614a3b57505050505090565b813567ffffffffffffffff811161061557602091614a5e878480948801016115b8565b815201910190614a2c565b610120813603126106155760405190614a81826114a5565b614a8a81610fd8565b8252602081013567ffffffffffffffff811161061557614aad90369083016149ea565b602083015260408101359067ffffffffffffffff821161061557614ad7614af892369083016115b8565b6040840152614ae9366060830161309e565b606084015260c036910161309e565b608082015290565b8151815460208401516040850151608091821b73ffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffff000000000000000000000000000000000000000000909416939093179290921791151560a01b74ff000000000000000000000000000000000000000016919091178355606084015193810151901b7fffffffffffffffffffffffffffffffff0000000000000000000000000000000016921691909117600190910155565b9190601f8111614bd657505050565b610fe3926000526020600020906020601f840160051c83019310614c02575b601f0160051c019061474c565b9091508190614bf5565b919091825167ffffffffffffffff811161148457614c3481614c2e8454614434565b84614bc7565b6020601f8211600114614c70578190614748939495600092614c65575b50506000198260011b9260031b1c19161790565b015190503880614c51565b601f19821690614c8584600052602060002090565b9160005b818110614cc157509583600195969710614ca8575b505050811b019055565b015160001960f88460031b161c19169055388080614c9e565b9192602060018192868b015181550194019201614c89565b614d3d614d08610fe39597969467ffffffffffffffff60a09516845261010060208501526101008401906109b0565b9660408301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b01906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff1615614da757565b7fe2517d3f0000000000000000000000000000000000000000000000000000000060005233600452600060245260446000fd5b80600052600060205260ff614e06336040600020906001600160a01b0316600052602052604060002090565b541615614e105750565b7fe2517d3f000000000000000000000000000000000000000000000000000000006000523360045260245260446000fd5b610a02907f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af616770565b600254906001600160a01b038216610e6c57610a02917fffffffffffffffffffffffff00000000000000000000000000000000000000006001600160a01b0383169116176002556000616770565b908115614eca575b610a0291616770565b600254916001600160a01b038316610e6c577fffffffffffffffffffffffff00000000000000000000000000000000000000009092166001600160a01b03821617600255614ec1565b9190811015613ac35760e0020190565b35610a0281612fb4565b919060005b818110614f3f5750509050565b614f4a818386614f13565b614f5381613ac8565b614f5f610f908261435b565b6127f45781614ff561506f92614ffb60206001979601614f87614f82368361309e565b6164ec565b614ff5614fa88467ffffffffffffffff16600052600b602052604060002090565b918254614fc8614fbf8263ffffffff9060801c1690565b63ffffffff1690565b159081615262575b8161523a575b8161521f575b8161520a575b50806151fb575b6151a6575b369061309e565b906168bd565b61502a6080840191615010614f82368561309e565b67ffffffffffffffff16600052600c602052604060002090565b928354615041614fbf8263ffffffff9060801c1690565b15908161516f575b81615147575b8161512c575b81615117575b5080615108575b615075575b50369061309e565b01614f32565b61508460a06150c19201614f23565b84906fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b82547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff000000000000000000000000000000001617835538615067565b5061511282613ae0565b615062565b615126915060a01c60ff161590565b3861505b565b6fffffffffffffffffffffffffffffffff8116159150615055565b90506fffffffffffffffffffffffffffffffff6151678987015460801c90565b16159061504f565b90506fffffffffffffffffffffffffffffffff61519e898701546fffffffffffffffffffffffffffffffff1690565b161590615049565b6151b561508460408901614f23565b82547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178355614fee565b5061520581613ae0565b614fe9565b615219915060a01c60ff161590565b38614fe2565b6fffffffffffffffffffffffffffffffff8116159150614fdc565b90506fffffffffffffffffffffffffffffffff61525a8c86015460801c90565b161590614fd6565b90506fffffffffffffffffffffffffffffffff6152918c8601546fffffffffffffffffffffffffffffffff1690565b161590614fd0565b90336001600160a01b038216036152b6576152b391616455565b50565b7f6697b2320000000000000000000000000000000000000000000000000000000060005260046000fd5b6000608082016152f5610f90610b0b83613e14565b6154ab57506020820191615396602061533b61531361357787613ac8565b60801b7fffffffffffffffffffffffffffffffff000000000000000000000000000000001690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116600482015291829081906024820190565b03816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa90811561130b57839161548c575b50615464576153e86153e384613ac8565b616b62565b6153f183613ac8565b9061540a610f9060a08301936116166111718686613d04565b615424575050610fe3929161541f9150613ac8565b616c1d565b61542e9250613d04565b9061241b6040519283927f24eb47e50000000000000000000000000000000000000000000000000000000084526004840161440f565b6004827f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b6154a5915060203d602011611304576112f68183611515565b386153d2565b6154b76154eb91613e14565b7f961c9a4f0000000000000000000000000000000000000000000000000000000083526001600160a01b0316600452602490565b90fd5b916080830190615503610f90610b0b84613e14565b61567a576020840193615520602061533b61531361357789613ac8565b03816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa90811561130b5760009161565b575b506156315761556e6153e386613ac8565b61557785613ac8565b90615590610f9060a08301936116166111718686613d04565b61542457505061ffff16156156245767ffffffffffffffff9061561f6155fa6155f486611fab6155e46150107f41a8aa8df7945f0fb8ac5f7d88279638d9dc2ef9a6bf4ec9a53b80681b34aff7999a613ac8565b896155ee88613e14565b91616cb4565b92613e14565b94604051938493169583602090939291936001600160a01b0360408201951681520152565b0390a2565b5061541f610fe392613ac8565b7f53ad11d80000000000000000000000000000000000000000000000000000000060005260046000fd5b615674915060203d602011611304576112f68183611515565b3861555d565b610f6a61568683613e14565b7f961c9a4f000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b604051906156c8826114dd565b6000825260208201600081526020820151917fffffffff000000000000000000000000000000000000000000000000000000006028602483015160e01c92015193167fb148ea5f0000000000000000000000000000000000000000000000000000000081141580615772575b615745575063ffffffff1683525290565b7fa176027f0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b507f3047587c00000000000000000000000000000000000000000000000000000000811415615734565b805161011881106159f6575060048101517f000000000000000000000000000000000000000000000000000000000000000063ffffffff811663ffffffff8316036159bd575050600881015190600c81015191608c82015191609081015193609482015160b88301519360f860d88501519401519161581f895163ffffffff1690565b63ffffffff811663ffffffff84160361598357507f000000000000000000000000000000000000000000000000000000000000000063ffffffff811663ffffffff86160361594957506107d063ffffffff89160361590f576107d063ffffffff8216036158d6575091615898959391602097959361616c565b9101518082036158a6575050565b7f7be225b60000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b7f0389caa2000000000000000000000000000000000000000000000000000000006000526107d060045263ffffffff1660245260446000fd5b7f22e102a0000000000000000000000000000000000000000000000000000000006000526107d060045263ffffffff881660245260446000fd5b7f77e480260000000000000000000000000000000000000000000000000000000060005263ffffffff908116600452841660245260446000fd5b7fe366a1170000000000000000000000000000000000000000000000000000000060005263ffffffff908116600452821660245260446000fd5b7f960693cd0000000000000000000000000000000000000000000000000000000060005263ffffffff9081166004521660245260446000fd5b7f758b22cc0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b80518015615ac857602003615a8b57615a4560208251830101602083016142b3565b9060ff8211615a55575060ff1690565b61241b906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352600483016109f1565b61241b906040519182917f953576f700000000000000000000000000000000000000000000000000000000835260206004840181815201906109b0565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff8211611cab57565b60ff16604d8111611cab57600a0a90565b8015615b20576000190490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b8115615b20570490565b907f000000000000000000000000000000000000000000000000000000000000000060ff811660ff8316818114615c3f5711615c1457615b998282615aee565b91604d60ff8416118015615bfb575b615bc157505090615bbb610a0292615b02565b9061471b565b7fa9cb113d0000000000000000000000000000000000000000000000000000000060005260ff908116600452166024525060445260646000fd5b50615c0d615c0884615b02565b615b13565b8411615ba8565b615c1e8183615aee565b91604d60ff841611615bc157505090615c39610a0292615b02565b90615b4f565b5050505090565b60409067ffffffffffffffff610a02949316815281602082015201906109b0565b90805115612ce8578051602082012067ffffffffffffffff831692836000526008602052615c9c826005604060002001617482565b15615cf0575081615ce47f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea93615cdf61561f946000526009602052604060002090565b614c0c565b604051918291826109f1565b905061241b6040519283927f393b8ad200000000000000000000000000000000000000000000000000000000845260048401615c46565b9065ffffffffffff8091169116019065ffffffffffff8211611cab57565b67ffffffffffffffff166000818152600760205260409020549092919015615e475791615e4460e092615e1085615d9c7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b976164ec565b846000526008602052615db38160406000206168bd565b615dbc836164ec565b846000526008602052615dd68360026040600020016168bd565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60009060808101615e8b610f90610b0b83613e14565b615f6a57506020810190615ea9602061533b61531361357786613ac8565b03816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa90811561130b578491615f4b575b50615f2357610fe39291606082615f09615f046040615f1e9601613e14565b616f27565b615f156153e384613ac8565b01359250613ac8565b616f9c565b6004837f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b615f64915060203d602011611304576112f68183611515565b38615ee5565b6154eb6154b78492613e14565b6080810190615f8b610f90610b0b84613e14565b61567a576020810190615fa8602061533b61531361357786613ac8565b03816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa90811561130b576000916160f4575b506156315780615ffc615f04604060609401613e14565b6160086153e384613ac8565b01359261ffff811690811515806160d5575b156160c457600a5461ffff1691821161608d5750507f7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b29161561f6155fa6155f484611fab6155e461607367ffffffffffffffff98613ac8565b67ffffffffffffffff16600052600b602052604060002090565b7fe08f03ef0000000000000000000000000000000000000000000000000000000060005261ffff9081166004521660245260446000fd5b5050610fe3929150615f1e90613ac8565b506160ed6160e6600a5461ffff1690565b61ffff1690565b151561601a565b61610d915060203d602011611304576112f68183611515565b38615fe5565b949290939163ffffffff90604051958260208801981688526040870152166060850152608084015260a083015260c0820152600060e08201526107d0610100820152610100815261616661012082611515565b51902090565b959263ffffffff8095929693604051978260208a019a168a526040890152166060870152608086015260a085015260c0840152600060e084015216610100820152610100815261616661012082611515565b602081519101517fffffffff00000000000000000000000000000000000000000000000000000000604051927fb148ea5f00000000000000000000000000000000000000000000000000000000602085015260e01b166024830152602882015260288152610a02604882611515565b906000198201918211611cab57565b91908203918211611cab57565b616251614568565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff82511690602083019163ffffffff8351164203428111611cab576162b5906fffffffffffffffffffffffffffffffff6080870151169061471b565b8101809111611cab576162db6fffffffffffffffffffffffffffffffff929183926174ed565b161682524263ffffffff16905290565b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016602082015260208152610a02604082611515565b91908201809211611cab57565b80519060005b82811061634557505050565b60018101808211611cab575b8381106163615750600101616339565b6001600160a01b036163738385614420565b5116616385610faa6117ed8487614420565b1461639257600101616351565b610f6a6163a26117ed8486614420565b7fa1726e40000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b610a02906001600160a01b03600254166001600160a01b038216146163fe575b6000617300565b7fffffffffffffffffffffffff0000000000000000000000000000000000000000600254166002556163f7565b610a02907f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af617300565b90610a0291801580616494575b15617300577fffffffffffffffffffffffff000000000000000000000000000000000000000060025416600255617300565b506001600160a01b03600254166001600160a01b03831614616462565b610fe39092919260608101936fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b8051156165905760408101516fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff61655061653b60208501516fffffffffffffffffffffffffffffffff1690565b6fffffffffffffffffffffffffffffffff1690565b91161161655a5750565b61241b906040519182917f8020d124000000000000000000000000000000000000000000000000000000008352600483016164b1565b6fffffffffffffffffffffffffffffffff6165be60408301516fffffffffffffffffffffffffffffffff1690565b1615801590616605575b6165cf5750565b61241b906040519182917fd68af9cc000000000000000000000000000000000000000000000000000000008352600483016164b1565b5061662661653b60208301516fffffffffffffffffffffffffffffffff1690565b15156165c8565b90616690610fe3926002548060d01c806166dc575b50507fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff79ffffffffffff00000000000000000000000000000000000000006002549260a01b16911617600255565b79ffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffffffff00000000000000000000000000000000000000000000000000006002549260d01b16911617600255565b4211156167455765ffffffffffff61673e9160a01c1679ffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffffffff00000000000000000000000000000000000000000000000000006001549260d01b16911617600155565b3880616642565b507f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5600080a161673e565b80600052600060205260ff61679c836040600020906001600160a01b0316600052602052604060002090565b541661682b578060005260006020526167cc826040600020906001600160a01b0316600052602052604060002090565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008254161790556001600160a01b03339216907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d600080a4600190565b5050600090565b1561683957565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1991616a93613e929280546169056168ff614fbf8363ffffffff9060801c1690565b4261623c565b9081616a9f575b5050616a4d600161693060208601516fffffffffffffffffffffffffffffffff1690565b926169bb61697e61653b6fffffffffffffffffffffffffffffffff61696585546fffffffffffffffffffffffffffffffff1690565b166fffffffffffffffffffffffffffffffff88166174ed565b82906fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b616a0e6169c88751151590565b82547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178255565b019182906fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b604083015181546fffffffffffffffffffffffffffffffff1660809190911b7fffffffffffffffffffffffffffffffff0000000000000000000000000000000016179055565b604051918291826164b1565b61653b61697e916fffffffffffffffffffffffffffffffff616b13616b1a9582616b0c60018a01549282616b05616afe616ae8876fffffffffffffffffffffffffffffffff1690565b996fffffffffffffffffffffffffffffffff1690565b9560801c90565b169061471b565b9116616326565b91166174ed565b80547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178155388061690c565b67ffffffffffffffff16616b83816000526007602052604060002054151590565b15616bf057503360009081527f04c57a7d2bd5d0e733fe996f5b5aecc738999f0a2f9ddc4137bc3e1665bdf893602052604090205460ff1615616bc257565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fa9902c7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b67ffffffffffffffff7f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c91169182600052600860205280616c9160026040600020016001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391616cb4565b604080516001600160a01b0390921682526020820192909252908190810161561f565b8054939290919060ff60a086901c16158015616ed5575b616ece57616cea6fffffffffffffffffffffffffffffffff861661653b565b9060018401958654616d246168ff614fbf616d1761653b856fffffffffffffffffffffffffffffffff1690565b9460801c63ffffffff1690565b80616e3a575b5050838110616dfc5750828210616d8a5750610fe3939450616d4f9161653b9161623c565b6fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b90616dc1610f6a93616dbc616dad84616da761653b8c5460801c90565b9361623c565b616db68361622d565b90616326565b615b4f565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024526001600160a01b0316604452606490565b7f1a76572a0000000000000000000000000000000000000000000000000000000060005260045260248390526001600160a01b031660445260646000fd5b828592939511616ea457616e5461653b616e5b9460801c90565b91856172e3565b84547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178555913880616d2a565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050509050565b508115616ccb565b65ffffffffffff8111616ef55765ffffffffffff1690565b7f6dfcc65000000000000000000000000000000000000000000000000000000000600052603060045260245260446000fd5b7f0000000000000000000000000000000000000000000000000000000000000000616f4f5750565b6001600160a01b031680600052600460205260406000205415616f6f5750565b7fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b67ffffffffffffffff7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894491169182600052600860205280616c9160406000206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391616cb4565b604051906006548083528260208101600660005260206000209260005b81811061703f575050610fe392500383611515565b845483526001948501948794506020909301920161702a565b906040519182815491828252602082019060005260206000209260005b81811061708a575050610fe392500383611515565b8454835260019485019487945060209093019201617075565b8054821015613ac35760005260206000200190600090565b805480156170e35760001901906170d282826170a3565b60001982549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60008181526004602052604090205490811561682b57600019820190828211611cab57600354926000198401938411611cab5783836000956171729503617178575b50505061716160036170bb565b600490600052602052604060002090565b55600190565b6171616171a49161719a6171906171aa9560036170a3565b90549060031b1c90565b92839160036170a3565b9061472e565b55388080617154565b60008181526007602052604090205490811561682b57600019820190828211611cab57600654926000198401938411611cab5783836000956171729503617213575b50505061720260066170bb565b600790600052602052604060002090565b6172026171a49161722b6171906172359560066170a3565b92839160066170a3565b553880806171f5565b60018101918060005282602052604060002054928315156000146172da576000198401848111611cab578354936000198501948511611cab5760009585836171729761729295036172a1575b5050506170bb565b90600052602052604060002090565b6172c16171a4916172b86171906172d195886170a3565b928391876170a3565b8590600052602052604060002090565b5538808061728a565b50505050600090565b926172ee919261471b565b8101809111611cab57610a02916174ed565b80600052600060205260ff61732c836040600020906001600160a01b0316600052602052604060002090565b54161561682b5780600052600060205261735d826040600020906001600160a01b0316600052602052604060002090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541690556001600160a01b03339216907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b600080a4600190565b60008181526004602052604090205461742657600354680100000000000000008110156114845761740d6173f682600185940160035560036170a3565b81939154906000199060031b92831b921b19161790565b9055600354906000526004602052604060002055600190565b50600090565b6000818152600760205260409020546174265760065468010000000000000000811015611484576174696173f682600185940160065560066170a3565b9055600654906000526007602052604060002055600190565b600082815260018201602052604090205461682b578054906801000000000000000082101561148457826174c06173f68460018096018555846170a3565b905580549260005201602052604060002055600190565b620697808110156174e55790565b506206978090565b90808210156174fa575090565b905090565b9192901561757a5750815115617513575090565b3b1561751c5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b82519091501561758d5750805190602001fd5b61241b906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352600483016109f156fea164736f6c634300081a000aad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5",
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) GetRequiredInboundCCVs(opts *bind.CallOpts, arg0 common.Address, sourceChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "getRequiredInboundCCVs", arg0, sourceChainSelector, arg2, arg3, arg4)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GetRequiredInboundCCVs(arg0 common.Address, sourceChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetRequiredInboundCCVs(&_USDCTokenPoolCCTPV2.CallOpts, arg0, sourceChainSelector, arg2, arg3, arg4)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) GetRequiredInboundCCVs(arg0 common.Address, sourceChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetRequiredInboundCCVs(&_USDCTokenPoolCCTPV2.CallOpts, arg0, sourceChainSelector, arg2, arg3, arg4)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) GetRequiredOutboundCCVs(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "getRequiredOutboundCCVs", arg0, destChainSelector, arg2, arg3, arg4)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GetRequiredOutboundCCVs(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetRequiredOutboundCCVs(&_USDCTokenPoolCCTPV2.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) GetRequiredOutboundCCVs(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetRequiredOutboundCCVs(&_USDCTokenPoolCCTPV2.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GetRouter() (common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetRouter(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) GetRouter() (common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetRouter(&_USDCTokenPoolCCTPV2.CallOpts)
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "setRouter", newRouter)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.SetRouter(&_USDCTokenPoolCCTPV2.TransactOpts, newRouter)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.SetRouter(&_USDCTokenPoolCCTPV2.TransactOpts, newRouter)
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
	RemoteChainSelector uint64
	OutboundCCVs        []common.Address
	InboundCCVs         []common.Address
	Raw                 types.Log
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

type USDCTokenPoolCCTPV2RouterUpdatedIterator struct {
	Event *USDCTokenPoolCCTPV2RouterUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2RouterUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2RouterUpdated)
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
		it.Event = new(USDCTokenPoolCCTPV2RouterUpdated)
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

func (it *USDCTokenPoolCCTPV2RouterUpdatedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2RouterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2RouterUpdated struct {
	OldRouter common.Address
	NewRouter common.Address
	Raw       types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterRouterUpdated(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2RouterUpdatedIterator, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2RouterUpdatedIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "RouterUpdated", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2RouterUpdated) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2RouterUpdated)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseRouterUpdated(log types.Log) (*USDCTokenPoolCCTPV2RouterUpdated, error) {
	event := new(USDCTokenPoolCCTPV2RouterUpdated)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
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
	return common.HexToHash("0xb0897119e8510f887b892cbc4c8506fc51d9849fd90afae4fd065e705f2d0f6c")
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

func (USDCTokenPoolCCTPV2RouterUpdated) Topic() common.Hash {
	return common.HexToHash("0x02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f1684")
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

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRequiredInboundCCVs(opts *bind.CallOpts, arg0 common.Address, sourceChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error)

	GetRequiredOutboundCCVs(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

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

	SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error)

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

	FilterRouterUpdated(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2RouterUpdatedIterator, error)

	WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2RouterUpdated) (event.Subscription, error)

	ParseRouterUpdated(log types.Log) (*USDCTokenPoolCCTPV2RouterUpdated, error)

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*USDCTokenPoolCCTPV2TokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2TokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*USDCTokenPoolCCTPV2TokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*USDCTokenPoolCCTPV2TokenTransferFeeConfigUpdatedIterator, error)

	WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2TokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigUpdated(log types.Log) (*USDCTokenPoolCCTPV2TokenTransferFeeConfigUpdated, error)

	Address() common.Address
}
