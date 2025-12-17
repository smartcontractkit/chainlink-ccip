package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/versioned_verifier_resolver"
	versioned_verifier_resolver_latest "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/versioned_verifier_resolver"
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

		if input.CREATE2Factory != (common.Address{}) {
			// Deployment flow via CREATE2Factory
			// First, check if the CommitteeVerifierResolver already exists
			var resolverRef *datastore.AddressRef
			for _, ref := range input.ExistingAddresses {
				if ref.Type == datastore.ContractType(committee_verifier.ResolverType) &&
					ref.Version.String() == semver.MustParse("1.7.0").String() &&
					ref.Qualifier == input.Params.Qualifier {
					resolverRef = &ref
				}
			}
			if resolverRef != nil {
				addresses = append(addresses, *resolverRef)
			} else {
				additionalAddresses, additionalWrites, err := deployResolverViaCREATE2(b, chain, input)
				if err != nil {
					return sequences.OnChainOutput{}, err
				}
				addresses = append(addresses, additionalAddresses...)
				writes = append(writes, additionalWrites...)
			}
		} else {
			// Otherwise, just deploy the CommitteeVerifierResolver
			committeeVerifierResolverRef, err := contract_utils.MaybeDeployContract(b, committee_verifier.DeployResolver, chain, contract_utils.DeployInput[committee_verifier.ResolverConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(committee_verifier.ResolverType, *semver.MustParse("1.7.0")),
				ChainSelector:  chain.Selector,
				Args:           committee_verifier.ResolverConstructorArgs{},
				Qualifier:      qualifierPtr,
			}, input.ExistingAddresses)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CommitteeVerifierResolver: %w", err)
			}
			addresses = append(addresses, committeeVerifierResolverRef)
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

// deployResolverViaCREATE2 deploys the CommitteeVerifierResolver via CREATE2Factory
// to ensure deterministic addresses. It computes the expected address, deploys and transfers
// ownership to the deployer key, then accepts ownership.
func deployResolverViaCREATE2(
	b cldf_ops.Bundle,
	chain evm.Chain,
	input DeployCommitteeVerifierInput,
) ([]datastore.AddressRef, []contract_utils.WriteOutput, error) {
	addresses := make([]datastore.AddressRef, 0)
	writes := make([]contract_utils.WriteOutput, 0)

	// Deploy CommitteeVerifierResolver via CREATE2Factory to ensure deterministic addresses.
	// Fetch and save the expected address of the CommitteeVerifierResolver.
	expectedAddressReport, err := cldf_ops.ExecuteOperation(b, create2_factory.ComputeAddress, chain, contract_utils.FunctionInput[create2_factory.ComputeAddressArgs]{
		Address:       input.CREATE2Factory,
		ChainSelector: chain.Selector,
		Args: create2_factory.ComputeAddressArgs{
			ABI:             versioned_verifier_resolver_latest.VersionedVerifierResolverMetaData.ABI,
			Bin:             versioned_verifier_resolver_latest.VersionedVerifierResolverMetaData.Bin,
			ConstructorArgs: []any{},
			Salt:            input.Params.Qualifier,
		},
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to compute address for CommitteeVerifierResolver: %w", err)
	}
	addresses = append(addresses, datastore.AddressRef{
		ChainSelector: chain.Selector,
		Address:       expectedAddressReport.Output.Hex(),
		Qualifier:     input.Params.Qualifier,
		Type:          datastore.ContractType(committee_verifier.ResolverType),
		Version:       semver.MustParse("1.7.0"),
	})
	// Then, actually deploy and transfer ownership of the CommitteeVerifierResolver to the deployer key.
	deployAndTransferResolverReport, err := cldf_ops.ExecuteOperation(b, create2_factory.CreateAndTransferOwnership, chain, contract_utils.FunctionInput[create2_factory.CreateAndTransferOwnershipArgs]{
		ChainSelector: chain.Selector,
		Address:       input.CREATE2Factory,
		Args: create2_factory.CreateAndTransferOwnershipArgs{
			ComputeAddressArgs: create2_factory.ComputeAddressArgs{
				ABI:             versioned_verifier_resolver_latest.VersionedVerifierResolverMetaData.ABI,
				Bin:             versioned_verifier_resolver_latest.VersionedVerifierResolverMetaData.Bin,
				ConstructorArgs: []any{},
				Salt:            input.Params.Qualifier,
			},
			To: chain.DeployerKey.From,
		},
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to deploy and transfer ownership of CommitteeVerifierResolver: %w", err)
	}
	writes = append(writes, deployAndTransferResolverReport.Output)
	// Finally, accept ownership of the CommitteeVerifierResolver
	acceptOwnershipReport, err := cldf_ops.ExecuteOperation(b, versioned_verifier_resolver.AcceptOwnership, chain, contract_utils.FunctionInput[versioned_verifier_resolver.AcceptOwnershipArgs]{
		ChainSelector: chain.Selector,
		Address:       common.HexToAddress(expectedAddressReport.Output.Hex()),
		Args: versioned_verifier_resolver.AcceptOwnershipArgs{
			IsProposedOwner: true,
		},
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to accept ownership of CommitteeVerifierResolver: %w", err)
	}
	writes = append(writes, acceptOwnershipReport.Output)

	return addresses, writes, nil
}
