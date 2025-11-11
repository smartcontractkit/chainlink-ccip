package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/contract_factory"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/proxy"
	proxy_latest "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
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
	ContractFactory   common.Address
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

		// Deploy CommitteeVerifierResolver
		committeeVerifierResolverRef, err := contract_utils.MaybeDeployContract(b, committee_verifier.DeployResolver, chain, contract.DeployInput[committee_verifier.ResolverConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(committee_verifier.ResolverType, *semver.MustParse("1.7.0")),
			ChainSelector:  chain.Selector,
			Args:           committee_verifier.ResolverConstructorArgs{},
			Qualifier:      qualifierPtr,
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CommitteeVerifierResolver: %w", err)
		}
		addresses = append(addresses, committeeVerifierResolverRef)

		if input.ContractFactory != (common.Address{}) {
			// Deployment flow via ContractFactory
			// First, check if the CommitteeVerifierResolverProxy already exists.
			var resolverProxyRef *datastore.AddressRef
			for _, ref := range input.ExistingAddresses {
				if ref.Type == datastore.ContractType(committee_verifier.ResolverProxyType) &&
					ref.Version.String() == semver.MustParse("1.7.0").String() &&
					ref.Qualifier == input.Params.Qualifier {
					resolverProxyRef = &ref
				}
			}
			if resolverProxyRef != nil {
				addresses = append(addresses, *resolverProxyRef)
			} else {
				// Deploy CommitteeVerifierResolverProxy via ContractFactory to ensure deterministic addresses.
				// Fetch and save the expected address of the CommitteeVerifierResolverProxy.
				expectedAddressReport, err := cldf_ops.ExecuteOperation(b, contract_factory.ComputeAddress, chain, contract.FunctionInput[contract_factory.ComputeAddressArgs]{
					Address:       input.ContractFactory,
					ChainSelector: chain.Selector,
					Args: contract_factory.ComputeAddressArgs{
						ABI:             proxy_latest.ProxyMetaData.ABI,
						Bin:             proxy_latest.ProxyMetaData.Bin,
						ConstructorArgs: []any{common.HexToAddress(committeeVerifierResolverRef.Address)},
						Salt:            input.Params.Qualifier,
					},
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to compute address for CommitteeVerifierResolverProxy: %w", err)
				}
				addresses = append(addresses, datastore.AddressRef{
					ChainSelector: chain.Selector,
					Address:       expectedAddressReport.Output.Hex(),
					Qualifier:     input.Params.Qualifier,
					Type:          datastore.ContractType(committee_verifier.ResolverProxyType),
					Version:       semver.MustParse("1.7.0"),
				})
				// Then, actually deploy and transfer ownership of the CommitteeVerifierResolverProxy to the deployer key.
				deployAndTransferResolverProxyReport, err := cldf_ops.ExecuteOperation(b, contract_factory.CreateAndTransferOwnership, chain, contract.FunctionInput[contract_factory.CreateAndTransferOwnershipArgs]{
					ChainSelector: chain.Selector,
					Address:       input.ContractFactory,
					Args: contract_factory.CreateAndTransferOwnershipArgs{
						ComputeAddressArgs: contract_factory.ComputeAddressArgs{
							ABI:             proxy_latest.ProxyMetaData.ABI,
							Bin:             proxy_latest.ProxyMetaData.Bin,
							ConstructorArgs: []any{common.HexToAddress(committeeVerifierResolverRef.Address)},
							Salt:            input.Params.Qualifier,
						},
						To: chain.DeployerKey.From,
					},
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy and transfer ownership of CommitteeVerifierResolverProxy: %w", err)
				}
				writes = append(writes, deployAndTransferResolverProxyReport.Output)
				// Finally, accept ownership of the CommitteeVerifierResolverProxy
				acceptOwnershipReport, err := cldf_ops.ExecuteOperation(b, proxy.AcceptOwnership, chain, contract.FunctionInput[proxy.AcceptOwnershipArgs]{
					ChainSelector: chain.Selector,
					Address:       common.HexToAddress(expectedAddressReport.Output.Hex()),
					Args: proxy.AcceptOwnershipArgs{
						IsProposedOwner: true,
					},
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to accept ownership of CommitteeVerifierResolverProxy: %w", err)
				}
				writes = append(writes, acceptOwnershipReport.Output)
			}
		} else {
			// Otherwise, just deploy the CommitteeVerifierResolverProxy
			committeeVerifierResolverProxyRef, err := contract_utils.MaybeDeployContract(b, committee_verifier.DeployResolverProxy, chain, contract.DeployInput[committee_verifier.ResolverProxyConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(committee_verifier.ResolverProxyType, *semver.MustParse("1.7.0")),
				ChainSelector:  chain.Selector,
				Args: committee_verifier.ResolverProxyConstructorArgs{
					ResolverAddress: common.HexToAddress(committeeVerifierResolverRef.Address),
				},
				Qualifier: qualifierPtr,
			}, input.ExistingAddresses)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CommitteeVerifierResolverProxy: %w", err)
			}
			addresses = append(addresses, committeeVerifierResolverProxyRef)
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
