package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/create2_factory"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
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
	Writes    []contract_utils.WriteOutput
}

// DeployContractViaCREATE2 deploys a contract via CREATE2Factory with the given ABI, BIN, and constructor args.
// The contract is deployed and ownership is transferred to the chain deployer key.
var DeployContractViaCREATE2 = cldf_ops.NewSequence(
	"deploy-contract-via-create2",
	semver.MustParse("1.7.0"),
	"Deploys a contract via CREATE2Factory with deterministic address",
	func(b cldf_ops.Bundle, chain evm.Chain, input DeployContractViaCREATE2Input) (output DeployContractViaCREATE2Output, err error) {
		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract_utils.WriteOutput, 0)

		computeArgs := create2_factory.ComputeAddressArgs{
			ABI:             input.ABI,
			Bin:             input.BIN,
			ConstructorArgs: input.ConstructorArgs,
			Salt:            input.Qualifier,
		}

		expectedAddressReport, err := cldf_ops.ExecuteOperation(b, create2_factory.ComputeAddress, chain, contract_utils.FunctionInput[create2_factory.ComputeAddressArgs]{
			Address:       input.CREATE2Factory,
			ChainSelector: chain.Selector,
			Args:          computeArgs,
		})
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

		deployReport, err := cldf_ops.ExecuteOperation(b, create2_factory.CreateAndTransferOwnership, chain, contract_utils.FunctionInput[create2_factory.CreateAndTransferOwnershipArgs]{
			ChainSelector: chain.Selector,
			Address:       input.CREATE2Factory,
			Args: create2_factory.CreateAndTransferOwnershipArgs{
				ComputeAddressArgs: computeArgs,
				To:                 chain.DeployerKey.From,
			},
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
