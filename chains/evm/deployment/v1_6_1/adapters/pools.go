package adapters

import (
	evmseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
)

var _ tokens.TokenAdapterRegistry = (*PoolsAdapter)(nil)

type PoolsAdapter struct {
	evm *evmseq.EVMAdapter
}
