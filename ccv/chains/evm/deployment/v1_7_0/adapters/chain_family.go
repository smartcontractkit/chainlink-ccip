package adapters

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	seq_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// ChainFamilyAdapter is the adapter for chains of the EVM family.
type ChainFamilyAdapter struct{}

// ConfigureChainForLanesSequence returns the sequence for configuring a chain of the EVM family for CCIP lanes.
func (c *ChainFamilyAdapter) ConfigureChainForLanes() *operations.Sequence[adapters.ConfigureChainForLanesInput, seq_core.OnChainOutput, chain.BlockChains] {
	return sequences.ConfigureChainForLanes
}

// AddressRefToBytes returns the byte representation of an address for this chain family.
func (c *ChainFamilyAdapter) AddressRefToBytes(ref datastore.AddressRef) ([]byte, error) {
	return common.HexToAddress(ref.Address).Bytes(), nil
}
