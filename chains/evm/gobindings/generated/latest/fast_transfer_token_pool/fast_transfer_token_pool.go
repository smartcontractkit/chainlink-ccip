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
	FillerAllowlistEnabled        bool
	FastTransferBpsFee            uint16
	SettlementOverheadGas         uint32
	RemoteChainSelector           uint64
	ChainFamilySelector           [4]byte
	MaxFillAmountPerRequest       *big.Int
	DestinationPool               []byte
	EvmToAnyMessageExtraArgsBytes []byte
	AddFillers                    []common.Address
	RemoveFillers                 []common.Address
}

type FastTransferTokenPoolAbstractDestChainConfigView struct {
	MaxFillAmountPerRequest       *big.Int
	FillerAllowlistEnabled        bool
	FastTransferBpsFee            uint16
	SettlementOverheadGas         uint32
	ChainFamilySelector           [4]byte
	DestinationPool               []byte
	EvmToAnyMessageExtraArgsBytes []byte
	AllowedFillers                []common.Address
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIBurnMintERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipReceive\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.Any2EVMMessage\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipSendToken\",\"inputs\":[{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"fillRequestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"computeFillId\",\"inputs\":[{\"name\":\"fillRequestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"fastFill\",\"inputs\":[{\"name\":\"fillRequestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"srcAmountToFill\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourceDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListedFillers\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCcipSendTokenFee\",\"inputs\":[{\"name\":\"settlementFeeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIFastTransferPool.Quote\",\"components\":[{\"name\":\"ccipSettlementFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fastTransferFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFastTransferTokenPoolAbstract.DestChainConfigView\",\"components\":[{\"name\":\"maxFillAmountPerRequest\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fillerAllowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"fastTransferBpsFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"settlementOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"destinationPool\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"evmToAnyMessageExtraArgsBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"allowedFillers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFillInfo\",\"inputs\":[{\"name\":\"fillId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFastTransferTokenPoolAbstract.FillInfo\",\"components\":[{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"enumFastTransferTokenPoolAbstract.FillState\"},{\"name\":\"filler\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isFillerAllowListed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"filler\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRouter\",\"inputs\":[{\"name\":\"newRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateDestChainConfig\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple\",\"internalType\":\"structFastTransferTokenPoolAbstract.DestChainConfigUpdateArgs\",\"components\":[{\"name\":\"fillerAllowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"fastTransferBpsFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"settlementOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"maxFillAmountPerRequest\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationPool\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"evmToAnyMessageExtraArgsBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"addFillers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removeFillers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateFillerAllowList\",\"inputs\":[{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addFillers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removeFillers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Burned\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigUpdated\",\"inputs\":[{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"bps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"maxFillAmountPerRequest\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"destinationPool\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"addFillers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"removeFillers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"},{\"name\":\"settlementOverheadGas\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fillerAllowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestinationPoolUpdated\",\"inputs\":[{\"name\":\"dst\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"destinationPool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FastTransferFilled\",\"inputs\":[{\"name\":\"fillRequestId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"fillId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"filler\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"destAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FastTransferRequested\",\"inputs\":[{\"name\":\"fillRequestId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"dstChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fastTransferFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FastTransferSettled\",\"inputs\":[{\"name\":\"fillRequestId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FillerAllowListUpdated\",\"inputs\":[{\"name\":\"dst\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"addFillers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"removeFillers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvalidFill\",\"inputs\":[{\"name\":\"fillRequestId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"filler\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"filledAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"expectedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Locked\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Minted\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Released\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RouterUpdated\",\"inputs\":[{\"name\":\"oldRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensConsumed\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AggregateValueMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"AggregateValueRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyFilled\",\"inputs\":[{\"name\":\"fillRequestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AlreadySettled\",\"inputs\":[{\"name\":\"fillRequestId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"FillerNotAllowlisted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"filler\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRouter\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"RateLimitMustBeDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TransferAmountExceedsMaxFillAmount\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x61012080604052346103d757617640803803809161001d8285610456565b8339810160a0828203126103d75781516001600160a01b038116908190036103d75761004b60208401610479565b60408401519091906001600160401b0381116103d75784019280601f850112156103d7578351936001600160401b038511610440578460051b9060208201956100976040519788610456565b86526020808701928201019283116103d757602001905b828210610428575050506100d060806100c960608701610487565b9501610487565b93331561041757600180546001600160a01b0319163317905581158015610406575b80156103f5575b6103e4578160209160049360805260c0526040519283809263313ce56760e01b82525afa600091816103a3575b50610378575b5060a052600480546001600160a01b0319166001600160a01b0384169081179091558151151560e0819052909190610256575b50156102405761010052604051617004908161063c823960805181818161206901528181612f5001528181613139015281816142e1015281816144b5015281816147c001528181614838015281816148f501528181615f5c01528181616017015261630a015260a0518181816118bf0152818161225e015281816127d301528181614747015281816159d30152615a56015260c051818181610ca90152818161180e015281816121050152818161271201528181612e5c015261437c015260e051818181610c390152818161408d015261684d01526101005181611bf20152f35b6335fdcccd60e21b600052600060045260246000fd5b90602090604051906102688383610456565b60008252600036813760e051156103675760005b82518110156102e3576001906001600160a01b0361029a828661049b565b5116856102a6826104dd565b6102b3575b50500161027c565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a138856102ab565b5092905060005b815181101561035e576001906001600160a01b03610308828561049b565b51168015610358578461031a826105db565b610328575b50505b016102ea565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a1388461031f565b50610322565b5050503861015f565b6335f4a7b360e01b60005260046000fd5b60ff1660ff821681810361038c575061012c565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d6020116103dc575b816103bf60209383610456565b810103126103d7576103d090610479565b9038610126565b600080fd5b3d91506103b2565b6342bcdf7f60e11b60005260046000fd5b506001600160a01b038116156100f9565b506001600160a01b038516156100f2565b639b15e16f60e01b60005260046000fd5b6020809161043584610487565b8152019101906100ae565b634e487b7160e01b600052604160045260246000fd5b601f909101601f19168101906001600160401b0382119082101761044057604052565b519060ff821682036103d757565b51906001600160a01b03821682036103d757565b80518210156104af5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b80548210156104af5760005260206000200190600090565b60008181526003602052604090205480156105d45760001981018181116105be576002546000198101919082116105be5781810361056d575b505050600254801561055757600019016105318160026104c5565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b6105a661057e61058f9360026104c5565b90549060031b1c92839260026104c5565b819391549060031b91821b91600019901b19161790565b90556000526003602052604060002055388080610516565b634e487b7160e01b600052601160045260246000fd5b5050600090565b8060005260036020526040600020541560001461063557600254680100000000000000008110156104405761061c61058f82600185940160025560026104c5565b9055600254906000526003602052604060002055600190565b5060009056fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a714614c0e575080630eb9eb1614614b86578063181f5a7714614ae1578063211ffa041461485c57806321df0da7146147ed578063240028e81461476b57806324f65ee71461470f578063390775371461423b5780634c5ef0ed146141d657806354c8a4f31461405957806362ddd3c414613fd5578063669c6f1614613e675780636d3d1a5814613e155780636def4ce714613c3757806372b477e1146134fb57806379ba5097146134165780637d54534e1461336957806385572ffb14612bcc5780638926f54f14612b685780638d4470e4146126425780638da5cb5b146125f0578063962d40201461244c5780639a4575b914612020578063a22f8f7e1461174b578063a42a7b8b146115c6578063a7cd63b714611533578063abe1c1e81461144f578063acfecf911461132b578063af58d59f146112c4578063b0f479a114611272578063b79465801461121b578063c0d786551461111a578063c4bffe2b14610fd1578063c75eea9c14610f0f578063ccc316fb14610e9c578063cf7401f314610ccd578063dc0bd97114610c5e578063e0351e1314610c03578063e0c118b414610b99578063e8a1da17146102d35763f2fde38b146101e457600080fd5b346102d05760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d05773ffffffffffffffffffffffffffffffffffffffff610230614e22565b610238615b60565b163381146102a857807fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b80fd5b50346102d0576102e2366150a5565b939190926102ee615b60565b82915b808310610a04575050508063ffffffff4216917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee1843603015b85821015610a00578160051b850135818112156109fc57850190610120823603126109fc576040519561035c87614eea565b61036583614e0d565b8752602083013567ffffffffffffffff81116109f85783019536601f880112156109f857863596610395886151a2565b976103a3604051998a614f3e565b8089526020808a019160051b830101903682116109f45760208301905b8282106109c1575050505060208801968752604084013567ffffffffffffffff81116109bd576103f39036908601615056565b9860408901998a5261041d61040b36606088016153e4565b9560608b0196875260c03691016153e4565b9660808a0197885261042f8651616178565b6104398851616178565b8a5151156109955761045567ffffffffffffffff8b51166163e7565b1561095e5767ffffffffffffffff8a5116815260076020526040812061059587516fffffffffffffffffffffffffffffffff604082015116906105506fffffffffffffffffffffffffffffffff602083015116915115158360806040516104bb81614eea565b858152602081018c905260408101849052606081018690520152855474ff000000000000000000000000000000000000000091151560a01b919091167fffffffffffffffffffffff0000000000000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff84161773ffffffff0000000000000000000000000000000060808b901b1617178555565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176001830155565b6106bb89516fffffffffffffffffffffffffffffffff604082015116906106766fffffffffffffffffffffffffffffffff602083015116915115158360806040516105df81614eea565b858152602081018c9052604081018490526060810186905201526002860180547fffffffffffffffffffffff000000000000000000000000000000000000000000166fffffffffffffffffffffffffffffffff85161773ffffffff0000000000000000000000000000000060808c901b161791151560a01b74ff000000000000000000000000000000000000000016919091179055565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176003830155565b60048c5191019080519067ffffffffffffffff8211610931576106de8354615558565b601f81116108f6575b50602090601f831160011461085757610735929185918361084c575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b805b89518051821015610770579061076a600192610763838f67ffffffffffffffff90511692615515565b5190615bab565b0161073a565b5050975097987f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c29295939661083e67ffffffffffffffff600197949c511692519351915161080a6107d560405196879687526101006020880152610100870190614fdc565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a101909394929161032a565b015190503880610703565b83855281852091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416865b8181106108de57509084600195949392106108a7575b505050811b019055610738565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c1916905538808061089a565b92936020600181928786015181550195019301610884565b6109219084865260208620601f850160051c81019160208610610927575b601f0160051c01906158ee565b386106e7565b9091508190610914565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60249067ffffffffffffffff8b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b807f8579befe0000000000000000000000000000000000000000000000000000000060049252fd5b8680fd5b813567ffffffffffffffff81116109f0576020916109e58392833691890101615056565b8152019101906103c0565b8a80fd5b8880fd5b8580fd5b8380fd5b8280f35b9092919367ffffffffffffffff610a24610a1f878588615686565b615696565b1695610a2f87616dd2565b15610b6d578684526007602052610a4b600560408620016165cd565b94845b8651811015610a84576001908987526007602052610a7d60056040892001610a76838b615515565b5190616ea2565b5001610a4e565b5093945094909580855260076020526005604086208681558660018201558660028201558660038201558660048201610abd8154615558565b80610b2c575b5050500180549086815581610b0e575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a10191909493946102f1565b865260208620908101905b81811015610ad357868155600101610b19565b601f8111600114610b425750555b863880610ac3565b81835260208320610b5d91601f01861c8101906001016158ee565b8082528160208120915555610b3a565b602484887f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50346102d05760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d0576044359073ffffffffffffffffffffffffffffffffffffffff821682036102d0576020610bfb83602435600435616136565b604051908152f35b50346102d057807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d05760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b50346102d057807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d057602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102d05760e07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d057610d05614df6565b9060607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc3601126102d057604051610d3c81614f22565b6024358015158103610e985781526044356fffffffffffffffffffffffffffffffff81168103610e985760208201526064356fffffffffffffffffffffffffffffffff81168103610e9857604082015260607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c360112610e945760405190610dc382614f22565b60843580151581036109fc57825260a4356fffffffffffffffffffffffffffffffff811681036109fc57602083015260c4356fffffffffffffffffffffffffffffffff811681036109fc57604083015273ffffffffffffffffffffffffffffffffffffffff6009541633141580610e72575b610e4657610e439293615e15565b80f35b6024837f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b5073ffffffffffffffffffffffffffffffffffffffff60015416331415610e35565b5080fd5b8280fd5b50346102d05760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d057610ef760036040610f0b9367ffffffffffffffff610ee8614df6565b168152600a60205220016165cd565b604051918291602083526020830190615229565b0390f35b50346102d05760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d057610f78610f736040610f0b9367ffffffffffffffff610f5c614df6565b610f6461583b565b50168152600760205220615866565b6160b1565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b50346102d057807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d057604051906005548083528260208101600584526020842092845b81811061110157505061102f92500383614f3e565b815161105361103d826151a2565b9161104b6040519384614f3e565b8083526151a2565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602083019301368437805b84518110156110b2578067ffffffffffffffff61109f60019388615515565b51166110ab8286615515565b5201611080565b50925090604051928392602084019060208552518091526040840192915b8181106110de575050500390f35b825167ffffffffffffffff168452859450602093840193909201916001016110d0565b845483526001948501948794506020909301920161101a565b50346102d05760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d057611152614e22565b61115a615b60565b73ffffffffffffffffffffffffffffffffffffffff81169081156111f357600480547fffffffffffffffffffffffff000000000000000000000000000000000000000081169390931790556040805173ffffffffffffffffffffffffffffffffffffffff93841681529190921660208201527f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f16849190a180f35b6004837f8579befe000000000000000000000000000000000000000000000000000000008152fd5b50346102d05760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d057610f0b61125e611259614df6565b6158cc565b604051918291602083526020830190614fdc565b50346102d057807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d057602073ffffffffffffffffffffffffffffffffffffffff60045416604051908152f35b50346102d05760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d057610f78610f7360026040610f0b9467ffffffffffffffff611313614df6565b61131b61583b565b5016815260076020522001615866565b50346102d05767ffffffffffffffff61134336615143565b92909161134e615b60565b1691611367836000526006602052604060002054151590565b156114235782845260076020526113966005604086200161138936848661501f565b6020815191012090616ea2565b156113db57907f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76916113d56040519283926020845260208401916157fc565b0390a280f35b8261141f836040519384937f74f23c7c00000000000000000000000000000000000000000000000000000000855260048501526040602485015260448401916157fc565b0390fd5b602484847f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50346102d05760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d05761148761566d565b506004358152600b60205260408120604051906114a382614e66565b54916114b260ff84168361542b565b73ffffffffffffffffffffffffffffffffffffffff602083019360081c1683526040519151906003821015611506575060409273ffffffffffffffffffffffffffffffffffffffff91835251166020820152f35b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526021600452fd5b50346102d057807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d05760405160028054808352908352909160208301917f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace915b8181106115b057610f0b85610ef781870382614f3e565b8254845260209093019260019283019201611599565b50346102d05760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d05767ffffffffffffffff611607614df6565b168152600760205261161e600560408320016165cd565b80517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061166361164d836151a2565b9261165b6040519485614f3e565b8084526151a2565b01835b81811061173a575050825b82518110156116b7578061168760019285615515565b518552600860205261169b604086206155ab565b6116a58285615515565b526116b08184615515565b5001611671565b81846040519182916020830160208452825180915260408401602060408360051b870101940192905b8282106116ef57505050500390f35b9193602061172a827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851614fdc565b96019201920185949391926116e0565b806060602080938601015201611666565b50611755366152e2565b5050919060009594955060405161176b81614e66565b60008152600060208201526060608060405161178681614eea565b8281528260208201528260408201528883820152015267ffffffffffffffff8516946040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff000000000000000000000000000000008260801b16600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115612015578891611fe6575b50611fbe5761184c3361684b565b611863866000526006602052604060002054151590565b15611f9257858752600a602052604087209081548511611f62576001820154928360081c61ffff166118959087615905565b61271090049960208201948b86526040516118af81614f06565b888152602081019c8d52604081017f000000000000000000000000000000000000000000000000000000000000000060ff168152366118ef908c8b61501f565b90606083019182526040519e8f9360208501602090525160408501525160608401525160ff1660808301525160a082016080905260c0820161193091614fdc565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018d52611960908d614f3e565b8060a81b7fffffffff00000000000000000000000000000000000000000000000000000000167f2812d52c0000000000000000000000000000000000000000000000000000000081148015611f39575b8015611f10575b15611e9e575060181c63ffffffff1680611e25575073ffffffffffffffffffffffffffffffffffffffff6119ed600586016155ab565b915b60405160209d8e8e611a018285614f3e565b8352611a1b60026040519a611a158c614eea565b016155ab565b895288015260408701521690816060860152608085015273ffffffffffffffffffffffffffffffffffffffff600454168b60405180927f20487ded0000000000000000000000000000000000000000000000000000000082528180611a848a8a600484016156bb565b03915afa908115611e1a57908c949392918c91611deb575b508252611aa98884615fe3565b611ab288615f45565b80611bdf575b505091611b0e9273ffffffffffffffffffffffffffffffffffffffff60045416906040518095819482937f96f4e9f9000000000000000000000000000000000000000000000000000000008452600484016156bb565b039134905af1968715611bd3578097611b77575b505091611b6c8694927f1199f551568004635134aaf2ae681cd3c00d4baf32ee26dd8b2b8d583b051d9a94519360405194859485528a8501526060604085015260608401916157fc565b0390a3604051908152f35b909194939296508782813d8311611bcc575b611b938183614f3e565b810103126102d0575051949192909190611b6c7f1199f551568004635134aaf2ae681cd3c00d4baf32ee26dd8b2b8d583b051d9a611b22565b503d611b89565b604051903d90823e3d90fd5b90919250611bf0825130338461603d565b7f000000000000000000000000000000000000000000000000000000000000000091519182158015611d48575b15611cc457611b0e94928c9492611cb7611cbc93611c8b6040519485927f095ea7b3000000000000000000000000000000000000000000000000000000008b850152602484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101845283614f3e565b616bc2565b909238611ab8565b60848c604051907f08c379a00000000000000000000000000000000000000000000000000000000082526004820152603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527f20746f206e6f6e2d7a65726f20616c6c6f77616e6365000000000000000000006064820152fd5b506040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff821660248201528c81604481865afa908115611de0578c91611dae575b5015611c1d565b90508c81813d8311611dd9575b611dc58183614f3e565b81010312611dd4575138611da7565b600080fd5b503d611dbb565b6040513d8e823e3d90fd5b85819692503d8311611e13575b611e028183614f3e565b81010312611dd4578b935138611a9c565b503d611df8565b6040513d8d823e3d90fd5b60405173ffffffffffffffffffffffffffffffffffffffff91611e4782614e66565b81526020810160018152604051917f181dcf10000000000000000000000000000000000000000000000000000000006020840152516024830152511515604482015260448152611e98606482614f3e565b916119ef565b7f1e10bdc40000000000000000000000000000000000000000000000000000000014159050611ee85773ffffffffffffffffffffffffffffffffffffffff611e98600586016155ab565b60048a7f382c0982000000000000000000000000000000000000000000000000000000008152fd5b507fc4e059530000000000000000000000000000000000000000000000000000000081146119b7565b507fac77ffec0000000000000000000000000000000000000000000000000000000081146119b0565b60448886897f58dd87c5000000000000000000000000000000000000000000000000000000008352600452602452fd5b602487877fa9902c7e000000000000000000000000000000000000000000000000000000008252600452fd5b6004877f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b612008915060203d60201161200e575b6120008183614f3e565b810190615947565b3861183e565b503d611ff6565b6040513d8a823e3d90fd5b50346102d05761202f36615273565b6060602060405161203f81614e66565b828152015260808101612051816154b7565b73ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116036124025750602081019177ffffffffffffffff000000000000000000000000000000006120b884615696565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156123845782916123e3575b506123bb5761214e612149604084016154b7565b61684b565b67ffffffffffffffff61216084615696565b16612178816000526006602052604060002054151590565b1561238f57602073ffffffffffffffffffffffffffffffffffffffff60045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015612384578290612321575b73ffffffffffffffffffffffffffffffffffffffff91501633036122f5576122c4612254611259858561221a606061221084615696565b9201358092615fe3565b61222381615f45565b6040519081527f696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df760203392a2615696565b610f0b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016602082015260208152612292604082614f3e565b6040519261229f84614e66565b8352602083019081526040519384936020855251604060208601526060850190614fdc565b90517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0848303016040850152614fdc565b807f728fe07b000000000000000000000000000000000000000000000000000000006024925233600452fd5b506020813d60201161237c575b8161233b60209383614f3e565b81010312610e94575173ffffffffffffffffffffffffffffffffffffffff81168103610e945773ffffffffffffffffffffffffffffffffffffffff906121d9565b3d915061232e565b6040513d84823e3d90fd5b602492507fa9902c7e000000000000000000000000000000000000000000000000000000008252600452fd5b807f53ad11d80000000000000000000000000000000000000000000000000000000060049252fd5b6123fc915060203d60201161200e576120008183614f3e565b38612135565b8273ffffffffffffffffffffffffffffffffffffffff6124236024936154b7565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835216600452fd5b50346102d05760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d05760043567ffffffffffffffff8111610e945761249c903690600401615074565b60243567ffffffffffffffff81116109fc576124bc903690600401615389565b60449291923567ffffffffffffffff81116109f8576124df903690600401615389565b91909273ffffffffffffffffffffffffffffffffffffffff60095416331415806125ce575b6125a257818114801590612598575b61257057865b818110612524578780f35b8061256a612538610a1f600194868c615686565b61254383878b6156ab565b61256461255c612554868b8d6156ab565b9236906153e4565b9136906153e4565b91615e15565b01612519565b6004877f568efce2000000000000000000000000000000000000000000000000000000008152fd5b5082811415612513565b6024877f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b5073ffffffffffffffffffffffffffffffffffffffff60015416331415612504565b50346102d057807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d057602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b50346102d057612651366152e2565b505094939261266192919261566d565b506040519561266f87614e66565b60008752600060208801526060608060405161268a81614eea565b8281528260208201528260408201528883820152015267ffffffffffffffff8216926040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff000000000000000000000000000000008460801b16600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115612015578891612b49575b50611fbe576127503361684b565b612767846000526006602052604060002054151590565b15612b1d57838752600a602052604087209384548211612aef57509061286184939261283560018b9701549360ff6127fc60206127106127ae61ffff8a60081c1688615905565b049a019a8a8c52604051956127c287614f06565b8652602086019a8b526040860193837f0000000000000000000000000000000000000000000000000000000000000000168552369161501f565b9160608501928352604051998a956020808801525160408701525160608601525116608084015251608060a084015260c0830190614fdc565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101865285614f3e565b7fffffffff000000000000000000000000000000000000000000000000000000008160a81b167f2812d52c0000000000000000000000000000000000000000000000000000000081148015612ac6575b8015612a9d575b15612a3a57506129829392919060181c63ffffffff16806129d757506128e0600583016155ab565b73ffffffffffffffffffffffffffffffffffffffff6020978895604051906129088883614f3e565b8b825261291d600260405198611a158a614eea565b8752878701526040860152166060840152608083015273ffffffffffffffffffffffffffffffffffffffff60045416906040518095819482937f20487ded000000000000000000000000000000000000000000000000000000008452600484016156bb565b03915afa938415611bd357936129a5575b50826040945283519283525190820152f35b9392508184813d83116129d0575b6129bd8183614f3e565b81010312611dd457604093519293612993565b503d6129b3565b604051906129e482614e66565b81526020810160018152604051917f181dcf10000000000000000000000000000000000000000000000000000000006020840152516024830152511515604482015260448152612a35606482614f3e565b6128e0565b7f1e10bdc40000000000000000000000000000000000000000000000000000000014159050612a7557906129829291612a35600583016155ab565b6004867f382c0982000000000000000000000000000000000000000000000000000000008152fd5b507fc4e059530000000000000000000000000000000000000000000000000000000081146128b8565b507fac77ffec0000000000000000000000000000000000000000000000000000000081146128b1565b7f58dd87c5000000000000000000000000000000000000000000000000000000008852600452602452604486fd5b602487857fa9902c7e000000000000000000000000000000000000000000000000000000008252600452fd5b612b62915060203d60201161200e576120008183614f3e565b38612742565b50346102d05760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d0576020612bc267ffffffffffffffff612bae614df6565b166000526006602052604060002054151590565b6040519015158152f35b50346102d057612bdb36615273565b9073ffffffffffffffffffffffffffffffffffffffff60045416330361333d5760a0823603126102d057604051612c1181614eea565b82358152612c2160208401614e0d565b60208201908152604084013567ffffffffffffffff81116109fc57612c499036908601615056565b9360408301948552606081013567ffffffffffffffff811161330157612c729036908301615056565b906060840191825260808101359067ffffffffffffffff82116109f8570136601f82011215613301578035612ca6816151a2565b91612cb46040519384614f3e565b81835260208084019260061b820101903682116132fd57602001915b81831061330557505050608084015251908151820160208101926020818303126109f85760208101519067ffffffffffffffff82116109bd570190608090829003126133015760405190612d2382614f06565b602081015182526040810151906020830191825260608101519060ff821682036132fd576040840191825260808101519067ffffffffffffffff82116109f4579060209101019385601f860112156132fd578451612d8081614f7f565b95612d8e6040519788614f3e565b8187526020870197602083830101116132f95790876020612daf9301614fb9565b846060850152519273ffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff85169660ff89519b51935194511694519651905190602081106132c9575b50169377ffffffffffffffff00000000000000000000000000000000604051917f2cbc26bb00000000000000000000000000000000000000000000000000000000835260801b16600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156132be57899161329f575b5061327757612e9b81876154d8565b156132395750612ec2612ebc83612eb6612ec8958795615a53565b96615a53565b856160a4565b88616136565b92838652600b602052604086209660405197612ee389614e66565b54612ef160ff82168a61542b565b73ffffffffffffffffffffffffffffffffffffffff60208a019160081c16815288519860038a101561320c5788998997989950156000146130c05750505082612f39916162d3565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691823b156109fc576040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff92909216600483015260248201529082908290604490829084905af18015612384576130ab575b50505b60405190612fea82614e66565b6002825260208201908482528452600b60205260408420915190600382101561307e577fffffffffffffffffffffff00000000000000000000000000000000000000000060ff74ffffffffffffffffffffffffffffffffffffffff008554935160081b169316911617179055517fd5d4e34f92d581a906fa8f24b4d8c5639f8931ede3f4ad88f6aaf2d634d653c98280a280f35b6024857f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b816130b591614f3e565b610e98578238612fda565b92509250959493505160038110156131df5760020361310557602486867fb196a44a000000000000000000000000000000000000000000000000000000008252600452fd5b859293945073ffffffffffffffffffffffffffffffffffffffff90511673ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691823b156109fc576040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff92909216600483015260248201529082908290604490829084905af18015612384576131ca575b5050612fdd565b816131d491614f3e565b610e985782386131c3565b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b6024897f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b61141f906040519182917f24eb47e5000000000000000000000000000000000000000000000000000000008352602060048401526024830190614fdc565b6004887f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b6132b8915060203d60201161200e576120008183614f3e565b38612e8c565b6040513d8b823e3d90fd5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1638612df3565b8980fd5b8780fd5b8480fd5b6040833603126132fd576020604091825161331f81614e66565b61332886614e45565b81528286013583820152815201920191612cd0565b807fd7f73334000000000000000000000000000000000000000000000000000000006024925233600452fd5b50346102d05760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d0577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174602073ffffffffffffffffffffffffffffffffffffffff6133d9614e22565b6133e1615b60565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000006009541617600955604051908152a180f35b50346102d057807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d057805473ffffffffffffffffffffffffffffffffffffffff811633036134d3577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516825573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b50346102d05760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d0576004359067ffffffffffffffff82116102d0576101407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc83360301126102d057613573615b60565b60405191610140830183811067ffffffffffffffff821117613c0a5760405261359e816004016153ba565b835260248101359061ffff82168203610e98576020840191825260448101359263ffffffff841684036102d057604085019384526135de60648301614e0d565b90606086019182526084830135937fffffffff0000000000000000000000000000000000000000000000000000000085168503610e94576080870194855260a0870160a4850135815260c485013567ffffffffffffffff81116109fc5761364b9060043691880101615056565b9060c0890191825260e486013567ffffffffffffffff8111613301576136779060043691890101615056565b9760e08a0198895261010487013567ffffffffffffffff81116109f8576136a490600436918a010161520e565b966101008b0197885261012481013567ffffffffffffffff81116109bd576136d19136910160040161520e565b966101208b0197885261271061ffff86511611612a755767ffffffffffffffff8751168652600a6020526040862099845160028c019080519067ffffffffffffffff8211613b0e576137238354615558565b601f8111613bda575b50602090601f8311600114613b3b5761377992918b918361084c5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b898b8d6affffffff000000000000008951927fffffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffff82600183019660ff88549651151516938c5190555160a81c16937fffffffffffffffffffffffffffffffffffffffffff00000000ffffffff00000062ffff0066ffffffff0000008c5160181b169760081b169116171617911617179055805160058c019080519067ffffffffffffffff8211613b0e5761382f8354615558565b601f8111613ade575b50602090601f8311600114613a3f5761388592918b918361084c5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b7f1e10bdc4000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008b511614613a32575b506003869b9a019a5b8b8951805183101561391d57829173ffffffffffffffffffffffffffffffffffffffff61390e60019561391694615515565b5116906164f1565b50016138dc565b50998b9150875b83518051821015613964579061395d8d73ffffffffffffffffffffffffffffffffffffffff61395584600196615515565b511690616420565b5001613924565b5050613a1c90613a0e8b613a0063ffffffff7fffffffff0000000000000000000000000000000000000000000000000000000061ffff7f391233e9b9f2fa088f64269fb39c6cf49aeeaad3183e75aafb2e2fd906f770e79d9e9f67ffffffffffffffff9051169e51169a519b519851935195511698511695511515966040519a8b9a8b5260208b015261010060408b01526101008a0190614fdc565b9088820360608a0152615229565b908682036080880152615229565b9260a085015260c084015260e08301520390a280f35b515115612a7557386138d3565b838b52818b2091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe084168c5b818110613ac65750908460019594939210613a8f575b505050811b019055613888565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388080613a82565b92936020600181928786015181550195019301613a6c565b613b0890848c5260208c20601f850160051c8101916020861061092757601f0160051c01906158ee565b38613838565b60248a7f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b838b52818b2091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe084168c5b818110613bc25750908460019594939210613b8b575b505050811b01905561377c565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388080613b7e565b92936020600181928786015181550195019301613b68565b613c0490848c5260208c20601f850160051c8101916020861061092757601f0160051c01906158ee565b3861372c565b6024837f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b50346102d05760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d0576040809167ffffffffffffffff613c7c614df6565b606060e08551613c8b81614ecd565b858152856020820152858782015285838201528560808201528260a08201528260c08201520152168152600a60205220610f0b815491613de3600182015493613db2613cd9600385016165cd565b93875192613ce684614ecd565b83527fffffffff00000000000000000000000000000000000000000000000000000000602084019760ff81161515895263ffffffff8a860161ffff8360081c16815261ffff6060880191838560181c1683528560808a019560a81b168552613d636005613d5560028a016155ab565b9860a08c01998a52016155ab565b9860c08101998a5260e081019b8c528e519e8f9e8f9260208452516020840152511515910152511660608c0152511660808a0152511660a08801525161010060c0880152610120870190614fdc565b90517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08683030160e0870152614fdc565b90517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe084830301610100850152615229565b50346102d057807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d057602073ffffffffffffffffffffffffffffffffffffffff60095416604051908152f35b50346102d05760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d057613e9f614df6565b9060243567ffffffffffffffff8111610e9457613ec090369060040161520e565b60443567ffffffffffffffff8111610e98579067ffffffffffffffff613eeb8593369060040161520e565b92613ef4615b60565b1691828452600a60205260408420936003819501945b8351811015613f475780613f4073ffffffffffffffffffffffffffffffffffffffff613f3860019488615515565b511688616420565b5001613f0a565b508493815b8351811015613f895780613f8273ffffffffffffffffffffffffffffffffffffffff613f7a60019488615515565b5116886164f1565b5001613f4c565b507fccc0f2211c115acfa175a7923abdeb4b0a7c376d1b9e43c74973efe83d7d9e22613fc7856113d586604051938493604085526040850190615229565b908382036020850152615229565b50346102d057613fe436615143565b613ff093929193615b60565b67ffffffffffffffff8216614012816000526006602052604060002054151590565b1561402e5750610e43929361402891369161501f565b90615bab565b7f1e670e4b000000000000000000000000000000000000000000000000000000008452600452602483fd5b50346102d0576140839061408b61406f366150a5565b959161407c939193615b60565b36916151ba565b9336916151ba565b7f0000000000000000000000000000000000000000000000000000000000000000156141ae57815b8351811015614126578073ffffffffffffffffffffffffffffffffffffffff6140de60019387615515565b51166140e981616d02565b6140f5575b50016140b3565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a1386140ee565b5090805b82518110156141aa578073ffffffffffffffffffffffffffffffffffffffff61415560019386615515565b511680156141a457614166816163a8565b614173575b505b0161412a565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a18461416b565b5061416d565b5080f35b6004827f35f4a7b3000000000000000000000000000000000000000000000000000000008152fd5b50346102d05760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d05761420e614df6565b906024359067ffffffffffffffff82116102d0576020612bc2846142353660048701615056565b906154d8565b50346102d05760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d05760043567ffffffffffffffff8111610e9457806004016101007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8336030112610e9857826040516142bb81614eb1565b52608482016142c9816154b7565b73ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116036146ee57506024820177ffffffffffffffff0000000000000000000000000000000061432f82615696565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156146715785916146cf575b506146a75767ffffffffffffffff6143c382615696565b166143db816000526006602052604060002054151590565b1561467c57602073ffffffffffffffffffffffffffffffffffffffff60045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115614671578591614652575b50156146265761445281615696565b61446e60a48501916142356144678487615466565b369161501f565b156145df5750906144a96144a461446761448a6144af95615696565b9361449a606488013580966162d3565b60c4870190615466565b61595f565b90615a53565b906044017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16836144f4836154b7565b823b15610e94576040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91909116600482015260248101859052918290604490829084905af180156145d457916020946145829273ffffffffffffffffffffffffffffffffffffffff946145c4575b50506154b7565b166040518281527f9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0843392a3806040516145bb81614eb1565b52604051908152f35b816145ce91614f3e565b3861457b565b6040513d86823e3d90fd5b6145e99083615466565b61141f6040519283927f24eb47e50000000000000000000000000000000000000000000000000000000084526020600485015260248401916157fc565b6024847f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b61466b915060203d60201161200e576120008183614f3e565b38614443565b6040513d87823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008552600452602484fd5b6004847f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b6146e8915060203d60201161200e576120008183614f3e565b386143ac565b8373ffffffffffffffffffffffffffffffffffffffff6124236024936154b7565b50346102d057807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d057602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102d05760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d0576020906147a6614e22565b905073ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b50346102d057807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d057602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102d05760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d05760243567ffffffffffffffff81166004358183036109fc5760643560ff81168103613301576084359073ffffffffffffffffffffffffffffffffffffffff8216918281036109bd576148ed6148e661492093604435615a53565b80976162d3565b6149198682337f000000000000000000000000000000000000000000000000000000000000000061603d565b8584616136565b92838652600b6020526040862073ffffffffffffffffffffffffffffffffffffffff6040519161494f83614e66565b5461495d60ff82168461542b565b60081c1660208201525160038110156131df57614ab557808652600a6020526040862060ff600182015416614a68575b505060405161499b81614e66565b6001815260208101338152848752600b602052604087209151906003821015614a3b577fffffffffffffffffffffff00000000000000000000000000000000000000000060ff74ffffffffffffffffffffffffffffffffffffffff008554935160081b16931691161717905560405193845260208401527fd6f70fb263bfe7d01ec6802b3c07b6bd32579760fe9fcb4e248a036debb8cdf160403394a480f35b6024887f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b3360009081526004909101602052604090205415614a86578061498d565b7f6c46a9b500000000000000000000000000000000000000000000000000000000865260045233602452604485fd5b602486847fcee81443000000000000000000000000000000000000000000000000000000008252600452fd5b50346102d057807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d05750610f0b604051614b22606082614f3e565b602381527f4275726e4d696e74466173745472616e73666572546f6b656e506f6f6c20312e60208201527f362e3100000000000000000000000000000000000000000000000000000000006040820152604051918291602083526020830190614fdc565b50346102d05760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d057614bbe614df6565b6024359073ffffffffffffffffffffffffffffffffffffffff8216809203610e985767ffffffffffffffff168252600a602090815260409283902060009283526004018152919020541515612bc2565b905034610e945760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610e94576004357fffffffff000000000000000000000000000000000000000000000000000000008116809103610e9857602092507faff2afbf000000000000000000000000000000000000000000000000000000008114908115614dcc575b8115614da2575b8115614cb0575b5015158152f35b7feeb51d2a00000000000000000000000000000000000000000000000000000000811491508115614d78575b8115614d4e575b8115614cf1575b5038614ca9565b7f85572ffb00000000000000000000000000000000000000000000000000000000811491508115614d24575b5038614cea565b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438614d1d565b7f01ffc9a70000000000000000000000000000000000000000000000000000000081149150614ce3565b7f85572ffb0000000000000000000000000000000000000000000000000000000081149150614cdc565b7f01ffc9a70000000000000000000000000000000000000000000000000000000081149150614ca2565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150614c9b565b6004359067ffffffffffffffff82168203611dd457565b359067ffffffffffffffff82168203611dd457565b6004359073ffffffffffffffffffffffffffffffffffffffff82168203611dd457565b359073ffffffffffffffffffffffffffffffffffffffff82168203611dd457565b6040810190811067ffffffffffffffff821117614e8257604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff821117614e8257604052565b610100810190811067ffffffffffffffff821117614e8257604052565b60a0810190811067ffffffffffffffff821117614e8257604052565b6080810190811067ffffffffffffffff821117614e8257604052565b6060810190811067ffffffffffffffff821117614e8257604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117614e8257604052565b67ffffffffffffffff8111614e8257601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60005b838110614fcc5750506000910152565b8181015183820152602001614fbc565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f60209361501881518092818752878088019101614fb9565b0116010190565b92919261502b82614f7f565b916150396040519384614f3e565b829481845281830111611dd4578281602093846000960137010152565b9080601f83011215611dd4578160206150719335910161501f565b90565b9181601f84011215611dd45782359167ffffffffffffffff8311611dd4576020808501948460051b010111611dd457565b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc820112611dd45760043567ffffffffffffffff8111611dd457816150ee91600401615074565b929092916024359067ffffffffffffffff8211611dd45761511191600401615074565b9091565b9181601f84011215611dd45782359167ffffffffffffffff8311611dd45760208381860195010111611dd457565b9060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc830112611dd45760043567ffffffffffffffff81168103611dd457916024359067ffffffffffffffff8211611dd45761511191600401615115565b67ffffffffffffffff8111614e825760051b60200190565b9291906151c6816151a2565b936151d46040519586614f3e565b602085838152019160051b8101928311611dd457905b8282106151f657505050565b6020809161520384614e45565b8152019101906151ea565b9080601f83011215611dd457816020615071933591016151ba565b906020808351928381520192019060005b8181106152475750505090565b825173ffffffffffffffffffffffffffffffffffffffff1684526020938401939092019160010161523a565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc820112611dd4576004359067ffffffffffffffff8211611dd4577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8260a092030112611dd45760040190565b9060a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc830112611dd45760043573ffffffffffffffffffffffffffffffffffffffff81168103611dd4579160243567ffffffffffffffff81168103611dd457916044359160643567ffffffffffffffff8111611dd4578161536691600401615115565b929092916084359067ffffffffffffffff8211611dd45761511191600401615115565b9181601f84011215611dd45782359167ffffffffffffffff8311611dd45760208085019460608502010111611dd457565b35908115158203611dd457565b35906fffffffffffffffffffffffffffffffff82168203611dd457565b9190826060910312611dd4576040516153fc81614f22565b604061542681839561540d816153ba565b855261541b602082016153c7565b6020860152016153c7565b910152565b60038210156154375752565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215611dd4570180359067ffffffffffffffff8211611dd457602001918136038313611dd457565b3573ffffffffffffffffffffffffffffffffffffffff81168103611dd45790565b9067ffffffffffffffff61507192166000526007602052600560406000200190602081519101209060019160005201602052604060002054151590565b80518210156155295760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600182811c921680156155a1575b602083101461557257565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691615567565b90604051918260008254926155bf84615558565b808452936001811690811561562d57506001146155e6575b506155e492500383614f3e565b565b90506000929192526020600020906000915b8183106156115750509060206155e492820101386155d7565b60209193508060019154838589010152019101909184926155f8565b602093506155e49592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b820101386155d7565b6040519061567a82614e66565b60006020838281520152565b91908110156155295760051b0190565b3567ffffffffffffffff81168103611dd45790565b9190811015615529576060020190565b9067ffffffffffffffff90939293168152604060208201526157206156ec845160a0604085015260e0840190614fdc565b60208501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0848303016060850152614fdc565b906040840151917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08282030160808301526020808451928381520193019060005b8181106157c45750505060808473ffffffffffffffffffffffffffffffffffffffff6060615071969701511660a084015201519060c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc082850301910152614fdc565b8251805173ffffffffffffffffffffffffffffffffffffffff1686526020908101518187015260409095019490920191600101615761565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b6040519061584882614eea565b60006080838281528260208201528260408201528260608201520152565b9060405161587381614eea565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff16600052600760205261507160046040600020016155ab565b8181106158f9575050565b600081556001016158ee565b8181029291811591840414171561591857565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b90816020910312611dd457518015158103611dd45790565b805180156159cf57602003615991578051602082810191830183900312611dd457519060ff8211615991575060ff1690565b61141f906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190614fdc565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff821161591857565b60ff16604d811161591857600a0a90565b8115615a24570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff811692828414615b5957828411615b2f5790615a98916159f5565b91604d60ff8416118015615af6575b615ac057505090615aba61507192615a09565b90615905565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b50615b0083615a09565b8015615a24577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048411615aa7565b615b38916159f5565b91604d60ff841611615ac057505090615b5361507192615a09565b90615a1a565b5050505090565b73ffffffffffffffffffffffffffffffffffffffff600154163303615b8157565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b90805115615deb5767ffffffffffffffff81516020830120921691826000526007602052615be0816005604060002001616420565b15615da75760005260086020526040600020815167ffffffffffffffff8111614e8257615c0d8254615558565b601f8111615d75575b506020601f8211600114615caf5791615c89827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9593615c9f95600091615ca4575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055604051918291602083526020830190614fdc565b0390a2565b905084015138615c58565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b818110615d5d575092615c9f9492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610615d26575b5050811b01905561125e565b8501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880615d1a565b9192602060018192868a015181550194019201615cdf565b615da190836000526020600020601f840160051c8101916020851061092757601f0160051c01906158ee565b38615c16565b509061141f6040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190614fdc565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b67ffffffffffffffff166000818152600660205260409020549092919015615f175791615f1460e092615ee085615e6c7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b97616178565b846000526007602052615e83816040600020616618565b615e8c83616178565b846000526007602052615ea6836002604060002001616618565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000615f878330338461603d565b1690813b15611dd457604051907f42966c680000000000000000000000000000000000000000000000000000000082528160248160008096819560048401525af1801561238457615fd6575050565b81615fe091614f3e565b50565b9067ffffffffffffffff6155e49216600052600760205260406000209073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016916168da565b90919273ffffffffffffffffffffffffffffffffffffffff6155e49481604051957f23b872dd000000000000000000000000000000000000000000000000000000006020880152166024860152166044840152606483015260648252611cb7608483614f3e565b9190820391821161591857565b6160b961583b565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff8083511691616116602085019361611061610363ffffffff875116426160a4565b8560808901511690615905565b906168cd565b8082101561612f57505b16825263ffffffff4216905290565b9050616120565b9173ffffffffffffffffffffffffffffffffffffffff906040519260208401948552604084015216606082015260608152616172608082614f3e565b51902090565b80511561622c576fffffffffffffffffffffffffffffffff6040820151166fffffffffffffffffffffffffffffffff602083015116811090811591616223575b506161c05750565b606490616221604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b905015386161b8565b6fffffffffffffffffffffffffffffffff604082015116158015906162b4575b6162535750565b606490616221604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff602082015116151561624c565b9067ffffffffffffffff6155e49216600052600760205260026040600020019073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016916168da565b80548210156155295760005260206000200190600090565b80549068010000000000000000821015614e82578161636f9160016163a494018155616330565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055565b806000526003602052604060002054156000146163e1576163ca816002616348565b600254906000526003602052604060002055600190565b50600090565b806000526006602052604060002054156000146163e157616409816005616348565b600554906000526006602052604060002055600190565b6000828152600182016020526040902054616457578061644283600193616348565b80549260005201602052604060002055600190565b5050600090565b805480156164c2577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906164938282616330565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b9060018201918160005282602052604060002054908115156000146165c4577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918083116159185781547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101908111615918578361657b94820361658d575b50505061645e565b60005260205260006040812055600190565b6165ad61659d61636f9386616330565b90549060031b1c92839286616330565b905560005284602052604060002055388080616573565b50505050600090565b906040519182815491828252602082019060005260206000209260005b8181106165ff5750506155e492500383614f3e565b84548352600194850194879450602090930192016165ea565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1991616751606092805461665563ffffffff8260801c16426160a4565b9081616790575b50506fffffffffffffffffffffffffffffffff600181602086015116928281541680851060001461678857508280855b16167fffffffffffffffffffffffffffffffff000000000000000000000000000000008254161781556167058651151582907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b60408601517fffffffffffffffffffffffffffffffff0000000000000000000000000000000060809190911b16939092166fffffffffffffffffffffffffffffffff1692909217910155565b615f1460405180926fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b83809161668c565b6fffffffffffffffffffffffffffffffff916167c58392836167be6001880154948286169560801c90615905565b91166168cd565b8082101561684457505b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff9290911692909216167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116174260801b73ffffffff0000000000000000000000000000000016178155388061665c565b90506167cf565b7f00000000000000000000000000000000000000000000000000000000000000006168735750565b73ffffffffffffffffffffffffffffffffffffffff16806000526003602052604060002054156168a05750565b7fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b9190820180921161591857565b929192805460ff8160a01c16158015616bba575b616bb3576fffffffffffffffffffffffffffffffff8116906001830190815461693363ffffffff6fffffffffffffffffffffffffffffffff83169360801c16426160a4565b9081616b15575b5050848110616a9357508382106169c257507f1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a939450906fffffffffffffffffffffffffffffffff8061699085602096956160a4565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055604051908152a1565b8194506169d492505460801c926160a4565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81019080821161591857616a22616a279273ffffffffffffffffffffffffffffffffffffffff946168cd565b615a1a565b9216918215616a63577fd0c8d23a0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b7f15279c080000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b8473ffffffffffffffffffffffffffffffffffffffff8816918215616ae5577f1a76572a0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b7ff94ebcd10000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b828592939511616b8957616b30926161109160801c90615905565b80831015616b845750815b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff000000000000000000000000000000001617845591388061693a565b616b3b565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050509050565b5082156168ee565b73ffffffffffffffffffffffffffffffffffffffff616c51911691604092600080855193616bf08786614f3e565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602086015260208151910182855af13d15616cfa573d91616c3583614f7f565b92616c4287519485614f3e565b83523d6000602085013e616f2b565b80519081616c5e57505050565b602080616c6f938301019101615947565b15616c775750565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606091616f2b565b6000818152600360205260409020548015616457577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161591857600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161591857818103616d98575b505050616d84600261645e565b600052600360205260006040812055600190565b616dba616da961636f936002616330565b90549060031b1c9283926002616330565b90556000526003602052604060002055388080616d77565b6000818152600660205260409020548015616457577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161591857600554907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161591857818103616e68575b505050616e54600561645e565b600052600660205260006040812055600190565b616e8a616e7961636f936005616330565b90549060031b1c9283926005616330565b90556000526006602052604060002055388080616e47565b9060018201918160005282602052604060002054908115156000146165c4577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918083116159185781547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810190811161591857838161657b950361658d5750505061645e565b91929015616fa65750815115616f3f575090565b3b15616f485790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015616fb95750805190602001fd5b61141f906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190614fdc56fea164736f6c634300081a000a",
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) GetAllowListedFillers(opts *bind.CallOpts, remoteChainSelector uint64) ([]common.Address, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "getAllowListedFillers", remoteChainSelector)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) GetAllowListedFillers(remoteChainSelector uint64) ([]common.Address, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetAllowListedFillers(&_BurnMintFastTransferTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) GetAllowListedFillers(remoteChainSelector uint64) ([]common.Address, error) {
	return _BurnMintFastTransferTokenPool.Contract.GetAllowListedFillers(&_BurnMintFastTransferTokenPool.CallOpts, remoteChainSelector)
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCaller) IsFillerAllowListed(opts *bind.CallOpts, remoteChainSelector uint64, filler common.Address) (bool, error) {
	var out []interface{}
	err := _BurnMintFastTransferTokenPool.contract.Call(opts, &out, "isFillerAllowListed", remoteChainSelector, filler)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) IsFillerAllowListed(remoteChainSelector uint64, filler common.Address) (bool, error) {
	return _BurnMintFastTransferTokenPool.Contract.IsFillerAllowListed(&_BurnMintFastTransferTokenPool.CallOpts, remoteChainSelector, filler)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolCallerSession) IsFillerAllowListed(remoteChainSelector uint64, filler common.Address) (bool, error) {
	return _BurnMintFastTransferTokenPool.Contract.IsFillerAllowListed(&_BurnMintFastTransferTokenPool.CallOpts, remoteChainSelector, filler)
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactor) CcipSendToken(opts *bind.TransactOpts, feeToken common.Address, destinationChainSelector uint64, amount *big.Int, receiver []byte, extraArgs []byte) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.contract.Transact(opts, "ccipSendToken", feeToken, destinationChainSelector, amount, receiver, extraArgs)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) CcipSendToken(feeToken common.Address, destinationChainSelector uint64, amount *big.Int, receiver []byte, extraArgs []byte) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.CcipSendToken(&_BurnMintFastTransferTokenPool.TransactOpts, feeToken, destinationChainSelector, amount, receiver, extraArgs)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorSession) CcipSendToken(feeToken common.Address, destinationChainSelector uint64, amount *big.Int, receiver []byte, extraArgs []byte) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.CcipSendToken(&_BurnMintFastTransferTokenPool.TransactOpts, feeToken, destinationChainSelector, amount, receiver, extraArgs)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactor) FastFill(opts *bind.TransactOpts, fillRequestId [32]byte, sourceChainSelector uint64, srcAmountToFill *big.Int, sourceDecimals uint8, receiver common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.contract.Transact(opts, "fastFill", fillRequestId, sourceChainSelector, srcAmountToFill, sourceDecimals, receiver)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) FastFill(fillRequestId [32]byte, sourceChainSelector uint64, srcAmountToFill *big.Int, sourceDecimals uint8, receiver common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.FastFill(&_BurnMintFastTransferTokenPool.TransactOpts, fillRequestId, sourceChainSelector, srcAmountToFill, sourceDecimals, receiver)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorSession) FastFill(fillRequestId [32]byte, sourceChainSelector uint64, srcAmountToFill *big.Int, sourceDecimals uint8, receiver common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.FastFill(&_BurnMintFastTransferTokenPool.TransactOpts, fillRequestId, sourceChainSelector, srcAmountToFill, sourceDecimals, receiver)
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactor) UpdateDestChainConfig(opts *bind.TransactOpts, destChainConfigArgs FastTransferTokenPoolAbstractDestChainConfigUpdateArgs) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.contract.Transact(opts, "updateDestChainConfig", destChainConfigArgs)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) UpdateDestChainConfig(destChainConfigArgs FastTransferTokenPoolAbstractDestChainConfigUpdateArgs) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.UpdateDestChainConfig(&_BurnMintFastTransferTokenPool.TransactOpts, destChainConfigArgs)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorSession) UpdateDestChainConfig(destChainConfigArgs FastTransferTokenPoolAbstractDestChainConfigUpdateArgs) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.UpdateDestChainConfig(&_BurnMintFastTransferTokenPool.TransactOpts, destChainConfigArgs)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactor) UpdateFillerAllowList(opts *bind.TransactOpts, destinationChainSelector uint64, addFillers []common.Address, removeFillers []common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.contract.Transact(opts, "updateFillerAllowList", destinationChainSelector, addFillers, removeFillers)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) UpdateFillerAllowList(destinationChainSelector uint64, addFillers []common.Address, removeFillers []common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.UpdateFillerAllowList(&_BurnMintFastTransferTokenPool.TransactOpts, destinationChainSelector, addFillers, removeFillers)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorSession) UpdateFillerAllowList(destinationChainSelector uint64, addFillers []common.Address, removeFillers []common.Address) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.UpdateFillerAllowList(&_BurnMintFastTransferTokenPool.TransactOpts, destinationChainSelector, addFillers, removeFillers)
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
	Bps                      uint16
	MaxFillAmountPerRequest  *big.Int
	DestinationPool          []byte
	AddFillers               []common.Address
	RemoveFillers            []common.Address
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
	FillRequestId [32]byte
	FillId        [32]byte
	Filler        common.Address
	DestAmount    *big.Int
	Receiver      common.Address
	Raw           types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterFastTransferFilled(opts *bind.FilterOpts, fillRequestId [][32]byte, fillId [][32]byte, filler []common.Address) (*BurnMintFastTransferTokenPoolFastTransferFilledIterator, error) {

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

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "FastTransferFilled", fillRequestIdRule, fillIdRule, fillerRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolFastTransferFilledIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "FastTransferFilled", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchFastTransferFilled(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolFastTransferFilled, fillRequestId [][32]byte, fillId [][32]byte, filler []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "FastTransferFilled", fillRequestIdRule, fillIdRule, fillerRule)
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
	FillRequestId    [32]byte
	DstChainSelector uint64
	Amount           *big.Int
	FastTransferFee  *big.Int
	Receiver         []byte
	Raw              types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterFastTransferRequested(opts *bind.FilterOpts, fillRequestId [][32]byte, dstChainSelector []uint64) (*BurnMintFastTransferTokenPoolFastTransferRequestedIterator, error) {

	var fillRequestIdRule []interface{}
	for _, fillRequestIdItem := range fillRequestId {
		fillRequestIdRule = append(fillRequestIdRule, fillRequestIdItem)
	}
	var dstChainSelectorRule []interface{}
	for _, dstChainSelectorItem := range dstChainSelector {
		dstChainSelectorRule = append(dstChainSelectorRule, dstChainSelectorItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "FastTransferRequested", fillRequestIdRule, dstChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolFastTransferRequestedIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "FastTransferRequested", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchFastTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolFastTransferRequested, fillRequestId [][32]byte, dstChainSelector []uint64) (event.Subscription, error) {

	var fillRequestIdRule []interface{}
	for _, fillRequestIdItem := range fillRequestId {
		fillRequestIdRule = append(fillRequestIdRule, fillRequestIdItem)
	}
	var dstChainSelectorRule []interface{}
	for _, dstChainSelectorItem := range dstChainSelector {
		dstChainSelectorRule = append(dstChainSelectorRule, dstChainSelectorItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "FastTransferRequested", fillRequestIdRule, dstChainSelectorRule)
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
	FillRequestId [32]byte
	Raw           types.Log
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) FilterFastTransferSettled(opts *bind.FilterOpts, fillRequestId [][32]byte) (*BurnMintFastTransferTokenPoolFastTransferSettledIterator, error) {

	var fillRequestIdRule []interface{}
	for _, fillRequestIdItem := range fillRequestId {
		fillRequestIdRule = append(fillRequestIdRule, fillRequestIdItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.FilterLogs(opts, "FastTransferSettled", fillRequestIdRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintFastTransferTokenPoolFastTransferSettledIterator{contract: _BurnMintFastTransferTokenPool.contract, event: "FastTransferSettled", logs: logs, sub: sub}, nil
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolFilterer) WatchFastTransferSettled(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolFastTransferSettled, fillRequestId [][32]byte) (event.Subscription, error) {

	var fillRequestIdRule []interface{}
	for _, fillRequestIdItem := range fillRequestId {
		fillRequestIdRule = append(fillRequestIdRule, fillRequestIdItem)
	}

	logs, sub, err := _BurnMintFastTransferTokenPool.contract.WatchLogs(opts, "FastTransferSettled", fillRequestIdRule)
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
	case _BurnMintFastTransferTokenPool.abi.Events["InvalidFill"].ID:
		return _BurnMintFastTransferTokenPool.ParseInvalidFill(log)
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

func (BurnMintFastTransferTokenPoolDestChainConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x391233e9b9f2fa088f64269fb39c6cf49aeeaad3183e75aafb2e2fd906f770e7")
}

func (BurnMintFastTransferTokenPoolDestinationPoolUpdated) Topic() common.Hash {
	return common.HexToHash("0xb760e03fa04c0e86fcff6d0046cdcf22fb5d5b6a17d1e6f890b3456e81c40fd8")
}

func (BurnMintFastTransferTokenPoolFastTransferFilled) Topic() common.Hash {
	return common.HexToHash("0xd6f70fb263bfe7d01ec6802b3c07b6bd32579760fe9fcb4e248a036debb8cdf1")
}

func (BurnMintFastTransferTokenPoolFastTransferRequested) Topic() common.Hash {
	return common.HexToHash("0x1199f551568004635134aaf2ae681cd3c00d4baf32ee26dd8b2b8d583b051d9a")
}

func (BurnMintFastTransferTokenPoolFastTransferSettled) Topic() common.Hash {
	return common.HexToHash("0xd5d4e34f92d581a906fa8f24b4d8c5639f8931ede3f4ad88f6aaf2d634d653c9")
}

func (BurnMintFastTransferTokenPoolFillerAllowListUpdated) Topic() common.Hash {
	return common.HexToHash("0xccc0f2211c115acfa175a7923abdeb4b0a7c376d1b9e43c74973efe83d7d9e22")
}

func (BurnMintFastTransferTokenPoolInvalidFill) Topic() common.Hash {
	return common.HexToHash("0xad64960fe1d28c88faed204b509e08b3c9e07d9c1cb84991addc205e6bfca42f")
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

	GetAllowListedFillers(opts *bind.CallOpts, remoteChainSelector uint64) ([]common.Address, error)

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

	IsFillerAllowListed(opts *bind.CallOpts, remoteChainSelector uint64, filler common.Address) (bool, error)

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

	CcipSendToken(opts *bind.TransactOpts, feeToken common.Address, destinationChainSelector uint64, amount *big.Int, receiver []byte, extraArgs []byte) (*types.Transaction, error)

	FastFill(opts *bind.TransactOpts, fillRequestId [32]byte, sourceChainSelector uint64, srcAmountToFill *big.Int, sourceDecimals uint8, receiver common.Address) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error)

	SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error)

	SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateDestChainConfig(opts *bind.TransactOpts, destChainConfigArgs FastTransferTokenPoolAbstractDestChainConfigUpdateArgs) (*types.Transaction, error)

	UpdateFillerAllowList(opts *bind.TransactOpts, destinationChainSelector uint64, addFillers []common.Address, removeFillers []common.Address) (*types.Transaction, error)

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

	FilterDestChainConfigUpdated(opts *bind.FilterOpts, destinationChainSelector []uint64) (*BurnMintFastTransferTokenPoolDestChainConfigUpdatedIterator, error)

	WatchDestChainConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolDestChainConfigUpdated, destinationChainSelector []uint64) (event.Subscription, error)

	ParseDestChainConfigUpdated(log types.Log) (*BurnMintFastTransferTokenPoolDestChainConfigUpdated, error)

	FilterDestinationPoolUpdated(opts *bind.FilterOpts, dst []uint64) (*BurnMintFastTransferTokenPoolDestinationPoolUpdatedIterator, error)

	WatchDestinationPoolUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolDestinationPoolUpdated, dst []uint64) (event.Subscription, error)

	ParseDestinationPoolUpdated(log types.Log) (*BurnMintFastTransferTokenPoolDestinationPoolUpdated, error)

	FilterFastTransferFilled(opts *bind.FilterOpts, fillRequestId [][32]byte, fillId [][32]byte, filler []common.Address) (*BurnMintFastTransferTokenPoolFastTransferFilledIterator, error)

	WatchFastTransferFilled(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolFastTransferFilled, fillRequestId [][32]byte, fillId [][32]byte, filler []common.Address) (event.Subscription, error)

	ParseFastTransferFilled(log types.Log) (*BurnMintFastTransferTokenPoolFastTransferFilled, error)

	FilterFastTransferRequested(opts *bind.FilterOpts, fillRequestId [][32]byte, dstChainSelector []uint64) (*BurnMintFastTransferTokenPoolFastTransferRequestedIterator, error)

	WatchFastTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolFastTransferRequested, fillRequestId [][32]byte, dstChainSelector []uint64) (event.Subscription, error)

	ParseFastTransferRequested(log types.Log) (*BurnMintFastTransferTokenPoolFastTransferRequested, error)

	FilterFastTransferSettled(opts *bind.FilterOpts, fillRequestId [][32]byte) (*BurnMintFastTransferTokenPoolFastTransferSettledIterator, error)

	WatchFastTransferSettled(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolFastTransferSettled, fillRequestId [][32]byte) (event.Subscription, error)

	ParseFastTransferSettled(log types.Log) (*BurnMintFastTransferTokenPoolFastTransferSettled, error)

	FilterFillerAllowListUpdated(opts *bind.FilterOpts, dst []uint64) (*BurnMintFastTransferTokenPoolFillerAllowListUpdatedIterator, error)

	WatchFillerAllowListUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolFillerAllowListUpdated, dst []uint64) (event.Subscription, error)

	ParseFillerAllowListUpdated(log types.Log) (*BurnMintFastTransferTokenPoolFillerAllowListUpdated, error)

	FilterInvalidFill(opts *bind.FilterOpts, fillRequestId [][32]byte, filler []common.Address) (*BurnMintFastTransferTokenPoolInvalidFillIterator, error)

	WatchInvalidFill(opts *bind.WatchOpts, sink chan<- *BurnMintFastTransferTokenPoolInvalidFill, fillRequestId [][32]byte, filler []common.Address) (event.Subscription, error)

	ParseInvalidFill(log types.Log) (*BurnMintFastTransferTokenPoolInvalidFill, error)

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
