// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package external_program_cpi_stub

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ComputeHeavy is the `computeHeavy` instruction.
type ComputeHeavy struct {
	Iterations *uint32

	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewComputeHeavyInstructionBuilder creates a new `ComputeHeavy` instruction builder.
func NewComputeHeavyInstructionBuilder() *ComputeHeavy {
	nd := &ComputeHeavy{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 0),
	}
	return nd
}

// SetIterations sets the "iterations" parameter.
func (inst *ComputeHeavy) SetIterations(iterations uint32) *ComputeHeavy {
	inst.Iterations = &iterations
	return inst
}

func (inst ComputeHeavy) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ComputeHeavy,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ComputeHeavy) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ComputeHeavy) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Iterations == nil {
			return errors.New("Iterations parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
	}
	return nil
}

func (inst *ComputeHeavy) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ComputeHeavy")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Iterations", *inst.Iterations))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=0]").ParentFunc(func(accountsBranch ag_treeout.Branches) {})
				})
		})
}

func (obj ComputeHeavy) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Iterations` param:
	err = encoder.Encode(obj.Iterations)
	if err != nil {
		return err
	}
	return nil
}
func (obj *ComputeHeavy) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Iterations`:
	err = decoder.Decode(&obj.Iterations)
	if err != nil {
		return err
	}
	return nil
}

// NewComputeHeavyInstruction declares a new ComputeHeavy instruction with the provided parameters and accounts.
func NewComputeHeavyInstruction(
	// Parameters:
	iterations uint32) *ComputeHeavy {
	return NewComputeHeavyInstructionBuilder().
		SetIterations(iterations)
}
