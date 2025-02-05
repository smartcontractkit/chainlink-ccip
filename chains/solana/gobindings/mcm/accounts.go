// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package mcm

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type ConfigSignersAccount struct {
	SignerAddresses [][20]uint8
	TotalSigners    uint8
	IsFinalized     bool
}

var ConfigSignersAccountDiscriminator = [8]byte{147, 137, 80, 98, 50, 225, 190, 163}

func (obj ConfigSignersAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(ConfigSignersAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `SignerAddresses` param:
	err = encoder.Encode(obj.SignerAddresses)
	if err != nil {
		return err
	}
	// Serialize `TotalSigners` param:
	err = encoder.Encode(obj.TotalSigners)
	if err != nil {
		return err
	}
	// Serialize `IsFinalized` param:
	err = encoder.Encode(obj.IsFinalized)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ConfigSignersAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(ConfigSignersAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[147 137 80 98 50 225 190 163]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `SignerAddresses`:
	err = decoder.Decode(&obj.SignerAddresses)
	if err != nil {
		return err
	}
	// Deserialize `TotalSigners`:
	err = decoder.Decode(&obj.TotalSigners)
	if err != nil {
		return err
	}
	// Deserialize `IsFinalized`:
	err = decoder.Decode(&obj.IsFinalized)
	if err != nil {
		return err
	}
	return nil
}

type ExpiringRootAndOpCountAccount struct {
	Root       [32]uint8
	ValidUntil uint32
	OpCount    uint64
}

var ExpiringRootAndOpCountAccountDiscriminator = [8]byte{196, 176, 71, 210, 134, 228, 202, 75}

func (obj ExpiringRootAndOpCountAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(ExpiringRootAndOpCountAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Root` param:
	err = encoder.Encode(obj.Root)
	if err != nil {
		return err
	}
	// Serialize `ValidUntil` param:
	err = encoder.Encode(obj.ValidUntil)
	if err != nil {
		return err
	}
	// Serialize `OpCount` param:
	err = encoder.Encode(obj.OpCount)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ExpiringRootAndOpCountAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(ExpiringRootAndOpCountAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[196 176 71 210 134 228 202 75]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Root`:
	err = decoder.Decode(&obj.Root)
	if err != nil {
		return err
	}
	// Deserialize `ValidUntil`:
	err = decoder.Decode(&obj.ValidUntil)
	if err != nil {
		return err
	}
	// Deserialize `OpCount`:
	err = decoder.Decode(&obj.OpCount)
	if err != nil {
		return err
	}
	return nil
}

type MultisigConfigAccount struct {
	ChainId       uint64
	MultisigId    [32]uint8
	Owner         ag_solanago.PublicKey
	ProposedOwner ag_solanago.PublicKey
	GroupQuorums  [32]uint8
	GroupParents  [32]uint8
	Signers       []McmSigner
}

var MultisigConfigAccountDiscriminator = [8]byte{44, 62, 172, 225, 246, 3, 178, 33}

func (obj MultisigConfigAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(MultisigConfigAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `ChainId` param:
	err = encoder.Encode(obj.ChainId)
	if err != nil {
		return err
	}
	// Serialize `MultisigId` param:
	err = encoder.Encode(obj.MultisigId)
	if err != nil {
		return err
	}
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
	// Serialize `GroupQuorums` param:
	err = encoder.Encode(obj.GroupQuorums)
	if err != nil {
		return err
	}
	// Serialize `GroupParents` param:
	err = encoder.Encode(obj.GroupParents)
	if err != nil {
		return err
	}
	// Serialize `Signers` param:
	err = encoder.Encode(obj.Signers)
	if err != nil {
		return err
	}
	return nil
}

func (obj *MultisigConfigAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(MultisigConfigAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[44 62 172 225 246 3 178 33]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `ChainId`:
	err = decoder.Decode(&obj.ChainId)
	if err != nil {
		return err
	}
	// Deserialize `MultisigId`:
	err = decoder.Decode(&obj.MultisigId)
	if err != nil {
		return err
	}
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
	// Deserialize `GroupQuorums`:
	err = decoder.Decode(&obj.GroupQuorums)
	if err != nil {
		return err
	}
	// Deserialize `GroupParents`:
	err = decoder.Decode(&obj.GroupParents)
	if err != nil {
		return err
	}
	// Deserialize `Signers`:
	err = decoder.Decode(&obj.Signers)
	if err != nil {
		return err
	}
	return nil
}

type RootMetadataAccount struct {
	ChainId              uint64
	Multisig             ag_solanago.PublicKey
	PreOpCount           uint64
	PostOpCount          uint64
	OverridePreviousRoot bool
}

var RootMetadataAccountDiscriminator = [8]byte{125, 211, 89, 150, 221, 6, 141, 205}

func (obj RootMetadataAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(RootMetadataAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
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

func (obj *RootMetadataAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(RootMetadataAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[125 211 89 150 221 6 141 205]",
				fmt.Sprint(discriminator[:]))
		}
	}
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

type RootSignaturesAccount struct {
	TotalSignatures uint8
	Signatures      []Signature
	IsFinalized     bool
}

var RootSignaturesAccountDiscriminator = [8]byte{21, 186, 10, 33, 117, 215, 246, 76}

func (obj RootSignaturesAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(RootSignaturesAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `TotalSignatures` param:
	err = encoder.Encode(obj.TotalSignatures)
	if err != nil {
		return err
	}
	// Serialize `Signatures` param:
	err = encoder.Encode(obj.Signatures)
	if err != nil {
		return err
	}
	// Serialize `IsFinalized` param:
	err = encoder.Encode(obj.IsFinalized)
	if err != nil {
		return err
	}
	return nil
}

func (obj *RootSignaturesAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(RootSignaturesAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[21 186 10 33 117 215 246 76]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `TotalSignatures`:
	err = decoder.Decode(&obj.TotalSignatures)
	if err != nil {
		return err
	}
	// Deserialize `Signatures`:
	err = decoder.Decode(&obj.Signatures)
	if err != nil {
		return err
	}
	// Deserialize `IsFinalized`:
	err = decoder.Decode(&obj.IsFinalized)
	if err != nil {
		return err
	}
	return nil
}

type SeenSignedHashAccount struct {
	Seen bool
}

var SeenSignedHashAccountDiscriminator = [8]byte{229, 115, 10, 185, 39, 100, 210, 151}

func (obj SeenSignedHashAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(SeenSignedHashAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Seen` param:
	err = encoder.Encode(obj.Seen)
	if err != nil {
		return err
	}
	return nil
}

func (obj *SeenSignedHashAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(SeenSignedHashAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[229 115 10 185 39 100 210 151]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Seen`:
	err = decoder.Decode(&obj.Seen)
	if err != nil {
		return err
	}
	return nil
}
