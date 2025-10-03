package adapters

import (
	"github.com/ethereum/go-ethereum/common"
	evm_tokens "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type TokenAdapter struct{}

func (*TokenAdapter) ConfigureTokenForTransfersSequence() *operations.Sequence[tokens.ConfigureTokenForTransfersInput, sequences.OnChainOutput, chain.BlockChains] {
	return evm_tokens.ConfigureTokenForTransfers
}

func (*TokenAdapter) ConvertRefToBytes(ref datastore.AddressRef) ([]byte, error) {
	return common.FromHex(ref.Address), nil
}
