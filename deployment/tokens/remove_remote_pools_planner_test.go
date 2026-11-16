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
	// reverseTestRemoteVersion is registered only for the peer adapter, so the local and remote
	// adapters do not collide in the first-write-wins (family, version) registry.
	reverseTestRemoteVersion = semver.MustParse("1.6.1")
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
	localToken      string
	remotePools     map[string][]string // pool address -> remote pools it lists for the other chain
	// remotePoolsBySelector, when set, takes precedence and keys the remote pools by the remote
	// selector as well, so one pool can list different pools per remote chain.
	remotePoolsBySelector map[uint64]map[string][]string
	removed               *[]RemoveRemotePoolsSequenceInput
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

func (a *reverseTestAdapter) GetRemotePools(_ cldf.Environment, _ uint64, poolAddr, _ []byte, remoteSelector uint64) ([][]byte, error) {
	if a.remotePoolsBySelector != nil {
		out := [][]byte{}
		for _, p := range a.remotePoolsBySelector[remoteSelector][string(poolAddr)] {
			out = append(out, []byte(p))
		}
		return out, nil
	}
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

// reverseTestWipeAdapter embeds reverseTestAdapter and additionally implements
// RemoteChainConfigRemover, recording the chain-config removals it is asked to perform.
type reverseTestWipeAdapter struct {
	*reverseTestAdapter
	wiped       *[]RemoveRemoteChainConfigsSequenceInput
	unsupported bool
}

var _ RemoteChainConfigRemover = (*reverseTestWipeAdapter)(nil)

func (a *reverseTestWipeAdapter) SupportsRemoteChainConfigRemoval() bool { return !a.unsupported }

func (a *reverseTestWipeAdapter) RemoveRemoteChainConfigs() *cldf_ops.Sequence[RemoveRemoteChainConfigsSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"reverse-test:remove-remote-chain-configs",
		semver.MustParse("1.0.0"),
		"Records remote chain config removals for the wipe unit tests",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, in RemoveRemoteChainConfigsSequenceInput) (sequences.OnChainOutput, error) {
			*a.wiped = append(*a.wiped, in)
			return sequences.OnChainOutput{}, nil
		},
	)
}

// Unused TokenAdapter methods.
func (a *reverseTestAdapter) ConfigureTokenForTransfersSequence() *cldf_ops.Sequence[ConfigureTokenForTransfersInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}

func (a *reverseTestAdapter) DeriveTokenAddress(cldf.Environment, uint64, datastore.AddressRef) (string, error) {
	if a.localToken == "" {
		return "", errors.New("not implemented")
	}
	return a.localToken, nil
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

// reverseTestReaderOnlyTAR is a TokenAdminRegistryReader that cannot unregister.
type reverseTestReaderOnlyTAR struct{ activePool string }

func (m *reverseTestReaderOnlyTAR) GetActivePool(cldf.Environment, uint64, datastore.AddressRef, ...datastore.AddressRef) ([]byte, error) {
	return []byte(m.activePool), nil
}

func (m *reverseTestReaderOnlyTAR) GetTokenAdminRegistryRef(cldf.Environment, uint64) (datastore.AddressRef, error) {
	return datastore.AddressRef{}, nil
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
	readerOnlyTAR        bool // the remote family registers a reader-only TAR
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
	switch {
	case setup.noTAR:
	case setup.readerOnlyTAR:
		reg.RegisterTokenAdminRegistryReader(family, &reverseTestReaderOnlyTAR{activePool: activePool})
	default:
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

// Test that a peer token with no TAR-active pool does not fail the reverse pass: the active pool is
// dropped as a target and the pool named by the remote entry is still cleaned. When the remote
// entry is also the zero address there is no target at all, so nothing is cleaned and no error is
// returned.
func TestRemoveRemotePoolsReverse_NoActivePool(t *testing.T) {
	t.Run("named pool is still cleaned", func(t *testing.T) {
		removed, err := runReverseTest(t, []reverseTestPool{
			{address: "B_NEW", version: reverseTestV2_0_0, lists: []string{"A_NEW", "A_OLD"}},
		}, "B_NEW", "", nil)
		require.NoError(t, err)
		require.Equal(t, []string{"B_NEW"}, removedFrom(t, removed))
	})

	t.Run("zero-address remote with no active pool cleans nothing", func(t *testing.T) {
		removed, err := runReverseTest(t, nil, "\x00\x00\x00\x00", "", nil)
		require.NoError(t, err)
		require.Empty(t, removed)
	})
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

// Test that the reverse pass only reads peer TARs, so a peer family with a reader-only TAR is cleaned.
func TestRemoveRemotePoolsReverse_ReaderOnlyPeerTAR(t *testing.T) {
	removed, err := runReverseTest(t, []reverseTestPool{
		{address: "B_NEW", version: reverseTestV2_0_0, lists: []string{"A_OLD"}},
	}, "B_NEW", "B_NEW", &reverseTestSetup{readerOnlyTAR: true})
	require.NoError(t, err)
	require.Equal(t, []string{"B_NEW"}, removedFrom(t, removed))
}

// Test that the TAR unregister needs a manager only when it has to write: a reader-only family can
// deactivate a pool the TAR no longer points at, but fails at planning when the pool is still active.
func TestRemoveRemotePoolsPlanUnregister_ReaderOnlyTAR(t *testing.T) {
	family, err := chainsel.GetSelectorFamily(reverseTestLocalSel)
	require.NoError(t, err)
	deploy.GetAddressNormalizerRegistry().RegisterAddressNormalizer(family, reverseTestIdentityNormalizer{})

	planUnregister := func(t *testing.T, tar TokenAdminRegistryReader) (*removeRemotePoolsPlanner, error) {
		reg := newTokenAdapterRegistry()
		if manager, ok := tar.(TokenAdminRegistryManager); ok {
			reg.RegisterTokenAdminRegistryManager(family, manager)
		} else {
			reg.RegisterTokenAdminRegistryReader(family, tar)
		}
		lggr := logger.Test(t)
		planner := newRemoveRemotePoolsPlanner(cldf.Environment{Logger: lggr}, reg)
		return planner, planner.planUnregister(&removeRemotePoolsEntry{
			label:     "A_OLD",
			selector:  reverseTestLocalSel,
			family:    family,
			poolRef:   datastore.AddressRef{ChainSelector: reverseTestLocalSel, Address: "A_OLD"},
			poolBytes: []byte("A_OLD"),
		})
	}

	t.Run("TAR points at another pool: nothing to unregister", func(t *testing.T) {
		planner, err := planUnregister(t, &reverseTestReaderOnlyTAR{activePool: "A_NEW"})
		require.NoError(t, err)
		require.Empty(t, planner.unregisters)
	})

	t.Run("TAR points at this pool: fails at planning", func(t *testing.T) {
		planner, err := planUnregister(t, &reverseTestReaderOnlyTAR{activePool: "A_OLD"})
		require.ErrorContains(t, err, "does not support unregistering tokens")
		require.Empty(t, planner.unregisters)
	})

	t.Run("TAR points at this pool with a manager: unregister queued", func(t *testing.T) {
		planner, err := planUnregister(t, &reverseTestTAR{activePool: "A_OLD"})
		require.NoError(t, err)
		require.Len(t, planner.unregisters, 1)
	})
}

// Test that a peer listing the local pool more than once (e.g. in both EVM encodings, which decode to
// the same address) gets a single removal: the remover removes every stored encoding of the address,
// and a duplicate entry would queue duplicate removals that revert under MCMS.
func TestRemoveRemotePoolsReverse_PeerListsPoolTwice(t *testing.T) {
	removed, err := runReverseTest(t, []reverseTestPool{
		{address: "B_NEW", version: reverseTestV2_0_0, lists: []string{"A_OLD", "A_NEW", "A_OLD"}},
	}, "B_NEW", "B_NEW", nil)
	require.NoError(t, err)
	require.Equal(t, []string{"B_NEW"}, removedFrom(t, removed))
}

// Test that removals queued for the same target pool run as one removal call (so a family that
// rewrites the whole list per call cannot undo its own removals under MCMS), that pools serving
// different tokens are not merged, and that peers are cleaned before entries' own pools.
func TestRemoveRemotePoolsPlanner_GroupsRemovalsPerPool(t *testing.T) {
	family, err := chainsel.GetSelectorFamily(reverseTestRemoteSel)
	require.NoError(t, err)
	deploy.GetAddressNormalizerRegistry().RegisterAddressNormalizer(family, reverseTestIdentityNormalizer{})

	lggr := logger.Test(t)
	env := cldf.Environment{
		Logger:           lggr,
		OperationsBundle: cldf_ops.NewBundle(func() context.Context { return t.Context() }, lggr, cldf_ops.NewMemoryReporter()),
		BlockChains:      cldf_chain.NewBlockChains(nil),
	}
	removed := []RemoveRemotePoolsSequenceInput{}
	remover := &reverseTestAdapter{removed: &removed}
	removal := func(pool, token string, remotes ...string) RemoveRemotePoolsSequenceInput {
		in := RemoveRemotePoolsSequenceInput{
			Selector:     reverseTestRemoteSel,
			TokenPoolRef: datastore.AddressRef{Address: pool},
			TokenRef:     datastore.AddressRef{Address: token},
		}
		for _, r := range remotes {
			in.RemotePoolsToRemove = append(in.RemotePoolsToRemove, RemotePoolToRemove{Selector: reverseTestLocalSel, Remote: datastore.AddressRef{Address: r}})
		}
		return in
	}
	remotesOf := func(in RemoveRemotePoolsSequenceInput) []string {
		out := []string{}
		for _, r := range in.RemotePoolsToRemove {
			out = append(out, r.Remote.Address)
		}
		return out
	}

	planner := newRemoveRemotePoolsPlanner(env, newTokenAdapterRegistry())
	require.NoError(t, planner.queueRemoval(remover, removal("P", "TOKEN", "X"), true))        // P is an entry's own pool
	require.NoError(t, planner.queueRemoval(remover, removal("Q", "TOKEN", "Y"), false))       // Q is only a peer
	require.NoError(t, planner.queueRemoval(remover, removal("P", "TOKEN", "Z"), false))       // another entry's reverse pass hits P
	require.NoError(t, planner.queueRemoval(remover, removal("P", "OTHER_TOKEN", "W"), false)) // same program, different token
	_, _, err = planner.execute()
	require.NoError(t, err)

	require.Len(t, removed, 3)
	require.Equal(t, "Q", removed[0].TokenPoolRef.Address, "a pool that is only a peer is cleaned first")
	require.Equal(t, "P", removed[1].TokenPoolRef.Address)
	require.Equal(t, "OTHER_TOKEN", removed[1].TokenRef.Address, "a different token's config on the same pool is its own call")
	require.Equal(t, "P", removed[2].TokenPoolRef.Address)
	require.Equal(t, "TOKEN", removed[2].TokenRef.Address)
	require.Equal(t, []string{"X", "Z"}, remotesOf(removed[2]), "both removals on P run as one call, after the peers")
}

// reverseTestOtherSel is a third Aptos selector, used to give the local pool a second supported
// remote chain whose remote-pool list is independent of reverseTestRemoteSel's.
var reverseTestOtherSel = chainsel.APTOS_MAINNET.Selector

// planPoolTestConfig configures runPlanPoolTest. The local pool is always A_OLD with token
// TOKEN_LOCAL; peers are on reverseTestRemoteSel with token TOKEN_REMOTE.
type planPoolTestConfig struct {
	entry                RemoveRemotePoolsPerPool
	extraEntries         []RemoveRemotePoolsPerPool
	localSupported       []uint64
	localRemotePools     map[uint64]map[string][]string // remote selector -> local pool -> remote pools it lists
	peerPools            []reverseTestPool
	localWithWipe        bool
	localWipeUnsupported bool
	remoteWithWipe       bool
	activePool           string
}

// runPlanPoolTest plans and executes a single pool entry against a memory datastore and identity
// normalizers, returning the recorded per-pool removals and chain-config wipes.
func runPlanPoolTest(t *testing.T, cfg planPoolTestConfig) ([]RemoveRemotePoolsSequenceInput, []RemoveRemoteChainConfigsSequenceInput, error) {
	t.Helper()

	localFamily, err := chainsel.GetSelectorFamily(reverseTestLocalSel)
	require.NoError(t, err)
	remoteFamily, err := chainsel.GetSelectorFamily(reverseTestRemoteSel)
	require.NoError(t, err)
	deploy.GetAddressNormalizerRegistry().RegisterAddressNormalizer(localFamily, reverseTestIdentityNormalizer{})
	deploy.GetAddressNormalizerRegistry().RegisterAddressNormalizer(remoteFamily, reverseTestIdentityNormalizer{})

	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: reverseTestLocalSel,
		Address:       "A_OLD",
		Type:          datastore.ContractType("TestTokenPool"),
		Version:       reverseTestV2_0_0,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: reverseTestLocalSel,
		Address:       "TOKEN_LOCAL",
		Type:          datastore.ContractType("TestToken"),
		Version:       semver.MustParse("1.0.0"),
	}))
	for _, p := range cfg.peerPools {
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

	loadedChains := map[uint64]cldf_chain.BlockChain{
		reverseTestLocalSel:  cldf_aptos.Chain{Selector: reverseTestLocalSel},
		reverseTestRemoteSel: cldf_aptos.Chain{Selector: reverseTestRemoteSel},
	}
	lggr := logger.Test(t)
	e := cldf.Environment{
		Logger:           lggr,
		OperationsBundle: cldf_ops.NewBundle(func() context.Context { return t.Context() }, lggr, cldf_ops.NewMemoryReporter()),
		DataStore:        ds.Seal(),
		BlockChains:      cldf_chain.NewBlockChains(loadedChains),
	}

	removed := []RemoveRemotePoolsSequenceInput{}
	wiped := []RemoveRemoteChainConfigsSequenceInput{}
	remotePools := map[string][]string{}
	for _, p := range cfg.peerPools {
		remotePools[p.address] = p.lists
	}
	local := &reverseTestAdapter{
		supportedChains:       cfg.localSupported,
		remoteToken:           "TOKEN_REMOTE",
		localToken:            "TOKEN_LOCAL",
		remotePoolsBySelector: cfg.localRemotePools,
		removed:               &removed,
	}
	var localAdapter TokenAdapter = local
	if cfg.localWithWipe {
		localAdapter = &reverseTestWipeAdapter{reverseTestAdapter: local, wiped: &wiped, unsupported: cfg.localWipeUnsupported}
	}
	remoteBase := &reverseTestAdapter{remotePools: remotePools, removed: &removed}
	var remoteAdapter TokenAdapter = remoteBase
	if cfg.remoteWithWipe {
		remoteAdapter = &reverseTestWipeAdapter{reverseTestAdapter: remoteBase, wiped: &wiped}
	}

	reg := newTokenAdapterRegistry()
	reg.RegisterTokenAdapter(localFamily, reverseTestV2_0_0, localAdapter)
	reg.RegisterTokenAdapter(remoteFamily, reverseTestRemoteVersion, remoteAdapter)
	// The local and remote selectors share one family, and TAR registrations are first-write-wins per
	// family, so a single manager serves both. It returns the peer's active pool.
	reg.RegisterTokenAdminRegistryManager(remoteFamily, &reverseTestTAR{activePool: cfg.activePool})

	planner := newRemoveRemotePoolsPlanner(e, reg)
	entries := append([]RemoveRemotePoolsPerPool{cfg.entry}, cfg.extraEntries...)
	for _, entry := range entries {
		if err := planner.planPool(entry); err != nil {
			return nil, nil, err
		}
	}
	_, _, err = planner.execute()
	return removed, wiped, err
}

// Test that wipeChainConfigs on its own removes every supported selector whose remote-pool list is
// empty and keeps the ones that still have a remote pool.
func TestRemoveRemotePoolsWipe_StandaloneSweepsEmptyLanes(t *testing.T) {
	removed, wiped, err := runPlanPoolTest(t, planPoolTestConfig{
		entry: RemoveRemotePoolsPerPool{
			ChainSelector:    reverseTestLocalSel,
			Pool:             datastore.AddressRef{Address: "A_OLD"},
			WipeChainConfigs: true,
		},
		localSupported: []uint64{reverseTestRemoteSel, reverseTestOtherSel},
		localRemotePools: map[uint64]map[string][]string{
			reverseTestRemoteSel: {"A_OLD": {}},
			reverseTestOtherSel:  {"A_OLD": {"KEEP"}},
		},
		localWithWipe: true,
	})
	require.NoError(t, err)
	require.Empty(t, removed, "a standalone wipe removes no pools directly")
	require.Len(t, wiped, 1)
	require.Equal(t, reverseTestLocalSel, wiped[0].Selector)
	require.Equal(t, []uint64{reverseTestRemoteSel}, wiped[0].RemoteChainSelectors)
}

// Test that an explicit removal which empties a selector also wipes its chain config, and that the
// selector's pools are not also queued for per-pool removal.
func TestRemoveRemotePoolsWipe_ExplicitRemovalEmptiesSelector(t *testing.T) {
	removed, wiped, err := runPlanPoolTest(t, planPoolTestConfig{
		entry: RemoveRemotePoolsPerPool{
			ChainSelector:       reverseTestLocalSel,
			Pool:                datastore.AddressRef{Address: "A_OLD"},
			WipeChainConfigs:    true,
			RemotePoolsToRemove: []RemotePoolToRemove{{Selector: reverseTestRemoteSel, Remote: datastore.AddressRef{Address: "B_OLD"}}},
		},
		localSupported: []uint64{reverseTestRemoteSel},
		localRemotePools: map[uint64]map[string][]string{
			reverseTestRemoteSel: {"A_OLD": {"B_OLD"}},
		},
		localWithWipe: true,
	})
	require.NoError(t, err)
	require.Empty(t, removed, "the wiped selector's pools are not queued for per-pool removal")
	require.Len(t, wiped, 1)
	require.Equal(t, []uint64{reverseTestRemoteSel}, wiped[0].RemoteChainSelectors)
}

// Test that removing one of several remote pools on a selector keeps its chain config and still
// queues the per-pool removal.
func TestRemoveRemotePoolsWipe_PartialRemovalKeepsChain(t *testing.T) {
	removed, wiped, err := runPlanPoolTest(t, planPoolTestConfig{
		entry: RemoveRemotePoolsPerPool{
			ChainSelector:       reverseTestLocalSel,
			Pool:                datastore.AddressRef{Address: "A_OLD"},
			WipeChainConfigs:    true,
			RemotePoolsToRemove: []RemotePoolToRemove{{Selector: reverseTestRemoteSel, Remote: datastore.AddressRef{Address: "B_OLD"}}},
		},
		localSupported: []uint64{reverseTestRemoteSel},
		localRemotePools: map[uint64]map[string][]string{
			reverseTestRemoteSel: {"A_OLD": {"B_OLD", "B_NEW"}},
		},
		localWithWipe: true,
	})
	require.NoError(t, err)
	require.Empty(t, wiped, "a selector that still has a remote pool keeps its chain config")
	require.Len(t, removed, 1)
	require.Equal(t, []RemotePoolToRemove{{Selector: reverseTestRemoteSel, Remote: datastore.AddressRef{Address: "B_OLD"}}}, removed[0].RemotePoolsToRemove)
}

// Test that a selector no longer supported is never wiped (idempotency): the wipe is gated on the
// plan-time supported-chain list, so a re-run after a wipe does not revert with NonExistentChain.
func TestRemoveRemotePoolsWipe_AlreadyGoneIsNotRewiped(t *testing.T) {
	_, wiped, err := runPlanPoolTest(t, planPoolTestConfig{
		entry: RemoveRemotePoolsPerPool{
			ChainSelector:       reverseTestLocalSel,
			Pool:                datastore.AddressRef{Address: "A_OLD"},
			WipeChainConfigs:    true,
			RemotePoolsToRemove: []RemotePoolToRemove{{Selector: reverseTestRemoteSel, Remote: datastore.AddressRef{Address: "B_OLD"}}},
		},
		localSupported: []uint64{},
		localRemotePools: map[uint64]map[string][]string{
			reverseTestRemoteSel: {"A_OLD": {}},
		},
		localWithWipe: true,
	})
	require.NoError(t, err)
	require.Empty(t, wiped)
}

// Test that wipeChainConfigs fails at planning when the adapter cannot remove chain configs.
func TestRemoveRemotePoolsWipe_UnsupportedAdapterFailsPlanning(t *testing.T) {
	_, _, err := runPlanPoolTest(t, planPoolTestConfig{
		entry: RemoveRemotePoolsPerPool{
			ChainSelector:    reverseTestLocalSel,
			Pool:             datastore.AddressRef{Address: "A_OLD"},
			WipeChainConfigs: true,
		},
		localSupported: []uint64{reverseTestRemoteSel},
		localWithWipe:  false,
	})
	require.ErrorContains(t, err, "does not support removing remote chain configs")
}

// Test that an explicitly-listed remote on a lane the local pool does not support is skipped with a
// warning rather than failing: the local-supported gate is not bypassed, and recovery after a
// one-sided wipe is done by adding the peer pool as its own entry (see remove_remote_pools.go).
func TestRemoveRemotePoolsWipe_ExplicitRemoteOnUnsupportedLaneIsSkipped(t *testing.T) {
	removed, wiped, err := runPlanPoolTest(t, planPoolTestConfig{
		entry: RemoveRemotePoolsPerPool{
			ChainSelector:       reverseTestLocalSel,
			Pool:                datastore.AddressRef{Address: "A_OLD"},
			WipeChainConfigs:    true,
			Bidirectional:       true,
			RemotePoolsToRemove: []RemotePoolToRemove{{Selector: reverseTestRemoteSel, Remote: datastore.AddressRef{Address: "B_NEW"}}},
		},
		localSupported: []uint64{}, // already wiped
		peerPools:      []reverseTestPool{{address: "B_NEW", version: reverseTestRemoteVersion, lists: []string{"A_OLD"}}},
		localWithWipe:  true,
		remoteWithWipe: true,
		activePool:     "B_NEW",
	})
	require.NoError(t, err)
	require.Empty(t, wiped, "the unsupported lane is not wiped again")
	for _, r := range removed {
		require.Equal(t, "A_OLD", r.TokenPoolRef.Address, "the peer is not touched when the local lane is unsupported")
	}
}

// Test that a peer which still lists another pool for the local chain keeps its chain config and
// gets a per-pool removal instead of a wipe.
func TestRemoveRemotePoolsWipe_PeerPartialKeepsChain(t *testing.T) {
	removed, wiped, err := runPlanPoolTest(t, planPoolTestConfig{
		entry: RemoveRemotePoolsPerPool{
			ChainSelector:       reverseTestLocalSel,
			Pool:                datastore.AddressRef{Address: "A_OLD"},
			WipeChainConfigs:    true,
			Bidirectional:       true,
			RemotePoolsToRemove: []RemotePoolToRemove{{Selector: reverseTestRemoteSel, Remote: datastore.AddressRef{Address: "B_NEW"}}},
		},
		localSupported: []uint64{reverseTestRemoteSel},
		localRemotePools: map[uint64]map[string][]string{
			reverseTestRemoteSel: {"A_OLD": {"B_NEW"}},
		},
		peerPools:      []reverseTestPool{{address: "B_NEW", version: reverseTestRemoteVersion, lists: []string{"A_OLD", "A_NEW"}}},
		localWithWipe:  true,
		remoteWithWipe: true,
		activePool:     "B_NEW",
	})
	require.NoError(t, err)
	require.Len(t, wiped, 1, "the local pool's own list emptied, so its chain config is wiped")
	require.Equal(t, "A_OLD", wiped[0].TokenPoolRef.Address)
	require.Equal(t, []uint64{reverseTestRemoteSel}, wiped[0].RemoteChainSelectors)
	require.Len(t, removed, 1)
	require.Equal(t, "B_NEW", removed[0].TokenPoolRef.Address)
}

// Test that a peer whose adapter cannot wipe is an unsupported peer: fail by default, and with
// skipUnsupportedPeers remove its pool entry while leaving its chain config.
func TestRemoveRemotePoolsWipe_UnsupportedPeer(t *testing.T) {
	cfg := planPoolTestConfig{
		entry: RemoveRemotePoolsPerPool{
			ChainSelector:       reverseTestLocalSel,
			Pool:                datastore.AddressRef{Address: "A_OLD"},
			WipeChainConfigs:    true,
			Bidirectional:       true,
			RemotePoolsToRemove: []RemotePoolToRemove{{Selector: reverseTestRemoteSel, Remote: datastore.AddressRef{Address: "B_NEW"}}},
		},
		localSupported: []uint64{reverseTestRemoteSel},
		localRemotePools: map[uint64]map[string][]string{
			reverseTestRemoteSel: {"A_OLD": {"B_NEW"}},
		},
		peerPools:      []reverseTestPool{{address: "B_NEW", version: reverseTestRemoteVersion, lists: []string{"A_OLD"}}},
		localWithWipe:  true,
		remoteWithWipe: false,
		activePool:     "B_NEW",
	}

	_, _, err := runPlanPoolTest(t, cfg)
	require.ErrorContains(t, err, "does not support removing remote chain configs")

	cfg.entry.SkipUnsupportedPeers = true
	removed, wiped, err := runPlanPoolTest(t, cfg)
	require.NoError(t, err)
	require.Len(t, wiped, 1, "the local pool's chain config is still wiped")
	require.Equal(t, "A_OLD", wiped[0].TokenPoolRef.Address)
	require.Len(t, removed, 1)
	require.Equal(t, "B_NEW", removed[0].TokenPoolRef.Address)
}

// TestRemoveRemotePoolsWipe_UnsupportedVersionFailsPlanning covers an adapter that implements the
// interface but reports the resolved version cannot wipe: the planner must fail before any write.
func TestRemoveRemotePoolsWipe_UnsupportedVersionFailsPlanning(t *testing.T) {
	_, _, err := runPlanPoolTest(t, planPoolTestConfig{
		entry: RemoveRemotePoolsPerPool{
			ChainSelector:    reverseTestLocalSel,
			Pool:             datastore.AddressRef{Address: "A_OLD"},
			WipeChainConfigs: true,
		},
		localSupported:       []uint64{reverseTestRemoteSel},
		localWithWipe:        true,
		localWipeUnsupported: true,
	})
	require.ErrorContains(t, err, "does not support removing remote chain configs")
}

// Test that an explicit remotePoolsToRemove scopes wipeChainConfigs to the lanes it names: an
// already-empty lane on the same pool is left alone, so the blast radius equals the removal's scope.
func TestRemoveRemotePoolsWipe_ExplicitScopesToNamedLane(t *testing.T) {
	removed, wiped, err := runPlanPoolTest(t, planPoolTestConfig{
		entry: RemoveRemotePoolsPerPool{
			ChainSelector:       reverseTestLocalSel,
			Pool:                datastore.AddressRef{Address: "A_OLD"},
			WipeChainConfigs:    true,
			RemotePoolsToRemove: []RemotePoolToRemove{{Selector: reverseTestRemoteSel, Remote: datastore.AddressRef{Address: "B_OLD"}}},
		},
		localSupported: []uint64{reverseTestRemoteSel, reverseTestOtherSel},
		localRemotePools: map[uint64]map[string][]string{
			reverseTestRemoteSel: {"A_OLD": {"B_OLD"}},
			reverseTestOtherSel:  {"A_OLD": {}},
		},
		localWithWipe: true,
	})
	require.NoError(t, err)
	require.Empty(t, removed)
	require.Len(t, wiped, 1)
	require.Equal(t, []uint64{reverseTestRemoteSel}, wiped[0].RemoteChainSelectors, "only the named lane is wiped, not the already-empty one")
}

// Test that two explicit entries on the same pool, naming different lanes, each wipe only their own
// lane; the wipes merge into one call.
func TestRemoveRemotePoolsWipe_TwoExplicitEntriesScopeToTheirLanes(t *testing.T) {
	removed, wiped, err := runPlanPoolTest(t, planPoolTestConfig{
		entry: RemoveRemotePoolsPerPool{
			ChainSelector:       reverseTestLocalSel,
			Pool:                datastore.AddressRef{Address: "A_OLD"},
			WipeChainConfigs:    true,
			RemotePoolsToRemove: []RemotePoolToRemove{{Selector: reverseTestRemoteSel, Remote: datastore.AddressRef{Address: "B_OLD"}}},
		},
		extraEntries: []RemoveRemotePoolsPerPool{{
			ChainSelector:       reverseTestLocalSel,
			Pool:                datastore.AddressRef{Address: "A_OLD"},
			WipeChainConfigs:    true,
			RemotePoolsToRemove: []RemotePoolToRemove{{Selector: reverseTestOtherSel, Remote: datastore.AddressRef{Address: "C_OLD"}}},
		}},
		localSupported: []uint64{reverseTestRemoteSel, reverseTestOtherSel},
		localRemotePools: map[uint64]map[string][]string{
			reverseTestRemoteSel: {"A_OLD": {"B_OLD"}},
			reverseTestOtherSel:  {"A_OLD": {"C_OLD"}},
		},
		localWithWipe: true,
	})
	require.NoError(t, err)
	require.Empty(t, removed)
	require.Len(t, wiped, 1)
	require.ElementsMatch(t, []uint64{reverseTestRemoteSel, reverseTestOtherSel}, wiped[0].RemoteChainSelectors)
}

// Test that allRemotes still considers every supported lane: the one with a remote pool is emptied by
// the run's removals and the already-empty one is swept.
func TestRemoveRemotePoolsWipe_AllRemotesWipesEveryLane(t *testing.T) {
	removed, wiped, err := runPlanPoolTest(t, planPoolTestConfig{
		entry: RemoveRemotePoolsPerPool{
			ChainSelector:    reverseTestLocalSel,
			Pool:             datastore.AddressRef{Address: "A_OLD"},
			WipeChainConfigs: true,
			AllRemotes:       true,
		},
		localSupported: []uint64{reverseTestRemoteSel, reverseTestOtherSel},
		localRemotePools: map[uint64]map[string][]string{
			reverseTestRemoteSel: {"A_OLD": {"B_OLD"}},
			reverseTestOtherSel:  {"A_OLD": {}},
		},
		localWithWipe: true,
	})
	require.NoError(t, err)
	require.Empty(t, removed)
	require.Len(t, wiped, 1)
	require.ElementsMatch(t, []uint64{reverseTestRemoteSel, reverseTestOtherSel}, wiped[0].RemoteChainSelectors)
}
