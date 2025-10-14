package v160

import (
	ccipapi "github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// high level API
func (a *SolanaAdapter) ConfigureLaneLegAsSource() *cldf_ops.Sequence[ccipapi.UpdateLanesInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil // Not implemented for Solana
}

func (a *SolanaAdapter) ConfigureLaneLegAsDest() *cldf_ops.Sequence[ccipapi.UpdateLanesInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil // Not implemented for Solana
}
