package common

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/fees"
)

func SendAndConfirm(ctx context.Context, rpcClient *rpc.Client, instructions []solana.Instruction,
	signer solana.PrivateKey, commitment rpc.CommitmentType, opts ...TxModifier) (*rpc.GetTransactionResult, error) {
	emptyLookupTables := map[solana.PublicKey]solana.PublicKeySlice{}
	return SendAndConfirmWithLookupTables(ctx, rpcClient, instructions, signer, commitment, emptyLookupTables, opts...)
}

func SendAndConfirmWithLookupTables(ctx context.Context, rpcClient *rpc.Client, instructions []solana.Instruction,
	signer solana.PrivateKey, commitment rpc.CommitmentType, lookupTables map[solana.PublicKey]solana.PublicKeySlice, opts ...TxModifier) (*rpc.GetTransactionResult, error) {
	txres, err := sendTransactionWithLookupTables(ctx, rpcClient, instructions, signer, commitment, false, lookupTables, opts...) // do not skipPreflight when expected to pass, preflight can help debug
	if err != nil {
		return nil, err
	}

	if txres.Meta == nil {
		return nil, fmt.Errorf("txres.Meta == nil")
	}

	if txres.Meta.Err != nil {
		return nil, fmt.Errorf("tx failed with: %+v", txres.Meta) // tx should not err, print meta if it does (contains logs)
	}
	return txres, nil
}

func SendAndFailWith(ctx context.Context, rpcClient *rpc.Client, instructions []solana.Instruction,
	signer solana.PrivateKey, commitment rpc.CommitmentType, expectedErrors []string, opts ...TxModifier) (*rpc.GetTransactionResult, error) {
	emptyLookupTables := map[solana.PublicKey]solana.PublicKeySlice{}
	return SendAndFailWithLookupTables(ctx, rpcClient, instructions, signer, commitment, emptyLookupTables, expectedErrors, opts...)
}

func SendAndFailWithLookupTables(ctx context.Context, rpcClient *rpc.Client, instructions []solana.Instruction,
	signer solana.PrivateKey, commitment rpc.CommitmentType, lookupTables map[solana.PublicKey]solana.PublicKeySlice, expectedErrors []string, opts ...TxModifier) (*rpc.GetTransactionResult, error) {
	txres, err := sendTransactionWithLookupTables(ctx, rpcClient, instructions, signer, commitment, true, lookupTables, opts...) // skipPreflight when expected to fail so revert captured onchain
	if err != nil {
		return nil, err
	}

	if txres.Meta == nil || txres.Meta.Err == nil {
		return nil, fmt.Errorf("txres.Meta == nil || txres.Meta.Err == nil")
	}
	logs := strings.Join(txres.Meta.LogMessages, " ")
	for _, expectedError := range expectedErrors {
		if !strings.Contains(logs, expectedError) {
			return nil, fmt.Errorf("The logs did not contain '%s'. The logs were: %s", expectedError, logs)
		}
	}
	return txres, nil
}

func SendAndFailWithRPCError(ctx context.Context, rpcClient *rpc.Client, instructions []solana.Instruction,
	signer solana.PrivateKey, commitment rpc.CommitmentType, expectedErrors []string, opts ...TxModifier) error {
	hashRes, err := rpcClient.GetLatestBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		return err
	}

	tx, err := solana.NewTransaction(
		instructions,
		hashRes.Value.Blockhash,
		solana.TransactionPayer(signer.PublicKey()),
	)
	if err != nil {
		return err
	}

	// build signers map
	signers := map[solana.PublicKey]solana.PrivateKey{}
	signers[signer.PublicKey()] = signer

	// set options before signing transaction
	for _, o := range opts {
		if err = o(tx, signers); err != nil {
			return err
		}
	}

	if _, err = tx.Sign(func(_ solana.PublicKey) *solana.PrivateKey {
		return &signer
	}); err != nil {
		return err
	}

	_, err = rpcClient.SendTransactionWithOpts(ctx, tx, rpc.TransactionOpts{SkipPreflight: false, PreflightCommitment: rpc.CommitmentProcessed})
	if err == nil {
		return fmt.Errorf("expected RPC error - none found")
	}

	errStr := err.Error()
	for _, expectedError := range expectedErrors {
		if !strings.Contains(errStr, expectedError) {
			return fmt.Errorf("The error did not contain '%s'. The error was: %s", expectedError, errStr)
		}
	}
	return nil
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

func AddComputeUnitPrice(v fees.ComputeUnitPrice) TxModifier {
	return func(tx *solana.Transaction, _ map[solana.PublicKey]solana.PrivateKey) error {
		return fees.SetComputeUnitPrice(tx, v)
	}
}

func sendTransactionWithLookupTables(ctx context.Context, rpcClient *rpc.Client, instructions []solana.Instruction,
	signerAndPayer solana.PrivateKey, commitment rpc.CommitmentType, skipPreflight bool, lookupTables map[solana.PublicKey]solana.PublicKeySlice, opts ...TxModifier) (*rpc.GetTransactionResult, error) {
	hashRes, err := rpcClient.GetLatestBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		return nil, err
	}

	tx, err := solana.NewTransaction(
		instructions,
		hashRes.Value.Blockhash,
		solana.TransactionAddressTables(lookupTables),
		solana.TransactionPayer(signerAndPayer.PublicKey()),
	)
	if err != nil {
		return nil, err
	}

	// build signers map
	signers := map[solana.PublicKey]solana.PrivateKey{}
	signers[signerAndPayer.PublicKey()] = signerAndPayer

	// set options before signing transaction
	for _, o := range opts {
		if err = o(tx, signers); err != nil {
			return nil, err
		}
	}

	if _, err = tx.Sign(func(pub solana.PublicKey) *solana.PrivateKey {
		priv, ok := signers[pub]
		if !ok {
			fmt.Printf("ERROR: Missing signer private key for %s\n", pub)
		}
		return &priv
	}); err != nil {
		return nil, err
	}

	txsig, err := rpcClient.SendTransactionWithOpts(ctx, tx, rpc.TransactionOpts{SkipPreflight: skipPreflight, PreflightCommitment: rpc.CommitmentProcessed})
	if err != nil {
		fmt.Println(tx) // debugging if tx errors
		return nil, err
	}

	var txStatus rpc.ConfirmationStatusType
	count := 0
	for txStatus != rpc.ConfirmationStatusConfirmed && txStatus != rpc.ConfirmationStatusFinalized {
		count++
		statusRes, sigErr := rpcClient.GetSignatureStatuses(ctx, true, txsig)
		if sigErr != nil {
			return nil, sigErr
		}
		if statusRes != nil && len(statusRes.Value) > 0 && statusRes.Value[0] != nil {
			txStatus = statusRes.Value[0].ConfirmationStatus
		}
		time.Sleep(50 * time.Millisecond)
		if count > 500 {
			return nil, fmt.Errorf("unable to find transaction within timeout")
		}
	}

	v := uint64(0)
	return rpcClient.GetTransaction(ctx, txsig, &rpc.GetTransactionOpts{
		Commitment:                     commitment,
		MaxSupportedTransactionVersion: &v,
	})
}

// shared function for transaction simulation
func buildSignedTx(ctx context.Context, rpcClient *rpc.Client, instructions []solana.Instruction,
	signer solana.PrivateKey) (*solana.Transaction, error) {
	hashRes, err := rpcClient.GetLatestBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		return nil, err
	}

	tx, err := solana.NewTransaction(
		instructions,
		hashRes.Value.Blockhash,
		solana.TransactionPayer(signer.PublicKey()),
	)
	if err != nil {
		return nil, err
	}

	if _, err = tx.Sign(func(_ solana.PublicKey) *solana.PrivateKey {
		return &signer
	}); err != nil {
		return nil, err
	}

	return tx, err
}

func buildSignedTxWithLookupTables(ctx context.Context, rpcClient *rpc.Client, instructions []solana.Instruction,
	signer solana.PrivateKey, lookupTables map[solana.PublicKey]solana.PublicKeySlice) (*solana.Transaction, error) {
	hashRes, err := rpcClient.GetLatestBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		return nil, err
	}

	tx, err := solana.NewTransaction(
		instructions,
		hashRes.Value.Blockhash,
		solana.TransactionAddressTables(lookupTables),
		solana.TransactionPayer(signer.PublicKey()),
	)
	if err != nil {
		return nil, err
	}

	if _, err = tx.Sign(func(_ solana.PublicKey) *solana.PrivateKey {
		return &signer
	}); err != nil {
		return nil, err
	}

	return tx, err
}

func SimulateTransaction(ctx context.Context, rpcClient *rpc.Client, instructions []solana.Instruction,
	signer solana.PrivateKey) (*rpc.SimulateTransactionResponse, error) {
	tx, err := buildSignedTx(ctx, rpcClient, instructions, signer)
	if err != nil {
		return nil, err
	}
	return rpcClient.SimulateTransaction(ctx, tx)
}

func SimulateTransactionWithOpts(ctx context.Context, rpcClient *rpc.Client, instructions []solana.Instruction,
	signer solana.PrivateKey, opts rpc.SimulateTransactionOpts) (*rpc.SimulateTransactionResponse, error) {
	tx, err := buildSignedTx(ctx, rpcClient, instructions, signer)
	if err != nil {
		return nil, err
	}
	return rpcClient.SimulateTransactionWithOpts(ctx, tx, &opts)
}

func SimulateTransactionWithLookupTables(ctx context.Context, rpcClient *rpc.Client, instructions []solana.Instruction,
	signer solana.PrivateKey, lookupTables map[solana.PublicKey]solana.PublicKeySlice, opts rpc.SimulateTransactionOpts) (*rpc.SimulateTransactionResponse, error) {
	tx, err := buildSignedTxWithLookupTables(ctx, rpcClient, instructions, signer, lookupTables)
	if err != nil {
		return nil, err
	}
	return rpcClient.SimulateTransactionWithOpts(ctx, tx, &opts)
}

func ExtractReturnValue(ctx context.Context, logs []string, programID string) ([]byte, error) {
	if logs == nil {
		return []byte{}, nil
	}

	for _, log := range logs {
		if strings.HasPrefix(log, "Program return: "+programID) {
			parts := strings.Split(log, " ")
			encoded := parts[len(parts)-1]
			return base64.StdEncoding.DecodeString(encoded)
		}
	}
	return []byte{}, nil
}

func ExtractReturnedError(ctx context.Context, logs []string, programID string) *string {
	if logs == nil {
		return nil
	}

	for _, log := range logs {
		if strings.Contains(log, "Error Code: ") {
			parts := strings.Split(log, "Error Code: ")
			if len(parts) > 1 {
				codePart := strings.Split(parts[1], ".")[0]
				return &codePart
			}
		}
	}
	return nil
}

type anchorType[T any] interface {
	*T
	UnmarshalWithDecoder(decoder *bin.Decoder) (err error)
}

func ExtractAnchorTypedReturnValue[T any, PT anchorType[T]](ctx context.Context, logs []string, programID string) (*T, error) {
	var result T
	bytes, err := ExtractReturnValue(ctx, logs, programID)
	if err != nil {
		return nil, err
	}
	err = PT(&result).UnmarshalWithDecoder(bin.NewBinDecoder(bytes))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func ExtractTypedReturnValue[T any](ctx context.Context, logs []string, programID string, decoderFn func([]byte) T) (T, error) {
	bytes, err := ExtractReturnValue(ctx, logs, programID)
	return decoderFn(bytes), err
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

func IsClosedAccount(ctx context.Context, solanaGoClient *rpc.Client, accountKey solana.PublicKey, commitment rpc.CommitmentType) bool {
	_, err := solanaGoClient.GetAccountInfoWithOpts(ctx, accountKey, &rpc.GetAccountInfoOpts{
		Commitment: commitment,
	})
	return err != nil
}
