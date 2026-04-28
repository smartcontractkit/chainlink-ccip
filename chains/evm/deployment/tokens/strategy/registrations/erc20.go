package registrations

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/tokens/strategy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

func init() {
	strategy.GetRegistry().RegisterEVM(erc20Strategy{})
}

// erc20Strategy is the plain (non-CCIP-aware) ERC20 strategy.
type erc20Strategy struct{}

func (erc20Strategy) ContractType() deployment.ContractType { return erc20.ContractType }

func (erc20Strategy) Capabilities() strategy.Capabilities {
	return strategy.Capabilities{}
}

func (erc20Strategy) Deploy(b cldf_ops.Bundle, chain evm.Chain, in tokensapi.DeployTokenInput) (datastore.AddressRef, []contract.WriteOutput, error) {
	qualifier := in.Symbol
	ref, err := contract.MaybeDeployContract(b, erc20.Deploy, chain, contract.DeployInput[erc20.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(erc20.ContractType, *common_utils.Version_1_0_0),
		ChainSelector:  chain.Selector,
		Args: erc20.ConstructorArgs{
			Name:   in.Name,
			Symbol: in.Symbol,
		},
		Qualifier: &qualifier,
	}, nil)
	if err != nil {
		return datastore.AddressRef{}, nil, fmt.Errorf("failed to deploy ERC20 token: %w", err)
	}
	return ref, nil, nil
}

func (erc20Strategy) GrantPoolRoles(_ cldf_ops.Bundle, _ evm.Chain, _, _ common.Address, _ uint64) ([]contract.WriteOutput, error) {
	return nil, nil
}

func (erc20Strategy) GrantExternalAdmin(_ cldf_ops.Bundle, _ evm.Chain, _, _ common.Address, _ uint64) ([]contract.WriteOutput, error) {
	return nil, nil
}
