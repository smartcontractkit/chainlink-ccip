// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package lockrelease_token_pool

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Initialize is the `initialize` instruction.
type Initialize struct {
	Router    *ag_solanago.PublicKey
	RmnRemote *ag_solanago.PublicKey

	// [0] = [WRITE] state
	//
	// [1] = [] mint
	//
	// [2] = [WRITE, SIGNER] authority
	//
	// [3] = [] systemProgram
	//
	// [4] = [] program
	//
	// [5] = [] programData
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewInitializeInstructionBuilder creates a new `Initialize` instruction builder.
func NewInitializeInstructionBuilder() *Initialize {
	nd := &Initialize{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	return nd
}

// SetRouter sets the "router" parameter.
func (inst *Initialize) SetRouter(router ag_solanago.PublicKey) *Initialize {
	inst.Router = &router
	return inst
}

// SetRmnRemote sets the "rmnRemote" parameter.
func (inst *Initialize) SetRmnRemote(rmnRemote ag_solanago.PublicKey) *Initialize {
	inst.RmnRemote = &rmnRemote
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *Initialize) SetStateAccount(state ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *Initialize) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetMintAccount sets the "mint" account.
func (inst *Initialize) SetMintAccount(mint ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *Initialize) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *Initialize) SetAuthorityAccount(authority ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *Initialize) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *Initialize) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *Initialize) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetProgramAccount sets the "program" account.
func (inst *Initialize) SetProgramAccount(program ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *Initialize) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetProgramDataAccount sets the "programData" account.
func (inst *Initialize) SetProgramDataAccount(programData ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(programData)
	return inst
}

// GetProgramDataAccount gets the "programData" account.
func (inst *Initialize) GetProgramDataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

func (inst Initialize) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Initialize,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Initialize) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Initialize) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Router == nil {
			return errors.New("Router parameter is not set")
		}
		if inst.RmnRemote == nil {
			return errors.New("RmnRemote parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Program is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.ProgramData is not set")
		}
	}
	return nil
}

func (inst *Initialize) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Initialize")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("   Router", *inst.Router))
						paramsBranch.Child(ag_format.Param("RmnRemote", *inst.RmnRemote))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        state", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("         mint", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("    authority", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("      program", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("  programData", inst.AccountMetaSlice[5]))
					})
				})
		})
}

func (obj Initialize) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Router` param:
	err = encoder.Encode(obj.Router)
	if err != nil {
		return err
	}
	// Serialize `RmnRemote` param:
	err = encoder.Encode(obj.RmnRemote)
	if err != nil {
		return err
	}
	return nil
}
func (obj *Initialize) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Router`:
	err = decoder.Decode(&obj.Router)
	if err != nil {
		return err
	}
	// Deserialize `RmnRemote`:
	err = decoder.Decode(&obj.RmnRemote)
	if err != nil {
		return err
	}
	return nil
}

// NewInitializeInstruction declares a new Initialize instruction with the provided parameters and accounts.
func NewInitializeInstruction(
	// Parameters:
	router ag_solanago.PublicKey,
	rmnRemote ag_solanago.PublicKey,
	// Accounts:
	state ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	program ag_solanago.PublicKey,
	programData ag_solanago.PublicKey) *Initialize {
	return NewInitializeInstructionBuilder().
		SetRouter(router).
		SetRmnRemote(rmnRemote).
		SetStateAccount(state).
		SetMintAccount(mint).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram).
		SetProgramAccount(program).
		SetProgramDataAccount(programData)
}
