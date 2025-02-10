// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package example_ccip_sender

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// InitChainConfig is the `init_chain_config` instruction.
type InitChainConfig struct {
	ChainSelector  *uint64
	Recipient      *[]byte
	ExtraArgsBytes *[]byte

	// [0] = [WRITE] state
	//
	// [1] = [WRITE] chain_config
	//
	// [2] = [WRITE, SIGNER] authority
	//
	// [3] = [] system_program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitChainConfigInstructionBuilder creates a new `InitChainConfig` instruction builder.
func NewInitChainConfigInstructionBuilder() *InitChainConfig {
	nd := &InitChainConfig{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetChainSelector sets the "_chain_selector" parameter.
func (inst *InitChainConfig) SetChainSelector(_chain_selector uint64) *InitChainConfig {
	inst.ChainSelector = &_chain_selector
	return inst
}

// SetRecipient sets the "recipient" parameter.
func (inst *InitChainConfig) SetRecipient(recipient []byte) *InitChainConfig {
	inst.Recipient = &recipient
	return inst
}

// SetExtraArgsBytes sets the "extra_args_bytes" parameter.
func (inst *InitChainConfig) SetExtraArgsBytes(extra_args_bytes []byte) *InitChainConfig {
	inst.ExtraArgsBytes = &extra_args_bytes
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *InitChainConfig) SetStateAccount(state ag_solanago.PublicKey) *InitChainConfig {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *InitChainConfig) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetChainConfigAccount sets the "chain_config" account.
func (inst *InitChainConfig) SetChainConfigAccount(chainConfig ag_solanago.PublicKey) *InitChainConfig {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(chainConfig).WRITE()
	return inst
}

// GetChainConfigAccount gets the "chain_config" account.
func (inst *InitChainConfig) GetChainConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *InitChainConfig) SetAuthorityAccount(authority ag_solanago.PublicKey) *InitChainConfig {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *InitChainConfig) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *InitChainConfig) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *InitChainConfig {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *InitChainConfig) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst InitChainConfig) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitChainConfig,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitChainConfig) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitChainConfig) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.ChainSelector == nil {
			return errors.New("ChainSelector parameter is not set")
		}
		if inst.Recipient == nil {
			return errors.New("Recipient parameter is not set")
		}
		if inst.ExtraArgsBytes == nil {
			return errors.New("ExtraArgsBytes parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ChainConfig is not set")
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

func (inst *InitChainConfig) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitChainConfig")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("   ChainSelector", *inst.ChainSelector))
						paramsBranch.Child(ag_format.Param("       Recipient", *inst.Recipient))
						paramsBranch.Child(ag_format.Param("  ExtraArgsBytes", *inst.ExtraArgsBytes))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("         state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("  chain_config", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("     authority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("system_program", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj InitChainConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ChainSelector` param:
	err = encoder.Encode(obj.ChainSelector)
	if err != nil {
		return err
	}
	// Serialize `Recipient` param:
	err = encoder.Encode(obj.Recipient)
	if err != nil {
		return err
	}
	// Serialize `ExtraArgsBytes` param:
	err = encoder.Encode(obj.ExtraArgsBytes)
	if err != nil {
		return err
	}
	return nil
}
func (obj *InitChainConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ChainSelector`:
	err = decoder.Decode(&obj.ChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `Recipient`:
	err = decoder.Decode(&obj.Recipient)
	if err != nil {
		return err
	}
	// Deserialize `ExtraArgsBytes`:
	err = decoder.Decode(&obj.ExtraArgsBytes)
	if err != nil {
		return err
	}
	return nil
}

// NewInitChainConfigInstruction declares a new InitChainConfig instruction with the provided parameters and accounts.
func NewInitChainConfigInstruction(
	// Parameters:
	_chain_selector uint64,
	recipient []byte,
	extra_args_bytes []byte,
	// Accounts:
	state ag_solanago.PublicKey,
	chainConfig ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *InitChainConfig {
	return NewInitChainConfigInstructionBuilder().
		SetChainSelector(_chain_selector).
		SetRecipient(recipient).
		SetExtraArgsBytes(extra_args_bytes).
		SetStateAccount(state).
		SetChainConfigAccount(chainConfig).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram)
}
