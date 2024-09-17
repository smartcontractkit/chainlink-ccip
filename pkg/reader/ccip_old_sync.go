package reader

import (
	"context"
	"fmt"
)

func (r *ccipChainReader) Sync(ctx context.Context, _ ContractAddresses) error {
	/*
		// Note: offramp must be bound first, otherwise onramp bindings
		// will fail.
		if err := r.bindOfframp(ctx); err != nil {
			return err
		}

		if err := r.bindOnramps(ctx); err != nil {
			return err
		}

		return r.bindNonceManager(ctx)
	*/
	contracts, err := r.DiscoverContracts(ctx, r.destChain)
	if err != nil {
		return fmt.Errorf("sync: %w", err)
	}

	return r.newSync(ctx, contracts)
}
