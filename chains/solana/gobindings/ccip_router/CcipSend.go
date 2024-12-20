// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ccip_router

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ON RAMP FLOW
// Sends a message to the destination chain.
//
// Request a message to be sent to the destination chain.
// The method name needs to be ccip_send with Anchor encoding.
// This function is called by the CCIP Sender Contract (or final user) to send a message to the CCIP Router.
// The message will be sent to the receiver on the destination chain selector.
// This message emits the event CCIPSendRequested with all the necessary data to be retrieved by the OffChain Code
//
// # Arguments
//
// * `ctx` - The context containing the accounts required for sending the message.
// * `dest_chain_selector` - The chain selector for the destination chain.
// * `message` - The message to be sent. The size limit of data is 256 bytes.
type CcipSend struct {
	DestChainSelector *uint64
	Message           *Solana2AnyMessage

	// [0] = [] config
	//
	// [1] = [WRITE] destChainState
	//
	// [2] = [WRITE] nonce
	//
	// [3] = [WRITE, SIGNER] authority
	//
	// [4] = [] systemProgram
	//
	// [5] = [] feeTokenProgram
	// ··········· type of a specific program (which would enforce its ID). Thus, it's an UncheckedAccount
	// ··········· with a constraint enforcing that it is one of the two allowed programs.
	//
	// [6] = [] feeTokenMint
	//
	// [7] = [] feeTokenConfig
	//
	// [8] = [] feeTokenUserAssociatedAccount
	// ··········· CHECK this is the associated token account for the user paying the fee.
	// ··········· If paying with native SOL, this must be the zero address.
	//
	// [9] = [WRITE] feeTokenReceiver
	//
	// [10] = [] feeBillingSigner
	//
	// [11] = [WRITE] tokenPoolsSigner
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewCcipSendInstructionBuilder creates a new `CcipSend` instruction builder.
func NewCcipSendInstructionBuilder() *CcipSend {
	nd := &CcipSend{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 12),
	}
	return nd
}

// SetDestChainSelector sets the "destChainSelector" parameter.
func (inst *CcipSend) SetDestChainSelector(destChainSelector uint64) *CcipSend {
	inst.DestChainSelector = &destChainSelector
	return inst
}

// SetMessage sets the "message" parameter.
func (inst *CcipSend) SetMessage(message Solana2AnyMessage) *CcipSend {
	inst.Message = &message
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *CcipSend) SetConfigAccount(config ag_solanago.PublicKey) *CcipSend {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *CcipSend) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetDestChainStateAccount sets the "destChainState" account.
func (inst *CcipSend) SetDestChainStateAccount(destChainState ag_solanago.PublicKey) *CcipSend {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(destChainState).WRITE()
	return inst
}

// GetDestChainStateAccount gets the "destChainState" account.
func (inst *CcipSend) GetDestChainStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetNonceAccount sets the "nonce" account.
func (inst *CcipSend) SetNonceAccount(nonce ag_solanago.PublicKey) *CcipSend {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(nonce).WRITE()
	return inst
}

// GetNonceAccount gets the "nonce" account.
func (inst *CcipSend) GetNonceAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *CcipSend) SetAuthorityAccount(authority ag_solanago.PublicKey) *CcipSend {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(authority).WRITE().SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *CcipSend) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *CcipSend) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *CcipSend {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *CcipSend) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetFeeTokenProgramAccount sets the "feeTokenProgram" account.
// type of a specific program (which would enforce its ID). Thus, it's an UncheckedAccount
// with a constraint enforcing that it is one of the two allowed programs.
func (inst *CcipSend) SetFeeTokenProgramAccount(feeTokenProgram ag_solanago.PublicKey) *CcipSend {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(feeTokenProgram)
	return inst
}

// GetFeeTokenProgramAccount gets the "feeTokenProgram" account.
// type of a specific program (which would enforce its ID). Thus, it's an UncheckedAccount
// with a constraint enforcing that it is one of the two allowed programs.
func (inst *CcipSend) GetFeeTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetFeeTokenMintAccount sets the "feeTokenMint" account.
func (inst *CcipSend) SetFeeTokenMintAccount(feeTokenMint ag_solanago.PublicKey) *CcipSend {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(feeTokenMint)
	return inst
}

// GetFeeTokenMintAccount gets the "feeTokenMint" account.
func (inst *CcipSend) GetFeeTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

// SetFeeTokenConfigAccount sets the "feeTokenConfig" account.
func (inst *CcipSend) SetFeeTokenConfigAccount(feeTokenConfig ag_solanago.PublicKey) *CcipSend {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(feeTokenConfig)
	return inst
}

// GetFeeTokenConfigAccount gets the "feeTokenConfig" account.
func (inst *CcipSend) GetFeeTokenConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[7]
}

// SetFeeTokenUserAssociatedAccountAccount sets the "feeTokenUserAssociatedAccount" account.
// CHECK this is the associated token account for the user paying the fee.
// If paying with native SOL, this must be the zero address.
func (inst *CcipSend) SetFeeTokenUserAssociatedAccountAccount(feeTokenUserAssociatedAccount ag_solanago.PublicKey) *CcipSend {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(feeTokenUserAssociatedAccount)
	return inst
}

// GetFeeTokenUserAssociatedAccountAccount gets the "feeTokenUserAssociatedAccount" account.
// CHECK this is the associated token account for the user paying the fee.
// If paying with native SOL, this must be the zero address.
func (inst *CcipSend) GetFeeTokenUserAssociatedAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[8]
}

// SetFeeTokenReceiverAccount sets the "feeTokenReceiver" account.
func (inst *CcipSend) SetFeeTokenReceiverAccount(feeTokenReceiver ag_solanago.PublicKey) *CcipSend {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(feeTokenReceiver).WRITE()
	return inst
}

// GetFeeTokenReceiverAccount gets the "feeTokenReceiver" account.
func (inst *CcipSend) GetFeeTokenReceiverAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[9]
}

// SetFeeBillingSignerAccount sets the "feeBillingSigner" account.
func (inst *CcipSend) SetFeeBillingSignerAccount(feeBillingSigner ag_solanago.PublicKey) *CcipSend {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(feeBillingSigner)
	return inst
}

// GetFeeBillingSignerAccount gets the "feeBillingSigner" account.
func (inst *CcipSend) GetFeeBillingSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[10]
}

// SetTokenPoolsSignerAccount sets the "tokenPoolsSigner" account.
func (inst *CcipSend) SetTokenPoolsSignerAccount(tokenPoolsSigner ag_solanago.PublicKey) *CcipSend {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(tokenPoolsSigner).WRITE()
	return inst
}

// GetTokenPoolsSignerAccount gets the "tokenPoolsSigner" account.
func (inst *CcipSend) GetTokenPoolsSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[11]
}

func (inst CcipSend) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CcipSend,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CcipSend) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CcipSend) Validate() error {
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
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.DestChainState is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Nonce is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.FeeTokenProgram is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.FeeTokenMint is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.FeeTokenConfig is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.FeeTokenUserAssociatedAccount is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.FeeTokenReceiver is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.FeeBillingSigner is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.TokenPoolsSigner is not set")
		}
	}
	return nil
}

func (inst *CcipSend) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CcipSend")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("DestChainSelector", *inst.DestChainSelector))
						paramsBranch.Child(ag_format.Param("          Message", *inst.Message))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=12]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                config", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("        destChainState", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("                 nonce", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("             authority", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("         systemProgram", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("       feeTokenProgram", inst.AccountMetaSlice[5]))
						accountsBranch.Child(ag_format.Meta("          feeTokenMint", inst.AccountMetaSlice[6]))
						accountsBranch.Child(ag_format.Meta("        feeTokenConfig", inst.AccountMetaSlice[7]))
						accountsBranch.Child(ag_format.Meta("feeTokenUserAssociated", inst.AccountMetaSlice[8]))
						accountsBranch.Child(ag_format.Meta("      feeTokenReceiver", inst.AccountMetaSlice[9]))
						accountsBranch.Child(ag_format.Meta("      feeBillingSigner", inst.AccountMetaSlice[10]))
						accountsBranch.Child(ag_format.Meta("      tokenPoolsSigner", inst.AccountMetaSlice[11]))
					})
				})
		})
}

func (obj CcipSend) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
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
func (obj *CcipSend) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
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

// NewCcipSendInstruction declares a new CcipSend instruction with the provided parameters and accounts.
func NewCcipSendInstruction(
	// Parameters:
	destChainSelector uint64,
	message Solana2AnyMessage,
	// Accounts:
	config ag_solanago.PublicKey,
	destChainState ag_solanago.PublicKey,
	nonce ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	feeTokenProgram ag_solanago.PublicKey,
	feeTokenMint ag_solanago.PublicKey,
	feeTokenConfig ag_solanago.PublicKey,
	feeTokenUserAssociatedAccount ag_solanago.PublicKey,
	feeTokenReceiver ag_solanago.PublicKey,
	feeBillingSigner ag_solanago.PublicKey,
	tokenPoolsSigner ag_solanago.PublicKey) *CcipSend {
	return NewCcipSendInstructionBuilder().
		SetDestChainSelector(destChainSelector).
		SetMessage(message).
		SetConfigAccount(config).
		SetDestChainStateAccount(destChainState).
		SetNonceAccount(nonce).
		SetAuthorityAccount(authority).
		SetSystemProgramAccount(systemProgram).
		SetFeeTokenProgramAccount(feeTokenProgram).
		SetFeeTokenMintAccount(feeTokenMint).
		SetFeeTokenConfigAccount(feeTokenConfig).
		SetFeeTokenUserAssociatedAccountAccount(feeTokenUserAssociatedAccount).
		SetFeeTokenReceiverAccount(feeTokenReceiver).
		SetFeeBillingSignerAccount(feeBillingSigner).
		SetTokenPoolsSignerAccount(tokenPoolsSigner)
}