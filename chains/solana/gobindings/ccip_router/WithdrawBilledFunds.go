// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ccip_router

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Transfers the accumulated billed fees in a particular token to an arbitrary token account.
// Only the CCIP Admin can withdraw billed funds.
//
// # Arguments
//
// * `ctx` - The context containing the accounts required for the transfer of billed fees.
// * `transfer_all` - A flag indicating whether to transfer all the accumulated fees in that token or not.
// * `desired_amount` - The amount to transfer. If `transfer_all` is true, this value must be 0.
type WithdrawBilledFunds struct {
	TransferAll   *bool
	DesiredAmount *uint64

	// [0] = [] feeTokenMint
	//
	// [1] = [WRITE] feeTokenAccum
	//
	// [2] = [WRITE] recipient
	//
	// [3] = [] tokenProgram
	//
	// [4] = [] feeBillingSigner
	//
	// [5] = [] config
	//
	// [6] = [WRITE, SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewWithdrawBilledFundsInstructionBuilder creates a new `WithdrawBilledFunds` instruction builder.
func NewWithdrawBilledFundsInstructionBuilder() *WithdrawBilledFunds {
	nd := &WithdrawBilledFunds{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	return nd
}

// SetTransferAll sets the "transferAll" parameter.
func (inst *WithdrawBilledFunds) SetTransferAll(transferAll bool) *WithdrawBilledFunds {
	inst.TransferAll = &transferAll
	return inst
}

// SetDesiredAmount sets the "desiredAmount" parameter.
func (inst *WithdrawBilledFunds) SetDesiredAmount(desiredAmount uint64) *WithdrawBilledFunds {
	inst.DesiredAmount = &desiredAmount
	return inst
}

// SetFeeTokenMintAccount sets the "feeTokenMint" account.
func (inst *WithdrawBilledFunds) SetFeeTokenMintAccount(feeTokenMint ag_solanago.PublicKey) *WithdrawBilledFunds {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(feeTokenMint)
	return inst
}

// GetFeeTokenMintAccount gets the "feeTokenMint" account.
func (inst *WithdrawBilledFunds) GetFeeTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetFeeTokenAccumAccount sets the "feeTokenAccum" account.
func (inst *WithdrawBilledFunds) SetFeeTokenAccumAccount(feeTokenAccum ag_solanago.PublicKey) *WithdrawBilledFunds {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(feeTokenAccum).WRITE()
	return inst
}

// GetFeeTokenAccumAccount gets the "feeTokenAccum" account.
func (inst *WithdrawBilledFunds) GetFeeTokenAccumAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetRecipientAccount sets the "recipient" account.
func (inst *WithdrawBilledFunds) SetRecipientAccount(recipient ag_solanago.PublicKey) *WithdrawBilledFunds {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(recipient).WRITE()
	return inst
}

// GetRecipientAccount gets the "recipient" account.
func (inst *WithdrawBilledFunds) GetRecipientAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *WithdrawBilledFunds) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *WithdrawBilledFunds {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *WithdrawBilledFunds) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetFeeBillingSignerAccount sets the "feeBillingSigner" account.
func (inst *WithdrawBilledFunds) SetFeeBillingSignerAccount(feeBillingSigner ag_solanago.PublicKey) *WithdrawBilledFunds {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(feeBillingSigner)
	return inst
}

// GetFeeBillingSignerAccount gets the "feeBillingSigner" account.
func (inst *WithdrawBilledFunds) GetFeeBillingSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetConfigAccount sets the "config" account.
func (inst *WithdrawBilledFunds) SetConfigAccount(config ag_solanago.PublicKey) *WithdrawBilledFunds {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *WithdrawBilledFunds) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *WithdrawBilledFunds) SetAuthorityAccount(authority ag_solanago.PublicKey) *WithdrawBilledFunds {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *WithdrawBilledFunds) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

func (inst WithdrawBilledFunds) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_WithdrawBilledFunds,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst WithdrawBilledFunds) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *WithdrawBilledFunds) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.TransferAll == nil {
			return errors.New("TransferAll parameter is not set")
		}
		if inst.DesiredAmount == nil {
			return errors.New("DesiredAmount parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.FeeTokenMint is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.FeeTokenAccum is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Recipient is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.FeeBillingSigner is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Authority is not set")
		}
	}
	return nil
}

func (inst *WithdrawBilledFunds) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("WithdrawBilledFunds")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("  TransferAll", *inst.TransferAll))
						paramsBranch.Child(ag_format.Param("DesiredAmount", *inst.DesiredAmount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("    feeTokenMint", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("   feeTokenAccum", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("       recipient", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("    tokenProgram", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("feeBillingSigner", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("          config", inst.AccountMetaSlice[5]))
						accountsBranch.Child(ag_format.Meta("       authority", inst.AccountMetaSlice[6]))
					})
				})
		})
}

func (obj WithdrawBilledFunds) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `TransferAll` param:
	err = encoder.Encode(obj.TransferAll)
	if err != nil {
		return err
	}
	// Serialize `DesiredAmount` param:
	err = encoder.Encode(obj.DesiredAmount)
	if err != nil {
		return err
	}
	return nil
}
func (obj *WithdrawBilledFunds) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `TransferAll`:
	err = decoder.Decode(&obj.TransferAll)
	if err != nil {
		return err
	}
	// Deserialize `DesiredAmount`:
	err = decoder.Decode(&obj.DesiredAmount)
	if err != nil {
		return err
	}
	return nil
}

// NewWithdrawBilledFundsInstruction declares a new WithdrawBilledFunds instruction with the provided parameters and accounts.
func NewWithdrawBilledFundsInstruction(
	// Parameters:
	transferAll bool,
	desiredAmount uint64,
	// Accounts:
	feeTokenMint ag_solanago.PublicKey,
	feeTokenAccum ag_solanago.PublicKey,
	recipient ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	feeBillingSigner ag_solanago.PublicKey,
	config ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *WithdrawBilledFunds {
	return NewWithdrawBilledFundsInstructionBuilder().
		SetTransferAll(transferAll).
		SetDesiredAmount(desiredAmount).
		SetFeeTokenMintAccount(feeTokenMint).
		SetFeeTokenAccumAccount(feeTokenAccum).
		SetRecipientAccount(recipient).
		SetTokenProgramAccount(tokenProgram).
		SetFeeBillingSignerAccount(feeBillingSigner).
		SetConfigAccount(config).
		SetAuthorityAccount(authority)
}