package hooks

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/changeset"
	cfgnet "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/config/network"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/domain"
	cldverification "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/verification"
	"golang.org/x/sync/errgroup"
)

const (
	verifyDeployedContractsHookName     = "verify-deployed-contracts"
	requireVerifiedEnvContractsHookName = "require-verified-env-contracts"
	DefaultVerifyPollInterval           = 1 * time.Second
	concurrentNetworksLimit             = 5
	concurrentVerificationsLimit        = 5
)

// Hook name constants for changeset wiring and chain-specific tests.
const (
	VerifyDeployedContractsHookName     = verifyDeployedContractsHookName
	RequireVerifiedEnvContractsHookName = requireVerifiedEnvContractsHookName
)

var (
	singletonContractVerificationRegistry *ContractVerificationRegistry
	onceContractVerificationRegistry      sync.Once
)

type VerifierBuilderForNetwork func(ctx context.Context, ref datastore.AddressRef) (v cldverification.Verifiable, err error)

type ContractVerification interface {
	FilterNetworks(envName string, dom domain.Domain, lggr logger.Logger) (*cfgnet.Config, error)
	NeedsVerification(ref datastore.AddressRef) bool
	ForEachNetwork(ctx context.Context, network cfgnet.Network, selector uint64, lggr logger.Logger, logPrefix string) (VerifierBuilderForNetwork, bool)
}

type ContractVerificationRegistry struct {
	providers map[string]ContractVerification
	mu        *sync.Mutex
}

func newContractVerificationRegistry() *ContractVerificationRegistry {
	return &ContractVerificationRegistry{
		providers: make(map[string]ContractVerification),
		mu:        &sync.Mutex{},
	}
}

func GetContractVerificationRegistry() *ContractVerificationRegistry {
	onceContractVerificationRegistry.Do(func() {
		singletonContractVerificationRegistry = newContractVerificationRegistry()
	})
	return singletonContractVerificationRegistry
}

func (r *ContractVerificationRegistry) Register(family string, provider ContractVerification) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.providers[family]; !exists {
		r.providers[family] = provider
	}
}

func (r *ContractVerificationRegistry) Get(family string) (ContractVerification, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	provider, ok := r.providers[family]
	return provider, ok
}

// VerifyDeployedContractsPostHookForMultipleChainFamilies returns a slice of post-apply hooks that verifies deployed contracts using chain-family verifiers (e.g. EVM).
// If multiple chain selectors from the same family are provided, only one hook will be returned for that family to avoid redundant verification.
// The hooks will verify all contracts in the datastore created by changeset output
// The addresses should match the criteria defined by the verifier's NeedsVerification and ForEachNetwork methods, for each network supported by the verifier's FilterNetworks method.
func VerifyDeployedContractsPostHookForMultipleChainFamilies(dom domain.Domain, chainFamilies []string) []changeset.PostHook {
	postHooksByFamily := make(map[string]changeset.PostHook)
	for _, family := range chainFamilies {
		if !isSelectorFamilySupported(family) {
			panic(fmt.Sprintf("unsupported chain family %s for contract verification hook", family))
		}
		_, exists := postHooksByFamily[family]
		if !exists {
			verifier, ok := GetContractVerificationRegistry().Get(family)
			if !ok {
				panic(fmt.Sprintf("no contract verification provider registered for chain family %s", family))
			}

			postHook := verifyDeployedContractsPostHook(dom, verifier)
			postHooksByFamily[family] = postHook
		}
	}
	var postHooks []changeset.PostHook
	for _, hook := range postHooksByFamily {
		postHooks = append(postHooks, hook)
	}
	return postHooks
}

// NewVerifyDeployedContractsPostHook returns a post-apply hook that verifies deployed contracts using a single chain-family verifier (e.g. EVM).
func NewVerifyDeployedContractsPostHook(dom domain.Domain, verifier ContractVerification) changeset.PostHook {
	return verifyDeployedContractsPostHook(dom, verifier)
}

// NewRequireVerifiedEnvContractsPreHook returns a pre-apply hook that requires contracts to already be verified on block explorers.
func NewRequireVerifiedEnvContractsPreHook(dom domain.Domain, verifier ContractVerification, refsToVerify []datastore.AddressRef, selectors []uint64) changeset.PreHook {
	return requireVerifiedEnvContractsPreHook(dom, verifier, refsToVerify, selectors)
}

func RequireVerifiedEnvContractsPreHookForMultipleChainFamilies(dom domain.Domain, chainSelectors []uint64, refsToVerify []datastore.AddressRef) []changeset.PreHook {
	preHooksByFamily := make(map[string]changeset.PreHook)
	chainselectorsByFamily := make(map[string][]uint64)
	// group chain selectors by family to avoid redundant pre-hooks for multiple chains of the same family
	for _, selector := range chainSelectors {
		family, err := chain_selectors.GetSelectorFamily(selector)
		if err != nil {
			panic(fmt.Sprintf("invalid chain selector %d: %v", selector, err))
		}
		chainselectorsByFamily[family] = append(chainselectorsByFamily[family], selector)
	}
	for family, selectors := range chainselectorsByFamily {
		_, exists := preHooksByFamily[family]
		if !exists {
			verifier, ok := GetContractVerificationRegistry().Get(family)
			if !ok {
				panic(fmt.Sprintf("no contract verification provider registered for chain family %s", family))
			}

			preHook := requireVerifiedEnvContractsPreHook(dom, verifier, refsToVerify, selectors)
			preHooksByFamily[family] = preHook
		}
	}
	var preHooks []changeset.PreHook
	for _, hook := range preHooksByFamily {
		preHooks = append(preHooks, hook)
	}
	return preHooks
}

func verifyDeployedContractsPostHook(
	dom domain.Domain,
	verifier ContractVerification,
) changeset.PostHook {
	return changeset.PostHook{
		HookDefinition: changeset.HookDefinition{
			Name:          verifyDeployedContractsHookName,
			FailurePolicy: changeset.Abort,
			Timeout:       15 * time.Minute,
		},
		Func: verifyDeployedContracts(dom, verifier),
	}
}

func requireVerifiedEnvContractsPreHook(
	dom domain.Domain,
	verifier ContractVerification,
	refsToVerify []datastore.AddressRef,
	selectors []uint64,
) changeset.PreHook {
	return changeset.PreHook{
		HookDefinition: changeset.HookDefinition{
			Name:          requireVerifiedEnvContractsHookName,
			FailurePolicy: changeset.Abort,
			Timeout:       5 * time.Minute,
		},
		Func: requireVerifiedEnvContracts(dom, verifier, refsToVerify, selectors),
	}
}

func verifyDeployedContracts(
	dom domain.Domain,
	verifier ContractVerification,
) changeset.PostHookFunc {
	return func(ctx context.Context, params changeset.PostHookParams) error {
		if params.Err != nil {
			// Skip verification when apply failed; returning the error would log a misleading post-hook failure.
			return nil //nolint:nilerr // apply error is already returned by the registry
		}
		if params.Output.DataStore == nil {
			return nil
		}
		// Get and filter networks for the provider's chain family; if this fails, return an error to log the failure and skip verification.
		networkCfg, err := verifier.FilterNetworks(params.Env.Name, dom, params.Env.Logger)
		if err != nil {
			return fmt.Errorf("verify hook: load networks: %w", err)
		}
		// Ensure all contracts that are added as part of changeset execution are verified
		// Note : certain contract types and versions may be filtered out by the verifier based on NeedsVerification and ForEachNetwork implementations,
		// For example, if a certain network is not supported by the verifier, ForEachNetwork can return skipNetwork=true to skip all addresses for that network.
		// Or if a certain contract type/version does not require verification, NeedsVerification can return false to skip that address.
		// so not all addresses in the datastore are guaranteed to be verified by this hook
		ds := params.Output.DataStore.Seal()
		return IterateVerifiers(ctx, ds, networkCfg, params.Env.Logger, "verify hook", verifier,
			func(ctx context.Context, v cldverification.Verifiable, ref datastore.AddressRef, selector uint64) error {
				// check if already verified before attempting verification,
				verified, err := v.IsVerified(ctx)
				if err != nil {
					return fmt.Errorf("verify hook: error checking verification status for %s %s (%s on %d): %w", ref.Type, ref.Version, ref.Address, selector, err)
				}
				if !verified {
					params.Env.Logger.Infof("verify hook: verifying %s %s (%s on %d)", ref.Type, ref.Version, ref.Address, selector)
					if err := v.Verify(ctx); err != nil {
						return fmt.Errorf("verify hook: error verifying %s %s (%s on %d): %w", ref.Type, ref.Version, ref.Address, selector, err)
					}
				}
				params.Env.Logger.Infof("verify hook: verified %s %s (%s on %d)", ref.Type, ref.Version, ref.Address, selector)

				return nil
			},
		)
	}
}

func requireVerifiedEnvContracts(
	dom domain.Domain,
	verifier ContractVerification,
	refsToVerify []datastore.AddressRef,
	selectors []uint64,
) changeset.PreHookFunc {
	return func(ctx context.Context, params changeset.PreHookParams) error {
		ds, err := dom.EnvDir(params.Env.Name).DataStore()
		if err != nil {
			return fmt.Errorf("require verified pre-hook: load datastore: %w", err)
		}
		// filter networks with selective chain selectors; if this fails, return an error to log the failure and skip verification.
		newDs := datastore.NewMemoryDataStore()
		for _, selector := range selectors {
			networkRefs := ds.Addresses().Filter(datastore.AddressRefByChainSelector(selector))
			for _, ref := range networkRefs {
				err := newDs.Addresses().Add(ref)
				if err != nil {
					return fmt.Errorf("require verified pre-hook: failed to add address ref %+v to filtered datastore: %w", ref, err)
				}
			}
		}
		ds = newDs.Seal()

		networkCfg, err := verifier.FilterNetworks(params.Env.Name, dom, params.Env.Logger)
		if err != nil {
			return fmt.Errorf("require verified pre-hook: load networks: %w", err)
		}
		// filter the datastore to only include refs that are in refsToVerify;
		// if refsToVerify is empty, verify all refs in the datastore that match the verifier's NeedsVerification and ForEachNetwork criteria
		ds, err = getFilteredAddressRefsForVerification(params.Env.Logger, ds, refsToVerify)
		if err != nil {
			return fmt.Errorf("require verified pre-hook: filter address refs for verification: %w", err)
		}
		return IterateVerifiers(
			ctx, ds, networkCfg, params.Env.Logger,
			"require verified pre-hook", verifier,
			func(ctx context.Context, v cldverification.Verifiable, ref datastore.AddressRef, selector uint64) error {
				params.Env.Logger.Infof("require verified pre-hook: checking is verified %+v", ref)
				verified, err := v.IsVerified(ctx)
				if err != nil {
					return fmt.Errorf("error checking verification status for address ref %+v on chain selector %d: %w", ref, selector, err)
				}
				if !verified {
					return fmt.Errorf("contract is not verified on explorer for address ref %+v on chain selector %d", ref, selector)
				}

				return nil
			},
		)
	}
}

func getFilteredAddressRefsForVerification(l logger.Logger, ds datastore.DataStore, refsToVerify []datastore.AddressRef) (datastore.DataStore, error) {
	if len(refsToVerify) > 0 {
		newDs := datastore.NewMemoryDataStore()
		for i := range refsToVerify {
			var filterFns []datastore.FilterFunc[datastore.AddressRefKey, datastore.AddressRef]
			var criteria []string
			if refsToVerify[i].Type != "" {
				filterFns = append(filterFns, datastore.AddressRefByType(refsToVerify[i].Type))
				criteria = append(criteria, fmt.Sprintf("type=%s", refsToVerify[i].Type))
			}
			if refsToVerify[i].Version != nil {
				filterFns = append(filterFns, datastore.AddressRefByVersion(refsToVerify[i].Version))
				criteria = append(criteria, fmt.Sprintf("version=%s", refsToVerify[i].Version.String()))
			}
			if refsToVerify[i].Address != "" {
				filterFns = append(filterFns, datastore.AddressRefByAddress(refsToVerify[i].Address))
				criteria = append(criteria, fmt.Sprintf("address=%s", refsToVerify[i].Address))
			}
			if refsToVerify[i].ChainSelector != 0 {
				filterFns = append(filterFns, datastore.AddressRefByChainSelector(refsToVerify[i].ChainSelector))
				criteria = append(criteria, fmt.Sprintf("chainSelector=%d", refsToVerify[i].ChainSelector))
			}
			if refsToVerify[i].Qualifier != "" {
				filterFns = append(filterFns, datastore.AddressRefByQualifier(refsToVerify[i].Qualifier))
				criteria = append(criteria, fmt.Sprintf("qualifier=%s", refsToVerify[i].Qualifier))
			}
			if len(filterFns) == 0 {
				return nil, fmt.Errorf("require verified pre-hook: invalid filter criteria %+v: "+
					"at least one field must be specified to filter address refs for verification", refsToVerify[i])
			}
			ref := ds.Addresses().Filter(filterFns...)
			if len(ref) == 0 {
				return nil, fmt.Errorf("require verified pre-hook: no address ref found for filter criteria %v", criteria)
			}
			for _, r := range ref {
				err := newDs.Addresses().Add(r)
				if err != nil {
					return nil, fmt.Errorf("require verified pre-hook: failed to add address ref %+v to filtered datastore: %w", r, err)
				}
			}
		}
		ds = newDs.Seal()
	}
	return ds, nil
}

// IterateVerifiers walks filtered networks and address refs, builds a verifier per ref
// when NeedsVerification returns true, and runs step for each.
func IterateVerifiers(
	ctx context.Context,
	ds datastore.DataStore,
	networkCfg *cfgnet.Config,
	lggr logger.Logger,
	logPrefix string,
	verifier ContractVerification,
	step func(ctx context.Context, v cldverification.Verifiable, ref datastore.AddressRef, selector uint64) error,
) error {
	var errs []error
	var errsMu sync.Mutex
	networkGrp, ctx := errgroup.WithContext(ctx)
	networkGrp.SetLimit(concurrentNetworksLimit)
	for _, network := range networkCfg.Networks() {
		network := network // capture loop variable
		networkGrp.Go(func() error {
			build, skipNetwork := verifier.ForEachNetwork(ctx, network, network.ChainSelector, lggr, logPrefix)
			if skipNetwork {
				lggr.Warnf("%s: skipping unsupported network %d", logPrefix, network.ChainSelector)
				return nil
			}

			addresses := ds.Addresses().Filter(datastore.AddressRefByChainSelector(network.ChainSelector))
			stepGrp, ctx := errgroup.WithContext(ctx)
			stepGrp.SetLimit(concurrentVerificationsLimit) // limit concurrent verifications to avoid overwhelming explorers; adjust as needed

			for _, ref := range addresses {
				ref := ref // capture loop variable
				if ref.Type == "" || ref.Version == nil || ref.Address == "" {
					errsMu.Lock()
					errs = append(errs, fmt.Errorf("%s: invalid address ref %+v, missing type, version, or address", logPrefix, ref))
					errsMu.Unlock()
					continue
				}
				if !verifier.NeedsVerification(ref) {
					continue
				}
				stepGrp.Go(func() error {
					lggr.Infof("%s: building verifier for %s %s (%s on %d)", logPrefix, ref.Type, ref.Version, ref.Address, network.ChainSelector)
					v, err := build(ctx, ref)
					if err != nil {
						errsMu.Lock()
						errs = append(errs, fmt.Errorf("%s: build verifier for %s %s (%s on %d): %w", logPrefix, ref.Type, ref.Version, ref.Address, network.ChainSelector, err))
						errsMu.Unlock()
						return nil
					}
					if err := step(ctx, v, ref, network.ChainSelector); err != nil {
						errsMu.Lock()
						errs = append(errs, fmt.Errorf("%s: step for %s %s (%s on %d): %w", logPrefix, ref.Type, ref.Version, ref.Address, network.ChainSelector, err))
						errsMu.Unlock()
						return nil
					}
					return nil
				})
			}
			if err := stepGrp.Wait(); err != nil {
				errsMu.Lock()
				errs = append(errs, fmt.Errorf("%s: wait for verifications: %w", logPrefix, err))
				errsMu.Unlock()
			}
			return nil
		})
	}
	if err := networkGrp.Wait(); err != nil {
		errsMu.Lock()
		errs = append(errs, fmt.Errorf("%s: wait for networks: %w", logPrefix, err))
		errsMu.Unlock()
	}
	return errors.Join(errs...)
}

// ResetContractVerificationRegistryForTest replaces the registry and resets sync.Once so tests
// can register providers in isolation.
func ResetContractVerificationRegistryForTest() {
	singletonContractVerificationRegistry = newContractVerificationRegistry()
	onceContractVerificationRegistry = sync.Once{}
}

func isSelectorFamilySupported(family string) bool {
	switch family {
	case chain_selectors.FamilyEVM, chain_selectors.FamilySolana, chain_selectors.FamilyStarknet,
		chain_selectors.FamilyCosmos, chain_selectors.FamilyAptos, chain_selectors.FamilySui,
		chain_selectors.FamilyTron, chain_selectors.FamilyTon, chain_selectors.FamilyCanton,
		chain_selectors.FamilyStellar:
		return true
	default:
		return false
	}
}
