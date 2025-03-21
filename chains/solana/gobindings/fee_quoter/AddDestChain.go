// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package fee_quoter

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Adds a new destination chain selector to the fee quoter.
//
// The Admin needs to add any new chain supported.
// When adding a new chain, the Admin needs to specify if it's enabled or not.
//
// # Arguments
//
// * `ctx` - The context containing the accounts required for adding the chain selector.
// * `chain_selector` - The new chain selector to be added.
// * `dest_chain_config` - The configuration for the chain as destination.
type AddDestChain struct {
	ChainSelector   *uint64
	DestChainConfig *DestChainConfig

	// [0] = [] config
	//
	// [1] = [WRITE] destChain
	//
	// [2] = [WRITE, SIGNER] authority
	//
	// [3] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewAddDestChainInstructionBuilder creates a new `AddDestChain` instruction builder.
func NewAddDestChainInstructionBuilder() *AddDestChain {
	nd := &AddDestChain{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetChainSelector sets the "chainSelector" parameter.
func (inst *AddDestChain) SetChainSelector(chainSelector uint64) *AddDestChain {
	inst.ChainSelector = &chainSelector
	return inst
}

// SetDestChainConfig sets the "destChainConfig" parameter.
func (inst *AddDestChain) SetDestChainConfig(destChainConfig DestChainConfig) *AddDestChain {
	inst.DestChainConfig = &destChainConfig
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *AddDestChain) SetConfigAccount(config ag_solanago.PublicKey) *AddDestChain {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *AddDestChain) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetDestChainAccount sets the "destChain" account.
func (inst *AddDestChain) SetDestChainAccount(destChain ag_solanago.PublicKey) *AddDestChain {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(destChain).WRITE()
	return inst
}

// GetDestChainAccount gets the "destChain" account.
func (inst *AddDestChain) GetDestChainAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *AddDestChain) SetAuthorityAccount(authority ag_solanago.PublicKey) *AddDestChain {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *AddDestChain) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *AddDestChain) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *AddDestChain {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *AddDestChain) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

func (inst AddDestChain) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_AddDestChain,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AddDestChain) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AddDestChain) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.ChainSelector == nil {
			return errors.New("ChainSelector parameter is not set")
		}
		if inst.DestChainConfig == nil {
			return errors.New("DestChainConfig parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.DestChain is not set")
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

func (inst *AddDestChain) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AddDestChain")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("  ChainSelector", *inst.ChainSelector))
						paramsBranch.Child(ag_format.Param("DestChainConfig", *inst.DestChainConfig))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       config", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("    destChain", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("    authority", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice[3]))
					})
				})
		})
}

func (obj AddDestChain) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ChainSelector` param:
	err = encoder.Encode(obj.ChainSelector)
	if err != nil {
		return err
	}
	// Serialize `DestChainConfig` param:
	err = encoder.Encode(obj.DestChainConfig)
	if err != nil {
		return err
	}
	return nil
}
func (obj *AddDestChain) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ChainSelector`:
	err = decoder.Decode(&obj.ChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `DestChainConfig`:
	err = decoder.Decode(&obj.DestChainConfig)
	if err != nil {
		return err
	}
	return nil
}

// NewAddDestChainInstruction declares a new AddDestChain instruction with the provided parameters and accounts.
func NewAddDestChainInstruction(
	// Parameters:
	chainSelector uint64,
	destChainConfig DestChainConfig,
	// Accounts:
	config ag_solanago.PublicKey,
	destChain ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *AddDestChain {
	return NewAddDestChainInstructionBuilder().
		SetChainSelector(chainSelector).
		SetDestChainConfig(destChainConfig).
		SetConfigAccount(config).
		SetDestChainAccount(destChain).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram)
}
