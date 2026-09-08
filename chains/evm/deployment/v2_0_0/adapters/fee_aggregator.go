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

	evm_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	executorops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/executor"
	onrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/onramp"
	proxyops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/proxy"
	usdcproxyops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/usdc_token_pool_proxy"
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

			refs, err := a.resolveRefs(e, input.ChainSelector, input.Contracts)
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

func (a *FeeAggregatorAdapter) resolveRefs(e cldf.Environment, chainSelector uint64, contracts []datastore.AddressRef) ([]datastore.AddressRef, error) {
	if len(contracts) == 0 {
		ref, err := datastore_utils.FindAndFormatRef(
			e.DataStore,
			datastore.AddressRef{Type: datastore.ContractType(proxyops.ContractType), Version: proxyops.Version},
			chainSelector,
			datastore_utils.FullRef,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to find default Proxy ref for chain %d: %w", chainSelector, err)
		}
		return []datastore.AddressRef{ref}, nil
	}

	resolved := make([]datastore.AddressRef, 0, len(contracts))
	for _, c := range contracts {
		if !supportedContractTypes[c.Type] {
			return nil, fmt.Errorf("unsupported contract type %q on chain %d", c.Type, chainSelector)
		}
		ref, err := datastore_utils.FindAndFormatRef(
			e.DataStore,
			c,
			chainSelector,
			datastore_utils.FullRef,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve ref for %s (version=%v, qualifier=%q) on chain %d: %w",
				c.Type, c.Version, c.Qualifier, chainSelector, err)
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

func (a *FeeAggregatorAdapter) WithdrawFeeTokens(e cldf.Environment) *operations.Sequence[fees.WithdrawFeeTokensForChain, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return operations.NewSequence(
		"WithdrawFeeTokens",
		semver.MustParse("2.0.0"),
		"Withdraws accumulated fee token balances to the fee aggregator on CCIP 2.0.0 contracts",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input fees.WithdrawFeeTokensForChain) (sequences.OnChainOutput, error) {
			var result sequences.OnChainOutput

			evmChain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return result, fmt.Errorf("EVM chain with selector %d not defined", input.ChainSelector)
			}

			feeTokens, err := evm_utils.EVMFeeTokenAddresses(input.FeeTokens, input.ChainSelector)
			if err != nil {
				return result, err
			}

			refs, err := a.resolveRefs(e, input.ChainSelector, input.Contracts)
			if err != nil {
				return result, err
			}

			if err := verifyFeeAggregatorsAreSet(e, input.ChainSelector, refs); err != nil {
				return result, err
			}

			for _, ref := range refs {
				writes, err := withdrawFeeTokensOnContract(b, evmChain, ref, feeTokens)
				if err != nil {
					return result, fmt.Errorf("failed to withdraw fee tokens on %s (%s) on chain %d: %w",
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

// feeAggregatorOnContract reads the fee aggregator currently configured on a single 2.0 contract.
// Each type stores it differently: Proxy exposes it directly, OnRamp and Executor keep it in their
// dynamic config.
func feeAggregatorOnContract(e cldf.Environment, chain cldf_evm.Chain, ref datastore.AddressRef) (common.Address, error) {
	addr := common.HexToAddress(ref.Address)
	opts := &bind.CallOpts{Context: e.GetContext()}

	switch ref.Type {
	case datastore.ContractType(proxyops.ContractType):
		proxy, err := proxyops.NewProxyContract(addr, chain.Client)
		if err != nil {
			return common.Address{}, fmt.Errorf("failed to instantiate Proxy at %s: %w", addr.Hex(), err)
		}
		feeAgg, err := proxy.GetFeeAggregator(opts)
		if err != nil {
			return common.Address{}, fmt.Errorf("failed to read fee aggregator from Proxy at %s: %w", addr.Hex(), err)
		}
		return feeAgg, nil

	case datastore.ContractType(onrampops.ContractType):
		onRamp, err := onrampops.NewOnRampContract(addr, chain.Client)
		if err != nil {
			return common.Address{}, fmt.Errorf("failed to instantiate OnRamp at %s: %w", addr.Hex(), err)
		}
		dynamicCfg, err := onRamp.GetDynamicConfig(opts)
		if err != nil {
			return common.Address{}, fmt.Errorf("failed to read OnRamp dynamic config at %s: %w", addr.Hex(), err)
		}
		return dynamicCfg.FeeAggregator, nil

	case datastore.ContractType(executorops.ContractType):
		executor, err := executorops.NewExecutorContract(addr, chain.Client)
		if err != nil {
			return common.Address{}, fmt.Errorf("failed to instantiate Executor at %s: %w", addr.Hex(), err)
		}
		dynamicCfg, err := executor.GetDynamicConfig(opts)
		if err != nil {
			return common.Address{}, fmt.Errorf("failed to read Executor dynamic config at %s: %w", addr.Hex(), err)
		}
		return dynamicCfg.FeeAggregator, nil

	default:
		// Kept in step with withdrawFeeTokensOnContract: a type we cannot sweep is also a type
		// whose fee aggregator we have no accessor for.
		return common.Address{}, fmt.Errorf("no fee aggregator accessor for contract type %q", ref.Type)
	}
}

// verifyFeeAggregatorsAreSet fails the withdrawal before any transaction is built if a target
// contract has no fee aggregator. Withdrawing to the zero address reverts on-chain with
// ZeroAddressNotAllowed, which is far harder to act on than this message.
func verifyFeeAggregatorsAreSet(e cldf.Environment, chainSelector uint64, refs []datastore.AddressRef) error {
	chain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return fmt.Errorf("EVM chain with selector %d not defined", chainSelector)
	}

	for _, ref := range refs {
		feeAgg, err := feeAggregatorOnContract(e, chain, ref)
		if err != nil {
			return fmt.Errorf("chain %d: %w", chainSelector, err)
		}
		if feeAgg == (common.Address{}) {
			return fmt.Errorf("fee aggregator is not set on %s at %s (chain %d); set it before withdrawing fee tokens",
				ref.Type, ref.Address, chainSelector)
		}
	}

	return nil
}

func withdrawFeeTokensOnContract(
	b operations.Bundle,
	chain cldf_evm.Chain,
	ref datastore.AddressRef,
	feeTokens []common.Address,
) ([]contract.WriteOutput, error) {
	addr := common.HexToAddress(ref.Address)

	switch ref.Type {
	case datastore.ContractType(proxyops.ContractType):
		return withdrawFeeTokensDirect(b, chain, addr, proxyops.WithdrawFeeTokens, feeTokens)

	case datastore.ContractType(onrampops.ContractType):
		return withdrawFeeTokensDirect(b, chain, addr, onrampops.WithdrawFeeTokens, feeTokens)

	case datastore.ContractType(executorops.ContractType):
		return withdrawFeeTokensDirect(b, chain, addr, executorops.WithdrawFeeTokens, feeTokens)

	default:
		// USDCTokenPoolProxy is a supported fee aggregator target but has no generated
		// withdrawFeeTokens operation, so it cannot be swept through this changeset.
		return nil, fmt.Errorf("no withdrawFeeTokens handler for contract type %q", ref.Type)
	}
}

func withdrawFeeTokensDirect(
	b operations.Bundle,
	chain cldf_evm.Chain,
	addr common.Address,
	op *operations.Operation[contract.FunctionInput[[]common.Address], contract.WriteOutput, cldf_evm.Chain],
	feeTokens []common.Address,
) ([]contract.WriteOutput, error) {
	report, err := operations.ExecuteOperation(
		b, op, chain,
		contract.FunctionInput[[]common.Address]{
			ChainSelector: chain.Selector,
			Address:       addr,
			Args:          feeTokens,
		},
	)
	if err != nil {
		return nil, err
	}
	return []contract.WriteOutput{report.Output}, nil
}
