// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ccip_offramp

import (
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type CommitInput struct {
	PriceUpdates  PriceUpdates
	MerkleRoot    *MerkleRoot `bin:"optional"`
	RmnSignatures [][64]uint8
}

func (obj CommitInput) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `PriceUpdates` param:
	err = encoder.Encode(obj.PriceUpdates)
	if err != nil {
		return err
	}
	// Serialize `MerkleRoot` param (optional):
	{
		if obj.MerkleRoot == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.MerkleRoot)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `RmnSignatures` param:
	err = encoder.Encode(obj.RmnSignatures)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CommitInput) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `PriceUpdates`:
	err = decoder.Decode(&obj.PriceUpdates)
	if err != nil {
		return err
	}
	// Deserialize `MerkleRoot` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.MerkleRoot)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `RmnSignatures`:
	err = decoder.Decode(&obj.RmnSignatures)
	if err != nil {
		return err
	}
	return nil
}

type PriceUpdates struct {
	TokenPriceUpdates []TokenPriceUpdate
	GasPriceUpdates   []GasPriceUpdate
}

func (obj PriceUpdates) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `TokenPriceUpdates` param:
	err = encoder.Encode(obj.TokenPriceUpdates)
	if err != nil {
		return err
	}
	// Serialize `GasPriceUpdates` param:
	err = encoder.Encode(obj.GasPriceUpdates)
	if err != nil {
		return err
	}
	return nil
}

func (obj *PriceUpdates) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `TokenPriceUpdates`:
	err = decoder.Decode(&obj.TokenPriceUpdates)
	if err != nil {
		return err
	}
	// Deserialize `GasPriceUpdates`:
	err = decoder.Decode(&obj.GasPriceUpdates)
	if err != nil {
		return err
	}
	return nil
}

type TokenPriceUpdate struct {
	SourceToken ag_solanago.PublicKey
	UsdPerToken [28]uint8
}

func (obj TokenPriceUpdate) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `SourceToken` param:
	err = encoder.Encode(obj.SourceToken)
	if err != nil {
		return err
	}
	// Serialize `UsdPerToken` param:
	err = encoder.Encode(obj.UsdPerToken)
	if err != nil {
		return err
	}
	return nil
}

func (obj *TokenPriceUpdate) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `SourceToken`:
	err = decoder.Decode(&obj.SourceToken)
	if err != nil {
		return err
	}
	// Deserialize `UsdPerToken`:
	err = decoder.Decode(&obj.UsdPerToken)
	if err != nil {
		return err
	}
	return nil
}

type GasPriceUpdate struct {
	DestChainSelector uint64
	UsdPerUnitGas     [28]uint8
}

func (obj GasPriceUpdate) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `DestChainSelector` param:
	err = encoder.Encode(obj.DestChainSelector)
	if err != nil {
		return err
	}
	// Serialize `UsdPerUnitGas` param:
	err = encoder.Encode(obj.UsdPerUnitGas)
	if err != nil {
		return err
	}
	return nil
}

func (obj *GasPriceUpdate) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `DestChainSelector`:
	err = decoder.Decode(&obj.DestChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `UsdPerUnitGas`:
	err = decoder.Decode(&obj.UsdPerUnitGas)
	if err != nil {
		return err
	}
	return nil
}

type MerkleRoot struct {
	SourceChainSelector uint64
	OnRampAddress       []byte
	MinSeqNr            uint64
	MaxSeqNr            uint64
	MerkleRoot          [32]uint8
}

func (obj MerkleRoot) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `SourceChainSelector` param:
	err = encoder.Encode(obj.SourceChainSelector)
	if err != nil {
		return err
	}
	// Serialize `OnRampAddress` param:
	err = encoder.Encode(obj.OnRampAddress)
	if err != nil {
		return err
	}
	// Serialize `MinSeqNr` param:
	err = encoder.Encode(obj.MinSeqNr)
	if err != nil {
		return err
	}
	// Serialize `MaxSeqNr` param:
	err = encoder.Encode(obj.MaxSeqNr)
	if err != nil {
		return err
	}
	// Serialize `MerkleRoot` param:
	err = encoder.Encode(obj.MerkleRoot)
	if err != nil {
		return err
	}
	return nil
}

func (obj *MerkleRoot) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `SourceChainSelector`:
	err = decoder.Decode(&obj.SourceChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `OnRampAddress`:
	err = decoder.Decode(&obj.OnRampAddress)
	if err != nil {
		return err
	}
	// Deserialize `MinSeqNr`:
	err = decoder.Decode(&obj.MinSeqNr)
	if err != nil {
		return err
	}
	// Deserialize `MaxSeqNr`:
	err = decoder.Decode(&obj.MaxSeqNr)
	if err != nil {
		return err
	}
	// Deserialize `MerkleRoot`:
	err = decoder.Decode(&obj.MerkleRoot)
	if err != nil {
		return err
	}
	return nil
}

type ExecutionReportSingleChain struct {
	SourceChainSelector uint64
	Message             Any2SVMRampMessage
	OffchainTokenData   [][]byte
	Root                [32]uint8
	Proofs              [][32]uint8
}

func (obj ExecutionReportSingleChain) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `SourceChainSelector` param:
	err = encoder.Encode(obj.SourceChainSelector)
	if err != nil {
		return err
	}
	// Serialize `Message` param:
	err = encoder.Encode(obj.Message)
	if err != nil {
		return err
	}
	// Serialize `OffchainTokenData` param:
	err = encoder.Encode(obj.OffchainTokenData)
	if err != nil {
		return err
	}
	// Serialize `Root` param:
	err = encoder.Encode(obj.Root)
	if err != nil {
		return err
	}
	// Serialize `Proofs` param:
	err = encoder.Encode(obj.Proofs)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ExecutionReportSingleChain) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `SourceChainSelector`:
	err = decoder.Decode(&obj.SourceChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `Message`:
	err = decoder.Decode(&obj.Message)
	if err != nil {
		return err
	}
	// Deserialize `OffchainTokenData`:
	err = decoder.Decode(&obj.OffchainTokenData)
	if err != nil {
		return err
	}
	// Deserialize `Root`:
	err = decoder.Decode(&obj.Root)
	if err != nil {
		return err
	}
	// Deserialize `Proofs`:
	err = decoder.Decode(&obj.Proofs)
	if err != nil {
		return err
	}
	return nil
}

type RampMessageHeader struct {
	MessageId           [32]uint8
	SourceChainSelector uint64
	DestChainSelector   uint64
	SequenceNumber      uint64
	Nonce               uint64
}

func (obj RampMessageHeader) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MessageId` param:
	err = encoder.Encode(obj.MessageId)
	if err != nil {
		return err
	}
	// Serialize `SourceChainSelector` param:
	err = encoder.Encode(obj.SourceChainSelector)
	if err != nil {
		return err
	}
	// Serialize `DestChainSelector` param:
	err = encoder.Encode(obj.DestChainSelector)
	if err != nil {
		return err
	}
	// Serialize `SequenceNumber` param:
	err = encoder.Encode(obj.SequenceNumber)
	if err != nil {
		return err
	}
	// Serialize `Nonce` param:
	err = encoder.Encode(obj.Nonce)
	if err != nil {
		return err
	}
	return nil
}

func (obj *RampMessageHeader) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MessageId`:
	err = decoder.Decode(&obj.MessageId)
	if err != nil {
		return err
	}
	// Deserialize `SourceChainSelector`:
	err = decoder.Decode(&obj.SourceChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `DestChainSelector`:
	err = decoder.Decode(&obj.DestChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `SequenceNumber`:
	err = decoder.Decode(&obj.SequenceNumber)
	if err != nil {
		return err
	}
	// Deserialize `Nonce`:
	err = decoder.Decode(&obj.Nonce)
	if err != nil {
		return err
	}
	return nil
}

type Any2SVMRampExtraArgs struct {
	ComputeUnits     uint32
	IsWritableBitmap uint64
}

func (obj Any2SVMRampExtraArgs) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ComputeUnits` param:
	err = encoder.Encode(obj.ComputeUnits)
	if err != nil {
		return err
	}
	// Serialize `IsWritableBitmap` param:
	err = encoder.Encode(obj.IsWritableBitmap)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Any2SVMRampExtraArgs) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ComputeUnits`:
	err = decoder.Decode(&obj.ComputeUnits)
	if err != nil {
		return err
	}
	// Deserialize `IsWritableBitmap`:
	err = decoder.Decode(&obj.IsWritableBitmap)
	if err != nil {
		return err
	}
	return nil
}

type Any2SVMRampMessage struct {
	Header        RampMessageHeader
	Sender        []byte
	Data          []byte
	TokenReceiver ag_solanago.PublicKey
	TokenAmounts  []Any2SVMTokenTransfer
	ExtraArgs     Any2SVMRampExtraArgs
	OnRampAddress []byte
}

func (obj Any2SVMRampMessage) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Header` param:
	err = encoder.Encode(obj.Header)
	if err != nil {
		return err
	}
	// Serialize `Sender` param:
	err = encoder.Encode(obj.Sender)
	if err != nil {
		return err
	}
	// Serialize `Data` param:
	err = encoder.Encode(obj.Data)
	if err != nil {
		return err
	}
	// Serialize `TokenReceiver` param:
	err = encoder.Encode(obj.TokenReceiver)
	if err != nil {
		return err
	}
	// Serialize `TokenAmounts` param:
	err = encoder.Encode(obj.TokenAmounts)
	if err != nil {
		return err
	}
	// Serialize `ExtraArgs` param:
	err = encoder.Encode(obj.ExtraArgs)
	if err != nil {
		return err
	}
	// Serialize `OnRampAddress` param:
	err = encoder.Encode(obj.OnRampAddress)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Any2SVMRampMessage) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Header`:
	err = decoder.Decode(&obj.Header)
	if err != nil {
		return err
	}
	// Deserialize `Sender`:
	err = decoder.Decode(&obj.Sender)
	if err != nil {
		return err
	}
	// Deserialize `Data`:
	err = decoder.Decode(&obj.Data)
	if err != nil {
		return err
	}
	// Deserialize `TokenReceiver`:
	err = decoder.Decode(&obj.TokenReceiver)
	if err != nil {
		return err
	}
	// Deserialize `TokenAmounts`:
	err = decoder.Decode(&obj.TokenAmounts)
	if err != nil {
		return err
	}
	// Deserialize `ExtraArgs`:
	err = decoder.Decode(&obj.ExtraArgs)
	if err != nil {
		return err
	}
	// Deserialize `OnRampAddress`:
	err = decoder.Decode(&obj.OnRampAddress)
	if err != nil {
		return err
	}
	return nil
}

type Any2SVMTokenTransfer struct {
	SourcePoolAddress []byte
	DestTokenAddress  ag_solanago.PublicKey
	DestGasAmount     uint32
	ExtraData         []byte
	Amount            CrossChainAmount
}

func (obj Any2SVMTokenTransfer) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `SourcePoolAddress` param:
	err = encoder.Encode(obj.SourcePoolAddress)
	if err != nil {
		return err
	}
	// Serialize `DestTokenAddress` param:
	err = encoder.Encode(obj.DestTokenAddress)
	if err != nil {
		return err
	}
	// Serialize `DestGasAmount` param:
	err = encoder.Encode(obj.DestGasAmount)
	if err != nil {
		return err
	}
	// Serialize `ExtraData` param:
	err = encoder.Encode(obj.ExtraData)
	if err != nil {
		return err
	}
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Any2SVMTokenTransfer) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `SourcePoolAddress`:
	err = decoder.Decode(&obj.SourcePoolAddress)
	if err != nil {
		return err
	}
	// Deserialize `DestTokenAddress`:
	err = decoder.Decode(&obj.DestTokenAddress)
	if err != nil {
		return err
	}
	// Deserialize `DestGasAmount`:
	err = decoder.Decode(&obj.DestGasAmount)
	if err != nil {
		return err
	}
	// Deserialize `ExtraData`:
	err = decoder.Decode(&obj.ExtraData)
	if err != nil {
		return err
	}
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

type SVMTokenAmount struct {
	Token  ag_solanago.PublicKey
	Amount uint64
}

func (obj SVMTokenAmount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Token` param:
	err = encoder.Encode(obj.Token)
	if err != nil {
		return err
	}
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (obj *SVMTokenAmount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Token`:
	err = decoder.Decode(&obj.Token)
	if err != nil {
		return err
	}
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

type CrossChainAmount struct {
	LeBytes [32]uint8
}

func (obj CrossChainAmount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `LeBytes` param:
	err = encoder.Encode(obj.LeBytes)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CrossChainAmount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `LeBytes`:
	err = decoder.Decode(&obj.LeBytes)
	if err != nil {
		return err
	}
	return nil
}

type Ocr3ConfigInfo struct {
	ConfigDigest                   [32]uint8
	F                              uint8
	N                              uint8
	IsSignatureVerificationEnabled uint8
}

func (obj Ocr3ConfigInfo) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ConfigDigest` param:
	err = encoder.Encode(obj.ConfigDigest)
	if err != nil {
		return err
	}
	// Serialize `F` param:
	err = encoder.Encode(obj.F)
	if err != nil {
		return err
	}
	// Serialize `N` param:
	err = encoder.Encode(obj.N)
	if err != nil {
		return err
	}
	// Serialize `IsSignatureVerificationEnabled` param:
	err = encoder.Encode(obj.IsSignatureVerificationEnabled)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Ocr3ConfigInfo) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ConfigDigest`:
	err = decoder.Decode(&obj.ConfigDigest)
	if err != nil {
		return err
	}
	// Deserialize `F`:
	err = decoder.Decode(&obj.F)
	if err != nil {
		return err
	}
	// Deserialize `N`:
	err = decoder.Decode(&obj.N)
	if err != nil {
		return err
	}
	// Deserialize `IsSignatureVerificationEnabled`:
	err = decoder.Decode(&obj.IsSignatureVerificationEnabled)
	if err != nil {
		return err
	}
	return nil
}

type Ocr3Config struct {
	PluginType   uint8
	ConfigInfo   Ocr3ConfigInfo
	Signers      [16][20]uint8
	Transmitters [16][32]uint8
}

func (obj Ocr3Config) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `PluginType` param:
	err = encoder.Encode(obj.PluginType)
	if err != nil {
		return err
	}
	// Serialize `ConfigInfo` param:
	err = encoder.Encode(obj.ConfigInfo)
	if err != nil {
		return err
	}
	// Serialize `Signers` param:
	err = encoder.Encode(obj.Signers)
	if err != nil {
		return err
	}
	// Serialize `Transmitters` param:
	err = encoder.Encode(obj.Transmitters)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Ocr3Config) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `PluginType`:
	err = decoder.Decode(&obj.PluginType)
	if err != nil {
		return err
	}
	// Deserialize `ConfigInfo`:
	err = decoder.Decode(&obj.ConfigInfo)
	if err != nil {
		return err
	}
	// Deserialize `Signers`:
	err = decoder.Decode(&obj.Signers)
	if err != nil {
		return err
	}
	// Deserialize `Transmitters`:
	err = decoder.Decode(&obj.Transmitters)
	if err != nil {
		return err
	}
	return nil
}

type SourceChainConfig struct {
	IsEnabled       bool
	LaneCodeVersion CodeVersion
	OnRamp          [2][64]uint8
}

func (obj SourceChainConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `IsEnabled` param:
	err = encoder.Encode(obj.IsEnabled)
	if err != nil {
		return err
	}
	// Serialize `LaneCodeVersion` param:
	err = encoder.Encode(obj.LaneCodeVersion)
	if err != nil {
		return err
	}
	// Serialize `OnRamp` param:
	err = encoder.Encode(obj.OnRamp)
	if err != nil {
		return err
	}
	return nil
}

func (obj *SourceChainConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `IsEnabled`:
	err = decoder.Decode(&obj.IsEnabled)
	if err != nil {
		return err
	}
	// Deserialize `LaneCodeVersion`:
	err = decoder.Decode(&obj.LaneCodeVersion)
	if err != nil {
		return err
	}
	// Deserialize `OnRamp`:
	err = decoder.Decode(&obj.OnRamp)
	if err != nil {
		return err
	}
	return nil
}

type SourceChainState struct {
	MinSeqNr uint64
}

func (obj SourceChainState) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MinSeqNr` param:
	err = encoder.Encode(obj.MinSeqNr)
	if err != nil {
		return err
	}
	return nil
}

func (obj *SourceChainState) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MinSeqNr`:
	err = decoder.Decode(&obj.MinSeqNr)
	if err != nil {
		return err
	}
	return nil
}

type OcrPluginType ag_binary.BorshEnum

const (
	Commit_OcrPluginType OcrPluginType = iota
	Execution_OcrPluginType
)

func (value OcrPluginType) String() string {
	switch value {
	case Commit_OcrPluginType:
		return "Commit"
	case Execution_OcrPluginType:
		return "Execution"
	default:
		return ""
	}
}

type MessageExecutionState ag_binary.BorshEnum

const (
	Untouched_MessageExecutionState MessageExecutionState = iota
	InProgress_MessageExecutionState
	Success_MessageExecutionState
	Failure_MessageExecutionState
)

func (value MessageExecutionState) String() string {
	switch value {
	case Untouched_MessageExecutionState:
		return "Untouched"
	case InProgress_MessageExecutionState:
		return "InProgress"
	case Success_MessageExecutionState:
		return "Success"
	case Failure_MessageExecutionState:
		return "Failure"
	default:
		return ""
	}
}

type CodeVersion ag_binary.BorshEnum

const (
	Default_CodeVersion CodeVersion = iota
	V1_CodeVersion
)

func (value CodeVersion) String() string {
	switch value {
	case Default_CodeVersion:
		return "Default"
	case V1_CodeVersion:
		return "V1"
	default:
		return ""
	}
}
