// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package test_ccip_invalid_receiver

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CcipReceive is the `ccipReceive` instruction.
type CcipReceive struct {
	Message *Any2SVMMessage

	// [0] = [WRITE, SIGNER] authority
	//
	// [1] = [] offrampProgram
	// ··········· CHECK offramp program: exists only to derive the allowed offramp PDA
	// ··········· and the authority PDA. Must be second.
	//
	// [2] = [] allowedOfframp
	// ··········· CHECK PDA of the router program verifying the signer is an allowed offramp.
	// ··········· If PDA does not exist, the router doesn't allow this offramp
	//
	// [3] = [WRITE] counter
	//
	// [4] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewCcipReceiveInstructionBuilder creates a new `CcipReceive` instruction builder.
func NewCcipReceiveInstructionBuilder() *CcipReceive {
	nd := &CcipReceive{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetMessage sets the "message" parameter.
func (inst *CcipReceive) SetMessage(message Any2SVMMessage) *CcipReceive {
	inst.Message = &message
	return inst
}

// SetAuthorityAccount sets the "authority" account.
func (inst *CcipReceive) SetAuthorityAccount(authority ag_solanago.PublicKey) *CcipReceive {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *CcipReceive) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetOfframpProgramAccount sets the "offrampProgram" account.
// CHECK offramp program: exists only to derive the allowed offramp PDA
// and the authority PDA. Must be second.
func (inst *CcipReceive) SetOfframpProgramAccount(offrampProgram ag_solanago.PublicKey) *CcipReceive {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(offrampProgram)
	return inst
}

// GetOfframpProgramAccount gets the "offrampProgram" account.
// CHECK offramp program: exists only to derive the allowed offramp PDA
// and the authority PDA. Must be second.
func (inst *CcipReceive) GetOfframpProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetAllowedOfframpAccount sets the "allowedOfframp" account.
// CHECK PDA of the router program verifying the signer is an allowed offramp.
// If PDA does not exist, the router doesn't allow this offramp
func (inst *CcipReceive) SetAllowedOfframpAccount(allowedOfframp ag_solanago.PublicKey) *CcipReceive {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(allowedOfframp)
	return inst
}

// GetAllowedOfframpAccount gets the "allowedOfframp" account.
// CHECK PDA of the router program verifying the signer is an allowed offramp.
// If PDA does not exist, the router doesn't allow this offramp
func (inst *CcipReceive) GetAllowedOfframpAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetCounterAccount sets the "counter" account.
func (inst *CcipReceive) SetCounterAccount(counter ag_solanago.PublicKey) *CcipReceive {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(counter).WRITE()
	return inst
}

// GetCounterAccount gets the "counter" account.
func (inst *CcipReceive) GetCounterAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *CcipReceive) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *CcipReceive {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *CcipReceive) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

func (inst CcipReceive) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CcipReceive,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CcipReceive) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CcipReceive) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Message == nil {
			return errors.New("Message parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.OfframpProgram is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.AllowedOfframp is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Counter is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *CcipReceive) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CcipReceive")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Message", *inst.Message))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     authority", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("offrampProgram", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("allowedOfframp", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("       counter", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta(" systemProgram", inst.AccountMetaSlice[4]))
					})
				})
		})
}

func (obj CcipReceive) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Message` param:
	err = encoder.Encode(obj.Message)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CcipReceive) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Message`:
	err = decoder.Decode(&obj.Message)
	if err != nil {
		return err
	}
	return nil
}

// NewCcipReceiveInstruction declares a new CcipReceive instruction with the provided parameters and accounts.
func NewCcipReceiveInstruction(
	// Parameters:
	message Any2SVMMessage,
	// Accounts:
	authority ag_solanago.PublicKey,
	offrampProgram ag_solanago.PublicKey,
	allowedOfframp ag_solanago.PublicKey,
	counter ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *CcipReceive {
	return NewCcipReceiveInstructionBuilder().
		SetMessage(message).
		SetAuthorityAccount(authority).
		SetOfframpProgramAccount(offrampProgram).
		SetAllowedOfframpAccount(allowedOfframp).
		SetCounterAccount(counter).
		SetSystemProgramAccount(systemProgram)
}
