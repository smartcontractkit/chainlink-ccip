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
// Initializes the CCIP Offramp, except for the config account (due to stack size limitations).
//
// The initialization of the Offramp is responsibility of Admin, nothing more than calling these
// initialization methods should be done first.
//
// # Arguments
//
// * `ctx` - The context containing the accounts required for initialization.
type Initialize struct {

	// [0] = [WRITE] referenceAddresses
	//
	// [1] = [] router
	//
	// [2] = [] feeQuoter
	//
	// [3] = [] offrampLookupTable
	//
	// [4] = [WRITE] state
	//
	// [5] = [WRITE] externalExecutionConfig
	//
	// [6] = [WRITE] tokenPoolsSigner
	//
	// [7] = [WRITE, SIGNER] authority
	//
	// [8] = [] systemProgram
	//
	// [9] = [] program
	//
	// [10] = [] programData
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewInitializeInstructionBuilder creates a new `Initialize` instruction builder.
func NewInitializeInstructionBuilder() *Initialize {
	nd := &Initialize{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 11),
	}
	return nd
}

// SetReferenceAddressesAccount sets the "referenceAddresses" account.
func (inst *Initialize) SetReferenceAddressesAccount(referenceAddresses ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(referenceAddresses).WRITE()
	return inst
}

// GetReferenceAddressesAccount gets the "referenceAddresses" account.
func (inst *Initialize) GetReferenceAddressesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetRouterAccount sets the "router" account.
func (inst *Initialize) SetRouterAccount(router ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(router)
	return inst
}

// GetRouterAccount gets the "router" account.
func (inst *Initialize) GetRouterAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetFeeQuoterAccount sets the "feeQuoter" account.
func (inst *Initialize) SetFeeQuoterAccount(feeQuoter ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(feeQuoter)
	return inst
}

// GetFeeQuoterAccount gets the "feeQuoter" account.
func (inst *Initialize) GetFeeQuoterAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetOfframpLookupTableAccount sets the "offrampLookupTable" account.
func (inst *Initialize) SetOfframpLookupTableAccount(offrampLookupTable ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(offrampLookupTable)
	return inst
}

// GetOfframpLookupTableAccount gets the "offrampLookupTable" account.
func (inst *Initialize) GetOfframpLookupTableAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetStateAccount sets the "state" account.
func (inst *Initialize) SetStateAccount(state ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *Initialize) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetExternalExecutionConfigAccount sets the "externalExecutionConfig" account.
func (inst *Initialize) SetExternalExecutionConfigAccount(externalExecutionConfig ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(externalExecutionConfig).WRITE()
	return inst
}

// GetExternalExecutionConfigAccount gets the "externalExecutionConfig" account.
func (inst *Initialize) GetExternalExecutionConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetTokenPoolsSignerAccount sets the "tokenPoolsSigner" account.
func (inst *Initialize) SetTokenPoolsSignerAccount(tokenPoolsSigner ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(tokenPoolsSigner).WRITE()
	return inst
}

// GetTokenPoolsSignerAccount gets the "tokenPoolsSigner" account.
func (inst *Initialize) GetTokenPoolsSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *Initialize) SetAuthorityAccount(authority ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *Initialize) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[7]
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *Initialize) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *Initialize) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[8]
}

// SetProgramAccount sets the "program" account.
func (inst *Initialize) SetProgramAccount(program ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *Initialize) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[9]
}

// SetProgramDataAccount sets the "programData" account.
func (inst *Initialize) SetProgramDataAccount(programData ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(programData)
	return inst
}

// GetProgramDataAccount gets the "programData" account.
func (inst *Initialize) GetProgramDataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[10]
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
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.ReferenceAddresses is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Router is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.FeeQuoter is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.OfframpLookupTable is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.ExternalExecutionConfig is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.TokenPoolsSigner is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.Program is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
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
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=11]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     referenceAddresses", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("                 router", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("              feeQuoter", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("     offrampLookupTable", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("                  state", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("externalExecutionConfig", inst.AccountMetaSlice[5]))
						accountsBranch.Child(ag_format.Meta("       tokenPoolsSigner", inst.AccountMetaSlice[6]))
						accountsBranch.Child(ag_format.Meta("              authority", inst.AccountMetaSlice[7]))
						accountsBranch.Child(ag_format.Meta("          systemProgram", inst.AccountMetaSlice[8]))
						accountsBranch.Child(ag_format.Meta("                program", inst.AccountMetaSlice[9]))
						accountsBranch.Child(ag_format.Meta("            programData", inst.AccountMetaSlice[10]))
					})
				})
		})
}

func (obj Initialize) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *Initialize) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewInitializeInstruction declares a new Initialize instruction with the provided parameters and accounts.
func NewInitializeInstruction(
	// Accounts:
	referenceAddresses ag_solanago.PublicKey,
	router ag_solanago.PublicKey,
	feeQuoter ag_solanago.PublicKey,
	offrampLookupTable ag_solanago.PublicKey,
	state ag_solanago.PublicKey,
	externalExecutionConfig ag_solanago.PublicKey,
	tokenPoolsSigner ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	program ag_solanago.PublicKey,
	programData ag_solanago.PublicKey) *Initialize {
	return NewInitializeInstructionBuilder().
		SetReferenceAddressesAccount(referenceAddresses).
		SetRouterAccount(router).
		SetFeeQuoterAccount(feeQuoter).
		SetOfframpLookupTableAccount(offrampLookupTable).
		SetStateAccount(state).
		SetExternalExecutionConfigAccount(externalExecutionConfig).
		SetTokenPoolsSignerAccount(tokenPoolsSigner).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram).
		SetProgramAccount(program).
		SetProgramDataAccount(programData)
}
