// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package mcm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Clear the temporary signature storage.
//
// Closes the account storing signatures, allowing it to be reinitialized if needed.
//
// # Parameters
//
// - `ctx`: The context containing required accounts.
// - `multisig_id`: The multisig instance identifier.
// - `root`: The Merkle root associated with the signatures.
// - `valid_until`: Timestamp until which the root would remain valid.
type ClearSignatures struct {
	MultisigId *[32]uint8
	Root       *[32]uint8
	ValidUntil *uint32

	// [0] = [WRITE] signatures
	//
	// [1] = [WRITE, SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewClearSignaturesInstructionBuilder creates a new `ClearSignatures` instruction builder.
func NewClearSignaturesInstructionBuilder() *ClearSignatures {
	nd := &ClearSignatures{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetMultisigId sets the "multisigId" parameter.
func (inst *ClearSignatures) SetMultisigId(multisigId [32]uint8) *ClearSignatures {
	inst.MultisigId = &multisigId
	return inst
}

// SetRoot sets the "root" parameter.
func (inst *ClearSignatures) SetRoot(root [32]uint8) *ClearSignatures {
	inst.Root = &root
	return inst
}

// SetValidUntil sets the "validUntil" parameter.
func (inst *ClearSignatures) SetValidUntil(validUntil uint32) *ClearSignatures {
	inst.ValidUntil = &validUntil
	return inst
}

// SetSignaturesAccount sets the "signatures" account.
func (inst *ClearSignatures) SetSignaturesAccount(signatures ag_solanago.PublicKey) *ClearSignatures {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(signatures).WRITE()
	return inst
}

// GetSignaturesAccount gets the "signatures" account.
func (inst *ClearSignatures) GetSignaturesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *ClearSignatures) SetAuthorityAccount(authority ag_solanago.PublicKey) *ClearSignatures {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *ClearSignatures) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

func (inst ClearSignatures) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ClearSignatures,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ClearSignatures) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ClearSignatures) Validate() error {
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

func (inst *ClearSignatures) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ClearSignatures")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("MultisigId", *inst.MultisigId))
						paramsBranch.Child(ag_format.Param("      Root", *inst.Root))
						paramsBranch.Child(ag_format.Param("ValidUntil", *inst.ValidUntil))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("signatures", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta(" authority", inst.AccountMetaSlice[1]))
					})
				})
		})
}

func (obj ClearSignatures) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
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
	return nil
}
func (obj *ClearSignatures) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
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
	return nil
}

// NewClearSignaturesInstruction declares a new ClearSignatures instruction with the provided parameters and accounts.
func NewClearSignaturesInstruction(
	// Parameters:
	multisigId [32]uint8,
	root [32]uint8,
	validUntil uint32,
	// Accounts:
	signatures ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *ClearSignatures {
	return NewClearSignaturesInstructionBuilder().
		SetMultisigId(multisigId).
		SetRoot(root).
		SetValidUntil(validUntil).
		SetSignaturesAccount(signatures).
		SetAuthorityAccount(authority)
}
