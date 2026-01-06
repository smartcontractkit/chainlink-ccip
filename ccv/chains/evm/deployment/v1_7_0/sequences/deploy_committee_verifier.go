package sequences

import (
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

type CommitteeVerifierParams struct {
	Version          *semver.Version
	AllowlistAdmin   common.Address
	FeeAggregator    common.Address
	StorageLocations []string
	// Qualifier distinguishes between multiple deployments of the committee verifier and resolver on the same chain.
	Qualifier string
}

type DeployCommitteeVerifierInput struct {
	ChainSelector     uint64
	ExistingAddresses []datastore.AddressRef
	CREATE2Factory    common.Address
	RMN               common.Address
	Params            CommitteeVerifierParams
}

var DeployCommitteeVerifier = cldf_ops.NewSequence(
	"deploy-committee-verifier",
	semver.MustParse("1.7.0"),
	"Deploys the CommitteeVerifier contract",
	func(b cldf_ops.Bundle, chain evm.Chain, input DeployCommitteeVerifierInput) (output sequences.OnChainOutput, err error) {
		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract_utils.WriteOutput, 0)

		// Deploy CommitteeVerifier
		var qualifierPtr *string
		if input.Params.Qualifier != "" {
			qualifierPtr = &input.Params.Qualifier
		}
		committeeVerifierRef, err := contract_utils.MaybeDeployContract(b, committee_verifier.Deploy, chain, contract_utils.DeployInput[committee_verifier.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(committee_verifier.ContractType, *input.Params.Version),
			ChainSelector:  chain.Selector,
			Args: committee_verifier.ConstructorArgs{
				DynamicConfig: committee_verifier.DynamicConfig{
					FeeAggregator:  input.Params.FeeAggregator,
					AllowlistAdmin: input.Params.AllowlistAdmin,
				},
				StorageLocations: input.Params.StorageLocations,
				RMN:              input.RMN,
			},
			Qualifier: qualifierPtr,
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CommitteeVerifier: %w", err)
		}
		addresses = append(addresses, committeeVerifierRef)

		if input.CREATE2Factory == (common.Address{}) {
			return sequences.OnChainOutput{}, errors.New("committee verifier requires CREATE2Factory")
		}
		// Deployment flow via CREATE2Factory
		// First, check if the CommitteeVerifierResolver already exists
		var resolverRef *datastore.AddressRef
		for _, ref := range input.ExistingAddresses {
			if ref.Type == datastore.ContractType(committee_verifier.ResolverType) &&
				ref.Version.String() == committee_verifier.Version.String() &&
				ref.Qualifier == input.Params.Qualifier {
				resolverRef = &ref
			}
		}
		if resolverRef != nil {
			addresses = append(addresses, *resolverRef)
		} else {
			// Otherwise, just deploy the CommitteeVerifierResolver
			committeeVerifierResolverRef, err := contract_utils.MaybeDeployContract(b, versioned_verifier_resolver.Deploy, chain, contract_utils.DeployInput[versioned_verifier_resolver.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(committee_verifier.ResolverType, *committee_verifier.Version),
				ChainSelector:  chain.Selector,
				CREATE2Factory: input.CREATE2Factory,
				Type:           datastore.ContractType(committee_verifier.ResolverType),
				Qualifier:      qualifierPtr,
			}, input.ExistingAddresses)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CommitteeVerifierResolver: %w", err)
			}
			addresses = append(addresses, deployVerifierResolverViaCREATE2Report.Output.Addresses...)
			writes = append(writes, deployVerifierResolverViaCREATE2Report.Output.Writes...)
		}

		batchOp, err := contract_utils.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		return sequences.OnChainOutput{
			Addresses: addresses,
			BatchOps:  []mcms_types.BatchOperation{batchOp},
		}, nil
	})
