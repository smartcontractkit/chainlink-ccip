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
	Mint *ag_solanago.PublicKey

	// [0] = [WRITE] tokenAdminRegistry
	//
	// [1] = [WRITE, SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewAcceptAdminRoleTokenAdminRegistryInstructionBuilder creates a new `AcceptAdminRoleTokenAdminRegistry` instruction builder.
func NewAcceptAdminRoleTokenAdminRegistryInstructionBuilder() *AcceptAdminRoleTokenAdminRegistry {
	nd := &AcceptAdminRoleTokenAdminRegistry{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetMint sets the "mint" parameter.
func (inst *AcceptAdminRoleTokenAdminRegistry) SetMint(mint ag_solanago.PublicKey) *AcceptAdminRoleTokenAdminRegistry {
	inst.Mint = &mint
	return inst
}

// SetTokenAdminRegistryAccount sets the "tokenAdminRegistry" account.
func (inst *AcceptAdminRoleTokenAdminRegistry) SetTokenAdminRegistryAccount(tokenAdminRegistry ag_solanago.PublicKey) *AcceptAdminRoleTokenAdminRegistry {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(tokenAdminRegistry).WRITE()
	return inst
}

// GetTokenAdminRegistryAccount gets the "tokenAdminRegistry" account.
func (inst *AcceptAdminRoleTokenAdminRegistry) GetTokenAdminRegistryAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *AcceptAdminRoleTokenAdminRegistry) SetAuthorityAccount(authority ag_solanago.PublicKey) *AcceptAdminRoleTokenAdminRegistry {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *AcceptAdminRoleTokenAdminRegistry) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
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
	// Check whether all (required) parameters are set:
	{
		if inst.Mint == nil {
			return errors.New("Mint parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.TokenAdminRegistry is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
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
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Mint", *inst.Mint))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("tokenAdminRegistry", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("         authority", inst.AccountMetaSlice[1]))
					})
				})
		})
}

func (obj AcceptAdminRoleTokenAdminRegistry) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Mint` param:
	err = encoder.Encode(obj.Mint)
	if err != nil {
		return err
	}
	return nil
}
func (obj *AcceptAdminRoleTokenAdminRegistry) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Mint`:
	err = decoder.Decode(&obj.Mint)
	if err != nil {
		return err
	}
	return nil
}

// NewAcceptAdminRoleTokenAdminRegistryInstruction declares a new AcceptAdminRoleTokenAdminRegistry instruction with the provided parameters and accounts.
func NewAcceptAdminRoleTokenAdminRegistryInstruction(
	// Parameters:
	mint ag_solanago.PublicKey,
	// Accounts:
	tokenAdminRegistry ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *AcceptAdminRoleTokenAdminRegistry {
	return NewAcceptAdminRoleTokenAdminRegistryInstructionBuilder().
		SetMint(mint).
		SetTokenAdminRegistryAccount(tokenAdminRegistry).
		SetAuthorityAccount(authority)
}
