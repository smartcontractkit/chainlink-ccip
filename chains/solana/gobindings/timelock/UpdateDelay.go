// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package timelock

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdateDelay is the `updateDelay` instruction.
type UpdateDelay struct {
	Delay *uint64

	// [0] = [WRITE] config
	//
	// [1] = [SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewUpdateDelayInstructionBuilder creates a new `UpdateDelay` instruction builder.
func NewUpdateDelayInstructionBuilder() *UpdateDelay {
	nd := &UpdateDelay{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetDelay sets the "delay" parameter.
func (inst *UpdateDelay) SetDelay(delay uint64) *UpdateDelay {
	inst.Delay = &delay
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *UpdateDelay) SetConfigAccount(config ag_solanago.PublicKey) *UpdateDelay {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config).WRITE()
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *UpdateDelay) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *UpdateDelay) SetAuthorityAccount(authority ag_solanago.PublicKey) *UpdateDelay {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *UpdateDelay) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

func (inst UpdateDelay) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateDelay,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateDelay) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateDelay) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Delay == nil {
			return errors.New("Delay parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
	}
	return nil
}

func (inst *UpdateDelay) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateDelay")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Delay", *inst.Delay))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("   config", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("authority", inst.AccountMetaSlice[1]))
					})
				})
		})
}

func (obj UpdateDelay) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Delay` param:
	err = encoder.Encode(obj.Delay)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdateDelay) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Delay`:
	err = decoder.Decode(&obj.Delay)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateDelayInstruction declares a new UpdateDelay instruction with the provided parameters and accounts.
func NewUpdateDelayInstruction(
	// Parameters:
	delay uint64,
	// Accounts:
	config ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *UpdateDelay {
	return NewUpdateDelayInstructionBuilder().
		SetDelay(delay).
		SetConfigAccount(config).
		SetAuthorityAccount(authority)
}