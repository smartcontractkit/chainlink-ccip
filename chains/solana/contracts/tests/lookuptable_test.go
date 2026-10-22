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

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/testutils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
)

func TestSVMLookupTables(t *testing.T) {
	ctx := t.Context()
	url := testutils.SetupLocalSolNode(t)
	c := rpc.New(url)

	sender, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)
	testutils.FundAccounts(ctx, []solana.PrivateKey{sender}, c, t)

	// transfer instructions
	pubkeys := solana.PublicKeySlice{}
	instructions := []solana.Instruction{}
	for range 32 {
		k, kerr := solana.NewRandomPrivateKey()
		require.NoError(t, kerr)
		pubkeys = append(pubkeys, k.PublicKey())
		instructions = append(instructions, system.NewTransferInstruction(1_000_000_000, sender.PublicKey(), k.PublicKey()).Build())
	}

	loadLookupTable := func(k []solana.PublicKey) solana.PublicKey {
		// NOTE: the ALT program requires the derivation slot to be present in the
		// slot-hashes sysvar of the bank that processes (and preflight-simulates) the
		// instruction. common.SetupLookupTable handles this by deriving the table from
		// the slot preceding the finalized slot. "not a recent slot" errors can still
		// occur transiently when finalized lags behind the bank tip, so retry (see
		// ccip_router_test.go for the same pattern).
		var table solana.PublicKey
		testutils.RetryIfErrContains(t, 5, "not a recent slot", func() error {
			var err error
			table, err = common.SetupLookupTable(ctx, c, sender, k)
			return err
		})
		return table
	}

	// use two separate tables + leave additional addresses that are not in tables
	table0 := loadLookupTable(pubkeys[0:15])
	table1 := loadLookupTable(append(pubkeys[15:30], solana.SystemProgramID)) // the program that is called is not used from a lookup table

	// fetch lookup table
	t0data, err := addresslookuptable.GetAddressLookupTableStateWithOpts(ctx, c, table0, &rpc.GetAccountInfoOpts{
		Commitment: rpc.CommitmentFinalized,
	})
	require.NoError(t, err)
	t1data, err := addresslookuptable.GetAddressLookupTableStateWithOpts(ctx, c, table1, &rpc.GetAccountInfoOpts{
		Commitment: rpc.CommitmentFinalized,
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
