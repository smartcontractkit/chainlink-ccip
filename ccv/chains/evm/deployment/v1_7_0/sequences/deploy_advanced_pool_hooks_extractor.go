package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/advanced_pool_hooks_extractor"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type DeployAdvancedPoolHooksExtractorInput struct {
	ChainSelector     uint64
	ExistingAddresses []datastore.AddressRef
}

var DeployAdvancedPoolHooksExtractor = cldf_ops.NewSequence(
	"deploy-advanced-pool-hooks-extractor",
	semver.MustParse("1.7.0"),
	"Deploys the AdvancedPoolHooksExtractor contract",
	func(b cldf_ops.Bundle, chain evm.Chain, input DeployAdvancedPoolHooksExtractorInput) (sequences.OnChainOutput, error) {
		ref, err := contract_utils.MaybeDeployContract(b, advanced_pool_hooks_extractor.Deploy, chain, contract_utils.DeployInput[advanced_pool_hooks_extractor.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(advanced_pool_hooks_extractor.ContractType, *advanced_pool_hooks_extractor.Version),
			ChainSelector:  chain.Selector,
			Args:           advanced_pool_hooks_extractor.ConstructorArgs{},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy AdvancedPoolHooksExtractor: %w", err)
		}

		return sequences.OnChainOutput{
			Addresses: []datastore.AddressRef{ref},
		}, nil
	},
)
