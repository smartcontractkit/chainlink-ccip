// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ping_pong_demo

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type Config struct {
	Owner                    ag_solanago.PublicKey
	Router                   ag_solanago.PublicKey
	CounterpartChainSelector uint64
	CounterpartAddress       CounterpartAddress
	IsPaused                 bool
	FeeTokenMint             ag_solanago.PublicKey
	ExtraArgs                []byte
}

var ConfigDiscriminator = [8]byte{155, 12, 170, 224, 30, 250, 204, 130}

func (obj Config) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(ConfigDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Owner` param:
	err = encoder.Encode(obj.Owner)
	if err != nil {
		return err
	}
	// Serialize `Router` param:
	err = encoder.Encode(obj.Router)
	if err != nil {
		return err
	}
	// Serialize `CounterpartChainSelector` param:
	err = encoder.Encode(obj.CounterpartChainSelector)
	if err != nil {
		return err
	}
	// Serialize `CounterpartAddress` param:
	err = encoder.Encode(obj.CounterpartAddress)
	if err != nil {
		return err
	}
	// Serialize `IsPaused` param:
	err = encoder.Encode(obj.IsPaused)
	if err != nil {
		return err
	}
	// Serialize `FeeTokenMint` param:
	err = encoder.Encode(obj.FeeTokenMint)
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

func (obj *Config) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(ConfigDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[155 12 170 224 30 250 204 130]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Owner`:
	err = decoder.Decode(&obj.Owner)
	if err != nil {
		return err
	}
	// Deserialize `Router`:
	err = decoder.Decode(&obj.Router)
	if err != nil {
		return err
	}
	// Deserialize `CounterpartChainSelector`:
	err = decoder.Decode(&obj.CounterpartChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `CounterpartAddress`:
	err = decoder.Decode(&obj.CounterpartAddress)
	if err != nil {
		return err
	}
	// Deserialize `IsPaused`:
	err = decoder.Decode(&obj.IsPaused)
	if err != nil {
		return err
	}
	// Deserialize `FeeTokenMint`:
	err = decoder.Decode(&obj.FeeTokenMint)
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

type NameVersion struct {
	Name    string
	Version string
}

var NameVersionDiscriminator = [8]byte{4, 169, 171, 229, 87, 69, 68, 244}

func (obj NameVersion) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(NameVersionDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
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

func (obj *NameVersion) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(NameVersionDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[4 169 171 229 87 69 68 244]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
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
