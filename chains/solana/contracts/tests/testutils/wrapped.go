package testutils

import (
	"context"
	"errors"
	"fmt"
	"testing"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	computebudget "github.com/gagliardetto/solana-go/programs/compute-budget"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/ccip_offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/ccip"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/fees"
)

// this files includes wrapped methods to be used in testing without additional error checks
// this is used to keep a consistent interface to introduce less code churn in the tests

func retryOnTimeouts(maxAttempts int, fn func() (*rpc.GetTransactionResult, error)) (res *rpc.GetTransactionResult, err error) {
	for i := range maxAttempts {
		res, err = fn()
		if err == nil || !errors.Is(err, &common.FindTxTimeoutError{}) {
			return res, err
		}
		// we've seen on CI that sometimes the transaction is not found due to a timeout error because some
		// underlying tool (likely the local validator there) is a little unreliable, so we retry
		fmt.Printf("Retrying due to timeout error: %v (attempt %d out of %d)\n", err, i+1, maxAttempts)
	}
	return nil, fmt.Errorf("max attempts reached: %w", err)
}

func SendAndConfirm(ctx context.Context, t *testing.T, rpcClient *rpc.Client, instructions []solana.Instruction,
	signer solana.PrivateKey, commitment rpc.CommitmentType, opts ...common.TxModifier) *rpc.GetTransactionResult {
	res, err := retryOnTimeouts(3, func() (*rpc.GetTransactionResult, error) {
		return common.SendAndConfirm(ctx, rpcClient, instructions, signer, commitment, opts...)
	})

	if err != nil {
		fmt.Printf("Transaction failed with error: %v\n", err)
		tx, terr := res.Transaction.GetTransaction()
		require.NoError(t, terr)
		fmt.Printf("Transaction details: %v\n", tx)
		for _, l := range res.Meta.LogMessages {
			fmt.Printf("    %s\n", l)
		}
	}
	require.NoError(t, err)

	return res
}

func SendAndConfirmWithLookupTables(ctx context.Context, t *testing.T, rpcClient *rpc.Client, instructions []solana.Instruction,
	signer solana.PrivateKey, commitment rpc.CommitmentType, lookupTables map[solana.PublicKey]solana.PublicKeySlice, opts ...common.TxModifier) *rpc.GetTransactionResult {
	res, err := retryOnTimeouts(3, func() (*rpc.GetTransactionResult, error) {
		return common.SendAndConfirmWithLookupTables(ctx, rpcClient, instructions, signer, commitment, lookupTables, opts...)
	})

	if err != nil {
		fmt.Printf("Transaction failed with error: %v\n", err)
		tx, terr := res.Transaction.GetTransaction()
		require.NoError(t, terr)
		fmt.Printf("Transaction details: %v\n", tx)
		for _, l := range res.Meta.LogMessages {
			fmt.Printf("    %s\n", l)
		}
	}
	require.NoError(t, err)

	return res
}

func SendAndFailWith(ctx context.Context, t *testing.T, rpcClient *rpc.Client, instructions []solana.Instruction,
	signer solana.PrivateKey, commitment rpc.CommitmentType, expectedErrors []string, opts ...common.TxModifier) *rpc.GetTransactionResult {
	res, err := common.SendAndFailWith(ctx, rpcClient, instructions, signer, commitment, expectedErrors, opts...)
	require.NoError(t, err)

	return res
}

func SendAndFailWithLookupTables(ctx context.Context, t *testing.T, rpcClient *rpc.Client, instructions []solana.Instruction,
	signer solana.PrivateKey, commitment rpc.CommitmentType, lookupTables map[solana.PublicKey]solana.PublicKeySlice, expectedErrors []string, opts ...common.TxModifier) *rpc.GetTransactionResult {
	res, err := common.SendAndFailWithLookupTables(ctx, rpcClient, instructions, signer, commitment, lookupTables, expectedErrors, opts...)
	require.NoError(t, err)

	return res
}

func SendAndFailWithRPCError(ctx context.Context, t *testing.T, rpcClient *rpc.Client, instructions []solana.Instruction,
	signer solana.PrivateKey, commitment rpc.CommitmentType, expectedErrors []string, opts ...common.TxModifier) {
	require.NoError(t, common.SendAndFailWithRPCError(ctx, rpcClient, instructions, signer, commitment, expectedErrors, opts...))
}

func SimulateTransaction(ctx context.Context, t *testing.T, rpcClient *rpc.Client, instructions []solana.Instruction, signer solana.PrivateKey) *rpc.SimulateTransactionResponse {
	simRes, err := common.SimulateTransaction(ctx, rpcClient, instructions, signer)
	require.NoError(t, err)

	return simRes
}

func GetRequiredCUWithLookupTables(ctx context.Context, t *testing.T, client *rpc.Client, ixs []solana.Instruction, signer solana.PrivateKey, commitment rpc.CommitmentType, lookupTables map[solana.PublicKey]solana.PublicKeySlice) fees.ComputeUnitLimit {
	// simulate the transaction with max cu to get the required CU
	cuIx, err := computebudget.NewSetComputeUnitLimitInstruction(uint32(computebudget.MAX_COMPUTE_UNIT_LIMIT)).ValidateAndBuild()
	require.NoError(t, err)
	simulateIxs := append([]solana.Instruction{cuIx}, ixs...)
	feeResult, err := common.SimulateTransactionWithLookupTables(ctx, client, simulateIxs, signer,
		lookupTables,
		rpc.SimulateTransactionOpts{
			ReplaceRecentBlockhash: true,
			Commitment:             commitment,
		})
	require.NoError(t, err)
	require.Nil(t, feeResult.Value.Err)
	// maximum cu is 1_400_000, so we can safely cast to uint32
	//nolint:gosec
	return fees.ComputeUnitLimit(*feeResult.Value.UnitsConsumed)
}

func GetRequiredCU(ctx context.Context, t *testing.T, client *rpc.Client, ixs []solana.Instruction, signer solana.PrivateKey, commitment rpc.CommitmentType) fees.ComputeUnitLimit {
	// simulate the transaction with max cu to get the required CU
	cuIx, err := computebudget.NewSetComputeUnitLimitInstruction(uint32(computebudget.MAX_COMPUTE_UNIT_LIMIT)).ValidateAndBuild()
	require.NoError(t, err)
	simulateIxs := append([]solana.Instruction{cuIx}, ixs...)
	feeResult, err := common.SimulateTransactionWithOpts(ctx, client, simulateIxs, signer, rpc.SimulateTransactionOpts{
		ReplaceRecentBlockhash: true,
		Commitment:             commitment,
	})
	require.NoError(t, err)
	require.Nil(t, feeResult.Value.Err)
	// maximum cu is 1_400_000, so we can safely cast to uint32
	//nolint:gosec
	return fees.ComputeUnitLimit(*feeResult.Value.UnitsConsumed)
}

func AssertClosedAccount(ctx context.Context, t *testing.T, solanaGoClient *rpc.Client, accountKey solana.PublicKey, commitment rpc.CommitmentType) {
	isClosed := common.IsClosedAccount(ctx, solanaGoClient, accountKey, commitment)
	require.True(t, isClosed)
}

func CreateNextMessage(ctx context.Context, solanaGoClient *rpc.Client, t *testing.T, remainingAccounts []solana.PublicKey) (ccip_offramp.Any2SVMRampMessage, [32]byte) {
	msg, hash, err := ccip.CreateNextMessage(ctx, solanaGoClient, remainingAccounts)
	require.NoError(t, err)
	return msg, hash
}

func NextSequenceNumber(ctx context.Context, solanaGoClient *rpc.Client, sourceChainStatePDA solana.PublicKey, t *testing.T) uint64 {
	num, err := ccip.NextSequenceNumber(ctx, solanaGoClient, sourceChainStatePDA)
	require.NoError(t, err)
	return num
}

func MakeAnyToSVMMessage(t *testing.T, tokenReceiver solana.PublicKey, evmChainSelector uint64, solanaChainSelector uint64, data []byte, msgAccounts []solana.PublicKey) (ccip_offramp.Any2SVMRampMessage, [32]byte) {
	msg, hash, err := ccip.MakeAnyToSVMMessage(tokenReceiver, evmChainSelector, solanaChainSelector, data, msgAccounts)
	require.NoError(t, err)
	return msg, hash
}

func MustMarshalBorsh(t *testing.T, v interface{}) []byte {
	bz, err := bin.MarshalBorsh(v)
	require.NoError(t, err)
	return bz
}

func MustSerializeExtraArgs(t *testing.T, data interface{}, tag string) []byte {
	b, err := ccip.SerializeExtraArgs(data, tag)
	require.NoError(t, err)
	return b
}

func MustDeserializeExtraArgs[A any](t *testing.T, obj A, data []byte, tag string) A {
	require.NoError(t, ccip.DeserializeExtraArgs(obj, data, tag))
	return obj
}
