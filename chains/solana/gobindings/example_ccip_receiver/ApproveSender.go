// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package example_ccip_receiver

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ApproveSender is the `approveSender` instruction.
type ApproveSender struct {
	ChainSelector *uint64
	RemoteAddress *[]byte

	// [0] = [WRITE] state
	//
	// [1] = [WRITE] approvedSender
	//
	// [2] = [WRITE, SIGNER] authority
	//
	// [3] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewApproveSenderInstructionBuilder creates a new `ApproveSender` instruction builder.
func NewApproveSenderInstructionBuilder() *ApproveSender {
	nd := &ApproveSender{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetChainSelector sets the "chainSelector" parameter.
func (inst *ApproveSender) SetChainSelector(chainSelector uint64) *ApproveSender {
	inst.ChainSelector = &chainSelector
	return inst
}

// SetRemoteAddress sets the "remoteAddress" parameter.
func (inst *ApproveSender) SetRemoteAddress(remoteAddress []byte) *ApproveSender {
	inst.RemoteAddress = &remoteAddress
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *ApproveSender) SetStateAccount(state ag_solanago.PublicKey) *ApproveSender {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *ApproveSender) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetApprovedSenderAccount sets the "approvedSender" account.
func (inst *ApproveSender) SetApprovedSenderAccount(approvedSender ag_solanago.PublicKey) *ApproveSender {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(approvedSender).WRITE()
	return inst
}

// GetApprovedSenderAccount gets the "approvedSender" account.
func (inst *ApproveSender) GetApprovedSenderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *ApproveSender) SetAuthorityAccount(authority ag_solanago.PublicKey) *ApproveSender {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *ApproveSender) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *ApproveSender) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *ApproveSender {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *ApproveSender) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

func (inst ApproveSender) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ApproveSender,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ApproveSender) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ApproveSender) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.ChainSelector == nil {
			return errors.New("ChainSelector parameter is not set")
		}
		if inst.RemoteAddress == nil {
			return errors.New("RemoteAddress parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ApprovedSender is not set")
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

func (inst *ApproveSender) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ApproveSender")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("ChainSelector", *inst.ChainSelector))
						paramsBranch.Child(ag_format.Param("RemoteAddress", *inst.RemoteAddress))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("         state", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("approvedSender", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("     authority", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta(" systemProgram", inst.AccountMetaSlice[3]))
					})
				})
		})
}

func (obj ApproveSender) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ChainSelector` param:
	err = encoder.Encode(obj.ChainSelector)
	if err != nil {
		return err
	}
	// Serialize `RemoteAddress` param:
	err = encoder.Encode(obj.RemoteAddress)
	if err != nil {
		return err
	}
	return nil
}
func (obj *ApproveSender) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ChainSelector`:
	err = decoder.Decode(&obj.ChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `RemoteAddress`:
	err = decoder.Decode(&obj.RemoteAddress)
	if err != nil {
		return err
	}
	return nil
}

// NewApproveSenderInstruction declares a new ApproveSender instruction with the provided parameters and accounts.
func NewApproveSenderInstruction(
	// Parameters:
	chainSelector uint64,
	remoteAddress []byte,
	// Accounts:
	state ag_solanago.PublicKey,
	approvedSender ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *ApproveSender {
	return NewApproveSenderInstructionBuilder().
		SetChainSelector(chainSelector).
		SetRemoteAddress(remoteAddress).
		SetStateAccount(state).
		SetApprovedSenderAccount(approvedSender).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram)
}
