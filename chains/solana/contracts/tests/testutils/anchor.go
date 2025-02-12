package testutils

import (
	"context"
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
	sigs := []solana.Signature{}
	for _, v := range accounts {
		sig, err := solanaGoClient.RequestAirdrop(ctx, v.PublicKey(), 1000*solana.LAMPORTS_PER_SOL, rpc.CommitmentFinalized)
		require.NoError(t, err)
		sigs = append(sigs, sig)
	}

	// wait for confirmation so later transactions don't fail
	remaining := len(sigs)
	count := 0
	for remaining > 0 {
		count++
		statusRes, sigErr := solanaGoClient.GetSignatureStatuses(ctx, true, sigs...)
		require.NoError(t, sigErr)
		require.NotNil(t, statusRes)
		require.NotNil(t, statusRes.Value)

		unconfirmedTxCount := 0
		for _, res := range statusRes.Value {
			if res == nil || res.ConfirmationStatus == rpc.ConfirmationStatusProcessed || res.ConfirmationStatus == rpc.ConfirmationStatusConfirmed {
				unconfirmedTxCount++
			}
		}
		remaining = unconfirmedTxCount

		time.Sleep(500 * time.Millisecond)
		if count > 120 {
			require.NoError(t, fmt.Errorf("unable to find transaction within timeout"))
		}
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
