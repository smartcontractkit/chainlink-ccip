package testutils

import (
	"context"
	"testing"

	"github.com/gagliardetto/solana-go"
	computebudget "github.com/gagliardetto/solana-go/programs/compute-budget"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/ccip"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/fees"
)

// this files includes wrapped methods to be used in testing without additional error checks
// this is used to keep a consistent interface to introduce less code churn in the tests

func SendAndConfirm(ctx context.Context, t *testing.T, rpcClient *rpc.Client, instructions []solana.Instruction, signer solana.PrivateKey, commitment rpc.CommitmentType, opts ...common.TxModifier) *rpc.GetTransactionResult {
	res, err := common.SendAndConfirm(ctx, rpcClient, instructions, signer, commitment, opts...)
	require.NoError(t, err)

	return res
}

func SendAndConfirmWithLookupTables(ctx context.Context, t *testing.T, rpcClient *rpc.Client, instructions []solana.Instruction,
	signer solana.PrivateKey, commitment rpc.CommitmentType, lookupTables map[solana.PublicKey]solana.PublicKeySlice, opts ...common.TxModifier) *rpc.GetTransactionResult {
	res, err := common.SendAndConfirmWithLookupTables(ctx, rpcClient, instructions, signer, commitment, lookupTables, opts...)
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
	signer solana.PrivateKey, commitment rpc.CommitmentType, expectedErrors []string) {
	require.NoError(t, common.SendAndFailWithRPCError(ctx, rpcClient, instructions, signer, commitment, expectedErrors))
}

func SimulateTransaction(ctx context.Context, t *testing.T, rpcClient *rpc.Client, instructions []solana.Instruction, signer solana.PrivateKey) *rpc.SimulateTransactionResponse {
	simRes, err := common.SimulateTransaction(ctx, rpcClient, instructions, signer)
	require.NoError(t, err)

	return simRes
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

func CreateNextMessage(ctx context.Context, solanaGoClient *rpc.Client, t *testing.T) (ccip_router.Any2SolanaRampMessage, [32]byte) {
	msg, hash, err := ccip.CreateNextMessage(ctx, solanaGoClient)
	require.NoError(t, err)
	return msg, hash
}

func NextSequenceNumber(ctx context.Context, solanaGoClient *rpc.Client, sourceChainStatePDA solana.PublicKey, t *testing.T) uint64 {
	num, err := ccip.NextSequenceNumber(ctx, solanaGoClient, sourceChainStatePDA)
	require.NoError(t, err)
	return num
}

func MakeEvmToSolanaMessage(t *testing.T, ccipReceiver solana.PublicKey, evmChainSelector uint64, solanaChainSelector uint64, data []byte) (ccip_router.Any2SolanaRampMessage, [32]byte) {
	msg, hash, err := ccip.MakeEvmToSolanaMessage(ccipReceiver, evmChainSelector, solanaChainSelector, data)
	require.NoError(t, err)
	return msg, hash
}
