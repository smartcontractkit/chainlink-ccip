// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package cctp_token_pool

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// EditChainRemoteConfigCctp is the `editChainRemoteConfigCctp` instruction.
type EditChainRemoteConfigCctp struct {
	RemoteChainSelector *uint64
	Mint                *ag_solanago.PublicKey
	Cfg                 *CctpChain

	// [0] = [] state
	//
	// [1] = [WRITE] chainConfig
	//
	// [2] = [WRITE, SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewEditChainRemoteConfigCctpInstructionBuilder creates a new `EditChainRemoteConfigCctp` instruction builder.
func NewEditChainRemoteConfigCctpInstructionBuilder() *EditChainRemoteConfigCctp {
	nd := &EditChainRemoteConfigCctp{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetRemoteChainSelector sets the "remoteChainSelector" parameter.
func (inst *EditChainRemoteConfigCctp) SetRemoteChainSelector(remoteChainSelector uint64) *EditChainRemoteConfigCctp {
	inst.RemoteChainSelector = &remoteChainSelector
	return inst
}

// SetMint sets the "mint" parameter.
func (inst *EditChainRemoteConfigCctp) SetMint(mint ag_solanago.PublicKey) *EditChainRemoteConfigCctp {
	inst.Mint = &mint
	return inst
}

// SetCfg sets the "cfg" parameter.
func (inst *EditChainRemoteConfigCctp) SetCfg(cfg CctpChain) *EditChainRemoteConfigCctp {
	inst.Cfg = &cfg
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *EditChainRemoteConfigCctp) SetStateAccount(state ag_solanago.PublicKey) *EditChainRemoteConfigCctp {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *EditChainRemoteConfigCctp) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetChainConfigAccount sets the "chainConfig" account.
func (inst *EditChainRemoteConfigCctp) SetChainConfigAccount(chainConfig ag_solanago.PublicKey) *EditChainRemoteConfigCctp {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(chainConfig).WRITE()
	return inst
}

// GetChainConfigAccount gets the "chainConfig" account.
func (inst *EditChainRemoteConfigCctp) GetChainConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *EditChainRemoteConfigCctp) SetAuthorityAccount(authority ag_solanago.PublicKey) *EditChainRemoteConfigCctp {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *EditChainRemoteConfigCctp) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

func (inst EditChainRemoteConfigCctp) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_EditChainRemoteConfigCctp,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst EditChainRemoteConfigCctp) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *EditChainRemoteConfigCctp) Validate() error {
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
	}
	return nil
}

func (inst *EditChainRemoteConfigCctp) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("EditChainRemoteConfigCctp")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("RemoteChainSelector", *inst.RemoteChainSelector))
						paramsBranch.Child(ag_format.Param("               Mint", *inst.Mint))
						paramsBranch.Child(ag_format.Param("                Cfg", *inst.Cfg))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("      state", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("chainConfig", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("  authority", inst.AccountMetaSlice[2]))
					})
				})
		})
}

func (obj EditChainRemoteConfigCctp) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
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
func (obj *EditChainRemoteConfigCctp) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
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

// NewEditChainRemoteConfigCctpInstruction declares a new EditChainRemoteConfigCctp instruction with the provided parameters and accounts.
func NewEditChainRemoteConfigCctpInstruction(
	// Parameters:
	remoteChainSelector uint64,
	mint ag_solanago.PublicKey,
	cfg CctpChain,
	// Accounts:
	state ag_solanago.PublicKey,
	chainConfig ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *EditChainRemoteConfigCctp {
	return NewEditChainRemoteConfigCctpInstructionBuilder().
		SetRemoteChainSelector(remoteChainSelector).
		SetMint(mint).
		SetCfg(cfg).
		SetStateAccount(state).
		SetChainConfigAccount(chainConfig).
		SetAuthorityAccount(authority)
}
