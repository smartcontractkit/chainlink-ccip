// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package offramp_over_superchain_interop

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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structOffRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structOffRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"sourceChainConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structOffRamp.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isRMNVerificationDisabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"crossL2Inbox\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainSelectorToChainIdConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structOffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[]\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainSelectorToChainIdConfigUpdates\",\"inputs\":[{\"name\":\"chainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainSelectorsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structOffRampOverSuperchainInterop.ChainSelectorToChainIdConfigArgs[]\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySourceChainConfigUpdates\",\"inputs\":[{\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structOffRamp.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isRMNVerificationDisabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipReceive\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structClient.Any2EVMMessage\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"commit\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32[2]\",\"internalType\":\"bytes32[2]\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"reportContext\",\"type\":\"bytes32[2]\",\"internalType\":\"bytes32[2]\"},{\"name\":\"report\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSingleMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInternal.Any2EVMRampMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.RampMessageHeader\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Any2EVMTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasAmount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"offchainTokenData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"tokenGasOverrides\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllSourceChainConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOffRamp.SourceChainConfig[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"minSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isRMNVerificationDisabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChainId\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCrossL2Inbox\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOffRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExecutionState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumInternal.MessageExecutionState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestPriceSequenceNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMerkleRoot\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSourceChainConfig\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOffRamp.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"minSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isRMNVerificationDisabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOffRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestConfigDetails\",\"inputs\":[{\"name\":\"ocrPluginType\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"ocrConfig\",\"type\":\"tuple\",\"internalType\":\"structMultiOCR3Base.OCRConfig\",\"components\":[{\"name\":\"configInfo\",\"type\":\"tuple\",\"internalType\":\"structMultiOCR3Base.ConfigInfo\",\"components\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"F\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"n\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isSignatureVerificationEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"transmitters\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"manuallyExecute\",\"inputs\":[{\"name\":\"reports\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.ExecutionReport[]\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messages\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Any2EVMRampMessage[]\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.RampMessageHeader\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Any2EVMTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasAmount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"offchainTokenData\",\"type\":\"bytes[][]\",\"internalType\":\"bytes[][]\"},{\"name\":\"proofs\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"proofFlagBits\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"gasLimitOverrides\",\"type\":\"tuple[][]\",\"internalType\":\"structOffRamp.GasLimitOverride[][]\",\"components\":[{\"name\":\"receiverExecutionGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenGasOverrides\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structOffRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOCR3Configs\",\"inputs\":[{\"name\":\"ocrConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structMultiOCR3Base.OCRConfigArgs[]\",\"components\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ocrPluginType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"F\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isSignatureVerificationEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"transmitters\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AlreadyAttempted\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainSelectorToChainIdConfigRemoved\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainSelectorToChainIdConfigUpdated\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CommitReportAccepted\",\"inputs\":[{\"name\":\"blessedMerkleRoots\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structInternal.MerkleRoot[]\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"minSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"unblessedMerkleRoots\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structInternal.MerkleRoot[]\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"minSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"priceUpdates\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInternal.PriceUpdates\",\"components\":[{\"name\":\"tokenPriceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.TokenPriceUpdate[]\",\"components\":[{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"usdPerToken\",\"type\":\"uint224\",\"internalType\":\"uint224\"}]},{\"name\":\"gasPriceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.GasPriceUpdate[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"usdPerUnitGas\",\"type\":\"uint224\",\"internalType\":\"uint224\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"ocrPluginType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"configDigest\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"signers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"transmitters\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"F\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOffRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutionStateChanged\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"state\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"gasUsed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RootRemoved\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SkippedAlreadyExecutedMessage\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SkippedReportExecution\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SourceChainConfigSet\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sourceConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOffRamp.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"minSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isRMNVerificationDisabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SourceChainSelectorAdded\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOffRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transmitted\",\"inputs\":[{\"name\":\"ocrPluginType\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"configDigest\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySelfCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainIdMismatch\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expectedChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ChainIdNotConfiguredForSelector\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CommitOnRampMismatch\",\"inputs\":[{\"name\":\"reportOnRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"configOnRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ConfigDigestMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"CrossL2InboxCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"EmptyBatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyReport\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ExecutionError\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ForkedChain\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasToCompleteTx\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[{\"name\":\"errorType\",\"type\":\"uint8\",\"internalType\":\"enumMultiOCR3Base.InvalidConfigErrorType\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainSelector\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"expected\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidEncodingOfIdentifierInProofs\",\"inputs\":[{\"name\":\"proofs\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]},{\"type\":\"error\",\"name\":\"InvalidInterval\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"min\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"max\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidManualExecutionGasLimit\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"newLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidManualExecutionTokenGasOverride\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tokenIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"oldLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenGasOverride\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageDestChainSelector\",\"inputs\":[{\"name\":\"messageDestChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNewState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newState\",\"type\":\"uint8\",\"internalType\":\"enumInternal.MessageExecutionState\"}]},{\"type\":\"error\",\"name\":\"InvalidOnRampUpdate\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidProofsWordLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSourceChainSelector\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"expected\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidSourceOnRamp\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceOnRamp\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ManualExecutionGasAmountCountMismatch\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ManualExecutionGasLimitMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ManualExecutionNotYetEnabled\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MessageValidationError\",\"inputs\":[{\"name\":\"errorReason\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonUniqueSignatures\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotACompatiblePool\",\"inputs\":[{\"name\":\"notPool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperationNotSupportedByThisOffRampType\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OracleCannotBeZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProofFlagBitsMustBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiverError\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ReleaseOrMintBalanceMismatch\",\"inputs\":[{\"name\":\"amountReleased\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePre\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePost\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ReportMustContainExactlyOneMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RootAlreadyCommitted\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"RootBlessingMismatch\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"isBlessed\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"type\":\"error\",\"name\":\"RootNotCommitted\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SignatureVerificationNotAllowedInExecutionPlugin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureVerificationRequiredInCommitPlugin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignaturesOutOfRegistration\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SourceChainNotEnabled\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SourceChainSelectorMismatch\",\"inputs\":[{\"name\":\"reportSourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageSourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"StaleCommitReport\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StaticConfigCannotBeChanged\",\"inputs\":[{\"name\":\"ocrPluginType\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"TokenDataMismatch\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenHandlingError\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedTransmitter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnexpectedTokenData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongMessageLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"WrongNumberOfSignatures\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainIdNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainSelectorNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610160806040523461093257615ff9803803809161001d8285610a65565b833981018181039061016082126109325760a08212610932576040519161004383610a4a565b61004c84610a88565b8352602084015161ffff8116810361093257602084019081526040850151936001600160a01b0385168503610932576040810194855261008e60608701610a9c565b946060820195865260606100a460808901610a9c565b6080840190815294609f1901126109325760405192606084016001600160401b03811185821017610615576040526100de60a08901610a9c565b845260c08801519263ffffffff84168403610932576020850193845261010660e08a01610a9c565b604086019081526101008a01519096906001600160401b038111610932578a019888601f8b01121561093257895161013d81610ab0565b9a61014b6040519c8d610a65565b818c526020808d019260051b820101908b82116109325760208101925b82841061093757505050506101806101208c01610a9c565b6101408c0151909b6001600160401b038211610932570189601f82011215610932578051906101ae82610ab0565b9a6101bc6040519c8d610a65565b828c526020808d019360061b8301019181831161093257602001925b8284106108e7575050505033156108d657600180546001600160a01b031916331790554660805284516001600160a01b03161580156108c4575b80156108b2575b61062b5782516001600160401b0316156106c85782516001600160401b0390811660a090815286516001600160a01b0390811660c0528351811660e0528451811661010052865161ffff90811661012052604080519751909416875296519096166020860152955185169084015251831660608301525190911660808201527fb0fa1fb01508c5097c502ad056fd77018870c9be9a86d9e56b6b471862d7c5b79190a181516001600160a01b03161561062b57905160048054835163ffffffff60a01b60a09190911b166001600160a01b039384166001600160c01b03199092168217179091558351600580549184166001600160a01b031990921691909117905560408051918252925163ffffffff16602082015292511690820152839083907fa1c15688cb2c24508e158f6942b9276c6f3028a85e1af8cf3fff0c3ff3d5fc8d90606090a16000925b81518410156106d9576103778483610ad4565b5160208101519094906001600160401b031680156106c85785516001600160a01b03161561062b57806000526008602052604060002060808701519660018201906103c28254610afe565b610668578254600160a81b600160e81b031916600160a81b1783556040518481527ff4c1390c70e5c0f491ae1ccbc06f9117cbbadf2767b247b3bc203280f24c0fb990602090a15b8851801590811561063c575b5061062b5788516001600160401b038111610615576104358354610afe565b601f81116105cd575b50602099601f821160011461054957926060600080516020615fd98339815191529593836105329460019b9c9d9e60ff9860009261053e575b5050600019600383901b1c1916908b1b1783555b604081015115158554908760a01b9060a01b16908760a01b1916178555898060a01b038151168a8060a01b0319865416178555015115158354908560e81b9060e81b16908560e81b19161783556104e186610bbb565b506040519384936020855254898060a01b0381166020860152818160a01c1615156040860152898060401b038160a81c16606086015260e81c161515608084015260a08084015260c0830190610b38565b0390a201929190610364565b015190508f80610477565b99601f1982169a84600052816000209b60005b8181106105b5575093600184819b9c9d9e60ff989560609561053298600080516020615fd98339815191529c9a1061059c575b505050811b01835561048b565b015160001960f88460031b161c191690558f808061058f565b828401518e556001909d019c6020938401930161055c565b836000526020600020601f830160051c8101916020841061060b575b601f0160051c01905b8181106105ff575061043e565b600081556001016105f2565b90915081906105e9565b634e487b7160e01b600052604160045260246000fd5b6342bcdf7f60e11b60005260046000fd5b905060208a01206040516020810190600082526020815261065e604082610a65565b519020148a610416565b825460a81c6001600160401b0316600114158061069a575b1561040a5783632105803760e11b60005260045260246000fd5b506040516106b3816106ac8186610b38565b0382610a65565b60208151910120895160208b01201415610680565b63c656089560e01b60005260046000fd5b6001600160a01b03831680156108a157610140526020906040516106fd8382610a65565b60008152600036813760005b8151811015610781576001906001600160401b036107278285610ad4565b511680600052600c86526040600020549081610746575b505001610709565b80600052600c8752600060408120557fb56b587763154465d175e8a2a97978dffe45711125145973789b83ab201e702a600080a3858061073e565b505060005b8151811015610818576107998183610ad4565b51908382018051156108075782516001600160401b0316156106c8576001928151848060401b03825116600052600c8752604060002055838060401b039051169051907f5b4e1378b67677cfb7d6b50a37fd7f632168140928b0f970077a372bdfb42a3e600080a301610786565b63488d765160e01b60005260046000fd5b6040516153a09081610c39823960805181612dfa015260a0518181816101cf015261469c015260c0518181816102250152613863015260e0518181816102540152613f5d0152610100518181816102830152613bcc0152610120518181816101f6015281816123e7015281816140500152614df201526101405181818161050601526147740152f35b6303b3fcf960e61b60005260046000fd5b5081516001600160a01b031615610219565b5080516001600160a01b031615610212565b639b15e16f60e01b60005260046000fd5b604084830312610932576040805191908201906001600160401b0382118383101761061557604092602092845261091d87610a88565b815282870151838201528152019301926101d8565b600080fd5b83516001600160401b03811161093257820160a0818f03601f190112610932576040519061096482610a4a565b60208101516001600160a01b038116810361093257825261098760408201610a88565b602083015261099860608201610ac7565b60408301526109a960808201610ac7565b606083015260a0810151906001600160401b03821161093257016020810190603f018f13156109325780516001600160401b038111610615578f91604051926109fc6020601f19601f8601160185610a65565b828452602083830101116109325760005b828110610a355750509181600060208096949581960101526080820152815201930192610168565b80602080928401015182828701015201610a0d565b60a081019081106001600160401b0382111761061557604052565b601f909101601f19168101906001600160401b0382119082101761061557604052565b51906001600160401b038216820361093257565b51906001600160a01b038216820361093257565b6001600160401b0381116106155760051b60200190565b5190811515820361093257565b8051821015610ae85760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b90600182811c92168015610b2e575b6020831014610b1857565b634e487b7160e01b600052602260045260246000fd5b91607f1691610b0d565b60009291815491610b4883610afe565b8083529260018116908115610b9e5750600114610b6457505050565b60009081526020812093945091925b838310610b84575060209250010190565b600181602092949394548385870101520191019190610b73565b915050602093945060ff929192191683830152151560051b010190565b80600052600760205260406000205415600014610c325760065468010000000000000000811015610615576001810180600655811015610ae8577ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f0181905560065460009182526007602052604090912055600190565b5060009056fe6080604052600436101561001257600080fd5b60003560e01c806306285c6914610187578063181f5a77146101825780633242d5a91461017d5780633f4b04aa146101785780635215505b146101735780635e36480c1461016e5780635e7bb0081461016957806360987c201461016457806361ac8aac1461015f5780636f9e320f1461015a5780637437ff9f1461015557806379ba50971461015057806385572ffb1461014b5780638b6cecf8146101465780638da5cb5b14610141578063c673e5841461013c578063ccd37ba314610137578063cd19723714610132578063de5e0b9a1461012d578063e9d68a8e14610128578063f2fde38b14610123578063f58e03fc1461011e5763f716f99f1461011957600080fd5b611aa4565b611987565b6118fc565b611857565b6117b7565b611659565b6115fa565b611535565b61144d565b611413565b6113dd565b61135d565b6112bd565b611148565b611099565b610f97565b610d90565b6107b5565b610646565b61052a565b6104e6565b610460565b61019c565b600091031261019757565b600080fd5b34610197576000366003190112610197576101b5611bdf565b506102fd6040516101c581610317565b6001600160401b037f000000000000000000000000000000000000000000000000000000000000000016815261ffff7f00000000000000000000000000000000000000000000000000000000000000001660208201526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001660408201526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001660608201526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001660808201526040519182918291909160806001600160a01b038160a08401956001600160401b03815116855261ffff6020820151166020860152826040820151166040860152826060820151166060860152015116910152565b0390f35b634e487b7160e01b600052604160045260246000fd5b60a081019081106001600160401b0382111761033257604052565b610301565b604081019081106001600160401b0382111761033257604052565b606081019081106001600160401b0382111761033257604052565b608081019081106001600160401b0382111761033257604052565b90601f801991011681019081106001600160401b0382111761033257604052565b604051906103b860c083610388565b565b604051906103b860a083610388565b604051906103b861010083610388565b604051906103b8604083610388565b6001600160401b03811161033257601f01601f191660200190565b60405190610412602083610388565b60008252565b60005b83811061042b5750506000910152565b818101518382015260200161041b565b9060209161045481518092818552858086019101610418565b601f01601f1916010190565b34610197576000366003190112610197576102fd604051610482606082610388565b602681527f4f666652616d704f7665725375706572636861696e496e7465726f7020312e3660208201527f2e322d6465760000000000000000000000000000000000000000000000000000604082015260405191829160208352602083019061043b565b346101975760003660031901126101975760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101975760003660031901126101975760206001600160401b03600b5416604051908152f35b9060a060806105a2936001600160a01b0381511684526020810151151560208501526001600160401b036040820151166040850152606081015115156060850152015191816080820152019061043b565b90565b6040810160408252825180915260206060830193019060005b818110610627575050506020818303910152815180825260208201916020808360051b8301019401926000915b8383106105fa57505050505090565b9091929394602080610618600193601f198682030187528951610551565b970193019301919392906105eb565b82516001600160401b03168552602094850194909201916001016105be565b3461019757600036600319011261019757600654610663816107fd565b906106716040519283610388565b808252601f19610680826107fd565b0160005b81811061074257505061069681611c38565b9060005b8181106106b25750506102fd604051928392836105a5565b806106e86106d06106c4600194613728565b6001600160401b031690565b6106da8387611cd2565b906001600160401b03169052565b6107266107216107086106fb8488611cd2565b516001600160401b031690565b6001600160401b03166000526008602052604060002090565b611dbe565b6107308287611cd2565b5261073b8186611cd2565b500161069a565b60209061074d611c0a565b82828701015201610684565b600435906001600160401b038216820361019757565b35906001600160401b038216820361019757565b634e487b7160e01b600052602160045260246000fd5b600411156107a357565b610783565b9060048210156107a35752565b34610197576040366003190112610197576107ce610759565b602435906001600160401b0382168203610197576020916107ee91611e67565b6107fb60405180926107a8565bf35b6001600160401b0381116103325760051b60200190565b91908260a09103126101975760405161082c81610317565b6080610871818395803585526108446020820161076f565b60208601526108556040820161076f565b60408601526108666060820161076f565b60608601520161076f565b910152565b929192610882826103e8565b916108906040519384610388565b829481845281830111610197578281602093846000960137010152565b9080601f83011215610197578160206105a293359101610876565b6001600160a01b0381160361019757565b35906103b8826108c8565b63ffffffff81160361019757565b35906103b8826108e4565b81601f8201121561019757803590610914826107fd565b926109226040519485610388565b82845260208085019360051b830101918183116101975760208101935b83851061094e57505050505090565b84356001600160401b03811161019757820160a0818503601f190112610197576040519161097b83610317565b60208201356001600160401b0381116101975785602061099d928501016108ad565b835260408201356109ad816108c8565b60208401526109be606083016108f2565b60408401526080820135926001600160401b0384116101975760a0836109eb8860208098819801016108ad565b60608401520135608082015281520194019361093f565b9190916101408184031261019757610a186103a9565b92610a238183610814565b845260a08201356001600160401b0381116101975781610a449184016108ad565b602085015260c08201356001600160401b0381116101975781610a689184016108ad565b6040850152610a7960e083016108d9565b606085015261010082013560808501526101208201356001600160401b03811161019757610aa792016108fd565b60a0830152565b9080601f83011215610197578135610ac5816107fd565b92610ad36040519485610388565b81845260208085019260051b820101918383116101975760208201905b838210610aff57505050505090565b81356001600160401b03811161019757602091610b2187848094880101610a02565b815201910190610af0565b81601f8201121561019757803590610b43826107fd565b92610b516040519485610388565b82845260208085019360051b830101918183116101975760208101935b838510610b7d57505050505090565b84356001600160401b03811161019757820183603f82011215610197576020810135610ba8816107fd565b91610bb66040519384610388565b8183526020808085019360051b83010101918683116101975760408201905b838210610bef575050509082525060209485019401610b6e565b81356001600160401b03811161019757602091610c138a84808095890101016108ad565b815201910190610bd5565b9080601f83011215610197578135610c35816107fd565b92610c436040519485610388565b81845260208085019260051b82010192831161019757602001905b828210610c6b5750505090565b8135815260209182019101610c5e565b81601f8201121561019757803590610c92826107fd565b92610ca06040519485610388565b82845260208085019360051b830101918183116101975760208101935b838510610ccc57505050505090565b84356001600160401b03811161019757820160a0818503601f19011261019757610cf46103ba565b91610d016020830161076f565b835260408201356001600160401b03811161019757856020610d2592850101610aae565b602084015260608201356001600160401b03811161019757856020610d4c92850101610b2c565b60408401526080820135926001600160401b0384116101975760a083610d79886020809881980101610c1e565b606084015201356080820152815201940193610cbd565b34610197576040366003190112610197576004356001600160401b03811161019757610dc0903690600401610c7b565b6024356001600160401b038111610197573660238201121561019757806004013591610deb836107fd565b91610df96040519384610388565b8383526024602084019460051b820101903682116101975760248101945b828610610e2a57610e288585611eaf565b005b85356001600160401b03811161019757820136604382011215610197576024810135610e55816107fd565b91610e636040519384610388565b818352602060248185019360051b83010101903682116101975760448101925b828410610e9d575050509082525060209586019501610e17565b83356001600160401b038111610197576024908301016040601f1982360301126101975760405190610ece82610337565b6020810135825260408101356001600160401b03811161019757602091010136601f8201121561019757803590610f04826107fd565b91610f126040519384610388565b80835260208084019160051b8301019136831161019757602001905b828210610f4d5750505091816020938480940152815201930192610e83565b602080918335610f5c816108e4565b815201910190610f2e565b9181601f84011215610197578235916001600160401b038311610197576020808501948460051b01011161019757565b34610197576060366003190112610197576004356001600160401b03811161019757610fc7903690600401610a02565b6024356001600160401b03811161019757610fe6903690600401610f67565b91604435926001600160401b0384116101975761100a610e28943690600401610f67565b9390926122bf565b81601f8201121561019757803590611029826107fd565b926110376040519485610388565b82845260208085019360061b8301019181831161019757602001925b828410611061575050505090565b604084830312610197576020604091825161107b81610337565b6110848761076f565b81528287013583820152815201930192611053565b34610197576040366003190112610197576004356001600160401b0381116101975736602382011215610197578060040135906110d5826107fd565b916110e36040519384610388565b8083526024602084019160051b8301019136831161019757602401905b82821061113057602435846001600160401b0382116101975761112a610e28923690600401611012565b9061259c565b6020809161113d8461076f565b815201910190611100565b3461019757606036600319011261019757600060405161116781610352565b600435611173816108c8565b8152602435611181816108e4565b6020820190815260443590611195826108c8565b604083019182526111a46130b0565b6001600160a01b03835116156112ae57916112706001600160a01b036112a8937fa1c15688cb2c24508e158f6942b9276c6f3028a85e1af8cf3fff0c3ff3d5fc8d95611209838651166001600160a01b03166001600160a01b03196004541617600455565b517fffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffff77ffffffff00000000000000000000000000000000000000006004549260a01b1691161760045551166001600160a01b03166001600160a01b03196005541617600555565b6040519182918291909160406001600160a01b0381606084019582815116855263ffffffff6020820151166020860152015116910152565b0390a180f35b6342bcdf7f60e11b8452600484fd5b34610197576000366003190112610197576000604080516112dd81610352565b82815282602082015201526102fd6040516112f781610352565b63ffffffff6004546001600160a01b038116835260a01c1660208201526001600160a01b036005541660408201526040519182918291909160406001600160a01b0381606084019582815116855263ffffffff6020820151166020860152015116910152565b34610197576000366003190112610197576000546001600160a01b03811633036113cc576001600160a01b0319600154913382841617600155166000556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b63015aa1e360e11b60005260046000fd5b34610197576020366003190112610197576004356001600160401b0381116101975760a090600319903603011261019757600080fd5b34610197576020366003190112610197576001600160401b03611434610759565b16600052600c6020526020604060002054604051908152f35b346101975760003660031901126101975760206001600160a01b0360015416604051908152f35b6004359060ff8216820361019757565b359060ff8216820361019757565b906020808351928381520192019060005b8181106114b05750505090565b82516001600160a01b03168452602093840193909201916001016114a3565b906105a29160208152606082518051602084015260ff602082015116604084015260ff604082015116828401520151151560808201526040611520602084015160c060a085015260e0840190611492565b9201519060c0601f1982850301910152611492565b346101975760203660031901126101975760ff611550611474565b60606040805161155f81610352565b815161156a8161036d565b6000815260006020820152600083820152600084820152815282602082015201521660005260026020526102fd604060002060036115e9604051926115ae84610352565b6115b78161271f565b84526040516115d4816115cd8160028601612758565b0382610388565b60208501526115cd6040518094819301612758565b6040820152604051918291826114cf565b3461019757604036600319011261019757611613610759565b6001600160401b036024359116600052600a6020526040600020906000526020526020604060002054604051908152f35b8015150361019757565b35906103b882611644565b34610197576020366003190112610197576004356001600160401b038111610197573660238201121561019757806004013590611695826107fd565b906116a36040519283610388565b8282526024602083019360051b820101903682116101975760248101935b8285106116d157610e28846127af565b84356001600160401b03811161019757820160a0602319823603011261019757604051916116fe83610317565b602482013561170c816108c8565b835261171a6044830161076f565b6020840152606482013561172d81611644565b6040840152608482013561174081611644565b606084015260a4820135926001600160401b0384116101975761176d6020949360248695369201016108ad565b60808201528152019401936116c1565b9060049160441161019757565b9181601f84011215610197578235916001600160401b038311610197576020838186019501011161019757565b346101975760c0366003190112610197576117d13661177d565b506044356001600160401b038111610197576117f190369060040161178a565b50506064356001600160401b03811161019757611812903690600401610f67565b50506084356001600160401b03811161019757611833903690600401610f67565b505063c4d744db60e01b60005260046000fd5b9060206105a2928181520190610551565b34610197576020366003190112610197576001600160401b03611878610759565b611880611c0a565b501660005260086020526102fd60406000206118eb6001604051926118a484610317565b6118e560ff82546001600160a01b0381168752818160a01c16151560208801526001600160401b038160a81c16604088015260e81c16606086019015159052565b01611da3565b608082015260405191829182611846565b34610197576020366003190112610197576001600160a01b03600435611921816108c8565b6119296130b0565b1633811461197657806001600160a01b031960005416176000556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b636d6c4ee560e11b60005260046000fd5b34610197576060366003190112610197576119a13661177d565b6044356001600160401b038111610197576119c090369060040161178a565b91828201602083820312610197578235906001600160401b038211610197576119ea918401610c7b565b6040519060206119fa8184610388565b60008352601f19810160005b818110611a2e57505050610e289491611a1e91612e3b565b611a26612a8d565b92839261332f565b60608582018401528201611a06565b9080601f83011215610197578135611a54816107fd565b92611a626040519485610388565b81845260208085019260051b82010192831161019757602001905b828210611a8a5750505090565b602080918335611a99816108c8565b815201910190611a7d565b34610197576020366003190112610197576004356001600160401b038111610197573660238201121561019757806004013590611ae0826107fd565b90611aee6040519283610388565b8282526024602083019360051b820101903682116101975760248101935b828510611b1c57610e2884612acb565b84356001600160401b03811161019757820160c0602319823603011261019757611b446103a9565b9160248201358352611b5860448301611484565b6020840152611b6960648301611484565b6040840152611b7a6084830161164e565b606084015260a48201356001600160401b03811161019757611ba29060243691850101611a3d565b608084015260c4820135926001600160401b03841161019757611bcf602094936024869536920101611a3d565b60a0820152815201940193611b0c565b60405190611bec82610317565b60006080838281528260208201528260408201528260608201520152565b60405190611c1782610317565b60606080836000815260006020820152600060408201526000838201520152565b90611c42826107fd565b611c4f6040519182610388565b8281528092611c60601f19916107fd565b0190602036910137565b634e487b7160e01b600052603260045260246000fd5b805115611c8d5760200190565b611c6a565b805160011015611c8d5760400190565b805160031015611c8d5760800190565b805160041015611c8d5760a00190565b805160021015611c8d5760600190565b8051821015611c8d5760209160051b010190565b90600182811c92168015611d16575b6020831014611d0057565b634e487b7160e01b600052602260045260246000fd5b91607f1691611cf5565b60009291815491611d3083611ce6565b8083529260018116908115611d865750600114611d4c57505050565b60009081526020812093945091925b838310611d6c575060209250010190565b600181602092949394548385870101520191019190611d5b565b915050602093945060ff929192191683830152151560051b010190565b906103b8611db79260405193848092611d20565b0383610388565b9060016080604051611dcf81610317565b611e25819560ff81546001600160a01b0381168552818160a01c16151560208601526001600160401b038160a81c16604086015260e81c1615156060840152611e1e6040518096819301611d20565b0384610388565b0152565b634e487b7160e01b600052601160045260246000fd5b908160051b9180830460201490151715611e5557565b611e29565b91908203918211611e5557565b611e7382607f92612db4565b9116906801fffffffffffffffe6001600160401b0383169260011b169180830460021490151715611e55576003911c1660048110156107a35790565b611eb7612df8565b8051825181036120b25760005b818110611ed7575050906103b891612e3b565b611ee18184611cd2565b516020810190815151611ef48488611cd2565b5192835182036120b25790916000925b808410611f18575050505050600101611ec4565b91949398611f2a848b98939598611cd2565b515198611f38888851611cd2565b519980612069575b5060a08a01988b6020611f568b8d515193611cd2565b51015151036120285760005b8a515181101561201357611f9e611f95611f8b8f6020611f838f8793611cd2565b510151611cd2565b5163ffffffff1690565b63ffffffff1690565b8b81611faf575b5050600101611f62565b611f956040611fc285611fce9451611cd2565b51015163ffffffff1690565b90818110611fdd57508b611fa5565b8d51516040516348e617b360e01b81526004810191909152602481019390935260448301919091526064820152608490fd5b0390fd5b50985098509893949095600101929091611f04565b6120658b51612043606082519201516001600160401b031690565b6370a193fd60e01b6000526004919091526001600160401b0316602452604490565b6000fd5b60808b0151811015611f4057612065908b61208b88516001600160401b031690565b905151633a98d46360e11b6000526001600160401b03909116600452602452604452606490565b6320f8fd5960e21b60005260046000fd5b604051906120d082610337565b60006020838281520152565b604051906120eb602083610388565b600080835282815b8281106120ff57505050565b60209061210a6120c3565b828285010152016120f3565b805182526001600160401b036020820151166020830152608061215d61214b604084015160a0604087015260a086019061043b565b6060840151858203606087015261043b565b9101519160808183039101526020808351928381520192019060005b8181106121865750505090565b825180516001600160a01b031685526020908101518186015260409094019390920191600101612179565b9060206105a2928181520190612116565b6040513d6000823e3d90fd5b3d156121f9573d906121df826103e8565b916121ed6040519384610388565b82523d6000602084013e565b606090565b9060206105a292818152019061043b565b909160608284031261019757815161222681611644565b9260208301516001600160401b0381116101975783019080601f8301121561019757815191612254836103e8565b916122626040519384610388565b83835260208483010111610197576040926122839160208085019101610418565b92015190565b9293606092959461ffff6122ad6001600160a01b0394608088526080880190612116565b97166020860152604085015216910152565b9290939130330361258b576122d26120dc565b9460a08501518051612544575b50505050508051916122fd602084519401516001600160401b031690565b90602083015191604084019261232a8451926123176103ba565b9788526001600160401b03166020880152565b6040860152606085015260808401526001600160a01b036123536005546001600160a01b031690565b16806124c7575b50515115806124bb575b80156124a5575b801561247c575b6124785761241091816123b56123a961239c610708602060009751016001600160401b0390511690565b546001600160a01b031690565b6001600160a01b031690565b90836123d0606060808401519301516001600160a01b031690565b604051633cf9798360e01b815296879586948593917f00000000000000000000000000000000000000000000000000000000000000009060048601612289565b03925af19081156124735760009060009261244c575b501561242f5750565b6040516302a35ba360e21b815290819061200f90600483016121fe565b905061246b91503d806000833e6124638183610388565b81019061220f565b509038612426565b6121c2565b5050565b506124a061249c61249760608401516001600160a01b031690565b613062565b1590565b612372565b5060608101516001600160a01b03163b1561236b565b50608081015115612364565b803b1561019757600060405180926308d450a160e01b82528183816124ef8a600483016121b1565b03925af19081612529575b506125235761200f61250a6121ce565b6040516309c2532560e01b8152918291600483016121fe565b3861235a565b80612538600061253e93610388565b8061018c565b386124fa565b859650602061258096015161256360608901516001600160a01b031690565b9061257a60208a51016001600160401b0390511690565b92612f49565b9038808080806122df565b6306e34e6560e31b60005260046000fd5b916125a56130b0565b60005b835181101561264457806125c16106fb60019387611cd2565b6125de816001600160401b0316600052600c602052604060002090565b5490816125ee575b5050016125a8565b8060006126166001600160401b03936001600160401b0316600052600c602052604060002090565b55167fb56b587763154465d175e8a2a97978dffe45711125145973789b83ab201e702a600080a338806125e6565b50915060005b81518110156124785761265d8183611cd2565b51906020820180511561270e5761267e6106c484516001600160401b031690565b156126fd576001600160401b036126ce60019483516126c06126a783516001600160401b031690565b6001600160401b0316600052600c602052604060002090565b55516001600160401b031690565b915191167f5b4e1378b67677cfb7d6b50a37fd7f632168140928b0f970077a372bdfb42a3e600080a30161264a565b63c656089560e01b60005260046000fd5b63488d765160e01b60005260046000fd5b9060405161272c8161036d565b606060ff600183958054855201548181166020850152818160081c16604085015260101c161515910152565b906020825491828152019160005260206000209060005b81811061277c5750505090565b82546001600160a01b031684526020909301926001928301920161276f565b906103b8611db79260405193848092612758565b6127b76130b0565b60005b8151811015612478576127cd8183611cd2565b51906127e360208301516001600160401b031690565b6001600160401b0381169081156126fd5761280b6123a96123a986516001600160a01b031690565b156129f85761282d816001600160401b03166000526008602052604060002090565b60808501519060018101926128428454611ce6565b612a1f576128c97ff4c1390c70e5c0f491ae1ccbc06f9117cbbadf2767b247b3bc203280f24c0fb9916128af8475010000000000000000000000000000000000000000007fffffff0000000000000000ffffffffffffffffffffffffffffffffffffffffff825416179055565b6040516001600160401b0390911681529081906020820190565b0390a15b81518015908115612a09575b506129f8576129d96129a460606001986129176129ef967fbd1ab25a0ff0a36a588597ba1af11e30f3f210de8b9e818cc9bbc457c94c8d8c98613152565b61296d6129276040830151151590565b86547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178655565b61299d61298182516001600160a01b031690565b86906001600160a01b03166001600160a01b0319825416179055565b0151151590565b82547fffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690151560e81b60ff60e81b16178255565b6129e284614f05565b5060405191829182613223565b0390a2016127ba565b6342bcdf7f60e11b60005260046000fd5b90506020830120612a186130d5565b14386128d9565b60016001600160401b03612a3e84546001600160401b039060a81c1690565b16141580612a6e575b612a5157506128cd565b632105803760e11b6000526001600160401b031660045260246000fd5b50612a7884611da3565b60208151910120835160208501201415612a47565b60405190612a9c602083610388565b6000808352366020840137565b60408051909190612aba8382610388565b6001815291601f1901366020840137565b612ad36130b0565b60005b815181101561247857612ae98183611cd2565b51906040820160ff612afc825160ff1690565b1615612d9e57602083015160ff1692612b228460ff166000526002602052604060002090565b9160018301918254612b3d612b378260ff1690565b60ff1690565b612d635750612b6a612b526060830151151590565b845462ff0000191690151560101b62ff000016178455565b60a08101918251610100815111612d0b57805115612d4d5760038601612b98612b928261279b565b8a6143d7565b6060840151612c28575b947fab8b1b57514019638d7b5ce9c638fe71366fe8e2be1c40a7a80f1733d0e9f54794600294612c04612bf4612c229a96612bed8760019f9c612be8612c1a9a8f614545565b61362b565b5160ff1690565b845460ff191660ff821617909455565b51908185555190604051958695019088866136b1565b0390a16145c7565b01612ad6565b97946002879395970196612c44612c3e8961279b565b886143d7565b608085015194610100865111612d37578551612c6c612b37612c678a5160ff1690565b613617565b1015612d21578551845111612d0b57612c04612bf47fab8b1b57514019638d7b5ce9c638fe71366fe8e2be1c40a7a80f1733d0e9f54798612bed8760019f612be8612c229f9a8f612cf360029f612ced612c1a9f8f90612be88492612cd2845160ff1690565b908054909161ff001990911660089190911b61ff0016179055565b8261446b565b505050979c9f50975050969a50505094509450612ba2565b631b3fab5160e11b600052600160045260246000fd5b631b3fab5160e11b600052600360045260246000fd5b631b3fab5160e11b600052600260045260246000fd5b631b3fab5160e11b600052600560045260246000fd5b60101c60ff16612d7e612d796060840151151590565b151590565b90151514612b6a576321fd80df60e21b60005260ff861660045260246000fd5b631b3fab5160e11b600090815260045260246000fd5b906001600160401b03612df4921660005260096020526701ffffffffffffff60406000209160071c166001600160401b0316600052602052604060002090565b5490565b7f0000000000000000000000000000000000000000000000000000000000000000468103612e235750565b630f01ce8560e01b6000526004524660245260446000fd5b919091805115612edd578251159260209160405192612e5a8185610388565b60008452601f19810160005b818110612eb95750505060005b8151811015612eb15780612e9a612e8c60019385611cd2565b518815612ea0578690613805565b01612e73565b612eaa8387611cd2565b5190613805565b505050509050565b8290604051612ec781610337565b6000815260608382015282828901015201612e66565b63c2e5347d60e01b60005260046000fd5b9190811015611c8d5760051b0190565b356105a2816108e4565b9190811015611c8d5760051b81013590601e19813603018212156101975701908135916001600160401b038311610197576020018236038113610197579190565b90929491939796815196612f5c886107fd565b97612f6a604051998a610388565b808952612f79601f19916107fd565b0160005b81811061304b57505060005b835181101561303e5780612fd08c8a8a8a612fca612fc3878d612fbc828f8f9d8f9e60019f81612fec575b505050611cd2565b5197612f08565b3691610876565b93613f0e565b612fda828c611cd2565b52612fe5818b611cd2565b5001612f89565b63ffffffff613004612fff858585612eee565b612efe565b1615612fb4576130349261301b92612fff92612eee565b60406130278585611cd2565b51019063ffffffff169052565b8f8f908391612fb4565b5096985050505050505050565b6020906130566120c3565b82828d01015201612f7d565b6130736385572ffb60e01b82614271565b908161308d575b81613083575090565b6105a29150614243565b9050613098816141c8565b159061307a565b61307363aff2afbf60e01b82614271565b6001600160a01b036001541633036130c457565b6315ae3a6f60e11b60005260046000fd5b604051602081019060008252602081526130f0604082610388565b51902090565b818110613101575050565b600081556001016130f6565b9190601f811161311c57505050565b6103b8926000526020600020906020601f840160051c83019310613148575b601f0160051c01906130f6565b909150819061313b565b91909182516001600160401b03811161033257613179816131738454611ce6565b8461310d565b6020601f82116001146131ba5781906131ab9394956000926131af575b50508160011b916000199060031b1c19161790565b9055565b015190503880613196565b601f198216906131cf84600052602060002090565b9160005b81811061320b575095836001959697106131f2575b505050811b019055565b015160001960f88460031b161c191690553880806131e8565b9192602060018192868b0151815501940192016131d3565b90600160c06105a2936020815260ff84546001600160a01b0381166020840152818160a01c16151560408401526001600160401b038160a81c16606084015260e81c161515608082015260a080820152019101611d20565b6084019081608411611e5557565b60a001908160a011611e5557565b91908201809211611e5557565b600311156107a357565b60038210156107a35752565b906103b86040516132ca81610337565b602060ff829554818116845260081c1691016132ae565b8054821015611c8d5760005260206000200190600090565b60ff60019116019060ff8211611e5557565b60ff601b9116019060ff8211611e5557565b90606092604091835260208301370190565b60016000526002602052936133637fe90b7bceb6e7df5418fb78d8ee546e97c83a08bbccc01a0644d599ccd2a7c2e061271f565b938535946133708561327b565b606082019061337f8251151590565b6135e9575b8036036135d1575081518781036135b8575061339e612df8565b3360009081527fa15bc60c955c405d20d9149c709e2460f1c2d9a497496a7f46004d1772c3054c602052604090206133da906132ba565b6132ba565b600260208201516133ea816132a4565b6133f3816132a4565b14908161355c575b5015613530575b51613467575b50505050507f198d6990ef96613a9026203077e422916918b03ff47f0be6bee7b02d8e139ef09061344b61343e60019460200190565b356001600160401b031690565b604080519283526001600160401b0391909116602083015290a2565b613488612b37613483602085969799989a955194015160ff1690565b6132f9565b0361351f57815183510361350e57613506600061344b9461343e946134d27f198d6990ef96613a9026203077e422916918b03ff47f0be6bee7b02d8e139ef09960019b3691610876565b602081519101206040516134fd816134ef8960208301958661331d565b03601f198101835282610388565b5190208a6142a1565b948394613408565b63a75d88af60e01b60005260046000fd5b6371253a2560e01b60005260046000fd5b72c11c11c11c11c11c11c11c11c11c11c11c11c133031561340257631b41e11d60e31b60005260046000fd5b60016000526002602052516135b091506123a99061359d9060ff167fe90b7bceb6e7df5418fb78d8ee546e97c83a08bbccc01a0644d599ccd2a7c2e36132e1565b90546001600160a01b039160031b1c1690565b3314386133fb565b6324f7d61360e21b600052600452602487905260446000fd5b638e1192e160e01b6000526004523660245260446000fd5b6136129061360c6136026135fd8751611e3f565b613289565b61360c8851611e3f565b90613297565b613384565b60ff166003029060ff8216918203611e5557565b8151916001600160401b03831161033257680100000000000000008311610332576020908254848455808510613694575b500190600052602060002060005b8381106136775750505050565b60019060206001600160a01b03855116940193818401550161366a565b6136ab9084600052858460002091820191016130f6565b3861365c565b95949392909160ff6136d693168752602087015260a0604087015260a0860190612758565b84810360608601526020808351928381520192019060005b818110613709575050509060806103b89294019060ff169052565b82516001600160a01b03168452602093840193909201916001016136ee565b600654811015611c8d5760066000527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f015490565b9081602091031261019757516105a281611644565b6001600160401b036105a2949381606094168352166020820152816040820152019061043b565b6040906105a293928152816020820152019061043b565b9291906001600160401b039081606495166004521660245260048110156107a357604452565b9493926137ef60609361380093885260208801906107a8565b60806040870152608086019061043b565b930152565b9061381782516001600160401b031690565b8151604051632cbc26bb60e01b815277ffffffffffffffff00000000000000000000000000000000608084901b1660048201529015159391906001600160401b038216906020816024817f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03165afa90811561247357600091613df7575b50613d98576020830191825151948515613d685760408501948551518703613d57576138c89083614636565b95909760005b8881106138e15750505050505050505050565b5a6138f66138f0838a51611cd2565b51614911565b80516060015161390f906001600160401b031688611e67565b61391881610799565b8015908d8283159384613d44575b15613d015760608815613c84575061394d6020613943898d611cd2565b5101519242611e5a565b6004546139629060a01c63ffffffff16611f95565b108015613c71575b15613c5357613979878b611cd2565b5151613c3d575b845160800151613998906001600160401b03166106c4565b613b75575b506139a9868951611cd2565b5160a085015151815103613b395793613a0e9695938c938f966139ee8e958c926139e86139e260608951016001600160401b0390511690565b8961495b565b86614ca1565b9a908096613a0860608851016001600160401b0390511690565b906149e3565b613ae7575b5050613a1e82610799565b60028203613a9f575b600196613a957f05665fe9ad095383d018353f4cbcba77e84db27dd215081bbf7cdf9ae6fbe48b936001600160401b03935192613a86613a7d8b613a7560608801516001600160401b031690565b96519b611cd2565b51985a90611e5a565b916040519586951698856137d6565b0390a45b016138ce565b91509193949250613aaf82610799565b60038203613ac3578b929493918a91613a27565b51606001516349362d1f60e11b60005261206591906001600160401b0316896137b0565b613af084610799565b60038403613a13579092949550613b08919350610799565b613b18578b92918a913880613a13565b5151604051632b11b8d960e01b815290819061200f90879060048401613799565b6120658b613b5360608851016001600160401b0390511690565b631cfe6d8b60e01b6000526001600160401b0391821660045216602452604490565b613b7e83610799565b613b89575b3861399d565b8351608001516001600160401b0316602080860151918c613bbe60405194859384936370701e5760e11b855260048501613772565b038160006001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165af190811561247357600091613c0f575b50613b83575050505050600190613a99565b613c30915060203d8111613c36575b613c288183610388565b81019061375d565b38613bfd565b503d613c1e565b613c47878b611cd2565b51516080860152613980565b6354e7e43160e11b6000526001600160401b038b1660045260246000fd5b50613c7b83610799565b6003831461396a565b915083613c9084610799565b1561398057506001959450613cf99250613cd791507f3ef2a99c550a751d4b0b261268f05a803dfb049ab43616a1ffb388f61fe651209351016001600160401b0390511690565b604080516001600160401b03808c168252909216602083015290918291820190565b0390a1613a99565b505050506001929150613cf9613cd760607f3b575419319662b2a6f5e2467d84521517a3382b908eb3d557bb3fdb0c50e23c9351016001600160401b0390511690565b50613d4e83610799565b60038314613926565b6357e0e08360e01b60005260046000fd5b612065613d7c86516001600160401b031690565b63676cf24b60e11b6000526001600160401b0316600452602490565b5092915050613dda576040516001600160401b039190911681527faab522ed53d887e56ed53dd37398a01aeef6a58e0fa77c2173beb9512d89493390602090a1565b637edeb53960e11b6000526001600160401b031660045260246000fd5b613e10915060203d602011613c3657613c288183610388565b3861389c565b9081602091031261019757516105a2816108c8565b906105a2916020815260e0613ec9613eb4613e548551610100602087015261012086019061043b565b60208601516001600160401b0316604086015260408601516001600160a01b0316606086015260608601516080860152613e9e608087015160a08701906001600160a01b03169052565b60a0860151858203601f190160c087015261043b565b60c0850151848203601f19018486015261043b565b92015190610100601f198285030191015261043b565b6040906001600160a01b036105a29493168152816020820152019061043b565b90816020910312610197575190565b91939293613f1a6120c3565b5060208301516001600160a01b031660405163bbe4f6db60e01b81526001600160a01b038216600482015290959092602084806024810103816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa93841561247357600094614197575b506001600160a01b0384169586158015614185575b6141675761404c614075926134ef92613fd0613fc9611f9560408c015163ffffffff1690565b8c89614dba565b9690996080810151613ffe6060835193015193613feb6103c9565b9687526001600160401b03166020870152565b6001600160a01b038a16604086015260608501526001600160a01b038d16608085015260a084015260c083015260e0820152604051633907753760e01b602082015292839160248301613e2b565b82857f000000000000000000000000000000000000000000000000000000000000000092614e48565b9490911561414b575080516020810361413257509061409e826020808a95518301019101613eff565b956001600160a01b038416036140d6575b50505050506140ce6140bf6103d9565b6001600160a01b039093168352565b602082015290565b6140e9936140e391611e5a565b91614dba565b5090808210801561411f575b614101578084816140af565b63a966e21f60e01b6000908152600493909352602452604452606490fd5b508261412b8284611e5a565b14156140f5565b631e3be00960e21b600052602060045260245260446000fd5b61200f604051928392634ff17cad60e11b845260048401613edf565b63ae9b4ce960e01b6000526001600160a01b03851660045260246000fd5b5061419261249c8661309f565b613fa3565b6141ba91945060203d6020116141c1575b6141b28183610388565b810190613e16565b9238613f8e565b503d6141a8565b60405160208101916301ffc9a760e01b835263ffffffff60e01b6024830152602482526141f6604483610388565b6179185a10614232576020926000925191617530fa6000513d82614226575b508161421f575090565b9050151590565b60201115915038614215565b63753fa58960e11b60005260046000fd5b60405160208101916301ffc9a760e01b83526301ffc9a760e01b6024830152602482526141f6604483610388565b6040519060208201926301ffc9a760e01b845263ffffffff60e01b166024830152602482526141f6604483610388565b919390926000948051946000965b8688106142c0575050505050505050565b6020881015611c8d57602060006142d8878b1a61330b565b6142e28b87611cd2565b51906143196142f18d8a611cd2565b5160405193849389859094939260ff6060936080840197845216602083015260408201520152565b838052039060015afa156124735761435f6133d56000516143478960ff166000526003602052604060002090565b906001600160a01b0316600052602052604060002090565b9060016020830151614370816132a4565b614379816132a4565b036143c65761439661438c835160ff1690565b60ff600191161b90565b81166143b5576143ac61438c6001935160ff1690565b179701966142af565b633d9ef1f160e21b60005260046000fd5b636518c33d60e11b60005260046000fd5b91909160005b83518110156144305760019060ff831660005260036020526000614429604082206001600160a01b03614410858a611cd2565b51166001600160a01b0316600052602052604060002090565b55016143dd565b50509050565b8151815460ff191660ff91909116178155906020015160038110156107a357815461ff00191660089190911b61ff0016179055565b919060005b8151811015614430576144936144868284611cd2565b516001600160a01b031690565b906144bc6144b2836143478860ff166000526003602052604060002090565b5460081c60ff1690565b6144c5816132a4565b614530576001600160a01b0382161561451f576145196001926145146144e96103d9565b60ff85168152916144fd86602085016132ae565b6143478960ff166000526003602052604060002090565b614436565b01614470565b63d6c62c9b60e01b60005260046000fd5b631b3fab5160e11b6000526004805260246000fd5b919060005b8151811015614430576145606144868284611cd2565b9061457f6144b2836143478860ff166000526003602052604060002090565b614588816132a4565b614530576001600160a01b0382161561451f576145c16001926145146145ac6103d9565b60ff85168152916144fd6002602085016132ae565b0161454a565b60ff1680600052600260205260ff60016040600020015460101c16908015600014614615575015614604576001600160401b0319600b5416600b55565b6317bd8dd160e11b60005260046000fd5b60011461461f5750565b61462557565b6307b8c74d60e51b60005260046000fd5b91909160208301600181515103614900576146519051611c80565b51926146686020855101516001600160401b031690565b6001600160401b0383166001600160401b038216036148dc575060808101516148cb578351604001516001600160401b03167f0000000000000000000000000000000000000000000000000000000000000000906001600160401b0382166001600160401b038216036148a857505060606146e591015184615037565b916147096123a96146fa60016118e5856151a8565b60208082518301019101613e16565b9061471b83516001600160a01b031690565b6001600160a01b0383166001600160a01b0382160361487d5750614752816001600160401b0316600052600c602052604060002090565b549081156148605760808401519081830361483d575050506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016803b15610197576040805163ab4d6f7560e01b815284516001600160a01b0316600482015260208501516024820152908401516044820152606084015160648201526080840151608482015260a4810194909452600090849060c490829084905af19081156124735760609361481892614828575b50614812612aa9565b956151e5565b61482185611c80565b5201519190565b80612538600061483793610388565b38614809565b63f64ea74360e01b6000526001600160401b031660045260245260445260646000fd5b63bd04062360e01b6000526001600160401b031660045260246000fd5b636db7ad9160e01b6000526001600160401b039091166004526001600160a01b031660245260446000fd5b632f0ccf7760e01b6000526001600160401b039081166004521660245260446000fd5b6322a4f36560e11b60005260046000fd5b633c3e620760e11b6000526001600160401b03908116600452821660245260446000fd5b630fc5267f60e11b60005260046000fd5b60405160c081018181106001600160401b038211176103325760609160a09160405261493b611bdf565b815282602082015282604082015260008382015260006080820152015290565b607f8216906801fffffffffffffffe6001600160401b0383169260011b169180830460021490151715611e55576149e0916001600160401b0361499e8584612db4565b921660005260096020526701ffffffffffffff60406000209460071c169160036001831b921b19161792906001600160401b0316600052602052604060002090565b55565b9091607f83166801fffffffffffffffe6001600160401b0382169160011b169080820460021490151715611e5557614a1b8484612db4565b60048310156107a3576001600160401b036149e09416600052600960205260036701ffffffffffffff60406000209660071c1693831b921b19161792906001600160401b0316600052602052604060002090565b9080602083519182815201916020808360051b8301019401926000915b838310614a9b57505050505090565b9091929394602080600192601f19858203018652885190608080614afe614acb855160a0865260a086019061043b565b6001600160a01b0387870151168786015263ffffffff60408701511660408601526060860151858203606087015261043b565b93015191015297019301930191939290614a8c565b6105a2916001600160401b036080835180518452826020820151166020850152826040820151166040850152826060820151166060850152015116608082015260a0614b84614b7260208501516101408486015261014085019061043b565b604085015184820360c086015261043b565b60608401516001600160a01b031660e0840152926080810151610100840152015190610120818403910152614a6f565b90614bc790606083526060830190614b13565b8181036020830152825180825260208201916020808360051b8301019501926000915b838310614c3757505050505060408183039101526020808351928381520192019060005b818110614c1b5750505090565b825163ffffffff16845260209384019390920191600101614c0e565b9091929395602080614c55600193601f198682030187528a5161043b565b98019301930191939290614bea565b80516020909101516001600160e01b0319811692919060048210614c86575050565b6001600160e01b031960049290920360031b82901b16169150565b90303b1561019757600091614cca6040519485938493630304c3e160e51b855260048501614bb4565b038183305af19081614da5575b50614d9a57614ce46121ce565b9072c11c11c11c11c11c11c11c11c11c11c11c11c13314614d06575b60039190565b614d1f614d1283614c64565b6001600160e01b03191690565b6337c3be2960e01b148015614d7f575b8015614d64575b15614d0057612065614d4783614c64565b632882569d60e01b6000526001600160e01b031916600452602490565b50614d71614d1283614c64565b63753fa58960e11b14614d36565b50614d8c614d1283614c64565b632be8ca8b60e21b14614d2f565b6002906105a2610403565b806125386000614db493610388565b38614cd7565b6040516370a0823160e01b60208201526001600160a01b039091166024820152919291614e1790614dee81604481016134ef565b84837f000000000000000000000000000000000000000000000000000000000000000092614e48565b9290911561414b5750805160208103614132575090614e42826020806105a295518301019101613eff565b93611e5a565b939193614e5560846103e8565b94614e636040519687610388565b60848652614e7160846103e8565b602087019590601f1901368737833b15614ef4575a90808210614ee3578291038060061c90031115614ed2576000918291825a9560208451940192f1905a9003923d9060848211614ec9575b6000908287523e929190565b60849150614ebd565b6337c3be2960e01b60005260046000fd5b632be8ca8b60e21b60005260046000fd5b63030ed58f60e21b60005260046000fd5b80600052600760205260406000205415600014614f83576006546801000000000000000081101561033257600181016006556000600654821015611c8d57600690527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f01819055600654906000526007602052604060002055600190565b50600090565b602060408183019282815284518094520192019060005b818110614fad5750505090565b8251845260209384019390920191600101614fa0565b9060206105a2928181520190614b13565b6080906103b89294936040519586927fb056c48de8fce43e6c30818e5f9c6a56a7014ef5944b6329ed5c76afff7e838a6020850152604084015260608301526150268151809260208686019101610418565b81010301601f198101845283610388565b919091615042611bdf565b5082516005810361518f575061505783611c80565b5115801561517e575b801561516d575b801561515c575b61514157615124615137916150886123a961448687611c80565b9461509281611c92565b519061509d81611cc2565b516150b16150aa83611ca2565b5192611cb2565b51926150cd6150be6103ba565b6001600160a01b03909a168a52565b602089015260408801526060870152608086015280516151326151146106c460606151056106c460408701516001600160401b031690565b9401516001600160401b031690565b9260405194859160208301614fc3565b03601f198101855284610388565b614fd4565b6020815191012090565b604051635649411d60e01b81528061200f8560048301614f89565b5061516683611cb2565b511561506e565b5061517783611ca2565b5115615067565b5061518883611c92565b5115615060565b630608814160e41b600052600452600560245260446000fd5b6001600160401b031680600052600860205260406000209060ff825460a01c16156151d1575090565b63ed053c5960e01b60005260045260246000fd5b6105a2918151906001600160401b0360408160208501511693015116906040516001600160a01b03602082019216825260208152615224604082610388565b5190206040519160208301937f2425b0b9f9054c76ff151b0a175b18f37a4a4e82013a72e9f15c9caa095ed21f85526040840152606083015260808201526080815261527160a082610388565b5190209061528b565b9060206105a2928181520190614a6f565b6130f0815180519061531f6152aa60608601516001600160a01b031690565b6134ef6152c160608501516001600160401b031690565b936152da6080808a01519201516001600160401b031690565b90604051958694602086019889936001600160401b036080946001600160a01b0382959998949960a089019a8952166020880152166040860152606085015216910152565b5190206134ef6020840151602081519101209360a0604082015160208151910120910151604051615358816134ef60208201948561527a565b51902090604051958694602086019889919260a093969594919660c0840197600085526020850152604084015260608301526080820152015256fea164736f6c634300081a000abd1ab25a0ff0a36a588597ba1af11e30f3f210de8b9e818cc9bbc457c94c8d8c",
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

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropCaller) GetChainId(opts *bind.CallOpts, chainSelector uint64) (*big.Int, error) {
	var out []interface{}
	err := _OffRampOverSuperchainInterop.contract.Call(opts, &out, "getChainId", chainSelector)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropSession) GetChainId(chainSelector uint64) (*big.Int, error) {
	return _OffRampOverSuperchainInterop.Contract.GetChainId(&_OffRampOverSuperchainInterop.CallOpts, chainSelector)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropCallerSession) GetChainId(chainSelector uint64) (*big.Int, error) {
	return _OffRampOverSuperchainInterop.Contract.GetChainId(&_OffRampOverSuperchainInterop.CallOpts, chainSelector)
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

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropTransactor) ApplyChainSelectorToChainIdConfigUpdates(opts *bind.TransactOpts, chainSelectorsToRemove []uint64, chainSelectorsToAdd []OffRampOverSuperchainInteropChainSelectorToChainIdConfigArgs) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.contract.Transact(opts, "applyChainSelectorToChainIdConfigUpdates", chainSelectorsToRemove, chainSelectorsToAdd)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropSession) ApplyChainSelectorToChainIdConfigUpdates(chainSelectorsToRemove []uint64, chainSelectorsToAdd []OffRampOverSuperchainInteropChainSelectorToChainIdConfigArgs) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.Contract.ApplyChainSelectorToChainIdConfigUpdates(&_OffRampOverSuperchainInterop.TransactOpts, chainSelectorsToRemove, chainSelectorsToAdd)
}

func (_OffRampOverSuperchainInterop *OffRampOverSuperchainInteropTransactorSession) ApplyChainSelectorToChainIdConfigUpdates(chainSelectorsToRemove []uint64, chainSelectorsToAdd []OffRampOverSuperchainInteropChainSelectorToChainIdConfigArgs) (*types.Transaction, error) {
	return _OffRampOverSuperchainInterop.Contract.ApplyChainSelectorToChainIdConfigUpdates(&_OffRampOverSuperchainInterop.TransactOpts, chainSelectorsToRemove, chainSelectorsToAdd)
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

	GetChainId(opts *bind.CallOpts, chainSelector uint64) (*big.Int, error)

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

	ApplyChainSelectorToChainIdConfigUpdates(opts *bind.TransactOpts, chainSelectorsToRemove []uint64, chainSelectorsToAdd []OffRampOverSuperchainInteropChainSelectorToChainIdConfigArgs) (*types.Transaction, error)

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

	Address() common.Address
}
