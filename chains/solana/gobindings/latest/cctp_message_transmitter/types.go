// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package message_transmitter

import (
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type AcceptOwnershipParams struct{}

func (obj AcceptOwnershipParams) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}

func (obj *AcceptOwnershipParams) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

type DisableAttesterParams struct {
	Attester ag_solanago.PublicKey
}

func (obj DisableAttesterParams) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Attester` param:
	err = encoder.Encode(obj.Attester)
	if err != nil {
		return err
	}
	return nil
}

func (obj *DisableAttesterParams) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Attester`:
	err = decoder.Decode(&obj.Attester)
	if err != nil {
		return err
	}
	return nil
}

type EnableAttesterParams struct {
	NewAttester ag_solanago.PublicKey
}

func (obj EnableAttesterParams) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NewAttester` param:
	err = encoder.Encode(obj.NewAttester)
	if err != nil {
		return err
	}
	return nil
}

func (obj *EnableAttesterParams) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NewAttester`:
	err = decoder.Decode(&obj.NewAttester)
	if err != nil {
		return err
	}
	return nil
}

type GetNoncePDAParams struct {
	Nonce        uint64
	SourceDomain uint32
}

func (obj GetNoncePDAParams) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Nonce` param:
	err = encoder.Encode(obj.Nonce)
	if err != nil {
		return err
	}
	// Serialize `SourceDomain` param:
	err = encoder.Encode(obj.SourceDomain)
	if err != nil {
		return err
	}
	return nil
}

func (obj *GetNoncePDAParams) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Nonce`:
	err = decoder.Decode(&obj.Nonce)
	if err != nil {
		return err
	}
	// Deserialize `SourceDomain`:
	err = decoder.Decode(&obj.SourceDomain)
	if err != nil {
		return err
	}
	return nil
}

type InitializeParams struct {
	LocalDomain        uint32
	Attester           ag_solanago.PublicKey
	MaxMessageBodySize uint64
	Version            uint32
}

func (obj InitializeParams) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `LocalDomain` param:
	err = encoder.Encode(obj.LocalDomain)
	if err != nil {
		return err
	}
	// Serialize `Attester` param:
	err = encoder.Encode(obj.Attester)
	if err != nil {
		return err
	}
	// Serialize `MaxMessageBodySize` param:
	err = encoder.Encode(obj.MaxMessageBodySize)
	if err != nil {
		return err
	}
	// Serialize `Version` param:
	err = encoder.Encode(obj.Version)
	if err != nil {
		return err
	}
	return nil
}

func (obj *InitializeParams) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `LocalDomain`:
	err = decoder.Decode(&obj.LocalDomain)
	if err != nil {
		return err
	}
	// Deserialize `Attester`:
	err = decoder.Decode(&obj.Attester)
	if err != nil {
		return err
	}
	// Deserialize `MaxMessageBodySize`:
	err = decoder.Decode(&obj.MaxMessageBodySize)
	if err != nil {
		return err
	}
	// Deserialize `Version`:
	err = decoder.Decode(&obj.Version)
	if err != nil {
		return err
	}
	return nil
}

type IsNonceUsedParams struct {
	Nonce uint64
}

func (obj IsNonceUsedParams) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Nonce` param:
	err = encoder.Encode(obj.Nonce)
	if err != nil {
		return err
	}
	return nil
}

func (obj *IsNonceUsedParams) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Nonce`:
	err = decoder.Decode(&obj.Nonce)
	if err != nil {
		return err
	}
	return nil
}

type PauseParams struct{}

func (obj PauseParams) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}

func (obj *PauseParams) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

type ReceiveMessageParams struct {
	Message     []byte
	Attestation []byte
}

func (obj ReceiveMessageParams) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Message` param:
	err = encoder.Encode(obj.Message)
	if err != nil {
		return err
	}
	// Serialize `Attestation` param:
	err = encoder.Encode(obj.Attestation)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ReceiveMessageParams) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Message`:
	err = decoder.Decode(&obj.Message)
	if err != nil {
		return err
	}
	// Deserialize `Attestation`:
	err = decoder.Decode(&obj.Attestation)
	if err != nil {
		return err
	}
	return nil
}

type HandleReceiveMessageParams struct {
	RemoteDomain  uint32
	Sender        ag_solanago.PublicKey
	MessageBody   []byte
	AuthorityBump uint8
}

func (obj HandleReceiveMessageParams) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `RemoteDomain` param:
	err = encoder.Encode(obj.RemoteDomain)
	if err != nil {
		return err
	}
	// Serialize `Sender` param:
	err = encoder.Encode(obj.Sender)
	if err != nil {
		return err
	}
	// Serialize `MessageBody` param:
	err = encoder.Encode(obj.MessageBody)
	if err != nil {
		return err
	}
	// Serialize `AuthorityBump` param:
	err = encoder.Encode(obj.AuthorityBump)
	if err != nil {
		return err
	}
	return nil
}

func (obj *HandleReceiveMessageParams) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `RemoteDomain`:
	err = decoder.Decode(&obj.RemoteDomain)
	if err != nil {
		return err
	}
	// Deserialize `Sender`:
	err = decoder.Decode(&obj.Sender)
	if err != nil {
		return err
	}
	// Deserialize `MessageBody`:
	err = decoder.Decode(&obj.MessageBody)
	if err != nil {
		return err
	}
	// Deserialize `AuthorityBump`:
	err = decoder.Decode(&obj.AuthorityBump)
	if err != nil {
		return err
	}
	return nil
}

type ReclaimEventAccountParams struct {
	Attestation []byte
}

func (obj ReclaimEventAccountParams) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Attestation` param:
	err = encoder.Encode(obj.Attestation)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ReclaimEventAccountParams) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Attestation`:
	err = decoder.Decode(&obj.Attestation)
	if err != nil {
		return err
	}
	return nil
}

type ReplaceMessageParams struct {
	OriginalMessage      []byte
	OriginalAttestation  []byte
	NewMessageBody       []byte
	NewDestinationCaller ag_solanago.PublicKey
}

func (obj ReplaceMessageParams) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `OriginalMessage` param:
	err = encoder.Encode(obj.OriginalMessage)
	if err != nil {
		return err
	}
	// Serialize `OriginalAttestation` param:
	err = encoder.Encode(obj.OriginalAttestation)
	if err != nil {
		return err
	}
	// Serialize `NewMessageBody` param:
	err = encoder.Encode(obj.NewMessageBody)
	if err != nil {
		return err
	}
	// Serialize `NewDestinationCaller` param:
	err = encoder.Encode(obj.NewDestinationCaller)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ReplaceMessageParams) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `OriginalMessage`:
	err = decoder.Decode(&obj.OriginalMessage)
	if err != nil {
		return err
	}
	// Deserialize `OriginalAttestation`:
	err = decoder.Decode(&obj.OriginalAttestation)
	if err != nil {
		return err
	}
	// Deserialize `NewMessageBody`:
	err = decoder.Decode(&obj.NewMessageBody)
	if err != nil {
		return err
	}
	// Deserialize `NewDestinationCaller`:
	err = decoder.Decode(&obj.NewDestinationCaller)
	if err != nil {
		return err
	}
	return nil
}

type SendMessageWithCallerParams struct {
	DestinationDomain uint32
	Recipient         ag_solanago.PublicKey
	MessageBody       []byte
	DestinationCaller ag_solanago.PublicKey
}

func (obj SendMessageWithCallerParams) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `DestinationDomain` param:
	err = encoder.Encode(obj.DestinationDomain)
	if err != nil {
		return err
	}
	// Serialize `Recipient` param:
	err = encoder.Encode(obj.Recipient)
	if err != nil {
		return err
	}
	// Serialize `MessageBody` param:
	err = encoder.Encode(obj.MessageBody)
	if err != nil {
		return err
	}
	// Serialize `DestinationCaller` param:
	err = encoder.Encode(obj.DestinationCaller)
	if err != nil {
		return err
	}
	return nil
}

func (obj *SendMessageWithCallerParams) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `DestinationDomain`:
	err = decoder.Decode(&obj.DestinationDomain)
	if err != nil {
		return err
	}
	// Deserialize `Recipient`:
	err = decoder.Decode(&obj.Recipient)
	if err != nil {
		return err
	}
	// Deserialize `MessageBody`:
	err = decoder.Decode(&obj.MessageBody)
	if err != nil {
		return err
	}
	// Deserialize `DestinationCaller`:
	err = decoder.Decode(&obj.DestinationCaller)
	if err != nil {
		return err
	}
	return nil
}

type SendMessageParams struct {
	DestinationDomain uint32
	Recipient         ag_solanago.PublicKey
	MessageBody       []byte
}

func (obj SendMessageParams) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `DestinationDomain` param:
	err = encoder.Encode(obj.DestinationDomain)
	if err != nil {
		return err
	}
	// Serialize `Recipient` param:
	err = encoder.Encode(obj.Recipient)
	if err != nil {
		return err
	}
	// Serialize `MessageBody` param:
	err = encoder.Encode(obj.MessageBody)
	if err != nil {
		return err
	}
	return nil
}

func (obj *SendMessageParams) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `DestinationDomain`:
	err = decoder.Decode(&obj.DestinationDomain)
	if err != nil {
		return err
	}
	// Deserialize `Recipient`:
	err = decoder.Decode(&obj.Recipient)
	if err != nil {
		return err
	}
	// Deserialize `MessageBody`:
	err = decoder.Decode(&obj.MessageBody)
	if err != nil {
		return err
	}
	return nil
}

type SetMaxMessageBodySizeParams struct {
	NewMaxMessageBodySize uint64
}

func (obj SetMaxMessageBodySizeParams) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NewMaxMessageBodySize` param:
	err = encoder.Encode(obj.NewMaxMessageBodySize)
	if err != nil {
		return err
	}
	return nil
}

func (obj *SetMaxMessageBodySizeParams) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NewMaxMessageBodySize`:
	err = decoder.Decode(&obj.NewMaxMessageBodySize)
	if err != nil {
		return err
	}
	return nil
}

type SetSignatureThresholdParams struct {
	NewSignatureThreshold uint32
}

func (obj SetSignatureThresholdParams) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NewSignatureThreshold` param:
	err = encoder.Encode(obj.NewSignatureThreshold)
	if err != nil {
		return err
	}
	return nil
}

func (obj *SetSignatureThresholdParams) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NewSignatureThreshold`:
	err = decoder.Decode(&obj.NewSignatureThreshold)
	if err != nil {
		return err
	}
	return nil
}

type TransferOwnershipParams struct {
	NewOwner ag_solanago.PublicKey
}

func (obj TransferOwnershipParams) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NewOwner` param:
	err = encoder.Encode(obj.NewOwner)
	if err != nil {
		return err
	}
	return nil
}

func (obj *TransferOwnershipParams) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NewOwner`:
	err = decoder.Decode(&obj.NewOwner)
	if err != nil {
		return err
	}
	return nil
}

type UnpauseParams struct{}

func (obj UnpauseParams) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}

func (obj *UnpauseParams) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

type UpdateAttesterManagerParams struct {
	NewAttesterManager ag_solanago.PublicKey
}

func (obj UpdateAttesterManagerParams) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NewAttesterManager` param:
	err = encoder.Encode(obj.NewAttesterManager)
	if err != nil {
		return err
	}
	return nil
}

func (obj *UpdateAttesterManagerParams) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NewAttesterManager`:
	err = decoder.Decode(&obj.NewAttesterManager)
	if err != nil {
		return err
	}
	return nil
}

type UpdatePauserParams struct {
	NewPauser ag_solanago.PublicKey
}

func (obj UpdatePauserParams) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NewPauser` param:
	err = encoder.Encode(obj.NewPauser)
	if err != nil {
		return err
	}
	return nil
}

func (obj *UpdatePauserParams) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NewPauser`:
	err = decoder.Decode(&obj.NewPauser)
	if err != nil {
		return err
	}
	return nil
}

type MathError ag_binary.BorshEnum

const (
	MathOverflow_MathError MathError = iota
	MathUnderflow_MathError
	ErrorInDivision_MathError
)

func (value MathError) String() string {
	switch value {
	case MathOverflow_MathError:
		return "MathOverflow"
	case MathUnderflow_MathError:
		return "MathUnderflow"
	case ErrorInDivision_MathError:
		return "ErrorInDivision"
	default:
		return ""
	}
}
