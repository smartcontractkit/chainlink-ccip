package utils

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"

	"github.com/pelletier/go-toml/v2"

	"github.com/stretchr/testify/require"
)

var PathToAnchorConfig = filepath.Join(ProjectRoot, "Anchor.toml")

var ZeroAddress = [32]byte{}

func MakeRandom32ByteArray() [32]byte {
	a := make([]byte, 32)
	if _, err := rand.Read(a); err != nil {
		panic(err) // should never panic but check in case
	}
	return [32]byte(a)
}

func Uint64ToLE(chain uint64) []byte {
	chainLE := make([]byte, 8)
	binary.LittleEndian.PutUint64(chainLE, chain)
	return chainLE
}

func To28BytesLE(value uint64) [28]byte {
	le := make([]byte, 28)
	binary.LittleEndian.PutUint64(le, value)
	return [28]byte(le)
}

func To28BytesBE(value uint64) [28]byte {
	be := make([]byte, 28)
	binary.BigEndian.PutUint64(be[20:], value)
	return [28]byte(be)
}

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func Discriminator(namespace, name string) []byte {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%s:%s", namespace, name)))
	return h.Sum(nil)[:8]
}

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
		if count > 60 {
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

func IsEvent(event string, data []byte) bool {
	if len(data) < 8 {
		return false
	}
	d := Discriminator("event", event)
	return bytes.Equal(d, data[:8])
}

func ParseEvent(logs []string, event string, obj interface{}, shouldPrint ...bool) error {
	for _, v := range logs {
		if strings.Contains(v, "Program data:") {
			encodedData := strings.TrimSpace(strings.TrimPrefix(v, "Program data:"))
			data, err := base64.StdEncoding.DecodeString(encodedData)
			if err != nil {
				return err
			}
			if IsEvent(event, data) {
				if err := bin.UnmarshalBorsh(obj, data); err != nil {
					return err
				}

				if len(shouldPrint) > 0 && shouldPrint[0] {
					fmt.Printf("%s: %+v\n", event, obj)
				}
				return nil
			}
		}
	}
	return fmt.Errorf("%s: event not found", event)
}

func ParseMultipleEvents[T any](logs []string, event string, shouldPrint bool) ([]T, error) {
	var results []T
	for _, v := range logs {
		if strings.Contains(v, "Program data:") {
			encodedData := strings.TrimSpace(strings.TrimPrefix(v, "Program data:"))
			data, err := base64.StdEncoding.DecodeString(encodedData)
			if err != nil {
				return nil, err
			}
			if IsEvent(event, data) {
				var obj T
				if err := bin.UnmarshalBorsh(&obj, data); err != nil {
					return nil, err
				}

				if shouldPrint {
					fmt.Printf("%s: %+v\n", event, obj)
				}

				results = append(results, obj)
			}
		}
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("%s: event not found", event)
	}

	return results, nil
}

func GetBlockTime(ctx context.Context, client *rpc.Client, commitment rpc.CommitmentType) (*solana.UnixTimeSeconds, error) {
	block, err := client.GetBlockHeight(ctx, commitment)
	if err != nil {
		return nil, fmt.Errorf("failed to get block height: %w", err)
	}

	blockTime, err := client.GetBlockTime(ctx, block)
	if err != nil {
		return nil, fmt.Errorf("failed to get block time: %w", err)
	}

	return blockTime, nil
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
