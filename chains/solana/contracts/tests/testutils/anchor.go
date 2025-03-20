package testutils

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"

	"github.com/pelletier/go-toml/v2"

	"github.com/stretchr/testify/require"
)

var PathToAnchorConfig = filepath.Join(ProjectRoot, "Anchor.toml")

func DeployAllPrograms(t *testing.T, pathToAnchorConfig string, admin solana.PrivateKey) *rpc.Client {
	return rpc.New(SetupTestValidatorWithAnchorPrograms(t, pathToAnchorConfig, admin.PublicKey().String()))
}

func FundAccounts(ctx context.Context, accounts []solana.PrivateKey, solanaGoClient *rpc.Client, t *testing.T) {
	fundAccounts(ctx, accounts, solanaGoClient, t, 3)
}

func fundAccounts(ctx context.Context, accounts []solana.PrivateKey, solanaGoClient *rpc.Client, t *testing.T, remainingAttempts uint) {
	if remainingAttempts == 0 {
		require.NoError(t, fmt.Errorf("[%s]: unable to find transactions after all attempts", t.Name()))
	}

	sigs := []solana.Signature{}
	fmt.Printf("[%s]: Requesting airdrop for %d accounts and remaining attempts %d\n", t.Name(), len(accounts), remainingAttempts)
	for i, v := range accounts {
		sig, err := solanaGoClient.RequestAirdrop(ctx, v.PublicKey(), 1000*solana.LAMPORTS_PER_SOL, rpc.CommitmentFinalized)
		require.NoError(t, err)
		fmt.Printf("[%s]: Requested airdrop for account #%d out of %d: %s\n", t.Name(), i, len(accounts), sig)
		sigs = append(sigs, sig)
	}

	// wait for confirmation so later transactions don't fail
	remaining := accounts
	initTime := time.Now()
	for elapsed := time.Since(initTime); elapsed < 1*time.Minute; elapsed = time.Since(initTime) {
		time.Sleep(500 * time.Millisecond)

		statusRes, sigErr := solanaGoClient.GetSignatureStatuses(ctx, true, sigs...)
		require.NoError(t, sigErr)
		require.NotNil(t, statusRes)
		require.NotNil(t, statusRes.Value)

		accountsWithNonFinalizedFunding := []solana.PrivateKey{}
		for i, res := range statusRes.Value {
			if res == nil || res.ConfirmationStatus == rpc.ConfirmationStatusProcessed || res.ConfirmationStatus == rpc.ConfirmationStatusConfirmed {
				accountsWithNonFinalizedFunding = append(accountsWithNonFinalizedFunding, accounts[i])
			}
		}
		remaining = accountsWithNonFinalizedFunding
		if len(remaining) == 0 {
			break
		}

		printableStatuses, err := json.Marshal(statusRes)
		require.NoError(t, err)
		fmt.Printf("[%s]: Waiting for airdrop confirmation, %d transactions remaining out of %d, elapsed time: %s\nSignatureStatuses: %s\n\n", t.Name(), len(remaining), len(accounts), elapsed, printableStatuses)
	}

	if len(remaining) > 0 {
		fmt.Printf("[%s]: unable to find %d transactions out of %d within timeout when remaining attempts is %d", t.Name(), len(remaining), len(accounts), remainingAttempts)
		fundAccounts(ctx, remaining, solanaGoClient, t, remainingAttempts-1) // recursive call with only remaining & with fewer attempts
	}
}

func SetupTestValidatorWithAnchorPrograms(t *testing.T, pathToAnchorConfig string, upgradeAuthority string) string {
	anchorData := struct {
		Programs struct {
			Localnet map[string]string
		}
	}{}

	// upload programs to validator
	anchorBytes, err := os.ReadFile(pathToAnchorConfig)
	require.NoError(t, err)
	require.NoError(t, toml.Unmarshal(anchorBytes, &anchorData))

	flags := []string{}
	for k, v := range anchorData.Programs.Localnet {
		flags = append(flags, "--upgradeable-program", v, filepath.Join(ContractsDir, k+".so"), upgradeAuthority)
	}
	url, _ := SetupLocalSolNodeWithFlags(t, flags...)
	return url
}

func WaitForTheNextBlock(client *rpc.Client, timeout time.Duration, commitment rpc.CommitmentType) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return WaitForNewBlock(ctx, client, 1, commitment)
}

func WaitForNewBlock(ctx context.Context, client *rpc.Client, height uint64, commitment rpc.CommitmentType) error {
	initialHeight, err := client.GetBlockHeight(ctx, commitment)
	if err != nil {
		return fmt.Errorf("failed to get initial block height: %w", err)
	}

	targetFinalHeight := initialHeight + height

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			currentHeight, err := client.GetBlockHeight(ctx, commitment)
			if err != nil {
				return fmt.Errorf("failed to get current block height: %w", err)
			}

			if currentHeight >= targetFinalHeight {
				return nil
			}
		}
	}
}
