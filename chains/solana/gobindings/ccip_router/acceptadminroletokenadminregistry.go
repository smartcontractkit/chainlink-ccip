// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ccip_router

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Accepts the admin role of the token admin registry.
//
// The Pending Admin must call this function to accept the admin role of the Token Admin Registry.
//
// # Arguments
//
// * `ctx` - The context containing the accounts required for accepting the admin role.
// * `mint` - The public key of the token mint.
type AcceptAdminRoleTokenAdminRegistry struct {

	// [0] = [] config
	//
	// [1] = [WRITE] token_admin_registry
	//
	// [2] = [] mint
	//
	// [3] = [WRITE, SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewAcceptAdminRoleTokenAdminRegistryInstructionBuilder creates a new `AcceptAdminRoleTokenAdminRegistry` instruction builder.
func NewAcceptAdminRoleTokenAdminRegistryInstructionBuilder() *AcceptAdminRoleTokenAdminRegistry {
	nd := &AcceptAdminRoleTokenAdminRegistry{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetConfigAccount sets the "config" account.
func (inst *AcceptAdminRoleTokenAdminRegistry) SetConfigAccount(config ag_solanago.PublicKey) *AcceptAdminRoleTokenAdminRegistry {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *AcceptAdminRoleTokenAdminRegistry) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetTokenAdminRegistryAccount sets the "token_admin_registry" account.
func (inst *AcceptAdminRoleTokenAdminRegistry) SetTokenAdminRegistryAccount(tokenAdminRegistry ag_solanago.PublicKey) *AcceptAdminRoleTokenAdminRegistry {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(tokenAdminRegistry).WRITE()
	return inst
}

// GetTokenAdminRegistryAccount gets the "token_admin_registry" account.
func (inst *AcceptAdminRoleTokenAdminRegistry) GetTokenAdminRegistryAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMintAccount sets the "mint" account.
func (inst *AcceptAdminRoleTokenAdminRegistry) SetMintAccount(mint ag_solanago.PublicKey) *AcceptAdminRoleTokenAdminRegistry {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *AcceptAdminRoleTokenAdminRegistry) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *AcceptAdminRoleTokenAdminRegistry) SetAuthorityAccount(authority ag_solanago.PublicKey) *AcceptAdminRoleTokenAdminRegistry {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *AcceptAdminRoleTokenAdminRegistry) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst AcceptAdminRoleTokenAdminRegistry) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_AcceptAdminRoleTokenAdminRegistry,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AcceptAdminRoleTokenAdminRegistry) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AcceptAdminRoleTokenAdminRegistry) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.TokenAdminRegistry is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Authority is not set")
		}
	}
	return nil
}

func (inst *AcceptAdminRoleTokenAdminRegistry) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AcceptAdminRoleTokenAdminRegistry")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("              config", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("token_admin_registry", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                mint", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("           authority", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj AcceptAdminRoleTokenAdminRegistry) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *AcceptAdminRoleTokenAdminRegistry) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewAcceptAdminRoleTokenAdminRegistryInstruction declares a new AcceptAdminRoleTokenAdminRegistry instruction with the provided parameters and accounts.
func NewAcceptAdminRoleTokenAdminRegistryInstruction(
	// Accounts:
	config ag_solanago.PublicKey,
	tokenAdminRegistry ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *AcceptAdminRoleTokenAdminRegistry {
	return NewAcceptAdminRoleTokenAdminRegistryInstructionBuilder().
		SetConfigAccount(config).
		SetTokenAdminRegistryAccount(tokenAdminRegistry).
		SetMintAccount(mint).
		SetAuthorityAccount(authority)
}
