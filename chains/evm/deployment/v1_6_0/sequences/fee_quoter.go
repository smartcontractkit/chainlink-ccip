package sequences

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"golang.org/x/exp/maps"
	"golang.org/x/sync/errgroup"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	ops2contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	fqops0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	fqbind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

type FeeQuoterApplyDestChainConfigUpdatesSequenceInput struct {
	Address        common.Address
	ChainSelector  uint64
	UpdatesByChain []fqbind.FeeQuoterDestChainConfigArgs
}

type FeeQuoterUpdatePricesSequenceInput struct {
	Address        common.Address
	ChainSelector  uint64
	UpdatesByChain fqbind.InternalPriceUpdates
}

type FeeQuoterApplyTokenTransferFeeConfigUpdatesSequenceInput struct {
	Address        common.Address
	ChainSelector  uint64
	UpdatesByChain fqops0.ApplyTokenTransferFeeConfigUpdatesArgs
}

type FeeQuoterImportConfigSequenceInput struct {
	Address              common.Address
	ChainSelector        uint64
	TokensPerRemoteChain map[uint64][]common.Address
	RemoteChains         []uint64
}

type FeeQuoterImportConfigSequenceOutput struct {
	RemoteChainCfgs map[uint64]FeeQuoterImportConfigSequenceOutputPerRemoteChain
	PriceUpdaters   []common.Address
	StaticCfg       fqbind.FeeQuoterStaticConfig
	TokenPrices     map[common.Address]*big.Int
}

type FeeQuoterImportConfigSequenceOutputPerRemoteChain struct {
	DestChainCfg         fqbind.FeeQuoterDestChainConfig
	TokenTransferFeeCfgs map[common.Address]fqbind.FeeQuoterTokenTransferFeeConfig
	GasPrice             *big.Int
}

var (
	FeeQuoterApplyDestChainConfigUpdatesSequence = operations.NewSequence(
		"FeeQuoterApplyDestChainConfigUpdatesSequence",
		semver.MustParse("1.6.0"),
		"Apply updates to destination chain configs on the FeeQuoter 1.6.0 contract across multiple EVM chains",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input FeeQuoterApplyDestChainConfigUpdatesSequenceInput) (sequences.OnChainOutput, error) {
			writes := make([]ops2contract.WriteOutput, 0)
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}
			fq, err := fqbind.NewFeeQuoter(input.Address, chain.Client)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("bind fee quoter: %w", err)
			}
			report, err := operations.ExecuteOperation(b, fqops0.NewWriteApplyDestChainConfigUpdates(fq), chain, ops2contract.FunctionInput[[]fqbind.FeeQuoterDestChainConfigArgs]{
				Args: input.UpdatesByChain,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute FeeQuoterApplyDestChainConfigUpdatesOp on %s: %w", chain, err)
			}
			writes = append(writes, report.Output)
			batch, err := ops2contract.NewBatchOperationFromWrites(writes)
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
			writes := make([]ops2contract.WriteOutput, 0)
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}
			fq, err := fqbind.NewFeeQuoter(input.Address, chain.Client)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("bind fee quoter: %w", err)
			}
			report, err := operations.ExecuteOperation(b, fqops0.NewWriteUpdatePrices(fq), chain, ops2contract.FunctionInput[fqbind.InternalPriceUpdates]{
				Args: input.UpdatesByChain,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute FeeQuoterUpdatePricesOp on %s: %w", chain, err)
			}
			writes = append(writes, report.Output)
			batch, err := ops2contract.NewBatchOperationFromWrites(writes)
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
			writes := make([]ops2contract.WriteOutput, 0)
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}
			fq, err := fqbind.NewFeeQuoter(input.Address, chain.Client)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("bind fee quoter: %w", err)
			}
			report, err := operations.ExecuteOperation(b, fqops0.NewWriteApplyTokenTransferFeeConfigUpdates(fq), chain, ops2contract.FunctionInput[fqops0.ApplyTokenTransferFeeConfigUpdatesArgs]{
				Args: input.UpdatesByChain,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute FeeQuoterApplyTokenTransferFeeConfigUpdatesOp on %s: %w", chain, err)
			}
			writes = append(writes, report.Output)
			batch, err := ops2contract.NewBatchOperationFromWrites(writes)
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

			fq, err := fqbind.NewFeeQuoter(fqAddress, evmChain.Client)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("bind fee quoter: %w", err)
			}

			fqOutput := make(map[uint64]FeeQuoterImportConfigSequenceOutputPerRemoteChain)
			destChainConfigs := make(map[uint64]fqbind.FeeQuoterDestChainConfig)
			var destChainMu sync.Mutex
			destGrp, _ := errgroup.WithContext(b.GetContext())
			destGrp.SetLimit(10)
			gasPricePerChain := make(map[uint64]*big.Int)
			gasPriceMu := sync.Mutex{}
			// fetch fee tokens
			feeTokensRep, err := operations.ExecuteOperation(b, fqops0.NewReadGetFeeTokens(fq), evmChain, ops2contract.FunctionInput[struct{}]{
				Args: struct{}{},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get fee tokens from feequoter %s on chain %s: %w",
					fqAddress.Hex(), evmChain.String(), err)
			}
			for _, remoteChain := range in.RemoteChains {
				remoteChain := remoteChain
				destGrp.Go(func() error {
					opsOutput, err := operations.ExecuteOperation(b, fqops0.NewReadGetDestChainConfig(fq), evmChain, ops2contract.FunctionInput[uint64]{
						Args: remoteChain,
					})
					if err != nil {
						return fmt.Errorf("failed to get dest chain config for "+
							"remote chain %d from feequoter %s on chain %d: %w",
							remoteChain, fqAddress.Hex(), chainSelector, err)
					}
					if !opsOutput.Output.IsEnabled {
						return nil // skip disabled dest chain configs
					}
					destChainMu.Lock()
					destChainConfigs[remoteChain] = opsOutput.Output
					destChainMu.Unlock()
					gasPriceOutput, err := operations.ExecuteOperation(b, fqops0.NewReadGetDestinationChainGasPrice(fq), evmChain, ops2contract.FunctionInput[uint64]{
						Args: remoteChain,
					})
					if err != nil {
						return fmt.Errorf("failed to get destination chain gas price for "+
							"remote chain %d from feequoter %s on chain %d: %w",
							remoteChain, fqAddress.Hex(), chainSelector, err)
					}
					gasPriceMu.Lock()
					gasPricePerChain[remoteChain] = gasPriceOutput.Output.Value
					gasPriceMu.Unlock()
					return nil
				})
			}
			if err := destGrp.Wait(); err != nil {
				return sequences.OnChainOutput{}, err
			}

			tokenTransferFeeCfgsPerChain := make(map[uint64]map[common.Address]fqbind.FeeQuoterTokenTransferFeeConfig)
			allTokens := make(map[common.Address]struct{})
			var ttfcMu sync.Mutex
			tokenGrp, _ := errgroup.WithContext(b.GetContext())
			tokenGrp.SetLimit(10)
			for remoteChain, tokens := range in.TokensPerRemoteChain {
				remoteChain := remoteChain
				tokens := tokens
				for _, token := range tokens {
					if token == (common.Address{}) {
						continue
					}
					if _, exists := allTokens[token]; !exists {
						allTokens[token] = struct{}{}
					}
				}

				tokenGrp.Go(func() error {
					destChainMu.Lock()
					_, enabled := destChainConfigs[remoteChain]
					destChainMu.Unlock()
					if !enabled {
						return nil // skip if dest chain config is not enabled
					}

					tokenTransferFeeCfgs := make(map[common.Address]fqbind.FeeQuoterTokenTransferFeeConfig)
					var tokenTransferFeeCfgsMu sync.Mutex
					innerTokenGrp, _ := errgroup.WithContext(b.GetContext())
					innerTokenGrp.SetLimit(10)
					for _, token := range tokens {
						token := token
						if token == (common.Address{}) {
							continue
						}
						innerTokenGrp.Go(func() error {
							opsOutput, err := operations.ExecuteOperation(b, fqops0.NewReadGetTokenTransferFeeConfig(fq), evmChain,
								ops2contract.FunctionInput[fqops0.GetTokenTransferFeeConfigArgs]{
									Args: fqops0.GetTokenTransferFeeConfigArgs{
										Token:             token,
										DestChainSelector: remoteChain,
									},
								})
							if err != nil {
								return fmt.Errorf("failed to get token transfer fee config for "+
									"token %s to remote chain %d from feequoter %s on chain %d: %w",
									token.Hex(), remoteChain, fqAddress.Hex(), chainSelector, err)
							}
							if opsOutput.Output.IsEnabled {
								tokenTransferFeeCfgsMu.Lock()
								tokenTransferFeeCfgs[token] = opsOutput.Output
								tokenTransferFeeCfgsMu.Unlock()
							}
							return nil
						})
					}
					if err := innerTokenGrp.Wait(); err != nil {
						return err
					}
					ttfcMu.Lock()
					tokenTransferFeeCfgsPerChain[remoteChain] = tokenTransferFeeCfgs
					ttfcMu.Unlock()
					return nil
				})
			}
			if err := tokenGrp.Wait(); err != nil {
				return sequences.OnChainOutput{}, err
			}
			for remoteChain, destCfg := range destChainConfigs {
				if !destCfg.IsEnabled {
					continue
				}
				fqOutput[remoteChain] = FeeQuoterImportConfigSequenceOutputPerRemoteChain{
					DestChainCfg:         destCfg,
					TokenTransferFeeCfgs: tokenTransferFeeCfgsPerChain[remoteChain],
					GasPrice:             gasPricePerChain[remoteChain],
				}
			}
			staticCfgOutput, err := operations.ExecuteOperation(b, fqops0.NewReadGetStaticConfig(fq), evmChain, ops2contract.FunctionInput[struct{}]{
				Args: struct{}{},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get static config from feequoter %s on chain %d: %w",
					fqAddress.Hex(), chainSelector, err)
			}
			priceUpdaters, err := operations.ExecuteOperation(b, fqops0.NewReadGetAllAuthorizedCallers(fq), evmChain, ops2contract.FunctionInput[struct{}]{
				Args: struct{}{},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get all authorized callers from feequoter %s on chain %d: %w",
					fqAddress.Hex(), chainSelector, err)
			}

			// add fee tokens to all tokens if not present
			for _, token := range feeTokensRep.Output {
				if _, exists := allTokens[token]; !exists {
					allTokens[token] = struct{}{}
				}
			}

			tokenSlice := maps.Keys(allTokens)
			// add fee tokens
			tokenPrices, err := operations.ExecuteOperation(b, fqops0.NewReadGetTokenPrices(fq), evmChain, ops2contract.FunctionInput[[]common.Address]{
				Args: tokenSlice,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token prices from feequoter %s on chain %d: %w",
					fqAddress.Hex(), chainSelector, err)
			}
			tokenPricesPerToken := make(map[common.Address]*big.Int)
			for i, token := range tokenSlice {
				tokenPricesPerToken[token] = tokenPrices.Output[i].Value
			}
			contractMetadata = []datastore.ContractMetadata{
				{
					Address:       fqAddress.Hex(),
					ChainSelector: chainSelector,
					Metadata: FeeQuoterImportConfigSequenceOutput{
						RemoteChainCfgs: fqOutput,
						StaticCfg:       staticCfgOutput.Output,
						PriceUpdaters:   priceUpdaters.Output,
						TokenPrices:     tokenPricesPerToken,
					},
				},
			}
			return sequences.OnChainOutput{
				Metadata: sequences.Metadata{
					Contracts: contractMetadata,
				},
			}, nil
		})
)
