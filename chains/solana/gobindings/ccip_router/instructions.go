// The `ccip_router` module contains the implementation of the Cross-Chain Interoperability Protocol (CCIP) Router.
//
// This is the Collapsed Router Program for CCIP.
// As it's upgradable persisting the same program id, there is no need to have an indirection of a Proxy Program.
// This Router handles the OnRamp flow of the CCIP Messages.
//
// NOTE to devs: This file however should contain *no logic*, only the entrypoints to the different versioned modules,
// thus making it easier to ensure later on that logic can be changed during upgrades without affecting the interface.
// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ccip_router

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

const ProgramName = "CcipRouter"

func init() {
	if !ProgramID.IsZero() {
		ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
	}
}

var (
	// Initialization Flow //
	// Initializes the CCIP Router.
	//
	// The initialization of the Router is responsibility of Admin, nothing more than calling this method should be done first.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for initialization.
	// * `svm_chain_selector` - The chain selector for SVM.
	// * `fee_aggregator` - The public key of the fee aggregator.
	// * `fee_quoter` - The public key of the fee quoter.
	// * `link_token_mint` - The public key of the LINK token mint.
	// * `rmn_remote` - The public key of the RMN remote.
	Instruction_Initialize = ag_binary.TypeID([8]byte{175, 175, 109, 31, 13, 152, 155, 237})

	// Print commit SHA
	Instruction_GitCommit = ag_binary.TypeID([8]byte{24, 10, 239, 86, 212, 22, 126, 255})

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

	// Config //
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

	// Updates the fee aggregator in the router configuration.
	// The Admin is the only one able to update the fee aggregator.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for updating the configuration.
	// * `fee_aggregator` - The new fee aggregator address (ATAs will be derived for it for each token).
	Instruction_UpdateFeeAggregator = ag_binary.TypeID([8]byte{85, 112, 115, 60, 22, 95, 230, 56})

	// Updates the RMN remote program in the router configuration.
	// The Admin is the only one able to update the RMN remote program.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for updating the configuration.
	// * `rmn_remote,` - The new RMN remote address.
	Instruction_UpdateRmnRemote = ag_binary.TypeID([8]byte{66, 12, 215, 147, 14, 176, 55, 214})

	// Adds a new chain selector to the router.
	//
	// The Admin needs to add any new chain supported (this means both OnRamp and OffRamp).
	// When adding a new chain, the Admin needs to specify if it's enabled or not.
	// They may enable only source, or only destination, or neither, or both.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for adding the chain selector.
	// * `new_chain_selector` - The new chain selector to be added.
	// * `source_chain_config` - The configuration for the chain as source.
	// * `dest_chain_config` - The configuration for the chain as destination.
	Instruction_AddChainSelector = ag_binary.TypeID([8]byte{28, 60, 171, 0, 195, 113, 56, 7})

	// Updates the configuration of the destination chain selector.
	//
	// The Admin is the only one able to update the destination chain config.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for updating the chain selector.
	// * `dest_chain_selector` - The destination chain selector to be updated.
	// * `dest_chain_config` - The new configuration for the destination chain.
	Instruction_UpdateDestChainConfig = ag_binary.TypeID([8]byte{215, 122, 81, 22, 190, 58, 219, 13})

	// Add an offramp address to the list of offramps allowed by the router, for a
	// particular source chain. External users will check this list before accepting
	// a `ccip_receive` CPI.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for this operation.
	// * `source_chain_selector` - The source chain for the offramp's lane.
	// * `offramp` - The offramp's address.
	Instruction_AddOfframp = ag_binary.TypeID([8]byte{164, 255, 154, 96, 204, 239, 24, 2})

	// Remove an offramp address from the list of offramps allowed by the router, for a
	// particular source chain. External users will check this list before accepting
	// a `ccip_receive` CPI.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for this operation.
	// * `source_chain_selector` - The source chain for the offramp's lane.
	// * `offramp` - The offramp's address.
	Instruction_RemoveOfframp = ag_binary.TypeID([8]byte{252, 152, 51, 170, 241, 13, 199, 8})

	// Updates the SVM chain selector in the router configuration.
	//
	// This method should only be used if there was an error with the initial configuration or if the solana chain selector changes.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for updating the configuration.
	// * `new_chain_selector` - The new chain selector for SVM.
	Instruction_UpdateSvmChainSelector = ag_binary.TypeID([8]byte{164, 212, 71, 101, 166, 113, 26, 93})

	// Bumps the CCIP version for a destination chain.
	// This effectively just resets the sequence number of the destination chain state.
	// If there had been a previous rollback, on re-upgrade the sequence number will resume from where it was
	// prior to the rollback.
	//
	// # Arguments
	// * `ctx` - The context containing the accounts required for the bump.
	// * `dest_chain_selector` - The destination chain selector to bump version for.
	Instruction_BumpCcipVersionForDestChain = ag_binary.TypeID([8]byte{120, 25, 6, 201, 42, 224, 235, 187})

	// Rolls back the CCIP version for a destination chain.
	// This effectively just restores the old version's sequence number of the destination chain state.
	// We only support 1 consecutive rollback. If a rollback has occurred for that lane, the version can't
	// be rolled back again without bumping the version first.
	//
	// # Arguments
	// * `ctx` - The context containing the accounts required for the rollback.
	// * `dest_chain_selector` - The destination chain selector to rollback the version for.
	Instruction_RollbackCcipVersionForDestChain = ag_binary.TypeID([8]byte{95, 107, 33, 138, 26, 57, 154, 110})

	// Token Admin Registry //
	// Registers the Token Admin Registry via the CCIP Admin
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for registration.
	// * `token_admin_registry_admin` - The public key of the token admin registry admin to propose.
	Instruction_CcipAdminProposeAdministrator = ag_binary.TypeID([8]byte{218, 37, 139, 107, 142, 228, 51, 219})

	// Overrides the pending admin of the Token Admin Registry
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for registration.
	// * `token_admin_registry_admin` - The public key of the token admin registry admin to propose.
	Instruction_CcipAdminOverridePendingAdministrator = ag_binary.TypeID([8]byte{163, 206, 164, 199, 248, 92, 36, 46})

	// Registers the Token Admin Registry by the token owner.
	//
	// The Authority of the Mint Token can claim the registry of the token.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for registration.
	// * `token_admin_registry_admin` - The public key of the token admin registry admin to propose.
	Instruction_OwnerProposeAdministrator = ag_binary.TypeID([8]byte{175, 81, 160, 246, 206, 132, 18, 22})

	// Overrides the pending admin of the Token Admin Registry by the token owner
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for registration.
	// * `token_admin_registry_admin` - The public key of the token admin registry admin to propose.
	Instruction_OwnerOverridePendingAdministrator = ag_binary.TypeID([8]byte{230, 111, 134, 149, 203, 168, 118, 201})

	// Accepts the admin role of the token admin registry.
	//
	// The Pending Admin must call this function to accept the admin role of the Token Admin Registry.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for accepting the admin role.
	// * `mint` - The public key of the token mint.
	Instruction_AcceptAdminRoleTokenAdminRegistry = ag_binary.TypeID([8]byte{106, 240, 16, 173, 137, 213, 163, 246})

	// Transfers the admin role of the token admin registry to a new admin.
	//
	// Only the Admin can transfer the Admin Role of the Token Admin Registry, this setups the Pending Admin and then it's their responsibility to accept the role.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for the transfer.
	// * `mint` - The public key of the token mint.
	// * `new_admin` - The public key of the new admin.
	Instruction_TransferAdminRoleTokenAdminRegistry = ag_binary.TypeID([8]byte{178, 98, 203, 181, 203, 107, 106, 14})

	// Sets the pool lookup table for a given token mint.
	//
	// The administrator of the token admin registry can set the pool lookup table for a given token mint.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for setting the pool.
	// * `writable_indexes` - a bit map of the indexes of the accounts in lookup table that are writable
	Instruction_SetPool = ag_binary.TypeID([8]byte{119, 30, 14, 180, 115, 225, 167, 238})

	// Billing //
	// Transfers the accumulated billed fees in a particular token to an arbitrary token account.
	// Only the CCIP Admin can withdraw billed funds.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for the transfer of billed fees.
	// * `transfer_all` - A flag indicating whether to transfer all the accumulated fees in that token or not.
	// * `desired_amount` - The amount to transfer. If `transfer_all` is true, this value must be 0.
	Instruction_WithdrawBilledFunds = ag_binary.TypeID([8]byte{16, 116, 73, 38, 77, 232, 6, 28})

	// On Ramp Flow //
	// Sends a message to the destination chain.
	//
	// Request a message to be sent to the destination chain.
	// The method name needs to be ccip_send with Anchor encoding.
	// This function is called by the CCIP Sender Contract (or final user) to send a message to the CCIP Router.
	// The message will be sent to the receiver on the destination chain selector.
	// This message emits the event CCIPMessageSent with all the necessary data to be retrieved by the OffChain Code
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for sending the message.
	// * `dest_chain_selector` - The chain selector for the destination chain.
	// * `message` - The message to be sent. The size limit of data is 256 bytes.
	// * `token_indexes` - Indices into the remaining accounts vector where the subslice for a token begins.
	Instruction_CcipSend = ag_binary.TypeID([8]byte{108, 216, 134, 191, 249, 234, 33, 84})

	// Queries the onramp for the fee required to send a message.
	//
	// This call is permissionless. Note it does not verify whether there's a curse active
	// in order to avoid the RMN CPI overhead.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for obtaining the message fee.
	// * `dest_chain_selector` - The chain selector for the destination chain.
	// * `message` - The message to be sent. The size limit of data is 256 bytes.
	Instruction_GetFee = ag_binary.TypeID([8]byte{115, 195, 235, 161, 25, 219, 60, 29})
)

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id ag_binary.TypeID) string {
	switch id {
	case Instruction_Initialize:
		return "Initialize"
	case Instruction_GitCommit:
		return "GitCommit"
	case Instruction_TransferOwnership:
		return "TransferOwnership"
	case Instruction_AcceptOwnership:
		return "AcceptOwnership"
	case Instruction_SetDefaultCodeVersion:
		return "SetDefaultCodeVersion"
	case Instruction_UpdateFeeAggregator:
		return "UpdateFeeAggregator"
	case Instruction_UpdateRmnRemote:
		return "UpdateRmnRemote"
	case Instruction_AddChainSelector:
		return "AddChainSelector"
	case Instruction_UpdateDestChainConfig:
		return "UpdateDestChainConfig"
	case Instruction_AddOfframp:
		return "AddOfframp"
	case Instruction_RemoveOfframp:
		return "RemoveOfframp"
	case Instruction_UpdateSvmChainSelector:
		return "UpdateSvmChainSelector"
	case Instruction_BumpCcipVersionForDestChain:
		return "BumpCcipVersionForDestChain"
	case Instruction_RollbackCcipVersionForDestChain:
		return "RollbackCcipVersionForDestChain"
	case Instruction_CcipAdminProposeAdministrator:
		return "CcipAdminProposeAdministrator"
	case Instruction_CcipAdminOverridePendingAdministrator:
		return "CcipAdminOverridePendingAdministrator"
	case Instruction_OwnerProposeAdministrator:
		return "OwnerProposeAdministrator"
	case Instruction_OwnerOverridePendingAdministrator:
		return "OwnerOverridePendingAdministrator"
	case Instruction_AcceptAdminRoleTokenAdminRegistry:
		return "AcceptAdminRoleTokenAdminRegistry"
	case Instruction_TransferAdminRoleTokenAdminRegistry:
		return "TransferAdminRoleTokenAdminRegistry"
	case Instruction_SetPool:
		return "SetPool"
	case Instruction_WithdrawBilledFunds:
		return "WithdrawBilledFunds"
	case Instruction_CcipSend:
		return "CcipSend"
	case Instruction_GetFee:
		return "GetFee"
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
			"git_commit", (*GitCommit)(nil),
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
			"update_fee_aggregator", (*UpdateFeeAggregator)(nil),
		},
		{
			"update_rmn_remote", (*UpdateRmnRemote)(nil),
		},
		{
			"add_chain_selector", (*AddChainSelector)(nil),
		},
		{
			"update_dest_chain_config", (*UpdateDestChainConfig)(nil),
		},
		{
			"add_offramp", (*AddOfframp)(nil),
		},
		{
			"remove_offramp", (*RemoveOfframp)(nil),
		},
		{
			"update_svm_chain_selector", (*UpdateSvmChainSelector)(nil),
		},
		{
			"bump_ccip_version_for_dest_chain", (*BumpCcipVersionForDestChain)(nil),
		},
		{
			"rollback_ccip_version_for_dest_chain", (*RollbackCcipVersionForDestChain)(nil),
		},
		{
			"ccip_admin_propose_administrator", (*CcipAdminProposeAdministrator)(nil),
		},
		{
			"ccip_admin_override_pending_administrator", (*CcipAdminOverridePendingAdministrator)(nil),
		},
		{
			"owner_propose_administrator", (*OwnerProposeAdministrator)(nil),
		},
		{
			"owner_override_pending_administrator", (*OwnerOverridePendingAdministrator)(nil),
		},
		{
			"accept_admin_role_token_admin_registry", (*AcceptAdminRoleTokenAdminRegistry)(nil),
		},
		{
			"transfer_admin_role_token_admin_registry", (*TransferAdminRoleTokenAdminRegistry)(nil),
		},
		{
			"set_pool", (*SetPool)(nil),
		},
		{
			"withdraw_billed_funds", (*WithdrawBilledFunds)(nil),
		},
		{
			"ccip_send", (*CcipSend)(nil),
		},
		{
			"get_fee", (*GetFee)(nil),
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
