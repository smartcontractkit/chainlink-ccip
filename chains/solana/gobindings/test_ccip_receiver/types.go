// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package test_ccip_receiver

import (
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type LockOrBurnInV1 struct {
	Receiver            []byte
	RemoteChainSelector uint64
	OriginalSender      ag_solanago.PublicKey
	Amount              uint64
	LocalToken          ag_solanago.PublicKey
}

func (obj LockOrBurnInV1) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Receiver` param:
	err = encoder.Encode(obj.Receiver)
	if err != nil {
		return err
	}
	// Serialize `RemoteChainSelector` param:
	err = encoder.Encode(obj.RemoteChainSelector)
	if err != nil {
		return err
	}
	// Serialize `OriginalSender` param:
	err = encoder.Encode(obj.OriginalSender)
	if err != nil {
		return err
	}
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	// Serialize `LocalToken` param:
	err = encoder.Encode(obj.LocalToken)
	if err != nil {
		return err
	}
	return nil
}

func (obj *LockOrBurnInV1) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Receiver`:
	err = decoder.Decode(&obj.Receiver)
	if err != nil {
		return err
	}
	// Deserialize `RemoteChainSelector`:
	err = decoder.Decode(&obj.RemoteChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `OriginalSender`:
	err = decoder.Decode(&obj.OriginalSender)
	if err != nil {
		return err
	}
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	// Deserialize `LocalToken`:
	err = decoder.Decode(&obj.LocalToken)
	if err != nil {
		return err
	}
	return nil
}

type ReleaseOrMintInV1 struct {
	OriginalSender      []byte
	RemoteChainSelector uint64
	Receiver            ag_solanago.PublicKey
	Amount              [32]uint8
	LocalToken          ag_solanago.PublicKey

	// @dev WARNING: sourcePoolAddress should be checked prior to any processing of funds. Make sure it matches the
	// expected pool address for the given remoteChainSelector.
	SourcePoolAddress []byte
	SourcePoolData    []byte

	// @dev WARNING: offchainTokenData is untrusted data.
	OffchainTokenData []byte
}

func (obj ReleaseOrMintInV1) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `OriginalSender` param:
	err = encoder.Encode(obj.OriginalSender)
	if err != nil {
		return err
	}
	// Serialize `RemoteChainSelector` param:
	err = encoder.Encode(obj.RemoteChainSelector)
	if err != nil {
		return err
	}
	// Serialize `Receiver` param:
	err = encoder.Encode(obj.Receiver)
	if err != nil {
		return err
	}
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	// Serialize `LocalToken` param:
	err = encoder.Encode(obj.LocalToken)
	if err != nil {
		return err
	}
	// Serialize `SourcePoolAddress` param:
	err = encoder.Encode(obj.SourcePoolAddress)
	if err != nil {
		return err
	}
	// Serialize `SourcePoolData` param:
	err = encoder.Encode(obj.SourcePoolData)
	if err != nil {
		return err
	}
	// Serialize `OffchainTokenData` param:
	err = encoder.Encode(obj.OffchainTokenData)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ReleaseOrMintInV1) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `OriginalSender`:
	err = decoder.Decode(&obj.OriginalSender)
	if err != nil {
		return err
	}
	// Deserialize `RemoteChainSelector`:
	err = decoder.Decode(&obj.RemoteChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `Receiver`:
	err = decoder.Decode(&obj.Receiver)
	if err != nil {
		return err
	}
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	// Deserialize `LocalToken`:
	err = decoder.Decode(&obj.LocalToken)
	if err != nil {
		return err
	}
	// Deserialize `SourcePoolAddress`:
	err = decoder.Decode(&obj.SourcePoolAddress)
	if err != nil {
		return err
	}
	// Deserialize `SourcePoolData`:
	err = decoder.Decode(&obj.SourcePoolData)
	if err != nil {
		return err
	}
	// Deserialize `OffchainTokenData`:
	err = decoder.Decode(&obj.OffchainTokenData)
	if err != nil {
		return err
	}
	return nil
}

type Behavior ag_binary.BorshEnum

const (
	Normal_Behavior Behavior = iota
	RejectAll_Behavior
	ExtraCUs_Behavior
)

func (value Behavior) String() string {
	switch value {
	case Normal_Behavior:
		return "Normal"
	case RejectAll_Behavior:
		return "RejectAll"
	case ExtraCUs_Behavior:
		return "ExtraCUs"
	default:
		return ""
	}
}
