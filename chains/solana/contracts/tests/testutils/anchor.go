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
	fundAccounts(ctx, accounts, solanaGoClient, t, waitAndRetryOpts{
		RemainingAttempts: 5,
		Timeout:           30 * time.Second,
		Timestep:          500 * time.Millisecond,
	})
}

type waitAndRetryOpts struct {
	RemainingAttempts uint
	Timeout           time.Duration
	Timestep          time.Duration
}

func (o waitAndRetryOpts) WithDecreasedAttempts() waitAndRetryOpts {
	return waitAndRetryOpts{
		RemainingAttempts: o.RemainingAttempts - 1,
		Timeout:           o.Timeout,
		Timestep:          o.Timestep,
	}
}

func fundAccounts(ctx context.Context, accounts []solana.PrivateKey, solanaGoClient *rpc.Client, t *testing.T, opts waitAndRetryOpts) {
	sigs := []solana.Signature{}
	for _, v := range accounts {
		sig, err := solanaGoClient.RequestAirdrop(ctx, v.PublicKey(), 1000*solana.LAMPORTS_PER_SOL, rpc.CommitmentFinalized)
		require.NoError(t, err)
		sigs = append(sigs, sig)
	}

	// wait for confirmation so later transactions don't fail
	remaining := accounts
	initTime := time.Now()
	for elapsed := time.Since(initTime); elapsed < opts.Timeout; elapsed = time.Since(initTime) {
		time.Sleep(opts.Timestep)

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
			return // all done!
		}
	}

	decreasedOpts := opts.WithDecreasedAttempts()
	if decreasedOpts.RemainingAttempts == 0 {
		require.NoError(t, fmt.Errorf("[%s]: unable to find transactions after all attempts", t.Name()))
	} else {
		fundAccounts(ctx, remaining, solanaGoClient, t, decreasedOpts) // recursive call with only remaining & with fewer attempts
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
