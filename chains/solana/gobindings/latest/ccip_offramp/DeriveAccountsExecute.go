// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ccip_offramp

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Automatically derives all acounts required to call `ccip_execute`.
//
// This method receives the bare minimum amount of information needed to construct
// the entire account list to execute a transaction, and builds it iteratively
// over the course of multiple calls.
//
// The return type contains:
//
// * `accounts_to_save`: The caller must append these accounts to a list they maintain.
// When complete, this list will contain all accounts needed to call `ccip_execute`.
// * `ask_again_with`: When `next_stage` is not empty, the caller must call `derive_accounts_execute`
// again, including exactly these accounts as the `remaining_accounts`.
// * `lookup_tables_to_save`: The caller must save those LUTs. They can be used for `ccip_execute`.
// * `current_stage`: A string describing the current stage of the derivation process. When the stage
// is "TokenTransferStaticAccounts/<N>/0", it means the `accounts_to_save` block in this response contains
// all accounts relating to the Nth token being transferred. Use this information to construct
// the `token_indexes` vector that `execute` requires.
// * `next_stage`: If nonempty, this means the instruction must get called again with this value
// as the `stage` argument.
//
// Therefore, and starting with an empty `remaining_accounts` list, the caller must repeteadly
// call `derive_accounts_execute` until `next_stage` is returned empty.
//
// # Arguments
//
// * `ctx`: Context containing only the offramp config.
// * `stage`: Requested derivation stage. Pass "Start" the first time, then for each subsequent
// call, pass the value returned in `response.next_stage` until empty.
// * `params`:
// * `execute_caller`: Public key of the account that will sign the call to `ccip_execute`.
// * `message_accounts`: If the transaction involves messaging, the message accounts.
// * `source_chain_selector`: CCIP chain selector for the source chain.
// * `mints_of_transferred_token`: List of all token mints for tokens being transferred (i.e.
// the entries in `report.message.token_amounts.destination_address`.)
// * `merkle_root`: Merkle root as per the commit report.
// * `buffer_id`: If the execution will be buffered, the buffer id that will be used by the
// `execute_caller`: If the execution will not be buffered, this should be empty.
// * `token_receiver`: Receiver of token transfers, if any (i.e. report.message.token_receiver)
type DeriveAccountsExecute struct {
	Params *DeriveAccountsExecuteParams
	Stage  *string

	// [0] = [] config
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewDeriveAccountsExecuteInstructionBuilder creates a new `DeriveAccountsExecute` instruction builder.
func NewDeriveAccountsExecuteInstructionBuilder() *DeriveAccountsExecute {
	nd := &DeriveAccountsExecute{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 1),
	}
	return nd
}

// SetParams sets the "params" parameter.
func (inst *DeriveAccountsExecute) SetParams(params DeriveAccountsExecuteParams) *DeriveAccountsExecute {
	inst.Params = &params
	return inst
}

// SetStage sets the "stage" parameter.
func (inst *DeriveAccountsExecute) SetStage(stage string) *DeriveAccountsExecute {
	inst.Stage = &stage
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *DeriveAccountsExecute) SetConfigAccount(config ag_solanago.PublicKey) *DeriveAccountsExecute {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *DeriveAccountsExecute) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

func (inst DeriveAccountsExecute) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_DeriveAccountsExecute,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst DeriveAccountsExecute) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *DeriveAccountsExecute) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Params == nil {
			return errors.New("Params parameter is not set")
		}
		if inst.Stage == nil {
			return errors.New("Stage parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Config is not set")
		}
	}
	return nil
}

func (inst *DeriveAccountsExecute) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DeriveAccountsExecute")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Params", *inst.Params))
						paramsBranch.Child(ag_format.Param(" Stage", *inst.Stage))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=1]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("config", inst.AccountMetaSlice[0]))
					})
				})
		})
}

func (obj DeriveAccountsExecute) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Params` param:
	err = encoder.Encode(obj.Params)
	if err != nil {
		return err
	}
	// Serialize `Stage` param:
	err = encoder.Encode(obj.Stage)
	if err != nil {
		return err
	}
	return nil
}
func (obj *DeriveAccountsExecute) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Params`:
	err = decoder.Decode(&obj.Params)
	if err != nil {
		return err
	}
	// Deserialize `Stage`:
	err = decoder.Decode(&obj.Stage)
	if err != nil {
		return err
	}
	return nil
}

// NewDeriveAccountsExecuteInstruction declares a new DeriveAccountsExecute instruction with the provided parameters and accounts.
func NewDeriveAccountsExecuteInstruction(
	// Parameters:
	params DeriveAccountsExecuteParams,
	stage string,
	// Accounts:
	config ag_solanago.PublicKey) *DeriveAccountsExecute {
	return NewDeriveAccountsExecuteInstructionBuilder().
		SetParams(params).
		SetStage(stage).
		SetConfigAccount(config)
}
