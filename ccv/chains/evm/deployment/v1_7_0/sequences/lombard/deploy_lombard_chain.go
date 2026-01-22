package lombard

import (
	"fmt"
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/lombard_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/type_and_version"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/versioned_verifier_resolver"
	v1_7_0_sequences "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

var indexAddressesByTypeAndVersion = type_and_version.IndexAddressesByTypeAndVersion

var (
	lombardQualifier = "Lombard"
)

var DeployLombardChain = cldf_ops.NewSequence(
	"deploy-lombard-chain",
	semver.MustParse("1.0.0"),
	"Deploys the Lombard chain with all required contracts and configurations",
	func(b cldf_ops.Bundle, chains chain.BlockChains, input adapters.DeployLombardInput[datastore.AddressRef, datastore.AddressRef]) (output sequences.OnChainOutput, err error) {
		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract_utils.WriteOutput, 0)

		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}

		verifierTypeAndVersionToAddr, err := indexAddressesByTypeAndVersion(b, chain, input.CCTPVerifier)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to index addresses by type and version: %w", err)
		}

		lombardVerifierAddress := verifierTypeAndVersionToAddr[deployment.NewTypeAndVersion(lombard_verifier.ContractType, *cctp_verifier.Version).String()]
		lombardVerifierResolverAddress := verifierTypeAndVersionToAddr[deployment.NewTypeAndVersion(versioned_verifier_resolver.ContractType, *versioned_verifier_resolver.Version).String()]

		if lombardVerifierAddress == "" {
			lombardVerifierReport, err := cldf_ops.ExecuteOperation(b, lombard_verifier.Deploy, chain, contract_utils.DeployInput[lombard_verifier.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(lombard_verifier.ContractType, *lombard_verifier.Version),
				ChainSelector:  input.Selector,
				Qualifier:      &lombardQualifier,
				Args: lombard_verifier.ConstructorArgs{
					Bridge:           common.HexToAddress(input.Bridge),
					StorageLocations: input.StorageLocations,
					DynamicConfig: lombard_verifier.DynamicConfig{
						FeeAggregator: common.HexToAddress(input.FeeAggregator),
					},
					RMN: common.HexToAddress(input.RMN),
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy LombardVerifier: %w", err)
			}
			addresses = append(addresses, lombardVerifierReport.Output)
			lombardVerifierAddress = lombardVerifierReport.Output.Address
		}

		if lombardVerifierResolverAddress == "" {
			if input.DeployerContract == "" {
				return sequences.OnChainOutput{}, fmt.Errorf("deployer contract is required")
			}

			deployVerifierResolverViaCREATE2Report, err := cldf_ops.ExecuteSequence(b, v1_7_0_sequences.DeployVerifierResolverViaCREATE2, chain, v1_7_0_sequences.DeployVerifierResolverViaCREATE2Input{
				ChainSelector:  input.Selector,
				Qualifier:      lombardQualifier,
				Type:           datastore.ContractType(lombard_verifier.ResolverType),
				Version:        lombard_verifier.Version,
				CREATE2Factory: common.HexToAddress(input.DeployerContract),
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy LombardVerifierResolver: %w", err)
			}
			addresses = append(addresses, deployVerifierResolverViaCREATE2Report.Output.Addresses...)
			writes = append(writes, deployVerifierResolverViaCREATE2Report.Output.Writes...)
			if len(deployVerifierResolverViaCREATE2Report.Output.Addresses) != 1 {
				return sequences.OnChainOutput{}, fmt.Errorf("expected 1 LombardVerifierResolver address, got %d", len(deployVerifierResolverViaCREATE2Report.Output.Addresses))
			}
			lombardVerifierResolverAddress = deployVerifierResolverViaCREATE2Report.Output.Addresses[0].Address
		}

		return sequences.OnChainOutput{
			Addresses: []datastore.AddressRef{},
			BatchOps:  []mcms_types.BatchOperation{},
		}, nil
	})
