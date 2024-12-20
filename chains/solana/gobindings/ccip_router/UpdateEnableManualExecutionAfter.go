// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ccip_router

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Updates the minimum amount of time required between a message being committed and when it can be manually executed.
//
// This is part of the OffRamp Configuration for Solana.
// The Admin is the only one able to update this config.
//
// # Arguments
//
// * `ctx` - The context containing the accounts required for updating the configuration.
// * `new_enable_manual_execution_after` - The new minimum amount of time required.
type UpdateEnableManualExecutionAfter struct {
	NewEnableManualExecutionAfter *int64

	// [0] = [WRITE] config
	//
	// [1] = [SIGNER] authority
	//
	// [2] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewUpdateEnableManualExecutionAfterInstructionBuilder creates a new `UpdateEnableManualExecutionAfter` instruction builder.
func NewUpdateEnableManualExecutionAfterInstructionBuilder() *UpdateEnableManualExecutionAfter {
	nd := &UpdateEnableManualExecutionAfter{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetNewEnableManualExecutionAfter sets the "newEnableManualExecutionAfter" parameter.
func (inst *UpdateEnableManualExecutionAfter) SetNewEnableManualExecutionAfter(newEnableManualExecutionAfter int64) *UpdateEnableManualExecutionAfter {
	inst.NewEnableManualExecutionAfter = &newEnableManualExecutionAfter
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *UpdateEnableManualExecutionAfter) SetConfigAccount(config ag_solanago.PublicKey) *UpdateEnableManualExecutionAfter {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config).WRITE()
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *UpdateEnableManualExecutionAfter) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *UpdateEnableManualExecutionAfter) SetAuthorityAccount(authority ag_solanago.PublicKey) *UpdateEnableManualExecutionAfter {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *UpdateEnableManualExecutionAfter) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *UpdateEnableManualExecutionAfter) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *UpdateEnableManualExecutionAfter {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *UpdateEnableManualExecutionAfter) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

func (inst UpdateEnableManualExecutionAfter) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateEnableManualExecutionAfter,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateEnableManualExecutionAfter) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateEnableManualExecutionAfter) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.NewEnableManualExecutionAfter == nil {
			return errors.New("NewEnableManualExecutionAfter parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *UpdateEnableManualExecutionAfter) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateEnableManualExecutionAfter")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("NewEnableManualExecutionAfter", *inst.NewEnableManualExecutionAfter))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       config", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("    authority", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice[2]))
					})
				})
		})
}

func (obj UpdateEnableManualExecutionAfter) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NewEnableManualExecutionAfter` param:
	err = encoder.Encode(obj.NewEnableManualExecutionAfter)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdateEnableManualExecutionAfter) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NewEnableManualExecutionAfter`:
	err = decoder.Decode(&obj.NewEnableManualExecutionAfter)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateEnableManualExecutionAfterInstruction declares a new UpdateEnableManualExecutionAfter instruction with the provided parameters and accounts.
func NewUpdateEnableManualExecutionAfterInstruction(
	// Parameters:
	newEnableManualExecutionAfter int64,
	// Accounts:
	config ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *UpdateEnableManualExecutionAfter {
	return NewUpdateEnableManualExecutionAfterInstructionBuilder().
		SetNewEnableManualExecutionAfter(newEnableManualExecutionAfter).
		SetConfigAccount(config).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram)
}