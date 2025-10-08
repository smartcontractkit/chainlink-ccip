// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fast_transfer_token_pool

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

type ClientAny2EVMMessage struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	Sender              []byte
	Data                []byte
	DestTokenAmounts    []ClientEVMTokenAmount
}

type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

type FastTransferTokenPoolAbstractDestChainConfig struct {
	MaxFillAmountPerRequest  *big.Int
	FillerAllowlistEnabled   bool
	FastTransferFillerFeeBps uint16
	FastTransferPoolFeeBps   uint16
	SettlementOverheadGas    uint32
	DestinationPool          []byte
	CustomExtraArgs          []byte
}

type FastTransferTokenPoolAbstractDestChainConfigUpdateArgs struct {
	FillerAllowlistEnabled   bool
	FastTransferFillerFeeBps uint16
	FastTransferPoolFeeBps   uint16
	SettlementOverheadGas    uint32
	RemoteChainSelector      uint64
	ChainFamilySelector      [4]byte
	MaxFillAmountPerRequest  *big.Int
	DestinationPool          []byte
	CustomExtraArgs          []byte
}

type FastTransferTokenPoolAbstractFillInfo struct {
	State  uint8
	Filler common.Address
}

type IFastTransferPoolQuote struct {
	CcipSettlementFee *big.Int
	FastTransferFee   *big.Int
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

var BurnMintFastTransferTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIBurnMintERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipReceive\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.Any2EVMMessage\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipSendToken\",\"inputs\":[{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFastTransferFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"settlementFeeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"settlementId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"computeFillId\",\"inputs\":[{\"name\":\"settlementId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceAmountNetFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourceDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"fastFill\",\"inputs\":[{\"name\":\"fillId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"settlementId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceAmountNetFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourceDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAccumulatedPoolFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowedFillers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCcipSendTokenFee\",\"inputs\":[{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"settlementFeeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"quote\",\"type\":\"tuple\",\"internalType\":\"structIFastTransferPool.Quote\",\"components\":[{\"name\":\"ccipSettlementFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fastTransferFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFastTransferTokenPoolAbstract.DestChainConfig\",\"components\":[{\"name\":\"maxFillAmountPerRequest\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fillerAllowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"fastTransferFillerFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fastTransferPoolFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"settlementOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destinationPool\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"customExtraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFillInfo\",\"inputs\":[{\"name\":\"fillId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFastTransferTokenPoolAbstract.FillInfo\",\"components\":[{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"enumIFastTransferPool.FillState\"},{\"name\":\"filler\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAllowedFiller\",\"inputs\":[{\"name\":\"filler\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRouter\",\"inputs\":[{\"name\":\"newRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateDestChainConfig\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFastTransferTokenPoolAbstract.DestChainConfigUpdateArgs[]\",\"components\":[{\"name\":\"fillerAllowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"fastTransferFillerFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fastTransferPoolFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"settlementOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"maxFillAmountPerRequest\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationPool\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"customExtraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateFillerAllowList\",\"inputs\":[{\"name\":\"fillersToAdd\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"fillersToRemove\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawPoolFees\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigUpdated\",\"inputs\":[{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"fastTransferFillerFeeBps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"fastTransferPoolFeeBps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"maxFillAmountPerRequest\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"destinationPool\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"},{\"name\":\"settlementOverheadGas\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fillerAllowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestinationPoolUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"destinationPool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FastTransferFilled\",\"inputs\":[{\"name\":\"fillId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"settlementId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"filler\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"destAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FastTransferRequested\",\"inputs\":[{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"fillId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"settlementId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sourceAmountNetFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"sourceDecimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"fillerFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"poolFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"destinationPool\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FastTransferSettled\",\"inputs\":[{\"name\":\"fillId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"settlementId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"fillerReimbursementAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"poolFeeAccumulated\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"prevState\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIFastTransferPool.FillState\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FillerAllowListUpdated\",\"inputs\":[{\"name\":\"addFillers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"removeFillers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeeWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RouterUpdated\",\"inputs\":[{\"name\":\"oldRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyFilledOrSettled\",\"inputs\":[{\"name\":\"fillId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AlreadySettled\",\"inputs\":[{\"name\":\"fillId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"FillerNotAllowlisted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"filler\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InsufficientPoolFees\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEncodedAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidFillId\",\"inputs\":[{\"name\":\"fillId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRouter\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"QuoteFeeExceedsUserMaxLimit\",\"inputs\":[{\"name\":\"quoteFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFastTransferFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TransferAmountExceedsMaxFillAmount\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x610140806040523461042c57616616803803809161001d82856104ab565b833981019060c08183031261042c5780516001600160a01b038116929083900361042c5761004d602083016104ce565b60408301516001600160401b03811161042c5783019180601f8401121561042c578251926001600160401b038411610495578360051b90602082019461009660405196876104ab565b855260208086019282010192831161042c57602001905b82821061047d575050506100c3606084016104dc565b9360a06100d2608086016104dc565b940151946001600160401b038616860361042c57331561046c57600180546001600160a01b031916331790558115801561045b575b801561044a575b610439578160209160049360805260c0526040519283809263313ce56760e01b82525afa600091816103f8575b506103cd575b5060a052600480546001600160a01b0319166001600160a01b0384169081179091558151151560e08190529091906102aa575b5015610294576101005261012052604051615f8590816106918239608051818181611343015281816113ad01528181611582015281816120d80152818161261501528181612d7601528181612f60015281816135400152818161358d015281816139b1015281816145dc01528181614a6101528181614d040152818161511f01526156fd015260a0518181816115c60152818161326d015281816134f60152818161385b01528181613ab501528181613b0901528181614b8b0152614bf5015260c051818181610ab00152818161143b0152818161242b01528181612e05015281816131700152613752015260e051818181610a6b01528181612b790152615c71015261010051818181613bf00152613cfb01526101205181613ad70152f35b6335fdcccd60e21b600052600060045260246000fd5b602091604051916102bb84846104ab565b60008352600036813760e051156103bc5760005b8351811015610336576001906001600160a01b036102ed82876104f0565b5116866102f982610532565b610306575b5050016102cf565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a138866102fe565b5091509260005b82518110156103b1576001906001600160a01b0361035b82866104f0565b511680156103ab578561036d82610630565b61037b575b50505b0161033d565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a13885610372565b50610375565b509291505038610174565b6335f4a7b360e01b60005260046000fd5b60ff1660ff82168181036103e15750610141565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d602011610431575b81610414602093836104ab565b8101031261042c57610425906104ce565b903861013b565b600080fd5b3d9150610407565b630a64406560e11b60005260046000fd5b506001600160a01b0381161561010e565b506001600160a01b03851615610107565b639b15e16f60e01b60005260046000fd5b6020809161048a846104dc565b8152019101906100ad565b634e487b7160e01b600052604160045260246000fd5b601f909101601f19168101906001600160401b0382119082101761049557604052565b519060ff8216820361042c57565b51906001600160a01b038216820361042c57565b80518210156105045760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b80548210156105045760005260206000200190600090565b600081815260036020526040902054801561062957600019810181811161061357600254600019810191908211610613578181036105c2575b50505060025480156105ac576000190161058681600261051a565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b6105fb6105d36105e493600261051a565b90549060031b1c928392600261051a565b819391549060031b91821b91600019901b19161790565b9055600052600360205260406000205538808061056b565b634e487b7160e01b600052601160045260246000fd5b5050600090565b8060005260036020526040600020541560001461068a5760025468010000000000000000811015610495576106716105e4826001859401600255600261051a565b9055600254906000526003602052604060002055600190565b5060009056fe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a714613ed157508063055befd414613637578063181f5a77146135b157806321df0da71461356d578063240028e81461351a57806324f65ee7146134dc5780632b2c0eb4146134c15780632e7aa8c8146130315780633907753714612cec5780634c5ef0ed14612ca757806354c8a4f314612b4757806362ddd3c414612ac45780636609f59914612aa85780636d3d1a5814612a815780636def4ce71461293d57806378b410f2146129035780637909b5491461288857806379ba5097146127d75780637d54534e1461275757806385572ffb146121cf57806387f060d014611f285780638926f54f14611ee35780638a18dcbd146119fb5780638da5cb5b146119d4578063929ea5ba146118ca578063962d40201461178e5780639a4575b9146113725780639fe280f5146112df578063a42a7b8b146111ad578063a7cd63b71461113f578063abe1c1e8146110d0578063acfecf9114610fab578063af58d59f14610f61578063b0f479a114610f3a578063b794658014610f02578063c0d7865514610e5e578063c4bffe2b14610d4c578063c607c82f14610cdd578063c75eea9c14610c3d578063cf7401f314610ad4578063dc0bd97114610a90578063e0351e1314610a53578063e8a1da17146102c05763f2fde38b1461020357600080fd5b346102bb5760206003193601126102bb576001600160a01b036102246140bf565b61022c614da1565b1633811461029157807fffffffffffffffffffffffff000000000000000000000000000000000000000060005416176000556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346102bb576102ce3661429e565b9190926102d9614da1565b6000905b8282106108ae5750505060009063ffffffff42165b8183106102fb57005b6103068383866147f9565b92610120843603126102bb576040519361031f85614105565b6103288161407c565b8552602081013567ffffffffffffffff81116102bb5781019336601f860112156102bb578435610357816143a1565b956103656040519788614191565b81875260208088019260051b820101903682116102bb5760208101925b82841061087f575050505060208601948552604082013567ffffffffffffffff81116102bb576103b5903690840161424f565b90604087019182526103df6103cd366060860161447a565b936060890194855260c036910161447a565b94608088019586526103f18451615252565b6103fb8651615252565b825151156108555761041767ffffffffffffffff895116615917565b1561081c5767ffffffffffffffff8851166000526007602052604060002061053585516001600160801b03604082015116906104f96001600160801b036020830151169151151583608060405161046d81614105565b858152602081018a905260408101849052606081018690520152855474ff000000000000000000000000000000000000000091151560a01b919091167fffffffffffffffffffffff0000000000000000000000000000000000000000009091166001600160801b0384161773ffffffff00000000000000000000000000000000608089901b1617178555565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166001600160801b0391909116176001830155565b61063787516001600160801b03604082015116906105fb6001600160801b036020830151169151151583608060405161056d81614105565b858152602081018a9052604081018490526060810186905201526002860180547fffffffffffffffffffffff000000000000000000000000000000000000000000166001600160801b0385161773ffffffff0000000000000000000000000000000060808a901b161791151560a01b74ff000000000000000000000000000000000000000016919091179055565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166001600160801b0391909116176003830155565b6004845191019080519067ffffffffffffffff8211610806576106648261065e855461470f565b856149c6565b602090601f831160011461079f57610694929160009183610794575b50506000198260011b9260031b1c19161790565b90555b60005b875180518210156106cf57906106c96001926106c28367ffffffffffffffff8e51169261484f565b5190614ddf565b0161069a565b505097967f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c293929196509461078967ffffffffffffffff600197511692519351915161075e610732604051968796875261010060208801526101008701906141f3565b9360408601906001600160801b0360408092805115158552826020820151166020860152015116910152565b60a08401906001600160801b0360408092805115158552826020820151166020860152015116910152565b0390a10191926102f2565b015190508d80610680565b90601f1983169184600052816000209260005b8181106107ee57509084600195949392106107d5575b505050811b019055610697565b015160001960f88460031b161c191690558c80806107c8565b929360206001819287860151815501950193016107b2565b634e487b7160e01b600052604160045260246000fd5b67ffffffffffffffff8851167f1d5ad3c50000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b833567ffffffffffffffff81116102bb576020916108a3839283369187010161424f565b815201930192610382565b9092919367ffffffffffffffff6108ce6108c9868886614863565b6146bd565b16926108d984615d78565b15610a25578360005260076020526108f76005604060002001615770565b9260005b84518110156109335760019086600052600760205261092c6005604060002001610925838961484f565b5190615e0c565b50016108fb565b5093909491959250806000526007602052600560406000206000815560006001820155600060028201556000600382015560048101610972815461470f565b90816109e2575b50500180549060008155816109c1575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a10190916102dd565b6000526020600020908101905b8181101561098957600081556001016109ce565b81601f600093116001146109fa5750555b8880610979565b81835260208320610a1591601f01861c81019060010161499c565b80825281602081209155556109f3565b837f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346102bb5760006003193601126102bb5760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b346102bb5760006003193601126102bb5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102bb5760e06003193601126102bb57610aed614065565b60607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc3601126102bb57604051610b2381614175565b60243580151581036102bb5781526044356001600160801b03811681036102bb5760208201526064356001600160801b03811681036102bb57604082015260607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c3601126102bb5760405190610b9882614175565b60843580151581036102bb57825260a4356001600160801b03811681036102bb57602083015260c4356001600160801b03811681036102bb5760408301526001600160a01b036009541633141580610c28575b610bfa57610bf892614ff7565b005b7f8e4a23d6000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b506001600160a01b0360015416331415610beb565b346102bb5760206003193601126102bb5767ffffffffffffffff610c5f614065565b610c676148a4565b50166000526007602052610cd9610c89610c8460406000206148cf565b6151df565b6040519182918291909160806001600160801b038160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b0390f35b346102bb5760a06003193601126102bb5760243567ffffffffffffffff811681036102bb576064359060ff821682036102bb576084359067ffffffffffffffff82116102bb57602092610d37610d4493369060040161424f565b9160443590600435614945565b604051908152f35b346102bb5760006003193601126102bb576040516005548082528160208101600560005260206000209260005b818110610e45575050610d8e92500382614191565b805190610db3610d9d836143a1565b92610dab6040519485614191565b8084526143a1565b90601f1960208401920136833760005b8151811015610df5578067ffffffffffffffff610de26001938561484f565b5116610dee828761484f565b5201610dc3565b5050906040519182916020830190602084525180915260408301919060005b818110610e22575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101610e14565b8454835260019485019486945060209093019201610d79565b346102bb5760206003193601126102bb57610e776140bf565b610e7f614da1565b6001600160a01b03811690811561085557600480547fffffffffffffffffffffffff000000000000000000000000000000000000000081169093179055604080516001600160a01b0393841681529190921660208201527f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f168491819081015b0390a1005b346102bb5760206003193601126102bb57610cd9610f26610f21614065565b614923565b6040519182916020835260208301906141f3565b346102bb5760006003193601126102bb5760206001600160a01b0360045416604051908152f35b346102bb5760206003193601126102bb5767ffffffffffffffff610f83614065565b610f8b6148a4565b50166000526007602052610cd9610c89610c8460026040600020016148cf565b346102bb5767ffffffffffffffff610fc2366142f0565b929091610fcd614da1565b1690610fe6826000526006602052604060002054151590565b156110a257816000526007602052611017600560406000200161100a368685614218565b6020815191012090615e0c565b1561105b577f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d769192611056604051928392602084526020840191614883565b0390a2005b61109e906040519384937f74f23c7c0000000000000000000000000000000000000000000000000000000085526004850152604060248501526044840191614883565b0390fd5b507f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346102bb5760206003193601126102bb576110e961463f565b50600435600052600d6020526040806000206001600160a01b0382519161110f83614121565b5461111d60ff8216846147ed565b81602084019160081c1681526111368451809451614459565b51166020820152f35b346102bb5760006003193601126102bb576040516002548082526020820190600260005260206000209060005b81811061119757610cd98561118381870382614191565b604051918291602083526020830190614331565b825484526020909301926001928301920161116c565b346102bb5760206003193601126102bb5767ffffffffffffffff6111cf614065565b1660005260076020526111e86005604060002001615770565b805190601f196112106111fa846143a1565b936112086040519586614191565b8085526143a1565b0160005b8181106112ce57505060005b815181101561126857806112366001928461484f565b51600052600860205261124c6040600020614749565b611256828661484f565b52611261818561484f565b5001611220565b826040518091602082016020835281518091526040830190602060408260051b8601019301916000905b8282106112a157505050500390f35b919360206112be82603f19600195979984950301865288516141f3565b9601920192018594939192611292565b806060602080938701015201611214565b346102bb5760206003193601126102bb576112f86140bf565b611300614da1565b6113086145a0565b908161131057005b60206001600160a01b0382611367857f738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b859599957f000000000000000000000000000000000000000000000000000000000000000061518d565b6040519485521692a2005b346102bb576113803661436e565b6060602060405161139081614121565b8281520152608081016113a2816146a9565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691160361174f57506020810177ffffffffffffffff000000000000000000000000000000006113fb826146bd565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156116bd57600091611720575b506116f657611485611480604084016146a9565b615c6f565b67ffffffffffffffff611497826146bd565b166114af816000526006602052604060002054151590565b156116c95760206001600160a01b0360045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa9081156116bd5760009161166d575b506001600160a01b0316330361163f57610f218161162c93611538606061152e6115bc966146bd565b9201358092614a18565b61154181615115565b7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff611574846146bd565b604080516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152336020820152908101949094521691606090a26146bd565b610cd960405160ff7f0000000000000000000000000000000000000000000000000000000000000000166020820152602081526115fa604082614191565b6040519261160784614121565b83526020830190815260405193849360208552516040602086015260608501906141f3565b9051601f198483030160408501526141f3565b7f728fe07b000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b6020813d6020116116b5575b8161168660209383614191565b810103126116b15751906001600160a01b03821682036116ae57506001600160a01b03611505565b80fd5b5080fd5b3d9150611679565b6040513d6000823e3d90fd5b7fa9902c7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f53ad11d80000000000000000000000000000000000000000000000000000000060005260046000fd5b611742915060203d602011611748575b61173a8183614191565b810190614ce1565b8361146c565b503d611730565b6117606001600160a01b03916146a9565b7f961c9a4f000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b346102bb5760606003193601126102bb5760043567ffffffffffffffff81116102bb576117bf90369060040161426d565b9060243567ffffffffffffffff81116102bb576117e0903690600401614428565b9060443567ffffffffffffffff81116102bb57611801903690600401614428565b6001600160a01b0360095416331415806118b5575b610bfa578386148015906118ab575b6118815760005b86811061183557005b8061187b6118496108c96001948b8b614863565b611854838989614873565b61187561186d61186586898b614873565b92369061447a565b91369061447a565b91614ff7565b0161182c565b7f568efce20000000000000000000000000000000000000000000000000000000060005260046000fd5b5080861415611825565b506001600160a01b0360015416331415611816565b346102bb5760406003193601126102bb5760043567ffffffffffffffff81116102bb576118fb90369060040161440d565b60243567ffffffffffffffff81116102bb5761191b90369060040161440d565b90611924614da1565b60005b8151811015611956578061194f6001600160a01b036119486001948661484f565b51166158de565b5001611927565b5060005b825181101561198957806119826001600160a01b0361197b6001948761484f565b51166159cc565b500161195a565b7ffd35c599d42a981cbb1bbf7d3e6d9855a59f5c994ec6b427118ee0c260e241936119c683610efd86604051938493604085526040850190614331565b908382036020850152614331565b346102bb5760006003193601126102bb5760206001600160a01b0360015416604051908152f35b346102bb5760206003193601126102bb5760043567ffffffffffffffff81116102bb57611a2c90369060040161426d565b611a34614da1565b60005b818110611a4057005b611a4b8183856147f9565b60a081017f1e10bdc4000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000611a9a83614f9d565b1614611ea2575b60208201611aae81614fdb565b90604084019161ffff80611ac185614fdb565b1691160161ffff8111611e8c5761ffff61271091161015611e62576080840167ffffffffffffffff611af2826146bd565b16600052600a60205260406000209460e0810194611b19611b138784614658565b906157bb565b611b238683614658565b600289019167ffffffffffffffff821161080657611b458261065e855461470f565b600090601f8311600114611dfe57611b74929160009183611df35750506000198260011b9260031b1c19161790565b90555b611b8084614fdb565b926001880197885498611b9288614fdb565b60181b64ffff0000001695611ba686614fea565b151560c087013597888555606088019c611bbf8e614fca565b60281b68ffffffff0000000000169360081b62ffff0016907fffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000016177fffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffffff16179060ff16171790556101008401611c349085614658565b90916003019167ffffffffffffffff821161080657611c578261065e855461470f565b600090601f8311600114611d7e579180611c8c92611c93969594600092611d735750506000198260011b9260031b1c19161790565b90556146bd565b93611c9d90614fdb565b94611ca790614fdb565b95611cb29083614658565b9091611cbd90614f9d565b97611cc790614fca565b92611cd190614fea565b936040519761ffff899816885261ffff16602088015260408701526060860160e0905260e0860190611d0292614883565b957fffffffff0000000000000000000000000000000000000000000000000000000016608085015263ffffffff1660a0840152151560c083015267ffffffffffffffff1692037f6cfec31453105612e33aed8011f0e249b68d55e4efa65374322eb7ceeee76fbd91a2600101611a37565b013590503880610680565b838252602082209a9e9d9c9b9a91601f198416815b818110611ddb5750919e9f9b9c9d9e6001939185611c939897969410611dc1575b505050811b0190556146bd565b60001960f88560031b161c199101351690558f8080611db4565b91936020600181928787013581550195019201611d93565b013590508e80610680565b8382526020822091601f198416815b818110611e4a5750908460019594939210611e30575b505050811b019055611b77565b60001960f88560031b161c199101351690558d8080611e23565b83830135855560019094019360209283019201611e0d565b7f382c09820000000000000000000000000000000000000000000000000000000060005260046000fd5b634e487b7160e01b600052601160045260246000fd5b63ffffffff611eb360608401614fca565b1615611aa1577f382c09820000000000000000000000000000000000000000000000000000000060005260046000fd5b346102bb5760206003193601126102bb576020611f1e67ffffffffffffffff611f0a614065565b166000526006602052604060002054151590565b6040519015158152f35b346102bb5760c06003193601126102bb576004356024356044359067ffffffffffffffff8216918281036102bb576064359260843560ff811681036102bb5760a435926001600160a01b038416928385036102bb5780600052600a60205260ff60016040600020015416612183575b50611fbd9082876040519286602085015260208452611fb7604085614191565b88614945565b86036121555785600052600d60205260406000206001600160a01b0360405191611fe683614121565b54611ff460ff8216846147ed565b60081c1660208201525194600386101561213f57600095612113579061201991614bf2565b926040519561202787614121565b600187526020870196338852818752600d6020526040872090519760038910156120ff5787986120fc985060ff60ff198454169116178255517fffffffffffffffffffffff0000000000000000000000000000000000000000ff74ffffffffffffffffffffffffffffffffffffffff0083549260081b1691161790556040519285845260208401527fd6f70fb263bfe7d01ec6802b3c07b6bd32579760fe9fcb4e248a036debb8cdf160403394a4337f0000000000000000000000000000000000000000000000000000000000000000614aad565b80f35b602488634e487b7160e01b81526021600452fd5b602486887f9b91b78c000000000000000000000000000000000000000000000000000000008252600452fd5b634e487b7160e01b600052602160045260246000fd5b857fcb537aa40000000000000000000000000000000000000000000000000000000060005260045260246000fd5b61219a33600052600c602052604060002054151590565b611f97577f6c46a9b5000000000000000000000000000000000000000000000000000000006000526004523360245260446000fd5b346102bb576121dd3661436e565b6001600160a01b036004541633036127295760a0813603126102bb5760405161220581614105565b813581526122156020830161407c565b9060208101918252604083013567ffffffffffffffff81116102bb5761223e903690850161424f565b9160408201928352606084013567ffffffffffffffff81116102bb57612267903690860161424f565b936060830194855260808101359067ffffffffffffffff82116102bb570136601f820112156102bb57803561229b816143a1565b916122a96040519384614191565b81835260208084019260061b820101903682116102bb57602001915b8183106126f1575050506080830152519067ffffffffffffffff821690519251935191825183019460208601936020818803126102bb57602081015167ffffffffffffffff81116102bb57019560a090879003126102bb576040519161232a83614105565b6020870151835261233d60408801614f8e565b916020840192835261235160608901614f8e565b916040850192835260808901519860ff8a168a036102bb5760608601998a5260a081015167ffffffffffffffff81116102bb57602091010187601f820112156102bb57805161239f816141b4565b986123ad6040519a8b614191565b818a52602082840101116102bb576123cb916020808b0191016141d0565b6080850196875277ffffffffffffffff00000000000000000000000000000000604051917f2cbc26bb00000000000000000000000000000000000000000000000000000000835260801b1660048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156116bd576000916126d2575b506116f65761246b81866146d2565b1561269457509560ff6124926124ce936124c1989961ffff80885193511691511691615363565b6124bc6124a7879a939a518587511690614bf2565b996124b58587511684614bf2565b9751614593565b614593565b9151168551918589614945565b9384600052600d60205260406000209161252f82604051946124ef86614121565b54956124fe60ff8816876147ed565b6001600160a01b03602087019760081c16875288600052600d6020526040600020600260ff198254161790556156b1565b60009383516003811015612680576125c2575050600093516020818051810103126125be57602001516001600160a01b0381168091036125be579061257391614cf9565b5190600382101561213f576125bb6060927f33e17439bb4d31426d9168fc32af3a69cfce0467ba0d532fa804c27b5ff2189c9460405193845260208401526040830190614459565ba3005b8480fd5b939094508251600381101561266c5760010361264057506125eb846001600160a01b0392614593565b9251166126016125fb8585614a0b565b30614cf9565b828061260f575b5050612573565b612639917f000000000000000000000000000000000000000000000000000000000000000061518d565b8582612608565b80867fb196a44a0000000000000000000000000000000000000000000000000000000060249352600452fd5b602482634e487b7160e01b81526021600452fd5b602486634e487b7160e01b81526021600452fd5b61109e906040519182917f24eb47e50000000000000000000000000000000000000000000000000000000083526020600484015260248301906141f3565b6126eb915060203d6020116117485761173a8183614191565b8961245c565b6040833603126102bb576020604091825161270b81614121565b612714866140d5565b815282860135838201528152019201916122c5565b7fd7f73334000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b346102bb5760206003193601126102bb577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d0917460206001600160a01b0361279b6140bf565b6127a3614da1565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000006009541617600955604051908152a1005b346102bb5760006003193601126102bb576000546001600160a01b038116330361285e577fffffffffffffffffffffffff0000000000000000000000000000000000000000600154913382841617600155166000556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346102bb5760206003193601126102bb576128a1614065565b506128eb60206128f8604051916128b88184614191565b6000835260003681376040516128ce8282614191565b600081526000368137604051948594606086526060860190614331565b9184830390850152614331565b600060408301520390f35b346102bb5760206003193601126102bb576020611f1e6001600160a01b036129296140bf565b16600052600c602052604060002054151590565b346102bb5760206003193601126102bb5767ffffffffffffffff61295f614065565b606060c060405161296f81614159565b600081526000602082015260006040820152600083820152600060808201528260a0820152015216600052600a60205260606040600020610cd96129b1615725565b6119c66040516129c081614159565b84548152612a6d60018601549563ffffffff602084019760ff81161515895261ffff60408601818360081c168152818c880191818560181c1683528560808a019560281c168552612a266003612a1860028a01614749565b9860a08c01998a5201614749565b9860c08101998a526040519e8f9e8f9260408452516040840152511515910152511660808c0152511660a08a0152511660c08801525160e0808801526101208701906141f3565b9051603f19868303016101008701526141f3565b346102bb5760006003193601126102bb5760206001600160a01b0360095416604051908152f35b346102bb5760006003193601126102bb57610cd9611183615725565b346102bb57612ad2366142f0565b612add929192614da1565b67ffffffffffffffff8216612aff816000526006602052604060002054151590565b15612b1a5750610bf892612b14913691614218565b90614ddf565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346102bb57612b6f612b77612b5b3661429e565b9491612b68939193614da1565b36916143b9565b9236916143b9565b7f000000000000000000000000000000000000000000000000000000000000000015612c7d5760005b8251811015612c0657806001600160a01b03612bbe6001938661484f565b5116612bc981615ce4565b612bd5575b5001612ba0565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a184612bce565b5060005b8151811015610bf857806001600160a01b03612c286001938561484f565b51168015612c7757612c398161589f565b612c46575b505b01612c0a565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a183612c3e565b50612c40565b7f35f4a7b30000000000000000000000000000000000000000000000000000000060005260046000fd5b346102bb5760406003193601126102bb57612cc0614065565b60243567ffffffffffffffff81116102bb57602091612ce6611f1e92369060040161424f565b906146d2565b346102bb5760206003193601126102bb5760043567ffffffffffffffff81116102bb578060040161010060031983360301126102bb576000604051612d308161413d565b52612d5d612d53612d4e612d4760c4860185614658565b3691614218565b614b17565b6064840135614bf2565b9060848301612d6b816146a9565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691160361174f5750602483019077ffffffffffffffff00000000000000000000000000000000612dc5836146bd565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156116bd57600091613012575b506116f65767ffffffffffffffff612e4d836146bd565b16612e65816000526006602052604060002054151590565b156116c95760206001600160a01b0360045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa9081156116bd57600091612ff3575b501561163f57612ed0826146bd565b90612ee660a4860192612ce6612d478585614658565b15612fac5750507ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0608067ffffffffffffffff612f50612f4a6044602098612f3689612f318a6146bd565b6156b1565b01956108c988612f45896146a9565b614cf9565b946146a9565b936001600160a01b0360405195817f000000000000000000000000000000000000000000000000000000000000000016875233898801521660408601528560608601521692a280604051612fa38161413d565b52604051908152f35b612fb69250614658565b61109e6040519283927f24eb47e5000000000000000000000000000000000000000000000000000000008452602060048501526024840191614883565b61300c915060203d6020116117485761173a8183614191565b85612ec1565b61302b915060203d6020116117485761173a8183614191565b85612e36565b346102bb5760a06003193601126102bb5761304a614065565b6024359060443567ffffffffffffffff81116102bb5761306e903690600401614091565b91606435916001600160a01b0383168093036102bb5760843567ffffffffffffffff81116102bb576130a4903690600401614091565b50506130ae61463f565b50604051936130bc856140e9565b600085526020850192600084526040860192600084526060870160008152606060806040516130ea81614105565b828152826020820152826040820152600083820152015261310b84836157bb565b67ffffffffffffffff8316936040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff000000000000000000000000000000008560801b1660048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156116bd576000916134a2575b506116f6576131af33615c6f565b6131c6856000526006602052604060002054151590565b156134745784600052600a60205260406000209485548b116134425750986132d86132e69260ff61336c9b9c63ffffffff60018a015461322861ffff8260081c169c8d9661321d61ffff8560181c1680998c615363565b928382935252614a0b565b8d5260281c16806133d7575061ffff61329661324660038c01614749565b985b6040519761325589614105565b8852602088019c8d52604088019586526060880193857f00000000000000000000000000000000000000000000000000000000000000001685523691614218565b9360808701948552816040519c8d986020808b01525160408a01525116606088015251166080860152511660a08401525160a060c084015260e08301906141f3565b03601f198101865285614191565b6020958694604051906132f98783614191565b6000825261331560026040519761330f89614105565b01614749565b8652868601526040850152606084015260808301526001600160a01b0360045416906040518097819482937f20487ded000000000000000000000000000000000000000000000000000000008452600484016144c6565b03915afa9283156116bd576000936133a5575b5082604094525181845161339281614121565b8481520190815283519283525190820152f35b9392508184813d83116133d0575b6133bd8183614191565b810103126102bb5760409351929361337f565b503d6133b3565b61329661ffff91604051906133eb82614121565b81526020810160018152604051917f181dcf1000000000000000000000000000000000000000000000000000000000602084015251602483015251151560448201526044815261343c606482614191565b98613248565b8a907f58dd87c50000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b847fa9902c7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6134bb915060203d6020116117485761173a8183614191565b8b6131a1565b346102bb5760006003193601126102bb576020610d446145a0565b346102bb5760006003193601126102bb57602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102bb5760206003193601126102bb5760206135356140bf565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b346102bb5760006003193601126102bb5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102bb5760006003193601126102bb57610cd96040516135d3606082614191565b602781527f4275726e4d696e74466173745472616e73666572546f6b656e506f6f6c20312e60208201527f362e332d6465760000000000000000000000000000000000000000000000000060408201526040519182916020835260208301906141f3565b60c06003193601126102bb5761364b614065565b60643567ffffffffffffffff81116102bb5761366b903690600401614091565b60843592916001600160a01b03841684036102bb5760a43567ffffffffffffffff81116102bb576136a0903690600401614091565b5050604051906136af826140e9565b60008252600060208301526000604083015260006060830152606060806040516136d881614105565b82815282602082015282604082015260008382015201526136f983826157bb565b6040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff000000000000000000000000000000008560801b1660048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156116bd57600091613eb2575b506116f65761379133615c6f565b6137b267ffffffffffffffff85166000526006602052604060002054151590565b15613e7a5767ffffffffffffffff8416600052600a6020526040600020805460243511613e3c5760018101549561ffff8760081c169663ffffffff61ffff8260181c1691613818613806848c602435615363565b908160608b01528060408b0152614a0b565b602088015260281c1680613dd8575061383360038401614749565b905b6040519061384282614105565b602435825260208201998a5260408201908152606082017f000000000000000000000000000000000000000000000000000000000000000060ff1681523661388b908a89614218565b91608084019283526040519b8c9460208601602090525160408601525161ffff1660608501525161ffff1660808401525160ff1660a08301525160c0820160a0905260e082016138da916141f3565b03601f19810189526138ec9089614191565b6040516020986138fc8a83614191565b6000825261391260026040519661330f88614105565b85528985015260408401526001600160a01b038216606084015260808301526001600160a01b03600454168760405180927f20487ded0000000000000000000000000000000000000000000000000000000082528180613976888d600484016144c6565b03915afa9081156116bd57600091613dab575b50845261399860243587614a18565b60208401516044358111613d7857506139d560243530337f0000000000000000000000000000000000000000000000000000000000000000614aad565b6139e0602435615115565b6001600160a01b038116613b98575b506001600160a01b0360045416948660405180977f96f4e9f90000000000000000000000000000000000000000000000000000000082528180613a368787600484016144c6565b039134905af19485156116bd578796600096613b63575b5091613b588697927ffa7d3740fa7611df3f0d8d8c3aa1ed57c4fffaf2dcd0c47535f18a4774b44acd9467ffffffffffffffff613b48613a9360208b0151602435614593565b95606060408c01519b0151905190613afc8d613ab18d883691614218565b908a7f0000000000000000000000000000000000000000000000000000000000000000917f000000000000000000000000000000000000000000000000000000000000000090614945565b9b604051998a998a5260ff7f000000000000000000000000000000000000000000000000000000000000000016908a01526040890152606088015260c0608088015260c08701906141f3565b9285840360a08701521696614883565b0390a4604051908152f35b878193989297503d8311613b91575b613b7c8183614191565b810103126102bb575186959094613b58613a4d565b503d613b72565b613bae845130336001600160a01b038516614aad565b835180158015613cc4575b15613c40576040517f095ea7b300000000000000000000000000000000000000000000000000000000898201526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081166024830152604480830193909352918152613c3a92909190613c34606484614191565b1661557e565b866139ef565b608488604051907f08c379a00000000000000000000000000000000000000000000000000000000082526004820152603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527f20746f206e6f6e2d7a65726f20616c6c6f77616e6365000000000000000000006064820152fd5b506040517fdd62ed3e0000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001660248201528881806044810103816001600160a01b0387165afa9081156116bd57600091613d4b575b5015613bb9565b90508881813d8311613d71575b613d628183614191565b810103126102bb575189613d44565b503d613d58565b7f61acdb930000000000000000000000000000000000000000000000000000000060005260045260443560245260446000fd5b90508781813d8311613dd1575b613dc28183614191565b810103126102bb575188613989565b503d613db8565b60405190613de582614121565b81526020810160018152604051917f181dcf10000000000000000000000000000000000000000000000000000000006020840152516024830152511515604482015260448152613e36606482614191565b90613835565b67ffffffffffffffff857f58dd87c5000000000000000000000000000000000000000000000000000000006000521660045260243560245260446000fd5b67ffffffffffffffff847fa9902c7e000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b613ecb915060203d6020116117485761173a8183614191565b86613783565b346102bb5760206003193601126102bb57600435907fffffffff0000000000000000000000000000000000000000000000000000000082168092036102bb57817fde1861a20000000000000000000000000000000000000000000000000000000060209314908115613fda575b8115613f4c575b5015158152f35b7f85572ffb00000000000000000000000000000000000000000000000000000000811491508115613fb0575b8115613f86575b5083613f45565b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483613f7f565b7f7909b5490000000000000000000000000000000000000000000000000000000081149150613f78565b90507faff2afbf000000000000000000000000000000000000000000000000000000008114801561403c575b8015614013575b90613f3e565b507f01ffc9a700000000000000000000000000000000000000000000000000000000811461400d565b507f0e64dd29000000000000000000000000000000000000000000000000000000008114614006565b6004359067ffffffffffffffff821682036102bb57565b359067ffffffffffffffff821682036102bb57565b9181601f840112156102bb5782359167ffffffffffffffff83116102bb57602083818601950101116102bb57565b600435906001600160a01b03821682036102bb57565b35906001600160a01b03821682036102bb57565b6080810190811067ffffffffffffffff82111761080657604052565b60a0810190811067ffffffffffffffff82111761080657604052565b6040810190811067ffffffffffffffff82111761080657604052565b6020810190811067ffffffffffffffff82111761080657604052565b60e0810190811067ffffffffffffffff82111761080657604052565b6060810190811067ffffffffffffffff82111761080657604052565b90601f601f19910116810190811067ffffffffffffffff82111761080657604052565b67ffffffffffffffff811161080657601f01601f191660200190565b60005b8381106141e35750506000910152565b81810151838201526020016141d3565b90601f19601f602093614211815180928187528780880191016141d0565b0116010190565b929192614224826141b4565b916142326040519384614191565b8294818452818301116102bb578281602093846000960137010152565b9080601f830112156102bb5781602061426a93359101614218565b90565b9181601f840112156102bb5782359167ffffffffffffffff83116102bb576020808501948460051b0101116102bb57565b60406003198201126102bb5760043567ffffffffffffffff81116102bb57816142c99160040161426d565b929092916024359067ffffffffffffffff82116102bb576142ec9160040161426d565b9091565b9060406003198301126102bb5760043567ffffffffffffffff811681036102bb57916024359067ffffffffffffffff82116102bb576142ec91600401614091565b906020808351928381520192019060005b81811061434f5750505090565b82516001600160a01b0316845260209384019390920191600101614342565b60206003198201126102bb576004359067ffffffffffffffff82116102bb576003198260a0920301126102bb5760040190565b67ffffffffffffffff81116108065760051b60200190565b9291906143c5816143a1565b936143d36040519586614191565b602085838152019160051b81019283116102bb57905b8282106143f557505050565b60208091614402846140d5565b8152019101906143e9565b9080601f830112156102bb5781602061426a933591016143b9565b9181601f840112156102bb5782359167ffffffffffffffff83116102bb57602080850194606085020101116102bb57565b90600382101561213f5752565b35906001600160801b03821682036102bb57565b91908260609103126102bb5760405161449281614175565b809280359081151582036102bb5760406144c191819385526144b660208201614466565b602086015201614466565b910152565b9067ffffffffffffffff909392931681526040602082015261450d6144f7845160a0604085015260e08401906141f3565b6020850151603f198483030160608501526141f3565b90604084015191603f198282030160808301526020808451928381520193019060005b818110614568575050506080846001600160a01b03606061426a969701511660a084015201519060c0603f19828503019101526141f3565b825180516001600160a01b031686526020908101518187015260409095019490920191600101614530565b91908203918211611e8c57565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156116bd57600091614610575090565b90506020813d602011614637575b8161462b60209383614191565b810103126102bb575190565b3d915061461e565b6040519061464c82614121565b60006020838281520152565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156102bb570180359067ffffffffffffffff82116102bb576020019181360383136102bb57565b356001600160a01b03811681036102bb5790565b3567ffffffffffffffff811681036102bb5790565b9067ffffffffffffffff61426a92166000526007602052600560406000200190602081519101209060019160005201602052604060002054151590565b90600182811c9216801561473f575b602083101461472957565b634e487b7160e01b600052602260045260246000fd5b91607f169161471e565b906040519182600082549261475d8461470f565b80845293600181169081156147cb5750600114614784575b5061478292500383614191565b565b90506000929192526020600020906000915b8183106147af5750509060206147829282010138614775565b6020919350806001915483858901015201910190918492614796565b6020935061478295925060ff1991501682840152151560051b82010138614775565b600382101561213f5752565b91908110156148395760051b810135907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee1813603018212156102bb570190565b634e487b7160e01b600052603260045260246000fd5b80518210156148395760209160051b010190565b91908110156148395760051b0190565b9190811015614839576060020190565b601f8260209493601f19938186528686013760008582860101520116010190565b604051906148b182614105565b60006080838281528260208201528260408201528260608201520152565b906040516148dc81614105565b60806001829460ff81546001600160801b038116865263ffffffff81861c16602087015260a01c161515604085015201546001600160801b0381166060840152811c910152565b67ffffffffffffffff16600052600760205261426a6004604060002001614749565b936149889060ff67ffffffffffffffff9561499695604051978896602088019a8b52166040870152606086015216608084015260a08084015260c08301906141f3565b03601f198101835282614191565b51902090565b8181106149a7575050565b6000815560010161499c565b81810292918115918404141715611e8c57565b9190601f81116149d557505050565b614782926000526020600020906020601f840160051c83019310614a01575b601f0160051c019061499c565b90915081906149f4565b91908201809211611e8c57565b67ffffffffffffffff7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894491169182600052600760205280614a8960406000206001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001692839161538a565b604080516001600160a01b039092168252602082019290925290819081015b0390a2565b6040517f23b872dd0000000000000000000000000000000000000000000000000000000060208201526001600160a01b039283166024820152929091166044830152606482019290925261478291614b1282608481015b03601f198101845283614191565b61557e565b80518015614b8757602003614b495780516020828101918301839003126102bb57519060ff8211614b49575060ff1690565b61109e906040519182917f953576f70000000000000000000000000000000000000000000000000000000083526020600484015260248301906141f3565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff8211611e8c57565b60ff16604d8111611e8c57600a0a90565b8115614bdc570490565b634e487b7160e01b600052601260045260246000fd5b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff811692828414614cda57828411614cb05790614c3791614bad565b91604d60ff8416118015614c95575b614c5f57505090614c5961426a92614bc1565b906149b3565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b50614c9f83614bc1565b8015614bdc57600019048411614c46565b614cb991614bad565b91604d60ff841611614c5f57505090614cd461426a92614bc1565b90614bd2565b5050505090565b908160209103126102bb575180151581036102bb5790565b906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016803b156102bb576040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b0393909316600484015260248301919091526000919082908290604490829084905af18015614d9657614d89575050565b81614d9391614191565b50565b6040513d84823e3d90fd5b6001600160a01b03600154163303614db557565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b908051156108555767ffffffffffffffff81516020830120921691826000526007602052614e14816005604060002001615950565b15614f4a5760005260086020526040600020815167ffffffffffffffff811161080657614e4b81614e45845461470f565b846149c6565b6020601f8211600114614ec05791614e9f827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9593614aa895600091614eb5575b506000198260011b9260031b1c19161790565b90556040519182916020835260208301906141f3565b905084015138614e8c565b601f1982169083600052806000209160005b818110614f32575092614aa89492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610614f19575b5050811b019055610f26565b85015160001960f88460031b161c191690553880614f0d565b9192602060018192868a015181550194019201614ed2565b509061109e6040519283927f393b8ad200000000000000000000000000000000000000000000000000000000845260048401526040602484015260448301906141f3565b519061ffff821682036102bb57565b357fffffffff00000000000000000000000000000000000000000000000000000000811681036102bb5790565b3563ffffffff811681036102bb5790565b3561ffff811681036102bb5790565b3580151581036102bb5790565b67ffffffffffffffff1660008181526006602052604090205490929190156150e757916150e460e0926150b98561504e7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b97615252565b846000526007602052615065816040600020615a60565b61506e83615252565b846000526007602052615088836002604060002001615a60565b60405194855260208501906001600160801b0360408092805115158552826020820151166020860152015116910152565b60808301906001600160801b0360408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690813b156102bb57604051907f42966c680000000000000000000000000000000000000000000000000000000082528160248160008096819560048401525af18015614d9657614d89575050565b6040517fa9059cbb0000000000000000000000000000000000000000000000000000000060208201526001600160a01b039092166024830152604482019290925261478291614b128260648101614b04565b6151e76148a4565b506001600160801b036060820151166001600160801b038083511691615232602085019361522c61521f63ffffffff87511642614593565b85608089015116906149b3565b90614a0b565b8082101561524b57505b16825263ffffffff4216905290565b905061523c565b8051156152d7576001600160801b036040820151166001600160801b036020830151161061527d5750565b6064906152d5604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906001600160801b0360408092805115158552826020820151166020860152015116910152565bfd5b6001600160801b036040820151161580159061534d575b6152f55750565b6064906152d5604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906001600160801b0360408092805115158552826020820151166020860152015116910152565b506001600160801b0360208201511615156152ee565b6153869061ffff61271061537d82829698979816846149b3565b049516906149b3565b0490565b9182549060ff8260a01c16158015615576575b615570576001600160801b03821691600185019081546153d063ffffffff6001600160801b0383169360801c1642614593565b90816154d2575b505084811061549357508383106154285750506153fd6001600160801b03928392614593565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b5460801c916154378185614593565b92600019810190808211611e8c5761545a61545f926001600160a01b0396614a0b565b614bd2565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b82856001600160a01b03927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611615546576154ed9261522c9160801c906149b3565b808410156155415750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff00000000000000000000000000000000161786559238806153d7565b6154f8565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b50821561539d565b6001600160a01b0361560091169160409260008085519361559f8786614191565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602086015260208151910182855af13d156156a9573d916155e4836141b4565b926155f187519485614191565b83523d6000602085013e615eac565b8051908161560d57505050565b60208061561e938301019101614ce1565b156156265750565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606091615eac565b67ffffffffffffffff7f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c91169182600052600760205280614a8960026040600020016001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001692839161538a565b60405190600b548083528260208101600b60005260206000209260005b81811061575757505061478292500383614191565b8454835260019485019487945060209093019201615742565b906040519182815491828252602082019060005260206000209260005b8181106157a257505061478292500383614191565b845483526001948501948794506020909301920161578d565b908015801561583b575b6157e45760008235615833575b60208211615821575b156157e4575050565b61109e6040519283927fec4c23ce000000000000000000000000000000000000000000000000000000008452602060048501526024840191614883565b6020830135156157db575060016157db565b5060016157d2565b50604081116157c5565b80548210156148395760005260206000200190600090565b80549068010000000000000000821015610806578161588491600161589b94018155615845565b81939154906000199060031b92831b921b19161790565b9055565b806000526003602052604060002054156000146158d8576158c181600261585d565b600254906000526003602052604060002055600190565b50600090565b80600052600c602052604060002054156000146158d85761590081600b61585d565b600b5490600052600c602052604060002055600190565b806000526006602052604060002054156000146158d85761593981600561585d565b600554906000526006602052604060002055600190565b600082815260018201602052604090205461598757806159728360019361585d565b80549260005201602052604060002055600190565b5050600090565b805480156159b65760001901906159a58282615845565b60001982549160031b1b1916905555565b634e487b7160e01b600052603160045260246000fd5b6000818152600c60205260409020548015615987576000198101818111611e8c57600b54906000198201918211611e8c57808203615a26575b505050615a12600b61598e565b600052600c60205260006040812055600190565b615a48615a3761588493600b615845565b90549060031b1c928392600b615845565b9055600052600c602052604060002055388080615a05565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1991615b876060928054615a9d63ffffffff8260801c1642614593565b9081615bbd575b50506001600160801b036001816020860151169282815416808510600014615bb557508280855b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416178155615b448651151582907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b60408601517fffffffffffffffffffffffffffffffff0000000000000000000000000000000060809190911b16939092166001600160801b031692909217910155565b6150e460405180926001600160801b0360408092805115158552826020820151166020860152015116910152565b838091615acb565b6001600160801b0391615be9839283615be26001880154948286169560801c906149b3565b9116614a0b565b80821015615c6857505b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff9290911692909216167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116174260801b73ffffffff00000000000000000000000000000000161781553880615aa4565b9050615bf3565b7f0000000000000000000000000000000000000000000000000000000000000000615c975750565b6001600160a01b031680600052600360205260406000205415615cb75750565b7fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6000818152600360205260409020548015615987576000198101818111611e8c57600254906000198201918211611e8c57818103615d3e575b505050615d2a600261598e565b600052600360205260006040812055600190565b615d60615d4f615884936002615845565b90549060031b1c9283926002615845565b90556000526003602052604060002055388080615d1d565b6000818152600660205260409020548015615987576000198101818111611e8c57600554906000198201918211611e8c57818103615dd2575b505050615dbe600561598e565b600052600660205260006040812055600190565b615df4615de3615884936005615845565b90549060031b1c9283926005615845565b90556000526006602052604060002055388080615db1565b906001820191816000528260205260406000205490811515600014615ea357600019820191808311611e8c5781546000198101908111611e8c578381615e5a9503615e6c575b50505061598e565b60005260205260006040812055600190565b615e8c615e7c6158849386615845565b90549060031b1c92839286615845565b905560005284602052604060002055388080615e52565b50505050600090565b91929015615f275750815115615ec0575090565b3b15615ec95790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015615f3a5750805190602001fd5b61109e906040519182917f08c379a00000000000000000000000000000000000000000000000000000000083526020600484015260248301906141f356fea164736f6c634300081a000a",
}

var BurnMintFastTransferTokenPoolABI = BurnMintFastTransferTokenPoolMetaData.ABI

var BurnMintFastTransferTokenPoolBin = BurnMintFastTransferTokenPoolMetaData.Bin

func DeployBurnMintFastTransferTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, localTokenDecimals uint8, allowlist []common.Address, rmnProxy common.Address, router common.Address, sourceChainSelector uint64) (common.Address, *types.Transaction, *BurnMintFastTransferTokenPool, error) {
	parsed, err := BurnMintFastTransferTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BurnMintFastTransferTokenPoolBin), backend, token, localTokenDecimals, allowlist, rmnProxy, router, sourceChainSelector)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnMintFastTransferTokenPool{address: address, abi: *parsed, BurnMintFastTransferTokenPoolCaller: BurnMintFastTransferTokenPoolCaller{contract: contract}, BurnMintFastTransferTokenPoolTransactor: BurnMintFastTransferTokenPoolTransactor{contract: contract}, BurnMintFastTransferTokenPoolFilterer: BurnMintFastTransferTokenPoolFilterer{contract: contract}}, nil
}

type BurnMintFastTransferTokenPool struct {
	address common.Address
	abi     abi.ABI
	BurnMintFastTransferTokenPoolCaller
	BurnMintFastTransferTokenPoolTransactor
	BurnMintFastTransferTokenPoolFilterer
}

type BurnMintFastTransferTokenPoolCaller struct {
	contract *bind.BoundContract
}

type BurnMintFastTransferTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type BurnMintFastTransferTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type BurnMintFastTransferTokenPoolSession struct {
	Contract     *BurnMintFastTransferTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type BurnMintFastTransferTokenPoolCallerSession struct {
	Contract *BurnMintFastTransferTokenPoolCaller
	CallOpts bind.CallOpts
}

type BurnMintFastTransferTokenPoolTransactorSession struct {
	Contract     *BurnMintFastTransferTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type BurnMintFastTransferTokenPoolRaw struct {
	Contract *BurnMintFastTransferTokenPool
}

type BurnMintFastTransferTokenPoolCallerRaw struct {
	Contract *BurnMintFastTransferTokenPoolCaller
}

type BurnMintFastTransferTokenPoolTransactorRaw struct {
	Contract *BurnMintFastTransferTokenPoolTransactor
}

func NewBurnMintFastTransferTokenPool(address common.Address, backend bind.ContractBackend) (*BurnMintFastTransferTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(BurnMintFastTransferTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindBurnMintFastTransferTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPool{address: address, abi: abi, BurnMintFastTransferTokenPoolCaller: BurnMintFastTransferTokenPoolCaller{contract: contract}, BurnMintFastTransferTokenPoolTransactor: BurnMintFastTransferTokenPoolTransactor{contract: contract}, BurnMintFastTransferTokenPoolFilterer: BurnMintFastTransferTokenPoolFilterer{contract: contract}}, nil
}

func NewBurnMintFastTransferTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*BurnMintFastTransferTokenPoolCaller, error) {
	contract, err := bindBurnMintFastTransferTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolCaller{contract: contract}, nil
}

func NewBurnMintFastTransferTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnMintFastTransferTokenPoolTransactor, error) {
	contract, err := bindBurnMintFastTransferTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolTransactor{contract: contract}, nil
}

func NewBurnMintFastTransferTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnMintFastTransferTokenPoolFilterer, error) {
	contract, err := bindBurnMintFastTransferTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolFilterer{contract: contract}, nil
}

func bindBurnMintFastTransferTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BurnMintFastTransferTokenPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintFastTransferTokenPool.Contract.BurnMintFastTransferTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.BurnMintFastTransferTokenPoolTransactor.contract.Transfer(opts)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.BurnMintFastTransferTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintFastTransferTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.contract.Transfer(opts)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) ComputeFillId(opts *bind.CallOpts, settlementId [32]byte, sourceChainSelector uint64, sourceAmountNetFee *big.Int, sourceDecimals uint8, receiver []byte) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "computeFillId", settlementId, sourceChainSelector, sourceAmountNetFee, sourceDecimals, receiver)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) ComputeFillId(settlementId [32]byte, sourceChainSelector uint64, sourceAmountNetFee *big.Int, sourceDecimals uint8, receiver []byte) ([32]byte, error) {
	return _BurnMintFastTransferTokenPool.Contract.ComputeFillId(&_BurnMintFastTransferTokenPool.CallOpts, settlementId, sourceChainSelector, sourceAmountNetFee, sourceDecimals, receiver)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) ComputeFillId(settlementId [32]byte, sourceChainSelector uint64, sourceAmountNetFee *big.Int, sourceDecimals uint8, receiver []byte) ([32]byte, error) {
	return _BurnMintFastTransferTokenPool.Contract.ComputeFillId(&_BurnMintFastTransferTokenPool.CallOpts, settlementId, sourceChainSelector, sourceAmountNetFee, sourceDecimals, receiver)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) GetAccumulatedPoolFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "getAccumulatedPoolFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) GetAccumulatedPoolFees() (*big.Int, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetAccumulatedPoolFees(&_BurnMintFastTransferTokenPool.CallOpts)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) GetAccumulatedPoolFees() (*big.Int, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetAccumulatedPoolFees(&_BurnMintFastTransferTokenPool.CallOpts)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) GetAllowList() ([]common.Address, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetAllowList(&_BurnMintFastTransferTokenPool.CallOpts)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) GetAllowList() ([]common.Address, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetAllowList(&_BurnMintFastTransferTokenPool.CallOpts)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) GetAllowListEnabled() (bool, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetAllowListEnabled(&_BurnMintFastTransferTokenPool.CallOpts)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) GetAllowListEnabled() (bool, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetAllowListEnabled(&_BurnMintFastTransferTokenPool.CallOpts)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) GetAllowedFillers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "getAllowedFillers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) GetAllowedFillers() ([]common.Address, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetAllowedFillers(&_BurnMintFastTransferTokenPool.CallOpts)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) GetAllowedFillers() ([]common.Address, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetAllowedFillers(&_BurnMintFastTransferTokenPool.CallOpts)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) GetCCVs(opts *bind.CallOpts, arg0 uint64) (GetCCVs,

	error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "getCCVs", arg0)

	outstruct := new(GetCCVs)
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequiredCCVs = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.OptionalCCVs = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.OptionalThreshold = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) GetCCVs(arg0 uint64) (GetCCVs,

	error) {
	return _BurnMintFastTransferTokenPool.Contract.GetCCVs(&_BurnMintFastTransferTokenPool.CallOpts, arg0)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) GetCCVs(arg0 uint64) (GetCCVs,

	error) {
	return _BurnMintFastTransferTokenPool.Contract.GetCCVs(&_BurnMintFastTransferTokenPool.CallOpts, arg0)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) GetCcipSendTokenFee(opts *bind.CallOpts, destinationChainSelector uint64, amount *big.Int, receiver []byte, settlementFeeToken common.Address, extraArgs []byte) (IFastTransferPoolQuote, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "getCcipSendTokenFee", destinationChainSelector, amount, receiver, settlementFeeToken, extraArgs)

	if err != nil {
		return *new(IFastTransferPoolQuote), err
	}

	out0 := *abi.ConvertType(out[0], new(IFastTransferPoolQuote)).(*IFastTransferPoolQuote)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) GetCcipSendTokenFee(destinationChainSelector uint64, amount *big.Int, receiver []byte, settlementFeeToken common.Address, extraArgs []byte) (IFastTransferPoolQuote, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetCcipSendTokenFee(&_BurnMintFastTransferTokenPool.CallOpts, destinationChainSelector, amount, receiver, settlementFeeToken, extraArgs)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) GetCcipSendTokenFee(destinationChainSelector uint64, amount *big.Int, receiver []byte, settlementFeeToken common.Address, extraArgs []byte) (IFastTransferPoolQuote, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetCcipSendTokenFee(&_BurnMintFastTransferTokenPool.CallOpts, destinationChainSelector, amount, receiver, settlementFeeToken, extraArgs)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "getCurrentInboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetCurrentInboundRateLimiterState(&_BurnMintFastTransferTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetCurrentInboundRateLimiterState(&_BurnMintFastTransferTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "getCurrentOutboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_BurnMintFastTransferTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_BurnMintFastTransferTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) GetDestChainConfig(opts *bind.CallOpts, remoteChainSelector uint64) (FastTransferTokenPoolAbstractDestChainConfig, []common.Address, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "getDestChainConfig", remoteChainSelector)

	if err != nil {
		return *new(FastTransferTokenPoolAbstractDestChainConfig), *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(FastTransferTokenPoolAbstractDestChainConfig)).(*FastTransferTokenPoolAbstractDestChainConfig)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return out0, out1, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) GetDestChainConfig(remoteChainSelector uint64) (FastTransferTokenPoolAbstractDestChainConfig, []common.Address, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetDestChainConfig(&_BurnMintFastTransferTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) GetDestChainConfig(remoteChainSelector uint64) (FastTransferTokenPoolAbstractDestChainConfig, []common.Address, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetDestChainConfig(&_BurnMintFastTransferTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) GetFillInfo(opts *bind.CallOpts, fillId [32]byte) (FastTransferTokenPoolAbstractFillInfo, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "getFillInfo", fillId)

	if err != nil {
		return *new(FastTransferTokenPoolAbstractFillInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(FastTransferTokenPoolAbstractFillInfo)).(*FastTransferTokenPoolAbstractFillInfo)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) GetFillInfo(fillId [32]byte) (FastTransferTokenPoolAbstractFillInfo, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetFillInfo(&_BurnMintFastTransferTokenPool.CallOpts, fillId)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) GetFillInfo(fillId [32]byte) (FastTransferTokenPoolAbstractFillInfo, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetFillInfo(&_BurnMintFastTransferTokenPool.CallOpts, fillId)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "getRateLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) GetRateLimitAdmin() (common.Address, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetRateLimitAdmin(&_BurnMintFastTransferTokenPool.CallOpts)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) GetRateLimitAdmin() (common.Address, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetRateLimitAdmin(&_BurnMintFastTransferTokenPool.CallOpts)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "getRemotePools", remoteChainSelector)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetRemotePools(&_BurnMintFastTransferTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetRemotePools(&_BurnMintFastTransferTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "getRemoteToken", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetRemoteToken(&_BurnMintFastTransferTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetRemoteToken(&_BurnMintFastTransferTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) GetRmnProxy() (common.Address, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetRmnProxy(&_BurnMintFastTransferTokenPool.CallOpts)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) GetRmnProxy() (common.Address, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetRmnProxy(&_BurnMintFastTransferTokenPool.CallOpts)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) GetRouter() (common.Address, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetRouter(&_BurnMintFastTransferTokenPool.CallOpts)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) GetRouter() (common.Address, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetRouter(&_BurnMintFastTransferTokenPool.CallOpts)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) GetSupportedChains() ([]uint64, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetSupportedChains(&_BurnMintFastTransferTokenPool.CallOpts)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) GetSupportedChains() ([]uint64, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetSupportedChains(&_BurnMintFastTransferTokenPool.CallOpts)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) GetToken() (common.Address, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetToken(&_BurnMintFastTransferTokenPool.CallOpts)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetToken(&_BurnMintFastTransferTokenPool.CallOpts)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) GetTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "getTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) GetTokenDecimals() (uint8, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetTokenDecimals(&_BurnMintFastTransferTokenPool.CallOpts)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) GetTokenDecimals() (uint8, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetTokenDecimals(&_BurnMintFastTransferTokenPool.CallOpts)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) IsAllowedFiller(opts *bind.CallOpts, filler common.Address) (bool, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "isAllowedFiller", filler)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) IsAllowedFiller(filler common.Address) (bool, error) {
	return _BurnMintFastTransferTokenPool.Contract.IsAllowedFiller(&_BurnMintFastTransferTokenPool.CallOpts, filler)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) IsAllowedFiller(filler common.Address) (bool, error) {
	return _BurnMintFastTransferTokenPool.Contract.IsAllowedFiller(&_BurnMintFastTransferTokenPool.CallOpts, filler)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "isRemotePool", remoteChainSelector, remotePoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _BurnMintFastTransferTokenPool.Contract.IsRemotePool(&_BurnMintFastTransferTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _BurnMintFastTransferTokenPool.Contract.IsRemotePool(&_BurnMintFastTransferTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _BurnMintFastTransferTokenPool.Contract.IsSupportedChain(&_BurnMintFastTransferTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _BurnMintFastTransferTokenPool.Contract.IsSupportedChain(&_BurnMintFastTransferTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) IsSupportedToken(token common.Address) (bool, error) {
	return _BurnMintFastTransferTokenPool.Contract.IsSupportedToken(&_BurnMintFastTransferTokenPool.CallOpts, token)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _BurnMintFastTransferTokenPool.Contract.IsSupportedToken(&_BurnMintFastTransferTokenPool.CallOpts, token)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) Owner() (common.Address, error) {
	return _BurnMintFastTransferTokenPool.Contract.Owner(&_BurnMintFastTransferTokenPool.CallOpts)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) Owner() (common.Address, error) {
	return _BurnMintFastTransferTokenPool.Contract.Owner(&_BurnMintFastTransferTokenPool.CallOpts)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintFastTransferTokenPool.Contract.SupportsInterface(&_BurnMintFastTransferTokenPool.CallOpts, interfaceId)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintFastTransferTokenPool.Contract.SupportsInterface(&_BurnMintFastTransferTokenPool.CallOpts, interfaceId)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) TypeAndVersion() (string, error) {
	return _BurnMintFastTransferTokenPool.Contract.TypeAndVersion(&_BurnMintFastTransferTokenPool.CallOpts)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) TypeAndVersion() (string, error) {
	return _BurnMintFastTransferTokenPool.Contract.TypeAndVersion(&_BurnMintFastTransferTokenPool.CallOpts)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.AcceptOwnership(&_BurnMintFastTransferTokenPool.TransactOpts)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.AcceptOwnership(&_BurnMintFastTransferTokenPool.TransactOpts)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactor) AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.contract.Transact(opts, "addRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.AddRemotePool(&_BurnMintFastTransferTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.AddRemotePool(&_BurnMintFastTransferTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.ApplyAllowListUpdates(&_BurnMintFastTransferTokenPool.TransactOpts, removes, adds)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.ApplyAllowListUpdates(&_BurnMintFastTransferTokenPool.TransactOpts, removes, adds)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactor) ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.contract.Transact(opts, "applyChainUpdates", remoteChainSelectorsToRemove, chainsToAdd)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.ApplyChainUpdates(&_BurnMintFastTransferTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.ApplyChainUpdates(&_BurnMintFastTransferTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactor) CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.contract.Transact(opts, "ccipReceive", message)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.CcipReceive(&_BurnMintFastTransferTokenPool.TransactOpts, message)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.CcipReceive(&_BurnMintFastTransferTokenPool.TransactOpts, message)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactor) CcipSendToken(opts *bind.TransactOpts, destinationChainSelector uint64, amount *big.Int, maxFastTransferFee *big.Int, receiver []byte, settlementFeeToken common.Address, extraArgs []byte) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.contract.Transact(opts, "ccipSendToken", destinationChainSelector, amount, maxFastTransferFee, receiver, settlementFeeToken, extraArgs)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) CcipSendToken(destinationChainSelector uint64, amount *big.Int, maxFastTransferFee *big.Int, receiver []byte, settlementFeeToken common.Address, extraArgs []byte) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.CcipSendToken(&_BurnMintFastTransferTokenPool.TransactOpts, destinationChainSelector, amount, maxFastTransferFee, receiver, settlementFeeToken, extraArgs)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorSession) CcipSendToken(destinationChainSelector uint64, amount *big.Int, maxFastTransferFee *big.Int, receiver []byte, settlementFeeToken common.Address, extraArgs []byte) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.CcipSendToken(&_BurnMintFastTransferTokenPool.TransactOpts, destinationChainSelector, amount, maxFastTransferFee, receiver, settlementFeeToken, extraArgs)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactor) FastFill(opts *bind.TransactOpts, fillId [32]byte, settlementId [32]byte, sourceChainSelector uint64, sourceAmountNetFee *big.Int, sourceDecimals uint8, receiver common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.contract.Transact(opts, "fastFill", fillId, settlementId, sourceChainSelector, sourceAmountNetFee, sourceDecimals, receiver)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) FastFill(fillId [32]byte, settlementId [32]byte, sourceChainSelector uint64, sourceAmountNetFee *big.Int, sourceDecimals uint8, receiver common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.FastFill(&_BurnMintFastTransferTokenPool.TransactOpts, fillId, settlementId, sourceChainSelector, sourceAmountNetFee, sourceDecimals, receiver)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorSession) FastFill(fillId [32]byte, settlementId [32]byte, sourceChainSelector uint64, sourceAmountNetFee *big.Int, sourceDecimals uint8, receiver common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.FastFill(&_BurnMintFastTransferTokenPool.TransactOpts, fillId, settlementId, sourceChainSelector, sourceAmountNetFee, sourceDecimals, receiver)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.LockOrBurn(&_BurnMintFastTransferTokenPool.TransactOpts, lockOrBurnIn)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.LockOrBurn(&_BurnMintFastTransferTokenPool.TransactOpts, lockOrBurnIn)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.ReleaseOrMint(&_BurnMintFastTransferTokenPool.TransactOpts, releaseOrMintIn)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.ReleaseOrMint(&_BurnMintFastTransferTokenPool.TransactOpts, releaseOrMintIn)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactor) RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.contract.Transact(opts, "removeRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.RemoveRemotePool(&_BurnMintFastTransferTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.RemoveRemotePool(&_BurnMintFastTransferTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactor) SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.contract.Transact(opts, "setChainRateLimiterConfig", remoteChainSelector, outboundConfig, inboundConfig)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.SetChainRateLimiterConfig(&_BurnMintFastTransferTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.SetChainRateLimiterConfig(&_BurnMintFastTransferTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactor) SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.contract.Transact(opts, "setChainRateLimiterConfigs", remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.SetChainRateLimiterConfigs(&_BurnMintFastTransferTokenPool.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.SetChainRateLimiterConfigs(&_BurnMintFastTransferTokenPool.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactor) SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.contract.Transact(opts, "setRateLimitAdmin", rateLimitAdmin)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.SetRateLimitAdmin(&_BurnMintFastTransferTokenPool.TransactOpts, rateLimitAdmin)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.SetRateLimitAdmin(&_BurnMintFastTransferTokenPool.TransactOpts, rateLimitAdmin)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactor) SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.contract.Transact(opts, "setRouter", newRouter)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.SetRouter(&_BurnMintFastTransferTokenPool.TransactOpts, newRouter)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.SetRouter(&_BurnMintFastTransferTokenPool.TransactOpts, newRouter)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.TransferOwnership(&_BurnMintFastTransferTokenPool.TransactOpts, to)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.TransferOwnership(&_BurnMintFastTransferTokenPool.TransactOpts, to)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactor) UpdateDestChainConfig(opts *bind.TransactOpts, destChainConfigArgs []FastTransferTokenPoolAbstractDestChainConfigUpdateArgs) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.contract.Transact(opts, "updateDestChainConfig", destChainConfigArgs)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) UpdateDestChainConfig(destChainConfigArgs []FastTransferTokenPoolAbstractDestChainConfigUpdateArgs) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.UpdateDestChainConfig(&_BurnMintFastTransferTokenPool.TransactOpts, destChainConfigArgs)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorSession) UpdateDestChainConfig(destChainConfigArgs []FastTransferTokenPoolAbstractDestChainConfigUpdateArgs) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.UpdateDestChainConfig(&_BurnMintFastTransferTokenPool.TransactOpts, destChainConfigArgs)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactor) UpdateFillerAllowList(opts *bind.TransactOpts, fillersToAdd []common.Address, fillersToRemove []common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.contract.Transact(opts, "updateFillerAllowList", fillersToAdd, fillersToRemove)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) UpdateFillerAllowList(fillersToAdd []common.Address, fillersToRemove []common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.UpdateFillerAllowList(&_BurnMintFastTransferTokenPool.TransactOpts, fillersToAdd, fillersToRemove)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorSession) UpdateFillerAllowList(fillersToAdd []common.Address, fillersToRemove []common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.UpdateFillerAllowList(&_BurnMintFastTransferTokenPool.TransactOpts, fillersToAdd, fillersToRemove)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactor) WithdrawPoolFees(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.contract.Transact(opts, "withdrawPoolFees", recipient)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) WithdrawPoolFees(recipient common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.WithdrawPoolFees(&_BurnMintFastTransferTokenPool.TransactOpts, recipient)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorSession) WithdrawPoolFees(recipient common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.WithdrawPoolFees(&_BurnMintFastTransferTokenPool.TransactOpts, recipient)
}

type BurnMintFastTransferTokenPoolAllowListAddIterator struct {
	Event *BurnMintFastTransferTokenPoolAllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolAllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolAllowListAdd)
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
		it.Event = new(BurnMintFastTransferTokenPoolAllowListAdd)
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

func (it *BurnMintFastTransferTokenPoolAllowListAddIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolAllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolAllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterAllowListAdd(opts *bind.FilterOpts) (*BurnMintFastTransferTokenPoolAllowListAddIterator, error) {

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolAllowListAddIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolAllowListAdd) (event.Subscription, error) {

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolAllowListAdd)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseAllowListAdd(log types.Log) (*BurnMintFastTransferTokenPoolAllowListAdd, error) {
	event := new(BurnMintFastTransferTokenPoolAllowListAdd)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintFastTransferTokenPoolAllowListRemoveIterator struct {
	Event *BurnMintFastTransferTokenPoolAllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolAllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolAllowListRemove)
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
		it.Event = new(BurnMintFastTransferTokenPoolAllowListRemove)
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

func (it *BurnMintFastTransferTokenPoolAllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolAllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolAllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterAllowListRemove(opts *bind.FilterOpts) (*BurnMintFastTransferTokenPoolAllowListRemoveIterator, error) {

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolAllowListRemoveIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolAllowListRemove) (event.Subscription, error) {

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolAllowListRemove)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseAllowListRemove(log types.Log) (*BurnMintFastTransferTokenPoolAllowListRemove, error) {
	event := new(BurnMintFastTransferTokenPoolAllowListRemove)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintFastTransferTokenPoolChainAddedIterator struct {
	Event *BurnMintFastTransferTokenPoolChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolChainAdded)
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
		it.Event = new(BurnMintFastTransferTokenPoolChainAdded)
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

func (it *BurnMintFastTransferTokenPoolChainAddedIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolChainAdded struct {
	RemoteChainSelector       uint64
	RemoteToken               []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterChainAdded(opts *bind.FilterOpts) (*BurnMintFastTransferTokenPoolChainAddedIterator, error) {

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolChainAddedIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolChainAdded) (event.Subscription, error) {

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolChainAdded)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseChainAdded(log types.Log) (*BurnMintFastTransferTokenPoolChainAdded, error) {
	event := new(BurnMintFastTransferTokenPoolChainAdded)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintFastTransferTokenPoolChainConfiguredIterator struct {
	Event *BurnMintFastTransferTokenPoolChainConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolChainConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolChainConfigured)
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
		it.Event = new(BurnMintFastTransferTokenPoolChainConfigured)
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

func (it *BurnMintFastTransferTokenPoolChainConfiguredIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolChainConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolChainConfigured struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterChainConfigured(opts *bind.FilterOpts) (*BurnMintFastTransferTokenPoolChainConfiguredIterator, error) {

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolChainConfiguredIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "ChainConfigured", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolChainConfigured) (event.Subscription, error) {

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolChainConfigured)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseChainConfigured(log types.Log) (*BurnMintFastTransferTokenPoolChainConfigured, error) {
	event := new(BurnMintFastTransferTokenPoolChainConfigured)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintFastTransferTokenPoolChainRemovedIterator struct {
	Event *BurnMintFastTransferTokenPoolChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolChainRemoved)
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
		it.Event = new(BurnMintFastTransferTokenPoolChainRemoved)
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

func (it *BurnMintFastTransferTokenPoolChainRemovedIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterChainRemoved(opts *bind.FilterOpts) (*BurnMintFastTransferTokenPoolChainRemovedIterator, error) {

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolChainRemovedIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolChainRemoved) (event.Subscription, error) {

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolChainRemoved)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseChainRemoved(log types.Log) (*BurnMintFastTransferTokenPoolChainRemoved, error) {
	event := new(BurnMintFastTransferTokenPoolChainRemoved)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintFastTransferTokenPoolConfigChangedIterator struct {
	Event *BurnMintFastTransferTokenPoolConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolConfigChanged)
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
		it.Event = new(BurnMintFastTransferTokenPoolConfigChanged)
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

func (it *BurnMintFastTransferTokenPoolConfigChangedIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolConfigChanged struct {
	Config RateLimiterConfig
	Raw    types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*BurnMintFastTransferTokenPoolConfigChangedIterator, error) {

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolConfigChangedIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolConfigChanged) (event.Subscription, error) {

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolConfigChanged)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseConfigChanged(log types.Log) (*BurnMintFastTransferTokenPoolConfigChanged, error) {
	event := new(BurnMintFastTransferTokenPoolConfigChanged)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintFastTransferTokenPoolDestChainConfigUpdatedIterator struct {
	Event *BurnMintFastTransferTokenPoolDestChainConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolDestChainConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolDestChainConfigUpdated)
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
		it.Event = new(BurnMintFastTransferTokenPoolDestChainConfigUpdated)
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

func (it *BurnMintFastTransferTokenPoolDestChainConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolDestChainConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolDestChainConfigUpdated struct {
	DestinationChainSelector uint64
	FastTransferFillerFeeBps uint16
	FastTransferPoolFeeBps   uint16
	MaxFillAmountPerRequest  *big.Int
	DestinationPool          []byte
	ChainFamilySelector      [4]byte
	SettlementOverheadGas    *big.Int
	FillerAllowlistEnabled   bool
	Raw                      types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterDestChainConfigUpdated(opts *bind.FilterOpts, destinationChainSelector []uint64) (*BurnMintFastTransferTokenPoolDestChainConfigUpdatedIterator, error) {

	var destinationChainSelectorRule []interface{}
	for _, destinationChainSelectorItem := range destinationChainSelector {
		destinationChainSelectorRule = append(destinationChainSelectorRule, destinationChainSelectorItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "DestChainConfigUpdated", destinationChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolDestChainConfigUpdatedIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "DestChainConfigUpdated", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchDestChainConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolDestChainConfigUpdated, destinationChainSelector []uint64) (event.Subscription, error) {

	var destinationChainSelectorRule []interface{}
	for _, destinationChainSelectorItem := range destinationChainSelector {
		destinationChainSelectorRule = append(destinationChainSelectorRule, destinationChainSelectorItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "DestChainConfigUpdated", destinationChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolDestChainConfigUpdated)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "DestChainConfigUpdated", log); err != nil {
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseDestChainConfigUpdated(log types.Log) (*BurnMintFastTransferTokenPoolDestChainConfigUpdated, error) {
	event := new(BurnMintFastTransferTokenPoolDestChainConfigUpdated)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "DestChainConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintFastTransferTokenPoolDestinationPoolUpdatedIterator struct {
	Event *BurnMintFastTransferTokenPoolDestinationPoolUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolDestinationPoolUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolDestinationPoolUpdated)
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
		it.Event = new(BurnMintFastTransferTokenPoolDestinationPoolUpdated)
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

func (it *BurnMintFastTransferTokenPoolDestinationPoolUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolDestinationPoolUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolDestinationPoolUpdated struct {
	DestChainSelector uint64
	DestinationPool   common.Address
	Raw               types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterDestinationPoolUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnMintFastTransferTokenPoolDestinationPoolUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "DestinationPoolUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolDestinationPoolUpdatedIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "DestinationPoolUpdated", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchDestinationPoolUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolDestinationPoolUpdated, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "DestinationPoolUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolDestinationPoolUpdated)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "DestinationPoolUpdated", log); err != nil {
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseDestinationPoolUpdated(log types.Log) (*BurnMintFastTransferTokenPoolDestinationPoolUpdated, error) {
	event := new(BurnMintFastTransferTokenPoolDestinationPoolUpdated)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "DestinationPoolUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintFastTransferTokenPoolFastTransferFilledIterator struct {
	Event *BurnMintFastTransferTokenPoolFastTransferFilled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolFastTransferFilledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolFastTransferFilled)
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
		it.Event = new(BurnMintFastTransferTokenPoolFastTransferFilled)
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

func (it *BurnMintFastTransferTokenPoolFastTransferFilledIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolFastTransferFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolFastTransferFilled struct {
	FillId       [32]byte
	SettlementId [32]byte
	Filler       common.Address
	DestAmount   *big.Int
	Receiver     common.Address
	Raw          types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterFastTransferFilled(opts *bind.FilterOpts, fillId [][32]byte, settlementId [][32]byte, filler []common.Address) (*BurnMintFastTransferTokenPoolFastTransferFilledIterator, error) {

	var fillIdRule []interface{}
	for _, fillIdItem := range fillId {
		fillIdRule = append(fillIdRule, fillIdItem)
	}
	var settlementIdRule []interface{}
	for _, settlementIdItem := range settlementId {
		settlementIdRule = append(settlementIdRule, settlementIdItem)
	}
	var fillerRule []interface{}
	for _, fillerItem := range filler {
		fillerRule = append(fillerRule, fillerItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "FastTransferFilled", fillIdRule, settlementIdRule, fillerRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolFastTransferFilledIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "FastTransferFilled", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchFastTransferFilled(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolFastTransferFilled, fillId [][32]byte, settlementId [][32]byte, filler []common.Address) (event.Subscription, error) {

	var fillIdRule []interface{}
	for _, fillIdItem := range fillId {
		fillIdRule = append(fillIdRule, fillIdItem)
	}
	var settlementIdRule []interface{}
	for _, settlementIdItem := range settlementId {
		settlementIdRule = append(settlementIdRule, settlementIdItem)
	}
	var fillerRule []interface{}
	for _, fillerItem := range filler {
		fillerRule = append(fillerRule, fillerItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "FastTransferFilled", fillIdRule, settlementIdRule, fillerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolFastTransferFilled)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "FastTransferFilled", log); err != nil {
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseFastTransferFilled(log types.Log) (*BurnMintFastTransferTokenPoolFastTransferFilled, error) {
	event := new(BurnMintFastTransferTokenPoolFastTransferFilled)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "FastTransferFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintFastTransferTokenPoolFastTransferRequestedIterator struct {
	Event *BurnMintFastTransferTokenPoolFastTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolFastTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolFastTransferRequested)
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
		it.Event = new(BurnMintFastTransferTokenPoolFastTransferRequested)
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

func (it *BurnMintFastTransferTokenPoolFastTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolFastTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolFastTransferRequested struct {
	DestinationChainSelector uint64
	FillId                   [32]byte
	SettlementId             [32]byte
	SourceAmountNetFee       *big.Int
	SourceDecimals           uint8
	FillerFee                *big.Int
	PoolFee                  *big.Int
	DestinationPool          []byte
	Receiver                 []byte
	Raw                      types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterFastTransferRequested(opts *bind.FilterOpts, destinationChainSelector []uint64, fillId [][32]byte, settlementId [][32]byte) (*BurnMintFastTransferTokenPoolFastTransferRequestedIterator, error) {

	var destinationChainSelectorRule []interface{}
	for _, destinationChainSelectorItem := range destinationChainSelector {
		destinationChainSelectorRule = append(destinationChainSelectorRule, destinationChainSelectorItem)
	}
	var fillIdRule []interface{}
	for _, fillIdItem := range fillId {
		fillIdRule = append(fillIdRule, fillIdItem)
	}
	var settlementIdRule []interface{}
	for _, settlementIdItem := range settlementId {
		settlementIdRule = append(settlementIdRule, settlementIdItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "FastTransferRequested", destinationChainSelectorRule, fillIdRule, settlementIdRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolFastTransferRequestedIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "FastTransferRequested", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchFastTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolFastTransferRequested, destinationChainSelector []uint64, fillId [][32]byte, settlementId [][32]byte) (event.Subscription, error) {

	var destinationChainSelectorRule []interface{}
	for _, destinationChainSelectorItem := range destinationChainSelector {
		destinationChainSelectorRule = append(destinationChainSelectorRule, destinationChainSelectorItem)
	}
	var fillIdRule []interface{}
	for _, fillIdItem := range fillId {
		fillIdRule = append(fillIdRule, fillIdItem)
	}
	var settlementIdRule []interface{}
	for _, settlementIdItem := range settlementId {
		settlementIdRule = append(settlementIdRule, settlementIdItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "FastTransferRequested", destinationChainSelectorRule, fillIdRule, settlementIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolFastTransferRequested)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "FastTransferRequested", log); err != nil {
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseFastTransferRequested(log types.Log) (*BurnMintFastTransferTokenPoolFastTransferRequested, error) {
	event := new(BurnMintFastTransferTokenPoolFastTransferRequested)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "FastTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintFastTransferTokenPoolFastTransferSettledIterator struct {
	Event *BurnMintFastTransferTokenPoolFastTransferSettled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolFastTransferSettledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolFastTransferSettled)
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
		it.Event = new(BurnMintFastTransferTokenPoolFastTransferSettled)
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

func (it *BurnMintFastTransferTokenPoolFastTransferSettledIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolFastTransferSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolFastTransferSettled struct {
	FillId                    [32]byte
	SettlementId              [32]byte
	FillerReimbursementAmount *big.Int
	PoolFeeAccumulated        *big.Int
	PrevState                 uint8
	Raw                       types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterFastTransferSettled(opts *bind.FilterOpts, fillId [][32]byte, settlementId [][32]byte) (*BurnMintFastTransferTokenPoolFastTransferSettledIterator, error) {

	var fillIdRule []interface{}
	for _, fillIdItem := range fillId {
		fillIdRule = append(fillIdRule, fillIdItem)
	}
	var settlementIdRule []interface{}
	for _, settlementIdItem := range settlementId {
		settlementIdRule = append(settlementIdRule, settlementIdItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "FastTransferSettled", fillIdRule, settlementIdRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolFastTransferSettledIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "FastTransferSettled", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchFastTransferSettled(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolFastTransferSettled, fillId [][32]byte, settlementId [][32]byte) (event.Subscription, error) {

	var fillIdRule []interface{}
	for _, fillIdItem := range fillId {
		fillIdRule = append(fillIdRule, fillIdItem)
	}
	var settlementIdRule []interface{}
	for _, settlementIdItem := range settlementId {
		settlementIdRule = append(settlementIdRule, settlementIdItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "FastTransferSettled", fillIdRule, settlementIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolFastTransferSettled)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "FastTransferSettled", log); err != nil {
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseFastTransferSettled(log types.Log) (*BurnMintFastTransferTokenPoolFastTransferSettled, error) {
	event := new(BurnMintFastTransferTokenPoolFastTransferSettled)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "FastTransferSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintFastTransferTokenPoolFillerAllowListUpdatedIterator struct {
	Event *BurnMintFastTransferTokenPoolFillerAllowListUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolFillerAllowListUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolFillerAllowListUpdated)
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
		it.Event = new(BurnMintFastTransferTokenPoolFillerAllowListUpdated)
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

func (it *BurnMintFastTransferTokenPoolFillerAllowListUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolFillerAllowListUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolFillerAllowListUpdated struct {
	AddFillers    []common.Address
	RemoveFillers []common.Address
	Raw           types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterFillerAllowListUpdated(opts *bind.FilterOpts) (*BurnMintFastTransferTokenPoolFillerAllowListUpdatedIterator, error) {

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "FillerAllowListUpdated")
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolFillerAllowListUpdatedIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "FillerAllowListUpdated", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchFillerAllowListUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolFillerAllowListUpdated) (event.Subscription, error) {

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "FillerAllowListUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolFillerAllowListUpdated)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "FillerAllowListUpdated", log); err != nil {
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseFillerAllowListUpdated(log types.Log) (*BurnMintFastTransferTokenPoolFillerAllowListUpdated, error) {
	event := new(BurnMintFastTransferTokenPoolFillerAllowListUpdated)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "FillerAllowListUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintFastTransferTokenPoolInboundRateLimitConsumedIterator struct {
	Event *BurnMintFastTransferTokenPoolInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolInboundRateLimitConsumed)
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
		it.Event = new(BurnMintFastTransferTokenPoolInboundRateLimitConsumed)
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

func (it *BurnMintFastTransferTokenPoolInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintFastTransferTokenPoolInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolInboundRateLimitConsumedIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "InboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolInboundRateLimitConsumed)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseInboundRateLimitConsumed(log types.Log) (*BurnMintFastTransferTokenPoolInboundRateLimitConsumed, error) {
	event := new(BurnMintFastTransferTokenPoolInboundRateLimitConsumed)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintFastTransferTokenPoolLockedOrBurnedIterator struct {
	Event *BurnMintFastTransferTokenPoolLockedOrBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolLockedOrBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolLockedOrBurned)
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
		it.Event = new(BurnMintFastTransferTokenPoolLockedOrBurned)
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

func (it *BurnMintFastTransferTokenPoolLockedOrBurnedIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolLockedOrBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolLockedOrBurned struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintFastTransferTokenPoolLockedOrBurnedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolLockedOrBurnedIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "LockedOrBurned", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolLockedOrBurned)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseLockedOrBurned(log types.Log) (*BurnMintFastTransferTokenPoolLockedOrBurned, error) {
	event := new(BurnMintFastTransferTokenPoolLockedOrBurned)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintFastTransferTokenPoolOutboundRateLimitConsumedIterator struct {
	Event *BurnMintFastTransferTokenPoolOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolOutboundRateLimitConsumed)
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
		it.Event = new(BurnMintFastTransferTokenPoolOutboundRateLimitConsumed)
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

func (it *BurnMintFastTransferTokenPoolOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintFastTransferTokenPoolOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolOutboundRateLimitConsumedIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "OutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolOutboundRateLimitConsumed)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseOutboundRateLimitConsumed(log types.Log) (*BurnMintFastTransferTokenPoolOutboundRateLimitConsumed, error) {
	event := new(BurnMintFastTransferTokenPoolOutboundRateLimitConsumed)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintFastTransferTokenPoolOwnershipTransferRequestedIterator struct {
	Event *BurnMintFastTransferTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolOwnershipTransferRequested)
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
		it.Event = new(BurnMintFastTransferTokenPoolOwnershipTransferRequested)
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

func (it *BurnMintFastTransferTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintFastTransferTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolOwnershipTransferRequestedIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolOwnershipTransferRequested)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*BurnMintFastTransferTokenPoolOwnershipTransferRequested, error) {
	event := new(BurnMintFastTransferTokenPoolOwnershipTransferRequested)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintFastTransferTokenPoolOwnershipTransferredIterator struct {
	Event *BurnMintFastTransferTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolOwnershipTransferred)
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
		it.Event = new(BurnMintFastTransferTokenPoolOwnershipTransferred)
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

func (it *BurnMintFastTransferTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintFastTransferTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolOwnershipTransferredIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolOwnershipTransferred)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*BurnMintFastTransferTokenPoolOwnershipTransferred, error) {
	event := new(BurnMintFastTransferTokenPoolOwnershipTransferred)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintFastTransferTokenPoolPoolFeeWithdrawnIterator struct {
	Event *BurnMintFastTransferTokenPoolPoolFeeWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolPoolFeeWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolPoolFeeWithdrawn)
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
		it.Event = new(BurnMintFastTransferTokenPoolPoolFeeWithdrawn)
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

func (it *BurnMintFastTransferTokenPoolPoolFeeWithdrawnIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolPoolFeeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolPoolFeeWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterPoolFeeWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*BurnMintFastTransferTokenPoolPoolFeeWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "PoolFeeWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolPoolFeeWithdrawnIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "PoolFeeWithdrawn", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchPoolFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolPoolFeeWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "PoolFeeWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolPoolFeeWithdrawn)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "PoolFeeWithdrawn", log); err != nil {
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParsePoolFeeWithdrawn(log types.Log) (*BurnMintFastTransferTokenPoolPoolFeeWithdrawn, error) {
	event := new(BurnMintFastTransferTokenPoolPoolFeeWithdrawn)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "PoolFeeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintFastTransferTokenPoolRateLimitAdminSetIterator struct {
	Event *BurnMintFastTransferTokenPoolRateLimitAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolRateLimitAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolRateLimitAdminSet)
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
		it.Event = new(BurnMintFastTransferTokenPoolRateLimitAdminSet)
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

func (it *BurnMintFastTransferTokenPoolRateLimitAdminSetIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolRateLimitAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolRateLimitAdminSet struct {
	RateLimitAdmin common.Address
	Raw            types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterRateLimitAdminSet(opts *bind.FilterOpts) (*BurnMintFastTransferTokenPoolRateLimitAdminSetIterator, error) {

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolRateLimitAdminSetIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "RateLimitAdminSet", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolRateLimitAdminSet) (event.Subscription, error) {

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolRateLimitAdminSet)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseRateLimitAdminSet(log types.Log) (*BurnMintFastTransferTokenPoolRateLimitAdminSet, error) {
	event := new(BurnMintFastTransferTokenPoolRateLimitAdminSet)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintFastTransferTokenPoolReleasedOrMintedIterator struct {
	Event *BurnMintFastTransferTokenPoolReleasedOrMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolReleasedOrMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolReleasedOrMinted)
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
		it.Event = new(BurnMintFastTransferTokenPoolReleasedOrMinted)
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

func (it *BurnMintFastTransferTokenPoolReleasedOrMintedIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolReleasedOrMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolReleasedOrMinted struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Recipient           common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintFastTransferTokenPoolReleasedOrMintedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolReleasedOrMintedIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "ReleasedOrMinted", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolReleasedOrMinted)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseReleasedOrMinted(log types.Log) (*BurnMintFastTransferTokenPoolReleasedOrMinted, error) {
	event := new(BurnMintFastTransferTokenPoolReleasedOrMinted)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintFastTransferTokenPoolRemotePoolAddedIterator struct {
	Event *BurnMintFastTransferTokenPoolRemotePoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolRemotePoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolRemotePoolAdded)
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
		it.Event = new(BurnMintFastTransferTokenPoolRemotePoolAdded)
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

func (it *BurnMintFastTransferTokenPoolRemotePoolAddedIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolRemotePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolRemotePoolAdded struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintFastTransferTokenPoolRemotePoolAddedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolRemotePoolAddedIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "RemotePoolAdded", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolRemotePoolAdded)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseRemotePoolAdded(log types.Log) (*BurnMintFastTransferTokenPoolRemotePoolAdded, error) {
	event := new(BurnMintFastTransferTokenPoolRemotePoolAdded)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintFastTransferTokenPoolRemotePoolRemovedIterator struct {
	Event *BurnMintFastTransferTokenPoolRemotePoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolRemotePoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolRemotePoolRemoved)
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
		it.Event = new(BurnMintFastTransferTokenPoolRemotePoolRemoved)
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

func (it *BurnMintFastTransferTokenPoolRemotePoolRemovedIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolRemotePoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolRemotePoolRemoved struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintFastTransferTokenPoolRemotePoolRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolRemotePoolRemovedIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "RemotePoolRemoved", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolRemotePoolRemoved)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseRemotePoolRemoved(log types.Log) (*BurnMintFastTransferTokenPoolRemotePoolRemoved, error) {
	event := new(BurnMintFastTransferTokenPoolRemotePoolRemoved)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintFastTransferTokenPoolRouterUpdatedIterator struct {
	Event *BurnMintFastTransferTokenPoolRouterUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolRouterUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolRouterUpdated)
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
		it.Event = new(BurnMintFastTransferTokenPoolRouterUpdated)
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

func (it *BurnMintFastTransferTokenPoolRouterUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolRouterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolRouterUpdated struct {
	OldRouter common.Address
	NewRouter common.Address
	Raw       types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterRouterUpdated(opts *bind.FilterOpts) (*BurnMintFastTransferTokenPoolRouterUpdatedIterator, error) {

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolRouterUpdatedIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "RouterUpdated", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolRouterUpdated) (event.Subscription, error) {

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolRouterUpdated)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseRouterUpdated(log types.Log) (*BurnMintFastTransferTokenPoolRouterUpdated, error) {
	event := new(BurnMintFastTransferTokenPoolRouterUpdated)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetCCVs struct {
	RequiredCCVs      []common.Address
	OptionalCCVs      []common.Address
	OptionalThreshold uint8
}

func (BurnMintFastTransferTokenPoolAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (BurnMintFastTransferTokenPoolAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (BurnMintFastTransferTokenPoolChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (BurnMintFastTransferTokenPoolChainConfigured) Topic() common.Hash {
	return common.HexToHash("0x0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b")
}

func (BurnMintFastTransferTokenPoolChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (BurnMintFastTransferTokenPoolConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (BurnMintFastTransferTokenPoolDestChainConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x6cfec31453105612e33aed8011f0e249b68d55e4efa65374322eb7ceeee76fbd")
}

func (BurnMintFastTransferTokenPoolDestinationPoolUpdated) Topic() common.Hash {
	return common.HexToHash("0xb760e03fa04c0e86fcff6d0046cdcf22fb5d5b6a17d1e6f890b3456e81c40fd8")
}

func (BurnMintFastTransferTokenPoolFastTransferFilled) Topic() common.Hash {
	return common.HexToHash("0xd6f70fb263bfe7d01ec6802b3c07b6bd32579760fe9fcb4e248a036debb8cdf1")
}

func (BurnMintFastTransferTokenPoolFastTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xfa7d3740fa7611df3f0d8d8c3aa1ed57c4fffaf2dcd0c47535f18a4774b44acd")
}

func (BurnMintFastTransferTokenPoolFastTransferSettled) Topic() common.Hash {
	return common.HexToHash("0x33e17439bb4d31426d9168fc32af3a69cfce0467ba0d532fa804c27b5ff2189c")
}

func (BurnMintFastTransferTokenPoolFillerAllowListUpdated) Topic() common.Hash {
	return common.HexToHash("0xfd35c599d42a981cbb1bbf7d3e6d9855a59f5c994ec6b427118ee0c260e24193")
}

func (BurnMintFastTransferTokenPoolInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (BurnMintFastTransferTokenPoolLockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
}

func (BurnMintFastTransferTokenPoolOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944")
}

func (BurnMintFastTransferTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (BurnMintFastTransferTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (BurnMintFastTransferTokenPoolPoolFeeWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b859599")
}

func (BurnMintFastTransferTokenPoolRateLimitAdminSet) Topic() common.Hash {
	return common.HexToHash("0x44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174")
}

func (BurnMintFastTransferTokenPoolReleasedOrMinted) Topic() common.Hash {
	return common.HexToHash("0xfc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0")
}

func (BurnMintFastTransferTokenPoolRemotePoolAdded) Topic() common.Hash {
	return common.HexToHash("0x7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea")
}

func (BurnMintFastTransferTokenPoolRemotePoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76")
}

func (BurnMintFastTransferTokenPoolRouterUpdated) Topic() common.Hash {
	return common.HexToHash("0x02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f1684")
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPool) Address() common.Address {
	return _BurnMintFastTransferTokenPool.address
}

type BurnMintFastTransferTokenPoolInterface interface {
	ComputeFillId(opts *bind.CallOpts, settlementId [32]byte, sourceChainSelector uint64, sourceAmountNetFee *big.Int, sourceDecimals uint8, receiver []byte) ([32]byte, error)

	GetAccumulatedPoolFees(opts *bind.CallOpts) (*big.Int, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetAllowedFillers(opts *bind.CallOpts) ([]common.Address, error)

	GetCCVs(opts *bind.CallOpts, arg0 uint64) (GetCCVs,

		error)

	GetCcipSendTokenFee(opts *bind.CallOpts, destinationChainSelector uint64, amount *big.Int, receiver []byte, settlementFeeToken common.Address, extraArgs []byte) (IFastTransferPoolQuote, error)

	GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetDestChainConfig(opts *bind.CallOpts, remoteChainSelector uint64) (FastTransferTokenPoolAbstractDestChainConfig, []common.Address, error)

	GetFillInfo(opts *bind.CallOpts, fillId [32]byte) (FastTransferTokenPoolAbstractFillInfo, error)

	GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	GetSupportedChains(opts *bind.CallOpts) ([]uint64, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	GetTokenDecimals(opts *bind.CallOpts) (uint8, error)

	IsAllowedFiller(opts *bind.CallOpts, filler common.Address) (bool, error)

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

	CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error)

	CcipSendToken(opts *bind.TransactOpts, destinationChainSelector uint64, amount *big.Int, maxFastTransferFee *big.Int, receiver []byte, settlementFeeToken common.Address, extraArgs []byte) (*types.Transaction, error)

	FastFill(opts *bind.TransactOpts, fillId [32]byte, settlementId [32]byte, sourceChainSelector uint64, sourceAmountNetFee *big.Int, sourceDecimals uint8, receiver common.Address) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error)

	SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error)

	SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateDestChainConfig(opts *bind.TransactOpts, destChainConfigArgs []FastTransferTokenPoolAbstractDestChainConfigUpdateArgs) (*types.Transaction, error)

	UpdateFillerAllowList(opts *bind.TransactOpts, fillersToAdd []common.Address, fillersToRemove []common.Address) (*types.Transaction, error)

	WithdrawPoolFees(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error)

	FilterAllowListAdd(opts *bind.FilterOpts) (*BurnMintFastTransferTokenPoolAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*BurnMintFastTransferTokenPoolAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*BurnMintFastTransferTokenPoolAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*BurnMintFastTransferTokenPoolAllowListRemove, error)

	FilterChainAdded(opts *bind.FilterOpts) (*BurnMintFastTransferTokenPoolChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*BurnMintFastTransferTokenPoolChainAdded, error)

	FilterChainConfigured(opts *bind.FilterOpts) (*BurnMintFastTransferTokenPoolChainConfiguredIterator, error)

	WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolChainConfigured) (event.Subscription, error)

	ParseChainConfigured(log types.Log) (*BurnMintFastTransferTokenPoolChainConfigured, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*BurnMintFastTransferTokenPoolChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*BurnMintFastTransferTokenPoolChainRemoved, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*BurnMintFastTransferTokenPoolConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*BurnMintFastTransferTokenPoolConfigChanged, error)

	FilterDestChainConfigUpdated(opts *bind.FilterOpts, destinationChainSelector []uint64) (*BurnMintFastTransferTokenPoolDestChainConfigUpdatedIterator, error)

	WatchDestChainConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolDestChainConfigUpdated, destinationChainSelector []uint64) (event.Subscription, error)

	ParseDestChainConfigUpdated(log types.Log) (*BurnMintFastTransferTokenPoolDestChainConfigUpdated, error)

	FilterDestinationPoolUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnMintFastTransferTokenPoolDestinationPoolUpdatedIterator, error)

	WatchDestinationPoolUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolDestinationPoolUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseDestinationPoolUpdated(log types.Log) (*BurnMintFastTransferTokenPoolDestinationPoolUpdated, error)

	FilterFastTransferFilled(opts *bind.FilterOpts, fillId [][32]byte, settlementId [][32]byte, filler []common.Address) (*BurnMintFastTransferTokenPoolFastTransferFilledIterator, error)

	WatchFastTransferFilled(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolFastTransferFilled, fillId [][32]byte, settlementId [][32]byte, filler []common.Address) (event.Subscription, error)

	ParseFastTransferFilled(log types.Log) (*BurnMintFastTransferTokenPoolFastTransferFilled, error)

	FilterFastTransferRequested(opts *bind.FilterOpts, destinationChainSelector []uint64, fillId [][32]byte, settlementId [][32]byte) (*BurnMintFastTransferTokenPoolFastTransferRequestedIterator, error)

	WatchFastTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolFastTransferRequested, destinationChainSelector []uint64, fillId [][32]byte, settlementId [][32]byte) (event.Subscription, error)

	ParseFastTransferRequested(log types.Log) (*BurnMintFastTransferTokenPoolFastTransferRequested, error)

	FilterFastTransferSettled(opts *bind.FilterOpts, fillId [][32]byte, settlementId [][32]byte) (*BurnMintFastTransferTokenPoolFastTransferSettledIterator, error)

	WatchFastTransferSettled(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolFastTransferSettled, fillId [][32]byte, settlementId [][32]byte) (event.Subscription, error)

	ParseFastTransferSettled(log types.Log) (*BurnMintFastTransferTokenPoolFastTransferSettled, error)

	FilterFillerAllowListUpdated(opts *bind.FilterOpts) (*BurnMintFastTransferTokenPoolFillerAllowListUpdatedIterator, error)

	WatchFillerAllowListUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolFillerAllowListUpdated) (event.Subscription, error)

	ParseFillerAllowListUpdated(log types.Log) (*BurnMintFastTransferTokenPoolFillerAllowListUpdated, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintFastTransferTokenPoolInboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*BurnMintFastTransferTokenPoolInboundRateLimitConsumed, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintFastTransferTokenPoolLockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*BurnMintFastTransferTokenPoolLockedOrBurned, error)

	FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintFastTransferTokenPoolOutboundRateLimitConsumedIterator, error)

	WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseOutboundRateLimitConsumed(log types.Log) (*BurnMintFastTransferTokenPoolOutboundRateLimitConsumed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintFastTransferTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*BurnMintFastTransferTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintFastTransferTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*BurnMintFastTransferTokenPoolOwnershipTransferred, error)

	FilterPoolFeeWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*BurnMintFastTransferTokenPoolPoolFeeWithdrawnIterator, error)

	WatchPoolFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolPoolFeeWithdrawn, recipient []common.Address) (event.Subscription, error)

	ParsePoolFeeWithdrawn(log types.Log) (*BurnMintFastTransferTokenPoolPoolFeeWithdrawn, error)

	FilterRateLimitAdminSet(opts *bind.FilterOpts) (*BurnMintFastTransferTokenPoolRateLimitAdminSetIterator, error)

	WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolRateLimitAdminSet) (event.Subscription, error)

	ParseRateLimitAdminSet(log types.Log) (*BurnMintFastTransferTokenPoolRateLimitAdminSet, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintFastTransferTokenPoolReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*BurnMintFastTransferTokenPoolReleasedOrMinted, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintFastTransferTokenPoolRemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*BurnMintFastTransferTokenPoolRemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintFastTransferTokenPoolRemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*BurnMintFastTransferTokenPoolRemotePoolRemoved, error)

	FilterRouterUpdated(opts *bind.FilterOpts) (*BurnMintFastTransferTokenPoolRouterUpdatedIterator, error)

	WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolRouterUpdated) (event.Subscription, error)

	ParseRouterUpdated(log types.Log) (*BurnMintFastTransferTokenPoolRouterUpdated, error)

	Address() common.Address
}
