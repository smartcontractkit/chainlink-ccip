package main

import (
	"context"
	"crypto/sha256"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	confirm "github.com/gagliardetto/solana-go/rpc/sendAndConfirmTransaction"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

const (
	TRIGGER_CCIP_ROUTER_EVENTS_OK      = "trigger_all_events"
	TRIGGER_CCIP_ROUTER_EVENTS_REVERTS = "trigger_all_events_reverts"
)

var programIDs = map[string]string{
	"ccip-router": "7sDY5A5S5NZe1zcqEuZybW6ZxAna1NWUZxU4ypdn8UQU",
}

func main() {
	// Define command-line flags
	contractName := flag.String("contract", "ccip-router", "Contract name to emit events from (e.g., ccip-router)")
	resultType := flag.String("result", "OK", "Result type (OK or REVERT)")
	keypairPath := flag.String("keypair", "/path/to/solana/keypair.json", "Path to the keypair file. Needs to be funded first: https://faucet.solana.com/")
	numTx := flag.Int("num-tx", 2, "Number of transactions to send")

	flag.Parse()

	// Validate contract name
	programIDStr, ok := programIDs[*contractName]
	if !ok {
		log.Fatalf("Invalid contract name: %s", *contractName)
	}

	// Determine method name based on result type
	var methodName string
	switch *resultType {
	case "OK":
		methodName = TRIGGER_CCIP_ROUTER_EVENTS_OK
	case "REVERT":
		methodName = TRIGGER_CCIP_ROUTER_EVENTS_REVERTS
	default:
		log.Fatalf("Invalid result type: %s", *resultType)
	}

	// Connect to Solana devnet
	rpcClient := rpc.New("https://api.devnet.solana.com")
	wsClient, err := ws.Connect(context.Background(), "wss://api.devnet.solana.com")
	if err != nil {
		log.Fatalf("Failed to connect to websocket: %v", err)
	}
	defer wsClient.Close()

	// Load keypair from file
	keyBytes, err := os.ReadFile(*keypairPath)
	if err != nil {
		log.Fatalf("Failed to read keypair file: %v", err)
	}

	wallet, err := solana.PrivateKeyFromSolanaKeygenFileBytes(keyBytes)
	if err != nil {
		log.Fatalf("Failed to load wallet: %v", err)
	}

	// Program ID from the contract name
	programID := solana.MustPublicKeyFromBase58(programIDStr)

	ctx := context.Background()
	// Loop to send multiple transactions
	for i := range *numTx {
		fmt.Printf("\nSending transaction %d of %d\n", i+1, *numTx)
		if err := sendTransaction(ctx, rpcClient, wsClient, wallet, programID, methodName); err != nil {
			log.Printf("Failed to send transaction %d: %v", i+1, err)
			continue // Continue with next transaction even if this one fails
		}
		time.Sleep(time.Second * 10)
	}
}

// anchorDiscriminator computes the 8-byte discriminator.
// It prepends "global:" to the method name.
func anchorDiscriminator(methodName string) []byte {
	hash := sha256.Sum256([]byte("global:" + methodName))
	return hash[:8]
}

// sendTransaction sends a transaction to the contract with SkipPreflight enabled.
// This is necessary so transactions that will fail preflight checks can still be broadcasted.
func sendTransaction(ctx context.Context, rpcClient *rpc.Client, wsClient *ws.Client, wallet solana.PrivateKey, programID solana.PublicKey, methodName string) error {
	// Create the instruction with the correct discriminator
	instruction := solana.NewInstruction(
		programID,
		solana.AccountMetaSlice{},
		anchorDiscriminator(methodName),
	)

	// Get latest blockhash
	latest, err := rpcClient.GetLatestBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		return fmt.Errorf("failed to get latest blockhash: %v", err)
	}

	// Build transaction
	tx, err := solana.NewTransaction(
		[]solana.Instruction{instruction},
		latest.Value.Blockhash,
		solana.TransactionPayer(wallet.PublicKey()),
	)
	if err != nil {
		return fmt.Errorf("failed to create transaction: %v", err)
	}

	// Sign transaction
	_, err = tx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		if wallet.PublicKey().Equals(key) {
			return &wallet
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to sign transaction: %v", err)
	}

	// Create transaction options to skip preflight
	txOpts := rpc.TransactionOpts{
		SkipPreflight:       true, // <-- Important: skip preflight
		PreflightCommitment: rpc.CommitmentFinalized,
	}

	// Send and confirm transaction
	sig, err := confirm.SendAndConfirmTransactionWithOpts(
		ctx,
		rpcClient,
		wsClient,
		tx,
		txOpts,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to send transaction: %v", err)
	}

	fmt.Printf("Transaction successful! Signature: %s\n", sig)

	return nil
}
