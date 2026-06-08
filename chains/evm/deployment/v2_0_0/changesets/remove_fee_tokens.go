package changesets

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	adapters1_2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/adapters"
	priceregistryops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/price_registry"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	seq1_6 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	fq1_6ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_3/operations/fee_quoter"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

// RemoveFeeTokensCfg is configuration for the RemoveFeeTokens changeset.
type RemoveFeeTokensCfg struct {
	ChainSels []uint64
}

var RemoveFeeTokens = changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[
	sequences.RemoveFeeTokensInput,
	cldf_chain.BlockChains,
	RemoveFeeTokensCfg,
]{
	Sequence: sequences.SequenceRemoveFeeTokens,
	ResolveInput: func(e cldf_deployment.Environment, cfg RemoveFeeTokensCfg) (sequences.RemoveFeeTokensInput, error) {
		if len(cfg.ChainSels) == 0 {
			return sequences.RemoveFeeTokensInput{}, fmt.Errorf("at least one chain selector is required")
		}

		seen := make(map[uint64]struct{}, len(cfg.ChainSels))
		chainUpdates := make([]sequences.RemoveFeeTokensPerChainInput, 0, len(cfg.ChainSels))
		for _, chainSel := range cfg.ChainSels {
			if _, exists := seen[chainSel]; exists {
				return sequences.RemoveFeeTokensInput{}, fmt.Errorf("duplicate chain selector %d", chainSel)
			}
			seen[chainSel] = struct{}{}

			chainUpdate, err := resolveRemoveFeeTokensPerChain(e, chainSel)
			if err != nil {
				return sequences.RemoveFeeTokensInput{}, err
			}
			chainUpdates = append(chainUpdates, chainUpdate)
		}

		return sequences.RemoveFeeTokensInput{ChainUpdates: chainUpdates}, nil
	},
	ResolveDep: func(e cldf_deployment.Environment, _ RemoveFeeTokensCfg) (cldf_chain.BlockChains, error) {
		return e.BlockChains, nil
	},
})

func resolveRemoveFeeTokensPerChain(e cldf_deployment.Environment, chainSel uint64) (sequences.RemoveFeeTokensPerChainInput, error) {
	addresses := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(chainSel))

	routerRefs := e.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(chainSel),
		datastore.AddressRefByType(datastore.ContractType(routerops.ContractType)),
		datastore.AddressRefByVersion(routerops.Version),
	)
	if len(routerRefs) != 1 {
		return sequences.RemoveFeeTokensPerChainInput{}, fmt.Errorf(
			"expected exactly one Router v%s on chain %d, found %d",
			routerops.Version, chainSel, len(routerRefs),
		)
	}

	laneResolver := &adapters1_2.LaneVersionResolver{}
	_, laneVersions, err := laneResolver.DeriveLaneVersionsForChain(e, chainSel)
	if err != nil {
		return sequences.RemoveFeeTokensPerChainInput{}, fmt.Errorf(
			"failed to derive lane versions for chain %d from router %s: %w",
			chainSel, routerRefs[0].Address, err,
		)
	}

	chain, ok := e.BlockChains.EVMChains()[chainSel]
	if !ok {
		return sequences.RemoveFeeTokensPerChainInput{}, fmt.Errorf("chain selector %d not found in environment EVM chains", chainSel)
	}

	legacyFeeTokens, err := collectLegacyFeeTokens(e, chain, addresses, laneVersions)
	if err != nil {
		return sequences.RemoveFeeTokensPerChainInput{}, err
	}

	fq20Ref := datastore_utils.GetAddressRef(
		addresses,
		chainSel,
		fqops.ContractType,
		fqops.Version,
		"",
	)
	if datastore_utils.IsAddressRefEmpty(fq20Ref) {
		return sequences.RemoveFeeTokensPerChainInput{}, fmt.Errorf("no FeeQuoter v%s found on chain selector %d", fqops.Version, chainSel)
	}

	fq20Tokens, err := queryFeeQuoter20FeeTokens(e, chain, common.HexToAddress(fq20Ref.Address))
	if err != nil {
		return sequences.RemoveFeeTokensPerChainInput{}, err
	}

	return sequences.RemoveFeeTokensPerChainInput{
		ChainSelector:     chainSel,
		FeeQuoter20Ref:    fq20Ref,
		FeeTokensToRemove: extraFeeTokens(fq20Tokens, legacyFeeTokens),
	}, nil
}

func collectLegacyFeeTokens(
	e cldf_deployment.Environment,
	chain evm.Chain,
	addresses []datastore.AddressRef,
	laneVersions []*semver.Version,
) ([]common.Address, error) {
	var hasV15, hasV16 bool
	for _, version := range laneVersions {
		if version.Major() == 1 && version.Minor() == 5 {
			hasV15 = true
		}
		if version.Major() == 1 && version.Minor() == 6 {
			hasV16 = true
		}
	}

	legacySet := make(map[common.Address]struct{})

	if hasV15 {
		priceRegistryRef := datastore_utils.GetAddressRef(
			addresses,
			chain.Selector,
			priceregistryops.ContractType,
			priceregistryops.Version,
			"",
		)
		if datastore_utils.IsAddressRefEmpty(priceRegistryRef) {
			return nil, fmt.Errorf("no PriceRegistry v%s found on chain selector %d for 1.5 lanes",
				priceregistryops.Version, chain.Selector)
		}

		feeTokens, err := queryPriceRegistryFeeTokens(e, chain, common.HexToAddress(priceRegistryRef.Address))
		if err != nil {
			return nil, err
		}
		for _, token := range feeTokens {
			legacySet[token] = struct{}{}
		}
	}

	if hasV16 {
		fq16Ref, err := seq1_6.GetFeeQuoterAddress(addresses, chain.Selector, fqops.Version)
		if err != nil {
			return nil, fmt.Errorf("no FeeQuoter v1.6.x found on chain selector %d for 1.6 lanes: %w", chain.Selector, err)
		}

		feeTokens, err := queryFeeQuoter16FeeTokens(e, chain, common.HexToAddress(fq16Ref.Address))
		if err != nil {
			return nil, err
		}
		for _, token := range feeTokens {
			legacySet[token] = struct{}{}
		}
	}

	legacyFeeTokens := make([]common.Address, 0, len(legacySet))
	for token := range legacySet {
		legacyFeeTokens = append(legacyFeeTokens, token)
	}
	return legacyFeeTokens, nil
}

func queryPriceRegistryFeeTokens(e cldf_deployment.Environment, chain evm.Chain, priceRegistry common.Address) ([]common.Address, error) {
	report, err := cldf_ops.ExecuteOperation(e.OperationsBundle, priceregistryops.PriceRegistryGetFeeToken, chain, contract.FunctionInput[any]{
		ChainSelector: chain.Selector,
		Address:       priceRegistry,
		Args:          nil,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get fee tokens from PriceRegistry %s on chain %d: %w",
			priceRegistry.Hex(), chain.Selector, err)
	}
	return report.Output, nil
}

func queryFeeQuoter16FeeTokens(e cldf_deployment.Environment, chain evm.Chain, feeQuoter common.Address) ([]common.Address, error) {
	report, err := cldf_ops.ExecuteOperation(e.OperationsBundle, fq1_6ops.GetFeeTokens, chain, contract.FunctionInput[struct{}]{
		ChainSelector: chain.Selector,
		Address:       feeQuoter,
		Args:          struct{}{},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get fee tokens from FeeQuoter 1.6 %s on chain %d: %w",
			feeQuoter.Hex(), chain.Selector, err)
	}
	return report.Output, nil
}

func queryFeeQuoter20FeeTokens(e cldf_deployment.Environment, chain evm.Chain, feeQuoter common.Address) ([]common.Address, error) {
	report, err := cldf_ops.ExecuteOperation(e.OperationsBundle, fqops.GetFeeTokens, chain, contract.FunctionInput[struct{}]{
		ChainSelector: chain.Selector,
		Address:       feeQuoter,
		Args:          struct{}{},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get fee tokens from FeeQuoter 2.0 %s on chain %d: %w",
			feeQuoter.Hex(), chain.Selector, err)
	}
	return report.Output, nil
}

func extraFeeTokens(fq20Tokens, legacyFeeTokens []common.Address) []common.Address {
	legacySet := make(map[common.Address]struct{}, len(legacyFeeTokens))
	for _, token := range legacyFeeTokens {
		legacySet[token] = struct{}{}
	}

	extra := make([]common.Address, 0)
	for _, token := range fq20Tokens {
		if _, exists := legacySet[token]; !exists {
			extra = append(extra, token)
		}
	}
	return extra
}
