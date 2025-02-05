// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package example_burnmint_token_pool

import (
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type Config struct {
	TokenProgram     ag_solanago.PublicKey
	Mint             ag_solanago.PublicKey
	Decimals         uint8
	PoolSigner       ag_solanago.PublicKey
	PoolTokenAccount ag_solanago.PublicKey
	Owner            ag_solanago.PublicKey
	ProposedOwner    ag_solanago.PublicKey
	RateLimitAdmin   ag_solanago.PublicKey
	RampAuthority    ag_solanago.PublicKey
	ListEnabled      bool
	AllowList        []ag_solanago.PublicKey
}

func (obj Config) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `TokenProgram` param:
	err = encoder.Encode(obj.TokenProgram)
	if err != nil {
		return err
	}
	// Serialize `Mint` param:
	err = encoder.Encode(obj.Mint)
	if err != nil {
		return err
	}
	// Serialize `Decimals` param:
	err = encoder.Encode(obj.Decimals)
	if err != nil {
		return err
	}
	// Serialize `PoolSigner` param:
	err = encoder.Encode(obj.PoolSigner)
	if err != nil {
		return err
	}
	// Serialize `PoolTokenAccount` param:
	err = encoder.Encode(obj.PoolTokenAccount)
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
	// Serialize `RateLimitAdmin` param:
	err = encoder.Encode(obj.RateLimitAdmin)
	if err != nil {
		return err
	}
	// Serialize `RampAuthority` param:
	err = encoder.Encode(obj.RampAuthority)
	if err != nil {
		return err
	}
	// Serialize `ListEnabled` param:
	err = encoder.Encode(obj.ListEnabled)
	if err != nil {
		return err
	}
	// Serialize `AllowList` param:
	err = encoder.Encode(obj.AllowList)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Config) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `TokenProgram`:
	err = decoder.Decode(&obj.TokenProgram)
	if err != nil {
		return err
	}
	// Deserialize `Mint`:
	err = decoder.Decode(&obj.Mint)
	if err != nil {
		return err
	}
	// Deserialize `Decimals`:
	err = decoder.Decode(&obj.Decimals)
	if err != nil {
		return err
	}
	// Deserialize `PoolSigner`:
	err = decoder.Decode(&obj.PoolSigner)
	if err != nil {
		return err
	}
	// Deserialize `PoolTokenAccount`:
	err = decoder.Decode(&obj.PoolTokenAccount)
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
	// Deserialize `RateLimitAdmin`:
	err = decoder.Decode(&obj.RateLimitAdmin)
	if err != nil {
		return err
	}
	// Deserialize `RampAuthority`:
	err = decoder.Decode(&obj.RampAuthority)
	if err != nil {
		return err
	}
	// Deserialize `ListEnabled`:
	err = decoder.Decode(&obj.ListEnabled)
	if err != nil {
		return err
	}
	// Deserialize `AllowList`:
	err = decoder.Decode(&obj.AllowList)
	if err != nil {
		return err
	}
	return nil
}

type RemoteConfig struct {
	PoolAddresses []RemoteAddress
	TokenAddress  RemoteAddress
	Decimals      uint8
}

func (obj RemoteConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `PoolAddresses` param:
	err = encoder.Encode(obj.PoolAddresses)
	if err != nil {
		return err
	}
	// Serialize `TokenAddress` param:
	err = encoder.Encode(obj.TokenAddress)
	if err != nil {
		return err
	}
	// Serialize `Decimals` param:
	err = encoder.Encode(obj.Decimals)
	if err != nil {
		return err
	}
	return nil
}

func (obj *RemoteConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `PoolAddresses`:
	err = decoder.Decode(&obj.PoolAddresses)
	if err != nil {
		return err
	}
	// Deserialize `TokenAddress`:
	err = decoder.Decode(&obj.TokenAddress)
	if err != nil {
		return err
	}
	// Deserialize `Decimals`:
	err = decoder.Decode(&obj.Decimals)
	if err != nil {
		return err
	}
	return nil
}

type RemoteAddress struct {
	Address []byte
}

func (obj RemoteAddress) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Address` param:
	err = encoder.Encode(obj.Address)
	if err != nil {
		return err
	}
	return nil
}

func (obj *RemoteAddress) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Address`:
	err = decoder.Decode(&obj.Address)
	if err != nil {
		return err
	}
	return nil
}

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

type LockOrBurnOutV1 struct {
	DestTokenAddress RemoteAddress
	DestPoolData     RemoteAddress
}

func (obj LockOrBurnOutV1) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `DestTokenAddress` param:
	err = encoder.Encode(obj.DestTokenAddress)
	if err != nil {
		return err
	}
	// Serialize `DestPoolData` param:
	err = encoder.Encode(obj.DestPoolData)
	if err != nil {
		return err
	}
	return nil
}

func (obj *LockOrBurnOutV1) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `DestTokenAddress`:
	err = decoder.Decode(&obj.DestTokenAddress)
	if err != nil {
		return err
	}
	// Deserialize `DestPoolData`:
	err = decoder.Decode(&obj.DestPoolData)
	if err != nil {
		return err
	}
	return nil
}

type ReleaseOrMintInV1 struct {
	OriginalSender      RemoteAddress
	RemoteChainSelector uint64
	Receiver            ag_solanago.PublicKey
	Amount              [32]uint8
	LocalToken          ag_solanago.PublicKey

	// @dev WARNING: sourcePoolAddress should be checked prior to any processing of funds. Make sure it matches the
	// expected pool address for the given remoteChainSelector.
	SourcePoolAddress RemoteAddress
	SourcePoolData    RemoteAddress

	// @dev WARNING: offchainTokenData is untrusted data.
	OffchainTokenData RemoteAddress
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

type ReleaseOrMintOutV1 struct {
	DestinationAmount uint64
}

func (obj ReleaseOrMintOutV1) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `DestinationAmount` param:
	err = encoder.Encode(obj.DestinationAmount)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ReleaseOrMintOutV1) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `DestinationAmount`:
	err = decoder.Decode(&obj.DestinationAmount)
	if err != nil {
		return err
	}
	return nil
}

type RateLimitTokenBucket struct {
	Tokens      uint64
	LastUpdated uint64
	Cfg         RateLimitConfig
}

func (obj RateLimitTokenBucket) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Tokens` param:
	err = encoder.Encode(obj.Tokens)
	if err != nil {
		return err
	}
	// Serialize `LastUpdated` param:
	err = encoder.Encode(obj.LastUpdated)
	if err != nil {
		return err
	}
	// Serialize `Cfg` param:
	err = encoder.Encode(obj.Cfg)
	if err != nil {
		return err
	}
	return nil
}

func (obj *RateLimitTokenBucket) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Tokens`:
	err = decoder.Decode(&obj.Tokens)
	if err != nil {
		return err
	}
	// Deserialize `LastUpdated`:
	err = decoder.Decode(&obj.LastUpdated)
	if err != nil {
		return err
	}
	// Deserialize `Cfg`:
	err = decoder.Decode(&obj.Cfg)
	if err != nil {
		return err
	}
	return nil
}

type RateLimitConfig struct {
	Enabled  bool
	Capacity uint64
	Rate     uint64
}

func (obj RateLimitConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Enabled` param:
	err = encoder.Encode(obj.Enabled)
	if err != nil {
		return err
	}
	// Serialize `Capacity` param:
	err = encoder.Encode(obj.Capacity)
	if err != nil {
		return err
	}
	// Serialize `Rate` param:
	err = encoder.Encode(obj.Rate)
	if err != nil {
		return err
	}
	return nil
}

func (obj *RateLimitConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Enabled`:
	err = decoder.Decode(&obj.Enabled)
	if err != nil {
		return err
	}
	// Deserialize `Capacity`:
	err = decoder.Decode(&obj.Capacity)
	if err != nil {
		return err
	}
	// Deserialize `Rate`:
	err = decoder.Decode(&obj.Rate)
	if err != nil {
		return err
	}
	return nil
}
