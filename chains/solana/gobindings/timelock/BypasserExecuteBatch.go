// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package timelock

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// BypasserExecuteBatch is the `bypasserExecuteBatch` instruction.
type BypasserExecuteBatch struct {
	Id *[32]uint8

	// [0] = [WRITE] operation
	//
	// [1] = [] config
	//
	// [2] = [] timelockSigner
	//
	// [3] = [] roleAccessController
	//
	// [4] = [WRITE, SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewBypasserExecuteBatchInstructionBuilder creates a new `BypasserExecuteBatch` instruction builder.
func NewBypasserExecuteBatchInstructionBuilder() *BypasserExecuteBatch {
	nd := &BypasserExecuteBatch{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetId sets the "id" parameter.
func (inst *BypasserExecuteBatch) SetId(id [32]uint8) *BypasserExecuteBatch {
	inst.Id = &id
	return inst
}

// SetOperationAccount sets the "operation" account.
func (inst *BypasserExecuteBatch) SetOperationAccount(operation ag_solanago.PublicKey) *BypasserExecuteBatch {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(operation).WRITE()
	return inst
}

// GetOperationAccount gets the "operation" account.
func (inst *BypasserExecuteBatch) GetOperationAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetConfigAccount sets the "config" account.
func (inst *BypasserExecuteBatch) SetConfigAccount(config ag_solanago.PublicKey) *BypasserExecuteBatch {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *BypasserExecuteBatch) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetTimelockSignerAccount sets the "timelockSigner" account.
func (inst *BypasserExecuteBatch) SetTimelockSignerAccount(timelockSigner ag_solanago.PublicKey) *BypasserExecuteBatch {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(timelockSigner)
	return inst
}

// GetTimelockSignerAccount gets the "timelockSigner" account.
func (inst *BypasserExecuteBatch) GetTimelockSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetRoleAccessControllerAccount sets the "roleAccessController" account.
func (inst *BypasserExecuteBatch) SetRoleAccessControllerAccount(roleAccessController ag_solanago.PublicKey) *BypasserExecuteBatch {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(roleAccessController)
	return inst
}

// GetRoleAccessControllerAccount gets the "roleAccessController" account.
func (inst *BypasserExecuteBatch) GetRoleAccessControllerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *BypasserExecuteBatch) SetAuthorityAccount(authority ag_solanago.PublicKey) *BypasserExecuteBatch {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *BypasserExecuteBatch) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

func (inst BypasserExecuteBatch) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_BypasserExecuteBatch,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst BypasserExecuteBatch) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *BypasserExecuteBatch) Validate() error {
	// Check whether all (required) parameters are set:
	{
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
			return errors.New("accounts.TimelockSigner is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.RoleAccessController is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Authority is not set")
		}
	}
	return nil
}

func (inst *BypasserExecuteBatch) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("BypasserExecuteBatch")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Id", *inst.Id))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("           operation", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("              config", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("      timelockSigner", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("roleAccessController", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("           authority", inst.AccountMetaSlice[4]))
					})
				})
		})
}

func (obj BypasserExecuteBatch) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Id` param:
	err = encoder.Encode(obj.Id)
	if err != nil {
		return err
	}
	return nil
}
func (obj *BypasserExecuteBatch) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Id`:
	err = decoder.Decode(&obj.Id)
	if err != nil {
		return err
	}
	return nil
}

// NewBypasserExecuteBatchInstruction declares a new BypasserExecuteBatch instruction with the provided parameters and accounts.
func NewBypasserExecuteBatchInstruction(
	// Parameters:
	id [32]uint8,
	// Accounts:
	operation ag_solanago.PublicKey,
	config ag_solanago.PublicKey,
	timelockSigner ag_solanago.PublicKey,
	roleAccessController ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *BypasserExecuteBatch {
	return NewBypasserExecuteBatchInstructionBuilder().
		SetId(id).
		SetOperationAccount(operation).
		SetConfigAccount(config).
		SetTimelockSignerAccount(timelockSigner).
		SetRoleAccessControllerAccount(roleAccessController).
		SetAuthorityAccount(authority)
}
