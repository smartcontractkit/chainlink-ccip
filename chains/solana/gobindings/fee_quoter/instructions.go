// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package fee_quoter

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

const ProgramName = "FeeQuoter"

func init() {
	if !ProgramID.IsZero() {
		ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
	}
}

var (
	// Initializes the Fee Quoter.
	//
	// The initialization is responsibility of Admin, nothing more than calling this method should be done first.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for initialization.
	// * `max_fee_juels_per_msg` - The maximum fee in juels that can be charged per message.
	// * `onramp` - The public key of the onramp.
	//
	// The function also uses the link_token_mint account from the context.
	Instruction_Initialize = ag_binary.TypeID([8]byte{175, 175, 109, 31, 13, 152, 155, 237})

	// Transfers the ownership of the fee quoter to a new proposed owner.
	//
	// Shared func signature with other programs
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for the transfer.
	// * `proposed_owner` - The public key of the new proposed owner.
	Instruction_TransferOwnership = ag_binary.TypeID([8]byte{65, 177, 215, 73, 53, 45, 99, 47})

	// Accepts the ownership of the fee quoter by the proposed owner.
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

	// Adds a billing token configuration.
	// Only CCIP Admin can add a billing token configuration.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for adding the billing token configuration.
	// * `config` - The billing token configuration to be added.
	Instruction_AddBillingTokenConfig = ag_binary.TypeID([8]byte{63, 156, 254, 216, 227, 53, 0, 69})

	// Updates the billing token configuration.
	// Only CCIP Admin can update a billing token configuration.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for updating the billing token configuration.
	// * `config` - The new billing token configuration.
	Instruction_UpdateBillingTokenConfig = ag_binary.TypeID([8]byte{140, 184, 124, 146, 204, 62, 244, 79})

	// Adds a new destination chain selector to the fee quoter.
	//
	// The Admin needs to add any new chain supported.
	// When adding a new chain, the Admin needs to specify if it's enabled or not.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for adding the chain selector.
	// * `chain_selector` - The new chain selector to be added.
	// * `dest_chain_config` - The configuration for the chain as destination.
	Instruction_AddDestChain = ag_binary.TypeID([8]byte{122, 202, 174, 155, 55, 100, 102, 36})

	// Disables the destination chain selector.
	//
	// The Admin is the only one able to disable the chain selector as destination. This method is thought of as an emergency kill-switch.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for disabling the chain selector.
	// * `chain_selector` - The destination chain selector to be disabled.
	Instruction_DisableDestChain = ag_binary.TypeID([8]byte{200, 195, 114, 206, 152, 86, 50, 41})

	// Updates the configuration of the destination chain selector.
	//
	// The Admin is the only one able to update the destination chain config.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for updating the chain selector.
	// * `chain_selector` - The destination chain selector to be updated.
	// * `dest_chain_config` - The new configuration for the destination chain.
	Instruction_UpdateDestChainConfig = ag_binary.TypeID([8]byte{215, 122, 81, 22, 190, 58, 219, 13})

	// Sets the token transfer fee configuration for a particular token when it's transferred to a particular dest chain.
	// It is an upsert, initializing the per-chain-per-token config account if it doesn't exist
	// and overwriting it if it does.
	//
	// Only the Admin can perform this operation.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for setting the token billing configuration.
	// * `chain_selector` - The chain selector.
	// * `mint` - The public key of the token mint.
	// * `cfg` - The token transfer fee configuration.
	Instruction_SetTokenTransferFeeConfig = ag_binary.TypeID([8]byte{76, 243, 16, 214, 126, 11, 254, 77})

	// Add a price updater address to the list of allowed price updaters.
	// On price updates, the fee quoter will check the that caller is allowed.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for this operation.
	// * `price_updater` - The price updater address.
	Instruction_AddPriceUpdater = ag_binary.TypeID([8]byte{200, 26, 13, 120, 226, 182, 64, 16})

	// Remove a price updater address from the list of allowed price updaters.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for this operation.
	// * `price_updater` - The price updater address.
	Instruction_RemovePriceUpdater = ag_binary.TypeID([8]byte{10, 61, 172, 48, 110, 8, 162, 198})

	// Calculates the fee for sending a message to the destination chain.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts required for the fee calculation.
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
	// GetFeeResult struct with:
	// - the fee token mint address,
	// - the fee amount of said token,
	// - the fee value in juels,
	// - additional data required when performing the cross-chain transfer of tokens in that message
	// - deserialized and processed extra args
	Instruction_GetFee = ag_binary.TypeID([8]byte{115, 195, 235, 161, 25, 219, 60, 29})

	// Updates prices for tokens and gas. This method may only be called by an allowed price updater.
	//
	// # Arguments
	//
	// * `ctx` - The context containing the accounts always required for the price updates
	// * `token_updates` - Vector of token price updates
	// * `gas_updates` - Vector of gas price updates
	//
	// # Additional accounts
	//
	// In addition to the fixed amount of accounts defined in the `UpdatePrices` context,
	// the following accounts must be provided:
	//
	// * First, the billing token config accounts for each token whose price is being updated, in the same order
	// as the token_updates vector.
	// * Then, the dest chain accounts of every chain whose gas price is being updated, in the same order as the
	// gas_updates vector.
	Instruction_UpdatePrices = ag_binary.TypeID([8]byte{62, 161, 234, 136, 106, 26, 18, 160})
)

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id ag_binary.TypeID) string {
	switch id {
	case Instruction_Initialize:
		return "Initialize"
	case Instruction_TransferOwnership:
		return "TransferOwnership"
	case Instruction_AcceptOwnership:
		return "AcceptOwnership"
	case Instruction_SetDefaultCodeVersion:
		return "SetDefaultCodeVersion"
	case Instruction_AddBillingTokenConfig:
		return "AddBillingTokenConfig"
	case Instruction_UpdateBillingTokenConfig:
		return "UpdateBillingTokenConfig"
	case Instruction_AddDestChain:
		return "AddDestChain"
	case Instruction_DisableDestChain:
		return "DisableDestChain"
	case Instruction_UpdateDestChainConfig:
		return "UpdateDestChainConfig"
	case Instruction_SetTokenTransferFeeConfig:
		return "SetTokenTransferFeeConfig"
	case Instruction_AddPriceUpdater:
		return "AddPriceUpdater"
	case Instruction_RemovePriceUpdater:
		return "RemovePriceUpdater"
	case Instruction_GetFee:
		return "GetFee"
	case Instruction_UpdatePrices:
		return "UpdatePrices"
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
			"transfer_ownership", (*TransferOwnership)(nil),
		},
		{
			"accept_ownership", (*AcceptOwnership)(nil),
		},
		{
			"set_default_code_version", (*SetDefaultCodeVersion)(nil),
		},
		{
			"add_billing_token_config", (*AddBillingTokenConfig)(nil),
		},
		{
			"update_billing_token_config", (*UpdateBillingTokenConfig)(nil),
		},
		{
			"add_dest_chain", (*AddDestChain)(nil),
		},
		{
			"disable_dest_chain", (*DisableDestChain)(nil),
		},
		{
			"update_dest_chain_config", (*UpdateDestChainConfig)(nil),
		},
		{
			"set_token_transfer_fee_config", (*SetTokenTransferFeeConfig)(nil),
		},
		{
			"add_price_updater", (*AddPriceUpdater)(nil),
		},
		{
			"remove_price_updater", (*RemovePriceUpdater)(nil),
		},
		{
			"get_fee", (*GetFee)(nil),
		},
		{
			"update_prices", (*UpdatePrices)(nil),
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
