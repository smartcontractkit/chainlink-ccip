// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package rmn_remote

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Uncurses an abstract subject. If the subject is CurseSubject::from_chain_selector(local_chain_selector),
// the entire chain curse will be lifted.
//
// # Only the CCIP Admin may perform this operation
//
// # Arguments
//
// * `ctx` - The context containing the accounts required for removing a curse.
// * `subject` - The subject to uncurse.
type Uncurse struct {
	Subject *CurseSubject

	// [0] = [] config
	//
	// [1] = [WRITE, SIGNER] authority
	//
	// [2] = [WRITE] curses
	//
	// [3] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewUncurseInstructionBuilder creates a new `Uncurse` instruction builder.
func NewUncurseInstructionBuilder() *Uncurse {
	nd := &Uncurse{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetSubject sets the "subject" parameter.
func (inst *Uncurse) SetSubject(subject CurseSubject) *Uncurse {
	inst.Subject = &subject
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *Uncurse) SetConfigAccount(config ag_solanago.PublicKey) *Uncurse {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *Uncurse) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *Uncurse) SetAuthorityAccount(authority ag_solanago.PublicKey) *Uncurse {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *Uncurse) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetCursesAccount sets the "curses" account.
func (inst *Uncurse) SetCursesAccount(curses ag_solanago.PublicKey) *Uncurse {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(curses).WRITE()
	return inst
}

// GetCursesAccount gets the "curses" account.
func (inst *Uncurse) GetCursesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *Uncurse) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Uncurse {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *Uncurse) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

func (inst Uncurse) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Uncurse,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Uncurse) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Uncurse) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Subject == nil {
			return errors.New("Subject parameter is not set")
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
			return errors.New("accounts.Curses is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *Uncurse) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Uncurse")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Subject", *inst.Subject))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       config", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("    authority", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("       curses", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice[3]))
					})
				})
		})
}

func (obj Uncurse) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Subject` param:
	err = encoder.Encode(obj.Subject)
	if err != nil {
		return err
	}
	return nil
}
func (obj *Uncurse) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Subject`:
	err = decoder.Decode(&obj.Subject)
	if err != nil {
		return err
	}
	return nil
}

// NewUncurseInstruction declares a new Uncurse instruction with the provided parameters and accounts.
func NewUncurseInstruction(
	// Parameters:
	subject CurseSubject,
	// Accounts:
	config ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	curses ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *Uncurse {
	return NewUncurseInstructionBuilder().
		SetSubject(subject).
		SetConfigAccount(config).
		SetAuthorityAccount(authority).
		SetCursesAccount(curses).
		SetSystemProgramAccount(systemProgram)
}
