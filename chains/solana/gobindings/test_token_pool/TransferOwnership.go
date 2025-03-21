// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package test_token_pool

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// TransferOwnership is the `transferOwnership` instruction.
type TransferOwnership struct {
	ProposedOwner *ag_solanago.PublicKey

	// [0] = [WRITE] state
	//
	// [1] = [] mint
	//
	// [2] = [SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewTransferOwnershipInstructionBuilder creates a new `TransferOwnership` instruction builder.
func NewTransferOwnershipInstructionBuilder() *TransferOwnership {
	nd := &TransferOwnership{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetProposedOwner sets the "proposedOwner" parameter.
func (inst *TransferOwnership) SetProposedOwner(proposedOwner ag_solanago.PublicKey) *TransferOwnership {
	inst.ProposedOwner = &proposedOwner
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *TransferOwnership) SetStateAccount(state ag_solanago.PublicKey) *TransferOwnership {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *TransferOwnership) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetMintAccount sets the "mint" account.
func (inst *TransferOwnership) SetMintAccount(mint ag_solanago.PublicKey) *TransferOwnership {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *TransferOwnership) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *TransferOwnership) SetAuthorityAccount(authority ag_solanago.PublicKey) *TransferOwnership {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *TransferOwnership) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

func (inst TransferOwnership) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_TransferOwnership,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst TransferOwnership) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *TransferOwnership) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.ProposedOwner == nil {
			return errors.New("ProposedOwner parameter is not set")
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

func (inst *TransferOwnership) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("TransferOwnership")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("ProposedOwner", *inst.ProposedOwner))
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

func (obj TransferOwnership) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ProposedOwner` param:
	err = encoder.Encode(obj.ProposedOwner)
	if err != nil {
		return err
	}
	return nil
}
func (obj *TransferOwnership) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ProposedOwner`:
	err = decoder.Decode(&obj.ProposedOwner)
	if err != nil {
		return err
	}
	return nil
}

// NewTransferOwnershipInstruction declares a new TransferOwnership instruction with the provided parameters and accounts.
func NewTransferOwnershipInstruction(
	// Parameters:
	proposedOwner ag_solanago.PublicKey,
	// Accounts:
	state ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *TransferOwnership {
	return NewTransferOwnershipInstructionBuilder().
		SetProposedOwner(proposedOwner).
		SetStateAccount(state).
		SetMintAccount(mint).
		SetAuthorityAccount(authority)
}
