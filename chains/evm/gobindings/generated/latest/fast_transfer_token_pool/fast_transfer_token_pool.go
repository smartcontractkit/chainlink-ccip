// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fast_transfer_token_pool

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated"
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
	MaxFillAmountPerRequest *big.Int
	FillerAllowlistEnabled  bool
	FastTransferBpsFee      uint16
	SettlementOverheadGas   uint32
	DestinationPool         []byte
	CustomExtraArgs         []byte
}

type FastTransferTokenPoolAbstractDestChainConfigUpdateArgs struct {
	FillerAllowlistEnabled  bool
	FastTransferBpsFee      uint16
	SettlementOverheadGas   uint32
	RemoteChainSelector     uint64
	ChainFamilySelector     [4]byte
	MaxFillAmountPerRequest *big.Int
	DestinationPool         []byte
	CustomExtraArgs         []byte
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
	OriginalSender      []byte
	RemoteChainSelector uint64
	Receiver            common.Address
	Amount              *big.Int
	LocalToken          common.Address
	SourcePoolAddress   []byte
	SourcePoolData      []byte
	OffchainTokenData   []byte
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIBurnMintERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipReceive\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.Any2EVMMessage\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipSendToken\",\"inputs\":[{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"settlementId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"computeFillId\",\"inputs\":[{\"name\":\"settlementId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceAmountNetFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourceDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"fastFill\",\"inputs\":[{\"name\":\"settlementId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fillId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceAmountNetFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourceDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowedFillers\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCcipSendTokenFee\",\"inputs\":[{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"settlementFeeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIFastTransferPool.Quote\",\"components\":[{\"name\":\"ccipSettlementFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fastTransferFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFastTransferTokenPoolAbstract.DestChainConfig\",\"components\":[{\"name\":\"maxFillAmountPerRequest\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fillerAllowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"fastTransferBpsFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"settlementOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destinationPool\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"customExtraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFillInfo\",\"inputs\":[{\"name\":\"fillId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFastTransferTokenPoolAbstract.FillInfo\",\"components\":[{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"enumFastTransferTokenPoolAbstract.FillState\"},{\"name\":\"filler\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAllowedFiller\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"filler\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRouter\",\"inputs\":[{\"name\":\"newRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateDestChainConfig\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFastTransferTokenPoolAbstract.DestChainConfigUpdateArgs[]\",\"components\":[{\"name\":\"fillerAllowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"fastTransferBpsFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"settlementOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"maxFillAmountPerRequest\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationPool\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"customExtraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateFillerAllowList\",\"inputs\":[{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fillersToAdd\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"fillersToRemove\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigUpdated\",\"inputs\":[{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"fastTransferBpsFee\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"maxFillAmountPerRequest\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"destinationPool\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"},{\"name\":\"settlementOverheadGas\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fillerAllowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestinationPoolUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"destinationPool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FastTransferFilled\",\"inputs\":[{\"name\":\"fillId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"settlementId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"filler\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"destAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FastTransferRequested\",\"inputs\":[{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"fillId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"settlementId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sourceAmountNetFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fastTransferFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FastTransferSettled\",\"inputs\":[{\"name\":\"fillId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"settlementId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FillerAllowListUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"addFillers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"removeFillers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RouterUpdated\",\"inputs\":[{\"name\":\"oldRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyFilled\",\"inputs\":[{\"name\":\"fillId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AlreadySettled\",\"inputs\":[{\"name\":\"fillId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"FillerNotAllowlisted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"filler\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRouter\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidSettlementId\",\"inputs\":[{\"name\":\"settlementId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TransferAmountExceedsMaxFillAmount\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x61012080604052346103de5761709d803803809161001d828561045d565b8339810160a0828203126103de5781516001600160a01b038116908190036103de5761004b60208401610480565b60408401519091906001600160401b0381116103de5784019280601f850112156103de578351936001600160401b038511610447578460051b906020820195610097604051978861045d565b86526020808701928201019283116103de57602001905b82821061042f575050506100d060806100c96060870161048e565b950161048e565b93331561041e57600180546001600160a01b031916331790558115801561040d575b80156103fc575b6103eb578160209160049360805260c0526040519283809263313ce56760e01b82525afa600091816103aa575b5061037f575b5060a052600480546001600160a01b0319166001600160a01b0384169081179091558151151560e081905290919061025d575b50156102475761010052604051616a5a908161064382396080518181816117f801528181611a01015281816120d70152818161274401528181613d1101528181613f4a015281816144f001528181614568015281816155b20152818161596801528181615a160152615bed015260a051818181611a4501528181611f3b015281816141d00152818161447b015281816154460152615490015260c051818181610d040152818161189301528181611e8501528181612baf01528181613dac01526140e9015260e051818181610c230152818161347701526166d8015261010051816122420152f35b6335fdcccd60e21b600052600060045260246000fd5b906020906040519061026f838361045d565b60008252600036813760e0511561036e5760005b82518110156102ea576001906001600160a01b036102a182866104a2565b5116856102ad826104e4565b6102ba575b505001610283565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a138856102b2565b5092905060005b8151811015610365576001906001600160a01b0361030f82856104a2565b5116801561035f5784610321826105e2565b61032f575b50505b016102f1565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a13884610326565b50610329565b5050503861015f565b6335f4a7b360e01b60005260046000fd5b60ff1660ff8216818103610393575061012c565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d6020116103e3575b816103c66020938361045d565b810103126103de576103d790610480565b9038610126565b600080fd5b3d91506103b9565b6342bcdf7f60e11b60005260046000fd5b506001600160a01b038116156100f9565b506001600160a01b038516156100f2565b639b15e16f60e01b60005260046000fd5b6020809161043c8461048e565b8152019101906100ae565b634e487b7160e01b600052604160045260246000fd5b601f909101601f19168101906001600160401b0382119082101761044757604052565b519060ff821682036103de57565b51906001600160a01b03821682036103de57565b80518210156104b65760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b80548210156104b65760005260206000200190600090565b60008181526003602052604090205480156105db5760001981018181116105c5576002546000198101919082116105c557818103610574575b505050600254801561055e57600019016105388160026104cc565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b6105ad6105856105969360026104cc565b90549060031b1c92839260026104cc565b819391549060031b91821b91600019901b19161790565b9055600052600360205260406000205538808061051d565b634e487b7160e01b600052601160045260246000fd5b5050600090565b8060005260036020526040600020541560001461063c57600254680100000000000000008110156104475761062361059682600185940160025560026104cc565b9055600254906000526003602052604060002055600190565b5060009056fe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a71461463057508063181f5a771461458c57806321df0da71461451d578063240028e81461449f57806324f65ee7146144435780632e7aa8c81461401b5780633907753714613c6a57806348c16059146136225780634c5ef0ed146135bf57806354c8a4f31461344557806362ddd3c4146133c2578063669c6f16146132675780636d3d1a58146132155780636def4ce71461309357806379ba509714612faa5780637d54534e14612eff57806385572ffb1461291d57806387f060d01461260a5780638926f54f146125b15780638da5cb5b1461255f57806392575c3b14611dc6578063962d402014611c525780639a4575b9146117b0578063a42a7b8b14611624578063a4e27e741461158f578063a7cd63b714611517578063abe1c1e814611432578063acfecf911461130d578063af58d59f146112a5578063b0f479a114611253578063b7946580146111fd578063c0d78655146110fc578063c4bffe2b14610fae578063c75eea9c14610eeb578063cf7401f314610d28578063dc0bd97114610cb9578063dd687aea14610c48578063e0351e1314610bed578063e8a1da171461034d578063eeebc674146102d75763f2fde38b146101e257600080fd5b346102d25760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d25773ffffffffffffffffffffffffffffffffffffffff61022e614969565b61023661565c565b163381146102a857807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346102d25760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d25760443560ff811681036102d25760643567ffffffffffffffff81116102d257602091610339610345923690600401614b1a565b9060243560043561531c565b604051908152f35b346102d25761035b36614b38565b91909261036661565c565b6000905b828210610a445750505060009063ffffffff4216907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee184360301925b81811015610a42576000918160051b86013585811215610a3e5786019061012082360312610a3e57604051956103db876147b1565b6103e4836149c4565b8752602083013567ffffffffffffffff8111610a3a5783019536601f88011215610a3a5786359661041488614c03565b97610422604051998a614888565b8089526020808a019160051b83010190368211610a365760208301905b828210610a03575050505060208801968752604084013567ffffffffffffffff81116109ff576104729036908601614b1a565b926040890193845261049c61048a3660608801614d91565b9560608b0196875260c0369101614d91565b9660808a019788526104ae8651615ce0565b6104b88851615ce0565b845151156109d7576104d467ffffffffffffffff8b5116615ede565b156109a05767ffffffffffffffff8a5116815260076020526040812061061487516fffffffffffffffffffffffffffffffff604082015116906105cf6fffffffffffffffffffffffffffffffff6020830151169151151583608060405161053a816147b1565b858152602081018c905260408101849052606081018690520152855474ff000000000000000000000000000000000000000091151560a01b919091167fffffffffffffffffffffff0000000000000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff84161773ffffffff0000000000000000000000000000000060808b901b1617178555565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176001830155565b61073a89516fffffffffffffffffffffffffffffffff604082015116906106f56fffffffffffffffffffffffffffffffff6020830151169151151583608060405161065e816147b1565b858152602081018c9052604081018490526060810186905201526002860180547fffffffffffffffffffffff000000000000000000000000000000000000000000166fffffffffffffffffffffffffffffffff85161773ffffffff0000000000000000000000000000000060808c901b161791151560a01b74ff000000000000000000000000000000000000000016919091179055565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176003830155565b6004865191019080519067ffffffffffffffff821161097357610767826107618554614efd565b856152d7565b602090601f83116001146108d4576107b492918591836108c9575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b885180518210156107ec57906107e66001926107df8367ffffffffffffffff8f511692614eba565b5190615701565b016107b7565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c29391999750956108ba67ffffffffffffffff600197969498511692519351915161088661085160405196879687526101006020880152610100870190614926565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a1019391939290926103a6565b015190508f80610782565b83855281852091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416865b81811061095b5750908460019594939210610924575b505050811b0190556107b7565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558e8080610917565b92936020600181928786015181550195019301610901565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60249067ffffffffffffffff8b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b807f8579befe0000000000000000000000000000000000000000000000000000000060049252fd5b8680fd5b813567ffffffffffffffff8111610a3257602091610a278392833691890101614b1a565b81520191019061043f565b8a80fd5b8880fd5b8580fd5b8380fd5b005b909267ffffffffffffffff610a65610a608686869997996151da565b614e68565b1692610a7084616828565b15610bbf57836000526007602052610a8e60056040600020016160c4565b9260005b8451811015610aca57600190866000526007602052610ac36005604060002001610abc8389614eba565b51906168f8565b5001610a92565b5093909491959250806000526007602052600560406000206000815560006001820155600060028201556000600382015560048101610b098154614efd565b9081610b7c575b5050018054906000815581610b5b575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a101909193929361036a565b6000526020600020908101905b81811015610b205760008155600101610b68565b81601f60009311600114610b945750555b8880610b10565b81835260208320610baf91601f01861c8101906001016152ad565b8082528160208120915555610b8d565b837f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346102d25760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d25760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b346102d25760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d25767ffffffffffffffff610c886149ad565b16600052600b602052610cb5610ca160406000206160c4565b604051918291602083526020830190614c8a565b0390f35b346102d25760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d257602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102d25760e07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d257610d5f6149ad565b60607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc3601126102d257604051610d958161486c565b60243580151581036102d25781526044356fffffffffffffffffffffffffffffffff811681036102d25760208201526064356fffffffffffffffffffffffffffffffff811681036102d257604082015260607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c3601126102d25760405190610e1c8261486c565b60843580151581036102d257825260a4356fffffffffffffffffffffffffffffffff811681036102d257602083015260c4356fffffffffffffffffffffffffffffffff811681036102d257604083015273ffffffffffffffffffffffffffffffffffffffff6009541633141580610ec9575b610e9b57610a4292615aa6565b7f8e4a23d6000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b5073ffffffffffffffffffffffffffffffffffffffff60015416331415610e8e565b346102d25760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d25767ffffffffffffffff610f2b6149ad565b610f336151fa565b50166000526007602052610cb5610f55610f506040600020615225565b615c5b565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b346102d25760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d2576040516005548082528160208101600560005260206000209260005b8181106110e357505061100e92500382614888565b80519061103361101d83614c03565b9261102b6040519485614888565b808452614c03565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe060208401920136833760005b8151811015611093578067ffffffffffffffff61108060019385614eba565b511661108c8287614eba565b5201611061565b5050906040519182916020830190602084525180915260408301919060005b8181106110c0575050500390f35b825167ffffffffffffffff168452859450602093840193909201916001016110b2565b8454835260019485019486945060209093019201610ff9565b346102d25760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d257611133614969565b61113b61565c565b73ffffffffffffffffffffffffffffffffffffffff81169081156111d357600480547fffffffffffffffffffffffff000000000000000000000000000000000000000081169390931790556040805173ffffffffffffffffffffffffffffffffffffffff93841681529190921660208201527f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f16849190a1005b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b346102d25760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d257610cb561123f61123a6149ad565b61528b565b604051918291602083526020830190614926565b346102d25760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d257602073ffffffffffffffffffffffffffffffffffffffff60045416604051908152f35b346102d25760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d25767ffffffffffffffff6112e56149ad565b6112ed6151fa565b50166000526007602052610cb5610f55610f506002604060002001615225565b346102d25767ffffffffffffffff61132436614ba4565b92909161132f61565c565b1690611348826000526006602052604060002054151590565b1561140457816000526007602052611379600560406000200161136c368685614ae3565b60208151910120906168f8565b156113bd577f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d7691926113b860405192839260208452602084019161519b565b0390a2005b611400906040519384937f74f23c7c000000000000000000000000000000000000000000000000000000008552600485015260406024850152604484019161519b565b0390fd5b507f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346102d25760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d257611469614ddd565b50600435600052600c602052604060002060405190611487826147fc565b549061149660ff831682615012565b73ffffffffffffffffffffffffffffffffffffffff602082019260081c168252604051905160038110156114e85760409273ffffffffffffffffffffffffffffffffffffffff91835251166020820152f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b346102d25760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d2576040516002548082526020820190600260005260206000209060005b81811061157957610cb585610ca181870382614888565b8254845260209093019260019283019201611562565b346102d25760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d2576115c66149ad565b6024359073ffffffffffffffffffffffffffffffffffffffff82168092036102d25760209167ffffffffffffffff61161a9216600052600b8352604060002060019160005201602052604060002054151590565b6040519015158152f35b346102d25760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d25767ffffffffffffffff6116646149ad565b16600052600760205261167d60056040600020016160c4565b8051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06116c36116ad84614c03565b936116bb6040519586614888565b808552614c03565b0160005b81811061179f57505060005b815181101561171b57806116e960019284614eba565b5160005260086020526116ff6040600020614f50565b6117098286614eba565b526117148185614eba565b50016116d3565b826040518091602082016020835281518091526040830190602060408260051b8601019301916000905b82821061175457505050500390f35b9193602061178f827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851614926565b9601920192018594939192611745565b8060606020809387010152016116c7565b346102d2576117be36614cd4565b606060206040516117ce816147fc565b8281520152608081016117e081614e47565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611c0657506020810177ffffffffffffffff0000000000000000000000000000000061184682614e68565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611b7457600091611bd7575b50611bad576118dd6118d860408401614e47565b6166d6565b67ffffffffffffffff6118ef82614e68565b16611907816000526006602052604060002054151590565b15611b8057602073ffffffffffffffffffffffffffffffffffffffff60045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa908115611b7457600091611b0a575b5073ffffffffffffffffffffffffffffffffffffffff163303611adc5761123a81611aab936119aa60606119a0611a3b96614e68565b92013580926159c0565b6119b381615bd6565b7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff6119e684614e68565b6040805173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152336020820152908101949094521691606090a2614e68565b610cb560405160ff7f000000000000000000000000000000000000000000000000000000000000000016602082015260208152611a79604082614888565b60405192611a86846147fc565b8352602083019081526040519384936020855251604060208601526060850190614926565b90517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0848303016040850152614926565b7f728fe07b000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b6020813d602011611b6c575b81611b2360209383614888565b81010312611b6857519073ffffffffffffffffffffffffffffffffffffffff82168203611b65575073ffffffffffffffffffffffffffffffffffffffff61196a565b80fd5b5080fd5b3d9150611b16565b6040513d6000823e3d90fd5b7fa9902c7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f53ad11d80000000000000000000000000000000000000000000000000000000060005260046000fd5b611bf9915060203d602011611bff575b611bf18183614888565b8101906153ba565b836118c4565b503d611be7565b611c2473ffffffffffffffffffffffffffffffffffffffff91614e47565b7f961c9a4f000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b346102d25760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d25760043567ffffffffffffffff81116102d257611ca1903690600401614ab2565b9060243567ffffffffffffffff81116102d257611cc2903690600401614d43565b9060443567ffffffffffffffff81116102d257611ce3903690600401614d43565b73ffffffffffffffffffffffffffffffffffffffff6009541633141580611da4575b610e9b57838614801590611d9a575b611d705760005b868110611d2457005b80611d6a611d38610a606001948b8b6151da565b611d438389896151ea565b611d64611d5c611d5486898b6151ea565b923690614d91565b913690614d91565b91615aa6565b01611d1b565b7f568efce20000000000000000000000000000000000000000000000000000000060005260046000fd5b5080861415611d14565b5073ffffffffffffffffffffffffffffffffffffffff60015416331415611d05565b611dcf36614a07565b505092909391604051611de1816147fc565b600081526000602082015260606080604051611dfc816147b1565b828152826020820152826040820152600083820152015267ffffffffffffffff8416916040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff000000000000000000000000000000008660801b16600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611b7457600091612540575b50611bad57611ec4336166d6565b611edb836000526006602052604060002054151590565b156125125782600052600a602052604060002093845481116124e15760018501549461ffff8660081c169663ffffffff612710611f188a866152c4565b04602087019081529760181c168061247d5750611f3760038301614f50565b945b7f000000000000000000000000000000000000000000000000000000000000000095604051611f6781614818565b858152602081019a8b526040810160ff891681528d36611f87918a614ae3565b90606083019182526040519c8d9360208501602090525160408501525161ffff1660608401525160ff1660808301525160a082016080905260c08201611fcc91614926565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018b52611ffc908b614888565b6020809b60405161200d8382614888565b60008082529c73ffffffffffffffffffffffffffffffffffffffff9361204160026040519a61203b8c6147b1565b01614f50565b895288015260408701521690816060860152608085015273ffffffffffffffffffffffffffffffffffffffff600454168b60405180927f20487ded00000000000000000000000000000000000000000000000000000000825281806120aa8a8a6004840161501e565b03915afa90811561247257908c949392918c9161243d575b5082526120cf86846159c0565b6120fb8630337f0000000000000000000000000000000000000000000000000000000000000000615a3e565b61210486615bd6565b8061222f575b5050916121609273ffffffffffffffffffffffffffffffffffffffff60045416906040518095819482937f96f4e9f90000000000000000000000000000000000000000000000000000000084526004840161501e565b039134905af19687156122235780976121ed575b5050916121e2869798926121c26121ae7fc185a458b75814b39eec5147bcb801e0ca4336f65fdcc99d90bd8405b54104539689519061515f565b936121ba368785614ae3565b90858b61531c565b96519360405194859485528b85015260606040850152606084019161519b565b0390a4604051908152f35b909196508782813d831161221c575b6122068183614888565b81010312611b65575051946121e26121ae612174565b503d6121fc565b604051903d90823e3d90fd5b909192506122408251303384615a3e565b7f000000000000000000000000000000000000000000000000000000000000000091519182158015612398575b156123145761216094928c949261230761230c936122db6040519485927f095ea7b3000000000000000000000000000000000000000000000000000000008b850152602484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101845283614888565b616363565b90928c61210a565b60848c604051907f08c379a00000000000000000000000000000000000000000000000000000000082526004820152603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527f20746f206e6f6e2d7a65726f20616c6c6f77616e6365000000000000000000006064820152fd5b506040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff821660248201528c81604481865afa908115612432578c916123fe575b501561226d565b90508c81818e3d821161242a575b6124168284614888565b508101031261242657518e6123f7565b8b80fd5b3d915061240c565b6040513d8e823e3d90fd5b85819394959692503d831161246b575b6124578183614888565b81010312610a3257908b939291518e6120c2565b503d61244d565b6040513d8d823e3d90fd5b6040519061248a826147fc565b81526020810160018152604051917f181dcf100000000000000000000000000000000000000000000000000000000060208401525160248301525115156044820152604481526124db606482614888565b94611f39565b837f58dd87c50000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b827fa9902c7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b612559915060203d602011611bff57611bf18183614888565b88611eb6565b346102d25760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d257602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346102d25760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d257602061161a67ffffffffffffffff6125f66149ad565b166000526006602052604060002054151590565b346102d25760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d25760443567ffffffffffffffff81166004356024358284036102d2576064359260843560ff811681036102d25760a4359073ffffffffffffffffffffffffffffffffffffffff8216928383036102d25780600052600a60205260ff600160406000200154166128c0575b506126c6604051846020820152602081526126be604082614888565b82888861531c565b84036128925783600052600c602052604060002073ffffffffffffffffffffffffffffffffffffffff604051916126fc836147fc565b5461270a60ff821684615012565b60081c166020820152519560038710156114e85760009661286657916127386127419798926127689461548d565b9687809361590f565b337f0000000000000000000000000000000000000000000000000000000000000000615a3e565b604051612774816147fc565b6001815260208101338152838752600c60205260408720915160038110156128395760ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008454169116178255517fffffffffffffffffffffff0000000000000000000000000000000000000000ff74ffffffffffffffffffffffffffffffffffffffff0083549260081b16911617905560405193845260208401527fd6f70fb263bfe7d01ec6802b3c07b6bd32579760fe9fcb4e248a036debb8cdf160403394a480f35b6024887f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b602487867fcee81443000000000000000000000000000000000000000000000000000000008252600452fd5b847f402c2fe80000000000000000000000000000000000000000000000000000000060005260045260246000fd5b80600052600b6020526128e86040600020339060019160005201602052604060002054151590565b6126a2577f6c46a9b5000000000000000000000000000000000000000000000000000000006000526004523360245260446000fd5b346102d25761292b36614cd4565b73ffffffffffffffffffffffffffffffffffffffff600454163303612ed15760a0813603126102d257604051612960816147b1565b81358152612970602083016149c4565b9160208201928352604081013567ffffffffffffffff81116102d2576129999036908301614b1a565b9160408101928352606082013567ffffffffffffffff81116102d2576129c29036908401614b1a565b916060820192835260808101359067ffffffffffffffff82116102d2570136601f820112156102d25780356129f681614c03565b91612a046040519384614888565b81835260208084019260061b820101903682116102d257602001915b818310612e99575050506080820152600093519167ffffffffffffffff83169151935190519081518201906020820192602081840312612e9557602081015167ffffffffffffffff8111610a36570191608090839003126109ff5760405190612a8882614818565b6020830151825260408301519361ffff85168503610a36576020830194855260608401519360ff85168503612e915760408401948552608081015167ffffffffffffffff8111610a325760209101019080601f83011215612e9157815191612aef836148c9565b90612afd6040519283614888565b8382526020820192602085830101116124265773ffffffffffffffffffffffffffffffffffffffff93836020612b339301614903565b8060608601525190519060208110612e61575b50169577ffffffffffffffff00000000000000000000000000000000604051917f2cbc26bb00000000000000000000000000000000000000000000000000000000835260801b16600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115612e56578991612e37575b50612e0f57612bee8186614e7d565b15612dd157509060ff612c25612710612c1e61ffff612c15612c479751868851169061548d565b975116876152c4565b048561515f565b9151166040519186602084015260208352612c41604084614888565b8761531c565b92838652600c6020526040862060405190612c61826147fc565b5490612c7060ff831682615012565b73ffffffffffffffffffffffffffffffffffffffff602082019260081c16825280516003811015612da457612d0f575050612cae82612cb39461590f565b61559a565b808352600c6020526040832060027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008254161790557fc43049430336d164648e213c98d9fbfdafe63ddac3a5e350ec1dd6f4fa555bf58380a380f35b9150919250516003811015612d7757600103612d4b579073ffffffffffffffffffffffffffffffffffffffff612d4692511661559a565b612cb3565b602485847fb196a44a000000000000000000000000000000000000000000000000000000008252600452fd5b6024867f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b6024897f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b611400906040519182917f24eb47e5000000000000000000000000000000000000000000000000000000008352602060048401526024830190614926565b6004887f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b612e50915060203d602011611bff57611bf18183614888565b89612bdf565b6040513d8b823e3d90fd5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b168a612b46565b8980fd5b8780fd5b6040833603126102d25760206040918251612eb3816147fc565b612ebc8661498c565b81528286013583820152815201920191612a20565b7fd7f73334000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b346102d25760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d2577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174602073ffffffffffffffffffffffffffffffffffffffff612f6e614969565b612f7661565c565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000006009541617600955604051908152a1005b346102d25760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d25760005473ffffffffffffffffffffffffffffffffffffffff81163303613069577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346102d25760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d25767ffffffffffffffff6130d36149ad565b606060a06040516130e381614850565b60008152600060208201526000604082015260008382015282608082015201521680600052600a602052604060002090600052600b60205261312860406000206160c4565b610cb56040519161313883614850565b8354835261320760018501546131d6602086019660ff83161515885263ffffffff604088019361ffff8160081c1685528160608a019160181c16815261ffff613196600361318860028701614f50565b9560808d0196875201614f50565b9560a08b019687526040519b8c9b60408d525160408d015251151560608c0152511660808a0152511660a08801525160c080880152610100870190614926565b90517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08683030160e0870152614926565b908382036020850152614c8a565b346102d25760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d257602073ffffffffffffffffffffffffffffffffffffffff60095416604051908152f35b346102d25760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d25761329e6149ad565b60243567ffffffffffffffff81116102d2576132be903690600401614c6f565b60443567ffffffffffffffff81116102d2576132e767ffffffffffffffff913690600401614c6f565b926132f061565c565b169081600052600b60205260406000209260005b8251811015613341578061333a73ffffffffffffffffffffffffffffffffffffffff61333260019487614eba565b511687615f17565b5001613304565b50909160005b8251811015613384578061337d73ffffffffffffffffffffffffffffffffffffffff61337560019487614eba565b511687615fe8565b5001613347565b507fccc0f2211c115acfa175a7923abdeb4b0a7c376d1b9e43c74973efe83d7d9e22613207846113b885604051938493604085526040850190614c8a565b346102d2576133d036614ba4565b6133db92919261565c565b67ffffffffffffffff82166133fd816000526006602052604060002054151590565b156134185750610a4292613412913691614ae3565b90615701565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346102d25761346d61347561345936614b38565b949161346693919361565c565b3691614c1b565b923691614c1b565b7f0000000000000000000000000000000000000000000000000000000000000000156135955760005b8251811015613511578073ffffffffffffffffffffffffffffffffffffffff6134c960019386614eba565b51166134d481616758565b6134e0575b500161349e565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a1846134d9565b5060005b8151811015610a42578073ffffffffffffffffffffffffffffffffffffffff61354060019385614eba565b5116801561358f5761355181615e9f565b61355e575b505b01613515565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a183613556565b50613558565b7f35f4a7b30000000000000000000000000000000000000000000000000000000060005260046000fd5b346102d25760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d2576135f66149ad565b60243567ffffffffffffffff81116102d25760209161361c61161a923690600401614b1a565b90614e7d565b346102d25760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d25760043567ffffffffffffffff81116102d257613671903690600401614ab2565b61367961565c565b60005b81811015610a425760008160051b840135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0185360301821215611b65577f1e10bdc4000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000061370a6080858901016156a7565b1614613c29575b61271061ffff6137256020858901016156e5565b1611613c015767ffffffffffffffff613742606084880101614e68565b168152600a602052604081209061375f86840160c0810190614df6565b9067ffffffffffffffff8211613bd457613789826137806002870154614efd565b600287016152d7565b8290601f8311600114613b0c576137d49291849183613b015750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b60028301555b6137e86020848801016156e5565b60018301908782547fffffffffffffffffffffffffffffffffffffffffffffffffff0000000000000062ffff0066ffffffff00000061384460408b60ff613830828a016156f4565b1515169760a082820101358d5501016156d4565b60181b169460081b169116171717905561386560e084880101848801614df6565b919067ffffffffffffffff8311613ad457613890836138876003870154614efd565b600387016152d7565b81601f84116001146139e457936003613909857f2a5e99c38151929c7be04f002e86e1441e1596df786d74cd6b18b0c6cb935fa29661ffff9660019b9a966139ba9a926139d95750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9101555b67ffffffffffffffff6139246060838b0101614e68565b8963ffffffff6139386020868401016156e5565b947fffffffff0000000000000000000000000000000000000000000000000000000061396a84830160c0810190614df6565b9a9061397a6080858801016156a7565b9b60a061398b6040878a01016156d4565b95613997818a016156f4565b986040519d8e9d168d5201013560208b015260c060408b015260c08a019161519b565b99166060870152166080850152151560a084015216930390a20161367c565b013590508d80610782565b9160038501835260208320925b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085168110613abc575093600360018561ffff95829a99956139ba99837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe07f2a5e99c38151929c7be04f002e86e1441e1596df786d74cd6b18b0c6cb935fa29b1610613a86575b505050811b0191015561390d565b01357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83861b60f8161c191690558c8080613a78565b909260206001819286860135815501940191016139f1565b6024827f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b013590508980610782565b600285939293018452602084209084935b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe084168510613bbc5760019450837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0811610613b84575b505050811b0160028301556137da565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c19910135169055888080613b74565b81810135835560209485019460019093019201613b1d565b6024837f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b807f382c09820000000000000000000000000000000000000000000000000000000060049252fd5b63ffffffff613c3c6040848801016156d4565b161561371157807f382c09820000000000000000000000000000000000000000000000000000000060049252fd5b346102d25760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d25760043567ffffffffffffffff81116102d25780600401906101007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126102d2576000604051613ceb81614834565b5260848101613cf981614e47565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611c0657506024810177ffffffffffffffff00000000000000000000000000000000613d5f82614e68565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611b7457600091613ffc575b50611bad5767ffffffffffffffff613df482614e68565b16613e0c816000526006602052604060002054151590565b15611b8057602073ffffffffffffffffffffffffffffffffffffffff60045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115611b7457600091613fdd575b5015611adc57613e8481614e68565b613ea060a484019161361c613e998488614df6565b3691614ae3565b15613f9657507ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0608067ffffffffffffffff613f2d613f276044613f1788613f11613f0c613e9960209d60c4613ef58e614e68565b95613f056064820135809861590f565b0190614df6565b6153d2565b9061548d565b970195610a6088612cae89614e47565b94614e47565b9373ffffffffffffffffffffffffffffffffffffffff60405195817f000000000000000000000000000000000000000000000000000000000000000016875233898801521660408601528560608601521692a280604051613f8d81614834565b52604051908152f35b613fa09084614df6565b6114006040519283927f24eb47e500000000000000000000000000000000000000000000000000000000845260206004850152602484019161519b565b613ff6915060203d602011611bff57611bf18183614888565b84613e75565b614015915060203d602011611bff57611bf18183614888565b84613ddd565b346102d25761402936614a07565b505092909193614037614ddd565b5060405194614045866147fc565b600086526000602087015260606080604051614060816147b1565b828152826020820152826040820152600083820152015267ffffffffffffffff8216926040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff000000000000000000000000000000008460801b16600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611b7457600091614424575b50611bad57614128336166d6565b61413f846000526006602052604060002054151590565b156143f65783600052600a602052604060002093845482116143c6575091839161426261430597969461423660018b9801549663ffffffff61ffff8960081c1698602061271061418f8c886152c4565b049b019a8b5260181c168061435c575060ff6141f96141b060038a01614f50565b965b604051956141bf87614818565b8652602086019a8b526040860193837f00000000000000000000000000000000000000000000000000000000000000001685523691614ae3565b916060850192835261ffff6040519a8b96602080890152516040880152511660608601525116608084015251608060a084015260c0830190614926565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101865285614888565b73ffffffffffffffffffffffffffffffffffffffff60209687956040519061428a8883614888565b600082526142a060026040519861203b8a6147b1565b8752878701526040860152166060840152608083015273ffffffffffffffffffffffffffffffffffffffff60045416906040518097819482937f20487ded0000000000000000000000000000000000000000000000000000000084526004840161501e565b03915afa928315611b745760009361432a575b50826040945283519283525190820152f35b9392508184813d8311614355575b6143428183614888565b810103126102d257604093519293614318565b503d614338565b6141f960ff916040519061436f826147fc565b81526020810160018152604051917f181dcf100000000000000000000000000000000000000000000000000000000060208401525160248301525115156044820152604481526143c0606482614888565b966141b2565b7f58dd87c50000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b837fa9902c7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b61443d915060203d602011611bff57611bf18183614888565b8861411a565b346102d25760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d257602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102d25760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d25760206144d8614969565b73ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b346102d25760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d257602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102d25760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d257610cb56040516145cc606082614888565b602781527f4275726e4d696e74466173745472616e73666572546f6b656e506f6f6c20312e60208201527f362e312d646576000000000000000000000000000000000000000000000000006040820152604051918291602083526020830190614926565b346102d25760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d257600435907fffffffff0000000000000000000000000000000000000000000000000000000082168092036102d257817fd53652570000000000000000000000000000000000000000000000000000000060209314908115614726575b81156146c9575b5015158152f35b7f85572ffb000000000000000000000000000000000000000000000000000000008114915081156146fc575b50836146c2565b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014836146f5565b90507faff2afbf0000000000000000000000000000000000000000000000000000000081148015614788575b801561475f575b906146bb565b507f01ffc9a7000000000000000000000000000000000000000000000000000000008114614759565b507f0e64dd29000000000000000000000000000000000000000000000000000000008114614752565b60a0810190811067ffffffffffffffff8211176147cd57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff8211176147cd57604052565b6080810190811067ffffffffffffffff8211176147cd57604052565b6020810190811067ffffffffffffffff8211176147cd57604052565b60c0810190811067ffffffffffffffff8211176147cd57604052565b6060810190811067ffffffffffffffff8211176147cd57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176147cd57604052565b67ffffffffffffffff81116147cd57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60005b8381106149165750506000910152565b8181015183820152602001614906565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f60209361496281518092818752878088019101614903565b0116010190565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036102d257565b359073ffffffffffffffffffffffffffffffffffffffff821682036102d257565b6004359067ffffffffffffffff821682036102d257565b359067ffffffffffffffff821682036102d257565b9181601f840112156102d25782359167ffffffffffffffff83116102d257602083818601950101116102d257565b9060a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126102d25760043567ffffffffffffffff811681036102d257916024359160443567ffffffffffffffff81116102d25782614a6a916004016149d9565b9290929160643573ffffffffffffffffffffffffffffffffffffffff811681036102d257916084359067ffffffffffffffff82116102d257614aae916004016149d9565b9091565b9181601f840112156102d25782359167ffffffffffffffff83116102d2576020808501948460051b0101116102d257565b929192614aef826148c9565b91614afd6040519384614888565b8294818452818301116102d2578281602093846000960137010152565b9080601f830112156102d257816020614b3593359101614ae3565b90565b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126102d25760043567ffffffffffffffff81116102d25781614b8191600401614ab2565b929092916024359067ffffffffffffffff82116102d257614aae91600401614ab2565b9060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126102d25760043567ffffffffffffffff811681036102d257916024359067ffffffffffffffff82116102d257614aae916004016149d9565b67ffffffffffffffff81116147cd5760051b60200190565b929190614c2781614c03565b93614c356040519586614888565b602085838152019160051b81019283116102d257905b828210614c5757505050565b60208091614c648461498c565b815201910190614c4b565b9080601f830112156102d257816020614b3593359101614c1b565b906020808351928381520192019060005b818110614ca85750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101614c9b565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126102d2576004359067ffffffffffffffff82116102d2577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8260a0920301126102d25760040190565b9181601f840112156102d25782359167ffffffffffffffff83116102d257602080850194606085020101116102d257565b35906fffffffffffffffffffffffffffffffff821682036102d257565b91908260609103126102d257604051614da98161486c565b809280359081151582036102d2576040614dd89181938552614dcd60208201614d74565b602086015201614d74565b910152565b60405190614dea826147fc565b60006020838281520152565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156102d2570180359067ffffffffffffffff82116102d2576020019181360383136102d257565b3573ffffffffffffffffffffffffffffffffffffffff811681036102d25790565b3567ffffffffffffffff811681036102d25790565b9067ffffffffffffffff614b3592166000526007602052600560406000200190602081519101209060019160005201602052604060002054151590565b8051821015614ece5760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600182811c92168015614f46575b6020831014614f1757565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691614f0c565b9060405191826000825492614f6484614efd565b8084529360018116908115614fd25750600114614f8b575b50614f8992500383614888565b565b90506000929192526020600020906000915b818310614fb6575050906020614f899282010138614f7c565b6020919350806001915483858901015201910190918492614f9d565b60209350614f899592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138614f7c565b60038210156114e85752565b9067ffffffffffffffff909392931681526040602082015261508361504f845160a0604085015260e0840190614926565b60208501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0848303016060850152614926565b906040840151917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08282030160808301526020808451928381520193019060005b8181106151275750505060808473ffffffffffffffffffffffffffffffffffffffff6060614b35969701511660a084015201519060c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc082850301910152614926565b8251805173ffffffffffffffffffffffffffffffffffffffff16865260209081015181870152604090950194909201916001016150c4565b9190820391821161516c57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b9190811015614ece5760051b0190565b9190811015614ece576060020190565b60405190615207826147b1565b60006080838281528260208201528260408201528260608201520152565b90604051615232816147b1565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff166000526007602052614b356004604060002001614f50565b8181106152b8575050565b600081556001016152ad565b8181029291811591840414171561516c57565b9190601f81116152e657505050565b614f89926000526020600020906020601f840160051c83019310615312575b601f0160051c01906152ad565b9091508190615305565b929061534f61537b9260ff60405195869460208601988952604086015216606084015260808084015260a0830190614926565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282614888565b51902090565b811561538b570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b908160209103126102d2575180151581036102d25790565b80518015615442576020036154045780516020828101918301839003126102d257519060ff8211615404575060ff1690565b611400906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190614926565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff821161516c57565b60ff16604d811161516c57600a0a90565b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff8116928284146155935782841161556957906154d291615468565b91604d60ff8416118015615530575b6154fa575050906154f4614b359261547c565b906152c4565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b5061553a8361547c565b801561538b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0484116154e1565b61557291615468565b91604d60ff8416116154fa5750509061558d614b359261547c565b90615381565b5050505090565b9073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803b156102d2576040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff93909316600484015260248301919091526000919082908290604490829084905af1801561565157615644575050565b8161564e91614888565b50565b6040513d84823e3d90fd5b73ffffffffffffffffffffffffffffffffffffffff60015416330361567d57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b357fffffffff00000000000000000000000000000000000000000000000000000000811681036102d25790565b3563ffffffff811681036102d25790565b3561ffff811681036102d25790565b3580151581036102d25790565b908051156111d35767ffffffffffffffff81516020830120921691826000526007602052615736816005604060002001615f17565b156158cb5760005260086020526040600020815167ffffffffffffffff81116147cd5761576d816157678454614efd565b846152d7565b6020601f821160011461580557916157df827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea95936157f5956000916157fa575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055604051918291602083526020830190614926565b0390a2565b9050840151386157ae565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b8181106158b35750926157f59492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea98961061587c575b5050811b01905561123f565b8501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880615870565b9192602060018192868a015181550194019201615835565b50906114006040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190614926565b67ffffffffffffffff7f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c91169182600052600760205280615990600260406000200173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692839161611c565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252602082019290925290819081016157f5565b67ffffffffffffffff7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894491169182600052600760205280615990604060002073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692839161611c565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff92831660248201529290911660448301526064820192909252614f899161230782608481016122db565b67ffffffffffffffff166000818152600660205260409020549092919015615ba85791615ba560e092615b7185615afd7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b97615ce0565b846000526007602052615b148160406000206164a3565b615b1d83615ce0565b846000526007602052615b378360026040600020016164a3565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690813b156102d257604051907f42966c680000000000000000000000000000000000000000000000000000000082528160248160008096819560048401525af1801561565157615644575050565b615c636151fa565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff8083511691615cc06020850193615cba615cad63ffffffff8751164261515f565b85608089015116906152c4565b9061610f565b80821015615cd957505b16825263ffffffff4216905290565b9050615cca565b805115615d80576fffffffffffffffffffffffffffffffff6040820151166fffffffffffffffffffffffffffffffff60208301511610615d1d5750565b606490615d7e604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff60408201511615801590615e08575b615da75750565b606490615d7e604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff6020820151161515615da0565b8054821015614ece5760005260206000200190600090565b805490680100000000000000008210156147cd5781615e66916001615e9b94018155615e27565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055565b80600052600360205260406000205415600014615ed857615ec1816002615e3f565b600254906000526003602052604060002055600190565b50600090565b80600052600660205260406000205415600014615ed857615f00816005615e3f565b600554906000526006602052604060002055600190565b6000828152600182016020526040902054615f4e5780615f3983600193615e3f565b80549260005201602052604060002055600190565b5050600090565b80548015615fb9577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190615f8a8282615e27565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b9060018201918160005282602052604060002054908115156000146160bb577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019180831161516c5781547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810190811161516c5783616072948203616084575b505050615f55565b60005260205260006040812055600190565b6160a4616094615e669386615e27565b90549060031b1c92839286615e27565b90556000528460205260406000205538808061606a565b50505050600090565b906040519182815491828252602082019060005260206000209260005b8181106160f6575050614f8992500383614888565b84548352600194850194879450602090930192016160e1565b9190820180921161516c57565b9182549060ff8260a01c1615801561635b575b616355576fffffffffffffffffffffffffffffffff8216916001850190815461617463ffffffff6fffffffffffffffffffffffffffffffff83169360801c164261515f565b90816162b7575b505084811061626b57508383106161d55750506161aa6fffffffffffffffffffffffffffffffff92839261515f565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b5460801c916161e4818561515f565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81019080821161516c576162326162379273ffffffffffffffffffffffffffffffffffffffff9661610f565b615381565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828573ffffffffffffffffffffffffffffffffffffffff927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b82869293961161632b576162d292615cba9160801c906152c4565b808410156163265750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff000000000000000000000000000000001617865592388061617b565b6162dd565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b50821561612f565b73ffffffffffffffffffffffffffffffffffffffff6163f29116916040926000808551936163918786614888565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602086015260208151910182855af13d1561649b573d916163d6836148c9565b926163e387519485614888565b83523d6000602085013e616981565b805190816163ff57505050565b6020806164109383010191016153ba565b156164185750565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606091616981565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19916165dc60609280546164e063ffffffff8260801c164261515f565b908161661b575b50506fffffffffffffffffffffffffffffffff600181602086015116928281541680851060001461661357508280855b16167fffffffffffffffffffffffffffffffff000000000000000000000000000000008254161781556165908651151582907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b60408601517fffffffffffffffffffffffffffffffff0000000000000000000000000000000060809190911b16939092166fffffffffffffffffffffffffffffffff1692909217910155565b615ba560405180926fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b838091616517565b6fffffffffffffffffffffffffffffffff916166508392836166496001880154948286169560801c906152c4565b911661610f565b808210156166cf57505b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff9290911692909216167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116174260801b73ffffffff000000000000000000000000000000001617815538806164e7565b905061665a565b7f00000000000000000000000000000000000000000000000000000000000000006166fe5750565b73ffffffffffffffffffffffffffffffffffffffff168060005260036020526040600020541561672b5750565b7fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6000818152600360205260409020548015615f4e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161516c57600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161516c578181036167ee575b5050506167da6002615f55565b600052600360205260006040812055600190565b6168106167ff615e66936002615e27565b90549060031b1c9283926002615e27565b905560005260036020526040600020553880806167cd565b6000818152600660205260409020548015615f4e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161516c57600554907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161516c578181036168be575b5050506168aa6005615f55565b600052600660205260006040812055600190565b6168e06168cf615e66936005615e27565b90549060031b1c9283926005615e27565b9055600052600660205260406000205538808061689d565b9060018201918160005282602052604060002054908115156000146160bb577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019180831161516c5781547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810190811161516c578381616072950361608457505050615f55565b919290156169fc5750815115616995575090565b3b1561699e5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015616a0f5750805190602001fd5b611400906040519182917f08c379a000000000000000000000000000000000000000000000000000000000835260206004840152602483019061492656fea164736f6c634300081a000a",
}

var BurnMintFastTransferTokenPoolABI = BurnMintFastTransferTokenPoolMetaData.ABI

var BurnMintFastTransferTokenPoolBin = BurnMintFastTransferTokenPoolMetaData.Bin

func DeployBurnMintFastTransferTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, localTokenDecimals uint8, allowlist []common.Address, rmnProxy common.Address, router common.Address) (common.Address, *types.Transaction, *BurnMintFastTransferTokenPool, error) {
	parsed, err := BurnMintFastTransferTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BurnMintFastTransferTokenPoolBin), backend, token, localTokenDecimals, allowlist, rmnProxy, router)
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) ComputeFillId(opts *bind.CallOpts, settlementId [32]byte, sourceAmountNetFee *big.Int, sourceDecimals uint8, receiver []byte) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "computeFillId", settlementId, sourceAmountNetFee, sourceDecimals, receiver)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) ComputeFillId(settlementId [32]byte, sourceAmountNetFee *big.Int, sourceDecimals uint8, receiver []byte) ([32]byte, error) {
	return _BurnMintFastTransferTokenPool.Contract.ComputeFillId(&_BurnMintFastTransferTokenPool.CallOpts, settlementId, sourceAmountNetFee, sourceDecimals, receiver)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) ComputeFillId(settlementId [32]byte, sourceAmountNetFee *big.Int, sourceDecimals uint8, receiver []byte) ([32]byte, error) {
	return _BurnMintFastTransferTokenPool.Contract.ComputeFillId(&_BurnMintFastTransferTokenPool.CallOpts, settlementId, sourceAmountNetFee, sourceDecimals, receiver)
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) GetAllowedFillers(opts *bind.CallOpts, remoteChainSelector uint64) ([]common.Address, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "getAllowedFillers", remoteChainSelector)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) GetAllowedFillers(remoteChainSelector uint64) ([]common.Address, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetAllowedFillers(&_BurnMintFastTransferTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) GetAllowedFillers(remoteChainSelector uint64) ([]common.Address, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetAllowedFillers(&_BurnMintFastTransferTokenPool.CallOpts, remoteChainSelector)
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) IsAllowedFiller(opts *bind.CallOpts, remoteChainSelector uint64, filler common.Address) (bool, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "isAllowedFiller", remoteChainSelector, filler)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) IsAllowedFiller(remoteChainSelector uint64, filler common.Address) (bool, error) {
	return _BurnMintFastTransferTokenPool.Contract.IsAllowedFiller(&_BurnMintFastTransferTokenPool.CallOpts, remoteChainSelector, filler)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) IsAllowedFiller(remoteChainSelector uint64, filler common.Address) (bool, error) {
	return _BurnMintFastTransferTokenPool.Contract.IsAllowedFiller(&_BurnMintFastTransferTokenPool.CallOpts, remoteChainSelector, filler)
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactor) CcipSendToken(opts *bind.TransactOpts, destinationChainSelector uint64, amount *big.Int, receiver []byte, feeToken common.Address, extraArgs []byte) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.contract.Transact(opts, "ccipSendToken", destinationChainSelector, amount, receiver, feeToken, extraArgs)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) CcipSendToken(destinationChainSelector uint64, amount *big.Int, receiver []byte, feeToken common.Address, extraArgs []byte) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.CcipSendToken(&_BurnMintFastTransferTokenPool.TransactOpts, destinationChainSelector, amount, receiver, feeToken, extraArgs)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorSession) CcipSendToken(destinationChainSelector uint64, amount *big.Int, receiver []byte, feeToken common.Address, extraArgs []byte) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.CcipSendToken(&_BurnMintFastTransferTokenPool.TransactOpts, destinationChainSelector, amount, receiver, feeToken, extraArgs)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactor) FastFill(opts *bind.TransactOpts, settlementId [32]byte, fillId [32]byte, sourceChainSelector uint64, sourceAmountNetFee *big.Int, sourceDecimals uint8, receiver common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.contract.Transact(opts, "fastFill", settlementId, fillId, sourceChainSelector, sourceAmountNetFee, sourceDecimals, receiver)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) FastFill(settlementId [32]byte, fillId [32]byte, sourceChainSelector uint64, sourceAmountNetFee *big.Int, sourceDecimals uint8, receiver common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.FastFill(&_BurnMintFastTransferTokenPool.TransactOpts, settlementId, fillId, sourceChainSelector, sourceAmountNetFee, sourceDecimals, receiver)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorSession) FastFill(settlementId [32]byte, fillId [32]byte, sourceChainSelector uint64, sourceAmountNetFee *big.Int, sourceDecimals uint8, receiver common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.FastFill(&_BurnMintFastTransferTokenPool.TransactOpts, settlementId, fillId, sourceChainSelector, sourceAmountNetFee, sourceDecimals, receiver)
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactor) UpdateFillerAllowList(opts *bind.TransactOpts, destinationChainSelector uint64, fillersToAdd []common.Address, fillersToRemove []common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.contract.Transact(opts, "updateFillerAllowList", destinationChainSelector, fillersToAdd, fillersToRemove)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) UpdateFillerAllowList(destinationChainSelector uint64, fillersToAdd []common.Address, fillersToRemove []common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.UpdateFillerAllowList(&_BurnMintFastTransferTokenPool.TransactOpts, destinationChainSelector, fillersToAdd, fillersToRemove)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorSession) UpdateFillerAllowList(destinationChainSelector uint64, fillersToAdd []common.Address, fillersToRemove []common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.UpdateFillerAllowList(&_BurnMintFastTransferTokenPool.TransactOpts, destinationChainSelector, fillersToAdd, fillersToRemove)
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
	FastTransferBpsFee       uint16
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
	FastTransferFee          *big.Int
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
	FillId       [32]byte
	SettlementId [32]byte
	Raw          types.Log
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
	DestChainSelector uint64
	AddFillers        []common.Address
	RemoveFillers     []common.Address
	Raw               types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterFillerAllowListUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnMintFastTransferTokenPoolFillerAllowListUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "FillerAllowListUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolFillerAllowListUpdatedIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "FillerAllowListUpdated", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchFillerAllowListUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolFillerAllowListUpdated, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "FillerAllowListUpdated", destChainSelectorRule)
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPool) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _BurnMintFastTransferTokenPool.abi.Events["AllowListAdd"].ID:
		return _BurnMintFastTransferTokenPool.ParseAllowListAdd(log)
	case _BurnMintFastTransferTokenPool.abi.Events["AllowListRemove"].ID:
		return _BurnMintFastTransferTokenPool.ParseAllowListRemove(log)
	case _BurnMintFastTransferTokenPool.abi.Events["ChainAdded"].ID:
		return _BurnMintFastTransferTokenPool.ParseChainAdded(log)
	case _BurnMintFastTransferTokenPool.abi.Events["ChainConfigured"].ID:
		return _BurnMintFastTransferTokenPool.ParseChainConfigured(log)
	case _BurnMintFastTransferTokenPool.abi.Events["ChainRemoved"].ID:
		return _BurnMintFastTransferTokenPool.ParseChainRemoved(log)
	case _BurnMintFastTransferTokenPool.abi.Events["ConfigChanged"].ID:
		return _BurnMintFastTransferTokenPool.ParseConfigChanged(log)
	case _BurnMintFastTransferTokenPool.abi.Events["DestChainConfigUpdated"].ID:
		return _BurnMintFastTransferTokenPool.ParseDestChainConfigUpdated(log)
	case _BurnMintFastTransferTokenPool.abi.Events["DestinationPoolUpdated"].ID:
		return _BurnMintFastTransferTokenPool.ParseDestinationPoolUpdated(log)
	case _BurnMintFastTransferTokenPool.abi.Events["FastTransferFilled"].ID:
		return _BurnMintFastTransferTokenPool.ParseFastTransferFilled(log)
	case _BurnMintFastTransferTokenPool.abi.Events["FastTransferRequested"].ID:
		return _BurnMintFastTransferTokenPool.ParseFastTransferRequested(log)
	case _BurnMintFastTransferTokenPool.abi.Events["FastTransferSettled"].ID:
		return _BurnMintFastTransferTokenPool.ParseFastTransferSettled(log)
	case _BurnMintFastTransferTokenPool.abi.Events["FillerAllowListUpdated"].ID:
		return _BurnMintFastTransferTokenPool.ParseFillerAllowListUpdated(log)
	case _BurnMintFastTransferTokenPool.abi.Events["InboundRateLimitConsumed"].ID:
		return _BurnMintFastTransferTokenPool.ParseInboundRateLimitConsumed(log)
	case _BurnMintFastTransferTokenPool.abi.Events["LockedOrBurned"].ID:
		return _BurnMintFastTransferTokenPool.ParseLockedOrBurned(log)
	case _BurnMintFastTransferTokenPool.abi.Events["OutboundRateLimitConsumed"].ID:
		return _BurnMintFastTransferTokenPool.ParseOutboundRateLimitConsumed(log)
	case _BurnMintFastTransferTokenPool.abi.Events["OwnershipTransferRequested"].ID:
		return _BurnMintFastTransferTokenPool.ParseOwnershipTransferRequested(log)
	case _BurnMintFastTransferTokenPool.abi.Events["OwnershipTransferred"].ID:
		return _BurnMintFastTransferTokenPool.ParseOwnershipTransferred(log)
	case _BurnMintFastTransferTokenPool.abi.Events["RateLimitAdminSet"].ID:
		return _BurnMintFastTransferTokenPool.ParseRateLimitAdminSet(log)
	case _BurnMintFastTransferTokenPool.abi.Events["ReleasedOrMinted"].ID:
		return _BurnMintFastTransferTokenPool.ParseReleasedOrMinted(log)
	case _BurnMintFastTransferTokenPool.abi.Events["RemotePoolAdded"].ID:
		return _BurnMintFastTransferTokenPool.ParseRemotePoolAdded(log)
	case _BurnMintFastTransferTokenPool.abi.Events["RemotePoolRemoved"].ID:
		return _BurnMintFastTransferTokenPool.ParseRemotePoolRemoved(log)
	case _BurnMintFastTransferTokenPool.abi.Events["RouterUpdated"].ID:
		return _BurnMintFastTransferTokenPool.ParseRouterUpdated(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
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
	return common.HexToHash("0x2a5e99c38151929c7be04f002e86e1441e1596df786d74cd6b18b0c6cb935fa2")
}

func (BurnMintFastTransferTokenPoolDestinationPoolUpdated) Topic() common.Hash {
	return common.HexToHash("0xb760e03fa04c0e86fcff6d0046cdcf22fb5d5b6a17d1e6f890b3456e81c40fd8")
}

func (BurnMintFastTransferTokenPoolFastTransferFilled) Topic() common.Hash {
	return common.HexToHash("0xd6f70fb263bfe7d01ec6802b3c07b6bd32579760fe9fcb4e248a036debb8cdf1")
}

func (BurnMintFastTransferTokenPoolFastTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xc185a458b75814b39eec5147bcb801e0ca4336f65fdcc99d90bd8405b5410453")
}

func (BurnMintFastTransferTokenPoolFastTransferSettled) Topic() common.Hash {
	return common.HexToHash("0xc43049430336d164648e213c98d9fbfdafe63ddac3a5e350ec1dd6f4fa555bf5")
}

func (BurnMintFastTransferTokenPoolFillerAllowListUpdated) Topic() common.Hash {
	return common.HexToHash("0xccc0f2211c115acfa175a7923abdeb4b0a7c376d1b9e43c74973efe83d7d9e22")
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
	ComputeFillId(opts *bind.CallOpts, settlementId [32]byte, sourceAmountNetFee *big.Int, sourceDecimals uint8, receiver []byte) ([32]byte, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetAllowedFillers(opts *bind.CallOpts, remoteChainSelector uint64) ([]common.Address, error)

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

	IsAllowedFiller(opts *bind.CallOpts, remoteChainSelector uint64, filler common.Address) (bool, error)

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

	CcipSendToken(opts *bind.TransactOpts, destinationChainSelector uint64, amount *big.Int, receiver []byte, feeToken common.Address, extraArgs []byte) (*types.Transaction, error)

	FastFill(opts *bind.TransactOpts, settlementId [32]byte, fillId [32]byte, sourceChainSelector uint64, sourceAmountNetFee *big.Int, sourceDecimals uint8, receiver common.Address) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error)

	SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error)

	SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateDestChainConfig(opts *bind.TransactOpts, destChainConfigArgs []FastTransferTokenPoolAbstractDestChainConfigUpdateArgs) (*types.Transaction, error)

	UpdateFillerAllowList(opts *bind.TransactOpts, destinationChainSelector uint64, fillersToAdd []common.Address, fillersToRemove []common.Address) (*types.Transaction, error)

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

	FilterFillerAllowListUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnMintFastTransferTokenPoolFillerAllowListUpdatedIterator, error)

	WatchFillerAllowListUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolFillerAllowListUpdated, destChainSelector []uint64) (event.Subscription, error)

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

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
