// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ping_pong_demo

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SetCounterpart is the `setCounterpart` instruction.
type SetCounterpart struct {
	CounterpartChainSelector *uint64
	CounterpartAddress       *[64]uint8

	// [0] = [WRITE] config
	//
	// [1] = [WRITE, SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewSetCounterpartInstructionBuilder creates a new `SetCounterpart` instruction builder.
func NewSetCounterpartInstructionBuilder() *SetCounterpart {
	nd := &SetCounterpart{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetCounterpartChainSelector sets the "counterpartChainSelector" parameter.
func (inst *SetCounterpart) SetCounterpartChainSelector(counterpartChainSelector uint64) *SetCounterpart {
	inst.CounterpartChainSelector = &counterpartChainSelector
	return inst
}

// SetCounterpartAddress sets the "counterpartAddress" parameter.
func (inst *SetCounterpart) SetCounterpartAddress(counterpartAddress [64]uint8) *SetCounterpart {
	inst.CounterpartAddress = &counterpartAddress
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *SetCounterpart) SetConfigAccount(config ag_solanago.PublicKey) *SetCounterpart {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config).WRITE()
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *SetCounterpart) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *SetCounterpart) SetAuthorityAccount(authority ag_solanago.PublicKey) *SetCounterpart {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *SetCounterpart) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

func (inst SetCounterpart) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SetCounterpart,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetCounterpart) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetCounterpart) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.CounterpartChainSelector == nil {
			return errors.New("CounterpartChainSelector parameter is not set")
		}
		if inst.CounterpartAddress == nil {
			return errors.New("CounterpartAddress parameter is not set")
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

func (inst *SetCounterpart) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetCounterpart")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("CounterpartChainSelector", *inst.CounterpartChainSelector))
						paramsBranch.Child(ag_format.Param("      CounterpartAddress", *inst.CounterpartAddress))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("   config", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("authority", inst.AccountMetaSlice[1]))
					})
				})
		})
}

func (obj SetCounterpart) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `CounterpartChainSelector` param:
	err = encoder.Encode(obj.CounterpartChainSelector)
	if err != nil {
		return err
	}
	// Serialize `CounterpartAddress` param:
	err = encoder.Encode(obj.CounterpartAddress)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SetCounterpart) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `CounterpartChainSelector`:
	err = decoder.Decode(&obj.CounterpartChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `CounterpartAddress`:
	err = decoder.Decode(&obj.CounterpartAddress)
	if err != nil {
		return err
	}
	return nil
}

// NewSetCounterpartInstruction declares a new SetCounterpart instruction with the provided parameters and accounts.
func NewSetCounterpartInstruction(
	// Parameters:
	counterpartChainSelector uint64,
	counterpartAddress [64]uint8,
	// Accounts:
	config ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *SetCounterpart {
	return NewSetCounterpartInstructionBuilder().
		SetCounterpartChainSelector(counterpartChainSelector).
		SetCounterpartAddress(counterpartAddress).
		SetConfigAccount(config).
		SetAuthorityAccount(authority)
}
