// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ccip_router

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Updates the configuration of the source chain selector.
//
// The Admin is the only one able to update the source chain config.
//
// # Arguments
//
// * `ctx` - The context containing the accounts required for updating the chain selector.
// * `source_chain_selector` - The source chain selector to be updated.
// * `source_chain_config` - The new configuration for the source chain.
type UpdateSourceChainConfig struct {
	SourceChainSelector *uint64
	SourceChainConfig   *SourceChainConfig

	// [0] = [WRITE] source_chain_state
	//
	// [1] = [] config
	//
	// [2] = [WRITE, SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdateSourceChainConfigInstructionBuilder creates a new `UpdateSourceChainConfig` instruction builder.
func NewUpdateSourceChainConfigInstructionBuilder() *UpdateSourceChainConfig {
	nd := &UpdateSourceChainConfig{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetSourceChainSelector sets the "source_chain_selector" parameter.
func (inst *UpdateSourceChainConfig) SetSourceChainSelector(source_chain_selector uint64) *UpdateSourceChainConfig {
	inst.SourceChainSelector = &source_chain_selector
	return inst
}

// SetSourceChainConfig sets the "source_chain_config" parameter.
func (inst *UpdateSourceChainConfig) SetSourceChainConfig(source_chain_config SourceChainConfig) *UpdateSourceChainConfig {
	inst.SourceChainConfig = &source_chain_config
	return inst
}

// SetSourceChainStateAccount sets the "source_chain_state" account.
func (inst *UpdateSourceChainConfig) SetSourceChainStateAccount(sourceChainState ag_solanago.PublicKey) *UpdateSourceChainConfig {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(sourceChainState).WRITE()
	return inst
}

// GetSourceChainStateAccount gets the "source_chain_state" account.
func (inst *UpdateSourceChainConfig) GetSourceChainStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetConfigAccount sets the "config" account.
func (inst *UpdateSourceChainConfig) SetConfigAccount(config ag_solanago.PublicKey) *UpdateSourceChainConfig {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *UpdateSourceChainConfig) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *UpdateSourceChainConfig) SetAuthorityAccount(authority ag_solanago.PublicKey) *UpdateSourceChainConfig {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *UpdateSourceChainConfig) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst UpdateSourceChainConfig) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateSourceChainConfig,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateSourceChainConfig) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateSourceChainConfig) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.SourceChainSelector == nil {
			return errors.New("SourceChainSelector parameter is not set")
		}
		if inst.SourceChainConfig == nil {
			return errors.New("SourceChainConfig parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.SourceChainState is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Authority is not set")
		}
	}
	return nil
}

func (inst *UpdateSourceChainConfig) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateSourceChainConfig")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("  SourceChainSelector", *inst.SourceChainSelector))
						paramsBranch.Child(ag_format.Param("    SourceChainConfig", *inst.SourceChainConfig))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("source_chain_state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("            config", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("         authority", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj UpdateSourceChainConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `SourceChainSelector` param:
	err = encoder.Encode(obj.SourceChainSelector)
	if err != nil {
		return err
	}
	// Serialize `SourceChainConfig` param:
	err = encoder.Encode(obj.SourceChainConfig)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdateSourceChainConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `SourceChainSelector`:
	err = decoder.Decode(&obj.SourceChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `SourceChainConfig`:
	err = decoder.Decode(&obj.SourceChainConfig)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateSourceChainConfigInstruction declares a new UpdateSourceChainConfig instruction with the provided parameters and accounts.
func NewUpdateSourceChainConfigInstruction(
	// Parameters:
	source_chain_selector uint64,
	source_chain_config SourceChainConfig,
	// Accounts:
	sourceChainState ag_solanago.PublicKey,
	config ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *UpdateSourceChainConfig {
	return NewUpdateSourceChainConfigInstructionBuilder().
		SetSourceChainSelector(source_chain_selector).
		SetSourceChainConfig(source_chain_config).
		SetSourceChainStateAccount(sourceChainState).
		SetConfigAccount(config).
		SetAuthorityAccount(authority)
}
