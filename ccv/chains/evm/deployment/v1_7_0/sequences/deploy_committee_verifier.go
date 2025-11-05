package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/ownable_deployer"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/proxy"
	proxy_latest "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-common/pkg/hashutil"
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
	OwnableDeployer     common.Address
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

		// Deploy CommitteeVerifierResolverProxy via OwnableDeployer to ensure deterministic addresses
		// We use the qualifier string as a differentiator between multiple committee verifier instances on the same chain.
		// First, pre-compute the expected address of the CommitteeVerifierResolverProxy using the OwnableDeployer contract.
		hasher := hashutil.NewKeccak()
		salt := hasher.Hash([]byte(input.Params.Qualifier))
		expectedAddressReport, err := cldf_ops.ExecuteOperation(b, ownable_deployer.ComputeAddress, chain, contract.FunctionInput[ownable_deployer.ComputeAddressArgs]{
			Address:       input.Params.OwnableDeployer,
			ChainSelector: chain.Selector,
			Args: ownable_deployer.ComputeAddressArgs{
				Sender:   chain.DeployerKey.From,
				InitCode: common.FromHex(proxy_latest.ProxyBin),
				Salt:     salt,
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
		// Then, actually deploy and transfer ownership of the CommitteeVerifierResolverProxy
		// Parse the ABI to properly encode constructor arguments
		parsedABI, err := proxy_latest.ProxyMetaData.GetAbi()
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to parse Proxy ABI: %w", err)
		}

		// ABI-encode the constructor arguments (target address)
		targetAddress := common.HexToAddress(committeeVerifierResolverRef.Address)
		constructorArgs, err := parsedABI.Pack("", targetAddress)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to pack constructor arguments: %w", err)
		}

		// Combine bytecode with encoded constructor arguments
		creationCode := common.FromHex(proxy_latest.ProxyBin)
		creationCode = append(creationCode, constructorArgs...)
		deployAndTransferResolverProxyReport, err := cldf_ops.ExecuteOperation(b, ownable_deployer.DeployAndTransferOwnership, chain, contract.FunctionInput[ownable_deployer.DeployAndTransferOwnershipArgs]{
			ChainSelector: chain.Selector,
			Address:       input.Params.OwnableDeployer,
			Args: ownable_deployer.DeployAndTransferOwnershipArgs{
				InitCode: creationCode,
				Salt:     salt,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy and transfer ownership of CommitteeVerifierResolverProxy: %w", err)
		}
		writes = append(writes, deployAndTransferResolverProxyReport.Output)
		// Finally, accept ownership of the CommitteeVerifierResolverProxy
		acceptOwnershipReport, err := cldf_ops.ExecuteOperation(b, proxy.AcceptOwnership, chain, contract.FunctionInput[any]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(expectedAddressReport.Output.Hex()),
			Args:          nil,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to accept ownership of CommitteeVerifierResolverProxy: %w", err)
		}
		writes = append(writes, acceptOwnershipReport.Output)

		batchOp, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		return sequences.OnChainOutput{
			Addresses: addresses,
			BatchOps:  []mcms_types.BatchOperation{batchOp},
		}, nil
	})
