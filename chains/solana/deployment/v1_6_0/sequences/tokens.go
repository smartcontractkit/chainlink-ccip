package sequences

import (
	tokenapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

func (a *SolanaAdapter) ConfigureTokenForTransfersSequence() *cldf_ops.Sequence[tokenapi.ConfigureTokenForTransfersInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}

func (a *SolanaAdapter) AddressRefToBytes(ref datastore.AddressRef) ([]byte, error) {
	return nil, nil
}
func (a *SolanaAdapter) DeriveTokenAddress(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef) ([]byte, error) {
	return nil, nil
}

func (a *SolanaAdapter) ManualRegistration() *cldf_ops.Sequence[tokenapi.ManualRegistrationInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}

func (a *SolanaAdapter) DeployToken() *cldf_ops.Sequence[tokenapi.DeployTokenInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}
func (a *SolanaAdapter) DeployTokenPoolForToken() *cldf_ops.Sequence[tokenapi.DeployTokenPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}
func (a *SolanaAdapter) RegisterToken() *cldf_ops.Sequence[tokenapi.RegisterTokenInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}
func (a *SolanaAdapter) SetPool() *cldf_ops.Sequence[tokenapi.SetPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}
func (a *SolanaAdapter) UpdateAuthorities() *cldf_ops.Sequence[tokenapi.UpdateAuthoritiesInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}
