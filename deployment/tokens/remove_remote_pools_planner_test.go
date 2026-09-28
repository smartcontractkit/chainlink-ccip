package tokens

import (
	"context"
	"errors"
	"testing"

	"github.com/Masterminds/semver/v3"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_aptos "github.com/smartcontractkit/chainlink-deployments-framework/chain/aptos"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

// The reverse-pass unit tests use Aptos selectors with an identity address normalizer so the mocks
// can use readable addresses. No other test in this package registers an Aptos normalizer, and the
// normalizer registry is first-write-wins, so this registration cannot collide with other tests.
var (
	reverseTestLocalSel  = chainsel.APTOS_LOCALNET.Selector
	reverseTestRemoteSel = chainsel.APTOS_TESTNET.Selector
	reverseTestV2_0_0    = semver.MustParse("2.0.0")
)

type reverseTestIdentityNormalizer struct{}

func (reverseTestIdentityNormalizer) NormalizeAddress(a string) (string, error) { return a, nil }

func (reverseTestIdentityNormalizer) BytesToString(b []byte) (string, error) { return string(b), nil }

func (reverseTestIdentityNormalizer) StringToBytes(a string) ([]byte, error) { return []byte(a), nil }

// reverseTestAdapter is a minimal TokenAdapter + TokenPoolMigrator + RemotePoolRemover. Addresses
// are plain strings (the identity normalizer maps them to bytes 1:1). It serves as the local adapter
// (GetSupportedChains/GetRemoteToken) and as the remote adapters (GetRemotePools/RemoveRemotePools).
type reverseTestAdapter struct {
	supportedChains []uint64
	remoteToken     string
	remotePools     map[string][]string // pool address -> remote pools it lists for the other chain
	removed         *[]RemoveRemotePoolsSequenceInput
}

var (
	_ TokenAdapter      = (*reverseTestAdapter)(nil)
	_ TokenPoolMigrator = (*reverseTestAdapter)(nil)
	_ RemotePoolRemover = (*reverseTestAdapter)(nil)
)

func (a *reverseTestAdapter) AddressRefToBytes(ref datastore.AddressRef) ([]byte, error) {
	return []byte(ref.Address), nil
}

func (a *reverseTestAdapter) DeriveTokenPoolCounterpart(_ cldf.Environment, _ uint64, tokenPool []byte, _ []byte) ([]byte, error) {
	return tokenPool, nil
}

func (a *reverseTestAdapter) GetSupportedChains(_ cldf.Environment, _ uint64, _, _ []byte) ([]uint64, error) {
	return a.supportedChains, nil
}

func (a *reverseTestAdapter) GetRemoteToken(_ cldf.Environment, _ uint64, _, _ []byte, _ uint64) ([]byte, error) {
	return []byte(a.remoteToken), nil
}

func (a *reverseTestAdapter) GetRemotePools(_ cldf.Environment, _ uint64, poolAddr, _ []byte, _ uint64) ([][]byte, error) {
	out := [][]byte{}
	for _, p := range a.remotePools[string(poolAddr)] {
		out = append(out, []byte(p))
	}
	return out, nil
}

func (a *reverseTestAdapter) RemoveRemotePools() *cldf_ops.Sequence[RemoveRemotePoolsSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"reverse-test:remove-remote-pools",
		semver.MustParse("1.0.0"),
		"Records remote pool removals for the reverse-pass unit tests",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, in RemoveRemotePoolsSequenceInput) (sequences.OnChainOutput, error) {
			*a.removed = append(*a.removed, in)
			return sequences.OnChainOutput{}, nil
		},
	)
}

// Unused TokenAdapter methods.
func (a *reverseTestAdapter) ConfigureTokenForTransfersSequence() *cldf_ops.Sequence[ConfigureTokenForTransfersInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}

func (a *reverseTestAdapter) DeriveTokenAddress(cldf.Environment, uint64, datastore.AddressRef) (string, error) {
	return "", errors.New("not implemented")
}

func (a *reverseTestAdapter) DeriveTokenDecimals(cldf.Environment, uint64, datastore.AddressRef, []byte) (uint8, error) {
	return 0, errors.New("not implemented")
}

func (a *reverseTestAdapter) ManualRegistration() *cldf_ops.Sequence[ManualRegistrationSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}

func (a *reverseTestAdapter) SetTokenPoolRateLimits() *cldf_ops.Sequence[TPRLRemotes, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}

func (a *reverseTestAdapter) DeployToken() *cldf_ops.Sequence[DeployTokenInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}

func (a *reverseTestAdapter) DeployTokenVerify(cldf.Environment, DeployTokenInput) error { return nil }

func (a *reverseTestAdapter) DeployTokenPoolForToken() *cldf_ops.Sequence[DeployTokenPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}

func (a *reverseTestAdapter) UpdateAuthorities() *cldf_ops.Sequence[UpdateAuthoritiesInput, sequences.OnChainOutput, *cldf.Environment] {
	return nil
}

func (a *reverseTestAdapter) MigrateLockReleasePoolLiquiditySequence() *cldf_ops.Sequence[MigrateLockReleasePoolLiquidityInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}

// reverseTestTAR is a TokenAdminRegistryManager that reports a fixed active pool.
type reverseTestTAR struct{ activePool string }

func (m *reverseTestTAR) GetActivePool(cldf.Environment, uint64, datastore.AddressRef, ...datastore.AddressRef) ([]byte, error) {
	return []byte(m.activePool), nil
}

func (m *reverseTestTAR) GetTokenAdminRegistryRef(cldf.Environment, uint64) (datastore.AddressRef, error) {
	return datastore.AddressRef{}, nil
}

func (m *reverseTestTAR) UnregisterToken() *cldf_ops.Sequence[UnregisterTokenSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}

// reverseTestPool describes a peer pool on the remote chain.
type reverseTestPool struct {
	address string
	version *semver.Version
	lists   []string // local pools this peer pool lists as remotes
}

// reverseTestSetup varies the environment of runReverseTest; the zero value is a fully supported peer.
type reverseTestSetup struct {
	remoteNotLoaded      bool // the remote chain is not loaded in the environment
	noTAR                bool // no TAR reader is registered for the remote family
	skipUnsupportedPeers bool
}

// runReverseTest plans and executes the reverse pass for local pool A_OLD against a single remote chain
// whose pools are described by peerPools. pairedPool is the peer pool named by the remote entry;
// activePool is the peer's TAR-active pool. It returns the recorded removals and the error.
func runReverseTest(t *testing.T, peerPools []reverseTestPool, pairedPool, activePool string, setup *reverseTestSetup) ([]RemoveRemotePoolsSequenceInput, error) {
	t.Helper()

	family, err := chainsel.GetSelectorFamily(reverseTestRemoteSel)
	require.NoError(t, err)
	deploy.GetAddressNormalizerRegistry().RegisterAddressNormalizer(family, reverseTestIdentityNormalizer{})
	if setup == nil {
		setup = &reverseTestSetup{}
	}

	// Seed the datastore with the peer pools and the peer token so ref resolution hits the cache.
	ds := datastore.NewMemoryDataStore()
	for _, p := range peerPools {
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: reverseTestRemoteSel,
			Address:       p.address,
			Type:          datastore.ContractType("TestTokenPool"),
			Version:       p.version,
			Qualifier:     p.address,
		}))
	}
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: reverseTestRemoteSel,
		Address:       "TOKEN_REMOTE",
		Type:          datastore.ContractType("TestToken"),
		Version:       semver.MustParse("1.0.0"),
	}))

	loadedChains := map[uint64]cldf_chain.BlockChain{}
	if !setup.remoteNotLoaded {
		loadedChains[reverseTestRemoteSel] = cldf_aptos.Chain{Selector: reverseTestRemoteSel}
	}

	lggr := logger.Test(t)
	e := cldf.Environment{
		Logger:           lggr,
		OperationsBundle: cldf_ops.NewBundle(func() context.Context { return t.Context() }, lggr, cldf_ops.NewMemoryReporter()),
		DataStore:        ds.Seal(),
		BlockChains:      cldf_chain.NewBlockChains(loadedChains),
	}

	removed := []RemoveRemotePoolsSequenceInput{}
	remotePools := map[string][]string{}
	for _, p := range peerPools {
		remotePools[p.address] = p.lists
	}
	local := &reverseTestAdapter{supportedChains: []uint64{reverseTestRemoteSel}, remoteToken: "TOKEN_REMOTE"}
	remoteV2 := &reverseTestAdapter{remotePools: remotePools, removed: &removed}

	reg := newTokenAdapterRegistry()
	reg.RegisterTokenAdapter(family, reverseTestV2_0_0, remoteV2)
	if !setup.noTAR {
		reg.RegisterTokenAdminRegistryManager(family, &reverseTestTAR{activePool: activePool})
	}

	planner := newRemoveRemotePoolsPlanner(e, reg)
	entry := &removeRemotePoolsEntry{
		pool:       RemoveRemotePoolsPerPool{SkipUnsupportedPeers: setup.skipUnsupportedPeers},
		label:      "A_OLD",
		selector:   reverseTestLocalSel,
		adapter:    local,
		migrator:   local,
		poolRef:    datastore.AddressRef{ChainSelector: reverseTestLocalSel, Address: "A_OLD", Version: reverseTestV2_0_0},
		tokenRef:   datastore.AddressRef{ChainSelector: reverseTestLocalSel, Address: "TOKEN_LOCAL"},
		poolBytes:  []byte("A_OLD"),
		tokenBytes: []byte("TOKEN_LOCAL"),
	}
	if err := planner.planReverse(entry, []RemotePoolToRemove{{Selector: reverseTestRemoteSel, Remote: datastore.AddressRef{Address: pairedPool}}}); err != nil {
		return nil, err
	}
	_, _, err = planner.execute()
	require.NoError(t, err)
	return removed, nil
}

// removedFrom returns the peer pool addresses that received a removal, requiring that each removal
// targeted only the local pool A_OLD.
func removedFrom(t *testing.T, removed []RemoveRemotePoolsSequenceInput) []string {
	t.Helper()
	pools := []string{}
	for _, r := range removed {
		require.Equal(t, reverseTestRemoteSel, r.Selector)
		require.Len(t, r.RemotePoolsToRemove, 1)
		require.Equal(t, "A_OLD", r.RemotePoolsToRemove[0].Remote.Address, "only the local pool being torn down may be removed")
		pools = append(pools, r.TokenPoolRef.Address)
	}
	return pools
}

// Test target dedupe. When the paired pool is the active pool (a peer that was never upgraded), it is
// cleaned exactly once; when they differ, both are cleaned.
func TestRemoveRemotePoolsReverse_TargetDedupe(t *testing.T) {
	t.Run("paired pool equals active pool: cleaned once", func(t *testing.T) {
		removed, err := runReverseTest(t, []reverseTestPool{
			{address: "B_OLD", version: reverseTestV2_0_0, lists: []string{"A_NEW", "A_OLD"}},
		}, "B_OLD", "B_OLD", nil)
		require.NoError(t, err)
		require.Equal(t, []string{"B_OLD"}, removedFrom(t, removed))
	})

	t.Run("paired pool differs from active pool: both cleaned", func(t *testing.T) {
		removed, err := runReverseTest(t, []reverseTestPool{
			{address: "B_OLD", version: reverseTestV2_0_0, lists: []string{"A_NEW", "A_OLD"}},
			{address: "B_NEW", version: reverseTestV2_0_0, lists: []string{"A_NEW", "A_OLD"}},
		}, "B_OLD", "B_NEW", nil)
		require.NoError(t, err)
		require.ElementsMatch(t, []string{"B_NEW", "B_OLD"}, removedFrom(t, removed))
	})
}

// Test that a zero-address remote entry is skipped as a reverse target: it names no peer pool, so
// only the peer's active pool is cleaned instead of failing to resolve the zero address.
func TestRemoveRemotePoolsReverse_ZeroAddressPairedPool(t *testing.T) {
	removed, err := runReverseTest(t, []reverseTestPool{
		{address: "B_NEW", version: reverseTestV2_0_0, lists: []string{"A_NEW", "A_OLD"}},
	}, "\x00\x00\x00\x00", "B_NEW", nil)
	require.NoError(t, err)
	require.Equal(t, []string{"B_NEW"}, removedFrom(t, removed))
}

// Test that a remote entry naming an unknown pool fails planning with the lane-only removal hint,
// and that nothing is removed.
func TestRemoveRemotePoolsReverse_UnresolvablePairedPool(t *testing.T) {
	removed, err := runReverseTest(t, []reverseTestPool{
		{address: "B_NEW", version: reverseTestV2_0_0, lists: []string{"A_NEW", "A_OLD"}},
	}, "NOT_A_POOL", "B_NEW", nil)
	require.ErrorContains(t, err, "remove it from this pool with a lane-only removal")
	require.Empty(t, removed)
}

// Test that peers this tooling cannot process fail the reverse pass, naming skipUnsupportedPeers, and
// that with skipUnsupportedPeers they are skipped without removing anything.
func TestRemoveRemotePoolsReverse_UnsupportedPeers(t *testing.T) {
	supportedPeer := []reverseTestPool{{address: "B_NEW", version: reverseTestV2_0_0, lists: []string{"A_OLD"}}}
	cases := []struct {
		name      string
		peerPools []reverseTestPool
		setup     reverseTestSetup
	}{
		{name: "remote chain not loaded", peerPools: supportedPeer, setup: reverseTestSetup{remoteNotLoaded: true}},
		{name: "no TAR reader for the remote family", peerPools: supportedPeer, setup: reverseTestSetup{noTAR: true}},
		{name: "no adapter for the peer pool version", peerPools: []reverseTestPool{{address: "B_NEW", version: semver.MustParse("9.9.9"), lists: []string{"A_OLD"}}}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			removed, err := runReverseTest(t, tc.peerPools, "B_NEW", "B_NEW", &tc.setup)
			require.ErrorContains(t, err, "skipUnsupportedPeers")
			require.Empty(t, removed)

			tc.setup.skipUnsupportedPeers = true
			removed, err = runReverseTest(t, tc.peerPools, "B_NEW", "B_NEW", &tc.setup)
			require.NoError(t, err)
			require.Empty(t, removed)
		})
	}
}
