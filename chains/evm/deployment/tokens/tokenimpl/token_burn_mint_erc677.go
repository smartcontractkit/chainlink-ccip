package tokenimpl

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	bnmERC677 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc677"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type tokenBurnMintERC677 struct{}

func (tokenBurnMintERC677) ContractType() deployment.ContractType {
	return cciputils.BurnMintToken
}

func (tokenBurnMintERC677) Capabilities() CapabilitySet {
	return CapabilitySet{
		ParticipatesInPoolRoleGrant: true,
		// Admin tidy uses BurnMintERC20 operations in tidyTokenRoles; ERC677 uses a
		// different binding until dedicated admin ops exist for this type.
		SupportsAdminRole: false,
		SupportsCCIPAdmin: false,
		SupportsPreMint:   false,
	}
}

func (tokenBurnMintERC677) RevokeAdminRole(_ cldf_ops.Bundle, _ evm.Chain, _, _ common.Address) ([]contract.WriteOutput, error) {
	return nil, fmt.Errorf("admin role revoke not supported for BurnMintERC677 token type")
}

func (tokenBurnMintERC677) GrantAdminRole(_ cldf_ops.Bundle, _ evm.Chain, _, _ common.Address) ([]contract.WriteOutput, error) {
	return nil, fmt.Errorf("admin role grant not supported for BurnMintERC677 token type")
}

func (tokenBurnMintERC677) GrantPoolRoles(
	b cldf_ops.Bundle,
	chain evm.Chain,
	token, pool, proposalExecutor common.Address,
) ([]contract.WriteOutput, error) {
	return bnmERC677.PrepareGrantMintAndBurnRoles(
		b,
		chain,
		contract.FunctionInput[common.Address]{
			ChainSelector: chain.Selector,
			Address:       token,
			Args:          pool,
		},
		proposalExecutor,
	)
}

func (tokenBurnMintERC677) SetCCIPAdmin(_ cldf_ops.Bundle, _ evm.Chain, _, _ common.Address) ([]contract.WriteOutput, error) {
	return nil, fmt.Errorf("CCIP admin not supported for BurnMintERC677 token type via this deployment path")
}

func (tokenBurnMintERC677) Transfer(b cldf_ops.Bundle, chain evm.Chain, token, to common.Address, scaledAmount *big.Int) ([]contract.WriteOutput, error) {
	return transferTokensERC20(b, chain, token, to, scaledAmount)
}

func (tokenBurnMintERC677) Deploy(_ cldf_ops.Bundle, _ evm.Chain, _ tokensapi.DeployTokenInput) (datastore.AddressRef, []contract.WriteOutput, error) {
	return datastore.AddressRef{}, nil, fmt.Errorf("deploy BurnMintERC677 token is not implemented in tokenimpl; deploy out of band and record in datastore")
}
