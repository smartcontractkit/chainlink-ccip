package solutils

import (
	"encoding/binary"
	"os"
	"testing"

	"github.com/gagliardetto/solana-go"
	solRpc "github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/require"
)

func borshString(s string) []byte {
	b := make([]byte, 4+len(s))
	binary.LittleEndian.PutUint32(b, uint32(len(s)))
	copy(b[4:], s)
	return b
}

func TestParseTypeVersion(t *testing.T) {
	t.Parallel()

	programID := solana.MustPublicKeyFromBase58("Ccip842gzYHhvdDkSyi2YVCoAWPbYJoApMFzSxQroE9C")

	tests := []struct {
		name        string
		returned    string
		wantName    string
		wantVersion string
	}{
		// Exactly what the live programs answer, taken from a sweep of all four environments.
		{"release", "ccip-router 1.6.2", "ccip-router", "1.6.2"},
		{"offramp", "ccip-offramp 1.6.3", "ccip-offramp", "1.6.3"},
		// Pre-release versions must survive intact: 22 mainnet token pools report 0.1.0-dev, and
		// flattening or rejecting those is how the datastore ended up claiming they were 1.6.0.
		{"dev prerelease", "burnmint-token-pool 0.1.0-dev", "burnmint-token-pool", "0.1.0-dev"},
		{"candidate prerelease", "fee-quoter 1.6.1-candidate", "fee-quoter", "1.6.1-candidate"},
		{"hyphenated name", "lockrelease-token-pool 0.1.1-dev", "lockrelease-token-pool", "0.1.1-dev"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := parseTypeVersion(borshString(tt.returned), programID)
			require.NoError(t, err)
			require.Equal(t, tt.wantName, got.Name)
			require.Equal(t, tt.wantVersion, got.Version.String())
			require.Equal(t, tt.returned, got.String(), "String() must round-trip what the program reported")
		})
	}
}

func TestParseTypeVersion_Rejects(t *testing.T) {
	t.Parallel()

	programID := solana.MustPublicKeyFromBase58("Ccip842gzYHhvdDkSyi2YVCoAWPbYJoApMFzSxQroE9C")

	tests := []struct {
		name    string
		data    []byte
		wantErr string
	}{
		{"empty", nil, "too short"},
		{"truncated length prefix", []byte{1, 2}, "too short"},
		{"length longer than payload", append([]byte{99, 0, 0, 0}, []byte("short")...), "declares 99 bytes"},
		{"no space separator", borshString("ccip-router"), "want \"<name> <version>\""},
		{"unparseable version", borshString("ccip-router not-a-version"), "returned version"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := parseTypeVersion(tt.data, programID)
			require.ErrorContains(t, err, tt.wantErr)
		})
	}
}

func TestReadTypeVersion_RequiresFeePayer(t *testing.T) {
	t.Parallel()

	// Without a fee payer the simulation fails before the program executes, which is
	// indistinguishable from a program that has no type_version -- so it is rejected up front
	// rather than silently reported as ErrNoTypeVersion.
	_, err := ReadTypeVersion(t.Context(), nil,
		solana.MustPublicKeyFromBase58("Ccip842gzYHhvdDkSyi2YVCoAWPbYJoApMFzSxQroE9C"),
		solana.PublicKey{})
	require.ErrorContains(t, err, "fee payer is required")
}

// TestReadTypeVersion_Live exercises the whole path against real deployed programs. It is opt-in
// because it needs network access and a funded, system-owned fee payer on the cluster:
//
//	SOLANA_TYPEVERSION_E2E=<funded devnet pubkey> go test ./solutils/... -run Live -v
func TestReadTypeVersion_Live(t *testing.T) {
	feePayer := os.Getenv("SOLANA_TYPEVERSION_E2E")
	if feePayer == "" {
		t.Skip("set SOLANA_TYPEVERSION_E2E to a funded devnet pubkey to run")
	}

	client := solRpc.New(solRpc.DevNet_RPC)
	payer := solana.MustPublicKeyFromBase58(feePayer)

	t.Run("program with type_version", func(t *testing.T) {
		router := solana.MustPublicKeyFromBase58("Ccip842gzYHhvdDkSyi2YVCoAWPbYJoApMFzSxQroE9C")
		tv, err := ReadTypeVersion(t.Context(), client, router, payer)
		require.NoError(t, err)
		require.Equal(t, "ccip-router", tv.Name)
		t.Logf("devnet router reports %s", tv)
	})

	t.Run("account without type_version", func(t *testing.T) {
		// WSOL: an SPL mint, not a program. Must come back as ErrNoTypeVersion rather than an error
		// a caller would fail a deployment over.
		wsol := solana.MustPublicKeyFromBase58("So11111111111111111111111111111111111111112")
		_, err := ReadTypeVersion(t.Context(), client, wsol, payer)
		require.ErrorIs(t, err, ErrNoTypeVersion)
	})
}
