// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package example_burnmint_token_pool

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ConfigureAllowList is the `configure_allow_list` instruction.
type ConfigureAllowList struct {
	Add     *[]ag_solanago.PublicKey
	Enabled *bool

	// [0] = [WRITE] state
	//
	// [1] = [WRITE, SIGNER] authority
	//
	// [2] = [] system_program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewConfigureAllowListInstructionBuilder creates a new `ConfigureAllowList` instruction builder.
func NewConfigureAllowListInstructionBuilder() *ConfigureAllowList {
	nd := &ConfigureAllowList{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetAdd sets the "add" parameter.
func (inst *ConfigureAllowList) SetAdd(add []ag_solanago.PublicKey) *ConfigureAllowList {
	inst.Add = &add
	return inst
}

// SetEnabled sets the "enabled" parameter.
func (inst *ConfigureAllowList) SetEnabled(enabled bool) *ConfigureAllowList {
	inst.Enabled = &enabled
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *ConfigureAllowList) SetStateAccount(state ag_solanago.PublicKey) *ConfigureAllowList {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *ConfigureAllowList) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *ConfigureAllowList) SetAuthorityAccount(authority ag_solanago.PublicKey) *ConfigureAllowList {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *ConfigureAllowList) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *ConfigureAllowList) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *ConfigureAllowList {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *ConfigureAllowList) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst ConfigureAllowList) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ConfigureAllowList,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ConfigureAllowList) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ConfigureAllowList) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Add == nil {
			return errors.New("Add parameter is not set")
		}
		if inst.Enabled == nil {
			return errors.New("Enabled parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
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

func (inst *ConfigureAllowList) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ConfigureAllowList")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("    Add", *inst.Add))
						paramsBranch.Child(ag_format.Param("Enabled", *inst.Enabled))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("         state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("     authority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("system_program", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj ConfigureAllowList) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Add` param:
	err = encoder.Encode(obj.Add)
	if err != nil {
		return err
	}
	// Serialize `Enabled` param:
	err = encoder.Encode(obj.Enabled)
	if err != nil {
		return err
	}
	return nil
}
func (obj *ConfigureAllowList) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Add`:
	err = decoder.Decode(&obj.Add)
	if err != nil {
		return err
	}
	// Deserialize `Enabled`:
	err = decoder.Decode(&obj.Enabled)
	if err != nil {
		return err
	}
	return nil
}

// NewConfigureAllowListInstruction declares a new ConfigureAllowList instruction with the provided parameters and accounts.
func NewConfigureAllowListInstruction(
	// Parameters:
	add []ag_solanago.PublicKey,
	enabled bool,
	// Accounts:
	state ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *ConfigureAllowList {
	return NewConfigureAllowListInstructionBuilder().
		SetAdd(add).
		SetEnabled(enabled).
		SetStateAccount(state).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram)
}
