// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_pool

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

const ProgramName = "TokenPool"

func init() {
	if !ProgramID.IsZero() {
		ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
	}
}

var (
	Instruction_Initialize = ag_binary.TypeID([8]byte{175, 175, 109, 31, 13, 152, 155, 237})

	Instruction_TransferOwnership = ag_binary.TypeID([8]byte{65, 177, 215, 73, 53, 45, 99, 47})

	Instruction_AcceptOwnership = ag_binary.TypeID([8]byte{172, 23, 43, 13, 238, 213, 85, 150})

	Instruction_SetRampAuthority = ag_binary.TypeID([8]byte{181, 180, 204, 162, 156, 188, 239, 153})

	Instruction_SetChainRemoteConfig = ag_binary.TypeID([8]byte{147, 161, 6, 246, 121, 59, 37, 28})

	Instruction_SetChainRateLimit = ag_binary.TypeID([8]byte{188, 188, 161, 37, 100, 249, 123, 170})

	Instruction_DeleteChainConfig = ag_binary.TypeID([8]byte{241, 159, 142, 210, 64, 173, 77, 179})

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
	case Instruction_SetRampAuthority:
		return "SetRampAuthority"
	case Instruction_SetChainRemoteConfig:
		return "SetChainRemoteConfig"
	case Instruction_SetChainRateLimit:
		return "SetChainRateLimit"
	case Instruction_DeleteChainConfig:
		return "DeleteChainConfig"
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
			"set_ramp_authority", (*SetRampAuthority)(nil),
		},
		{
			"set_chain_remote_config", (*SetChainRemoteConfig)(nil),
		},
		{
			"set_chain_rate_limit", (*SetChainRateLimit)(nil),
		},
		{
			"delete_chain_config", (*DeleteChainConfig)(nil),
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