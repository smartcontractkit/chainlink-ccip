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

var BurnMintFastTransferTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIBurnMintERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipReceive\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.Any2EVMMessage\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipSendToken\",\"inputs\":[{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"fillRequestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"computeFillId\",\"inputs\":[{\"name\":\"fillRequestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"fastFill\",\"inputs\":[{\"name\":\"fillRequestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"srcAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"srcDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCcipSendTokenFee\",\"inputs\":[{\"name\":\"settlementFeeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIFastTransferPool.Quote\",\"components\":[{\"name\":\"ccipSettlementFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fastTransferFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFastTransferTokenPoolAbstract.DestChainConfigView\",\"components\":[{\"name\":\"maxFillAmountPerRequest\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fastTransferBpsFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fillerAllowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"destinationPool\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFillInfo\",\"inputs\":[{\"name\":\"fillId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFastTransferTokenPoolAbstract.FillInfo\",\"components\":[{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"enumFastTransferTokenPoolAbstract.FillState\"},{\"name\":\"filler\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isfillerAllowListed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"filler\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRouter\",\"inputs\":[{\"name\":\"newRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateDestChainConfig\",\"inputs\":[{\"name\":\"laneConfigArgs\",\"type\":\"tuple\",\"internalType\":\"structFastTransferTokenPoolAbstract.DestChainConfigUpdateArgs\",\"components\":[{\"name\":\"maxFillAmountPerRequest\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"addFillers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removeFillers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fastTransferBpsFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fillerAllowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"destinationPool\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updatefillerAllowList\",\"inputs\":[{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addFillers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removeFillers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Burned\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestinationPoolUpdated\",\"inputs\":[{\"name\":\"dst\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"destinationPool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FastFill\",\"inputs\":[{\"name\":\"fillRequestId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"fillId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"filler\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"destAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FastFillRequest\",\"inputs\":[{\"name\":\"fillRequestId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"dstChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fastTransferFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FastFillSettled\",\"inputs\":[{\"name\":\"fillRequestId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FillerAllowListUpdated\",\"inputs\":[{\"name\":\"dst\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"addFillers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"removeFillers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvalidFill\",\"inputs\":[{\"name\":\"fillRequestId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"filler\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"filledAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"expectedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LaneUpdated\",\"inputs\":[{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"bps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"maxFillAmountPerRequest\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"destinationPool\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"addFillers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"removeFillers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Locked\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Minted\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Released\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RouterUpdated\",\"inputs\":[{\"name\":\"oldRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensConsumed\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AggregateValueMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"AggregateValueRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyFilled\",\"inputs\":[{\"name\":\"fillRequestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AlreadySettled\",\"inputs\":[{\"name\":\"fillRequestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"FillerNotWhitelisted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"filler\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRouter\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"LaneDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"RateLimitMustBeDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"WhitelistNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x61012080604052346103d057616ec0803803809161001d828561044f565b8339810160a0828203126103d05781516001600160a01b038116908190036103d05761004b60208401610472565b60408401519091906001600160401b0381116103d05784019280601f850112156103d0578351936001600160401b038511610439578460051b906020820195610097604051978861044f565b86526020808701928201019283116103d057602001905b828210610421575050506100d060806100c960608701610480565b9501610480565b93331561041057600180546001600160a01b03191633179055811580156103ff575b80156103ee575b6103dd578160209160049360805260c0526040519283809263313ce56760e01b82525afa6000918161039c575b50610371575b5060a052600480546001600160a01b0319166001600160a01b0384169081179091558151151560e081905290919061024f575b5015610239576101005260405161688b908161063582396080518181816123e401528181612ff3015281816131dc015281816139e201528181613bb601528181613ec101528181613f3901528181614075015281816157a00152818161585b0152615b48015260a051818181611d85015281816125d901528181612a8901528181613e480152818161526201526152e5015260c05181818161127601528181611f290152818161248001528181612efd0152613a7d015260e0518181816112060152818161378e0152615da3015261010051816120e40152f35b6335fdcccd60e21b600052600060045260246000fd5b9060209060405190610261838361044f565b60008252600036813760e051156103605760005b82518110156102dc576001906001600160a01b036102938286610494565b51168561029f826104d6565b6102ac575b505001610275565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a138856102a4565b5092905060005b8151811015610357576001906001600160a01b036103018285610494565b511680156103515784610313826105d4565b610321575b50505b016102e3565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a13884610318565b5061031b565b5050503861015f565b6335f4a7b360e01b60005260046000fd5b60ff1660ff8216818103610385575061012c565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d6020116103d5575b816103b86020938361044f565b810103126103d0576103c990610472565b9038610126565b600080fd5b3d91506103ab565b6342bcdf7f60e11b60005260046000fd5b506001600160a01b038116156100f9565b506001600160a01b038516156100f2565b639b15e16f60e01b60005260046000fd5b6020809161042e84610480565b8152019101906100ae565b634e487b7160e01b600052604160045260246000fd5b601f909101601f19168101906001600160401b0382119082101761043957604052565b519060ff821682036103d057565b51906001600160a01b03821682036103d057565b80518210156104a85760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b80548210156104a85760005260206000200190600090565b60008181526003602052604090205480156105cd5760001981018181116105b7576002546000198101919082116105b757818103610566575b5050506002548015610550576000190161052a8160026104be565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61059f6105776105889360026104be565b90549060031b1c92839260026104be565b819391549060031b91821b91600019901b19161790565b9055600052600360205260406000205538808061050f565b634e487b7160e01b600052601160045260246000fd5b5050600090565b8060005260036020526040600020541560001461062e57600254680100000000000000008110156104395761061561058882600185940160025560026104be565b9055600254906000526003602052604060002055600190565b5060009056fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a71461444e575080630a5fc2131461428e578063181f5a77146141e9578063211ffa0414613f5d57806321df0da714613eee578063240028e814613e6c57806324f65ee714613e10578063390775371461393c5780634c5ef0ed146138d757806354c8a4f31461375a57806362ddd3c4146136d65780636d3d1a58146136845780636def4ce71461359e57806379ba5097146134b95780637d54534e1461340c57806385572ffb14612c735780638926f54f14612c0f5780638d4470e4146129bd5780638da5cb5b1461296b578063962d4020146127c75780639a4575b91461239b578063a22f8f7e14611cb3578063a42a7b8b14611b2e578063a7cd63b714611a87578063abe1c1e8146119a3578063acfecf9114611885578063af58d59f1461181e578063b0f479a1146117cc578063b794658014611775578063c0d7865514611674578063c4bffe2b1461152b578063c75eea9c14611465578063cf7401f31461129a578063dc0bd9711461122b578063e0351e13146111d0578063e0c118b414611166578063e8a1da17146108bb578063f2fde38b146107ce578063f49570a6146107295763f5faf629146101d957600080fd5b346107265760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126107265760043567ffffffffffffffff81116106f55760e07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126106f55761024e61518b565b6040519160e0830183811067ffffffffffffffff8211176106f95760405281600401358352602482013567ffffffffffffffff81116106f5576102979060043691850101614822565b9060208401918252604483013567ffffffffffffffff81116106f5576102c39060043691860101614822565b90604085019182526102d760648501614661565b9060608601918252608485013561ffff811681036106f5576080870190815261030260a48701614bf1565b9560a0880196875260c481013567ffffffffffffffff81116106f15761032d91369101600401614917565b60c0880190815261271061ffff835116116106c95767ffffffffffffffff8451168352600a60205260408320968151600289019080519067ffffffffffffffff821161069c5761037d8354614d8f565b601f8111610661575b50602090601f83116001146105bf576103d492918891836105b4575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90559796975b61ffff8351169060018a01917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000062ff000084549351151560101b16921617179055865188556003839801975b855180518210156104a35790604073ffffffffffffffffffffffffffffffffffffffff61045583600195614c62565b511673ffffffffffffffffffffffffffffffffffffffff6000911681528b602052207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00815416905501610426565b5050868896845b815180518210156105295790604073ffffffffffffffffffffffffffffffffffffffff6104d983600195614c62565b511673ffffffffffffffffffffffffffffffffffffffff6000911681528b60205220827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055016104aa565b50506105ae906105a08861059261ffff67ffffffffffffffff7fbed278e7a8c5c763baafe3f3497295e65fd1dec8d51555c7b72c665e219d2deb999a9b51169951169551965193519151936040519788978852602088015260a0604088015260a087019061489d565b908582036060870152614ba7565b908382036080850152614ba7565b0390a280f35b0151905038806103a2565b83885281882091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416895b8181106106495750908460019594939210610612575b505050811b0190559796976103da565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388080610602565b929360206001819287860151815501950193016105ec565b61068c9084895260208920601f850160051c81019160208610610692575b601f0160051c0190615161565b38610386565b909150819061067f565b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b6004837f382c0982000000000000000000000000000000000000000000000000000000008152fd5b8380fd5b5080fd5b6024827f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b80fd5b50346107265760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126107265761076161464a565b906024359173ffffffffffffffffffffffffffffffffffffffff831683036106f557604073ffffffffffffffffffffffffffffffffffffffff9267ffffffffffffffff600393168152600a60205220019116600052602052602060ff604060002054166040519015158152f35b50346107265760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126107265773ffffffffffffffffffffffffffffffffffffffff61081b61478a565b61082361518b565b1633811461089357807fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b5034610726576108ca36614963565b939190926108d661518b565b82915b808310610fd1575050508063ffffffff4216917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee1843603015b85821015610fcd578160051b850135818112156106f157850190610120823603126106f15760405195610944876146f9565b61094d83614661565b8752602083013567ffffffffffffffff8111610fc95783019536601f88011215610fc95786359661097d88614772565b9761098b604051998a614731565b8089526020808a019160051b83010190368211610fc55760208301905b828210610f92575050505060208801968752604084013567ffffffffffffffff8111610f8e576109db9036908601614917565b9860408901998a52610a056109f33660608801614c1b565b9560608b0196875260c0369101614c1b565b9660808a01978852610a1786516159b6565b610a2188516159b6565b8a515115610f6657610a3d67ffffffffffffffff8b5116616703565b15610f2f5767ffffffffffffffff8a51168152600760205260408120610b7d87516fffffffffffffffffffffffffffffffff60408201511690610b386fffffffffffffffffffffffffffffffff60208301511691511515836080604051610aa3816146f9565b858152602081018c905260408101849052606081018690520152855474ff000000000000000000000000000000000000000091151560a01b919091167fffffffffffffffffffffff0000000000000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff84161773ffffffff0000000000000000000000000000000060808b901b1617178555565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176001830155565b610ca389516fffffffffffffffffffffffffffffffff60408201511690610c5e6fffffffffffffffffffffffffffffffff60208301511691511515836080604051610bc7816146f9565b858152602081018c9052604081018490526060810186905201526002860180547fffffffffffffffffffffff000000000000000000000000000000000000000000166fffffffffffffffffffffffffffffffff85161773ffffffff0000000000000000000000000000000060808c901b161791151560a01b74ff000000000000000000000000000000000000000016919091179055565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176003830155565b60048c5191019080519067ffffffffffffffff8211610f0257610cc68354614d8f565b601f8111610ed2575b50602090601f8311600114610e3357610d1c92918591836105b45750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b805b89518051821015610d575790610d51600192610d4a838f67ffffffffffffffff90511692614c62565b51906153ef565b01610d21565b5050975097987f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c292959396610e2567ffffffffffffffff600197949c5116925193519151610df1610dbc6040519687968752610100602088015261010087019061489d565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a1019093949291610912565b83855281852091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416865b818110610eba5750908460019594939210610e83575b505050811b019055610d1f565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388080610e76565b92936020600181928786015181550195019301610e60565b610efc9084865260208620601f850160051c8101916020861061069257601f0160051c0190615161565b38610ccf565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60249067ffffffffffffffff8b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b807f8579befe0000000000000000000000000000000000000000000000000000000060049252fd5b8680fd5b813567ffffffffffffffff8111610fc157602091610fb68392833691890101614917565b8152019101906109a8565b8a80fd5b8880fd5b8580fd5b8280f35b9092919367ffffffffffffffff610ff1610fec878588614ebd565b614ecd565b1695610ffc87616444565b1561113a5786845260076020526110186005604086200161624b565b94845b865181101561105157600190898752600760205261104a60056040892001611043838b614c62565b519061656f565b500161101b565b509394509490958085526007602052600560408620868155866001820155866002820155866003820155866004820161108a8154614d8f565b806110f9575b50505001805490868155816110db575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a10191909493946108d9565b865260208620908101905b818110156110a0578681556001016110e6565b601f811160011461110f5750555b863880611090565b8183526020832061112a91601f01861c810190600101615161565b8082528160208120915555611107565b602484887f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50346107265760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610726576044359073ffffffffffffffffffffffffffffffffffffffff821682036107265760206111c883602435600435615974565b604051908152f35b503461072657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126107265760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b503461072657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261072657602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346107265760e07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610726576112d261464a565b9060607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc3601126107265760405161130981614715565b60243580151581036114615781526044356fffffffffffffffffffffffffffffffff811681036114615760208201526064356fffffffffffffffffffffffffffffffff8116810361146157604082015260607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c3601126106f5576040519061139082614715565b60843580151581036106f157825260a4356fffffffffffffffffffffffffffffffff811681036106f157602083015260c4356fffffffffffffffffffffffffffffffff811681036106f157604083015273ffffffffffffffffffffffffffffffffffffffff600954163314158061143f575b611413576114109293615659565b80f35b6024837f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b5073ffffffffffffffffffffffffffffffffffffffff60015416331415611402565b8280fd5b50346107265760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610726576114ce6114c960406115279367ffffffffffffffff6114b261464a565b6114ba6150ae565b501681526007602052206150d9565b6158f5565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b0390f35b503461072657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261072657604051906005548083528260208101600584526020842092845b81811061165b57505061158992500383614731565b81516115ad61159782614772565b916115a56040519384614731565b808352614772565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602083019301368437805b845181101561160c578067ffffffffffffffff6115f960019388614c62565b51166116058286614c62565b52016115da565b50925090604051928392602084019060208552518091526040840192915b818110611638575050500390f35b825167ffffffffffffffff1684528594506020938401939092019160010161162a565b8454835260019485019487945060209093019201611574565b50346107265760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610726576116ac61478a565b6116b461518b565b73ffffffffffffffffffffffffffffffffffffffff811690811561174d57600480547fffffffffffffffffffffffff000000000000000000000000000000000000000081169390931790556040805173ffffffffffffffffffffffffffffffffffffffff93841681529190921660208201527f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f16849190a180f35b6004837f8579befe000000000000000000000000000000000000000000000000000000008152fd5b50346107265760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610726576115276117b86117b361464a565b61513f565b60405191829160208352602083019061489d565b503461072657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261072657602073ffffffffffffffffffffffffffffffffffffffff60045416604051908152f35b50346107265760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610726576114ce6114c9600260406115279467ffffffffffffffff61186d61464a565b6118756150ae565b50168152600760205220016150d9565b50346107265767ffffffffffffffff61189d36614a01565b9290916118a861518b565b16916118c1836000526006602052604060002054151590565b156119775782845260076020526118f0600560408620016118e33684866148e0565b602081519101209061656f565b1561192f57907f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76916105ae60405192839260208452602084019161506f565b82611973836040519384937f74f23c7c000000000000000000000000000000000000000000000000000000008552600485015260406024850152604484019161506f565b0390fd5b602484847f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50346107265760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610726576119db614ea4565b506004358152600b60205260408120604051906119f782614676565b5491611a0660ff841683614ca5565b73ffffffffffffffffffffffffffffffffffffffff602083019360081c1683526040519151906003821015611a5a575060409273ffffffffffffffffffffffffffffffffffffffff91835251166020820152f35b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526021600452fd5b503461072657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126107265760405160028054808352908352909160208301917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace915b818110611b185761152785611b0481870382614731565b604051918291602083526020830190614ba7565b8254845260209093019260019283019201611aed565b50346107265760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126107265767ffffffffffffffff611b6f61464a565b1681526007602052611b866005604083200161624b565b80517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611bcb611bb583614772565b92611bc36040519485614731565b808452614772565b01835b818110611ca2575050825b8251811015611c1f5780611bef60019285614c62565b5185526008602052611c0360408620614de2565b611c0d8285614c62565b52611c188184614c62565b5001611bd9565b81846040519182916020830160208452825180915260408401602060408360051b870101940192905b828210611c5757505050500390f35b91936020611c92827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc06001959799849503018652885161489d565b9601920192018594939192611c48565b806060602080938601015201611bce565b600182611cbf36614acf565b50509293604051611ccf81614676565b600081526000602082015260606080604051611cea816146f9565b8281528260208201528260408201528983820152015267ffffffffffffffff861695868852600a602052604088209173ffffffffffffffffffffffffffffffffffffffff612710611d4261ffff8d8701541689615178565b04948a6020840196808852611e04604051611d5c816146dd565b8b815260208101928352611d71368e8d6148e0565b6040820190815260ff611dd06060840192827f00000000000000000000000000000000000000000000000000000000000000001684526040519687956020808801525160408701525160608601525160808086015260c085019061489d565b91511660a0830152037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282614731565b60209d8e60405193611e168286614731565b845250612396575b611e36600260405198611e308a6146f9565b01614de2565b87528d870152604086015216806060850152604051611e558c82614731565b8a8152608085015273ffffffffffffffffffffffffffffffffffffffff600454168b60405180927f20487ded0000000000000000000000000000000000000000000000000000000082528180611eaf8a8a60048401614f2e565b03915afa90811561235e578b91612369575b5082526040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff000000000000000000000000000000008460801b1660048201528b8160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561235e578b91612331575b5061230957611f6733615da1565b611f7e896000526006602052604060002054151590565b156122dd57908291611f91888d95615827565b611fa4611f9f87518a614ef2565b615789565b806120d1575b5050916120009273ffffffffffffffffffffffffffffffffffffffff60045416906040518095819482937f96f4e9f900000000000000000000000000000000000000000000000000000000845260048401614f2e565b039134905af19687156120c5578097612069575b50509161205e8694927fbe96da9e73004694b4d32649f81f31784bd7c37daf64113f0ce1ebdd49a99a9394519360405194859485528a85015260606040850152606084019161506f565b0390a3604051908152f35b909194939296508782813d83116120be575b6120858183614731565b8101031261072657505194919290919061205e7fbe96da9e73004694b4d32649f81f31784bd7c37daf64113f0ce1ebdd49a99a93612014565b503d61207b565b604051903d90823e3d90fd5b909192506120e28251303384615881565b7f00000000000000000000000000000000000000000000000000000000000000009151918215801561223a575b156121b65761200094928c94926121a96121ae9361217d6040519485927f095ea7b3000000000000000000000000000000000000000000000000000000008b850152602484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101845283614731565b61610b565b90928b611faa565b60848c604051907f08c379a00000000000000000000000000000000000000000000000000000000082526004820152603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527f20746f206e6f6e2d7a65726f20616c6c6f77616e6365000000000000000000006064820152fd5b506040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff821660248201528c81604481865afa9081156122d2578c916122a0575b501561210f565b90508c81813d83116122cb575b6122b78183614731565b810103126122c657518d612299565b600080fd5b503d6122ad565b6040513d8e823e3d90fd5b60248a8a7fa9902c7e000000000000000000000000000000000000000000000000000000008252600452fd5b60048a7f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b61235191508c8d3d10612357575b6123498183614731565b8101906151d6565b8c611f59565b503d61233f565b6040513d8d823e3d90fd5b90508b81813d831161238f575b6123808183614731565b810103126122c657518c611ec1565b503d612376565b611e1e565b5034610726576123aa36614a60565b606060206040516123ba81614676565b8281520152608081016123cc81614d31565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691160361277d5750602081019177ffffffffffffffff0000000000000000000000000000000061243384614ecd565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156126ff57829161275e575b50612736576124c96124c460408401614d31565b615da1565b67ffffffffffffffff6124db84614ecd565b166124f3816000526006602052604060002054151590565b1561270a57602073ffffffffffffffffffffffffffffffffffffffff60045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa80156126ff57829061269c575b73ffffffffffffffffffffffffffffffffffffffff91501633036126705761263f6125cf6117b38585612595606061258b84614ecd565b9201358092615827565b61259e81615789565b6040519081527f696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df760203392a2614ecd565b61152760405160ff7f00000000000000000000000000000000000000000000000000000000000000001660208201526020815261260d604082614731565b6040519261261a84614676565b835260208301908152604051938493602085525160406020860152606085019061489d565b90517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe084830301604085015261489d565b807f728fe07b000000000000000000000000000000000000000000000000000000006024925233600452fd5b506020813d6020116126f7575b816126b660209383614731565b810103126106f5575173ffffffffffffffffffffffffffffffffffffffff811681036106f55773ffffffffffffffffffffffffffffffffffffffff90612554565b3d91506126a9565b6040513d84823e3d90fd5b602492507fa9902c7e000000000000000000000000000000000000000000000000000000008252600452fd5b807f53ad11d80000000000000000000000000000000000000000000000000000000060049252fd5b612777915060203d602011612357576123498183614731565b386124b0565b8273ffffffffffffffffffffffffffffffffffffffff61279e602493614d31565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835216600452fd5b50346107265760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126107265760043567ffffffffffffffff81116106f557612817903690600401614932565b60243567ffffffffffffffff81116106f157612837903690600401614b76565b60449291923567ffffffffffffffff8111610fc95761285a903690600401614b76565b91909273ffffffffffffffffffffffffffffffffffffffff6009541633141580612949575b61291d57818114801590612913575b6128eb57865b81811061289f578780f35b806128e56128b3610fec600194868c614ebd565b6128be83878b614ee2565b6128df6128d76128cf868b8d614ee2565b923690614c1b565b913690614c1b565b91615659565b01612894565b6004877f568efce2000000000000000000000000000000000000000000000000000000008152fd5b508281141561288e565b6024877f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b5073ffffffffffffffffffffffffffffffffffffffff6001541633141561287f565b503461072657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261072657602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b503461072657612bba906129d036614acf565b50509295916129dd614ea4565b50604051966129eb88614676565b600088526000602089015260606080604051612a06816146f9565b8281528260208201528260408201528983820152015267ffffffffffffffff81168752600a602052612b086040882092612a75612710612a4e61ffff60018801541684615178565b049560208c0198878a5260405193612a65856146dd565b84526020840197885236916148e0565b6040820190815260ff612ad46060840192827f00000000000000000000000000000000000000000000000000000000000000001684526040519889956020808801525160408701525160608601525160808086015260c085019061489d565b91511660a0830152037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101855284614731565b73ffffffffffffffffffffffffffffffffffffffff602096879460405190612b308783614731565b8a8252612b45600260405197611e30896146f9565b8652868601526040850152166060830152604051612b638482614731565b878152608083015273ffffffffffffffffffffffffffffffffffffffff60045416906040518095819482937f20487ded00000000000000000000000000000000000000000000000000000000845260048401614f2e565b03915afa9384156120c55793612bdd575b50826040945283519283525190820152f35b9392508184813d8311612c08575b612bf58183614731565b810103126122c657604093519293612bcb565b503d612beb565b50346107265760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610726576020612c6967ffffffffffffffff612c5561464a565b166000526006602052604060002054151590565b6040519015158152f35b503461072657612c8236614a60565b9073ffffffffffffffffffffffffffffffffffffffff6004541633036133e05760a08236031261072657604051612cb8816146f9565b82358152612cc860208401614661565b60208201908152604084013567ffffffffffffffff81116106f157612cf09036908601614917565b9360408301948552606081013567ffffffffffffffff81116133a457612d199036908301614917565b906060840191825260808101359067ffffffffffffffff8211610fc9570136601f820112156133a4578035612d4d81614772565b91612d5b6040519384614731565b81835260208084019260061b820101903682116133a057602001915b8183106133a85750505060808401525190815182016020810192602081830312610fc95760208101519067ffffffffffffffff8211610f8e570190608090829003126133a45760405190612dca826146dd565b6020810151825260408101519260208301938452606082015167ffffffffffffffff81116133a0576020908301019185601f840112156133a0578251612e0f81614840565b93612e1d6040519586614731565b81855260208501976020838301011161339c5760809291886020612e41930161487a565b83604086015201519060ff821691828103610fc5576060850152519273ffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff85169688519a519251965194519051906020811061336c575b50169377ffffffffffffffff00000000000000000000000000000000604051917f2cbc26bb00000000000000000000000000000000000000000000000000000000835260801b16600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115613361578991613342575b5061331a57612f3c8187614d52565b156132dc5750612f64612f6b92612f5e83612f588795896152e2565b926152e2565b90614ef2565b9388615974565b92838652600b602052604086209660405197612f8689614676565b54612f9460ff82168a614ca5565b73ffffffffffffffffffffffffffffffffffffffff60208a019160081c16815288519860038a10156132af5788998997989950156000146131635750505082612fdc91615b11565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691823b156106f1576040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff92909216600483015260248201529082908290604490829084905af180156126ff5761314e575b50505b6040519061308d82614676565b6002825260208201908482528452600b602052604084209151906003821015613121577fffffffffffffffffffffff00000000000000000000000000000000000000000060ff74ffffffffffffffffffffffffffffffffffffffff008554935160081b169316911617179055517f9a3acb7ef95f9a8a2384f156d99ec2f035807c667ae66326823feead6d08fdbd8280a280f35b6024857f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b8161315891614731565b61146157823861307d565b9250925095949350516003811015613282576002036131a857602486867fb196a44a000000000000000000000000000000000000000000000000000000008252600452fd5b859293945073ffffffffffffffffffffffffffffffffffffffff90511673ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691823b156106f1576040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff92909216600483015260248201529082908290604490829084905af180156126ff5761326d575b5050613080565b8161327791614731565b611461578238613266565b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b6024897f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b611973906040519182917f24eb47e500000000000000000000000000000000000000000000000000000000835260206004840152602483019061489d565b6004887f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b61335b915060203d602011612357576123498183614731565b38612f2d565b6040513d8b823e3d90fd5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1638612e94565b8980fd5b8780fd5b8480fd5b6040833603126133a057602060409182516133c281614676565b6133cb866147ad565b81528286013583820152815201920191612d77565b807fd7f73334000000000000000000000000000000000000000000000000000000006024925233600452fd5b50346107265760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610726577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174602073ffffffffffffffffffffffffffffffffffffffff61347c61478a565b61348461518b565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000006009541617600955604051908152a180f35b503461072657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261072657805473ffffffffffffffffffffffffffffffffffffffff81163303613576577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516825573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b50346107265760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126107265760409067ffffffffffffffff6135e261464a565b60608085516135f0816146dd565b85815285602082015285878201520152168152600a60205220600181015461152782549160405192613621846146dd565b83526136496002602085019561ffff8416875260ff604087019460101c161515845201614de2565b906060840191825261ffff604051958695602087525160208701525116604085015251151560608401525160808084015260a083019061489d565b503461072657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261072657602073ffffffffffffffffffffffffffffffffffffffff60095416604051908152f35b5034610726576136e536614a01565b6136f19392919361518b565b67ffffffffffffffff8216613713816000526006602052604060002054151590565b1561372f575061141092936137299136916148e0565b906153ef565b7f1e670e4b000000000000000000000000000000000000000000000000000000008452600452602483fd5b5034610726576137849061378c61377036614963565b959161377d93919361518b565b36916147ce565b9336916147ce565b7f0000000000000000000000000000000000000000000000000000000000000000156138af57815b8351811015613827578073ffffffffffffffffffffffffffffffffffffffff6137df60019387614c62565b51166137ea816162ae565b6137f6575b50016137b4565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a1386137ef565b5090805b82518110156138ab578073ffffffffffffffffffffffffffffffffffffffff61385660019386614c62565b511680156138a557613867816166a3565b613874575b505b0161382b565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a18461386c565b5061386e565b5080f35b6004827f35f4a7b3000000000000000000000000000000000000000000000000000000008152fd5b50346107265760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126107265761390f61464a565b906024359067ffffffffffffffff8211610726576020612c69846139363660048701614917565b90614d52565b50346107265760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126107265760043567ffffffffffffffff81116106f557806004016101007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc833603011261146157826040516139bc816146c1565b52608482016139ca81614d31565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603613def57506024820177ffffffffffffffff00000000000000000000000000000000613a3082614ecd565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115613d72578591613dd0575b50613da85767ffffffffffffffff613ac482614ecd565b16613adc816000526006602052604060002054151590565b15613d7d57602073ffffffffffffffffffffffffffffffffffffffff60045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115613d72578591613d53575b5015613d2757613b5381614ecd565b613b6f60a4850191613936613b688487614ce0565b36916148e0565b15613ce0575090613baa613ba5613b68613b8b613bb095614ecd565b93613b9b60648801358096615b11565b60c4870190614ce0565b6151ee565b906152e2565b906044017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1683613bf583614d31565b823b156106f5576040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91909116600482015260248101859052918290604490829084905af18015613cd55791602094613c839273ffffffffffffffffffffffffffffffffffffffff94613cc5575b5050614d31565b166040518281527f9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0843392a380604051613cbc816146c1565b52604051908152f35b81613ccf91614731565b38613c7c565b6040513d86823e3d90fd5b613cea9083614ce0565b6119736040519283927f24eb47e500000000000000000000000000000000000000000000000000000000845260206004850152602484019161506f565b6024847f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b613d6c915060203d602011612357576123498183614731565b38613b44565b6040513d87823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008552600452602484fd5b6004847f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b613de9915060203d602011612357576123498183614731565b38613aad565b8373ffffffffffffffffffffffffffffffffffffffff61279e602493614d31565b503461072657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261072657602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346107265760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261072657602090613ea761478a565b905073ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b503461072657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261072657602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346107265760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610726576024359067ffffffffffffffff82166004358184036114615760443560643560ff811681036133a4576084359073ffffffffffffffffffffffffffffffffffffffff821692838303610f8e57613fe4838287615974565b95868852600b6020526040882073ffffffffffffffffffffffffffffffffffffffff6040519161401383614676565b5461402160ff821684614ca5565b60081c1660208201525160038110156132af576141bd57808852600a6020526040882060ff600182015460101c16614172575b505091614069614072979892614099946152e2565b96878093615b11565b337f0000000000000000000000000000000000000000000000000000000000000000615881565b6040516140a581614676565b6001815260208101338152848752600b602052604087209151906003821015614145577fffffffffffffffffffffff00000000000000000000000000000000000000000060ff74ffffffffffffffffffffffffffffffffffffffff008554935160081b16931691161717905560405193845260208401527f05abb4c63204ac64d430895792b9a147e544fbb4591f7c1bd1b4436dbf63355f60403394a480f35b6024887f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b338952600301602052604088205460ff161561418e5780614054565b7f4248ccdc00000000000000000000000000000000000000000000000000000000885260045233602452604487fd5b602488877fcee81443000000000000000000000000000000000000000000000000000000008252600452fd5b503461072657807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610726575061152760405161422a606082614731565b602381527f4275726e4d696e74466173745472616e73666572546f6b656e506f6f6c20312e60208201527f362e310000000000000000000000000000000000000000000000000000000000604082015260405191829160208352602083019061489d565b50346107265760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610726576142c661464a565b60243567ffffffffffffffff8111611461576142e6903690600401614822565b60443567ffffffffffffffff81116106f1579167ffffffffffffffff61431185943690600401614822565b9161431a61518b565b1691828452600a60205260408420916003859301925b8251811015614397578073ffffffffffffffffffffffffffffffffffffffff8061435c60019487614c62565b51161687528460205260408720827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082541617905501614330565b509092845b8251811015614401578073ffffffffffffffffffffffffffffffffffffffff806143c860019487614c62565b511616875284602052604087207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541690550161439c565b507fccc0f2211c115acfa175a7923abdeb4b0a7c376d1b9e43c74973efe83d7d9e22614440856105ae8895604051938493604085526040850190614ba7565b908382036020850152614ba7565b9050346106f55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126106f5576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361146157602092507feeb51d2a000000000000000000000000000000000000000000000000000000008114908115908282614620575b83156145f6575b83156144f5575b50505015158152f35b9250906145cc575b81156145a2575b8115614578575b811561451b575b503880806144ec565b7f85572ffb0000000000000000000000000000000000000000000000000000000081149150811561454e575b5038614512565b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438614547565b7f01ffc9a7000000000000000000000000000000000000000000000000000000008114915061450b565b7f85572ffb0000000000000000000000000000000000000000000000000000000081149150614504565b7f181f5a7700000000000000000000000000000000000000000000000000000000811491506144fd565b7f85572ffb00000000000000000000000000000000000000000000000000000000821493506144e5565b7f01ffc9a700000000000000000000000000000000000000000000000000000000821493506144de565b6004359067ffffffffffffffff821682036122c657565b359067ffffffffffffffff821682036122c657565b6040810190811067ffffffffffffffff82111761469257604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff82111761469257604052565b6080810190811067ffffffffffffffff82111761469257604052565b60a0810190811067ffffffffffffffff82111761469257604052565b6060810190811067ffffffffffffffff82111761469257604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761469257604052565b67ffffffffffffffff81116146925760051b60200190565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036122c657565b359073ffffffffffffffffffffffffffffffffffffffff821682036122c657565b9291906147da81614772565b936147e86040519586614731565b602085838152019160051b81019283116122c657905b82821061480a57505050565b60208091614817846147ad565b8152019101906147fe565b9080601f830112156122c65781602061483d933591016147ce565b90565b67ffffffffffffffff811161469257601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60005b83811061488d5750506000910152565b818101518382015260200161487d565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f6020936148d98151809281875287808801910161487a565b0116010190565b9291926148ec82614840565b916148fa6040519384614731565b8294818452818301116122c6578281602093846000960137010152565b9080601f830112156122c65781602061483d933591016148e0565b9181601f840112156122c65782359167ffffffffffffffff83116122c6576020808501948460051b0101116122c657565b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126122c65760043567ffffffffffffffff81116122c657816149ac91600401614932565b929092916024359067ffffffffffffffff82116122c6576149cf91600401614932565b9091565b9181601f840112156122c65782359167ffffffffffffffff83116122c657602083818601950101116122c657565b9060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126122c65760043567ffffffffffffffff811681036122c657916024359067ffffffffffffffff82116122c6576149cf916004016149d3565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126122c6576004359067ffffffffffffffff82116122c6577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8260a0920301126122c65760040190565b9060a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126122c65760043573ffffffffffffffffffffffffffffffffffffffff811681036122c6579160243567ffffffffffffffff811681036122c657916044359160643567ffffffffffffffff81116122c65781614b53916004016149d3565b929092916084359067ffffffffffffffff82116122c6576149cf916004016149d3565b9181601f840112156122c65782359167ffffffffffffffff83116122c657602080850194606085020101116122c657565b906020808351928381520192019060005b818110614bc55750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101614bb8565b359081151582036122c657565b35906fffffffffffffffffffffffffffffffff821682036122c657565b91908260609103126122c657604051614c3381614715565b6040614c5d818395614c4481614bf1565b8552614c5260208201614bfe565b602086015201614bfe565b910152565b8051821015614c765760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6003821015614cb15752565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156122c6570180359067ffffffffffffffff82116122c6576020019181360383136122c657565b3573ffffffffffffffffffffffffffffffffffffffff811681036122c65790565b9067ffffffffffffffff61483d92166000526007602052600560406000200190602081519101209060019160005201602052604060002054151590565b90600182811c92168015614dd8575b6020831014614da957565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691614d9e565b9060405191826000825492614df684614d8f565b8084529360018116908115614e645750600114614e1d575b50614e1b92500383614731565b565b90506000929192526020600020906000915b818310614e48575050906020614e1b9282010138614e0e565b6020919350806001915483858901015201910190918492614e2f565b60209350614e1b9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138614e0e565b60405190614eb182614676565b60006020838281520152565b9190811015614c765760051b0190565b3567ffffffffffffffff811681036122c65790565b9190811015614c76576060020190565b91908201809211614eff57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9067ffffffffffffffff9093929316815260406020820152614f93614f5f845160a0604085015260e084019061489d565b60208501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc084830301606085015261489d565b906040840151917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08282030160808301526020808451928381520193019060005b8181106150375750505060808473ffffffffffffffffffffffffffffffffffffffff606061483d969701511660a084015201519060c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08285030191015261489d565b8251805173ffffffffffffffffffffffffffffffffffffffff1686526020908101518187015260409095019490920191600101614fd4565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b604051906150bb826146f9565b60006080838281528260208201528260408201528260608201520152565b906040516150e6816146f9565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff16600052600760205261483d6004604060002001614de2565b81811061516c575050565b60008155600101615161565b81810292918115918404141715614eff57565b73ffffffffffffffffffffffffffffffffffffffff6001541633036151ac57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b908160209103126122c6575180151581036122c65790565b8051801561525e576020036152205780516020828101918301839003126122c657519060ff8211615220575060ff1690565b611973906040519182917f953576f700000000000000000000000000000000000000000000000000000000835260206004840152602483019061489d565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff8211614eff57565b60ff16604d8111614eff57600a0a90565b81156152b3570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff8116928284146153e8578284116153be579061532791615284565b91604d60ff8416118015615385575b61534f5750509061534961483d92615298565b90615178565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b5061538f83615298565b80156152b3577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048411615336565b6153c791615284565b91604d60ff84161161534f575050906153e261483d92615298565b906152a9565b5050505090565b9080511561562f5767ffffffffffffffff8151602083012092169182600052600760205261542481600560406000200161675d565b156155eb5760005260086020526040600020815167ffffffffffffffff8111614692576154518254614d8f565b601f81116155b9575b506020601f82116001146154f357916154cd827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea95936154e3956000916154e8575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b905560405191829160208352602083019061489d565b0390a2565b90508401513861549c565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b8181106155a15750926154e39492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea98961061556a575b5050811b0190556117b8565b8501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388061555e565b9192602060018192868a015181550194019201615523565b6155e590836000526020600020601f840160051c8101916020851061069257601f0160051c0190615161565b3861545a565b50906119736040519283927f393b8ad2000000000000000000000000000000000000000000000000000000008452600484015260406024840152604483019061489d565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b67ffffffffffffffff16600081815260066020526040902054909291901561575b579161575860e092615724856156b07f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b976159b6565b8460005260076020526156c7816040600020615b6e565b6156d0836159b6565b8460005260076020526156ea836002604060002001615b6e565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000006157cb83303384615881565b1690813b156122c657604051907f42966c680000000000000000000000000000000000000000000000000000000082528160248160008096819560048401525af180156126ff5761581a575050565b8161582491614731565b50565b9067ffffffffffffffff614e1b9216600052600760205260406000209073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691615e23565b90919273ffffffffffffffffffffffffffffffffffffffff614e1b9481604051957f23b872dd0000000000000000000000000000000000000000000000000000000060208801521660248601521660448401526064830152606482526121a9608483614731565b91908203918211614eff57565b6158fd6150ae565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff80835116916159546020850193612f5e61594763ffffffff875116426158e8565b8560808901511690615178565b8082101561596d57505b16825263ffffffff4216905290565b905061595e565b9173ffffffffffffffffffffffffffffffffffffffff9060405192602084019485526040840152166060820152606081526159b0608082614731565b51902090565b805115615a6a576fffffffffffffffffffffffffffffffff6040820151166fffffffffffffffffffffffffffffffff602083015116811090811591615a61575b506159fe5750565b606490615a5f604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b905015386159f6565b6fffffffffffffffffffffffffffffffff60408201511615801590615af2575b615a915750565b606490615a5f604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff6020820151161515615a8a565b9067ffffffffffffffff614e1b9216600052600760205260026040600020019073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691615e23565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1991615ca76060928054615bab63ffffffff8260801c16426158e8565b9081615ce6575b50506fffffffffffffffffffffffffffffffff6001816020860151169282815416808510600014615cde57508280855b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416178155615c5b8651151582907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b60408601517fffffffffffffffffffffffffffffffff0000000000000000000000000000000060809190911b16939092166fffffffffffffffffffffffffffffffff1692909217910155565b61575860405180926fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b838091615be2565b6fffffffffffffffffffffffffffffffff91615d1b839283615d146001880154948286169560801c90615178565b9116614ef2565b80821015615d9a57505b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff9290911692909216167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116174260801b73ffffffff00000000000000000000000000000000161781553880615bb2565b9050615d25565b7f0000000000000000000000000000000000000000000000000000000000000000615dc95750565b73ffffffffffffffffffffffffffffffffffffffff1680600052600360205260406000205415615df65750565b7fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b929192805460ff8160a01c16158015616103575b6160fc576fffffffffffffffffffffffffffffffff81169060018301908154615e7c63ffffffff6fffffffffffffffffffffffffffffffff83169360801c16426158e8565b908161605e575b5050848110615fdc5750838210615f0b57507f1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a939450906fffffffffffffffffffffffffffffffff80615ed985602096956158e8565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055604051908152a1565b819450615f1d92505460801c926158e8565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810190808211614eff57615f6b615f709273ffffffffffffffffffffffffffffffffffffffff94614ef2565b6152a9565b9216918215615fac577fd0c8d23a0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b7f15279c080000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b8473ffffffffffffffffffffffffffffffffffffffff881691821561602e577f1a76572a0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b7ff94ebcd10000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b8285929395116160d25761607992612f5e9160801c90615178565b808310156160cd5750815b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178455913880615e83565b616084565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050509050565b508215615e37565b73ffffffffffffffffffffffffffffffffffffffff61619a9116916040926000808551936161398786614731565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602086015260208151910182855af13d15616243573d9161617e83614840565b9261618b87519485614731565b83523d6000602085013e6167b2565b805190816161a757505050565b6020806161b89383010191016151d6565b156161c05750565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b6060916167b2565b906040519182815491828252602082019060005260206000209260005b81811061627d575050614e1b92500383614731565b8454835260019485019487945060209093019201616268565b8054821015614c765760005260206000200190600090565b600081815260036020526040902054801561643d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111614eff57600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211614eff578181036163ce575b505050600254801561639f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161635c816002616296565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6164256163df6163f0936002616296565b90549060031b1c9283926002616296565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526003602052604060002055388080616323565b5050600090565b600081815260066020526040902054801561643d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111614eff57600554907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211614eff57818103616535575b505050600554801561639f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016164f2816005616296565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600555600052600660205260006040812055600190565b6165576165466163f0936005616296565b90549060031b1c9283926005616296565b905560005260066020526040600020553880806164b9565b906001820191816000528260205260406000205480151560001461669a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111614eff578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211614eff57818103616663575b5050508054801561639f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906166248282616296565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b6166836166736163f09386616296565b90549060031b1c92839286616296565b9055600052836020526040600020553880806165ec565b50505050600090565b806000526003602052604060002054156000146166fd5760025468010000000000000000811015614692576166e46163f08260018594016002556002616296565b9055600254906000526003602052604060002055600190565b50600090565b806000526006602052604060002054156000146166fd5760055468010000000000000000811015614692576167446163f08260018594016005556005616296565b9055600554906000526006602052604060002055600190565b600082815260018201602052604090205461643d5780549068010000000000000000821015614692578261679b6163f0846001809601855584616296565b905580549260005201602052604060002055600190565b9192901561682d57508151156167c6575090565b3b156167cf5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156168405750805190602001fd5b611973906040519182917f08c379a000000000000000000000000000000000000000000000000000000000835260206004840152602483019061489d56fea164736f6c634300081a000a",
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) ComputeFillId(opts *bind.CallOpts, fillRequestId [32]byte, amount *big.Int, receiver common.Address) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "computeFillId", fillRequestId, amount, receiver)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) ComputeFillId(fillRequestId [32]byte, amount *big.Int, receiver common.Address) ([32]byte, error) {
	return _BurnMintFastTransferTokenPool.Contract.ComputeFillId(&_BurnMintFastTransferTokenPool.CallOpts, fillRequestId, amount, receiver)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) ComputeFillId(fillRequestId [32]byte, amount *big.Int, receiver common.Address) ([32]byte, error) {
	return _BurnMintFastTransferTokenPool.Contract.ComputeFillId(&_BurnMintFastTransferTokenPool.CallOpts, fillRequestId, amount, receiver)
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) GetCcipSendTokenFee(opts *bind.CallOpts, settlementFeeToken common.Address, destinationChainSelector uint64, amount *big.Int, receiver []byte, extraArgs []byte) (IFastTransferPoolQuote, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "getCcipSendTokenFee", settlementFeeToken, destinationChainSelector, amount, receiver, extraArgs)

	if err != nil {
		return *new(IFastTransferPoolQuote), err
	}

	out0 := *abi.ConvertType(out[0], new(IFastTransferPoolQuote)).(*IFastTransferPoolQuote)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) GetCcipSendTokenFee(settlementFeeToken common.Address, destinationChainSelector uint64, amount *big.Int, receiver []byte, extraArgs []byte) (IFastTransferPoolQuote, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetCcipSendTokenFee(&_BurnMintFastTransferTokenPool.CallOpts, settlementFeeToken, destinationChainSelector, amount, receiver, extraArgs)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) GetCcipSendTokenFee(settlementFeeToken common.Address, destinationChainSelector uint64, amount *big.Int, receiver []byte, extraArgs []byte) (IFastTransferPoolQuote, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetCcipSendTokenFee(&_BurnMintFastTransferTokenPool.CallOpts, settlementFeeToken, destinationChainSelector, amount, receiver, extraArgs)
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) GetDestChainConfig(opts *bind.CallOpts, remoteChainSelector uint64) (FastTransferTokenPoolAbstractDestChainConfigView, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "getDestChainConfig", remoteChainSelector)

	if err != nil {
		return *new(FastTransferTokenPoolAbstractDestChainConfigView), err
	}

	out0 := *abi.ConvertType(out[0], new(FastTransferTokenPoolAbstractDestChainConfigView)).(*FastTransferTokenPoolAbstractDestChainConfigView)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) GetDestChainConfig(remoteChainSelector uint64) (FastTransferTokenPoolAbstractDestChainConfigView, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetDestChainConfig(&_BurnMintFastTransferTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) GetDestChainConfig(remoteChainSelector uint64) (FastTransferTokenPoolAbstractDestChainConfigView, error) {
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) IsfillerAllowListed(opts *bind.CallOpts, remoteChainSelector uint64, filler common.Address) (bool, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "isfillerAllowListed", remoteChainSelector, filler)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) IsfillerAllowListed(remoteChainSelector uint64, filler common.Address) (bool, error) {
	return _BurnMintFastTransferTokenPool.Contract.IsfillerAllowListed(&_BurnMintFastTransferTokenPool.CallOpts, remoteChainSelector, filler)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) IsfillerAllowListed(remoteChainSelector uint64, filler common.Address) (bool, error) {
	return _BurnMintFastTransferTokenPool.Contract.IsfillerAllowListed(&_BurnMintFastTransferTokenPool.CallOpts, remoteChainSelector, filler)
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactor) CcipSendToken(opts *bind.TransactOpts, feeToken common.Address, destinationChainSelector uint64, amount *big.Int, receiver []byte, extraArgs []byte) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.contract.Transact(opts, "ccipSendToken", feeToken, destinationChainSelector, amount, receiver, extraArgs)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) CcipSendToken(feeToken common.Address, destinationChainSelector uint64, amount *big.Int, receiver []byte, extraArgs []byte) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.CcipSendToken(&_BurnMintFastTransferTokenPool.TransactOpts, feeToken, destinationChainSelector, amount, receiver, extraArgs)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorSession) CcipSendToken(feeToken common.Address, destinationChainSelector uint64, amount *big.Int, receiver []byte, extraArgs []byte) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.CcipSendToken(&_BurnMintFastTransferTokenPool.TransactOpts, feeToken, destinationChainSelector, amount, receiver, extraArgs)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactor) FastFill(opts *bind.TransactOpts, fillRequestId [32]byte, sourceChainSelector uint64, srcAmount *big.Int, srcDecimals uint8, receiver common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.contract.Transact(opts, "fastFill", fillRequestId, sourceChainSelector, srcAmount, srcDecimals, receiver)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) FastFill(fillRequestId [32]byte, sourceChainSelector uint64, srcAmount *big.Int, srcDecimals uint8, receiver common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.FastFill(&_BurnMintFastTransferTokenPool.TransactOpts, fillRequestId, sourceChainSelector, srcAmount, srcDecimals, receiver)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorSession) FastFill(fillRequestId [32]byte, sourceChainSelector uint64, srcAmount *big.Int, srcDecimals uint8, receiver common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.FastFill(&_BurnMintFastTransferTokenPool.TransactOpts, fillRequestId, sourceChainSelector, srcAmount, srcDecimals, receiver)
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactor) UpdateDestChainConfig(opts *bind.TransactOpts, laneConfigArgs FastTransferTokenPoolAbstractDestChainConfigUpdateArgs) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.contract.Transact(opts, "updateDestChainConfig", laneConfigArgs)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) UpdateDestChainConfig(laneConfigArgs FastTransferTokenPoolAbstractDestChainConfigUpdateArgs) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.UpdateDestChainConfig(&_BurnMintFastTransferTokenPool.TransactOpts, laneConfigArgs)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorSession) UpdateDestChainConfig(laneConfigArgs FastTransferTokenPoolAbstractDestChainConfigUpdateArgs) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.UpdateDestChainConfig(&_BurnMintFastTransferTokenPool.TransactOpts, laneConfigArgs)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactor) UpdatefillerAllowList(opts *bind.TransactOpts, destinationChainSelector uint64, addFillers []common.Address, removeFillers []common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.contract.Transact(opts, "updatefillerAllowList", destinationChainSelector, addFillers, removeFillers)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) UpdatefillerAllowList(destinationChainSelector uint64, addFillers []common.Address, removeFillers []common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.UpdatefillerAllowList(&_BurnMintFastTransferTokenPool.TransactOpts, destinationChainSelector, addFillers, removeFillers)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorSession) UpdatefillerAllowList(destinationChainSelector uint64, addFillers []common.Address, removeFillers []common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.UpdatefillerAllowList(&_BurnMintFastTransferTokenPool.TransactOpts, destinationChainSelector, addFillers, removeFillers)
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

type BurnMintFastTransferTokenPoolBurnedIterator struct {
	Event *BurnMintFastTransferTokenPoolBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolBurned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintFastTransferTokenPoolBurned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintFastTransferTokenPoolBurnedIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolBurned struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterBurned(opts *bind.FilterOpts, sender []common.Address) (*BurnMintFastTransferTokenPoolBurnedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "Burned", senderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolBurnedIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "Burned", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolBurned, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "Burned", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolBurned)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "Burned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseBurned(log types.Log) (*BurnMintFastTransferTokenPoolBurned, error) {
	event := new(BurnMintFastTransferTokenPoolBurned)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "Burned", log); err != nil {
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
	Dst             uint64
	DestinationPool common.Address
	Raw             types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterDestinationPoolUpdated(opts *bind.FilterOpts, dst []uint64) (*BurnMintFastTransferTokenPoolDestinationPoolUpdatedIterator, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "DestinationPoolUpdated", dstRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolDestinationPoolUpdatedIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "DestinationPoolUpdated", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchDestinationPoolUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolDestinationPoolUpdated, dst []uint64) (event.Subscription, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "DestinationPoolUpdated", dstRule)
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

type BurnMintFastTransferTokenPoolFastFillIterator struct {
	Event *BurnMintFastTransferTokenPoolFastFill

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolFastFillIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolFastFill)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintFastTransferTokenPoolFastFill)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintFastTransferTokenPoolFastFillIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolFastFillIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolFastFill struct {
	FillRequestId [32]byte
	FillId        [32]byte
	Filler        common.Address
	DestAmount    *big.Int
	Receiver      common.Address
	Raw           types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterFastFill(opts *bind.FilterOpts, fillRequestId [][32]byte, fillId [][32]byte, filler []common.Address) (*BurnMintFastTransferTokenPoolFastFillIterator, error) {

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

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "FastFill", fillRequestIdRule, fillIdRule, fillerRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolFastFillIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "FastFill", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchFastFill(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolFastFill, fillRequestId [][32]byte, fillId [][32]byte, filler []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "FastFill", fillRequestIdRule, fillIdRule, fillerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolFastFill)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "FastFill", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseFastFill(log types.Log) (*BurnMintFastTransferTokenPoolFastFill, error) {
	event := new(BurnMintFastTransferTokenPoolFastFill)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "FastFill", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintFastTransferTokenPoolFastFillRequestIterator struct {
	Event *BurnMintFastTransferTokenPoolFastFillRequest

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolFastFillRequestIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolFastFillRequest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintFastTransferTokenPoolFastFillRequest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintFastTransferTokenPoolFastFillRequestIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolFastFillRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolFastFillRequest struct {
	FillRequestId    [32]byte
	DstChainSelector uint64
	Amount           *big.Int
	FastTransferFee  *big.Int
	Receiver         []byte
	Raw              types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterFastFillRequest(opts *bind.FilterOpts, fillRequestId [][32]byte, dstChainSelector []uint64) (*BurnMintFastTransferTokenPoolFastFillRequestIterator, error) {

	var fillRequestIdRule []interface{}
	for _, fillRequestIdItem := range fillRequestId {
		fillRequestIdRule = append(fillRequestIdRule, fillRequestIdItem)
	}
	var dstChainSelectorRule []interface{}
	for _, dstChainSelectorItem := range dstChainSelector {
		dstChainSelectorRule = append(dstChainSelectorRule, dstChainSelectorItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "FastFillRequest", fillRequestIdRule, dstChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolFastFillRequestIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "FastFillRequest", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchFastFillRequest(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolFastFillRequest, fillRequestId [][32]byte, dstChainSelector []uint64) (event.Subscription, error) {

	var fillRequestIdRule []interface{}
	for _, fillRequestIdItem := range fillRequestId {
		fillRequestIdRule = append(fillRequestIdRule, fillRequestIdItem)
	}
	var dstChainSelectorRule []interface{}
	for _, dstChainSelectorItem := range dstChainSelector {
		dstChainSelectorRule = append(dstChainSelectorRule, dstChainSelectorItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "FastFillRequest", fillRequestIdRule, dstChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolFastFillRequest)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "FastFillRequest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseFastFillRequest(log types.Log) (*BurnMintFastTransferTokenPoolFastFillRequest, error) {
	event := new(BurnMintFastTransferTokenPoolFastFillRequest)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "FastFillRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintFastTransferTokenPoolFastFillSettledIterator struct {
	Event *BurnMintFastTransferTokenPoolFastFillSettled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolFastFillSettledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolFastFillSettled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintFastTransferTokenPoolFastFillSettled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintFastTransferTokenPoolFastFillSettledIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolFastFillSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolFastFillSettled struct {
	FillRequestId [32]byte
	Raw           types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterFastFillSettled(opts *bind.FilterOpts, fillRequestId [][32]byte) (*BurnMintFastTransferTokenPoolFastFillSettledIterator, error) {

	var fillRequestIdRule []interface{}
	for _, fillRequestIdItem := range fillRequestId {
		fillRequestIdRule = append(fillRequestIdRule, fillRequestIdItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "FastFillSettled", fillRequestIdRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolFastFillSettledIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "FastFillSettled", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchFastFillSettled(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolFastFillSettled, fillRequestId [][32]byte) (event.Subscription, error) {

	var fillRequestIdRule []interface{}
	for _, fillRequestIdItem := range fillRequestId {
		fillRequestIdRule = append(fillRequestIdRule, fillRequestIdItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "FastFillSettled", fillRequestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolFastFillSettled)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "FastFillSettled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseFastFillSettled(log types.Log) (*BurnMintFastTransferTokenPoolFastFillSettled, error) {
	event := new(BurnMintFastTransferTokenPoolFastFillSettled)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "FastFillSettled", log); err != nil {
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
	Dst           uint64
	AddFillers    []common.Address
	RemoveFillers []common.Address
	Raw           types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterFillerAllowListUpdated(opts *bind.FilterOpts, dst []uint64) (*BurnMintFastTransferTokenPoolFillerAllowListUpdatedIterator, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "FillerAllowListUpdated", dstRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolFillerAllowListUpdatedIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "FillerAllowListUpdated", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchFillerAllowListUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolFillerAllowListUpdated, dst []uint64) (event.Subscription, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "FillerAllowListUpdated", dstRule)
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

type BurnMintFastTransferTokenPoolInvalidFillIterator struct {
	Event *BurnMintFastTransferTokenPoolInvalidFill

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolInvalidFillIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolInvalidFill)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintFastTransferTokenPoolInvalidFill)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintFastTransferTokenPoolInvalidFillIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolInvalidFillIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolInvalidFill struct {
	FillRequestId  [32]byte
	Filler         common.Address
	FilledAmount   *big.Int
	ExpectedAmount *big.Int
	Raw            types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterInvalidFill(opts *bind.FilterOpts, fillRequestId [][32]byte, filler []common.Address) (*BurnMintFastTransferTokenPoolInvalidFillIterator, error) {

	var fillRequestIdRule []interface{}
	for _, fillRequestIdItem := range fillRequestId {
		fillRequestIdRule = append(fillRequestIdRule, fillRequestIdItem)
	}
	var fillerRule []interface{}
	for _, fillerItem := range filler {
		fillerRule = append(fillerRule, fillerItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "InvalidFill", fillRequestIdRule, fillerRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolInvalidFillIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "InvalidFill", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchInvalidFill(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolInvalidFill, fillRequestId [][32]byte, filler []common.Address) (event.Subscription, error) {

	var fillRequestIdRule []interface{}
	for _, fillRequestIdItem := range fillRequestId {
		fillRequestIdRule = append(fillRequestIdRule, fillRequestIdItem)
	}
	var fillerRule []interface{}
	for _, fillerItem := range filler {
		fillerRule = append(fillerRule, fillerItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "InvalidFill", fillRequestIdRule, fillerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolInvalidFill)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "InvalidFill", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseInvalidFill(log types.Log) (*BurnMintFastTransferTokenPoolInvalidFill, error) {
	event := new(BurnMintFastTransferTokenPoolInvalidFill)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "InvalidFill", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintFastTransferTokenPoolLaneUpdatedIterator struct {
	Event *BurnMintFastTransferTokenPoolLaneUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolLaneUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolLaneUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintFastTransferTokenPoolLaneUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintFastTransferTokenPoolLaneUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolLaneUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolLaneUpdated struct {
	DestinationChainSelector uint64
	Bps                      uint16
	MaxFillAmountPerRequest  *big.Int
	DestinationPool          []byte
	AddFillers               []common.Address
	RemoveFillers            []common.Address
	Raw                      types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterLaneUpdated(opts *bind.FilterOpts, destinationChainSelector []uint64) (*BurnMintFastTransferTokenPoolLaneUpdatedIterator, error) {

	var destinationChainSelectorRule []interface{}
	for _, destinationChainSelectorItem := range destinationChainSelector {
		destinationChainSelectorRule = append(destinationChainSelectorRule, destinationChainSelectorItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "LaneUpdated", destinationChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolLaneUpdatedIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "LaneUpdated", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchLaneUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolLaneUpdated, destinationChainSelector []uint64) (event.Subscription, error) {

	var destinationChainSelectorRule []interface{}
	for _, destinationChainSelectorItem := range destinationChainSelector {
		destinationChainSelectorRule = append(destinationChainSelectorRule, destinationChainSelectorItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "LaneUpdated", destinationChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolLaneUpdated)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "LaneUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseLaneUpdated(log types.Log) (*BurnMintFastTransferTokenPoolLaneUpdated, error) {
	event := new(BurnMintFastTransferTokenPoolLaneUpdated)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "LaneUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintFastTransferTokenPoolLockedIterator struct {
	Event *BurnMintFastTransferTokenPoolLocked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolLockedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolLocked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintFastTransferTokenPoolLocked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintFastTransferTokenPoolLockedIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolLocked struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*BurnMintFastTransferTokenPoolLockedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolLockedIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "Locked", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolLocked, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolLocked)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "Locked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseLocked(log types.Log) (*BurnMintFastTransferTokenPoolLocked, error) {
	event := new(BurnMintFastTransferTokenPoolLocked)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintFastTransferTokenPoolMintedIterator struct {
	Event *BurnMintFastTransferTokenPoolMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintFastTransferTokenPoolMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintFastTransferTokenPoolMintedIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolMinted struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*BurnMintFastTransferTokenPoolMintedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolMintedIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "Minted", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolMinted)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "Minted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseMinted(log types.Log) (*BurnMintFastTransferTokenPoolMinted, error) {
	event := new(BurnMintFastTransferTokenPoolMinted)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "Minted", log); err != nil {
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

type BurnMintFastTransferTokenPoolReleasedIterator struct {
	Event *BurnMintFastTransferTokenPoolReleased

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolReleasedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolReleased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintFastTransferTokenPoolReleased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintFastTransferTokenPoolReleasedIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolReleased struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*BurnMintFastTransferTokenPoolReleasedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolReleasedIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "Released", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchReleased(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolReleased)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "Released", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseReleased(log types.Log) (*BurnMintFastTransferTokenPoolReleased, error) {
	event := new(BurnMintFastTransferTokenPoolReleased)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "Released", log); err != nil {
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

type BurnMintFastTransferTokenPoolTokensConsumedIterator struct {
	Event *BurnMintFastTransferTokenPoolTokensConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintFastTransferTokenPoolTokensConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintFastTransferTokenPoolTokensConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintFastTransferTokenPoolTokensConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintFastTransferTokenPoolTokensConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnMintFastTransferTokenPoolTokensConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintFastTransferTokenPoolTokensConsumed struct {
	Tokens *big.Int
	Raw    types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterTokensConsumed(opts *bind.FilterOpts) (*BurnMintFastTransferTokenPoolTokensConsumedIterator, error) {

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "TokensConsumed")
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolTokensConsumedIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "TokensConsumed", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchTokensConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolTokensConsumed) (event.Subscription, error) {

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "TokensConsumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintFastTransferTokenPoolTokensConsumed)
				if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "TokensConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) ParseTokensConsumed(log types.Log) (*BurnMintFastTransferTokenPoolTokensConsumed, error) {
	event := new(BurnMintFastTransferTokenPoolTokensConsumed)
	if err := _BurnMintFastTransferTokenPool.contract.UnpackLog(event, "TokensConsumed", log); err != nil {
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
	case _BurnMintFastTransferTokenPool.abi.Events["Burned"].ID:
		return _BurnMintFastTransferTokenPool.ParseBurned(log)
	case _BurnMintFastTransferTokenPool.abi.Events["ChainAdded"].ID:
		return _BurnMintFastTransferTokenPool.ParseChainAdded(log)
	case _BurnMintFastTransferTokenPool.abi.Events["ChainConfigured"].ID:
		return _BurnMintFastTransferTokenPool.ParseChainConfigured(log)
	case _BurnMintFastTransferTokenPool.abi.Events["ChainRemoved"].ID:
		return _BurnMintFastTransferTokenPool.ParseChainRemoved(log)
	case _BurnMintFastTransferTokenPool.abi.Events["ConfigChanged"].ID:
		return _BurnMintFastTransferTokenPool.ParseConfigChanged(log)
	case _BurnMintFastTransferTokenPool.abi.Events["DestinationPoolUpdated"].ID:
		return _BurnMintFastTransferTokenPool.ParseDestinationPoolUpdated(log)
	case _BurnMintFastTransferTokenPool.abi.Events["FastFill"].ID:
		return _BurnMintFastTransferTokenPool.ParseFastFill(log)
	case _BurnMintFastTransferTokenPool.abi.Events["FastFillRequest"].ID:
		return _BurnMintFastTransferTokenPool.ParseFastFillRequest(log)
	case _BurnMintFastTransferTokenPool.abi.Events["FastFillSettled"].ID:
		return _BurnMintFastTransferTokenPool.ParseFastFillSettled(log)
	case _BurnMintFastTransferTokenPool.abi.Events["FillerAllowListUpdated"].ID:
		return _BurnMintFastTransferTokenPool.ParseFillerAllowListUpdated(log)
	case _BurnMintFastTransferTokenPool.abi.Events["InvalidFill"].ID:
		return _BurnMintFastTransferTokenPool.ParseInvalidFill(log)
	case _BurnMintFastTransferTokenPool.abi.Events["LaneUpdated"].ID:
		return _BurnMintFastTransferTokenPool.ParseLaneUpdated(log)
	case _BurnMintFastTransferTokenPool.abi.Events["Locked"].ID:
		return _BurnMintFastTransferTokenPool.ParseLocked(log)
	case _BurnMintFastTransferTokenPool.abi.Events["Minted"].ID:
		return _BurnMintFastTransferTokenPool.ParseMinted(log)
	case _BurnMintFastTransferTokenPool.abi.Events["OwnershipTransferRequested"].ID:
		return _BurnMintFastTransferTokenPool.ParseOwnershipTransferRequested(log)
	case _BurnMintFastTransferTokenPool.abi.Events["OwnershipTransferred"].ID:
		return _BurnMintFastTransferTokenPool.ParseOwnershipTransferred(log)
	case _BurnMintFastTransferTokenPool.abi.Events["RateLimitAdminSet"].ID:
		return _BurnMintFastTransferTokenPool.ParseRateLimitAdminSet(log)
	case _BurnMintFastTransferTokenPool.abi.Events["Released"].ID:
		return _BurnMintFastTransferTokenPool.ParseReleased(log)
	case _BurnMintFastTransferTokenPool.abi.Events["RemotePoolAdded"].ID:
		return _BurnMintFastTransferTokenPool.ParseRemotePoolAdded(log)
	case _BurnMintFastTransferTokenPool.abi.Events["RemotePoolRemoved"].ID:
		return _BurnMintFastTransferTokenPool.ParseRemotePoolRemoved(log)
	case _BurnMintFastTransferTokenPool.abi.Events["RouterUpdated"].ID:
		return _BurnMintFastTransferTokenPool.ParseRouterUpdated(log)
	case _BurnMintFastTransferTokenPool.abi.Events["TokensConsumed"].ID:
		return _BurnMintFastTransferTokenPool.ParseTokensConsumed(log)

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

func (BurnMintFastTransferTokenPoolBurned) Topic() common.Hash {
	return common.HexToHash("0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7")
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

func (BurnMintFastTransferTokenPoolDestinationPoolUpdated) Topic() common.Hash {
	return common.HexToHash("0xb760e03fa04c0e86fcff6d0046cdcf22fb5d5b6a17d1e6f890b3456e81c40fd8")
}

func (BurnMintFastTransferTokenPoolFastFill) Topic() common.Hash {
	return common.HexToHash("0x05abb4c63204ac64d430895792b9a147e544fbb4591f7c1bd1b4436dbf63355f")
}

func (BurnMintFastTransferTokenPoolFastFillRequest) Topic() common.Hash {
	return common.HexToHash("0xbe96da9e73004694b4d32649f81f31784bd7c37daf64113f0ce1ebdd49a99a93")
}

func (BurnMintFastTransferTokenPoolFastFillSettled) Topic() common.Hash {
	return common.HexToHash("0x9a3acb7ef95f9a8a2384f156d99ec2f035807c667ae66326823feead6d08fdbd")
}

func (BurnMintFastTransferTokenPoolFillerAllowListUpdated) Topic() common.Hash {
	return common.HexToHash("0xccc0f2211c115acfa175a7923abdeb4b0a7c376d1b9e43c74973efe83d7d9e22")
}

func (BurnMintFastTransferTokenPoolInvalidFill) Topic() common.Hash {
	return common.HexToHash("0xad64960fe1d28c88faed204b509e08b3c9e07d9c1cb84991addc205e6bfca42f")
}

func (BurnMintFastTransferTokenPoolLaneUpdated) Topic() common.Hash {
	return common.HexToHash("0xbed278e7a8c5c763baafe3f3497295e65fd1dec8d51555c7b72c665e219d2deb")
}

func (BurnMintFastTransferTokenPoolLocked) Topic() common.Hash {
	return common.HexToHash("0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008")
}

func (BurnMintFastTransferTokenPoolMinted) Topic() common.Hash {
	return common.HexToHash("0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0")
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

func (BurnMintFastTransferTokenPoolReleased) Topic() common.Hash {
	return common.HexToHash("0x2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52")
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

func (BurnMintFastTransferTokenPoolTokensConsumed) Topic() common.Hash {
	return common.HexToHash("0x1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a")
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPool) Address() common.Address {
	return _BurnMintFastTransferTokenPool.address
}

type BurnMintFastTransferTokenPoolInterface interface {
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

	FilterAllowListAdd(opts *bind.FilterOpts) (*BurnMintFastTransferTokenPoolAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*BurnMintFastTransferTokenPoolAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*BurnMintFastTransferTokenPoolAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*BurnMintFastTransferTokenPoolAllowListRemove, error)

	FilterBurned(opts *bind.FilterOpts, sender []common.Address) (*BurnMintFastTransferTokenPoolBurnedIterator, error)

	WatchBurned(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolBurned, sender []common.Address) (event.Subscription, error)

	ParseBurned(log types.Log) (*BurnMintFastTransferTokenPoolBurned, error)

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

	FilterDestinationPoolUpdated(opts *bind.FilterOpts, dst []uint64) (*BurnMintFastTransferTokenPoolDestinationPoolUpdatedIterator, error)

	WatchDestinationPoolUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolDestinationPoolUpdated, dst []uint64) (event.Subscription, error)

	ParseDestinationPoolUpdated(log types.Log) (*BurnMintFastTransferTokenPoolDestinationPoolUpdated, error)

	FilterFastFill(opts *bind.FilterOpts, fillRequestId [][32]byte, fillId [][32]byte, filler []common.Address) (*BurnMintFastTransferTokenPoolFastFillIterator, error)

	WatchFastFill(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolFastFill, fillRequestId [][32]byte, fillId [][32]byte, filler []common.Address) (event.Subscription, error)

	ParseFastFill(log types.Log) (*BurnMintFastTransferTokenPoolFastFill, error)

	FilterFastFillRequest(opts *bind.FilterOpts, fillRequestId [][32]byte, dstChainSelector []uint64) (*BurnMintFastTransferTokenPoolFastFillRequestIterator, error)

	WatchFastFillRequest(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolFastFillRequest, fillRequestId [][32]byte, dstChainSelector []uint64) (event.Subscription, error)

	ParseFastFillRequest(log types.Log) (*BurnMintFastTransferTokenPoolFastFillRequest, error)

	FilterFastFillSettled(opts *bind.FilterOpts, fillRequestId [][32]byte) (*BurnMintFastTransferTokenPoolFastFillSettledIterator, error)

	WatchFastFillSettled(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolFastFillSettled, fillRequestId [][32]byte) (event.Subscription, error)

	ParseFastFillSettled(log types.Log) (*BurnMintFastTransferTokenPoolFastFillSettled, error)

	FilterFillerAllowListUpdated(opts *bind.FilterOpts, dst []uint64) (*BurnMintFastTransferTokenPoolFillerAllowListUpdatedIterator, error)

	WatchFillerAllowListUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolFillerAllowListUpdated, dst []uint64) (event.Subscription, error)

	ParseFillerAllowListUpdated(log types.Log) (*BurnMintFastTransferTokenPoolFillerAllowListUpdated, error)

	FilterInvalidFill(opts *bind.FilterOpts, fillRequestId [][32]byte, filler []common.Address) (*BurnMintFastTransferTokenPoolInvalidFillIterator, error)

	WatchInvalidFill(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolInvalidFill, fillRequestId [][32]byte, filler []common.Address) (event.Subscription, error)

	ParseInvalidFill(log types.Log) (*BurnMintFastTransferTokenPoolInvalidFill, error)

	FilterLaneUpdated(opts *bind.FilterOpts, destinationChainSelector []uint64) (*BurnMintFastTransferTokenPoolLaneUpdatedIterator, error)

	WatchLaneUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolLaneUpdated, destinationChainSelector []uint64) (event.Subscription, error)

	ParseLaneUpdated(log types.Log) (*BurnMintFastTransferTokenPoolLaneUpdated, error)

	FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*BurnMintFastTransferTokenPoolLockedIterator, error)

	WatchLocked(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolLocked, sender []common.Address) (event.Subscription, error)

	ParseLocked(log types.Log) (*BurnMintFastTransferTokenPoolLocked, error)

	FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*BurnMintFastTransferTokenPoolMintedIterator, error)

	WatchMinted(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseMinted(log types.Log) (*BurnMintFastTransferTokenPoolMinted, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintFastTransferTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*BurnMintFastTransferTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintFastTransferTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*BurnMintFastTransferTokenPoolOwnershipTransferred, error)

	FilterRateLimitAdminSet(opts *bind.FilterOpts) (*BurnMintFastTransferTokenPoolRateLimitAdminSetIterator, error)

	WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolRateLimitAdminSet) (event.Subscription, error)

	ParseRateLimitAdminSet(log types.Log) (*BurnMintFastTransferTokenPoolRateLimitAdminSet, error)

	FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*BurnMintFastTransferTokenPoolReleasedIterator, error)

	WatchReleased(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseReleased(log types.Log) (*BurnMintFastTransferTokenPoolReleased, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintFastTransferTokenPoolRemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*BurnMintFastTransferTokenPoolRemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintFastTransferTokenPoolRemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*BurnMintFastTransferTokenPoolRemotePoolRemoved, error)

	FilterRouterUpdated(opts *bind.FilterOpts) (*BurnMintFastTransferTokenPoolRouterUpdatedIterator, error)

	WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolRouterUpdated) (event.Subscription, error)

	ParseRouterUpdated(log types.Log) (*BurnMintFastTransferTokenPoolRouterUpdated, error)

	FilterTokensConsumed(opts *bind.FilterOpts) (*BurnMintFastTransferTokenPoolTokensConsumedIterator, error)

	WatchTokensConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolTokensConsumed) (event.Subscription, error)

	ParseTokensConsumed(log types.Log) (*BurnMintFastTransferTokenPoolTokensConsumed, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
