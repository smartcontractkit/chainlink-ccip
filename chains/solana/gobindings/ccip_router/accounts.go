// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ccip_router

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type AllowedOfframpAccount struct{}

var AllowedOfframpAccountDiscriminator = [8]byte{247, 97, 179, 16, 207, 36, 236, 132}

func (obj AllowedOfframpAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(AllowedOfframpAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	return nil
}

func (obj *AllowedOfframpAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(AllowedOfframpAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[247 97 179 16 207 36 236 132]",
				fmt.Sprint(discriminator[:]))
		}
	}
	return nil
}

type ConfigAccount struct {
	Version          uint8
	Padding0         [7]uint8
	SvmChainSelector uint64
	Padding1         [8]uint8
	Owner            ag_solanago.PublicKey
	ProposedOwner    ag_solanago.PublicKey
	Padding2         [8]uint8
	FeeQuoter        ag_solanago.PublicKey
	LinkTokenMint    ag_solanago.PublicKey
	FeeAggregator    ag_solanago.PublicKey
}

var ConfigAccountDiscriminator = [8]byte{155, 12, 170, 224, 30, 250, 204, 130}

func (obj ConfigAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(ConfigAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Version` param:
	err = encoder.Encode(obj.Version)
	if err != nil {
		return err
	}
	// Serialize `Padding0` param:
	err = encoder.Encode(obj.Padding0)
	if err != nil {
		return err
	}
	// Serialize `SvmChainSelector` param:
	err = encoder.Encode(obj.SvmChainSelector)
	if err != nil {
		return err
	}
	// Serialize `Padding1` param:
	err = encoder.Encode(obj.Padding1)
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
	// Serialize `Padding2` param:
	err = encoder.Encode(obj.Padding2)
	if err != nil {
		return err
	}
	// Serialize `FeeQuoter` param:
	err = encoder.Encode(obj.FeeQuoter)
	if err != nil {
		return err
	}
	// Serialize `LinkTokenMint` param:
	err = encoder.Encode(obj.LinkTokenMint)
	if err != nil {
		return err
	}
	// Serialize `FeeAggregator` param:
	err = encoder.Encode(obj.FeeAggregator)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ConfigAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(ConfigAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[155 12 170 224 30 250 204 130]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Version`:
	err = decoder.Decode(&obj.Version)
	if err != nil {
		return err
	}
	// Deserialize `Padding0`:
	err = decoder.Decode(&obj.Padding0)
	if err != nil {
		return err
	}
	// Deserialize `SvmChainSelector`:
	err = decoder.Decode(&obj.SvmChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `Padding1`:
	err = decoder.Decode(&obj.Padding1)
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
	// Deserialize `Padding2`:
	err = decoder.Decode(&obj.Padding2)
	if err != nil {
		return err
	}
	// Deserialize `FeeQuoter`:
	err = decoder.Decode(&obj.FeeQuoter)
	if err != nil {
		return err
	}
	// Deserialize `LinkTokenMint`:
	err = decoder.Decode(&obj.LinkTokenMint)
	if err != nil {
		return err
	}
	// Deserialize `FeeAggregator`:
	err = decoder.Decode(&obj.FeeAggregator)
	if err != nil {
		return err
	}
	return nil
}

type DestChainAccount struct {
	Version       uint8
	ChainSelector uint64
	State         DestChainState
	Config        DestChainConfig
}

var DestChainAccountDiscriminator = [8]byte{77, 18, 241, 132, 212, 54, 218, 16}

func (obj DestChainAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(DestChainAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Version` param:
	err = encoder.Encode(obj.Version)
	if err != nil {
		return err
	}
	// Serialize `ChainSelector` param:
	err = encoder.Encode(obj.ChainSelector)
	if err != nil {
		return err
	}
	// Serialize `State` param:
	err = encoder.Encode(obj.State)
	if err != nil {
		return err
	}
	// Serialize `Config` param:
	err = encoder.Encode(obj.Config)
	if err != nil {
		return err
	}
	return nil
}

func (obj *DestChainAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(DestChainAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[77 18 241 132 212 54 218 16]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Version`:
	err = decoder.Decode(&obj.Version)
	if err != nil {
		return err
	}
	// Deserialize `ChainSelector`:
	err = decoder.Decode(&obj.ChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `State`:
	err = decoder.Decode(&obj.State)
	if err != nil {
		return err
	}
	// Deserialize `Config`:
	err = decoder.Decode(&obj.Config)
	if err != nil {
		return err
	}
	return nil
}

type ExternalExecutionConfigAccount struct{}

var ExternalExecutionConfigAccountDiscriminator = [8]byte{159, 157, 150, 212, 168, 103, 117, 39}

func (obj ExternalExecutionConfigAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(ExternalExecutionConfigAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ExternalExecutionConfigAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(ExternalExecutionConfigAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[159 157 150 212 168 103 117 39]",
				fmt.Sprint(discriminator[:]))
		}
	}
	return nil
}

type NonceAccount struct {
	Version uint8
	Counter uint64
}

var NonceAccountDiscriminator = [8]byte{143, 197, 147, 95, 106, 165, 50, 43}

func (obj NonceAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(NonceAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Version` param:
	err = encoder.Encode(obj.Version)
	if err != nil {
		return err
	}
	// Serialize `Counter` param:
	err = encoder.Encode(obj.Counter)
	if err != nil {
		return err
	}
	return nil
}

func (obj *NonceAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(NonceAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[143 197 147 95 106 165 50 43]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Version`:
	err = decoder.Decode(&obj.Version)
	if err != nil {
		return err
	}
	// Deserialize `Counter`:
	err = decoder.Decode(&obj.Counter)
	if err != nil {
		return err
	}
	return nil
}

type TokenAdminRegistryAccount struct {
	Version              uint8
	Administrator        ag_solanago.PublicKey
	PendingAdministrator ag_solanago.PublicKey
	LookupTable          ag_solanago.PublicKey
	WritableIndexes      [2]ag_binary.Uint128
	Mint                 ag_solanago.PublicKey
}

var TokenAdminRegistryAccountDiscriminator = [8]byte{70, 92, 207, 200, 76, 17, 57, 114}

func (obj TokenAdminRegistryAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(TokenAdminRegistryAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Version` param:
	err = encoder.Encode(obj.Version)
	if err != nil {
		return err
	}
	// Serialize `Administrator` param:
	err = encoder.Encode(obj.Administrator)
	if err != nil {
		return err
	}
	// Serialize `PendingAdministrator` param:
	err = encoder.Encode(obj.PendingAdministrator)
	if err != nil {
		return err
	}
	// Serialize `LookupTable` param:
	err = encoder.Encode(obj.LookupTable)
	if err != nil {
		return err
	}
	// Serialize `WritableIndexes` param:
	err = encoder.Encode(obj.WritableIndexes)
	if err != nil {
		return err
	}
	// Serialize `Mint` param:
	err = encoder.Encode(obj.Mint)
	if err != nil {
		return err
	}
	return nil
}

func (obj *TokenAdminRegistryAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(TokenAdminRegistryAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[70 92 207 200 76 17 57 114]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Version`:
	err = decoder.Decode(&obj.Version)
	if err != nil {
		return err
	}
	// Deserialize `Administrator`:
	err = decoder.Decode(&obj.Administrator)
	if err != nil {
		return err
	}
	// Deserialize `PendingAdministrator`:
	err = decoder.Decode(&obj.PendingAdministrator)
	if err != nil {
		return err
	}
	// Deserialize `LookupTable`:
	err = decoder.Decode(&obj.LookupTable)
	if err != nil {
		return err
	}
	// Deserialize `WritableIndexes`:
	err = decoder.Decode(&obj.WritableIndexes)
	if err != nil {
		return err
	}
	// Deserialize `Mint`:
	err = decoder.Decode(&obj.Mint)
	if err != nil {
		return err
	}
	return nil
}
