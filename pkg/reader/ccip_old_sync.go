package reader

import (
	"context"
	"fmt"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	typeconv "github.com/smartcontractkit/chainlink-ccip/internal/libs/typeconv"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
)

func (r *ccipChainReader) bindOfframp(ctx context.Context) error {
	if err := r.validateReaderExistence(r.destChain); err != nil {
		return err
	}

	// Bind the offRamp contract address to the reader.
	// If the same address exists -> no-op
	// If the address is changed -> updates the address, overwrites the existing one
	// If the contract not binded -> binds to the new address
	if err := r.contractReaders[r.destChain].Bind(ctx, []types.BoundContract{
		{
			Address: r.offrampAddress,
			Name:    consts.ContractNameOffRamp,
		},
	}); err != nil {
		return fmt.Errorf("bind offRamp: %w", err)
	}

	return nil
}

// bindOnRamps reads the onchain configuration to discover source ramp addresses.
func (r *ccipChainReader) bindOnramps(
	ctx context.Context,
) error {
	chains := make([]cciptypes.ChainSelector, 0, len(r.contractReaders))
	for chain := range r.contractReaders {
		chains = append(chains, chain)
	}
	sourceConfigs, err := r.getSourceChainsConfig(ctx, chains)
	if err != nil {
		return fmt.Errorf("get onramps: %w", err)
	}

	r.lggr.Infow("got source chain configs", "onramps", func() []string {
		var r []string
		for chainSelector, scc := range sourceConfigs {
			r = append(r, typeconv.AddressBytesToString(scc.OnRamp, uint64(chainSelector)))
		}
		return r
	}())

	for chain, cfg := range sourceConfigs {
		if len(cfg.OnRamp) == 0 {
			return fmt.Errorf("onRamp address not found for chain %d", chain)
		}

		// We only want to produce reports for enabled source chains.
		if !cfg.IsEnabled {
			continue
		}

		// Bind the onRamp contract address to the reader.
		// If the same address exists -> no-op
		// If the address is changed -> updates the address, overwrites the existing one
		// If the contract not binded -> binds to the new address
		if err := r.contractReaders[chain].Bind(ctx, []types.BoundContract{
			{
				Address: typeconv.AddressBytesToString(cfg.OnRamp, uint64(chain)),
				Name:    consts.ContractNameOnRamp,
			},
		}); err != nil {
			return fmt.Errorf("bind onRamp: %w", err)
		}
	}

	return nil
}

func (r *ccipChainReader) bindNonceManager(ctx context.Context) error {
	destReader, ok := r.contractReaders[r.destChain]
	if !ok {
		r.lggr.Debugw("skipping nonce manager, dest chain not configured for this deployment",
			"destChain", r.destChain)
		return nil
	}

	staticConfig, err := r.getOfframpStaticConfig(ctx, r.destChain)
	if err != nil {
		return fmt.Errorf("get offramp static config: %w", err)
	}

	if staticConfig.ChainSelector != r.destChain {
		return fmt.Errorf("invalid configuration detected, somehow reading from non-dest offramp")
	}

	// Bind the nonceManager contract address to the reader.
	// If the same address exists -> no-op
	// If the address is changed -> updates the address, overwrites the existing one
	// If the contract not binded -> binds to the new address
	if err := destReader.Bind(ctx, []types.BoundContract{
		{
			Address: typeconv.AddressBytesToString(staticConfig.NonceManager, uint64(r.destChain)),
			Name:    consts.ContractNameNonceManager,
		},
	}); err != nil {
		return fmt.Errorf("bind nonce manager: %w", err)
	}

	return nil
}

func (r *ccipChainReader) Sync(ctx context.Context, _ ContractAddresses) error {
	// Note: offramp must be bound first, otherwise onramp bindings
	// will fail.
	if err := r.bindOfframp(ctx); err != nil {
		return err
	}

	if err := r.bindOnramps(ctx); err != nil {
		return err
	}

	return r.bindNonceManager(ctx)
}
