// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package example_burnmint_token_pool

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// InitChainRemoteConfig is the `initChainRemoteConfig` instruction.
type InitChainRemoteConfig struct {
	RemoteChainSelector *uint64
	Mint                *ag_solanago.PublicKey
	Cfg                 *RemoteConfig

	// [0] = [] state
	//
	// [1] = [WRITE] chainConfig
	//
	// [2] = [WRITE, SIGNER] authority
	//
	// [3] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewInitChainRemoteConfigInstructionBuilder creates a new `InitChainRemoteConfig` instruction builder.
func NewInitChainRemoteConfigInstructionBuilder() *InitChainRemoteConfig {
	nd := &InitChainRemoteConfig{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetRemoteChainSelector sets the "remoteChainSelector" parameter.
func (inst *InitChainRemoteConfig) SetRemoteChainSelector(remoteChainSelector uint64) *InitChainRemoteConfig {
	inst.RemoteChainSelector = &remoteChainSelector
	return inst
}

// SetMint sets the "mint" parameter.
func (inst *InitChainRemoteConfig) SetMint(mint ag_solanago.PublicKey) *InitChainRemoteConfig {
	inst.Mint = &mint
	return inst
}

// SetCfg sets the "cfg" parameter.
func (inst *InitChainRemoteConfig) SetCfg(cfg RemoteConfig) *InitChainRemoteConfig {
	inst.Cfg = &cfg
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *InitChainRemoteConfig) SetStateAccount(state ag_solanago.PublicKey) *InitChainRemoteConfig {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *InitChainRemoteConfig) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetChainConfigAccount sets the "chainConfig" account.
func (inst *InitChainRemoteConfig) SetChainConfigAccount(chainConfig ag_solanago.PublicKey) *InitChainRemoteConfig {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(chainConfig).WRITE()
	return inst
}

// GetChainConfigAccount gets the "chainConfig" account.
func (inst *InitChainRemoteConfig) GetChainConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *InitChainRemoteConfig) SetAuthorityAccount(authority ag_solanago.PublicKey) *InitChainRemoteConfig {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *InitChainRemoteConfig) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *InitChainRemoteConfig) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *InitChainRemoteConfig {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *InitChainRemoteConfig) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

func (inst InitChainRemoteConfig) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitChainRemoteConfig,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitChainRemoteConfig) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitChainRemoteConfig) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.RemoteChainSelector == nil {
			return errors.New("RemoteChainSelector parameter is not set")
		}
		if inst.Mint == nil {
			return errors.New("Mint parameter is not set")
		}
		if inst.Cfg == nil {
			return errors.New("Cfg parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
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

func (inst *InitChainRemoteConfig) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitChainRemoteConfig")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("RemoteChainSelector", *inst.RemoteChainSelector))
						paramsBranch.Child(ag_format.Param("               Mint", *inst.Mint))
						paramsBranch.Child(ag_format.Param("                Cfg", *inst.Cfg))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        state", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("  chainConfig", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("    authority", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice[3]))
					})
				})
		})
}

func (obj InitChainRemoteConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
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
	// Serialize `Cfg` param:
	err = encoder.Encode(obj.Cfg)
	if err != nil {
		return err
	}
	return nil
}
func (obj *InitChainRemoteConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
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
	// Deserialize `Cfg`:
	err = decoder.Decode(&obj.Cfg)
	if err != nil {
		return err
	}
	return nil
}

// NewInitChainRemoteConfigInstruction declares a new InitChainRemoteConfig instruction with the provided parameters and accounts.
func NewInitChainRemoteConfigInstruction(
	// Parameters:
	remoteChainSelector uint64,
	mint ag_solanago.PublicKey,
	cfg RemoteConfig,
	// Accounts:
	state ag_solanago.PublicKey,
	chainConfig ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *InitChainRemoteConfig {
	return NewInitChainRemoteConfigInstructionBuilder().
		SetRemoteChainSelector(remoteChainSelector).
		SetMint(mint).
		SetCfg(cfg).
		SetStateAccount(state).
		SetChainConfigAccount(chainConfig).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram)
}
