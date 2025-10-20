package utils

import (
	"errors"
	"fmt"

	sequencescommon "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

// ResolveSolanaChainDep resolves a solana.Chain dependency from the environment based on the chain selector provided by the config.
func ResolveSolanaChainDep[CFG sequencescommon.WithChainSelector](e deployment.Environment, cfg CFG) (solana.Chain, error) {
	solanaChains := e.BlockChains.SolanaChains()
	if solanaChains == nil {
		return solana.Chain{}, errors.New("no Solana chains found in environment")
	}
	if _, exists := solanaChains[cfg.ChainSelector()]; !exists {
		return solana.Chain{}, fmt.Errorf("no Solana chain with selector %d found in environment", cfg.ChainSelector())
	}
	return solanaChains[cfg.ChainSelector()], nil
}
