package adapters

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/token_pool"
	evm_seq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// TokenAdapter is the adapter for EVM tokens using 1.6.1 token pools.
type TokenAdapter struct{}

// ConfigureTokenForTransfersSequence returns the sequence for configuring an EVM token with a 1.6.1 token pool.
func (t *TokenAdapter) ConfigureTokenForTransfersSequence() *operations.Sequence[tokens.ConfigureTokenForTransfersInput, sequences.OnChainOutput, chain.BlockChains] {
	return evm_seq.ConfigureTokenForTransfers
}

// AddressRefToBytes returns an EVM address reference as an EVM address.
func (t *TokenAdapter) AddressRefToBytes(ref datastore.AddressRef) ([]byte, error) {
	return common.HexToAddress(ref.Address).Bytes(), nil
}

// DeriveTokenAddress derives the token address from a token pool reference, returning it as an EVM address.
func (t *TokenAdapter) DeriveTokenAddress(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef) ([]byte, error) {
	chain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return nil, fmt.Errorf("chain with selector %d not found", chainSelector)
	}
	getTokenReport, err := cldf_ops.ExecuteOperation(e.OperationsBundle, token_pool.GetToken, chain, contract.FunctionInput[struct{}]{
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

func (t *TokenAdapter) ManualRegistration() *cldf_ops.Sequence[tokens.ManualRegistrationInput, sequences.OnChainOutput, chain.BlockChains] {
	// TODO implement me
	return nil
}

func (t *TokenAdapter) DeployToken() *cldf_ops.Sequence[tokens.DeployTokenInput, sequences.OnChainOutput, chain.BlockChains] {
	// TODO implement me
	return nil
}

func (t *TokenAdapter) DeployTokenVerify(e deployment.Environment, in any) error {
	// TODO implement me
	return nil
}

func (t *TokenAdapter) DeployTokenPoolForToken() *cldf_ops.Sequence[tokens.DeployTokenPoolInput, sequences.OnChainOutput, chain.BlockChains] {
	// TODO implement me
	return nil
}

func (t *TokenAdapter) RegisterToken() *cldf_ops.Sequence[tokens.RegisterTokenInput, sequences.OnChainOutput, chain.BlockChains] {
	// TODO implement me
	return nil
}

func (t *TokenAdapter) SetPool() *cldf_ops.Sequence[tokens.SetPoolInput, sequences.OnChainOutput, chain.BlockChains] {
	// TODO implement me
	return nil
}

func (t *TokenAdapter) UpdateAuthorities() *cldf_ops.Sequence[tokens.UpdateAuthoritiesInput, sequences.OnChainOutput, chain.BlockChains] {
	// TODO implement me
	return nil
}
