// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package mcm

import (
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type Signature struct {
	V uint8
	R [32]uint8
	S [32]uint8
}

func (obj Signature) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `V` param:
	err = encoder.Encode(obj.V)
	if err != nil {
		return err
	}
	// Serialize `R` param:
	err = encoder.Encode(obj.R)
	if err != nil {
		return err
	}
	// Serialize `S` param:
	err = encoder.Encode(obj.S)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Signature) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `V`:
	err = decoder.Decode(&obj.V)
	if err != nil {
		return err
	}
	// Deserialize `R`:
	err = decoder.Decode(&obj.R)
	if err != nil {
		return err
	}
	// Deserialize `S`:
	err = decoder.Decode(&obj.S)
	if err != nil {
		return err
	}
	return nil
}

type McmSigner struct {
	EvmAddress [20]uint8
	Index      uint8
	Group      uint8
}

func (obj McmSigner) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `EvmAddress` param:
	err = encoder.Encode(obj.EvmAddress)
	if err != nil {
		return err
	}
	// Serialize `Index` param:
	err = encoder.Encode(obj.Index)
	if err != nil {
		return err
	}
	// Serialize `Group` param:
	err = encoder.Encode(obj.Group)
	if err != nil {
		return err
	}
	return nil
}

func (obj *McmSigner) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `EvmAddress`:
	err = decoder.Decode(&obj.EvmAddress)
	if err != nil {
		return err
	}
	// Deserialize `Index`:
	err = decoder.Decode(&obj.Index)
	if err != nil {
		return err
	}
	// Deserialize `Group`:
	err = decoder.Decode(&obj.Group)
	if err != nil {
		return err
	}
	return nil
}

type RootMetadataInput struct {
	ChainId              uint64
	Multisig             ag_solanago.PublicKey
	PreOpCount           uint64
	PostOpCount          uint64
	OverridePreviousRoot bool
}

func (obj RootMetadataInput) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ChainId` param:
	err = encoder.Encode(obj.ChainId)
	if err != nil {
		return err
	}
	// Serialize `Multisig` param:
	err = encoder.Encode(obj.Multisig)
	if err != nil {
		return err
	}
	// Serialize `PreOpCount` param:
	err = encoder.Encode(obj.PreOpCount)
	if err != nil {
		return err
	}
	// Serialize `PostOpCount` param:
	err = encoder.Encode(obj.PostOpCount)
	if err != nil {
		return err
	}
	// Serialize `OverridePreviousRoot` param:
	err = encoder.Encode(obj.OverridePreviousRoot)
	if err != nil {
		return err
	}
	return nil
}

func (obj *RootMetadataInput) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ChainId`:
	err = decoder.Decode(&obj.ChainId)
	if err != nil {
		return err
	}
	// Deserialize `Multisig`:
	err = decoder.Decode(&obj.Multisig)
	if err != nil {
		return err
	}
	// Deserialize `PreOpCount`:
	err = decoder.Decode(&obj.PreOpCount)
	if err != nil {
		return err
	}
	// Deserialize `PostOpCount`:
	err = decoder.Decode(&obj.PostOpCount)
	if err != nil {
		return err
	}
	// Deserialize `OverridePreviousRoot`:
	err = decoder.Decode(&obj.OverridePreviousRoot)
	if err != nil {
		return err
	}
	return nil
}
