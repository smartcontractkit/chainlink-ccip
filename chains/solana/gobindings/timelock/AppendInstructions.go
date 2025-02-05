// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package timelock

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// AppendInstructions is the `append_instructions` instruction.
type AppendInstructions struct {
	TimelockId        *[32]uint8
	Id                *[32]uint8
	InstructionsBatch *[]InstructionData

	// [0] = [WRITE] operation
	//
	// [1] = [] config
	//
	// [2] = [] role_access_controller
	//
	// [3] = [WRITE, SIGNER] authority
	//
	// [4] = [] system_program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewAppendInstructionsInstructionBuilder creates a new `AppendInstructions` instruction builder.
func NewAppendInstructionsInstructionBuilder() *AppendInstructions {
	nd := &AppendInstructions{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetTimelockId sets the "timelock_id" parameter.
func (inst *AppendInstructions) SetTimelockId(timelock_id [32]uint8) *AppendInstructions {
	inst.TimelockId = &timelock_id
	return inst
}

// SetId sets the "id" parameter.
func (inst *AppendInstructions) SetId(id [32]uint8) *AppendInstructions {
	inst.Id = &id
	return inst
}

// SetInstructionsBatch sets the "instructions_batch" parameter.
func (inst *AppendInstructions) SetInstructionsBatch(instructions_batch []InstructionData) *AppendInstructions {
	inst.InstructionsBatch = &instructions_batch
	return inst
}

// SetOperationAccount sets the "operation" account.
func (inst *AppendInstructions) SetOperationAccount(operation ag_solanago.PublicKey) *AppendInstructions {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(operation).WRITE()
	return inst
}

// GetOperationAccount gets the "operation" account.
func (inst *AppendInstructions) GetOperationAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetConfigAccount sets the "config" account.
func (inst *AppendInstructions) SetConfigAccount(config ag_solanago.PublicKey) *AppendInstructions {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *AppendInstructions) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetRoleAccessControllerAccount sets the "role_access_controller" account.
func (inst *AppendInstructions) SetRoleAccessControllerAccount(roleAccessController ag_solanago.PublicKey) *AppendInstructions {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(roleAccessController)
	return inst
}

// GetRoleAccessControllerAccount gets the "role_access_controller" account.
func (inst *AppendInstructions) GetRoleAccessControllerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *AppendInstructions) SetAuthorityAccount(authority ag_solanago.PublicKey) *AppendInstructions {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *AppendInstructions) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *AppendInstructions) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *AppendInstructions {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *AppendInstructions) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst AppendInstructions) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_AppendInstructions,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AppendInstructions) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AppendInstructions) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.TimelockId == nil {
			return errors.New("TimelockId parameter is not set")
		}
		if inst.Id == nil {
			return errors.New("Id parameter is not set")
		}
		if inst.InstructionsBatch == nil {
			return errors.New("InstructionsBatch parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Operation is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.RoleAccessController is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *AppendInstructions) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AppendInstructions")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("        TimelockId", *inst.TimelockId))
						paramsBranch.Child(ag_format.Param("                Id", *inst.Id))
						paramsBranch.Child(ag_format.Param(" InstructionsBatch", *inst.InstructionsBatch))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("             operation", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                config", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("role_access_controller", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("             authority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("        system_program", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj AppendInstructions) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `TimelockId` param:
	err = encoder.Encode(obj.TimelockId)
	if err != nil {
		return err
	}
	// Serialize `Id` param:
	err = encoder.Encode(obj.Id)
	if err != nil {
		return err
	}
	// Serialize `InstructionsBatch` param:
	err = encoder.Encode(obj.InstructionsBatch)
	if err != nil {
		return err
	}
	return nil
}
func (obj *AppendInstructions) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `TimelockId`:
	err = decoder.Decode(&obj.TimelockId)
	if err != nil {
		return err
	}
	// Deserialize `Id`:
	err = decoder.Decode(&obj.Id)
	if err != nil {
		return err
	}
	// Deserialize `InstructionsBatch`:
	err = decoder.Decode(&obj.InstructionsBatch)
	if err != nil {
		return err
	}
	return nil
}

// NewAppendInstructionsInstruction declares a new AppendInstructions instruction with the provided parameters and accounts.
func NewAppendInstructionsInstruction(
	// Parameters:
	timelock_id [32]uint8,
	id [32]uint8,
	instructions_batch []InstructionData,
	// Accounts:
	operation ag_solanago.PublicKey,
	config ag_solanago.PublicKey,
	roleAccessController ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *AppendInstructions {
	return NewAppendInstructionsInstructionBuilder().
		SetTimelockId(timelock_id).
		SetId(id).
		SetInstructionsBatch(instructions_batch).
		SetOperationAccount(operation).
		SetConfigAccount(config).
		SetRoleAccessControllerAccount(roleAccessController).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram)
}
