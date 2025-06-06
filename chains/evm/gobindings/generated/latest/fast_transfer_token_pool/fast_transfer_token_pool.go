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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIBurnMintERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipReceive\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.Any2EVMMessage\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipSendToken\",\"inputs\":[{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"settlementId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"computeFillId\",\"inputs\":[{\"name\":\"settlementId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceAmountNetFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourceDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"fastFill\",\"inputs\":[{\"name\":\"fillId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"settlementId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceAmountNetFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourceDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAccumulatedPoolFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowedFillers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCcipSendTokenFee\",\"inputs\":[{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"settlementFeeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIFastTransferPool.Quote\",\"components\":[{\"name\":\"ccipSettlementFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fastTransferFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFastTransferTokenPoolAbstract.DestChainConfig\",\"components\":[{\"name\":\"maxFillAmountPerRequest\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fillerAllowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"fastTransferFillerFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fastTransferPoolFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"settlementOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destinationPool\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"customExtraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFillInfo\",\"inputs\":[{\"name\":\"fillId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFastTransferTokenPoolAbstract.FillInfo\",\"components\":[{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"enumIFastTransferPool.FillState\"},{\"name\":\"filler\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAllowedFiller\",\"inputs\":[{\"name\":\"filler\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRouter\",\"inputs\":[{\"name\":\"newRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateDestChainConfig\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structFastTransferTokenPoolAbstract.DestChainConfigUpdateArgs[]\",\"components\":[{\"name\":\"fillerAllowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"fastTransferFillerFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fastTransferPoolFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"settlementOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"maxFillAmountPerRequest\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationPool\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"customExtraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateFillerAllowList\",\"inputs\":[{\"name\":\"fillersToAdd\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"fillersToRemove\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawPoolFees\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigUpdated\",\"inputs\":[{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"fastTransferFillerFeeBps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"fastTransferPoolFeeBps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"maxFillAmountPerRequest\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"destinationPool\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"chainFamilySelector\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"},{\"name\":\"settlementOverheadGas\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fillerAllowlistEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestinationPoolUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"destinationPool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FastTransferFilled\",\"inputs\":[{\"name\":\"fillId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"settlementId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"filler\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"destAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FastTransferRequested\",\"inputs\":[{\"name\":\"destinationChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"fillId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"settlementId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sourceAmountNetFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"sourceDecimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"fastTransferFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FastTransferSettled\",\"inputs\":[{\"name\":\"fillId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"settlementId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"fillerReimbursementAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"poolFeeAccumulated\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"prevState\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIFastTransferPool.FillState\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FillerAllowListUpdated\",\"inputs\":[{\"name\":\"addFillers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"removeFillers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeeWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RouterUpdated\",\"inputs\":[{\"name\":\"oldRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyFilled\",\"inputs\":[{\"name\":\"fillId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AlreadySettled\",\"inputs\":[{\"name\":\"fillId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"FillerNotAllowlisted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"filler\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InsufficientPoolFees\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFillId\",\"inputs\":[{\"name\":\"fillId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRouter\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TransferAmountExceedsMaxFillAmount\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x61012080604052346103f357616e81803803809161001d8285610472565b8339810160a0828203126103f35781516001600160a01b038116908190036103f35761004b60208401610495565b60408401519091906001600160401b0381116103f35784019280601f850112156103f3578351936001600160401b03851161045c578460051b9060208201956100976040519788610472565b86526020808701928201019283116103f357602001905b828210610444575050506100d060806100c9606087016104a3565b95016104a3565b93331561043357600180546001600160a01b0319163317905581158015610422575b8015610411575b610400578160209160049360805260c0526040519283809263313ce56760e01b82525afa600091816103bf575b50610394575b5060a052600480546001600160a01b0319166001600160a01b0384169081179091558151151560e0819052909190610272575b501561025c5761010052604051616829908161065882396080518181816114f70152818161156e0152818161177701528181611f7501528181612c3701528181613240015281816139fc01528181613c35015281816141d601528181614230015281816149de015281816152340152818161559f015281816156a7015261587e015260a0518181816117bb01528181611dc501528181613ecf0152818161417f0152818161508f0152615112015260c051818181610bf10152818161160901528181611d0001528181612fde01528181613a970152613dd4015260e051818181610b9f015281816138040152616454015261010051816121010152f35b6335fdcccd60e21b600052600060045260246000fd5b90602090604051906102848383610472565b60008252600036813760e051156103835760005b82518110156102ff576001906001600160a01b036102b682866104b7565b5116856102c2826104f9565b6102cf575b505001610298565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a138856102c7565b5092905060005b815181101561037a576001906001600160a01b0361032482856104b7565b511680156103745784610336826105f7565b610344575b50505b01610306565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a1388461033b565b5061033e565b5050503861015f565b6335f4a7b360e01b60005260046000fd5b60ff1660ff82168181036103a8575061012c565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d6020116103f8575b816103db60209383610472565b810103126103f3576103ec90610495565b9038610126565b600080fd5b3d91506103ce565b6342bcdf7f60e11b60005260046000fd5b506001600160a01b038116156100f9565b506001600160a01b038516156100f2565b639b15e16f60e01b60005260046000fd5b60208091610451846104a3565b8152019101906100ae565b634e487b7160e01b600052604160045260246000fd5b601f909101601f19168101906001600160401b0382119082101761045c57604052565b519060ff821682036103f357565b51906001600160a01b03821682036103f357565b80518210156104cb5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b80548210156104cb5760005260206000200190600090565b60008181526003602052604090205480156105f05760001981018181116105da576002546000198101919082116105da57818103610589575b5050506002548015610573576000190161054d8160026104e1565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b6105c261059a6105ab9360026104e1565b90549060031b1c92839260026104e1565b819391549060031b91821b91600019901b19161790565b90556000526003602052604060002055388080610532565b634e487b7160e01b600052601160045260246000fd5b5050600090565b80600052600360205260406000205415600014610651576002546801000000000000000081101561045c576106386105ab82600185940160025560026104e1565b9055600254906000526003602052604060002055600190565b5060009056fe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a7146142da57508063181f5a771461425457806321df0da714614203578063240028e8146141a357806324f65ee7146141655780632b2c0eb41461414a5780632e7aa8c814613d0657806339077537146139915780634c5ef0ed1461394c57806354c8a4f3146137d257806362ddd3c41461374f5780636609f599146137335780636d3d1a58146136ff5780636def4ce71461359d57806378b410f21461355657806379ba50971461348b5780637d54534e146133fe57806385572ffb14612d6057806387f060d014612a485780638926f54f14612a035780638a18dcbd146124625780638da5cb5b1461242e57806392575c3b14611c42578063929ea5ba14611b1e578063962d4020146119c85780639a4575b9146115265780639fe280f514611486578063a42a7b8b14611318578063a7cd63b7146112aa578063abe1c1e81461122e578063acfecf9114611109578063af58d59f146110bf578063b0f479a11461108b578063b794658014611053578063c0d7865514610f95578063c4bffe2b14610e65578063c75eea9c14610dbc578063cf7401f314610c15578063dc0bd97114610bc4578063e0351e1314610b87578063e8a1da1714610327578063eeebc674146102cf5763f2fde38b146101f857600080fd5b346102ca5760206003193601126102ca5773ffffffffffffffffffffffffffffffffffffffff6102266145aa565b61022e6152de565b163381146102a057807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346102ca5760806003193601126102ca5760443560ff811681036102ca5760643567ffffffffffffffff81116102ca5760209161031361031f92369060040161470c565b90602435600435614f91565b604051908152f35b346102ca576103353661475b565b9190926103406152de565b6000905b8282106109e25750505060009063ffffffff42165b81831061036257005b61036d838386614c3f565b92610120843603126102ca57604051936103868561443d565b61038f81614605565b8552602081013567ffffffffffffffff81116102ca5781019336601f860112156102ca5784356103be81614867565b956103cc60405197886144c9565b81875260208088019260051b820101903682116102ca5760208101925b8284106109b3575050505060208601948552604082013567ffffffffffffffff81116102ca5761041c903690840161470c565b90604087019182526104466104343660608601614949565b936060890194855260c0369101614949565b946080880195865261045884516159ca565b61046286516159ca565b825151156109895761047e67ffffffffffffffff895116616045565b156109505767ffffffffffffffff885116600052600760205260406000206105c085516fffffffffffffffffffffffffffffffff6040820151169061057b6fffffffffffffffffffffffffffffffff602083015116915115158360806040516104e68161443d565b858152602081018a905260408101849052606081018690520152855474ff000000000000000000000000000000000000000091151560a01b919091167fffffffffffffffffffffff0000000000000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff84161773ffffffff00000000000000000000000000000000608089901b1617178555565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176001830155565b6106e687516fffffffffffffffffffffffffffffffff604082015116906106a16fffffffffffffffffffffffffffffffff6020830151169151151583608060405161060a8161443d565b858152602081018a9052604081018490526060810186905201526002860180547fffffffffffffffffffffff000000000000000000000000000000000000000000166fffffffffffffffffffffffffffffffff85161773ffffffff0000000000000000000000000000000060808a901b161791151560a01b74ff000000000000000000000000000000000000000016919091179055565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176003830155565b6004845191019080519067ffffffffffffffff8211610921576107138261070d8554614b1e565b85614f4c565b602090601f831160011461087e57610761929160009183610873575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b60005b8751805182101561079c579061079660019261078f8367ffffffffffffffff8e511692614e3b565b5190615329565b01610767565b505097967f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c293929196509461086867ffffffffffffffff60019751169251935191516108346107ff60405196879687526101006020880152610100870190614567565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a1019192610359565b015190508d8061072f565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083169184600052816000209260005b81811061090957509084600195949392106108d2575b505050811b019055610764565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558c80806108c5565b929360206001819287860151815501950193016108af565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b67ffffffffffffffff8851167f1d5ad3c50000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b833567ffffffffffffffff81116102ca576020916109d7839283369187010161470c565b8152019301926103e9565b9092919367ffffffffffffffff610a026109fd868886614e4f565b614acc565b1692610a0d846165a4565b15610b5957836000526007602052610a2b6005604060002001615b83565b9260005b8451811015610a6757600190866000526007602052610a606005604060002001610a598389614e3b565b5190616674565b5001610a2f565b5093909491959250806000526007602052600560406000206000815560006001820155600060028201556000600382015560048101610aa68154614b1e565b9081610b16575b5050018054906000815581610af5575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a1019091610344565b6000526020600020908101905b81811015610abd5760008155600101610b02565b81601f60009311600114610b2e5750555b8880610aad565b81835260208320610b4991601f01861c810190600101614f22565b8082528160208120915555610b27565b837f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346102ca5760006003193601126102ca5760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b346102ca5760006003193601126102ca57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102ca5760e06003193601126102ca57610c2e6145ee565b60607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc3601126102ca57604051610c64816144ad565b60243580151581036102ca5781526044356fffffffffffffffffffffffffffffffff811681036102ca5760208201526064356fffffffffffffffffffffffffffffffff811681036102ca57604082015260607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c3601126102ca5760405190610ceb826144ad565b60843580151581036102ca57825260a4356fffffffffffffffffffffffffffffffff811681036102ca57602083015260c4356fffffffffffffffffffffffffffffffff811681036102ca57604083015273ffffffffffffffffffffffffffffffffffffffff6009541633141580610d9a575b610d6c57610d6a92615737565b005b7f8e4a23d6000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b5073ffffffffffffffffffffffffffffffffffffffff60015416331415610d5d565b346102ca5760206003193601126102ca5767ffffffffffffffff610dde6145ee565b610de6614e6f565b50166000526007602052610e61610e08610e036040600020614e9a565b61594b565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b0390f35b346102ca5760006003193601126102ca576040516005548082528160208101600560005260206000209260005b818110610f7c575050610ea7925003826144c9565b805190610ecc610eb683614867565b92610ec460405194856144c9565b808452614867565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe060208401920136833760005b8151811015610f2c578067ffffffffffffffff610f1960019385614e3b565b5116610f258287614e3b565b5201610efa565b5050906040519182916020830190602084525180915260408301919060005b818110610f59575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101610f4b565b8454835260019485019486945060209093019201610e92565b346102ca5760206003193601126102ca57610fae6145aa565b610fb66152de565b73ffffffffffffffffffffffffffffffffffffffff811690811561098957600480547fffffffffffffffffffffffff0000000000000000000000000000000000000000811690931790556040805173ffffffffffffffffffffffffffffffffffffffff93841681529190921660208201527f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f168491819081015b0390a1005b346102ca5760206003193601126102ca57610e616110776110726145ee565b614f00565b604051918291602083526020830190614567565b346102ca5760006003193601126102ca57602073ffffffffffffffffffffffffffffffffffffffff60045416604051908152f35b346102ca5760206003193601126102ca5767ffffffffffffffff6110e16145ee565b6110e9614e6f565b50166000526007602052610e61610e08610e036002604060002001614e9a565b346102ca5767ffffffffffffffff611120366147a9565b92909161112b6152de565b1690611144826000526006602052604060002054151590565b156112005781600052600760205261117560056040600020016111683686856146d5565b6020815191012090616674565b156111b9577f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d7691926111b4604051928392602084526020840191614dfc565b0390a2005b6111fc906040519384937f74f23c7c0000000000000000000000000000000000000000000000000000000085526004850152604060248501526044840191614dfc565b0390fd5b507f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346102ca5760206003193601126102ca57611247614a41565b50600435600052600d60205260408060002073ffffffffffffffffffffffffffffffffffffffff82519161127a83614459565b5461128860ff821684614c33565b81602084019160081c1681526112a1845180945161491f565b51166020820152f35b346102ca5760006003193601126102ca576040516002548082526020820190600260005260206000209060005b81811061130257610e61856112ee818703826144c9565b6040519182916020835260208301906147ea565b82548452602090930192600192830192016112d7565b346102ca5760206003193601126102ca5767ffffffffffffffff61133a6145ee565b1660005260076020526113536005604060002001615b83565b8051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061139961138384614867565b9361139160405195866144c9565b808552614867565b0160005b81811061147557505060005b81518110156113f157806113bf60019284614e3b565b5160005260086020526113d56040600020614b71565b6113df8286614e3b565b526113ea8185614e3b565b50016113a9565b826040518091602082016020835281518091526040830190602060408260051b8601019301916000905b82821061142a57505050500390f35b91936020611465827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851614567565b960192019201859493919261141b565b80606060208093870101520161139d565b346102ca5760206003193601126102ca5761149f6145aa565b6114a76152de565b6114af614995565b90816114b757005b602073ffffffffffffffffffffffffffffffffffffffff8261151b857f738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b859599957f00000000000000000000000000000000000000000000000000000000000000006158ec565b6040519485521692a2005b346102ca5761153436614834565b6060602060405161154481614459565b82815201526080810161155681614aab565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691160361197c57506020810177ffffffffffffffff000000000000000000000000000000006115bc82614acc565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156118ea5760009161194d575b506119235761165361164e60408401614aab565b616452565b67ffffffffffffffff61166582614acc565b1661167d816000526006602052604060002054151590565b156118f657602073ffffffffffffffffffffffffffffffffffffffff60045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa9081156118ea57600091611880575b5073ffffffffffffffffffffffffffffffffffffffff16330361185257611072816118219361172060606117166117b196614acc565b9201358092615651565b61172981615867565b7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff61175c84614acc565b6040805173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152336020820152908101949094521691606090a2614acc565b610e6160405160ff7f0000000000000000000000000000000000000000000000000000000000000000166020820152602081526117ef6040826144c9565b604051926117fc84614459565b8352602083019081526040519384936020855251604060208601526060850190614567565b90517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0848303016040850152614567565b7f728fe07b000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b6020813d6020116118e2575b81611899602093836144c9565b810103126118de57519073ffffffffffffffffffffffffffffffffffffffff821682036118db575073ffffffffffffffffffffffffffffffffffffffff6116e0565b80fd5b5080fd5b3d915061188c565b6040513d6000823e3d90fd5b7fa9902c7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f53ad11d80000000000000000000000000000000000000000000000000000000060005260046000fd5b61196f915060203d602011611975575b61196781836144c9565b810190615003565b8361163a565b503d61195d565b61199a73ffffffffffffffffffffffffffffffffffffffff91614aab565b7f961c9a4f000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b346102ca5760606003193601126102ca5760043567ffffffffffffffff81116102ca576119f990369060040161472a565b9060243567ffffffffffffffff81116102ca57611a1a9036906004016148ee565b9060443567ffffffffffffffff81116102ca57611a3b9036906004016148ee565b73ffffffffffffffffffffffffffffffffffffffff6009541633141580611afc575b610d6c57838614801590611af2575b611ac85760005b868110611a7c57005b80611ac2611a906109fd6001948b8b614e4f565b611a9b838989614e5f565b611abc611ab4611aac86898b614e5f565b923690614949565b913690614949565b91615737565b01611a73565b7f568efce20000000000000000000000000000000000000000000000000000000060005260046000fd5b5080861415611a6c565b5073ffffffffffffffffffffffffffffffffffffffff60015416331415611a5d565b346102ca5760406003193601126102ca5760043567ffffffffffffffff81116102ca57611b4f9036906004016148d3565b60243567ffffffffffffffff81116102ca57611b6f9036906004016148d3565b90611b786152de565b60005b8151811015611bb75780611bb073ffffffffffffffffffffffffffffffffffffffff611ba960019486614e3b565b511661600c565b5001611b7b565b5060005b8251811015611bf75780611bf073ffffffffffffffffffffffffffffffffffffffff611be960019487614e3b565b511661614f565b5001611bbb565b7ffd35c599d42a981cbb1bbf7d3e6d9855a59f5c994ec6b427118ee0c260e24193611c348361104e866040519384936040855260408501906147ea565b9083820360208501526147ea565b611c4b36614648565b5050929091604051611c5c81614459565b600081526000602082015260606080604051611c778161443d565b828152826020820152826040820152600083820152015267ffffffffffffffff8616916040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff000000000000000000000000000000008860801b16600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156118ea5760009161240f575b5061192357611d3f33616452565b611d56836000526006602052604060002054151590565b156123e15782600052600a602052604060002094855485116123ae5760018601549461ffff8660081c169661ffff8760181c169963ffffffff611da3611d9d8d8c87615b11565b90614ff6565b602088019081529860281c168061234a5750611dc160038301614b71565b955b7f000000000000000000000000000000000000000000000000000000000000000096604051611df18161443d565b858152602081019b8c52604081019d8e52606081019d60ff8a169e8f815236611e1b908a8c6146d5565b91608084019283526040519e8f9460208601602090525160408601525161ffff1660608501525161ffff1660808401525160ff1660a08301525160c0820160a0905260e08201611e6a91614567565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018c52611e9a908c6144c9565b6020809c604051611eab83826144c9565b60008082529d73ffffffffffffffffffffffffffffffffffffffff93611edf60026040519a611ed98c61443d565b01614b71565b895288015260408701521690816060860152608085015273ffffffffffffffffffffffffffffffffffffffff600454168c60405180927f20487ded0000000000000000000000000000000000000000000000000000000082528180611f488a8a60048401614cae565b03915afa90811561233f57908d949392918d91612302575b508252611f6d8684615651565b611f998630337f00000000000000000000000000000000000000000000000000000000000000006156cf565b611fa286615867565b806120ee575b505091611ffe9273ffffffffffffffffffffffffffffffffffffffff60045416906040518095819482937f96f4e9f900000000000000000000000000000000000000000000000000000000845260048401614cae565b039134905af19788156120e2578098612090575b5050869798927f240a1286fd41f1034c4032dcd6b93fc09e81be4a0b64c7ecee6260b605a8e016949261206161204d612085948a5190614def565b946120593684866146d5565b90868c614f91565b975160405195869586528c8601526040850152608060608501526080840191614dfc565b0390a4604051908152f35b909197508882813d83116120db575b6120a981836144c9565b810103126118db57505195827f240a1286fd41f1034c4032dcd6b93fc09e81be4a0b64c7ecee6260b605a8e016612012565b503d61209f565b604051903d90823e3d90fd5b909192506120ff82513033846156cf565b7f000000000000000000000000000000000000000000000000000000000000000091519182158015612257575b156121d357611ffe94928d94926121c66121cb9361219a6040519485927f095ea7b3000000000000000000000000000000000000000000000000000000008b850152602484016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018452836144c9565b615e15565b90928d611fa8565b60848d604051907f08c379a00000000000000000000000000000000000000000000000000000000082526004820152603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527f20746f206e6f6e2d7a65726f20616c6c6f77616e6365000000000000000000006064820152fd5b506040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff821660248201528d81604481865afa80156122f7578e8e80926122c1575b505090501561212c565b8184929350903d84116122ef575b6122d982846144c9565b50810103126122eb5751808e8e6122b7565b8c80fd5b3d91506122cf565b6040513d8f823e3d90fd5b8591508d8193949596903d8411612337575b61231e82846144c9565b508101031261233357908c939291518f611f60565b8b80fd5b3d9150612314565b6040513d8e823e3d90fd5b6040519061235782614459565b81526020810160018152604051917f181dcf100000000000000000000000000000000000000000000000000000000060208401525160248301525115156044820152604481526123a86064826144c9565b95611dc3565b5050507f58dd87c50000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b827fa9902c7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b612428915060203d6020116119755761196781836144c9565b88611d31565b346102ca5760006003193601126102ca57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346102ca5760206003193601126102ca5760043567ffffffffffffffff81116102ca5761249390369060040161472a565b61249b6152de565b60005b8181106124a757005b6124b2818385614c3f565b60a081017f1e10bdc4000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000612501836155f7565b16146129c2575b6020820161251581615635565b90604084019161ffff8061252885615635565b1691160161ffff81116129935761ffff612710911611612969576080840167ffffffffffffffff61255882614acc565b16600052600a60205260406000209460e08101946125768683614a5a565b600289019167ffffffffffffffff8211610921576125988261070d8554614b1e565b600090601f83116001146128c9576125e59291600091836128be5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b6125f184615635565b92600188019788549861260388615635565b60181b64ffff000000169561261786615644565b151560c087013597888555606088019c6126308e615624565b60281b68ffffffff0000000000169360081b62ffff0016907fffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000016177fffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffffff16179060ff161717905561010084016126a59085614a5a565b90916003019167ffffffffffffffff8211610921576126c88261070d8554614b1e565b600090601f831160011461280d57918061271b926127229695946000926128025750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055614acc565b9361272c90615635565b9461273690615635565b956127419083614a5a565b909161274c906155f7565b9761275690615624565b9261276090615644565b936040519761ffff899816885261ffff16602088015260408701526060860160e0905260e086019061279192614dfc565b957fffffffff0000000000000000000000000000000000000000000000000000000016608085015263ffffffff1660a0840152151560c083015267ffffffffffffffff1692037f6cfec31453105612e33aed8011f0e249b68d55e4efa65374322eb7ceeee76fbd91a260010161249e565b01359050388061072f565b838252602082209a9e9d9c9b9a917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416815b8181106128a65750919e9f9b9c9d9e6001939185612722989796941061286e575b505050811b019055614acc565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c199101351690558f8080612861565b91936020600181928787013581550195019201612840565b013590508e8061072f565b83825260208220917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416815b8181106129515750908460019594939210612919575b505050811b0190556125e8565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c199101351690558d808061290c565b838301358555600190940193602092830192016128f6565b7f382c09820000000000000000000000000000000000000000000000000000000060005260046000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b63ffffffff6129d360608401615624565b1615612508577f382c09820000000000000000000000000000000000000000000000000000000060005260046000fd5b346102ca5760206003193601126102ca576020612a3e67ffffffffffffffff612a2a6145ee565b166000526006602052604060002054151590565b6040519015158152f35b346102ca5760c06003193601126102ca5760043560243560443567ffffffffffffffff8116918282036102ca57606435926084359160ff831683036102ca5760a4359273ffffffffffffffffffffffffffffffffffffffff8416928385036102ca5780600052600a60205260ff60016040600020015416612d14575b50612ae860405184602082015260208152612ae06040826144c9565b828885614f91565b8703612ce65786600052600d602052604060002073ffffffffffffffffffffffffffffffffffffffff60405191612b1e83614459565b54612b2c60ff821684614c33565b60081c16602082015251956003871015612cb757600096612c8b57612b5b91612b549161510f565b8095615546565b60405195612b6887614459565b600187526020870196338852818752600d602052604087209051976003891015612c5e578798612c5b985060ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008454169116178255517fffffffffffffffffffffff0000000000000000000000000000000000000000ff74ffffffffffffffffffffffffffffffffffffffff0083549260081b1691161790556040519285845260208401527fd6f70fb263bfe7d01ec6802b3c07b6bd32579760fe9fcb4e248a036debb8cdf160403394a4337f00000000000000000000000000000000000000000000000000000000000000006156cf565b80f35b6024887f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b602487897fcee81443000000000000000000000000000000000000000000000000000000008252600452fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b867fcb537aa40000000000000000000000000000000000000000000000000000000060005260045260246000fd5b612d2b33600052600c602052604060002054151590565b612ac4577f6c46a9b5000000000000000000000000000000000000000000000000000000006000526004523360245260446000fd5b346102ca57612d6e36614834565b73ffffffffffffffffffffffffffffffffffffffff6004541633036133d05760a0813603126102ca57604051612da38161443d565b81358152612db360208301614605565b9060208101918252604083013567ffffffffffffffff81116102ca57612ddc903690850161470c565b9060408101918252606084013567ffffffffffffffff81116102ca57612e05903690860161470c565b936060820194855260808101359067ffffffffffffffff82116102ca570136601f820112156102ca578035612e3981614867565b91612e4760405193846144c9565b81835260208084019260061b820101903682116102ca57602001915b818310613398575050506080820152600092519067ffffffffffffffff821690519251945191825183019560208701936020818903126133945760208101519067ffffffffffffffff821161339057019660a0908890031261338c5760405193612ecc8561443d565b60208801518552612edf60408901615537565b9160208601928352612ef360608a01615537565b936040870194855260808a01519960ff8b168b0361338857606088019a8b5260a08101519067ffffffffffffffff82116133845790602091010183601f8201121561338857805190612f448261450a565b94612f5260405196876144c9565b828652602083830101116133845790612f719160208087019101614544565b6080870192835277ffffffffffffffff00000000000000000000000000000000604051917f2cbc26bb00000000000000000000000000000000000000000000000000000000835260801b16600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561337957899161335a575b506133325761301d8186614ae1565b156132f457509660ff61306f6130c09373ffffffffffffffffffffffffffffffffffffffff61309e999a9b516020815191015190602081106132c4575b50169561ffff808a5193511691511691615b11565b613099613084899a939a51858751169061510f565b99613092858751168461510f565b9951614def565b614def565b91511660405191846020840152602083526130ba6040846144c9565b88614f91565b93848752600d6020526040872091604051926130db84614459565b546130e960ff821685614c33565b73ffffffffffffffffffffffffffffffffffffffff602085019160081c168152889584516003811015613297576131e0575050906131319161312c828a96615546565b61521c565b838652600d6020526040862060027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055519060038210156131b357916131af6060927f33e17439bb4d31426d9168fc32af3a69cfce0467ba0d532fa804c27b5ff2189c946040519384526020840152604083019061491f565ba380f35b6024867f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b945094505081516003811015612c5e5760010361326b576132168373ffffffffffffffffffffffffffffffffffffffff92614def565b93511661322c6132268486614ff6565b3061521c565b838061323a575b5050613131565b613264917f00000000000000000000000000000000000000000000000000000000000000006158ec565b8683613233565b602487867fb196a44a000000000000000000000000000000000000000000000000000000008252600452fd5b60248b7f4e487b710000000000000000000000000000000000000000000000000000000081526021600452fd5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b168d61305a565b6111fc906040519182917f24eb47e5000000000000000000000000000000000000000000000000000000008352602060048401526024830190614567565b6004887f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b613373915060203d6020116119755761196781836144c9565b8a61300e565b6040513d8b823e3d90fd5b8a80fd5b8980fd5b8580fd5b8780fd5b8680fd5b6040833603126102ca57602060409182516133b281614459565b6133bb866145cd565b81528286013583820152815201920191612e63565b7fd7f73334000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b346102ca5760206003193601126102ca577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174602073ffffffffffffffffffffffffffffffffffffffff61344f6145aa565b6134576152de565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000006009541617600955604051908152a1005b346102ca5760006003193601126102ca5760005473ffffffffffffffffffffffffffffffffffffffff8116330361352c577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346102ca5760206003193601126102ca576020612a3e73ffffffffffffffffffffffffffffffffffffffff6135896145aa565b16600052600c602052604060002054151590565b346102ca5760206003193601126102ca5767ffffffffffffffff6135bf6145ee565b606060c06040516135cf81614491565b600081526000602082015260006040820152600083820152600060808201528260a0820152015216600052600a60205260606040600020610e61613611615b38565b611c3460405161362081614491565b845481526136cd60018601549563ffffffff602084019760ff81161515895261ffff60408601818360081c168152818c880191818560181c1683528560808a019560281c168552613686600361367860028a01614b71565b9860a08c01998a5201614b71565b9860c08101998a526040519e8f9e8f9260408452516040840152511515910152511660808c0152511660a08a0152511660c08801525160e080880152610120870190614567565b90517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc086830301610100870152614567565b346102ca5760006003193601126102ca57602073ffffffffffffffffffffffffffffffffffffffff60095416604051908152f35b346102ca5760006003193601126102ca57610e616112ee615b38565b346102ca5761375d366147a9565b6137689291926152de565b67ffffffffffffffff821661378a816000526006602052604060002054151590565b156137a55750610d6a9261379f9136916146d5565b90615329565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346102ca576137fa6138026137e63661475b565b94916137f39391936152de565b369161487f565b92369161487f565b7f0000000000000000000000000000000000000000000000000000000000000000156139225760005b825181101561389e578073ffffffffffffffffffffffffffffffffffffffff61385660019386614e3b565b5116613861816164d4565b61386d575b500161382b565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a184613866565b5060005b8151811015610d6a578073ffffffffffffffffffffffffffffffffffffffff6138cd60019385614e3b565b5116801561391c576138de81615fcd565b6138eb575b505b016138a2565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a1836138e3565b506138e5565b7f35f4a7b30000000000000000000000000000000000000000000000000000000060005260046000fd5b346102ca5760406003193601126102ca576139656145ee565b60243567ffffffffffffffff81116102ca5760209161398b612a3e92369060040161470c565b90614ae1565b346102ca5760206003193601126102ca5760043567ffffffffffffffff81116102ca57806004019061010060031982360301126102ca5760006040516139d681614475565b52608481016139e481614aab565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691160361197c57506024810177ffffffffffffffff00000000000000000000000000000000613a4a82614acc565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156118ea57600091613ce7575b506119235767ffffffffffffffff613adf82614acc565b16613af7816000526006602052604060002054151590565b156118f657602073ffffffffffffffffffffffffffffffffffffffff60045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa9081156118ea57600091613cc8575b501561185257613b6f81614acc565b613b8b60a484019161398b613b848488614a5a565b36916146d5565b15613c8157507ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0608067ffffffffffffffff613c18613c126044613c0288613bfc613bf7613b8460209d60c4613be08e614acc565b95613bf060648201358098615546565b0190614a5a565b61501b565b9061510f565b9701956109fd8861312c89614aab565b94614aab565b9373ffffffffffffffffffffffffffffffffffffffff60405195817f000000000000000000000000000000000000000000000000000000000000000016875233898801521660408601528560608601521692a280604051613c7881614475565b52604051908152f35b613c8b9084614a5a565b6111fc6040519283927f24eb47e5000000000000000000000000000000000000000000000000000000008452602060048501526024840191614dfc565b613ce1915060203d6020116119755761196781836144c9565b84613b60565b613d00915060203d6020116119755761196781836144c9565b84613ac8565b346102ca57613d1436614648565b505092939091613d22614a41565b5060405194613d3086614459565b600086526000602087015260606080604051613d4b8161443d565b828152826020820152826040820152600083820152015267ffffffffffffffff8316936040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff000000000000000000000000000000008560801b16600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156118ea5760009161412b575b5061192357613e1333616452565b613e2a856000526006602052604060002054151590565b156140fd5784600052600a602052604060002094855483116140cb57509263ffffffff92613f668693613f3a8a9760ff60016140099c9b01549161ffff8360081c1699836020613e87611d9d61ffff8f9860181c1680988b615b11565b9d019c8d5260281c1680614060575061ffff613ef8613ea860038c01614b71565b985b60405197613eb78961443d565b8852602088019c8d52604088019586526060880193857f000000000000000000000000000000000000000000000000000000000000000016855236916146d5565b9360808701948552816040519c8d986020808b01525160408a01525116606088015251166080860152511660a08401525160a060c084015260e0830190614567565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018652856144c9565b73ffffffffffffffffffffffffffffffffffffffff602096879560405190613f8e88836144c9565b60008252613fa4600260405198611ed98a61443d565b8752878701526040860152166060840152608083015273ffffffffffffffffffffffffffffffffffffffff60045416906040518097819482937f20487ded00000000000000000000000000000000000000000000000000000000845260048401614cae565b03915afa9283156118ea5760009361402e575b50826040945283519283525190820152f35b9392508184813d8311614059575b61404681836144c9565b810103126102ca5760409351929361401c565b503d61403c565b613ef861ffff916040519061407482614459565b81526020810160018152604051917f181dcf100000000000000000000000000000000000000000000000000000000060208401525160248301525115156044820152604481526140c56064826144c9565b98613eaa565b90507f58dd87c50000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b847fa9902c7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b614144915060203d6020116119755761196781836144c9565b88613e05565b346102ca5760006003193601126102ca57602061031f614995565b346102ca5760006003193601126102ca57602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102ca5760206003193601126102ca5760206141be6145aa565b73ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b346102ca5760006003193601126102ca57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102ca5760006003193601126102ca57610e616040516142766060826144c9565b602781527f4275726e4d696e74466173745472616e73666572546f6b656e506f6f6c20312e60208201527f362e312d646576000000000000000000000000000000000000000000000000006040820152604051918291602083526020830190614567565b346102ca5760206003193601126102ca57600435907fffffffff0000000000000000000000000000000000000000000000000000000082168092036102ca57817f61f8dc1600000000000000000000000000000000000000000000000000000000602093149081156143b2575b8115614355575b5015158152f35b7f85572ffb00000000000000000000000000000000000000000000000000000000811491508115614388575b508361434e565b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483614381565b90507faff2afbf0000000000000000000000000000000000000000000000000000000081148015614414575b80156143eb575b90614347565b507f01ffc9a70000000000000000000000000000000000000000000000000000000081146143e5565b507f0e64dd290000000000000000000000000000000000000000000000000000000081146143de565b60a0810190811067ffffffffffffffff82111761092157604052565b6040810190811067ffffffffffffffff82111761092157604052565b6020810190811067ffffffffffffffff82111761092157604052565b60e0810190811067ffffffffffffffff82111761092157604052565b6060810190811067ffffffffffffffff82111761092157604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761092157604052565b67ffffffffffffffff811161092157601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60005b8381106145575750506000910152565b8181015183820152602001614547565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f6020936145a381518092818752878088019101614544565b0116010190565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036102ca57565b359073ffffffffffffffffffffffffffffffffffffffff821682036102ca57565b6004359067ffffffffffffffff821682036102ca57565b359067ffffffffffffffff821682036102ca57565b9181601f840112156102ca5782359167ffffffffffffffff83116102ca57602083818601950101116102ca57565b9060a06003198301126102ca5760043567ffffffffffffffff811681036102ca57916024359160443567ffffffffffffffff81116102ca578261468d9160040161461a565b9290929160643573ffffffffffffffffffffffffffffffffffffffff811681036102ca57916084359067ffffffffffffffff82116102ca576146d19160040161461a565b9091565b9291926146e18261450a565b916146ef60405193846144c9565b8294818452818301116102ca578281602093846000960137010152565b9080601f830112156102ca57816020614727933591016146d5565b90565b9181601f840112156102ca5782359167ffffffffffffffff83116102ca576020808501948460051b0101116102ca57565b60406003198201126102ca5760043567ffffffffffffffff81116102ca57816147869160040161472a565b929092916024359067ffffffffffffffff82116102ca576146d19160040161472a565b9060406003198301126102ca5760043567ffffffffffffffff811681036102ca57916024359067ffffffffffffffff82116102ca576146d19160040161461a565b906020808351928381520192019060005b8181106148085750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016147fb565b60206003198201126102ca576004359067ffffffffffffffff82116102ca576003198260a0920301126102ca5760040190565b67ffffffffffffffff81116109215760051b60200190565b92919061488b81614867565b9361489960405195866144c9565b602085838152019160051b81019283116102ca57905b8282106148bb57505050565b602080916148c8846145cd565b8152019101906148af565b9080601f830112156102ca578160206147279335910161487f565b9181601f840112156102ca5782359167ffffffffffffffff83116102ca57602080850194606085020101116102ca57565b906003821015612cb75752565b35906fffffffffffffffffffffffffffffffff821682036102ca57565b91908260609103126102ca57604051614961816144ad565b809280359081151582036102ca57604061499091819385526149856020820161492c565b60208601520161492c565b910152565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156118ea57600091614a12575090565b90506020813d602011614a39575b81614a2d602093836144c9565b810103126102ca575190565b3d9150614a20565b60405190614a4e82614459565b60006020838281520152565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156102ca570180359067ffffffffffffffff82116102ca576020019181360383136102ca57565b3573ffffffffffffffffffffffffffffffffffffffff811681036102ca5790565b3567ffffffffffffffff811681036102ca5790565b9067ffffffffffffffff61472792166000526007602052600560406000200190602081519101209060019160005201602052604060002054151590565b90600182811c92168015614b67575b6020831014614b3857565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691614b2d565b9060405191826000825492614b8584614b1e565b8084529360018116908115614bf35750600114614bac575b50614baa925003836144c9565b565b90506000929192526020600020906000915b818310614bd7575050906020614baa9282010138614b9d565b6020919350806001915483858901015201910190918492614bbe565b60209350614baa9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138614b9d565b6003821015612cb75752565b9190811015614c7f5760051b810135907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee1813603018212156102ca570190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9067ffffffffffffffff9093929316815260406020820152614d13614cdf845160a0604085015260e0840190614567565b60208501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0848303016060850152614567565b906040840151917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08282030160808301526020808451928381520193019060005b818110614db75750505060808473ffffffffffffffffffffffffffffffffffffffff6060614727969701511660a084015201519060c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc082850301910152614567565b8251805173ffffffffffffffffffffffffffffffffffffffff1686526020908101518187015260409095019490920191600101614d54565b9190820391821161299357565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b8051821015614c7f5760209160051b010190565b9190811015614c7f5760051b0190565b9190811015614c7f576060020190565b60405190614e7c8261443d565b60006080838281528260208201528260408201528260608201520152565b90604051614ea78161443d565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff1660005260076020526147276004604060002001614b71565b818110614f2d575050565b60008155600101614f22565b8181029291811591840414171561299357565b9190601f8111614f5b57505050565b614baa926000526020600020906020601f840160051c83019310614f87575b601f0160051c0190614f22565b9091508190614f7a565b9290614fc4614ff09260ff60405195869460208601988952604086015216606084015260808084015260a0830190614567565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826144c9565b51902090565b9190820180921161299357565b908160209103126102ca575180151581036102ca5790565b8051801561508b5760200361504d5780516020828101918301839003126102ca57519060ff821161504d575060ff1690565b6111fc906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190614567565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff821161299357565b60ff16604d811161299357600a0a90565b81156150e0570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff811692828414615215578284116151eb5790615154916150b1565b91604d60ff84161180156151b2575b61517c57505090615176614727926150c5565b90614f39565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b506151bc836150c5565b80156150e0577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048411615163565b6151f4916150b1565b91604d60ff84161161517c5750509061520f614727926150c5565b906150d6565b5050505090565b9073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803b156102ca576040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff93909316600484015260248301919091526000919082908290604490829084905af180156152d3576152c6575050565b816152d0916144c9565b50565b6040513d84823e3d90fd5b73ffffffffffffffffffffffffffffffffffffffff6001541633036152ff57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b908051156109895767ffffffffffffffff8151602083012092169182600052600760205261535e81600560406000200161607e565b156154f35760005260086020526040600020815167ffffffffffffffff8111610921576153958161538f8454614b1e565b84614f4c565b6020601f821160011461542d5791615407827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea959361541d95600091615422575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055604051918291602083526020830190614567565b0390a2565b9050840151386153d6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b8181106154db57509261541d9492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9896106154a4575b5050811b019055611077565b8501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880615498565b9192602060018192868a01518155019401920161545d565b50906111fc6040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190614567565b519061ffff821682036102ca57565b67ffffffffffffffff7f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c911691826000526007602052806155c7600260406000200173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615bce565b6040805173ffffffffffffffffffffffffffffffffffffffff90921682526020820192909252908190810161541d565b357fffffffff00000000000000000000000000000000000000000000000000000000811681036102ca5790565b3563ffffffff811681036102ca5790565b3561ffff811681036102ca5790565b3580151581036102ca5790565b67ffffffffffffffff7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944911691826000526007602052806155c7604060002073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615bce565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff92831660248201529290911660448301526064820192909252614baa916121c6826084810161219a565b67ffffffffffffffff166000818152600660205260409020549092919015615839579161583660e0926158028561578e7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b976159ca565b8460005260076020526157a581604060002061621f565b6157ae836159ca565b8460005260076020526157c883600260406000200161621f565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690813b156102ca57604051907f42966c680000000000000000000000000000000000000000000000000000000082528160248160008096819560048401525af180156152d3576152c6575050565b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff90921660248301526044820192909252614baa916121c6826064810161219a565b615953614e6f565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff80835116916159aa6020850193611d9d61599d63ffffffff87511642614def565b8560808901511690614f39565b808210156159c357505b16825263ffffffff4216905290565b90506159b4565b805115615a6a576fffffffffffffffffffffffffffffffff6040820151166fffffffffffffffffffffffffffffffff60208301511610615a075750565b606490615a68604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff60408201511615801590615af2575b615a915750565b606490615a68604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff6020820151161515615a8a565b615b349061ffff612710615b2b8282969897981684614f39565b04951690614f39565b0490565b60405190600b548083528260208101600b60005260206000209260005b818110615b6a575050614baa925003836144c9565b8454835260019485019487945060209093019201615b55565b906040519182815491828252602082019060005260206000209260005b818110615bb5575050614baa925003836144c9565b8454835260019485019487945060209093019201615ba0565b9182549060ff8260a01c16158015615e0d575b615e07576fffffffffffffffffffffffffffffffff82169160018501908154615c2663ffffffff6fffffffffffffffffffffffffffffffff83169360801c1642614def565b9081615d69575b5050848110615d1d5750838310615c87575050615c5c6fffffffffffffffffffffffffffffffff928392614def565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b5460801c91615c968185614def565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81019080821161299357615ce4615ce99273ffffffffffffffffffffffffffffffffffffffff96614ff6565b6150d6565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828573ffffffffffffffffffffffffffffffffffffffff927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611615ddd57615d8492611d9d9160801c90614f39565b80841015615dd85750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178655923880615c2d565b615d8f565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b508215615be1565b73ffffffffffffffffffffffffffffffffffffffff615ea4911691604092600080855193615e4387866144c9565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602086015260208151910182855af13d15615f4d573d91615e888361450a565b92615e95875194856144c9565b83523d6000602085013e616750565b80519081615eb157505050565b602080615ec2938301019101615003565b15615eca5750565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606091616750565b8054821015614c7f5760005260206000200190600090565b805490680100000000000000008210156109215781615f94916001615fc994018155615f55565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055565b8060005260036020526040600020541560001461600657615fef816002615f6d565b600254906000526003602052604060002055600190565b50600090565b80600052600c602052604060002054156000146160065761602e81600b615f6d565b600b5490600052600c602052604060002055600190565b8060005260066020526040600020541560001461600657616067816005615f6d565b600554906000526006602052604060002055600190565b60008281526001820160205260409020546160b557806160a083600193615f6d565b80549260005201602052604060002055600190565b5050600090565b80548015616120577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906160f18282615f55565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6000818152600c602052604090205480156160b5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161299357600b54907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211612993578082036161e5575b5050506161d1600b6160bc565b600052600c60205260006040812055600190565b6162076161f6615f9493600b615f55565b90549060031b1c928392600b615f55565b9055600052600c6020526040600020553880806161c4565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1991616358606092805461625c63ffffffff8260801c1642614def565b9081616397575b50506fffffffffffffffffffffffffffffffff600181602086015116928281541680851060001461638f57508280855b16167fffffffffffffffffffffffffffffffff0000000000000000000000000000000082541617815561630c8651151582907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b60408601517fffffffffffffffffffffffffffffffff0000000000000000000000000000000060809190911b16939092166fffffffffffffffffffffffffffffffff1692909217910155565b61583660405180926fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b838091616293565b6fffffffffffffffffffffffffffffffff916163cc8392836163c56001880154948286169560801c90614f39565b9116614ff6565b8082101561644b57505b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff9290911692909216167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116174260801b73ffffffff00000000000000000000000000000000161781553880616263565b90506163d6565b7f000000000000000000000000000000000000000000000000000000000000000061647a5750565b73ffffffffffffffffffffffffffffffffffffffff16806000526003602052604060002054156164a75750565b7fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60008181526003602052604090205480156160b5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161299357600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116129935781810361656a575b50505061655660026160bc565b600052600360205260006040812055600190565b61658c61657b615f94936002615f55565b90549060031b1c9283926002615f55565b90556000526003602052604060002055388080616549565b60008181526006602052604090205480156160b5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161299357600554907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116129935781810361663a575b50505061662660056160bc565b600052600660205260006040812055600190565b61665c61664b615f94936005615f55565b90549060031b1c9283926005615f55565b90556000526006602052604060002055388080616619565b906001820191816000528260205260406000205490811515600014616747577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918083116129935781547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81019081116129935783816166fe9503616710575b5050506160bc565b60005260205260006040812055600190565b616730616720615f949386615f55565b90549060031b1c92839286615f55565b9055600052846020526040600020553880806166f6565b50505050600090565b919290156167cb5750815115616764575090565b3b1561676d5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156167de5750805190602001fd5b6111fc906040519182917f08c379a000000000000000000000000000000000000000000000000000000000835260206004840152602483019061456756fea164736f6c634300081a000a",
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

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactor) CcipSendToken(opts *bind.TransactOpts, destinationChainSelector uint64, amount *big.Int, receiver []byte, feeToken common.Address, extraArgs []byte) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.contract.Transact(opts, "ccipSendToken", destinationChainSelector, amount, receiver, feeToken, extraArgs)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolSession) CcipSendToken(destinationChainSelector uint64, amount *big.Int, receiver []byte, feeToken common.Address, extraArgs []byte) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.CcipSendToken(&_BurnMintFastTransferTokenPool.TransactOpts, destinationChainSelector, amount, receiver, feeToken, extraArgs)
}

func (_BurnMintFastTransferTokenPool *BurnMintFastTransferTokenPoolTransactorSession) CcipSendToken(destinationChainSelector uint64, amount *big.Int, receiver []byte, feeToken common.Address, extraArgs []byte) (*types.Transaction, error) {
	return _BurnMintFastTransferTokenPool.Contract.CcipSendToken(&_BurnMintFastTransferTokenPool.TransactOpts, destinationChainSelector, amount, receiver, feeToken, extraArgs)
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
	case _BurnMintFastTransferTokenPool.abi.Events["PoolFeeWithdrawn"].ID:
		return _BurnMintFastTransferTokenPool.ParsePoolFeeWithdrawn(log)
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
	return common.HexToHash("0x6cfec31453105612e33aed8011f0e249b68d55e4efa65374322eb7ceeee76fbd")
}

func (BurnMintFastTransferTokenPoolDestinationPoolUpdated) Topic() common.Hash {
	return common.HexToHash("0xb760e03fa04c0e86fcff6d0046cdcf22fb5d5b6a17d1e6f890b3456e81c40fd8")
}

func (BurnMintFastTransferTokenPoolFastTransferFilled) Topic() common.Hash {
	return common.HexToHash("0xd6f70fb263bfe7d01ec6802b3c07b6bd32579760fe9fcb4e248a036debb8cdf1")
}

func (BurnMintFastTransferTokenPoolFastTransferRequested) Topic() common.Hash {
	return common.HexToHash("0x240a1286fd41f1034c4032dcd6b93fc09e81be4a0b64c7ecee6260b605a8e016")
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
	ComputeFillId(opts *bind.CallOpts, settlementId [32]byte, sourceAmountNetFee *big.Int, sourceDecimals uint8, receiver []byte) ([32]byte, error)

	GetAccumulatedPoolFees(opts *bind.CallOpts) (*big.Int, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetAllowedFillers(opts *bind.CallOpts) ([]common.Address, error)

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

	CcipSendToken(opts *bind.TransactOpts, destinationChainSelector uint64, amount *big.Int, receiver []byte, feeToken common.Address, extraArgs []byte) (*types.Transaction, error)

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

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
