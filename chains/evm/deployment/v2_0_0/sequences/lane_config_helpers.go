package sequences

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/executor"
)

// FilterOffRampAdds reads all currently registered OffRamps from the Router in a single call,
// then removes entries that are already present. This avoids doing no-op transactions.
func FilterOffRampAdds(
	b cldf_ops.Bundle,
	chain evm.Chain,
	routerAddr common.Address,
	offRampAdds []router.OffRamp,
) ([]router.OffRamp, error) {
	currentReport, err := cldf_ops.ExecuteOperation(b, router.GetOffRamps, chain, contract.FunctionInput[any]{
		ChainSelector: chain.Selector,
		Address:       routerAddr,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get off ramps from Router(%s) on chain %v: %w", routerAddr, chain, err)
	}
	currentSet := make(map[router.OffRamp]struct{}, len(currentReport.Output))
	for _, current := range currentReport.Output {
		currentSet[current] = struct{}{}
	}
	filtered := offRampAdds[:0]
	for _, add := range offRampAdds {
		if _, exists := currentSet[add]; !exists {
			filtered = append(filtered, add)
		}
	}
	return filtered, nil
}

// FilterExecutorDestChains reads each Executor's current dest chain list and removes entries
// whose on-chain config already matches the desired state. This is done per-executor (not
// per-remote-chain) because the Executor exposes a getDestChains bulk getter.
func FilterExecutorDestChains(
	b cldf_ops.Bundle,
	chain evm.Chain,
	destChainSelectorsPerExecutor map[common.Address][]ExecutorRemoteChainConfigArgs,
) (map[common.Address][]ExecutorRemoteChainConfigArgs, error) {
	out := make(map[common.Address][]ExecutorRemoteChainConfigArgs, len(destChainSelectorsPerExecutor))
	for executorAddr, toAdd := range destChainSelectorsPerExecutor {
		currentReport, err := cldf_ops.ExecuteOperation(b, executor.GetDestChains, chain, contract.FunctionInput[struct{}]{
			ChainSelector: chain.Selector,
			Address:       executorAddr,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get dest chains from Executor(%s) on chain %v: %w", executorAddr, chain, err)
		}
		currentMap := make(map[uint64]executor.RemoteChainConfigArgs, len(currentReport.Output))
		for _, current := range currentReport.Output {
			currentMap[current.DestChainSelector] = current
		}
		filtered := toAdd[:0]
		for _, add := range toAdd {
			cur, ok := currentMap[add.DestChainSelector]
			if ok && cur.Config.UsdCentsFee == add.Config.USDCentsFee && cur.Config.Enabled == add.Config.Enabled {
				continue
			}
			filtered = append(filtered, add)
		}
		out[executorAddr] = filtered
	}
	return out, nil
}

// extractCommitteeVerifierAddresses picks the CommitteeVerifier and its resolver out of a
// set of refs, requiring exactly one of each.
func extractCommitteeVerifierAddresses(refs []datastore.AddressRef, chainSelector uint64) (verifier string, resolver string, err error) {
	for _, addr := range refs {
		switch addr.Type {
		case datastore.ContractType(committee_verifier.ContractType):
			if verifier != "" {
				return "", "", fmt.Errorf("duplicate committee verifier contract on chain %d", chainSelector)
			}
			verifier = addr.Address
		case datastore.ContractType(CommitteeVerifierResolverType):
			if resolver != "" {
				return "", "", fmt.Errorf("duplicate committee verifier resolver contract on chain %d", chainSelector)
			}
			resolver = addr.Address
		}
	}
	if verifier == "" {
		return "", "", fmt.Errorf("committee verifier contract not found on chain %d", chainSelector)
	}
	if resolver == "" {
		return "", "", fmt.Errorf("committee verifier resolver contract not found on chain %d", chainSelector)
	}
	return verifier, resolver, nil
}

// makeAllowlistUpdates returns the adds and removes to apply so the allowlist becomes (current ∪ added) \ removed.
// It takes the current on-chain allowlist and the config's added/removed sender lists (hex addresses).
func makeAllowlistUpdates(current []common.Address, added, removed []string) (toAdd, toRemove []common.Address, err error) {
	curSet := make(map[common.Address]struct{}, len(current))
	for _, a := range current {
		curSet[a] = struct{}{}
	}
	addedSet := make(map[common.Address]struct{}, len(added))
	for _, s := range added {
		if !common.IsHexAddress(s) {
			return nil, nil, fmt.Errorf("invalid hex address in added allowlist: %q", s)
		}
		addedSet[common.HexToAddress(s)] = struct{}{}
	}
	removedSet := make(map[common.Address]struct{}, len(removed))
	for _, s := range removed {
		if !common.IsHexAddress(s) {
			return nil, nil, fmt.Errorf("invalid hex address in removed allowlist: %q", s)
		}
		removedSet[common.HexToAddress(s)] = struct{}{}
	}
	desiredSet := make(map[common.Address]struct{})
	for a := range curSet {
		if _, remove := removedSet[a]; !remove {
			desiredSet[a] = struct{}{}
		}
	}
	for a := range addedSet {
		desiredSet[a] = struct{}{}
	}
	desired := make([]common.Address, 0, len(desiredSet))
	for a := range desiredSet {
		desired = append(desired, a)
	}
	toAdd = AddressesNotIn(desired, current)
	toRemove = AddressesNotIn(current, desired)
	return toAdd, toRemove, nil
}
