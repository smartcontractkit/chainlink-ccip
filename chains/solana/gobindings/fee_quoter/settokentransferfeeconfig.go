// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package fee_quoter

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Sets the token transfer fee configuration for a particular token when it's transferred to a particular dest chain.
// It is an upsert, initializing the per-chain-per-token config account if it doesn't exist
// and overwriting it if it does.
//
// Only the Admin can perform this operation.
//
// # Arguments
//
// * `ctx` - The context containing the accounts required for setting the token billing configuration.
// * `chain_selector` - The chain selector.
// * `mint` - The public key of the token mint.
// * `cfg` - The token transfer fee configuration.
type SetTokenTransferFeeConfig struct {
	ChainSelector *uint64
	Mint          *ag_solanago.PublicKey
	Cfg           *TokenTransferFeeConfig

	// [0] = [] config
	//
	// [1] = [WRITE] per_chain_per_token_config
	//
	// [2] = [WRITE, SIGNER] authority
	//
	// [3] = [] system_program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSetTokenTransferFeeConfigInstructionBuilder creates a new `SetTokenTransferFeeConfig` instruction builder.
func NewSetTokenTransferFeeConfigInstructionBuilder() *SetTokenTransferFeeConfig {
	nd := &SetTokenTransferFeeConfig{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetChainSelector sets the "chain_selector" parameter.
func (inst *SetTokenTransferFeeConfig) SetChainSelector(chain_selector uint64) *SetTokenTransferFeeConfig {
	inst.ChainSelector = &chain_selector
	return inst
}

// SetMint sets the "mint" parameter.
func (inst *SetTokenTransferFeeConfig) SetMint(mint ag_solanago.PublicKey) *SetTokenTransferFeeConfig {
	inst.Mint = &mint
	return inst
}

// SetCfg sets the "cfg" parameter.
func (inst *SetTokenTransferFeeConfig) SetCfg(cfg TokenTransferFeeConfig) *SetTokenTransferFeeConfig {
	inst.Cfg = &cfg
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *SetTokenTransferFeeConfig) SetConfigAccount(config ag_solanago.PublicKey) *SetTokenTransferFeeConfig {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *SetTokenTransferFeeConfig) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetPerChainPerTokenConfigAccount sets the "per_chain_per_token_config" account.
func (inst *SetTokenTransferFeeConfig) SetPerChainPerTokenConfigAccount(perChainPerTokenConfig ag_solanago.PublicKey) *SetTokenTransferFeeConfig {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(perChainPerTokenConfig).WRITE()
	return inst
}

// GetPerChainPerTokenConfigAccount gets the "per_chain_per_token_config" account.
func (inst *SetTokenTransferFeeConfig) GetPerChainPerTokenConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *SetTokenTransferFeeConfig) SetAuthorityAccount(authority ag_solanago.PublicKey) *SetTokenTransferFeeConfig {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *SetTokenTransferFeeConfig) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *SetTokenTransferFeeConfig) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *SetTokenTransferFeeConfig {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *SetTokenTransferFeeConfig) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst SetTokenTransferFeeConfig) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SetTokenTransferFeeConfig,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetTokenTransferFeeConfig) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetTokenTransferFeeConfig) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.ChainSelector == nil {
			return errors.New("ChainSelector parameter is not set")
		}
		if inst.Mint == nil {
			return errors.New("Mint parameter is not set")
		}
		if inst.Cfg == nil {
			return errors.New("Cfg parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.PerChainPerTokenConfig is not set")
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

func (inst *SetTokenTransferFeeConfig) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetTokenTransferFeeConfig")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param(" ChainSelector", *inst.ChainSelector))
						paramsBranch.Child(ag_format.Param("          Mint", *inst.Mint))
						paramsBranch.Child(ag_format.Param("           Cfg", *inst.Cfg))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                    config", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("per_chain_per_token_config", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                 authority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("            system_program", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj SetTokenTransferFeeConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ChainSelector` param:
	err = encoder.Encode(obj.ChainSelector)
	if err != nil {
		return err
	}
	// Serialize `Mint` param:
	err = encoder.Encode(obj.Mint)
	if err != nil {
		return err
	}
	// Serialize `Cfg` param:
	err = encoder.Encode(obj.Cfg)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SetTokenTransferFeeConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ChainSelector`:
	err = decoder.Decode(&obj.ChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `Mint`:
	err = decoder.Decode(&obj.Mint)
	if err != nil {
		return err
	}
	// Deserialize `Cfg`:
	err = decoder.Decode(&obj.Cfg)
	if err != nil {
		return err
	}
	return nil
}

// NewSetTokenTransferFeeConfigInstruction declares a new SetTokenTransferFeeConfig instruction with the provided parameters and accounts.
func NewSetTokenTransferFeeConfigInstruction(
	// Parameters:
	chain_selector uint64,
	mint ag_solanago.PublicKey,
	cfg TokenTransferFeeConfig,
	// Accounts:
	config ag_solanago.PublicKey,
	perChainPerTokenConfig ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *SetTokenTransferFeeConfig {
	return NewSetTokenTransferFeeConfigInstructionBuilder().
		SetChainSelector(chain_selector).
		SetMint(mint).
		SetCfg(cfg).
		SetConfigAccount(config).
		SetPerChainPerTokenConfigAccount(perChainPerTokenConfig).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram)
}
