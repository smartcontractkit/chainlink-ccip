package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	confirm "github.com/gagliardetto/solana-go/rpc/sendAndConfirmTransaction"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

func sendTransaction(ctx context.Context, rpcClient *rpc.Client, wsClient *ws.Client, wallet solana.PrivateKey, programID solana.PublicKey) error {
	// Create the instruction with the correct discriminator
	instruction := solana.NewInstruction(
		programID,
		solana.AccountMetaSlice{},
		anchorDiscriminator("trigger_all_events"),
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

	// Send and confirm transaction
	sig, err := confirm.SendAndConfirmTransaction(
		ctx,
		rpcClient,
		wsClient,
		tx,
	)
	if err != nil {
		return fmt.Errorf("failed to send transaction: %v", err)
	}

	fmt.Printf("Transaction successful! Signature: %s\n", sig)

	return nil
}

func main() {
	// Connect to Solana devnet
	rpcClient := rpc.New("https://api.devnet.solana.com")
	wsClient, err := ws.Connect(context.Background(), "wss://api.devnet.solana.com")
	if err != nil {
		log.Fatalf("Failed to connect to websocket: %v", err)
	}
	defer wsClient.Close()

	// Load keypair from file
	keyBytes, err := os.ReadFile("/Users/chainlink/.config/solana/id.json")
	if err != nil {
		log.Fatalf("Failed to read keypair file: %v", err)
	}

	wallet, err := solana.PrivateKeyFromSolanaKeygenFileBytes(keyBytes)
	if err != nil {
		log.Fatalf("Failed to load wallet: %v", err)
	}

	// Program ID from your deployed contract
	programID := solana.MustPublicKeyFromBase58("7sDY5A5S5NZe1zcqEuZybW6ZxAna1NWUZxU4ypdn8UQU")

	// Number of transactions to send
	numTx := 1 // Change this to your desired number
	ctx := context.Background()
	// Loop to send multiple transactions
	for i := range numTx {
		fmt.Printf("\nSending transaction %d of %d\n", i+1, numTx)
		if err := sendTransaction(ctx, rpcClient, wsClient, wallet, programID); err != nil {
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
