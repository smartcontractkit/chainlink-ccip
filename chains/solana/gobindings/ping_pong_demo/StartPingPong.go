// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ping_pong_demo

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// StartPingPong is the `startPingPong` instruction.
type StartPingPong struct {

	// [0] = [WRITE] config
	//
	// [1] = [WRITE, SIGNER] authority
	//
	// [2] = [WRITE] ccipSendSigner
	// ··········· CHECK
	//
	// [3] = [] feeTokenProgram
	//
	// [4] = [] feeTokenMint
	//
	// [5] = [WRITE] feeTokenAta
	//
	// [6] = [] ccipRouterProgram
	// ··········· CHECK
	//
	// [7] = [] ccipRouterConfig
	// ··········· CHECK
	//
	// [8] = [WRITE] ccipRouterDestChainState
	// ··········· CHECK
	//
	// [9] = [WRITE] ccipRouterNonce
	// ··········· CHECK
	//
	// [10] = [WRITE] ccipRouterFeeReceiver
	// ··········· CHECK
	//
	// [11] = [] ccipRouterFeeBillingSigner
	// ··········· CHECK
	//
	// [12] = [] feeQuoter
	// ··········· CHECK
	//
	// [13] = [] feeQuoterConfig
	// ··········· CHECK
	//
	// [14] = [] feeQuoterDestChain
	// ··········· CHECK
	//
	// [15] = [] feeQuoterBillingTokenConfig
	// ··········· CHECK
	//
	// [16] = [] feeQuoterLinkTokenConfig
	// ··········· CHECK
	//
	// [17] = [] rmnRemote
	// ··········· CHECK
	//
	// [18] = [] rmnRemoteCurses
	// ··········· CHECK
	//
	// [19] = [] rmnRemoteConfig
	// ··········· CHECK
	//
	// [20] = [WRITE] tokenPoolsSigner
	// ··········· CHECK
	//
	// [21] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewStartPingPongInstructionBuilder creates a new `StartPingPong` instruction builder.
func NewStartPingPongInstructionBuilder() *StartPingPong {
	nd := &StartPingPong{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 22),
	}
	return nd
}

// SetConfigAccount sets the "config" account.
func (inst *StartPingPong) SetConfigAccount(config ag_solanago.PublicKey) *StartPingPong {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config).WRITE()
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *StartPingPong) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *StartPingPong) SetAuthorityAccount(authority ag_solanago.PublicKey) *StartPingPong {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *StartPingPong) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetCcipSendSignerAccount sets the "ccipSendSigner" account.
// CHECK
func (inst *StartPingPong) SetCcipSendSignerAccount(ccipSendSigner ag_solanago.PublicKey) *StartPingPong {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(ccipSendSigner).WRITE()
	return inst
}

// GetCcipSendSignerAccount gets the "ccipSendSigner" account.
// CHECK
func (inst *StartPingPong) GetCcipSendSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetFeeTokenProgramAccount sets the "feeTokenProgram" account.
func (inst *StartPingPong) SetFeeTokenProgramAccount(feeTokenProgram ag_solanago.PublicKey) *StartPingPong {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(feeTokenProgram)
	return inst
}

// GetFeeTokenProgramAccount gets the "feeTokenProgram" account.
func (inst *StartPingPong) GetFeeTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetFeeTokenMintAccount sets the "feeTokenMint" account.
func (inst *StartPingPong) SetFeeTokenMintAccount(feeTokenMint ag_solanago.PublicKey) *StartPingPong {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(feeTokenMint)
	return inst
}

// GetFeeTokenMintAccount gets the "feeTokenMint" account.
func (inst *StartPingPong) GetFeeTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetFeeTokenAtaAccount sets the "feeTokenAta" account.
func (inst *StartPingPong) SetFeeTokenAtaAccount(feeTokenAta ag_solanago.PublicKey) *StartPingPong {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(feeTokenAta).WRITE()
	return inst
}

// GetFeeTokenAtaAccount gets the "feeTokenAta" account.
func (inst *StartPingPong) GetFeeTokenAtaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetCcipRouterProgramAccount sets the "ccipRouterProgram" account.
// CHECK
func (inst *StartPingPong) SetCcipRouterProgramAccount(ccipRouterProgram ag_solanago.PublicKey) *StartPingPong {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(ccipRouterProgram)
	return inst
}

// GetCcipRouterProgramAccount gets the "ccipRouterProgram" account.
// CHECK
func (inst *StartPingPong) GetCcipRouterProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

// SetCcipRouterConfigAccount sets the "ccipRouterConfig" account.
// CHECK
func (inst *StartPingPong) SetCcipRouterConfigAccount(ccipRouterConfig ag_solanago.PublicKey) *StartPingPong {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(ccipRouterConfig)
	return inst
}

// GetCcipRouterConfigAccount gets the "ccipRouterConfig" account.
// CHECK
func (inst *StartPingPong) GetCcipRouterConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[7]
}

// SetCcipRouterDestChainStateAccount sets the "ccipRouterDestChainState" account.
// CHECK
func (inst *StartPingPong) SetCcipRouterDestChainStateAccount(ccipRouterDestChainState ag_solanago.PublicKey) *StartPingPong {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(ccipRouterDestChainState).WRITE()
	return inst
}

// GetCcipRouterDestChainStateAccount gets the "ccipRouterDestChainState" account.
// CHECK
func (inst *StartPingPong) GetCcipRouterDestChainStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[8]
}

// SetCcipRouterNonceAccount sets the "ccipRouterNonce" account.
// CHECK
func (inst *StartPingPong) SetCcipRouterNonceAccount(ccipRouterNonce ag_solanago.PublicKey) *StartPingPong {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(ccipRouterNonce).WRITE()
	return inst
}

// GetCcipRouterNonceAccount gets the "ccipRouterNonce" account.
// CHECK
func (inst *StartPingPong) GetCcipRouterNonceAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[9]
}

// SetCcipRouterFeeReceiverAccount sets the "ccipRouterFeeReceiver" account.
// CHECK
func (inst *StartPingPong) SetCcipRouterFeeReceiverAccount(ccipRouterFeeReceiver ag_solanago.PublicKey) *StartPingPong {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(ccipRouterFeeReceiver).WRITE()
	return inst
}

// GetCcipRouterFeeReceiverAccount gets the "ccipRouterFeeReceiver" account.
// CHECK
func (inst *StartPingPong) GetCcipRouterFeeReceiverAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[10]
}

// SetCcipRouterFeeBillingSignerAccount sets the "ccipRouterFeeBillingSigner" account.
// CHECK
func (inst *StartPingPong) SetCcipRouterFeeBillingSignerAccount(ccipRouterFeeBillingSigner ag_solanago.PublicKey) *StartPingPong {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(ccipRouterFeeBillingSigner)
	return inst
}

// GetCcipRouterFeeBillingSignerAccount gets the "ccipRouterFeeBillingSigner" account.
// CHECK
func (inst *StartPingPong) GetCcipRouterFeeBillingSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[11]
}

// SetFeeQuoterAccount sets the "feeQuoter" account.
// CHECK
func (inst *StartPingPong) SetFeeQuoterAccount(feeQuoter ag_solanago.PublicKey) *StartPingPong {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(feeQuoter)
	return inst
}

// GetFeeQuoterAccount gets the "feeQuoter" account.
// CHECK
func (inst *StartPingPong) GetFeeQuoterAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[12]
}

// SetFeeQuoterConfigAccount sets the "feeQuoterConfig" account.
// CHECK
func (inst *StartPingPong) SetFeeQuoterConfigAccount(feeQuoterConfig ag_solanago.PublicKey) *StartPingPong {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(feeQuoterConfig)
	return inst
}

// GetFeeQuoterConfigAccount gets the "feeQuoterConfig" account.
// CHECK
func (inst *StartPingPong) GetFeeQuoterConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[13]
}

// SetFeeQuoterDestChainAccount sets the "feeQuoterDestChain" account.
// CHECK
func (inst *StartPingPong) SetFeeQuoterDestChainAccount(feeQuoterDestChain ag_solanago.PublicKey) *StartPingPong {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(feeQuoterDestChain)
	return inst
}

// GetFeeQuoterDestChainAccount gets the "feeQuoterDestChain" account.
// CHECK
func (inst *StartPingPong) GetFeeQuoterDestChainAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[14]
}

// SetFeeQuoterBillingTokenConfigAccount sets the "feeQuoterBillingTokenConfig" account.
// CHECK
func (inst *StartPingPong) SetFeeQuoterBillingTokenConfigAccount(feeQuoterBillingTokenConfig ag_solanago.PublicKey) *StartPingPong {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(feeQuoterBillingTokenConfig)
	return inst
}

// GetFeeQuoterBillingTokenConfigAccount gets the "feeQuoterBillingTokenConfig" account.
// CHECK
func (inst *StartPingPong) GetFeeQuoterBillingTokenConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[15]
}

// SetFeeQuoterLinkTokenConfigAccount sets the "feeQuoterLinkTokenConfig" account.
// CHECK
func (inst *StartPingPong) SetFeeQuoterLinkTokenConfigAccount(feeQuoterLinkTokenConfig ag_solanago.PublicKey) *StartPingPong {
	inst.AccountMetaSlice[16] = ag_solanago.Meta(feeQuoterLinkTokenConfig)
	return inst
}

// GetFeeQuoterLinkTokenConfigAccount gets the "feeQuoterLinkTokenConfig" account.
// CHECK
func (inst *StartPingPong) GetFeeQuoterLinkTokenConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[16]
}

// SetRmnRemoteAccount sets the "rmnRemote" account.
// CHECK
func (inst *StartPingPong) SetRmnRemoteAccount(rmnRemote ag_solanago.PublicKey) *StartPingPong {
	inst.AccountMetaSlice[17] = ag_solanago.Meta(rmnRemote)
	return inst
}

// GetRmnRemoteAccount gets the "rmnRemote" account.
// CHECK
func (inst *StartPingPong) GetRmnRemoteAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[17]
}

// SetRmnRemoteCursesAccount sets the "rmnRemoteCurses" account.
// CHECK
func (inst *StartPingPong) SetRmnRemoteCursesAccount(rmnRemoteCurses ag_solanago.PublicKey) *StartPingPong {
	inst.AccountMetaSlice[18] = ag_solanago.Meta(rmnRemoteCurses)
	return inst
}

// GetRmnRemoteCursesAccount gets the "rmnRemoteCurses" account.
// CHECK
func (inst *StartPingPong) GetRmnRemoteCursesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[18]
}

// SetRmnRemoteConfigAccount sets the "rmnRemoteConfig" account.
// CHECK
func (inst *StartPingPong) SetRmnRemoteConfigAccount(rmnRemoteConfig ag_solanago.PublicKey) *StartPingPong {
	inst.AccountMetaSlice[19] = ag_solanago.Meta(rmnRemoteConfig)
	return inst
}

// GetRmnRemoteConfigAccount gets the "rmnRemoteConfig" account.
// CHECK
func (inst *StartPingPong) GetRmnRemoteConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[19]
}

// SetTokenPoolsSignerAccount sets the "tokenPoolsSigner" account.
// CHECK
func (inst *StartPingPong) SetTokenPoolsSignerAccount(tokenPoolsSigner ag_solanago.PublicKey) *StartPingPong {
	inst.AccountMetaSlice[20] = ag_solanago.Meta(tokenPoolsSigner).WRITE()
	return inst
}

// GetTokenPoolsSignerAccount gets the "tokenPoolsSigner" account.
// CHECK
func (inst *StartPingPong) GetTokenPoolsSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[20]
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *StartPingPong) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *StartPingPong {
	inst.AccountMetaSlice[21] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *StartPingPong) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[21]
}

func (inst StartPingPong) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_StartPingPong,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst StartPingPong) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *StartPingPong) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.CcipSendSigner is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.FeeTokenProgram is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.FeeTokenMint is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.FeeTokenAta is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.CcipRouterProgram is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.CcipRouterConfig is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.CcipRouterDestChainState is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.CcipRouterNonce is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.CcipRouterFeeReceiver is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.CcipRouterFeeBillingSigner is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.FeeQuoter is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.FeeQuoterConfig is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.FeeQuoterDestChain is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.FeeQuoterBillingTokenConfig is not set")
		}
		if inst.AccountMetaSlice[16] == nil {
			return errors.New("accounts.FeeQuoterLinkTokenConfig is not set")
		}
		if inst.AccountMetaSlice[17] == nil {
			return errors.New("accounts.RmnRemote is not set")
		}
		if inst.AccountMetaSlice[18] == nil {
			return errors.New("accounts.RmnRemoteCurses is not set")
		}
		if inst.AccountMetaSlice[19] == nil {
			return errors.New("accounts.RmnRemoteConfig is not set")
		}
		if inst.AccountMetaSlice[20] == nil {
			return errors.New("accounts.TokenPoolsSigner is not set")
		}
		if inst.AccountMetaSlice[21] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *StartPingPong) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("StartPingPong")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=22]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                     config", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("                  authority", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("             ccipSendSigner", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("            feeTokenProgram", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("               feeTokenMint", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("                feeTokenAta", inst.AccountMetaSlice[5]))
						accountsBranch.Child(ag_format.Meta("          ccipRouterProgram", inst.AccountMetaSlice[6]))
						accountsBranch.Child(ag_format.Meta("           ccipRouterConfig", inst.AccountMetaSlice[7]))
						accountsBranch.Child(ag_format.Meta("   ccipRouterDestChainState", inst.AccountMetaSlice[8]))
						accountsBranch.Child(ag_format.Meta("            ccipRouterNonce", inst.AccountMetaSlice[9]))
						accountsBranch.Child(ag_format.Meta("      ccipRouterFeeReceiver", inst.AccountMetaSlice[10]))
						accountsBranch.Child(ag_format.Meta(" ccipRouterFeeBillingSigner", inst.AccountMetaSlice[11]))
						accountsBranch.Child(ag_format.Meta("                  feeQuoter", inst.AccountMetaSlice[12]))
						accountsBranch.Child(ag_format.Meta("            feeQuoterConfig", inst.AccountMetaSlice[13]))
						accountsBranch.Child(ag_format.Meta("         feeQuoterDestChain", inst.AccountMetaSlice[14]))
						accountsBranch.Child(ag_format.Meta("feeQuoterBillingTokenConfig", inst.AccountMetaSlice[15]))
						accountsBranch.Child(ag_format.Meta("   feeQuoterLinkTokenConfig", inst.AccountMetaSlice[16]))
						accountsBranch.Child(ag_format.Meta("                  rmnRemote", inst.AccountMetaSlice[17]))
						accountsBranch.Child(ag_format.Meta("            rmnRemoteCurses", inst.AccountMetaSlice[18]))
						accountsBranch.Child(ag_format.Meta("            rmnRemoteConfig", inst.AccountMetaSlice[19]))
						accountsBranch.Child(ag_format.Meta("           tokenPoolsSigner", inst.AccountMetaSlice[20]))
						accountsBranch.Child(ag_format.Meta("              systemProgram", inst.AccountMetaSlice[21]))
					})
				})
		})
}

func (obj StartPingPong) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *StartPingPong) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewStartPingPongInstruction declares a new StartPingPong instruction with the provided parameters and accounts.
func NewStartPingPongInstruction(
	// Accounts:
	config ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	ccipSendSigner ag_solanago.PublicKey,
	feeTokenProgram ag_solanago.PublicKey,
	feeTokenMint ag_solanago.PublicKey,
	feeTokenAta ag_solanago.PublicKey,
	ccipRouterProgram ag_solanago.PublicKey,
	ccipRouterConfig ag_solanago.PublicKey,
	ccipRouterDestChainState ag_solanago.PublicKey,
	ccipRouterNonce ag_solanago.PublicKey,
	ccipRouterFeeReceiver ag_solanago.PublicKey,
	ccipRouterFeeBillingSigner ag_solanago.PublicKey,
	feeQuoter ag_solanago.PublicKey,
	feeQuoterConfig ag_solanago.PublicKey,
	feeQuoterDestChain ag_solanago.PublicKey,
	feeQuoterBillingTokenConfig ag_solanago.PublicKey,
	feeQuoterLinkTokenConfig ag_solanago.PublicKey,
	rmnRemote ag_solanago.PublicKey,
	rmnRemoteCurses ag_solanago.PublicKey,
	rmnRemoteConfig ag_solanago.PublicKey,
	tokenPoolsSigner ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *StartPingPong {
	return NewStartPingPongInstructionBuilder().
		SetConfigAccount(config).
		SetAuthorityAccount(authority).
		SetCcipSendSignerAccount(ccipSendSigner).
		SetFeeTokenProgramAccount(feeTokenProgram).
		SetFeeTokenMintAccount(feeTokenMint).
		SetFeeTokenAtaAccount(feeTokenAta).
		SetCcipRouterProgramAccount(ccipRouterProgram).
		SetCcipRouterConfigAccount(ccipRouterConfig).
		SetCcipRouterDestChainStateAccount(ccipRouterDestChainState).
		SetCcipRouterNonceAccount(ccipRouterNonce).
		SetCcipRouterFeeReceiverAccount(ccipRouterFeeReceiver).
		SetCcipRouterFeeBillingSignerAccount(ccipRouterFeeBillingSigner).
		SetFeeQuoterAccount(feeQuoter).
		SetFeeQuoterConfigAccount(feeQuoterConfig).
		SetFeeQuoterDestChainAccount(feeQuoterDestChain).
		SetFeeQuoterBillingTokenConfigAccount(feeQuoterBillingTokenConfig).
		SetFeeQuoterLinkTokenConfigAccount(feeQuoterLinkTokenConfig).
		SetRmnRemoteAccount(rmnRemote).
		SetRmnRemoteCursesAccount(rmnRemoteCurses).
		SetRmnRemoteConfigAccount(rmnRemoteConfig).
		SetTokenPoolsSignerAccount(tokenPoolsSigner).
		SetSystemProgramAccount(systemProgram)
}
