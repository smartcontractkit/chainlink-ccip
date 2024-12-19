package timelock

import (
	"bytes"

	"github.com/gagliardetto/solana-go"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/timelock"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/mcms"
)

// represents a single instruction with its required accounts
type Instruction struct {
	Ix       solana.Instruction   // instruction to be scheduled / executed
	Accounts []solana.AccountMeta // required accounts for the instruction(should be provided in execute stage)
}

// represents a batch of instructions that having atomicy to be scheduled and executed via timelock
type Operation struct {
	Predecessor  [32]byte      // hashed id of the previous operation
	Salt         [32]byte      // random salt for the operation
	Delay        uint64        // delay in seconds
	instructions []Instruction // instruction data slice, use Add method to add instructions and accounts
}

// add instruction and required accounts to operation
func (op *Operation) AddInstruction(ix solana.Instruction, additionalPrograms []solana.PublicKey) {
	accounts := make([]solana.AccountMeta, len(ix.Accounts()))
	// anchor ix builder doesn't include program
	for _, program := range additionalPrograms {
		accounts = append(accounts, solana.AccountMeta{PublicKey: program, IsSigner: false, IsWritable: false})
	}

	for i, acc := range ix.Accounts() {
		accounts[i] = *acc
		// for debugging
		// fmt.Printf("Account %d: %s (signer: %v, writable: %v)\n",
		// 	i, acc.PublicKey, acc.IsSigner, acc.IsWritable)
	}

	op.instructions = append(op.instructions, Instruction{
		Ix:       ix,
		Accounts: accounts,
	})
}

func (op *Operation) IxsCountU32() uint32 {
	ixsCount, err := mcms.SafeToUint32(len(op.instructions))
	if err != nil {
		panic(err)
	}
	return ixsCount
}

// convert operation to timelock instruction data slice
func (op *Operation) ToInstructionData() []timelock.InstructionData {
	ixs := make([]timelock.InstructionData, len(op.instructions))
	for i, ix := range op.instructions {
		ixData, err := convertToInstructionData(ix.Ix)
		if err != nil {
			panic(err)
		}
		ixs[i] = ixData
	}
	return ixs
}

// get required accounts for the operation
// it merges the required accounts of all instructions and removes duplicates
func (op *Operation) RemainingAccounts() []*solana.AccountMeta {
	accountMap := make(map[string]*solana.AccountMeta)
	for _, instr := range op.instructions {
		for _, acc := range instr.Accounts {
			key := acc.PublicKey.String()
			if existing, exists := accountMap[key]; exists {
				existing.IsWritable = existing.IsWritable || acc.IsWritable
			} else {
				accCopy := acc
				// todo: maybe keep it as it is and override on chain(in case if we gonna calc root from this)
				accCopy.IsSigner = false // force false for CPI
				accountMap[key] = &accCopy
			}
		}
	}

	accounts := make([]*solana.AccountMeta, 0, len(accountMap))
	for _, acc := range accountMap {
		accounts = append(accounts, acc)
	}
	return accounts
}

// hash the operation and return operation id
func (op *Operation) OperationID() [32]byte {
	return hashOperation(op.ToInstructionData(), op.Predecessor, op.Salt)
}

func (op *Operation) OperationPDA() solana.PublicKey {
	id := op.OperationID()
	return config.TimelockOperationPDA(id)
}

// type conversion from solana instruction to timelock instruction data
func convertToInstructionData(ix solana.Instruction) (timelock.InstructionData, error) {
	accounts := make([]timelock.InstructionAccount, len(ix.Accounts()))
	for i, acc := range ix.Accounts() {
		accounts[i] = timelock.InstructionAccount{
			Pubkey:     acc.PublicKey,
			IsSigner:   acc.IsSigner,
			IsWritable: acc.IsWritable,
		}
	}

	data, err := ix.Data()
	if err != nil {
		return timelock.InstructionData{}, err
	}

	return timelock.InstructionData{
		ProgramId: ix.ProgramID(),
		Accounts:  accounts,
		Data:      data,
	}, nil
}

func hashOperation(instructions []timelock.InstructionData, predecessor [32]byte, salt [32]byte) [32]byte {
	var encodedData bytes.Buffer

	for _, ix := range instructions {
		encodedData.Write(ix.ProgramId[:])

		for _, acc := range ix.Accounts {
			encodedData.Write(acc.Pubkey[:])
			if acc.IsSigner {
				encodedData.WriteByte(1)
			} else {
				encodedData.WriteByte(0)
			}
			if acc.IsWritable {
				encodedData.WriteByte(1)
			} else {
				encodedData.WriteByte(0)
			}
		}
		encodedData.Write(ix.Data)
	}

	encodedData.Write(predecessor[:])
	encodedData.Write(salt[:])

	result := mcms.Keccak256(encodedData.Bytes())

	var hash [32]byte
	copy(hash[:], result)

	return hash
}
