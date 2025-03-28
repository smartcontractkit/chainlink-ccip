// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ping_pong_demo

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SetPaused is the `setPaused` instruction.
type SetPaused struct {
	CounterpartChainSelector *uint64
	Pause                    *bool

	// [0] = [] globalConfig
	//
	// [1] = [WRITE] config
	//
	// [2] = [WRITE, SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewSetPausedInstructionBuilder creates a new `SetPaused` instruction builder.
func NewSetPausedInstructionBuilder() *SetPaused {
	nd := &SetPaused{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetCounterpartChainSelector sets the "counterpartChainSelector" parameter.
func (inst *SetPaused) SetCounterpartChainSelector(counterpartChainSelector uint64) *SetPaused {
	inst.CounterpartChainSelector = &counterpartChainSelector
	return inst
}

// SetPause sets the "pause" parameter.
func (inst *SetPaused) SetPause(pause bool) *SetPaused {
	inst.Pause = &pause
	return inst
}

// SetGlobalConfigAccount sets the "globalConfig" account.
func (inst *SetPaused) SetGlobalConfigAccount(globalConfig ag_solanago.PublicKey) *SetPaused {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(globalConfig)
	return inst
}

// GetGlobalConfigAccount gets the "globalConfig" account.
func (inst *SetPaused) GetGlobalConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetConfigAccount sets the "config" account.
func (inst *SetPaused) SetConfigAccount(config ag_solanago.PublicKey) *SetPaused {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(config).WRITE()
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *SetPaused) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *SetPaused) SetAuthorityAccount(authority ag_solanago.PublicKey) *SetPaused {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *SetPaused) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

func (inst SetPaused) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SetPaused,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetPaused) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetPaused) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.CounterpartChainSelector == nil {
			return errors.New("CounterpartChainSelector parameter is not set")
		}
		if inst.Pause == nil {
			return errors.New("Pause parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.GlobalConfig is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Authority is not set")
		}
	}
	return nil
}

func (inst *SetPaused) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetPaused")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("CounterpartChainSelector", *inst.CounterpartChainSelector))
						paramsBranch.Child(ag_format.Param("                   Pause", *inst.Pause))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("globalConfig", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("      config", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("   authority", inst.AccountMetaSlice[2]))
					})
				})
		})
}

func (obj SetPaused) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `CounterpartChainSelector` param:
	err = encoder.Encode(obj.CounterpartChainSelector)
	if err != nil {
		return err
	}
	// Serialize `Pause` param:
	err = encoder.Encode(obj.Pause)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SetPaused) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `CounterpartChainSelector`:
	err = decoder.Decode(&obj.CounterpartChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `Pause`:
	err = decoder.Decode(&obj.Pause)
	if err != nil {
		return err
	}
	return nil
}

// NewSetPausedInstruction declares a new SetPaused instruction with the provided parameters and accounts.
func NewSetPausedInstruction(
	// Parameters:
	counterpartChainSelector uint64,
	pause bool,
	// Accounts:
	globalConfig ag_solanago.PublicKey,
	config ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *SetPaused {
	return NewSetPausedInstructionBuilder().
		SetCounterpartChainSelector(counterpartChainSelector).
		SetPause(pause).
		SetGlobalConfigAccount(globalConfig).
		SetConfigAccount(config).
		SetAuthorityAccount(authority)
}
