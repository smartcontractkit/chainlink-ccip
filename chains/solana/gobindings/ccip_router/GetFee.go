// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ccip_router

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Calculates the fee for sending a message to the destination chain.
//
// # Arguments
//
// * `_ctx` - The context containing the accounts required for the fee calculation.
// * `dest_chain_selector` - The chain selector for the destination chain.
// * `message` - The message to be sent.
//
// # Additional accounts
//
// In addition to the fixed amount of accounts defined in the `GetFee` context,
// the following accounts must be provided:
//
// * First, the billing token config accounts for each token sent with the message, sequentially.
// For each token with no billing config account (i.e. tokens that cannot be possibly used as fee
// tokens, which also have no BPS fees enabled) the ZERO address must be provided instead.
// * Then, the per chain / per token config of every token sent with the message, sequentially
// in the same order.
//
// # Returns
//
// The fee amount in u64.
type GetFee struct {
	DestChainSelector *uint64
	Message           *Solana2AnyMessage

	// [0] = [] destChainState
	//
	// [1] = [] billingTokenConfig
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewGetFeeInstructionBuilder creates a new `GetFee` instruction builder.
func NewGetFeeInstructionBuilder() *GetFee {
	nd := &GetFee{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetDestChainSelector sets the "destChainSelector" parameter.
func (inst *GetFee) SetDestChainSelector(destChainSelector uint64) *GetFee {
	inst.DestChainSelector = &destChainSelector
	return inst
}

// SetMessage sets the "message" parameter.
func (inst *GetFee) SetMessage(message Solana2AnyMessage) *GetFee {
	inst.Message = &message
	return inst
}

// SetDestChainStateAccount sets the "destChainState" account.
func (inst *GetFee) SetDestChainStateAccount(destChainState ag_solanago.PublicKey) *GetFee {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(destChainState)
	return inst
}

// GetDestChainStateAccount gets the "destChainState" account.
func (inst *GetFee) GetDestChainStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetBillingTokenConfigAccount sets the "billingTokenConfig" account.
func (inst *GetFee) SetBillingTokenConfigAccount(billingTokenConfig ag_solanago.PublicKey) *GetFee {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(billingTokenConfig)
	return inst
}

// GetBillingTokenConfigAccount gets the "billingTokenConfig" account.
func (inst *GetFee) GetBillingTokenConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

func (inst GetFee) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_GetFee,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst GetFee) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *GetFee) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.DestChainSelector == nil {
			return errors.New("DestChainSelector parameter is not set")
		}
		if inst.Message == nil {
			return errors.New("Message parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.DestChainState is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.BillingTokenConfig is not set")
		}
	}
	return nil
}

func (inst *GetFee) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("GetFee")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("DestChainSelector", *inst.DestChainSelector))
						paramsBranch.Child(ag_format.Param("          Message", *inst.Message))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("    destChainState", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("billingTokenConfig", inst.AccountMetaSlice[1]))
					})
				})
		})
}

func (obj GetFee) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `DestChainSelector` param:
	err = encoder.Encode(obj.DestChainSelector)
	if err != nil {
		return err
	}
	// Serialize `Message` param:
	err = encoder.Encode(obj.Message)
	if err != nil {
		return err
	}
	return nil
}
func (obj *GetFee) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `DestChainSelector`:
	err = decoder.Decode(&obj.DestChainSelector)
	if err != nil {
		return err
	}
	// Deserialize `Message`:
	err = decoder.Decode(&obj.Message)
	if err != nil {
		return err
	}
	return nil
}

// NewGetFeeInstruction declares a new GetFee instruction with the provided parameters and accounts.
func NewGetFeeInstruction(
	// Parameters:
	destChainSelector uint64,
	message Solana2AnyMessage,
	// Accounts:
	destChainState ag_solanago.PublicKey,
	billingTokenConfig ag_solanago.PublicKey) *GetFee {
	return NewGetFeeInstructionBuilder().
		SetDestChainSelector(destChainSelector).
		SetMessage(message).
		SetDestChainStateAccount(destChainState).
		SetBillingTokenConfigAccount(billingTokenConfig)
}
