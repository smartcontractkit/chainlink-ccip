// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package test_ccip_receiver

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SetRejectAll is the `setRejectAll` instruction.
type SetRejectAll struct {
	RejectAll *bool

	// [0] = [WRITE] counter
	//
	// [1] = [SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewSetRejectAllInstructionBuilder creates a new `SetRejectAll` instruction builder.
func NewSetRejectAllInstructionBuilder() *SetRejectAll {
	nd := &SetRejectAll{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetRejectAll sets the "rejectAll" parameter.
func (inst *SetRejectAll) SetRejectAll(rejectAll bool) *SetRejectAll {
	inst.RejectAll = &rejectAll
	return inst
}

// SetCounterAccount sets the "counter" account.
func (inst *SetRejectAll) SetCounterAccount(counter ag_solanago.PublicKey) *SetRejectAll {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(counter).WRITE()
	return inst
}

// GetCounterAccount gets the "counter" account.
func (inst *SetRejectAll) GetCounterAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *SetRejectAll) SetAuthorityAccount(authority ag_solanago.PublicKey) *SetRejectAll {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *SetRejectAll) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

func (inst SetRejectAll) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SetRejectAll,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetRejectAll) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetRejectAll) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.RejectAll == nil {
			return errors.New("RejectAll parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Counter is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
	}
	return nil
}

func (inst *SetRejectAll) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetRejectAll")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("RejectAll", *inst.RejectAll))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("  counter", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("authority", inst.AccountMetaSlice[1]))
					})
				})
		})
}

func (obj SetRejectAll) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `RejectAll` param:
	err = encoder.Encode(obj.RejectAll)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SetRejectAll) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `RejectAll`:
	err = decoder.Decode(&obj.RejectAll)
	if err != nil {
		return err
	}
	return nil
}

// NewSetRejectAllInstruction declares a new SetRejectAll instruction with the provided parameters and accounts.
func NewSetRejectAllInstruction(
	// Parameters:
	rejectAll bool,
	// Accounts:
	counter ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *SetRejectAll {
	return NewSetRejectAllInstructionBuilder().
		SetRejectAll(rejectAll).
		SetCounterAccount(counter).
		SetAuthorityAccount(authority)
}
