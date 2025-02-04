// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ccip_offramp

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Initialization Flow //
// Initializes the CCIP Offramp.
//
// The initialization of the Offramp is responsibility of Admin, nothing more than calling this method should be done first.
//
// # Arguments
//
// * `ctx` - The context containing the accounts required for initialization.
// * `svm_chain_selector` - The chain selector for SVM.
// * `enable_execution_after` - The minimum amount of time required between a message has been committed and can be manually executed.
type Initialize struct {
	SvmChainSelector     *uint64
	EnableExecutionAfter *int64
	Router               *ag_solanago.PublicKey
	FeeQuoter            *ag_solanago.PublicKey
	OfframpLookupTable   *ag_solanago.PublicKey

	// [0] = [WRITE] config
	//
	// [1] = [WRITE] referenceAddresses
	//
	// [2] = [WRITE] state
	//
	// [3] = [WRITE] externalExecutionConfig
	//
	// [4] = [WRITE] tokenPoolsSigner
	//
	// [5] = [WRITE, SIGNER] authority
	//
	// [6] = [] systemProgram
	//
	// [7] = [] program
	//
	// [8] = [] programData
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewInitializeInstructionBuilder creates a new `Initialize` instruction builder.
func NewInitializeInstructionBuilder() *Initialize {
	nd := &Initialize{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 9),
	}
	return nd
}

// SetSvmChainSelector sets the "svmChainSelector" parameter.
func (inst *Initialize) SetSvmChainSelector(svmChainSelector uint64) *Initialize {
	inst.SvmChainSelector = &svmChainSelector
	return inst
}

// SetEnableExecutionAfter sets the "enableExecutionAfter" parameter.
func (inst *Initialize) SetEnableExecutionAfter(enableExecutionAfter int64) *Initialize {
	inst.EnableExecutionAfter = &enableExecutionAfter
	return inst
}

// SetRouter sets the "router" parameter.
func (inst *Initialize) SetRouter(router ag_solanago.PublicKey) *Initialize {
	inst.Router = &router
	return inst
}

// SetFeeQuoter sets the "feeQuoter" parameter.
func (inst *Initialize) SetFeeQuoter(feeQuoter ag_solanago.PublicKey) *Initialize {
	inst.FeeQuoter = &feeQuoter
	return inst
}

// SetOfframpLookupTable sets the "offrampLookupTable" parameter.
func (inst *Initialize) SetOfframpLookupTable(offrampLookupTable ag_solanago.PublicKey) *Initialize {
	inst.OfframpLookupTable = &offrampLookupTable
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *Initialize) SetConfigAccount(config ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config).WRITE()
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *Initialize) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetReferenceAddressesAccount sets the "referenceAddresses" account.
func (inst *Initialize) SetReferenceAddressesAccount(referenceAddresses ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(referenceAddresses).WRITE()
	return inst
}

// GetReferenceAddressesAccount gets the "referenceAddresses" account.
func (inst *Initialize) GetReferenceAddressesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetStateAccount sets the "state" account.
func (inst *Initialize) SetStateAccount(state ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *Initialize) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetExternalExecutionConfigAccount sets the "externalExecutionConfig" account.
func (inst *Initialize) SetExternalExecutionConfigAccount(externalExecutionConfig ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(externalExecutionConfig).WRITE()
	return inst
}

// GetExternalExecutionConfigAccount gets the "externalExecutionConfig" account.
func (inst *Initialize) GetExternalExecutionConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetTokenPoolsSignerAccount sets the "tokenPoolsSigner" account.
func (inst *Initialize) SetTokenPoolsSignerAccount(tokenPoolsSigner ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(tokenPoolsSigner).WRITE()
	return inst
}

// GetTokenPoolsSignerAccount gets the "tokenPoolsSigner" account.
func (inst *Initialize) GetTokenPoolsSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *Initialize) SetAuthorityAccount(authority ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *Initialize) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *Initialize) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *Initialize) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

// SetProgramAccount sets the "program" account.
func (inst *Initialize) SetProgramAccount(program ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *Initialize) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[7]
}

// SetProgramDataAccount sets the "programData" account.
func (inst *Initialize) SetProgramDataAccount(programData ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(programData)
	return inst
}

// GetProgramDataAccount gets the "programData" account.
func (inst *Initialize) GetProgramDataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[8]
}

func (inst Initialize) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Initialize,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Initialize) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Initialize) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.SvmChainSelector == nil {
			return errors.New("SvmChainSelector parameter is not set")
		}
		if inst.EnableExecutionAfter == nil {
			return errors.New("EnableExecutionAfter parameter is not set")
		}
		if inst.Router == nil {
			return errors.New("Router parameter is not set")
		}
		if inst.FeeQuoter == nil {
			return errors.New("FeeQuoter parameter is not set")
		}
		if inst.OfframpLookupTable == nil {
			return errors.New("OfframpLookupTable parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ReferenceAddresses is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.ExternalExecutionConfig is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.TokenPoolsSigner is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Program is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.ProgramData is not set")
		}
	}
	return nil
}

func (inst *Initialize) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Initialize")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=5]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("    SvmChainSelector", *inst.SvmChainSelector))
						paramsBranch.Child(ag_format.Param("EnableExecutionAfter", *inst.EnableExecutionAfter))
						paramsBranch.Child(ag_format.Param("              Router", *inst.Router))
						paramsBranch.Child(ag_format.Param("           FeeQuoter", *inst.FeeQuoter))
						paramsBranch.Child(ag_format.Param("  OfframpLookupTable", *inst.OfframpLookupTable))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=9]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                 config", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("     referenceAddresses", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("                  state", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("externalExecutionConfig", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("       tokenPoolsSigner", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("              authority", inst.AccountMetaSlice[5]))
						accountsBranch.Child(ag_format.Meta("          systemProgram", inst.AccountMetaSlice[6]))
						accountsBranch.Child(ag_format.Meta("                program", inst.AccountMetaSlice[7]))
						accountsBranch.Child(ag_format.Meta("            programData", inst.AccountMetaSlice[8]))
					})
				})
		})
}

func (obj Initialize) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `SvmChainSelector` param:
	err = encoder.Encode(obj.SvmChainSelector)
	if err != nil {
		return err
	}
	// Serialize `EnableExecutionAfter` param:
	err = encoder.Encode(obj.EnableExecutionAfter)
	if err != nil {
		return err
	}
	// Serialize `Router` param:
	err = encoder.Encode(obj.Router)
	if err != nil {
		return err
	}
	// Serialize `FeeQuoter` param:
	err = encoder.Encode(obj.FeeQuoter)
	if err != nil {
		return err
	}
	// Serialize `OfframpLookupTable` param:
	err = encoder.Encode(obj.OfframpLookupTable)
	if err != nil {
		return err
	}
	return nil
}
func (obj *Initialize) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `SvmChainSelector`:
	err = decoder.Decode(&obj.SvmChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `EnableExecutionAfter`:
	err = decoder.Decode(&obj.EnableExecutionAfter)
	if err != nil {
		return err
	}
	// Deserialize `Router`:
	err = decoder.Decode(&obj.Router)
	if err != nil {
		return err
	}
	// Deserialize `FeeQuoter`:
	err = decoder.Decode(&obj.FeeQuoter)
	if err != nil {
		return err
	}
	// Deserialize `OfframpLookupTable`:
	err = decoder.Decode(&obj.OfframpLookupTable)
	if err != nil {
		return err
	}
	return nil
}

// NewInitializeInstruction declares a new Initialize instruction with the provided parameters and accounts.
func NewInitializeInstruction(
	// Parameters:
	svmChainSelector uint64,
	enableExecutionAfter int64,
	router ag_solanago.PublicKey,
	feeQuoter ag_solanago.PublicKey,
	offrampLookupTable ag_solanago.PublicKey,
	// Accounts:
	config ag_solanago.PublicKey,
	referenceAddresses ag_solanago.PublicKey,
	state ag_solanago.PublicKey,
	externalExecutionConfig ag_solanago.PublicKey,
	tokenPoolsSigner ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	program ag_solanago.PublicKey,
	programData ag_solanago.PublicKey) *Initialize {
	return NewInitializeInstructionBuilder().
		SetSvmChainSelector(svmChainSelector).
		SetEnableExecutionAfter(enableExecutionAfter).
		SetRouter(router).
		SetFeeQuoter(feeQuoter).
		SetOfframpLookupTable(offrampLookupTable).
		SetConfigAccount(config).
		SetReferenceAddressesAccount(referenceAddresses).
		SetStateAccount(state).
		SetExternalExecutionConfigAccount(externalExecutionConfig).
		SetTokenPoolsSignerAccount(tokenPoolsSigner).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram).
		SetProgramAccount(program).
		SetProgramDataAccount(programData)
}
