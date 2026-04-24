package hooks

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/changeset"
	cfgnet "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/config/network"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/domain"
	"golang.org/x/sync/errgroup"
)

const (
	requireOwnedEnvContractsHookName = "require-owned-env-contracts"
)

// Hook name constants for changeset wiring and chain-specific tests.
const (
	RequireOwnedEnvContractsHookName = requireOwnedEnvContractsHookName
)

var (
	singletonContractOwnershipRegistry *ContractOwnershipRegistry
	onceContractOwnershipRegistry      sync.Once
)

// ContractOwnership checks ownership for contracts in a specific chain family.
type ContractOwnership interface {
	FilterNetworks(envName string, dom domain.Domain, lggr logger.Logger) (*cfgnet.Config, error)
	NeedsOwnershipCheck(ref datastore.AddressRef) bool
	VerifyContractOwnership(
		ctx context.Context,
		lggr logger.Logger,
		network cfgnet.Network,
		refsToCheck []datastore.AddressRef,
	) error
}

type ContractOwnershipRegistry struct {
	providers map[string]ContractOwnership
	mu        *sync.Mutex
}

func newContractOwnershipRegistry() *ContractOwnershipRegistry {
	return &ContractOwnershipRegistry{
		providers: make(map[string]ContractOwnership),
		mu:        &sync.Mutex{},
	}
}

func GetContractOwnershipRegistry() *ContractOwnershipRegistry {
	onceContractOwnershipRegistry.Do(func() {
		singletonContractOwnershipRegistry = newContractOwnershipRegistry()
	})
	return singletonContractOwnershipRegistry
}

func (r *ContractOwnershipRegistry) Register(family string, provider ContractOwnership) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.providers[family]; !exists {
		r.providers[family] = provider
	}
}

func (r *ContractOwnershipRegistry) Get(family string) (ContractOwnership, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	provider, ok := r.providers[family]
	return provider, ok
}

// NewRequireOwnedEnvContractsPreHook returns a pre-apply hook that checks ownership
// of contracts in the environment datastore before changeset execution.
func NewRequireOwnedEnvContractsPreHook(dom domain.Domain, verifier ContractOwnership) changeset.PreHook {
	return requireOwnedEnvContractsPreHook(dom, verifier)
}

// RequireOwnedEnvContractsPreHookForMultipleChainFamilies returns one pre-hook per chain family.
// If duplicate families are provided, only one hook is returned for that family.
func RequireOwnedEnvContractsPreHookForMultipleChainFamilies(dom domain.Domain, chainFamilies []string) []changeset.PreHook {
	preHooksByFamily := make(map[string]changeset.PreHook)
	for _, family := range chainFamilies {
		if _, exists := preHooksByFamily[family]; exists {
			continue
		}
		verifier, ok := GetContractOwnershipRegistry().Get(family)
		if !ok {
			panic(fmt.Sprintf("no contract ownership provider registered for chain family %s", family))
		}
		preHooksByFamily[family] = requireOwnedEnvContractsPreHook(dom, verifier)
	}

	var preHooks []changeset.PreHook
	for _, hook := range preHooksByFamily {
		preHooks = append(preHooks, hook)
	}
	return preHooks
}

func requireOwnedEnvContractsPreHook(
	dom domain.Domain,
	verifier ContractOwnership,
) changeset.PreHook {
	return changeset.PreHook{
		HookDefinition: changeset.HookDefinition{
			Name:          requireOwnedEnvContractsHookName,
			FailurePolicy: changeset.Abort,
			Timeout:       15 * time.Minute,
		},
		Func: requireOwnedEnvContracts(dom, verifier),
	}
}

func requireOwnedEnvContracts(
	dom domain.Domain,
	verifier ContractOwnership,
) changeset.PreHookFunc {
	return func(ctx context.Context, params changeset.PreHookParams) error {
		ds, err := dom.EnvDir(params.Env.Name).DataStore()
		if err != nil {
			return fmt.Errorf("require ownership pre-hook: load datastore: %w", err)
		}

		networkCfg, err := verifier.FilterNetworks(params.Env.Name, dom, params.Env.Logger)
		if err != nil {
			return fmt.Errorf("require ownership pre-hook: load networks: %w", err)
		}

		return IterateOwnershipCheckers(ctx, params.Env.Logger, ds, networkCfg, verifier)
	}
}

// IterateOwnershipCheckers walks configured networks and runs family-specific ownership checks.
func IterateOwnershipCheckers(
	ctx context.Context,
	lggr logger.Logger,
	ds datastore.DataStore,
	networkCfg *cfgnet.Config,
	verifier ContractOwnership,
) error {
	var errs []error
	var errsMu sync.Mutex
	networkGrp, ctx := errgroup.WithContext(ctx)
	networkGrp.SetLimit(concurrentNetworksLimit)
	for _, network := range networkCfg.Networks() {
		network := network
		networkGrp.Go(func() error {
			// TODO: filter this further when chains are available in hookEnv
			chainRefs := ds.Addresses().Filter(datastore.AddressRefByChainSelector(network.ChainSelector))
			if len(chainRefs) == 0 {
				lggr.Infof("no address refs for network %d in datastore, skipping ownership check", network.ChainSelector)
				return nil
			}
			refsToCheck := make([]datastore.AddressRef, 0, len(chainRefs))
			for _, ref := range chainRefs {
				if ref.Type == "" || ref.Version == nil || ref.Address == "" {
					errsMu.Lock()
					errs = append(errs, fmt.Errorf("ownership check err: invalid address ref %+v, missing type, version, or address", ref))
					errsMu.Unlock()
					continue
				}
				if !verifier.NeedsOwnershipCheck(ref) {
					lggr.Infof("skipping ownership check for %s %s (%s on %d) based on NeedsOwnershipCheck criteria", ref.Type, ref.Version, ref.Address, network.ChainSelector)
					continue
				}
				refsToCheck = append(refsToCheck, ref)
			}
			if len(refsToCheck) == 0 {
				lggr.Infof("no ownership checks required for network %d after filtering", network.ChainSelector)
				return nil
			}
			if err := verifier.VerifyContractOwnership(ctx, lggr, network, refsToCheck); err != nil {
				errsMu.Lock()
				errs = append(errs, fmt.Errorf("ownership check err: network %d: %w", network.ChainSelector, err))
				errsMu.Unlock()
			}
			return nil
		})
	}
	if err := networkGrp.Wait(); err != nil {
		errsMu.Lock()
		errs = append(errs, fmt.Errorf("ownership check err: wait for networks: %w", err))
		errsMu.Unlock()
	}
	return errors.Join(errs...)
}

// ResetContractOwnershipRegistryForTest replaces the registry and resets sync.Once so tests
// can register providers in isolation.
func ResetContractOwnershipRegistryForTest() {
	singletonContractOwnershipRegistry = newContractOwnershipRegistry()
	onceContractOwnershipRegistry = sync.Once{}
}
