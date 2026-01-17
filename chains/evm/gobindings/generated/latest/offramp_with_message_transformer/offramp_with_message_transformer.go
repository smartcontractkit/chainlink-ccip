// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package offramp_with_message_transformer

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

var OffRampWithMessageTransformerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structOffRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structOffRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"sourceChainConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structOffRamp.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isRMNVerificationDisabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageTransformerAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySourceChainConfigUpdates\",\"inputs\":[{\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structOffRamp.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isRMNVerificationDisabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipReceive\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structClient.Any2EVMMessage\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"commit\",\"inputs\":[{\"name\":\"reportContext\",\"type\":\"bytes32[2]\",\"internalType\":\"bytes32[2]\"},{\"name\":\"report\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"rs\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"ss\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"rawVs\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"reportContext\",\"type\":\"bytes32[2]\",\"internalType\":\"bytes32[2]\"},{\"name\":\"report\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSingleMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInternal.Any2EVMRampMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.RampMessageHeader\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Any2EVMTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasAmount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"offchainTokenData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"tokenGasOverrides\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllSourceChainConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOffRamp.SourceChainConfig[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"minSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isRMNVerificationDisabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOffRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExecutionState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumInternal.MessageExecutionState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestPriceSequenceNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMerkleRoot\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMessageTransformer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSourceChainConfig\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOffRamp.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"minSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isRMNVerificationDisabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOffRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestConfigDetails\",\"inputs\":[{\"name\":\"ocrPluginType\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"ocrConfig\",\"type\":\"tuple\",\"internalType\":\"structMultiOCR3Base.OCRConfig\",\"components\":[{\"name\":\"configInfo\",\"type\":\"tuple\",\"internalType\":\"structMultiOCR3Base.ConfigInfo\",\"components\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"F\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"n\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isSignatureVerificationEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"transmitters\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"manuallyExecute\",\"inputs\":[{\"name\":\"reports\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.ExecutionReport[]\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messages\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Any2EVMRampMessage[]\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.RampMessageHeader\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Any2EVMTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasAmount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"offchainTokenData\",\"type\":\"bytes[][]\",\"internalType\":\"bytes[][]\"},{\"name\":\"proofs\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"proofFlagBits\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"gasLimitOverrides\",\"type\":\"tuple[][]\",\"internalType\":\"structOffRamp.GasLimitOverride[][]\",\"components\":[{\"name\":\"receiverExecutionGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenGasOverrides\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structOffRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMessageTransformer\",\"inputs\":[{\"name\":\"messageTransformerAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOCR3Configs\",\"inputs\":[{\"name\":\"ocrConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structMultiOCR3Base.OCRConfigArgs[]\",\"components\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ocrPluginType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"F\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isSignatureVerificationEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"transmitters\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"AlreadyAttempted\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CommitReportAccepted\",\"inputs\":[{\"name\":\"blessedMerkleRoots\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structInternal.MerkleRoot[]\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"minSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"unblessedMerkleRoots\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structInternal.MerkleRoot[]\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"minSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"priceUpdates\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInternal.PriceUpdates\",\"components\":[{\"name\":\"tokenPriceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.TokenPriceUpdate[]\",\"components\":[{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"usdPerToken\",\"type\":\"uint224\",\"internalType\":\"uint224\"}]},{\"name\":\"gasPriceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.GasPriceUpdate[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"usdPerUnitGas\",\"type\":\"uint224\",\"internalType\":\"uint224\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"ocrPluginType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"configDigest\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"signers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"transmitters\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"F\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOffRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutionStateChanged\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"state\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"gasUsed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RootRemoved\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SkippedAlreadyExecutedMessage\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SkippedReportExecution\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SourceChainConfigSet\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sourceConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOffRamp.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"minSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isRMNVerificationDisabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SourceChainSelectorAdded\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOffRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transmitted\",\"inputs\":[{\"name\":\"ocrPluginType\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"configDigest\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySelfCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CommitOnRampMismatch\",\"inputs\":[{\"name\":\"reportOnRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"configOnRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ConfigDigestMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"EmptyBatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyReport\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ExecutionError\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ForkedChain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasToCompleteTx\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[{\"name\":\"errorType\",\"type\":\"uint8\",\"internalType\":\"enumMultiOCR3Base.InvalidConfigErrorType\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidInterval\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"min\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"max\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidManualExecutionGasLimit\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"newLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidManualExecutionTokenGasOverride\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tokenIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"oldLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenGasOverride\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageDestChainSelector\",\"inputs\":[{\"name\":\"messageDestChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNewState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newState\",\"type\":\"uint8\",\"internalType\":\"enumInternal.MessageExecutionState\"}]},{\"type\":\"error\",\"name\":\"InvalidOnRampUpdate\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LeavesCannotBeEmpty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ManualExecutionGasAmountCountMismatch\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ManualExecutionGasLimitMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ManualExecutionNotYetEnabled\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MessageTransformError\",\"inputs\":[{\"name\":\"errorReason\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"MessageValidationError\",\"inputs\":[{\"name\":\"errorReason\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonUniqueSignatures\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotACompatiblePool\",\"inputs\":[{\"name\":\"notPool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OracleCannotBeZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiverError\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ReleaseOrMintBalanceMismatch\",\"inputs\":[{\"name\":\"amountReleased\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePre\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePost\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"RootAlreadyCommitted\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"RootBlessingMismatch\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"isBlessed\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"type\":\"error\",\"name\":\"RootNotCommitted\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SignatureVerificationNotAllowedInExecutionPlugin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureVerificationRequiredInCommitPlugin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignaturesOutOfRegistration\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SourceChainNotEnabled\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SourceChainSelectorMismatch\",\"inputs\":[{\"name\":\"reportSourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageSourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"StaleCommitReport\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StaticConfigCannotBeChanged\",\"inputs\":[{\"name\":\"ocrPluginType\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"TokenDataMismatch\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenHandlingError\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedTransmitter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnexpectedTokenData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongMessageLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"WrongNumberOfSignatures\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainSelectorNotAllowed\",\"inputs\":[]}]",
	Bin: "0x61014080604052346108a157616a9d803803809161001d82856108d7565b83398101908082039061014082126108a15760a082126108a157604051610043816108bc565b61004c826108fa565b815260208201519261ffff841684036108a1576020820193845260408301516001600160a01b03811681036108a1576040830190815261008e6060850161090e565b946060840195865260606100a46080870161090e565b6080860190815293609f1901126108a15760405193606085016001600160401b038111868210176108a6576040526100de60a0870161090e565b855260c08601519363ffffffff851685036108a1576020860194855261010660e0880161090e565b604087019081526101008801519097906001600160401b0381116108a15781018a601f820112156108a15780519a6001600160401b038c116108a6578b60051b916020806040519e8f9061015c838801836108d7565b81520193820101908282116108a15760208101935b82851061079057505050505061012061018a910161090e565b97331561077f57600180546001600160a01b031916331790554660805284516001600160a01b031615801561076d575b801561075b575b6107395782516001600160401b03161561074a5782516001600160401b0390811660a090815286516001600160a01b0390811660c0528351811660e0528451811661010052865161ffff90811661012052604080519751909416875296519096166020860152955185169084015251831660608301525190911660808201527fb0fa1fb01508c5097c502ad056fd77018870c9be9a86d9e56b6b471862d7c5b79190a181516001600160a01b03161561073957905160048054835163ffffffff60a01b60a09190911b166001600160a01b039384166001600160c01b03199092168217179091558351600580549184166001600160a01b031990921691909117905560408051918252925163ffffffff1660208201529251169082015282907fa1c15688cb2c24508e158f6942b9276c6f3028a85e1af8cf3fff0c3ff3d5fc8d90606090a16000915b815183101561067a5760009260208160051b8401015160018060401b0360208201511690811561066b5780516001600160a01b03161561065c57818652600860205260408620906080810151906001830191610366835461092f565b6105fd578354600160a81b600160e81b031916600160a81b1784556040518581527ff4c1390c70e5c0f491ae1ccbc06f9117cbbadf2767b247b3bc203280f24c0fb990602090a15b805180159081156105d2575b506105c3578051906001600160401b0382116105af576103da845461092f565b601f811161056a575b50602090601f83116001146104eb57600080516020616a7d8339815191529593836104d59460ff979460609460019d9e9f926104e0575b5050600019600383901b1c1916908b1b1783555b604081015115158554908760a01b9060a01b16908760a01b1916178555898060a01b038151168a8060a01b0319865416178555015115158354908560e81b9060e81b16908560e81b1916178355610484866109ec565b506040519384936020855254898060a01b0381166020860152818160a01c1615156040860152898060401b038160a81c16606086015260e81c161515608084015260a08084015260c0830190610969565b0390a201919061030a565b015190508e8061041a565b848b52818b20919a601f198416905b81811061055257509360018460ff9794829c9d9e6060956104d598600080516020616a7d8339815191529c9a10610539575b505050811b01835561042e565b015160001960f88460031b161c191690558e808061052c565b828d0151845560209c8d019c600190940193016104fa565b848b5260208b20601f840160051c810191602085106105a5575b601f0160051c01905b81811061059a57506103e3565b8b815560010161058d565b9091508190610584565b634e487b7160e01b8a52604160045260248afd5b6342bcdf7f60e11b8952600489fd5b9050602082012060405160208101908b8252602081526105f36040826108d7565b519020148a6103ba565b835460a81c6001600160401b0316600114158061062e575b156103ae57632105803760e11b89526004859052602489fd5b50604051610647816106408187610969565b03826108d7565b60208151910120815160208301201415610615565b6342bcdf7f60e11b8652600486fd5b63c656089560e01b8652600486fd5b6001600160a01b0381161561073957600b8054600160401b600160e01b031916604092831b600160401b600160e01b031617905551615ffd9081610a80823960805181613377015260a0518181816101bf015261521b015260c05181818161021501528181612fb60152818161388b01528181613b5f01526143fa015260e0518181816102440152614aef01526101005181818161027301526147630152610120518181816101e6015281816122ae01528181614be20152615b630152f35b6342bcdf7f60e11b60005260046000fd5b63c656089560e01b60005260046000fd5b5081516001600160a01b0316156101c1565b5080516001600160a01b0316156101ba565b639b15e16f60e01b60005260046000fd5b84516001600160401b0381116108a157820160a0818603601f1901126108a157604051906107bd826108bc565b60208101516001600160a01b03811681036108a15782526107e0604082016108fa565b60208301526107f160608201610922565b604083015261080260808201610922565b606083015260a08101516001600160401b0381116108a157602091010185601f820112156108a15780516001600160401b0381116108a65760405191610852601f8301601f1916602001846108d7565b81835287602083830101116108a15760005b82811061088c5750509181600060208096949581960101526080820152815201940193610171565b80602080928401015182828701015201610864565b600080fd5b634e487b7160e01b600052604160045260246000fd5b60a081019081106001600160401b038211176108a657604052565b601f909101601f19168101906001600160401b038211908210176108a657604052565b51906001600160401b03821682036108a157565b51906001600160a01b03821682036108a157565b519081151582036108a157565b90600182811c9216801561095f575b602083101461094957565b634e487b7160e01b600052602260045260246000fd5b91607f169161093e565b600092918154916109798361092f565b80835292600181169081156109cf575060011461099557505050565b60009081526020812093945091925b8383106109b5575060209250010190565b6001816020929493945483858701015201910191906109a4565b915050602093945060ff929192191683830152151560051b010190565b80600052600760205260406000205415600014610a7957600654680100000000000000008110156108a6576001810180600655811015610a63577ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f0181905560065460009182526007602052604090912055600190565b634e487b7160e01b600052603260045260246000fd5b5060009056fe6080604052600436101561001257600080fd5b60003560e01c806306285c691461017757806315777ab214610172578063181f5a771461016d5780633f4b04aa146101685780635215505b146101635780635e36480c1461015e5780635e7bb0081461015957806360987c201461015457806365b81aab1461014f5780636f9e320f1461014a5780637437ff9f1461014557806379ba50971461014057806385572ffb1461013b5780638da5cb5b14610136578063c673e58414610131578063ccd37ba31461012c578063cd19723714610127578063de5e0b9a14610122578063e9d68a8e1461011d578063f2fde38b14610118578063f58e03fc146101135763f716f99f1461010e57600080fd5b611993565b611876565b6117eb565b611742565b6116a6565b611546565b6114e3565b61141e565b611336565b611300565b611280565b6111e0565b61106b565b610fd5565b610f5a565b610d53565b610641565b6104de565b6103c2565b610363565b6102f1565b61018c565b600091031261018757565b600080fd5b34610187576000366003190112610187576101a5611ace565b506102ed6040516101b58161069a565b6001600160401b037f000000000000000000000000000000000000000000000000000000000000000016815261ffff7f00000000000000000000000000000000000000000000000000000000000000001660208201526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001660408201526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001660608201526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001660808201526040519182918291909160806001600160a01b038160a08401956001600160401b03815116855261ffff6020820151166020860152826040820151166040860152826060820151166060860152015116910152565b0390f35b346101875760003660031901126101875760206001600160a01b03600b5460401c16604051908152f35b60005b83811061032e5750506000910152565b818101518382015260200161031e565b906020916103578151809281855285808601910161031b565b601f01601f1916010190565b34610187576000366003190112610187576102ed60408051906103868183610726565b601182527f4f666652616d7020312e362e322d64657600000000000000000000000000000060208301525191829160208352602083019061033e565b346101875760003660031901126101875760206001600160401b03600b5416604051908152f35b9060a0608061043a936001600160a01b0381511684526020810151151560208501526001600160401b036040820151166040850152606081015115156060850152015191816080820152019061033e565b90565b6040810160408252825180915260206060830193019060005b8181106104bf575050506020818303910152815180825260208201916020808360051b8301019401926000915b83831061049257505050505090565b90919293946020806104b0600193601f1986820301875289516103e9565b97019301930191939290610483565b82516001600160401b0316855260209485019490920191600101610456565b34610187576000366003190112610187576006546104fb81610793565b906105096040519283610726565b808252601f1961051882610793565b0160005b8181106105da57505061052e81611b3c565b9060005b81811061054a5750506102ed6040519283928361043d565b8061058061056861055c6001946142e1565b6001600160401b031690565b6105728387611b96565b906001600160401b03169052565b6105be6105b96105a06105938488611b96565b516001600160401b031690565b6001600160401b03166000526008602052604060002090565b611c82565b6105c88287611b96565b526105d38186611b96565b5001610532565b6020906105e5611b0e565b8282870101520161051c565b6001600160401b0381160361018757565b359061060d826105f1565b565b634e487b7160e01b600052602160045260246000fd5b6004111561062f57565b61060f565b90600482101561062f5752565b34610187576040366003190112610187576020610675600435610663816105f1565b60243590610670826105f1565b611d27565b6106826040518092610634565bf35b634e487b7160e01b600052604160045260246000fd5b60a081019081106001600160401b038211176106b557604052565b610684565b604081019081106001600160401b038211176106b557604052565b606081019081106001600160401b038211176106b557604052565b608081019081106001600160401b038211176106b557604052565b60c081019081106001600160401b038211176106b557604052565b90601f801991011681019081106001600160401b038211176106b557604052565b6040519061060d60c083610726565b6040519061060d60a083610726565b6040519061060d608083610726565b6040519061060d61010083610726565b6040519061060d604083610726565b6001600160401b0381116106b55760051b60200190565b91908260a0910312610187576040516107c28161069a565b60808082948035845260208101356107d9816105f1565b602085015260408101356107ec816105f1565b604085015260608101356107ff816105f1565b6060850152013591610810836105f1565b0152565b6001600160401b0381116106b557601f01601f191660200190565b92919261083b82610814565b916108496040519384610726565b829481845281830111610187578281602093846000960137010152565b9080601f830112156101875781602061043a9335910161082f565b6001600160a01b0381160361018757565b359061060d82610881565b63ffffffff81160361018757565b359061060d8261089d565b81601f82011215610187578035906108cd82610793565b926108db6040519485610726565b82845260208085019360051b830101918183116101875760208101935b83851061090757505050505090565b84356001600160401b03811161018757820160a0818503601f19011261018757604051916109348361069a565b60208201356001600160401b0381116101875785602061095692850101610866565b8352604082013561096681610881565b6020840152610977606083016108ab565b60408401526080820135926001600160401b0384116101875760a0836109a4886020809881980101610866565b6060840152013560808201528152019401936108f8565b91909161014081840312610187576109d1610747565b926109dc81836107aa565b845260a08201356001600160401b03811161018757816109fd918401610866565b602085015260c08201356001600160401b0381116101875781610a21918401610866565b6040850152610a3260e08301610892565b606085015261010082013560808501526101208201356001600160401b03811161018757610a6092016108b6565b60a0830152565b9080601f83011215610187578135610a7e81610793565b92610a8c6040519485610726565b81845260208085019260051b820101918383116101875760208201905b838210610ab857505050505090565b81356001600160401b03811161018757602091610ada878480948801016109bb565b815201910190610aa9565b81601f8201121561018757803590610afc82610793565b92610b0a6040519485610726565b82845260208085019360051b830101918183116101875760208101935b838510610b3657505050505090565b84356001600160401b03811161018757820183603f82011215610187576020810135610b6181610793565b91610b6f6040519384610726565b8183526020808085019360051b83010101918683116101875760408201905b838210610ba8575050509082525060209485019401610b27565b81356001600160401b03811161018757602091610bcc8a8480809589010101610866565b815201910190610b8e565b929190610be381610793565b93610bf16040519586610726565b602085838152019160051b810192831161018757905b828210610c1357505050565b8135815260209182019101610c07565b9080601f830112156101875781602061043a93359101610bd7565b81601f8201121561018757803590610c5582610793565b92610c636040519485610726565b82845260208085019360051b830101918183116101875760208101935b838510610c8f57505050505090565b84356001600160401b03811161018757820160a0818503601f19011261018757610cb7610756565b91610cc460208301610602565b835260408201356001600160401b03811161018757856020610ce892850101610a67565b602084015260608201356001600160401b03811161018757856020610d0f92850101610ae5565b60408401526080820135926001600160401b0384116101875760a083610d3c886020809881980101610c23565b606084015201356080820152815201940193610c80565b34610187576040366003190112610187576004356001600160401b03811161018757610d83903690600401610c3e565b6024356001600160401b038111610187573660238201121561018757806004013591610dae83610793565b91610dbc6040519384610726565b8383526024602084019460051b820101903682116101875760248101945b828610610ded57610deb8585611d6f565b005b85356001600160401b03811161018757820136604382011215610187576024810135610e1881610793565b91610e266040519384610726565b818352602060248185019360051b83010101903682116101875760448101925b828410610e60575050509082525060209586019501610dda565b83356001600160401b038111610187576024908301016040601f1982360301126101875760405190610e91826106ba565b6020810135825260408101356001600160401b03811161018757602091010136601f8201121561018757803590610ec782610793565b91610ed56040519384610726565b80835260208084019160051b8301019136831161018757602001905b828210610f105750505091816020938480940152815201930192610e46565b602080918335610f1f8161089d565b815201910190610ef1565b9181601f84011215610187578235916001600160401b038311610187576020808501948460051b01011161018757565b34610187576060366003190112610187576004356001600160401b03811161018757610f8a9036906004016109bb565b6024356001600160401b03811161018757610fa9903690600401610f2a565b91604435926001600160401b03841161018757610fcd610deb943690600401610f2a565b939092612186565b3461018757602036600319011261018757600435610ff281610881565b610ffa61362d565b6001600160a01b0381161561105a577fffffffff0000000000000000000000000000000000000000ffffffffffffffff7bffffffffffffffffffffffffffffffffffffffff0000000000000000600b549260401b16911617600b55600080f35b6342bcdf7f60e11b60005260046000fd5b3461018757606036600319011261018757600060405161108a816106d5565b60043561109681610881565b81526024356110a48161089d565b60208201908152604435906110b882610881565b604083019182526110c761362d565b6001600160a01b03835116156111d157916111936001600160a01b036111cb937fa1c15688cb2c24508e158f6942b9276c6f3028a85e1af8cf3fff0c3ff3d5fc8d9561112c838651166001600160a01b03166001600160a01b03196004541617600455565b517fffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffff77ffffffff00000000000000000000000000000000000000006004549260a01b1691161760045551166001600160a01b03166001600160a01b03196005541617600555565b6040519182918291909160406001600160a01b0381606084019582815116855263ffffffff6020820151166020860152015116910152565b0390a180f35b6342bcdf7f60e11b8452600484fd5b3461018757600036600319011261018757600060408051611200816106d5565b82815282602082015201526102ed60405161121a816106d5565b63ffffffff6004546001600160a01b038116835260a01c1660208201526001600160a01b036005541660408201526040519182918291909160406001600160a01b0381606084019582815116855263ffffffff6020820151166020860152015116910152565b34610187576000366003190112610187576000546001600160a01b03811633036112ef576001600160a01b0319600154913382841617600155166000556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b63015aa1e360e11b60005260046000fd5b34610187576020366003190112610187576004356001600160401b0381116101875760a090600319903603011261018757600080fd5b346101875760003660031901126101875760206001600160a01b0360015416604051908152f35b6004359060ff8216820361018757565b359060ff8216820361018757565b906020808351928381520192019060005b8181106113995750505090565b82516001600160a01b031684526020938401939092019160010161138c565b9061043a9160208152606082518051602084015260ff602082015116604084015260ff604082015116828401520151151560808201526040611409602084015160c060a085015260e084019061137b565b9201519060c0601f198285030191015261137b565b346101875760203660031901126101875760ff61143961135d565b606060408051611448816106d5565b8151611453816106f0565b6000815260006020820152600083820152600084820152815282602082015201521660005260026020526102ed604060002060036114d260405192611497846106d5565b6114a081612463565b84526040516114bd816114b6816002860161249c565b0382610726565b60208501526114b6604051809481930161249c565b6040820152604051918291826113b8565b3461018757604036600319011261018757600435611500816105f1565b6001600160401b036024359116600052600a6020526040600020906000526020526020604060002054604051908152f35b8015150361018757565b359061060d82611531565b34610187576020366003190112610187576004356001600160401b03811161018757366023820112156101875780600401359061158282610793565b906115906040519283610726565b8282526024602083019360051b820101903682116101875760248101935b8285106115be57610deb846124f3565b84356001600160401b03811161018757820160a0602319823603011261018757604051916115eb8361069a565b60248201356115f981610881565b83526044820135611609816105f1565b6020840152606482013561161c81611531565b6040840152608482013561162f81611531565b606084015260a4820135926001600160401b0384116101875761165c602094936024869536920101610866565b60808201528152019401936115ae565b9060049160441161018757565b9181601f84011215610187578235916001600160401b038311610187576020838186019501011161018757565b346101875760c0366003190112610187576116c03661166c565b6044356001600160401b038111610187576116df903690600401611679565b6064929192356001600160401b03811161018757611701903690600401610f2a565b60843594916001600160401b03861161018757611725610deb963690600401610f2a565b94909360a43596612db2565b90602061043a9281815201906103e9565b34610187576020366003190112610187576001600160401b03600435611767816105f1565b61176f611b0e565b501660005260086020526102ed60406000206117da6001604051926117938461069a565b6117d460ff82546001600160a01b0381168752818160a01c16151560208801526001600160401b038160a81c16604088015260e81c16606086019015159052565b01611c67565b608082015260405191829182611731565b34610187576020366003190112610187576001600160a01b0360043561181081610881565b61181861362d565b1633811461186557806001600160a01b031960005416176000556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b636d6c4ee560e11b60005260046000fd5b34610187576060366003190112610187576118903661166c565b6044356001600160401b038111610187576118af903690600401611679565b91828201602083820312610187578235906001600160401b038211610187576118d9918401610c3e565b6040519060206118e98184610726565b60008352601f19810160005b81811061191d57505050610deb949161190d916133b8565b61191561302c565b928392613ca9565b606085820184015282016118f5565b9080601f8301121561018757813561194381610793565b926119516040519485610726565b81845260208085019260051b82010192831161018757602001905b8282106119795750505090565b60208091833561198881610881565b81520191019061196c565b34610187576020366003190112610187576004356001600160401b0381116101875736602382011215610187578060040135906119cf82610793565b906119dd6040519283610726565b8282526024602083019360051b820101903682116101875760248101935b828510611a0b57610deb84613048565b84356001600160401b03811161018757820160c0602319823603011261018757611a33610747565b9160248201358352611a476044830161136d565b6020840152611a586064830161136d565b6040840152611a696084830161153b565b606084015260a48201356001600160401b03811161018757611a91906024369185010161192c565b608084015260c4820135926001600160401b03841161018757611abe60209493602486953692010161192c565b60a08201528152019401936119fb565b60405190611adb8261069a565b60006080838281528260208201528260408201528260608201520152565b60405190611b08602083610726565b60008252565b60405190611b1b8261069a565b60606080836000815260006020820152600060408201526000838201520152565b90611b4682610793565b611b536040519182610726565b8281528092611b64601f1991610793565b0190602036910137565b634e487b7160e01b600052603260045260246000fd5b805115611b915760200190565b611b6e565b8051821015611b915760209160051b010190565b90600182811c92168015611bda575b6020831014611bc457565b634e487b7160e01b600052602260045260246000fd5b91607f1691611bb9565b60009291815491611bf483611baa565b8083529260018116908115611c4a5750600114611c1057505050565b60009081526020812093945091925b838310611c30575060209250010190565b600181602092949394548385870101520191019190611c1f565b915050602093945060ff929192191683830152151560051b010190565b9061060d611c7b9260405193848092611be4565b0383610726565b9060016080604051611c938161069a565b610810819560ff81546001600160a01b0381168552818160a01c16151560208601526001600160401b038160a81c16604086015260e81c1615156060840152611ce26040518096819301611be4565b0384610726565b634e487b7160e01b600052601160045260246000fd5b908160051b9180830460201490151715611d1557565b611ce9565b91908203918211611d1557565b611d3382607f92613331565b9116906801fffffffffffffffe6001600160401b0383169260011b169180830460021490151715611d15576003911c16600481101561062f5790565b611d77613375565b805182518103611f725760005b818110611d975750509061060d916133b8565b611da18184611b96565b516020810190815151611db48488611b96565b519283518203611f725790916000925b808410611dd8575050505050600101611d84565b91949398611dea848b98939598611b96565b515198611df8888851611b96565b519980611f29575b5060a08a01988b6020611e168b8d515193611b96565b5101515103611ee85760005b8a5151811015611ed357611e5e611e55611e4b8f6020611e438f8793611b96565b510151611b96565b5163ffffffff1690565b63ffffffff1690565b8b81611e6f575b5050600101611e22565b611e556040611e8285611e8e9451611b96565b51015163ffffffff1690565b90818110611e9d57508b611e65565b8d51516040516348e617b360e01b81526004810191909152602481019390935260448301919091526064820152608490fd5b0390fd5b50985098509893949095600101929091611dc4565b611f258b51611f03606082519201516001600160401b031690565b6370a193fd60e01b6000526004919091526001600160401b0316602452604490565b6000fd5b60808b0151811015611e0057611f25908b611f4b88516001600160401b031690565b905151633a98d46360e11b6000526001600160401b03909116600452602452604452606490565b6320f8fd5960e21b60005260046000fd5b60405190611f90826106ba565b60006020838281520152565b60405190611fab602083610726565b600080835282815b828110611fbf57505050565b602090611fca611f83565b82828501015201611fb3565b805182526001600160401b036020820151166020830152608061201d61200b604084015160a0604087015260a086019061033e565b6060840151858203606087015261033e565b9101519160808183039101526020808351928381520192019060005b8181106120465750505090565b825180516001600160a01b031685526020908101518186015260409094019390920191600101612039565b90602061043a928181520190611fd6565b6040513d6000823e3d90fd5b3d156120b9573d9061209f82610814565b916120ad6040519384610726565b82523d6000602084013e565b606090565b90602061043a92818152019061033e565b81601f820112156101875780516120e581610814565b926120f36040519485610726565b818452602082840101116101875761043a916020808501910161031b565b909160608284031261018757815161212881611531565b9260208301516001600160401b0381116101875760409161214a9185016120cf565b92015190565b9293606092959461ffff6121746001600160a01b0394608088526080880190611fd6565b97166020860152604085015216910152565b9290939130330361245257612199611f9c565b9460a0850151805161240b575b50505050508051916121c4602084519401516001600160401b031690565b9060208301519160408401926121f18451926121de610756565b9788526001600160401b03166020880152565b6040860152606085015260808401526001600160a01b0361221a6005546001600160a01b031690565b168061238e575b5051511580612382575b801561236c575b8015612343575b61233f576122d7918161227c6122706122636105a0602060009751016001600160401b0390511690565b546001600160a01b031690565b6001600160a01b031690565b9083612297606060808401519301516001600160a01b031690565b604051633cf9798360e01b815296879586948593917f00000000000000000000000000000000000000000000000000000000000000009060048601612150565b03925af190811561233a57600090600092612313575b50156122f65750565b6040516302a35ba360e21b8152908190611ecf90600483016120be565b905061233291503d806000833e61232a8183610726565b810190612111565b5090386122ed565b612082565b5050565b5061236761236361235e60608401516001600160a01b031690565b6135df565b1590565b612239565b5060608101516001600160a01b03163b15612232565b5060808101511561222b565b803b1561018757600060405180926308d450a160e01b82528183816123b68a60048301612071565b03925af190816123f0575b506123ea57611ecf6123d161208e565b6040516309c2532560e01b8152918291600483016120be565b38612221565b806123ff600061240593610726565b8061017c565b386123c1565b859650602061244796015161242a60608901516001600160a01b031690565b9061244160208a51016001600160401b0390511690565b926134c6565b9038808080806121a6565b6306e34e6560e31b60005260046000fd5b90604051612470816106f0565b606060ff600183958054855201548181166020850152818160081c16604085015260101c161515910152565b906020825491828152019160005260206000209060005b8181106124c05750505090565b82546001600160a01b03168452602090930192600192830192016124b3565b9061060d611c7b926040519384809261249c565b6124fb61362d565b60005b815181101561233f576125118183611b96565b519061252760208301516001600160401b031690565b6001600160401b0381169081156127ac5761254f61227061227086516001600160a01b031690565b1561105a57612571816001600160401b03166000526008602052604060002090565b60808501519060018101926125868454611baa565b61273e576125f97ff4c1390c70e5c0f491ae1ccbc06f9117cbbadf2767b247b3bc203280f24c0fb9916125df84750100000000000000000000000000000000000000000067ffffffffffffffff60a81b19825416179055565b6040516001600160401b0390911681529081906020820190565b0390a15b81518015908115612728575b5061105a576127096126d4606060019861264761271f967fbd1ab25a0ff0a36a588597ba1af11e30f3f210de8b9e818cc9bbc457c94c8d8c986136cf565b61269d6126576040830151151590565b86547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178655565b6126cd6126b182516001600160a01b031690565b86906001600160a01b03166001600160a01b0319825416179055565b0151151590565b82547fffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690151560e81b60ff60e81b16178255565b61271284615c76565b50604051918291826137a0565b0390a2016124fe565b90506020830120612737613652565b1438612609565b60016001600160401b0361275d84546001600160401b039060a81c1690565b1614158061278d575b61277057506125fd565b632105803760e11b6000526001600160401b031660045260246000fd5b5061279784611c67565b60208151910120835160208501201415612766565b63c656089560e01b60005260046000fd5b35906001600160e01b038216820361018757565b81601f82011215610187578035906127e882610793565b926127f66040519485610726565b82845260208085019360061b8301019181831161018757602001925b828410612820575050505090565b604084830312610187576020604091825161283a816106ba565b8635612845816105f1565b81526128528388016127bd565b83820152815201930192612812565b9190604083820312610187576040519061287a826106ba565b819380356001600160401b03811161018757810182601f820112156101875780356128a481610793565b916128b26040519384610726565b81835260208084019260061b8201019085821161018757602001915b8183106129005750505083526020810135916001600160401b038311610187576020926128fb92016127d1565b910152565b604083870312610187576020604091825161291a816106ba565b853561292581610881565b81526129328387016127bd565b838201528152019201916128ce565b81601f820112156101875780359061295882610793565b926129666040519485610726565b82845260208085019360051b830101918183116101875760208101935b83851061299257505050505090565b84356001600160401b03811161018757820160a0818503601f19011261018757604051916129bf8361069a565b60208201356129cd816105f1565b83526040820135926001600160401b0384116101875760a0836129f7886020809881980101610866565b858401526060810135612a09816105f1565b60408401526080810135612a1c816105f1565b606084015201356080820152815201940193612983565b81601f8201121561018757803590612a4a82610793565b92612a586040519485610726565b82845260208085019360061b8301019181831161018757602001925b828410612a82575050505090565b6040848303126101875760206040918251612a9c816106ba565b863581528287013583820152815201930192612a74565b602081830312610187578035906001600160401b038211610187570160808183031261018757612ae1610765565b9181356001600160401b0381116101875781612afe918401612861565b835260208201356001600160401b0381116101875781612b1f918401612941565b602084015260408201356001600160401b0381116101875781612b43918401612941565b604084015260608201356001600160401b03811161018757612b659201612a33565b606082015290565b9080602083519182815201916020808360051b8301019401926000915b838310612b9957505050505090565b9091929394602080600192601f198582030186528851906001600160401b038251168152608080612bd78585015160a08786015260a085019061033e565b936001600160401b0360408201511660408501526001600160401b036060820151166060850152015191015297019301930191939290612b8a565b916001600160a01b03612c3392168352606060208401526060830190612b6d565b9060408183039101526020808351928381520192019060005b818110612c595750505090565b8251805185526020908101518186015260409094019390920191600101612c4c565b6084019081608411611d1557565b60a001908160a011611d1557565b91908201809211611d1557565b906020808351928381520192019060005b818110612cc25750505090565b825180516001600160401b031685526020908101516001600160e01b03168186015260409094019390920191600101612cb5565b9190604081019083519160408252825180915260206060830193019060005b818110612d3657505050602061043a93940151906020818403910152612ca4565b825180516001600160a01b031686526020908101516001600160e01b03168187015260409095019490920191600101612d15565b90602061043a928181520190612cf6565b91612da490612d9661043a9593606086526060860190612b6d565b908482036020860152612b6d565b916040818403910152612cf6565b9197939796929695909495612dc981870187612ab3565b95602087019788518051612fac575b5087518051511590811591612f9d575b50612eb8575b60005b89518051821015612e185790612e12612e0c82600194611b96565b51613850565b01612df1565b50509193959799989092949698600099604081019a5b8b518051821015612e555790612e4f612e4982600194611b96565b51613b24565b01612e2e565b5050907fb967c9b9e1b7af9a61ca71ff00e9f5b89ec6f2e268de8dacf12f0de8e51f3e47612eaa939261060d9c612ea0612eb298999a9b9c9d9f519151925160405193849384612d7b565b0390a13691610bd7565b943691610bd7565b93613fa3565b612ecd602086015b356001600160401b031690565b600b546001600160401b0382811691161015612f7557612f03906001600160401b03166001600160401b0319600b541617600b55565b612f1b6122706122706004546001600160a01b031690565b885190803b1561018757604051633937306f60e01b8152916000918391829084908290612f4b9060048301612d6a565b03925af1801561233a57612f60575b50612dee565b806123ff6000612f6f93610726565b38612f5a565b50612f8889515160408a01515190612c97565b612dee57632261116760e01b60005260046000fd5b60209150015151151538612de8565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169060608a0151823b1561018757604051633854844f60e11b815292600092849283918291613008913060048501612c12565b03915afa801561233a5715612dd857806123ff600061302693610726565b38612dd8565b6040519061303b602083610726565b6000808352366020840137565b61305061362d565b60005b815181101561233f576130668183611b96565b51906040820160ff613079825160ff1690565b161561331b57602083015160ff169261309f8460ff166000526002602052604060002090565b91600183019182546130ba6130b48260ff1690565b60ff1690565b6132e057506130e76130cf6060830151151590565b845462ff0000191690151560101b62ff000016178455565b60a08101918251610100815111613288578051156132ca576003860161311561310f826124df565b8a614fa6565b60608401516131a5575b947fab8b1b57514019638d7b5ce9c638fe71366fe8e2be1c40a7a80f1733d0e9f5479460029461318161317161319f9a9661316a8760019f9c6131656131979a8f615114565b6141e4565b5160ff1690565b845460ff191660ff821617909455565b519081855551906040519586950190888661426a565b0390a1615196565b01613053565b979460028793959701966131c16131bb896124df565b88614fa6565b6080850151946101008651116132b45785516131e96130b46131e48a5160ff1690565b6141d0565b101561329e578551845111613288576131816131717fab8b1b57514019638d7b5ce9c638fe71366fe8e2be1c40a7a80f1733d0e9f5479861316a8760019f61316561319f9f9a8f61327060029f61326a6131979f8f90613165849261324f845160ff1690565b908054909161ff001990911660089190911b61ff0016179055565b8261503a565b505050979c9f50975050969a5050509450945061311f565b631b3fab5160e11b600052600160045260246000fd5b631b3fab5160e11b600052600360045260246000fd5b631b3fab5160e11b600052600260045260246000fd5b631b3fab5160e11b600052600560045260246000fd5b60101c60ff166132fb6132f66060840151151590565b151590565b901515146130e7576321fd80df60e21b60005260ff861660045260246000fd5b631b3fab5160e11b600090815260045260246000fd5b906001600160401b03613371921660005260096020526701ffffffffffffff60406000209160071c166001600160401b0316600052602052604060002090565b5490565b7f00000000000000000000000000000000000000000000000000000000000000004681036133a05750565b630f01ce8560e01b6000526004524660245260446000fd5b91909180511561345a5782511592602091604051926133d78185610726565b60008452601f19810160005b8181106134365750505060005b815181101561342e578061341761340960019385611b96565b51881561341d5786906143a9565b016133f0565b6134278387611b96565b51906143a9565b505050509050565b8290604051613444816106ba565b60008152606083820152828289010152016133e3565b63c2e5347d60e01b60005260046000fd5b9190811015611b915760051b0190565b3561043a8161089d565b9190811015611b915760051b81013590601e19813603018212156101875701908135916001600160401b038311610187576020018236038113610187579190565b909294919397968151966134d988610793565b976134e7604051998a610726565b8089526134f6601f1991610793565b0160005b8181106135c857505060005b83518110156135bb578061354d8c8a8a8a613547613540878d613539828f8f9d8f9e60019f81613569575b505050611b96565b5197613485565b369161082f565b93614aa0565b613557828c611b96565b52613562818b611b96565b5001613506565b63ffffffff61358161357c85858561346b565b61347b565b1615613531576135b1926135989261357c9261346b565b60406135a48585611b96565b51019063ffffffff169052565b8f8f908391613531565b5096985050505050505050565b6020906135d3611f83565b82828d010152016134fa565b6135f06385572ffb60e01b82614e03565b908161360a575b81613600575090565b61043a9150614dd5565b905061361581614d5a565b15906135f7565b6135f063aff2afbf60e01b82614e03565b6001600160a01b0360015416330361364157565b6315ae3a6f60e11b60005260046000fd5b6040516020810190600082526020815261366d604082610726565b51902090565b81811061367e575050565b60008155600101613673565b9190601f811161369957505050565b61060d926000526020600020906020601f840160051c830193106136c5575b601f0160051c0190613673565b90915081906136b8565b91909182516001600160401b0381116106b5576136f6816136f08454611baa565b8461368a565b6020601f821160011461373757819061372893949560009261372c575b50508160011b916000199060031b1c19161790565b9055565b015190503880613713565b601f1982169061374c84600052602060002090565b9160005b8181106137885750958360019596971061376f575b505050811b019055565b015160001960f88460031b161c19169055388080613765565b9192602060018192868b015181550194019201613750565b90600160c061043a936020815260ff84546001600160a01b0381166020840152818160a01c16151560408401526001600160401b038160a81c16606084015260e81c161515608082015260a080820152019101611be4565b90816020910312610187575161043a81611531565b909161382461043a9360408452604084019061033e565b916020818403910152611be4565b6001600160401b036001911601906001600160401b038211611d1557565b8051604051632cbc26bb60e01b815267ffffffffffffffff60801b608083901b1660048201526001600160401b0390911691906020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561233a57600091613af5575b50613ad7576138d282614e33565b805460ff60e882901c161515600114613aac576020830180516020815191012090600184019161390183611c67565b6020815191012003613a8f57505060408301516001600160401b039081169160a81c168114801590613a67575b613a2657506080820151918215613a155761396f83613960866001600160401b0316600052600a602052604060002090565b90600052602052604060002090565b546139f2576139ef929161399861399360606139d19401516001600160401b031690565b613832565b67ffffffffffffffff60a81b197cffffffffffffffff00000000000000000000000000000000000000000083549260a81b169116179055565b61396042936001600160401b0316600052600a602052604060002090565b55565b6332cf0cbf60e01b6000526001600160401b038416600452602483905260446000fd5b63504570e360e01b60005260046000fd5b83611f2591613a3f60608601516001600160401b031690565b636af0786b60e11b6000526001600160401b0392831660045290821660245216604452606490565b50613a7f61055c60608501516001600160401b031690565b6001600160401b0382161161392e565b51611ecf60405192839263b80d8fa960e01b84526004840161380d565b60808301516348e2b93360e11b6000526001600160401b038516600452602452600160445260646000fd5b637edeb53960e11b6000526001600160401b03821660045260246000fd5b613b17915060203d602011613b1d575b613b0f8183610726565b8101906137f8565b386138c4565b503d613b05565b8051604051632cbc26bb60e01b815267ffffffffffffffff60801b608083901b1660048201526001600160401b0390911691906020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561233a57600091613bff575b50613ad757613ba682614e33565b805460ff60e882901c1615613bd1576020830180516020815191012090600184019161390183611c67565b60808301516348e2b93360e11b60009081526001600160401b03861660045260249190915260445260646000fd5b613c18915060203d602011613b1d57613b0f8183610726565b38613b98565b6003111561062f57565b600382101561062f5752565b9061060d604051613c44816106ba565b602060ff829554818116845260081c169101613c28565b8054821015611b915760005260206000200190600090565b60ff60019116019060ff8211611d1557565b60ff601b9116019060ff8211611d1557565b90606092604091835260208301370190565b6001600052600260205293613cdd7fe90b7bceb6e7df5418fb78d8ee546e97c83a08bbccc01a0644d599ccd2a7c2e0612463565b93853594613cea85612c7b565b6060820190613cf98251151590565b613f75575b803603613f5d57508151878103613f445750613d18613375565b60016000526003602052613d67613d627fa15bc60c955c405d20d9149c709e2460f1c2d9a497496a7f46004d1772c3054c5b336001600160a01b0316600052602052604060002090565b613c34565b60026020820151613d7781613c1e565b613d8081613c1e565b149081613edc575b5015613eb0575b51613de7575b50505050507f198d6990ef96613a9026203077e422916918b03ff47f0be6bee7b02d8e139ef090613dcb612ec060019460200190565b604080519283526001600160401b0391909116602083015290a2565b613e086130b4613e03602085969799989a955194015160ff1690565b613c73565b03613e9f578151835103613e8e57613e866000613dcb94612ec094613e527f198d6990ef96613a9026203077e422916918b03ff47f0be6bee7b02d8e139ef09960019b369161082f565b60208151910120604051613e7d81613e6f89602083019586613c97565b03601f198101835282610726565b5190208a614e70565b948394613d95565b63a75d88af60e01b60005260046000fd5b6371253a2560e01b60005260046000fd5b72c11c11c11c11c11c11c11c11c11c11c11c11c1330315613d8f57631b41e11d60e31b60005260046000fd5b60016000526002602052613f3c915061227090613f2990613f2360037fe90b7bceb6e7df5418fb78d8ee546e97c83a08bbccc01a0644d599ccd2a7c2e05b01915160ff1690565b90613c5b565b90546001600160a01b039160031b1c1690565b331438613d88565b6324f7d61360e21b600052600452602487905260446000fd5b638e1192e160e01b6000526004523660245260446000fd5b613f9e90613f98613f8e613f898751611cff565b612c89565b613f988851611cff565b90612c97565b613cfe565b60008052600260205294909390929091613fdc7fac33ff75c19e70fe83507db0d683fd3465c996598dc972688b7ace676c89077b612463565b94863595613fe983612c7b565b6060820190613ff88251151590565b6141ad575b803603613f5d575081518881036141945750614017613375565b60008052600360205261404c613d627f3617319a054d772f909f7c479a2cebe5066e836a939412e32403c99029b92eff613d4a565b6002602082015161405c81613c1e565b61406581613c1e565b14908161414b575b501561411f575b516140b1575b5050505050507f198d6990ef96613a9026203077e422916918b03ff47f0be6bee7b02d8e139ef090613dcb612ec060009460200190565b6140cd6130b4613e03602087989a999b96975194015160ff1690565b03613e9f578351865103613e8e576000967f198d6990ef96613a9026203077e422916918b03ff47f0be6bee7b02d8e139ef096613dcb95613e5261411694612ec097369161082f565b9483943861407a565b72c11c11c11c11c11c11c11c11c11c11c11c11c133031561407457631b41e11d60e31b60005260046000fd5b60008052600260205261418c915061227090613f2990613f2360037fac33ff75c19e70fe83507db0d683fd3465c996598dc972688b7ace676c89077b613f1a565b33143861406d565b6324f7d61360e21b600052600452602488905260446000fd5b6141cb90613f986141c1613f898951611cff565b613f988a51611cff565b613ffd565b60ff166003029060ff8216918203611d1557565b8151916001600160401b0383116106b5576801000000000000000083116106b557602090825484845580851061424d575b500190600052602060002060005b8381106142305750505050565b60019060206001600160a01b038551169401938184015501614223565b614264908460005285846000209182019101613673565b38614215565b95949392909160ff61428f93168752602087015260a0604087015260a086019061249c565b84810360608601526020808351928381520192019060005b8181106142c25750505090608061060d9294019060ff169052565b82516001600160a01b03168452602093840193909201916001016142a7565b600654811015611b915760066000527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f015490565b6001600160401b0361043a949381606094168352166020820152816040820152019061033e565b60409061043a93928152816020820152019061033e565b9291906001600160401b0390816064951660045216602452600481101561062f57604452565b9493926143936060936143a49388526020880190610634565b60806040870152608086019061033e565b930152565b906143bb82516001600160401b031690565b8151604051632cbc26bb60e01b815267ffffffffffffffff60801b608084901b1660048201529015159391906001600160401b038216906020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561233a5760009161497e575b5061491f5760208301918251519485156148ef57604085019485515187036148de5761445f9083615205565b95909760005b8881106144785750505050505050505050565b5a61448d614487838a51611b96565b51615741565b8051606001516144a6906001600160401b031688611d27565b6144af81610625565b8015908d82831593846148cb575b15614888576060881561480b57506144e460206144da898d611b96565b5101519242611d1a565b6004546144f99060a01c63ffffffff16611e55565b1080156147f8575b156147da57614510878b611b96565b51516147c4575b84516080015161452f906001600160401b031661055c565b61470c575b50614540868951611b96565b5160a0850151518151036146d057936145a59695938c938f966145858e958c9261457f61457960608951016001600160401b0390511690565b89615814565b86615a12565b9a90809661459f60608851016001600160401b0390511690565b90615899565b61467e575b50506145b582610625565b60028203614636575b60019661462c7f05665fe9ad095383d018353f4cbcba77e84db27dd215081bbf7cdf9ae6fbe48b936001600160401b0393519261461d6146148b61460c60608801516001600160401b031690565b96519b611b96565b51985a90611d1a565b9160405195869516988561437a565b0390a45b01614465565b9150919394925061464682610625565b6003820361465a578b929493918a916145be565b51606001516349362d1f60e11b600052611f2591906001600160401b031689614354565b61468784610625565b600384036145aa57909294955061469f919350610625565b6146af578b92918a9138806145aa565b5151604051632b11b8d960e01b8152908190611ecf9087906004840161433d565b611f258b6146ea60608851016001600160401b0390511690565b631cfe6d8b60e01b6000526001600160401b0391821660045216602452604490565b61471583610625565b614720575b38614534565b8351608001516001600160401b0316602080860151918c61475560405194859384936370701e5760e11b855260048501614316565b038160006001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165af190811561233a576000916147a6575b5061471a575050505050600190614630565b6147be915060203d8111613b1d57613b0f8183610726565b38614794565b6147ce878b611b96565b51516080860152614517565b6354e7e43160e11b6000526001600160401b038b1660045260246000fd5b5061480283610625565b60038314614501565b91508361481784610625565b1561451757506001959450614880925061485e91507f3ef2a99c550a751d4b0b261268f05a803dfb049ab43616a1ffb388f61fe651209351016001600160401b0390511690565b604080516001600160401b03808c168252909216602083015290918291820190565b0390a1614630565b50505050600192915061488061485e60607f3b575419319662b2a6f5e2467d84521517a3382b908eb3d557bb3fdb0c50e23c9351016001600160401b0390511690565b506148d583610625565b600383146144bd565b6357e0e08360e01b60005260046000fd5b611f2561490386516001600160401b031690565b63676cf24b60e11b6000526001600160401b0316600452602490565b5092915050614961576040516001600160401b039190911681527faab522ed53d887e56ed53dd37398a01aeef6a58e0fa77c2173beb9512d89493390602090a1565b637edeb53960e11b6000526001600160401b031660045260246000fd5b614997915060203d602011613b1d57613b0f8183610726565b38614433565b519061060d82610881565b90816020910312610187575161043a81610881565b9061043a916020815260e0614a5b614a466149e68551610100602087015261012086019061033e565b60208601516001600160401b0316604086015260408601516001600160a01b0316606086015260608601516080860152614a30608087015160a08701906001600160a01b03169052565b60a0860151858203601f190160c087015261033e565b60c0850151848203601f19018486015261033e565b92015190610100601f198285030191015261033e565b6040906001600160a01b0361043a9493168152816020820152019061033e565b90816020910312610187575190565b91939293614aac611f83565b5060208301516001600160a01b031660405163bbe4f6db60e01b81526001600160a01b038216600482015290959092602084806024810103816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa93841561233a57600094614d29575b506001600160a01b0384169586158015614d17575b614cf957614bde614c0792613e6f92614b62614b5b611e5560408c015163ffffffff1690565b8c89615b2b565b9690996080810151614b906060835193015193614b7d610774565b9687526001600160401b03166020870152565b6001600160a01b038a16604086015260608501526001600160a01b038d16608085015260a084015260c083015260e0820152604051633907753760e01b6020820152928391602483016149bd565b82857f000000000000000000000000000000000000000000000000000000000000000092615bb9565b94909115614cdd5750805160208103614cc4575090614c30826020808a95518301019101614a91565b956001600160a01b03841603614c68575b5050505050614c60614c51610784565b6001600160a01b039093168352565b602082015290565b614c7b93614c7591611d1a565b91615b2b565b50908082108015614cb1575b614c9357808481614c41565b63a966e21f60e01b6000908152600493909352602452604452606490fd5b5082614cbd8284611d1a565b1415614c87565b631e3be00960e21b600052602060045260245260446000fd5b611ecf604051928392634ff17cad60e11b845260048401614a71565b63ae9b4ce960e01b6000526001600160a01b03851660045260246000fd5b50614d246123638661361c565b614b35565b614d4c91945060203d602011614d53575b614d448183610726565b8101906149a8565b9238614b20565b503d614d3a565b60405160208101916301ffc9a760e01b835263ffffffff60e01b602483015260248252614d88604483610726565b6179185a10614dc4576020926000925191617530fa6000513d82614db8575b5081614db1575090565b9050151590565b60201115915038614da7565b63753fa58960e11b60005260046000fd5b60405160208101916301ffc9a760e01b83526301ffc9a760e01b602483015260248252614d88604483610726565b6040519060208201926301ffc9a760e01b845263ffffffff60e01b16602483015260248252614d88604483610726565b6001600160401b031680600052600860205260406000209060ff825460a01c1615614e5c575090565b63ed053c5960e01b60005260045260246000fd5b919390926000948051946000965b868810614e8f575050505050505050565b6020881015611b915760206000614ea7878b1a613c85565b614eb18b87611b96565b5190614ee8614ec08d8a611b96565b5160405193849389859094939260ff6060936080840197845216602083015260408201520152565b838052039060015afa1561233a57614f2e613d62600051614f168960ff166000526003602052604060002090565b906001600160a01b0316600052602052604060002090565b9060016020830151614f3f81613c1e565b614f4881613c1e565b03614f9557614f65614f5b835160ff1690565b60ff600191161b90565b8116614f8457614f7b614f5b6001935160ff1690565b17970196614e7e565b633d9ef1f160e21b60005260046000fd5b636518c33d60e11b60005260046000fd5b91909160005b8351811015614fff5760019060ff831660005260036020526000614ff8604082206001600160a01b03614fdf858a611b96565b51166001600160a01b0316600052602052604060002090565b5501614fac565b50509050565b8151815460ff191660ff919091161781559060200151600381101561062f57815461ff00191660089190911b61ff0016179055565b919060005b8151811015614fff576150626150558284611b96565b516001600160a01b031690565b9061508b61508183614f168860ff166000526003602052604060002090565b5460081c60ff1690565b61509481613c1e565b6150ff576001600160a01b038216156150ee576150e86001926150e36150b8610784565b60ff85168152916150cc8660208501613c28565b614f168960ff166000526003602052604060002090565b615005565b0161503f565b63d6c62c9b60e01b60005260046000fd5b631b3fab5160e11b6000526004805260246000fd5b919060005b8151811015614fff5761512f6150558284611b96565b9061514e61508183614f168860ff166000526003602052604060002090565b61515781613c1e565b6150ff576001600160a01b038216156150ee576151906001926150e361517b610784565b60ff85168152916150cc600260208501613c28565b01615119565b60ff1680600052600260205260ff60016040600020015460101c169080156000146151e45750156151d3576001600160401b0319600b5416600b55565b6317bd8dd160e11b60005260046000fd5b6001146151ee5750565b6151f457565b6307b8c74d60e51b60005260046000fd5b6020820180515193929061521885611b3c565b947f000000000000000000000000000000000000000000000000000000000000000061524860016117d487614e33565b602081519101206040516152a881613e6f6020820194868b876001600160401b036060929594938160808401977f2425b0b9f9054c76ff151b0a175b18f37a4a4e82013a72e9f15c9caa095ed21f85521660208401521660408201520152565b519020926001600160401b0360009216915b83811061530557505050505080608060606152dc930151910151908584615e13565b9081156152e857509190565b633ee8bd3f60e11b6000526001600160401b031660045260246000fd5b615310818351611b96565b518051604001516001600160401b031684810361539257508051602001516001600160401b03166001600160401b0389166001600160401b0382160361536f57509061535e86600193615d0b565b615368828c611b96565b52016152ba565b636c95f1eb60e01b6000526001600160401b03808a166004521660245260446000fd5b631c21951160e11b6000526001600160401b031660045260246000fd5b91908260a0910312610187576040516153c78161069a565b60808082948051845260208101516153de816105f1565b602085015260408101516153f1816105f1565b60408501526060810151615404816105f1565b6060850152015191610810836105f1565b519061060d8261089d565b81601f820112156101875780519061543782610793565b926154456040519485610726565b82845260208085019360051b830101918183116101875760208101935b83851061547157505050505090565b84516001600160401b03811161018757820160a0818503601f190112610187576040519161549e8361069a565b60208201516001600160401b038111610187578560206154c0928501016120cf565b835260408201516154d081610881565b60208401526154e160608301615415565b60408401526080820151926001600160401b0384116101875760a08361550e8860208098819801016120cf565b606084015201516080820152815201940193615462565b602081830312610187578051906001600160401b03821161018757016101408183031261018757615554610747565b9161555f81836153af565b835260a08201516001600160401b03811161018757816155809184016120cf565b602084015260c08201516001600160401b03811161018757816155a49184016120cf565b60408401526155b560e0830161499d565b606084015261010082015160808401526101208201516001600160401b038111610187576155e39201615420565b60a082015290565b9080602083519182815201916020808360051b8301019401926000915b83831061561757505050505090565b9091929394602080600192601f1985820301865288519060808061567a615647855160a0865260a086019061033e565b6001600160a01b0387870151168786015263ffffffff60408701511660408601526060860151858203606087015261033e565b93015191015297019301930191939290615608565b61043a916001600160401b036080835180518452826020820151166020850152826040820151166040850152826060820151166060850152015116608082015260a06157006156ee60208501516101408486015261014085019061033e565b604085015184820360c086015261033e565b60608401516001600160a01b031660e08401529260808101516101008401520151906101208184039101526155eb565b90602061043a92818152019061568f565b60006157b981926040516157548161070b565b61575c611ace565b81526060602082015260606040820152836060820152836080820152606060a08201525061579c612270612270600b546001600160a01b039060401c1690565b90604051948580948193634546c6e560e01b835260048301615730565b03925af1600091816157ef575b5061043a57611ecf6157d661208e565b60405163828ebdfb60e01b8152918291600483016120be565b61580d9192503d806000833e6158058183610726565b810190615525565b90386157c6565b607f8216906801fffffffffffffffe6001600160401b0383169260011b169180830460021490151715611d15576139ef916001600160401b036158578584613331565b921660005260096020526701ffffffffffffff60406000209460071c169160036001831b921b19161792906001600160401b0316600052602052604060002090565b9091607f83166801fffffffffffffffe6001600160401b0382169160011b169080820460021490151715611d15576158d18484613331565b600483101561062f576001600160401b036139ef9416600052600960205260036701ffffffffffffff60406000209660071c1693831b921b19161792906001600160401b0316600052602052604060002090565b906159389060608352606083019061568f565b8181036020830152825180825260208201916020808360051b8301019501926000915b8383106159a857505050505060408183039101526020808351928381520192019060005b81811061598c5750505090565b825163ffffffff1684526020938401939092019160010161597f565b90919293956020806159c6600193601f198682030187528a5161033e565b9801930193019193929061595b565b80516020909101516001600160e01b03198116929190600482106159f7575050565b6001600160e01b031960049290920360031b82901b16169150565b90303b1561018757600091615a3b6040519485938493630304c3e160e51b855260048501615925565b038183305af19081615b16575b50615b0b57615a5561208e565b9072c11c11c11c11c11c11c11c11c11c11c11c11c13314615a77575b60039190565b615a90615a83836159d5565b6001600160e01b03191690565b6337c3be2960e01b148015615af0575b8015615ad5575b15615a7157611f25615ab8836159d5565b632882569d60e01b6000526001600160e01b031916600452602490565b50615ae2615a83836159d5565b63753fa58960e11b14615aa7565b50615afd615a83836159d5565b632be8ca8b60e21b14615aa0565b60029061043a611af9565b806123ff6000615b2593610726565b38615a48565b6040516370a0823160e01b60208201526001600160a01b039091166024820152919291615b8890615b5f8160448101613e6f565b84837f000000000000000000000000000000000000000000000000000000000000000092615bb9565b92909115614cdd5750805160208103614cc4575090615bb38260208061043a95518301019101614a91565b93611d1a565b939193615bc66084610814565b94615bd46040519687610726565b60848652615be26084610814565b602087019590601f1901368737833b15615c65575a90808210615c54578291038060061c90031115615c43576000918291825a9560208451940192f1905a9003923d9060848211615c3a575b6000908287523e929190565b60849150615c2e565b6337c3be2960e01b60005260046000fd5b632be8ca8b60e21b60005260046000fd5b63030ed58f60e21b60005260046000fd5b80600052600760205260406000205415600014615cf457600654680100000000000000008110156106b557600181016006556000600654821015611b9157600690527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f01819055600654906000526007602052604060002055600190565b50600090565b90602061043a9281815201906155eb565b61366d8151805190615d9f615d2a60608601516001600160a01b031690565b613e6f615d4160608501516001600160401b031690565b93615d5a6080808a01519201516001600160401b031690565b90604051958694602086019889936001600160401b036080946001600160a01b0382959998949960a089019a8952166020880152166040860152606085015216910152565b519020613e6f6020840151602081519101209360a0604082015160208151910120910151604051615dd881613e6f602082019485615cfa565b51902090604051958694602086019889919260a093969594919660c08401976000855260208501526040840152606083015260808201520152565b926001600160401b0392615e2692615e45565b9116600052600a60205260406000209060005260205260406000205490565b8051928251908415615fa15761010185111580615f95575b15615ec457818501946000198601956101008711615ec4578615615f8557615e8487611b3c565b9660009586978795885b848110615ee9575050505050600119018095149384615edf575b505082615ed5575b505015615ec457615ec091611b96565b5190565b6309bde33960e01b60005260046000fd5b1490503880615eb0565b1492503880615ea8565b6001811b82811603615f7757868a1015615f6257615f0b60018b019a85611b96565b51905b8c888c1015615f4e5750615f2660018c019b86611b96565b515b818d11615ec457615f47828f92615f4190600196615fb2565b92611b96565b5201615e8e565b60018d019c615f5c91611b96565b51615f28565b615f7060018c019b8d611b96565b5190615f0e565b615f70600189019884611b96565b505050509050615ec09150611b84565b50610101821115615e5d565b630469ac9960e21b60005260046000fd5b81811015615fc4579061043a91615fc9565b61043a915b9060405190602082019260018452604083015260608201526060815261366d60808261072656fea164736f6c634300081a000abd1ab25a0ff0a36a588597ba1af11e30f3f210de8b9e818cc9bbc457c94c8d8c",
}

var OffRampWithMessageTransformerABI = OffRampWithMessageTransformerMetaData.ABI

var OffRampWithMessageTransformerBin = OffRampWithMessageTransformerMetaData.Bin

func DeployOffRampWithMessageTransformer(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig OffRampStaticConfig, dynamicConfig OffRampDynamicConfig, sourceChainConfigs []OffRampSourceChainConfigArgs, messageTransformerAddr common.Address) (common.Address, *types.Transaction, *OffRampWithMessageTransformer, error) {
	parsed, err := OffRampWithMessageTransformerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OffRampWithMessageTransformerBin), backend, staticConfig, dynamicConfig, sourceChainConfigs, messageTransformerAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OffRampWithMessageTransformer{address: address, abi: *parsed, OffRampWithMessageTransformerCaller: OffRampWithMessageTransformerCaller{contract: contract}, OffRampWithMessageTransformerTransactor: OffRampWithMessageTransformerTransactor{contract: contract}, OffRampWithMessageTransformerFilterer: OffRampWithMessageTransformerFilterer{contract: contract}}, nil
}

type OffRampWithMessageTransformer struct {
	address common.Address
	abi     abi.ABI
	OffRampWithMessageTransformerCaller
	OffRampWithMessageTransformerTransactor
	OffRampWithMessageTransformerFilterer
}

type OffRampWithMessageTransformerCaller struct {
	contract *bind.BoundContract
}

type OffRampWithMessageTransformerTransactor struct {
	contract *bind.BoundContract
}

type OffRampWithMessageTransformerFilterer struct {
	contract *bind.BoundContract
}

type OffRampWithMessageTransformerSession struct {
	Contract     *OffRampWithMessageTransformer
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type OffRampWithMessageTransformerCallerSession struct {
	Contract *OffRampWithMessageTransformerCaller
	CallOpts bind.CallOpts
}

type OffRampWithMessageTransformerTransactorSession struct {
	Contract     *OffRampWithMessageTransformerTransactor
	TransactOpts bind.TransactOpts
}

type OffRampWithMessageTransformerRaw struct {
	Contract *OffRampWithMessageTransformer
}

type OffRampWithMessageTransformerCallerRaw struct {
	Contract *OffRampWithMessageTransformerCaller
}

type OffRampWithMessageTransformerTransactorRaw struct {
	Contract *OffRampWithMessageTransformerTransactor
}

func NewOffRampWithMessageTransformer(address common.Address, backend bind.ContractBackend) (*OffRampWithMessageTransformer, error) {
	abi, err := abi.JSON(strings.NewReader(OffRampWithMessageTransformerABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindOffRampWithMessageTransformer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OffRampWithMessageTransformer{address: address, abi: abi, OffRampWithMessageTransformerCaller: OffRampWithMessageTransformerCaller{contract: contract}, OffRampWithMessageTransformerTransactor: OffRampWithMessageTransformerTransactor{contract: contract}, OffRampWithMessageTransformerFilterer: OffRampWithMessageTransformerFilterer{contract: contract}}, nil
}

func NewOffRampWithMessageTransformerCaller(address common.Address, caller bind.ContractCaller) (*OffRampWithMessageTransformerCaller, error) {
	contract, err := bindOffRampWithMessageTransformer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OffRampWithMessageTransformerCaller{contract: contract}, nil
}

func NewOffRampWithMessageTransformerTransactor(address common.Address, transactor bind.ContractTransactor) (*OffRampWithMessageTransformerTransactor, error) {
	contract, err := bindOffRampWithMessageTransformer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OffRampWithMessageTransformerTransactor{contract: contract}, nil
}

func NewOffRampWithMessageTransformerFilterer(address common.Address, filterer bind.ContractFilterer) (*OffRampWithMessageTransformerFilterer, error) {
	contract, err := bindOffRampWithMessageTransformer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OffRampWithMessageTransformerFilterer{contract: contract}, nil
}

func bindOffRampWithMessageTransformer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OffRampWithMessageTransformerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OffRampWithMessageTransformer.Contract.OffRampWithMessageTransformerCaller.contract.Call(opts, result, method, params...)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.Contract.OffRampWithMessageTransformerTransactor.contract.Transfer(opts)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.Contract.OffRampWithMessageTransformerTransactor.contract.Transact(opts, method, params...)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OffRampWithMessageTransformer.Contract.contract.Call(opts, result, method, params...)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.Contract.contract.Transfer(opts)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.Contract.contract.Transact(opts, method, params...)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerCaller) CcipReceive(opts *bind.CallOpts, arg0 ClientAny2EVMMessage) error {
	var out []interface{}
	err := _OffRampWithMessageTransformer.contract.Call(opts, &out, "ccipReceive", arg0)

	if err != nil {
		return err
	}

	return err

}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerSession) CcipReceive(arg0 ClientAny2EVMMessage) error {
	return _OffRampWithMessageTransformer.Contract.CcipReceive(&_OffRampWithMessageTransformer.CallOpts, arg0)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerCallerSession) CcipReceive(arg0 ClientAny2EVMMessage) error {
	return _OffRampWithMessageTransformer.Contract.CcipReceive(&_OffRampWithMessageTransformer.CallOpts, arg0)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerCaller) GetAllSourceChainConfigs(opts *bind.CallOpts) ([]uint64, []OffRampSourceChainConfig, error) {
	var out []interface{}
	err := _OffRampWithMessageTransformer.contract.Call(opts, &out, "getAllSourceChainConfigs")

	if err != nil {
		return *new([]uint64), *new([]OffRampSourceChainConfig), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)
	out1 := *abi.ConvertType(out[1], new([]OffRampSourceChainConfig)).(*[]OffRampSourceChainConfig)

	return out0, out1, err

}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerSession) GetAllSourceChainConfigs() ([]uint64, []OffRampSourceChainConfig, error) {
	return _OffRampWithMessageTransformer.Contract.GetAllSourceChainConfigs(&_OffRampWithMessageTransformer.CallOpts)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerCallerSession) GetAllSourceChainConfigs() ([]uint64, []OffRampSourceChainConfig, error) {
	return _OffRampWithMessageTransformer.Contract.GetAllSourceChainConfigs(&_OffRampWithMessageTransformer.CallOpts)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerCaller) GetDynamicConfig(opts *bind.CallOpts) (OffRampDynamicConfig, error) {
	var out []interface{}
	err := _OffRampWithMessageTransformer.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(OffRampDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(OffRampDynamicConfig)).(*OffRampDynamicConfig)

	return out0, err

}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerSession) GetDynamicConfig() (OffRampDynamicConfig, error) {
	return _OffRampWithMessageTransformer.Contract.GetDynamicConfig(&_OffRampWithMessageTransformer.CallOpts)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerCallerSession) GetDynamicConfig() (OffRampDynamicConfig, error) {
	return _OffRampWithMessageTransformer.Contract.GetDynamicConfig(&_OffRampWithMessageTransformer.CallOpts)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerCaller) GetExecutionState(opts *bind.CallOpts, sourceChainSelector uint64, sequenceNumber uint64) (uint8, error) {
	var out []interface{}
	err := _OffRampWithMessageTransformer.contract.Call(opts, &out, "getExecutionState", sourceChainSelector, sequenceNumber)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerSession) GetExecutionState(sourceChainSelector uint64, sequenceNumber uint64) (uint8, error) {
	return _OffRampWithMessageTransformer.Contract.GetExecutionState(&_OffRampWithMessageTransformer.CallOpts, sourceChainSelector, sequenceNumber)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerCallerSession) GetExecutionState(sourceChainSelector uint64, sequenceNumber uint64) (uint8, error) {
	return _OffRampWithMessageTransformer.Contract.GetExecutionState(&_OffRampWithMessageTransformer.CallOpts, sourceChainSelector, sequenceNumber)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerCaller) GetLatestPriceSequenceNumber(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OffRampWithMessageTransformer.contract.Call(opts, &out, "getLatestPriceSequenceNumber")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerSession) GetLatestPriceSequenceNumber() (uint64, error) {
	return _OffRampWithMessageTransformer.Contract.GetLatestPriceSequenceNumber(&_OffRampWithMessageTransformer.CallOpts)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerCallerSession) GetLatestPriceSequenceNumber() (uint64, error) {
	return _OffRampWithMessageTransformer.Contract.GetLatestPriceSequenceNumber(&_OffRampWithMessageTransformer.CallOpts)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerCaller) GetMerkleRoot(opts *bind.CallOpts, sourceChainSelector uint64, root [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _OffRampWithMessageTransformer.contract.Call(opts, &out, "getMerkleRoot", sourceChainSelector, root)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerSession) GetMerkleRoot(sourceChainSelector uint64, root [32]byte) (*big.Int, error) {
	return _OffRampWithMessageTransformer.Contract.GetMerkleRoot(&_OffRampWithMessageTransformer.CallOpts, sourceChainSelector, root)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerCallerSession) GetMerkleRoot(sourceChainSelector uint64, root [32]byte) (*big.Int, error) {
	return _OffRampWithMessageTransformer.Contract.GetMerkleRoot(&_OffRampWithMessageTransformer.CallOpts, sourceChainSelector, root)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerCaller) GetMessageTransformer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OffRampWithMessageTransformer.contract.Call(opts, &out, "getMessageTransformer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerSession) GetMessageTransformer() (common.Address, error) {
	return _OffRampWithMessageTransformer.Contract.GetMessageTransformer(&_OffRampWithMessageTransformer.CallOpts)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerCallerSession) GetMessageTransformer() (common.Address, error) {
	return _OffRampWithMessageTransformer.Contract.GetMessageTransformer(&_OffRampWithMessageTransformer.CallOpts)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerCaller) GetSourceChainConfig(opts *bind.CallOpts, sourceChainSelector uint64) (OffRampSourceChainConfig, error) {
	var out []interface{}
	err := _OffRampWithMessageTransformer.contract.Call(opts, &out, "getSourceChainConfig", sourceChainSelector)

	if err != nil {
		return *new(OffRampSourceChainConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(OffRampSourceChainConfig)).(*OffRampSourceChainConfig)

	return out0, err

}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerSession) GetSourceChainConfig(sourceChainSelector uint64) (OffRampSourceChainConfig, error) {
	return _OffRampWithMessageTransformer.Contract.GetSourceChainConfig(&_OffRampWithMessageTransformer.CallOpts, sourceChainSelector)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerCallerSession) GetSourceChainConfig(sourceChainSelector uint64) (OffRampSourceChainConfig, error) {
	return _OffRampWithMessageTransformer.Contract.GetSourceChainConfig(&_OffRampWithMessageTransformer.CallOpts, sourceChainSelector)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerCaller) GetStaticConfig(opts *bind.CallOpts) (OffRampStaticConfig, error) {
	var out []interface{}
	err := _OffRampWithMessageTransformer.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(OffRampStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(OffRampStaticConfig)).(*OffRampStaticConfig)

	return out0, err

}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerSession) GetStaticConfig() (OffRampStaticConfig, error) {
	return _OffRampWithMessageTransformer.Contract.GetStaticConfig(&_OffRampWithMessageTransformer.CallOpts)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerCallerSession) GetStaticConfig() (OffRampStaticConfig, error) {
	return _OffRampWithMessageTransformer.Contract.GetStaticConfig(&_OffRampWithMessageTransformer.CallOpts)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerCaller) LatestConfigDetails(opts *bind.CallOpts, ocrPluginType uint8) (MultiOCR3BaseOCRConfig, error) {
	var out []interface{}
	err := _OffRampWithMessageTransformer.contract.Call(opts, &out, "latestConfigDetails", ocrPluginType)

	if err != nil {
		return *new(MultiOCR3BaseOCRConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(MultiOCR3BaseOCRConfig)).(*MultiOCR3BaseOCRConfig)

	return out0, err

}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerSession) LatestConfigDetails(ocrPluginType uint8) (MultiOCR3BaseOCRConfig, error) {
	return _OffRampWithMessageTransformer.Contract.LatestConfigDetails(&_OffRampWithMessageTransformer.CallOpts, ocrPluginType)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerCallerSession) LatestConfigDetails(ocrPluginType uint8) (MultiOCR3BaseOCRConfig, error) {
	return _OffRampWithMessageTransformer.Contract.LatestConfigDetails(&_OffRampWithMessageTransformer.CallOpts, ocrPluginType)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OffRampWithMessageTransformer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerSession) Owner() (common.Address, error) {
	return _OffRampWithMessageTransformer.Contract.Owner(&_OffRampWithMessageTransformer.CallOpts)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerCallerSession) Owner() (common.Address, error) {
	return _OffRampWithMessageTransformer.Contract.Owner(&_OffRampWithMessageTransformer.CallOpts)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OffRampWithMessageTransformer.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerSession) TypeAndVersion() (string, error) {
	return _OffRampWithMessageTransformer.Contract.TypeAndVersion(&_OffRampWithMessageTransformer.CallOpts)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerCallerSession) TypeAndVersion() (string, error) {
	return _OffRampWithMessageTransformer.Contract.TypeAndVersion(&_OffRampWithMessageTransformer.CallOpts)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.contract.Transact(opts, "acceptOwnership")
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerSession) AcceptOwnership() (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.Contract.AcceptOwnership(&_OffRampWithMessageTransformer.TransactOpts)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.Contract.AcceptOwnership(&_OffRampWithMessageTransformer.TransactOpts)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerTransactor) ApplySourceChainConfigUpdates(opts *bind.TransactOpts, sourceChainConfigUpdates []OffRampSourceChainConfigArgs) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.contract.Transact(opts, "applySourceChainConfigUpdates", sourceChainConfigUpdates)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerSession) ApplySourceChainConfigUpdates(sourceChainConfigUpdates []OffRampSourceChainConfigArgs) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.Contract.ApplySourceChainConfigUpdates(&_OffRampWithMessageTransformer.TransactOpts, sourceChainConfigUpdates)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerTransactorSession) ApplySourceChainConfigUpdates(sourceChainConfigUpdates []OffRampSourceChainConfigArgs) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.Contract.ApplySourceChainConfigUpdates(&_OffRampWithMessageTransformer.TransactOpts, sourceChainConfigUpdates)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerTransactor) Commit(opts *bind.TransactOpts, reportContext [2][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.contract.Transact(opts, "commit", reportContext, report, rs, ss, rawVs)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerSession) Commit(reportContext [2][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.Contract.Commit(&_OffRampWithMessageTransformer.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerTransactorSession) Commit(reportContext [2][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.Contract.Commit(&_OffRampWithMessageTransformer.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerTransactor) Execute(opts *bind.TransactOpts, reportContext [2][32]byte, report []byte) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.contract.Transact(opts, "execute", reportContext, report)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerSession) Execute(reportContext [2][32]byte, report []byte) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.Contract.Execute(&_OffRampWithMessageTransformer.TransactOpts, reportContext, report)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerTransactorSession) Execute(reportContext [2][32]byte, report []byte) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.Contract.Execute(&_OffRampWithMessageTransformer.TransactOpts, reportContext, report)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message InternalAny2EVMRampMessage, offchainTokenData [][]byte, tokenGasOverrides []uint32) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.contract.Transact(opts, "executeSingleMessage", message, offchainTokenData, tokenGasOverrides)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerSession) ExecuteSingleMessage(message InternalAny2EVMRampMessage, offchainTokenData [][]byte, tokenGasOverrides []uint32) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.Contract.ExecuteSingleMessage(&_OffRampWithMessageTransformer.TransactOpts, message, offchainTokenData, tokenGasOverrides)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerTransactorSession) ExecuteSingleMessage(message InternalAny2EVMRampMessage, offchainTokenData [][]byte, tokenGasOverrides []uint32) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.Contract.ExecuteSingleMessage(&_OffRampWithMessageTransformer.TransactOpts, message, offchainTokenData, tokenGasOverrides)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerTransactor) ManuallyExecute(opts *bind.TransactOpts, reports []InternalExecutionReport, gasLimitOverrides [][]OffRampGasLimitOverride) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.contract.Transact(opts, "manuallyExecute", reports, gasLimitOverrides)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerSession) ManuallyExecute(reports []InternalExecutionReport, gasLimitOverrides [][]OffRampGasLimitOverride) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.Contract.ManuallyExecute(&_OffRampWithMessageTransformer.TransactOpts, reports, gasLimitOverrides)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerTransactorSession) ManuallyExecute(reports []InternalExecutionReport, gasLimitOverrides [][]OffRampGasLimitOverride) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.Contract.ManuallyExecute(&_OffRampWithMessageTransformer.TransactOpts, reports, gasLimitOverrides)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig OffRampDynamicConfig) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerSession) SetDynamicConfig(dynamicConfig OffRampDynamicConfig) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.Contract.SetDynamicConfig(&_OffRampWithMessageTransformer.TransactOpts, dynamicConfig)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerTransactorSession) SetDynamicConfig(dynamicConfig OffRampDynamicConfig) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.Contract.SetDynamicConfig(&_OffRampWithMessageTransformer.TransactOpts, dynamicConfig)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerTransactor) SetMessageTransformer(opts *bind.TransactOpts, messageTransformerAddr common.Address) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.contract.Transact(opts, "setMessageTransformer", messageTransformerAddr)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerSession) SetMessageTransformer(messageTransformerAddr common.Address) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.Contract.SetMessageTransformer(&_OffRampWithMessageTransformer.TransactOpts, messageTransformerAddr)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerTransactorSession) SetMessageTransformer(messageTransformerAddr common.Address) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.Contract.SetMessageTransformer(&_OffRampWithMessageTransformer.TransactOpts, messageTransformerAddr)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerTransactor) SetOCR3Configs(opts *bind.TransactOpts, ocrConfigArgs []MultiOCR3BaseOCRConfigArgs) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.contract.Transact(opts, "setOCR3Configs", ocrConfigArgs)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerSession) SetOCR3Configs(ocrConfigArgs []MultiOCR3BaseOCRConfigArgs) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.Contract.SetOCR3Configs(&_OffRampWithMessageTransformer.TransactOpts, ocrConfigArgs)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerTransactorSession) SetOCR3Configs(ocrConfigArgs []MultiOCR3BaseOCRConfigArgs) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.Contract.SetOCR3Configs(&_OffRampWithMessageTransformer.TransactOpts, ocrConfigArgs)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.contract.Transact(opts, "transferOwnership", to)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.Contract.TransferOwnership(&_OffRampWithMessageTransformer.TransactOpts, to)
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OffRampWithMessageTransformer.Contract.TransferOwnership(&_OffRampWithMessageTransformer.TransactOpts, to)
}

type OffRampWithMessageTransformerAlreadyAttemptedIterator struct {
	Event *OffRampWithMessageTransformerAlreadyAttempted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampWithMessageTransformerAlreadyAttemptedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampWithMessageTransformerAlreadyAttempted)
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
		it.Event = new(OffRampWithMessageTransformerAlreadyAttempted)
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

func (it *OffRampWithMessageTransformerAlreadyAttemptedIterator) Error() error {
	return it.fail
}

func (it *OffRampWithMessageTransformerAlreadyAttemptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampWithMessageTransformerAlreadyAttempted struct {
	SourceChainSelector uint64
	SequenceNumber      uint64
	Raw                 types.Log
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) FilterAlreadyAttempted(opts *bind.FilterOpts) (*OffRampWithMessageTransformerAlreadyAttemptedIterator, error) {

	logs, sub, err := _OffRampWithMessageTransformer.contract.FilterLogs(opts, "AlreadyAttempted")
	if err != nil {
		return nil, err
	}
	return &OffRampWithMessageTransformerAlreadyAttemptedIterator{contract: _OffRampWithMessageTransformer.contract, event: "AlreadyAttempted", logs: logs, sub: sub}, nil
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) WatchAlreadyAttempted(opts *bind.WatchOpts, sink chan<- *OffRampWithMessageTransformerAlreadyAttempted) (event.Subscription, error) {

	logs, sub, err := _OffRampWithMessageTransformer.contract.WatchLogs(opts, "AlreadyAttempted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampWithMessageTransformerAlreadyAttempted)
				if err := _OffRampWithMessageTransformer.contract.UnpackLog(event, "AlreadyAttempted", log); err != nil {
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

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) ParseAlreadyAttempted(log types.Log) (*OffRampWithMessageTransformerAlreadyAttempted, error) {
	event := new(OffRampWithMessageTransformerAlreadyAttempted)
	if err := _OffRampWithMessageTransformer.contract.UnpackLog(event, "AlreadyAttempted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampWithMessageTransformerCommitReportAcceptedIterator struct {
	Event *OffRampWithMessageTransformerCommitReportAccepted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampWithMessageTransformerCommitReportAcceptedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampWithMessageTransformerCommitReportAccepted)
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
		it.Event = new(OffRampWithMessageTransformerCommitReportAccepted)
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

func (it *OffRampWithMessageTransformerCommitReportAcceptedIterator) Error() error {
	return it.fail
}

func (it *OffRampWithMessageTransformerCommitReportAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampWithMessageTransformerCommitReportAccepted struct {
	BlessedMerkleRoots   []InternalMerkleRoot
	UnblessedMerkleRoots []InternalMerkleRoot
	PriceUpdates         InternalPriceUpdates
	Raw                  types.Log
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) FilterCommitReportAccepted(opts *bind.FilterOpts) (*OffRampWithMessageTransformerCommitReportAcceptedIterator, error) {

	logs, sub, err := _OffRampWithMessageTransformer.contract.FilterLogs(opts, "CommitReportAccepted")
	if err != nil {
		return nil, err
	}
	return &OffRampWithMessageTransformerCommitReportAcceptedIterator{contract: _OffRampWithMessageTransformer.contract, event: "CommitReportAccepted", logs: logs, sub: sub}, nil
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) WatchCommitReportAccepted(opts *bind.WatchOpts, sink chan<- *OffRampWithMessageTransformerCommitReportAccepted) (event.Subscription, error) {

	logs, sub, err := _OffRampWithMessageTransformer.contract.WatchLogs(opts, "CommitReportAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampWithMessageTransformerCommitReportAccepted)
				if err := _OffRampWithMessageTransformer.contract.UnpackLog(event, "CommitReportAccepted", log); err != nil {
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

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) ParseCommitReportAccepted(log types.Log) (*OffRampWithMessageTransformerCommitReportAccepted, error) {
	event := new(OffRampWithMessageTransformerCommitReportAccepted)
	if err := _OffRampWithMessageTransformer.contract.UnpackLog(event, "CommitReportAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampWithMessageTransformerConfigSetIterator struct {
	Event *OffRampWithMessageTransformerConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampWithMessageTransformerConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampWithMessageTransformerConfigSet)
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
		it.Event = new(OffRampWithMessageTransformerConfigSet)
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

func (it *OffRampWithMessageTransformerConfigSetIterator) Error() error {
	return it.fail
}

func (it *OffRampWithMessageTransformerConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampWithMessageTransformerConfigSet struct {
	OcrPluginType uint8
	ConfigDigest  [32]byte
	Signers       []common.Address
	Transmitters  []common.Address
	F             uint8
	Raw           types.Log
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) FilterConfigSet(opts *bind.FilterOpts) (*OffRampWithMessageTransformerConfigSetIterator, error) {

	logs, sub, err := _OffRampWithMessageTransformer.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &OffRampWithMessageTransformerConfigSetIterator{contract: _OffRampWithMessageTransformer.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampWithMessageTransformerConfigSet) (event.Subscription, error) {

	logs, sub, err := _OffRampWithMessageTransformer.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampWithMessageTransformerConfigSet)
				if err := _OffRampWithMessageTransformer.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) ParseConfigSet(log types.Log) (*OffRampWithMessageTransformerConfigSet, error) {
	event := new(OffRampWithMessageTransformerConfigSet)
	if err := _OffRampWithMessageTransformer.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampWithMessageTransformerDynamicConfigSetIterator struct {
	Event *OffRampWithMessageTransformerDynamicConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampWithMessageTransformerDynamicConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampWithMessageTransformerDynamicConfigSet)
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
		it.Event = new(OffRampWithMessageTransformerDynamicConfigSet)
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

func (it *OffRampWithMessageTransformerDynamicConfigSetIterator) Error() error {
	return it.fail
}

func (it *OffRampWithMessageTransformerDynamicConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampWithMessageTransformerDynamicConfigSet struct {
	DynamicConfig OffRampDynamicConfig
	Raw           types.Log
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) FilterDynamicConfigSet(opts *bind.FilterOpts) (*OffRampWithMessageTransformerDynamicConfigSetIterator, error) {

	logs, sub, err := _OffRampWithMessageTransformer.contract.FilterLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return &OffRampWithMessageTransformerDynamicConfigSetIterator{contract: _OffRampWithMessageTransformer.contract, event: "DynamicConfigSet", logs: logs, sub: sub}, nil
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampWithMessageTransformerDynamicConfigSet) (event.Subscription, error) {

	logs, sub, err := _OffRampWithMessageTransformer.contract.WatchLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampWithMessageTransformerDynamicConfigSet)
				if err := _OffRampWithMessageTransformer.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
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

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) ParseDynamicConfigSet(log types.Log) (*OffRampWithMessageTransformerDynamicConfigSet, error) {
	event := new(OffRampWithMessageTransformerDynamicConfigSet)
	if err := _OffRampWithMessageTransformer.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampWithMessageTransformerExecutionStateChangedIterator struct {
	Event *OffRampWithMessageTransformerExecutionStateChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampWithMessageTransformerExecutionStateChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampWithMessageTransformerExecutionStateChanged)
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
		it.Event = new(OffRampWithMessageTransformerExecutionStateChanged)
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

func (it *OffRampWithMessageTransformerExecutionStateChangedIterator) Error() error {
	return it.fail
}

func (it *OffRampWithMessageTransformerExecutionStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampWithMessageTransformerExecutionStateChanged struct {
	SourceChainSelector uint64
	SequenceNumber      uint64
	MessageId           [32]byte
	MessageHash         [32]byte
	State               uint8
	ReturnData          []byte
	GasUsed             *big.Int
	Raw                 types.Log
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) FilterExecutionStateChanged(opts *bind.FilterOpts, sourceChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (*OffRampWithMessageTransformerExecutionStateChangedIterator, error) {

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

	logs, sub, err := _OffRampWithMessageTransformer.contract.FilterLogs(opts, "ExecutionStateChanged", sourceChainSelectorRule, sequenceNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return &OffRampWithMessageTransformerExecutionStateChangedIterator{contract: _OffRampWithMessageTransformer.contract, event: "ExecutionStateChanged", logs: logs, sub: sub}, nil
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *OffRampWithMessageTransformerExecutionStateChanged, sourceChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _OffRampWithMessageTransformer.contract.WatchLogs(opts, "ExecutionStateChanged", sourceChainSelectorRule, sequenceNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampWithMessageTransformerExecutionStateChanged)
				if err := _OffRampWithMessageTransformer.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
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

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) ParseExecutionStateChanged(log types.Log) (*OffRampWithMessageTransformerExecutionStateChanged, error) {
	event := new(OffRampWithMessageTransformerExecutionStateChanged)
	if err := _OffRampWithMessageTransformer.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampWithMessageTransformerOwnershipTransferRequestedIterator struct {
	Event *OffRampWithMessageTransformerOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampWithMessageTransformerOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampWithMessageTransformerOwnershipTransferRequested)
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
		it.Event = new(OffRampWithMessageTransformerOwnershipTransferRequested)
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

func (it *OffRampWithMessageTransformerOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *OffRampWithMessageTransformerOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampWithMessageTransformerOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OffRampWithMessageTransformerOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffRampWithMessageTransformer.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OffRampWithMessageTransformerOwnershipTransferRequestedIterator{contract: _OffRampWithMessageTransformer.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OffRampWithMessageTransformerOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffRampWithMessageTransformer.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampWithMessageTransformerOwnershipTransferRequested)
				if err := _OffRampWithMessageTransformer.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) ParseOwnershipTransferRequested(log types.Log) (*OffRampWithMessageTransformerOwnershipTransferRequested, error) {
	event := new(OffRampWithMessageTransformerOwnershipTransferRequested)
	if err := _OffRampWithMessageTransformer.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampWithMessageTransformerOwnershipTransferredIterator struct {
	Event *OffRampWithMessageTransformerOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampWithMessageTransformerOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampWithMessageTransformerOwnershipTransferred)
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
		it.Event = new(OffRampWithMessageTransformerOwnershipTransferred)
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

func (it *OffRampWithMessageTransformerOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *OffRampWithMessageTransformerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampWithMessageTransformerOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OffRampWithMessageTransformerOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffRampWithMessageTransformer.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OffRampWithMessageTransformerOwnershipTransferredIterator{contract: _OffRampWithMessageTransformer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OffRampWithMessageTransformerOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffRampWithMessageTransformer.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampWithMessageTransformerOwnershipTransferred)
				if err := _OffRampWithMessageTransformer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) ParseOwnershipTransferred(log types.Log) (*OffRampWithMessageTransformerOwnershipTransferred, error) {
	event := new(OffRampWithMessageTransformerOwnershipTransferred)
	if err := _OffRampWithMessageTransformer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampWithMessageTransformerRootRemovedIterator struct {
	Event *OffRampWithMessageTransformerRootRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampWithMessageTransformerRootRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampWithMessageTransformerRootRemoved)
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
		it.Event = new(OffRampWithMessageTransformerRootRemoved)
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

func (it *OffRampWithMessageTransformerRootRemovedIterator) Error() error {
	return it.fail
}

func (it *OffRampWithMessageTransformerRootRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampWithMessageTransformerRootRemoved struct {
	Root [32]byte
	Raw  types.Log
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) FilterRootRemoved(opts *bind.FilterOpts) (*OffRampWithMessageTransformerRootRemovedIterator, error) {

	logs, sub, err := _OffRampWithMessageTransformer.contract.FilterLogs(opts, "RootRemoved")
	if err != nil {
		return nil, err
	}
	return &OffRampWithMessageTransformerRootRemovedIterator{contract: _OffRampWithMessageTransformer.contract, event: "RootRemoved", logs: logs, sub: sub}, nil
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) WatchRootRemoved(opts *bind.WatchOpts, sink chan<- *OffRampWithMessageTransformerRootRemoved) (event.Subscription, error) {

	logs, sub, err := _OffRampWithMessageTransformer.contract.WatchLogs(opts, "RootRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampWithMessageTransformerRootRemoved)
				if err := _OffRampWithMessageTransformer.contract.UnpackLog(event, "RootRemoved", log); err != nil {
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

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) ParseRootRemoved(log types.Log) (*OffRampWithMessageTransformerRootRemoved, error) {
	event := new(OffRampWithMessageTransformerRootRemoved)
	if err := _OffRampWithMessageTransformer.contract.UnpackLog(event, "RootRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampWithMessageTransformerSkippedAlreadyExecutedMessageIterator struct {
	Event *OffRampWithMessageTransformerSkippedAlreadyExecutedMessage

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampWithMessageTransformerSkippedAlreadyExecutedMessageIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampWithMessageTransformerSkippedAlreadyExecutedMessage)
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
		it.Event = new(OffRampWithMessageTransformerSkippedAlreadyExecutedMessage)
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

func (it *OffRampWithMessageTransformerSkippedAlreadyExecutedMessageIterator) Error() error {
	return it.fail
}

func (it *OffRampWithMessageTransformerSkippedAlreadyExecutedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampWithMessageTransformerSkippedAlreadyExecutedMessage struct {
	SourceChainSelector uint64
	SequenceNumber      uint64
	Raw                 types.Log
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) FilterSkippedAlreadyExecutedMessage(opts *bind.FilterOpts) (*OffRampWithMessageTransformerSkippedAlreadyExecutedMessageIterator, error) {

	logs, sub, err := _OffRampWithMessageTransformer.contract.FilterLogs(opts, "SkippedAlreadyExecutedMessage")
	if err != nil {
		return nil, err
	}
	return &OffRampWithMessageTransformerSkippedAlreadyExecutedMessageIterator{contract: _OffRampWithMessageTransformer.contract, event: "SkippedAlreadyExecutedMessage", logs: logs, sub: sub}, nil
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) WatchSkippedAlreadyExecutedMessage(opts *bind.WatchOpts, sink chan<- *OffRampWithMessageTransformerSkippedAlreadyExecutedMessage) (event.Subscription, error) {

	logs, sub, err := _OffRampWithMessageTransformer.contract.WatchLogs(opts, "SkippedAlreadyExecutedMessage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampWithMessageTransformerSkippedAlreadyExecutedMessage)
				if err := _OffRampWithMessageTransformer.contract.UnpackLog(event, "SkippedAlreadyExecutedMessage", log); err != nil {
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

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) ParseSkippedAlreadyExecutedMessage(log types.Log) (*OffRampWithMessageTransformerSkippedAlreadyExecutedMessage, error) {
	event := new(OffRampWithMessageTransformerSkippedAlreadyExecutedMessage)
	if err := _OffRampWithMessageTransformer.contract.UnpackLog(event, "SkippedAlreadyExecutedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampWithMessageTransformerSkippedReportExecutionIterator struct {
	Event *OffRampWithMessageTransformerSkippedReportExecution

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampWithMessageTransformerSkippedReportExecutionIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampWithMessageTransformerSkippedReportExecution)
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
		it.Event = new(OffRampWithMessageTransformerSkippedReportExecution)
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

func (it *OffRampWithMessageTransformerSkippedReportExecutionIterator) Error() error {
	return it.fail
}

func (it *OffRampWithMessageTransformerSkippedReportExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampWithMessageTransformerSkippedReportExecution struct {
	SourceChainSelector uint64
	Raw                 types.Log
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) FilterSkippedReportExecution(opts *bind.FilterOpts) (*OffRampWithMessageTransformerSkippedReportExecutionIterator, error) {

	logs, sub, err := _OffRampWithMessageTransformer.contract.FilterLogs(opts, "SkippedReportExecution")
	if err != nil {
		return nil, err
	}
	return &OffRampWithMessageTransformerSkippedReportExecutionIterator{contract: _OffRampWithMessageTransformer.contract, event: "SkippedReportExecution", logs: logs, sub: sub}, nil
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) WatchSkippedReportExecution(opts *bind.WatchOpts, sink chan<- *OffRampWithMessageTransformerSkippedReportExecution) (event.Subscription, error) {

	logs, sub, err := _OffRampWithMessageTransformer.contract.WatchLogs(opts, "SkippedReportExecution")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampWithMessageTransformerSkippedReportExecution)
				if err := _OffRampWithMessageTransformer.contract.UnpackLog(event, "SkippedReportExecution", log); err != nil {
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

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) ParseSkippedReportExecution(log types.Log) (*OffRampWithMessageTransformerSkippedReportExecution, error) {
	event := new(OffRampWithMessageTransformerSkippedReportExecution)
	if err := _OffRampWithMessageTransformer.contract.UnpackLog(event, "SkippedReportExecution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampWithMessageTransformerSourceChainConfigSetIterator struct {
	Event *OffRampWithMessageTransformerSourceChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampWithMessageTransformerSourceChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampWithMessageTransformerSourceChainConfigSet)
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
		it.Event = new(OffRampWithMessageTransformerSourceChainConfigSet)
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

func (it *OffRampWithMessageTransformerSourceChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *OffRampWithMessageTransformerSourceChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampWithMessageTransformerSourceChainConfigSet struct {
	SourceChainSelector uint64
	SourceConfig        OffRampSourceChainConfig
	Raw                 types.Log
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) FilterSourceChainConfigSet(opts *bind.FilterOpts, sourceChainSelector []uint64) (*OffRampWithMessageTransformerSourceChainConfigSetIterator, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}

	logs, sub, err := _OffRampWithMessageTransformer.contract.FilterLogs(opts, "SourceChainConfigSet", sourceChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &OffRampWithMessageTransformerSourceChainConfigSetIterator{contract: _OffRampWithMessageTransformer.contract, event: "SourceChainConfigSet", logs: logs, sub: sub}, nil
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) WatchSourceChainConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampWithMessageTransformerSourceChainConfigSet, sourceChainSelector []uint64) (event.Subscription, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}

	logs, sub, err := _OffRampWithMessageTransformer.contract.WatchLogs(opts, "SourceChainConfigSet", sourceChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampWithMessageTransformerSourceChainConfigSet)
				if err := _OffRampWithMessageTransformer.contract.UnpackLog(event, "SourceChainConfigSet", log); err != nil {
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

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) ParseSourceChainConfigSet(log types.Log) (*OffRampWithMessageTransformerSourceChainConfigSet, error) {
	event := new(OffRampWithMessageTransformerSourceChainConfigSet)
	if err := _OffRampWithMessageTransformer.contract.UnpackLog(event, "SourceChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampWithMessageTransformerSourceChainSelectorAddedIterator struct {
	Event *OffRampWithMessageTransformerSourceChainSelectorAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampWithMessageTransformerSourceChainSelectorAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampWithMessageTransformerSourceChainSelectorAdded)
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
		it.Event = new(OffRampWithMessageTransformerSourceChainSelectorAdded)
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

func (it *OffRampWithMessageTransformerSourceChainSelectorAddedIterator) Error() error {
	return it.fail
}

func (it *OffRampWithMessageTransformerSourceChainSelectorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampWithMessageTransformerSourceChainSelectorAdded struct {
	SourceChainSelector uint64
	Raw                 types.Log
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) FilterSourceChainSelectorAdded(opts *bind.FilterOpts) (*OffRampWithMessageTransformerSourceChainSelectorAddedIterator, error) {

	logs, sub, err := _OffRampWithMessageTransformer.contract.FilterLogs(opts, "SourceChainSelectorAdded")
	if err != nil {
		return nil, err
	}
	return &OffRampWithMessageTransformerSourceChainSelectorAddedIterator{contract: _OffRampWithMessageTransformer.contract, event: "SourceChainSelectorAdded", logs: logs, sub: sub}, nil
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) WatchSourceChainSelectorAdded(opts *bind.WatchOpts, sink chan<- *OffRampWithMessageTransformerSourceChainSelectorAdded) (event.Subscription, error) {

	logs, sub, err := _OffRampWithMessageTransformer.contract.WatchLogs(opts, "SourceChainSelectorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampWithMessageTransformerSourceChainSelectorAdded)
				if err := _OffRampWithMessageTransformer.contract.UnpackLog(event, "SourceChainSelectorAdded", log); err != nil {
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

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) ParseSourceChainSelectorAdded(log types.Log) (*OffRampWithMessageTransformerSourceChainSelectorAdded, error) {
	event := new(OffRampWithMessageTransformerSourceChainSelectorAdded)
	if err := _OffRampWithMessageTransformer.contract.UnpackLog(event, "SourceChainSelectorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampWithMessageTransformerStaticConfigSetIterator struct {
	Event *OffRampWithMessageTransformerStaticConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampWithMessageTransformerStaticConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampWithMessageTransformerStaticConfigSet)
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
		it.Event = new(OffRampWithMessageTransformerStaticConfigSet)
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

func (it *OffRampWithMessageTransformerStaticConfigSetIterator) Error() error {
	return it.fail
}

func (it *OffRampWithMessageTransformerStaticConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampWithMessageTransformerStaticConfigSet struct {
	StaticConfig OffRampStaticConfig
	Raw          types.Log
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) FilterStaticConfigSet(opts *bind.FilterOpts) (*OffRampWithMessageTransformerStaticConfigSetIterator, error) {

	logs, sub, err := _OffRampWithMessageTransformer.contract.FilterLogs(opts, "StaticConfigSet")
	if err != nil {
		return nil, err
	}
	return &OffRampWithMessageTransformerStaticConfigSetIterator{contract: _OffRampWithMessageTransformer.contract, event: "StaticConfigSet", logs: logs, sub: sub}, nil
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) WatchStaticConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampWithMessageTransformerStaticConfigSet) (event.Subscription, error) {

	logs, sub, err := _OffRampWithMessageTransformer.contract.WatchLogs(opts, "StaticConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampWithMessageTransformerStaticConfigSet)
				if err := _OffRampWithMessageTransformer.contract.UnpackLog(event, "StaticConfigSet", log); err != nil {
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

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) ParseStaticConfigSet(log types.Log) (*OffRampWithMessageTransformerStaticConfigSet, error) {
	event := new(OffRampWithMessageTransformerStaticConfigSet)
	if err := _OffRampWithMessageTransformer.contract.UnpackLog(event, "StaticConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampWithMessageTransformerTransmittedIterator struct {
	Event *OffRampWithMessageTransformerTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampWithMessageTransformerTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampWithMessageTransformerTransmitted)
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
		it.Event = new(OffRampWithMessageTransformerTransmitted)
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

func (it *OffRampWithMessageTransformerTransmittedIterator) Error() error {
	return it.fail
}

func (it *OffRampWithMessageTransformerTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampWithMessageTransformerTransmitted struct {
	OcrPluginType  uint8
	ConfigDigest   [32]byte
	SequenceNumber uint64
	Raw            types.Log
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) FilterTransmitted(opts *bind.FilterOpts, ocrPluginType []uint8) (*OffRampWithMessageTransformerTransmittedIterator, error) {

	var ocrPluginTypeRule []interface{}
	for _, ocrPluginTypeItem := range ocrPluginType {
		ocrPluginTypeRule = append(ocrPluginTypeRule, ocrPluginTypeItem)
	}

	logs, sub, err := _OffRampWithMessageTransformer.contract.FilterLogs(opts, "Transmitted", ocrPluginTypeRule)
	if err != nil {
		return nil, err
	}
	return &OffRampWithMessageTransformerTransmittedIterator{contract: _OffRampWithMessageTransformer.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *OffRampWithMessageTransformerTransmitted, ocrPluginType []uint8) (event.Subscription, error) {

	var ocrPluginTypeRule []interface{}
	for _, ocrPluginTypeItem := range ocrPluginType {
		ocrPluginTypeRule = append(ocrPluginTypeRule, ocrPluginTypeItem)
	}

	logs, sub, err := _OffRampWithMessageTransformer.contract.WatchLogs(opts, "Transmitted", ocrPluginTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampWithMessageTransformerTransmitted)
				if err := _OffRampWithMessageTransformer.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformerFilterer) ParseTransmitted(log types.Log) (*OffRampWithMessageTransformerTransmitted, error) {
	event := new(OffRampWithMessageTransformerTransmitted)
	if err := _OffRampWithMessageTransformer.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (OffRampWithMessageTransformerAlreadyAttempted) Topic() common.Hash {
	return common.HexToHash("0x3ef2a99c550a751d4b0b261268f05a803dfb049ab43616a1ffb388f61fe65120")
}

func (OffRampWithMessageTransformerCommitReportAccepted) Topic() common.Hash {
	return common.HexToHash("0xb967c9b9e1b7af9a61ca71ff00e9f5b89ec6f2e268de8dacf12f0de8e51f3e47")
}

func (OffRampWithMessageTransformerConfigSet) Topic() common.Hash {
	return common.HexToHash("0xab8b1b57514019638d7b5ce9c638fe71366fe8e2be1c40a7a80f1733d0e9f547")
}

func (OffRampWithMessageTransformerDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0xa1c15688cb2c24508e158f6942b9276c6f3028a85e1af8cf3fff0c3ff3d5fc8d")
}

func (OffRampWithMessageTransformerExecutionStateChanged) Topic() common.Hash {
	return common.HexToHash("0x05665fe9ad095383d018353f4cbcba77e84db27dd215081bbf7cdf9ae6fbe48b")
}

func (OffRampWithMessageTransformerOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (OffRampWithMessageTransformerOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (OffRampWithMessageTransformerRootRemoved) Topic() common.Hash {
	return common.HexToHash("0x202f1139a3e334b6056064c0e9b19fd07e44a88d8f6e5ded571b24cf8c371f12")
}

func (OffRampWithMessageTransformerSkippedAlreadyExecutedMessage) Topic() common.Hash {
	return common.HexToHash("0x3b575419319662b2a6f5e2467d84521517a3382b908eb3d557bb3fdb0c50e23c")
}

func (OffRampWithMessageTransformerSkippedReportExecution) Topic() common.Hash {
	return common.HexToHash("0xaab522ed53d887e56ed53dd37398a01aeef6a58e0fa77c2173beb9512d894933")
}

func (OffRampWithMessageTransformerSourceChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0xbd1ab25a0ff0a36a588597ba1af11e30f3f210de8b9e818cc9bbc457c94c8d8c")
}

func (OffRampWithMessageTransformerSourceChainSelectorAdded) Topic() common.Hash {
	return common.HexToHash("0xf4c1390c70e5c0f491ae1ccbc06f9117cbbadf2767b247b3bc203280f24c0fb9")
}

func (OffRampWithMessageTransformerStaticConfigSet) Topic() common.Hash {
	return common.HexToHash("0xb0fa1fb01508c5097c502ad056fd77018870c9be9a86d9e56b6b471862d7c5b7")
}

func (OffRampWithMessageTransformerTransmitted) Topic() common.Hash {
	return common.HexToHash("0x198d6990ef96613a9026203077e422916918b03ff47f0be6bee7b02d8e139ef0")
}

func (_OffRampWithMessageTransformer *OffRampWithMessageTransformer) Address() common.Address {
	return _OffRampWithMessageTransformer.address
}

type OffRampWithMessageTransformerInterface interface {
	CcipReceive(opts *bind.CallOpts, arg0 ClientAny2EVMMessage) error

	GetAllSourceChainConfigs(opts *bind.CallOpts) ([]uint64, []OffRampSourceChainConfig, error)

	GetDynamicConfig(opts *bind.CallOpts) (OffRampDynamicConfig, error)

	GetExecutionState(opts *bind.CallOpts, sourceChainSelector uint64, sequenceNumber uint64) (uint8, error)

	GetLatestPriceSequenceNumber(opts *bind.CallOpts) (uint64, error)

	GetMerkleRoot(opts *bind.CallOpts, sourceChainSelector uint64, root [32]byte) (*big.Int, error)

	GetMessageTransformer(opts *bind.CallOpts) (common.Address, error)

	GetSourceChainConfig(opts *bind.CallOpts, sourceChainSelector uint64) (OffRampSourceChainConfig, error)

	GetStaticConfig(opts *bind.CallOpts) (OffRampStaticConfig, error)

	LatestConfigDetails(opts *bind.CallOpts, ocrPluginType uint8) (MultiOCR3BaseOCRConfig, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplySourceChainConfigUpdates(opts *bind.TransactOpts, sourceChainConfigUpdates []OffRampSourceChainConfigArgs) (*types.Transaction, error)

	Commit(opts *bind.TransactOpts, reportContext [2][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error)

	Execute(opts *bind.TransactOpts, reportContext [2][32]byte, report []byte) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message InternalAny2EVMRampMessage, offchainTokenData [][]byte, tokenGasOverrides []uint32) (*types.Transaction, error)

	ManuallyExecute(opts *bind.TransactOpts, reports []InternalExecutionReport, gasLimitOverrides [][]OffRampGasLimitOverride) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig OffRampDynamicConfig) (*types.Transaction, error)

	SetMessageTransformer(opts *bind.TransactOpts, messageTransformerAddr common.Address) (*types.Transaction, error)

	SetOCR3Configs(opts *bind.TransactOpts, ocrConfigArgs []MultiOCR3BaseOCRConfigArgs) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterAlreadyAttempted(opts *bind.FilterOpts) (*OffRampWithMessageTransformerAlreadyAttemptedIterator, error)

	WatchAlreadyAttempted(opts *bind.WatchOpts, sink chan<- *OffRampWithMessageTransformerAlreadyAttempted) (event.Subscription, error)

	ParseAlreadyAttempted(log types.Log) (*OffRampWithMessageTransformerAlreadyAttempted, error)

	FilterCommitReportAccepted(opts *bind.FilterOpts) (*OffRampWithMessageTransformerCommitReportAcceptedIterator, error)

	WatchCommitReportAccepted(opts *bind.WatchOpts, sink chan<- *OffRampWithMessageTransformerCommitReportAccepted) (event.Subscription, error)

	ParseCommitReportAccepted(log types.Log) (*OffRampWithMessageTransformerCommitReportAccepted, error)

	FilterConfigSet(opts *bind.FilterOpts) (*OffRampWithMessageTransformerConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampWithMessageTransformerConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*OffRampWithMessageTransformerConfigSet, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*OffRampWithMessageTransformerDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampWithMessageTransformerDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*OffRampWithMessageTransformerDynamicConfigSet, error)

	FilterExecutionStateChanged(opts *bind.FilterOpts, sourceChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (*OffRampWithMessageTransformerExecutionStateChangedIterator, error)

	WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *OffRampWithMessageTransformerExecutionStateChanged, sourceChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error)

	ParseExecutionStateChanged(log types.Log) (*OffRampWithMessageTransformerExecutionStateChanged, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OffRampWithMessageTransformerOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OffRampWithMessageTransformerOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*OffRampWithMessageTransformerOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OffRampWithMessageTransformerOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OffRampWithMessageTransformerOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*OffRampWithMessageTransformerOwnershipTransferred, error)

	FilterRootRemoved(opts *bind.FilterOpts) (*OffRampWithMessageTransformerRootRemovedIterator, error)

	WatchRootRemoved(opts *bind.WatchOpts, sink chan<- *OffRampWithMessageTransformerRootRemoved) (event.Subscription, error)

	ParseRootRemoved(log types.Log) (*OffRampWithMessageTransformerRootRemoved, error)

	FilterSkippedAlreadyExecutedMessage(opts *bind.FilterOpts) (*OffRampWithMessageTransformerSkippedAlreadyExecutedMessageIterator, error)

	WatchSkippedAlreadyExecutedMessage(opts *bind.WatchOpts, sink chan<- *OffRampWithMessageTransformerSkippedAlreadyExecutedMessage) (event.Subscription, error)

	ParseSkippedAlreadyExecutedMessage(log types.Log) (*OffRampWithMessageTransformerSkippedAlreadyExecutedMessage, error)

	FilterSkippedReportExecution(opts *bind.FilterOpts) (*OffRampWithMessageTransformerSkippedReportExecutionIterator, error)

	WatchSkippedReportExecution(opts *bind.WatchOpts, sink chan<- *OffRampWithMessageTransformerSkippedReportExecution) (event.Subscription, error)

	ParseSkippedReportExecution(log types.Log) (*OffRampWithMessageTransformerSkippedReportExecution, error)

	FilterSourceChainConfigSet(opts *bind.FilterOpts, sourceChainSelector []uint64) (*OffRampWithMessageTransformerSourceChainConfigSetIterator, error)

	WatchSourceChainConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampWithMessageTransformerSourceChainConfigSet, sourceChainSelector []uint64) (event.Subscription, error)

	ParseSourceChainConfigSet(log types.Log) (*OffRampWithMessageTransformerSourceChainConfigSet, error)

	FilterSourceChainSelectorAdded(opts *bind.FilterOpts) (*OffRampWithMessageTransformerSourceChainSelectorAddedIterator, error)

	WatchSourceChainSelectorAdded(opts *bind.WatchOpts, sink chan<- *OffRampWithMessageTransformerSourceChainSelectorAdded) (event.Subscription, error)

	ParseSourceChainSelectorAdded(log types.Log) (*OffRampWithMessageTransformerSourceChainSelectorAdded, error)

	FilterStaticConfigSet(opts *bind.FilterOpts) (*OffRampWithMessageTransformerStaticConfigSetIterator, error)

	WatchStaticConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampWithMessageTransformerStaticConfigSet) (event.Subscription, error)

	ParseStaticConfigSet(log types.Log) (*OffRampWithMessageTransformerStaticConfigSet, error)

	FilterTransmitted(opts *bind.FilterOpts, ocrPluginType []uint8) (*OffRampWithMessageTransformerTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *OffRampWithMessageTransformerTransmitted, ocrPluginType []uint8) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*OffRampWithMessageTransformerTransmitted, error)

	Address() common.Address
}
