package testutils

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"testing"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/utils"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
)

// SetupLocalSolNode sets up a local solana node via solana cli, and returns the url
func SetupLocalSolNode(t *testing.T) string {
	t.Helper()

	url, _ := SetupLocalSolNodeWithFlags(t)

	return url
}

// helper function to get a set of different random open ports
func getPorts(t *testing.T) (port string, wsPort string, faucetPort string) {
	t.Helper()

	attempts := 5

	for i := 0; i < attempts; i++ {
		port = utils.MustRandomPort(t)

		portInt, _ := strconv.Atoi(port)
		wsPort = strconv.Itoa(portInt + 1) // ws port is always +1 from the rpc port, required by solana
		if !utils.IsPortOpen(t, wsPort) {
			continue
		}

		faucetPort = utils.MustRandomPort(t)
		if faucetPort != port && faucetPort != wsPort {
			return
		}
	}

	panic(fmt.Sprintf("unable to find unique open ports after %d attempts", attempts))
}

// SetupLocalSolNode sets up a local solana node via solana cli, and returns the url
func SetupLocalSolNodeWithFlags(t *testing.T, flags ...string) (string, string) {
	t.Helper()

	port, wsPort, faucetPort := getPorts(t)

	url := "http://127.0.0.1:" + port
	wsURL := "ws://127.0.0.1:" + wsPort

	args := append([]string{
		"--reset",
		"--rpc-port", port,
		"--faucet-port", faucetPort,
		"--ledger", t.TempDir(),
		// Configurations to make the local cluster faster
		"--ticks-per-slot", "8", // value in mainnet: 64
		// account data direct mapping feature is disabled on mainnet,
		// so we disable it here to make the local cluster more similar to mainnet
		"--deactivate-feature", "EenyoWx9UMXYKpR8mW5Jmfmy2fRjzUtM7NduYMY8bx33",
	}, flags...)

	cmd := exec.Command("solana-test-validator", args...)

	var stdErr bytes.Buffer
	cmd.Stderr = &stdErr
	var stdOut bytes.Buffer
	cmd.Stdout = &stdOut
	require.NoError(t, cmd.Start())
	t.Cleanup(func() {
		assert.NoError(t, cmd.Process.Kill())
		if err2 := cmd.Wait(); assert.Error(t, err2) {
			if !assert.Contains(t, err2.Error(), "signal: killed", cmd.ProcessState.String()) {
				t.Logf("solana-test-validator\n stdout: %s\n stderr: %s", stdOut.String(), stdErr.String())
			}
		}
	})

	// Wait for api server to boot
	var ready bool
	for i := 0; i < 30; i++ {
		time.Sleep(time.Second)
		client := rpc.New(url)
		out, err := client.GetHealth(tests.Context(t))
		if err != nil || out != rpc.HealthOk {
			t.Logf("API server not ready yet (attempt %d)\n", i+1)
			continue
		}
		ready = true
		break
	}
	if !ready {
		t.Logf("Cmd output: %s\nCmd error: %s\n", stdOut.String(), stdErr.String())
	}
	require.True(t, ready)

	return url, wsURL
}

func FundTestAccounts(t *testing.T, keys []solana.PublicKey, url string) {
	t.Helper()

	for i := range keys {
		account := keys[i].String()
		_, err := exec.Command("solana", "airdrop", "100",
			account,
			"--url", url,
		).Output()
		require.NoError(t, err)
	}
}
