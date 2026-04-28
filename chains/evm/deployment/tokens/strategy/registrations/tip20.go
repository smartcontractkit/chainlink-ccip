package registrations

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/tokens/strategy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/tip20"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

func init() {
	strategy.GetRegistry().RegisterEVM(tip20Strategy{})
}

// tip20Strategy is the Tempo-only TIP-20 token strategy. Deployed via a
// factory sequence rather than MaybeDeployContract; CCIPAdmin and pre-mint
// do not apply.
type tip20Strategy struct{}

func (tip20Strategy) ContractType() deployment.ContractType { return tip20.ContractType }

func (tip20Strategy) Capabilities() strategy.Capabilities {
	return strategy.Capabilities{
		SupportsAdminRole:           true,
		SupportsCCIPAdmin:           false,
		SupportsPreMint:             false,
		ParticipatesInPoolRoleGrant: true,
	}
}

func (tip20Strategy) Deploy(b cldf_ops.Bundle, chain evm.Chain, in tokensapi.DeployTokenInput) (datastore.AddressRef, []contract.WriteOutput, error) {
	// Initial admin must be the deployer so subsequent ops (e.g. GrantIssuerRole)
	// pass IsAllowedCaller; ExternalAdmin receives DEFAULT_ADMIN_ROLE in a
	// follow-up grant performed by the orchestrating sequence.
	report, err := cldf_ops.ExecuteSequence(b, tip20.Deploy, chain, tip20.FactoryDeployArgs{
		QuoteToken: common.Address{},
		Currency:   "",
		Salt:       [32]byte{},
		Symbol:     in.Symbol,
		Admin:      chain.DeployerKey.From,
		Name:       in.Name,
	})
	if err != nil {
		return datastore.AddressRef{}, nil, fmt.Errorf("failed to deploy TIP20 token via factory: %w", err)
	}
	if len(report.Output.Addresses) == 0 {
		return datastore.AddressRef{}, nil, errors.New("no address returned from TIP20 factory deployment")
	}
	return report.Output.Addresses[0], nil, nil
}

func (tip20Strategy) GrantPoolRoles(b cldf_ops.Bundle, chain evm.Chain, token, pool common.Address, chainSelector uint64) ([]contract.WriteOutput, error) {
	report, err := cldf_ops.ExecuteOperation(b, tip20.GrantIssuerRole, chain, contract.FunctionInput[common.Address]{
		ChainSelector: chainSelector,
		Address:       token,
		Args:          pool,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to grant TIP-20 issuer role: %w", err)
	}
	return []contract.WriteOutput{report.Output}, nil
}

func (tip20Strategy) GrantExternalAdmin(b cldf_ops.Bundle, chain evm.Chain, token, externalAdmin common.Address, chainSelector uint64) ([]contract.WriteOutput, error) {
	report, err := cldf_ops.ExecuteOperation(b, tip20.GrantAdminRole, chain, contract.FunctionInput[common.Address]{
		ChainSelector: chainSelector,
		Address:       token,
		Args:          externalAdmin,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to grant TIP-20 admin role: %w", err)
	}
	return []contract.WriteOutput{report.Output}, nil
}
