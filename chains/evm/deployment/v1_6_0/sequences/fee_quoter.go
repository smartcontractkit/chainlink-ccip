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
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_3/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

type FeeQuoterApplyDestChainConfigUpdatesSequenceInput struct {
	Address        common.Address
	ChainSelector  uint64
	UpdatesByChain []fee_quoter.FeeQuoterDestChainConfigArgs
}

type FeeQuoterUpdatePricesSequenceInput struct {
	Address        common.Address
	ChainSelector  uint64
	UpdatesByChain fee_quoter.InternalPriceUpdates
}

type FeeQuoterApplyTokenTransferFeeConfigUpdatesSequenceInput struct {
	Address        common.Address
	ChainSelector  uint64
	UpdatesByChain fqops.ApplyTokenTransferFeeConfigUpdatesInput
}

type FeeQuoterImportConfigSequenceInput struct {
	Address              common.Address
	ChainSelector        uint64
	TokensPerRemoteChain map[uint64][]common.Address
	RemoteChains         []uint64
}

type FqOutput struct {
	DestChainCfg         fee_quoter.FeeQuoterDestChainConfig
	TokenTransferFeeCfgs map[common.Address]fee_quoter.FeeQuoterTokenTransferFeeConfig
}

var (
	FeeQuoterApplyDestChainConfigUpdatesSequence = operations.NewSequence(
		"FeeQuoterApplyDestChainConfigUpdatesSequence",
		semver.MustParse("1.6.0"),
		"Apply updates to destination chain configs on the FeeQuoter 1.6.0 contract across multiple EVM chains",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input FeeQuoterApplyDestChainConfigUpdatesSequenceInput) (sequences.OnChainOutput, error) {
			writes := make([]contract.WriteOutput, 0)
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}
			report, err := operations.ExecuteOperation(b, fqops.FeeQuoterApplyDestChainConfigUpdates, chain, contract.FunctionInput[[]fee_quoter.FeeQuoterDestChainConfigArgs]{
				ChainSelector: chain.Selector,
				Address:       input.Address,
				Args:          input.UpdatesByChain,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute FeeQuoterApplyDestChainConfigUpdatesOp on %s: %w", chain, err)
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

	FeeQuoterUpdatePricesSequence = operations.NewSequence(
		"FeeQuoterUpdatePricesSequence",
		semver.MustParse("1.6.0"),
		"Update token and gas prices on FeeQuoter 1.6.0 contracts on multiple EVM chains",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input FeeQuoterUpdatePricesSequenceInput) (sequences.OnChainOutput, error) {
			writes := make([]contract.WriteOutput, 0)
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}
			report, err := operations.ExecuteOperation(b, fqops.FeeQuoterUpdatePrices, chain, contract.FunctionInput[fee_quoter.InternalPriceUpdates]{
				ChainSelector: chain.Selector,
				Address:       input.Address,
				Args:          input.UpdatesByChain,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute FeeQuoterUpdatePricesOp on %s: %w", chain, err)
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

	FeeQuoterApplyTokenTransferFeeConfigUpdatesSequence = operations.NewSequence(
		"FeeQuoterApplyTokenTransferFeeConfigUpdatesSequence",
		semver.MustParse("1.6.0"),
		"Update token transfer fee configs on FeeQuoter 1.6.0 contracts on multiple EVM chains",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input FeeQuoterApplyTokenTransferFeeConfigUpdatesSequenceInput) (sequences.OnChainOutput, error) {
			writes := make([]contract.WriteOutput, 0)
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}
			report, err := operations.ExecuteOperation(b, fqops.FeeQuoterApplyTokenTransferFeeConfigUpdates, chain, contract.FunctionInput[fqops.ApplyTokenTransferFeeConfigUpdatesInput]{
				ChainSelector: chain.Selector,
				Address:       input.Address,
				Args:          input.UpdatesByChain,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute FeeQuoterApplyTokenTransferFeeConfigUpdatesOp on %s: %w", chain, err)
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

	FeeQuoterImportConfigSequence = operations.NewSequence(
		"FeeQuoterImportConfigSequence",
		semver.MustParse("1.6.0"),
		"Imports FeeQuoter configuration from on-chain contracts across multiple EVM chains",
		func(b operations.Bundle, chains cldf_chain.BlockChains, in FeeQuoterImportConfigSequenceInput) (sequences.OnChainOutput, error) {
			var contractMetadata []datastore.ContractMetadata
			evmChain, ok := chains.EVMChains()[in.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found in environment", in.ChainSelector)
			}
			fqAddress := in.Address
			chainSelector := in.ChainSelector
			b.Logger.Infof("Importing configuration for FeeQuoter %s on chain %d (%s)", fqAddress.Hex(), chainSelector, evmChain.Name())
			fqOutput := make(map[uint64]FqOutput)
			destChainConfigs := make(map[uint64]fee_quoter.FeeQuoterDestChainConfig)
			for _, remoteChain := range in.RemoteChains {
				opsOutput, err := operations.ExecuteOperation(b, fqops.GetDestChainConfig, evmChain, contract.FunctionInput[uint64]{
					Address:       fqAddress,
					ChainSelector: chainSelector,
					Args:          remoteChain,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get dest chain config for "+
						"remote chain %d from feequoter %s on chain %d: %w",
						remoteChain, fqAddress.Hex(), chainSelector, err)
				}
				destChainConfigs[remoteChain] = opsOutput.Output
			}
			tokenTransferFeeCfgs := make(map[common.Address]fee_quoter.FeeQuoterTokenTransferFeeConfig)
			tokenTransferFeeCfgsPerChain := make(map[uint64]map[common.Address]fee_quoter.FeeQuoterTokenTransferFeeConfig)

			for remoteChain, tokens := range in.TokensPerRemoteChain {
				for _, token := range tokens {
					opsOutput, err := operations.ExecuteOperation(b, fqops.GetTokenTransferFeeConfig, evmChain,
						contract.FunctionInput[fqops.GetTokenTransferFeeConfigInput]{
							Address:       fqAddress,
							ChainSelector: chainSelector,
							Args: fqops.GetTokenTransferFeeConfigInput{
								Token:             token,
								DestChainSelector: remoteChain,
							},
						})
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to get token transfer fee config for "+
							"token %s to remote chain %d from feequoter %s on chain %d: %w",
							token.Hex(), remoteChain, fqAddress.Hex(), chainSelector, err)
					}
					if opsOutput.Output.IsEnabled {
						tokenTransferFeeCfgs[token] = opsOutput.Output
					}
				}
				tokenTransferFeeCfgsPerChain[remoteChain] = tokenTransferFeeCfgs
			}
			for remoteChain, destCfg := range destChainConfigs {
				fqOutput[remoteChain] = FqOutput{
					DestChainCfg:         destCfg,
					TokenTransferFeeCfgs: tokenTransferFeeCfgsPerChain[remoteChain],
				}
			}
			contractMetadata = []datastore.ContractMetadata{
				{
					Address:       fqAddress.Hex(),
					ChainSelector: chainSelector,
					Metadata:      fqOutput,
				},
			}
			return sequences.OnChainOutput{
				Metadata: sequences.Metadata{
					Contracts: contractMetadata,
				},
			}, nil
		})
)
