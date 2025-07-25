// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_messenger_minter

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// RemoveRemoteTokenMessenger is the `removeRemoteTokenMessenger` instruction.
type RemoveRemoteTokenMessenger struct {
	Params *RemoveRemoteTokenMessengerParams

	// [0] = [WRITE, SIGNER] payee
	//
	// [1] = [SIGNER] owner
	//
	// [2] = [] tokenMessenger
	//
	// [3] = [WRITE] remoteTokenMessenger
	//
	// [4] = [] eventAuthority
	//
	// [5] = [] program
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewRemoveRemoteTokenMessengerInstructionBuilder creates a new `RemoveRemoteTokenMessenger` instruction builder.
func NewRemoveRemoteTokenMessengerInstructionBuilder() *RemoveRemoteTokenMessenger {
	nd := &RemoveRemoteTokenMessenger{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	return nd
}

// SetParams sets the "params" parameter.
func (inst *RemoveRemoteTokenMessenger) SetParams(params RemoveRemoteTokenMessengerParams) *RemoveRemoteTokenMessenger {
	inst.Params = &params
	return inst
}

// SetPayeeAccount sets the "payee" account.
func (inst *RemoveRemoteTokenMessenger) SetPayeeAccount(payee ag_solanago.PublicKey) *RemoveRemoteTokenMessenger {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(payee).WRITE().SIGNER()
	return inst
}

// GetPayeeAccount gets the "payee" account.
func (inst *RemoveRemoteTokenMessenger) GetPayeeAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetOwnerAccount sets the "owner" account.
func (inst *RemoveRemoteTokenMessenger) SetOwnerAccount(owner ag_solanago.PublicKey) *RemoveRemoteTokenMessenger {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(owner).SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *RemoveRemoteTokenMessenger) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetTokenMessengerAccount sets the "tokenMessenger" account.
func (inst *RemoveRemoteTokenMessenger) SetTokenMessengerAccount(tokenMessenger ag_solanago.PublicKey) *RemoveRemoteTokenMessenger {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(tokenMessenger)
	return inst
}

// GetTokenMessengerAccount gets the "tokenMessenger" account.
func (inst *RemoveRemoteTokenMessenger) GetTokenMessengerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetRemoteTokenMessengerAccount sets the "remoteTokenMessenger" account.
func (inst *RemoveRemoteTokenMessenger) SetRemoteTokenMessengerAccount(remoteTokenMessenger ag_solanago.PublicKey) *RemoveRemoteTokenMessenger {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(remoteTokenMessenger).WRITE()
	return inst
}

// GetRemoteTokenMessengerAccount gets the "remoteTokenMessenger" account.
func (inst *RemoveRemoteTokenMessenger) GetRemoteTokenMessengerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetEventAuthorityAccount sets the "eventAuthority" account.
func (inst *RemoveRemoteTokenMessenger) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *RemoveRemoteTokenMessenger {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(eventAuthority)
	return inst
}

// GetEventAuthorityAccount gets the "eventAuthority" account.
func (inst *RemoveRemoteTokenMessenger) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetProgramAccount sets the "program" account.
func (inst *RemoveRemoteTokenMessenger) SetProgramAccount(program ag_solanago.PublicKey) *RemoveRemoteTokenMessenger {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *RemoveRemoteTokenMessenger) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

func (inst RemoveRemoteTokenMessenger) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_RemoveRemoteTokenMessenger,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst RemoveRemoteTokenMessenger) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *RemoveRemoteTokenMessenger) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Params == nil {
			return errors.New("Params parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Payee is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.TokenMessenger is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.RemoteTokenMessenger is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *RemoveRemoteTokenMessenger) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("RemoveRemoteTokenMessenger")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Params", *inst.Params))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("               payee", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("               owner", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("      tokenMessenger", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("remoteTokenMessenger", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("      eventAuthority", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("             program", inst.AccountMetaSlice[5]))
					})
				})
		})
}

func (obj RemoveRemoteTokenMessenger) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Params` param:
	err = encoder.Encode(obj.Params)
	if err != nil {
		return err
	}
	return nil
}
func (obj *RemoveRemoteTokenMessenger) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Params`:
	err = decoder.Decode(&obj.Params)
	if err != nil {
		return err
	}
	return nil
}

// NewRemoveRemoteTokenMessengerInstruction declares a new RemoveRemoteTokenMessenger instruction with the provided parameters and accounts.
func NewRemoveRemoteTokenMessengerInstruction(
	// Parameters:
	params RemoveRemoteTokenMessengerParams,
	// Accounts:
	payee ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	tokenMessenger ag_solanago.PublicKey,
	remoteTokenMessenger ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey,
	program ag_solanago.PublicKey) *RemoveRemoteTokenMessenger {
	return NewRemoveRemoteTokenMessengerInstructionBuilder().
		SetParams(params).
		SetPayeeAccount(payee).
		SetOwnerAccount(owner).
		SetTokenMessengerAccount(tokenMessenger).
		SetRemoteTokenMessengerAccount(remoteTokenMessenger).
		SetEventAuthorityAccount(eventAuthority).
		SetProgramAccount(program)
}
