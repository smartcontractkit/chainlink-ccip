// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package mcm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SetRoot is the `setRoot` instruction.
type SetRoot struct {
	MultisigName  *[32]uint8
	Root          *[32]uint8
	ValidUntil    *uint32
	Metadata      *RootMetadataInput
	MetadataProof *[][32]uint8

	// [0] = [WRITE] rootSignatures
	//
	// [1] = [WRITE] rootMetadata
	//
	// [2] = [WRITE] seenSignedHashes
	//
	// [3] = [WRITE] expiringRootAndOpCount
	//
	// [4] = [] multisigConfig
	//
	// [5] = [WRITE, SIGNER] authority
	//
	// [6] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewSetRootInstructionBuilder creates a new `SetRoot` instruction builder.
func NewSetRootInstructionBuilder() *SetRoot {
	nd := &SetRoot{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	return nd
}

// SetMultisigName sets the "multisigName" parameter.
func (inst *SetRoot) SetMultisigName(multisigName [32]uint8) *SetRoot {
	inst.MultisigName = &multisigName
	return inst
}

// SetRoot sets the "root" parameter.
func (inst *SetRoot) SetRoot(root [32]uint8) *SetRoot {
	inst.Root = &root
	return inst
}

// SetValidUntil sets the "validUntil" parameter.
func (inst *SetRoot) SetValidUntil(validUntil uint32) *SetRoot {
	inst.ValidUntil = &validUntil
	return inst
}

// SetMetadata sets the "metadata" parameter.
func (inst *SetRoot) SetMetadata(metadata RootMetadataInput) *SetRoot {
	inst.Metadata = &metadata
	return inst
}

// SetMetadataProof sets the "metadataProof" parameter.
func (inst *SetRoot) SetMetadataProof(metadataProof [][32]uint8) *SetRoot {
	inst.MetadataProof = &metadataProof
	return inst
}

// SetRootSignaturesAccount sets the "rootSignatures" account.
func (inst *SetRoot) SetRootSignaturesAccount(rootSignatures ag_solanago.PublicKey) *SetRoot {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(rootSignatures).WRITE()
	return inst
}

// GetRootSignaturesAccount gets the "rootSignatures" account.
func (inst *SetRoot) GetRootSignaturesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetRootMetadataAccount sets the "rootMetadata" account.
func (inst *SetRoot) SetRootMetadataAccount(rootMetadata ag_solanago.PublicKey) *SetRoot {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(rootMetadata).WRITE()
	return inst
}

// GetRootMetadataAccount gets the "rootMetadata" account.
func (inst *SetRoot) GetRootMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetSeenSignedHashesAccount sets the "seenSignedHashes" account.
func (inst *SetRoot) SetSeenSignedHashesAccount(seenSignedHashes ag_solanago.PublicKey) *SetRoot {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(seenSignedHashes).WRITE()
	return inst
}

// GetSeenSignedHashesAccount gets the "seenSignedHashes" account.
func (inst *SetRoot) GetSeenSignedHashesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetExpiringRootAndOpCountAccount sets the "expiringRootAndOpCount" account.
func (inst *SetRoot) SetExpiringRootAndOpCountAccount(expiringRootAndOpCount ag_solanago.PublicKey) *SetRoot {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(expiringRootAndOpCount).WRITE()
	return inst
}

// GetExpiringRootAndOpCountAccount gets the "expiringRootAndOpCount" account.
func (inst *SetRoot) GetExpiringRootAndOpCountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetMultisigConfigAccount sets the "multisigConfig" account.
func (inst *SetRoot) SetMultisigConfigAccount(multisigConfig ag_solanago.PublicKey) *SetRoot {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(multisigConfig)
	return inst
}

// GetMultisigConfigAccount gets the "multisigConfig" account.
func (inst *SetRoot) GetMultisigConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *SetRoot) SetAuthorityAccount(authority ag_solanago.PublicKey) *SetRoot {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *SetRoot) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *SetRoot) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *SetRoot {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *SetRoot) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

func (inst SetRoot) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SetRoot,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetRoot) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetRoot) Validate() error {
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
		if inst.Metadata == nil {
			return errors.New("Metadata parameter is not set")
		}
		if inst.MetadataProof == nil {
			return errors.New("MetadataProof parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.RootSignatures is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.RootMetadata is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SeenSignedHashes is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.ExpiringRootAndOpCount is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.MultisigConfig is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *SetRoot) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetRoot")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=5]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param(" MultisigName", *inst.MultisigName))
						paramsBranch.Child(ag_format.Param("         Root", *inst.Root))
						paramsBranch.Child(ag_format.Param("   ValidUntil", *inst.ValidUntil))
						paramsBranch.Child(ag_format.Param("     Metadata", *inst.Metadata))
						paramsBranch.Child(ag_format.Param("MetadataProof", *inst.MetadataProof))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        rootSignatures", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("          rootMetadata", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("      seenSignedHashes", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("expiringRootAndOpCount", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("        multisigConfig", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("             authority", inst.AccountMetaSlice[5]))
						accountsBranch.Child(ag_format.Meta("         systemProgram", inst.AccountMetaSlice[6]))
					})
				})
		})
}

func (obj SetRoot) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
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
	// Serialize `Metadata` param:
	err = encoder.Encode(obj.Metadata)
	if err != nil {
		return err
	}
	// Serialize `MetadataProof` param:
	err = encoder.Encode(obj.MetadataProof)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SetRoot) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
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
	// Deserialize `Metadata`:
	err = decoder.Decode(&obj.Metadata)
	if err != nil {
		return err
	}
	// Deserialize `MetadataProof`:
	err = decoder.Decode(&obj.MetadataProof)
	if err != nil {
		return err
	}
	return nil
}

// NewSetRootInstruction declares a new SetRoot instruction with the provided parameters and accounts.
func NewSetRootInstruction(
	// Parameters:
	multisigName [32]uint8,
	root [32]uint8,
	validUntil uint32,
	metadata RootMetadataInput,
	metadataProof [][32]uint8,
	// Accounts:
	rootSignatures ag_solanago.PublicKey,
	rootMetadata ag_solanago.PublicKey,
	seenSignedHashes ag_solanago.PublicKey,
	expiringRootAndOpCount ag_solanago.PublicKey,
	multisigConfig ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *SetRoot {
	return NewSetRootInstructionBuilder().
		SetMultisigName(multisigName).
		SetRoot(root).
		SetValidUntil(validUntil).
		SetMetadata(metadata).
		SetMetadataProof(metadataProof).
		SetRootSignaturesAccount(rootSignatures).
		SetRootMetadataAccount(rootMetadata).
		SetSeenSignedHashesAccount(seenSignedHashes).
		SetExpiringRootAndOpCountAccount(expiringRootAndOpCount).
		SetMultisigConfigAccount(multisigConfig).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram)
}
