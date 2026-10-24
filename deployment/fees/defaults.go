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
	srcFam, srcFamErr := chainsel.GetSelectorFamily(src)
	dstFam, dstFamErr := chainsel.GetSelectorFamily(dst)

	isEth := func(sel uint64) bool {
		return sel == chainsel.ETHEREUM_TESTNET_SEPOLIA.Selector ||
			sel == chainsel.ETHEREUM_TESTNET_HOODI.Selector ||
			sel == chainsel.ETHEREUM_MAINNET.Selector
	}

	minFeeUSDCents := uint32(25)
	if src != dst && srcFamErr == nil && dstFamErr == nil {
		switch {
		case srcFam == chainsel.FamilyEVM && dstFam == chainsel.FamilySolana:
			if isEth(src) {
				minFeeUSDCents = 60
			} else {
				minFeeUSDCents = 35
			}
		case srcFam == chainsel.FamilySolana && dstFam == chainsel.FamilyEVM:
			if isEth(dst) {
				minFeeUSDCents = 150
			}
		case srcFam == chainsel.FamilyEVM && dstFam == chainsel.FamilyEVM:
			if isEth(src) {
				minFeeUSDCents = 50
			}
			if isEth(dst) {
				minFeeUSDCents = 150
			}
		case dstFam == chainsel.FamilySolana:
			minFeeUSDCents = 35
		case dstFam == chainsel.FamilyEVM:
			if isEth(dst) {
				minFeeUSDCents = 150
			} else {
				minFeeUSDCents = 25
			}
		}
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

// ResolveOnRampAndFeeQuoter resolves the OnRamp adapter, the OnRamp ref, the FeeQuoter adapter and the
// FeeQuoter ref for a lane. The OnRamp and FeeQuoter versions are not guaranteed to be the same (mixed
// lanes exist), so both adapters are returned: use the FeeQuoter adapter to read/write the fee config,
// and the OnRamp adapter to obtain the OnRamp-version-appropriate default.
func ResolveOnRampAndFeeQuoter(b cldf_ops.Bundle, chains cldf_chain.BlockChains, ds datastore.DataStore, src, dst uint64) (FeeAdapter, datastore.AddressRef, FeeAdapter, datastore.AddressRef, error) {
	registry := GetRegistry()

	fam, err := chainsel.GetSelectorFamily(src)
	if err != nil {
		return nil, datastore.AddressRef{}, nil, datastore.AddressRef{}, fmt.Errorf("failed to get chain selector family for selector %d: %w", src, err)
	}

	// Fee Quoter resolution part 1: get the current on ramp from the router
	resolver, ok := registry.GetFeeResolver(fam)
	if !ok {
		return nil, datastore.AddressRef{}, nil, datastore.AddressRef{}, fmt.Errorf("no fee resolver found for chain selector %d", src)
	}
	onRampRef, err := resolver.GetOnRampRef(b, chains, ds, src, dst)
	if err != nil {
		return nil, datastore.AddressRef{}, nil, datastore.AddressRef{}, fmt.Errorf("failed to resolve on-ramp ref for chain selector %d and remote chain selector %d: %w", src, dst, err)
	}

	// Fee Quoter resolution part 2: get the current fee quoter from the on ramp
	onRampAdp, ok := registry.GetFeeAdapter(fam, onRampRef.Version)
	if !ok {
		return nil, datastore.AddressRef{}, nil, datastore.AddressRef{}, fmt.Errorf("no fee adapter found for chain selector %d, version %s", src, onRampRef.Version)
	}
	feeRef, err := onRampAdp.GetFeeContractRef(b, chains, ds, onRampRef, src, dst)
	if err != nil {
		return nil, datastore.AddressRef{}, nil, datastore.AddressRef{}, fmt.Errorf("failed to get fee contract ref for chain selector %d and remote chain selector %d: %w", src, dst, err)
	}

	// Fee Quoter resolution part 3: lookup the adapter using the fee quoter version
	feeQuoterAdp, ok := registry.GetFeeAdapter(fam, feeRef.Version)
	if !ok {
		return nil, datastore.AddressRef{}, nil, datastore.AddressRef{}, fmt.Errorf("no fee adapter found for chain selector %d and fee quoter version %s", src, feeRef.Version)
	}

	return onRampAdp, onRampRef, feeQuoterAdp, feeRef, nil
}

func ResolveFeeAdapter(b cldf_ops.Bundle, chains cldf_chain.BlockChains, ds datastore.DataStore, src, dst uint64) (FeeAdapter, datastore.AddressRef, error) {
	_, _, feeAdapter, feeRef, err := ResolveOnRampAndFeeQuoter(b, chains, ds, src, dst)
	return feeAdapter, feeRef, err
}
