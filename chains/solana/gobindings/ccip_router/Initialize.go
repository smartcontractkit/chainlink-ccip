// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ccip_router

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Initialization Flow //
// Initializes the CCIP Router.
//
// The initialization of the Router is responsibility of Admin, nothing more than calling this method should be done first.
//
// # Arguments
//
// * `ctx` - The context containing the accounts required for initialization.
// * `svm_chain_selector` - The chain selector for SVM.
// * `enable_execution_after` - The minimum amount of time required between a message has been committed and can be manually executed.
type Initialize struct {
	SvmChainSelector     *uint64
	EnableExecutionAfter *int64
	FeeAggregator        *ag_solanago.PublicKey
	FeeQuoter            *ag_solanago.PublicKey
	LinkTokenMint        *ag_solanago.PublicKey
	MaxFeeJuelsPerMsg    *ag_binary.Uint128

	// [0] = [WRITE] config
	//
	// [1] = [WRITE] state
	//
	// [2] = [WRITE, SIGNER] authority
	//
	// [3] = [] system_program
	//
	// [4] = [] program
	//
	// [5] = [] program_data
	//
	// [6] = [WRITE] external_execution_config
	//
	// [7] = [WRITE] token_pools_signer
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitializeInstructionBuilder creates a new `Initialize` instruction builder.
func NewInitializeInstructionBuilder() *Initialize {
	nd := &Initialize{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 8),
	}
	return nd
}

// SetSvmChainSelector sets the "svm_chain_selector" parameter.
func (inst *Initialize) SetSvmChainSelector(svm_chain_selector uint64) *Initialize {
	inst.SvmChainSelector = &svm_chain_selector
	return inst
}

// SetEnableExecutionAfter sets the "enable_execution_after" parameter.
func (inst *Initialize) SetEnableExecutionAfter(enable_execution_after int64) *Initialize {
	inst.EnableExecutionAfter = &enable_execution_after
	return inst
}

// SetFeeAggregator sets the "fee_aggregator" parameter.
func (inst *Initialize) SetFeeAggregator(fee_aggregator ag_solanago.PublicKey) *Initialize {
	inst.FeeAggregator = &fee_aggregator
	return inst
}

// SetFeeQuoter sets the "fee_quoter" parameter.
func (inst *Initialize) SetFeeQuoter(fee_quoter ag_solanago.PublicKey) *Initialize {
	inst.FeeQuoter = &fee_quoter
	return inst
}

// SetLinkTokenMint sets the "link_token_mint" parameter.
func (inst *Initialize) SetLinkTokenMint(link_token_mint ag_solanago.PublicKey) *Initialize {
	inst.LinkTokenMint = &link_token_mint
	return inst
}

// SetMaxFeeJuelsPerMsg sets the "max_fee_juels_per_msg" parameter.
func (inst *Initialize) SetMaxFeeJuelsPerMsg(max_fee_juels_per_msg ag_binary.Uint128) *Initialize {
	inst.MaxFeeJuelsPerMsg = &max_fee_juels_per_msg
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *Initialize) SetConfigAccount(config ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config).WRITE()
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *Initialize) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetStateAccount sets the "state" account.
func (inst *Initialize) SetStateAccount(state ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *Initialize) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *Initialize) SetAuthorityAccount(authority ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *Initialize) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *Initialize) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *Initialize) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetProgramAccount sets the "program" account.
func (inst *Initialize) SetProgramAccount(program ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *Initialize) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetProgramDataAccount sets the "program_data" account.
func (inst *Initialize) SetProgramDataAccount(programData ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(programData)
	return inst
}

// GetProgramDataAccount gets the "program_data" account.
func (inst *Initialize) GetProgramDataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetExternalExecutionConfigAccount sets the "external_execution_config" account.
func (inst *Initialize) SetExternalExecutionConfigAccount(externalExecutionConfig ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(externalExecutionConfig).WRITE()
	return inst
}

// GetExternalExecutionConfigAccount gets the "external_execution_config" account.
func (inst *Initialize) GetExternalExecutionConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetTokenPoolsSignerAccount sets the "token_pools_signer" account.
func (inst *Initialize) SetTokenPoolsSignerAccount(tokenPoolsSigner ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(tokenPoolsSigner).WRITE()
	return inst
}

// GetTokenPoolsSignerAccount gets the "token_pools_signer" account.
func (inst *Initialize) GetTokenPoolsSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

func (inst Initialize) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Initialize,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Initialize) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Initialize) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.SvmChainSelector == nil {
			return errors.New("SvmChainSelector parameter is not set")
		}
		if inst.EnableExecutionAfter == nil {
			return errors.New("EnableExecutionAfter parameter is not set")
		}
		if inst.FeeAggregator == nil {
			return errors.New("FeeAggregator parameter is not set")
		}
		if inst.FeeQuoter == nil {
			return errors.New("FeeQuoter parameter is not set")
		}
		if inst.LinkTokenMint == nil {
			return errors.New("LinkTokenMint parameter is not set")
		}
		if inst.MaxFeeJuelsPerMsg == nil {
			return errors.New("MaxFeeJuelsPerMsg parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Program is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.ProgramData is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.ExternalExecutionConfig is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.TokenPoolsSigner is not set")
		}
	}
	return nil
}

func (inst *Initialize) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Initialize")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=6]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("      SvmChainSelector", *inst.SvmChainSelector))
						paramsBranch.Child(ag_format.Param("  EnableExecutionAfter", *inst.EnableExecutionAfter))
						paramsBranch.Child(ag_format.Param("         FeeAggregator", *inst.FeeAggregator))
						paramsBranch.Child(ag_format.Param("             FeeQuoter", *inst.FeeQuoter))
						paramsBranch.Child(ag_format.Param("         LinkTokenMint", *inst.LinkTokenMint))
						paramsBranch.Child(ag_format.Param("     MaxFeeJuelsPerMsg", *inst.MaxFeeJuelsPerMsg))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=8]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                   config", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                    state", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                authority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("           system_program", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                  program", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("             program_data", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("external_execution_config", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("       token_pools_signer", inst.AccountMetaSlice.Get(7)))
					})
				})
		})
}

func (obj Initialize) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `SvmChainSelector` param:
	err = encoder.Encode(obj.SvmChainSelector)
	if err != nil {
		return err
	}
	// Serialize `EnableExecutionAfter` param:
	err = encoder.Encode(obj.EnableExecutionAfter)
	if err != nil {
		return err
	}
	// Serialize `FeeAggregator` param:
	err = encoder.Encode(obj.FeeAggregator)
	if err != nil {
		return err
	}
	// Serialize `FeeQuoter` param:
	err = encoder.Encode(obj.FeeQuoter)
	if err != nil {
		return err
	}
	// Serialize `LinkTokenMint` param:
	err = encoder.Encode(obj.LinkTokenMint)
	if err != nil {
		return err
	}
	// Serialize `MaxFeeJuelsPerMsg` param:
	err = encoder.Encode(obj.MaxFeeJuelsPerMsg)
	if err != nil {
		return err
	}
	return nil
}
func (obj *Initialize) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `SvmChainSelector`:
	err = decoder.Decode(&obj.SvmChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `EnableExecutionAfter`:
	err = decoder.Decode(&obj.EnableExecutionAfter)
	if err != nil {
		return err
	}
	// Deserialize `FeeAggregator`:
	err = decoder.Decode(&obj.FeeAggregator)
	if err != nil {
		return err
	}
	// Deserialize `FeeQuoter`:
	err = decoder.Decode(&obj.FeeQuoter)
	if err != nil {
		return err
	}
	// Deserialize `LinkTokenMint`:
	err = decoder.Decode(&obj.LinkTokenMint)
	if err != nil {
		return err
	}
	// Deserialize `MaxFeeJuelsPerMsg`:
	err = decoder.Decode(&obj.MaxFeeJuelsPerMsg)
	if err != nil {
		return err
	}
	return nil
}

// NewInitializeInstruction declares a new Initialize instruction with the provided parameters and accounts.
func NewInitializeInstruction(
	// Parameters:
	svm_chain_selector uint64,
	enable_execution_after int64,
	fee_aggregator ag_solanago.PublicKey,
	fee_quoter ag_solanago.PublicKey,
	link_token_mint ag_solanago.PublicKey,
	max_fee_juels_per_msg ag_binary.Uint128,
	// Accounts:
	config ag_solanago.PublicKey,
	state ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	program ag_solanago.PublicKey,
	programData ag_solanago.PublicKey,
	externalExecutionConfig ag_solanago.PublicKey,
	tokenPoolsSigner ag_solanago.PublicKey) *Initialize {
	return NewInitializeInstructionBuilder().
		SetSvmChainSelector(svm_chain_selector).
		SetEnableExecutionAfter(enable_execution_after).
		SetFeeAggregator(fee_aggregator).
		SetFeeQuoter(fee_quoter).
		SetLinkTokenMint(link_token_mint).
		SetMaxFeeJuelsPerMsg(max_fee_juels_per_msg).
		SetConfigAccount(config).
		SetStateAccount(state).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram).
		SetProgramAccount(program).
		SetProgramDataAccount(programData).
		SetExternalExecutionConfigAccount(externalExecutionConfig).
		SetTokenPoolsSignerAccount(tokenPoolsSigner)
}
