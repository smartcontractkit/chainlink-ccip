package fees

import (
	"fmt"
	"math"

	chainsel "github.com/smartcontractkit/chain-selectors"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

func GetDefaultChainAgnosticTokenTransferFeeConfig(src uint64, dst uint64, overrides ...func(*TokenTransferFeeArgs)) TokenTransferFeeArgs {
	minFeeUSDCents := uint32(25)

	// NOTE: we validate that src != dst so only one of these if statements will execute
	if src == chainsel.ETHEREUM_MAINNET.Selector {
		minFeeUSDCents = 50
	}
	if dst == chainsel.ETHEREUM_MAINNET.Selector {
		minFeeUSDCents = 150
	}

	cfg := TokenTransferFeeArgs{
		DestBytesOverhead: 32,
		DestGasOverhead:   90_000,
		MinFeeUSDCents:    minFeeUSDCents,
		MaxFeeUSDCents:    math.MaxUint32,
		DeciBps:           0,
		IsEnabled:         true,
	}

	for _, override := range overrides {
		override(&cfg)
	}

	return cfg
}

func ResolveFeeAdapter(b cldf_ops.Bundle, chains cldf_chain.BlockChains, ds datastore.DataStore, src, dst uint64) (FeeAdapter, datastore.AddressRef, error) {
	// NOTE: the OnRamp version and FeeQuoter versions are not guaranteed to be the same.
	// We need to do a multi-step lookup in order to get the fee quoter that's currently
	// configured on-chain.
	registry := GetRegistry()

	fam, err := chainsel.GetSelectorFamily(src)
	if err != nil {
		return nil, datastore.AddressRef{}, fmt.Errorf("failed to get chain selector family for selector %d: %w", src, err)
	}

	// Fee Quoter resolution part 1: get the current on ramp from the router
	resolver, ok := registry.GetFeeResolver(fam)
	if !ok {
		return nil, datastore.AddressRef{}, fmt.Errorf("no fee resolver found for chain selector %d", src)
	}
	onRampRef, err := resolver.GetOnRampRef(b, chains, ds, src, dst)
	if err != nil {
		return nil, datastore.AddressRef{}, fmt.Errorf("failed to resolve on-ramp ref for chain selector %d and remote chain selector %d: %w", src, dst, err)
	}

	// Fee Quoter resolution part 2: get the current fee quoter from the on ramp
	onRampAdp, ok := registry.GetFeeAdapter(fam, onRampRef.Version)
	if !ok {
		return nil, datastore.AddressRef{}, fmt.Errorf("no fee adapter found for chain selector %d, version %s", src, onRampRef.Version)
	}
	feeRef, err := onRampAdp.GetFeeContractRef(b, chains, ds, onRampRef, src, dst)
	if err != nil {
		return nil, datastore.AddressRef{}, fmt.Errorf("failed to get fee contract ref for chain selector %d and remote chain selector %d: %w", src, dst, err)
	}

	// Fee Quoter resolution part 3: lookup the adapter using the fee quoter version
	feeQuoterAdp, ok := registry.GetFeeAdapter(fam, feeRef.Version)
	if !ok {
		return nil, datastore.AddressRef{}, fmt.Errorf("no fee adapter found for chain selector %d and fee quoter version %s", src, feeRef.Version)
	}

	return feeQuoterAdp, feeRef, nil
}
