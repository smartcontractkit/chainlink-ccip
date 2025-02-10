// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package example_ccip_receiver

import (
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type Any2SVMMessage struct {
	MessageId           [32]uint8
	SourceChainSelector uint64
	Sender              []byte
	Data                []byte
	TokenAmounts        []SVMTokenAmount
}

func (obj Any2SVMMessage) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
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
	// Serialize `TokenAmounts` param:
	err = encoder.Encode(obj.TokenAmounts)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Any2SVMMessage) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
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
	// Deserialize `TokenAmounts`:
	err = decoder.Decode(&obj.TokenAmounts)
	if err != nil {
		return err
	}
	return nil
}

type ApprovedSender struct{}

func (obj ApprovedSender) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}

func (obj *ApprovedSender) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

type BaseState struct {
	Owner         ag_solanago.PublicKey
	ProposedOwner ag_solanago.PublicKey
	Router        ag_solanago.PublicKey
}

func (obj BaseState) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Owner` param:
	err = encoder.Encode(obj.Owner)
	if err != nil {
		return err
	}
	// Serialize `ProposedOwner` param:
	err = encoder.Encode(obj.ProposedOwner)
	if err != nil {
		return err
	}
	// Serialize `Router` param:
	err = encoder.Encode(obj.Router)
	if err != nil {
		return err
	}
	return nil
}

func (obj *BaseState) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Owner`:
	err = decoder.Decode(&obj.Owner)
	if err != nil {
		return err
	}
	// Deserialize `ProposedOwner`:
	err = decoder.Decode(&obj.ProposedOwner)
	if err != nil {
		return err
	}
	// Deserialize `Router`:
	err = decoder.Decode(&obj.Router)
	if err != nil {
		return err
	}
	return nil
}

type MessageReceived struct {
	MessageId [32]uint8
}

func (obj MessageReceived) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MessageId` param:
	err = encoder.Encode(obj.MessageId)
	if err != nil {
		return err
	}
	return nil
}

func (obj *MessageReceived) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MessageId`:
	err = decoder.Decode(&obj.MessageId)
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
