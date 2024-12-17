package utils

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"
	"testing"
	"time"

	"strconv"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/utils/fees"
)

func SendAndConfirm(ctx context.Context, t *testing.T, rpcClient *rpc.Client, instructions []solana.Instruction,
	signer solana.PrivateKey, commitment rpc.CommitmentType, opts ...TxModifier) *rpc.GetTransactionResult {
	emptyLookupTables := map[solana.PublicKey]solana.PublicKeySlice{}
	txres := sendTransactionWithLookupTables(ctx, rpcClient, t, instructions, signer, commitment, false, emptyLookupTables, opts...) // do not skipPreflight when expected to pass, preflight can help debug

	require.NotNil(t, txres.Meta)
	require.Nil(t, txres.Meta.Err, fmt.Sprintf("tx failed with: %+v", txres.Meta)) // tx should not err, print meta if it does (contains logs)
	return txres
}

func SendAndConfirmWithLookupTables(ctx context.Context, t *testing.T, rpcClient *rpc.Client, instructions []solana.Instruction,
	signer solana.PrivateKey, commitment rpc.CommitmentType, lookupTables map[solana.PublicKey]solana.PublicKeySlice, opts ...TxModifier) *rpc.GetTransactionResult {
	txres := sendTransactionWithLookupTables(ctx, rpcClient, t, instructions, signer, commitment, false, lookupTables, opts...) // do not skipPreflight when expected to pass, preflight can help debug

	require.NotNil(t, txres.Meta)
	require.Nil(t, txres.Meta.Err, fmt.Sprintf("tx failed with: %+v", txres.Meta)) // tx should not err, print meta if it does (contains logs)
	return txres
}

func SendAndFailWith(ctx context.Context, t *testing.T, rpcClient *rpc.Client, instructions []solana.Instruction,
	signer solana.PrivateKey, commitment rpc.CommitmentType, expectedErrors []string, opts ...TxModifier) *rpc.GetTransactionResult {
	emptyLookupTables := map[solana.PublicKey]solana.PublicKeySlice{}
	txres := sendTransactionWithLookupTables(ctx, rpcClient, t, instructions, signer, commitment, true, emptyLookupTables, opts...) // skipPreflight when expected to fail so revert captured onchain

	require.NotNil(t, txres.Meta)
	require.NotNil(t, txres.Meta.Err)
	logs := strings.Join(txres.Meta.LogMessages, " ")
	for _, expectedError := range expectedErrors {
		require.Contains(t, logs, expectedError, fmt.Sprintf("The logs did not contain '%s'. The logs were: %s", expectedError, logs))
	}
	return txres
}

func SendAndFailWithLookupTables(ctx context.Context, t *testing.T, rpcClient *rpc.Client, instructions []solana.Instruction,
	signer solana.PrivateKey, commitment rpc.CommitmentType, lookupTables map[solana.PublicKey]solana.PublicKeySlice, expectedErrors []string, opts ...TxModifier) *rpc.GetTransactionResult {
	txres := sendTransactionWithLookupTables(ctx, rpcClient, t, instructions, signer, commitment, true, lookupTables, opts...) // skipPreflight when expected to fail so revert captured onchain

	require.NotNil(t, txres.Meta)
	require.NotNil(t, txres.Meta.Err)
	logs := strings.Join(txres.Meta.LogMessages, " ")
	for _, expectedError := range expectedErrors {
		require.Contains(t, logs, expectedError, fmt.Sprintf("The logs did not contain '%s'. The logs were: %s", expectedError, logs))
	}
	return txres
}

func SendAndFailWithRPCError(ctx context.Context, t *testing.T, rpcClient *rpc.Client, instructions []solana.Instruction,
	signer solana.PrivateKey, commitment rpc.CommitmentType, expectedErrors []string) {
	hashRes, err := rpcClient.GetLatestBlockhash(ctx, rpc.CommitmentFinalized)
	require.NoError(t, err)

	tx, err := solana.NewTransaction(
		instructions,
		hashRes.Value.Blockhash,
		solana.TransactionPayer(signer.PublicKey()),
	)
	require.NoError(t, err)

	_, err = tx.Sign(func(_ solana.PublicKey) *solana.PrivateKey {
		return &signer
	})
	require.NoError(t, err)

	_, err = rpcClient.SendTransactionWithOpts(ctx, tx, rpc.TransactionOpts{SkipPreflight: false, PreflightCommitment: rpc.CommitmentProcessed})
	require.NotNil(t, err)

	errStr := err.Error()

	for _, expectedError := range expectedErrors {
		require.Contains(t, errStr, expectedError)
	}
}

// TxModifier is a dynamic function used to flexibly add components to a transaction such as additional signers, and compute budget parameters
type TxModifier func(tx *solana.Transaction, signers map[solana.PublicKey]solana.PrivateKey) error

func AddSigners(additionalSigners ...solana.PrivateKey) TxModifier {
	return func(_ *solana.Transaction, s map[solana.PublicKey]solana.PrivateKey) error {
		for _, v := range additionalSigners {
			s[v.PublicKey()] = v
		}
		return nil
	}
}

// AddComputeUnitLimit allows for configuring the total compute unit limit for a transaction - solana network default is 200K, maximum is 1.4M
// signature verification compute units can vary depending on searching for signatures
func AddComputeUnitLimit(v fees.ComputeUnitLimit) TxModifier {
	return func(tx *solana.Transaction, _ map[solana.PublicKey]solana.PrivateKey) error {
		return fees.SetComputeUnitLimit(tx, v)
	}
}

func sendTransactionWithLookupTables(ctx context.Context, rpcClient *rpc.Client, t *testing.T, instructions []solana.Instruction,
	signerAndPayer solana.PrivateKey, commitment rpc.CommitmentType, skipPreflight bool, lookupTables map[solana.PublicKey]solana.PublicKeySlice, opts ...TxModifier) *rpc.GetTransactionResult {
	hashRes, err := rpcClient.GetLatestBlockhash(ctx, rpc.CommitmentFinalized)
	require.NoError(t, err)

	tx, err := solana.NewTransaction(
		instructions,
		hashRes.Value.Blockhash,
		solana.TransactionAddressTables(lookupTables),
		solana.TransactionPayer(signerAndPayer.PublicKey()),
	)

	require.NoError(t, err)

	// build signers map
	signers := map[solana.PublicKey]solana.PrivateKey{}
	signers[signerAndPayer.PublicKey()] = signerAndPayer

	// set options before signing transaction
	for _, o := range opts {
		require.NoError(t, o(tx, signers))
	}

	_, err = tx.Sign(func(pub solana.PublicKey) *solana.PrivateKey {
		priv, ok := signers[pub]
		require.True(t, ok, fmt.Sprintf("Missing signer private key for %s", pub))
		return &priv
	})
	require.NoError(t, err)

	txsig, err := rpcClient.SendTransactionWithOpts(ctx, tx, rpc.TransactionOpts{SkipPreflight: skipPreflight, PreflightCommitment: rpc.CommitmentProcessed})
	require.NoError(t, err)

	var txStatus rpc.ConfirmationStatusType
	count := 0
	for txStatus != rpc.ConfirmationStatusConfirmed && txStatus != rpc.ConfirmationStatusFinalized {
		count++
		statusRes, sigErr := rpcClient.GetSignatureStatuses(ctx, true, txsig)
		require.NoError(t, sigErr)
		if statusRes != nil && len(statusRes.Value) > 0 && statusRes.Value[0] != nil {
			txStatus = statusRes.Value[0].ConfirmationStatus
		}
		time.Sleep(50 * time.Millisecond)
		if count > 500 {
			require.NoError(t, fmt.Errorf("unable to find transaction within timeout"))
		}
	}

	v := uint64(0)
	txres, err := rpcClient.GetTransaction(ctx, txsig, &rpc.GetTransactionOpts{
		Commitment:                     commitment,
		MaxSupportedTransactionVersion: &v,
	})
	require.NoError(t, err)
	return txres
}

func SimulateTransaction(ctx context.Context, t *testing.T, rpcClient *rpc.Client, instructions []solana.Instruction, signer solana.PrivateKey) *rpc.SimulateTransactionResponse {
	simRes, err := simulateTransaction(ctx, t, rpcClient, instructions, signer)
	require.NoError(t, err)

	return simRes
}

func simulateTransaction(ctx context.Context, t *testing.T, rpcClient *rpc.Client, instructions []solana.Instruction,
	signer solana.PrivateKey) (*rpc.SimulateTransactionResponse, error) {
	hashRes, err := rpcClient.GetLatestBlockhash(ctx, rpc.CommitmentFinalized)
	require.NoError(t, err)

	tx, err := solana.NewTransaction(
		instructions,
		hashRes.Value.Blockhash,
		solana.TransactionPayer(signer.PublicKey()),
	)
	require.NoError(t, err)

	_, err = tx.Sign(func(_ solana.PublicKey) *solana.PrivateKey {
		return &signer
	})
	require.NoError(t, err)

	return rpcClient.SimulateTransaction(ctx, tx)
}

func ExtractReturnValue(ctx context.Context, t *testing.T, logs []string, programID string) []byte {
	if logs == nil {
		return []byte{}
	}
	for _, log := range logs {
		if strings.HasPrefix(log, "Program return: "+programID) {
			parts := strings.Split(log, " ")
			encoded := parts[len(parts)-1]
			ret, err := base64.StdEncoding.DecodeString(encoded)
			require.NoError(t, err)
			return ret
		}
	}
	return []byte{}
}

func ExtractReturnedError(ctx context.Context, t *testing.T, logs []string, programID string) *int {
	if logs == nil {
		return nil
	}

	for _, log := range logs {
		if strings.Contains(log, "Error Number: ") {
			// extract error number from the string
			parts := strings.Split(log, "Error Number: ")
			if len(parts) > 1 {
				numberPart := strings.Split(parts[1], ".")[0]
				errorNumber, err := strconv.Atoi(strings.TrimSpace(numberPart))
				require.NoError(t, err)
				return &errorNumber
			}
		}
	}
	return nil
}

func ExtractTypedReturnValue[T any](ctx context.Context, t *testing.T, logs []string, programID string, decoderFn func([]byte) T) T {
	bytes := ExtractReturnValue(ctx, t, logs, programID)
	return decoderFn(bytes)
}

func GetAccountDataBorshInto(ctx context.Context, solanaGoClient *rpc.Client, account solana.PublicKey, commitment rpc.CommitmentType, data interface{}) error {
	resp, err := solanaGoClient.GetAccountInfoWithOpts(
		ctx,
		account,
		&rpc.GetAccountInfoOpts{
			Commitment: commitment,
			DataSlice:  nil,
		},
	)
	if err != nil {
		return err
	}
	return bin.NewBorshDecoder(resp.Value.Data.GetBinary()).Decode(data)
}

func AssertClosedAccount(ctx context.Context, t *testing.T, solanaGoClient *rpc.Client, accountKey solana.PublicKey, commitment rpc.CommitmentType) {
	_, err := solanaGoClient.GetAccountInfoWithOpts(ctx, accountKey, &rpc.GetAccountInfoOpts{
		Commitment: commitment,
	})
	require.Error(t, err)
}
