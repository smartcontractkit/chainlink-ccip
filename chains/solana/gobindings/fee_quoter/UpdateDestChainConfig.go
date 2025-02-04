// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package fee_quoter

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Updates the configuration of the destination chain selector.
//
// The Admin is the only one able to update the destination chain config.
//
// # Arguments
//
// * `ctx` - The context containing the accounts required for updating the chain selector.
// * `chain_selector` - The destination chain selector to be updated.
// * `dest_chain_config` - The new configuration for the destination chain.
type UpdateDestChainConfig struct {
	ChainSelector   *uint64
	DestChainConfig *DestChainConfig

	// [0] = [] config
	//
	// [1] = [WRITE] destChain
	//
	// [2] = [WRITE, SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewUpdateDestChainConfigInstructionBuilder creates a new `UpdateDestChainConfig` instruction builder.
func NewUpdateDestChainConfigInstructionBuilder() *UpdateDestChainConfig {
	nd := &UpdateDestChainConfig{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetChainSelector sets the "chainSelector" parameter.
func (inst *UpdateDestChainConfig) SetChainSelector(chainSelector uint64) *UpdateDestChainConfig {
	inst.ChainSelector = &chainSelector
	return inst
}

// SetDestChainConfig sets the "destChainConfig" parameter.
func (inst *UpdateDestChainConfig) SetDestChainConfig(destChainConfig DestChainConfig) *UpdateDestChainConfig {
	inst.DestChainConfig = &destChainConfig
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *UpdateDestChainConfig) SetConfigAccount(config ag_solanago.PublicKey) *UpdateDestChainConfig {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *UpdateDestChainConfig) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetDestChainAccount sets the "destChain" account.
func (inst *UpdateDestChainConfig) SetDestChainAccount(destChain ag_solanago.PublicKey) *UpdateDestChainConfig {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(destChain).WRITE()
	return inst
}

// GetDestChainAccount gets the "destChain" account.
func (inst *UpdateDestChainConfig) GetDestChainAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *UpdateDestChainConfig) SetAuthorityAccount(authority ag_solanago.PublicKey) *UpdateDestChainConfig {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *UpdateDestChainConfig) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

func (inst UpdateDestChainConfig) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateDestChainConfig,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateDestChainConfig) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateDestChainConfig) Validate() error {
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
	}
	return nil
}

func (inst *UpdateDestChainConfig) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateDestChainConfig")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("  ChainSelector", *inst.ChainSelector))
						paramsBranch.Child(ag_format.Param("DestChainConfig", *inst.DestChainConfig))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("   config", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("destChain", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("authority", inst.AccountMetaSlice[2]))
					})
				})
		})
}

func (obj UpdateDestChainConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
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
func (obj *UpdateDestChainConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
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

// NewUpdateDestChainConfigInstruction declares a new UpdateDestChainConfig instruction with the provided parameters and accounts.
func NewUpdateDestChainConfigInstruction(
	// Parameters:
	chainSelector uint64,
	destChainConfig DestChainConfig,
	// Accounts:
	config ag_solanago.PublicKey,
	destChain ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *UpdateDestChainConfig {
	return NewUpdateDestChainConfigInstructionBuilder().
		SetChainSelector(chainSelector).
		SetDestChainConfig(destChainConfig).
		SetConfigAccount(config).
		SetDestChainAccount(destChain).
		SetAuthorityAccount(authority)
}
