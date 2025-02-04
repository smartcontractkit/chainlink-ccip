// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package timelock

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// InitializeBypasserInstruction is the `initializeBypasserInstruction` instruction.
type InitializeBypasserInstruction struct {
	TimelockId *[32]uint8
	Id         *[32]uint8
	ProgramId  *ag_solanago.PublicKey
	Accounts   *[]InstructionAccount

	// [0] = [WRITE] operation
	//
	// [1] = [] config
	//
	// [2] = [] roleAccessController
	//
	// [3] = [WRITE, SIGNER] authority
	//
	// [4] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewInitializeBypasserInstructionInstructionBuilder creates a new `InitializeBypasserInstruction` instruction builder.
func NewInitializeBypasserInstructionInstructionBuilder() *InitializeBypasserInstruction {
	nd := &InitializeBypasserInstruction{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetTimelockId sets the "timelockId" parameter.
func (inst *InitializeBypasserInstruction) SetTimelockId(timelockId [32]uint8) *InitializeBypasserInstruction {
	inst.TimelockId = &timelockId
	return inst
}

// SetId sets the "id" parameter.
func (inst *InitializeBypasserInstruction) SetId(id [32]uint8) *InitializeBypasserInstruction {
	inst.Id = &id
	return inst
}

// SetProgramId sets the "programId" parameter.
func (inst *InitializeBypasserInstruction) SetProgramId(programId ag_solanago.PublicKey) *InitializeBypasserInstruction {
	inst.ProgramId = &programId
	return inst
}

// SetAccounts sets the "accounts" parameter.
func (inst *InitializeBypasserInstruction) SetAccounts(accounts []InstructionAccount) *InitializeBypasserInstruction {
	inst.Accounts = &accounts
	return inst
}

// SetOperationAccount sets the "operation" account.
func (inst *InitializeBypasserInstruction) SetOperationAccount(operation ag_solanago.PublicKey) *InitializeBypasserInstruction {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(operation).WRITE()
	return inst
}

// GetOperationAccount gets the "operation" account.
func (inst *InitializeBypasserInstruction) GetOperationAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetConfigAccount sets the "config" account.
func (inst *InitializeBypasserInstruction) SetConfigAccount(config ag_solanago.PublicKey) *InitializeBypasserInstruction {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *InitializeBypasserInstruction) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetRoleAccessControllerAccount sets the "roleAccessController" account.
func (inst *InitializeBypasserInstruction) SetRoleAccessControllerAccount(roleAccessController ag_solanago.PublicKey) *InitializeBypasserInstruction {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(roleAccessController)
	return inst
}

// GetRoleAccessControllerAccount gets the "roleAccessController" account.
func (inst *InitializeBypasserInstruction) GetRoleAccessControllerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *InitializeBypasserInstruction) SetAuthorityAccount(authority ag_solanago.PublicKey) *InitializeBypasserInstruction {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *InitializeBypasserInstruction) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *InitializeBypasserInstruction) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *InitializeBypasserInstruction {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *InitializeBypasserInstruction) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

func (inst InitializeBypasserInstruction) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitializeBypasserInstruction,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitializeBypasserInstruction) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitializeBypasserInstruction) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.TimelockId == nil {
			return errors.New("TimelockId parameter is not set")
		}
		if inst.Id == nil {
			return errors.New("Id parameter is not set")
		}
		if inst.ProgramId == nil {
			return errors.New("ProgramId parameter is not set")
		}
		if inst.Accounts == nil {
			return errors.New("Accounts parameter is not set")
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

func (inst *InitializeBypasserInstruction) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitializeBypasserInstruction")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=4]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("TimelockId", *inst.TimelockId))
						paramsBranch.Child(ag_format.Param("        Id", *inst.Id))
						paramsBranch.Child(ag_format.Param(" ProgramId", *inst.ProgramId))
						paramsBranch.Child(ag_format.Param("  Accounts", *inst.Accounts))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("           operation", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("              config", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("roleAccessController", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("           authority", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("       systemProgram", inst.AccountMetaSlice[4]))
					})
				})
		})
}

func (obj InitializeBypasserInstruction) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
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
	// Serialize `ProgramId` param:
	err = encoder.Encode(obj.ProgramId)
	if err != nil {
		return err
	}
	// Serialize `Accounts` param:
	err = encoder.Encode(obj.Accounts)
	if err != nil {
		return err
	}
	return nil
}
func (obj *InitializeBypasserInstruction) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
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
	// Deserialize `ProgramId`:
	err = decoder.Decode(&obj.ProgramId)
	if err != nil {
		return err
	}
	// Deserialize `Accounts`:
	err = decoder.Decode(&obj.Accounts)
	if err != nil {
		return err
	}
	return nil
}

// NewInitializeBypasserInstructionInstruction declares a new InitializeBypasserInstruction instruction with the provided parameters and accounts.
func NewInitializeBypasserInstructionInstruction(
	// Parameters:
	timelockId [32]uint8,
	id [32]uint8,
	programId ag_solanago.PublicKey,
	accounts []InstructionAccount,
	// Accounts:
	operation ag_solanago.PublicKey,
	config ag_solanago.PublicKey,
	roleAccessController ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *InitializeBypasserInstruction {
	return NewInitializeBypasserInstructionInstructionBuilder().
		SetTimelockId(timelockId).
		SetId(id).
		SetProgramId(programId).
		SetAccounts(accounts).
		SetOperationAccount(operation).
		SetConfigAccount(config).
		SetRoleAccessControllerAccount(roleAccessController).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram)
}
