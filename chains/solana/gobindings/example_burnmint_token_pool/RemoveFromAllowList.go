// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package example_burnmint_token_pool

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// RemoveFromAllowList is the `removeFromAllowList` instruction.
type RemoveFromAllowList struct {
	Remove *[]ag_solanago.PublicKey

	// [0] = [WRITE] state
	//
	// [1] = [] mint
	//
	// [2] = [SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewRemoveFromAllowListInstructionBuilder creates a new `RemoveFromAllowList` instruction builder.
func NewRemoveFromAllowListInstructionBuilder() *RemoveFromAllowList {
	nd := &RemoveFromAllowList{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetRemove sets the "remove" parameter.
func (inst *RemoveFromAllowList) SetRemove(remove []ag_solanago.PublicKey) *RemoveFromAllowList {
	inst.Remove = &remove
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *RemoveFromAllowList) SetStateAccount(state ag_solanago.PublicKey) *RemoveFromAllowList {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *RemoveFromAllowList) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetMintAccount sets the "mint" account.
func (inst *RemoveFromAllowList) SetMintAccount(mint ag_solanago.PublicKey) *RemoveFromAllowList {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *RemoveFromAllowList) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *RemoveFromAllowList) SetAuthorityAccount(authority ag_solanago.PublicKey) *RemoveFromAllowList {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *RemoveFromAllowList) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

func (inst RemoveFromAllowList) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_RemoveFromAllowList,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst RemoveFromAllowList) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *RemoveFromAllowList) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Remove == nil {
			return errors.New("Remove parameter is not set")
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

func (inst *RemoveFromAllowList) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("RemoveFromAllowList")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Remove", *inst.Remove))
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

func (obj RemoveFromAllowList) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Remove` param:
	err = encoder.Encode(obj.Remove)
	if err != nil {
		return err
	}
	return nil
}
func (obj *RemoveFromAllowList) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Remove`:
	err = decoder.Decode(&obj.Remove)
	if err != nil {
		return err
	}
	return nil
}

// NewRemoveFromAllowListInstruction declares a new RemoveFromAllowList instruction with the provided parameters and accounts.
func NewRemoveFromAllowListInstruction(
	// Parameters:
	remove []ag_solanago.PublicKey,
	// Accounts:
	state ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *RemoveFromAllowList {
	return NewRemoveFromAllowListInstructionBuilder().
		SetRemove(remove).
		SetStateAccount(state).
		SetMintAccount(mint).
		SetAuthorityAccount(authority)
}
