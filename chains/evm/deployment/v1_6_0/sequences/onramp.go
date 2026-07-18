package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/onramp"
	evmops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations"
	onrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

type OnRampApplyDestChainConfigUpdatesSequenceInput struct {
	Address        common.Address
	ChainSelector  uint64
	UpdatesByChain []gobindings.OnRampDestChainConfigArgs
}

type OnRampImportConfigSequenceInput struct {
	Address       common.Address
	ChainSelector uint64
	RemoteChains  []uint64
}

type OnRampImportConfigSequenceOutput struct {
	DestChainCfgs map[uint64]gobindings.GetDestChainConfig
	StaticConfig  gobindings.OnRampStaticConfig
	DynamicConfig gobindings.OnRampDynamicConfig
}

var (
	OnRampApplyDestChainConfigUpdatesSequence = operations.NewSequence(
		"OnRampApplyDestChainConfigUpdatesSequence",
		semver.MustParse("1.6.0"),
		"Applies updates to destination chain configurations stored on OnRamp contracts on multiple EVM chains",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input OnRampApplyDestChainConfigUpdatesSequenceInput) (sequences.OnChainOutput, error) {
			writes := make([]contract.WriteOutput, 0)
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}
			report, err := evmops.ExecuteWrite(b, chain, input.Address, evmops.BindAs[gobindings.OnRampInterface](gobindings.NewOnRamp), onrampops.NewWriteApplyDestChainConfigUpdates, input.UpdatesByChain)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute OnRampApplyDestChainConfigUpdatesOp on %s: %w", chain, err)
			}
			writes = append(writes, report.Output)
			batch, err := contract.NewBatchOperationFromWrites(writes)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
			}
			return sequences.OnChainOutput{
				BatchOps: []mcms_types.BatchOperation{batch},
			}, nil
		})

	OnRampImportConfigSequence = operations.NewSequence(
		"OnRampImportConfigSequence",
		semver.MustParse("1.6.0"),
		"Imports OnRamp configuration from TokenAdminRegistry and FeeQuoter contracts into OnRamp contracts on multiple EVM chains",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input OnRampImportConfigSequenceInput) (sequences.OnChainOutput, error) {
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}
			onRampDestConfigs := make(map[uint64]gobindings.GetDestChainConfig)
			for _, remoteChain := range input.RemoteChains {
				report, err := evmops.ExecuteRead(b, chain, input.Address, evmops.BindAs[gobindings.OnRampInterface](gobindings.NewOnRamp), onrampops.NewReadGetDestChainConfig, remoteChain)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get dest chain config for chain %d from OnRamp at %s on %s: %w", remoteChain, input.Address.String(), chain, err)
				}
				onRampDestConfigs[remoteChain] = report.Output
			}
			report, err := evmops.ExecuteRead(b, chain, input.Address, evmops.BindAs[gobindings.OnRampInterface](gobindings.NewOnRamp), onrampops.NewReadGetStaticConfig, struct{}{})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get static config from OnRamp at %s on %s: %w", input.Address.String(), chain, err)
			}
			staticConfig := report.Output
			out, err := evmops.ExecuteRead(b, chain, input.Address, evmops.BindAs[gobindings.OnRampInterface](gobindings.NewOnRamp), onrampops.NewReadGetDynamicConfig, struct{}{})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get dynamic config from OnRamp at %s on %s: %w", input.Address.String(), chain, err)
			}
			dynamicConfig := out.Output
			contractMetadata := datastore.ContractMetadata{
				Address:       input.Address.Hex(),
				ChainSelector: chain.Selector,
				Metadata: OnRampImportConfigSequenceOutput{
					StaticConfig:  staticConfig,
					DynamicConfig: dynamicConfig,
					DestChainCfgs: onRampDestConfigs,
				},
			}
			return sequences.OnChainOutput{
				Metadata: sequences.Metadata{
					Contracts: []datastore.ContractMetadata{contractMetadata},
				},
			}, nil
		})
)
