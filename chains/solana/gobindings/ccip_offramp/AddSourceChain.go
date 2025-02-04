// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ccip_offramp

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Config //
// Adds a new source chain selector with its config to the offramp.
//
// The Admin needs to add any new chain supported.
// When adding a new chain, the Admin needs to specify if it's enabled or not.
//
// # Arguments
type AddSourceChain struct {
	NewChainSelector  *uint64
	SourceChainConfig *SourceChainConfig

	// [0] = [WRITE] sourceChainState
	// ··········· Adding a chain selector implies initializing the state for a new chain,
	// ··········· hence the need to initialize two accounts.
	//
	// [1] = [] config
	//
	// [2] = [WRITE, SIGNER] authority
	//
	// [3] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewAddSourceChainInstructionBuilder creates a new `AddSourceChain` instruction builder.
func NewAddSourceChainInstructionBuilder() *AddSourceChain {
	nd := &AddSourceChain{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetNewChainSelector sets the "newChainSelector" parameter.
func (inst *AddSourceChain) SetNewChainSelector(newChainSelector uint64) *AddSourceChain {
	inst.NewChainSelector = &newChainSelector
	return inst
}

// SetSourceChainConfig sets the "sourceChainConfig" parameter.
func (inst *AddSourceChain) SetSourceChainConfig(sourceChainConfig SourceChainConfig) *AddSourceChain {
	inst.SourceChainConfig = &sourceChainConfig
	return inst
}

// SetSourceChainStateAccount sets the "sourceChainState" account.
// Adding a chain selector implies initializing the state for a new chain,
// hence the need to initialize two accounts.
func (inst *AddSourceChain) SetSourceChainStateAccount(sourceChainState ag_solanago.PublicKey) *AddSourceChain {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(sourceChainState).WRITE()
	return inst
}

// GetSourceChainStateAccount gets the "sourceChainState" account.
// Adding a chain selector implies initializing the state for a new chain,
// hence the need to initialize two accounts.
func (inst *AddSourceChain) GetSourceChainStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetConfigAccount sets the "config" account.
func (inst *AddSourceChain) SetConfigAccount(config ag_solanago.PublicKey) *AddSourceChain {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *AddSourceChain) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *AddSourceChain) SetAuthorityAccount(authority ag_solanago.PublicKey) *AddSourceChain {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *AddSourceChain) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *AddSourceChain) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *AddSourceChain {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *AddSourceChain) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

func (inst AddSourceChain) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_AddSourceChain,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AddSourceChain) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AddSourceChain) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.NewChainSelector == nil {
			return errors.New("NewChainSelector parameter is not set")
		}
		if inst.SourceChainConfig == nil {
			return errors.New("SourceChainConfig parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.SourceChainState is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Config is not set")
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

func (inst *AddSourceChain) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AddSourceChain")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param(" NewChainSelector", *inst.NewChainSelector))
						paramsBranch.Child(ag_format.Param("SourceChainConfig", *inst.SourceChainConfig))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("sourceChainState", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("          config", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("       authority", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("   systemProgram", inst.AccountMetaSlice[3]))
					})
				})
		})
}

func (obj AddSourceChain) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NewChainSelector` param:
	err = encoder.Encode(obj.NewChainSelector)
	if err != nil {
		return err
	}
	// Serialize `SourceChainConfig` param:
	err = encoder.Encode(obj.SourceChainConfig)
	if err != nil {
		return err
	}
	return nil
}
func (obj *AddSourceChain) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NewChainSelector`:
	err = decoder.Decode(&obj.NewChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `SourceChainConfig`:
	err = decoder.Decode(&obj.SourceChainConfig)
	if err != nil {
		return err
	}
	return nil
}

// NewAddSourceChainInstruction declares a new AddSourceChain instruction with the provided parameters and accounts.
func NewAddSourceChainInstruction(
	// Parameters:
	newChainSelector uint64,
	sourceChainConfig SourceChainConfig,
	// Accounts:
	sourceChainState ag_solanago.PublicKey,
	config ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *AddSourceChain {
	return NewAddSourceChainInstructionBuilder().
		SetNewChainSelector(newChainSelector).
		SetSourceChainConfig(sourceChainConfig).
		SetSourceChainStateAccount(sourceChainState).
		SetConfigAccount(config).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram)
}
