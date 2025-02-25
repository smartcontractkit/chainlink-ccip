// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package mcm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Append a batch of signer addresses to the temporary storage.
//
// Allows adding multiple signer addresses in batches to overcome transaction size limits.
//
// # Parameters
//
// - `ctx`: The context containing required accounts.
// - `multisig_id`: The multisig instance identifier.
// - `signers_batch`: A batch of Ethereum addresses (20 bytes each) to be added as signers.
type AppendSigners struct {
	MultisigId   *[32]uint8
	SignersBatch *[][20]uint8

	// [0] = [] multisigConfig
	//
	// [1] = [WRITE] configSigners
	//
	// [2] = [WRITE, SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewAppendSignersInstructionBuilder creates a new `AppendSigners` instruction builder.
func NewAppendSignersInstructionBuilder() *AppendSigners {
	nd := &AppendSigners{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetMultisigId sets the "multisigId" parameter.
func (inst *AppendSigners) SetMultisigId(multisigId [32]uint8) *AppendSigners {
	inst.MultisigId = &multisigId
	return inst
}

// SetSignersBatch sets the "signersBatch" parameter.
func (inst *AppendSigners) SetSignersBatch(signersBatch [][20]uint8) *AppendSigners {
	inst.SignersBatch = &signersBatch
	return inst
}

// SetMultisigConfigAccount sets the "multisigConfig" account.
func (inst *AppendSigners) SetMultisigConfigAccount(multisigConfig ag_solanago.PublicKey) *AppendSigners {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(multisigConfig)
	return inst
}

// GetMultisigConfigAccount gets the "multisigConfig" account.
func (inst *AppendSigners) GetMultisigConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetConfigSignersAccount sets the "configSigners" account.
func (inst *AppendSigners) SetConfigSignersAccount(configSigners ag_solanago.PublicKey) *AppendSigners {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(configSigners).WRITE()
	return inst
}

// GetConfigSignersAccount gets the "configSigners" account.
func (inst *AppendSigners) GetConfigSignersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *AppendSigners) SetAuthorityAccount(authority ag_solanago.PublicKey) *AppendSigners {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *AppendSigners) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

func (inst AppendSigners) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_AppendSigners,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AppendSigners) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AppendSigners) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.MultisigId == nil {
			return errors.New("MultisigId parameter is not set")
		}
		if inst.SignersBatch == nil {
			return errors.New("SignersBatch parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.MultisigConfig is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ConfigSigners is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Authority is not set")
		}
	}
	return nil
}

func (inst *AppendSigners) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AppendSigners")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("  MultisigId", *inst.MultisigId))
						paramsBranch.Child(ag_format.Param("SignersBatch", *inst.SignersBatch))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("multisigConfig", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta(" configSigners", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("     authority", inst.AccountMetaSlice[2]))
					})
				})
		})
}

func (obj AppendSigners) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MultisigId` param:
	err = encoder.Encode(obj.MultisigId)
	if err != nil {
		return err
	}
	// Serialize `SignersBatch` param:
	err = encoder.Encode(obj.SignersBatch)
	if err != nil {
		return err
	}
	return nil
}
func (obj *AppendSigners) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MultisigId`:
	err = decoder.Decode(&obj.MultisigId)
	if err != nil {
		return err
	}
	// Deserialize `SignersBatch`:
	err = decoder.Decode(&obj.SignersBatch)
	if err != nil {
		return err
	}
	return nil
}

// NewAppendSignersInstruction declares a new AppendSigners instruction with the provided parameters and accounts.
func NewAppendSignersInstruction(
	// Parameters:
	multisigId [32]uint8,
	signersBatch [][20]uint8,
	// Accounts:
	multisigConfig ag_solanago.PublicKey,
	configSigners ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *AppendSigners {
	return NewAppendSignersInstructionBuilder().
		SetMultisigId(multisigId).
		SetSignersBatch(signersBatch).
		SetMultisigConfigAccount(multisigConfig).
		SetConfigSignersAccount(configSigners).
		SetAuthorityAccount(authority)
}
