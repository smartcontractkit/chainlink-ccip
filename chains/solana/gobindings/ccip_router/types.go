// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ccip_router

import (
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

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

type SVM2AnyRampMessage struct {
	Header         RampMessageHeader
	Sender         ag_solanago.PublicKey
	Data           []byte
	Receiver       []byte
	ExtraArgs      []byte
	FeeToken       ag_solanago.PublicKey
	TokenAmounts   []SVM2AnyTokenTransfer
	FeeTokenAmount CrossChainAmount
	FeeValueJuels  CrossChainAmount
}

func (obj SVM2AnyRampMessage) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
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
	// Serialize `Receiver` param:
	err = encoder.Encode(obj.Receiver)
	if err != nil {
		return err
	}
	// Serialize `ExtraArgs` param:
	err = encoder.Encode(obj.ExtraArgs)
	if err != nil {
		return err
	}
	// Serialize `FeeToken` param:
	err = encoder.Encode(obj.FeeToken)
	if err != nil {
		return err
	}
	// Serialize `TokenAmounts` param:
	err = encoder.Encode(obj.TokenAmounts)
	if err != nil {
		return err
	}
	// Serialize `FeeTokenAmount` param:
	err = encoder.Encode(obj.FeeTokenAmount)
	if err != nil {
		return err
	}
	// Serialize `FeeValueJuels` param:
	err = encoder.Encode(obj.FeeValueJuels)
	if err != nil {
		return err
	}
	return nil
}

func (obj *SVM2AnyRampMessage) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
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
	// Deserialize `Receiver`:
	err = decoder.Decode(&obj.Receiver)
	if err != nil {
		return err
	}
	// Deserialize `ExtraArgs`:
	err = decoder.Decode(&obj.ExtraArgs)
	if err != nil {
		return err
	}
	// Deserialize `FeeToken`:
	err = decoder.Decode(&obj.FeeToken)
	if err != nil {
		return err
	}
	// Deserialize `TokenAmounts`:
	err = decoder.Decode(&obj.TokenAmounts)
	if err != nil {
		return err
	}
	// Deserialize `FeeTokenAmount`:
	err = decoder.Decode(&obj.FeeTokenAmount)
	if err != nil {
		return err
	}
	// Deserialize `FeeValueJuels`:
	err = decoder.Decode(&obj.FeeValueJuels)
	if err != nil {
		return err
	}
	return nil
}

type SVM2AnyTokenTransfer struct {
	SourcePoolAddress ag_solanago.PublicKey
	DestTokenAddress  []byte
	ExtraData         []byte
	Amount            CrossChainAmount
	DestExecData      []byte
}

func (obj SVM2AnyTokenTransfer) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
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
	// Serialize `DestExecData` param:
	err = encoder.Encode(obj.DestExecData)
	if err != nil {
		return err
	}
	return nil
}

func (obj *SVM2AnyTokenTransfer) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
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
	// Deserialize `DestExecData`:
	err = decoder.Decode(&obj.DestExecData)
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

type DestChainState struct {
	SequenceNumber         uint64
	RollbackSequenceNumber uint64
}

func (obj DestChainState) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `SequenceNumber` param:
	err = encoder.Encode(obj.SequenceNumber)
	if err != nil {
		return err
	}
	// Serialize `RollbackSequenceNumber` param:
	err = encoder.Encode(obj.RollbackSequenceNumber)
	if err != nil {
		return err
	}
	return nil
}

func (obj *DestChainState) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `SequenceNumber`:
	err = decoder.Decode(&obj.SequenceNumber)
	if err != nil {
		return err
	}
	// Deserialize `RollbackSequenceNumber`:
	err = decoder.Decode(&obj.RollbackSequenceNumber)
	if err != nil {
		return err
	}
	return nil
}

type DestChainConfig struct {
	AllowedSenders   []ag_solanago.PublicKey
	AllowListEnabled bool
}

func (obj DestChainConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `AllowedSenders` param:
	err = encoder.Encode(obj.AllowedSenders)
	if err != nil {
		return err
	}
	// Serialize `AllowListEnabled` param:
	err = encoder.Encode(obj.AllowListEnabled)
	if err != nil {
		return err
	}
	return nil
}

func (obj *DestChainConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `AllowedSenders`:
	err = decoder.Decode(&obj.AllowedSenders)
	if err != nil {
		return err
	}
	// Deserialize `AllowListEnabled`:
	err = decoder.Decode(&obj.AllowListEnabled)
	if err != nil {
		return err
	}
	return nil
}

type CcipRouterError ag_binary.BorshEnum

const (
	Unauthorized_CcipRouterError CcipRouterError = iota
	InvalidInputsMint_CcipRouterError
	InvalidVersion_CcipRouterError
	FeeTokenMismatch_CcipRouterError
	RedundantOwnerProposal_CcipRouterError
	ReachedMaxSequenceNumber_CcipRouterError
	InvalidInputsTokenIndices_CcipRouterError
	InvalidInputsPoolAccounts_CcipRouterError
	InvalidInputsTokenAccounts_CcipRouterError
	InvalidInputsConfigAccounts_CcipRouterError
	InvalidInputsTokenAdminRegistryAccounts_CcipRouterError
	InvalidInputsLookupTableAccounts_CcipRouterError
	InvalidInputsLookupTableAccountWritable_CcipRouterError
	InvalidInputsTokenAmount_CcipRouterError
	InvalidInputsTransferAllAmount_CcipRouterError
	InvalidInputsAtaAddress_CcipRouterError
	InvalidInputsAtaWritable_CcipRouterError
	InvalidInputsChainSelector_CcipRouterError
	InsufficientLamports_CcipRouterError
	InsufficientFunds_CcipRouterError
	SourceTokenDataTooLarge_CcipRouterError
	InvalidTokenAdminRegistryInputsZeroAddress_CcipRouterError
	InvalidTokenAdminRegistryProposedAdmin_CcipRouterError
	SenderNotAllowed_CcipRouterError
	InvalidCcipVersionRollback_CcipRouterError
)

func (value CcipRouterError) String() string {
	switch value {
	case Unauthorized_CcipRouterError:
		return "Unauthorized"
	case InvalidInputsMint_CcipRouterError:
		return "InvalidInputsMint"
	case InvalidVersion_CcipRouterError:
		return "InvalidVersion"
	case FeeTokenMismatch_CcipRouterError:
		return "FeeTokenMismatch"
	case RedundantOwnerProposal_CcipRouterError:
		return "RedundantOwnerProposal"
	case ReachedMaxSequenceNumber_CcipRouterError:
		return "ReachedMaxSequenceNumber"
	case InvalidInputsTokenIndices_CcipRouterError:
		return "InvalidInputsTokenIndices"
	case InvalidInputsPoolAccounts_CcipRouterError:
		return "InvalidInputsPoolAccounts"
	case InvalidInputsTokenAccounts_CcipRouterError:
		return "InvalidInputsTokenAccounts"
	case InvalidInputsConfigAccounts_CcipRouterError:
		return "InvalidInputsConfigAccounts"
	case InvalidInputsTokenAdminRegistryAccounts_CcipRouterError:
		return "InvalidInputsTokenAdminRegistryAccounts"
	case InvalidInputsLookupTableAccounts_CcipRouterError:
		return "InvalidInputsLookupTableAccounts"
	case InvalidInputsLookupTableAccountWritable_CcipRouterError:
		return "InvalidInputsLookupTableAccountWritable"
	case InvalidInputsTokenAmount_CcipRouterError:
		return "InvalidInputsTokenAmount"
	case InvalidInputsTransferAllAmount_CcipRouterError:
		return "InvalidInputsTransferAllAmount"
	case InvalidInputsAtaAddress_CcipRouterError:
		return "InvalidInputsAtaAddress"
	case InvalidInputsAtaWritable_CcipRouterError:
		return "InvalidInputsAtaWritable"
	case InvalidInputsChainSelector_CcipRouterError:
		return "InvalidInputsChainSelector"
	case InsufficientLamports_CcipRouterError:
		return "InsufficientLamports"
	case InsufficientFunds_CcipRouterError:
		return "InsufficientFunds"
	case SourceTokenDataTooLarge_CcipRouterError:
		return "SourceTokenDataTooLarge"
	case InvalidTokenAdminRegistryInputsZeroAddress_CcipRouterError:
		return "InvalidTokenAdminRegistryInputsZeroAddress"
	case InvalidTokenAdminRegistryProposedAdmin_CcipRouterError:
		return "InvalidTokenAdminRegistryProposedAdmin"
	case SenderNotAllowed_CcipRouterError:
		return "SenderNotAllowed"
	case InvalidCcipVersionRollback_CcipRouterError:
		return "InvalidCcipVersionRollback"
	default:
		return ""
	}
}
