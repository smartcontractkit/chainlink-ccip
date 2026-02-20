package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

type OffRampImportConfigSequenceInput struct {
	ChainSelector          uint64
	OffRampsPerRemoteChain map[uint64]common.Address
}

type OffRampImportConfigSequenceOutput struct {
	RemoteChainSelector uint64
	StaticConfig        evm_2_evm_offramp.EVM2EVMOffRampStaticConfig
	DynamicConfig       evm_2_evm_offramp.EVM2EVMOffRampDynamicConfig
}

var OffRampImportConfigSequence = operations.NewSequence(
	"offramp:import-config",
	semver.MustParse("1.5.0"),
	"Import OffRamp 1.5.0 config across multiple EVM chains",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input OffRampImportConfigSequenceInput) (sequences.OnChainOutput, error) {
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
		}
		contractMeta := make([]datastore.ContractMetadata, 0)
		for remoteChain, offRampAddress := range input.OffRampsPerRemoteChain {
			sCfg, err := operations.ExecuteOperation(b, offrampops.OffRampStaticConfig, chain, contract.FunctionInput[any]{
				ChainSelector: chain.Selector,
				Address:       offRampAddress,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get static config from OffRamp %s on chain %s "+
					"for remote chain %d: %w", offRampAddress.Hex(), chain.String(), remoteChain, err)
			}
			dCfg, err := operations.ExecuteOperation(b, offrampops.OffRampDynamicConfig, chain, contract.FunctionInput[any]{
				ChainSelector: chain.Selector,
				Address:       offRampAddress,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get dynamic config from OffRamp %s "+
					"on chain %s for remote chain %d: %w", offRampAddress.Hex(), chain.String(), remoteChain, err)
			}
			contractMeta = append(contractMeta, datastore.ContractMetadata{
				ChainSelector: chain.Selector,
				Address:       offRampAddress.Hex(),
				Metadata: OffRampImportConfigSequenceOutput{
					RemoteChainSelector: remoteChain,
					StaticConfig:        sCfg.Output,
					DynamicConfig:       dCfg.Output,
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
