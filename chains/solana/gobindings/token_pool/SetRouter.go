// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_pool

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SetRouter is the `setRouter` instruction.
type SetRouter struct {
	NewRouter *ag_solanago.PublicKey

	// [0] = [WRITE] config
	//
	// [1] = [SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewSetRouterInstructionBuilder creates a new `SetRouter` instruction builder.
func NewSetRouterInstructionBuilder() *SetRouter {
	nd := &SetRouter{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetNewRouter sets the "newRouter" parameter.
func (inst *SetRouter) SetNewRouter(newRouter ag_solanago.PublicKey) *SetRouter {
	inst.NewRouter = &newRouter
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *SetRouter) SetConfigAccount(config ag_solanago.PublicKey) *SetRouter {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config).WRITE()
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *SetRouter) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *SetRouter) SetAuthorityAccount(authority ag_solanago.PublicKey) *SetRouter {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *SetRouter) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

func (inst SetRouter) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SetRouter,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetRouter) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetRouter) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.NewRouter == nil {
			return errors.New("NewRouter parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
	}
	return nil
}

func (inst *SetRouter) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetRouter")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("NewRouter", *inst.NewRouter))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("   config", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("authority", inst.AccountMetaSlice[1]))
					})
				})
		})
}

func (obj SetRouter) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NewRouter` param:
	err = encoder.Encode(obj.NewRouter)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SetRouter) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NewRouter`:
	err = decoder.Decode(&obj.NewRouter)
	if err != nil {
		return err
	}
	return nil
}

// NewSetRouterInstruction declares a new SetRouter instruction with the provided parameters and accounts.
func NewSetRouterInstruction(
	// Parameters:
	newRouter ag_solanago.PublicKey,
	// Accounts:
	config ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *SetRouter {
	return NewSetRouterInstructionBuilder().
		SetNewRouter(newRouter).
		SetConfigAccount(config).
		SetAuthorityAccount(authority)
}
