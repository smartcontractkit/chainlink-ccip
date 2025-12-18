package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/versioned_verifier_resolver"
	versioned_verifier_resolver_latest "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/versioned_verifier_resolver"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type DeployVerifierResolverViaCREATE2Input struct {
	CREATE2Factory common.Address
	ChainSelector  uint64
	Qualifier      string
	Type           datastore.ContractType
	Version        *semver.Version
}

type DeployVerifierResolverViaCREATE2Output struct {
	Addresses []datastore.AddressRef
	Writes    []contract_utils.WriteOutput
}

var DeployVerifierResolverViaCREATE2 = cldf_ops.NewSequence(
	"deploy-verifier-resolver-via-create2",
	semver.MustParse("1.7.0"),
	"Deploys the VerifierResolver contract via CREATE2Factory",
	func(b cldf_ops.Bundle, chain evm.Chain, input DeployVerifierResolverViaCREATE2Input) (output DeployVerifierResolverViaCREATE2Output, err error) {
		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract_utils.WriteOutput, 0)

		// Deploy resolver via CREATE2Factory to ensure deterministic addresses.
		// Fetch and save the expected address of the resolver.
		expectedAddressReport, err := cldf_ops.ExecuteOperation(b, create2_factory.ComputeAddress, chain, contract_utils.FunctionInput[create2_factory.ComputeAddressArgs]{
			Address:       input.CREATE2Factory,
			ChainSelector: chain.Selector,
			Args: create2_factory.ComputeAddressArgs{
				ABI:             versioned_verifier_resolver_latest.VersionedVerifierResolverMetaData.ABI,
				Bin:             versioned_verifier_resolver_latest.VersionedVerifierResolverMetaData.Bin,
				ConstructorArgs: []any{},
				Salt:            input.Qualifier,
			},
		})
		if err != nil {
			return DeployVerifierResolverViaCREATE2Output{}, fmt.Errorf("failed to compute address for resolver: %w", err)
		}
		addresses = append(addresses, datastore.AddressRef{
			ChainSelector: chain.Selector,
			Address:       expectedAddressReport.Output.Hex(),
			Qualifier:     input.Qualifier,
			Type:          input.Type,
			Version:       input.Version,
		})

		// Then, actually deploy and transfer ownership of the resolver to the deployer key.
		deployAndTransferResolverReport, err := cldf_ops.ExecuteOperation(b, create2_factory.CreateAndTransferOwnership, chain, contract_utils.FunctionInput[create2_factory.CreateAndTransferOwnershipArgs]{
			ChainSelector: chain.Selector,
			Address:       input.CREATE2Factory,
			Args: create2_factory.CreateAndTransferOwnershipArgs{
				ComputeAddressArgs: create2_factory.ComputeAddressArgs{
					ABI:             versioned_verifier_resolver_latest.VersionedVerifierResolverMetaData.ABI,
					Bin:             versioned_verifier_resolver_latest.VersionedVerifierResolverMetaData.Bin,
					ConstructorArgs: []any{},
					Salt:            input.Qualifier,
				},
				To: chain.DeployerKey.From,
			},
		})
		if err != nil {
			return DeployVerifierResolverViaCREATE2Output{}, fmt.Errorf("failed to deploy and transfer ownership of resolver: %w", err)
		}
		writes = append(writes, deployAndTransferResolverReport.Output)

		// Finally, accept ownership of the resolver
		acceptOwnershipReport, err := cldf_ops.ExecuteOperation(b, versioned_verifier_resolver.AcceptOwnership, chain, contract_utils.FunctionInput[versioned_verifier_resolver.AcceptOwnershipArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(expectedAddressReport.Output.Hex()),
			Args: versioned_verifier_resolver.AcceptOwnershipArgs{
				IsProposedOwner: true,
			},
		})
		if err != nil {
			return DeployVerifierResolverViaCREATE2Output{}, fmt.Errorf("failed to accept ownership of resolver: %w", err)
		}
		writes = append(writes, acceptOwnershipReport.Output)

		return DeployVerifierResolverViaCREATE2Output{
			Addresses: addresses,
			Writes:    writes,
		}, nil
	})
