// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_pool

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SetChainRateLimit is the `setChainRateLimit` instruction.
type SetChainRateLimit struct {
	RemoteChainSelector *uint64
	Mint                *ag_solanago.PublicKey
	Inbound             *RateLimitConfig
	Outbound            *RateLimitConfig

	// [0] = [] config
	//
	// [1] = [WRITE] chainConfig
	//
	// [2] = [WRITE, SIGNER] authority
	//
	// [3] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewSetChainRateLimitInstructionBuilder creates a new `SetChainRateLimit` instruction builder.
func NewSetChainRateLimitInstructionBuilder() *SetChainRateLimit {
	nd := &SetChainRateLimit{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetRemoteChainSelector sets the "remoteChainSelector" parameter.
func (inst *SetChainRateLimit) SetRemoteChainSelector(remoteChainSelector uint64) *SetChainRateLimit {
	inst.RemoteChainSelector = &remoteChainSelector
	return inst
}

// SetMint sets the "mint" parameter.
func (inst *SetChainRateLimit) SetMint(mint ag_solanago.PublicKey) *SetChainRateLimit {
	inst.Mint = &mint
	return inst
}

// SetInbound sets the "inbound" parameter.
func (inst *SetChainRateLimit) SetInbound(inbound RateLimitConfig) *SetChainRateLimit {
	inst.Inbound = &inbound
	return inst
}

// SetOutbound sets the "outbound" parameter.
func (inst *SetChainRateLimit) SetOutbound(outbound RateLimitConfig) *SetChainRateLimit {
	inst.Outbound = &outbound
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *SetChainRateLimit) SetConfigAccount(config ag_solanago.PublicKey) *SetChainRateLimit {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *SetChainRateLimit) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetChainConfigAccount sets the "chainConfig" account.
func (inst *SetChainRateLimit) SetChainConfigAccount(chainConfig ag_solanago.PublicKey) *SetChainRateLimit {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(chainConfig).WRITE()
	return inst
}

// GetChainConfigAccount gets the "chainConfig" account.
func (inst *SetChainRateLimit) GetChainConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *SetChainRateLimit) SetAuthorityAccount(authority ag_solanago.PublicKey) *SetChainRateLimit {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *SetChainRateLimit) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *SetChainRateLimit) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *SetChainRateLimit {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *SetChainRateLimit) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

func (inst SetChainRateLimit) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SetChainRateLimit,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetChainRateLimit) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetChainRateLimit) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.RemoteChainSelector == nil {
			return errors.New("RemoteChainSelector parameter is not set")
		}
		if inst.Mint == nil {
			return errors.New("Mint parameter is not set")
		}
		if inst.Inbound == nil {
			return errors.New("Inbound parameter is not set")
		}
		if inst.Outbound == nil {
			return errors.New("Outbound parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ChainConfig is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *SetChainRateLimit) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetChainRateLimit")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=4]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("RemoteChainSelector", *inst.RemoteChainSelector))
						paramsBranch.Child(ag_format.Param("               Mint", *inst.Mint))
						paramsBranch.Child(ag_format.Param("            Inbound", *inst.Inbound))
						paramsBranch.Child(ag_format.Param("           Outbound", *inst.Outbound))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       config", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("  chainConfig", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("    authority", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice[3]))
					})
				})
		})
}

func (obj SetChainRateLimit) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `RemoteChainSelector` param:
	err = encoder.Encode(obj.RemoteChainSelector)
	if err != nil {
		return err
	}
	// Serialize `Mint` param:
	err = encoder.Encode(obj.Mint)
	if err != nil {
		return err
	}
	// Serialize `Inbound` param:
	err = encoder.Encode(obj.Inbound)
	if err != nil {
		return err
	}
	// Serialize `Outbound` param:
	err = encoder.Encode(obj.Outbound)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SetChainRateLimit) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `RemoteChainSelector`:
	err = decoder.Decode(&obj.RemoteChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `Mint`:
	err = decoder.Decode(&obj.Mint)
	if err != nil {
		return err
	}
	// Deserialize `Inbound`:
	err = decoder.Decode(&obj.Inbound)
	if err != nil {
		return err
	}
	// Deserialize `Outbound`:
	err = decoder.Decode(&obj.Outbound)
	if err != nil {
		return err
	}
	return nil
}

// NewSetChainRateLimitInstruction declares a new SetChainRateLimit instruction with the provided parameters and accounts.
func NewSetChainRateLimitInstruction(
	// Parameters:
	remoteChainSelector uint64,
	mint ag_solanago.PublicKey,
	inbound RateLimitConfig,
	outbound RateLimitConfig,
	// Accounts:
	config ag_solanago.PublicKey,
	chainConfig ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *SetChainRateLimit {
	return NewSetChainRateLimitInstructionBuilder().
		SetRemoteChainSelector(remoteChainSelector).
		SetMint(mint).
		SetInbound(inbound).
		SetOutbound(outbound).
		SetConfigAccount(config).
		SetChainConfigAccount(chainConfig).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram)
}
