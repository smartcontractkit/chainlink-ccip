// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package timelock

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// FinalizeOperation is the `finalizeOperation` instruction.
type FinalizeOperation struct {
	TimelockId *[32]uint8
	Id         *[32]uint8

	// [0] = [WRITE] operation
	//
	// [1] = [] config
	//
	// [2] = [] roleAccessController
	//
	// [3] = [WRITE, SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewFinalizeOperationInstructionBuilder creates a new `FinalizeOperation` instruction builder.
func NewFinalizeOperationInstructionBuilder() *FinalizeOperation {
	nd := &FinalizeOperation{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetTimelockId sets the "timelockId" parameter.
func (inst *FinalizeOperation) SetTimelockId(timelockId [32]uint8) *FinalizeOperation {
	inst.TimelockId = &timelockId
	return inst
}

// SetId sets the "id" parameter.
func (inst *FinalizeOperation) SetId(id [32]uint8) *FinalizeOperation {
	inst.Id = &id
	return inst
}

// SetOperationAccount sets the "operation" account.
func (inst *FinalizeOperation) SetOperationAccount(operation ag_solanago.PublicKey) *FinalizeOperation {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(operation).WRITE()
	return inst
}

// GetOperationAccount gets the "operation" account.
func (inst *FinalizeOperation) GetOperationAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetConfigAccount sets the "config" account.
func (inst *FinalizeOperation) SetConfigAccount(config ag_solanago.PublicKey) *FinalizeOperation {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *FinalizeOperation) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetRoleAccessControllerAccount sets the "roleAccessController" account.
func (inst *FinalizeOperation) SetRoleAccessControllerAccount(roleAccessController ag_solanago.PublicKey) *FinalizeOperation {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(roleAccessController)
	return inst
}

// GetRoleAccessControllerAccount gets the "roleAccessController" account.
func (inst *FinalizeOperation) GetRoleAccessControllerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *FinalizeOperation) SetAuthorityAccount(authority ag_solanago.PublicKey) *FinalizeOperation {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *FinalizeOperation) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

func (inst FinalizeOperation) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_FinalizeOperation,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst FinalizeOperation) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *FinalizeOperation) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.TimelockId == nil {
			return errors.New("TimelockId parameter is not set")
		}
		if inst.Id == nil {
			return errors.New("Id parameter is not set")
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
	}
	return nil
}

func (inst *FinalizeOperation) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("FinalizeOperation")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("TimelockId", *inst.TimelockId))
						paramsBranch.Child(ag_format.Param("        Id", *inst.Id))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("           operation", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("              config", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("roleAccessController", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("           authority", inst.AccountMetaSlice[3]))
					})
				})
		})
}

func (obj FinalizeOperation) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
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
	return nil
}
func (obj *FinalizeOperation) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
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
	return nil
}

// NewFinalizeOperationInstruction declares a new FinalizeOperation instruction with the provided parameters and accounts.
func NewFinalizeOperationInstruction(
	// Parameters:
	timelockId [32]uint8,
	id [32]uint8,
	// Accounts:
	operation ag_solanago.PublicKey,
	config ag_solanago.PublicKey,
	roleAccessController ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *FinalizeOperation {
	return NewFinalizeOperationInstructionBuilder().
		SetTimelockId(timelockId).
		SetId(id).
		SetOperationAccount(operation).
		SetConfigAccount(config).
		SetRoleAccessControllerAccount(roleAccessController).
		SetAuthorityAccount(authority)
}
