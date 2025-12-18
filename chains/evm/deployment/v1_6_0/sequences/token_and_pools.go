package sequences

import (
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

func (a *EVMAdapter) ConfigureTokenForTransfersSequence() *cldf_ops.Sequence[tokensapi.ConfigureTokenForTransfersInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	// TODO implement me
	return nil
}

func (a *EVMAdapter) AddressRefToBytes(ref datastore.AddressRef) ([]byte, error) {
	// TODO implement me
	return nil, nil
}
func (a *EVMAdapter) DeriveTokenAddress(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef) ([]byte, error) {
	// TODO implement me
	return nil, nil
}

func (a *EVMAdapter) ManualRegistration() *cldf_ops.Sequence[tokensapi.ManualRegistrationInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	// TODO implement me
	return nil
}

func (a *EVMAdapter) DeployToken() *cldf_ops.Sequence[tokensapi.DeployTokenInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return DeployToken
}

func (a *EVMAdapter) DeployTokenVerify(in any) error {
	// TODO implement me
	return nil
}

func (a *EVMAdapter) DeployTokenPoolForToken() *cldf_ops.Sequence[tokensapi.DeployTokenPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	// TODO implement me
	return nil
}

func (a *EVMAdapter) RegisterToken() *cldf_ops.Sequence[tokensapi.RegisterTokenInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	// TODO implement me
	return nil
}

func (a *EVMAdapter) SetPool() *cldf_ops.Sequence[tokensapi.SetPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	// TODO implement me
	return nil
}

func (a *EVMAdapter) UpdateAuthorities() *cldf_ops.Sequence[tokensapi.UpdateAuthoritiesInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	// TODO implement me
	return nil
}
