package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	evmseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	onrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var _ fees.FeeAggregatorAdapter = (*FeeAggregatorAdapter)(nil)

type FeeAggregatorAdapter struct {
	evm *evmseq.EVMAdapter
}

func NewFeeAggregatorAdapter(evmAdapter *evmseq.EVMAdapter) *FeeAggregatorAdapter {
	return &FeeAggregatorAdapter{
		evm: evmAdapter,
	}
}

func (a *FeeAggregatorAdapter) getOnRampRef(e cldf.Environment, chainSelector uint64) (datastore.AddressRef, error) {
	filter := datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(onrampops.ContractType),
	}
	ref, err := datastore_utils.FindAndFormatRef(
		e.DataStore,
		filter,
		chainSelector,
		datastore_utils.FullRef,
	)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to find OnRamp address ref for chain selector %d: %w", chainSelector, err)
	}
	return ref, nil
}

func (a *FeeAggregatorAdapter) GetFeeAggregator(e cldf.Environment, chainSelector uint64) (string, error) {
	chain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return "", fmt.Errorf("EVM chain with selector %d not defined", chainSelector)
	}

	onRampRef, err := a.getOnRampRef(e, chainSelector)
	if err != nil {
		return "", err
	}

	onRampAddr := common.HexToAddress(onRampRef.Address)
	onRamp, err := onrampops.NewOnRampContract(onRampAddr, chain.Client)
	if err != nil {
		return "", fmt.Errorf("failed to instantiate OnRamp at %s on chain %d: %w", onRampAddr.Hex(), chainSelector, err)
	}

	dynamicCfg, err := onRamp.GetDynamicConfig(&bind.CallOpts{Context: e.GetContext()})
	if err != nil {
		return "", fmt.Errorf("failed to get OnRamp dynamic config on chain %d: %w", chainSelector, err)
	}

	return dynamicCfg.FeeAggregator.Hex(), nil
}

func (a *FeeAggregatorAdapter) resolveOnRampRef(e cldf.Environment, input fees.SetFeeAggregatorSequenceInput) (datastore.AddressRef, error) {
	if len(input.Contracts) > 0 {
		if len(input.Contracts) != 1 {
			return datastore.AddressRef{}, fmt.Errorf("EVM 1.6 adapter supports exactly one contract ref, got %d", len(input.Contracts))
		}
		ref := input.Contracts[0]
		if ref.Type != datastore.ContractType(onrampops.ContractType) {
			return datastore.AddressRef{}, fmt.Errorf("EVM 1.6 adapter only supports contract type %q, got %q", onrampops.ContractType, ref.Type)
		}
		return datastore_utils.FindAndFormatRef(e.DataStore, ref, input.ChainSelector, datastore_utils.FullRef)
	}
	return a.getOnRampRef(e, input.ChainSelector)
}

func (a *FeeAggregatorAdapter) SetFeeAggregator(e cldf.Environment) *operations.Sequence[fees.SetFeeAggregatorSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return operations.NewSequence(
		"SetFeeAggregator",
		semver.MustParse("1.6.0"),
		"Sets the fee aggregator address on CCIP 1.6.0 OnRamp dynamic config",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input fees.SetFeeAggregatorSequenceInput) (sequences.OnChainOutput, error) {
			var result sequences.OnChainOutput

			evmChain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return result, fmt.Errorf("EVM chain with selector %d not defined", input.ChainSelector)
			}

			if !common.IsHexAddress(input.FeeAggregator) {
				return result, fmt.Errorf("invalid fee aggregator address: %s", input.FeeAggregator)
			}
			newFeeAggregator := common.HexToAddress(input.FeeAggregator)

			onRampRef, err := a.resolveOnRampRef(e, input)
			if err != nil {
				return result, err
			}
			onRampAddr := common.HexToAddress(onRampRef.Address)

			readReport, err := operations.ExecuteOperation(
				b, onrampops.GetDynamicConfig, evmChain,
				contract.FunctionInput[struct{}]{
					ChainSelector: evmChain.Selector,
					Address:       onRampAddr,
				},
			)
			if err != nil {
				return result, fmt.Errorf("failed to read OnRamp dynamic config on chain %d: %w", input.ChainSelector, err)
			}

			currentCfg := readReport.Output
			currentCfg.FeeAggregator = newFeeAggregator

			writeReport, err := operations.ExecuteOperation(
				b, onrampops.SetDynamicConfig, evmChain,
				contract.FunctionInput[onrampops.DynamicConfig]{
					ChainSelector: evmChain.Selector,
					Address:       onRampAddr,
					Args:          currentCfg,
				},
			)
			if err != nil {
				return result, fmt.Errorf("failed to set OnRamp dynamic config on chain %d: %w", input.ChainSelector, err)
			}

			batch, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{writeReport.Output})
			if err != nil {
				return result, fmt.Errorf("failed to create batch operation from writes for chain %d: %w", input.ChainSelector, err)
			}
			result.BatchOps = append(result.BatchOps, batch)

			return result, nil
		},
	)
}
