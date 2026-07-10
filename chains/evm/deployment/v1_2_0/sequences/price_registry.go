package sequences

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	evmops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations"
	priceregistryops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/price_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/price_registry"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

type PriceRegistryImportConfigSequenceInput struct {
	ChainSelector uint64
	PriceRegistry common.Address
	RemoteChains  []uint64
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

		gasPrices := make(map[uint64]*big.Int)
		for _, remoteChainSelector := range input.RemoteChains {
			gasPricesOutput, err := evmops.ExecuteRead(b, chain, input.PriceRegistry, evmops.BindAs[price_registry.PriceRegistryInterface](price_registry.NewPriceRegistry), priceregistryops.NewReadGetDestinationChainGasPrice, remoteChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute PriceRegistryGetDestinationChainGasPriceOp "+
					"on %s for price registry %s and remote chain %d: %w", chain.String(), input.PriceRegistry.String(), remoteChainSelector, err)
			}
			gasPrices[remoteChainSelector] = gasPricesOutput.Output.Value
		}
		feetokensRep, err := evmops.ExecuteRead(b, chain, input.PriceRegistry, evmops.BindAs[price_registry.PriceRegistryInterface](price_registry.NewPriceRegistry), priceregistryops.NewReadGetFeeToken, struct{}{})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to execute PriceRegistryGetFeeTokenOp "+
				"on %s for price registry %s: %w", chain.String(), input.PriceRegistry.String(), err)
		}

		tokenPrices := make(map[common.Address]*big.Int)
		tokenPriceOutput, err := evmops.ExecuteRead(b, chain, input.PriceRegistry, evmops.BindAs[price_registry.PriceRegistryInterface](price_registry.NewPriceRegistry), priceregistryops.NewReadGetTokenPrices, feetokensRep.Output)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to execute PriceRegistryGetTokenPricesOp "+
				"on %s for price registry %s and fee tokens %v: %w", chain.String(), input.PriceRegistry.String(), feetokensRep.Output, err)
		}
		for i, token := range feetokensRep.Output {
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
