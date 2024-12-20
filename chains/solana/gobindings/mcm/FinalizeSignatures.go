// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package mcm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// FinalizeSignatures is the `finalizeSignatures` instruction.
type FinalizeSignatures struct {
	MultisigName *[32]uint8
	Root         *[32]uint8
	ValidUntil   *uint32

	// [0] = [WRITE] signatures
	//
	// [1] = [WRITE, SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewFinalizeSignaturesInstructionBuilder creates a new `FinalizeSignatures` instruction builder.
func NewFinalizeSignaturesInstructionBuilder() *FinalizeSignatures {
	nd := &FinalizeSignatures{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetMultisigName sets the "multisigName" parameter.
func (inst *FinalizeSignatures) SetMultisigName(multisigName [32]uint8) *FinalizeSignatures {
	inst.MultisigName = &multisigName
	return inst
}

// SetRoot sets the "root" parameter.
func (inst *FinalizeSignatures) SetRoot(root [32]uint8) *FinalizeSignatures {
	inst.Root = &root
	return inst
}

// SetValidUntil sets the "validUntil" parameter.
func (inst *FinalizeSignatures) SetValidUntil(validUntil uint32) *FinalizeSignatures {
	inst.ValidUntil = &validUntil
	return inst
}

// SetSignaturesAccount sets the "signatures" account.
func (inst *FinalizeSignatures) SetSignaturesAccount(signatures ag_solanago.PublicKey) *FinalizeSignatures {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(signatures).WRITE()
	return inst
}

// GetSignaturesAccount gets the "signatures" account.
func (inst *FinalizeSignatures) GetSignaturesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *FinalizeSignatures) SetAuthorityAccount(authority ag_solanago.PublicKey) *FinalizeSignatures {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *FinalizeSignatures) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

func (inst FinalizeSignatures) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_FinalizeSignatures,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst FinalizeSignatures) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *FinalizeSignatures) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.MultisigName == nil {
			return errors.New("MultisigName parameter is not set")
		}
		if inst.Root == nil {
			return errors.New("Root parameter is not set")
		}
		if inst.ValidUntil == nil {
			return errors.New("ValidUntil parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Signatures is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
	}
	return nil
}

func (inst *FinalizeSignatures) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("FinalizeSignatures")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("MultisigName", *inst.MultisigName))
						paramsBranch.Child(ag_format.Param("        Root", *inst.Root))
						paramsBranch.Child(ag_format.Param("  ValidUntil", *inst.ValidUntil))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("signatures", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta(" authority", inst.AccountMetaSlice[1]))
					})
				})
		})
}

func (obj FinalizeSignatures) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MultisigName` param:
	err = encoder.Encode(obj.MultisigName)
	if err != nil {
		return err
	}
	// Serialize `Root` param:
	err = encoder.Encode(obj.Root)
	if err != nil {
		return err
	}
	// Serialize `ValidUntil` param:
	err = encoder.Encode(obj.ValidUntil)
	if err != nil {
		return err
	}
	return nil
}
func (obj *FinalizeSignatures) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MultisigName`:
	err = decoder.Decode(&obj.MultisigName)
	if err != nil {
		return err
	}
	// Deserialize `Root`:
	err = decoder.Decode(&obj.Root)
	if err != nil {
		return err
	}
	// Deserialize `ValidUntil`:
	err = decoder.Decode(&obj.ValidUntil)
	if err != nil {
		return err
	}
	return nil
}

// NewFinalizeSignaturesInstruction declares a new FinalizeSignatures instruction with the provided parameters and accounts.
func NewFinalizeSignaturesInstruction(
	// Parameters:
	multisigName [32]uint8,
	root [32]uint8,
	validUntil uint32,
	// Accounts:
	signatures ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *FinalizeSignatures {
	return NewFinalizeSignaturesInstructionBuilder().
		SetMultisigName(multisigName).
		SetRoot(root).
		SetValidUntil(validUntil).
		SetSignaturesAccount(signatures).
		SetAuthorityAccount(authority)
}