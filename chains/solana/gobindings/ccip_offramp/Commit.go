// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ccip_offramp

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Off Ramp Flow //
// Commits a report to the router, containing a Merkle Root.
//
// The method name needs to be commit with Anchor encoding.
//
// This function is called by the OffChain when committing one Report to the SVM Router.
// In this Flow only one report is sent, the Commit Report. This is different as EVM does,
// this is because here all the chain state is stored in one account per Merkle Tree Root.
// So, to avoid having to send a dynamic size array of accounts, in this message only one Commit Report Account is sent.
// This message validates the signatures of the report and stores the Merkle Root in the Commit Report Account.
// The Report must contain an interval of messages, and the min of them must be the next sequence number expected.
// The max size of the interval is 64.
// This message emits two events: CommitReportAccepted and Transmitted.
//
// # Arguments
//
// * `ctx` - The context containing the accounts required for the commit.
// * `report_context_byte_words` - consists of:
// * report_context_byte_words[0]: ConfigDigest
// * report_context_byte_words[1]: 24 byte padding, 8 byte sequence number
// * `raw_report` - The serialized commit input report, single merkle root with RMN signatures and price updates
// * `rs` - slice of R components of signatures
// * `ss` - slice of S components of signatures
// * `raw_vs` - array of V components of signatures
type Commit struct {
	ReportContextByteWords *[2][32]uint8
	RawReport              *[]byte
	Rs                     *[][32]uint8
	Ss                     *[][32]uint8
	RawVs                  *[32]uint8

	// [0] = [] config
	//
	// [1] = [] referenceAddresses
	//
	// [2] = [WRITE] sourceChain
	//
	// [3] = [WRITE] commitReport
	//
	// [4] = [WRITE, SIGNER] authority
	//
	// [5] = [] systemProgram
	//
	// [6] = [] sysvarInstructions
	//
	// [7] = [] feeBillingSigner
	//
	// [8] = [] feeQuoter
	//
	// [9] = [] feeQuoterAllowedPriceUpdater
	// ··········· so that it can authorize the call made by this offramp
	//
	// [10] = [] feeQuoterConfig
	//
	// [11] = [] rmnRemote
	//
	// [12] = [] rmnRemoteConfigAndCurses
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewCommitInstructionBuilder creates a new `Commit` instruction builder.
func NewCommitInstructionBuilder() *Commit {
	nd := &Commit{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 13),
	}
	return nd
}

// SetReportContextByteWords sets the "reportContextByteWords" parameter.
func (inst *Commit) SetReportContextByteWords(reportContextByteWords [2][32]uint8) *Commit {
	inst.ReportContextByteWords = &reportContextByteWords
	return inst
}

// SetRawReport sets the "rawReport" parameter.
func (inst *Commit) SetRawReport(rawReport []byte) *Commit {
	inst.RawReport = &rawReport
	return inst
}

// SetRs sets the "rs" parameter.
func (inst *Commit) SetRs(rs [][32]uint8) *Commit {
	inst.Rs = &rs
	return inst
}

// SetSs sets the "ss" parameter.
func (inst *Commit) SetSs(ss [][32]uint8) *Commit {
	inst.Ss = &ss
	return inst
}

// SetRawVs sets the "rawVs" parameter.
func (inst *Commit) SetRawVs(rawVs [32]uint8) *Commit {
	inst.RawVs = &rawVs
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *Commit) SetConfigAccount(config ag_solanago.PublicKey) *Commit {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *Commit) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetReferenceAddressesAccount sets the "referenceAddresses" account.
func (inst *Commit) SetReferenceAddressesAccount(referenceAddresses ag_solanago.PublicKey) *Commit {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(referenceAddresses)
	return inst
}

// GetReferenceAddressesAccount gets the "referenceAddresses" account.
func (inst *Commit) GetReferenceAddressesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetSourceChainAccount sets the "sourceChain" account.
func (inst *Commit) SetSourceChainAccount(sourceChain ag_solanago.PublicKey) *Commit {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(sourceChain).WRITE()
	return inst
}

// GetSourceChainAccount gets the "sourceChain" account.
func (inst *Commit) GetSourceChainAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetCommitReportAccount sets the "commitReport" account.
func (inst *Commit) SetCommitReportAccount(commitReport ag_solanago.PublicKey) *Commit {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(commitReport).WRITE()
	return inst
}

// GetCommitReportAccount gets the "commitReport" account.
func (inst *Commit) GetCommitReportAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *Commit) SetAuthorityAccount(authority ag_solanago.PublicKey) *Commit {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *Commit) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *Commit) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Commit {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *Commit) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetSysvarInstructionsAccount sets the "sysvarInstructions" account.
func (inst *Commit) SetSysvarInstructionsAccount(sysvarInstructions ag_solanago.PublicKey) *Commit {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(sysvarInstructions)
	return inst
}

// GetSysvarInstructionsAccount gets the "sysvarInstructions" account.
func (inst *Commit) GetSysvarInstructionsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

// SetFeeBillingSignerAccount sets the "feeBillingSigner" account.
func (inst *Commit) SetFeeBillingSignerAccount(feeBillingSigner ag_solanago.PublicKey) *Commit {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(feeBillingSigner)
	return inst
}

// GetFeeBillingSignerAccount gets the "feeBillingSigner" account.
func (inst *Commit) GetFeeBillingSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[7]
}

// SetFeeQuoterAccount sets the "feeQuoter" account.
func (inst *Commit) SetFeeQuoterAccount(feeQuoter ag_solanago.PublicKey) *Commit {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(feeQuoter)
	return inst
}

// GetFeeQuoterAccount gets the "feeQuoter" account.
func (inst *Commit) GetFeeQuoterAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[8]
}

// SetFeeQuoterAllowedPriceUpdaterAccount sets the "feeQuoterAllowedPriceUpdater" account.
// so that it can authorize the call made by this offramp
func (inst *Commit) SetFeeQuoterAllowedPriceUpdaterAccount(feeQuoterAllowedPriceUpdater ag_solanago.PublicKey) *Commit {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(feeQuoterAllowedPriceUpdater)
	return inst
}

// GetFeeQuoterAllowedPriceUpdaterAccount gets the "feeQuoterAllowedPriceUpdater" account.
// so that it can authorize the call made by this offramp
func (inst *Commit) GetFeeQuoterAllowedPriceUpdaterAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[9]
}

// SetFeeQuoterConfigAccount sets the "feeQuoterConfig" account.
func (inst *Commit) SetFeeQuoterConfigAccount(feeQuoterConfig ag_solanago.PublicKey) *Commit {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(feeQuoterConfig)
	return inst
}

// GetFeeQuoterConfigAccount gets the "feeQuoterConfig" account.
func (inst *Commit) GetFeeQuoterConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[10]
}

// SetRmnRemoteAccount sets the "rmnRemote" account.
func (inst *Commit) SetRmnRemoteAccount(rmnRemote ag_solanago.PublicKey) *Commit {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(rmnRemote)
	return inst
}

// GetRmnRemoteAccount gets the "rmnRemote" account.
func (inst *Commit) GetRmnRemoteAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[11]
}

// SetRmnRemoteConfigAndCursesAccount sets the "rmnRemoteConfigAndCurses" account.
func (inst *Commit) SetRmnRemoteConfigAndCursesAccount(rmnRemoteConfigAndCurses ag_solanago.PublicKey) *Commit {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(rmnRemoteConfigAndCurses)
	return inst
}

// GetRmnRemoteConfigAndCursesAccount gets the "rmnRemoteConfigAndCurses" account.
func (inst *Commit) GetRmnRemoteConfigAndCursesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[12]
}

func (inst Commit) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Commit,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Commit) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Commit) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.ReportContextByteWords == nil {
			return errors.New("ReportContextByteWords parameter is not set")
		}
		if inst.RawReport == nil {
			return errors.New("RawReport parameter is not set")
		}
		if inst.Rs == nil {
			return errors.New("Rs parameter is not set")
		}
		if inst.Ss == nil {
			return errors.New("Ss parameter is not set")
		}
		if inst.RawVs == nil {
			return errors.New("RawVs parameter is not set")
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
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.SysvarInstructions is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.FeeBillingSigner is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.FeeQuoter is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.FeeQuoterAllowedPriceUpdater is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.FeeQuoterConfig is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.RmnRemote is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.RmnRemoteConfigAndCurses is not set")
		}
	}
	return nil
}

func (inst *Commit) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Commit")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=5]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("ReportContextByteWords", *inst.ReportContextByteWords))
						paramsBranch.Child(ag_format.Param("             RawReport", *inst.RawReport))
						paramsBranch.Child(ag_format.Param("                    Rs", *inst.Rs))
						paramsBranch.Child(ag_format.Param("                    Ss", *inst.Ss))
						paramsBranch.Child(ag_format.Param("                 RawVs", *inst.RawVs))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=13]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                      config", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("          referenceAddresses", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("                 sourceChain", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("                commitReport", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("                   authority", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("               systemProgram", inst.AccountMetaSlice[5]))
						accountsBranch.Child(ag_format.Meta("          sysvarInstructions", inst.AccountMetaSlice[6]))
						accountsBranch.Child(ag_format.Meta("            feeBillingSigner", inst.AccountMetaSlice[7]))
						accountsBranch.Child(ag_format.Meta("                   feeQuoter", inst.AccountMetaSlice[8]))
						accountsBranch.Child(ag_format.Meta("feeQuoterAllowedPriceUpdater", inst.AccountMetaSlice[9]))
						accountsBranch.Child(ag_format.Meta("             feeQuoterConfig", inst.AccountMetaSlice[10]))
						accountsBranch.Child(ag_format.Meta("                   rmnRemote", inst.AccountMetaSlice[11]))
						accountsBranch.Child(ag_format.Meta("    rmnRemoteConfigAndCurses", inst.AccountMetaSlice[12]))
					})
				})
		})
}

func (obj Commit) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ReportContextByteWords` param:
	err = encoder.Encode(obj.ReportContextByteWords)
	if err != nil {
		return err
	}
	// Serialize `RawReport` param:
	err = encoder.Encode(obj.RawReport)
	if err != nil {
		return err
	}
	// Serialize `Rs` param:
	err = encoder.Encode(obj.Rs)
	if err != nil {
		return err
	}
	// Serialize `Ss` param:
	err = encoder.Encode(obj.Ss)
	if err != nil {
		return err
	}
	// Serialize `RawVs` param:
	err = encoder.Encode(obj.RawVs)
	if err != nil {
		return err
	}
	return nil
}
func (obj *Commit) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ReportContextByteWords`:
	err = decoder.Decode(&obj.ReportContextByteWords)
	if err != nil {
		return err
	}
	// Deserialize `RawReport`:
	err = decoder.Decode(&obj.RawReport)
	if err != nil {
		return err
	}
	// Deserialize `Rs`:
	err = decoder.Decode(&obj.Rs)
	if err != nil {
		return err
	}
	// Deserialize `Ss`:
	err = decoder.Decode(&obj.Ss)
	if err != nil {
		return err
	}
	// Deserialize `RawVs`:
	err = decoder.Decode(&obj.RawVs)
	if err != nil {
		return err
	}
	return nil
}

// NewCommitInstruction declares a new Commit instruction with the provided parameters and accounts.
func NewCommitInstruction(
	// Parameters:
	reportContextByteWords [2][32]uint8,
	rawReport []byte,
	rs [][32]uint8,
	ss [][32]uint8,
	rawVs [32]uint8,
	// Accounts:
	config ag_solanago.PublicKey,
	referenceAddresses ag_solanago.PublicKey,
	sourceChain ag_solanago.PublicKey,
	commitReport ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	sysvarInstructions ag_solanago.PublicKey,
	feeBillingSigner ag_solanago.PublicKey,
	feeQuoter ag_solanago.PublicKey,
	feeQuoterAllowedPriceUpdater ag_solanago.PublicKey,
	feeQuoterConfig ag_solanago.PublicKey,
	rmnRemote ag_solanago.PublicKey,
	rmnRemoteConfigAndCurses ag_solanago.PublicKey) *Commit {
	return NewCommitInstructionBuilder().
		SetReportContextByteWords(reportContextByteWords).
		SetRawReport(rawReport).
		SetRs(rs).
		SetSs(ss).
		SetRawVs(rawVs).
		SetConfigAccount(config).
		SetReferenceAddressesAccount(referenceAddresses).
		SetSourceChainAccount(sourceChain).
		SetCommitReportAccount(commitReport).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram).
		SetSysvarInstructionsAccount(sysvarInstructions).
		SetFeeBillingSignerAccount(feeBillingSigner).
		SetFeeQuoterAccount(feeQuoter).
		SetFeeQuoterAllowedPriceUpdaterAccount(feeQuoterAllowedPriceUpdater).
		SetFeeQuoterConfigAccount(feeQuoterConfig).
		SetRmnRemoteAccount(rmnRemote).
		SetRmnRemoteConfigAndCursesAccount(rmnRemoteConfigAndCurses)
}
