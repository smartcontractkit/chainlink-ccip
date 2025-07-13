package contracts

import (
	"fmt"
	"testing"
	"time"

	"github.com/gagliardetto/solana-go"
	addresslookuptable "github.com/gagliardetto/solana-go/programs/address-lookup-table"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/testutils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
)

func TestSVMLookupTables(t *testing.T) {
	t.Parallel()

	ctx := tests.Context(t)
	url := testutils.SetupLocalSolNode(t)
	c := rpc.New(url)

	sender, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)
	testutils.FundAccounts(ctx, []solana.PrivateKey{sender}, c, t)

	// transfer instructions
	pubkeys := solana.PublicKeySlice{}
	instructions := []solana.Instruction{}
	for i := 0; i < 32; i++ {
		k, kerr := solana.NewRandomPrivateKey()
		require.NoError(t, kerr)
		pubkeys = append(pubkeys, k.PublicKey())
		instructions = append(instructions, system.NewTransferInstruction(1_000_000_000, sender.PublicKey(), k.PublicKey()).Build())
	}

	loadLookupTable := func(k []solana.PublicKey) solana.PublicKey {
		// create lookup table
		slot, serr := c.GetSlot(ctx, rpc.CommitmentFinalized)
		require.NoError(t, serr)
		table, instruction, ierr := common.NewCreateLookupTableInstruction(
			sender.PublicKey(),
			sender.PublicKey(),
			slot,
		)
		require.NoError(t, ierr)
		testutils.SendAndConfirm(ctx, t, c, []solana.Instruction{instruction}, sender, rpc.CommitmentConfirmed)

		// add entries to lookup table
		testutils.SendAndConfirm(ctx, t, c, []solana.Instruction{
			common.NewExtendLookupTableInstruction(
				table, sender.PublicKey(), sender.PublicKey(),
				k,
			),
		}, sender, rpc.CommitmentConfirmed)

		return table
	}

	// use two separate tables + leave additional addresses that are not in tables
	table0 := loadLookupTable(pubkeys[0:15])
	table1 := loadLookupTable(append(pubkeys[15:30], solana.SystemProgramID)) // the program that is called is not used from a lookup table

	// fetch lookup table
	t0data, err := addresslookuptable.GetAddressLookupTableStateWithOpts(ctx, c, table0, &rpc.GetAccountInfoOpts{
		Commitment: rpc.CommitmentConfirmed,
	})
	require.NoError(t, err)
	t1data, err := addresslookuptable.GetAddressLookupTableStateWithOpts(ctx, c, table1, &rpc.GetAccountInfoOpts{
		Commitment: rpc.CommitmentConfirmed,
	})
	require.NoError(t, err)

	// tx should fail without lookup tables (too large)
	blockhash, err := c.GetLatestBlockhash(ctx, rpc.CommitmentFinalized)
	require.NoError(t, err)
	tx, err := solana.NewTransaction(
		instructions,
		blockhash.Value.Blockhash,
		solana.TransactionPayer(sender.PublicKey()),
	)
	require.NoError(t, err)
	_, err = tx.Sign(func(_ solana.PublicKey) *solana.PrivateKey {
		return &sender
	})
	require.NoError(t, err)
	_, err = c.SendTransactionWithOpts(ctx, tx, rpc.TransactionOpts{
		PreflightCommitment: rpc.CommitmentProcessed,
	})
	require.ErrorContains(t, err, "VersionedTransaction too large: 2312 bytes")

	// tx should not fail with lookup tables (874 bytes) - takes a few moments for the tables to load on state?
	err = fmt.Errorf("waiting for successful run")
	count := 0
	for err != nil && count < 5 {
		count++
		time.Sleep(time.Second)
		blockhash, err = c.GetLatestBlockhash(ctx, rpc.CommitmentFinalized)
		require.NoError(t, err)
		tx, err = solana.NewTransaction(
			instructions,
			blockhash.Value.Blockhash,
			solana.TransactionAddressTables(map[solana.PublicKey]solana.PublicKeySlice{
				table0: t0data.Addresses,
				table1: t1data.Addresses,
			}),
			solana.TransactionPayer(sender.PublicKey()),
		)
		require.NoError(t, err)
		_, err = tx.Sign(func(_ solana.PublicKey) *solana.PrivateKey {
			return &sender
		})
		require.NoError(t, err)
		if _, err = c.SendTransactionWithOpts(ctx, tx, rpc.TransactionOpts{
			PreflightCommitment: rpc.CommitmentProcessed,
		}); err != nil {
			t.Log("failed: retrying in 1s")
		}
	}
	require.NoError(t, err)
}
