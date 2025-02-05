// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ccip_router

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Executes a message on the destination chain.
//
// The method name needs to be execute with Anchor encoding.
//
// This function is called by the OffChain when executing one Report to the SVM Router.
// In this Flow only one message is sent, the Execution Report. This is different as EVM does,
// this is because there is no try/catch mechanism to allow batch execution.
// This message validates that the Merkle Tree Proof of the given message is correct and is stored in the Commit Report Account.
// The message must be untouched to be executed.
// This message emits the event ExecutionStateChanged with the new state of the message.
// Finally, executes the CPI instruction to the receiver program in the ccip_receive message.
//
// # Arguments
//
// * `ctx` - The context containing the accounts required for the execute.
// * `raw_execution_report` - the serialized execution report containing only one message and proofs
// * `report_context_byte_words` - report_context after execution_report to match context for manually execute (proper decoding order)
// *  consists of:
// * report_context_byte_words[0]: ConfigDigest
// * report_context_byte_words[1]: 24 byte padding, 8 byte sequence number
type Execute struct {
	RawExecutionReport     *[]byte
	ReportContextByteWords *[2][32]uint8
	TokenIndexes           *[]byte

	// [0] = [] config
	//
	// [1] = [] source_chain_state
	//
	// [2] = [WRITE] commit_report
	//
	// [3] = [] external_execution_config
	//
	// [4] = [WRITE, SIGNER] authority
	//
	// [5] = [] system_program
	//
	// [6] = [] sysvar_instructions
	//
	// [7] = [] token_pools_signer
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewExecuteInstructionBuilder creates a new `Execute` instruction builder.
func NewExecuteInstructionBuilder() *Execute {
	nd := &Execute{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 8),
	}
	return nd
}

// SetRawExecutionReport sets the "raw_execution_report" parameter.
func (inst *Execute) SetRawExecutionReport(raw_execution_report []byte) *Execute {
	inst.RawExecutionReport = &raw_execution_report
	return inst
}

// SetReportContextByteWords sets the "report_context_byte_words" parameter.
func (inst *Execute) SetReportContextByteWords(report_context_byte_words [2][32]uint8) *Execute {
	inst.ReportContextByteWords = &report_context_byte_words
	return inst
}

// SetTokenIndexes sets the "token_indexes" parameter.
func (inst *Execute) SetTokenIndexes(token_indexes []byte) *Execute {
	inst.TokenIndexes = &token_indexes
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *Execute) SetConfigAccount(config ag_solanago.PublicKey) *Execute {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *Execute) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetSourceChainStateAccount sets the "source_chain_state" account.
func (inst *Execute) SetSourceChainStateAccount(sourceChainState ag_solanago.PublicKey) *Execute {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(sourceChainState)
	return inst
}

// GetSourceChainStateAccount gets the "source_chain_state" account.
func (inst *Execute) GetSourceChainStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetCommitReportAccount sets the "commit_report" account.
func (inst *Execute) SetCommitReportAccount(commitReport ag_solanago.PublicKey) *Execute {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(commitReport).WRITE()
	return inst
}

// GetCommitReportAccount gets the "commit_report" account.
func (inst *Execute) GetCommitReportAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetExternalExecutionConfigAccount sets the "external_execution_config" account.
func (inst *Execute) SetExternalExecutionConfigAccount(externalExecutionConfig ag_solanago.PublicKey) *Execute {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(externalExecutionConfig)
	return inst
}

// GetExternalExecutionConfigAccount gets the "external_execution_config" account.
func (inst *Execute) GetExternalExecutionConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *Execute) SetAuthorityAccount(authority ag_solanago.PublicKey) *Execute {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *Execute) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *Execute) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Execute {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *Execute) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetSysvarInstructionsAccount sets the "sysvar_instructions" account.
func (inst *Execute) SetSysvarInstructionsAccount(sysvarInstructions ag_solanago.PublicKey) *Execute {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(sysvarInstructions)
	return inst
}

// GetSysvarInstructionsAccount gets the "sysvar_instructions" account.
func (inst *Execute) GetSysvarInstructionsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetTokenPoolsSignerAccount sets the "token_pools_signer" account.
func (inst *Execute) SetTokenPoolsSignerAccount(tokenPoolsSigner ag_solanago.PublicKey) *Execute {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(tokenPoolsSigner)
	return inst
}

// GetTokenPoolsSignerAccount gets the "token_pools_signer" account.
func (inst *Execute) GetTokenPoolsSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

func (inst Execute) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Execute,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Execute) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Execute) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.RawExecutionReport == nil {
			return errors.New("RawExecutionReport parameter is not set")
		}
		if inst.ReportContextByteWords == nil {
			return errors.New("ReportContextByteWords parameter is not set")
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

func (inst *Execute) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Execute")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("       RawExecutionReport", *inst.RawExecutionReport))
						paramsBranch.Child(ag_format.Param("   ReportContextByteWords", *inst.ReportContextByteWords))
						paramsBranch.Child(ag_format.Param("             TokenIndexes", *inst.TokenIndexes))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=8]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                   config", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("       source_chain_state", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("            commit_report", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("external_execution_config", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                authority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("           system_program", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("      sysvar_instructions", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("       token_pools_signer", inst.AccountMetaSlice.Get(7)))
					})
				})
		})
}

func (obj Execute) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `RawExecutionReport` param:
	err = encoder.Encode(obj.RawExecutionReport)
	if err != nil {
		return err
	}
	// Serialize `ReportContextByteWords` param:
	err = encoder.Encode(obj.ReportContextByteWords)
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
func (obj *Execute) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `RawExecutionReport`:
	err = decoder.Decode(&obj.RawExecutionReport)
	if err != nil {
		return err
	}
	// Deserialize `ReportContextByteWords`:
	err = decoder.Decode(&obj.ReportContextByteWords)
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

// NewExecuteInstruction declares a new Execute instruction with the provided parameters and accounts.
func NewExecuteInstruction(
	// Parameters:
	raw_execution_report []byte,
	report_context_byte_words [2][32]uint8,
	token_indexes []byte,
	// Accounts:
	config ag_solanago.PublicKey,
	sourceChainState ag_solanago.PublicKey,
	commitReport ag_solanago.PublicKey,
	externalExecutionConfig ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	sysvarInstructions ag_solanago.PublicKey,
	tokenPoolsSigner ag_solanago.PublicKey) *Execute {
	return NewExecuteInstructionBuilder().
		SetRawExecutionReport(raw_execution_report).
		SetReportContextByteWords(report_context_byte_words).
		SetTokenIndexes(token_indexes).
		SetConfigAccount(config).
		SetSourceChainStateAccount(sourceChainState).
		SetCommitReportAccount(commitReport).
		SetExternalExecutionConfigAccount(externalExecutionConfig).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram).
		SetSysvarInstructionsAccount(sysvarInstructions).
		SetTokenPoolsSignerAccount(tokenPoolsSigner)
}
