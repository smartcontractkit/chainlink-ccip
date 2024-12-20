// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ccip_router

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Adds a billing token configuration.
// Only CCIP Admin can add a billing token configuration.
//
// # Arguments
//
// * `ctx` - The context containing the accounts required for adding the billing token configuration.
// * `config` - The billing token configuration to be added.
type AddBillingTokenConfig struct {
	Config *BillingTokenConfig

	// [0] = [] config
	//
	// [1] = [WRITE] billingTokenConfig
	//
	// [2] = [] tokenProgram
	// ··········· type of a specific program (which would enforce its ID). Thus, it's an UncheckedAccount
	// ··········· with a constraint enforcing that it is one of the two allowed programs.
	//
	// [3] = [] feeTokenMint
	//
	// [4] = [WRITE] feeTokenReceiver
	//
	// [5] = [WRITE, SIGNER] authority
	//
	// [6] = [] feeBillingSigner
	//
	// [7] = [] associatedTokenProgram
	//
	// [8] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewAddBillingTokenConfigInstructionBuilder creates a new `AddBillingTokenConfig` instruction builder.
func NewAddBillingTokenConfigInstructionBuilder() *AddBillingTokenConfig {
	nd := &AddBillingTokenConfig{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 9),
	}
	return nd
}

// SetConfig sets the "config" parameter.
func (inst *AddBillingTokenConfig) SetConfig(config BillingTokenConfig) *AddBillingTokenConfig {
	inst.Config = &config
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *AddBillingTokenConfig) SetConfigAccount(config ag_solanago.PublicKey) *AddBillingTokenConfig {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *AddBillingTokenConfig) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetBillingTokenConfigAccount sets the "billingTokenConfig" account.
func (inst *AddBillingTokenConfig) SetBillingTokenConfigAccount(billingTokenConfig ag_solanago.PublicKey) *AddBillingTokenConfig {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(billingTokenConfig).WRITE()
	return inst
}

// GetBillingTokenConfigAccount gets the "billingTokenConfig" account.
func (inst *AddBillingTokenConfig) GetBillingTokenConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetTokenProgramAccount sets the "tokenProgram" account.
// type of a specific program (which would enforce its ID). Thus, it's an UncheckedAccount
// with a constraint enforcing that it is one of the two allowed programs.
func (inst *AddBillingTokenConfig) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *AddBillingTokenConfig {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
// type of a specific program (which would enforce its ID). Thus, it's an UncheckedAccount
// with a constraint enforcing that it is one of the two allowed programs.
func (inst *AddBillingTokenConfig) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetFeeTokenMintAccount sets the "feeTokenMint" account.
func (inst *AddBillingTokenConfig) SetFeeTokenMintAccount(feeTokenMint ag_solanago.PublicKey) *AddBillingTokenConfig {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(feeTokenMint)
	return inst
}

// GetFeeTokenMintAccount gets the "feeTokenMint" account.
func (inst *AddBillingTokenConfig) GetFeeTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetFeeTokenReceiverAccount sets the "feeTokenReceiver" account.
func (inst *AddBillingTokenConfig) SetFeeTokenReceiverAccount(feeTokenReceiver ag_solanago.PublicKey) *AddBillingTokenConfig {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(feeTokenReceiver).WRITE()
	return inst
}

// GetFeeTokenReceiverAccount gets the "feeTokenReceiver" account.
func (inst *AddBillingTokenConfig) GetFeeTokenReceiverAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *AddBillingTokenConfig) SetAuthorityAccount(authority ag_solanago.PublicKey) *AddBillingTokenConfig {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *AddBillingTokenConfig) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetFeeBillingSignerAccount sets the "feeBillingSigner" account.
func (inst *AddBillingTokenConfig) SetFeeBillingSignerAccount(feeBillingSigner ag_solanago.PublicKey) *AddBillingTokenConfig {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(feeBillingSigner)
	return inst
}

// GetFeeBillingSignerAccount gets the "feeBillingSigner" account.
func (inst *AddBillingTokenConfig) GetFeeBillingSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

// SetAssociatedTokenProgramAccount sets the "associatedTokenProgram" account.
func (inst *AddBillingTokenConfig) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *AddBillingTokenConfig {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associatedTokenProgram" account.
func (inst *AddBillingTokenConfig) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[7]
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *AddBillingTokenConfig) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *AddBillingTokenConfig {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *AddBillingTokenConfig) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[8]
}

func (inst AddBillingTokenConfig) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_AddBillingTokenConfig,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AddBillingTokenConfig) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AddBillingTokenConfig) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Config == nil {
			return errors.New("Config parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.BillingTokenConfig is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.FeeTokenMint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.FeeTokenReceiver is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.FeeBillingSigner is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *AddBillingTokenConfig) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AddBillingTokenConfig")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Config", *inst.Config))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=9]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                config", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("    billingTokenConfig", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("          tokenProgram", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("          feeTokenMint", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("      feeTokenReceiver", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("             authority", inst.AccountMetaSlice[5]))
						accountsBranch.Child(ag_format.Meta("      feeBillingSigner", inst.AccountMetaSlice[6]))
						accountsBranch.Child(ag_format.Meta("associatedTokenProgram", inst.AccountMetaSlice[7]))
						accountsBranch.Child(ag_format.Meta("         systemProgram", inst.AccountMetaSlice[8]))
					})
				})
		})
}

func (obj AddBillingTokenConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Config` param:
	err = encoder.Encode(obj.Config)
	if err != nil {
		return err
	}
	return nil
}
func (obj *AddBillingTokenConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Config`:
	err = decoder.Decode(&obj.Config)
	if err != nil {
		return err
	}
	return nil
}

// NewAddBillingTokenConfigInstruction declares a new AddBillingTokenConfig instruction with the provided parameters and accounts.
func NewAddBillingTokenConfigInstruction(
	// Parameters:
	config BillingTokenConfig,
	// Accounts:
	configAccount ag_solanago.PublicKey,
	billingTokenConfig ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	feeTokenMint ag_solanago.PublicKey,
	feeTokenReceiver ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	feeBillingSigner ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *AddBillingTokenConfig {
	return NewAddBillingTokenConfigInstructionBuilder().
		SetConfig(config).
		SetConfigAccount(configAccount).
		SetBillingTokenConfigAccount(billingTokenConfig).
		SetTokenProgramAccount(tokenProgram).
		SetFeeTokenMintAccount(feeTokenMint).
		SetFeeTokenReceiverAccount(feeTokenReceiver).
		SetAuthorityAccount(authority).
		SetFeeBillingSignerAccount(feeBillingSigner).
		SetAssociatedTokenProgramAccount(associatedTokenProgram).
		SetSystemProgramAccount(systemProgram)
}