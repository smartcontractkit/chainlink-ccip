// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ccip_router

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Config //
// Sets the default code version to be used. This is then used by the slim routing layer to determine
// which version of the versioned business logic module (`instructions`) to use. Only the admin may set this.
//
// # Shared func signature with other programs
//
// # Arguments
//
// * `ctx` - The context containing the accounts required for updating the configuration.
// * `code_version` - The new code version to be set as default.
type SetDefaultCodeVersion struct {
	CodeVersion *CodeVersion

	// [0] = [WRITE] config
	//
	// [1] = [SIGNER] authority
	//
	// [2] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewSetDefaultCodeVersionInstructionBuilder creates a new `SetDefaultCodeVersion` instruction builder.
func NewSetDefaultCodeVersionInstructionBuilder() *SetDefaultCodeVersion {
	nd := &SetDefaultCodeVersion{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetCodeVersion sets the "codeVersion" parameter.
func (inst *SetDefaultCodeVersion) SetCodeVersion(codeVersion CodeVersion) *SetDefaultCodeVersion {
	inst.CodeVersion = &codeVersion
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *SetDefaultCodeVersion) SetConfigAccount(config ag_solanago.PublicKey) *SetDefaultCodeVersion {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config).WRITE()
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *SetDefaultCodeVersion) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *SetDefaultCodeVersion) SetAuthorityAccount(authority ag_solanago.PublicKey) *SetDefaultCodeVersion {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *SetDefaultCodeVersion) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *SetDefaultCodeVersion) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *SetDefaultCodeVersion {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *SetDefaultCodeVersion) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

func (inst SetDefaultCodeVersion) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SetDefaultCodeVersion,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetDefaultCodeVersion) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetDefaultCodeVersion) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.CodeVersion == nil {
			return errors.New("CodeVersion parameter is not set")
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

func (inst *SetDefaultCodeVersion) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetDefaultCodeVersion")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("CodeVersion", *inst.CodeVersion))
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

func (obj SetDefaultCodeVersion) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `CodeVersion` param:
	err = encoder.Encode(obj.CodeVersion)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SetDefaultCodeVersion) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `CodeVersion`:
	err = decoder.Decode(&obj.CodeVersion)
	if err != nil {
		return err
	}
	return nil
}

// NewSetDefaultCodeVersionInstruction declares a new SetDefaultCodeVersion instruction with the provided parameters and accounts.
func NewSetDefaultCodeVersionInstruction(
	// Parameters:
	codeVersion CodeVersion,
	// Accounts:
	config ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *SetDefaultCodeVersion {
	return NewSetDefaultCodeVersionInstructionBuilder().
		SetCodeVersion(codeVersion).
		SetConfigAccount(config).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram)
}
