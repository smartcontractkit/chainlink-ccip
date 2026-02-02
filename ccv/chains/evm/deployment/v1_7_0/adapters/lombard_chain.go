package adapters

import (
	"github.com/ethereum/go-ethereum/common"

	seq_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences/lombard"
)

var _ adapters.LombardChain = &LombardChainAdapter{}

// LombardChainAdapter is the adapter for Lombard chains.
type LombardChainAdapter struct{}

// AddressRefToBytes returns the byte representation of an address for this chain family.
func (c *LombardChainAdapter) AddressRefToBytes(ref datastore.AddressRef) ([]byte, error) {
	return common.HexToAddress(ref.Address).Bytes(), nil
}

// DeployLombardChain returns the sequence for deploying a Lombard chain.
func (c *LombardChainAdapter) DeployLombardChain() *operations.Sequence[adapters.DeployLombardInput, seq_core.OnChainOutput, adapters.DeployLombardChainDeps] {
	return lombard.DeployLombardChain
}
