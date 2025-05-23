// This program an example of a CCIP Receiver Program.
// Used to test CCIP Router execute.
// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package test_ccip_receiver

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

const ProgramName = "TestCcipReceiver"

func init() {
	if !ProgramID.IsZero() {
		ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
	}
}

var (
	// The initialization is responsibility of the External User, CCIP is not handling initialization of Accounts
	Instruction_Initialize = ag_binary.TypeID([8]byte{175, 175, 109, 31, 13, 152, 155, 237})

	Instruction_SetRejectAll = ag_binary.TypeID([8]byte{42, 90, 30, 32, 7, 99, 130, 151})

	// This function is called by the CCIP Router to execute the CCIP message.
	// The method name needs to be ccip_receive with Anchor encoding,
	// if not using Anchor the discriminator needs to be [0x0b, 0xf4, 0x09, 0xf9, 0x2c, 0x53, 0x2f, 0xf5]
	// You can send as many accounts as you need, specifying if mutable or not.
	// But none of them could be an init, realloc or close.
	// In this case, it increments the counter value by 1 and logs the parsed message.
	Instruction_CcipReceive = ag_binary.TypeID([8]byte{11, 244, 9, 249, 44, 83, 47, 245})

	Instruction_CcipTokenReleaseMint = ag_binary.TypeID([8]byte{20, 148, 113, 198, 229, 170, 71, 48})

	Instruction_CcipTokenLockBurn = ag_binary.TypeID([8]byte{200, 14, 50, 9, 44, 91, 121, 37})
)

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id ag_binary.TypeID) string {
	switch id {
	case Instruction_Initialize:
		return "Initialize"
	case Instruction_SetRejectAll:
		return "SetRejectAll"
	case Instruction_CcipReceive:
		return "CcipReceive"
	case Instruction_CcipTokenReleaseMint:
		return "CcipTokenReleaseMint"
	case Instruction_CcipTokenLockBurn:
		return "CcipTokenLockBurn"
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
			"set_reject_all", (*SetRejectAll)(nil),
		},
		{
			"ccip_receive", (*CcipReceive)(nil),
		},
		{
			"ccip_token_release_mint", (*CcipTokenReleaseMint)(nil),
		},
		{
			"ccip_token_lock_burn", (*CcipTokenLockBurn)(nil),
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
