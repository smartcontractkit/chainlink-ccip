// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package offramp_over_superchain_interop

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

type InternalAny2EVMRampMessage struct {
	Header       InternalRampMessageHeader
	Sender       []byte
	Data         []byte
	Receiver     common.Address
	GasLimit     *big.Int
	TokenAmounts []InternalAny2EVMTokenTransfer
}

type InternalAny2EVMTokenTransfer struct {
	SourcePoolAddress []byte
	DestTokenAddress  common.Address
	DestGasAmount     uint32
	ExtraData         []byte
	Amount            *big.Int
}

type InternalExecutionReport struct {
	SourceChainSelector uint64
	Messages            []InternalAny2EVMRampMessage
	OffchainTokenData   [][][]byte
	Proofs              [][32]byte
	ProofFlagBits       *big.Int
}

type InternalGasPriceUpdate struct {
	DestChainSelector uint64
	UsdPerUnitGas     *big.Int
}

type InternalMerkleRoot struct {
	SourceChainSelector uint64
	OnRampAddress       []byte
	MinSeqNr            uint64
	MaxSeqNr            uint64
	MerkleRoot          [32]byte
}

type InternalPriceUpdates struct {
	TokenPriceUpdates []InternalTokenPriceUpdate
	GasPriceUpdates   []InternalGasPriceUpdate
}

type InternalRampMessageHeader struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	DestChainSelector   uint64
	SequenceNumber      uint64
	Nonce               uint64
}

type InternalTokenPriceUpdate struct {
	SourceToken common.Address
	UsdPerToken *big.Int
}

type MultiOCR3BaseConfigInfo struct {
	ConfigDigest                   [32]byte
	F                              uint8
	N                              uint8
	IsSignatureVerificationEnabled bool
}

type MultiOCR3BaseOCRConfig struct {
	ConfigInfo   MultiOCR3BaseConfigInfo
	Signers      []common.Address
	Transmitters []common.Address
}

type MultiOCR3BaseOCRConfigArgs struct {
	ConfigDigest                   [32]byte
	OcrPluginType                  uint8
	F                              uint8
	IsSignatureVerificationEnabled bool
	Signers                        []common.Address
	Transmitters                   []common.Address
}

type OffRampDynamicConfig struct {
	FeeQuoter                               common.Address
	PermissionLessExecutionThresholdSeconds uint32
	MessageInterceptor                      common.Address
}

type OffRampGasLimitOverride struct {
	ReceiverExecutionGasLimit *big.Int
	TokenGasOverrides         []uint32
}

type OffRampOverSuperchainInteropChainSelectorToChainIdConfigArgs struct {
	ChainSelector uint64
	ChainId       *big.Int
}

type OffRampSourceChainConfig struct {
	Router                    common.Address
	IsEnabled                 bool
	MinSeqNr                  uint64
	IsRMNVerificationDisabled bool
	OnRamp                    []byte
}

type OffRampSourceChainConfigArgs struct {
	Router                    common.Address
	SourceChainSelector       uint64
	IsEnabled                 bool
	IsRMNVerificationDisabled bool
	OnRamp                    []byte
}

type OffRampStaticConfig struct {
	ChainSelector        uint64
	GasForCallExactCheck uint16
	RmnRemote            common.Address
	TokenAdminRegistry   common.Address
	NonceManager         common.Address
}

var OffRampOverSuperchainInteropMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structOffRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structOffRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"sourceChainConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structOffRamp.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isRMNVerificationDisabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"crossL2Inbox\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainSelectorToChainIdConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structOffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[]\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainSelectorToChainIdConfigUpdates\",\"inputs\":[{\"name\":\"chainSelectorsToUnset\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainSelectorsToSet\",\"type\":\"tuple[]\",\"internalType\":\"structOffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[]\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySourceChainConfigUpdates\",\"inputs\":[{\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structOffRamp.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isRMNVerificationDisabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipReceive\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structClient.Any2EVMMessage\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"commit\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32[2]\",\"internalType\":\"bytes32[2]\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"reportContext\",\"type\":\"bytes32[2]\",\"internalType\":\"bytes32[2]\"},{\"name\":\"report\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSingleMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInternal.Any2EVMRampMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.RampMessageHeader\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Any2EVMTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasAmount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"offchainTokenData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"tokenGasOverrides\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllSourceChainConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOffRamp.SourceChainConfig[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"minSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isRMNVerificationDisabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChainId\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCrossL2Inbox\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOffRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExecutionState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumInternal.MessageExecutionState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestPriceSequenceNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMerkleRoot\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSourceChainConfig\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOffRamp.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"minSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isRMNVerificationDisabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOffRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestConfigDetails\",\"inputs\":[{\"name\":\"ocrPluginType\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"ocrConfig\",\"type\":\"tuple\",\"internalType\":\"structMultiOCR3Base.OCRConfig\",\"components\":[{\"name\":\"configInfo\",\"type\":\"tuple\",\"internalType\":\"structMultiOCR3Base.ConfigInfo\",\"components\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"F\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"n\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isSignatureVerificationEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"transmitters\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"manuallyExecute\",\"inputs\":[{\"name\":\"reports\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.ExecutionReport[]\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messages\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Any2EVMRampMessage[]\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.RampMessageHeader\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Any2EVMTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasAmount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"offchainTokenData\",\"type\":\"bytes[][]\",\"internalType\":\"bytes[][]\"},{\"name\":\"proofs\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"proofFlagBits\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"gasLimitOverrides\",\"type\":\"tuple[][]\",\"internalType\":\"structOffRamp.GasLimitOverride[][]\",\"components\":[{\"name\":\"receiverExecutionGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenGasOverrides\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structOffRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOCR3Configs\",\"inputs\":[{\"name\":\"ocrConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structMultiOCR3Base.OCRConfigArgs[]\",\"components\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ocrPluginType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"F\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isSignatureVerificationEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"transmitters\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AlreadyAttempted\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainSelectorToChainIdConfigRemoved\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainSelectorToChainIdConfigUpdated\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CommitReportAccepted\",\"inputs\":[{\"name\":\"blessedMerkleRoots\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structInternal.MerkleRoot[]\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"minSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"unblessedMerkleRoots\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structInternal.MerkleRoot[]\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"minSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"priceUpdates\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInternal.PriceUpdates\",\"components\":[{\"name\":\"tokenPriceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.TokenPriceUpdate[]\",\"components\":[{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"usdPerToken\",\"type\":\"uint224\",\"internalType\":\"uint224\"}]},{\"name\":\"gasPriceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.GasPriceUpdate[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"usdPerUnitGas\",\"type\":\"uint224\",\"internalType\":\"uint224\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"ocrPluginType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"configDigest\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"signers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"transmitters\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"F\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOffRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutionStateChanged\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"state\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"gasUsed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RootRemoved\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SkippedAlreadyExecutedMessage\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SkippedReportExecution\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SourceChainConfigSet\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sourceConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOffRamp.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"minSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isRMNVerificationDisabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SourceChainSelectorAdded\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOffRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transmitted\",\"inputs\":[{\"name\":\"ocrPluginType\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"configDigest\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySelfCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainIdMismatch\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expectedChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ChainIdNotConfiguredForSelector\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CommitOnRampMismatch\",\"inputs\":[{\"name\":\"reportOnRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"configOnRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ConfigDigestMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"EmptyBatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyReport\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ExecutionError\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ForkedChain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasToCompleteTx\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[{\"name\":\"errorType\",\"type\":\"uint8\",\"internalType\":\"enumMultiOCR3Base.InvalidConfigErrorType\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainSelector\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"expected\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidInterval\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"min\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"max\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidManualExecutionGasLimit\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"newLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidManualExecutionTokenGasOverride\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tokenIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"oldLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenGasOverride\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageCountInReport\",\"inputs\":[{\"name\":\"numMessages\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageDestChainSelector\",\"inputs\":[{\"name\":\"messageDestChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNewState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newState\",\"type\":\"uint8\",\"internalType\":\"enumInternal.MessageExecutionState\"}]},{\"type\":\"error\",\"name\":\"InvalidOnRampUpdate\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidProofsWordLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSourceChainSelector\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"expected\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidSourceOnRamp\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceOnRamp\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ManualExecutionGasAmountCountMismatch\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ManualExecutionGasLimitMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ManualExecutionNotYetEnabled\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MessageValidationError\",\"inputs\":[{\"name\":\"errorReason\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonUniqueSignatures\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotACompatiblePool\",\"inputs\":[{\"name\":\"notPool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperationNotSupportedByThisOffRampType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OracleCannotBeZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiverError\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ReleaseOrMintBalanceMismatch\",\"inputs\":[{\"name\":\"amountReleased\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePre\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePost\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"RootAlreadyCommitted\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"RootBlessingMismatch\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"isBlessed\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"type\":\"error\",\"name\":\"RootNotCommitted\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SignatureVerificationNotAllowedInExecutionPlugin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureVerificationRequiredInCommitPlugin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignaturesOutOfRegistration\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SourceChainNotEnabled\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SourceChainSelectorMismatch\",\"inputs\":[{\"name\":\"reportSourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageSourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"StaleCommitReport\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StaticConfigCannotBeChanged\",\"inputs\":[{\"name\":\"ocrPluginType\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"TokenDataMismatch\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenHandlingError\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedTransmitter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnexpectedTokenData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongMessageLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"WrongNumberOfSignatures\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainIdNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainSelectorNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610160806040523461091c57615eb3803803809161001d8285610a4f565b8339810181810390610160821261091c5760a0821261091c576040519161004383610a34565b61004c84610a72565b8352602084015161ffff8116810361091c57602084019081526040850151936001600160a01b038516850361091c576040810194855261008e60608701610a86565b946060820195865260606100a460808901610a86565b6080840190815294609f19011261091c5760405192606084016001600160401b03811185821017610615576040526100de60a08901610a86565b845260c08801519263ffffffff8416840361091c576020850193845261010660e08a01610a86565b604086019081526101008a01519096906001600160401b03811161091c578a019888601f8b01121561091c57895161013d81610a9a565b9a61014b6040519c8d610a4f565b818c526020808d019260051b820101908b821161091c5760208101925b82841061092157505050506101806101208c01610a86565b6101408c0151909b6001600160401b03821161091c570189601f8201121561091c578051906101ae82610a9a565b9a6101bc6040519c8d610a4f565b828c526020808d019360061b8301019181831161091c57602001925b8284106108d1575050505033156108c057600180546001600160a01b031916331790554660805284516001600160a01b03161580156108ae575b801561089c575b61062b5782516001600160401b0316156106c85782516001600160401b0390811660a090815286516001600160a01b0390811660c0528351811660e0528451811661010052865161ffff90811661012052604080519751909416875296519096166020860152955185169084015251831660608301525190911660808201527fb0fa1fb01508c5097c502ad056fd77018870c9be9a86d9e56b6b471862d7c5b79190a181516001600160a01b03161561062b57905160048054835163ffffffff60a01b60a09190911b166001600160a01b039384166001600160c01b03199092168217179091558351600580549184166001600160a01b031990921691909117905560408051918252925163ffffffff16602082015292511690820152839083907fa1c15688cb2c24508e158f6942b9276c6f3028a85e1af8cf3fff0c3ff3d5fc8d90606090a16000925b81518410156106d9576103778483610abe565b5160208101519094906001600160401b031680156106c85785516001600160a01b03161561062b57806000526008602052604060002060808701519660018201906103c28254610ae8565b610668578254600160a81b600160e81b031916600160a81b1783556040518481527ff4c1390c70e5c0f491ae1ccbc06f9117cbbadf2767b247b3bc203280f24c0fb990602090a15b8851801590811561063c575b5061062b5788516001600160401b038111610615576104358354610ae8565b601f81116105cd575b50602099601f821160011461054957926060600080516020615e938339815191529593836105329460019b9c9d9e60ff9860009261053e575b5050600019600383901b1c1916908b1b1783555b604081015115158554908760a01b9060a01b16908760a01b1916178555898060a01b038151168a8060a01b0319865416178555015115158354908560e81b9060e81b16908560e81b19161783556104e186610ba5565b506040519384936020855254898060a01b0381166020860152818160a01c1615156040860152898060401b038160a81c16606086015260e81c161515608084015260a08084015260c0830190610b22565b0390a201929190610364565b015190508f80610477565b99601f1982169a84600052816000209b60005b8181106105b5575093600184819b9c9d9e60ff989560609561053298600080516020615e938339815191529c9a1061059c575b505050811b01835561048b565b015160001960f88460031b161c191690558f808061058f565b828401518e556001909d019c6020938401930161055c565b836000526020600020601f830160051c8101916020841061060b575b601f0160051c01905b8181106105ff575061043e565b600081556001016105f2565b90915081906105e9565b634e487b7160e01b600052604160045260246000fd5b6342bcdf7f60e11b60005260046000fd5b905060208a01206040516020810190600082526020815261065e604082610a4f565b519020148a610416565b825460a81c6001600160401b0316600114158061069a575b1561040a5783632105803760e11b60005260045260246000fd5b506040516106b3816106ac8186610b22565b0382610a4f565b60208151910120895160208b01201415610680565b63c656089560e01b60005260046000fd5b6001600160a01b03831661014052604051602091906106f88382610a4f565b60008152600036813760005b815181101561077c576001906001600160401b036107228285610abe565b511680600052600c86526040600020549080600052600c87526000604081205581610750575b505001610704565b7fb56b587763154465d175e8a2a97978dffe45711125145973789b83ab201e702a600080a38580610748565b505060005b8151811015610813576107948183610abe565b51908382018051156108025782516001600160401b0316156106c8576001928151848060401b03825116600052600c8752604060002055838060401b039051169051907f5b4e1378b67677cfb7d6b50a37fd7f632168140928b0f970077a372bdfb42a3e600080a301610781565b63488d765160e01b60005260046000fd5b6040516152709081610c23823960805181612dd2015260a0518181816101cf015261466d015260c051818181610225015261383b015260e0518181816102540152613f350152610100518181816102830152613ba40152610120518181816101f6015281816123c0015281816140280152614dba0152610140518181816104df01526147450152f35b5081516001600160a01b031615610219565b5080516001600160a01b031615610212565b639b15e16f60e01b60005260046000fd5b60408483031261091c576040805191908201906001600160401b0382118383101761061557604092602092845261090787610a72565b815282870151838201528152019301926101d8565b600080fd5b83516001600160401b03811161091c57820160a0818f03601f19011261091c576040519061094e82610a34565b60208101516001600160a01b038116810361091c57825261097160408201610a72565b602083015261098260608201610ab1565b604083015261099360808201610ab1565b606083015260a0810151906001600160401b03821161091c57016020810190603f018f131561091c5780516001600160401b038111610615578f91604051926109e66020601f19601f8601160185610a4f565b8284526020838301011161091c5760005b828110610a1f5750509181600060208096949581960101526080820152815201930192610168565b806020809284010151828287010152016109f7565b60a081019081106001600160401b0382111761061557604052565b601f909101601f19168101906001600160401b0382119082101761061557604052565b51906001600160401b038216820361091c57565b51906001600160a01b038216820361091c57565b6001600160401b0381116106155760051b60200190565b5190811515820361091c57565b8051821015610ad25760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b90600182811c92168015610b18575b6020831014610b0257565b634e487b7160e01b600052602260045260246000fd5b91607f1691610af7565b60009291815491610b3283610ae8565b8083529260018116908115610b885750600114610b4e57505050565b60009081526020812093945091925b838310610b6e575060209250010190565b600181602092949394548385870101520191019190610b5d565b915050602093945060ff929192191683830152151560051b010190565b80600052600760205260406000205415600014610c1c5760065468010000000000000000811015610615576001810180600655811015610ad2577ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f0181905560065460009182526007602052604090912055600190565b5060009056fe6080604052600436101561001257600080fd5b60003560e01c806306285c6914610187578063181f5a77146101825780633242d5a91461017d5780633f4b04aa146101785780635215505b146101735780635e36480c1461016e5780635e7bb0081461016957806360987c201461016457806361ac8aac1461015f5780636f9e320f1461015a5780637437ff9f1461015557806379ba50971461015057806385572ffb1461014b5780638b6cecf8146101465780638da5cb5b14610141578063c673e5841461013c578063ccd37ba314610137578063cd19723714610132578063de5e0b9a1461012d578063e9d68a8e14610128578063f2fde38b14610123578063f58e03fc1461011e5763f716f99f1461011957600080fd5b611a7d565b611960565b6118d5565b611830565b611790565b611632565b6115d3565b61150e565b611426565b6113ec565b6113b6565b611336565b611296565b611121565b611072565b610f70565b610d69565b61078e565b61061f565b610503565b6104bf565b610460565b61019c565b600091031261019757565b600080fd5b34610197576000366003190112610197576101b5611bb8565b506102fd6040516101c581610317565b6001600160401b037f000000000000000000000000000000000000000000000000000000000000000016815261ffff7f00000000000000000000000000000000000000000000000000000000000000001660208201526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001660408201526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001660608201526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001660808201526040519182918291909160806001600160a01b038160a08401956001600160401b03815116855261ffff6020820151166020860152826040820151166040860152826060820151166060860152015116910152565b0390f35b634e487b7160e01b600052604160045260246000fd5b60a081019081106001600160401b0382111761033257604052565b610301565b604081019081106001600160401b0382111761033257604052565b606081019081106001600160401b0382111761033257604052565b608081019081106001600160401b0382111761033257604052565b90601f801991011681019081106001600160401b0382111761033257604052565b604051906103b860c083610388565b565b604051906103b860a083610388565b604051906103b861010083610388565b604051906103b8604083610388565b6001600160401b03811161033257601f01601f191660200190565b60405190610412602083610388565b60008252565b60005b83811061042b5750506000910152565b818101518382015260200161041b565b9060209161045481518092818552858086019101610418565b601f01601f1916010190565b34610197576000366003190112610197576102fd60408051906104838183610388565b600d82527f4f666652616d7020312e362e300000000000000000000000000000000000000060208301525191829160208352602083019061043b565b346101975760003660031901126101975760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101975760003660031901126101975760206001600160401b03600b5416604051908152f35b9060a0608061057b936001600160a01b0381511684526020810151151560208501526001600160401b036040820151166040850152606081015115156060850152015191816080820152019061043b565b90565b6040810160408252825180915260206060830193019060005b818110610600575050506020818303910152815180825260208201916020808360051b8301019401926000915b8383106105d357505050505090565b90919293946020806105f1600193601f19868203018752895161052a565b970193019301919392906105c4565b82516001600160401b0316855260209485019490920191600101610597565b346101975760003660031901126101975760065461063c816107d6565b9061064a6040519283610388565b808252601f19610659826107d6565b0160005b81811061071b57505061066f81611c11565b9060005b81811061068b5750506102fd6040519283928361057e565b806106c16106a961069d600194613700565b6001600160401b031690565b6106b38387611cab565b906001600160401b03169052565b6106ff6106fa6106e16106d48488611cab565b516001600160401b031690565b6001600160401b03166000526008602052604060002090565b611d97565b6107098287611cab565b526107148186611cab565b5001610673565b602090610726611be3565b8282870101520161065d565b600435906001600160401b038216820361019757565b35906001600160401b038216820361019757565b634e487b7160e01b600052602160045260246000fd5b6004111561077c57565b61075c565b90600482101561077c5752565b34610197576040366003190112610197576107a7610732565b602435906001600160401b0382168203610197576020916107c791611e40565b6107d46040518092610781565bf35b6001600160401b0381116103325760051b60200190565b91908260a09103126101975760405161080581610317565b608061084a8183958035855261081d60208201610748565b602086015261082e60408201610748565b604086015261083f60608201610748565b606086015201610748565b910152565b92919261085b826103e8565b916108696040519384610388565b829481845281830111610197578281602093846000960137010152565b9080601f830112156101975781602061057b9335910161084f565b6001600160a01b0381160361019757565b35906103b8826108a1565b63ffffffff81160361019757565b35906103b8826108bd565b81601f82011215610197578035906108ed826107d6565b926108fb6040519485610388565b82845260208085019360051b830101918183116101975760208101935b83851061092757505050505090565b84356001600160401b03811161019757820160a0818503601f190112610197576040519161095483610317565b60208201356001600160401b0381116101975785602061097692850101610886565b83526040820135610986816108a1565b6020840152610997606083016108cb565b60408401526080820135926001600160401b0384116101975760a0836109c4886020809881980101610886565b606084015201356080820152815201940193610918565b91909161014081840312610197576109f16103a9565b926109fc81836107ed565b845260a08201356001600160401b0381116101975781610a1d918401610886565b602085015260c08201356001600160401b0381116101975781610a41918401610886565b6040850152610a5260e083016108b2565b606085015261010082013560808501526101208201356001600160401b03811161019757610a8092016108d6565b60a0830152565b9080601f83011215610197578135610a9e816107d6565b92610aac6040519485610388565b81845260208085019260051b820101918383116101975760208201905b838210610ad857505050505090565b81356001600160401b03811161019757602091610afa878480948801016109db565b815201910190610ac9565b81601f8201121561019757803590610b1c826107d6565b92610b2a6040519485610388565b82845260208085019360051b830101918183116101975760208101935b838510610b5657505050505090565b84356001600160401b03811161019757820183603f82011215610197576020810135610b81816107d6565b91610b8f6040519384610388565b8183526020808085019360051b83010101918683116101975760408201905b838210610bc8575050509082525060209485019401610b47565b81356001600160401b03811161019757602091610bec8a8480809589010101610886565b815201910190610bae565b9080601f83011215610197578135610c0e816107d6565b92610c1c6040519485610388565b81845260208085019260051b82010192831161019757602001905b828210610c445750505090565b8135815260209182019101610c37565b81601f8201121561019757803590610c6b826107d6565b92610c796040519485610388565b82845260208085019360051b830101918183116101975760208101935b838510610ca557505050505090565b84356001600160401b03811161019757820160a0818503601f19011261019757610ccd6103ba565b91610cda60208301610748565b835260408201356001600160401b03811161019757856020610cfe92850101610a87565b602084015260608201356001600160401b03811161019757856020610d2592850101610b05565b60408401526080820135926001600160401b0384116101975760a083610d52886020809881980101610bf7565b606084015201356080820152815201940193610c96565b34610197576040366003190112610197576004356001600160401b03811161019757610d99903690600401610c54565b6024356001600160401b038111610197573660238201121561019757806004013591610dc4836107d6565b91610dd26040519384610388565b8383526024602084019460051b820101903682116101975760248101945b828610610e0357610e018585611e88565b005b85356001600160401b03811161019757820136604382011215610197576024810135610e2e816107d6565b91610e3c6040519384610388565b818352602060248185019360051b83010101903682116101975760448101925b828410610e76575050509082525060209586019501610df0565b83356001600160401b038111610197576024908301016040601f1982360301126101975760405190610ea782610337565b6020810135825260408101356001600160401b03811161019757602091010136601f8201121561019757803590610edd826107d6565b91610eeb6040519384610388565b80835260208084019160051b8301019136831161019757602001905b828210610f265750505091816020938480940152815201930192610e5c565b602080918335610f35816108bd565b815201910190610f07565b9181601f84011215610197578235916001600160401b038311610197576020808501948460051b01011161019757565b34610197576060366003190112610197576004356001600160401b03811161019757610fa09036906004016109db565b6024356001600160401b03811161019757610fbf903690600401610f40565b91604435926001600160401b03841161019757610fe3610e01943690600401610f40565b939092612298565b81601f8201121561019757803590611002826107d6565b926110106040519485610388565b82845260208085019360061b8301019181831161019757602001925b82841061103a575050505090565b604084830312610197576020604091825161105481610337565b61105d87610748565b8152828701358382015281520193019261102c565b34610197576040366003190112610197576004356001600160401b0381116101975736602382011215610197578060040135906110ae826107d6565b916110bc6040519384610388565b8083526024602084019160051b8301019136831161019757602401905b82821061110957602435846001600160401b03821161019757611103610e01923690600401610feb565b90612575565b6020809161111684610748565b8152019101906110d9565b3461019757606036600319011261019757600060405161114081610352565b60043561114c816108a1565b815260243561115a816108bd565b602082019081526044359061116e826108a1565b6040830191825261117d613088565b6001600160a01b038351161561128757916112496001600160a01b03611281937fa1c15688cb2c24508e158f6942b9276c6f3028a85e1af8cf3fff0c3ff3d5fc8d956111e2838651166001600160a01b03166001600160a01b03196004541617600455565b517fffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffff77ffffffff00000000000000000000000000000000000000006004549260a01b1691161760045551166001600160a01b03166001600160a01b03196005541617600555565b6040519182918291909160406001600160a01b0381606084019582815116855263ffffffff6020820151166020860152015116910152565b0390a180f35b6342bcdf7f60e11b8452600484fd5b34610197576000366003190112610197576000604080516112b681610352565b82815282602082015201526102fd6040516112d081610352565b63ffffffff6004546001600160a01b038116835260a01c1660208201526001600160a01b036005541660408201526040519182918291909160406001600160a01b0381606084019582815116855263ffffffff6020820151166020860152015116910152565b34610197576000366003190112610197576000546001600160a01b03811633036113a5576001600160a01b0319600154913382841617600155166000556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b63015aa1e360e11b60005260046000fd5b34610197576020366003190112610197576004356001600160401b0381116101975760a090600319903603011261019757600080fd5b34610197576020366003190112610197576001600160401b0361140d610732565b16600052600c6020526020604060002054604051908152f35b346101975760003660031901126101975760206001600160a01b0360015416604051908152f35b6004359060ff8216820361019757565b359060ff8216820361019757565b906020808351928381520192019060005b8181106114895750505090565b82516001600160a01b031684526020938401939092019160010161147c565b9061057b9160208152606082518051602084015260ff602082015116604084015260ff6040820151168284015201511515608082015260406114f9602084015160c060a085015260e084019061146b565b9201519060c0601f198285030191015261146b565b346101975760203660031901126101975760ff61152961144d565b60606040805161153881610352565b81516115438161036d565b6000815260006020820152600083820152600084820152815282602082015201521660005260026020526102fd604060002060036115c26040519261158784610352565b611590816126f7565b84526040516115ad816115a68160028601612730565b0382610388565b60208501526115a66040518094819301612730565b6040820152604051918291826114a8565b34610197576040366003190112610197576115ec610732565b6001600160401b036024359116600052600a6020526040600020906000526020526020604060002054604051908152f35b8015150361019757565b35906103b88261161d565b34610197576020366003190112610197576004356001600160401b03811161019757366023820112156101975780600401359061166e826107d6565b9061167c6040519283610388565b8282526024602083019360051b820101903682116101975760248101935b8285106116aa57610e0184612787565b84356001600160401b03811161019757820160a0602319823603011261019757604051916116d783610317565b60248201356116e5816108a1565b83526116f360448301610748565b602084015260648201356117068161161d565b604084015260848201356117198161161d565b606084015260a4820135926001600160401b03841161019757611746602094936024869536920101610886565b608082015281520194019361169a565b9060049160441161019757565b9181601f84011215610197578235916001600160401b038311610197576020838186019501011161019757565b346101975760c0366003190112610197576117aa36611756565b506044356001600160401b038111610197576117ca903690600401611763565b50506064356001600160401b038111610197576117eb903690600401610f40565b50506084356001600160401b0381116101975761180c903690600401610f40565b505063c4d744db60e01b60005260046000fd5b90602061057b92818152019061052a565b34610197576020366003190112610197576001600160401b03611851610732565b611859611be3565b501660005260086020526102fd60406000206118c460016040519261187d84610317565b6118be60ff82546001600160a01b0381168752818160a01c16151560208801526001600160401b038160a81c16604088015260e81c16606086019015159052565b01611d7c565b60808201526040519182918261181f565b34610197576020366003190112610197576001600160a01b036004356118fa816108a1565b611902613088565b1633811461194f57806001600160a01b031960005416176000556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b636d6c4ee560e11b60005260046000fd5b346101975760603660031901126101975761197a36611756565b6044356001600160401b03811161019757611999903690600401611763565b91828201602083820312610197578235906001600160401b038211610197576119c3918401610c54565b6040519060206119d38184610388565b60008352601f19810160005b818110611a0757505050610e0194916119f791612e13565b6119ff612a65565b928392613307565b606085820184015282016119df565b9080601f83011215610197578135611a2d816107d6565b92611a3b6040519485610388565b81845260208085019260051b82010192831161019757602001905b828210611a635750505090565b602080918335611a72816108a1565b815201910190611a56565b34610197576020366003190112610197576004356001600160401b038111610197573660238201121561019757806004013590611ab9826107d6565b90611ac76040519283610388565b8282526024602083019360051b820101903682116101975760248101935b828510611af557610e0184612aa3565b84356001600160401b03811161019757820160c0602319823603011261019757611b1d6103a9565b9160248201358352611b316044830161145d565b6020840152611b426064830161145d565b6040840152611b5360848301611627565b606084015260a48201356001600160401b03811161019757611b7b9060243691850101611a16565b608084015260c4820135926001600160401b03841161019757611ba8602094936024869536920101611a16565b60a0820152815201940193611ae5565b60405190611bc582610317565b60006080838281528260208201528260408201528260608201520152565b60405190611bf082610317565b60606080836000815260006020820152600060408201526000838201520152565b90611c1b826107d6565b611c286040519182610388565b8281528092611c39601f19916107d6565b0190602036910137565b634e487b7160e01b600052603260045260246000fd5b805115611c665760200190565b611c43565b805160011015611c665760400190565b805160021015611c665760600190565b805160031015611c665760800190565b805160041015611c665760a00190565b8051821015611c665760209160051b010190565b90600182811c92168015611cef575b6020831014611cd957565b634e487b7160e01b600052602260045260246000fd5b91607f1691611cce565b60009291815491611d0983611cbf565b8083529260018116908115611d5f5750600114611d2557505050565b60009081526020812093945091925b838310611d45575060209250010190565b600181602092949394548385870101520191019190611d34565b915050602093945060ff929192191683830152151560051b010190565b906103b8611d909260405193848092611cf9565b0383610388565b9060016080604051611da881610317565b611dfe819560ff81546001600160a01b0381168552818160a01c16151560208601526001600160401b038160a81c16604086015260e81c1615156060840152611df76040518096819301611cf9565b0384610388565b0152565b634e487b7160e01b600052601160045260246000fd5b908160051b9180830460201490151715611e2e57565b611e02565b91908203918211611e2e57565b611e4c82607f92612d8c565b9116906801fffffffffffffffe6001600160401b0383169260011b169180830460021490151715611e2e576003911c16600481101561077c5790565b611e90612dd0565b80518251810361208b5760005b818110611eb0575050906103b891612e13565b611eba8184611cab565b516020810190815151611ecd8488611cab565b51928351820361208b5790916000925b808410611ef1575050505050600101611e9d565b91949398611f03848b98939598611cab565b515198611f11888851611cab565b519980612042575b5060a08a01988b6020611f2f8b8d515193611cab565b51015151036120015760005b8a5151811015611fec57611f77611f6e611f648f6020611f5c8f8793611cab565b510151611cab565b5163ffffffff1690565b63ffffffff1690565b8b81611f88575b5050600101611f3b565b611f6e6040611f9b85611fa79451611cab565b51015163ffffffff1690565b90818110611fb657508b611f7e565b8d51516040516348e617b360e01b81526004810191909152602481019390935260448301919091526064820152608490fd5b0390fd5b50985098509893949095600101929091611edd565b61203e8b5161201c606082519201516001600160401b031690565b6370a193fd60e01b6000526004919091526001600160401b0316602452604490565b6000fd5b60808b0151811015611f195761203e908b61206488516001600160401b031690565b905151633a98d46360e11b6000526001600160401b03909116600452602452604452606490565b6320f8fd5960e21b60005260046000fd5b604051906120a982610337565b60006020838281520152565b604051906120c4602083610388565b600080835282815b8281106120d857505050565b6020906120e361209c565b828285010152016120cc565b805182526001600160401b0360208201511660208301526080612136612124604084015160a0604087015260a086019061043b565b6060840151858203606087015261043b565b9101519160808183039101526020808351928381520192019060005b81811061215f5750505090565b825180516001600160a01b031685526020908101518186015260409094019390920191600101612152565b90602061057b9281815201906120ef565b6040513d6000823e3d90fd5b3d156121d2573d906121b8826103e8565b916121c66040519384610388565b82523d6000602084013e565b606090565b90602061057b92818152019061043b565b90916060828403126101975781516121ff8161161d565b9260208301516001600160401b0381116101975783019080601f830112156101975781519161222d836103e8565b9161223b6040519384610388565b838352602084830101116101975760409261225c9160208085019101610418565b92015190565b9293606092959461ffff6122866001600160a01b03946080885260808801906120ef565b97166020860152604085015216910152565b92909391303303612564576122ab6120b5565b9460a0850151805161251d575b50505050508051916122d6602084519401516001600160401b031690565b9060208301519160408401926123038451926122f06103ba565b9788526001600160401b03166020880152565b6040860152606085015260808401526001600160a01b0361232c6005546001600160a01b031690565b16806124a0575b5051511580612494575b801561247e575b8015612455575b612451576123e9918161238e6123826123756106e1602060009751016001600160401b0390511690565b546001600160a01b031690565b6001600160a01b031690565b90836123a9606060808401519301516001600160a01b031690565b604051633cf9798360e01b815296879586948593917f00000000000000000000000000000000000000000000000000000000000000009060048601612262565b03925af190811561244c57600090600092612425575b50156124085750565b6040516302a35ba360e21b8152908190611fe890600483016121d7565b905061244491503d806000833e61243c8183610388565b8101906121e8565b5090386123ff565b61219b565b5050565b5061247961247561247060608401516001600160a01b031690565b61303a565b1590565b61234b565b5060608101516001600160a01b03163b15612344565b5060808101511561233d565b803b1561019757600060405180926308d450a160e01b82528183816124c88a6004830161218a565b03925af19081612502575b506124fc57611fe86124e36121a7565b6040516309c2532560e01b8152918291600483016121d7565b38612333565b80612511600061251793610388565b8061018c565b386124d3565b859650602061255996015161253c60608901516001600160a01b031690565b9061255360208a51016001600160401b0390511690565b92612f21565b9038808080806122b8565b6306e34e6560e31b60005260046000fd5b9161257e613088565b60005b835181101561261c578061259a6106d460019387611cab565b6125b7816001600160401b0316600052600c602052604060002090565b549060006125d8826001600160401b0316600052600c602052604060002090565b55816125e7575b505001612581565b6001600160401b03167fb56b587763154465d175e8a2a97978dffe45711125145973789b83ab201e702a600080a338806125df565b50915060005b8151811015612451576126358183611cab565b5190602082018051156126e65761265661069d84516001600160401b031690565b156126d5576001600160401b036126a6600194835161269861267f83516001600160401b031690565b6001600160401b0316600052600c602052604060002090565b55516001600160401b031690565b915191167f5b4e1378b67677cfb7d6b50a37fd7f632168140928b0f970077a372bdfb42a3e600080a301612622565b63c656089560e01b60005260046000fd5b63488d765160e01b60005260046000fd5b906040516127048161036d565b606060ff600183958054855201548181166020850152818160081c16604085015260101c161515910152565b906020825491828152019160005260206000209060005b8181106127545750505090565b82546001600160a01b0316845260209093019260019283019201612747565b906103b8611d909260405193848092612730565b61278f613088565b60005b8151811015612451576127a58183611cab565b51906127bb60208301516001600160401b031690565b6001600160401b0381169081156126d5576127e361238261238286516001600160a01b031690565b156129d057612805816001600160401b03166000526008602052604060002090565b608085015190600181019261281a8454611cbf565b6129f7576128a17ff4c1390c70e5c0f491ae1ccbc06f9117cbbadf2767b247b3bc203280f24c0fb9916128878475010000000000000000000000000000000000000000007fffffff0000000000000000ffffffffffffffffffffffffffffffffffffffffff825416179055565b6040516001600160401b0390911681529081906020820190565b0390a15b815180159081156129e1575b506129d0576129b161297c60606001986128ef6129c7967fbd1ab25a0ff0a36a588597ba1af11e30f3f210de8b9e818cc9bbc457c94c8d8c9861312a565b6129456128ff6040830151151590565b86547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178655565b61297561295982516001600160a01b031690565b86906001600160a01b03166001600160a01b0319825416179055565b0151151590565b82547fffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690151560e81b60ff60e81b16178255565b6129ba84614ecd565b50604051918291826131fb565b0390a201612792565b6342bcdf7f60e11b60005260046000fd5b905060208301206129f06130ad565b14386128b1565b60016001600160401b03612a1684546001600160401b039060a81c1690565b16141580612a46575b612a2957506128a5565b632105803760e11b6000526001600160401b031660045260246000fd5b50612a5084611d7c565b60208151910120835160208501201415612a1f565b60405190612a74602083610388565b6000808352366020840137565b60408051909190612a928382610388565b6001815291601f1901366020840137565b612aab613088565b60005b815181101561245157612ac18183611cab565b51906040820160ff612ad4825160ff1690565b1615612d7657602083015160ff1692612afa8460ff166000526002602052604060002090565b9160018301918254612b15612b0f8260ff1690565b60ff1690565b612d3b5750612b42612b2a6060830151151590565b845462ff0000191690151560101b62ff000016178455565b60a08101918251610100815111612ce357805115612d255760038601612b70612b6a82612773565b8a6143af565b6060840151612c00575b947fab8b1b57514019638d7b5ce9c638fe71366fe8e2be1c40a7a80f1733d0e9f54794600294612bdc612bcc612bfa9a96612bc58760019f9c612bc0612bf29a8f61451d565b613603565b5160ff1690565b845460ff191660ff821617909455565b5190818555519060405195869501908886613689565b0390a161459f565b01612aae565b97946002879395970196612c1c612c1689612773565b886143af565b608085015194610100865111612d0f578551612c44612b0f612c3f8a5160ff1690565b6135ef565b1015612cf9578551845111612ce357612bdc612bcc7fab8b1b57514019638d7b5ce9c638fe71366fe8e2be1c40a7a80f1733d0e9f54798612bc58760019f612bc0612bfa9f9a8f612ccb60029f612cc5612bf29f8f90612bc08492612caa845160ff1690565b908054909161ff001990911660089190911b61ff0016179055565b82614443565b505050979c9f50975050969a50505094509450612b7a565b631b3fab5160e11b600052600160045260246000fd5b631b3fab5160e11b600052600360045260246000fd5b631b3fab5160e11b600052600260045260246000fd5b631b3fab5160e11b600052600560045260246000fd5b60101c60ff16612d56612d516060840151151590565b151590565b90151514612b42576321fd80df60e21b60005260ff861660045260246000fd5b631b3fab5160e11b600090815260045260246000fd5b906001600160401b03612dcc921660005260096020526701ffffffffffffff60406000209160071c166001600160401b0316600052602052604060002090565b5490565b7f0000000000000000000000000000000000000000000000000000000000000000468103612dfb5750565b630f01ce8560e01b6000526004524660245260446000fd5b919091805115612eb5578251159260209160405192612e328185610388565b60008452601f19810160005b818110612e915750505060005b8151811015612e895780612e72612e6460019385611cab565b518815612e785786906137dd565b01612e4b565b612e828387611cab565b51906137dd565b505050509050565b8290604051612e9f81610337565b6000815260608382015282828901015201612e3e565b63c2e5347d60e01b60005260046000fd5b9190811015611c665760051b0190565b3561057b816108bd565b9190811015611c665760051b81013590601e19813603018212156101975701908135916001600160401b038311610197576020018236038113610197579190565b90929491939796815196612f34886107d6565b97612f42604051998a610388565b808952612f51601f19916107d6565b0160005b81811061302357505060005b83518110156130165780612fa88c8a8a8a612fa2612f9b878d612f94828f8f9d8f9e60019f81612fc4575b505050611cab565b5197612ee0565b369161084f565b93613ee6565b612fb2828c611cab565b52612fbd818b611cab565b5001612f61565b63ffffffff612fdc612fd7858585612ec6565b612ed6565b1615612f8c5761300c92612ff392612fd792612ec6565b6040612fff8585611cab565b51019063ffffffff169052565b8f8f908391612f8c565b5096985050505050505050565b60209061302e61209c565b82828d01015201612f55565b61304b6385572ffb60e01b82614249565b9081613065575b8161305b575090565b61057b915061421b565b9050613070816141a0565b1590613052565b61304b63aff2afbf60e01b82614249565b6001600160a01b0360015416330361309c57565b6315ae3a6f60e11b60005260046000fd5b604051602081019060008252602081526130c8604082610388565b51902090565b8181106130d9575050565b600081556001016130ce565b9190601f81116130f457505050565b6103b8926000526020600020906020601f840160051c83019310613120575b601f0160051c01906130ce565b9091508190613113565b91909182516001600160401b038111610332576131518161314b8454611cbf565b846130e5565b6020601f8211600114613192578190613183939495600092613187575b50508160011b916000199060031b1c19161790565b9055565b01519050388061316e565b601f198216906131a784600052602060002090565b9160005b8181106131e3575095836001959697106131ca575b505050811b019055565b015160001960f88460031b161c191690553880806131c0565b9192602060018192868b0151815501940192016131ab565b90600160c061057b936020815260ff84546001600160a01b0381166020840152818160a01c16151560408401526001600160401b038160a81c16606084015260e81c161515608082015260a080820152019101611cf9565b6084019081608411611e2e57565b60a001908160a011611e2e57565b91908201809211611e2e57565b6003111561077c57565b600382101561077c5752565b906103b86040516132a281610337565b602060ff829554818116845260081c169101613286565b8054821015611c665760005260206000200190600090565b60ff60019116019060ff8211611e2e57565b60ff601b9116019060ff8211611e2e57565b90606092604091835260208301370190565b600160005260026020529361333b7fe90b7bceb6e7df5418fb78d8ee546e97c83a08bbccc01a0644d599ccd2a7c2e06126f7565b9385359461334885613253565b60608201906133578251151590565b6135c1575b8036036135a9575081518781036135905750613376612dd0565b3360009081527fa15bc60c955c405d20d9149c709e2460f1c2d9a497496a7f46004d1772c3054c602052604090206133b290613292565b613292565b600260208201516133c28161327c565b6133cb8161327c565b149081613534575b5015613508575b5161343f575b50505050507f198d6990ef96613a9026203077e422916918b03ff47f0be6bee7b02d8e139ef09061342361341660019460200190565b356001600160401b031690565b604080519283526001600160401b0391909116602083015290a2565b613460612b0f61345b602085969799989a955194015160ff1690565b6132d1565b036134f75781518351036134e6576134de600061342394613416946134aa7f198d6990ef96613a9026203077e422916918b03ff47f0be6bee7b02d8e139ef09960019b369161084f565b602081519101206040516134d5816134c7896020830195866132f5565b03601f198101835282610388565b5190208a614279565b9483946133e0565b63a75d88af60e01b60005260046000fd5b6371253a2560e01b60005260046000fd5b72c11c11c11c11c11c11c11c11c11c11c11c11c13303156133da57631b41e11d60e31b60005260046000fd5b60016000526002602052516135889150612382906135759060ff167fe90b7bceb6e7df5418fb78d8ee546e97c83a08bbccc01a0644d599ccd2a7c2e36132b9565b90546001600160a01b039160031b1c1690565b3314386133d3565b6324f7d61360e21b600052600452602487905260446000fd5b638e1192e160e01b6000526004523660245260446000fd5b6135ea906135e46135da6135d58751611e18565b613261565b6135e48851611e18565b9061326f565b61335c565b60ff166003029060ff8216918203611e2e57565b8151916001600160401b0383116103325768010000000000000000831161033257602090825484845580851061366c575b500190600052602060002060005b83811061364f5750505050565b60019060206001600160a01b038551169401938184015501613642565b6136839084600052858460002091820191016130ce565b38613634565b95949392909160ff6136ae93168752602087015260a0604087015260a0860190612730565b84810360608601526020808351928381520192019060005b8181106136e1575050509060806103b89294019060ff169052565b82516001600160a01b03168452602093840193909201916001016136c6565b600654811015611c665760066000527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f015490565b90816020910312610197575161057b8161161d565b6001600160401b0361057b949381606094168352166020820152816040820152019061043b565b60409061057b93928152816020820152019061043b565b9291906001600160401b0390816064951660045216602452600481101561077c57604452565b9493926137c76060936137d89388526020880190610781565b60806040870152608086019061043b565b930152565b906137ef82516001600160401b031690565b8151604051632cbc26bb60e01b815277ffffffffffffffff00000000000000000000000000000000608084901b1660048201529015159391906001600160401b038216906020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561244c57600091613dcf575b50613d70576020830191825151948515613d405760408501948551518703613d2f576138a0908361460e565b95909760005b8881106138b95750505050505050505050565b5a6138ce6138c8838a51611cab565b516148d9565b8051606001516138e7906001600160401b031688611e40565b6138f081610772565b8015908d8283159384613d1c575b15613cd95760608815613c5c5750613925602061391b898d611cab565b5101519242611e33565b60045461393a9060a01c63ffffffff16611f6e565b108015613c49575b15613c2b57613951878b611cab565b5151613c15575b845160800151613970906001600160401b031661069d565b613b4d575b50613981868951611cab565b5160a085015151815103613b1157936139e69695938c938f966139c68e958c926139c06139ba60608951016001600160401b0390511690565b89614923565b86614c69565b9a9080966139e060608851016001600160401b0390511690565b906149ab565b613abf575b50506139f682610772565b60028203613a77575b600196613a6d7f05665fe9ad095383d018353f4cbcba77e84db27dd215081bbf7cdf9ae6fbe48b936001600160401b03935192613a5e613a558b613a4d60608801516001600160401b031690565b96519b611cab565b51985a90611e33565b916040519586951698856137ae565b0390a45b016138a6565b91509193949250613a8782610772565b60038203613a9b578b929493918a916139ff565b51606001516349362d1f60e11b60005261203e91906001600160401b031689613788565b613ac884610772565b600384036139eb579092949550613ae0919350610772565b613af0578b92918a9138806139eb565b5151604051632b11b8d960e01b8152908190611fe890879060048401613771565b61203e8b613b2b60608851016001600160401b0390511690565b631cfe6d8b60e01b6000526001600160401b0391821660045216602452604490565b613b5683610772565b613b61575b38613975565b8351608001516001600160401b0316602080860151918c613b9660405194859384936370701e5760e11b85526004850161374a565b038160006001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165af190811561244c57600091613be7575b50613b5b575050505050600190613a71565b613c08915060203d8111613c0e575b613c008183610388565b810190613735565b38613bd5565b503d613bf6565b613c1f878b611cab565b51516080860152613958565b6354e7e43160e11b6000526001600160401b038b1660045260246000fd5b50613c5383610772565b60038314613942565b915083613c6884610772565b1561395857506001959450613cd19250613caf91507f3ef2a99c550a751d4b0b261268f05a803dfb049ab43616a1ffb388f61fe651209351016001600160401b0390511690565b604080516001600160401b03808c168252909216602083015290918291820190565b0390a1613a71565b505050506001929150613cd1613caf60607f3b575419319662b2a6f5e2467d84521517a3382b908eb3d557bb3fdb0c50e23c9351016001600160401b0390511690565b50613d2683610772565b600383146138fe565b6357e0e08360e01b60005260046000fd5b61203e613d5486516001600160401b031690565b63676cf24b60e11b6000526001600160401b0316600452602490565b5092915050613db2576040516001600160401b039190911681527faab522ed53d887e56ed53dd37398a01aeef6a58e0fa77c2173beb9512d89493390602090a1565b637edeb53960e11b6000526001600160401b031660045260246000fd5b613de8915060203d602011613c0e57613c008183610388565b38613874565b90816020910312610197575161057b816108a1565b9061057b916020815260e0613ea1613e8c613e2c8551610100602087015261012086019061043b565b60208601516001600160401b0316604086015260408601516001600160a01b0316606086015260608601516080860152613e76608087015160a08701906001600160a01b03169052565b60a0860151858203601f190160c087015261043b565b60c0850151848203601f19018486015261043b565b92015190610100601f198285030191015261043b565b6040906001600160a01b0361057b9493168152816020820152019061043b565b90816020910312610197575190565b91939293613ef261209c565b5060208301516001600160a01b031660405163bbe4f6db60e01b81526001600160a01b038216600482015290959092602084806024810103816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa93841561244c5760009461416f575b506001600160a01b038416958615801561415d575b61413f5761402461404d926134c792613fa8613fa1611f6e60408c015163ffffffff1690565b8c89614d82565b9690996080810151613fd66060835193015193613fc36103c9565b9687526001600160401b03166020870152565b6001600160a01b038a16604086015260608501526001600160a01b038d16608085015260a084015260c083015260e0820152604051633907753760e01b602082015292839160248301613e03565b82857f000000000000000000000000000000000000000000000000000000000000000092614e10565b94909115614123575080516020810361410a575090614076826020808a95518301019101613ed7565b956001600160a01b038416036140ae575b50505050506140a66140976103d9565b6001600160a01b039093168352565b602082015290565b6140c1936140bb91611e33565b91614d82565b509080821080156140f7575b6140d957808481614087565b63a966e21f60e01b6000908152600493909352602452604452606490fd5b50826141038284611e33565b14156140cd565b631e3be00960e21b600052602060045260245260446000fd5b611fe8604051928392634ff17cad60e11b845260048401613eb7565b63ae9b4ce960e01b6000526001600160a01b03851660045260246000fd5b5061416a61247586613077565b613f7b565b61419291945060203d602011614199575b61418a8183610388565b810190613dee565b9238613f66565b503d614180565b60405160208101916301ffc9a760e01b835263ffffffff60e01b6024830152602482526141ce604483610388565b6179185a1061420a576020926000925191617530fa6000513d826141fe575b50816141f7575090565b9050151590565b602011159150386141ed565b63753fa58960e11b60005260046000fd5b60405160208101916301ffc9a760e01b83526301ffc9a760e01b6024830152602482526141ce604483610388565b6040519060208201926301ffc9a760e01b845263ffffffff60e01b166024830152602482526141ce604483610388565b919390926000948051946000965b868810614298575050505050505050565b6020881015611c6657602060006142b0878b1a6132e3565b6142ba8b87611cab565b51906142f16142c98d8a611cab565b5160405193849389859094939260ff6060936080840197845216602083015260408201520152565b838052039060015afa1561244c576143376133ad60005161431f8960ff166000526003602052604060002090565b906001600160a01b0316600052602052604060002090565b90600160208301516143488161327c565b6143518161327c565b0361439e5761436e614364835160ff1690565b60ff600191161b90565b811661438d576143846143646001935160ff1690565b17970196614287565b633d9ef1f160e21b60005260046000fd5b636518c33d60e11b60005260046000fd5b91909160005b83518110156144085760019060ff831660005260036020526000614401604082206001600160a01b036143e8858a611cab565b51166001600160a01b0316600052602052604060002090565b55016143b5565b50509050565b8151815460ff191660ff919091161781559060200151600381101561077c57815461ff00191660089190911b61ff0016179055565b919060005b81518110156144085761446b61445e8284611cab565b516001600160a01b031690565b9061449461448a8361431f8860ff166000526003602052604060002090565b5460081c60ff1690565b61449d8161327c565b614508576001600160a01b038216156144f7576144f16001926144ec6144c16103d9565b60ff85168152916144d58660208501613286565b61431f8960ff166000526003602052604060002090565b61440e565b01614448565b63d6c62c9b60e01b60005260046000fd5b631b3fab5160e11b6000526004805260246000fd5b919060005b81518110156144085761453861445e8284611cab565b9061455761448a8361431f8860ff166000526003602052604060002090565b6145608161327c565b614508576001600160a01b038216156144f7576145996001926144ec6145846103d9565b60ff85168152916144d5600260208501613286565b01614522565b60ff1680600052600260205260ff60016040600020015460101c169080156000146145ed5750156145dc576001600160401b0319600b5416600b55565b6317bd8dd160e11b60005260046000fd5b6001146145f75750565b6145fd57565b6307b8c74d60e51b60005260046000fd5b91909160208301805151600181036148c0575061462b9051611c59565b51926146426020855101516001600160401b031690565b6001600160401b0383166001600160401b0382160361489c57508351604001516001600160401b03167f0000000000000000000000000000000000000000000000000000000000000000906001600160401b0382166001600160401b0382160361487957505060606146b691015184614f9e565b916146da6123826146cb60016118be85615078565b60208082518301019101613dee565b906146ec83516001600160a01b031690565b6001600160a01b0383166001600160a01b0382160361484e5750614723816001600160401b0316600052600c602052604060002090565b549081156148315760808401519081830361480e575050506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016803b15610197576040805163ab4d6f7560e01b815284516001600160a01b0316600482015260208501516024820152908401516044820152606084015160648201526080840151608482015260a4810194909452600090849060c490829084905af190811561244c576060936147e9926147f9575b506147e3612a81565b956150b5565b6147f285611c59565b5201519190565b80612511600061480893610388565b386147da565b63f64ea74360e01b6000526001600160401b031660045260245260445260646000fd5b63bd04062360e01b6000526001600160401b031660045260246000fd5b636db7ad9160e01b6000526001600160401b039091166004526001600160a01b031660245260446000fd5b632f0ccf7760e01b6000526001600160401b039081166004521660245260446000fd5b633c3e620760e11b6000526001600160401b03908116600452821660245260446000fd5b6329852a5160e21b600052600452600160245260446000fd5b60405160c081018181106001600160401b038211176103325760609160a091604052614903611bb8565b815282602082015282604082015260008382015260006080820152015290565b607f8216906801fffffffffffffffe6001600160401b0383169260011b169180830460021490151715611e2e576149a8916001600160401b036149668584612d8c565b921660005260096020526701ffffffffffffff60406000209460071c169160036001831b921b19161792906001600160401b0316600052602052604060002090565b55565b9091607f83166801fffffffffffffffe6001600160401b0382169160011b169080820460021490151715611e2e576149e38484612d8c565b600483101561077c576001600160401b036149a89416600052600960205260036701ffffffffffffff60406000209660071c1693831b921b19161792906001600160401b0316600052602052604060002090565b9080602083519182815201916020808360051b8301019401926000915b838310614a6357505050505090565b9091929394602080600192601f19858203018652885190608080614ac6614a93855160a0865260a086019061043b565b6001600160a01b0387870151168786015263ffffffff60408701511660408601526060860151858203606087015261043b565b93015191015297019301930191939290614a54565b61057b916001600160401b036080835180518452826020820151166020850152826040820151166040850152826060820151166060850152015116608082015260a0614b4c614b3a60208501516101408486015261014085019061043b565b604085015184820360c086015261043b565b60608401516001600160a01b031660e0840152926080810151610100840152015190610120818403910152614a37565b90614b8f90606083526060830190614adb565b8181036020830152825180825260208201916020808360051b8301019501926000915b838310614bff57505050505060408183039101526020808351928381520192019060005b818110614be35750505090565b825163ffffffff16845260209384019390920191600101614bd6565b9091929395602080614c1d600193601f198682030187528a5161043b565b98019301930191939290614bb2565b80516020909101516001600160e01b0319811692919060048210614c4e575050565b6001600160e01b031960049290920360031b82901b16169150565b90303b1561019757600091614c926040519485938493630304c3e160e51b855260048501614b7c565b038183305af19081614d6d575b50614d6257614cac6121a7565b9072c11c11c11c11c11c11c11c11c11c11c11c11c13314614cce575b60039190565b614ce7614cda83614c2c565b6001600160e01b03191690565b6337c3be2960e01b148015614d47575b8015614d2c575b15614cc85761203e614d0f83614c2c565b632882569d60e01b6000526001600160e01b031916600452602490565b50614d39614cda83614c2c565b63753fa58960e11b14614cfe565b50614d54614cda83614c2c565b632be8ca8b60e21b14614cf7565b60029061057b610403565b806125116000614d7c93610388565b38614c9f565b6040516370a0823160e01b60208201526001600160a01b039091166024820152919291614ddf90614db681604481016134c7565b84837f000000000000000000000000000000000000000000000000000000000000000092614e10565b92909115614123575080516020810361410a575090614e0a8260208061057b95518301019101613ed7565b93611e33565b939193614e1d60846103e8565b94614e2b6040519687610388565b60848652614e3960846103e8565b602087019590601f1901368737833b15614ebc575a90808210614eab578291038060061c90031115614e9a576000918291825a9560208451940192f1905a9003923d9060848211614e91575b6000908287523e929190565b60849150614e85565b6337c3be2960e01b60005260046000fd5b632be8ca8b60e21b60005260046000fd5b63030ed58f60e21b60005260046000fd5b80600052600760205260406000205415600014614f4b576006546801000000000000000081101561033257600181016006556000600654821015611c6657600690527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f01819055600654906000526007602052604060002055600190565b50600090565b6001600160401b0361057b9493816080947fb056c48de8fce43e6c30818e5f9c6a56a7014ef5944b6329ed5c76afff7e838a85521660208401521660408201528160608201520190614adb565b919091614fa9611bb8565b5082516005810361505f5750614fc461238261445e85611c59565b92614fce81611c6b565b5190614fd981611c7b565b51614fed614fe683611c8b565b5192611c9b565b5192615009614ffa6103ba565b6001600160a01b039098168852565b60208701526040860152606085015260808401528051906130c861504b606061503c60408601516001600160401b031690565b9401516001600160401b031690565b6134c7604051938492602084019687614f51565b630608814160e41b600052600452600560245260446000fd5b6001600160401b031680600052600860205260406000209060ff825460a01c16156150a1575090565b63ed053c5960e01b60005260045260246000fd5b61057b918151906001600160401b0360408160208501511693015116906040516001600160a01b036020820192168252602081526150f4604082610388565b5190206040519160208301937f2425b0b9f9054c76ff151b0a175b18f37a4a4e82013a72e9f15c9caa095ed21f85526040840152606083015260808201526080815261514160a082610388565b5190209061515b565b90602061057b928181520190614a37565b6130c881518051906151ef61517a60608601516001600160a01b031690565b6134c761519160608501516001600160401b031690565b936151aa6080808a01519201516001600160401b031690565b90604051958694602086019889936001600160401b036080946001600160a01b0382959998949960a089019a8952166020880152166040860152606085015216910152565b5190206134c76020840151602081519101209360a0604082015160208151910120910151604051615228816134c760208201948561514a565b51902090604051958694602086019889919260a093969594919660c0840197600085526020850152604084015260608301526080820152015256fea164736f6c634300081a000abd1ab25a0ff0a36a588597ba1af11e30f3f210de8b9e818cc9bbc457c94c8d8c",
}

var OffRampOverSuperchainInteropABI = OffRampOverSuperchainInteropMetaData.ABI

var OffRampOverSuperchainInteropBin = OffRampOverSuperchainInteropMetaData.Bin

func DeployOffRampOverSuperchainInterop(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig OffRampStaticConfig, dynamicConfig OffRampDynamicConfig, sourceChainConfigs []OffRampSourceChainConfigArgs, crossL2Inbox common.Address, chainSelectorToChainIdConfigArgs []OffRampOverSuperchainInteropChainSelectorToChainIdConfigArgs) (common.Address, *types.Transaction, *OffRampOverSuperchainInterop, error) {
	parsed, err := OffRampOverSuperchainInteropMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OffRampOverSuperchainInteropBin), backend, staticConfig, dynamicConfig, sourceChainConfigs, crossL2Inbox, chainSelectorToChainIdConfigArgs)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OffRampOverSuperchainInterop{address: address, abi: *parsed, OffRampOverSuperchainInteropCaller: OffRampOverSuperchainInteropCaller{contract: contract}, OffRampOverSuperchainInteropTransactor: OffRampOverSuperchainInteropTransactor{contract: contract}, OffRampOverSuperchainInteropFilterer: OffRampOverSuperchainInteropFilterer{contract: contract}}, nil
}

type OffRampOverSuperchainInterop struct {
	address common.Address
	abi     abi.ABI
	OffRampOverSuperchainInteropCaller
	OffRampOverSuperchainInteropTransactor
	OffRampOverSuperchainInteropFilterer
}

type OffRampOverSuperchainInteropCaller struct {
	contract *bind.BoundContract
}

type OffRampOverSuperchainInteropTransactor struct {
	contract *bind.BoundContract
}

type OffRampOverSuperchainInteropFilterer struct {
	contract *bind.BoundContract
}

type OffRampOverSuperchainInteropSession struct {
	Contract     *OffRampOverSuperchainInterop
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type OffRampOverSuperchainInteropCallerSession struct {
	Contract *OffRampOverSuperchainInteropCaller
	CallOpts bind.CallOpts
}

type OffRampOverSuperchainInteropTransactorSession struct {
	Contract     *OffRampOverSuperchainInteropTransactor
	TransactOpts bind.TransactOpts
}

type OffRampOverSuperchainInteropRaw struct {
	Contract *OffRampOverSuperchainInterop
}

type OffRampOverSuperchainInteropCallerRaw struct {
	Contract *OffRampOverSuperchainInteropCaller
}

type OffRampOverSuperchainInteropTransactorRaw struct {
	Contract *OffRampOverSuperchainInteropTransactor
}

func NewOffRampOverSuperchainInterop(address common.Address, backend bind.ContractBackend) (*OffRampOverSuperchainInterop, error) {
	abi, err := abi.JSON(strings.NewReader(OffRampOverSuperchainInteropABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindOffRampOverSuperchainInterop(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OffRampOverSuperchainInterop{address: address, abi: abi, OffRampOverSuperchainInteropCaller: OffRampOverSuperchainInteropCaller{contract: contract}, OffRampOverSuperchainInteropTransactor: OffRampOverSuperchainInteropTransactor{contract: contract}, OffRampOverSuperchainInteropFilterer: OffRampOverSuperchainInteropFilterer{contract: contract}}, nil
}

func NewOffRampOverSuperchainInteropCaller(address common.Address, caller bind.ContractCaller) (*OffRampOverSuperchainInteropCaller, error) {
	contract, err := bindOffRampOverSuperchainInterop(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OffRampOverSuperchainInteropCaller{contract: contract}, nil
}

func NewOffRampOverSuperchainInteropTransactor(address common.Address, transactor bind.ContractTransactor) (*OffRampOverSuperchainInteropTransactor, error) {
	contract, err := bindOffRampOverSuperchainInterop(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OffRampOverSuperchainInteropTransactor{contract: contract}, nil
}

func NewOffRampOverSuperchainInteropFilterer(address common.Address, filterer bind.ContractFilterer) (*OffRampOverSuperchainInteropFilterer, error) {
	contract, err := bindOffRampOverSuperchainInterop(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OffRampOverSuperchainInteropFilterer{contract: contract}, nil
}

func bindOffRampOverSuperchainInterop(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OffRampOverSuperchainInteropMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OffRampOverSuperchainInterop.Contract.OffRampOverSuperchainInteropCaller.contract.Call(opts, result, method, params...)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.Contract.OffRampOverSuperchainInteropTransactor.contract.Transfer(opts)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.Contract.OffRampOverSuperchainInteropTransactor.contract.Transact(opts, method, params...)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OffRampOverSuperchainInterop.Contract.contract.Call(opts, result, method, params...)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.Contract.contract.Transfer(opts)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.Contract.contract.Transact(opts, method, params...)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropCaller) CcipReceive(opts *bind.CallOpts, arg0 ClientAny2EVMMessage) error {
	var out []interface{}
	err := _OffRampOverSuperchainInterop.contract.Call(opts, &out, "ccipReceive", arg0)

	if err != nil {
		return err
	}

	return err

}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropSession) CcipReceive(arg0 ClientAny2EVMMessage) error {
	return _OffRampOverSuperchainInterop.Contract.CcipReceive(&_OffRampOverSuperchainInterop.CallOpts, arg0)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropCallerSession) CcipReceive(arg0 ClientAny2EVMMessage) error {
	return _OffRampOverSuperchainInterop.Contract.CcipReceive(&_OffRampOverSuperchainInterop.CallOpts, arg0)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropCaller) Commit(opts *bind.CallOpts, arg0 [2][32]byte, arg1 []byte, arg2 [][32]byte, arg3 [][32]byte, arg4 [32]byte) error {
	var out []interface{}
	err := _OffRampOverSuperchainInterop.contract.Call(opts, &out, "commit", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return err
	}

	return err

}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropSession) Commit(arg0 [2][32]byte, arg1 []byte, arg2 [][32]byte, arg3 [][32]byte, arg4 [32]byte) error {
	return _OffRampOverSuperchainInterop.Contract.Commit(&_OffRampOverSuperchainInterop.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropCallerSession) Commit(arg0 [2][32]byte, arg1 []byte, arg2 [][32]byte, arg3 [][32]byte, arg4 [32]byte) error {
	return _OffRampOverSuperchainInterop.Contract.Commit(&_OffRampOverSuperchainInterop.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropCaller) GetAllSourceChainConfigs(opts *bind.CallOpts) ([]uint64, []OffRampSourceChainConfig, error) {
	var out []interface{}
	err := _OffRampOverSuperchainInterop.contract.Call(opts, &out, "getAllSourceChainConfigs")

	if err != nil {
		return *new([]uint64), *new([]OffRampSourceChainConfig), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)
	out1 := *abi.ConvertType(out[1], new([]OffRampSourceChainConfig)).(*[]OffRampSourceChainConfig)

	return out0, out1, err

}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropSession) GetAllSourceChainConfigs() ([]uint64, []OffRampSourceChainConfig, error) {
	return _OffRampOverSuperchainInterop.Contract.GetAllSourceChainConfigs(&_OffRampOverSuperchainInterop.CallOpts)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropCallerSession) GetAllSourceChainConfigs() ([]uint64, []OffRampSourceChainConfig, error) {
	return _OffRampOverSuperchainInterop.Contract.GetAllSourceChainConfigs(&_OffRampOverSuperchainInterop.CallOpts)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropCaller) GetChainId(opts *bind.CallOpts, sourceChainSelector uint64) (*big.Int, error) {
	var out []interface{}
	err := _OffRampOverSuperchainInterop.contract.Call(opts, &out, "getChainId", sourceChainSelector)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropSession) GetChainId(sourceChainSelector uint64) (*big.Int, error) {
	return _OffRampOverSuperchainInterop.Contract.GetChainId(&_OffRampOverSuperchainInterop.CallOpts, sourceChainSelector)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropCallerSession) GetChainId(sourceChainSelector uint64) (*big.Int, error) {
	return _OffRampOverSuperchainInterop.Contract.GetChainId(&_OffRampOverSuperchainInterop.CallOpts, sourceChainSelector)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropCaller) GetCrossL2Inbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OffRampOverSuperchainInterop.contract.Call(opts, &out, "getCrossL2Inbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropSession) GetCrossL2Inbox() (common.Address, error) {
	return _OffRampOverSuperchainInterop.Contract.GetCrossL2Inbox(&_OffRampOverSuperchainInterop.CallOpts)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropCallerSession) GetCrossL2Inbox() (common.Address, error) {
	return _OffRampOverSuperchainInterop.Contract.GetCrossL2Inbox(&_OffRampOverSuperchainInterop.CallOpts)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropCaller) GetDynamicConfig(opts *bind.CallOpts) (OffRampDynamicConfig, error) {
	var out []interface{}
	err := _OffRampOverSuperchainInterop.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(OffRampDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(OffRampDynamicConfig)).(*OffRampDynamicConfig)

	return out0, err

}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropSession) GetDynamicConfig() (OffRampDynamicConfig, error) {
	return _OffRampOverSuperchainInterop.Contract.GetDynamicConfig(&_OffRampOverSuperchainInterop.CallOpts)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropCallerSession) GetDynamicConfig() (OffRampDynamicConfig, error) {
	return _OffRampOverSuperchainInterop.Contract.GetDynamicConfig(&_OffRampOverSuperchainInterop.CallOpts)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropCaller) GetExecutionState(opts *bind.CallOpts, sourceChainSelector uint64, sequenceNumber uint64) (uint8, error) {
	var out []interface{}
	err := _OffRampOverSuperchainInterop.contract.Call(opts, &out, "getExecutionState", sourceChainSelector, sequenceNumber)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropSession) GetExecutionState(sourceChainSelector uint64, sequenceNumber uint64) (uint8, error) {
	return _OffRampOverSuperchainInterop.Contract.GetExecutionState(&_OffRampOverSuperchainInterop.CallOpts, sourceChainSelector, sequenceNumber)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropCallerSession) GetExecutionState(sourceChainSelector uint64, sequenceNumber uint64) (uint8, error) {
	return _OffRampOverSuperchainInterop.Contract.GetExecutionState(&_OffRampOverSuperchainInterop.CallOpts, sourceChainSelector, sequenceNumber)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropCaller) GetLatestPriceSequenceNumber(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OffRampOverSuperchainInterop.contract.Call(opts, &out, "getLatestPriceSequenceNumber")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropSession) GetLatestPriceSequenceNumber() (uint64, error) {
	return _OffRampOverSuperchainInterop.Contract.GetLatestPriceSequenceNumber(&_OffRampOverSuperchainInterop.CallOpts)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropCallerSession) GetLatestPriceSequenceNumber() (uint64, error) {
	return _OffRampOverSuperchainInterop.Contract.GetLatestPriceSequenceNumber(&_OffRampOverSuperchainInterop.CallOpts)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropCaller) GetMerkleRoot(opts *bind.CallOpts, sourceChainSelector uint64, root [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _OffRampOverSuperchainInterop.contract.Call(opts, &out, "getMerkleRoot", sourceChainSelector, root)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropSession) GetMerkleRoot(sourceChainSelector uint64, root [32]byte) (*big.Int, error) {
	return _OffRampOverSuperchainInterop.Contract.GetMerkleRoot(&_OffRampOverSuperchainInterop.CallOpts, sourceChainSelector, root)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropCallerSession) GetMerkleRoot(sourceChainSelector uint64, root [32]byte) (*big.Int, error) {
	return _OffRampOverSuperchainInterop.Contract.GetMerkleRoot(&_OffRampOverSuperchainInterop.CallOpts, sourceChainSelector, root)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropCaller) GetSourceChainConfig(opts *bind.CallOpts, sourceChainSelector uint64) (OffRampSourceChainConfig, error) {
	var out []interface{}
	err := _OffRampOverSuperchainInterop.contract.Call(opts, &out, "getSourceChainConfig", sourceChainSelector)

	if err != nil {
		return *new(OffRampSourceChainConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(OffRampSourceChainConfig)).(*OffRampSourceChainConfig)

	return out0, err

}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropSession) GetSourceChainConfig(sourceChainSelector uint64) (OffRampSourceChainConfig, error) {
	return _OffRampOverSuperchainInterop.Contract.GetSourceChainConfig(&_OffRampOverSuperchainInterop.CallOpts, sourceChainSelector)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropCallerSession) GetSourceChainConfig(sourceChainSelector uint64) (OffRampSourceChainConfig, error) {
	return _OffRampOverSuperchainInterop.Contract.GetSourceChainConfig(&_OffRampOverSuperchainInterop.CallOpts, sourceChainSelector)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropCaller) GetStaticConfig(opts *bind.CallOpts) (OffRampStaticConfig, error) {
	var out []interface{}
	err := _OffRampOverSuperchainInterop.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(OffRampStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(OffRampStaticConfig)).(*OffRampStaticConfig)

	return out0, err

}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropSession) GetStaticConfig() (OffRampStaticConfig, error) {
	return _OffRampOverSuperchainInterop.Contract.GetStaticConfig(&_OffRampOverSuperchainInterop.CallOpts)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropCallerSession) GetStaticConfig() (OffRampStaticConfig, error) {
	return _OffRampOverSuperchainInterop.Contract.GetStaticConfig(&_OffRampOverSuperchainInterop.CallOpts)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropCaller) LatestConfigDetails(opts *bind.CallOpts, ocrPluginType uint8) (MultiOCR3BaseOCRConfig, error) {
	var out []interface{}
	err := _OffRampOverSuperchainInterop.contract.Call(opts, &out, "latestConfigDetails", ocrPluginType)

	if err != nil {
		return *new(MultiOCR3BaseOCRConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(MultiOCR3BaseOCRConfig)).(*MultiOCR3BaseOCRConfig)

	return out0, err

}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropSession) LatestConfigDetails(ocrPluginType uint8) (MultiOCR3BaseOCRConfig, error) {
	return _OffRampOverSuperchainInterop.Contract.LatestConfigDetails(&_OffRampOverSuperchainInterop.CallOpts, ocrPluginType)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropCallerSession) LatestConfigDetails(ocrPluginType uint8) (MultiOCR3BaseOCRConfig, error) {
	return _OffRampOverSuperchainInterop.Contract.LatestConfigDetails(&_OffRampOverSuperchainInterop.CallOpts, ocrPluginType)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OffRampOverSuperchainInterop.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropSession) Owner() (common.Address, error) {
	return _OffRampOverSuperchainInterop.Contract.Owner(&_OffRampOverSuperchainInterop.CallOpts)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropCallerSession) Owner() (common.Address, error) {
	return _OffRampOverSuperchainInterop.Contract.Owner(&_OffRampOverSuperchainInterop.CallOpts)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OffRampOverSuperchainInterop.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropSession) TypeAndVersion() (string, error) {
	return _OffRampOverSuperchainInterop.Contract.TypeAndVersion(&_OffRampOverSuperchainInterop.CallOpts)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropCallerSession) TypeAndVersion() (string, error) {
	return _OffRampOverSuperchainInterop.Contract.TypeAndVersion(&_OffRampOverSuperchainInterop.CallOpts)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.contract.Transact(opts, "acceptOwnership")
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropSession) AcceptOwnership() (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.Contract.AcceptOwnership(&_OffRampOverSuperchainInterop.TransactOpts)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.Contract.AcceptOwnership(&_OffRampOverSuperchainInterop.TransactOpts)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropTransactor) ApplyChainSelectorToChainIdConfigUpdates(opts *bind.TransactOpts, chainSelectorsToUnset []uint64, chainSelectorsToSet []OffRampOverSuperchainInteropChainSelectorToChainIdConfigArgs) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.contract.Transact(opts, "applyChainSelectorToChainIdConfigUpdates", chainSelectorsToUnset, chainSelectorsToSet)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropSession) ApplyChainSelectorToChainIdConfigUpdates(chainSelectorsToUnset []uint64, chainSelectorsToSet []OffRampOverSuperchainInteropChainSelectorToChainIdConfigArgs) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.Contract.ApplyChainSelectorToChainIdConfigUpdates(&_OffRampOverSuperchainInterop.TransactOpts, chainSelectorsToUnset, chainSelectorsToSet)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropTransactorSession) ApplyChainSelectorToChainIdConfigUpdates(chainSelectorsToUnset []uint64, chainSelectorsToSet []OffRampOverSuperchainInteropChainSelectorToChainIdConfigArgs) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.Contract.ApplyChainSelectorToChainIdConfigUpdates(&_OffRampOverSuperchainInterop.TransactOpts, chainSelectorsToUnset, chainSelectorsToSet)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropTransactor) ApplySourceChainConfigUpdates(opts *bind.TransactOpts, sourceChainConfigUpdates []OffRampSourceChainConfigArgs) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.contract.Transact(opts, "applySourceChainConfigUpdates", sourceChainConfigUpdates)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropSession) ApplySourceChainConfigUpdates(sourceChainConfigUpdates []OffRampSourceChainConfigArgs) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.Contract.ApplySourceChainConfigUpdates(&_OffRampOverSuperchainInterop.TransactOpts, sourceChainConfigUpdates)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropTransactorSession) ApplySourceChainConfigUpdates(sourceChainConfigUpdates []OffRampSourceChainConfigArgs) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.Contract.ApplySourceChainConfigUpdates(&_OffRampOverSuperchainInterop.TransactOpts, sourceChainConfigUpdates)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropTransactor) Execute(opts *bind.TransactOpts, reportContext [2][32]byte, report []byte) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.contract.Transact(opts, "execute", reportContext, report)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropSession) Execute(reportContext [2][32]byte, report []byte) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.Contract.Execute(&_OffRampOverSuperchainInterop.TransactOpts, reportContext, report)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropTransactorSession) Execute(reportContext [2][32]byte, report []byte) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.Contract.Execute(&_OffRampOverSuperchainInterop.TransactOpts, reportContext, report)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message InternalAny2EVMRampMessage, offchainTokenData [][]byte, tokenGasOverrides []uint32) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.contract.Transact(opts, "executeSingleMessage", message, offchainTokenData, tokenGasOverrides)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropSession) ExecuteSingleMessage(message InternalAny2EVMRampMessage, offchainTokenData [][]byte, tokenGasOverrides []uint32) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.Contract.ExecuteSingleMessage(&_OffRampOverSuperchainInterop.TransactOpts, message, offchainTokenData, tokenGasOverrides)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropTransactorSession) ExecuteSingleMessage(message InternalAny2EVMRampMessage, offchainTokenData [][]byte, tokenGasOverrides []uint32) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.Contract.ExecuteSingleMessage(&_OffRampOverSuperchainInterop.TransactOpts, message, offchainTokenData, tokenGasOverrides)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropTransactor) ManuallyExecute(opts *bind.TransactOpts, reports []InternalExecutionReport, gasLimitOverrides [][]OffRampGasLimitOverride) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.contract.Transact(opts, "manuallyExecute", reports, gasLimitOverrides)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropSession) ManuallyExecute(reports []InternalExecutionReport, gasLimitOverrides [][]OffRampGasLimitOverride) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.Contract.ManuallyExecute(&_OffRampOverSuperchainInterop.TransactOpts, reports, gasLimitOverrides)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropTransactorSession) ManuallyExecute(reports []InternalExecutionReport, gasLimitOverrides [][]OffRampGasLimitOverride) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.Contract.ManuallyExecute(&_OffRampOverSuperchainInterop.TransactOpts, reports, gasLimitOverrides)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig OffRampDynamicConfig) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropSession) SetDynamicConfig(dynamicConfig OffRampDynamicConfig) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.Contract.SetDynamicConfig(&_OffRampOverSuperchainInterop.TransactOpts, dynamicConfig)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropTransactorSession) SetDynamicConfig(dynamicConfig OffRampDynamicConfig) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.Contract.SetDynamicConfig(&_OffRampOverSuperchainInterop.TransactOpts, dynamicConfig)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropTransactor) SetOCR3Configs(opts *bind.TransactOpts, ocrConfigArgs []MultiOCR3BaseOCRConfigArgs) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.contract.Transact(opts, "setOCR3Configs", ocrConfigArgs)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropSession) SetOCR3Configs(ocrConfigArgs []MultiOCR3BaseOCRConfigArgs) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.Contract.SetOCR3Configs(&_OffRampOverSuperchainInterop.TransactOpts, ocrConfigArgs)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropTransactorSession) SetOCR3Configs(ocrConfigArgs []MultiOCR3BaseOCRConfigArgs) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.Contract.SetOCR3Configs(&_OffRampOverSuperchainInterop.TransactOpts, ocrConfigArgs)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.contract.Transact(opts, "transferOwnership", to)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.Contract.TransferOwnership(&_OffRampOverSuperchainInterop.TransactOpts, to)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.Contract.TransferOwnership(&_OffRampOverSuperchainInterop.TransactOpts, to)
}

type OffRampOverSuperchainInteropAlreadyAttemptedIterator struct {
	Event *OffRampOverSuperchainInteropAlreadyAttempted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampOverSuperchainInteropAlreadyAttemptedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampOverSuperchainInteropAlreadyAttempted)
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
		it.Event = new(OffRampOverSuperchainInteropAlreadyAttempted)
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

func (it *OffRampOverSuperchainInteropAlreadyAttemptedIterator) Error() error {
	return it.fail
}

func (it *OffRampOverSuperchainInteropAlreadyAttemptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampOverSuperchainInteropAlreadyAttempted struct {
	SourceChainSelector uint64
	SequenceNumber      uint64
	Raw                 types.Log
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) FilterAlreadyAttempted(opts *bind.FilterOpts) (*OffRampOverSuperchainInteropAlreadyAttemptedIterator, error) {

	logs, sub, err := _OffRampOverSuperchainInterop.contract.FilterLogs(opts, "AlreadyAttempted")
	if err != nil {
		return nil, err
	}
	return &OffRampOverSuperchainInteropAlreadyAttemptedIterator{contract: _OffRampOverSuperchainInterop.contract, event: "AlreadyAttempted", logs: logs, sub: sub}, nil
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) WatchAlreadyAttempted(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropAlreadyAttempted) (event.Subscription, error) {

	logs, sub, err := _OffRampOverSuperchainInterop.contract.WatchLogs(opts, "AlreadyAttempted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampOverSuperchainInteropAlreadyAttempted)
				if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "AlreadyAttempted", log); err != nil {
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

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) ParseAlreadyAttempted(log types.Log) (*OffRampOverSuperchainInteropAlreadyAttempted, error) {
	event := new(OffRampOverSuperchainInteropAlreadyAttempted)
	if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "AlreadyAttempted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampOverSuperchainInteropChainSelectorToChainIdConfigRemovedIterator struct {
	Event *OffRampOverSuperchainInteropChainSelectorToChainIdConfigRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampOverSuperchainInteropChainSelectorToChainIdConfigRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampOverSuperchainInteropChainSelectorToChainIdConfigRemoved)
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
		it.Event = new(OffRampOverSuperchainInteropChainSelectorToChainIdConfigRemoved)
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

func (it *OffRampOverSuperchainInteropChainSelectorToChainIdConfigRemovedIterator) Error() error {
	return it.fail
}

func (it *OffRampOverSuperchainInteropChainSelectorToChainIdConfigRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampOverSuperchainInteropChainSelectorToChainIdConfigRemoved struct {
	ChainSelector uint64
	ChainId       *big.Int
	Raw           types.Log
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) FilterChainSelectorToChainIdConfigRemoved(opts *bind.FilterOpts, chainSelector []uint64, chainId []*big.Int) (*OffRampOverSuperchainInteropChainSelectorToChainIdConfigRemovedIterator, error) {

	var chainSelectorRule []interface{}
	for _, chainSelectorItem := range chainSelector {
		chainSelectorRule = append(chainSelectorRule, chainSelectorItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _OffRampOverSuperchainInterop.contract.FilterLogs(opts, "ChainSelectorToChainIdConfigRemoved", chainSelectorRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &OffRampOverSuperchainInteropChainSelectorToChainIdConfigRemovedIterator{contract: _OffRampOverSuperchainInterop.contract, event: "ChainSelectorToChainIdConfigRemoved", logs: logs, sub: sub}, nil
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) WatchChainSelectorToChainIdConfigRemoved(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropChainSelectorToChainIdConfigRemoved, chainSelector []uint64, chainId []*big.Int) (event.Subscription, error) {

	var chainSelectorRule []interface{}
	for _, chainSelectorItem := range chainSelector {
		chainSelectorRule = append(chainSelectorRule, chainSelectorItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _OffRampOverSuperchainInterop.contract.WatchLogs(opts, "ChainSelectorToChainIdConfigRemoved", chainSelectorRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampOverSuperchainInteropChainSelectorToChainIdConfigRemoved)
				if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "ChainSelectorToChainIdConfigRemoved", log); err != nil {
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

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) ParseChainSelectorToChainIdConfigRemoved(log types.Log) (*OffRampOverSuperchainInteropChainSelectorToChainIdConfigRemoved, error) {
	event := new(OffRampOverSuperchainInteropChainSelectorToChainIdConfigRemoved)
	if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "ChainSelectorToChainIdConfigRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampOverSuperchainInteropChainSelectorToChainIdConfigUpdatedIterator struct {
	Event *OffRampOverSuperchainInteropChainSelectorToChainIdConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampOverSuperchainInteropChainSelectorToChainIdConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampOverSuperchainInteropChainSelectorToChainIdConfigUpdated)
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
		it.Event = new(OffRampOverSuperchainInteropChainSelectorToChainIdConfigUpdated)
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

func (it *OffRampOverSuperchainInteropChainSelectorToChainIdConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *OffRampOverSuperchainInteropChainSelectorToChainIdConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampOverSuperchainInteropChainSelectorToChainIdConfigUpdated struct {
	ChainSelector uint64
	ChainId       *big.Int
	Raw           types.Log
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) FilterChainSelectorToChainIdConfigUpdated(opts *bind.FilterOpts, chainSelector []uint64, chainId []*big.Int) (*OffRampOverSuperchainInteropChainSelectorToChainIdConfigUpdatedIterator, error) {

	var chainSelectorRule []interface{}
	for _, chainSelectorItem := range chainSelector {
		chainSelectorRule = append(chainSelectorRule, chainSelectorItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _OffRampOverSuperchainInterop.contract.FilterLogs(opts, "ChainSelectorToChainIdConfigUpdated", chainSelectorRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &OffRampOverSuperchainInteropChainSelectorToChainIdConfigUpdatedIterator{contract: _OffRampOverSuperchainInterop.contract, event: "ChainSelectorToChainIdConfigUpdated", logs: logs, sub: sub}, nil
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) WatchChainSelectorToChainIdConfigUpdated(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropChainSelectorToChainIdConfigUpdated, chainSelector []uint64, chainId []*big.Int) (event.Subscription, error) {

	var chainSelectorRule []interface{}
	for _, chainSelectorItem := range chainSelector {
		chainSelectorRule = append(chainSelectorRule, chainSelectorItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _OffRampOverSuperchainInterop.contract.WatchLogs(opts, "ChainSelectorToChainIdConfigUpdated", chainSelectorRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampOverSuperchainInteropChainSelectorToChainIdConfigUpdated)
				if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "ChainSelectorToChainIdConfigUpdated", log); err != nil {
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

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) ParseChainSelectorToChainIdConfigUpdated(log types.Log) (*OffRampOverSuperchainInteropChainSelectorToChainIdConfigUpdated, error) {
	event := new(OffRampOverSuperchainInteropChainSelectorToChainIdConfigUpdated)
	if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "ChainSelectorToChainIdConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampOverSuperchainInteropCommitReportAcceptedIterator struct {
	Event *OffRampOverSuperchainInteropCommitReportAccepted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampOverSuperchainInteropCommitReportAcceptedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampOverSuperchainInteropCommitReportAccepted)
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
		it.Event = new(OffRampOverSuperchainInteropCommitReportAccepted)
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

func (it *OffRampOverSuperchainInteropCommitReportAcceptedIterator) Error() error {
	return it.fail
}

func (it *OffRampOverSuperchainInteropCommitReportAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampOverSuperchainInteropCommitReportAccepted struct {
	BlessedMerkleRoots   []InternalMerkleRoot
	UnblessedMerkleRoots []InternalMerkleRoot
	PriceUpdates         InternalPriceUpdates
	Raw                  types.Log
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) FilterCommitReportAccepted(opts *bind.FilterOpts) (*OffRampOverSuperchainInteropCommitReportAcceptedIterator, error) {

	logs, sub, err := _OffRampOverSuperchainInterop.contract.FilterLogs(opts, "CommitReportAccepted")
	if err != nil {
		return nil, err
	}
	return &OffRampOverSuperchainInteropCommitReportAcceptedIterator{contract: _OffRampOverSuperchainInterop.contract, event: "CommitReportAccepted", logs: logs, sub: sub}, nil
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) WatchCommitReportAccepted(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropCommitReportAccepted) (event.Subscription, error) {

	logs, sub, err := _OffRampOverSuperchainInterop.contract.WatchLogs(opts, "CommitReportAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampOverSuperchainInteropCommitReportAccepted)
				if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "CommitReportAccepted", log); err != nil {
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

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) ParseCommitReportAccepted(log types.Log) (*OffRampOverSuperchainInteropCommitReportAccepted, error) {
	event := new(OffRampOverSuperchainInteropCommitReportAccepted)
	if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "CommitReportAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampOverSuperchainInteropConfigSetIterator struct {
	Event *OffRampOverSuperchainInteropConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampOverSuperchainInteropConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampOverSuperchainInteropConfigSet)
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
		it.Event = new(OffRampOverSuperchainInteropConfigSet)
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

func (it *OffRampOverSuperchainInteropConfigSetIterator) Error() error {
	return it.fail
}

func (it *OffRampOverSuperchainInteropConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampOverSuperchainInteropConfigSet struct {
	OcrPluginType uint8
	ConfigDigest  [32]byte
	Signers       []common.Address
	Transmitters  []common.Address
	F             uint8
	Raw           types.Log
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) FilterConfigSet(opts *bind.FilterOpts) (*OffRampOverSuperchainInteropConfigSetIterator, error) {

	logs, sub, err := _OffRampOverSuperchainInterop.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &OffRampOverSuperchainInteropConfigSetIterator{contract: _OffRampOverSuperchainInterop.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropConfigSet) (event.Subscription, error) {

	logs, sub, err := _OffRampOverSuperchainInterop.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampOverSuperchainInteropConfigSet)
				if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) ParseConfigSet(log types.Log) (*OffRampOverSuperchainInteropConfigSet, error) {
	event := new(OffRampOverSuperchainInteropConfigSet)
	if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampOverSuperchainInteropDynamicConfigSetIterator struct {
	Event *OffRampOverSuperchainInteropDynamicConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampOverSuperchainInteropDynamicConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampOverSuperchainInteropDynamicConfigSet)
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
		it.Event = new(OffRampOverSuperchainInteropDynamicConfigSet)
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

func (it *OffRampOverSuperchainInteropDynamicConfigSetIterator) Error() error {
	return it.fail
}

func (it *OffRampOverSuperchainInteropDynamicConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampOverSuperchainInteropDynamicConfigSet struct {
	DynamicConfig OffRampDynamicConfig
	Raw           types.Log
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) FilterDynamicConfigSet(opts *bind.FilterOpts) (*OffRampOverSuperchainInteropDynamicConfigSetIterator, error) {

	logs, sub, err := _OffRampOverSuperchainInterop.contract.FilterLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return &OffRampOverSuperchainInteropDynamicConfigSetIterator{contract: _OffRampOverSuperchainInterop.contract, event: "DynamicConfigSet", logs: logs, sub: sub}, nil
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropDynamicConfigSet) (event.Subscription, error) {

	logs, sub, err := _OffRampOverSuperchainInterop.contract.WatchLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampOverSuperchainInteropDynamicConfigSet)
				if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
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

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) ParseDynamicConfigSet(log types.Log) (*OffRampOverSuperchainInteropDynamicConfigSet, error) {
	event := new(OffRampOverSuperchainInteropDynamicConfigSet)
	if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampOverSuperchainInteropExecutionStateChangedIterator struct {
	Event *OffRampOverSuperchainInteropExecutionStateChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampOverSuperchainInteropExecutionStateChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampOverSuperchainInteropExecutionStateChanged)
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
		it.Event = new(OffRampOverSuperchainInteropExecutionStateChanged)
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

func (it *OffRampOverSuperchainInteropExecutionStateChangedIterator) Error() error {
	return it.fail
}

func (it *OffRampOverSuperchainInteropExecutionStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampOverSuperchainInteropExecutionStateChanged struct {
	SourceChainSelector uint64
	SequenceNumber      uint64
	MessageId           [32]byte
	MessageHash         [32]byte
	State               uint8
	ReturnData          []byte
	GasUsed             *big.Int
	Raw                 types.Log
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) FilterExecutionStateChanged(opts *bind.FilterOpts, sourceChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (*OffRampOverSuperchainInteropExecutionStateChangedIterator, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}
	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _OffRampOverSuperchainInterop.contract.FilterLogs(opts, "ExecutionStateChanged", sourceChainSelectorRule, sequenceNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return &OffRampOverSuperchainInteropExecutionStateChangedIterator{contract: _OffRampOverSuperchainInterop.contract, event: "ExecutionStateChanged", logs: logs, sub: sub}, nil
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropExecutionStateChanged, sourceChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}
	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _OffRampOverSuperchainInterop.contract.WatchLogs(opts, "ExecutionStateChanged", sourceChainSelectorRule, sequenceNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampOverSuperchainInteropExecutionStateChanged)
				if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
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

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) ParseExecutionStateChanged(log types.Log) (*OffRampOverSuperchainInteropExecutionStateChanged, error) {
	event := new(OffRampOverSuperchainInteropExecutionStateChanged)
	if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampOverSuperchainInteropOwnershipTransferRequestedIterator struct {
	Event *OffRampOverSuperchainInteropOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampOverSuperchainInteropOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampOverSuperchainInteropOwnershipTransferRequested)
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
		it.Event = new(OffRampOverSuperchainInteropOwnershipTransferRequested)
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

func (it *OffRampOverSuperchainInteropOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *OffRampOverSuperchainInteropOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampOverSuperchainInteropOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OffRampOverSuperchainInteropOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffRampOverSuperchainInterop.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OffRampOverSuperchainInteropOwnershipTransferRequestedIterator{contract: _OffRampOverSuperchainInterop.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffRampOverSuperchainInterop.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampOverSuperchainInteropOwnershipTransferRequested)
				if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) ParseOwnershipTransferRequested(log types.Log) (*OffRampOverSuperchainInteropOwnershipTransferRequested, error) {
	event := new(OffRampOverSuperchainInteropOwnershipTransferRequested)
	if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampOverSuperchainInteropOwnershipTransferredIterator struct {
	Event *OffRampOverSuperchainInteropOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampOverSuperchainInteropOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampOverSuperchainInteropOwnershipTransferred)
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
		it.Event = new(OffRampOverSuperchainInteropOwnershipTransferred)
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

func (it *OffRampOverSuperchainInteropOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *OffRampOverSuperchainInteropOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampOverSuperchainInteropOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OffRampOverSuperchainInteropOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffRampOverSuperchainInterop.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OffRampOverSuperchainInteropOwnershipTransferredIterator{contract: _OffRampOverSuperchainInterop.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffRampOverSuperchainInterop.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampOverSuperchainInteropOwnershipTransferred)
				if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) ParseOwnershipTransferred(log types.Log) (*OffRampOverSuperchainInteropOwnershipTransferred, error) {
	event := new(OffRampOverSuperchainInteropOwnershipTransferred)
	if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampOverSuperchainInteropRootRemovedIterator struct {
	Event *OffRampOverSuperchainInteropRootRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampOverSuperchainInteropRootRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampOverSuperchainInteropRootRemoved)
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
		it.Event = new(OffRampOverSuperchainInteropRootRemoved)
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

func (it *OffRampOverSuperchainInteropRootRemovedIterator) Error() error {
	return it.fail
}

func (it *OffRampOverSuperchainInteropRootRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampOverSuperchainInteropRootRemoved struct {
	Root [32]byte
	Raw  types.Log
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) FilterRootRemoved(opts *bind.FilterOpts) (*OffRampOverSuperchainInteropRootRemovedIterator, error) {

	logs, sub, err := _OffRampOverSuperchainInterop.contract.FilterLogs(opts, "RootRemoved")
	if err != nil {
		return nil, err
	}
	return &OffRampOverSuperchainInteropRootRemovedIterator{contract: _OffRampOverSuperchainInterop.contract, event: "RootRemoved", logs: logs, sub: sub}, nil
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) WatchRootRemoved(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropRootRemoved) (event.Subscription, error) {

	logs, sub, err := _OffRampOverSuperchainInterop.contract.WatchLogs(opts, "RootRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampOverSuperchainInteropRootRemoved)
				if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "RootRemoved", log); err != nil {
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

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) ParseRootRemoved(log types.Log) (*OffRampOverSuperchainInteropRootRemoved, error) {
	event := new(OffRampOverSuperchainInteropRootRemoved)
	if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "RootRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampOverSuperchainInteropSkippedAlreadyExecutedMessageIterator struct {
	Event *OffRampOverSuperchainInteropSkippedAlreadyExecutedMessage

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampOverSuperchainInteropSkippedAlreadyExecutedMessageIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampOverSuperchainInteropSkippedAlreadyExecutedMessage)
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
		it.Event = new(OffRampOverSuperchainInteropSkippedAlreadyExecutedMessage)
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

func (it *OffRampOverSuperchainInteropSkippedAlreadyExecutedMessageIterator) Error() error {
	return it.fail
}

func (it *OffRampOverSuperchainInteropSkippedAlreadyExecutedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampOverSuperchainInteropSkippedAlreadyExecutedMessage struct {
	SourceChainSelector uint64
	SequenceNumber      uint64
	Raw                 types.Log
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) FilterSkippedAlreadyExecutedMessage(opts *bind.FilterOpts) (*OffRampOverSuperchainInteropSkippedAlreadyExecutedMessageIterator, error) {

	logs, sub, err := _OffRampOverSuperchainInterop.contract.FilterLogs(opts, "SkippedAlreadyExecutedMessage")
	if err != nil {
		return nil, err
	}
	return &OffRampOverSuperchainInteropSkippedAlreadyExecutedMessageIterator{contract: _OffRampOverSuperchainInterop.contract, event: "SkippedAlreadyExecutedMessage", logs: logs, sub: sub}, nil
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) WatchSkippedAlreadyExecutedMessage(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropSkippedAlreadyExecutedMessage) (event.Subscription, error) {

	logs, sub, err := _OffRampOverSuperchainInterop.contract.WatchLogs(opts, "SkippedAlreadyExecutedMessage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampOverSuperchainInteropSkippedAlreadyExecutedMessage)
				if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "SkippedAlreadyExecutedMessage", log); err != nil {
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

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) ParseSkippedAlreadyExecutedMessage(log types.Log) (*OffRampOverSuperchainInteropSkippedAlreadyExecutedMessage, error) {
	event := new(OffRampOverSuperchainInteropSkippedAlreadyExecutedMessage)
	if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "SkippedAlreadyExecutedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampOverSuperchainInteropSkippedReportExecutionIterator struct {
	Event *OffRampOverSuperchainInteropSkippedReportExecution

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampOverSuperchainInteropSkippedReportExecutionIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampOverSuperchainInteropSkippedReportExecution)
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
		it.Event = new(OffRampOverSuperchainInteropSkippedReportExecution)
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

func (it *OffRampOverSuperchainInteropSkippedReportExecutionIterator) Error() error {
	return it.fail
}

func (it *OffRampOverSuperchainInteropSkippedReportExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampOverSuperchainInteropSkippedReportExecution struct {
	SourceChainSelector uint64
	Raw                 types.Log
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) FilterSkippedReportExecution(opts *bind.FilterOpts) (*OffRampOverSuperchainInteropSkippedReportExecutionIterator, error) {

	logs, sub, err := _OffRampOverSuperchainInterop.contract.FilterLogs(opts, "SkippedReportExecution")
	if err != nil {
		return nil, err
	}
	return &OffRampOverSuperchainInteropSkippedReportExecutionIterator{contract: _OffRampOverSuperchainInterop.contract, event: "SkippedReportExecution", logs: logs, sub: sub}, nil
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) WatchSkippedReportExecution(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropSkippedReportExecution) (event.Subscription, error) {

	logs, sub, err := _OffRampOverSuperchainInterop.contract.WatchLogs(opts, "SkippedReportExecution")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampOverSuperchainInteropSkippedReportExecution)
				if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "SkippedReportExecution", log); err != nil {
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

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) ParseSkippedReportExecution(log types.Log) (*OffRampOverSuperchainInteropSkippedReportExecution, error) {
	event := new(OffRampOverSuperchainInteropSkippedReportExecution)
	if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "SkippedReportExecution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampOverSuperchainInteropSourceChainConfigSetIterator struct {
	Event *OffRampOverSuperchainInteropSourceChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampOverSuperchainInteropSourceChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampOverSuperchainInteropSourceChainConfigSet)
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
		it.Event = new(OffRampOverSuperchainInteropSourceChainConfigSet)
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

func (it *OffRampOverSuperchainInteropSourceChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *OffRampOverSuperchainInteropSourceChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampOverSuperchainInteropSourceChainConfigSet struct {
	SourceChainSelector uint64
	SourceConfig        OffRampSourceChainConfig
	Raw                 types.Log
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) FilterSourceChainConfigSet(opts *bind.FilterOpts, sourceChainSelector []uint64) (*OffRampOverSuperchainInteropSourceChainConfigSetIterator, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}

	logs, sub, err := _OffRampOverSuperchainInterop.contract.FilterLogs(opts, "SourceChainConfigSet", sourceChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &OffRampOverSuperchainInteropSourceChainConfigSetIterator{contract: _OffRampOverSuperchainInterop.contract, event: "SourceChainConfigSet", logs: logs, sub: sub}, nil
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) WatchSourceChainConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropSourceChainConfigSet, sourceChainSelector []uint64) (event.Subscription, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}

	logs, sub, err := _OffRampOverSuperchainInterop.contract.WatchLogs(opts, "SourceChainConfigSet", sourceChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampOverSuperchainInteropSourceChainConfigSet)
				if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "SourceChainConfigSet", log); err != nil {
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

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) ParseSourceChainConfigSet(log types.Log) (*OffRampOverSuperchainInteropSourceChainConfigSet, error) {
	event := new(OffRampOverSuperchainInteropSourceChainConfigSet)
	if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "SourceChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampOverSuperchainInteropSourceChainSelectorAddedIterator struct {
	Event *OffRampOverSuperchainInteropSourceChainSelectorAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampOverSuperchainInteropSourceChainSelectorAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampOverSuperchainInteropSourceChainSelectorAdded)
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
		it.Event = new(OffRampOverSuperchainInteropSourceChainSelectorAdded)
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

func (it *OffRampOverSuperchainInteropSourceChainSelectorAddedIterator) Error() error {
	return it.fail
}

func (it *OffRampOverSuperchainInteropSourceChainSelectorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampOverSuperchainInteropSourceChainSelectorAdded struct {
	SourceChainSelector uint64
	Raw                 types.Log
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) FilterSourceChainSelectorAdded(opts *bind.FilterOpts) (*OffRampOverSuperchainInteropSourceChainSelectorAddedIterator, error) {

	logs, sub, err := _OffRampOverSuperchainInterop.contract.FilterLogs(opts, "SourceChainSelectorAdded")
	if err != nil {
		return nil, err
	}
	return &OffRampOverSuperchainInteropSourceChainSelectorAddedIterator{contract: _OffRampOverSuperchainInterop.contract, event: "SourceChainSelectorAdded", logs: logs, sub: sub}, nil
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) WatchSourceChainSelectorAdded(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropSourceChainSelectorAdded) (event.Subscription, error) {

	logs, sub, err := _OffRampOverSuperchainInterop.contract.WatchLogs(opts, "SourceChainSelectorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampOverSuperchainInteropSourceChainSelectorAdded)
				if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "SourceChainSelectorAdded", log); err != nil {
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

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) ParseSourceChainSelectorAdded(log types.Log) (*OffRampOverSuperchainInteropSourceChainSelectorAdded, error) {
	event := new(OffRampOverSuperchainInteropSourceChainSelectorAdded)
	if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "SourceChainSelectorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampOverSuperchainInteropStaticConfigSetIterator struct {
	Event *OffRampOverSuperchainInteropStaticConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampOverSuperchainInteropStaticConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampOverSuperchainInteropStaticConfigSet)
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
		it.Event = new(OffRampOverSuperchainInteropStaticConfigSet)
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

func (it *OffRampOverSuperchainInteropStaticConfigSetIterator) Error() error {
	return it.fail
}

func (it *OffRampOverSuperchainInteropStaticConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampOverSuperchainInteropStaticConfigSet struct {
	StaticConfig OffRampStaticConfig
	Raw          types.Log
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) FilterStaticConfigSet(opts *bind.FilterOpts) (*OffRampOverSuperchainInteropStaticConfigSetIterator, error) {

	logs, sub, err := _OffRampOverSuperchainInterop.contract.FilterLogs(opts, "StaticConfigSet")
	if err != nil {
		return nil, err
	}
	return &OffRampOverSuperchainInteropStaticConfigSetIterator{contract: _OffRampOverSuperchainInterop.contract, event: "StaticConfigSet", logs: logs, sub: sub}, nil
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) WatchStaticConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropStaticConfigSet) (event.Subscription, error) {

	logs, sub, err := _OffRampOverSuperchainInterop.contract.WatchLogs(opts, "StaticConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampOverSuperchainInteropStaticConfigSet)
				if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "StaticConfigSet", log); err != nil {
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

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) ParseStaticConfigSet(log types.Log) (*OffRampOverSuperchainInteropStaticConfigSet, error) {
	event := new(OffRampOverSuperchainInteropStaticConfigSet)
	if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "StaticConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampOverSuperchainInteropTransmittedIterator struct {
	Event *OffRampOverSuperchainInteropTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampOverSuperchainInteropTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampOverSuperchainInteropTransmitted)
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
		it.Event = new(OffRampOverSuperchainInteropTransmitted)
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

func (it *OffRampOverSuperchainInteropTransmittedIterator) Error() error {
	return it.fail
}

func (it *OffRampOverSuperchainInteropTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampOverSuperchainInteropTransmitted struct {
	OcrPluginType  uint8
	ConfigDigest   [32]byte
	SequenceNumber uint64
	Raw            types.Log
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) FilterTransmitted(opts *bind.FilterOpts, ocrPluginType []uint8) (*OffRampOverSuperchainInteropTransmittedIterator, error) {

	var ocrPluginTypeRule []interface{}
	for _, ocrPluginTypeItem := range ocrPluginType {
		ocrPluginTypeRule = append(ocrPluginTypeRule, ocrPluginTypeItem)
	}

	logs, sub, err := _OffRampOverSuperchainInterop.contract.FilterLogs(opts, "Transmitted", ocrPluginTypeRule)
	if err != nil {
		return nil, err
	}
	return &OffRampOverSuperchainInteropTransmittedIterator{contract: _OffRampOverSuperchainInterop.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropTransmitted, ocrPluginType []uint8) (event.Subscription, error) {

	var ocrPluginTypeRule []interface{}
	for _, ocrPluginTypeItem := range ocrPluginType {
		ocrPluginTypeRule = append(ocrPluginTypeRule, ocrPluginTypeItem)
	}

	logs, sub, err := _OffRampOverSuperchainInterop.contract.WatchLogs(opts, "Transmitted", ocrPluginTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampOverSuperchainInteropTransmitted)
				if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropFilterer) ParseTransmitted(log types.Log) (*OffRampOverSuperchainInteropTransmitted, error) {
	event := new(OffRampOverSuperchainInteropTransmitted)
	if err := _OffRampOverSuperchainInterop.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInterop) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _OffRampOverSuperchainInterop.abi.Events["AlreadyAttempted"].ID:
		return _OffRampOverSuperchainInterop.ParseAlreadyAttempted(log)
	case _OffRampOverSuperchainInterop.abi.Events["ChainSelectorToChainIdConfigRemoved"].ID:
		return _OffRampOverSuperchainInterop.ParseChainSelectorToChainIdConfigRemoved(log)
	case _OffRampOverSuperchainInterop.abi.Events["ChainSelectorToChainIdConfigUpdated"].ID:
		return _OffRampOverSuperchainInterop.ParseChainSelectorToChainIdConfigUpdated(log)
	case _OffRampOverSuperchainInterop.abi.Events["CommitReportAccepted"].ID:
		return _OffRampOverSuperchainInterop.ParseCommitReportAccepted(log)
	case _OffRampOverSuperchainInterop.abi.Events["ConfigSet"].ID:
		return _OffRampOverSuperchainInterop.ParseConfigSet(log)
	case _OffRampOverSuperchainInterop.abi.Events["DynamicConfigSet"].ID:
		return _OffRampOverSuperchainInterop.ParseDynamicConfigSet(log)
	case _OffRampOverSuperchainInterop.abi.Events["ExecutionStateChanged"].ID:
		return _OffRampOverSuperchainInterop.ParseExecutionStateChanged(log)
	case _OffRampOverSuperchainInterop.abi.Events["OwnershipTransferRequested"].ID:
		return _OffRampOverSuperchainInterop.ParseOwnershipTransferRequested(log)
	case _OffRampOverSuperchainInterop.abi.Events["OwnershipTransferred"].ID:
		return _OffRampOverSuperchainInterop.ParseOwnershipTransferred(log)
	case _OffRampOverSuperchainInterop.abi.Events["RootRemoved"].ID:
		return _OffRampOverSuperchainInterop.ParseRootRemoved(log)
	case _OffRampOverSuperchainInterop.abi.Events["SkippedAlreadyExecutedMessage"].ID:
		return _OffRampOverSuperchainInterop.ParseSkippedAlreadyExecutedMessage(log)
	case _OffRampOverSuperchainInterop.abi.Events["SkippedReportExecution"].ID:
		return _OffRampOverSuperchainInterop.ParseSkippedReportExecution(log)
	case _OffRampOverSuperchainInterop.abi.Events["SourceChainConfigSet"].ID:
		return _OffRampOverSuperchainInterop.ParseSourceChainConfigSet(log)
	case _OffRampOverSuperchainInterop.abi.Events["SourceChainSelectorAdded"].ID:
		return _OffRampOverSuperchainInterop.ParseSourceChainSelectorAdded(log)
	case _OffRampOverSuperchainInterop.abi.Events["StaticConfigSet"].ID:
		return _OffRampOverSuperchainInterop.ParseStaticConfigSet(log)
	case _OffRampOverSuperchainInterop.abi.Events["Transmitted"].ID:
		return _OffRampOverSuperchainInterop.ParseTransmitted(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (OffRampOverSuperchainInteropAlreadyAttempted) Topic() common.Hash {
	return common.HexToHash("0x3ef2a99c550a751d4b0b261268f05a803dfb049ab43616a1ffb388f61fe65120")
}

func (OffRampOverSuperchainInteropChainSelectorToChainIdConfigRemoved) Topic() common.Hash {
	return common.HexToHash("0xb56b587763154465d175e8a2a97978dffe45711125145973789b83ab201e702a")
}

func (OffRampOverSuperchainInteropChainSelectorToChainIdConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x5b4e1378b67677cfb7d6b50a37fd7f632168140928b0f970077a372bdfb42a3e")
}

func (OffRampOverSuperchainInteropCommitReportAccepted) Topic() common.Hash {
	return common.HexToHash("0xb967c9b9e1b7af9a61ca71ff00e9f5b89ec6f2e268de8dacf12f0de8e51f3e47")
}

func (OffRampOverSuperchainInteropConfigSet) Topic() common.Hash {
	return common.HexToHash("0xab8b1b57514019638d7b5ce9c638fe71366fe8e2be1c40a7a80f1733d0e9f547")
}

func (OffRampOverSuperchainInteropDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0xa1c15688cb2c24508e158f6942b9276c6f3028a85e1af8cf3fff0c3ff3d5fc8d")
}

func (OffRampOverSuperchainInteropExecutionStateChanged) Topic() common.Hash {
	return common.HexToHash("0x05665fe9ad095383d018353f4cbcba77e84db27dd215081bbf7cdf9ae6fbe48b")
}

func (OffRampOverSuperchainInteropOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (OffRampOverSuperchainInteropOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (OffRampOverSuperchainInteropRootRemoved) Topic() common.Hash {
	return common.HexToHash("0x202f1139a3e334b6056064c0e9b19fd07e44a88d8f6e5ded571b24cf8c371f12")
}

func (OffRampOverSuperchainInteropSkippedAlreadyExecutedMessage) Topic() common.Hash {
	return common.HexToHash("0x3b575419319662b2a6f5e2467d84521517a3382b908eb3d557bb3fdb0c50e23c")
}

func (OffRampOverSuperchainInteropSkippedReportExecution) Topic() common.Hash {
	return common.HexToHash("0xaab522ed53d887e56ed53dd37398a01aeef6a58e0fa77c2173beb9512d894933")
}

func (OffRampOverSuperchainInteropSourceChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0xbd1ab25a0ff0a36a588597ba1af11e30f3f210de8b9e818cc9bbc457c94c8d8c")
}

func (OffRampOverSuperchainInteropSourceChainSelectorAdded) Topic() common.Hash {
	return common.HexToHash("0xf4c1390c70e5c0f491ae1ccbc06f9117cbbadf2767b247b3bc203280f24c0fb9")
}

func (OffRampOverSuperchainInteropStaticConfigSet) Topic() common.Hash {
	return common.HexToHash("0xb0fa1fb01508c5097c502ad056fd77018870c9be9a86d9e56b6b471862d7c5b7")
}

func (OffRampOverSuperchainInteropTransmitted) Topic() common.Hash {
	return common.HexToHash("0x198d6990ef96613a9026203077e422916918b03ff47f0be6bee7b02d8e139ef0")
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInterop) Address() common.Address {
	return _OffRampOverSuperchainInterop.address
}

type OffRampOverSuperchainInteropInterface interface {
	CcipReceive(opts *bind.CallOpts, arg0 ClientAny2EVMMessage) error

	Commit(opts *bind.CallOpts, arg0 [2][32]byte, arg1 []byte, arg2 [][32]byte, arg3 [][32]byte, arg4 [32]byte) error

	GetAllSourceChainConfigs(opts *bind.CallOpts) ([]uint64, []OffRampSourceChainConfig, error)

	GetChainId(opts *bind.CallOpts, sourceChainSelector uint64) (*big.Int, error)

	GetCrossL2Inbox(opts *bind.CallOpts) (common.Address, error)

	GetDynamicConfig(opts *bind.CallOpts) (OffRampDynamicConfig, error)

	GetExecutionState(opts *bind.CallOpts, sourceChainSelector uint64, sequenceNumber uint64) (uint8, error)

	GetLatestPriceSequenceNumber(opts *bind.CallOpts) (uint64, error)

	GetMerkleRoot(opts *bind.CallOpts, sourceChainSelector uint64, root [32]byte) (*big.Int, error)

	GetSourceChainConfig(opts *bind.CallOpts, sourceChainSelector uint64) (OffRampSourceChainConfig, error)

	GetStaticConfig(opts *bind.CallOpts) (OffRampStaticConfig, error)

	LatestConfigDetails(opts *bind.CallOpts, ocrPluginType uint8) (MultiOCR3BaseOCRConfig, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyChainSelectorToChainIdConfigUpdates(opts *bind.TransactOpts, chainSelectorsToUnset []uint64, chainSelectorsToSet []OffRampOverSuperchainInteropChainSelectorToChainIdConfigArgs) (*types.Transaction, error)

	ApplySourceChainConfigUpdates(opts *bind.TransactOpts, sourceChainConfigUpdates []OffRampSourceChainConfigArgs) (*types.Transaction, error)

	Execute(opts *bind.TransactOpts, reportContext [2][32]byte, report []byte) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message InternalAny2EVMRampMessage, offchainTokenData [][]byte, tokenGasOverrides []uint32) (*types.Transaction, error)

	ManuallyExecute(opts *bind.TransactOpts, reports []InternalExecutionReport, gasLimitOverrides [][]OffRampGasLimitOverride) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig OffRampDynamicConfig) (*types.Transaction, error)

	SetOCR3Configs(opts *bind.TransactOpts, ocrConfigArgs []MultiOCR3BaseOCRConfigArgs) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterAlreadyAttempted(opts *bind.FilterOpts) (*OffRampOverSuperchainInteropAlreadyAttemptedIterator, error)

	WatchAlreadyAttempted(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropAlreadyAttempted) (event.Subscription, error)

	ParseAlreadyAttempted(log types.Log) (*OffRampOverSuperchainInteropAlreadyAttempted, error)

	FilterChainSelectorToChainIdConfigRemoved(opts *bind.FilterOpts, chainSelector []uint64, chainId []*big.Int) (*OffRampOverSuperchainInteropChainSelectorToChainIdConfigRemovedIterator, error)

	WatchChainSelectorToChainIdConfigRemoved(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropChainSelectorToChainIdConfigRemoved, chainSelector []uint64, chainId []*big.Int) (event.Subscription, error)

	ParseChainSelectorToChainIdConfigRemoved(log types.Log) (*OffRampOverSuperchainInteropChainSelectorToChainIdConfigRemoved, error)

	FilterChainSelectorToChainIdConfigUpdated(opts *bind.FilterOpts, chainSelector []uint64, chainId []*big.Int) (*OffRampOverSuperchainInteropChainSelectorToChainIdConfigUpdatedIterator, error)

	WatchChainSelectorToChainIdConfigUpdated(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropChainSelectorToChainIdConfigUpdated, chainSelector []uint64, chainId []*big.Int) (event.Subscription, error)

	ParseChainSelectorToChainIdConfigUpdated(log types.Log) (*OffRampOverSuperchainInteropChainSelectorToChainIdConfigUpdated, error)

	FilterCommitReportAccepted(opts *bind.FilterOpts) (*OffRampOverSuperchainInteropCommitReportAcceptedIterator, error)

	WatchCommitReportAccepted(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropCommitReportAccepted) (event.Subscription, error)

	ParseCommitReportAccepted(log types.Log) (*OffRampOverSuperchainInteropCommitReportAccepted, error)

	FilterConfigSet(opts *bind.FilterOpts) (*OffRampOverSuperchainInteropConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*OffRampOverSuperchainInteropConfigSet, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*OffRampOverSuperchainInteropDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*OffRampOverSuperchainInteropDynamicConfigSet, error)

	FilterExecutionStateChanged(opts *bind.FilterOpts, sourceChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (*OffRampOverSuperchainInteropExecutionStateChangedIterator, error)

	WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropExecutionStateChanged, sourceChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error)

	ParseExecutionStateChanged(log types.Log) (*OffRampOverSuperchainInteropExecutionStateChanged, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OffRampOverSuperchainInteropOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*OffRampOverSuperchainInteropOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OffRampOverSuperchainInteropOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*OffRampOverSuperchainInteropOwnershipTransferred, error)

	FilterRootRemoved(opts *bind.FilterOpts) (*OffRampOverSuperchainInteropRootRemovedIterator, error)

	WatchRootRemoved(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropRootRemoved) (event.Subscription, error)

	ParseRootRemoved(log types.Log) (*OffRampOverSuperchainInteropRootRemoved, error)

	FilterSkippedAlreadyExecutedMessage(opts *bind.FilterOpts) (*OffRampOverSuperchainInteropSkippedAlreadyExecutedMessageIterator, error)

	WatchSkippedAlreadyExecutedMessage(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropSkippedAlreadyExecutedMessage) (event.Subscription, error)

	ParseSkippedAlreadyExecutedMessage(log types.Log) (*OffRampOverSuperchainInteropSkippedAlreadyExecutedMessage, error)

	FilterSkippedReportExecution(opts *bind.FilterOpts) (*OffRampOverSuperchainInteropSkippedReportExecutionIterator, error)

	WatchSkippedReportExecution(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropSkippedReportExecution) (event.Subscription, error)

	ParseSkippedReportExecution(log types.Log) (*OffRampOverSuperchainInteropSkippedReportExecution, error)

	FilterSourceChainConfigSet(opts *bind.FilterOpts, sourceChainSelector []uint64) (*OffRampOverSuperchainInteropSourceChainConfigSetIterator, error)

	WatchSourceChainConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropSourceChainConfigSet, sourceChainSelector []uint64) (event.Subscription, error)

	ParseSourceChainConfigSet(log types.Log) (*OffRampOverSuperchainInteropSourceChainConfigSet, error)

	FilterSourceChainSelectorAdded(opts *bind.FilterOpts) (*OffRampOverSuperchainInteropSourceChainSelectorAddedIterator, error)

	WatchSourceChainSelectorAdded(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropSourceChainSelectorAdded) (event.Subscription, error)

	ParseSourceChainSelectorAdded(log types.Log) (*OffRampOverSuperchainInteropSourceChainSelectorAdded, error)

	FilterStaticConfigSet(opts *bind.FilterOpts) (*OffRampOverSuperchainInteropStaticConfigSetIterator, error)

	WatchStaticConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropStaticConfigSet) (event.Subscription, error)

	ParseStaticConfigSet(log types.Log) (*OffRampOverSuperchainInteropStaticConfigSet, error)

	FilterTransmitted(opts *bind.FilterOpts, ocrPluginType []uint8) (*OffRampOverSuperchainInteropTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *OffRampOverSuperchainInteropTransmitted, ocrPluginType []uint8) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*OffRampOverSuperchainInteropTransmitted, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
