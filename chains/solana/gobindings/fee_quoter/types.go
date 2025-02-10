// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package fee_quoter

import (
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

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

type EVMExtraArgsV2 struct {
	GasLimit                 ag_binary.Uint128
	AllowOutOfOrderExecution bool
}

func (obj EVMExtraArgsV2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `GasLimit` param:
	err = encoder.Encode(obj.GasLimit)
	if err != nil {
		return err
	}
	// Serialize `AllowOutOfOrderExecution` param:
	err = encoder.Encode(obj.AllowOutOfOrderExecution)
	if err != nil {
		return err
	}
	return nil
}

func (obj *EVMExtraArgsV2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `GasLimit`:
	err = decoder.Decode(&obj.GasLimit)
	if err != nil {
		return err
	}
	// Deserialize `AllowOutOfOrderExecution`:
	err = decoder.Decode(&obj.AllowOutOfOrderExecution)
	if err != nil {
		return err
	}
	return nil
}

type SVMExtraArgsV1 struct {
	ComputeUnits             uint32
	AccountIsWritableBitmap  uint64
	AllowOutOfOrderExecution bool
	TokenReceiver            [32]uint8
	Accounts                 [][32]uint8
}

func (obj SVMExtraArgsV1) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ComputeUnits` param:
	err = encoder.Encode(obj.ComputeUnits)
	if err != nil {
		return err
	}
	// Serialize `AccountIsWritableBitmap` param:
	err = encoder.Encode(obj.AccountIsWritableBitmap)
	if err != nil {
		return err
	}
	// Serialize `AllowOutOfOrderExecution` param:
	err = encoder.Encode(obj.AllowOutOfOrderExecution)
	if err != nil {
		return err
	}
	// Serialize `TokenReceiver` param:
	err = encoder.Encode(obj.TokenReceiver)
	if err != nil {
		return err
	}
	// Serialize `Accounts` param:
	err = encoder.Encode(obj.Accounts)
	if err != nil {
		return err
	}
	return nil
}

func (obj *SVMExtraArgsV1) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ComputeUnits`:
	err = decoder.Decode(&obj.ComputeUnits)
	if err != nil {
		return err
	}
	// Deserialize `AccountIsWritableBitmap`:
	err = decoder.Decode(&obj.AccountIsWritableBitmap)
	if err != nil {
		return err
	}
	// Deserialize `AllowOutOfOrderExecution`:
	err = decoder.Decode(&obj.AllowOutOfOrderExecution)
	if err != nil {
		return err
	}
	// Deserialize `TokenReceiver`:
	err = decoder.Decode(&obj.TokenReceiver)
	if err != nil {
		return err
	}
	// Deserialize `Accounts`:
	err = decoder.Decode(&obj.Accounts)
	if err != nil {
		return err
	}
	return nil
}

type SVM2AnyMessage struct {
	Receiver     []byte
	Data         []byte
	TokenAmounts []SVMTokenAmount
	FeeToken     ag_solanago.PublicKey
	ExtraArgs    []byte
}

func (obj SVM2AnyMessage) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Receiver` param:
	err = encoder.Encode(obj.Receiver)
	if err != nil {
		return err
	}
	// Serialize `Data` param:
	err = encoder.Encode(obj.Data)
	if err != nil {
		return err
	}
	// Serialize `TokenAmounts` param:
	err = encoder.Encode(obj.TokenAmounts)
	if err != nil {
		return err
	}
	// Serialize `FeeToken` param:
	err = encoder.Encode(obj.FeeToken)
	if err != nil {
		return err
	}
	// Serialize `ExtraArgs` param:
	err = encoder.Encode(obj.ExtraArgs)
	if err != nil {
		return err
	}
	return nil
}

func (obj *SVM2AnyMessage) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Receiver`:
	err = decoder.Decode(&obj.Receiver)
	if err != nil {
		return err
	}
	// Deserialize `Data`:
	err = decoder.Decode(&obj.Data)
	if err != nil {
		return err
	}
	// Deserialize `TokenAmounts`:
	err = decoder.Decode(&obj.TokenAmounts)
	if err != nil {
		return err
	}
	// Deserialize `FeeToken`:
	err = decoder.Decode(&obj.FeeToken)
	if err != nil {
		return err
	}
	// Deserialize `ExtraArgs`:
	err = decoder.Decode(&obj.ExtraArgs)
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

type TokenTransferAdditionalData struct {
	DestBytesOverhead uint32
	DestGasOverhead   uint32
}

func (obj TokenTransferAdditionalData) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `DestBytesOverhead` param:
	err = encoder.Encode(obj.DestBytesOverhead)
	if err != nil {
		return err
	}
	// Serialize `DestGasOverhead` param:
	err = encoder.Encode(obj.DestGasOverhead)
	if err != nil {
		return err
	}
	return nil
}

func (obj *TokenTransferAdditionalData) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `DestBytesOverhead`:
	err = decoder.Decode(&obj.DestBytesOverhead)
	if err != nil {
		return err
	}
	// Deserialize `DestGasOverhead`:
	err = decoder.Decode(&obj.DestGasOverhead)
	if err != nil {
		return err
	}
	return nil
}

type GetFeeResult struct {
	Token                       ag_solanago.PublicKey
	Amount                      uint64
	Juels                       uint64
	TokenTransferAdditionalData []TokenTransferAdditionalData
	ProcessedExtraArgs          ProcessedExtraArgs
}

func (obj GetFeeResult) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
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
	// Serialize `Juels` param:
	err = encoder.Encode(obj.Juels)
	if err != nil {
		return err
	}
	// Serialize `TokenTransferAdditionalData` param:
	err = encoder.Encode(obj.TokenTransferAdditionalData)
	if err != nil {
		return err
	}
	// Serialize `ProcessedExtraArgs` param:
	err = encoder.Encode(obj.ProcessedExtraArgs)
	if err != nil {
		return err
	}
	return nil
}

func (obj *GetFeeResult) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
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
	// Deserialize `Juels`:
	err = decoder.Decode(&obj.Juels)
	if err != nil {
		return err
	}
	// Deserialize `TokenTransferAdditionalData`:
	err = decoder.Decode(&obj.TokenTransferAdditionalData)
	if err != nil {
		return err
	}
	// Deserialize `ProcessedExtraArgs`:
	err = decoder.Decode(&obj.ProcessedExtraArgs)
	if err != nil {
		return err
	}
	return nil
}

type ProcessedExtraArgs struct {
	Bytes                    []byte
	GasLimit                 ag_binary.Uint128
	AllowOutOfOrderExecution bool
}

func (obj ProcessedExtraArgs) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Bytes` param:
	err = encoder.Encode(obj.Bytes)
	if err != nil {
		return err
	}
	// Serialize `GasLimit` param:
	err = encoder.Encode(obj.GasLimit)
	if err != nil {
		return err
	}
	// Serialize `AllowOutOfOrderExecution` param:
	err = encoder.Encode(obj.AllowOutOfOrderExecution)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ProcessedExtraArgs) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Bytes`:
	err = decoder.Decode(&obj.Bytes)
	if err != nil {
		return err
	}
	// Deserialize `GasLimit`:
	err = decoder.Decode(&obj.GasLimit)
	if err != nil {
		return err
	}
	// Deserialize `AllowOutOfOrderExecution`:
	err = decoder.Decode(&obj.AllowOutOfOrderExecution)
	if err != nil {
		return err
	}
	return nil
}

type DestChainConfig struct {
	IsEnabled                         bool
	MaxNumberOfTokensPerMsg           uint16
	MaxDataBytes                      uint32
	MaxPerMsgGasLimit                 uint32
	DestGasOverhead                   uint32
	DestGasPerPayloadByteBase         uint32
	DestGasPerPayloadByteHigh         uint32
	DestGasPerPayloadByteThreshold    uint32
	DestDataAvailabilityOverheadGas   uint32
	DestGasPerDataAvailabilityByte    uint16
	DestDataAvailabilityMultiplierBps uint16
	DefaultTokenFeeUsdcents           uint16
	DefaultTokenDestGasOverhead       uint32
	DefaultTxGasLimit                 uint32
	GasMultiplierWeiPerEth            uint64
	NetworkFeeUsdcents                uint32
	GasPriceStalenessThreshold        uint32
	EnforceOutOfOrder                 bool
	ChainFamilySelector               [4]uint8
}

func (obj DestChainConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `IsEnabled` param:
	err = encoder.Encode(obj.IsEnabled)
	if err != nil {
		return err
	}
	// Serialize `MaxNumberOfTokensPerMsg` param:
	err = encoder.Encode(obj.MaxNumberOfTokensPerMsg)
	if err != nil {
		return err
	}
	// Serialize `MaxDataBytes` param:
	err = encoder.Encode(obj.MaxDataBytes)
	if err != nil {
		return err
	}
	// Serialize `MaxPerMsgGasLimit` param:
	err = encoder.Encode(obj.MaxPerMsgGasLimit)
	if err != nil {
		return err
	}
	// Serialize `DestGasOverhead` param:
	err = encoder.Encode(obj.DestGasOverhead)
	if err != nil {
		return err
	}
	// Serialize `DestGasPerPayloadByteBase` param:
	err = encoder.Encode(obj.DestGasPerPayloadByteBase)
	if err != nil {
		return err
	}
	// Serialize `DestGasPerPayloadByteHigh` param:
	err = encoder.Encode(obj.DestGasPerPayloadByteHigh)
	if err != nil {
		return err
	}
	// Serialize `DestGasPerPayloadByteThreshold` param:
	err = encoder.Encode(obj.DestGasPerPayloadByteThreshold)
	if err != nil {
		return err
	}
	// Serialize `DestDataAvailabilityOverheadGas` param:
	err = encoder.Encode(obj.DestDataAvailabilityOverheadGas)
	if err != nil {
		return err
	}
	// Serialize `DestGasPerDataAvailabilityByte` param:
	err = encoder.Encode(obj.DestGasPerDataAvailabilityByte)
	if err != nil {
		return err
	}
	// Serialize `DestDataAvailabilityMultiplierBps` param:
	err = encoder.Encode(obj.DestDataAvailabilityMultiplierBps)
	if err != nil {
		return err
	}
	// Serialize `DefaultTokenFeeUsdcents` param:
	err = encoder.Encode(obj.DefaultTokenFeeUsdcents)
	if err != nil {
		return err
	}
	// Serialize `DefaultTokenDestGasOverhead` param:
	err = encoder.Encode(obj.DefaultTokenDestGasOverhead)
	if err != nil {
		return err
	}
	// Serialize `DefaultTxGasLimit` param:
	err = encoder.Encode(obj.DefaultTxGasLimit)
	if err != nil {
		return err
	}
	// Serialize `GasMultiplierWeiPerEth` param:
	err = encoder.Encode(obj.GasMultiplierWeiPerEth)
	if err != nil {
		return err
	}
	// Serialize `NetworkFeeUsdcents` param:
	err = encoder.Encode(obj.NetworkFeeUsdcents)
	if err != nil {
		return err
	}
	// Serialize `GasPriceStalenessThreshold` param:
	err = encoder.Encode(obj.GasPriceStalenessThreshold)
	if err != nil {
		return err
	}
	// Serialize `EnforceOutOfOrder` param:
	err = encoder.Encode(obj.EnforceOutOfOrder)
	if err != nil {
		return err
	}
	// Serialize `ChainFamilySelector` param:
	err = encoder.Encode(obj.ChainFamilySelector)
	if err != nil {
		return err
	}
	return nil
}

func (obj *DestChainConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `IsEnabled`:
	err = decoder.Decode(&obj.IsEnabled)
	if err != nil {
		return err
	}
	// Deserialize `MaxNumberOfTokensPerMsg`:
	err = decoder.Decode(&obj.MaxNumberOfTokensPerMsg)
	if err != nil {
		return err
	}
	// Deserialize `MaxDataBytes`:
	err = decoder.Decode(&obj.MaxDataBytes)
	if err != nil {
		return err
	}
	// Deserialize `MaxPerMsgGasLimit`:
	err = decoder.Decode(&obj.MaxPerMsgGasLimit)
	if err != nil {
		return err
	}
	// Deserialize `DestGasOverhead`:
	err = decoder.Decode(&obj.DestGasOverhead)
	if err != nil {
		return err
	}
	// Deserialize `DestGasPerPayloadByteBase`:
	err = decoder.Decode(&obj.DestGasPerPayloadByteBase)
	if err != nil {
		return err
	}
	// Deserialize `DestGasPerPayloadByteHigh`:
	err = decoder.Decode(&obj.DestGasPerPayloadByteHigh)
	if err != nil {
		return err
	}
	// Deserialize `DestGasPerPayloadByteThreshold`:
	err = decoder.Decode(&obj.DestGasPerPayloadByteThreshold)
	if err != nil {
		return err
	}
	// Deserialize `DestDataAvailabilityOverheadGas`:
	err = decoder.Decode(&obj.DestDataAvailabilityOverheadGas)
	if err != nil {
		return err
	}
	// Deserialize `DestGasPerDataAvailabilityByte`:
	err = decoder.Decode(&obj.DestGasPerDataAvailabilityByte)
	if err != nil {
		return err
	}
	// Deserialize `DestDataAvailabilityMultiplierBps`:
	err = decoder.Decode(&obj.DestDataAvailabilityMultiplierBps)
	if err != nil {
		return err
	}
	// Deserialize `DefaultTokenFeeUsdcents`:
	err = decoder.Decode(&obj.DefaultTokenFeeUsdcents)
	if err != nil {
		return err
	}
	// Deserialize `DefaultTokenDestGasOverhead`:
	err = decoder.Decode(&obj.DefaultTokenDestGasOverhead)
	if err != nil {
		return err
	}
	// Deserialize `DefaultTxGasLimit`:
	err = decoder.Decode(&obj.DefaultTxGasLimit)
	if err != nil {
		return err
	}
	// Deserialize `GasMultiplierWeiPerEth`:
	err = decoder.Decode(&obj.GasMultiplierWeiPerEth)
	if err != nil {
		return err
	}
	// Deserialize `NetworkFeeUsdcents`:
	err = decoder.Decode(&obj.NetworkFeeUsdcents)
	if err != nil {
		return err
	}
	// Deserialize `GasPriceStalenessThreshold`:
	err = decoder.Decode(&obj.GasPriceStalenessThreshold)
	if err != nil {
		return err
	}
	// Deserialize `EnforceOutOfOrder`:
	err = decoder.Decode(&obj.EnforceOutOfOrder)
	if err != nil {
		return err
	}
	// Deserialize `ChainFamilySelector`:
	err = decoder.Decode(&obj.ChainFamilySelector)
	if err != nil {
		return err
	}
	return nil
}

type DestChainState struct {
	UsdPerUnitGas TimestampedPackedU224
}

func (obj DestChainState) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `UsdPerUnitGas` param:
	err = encoder.Encode(obj.UsdPerUnitGas)
	if err != nil {
		return err
	}
	return nil
}

func (obj *DestChainState) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `UsdPerUnitGas`:
	err = decoder.Decode(&obj.UsdPerUnitGas)
	if err != nil {
		return err
	}
	return nil
}

type TimestampedPackedU224 struct {
	Value     [28]uint8
	Timestamp int64
}

func (obj TimestampedPackedU224) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Value` param:
	err = encoder.Encode(obj.Value)
	if err != nil {
		return err
	}
	// Serialize `Timestamp` param:
	err = encoder.Encode(obj.Timestamp)
	if err != nil {
		return err
	}
	return nil
}

func (obj *TimestampedPackedU224) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Value`:
	err = decoder.Decode(&obj.Value)
	if err != nil {
		return err
	}
	// Deserialize `Timestamp`:
	err = decoder.Decode(&obj.Timestamp)
	if err != nil {
		return err
	}
	return nil
}

type BillingTokenConfig struct {
	Enabled                    bool
	Mint                       ag_solanago.PublicKey
	UsdPerToken                TimestampedPackedU224
	PremiumMultiplierWeiPerEth uint64
}

func (obj BillingTokenConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Enabled` param:
	err = encoder.Encode(obj.Enabled)
	if err != nil {
		return err
	}
	// Serialize `Mint` param:
	err = encoder.Encode(obj.Mint)
	if err != nil {
		return err
	}
	// Serialize `UsdPerToken` param:
	err = encoder.Encode(obj.UsdPerToken)
	if err != nil {
		return err
	}
	// Serialize `PremiumMultiplierWeiPerEth` param:
	err = encoder.Encode(obj.PremiumMultiplierWeiPerEth)
	if err != nil {
		return err
	}
	return nil
}

func (obj *BillingTokenConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Enabled`:
	err = decoder.Decode(&obj.Enabled)
	if err != nil {
		return err
	}
	// Deserialize `Mint`:
	err = decoder.Decode(&obj.Mint)
	if err != nil {
		return err
	}
	// Deserialize `UsdPerToken`:
	err = decoder.Decode(&obj.UsdPerToken)
	if err != nil {
		return err
	}
	// Deserialize `PremiumMultiplierWeiPerEth`:
	err = decoder.Decode(&obj.PremiumMultiplierWeiPerEth)
	if err != nil {
		return err
	}
	return nil
}

type TokenTransferFeeConfig struct {
	MinFeeUsdcents    uint32
	MaxFeeUsdcents    uint32
	DeciBps           uint16
	DestGasOverhead   uint32
	DestBytesOverhead uint32
	IsEnabled         bool
}

func (obj TokenTransferFeeConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MinFeeUsdcents` param:
	err = encoder.Encode(obj.MinFeeUsdcents)
	if err != nil {
		return err
	}
	// Serialize `MaxFeeUsdcents` param:
	err = encoder.Encode(obj.MaxFeeUsdcents)
	if err != nil {
		return err
	}
	// Serialize `DeciBps` param:
	err = encoder.Encode(obj.DeciBps)
	if err != nil {
		return err
	}
	// Serialize `DestGasOverhead` param:
	err = encoder.Encode(obj.DestGasOverhead)
	if err != nil {
		return err
	}
	// Serialize `DestBytesOverhead` param:
	err = encoder.Encode(obj.DestBytesOverhead)
	if err != nil {
		return err
	}
	// Serialize `IsEnabled` param:
	err = encoder.Encode(obj.IsEnabled)
	if err != nil {
		return err
	}
	return nil
}

func (obj *TokenTransferFeeConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MinFeeUsdcents`:
	err = decoder.Decode(&obj.MinFeeUsdcents)
	if err != nil {
		return err
	}
	// Deserialize `MaxFeeUsdcents`:
	err = decoder.Decode(&obj.MaxFeeUsdcents)
	if err != nil {
		return err
	}
	// Deserialize `DeciBps`:
	err = decoder.Decode(&obj.DeciBps)
	if err != nil {
		return err
	}
	// Deserialize `DestGasOverhead`:
	err = decoder.Decode(&obj.DestGasOverhead)
	if err != nil {
		return err
	}
	// Deserialize `DestBytesOverhead`:
	err = decoder.Decode(&obj.DestBytesOverhead)
	if err != nil {
		return err
	}
	// Deserialize `IsEnabled`:
	err = decoder.Decode(&obj.IsEnabled)
	if err != nil {
		return err
	}
	return nil
}

type FeeQuoterError ag_binary.BorshEnum

const (
	InvalidSequenceInterval_FeeQuoterError FeeQuoterError = iota
	RootNotCommitted_FeeQuoterError
	ExistingMerkleRoot_FeeQuoterError
	Unauthorized_FeeQuoterError
	InvalidInputs_FeeQuoterError
	UnsupportedSourceChainSelector_FeeQuoterError
	UnsupportedDestinationChainSelector_FeeQuoterError
	InvalidProof_FeeQuoterError
	InvalidMessage_FeeQuoterError
	ReachedMaxSequenceNumber_FeeQuoterError
	ManualExecutionNotAllowed_FeeQuoterError
	InvalidInputsTokenIndices_FeeQuoterError
	InvalidInputsPoolAccounts_FeeQuoterError
	InvalidInputsTokenAccounts_FeeQuoterError
	InvalidInputsConfigAccounts_FeeQuoterError
	InvalidInputsTokenAdminRegistryAccounts_FeeQuoterError
	InvalidInputsLookupTableAccounts_FeeQuoterError
	InvalidInputsLookupTableAccountWritable_FeeQuoterError
	InvalidInputsTokenAmount_FeeQuoterError
	OfframpReleaseMintBalanceMismatch_FeeQuoterError
	OfframpInvalidDataLength_FeeQuoterError
	StaleCommitReport_FeeQuoterError
	DestinationChainDisabled_FeeQuoterError
	FeeTokenDisabled_FeeQuoterError
	MessageTooLarge_FeeQuoterError
	UnsupportedNumberOfTokens_FeeQuoterError
	UnsupportedChainFamilySelector_FeeQuoterError
	InvalidEVMAddress_FeeQuoterError
	InvalidEncoding_FeeQuoterError
	InvalidInputsAtaAddress_FeeQuoterError
	InvalidInputsAtaWritable_FeeQuoterError
	InvalidTokenPrice_FeeQuoterError
	StaleGasPrice_FeeQuoterError
	InsufficientLamports_FeeQuoterError
	InsufficientFunds_FeeQuoterError
	UnsupportedToken_FeeQuoterError
	InvalidInputsMissingTokenConfig_FeeQuoterError
	MessageFeeTooHigh_FeeQuoterError
	SourceTokenDataTooLarge_FeeQuoterError
	MessageGasLimitTooHigh_FeeQuoterError
	ExtraArgOutOfOrderExecutionMustBeTrue_FeeQuoterError
	InvalidExtraArgsTag_FeeQuoterError
	InvalidChainFamilySelector_FeeQuoterError
	InvalidTokenReceiver_FeeQuoterError
	InvalidSVMAddress_FeeQuoterError
	UnauthorizedPriceUpdater_FeeQuoterError
)

func (value FeeQuoterError) String() string {
	switch value {
	case InvalidSequenceInterval_FeeQuoterError:
		return "InvalidSequenceInterval"
	case RootNotCommitted_FeeQuoterError:
		return "RootNotCommitted"
	case ExistingMerkleRoot_FeeQuoterError:
		return "ExistingMerkleRoot"
	case Unauthorized_FeeQuoterError:
		return "Unauthorized"
	case InvalidInputs_FeeQuoterError:
		return "InvalidInputs"
	case UnsupportedSourceChainSelector_FeeQuoterError:
		return "UnsupportedSourceChainSelector"
	case UnsupportedDestinationChainSelector_FeeQuoterError:
		return "UnsupportedDestinationChainSelector"
	case InvalidProof_FeeQuoterError:
		return "InvalidProof"
	case InvalidMessage_FeeQuoterError:
		return "InvalidMessage"
	case ReachedMaxSequenceNumber_FeeQuoterError:
		return "ReachedMaxSequenceNumber"
	case ManualExecutionNotAllowed_FeeQuoterError:
		return "ManualExecutionNotAllowed"
	case InvalidInputsTokenIndices_FeeQuoterError:
		return "InvalidInputsTokenIndices"
	case InvalidInputsPoolAccounts_FeeQuoterError:
		return "InvalidInputsPoolAccounts"
	case InvalidInputsTokenAccounts_FeeQuoterError:
		return "InvalidInputsTokenAccounts"
	case InvalidInputsConfigAccounts_FeeQuoterError:
		return "InvalidInputsConfigAccounts"
	case InvalidInputsTokenAdminRegistryAccounts_FeeQuoterError:
		return "InvalidInputsTokenAdminRegistryAccounts"
	case InvalidInputsLookupTableAccounts_FeeQuoterError:
		return "InvalidInputsLookupTableAccounts"
	case InvalidInputsLookupTableAccountWritable_FeeQuoterError:
		return "InvalidInputsLookupTableAccountWritable"
	case InvalidInputsTokenAmount_FeeQuoterError:
		return "InvalidInputsTokenAmount"
	case OfframpReleaseMintBalanceMismatch_FeeQuoterError:
		return "OfframpReleaseMintBalanceMismatch"
	case OfframpInvalidDataLength_FeeQuoterError:
		return "OfframpInvalidDataLength"
	case StaleCommitReport_FeeQuoterError:
		return "StaleCommitReport"
	case DestinationChainDisabled_FeeQuoterError:
		return "DestinationChainDisabled"
	case FeeTokenDisabled_FeeQuoterError:
		return "FeeTokenDisabled"
	case MessageTooLarge_FeeQuoterError:
		return "MessageTooLarge"
	case UnsupportedNumberOfTokens_FeeQuoterError:
		return "UnsupportedNumberOfTokens"
	case UnsupportedChainFamilySelector_FeeQuoterError:
		return "UnsupportedChainFamilySelector"
	case InvalidEVMAddress_FeeQuoterError:
		return "InvalidEVMAddress"
	case InvalidEncoding_FeeQuoterError:
		return "InvalidEncoding"
	case InvalidInputsAtaAddress_FeeQuoterError:
		return "InvalidInputsAtaAddress"
	case InvalidInputsAtaWritable_FeeQuoterError:
		return "InvalidInputsAtaWritable"
	case InvalidTokenPrice_FeeQuoterError:
		return "InvalidTokenPrice"
	case StaleGasPrice_FeeQuoterError:
		return "StaleGasPrice"
	case InsufficientLamports_FeeQuoterError:
		return "InsufficientLamports"
	case InsufficientFunds_FeeQuoterError:
		return "InsufficientFunds"
	case UnsupportedToken_FeeQuoterError:
		return "UnsupportedToken"
	case InvalidInputsMissingTokenConfig_FeeQuoterError:
		return "InvalidInputsMissingTokenConfig"
	case MessageFeeTooHigh_FeeQuoterError:
		return "MessageFeeTooHigh"
	case SourceTokenDataTooLarge_FeeQuoterError:
		return "SourceTokenDataTooLarge"
	case MessageGasLimitTooHigh_FeeQuoterError:
		return "MessageGasLimitTooHigh"
	case ExtraArgOutOfOrderExecutionMustBeTrue_FeeQuoterError:
		return "ExtraArgOutOfOrderExecutionMustBeTrue"
	case InvalidExtraArgsTag_FeeQuoterError:
		return "InvalidExtraArgsTag"
	case InvalidChainFamilySelector_FeeQuoterError:
		return "InvalidChainFamilySelector"
	case InvalidTokenReceiver_FeeQuoterError:
		return "InvalidTokenReceiver"
	case InvalidSVMAddress_FeeQuoterError:
		return "InvalidSVMAddress"
	case UnauthorizedPriceUpdater_FeeQuoterError:
		return "UnauthorizedPriceUpdater"
	default:
		return ""
	}
}
