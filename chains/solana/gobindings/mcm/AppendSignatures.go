// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package mcm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Append a batch of ECDSA signatures to the temporary storage.
//
// Allows adding multiple signatures in batches to overcome transaction size limits.
//
// # Parameters
//
// - `ctx`: The context containing required accounts.
// - `multisig_id`: The multisig instance identifier.
// - `root`: The Merkle root being approved.
// - `valid_until`: Timestamp until which the root will remain valid.
// - `signatures_batch`: A batch of ECDSA signatures to be verified.
type AppendSignatures struct {
	MultisigId      *[32]uint8
	Root            *[32]uint8
	ValidUntil      *uint32
	SignaturesBatch *[]Signature

	// [0] = [WRITE] signatures
	//
	// [1] = [WRITE, SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewAppendSignaturesInstructionBuilder creates a new `AppendSignatures` instruction builder.
func NewAppendSignaturesInstructionBuilder() *AppendSignatures {
	nd := &AppendSignatures{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetMultisigId sets the "multisigId" parameter.
func (inst *AppendSignatures) SetMultisigId(multisigId [32]uint8) *AppendSignatures {
	inst.MultisigId = &multisigId
	return inst
}

// SetRoot sets the "root" parameter.
func (inst *AppendSignatures) SetRoot(root [32]uint8) *AppendSignatures {
	inst.Root = &root
	return inst
}

// SetValidUntil sets the "validUntil" parameter.
func (inst *AppendSignatures) SetValidUntil(validUntil uint32) *AppendSignatures {
	inst.ValidUntil = &validUntil
	return inst
}

// SetSignaturesBatch sets the "signaturesBatch" parameter.
func (inst *AppendSignatures) SetSignaturesBatch(signaturesBatch []Signature) *AppendSignatures {
	inst.SignaturesBatch = &signaturesBatch
	return inst
}

// SetSignaturesAccount sets the "signatures" account.
func (inst *AppendSignatures) SetSignaturesAccount(signatures ag_solanago.PublicKey) *AppendSignatures {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(signatures).WRITE()
	return inst
}

// GetSignaturesAccount gets the "signatures" account.
func (inst *AppendSignatures) GetSignaturesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *AppendSignatures) SetAuthorityAccount(authority ag_solanago.PublicKey) *AppendSignatures {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *AppendSignatures) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

func (inst AppendSignatures) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_AppendSignatures,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AppendSignatures) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AppendSignatures) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.MultisigId == nil {
			return errors.New("MultisigId parameter is not set")
		}
		if inst.Root == nil {
			return errors.New("Root parameter is not set")
		}
		if inst.ValidUntil == nil {
			return errors.New("ValidUntil parameter is not set")
		}
		if inst.SignaturesBatch == nil {
			return errors.New("SignaturesBatch parameter is not set")
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

func (inst *AppendSignatures) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AppendSignatures")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=4]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("     MultisigId", *inst.MultisigId))
						paramsBranch.Child(ag_format.Param("           Root", *inst.Root))
						paramsBranch.Child(ag_format.Param("     ValidUntil", *inst.ValidUntil))
						paramsBranch.Child(ag_format.Param("SignaturesBatch", *inst.SignaturesBatch))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("signatures", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta(" authority", inst.AccountMetaSlice[1]))
					})
				})
		})
}

func (obj AppendSignatures) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MultisigId` param:
	err = encoder.Encode(obj.MultisigId)
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
	// Serialize `SignaturesBatch` param:
	err = encoder.Encode(obj.SignaturesBatch)
	if err != nil {
		return err
	}
	return nil
}
func (obj *AppendSignatures) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MultisigId`:
	err = decoder.Decode(&obj.MultisigId)
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
	// Deserialize `SignaturesBatch`:
	err = decoder.Decode(&obj.SignaturesBatch)
	if err != nil {
		return err
	}
	return nil
}

// NewAppendSignaturesInstruction declares a new AppendSignatures instruction with the provided parameters and accounts.
func NewAppendSignaturesInstruction(
	// Parameters:
	multisigId [32]uint8,
	root [32]uint8,
	validUntil uint32,
	signaturesBatch []Signature,
	// Accounts:
	signatures ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *AppendSignatures {
	return NewAppendSignaturesInstructionBuilder().
		SetMultisigId(multisigId).
		SetRoot(root).
		SetValidUntil(validUntil).
		SetSignaturesBatch(signaturesBatch).
		SetSignaturesAccount(signatures).
		SetAuthorityAccount(authority)
}
