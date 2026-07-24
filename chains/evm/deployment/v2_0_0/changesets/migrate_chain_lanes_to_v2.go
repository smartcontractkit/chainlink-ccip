package changesets

import (
	"context"
	"fmt"
	"slices"
	"strings"
	"sync"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/erc20"
	"golang.org/x/sync/errgroup"

	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	changesetscore "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
	v2changesets "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/offchain"
)

// v2MajorVersion is the major version of a lane already on CCIP 2.0 (no migration needed).
const v2MajorVersion = 2

// tokenSymbolLookup resolves the ERC20 symbol of a token deployed on the given chain.
type tokenSymbolLookup func(chainSel uint64, token common.Address) (string, error)

// MigrateChainLanesToV2Input is the durable-pipeline payload for migrate_chain_lanes_to_v2.
//
// It intentionally omits the topology: the topology is injected by the config resolver (which
// loads it from the environment) into MigrateChainLanesToV2Config before the changeset runs. This
// mirrors how v2changesets.BuildLanesCrossFamilyConfig embeds into
// v2changesets.ConfigureChainsForLanesFromTopologyConfig.
//
// Unlike build_lanes_cross_family (which takes an explicit list of lane pairs), this changeset
// discovers the lanes to migrate on-chain: for each chain in ChainSelectors it uses the registered
// per-family LaneVersionResolver to read the router's connected remote chains and their current
// lane version, then migrates every EVM lane with a supported source version (one that has a
// registered config importer, e.g. 1.5.0/1.6.0) that is not already on CCIP 2.0 by delegating to
// ConfigureChainsForLanesFromTopology.
//
// The resolved lane list is bidirectional, so a lane discovered from both endpoints (when both
// chains appear in ChainSelectors) is configured exactly once.
type MigrateChainLanesToV2Input struct {
	// ChainSelectors are the chains whose existing lanes should be migrated to CCIP 2.0.
	ChainSelectors []uint64 `json:"chainSelectors" yaml:"chainSelectors"`
	// ExcludedRemoteChains is an optional blocklist of remote chain selectors to skip during
	// discovery. Any lane to one of these remotes is left untouched. Useful for routing around
	// a flaky/unreachable remote chain RPC or to intentionally hold a lane back from migration.
	ExcludedRemoteChains []uint64 `json:"excludedRemoteChains,omitempty" yaml:"excludedRemoteChains,omitempty"`
	// ExcludeLanesWithTokenSymbols optionally skips any lane whose local chain has a token with one
	// of these symbols configured (supported) for the remote chain — e.g. []string{"USDC", "LBTC"}
	// to hold back token lanes that need a dedicated migration path. Symbols are matched
	// case-insensitively against each configured token's on-chain ERC20 symbol(). Reading which
	// tokens are configured requires a non-nil fee-quoter/ramp updater registry.
	ExcludeLanesWithTokenSymbols []string   `json:"excludeLanesWithTokenSymbols,omitempty" yaml:"excludeLanesWithTokenSymbols,omitempty"`
	MCMS                         mcms.Input `json:"mcms" yaml:"mcms"`
	TestRouter                   *bool      `json:"testRouter,omitempty" yaml:"testRouter,omitempty"`
}

// MigrateChainLanesToV2Config is the full changeset config: the durable-pipeline payload plus the
// topology that the config resolver injects.
type MigrateChainLanesToV2Config struct {
	Topology *offchain.EnvironmentTopology `json:"topology" yaml:"topology"`
	MigrateChainLanesToV2Input
}

// MigrateChainLanesToV2 is the canonical changeset for migrating all pre-2.0 EVM lanes on a set of
// chains to CCIP 2.0. It resolves the lane pairs from on-chain state and then reuses the
// enrichment / resolution / dispatch path of ConfigureChainsForLanesFromTopology.
//
// deployChainContractsRegistry is the source of the per-family LaneVersionResolver used to
// enumerate connected remote chains. fqRegistry supplies the per-family, per-version ConfigImporter;
// it determines which lane source versions can be migrated (only versions with a registered importer,
// e.g. 1.5.0 and 1.6.0) and is also used to read which tokens are configured on a lane. The
// remaining two registries are forwarded verbatim to ConfigureChainsForLanesFromTopology. They are
// injected rather than looked up from globals so the changeset is testable with mocks.
func MigrateChainLanesToV2(
	committeeVerifierContractRegistry *adapters.CommitteeVerifierContractRegistry,
	chainFamilyRegistry *adapters.ChainFamilyRegistry,
	mcmsRegistry *changesetscore.MCMSReaderRegistry,
	deployChainContractsRegistry *adapters.DeployChainContractsRegistry,
	fqRegistry *deploy.FQAndRampUpdaterRegistry,
) deployment.ChangeSetV2[MigrateChainLanesToV2Config] {
	underlying := v2changesets.ConfigureChainsForLanesFromTopology(committeeVerifierContractRegistry, chainFamilyRegistry, mcmsRegistry)

	validate := func(e deployment.Environment, cfg MigrateChainLanesToV2Config) error {
		switch {
		case committeeVerifierContractRegistry == nil:
			return fmt.Errorf("committee verifier contract registry is required")
		case chainFamilyRegistry == nil:
			return fmt.Errorf("chain family registry is required")
		case mcmsRegistry == nil:
			return fmt.Errorf("mcms registry is required")
		case deployChainContractsRegistry == nil:
			return fmt.Errorf("deploy chain contracts registry is required")
		case cfg.Topology == nil:
			return fmt.Errorf("topology is required")
		case len(cfg.ChainSelectors) == 0:
			return fmt.Errorf("at least one chain selector is required")
		case fqRegistry == nil:
			return fmt.Errorf("fee-quoter/ramp updater registry is required")
		}

		// Cheap, non-RPC checks only: on-chain lane discovery happens in Apply. A chain being
		// migrated must not also be excluded, and each EVM chain must have a lane version resolver
		// so we fail fast on missing wiring. Non-EVM chains (e.g. Solana) are not migrated by this
		// changeset and are skipped here.
		excludedRemotes := newUint64Set(cfg.ExcludedRemoteChains)
		for _, chainSel := range cfg.ChainSelectors {
			if _, excluded := excludedRemotes[chainSel]; excluded {
				return fmt.Errorf("chain %d cannot be in both chainSelectors and excludedRemoteChains", chainSel)
			}
			if !isEVMChain(chainSel) {
				continue
			}
			if _, ok := deployChainContractsRegistry.GetLaneVersionResolver(chainSel); !ok {
				return fmt.Errorf("no lane version resolver registered for chain %d", chainSel)
			}
		}
		return nil
	}

	apply := func(e deployment.Environment, cfg MigrateChainLanesToV2Config) (deployment.ChangesetOutput, error) {
		lanes, err := discoverLanesToMigrate(e, deployChainContractsRegistry, fqRegistry, makeEVMSymbolLookup(e), cfg)
		if err != nil {
			return deployment.ChangesetOutput{}, err
		}
		if len(lanes) == 0 {
			e.Logger.Infow("migrate_chain_lanes_to_v2: no pre-2.0 lanes to migrate", "chainSelectors", cfg.ChainSelectors)
			return deployment.ChangesetOutput{}, nil
		}

		laneDescriptions := make([]string, len(lanes))
		for i, lane := range lanes {
			laneDescriptions[i] = fmt.Sprintf("%d<->%d", lane.ChainA, lane.ChainB)
		}
		e.Logger.Infow("migrate_chain_lanes_to_v2: migrating lanes to CCIP 2.0",
			"count", len(lanes), "lanes", laneDescriptions)

		resolvedCfg := v2changesets.ConfigureChainsForLanesFromTopologyConfig{
			Topology: cfg.Topology,
			BuildLanesCrossFamilyConfig: v2changesets.BuildLanesCrossFamilyConfig{
				Lanes:      lanes,
				MCMS:       cfg.MCMS,
				TestRouter: cfg.TestRouter,
				// A migration's whole purpose is to swap each lane's pre-2.0 OnRamp for the CCIP 2.0
				// OnRamp on the (prod) router, so it must be allowed to overwrite the existing mapping.
				AllowOnrampOverride: true,
			},
		}
		if err := underlying.VerifyPreconditions(e, resolvedCfg); err != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("resolved lane config failed preconditions: %w", err)
		}
		return underlying.Apply(e, resolvedCfg)
	}

	return deployment.CreateChangeSet(apply, validate)
}

// laneDiscoverer resolves, from live on-chain state, the set of lanes to migrate to CCIP 2.0. It
// holds the shared dependencies and per-run filters so the discovery steps read as small methods
// rather than long parameter lists.
type laneDiscoverer struct {
	env             deployment.Environment
	resolvers       *adapters.DeployChainContractsRegistry
	fqRegistry      *deploy.FQAndRampUpdaterRegistry
	symbolOf        tokenSymbolLookup
	excludedRemotes map[uint64]struct{}
	excludedSymbols map[string]struct{}
}

// discoverLanesToMigrate returns a deduplicated, deterministically ordered set of bidirectional
// lane pairs to migrate: connected, EVM-only, not already on CCIP 2.0, not blocklisted, and (when
// ExcludeLanesWithTokenSymbols is set) not carrying a token with an excluded symbol.
func discoverLanesToMigrate(
	e deployment.Environment,
	resolvers *adapters.DeployChainContractsRegistry,
	fqRegistry *deploy.FQAndRampUpdaterRegistry,
	symbolOf tokenSymbolLookup,
	cfg MigrateChainLanesToV2Config,
) ([]v2changesets.CrossFamilyLanePair, error) {
	d := &laneDiscoverer{
		env:             e,
		resolvers:       resolvers,
		fqRegistry:      fqRegistry,
		symbolOf:        symbolOf,
		excludedRemotes: newUint64Set(cfg.ExcludedRemoteChains),
		excludedSymbols: canonicalSymbolSet(cfg.ExcludeLanesWithTokenSymbols),
	}
	return d.run(cfg.ChainSelectors)
}

// run discovers candidate remotes per chain concurrently (the RPC-bound work) — one goroutine per
// chain, each against its own chain's RPC — then merges them into a deduplicated set of
// bidirectional lane pairs sequentially in input order. The dedup (keyed by canonicalLaneKey)
// guarantees no overlapping lanes are produced when multiple chains that share a lane are given.
func (d *laneDiscoverer) run(chainSelectors []uint64) ([]v2changesets.CrossFamilyLanePair, error) {
	candidates := make([][]uint64, len(chainSelectors))
	var g errgroup.Group
	for i, chainSel := range chainSelectors {
		g.Go(func() error {
			remotes, err := d.chainCandidates(chainSel)
			if err != nil {
				return err
			}
			candidates[i] = remotes
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}

	seen := make(map[[2]uint64]struct{})
	var pairs []v2changesets.CrossFamilyLanePair
	for i, chainSel := range chainSelectors {
		for _, remote := range candidates[i] {
			key := canonicalLaneKey(chainSel, remote)
			if _, dup := seen[key]; dup {
				continue
			}
			seen[key] = struct{}{}
			pairs = append(pairs, v2changesets.CrossFamilyLanePair{ChainA: chainSel, ChainB: remote})
		}
	}
	return pairs, nil
}

// chainCandidates returns the sorted remotes on chainSel whose lane should be migrated: connected,
// EVM, not already on CCIP 2.0, not blocklisted, and not carrying an excluded token symbol. Token
// detection runs only over the surviving candidates, so we never read token config for lanes that
// are already 2.0 or blocklisted.
func (d *laneDiscoverer) chainCandidates(chainSel uint64) ([]uint64, error) {
	// This changeset only migrates EVM lanes; skip non-EVM local chains (e.g. Solana) entirely.
	if !isEVMChain(chainSel) {
		return nil, nil
	}

	resolver, ok := d.resolvers.GetLaneVersionResolver(chainSel)
	if !ok {
		return nil, fmt.Errorf("no lane version resolver registered for chain %d", chainSel)
	}
	if !resolver.IsSupportedChain(d.env, chainSel) {
		return nil, fmt.Errorf("chain %d is not supported by its lane version resolver", chainSel)
	}
	laneVersions, _, err := resolver.DeriveLaneVersionsForChain(d.env, chainSel)
	if err != nil {
		return nil, fmt.Errorf("failed to derive lane versions for chain %d: %w", chainSel, err)
	}

	// Keep only EVM remotes that are connected, not blocklisted, not already on 2.0, and whose
	// source lane version we can migrate (i.e. has a registered config importer). Lanes with an
	// unknown or unsupported version are skipped — we never migrate a lane without a resolver.
	candidateVersions := make(map[uint64]*semver.Version, len(laneVersions))
	for remote, version := range laneVersions {
		switch {
		case isExcluded(d.excludedRemotes, remote):
			continue
		case !isEVMChain(remote):
			continue
		case version != nil && version.Major() >= v2MajorVersion:
			continue
		case !d.canMigrateVersion(chainSel, version):
			continue
		}
		candidateVersions[remote] = version
	}

	tokenExcluded, err := d.tokenExcludedRemotes(chainSel, candidateVersions)
	if err != nil {
		return nil, err
	}

	remotes := make([]uint64, 0, len(candidateVersions))
	for remote := range candidateVersions {
		if !isExcluded(tokenExcluded, remote) {
			remotes = append(remotes, remote)
		}
	}
	slices.Sort(remotes)
	return remotes, nil
}

// tokenExcludedRemotes returns the remotes whose lane on chainSel carries a token with an excluded
// symbol. Token support is read on-chain via the chain-family ConfigImporter (one read per lane
// version), then each token's symbol is resolved via symbolOf. Returns nil when no symbols are
// excluded.
func (d *laneDiscoverer) tokenExcludedRemotes(chainSel uint64, laneVersions map[uint64]*semver.Version) (map[uint64]struct{}, error) {
	if len(d.excludedSymbols) == 0 {
		return nil, nil
	}
	if d.fqRegistry == nil {
		return nil, fmt.Errorf("fee-quoter/ramp updater registry is required to evaluate token symbol exclusion for chain %d", chainSel)
	}
	if d.symbolOf == nil {
		return nil, fmt.Errorf("token symbol lookup is required to evaluate token symbol exclusion for chain %d", chainSel)
	}

	// One ConfigImporter read per lane version (not per remote) yields the token addresses
	// configured for each remote.
	supportedTokens := make(map[uint64][]common.Address)
	for _, grp := range groupRemotesByVersion(laneVersions) {
		importer, ok := d.fqRegistry.GetConfigImporter(chainSel, grp.version)
		if !ok {
			return nil, fmt.Errorf("no config importer registered for chain %d lane version %s; cannot evaluate token symbol exclusion", chainSel, grp.version.String())
		}
		if err := importer.InitializeAdapter(d.env, chainSel); err != nil {
			return nil, fmt.Errorf("failed to initialize config importer for chain %d version %s: %w", chainSel, grp.version.String(), err)
		}
		tokensPerRemote, err := importer.SupportedTokensPerRemoteChain(d.env, chainSel, grp.remotes)
		if err != nil {
			return nil, fmt.Errorf("failed to read supported tokens for chain %d version %s: %w", chainSel, grp.version.String(), err)
		}
		for remote, tokens := range tokensPerRemote {
			supportedTokens[remote] = append(supportedTokens[remote], tokens...)
		}
	}

	// A remote is excluded if any of its configured tokens has an excluded symbol. symbolOf caches,
	// so tokens shared across remotes are read once.
	excluded := make(map[uint64]struct{})
	for remote, tokens := range supportedTokens {
		for _, token := range tokens {
			symbol, err := d.symbolOf(chainSel, token)
			if err != nil {
				return nil, fmt.Errorf("failed to read symbol for token %s on chain %d: %w", token.Hex(), chainSel, err)
			}
			if _, bad := d.excludedSymbols[strings.ToUpper(symbol)]; bad {
				excluded[remote] = struct{}{}
				break
			}
		}
	}
	return excluded, nil
}

// canMigrateVersion reports whether a lane's source version can be migrated: we require a registered
// config importer for that version (e.g. 1.5.0, 1.6.0). An unknown (nil) version is never migrated.
func (d *laneDiscoverer) canMigrateVersion(chainSel uint64, version *semver.Version) bool {
	if version == nil || d.fqRegistry == nil {
		return false
	}
	_, ok := d.fqRegistry.GetConfigImporter(chainSel, version)
	return ok
}

// versionGroup is the set of remotes sharing a single lane version.
type versionGroup struct {
	version *semver.Version
	remotes []uint64
}

// groupRemotesByVersion groups remotes by their lane version, skipping remotes with an unknown
// (nil) version (no importer can be resolved for them; they are skipped defensively).
func groupRemotesByVersion(laneVersions map[uint64]*semver.Version) []versionGroup {
	byVersion := make(map[string]*versionGroup)
	for remote, version := range laneVersions {
		if version == nil {
			continue
		}
		grp, ok := byVersion[version.String()]
		if !ok {
			grp = &versionGroup{version: version}
			byVersion[version.String()] = grp
		}
		grp.remotes = append(grp.remotes, remote)
	}
	groups := make([]versionGroup, 0, len(byVersion))
	for _, grp := range byVersion {
		groups = append(groups, *grp)
	}
	return groups
}

// symbolCacheKey identifies a token on a specific chain for symbol-lookup caching.
type symbolCacheKey struct {
	chainSel uint64
	token    common.Address
}

// makeEVMSymbolLookup returns a tokenSymbolLookup that reads a token's ERC20 symbol() on-chain via
// the environment's EVM chain client. Non-EVM chains return an error, since token symbol exclusion
// is only supported for EVM local chains. Results are cached, and the returned closure is safe for
// concurrent use (chains are discovered in parallel).
func makeEVMSymbolLookup(e deployment.Environment) tokenSymbolLookup {
	var mu sync.Mutex
	cache := make(map[symbolCacheKey]string)
	return func(chainSel uint64, token common.Address) (string, error) {
		key := symbolCacheKey{chainSel: chainSel, token: token}

		mu.Lock()
		symbol, ok := cache[key]
		mu.Unlock()
		if ok {
			return symbol, nil
		}

		chain, ok := e.BlockChains.EVMChains()[chainSel]
		if !ok {
			return "", fmt.Errorf("token symbol exclusion is only supported for EVM chains; chain %d is not an EVM chain", chainSel)
		}
		symbol, err := readERC20Symbol(e.GetContext(), chain.Client, token)
		if err != nil {
			return "", err
		}

		mu.Lock()
		cache[key] = symbol
		mu.Unlock()
		return symbol, nil
	}
}

// readERC20Symbol reads the symbol() view function of an ERC20 token at the given address.
func readERC20Symbol(ctx context.Context, backend bind.ContractBackend, token common.Address) (string, error) {
	erc20C, err := erc20.NewERC20(token, backend)
	if err != nil {
		return "", fmt.Errorf("failed to bind ERC20 at %s: %w", token.Hex(), err)
	}
	return erc20C.Symbol(&bind.CallOpts{Context: ctx})
}

// newUint64Set returns a set of the given selectors, or nil when empty.
func newUint64Set(vals []uint64) map[uint64]struct{} {
	if len(vals) == 0 {
		return nil
	}
	set := make(map[uint64]struct{}, len(vals))
	for _, v := range vals {
		set[v] = struct{}{}
	}
	return set
}

// canonicalSymbolSet upper-cases token symbols into a set for case-insensitive comparison, or nil
// when empty.
func canonicalSymbolSet(symbols []string) map[string]struct{} {
	if len(symbols) == 0 {
		return nil
	}
	set := make(map[string]struct{}, len(symbols))
	for _, s := range symbols {
		set[strings.ToUpper(s)] = struct{}{}
	}
	return set
}

// isExcluded reports whether sel is present in the set (safe on a nil set).
func isExcluded(set map[uint64]struct{}, sel uint64) bool {
	_, ok := set[sel]
	return ok
}

// isEVMChain reports whether the chain selector belongs to the EVM family. Unknown selectors are
// treated as non-EVM.
func isEVMChain(sel uint64) bool {
	family, err := chainsel.GetSelectorFamily(sel)
	return err == nil && family == chainsel.FamilyEVM
}

// canonicalLaneKey returns an order-independent key for a bidirectional lane so that discovering
// the same lane from both endpoints collapses to a single entry.
func canonicalLaneKey(a, b uint64) [2]uint64 {
	if a <= b {
		return [2]uint64{a, b}
	}
	return [2]uint64{b, a}
}
