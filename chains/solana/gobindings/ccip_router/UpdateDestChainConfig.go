// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ccip_router

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Updates the configuration of the destination chain selector.
//
// The Admin is the only one able to update the destination chain config.
//
// # Arguments
//
// * `ctx` - The context containing the accounts required for updating the chain selector.
// * `dest_chain_selector` - The destination chain selector to be updated.
// * `dest_chain_config` - The new configuration for the destination chain.
type UpdateDestChainConfig struct {
	DestChainSelector *uint64
	DestChainConfig   *DestChainConfig

	// [0] = [WRITE] destChainState
	//
	// [1] = [] config
	//
	// [2] = [WRITE, SIGNER] authority
	//
	// [3] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewUpdateDestChainConfigInstructionBuilder creates a new `UpdateDestChainConfig` instruction builder.
func NewUpdateDestChainConfigInstructionBuilder() *UpdateDestChainConfig {
	nd := &UpdateDestChainConfig{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetDestChainSelector sets the "destChainSelector" parameter.
func (inst *UpdateDestChainConfig) SetDestChainSelector(destChainSelector uint64) *UpdateDestChainConfig {
	inst.DestChainSelector = &destChainSelector
	return inst
}

// SetDestChainConfig sets the "destChainConfig" parameter.
func (inst *UpdateDestChainConfig) SetDestChainConfig(destChainConfig DestChainConfig) *UpdateDestChainConfig {
	inst.DestChainConfig = &destChainConfig
	return inst
}

// SetDestChainStateAccount sets the "destChainState" account.
func (inst *UpdateDestChainConfig) SetDestChainStateAccount(destChainState ag_solanago.PublicKey) *UpdateDestChainConfig {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(destChainState).WRITE()
	return inst
}

// GetDestChainStateAccount gets the "destChainState" account.
func (inst *UpdateDestChainConfig) GetDestChainStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetConfigAccount sets the "config" account.
func (inst *UpdateDestChainConfig) SetConfigAccount(config ag_solanago.PublicKey) *UpdateDestChainConfig {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *UpdateDestChainConfig) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *UpdateDestChainConfig) SetAuthorityAccount(authority ag_solanago.PublicKey) *UpdateDestChainConfig {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *UpdateDestChainConfig) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *UpdateDestChainConfig) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *UpdateDestChainConfig {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *UpdateDestChainConfig) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

func (inst UpdateDestChainConfig) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateDestChainConfig,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateDestChainConfig) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateDestChainConfig) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.DestChainSelector == nil {
			return errors.New("DestChainSelector parameter is not set")
		}
		if inst.DestChainConfig == nil {
			return errors.New("DestChainConfig parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.DestChainState is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Config is not set")
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

func (inst *UpdateDestChainConfig) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateDestChainConfig")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("DestChainSelector", *inst.DestChainSelector))
						paramsBranch.Child(ag_format.Param("  DestChainConfig", *inst.DestChainConfig))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("destChainState", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("        config", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("     authority", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta(" systemProgram", inst.AccountMetaSlice[3]))
					})
				})
		})
}

func (obj UpdateDestChainConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `DestChainSelector` param:
	err = encoder.Encode(obj.DestChainSelector)
	if err != nil {
		return err
	}
	// Serialize `DestChainConfig` param:
	err = encoder.Encode(obj.DestChainConfig)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdateDestChainConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `DestChainSelector`:
	err = decoder.Decode(&obj.DestChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `DestChainConfig`:
	err = decoder.Decode(&obj.DestChainConfig)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateDestChainConfigInstruction declares a new UpdateDestChainConfig instruction with the provided parameters and accounts.
func NewUpdateDestChainConfigInstruction(
	// Parameters:
	destChainSelector uint64,
	destChainConfig DestChainConfig,
	// Accounts:
	destChainState ag_solanago.PublicKey,
	config ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *UpdateDestChainConfig {
	return NewUpdateDestChainConfigInstructionBuilder().
		SetDestChainSelector(destChainSelector).
		SetDestChainConfig(destChainConfig).
		SetDestChainStateAccount(destChainState).
		SetConfigAccount(config).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram)
}
