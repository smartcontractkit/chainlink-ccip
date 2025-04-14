// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package lockrelease_token_pool

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// InitializeStateVersion is the `initializeStateVersion` instruction.
type InitializeStateVersion struct {
	Mint *ag_solanago.PublicKey

	// [0] = [WRITE] state
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewInitializeStateVersionInstructionBuilder creates a new `InitializeStateVersion` instruction builder.
func NewInitializeStateVersionInstructionBuilder() *InitializeStateVersion {
	nd := &InitializeStateVersion{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 1),
	}
	return nd
}

// SetMint sets the "mint" parameter.
func (inst *InitializeStateVersion) SetMint(mint ag_solanago.PublicKey) *InitializeStateVersion {
	inst.Mint = &mint
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *InitializeStateVersion) SetStateAccount(state ag_solanago.PublicKey) *InitializeStateVersion {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *InitializeStateVersion) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

func (inst InitializeStateVersion) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitializeStateVersion,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitializeStateVersion) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitializeStateVersion) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Mint == nil {
			return errors.New("Mint parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
		}
	}
	return nil
}

func (inst *InitializeStateVersion) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitializeStateVersion")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Mint", *inst.Mint))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=1]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("state", inst.AccountMetaSlice[0]))
					})
				})
		})
}

func (obj InitializeStateVersion) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Mint` param:
	err = encoder.Encode(obj.Mint)
	if err != nil {
		return err
	}
	return nil
}
func (obj *InitializeStateVersion) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Mint`:
	err = decoder.Decode(&obj.Mint)
	if err != nil {
		return err
	}
	return nil
}

// NewInitializeStateVersionInstruction declares a new InitializeStateVersion instruction with the provided parameters and accounts.
func NewInitializeStateVersionInstruction(
	// Parameters:
	mint ag_solanago.PublicKey,
	// Accounts:
	state ag_solanago.PublicKey) *InitializeStateVersion {
	return NewInitializeStateVersionInstructionBuilder().
		SetMint(mint).
		SetStateAccount(state)
}
