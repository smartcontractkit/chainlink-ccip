// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ccip_router

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Removes the billing token configuration.
// Only CCIP Admin can remove a billing token configuration.
//
// # Arguments
//
// * `ctx` - The context containing the accounts required for removing the billing token configuration.
type RemoveBillingTokenConfig struct {

	// [0] = [] config
	//
	// [1] = [WRITE] billingTokenConfig
	//
	// [2] = [] tokenProgram
	//
	// [3] = [] feeTokenMint
	//
	// [4] = [WRITE] feeTokenReceiver
	//
	// [5] = [WRITE] feeBillingSigner
	//
	// [6] = [WRITE, SIGNER] authority
	//
	// [7] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewRemoveBillingTokenConfigInstructionBuilder creates a new `RemoveBillingTokenConfig` instruction builder.
func NewRemoveBillingTokenConfigInstructionBuilder() *RemoveBillingTokenConfig {
	nd := &RemoveBillingTokenConfig{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 8),
	}
	return nd
}

// SetConfigAccount sets the "config" account.
func (inst *RemoveBillingTokenConfig) SetConfigAccount(config ag_solanago.PublicKey) *RemoveBillingTokenConfig {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *RemoveBillingTokenConfig) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetBillingTokenConfigAccount sets the "billingTokenConfig" account.
func (inst *RemoveBillingTokenConfig) SetBillingTokenConfigAccount(billingTokenConfig ag_solanago.PublicKey) *RemoveBillingTokenConfig {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(billingTokenConfig).WRITE()
	return inst
}

// GetBillingTokenConfigAccount gets the "billingTokenConfig" account.
func (inst *RemoveBillingTokenConfig) GetBillingTokenConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *RemoveBillingTokenConfig) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *RemoveBillingTokenConfig {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *RemoveBillingTokenConfig) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetFeeTokenMintAccount sets the "feeTokenMint" account.
func (inst *RemoveBillingTokenConfig) SetFeeTokenMintAccount(feeTokenMint ag_solanago.PublicKey) *RemoveBillingTokenConfig {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(feeTokenMint)
	return inst
}

// GetFeeTokenMintAccount gets the "feeTokenMint" account.
func (inst *RemoveBillingTokenConfig) GetFeeTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetFeeTokenReceiverAccount sets the "feeTokenReceiver" account.
func (inst *RemoveBillingTokenConfig) SetFeeTokenReceiverAccount(feeTokenReceiver ag_solanago.PublicKey) *RemoveBillingTokenConfig {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(feeTokenReceiver).WRITE()
	return inst
}

// GetFeeTokenReceiverAccount gets the "feeTokenReceiver" account.
func (inst *RemoveBillingTokenConfig) GetFeeTokenReceiverAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetFeeBillingSignerAccount sets the "feeBillingSigner" account.
func (inst *RemoveBillingTokenConfig) SetFeeBillingSignerAccount(feeBillingSigner ag_solanago.PublicKey) *RemoveBillingTokenConfig {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(feeBillingSigner).WRITE()
	return inst
}

// GetFeeBillingSignerAccount gets the "feeBillingSigner" account.
func (inst *RemoveBillingTokenConfig) GetFeeBillingSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *RemoveBillingTokenConfig) SetAuthorityAccount(authority ag_solanago.PublicKey) *RemoveBillingTokenConfig {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *RemoveBillingTokenConfig) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *RemoveBillingTokenConfig) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *RemoveBillingTokenConfig {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *RemoveBillingTokenConfig) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[7]
}

func (inst RemoveBillingTokenConfig) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_RemoveBillingTokenConfig,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst RemoveBillingTokenConfig) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *RemoveBillingTokenConfig) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.BillingTokenConfig is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.FeeTokenMint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.FeeTokenReceiver is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.FeeBillingSigner is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *RemoveBillingTokenConfig) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("RemoveBillingTokenConfig")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=8]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("            config", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("billingTokenConfig", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("      tokenProgram", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("      feeTokenMint", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("  feeTokenReceiver", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("  feeBillingSigner", inst.AccountMetaSlice[5]))
						accountsBranch.Child(ag_format.Meta("         authority", inst.AccountMetaSlice[6]))
						accountsBranch.Child(ag_format.Meta("     systemProgram", inst.AccountMetaSlice[7]))
					})
				})
		})
}

func (obj RemoveBillingTokenConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *RemoveBillingTokenConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewRemoveBillingTokenConfigInstruction declares a new RemoveBillingTokenConfig instruction with the provided parameters and accounts.
func NewRemoveBillingTokenConfigInstruction(
	// Accounts:
	config ag_solanago.PublicKey,
	billingTokenConfig ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	feeTokenMint ag_solanago.PublicKey,
	feeTokenReceiver ag_solanago.PublicKey,
	feeBillingSigner ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *RemoveBillingTokenConfig {
	return NewRemoveBillingTokenConfigInstructionBuilder().
		SetConfigAccount(config).
		SetBillingTokenConfigAccount(billingTokenConfig).
		SetTokenProgramAccount(tokenProgram).
		SetFeeTokenMintAccount(feeTokenMint).
		SetFeeTokenReceiverAccount(feeTokenReceiver).
		SetFeeBillingSignerAccount(feeBillingSigner).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram)
}
