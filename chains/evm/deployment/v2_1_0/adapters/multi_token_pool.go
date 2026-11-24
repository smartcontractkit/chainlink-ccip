package adapters

import (
	"fmt"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	sequence_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var _ tokens.MultiTokenPoolAdapter = (*MultiTokenPoolAdapter)(nil)

// MultiTokenPoolAdapter implements EVM multi-token pool deployment.
type MultiTokenPoolAdapter struct{}

// ValidateDeployMultiTokenPool validates the input and the existing state. It sends no transactions.
func (a *MultiTokenPoolAdapter) ValidateDeployMultiTokenPool(e cldf.Environment, input tokens.DeployMultiTokenPoolInput) error {
	if err := sequences.ValidateDeployMultiTokenPoolInput(input); err != nil {
		return err
	}
	chain, ok := e.BlockChains.EVMChains()[input.ChainSelector]
	if !ok {
		return fmt.Errorf("EVM chain with selector %d not found", input.ChainSelector)
	}
	return sequences.ValidateDeployMultiTokenPoolState(e.OperationsBundle, chain, e.DataStore, input)
}

// DeployMultiTokenPool returns the EVM multi-token pool deployment sequence.
func (a *MultiTokenPoolAdapter) DeployMultiTokenPool() *cldf_ops.Sequence[tokens.DeployMultiTokenPoolInput, sequence_utils.OnChainOutput, cldf_chain.BlockChains] {
	return sequences.DeployMultiTokenPool
}
