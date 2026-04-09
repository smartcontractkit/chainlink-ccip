package sequences

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	priceregistryops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/price_registry"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

type PriceRegistryImportConfigSequenceInput struct {
	ChainSelector   uint64
	PriceRegistry   common.Address
	SupportedTokens []common.Address
	RemoteChains    []uint64
}

type PriceRegistryImportConfigSequenceOutput struct {
	GasPrices   map[uint64]*big.Int
	TokenPrices map[common.Address]*big.Int
}

var PriceRegistryImportConfigSequence = operations.NewSequence(
	"PriceRegistryImportTokenAndGasPricesSequence",
	semver.MustParse("1.2.0"),
	"Imports token and gas prices from the Price Registry contract on an EVM chain",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input PriceRegistryImportConfigSequenceInput) (sequences.OnChainOutput, error) {
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
		}
		var allTokens []common.Address
		// check if any supported token is blank address, if so delete
		for i, token := range input.SupportedTokens {
			if token == (common.Address{}) {
				continue
			}
			allTokens = append(allTokens, input.SupportedTokens[i])
		}
		gasPrices := make(map[uint64]*big.Int)
		for _, remoteChainSelector := range input.RemoteChains {
			gasPricesOutput, err := operations.ExecuteOperation(b, priceregistryops.PriceRegistryGetDestinationChainGasPrice, chain, contract.FunctionInput[uint64]{
				ChainSelector: chain.Selector,
				Address:       input.PriceRegistry,
				Args:          remoteChainSelector,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute PriceRegistryGetDestinationChainGasPriceOp "+
					"on %s for price registry %s and remote chain %d: %w", chain.String(), input.PriceRegistry.String(), remoteChainSelector, err)
			}
			gasPrices[remoteChainSelector] = gasPricesOutput.Output.Value
		}
		// get fee tokens
		feetokensRep, err := operations.ExecuteOperation(b, priceregistryops.PriceRegistryGetFeeToken, chain, contract.FunctionInput[any]{
			ChainSelector: chain.Selector,
			Address:       input.PriceRegistry,
			Args:          nil,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to execute PriceRegistryGetFeeTokenOp "+
				"on %s for price registry %s: %w", chain.String(), input.PriceRegistry.String(), err)
		}
		// check if fee tokens are already present in allTokens
		allTokenMap := make(map[common.Address]struct{})
		for _, token := range allTokens {
			allTokenMap[token] = struct{}{}
		}
		for _, token := range feetokensRep.Output {
			if _, exists := allTokenMap[token]; !exists {
				allTokenMap[token] = struct{}{}
				allTokens = append(allTokens, token)
			}
		}
		tokenPrices := make(map[common.Address]*big.Int)
		tokenPriceOutput, err := operations.ExecuteOperation(b, priceregistryops.PriceRegistryGetTokenPrices, chain, contract.FunctionInput[[]common.Address]{
			ChainSelector: chain.Selector,
			Address:       input.PriceRegistry,
			Args:          allTokens,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to execute PriceRegistryGetTokenPricesOp "+
				"on %s for price registry %s and tokens %v: %w", chain.String(), input.PriceRegistry.String(), allTokens, err)
		}
		for i, token := range allTokens {
			tokenPrices[token] = tokenPriceOutput.Output[i].Value
		}

		return sequences.OnChainOutput{
			Metadata: sequences.Metadata{
				Contracts: []datastore.ContractMetadata{
					{
						ChainSelector: input.ChainSelector,
						Address:       input.PriceRegistry.Hex(),
						Metadata: PriceRegistryImportConfigSequenceOutput{
							TokenPrices: tokenPrices,
							GasPrices:   gasPrices,
						},
					},
				},
			},
		}, nil
	},
)
