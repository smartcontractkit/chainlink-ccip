// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ccip_offramp

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Updates reference addresses in the offramp contract, such as
// the CCIP router, Fee Quoter, and the Offramp Lookup Table.
// Only the Admin may update these addresses.
//
// # Arguments
//
// * `ctx` - The context containing the accounts required for updating the reference addresses.
// * `reference_addresses` - The new reference addresses to be set.
type UpdateReferenceAddresses struct {
	ReferenceAddresses *ReferenceAddresses

	// [0] = [] config
	//
	// [1] = [WRITE] referenceAddresses
	//
	// [2] = [SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewUpdateReferenceAddressesInstructionBuilder creates a new `UpdateReferenceAddresses` instruction builder.
func NewUpdateReferenceAddressesInstructionBuilder() *UpdateReferenceAddresses {
	nd := &UpdateReferenceAddresses{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetReferenceAddresses sets the "referenceAddresses" parameter.
func (inst *UpdateReferenceAddresses) SetReferenceAddresses(referenceAddresses ReferenceAddresses) *UpdateReferenceAddresses {
	inst.ReferenceAddresses = &referenceAddresses
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *UpdateReferenceAddresses) SetConfigAccount(config ag_solanago.PublicKey) *UpdateReferenceAddresses {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *UpdateReferenceAddresses) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetReferenceAddressesAccount sets the "referenceAddresses" account.
func (inst *UpdateReferenceAddresses) SetReferenceAddressesAccount(referenceAddresses ag_solanago.PublicKey) *UpdateReferenceAddresses {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(referenceAddresses).WRITE()
	return inst
}

// GetReferenceAddressesAccount gets the "referenceAddresses" account.
func (inst *UpdateReferenceAddresses) GetReferenceAddressesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *UpdateReferenceAddresses) SetAuthorityAccount(authority ag_solanago.PublicKey) *UpdateReferenceAddresses {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *UpdateReferenceAddresses) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

func (inst UpdateReferenceAddresses) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateReferenceAddresses,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateReferenceAddresses) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateReferenceAddresses) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.ReferenceAddresses == nil {
			return errors.New("ReferenceAddresses parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ReferenceAddresses is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Authority is not set")
		}
	}
	return nil
}

func (inst *UpdateReferenceAddresses) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateReferenceAddresses")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("ReferenceAddresses", *inst.ReferenceAddresses))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("            config", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("referenceAddresses", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("         authority", inst.AccountMetaSlice[2]))
					})
				})
		})
}

func (obj UpdateReferenceAddresses) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ReferenceAddresses` param:
	err = encoder.Encode(obj.ReferenceAddresses)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdateReferenceAddresses) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ReferenceAddresses`:
	err = decoder.Decode(&obj.ReferenceAddresses)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateReferenceAddressesInstruction declares a new UpdateReferenceAddresses instruction with the provided parameters and accounts.
func NewUpdateReferenceAddressesInstruction(
	// Parameters:
	referenceAddresses ReferenceAddresses,
	// Accounts:
	config ag_solanago.PublicKey,
	referenceAddressesAccount ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *UpdateReferenceAddresses {
	return NewUpdateReferenceAddressesInstructionBuilder().
		SetReferenceAddresses(referenceAddresses).
		SetConfigAccount(config).
		SetReferenceAddressesAccount(referenceAddressesAccount).
		SetAuthorityAccount(authority)
}
