package sequences

import (
	"bytes"
	"testing"

	"github.com/Masterminds/semver/v3"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v1_6_4/burnmint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
)

func TestSelectTokenPoolRef(t *testing.T) {
	const sel = uint64(16423721717087811551)
	program := solana.NewWallet().PublicKey()
	migratedMint := solana.NewWallet().PublicKey()
	legacyMint := solana.NewWallet().PublicKey()
	unknownMint := solana.NewWallet().PublicKey()

	ref := func(address solana.PublicKey, typ, qualifier, version string) datastore.AddressRef {
		return datastore.AddressRef{
			ChainSelector: sel,
			Address:       address.String(),
			Type:          datastore.ContractType(typ),
			Qualifier:     qualifier,
			Version:       semver.MustParse(version),
		}
	}
	// A shared BurnMint program upgraded to 2.x, with one token migrated (CCIP-BnM) and one still on
	// the 1.x interface (LEGACY).
	ds := datastore.NewMemoryDataStore()
	for _, r := range []datastore.AddressRef{
		ref(program, "BurnMintTokenPool", "CLLCCIP", "1.6.0"),
		ref(program, "BurnMintTokenPool", "CLLCCIP", "2.0.0"),
		ref(program, "BurnMintTokenPool", "CLLCCIP", "2.1.0"),
		ref(program, "BurnMintTokenPool", "TokenPoolProgram", "2.1.0"),
		ref(program, "BurnMintTokenPool", "CCIP-BnM", "2.0.0"),
		ref(migratedMint, "SPLTokens", "CCIP-BnM", "1.0.0"),
		ref(legacyMint, "SPL2022Tokens", "LEGACY", "1.0.0"),
	} {
		require.NoError(t, ds.Addresses().Add(r))
	}
	sealed := ds.Seal()

	on := func(mint solana.PublicKey, interfaceVersion uint8) *tokenPoolInstance {
		return &tokenPoolInstance{Mint: mint, InterfaceVersion: interfaceVersion}
	}

	t.Run("migrated token resolves to its 2.x instance ref", func(t *testing.T) {
		got, err := selectTokenPoolRef(sealed, sel, program, on(migratedMint, 2))
		require.NoError(t, err)
		require.Equal(t, "CCIP-BnM", got.Qualifier)
		require.Equal(t, "2.0.0", got.Version.String())
	})
	t.Run("token on interface 1 resolves to the latest 1.x ref", func(t *testing.T) {
		got, err := selectTokenPoolRef(sealed, sel, program, on(legacyMint, 1))
		require.NoError(t, err)
		require.Equal(t, "CLLCCIP", got.Qualifier)
		require.Equal(t, "1.6.0", got.Version.String())
	})
	t.Run("instance ref is ignored while the registry is on interface 1", func(t *testing.T) {
		got, err := selectTokenPoolRef(sealed, sel, program, on(migratedMint, 1))
		require.NoError(t, err)
		require.Equal(t, "1.6.0", got.Version.String())
	})
	t.Run("token without token ref falls back to the latest 1.x ref", func(t *testing.T) {
		got, err := selectTokenPoolRef(sealed, sel, program, on(unknownMint, 1))
		require.NoError(t, err)
		require.Equal(t, "1.6.0", got.Version.String())
	})
	t.Run("interface 2 without an instance ref under the token qualifier fails", func(t *testing.T) {
		_, err := selectTokenPoolRef(sealed, sel, program, on(legacyMint, 2))
		require.ErrorContains(t, err, "uses pool interface 2 but no 2.x ref")
		require.ErrorContains(t, err, "LEGACY")
	})
	t.Run("program only resolves to the latest ref", func(t *testing.T) {
		got, err := selectTokenPoolRef(sealed, sel, program, nil)
		require.NoError(t, err)
		require.Equal(t, "2.1.0", got.Version.String())
	})
	t.Run("unknown program", func(t *testing.T) {
		_, err := selectTokenPoolRef(sealed, sel, solana.NewWallet().PublicKey(), on(migratedMint, 2))
		require.ErrorContains(t, err, "no ref at the program address")
	})
	t.Run("only 2.x refs and interface 1", func(t *testing.T) {
		v2Only := datastore.NewMemoryDataStore()
		require.NoError(t, v2Only.Addresses().Add(ref(program, "BurnMintTokenPool", "CLLCCIP", "2.0.0")))
		_, err := selectTokenPoolRef(v2Only.Seal(), sel, program, on(legacyMint, 1))
		require.ErrorContains(t, err, "no 1.x ref at the program address")
	})
}

func TestTokenAdminRegistryInterfaceVersion(t *testing.T) {
	const size = 8 + 1 + 32*5 + 1 + 1 + 1 + 1
	tar := func(layout, interfaceVersion uint8, length int) []byte {
		data := make([]byte, length)
		data[8] = layout
		if length > size-2 {
			data[size-2] = interfaceVersion
		}
		return data
	}

	got, err := tokenAdminRegistryInterfaceVersion(tar(4, 2, size))
	require.NoError(t, err)
	require.Equal(t, uint8(2), got)

	got, err = tokenAdminRegistryInterfaceVersion(tar(3, 1, size-1))
	require.NoError(t, err)
	require.Equal(t, uint8(1), got)

	got, err = tokenAdminRegistryInterfaceVersion(tar(2, 0, size-2))
	require.NoError(t, err)
	require.Equal(t, uint8(1), got, "layouts before 3 serve interface 1")

	_, err = tokenAdminRegistryInterfaceVersion(tar(4, 0, size-2))
	require.ErrorContains(t, err, "too short for interface_version")
	_, err = tokenAdminRegistryInterfaceVersion(make([]byte, 8))
	require.ErrorContains(t, err, "too short")
}

func TestPoolConfigMint(t *testing.T) {
	program := solana.NewWallet().PublicKey()
	mint := solana.NewWallet().PublicKey()
	configPDA, err := tokens.TokenPoolConfigAddress(mint, program)
	require.NoError(t, err)

	var state burnmint_token_pool.State
	state.Config.Mint = mint
	var buf bytes.Buffer
	require.NoError(t, bin.NewBorshEncoder(&buf).Encode(&state))

	got, err := poolConfigMint(configPDA, program, buf.Bytes())
	require.NoError(t, err)
	require.Equal(t, mint, got)

	_, err = poolConfigMint(configPDA, solana.NewWallet().PublicKey(), buf.Bytes())
	require.ErrorContains(t, err, "is not the pool config")
}
