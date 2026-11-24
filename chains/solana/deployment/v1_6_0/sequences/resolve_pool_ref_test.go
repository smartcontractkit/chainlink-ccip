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

	t.Run("migrated token resolves to its instance ref", func(t *testing.T) {
		got, err := selectTokenPoolRef(sealed, sel, program, &migratedMint)
		require.NoError(t, err)
		require.Equal(t, "CCIP-BnM", got.Qualifier)
		require.Equal(t, "2.0.0", got.Version.String())
	})
	t.Run("token without instance ref resolves to the latest 1.x ref", func(t *testing.T) {
		got, err := selectTokenPoolRef(sealed, sel, program, &legacyMint)
		require.NoError(t, err)
		require.Equal(t, "CLLCCIP", got.Qualifier)
		require.Equal(t, "1.6.0", got.Version.String())
	})
	t.Run("token without token ref falls back to the latest 1.x ref", func(t *testing.T) {
		got, err := selectTokenPoolRef(sealed, sel, program, &unknownMint)
		require.NoError(t, err)
		require.Equal(t, "1.6.0", got.Version.String())
	})
	t.Run("program only resolves to the latest ref", func(t *testing.T) {
		got, err := selectTokenPoolRef(sealed, sel, program, nil)
		require.NoError(t, err)
		require.Equal(t, "2.1.0", got.Version.String())
	})
	t.Run("unknown program", func(t *testing.T) {
		_, err := selectTokenPoolRef(sealed, sel, solana.NewWallet().PublicKey(), &migratedMint)
		require.ErrorContains(t, err, "no ref at the program address")
	})
	t.Run("only 2.x refs and no instance ref", func(t *testing.T) {
		v2Only := datastore.NewMemoryDataStore()
		require.NoError(t, v2Only.Addresses().Add(ref(program, "BurnMintTokenPool", "CLLCCIP", "2.0.0")))
		_, err := selectTokenPoolRef(v2Only.Seal(), sel, program, &legacyMint)
		require.ErrorContains(t, err, "no ref under the token qualifier")
	})
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
