package sequences

import (
	"errors"
	"fmt"

	sequencescommon "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

// ResolveEVMChainDep resolves an evm.Chain dependency from the environment based on the chain selector provided by the config.
func ResolveEVMChainDep[CFG sequencescommon.WithChainSelector](e deployment.Environment, cfg CFG) (evm.Chain, error) {
	evmChains := e.BlockChains.EVMChains()
	if evmChains == nil {
		return evm.Chain{}, errors.New("no EVM chains found in environment")
	}
	if _, exists := evmChains[cfg.ChainSelector()]; !exists {
		return evm.Chain{}, fmt.Errorf("no EVM chain with selector %d found in environment", cfg.ChainSelector())
	}
	return evmChains[cfg.ChainSelector()], nil
}
