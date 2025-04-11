// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package burnmint_token_pool

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

const ProgramName = "BurnmintTokenPool"

func init() {
	if !ProgramID.IsZero() {
		ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
	}
}

var (
	Instruction_Initialize = ag_binary.TypeID([8]byte{175, 175, 109, 31, 13, 152, 155, 237})

	Instruction_TransferOwnership = ag_binary.TypeID([8]byte{65, 177, 215, 73, 53, 45, 99, 47})

	Instruction_AcceptOwnership = ag_binary.TypeID([8]byte{172, 23, 43, 13, 238, 213, 85, 150})

	Instruction_SetRouter = ag_binary.TypeID([8]byte{236, 248, 107, 200, 151, 160, 44, 250})

	Instruction_InitializeStateVersion = ag_binary.TypeID([8]byte{54, 186, 181, 26, 2, 198, 200, 158})

	Instruction_InitChainRemoteConfig = ag_binary.TypeID([8]byte{21, 150, 133, 36, 2, 116, 199, 129})

	Instruction_EditChainRemoteConfig = ag_binary.TypeID([8]byte{149, 112, 186, 72, 116, 217, 159, 175})

	Instruction_AppendRemotePoolAddresses = ag_binary.TypeID([8]byte{172, 57, 83, 55, 70, 112, 26, 197})

	Instruction_SetChainRateLimit = ag_binary.TypeID([8]byte{188, 188, 161, 37, 100, 249, 123, 170})

	Instruction_DeleteChainConfig = ag_binary.TypeID([8]byte{241, 159, 142, 210, 64, 173, 77, 179})

	Instruction_ConfigureAllowList = ag_binary.TypeID([8]byte{18, 180, 102, 187, 209, 0, 130, 191})

	Instruction_RemoveFromAllowList = ag_binary.TypeID([8]byte{44, 46, 123, 213, 40, 11, 107, 18})

	Instruction_ReleaseOrMintTokens = ag_binary.TypeID([8]byte{92, 100, 150, 198, 252, 63, 164, 228})

	Instruction_LockOrBurnTokens = ag_binary.TypeID([8]byte{114, 161, 94, 29, 147, 25, 232, 191})
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
	case Instruction_SetRouter:
		return "SetRouter"
	case Instruction_InitializeStateVersion:
		return "InitializeStateVersion"
	case Instruction_InitChainRemoteConfig:
		return "InitChainRemoteConfig"
	case Instruction_EditChainRemoteConfig:
		return "EditChainRemoteConfig"
	case Instruction_AppendRemotePoolAddresses:
		return "AppendRemotePoolAddresses"
	case Instruction_SetChainRateLimit:
		return "SetChainRateLimit"
	case Instruction_DeleteChainConfig:
		return "DeleteChainConfig"
	case Instruction_ConfigureAllowList:
		return "ConfigureAllowList"
	case Instruction_RemoveFromAllowList:
		return "RemoveFromAllowList"
	case Instruction_ReleaseOrMintTokens:
		return "ReleaseOrMintTokens"
	case Instruction_LockOrBurnTokens:
		return "LockOrBurnTokens"
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
			"set_router", (*SetRouter)(nil),
		},
		{
			"initialize_state_version", (*InitializeStateVersion)(nil),
		},
		{
			"init_chain_remote_config", (*InitChainRemoteConfig)(nil),
		},
		{
			"edit_chain_remote_config", (*EditChainRemoteConfig)(nil),
		},
		{
			"append_remote_pool_addresses", (*AppendRemotePoolAddresses)(nil),
		},
		{
			"set_chain_rate_limit", (*SetChainRateLimit)(nil),
		},
		{
			"delete_chain_config", (*DeleteChainConfig)(nil),
		},
		{
			"configure_allow_list", (*ConfigureAllowList)(nil),
		},
		{
			"remove_from_allow_list", (*RemoveFromAllowList)(nil),
		},
		{
			"release_or_mint_tokens", (*ReleaseOrMintTokens)(nil),
		},
		{
			"lock_or_burn_tokens", (*LockOrBurnTokens)(nil),
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
