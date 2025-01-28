// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package timelock

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// FinalizeBypasserOperation is the `finalizeBypasserOperation` instruction.
type FinalizeBypasserOperation struct {
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

// NewFinalizeBypasserOperationInstructionBuilder creates a new `FinalizeBypasserOperation` instruction builder.
func NewFinalizeBypasserOperationInstructionBuilder() *FinalizeBypasserOperation {
	nd := &FinalizeBypasserOperation{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetTimelockId sets the "timelockId" parameter.
func (inst *FinalizeBypasserOperation) SetTimelockId(timelockId [32]uint8) *FinalizeBypasserOperation {
	inst.TimelockId = &timelockId
	return inst
}

// SetId sets the "id" parameter.
func (inst *FinalizeBypasserOperation) SetId(id [32]uint8) *FinalizeBypasserOperation {
	inst.Id = &id
	return inst
}

// SetOperationAccount sets the "operation" account.
func (inst *FinalizeBypasserOperation) SetOperationAccount(operation ag_solanago.PublicKey) *FinalizeBypasserOperation {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(operation).WRITE()
	return inst
}

// GetOperationAccount gets the "operation" account.
func (inst *FinalizeBypasserOperation) GetOperationAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetConfigAccount sets the "config" account.
func (inst *FinalizeBypasserOperation) SetConfigAccount(config ag_solanago.PublicKey) *FinalizeBypasserOperation {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *FinalizeBypasserOperation) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetRoleAccessControllerAccount sets the "roleAccessController" account.
func (inst *FinalizeBypasserOperation) SetRoleAccessControllerAccount(roleAccessController ag_solanago.PublicKey) *FinalizeBypasserOperation {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(roleAccessController)
	return inst
}

// GetRoleAccessControllerAccount gets the "roleAccessController" account.
func (inst *FinalizeBypasserOperation) GetRoleAccessControllerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *FinalizeBypasserOperation) SetAuthorityAccount(authority ag_solanago.PublicKey) *FinalizeBypasserOperation {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *FinalizeBypasserOperation) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

func (inst FinalizeBypasserOperation) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_FinalizeBypasserOperation,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst FinalizeBypasserOperation) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *FinalizeBypasserOperation) Validate() error {
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

func (inst *FinalizeBypasserOperation) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("FinalizeBypasserOperation")).
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

func (obj FinalizeBypasserOperation) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
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
func (obj *FinalizeBypasserOperation) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
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

// NewFinalizeBypasserOperationInstruction declares a new FinalizeBypasserOperation instruction with the provided parameters and accounts.
func NewFinalizeBypasserOperationInstruction(
	// Parameters:
	timelockId [32]uint8,
	id [32]uint8,
	// Accounts:
	operation ag_solanago.PublicKey,
	config ag_solanago.PublicKey,
	roleAccessController ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *FinalizeBypasserOperation {
	return NewFinalizeBypasserOperationInstructionBuilder().
		SetTimelockId(timelockId).
		SetId(id).
		SetOperationAccount(operation).
		SetConfigAccount(config).
		SetRoleAccessControllerAccount(roleAccessController).
		SetAuthorityAccount(authority)
}
