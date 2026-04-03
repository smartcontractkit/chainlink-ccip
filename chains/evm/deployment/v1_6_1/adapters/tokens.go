package adapters

import (
	v1_6_0_seq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	evm_seq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var _ tokens.TokenAdapter = &TokenAdapter{}

// TokenAdapter is the adapter for EVM tokens using 1.6.1 token pools.
// It embeds the v1.6.0 EVMAdapter (which itself embeds EVMTokenBase) and
// overrides only the methods that differ for 1.6.1 pools.
type TokenAdapter struct {
	v1_6_0_seq.EVMAdapter
}

// ConfigureTokenForTransfersSequence returns the v1.6.1-specific sequence.
func (t *TokenAdapter) ConfigureTokenForTransfersSequence() *cldf_ops.Sequence[tokens.ConfigureTokenForTransfersInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return evm_seq.ConfigureTokenForTransfers
}
