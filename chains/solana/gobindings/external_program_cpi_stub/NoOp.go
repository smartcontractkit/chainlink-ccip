// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package external_program_cpi_stub

import (
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// no-op instruction that does nothing, also can be used to test maximum account references(remaining_accounts)
type NoOp struct {
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewNoOpInstructionBuilder creates a new `NoOp` instruction builder.
func NewNoOpInstructionBuilder() *NoOp {
	nd := &NoOp{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 0),
	}
	return nd
}

func (inst NoOp) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_NoOp,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst NoOp) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *NoOp) Validate() error {
	// Check whether all (required) accounts are set:
	{
	}
	return nil
}

func (inst *NoOp) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("NoOp")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=0]").ParentFunc(func(accountsBranch ag_treeout.Branches) {})
				})
		})
}

func (obj NoOp) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *NoOp) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewNoOpInstruction declares a new NoOp instruction with the provided parameters and accounts.
func NewNoOpInstruction() *NoOp {
	return NewNoOpInstructionBuilder()
}
