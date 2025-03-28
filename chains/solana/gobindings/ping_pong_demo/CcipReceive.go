// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ping_pong_demo

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CcipReceive is the `ccipReceive` instruction.
type CcipReceive struct {
	Message *Any2SVMMessage

	// [0] = [SIGNER] authority
	//
	// [1] = [] offrampProgram
	// ··········· CHECK offramp program: exists only to derive the allowed offramp PDA
	// ··········· and the authority PDA. Must be second.
	//
	// [2] = [] allowedOfframp
	// ··········· CHECK PDA of the router program verifying the signer is an allowed offramp.
	// ··········· If PDA does not exist, the router doesn't allow this offramp
	//
	// [3] = [] globalConfig
	//
	// [4] = [] config
	//
	// [5] = [WRITE] ccipSendSigner
	// ··········· CHECK
	//
	// [6] = [] feeTokenProgram
	//
	// [7] = [] feeTokenMint
	//
	// [8] = [WRITE] feeTokenAta
	//
	// [9] = [] ccipRouterProgram
	// ··········· CHECK
	//
	// [10] = [] ccipRouterConfig
	// ··········· CHECK
	//
	// [11] = [WRITE] ccipRouterDestChainState
	// ··········· CHECK
	//
	// [12] = [WRITE] ccipRouterNonce
	// ··········· CHECK
	//
	// [13] = [WRITE] ccipRouterFeeReceiver
	// ··········· CHECK
	//
	// [14] = [] ccipRouterFeeBillingSigner
	// ··········· CHECK
	//
	// [15] = [] feeQuoter
	// ··········· CHECK
	//
	// [16] = [] feeQuoterConfig
	// ··········· CHECK
	//
	// [17] = [] feeQuoterDestChain
	// ··········· CHECK
	//
	// [18] = [] feeQuoterBillingTokenConfig
	// ··········· CHECK
	//
	// [19] = [] feeQuoterLinkTokenConfig
	// ··········· CHECK
	//
	// [20] = [] rmnRemote
	// ··········· CHECK
	//
	// [21] = [] rmnRemoteCurses
	// ··········· CHECK
	//
	// [22] = [] rmnRemoteConfig
	// ··········· CHECK
	//
	// [23] = [WRITE] tokenPoolsSigner
	// ··········· CHECK
	//
	// [24] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewCcipReceiveInstructionBuilder creates a new `CcipReceive` instruction builder.
func NewCcipReceiveInstructionBuilder() *CcipReceive {
	nd := &CcipReceive{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 25),
	}
	return nd
}

// SetMessage sets the "message" parameter.
func (inst *CcipReceive) SetMessage(message Any2SVMMessage) *CcipReceive {
	inst.Message = &message
	return inst
}

// SetAuthorityAccount sets the "authority" account.
func (inst *CcipReceive) SetAuthorityAccount(authority ag_solanago.PublicKey) *CcipReceive {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *CcipReceive) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetOfframpProgramAccount sets the "offrampProgram" account.
// CHECK offramp program: exists only to derive the allowed offramp PDA
// and the authority PDA. Must be second.
func (inst *CcipReceive) SetOfframpProgramAccount(offrampProgram ag_solanago.PublicKey) *CcipReceive {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(offrampProgram)
	return inst
}

// GetOfframpProgramAccount gets the "offrampProgram" account.
// CHECK offramp program: exists only to derive the allowed offramp PDA
// and the authority PDA. Must be second.
func (inst *CcipReceive) GetOfframpProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetAllowedOfframpAccount sets the "allowedOfframp" account.
// CHECK PDA of the router program verifying the signer is an allowed offramp.
// If PDA does not exist, the router doesn't allow this offramp
func (inst *CcipReceive) SetAllowedOfframpAccount(allowedOfframp ag_solanago.PublicKey) *CcipReceive {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(allowedOfframp)
	return inst
}

// GetAllowedOfframpAccount gets the "allowedOfframp" account.
// CHECK PDA of the router program verifying the signer is an allowed offramp.
// If PDA does not exist, the router doesn't allow this offramp
func (inst *CcipReceive) GetAllowedOfframpAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetGlobalConfigAccount sets the "globalConfig" account.
func (inst *CcipReceive) SetGlobalConfigAccount(globalConfig ag_solanago.PublicKey) *CcipReceive {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(globalConfig)
	return inst
}

// GetGlobalConfigAccount gets the "globalConfig" account.
func (inst *CcipReceive) GetGlobalConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetConfigAccount sets the "config" account.
func (inst *CcipReceive) SetConfigAccount(config ag_solanago.PublicKey) *CcipReceive {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *CcipReceive) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetCcipSendSignerAccount sets the "ccipSendSigner" account.
// CHECK
func (inst *CcipReceive) SetCcipSendSignerAccount(ccipSendSigner ag_solanago.PublicKey) *CcipReceive {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(ccipSendSigner).WRITE()
	return inst
}

// GetCcipSendSignerAccount gets the "ccipSendSigner" account.
// CHECK
func (inst *CcipReceive) GetCcipSendSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetFeeTokenProgramAccount sets the "feeTokenProgram" account.
func (inst *CcipReceive) SetFeeTokenProgramAccount(feeTokenProgram ag_solanago.PublicKey) *CcipReceive {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(feeTokenProgram)
	return inst
}

// GetFeeTokenProgramAccount gets the "feeTokenProgram" account.
func (inst *CcipReceive) GetFeeTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

// SetFeeTokenMintAccount sets the "feeTokenMint" account.
func (inst *CcipReceive) SetFeeTokenMintAccount(feeTokenMint ag_solanago.PublicKey) *CcipReceive {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(feeTokenMint)
	return inst
}

// GetFeeTokenMintAccount gets the "feeTokenMint" account.
func (inst *CcipReceive) GetFeeTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[7]
}

// SetFeeTokenAtaAccount sets the "feeTokenAta" account.
func (inst *CcipReceive) SetFeeTokenAtaAccount(feeTokenAta ag_solanago.PublicKey) *CcipReceive {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(feeTokenAta).WRITE()
	return inst
}

// GetFeeTokenAtaAccount gets the "feeTokenAta" account.
func (inst *CcipReceive) GetFeeTokenAtaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[8]
}

// SetCcipRouterProgramAccount sets the "ccipRouterProgram" account.
// CHECK
func (inst *CcipReceive) SetCcipRouterProgramAccount(ccipRouterProgram ag_solanago.PublicKey) *CcipReceive {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(ccipRouterProgram)
	return inst
}

// GetCcipRouterProgramAccount gets the "ccipRouterProgram" account.
// CHECK
func (inst *CcipReceive) GetCcipRouterProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[9]
}

// SetCcipRouterConfigAccount sets the "ccipRouterConfig" account.
// CHECK
func (inst *CcipReceive) SetCcipRouterConfigAccount(ccipRouterConfig ag_solanago.PublicKey) *CcipReceive {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(ccipRouterConfig)
	return inst
}

// GetCcipRouterConfigAccount gets the "ccipRouterConfig" account.
// CHECK
func (inst *CcipReceive) GetCcipRouterConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[10]
}

// SetCcipRouterDestChainStateAccount sets the "ccipRouterDestChainState" account.
// CHECK
func (inst *CcipReceive) SetCcipRouterDestChainStateAccount(ccipRouterDestChainState ag_solanago.PublicKey) *CcipReceive {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(ccipRouterDestChainState).WRITE()
	return inst
}

// GetCcipRouterDestChainStateAccount gets the "ccipRouterDestChainState" account.
// CHECK
func (inst *CcipReceive) GetCcipRouterDestChainStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[11]
}

// SetCcipRouterNonceAccount sets the "ccipRouterNonce" account.
// CHECK
func (inst *CcipReceive) SetCcipRouterNonceAccount(ccipRouterNonce ag_solanago.PublicKey) *CcipReceive {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(ccipRouterNonce).WRITE()
	return inst
}

// GetCcipRouterNonceAccount gets the "ccipRouterNonce" account.
// CHECK
func (inst *CcipReceive) GetCcipRouterNonceAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[12]
}

// SetCcipRouterFeeReceiverAccount sets the "ccipRouterFeeReceiver" account.
// CHECK
func (inst *CcipReceive) SetCcipRouterFeeReceiverAccount(ccipRouterFeeReceiver ag_solanago.PublicKey) *CcipReceive {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(ccipRouterFeeReceiver).WRITE()
	return inst
}

// GetCcipRouterFeeReceiverAccount gets the "ccipRouterFeeReceiver" account.
// CHECK
func (inst *CcipReceive) GetCcipRouterFeeReceiverAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[13]
}

// SetCcipRouterFeeBillingSignerAccount sets the "ccipRouterFeeBillingSigner" account.
// CHECK
func (inst *CcipReceive) SetCcipRouterFeeBillingSignerAccount(ccipRouterFeeBillingSigner ag_solanago.PublicKey) *CcipReceive {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(ccipRouterFeeBillingSigner)
	return inst
}

// GetCcipRouterFeeBillingSignerAccount gets the "ccipRouterFeeBillingSigner" account.
// CHECK
func (inst *CcipReceive) GetCcipRouterFeeBillingSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[14]
}

// SetFeeQuoterAccount sets the "feeQuoter" account.
// CHECK
func (inst *CcipReceive) SetFeeQuoterAccount(feeQuoter ag_solanago.PublicKey) *CcipReceive {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(feeQuoter)
	return inst
}

// GetFeeQuoterAccount gets the "feeQuoter" account.
// CHECK
func (inst *CcipReceive) GetFeeQuoterAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[15]
}

// SetFeeQuoterConfigAccount sets the "feeQuoterConfig" account.
// CHECK
func (inst *CcipReceive) SetFeeQuoterConfigAccount(feeQuoterConfig ag_solanago.PublicKey) *CcipReceive {
	inst.AccountMetaSlice[16] = ag_solanago.Meta(feeQuoterConfig)
	return inst
}

// GetFeeQuoterConfigAccount gets the "feeQuoterConfig" account.
// CHECK
func (inst *CcipReceive) GetFeeQuoterConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[16]
}

// SetFeeQuoterDestChainAccount sets the "feeQuoterDestChain" account.
// CHECK
func (inst *CcipReceive) SetFeeQuoterDestChainAccount(feeQuoterDestChain ag_solanago.PublicKey) *CcipReceive {
	inst.AccountMetaSlice[17] = ag_solanago.Meta(feeQuoterDestChain)
	return inst
}

// GetFeeQuoterDestChainAccount gets the "feeQuoterDestChain" account.
// CHECK
func (inst *CcipReceive) GetFeeQuoterDestChainAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[17]
}

// SetFeeQuoterBillingTokenConfigAccount sets the "feeQuoterBillingTokenConfig" account.
// CHECK
func (inst *CcipReceive) SetFeeQuoterBillingTokenConfigAccount(feeQuoterBillingTokenConfig ag_solanago.PublicKey) *CcipReceive {
	inst.AccountMetaSlice[18] = ag_solanago.Meta(feeQuoterBillingTokenConfig)
	return inst
}

// GetFeeQuoterBillingTokenConfigAccount gets the "feeQuoterBillingTokenConfig" account.
// CHECK
func (inst *CcipReceive) GetFeeQuoterBillingTokenConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[18]
}

// SetFeeQuoterLinkTokenConfigAccount sets the "feeQuoterLinkTokenConfig" account.
// CHECK
func (inst *CcipReceive) SetFeeQuoterLinkTokenConfigAccount(feeQuoterLinkTokenConfig ag_solanago.PublicKey) *CcipReceive {
	inst.AccountMetaSlice[19] = ag_solanago.Meta(feeQuoterLinkTokenConfig)
	return inst
}

// GetFeeQuoterLinkTokenConfigAccount gets the "feeQuoterLinkTokenConfig" account.
// CHECK
func (inst *CcipReceive) GetFeeQuoterLinkTokenConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[19]
}

// SetRmnRemoteAccount sets the "rmnRemote" account.
// CHECK
func (inst *CcipReceive) SetRmnRemoteAccount(rmnRemote ag_solanago.PublicKey) *CcipReceive {
	inst.AccountMetaSlice[20] = ag_solanago.Meta(rmnRemote)
	return inst
}

// GetRmnRemoteAccount gets the "rmnRemote" account.
// CHECK
func (inst *CcipReceive) GetRmnRemoteAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[20]
}

// SetRmnRemoteCursesAccount sets the "rmnRemoteCurses" account.
// CHECK
func (inst *CcipReceive) SetRmnRemoteCursesAccount(rmnRemoteCurses ag_solanago.PublicKey) *CcipReceive {
	inst.AccountMetaSlice[21] = ag_solanago.Meta(rmnRemoteCurses)
	return inst
}

// GetRmnRemoteCursesAccount gets the "rmnRemoteCurses" account.
// CHECK
func (inst *CcipReceive) GetRmnRemoteCursesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[21]
}

// SetRmnRemoteConfigAccount sets the "rmnRemoteConfig" account.
// CHECK
func (inst *CcipReceive) SetRmnRemoteConfigAccount(rmnRemoteConfig ag_solanago.PublicKey) *CcipReceive {
	inst.AccountMetaSlice[22] = ag_solanago.Meta(rmnRemoteConfig)
	return inst
}

// GetRmnRemoteConfigAccount gets the "rmnRemoteConfig" account.
// CHECK
func (inst *CcipReceive) GetRmnRemoteConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[22]
}

// SetTokenPoolsSignerAccount sets the "tokenPoolsSigner" account.
// CHECK
func (inst *CcipReceive) SetTokenPoolsSignerAccount(tokenPoolsSigner ag_solanago.PublicKey) *CcipReceive {
	inst.AccountMetaSlice[23] = ag_solanago.Meta(tokenPoolsSigner).WRITE()
	return inst
}

// GetTokenPoolsSignerAccount gets the "tokenPoolsSigner" account.
// CHECK
func (inst *CcipReceive) GetTokenPoolsSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[23]
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *CcipReceive) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *CcipReceive {
	inst.AccountMetaSlice[24] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *CcipReceive) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[24]
}

func (inst CcipReceive) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CcipReceive,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CcipReceive) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CcipReceive) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Message == nil {
			return errors.New("Message parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.OfframpProgram is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.AllowedOfframp is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.GlobalConfig is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.CcipSendSigner is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.FeeTokenProgram is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.FeeTokenMint is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.FeeTokenAta is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.CcipRouterProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.CcipRouterConfig is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.CcipRouterDestChainState is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.CcipRouterNonce is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.CcipRouterFeeReceiver is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.CcipRouterFeeBillingSigner is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.FeeQuoter is not set")
		}
		if inst.AccountMetaSlice[16] == nil {
			return errors.New("accounts.FeeQuoterConfig is not set")
		}
		if inst.AccountMetaSlice[17] == nil {
			return errors.New("accounts.FeeQuoterDestChain is not set")
		}
		if inst.AccountMetaSlice[18] == nil {
			return errors.New("accounts.FeeQuoterBillingTokenConfig is not set")
		}
		if inst.AccountMetaSlice[19] == nil {
			return errors.New("accounts.FeeQuoterLinkTokenConfig is not set")
		}
		if inst.AccountMetaSlice[20] == nil {
			return errors.New("accounts.RmnRemote is not set")
		}
		if inst.AccountMetaSlice[21] == nil {
			return errors.New("accounts.RmnRemoteCurses is not set")
		}
		if inst.AccountMetaSlice[22] == nil {
			return errors.New("accounts.RmnRemoteConfig is not set")
		}
		if inst.AccountMetaSlice[23] == nil {
			return errors.New("accounts.TokenPoolsSigner is not set")
		}
		if inst.AccountMetaSlice[24] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *CcipReceive) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CcipReceive")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Message", *inst.Message))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=25]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                  authority", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("             offrampProgram", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("             allowedOfframp", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("               globalConfig", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("                     config", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("             ccipSendSigner", inst.AccountMetaSlice[5]))
						accountsBranch.Child(ag_format.Meta("            feeTokenProgram", inst.AccountMetaSlice[6]))
						accountsBranch.Child(ag_format.Meta("               feeTokenMint", inst.AccountMetaSlice[7]))
						accountsBranch.Child(ag_format.Meta("                feeTokenAta", inst.AccountMetaSlice[8]))
						accountsBranch.Child(ag_format.Meta("          ccipRouterProgram", inst.AccountMetaSlice[9]))
						accountsBranch.Child(ag_format.Meta("           ccipRouterConfig", inst.AccountMetaSlice[10]))
						accountsBranch.Child(ag_format.Meta("   ccipRouterDestChainState", inst.AccountMetaSlice[11]))
						accountsBranch.Child(ag_format.Meta("            ccipRouterNonce", inst.AccountMetaSlice[12]))
						accountsBranch.Child(ag_format.Meta("      ccipRouterFeeReceiver", inst.AccountMetaSlice[13]))
						accountsBranch.Child(ag_format.Meta(" ccipRouterFeeBillingSigner", inst.AccountMetaSlice[14]))
						accountsBranch.Child(ag_format.Meta("                  feeQuoter", inst.AccountMetaSlice[15]))
						accountsBranch.Child(ag_format.Meta("            feeQuoterConfig", inst.AccountMetaSlice[16]))
						accountsBranch.Child(ag_format.Meta("         feeQuoterDestChain", inst.AccountMetaSlice[17]))
						accountsBranch.Child(ag_format.Meta("feeQuoterBillingTokenConfig", inst.AccountMetaSlice[18]))
						accountsBranch.Child(ag_format.Meta("   feeQuoterLinkTokenConfig", inst.AccountMetaSlice[19]))
						accountsBranch.Child(ag_format.Meta("                  rmnRemote", inst.AccountMetaSlice[20]))
						accountsBranch.Child(ag_format.Meta("            rmnRemoteCurses", inst.AccountMetaSlice[21]))
						accountsBranch.Child(ag_format.Meta("            rmnRemoteConfig", inst.AccountMetaSlice[22]))
						accountsBranch.Child(ag_format.Meta("           tokenPoolsSigner", inst.AccountMetaSlice[23]))
						accountsBranch.Child(ag_format.Meta("              systemProgram", inst.AccountMetaSlice[24]))
					})
				})
		})
}

func (obj CcipReceive) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Message` param:
	err = encoder.Encode(obj.Message)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CcipReceive) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Message`:
	err = decoder.Decode(&obj.Message)
	if err != nil {
		return err
	}
	return nil
}

// NewCcipReceiveInstruction declares a new CcipReceive instruction with the provided parameters and accounts.
func NewCcipReceiveInstruction(
	// Parameters:
	message Any2SVMMessage,
	// Accounts:
	authority ag_solanago.PublicKey,
	offrampProgram ag_solanago.PublicKey,
	allowedOfframp ag_solanago.PublicKey,
	globalConfig ag_solanago.PublicKey,
	config ag_solanago.PublicKey,
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
	systemProgram ag_solanago.PublicKey) *CcipReceive {
	return NewCcipReceiveInstructionBuilder().
		SetMessage(message).
		SetAuthorityAccount(authority).
		SetOfframpProgramAccount(offrampProgram).
		SetAllowedOfframpAccount(allowedOfframp).
		SetGlobalConfigAccount(globalConfig).
		SetConfigAccount(config).
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
