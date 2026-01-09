package sequences

import (
	"fmt"
	"math/big"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
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

func (a *EVMAdapter) DeployTokenVerify(e deployment.Environment, in any) error {
	input := in.(tokensapi.DeployTokenInput)

	tokenAddr, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, datastore.AddressRef{
		ChainSelector: input.ChainSelector,
		Type:          datastore.ContractType(input.Type),
		Qualifier:     input.Symbol,
	}, input.ChainSelector, datastore_utils.FullRef)
	if err == nil {
		e.OperationsBundle.Logger.Info("Token already deployed at address:", tokenAddr.Address)
		return nil
	}

	// Validate EVM addresses from chain-agnostic input
	if err := utils.ValidateEVMAddress(input.CCIPAdmin, "CCIPAdmin"); err != nil {
		return err
	}
	if len(input.ExternalAdmin) > 1 {
		return fmt.Errorf("only one ExternalAdmin address is supported for EVM chains")
	}
	if err := utils.ValidateEVMAddress(input.ExternalAdmin[0], "ExternalAdmin"); err != nil {
		return err
	}
	// ensuring that decimals is not more than 18
	if input.Decimals > 18 {
		return fmt.Errorf("EVM tokens cannot have more than 18 decimals, got %d", input.Decimals)
	}
	// ensuring that supply and pre-mint are not negative
	if input.Supply != nil && input.Supply.Cmp(big.NewInt(0)) < 0 {
		return fmt.Errorf("token supply cannot be negative, got %v", *input.Supply)
	}
	if input.PreMint != nil && input.PreMint.Cmp(big.NewInt(0)) < 0 {
		return fmt.Errorf("token pre-mint cannot be negative, got %v", *input.PreMint)
	}

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
