package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

type CommitteeVerifierParams struct {
	Version             *semver.Version
	AllowlistAdmin      common.Address
	FeeAggregator       common.Address
	SignatureConfigArgs committee_verifier.SetSignatureConfigArgs
	StorageLocation     string
	// Qualifier distinguishes between multiple deployments of the committee verifier and proxy
	// on the same chain.
	Qualifier string
}

type DeployCommitteeVerifierInput struct {
	ChainSelector     uint64
	ExistingAddresses []datastore.AddressRef
	Params            CommitteeVerifierParams
}

var DeployCommitteeVerifier = cldf_ops.NewSequence(
	"deploy-committee-verifier",
	semver.MustParse("1.7.0"),
	"Deploys the CommitteeVerifier contract",
	func(b operations.Bundle, chain evm.Chain, input DeployCommitteeVerifierInput) (output sequences.OnChainOutput, err error) {
		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract.WriteOutput, 0)

		// Deploy CommitteeVerifier
		var qualifierPtr *string
		if input.Params.Qualifier != "" {
			qualifierPtr = &input.Params.Qualifier
		}
		committeeVerifierRef, err := contract_utils.MaybeDeployContract(b, committee_verifier.Deploy, chain, contract.DeployInput[committee_verifier.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(committee_verifier.ContractType, *input.Params.Version),
			ChainSelector:  chain.Selector,
			Args: committee_verifier.ConstructorArgs{
				DynamicConfig: committee_verifier.DynamicConfig{
					FeeAggregator:  input.Params.FeeAggregator,
					AllowlistAdmin: input.Params.AllowlistAdmin,
				},
				StorageLocation: input.Params.StorageLocation,
			},
			Qualifier: qualifierPtr,
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CommitteeVerifier: %w", err)
		}
		addresses = append(addresses, committeeVerifierRef)

		// Set signature config on the CommitteeVerifier
		setSignatureConfigReport, err := cldf_ops.ExecuteOperation(b, committee_verifier.SetSignatureConfigs, chain, contract.FunctionInput[committee_verifier.SetSignatureConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(committeeVerifierRef.Address),
			Args:          input.Params.SignatureConfigArgs,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to set signature config on CommitteeVerifier: %w", err)
		}
		writes = append(writes, setSignatureConfigReport.Output)

		// Deploy CommitteeVerifierProxy
		committeeVerifierProxyRef, err := contract_utils.MaybeDeployContract(b, committee_verifier.DeployProxy, chain, contract.DeployInput[committee_verifier.ProxyConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(committee_verifier.ProxyType, *semver.MustParse("1.7.0")),
			ChainSelector:  chain.Selector,
			Args: committee_verifier.ProxyConstructorArgs{
				RampAddress: common.HexToAddress(committeeVerifierRef.Address),
			},
			Qualifier: qualifierPtr,
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CommitteeVerifierProxy: %w", err)
		}
		addresses = append(addresses, committeeVerifierProxyRef)

		batchOp, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		return sequences.OnChainOutput{
			Addresses: addresses,
			BatchOps:  []mcms_types.BatchOperation{batchOp},
		}, nil
	})
