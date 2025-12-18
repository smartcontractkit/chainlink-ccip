package adapters

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences/cctp"
	seq_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var _ adapters.CCTPChain = &CCTPChainAdapter{}

// CCTPChainAdapter is the adapter for CCTP chains.
type CCTPChainAdapter struct{}

// AddressRefToBytes returns the byte representation of an address for this chain family.
func (c *CCTPChainAdapter) AddressRefToBytes(ref datastore.AddressRef) ([]byte, error) {
	return common.HexToAddress(ref.Address).Bytes(), nil
}

// DeployCCTPChain returns the sequence for deploying a CCTP chain.
func (c *CCTPChainAdapter) DeployCCTPChain() *operations.Sequence[adapters.DeployCCTPInput[string, []byte], seq_core.OnChainOutput, chain.BlockChains] {
	return cctp.DeployCCTPChain
}
