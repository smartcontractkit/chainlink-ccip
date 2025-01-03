// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ccip_router

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Initializes the CCIP Router.
//
// The initialization of the Router is responsibility of Admin, nothing more than calling this method should be done first.
//
// # Arguments
//
// * `ctx` - The context containing the accounts required for initialization.
// * `solana_chain_selector` - The chain selector for Solana.
// * `default_gas_limit` - The default gas limit for other destination chains.
// * `default_allow_out_of_order_execution` - Whether out-of-order execution is allowed by default for other destination chains.
// * `enable_execution_after` - The minimum amount of time required between a message has been committed and can be manually executed.
type Initialize struct {
	SolanaChainSelector             *uint64
	DefaultGasLimit                 *ag_binary.Uint128
	DefaultAllowOutOfOrderExecution *bool
	EnableExecutionAfter            *int64
	FeeAggregator                   *ag_solanago.PublicKey

	// [0] = [WRITE] config
	//
	// [1] = [WRITE] state
	//
	// [2] = [WRITE, SIGNER] authority
	//
	// [3] = [] systemProgram
	//
	// [4] = [] program
	//
	// [5] = [] programData
	//
	// [6] = [WRITE] externalExecutionConfig
	//
	// [7] = [WRITE] tokenPoolsSigner
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewInitializeInstructionBuilder creates a new `Initialize` instruction builder.
func NewInitializeInstructionBuilder() *Initialize {
	nd := &Initialize{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 8),
	}
	return nd
}

// SetSolanaChainSelector sets the "solanaChainSelector" parameter.
func (inst *Initialize) SetSolanaChainSelector(solanaChainSelector uint64) *Initialize {
	inst.SolanaChainSelector = &solanaChainSelector
	return inst
}

// SetDefaultGasLimit sets the "defaultGasLimit" parameter.
func (inst *Initialize) SetDefaultGasLimit(defaultGasLimit ag_binary.Uint128) *Initialize {
	inst.DefaultGasLimit = &defaultGasLimit
	return inst
}

// SetDefaultAllowOutOfOrderExecution sets the "defaultAllowOutOfOrderExecution" parameter.
func (inst *Initialize) SetDefaultAllowOutOfOrderExecution(defaultAllowOutOfOrderExecution bool) *Initialize {
	inst.DefaultAllowOutOfOrderExecution = &defaultAllowOutOfOrderExecution
	return inst
}

// SetEnableExecutionAfter sets the "enableExecutionAfter" parameter.
func (inst *Initialize) SetEnableExecutionAfter(enableExecutionAfter int64) *Initialize {
	inst.EnableExecutionAfter = &enableExecutionAfter
	return inst
}

// SetFeeAggregator sets the "feeAggregator" parameter.
func (inst *Initialize) SetFeeAggregator(feeAggregator ag_solanago.PublicKey) *Initialize {
	inst.FeeAggregator = &feeAggregator
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *Initialize) SetConfigAccount(config ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config).WRITE()
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *Initialize) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetStateAccount sets the "state" account.
func (inst *Initialize) SetStateAccount(state ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *Initialize) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *Initialize) SetAuthorityAccount(authority ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *Initialize) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *Initialize) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *Initialize) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetProgramAccount sets the "program" account.
func (inst *Initialize) SetProgramAccount(program ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *Initialize) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetProgramDataAccount sets the "programData" account.
func (inst *Initialize) SetProgramDataAccount(programData ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(programData)
	return inst
}

// GetProgramDataAccount gets the "programData" account.
func (inst *Initialize) GetProgramDataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetExternalExecutionConfigAccount sets the "externalExecutionConfig" account.
func (inst *Initialize) SetExternalExecutionConfigAccount(externalExecutionConfig ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(externalExecutionConfig).WRITE()
	return inst
}

// GetExternalExecutionConfigAccount gets the "externalExecutionConfig" account.
func (inst *Initialize) GetExternalExecutionConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

// SetTokenPoolsSignerAccount sets the "tokenPoolsSigner" account.
func (inst *Initialize) SetTokenPoolsSignerAccount(tokenPoolsSigner ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(tokenPoolsSigner).WRITE()
	return inst
}

// GetTokenPoolsSignerAccount gets the "tokenPoolsSigner" account.
func (inst *Initialize) GetTokenPoolsSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[7]
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
		if inst.SolanaChainSelector == nil {
			return errors.New("SolanaChainSelector parameter is not set")
		}
		if inst.DefaultGasLimit == nil {
			return errors.New("DefaultGasLimit parameter is not set")
		}
		if inst.DefaultAllowOutOfOrderExecution == nil {
			return errors.New("DefaultAllowOutOfOrderExecution parameter is not set")
		}
		if inst.EnableExecutionAfter == nil {
			return errors.New("EnableExecutionAfter parameter is not set")
		}
		if inst.FeeAggregator == nil {
			return errors.New("FeeAggregator parameter is not set")
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
					instructionBranch.Child("Params[len=5]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("            SolanaChainSelector", *inst.SolanaChainSelector))
						paramsBranch.Child(ag_format.Param("                DefaultGasLimit", *inst.DefaultGasLimit))
						paramsBranch.Child(ag_format.Param("DefaultAllowOutOfOrderExecution", *inst.DefaultAllowOutOfOrderExecution))
						paramsBranch.Child(ag_format.Param("           EnableExecutionAfter", *inst.EnableExecutionAfter))
						paramsBranch.Child(ag_format.Param("                  FeeAggregator", *inst.FeeAggregator))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=8]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                 config", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("                  state", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("              authority", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("          systemProgram", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("                program", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("            programData", inst.AccountMetaSlice[5]))
						accountsBranch.Child(ag_format.Meta("externalExecutionConfig", inst.AccountMetaSlice[6]))
						accountsBranch.Child(ag_format.Meta("       tokenPoolsSigner", inst.AccountMetaSlice[7]))
					})
				})
		})
}

func (obj Initialize) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `SolanaChainSelector` param:
	err = encoder.Encode(obj.SolanaChainSelector)
	if err != nil {
		return err
	}
	// Serialize `DefaultGasLimit` param:
	err = encoder.Encode(obj.DefaultGasLimit)
	if err != nil {
		return err
	}
	// Serialize `DefaultAllowOutOfOrderExecution` param:
	err = encoder.Encode(obj.DefaultAllowOutOfOrderExecution)
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
	return nil
}
func (obj *Initialize) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `SolanaChainSelector`:
	err = decoder.Decode(&obj.SolanaChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `DefaultGasLimit`:
	err = decoder.Decode(&obj.DefaultGasLimit)
	if err != nil {
		return err
	}
	// Deserialize `DefaultAllowOutOfOrderExecution`:
	err = decoder.Decode(&obj.DefaultAllowOutOfOrderExecution)
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
	return nil
}

// NewInitializeInstruction declares a new Initialize instruction with the provided parameters and accounts.
func NewInitializeInstruction(
	// Parameters:
	solanaChainSelector uint64,
	defaultGasLimit ag_binary.Uint128,
	defaultAllowOutOfOrderExecution bool,
	enableExecutionAfter int64,
	feeAggregator ag_solanago.PublicKey,
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
		SetSolanaChainSelector(solanaChainSelector).
		SetDefaultGasLimit(defaultGasLimit).
		SetDefaultAllowOutOfOrderExecution(defaultAllowOutOfOrderExecution).
		SetEnableExecutionAfter(enableExecutionAfter).
		SetFeeAggregator(feeAggregator).
		SetConfigAccount(config).
		SetStateAccount(state).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram).
		SetProgramAccount(program).
		SetProgramDataAccount(programData).
		SetExternalExecutionConfigAccount(externalExecutionConfig).
		SetTokenPoolsSignerAccount(tokenPoolsSigner)
}
