// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fast_transfer_token_pool_abstract

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

type FastTransferTokenPoolAbstractDestChainConfigUpdateArgs struct {
	MaxFillAmountPerRequest *big.Int
	AddFillers              []common.Address
	RemoveFillers           []common.Address
	RemoteChainSelector     uint64
	FastTransferBpsFee      uint16
	FillerAllowlistEnabled  bool
	DestinationPool         []byte
}

type FastTransferTokenPoolAbstractDestChainConfigView struct {
	MaxFillAmountPerRequest *big.Int
	FastTransferBpsFee      uint16
	FillerAllowlistEnabled  bool
	DestinationPool         []byte
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

var FastTransferTokenPoolAbstractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipReceive\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.Any2EVMMessage\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipSendToken\",\"inputs\":[{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"fillRequestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"computeFillId\",\"inputs\":[{\"name\":\"fillRequestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"fastFill\",\"inputs\":[{\"name\":\"fillRequestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"srcAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"srcDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCcipSendTokenFee\",\"inputs\":[{\"name\":\"settlementFeeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIFastTransferPool.Quote\",\"components\":[{\"name\":\"ccipSettlementFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fastTransferFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFastTransferTokenPoolAbstract.DestChainConfigView\",\"components\":[{\"name\":\"maxFillAmountPerRequest\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fastTransferBpsFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fillerAllowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"destinationPool\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFillInfo\",\"inputs\":[{\"name\":\"fillId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFastTransferTokenPoolAbstract.FillInfo\",\"components\":[{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"enumFastTransferTokenPoolAbstract.FillState\"},{\"name\":\"filler\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isfillerAllowListed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"filler\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"lockOrBurnOut\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRouter\",\"inputs\":[{\"name\":\"newRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"updateDestChainConfig\",\"inputs\":[{\"name\":\"laneConfigArgs\",\"type\":\"tuple\",\"internalType\":\"structFastTransferTokenPoolAbstract.DestChainConfigUpdateArgs\",\"components\":[{\"name\":\"maxFillAmountPerRequest\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"addFillers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removeFillers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fastTransferBpsFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fillerAllowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"destinationPool\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updatefillerAllowList\",\"inputs\":[{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addFillers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removeFillers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Burned\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestinationPoolUpdated\",\"inputs\":[{\"name\":\"dst\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"destinationPool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FastFill\",\"inputs\":[{\"name\":\"fillRequestId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"fillId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"filler\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"destAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FastFillRequest\",\"inputs\":[{\"name\":\"fillRequestId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"dstChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fastTransferFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FastFillSettled\",\"inputs\":[{\"name\":\"fillRequestId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FillerAllowListUpdated\",\"inputs\":[{\"name\":\"dst\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"addFillers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"removeFillers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvalidFill\",\"inputs\":[{\"name\":\"fillRequestId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"filler\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"filledAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"expectedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LaneUpdated\",\"inputs\":[{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"bps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"maxFillAmountPerRequest\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"destinationPool\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"addFillers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"removeFillers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Locked\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Minted\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Released\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RouterUpdated\",\"inputs\":[{\"name\":\"oldRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensConsumed\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AggregateValueMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"AggregateValueRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyFilled\",\"inputs\":[{\"name\":\"fillRequestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AlreadySettled\",\"inputs\":[{\"name\":\"fillRequestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"FillerNotWhitelisted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"filler\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRouter\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"LaneDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"RateLimitMustBeDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"WhitelistNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
}

var FastTransferTokenPoolAbstractABI = FastTransferTokenPoolAbstractMetaData.ABI

type FastTransferTokenPoolAbstract struct {
	address common.Address
	abi     abi.ABI
	FastTransferTokenPoolAbstractCaller
	FastTransferTokenPoolAbstractTransactor
	FastTransferTokenPoolAbstractFilterer
}

type FastTransferTokenPoolAbstractCaller struct {
	contract *bind.BoundContract
}

type FastTransferTokenPoolAbstractTransactor struct {
	contract *bind.BoundContract
}

type FastTransferTokenPoolAbstractFilterer struct {
	contract *bind.BoundContract
}

type FastTransferTokenPoolAbstractSession struct {
	Contract     *FastTransferTokenPoolAbstract
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type FastTransferTokenPoolAbstractCallerSession struct {
	Contract *FastTransferTokenPoolAbstractCaller
	CallOpts bind.CallOpts
}

type FastTransferTokenPoolAbstractTransactorSession struct {
	Contract     *FastTransferTokenPoolAbstractTransactor
	TransactOpts bind.TransactOpts
}

type FastTransferTokenPoolAbstractRaw struct {
	Contract *FastTransferTokenPoolAbstract
}

type FastTransferTokenPoolAbstractCallerRaw struct {
	Contract *FastTransferTokenPoolAbstractCaller
}

type FastTransferTokenPoolAbstractTransactorRaw struct {
	Contract *FastTransferTokenPoolAbstractTransactor
}

func NewFastTransferTokenPoolAbstract(address common.Address, backend bind.ContractBackend) (*FastTransferTokenPoolAbstract, error) {
	abi, err := abi.JSON(strings.NewReader(FastTransferTokenPoolAbstractABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindFastTransferTokenPoolAbstract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FastTransferTokenPoolAbstract{address: address, abi: abi, FastTransferTokenPoolAbstractCaller: FastTransferTokenPoolAbstractCaller{contract: contract}, FastTransferTokenPoolAbstractTransactor: FastTransferTokenPoolAbstractTransactor{contract: contract}, FastTransferTokenPoolAbstractFilterer: FastTransferTokenPoolAbstractFilterer{contract: contract}}, nil
}

func NewFastTransferTokenPoolAbstractCaller(address common.Address, caller bind.ContractCaller) (*FastTransferTokenPoolAbstractCaller, error) {
	contract, err := bindFastTransferTokenPoolAbstract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FastTransferTokenPoolAbstractCaller{contract: contract}, nil
}

func NewFastTransferTokenPoolAbstractTransactor(address common.Address, transactor bind.ContractTransactor) (*FastTransferTokenPoolAbstractTransactor, error) {
	contract, err := bindFastTransferTokenPoolAbstract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FastTransferTokenPoolAbstractTransactor{contract: contract}, nil
}

func NewFastTransferTokenPoolAbstractFilterer(address common.Address, filterer bind.ContractFilterer) (*FastTransferTokenPoolAbstractFilterer, error) {
	contract, err := bindFastTransferTokenPoolAbstract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FastTransferTokenPoolAbstractFilterer{contract: contract}, nil
}

func bindFastTransferTokenPoolAbstract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FastTransferTokenPoolAbstractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FastTransferTokenPoolAbstract.Contract.FastTransferTokenPoolAbstractCaller.contract.Call(opts, result, method, params...)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.FastTransferTokenPoolAbstractTransactor.contract.Transfer(opts)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.FastTransferTokenPoolAbstractTransactor.contract.Transact(opts, method, params...)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FastTransferTokenPoolAbstract.Contract.contract.Call(opts, result, method, params...)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.contract.Transfer(opts)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.contract.Transact(opts, method, params...)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCaller) ComputeFillId(opts *bind.CallOpts, fillRequestId [32]byte, amount *big.Int, receiver common.Address) ([32]byte, error) {
	var out []interface{}
	err := _FastTransferTokenPoolAbstract.contract.Call(opts, &out, "computeFillId", fillRequestId, amount, receiver)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) ComputeFillId(fillRequestId [32]byte, amount *big.Int, receiver common.Address) ([32]byte, error) {
	return _FastTransferTokenPoolAbstract.Contract.ComputeFillId(&_FastTransferTokenPoolAbstract.CallOpts, fillRequestId, amount, receiver)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCallerSession) ComputeFillId(fillRequestId [32]byte, amount *big.Int, receiver common.Address) ([32]byte, error) {
	return _FastTransferTokenPoolAbstract.Contract.ComputeFillId(&_FastTransferTokenPoolAbstract.CallOpts, fillRequestId, amount, receiver)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _FastTransferTokenPoolAbstract.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) GetAllowList() ([]common.Address, error) {
	return _FastTransferTokenPoolAbstract.Contract.GetAllowList(&_FastTransferTokenPoolAbstract.CallOpts)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCallerSession) GetAllowList() ([]common.Address, error) {
	return _FastTransferTokenPoolAbstract.Contract.GetAllowList(&_FastTransferTokenPoolAbstract.CallOpts)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCaller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FastTransferTokenPoolAbstract.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) GetAllowListEnabled() (bool, error) {
	return _FastTransferTokenPoolAbstract.Contract.GetAllowListEnabled(&_FastTransferTokenPoolAbstract.CallOpts)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCallerSession) GetAllowListEnabled() (bool, error) {
	return _FastTransferTokenPoolAbstract.Contract.GetAllowListEnabled(&_FastTransferTokenPoolAbstract.CallOpts)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCaller) GetCcipSendTokenFee(opts *bind.CallOpts, settlementFeeToken common.Address, destinationChainSelector uint64, amount *big.Int, receiver []byte, extraArgs []byte) (IFastTransferPoolQuote, error) {
	var out []interface{}
	err := _FastTransferTokenPoolAbstract.contract.Call(opts, &out, "getCcipSendTokenFee", settlementFeeToken, destinationChainSelector, amount, receiver, extraArgs)

	if err != nil {
		return *new(IFastTransferPoolQuote), err
	}

	out0 := *abi.ConvertType(out[0], new(IFastTransferPoolQuote)).(*IFastTransferPoolQuote)

	return out0, err

}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) GetCcipSendTokenFee(settlementFeeToken common.Address, destinationChainSelector uint64, amount *big.Int, receiver []byte, extraArgs []byte) (IFastTransferPoolQuote, error) {
	return _FastTransferTokenPoolAbstract.Contract.GetCcipSendTokenFee(&_FastTransferTokenPoolAbstract.CallOpts, settlementFeeToken, destinationChainSelector, amount, receiver, extraArgs)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCallerSession) GetCcipSendTokenFee(settlementFeeToken common.Address, destinationChainSelector uint64, amount *big.Int, receiver []byte, extraArgs []byte) (IFastTransferPoolQuote, error) {
	return _FastTransferTokenPoolAbstract.Contract.GetCcipSendTokenFee(&_FastTransferTokenPoolAbstract.CallOpts, settlementFeeToken, destinationChainSelector, amount, receiver, extraArgs)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCaller) GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _FastTransferTokenPoolAbstract.contract.Call(opts, &out, "getCurrentInboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _FastTransferTokenPoolAbstract.Contract.GetCurrentInboundRateLimiterState(&_FastTransferTokenPoolAbstract.CallOpts, remoteChainSelector)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCallerSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _FastTransferTokenPoolAbstract.Contract.GetCurrentInboundRateLimiterState(&_FastTransferTokenPoolAbstract.CallOpts, remoteChainSelector)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCaller) GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _FastTransferTokenPoolAbstract.contract.Call(opts, &out, "getCurrentOutboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _FastTransferTokenPoolAbstract.Contract.GetCurrentOutboundRateLimiterState(&_FastTransferTokenPoolAbstract.CallOpts, remoteChainSelector)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCallerSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _FastTransferTokenPoolAbstract.Contract.GetCurrentOutboundRateLimiterState(&_FastTransferTokenPoolAbstract.CallOpts, remoteChainSelector)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCaller) GetDestChainConfig(opts *bind.CallOpts, remoteChainSelector uint64) (FastTransferTokenPoolAbstractDestChainConfigView, error) {
	var out []interface{}
	err := _FastTransferTokenPoolAbstract.contract.Call(opts, &out, "getDestChainConfig", remoteChainSelector)

	if err != nil {
		return *new(FastTransferTokenPoolAbstractDestChainConfigView), err
	}

	out0 := *abi.ConvertType(out[0], new(FastTransferTokenPoolAbstractDestChainConfigView)).(*FastTransferTokenPoolAbstractDestChainConfigView)

	return out0, err

}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) GetDestChainConfig(remoteChainSelector uint64) (FastTransferTokenPoolAbstractDestChainConfigView, error) {
	return _FastTransferTokenPoolAbstract.Contract.GetDestChainConfig(&_FastTransferTokenPoolAbstract.CallOpts, remoteChainSelector)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCallerSession) GetDestChainConfig(remoteChainSelector uint64) (FastTransferTokenPoolAbstractDestChainConfigView, error) {
	return _FastTransferTokenPoolAbstract.Contract.GetDestChainConfig(&_FastTransferTokenPoolAbstract.CallOpts, remoteChainSelector)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCaller) GetFillInfo(opts *bind.CallOpts, fillId [32]byte) (FastTransferTokenPoolAbstractFillInfo, error) {
	var out []interface{}
	err := _FastTransferTokenPoolAbstract.contract.Call(opts, &out, "getFillInfo", fillId)

	if err != nil {
		return *new(FastTransferTokenPoolAbstractFillInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(FastTransferTokenPoolAbstractFillInfo)).(*FastTransferTokenPoolAbstractFillInfo)

	return out0, err

}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) GetFillInfo(fillId [32]byte) (FastTransferTokenPoolAbstractFillInfo, error) {
	return _FastTransferTokenPoolAbstract.Contract.GetFillInfo(&_FastTransferTokenPoolAbstract.CallOpts, fillId)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCallerSession) GetFillInfo(fillId [32]byte) (FastTransferTokenPoolAbstractFillInfo, error) {
	return _FastTransferTokenPoolAbstract.Contract.GetFillInfo(&_FastTransferTokenPoolAbstract.CallOpts, fillId)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCaller) GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FastTransferTokenPoolAbstract.contract.Call(opts, &out, "getRateLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) GetRateLimitAdmin() (common.Address, error) {
	return _FastTransferTokenPoolAbstract.Contract.GetRateLimitAdmin(&_FastTransferTokenPoolAbstract.CallOpts)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCallerSession) GetRateLimitAdmin() (common.Address, error) {
	return _FastTransferTokenPoolAbstract.Contract.GetRateLimitAdmin(&_FastTransferTokenPoolAbstract.CallOpts)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCaller) GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error) {
	var out []interface{}
	err := _FastTransferTokenPoolAbstract.contract.Call(opts, &out, "getRemotePools", remoteChainSelector)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _FastTransferTokenPoolAbstract.Contract.GetRemotePools(&_FastTransferTokenPoolAbstract.CallOpts, remoteChainSelector)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCallerSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _FastTransferTokenPoolAbstract.Contract.GetRemotePools(&_FastTransferTokenPoolAbstract.CallOpts, remoteChainSelector)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCaller) GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _FastTransferTokenPoolAbstract.contract.Call(opts, &out, "getRemoteToken", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _FastTransferTokenPoolAbstract.Contract.GetRemoteToken(&_FastTransferTokenPoolAbstract.CallOpts, remoteChainSelector)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCallerSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _FastTransferTokenPoolAbstract.Contract.GetRemoteToken(&_FastTransferTokenPoolAbstract.CallOpts, remoteChainSelector)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCaller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FastTransferTokenPoolAbstract.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) GetRmnProxy() (common.Address, error) {
	return _FastTransferTokenPoolAbstract.Contract.GetRmnProxy(&_FastTransferTokenPoolAbstract.CallOpts)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCallerSession) GetRmnProxy() (common.Address, error) {
	return _FastTransferTokenPoolAbstract.Contract.GetRmnProxy(&_FastTransferTokenPoolAbstract.CallOpts)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FastTransferTokenPoolAbstract.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) GetRouter() (common.Address, error) {
	return _FastTransferTokenPoolAbstract.Contract.GetRouter(&_FastTransferTokenPoolAbstract.CallOpts)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCallerSession) GetRouter() (common.Address, error) {
	return _FastTransferTokenPoolAbstract.Contract.GetRouter(&_FastTransferTokenPoolAbstract.CallOpts)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _FastTransferTokenPoolAbstract.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) GetSupportedChains() ([]uint64, error) {
	return _FastTransferTokenPoolAbstract.Contract.GetSupportedChains(&_FastTransferTokenPoolAbstract.CallOpts)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCallerSession) GetSupportedChains() ([]uint64, error) {
	return _FastTransferTokenPoolAbstract.Contract.GetSupportedChains(&_FastTransferTokenPoolAbstract.CallOpts)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FastTransferTokenPoolAbstract.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) GetToken() (common.Address, error) {
	return _FastTransferTokenPoolAbstract.Contract.GetToken(&_FastTransferTokenPoolAbstract.CallOpts)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCallerSession) GetToken() (common.Address, error) {
	return _FastTransferTokenPoolAbstract.Contract.GetToken(&_FastTransferTokenPoolAbstract.CallOpts)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCaller) GetTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _FastTransferTokenPoolAbstract.contract.Call(opts, &out, "getTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) GetTokenDecimals() (uint8, error) {
	return _FastTransferTokenPoolAbstract.Contract.GetTokenDecimals(&_FastTransferTokenPoolAbstract.CallOpts)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCallerSession) GetTokenDecimals() (uint8, error) {
	return _FastTransferTokenPoolAbstract.Contract.GetTokenDecimals(&_FastTransferTokenPoolAbstract.CallOpts)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCaller) IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	var out []interface{}
	err := _FastTransferTokenPoolAbstract.contract.Call(opts, &out, "isRemotePool", remoteChainSelector, remotePoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _FastTransferTokenPoolAbstract.Contract.IsRemotePool(&_FastTransferTokenPoolAbstract.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCallerSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _FastTransferTokenPoolAbstract.Contract.IsRemotePool(&_FastTransferTokenPoolAbstract.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCaller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _FastTransferTokenPoolAbstract.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _FastTransferTokenPoolAbstract.Contract.IsSupportedChain(&_FastTransferTokenPoolAbstract.CallOpts, remoteChainSelector)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _FastTransferTokenPoolAbstract.Contract.IsSupportedChain(&_FastTransferTokenPoolAbstract.CallOpts, remoteChainSelector)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _FastTransferTokenPoolAbstract.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) IsSupportedToken(token common.Address) (bool, error) {
	return _FastTransferTokenPoolAbstract.Contract.IsSupportedToken(&_FastTransferTokenPoolAbstract.CallOpts, token)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _FastTransferTokenPoolAbstract.Contract.IsSupportedToken(&_FastTransferTokenPoolAbstract.CallOpts, token)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCaller) IsfillerAllowListed(opts *bind.CallOpts, remoteChainSelector uint64, filler common.Address) (bool, error) {
	var out []interface{}
	err := _FastTransferTokenPoolAbstract.contract.Call(opts, &out, "isfillerAllowListed", remoteChainSelector, filler)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) IsfillerAllowListed(remoteChainSelector uint64, filler common.Address) (bool, error) {
	return _FastTransferTokenPoolAbstract.Contract.IsfillerAllowListed(&_FastTransferTokenPoolAbstract.CallOpts, remoteChainSelector, filler)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCallerSession) IsfillerAllowListed(remoteChainSelector uint64, filler common.Address) (bool, error) {
	return _FastTransferTokenPoolAbstract.Contract.IsfillerAllowListed(&_FastTransferTokenPoolAbstract.CallOpts, remoteChainSelector, filler)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FastTransferTokenPoolAbstract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) Owner() (common.Address, error) {
	return _FastTransferTokenPoolAbstract.Contract.Owner(&_FastTransferTokenPoolAbstract.CallOpts)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCallerSession) Owner() (common.Address, error) {
	return _FastTransferTokenPoolAbstract.Contract.Owner(&_FastTransferTokenPoolAbstract.CallOpts)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _FastTransferTokenPoolAbstract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FastTransferTokenPoolAbstract.Contract.SupportsInterface(&_FastTransferTokenPoolAbstract.CallOpts, interfaceId)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FastTransferTokenPoolAbstract.Contract.SupportsInterface(&_FastTransferTokenPoolAbstract.CallOpts, interfaceId)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FastTransferTokenPoolAbstract.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) TypeAndVersion() (string, error) {
	return _FastTransferTokenPoolAbstract.Contract.TypeAndVersion(&_FastTransferTokenPoolAbstract.CallOpts)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractCallerSession) TypeAndVersion() (string, error) {
	return _FastTransferTokenPoolAbstract.Contract.TypeAndVersion(&_FastTransferTokenPoolAbstract.CallOpts)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.contract.Transact(opts, "acceptOwnership")
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) AcceptOwnership() (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.AcceptOwnership(&_FastTransferTokenPoolAbstract.TransactOpts)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.AcceptOwnership(&_FastTransferTokenPoolAbstract.TransactOpts)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactor) AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.contract.Transact(opts, "addRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.AddRemotePool(&_FastTransferTokenPoolAbstract.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactorSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.AddRemotePool(&_FastTransferTokenPoolAbstract.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.ApplyAllowListUpdates(&_FastTransferTokenPoolAbstract.TransactOpts, removes, adds)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.ApplyAllowListUpdates(&_FastTransferTokenPoolAbstract.TransactOpts, removes, adds)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactor) ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.contract.Transact(opts, "applyChainUpdates", remoteChainSelectorsToRemove, chainsToAdd)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.ApplyChainUpdates(&_FastTransferTokenPoolAbstract.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactorSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.ApplyChainUpdates(&_FastTransferTokenPoolAbstract.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactor) CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.contract.Transact(opts, "ccipReceive", message)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.CcipReceive(&_FastTransferTokenPoolAbstract.TransactOpts, message)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactorSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.CcipReceive(&_FastTransferTokenPoolAbstract.TransactOpts, message)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactor) CcipSendToken(opts *bind.TransactOpts, feeToken common.Address, destinationChainSelector uint64, amount *big.Int, receiver []byte, extraArgs []byte) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.contract.Transact(opts, "ccipSendToken", feeToken, destinationChainSelector, amount, receiver, extraArgs)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) CcipSendToken(feeToken common.Address, destinationChainSelector uint64, amount *big.Int, receiver []byte, extraArgs []byte) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.CcipSendToken(&_FastTransferTokenPoolAbstract.TransactOpts, feeToken, destinationChainSelector, amount, receiver, extraArgs)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactorSession) CcipSendToken(feeToken common.Address, destinationChainSelector uint64, amount *big.Int, receiver []byte, extraArgs []byte) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.CcipSendToken(&_FastTransferTokenPoolAbstract.TransactOpts, feeToken, destinationChainSelector, amount, receiver, extraArgs)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactor) FastFill(opts *bind.TransactOpts, fillRequestId [32]byte, sourceChainSelector uint64, srcAmount *big.Int, srcDecimals uint8, receiver common.Address) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.contract.Transact(opts, "fastFill", fillRequestId, sourceChainSelector, srcAmount, srcDecimals, receiver)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) FastFill(fillRequestId [32]byte, sourceChainSelector uint64, srcAmount *big.Int, srcDecimals uint8, receiver common.Address) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.FastFill(&_FastTransferTokenPoolAbstract.TransactOpts, fillRequestId, sourceChainSelector, srcAmount, srcDecimals, receiver)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactorSession) FastFill(fillRequestId [32]byte, sourceChainSelector uint64, srcAmount *big.Int, srcDecimals uint8, receiver common.Address) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.FastFill(&_FastTransferTokenPoolAbstract.TransactOpts, fillRequestId, sourceChainSelector, srcAmount, srcDecimals, receiver)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.LockOrBurn(&_FastTransferTokenPoolAbstract.TransactOpts, lockOrBurnIn)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.LockOrBurn(&_FastTransferTokenPoolAbstract.TransactOpts, lockOrBurnIn)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.ReleaseOrMint(&_FastTransferTokenPoolAbstract.TransactOpts, releaseOrMintIn)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.ReleaseOrMint(&_FastTransferTokenPoolAbstract.TransactOpts, releaseOrMintIn)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactor) RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.contract.Transact(opts, "removeRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.RemoveRemotePool(&_FastTransferTokenPoolAbstract.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactorSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.RemoveRemotePool(&_FastTransferTokenPoolAbstract.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactor) SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.contract.Transact(opts, "setChainRateLimiterConfig", remoteChainSelector, outboundConfig, inboundConfig)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.SetChainRateLimiterConfig(&_FastTransferTokenPoolAbstract.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactorSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.SetChainRateLimiterConfig(&_FastTransferTokenPoolAbstract.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactor) SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.contract.Transact(opts, "setChainRateLimiterConfigs", remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.SetChainRateLimiterConfigs(&_FastTransferTokenPoolAbstract.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactorSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.SetChainRateLimiterConfigs(&_FastTransferTokenPoolAbstract.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactor) SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.contract.Transact(opts, "setRateLimitAdmin", rateLimitAdmin)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.SetRateLimitAdmin(&_FastTransferTokenPoolAbstract.TransactOpts, rateLimitAdmin)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactorSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.SetRateLimitAdmin(&_FastTransferTokenPoolAbstract.TransactOpts, rateLimitAdmin)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactor) SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.contract.Transact(opts, "setRouter", newRouter)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.SetRouter(&_FastTransferTokenPoolAbstract.TransactOpts, newRouter)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactorSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.SetRouter(&_FastTransferTokenPoolAbstract.TransactOpts, newRouter)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.contract.Transact(opts, "transferOwnership", to)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.TransferOwnership(&_FastTransferTokenPoolAbstract.TransactOpts, to)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.TransferOwnership(&_FastTransferTokenPoolAbstract.TransactOpts, to)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactor) UpdateDestChainConfig(opts *bind.TransactOpts, laneConfigArgs FastTransferTokenPoolAbstractDestChainConfigUpdateArgs) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.contract.Transact(opts, "updateDestChainConfig", laneConfigArgs)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) UpdateDestChainConfig(laneConfigArgs FastTransferTokenPoolAbstractDestChainConfigUpdateArgs) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.UpdateDestChainConfig(&_FastTransferTokenPoolAbstract.TransactOpts, laneConfigArgs)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactorSession) UpdateDestChainConfig(laneConfigArgs FastTransferTokenPoolAbstractDestChainConfigUpdateArgs) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.UpdateDestChainConfig(&_FastTransferTokenPoolAbstract.TransactOpts, laneConfigArgs)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactor) UpdatefillerAllowList(opts *bind.TransactOpts, destinationChainSelector uint64, addFillers []common.Address, removeFillers []common.Address) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.contract.Transact(opts, "updatefillerAllowList", destinationChainSelector, addFillers, removeFillers)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractSession) UpdatefillerAllowList(destinationChainSelector uint64, addFillers []common.Address, removeFillers []common.Address) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.UpdatefillerAllowList(&_FastTransferTokenPoolAbstract.TransactOpts, destinationChainSelector, addFillers, removeFillers)
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractTransactorSession) UpdatefillerAllowList(destinationChainSelector uint64, addFillers []common.Address, removeFillers []common.Address) (*types.Transaction, error) {
	return _FastTransferTokenPoolAbstract.Contract.UpdatefillerAllowList(&_FastTransferTokenPoolAbstract.TransactOpts, destinationChainSelector, addFillers, removeFillers)
}

type FastTransferTokenPoolAbstractAllowListAddIterator struct {
	Event *FastTransferTokenPoolAbstractAllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FastTransferTokenPoolAbstractAllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastTransferTokenPoolAbstractAllowListAdd)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FastTransferTokenPoolAbstractAllowListAdd)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FastTransferTokenPoolAbstractAllowListAddIterator) Error() error {
	return it.fail
}

func (it *FastTransferTokenPoolAbstractAllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FastTransferTokenPoolAbstractAllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) FilterAllowListAdd(opts *bind.FilterOpts) (*FastTransferTokenPoolAbstractAllowListAddIterator, error) {

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &FastTransferTokenPoolAbstractAllowListAddIterator{contract: _FastTransferTokenPoolAbstract.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractAllowListAdd) (event.Subscription, error) {

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FastTransferTokenPoolAbstractAllowListAdd)
				if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) ParseAllowListAdd(log types.Log) (*FastTransferTokenPoolAbstractAllowListAdd, error) {
	event := new(FastTransferTokenPoolAbstractAllowListAdd)
	if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FastTransferTokenPoolAbstractAllowListRemoveIterator struct {
	Event *FastTransferTokenPoolAbstractAllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FastTransferTokenPoolAbstractAllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastTransferTokenPoolAbstractAllowListRemove)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FastTransferTokenPoolAbstractAllowListRemove)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FastTransferTokenPoolAbstractAllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *FastTransferTokenPoolAbstractAllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FastTransferTokenPoolAbstractAllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) FilterAllowListRemove(opts *bind.FilterOpts) (*FastTransferTokenPoolAbstractAllowListRemoveIterator, error) {

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &FastTransferTokenPoolAbstractAllowListRemoveIterator{contract: _FastTransferTokenPoolAbstract.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractAllowListRemove) (event.Subscription, error) {

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FastTransferTokenPoolAbstractAllowListRemove)
				if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) ParseAllowListRemove(log types.Log) (*FastTransferTokenPoolAbstractAllowListRemove, error) {
	event := new(FastTransferTokenPoolAbstractAllowListRemove)
	if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FastTransferTokenPoolAbstractBurnedIterator struct {
	Event *FastTransferTokenPoolAbstractBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FastTransferTokenPoolAbstractBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastTransferTokenPoolAbstractBurned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FastTransferTokenPoolAbstractBurned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FastTransferTokenPoolAbstractBurnedIterator) Error() error {
	return it.fail
}

func (it *FastTransferTokenPoolAbstractBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FastTransferTokenPoolAbstractBurned struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) FilterBurned(opts *bind.FilterOpts, sender []common.Address) (*FastTransferTokenPoolAbstractBurnedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.FilterLogs(opts, "Burned", senderRule)
	if err != nil {
		return nil, err
	}
	return &FastTransferTokenPoolAbstractBurnedIterator{contract: _FastTransferTokenPoolAbstract.contract, event: "Burned", logs: logs, sub: sub}, nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractBurned, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.WatchLogs(opts, "Burned", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FastTransferTokenPoolAbstractBurned)
				if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "Burned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) ParseBurned(log types.Log) (*FastTransferTokenPoolAbstractBurned, error) {
	event := new(FastTransferTokenPoolAbstractBurned)
	if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FastTransferTokenPoolAbstractChainAddedIterator struct {
	Event *FastTransferTokenPoolAbstractChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FastTransferTokenPoolAbstractChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastTransferTokenPoolAbstractChainAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FastTransferTokenPoolAbstractChainAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FastTransferTokenPoolAbstractChainAddedIterator) Error() error {
	return it.fail
}

func (it *FastTransferTokenPoolAbstractChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FastTransferTokenPoolAbstractChainAdded struct {
	RemoteChainSelector       uint64
	RemoteToken               []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) FilterChainAdded(opts *bind.FilterOpts) (*FastTransferTokenPoolAbstractChainAddedIterator, error) {

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &FastTransferTokenPoolAbstractChainAddedIterator{contract: _FastTransferTokenPoolAbstract.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractChainAdded) (event.Subscription, error) {

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FastTransferTokenPoolAbstractChainAdded)
				if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "ChainAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) ParseChainAdded(log types.Log) (*FastTransferTokenPoolAbstractChainAdded, error) {
	event := new(FastTransferTokenPoolAbstractChainAdded)
	if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FastTransferTokenPoolAbstractChainConfiguredIterator struct {
	Event *FastTransferTokenPoolAbstractChainConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FastTransferTokenPoolAbstractChainConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastTransferTokenPoolAbstractChainConfigured)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FastTransferTokenPoolAbstractChainConfigured)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FastTransferTokenPoolAbstractChainConfiguredIterator) Error() error {
	return it.fail
}

func (it *FastTransferTokenPoolAbstractChainConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FastTransferTokenPoolAbstractChainConfigured struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) FilterChainConfigured(opts *bind.FilterOpts) (*FastTransferTokenPoolAbstractChainConfiguredIterator, error) {

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.FilterLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return &FastTransferTokenPoolAbstractChainConfiguredIterator{contract: _FastTransferTokenPoolAbstract.contract, event: "ChainConfigured", logs: logs, sub: sub}, nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractChainConfigured) (event.Subscription, error) {

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.WatchLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FastTransferTokenPoolAbstractChainConfigured)
				if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) ParseChainConfigured(log types.Log) (*FastTransferTokenPoolAbstractChainConfigured, error) {
	event := new(FastTransferTokenPoolAbstractChainConfigured)
	if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FastTransferTokenPoolAbstractChainRemovedIterator struct {
	Event *FastTransferTokenPoolAbstractChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FastTransferTokenPoolAbstractChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastTransferTokenPoolAbstractChainRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FastTransferTokenPoolAbstractChainRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FastTransferTokenPoolAbstractChainRemovedIterator) Error() error {
	return it.fail
}

func (it *FastTransferTokenPoolAbstractChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FastTransferTokenPoolAbstractChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) FilterChainRemoved(opts *bind.FilterOpts) (*FastTransferTokenPoolAbstractChainRemovedIterator, error) {

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &FastTransferTokenPoolAbstractChainRemovedIterator{contract: _FastTransferTokenPoolAbstract.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractChainRemoved) (event.Subscription, error) {

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FastTransferTokenPoolAbstractChainRemoved)
				if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) ParseChainRemoved(log types.Log) (*FastTransferTokenPoolAbstractChainRemoved, error) {
	event := new(FastTransferTokenPoolAbstractChainRemoved)
	if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FastTransferTokenPoolAbstractConfigChangedIterator struct {
	Event *FastTransferTokenPoolAbstractConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FastTransferTokenPoolAbstractConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastTransferTokenPoolAbstractConfigChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FastTransferTokenPoolAbstractConfigChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FastTransferTokenPoolAbstractConfigChangedIterator) Error() error {
	return it.fail
}

func (it *FastTransferTokenPoolAbstractConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FastTransferTokenPoolAbstractConfigChanged struct {
	Config RateLimiterConfig
	Raw    types.Log
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*FastTransferTokenPoolAbstractConfigChangedIterator, error) {

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &FastTransferTokenPoolAbstractConfigChangedIterator{contract: _FastTransferTokenPoolAbstract.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractConfigChanged) (event.Subscription, error) {

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FastTransferTokenPoolAbstractConfigChanged)
				if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) ParseConfigChanged(log types.Log) (*FastTransferTokenPoolAbstractConfigChanged, error) {
	event := new(FastTransferTokenPoolAbstractConfigChanged)
	if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FastTransferTokenPoolAbstractDestinationPoolUpdatedIterator struct {
	Event *FastTransferTokenPoolAbstractDestinationPoolUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FastTransferTokenPoolAbstractDestinationPoolUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastTransferTokenPoolAbstractDestinationPoolUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FastTransferTokenPoolAbstractDestinationPoolUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FastTransferTokenPoolAbstractDestinationPoolUpdatedIterator) Error() error {
	return it.fail
}

func (it *FastTransferTokenPoolAbstractDestinationPoolUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FastTransferTokenPoolAbstractDestinationPoolUpdated struct {
	Dst             uint64
	DestinationPool common.Address
	Raw             types.Log
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) FilterDestinationPoolUpdated(opts *bind.FilterOpts, dst []uint64) (*FastTransferTokenPoolAbstractDestinationPoolUpdatedIterator, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.FilterLogs(opts, "DestinationPoolUpdated", dstRule)
	if err != nil {
		return nil, err
	}
	return &FastTransferTokenPoolAbstractDestinationPoolUpdatedIterator{contract: _FastTransferTokenPoolAbstract.contract, event: "DestinationPoolUpdated", logs: logs, sub: sub}, nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) WatchDestinationPoolUpdated(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractDestinationPoolUpdated, dst []uint64) (event.Subscription, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.WatchLogs(opts, "DestinationPoolUpdated", dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FastTransferTokenPoolAbstractDestinationPoolUpdated)
				if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "DestinationPoolUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) ParseDestinationPoolUpdated(log types.Log) (*FastTransferTokenPoolAbstractDestinationPoolUpdated, error) {
	event := new(FastTransferTokenPoolAbstractDestinationPoolUpdated)
	if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "DestinationPoolUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FastTransferTokenPoolAbstractFastFillIterator struct {
	Event *FastTransferTokenPoolAbstractFastFill

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FastTransferTokenPoolAbstractFastFillIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastTransferTokenPoolAbstractFastFill)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FastTransferTokenPoolAbstractFastFill)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FastTransferTokenPoolAbstractFastFillIterator) Error() error {
	return it.fail
}

func (it *FastTransferTokenPoolAbstractFastFillIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FastTransferTokenPoolAbstractFastFill struct {
	FillRequestId [32]byte
	FillId        [32]byte
	Filler        common.Address
	DestAmount    *big.Int
	Receiver      common.Address
	Raw           types.Log
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) FilterFastFill(opts *bind.FilterOpts, fillRequestId [][32]byte, fillId [][32]byte, filler []common.Address) (*FastTransferTokenPoolAbstractFastFillIterator, error) {

	var fillRequestIdRule []interface{}
	for _, fillRequestIdItem := range fillRequestId {
		fillRequestIdRule = append(fillRequestIdRule, fillRequestIdItem)
	}
	var fillIdRule []interface{}
	for _, fillIdItem := range fillId {
		fillIdRule = append(fillIdRule, fillIdItem)
	}
	var fillerRule []interface{}
	for _, fillerItem := range filler {
		fillerRule = append(fillerRule, fillerItem)
	}

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.FilterLogs(opts, "FastFill", fillRequestIdRule, fillIdRule, fillerRule)
	if err != nil {
		return nil, err
	}
	return &FastTransferTokenPoolAbstractFastFillIterator{contract: _FastTransferTokenPoolAbstract.contract, event: "FastFill", logs: logs, sub: sub}, nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) WatchFastFill(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractFastFill, fillRequestId [][32]byte, fillId [][32]byte, filler []common.Address) (event.Subscription, error) {

	var fillRequestIdRule []interface{}
	for _, fillRequestIdItem := range fillRequestId {
		fillRequestIdRule = append(fillRequestIdRule, fillRequestIdItem)
	}
	var fillIdRule []interface{}
	for _, fillIdItem := range fillId {
		fillIdRule = append(fillIdRule, fillIdItem)
	}
	var fillerRule []interface{}
	for _, fillerItem := range filler {
		fillerRule = append(fillerRule, fillerItem)
	}

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.WatchLogs(opts, "FastFill", fillRequestIdRule, fillIdRule, fillerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FastTransferTokenPoolAbstractFastFill)
				if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "FastFill", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) ParseFastFill(log types.Log) (*FastTransferTokenPoolAbstractFastFill, error) {
	event := new(FastTransferTokenPoolAbstractFastFill)
	if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "FastFill", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FastTransferTokenPoolAbstractFastFillRequestIterator struct {
	Event *FastTransferTokenPoolAbstractFastFillRequest

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FastTransferTokenPoolAbstractFastFillRequestIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastTransferTokenPoolAbstractFastFillRequest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FastTransferTokenPoolAbstractFastFillRequest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FastTransferTokenPoolAbstractFastFillRequestIterator) Error() error {
	return it.fail
}

func (it *FastTransferTokenPoolAbstractFastFillRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FastTransferTokenPoolAbstractFastFillRequest struct {
	FillRequestId    [32]byte
	DstChainSelector uint64
	Amount           *big.Int
	FastTransferFee  *big.Int
	Receiver         []byte
	Raw              types.Log
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) FilterFastFillRequest(opts *bind.FilterOpts, fillRequestId [][32]byte, dstChainSelector []uint64) (*FastTransferTokenPoolAbstractFastFillRequestIterator, error) {

	var fillRequestIdRule []interface{}
	for _, fillRequestIdItem := range fillRequestId {
		fillRequestIdRule = append(fillRequestIdRule, fillRequestIdItem)
	}
	var dstChainSelectorRule []interface{}
	for _, dstChainSelectorItem := range dstChainSelector {
		dstChainSelectorRule = append(dstChainSelectorRule, dstChainSelectorItem)
	}

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.FilterLogs(opts, "FastFillRequest", fillRequestIdRule, dstChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &FastTransferTokenPoolAbstractFastFillRequestIterator{contract: _FastTransferTokenPoolAbstract.contract, event: "FastFillRequest", logs: logs, sub: sub}, nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) WatchFastFillRequest(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractFastFillRequest, fillRequestId [][32]byte, dstChainSelector []uint64) (event.Subscription, error) {

	var fillRequestIdRule []interface{}
	for _, fillRequestIdItem := range fillRequestId {
		fillRequestIdRule = append(fillRequestIdRule, fillRequestIdItem)
	}
	var dstChainSelectorRule []interface{}
	for _, dstChainSelectorItem := range dstChainSelector {
		dstChainSelectorRule = append(dstChainSelectorRule, dstChainSelectorItem)
	}

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.WatchLogs(opts, "FastFillRequest", fillRequestIdRule, dstChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FastTransferTokenPoolAbstractFastFillRequest)
				if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "FastFillRequest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) ParseFastFillRequest(log types.Log) (*FastTransferTokenPoolAbstractFastFillRequest, error) {
	event := new(FastTransferTokenPoolAbstractFastFillRequest)
	if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "FastFillRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FastTransferTokenPoolAbstractFastFillSettledIterator struct {
	Event *FastTransferTokenPoolAbstractFastFillSettled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FastTransferTokenPoolAbstractFastFillSettledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastTransferTokenPoolAbstractFastFillSettled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FastTransferTokenPoolAbstractFastFillSettled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FastTransferTokenPoolAbstractFastFillSettledIterator) Error() error {
	return it.fail
}

func (it *FastTransferTokenPoolAbstractFastFillSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FastTransferTokenPoolAbstractFastFillSettled struct {
	FillRequestId [32]byte
	Raw           types.Log
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) FilterFastFillSettled(opts *bind.FilterOpts, fillRequestId [][32]byte) (*FastTransferTokenPoolAbstractFastFillSettledIterator, error) {

	var fillRequestIdRule []interface{}
	for _, fillRequestIdItem := range fillRequestId {
		fillRequestIdRule = append(fillRequestIdRule, fillRequestIdItem)
	}

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.FilterLogs(opts, "FastFillSettled", fillRequestIdRule)
	if err != nil {
		return nil, err
	}
	return &FastTransferTokenPoolAbstractFastFillSettledIterator{contract: _FastTransferTokenPoolAbstract.contract, event: "FastFillSettled", logs: logs, sub: sub}, nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) WatchFastFillSettled(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractFastFillSettled, fillRequestId [][32]byte) (event.Subscription, error) {

	var fillRequestIdRule []interface{}
	for _, fillRequestIdItem := range fillRequestId {
		fillRequestIdRule = append(fillRequestIdRule, fillRequestIdItem)
	}

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.WatchLogs(opts, "FastFillSettled", fillRequestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FastTransferTokenPoolAbstractFastFillSettled)
				if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "FastFillSettled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) ParseFastFillSettled(log types.Log) (*FastTransferTokenPoolAbstractFastFillSettled, error) {
	event := new(FastTransferTokenPoolAbstractFastFillSettled)
	if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "FastFillSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FastTransferTokenPoolAbstractFillerAllowListUpdatedIterator struct {
	Event *FastTransferTokenPoolAbstractFillerAllowListUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FastTransferTokenPoolAbstractFillerAllowListUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastTransferTokenPoolAbstractFillerAllowListUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FastTransferTokenPoolAbstractFillerAllowListUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FastTransferTokenPoolAbstractFillerAllowListUpdatedIterator) Error() error {
	return it.fail
}

func (it *FastTransferTokenPoolAbstractFillerAllowListUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FastTransferTokenPoolAbstractFillerAllowListUpdated struct {
	Dst           uint64
	AddFillers    []common.Address
	RemoveFillers []common.Address
	Raw           types.Log
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) FilterFillerAllowListUpdated(opts *bind.FilterOpts, dst []uint64) (*FastTransferTokenPoolAbstractFillerAllowListUpdatedIterator, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.FilterLogs(opts, "FillerAllowListUpdated", dstRule)
	if err != nil {
		return nil, err
	}
	return &FastTransferTokenPoolAbstractFillerAllowListUpdatedIterator{contract: _FastTransferTokenPoolAbstract.contract, event: "FillerAllowListUpdated", logs: logs, sub: sub}, nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) WatchFillerAllowListUpdated(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractFillerAllowListUpdated, dst []uint64) (event.Subscription, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.WatchLogs(opts, "FillerAllowListUpdated", dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FastTransferTokenPoolAbstractFillerAllowListUpdated)
				if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "FillerAllowListUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) ParseFillerAllowListUpdated(log types.Log) (*FastTransferTokenPoolAbstractFillerAllowListUpdated, error) {
	event := new(FastTransferTokenPoolAbstractFillerAllowListUpdated)
	if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "FillerAllowListUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FastTransferTokenPoolAbstractInvalidFillIterator struct {
	Event *FastTransferTokenPoolAbstractInvalidFill

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FastTransferTokenPoolAbstractInvalidFillIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastTransferTokenPoolAbstractInvalidFill)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FastTransferTokenPoolAbstractInvalidFill)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FastTransferTokenPoolAbstractInvalidFillIterator) Error() error {
	return it.fail
}

func (it *FastTransferTokenPoolAbstractInvalidFillIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FastTransferTokenPoolAbstractInvalidFill struct {
	FillRequestId  [32]byte
	Filler         common.Address
	FilledAmount   *big.Int
	ExpectedAmount *big.Int
	Raw            types.Log
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) FilterInvalidFill(opts *bind.FilterOpts, fillRequestId [][32]byte, filler []common.Address) (*FastTransferTokenPoolAbstractInvalidFillIterator, error) {

	var fillRequestIdRule []interface{}
	for _, fillRequestIdItem := range fillRequestId {
		fillRequestIdRule = append(fillRequestIdRule, fillRequestIdItem)
	}
	var fillerRule []interface{}
	for _, fillerItem := range filler {
		fillerRule = append(fillerRule, fillerItem)
	}

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.FilterLogs(opts, "InvalidFill", fillRequestIdRule, fillerRule)
	if err != nil {
		return nil, err
	}
	return &FastTransferTokenPoolAbstractInvalidFillIterator{contract: _FastTransferTokenPoolAbstract.contract, event: "InvalidFill", logs: logs, sub: sub}, nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) WatchInvalidFill(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractInvalidFill, fillRequestId [][32]byte, filler []common.Address) (event.Subscription, error) {

	var fillRequestIdRule []interface{}
	for _, fillRequestIdItem := range fillRequestId {
		fillRequestIdRule = append(fillRequestIdRule, fillRequestIdItem)
	}
	var fillerRule []interface{}
	for _, fillerItem := range filler {
		fillerRule = append(fillerRule, fillerItem)
	}

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.WatchLogs(opts, "InvalidFill", fillRequestIdRule, fillerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FastTransferTokenPoolAbstractInvalidFill)
				if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "InvalidFill", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) ParseInvalidFill(log types.Log) (*FastTransferTokenPoolAbstractInvalidFill, error) {
	event := new(FastTransferTokenPoolAbstractInvalidFill)
	if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "InvalidFill", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FastTransferTokenPoolAbstractLaneUpdatedIterator struct {
	Event *FastTransferTokenPoolAbstractLaneUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FastTransferTokenPoolAbstractLaneUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastTransferTokenPoolAbstractLaneUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FastTransferTokenPoolAbstractLaneUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FastTransferTokenPoolAbstractLaneUpdatedIterator) Error() error {
	return it.fail
}

func (it *FastTransferTokenPoolAbstractLaneUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FastTransferTokenPoolAbstractLaneUpdated struct {
	DestinationChainSelector uint64
	Bps                      uint16
	MaxFillAmountPerRequest  *big.Int
	DestinationPool          []byte
	AddFillers               []common.Address
	RemoveFillers            []common.Address
	Raw                      types.Log
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) FilterLaneUpdated(opts *bind.FilterOpts, destinationChainSelector []uint64) (*FastTransferTokenPoolAbstractLaneUpdatedIterator, error) {

	var destinationChainSelectorRule []interface{}
	for _, destinationChainSelectorItem := range destinationChainSelector {
		destinationChainSelectorRule = append(destinationChainSelectorRule, destinationChainSelectorItem)
	}

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.FilterLogs(opts, "LaneUpdated", destinationChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &FastTransferTokenPoolAbstractLaneUpdatedIterator{contract: _FastTransferTokenPoolAbstract.contract, event: "LaneUpdated", logs: logs, sub: sub}, nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) WatchLaneUpdated(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractLaneUpdated, destinationChainSelector []uint64) (event.Subscription, error) {

	var destinationChainSelectorRule []interface{}
	for _, destinationChainSelectorItem := range destinationChainSelector {
		destinationChainSelectorRule = append(destinationChainSelectorRule, destinationChainSelectorItem)
	}

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.WatchLogs(opts, "LaneUpdated", destinationChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FastTransferTokenPoolAbstractLaneUpdated)
				if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "LaneUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) ParseLaneUpdated(log types.Log) (*FastTransferTokenPoolAbstractLaneUpdated, error) {
	event := new(FastTransferTokenPoolAbstractLaneUpdated)
	if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "LaneUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FastTransferTokenPoolAbstractLockedIterator struct {
	Event *FastTransferTokenPoolAbstractLocked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FastTransferTokenPoolAbstractLockedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastTransferTokenPoolAbstractLocked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FastTransferTokenPoolAbstractLocked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FastTransferTokenPoolAbstractLockedIterator) Error() error {
	return it.fail
}

func (it *FastTransferTokenPoolAbstractLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FastTransferTokenPoolAbstractLocked struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*FastTransferTokenPoolAbstractLockedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.FilterLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return &FastTransferTokenPoolAbstractLockedIterator{contract: _FastTransferTokenPoolAbstract.contract, event: "Locked", logs: logs, sub: sub}, nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractLocked, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.WatchLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FastTransferTokenPoolAbstractLocked)
				if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "Locked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) ParseLocked(log types.Log) (*FastTransferTokenPoolAbstractLocked, error) {
	event := new(FastTransferTokenPoolAbstractLocked)
	if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FastTransferTokenPoolAbstractMintedIterator struct {
	Event *FastTransferTokenPoolAbstractMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FastTransferTokenPoolAbstractMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastTransferTokenPoolAbstractMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FastTransferTokenPoolAbstractMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FastTransferTokenPoolAbstractMintedIterator) Error() error {
	return it.fail
}

func (it *FastTransferTokenPoolAbstractMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FastTransferTokenPoolAbstractMinted struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*FastTransferTokenPoolAbstractMintedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.FilterLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &FastTransferTokenPoolAbstractMintedIterator{contract: _FastTransferTokenPoolAbstract.contract, event: "Minted", logs: logs, sub: sub}, nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.WatchLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FastTransferTokenPoolAbstractMinted)
				if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "Minted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) ParseMinted(log types.Log) (*FastTransferTokenPoolAbstractMinted, error) {
	event := new(FastTransferTokenPoolAbstractMinted)
	if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FastTransferTokenPoolAbstractOwnershipTransferRequestedIterator struct {
	Event *FastTransferTokenPoolAbstractOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FastTransferTokenPoolAbstractOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastTransferTokenPoolAbstractOwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FastTransferTokenPoolAbstractOwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FastTransferTokenPoolAbstractOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *FastTransferTokenPoolAbstractOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FastTransferTokenPoolAbstractOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FastTransferTokenPoolAbstractOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FastTransferTokenPoolAbstractOwnershipTransferRequestedIterator{contract: _FastTransferTokenPoolAbstract.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FastTransferTokenPoolAbstractOwnershipTransferRequested)
				if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) ParseOwnershipTransferRequested(log types.Log) (*FastTransferTokenPoolAbstractOwnershipTransferRequested, error) {
	event := new(FastTransferTokenPoolAbstractOwnershipTransferRequested)
	if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FastTransferTokenPoolAbstractOwnershipTransferredIterator struct {
	Event *FastTransferTokenPoolAbstractOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FastTransferTokenPoolAbstractOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastTransferTokenPoolAbstractOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FastTransferTokenPoolAbstractOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FastTransferTokenPoolAbstractOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *FastTransferTokenPoolAbstractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FastTransferTokenPoolAbstractOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FastTransferTokenPoolAbstractOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FastTransferTokenPoolAbstractOwnershipTransferredIterator{contract: _FastTransferTokenPoolAbstract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FastTransferTokenPoolAbstractOwnershipTransferred)
				if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) ParseOwnershipTransferred(log types.Log) (*FastTransferTokenPoolAbstractOwnershipTransferred, error) {
	event := new(FastTransferTokenPoolAbstractOwnershipTransferred)
	if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FastTransferTokenPoolAbstractRateLimitAdminSetIterator struct {
	Event *FastTransferTokenPoolAbstractRateLimitAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FastTransferTokenPoolAbstractRateLimitAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastTransferTokenPoolAbstractRateLimitAdminSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FastTransferTokenPoolAbstractRateLimitAdminSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FastTransferTokenPoolAbstractRateLimitAdminSetIterator) Error() error {
	return it.fail
}

func (it *FastTransferTokenPoolAbstractRateLimitAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FastTransferTokenPoolAbstractRateLimitAdminSet struct {
	RateLimitAdmin common.Address
	Raw            types.Log
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) FilterRateLimitAdminSet(opts *bind.FilterOpts) (*FastTransferTokenPoolAbstractRateLimitAdminSetIterator, error) {

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.FilterLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return &FastTransferTokenPoolAbstractRateLimitAdminSetIterator{contract: _FastTransferTokenPoolAbstract.contract, event: "RateLimitAdminSet", logs: logs, sub: sub}, nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractRateLimitAdminSet) (event.Subscription, error) {

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.WatchLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FastTransferTokenPoolAbstractRateLimitAdminSet)
				if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) ParseRateLimitAdminSet(log types.Log) (*FastTransferTokenPoolAbstractRateLimitAdminSet, error) {
	event := new(FastTransferTokenPoolAbstractRateLimitAdminSet)
	if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FastTransferTokenPoolAbstractReleasedIterator struct {
	Event *FastTransferTokenPoolAbstractReleased

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FastTransferTokenPoolAbstractReleasedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastTransferTokenPoolAbstractReleased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FastTransferTokenPoolAbstractReleased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FastTransferTokenPoolAbstractReleasedIterator) Error() error {
	return it.fail
}

func (it *FastTransferTokenPoolAbstractReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FastTransferTokenPoolAbstractReleased struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*FastTransferTokenPoolAbstractReleasedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.FilterLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &FastTransferTokenPoolAbstractReleasedIterator{contract: _FastTransferTokenPoolAbstract.contract, event: "Released", logs: logs, sub: sub}, nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) WatchReleased(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.WatchLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FastTransferTokenPoolAbstractReleased)
				if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "Released", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) ParseReleased(log types.Log) (*FastTransferTokenPoolAbstractReleased, error) {
	event := new(FastTransferTokenPoolAbstractReleased)
	if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "Released", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FastTransferTokenPoolAbstractRemotePoolAddedIterator struct {
	Event *FastTransferTokenPoolAbstractRemotePoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FastTransferTokenPoolAbstractRemotePoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastTransferTokenPoolAbstractRemotePoolAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FastTransferTokenPoolAbstractRemotePoolAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FastTransferTokenPoolAbstractRemotePoolAddedIterator) Error() error {
	return it.fail
}

func (it *FastTransferTokenPoolAbstractRemotePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FastTransferTokenPoolAbstractRemotePoolAdded struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*FastTransferTokenPoolAbstractRemotePoolAddedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.FilterLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &FastTransferTokenPoolAbstractRemotePoolAddedIterator{contract: _FastTransferTokenPoolAbstract.contract, event: "RemotePoolAdded", logs: logs, sub: sub}, nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.WatchLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FastTransferTokenPoolAbstractRemotePoolAdded)
				if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) ParseRemotePoolAdded(log types.Log) (*FastTransferTokenPoolAbstractRemotePoolAdded, error) {
	event := new(FastTransferTokenPoolAbstractRemotePoolAdded)
	if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FastTransferTokenPoolAbstractRemotePoolRemovedIterator struct {
	Event *FastTransferTokenPoolAbstractRemotePoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FastTransferTokenPoolAbstractRemotePoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastTransferTokenPoolAbstractRemotePoolRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FastTransferTokenPoolAbstractRemotePoolRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FastTransferTokenPoolAbstractRemotePoolRemovedIterator) Error() error {
	return it.fail
}

func (it *FastTransferTokenPoolAbstractRemotePoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FastTransferTokenPoolAbstractRemotePoolRemoved struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*FastTransferTokenPoolAbstractRemotePoolRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.FilterLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &FastTransferTokenPoolAbstractRemotePoolRemovedIterator{contract: _FastTransferTokenPoolAbstract.contract, event: "RemotePoolRemoved", logs: logs, sub: sub}, nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.WatchLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FastTransferTokenPoolAbstractRemotePoolRemoved)
				if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) ParseRemotePoolRemoved(log types.Log) (*FastTransferTokenPoolAbstractRemotePoolRemoved, error) {
	event := new(FastTransferTokenPoolAbstractRemotePoolRemoved)
	if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FastTransferTokenPoolAbstractRouterUpdatedIterator struct {
	Event *FastTransferTokenPoolAbstractRouterUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FastTransferTokenPoolAbstractRouterUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastTransferTokenPoolAbstractRouterUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FastTransferTokenPoolAbstractRouterUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FastTransferTokenPoolAbstractRouterUpdatedIterator) Error() error {
	return it.fail
}

func (it *FastTransferTokenPoolAbstractRouterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FastTransferTokenPoolAbstractRouterUpdated struct {
	OldRouter common.Address
	NewRouter common.Address
	Raw       types.Log
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) FilterRouterUpdated(opts *bind.FilterOpts) (*FastTransferTokenPoolAbstractRouterUpdatedIterator, error) {

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.FilterLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return &FastTransferTokenPoolAbstractRouterUpdatedIterator{contract: _FastTransferTokenPoolAbstract.contract, event: "RouterUpdated", logs: logs, sub: sub}, nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractRouterUpdated) (event.Subscription, error) {

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.WatchLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FastTransferTokenPoolAbstractRouterUpdated)
				if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) ParseRouterUpdated(log types.Log) (*FastTransferTokenPoolAbstractRouterUpdated, error) {
	event := new(FastTransferTokenPoolAbstractRouterUpdated)
	if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FastTransferTokenPoolAbstractTokensConsumedIterator struct {
	Event *FastTransferTokenPoolAbstractTokensConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FastTransferTokenPoolAbstractTokensConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastTransferTokenPoolAbstractTokensConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(FastTransferTokenPoolAbstractTokensConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *FastTransferTokenPoolAbstractTokensConsumedIterator) Error() error {
	return it.fail
}

func (it *FastTransferTokenPoolAbstractTokensConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FastTransferTokenPoolAbstractTokensConsumed struct {
	Tokens *big.Int
	Raw    types.Log
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) FilterTokensConsumed(opts *bind.FilterOpts) (*FastTransferTokenPoolAbstractTokensConsumedIterator, error) {

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.FilterLogs(opts, "TokensConsumed")
	if err != nil {
		return nil, err
	}
	return &FastTransferTokenPoolAbstractTokensConsumedIterator{contract: _FastTransferTokenPoolAbstract.contract, event: "TokensConsumed", logs: logs, sub: sub}, nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) WatchTokensConsumed(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractTokensConsumed) (event.Subscription, error) {

	logs, sub, err := _FastTransferTokenPoolAbstract.contract.WatchLogs(opts, "TokensConsumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FastTransferTokenPoolAbstractTokensConsumed)
				if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "TokensConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstractFilterer) ParseTokensConsumed(log types.Log) (*FastTransferTokenPoolAbstractTokensConsumed, error) {
	event := new(FastTransferTokenPoolAbstractTokensConsumed)
	if err := _FastTransferTokenPoolAbstract.contract.UnpackLog(event, "TokensConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstract) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _FastTransferTokenPoolAbstract.abi.Events["AllowListAdd"].ID:
		return _FastTransferTokenPoolAbstract.ParseAllowListAdd(log)
	case _FastTransferTokenPoolAbstract.abi.Events["AllowListRemove"].ID:
		return _FastTransferTokenPoolAbstract.ParseAllowListRemove(log)
	case _FastTransferTokenPoolAbstract.abi.Events["Burned"].ID:
		return _FastTransferTokenPoolAbstract.ParseBurned(log)
	case _FastTransferTokenPoolAbstract.abi.Events["ChainAdded"].ID:
		return _FastTransferTokenPoolAbstract.ParseChainAdded(log)
	case _FastTransferTokenPoolAbstract.abi.Events["ChainConfigured"].ID:
		return _FastTransferTokenPoolAbstract.ParseChainConfigured(log)
	case _FastTransferTokenPoolAbstract.abi.Events["ChainRemoved"].ID:
		return _FastTransferTokenPoolAbstract.ParseChainRemoved(log)
	case _FastTransferTokenPoolAbstract.abi.Events["ConfigChanged"].ID:
		return _FastTransferTokenPoolAbstract.ParseConfigChanged(log)
	case _FastTransferTokenPoolAbstract.abi.Events["DestinationPoolUpdated"].ID:
		return _FastTransferTokenPoolAbstract.ParseDestinationPoolUpdated(log)
	case _FastTransferTokenPoolAbstract.abi.Events["FastFill"].ID:
		return _FastTransferTokenPoolAbstract.ParseFastFill(log)
	case _FastTransferTokenPoolAbstract.abi.Events["FastFillRequest"].ID:
		return _FastTransferTokenPoolAbstract.ParseFastFillRequest(log)
	case _FastTransferTokenPoolAbstract.abi.Events["FastFillSettled"].ID:
		return _FastTransferTokenPoolAbstract.ParseFastFillSettled(log)
	case _FastTransferTokenPoolAbstract.abi.Events["FillerAllowListUpdated"].ID:
		return _FastTransferTokenPoolAbstract.ParseFillerAllowListUpdated(log)
	case _FastTransferTokenPoolAbstract.abi.Events["InvalidFill"].ID:
		return _FastTransferTokenPoolAbstract.ParseInvalidFill(log)
	case _FastTransferTokenPoolAbstract.abi.Events["LaneUpdated"].ID:
		return _FastTransferTokenPoolAbstract.ParseLaneUpdated(log)
	case _FastTransferTokenPoolAbstract.abi.Events["Locked"].ID:
		return _FastTransferTokenPoolAbstract.ParseLocked(log)
	case _FastTransferTokenPoolAbstract.abi.Events["Minted"].ID:
		return _FastTransferTokenPoolAbstract.ParseMinted(log)
	case _FastTransferTokenPoolAbstract.abi.Events["OwnershipTransferRequested"].ID:
		return _FastTransferTokenPoolAbstract.ParseOwnershipTransferRequested(log)
	case _FastTransferTokenPoolAbstract.abi.Events["OwnershipTransferred"].ID:
		return _FastTransferTokenPoolAbstract.ParseOwnershipTransferred(log)
	case _FastTransferTokenPoolAbstract.abi.Events["RateLimitAdminSet"].ID:
		return _FastTransferTokenPoolAbstract.ParseRateLimitAdminSet(log)
	case _FastTransferTokenPoolAbstract.abi.Events["Released"].ID:
		return _FastTransferTokenPoolAbstract.ParseReleased(log)
	case _FastTransferTokenPoolAbstract.abi.Events["RemotePoolAdded"].ID:
		return _FastTransferTokenPoolAbstract.ParseRemotePoolAdded(log)
	case _FastTransferTokenPoolAbstract.abi.Events["RemotePoolRemoved"].ID:
		return _FastTransferTokenPoolAbstract.ParseRemotePoolRemoved(log)
	case _FastTransferTokenPoolAbstract.abi.Events["RouterUpdated"].ID:
		return _FastTransferTokenPoolAbstract.ParseRouterUpdated(log)
	case _FastTransferTokenPoolAbstract.abi.Events["TokensConsumed"].ID:
		return _FastTransferTokenPoolAbstract.ParseTokensConsumed(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (FastTransferTokenPoolAbstractAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (FastTransferTokenPoolAbstractAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (FastTransferTokenPoolAbstractBurned) Topic() common.Hash {
	return common.HexToHash("0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7")
}

func (FastTransferTokenPoolAbstractChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (FastTransferTokenPoolAbstractChainConfigured) Topic() common.Hash {
	return common.HexToHash("0x0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b")
}

func (FastTransferTokenPoolAbstractChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (FastTransferTokenPoolAbstractConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (FastTransferTokenPoolAbstractDestinationPoolUpdated) Topic() common.Hash {
	return common.HexToHash("0xb760e03fa04c0e86fcff6d0046cdcf22fb5d5b6a17d1e6f890b3456e81c40fd8")
}

func (FastTransferTokenPoolAbstractFastFill) Topic() common.Hash {
	return common.HexToHash("0x05abb4c63204ac64d430895792b9a147e544fbb4591f7c1bd1b4436dbf63355f")
}

func (FastTransferTokenPoolAbstractFastFillRequest) Topic() common.Hash {
	return common.HexToHash("0xbe96da9e73004694b4d32649f81f31784bd7c37daf64113f0ce1ebdd49a99a93")
}

func (FastTransferTokenPoolAbstractFastFillSettled) Topic() common.Hash {
	return common.HexToHash("0x9a3acb7ef95f9a8a2384f156d99ec2f035807c667ae66326823feead6d08fdbd")
}

func (FastTransferTokenPoolAbstractFillerAllowListUpdated) Topic() common.Hash {
	return common.HexToHash("0xccc0f2211c115acfa175a7923abdeb4b0a7c376d1b9e43c74973efe83d7d9e22")
}

func (FastTransferTokenPoolAbstractInvalidFill) Topic() common.Hash {
	return common.HexToHash("0xad64960fe1d28c88faed204b509e08b3c9e07d9c1cb84991addc205e6bfca42f")
}

func (FastTransferTokenPoolAbstractLaneUpdated) Topic() common.Hash {
	return common.HexToHash("0xbed278e7a8c5c763baafe3f3497295e65fd1dec8d51555c7b72c665e219d2deb")
}

func (FastTransferTokenPoolAbstractLocked) Topic() common.Hash {
	return common.HexToHash("0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008")
}

func (FastTransferTokenPoolAbstractMinted) Topic() common.Hash {
	return common.HexToHash("0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0")
}

func (FastTransferTokenPoolAbstractOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (FastTransferTokenPoolAbstractOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (FastTransferTokenPoolAbstractRateLimitAdminSet) Topic() common.Hash {
	return common.HexToHash("0x44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174")
}

func (FastTransferTokenPoolAbstractReleased) Topic() common.Hash {
	return common.HexToHash("0x2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52")
}

func (FastTransferTokenPoolAbstractRemotePoolAdded) Topic() common.Hash {
	return common.HexToHash("0x7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea")
}

func (FastTransferTokenPoolAbstractRemotePoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76")
}

func (FastTransferTokenPoolAbstractRouterUpdated) Topic() common.Hash {
	return common.HexToHash("0x02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f1684")
}

func (FastTransferTokenPoolAbstractTokensConsumed) Topic() common.Hash {
	return common.HexToHash("0x1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a")
}

func (_FastTransferTokenPoolAbstract *FastTransferTokenPoolAbstract) Address() common.Address {
	return _FastTransferTokenPoolAbstract.address
}

type FastTransferTokenPoolAbstractInterface interface {
	ComputeFillId(opts *bind.CallOpts, fillRequestId [32]byte, amount *big.Int, receiver common.Address) ([32]byte, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetCcipSendTokenFee(opts *bind.CallOpts, settlementFeeToken common.Address, destinationChainSelector uint64, amount *big.Int, receiver []byte, extraArgs []byte) (IFastTransferPoolQuote, error)

	GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetDestChainConfig(opts *bind.CallOpts, remoteChainSelector uint64) (FastTransferTokenPoolAbstractDestChainConfigView, error)

	GetFillInfo(opts *bind.CallOpts, fillId [32]byte) (FastTransferTokenPoolAbstractFillInfo, error)

	GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	GetSupportedChains(opts *bind.CallOpts) ([]uint64, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	GetTokenDecimals(opts *bind.CallOpts) (uint8, error)

	IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error)

	IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error)

	IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error)

	IsfillerAllowListed(opts *bind.CallOpts, remoteChainSelector uint64, filler common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error)

	ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error)

	CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error)

	CcipSendToken(opts *bind.TransactOpts, feeToken common.Address, destinationChainSelector uint64, amount *big.Int, receiver []byte, extraArgs []byte) (*types.Transaction, error)

	FastFill(opts *bind.TransactOpts, fillRequestId [32]byte, sourceChainSelector uint64, srcAmount *big.Int, srcDecimals uint8, receiver common.Address) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error)

	SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error)

	SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateDestChainConfig(opts *bind.TransactOpts, laneConfigArgs FastTransferTokenPoolAbstractDestChainConfigUpdateArgs) (*types.Transaction, error)

	UpdatefillerAllowList(opts *bind.TransactOpts, destinationChainSelector uint64, addFillers []common.Address, removeFillers []common.Address) (*types.Transaction, error)

	FilterAllowListAdd(opts *bind.FilterOpts) (*FastTransferTokenPoolAbstractAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*FastTransferTokenPoolAbstractAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*FastTransferTokenPoolAbstractAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*FastTransferTokenPoolAbstractAllowListRemove, error)

	FilterBurned(opts *bind.FilterOpts, sender []common.Address) (*FastTransferTokenPoolAbstractBurnedIterator, error)

	WatchBurned(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractBurned, sender []common.Address) (event.Subscription, error)

	ParseBurned(log types.Log) (*FastTransferTokenPoolAbstractBurned, error)

	FilterChainAdded(opts *bind.FilterOpts) (*FastTransferTokenPoolAbstractChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*FastTransferTokenPoolAbstractChainAdded, error)

	FilterChainConfigured(opts *bind.FilterOpts) (*FastTransferTokenPoolAbstractChainConfiguredIterator, error)

	WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractChainConfigured) (event.Subscription, error)

	ParseChainConfigured(log types.Log) (*FastTransferTokenPoolAbstractChainConfigured, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*FastTransferTokenPoolAbstractChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*FastTransferTokenPoolAbstractChainRemoved, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*FastTransferTokenPoolAbstractConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*FastTransferTokenPoolAbstractConfigChanged, error)

	FilterDestinationPoolUpdated(opts *bind.FilterOpts, dst []uint64) (*FastTransferTokenPoolAbstractDestinationPoolUpdatedIterator, error)

	WatchDestinationPoolUpdated(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractDestinationPoolUpdated, dst []uint64) (event.Subscription, error)

	ParseDestinationPoolUpdated(log types.Log) (*FastTransferTokenPoolAbstractDestinationPoolUpdated, error)

	FilterFastFill(opts *bind.FilterOpts, fillRequestId [][32]byte, fillId [][32]byte, filler []common.Address) (*FastTransferTokenPoolAbstractFastFillIterator, error)

	WatchFastFill(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractFastFill, fillRequestId [][32]byte, fillId [][32]byte, filler []common.Address) (event.Subscription, error)

	ParseFastFill(log types.Log) (*FastTransferTokenPoolAbstractFastFill, error)

	FilterFastFillRequest(opts *bind.FilterOpts, fillRequestId [][32]byte, dstChainSelector []uint64) (*FastTransferTokenPoolAbstractFastFillRequestIterator, error)

	WatchFastFillRequest(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractFastFillRequest, fillRequestId [][32]byte, dstChainSelector []uint64) (event.Subscription, error)

	ParseFastFillRequest(log types.Log) (*FastTransferTokenPoolAbstractFastFillRequest, error)

	FilterFastFillSettled(opts *bind.FilterOpts, fillRequestId [][32]byte) (*FastTransferTokenPoolAbstractFastFillSettledIterator, error)

	WatchFastFillSettled(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractFastFillSettled, fillRequestId [][32]byte) (event.Subscription, error)

	ParseFastFillSettled(log types.Log) (*FastTransferTokenPoolAbstractFastFillSettled, error)

	FilterFillerAllowListUpdated(opts *bind.FilterOpts, dst []uint64) (*FastTransferTokenPoolAbstractFillerAllowListUpdatedIterator, error)

	WatchFillerAllowListUpdated(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractFillerAllowListUpdated, dst []uint64) (event.Subscription, error)

	ParseFillerAllowListUpdated(log types.Log) (*FastTransferTokenPoolAbstractFillerAllowListUpdated, error)

	FilterInvalidFill(opts *bind.FilterOpts, fillRequestId [][32]byte, filler []common.Address) (*FastTransferTokenPoolAbstractInvalidFillIterator, error)

	WatchInvalidFill(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractInvalidFill, fillRequestId [][32]byte, filler []common.Address) (event.Subscription, error)

	ParseInvalidFill(log types.Log) (*FastTransferTokenPoolAbstractInvalidFill, error)

	FilterLaneUpdated(opts *bind.FilterOpts, destinationChainSelector []uint64) (*FastTransferTokenPoolAbstractLaneUpdatedIterator, error)

	WatchLaneUpdated(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractLaneUpdated, destinationChainSelector []uint64) (event.Subscription, error)

	ParseLaneUpdated(log types.Log) (*FastTransferTokenPoolAbstractLaneUpdated, error)

	FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*FastTransferTokenPoolAbstractLockedIterator, error)

	WatchLocked(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractLocked, sender []common.Address) (event.Subscription, error)

	ParseLocked(log types.Log) (*FastTransferTokenPoolAbstractLocked, error)

	FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*FastTransferTokenPoolAbstractMintedIterator, error)

	WatchMinted(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseMinted(log types.Log) (*FastTransferTokenPoolAbstractMinted, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FastTransferTokenPoolAbstractOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*FastTransferTokenPoolAbstractOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FastTransferTokenPoolAbstractOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*FastTransferTokenPoolAbstractOwnershipTransferred, error)

	FilterRateLimitAdminSet(opts *bind.FilterOpts) (*FastTransferTokenPoolAbstractRateLimitAdminSetIterator, error)

	WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractRateLimitAdminSet) (event.Subscription, error)

	ParseRateLimitAdminSet(log types.Log) (*FastTransferTokenPoolAbstractRateLimitAdminSet, error)

	FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*FastTransferTokenPoolAbstractReleasedIterator, error)

	WatchReleased(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseReleased(log types.Log) (*FastTransferTokenPoolAbstractReleased, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*FastTransferTokenPoolAbstractRemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*FastTransferTokenPoolAbstractRemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*FastTransferTokenPoolAbstractRemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*FastTransferTokenPoolAbstractRemotePoolRemoved, error)

	FilterRouterUpdated(opts *bind.FilterOpts) (*FastTransferTokenPoolAbstractRouterUpdatedIterator, error)

	WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractRouterUpdated) (event.Subscription, error)

	ParseRouterUpdated(log types.Log) (*FastTransferTokenPoolAbstractRouterUpdated, error)

	FilterTokensConsumed(opts *bind.FilterOpts) (*FastTransferTokenPoolAbstractTokensConsumedIterator, error)

	WatchTokensConsumed(opts *bind.WatchOpts, sink chan<- *FastTransferTokenPoolAbstractTokensConsumed) (event.Subscription, error)

	ParseTokensConsumed(log types.Log) (*FastTransferTokenPoolAbstractTokensConsumed, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
