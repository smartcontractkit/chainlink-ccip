package changesets

import (
	"errors"
	"fmt"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

// WithChainSelector is implemented by configs that specify a target chain selector.
type WithChainSelector interface {
	ChainSelector() uint64
}

// ResolveEVMChainDep resolves an evm.Chain dependency from the environment based on the chain selector provided by the config.
func ResolveEVMChainDep[CFG WithChainSelector](e deployment.Environment, cfg CFG) (evm.Chain, error) {
	evmChains := e.BlockChains.EVMChains()
	if evmChains == nil {
		return evm.Chain{}, errors.New("no EVM chains found in environment")
	}
	if _, exists := evmChains[cfg.ChainSelector()]; !exists {
		return evm.Chain{}, fmt.Errorf("no EVM chain with selector %d found in environment", cfg.ChainSelector())
	}
	return evmChains[cfg.ChainSelector()], nil
}
