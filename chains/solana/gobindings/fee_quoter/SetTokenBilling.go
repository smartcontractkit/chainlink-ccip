// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package fee_quoter

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Sets the token billing configuration.
// It is an upsert, initializing the token billing config account if it doesn't exist
// and overwriting it if it does.
//
// Only the Admin can set the token billing configuration.
//
// # Arguments
//
// * `ctx` - The context containing the accounts required for setting the token billing configuration.
// * `chain_selector` - The chain selector.
// * `mint` - The public key of the token mint.
// * `cfg` - The token billing configuration.
type SetTokenBilling struct {
	ChainSelector *uint64
	Mint          *ag_solanago.PublicKey
	Cfg           *TokenBilling

	// [0] = [] config
	//
	// [1] = [WRITE] perChainPerTokenConfig
	//
	// [2] = [WRITE, SIGNER] authority
	//
	// [3] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewSetTokenBillingInstructionBuilder creates a new `SetTokenBilling` instruction builder.
func NewSetTokenBillingInstructionBuilder() *SetTokenBilling {
	nd := &SetTokenBilling{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetChainSelector sets the "chainSelector" parameter.
func (inst *SetTokenBilling) SetChainSelector(chainSelector uint64) *SetTokenBilling {
	inst.ChainSelector = &chainSelector
	return inst
}

// SetMint sets the "mint" parameter.
func (inst *SetTokenBilling) SetMint(mint ag_solanago.PublicKey) *SetTokenBilling {
	inst.Mint = &mint
	return inst
}

// SetCfg sets the "cfg" parameter.
func (inst *SetTokenBilling) SetCfg(cfg TokenBilling) *SetTokenBilling {
	inst.Cfg = &cfg
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *SetTokenBilling) SetConfigAccount(config ag_solanago.PublicKey) *SetTokenBilling {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *SetTokenBilling) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetPerChainPerTokenConfigAccount sets the "perChainPerTokenConfig" account.
func (inst *SetTokenBilling) SetPerChainPerTokenConfigAccount(perChainPerTokenConfig ag_solanago.PublicKey) *SetTokenBilling {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(perChainPerTokenConfig).WRITE()
	return inst
}

// GetPerChainPerTokenConfigAccount gets the "perChainPerTokenConfig" account.
func (inst *SetTokenBilling) GetPerChainPerTokenConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *SetTokenBilling) SetAuthorityAccount(authority ag_solanago.PublicKey) *SetTokenBilling {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *SetTokenBilling) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *SetTokenBilling) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *SetTokenBilling {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *SetTokenBilling) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

func (inst SetTokenBilling) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SetTokenBilling,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetTokenBilling) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetTokenBilling) Validate() error {
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

func (inst *SetTokenBilling) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetTokenBilling")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("ChainSelector", *inst.ChainSelector))
						paramsBranch.Child(ag_format.Param("         Mint", *inst.Mint))
						paramsBranch.Child(ag_format.Param("          Cfg", *inst.Cfg))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                config", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("perChainPerTokenConfig", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("             authority", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("         systemProgram", inst.AccountMetaSlice[3]))
					})
				})
		})
}

func (obj SetTokenBilling) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
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
func (obj *SetTokenBilling) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
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

// NewSetTokenBillingInstruction declares a new SetTokenBilling instruction with the provided parameters and accounts.
func NewSetTokenBillingInstruction(
	// Parameters:
	chainSelector uint64,
	mint ag_solanago.PublicKey,
	cfg TokenBilling,
	// Accounts:
	config ag_solanago.PublicKey,
	perChainPerTokenConfig ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *SetTokenBilling {
	return NewSetTokenBillingInstructionBuilder().
		SetChainSelector(chainSelector).
		SetMint(mint).
		SetCfg(cfg).
		SetConfigAccount(config).
		SetPerChainPerTokenConfigAccount(perChainPerTokenConfig).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram)
}
