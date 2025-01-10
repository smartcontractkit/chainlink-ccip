// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package timelock

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// BlockFunctionSelector is the `blockFunctionSelector` instruction.
type BlockFunctionSelector struct {
	TimelockId *[32]uint8
	Selector   *[8]uint8

	// [0] = [WRITE] config
	//
	// [1] = [SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewBlockFunctionSelectorInstructionBuilder creates a new `BlockFunctionSelector` instruction builder.
func NewBlockFunctionSelectorInstructionBuilder() *BlockFunctionSelector {
	nd := &BlockFunctionSelector{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetTimelockId sets the "timelockId" parameter.
func (inst *BlockFunctionSelector) SetTimelockId(timelockId [32]uint8) *BlockFunctionSelector {
	inst.TimelockId = &timelockId
	return inst
}

// SetSelector sets the "selector" parameter.
func (inst *BlockFunctionSelector) SetSelector(selector [8]uint8) *BlockFunctionSelector {
	inst.Selector = &selector
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *BlockFunctionSelector) SetConfigAccount(config ag_solanago.PublicKey) *BlockFunctionSelector {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config).WRITE()
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *BlockFunctionSelector) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *BlockFunctionSelector) SetAuthorityAccount(authority ag_solanago.PublicKey) *BlockFunctionSelector {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *BlockFunctionSelector) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

func (inst BlockFunctionSelector) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_BlockFunctionSelector,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst BlockFunctionSelector) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *BlockFunctionSelector) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.TimelockId == nil {
			return errors.New("TimelockId parameter is not set")
		}
		if inst.Selector == nil {
			return errors.New("Selector parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
	}
	return nil
}

func (inst *BlockFunctionSelector) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("BlockFunctionSelector")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("TimelockId", *inst.TimelockId))
						paramsBranch.Child(ag_format.Param("  Selector", *inst.Selector))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("   config", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("authority", inst.AccountMetaSlice[1]))
					})
				})
		})
}

func (obj BlockFunctionSelector) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `TimelockId` param:
	err = encoder.Encode(obj.TimelockId)
	if err != nil {
		return err
	}
	// Serialize `Selector` param:
	err = encoder.Encode(obj.Selector)
	if err != nil {
		return err
	}
	return nil
}
func (obj *BlockFunctionSelector) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `TimelockId`:
	err = decoder.Decode(&obj.TimelockId)
	if err != nil {
		return err
	}
	// Deserialize `Selector`:
	err = decoder.Decode(&obj.Selector)
	if err != nil {
		return err
	}
	return nil
}

// NewBlockFunctionSelectorInstruction declares a new BlockFunctionSelector instruction with the provided parameters and accounts.
func NewBlockFunctionSelectorInstruction(
	// Parameters:
	timelockId [32]uint8,
	selector [8]uint8,
	// Accounts:
	config ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *BlockFunctionSelector {
	return NewBlockFunctionSelectorInstructionBuilder().
		SetTimelockId(timelockId).
		SetSelector(selector).
		SetConfigAccount(config).
		SetAuthorityAccount(authority)
}
