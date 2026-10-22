package adapters

import (
	"bytes"
	"testing"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"

	ccip_common_v0_1_0 "github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_0/ccip_common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v1_6_4/ccip_common"
)

// TestDecodeTokenAdminRegistry covers the V1/V2 layout dispatch in
// decodeTokenAdminRegistry. V1 accounts (169 bytes) predate the trailing
// supports_auto_derivation bool and cannot be decoded by the v1.6.4 binding.
func TestDecodeTokenAdminRegistry(t *testing.T) {
	t.Parallel()

	mint := solana.MustPublicKeyFromBase58("4wgrRzSPdQDph7zW2rUUwyQprgmL3xpBvMFAewmuEccn")
	admin := solana.MustPublicKeyFromBase58("7oZnxiocDK1aa9XAQC3CZ1VHKFkKwLuwRK8NddhU3FT2")
	lookupTable := solana.MustPublicKeyFromBase58("iSwp6Dp8qMy89K69oWwQERLEYkjzxKDnWiRgqLodnCR")

	t.Run("v2 layout decodes as-is", func(t *testing.T) {
		t.Parallel()

		want := ccip_common.TokenAdminRegistry{
			Version:                1,
			Administrator:          admin,
			LookupTable:            lookupTable,
			Mint:                   mint,
			SupportsAutoDerivation: true,
		}
		var buf bytes.Buffer
		require.NoError(t, want.MarshalWithEncoder(bin.NewBorshEncoder(&buf)))
		require.Len(t, buf.Bytes(), 170)

		got, err := decodeTokenAdminRegistry(buf.Bytes())
		require.NoError(t, err)
		require.Equal(t, want, got)
	})

	t.Run("v1 layout falls back and defaults SupportsAutoDerivation to false", func(t *testing.T) {
		t.Parallel()

		legacy := ccip_common_v0_1_0.TokenAdminRegistry{
			Version:       1,
			Administrator: admin,
			LookupTable:   lookupTable,
			Mint:          mint,
		}
		var buf bytes.Buffer
		require.NoError(t, legacy.MarshalWithEncoder(bin.NewBorshEncoder(&buf)))
		require.Len(t, buf.Bytes(), 169)

		got, err := decodeTokenAdminRegistry(buf.Bytes())
		require.NoError(t, err)
		require.Equal(t, uint8(1), got.Version)
		require.Equal(t, admin, got.Administrator)
		require.Equal(t, lookupTable, got.LookupTable)
		require.Equal(t, mint, got.Mint)
		require.False(t, got.SupportsAutoDerivation)
	})

	t.Run("undecodable data returns an error", func(t *testing.T) {
		t.Parallel()

		_, err := decodeTokenAdminRegistry([]byte{0x01, 0x02, 0x03})
		require.Error(t, err)
	})
}
