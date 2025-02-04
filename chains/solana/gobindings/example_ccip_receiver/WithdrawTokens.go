// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package example_ccip_receiver

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// WithdrawTokens is the `withdrawTokens` instruction.
type WithdrawTokens struct {
	Amount   *uint64
	Decimals *uint8

	// [0] = [WRITE] state
	//
	// [1] = [WRITE] programTokenAccount
	//
	// [2] = [WRITE] toTokenAccount
	//
	// [3] = [] mint
	//
	// [4] = [] tokenProgram
	//
	// [5] = [] tokenAdmin
	//
	// [6] = [SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewWithdrawTokensInstructionBuilder creates a new `WithdrawTokens` instruction builder.
func NewWithdrawTokensInstructionBuilder() *WithdrawTokens {
	nd := &WithdrawTokens{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	return nd
}

// SetAmount sets the "amount" parameter.
func (inst *WithdrawTokens) SetAmount(amount uint64) *WithdrawTokens {
	inst.Amount = &amount
	return inst
}

// SetDecimals sets the "decimals" parameter.
func (inst *WithdrawTokens) SetDecimals(decimals uint8) *WithdrawTokens {
	inst.Decimals = &decimals
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *WithdrawTokens) SetStateAccount(state ag_solanago.PublicKey) *WithdrawTokens {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *WithdrawTokens) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetProgramTokenAccountAccount sets the "programTokenAccount" account.
func (inst *WithdrawTokens) SetProgramTokenAccountAccount(programTokenAccount ag_solanago.PublicKey) *WithdrawTokens {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(programTokenAccount).WRITE()
	return inst
}

// GetProgramTokenAccountAccount gets the "programTokenAccount" account.
func (inst *WithdrawTokens) GetProgramTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetToTokenAccountAccount sets the "toTokenAccount" account.
func (inst *WithdrawTokens) SetToTokenAccountAccount(toTokenAccount ag_solanago.PublicKey) *WithdrawTokens {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(toTokenAccount).WRITE()
	return inst
}

// GetToTokenAccountAccount gets the "toTokenAccount" account.
func (inst *WithdrawTokens) GetToTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetMintAccount sets the "mint" account.
func (inst *WithdrawTokens) SetMintAccount(mint ag_solanago.PublicKey) *WithdrawTokens {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *WithdrawTokens) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *WithdrawTokens) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *WithdrawTokens {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *WithdrawTokens) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetTokenAdminAccount sets the "tokenAdmin" account.
func (inst *WithdrawTokens) SetTokenAdminAccount(tokenAdmin ag_solanago.PublicKey) *WithdrawTokens {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(tokenAdmin)
	return inst
}

// GetTokenAdminAccount gets the "tokenAdmin" account.
func (inst *WithdrawTokens) GetTokenAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *WithdrawTokens) SetAuthorityAccount(authority ag_solanago.PublicKey) *WithdrawTokens {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *WithdrawTokens) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

func (inst WithdrawTokens) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_WithdrawTokens,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst WithdrawTokens) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *WithdrawTokens) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Amount == nil {
			return errors.New("Amount parameter is not set")
		}
		if inst.Decimals == nil {
			return errors.New("Decimals parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ProgramTokenAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.ToTokenAccount is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.TokenAdmin is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Authority is not set")
		}
	}
	return nil
}

func (inst *WithdrawTokens) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("WithdrawTokens")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("  Amount", *inst.Amount))
						paramsBranch.Child(ag_format.Param("Decimals", *inst.Decimals))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       state", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("programToken", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("     toToken", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("        mint", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("tokenProgram", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("  tokenAdmin", inst.AccountMetaSlice[5]))
						accountsBranch.Child(ag_format.Meta("   authority", inst.AccountMetaSlice[6]))
					})
				})
		})
}

func (obj WithdrawTokens) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	// Serialize `Decimals` param:
	err = encoder.Encode(obj.Decimals)
	if err != nil {
		return err
	}
	return nil
}
func (obj *WithdrawTokens) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	// Deserialize `Decimals`:
	err = decoder.Decode(&obj.Decimals)
	if err != nil {
		return err
	}
	return nil
}

// NewWithdrawTokensInstruction declares a new WithdrawTokens instruction with the provided parameters and accounts.
func NewWithdrawTokensInstruction(
	// Parameters:
	amount uint64,
	decimals uint8,
	// Accounts:
	state ag_solanago.PublicKey,
	programTokenAccount ag_solanago.PublicKey,
	toTokenAccount ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	tokenAdmin ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *WithdrawTokens {
	return NewWithdrawTokensInstructionBuilder().
		SetAmount(amount).
		SetDecimals(decimals).
		SetStateAccount(state).
		SetProgramTokenAccountAccount(programTokenAccount).
		SetToTokenAccountAccount(toTokenAccount).
		SetMintAccount(mint).
		SetTokenProgramAccount(tokenProgram).
		SetTokenAdminAccount(tokenAdmin).
		SetAuthorityAccount(authority)
}
