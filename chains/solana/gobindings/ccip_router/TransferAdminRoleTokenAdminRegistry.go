// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ccip_router

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Transfers the admin role of the token admin registry to a new admin.
//
// Only the Admin can transfer the Admin Role of the Token Admin Registry, this setups the Pending Admin and then it's their responsibility to accept the role.
//
// # Arguments
//
// * `ctx` - The context containing the accounts required for the transfer.
// * `mint` - The public key of the token mint.
// * `new_admin` - The public key of the new admin.
type TransferAdminRoleTokenAdminRegistry struct {
	Mint     *ag_solanago.PublicKey
	NewAdmin *ag_solanago.PublicKey

	// [0] = [WRITE] tokenAdminRegistry
	//
	// [1] = [WRITE, SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewTransferAdminRoleTokenAdminRegistryInstructionBuilder creates a new `TransferAdminRoleTokenAdminRegistry` instruction builder.
func NewTransferAdminRoleTokenAdminRegistryInstructionBuilder() *TransferAdminRoleTokenAdminRegistry {
	nd := &TransferAdminRoleTokenAdminRegistry{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetMint sets the "mint" parameter.
func (inst *TransferAdminRoleTokenAdminRegistry) SetMint(mint ag_solanago.PublicKey) *TransferAdminRoleTokenAdminRegistry {
	inst.Mint = &mint
	return inst
}

// SetNewAdmin sets the "newAdmin" parameter.
func (inst *TransferAdminRoleTokenAdminRegistry) SetNewAdmin(newAdmin ag_solanago.PublicKey) *TransferAdminRoleTokenAdminRegistry {
	inst.NewAdmin = &newAdmin
	return inst
}

// SetTokenAdminRegistryAccount sets the "tokenAdminRegistry" account.
func (inst *TransferAdminRoleTokenAdminRegistry) SetTokenAdminRegistryAccount(tokenAdminRegistry ag_solanago.PublicKey) *TransferAdminRoleTokenAdminRegistry {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(tokenAdminRegistry).WRITE()
	return inst
}

// GetTokenAdminRegistryAccount gets the "tokenAdminRegistry" account.
func (inst *TransferAdminRoleTokenAdminRegistry) GetTokenAdminRegistryAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *TransferAdminRoleTokenAdminRegistry) SetAuthorityAccount(authority ag_solanago.PublicKey) *TransferAdminRoleTokenAdminRegistry {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *TransferAdminRoleTokenAdminRegistry) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

func (inst TransferAdminRoleTokenAdminRegistry) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_TransferAdminRoleTokenAdminRegistry,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst TransferAdminRoleTokenAdminRegistry) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *TransferAdminRoleTokenAdminRegistry) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Mint == nil {
			return errors.New("Mint parameter is not set")
		}
		if inst.NewAdmin == nil {
			return errors.New("NewAdmin parameter is not set")
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

func (inst *TransferAdminRoleTokenAdminRegistry) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("TransferAdminRoleTokenAdminRegistry")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("    Mint", *inst.Mint))
						paramsBranch.Child(ag_format.Param("NewAdmin", *inst.NewAdmin))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("tokenAdminRegistry", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("         authority", inst.AccountMetaSlice[1]))
					})
				})
		})
}

func (obj TransferAdminRoleTokenAdminRegistry) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Mint` param:
	err = encoder.Encode(obj.Mint)
	if err != nil {
		return err
	}
	// Serialize `NewAdmin` param:
	err = encoder.Encode(obj.NewAdmin)
	if err != nil {
		return err
	}
	return nil
}
func (obj *TransferAdminRoleTokenAdminRegistry) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Mint`:
	err = decoder.Decode(&obj.Mint)
	if err != nil {
		return err
	}
	// Deserialize `NewAdmin`:
	err = decoder.Decode(&obj.NewAdmin)
	if err != nil {
		return err
	}
	return nil
}

// NewTransferAdminRoleTokenAdminRegistryInstruction declares a new TransferAdminRoleTokenAdminRegistry instruction with the provided parameters and accounts.
func NewTransferAdminRoleTokenAdminRegistryInstruction(
	// Parameters:
	mint ag_solanago.PublicKey,
	newAdmin ag_solanago.PublicKey,
	// Accounts:
	tokenAdminRegistry ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *TransferAdminRoleTokenAdminRegistry {
	return NewTransferAdminRoleTokenAdminRegistryInstructionBuilder().
		SetMint(mint).
		SetNewAdmin(newAdmin).
		SetTokenAdminRegistryAccount(tokenAdminRegistry).
		SetAuthorityAccount(authority)
}
