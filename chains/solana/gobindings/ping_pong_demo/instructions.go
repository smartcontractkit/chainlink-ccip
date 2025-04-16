// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ping_pong_demo

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

const ProgramName = "PingPongDemo"

func init() {
	if !ProgramID.IsZero() {
		ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
	}
}

var (
	Instruction_InitializeConfig = ag_binary.TypeID([8]byte{208, 127, 21, 1, 194, 190, 196, 70})

	Instruction_Initialize = ag_binary.TypeID([8]byte{175, 175, 109, 31, 13, 152, 155, 237})

	// Returns the program type (name) and version.
	// Used by offchain code to easily determine which program & version is being interacted with.
	//
	// # Arguments
	// * `ctx`` - The context, which contains no accounts.
	Instruction_TypeVersion = ag_binary.TypeID([8]byte{129, 251, 8, 243, 122, 229, 252, 164})

	Instruction_SetCounterpart = ag_binary.TypeID([8]byte{118, 28, 243, 127, 218, 176, 228, 228})

	Instruction_SetPaused = ag_binary.TypeID([8]byte{91, 60, 125, 192, 176, 225, 166, 218})

	Instruction_SetExtraArgs = ag_binary.TypeID([8]byte{103, 87, 237, 252, 141, 176, 81, 193})

	Instruction_StartPingPong = ag_binary.TypeID([8]byte{53, 36, 169, 135, 221, 239, 52, 103})

	Instruction_CcipReceive = ag_binary.TypeID([8]byte{11, 244, 9, 249, 44, 83, 47, 245})
)

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id ag_binary.TypeID) string {
	switch id {
	case Instruction_InitializeConfig:
		return "InitializeConfig"
	case Instruction_Initialize:
		return "Initialize"
	case Instruction_TypeVersion:
		return "TypeVersion"
	case Instruction_SetCounterpart:
		return "SetCounterpart"
	case Instruction_SetPaused:
		return "SetPaused"
	case Instruction_SetExtraArgs:
		return "SetExtraArgs"
	case Instruction_StartPingPong:
		return "StartPingPong"
	case Instruction_CcipReceive:
		return "CcipReceive"
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
			"initialize_config", (*InitializeConfig)(nil),
		},
		{
			"initialize", (*Initialize)(nil),
		},
		{
			"type_version", (*TypeVersion)(nil),
		},
		{
			"set_counterpart", (*SetCounterpart)(nil),
		},
		{
			"set_paused", (*SetPaused)(nil),
		},
		{
			"set_extra_args", (*SetExtraArgs)(nil),
		},
		{
			"start_ping_pong", (*StartPingPong)(nil),
		},
		{
			"ccip_receive", (*CcipReceive)(nil),
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
