// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package external_program_cpi_stub

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// StructInstructionData is the `structInstructionData` instruction.
type StructInstructionData struct {
	Data *Value

	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewStructInstructionDataInstructionBuilder creates a new `StructInstructionData` instruction builder.
func NewStructInstructionDataInstructionBuilder() *StructInstructionData {
	nd := &StructInstructionData{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 0),
	}
	return nd
}

// SetData sets the "data" parameter.
func (inst *StructInstructionData) SetData(data Value) *StructInstructionData {
	inst.Data = &data
	return inst
}

func (inst StructInstructionData) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_StructInstructionData,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst StructInstructionData) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *StructInstructionData) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Data == nil {
			return errors.New("Data parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
	}
	return nil
}

func (inst *StructInstructionData) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("StructInstructionData")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Data", *inst.Data))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=0]").ParentFunc(func(accountsBranch ag_treeout.Branches) {})
				})
		})
}

func (obj StructInstructionData) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Data` param:
	err = encoder.Encode(obj.Data)
	if err != nil {
		return err
	}
	return nil
}
func (obj *StructInstructionData) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Data`:
	err = decoder.Decode(&obj.Data)
	if err != nil {
		return err
	}
	return nil
}

// NewStructInstructionDataInstruction declares a new StructInstructionData instruction with the provided parameters and accounts.
func NewStructInstructionDataInstruction(
	// Parameters:
	data Value) *StructInstructionData {
	return NewStructInstructionDataInstructionBuilder().
		SetData(data)
}