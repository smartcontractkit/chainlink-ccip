// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package example_ccip_receiver

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type ApprovedSenderAccount struct{}

var ApprovedSenderAccountDiscriminator = [8]byte{141, 66, 47, 213, 85, 194, 71, 166}

func (obj ApprovedSenderAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(ApprovedSenderAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ApprovedSenderAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(ApprovedSenderAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[141 66 47 213 85 194 71 166]",
				fmt.Sprint(discriminator[:]))
		}
	}
	return nil
}

type BaseStateAccount struct {
	Owner         ag_solanago.PublicKey
	ProposedOwner ag_solanago.PublicKey
	Router        ag_solanago.PublicKey
}

var BaseStateAccountDiscriminator = [8]byte{46, 139, 13, 192, 80, 181, 96, 46}

func (obj BaseStateAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(BaseStateAccountDiscriminator[:], false)
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
	// Serialize `Router` param:
	err = encoder.Encode(obj.Router)
	if err != nil {
		return err
	}
	return nil
}

func (obj *BaseStateAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(BaseStateAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[46 139 13 192 80 181 96 46]",
				fmt.Sprint(discriminator[:]))
		}
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
	// Deserialize `Router`:
	err = decoder.Decode(&obj.Router)
	if err != nil {
		return err
	}
	return nil
}
