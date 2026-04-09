package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	executorops "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/operations/executor"
	onrampops "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/operations/onramp"
	proxyops "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/operations/proxy"
	usdcproxyops "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/operations/usdc_token_pool_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var _ fees.FeeAggregatorAdapter = (*FeeAggregatorAdapter)(nil)

var supportedContractTypes = map[datastore.ContractType]bool{
	datastore.ContractType(onrampops.ContractType):    true,
	datastore.ContractType(proxyops.ContractType):     true,
	datastore.ContractType(executorops.ContractType):  true,
	datastore.ContractType(usdcproxyops.ContractType): true,
}

type FeeAggregatorAdapter struct{}

func NewFeeAggregatorAdapter() *FeeAggregatorAdapter {
	return &FeeAggregatorAdapter{}
}

func (a *FeeAggregatorAdapter) GetFeeAggregator(e cldf.Environment, chainSelector uint64) (string, error) {
	chain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return "", fmt.Errorf("EVM chain with selector %d not defined", chainSelector)
	}

	ref, err := datastore_utils.FindAndFormatRef(
		e.DataStore,
		datastore.AddressRef{
			ChainSelector: chainSelector,
			Type:          datastore.ContractType(proxyops.ContractType),
			Version:       proxyops.Version,
		},
		chainSelector,
		datastore_utils.FullRef,
	)
	if err != nil {
		return "", fmt.Errorf("failed to find Proxy address ref for chain selector %d: %w", chainSelector, err)
	}

	proxyAddr := common.HexToAddress(ref.Address)
	proxyContract, err := proxyops.NewProxyContract(proxyAddr, chain.Client)
	if err != nil {
		return "", fmt.Errorf("failed to instantiate Proxy at %s on chain %d: %w", proxyAddr.Hex(), chainSelector, err)
	}

	feeAgg, err := proxyContract.GetFeeAggregator(&bind.CallOpts{Context: e.GetContext()})
	if err != nil {
		return "", fmt.Errorf("failed to read fee aggregator from Proxy at %s on chain %d: %w", proxyAddr.Hex(), chainSelector, err)
	}

	return feeAgg.Hex(), nil
}

func (a *FeeAggregatorAdapter) SetFeeAggregator(e cldf.Environment) *operations.Sequence[fees.FeeAggregatorForChain, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return operations.NewSequence(
		"SetFeeAggregator",
		semver.MustParse("2.0.0"),
		"Sets the fee aggregator address on CCIP 2.0.0 contracts",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input fees.FeeAggregatorForChain) (sequences.OnChainOutput, error) {
			var result sequences.OnChainOutput

			evmChain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return result, fmt.Errorf("EVM chain with selector %d not defined", input.ChainSelector)
			}

			if !common.IsHexAddress(input.FeeAggregator) {
				return result, fmt.Errorf("invalid fee aggregator address: %s", input.FeeAggregator)
			}
			newFeeAggregator := common.HexToAddress(input.FeeAggregator)

			refs, err := a.resolveRefs(e, input)
			if err != nil {
				return result, err
			}

			for _, ref := range refs {
				writes, err := setFeeAggregatorOnContract(b, evmChain, ref, newFeeAggregator)
				if err != nil {
					return result, fmt.Errorf("failed to set fee aggregator on %s (%s) on chain %d: %w",
						ref.Type, ref.Address, input.ChainSelector, err)
				}
				if len(writes) > 0 {
					batch, err := contract.NewBatchOperationFromWrites(writes)
					if err != nil {
						return result, fmt.Errorf("failed to create batch operation for %s on chain %d: %w",
							ref.Type, input.ChainSelector, err)
					}
					result.BatchOps = append(result.BatchOps, batch)
				}
			}

			return result, nil
		},
	)
}

func (a *FeeAggregatorAdapter) resolveRefs(e cldf.Environment, input fees.FeeAggregatorForChain) ([]datastore.AddressRef, error) {
	if len(input.Contracts) == 0 {
		ref, err := datastore_utils.FindAndFormatRef(
			e.DataStore,
			datastore.AddressRef{Type: datastore.ContractType(proxyops.ContractType), Version: proxyops.Version},
			input.ChainSelector,
			datastore_utils.FullRef,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to find default Proxy ref for chain %d: %w", input.ChainSelector, err)
		}
		return []datastore.AddressRef{ref}, nil
	}

	resolved := make([]datastore.AddressRef, 0, len(input.Contracts))
	for _, c := range input.Contracts {
		if !supportedContractTypes[c.Type] {
			return nil, fmt.Errorf("unsupported contract type %q for fee aggregator on chain %d", c.Type, input.ChainSelector)
		}
		ref, err := datastore_utils.FindAndFormatRef(
			e.DataStore,
			c,
			input.ChainSelector,
			datastore_utils.FullRef,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve ref for %s (version=%v, qualifier=%q) on chain %d: %w",
				c.Type, c.Version, c.Qualifier, input.ChainSelector, err)
		}
		resolved = append(resolved, ref)
	}
	return resolved, nil
}

func setFeeAggregatorOnContract(
	b operations.Bundle,
	chain cldf_evm.Chain,
	ref datastore.AddressRef,
	newFeeAggregator common.Address,
) ([]contract.WriteOutput, error) {
	addr := common.HexToAddress(ref.Address)

	switch ref.Type {
	case datastore.ContractType(proxyops.ContractType):
		return setFeeAggregatorDirect(b, chain, addr, proxyops.SetFeeAggregator, newFeeAggregator)

	case datastore.ContractType(usdcproxyops.ContractType):
		return setFeeAggregatorDirect(b, chain, addr, usdcproxyops.SetFeeAggregator, newFeeAggregator)

	case datastore.ContractType(onrampops.ContractType):
		return setFeeAggregatorViaOnRampDynamicConfig(b, chain, addr, newFeeAggregator)

	case datastore.ContractType(executorops.ContractType):
		return setFeeAggregatorViaExecutorDynamicConfig(b, chain, addr, newFeeAggregator)

	default:
		return nil, fmt.Errorf("no handler for contract type %q", ref.Type)
	}
}

func setFeeAggregatorDirect(
	b operations.Bundle,
	chain cldf_evm.Chain,
	addr common.Address,
	op *operations.Operation[contract.FunctionInput[common.Address], contract.WriteOutput, cldf_evm.Chain],
	newFeeAggregator common.Address,
) ([]contract.WriteOutput, error) {
	report, err := operations.ExecuteOperation(
		b, op, chain,
		contract.FunctionInput[common.Address]{
			ChainSelector: chain.Selector,
			Address:       addr,
			Args:          newFeeAggregator,
		},
	)
	if err != nil {
		return nil, err
	}
	return []contract.WriteOutput{report.Output}, nil
}

func setFeeAggregatorViaOnRampDynamicConfig(
	b operations.Bundle,
	chain cldf_evm.Chain,
	addr common.Address,
	newFeeAggregator common.Address,
) ([]contract.WriteOutput, error) {
	readReport, err := operations.ExecuteOperation(
		b, onrampops.GetDynamicConfig, chain,
		contract.FunctionInput[struct{}]{
			ChainSelector: chain.Selector,
			Address:       addr,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to read OnRamp DynamicConfig: %w", err)
	}

	cfg := readReport.Output
	cfg.FeeAggregator = newFeeAggregator

	writeReport, err := operations.ExecuteOperation(
		b, onrampops.SetDynamicConfig, chain,
		contract.FunctionInput[onrampops.DynamicConfig]{
			ChainSelector: chain.Selector,
			Address:       addr,
			Args:          cfg,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to write OnRamp DynamicConfig: %w", err)
	}
	return []contract.WriteOutput{writeReport.Output}, nil
}

func setFeeAggregatorViaExecutorDynamicConfig(
	b operations.Bundle,
	chain cldf_evm.Chain,
	addr common.Address,
	newFeeAggregator common.Address,
) ([]contract.WriteOutput, error) {
	readReport, err := operations.ExecuteOperation(
		b, executorops.GetDynamicConfig, chain,
		contract.FunctionInput[struct{}]{
			ChainSelector: chain.Selector,
			Address:       addr,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to read Executor DynamicConfig: %w", err)
	}

	cfg := readReport.Output
	cfg.FeeAggregator = newFeeAggregator

	writeReport, err := operations.ExecuteOperation(
		b, executorops.SetDynamicConfig, chain,
		contract.FunctionInput[executorops.DynamicConfig]{
			ChainSelector: chain.Selector,
			Address:       addr,
			Args:          cfg,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to write Executor DynamicConfig: %w", err)
	}
	return []contract.WriteOutput{writeReport.Output}, nil
}
