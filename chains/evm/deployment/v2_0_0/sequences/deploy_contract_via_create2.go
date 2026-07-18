package sequences

import (
	evmops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	create2_factory_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/create2_factory"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/create2_factory"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// DeployContractViaCREATE2Input is input for deploying a contract via CREATE2Factory.
type DeployContractViaCREATE2Input struct {
	CREATE2Factory  common.Address
	ChainSelector   uint64
	Qualifier       string
	Type            datastore.ContractType
	Version         *semver.Version
	ABI             string
	BIN             string
	ConstructorArgs []any
}

// DeployContractViaCREATE2Output is the output of the sequence.
type DeployContractViaCREATE2Output struct {
	Addresses []datastore.AddressRef
	Writes    []contract.WriteOutput
}

// DeployContractViaCREATE2 deploys a contract via CREATE2Factory with the given ABI, BIN, and constructor args.
// The contract is deployed and ownership is transferred to the chain deployer key.
var DeployContractViaCREATE2 = cldf_ops.NewSequence(
	"deploy-contract-via-create2",
	semver.MustParse("2.0.0"),
	"Deploys a contract via CREATE2Factory with deterministic address",
	func(b cldf_ops.Bundle, chain evm.Chain, input DeployContractViaCREATE2Input) (output DeployContractViaCREATE2Output, err error) {
		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract.WriteOutput, 0)

		computeArgs := create2_factory.ComputeAddressArgs{
			ABI:             input.ABI,
			Bin:             input.BIN,
			ConstructorArgs: input.ConstructorArgs,
			Salt:            input.Qualifier,
		}

		expectedAddressReport, err := evmops.ExecuteRead(b, chain, input.CREATE2Factory, create2_factory_bindings.NewCREATE2Factory, create2_factory.NewReadComputeAddress, computeArgs)
		if err != nil {
			return DeployContractViaCREATE2Output{}, fmt.Errorf("failed to compute address: %w", err)
		}
		addresses = append(addresses, datastore.AddressRef{
			ChainSelector: chain.Selector,
			Address:       expectedAddressReport.Output.Hex(),
			Qualifier:     input.Qualifier,
			Type:          input.Type,
			Version:       input.Version,
		})

		deployReport, err := evmops.ExecuteWrite(b, chain, input.CREATE2Factory, create2_factory_bindings.NewCREATE2Factory, create2_factory.NewWriteCreateAndTransferOwnership, create2_factory.CreateAndTransferOwnershipArgs{
			ComputeAddressArgs: computeArgs,
				To:                 chain.DeployerKey.From,
		})
		if err != nil {
			return DeployContractViaCREATE2Output{}, fmt.Errorf("failed to deploy and transfer ownership: %w", err)
		}
		writes = append(writes, deployReport.Output)

		return DeployContractViaCREATE2Output{
			Addresses: addresses,
			Writes:    writes,
		}, nil
	},
)
