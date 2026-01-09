package tokens

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var DeployBurnMintTokenPool = cldf_ops.NewSequence(
	"deploy-burn-mint-token-pool",
	semver.MustParse("1.7.0"),
	"Deploys a burn mint token pool to an EVM chain",
	func(b cldf_ops.Bundle, chain evm.Chain, input DeployTokenPoolInput) (output sequences.OnChainOutput, err error) {
		if err := input.Validate(chain); err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("invalid input: %w", err)
		}

		typeAndVersion := deployment.NewTypeAndVersion(
			deployment.ContractType(input.TokenPoolType),
			*input.TokenPoolVersion,
		)
		tpDeployReport, err := cldf_ops.ExecuteOperation(b, burn_mint_token_pool.Deploy, chain, evm_contract.DeployInput[burn_mint_token_pool.ConstructorArgs]{
			ChainSelector:  input.ChainSel,
			TypeAndVersion: typeAndVersion,
			Args: burn_mint_token_pool.ConstructorArgs{
				Token:              input.ConstructorArgs.Token,
				LocalTokenDecimals: input.ConstructorArgs.Decimals,
				Allowlist:          input.ConstructorArgs.Allowlist,
				RMNProxy:           input.ConstructorArgs.RMNProxy,
				Router:             input.ConstructorArgs.Router,
			},
			Qualifier: &input.TokenSymbol,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy %s to %s: %w", typeAndVersion, chain, err)
		}

		configureReport, err := cldf_ops.ExecuteSequence(b, ConfigureTokenPool, chain, ConfigureTokenPoolInput{
			ChainSelector:    input.ChainSel,
			TokenPoolAddress: common.HexToAddress(tpDeployReport.Output.Address),
			RateLimitAdmin:   input.RateLimitAdmin,
			AllowList:        input.ConstructorArgs.Allowlist,
			RouterAddress:    input.ConstructorArgs.Router,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token pool with address %s on %s: %w", tpDeployReport.Output.Address, chain, err)
		}

		return sequences.OnChainOutput{
			Addresses: []datastore.AddressRef{tpDeployReport.Output},
			BatchOps:  configureReport.Output.BatchOps,
		}, nil
	},
)
