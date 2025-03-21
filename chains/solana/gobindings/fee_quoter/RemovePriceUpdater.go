// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package fee_quoter

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Remove a price updater address from the list of allowed price updaters.
//
// # Arguments
//
// * `ctx` - The context containing the accounts required for this operation.
// * `price_updater` - The price updater address.
type RemovePriceUpdater struct {
	PriceUpdater *ag_solanago.PublicKey

	// [0] = [WRITE] allowedPriceUpdater
	//
	// [1] = [] config
	//
	// [2] = [WRITE, SIGNER] authority
	//
	// [3] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewRemovePriceUpdaterInstructionBuilder creates a new `RemovePriceUpdater` instruction builder.
func NewRemovePriceUpdaterInstructionBuilder() *RemovePriceUpdater {
	nd := &RemovePriceUpdater{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetPriceUpdater sets the "priceUpdater" parameter.
func (inst *RemovePriceUpdater) SetPriceUpdater(priceUpdater ag_solanago.PublicKey) *RemovePriceUpdater {
	inst.PriceUpdater = &priceUpdater
	return inst
}

// SetAllowedPriceUpdaterAccount sets the "allowedPriceUpdater" account.
func (inst *RemovePriceUpdater) SetAllowedPriceUpdaterAccount(allowedPriceUpdater ag_solanago.PublicKey) *RemovePriceUpdater {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(allowedPriceUpdater).WRITE()
	return inst
}

// GetAllowedPriceUpdaterAccount gets the "allowedPriceUpdater" account.
func (inst *RemovePriceUpdater) GetAllowedPriceUpdaterAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetConfigAccount sets the "config" account.
func (inst *RemovePriceUpdater) SetConfigAccount(config ag_solanago.PublicKey) *RemovePriceUpdater {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *RemovePriceUpdater) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *RemovePriceUpdater) SetAuthorityAccount(authority ag_solanago.PublicKey) *RemovePriceUpdater {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *RemovePriceUpdater) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *RemovePriceUpdater) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *RemovePriceUpdater {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *RemovePriceUpdater) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

func (inst RemovePriceUpdater) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_RemovePriceUpdater,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst RemovePriceUpdater) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *RemovePriceUpdater) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.PriceUpdater == nil {
			return errors.New("PriceUpdater parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.AllowedPriceUpdater is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *RemovePriceUpdater) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("RemovePriceUpdater")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("PriceUpdater", *inst.PriceUpdater))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("allowedPriceUpdater", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("             config", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("          authority", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("      systemProgram", inst.AccountMetaSlice[3]))
					})
				})
		})
}

func (obj RemovePriceUpdater) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `PriceUpdater` param:
	err = encoder.Encode(obj.PriceUpdater)
	if err != nil {
		return err
	}
	return nil
}
func (obj *RemovePriceUpdater) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `PriceUpdater`:
	err = decoder.Decode(&obj.PriceUpdater)
	if err != nil {
		return err
	}
	return nil
}

// NewRemovePriceUpdaterInstruction declares a new RemovePriceUpdater instruction with the provided parameters and accounts.
func NewRemovePriceUpdaterInstruction(
	// Parameters:
	priceUpdater ag_solanago.PublicKey,
	// Accounts:
	allowedPriceUpdater ag_solanago.PublicKey,
	config ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *RemovePriceUpdater {
	return NewRemovePriceUpdaterInstructionBuilder().
		SetPriceUpdater(priceUpdater).
		SetAllowedPriceUpdaterAccount(allowedPriceUpdater).
		SetConfigAccount(config).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram)
}
