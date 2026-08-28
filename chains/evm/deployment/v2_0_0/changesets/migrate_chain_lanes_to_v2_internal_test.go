package changesets

import (
	"errors"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_erc20_with_drip"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	changesetscore "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
	v2changesets "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/offchain"
)

// fakeLaneVersionResolver is a stand-in for a family's on-chain LaneVersionResolver.
type fakeLaneVersionResolver struct {
	supported map[uint64]bool
	lanes     map[uint64]map[uint64]*semver.Version
	err       error
	// errs returns a per-chain derivation error; checked before err.
	errs map[uint64]error
}

func (f *fakeLaneVersionResolver) IsSupportedChain(_ cldf.Environment, sel uint64) bool {
	return f.supported[sel]
}

func (f *fakeLaneVersionResolver) DeriveLaneVersionsForChain(_ cldf.Environment, sel uint64) (map[uint64]*semver.Version, []*semver.Version, error) {
	if perChainErr, ok := f.errs[sel]; ok {
		return nil, nil, perChainErr
	}
	if f.err != nil {
		return nil, nil, f.err
	}
	return f.lanes[sel], nil, nil
}

// fakeConfigImporter is a stand-in for a family/version ConfigImporter, reporting a canned set of
// supported tokens per remote chain.
type fakeConfigImporter struct {
	tokensPerRemote map[uint64][]common.Address
	initErr         error
	tokensErr       error
}

func (f *fakeConfigImporter) InitializeAdapter(_ cldf.Environment, _ uint64) error {
	return f.initErr
}

func (f *fakeConfigImporter) ConnectedChains(_ cldf.Environment, _ uint64) ([]uint64, error) {
	return nil, nil
}

func (f *fakeConfigImporter) SupportedTokensPerRemoteChain(_ cldf.Environment, _ uint64, remotes []uint64) (map[uint64][]common.Address, error) {
	if f.tokensErr != nil {
		return nil, f.tokensErr
	}
	out := make(map[uint64][]common.Address, len(remotes))
	for _, remote := range remotes {
		if toks, ok := f.tokensPerRemote[remote]; ok {
			out[remote] = toks
		}
	}
	return out, nil
}

func (f *fakeConfigImporter) SequenceImportConfig() *cldf_ops.Sequence[deploy.ImportConfigPerChainInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}

func registryWithResolver(t *testing.T, r *fakeLaneVersionResolver) *adapters.DeployChainContractsRegistry {
	t.Helper()
	reg := adapters.NewDeployChainContractsRegistry()
	reg.RegisterLaneVersionResolver(chainsel.FamilyEVM, r)
	return reg
}

func newFQRegistry() *deploy.FQAndRampUpdaterRegistry {
	return &deploy.FQAndRampUpdaterRegistry{
		FeeQuoterUpdater:            make(map[string]deploy.FeeQuoterUpdater[any]),
		RampUpdater:                 make(map[string]deploy.RampUpdater),
		ConfigImporter:              make(map[string]deploy.ConfigImporter),
		ImportconfigVersionResolver: make(map[string]deploy.LaneVersionResolver),
	}
}

// fqRegistryWithImporter registers a specific importer for a single version.
func fqRegistryWithImporter(t *testing.T, version *semver.Version, importer deploy.ConfigImporter) *deploy.FQAndRampUpdaterRegistry {
	t.Helper()
	reg := newFQRegistry()
	reg.RegisterConfigImporter(chainsel.FamilyEVM, version, importer)
	return reg
}

// supportedFQRegistry registers empty importers for the migratable source versions (1.5.0, 1.6.0),
// which is what the discovery version-gate checks for.
func supportedFQRegistry(t *testing.T) *deploy.FQAndRampUpdaterRegistry {
	t.Helper()
	reg := newFQRegistry()
	reg.RegisterConfigImporter(chainsel.FamilyEVM, semver.MustParse("1.5.0"), &fakeConfigImporter{})
	reg.RegisterConfigImporter(chainsel.FamilyEVM, semver.MustParse("1.6.0"), &fakeConfigImporter{})
	return reg
}

func TestDiscoverLanesToMigrate_SkipsAlreadyV2Lanes(t *testing.T) {
	chainA := chainsel.TEST_90000001.Selector
	remoteOld := chainsel.TEST_90000002.Selector
	remoteV2 := chainsel.TEST_90000003.Selector

	resolver := &fakeLaneVersionResolver{
		supported: map[uint64]bool{chainA: true},
		lanes: map[uint64]map[uint64]*semver.Version{
			chainA: {
				remoteOld: semver.MustParse("1.6.0"),
				remoteV2:  semver.MustParse("2.0.0"),
			},
		},
	}

	lanes, err := discoverLanesToMigrate(cldf.Environment{}, registryWithResolver(t, resolver), supportedFQRegistry(t), nil, MigrateChainLanesToV2Config{
		MigrateChainLanesToV2Input: MigrateChainLanesToV2Input{ChainSelectors: []uint64{chainA}},
	})
	require.NoError(t, err)
	require.Len(t, lanes, 1)
	assert.Equal(t, chainA, lanes[0].ChainA)
	assert.Equal(t, remoteOld, lanes[0].ChainB)
}

func TestDiscoverLanesToMigrate_SkipsUnsupportedVersionLane(t *testing.T) {
	chainA := chainsel.TEST_90000001.Selector
	remoteSupported := chainsel.TEST_90000002.Selector
	remoteUnsupported := chainsel.TEST_90000003.Selector

	resolver := &fakeLaneVersionResolver{
		supported: map[uint64]bool{chainA: true},
		lanes: map[uint64]map[uint64]*semver.Version{
			chainA: {
				remoteSupported:   semver.MustParse("1.6.0"),
				remoteUnsupported: semver.MustParse("1.2.0"), // no config importer registered
			},
		},
	}

	lanes, err := discoverLanesToMigrate(cldf.Environment{}, registryWithResolver(t, resolver), supportedFQRegistry(t), nil, MigrateChainLanesToV2Config{
		MigrateChainLanesToV2Input: MigrateChainLanesToV2Input{ChainSelectors: []uint64{chainA}},
	})
	require.NoError(t, err)
	require.Len(t, lanes, 1)
	assert.Equal(t, remoteSupported, lanes[0].ChainB)
}

func TestDiscoverLanesToMigrate_ExcludesBlocklistedRemotes(t *testing.T) {
	chainA := chainsel.TEST_90000001.Selector
	remoteKeep := chainsel.TEST_90000002.Selector
	remoteDrop := chainsel.TEST_90000003.Selector

	resolver := &fakeLaneVersionResolver{
		supported: map[uint64]bool{chainA: true},
		lanes: map[uint64]map[uint64]*semver.Version{
			chainA: {
				remoteKeep: semver.MustParse("1.6.0"),
				remoteDrop: semver.MustParse("1.6.0"),
			},
		},
	}

	lanes, err := discoverLanesToMigrate(cldf.Environment{}, registryWithResolver(t, resolver), supportedFQRegistry(t), nil, MigrateChainLanesToV2Config{
		MigrateChainLanesToV2Input: MigrateChainLanesToV2Input{
			ChainSelectors:       []uint64{chainA},
			ExcludedRemoteChains: []uint64{remoteDrop},
		},
	})
	require.NoError(t, err)
	require.Len(t, lanes, 1)
	assert.Equal(t, remoteKeep, lanes[0].ChainB)
}

func TestDiscoverLanesToMigrate_IncludesOnlyAllowlistedRemotes(t *testing.T) {
	chainA := chainsel.TEST_90000001.Selector
	remoteDrop := chainsel.TEST_90000002.Selector
	remoteKeep := chainsel.TEST_90000003.Selector

	resolver := &fakeLaneVersionResolver{
		supported: map[uint64]bool{chainA: true},
		lanes: map[uint64]map[uint64]*semver.Version{
			chainA: {
				remoteDrop: semver.MustParse("1.6.0"),
				remoteKeep: semver.MustParse("1.6.0"),
			},
		},
	}

	lanes, err := discoverLanesToMigrate(cldf.Environment{}, registryWithResolver(t, resolver), supportedFQRegistry(t), nil, MigrateChainLanesToV2Config{
		MigrateChainLanesToV2Input: MigrateChainLanesToV2Input{
			ChainSelectors:       []uint64{chainA},
			IncludedRemoteChains: []uint64{remoteKeep},
		},
	})
	require.NoError(t, err)
	require.Len(t, lanes, 1)
	assert.Equal(t, remoteKeep, lanes[0].ChainB)
}

func TestDiscoverLanesToMigrate_DedupsAcrossMultipleChains(t *testing.T) {
	chainA := chainsel.TEST_90000001.Selector
	chainB := chainsel.TEST_90000002.Selector
	chainC := chainsel.TEST_90000003.Selector
	v := semver.MustParse("1.6.0")

	// Fully-meshed lanes: every chain reports the other two as remotes, so each undirected lane is
	// discovered from both endpoints and must collapse to a single entry.
	resolver := &fakeLaneVersionResolver{
		supported: map[uint64]bool{chainA: true, chainB: true, chainC: true},
		lanes: map[uint64]map[uint64]*semver.Version{
			chainA: {chainB: v, chainC: v},
			chainB: {chainA: v, chainC: v},
			chainC: {chainA: v, chainB: v},
		},
	}

	lanes, err := discoverLanesToMigrate(cldf.Environment{}, registryWithResolver(t, resolver), supportedFQRegistry(t), nil, MigrateChainLanesToV2Config{
		MigrateChainLanesToV2Input: MigrateChainLanesToV2Input{ChainSelectors: []uint64{chainA, chainB, chainC}},
	})
	require.NoError(t, err)

	keys := make(map[[2]uint64]struct{})
	for _, lane := range lanes {
		keys[canonicalLaneKey(lane.ChainA, lane.ChainB)] = struct{}{}
	}
	require.Len(t, lanes, 3, "each undirected lane must appear exactly once")
	assert.Len(t, keys, 3, "lanes must not overlap")
	assert.Contains(t, keys, canonicalLaneKey(chainA, chainB))
	assert.Contains(t, keys, canonicalLaneKey(chainA, chainC))
	assert.Contains(t, keys, canonicalLaneKey(chainB, chainC))
}

func TestMigrateChainLanesToV2_ValidateRejectsExcludedChainSelector(t *testing.T) {
	chainA := chainsel.TEST_90000001.Selector

	cs := MigrateChainLanesToV2(
		adapters.NewCommitteeVerifierContractRegistry(),
		adapters.NewChainFamilyRegistry(),
		changesetscore.GetRegistry(),
		adapters.NewDeployChainContractsRegistry(),
		newFQRegistry(),
	)

	err := cs.VerifyPreconditions(cldf.Environment{}, MigrateChainLanesToV2Config{
		Topology: &offchain.EnvironmentTopology{},
		MigrateChainLanesToV2Input: MigrateChainLanesToV2Input{
			ChainSelectors:       []uint64{chainA},
			ExcludedRemoteChains: []uint64{chainA},
		},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "cannot be in both chainSelectors and excludedRemoteChains")
}

func TestMigrateChainLanesToV2_ValidateRejectsIncludedExcludedRemoteChain(t *testing.T) {
	chainA := chainsel.TEST_90000001.Selector
	remote := chainsel.TEST_90000002.Selector

	cs := MigrateChainLanesToV2(
		adapters.NewCommitteeVerifierContractRegistry(),
		adapters.NewChainFamilyRegistry(),
		changesetscore.GetRegistry(),
		adapters.NewDeployChainContractsRegistry(),
		newFQRegistry(),
	)

	err := cs.VerifyPreconditions(cldf.Environment{}, MigrateChainLanesToV2Config{
		Topology: &offchain.EnvironmentTopology{},
		MigrateChainLanesToV2Input: MigrateChainLanesToV2Input{
			ChainSelectors:       []uint64{chainA},
			IncludedRemoteChains: []uint64{remote},
			ExcludedRemoteChains: []uint64{remote},
		},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "cannot be in both includedRemoteChains and excludedRemoteChains")
}

func TestDiscoverLanesToMigrate_DedupsBidirectionalLane(t *testing.T) {
	chainA := chainsel.TEST_90000001.Selector
	chainB := chainsel.TEST_90000002.Selector

	resolver := &fakeLaneVersionResolver{
		supported: map[uint64]bool{chainA: true, chainB: true},
		lanes: map[uint64]map[uint64]*semver.Version{
			chainA: {chainB: semver.MustParse("1.6.0")},
			chainB: {chainA: semver.MustParse("1.6.0")},
		},
	}

	lanes, err := discoverLanesToMigrate(cldf.Environment{}, registryWithResolver(t, resolver), supportedFQRegistry(t), nil, MigrateChainLanesToV2Config{
		MigrateChainLanesToV2Input: MigrateChainLanesToV2Input{ChainSelectors: []uint64{chainA, chainB}},
	})
	require.NoError(t, err)
	require.Len(t, lanes, 1)
	assert.Equal(t, canonicalLaneKey(chainA, chainB), canonicalLaneKey(lanes[0].ChainA, lanes[0].ChainB))
}

func TestDiscoverLanesToMigrate_SkipsUnknownVersionLane(t *testing.T) {
	chainA := chainsel.TEST_90000001.Selector
	remote := chainsel.TEST_90000002.Selector

	resolver := &fakeLaneVersionResolver{
		supported: map[uint64]bool{chainA: true},
		lanes: map[uint64]map[uint64]*semver.Version{
			chainA: {remote: nil}, // unknown version -> no resolver -> not migrated
		},
	}

	lanes, err := discoverLanesToMigrate(cldf.Environment{}, registryWithResolver(t, resolver), supportedFQRegistry(t), nil, MigrateChainLanesToV2Config{
		MigrateChainLanesToV2Input: MigrateChainLanesToV2Input{ChainSelectors: []uint64{chainA}},
	})
	require.NoError(t, err)
	require.Empty(t, lanes)
}

func TestDiscoverLanesToMigrate_SkipsAllWithoutFQRegistry(t *testing.T) {
	chainA := chainsel.TEST_90000001.Selector
	remote := chainsel.TEST_90000002.Selector

	resolver := &fakeLaneVersionResolver{
		supported: map[uint64]bool{chainA: true},
		lanes: map[uint64]map[uint64]*semver.Version{
			chainA: {remote: semver.MustParse("1.6.0")},
		},
	}

	// Without a fee-quoter/ramp updater registry we cannot confirm a version is supported, so nothing
	// is migrated. (The hard requirement is enforced by the changeset's VerifyPreconditions.)
	lanes, err := discoverLanesToMigrate(cldf.Environment{}, registryWithResolver(t, resolver), nil, nil, MigrateChainLanesToV2Config{
		MigrateChainLanesToV2Input: MigrateChainLanesToV2Input{ChainSelectors: []uint64{chainA}},
	})
	require.NoError(t, err)
	require.Empty(t, lanes)
}

func TestDiscoverLanesToMigrate_ErrorsWhenResolverMissing(t *testing.T) {
	chainA := chainsel.TEST_90000001.Selector

	// Empty registry: no resolver registered for any family.
	_, err := discoverLanesToMigrate(cldf.Environment{}, adapters.NewDeployChainContractsRegistry(), supportedFQRegistry(t), nil, MigrateChainLanesToV2Config{
		MigrateChainLanesToV2Input: MigrateChainLanesToV2Input{ChainSelectors: []uint64{chainA}},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "no lane version resolver registered")
}

func TestDiscoverLanesToMigrate_PropagatesResolverError(t *testing.T) {
	chainA := chainsel.TEST_90000001.Selector

	resolver := &fakeLaneVersionResolver{
		supported: map[uint64]bool{chainA: true},
		err:       errors.New("rpc unavailable"),
	}

	_, err := discoverLanesToMigrate(cldf.Environment{}, registryWithResolver(t, resolver), supportedFQRegistry(t), nil, MigrateChainLanesToV2Config{
		MigrateChainLanesToV2Input: MigrateChainLanesToV2Input{ChainSelectors: []uint64{chainA}},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "rpc unavailable")
}

func TestDiscoverLanesToMigrate_ErrorsWhenChainUnsupported(t *testing.T) {
	chainA := chainsel.TEST_90000001.Selector

	resolver := &fakeLaneVersionResolver{
		supported: map[uint64]bool{chainA: false},
	}

	_, err := discoverLanesToMigrate(cldf.Environment{}, registryWithResolver(t, resolver), supportedFQRegistry(t), nil, MigrateChainLanesToV2Config{
		MigrateChainLanesToV2Input: MigrateChainLanesToV2Input{ChainSelectors: []uint64{chainA}},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not supported")
}

func TestDiscoverLanesToMigrate_SkipsDeprecatedRemotes(t *testing.T) {
	chainA := chainsel.TEST_90000001.Selector
	remoteLive := chainsel.TEST_90000002.Selector
	remoteDeprecated := chainsel.CORN_MAINNET.Selector

	require.True(t, deprecatedChain(remoteDeprecated))
	require.False(t, deprecatedChain(remoteLive))

	resolver := &fakeLaneVersionResolver{
		supported: map[uint64]bool{chainA: true},
		lanes: map[uint64]map[uint64]*semver.Version{
			chainA: {
				remoteLive:       semver.MustParse("1.6.0"),
				remoteDeprecated: semver.MustParse("1.6.0"),
			},
		},
	}

	lanes, err := discoverLanesToMigrate(
		cldf.Environment{},
		registryWithResolver(t, resolver),
		supportedFQRegistry(t),
		nil,
		MigrateChainLanesToV2Config{
			MigrateChainLanesToV2Input: MigrateChainLanesToV2Input{ChainSelectors: []uint64{chainA}},
		},
	)
	require.NoError(t, err)
	require.Len(t, lanes, 1)
	assert.Equal(t, remoteLive, lanes[0].ChainB)
	for _, lane := range lanes {
		assert.NotEqual(t, remoteDeprecated, lane.ChainB, "deprecated remote must not be migrated")
	}
}

func TestDiscoverLanesToMigrate_SkipsDeprecatedLocalChain(t *testing.T) {
	deprecated := chainsel.CORN_MAINNET.Selector

	lanes, err := discoverLanesToMigrate(
		cldf.Environment{},
		adapters.NewDeployChainContractsRegistry(),
		supportedFQRegistry(t),
		nil,
		MigrateChainLanesToV2Config{
			MigrateChainLanesToV2Input: MigrateChainLanesToV2Input{ChainSelectors: []uint64{deprecated}},
		},
	)
	require.NoError(t, err)
	require.Empty(t, lanes)
}

func TestDiscoverLanesToMigrate_ExcludesNonEVMRemotes(t *testing.T) {
	chainA := chainsel.TEST_90000001.Selector
	remoteEVM := chainsel.TEST_90000002.Selector
	remoteSolana := chainsel.SOLANA_DEVNET.Selector

	// Guard the assumptions the test relies on: chainA and remoteEVM are EVM, remoteSolana is not.
	// Both remotes are otherwise identical migration candidates (same, supported lane version), so
	// the only thing that can drop the Solana lane is the non-EVM filter.
	require.True(t, isEVMChain(chainA))
	require.True(t, isEVMChain(remoteEVM))
	require.False(t, isEVMChain(remoteSolana))

	resolver := &fakeLaneVersionResolver{
		supported: map[uint64]bool{chainA: true},
		lanes: map[uint64]map[uint64]*semver.Version{
			chainA: {
				remoteEVM:    semver.MustParse("1.6.0"),
				remoteSolana: semver.MustParse("1.6.0"),
			},
		},
	}

	lanes, err := discoverLanesToMigrate(cldf.Environment{}, registryWithResolver(t, resolver), supportedFQRegistry(t), nil, MigrateChainLanesToV2Config{
		MigrateChainLanesToV2Input: MigrateChainLanesToV2Input{ChainSelectors: []uint64{chainA}},
	})
	require.NoError(t, err)
	require.Len(t, lanes, 1)
	assert.Equal(t, remoteEVM, lanes[0].ChainB)
	for _, lane := range lanes {
		assert.NotEqual(t, remoteSolana, lane.ChainB, "non-EVM remote must not be migrated")
	}
}

func TestDiscoverLanesToMigrate_SkipsNonEVMLocalChain(t *testing.T) {
	solana := chainsel.SOLANA_DEVNET.Selector

	// No resolver needed: non-EVM local chains are skipped before any resolver lookup.
	lanes, err := discoverLanesToMigrate(cldf.Environment{}, adapters.NewDeployChainContractsRegistry(), supportedFQRegistry(t), nil, MigrateChainLanesToV2Config{
		MigrateChainLanesToV2Input: MigrateChainLanesToV2Input{ChainSelectors: []uint64{solana}},
	})
	require.NoError(t, err)
	require.Empty(t, lanes)
}

// symbolLookupFrom builds a fake tokenSymbolLookup from a token-address -> symbol map.
func symbolLookupFrom(symbols map[common.Address]string) tokenSymbolLookup {
	return func(_ uint64, token common.Address) (string, error) {
		return symbols[token], nil
	}
}

func TestDiscoverLanesToMigrate_SkipsLanesWithExcludedTokenSymbol(t *testing.T) {
	chainA := chainsel.TEST_90000001.Selector
	remoteWithUSDC := chainsel.TEST_90000002.Selector
	remotePlain := chainsel.TEST_90000003.Selector

	version := semver.MustParse("1.6.0")
	usdc := common.HexToAddress("0x000000000000000000000000000000000000abcd")
	other := common.HexToAddress("0x0000000000000000000000000000000000001234")

	resolver := &fakeLaneVersionResolver{
		supported: map[uint64]bool{chainA: true, remotePlain: true},
		lanes: map[uint64]map[uint64]*semver.Version{
			chainA: {
				remoteWithUSDC: version,
				remotePlain:    version,
			},
			remotePlain: {chainA: version},
		},
	}
	importer := &fakeConfigImporter{
		tokensPerRemote: map[uint64][]common.Address{
			remoteWithUSDC: {other, usdc},
			remotePlain:    {other},
		},
	}
	symbolOf := symbolLookupFrom(map[common.Address]string{
		usdc:  "USDC",
		other: "WETH",
	})

	lanes, err := discoverLanesToMigrate(
		cldf.Environment{},
		registryWithResolver(t, resolver),
		fqRegistryWithImporter(t, version, importer),
		symbolOf,
		MigrateChainLanesToV2Config{
			MigrateChainLanesToV2Input: MigrateChainLanesToV2Input{
				ChainSelectors: []uint64{chainA},
				// Lowercase to confirm case-insensitive matching against on-chain "USDC".
				ExcludeLanesWithTokenSymbols: []string{"usdc"},
			},
		},
	)
	require.NoError(t, err)
	require.Len(t, lanes, 1)
	assert.Equal(t, remotePlain, lanes[0].ChainB)
}

func TestDiscoverLanesToMigrate_SkipsLaneWithExcludedTokenOnlyInReverseDirection(t *testing.T) {
	chainA := chainsel.TEST_90000001.Selector
	chainB := chainsel.TEST_90000002.Selector
	version := semver.MustParse("1.6.0")
	usdc := common.HexToAddress("0x000000000000000000000000000000000000abcd")
	other := common.HexToAddress("0x0000000000000000000000000000000000001234")

	resolver := &fakeLaneVersionResolver{
		supported: map[uint64]bool{chainA: true, chainB: true},
		lanes: map[uint64]map[uint64]*semver.Version{
			chainA: {chainB: version},
			chainB: {chainA: version},
		},
	}
	importer := &fakeConfigImporter{
		tokensPerRemote: map[uint64][]common.Address{
			chainB: {other}, // chainA -> chainB has no excluded token.
			chainA: {usdc},  // chainB -> chainA carries USDC.
		},
	}

	lanes, err := discoverLanesToMigrate(
		cldf.Environment{},
		registryWithResolver(t, resolver),
		fqRegistryWithImporter(t, version, importer),
		symbolLookupFrom(map[common.Address]string{usdc: "USDC", other: "WETH"}),
		MigrateChainLanesToV2Config{
			MigrateChainLanesToV2Input: MigrateChainLanesToV2Input{
				ChainSelectors:               []uint64{chainA},
				ExcludeLanesWithTokenSymbols: []string{"USDC"},
			},
		},
	)
	require.NoError(t, err)
	require.Empty(t, lanes)
}

func TestDiscoverLanesToMigrate_SkipsReverseCheckWhenResolverErrors(t *testing.T) {
	chainA := chainsel.TEST_90000001.Selector
	chainB := chainsel.TEST_90000002.Selector
	version := semver.MustParse("1.6.0")

	// chainA -> chainB is a valid candidate, but resolving chainB's own lanes (reverse direction)
	// fails — e.g. chainB is deprecated with dead RPCs. Discovery must skip, not abort.
	resolver := &fakeLaneVersionResolver{
		supported: map[uint64]bool{chainA: true, chainB: true},
		lanes: map[uint64]map[uint64]*semver.Version{
			chainA: {chainB: version},
		},
		errs: map[uint64]error{
			chainB: errors.New("EVM chain with selector 9043146809313071210 not found in environment"),
		},
	}

	lanes, err := discoverLanesToMigrate(
		cldf.Environment{},
		registryWithResolver(t, resolver),
		fqRegistryWithImporter(t, version, &fakeConfigImporter{tokensPerRemote: map[uint64][]common.Address{chainB: {}}}),
		symbolLookupFrom(map[common.Address]string{}),
		MigrateChainLanesToV2Config{
			MigrateChainLanesToV2Input: MigrateChainLanesToV2Input{
				ChainSelectors:               []uint64{chainA},
				ExcludeLanesWithTokenSymbols: []string{"USDC"},
			},
		},
	)
	require.NoError(t, err)
	require.Empty(t, lanes)
}

func TestBuildTestTokenExpansionInput_DeploysTestTokenPerEVMChain(t *testing.T) {
	chainA := chainsel.TEST_90000001.Selector
	chainB := chainsel.TEST_90000002.Selector
	chainSolana := chainsel.SOLANA_DEVNET.Selector

	in, err := buildTestTokenExpansionInput(cldf.Environment{}, MigrateChainLanesToV2Config{}, []v2changesets.CrossFamilyLanePair{
		{ChainA: chainA, ChainB: chainB},
		{ChainA: chainA, ChainB: chainSolana}, // non-EVM lane: no test token wiring
	})
	require.NoError(t, err)
	require.Len(t, in.TokenExpansionInputPerChain, 2, "only EVM chains participating in EVM lanes get a test token")

	for _, sel := range []uint64{chainA, chainB} {
		perChain := in.TokenExpansionInputPerChain[sel]
		require.NotNil(t, perChain.DeployTokenInput)
		assert.Equal(t, testTokenSymbol, perChain.DeployTokenInput.Symbol)
		assert.Equal(t, testTokenSymbol, perChain.DeployTokenInput.Name)
		assert.Equal(t, uint8(testTokenDecimals), perChain.DeployTokenInput.Decimals)
		assert.Equal(t, cldf.ContractType(burn_mint_erc20_with_drip.ContractType), perChain.DeployTokenInput.Type)

		require.NotNil(t, perChain.DeployTokenPoolInput)
		assert.Equal(t, testTokenSymbol, perChain.DeployTokenPoolInput.TokenPoolQualifier)
		require.NotNil(t, perChain.DeployTokenPoolInput.RouterRef)
		assert.Equal(t, datastore.ContractType(router.TestRouterContractType), perChain.DeployTokenPoolInput.RouterRef.Type)

		require.NotNil(t, perChain.TokenTransferConfig)
		assert.Equal(t, testTokenSymbol, perChain.TokenTransferConfig.TokenRef.Qualifier)
		assert.True(t, perChain.SkipOwnershipTransfer)
	}
	assert.Contains(t, in.TokenExpansionInputPerChain[chainA].TokenTransferConfig.RemoteChains, chainB)
	assert.Contains(t, in.TokenExpansionInputPerChain[chainB].TokenTransferConfig.RemoteChains, chainA)
	assert.NotContains(t, in.TokenExpansionInputPerChain[chainA].TokenTransferConfig.RemoteChains, chainSolana)
}

func TestBuildTestTokenExpansionInput_ReusesExistingToken(t *testing.T) {
	chainA := chainsel.TEST_90000001.Selector
	chainB := chainsel.TEST_90000002.Selector

	// chainA already has the TESTTR token in the datastore; chainB does not.
	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainA,
		Type:          datastore.ContractType(burn_mint_erc20_with_drip.ContractType),
		Version:       burn_mint_erc20_with_drip.Version,
		Qualifier:     testTokenSymbol,
		Address:       "0x000000000000000000000000000000000000babe",
	}))

	in, err := buildTestTokenExpansionInput(cldf.Environment{DataStore: ds.Seal()}, MigrateChainLanesToV2Config{}, []v2changesets.CrossFamilyLanePair{
		{ChainA: chainA, ChainB: chainB},
	})
	require.NoError(t, err)
	assert.Nil(t, in.TokenExpansionInputPerChain[chainA].DeployTokenInput, "existing TESTTR token must be reused, not redeployed")
	assert.NotNil(t, in.TokenExpansionInputPerChain[chainB].DeployTokenInput, "chain without TESTTR must deploy a fresh token")
}

func TestDiscoverLanesToMigrate_SkipsVersionWithoutConfigImporter(t *testing.T) {
	chainA := chainsel.TEST_90000001.Selector
	remote := chainsel.TEST_90000002.Selector

	resolver := &fakeLaneVersionResolver{
		supported: map[uint64]bool{chainA: true},
		lanes: map[uint64]map[uint64]*semver.Version{
			chainA: {remote: semver.MustParse("1.6.0")},
		},
	}
	// Registry only has an importer for 1.5.0, but the lane is 1.6.0 -> unsupported -> skipped.
	lanes, err := discoverLanesToMigrate(
		cldf.Environment{},
		registryWithResolver(t, resolver),
		fqRegistryWithImporter(t, semver.MustParse("1.5.0"), &fakeConfigImporter{}),
		symbolLookupFrom(nil),
		MigrateChainLanesToV2Config{
			MigrateChainLanesToV2Input: MigrateChainLanesToV2Input{ChainSelectors: []uint64{chainA}},
		},
	)
	require.NoError(t, err)
	require.Empty(t, lanes)
}
