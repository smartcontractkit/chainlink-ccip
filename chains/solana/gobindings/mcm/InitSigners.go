// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package mcm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// InitSigners is the `initSigners` instruction.
type InitSigners struct {
	MultisigId   *[32]uint8
	TotalSigners *uint8

	// [0] = [] multisigConfig
	//
	// [1] = [WRITE] configSigners
	//
	// [2] = [WRITE, SIGNER] authority
	//
	// [3] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewInitSignersInstructionBuilder creates a new `InitSigners` instruction builder.
func NewInitSignersInstructionBuilder() *InitSigners {
	nd := &InitSigners{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetMultisigId sets the "multisigId" parameter.
func (inst *InitSigners) SetMultisigId(multisigId [32]uint8) *InitSigners {
	inst.MultisigId = &multisigId
	return inst
}

// SetTotalSigners sets the "totalSigners" parameter.
func (inst *InitSigners) SetTotalSigners(totalSigners uint8) *InitSigners {
	inst.TotalSigners = &totalSigners
	return inst
}

// SetMultisigConfigAccount sets the "multisigConfig" account.
func (inst *InitSigners) SetMultisigConfigAccount(multisigConfig ag_solanago.PublicKey) *InitSigners {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(multisigConfig)
	return inst
}

// GetMultisigConfigAccount gets the "multisigConfig" account.
func (inst *InitSigners) GetMultisigConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetConfigSignersAccount sets the "configSigners" account.
func (inst *InitSigners) SetConfigSignersAccount(configSigners ag_solanago.PublicKey) *InitSigners {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(configSigners).WRITE()
	return inst
}

// GetConfigSignersAccount gets the "configSigners" account.
func (inst *InitSigners) GetConfigSignersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *InitSigners) SetAuthorityAccount(authority ag_solanago.PublicKey) *InitSigners {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *InitSigners) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *InitSigners) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *InitSigners {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *InitSigners) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

func (inst InitSigners) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitSigners,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitSigners) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitSigners) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.MultisigId == nil {
			return errors.New("MultisigId parameter is not set")
		}
		if inst.TotalSigners == nil {
			return errors.New("TotalSigners parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.MultisigConfig is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ConfigSigners is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *InitSigners) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitSigners")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("  MultisigId", *inst.MultisigId))
						paramsBranch.Child(ag_format.Param("TotalSigners", *inst.TotalSigners))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("multisigConfig", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta(" configSigners", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("     authority", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta(" systemProgram", inst.AccountMetaSlice[3]))
					})
				})
		})
}

func (obj InitSigners) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MultisigId` param:
	err = encoder.Encode(obj.MultisigId)
	if err != nil {
		return err
	}
	// Serialize `TotalSigners` param:
	err = encoder.Encode(obj.TotalSigners)
	if err != nil {
		return err
	}
	return nil
}
func (obj *InitSigners) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MultisigId`:
	err = decoder.Decode(&obj.MultisigId)
	if err != nil {
		return err
	}
	// Deserialize `TotalSigners`:
	err = decoder.Decode(&obj.TotalSigners)
	if err != nil {
		return err
	}
	return nil
}

// NewInitSignersInstruction declares a new InitSigners instruction with the provided parameters and accounts.
func NewInitSignersInstruction(
	// Parameters:
	multisigId [32]uint8,
	totalSigners uint8,
	// Accounts:
	multisigConfig ag_solanago.PublicKey,
	configSigners ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *InitSigners {
	return NewInitSignersInstructionBuilder().
		SetMultisigId(multisigId).
		SetTotalSigners(totalSigners).
		SetMultisigConfigAccount(multisigConfig).
		SetConfigSignersAccount(configSigners).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram)
}
