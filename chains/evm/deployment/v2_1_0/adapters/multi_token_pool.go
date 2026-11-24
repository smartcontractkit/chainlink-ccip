package adapters

import (
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	sequence_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var _ tokens.MultiTokenPoolAdapter = (*MultiTokenPoolAdapter)(nil)

// MultiTokenPoolAdapter implements EVM multi-token pool deployment.
type MultiTokenPoolAdapter struct{}

// DeployMultiTokenPool returns the EVM multi-token pool deployment sequence.
func (a *MultiTokenPoolAdapter) DeployMultiTokenPool() *cldf_ops.Sequence[tokens.DeployMultiTokenPoolInput, sequence_utils.OnChainOutput, cldf_chain.BlockChains] {
	return sequences.DeployMultiTokenPool
}
