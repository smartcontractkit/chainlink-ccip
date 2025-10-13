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
	RemoteChainSelector    uint64
	OutboundCCVs           []common.Address
	AdditionalOutboundCCVs []common.Address
	InboundCCVs            []common.Address
	AdditionalInboundCCVs  []common.Address
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"internalType\":\"contractITokenMessenger\"},{\"name\":\"cctpMessageTransmitterProxy\",\"type\":\"address\",\"internalType\":\"contractCCTPMessageTransmitterProxy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"FINALITY_THRESHOLD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_FEE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_USDC_MESSAGE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"structAuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"additionalOutboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"additionalInboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyFinalityConfigUpdates\",\"inputs\":[{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"destToUseDefaultFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAccumulatedFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDomain\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structUSDCTokenPool.Domain\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useLegacySourcePoolDataFormat\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredInboundCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredOutboundCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_localDomainIdentifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_messageTransmitterProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractCCTPMessageTransmitterProxy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_supportedUSDCVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_tokenMessenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITokenMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustomFinalityRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDomains\",\"inputs\":[{\"name\":\"domains\",\"type\":\"tuple[]\",\"internalType\":\"structUSDCTokenPool.DomainUpdate[]\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useLegacySourcePoolDataFormat\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"additionalOutboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"additionalInboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityTransferInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DomainsSet\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structUSDCTokenPool.DomainUpdate[]\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useLegacySourcePoolDataFormat\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalityConfigUpdated\",\"inputs\":[{\"name\":\"finalityConfig\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeeWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidBurnToken\",\"inputs\":[{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"got\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDepositHash\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"got\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidDestinationDomain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidDomain\",\"inputs\":[{\"name\":\"domain\",\"type\":\"tuple\",\"internalType\":\"structUSDCTokenPool.DomainUpdate\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useLegacySourcePoolDataFormat\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"type\":\"error\",\"name\":\"InvalidExecutionFinalityThreshold\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidFinality\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidFinalityConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMessageLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMinFinalityThreshold\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidNonce\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"got\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidPreviousPool\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourceDomain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenMessengerVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidTransmitterInProxy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenTransferFeeConfigNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnknownDomain\",\"inputs\":[{\"name\":\"domain\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"UnlockingUSDCFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x6101808060405234610684576179cd803803809161001d8285610a0d565b8339810160c082820312610684578151906001600160a01b038216808303610684576020840151926001600160a01b038416908185036106845760408601516001600160a01b038116969094908786036106845760608101516001600160401b0381116106845781019180601f84011215610684578251926001600160401b038411610689578360051b9060208201946100ba6040519687610a0d565b855260208086019282010192831161068457602001905b8282106109f5575050506100f360a06100ec60808401610a30565b9201610a30565b9060209660405199610105898c610a0d565b60008b52600036813733156109e457600180546001600160a01b03191633179055801580156109d3575b80156109c2575b6109b157600492899260805260c0526040519283809263313ce56760e01b82525afa809160009161097a575b5090610956575b50600660a052600480546001600160a01b0319166001600160a01b03929092169190911790558051151560e0819052610830575b50604051926101ac8585610a0d565b60008452600036813760408051979088016001600160401b03811189821017610689576040528752838588015260005b8451811015610243576001906001600160a01b036101fa8288610a60565b51168761020682610ae1565b610213575b5050016101dc565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a1388761020b565b508493508587519260005b84518110156102bf576001600160a01b036102698287610a60565b51169081156102ae577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef88836102a0600195610bfd565b50604051908152a10161024e565b6342bcdf7f60e11b60005260046000fd5b5090859185600161010052841561081f57604051632c12192160e01b81528481600481895afa9081156106ef576000916107ea575b5060405163054fd4d560e41b81526001600160a01b039190911691908581600481865afa9081156106ef576000916107cd575b5063ffffffff8061010051169116908082036107b6575050604051639cdbb18160e01b815285816004818a5afa9081156106ef57600091610799575b5063ffffffff80610100511691169080820361078257505084600491604051928380926367e0ed8360e11b82525afa80156106ef578291600091610739575b506001600160a01b03160361072857600492849261012052610140526040519283809263234d8e3d60e21b82525afa9081156106ef576000916106fb575b506101605260805161012051604051636eb1769f60e11b81523060048201526001600160a01b03918216602482018190529492909116908381604481855afa9081156106ef576000916106c2575b5060001981018091116106ac57604051908482019563095ea7b360e01b87526024830152604482015260448152610466606482610a0d565b6000806040968751936104798986610a0d565b8785527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656488860152519082865af13d1561069f573d906001600160401b0382116106895786516104e69490926104d8601f8201601f1916890185610a0d565b83523d60008885013e610cca565b80518061060b575b847f2e902d38f15b233cbb63711add0fca4545334d3a169d60c0a616494d7eea954485858351908152a151616c329081610d9b823960805181818161059e015281816107a40152818161080201528181610df801528181611e0e015281816124eb01528181613b97015281816163ab015261671e015260a05181818161090201528181615464015281816154f40152615c86015260c051818181612ae101528181614d0601528181614eaa0152818161583c0152615948015260e0518181816112c501528181612be7015261664801526101005181818161075301526151470152610120518181816114e80152611dbf015261014051818181610d600152611c990152610160518181816116d601528181611f0a01526151cd0152f35b81849181010312610684578201518015908115036106845761062e5783806104ee565b50608491519062461bcd60e51b82526004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b916104e692606091610cca565b634e487b7160e01b600052601160045260246000fd5b90508381813d83116106e8575b6106d98183610a0d565b8101031261068457518561042e565b503d6106cf565b6040513d6000823e3d90fd5b61071b9150823d8411610721575b6107138183610a0d565b810190610a44565b836103e0565b503d610709565b632a32133b60e11b60005260046000fd5b9091508581813d831161077b575b6107518183610a0d565b810103126107775751906001600160a01b038216820361077457508190876103a2565b80fd5b5080fd5b503d610747565b633785f8f160e01b60005260045260245260446000fd5b6107b09150863d8811610721576107138183610a0d565b87610363565b63960693cd60e01b60005260045260245260446000fd5b6107e49150863d8811610721576107138183610a0d565b87610327565b90508481813d8311610818575b6108018183610a0d565b810103126106845761081290610a30565b866102f4565b503d6107f7565b6306b7c75960e31b60005260046000fd5b909194604051936108418686610a0d565b60008552600036813760e051156109455760005b85518110156108bc576001906001600160a01b036108738289610a60565b51168861087f82610c36565b61088c575b505001610855565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a13888610884565b50919350919460005b8451811015610939576001906001600160a01b036108e38288610a60565b5116801561093357876108f582610bbe565b610903575b50505b016108c5565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a138876108fa565b506108fd565b5091949092503861019d565b6335f4a7b360e01b60005260046000fd5b60ff1660068114610169576332ad3e0760e11b600052600660045260245260446000fd5b8881813d83116109aa575b61098f8183610a0d565b8101031261077757519060ff82168203610774575038610162565b503d610985565b630a64406560e11b60005260046000fd5b506001600160a01b03831615610136565b506001600160a01b0384161561012f565b639b15e16f60e01b60005260046000fd5b60208091610a0284610a30565b8152019101906100d1565b601f909101601f19168101906001600160401b0382119082101761068957604052565b51906001600160a01b038216820361068457565b90816020910312610684575163ffffffff811681036106845790565b8051821015610a745760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b8054821015610a745760005260206000200190600090565b80548015610acb576000190190610ab98282610a8a565b8154906000199060031b1b1916905555565b634e487b7160e01b600052603160045260246000fd5b6000818152601160205260409020548015610b8c5760001981018181116106ac576010546000198101919082116106ac57808203610b3b575b505050610b276010610aa2565b600052601160205260006040812055600190565b610b74610b4c610b5d936010610a8a565b90549060031b1c9283926010610a8a565b819391549060031b91821b91600019901b19161790565b90556000526011602052604060002055388080610b1a565b5050600090565b805490680100000000000000008210156106895781610b5d916001610bba94018155610a8a565b9055565b80600052600360205260406000205415600014610bf757610be0816002610b93565b600254906000526003602052604060002055600190565b50600090565b80600052601160205260406000205415600014610bf757610c1f816010610b93565b601054906000526011602052604060002055600190565b6000818152600360205260409020548015610b8c5760001981018181116106ac576002546000198101919082116106ac57818103610c90575b505050610c7c6002610aa2565b600052600360205260006040812055600190565b610cb2610ca1610b5d936002610a8a565b90549060031b1c9283926002610a8a565b90556000526003602052604060002055388080610c6f565b91929015610d2c5750815115610cde575090565b3b15610ce75790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b825190915015610d3f5750805190602001fd5b6040519062461bcd60e51b8252602060048301528181519182602483015260005b838110610d825750508160006044809484010152601f80199101168101030190fd5b60208282018101516044878401015285935001610d6056fe6080604052600436101561001257600080fd5b60003560e01c806301ffc9a714610357578063164e68de14610352578063181f5a771461034d578063212a052e1461034857806321df0da714610343578063240028e81461033e5780632451a6271461033957806324f65ee7146103345780632a10097b1461032f5780632c286daf1461032a57806337b19247146103255780633907753714610320578063489a68f21461031b5780634ac8bd5f146103165780634c5ef0ed146103115780634f71592c1461030c57806354c8a4f3146103075780635df45a37146103025780636155cda0146102fd57806362ddd3c4146102f8578063698c2c66146102f35780636b716b0d146102ee5780636d3d1a58146102e95780637437ff9f146102e457806379ba5097146102df5780637d54534e146102da578063804ba5a9146102d55780638926f54f146102d05780638da5cb5b146102cb57806391a2749a146102c6578063962d4020146102c157806398db9643146102bc5780639a4575b9146102b75780639f68f673146102b2578063a42a7b8b146102ad578063a7cd63b7146102a8578063acfecf91146102a3578063af58d59f1461029e578063b1c71c6514610299578063b794658014610294578063bc063e1a1461028f578063c4bffe2b1461028a578063c75eea9c14610285578063c8c8fd1914610280578063cf7401f31461027b578063d966866b14610276578063da4b05e714610271578063dc0bd9711461026c578063dfadfa3514610267578063e0351e1314610262578063e8a1da171461025d5763f2fde38b1461025857600080fd5b61302b565b612c0c565b612bcf565b612b05565b612ab4565b612a97565b6128e9565b61280b565b612717565b6126cc565b612657565b6125f7565b6125c0565b612453565b61238b565b612264565b61220a565b612126565b61208a565b611cf8565b611c6c565b611acd565b611a08565b61194b565b61190c565b6118ca565b611839565b61176e565b61172e565b6116fa565b6116b9565b6115c8565b611545565b6114bb565b611498565b611293565b611227565b611182565b610f7c565b610eb0565b610c7b565b610ba7565b610a3f565b610957565b6108e8565b610882565b6107c8565b610777565b610736565b6106d7565b6104bb565b34610498576020600319360112610498576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361049857807ff208a58f000000000000000000000000000000000000000000000000000000006103ec921490811561046e575b8115610444575b811561041a575b81156103f0575b5060405190151581529081906020820190565b0390f35b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014386103d9565b7f0e64dd2900000000000000000000000000000000000000000000000000000000811491506103d2565b7f1ef5498f00000000000000000000000000000000000000000000000000000000811491506103cb565b7faff2afbf00000000000000000000000000000000000000000000000000000000811491506103c4565b600080fd5b73ffffffffffffffffffffffffffffffffffffffff81160361049857565b34610498576020600319360112610498576004356104d88161049d565b6104e0614866565b6104e8613b4e565b90816104f057005b73ffffffffffffffffffffffffffffffffffffffff7f738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b8595999161064960405161062e60208201917fa9059cbb00000000000000000000000000000000000000000000000000000000835261059b8161058d8a88602484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b03601f1981018352826110b5565b857f0000000000000000000000000000000000000000000000000000000000000000166000806040958651946105d188876110b5565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020870152519082855af13d1561066f573d9161061283611114565b9261061f875194856110b5565b83523d6000602085013e616b61565b80518061064e575b5050519283921694829190602083019252565b0390a2005b81602080610663936106689501019101613564565b615ef4565b3880610636565b606091616b61565b600091031261049857565b919082519283825260005b8481106106ae575050601f19601f8460006020809697860101520116010190565b8060208092840101518282860101520161068d565b9060206106d4928181520190610682565b90565b34610498576000600319360112610498576103ec60408051906106fa81836110b5565b601d82527f55534443546f6b656e506f6f6c43435450563220312e372e302d646576000000602083015251918291602083526020830190610682565b3461049857600060031936011261049857602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461049857600060031936011261049857602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346104985760206003193601126104985760206108286004356107ea8161049d565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691161490565b6040519015158152f35b602060408183019282815284518094520192019060005b8181106108565750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101610849565b34610498576000600319360112610498576040516010548082526020820190601060005260206000209060005b8181106108d2576103ec856108c6818703826110b5565b60405191829182610832565b82548452602090930192600192830192016108af565b3461049857600060031936011261049857602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b9181601f840112156104985782359167ffffffffffffffff8311610498576020808501948460051b01011161049857565b346104985760406003193601126104985760043567ffffffffffffffff811161049857366023820112156104985780600401359067ffffffffffffffff82116104985736602460a0840283010111610498576024359067ffffffffffffffff8211610498576109d9926109d06024933690600401610926565b93909201613101565b005b6004359061ffff8216820361049857565b6024359061ffff8216820361049857565b6064359061ffff8216820361049857565b9181601f840112156104985782359167ffffffffffffffff83116104985760208085019460e0850201011161049857565b3461049857606060031936011261049857610a586109db565b610a606109ec565b60443567ffffffffffffffff811161049857610a80903690600401610a0e565b9290610a8a614866565b61ffff8316612710811015610b1f57507f52aa194b292c8bfb5aaca8ee2000a965c3a051b306ff841873b16147526a39ba93610b0091600b5463ffff00008660101b16907fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000061ffff871691161717600b556148cb565b6040805161ffff9283168152929091166020830152819081015b0390a1005b7f95f3517a0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b67ffffffffffffffff81160361049857565b3590610b6982610b4c565b565b908160a09103126104985790565b9181601f840112156104985782359167ffffffffffffffff8311610498576020838186019501011161049857565b346104985760a060031936011261049857610bc360043561049d565b602435610bcf81610b4c565b60443567ffffffffffffffff811161049857610bef903690600401610b6b565b50610bf86109fd565b5060843567ffffffffffffffff8111610498576103ec91610c20610c27923690600401610b79565b5050613409565b60405191829182919091606080608083019463ffffffff815116845263ffffffff602082015116602085015263ffffffff604082015116604085015201511515910152565b90816101009103126104985790565b346104985760206003193601126104985760043567ffffffffffffffff811161049857610cac903690600401610c6c565b610cb4613487565b506060810135610cc48183614c37565b610d456020610ce1610cd960e086018661349a565b8101906134eb565b610d0a610d03610cfe610cf760c089018961349a565b3691611130565b615053565b8251615134565b8181519101519060405193849283927f57ecfd2800000000000000000000000000000000000000000000000000000000845260048401613579565b0381600073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af1908115610eab57600091610e7c575b5015610e5257817ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc067ffffffffffffffff610ddd6040610dd660206103ec980161325e565b94016135aa565b6040805173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000081168252336020830152929092169082015260608101859052921691608090a2610e3f6110d8565b8190526040519081529081906020820190565b7fbf969f220000000000000000000000000000000000000000000000000000000060005260046000fd5b610e9e915060203d602011610ea4575b610e9681836110b5565b810190613564565b38610d91565b503d610e8c565b61359e565b346104985760406003193601126104985760043567ffffffffffffffff811161049857610ee46103ec913690600401610c6c565b610eec6109ec565b906000604051610efb81611008565b52610f2d610f256060830135610f1f610f1a610cf760c087018761349a565b6153bb565b906154f1565b928383614e5f565b7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc067ffffffffffffffff610ddd60206040850194610f6b863561049d565b013593610f7785610b4c565b6135aa565b346104985760206003193601126104985760043567ffffffffffffffff8111610498573660238201121561049857806004013567ffffffffffffffff81116104985736602460c08302840101116104985760246109d992016135b4565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff82111761102457604052565b610fd9565b6040810190811067ffffffffffffffff82111761102457604052565b6060810190811067ffffffffffffffff82111761102457604052565b60a0810190811067ffffffffffffffff82111761102457604052565b6080810190811067ffffffffffffffff82111761102457604052565b60c0810190811067ffffffffffffffff82111761102457604052565b90601f601f19910116810190811067ffffffffffffffff82111761102457604052565b60405190610b696020836110b5565b60405190610b696040836110b5565b60405190610b696080836110b5565b60405190610b6960a0836110b5565b67ffffffffffffffff811161102457601f01601f191660200190565b92919261113c82611114565b9161114a60405193846110b5565b829481845281830111610498578281602093846000960137010152565b9080601f83011215610498578160206106d493359101611130565b346104985760406003193601126104985760043561119f81610b4c565b60243567ffffffffffffffff8111610498576020916111c5610828923690600401611167565b9061391b565b60a0600319820112610498576004356111e38161049d565b916024356111f081610b4c565b916044359160643561ffff8116810361049857916084359067ffffffffffffffff82116104985761122391600401610b79565b9091565b34610498576103ec6108c661123b366111cb565b5050509150613a3a565b60406003198201126104985760043567ffffffffffffffff8111610498578161127091600401610926565b929092916024359067ffffffffffffffff82116104985761122391600401610926565b34610498576112bb6112c36112a736611245565b94916112b4939193614866565b3691611997565b923691611997565b7f00000000000000000000000000000000000000000000000000000000000000001561146e5760005b82518110156113b1578061131f61130560019386613a26565b5173ffffffffffffffffffffffffffffffffffffffff1690565b61135b61135673ffffffffffffffffffffffffffffffffffffffff83165b73ffffffffffffffffffffffffffffffffffffffff1690565b6167b5565b611367575b50016112ec565b60405173ffffffffffffffffffffffffffffffffffffffff9190911681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756690602090a138611360565b5060005b81518110156109d957806113ce61130560019385613a26565b73ffffffffffffffffffffffffffffffffffffffff8116156114685761141161140c73ffffffffffffffffffffffffffffffffffffffff831661133d565b6169e8565b61141e575b505b016113b5565b60405173ffffffffffffffffffffffffffffffffffffffff9190911681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d890602090a183611416565b50611418565b7f35f4a7b30000000000000000000000000000000000000000000000000000000060005260046000fd5b346104985760006003193601126104985760206114b3613b4e565b604051908152f35b3461049857600060031936011261049857602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b9060406003198301126104985760043561152581610b4c565b916024359067ffffffffffffffff82116104985761122391600401610b79565b34610498576115533661150c565b61155e929192614866565b67ffffffffffffffff8216611580816000526007602052604060002054151590565b1561159b57506109d992611595913691611130565b906155ff565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b34610498576040600319360112610498576004356115e58161049d565b6024356115f0614866565b73ffffffffffffffffffffffffffffffffffffffff821691821561168f577f78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f0238927fffffffffffffffffffffffff0000000000000000000000000000000000000000600454161760045581600555610b1a604051928392836020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461049857600060031936011261049857602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461049857600060031936011261049857602073ffffffffffffffffffffffffffffffffffffffff600a5416604051908152f35b34610498576000600319360112610498576004546005546040805173ffffffffffffffffffffffffffffffffffffffff9093168352602083019190915290f35b346104985760006003193601126104985760005473ffffffffffffffffffffffffffffffffffffffff8116330361180f577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b34610498576020600319360112610498577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174602073ffffffffffffffffffffffffffffffffffffffff60043561188e8161049d565b611896614866565b16807fffffffffffffffffffffffff0000000000000000000000000000000000000000600a541617600a55604051908152a1005b346104985760206003193601126104985760043567ffffffffffffffff8111610498576118fe6109d9913690600401610a0e565b90611907614866565b6148cb565b3461049857602060031936011261049857602061082867ffffffffffffffff60043561193781610b4c565b166000526007602052604060002054151590565b3461049857600060031936011261049857602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b67ffffffffffffffff81116110245760051b60200190565b9291906119a38161197f565b936119b160405195866110b5565b602085838152019160051b810192831161049857905b8282106119d357505050565b6020809183356119e28161049d565b8152019101906119c7565b9080601f83011215610498578160206106d493359101611997565b346104985760206003193601126104985760043567ffffffffffffffff8111610498576040600319823603011261049857604051611a4581611029565b816004013567ffffffffffffffff811161049857611a6990600436918501016119ed565b8152602482013567ffffffffffffffff8111610498576109d9926004611a9292369201016119ed565b6020820152613c15565b9181601f840112156104985782359167ffffffffffffffff8311610498576020808501946060850201011161049857565b346104985760606003193601126104985760043567ffffffffffffffff811161049857611afe903690600401610926565b9060243567ffffffffffffffff811161049857611b1f903690600401611a9c565b9060443567ffffffffffffffff811161049857611b40903690600401611a9c565b611b6261133d600a5473ffffffffffffffffffffffffffffffffffffffff1690565b33141580611c41575b611c0f57838614801590611c05575b611bdb5760005b868110611b8a57005b80611bd5611ba3611b9e6001948b8b6133f9565b61325e565b611bae838989613da4565b611bcf611bc7611bbf86898b613da4565b9236906127c2565b9136906127c2565b916156bf565b01611b81565b7f568efce20000000000000000000000000000000000000000000000000000000060005260046000fd5b5080861415611b7a565b7f8e4a23d6000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b6000fd5b50611c6461133d60015473ffffffffffffffffffffffffffffffffffffffff1690565b331415611b6b565b3461049857600060031936011261049857602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b6106d4916020611cd68351604084526040840190610682565b920151906020818403910152610682565b9060206106d4928181520190611cbd565b346104985760206003193601126104985760043567ffffffffffffffff811161049857611d29903690600401610b6b565b611d31613db4565b50611d3b816157ef565b60208101611d6d611d68611d4e8361325e565b67ffffffffffffffff166000526012602052604060002090565b613dcd565b91611d82611d7e6060850151151590565b1590565b612048576020611d92828061349a565b9050036120045760208301518015611fe857925b606073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016920135926040820192611df6845163ffffffff1690565b9273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016938151833b15610498576040517fd04857b00000000000000000000000000000000000000000000000000000000081526004810189905263ffffffff9290921660248301526044820189905273ffffffffffffffffffffffffffffffffffffffff861660648301526084820152600060a482018190526107d060c4830152909492859060e490829084905af18015610eab57611f5f611fb096611f3e7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1094866103ec9c611fab9a67ffffffffffffffff97611fcd575b50611f347f0000000000000000000000000000000000000000000000000000000000000000955163ffffffff1690565b9251928d86615aa7565b611f55611f496110e7565b63ffffffff9093168352565b6020820152615b52565b96611fa3611f6c8661325e565b6040805173ffffffffffffffffffffffffffffffffffffffff90971687523360208801528601929092529116929081906060820190565b0390a261325e565b6140df565b90611fb96110e7565b918252602082015260405191829182611ce7565b80611fdc6000611fe2936110b5565b80610677565b38611f04565b50611ffe611ff6828061349a565b810190613e48565b92611da6565b8061200e9161349a565b906120446040519283927fa3c8cf0900000000000000000000000000000000000000000000000000000000845260048401613e37565b0390fd5b611c3d6120548361325e565b7fd201c48a0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b34610498576103ec6108c661209e366111cb565b5050509150613e57565b602081016020825282518091526040820191602060408360051b8301019401926000915b8383106120db57505050505090565b9091929394602080612117837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc086600196030187528951610682565b970193019301919392906120cc565b346104985760206003193601126104985767ffffffffffffffff60043561214c81610b4c565b1660005260086020526121656005604060002001615fca565b805190601f1961218d6121778461197f565b9361218560405195866110b5565b80855261197f565b0160005b8181106121f957505060005b81518110156121eb57806121cf6121ca6121b960019486613a26565b516000526009602052604060002090565b613f79565b6121d98286613a26565b526121e48185613a26565b500161219d565b604051806103ec85826120a8565b806060602080938701015201612191565b34610498576000600319360112610498576040516002548082526020820190600260005260206000209060005b81811061224e576103ec856108c6818703826110b5565b8254845260209093019260019283019201612237565b34610498576122723661150c565b61227d929192614866565b67ffffffffffffffff8216916122a3611d7e846000526007602052604060002054151590565b612354576122e6611d7e60056122cd8467ffffffffffffffff166000526008602052604060002090565b016122d9368689611130565b60208151910120906168e8565b61231d57507f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76919261064960405192839283613e37565b61204484926040519384937f74f23c7c00000000000000000000000000000000000000000000000000000000855260048501614039565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b346104985760206003193601126104985767ffffffffffffffff6004356123b181610b4c565b6123b961405a565b501660005260086020526103ec6123de6123d96002604060002001614085565b615bdd565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b92919061244e602091604086526040860190611cbd565b930152565b346104985760606003193601126104985760043567ffffffffffffffff811161049857612484903690600401610b6b565b61248c6109ec565b9060443567ffffffffffffffff8111610498576124ad903690600401611167565b506124b6613db4565b506124c182826158fe565b606081013561ffff81931661258f575b50611fab602061256892016124e58161325e565b604080517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16815233602082015290810186905267ffffffffffffffff91909116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10908060608101611fa3565b612570615c7f565b6125786110e7565b91825260208201526103ec60405192839283612437565b9091506127106125a861ffff600b5460101c1683614195565b0481039081116125bb5790611fab6124d1565b6139b9565b34610498576020600319360112610498576103ec6125e3600435611fab81610b4c565b604051918291602083526020830190610682565b3461049857600060031936011261049857602060405160008152f35b602060408183019282815284518094520192019060005b8181106126375750505090565b825167ffffffffffffffff1684526020938401939092019160010161262a565b3461049857600060031936011261049857612670615f7f565b805190601f196126826121778461197f565b0136602084013760005b81518110156126be578067ffffffffffffffff6126ab60019385613a26565b51166126b78286613a26565b520161268c565b604051806103ec8582612613565b346104985760206003193601126104985767ffffffffffffffff6004356126f281610b4c565b6126fa61405a565b501660005260086020526103ec6123de6123d96040600020614085565b346104985760006003193601126104985760206040516101188152f35b8015150361049857565b6fffffffffffffffffffffffffffffffff81160361049857565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c6060910112610498576040519061278f82611045565b8160843561279c81612734565b815260a4356127aa8161273e565b6020820152604060c435916127be8361273e565b0152565b9190826060910312610498576040516127da81611045565b604080829480356127ea81612734565b845260208101356127fa8161273e565b60208501520135916127be8361273e565b346104985760e06003193601126104985760043561282881610b4c565b60607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc3601126104985760405161285e81611045565b60243561286a81612734565b81526044356128788161273e565b60208201526064356128898161273e565b604082015261289736612758565b9073ffffffffffffffffffffffffffffffffffffffff600a5416331415806128c7575b611c0f576109d9926156bf565b5073ffffffffffffffffffffffffffffffffffffffff600154163314156128ba565b346104985760206003193601126104985760043567ffffffffffffffff81116104985761291a903690600401610926565b612922614866565b60005b81811061292e57005b80837fece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d67ffffffffffffffff61296a611b9e6001968886614101565b612a8e61298561297b878a88614101565b6020810190614141565b896129a06129968a838b969b614101565b6040810190614141565b6129d36129c98c6129c16129b782888b989b614101565b6060810190614141565b969095614101565b6080810190614141565b9590946129e96129e436838f611997565b615cba565b6129f76129e4368585611997565b612a056129e4368787611997565b612a136129e4368989611997565b612a808c612a2b612a226110f6565b91843691611997565b8152612a38368686611997565b6020820152612a48368888611997565b6040820152612a58368a8a611997565b6060820152612a7b8b67ffffffffffffffff16600052600e602052604060002090565b614271565b604051998a99169b89614391565b0390a201612925565b346104985760006003193601126104985760206040516107d08152f35b3461049857600060031936011261049857602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346104985760206003193601126104985767ffffffffffffffff600435612b2b81610b4c565b612b3361405a565b501660005260126020526103ec604060002060ff600260405192612b5684611061565b8054845260018101546020850152015463ffffffff81166040840152818160201c161515606084015260281c16151560808201526040519182918291909160808060a0830194805184526020810151602085015263ffffffff604082015116604085015260608101511515606085015201511515910152565b346104985760006003193601126104985760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b3461049857612c1a36611245565b919092612c25614866565b6000915b808310612ed75750505060009163ffffffff4216925b828110612c4857005b612c5b612c568285856144a2565b614561565b9060608201612c6a8151615db3565b6080830193612c798551615db3565b60408401908151511561168f57612cb3611d7e612cae612ca1885167ffffffffffffffff1690565b67ffffffffffffffff1690565b616a23565b612e8c57612dec612cec612cd2879a999a5167ffffffffffffffff1690565b67ffffffffffffffff166000526008602052604060002090565b612daf89612da98751612d90612d1560408301516fffffffffffffffffffffffffffffffff1690565b91612d77612d40612d3960208401516fffffffffffffffffffffffffffffffff1690565b9251151590565b612d6e612d4b611105565b6fffffffffffffffffffffffffffffffff851681529763ffffffff166020890152565b15156040870152565b6fffffffffffffffffffffffffffffffff166060850152565b6fffffffffffffffffffffffffffffffff166080830152565b826145f8565b612de189612dd88a51612d90612d1560408301516fffffffffffffffffffffffffffffffff1690565b600283016145f8565b600484519101614704565b602085019660005b88518051821015612e2f5790612e29600192612e2283612e1c8c5167ffffffffffffffff1690565b92613a26565b51906155ff565b01612df4565b50509796509490612e837f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c29392612e706001975167ffffffffffffffff1690565b92519351905190604051948594856147d1565b0390a101612c3f565b611c3d612ea1865167ffffffffffffffff1690565b7f1d5ad3c50000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b909192612ee8611b9e8584866133f9565b94612eff611d7e67ffffffffffffffff881661685d565b612ff357612f2c6005612f268867ffffffffffffffff166000526008602052604060002090565b01615fca565b9360005b8551811015612f7857600190612f716005612f5f8b67ffffffffffffffff166000526008602052604060002090565b01612f6a838a613a26565b51906168e8565b5001612f30565b509350937f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916612fe560019397612fca612fc58267ffffffffffffffff166000526008602052604060002090565b614411565b60405167ffffffffffffffff90911681529081906020820190565b0390a1019190939293612c29565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff861660045260246000fd5b346104985760206003193601126104985773ffffffffffffffffffffffffffffffffffffffff60043561305d8161049d565b613065614866565b163381146130d757807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b92919261310c614866565b60005b82811061318f5750505060005b81811061312857505050565b8067ffffffffffffffff613142611b9e60019486886133f9565b60006131628267ffffffffffffffff16600052600f602052604060002090565b55167f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8600080a20161311c565b806131a0611b9e6001938686613249565b7f56f77aeff2def50c8b5f5a0df3bab7183df09bf36c6feba496bb42551db77d7067ffffffffffffffff60206131d7858989613249565b0192613200846131fb8367ffffffffffffffff16600052600f602052604060002090565b613280565b613211604051928392169482613397565b0390a20161310f565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b91908110156132595760a0020190565b61321a565b356106d481610b4c565b63ffffffff81160361049857565b356106d481612734565b906060906132c3813561329281613268565b849063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b60208101356132d181613268565b67ffffffff0000000084549160201b16807fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff83161785557fffffffffffffffffffffffffffffffffffffffff0000000000000000ffffffff6bffffffff0000000000000000604085013561334481613268565b60401b16921617178355013561335981612734565b81547fffffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffff1690151560601b6cff00000000000000000000000016179055565b919091606080608083019463ffffffff81356133b281613268565b16845263ffffffff60208201356133c881613268565b16602085015263ffffffff60408201356133e181613268565b16604085015201356133f281612734565b1515910152565b91908110156132595760051b0190565b67ffffffffffffffff90600060606040516134238161107d565b828152826020820152826040820152015216600052600f602052604060002060ff604051916134518361107d565b5463ffffffff8116835263ffffffff8160201c16602084015263ffffffff8160401c16604084015260601c161515606082015290565b6040519061349482611008565b60008252565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610498570180359067ffffffffffffffff82116104985760200191813603831361049857565b6020818303126104985780359067ffffffffffffffff82116104985701604081830312610498576040519161351f83611029565b813567ffffffffffffffff8111610498578161353c918401611167565b8352602082013567ffffffffffffffff81116104985761355c9201611167565b602082015290565b9081602091031261049857516106d481612734565b90916135906106d493604084526040840190610682565b916020818403910152610682565b6040513d6000823e3d90fd5b356106d48161049d565b6135bc614866565b60005b8281106135fe5750907fc97f93e817584952f1c1d633f93784b8430f0633d002f9dcc4de4fe2780424d0916135f96040519283928361387b565b0390a1565b61361161360c828585613753565b613763565b805115801561372d575b6136b457906136ae826136a9611d4e6060600196519361369a602082015161369161364d604085015163ffffffff1690565b61368961365d6080870151151590565b9161366b60a0880151151590565b94613674611105565b9b8c5260208c015263ffffffff1660408b0152565b151588860152565b15156080870152565b015167ffffffffffffffff1690565b6137d6565b016135bf565b604080517f19d7585700000000000000000000000000000000000000000000000000000000815282516004820152602083015160248201529082015163ffffffff166044820152606082015167ffffffffffffffff16606482015260808201511515608482015260a090910151151560a482015260c490fd5b5067ffffffffffffffff61374c606083015167ffffffffffffffff1690565b161561361b565b91908110156132595760c0020190565b60c0813603126104985760a06040519161377c83611099565b8035835260208101356020840152604081013561379881613268565b604084015260608101356137ab81610b4c565b606084015260808101356137be81612734565b608084015201356137ce81612734565b60a082015290565b60026080918351815560208401516001820155019161382a63ffffffff604083015116849063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b6060810151835492909101517fffffffffffffffffffffffffffffffffffffffffffffffffffff0000ffffffff90921690151560201b64ff00000000161790151560281b65ff000000000016179055565b602080825281018390526040019160005b8181106138995750505090565b90919260c080600192863581526020870135602082015263ffffffff60408801356138c381613268565b16604082015267ffffffffffffffff60608801356138e081610b4c565b16606082015260808701356138f481612734565b1515608082015260a087013561390981612734565b151560a082015201940192910161388c565b9067ffffffffffffffff6106d492166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b906040519182815491828252602082019060005260206000209260005b81811061398a575050610b69925003836110b5565b845473ffffffffffffffffffffffffffffffffffffffff16835260019485019487945060209093019201613975565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b919082018092116125bb57565b906139ff8261197f565b613a0c60405191826110b5565b828152601f19613a1c829461197f565b0190602036910137565b80518210156132595760209160051b010190565b613a589067ffffffffffffffff16600052600e602052604060002090565b613a6181613958565b916005548015159182613b34575b5050613a79575090565b6001613a859101613958565b805180613a9157505090565b613aa2613aa79184959394516139e8565b6139f5565b9060005b8451811015613af25780613aec613ac761130560019489613a26565b613ad18387613a26565b9073ffffffffffffffffffffffffffffffffffffffff169052565b01613aab565b509160005b8151811015613b2d5780613b27613b1361130560019486613a26565b613ad1613b21848a516139e8565b87613a26565b01613af7565b5090925050565b101590503880613a6f565b90816020910312610498575190565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610eab57600091613bcb575090565b6106d4915060203d602011613bed575b613be581836110b5565b810190613b3f565b503d613bdb565b67ffffffffffffffff6106d491166000526007602052604060002054151590565b613c1d614866565b60208101519160005b8351811015613cbc5780613c3f61130560019387613a26565b613c66613c6173ffffffffffffffffffffffffffffffffffffffff831661133d565b616ad6565b613c72575b5001613c26565b60405173ffffffffffffffffffffffffffffffffffffffff9190911681527fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758090602090a138613c6b565b5091505160005b8151811015613da057613cd96113058284613a26565b9073ffffffffffffffffffffffffffffffffffffffff821615613d76577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef613d6d83613d45613d4061133d60019773ffffffffffffffffffffffffffffffffffffffff1690565b616a58565b5060405173ffffffffffffffffffffffffffffffffffffffff90911681529081906020820190565b0390a101613cc3565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050565b9190811015613259576060020190565b60405190613dc182611029565b60606020838281520152565b90604051613dda81611061565b608060ff600283958054855260018101546020860152015463ffffffff81166040850152818160201c161515606085015260281c161515910152565b601f8260209493601f19938186528686013760008582860101520116010190565b9160206106d4938181520191613e16565b90816020910312610498573590565b613e759067ffffffffffffffff16600052600e602052604060002090565b906003613e8460028401613958565b920190613e9082613958565b506005548015159182613f1b575b5050613ea8575090565b613eb190613958565b805180613ebd57505090565b613aa2613ece9184959394516139e8565b9060005b8451811015613ef45780613eee613ac761130560019489613a26565b01613ed2565b509160005b8151811015613b2d5780613f15613b1361130560019486613a26565b01613ef9565b101590503880613e9e565b90600182811c92168015613f6f575b6020831014613f4057565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691613f35565b9060405191826000825492613f8d84613f26565b8084529360018116908115613ff95750600114613fb2575b50610b69925003836110b5565b90506000929192526020600020906000915b818310613fdd575050906020610b699282010138613fa5565b6020919350806001915483858901015201910190918492613fc4565b60209350610b699592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138613fa5565b60409067ffffffffffffffff6106d495931681528160208201520191613e16565b6040519061406782611061565b60006080838281528260208201528260408201528260608201520152565b9060405161409281611061565b60806fffffffffffffffffffffffffffffffff6001839560ff8154848116875263ffffffff81871c16602088015260a01c1615156040860152015481808216166060850152821c16910152565b67ffffffffffffffff1660005260086020526106d46004604060002001613f79565b91908110156132595760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6181360301821215610498570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610498570180359067ffffffffffffffff821161049857602001918160051b3603831361049857565b818102929181159184041417156125bb57565b916141c2918354906000199060031b92831b921b19161790565b9055565b8181106141d1575050565b600081556001016141c6565b81519167ffffffffffffffff831161102457680100000000000000008311611024576020908254848455808510614254575b500190600052602060002060005b83811061422a5750505050565b600190602073ffffffffffffffffffffffffffffffffffffffff855116940193818401550161421d565b61426b9084600052858460002091820191016141c6565b3861420f565b90805180519067ffffffffffffffff821161102457680100000000000000008211611024576020908454838655808410614323575b500183600052602060002060005b8381106142f357505050509060036060836142d96020610b69960151600186016141dd565b6142ea6040820151600286016141dd565b015191016141dd565b6001906020614316855173ffffffffffffffffffffffffffffffffffffffff1690565b94019381840155016142b4565b61433a9086600052848460002091820191016141c6565b386142a6565b9160209082815201919060005b81811061435a5750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff87356143838161049d565b16815201940192910161434d565b9694926143cf946143b36143c1936106d49b999560808c5260808c0191614340565b9189830360208b0152614340565b918683036040880152614340565b926060818503910152614340565b80549060008155816143ed575050565b6000526020600020908101905b818110614405575050565b600081556001016143fa565b6005610b6991600081556000600182015560006002820155600060038201556004810161443e8154613f26565b908161444d575b5050016143dd565b81601f600093116001146144655750555b3880614445565b8183526020832061448091601f01861c8101906001016141c6565b808252602082209081548360011b906000198560031b1c19161790555561445e565b91908110156132595760051b810135907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee181360301821215610498570190565b9080601f830112156104985781356144f98161197f565b9261450760405194856110b5565b81845260208085019260051b820101918383116104985760208201905b83821061453357505050505090565b813567ffffffffffffffff81116104985760209161455687848094880101611167565b815201910190614524565b61012081360312610498576040519061457982611061565b61458281610b5e565b8252602081013567ffffffffffffffff8111610498576145a590369083016144e2565b602083015260408101359067ffffffffffffffff8211610498576145cf6145f09236908301611167565b60408401526145e136606083016127c2565b606084015260c03691016127c2565b608082015290565b8151815460208401516040850151608091821b73ffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffff000000000000000000000000000000000000000000909416939093179290921791151560a01b74ff000000000000000000000000000000000000000016919091178355606084015193810151901b7fffffffffffffffffffffffffffffffff0000000000000000000000000000000016921691909117600190910155565b9190601f81116146ce57505050565b610b69926000526020600020906020601f840160051c830193106146fa575b601f0160051c01906141c6565b90915081906146ed565b919091825167ffffffffffffffff81116110245761472c816147268454613f26565b846146bf565b6020601f82116001146147685781906141c293949560009261475d575b50506000198260011b9260031b1c19161790565b015190503880614749565b601f1982169061477d84600052602060002090565b9160005b8181106147b9575095836001959697106147a0575b505050811b019055565b015160001960f88460031b161c19169055388080614796565b9192602060018192868b015181550194019201614781565b614835614800610b699597969467ffffffffffffffff60a0951684526101006020850152610100840190610682565b9660408301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b01906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b73ffffffffffffffffffffffffffffffffffffffff60015416330361488757565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b91908110156132595760e0020190565b356106d48161273e565b919060005b8181106148dd5750509050565b6148e88183866148b1565b6148f18161325e565b6148fd611d7e82613bf4565b6123545781614993614a0d926149996020600197960161492561492036836127c2565b615db3565b6149936149468467ffffffffffffffff16600052600c602052604060002090565b91825461496661495d8263ffffffff9060801c1690565b63ffffffff1690565b159081614c00575b81614bd8575b81614bbd575b81614ba8575b5080614b99575b614b44575b36906127c2565b90616015565b6149c860808401916149ae61492036856127c2565b67ffffffffffffffff16600052600d602052604060002090565b9283546149df61495d8263ffffffff9060801c1690565b159081614b0d575b81614ae5575b81614aca575b81614ab5575b5080614aa6575b614a13575b5036906127c2565b016148d0565b614a2260a0614a5f92016148c1565b84906fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b82547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff000000000000000000000000000000001617835538614a05565b50614ab082613276565b614a00565b614ac4915060a01c60ff161590565b386149f9565b6fffffffffffffffffffffffffffffffff81161591506149f3565b90506fffffffffffffffffffffffffffffffff614b058987015460801c90565b1615906149ed565b90506fffffffffffffffffffffffffffffffff614b3c898701546fffffffffffffffffffffffffffffffff1690565b1615906149e7565b614b53614a22604089016148c1565b82547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff000000000000000000000000000000001617835561498c565b50614ba381613276565b614987565b614bb7915060a01c60ff161590565b38614980565b6fffffffffffffffffffffffffffffffff811615915061497a565b90506fffffffffffffffffffffffffffffffff614bf88c86015460801c90565b161590614974565b90506fffffffffffffffffffffffffffffffff614c2f8c8601546fffffffffffffffffffffffffffffffff1690565b16159061496e565b600060808201614c4c611d7e6107ea836135aa565b614e0f57506020820191614ced6020614c92614c6a612ca18761325e565b60801b7fffffffffffffffffffffffffffffffff000000000000000000000000000000001690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116600482015291829081906024820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610eab578391614df0575b50614dc857614d4c614d478461325e565b6162ba565b614d558361325e565b90614d6e611d7e60a08301936111c5610cf7868661349a565b614d88575050610b699291614d83915061325e565b616352565b614d92925061349a565b906120446040519283927f24eb47e500000000000000000000000000000000000000000000000000000000845260048401613e37565b6004827f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b614e09915060203d602011610ea457610e9681836110b5565b38614d36565b614e1b614e5c916135aa565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835273ffffffffffffffffffffffffffffffffffffffff16600452602490565b90fd5b916080830190614e74611d7e6107ea846135aa565b615005576020840193614e916020614c92614c6a612ca18961325e565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610eab57600091614fe6575b50614fbc57614eec614d478661325e565b614ef58561325e565b90614f0e611d7e60a08301936111c5610cf7868661349a565b614d8857505061ffff1615614faf5767ffffffffffffffff90614faa614f78614f7286611b9e614f626149ae7f41a8aa8df7945f0fb8ac5f7d88279638d9dc2ef9a6bf4ec9a53b80681b34aff7999a61325e565b89614f6c886135aa565b91616403565b926135aa565b946040519384931695836020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b0390a2565b50614d83610b699261325e565b7f53ad11d80000000000000000000000000000000000000000000000000000000060005260046000fd5b614fff915060203d602011610ea457610e9681836110b5565b38614edb565b611c3d615011836135aa565b7f961c9a4f0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b6040519061506082611029565b6000825260208201600081526020820151917fffffffff000000000000000000000000000000000000000000000000000000006028602483015160e01c92015193167fb148ea5f000000000000000000000000000000000000000000000000000000008114158061510a575b6150dd575063ffffffff1683525290565b7fa176027f0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b507f3047587c000000000000000000000000000000000000000000000000000000008114156150cc565b8051610118811061538e575060048101517f000000000000000000000000000000000000000000000000000000000000000063ffffffff811663ffffffff831603615355575050600881015190600c81015191608c82015191609081015193609482015160b88301519360f860d8850151940151916151b7895163ffffffff1690565b63ffffffff811663ffffffff84160361531b57507f000000000000000000000000000000000000000000000000000000000000000063ffffffff811663ffffffff8616036152e157506107d063ffffffff8916036152a7576107d063ffffffff82160361526e5750916152309593916020979593615b00565b91015180820361523e575050565b7f7be225b60000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b7f0389caa2000000000000000000000000000000000000000000000000000000006000526107d060045263ffffffff1660245260446000fd5b7f22e102a0000000000000000000000000000000000000000000000000000000006000526107d060045263ffffffff881660245260446000fd5b7f77e480260000000000000000000000000000000000000000000000000000000060005263ffffffff908116600452841660245260446000fd5b7fe366a1170000000000000000000000000000000000000000000000000000000060005263ffffffff908116600452821660245260446000fd5b7f960693cd0000000000000000000000000000000000000000000000000000000060005263ffffffff9081166004521660245260446000fd5b7f758b22cc0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b8051801561546057602003615423576153dd6020825183010160208301613b3f565b9060ff82116153ed575060ff1690565b612044906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352600483016106c3565b612044906040519182917f953576f70000000000000000000000000000000000000000000000000000000083526020600484018181520190610682565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff82116125bb57565b60ff16604d81116125bb57600a0a90565b80156154b8576000190490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b81156154b8570490565b907f000000000000000000000000000000000000000000000000000000000000000060ff811660ff83168181146155d757116155ac576155318282615486565b91604d60ff8416118015615593575b615559575050906155536106d49261549a565b90614195565b7fa9cb113d0000000000000000000000000000000000000000000000000000000060005260ff908116600452166024525060445260646000fd5b506155a56155a08461549a565b6154ab565b8411615540565b6155b68183615486565b91604d60ff841611615559575050906155d16106d49261549a565b906154e7565b5050505090565b60409067ffffffffffffffff6106d494931681528160208201520190610682565b9080511561168f578051602082012067ffffffffffffffff831692836000526008602052615634826005604060002001616a8d565b1561568857508161567c7f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea93615677614faa946000526009602052604060002090565b614704565b604051918291826106c3565b90506120446040519283927f393b8ad2000000000000000000000000000000000000000000000000000000008452600484016155de565b67ffffffffffffffff1660008181526007602052604090205490929190156157c157916157be60e09261578a856157167f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b97615db3565b84600052600860205261572d816040600020616015565b61573683615db3565b846000526008602052615750836002604060002001616015565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60009060808101615805611d7e6107ea836135aa565b6158f1575060208101906158236020614c92614c6a612ca18661325e565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610eab5784916158d2575b506158aa57610b69929160608261589061588b60406158a596016135aa565b616646565b61589c614d478461325e565b0135925061325e565b6166c8565b6004837f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b6158eb915060203d602011610ea457610e9681836110b5565b3861586c565b614e5c614e1b84926135aa565b6080810190615912611d7e6107ea846135aa565b61500557602081019061592f6020614c92614c6a612ca18661325e565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610eab57600091615a88575b50614fbc578061599061588b6040606094016135aa565b61599c614d478461325e565b01359261ffff81169081151580615a69575b15615a5857600b5461ffff16918211615a215750507f7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b291614faa614f78614f7284611b9e614f62615a0767ffffffffffffffff9861325e565b67ffffffffffffffff16600052600c602052604060002090565b7fe08f03ef0000000000000000000000000000000000000000000000000000000060005261ffff9081166004521660245260446000fd5b5050610b699291506158a59061325e565b50615a81615a7a600b5461ffff1690565b61ffff1690565b15156159ae565b615aa1915060203d602011610ea457610e9681836110b5565b38615979565b949290939163ffffffff90604051958260208801981688526040870152166060850152608084015260a083015260c0820152600060e08201526107d06101008201526101008152615afa610120826110b5565b51902090565b959263ffffffff8095929693604051978260208a019a168a526040890152166060870152608086015260a085015260c0840152600060e0840152166101008201526101008152615afa610120826110b5565b602081519101517fffffffff00000000000000000000000000000000000000000000000000000000604051927fb148ea5f00000000000000000000000000000000000000000000000000000000602085015260e01b1660248301526028820152602881526106d46048826110b5565b9060001982019182116125bb57565b919082039182116125bb57565b615be561405a565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff82511690602083019163ffffffff83511642034281116125bb57615c49906fffffffffffffffffffffffffffffffff60808701511690614195565b81018091116125bb57615c6f6fffffffffffffffffffffffffffffffff92918392616ac4565b161682524263ffffffff16905290565b60405160ff7f0000000000000000000000000000000000000000000000000000000000000000166020820152602081526106d46040826110b5565b80519060005b828110615ccc57505050565b600181018082116125bb575b838110615ce85750600101615cc0565b73ffffffffffffffffffffffffffffffffffffffff615d078385613a26565b5116615d1961133d6113058487613a26565b14615d2657600101615cd8565b611c3d615d366113058486613a26565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b610b699092919260608101936fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b805115615e575760408101516fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff615e17615e0260208501516fffffffffffffffffffffffffffffffff1690565b6fffffffffffffffffffffffffffffffff1690565b911611615e215750565b612044906040519182917f8020d12400000000000000000000000000000000000000000000000000000000835260048301615d78565b6fffffffffffffffffffffffffffffffff615e8560408301516fffffffffffffffffffffffffffffffff1690565b1615801590615ecc575b615e965750565b612044906040519182917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301615d78565b50615eed615e0260208301516fffffffffffffffffffffffffffffffff1690565b1515615e8f565b15615efb57565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b604051906006548083528260208101600660005260206000209260005b818110615fb1575050610b69925003836110b5565b8454835260019485019487945060209093019201615f9c565b906040519182815491828252602082019060005260206000209260005b818110615ffc575050610b69925003836110b5565b8454835260019485019487945060209093019201615fe7565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19916161eb6135f992805461605d61605761495d8363ffffffff9060801c1690565b42615bd0565b90816161f7575b50506161a5600161608860208601516fffffffffffffffffffffffffffffffff1690565b926161136160d6615e026fffffffffffffffffffffffffffffffff6160bd85546fffffffffffffffffffffffffffffffff1690565b166fffffffffffffffffffffffffffffffff8816616ac4565b82906fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b6161666161208751151590565b82547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178255565b019182906fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b604083015181546fffffffffffffffffffffffffffffffff1660809190911b7fffffffffffffffffffffffffffffffff0000000000000000000000000000000016179055565b60405191829182615d78565b615e026160d6916fffffffffffffffffffffffffffffffff61626b616272958261626460018a0154928261625d616256616240876fffffffffffffffffffffffffffffffff1690565b996fffffffffffffffffffffffffffffffff1690565b9560801c90565b1690614195565b91166139e8565b9116616ac4565b80547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff00000000000000000000000000000000161781553880616064565b67ffffffffffffffff166162db816000526007602052604060002054151590565b156163255750336000526011602052604060002054156162f757565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fa9902c7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b67ffffffffffffffff7f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c911691826000526008602052806163d3600260406000200173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391616403565b6040805173ffffffffffffffffffffffffffffffffffffffff909216825260208201929092529081908101614faa565b8054939290919060ff60a086901c1615801561663e575b616637576164396fffffffffffffffffffffffffffffffff8616615e02565b906001840195865461647361605761495d616466615e02856fffffffffffffffffffffffffffffffff1690565b9460801c63ffffffff1690565b806165a3575b505083811061655857508282106164d95750610b6993945061649e91615e0291615bd0565b6fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b90616510611c3d9361650b6164fc846164f6615e028c5460801c90565b93615bd0565b61650583615bc1565b906139e8565b6154e7565b7fd0c8d23a0000000000000000000000000000000000000000000000000000000060005260045260245273ffffffffffffffffffffffffffffffffffffffff16604452606490565b7f1a76572a00000000000000000000000000000000000000000000000000000000600052600452602483905273ffffffffffffffffffffffffffffffffffffffff1660445260646000fd5b82859293951161660d576165bd615e026165c49460801c90565b918561698d565b84547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178555913880616479565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050509050565b50811561641a565b7f000000000000000000000000000000000000000000000000000000000000000061666e5750565b73ffffffffffffffffffffffffffffffffffffffff168060005260036020526040600020541561669b5750565b7fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b67ffffffffffffffff7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944911691826000526008602052806163d3604060002073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391616403565b80548210156132595760005260206000200190600090565b805480156167865760001901906167758282616746565b60001982549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600081815260036020526040902054908115616856576000198201908282116125bb576002549260001984019384116125bb578383600095616815950361681b575b505050616804600261675e565b600390600052602052604060002090565b55600190565b6168046168479161683d61683361684d956002616746565b90549060031b1c90565b9283916002616746565b906141a8565b553880806167f7565b5050600090565b600081815260076020526040902054908115616856576000198201908282116125bb576006549260001984019384116125bb57838360009561681595036168bd575b5050506168ac600661675e565b600790600052602052604060002090565b6168ac616847916168d56168336168df956006616746565b9283916006616746565b5538808061689f565b60018101918060005282602052604060002054928315156000146169845760001984018481116125bb5783549360001985019485116125bb5760009585836168159761693c950361694b575b50505061675e565b90600052602052604060002090565b61696b6168479161696261683361697b9588616746565b92839187616746565b8590600052602052604060002090565b55388080616934565b50505050600090565b926169989192614195565b81018091116125bb576106d491616ac4565b8054906801000000000000000082101561102457816169d19160016141c294018155616746565b81939154906000199060031b92831b921b19161790565b600081815260036020526040902054616a1d57616a068160026169aa565b600254906000526003602052604060002055600190565b50600090565b600081815260076020526040902054616a1d57616a418160066169aa565b600654906000526007602052604060002055600190565b600081815260116020526040902054616a1d57616a768160106169aa565b601054906000526011602052604060002055600190565b60008281526001820160205260409020546168565780616aaf836001936169aa565b80549260005201602052604060002055600190565b9080821015616ad1575090565b905090565b600081815260116020526040902054908115616856576000198201908282116125bb576010549260001984019384116125bb5783836168159460009603616b36575b505050616b25601061675e565b601190600052602052604060002090565b616b2561684791616b4e616833616b58956010616746565b9283916010616746565b55388080616b18565b91929015616bdc5750815115616b75575090565b3b15616b7e5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015616bef5750805190602001fd5b612044906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352600483016106c356fea164736f6c634300081a000a",
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) GetRequiredInboundCCVs(opts *bind.CallOpts, arg0 common.Address, sourceChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "getRequiredInboundCCVs", arg0, sourceChainSelector, amount, arg3, arg4)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GetRequiredInboundCCVs(arg0 common.Address, sourceChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetRequiredInboundCCVs(&_USDCTokenPoolCCTPV2.CallOpts, arg0, sourceChainSelector, amount, arg3, arg4)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) GetRequiredInboundCCVs(arg0 common.Address, sourceChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetRequiredInboundCCVs(&_USDCTokenPoolCCTPV2.CallOpts, arg0, sourceChainSelector, amount, arg3, arg4)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) GetRequiredOutboundCCVs(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "getRequiredOutboundCCVs", arg0, destChainSelector, amount, arg3, arg4)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GetRequiredOutboundCCVs(arg0 common.Address, destChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetRequiredOutboundCCVs(&_USDCTokenPoolCCTPV2.CallOpts, arg0, destChainSelector, amount, arg3, arg4)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) GetRequiredOutboundCCVs(arg0 common.Address, destChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetRequiredOutboundCCVs(&_USDCTokenPoolCCTPV2.CallOpts, arg0, destChainSelector, amount, arg3, arg4)
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
	RemoteChainSelector    uint64
	OutboundCCVs           []common.Address
	AdditionalOutboundCCVs []common.Address
	InboundCCVs            []common.Address
	AdditionalInboundCCVs  []common.Address
	Raw                    types.Log
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
	return common.HexToHash("0x56f77aeff2def50c8b5f5a0df3bab7183df09bf36c6feba496bb42551db77d70")
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

	GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRequiredInboundCCVs(opts *bind.CallOpts, arg0 common.Address, sourceChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error)

	GetRequiredOutboundCCVs(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error)

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
