package changesets

import (
	"errors"
	"fmt"
	"maps"
	"slices"
	"strconv"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/offchain"
)

var (
	// ErrChainMissingFromExecutorPools means a lane chain has no chain_configs entry in any
	// executor pool, so no executor is configured to serve the lane.
	ErrChainMissingFromExecutorPools = errors.New("chain missing from all executor pool chain_configs")
	// ErrNoChainFamilyAdapter means no adapter is registered for a lane chain's family.
	ErrNoChainFamilyAdapter = errors.New("no chain family adapter registered")
	// ErrLaneAddressUnresolved means a contract the lane needs is not in the datastore.
	ErrLaneAddressUnresolved = errors.New("lane contract address unresolved")
	// ErrNOPMissingSigner means a committee NOP has no signer address for a chain family,
	// in the topology or in JD.
	ErrNOPMissingSigner = errors.New("NOP has no signer address for chain family")
)

// laneContract pairs a contract name with the adapter call that resolves its address,
// so failures can name the contract that is missing.
type laneContract struct {
	name    string
	resolve func(datastore.DataStore, uint64) ([]byte, error)
}

// validateLaneAddressesResolvable resolves every contract address the apply path needs,
// without writing anything.
//
// Apply resolves these addresses one chain at a time, so a chain missing its contracts
// only surfaces after the JD round trip and after earlier chains have been processed.
// Resolving them up front turns that into a precondition failure naming the chain and
// the contract.
func validateLaneAddressesResolvable(
	e deployment.Environment,
	chainFamilyRegistry *adapters.ChainFamilyRegistry,
	chains []partialChainConfig,
	useTestRouter bool,
) error {
	for _, chainCfg := range chains {
		local := chainCfg.ChainSelector
		localAdapter, err := adapterForChain(chainFamilyRegistry, local)
		if err != nil {
			return err
		}

		router := laneContract{name: "router", resolve: localAdapter.GetRouterAddress}
		if useTestRouter {
			router = laneContract{name: "testRouter", resolve: localAdapter.GetTestRouter}
		}
		localContracts := []laneContract{
			router,
			{name: "onRamp", resolve: localAdapter.GetOnRampAddress},
			{name: "feeQuoter", resolve: localAdapter.GetFQAddress},
			{name: "offRamp", resolve: localAdapter.GetOffRampAddress},
		}
		for _, contract := range localContracts {
			if _, err := contract.resolve(e.DataStore, local); err != nil {
				return fmt.Errorf("chain %d %s: %w: %w", local, contract.name, ErrLaneAddressUnresolved, err)
			}
		}

		// Sorted so a topology with several remotes reports the same chain first every run.
		for _, remote := range slices.Sorted(maps.Keys(chainCfg.RemoteChains)) {
			remoteAdapter, err := adapterForChain(chainFamilyRegistry, remote)
			if err != nil {
				return fmt.Errorf("chain %d remote chain %d: %w", local, remote, err)
			}
			remoteContracts := []laneContract{
				{name: "onRamp", resolve: remoteAdapter.GetOnRampAddress},
				{name: "offRamp", resolve: remoteAdapter.GetOffRampAddress},
			}
			for _, contract := range remoteContracts {
				if _, err := contract.resolve(e.DataStore, remote); err != nil {
					return fmt.Errorf("chain %d remote chain %d %s: %w: %w",
						local, remote, contract.name, ErrLaneAddressUnresolved, err)
				}
			}

			// Executor resolution is a separate concern from the contract addresses above.
			qualifier := defaultQualifier
			if q := chainCfg.RemoteChains[remote].DefaultExecutorQualifier; q != nil {
				qualifier = *q
			}
			if _, err := localAdapter.ResolveExecutor(e.DataStore, local, qualifier); err != nil {
				return fmt.Errorf("chain %d executor (qualifier %q): %w: %w",
					local, qualifier, ErrLaneAddressUnresolved, err)
			}
		}
	}
	return nil
}

// validateLaneSignersResolvable checks that every lane leg can build its committee
// verifier signature quorum from the topology, falling back to JD-registered keys.
//
// Apply does this during enrichment, after the JD round trip, and fails partway through
// the chain loop. Doing it here reports the NOP and lane leg as a precondition failure.
func validateLaneSignersResolvable(
	e deployment.Environment,
	topology *offchain.EnvironmentTopology,
	chains []partialChainConfig,
) error {
	committeeNOPs := filterNOPsToCommitteeMembers(topology.NOPTopology, chains)
	signingKeysByNOP, err := fetchSigningKeysForNOPsByFamilies(e, committeeNOPs, deriveFamiliesFromChains(chains))
	if err != nil {
		return fmt.Errorf("failed to fetch signing keys: %w", err)
	}

	for _, chainCfg := range chains {
		for _, verifier := range chainCfg.CommitteeVerifiers {
			for _, remote := range slices.Sorted(maps.Keys(verifier.RemoteChains)) {
				_, err := getSignatureConfigForLane(
					e, topology, verifier.CommitteeQualifier, chainCfg.ChainSelector, remote, signingKeysByNOP,
				)
				if err != nil {
					return fmt.Errorf("lane %d -> %d: %w", chainCfg.ChainSelector, remote, err)
				}
			}
		}
	}
	return nil
}

// adapterForChain returns the registered chain family adapter for a chain selector.
func adapterForChain(registry *adapters.ChainFamilyRegistry, chainSelector uint64) (adapters.ChainFamily, error) {
	family, err := chainsel.GetSelectorFamily(chainSelector)
	if err != nil {
		return nil, fmt.Errorf("chain %d: %w", chainSelector, err)
	}
	adapter, ok := registry.GetChainFamily(family)
	if !ok {
		return nil, fmt.Errorf("chain %d family %q: %w", chainSelector, family, ErrNoChainFamilyAdapter)
	}
	return adapter, nil
}

// validateExecutorPoolCoverage checks that every given chain appears in at least one
// executor pool's chain_configs.
//
// Executor pools drive the off-chain executor jobs, so a lane chain absent from every
// pool has no executor serving it. Nothing else on the lane path reads executor pools.
//
// A topology that declares no pools at all is left alone: there is nothing to check
// against, and environments that do not model executor pools are still valid.
func validateExecutorPoolCoverage(topology *offchain.EnvironmentTopology, chainSelectors []uint64) error {
	if topology == nil || len(topology.ExecutorPools) == 0 {
		return nil
	}

	covered := make(map[string]struct{})
	for _, pool := range topology.ExecutorPools {
		for chainKey := range pool.ChainConfigs {
			covered[chainKey] = struct{}{}
		}
	}

	missing := make([]uint64, 0, len(chainSelectors))
	for _, selector := range chainSelectors {
		if _, ok := covered[strconv.FormatUint(selector, 10)]; !ok {
			missing = append(missing, selector)
		}
	}
	if len(missing) > 0 {
		slices.Sort(missing)
		return fmt.Errorf("chains %v: %w", missing, ErrChainMissingFromExecutorPools)
	}
	return nil
}
