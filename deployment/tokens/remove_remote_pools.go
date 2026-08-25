package tokens

import (
	"errors"
	"fmt"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

// RemoveRemotePoolsInput is the input for the RemoveRemotePools changeset. It removes remote
// pool entries from existing token pools.
type RemoveRemotePoolsInput struct {
	Pools []RemoveRemotePoolsPerPool `yaml:"pools" json:"pools"`
	MCMS  mcms.Input                 `yaml:"mcms,omitempty" json:"mcms"`
}

// RemoveRemotePoolsPerPool groups remote pool removals for a single token pool.
type RemoveRemotePoolsPerPool struct {
	ChainSelector       uint64               `yaml:"selector" json:"selector,string"`
	Address             string               `yaml:"address" json:"address"`
	RemotePoolsToRemove []RemotePoolToRemove `yaml:"remotePoolsToRemove" json:"remotePoolsToRemove"`
}

// RemoveRemotePools returns a changeset that removes remote pool entries from existing token
// pools. The operation version is inferred from the token pool (via the datastore), so the
// top-level changeset does not require a version field.
func RemoveRemotePools() cldf.ChangeSetV2[RemoveRemotePoolsInput] {
	return cldf.CreateChangeSet(removeRemotePoolsApply(), removeRemotePoolsVerify())
}

func removeRemotePoolsVerify() func(cldf.Environment, RemoveRemotePoolsInput) error {
	return func(_ cldf.Environment, cfg RemoveRemotePoolsInput) error {
		if len(cfg.Pools) == 0 {
			return errors.New("input must contain at least one pool entry")
		}

		type poolKey struct {
			selector uint64
			address  string
		}

		seenPools := make(map[poolKey]struct{})
		for _, pool := range cfg.Pools {
			if _, err := chain_selectors.GetSelectorFamily(pool.ChainSelector); err != nil {
				return fmt.Errorf("invalid chain selector %d: %w", pool.ChainSelector, err)
			}

			if pool.Address == "" {
				return fmt.Errorf("pool entry on chain selector %d has an empty address", pool.ChainSelector)
			}

			if len(pool.RemotePoolsToRemove) == 0 {
				return fmt.Errorf("pool entry %s on chain selector %d has no remote pools to remove", pool.Address, pool.ChainSelector)
			}

			seenRemotes := make(map[uint64]struct{})
			for _, remote := range pool.RemotePoolsToRemove {
				if remote.Selector == pool.ChainSelector {
					return fmt.Errorf("remote chain selector %d must not equal the pool's own chain selector", remote.Selector)
				}

				if _, err := chain_selectors.GetSelectorFamily(remote.Selector); err != nil {
					return fmt.Errorf("invalid remote chain selector %d: %w", remote.Selector, err)
				}

				if remote.Address == "" {
					return fmt.Errorf("remote pool entry for chain selector %d has an empty address", remote.Selector)
				}

				if _, dup := seenRemotes[remote.Selector]; dup {
					return fmt.Errorf("duplicate remote chain selector %d for pool on chain selector %d", remote.Selector, pool.ChainSelector)
				}

				seenRemotes[remote.Selector] = struct{}{}
			}

			key := poolKey{selector: pool.ChainSelector, address: pool.Address}
			if _, dup := seenPools[key]; dup {
				return fmt.Errorf("duplicate pool entry for chain selector %d and address %s", pool.ChainSelector, pool.Address)
			}

			seenPools[key] = struct{}{}
		}
		return nil
	}
}

func removeRemotePoolsApply() func(cldf.Environment, RemoveRemotePoolsInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg RemoveRemotePoolsInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)
		tokenRegistry := GetTokenAdapterRegistry()
		mcmsRegistry := changesets.GetRegistry()

		for _, pool := range cfg.Pools {
			selector := pool.ChainSelector
			poolRef := datastore.AddressRef{ChainSelector: selector, Address: pool.Address}

			adapter, family, fullPoolRef, fullTokenRef, err := ResolveAdapterAndRefs(e, tokenRegistry, selector, poolRef, datastore.AddressRef{})
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to resolve pool %s on chain selector %d: %w", pool.Address, selector, err)
			}

			remover, ok := adapter.(RemotePoolRemover)
			if !ok {
				return cldf.ChangesetOutput{}, fmt.Errorf(
					"adapter for chain selector %d (family %s, version %s) does not support remote pool removal",
					selector, family, fullPoolRef.Version,
				)
			}

			report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, remover.RemoveRemotePool(), e.BlockChains, RemoveRemotePoolSequenceInput{
				Selector:            selector,
				TokenPoolRef:        fullPoolRef,
				TokenRef:            fullTokenRef,
				RemotePoolsToRemove: pool.RemotePoolsToRemove,
			})
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to remove remote pools from pool %s on chain selector %d: %w", pool.Address, selector, err)
			}

			batchOps = append(batchOps, report.Output.BatchOps...)
			reports = append(reports, report.ExecutionReports...)
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}
}
