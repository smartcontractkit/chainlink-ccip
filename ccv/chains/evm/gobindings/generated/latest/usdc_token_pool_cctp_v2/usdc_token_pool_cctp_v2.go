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

type TokenPoolChainUpdate struct {
	RemoteChainSelector       uint64
	RemotePoolAddresses       [][]byte
	RemoteTokenAddress        []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
}

type TokenPoolRateLimitConfigArgs struct {
	RemoteChainSelector       uint64
	CustomBlockConfirmation   bool
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"internalType\":\"contract ITokenMessenger\"},{\"name\":\"cctpMessageTransmitterProxy\",\"type\":\"address\",\"internalType\":\"contract CCTPMessageTransmitterProxy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"advancedPoolHooks\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"FINALITY_THRESHOLD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_FEE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_USDC_MESSAGE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"struct AuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"disableTokenTransferFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAdvancedPoolHooks\",\"inputs\":[],\"outputs\":[{\"name\":\"advancedPoolHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDomain\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct USDCTokenPool.Domain\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useLegacySourcePoolDataFormat\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinBlockConfirmation\",\"inputs\":[],\"outputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_localDomainIdentifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_messageTransmitterProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract CCTPMessageTransmitterProxy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_supportedUSDCVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_tokenMessenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract ITokenMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDomains\",\"inputs\":[{\"name\":\"domains\",\"type\":\"tuple[]\",\"internalType\":\"struct USDCTokenPool.DomainUpdate[]\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useLegacySourcePoolDataFormat\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinBlockConfirmation\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.RateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"updateAdvancedPoolHooks\",\"inputs\":[{\"name\":\"newHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AdvancedPoolHooksUpdated\",\"inputs\":[{\"name\":\"oldHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"},{\"name\":\"newHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DomainsSet\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"struct USDCTokenPool.DomainUpdate[]\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useLegacySourcePoolDataFormat\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinBlockConfirmationSet\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustomBlockConfirmationsNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDepositHash\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"got\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidDestinationDomain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidDomain\",\"inputs\":[{\"name\":\"domain\",\"type\":\"tuple\",\"internalType\":\"struct USDCTokenPool.DomainUpdate\",\"components\":[{\"name\":\"allowedCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"domainIdentifier\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"useLegacySourcePoolDataFormat\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"type\":\"error\",\"name\":\"InvalidExecutionFinalityThreshold\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmation\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidMinFinalityThreshold\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidNonce\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"got\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourceDomain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenMessengerVersion\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"got\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidTransmitterInProxy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnknownDomain\",\"inputs\":[{\"name\":\"domain\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"UnlockingUSDCFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x61016080604052346104fc5760c0816167d6803803809161002082856106fe565b8339810103126104fc5780516001600160a01b038116908181036104fc576020830151916001600160a01b0383168084036104fc5760408501516001600160a01b038116959093908685036104fc5761007b60608201610721565b61009360a061008c60808501610721565b9301610721565b91602096604051996100a5898c6106fe565b60008b52600036813733156106ed57600180546001600160a01b03191633179055801580156106dc575b80156106cb575b6106ba57600492899260805260c0526040519283809263313ce56760e01b82525afa8091600091610683575b509061065f575b50600660a052600380546001600160a01b039283166001600160a01b031991821617909155600280549390921692169190911790556040519261014c85856106fe565b60008452600036813760408051979088016001600160401b03811189821017610649576040528752838588015260005b84518110156101e3576001906001600160a01b0361019a8288610751565b5116876101a682610793565b6101b3575b50500161017c565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a138876101ab565b508493508587519260005b845181101561025f576001600160a01b036102098287610751565b511690811561024e577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef8883610240600195610891565b50604051908152a1016101ee565b6342bcdf7f60e11b60005260046000fd5b5090859185600160e052841561063857604051632c12192160e01b81528481600481895afa90811561050857600091610603575b5060405163054fd4d560e41b81526001600160a01b039190911691908581600481865afa908115610508576000916105e6575b5063ffffffff8060e051169116908082036105cf575050604051639cdbb18160e01b815285816004818a5afa908115610508576000916105b2575b5063ffffffff8060e0511691169080820361059b57505084600491604051928380926367e0ed8360e11b82525afa8015610508578291600091610552575b506001600160a01b03160361054157600492849261010052610120526040519283809263234d8e3d60e21b82525afa90811561050857600091610514575b50610140526080516101005160405163095ea7b360e01b81526001600160a01b039182166004820152600019602482015291839183916044918391600091165af18015610508576104be575b506000805160206167b683398151915291604051908152a1604051615ec490816108f2823960805181818161050c0152818161058801528181610a64015281816114ca01528181611c0101528181615618015281816156c60152818161574401526157c2015260a05181818161079c01528181614b2101528181614bb101526150a1015260c051818181611b4d015281816125f6015281816144ea015281816146890152614ece015260e0518181816104bb01526147f5015261010051818181610edd015261147b0152610120518181816109cc0152611335015261014051818181610fde015281816115c6015261487b0152f35b8181813d8311610501575b6104d381836106fe565b810103126104fc57519182151583036104fc5791506000805160206167b68339815191526103c9565b600080fd5b503d6104c9565b6040513d6000823e3d90fd5b6105349150823d841161053a575b61052c81836106fe565b810190610735565b8361037d565b503d610522565b632a32133b60e11b60005260046000fd5b9091508581813d8311610594575b61056a81836106fe565b810103126105905751906001600160a01b038216820361058d575081908761033f565b80fd5b5080fd5b503d610560565b633785f8f160e01b60005260045260245260446000fd5b6105c99150863d881161053a5761052c81836106fe565b87610301565b63960693cd60e01b60005260045260245260446000fd5b6105fd9150863d881161053a5761052c81836106fe565b876102c6565b90508481813d8311610631575b61061a81836106fe565b810103126104fc5761062b90610721565b86610293565b503d610610565b6306b7c75960e31b60005260046000fd5b634e487b7160e01b600052604160045260246000fd5b60ff1660068114610109576332ad3e0760e11b600052600660045260245260446000fd5b8881813d83116106b3575b61069881836106fe565b8101031261059057519060ff8216820361058d575038610102565b503d61068e565b630a64406560e11b60005260046000fd5b506001600160a01b038316156100d6565b506001600160a01b038516156100cf565b639b15e16f60e01b60005260046000fd5b601f909101601f19168101906001600160401b0382119082101761064957604052565b51906001600160a01b03821682036104fc57565b908160209103126104fc575163ffffffff811681036104fc5790565b80518210156107655760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b80548210156107655760005260206000200190600090565b6000818152600d6020526040902054801561088a57600019810181811161087457600c5460001981019190821161087457808203610823575b505050600c54801561080d57600019016107e781600c61077b565b8154906000199060031b1b19169055600c55600052600d60205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61085c61083461084593600c61077b565b90549060031b1c928392600c61077b565b819391549060031b91821b91600019901b19161790565b9055600052600d6020526040600020553880806107cc565b634e487b7160e01b600052601160045260246000fd5b5050600090565b80600052600d602052604060002054156000146108eb57600c5468010000000000000000811015610649576108d2610845826001859401600c55600c61077b565b9055600c5490600052600d602052604060002055600190565b5060009056fe6080604052600436101561001257600080fd5b60003560e01c806301ffc9a7146102e7578063181f5a77146102e2578063212a052e146102dd57806321df0da7146102d8578063240028e8146102d35780632422ac45146102ce5780632451a627146102c957806324f65ee7146102c45780632c063404146102bf57806337a3210d146102ba57806339077537146102b5578063489a68f2146102b05780634ac8bd5f146102ab5780634c5ef0ed146102a65780634e921c30146102a15780636155cda01461029c57806362ddd3c4146102975780636b716b0d146102925780637437ff9f1461028d57806379ba5097146102885780638926f54f1461028357806389720a621461027e5780638da5cb5b1461027957806391a2749a1461027457806398db96431461026f5780639a4575b91461026a578063a42a7b8b14610265578063acfecf9114610260578063b1c71c651461025b578063b794658014610256578063bc063e1a14610251578063bfeffd3f1461024c578063c4bffe2b14610247578063c7230a6014610242578063c8c8fd191461023d578063d8aa3f4014610238578063da4b05e714610233578063dc04fa1f1461022e578063dc0bd97114610229578063dcbd41bc14610224578063dfadfa351461021f578063e8a1da171461021a578063f2fde38b14610215578063fa41d79c146102105763ff8e03f31461020b57600080fd5b612d2e565b612d09565b612c33565b6128ac565b6127e2565b61261a565b6125c9565b612346565b6122f8565b6121af565b61211d565b611fa3565b611efd565b611e20565b611e04565b611dcd565b6119f4565b6118ac565b6117c8565b6113a2565b611308565b611274565b6111c1565b61114d565b61110e565b611043565b611002565b610fc1565b610f3e565b610eb0565b610e15565b610dcc565b610bec565b610b1c565b6108e7565b6108a4565b610805565b610782565b61071c565b6105ec565b61054e565b6104df565b61049e565b61043b565b346103ea5760206003193601126103ea576004357fffffffff0000000000000000000000000000000000000000000000000000000081168091036103ea577faff2afbf0000000000000000000000000000000000000000000000000000000081149081156103c0575b8115610396575b811561036c575b506040519015158152602090f35b7f01ffc9a7000000000000000000000000000000000000000000000000000000009150143861035e565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150610357565b7f331710310000000000000000000000000000000000000000000000000000000081149150610350565b600080fd5b60009103126103ea57565b919082519283825260005b848110610426575050601f19601f8460006020809697860101520116010190565b80602080928401015182828601015201610405565b346103ea5760006003193601126103ea5761049a604080519061045e8183610d0b565b601d82527f55534443546f6b656e506f6f6c43435450563220312e372e302d6465760000006020830152519182916020835260208301906103fa565b0390f35b346103ea5760006003193601126103ea57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346103ea5760006003193601126103ea57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b73ffffffffffffffffffffffffffffffffffffffff8116036103ea57565b346103ea5760206003193601126103ea5760206105ae60043561057081610530565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691161490565b6040519015158152f35b67ffffffffffffffff8116036103ea57565b35906105d5826105b8565b565b801515036103ea57565b35906105d5826105d7565b346103ea5760406003193601126103ea576101406106ca610624600435610612816105b8565b6024359061061f826105d7565b612edb565b61067a60409392935180946fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b60a08301906fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565bf35b602060408183019282815284518094520192019060005b8181106106f05750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016106e3565b346103ea5760006003193601126103ea57604051600c548082526020820190600c60005260206000209060005b81811061076c5761049a8561076081870382610d0b565b604051918291826106cc565b8254845260209093019260019283019201610749565b346103ea5760006003193601126103ea57602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b61ffff8116036103ea57565b35906105d5826107c0565b9181601f840112156103ea5782359167ffffffffffffffff83116103ea57602083818601950101116103ea57565b346103ea5760c06003193601126103ea57610821600435610530565b60243561082d816105b8565b610838606435610530565b608435610844816107c0565b60a4359167ffffffffffffffff83116103ea5761ffff61087a63ffffffff9384936108738736906004016107d7565b5050612fcd565b6040805195865293909716602085015294169082015291166060820152901515608082015260a090f35b346103ea5760006003193601126103ea57602073ffffffffffffffffffffffffffffffffffffffff60035416604051908152f35b90816101009103126103ea5790565b346103ea5760206003193601126103ea5760043567ffffffffffffffff81116103ea576109189036906004016108d8565b6109206130d8565b506060810135610930818361449e565b6109b1602061094d61094560e08601866130e5565b810190613136565b61097661096f61096a61096360c08901896130e5565b3691610d77565b614738565b82516147e2565b8181519101519060405193849283927f57ecfd28000000000000000000000000000000000000000000000000000000008452600484016131c4565b0381600073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af1908115610b1757600091610ae8575b5015610abe57817ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc067ffffffffffffffff610a496040610a42602061049a98016131f5565b94016131ff565b6040805173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000081168252336020830152929092169082015260608101859052921691608090a2610aab610d2e565b8190526040519081529081906020820190565b7fbf969f220000000000000000000000000000000000000000000000000000000060005260046000fd5b610b0a915060203d602011610b10575b610b028183610d0b565b8101906131af565b386109fd565b503d610af8565b6131e9565b346103ea5760406003193601126103ea5760043567ffffffffffffffff81116103ea57610b5061049a9136906004016108d8565b60243590610b5d826107c0565b6000604051610b6b81610c7a565b52610b9d610b956060830135610b8f610b8a61096360c08701876130e5565b614a78565b90614bae565b92838361463e565b7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc067ffffffffffffffff610a4960206040850194610bdb8635610530565b013593610be7856105b8565b6131ff565b346103ea5760206003193601126103ea5760043567ffffffffffffffff81116103ea57366023820112156103ea57806004013567ffffffffffffffff81116103ea5736602460c08302840101116103ea576024610c499201613209565b005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff821117610c9657604052565b610c4b565b6040810190811067ffffffffffffffff821117610c9657604052565b60e0810190811067ffffffffffffffff821117610c9657604052565b60a0810190811067ffffffffffffffff821117610c9657604052565b60c0810190811067ffffffffffffffff821117610c9657604052565b90601f601f19910116810190811067ffffffffffffffff821117610c9657604052565b604051906105d5602083610d0b565b604051906105d5604083610d0b565b604051906105d560a083610d0b565b67ffffffffffffffff8111610c9657601f01601f191660200190565b929192610d8382610d5b565b91610d916040519384610d0b565b8294818452818301116103ea578281602093846000960137010152565b9080601f830112156103ea57816020610dc993359101610d77565b90565b346103ea5760406003193601126103ea57600435610de9816105b8565b60243567ffffffffffffffff81116103ea57602091610e0f6105ae923690600401610dae565b906135bd565b346103ea5760206003193601126103ea577fa7f8dbba8cdb126ce4a0e7939ec58e0161b70d808b585dd651d68e59d27e11fb6020600435610e55816107c0565b610e5d614c9b565b6003547fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff75ffff00000000000000000000000000000000000000008360a01b1691161760035561ffff60405191168152a1005b346103ea5760006003193601126103ea57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b9060406003198301126103ea57600435610f1a816105b8565b916024359067ffffffffffffffff82116103ea57610f3a916004016107d7565b9091565b346103ea57610f4c36610f01565b610f57929192614c9b565b67ffffffffffffffff8216610f79816000526007602052604060002054151590565b15610f945750610c4992610f8e913691610d77565b90614d07565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346103ea5760006003193601126103ea57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346103ea5760006003193601126103ea57600254600a546040805173ffffffffffffffffffffffffffffffffffffffff938416815292909116602083015290f35b346103ea5760006003193601126103ea5760005473ffffffffffffffffffffffffffffffffffffffff811633036110e4577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346103ea5760206003193601126103ea5760206105ae67ffffffffffffffff600435611139816105b8565b166000526007602052604060002054151590565b346103ea5760c06003193601126103ea5760043561116a81610530565b60243590611177826105b8565b604435606435611186816107c0565b60843567ffffffffffffffff81116103ea576111a69036906004016107d7565b9160a4359360028510156103ea5761049a9661076096613750565b346103ea5760006003193601126103ea57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b67ffffffffffffffff8111610c965760051b60200190565b9080601f830112156103ea578135611224816111f5565b926112326040519485610d0b565b81845260208085019260051b8201019283116103ea57602001905b82821061125a5750505090565b60208091833561126981610530565b81520191019061124d565b346103ea5760206003193601126103ea5760043567ffffffffffffffff81116103ea57604060031982360301126103ea576040516112b181610c9b565b816004013567ffffffffffffffff81116103ea576112d5906004369185010161120d565b8152602482013567ffffffffffffffff81116103ea57610c499260046112fe923692010161120d565b6020820152613826565b346103ea5760006003193601126103ea57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b908160a09103126103ea5790565b610dc991602061138083516040845260408401906103fa565b9201519060208184039101526103fa565b906020610dc9928181520190611367565b346103ea5760206003193601126103ea5760043567ffffffffffffffff81116103ea576113d3903690600401611359565b6113db6139cf565b506113f76113e882614dcc565b6113f0612e41565b5082614e81565b6020810161142961142461140a836131f5565b67ffffffffffffffff16600052600e602052604060002090565b6139e8565b9161143e61143a6060850151151590565b1590565b61170457602061144e82806130e5565b9050036116c057602083015180156116a457925b606073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169201359260408201926114b2845163ffffffff1690565b9273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016938151833b156103ea576040517fd04857b00000000000000000000000000000000000000000000000000000000081526004810189905263ffffffff9290921660248301526044820189905273ffffffffffffffffffffffffffffffffffffffff861660648301526084820152600060a482018190526107d060c4830152909492859060e490829084905af18015610b175761161b61166c966115fa7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10948661049a9c6116679a67ffffffffffffffff97611689575b506115f07f0000000000000000000000000000000000000000000000000000000000000000955163ffffffff1690565b9251928d86614f80565b611611611605610d3d565b63ffffffff9093168352565b602082015261502b565b9661165f611628866131f5565b6040805173ffffffffffffffffffffffffffffffffffffffff90971687523360208801528601929092529116929081906060820190565b0390a26131f5565b613be9565b90611675610d3d565b918252602082015260405191829182611391565b80611698600061169e93610d0b565b806103ef565b386115c0565b506116ba6116b282806130e5565b810190613a42565b92611462565b806116ca916130e5565b906117006040519283927fa3c8cf0900000000000000000000000000000000000000000000000000000000845260048401613a31565b0390fd5b611746611710836131f5565b7fd201c48a0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b6000fd5b602081016020825282518091526040820191602060408360051b8301019401926000915b83831061177d57505050505090565b90919293946020806117b9837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0866001960301875289516103fa565b9701930193019193929061176e565b346103ea5760206003193601126103ea5767ffffffffffffffff6004356117ee816105b8565b16600052600860205261180760056040600020016154dc565b805190601f1961182f611819846111f5565b936118276040519586610d0b565b8085526111f5565b0160005b81811061189b57505060005b815181101561188d578061187161186c61185b60019486613a51565b516000526009602052604060002090565b613ab8565b61187b8286613a51565b526118868185613a51565b500161183f565b6040518061049a858261174a565b806060602080938701015201611833565b346103ea576118ba36610f01565b6118c5929192614c9b565b67ffffffffffffffff8216916118eb61143a846000526007602052604060002054151590565b6119a15761192e61143a60056119158467ffffffffffffffff166000526008602052604060002090565b01611921368689610d77565b6020815191012090615901565b61196a57507f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76919261196560405192839283613a31565b0390a2005b61170084926040519384937f74f23c7c00000000000000000000000000000000000000000000000000000000855260048501613b78565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b9291906119ef602091604086526040860190611367565b930152565b346103ea5760606003193601126103ea5760043567ffffffffffffffff81116103ea57611a25903690600401611359565b602435611a31816107c0565b6044359067ffffffffffffffff82116103ea57611a55611a729236906004016107d7565b9290611a5f6139cf565b50611a6a8386614e0f565b933691610d77565b5060808301611a8661143a610570836131ff565b611d6e57506020830192611b346020611ad9611ab1611aa4886131f5565b67ffffffffffffffff1690565b60801b7fffffffffffffffffffffffffffffffff000000000000000000000000000000001690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116600482015291829081906024820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610b1757600091611d4f575b50611d2557606090611b97611b92866131f5565b615527565b013590611ba48383613bdc565b9061ffff8116908115611d065760035460a01c61ffff169161ffff8316908115611cdc5710611ca5575050611c7e9261166792611bec611bf193611be7886131f5565b61576c565b613bdc565b92611bfb816131f5565b604080517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16815233602082015290810186905267ffffffffffffffff91909116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1090806060810161165f565b611c8661509a565b611c8e610d3d565b918252602082015261049a604051928392836119d8565b7f7911d95b0000000000000000000000000000000000000000000000000000000060005261ffff9081166004521660245260446000fd5b7f98d50fd70000000000000000000000000000000000000000000000000000000060005260046000fd5b5050611c7e9261166792611bec611bf193611d20886131f5565b6156ee565b7f53ad11d80000000000000000000000000000000000000000000000000000000060005260046000fd5b611d68915060203d602011610b1057610b028183610d0b565b38611b7e565b611d7a611746916131ff565b7f961c9a4f0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b906020610dc99281815201906103fa565b346103ea5760206003193601126103ea5761049a611df0600435611667816105b8565b6040519182916020835260208301906103fa565b346103ea5760006003193601126103ea57602060405160008152f35b346103ea5760206003193601126103ea57600435611e3d81610530565b611e45614c9b565b7fffffffffffffffffffffffff00000000000000000000000000000000000000006003547fbaff46844acf36d6ee996f489a1a288709c4542bd33cd557770afd267d581209604073ffffffffffffffffffffffffffffffffffffffff81519581851687521694856020820152a11617600355005b602060408183019282815284518094520192019060005b818110611edd5750505090565b825167ffffffffffffffff16845260209384019390920191600101611ed0565b346103ea5760006003193601126103ea57611f16615491565b805190601f19611f28611819846111f5565b0136602084013760005b8151811015611f64578067ffffffffffffffff611f5160019385613a51565b5116611f5d8286613a51565b5201611f32565b6040518061049a8582611eb9565b9181601f840112156103ea5782359167ffffffffffffffff83116103ea576020808501948460051b0101116103ea57565b346103ea5760406003193601126103ea5760043567ffffffffffffffff81116103ea57611fd4903690600401611f72565b9060243591611fe283610530565b611fea614c9b565b60005b818110611ff657005b612020612007610be7838587613efd565b73ffffffffffffffffffffffffffffffffffffffff1690565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290919073ffffffffffffffffffffffffffffffffffffffff831690602081602481855afa908115610b175760019488916000936120ed575b5082612095575b5050505001611fed565b6120a09183916159a6565b60405190815273ffffffffffffffffffffffffffffffffffffffff8716907f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e90602090a33880868161208b565b61210f91935060203d8111612116575b6121078183610d0b565b810190614a69565b9138612084565b503d6120fd565b346103ea5760006003193601126103ea5760206040516101188152f35b6105d59092919260c08060e083019563ffffffff815116845263ffffffff602082015116602085015263ffffffff604082015116604085015263ffffffff606082015116606085015261ffff60808201511660808501526121a660a082015160a086019061ffff169052565b01511515910152565b346103ea5760806003193601126103ea576121cb600435610530565b6024356121d7816105b8565b6121e26044356107c0565b6064359067ffffffffffffffff82116103ea5761220c67ffffffffffffffff9236906004016107d7565b5050600060c060405161221e81610cb7565b8281528260208201528260408201528260608201528260808201528260a0820152015216600052600b60205261049a60406000206122ec6122e36040519261226584610cb7565b5463ffffffff8116845263ffffffff8160201c16602085015261229a63ffffffff8260401c16604086019063ffffffff169052565b6122b5606082901c63ffffffff1663ffffffff166060860152565b6122cc608082901c61ffff1661ffff166080860152565b61ffff609082901c1660a085015260a01c60ff1690565b151560c0830152565b6040519182918261213a565b346103ea5760006003193601126103ea5760206040516107d08152f35b9181601f840112156103ea5782359167ffffffffffffffff83116103ea576020808501948460081b0101116103ea57565b346103ea5760406003193601126103ea5760043567ffffffffffffffff81116103ea57612377903690600401612315565b9060243567ffffffffffffffff81116103ea57612398903690600401611f72565b9190926123a3614c9b565b60005b8181106124285750505060005b8181106123bc57005b8067ffffffffffffffff6123db6123d66001948688613efd565b6131f5565b60006123fb8267ffffffffffffffff16600052600b602052604060002090565b55167f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8600080a2016123b3565b6124366123d6828486613c0b565b612441828486613c0b565b90602082019161245661143a60e08301613c1b565b6125915760a0810161271061247461246d83613c25565b61ffff1690565b1015612555575060c00161271061248d61246d83613c25565b101561255557506124a96124a083613c2f565b63ffffffff1690565b1561251e57907ffae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa104167ffffffffffffffff83612504846124ff6001989767ffffffffffffffff16600052600b602052604060002090565b613c39565b612515604051928392169482613e50565b0390a2016123a6565b7f123322650000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b61256161174691613c25565b7f95f3517a0000000000000000000000000000000000000000000000000000000060005261ffff16600452602490565b7f123322650000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b346103ea5760006003193601126103ea57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346103ea5760206003193601126103ea5760043567ffffffffffffffff81116103ea5761264b903690600401612315565b906126546150d5565b6000915b80831061266157005b61266c838284613c0b565b92612676846131f5565b9361268361143a866135fa565b6127aa577f41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb67ffffffffffffffff600194959661274d612735602086016126c981613c1b565b15612758576127026126ef8567ffffffffffffffff166000526004602052604060002090565b6126fc3660408b01613f2a565b90615188565b6127306127238567ffffffffffffffff166000526005602052604060002090565b6126fc3660a08b01613f2a565b613c1b565b946040519384931695604060a0830192019084613fca565b0390a2019190612658565b6127796126ef8567ffffffffffffffff166000526008602052604060002090565b612730600261279c8667ffffffffffffffff166000526008602052604060002090565b016126fc3660a08b01613f2a565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff851660045260246000fd5b346103ea5760206003193601126103ea5767ffffffffffffffff600435612808816105b8565b612810612e56565b5016600052600e60205261049a604060002060ff60026040519261283384610cd3565b8054845260018101546020850152015463ffffffff81166040840152818160201c161515606084015260281c16151560808201526040519182918291909160808060a0830194805184526020810151602085015263ffffffff604082015116604085015260608101511515606085015201511515910152565b346103ea5760406003193601126103ea5760043567ffffffffffffffff81116103ea576128dd903690600401611f72565b60243567ffffffffffffffff81116103ea576128fd903690600401611f72565b919092612908614c9b565b6000915b808310612adf5750505060005b81811061292257005b6129356129308284866140ff565b6141be565b604081019182515115612ab55761296261143a61295d611aa4855167ffffffffffffffff1690565b615acc565b612a6a5761299861297e839695965167ffffffffffffffff1690565b67ffffffffffffffff166000526008602052604060002090565b9260608301906129a9825186615188565b6129ca60808501956129bf875160028301615188565b60048351910161429a565b602084019560005b87518051821015612a0d5790612a07600192612a00836129fa8b5167ffffffffffffffff1690565b92613a51565b5190614d07565b016129d2565b50509550612a5e600195612a4b7f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c295965167ffffffffffffffff1690565b9251935190519060405194859485614367565b0390a101919091612919565b611746612a7f835167ffffffffffffffff1690565b7f1d5ad3c50000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b909192612af06123d6858486613efd565b94612b0761143a67ffffffffffffffff8816615859565b612bfb57612b346005612b2e8867ffffffffffffffff166000526008602052604060002090565b016154dc565b9360005b8551811015612b8057600190612b796005612b678b67ffffffffffffffff166000526008602052604060002090565b01612b72838a613a51565b5190615901565b5001612b38565b509350937f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916612bed60019397612bd2612bcd8267ffffffffffffffff166000526008602052604060002090565b61406e565b60405167ffffffffffffffff90911681529081906020820190565b0390a101919093929361290c565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff861660045260246000fd5b346103ea5760206003193601126103ea5773ffffffffffffffffffffffffffffffffffffffff600435612c6581610530565b612c6d614c9b565b16338114612cdf57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346103ea5760006003193601126103ea57602061ffff60035460a01c16604051908152f35b346103ea5760406003193601126103ea57600435612d4b81610530565b602435612d5781610530565b612d5f614c9b565b73ffffffffffffffffffffffffffffffffffffffff8216918215612ab5577f22a0dbb8195755fbfc99667a86ae684c568e9dfbb1eccf7f90084e6166447970927fffffffffffffffffffffffff0000000000000000000000000000000000000000600254161760025573ffffffffffffffffffffffffffffffffffffffff82167fffffffffffffffffffffffff0000000000000000000000000000000000000000600a541617600a55612e3c6040519283928390929173ffffffffffffffffffffffffffffffffffffffff60209181604085019616845216910152565b0390a1005b60405190612e50602083610d0b565b60008252565b60405190612e6382610cd3565b60006080838281528260208201528260408201528260608201520152565b90604051612e8e81610cd3565b60806fffffffffffffffffffffffffffffffff6001839560ff8154848116875263ffffffff81871c16602088015260a01c1615156040860152015481808216166060850152821c16910152565b67ffffffffffffffff91612eed612e56565b50612ef6612e56565b50612f2a57166000526008602052604060002090610dc9612f1e6002612f23612f1e86612e81565b6143fc565b9401612e81565b1690816000526004602052612f45612f1e6040600020612e81565b916000526005602052610dc9612f1e6040600020612e81565b906105d5604051612f6e81610cb7565b60c0612fc682955463ffffffff8116845263ffffffff8160201c16602085015263ffffffff8160401c16604085015263ffffffff808260601c161660608501526122cc61ffff8260801c16608086019061ffff169052565b1515910152565b90612fdf60035461ffff9060a01c1690565b9061ffff8116801515938480956130cc575b611cdc5761301661301b9167ffffffffffffffff16600052600b602052604060002090565b612f5e565b9361302c61143a60c0870151151590565b6130b65761307b57505050604081015163ffffffff16815163ffffffff169263ffffffff6130706080613066602087015163ffffffff1690565b95015161ffff1690565b921693929190600190565b61ffff831611611ca5575050606081015163ffffffff16815163ffffffff169263ffffffff61307060a0613066602087015163ffffffff1690565b5050505050600090600090600090600090600090565b5061ffff841615612ff1565b60405190612e5082610c7a565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156103ea570180359067ffffffffffffffff82116103ea576020019181360383136103ea57565b6020818303126103ea5780359067ffffffffffffffff82116103ea57016040818303126103ea576040519161316a83610c9b565b813567ffffffffffffffff81116103ea5781613187918401610dae565b8352602082013567ffffffffffffffff81116103ea576131a79201610dae565b602082015290565b908160209103126103ea5751610dc9816105d7565b90916131db610dc9936040845260408401906103fa565b9160208184039101526103fa565b6040513d6000823e3d90fd5b35610dc9816105b8565b35610dc981610530565b613211614c9b565b60005b8281106132535750907fc97f93e817584952f1c1d633f93784b8430f0633d002f9dcc4de4fe2780424d09161324e6040519283928361351d565b0390a1565b6132666132618285856133d7565b613405565b8051158015613382575b6133095790613303826132fe61140a606060019651936132ef60208201516132e66132a2604085015163ffffffff1690565b6132de6132b26080870151151590565b916132c060a0880151151590565b946132c9610d4c565b9b8c5260208c015263ffffffff1660408b0152565b151588860152565b15156080870152565b015167ffffffffffffffff1690565b613478565b01613214565b604080517f19d7585700000000000000000000000000000000000000000000000000000000815282516004820152602083015160248201529082015163ffffffff166044820152606082015167ffffffffffffffff16606482015260808201511515608482015260a090910151151560a482015260c490fd5b5067ffffffffffffffff6133a1606083015167ffffffffffffffff1690565b1615613270565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b91908110156133e75760c0020190565b6133a8565b63ffffffff8116036103ea57565b35906105d5826133ec565b60c0813603126103ea5760a06040519161341e83610cef565b8035835260208101356020840152604081013561343a816133ec565b6040840152606081013561344d816105b8565b60608401526080810135613460816105d7565b60808401520135613470816105d7565b60a082015290565b6002608091835181556020840151600182015501916134cc63ffffffff604083015116849063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b6060810151835492909101517fffffffffffffffffffffffffffffffffffffffffffffffffffff0000ffffffff90921690151560201b64ff00000000161790151560281b65ff000000000016179055565b602080825281018390526040019160005b81811061353b5750505090565b90919260c080600192863581526020870135602082015263ffffffff6040880135613565816133ec565b16604082015267ffffffffffffffff6060880135613582816105b8565b1660608201526080870135613596816105d7565b1515608082015260a08701356135ab816105d7565b151560a082015201940192910161352e565b9067ffffffffffffffff610dc992166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b67ffffffffffffffff610dc991166000526007602052604060002054151590565b6020818303126103ea5780519067ffffffffffffffff82116103ea57019080601f830112156103ea57815161364f816111f5565b9261365d6040519485610d0b565b81845260208085019260051b8201019283116103ea57602001905b8282106136855750505090565b60208091835161369481610530565b815201910190613678565b601f8260209493601f19938186528686013760008582860101520116010190565b9796959267ffffffffffffffff61ffff9373ffffffffffffffffffffffffffffffffffffffff613711989794168b521660208a0152604089015216606087015260c0608087015260c086019161369f565b9260028210156137215760a00152565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b90919392956003549073ffffffffffffffffffffffffffffffffffffffff8216156137fc576000966137c99273ffffffffffffffffffffffffffffffffffffffff1695604051998a98899788977f89720a62000000000000000000000000000000000000000000000000000000008952600489016136c0565b03915afa908115610b17576000916137df575090565b610dc991503d806000833e6137f48183610d0b565b81019061361b565b50505050505050506020604051906138148183610d0b565b60008252601f19810190369083013790565b61382e614c9b565b60208101519160005b83518110156138e7578061386a61385060019387613a51565b5173ffffffffffffffffffffffffffffffffffffffff1690565b61389161388c73ffffffffffffffffffffffffffffffffffffffff8316612007565b615e2c565b61389d575b5001613837565b60405173ffffffffffffffffffffffffffffffffffffffff9190911681527fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758090602090a138613896565b5091505160005b81518110156139cb576139046138508284613a51565b9073ffffffffffffffffffffffffffffffffffffffff8216156139a1577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef6139988361397061396b61200760019773ffffffffffffffffffffffffffffffffffffffff1690565b615b07565b5060405173ffffffffffffffffffffffffffffffffffffffff90911681529081906020820190565b0390a1016138ee565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050565b604051906139dc82610c9b565b60606020838281520152565b906040516139f581610cd3565b608060ff600283958054855260018101546020860152015463ffffffff81166040850152818160201c161515606085015260281c161515910152565b916020610dc993818152019161369f565b908160209103126103ea573590565b80518210156133e75760209160051b010190565b90600182811c92168015613aae575b6020831014613a7f57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691613a74565b9060405191826000825492613acc84613a65565b8084529360018116908115613b385750600114613af1575b506105d592500383610d0b565b90506000929192526020600020906000915b818310613b1c5750509060206105d59282010138613ae4565b6020919350806001915483858901015201910190918492613b03565b602093506105d59592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138613ae4565b60409067ffffffffffffffff610dc99593168152816020820152019161369f565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b906000198201918211613bd757565b613b99565b91908203918211613bd757565b67ffffffffffffffff166000526008602052610dc96004604060002001613ab8565b91908110156133e75760081b0190565b35610dc9816105d7565b35610dc9816107c0565b35610dc9816133ec565b613e0a60c06105d593613c818135613c50816133ec565b859063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b6020810135613c8f816133ec565b67ffffffff0000000085549160201b16807fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff83161786557fffffffffffffffffffffffffffffffffffffffff0000000000000000ffffffff6bffffffff00000000000000006040850135613d02816133ec565b60401b16921617178455613d5e6060820135613d1d816133ec565b85547fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff1660609190911b6fffffffff00000000000000000000000016178555565b613db0613d6d60808301613c25565b85547fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff1660809190911b71ffff0000000000000000000000000000000016178555565b613e04613dbf60a08301613c25565b85547fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff1660909190911b73ffff00000000000000000000000000000000000016178555565b01613c1b565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016179055565b6105d59092919260c0612fc68160e084019663ffffffff8135613e72816133ec565b16855263ffffffff6020820135613e88816133ec565b16602086015263ffffffff6040820135613ea1816133ec565b166040860152613ec3613eb6606083016133fa565b63ffffffff166060870152565b613edd613ed2608083016107cc565b61ffff166080870152565b613ef7613eec60a083016107cc565b61ffff1660a0870152565b016105e1565b91908110156133e75760051b0190565b35906fffffffffffffffffffffffffffffffff821682036103ea57565b91908260609103126103ea576040516060810181811067ffffffffffffffff821117610c96576040526040613f7f8183958035613f66816105d7565b8552613f7460208201613f0d565b602086015201613f0d565b910152565b6fffffffffffffffffffffffffffffffff613fc4604080938035613fa7816105d7565b1515865283613fb860208301613f0d565b16602087015201613f0d565b16910152565b608090613feb6105d5949695939660e0830197151583526020830190613f84565b0190613f84565b9161400c918354906000199060031b92831b921b19161790565b9055565b81811061401b575050565b60008155600101614010565b81810292918115918404141715613bd757565b805490600081558161404a575050565b6000526020600020908101905b818110614062575050565b60008155600101614057565b60056105d591600081556000600182015560006002820155600060038201556004810161409b8154613a65565b90816140aa575b50500161403a565b81601f600093116001146140c25750555b38806140a2565b818352602083206140dd91601f01861c810190600101614010565b808252602082209081548360011b906000198560031b1c1916179055556140bb565b91908110156133e75760051b810135907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee1813603018212156103ea570190565b9080601f830112156103ea578135614156816111f5565b926141646040519485610d0b565b81845260208085019260051b820101918383116103ea5760208201905b83821061419057505050505090565b813567ffffffffffffffff81116103ea576020916141b387848094880101610dae565b815201910190614181565b610120813603126103ea57604051906141d682610cd3565b6141df816105ca565b8252602081013567ffffffffffffffff81116103ea57614202903690830161413f565b602083015260408101359067ffffffffffffffff82116103ea5761422c61424d9236908301610dae565b604084015261423e3660608301613f2a565b606084015260c0369101613f2a565b608082015290565b9190601f811161426457505050565b6105d5926000526020600020906020601f840160051c83019310614290575b601f0160051c0190614010565b9091508190614283565b919091825167ffffffffffffffff8111610c96576142c2816142bc8454613a65565b84614255565b6020601f82116001146142fe57819061400c9394956000926142f3575b50506000198260011b9260031b1c19161790565b0151905038806142df565b601f1982169061431384600052602060002090565b9160005b81811061434f57509583600195969710614336575b505050811b019055565b015160001960f88460031b161c1916905538808061432c565b9192602060018192868b015181550194019201614317565b6143cb6143966105d59597969467ffffffffffffffff60a09516845261010060208501526101008401906103fa565b9660408301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b01906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b614404612e56565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff82511690602083019163ffffffff8351164203428111613bd757614468906fffffffffffffffffffffffffffffffff60808701511690614027565b8101809111613bd75761448e6fffffffffffffffffffffffffffffffff92918392615b73565b161682524263ffffffff16905290565b6000608082016144b361143a610570836131ff565b6145ee575060208201916144d16020611ad9611ab1611aa4876131f5565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610b175783916145cf575b506145a75761452b611b92846131f5565b614534836131f5565b9061454d61143a60a0830193610e0f61096386866130e5565b6145675750506105d5929161456291506131f5565b6155bf565b61457192506130e5565b906117006040519283927f24eb47e500000000000000000000000000000000000000000000000000000000845260048401613a31565b6004827f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b6145e8915060203d602011610b1057610b028183610d0b565b3861451a565b6145fa61463b916131ff565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835273ffffffffffffffffffffffffffffffffffffffff16600452602490565b90fd5b916080830161465261143a610570836131ff565b611d6e575060208301926146706020611ad9611ab1611aa4886131f5565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610b1757600091614719575b50611d25576146cb611b92856131f5565b6146d4846131f5565b906146ed61143a60a0830193610e0f61096386866130e5565b61456757505061ffff161561470d576147086105d5926131f5565b615670565b6145626105d5926131f5565b614732915060203d602011610b1057610b028183610d0b565b386146ba565b6040519061474582610c9b565b6000825260208201600081526020820151917fffffffff000000000000000000000000000000000000000000000000000000006028602483015160e01c92015193167fb148ea5f0000000000000000000000000000000000000000000000000000000081036147b5575083525290565b7fa176027f0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b80516101188110614a3c575060048101517f000000000000000000000000000000000000000000000000000000000000000063ffffffff811663ffffffff831603614a03575050600881015190600c81015191608c82015191609081015193609482015160b88301519360f860d885015194015191614865895163ffffffff1690565b63ffffffff811663ffffffff8416036149c957507f000000000000000000000000000000000000000000000000000000000000000063ffffffff811663ffffffff86160361498f57506107d063ffffffff891603614955576107d063ffffffff82160361491c5750916148de9593916020979593614fd9565b9101518082036148ec575050565b7f7be225b60000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b7f0389caa2000000000000000000000000000000000000000000000000000000006000526107d060045263ffffffff1660245260446000fd5b7f22e102a0000000000000000000000000000000000000000000000000000000006000526107d060045263ffffffff881660245260446000fd5b7f77e480260000000000000000000000000000000000000000000000000000000060005263ffffffff908116600452841660245260446000fd5b7fe366a1170000000000000000000000000000000000000000000000000000000060005263ffffffff908116600452821660245260446000fd5b7f960693cd0000000000000000000000000000000000000000000000000000000060005263ffffffff9081166004521660245260446000fd5b7f758b22cc0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b908160209103126103ea575190565b80518015614b1d57602003614ae057614a9a6020825183010160208301614a69565b9060ff8211614aaa575060ff1690565b611700906040519182917f953576f700000000000000000000000000000000000000000000000000000000835260048301611dbc565b611700906040519182917f953576f700000000000000000000000000000000000000000000000000000000835260206004840181815201906103fa565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff8211613bd757565b60ff16604d8111613bd757600a0a90565b8015614b75576000190490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b8115614b75570490565b907f000000000000000000000000000000000000000000000000000000000000000060ff811660ff8316818114614c945711614c6957614bee8282614b43565b91604d60ff8416118015614c50575b614c1657505090614c10610dc992614b57565b90614027565b7fa9cb113d0000000000000000000000000000000000000000000000000000000060005260ff908116600452166024525060445260646000fd5b50614c62614c5d84614b57565b614b68565b8411614bfd565b614c738183614b43565b91604d60ff841611614c1657505090614c8e610dc992614b57565b90614ba4565b5050505090565b73ffffffffffffffffffffffffffffffffffffffff600154163303614cbc57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b60409067ffffffffffffffff610dc9949316815281602082015201906103fa565b90805115612ab5578051602082012067ffffffffffffffff831692836000526008602052614d3c826005604060002001615b3c565b15614d95575081614d847f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea93614d7f614d90946000526009602052604060002090565b61429a565b60405191829182611dbc565b0390a2565b90506117006040519283927f393b8ad200000000000000000000000000000000000000000000000000000000845260048401614ce6565b614e0b6127109167ffffffffffffffff6020820135614dea816105b8565b166000908152600b602052604090205460801c61ffff169060600135614027565b0490565b67ffffffffffffffff6020820135614e26816105b8565b166000908152600b602052604090209161ffff1615614e6657610dc991614c1061246d6060614e5e940135925461ffff9060901c1690565b612710900490565b9054610dc991614e5e9160801c61ffff169060600135614027565b60009160808201614e9761143a610570836131ff565b614f7357506020820191614eb56020611ad9611ab1611aa4876131f5565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610b17578591614f54575b50614f2c5791614f25611d209260606105d59695614f1e611b92866131f5565b0135613bdc565b92506131f5565b6004847f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b614f6d915060203d602011610b1057610b028183610d0b565b38614efe565b61463b6145fa85926131ff565b949290939163ffffffff90604051958260208801981688526040870152166060850152608084015260a083015260c0820152600060e08201526107d06101008201526101008152614fd361012082610d0b565b51902090565b959263ffffffff8095929693604051978260208a019a168a526040890152166060870152608086015260a085015260c0840152600060e0840152166101008201526101008152614fd361012082610d0b565b602081519101517fffffffff00000000000000000000000000000000000000000000000000000000604051927fb148ea5f00000000000000000000000000000000000000000000000000000000602085015260e01b166024830152602882015260288152610dc9604882610d0b565b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016602082015260208152610dc9604082610d0b565b73ffffffffffffffffffffffffffffffffffffffff600a54163314158061512b575b6150fd57565b7f8e4a23d6000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b5073ffffffffffffffffffffffffffffffffffffffff600154163314156150f7565b6105d59092919260608101936fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b8151919291156153c25760408301516fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff6151ef6151da60208701516fffffffffffffffffffffffffffffffff1690565b6fffffffffffffffffffffffffffffffff1690565b91161161538e5761534b6105d592935b61525261520c8251151590565b84547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178455565b61530b604061527460208401516fffffffffffffffffffffffffffffffff1690565b85547fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff8216178655926152f4600187019485906fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b01516fffffffffffffffffffffffffffffffff1690565b6fffffffffffffffffffffffffffffffff7fffffffffffffffffffffffffffffffff0000000000000000000000000000000083549260801b169116179055565b80547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016179055565b6040517f8020d12400000000000000000000000000000000000000000000000000000000815280611700856004830161514d565b6fffffffffffffffffffffffffffffffff6153f060408501516fffffffffffffffffffffffffffffffff1690565b161580159061543f575b61540b5761534b6105d592936151ff565b6040517fd68af9cc00000000000000000000000000000000000000000000000000000000815280611700856004830161514d565b506154606151da60208501516fffffffffffffffffffffffffffffffff1690565b15156153fa565b91908201809211613bd757565b9261547f9192614027565b8101809111613bd757610dc991615b73565b604051906006548083528260208101600660005260206000209260005b8181106154c35750506105d592500383610d0b565b84548352600194850194879450602090930192016154ae565b906040519182815491828252602082019060005260206000209260005b81811061550e5750506105d592500383610d0b565b84548352600194850194879450602090930192016154f9565b67ffffffffffffffff16615548816000526007602052604060002054151590565b15615592575033600052600d6020526040600020541561556457565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fa9902c7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b67ffffffffffffffff7f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c91169182600052600860205280615640600260406000200173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615b85565b6040805173ffffffffffffffffffffffffffffffffffffffff909216825260208201929092529081908101614d90565b67ffffffffffffffff7f0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce61591169182600052600560205280615640604060002073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615b85565b67ffffffffffffffff7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894491169182600052600860205280615640604060002073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615b85565b67ffffffffffffffff7f61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac2097932891169182600052600460205280615640604060002073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615b85565b80548210156133e75760005260206000200190600090565b8054801561582a57600019019061581982826157ea565b60001982549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6000818152600760205260409020549081156158fa57600019820190828211613bd757600654926000198401938411613bd75783836000956158b995036158bf575b5050506158a86006615802565b600790600052602052604060002090565b55600190565b6158a86158eb916158e16158d76158f19560066157ea565b90549060031b1c90565b92839160066157ea565b90613ff2565b5538808061589b565b5050600090565b600181019180600052826020526040600020549283151560001461599d576000198401848111613bd7578354936000198501948511613bd75760009585836158b9976159559503615964575b505050615802565b90600052602052604060002090565b6159846158eb9161597b6158d761599495886157ea565b928391876157ea565b8590600052602052604060002090565b5538808061594d565b50505050600090565b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000602080830191825273ffffffffffffffffffffffffffffffffffffffff9490941660248301526044808301959095529381529092600091615a0c606482610d0b565b519082855af1156131e9576000513d615a85575073ffffffffffffffffffffffffffffffffffffffff81163b155b615a415750565b73ffffffffffffffffffffffffffffffffffffffff907f5274afe7000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60011415615a3a565b80549068010000000000000000821015610c965781615ab591600161400c940181556157ea565b81939154906000199060031b92831b921b19161790565b600081815260076020526040902054615b0157615aea816006615a8e565b600654906000526007602052604060002055600190565b50600090565b6000818152600d6020526040902054615b0157615b2581600c615a8e565b600c5490600052600d602052604060002055600190565b60008281526001820160205260409020546158fa5780615b5e83600193615a8e565b80549260005201602052604060002055600190565b9080821015615b80575090565b905090565b909291815490615b9c61143a8360ff9060a01c1690565b8015615e24575b615e1d57615bc26fffffffffffffffffffffffffffffffff83166151da565b9160018401908154615c02615bfc6124a0615bef6151da856fffffffffffffffffffffffffffffffff1690565b9460801c63ffffffff1690565b42613bdc565b80615d89575b5050868110615d3d5750858310615c675750506151da6105d59394615c2c92613bdc565b6fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b91615c776151da87945460801c90565b928315615cee5761174693615ca1615c9284615ca694613bdc565b615c9b83613bc8565b90615467565b614ba4565b7fd0c8d23a0000000000000000000000000000000000000000000000000000000060005260045260245273ffffffffffffffffffffffffffffffffffffffff16604452606490565b7fd0c8d23a00000000000000000000000000000000000000000000000000000000600052600019600452602482905273ffffffffffffffffffffffffffffffffffffffff831660445260646000fd5b7f1a76572a00000000000000000000000000000000000000000000000000000000600052600452602486905273ffffffffffffffffffffffffffffffffffffffff821660445260646000fd5b828692939611615df357615da36151da615daa9460801c90565b9186615474565b84547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178555923880615c08565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050509050565b508415615ba3565b6000818152600d60205260409020549081156158fa57600019820190828211613bd757600c54926000198401938411613bd75783836158b99460009603615e8c575b505050615e7b600c615802565b600d90600052602052604060002090565b615e7b6158eb91615ea46158d7615eae95600c6157ea565b928391600c6157ea565b55388080615e6e56fea164736f6c634300081a000a2e902d38f15b233cbb63711add0fca4545334d3a169d60c0a616494d7eea9544",
}

var USDCTokenPoolCCTPV2ABI = USDCTokenPoolCCTPV2MetaData.ABI

var USDCTokenPoolCCTPV2Bin = USDCTokenPoolCCTPV2MetaData.Bin

func DeployUSDCTokenPoolCCTPV2(auth *bind.TransactOpts, backend bind.ContractBackend, tokenMessenger common.Address, cctpMessageTransmitterProxy common.Address, token common.Address, advancedPoolHooks common.Address, rmnProxy common.Address, router common.Address) (common.Address, *types.Transaction, *USDCTokenPoolCCTPV2, error) {
	parsed, err := USDCTokenPoolCCTPV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(USDCTokenPoolCCTPV2Bin), backend, tokenMessenger, cctpMessageTransmitterProxy, token, advancedPoolHooks, rmnProxy, router)
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) GetAdvancedPoolHooks(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "getAdvancedPoolHooks")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GetAdvancedPoolHooks() (common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetAdvancedPoolHooks(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) GetAdvancedPoolHooks() (common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetAdvancedPoolHooks(&_USDCTokenPoolCCTPV2.CallOpts)
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "getCurrentRateLimiterState", remoteChainSelector, customBlockConfirmation)

	outstruct := new(GetCurrentRateLimiterState)
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutboundRateLimiterState = *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)
	outstruct.InboundRateLimiterState = *abi.ConvertType(out[1], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return *outstruct, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GetCurrentRateLimiterState(remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	return _USDCTokenPoolCCTPV2.Contract.GetCurrentRateLimiterState(&_USDCTokenPoolCCTPV2.CallOpts, remoteChainSelector, customBlockConfirmation)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) GetCurrentRateLimiterState(remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	return _USDCTokenPoolCCTPV2.Contract.GetCurrentRateLimiterState(&_USDCTokenPoolCCTPV2.CallOpts, remoteChainSelector, customBlockConfirmation)
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
	outstruct.RateLimitAdmin = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "getFee", arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)

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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _USDCTokenPoolCCTPV2.Contract.GetFee(&_USDCTokenPoolCCTPV2.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _USDCTokenPoolCCTPV2.Contract.GetFee(&_USDCTokenPoolCCTPV2.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) GetMinBlockConfirmation(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "getMinBlockConfirmation")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GetMinBlockConfirmation() (uint16, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetMinBlockConfirmation(&_USDCTokenPoolCCTPV2.CallOpts)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) GetMinBlockConfirmation() (uint16, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetMinBlockConfirmation(&_USDCTokenPoolCCTPV2.CallOpts)
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) GetRequiredCCVs(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "getRequiredCCVs", localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetRequiredCCVs(&_USDCTokenPoolCCTPV2.CallOpts, localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetRequiredCCVs(&_USDCTokenPoolCCTPV2.CallOpts, localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Caller) GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _USDCTokenPoolCCTPV2.contract.Call(opts, &out, "getTokenTransferFeeConfig", arg0, destChainSelector, arg2, arg3)

	if err != nil {
		return *new(IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolV2TokenTransferFeeConfig)).(*IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetTokenTransferFeeConfig(&_USDCTokenPoolCCTPV2.CallOpts, arg0, destChainSelector, arg2, arg3)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2CallerSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _USDCTokenPoolCCTPV2.Contract.GetTokenTransferFeeConfig(&_USDCTokenPoolCCTPV2.CallOpts, arg0, destChainSelector, arg2, arg3)
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "applyAuthorizedCallerUpdates", authorizedCallerArgs)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.ApplyAuthorizedCallerUpdates(&_USDCTokenPoolCCTPV2.TransactOpts, authorizedCallerArgs)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.ApplyAuthorizedCallerUpdates(&_USDCTokenPoolCCTPV2.TransactOpts, authorizedCallerArgs)
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "applyTokenTransferFeeConfigUpdates", tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.ApplyTokenTransferFeeConfigUpdates(&_USDCTokenPoolCCTPV2.TransactOpts, tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.ApplyTokenTransferFeeConfigUpdates(&_USDCTokenPoolCCTPV2.TransactOpts, tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.LockOrBurn0(&_USDCTokenPoolCCTPV2.TransactOpts, lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.LockOrBurn0(&_USDCTokenPoolCCTPV2.TransactOpts, lockOrBurnIn, blockConfirmationRequested, tokenArgs)
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn, blockConfirmationRequested)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.ReleaseOrMint0(&_USDCTokenPoolCCTPV2.TransactOpts, releaseOrMintIn, blockConfirmationRequested)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.ReleaseOrMint0(&_USDCTokenPoolCCTPV2.TransactOpts, releaseOrMintIn, blockConfirmationRequested)
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) SetDomains(opts *bind.TransactOpts, domains []USDCTokenPoolDomainUpdate) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "setDomains", domains)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) SetDomains(domains []USDCTokenPoolDomainUpdate) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.SetDomains(&_USDCTokenPoolCCTPV2.TransactOpts, domains)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) SetDomains(domains []USDCTokenPoolDomainUpdate) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.SetDomains(&_USDCTokenPoolCCTPV2.TransactOpts, domains)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) SetDynamicConfig(opts *bind.TransactOpts, router common.Address, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "setDynamicConfig", router, rateLimitAdmin)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) SetDynamicConfig(router common.Address, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.SetDynamicConfig(&_USDCTokenPoolCCTPV2.TransactOpts, router, rateLimitAdmin)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) SetDynamicConfig(router common.Address, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.SetDynamicConfig(&_USDCTokenPoolCCTPV2.TransactOpts, router, rateLimitAdmin)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) SetMinBlockConfirmation(opts *bind.TransactOpts, minBlockConfirmation uint16) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "setMinBlockConfirmation", minBlockConfirmation)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) SetMinBlockConfirmation(minBlockConfirmation uint16) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.SetMinBlockConfirmation(&_USDCTokenPoolCCTPV2.TransactOpts, minBlockConfirmation)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) SetMinBlockConfirmation(minBlockConfirmation uint16) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.SetMinBlockConfirmation(&_USDCTokenPoolCCTPV2.TransactOpts, minBlockConfirmation)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) SetRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "setRateLimitConfig", rateLimitConfigArgs)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) SetRateLimitConfig(rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.SetRateLimitConfig(&_USDCTokenPoolCCTPV2.TransactOpts, rateLimitConfigArgs)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) SetRateLimitConfig(rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.SetRateLimitConfig(&_USDCTokenPoolCCTPV2.TransactOpts, rateLimitConfigArgs)
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) UpdateAdvancedPoolHooks(opts *bind.TransactOpts, newHook common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "updateAdvancedPoolHooks", newHook)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) UpdateAdvancedPoolHooks(newHook common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.UpdateAdvancedPoolHooks(&_USDCTokenPoolCCTPV2.TransactOpts, newHook)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) UpdateAdvancedPoolHooks(newHook common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.UpdateAdvancedPoolHooks(&_USDCTokenPoolCCTPV2.TransactOpts, newHook)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Transactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.contract.Transact(opts, "withdrawFeeTokens", feeTokens, recipient)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Session) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.WithdrawFeeTokens(&_USDCTokenPoolCCTPV2.TransactOpts, feeTokens, recipient)
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2TransactorSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolCCTPV2.Contract.WithdrawFeeTokens(&_USDCTokenPoolCCTPV2.TransactOpts, feeTokens, recipient)
}

type USDCTokenPoolCCTPV2AdvancedPoolHooksUpdatedIterator struct {
	Event *USDCTokenPoolCCTPV2AdvancedPoolHooksUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2AdvancedPoolHooksUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2AdvancedPoolHooksUpdated)
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
		it.Event = new(USDCTokenPoolCCTPV2AdvancedPoolHooksUpdated)
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

func (it *USDCTokenPoolCCTPV2AdvancedPoolHooksUpdatedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2AdvancedPoolHooksUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2AdvancedPoolHooksUpdated struct {
	OldHook common.Address
	NewHook common.Address
	Raw     types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterAdvancedPoolHooksUpdated(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2AdvancedPoolHooksUpdatedIterator, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "AdvancedPoolHooksUpdated")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2AdvancedPoolHooksUpdatedIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "AdvancedPoolHooksUpdated", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchAdvancedPoolHooksUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2AdvancedPoolHooksUpdated) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "AdvancedPoolHooksUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2AdvancedPoolHooksUpdated)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "AdvancedPoolHooksUpdated", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseAdvancedPoolHooksUpdated(log types.Log) (*USDCTokenPoolCCTPV2AdvancedPoolHooksUpdated, error) {
	event := new(USDCTokenPoolCCTPV2AdvancedPoolHooksUpdated)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "AdvancedPoolHooksUpdated", log); err != nil {
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

type USDCTokenPoolCCTPV2CustomBlockConfirmationInboundRateLimitConsumedIterator struct {
	Event *USDCTokenPoolCCTPV2CustomBlockConfirmationInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2CustomBlockConfirmationInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2CustomBlockConfirmationInboundRateLimitConsumed)
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
		it.Event = new(USDCTokenPoolCCTPV2CustomBlockConfirmationInboundRateLimitConsumed)
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

func (it *USDCTokenPoolCCTPV2CustomBlockConfirmationInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2CustomBlockConfirmationInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2CustomBlockConfirmationInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2CustomBlockConfirmationInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2CustomBlockConfirmationInboundRateLimitConsumedIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "CustomBlockConfirmationInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2CustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2CustomBlockConfirmationInboundRateLimitConsumed)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*USDCTokenPoolCCTPV2CustomBlockConfirmationInboundRateLimitConsumed, error) {
	event := new(USDCTokenPoolCCTPV2CustomBlockConfirmationInboundRateLimitConsumed)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolCCTPV2CustomBlockConfirmationOutboundRateLimitConsumedIterator struct {
	Event *USDCTokenPoolCCTPV2CustomBlockConfirmationOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2CustomBlockConfirmationOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2CustomBlockConfirmationOutboundRateLimitConsumed)
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
		it.Event = new(USDCTokenPoolCCTPV2CustomBlockConfirmationOutboundRateLimitConsumed)
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

func (it *USDCTokenPoolCCTPV2CustomBlockConfirmationOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2CustomBlockConfirmationOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2CustomBlockConfirmationOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2CustomBlockConfirmationOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2CustomBlockConfirmationOutboundRateLimitConsumedIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "CustomBlockConfirmationOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2CustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2CustomBlockConfirmationOutboundRateLimitConsumed)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*USDCTokenPoolCCTPV2CustomBlockConfirmationOutboundRateLimitConsumed, error) {
	event := new(USDCTokenPoolCCTPV2CustomBlockConfirmationOutboundRateLimitConsumed)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
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
	Router         common.Address
	RateLimitAdmin common.Address
	Raw            types.Log
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

type USDCTokenPoolCCTPV2FeeTokenWithdrawnIterator struct {
	Event *USDCTokenPoolCCTPV2FeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2FeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2FeeTokenWithdrawn)
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
		it.Event = new(USDCTokenPoolCCTPV2FeeTokenWithdrawn)
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

func (it *USDCTokenPoolCCTPV2FeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2FeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2FeeTokenWithdrawn struct {
	Receiver common.Address
	FeeToken common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*USDCTokenPoolCCTPV2FeeTokenWithdrawnIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2FeeTokenWithdrawnIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2FeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2FeeTokenWithdrawn)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseFeeTokenWithdrawn(log types.Log) (*USDCTokenPoolCCTPV2FeeTokenWithdrawn, error) {
	event := new(USDCTokenPoolCCTPV2FeeTokenWithdrawn)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

type USDCTokenPoolCCTPV2MinBlockConfirmationSetIterator struct {
	Event *USDCTokenPoolCCTPV2MinBlockConfirmationSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2MinBlockConfirmationSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2MinBlockConfirmationSet)
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
		it.Event = new(USDCTokenPoolCCTPV2MinBlockConfirmationSet)
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

func (it *USDCTokenPoolCCTPV2MinBlockConfirmationSetIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2MinBlockConfirmationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2MinBlockConfirmationSet struct {
	MinBlockConfirmation uint16
	Raw                  types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterMinBlockConfirmationSet(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2MinBlockConfirmationSetIterator, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "MinBlockConfirmationSet")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2MinBlockConfirmationSetIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "MinBlockConfirmationSet", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchMinBlockConfirmationSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2MinBlockConfirmationSet) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "MinBlockConfirmationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2MinBlockConfirmationSet)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "MinBlockConfirmationSet", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseMinBlockConfirmationSet(log types.Log) (*USDCTokenPoolCCTPV2MinBlockConfirmationSet, error) {
	event := new(USDCTokenPoolCCTPV2MinBlockConfirmationSet)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "MinBlockConfirmationSet", log); err != nil {
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

type USDCTokenPoolCCTPV2RateLimitConfiguredIterator struct {
	Event *USDCTokenPoolCCTPV2RateLimitConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolCCTPV2RateLimitConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolCCTPV2RateLimitConfigured)
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
		it.Event = new(USDCTokenPoolCCTPV2RateLimitConfigured)
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

func (it *USDCTokenPoolCCTPV2RateLimitConfiguredIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolCCTPV2RateLimitConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolCCTPV2RateLimitConfigured struct {
	RemoteChainSelector       uint64
	CustomBlockConfirmation   bool
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) FilterRateLimitConfigured(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2RateLimitConfiguredIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.FilterLogs(opts, "RateLimitConfigured", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolCCTPV2RateLimitConfiguredIterator{contract: _USDCTokenPoolCCTPV2.contract, event: "RateLimitConfigured", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) WatchRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2RateLimitConfigured, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolCCTPV2.contract.WatchLogs(opts, "RateLimitConfigured", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolCCTPV2RateLimitConfigured)
				if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "RateLimitConfigured", log); err != nil {
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

func (_USDCTokenPoolCCTPV2 *USDCTokenPoolCCTPV2Filterer) ParseRateLimitConfigured(log types.Log) (*USDCTokenPoolCCTPV2RateLimitConfigured, error) {
	event := new(USDCTokenPoolCCTPV2RateLimitConfigured)
	if err := _USDCTokenPoolCCTPV2.contract.UnpackLog(event, "RateLimitConfigured", log); err != nil {
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

type GetCurrentRateLimiterState struct {
	OutboundRateLimiterState RateLimiterTokenBucket
	InboundRateLimiterState  RateLimiterTokenBucket
}
type GetDynamicConfig struct {
	Router         common.Address
	RateLimitAdmin common.Address
}
type GetFee struct {
	FeeUSDCents       *big.Int
	DestGasOverhead   uint32
	DestBytesOverhead uint32
	TokenFeeBps       uint16
	IsEnabled         bool
}

func (USDCTokenPoolCCTPV2AdvancedPoolHooksUpdated) Topic() common.Hash {
	return common.HexToHash("0xbaff46844acf36d6ee996f489a1a288709c4542bd33cd557770afd267d581209")
}

func (USDCTokenPoolCCTPV2AuthorizedCallerAdded) Topic() common.Hash {
	return common.HexToHash("0xeb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef")
}

func (USDCTokenPoolCCTPV2AuthorizedCallerRemoved) Topic() common.Hash {
	return common.HexToHash("0xc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda77580")
}

func (USDCTokenPoolCCTPV2ChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (USDCTokenPoolCCTPV2ChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (USDCTokenPoolCCTPV2ConfigSet) Topic() common.Hash {
	return common.HexToHash("0x2e902d38f15b233cbb63711add0fca4545334d3a169d60c0a616494d7eea9544")
}

func (USDCTokenPoolCCTPV2CustomBlockConfirmationInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce615")
}

func (USDCTokenPoolCCTPV2CustomBlockConfirmationOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac20979328")
}

func (USDCTokenPoolCCTPV2DomainsSet) Topic() common.Hash {
	return common.HexToHash("0xc97f93e817584952f1c1d633f93784b8430f0633d002f9dcc4de4fe2780424d0")
}

func (USDCTokenPoolCCTPV2DynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x22a0dbb8195755fbfc99667a86ae684c568e9dfbb1eccf7f90084e6166447970")
}

func (USDCTokenPoolCCTPV2FeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (USDCTokenPoolCCTPV2InboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (USDCTokenPoolCCTPV2LockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
}

func (USDCTokenPoolCCTPV2MinBlockConfirmationSet) Topic() common.Hash {
	return common.HexToHash("0xa7f8dbba8cdb126ce4a0e7939ec58e0161b70d808b585dd651d68e59d27e11fb")
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

func (USDCTokenPoolCCTPV2RateLimitConfigured) Topic() common.Hash {
	return common.HexToHash("0x41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb")
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

	GetAdvancedPoolHooks(opts *bind.CallOpts) (common.Address, error)

	GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error)

	GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

		error)

	GetDomain(opts *bind.CallOpts, chainSelector uint64) (USDCTokenPoolDomain, error)

	GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

		error)

	GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

		error)

	GetMinBlockConfirmation(opts *bind.CallOpts) (uint16, error)

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRequiredCCVs(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetSupportedChains(opts *bind.CallOpts) ([]uint64, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	GetTokenDecimals(opts *bind.CallOpts) (uint8, error)

	GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error)

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

	ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error)

	ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error)

	ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetDomains(opts *bind.TransactOpts, domains []USDCTokenPoolDomainUpdate) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, router common.Address, rateLimitAdmin common.Address) (*types.Transaction, error)

	SetMinBlockConfirmation(opts *bind.TransactOpts, minBlockConfirmation uint16) (*types.Transaction, error)

	SetRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateAdvancedPoolHooks(opts *bind.TransactOpts, newHook common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error)

	FilterAdvancedPoolHooksUpdated(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2AdvancedPoolHooksUpdatedIterator, error)

	WatchAdvancedPoolHooksUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2AdvancedPoolHooksUpdated) (event.Subscription, error)

	ParseAdvancedPoolHooksUpdated(log types.Log) (*USDCTokenPoolCCTPV2AdvancedPoolHooksUpdated, error)

	FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2AuthorizedCallerAddedIterator, error)

	WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2AuthorizedCallerAdded) (event.Subscription, error)

	ParseAuthorizedCallerAdded(log types.Log) (*USDCTokenPoolCCTPV2AuthorizedCallerAdded, error)

	FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2AuthorizedCallerRemovedIterator, error)

	WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2AuthorizedCallerRemoved) (event.Subscription, error)

	ParseAuthorizedCallerRemoved(log types.Log) (*USDCTokenPoolCCTPV2AuthorizedCallerRemoved, error)

	FilterChainAdded(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2ChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2ChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*USDCTokenPoolCCTPV2ChainAdded, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2ChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2ChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*USDCTokenPoolCCTPV2ChainRemoved, error)

	FilterConfigSet(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2ConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2ConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*USDCTokenPoolCCTPV2ConfigSet, error)

	FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2CustomBlockConfirmationInboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2CustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*USDCTokenPoolCCTPV2CustomBlockConfirmationInboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2CustomBlockConfirmationOutboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2CustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*USDCTokenPoolCCTPV2CustomBlockConfirmationOutboundRateLimitConsumed, error)

	FilterDomainsSet(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2DomainsSetIterator, error)

	WatchDomainsSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2DomainsSet) (event.Subscription, error)

	ParseDomainsSet(log types.Log) (*USDCTokenPoolCCTPV2DomainsSet, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2DynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2DynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*USDCTokenPoolCCTPV2DynamicConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*USDCTokenPoolCCTPV2FeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2FeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*USDCTokenPoolCCTPV2FeeTokenWithdrawn, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2InboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2InboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*USDCTokenPoolCCTPV2InboundRateLimitConsumed, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2LockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2LockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*USDCTokenPoolCCTPV2LockedOrBurned, error)

	FilterMinBlockConfirmationSet(opts *bind.FilterOpts) (*USDCTokenPoolCCTPV2MinBlockConfirmationSetIterator, error)

	WatchMinBlockConfirmationSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2MinBlockConfirmationSet) (event.Subscription, error)

	ParseMinBlockConfirmationSet(log types.Log) (*USDCTokenPoolCCTPV2MinBlockConfirmationSet, error)

	FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2OutboundRateLimitConsumedIterator, error)

	WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2OutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseOutboundRateLimitConsumed(log types.Log) (*USDCTokenPoolCCTPV2OutboundRateLimitConsumed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*USDCTokenPoolCCTPV2OwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*USDCTokenPoolCCTPV2OwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*USDCTokenPoolCCTPV2OwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*USDCTokenPoolCCTPV2OwnershipTransferred, error)

	FilterRateLimitConfigured(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolCCTPV2RateLimitConfiguredIterator, error)

	WatchRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolCCTPV2RateLimitConfigured, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRateLimitConfigured(log types.Log) (*USDCTokenPoolCCTPV2RateLimitConfigured, error)

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
