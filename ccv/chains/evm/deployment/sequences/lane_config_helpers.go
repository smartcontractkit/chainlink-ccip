package sequences

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/executor"
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
