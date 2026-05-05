package tokenimpl

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	drip_v100 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20_with_drip"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// Deprecated: BurnMintERC20WithDripToken has no actual drip functionality - it
// is retained only for compatibility with existing tests and should be removed
// in a future cleanup.
type tokenBurnMintERC20WithDripV1_0_0 struct{}

func (tokenBurnMintERC20WithDripV1_0_0) ContractType() deployment.ContractType {
	return drip_v100.ContractType
}

func (tokenBurnMintERC20WithDripV1_0_0) Capabilities() CapabilitySet {
	return CapabilitySet{
		ParticipatesInPoolRoleGrant: true,
		SupportsAdminRole:           true,
		SupportsCCIPAdmin:           true,
		SupportsPreMint:             true,
	}
}

func (tokenBurnMintERC20WithDripV1_0_0) RevokeAdminRole(b cldf_ops.Bundle, chain evm.Chain, token, externalAdmin common.Address) ([]contract.WriteOutput, error) {
	return revokeDefaultAdminRoleBurnMintERC20(b, chain, token, externalAdmin)
}

func (tokenBurnMintERC20WithDripV1_0_0) GrantAdminRole(b cldf_ops.Bundle, chain evm.Chain, token, externalAdmin common.Address) ([]contract.WriteOutput, error) {
	return grantDefaultAdminRoleBurnMintERC20(b, chain, token, externalAdmin)
}

func (tokenBurnMintERC20WithDripV1_0_0) GrantPoolRoles(b cldf_ops.Bundle, chain evm.Chain, token, pool, _ common.Address) ([]contract.WriteOutput, error) {
	return grantMintAndBurnRolesBurnMintERC20(b, chain, token, pool)
}

func (tokenBurnMintERC20WithDripV1_0_0) SetCCIPAdmin(b cldf_ops.Bundle, chain evm.Chain, token, ccipAdmin common.Address) ([]contract.WriteOutput, error) {
	return setCCIPAdminBurnMintERC20(b, chain, token, ccipAdmin)
}

func (tokenBurnMintERC20WithDripV1_0_0) Transfer(b cldf_ops.Bundle, chain evm.Chain, token, to common.Address, scaledAmount *big.Int) ([]contract.WriteOutput, error) {
	// NOTE: BnM ERC20 drip tokens inherit from a standard ERC20 implementation, so we can use the same transfer helper function as the plain ERC20 token.
	return transferTokensERC20(b, chain, token, to, scaledAmount)
}

func (tokenBurnMintERC20WithDripV1_0_0) Deploy(b cldf_ops.Bundle, chain evm.Chain, in tokensapi.DeployTokenInput) (datastore.AddressRef, []contract.WriteOutput, error) {
	maxSupply := big.NewInt(0)
	if in.Supply != nil {
		maxSupply = tokensapi.ScaleTokenAmount(new(big.Int).SetUint64(*in.Supply), in.Decimals)
	}

	preMint := big.NewInt(0)
	if in.PreMint != nil {
		preMint = tokensapi.ScaleTokenAmount(new(big.Int).SetUint64(*in.PreMint), in.Decimals)
	}

	ref, err := contract.MaybeDeployContract(b, drip_v100.Deploy, chain,
		contract.DeployInput[drip_v100.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(drip_v100.ContractType, *common_utils.Version_1_0_0),
			ChainSelector:  chain.Selector,
			Qualifier:      &in.Symbol,
			Args: drip_v100.ConstructorArgs{
				Name:      in.Name,
				Symbol:    in.Symbol,
				Decimals:  in.Decimals,
				MaxSupply: maxSupply,
				PreMint:   preMint,
			},
		},
		nil,
	)
	if err != nil {
		return datastore.AddressRef{}, nil, fmt.Errorf("failed to deploy BurnMintERC20WithDripToken token: %w", err)
	}

	return ref, nil, nil
}
