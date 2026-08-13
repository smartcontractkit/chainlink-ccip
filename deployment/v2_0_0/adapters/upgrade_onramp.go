package adapters

import (
	"sync"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// OnRampUpgrader is implemented by chain family adapters that can upgrade an
// OnRamp contract and promote it from TestRouter to
// the production Router.
//
// Deployer-key writes on the fresh OnRamp (ApplyDestChainConfigUpdates,
// SetDynamicConfig) are executed directly. The returned BatchOps contain only
// writes to timelock-owned contracts (if any).
type OnRampUpgrader interface {
	// VerifyOnrampRequireUpgrade returns no error if the canonical OnRamp is not yet upgraded to the new version, and an error otherwise
	VerifyOnrampRequireUpgrade(e cldf.Environment, chainSelector uint64) error
	DeployNewOnRamp(e cldf.Environment, chainSelector uint64) (OnRampUpgradeResult, error)
	// ClassifyDestChains partitions the legacy OnRamp's dest chains by which router
	// currently fronts them (prod Router vs TestRouter). A dest chain routed through
	// neither is an error: every lane must be classifiable before an upgrade starts.
	ClassifyDestChains(e cldf.Environment, chainSelector uint64) (LaneClass, error)
	// DestChainSelectors returns the remote chain selectors that the canonical
	// OnRamp has lanes to. The implementation reads the OnRamp's dest chain
	// configs on-chain in a single RPC call (e.g. getAllDestChainConfigs).
	DestChainSelectors(e cldf.Environment, chainSelector uint64) ([]uint64, error)
	// LegacyOnRampRef returns the datastore ref of the legacy-qualified OnRamp
	// left behind by DeployNewOnRamp. It returns an error when no legacy OnRamp
	// exists (e.g. cleanup already completed and the ref was removed).
	LegacyOnRampRef(e cldf.Environment, chainSelector uint64) (datastore.AddressRef, error)
	// VerifyNewOnRampOwner returns an error unless the canonical (new) OnRamp is owned
	// by expectedOwner. Phase 2/3 use it as a pre-flight check to confirm Phase 1's
	// ownership-transfer proposal has executed; otherwise its writes would execute
	// directly with the deployer key instead of going through MCMS.
	VerifyNewOnRampOwner(e cldf.Environment, chainSelector uint64, expectedOwner string) error
	// PromoteOnrampToTestRouter points destSelectors' dest chain configs and the TestRouter at
	// the new OnRamp. Used for TestRouter lanes, whose only promotion is this one (their
	// production path is the TestRouter itself).
	PromoteOnrampToTestRouter(e cldf.Environment, chainSelector uint64, destSelectors []uint64) ([]mcms_types.BatchOperation, error)
	// PromoteOnrampToProdRouter points destSelectors' dest chain configs and the prod Router at
	// the new OnRamp. Used for ProdRouter lanes after they have been staged and smoke
	// tested behind the TestRouter.
	PromoteOnrampToProdRouter(e cldf.Environment, chainSelector uint64, destSelectors []uint64) ([]mcms_types.BatchOperation, error)
	// VerifyPromotedToRouters returns an error unless the production Router routes
	// destSelectors of the canonical OnRamp through it (i.e. Phase 3 has executed for them).
	// Cleanup uses it as a pre-flight check before removing the legacy OnRamp from remote
	// OffRamp whitelists.
	VerifyPromotedToRouters(e cldf.Environment, chainSelector uint64, prodDestSelectors []uint64, testDestSelectors []uint64) error
	// VerifyLegacyOnRampOnProdRouter returns an error unless the production Router still
	// routes destSelectors through the legacy OnRamp. Phase 3 uses it as a pre-flight check:
	// anything else means Phase 3 already ran or an unexpected OnRamp is live for that dest.
	VerifyLegacyOnRampOnProdRouter(e cldf.Environment, chainSelector uint64, destSelectors []uint64) error
}

// LaneClass partitions a chain's dest chains by which router currently fronts them.
// ProdRouter lanes are staged behind the TestRouter before being promoted, so they
// can be smoke tested; TestRouter lanes are production traffic for that lane and are
// promoted straight to the TestRouter once the verifier jobs observe both OnRamps.
type LaneClass struct {
	ProdRouterDests []uint64
	TestRouterDests []uint64
}

// OnRampUpgradeResult is returned by OnRampUpgrader.DeployNewOnRamp.
type OnRampUpgradeResult struct {
	NewOnRampRef    datastore.AddressRef
	LegacyOnRampRef datastore.AddressRef
	BatchOps        []mcms_types.BatchOperation
}

// OnRampUpgraderRegistry maps chain families to OnRampUpgrader implementations.
type OnRampUpgraderRegistry struct {
	mu        sync.Mutex
	upgraders map[string]OnRampUpgrader
}

// NewOnRampUpgraderRegistry creates an empty registry.
func NewOnRampUpgraderRegistry() *OnRampUpgraderRegistry {
	return &OnRampUpgraderRegistry{
		upgraders: make(map[string]OnRampUpgrader),
	}
}

func GetOnRampUpgraderRegistry() *OnRampUpgraderRegistry {
	upgraderRegistrySingletonOnce.Do(func() {
		upgraderRegistrySingleton = NewOnRampUpgraderRegistry()
	})
	return upgraderRegistrySingleton
}

var (
	upgraderRegistrySingleton     *OnRampUpgraderRegistry
	upgraderRegistrySingletonOnce sync.Once
)

// Register associates an OnRampUpgrader with a chain family.
func (r *OnRampUpgraderRegistry) Register(family string, upgrader OnRampUpgrader) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.upgraders[family]; !exists {
		r.upgraders[family] = upgrader
	}
}

// Get returns the OnRampUpgrader for the given chain family.
func (r *OnRampUpgraderRegistry) Get(family string) (OnRampUpgrader, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	a, ok := r.upgraders[family]
	return a, ok
}
