// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package mcm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// FinalizeSigners is the `finalizeSigners` instruction.
type FinalizeSigners struct {
	MultisigName *[32]uint8

	// [0] = [] multisigConfig
	//
	// [1] = [WRITE] configSigners
	//
	// [2] = [WRITE, SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewFinalizeSignersInstructionBuilder creates a new `FinalizeSigners` instruction builder.
func NewFinalizeSignersInstructionBuilder() *FinalizeSigners {
	nd := &FinalizeSigners{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetMultisigName sets the "multisigName" parameter.
func (inst *FinalizeSigners) SetMultisigName(multisigName [32]uint8) *FinalizeSigners {
	inst.MultisigName = &multisigName
	return inst
}

// SetMultisigConfigAccount sets the "multisigConfig" account.
func (inst *FinalizeSigners) SetMultisigConfigAccount(multisigConfig ag_solanago.PublicKey) *FinalizeSigners {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(multisigConfig)
	return inst
}

// GetMultisigConfigAccount gets the "multisigConfig" account.
func (inst *FinalizeSigners) GetMultisigConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetConfigSignersAccount sets the "configSigners" account.
func (inst *FinalizeSigners) SetConfigSignersAccount(configSigners ag_solanago.PublicKey) *FinalizeSigners {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(configSigners).WRITE()
	return inst
}

// GetConfigSignersAccount gets the "configSigners" account.
func (inst *FinalizeSigners) GetConfigSignersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *FinalizeSigners) SetAuthorityAccount(authority ag_solanago.PublicKey) *FinalizeSigners {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *FinalizeSigners) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

func (inst FinalizeSigners) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_FinalizeSigners,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst FinalizeSigners) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *FinalizeSigners) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.MultisigName == nil {
			return errors.New("MultisigName parameter is not set")
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
	}
	return nil
}

func (inst *FinalizeSigners) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("FinalizeSigners")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("MultisigName", *inst.MultisigName))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("multisigConfig", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta(" configSigners", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("     authority", inst.AccountMetaSlice[2]))
					})
				})
		})
}

func (obj FinalizeSigners) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MultisigName` param:
	err = encoder.Encode(obj.MultisigName)
	if err != nil {
		return err
	}
	return nil
}
func (obj *FinalizeSigners) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MultisigName`:
	err = decoder.Decode(&obj.MultisigName)
	if err != nil {
		return err
	}
	return nil
}

// NewFinalizeSignersInstruction declares a new FinalizeSigners instruction with the provided parameters and accounts.
func NewFinalizeSignersInstruction(
	// Parameters:
	multisigName [32]uint8,
	// Accounts:
	multisigConfig ag_solanago.PublicKey,
	configSigners ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *FinalizeSigners {
	return NewFinalizeSignersInstructionBuilder().
		SetMultisigName(multisigName).
		SetMultisigConfigAccount(multisigConfig).
		SetConfigSignersAccount(configSigners).
		SetAuthorityAccount(authority)
}
