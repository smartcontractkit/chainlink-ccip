// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ccip_router

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Registers the Token Admin Registry via the token owner.
//
// The Authority of the Mint Token can claim the registry of the token.
//
// # Arguments
//
// * `ctx` - The context containing the accounts required for registration.
type RegisterTokenAdminRegistryViaOwner struct {

	// [0] = [WRITE] tokenAdminRegistry
	//
	// [1] = [WRITE] mint
	//
	// [2] = [WRITE, SIGNER] authority
	//
	// [3] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewRegisterTokenAdminRegistryViaOwnerInstructionBuilder creates a new `RegisterTokenAdminRegistryViaOwner` instruction builder.
func NewRegisterTokenAdminRegistryViaOwnerInstructionBuilder() *RegisterTokenAdminRegistryViaOwner {
	nd := &RegisterTokenAdminRegistryViaOwner{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetTokenAdminRegistryAccount sets the "tokenAdminRegistry" account.
func (inst *RegisterTokenAdminRegistryViaOwner) SetTokenAdminRegistryAccount(tokenAdminRegistry ag_solanago.PublicKey) *RegisterTokenAdminRegistryViaOwner {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(tokenAdminRegistry).WRITE()
	return inst
}

// GetTokenAdminRegistryAccount gets the "tokenAdminRegistry" account.
func (inst *RegisterTokenAdminRegistryViaOwner) GetTokenAdminRegistryAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetMintAccount sets the "mint" account.
func (inst *RegisterTokenAdminRegistryViaOwner) SetMintAccount(mint ag_solanago.PublicKey) *RegisterTokenAdminRegistryViaOwner {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(mint).WRITE()
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *RegisterTokenAdminRegistryViaOwner) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *RegisterTokenAdminRegistryViaOwner) SetAuthorityAccount(authority ag_solanago.PublicKey) *RegisterTokenAdminRegistryViaOwner {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *RegisterTokenAdminRegistryViaOwner) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *RegisterTokenAdminRegistryViaOwner) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *RegisterTokenAdminRegistryViaOwner {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *RegisterTokenAdminRegistryViaOwner) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

func (inst RegisterTokenAdminRegistryViaOwner) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_RegisterTokenAdminRegistryViaOwner,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst RegisterTokenAdminRegistryViaOwner) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *RegisterTokenAdminRegistryViaOwner) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.TokenAdminRegistry is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *RegisterTokenAdminRegistryViaOwner) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("RegisterTokenAdminRegistryViaOwner")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("tokenAdminRegistry", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("              mint", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("         authority", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("     systemProgram", inst.AccountMetaSlice[3]))
					})
				})
		})
}

func (obj RegisterTokenAdminRegistryViaOwner) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *RegisterTokenAdminRegistryViaOwner) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewRegisterTokenAdminRegistryViaOwnerInstruction declares a new RegisterTokenAdminRegistryViaOwner instruction with the provided parameters and accounts.
func NewRegisterTokenAdminRegistryViaOwnerInstruction(
	// Accounts:
	tokenAdminRegistry ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *RegisterTokenAdminRegistryViaOwner {
	return NewRegisterTokenAdminRegistryViaOwnerInstructionBuilder().
		SetTokenAdminRegistryAccount(tokenAdminRegistry).
		SetMintAccount(mint).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram)
}
