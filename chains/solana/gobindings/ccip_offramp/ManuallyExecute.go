// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ccip_offramp

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
// * `raw_execution_report` - The serialized execution report containing the message and proofs.
type ManuallyExecute struct {
	RawExecutionReport *[]byte
	TokenIndexes       *[]byte

	// [0] = [] config
	//
	// [1] = [] reference_addresses
	//
	// [2] = [] source_chain
	//
	// [3] = [WRITE] commit_report
	//
	// [4] = [] offramp
	//
	// [5] = [] allowed_offramp
	// ··········· CHECK PDA of the router program verifying the signer is an allowed offramp.
	// ··········· If PDA does not exist, the router doesn't allow this offramp. This is just used
	// ··········· so that token pools and receivers can then check that the caller is an actual offramp that
	// ··········· has been registered in the router as such for that source chain.
	//
	// [6] = [] external_execution_config
	//
	// [7] = [WRITE, SIGNER] authority
	//
	// [8] = [] system_program
	//
	// [9] = [] sysvar_instructions
	//
	// [10] = [] token_pools_signer
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewManuallyExecuteInstructionBuilder creates a new `ManuallyExecute` instruction builder.
func NewManuallyExecuteInstructionBuilder() *ManuallyExecute {
	nd := &ManuallyExecute{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 11),
	}
	return nd
}

// SetRawExecutionReport sets the "raw_execution_report" parameter.
func (inst *ManuallyExecute) SetRawExecutionReport(raw_execution_report []byte) *ManuallyExecute {
	inst.RawExecutionReport = &raw_execution_report
	return inst
}

// SetTokenIndexes sets the "token_indexes" parameter.
func (inst *ManuallyExecute) SetTokenIndexes(token_indexes []byte) *ManuallyExecute {
	inst.TokenIndexes = &token_indexes
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *ManuallyExecute) SetConfigAccount(config ag_solanago.PublicKey) *ManuallyExecute {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *ManuallyExecute) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetReferenceAddressesAccount sets the "reference_addresses" account.
func (inst *ManuallyExecute) SetReferenceAddressesAccount(referenceAddresses ag_solanago.PublicKey) *ManuallyExecute {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(referenceAddresses)
	return inst
}

// GetReferenceAddressesAccount gets the "reference_addresses" account.
func (inst *ManuallyExecute) GetReferenceAddressesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSourceChainAccount sets the "source_chain" account.
func (inst *ManuallyExecute) SetSourceChainAccount(sourceChain ag_solanago.PublicKey) *ManuallyExecute {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(sourceChain)
	return inst
}

// GetSourceChainAccount gets the "source_chain" account.
func (inst *ManuallyExecute) GetSourceChainAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetCommitReportAccount sets the "commit_report" account.
func (inst *ManuallyExecute) SetCommitReportAccount(commitReport ag_solanago.PublicKey) *ManuallyExecute {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(commitReport).WRITE()
	return inst
}

// GetCommitReportAccount gets the "commit_report" account.
func (inst *ManuallyExecute) GetCommitReportAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetOfframpAccount sets the "offramp" account.
func (inst *ManuallyExecute) SetOfframpAccount(offramp ag_solanago.PublicKey) *ManuallyExecute {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(offramp)
	return inst
}

// GetOfframpAccount gets the "offramp" account.
func (inst *ManuallyExecute) GetOfframpAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetAllowedOfframpAccount sets the "allowed_offramp" account.
// CHECK PDA of the router program verifying the signer is an allowed offramp.
// If PDA does not exist, the router doesn't allow this offramp. This is just used
// so that token pools and receivers can then check that the caller is an actual offramp that
// has been registered in the router as such for that source chain.
func (inst *ManuallyExecute) SetAllowedOfframpAccount(allowedOfframp ag_solanago.PublicKey) *ManuallyExecute {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(allowedOfframp)
	return inst
}

// GetAllowedOfframpAccount gets the "allowed_offramp" account.
// CHECK PDA of the router program verifying the signer is an allowed offramp.
// If PDA does not exist, the router doesn't allow this offramp. This is just used
// so that token pools and receivers can then check that the caller is an actual offramp that
// has been registered in the router as such for that source chain.
func (inst *ManuallyExecute) GetAllowedOfframpAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetExternalExecutionConfigAccount sets the "external_execution_config" account.
func (inst *ManuallyExecute) SetExternalExecutionConfigAccount(externalExecutionConfig ag_solanago.PublicKey) *ManuallyExecute {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(externalExecutionConfig)
	return inst
}

// GetExternalExecutionConfigAccount gets the "external_execution_config" account.
func (inst *ManuallyExecute) GetExternalExecutionConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *ManuallyExecute) SetAuthorityAccount(authority ag_solanago.PublicKey) *ManuallyExecute {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *ManuallyExecute) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *ManuallyExecute) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *ManuallyExecute {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *ManuallyExecute) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetSysvarInstructionsAccount sets the "sysvar_instructions" account.
func (inst *ManuallyExecute) SetSysvarInstructionsAccount(sysvarInstructions ag_solanago.PublicKey) *ManuallyExecute {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(sysvarInstructions)
	return inst
}

// GetSysvarInstructionsAccount gets the "sysvar_instructions" account.
func (inst *ManuallyExecute) GetSysvarInstructionsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetTokenPoolsSignerAccount sets the "token_pools_signer" account.
func (inst *ManuallyExecute) SetTokenPoolsSignerAccount(tokenPoolsSigner ag_solanago.PublicKey) *ManuallyExecute {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(tokenPoolsSigner)
	return inst
}

// GetTokenPoolsSignerAccount gets the "token_pools_signer" account.
func (inst *ManuallyExecute) GetTokenPoolsSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
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
		if inst.RawExecutionReport == nil {
			return errors.New("RawExecutionReport parameter is not set")
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
			return errors.New("accounts.ReferenceAddresses is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SourceChain is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.CommitReport is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Offramp is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.AllowedOfframp is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.ExternalExecutionConfig is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.SysvarInstructions is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
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
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("  RawExecutionReport", *inst.RawExecutionReport))
						paramsBranch.Child(ag_format.Param("        TokenIndexes", *inst.TokenIndexes))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=11]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                   config", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("      reference_addresses", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("             source_chain", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("            commit_report", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                  offramp", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("          allowed_offramp", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("external_execution_config", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("                authority", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("           system_program", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("      sysvar_instructions", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("       token_pools_signer", inst.AccountMetaSlice.Get(10)))
					})
				})
		})
}

func (obj ManuallyExecute) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `RawExecutionReport` param:
	err = encoder.Encode(obj.RawExecutionReport)
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
	// Deserialize `RawExecutionReport`:
	err = decoder.Decode(&obj.RawExecutionReport)
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
	raw_execution_report []byte,
	token_indexes []byte,
	// Accounts:
	config ag_solanago.PublicKey,
	referenceAddresses ag_solanago.PublicKey,
	sourceChain ag_solanago.PublicKey,
	commitReport ag_solanago.PublicKey,
	offramp ag_solanago.PublicKey,
	allowedOfframp ag_solanago.PublicKey,
	externalExecutionConfig ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	sysvarInstructions ag_solanago.PublicKey,
	tokenPoolsSigner ag_solanago.PublicKey) *ManuallyExecute {
	return NewManuallyExecuteInstructionBuilder().
		SetRawExecutionReport(raw_execution_report).
		SetTokenIndexes(token_indexes).
		SetConfigAccount(config).
		SetReferenceAddressesAccount(referenceAddresses).
		SetSourceChainAccount(sourceChain).
		SetCommitReportAccount(commitReport).
		SetOfframpAccount(offramp).
		SetAllowedOfframpAccount(allowedOfframp).
		SetExternalExecutionConfigAccount(externalExecutionConfig).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram).
		SetSysvarInstructionsAccount(sysvarInstructions).
		SetTokenPoolsSignerAccount(tokenPoolsSigner)
}
