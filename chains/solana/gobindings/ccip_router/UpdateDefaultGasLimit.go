// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ccip_router

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Updates the default gas limit in the router configuration.
//
// This change affects the default value for gas limit on every other destination chain.
// The Admin is the only one able to update the default gas limit.
//
// # Arguments
//
// * `ctx` - The context containing the accounts required for updating the configuration.
// * `new_gas_limit` - The new default gas limit.
type UpdateDefaultGasLimit struct {
	NewGasLimit *ag_binary.Uint128

	// [0] = [WRITE] config
	//
	// [1] = [SIGNER] authority
	//
	// [2] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewUpdateDefaultGasLimitInstructionBuilder creates a new `UpdateDefaultGasLimit` instruction builder.
func NewUpdateDefaultGasLimitInstructionBuilder() *UpdateDefaultGasLimit {
	nd := &UpdateDefaultGasLimit{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetNewGasLimit sets the "newGasLimit" parameter.
func (inst *UpdateDefaultGasLimit) SetNewGasLimit(newGasLimit ag_binary.Uint128) *UpdateDefaultGasLimit {
	inst.NewGasLimit = &newGasLimit
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *UpdateDefaultGasLimit) SetConfigAccount(config ag_solanago.PublicKey) *UpdateDefaultGasLimit {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config).WRITE()
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *UpdateDefaultGasLimit) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *UpdateDefaultGasLimit) SetAuthorityAccount(authority ag_solanago.PublicKey) *UpdateDefaultGasLimit {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *UpdateDefaultGasLimit) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *UpdateDefaultGasLimit) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *UpdateDefaultGasLimit {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *UpdateDefaultGasLimit) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

func (inst UpdateDefaultGasLimit) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateDefaultGasLimit,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateDefaultGasLimit) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateDefaultGasLimit) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.NewGasLimit == nil {
			return errors.New("NewGasLimit parameter is not set")
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

func (inst *UpdateDefaultGasLimit) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateDefaultGasLimit")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("NewGasLimit", *inst.NewGasLimit))
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

func (obj UpdateDefaultGasLimit) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NewGasLimit` param:
	err = encoder.Encode(obj.NewGasLimit)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdateDefaultGasLimit) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NewGasLimit`:
	err = decoder.Decode(&obj.NewGasLimit)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateDefaultGasLimitInstruction declares a new UpdateDefaultGasLimit instruction with the provided parameters and accounts.
func NewUpdateDefaultGasLimitInstruction(
	// Parameters:
	newGasLimit ag_binary.Uint128,
	// Accounts:
	config ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *UpdateDefaultGasLimit {
	return NewUpdateDefaultGasLimitInstructionBuilder().
		SetNewGasLimit(newGasLimit).
		SetConfigAccount(config).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram)
}
