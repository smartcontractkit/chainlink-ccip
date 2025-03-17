// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package example_lockrelease_token_pool

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SetCanAcceptLiquidity is the `setCanAcceptLiquidity` instruction.
type SetCanAcceptLiquidity struct {
	Allow *bool

	// [0] = [WRITE] state
	//
	// [1] = [] mint
	//
	// [2] = [SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewSetCanAcceptLiquidityInstructionBuilder creates a new `SetCanAcceptLiquidity` instruction builder.
func NewSetCanAcceptLiquidityInstructionBuilder() *SetCanAcceptLiquidity {
	nd := &SetCanAcceptLiquidity{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetAllow sets the "allow" parameter.
func (inst *SetCanAcceptLiquidity) SetAllow(allow bool) *SetCanAcceptLiquidity {
	inst.Allow = &allow
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *SetCanAcceptLiquidity) SetStateAccount(state ag_solanago.PublicKey) *SetCanAcceptLiquidity {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *SetCanAcceptLiquidity) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetMintAccount sets the "mint" account.
func (inst *SetCanAcceptLiquidity) SetMintAccount(mint ag_solanago.PublicKey) *SetCanAcceptLiquidity {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *SetCanAcceptLiquidity) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *SetCanAcceptLiquidity) SetAuthorityAccount(authority ag_solanago.PublicKey) *SetCanAcceptLiquidity {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *SetCanAcceptLiquidity) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

func (inst SetCanAcceptLiquidity) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SetCanAcceptLiquidity,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetCanAcceptLiquidity) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetCanAcceptLiquidity) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Allow == nil {
			return errors.New("Allow parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Authority is not set")
		}
	}
	return nil
}

func (inst *SetCanAcceptLiquidity) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetCanAcceptLiquidity")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Allow", *inst.Allow))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("    state", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("     mint", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("authority", inst.AccountMetaSlice[2]))
					})
				})
		})
}

func (obj SetCanAcceptLiquidity) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Allow` param:
	err = encoder.Encode(obj.Allow)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SetCanAcceptLiquidity) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Allow`:
	err = decoder.Decode(&obj.Allow)
	if err != nil {
		return err
	}
	return nil
}

// NewSetCanAcceptLiquidityInstruction declares a new SetCanAcceptLiquidity instruction with the provided parameters and accounts.
func NewSetCanAcceptLiquidityInstruction(
	// Parameters:
	allow bool,
	// Accounts:
	state ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *SetCanAcceptLiquidity {
	return NewSetCanAcceptLiquidityInstructionBuilder().
		SetAllow(allow).
		SetStateAccount(state).
		SetMintAccount(mint).
		SetAuthorityAccount(authority)
}
