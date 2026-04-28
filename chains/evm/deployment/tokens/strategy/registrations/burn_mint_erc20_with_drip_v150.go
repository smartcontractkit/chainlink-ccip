package registrations

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/tokens/strategy"
	drip_v150 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_erc20_with_drip"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

func init() {
	strategy.GetRegistry().RegisterEVM(burnMintERC20WithDripV150Strategy{})
}

// burnMintERC20WithDripV150Strategy is the v1.5.0 variant of BurnMintERC20WithDrip.
// Pre-mint is unsupported because the v1.5.0 constructor takes neither
// supply nor decimals; matches the historical tokenSupportsPreMint table.
type burnMintERC20WithDripV150Strategy struct{}

func (burnMintERC20WithDripV150Strategy) ContractType() deployment.ContractType {
	return drip_v150.ContractType
}

func (burnMintERC20WithDripV150Strategy) Capabilities() strategy.Capabilities {
	return strategy.Capabilities{
		SupportsAdminRole:           true,
		SupportsCCIPAdmin:           true,
		SupportsPreMint:             false,
		ParticipatesInPoolRoleGrant: true,
	}
}

func (burnMintERC20WithDripV150Strategy) Deploy(b cldf_ops.Bundle, chain evm.Chain, in tokensapi.DeployTokenInput) (datastore.AddressRef, []contract.WriteOutput, error) {
	qualifier := in.Symbol
	ref, err := contract.MaybeDeployContract(b, drip_v150.Deploy, chain, contract.DeployInput[drip_v150.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(drip_v150.ContractType, *drip_v150.Version),
		ChainSelector:  chain.Selector,
		Args: drip_v150.ConstructorArgs{
			Name:   in.Name,
			Symbol: in.Symbol,
		},
		Qualifier: &qualifier,
	}, nil)
	if err != nil {
		return datastore.AddressRef{}, nil, fmt.Errorf("failed to deploy BurnMintERC20WithDrip (v1.5.0) token: %w", err)
	}
	return ref, nil, nil
}

func (burnMintERC20WithDripV150Strategy) GrantPoolRoles(b cldf_ops.Bundle, chain evm.Chain, token, pool common.Address, chainSelector uint64) ([]contract.WriteOutput, error) {
	return grantBnMMintAndBurnRoles(b, chain, token, pool, chainSelector)
}

func (burnMintERC20WithDripV150Strategy) GrantExternalAdmin(b cldf_ops.Bundle, chain evm.Chain, token, externalAdmin common.Address, chainSelector uint64) ([]contract.WriteOutput, error) {
	return grantBnMDefaultAdminRole(b, chain, token, externalAdmin, chainSelector)
}
