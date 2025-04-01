// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ccip_offramp

import (
	"bytes"
	"fmt"
	ag_spew "github.com/davecgh/go-spew/spew"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_text "github.com/gagliardetto/solana-go/text"
	ag_treeout "github.com/gagliardetto/treeout"
)

var ProgramID ag_solanago.PublicKey

func SetProgramID(pubkey ag_solanago.PublicKey) {
	ProgramID = pubkey
	ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
}

const ProgramName = "CcipOfframp"

func init() {
	if !ProgramID.IsZero() {
		ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
	}
}

var (
	// Initialization Flow //
	// Initializes the CCIP Offramp, except for the config account (due to stack size limitations).
	//
	// The initialization of the Offramp is responsibility of Admin, nothing more than calling these
	// initialization methods should be done first.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for initialization.
	Instruction_Initialize = ag_binary.TypeID([8]byte{175, 175, 109, 31, 13, 152, 155, 237})

	// Initializes the CCIP Offramp Config account.
	//
	// The initialization of the Offramp is responsibility of Admin, nothing more than calling these
	// initialization methods should be done first.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for initialization of the config.
	// * `svm_chain_selector` - The chain selector for SVM.
	// * `enable_execution_after` - The minimum amount of time required between a message has been committed and can be manually executed.
	Instruction_InitializeConfig = ag_binary.TypeID([8]byte{208, 127, 21, 1, 194, 190, 196, 70})

	// Transfers the ownership of the router to a new proposed owner.
	//
	// Shared func signature with other programs
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for the transfer.
	// * `proposed_owner` - The public key of the new proposed owner.
	Instruction_TransferOwnership = ag_binary.TypeID([8]byte{65, 177, 215, 73, 53, 45, 99, 47})

	// Accepts the ownership of the router by the proposed owner.
	//
	// Shared func signature with other programs
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for accepting ownership.
	// The new owner must be a signer of the transaction.
	Instruction_AcceptOwnership = ag_binary.TypeID([8]byte{172, 23, 43, 13, 238, 213, 85, 150})

	// Sets the default code version to be used. This is then used by the slim routing layer to determine
	// which version of the versioned business logic module (`instructions`) to use. Only the admin may set this.
	//
	// Shared func signature with other programs
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for updating the configuration.
	// * `code_version` - The new code version to be set as default.
	Instruction_SetDefaultCodeVersion = ag_binary.TypeID([8]byte{47, 151, 233, 254, 121, 82, 206, 152})

	// Updates reference addresses in the offramp contract, such as
	// the CCIP router, Fee Quoter, and the Offramp Lookup Table.
	// Only the Admin may update these addresses.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for updating the reference addresses.
	// * `router` - The router address to be set.
	// * `fee_quoter` - The fee_quoter address to be set.
	// * `offramp_lookup_table` - The offramp_lookup_table address to be set.
	// * `rmn_remote` - The rmn_remote address to be set.
	Instruction_UpdateReferenceAddresses = ag_binary.TypeID([8]byte{119, 179, 218, 249, 217, 184, 181, 9})

	// Adds a new source chain selector with its config to the offramp.
	//
	// The Admin needs to add any new chain supported.
	// When adding a new chain, the Admin needs to specify if it's enabled or not.
	//
	// # Arguments
	Instruction_AddSourceChain = ag_binary.TypeID([8]byte{26, 58, 148, 88, 190, 27, 2, 144})

	// Disables the source chain selector.
	//
	// The Admin is the only one able to disable the chain selector as source. This method is thought of as an emergency kill-switch.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for disabling the chain selector.
	// * `source_chain_selector` - The source chain selector to be disabled.
	Instruction_DisableSourceChainSelector = ag_binary.TypeID([8]byte{58, 101, 54, 252, 248, 31, 226, 121})

	// Updates the configuration of the source chain selector.
	//
	// The Admin is the only one able to update the source chain config.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for updating the chain selector.
	// * `source_chain_selector` - The source chain selector to be updated.
	// * `source_chain_config` - The new configuration for the source chain.
	Instruction_UpdateSourceChainConfig = ag_binary.TypeID([8]byte{52, 85, 37, 124, 209, 140, 181, 104})

	// Updates the SVM chain selector in the offramp configuration.
	//
	// This method should only be used if there was an error with the initial configuration or if the solana chain selector changes.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for updating the configuration.
	// * `new_chain_selector` - The new chain selector for SVM.
	Instruction_UpdateSvmChainSelector = ag_binary.TypeID([8]byte{164, 212, 71, 101, 166, 113, 26, 93})

	// Updates the minimum amount of time required between a message being committed and when it can be manually executed.
	//
	// This is part of the OffRamp Configuration for SVM.
	// The Admin is the only one able to update this config.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for updating the configuration.
	// * `new_enable_manual_execution_after` - The new minimum amount of time required.
	Instruction_UpdateEnableManualExecutionAfter = ag_binary.TypeID([8]byte{157, 236, 73, 92, 84, 197, 152, 105})

	// Sets the OCR configuration.
	// Only CCIP Admin can set the OCR configuration.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for setting the OCR configuration.
	// * `plugin_type` - The type of OCR plugin [0: Commit, 1: Execution].
	// * `config_info` - The OCR configuration information.
	// * `signers` - The list of signers.
	// * `transmitters` - The list of transmitters.
	Instruction_SetOcrConfig = ag_binary.TypeID([8]byte{4, 131, 107, 110, 250, 158, 244, 200})

	// Off Ramp Flow //
	// Commits a report to the router, containing a Merkle Root.
	//
	// The method name needs to be commit with Anchor encoding.
	//
	// This function is called by the OffChain when committing one Report to the SVM Router.
	// In this Flow only one report is sent, the Commit Report. This is different as EVM does,
	// this is because here all the chain state is stored in one account per Merkle Tree Root.
	// So, to avoid having to send a dynamic size array of accounts, in this message only one Commit Report Account is sent.
	// This message validates the signatures of the report and stores the Merkle Root in the Commit Report Account.
	// The Report must contain an interval of messages, and the min of them must be the next sequence number expected.
	// The max size of the interval is 64.
	// This message emits two events: CommitReportAccepted and Transmitted.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for the commit.
	// * `report_context_byte_words` - consists of:
	// * report_context_byte_words[0]: ConfigDigest
	// * report_context_byte_words[1]: 24 byte padding, 8 byte sequence number
	// * `raw_report` - The serialized commit input report, single merkle root with RMN signatures and price updates
	// * `rs` - slice of R components of signatures
	// * `ss` - slice of S components of signatures
	// * `raw_vs` - array of V components of signatures
	Instruction_Commit = ag_binary.TypeID([8]byte{223, 140, 142, 165, 229, 208, 156, 74})

	// Commits a report to the router, with price updates only.
	//
	// The method name needs to be commit with Anchor encoding.
	//
	// This function is called by the OffChain when committing one Report to the SVM Router,
	// containing only price updates and no merkle root.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for the commit.
	// * `report_context_byte_words` - consists of:
	// * report_context_byte_words[0]: ConfigDigest
	// * report_context_byte_words[1]: 24 byte padding, 8 byte sequence number
	// * `raw_report` - The serialized commit input report containing the price updates,
	// with no merkle root.
	// * `rs` - slice of R components of signatures
	// * `ss` - slice of S components of signatures
	// * `raw_vs` - array of V components of signatures
	Instruction_CommitPriceOnly = ag_binary.TypeID([8]byte{186, 145, 195, 227, 207, 211, 226, 134})

	// Executes a message on the destination chain.
	//
	// The method name needs to be execute with Anchor encoding.
	//
	// This function is called by the OffChain when executing one Report to the SVM Router.
	// In this Flow only one message is sent, the Execution Report. This is different as EVM does,
	// this is because there is no try/catch mechanism to allow batch execution.
	// This message validates that the Merkle Tree Proof of the given message is correct and is stored in the Commit Report Account.
	// The message must be untouched to be executed.
	// This message emits the event ExecutionStateChanged with the new state of the message.
	// Finally, executes the CPI instruction to the receiver program in the ccip_receive message.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for the execute.
	// * `raw_execution_report` - the serialized execution report containing only one message and proofs
	// * `report_context_byte_words` - report_context after execution_report to match context for manually execute (proper decoding order)
	// *  consists of:
	// * report_context_byte_words[0]: ConfigDigest
	// * report_context_byte_words[1]: 24 byte padding, 8 byte sequence number
	Instruction_Execute = ag_binary.TypeID([8]byte{130, 221, 242, 154, 13, 193, 189, 29})

	// Manually executes a report to the router.
	//
	// When a message is not being executed, then the user can trigger the execution manually.
	// No verification over the transmitter, but the message needs to be in some commit report.
	// It validates that the required time has passed since the commit and then executes the report.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for the execution.
	// * `raw_execution_report` - The serialized execution report containing the message and proofs.
	Instruction_ManuallyExecute = ag_binary.TypeID([8]byte{238, 219, 224, 11, 226, 248, 47, 192})

	Instruction_CloseCommitReportAccount = ag_binary.TypeID([8]byte{109, 145, 129, 64, 226, 172, 61, 106})
)

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id ag_binary.TypeID) string {
	switch id {
	case Instruction_Initialize:
		return "Initialize"
	case Instruction_InitializeConfig:
		return "InitializeConfig"
	case Instruction_TransferOwnership:
		return "TransferOwnership"
	case Instruction_AcceptOwnership:
		return "AcceptOwnership"
	case Instruction_SetDefaultCodeVersion:
		return "SetDefaultCodeVersion"
	case Instruction_UpdateReferenceAddresses:
		return "UpdateReferenceAddresses"
	case Instruction_AddSourceChain:
		return "AddSourceChain"
	case Instruction_DisableSourceChainSelector:
		return "DisableSourceChainSelector"
	case Instruction_UpdateSourceChainConfig:
		return "UpdateSourceChainConfig"
	case Instruction_UpdateSvmChainSelector:
		return "UpdateSvmChainSelector"
	case Instruction_UpdateEnableManualExecutionAfter:
		return "UpdateEnableManualExecutionAfter"
	case Instruction_SetOcrConfig:
		return "SetOcrConfig"
	case Instruction_Commit:
		return "Commit"
	case Instruction_CommitPriceOnly:
		return "CommitPriceOnly"
	case Instruction_Execute:
		return "Execute"
	case Instruction_ManuallyExecute:
		return "ManuallyExecute"
	case Instruction_CloseCommitReportAccount:
		return "CloseCommitReportAccount"
	default:
		return ""
	}
}

type Instruction struct {
	ag_binary.BaseVariant
}

func (inst *Instruction) EncodeToTree(parent ag_treeout.Branches) {
	if enToTree, ok := inst.Impl.(ag_text.EncodableToTree); ok {
		enToTree.EncodeToTree(parent)
	} else {
		parent.Child(ag_spew.Sdump(inst))
	}
}

var InstructionImplDef = ag_binary.NewVariantDefinition(
	ag_binary.AnchorTypeIDEncoding,
	[]ag_binary.VariantType{
		{
			"initialize", (*Initialize)(nil),
		},
		{
			"initialize_config", (*InitializeConfig)(nil),
		},
		{
			"transfer_ownership", (*TransferOwnership)(nil),
		},
		{
			"accept_ownership", (*AcceptOwnership)(nil),
		},
		{
			"set_default_code_version", (*SetDefaultCodeVersion)(nil),
		},
		{
			"update_reference_addresses", (*UpdateReferenceAddresses)(nil),
		},
		{
			"add_source_chain", (*AddSourceChain)(nil),
		},
		{
			"disable_source_chain_selector", (*DisableSourceChainSelector)(nil),
		},
		{
			"update_source_chain_config", (*UpdateSourceChainConfig)(nil),
		},
		{
			"update_svm_chain_selector", (*UpdateSvmChainSelector)(nil),
		},
		{
			"update_enable_manual_execution_after", (*UpdateEnableManualExecutionAfter)(nil),
		},
		{
			"set_ocr_config", (*SetOcrConfig)(nil),
		},
		{
			"commit", (*Commit)(nil),
		},
		{
			"commit_price_only", (*CommitPriceOnly)(nil),
		},
		{
			"execute", (*Execute)(nil),
		},
		{
			"manually_execute", (*ManuallyExecute)(nil),
		},
		{
			"close_commit_report_account", (*CloseCommitReportAccount)(nil),
		},
	},
)

func (inst *Instruction) ProgramID() ag_solanago.PublicKey {
	return ProgramID
}

func (inst *Instruction) Accounts() (out []*ag_solanago.AccountMeta) {
	return inst.Impl.(ag_solanago.AccountsGettable).GetAccounts()
}

func (inst *Instruction) Data() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := ag_binary.NewBorshEncoder(buf).Encode(inst); err != nil {
		return nil, fmt.Errorf("unable to encode instruction: %w", err)
	}
	return buf.Bytes(), nil
}

func (inst *Instruction) TextEncode(encoder *ag_text.Encoder, option *ag_text.Option) error {
	return encoder.Encode(inst.Impl, option)
}

func (inst *Instruction) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	return inst.BaseVariant.UnmarshalBinaryVariant(decoder, InstructionImplDef)
}

func (inst *Instruction) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	err := encoder.WriteBytes(inst.TypeID.Bytes(), false)
	if err != nil {
		return fmt.Errorf("unable to write variant type: %w", err)
	}
	return encoder.Encode(inst.Impl)
}

func registryDecodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (interface{}, error) {
	inst, err := DecodeInstruction(accounts, data)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func DecodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (*Instruction, error) {
	inst := new(Instruction)
	if err := ag_binary.NewBorshDecoder(data).Decode(inst); err != nil {
		return nil, fmt.Errorf("unable to decode instruction: %w", err)
	}
	if v, ok := inst.Impl.(ag_solanago.AccountsSettable); ok {
		err := v.SetAccounts(accounts)
		if err != nil {
			return nil, fmt.Errorf("unable to set accounts for instruction: %w", err)
		}
	}
	return inst, nil
}
