package sequences

import (
	evmops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	cv_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/committee_verifier"
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/verifier_tags"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/versioned_verifier_resolver"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var CommitteeVerifierResolverType = versioned_verifier_resolver.CommitteeVerifierResolverType

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
	semver.MustParse("2.0.0"),
	"Deploys the CommitteeVerifier contract",
	func(b cldf_ops.Bundle, chain evm.Chain, input DeployCommitteeVerifierInput) (output sequences.OnChainOutput, err error) {
		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract.WriteOutput, 0)

		// Deploy CommitteeVerifier
		var qualifierPtr *string
		if input.Params.Qualifier != "" {
			qualifierPtr = &input.Params.Qualifier
		}
		committeeVerifierRef, err := evmops.MaybeDeployContract(b, committee_verifier.Deploy, chain, contract.DeployInput[committee_verifier.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(committee_verifier.ContractType, *input.Params.Version),
			Args: committee_verifier.ConstructorArgs{
				DynamicConfig: cv_bindings.CommitteeVerifierDynamicConfig{
					FeeAggregator:  input.Params.FeeAggregator,
					AllowlistAdmin: input.Params.AllowlistAdmin,
				},
				StorageLocations: input.Params.StorageLocations,
				Rmn:              input.RMN,
				VersionTag:       verifier_tags.CommitteeVerifierV2(),
			},
			Qualifier: qualifierPtr,
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CommitteeVerifier: %w", err)
		}
		addresses = append(addresses, committeeVerifierRef)

		// Fetch dynamic config on the CommitteeVerifier
		dynamicConfigReport, err := evmops.ExecuteRead(b, chain, common.HexToAddress(committeeVerifierRef.Address), evmops.BindAs[cv_bindings.CommitteeVerifierInterface](cv_bindings.NewCommitteeVerifier), committee_verifier.NewReadGetDynamicConfig, struct{}{})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get dynamic config on CommitteeVerifier: %w", err)
		}

		// Set dynamic config on the CommitteeVerifier if there is a diff
		desiredFeeAggregator := dynamicConfigReport.Output.FeeAggregator
		if input.Params.FeeAggregator != (common.Address{}) {
			desiredFeeAggregator = input.Params.FeeAggregator
		}
		desiredAllowlistAdmin := dynamicConfigReport.Output.AllowlistAdmin
		if input.Params.AllowlistAdmin != (common.Address{}) {
			desiredAllowlistAdmin = input.Params.AllowlistAdmin
		}
		if desiredFeeAggregator != dynamicConfigReport.Output.FeeAggregator || desiredAllowlistAdmin != dynamicConfigReport.Output.AllowlistAdmin {
			setDynamicConfigReport, err := evmops.ExecuteWrite(b, chain, common.HexToAddress(committeeVerifierRef.Address), evmops.BindAs[cv_bindings.CommitteeVerifierInterface](cv_bindings.NewCommitteeVerifier), committee_verifier.NewWriteSetDynamicConfig, cv_bindings.CommitteeVerifierDynamicConfig{
				AllowlistAdmin: desiredAllowlistAdmin,
				FeeAggregator:  desiredFeeAggregator,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to set dynamic config on CommitteeVerifier: %w", err)
			}
			writes = append(writes, setDynamicConfigReport.Output)
		}

		if input.CREATE2Factory == (common.Address{}) {
			return sequences.OnChainOutput{}, errors.New("committee verifier requires CREATE2Factory")
		}
		// Deployment flow via CREATE2Factory
		// First, check if the CommitteeVerifierResolver already exists
		var resolverRef *datastore.AddressRef
		for _, ref := range input.ExistingAddresses {
			if ref.Type == datastore.ContractType(CommitteeVerifierResolverType) &&
				ref.Version.String() == semver.MustParse("2.0.0").String() &&
				ref.Qualifier == input.Params.Qualifier {
				resolverRef = &ref
			}
		}
		if resolverRef != nil {
			addresses = append(addresses, *resolverRef)
		} else {
			deployVerifierResolverViaCREATE2Report, err := cldf_ops.ExecuteSequence(b, DeployVerifierResolverViaCREATE2, chain, DeployVerifierResolverViaCREATE2Input{
				ChainSelector:  input.ChainSelector,
				Qualifier:      input.Params.Qualifier,
				Type:           datastore.ContractType(CommitteeVerifierResolverType),
				Version:        semver.MustParse("2.0.0"),
				CREATE2Factory: input.CREATE2Factory,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CommitteeVerifierResolver: %w", err)
			}
			addresses = append(addresses, deployVerifierResolverViaCREATE2Report.Output.Addresses...)
			writes = append(writes, deployVerifierResolverViaCREATE2Report.Output.Writes...)
		}

		batchOp, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		return sequences.OnChainOutput{
			Addresses: addresses,
			BatchOps:  []mcms_types.BatchOperation{batchOp},
		}, nil
	})
