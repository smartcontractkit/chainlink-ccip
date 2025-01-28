package common

import (
	"context"
	"encoding/binary"
	"time"

	"github.com/gagliardetto/solana-go"
	addresslookuptable "github.com/gagliardetto/solana-go/programs/address-lookup-table"
	"github.com/gagliardetto/solana-go/rpc"
)

var (
	AddressLookupTableProgram = solana.MustPublicKeyFromBase58("AddressLookupTab1e1111111111111111111111111")
)

// https://github.com/anza-xyz/agave/blob/master/programs/address-lookup-table/src/processor.rs
// https://github.com/anza-xyz/agave/blob/489f483e1d7b30ef114e0123994818b2accfa389/sdk/program/src/address_lookup_table/instruction.rs
const (
	InstructionCreateLookupTable uint32 = iota
	InstructionFreezeLookupTable
	InstructionExtendLookupTable
	InstructionDeactivateLookupTable
	InstructionCloseLookupTable
)

func NewCreateLookupTableInstruction(
	authority,
	funder solana.PublicKey,
	slot uint64,
) (solana.PublicKey, solana.Instruction, error) {
	account, bumpSeed, err := FindLookupTablePDA(authority, slot)
	if err != nil {
		return solana.PublicKey{}, nil, err
	}

	data := binary.LittleEndian.AppendUint32([]byte{}, InstructionCreateLookupTable)
	data = binary.LittleEndian.AppendUint64(data, slot)
	data = append(data, bumpSeed)
	return account, solana.NewInstruction(
		AddressLookupTableProgram,
		solana.AccountMetaSlice{
			solana.Meta(account).WRITE(),
			solana.Meta(authority).SIGNER(),
			solana.Meta(funder).SIGNER().WRITE(),
			solana.Meta(solana.SystemProgramID),
		},
		data,
	), nil
}

func NewExtendLookupTableInstruction(
	table, authority, funder solana.PublicKey,
	accounts []solana.PublicKey,
) solana.Instruction {
	// https://github.com/solana-labs/solana-web3.js/blob/c1c98715b0c7900ce37c59bffd2056fa0037213d/src/programs/address-lookup-table/index.ts#L113

	data := binary.LittleEndian.AppendUint32([]byte{}, InstructionExtendLookupTable)
	data = binary.LittleEndian.AppendUint64(data, uint64(len(accounts))) // note: this is usually u32 + 8 byte buffer
	for _, a := range accounts {
		data = append(data, a.Bytes()...)
	}

	return solana.NewInstruction(
		AddressLookupTableProgram,
		solana.AccountMetaSlice{
			solana.Meta(table).WRITE(),
			solana.Meta(authority).SIGNER(),
			solana.Meta(funder).SIGNER().WRITE(),
			solana.Meta(solana.SystemProgramID),
		},
		data,
	)
}

func CreateLookupTable(ctx context.Context, client *rpc.Client, admin solana.PrivateKey) (solana.PublicKey, error) {
	slot, serr := client.GetSlot(ctx, rpc.CommitmentFinalized)
	if serr != nil {
		return solana.PublicKey{}, serr
	}

	table, instruction, ierr := NewCreateLookupTableInstruction(
		admin.PublicKey(),
		admin.PublicKey(),
		slot-1, // Using the most recent slot sometimes results in errors when submitting the transaction.
	)
	if ierr != nil {
		return solana.PublicKey{}, ierr
	}

	_, err := SendAndConfirm(ctx, client, []solana.Instruction{instruction}, admin, rpc.CommitmentConfirmed)
	return table, err
}

func ExtendLookupTable(ctx context.Context, client *rpc.Client, table solana.PublicKey, admin solana.PrivateKey, entries []solana.PublicKey) error {
	_, err := SendAndConfirm(ctx, client, []solana.Instruction{
		NewExtendLookupTableInstruction(
			table,
			admin.PublicKey(),
			admin.PublicKey(),
			entries,
		),
	}, admin, rpc.CommitmentConfirmed)
	return err
}

func AwaitSlotChange(ctx context.Context, client *rpc.Client) error {
	originalSlot, err := client.GetSlot(ctx, rpc.CommitmentConfirmed)
	if err != nil {
		return err
	}
	newSlot := originalSlot
	for newSlot == originalSlot {
		newSlot, err = client.GetSlot(ctx, rpc.CommitmentConfirmed)
		if err != nil {
			return err
		}
		time.Sleep(10 * time.Millisecond)
	}
	return nil
}

func SetupLookupTable(ctx context.Context, client *rpc.Client, admin solana.PrivateKey, entries []solana.PublicKey) (solana.PublicKey, error) {
	table, err := CreateLookupTable(ctx, client, admin)
	if err != nil {
		return solana.PublicKey{}, err
	}

	err = ExtendLookupTable(ctx, client, table, admin, entries)
	if err != nil {
		return solana.PublicKey{}, err
	}

	// Address lookup tables have to "warm up" for at least 1 slot before they can be used.
	// So, we wait for a new slot to be produced before returning the table, so it's available
	// and warmed up as soon as this method returns it.
	err = AwaitSlotChange(ctx, client)
	if err != nil {
		return solana.PublicKey{}, err
	}

	return table, nil
}

func GetAddressLookupTable(ctx context.Context, client *rpc.Client, lookupTablePublicKey solana.PublicKey) ([]solana.PublicKey, error) {
	lookupTableState, err := addresslookuptable.GetAddressLookupTableStateWithOpts(ctx, client, lookupTablePublicKey, &rpc.GetAccountInfoOpts{
		Commitment: rpc.CommitmentConfirmed,
	})
	if err != nil {
		return []solana.PublicKey{}, err
	}

	return lookupTableState.Addresses, nil
}

// https://github.com/solana-labs/solana-web3.js/blob/c1c98715b0c7900ce37c59bffd2056fa0037213d/src/programs/address-lookup-table/index.ts#L274
func FindLookupTablePDA(authority solana.PublicKey, slot uint64) (solana.PublicKey, uint8, error) {
	slotLE := make([]byte, 8)
	binary.LittleEndian.PutUint64(slotLE, slot)
	return solana.FindProgramAddress([][]byte{authority.Bytes(), slotLE}, AddressLookupTableProgram)
}
