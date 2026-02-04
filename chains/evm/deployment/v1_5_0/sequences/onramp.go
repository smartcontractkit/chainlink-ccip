package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

type OnRampSetTokenTransferFeeConfigSequenceInput struct {
	Address        common.Address
	ChainSelector  uint64
	UpdatesByChain onramp.SetTokenTransferFeeConfigInput
}

type OnRampImportConfigSequenceInput struct {
	ChainSelector           uint64
	OnRampsPerRemoteChain   map[uint64]common.Address
	SupportedTokensPerChain map[uint64][]common.Address
}

type OnRampImportConfigSequenceOutput struct {
	RemoteChainSelector    uint64
	TokenTransferFeeConfig map[common.Address]evm_2_evm_onramp.EVM2EVMOnRampTokenTransferFeeConfig
	StaticConfig           evm_2_evm_onramp.EVM2EVMOnRampStaticConfig
	DynamicConfig          evm_2_evm_onramp.EVM2EVMOnRampDynamicConfig
}

var (
	OnRampSetTokenTransferFeeConfigSequence = operations.NewSequence(
		"onramp:set-token-transfer-fee-config",
		semver.MustParse("1.5.0"),
		"Set token transfer fee config on the OnRamp 1.5.0 contract across multiple EVM chains",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input OnRampSetTokenTransferFeeConfigSequenceInput) (sequences.OnChainOutput, error) {
			writes := make([]contract.WriteOutput, 0)
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}
			report, err := operations.ExecuteOperation(b, onramp.OnRampSetTokenTransferFeeConfig, chain, contract.FunctionInput[onramp.SetTokenTransferFeeConfigInput]{
				ChainSelector: chain.Selector,
				Address:       input.Address,
				Args:          input.UpdatesByChain,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute OnRampSetTokenTransferFeeConfigOp on %s: %w", chain, err)
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
		semver.MustParse("1.5.0"),
		"Imports the OnRamp configuration from remote chains into the OnRamp 1.5.0 contract on a given EVM chain",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input OnRampImportConfigSequenceInput) (sequences.OnChainOutput, error) {
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}
			contractMeta := make([]datastore.ContractMetadata, 0)
			for remoteChainSelector, onRampAddress := range input.OnRampsPerRemoteChain {
				sCfgOut, err := operations.ExecuteOperation(b, onramp.OnRampStaticConfig, chain, contract.FunctionInput[any]{
					ChainSelector: chain.Selector,
					Address:       onRampAddress,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to execute OnRampStaticConfigOp "+
						"on %s for remote chain %d: %w", chain.String(), remoteChainSelector, err)
				}
				dCfgOut, err := operations.ExecuteOperation(b, onramp.OnRampDynamicConfig, chain, contract.FunctionInput[any]{
					ChainSelector: chain.Selector,
					Address:       onRampAddress,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to execute OnRampDynamicConfigOp "+
						"on %s for remote chain %d: %w", chain.String(), remoteChainSelector, err)
				}
				tokenTransferFeeConfig := make(map[common.Address]evm_2_evm_onramp.EVM2EVMOnRampTokenTransferFeeConfig)
				for _, token := range input.SupportedTokensPerChain[remoteChainSelector] {
					ttfcOut, err := operations.ExecuteOperation(b, onramp.OnRampGetTokenTransferFeeConfig, chain, contract.FunctionInput[common.Address]{
						ChainSelector: chain.Selector,
						Address:       onRampAddress,
						Args:          token,
					})
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to execute OnRampGetTokenTransferFeeConfigOp "+
							"on %s for remote chain %d and token %s: %w", chain.String(), remoteChainSelector, token.String(), err)
					}
					tokenTransferFeeConfig[token] = ttfcOut.Output
				}
				contractMeta = append(contractMeta, datastore.ContractMetadata{
					ChainSelector: input.ChainSelector,
					Address:       onRampAddress.Hex(),
					Metadata: OnRampImportConfigSequenceOutput{
						RemoteChainSelector:    remoteChainSelector,
						StaticConfig:           sCfgOut.Output,
						DynamicConfig:          dCfgOut.Output,
						TokenTransferFeeConfig: tokenTransferFeeConfig,
					},
				})
			}
			return sequences.OnChainOutput{
				Metadata: sequences.Metadata{
					Contracts: contractMeta,
				},
			}, nil
		},
	)
)
