package adapters

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/token_pool"
	evm_tokens "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences/tokens"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// TokenAdapter is the adapter for EVM tokens using 1.7.0 token pools.
type TokenAdapter struct{}

// ConfigureTokenForTransfersSequence returns the sequence for configuring an EVM token with a 1.7.0 token pool.
func (t *TokenAdapter) ConfigureTokenForTransfersSequence() *operations.Sequence[tokens.ConfigureTokenForTransfersInput, sequences.OnChainOutput, chain.BlockChains] {
	return evm_tokens.ConfigureTokenForTransfers
}

// AddressRefToBytes returns an EVM address reference as an EVM address padded to 32 bytes.
func (t *TokenAdapter) AddressRefToBytes(ref datastore.AddressRef) ([]byte, error) {
	return common.LeftPadBytes(common.HexToAddress(ref.Address).Bytes(), 32), nil
}

// DeriveTokenAddress derives the token address from a token pool reference, returning it as an EVM address padded to 32 bytes.
func (t *TokenAdapter) DeriveTokenAddress(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef) ([]byte, error) {
	chain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return nil, fmt.Errorf("chain with selector %d not found", chainSelector)
	}
	getTokenReport, err := cldf_ops.ExecuteOperation(e.OperationsBundle, token_pool.GetToken, chain, contract.FunctionInput[any]{
		ChainSelector: chainSelector,
		Address:       common.HexToAddress(poolRef.Address),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get token address from token pool with address %s on %s: %w", poolRef.Address, chain, err)
	}

	return t.AddressRefToBytes(datastore.AddressRef{
		Address: getTokenReport.Output.Hex(),
	})
}
