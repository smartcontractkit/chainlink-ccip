// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package timelock

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

const ProgramName = "Timelock"

func init() {
	if !ProgramID.IsZero() {
		ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
	}
}

var (
	Instruction_Initialize = ag_binary.TypeID([8]byte{175, 175, 109, 31, 13, 152, 155, 237})

	Instruction_BatchAddAccess = ag_binary.TypeID([8]byte{73, 141, 223, 79, 66, 154, 226, 67})

	Instruction_InitializeOperation = ag_binary.TypeID([8]byte{15, 96, 217, 171, 124, 4, 113, 243})

	Instruction_AppendInstructions = ag_binary.TypeID([8]byte{58, 58, 137, 122, 115, 51, 144, 134})

	Instruction_FinalizeOperation = ag_binary.TypeID([8]byte{63, 208, 32, 98, 85, 182, 236, 140})

	Instruction_ClearOperation = ag_binary.TypeID([8]byte{111, 217, 62, 240, 224, 75, 60, 58})

	Instruction_ScheduleBatch = ag_binary.TypeID([8]byte{242, 140, 87, 106, 71, 226, 86, 32})

	Instruction_Cancel = ag_binary.TypeID([8]byte{232, 219, 223, 41, 219, 236, 220, 190})

	Instruction_ExecuteBatch = ag_binary.TypeID([8]byte{112, 159, 211, 51, 238, 70, 212, 60})

	Instruction_InitializeBypasserOperation = ag_binary.TypeID([8]byte{58, 27, 48, 204, 19, 197, 63, 26})

	Instruction_AppendBypasserInstructions = ag_binary.TypeID([8]byte{127, 68, 8, 210, 106, 213, 25, 215})

	Instruction_FinalizeBypasserOperation = ag_binary.TypeID([8]byte{45, 55, 198, 51, 124, 24, 169, 250})

	Instruction_ClearBypasserOperation = ag_binary.TypeID([8]byte{200, 21, 249, 130, 56, 13, 128, 32})

	Instruction_BypasserExecuteBatch = ag_binary.TypeID([8]byte{90, 62, 66, 6, 227, 174, 30, 194})

	Instruction_UpdateDelay = ag_binary.TypeID([8]byte{164, 186, 80, 62, 85, 88, 182, 147})

	Instruction_BlockFunctionSelector = ag_binary.TypeID([8]byte{119, 89, 101, 41, 72, 143, 218, 185})

	Instruction_UnblockFunctionSelector = ag_binary.TypeID([8]byte{53, 84, 245, 196, 149, 52, 30, 57})

	Instruction_TransferOwnership = ag_binary.TypeID([8]byte{65, 177, 215, 73, 53, 45, 99, 47})

	Instruction_AcceptOwnership = ag_binary.TypeID([8]byte{172, 23, 43, 13, 238, 213, 85, 150})
)

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id ag_binary.TypeID) string {
	switch id {
	case Instruction_Initialize:
		return "Initialize"
	case Instruction_BatchAddAccess:
		return "BatchAddAccess"
	case Instruction_InitializeOperation:
		return "InitializeOperation"
	case Instruction_AppendInstructions:
		return "AppendInstructions"
	case Instruction_FinalizeOperation:
		return "FinalizeOperation"
	case Instruction_ClearOperation:
		return "ClearOperation"
	case Instruction_ScheduleBatch:
		return "ScheduleBatch"
	case Instruction_Cancel:
		return "Cancel"
	case Instruction_ExecuteBatch:
		return "ExecuteBatch"
	case Instruction_InitializeBypasserOperation:
		return "InitializeBypasserOperation"
	case Instruction_AppendBypasserInstructions:
		return "AppendBypasserInstructions"
	case Instruction_FinalizeBypasserOperation:
		return "FinalizeBypasserOperation"
	case Instruction_ClearBypasserOperation:
		return "ClearBypasserOperation"
	case Instruction_BypasserExecuteBatch:
		return "BypasserExecuteBatch"
	case Instruction_UpdateDelay:
		return "UpdateDelay"
	case Instruction_BlockFunctionSelector:
		return "BlockFunctionSelector"
	case Instruction_UnblockFunctionSelector:
		return "UnblockFunctionSelector"
	case Instruction_TransferOwnership:
		return "TransferOwnership"
	case Instruction_AcceptOwnership:
		return "AcceptOwnership"
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
			"batch_add_access", (*BatchAddAccess)(nil),
		},
		{
			"initialize_operation", (*InitializeOperation)(nil),
		},
		{
			"append_instructions", (*AppendInstructions)(nil),
		},
		{
			"finalize_operation", (*FinalizeOperation)(nil),
		},
		{
			"clear_operation", (*ClearOperation)(nil),
		},
		{
			"schedule_batch", (*ScheduleBatch)(nil),
		},
		{
			"cancel", (*Cancel)(nil),
		},
		{
			"execute_batch", (*ExecuteBatch)(nil),
		},
		{
			"initialize_bypasser_operation", (*InitializeBypasserOperation)(nil),
		},
		{
			"append_bypasser_instructions", (*AppendBypasserInstructions)(nil),
		},
		{
			"finalize_bypasser_operation", (*FinalizeBypasserOperation)(nil),
		},
		{
			"clear_bypasser_operation", (*ClearBypasserOperation)(nil),
		},
		{
			"bypasser_execute_batch", (*BypasserExecuteBatch)(nil),
		},
		{
			"update_delay", (*UpdateDelay)(nil),
		},
		{
			"block_function_selector", (*BlockFunctionSelector)(nil),
		},
		{
			"unblock_function_selector", (*UnblockFunctionSelector)(nil),
		},
		{
			"transfer_ownership", (*TransferOwnership)(nil),
		},
		{
			"accept_ownership", (*AcceptOwnership)(nil),
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
