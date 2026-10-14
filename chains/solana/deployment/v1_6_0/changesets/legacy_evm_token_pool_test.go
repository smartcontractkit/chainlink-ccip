package changesets

import (
	"errors"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/shared"
	solanastateview "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/shared/stateview/solana"
)

const (
	poolTEST      = "0x00000000000000000000000000000000000000a1"
	poolOTHER     = "0x00000000000000000000000000000000000000a2"
	poolTEST2     = "0x00000000000000000000000000000000000000a3"
	poolNoToken   = "0x00000000000000000000000000000000000000a4"
	poolNoDecimal = "0x00000000000000000000000000000000000000a5"
	poolUSDC      = "0x00000000000000000000000000000000000000a6"
	tokenTEST     = "0x00000000000000000000000000000000000000b1"
	tokenOTHER    = "0x00000000000000000000000000000000000000b2"
	tokenUSDC     = "0x00000000000000000000000000000000000000b6"
)

// fakeEVMRefResolver stands in for the EVM family's registered TokenRefResolver: the on-chain
// typeAndVersion and token of each pool, and the ERC20 symbol of each token.
type fakeEVMRefResolver struct {
	pools   map[string]struct{ version, token string } // token "" means getToken fails
	symbols map[string]string
}

func (f fakeEVMRefResolver) ResolveTokenPoolRef(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, _ datastore.DataStore, sel uint64, address string) (datastore.AddressRef, error) {
	p, ok := f.pools[address]
	if !ok {
		return datastore.AddressRef{}, errors.New("no contract")
	}
	qualifier := p.token
	if qualifier == "" {
		qualifier = address + "-BurnMintTokenPool" // the resolver's placeholder when getToken fails
	}
	return datastore.AddressRef{ChainSelector: sel, Address: address, Version: semver.MustParse(p.version), Qualifier: qualifier}, nil
}

func (f fakeEVMRefResolver) ResolveTokenRef(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, _ datastore.DataStore, _ uint64, address string) (datastore.AddressRef, error) {
	return datastore.AddressRef{Address: address, Qualifier: f.symbols[address]}, nil
}

// fakeEVMTokenAdapter stands in for the registered 1.5.1 EVM token adapter: DeriveTokenDecimals
// reads the token when one is given, and otherwise the pool.
type fakeEVMTokenAdapter struct {
	tokensapi.TokenAdapter
	poolDecimals  map[string]uint8 // missing: the pool can't report them
	tokenDecimals uint8
}

func (f fakeEVMTokenAdapter) DeriveTokenDecimals(_ cldf.Environment, _ uint64, poolRef datastore.AddressRef, token []byte) (uint8, error) {
	if token != nil {
		return f.tokenDecimals, nil
	}
	d, ok := f.poolDecimals[poolRef.Address]
	if !ok {
		return 0, errors.New("getTokenDecimals reverted")
	}
	return d, nil
}

func evmPoolRef(sel uint64, poolType cldf.ContractType, address, version string, labels ...string) datastore.AddressRef {
	ref := datastore.AddressRef{ChainSelector: sel, Type: datastore.ContractType(poolType), Address: address, Version: semver.MustParse(version), Qualifier: address}
	if len(labels) > 0 {
		ref.Labels = datastore.NewLabelSet(labels...)
	}
	return ref
}

func TestLegacyEVMPoolTiers(t *testing.T) {
	t.Parallel()
	v := func(s string) semver.Version { return *semver.MustParse(s) }
	tiers, bySymbol := legacyEVMPoolTiers(shared.BurnMintTokenPool, v("1.6.1"))
	require.Equal(t, [][]string{{"1.6.1"}, {"1.5.1"}, {"1.5.0"}}, tiers)
	require.True(t, bySymbol)
	tiers, _ = legacyEVMPoolTiers(shared.BurnMintTokenPool, v("1.5.1"))
	require.Equal(t, [][]string{{"1.5.1"}, {"1.5.0"}}, tiers)
	tiers, _ = legacyEVMPoolTiers(shared.LockReleaseTokenPool, v("1.6.1"))
	require.Equal(t, [][]string{{"1.6.1"}, {"1.5.1"}}, tiers)
	tiers, _ = legacyEVMPoolTiers("BurnMintFastTransferTokenPool", v("1.6.1"))
	require.Equal(t, [][]string{{"1.6.1", "1.6.3-dev"}}, tiers)
	tiers, bySymbol = legacyEVMPoolTiers(shared.USDCTokenPool, v("1.6.2"))
	require.Equal(t, [][]string{{"1.6.2"}}, tiers)
	require.False(t, bySymbol)
	tiers, _ = legacyEVMPoolTiers("Unknown", v("1.5.1"))
	require.Nil(t, tiers)
	require.Len(t, legacyEVMPoolRefTypes("HybridLockReleaseUSDCTokenPool"), 2)
}

func TestEVMAddressBytes(t *testing.T) {
	t.Parallel()
	b, err := evmAddressBytes(tokenTEST)
	require.NoError(t, err)
	require.Len(t, b, 20)
	_, err = evmAddressBytes(poolTEST + "-BurnMintTokenPool")
	require.Error(t, err)
	require.Equal(t, append(make([]byte, 30), 1, 2), leftPad32([]byte{1, 2}))
}

// The token adapter registry is global, so the unregistered case and the registered cases run in
// order in one test.
func TestLegacyEVMTokenPoolLookup(t *testing.T) {
	evmSel := chainsel.TEST_90000001.Selector
	env := func(refs ...datastore.AddressRef) cldf.Environment {
		ds := datastore.NewMemoryDataStore()
		for _, r := range refs {
			require.NoError(t, ds.Addresses().Add(r))
		}
		return cldf.Environment{
			DataStore:   ds.Seal(),
			BlockChains: cldf_chain.NewBlockChainsFromSlice([]cldf_chain.BlockChain{cldf_evm.Chain{Selector: evmSel}}),
		}
	}
	v151 := *semver.MustParse("1.5.1")
	e := env(
		evmPoolRef(evmSel, shared.BurnMintTokenPool, poolTEST, "1.5.1"),
		evmPoolRef(evmSel, shared.BurnMintTokenPool, poolOTHER, "1.5.1"),
		evmPoolRef(evmSel, shared.BurnMintTokenPool, poolTEST2, "1.5.1", shared.SupersededLabel),
		evmPoolRef(evmSel, shared.USDCTokenPool, poolUSDC, "1.6.2"),
	)

	_, _, err := FindLegacyEVMTokenPool(e, evmSel, "TEST", shared.BurnMintTokenPool, v151)
	require.ErrorContains(t, err, "no token ref resolver registered for the EVM family")

	tokensapi.GetTokenAdapterRegistry().RegisterTokenRefResolver(chainsel.FamilyEVM, fakeEVMRefResolver{
		pools: map[string]struct{ version, token string }{
			poolTEST: {"1.5.1", tokenTEST}, poolOTHER: {"1.5.1", tokenOTHER}, poolTEST2: {"1.5.1", tokenTEST},
			poolNoToken: {"1.5.1", ""}, poolNoDecimal: {"1.5.1", tokenTEST}, poolUSDC: {"1.6.2", tokenUSDC},
		},
		symbols: map[string]string{tokenTEST: "TEST", tokenOTHER: "OTHER", tokenUSDC: "USDC"},
	})

	// matched by the token's on-chain symbol; the superseded pool for the same token is skipped
	pool, found, err := FindLegacyEVMTokenPool(e, evmSel, "TEST", shared.BurnMintTokenPool, v151)
	require.NoError(t, err)
	require.True(t, found)
	require.Equal(t, poolTEST, pool.PoolRef.Address)
	wantToken, _ := evmAddressBytes(tokenTEST)
	require.Equal(t, wantToken, pool.TokenAddress)

	_, found, err = FindLegacyEVMTokenPool(e, evmSel, "MISSING", shared.BurnMintTokenPool, v151)
	require.NoError(t, err)
	require.False(t, found)
	// the on-chain version must match too
	_, found, err = FindLegacyEVMTokenPool(e, evmSel, "TEST", shared.BurnMintTokenPool, *semver.MustParse("1.5.0"))
	require.NoError(t, err)
	require.False(t, found)
	// USDC pools are found by version alone
	pool, found, err = FindLegacyEVMTokenPool(e, evmSel, "anything", shared.USDCTokenPool, *semver.MustParse("1.6.2"))
	require.NoError(t, err)
	require.True(t, found)
	require.Equal(t, poolUSDC, pool.PoolRef.Address)

	// two live pools for the same symbol and version: refuse to guess
	_, _, err = FindLegacyEVMTokenPool(env(
		evmPoolRef(evmSel, shared.BurnMintTokenPool, poolTEST, "1.5.1"),
		evmPoolRef(evmSel, shared.BurnMintTokenPool, poolTEST2, "1.5.1"),
	), evmSel, "TEST", shared.BurnMintTokenPool, v151)
	require.ErrorContains(t, err, "ambiguous")
	// a pool whose getToken fails is an error, as in the legacy loader
	_, _, err = FindLegacyEVMTokenPool(env(evmPoolRef(evmSel, shared.BurnMintTokenPool, poolNoToken, "1.5.1")), evmSel, "TEST", shared.BurnMintTokenPool, v151)
	require.ErrorContains(t, err, "failed to get token address")

	// decimals
	_, err = GetLegacyEVMTokenPoolDecimals(e, evmSel, pool, false)
	require.ErrorContains(t, err, "no 1.5.1 token adapter registered")
	tokensapi.GetTokenAdapterRegistry().RegisterTokenAdapter(chainsel.FamilyEVM, decimalsAdapterVersion, fakeEVMTokenAdapter{
		poolDecimals:  map[string]uint8{poolTEST: 18},
		tokenDecimals: 6,
	})
	testPool := LegacyEVMTokenPool{PoolRef: evmPoolRef(evmSel, shared.BurnMintTokenPool, poolTEST, "1.5.1"), TokenAddress: wantToken}
	noDecimals := LegacyEVMTokenPool{PoolRef: evmPoolRef(evmSel, shared.BurnMintTokenPool, poolNoDecimal, "1.5.1"), TokenAddress: wantToken}
	d, err := GetLegacyEVMTokenPoolDecimals(e, evmSel, testPool, false)
	require.NoError(t, err)
	require.Equal(t, uint8(18), d, "the pool's decimals win over the token's")
	_, err = GetLegacyEVMTokenPoolDecimals(e, evmSel, noDecimals, false)
	require.ErrorContains(t, err, "failed to get token decimals", "without the fallback, never read the token")
	d, err = GetLegacyEVMTokenPoolDecimals(e, evmSel, noDecimals, true)
	require.NoError(t, err)
	require.Equal(t, uint8(6), d)

	// the v0_1_1 EVMRemoteConfig path end to end
	cfg := EVMRemoteConfig{TokenSymbol: "TEST", PoolType: shared.BurnMintTokenPool, PoolVersion: v151}
	state := solanastateview.CCIPOnChainState{}
	require.NoError(t, cfg.Validate(e, state, evmSel))
	missing := cfg
	missing.TokenSymbol = "MISSING"
	require.ErrorContains(t, missing.Validate(e, state, evmSel), "token pool does not exist")
	require.ErrorContains(t, cfg.Validate(e, state, chainsel.TEST_90000002.Selector), "does not exist in environment")
	remote, err := getOnChainEVMPoolConfig(e, state, evmSel, cfg)
	require.NoError(t, err)
	require.Equal(t, leftPad32(wantToken), remote.TokenAddress.Address)
	wantPool, _ := evmAddressBytes(poolTEST)
	require.Equal(t, wantPool, remote.PoolAddresses[0].Address)
	require.Equal(t, uint8(18), remote.Decimals)
}
