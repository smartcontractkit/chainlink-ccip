// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ccip_router

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Manually executes a report to the router.
//
// When a message is not being executed, then the user can trigger the execution manually.
// No verification over the transmitter, but the message needs to be in some commit report.
// It validates that the required time has passed since the commit and then executes the report.
//
// # Arguments
//
// * `ctx` - The context containing the accounts required for the execution.
// * `execution_report` - The execution report containing the message and proofs.
type ManuallyExecute struct {
	CcipVersion     *CcipVersion
	ExecutionReport *ExecutionReportSingleChain
	TokenIndexes    *[]byte

	// [0] = [] config
	//
	// [1] = [] sourceChainState
	//
	// [2] = [WRITE] commitReport
	//
	// [3] = [] externalExecutionConfig
	//
	// [4] = [WRITE, SIGNER] authority
	//
	// [5] = [] systemProgram
	//
	// [6] = [] sysvarInstructions
	//
	// [7] = [] tokenPoolsSigner
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewManuallyExecuteInstructionBuilder creates a new `ManuallyExecute` instruction builder.
func NewManuallyExecuteInstructionBuilder() *ManuallyExecute {
	nd := &ManuallyExecute{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 8),
	}
	return nd
}

// SetCcipVersion sets the "ccipVersion" parameter.
func (inst *ManuallyExecute) SetCcipVersion(ccipVersion CcipVersion) *ManuallyExecute {
	inst.CcipVersion = &ccipVersion
	return inst
}

// SetExecutionReport sets the "executionReport" parameter.
func (inst *ManuallyExecute) SetExecutionReport(executionReport ExecutionReportSingleChain) *ManuallyExecute {
	inst.ExecutionReport = &executionReport
	return inst
}

// SetTokenIndexes sets the "tokenIndexes" parameter.
func (inst *ManuallyExecute) SetTokenIndexes(tokenIndexes []byte) *ManuallyExecute {
	inst.TokenIndexes = &tokenIndexes
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *ManuallyExecute) SetConfigAccount(config ag_solanago.PublicKey) *ManuallyExecute {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *ManuallyExecute) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetSourceChainStateAccount sets the "sourceChainState" account.
func (inst *ManuallyExecute) SetSourceChainStateAccount(sourceChainState ag_solanago.PublicKey) *ManuallyExecute {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(sourceChainState)
	return inst
}

// GetSourceChainStateAccount gets the "sourceChainState" account.
func (inst *ManuallyExecute) GetSourceChainStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetCommitReportAccount sets the "commitReport" account.
func (inst *ManuallyExecute) SetCommitReportAccount(commitReport ag_solanago.PublicKey) *ManuallyExecute {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(commitReport).WRITE()
	return inst
}

// GetCommitReportAccount gets the "commitReport" account.
func (inst *ManuallyExecute) GetCommitReportAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetExternalExecutionConfigAccount sets the "externalExecutionConfig" account.
func (inst *ManuallyExecute) SetExternalExecutionConfigAccount(externalExecutionConfig ag_solanago.PublicKey) *ManuallyExecute {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(externalExecutionConfig)
	return inst
}

// GetExternalExecutionConfigAccount gets the "externalExecutionConfig" account.
func (inst *ManuallyExecute) GetExternalExecutionConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *ManuallyExecute) SetAuthorityAccount(authority ag_solanago.PublicKey) *ManuallyExecute {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *ManuallyExecute) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *ManuallyExecute) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *ManuallyExecute {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *ManuallyExecute) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetSysvarInstructionsAccount sets the "sysvarInstructions" account.
func (inst *ManuallyExecute) SetSysvarInstructionsAccount(sysvarInstructions ag_solanago.PublicKey) *ManuallyExecute {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(sysvarInstructions)
	return inst
}

// GetSysvarInstructionsAccount gets the "sysvarInstructions" account.
func (inst *ManuallyExecute) GetSysvarInstructionsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

// SetTokenPoolsSignerAccount sets the "tokenPoolsSigner" account.
func (inst *ManuallyExecute) SetTokenPoolsSignerAccount(tokenPoolsSigner ag_solanago.PublicKey) *ManuallyExecute {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(tokenPoolsSigner)
	return inst
}

// GetTokenPoolsSignerAccount gets the "tokenPoolsSigner" account.
func (inst *ManuallyExecute) GetTokenPoolsSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[7]
}

func (inst ManuallyExecute) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ManuallyExecute,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ManuallyExecute) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ManuallyExecute) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.CcipVersion == nil {
			return errors.New("CcipVersion parameter is not set")
		}
		if inst.ExecutionReport == nil {
			return errors.New("ExecutionReport parameter is not set")
		}
		if inst.TokenIndexes == nil {
			return errors.New("TokenIndexes parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.SourceChainState is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.CommitReport is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.ExternalExecutionConfig is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.SysvarInstructions is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.TokenPoolsSigner is not set")
		}
	}
	return nil
}

func (inst *ManuallyExecute) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ManuallyExecute")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("    CcipVersion", *inst.CcipVersion))
						paramsBranch.Child(ag_format.Param("ExecutionReport", *inst.ExecutionReport))
						paramsBranch.Child(ag_format.Param("   TokenIndexes", *inst.TokenIndexes))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=8]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                 config", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("       sourceChainState", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("           commitReport", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("externalExecutionConfig", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("              authority", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("          systemProgram", inst.AccountMetaSlice[5]))
						accountsBranch.Child(ag_format.Meta("     sysvarInstructions", inst.AccountMetaSlice[6]))
						accountsBranch.Child(ag_format.Meta("       tokenPoolsSigner", inst.AccountMetaSlice[7]))
					})
				})
		})
}

func (obj ManuallyExecute) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `CcipVersion` param:
	err = encoder.Encode(obj.CcipVersion)
	if err != nil {
		return err
	}
	// Serialize `ExecutionReport` param:
	err = encoder.Encode(obj.ExecutionReport)
	if err != nil {
		return err
	}
	// Serialize `TokenIndexes` param:
	err = encoder.Encode(obj.TokenIndexes)
	if err != nil {
		return err
	}
	return nil
}
func (obj *ManuallyExecute) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `CcipVersion`:
	err = decoder.Decode(&obj.CcipVersion)
	if err != nil {
		return err
	}
	// Deserialize `ExecutionReport`:
	err = decoder.Decode(&obj.ExecutionReport)
	if err != nil {
		return err
	}
	// Deserialize `TokenIndexes`:
	err = decoder.Decode(&obj.TokenIndexes)
	if err != nil {
		return err
	}
	return nil
}

// NewManuallyExecuteInstruction declares a new ManuallyExecute instruction with the provided parameters and accounts.
func NewManuallyExecuteInstruction(
	// Parameters:
	ccipVersion CcipVersion,
	executionReport ExecutionReportSingleChain,
	tokenIndexes []byte,
	// Accounts:
	config ag_solanago.PublicKey,
	sourceChainState ag_solanago.PublicKey,
	commitReport ag_solanago.PublicKey,
	externalExecutionConfig ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	sysvarInstructions ag_solanago.PublicKey,
	tokenPoolsSigner ag_solanago.PublicKey) *ManuallyExecute {
	return NewManuallyExecuteInstructionBuilder().
		SetCcipVersion(ccipVersion).
		SetExecutionReport(executionReport).
		SetTokenIndexes(tokenIndexes).
		SetConfigAccount(config).
		SetSourceChainStateAccount(sourceChainState).
		SetCommitReportAccount(commitReport).
		SetExternalExecutionConfigAccount(externalExecutionConfig).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram).
		SetSysvarInstructionsAccount(sysvarInstructions).
		SetTokenPoolsSignerAccount(tokenPoolsSigner)
}
