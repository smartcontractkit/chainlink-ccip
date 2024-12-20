// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_pool

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// DeleteChainConfig is the `deleteChainConfig` instruction.
type DeleteChainConfig struct {
	RemoteChainSelector *uint64
	Mint                *ag_solanago.PublicKey

	// [0] = [] config
	//
	// [1] = [WRITE] chainConfig
	//
	// [2] = [WRITE, SIGNER] authority
	//
	// [3] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewDeleteChainConfigInstructionBuilder creates a new `DeleteChainConfig` instruction builder.
func NewDeleteChainConfigInstructionBuilder() *DeleteChainConfig {
	nd := &DeleteChainConfig{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetRemoteChainSelector sets the "remoteChainSelector" parameter.
func (inst *DeleteChainConfig) SetRemoteChainSelector(remoteChainSelector uint64) *DeleteChainConfig {
	inst.RemoteChainSelector = &remoteChainSelector
	return inst
}

// SetMint sets the "mint" parameter.
func (inst *DeleteChainConfig) SetMint(mint ag_solanago.PublicKey) *DeleteChainConfig {
	inst.Mint = &mint
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *DeleteChainConfig) SetConfigAccount(config ag_solanago.PublicKey) *DeleteChainConfig {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *DeleteChainConfig) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetChainConfigAccount sets the "chainConfig" account.
func (inst *DeleteChainConfig) SetChainConfigAccount(chainConfig ag_solanago.PublicKey) *DeleteChainConfig {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(chainConfig).WRITE()
	return inst
}

// GetChainConfigAccount gets the "chainConfig" account.
func (inst *DeleteChainConfig) GetChainConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *DeleteChainConfig) SetAuthorityAccount(authority ag_solanago.PublicKey) *DeleteChainConfig {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *DeleteChainConfig) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *DeleteChainConfig) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *DeleteChainConfig {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *DeleteChainConfig) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

func (inst DeleteChainConfig) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_DeleteChainConfig,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst DeleteChainConfig) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *DeleteChainConfig) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.RemoteChainSelector == nil {
			return errors.New("RemoteChainSelector parameter is not set")
		}
		if inst.Mint == nil {
			return errors.New("Mint parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ChainConfig is not set")
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

func (inst *DeleteChainConfig) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DeleteChainConfig")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("RemoteChainSelector", *inst.RemoteChainSelector))
						paramsBranch.Child(ag_format.Param("               Mint", *inst.Mint))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       config", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("  chainConfig", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("    authority", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice[3]))
					})
				})
		})
}

func (obj DeleteChainConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `RemoteChainSelector` param:
	err = encoder.Encode(obj.RemoteChainSelector)
	if err != nil {
		return err
	}
	// Serialize `Mint` param:
	err = encoder.Encode(obj.Mint)
	if err != nil {
		return err
	}
	return nil
}
func (obj *DeleteChainConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `RemoteChainSelector`:
	err = decoder.Decode(&obj.RemoteChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `Mint`:
	err = decoder.Decode(&obj.Mint)
	if err != nil {
		return err
	}
	return nil
}

// NewDeleteChainConfigInstruction declares a new DeleteChainConfig instruction with the provided parameters and accounts.
func NewDeleteChainConfigInstruction(
	// Parameters:
	remoteChainSelector uint64,
	mint ag_solanago.PublicKey,
	// Accounts:
	config ag_solanago.PublicKey,
	chainConfig ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *DeleteChainConfig {
	return NewDeleteChainConfigInstructionBuilder().
		SetRemoteChainSelector(remoteChainSelector).
		SetMint(mint).
		SetConfigAccount(config).
		SetChainConfigAccount(chainConfig).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram)
}