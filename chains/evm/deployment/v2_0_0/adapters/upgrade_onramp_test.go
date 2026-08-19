package adapters_test

import (
	"encoding/hex"
	"fmt"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	rmnproxyops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/adapters"
	evmchangesets "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/testsetup"
	changesets "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	ccvadapters "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
	v2changesets "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/changesets"
)

var upgradeTestChains = []uint64{
	chainsel.TEST_90000001.Selector,
	chainsel.TEST_90000002.Selector,
	chainsel.TEST_90000003.Selector,
}

func deployUpgradeLaneContracts(
	t *testing.T,
	e *deployment.Environment,
	chainSelector uint64,
	ds datastore.MutableDataStore,
	deployTestRouter bool,
) {
	t.Helper()
	evmChain := e.BlockChains.EVMChains()[chainSelector]

	create2Ref, err := contract.MaybeDeployContract(
		e.OperationsBundle, create2_factory.Deploy, evmChain,
		contract.DeployInput[create2_factory.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("2.0.0")),
			ChainSelector:  chainSelector,
			Args:           create2_factory.ConstructorArgs{AllowList: []common.Address{evmChain.DeployerKey.From}},
		}, nil,
	)
	require.NoError(t, err)

	deployOut, err := evmchangesets.DeployChainContracts(changesets.GetRegistry()).Apply(*e, changesets.WithMCMS[evmchangesets.DeployChainContractsCfg]{
		Cfg: evmchangesets.DeployChainContractsCfg{
			ChainSel:         chainSelector,
			CREATE2Factory:   common.HexToAddress(create2Ref.Address),
			Params:           contractParamsForLaneTopologyTest(),
			DeployTestRouter: deployTestRouter,
			DeployerKeyOwned: true,
		},
		MCMS: mcms.Input{},
	})
	require.NoError(t, err)
	require.NoError(t, ds.Merge(deployOut.DataStore.Seal()))
}

// setupDeployNewOnRampTest deploys 2.0.0 contracts on 3 simulated chains connected to each other,
// configures lanes between all pairs, and returns a ready-to-use environment.
func setupDeployNewOnRampTest(t *testing.T) *deployment.Environment {
	t.Helper()
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, upgradeTestChains),
	)
	require.NoError(t, err)

	ds := datastore.NewMemoryDataStore()

	// Deploy chain contracts on each chain (with TestRouter for the upgrade target).
	for _, sel := range upgradeTestChains {
		e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
		err = testsetup.SeedUltraFastCurseMCMS(ds, sel)
		require.NoError(t, err)

		err = ds.Merge(e.DataStore)
		require.NoError(t, err)
		e.DataStore = ds.Seal()

		deployUpgradeLaneContracts(t, e, sel, ds, sel == upgradeTestChains[0])
	}
	e.DataStore = ds.Seal()

	// Configure lanes between all pairs so each OnRamp has dest chain configs.
	deployer := e.BlockChains.EVMChains()[upgradeTestChains[0]].DeployerKey.From.Hex()
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	cs := v2changesets.ConfigureChainsForLanesFromTopology(
		ccvadapters.GetCommitteeVerifierContractRegistry(),
		ccvadapters.GetChainFamilyRegistry(),
		changesets.GetRegistry(),
	)
	lanes := make([]v2changesets.CrossFamilyLanePair, 0, len(upgradeTestChains)-1)
	for _, remoteSel := range upgradeTestChains[1:] {
		lanes = append(lanes, v2changesets.CrossFamilyLanePair{
			ChainA: upgradeTestChains[0],
			ChainB: remoteSel,
		})
	}
	_, err = cs.Apply(*e, v2changesets.ConfigureChainsForLanesFromTopologyConfig{
		Topology: bidirectionalLaneTopology(deployer, upgradeTestChains...),
		BuildLanesCrossFamilyConfig: v2changesets.BuildLanesCrossFamilyConfig{
			Lanes: lanes,
			MCMS:  mcms.Input{},
		},
	})
	require.NoError(t, err)

	return e
}

func initialStateAssertions(t *testing.T, e deployment.Environment) datastore.AddressRef {
	t.Helper()
	ref, err := datastore_utils.FindAndFormatCanonicalRef(e.DataStore, datastore.AddressRef{
		Type:    datastore.ContractType(onramp.ContractType),
		Version: onramp.Version,
	}, upgradeTestChains[0], func(r datastore.AddressRef) (datastore.AddressRef, error) {
		return r, nil
	})
	require.NoError(t, err)
	return ref
}

func upgradeStateAssertions(t *testing.T, e deployment.Environment, result ccvadapters.OnRampUpgradeResult) {
	t.Helper()
	chain := e.BlockChains.EVMChains()[upgradeTestChains[0]]
	b := testsetup.BundleWithFreshReporter(e.OperationsBundle)

	// New OnRamp static config.
	newOnRampAddr := common.HexToAddress(result.NewOnRampRef.Address)
	staticReport, err := operations.ExecuteOperation(b, onramp.GetStaticConfig, chain, contract.FunctionInput[struct{}]{
		ChainSelector: chain.Selector,
		Address:       newOnRampAddr,
		Args:          struct{}{},
	})
	require.NoError(t, err)

	rmnProxyAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		Type:    datastore.ContractType(rmnproxyops.ContractType),
		Version: semver.MustParse("1.0.0"),
	}, upgradeTestChains[0], evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err)
	assert.Equal(t, rmnProxyAddr, staticReport.Output.RmnRemote, "new OnRamp should use RMNProxy address")

	// Phase 1 makes no router writes: TestRouter must not yet route to the new OnRamp for
	// any dest chain, since that would cut prod-router-class lanes over before the verifier
	// jobs observe both OnRamps.
	testRouterAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		Type:    datastore.ContractType(router.TestRouterContractType),
		Version: router.Version,
	}, upgradeTestChains[0], evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err)

	for _, remoteSel := range upgradeTestChains[1:] {
		onRampReport, err := operations.ExecuteOperation(b, router.GetOnRamp, chain, contract.FunctionInput[uint64]{
			ChainSelector: chain.Selector,
			Address:       testRouterAddr,
			Args:          remoteSel,
		})
		require.NoError(t, err)
		assert.Equal(t, common.Address{}, onRampReport.Output,
			"TestRouter should not yet route dest %d after Phase 1", remoteSel)
	}

	// New OnRamp dest chain configs should be copied verbatim from the legacy OnRamp,
	// i.e. still use the prod Router (this test's lanes are all prod-router class).
	prodRouterAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		Type:    datastore.ContractType(router.ContractType),
		Version: router.Version,
	}, upgradeTestChains[0], evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err)
	for _, remoteSel := range upgradeTestChains[1:] {
		destCfg, err := operations.ExecuteOperation(b, onramp.GetDestChainConfig, chain, contract.FunctionInput[uint64]{
			ChainSelector: chain.Selector,
			Address:       newOnRampAddr,
			Args:          remoteSel,
		})
		require.NoError(t, err)
		assert.Equal(t, prodRouterAddr, destCfg.Output.Router,
			"new OnRamp dest config for %d should still use the prod Router after Phase 1", remoteSel)
	}
}

// stagedOnTestRouterAssertions asserts that PromoteOnrampToTestRouter has wired the
// TestRouter (and the new OnRamp's dest chain configs) to route dest chains to the new OnRamp,
// without touching the prod Router.
func stagedOnTestRouterAssertions(t *testing.T, e deployment.Environment, result ccvadapters.OnRampUpgradeResult) {
	t.Helper()
	chain := e.BlockChains.EVMChains()[upgradeTestChains[0]]
	b := testsetup.BundleWithFreshReporter(e.OperationsBundle)

	newOnRampAddr := common.HexToAddress(result.NewOnRampRef.Address)

	testRouterAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		Type:    datastore.ContractType(router.TestRouterContractType),
		Version: router.Version,
	}, upgradeTestChains[0], evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err)

	for _, remoteSel := range upgradeTestChains[1:] {
		onRampReport, err := operations.ExecuteOperation(b, router.GetOnRamp, chain, contract.FunctionInput[uint64]{
			ChainSelector: chain.Selector,
			Address:       testRouterAddr,
			Args:          remoteSel,
		})
		require.NoError(t, err)
		assert.Equal(t, newOnRampAddr, onRampReport.Output,
			"TestRouter should map dest %d to new OnRamp", remoteSel)
	}

	for _, remoteSel := range upgradeTestChains[1:] {
		destCfg, err := operations.ExecuteOperation(b, onramp.GetDestChainConfig, chain, contract.FunctionInput[uint64]{
			ChainSelector: chain.Selector,
			Address:       newOnRampAddr,
			Args:          remoteSel,
		})
		require.NoError(t, err)
		assert.Equal(t, testRouterAddr, destCfg.Output.Router,
			"new OnRamp dest config for %d should use TestRouter", remoteSel)
	}
}

func promoteStateAssertions(t *testing.T, e deployment.Environment, result ccvadapters.OnRampUpgradeResult) {
	t.Helper()
	chain := e.BlockChains.EVMChains()[upgradeTestChains[0]]
	b := testsetup.BundleWithFreshReporter(e.OperationsBundle)

	newOnRampAddr := common.HexToAddress(result.NewOnRampRef.Address)

	prodRouterAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		Type:    datastore.ContractType(router.ContractType),
		Version: router.Version,
	}, upgradeTestChains[0], evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err)

	// Prod Router should route to the new OnRamp for each dest chain.
	for _, remoteSel := range upgradeTestChains[1:] {
		onRampReport, err := operations.ExecuteOperation(b, router.GetOnRamp, chain, contract.FunctionInput[uint64]{
			ChainSelector: chain.Selector,
			Address:       prodRouterAddr,
			Args:          remoteSel,
		})
		require.NoError(t, err)
		assert.Equal(t, newOnRampAddr, onRampReport.Output,
			"prod Router should map dest %d to new OnRamp", remoteSel)
	}

	// New OnRamp dest chain configs should use prod Router.
	for _, remoteSel := range upgradeTestChains[1:] {
		destCfg, err := operations.ExecuteOperation(b, onramp.GetDestChainConfig, chain, contract.FunctionInput[uint64]{
			ChainSelector: chain.Selector,
			Address:       newOnRampAddr,
			Args:          remoteSel,
		})
		require.NoError(t, err)
		assert.Equal(t, prodRouterAddr, destCfg.Output.Router,
			"new OnRamp dest config for %d should use prod Router", remoteSel)
	}
}

// setRouterOnRamp points routerAddr's onRamp mapping for destSel at onRampAddr. Passing the zero
// address disables the lane on that router. Unlike setDestChainRouter (which writes the OnRamp's
// dest config), this writes the Router contract itself, so the two together build a fixture whose
// OnRamp and Router agree about which router fronts a lane.
func setRouterOnRamp(
	t *testing.T,
	e *deployment.Environment,
	routerAddr common.Address,
	destSel uint64,
	onRampAddr common.Address,
) {
	t.Helper()
	chain := e.BlockChains.EVMChains()[upgradeTestChains[0]]

	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	_, err := operations.ExecuteOperation(e.OperationsBundle, router.ApplyRampUpdates, chain, contract.FunctionInput[router.ApplyRampsUpdatesArgs]{
		ChainSelector: chain.Selector,
		Address:       routerAddr,
		Args: router.ApplyRampsUpdatesArgs{
			OnRampUpdates: []router.OnRamp{{DestChainSelector: destSel, OnRamp: onRampAddr}},
		},
	})
	require.NoError(t, err)

	// The reporter memoizes reads by (operation, chain, address, args), so a stale pre-write
	// route would otherwise be served to the next caller.
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
}

// assertOnRampDestRouter asserts that onRampAddr's dest chain config for destSel routes through
// wantRouter. phase names the upgrade phase under test so failures identify themselves.
func assertOnRampDestRouter(
	t *testing.T,
	e deployment.Environment,
	onRampAddr common.Address,
	destSel uint64,
	wantRouter common.Address,
	phase string,
) {
	t.Helper()
	chain := e.BlockChains.EVMChains()[upgradeTestChains[0]]

	destCfg, err := operations.ExecuteOperation(e.OperationsBundle, onramp.GetDestChainConfig, chain, contract.FunctionInput[uint64]{
		ChainSelector: chain.Selector,
		Address:       onRampAddr,
		Args:          destSel,
	})
	require.NoError(t, err)
	assert.Equal(t, wantRouter, destCfg.Output.Router,
		"%s: OnRamp %s dest config for %d should route through %s",
		phase, onRampAddr.Hex(), destSel, wantRouter.Hex())
}

// assertRouterRoutes asserts that routerAddr maps destSel to wantOnRamp. A zero wantOnRamp asserts
// the router does not route that dest at all.
func assertRouterRoutes(
	t *testing.T,
	e deployment.Environment,
	routerAddr common.Address,
	destSel uint64,
	wantOnRamp common.Address,
	phase string,
) {
	t.Helper()
	chain := e.BlockChains.EVMChains()[upgradeTestChains[0]]

	report, err := operations.ExecuteOperation(e.OperationsBundle, router.GetOnRamp, chain, contract.FunctionInput[uint64]{
		ChainSelector: chain.Selector,
		Address:       routerAddr,
		Args:          destSel,
	})
	require.NoError(t, err)
	assert.Equal(t, wantOnRamp, report.Output,
		"%s: Router %s should map dest %d to OnRamp %s",
		phase, routerAddr.Hex(), destSel, wantOnRamp.Hex())
}

// addRandomQualifiersToContracts adds timestamp-based qualifiers to test router, prod router, offramp, and RMN proxy
// on all chains, simulating the real datastore where multiple instances of the same contract type coexist.
// The onramp is left unqualified to remain canonical during the upgrade.
func addRandomQualifiersToContracts(t *testing.T, e *deployment.Environment) {
	t.Helper()

	// Generate timestamp-based qualifiers for each contract type
	timestamp := fmt.Sprintf("%d", time.Now().UnixMilli())

	// Get all current refs from e.DataStore
	allRefs := e.DataStore.Addresses().Filter()

	// Process refs: modify qualifiers for specific contract types
	for i := range allRefs {
		ref := &allRefs[i]

		if ref.Type != datastore.ContractType(onramp.ContractType) && ref.Qualifier == "" { // Leave onramp unqualified
			ref.Qualifier = fmt.Sprintf("qualifier-%s", timestamp)
		}
	}

	// Create a new datastore with the modified refs
	newDS := datastore.NewMemoryDataStore()
	for _, ref := range allRefs {
		require.NoError(t, newDS.Addresses().Add(ref))
	}

	// Update environment datastore with the qualified versions
	e.DataStore = newDS.Seal()
}

// TestUpgradeOnrampFullFlowForInitialOnrampBehindTestRouterAndProdRouterUsingRefWithQualifier drives
// all three upgrade phases on a chain that serves both lane classes at once, against a datastore
// whose non-OnRamp refs all carry qualifiers.
//
// The qualifier step reproduces the production datastore, where Router/TestRouter/ARMProxy are
// stored under an "<address>-<Type>" qualifier with no canonical copy, while the 2.0.0 OnRamp is
// canonical. Any lookup that insists on an empty qualifier for a non-OnRamp contract therefore
// resolves nothing here, exactly as it would against prod.
//
// The lane fixture is mixed on purpose: one dest is fronted by the TestRouter (where the TestRouter
// *is* that lane's production path and it never stages), the other by the prod Router (which stages
// on the TestRouter in Phase 2 before cutting over in Phase 3). Phases are driven off
// ClassifyDestChains the same way UpgradeOnrampPhase2/Phase3 drive them, so each class takes its
// own workflow rather than both being promoted everywhere.
func TestUpgradeOnrampFullFlowForInitialOnrampBehindTestRouterAndProdRouterUsingRefWithQualifier(t *testing.T) {
	e := setupDeployNewOnRampTest(t)

	// Step 1 — before any deploy or upgrade: qualify every non-OnRamp ref.
	addRandomQualifiersToContracts(t, e)

	legacyOnRamp := common.HexToAddress(initialStateAssertions(t, *e).Address)

	// Both routers resolve from qualified refs.
	testRouterAddr := resolveRouterAddrForTest(t, *e, datastore.ContractType(router.TestRouterContractType))
	prodRouterAddr := resolveRouterAddrForTest(t, *e, datastore.ContractType(router.ContractType))

	// Build the mixed fixture. testDest becomes a genuine TestRouter-class lane: its OnRamp dest
	// config points at the TestRouter, the TestRouter routes it to the legacy OnRamp, and the prod
	// Router does not route it at all. prodDest is left as the lane setup created it.
	testDest, prodDest := upgradeTestChains[1], upgradeTestChains[2]
	setDestChainRouter(t, e, legacyOnRamp, testDest, testRouterAddr)
	setRouterOnRamp(t, e, testRouterAddr, testDest, legacyOnRamp)
	setRouterOnRamp(t, e, prodRouterAddr, testDest, common.Address{})

	adapter := adapters.EVMOnRampUpgrader{}

	class, err := adapter.ClassifyDestChains(*e, upgradeTestChains[0])
	require.NoError(t, err)
	require.ElementsMatch(t, []uint64{prodDest}, class.ProdRouterDests)
	require.ElementsMatch(t, []uint64{testDest}, class.TestRouterDests)

	// ---- Phase 1: deploy the replacement OnRamp ----
	result, err := adapter.DeployNewOnRamp(*e, upgradeTestChains[0])
	require.NoError(t, err)
	mergeUpgradeRefs(t, e, result)
	newOnRamp := common.HexToAddress(result.NewOnRampRef.Address)

	assert.Equal(t, legacyOnRamp, common.HexToAddress(result.LegacyOnRampRef.Address))
	assert.Equal(t, "legacy", result.LegacyOnRampRef.Qualifier)
	assert.Equal(t, "", result.NewOnRampRef.Qualifier)
	assert.NotEqual(t, legacyOnRamp, newOnRamp)

	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)

	// The replacement OnRamp must adopt the RMNProxy — resolved here from a qualified ref.
	rmnProxyAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		Type:    datastore.ContractType(rmnproxyops.ContractType),
		Version: semver.MustParse("1.0.0"),
	}, upgradeTestChains[0], evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err)
	staticReport, err := operations.ExecuteOperation(e.OperationsBundle, onramp.GetStaticConfig,
		e.BlockChains.EVMChains()[upgradeTestChains[0]], contract.FunctionInput[struct{}]{
			ChainSelector: upgradeTestChains[0],
			Address:       newOnRamp,
			Args:          struct{}{},
		})
	require.NoError(t, err)
	assert.Equal(t, rmnProxyAddr, staticReport.Output.RmnRemote, "new OnRamp should use the RMNProxy")

	// Phase 1 copies dest configs verbatim and makes no router writes: each class keeps its own
	// router on the new OnRamp, and both routers still point at the legacy OnRamp.
	assertOnRampDestRouter(t, *e, newOnRamp, prodDest, prodRouterAddr, "after Phase 1")
	assertOnRampDestRouter(t, *e, newOnRamp, testDest, testRouterAddr, "after Phase 1")
	assertRouterRoutes(t, *e, prodRouterAddr, prodDest, legacyOnRamp, "after Phase 1")
	assertRouterRoutes(t, *e, testRouterAddr, testDest, legacyOnRamp, "after Phase 1")
	assertRouterRoutes(t, *e, testRouterAddr, prodDest, common.Address{}, "after Phase 1")

	// ---- Phase 2: stage the prod-class lane on the TestRouter ----
	// Only ProdRouterDests stage. The TestRouter-class lane is untouched, since cutting it over
	// here would put it into production before the verifier jobs observe both OnRamps.
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	_, err = adapter.PromoteOnrampToTestRouter(*e, upgradeTestChains[0], class.ProdRouterDests)
	require.NoError(t, err)

	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	assertRouterRoutes(t, *e, testRouterAddr, prodDest, newOnRamp, "after Phase 2")
	assertOnRampDestRouter(t, *e, newOnRamp, prodDest, testRouterAddr, "after Phase 2")
	// The prod-class lane's production path is untouched.
	assertRouterRoutes(t, *e, prodRouterAddr, prodDest, legacyOnRamp, "after Phase 2")
	require.NoError(t, adapter.VerifyLegacyOnRampOnProdRouter(*e, upgradeTestChains[0], class.ProdRouterDests))
	// The test-class lane is untouched.
	assertRouterRoutes(t, *e, testRouterAddr, testDest, legacyOnRamp, "after Phase 2")
	assertOnRampDestRouter(t, *e, newOnRamp, testDest, testRouterAddr, "after Phase 2")

	// Staging must not reclassify the prod-class lane. Classification reads the legacy OnRamp
	// precisely so the Phase 2 write above cannot flip prodDest into the TestRouter bucket.
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	class, err = adapter.ClassifyDestChains(*e, upgradeTestChains[0])
	require.NoError(t, err)
	require.ElementsMatch(t, []uint64{prodDest}, class.ProdRouterDests)
	require.ElementsMatch(t, []uint64{testDest}, class.TestRouterDests)

	// ---- Phase 3: promote both classes to their production router ----
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	_, err = adapter.PromoteOnrampToTestRouter(*e, upgradeTestChains[0], class.TestRouterDests)
	require.NoError(t, err)

	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	_, err = adapter.PromoteOnrampToProdRouter(*e, upgradeTestChains[0], class.ProdRouterDests)
	require.NoError(t, err)

	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	assertRouterRoutes(t, *e, prodRouterAddr, prodDest, newOnRamp, "after Phase 3")
	assertOnRampDestRouter(t, *e, newOnRamp, prodDest, prodRouterAddr, "after Phase 3")
	assertRouterRoutes(t, *e, testRouterAddr, testDest, newOnRamp, "after Phase 3")
	assertOnRampDestRouter(t, *e, newOnRamp, testDest, testRouterAddr, "after Phase 3")

	// The adapter's own end-state check must agree, per class.
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	require.NoError(t, adapter.VerifyPromotedToRouters(
		*e, upgradeTestChains[0], class.ProdRouterDests, class.TestRouterDests))
}

func TestUpgradeOnrampFullFlowForInitialOnrampBehindTestRouter(t *testing.T) {
	e := setupDeployNewOnRampTest(t)

	existingOnRamp := initialStateAssertions(t, *e)
	adapter := adapters.EVMOnRampUpgrader{}

	result, err := adapter.DeployNewOnRamp(*e, upgradeTestChains[0])
	require.NoError(t, err)

	// Merge new OnRamp ref into the environment datastore so subsequent calls
	// (PromoteOnrampToProdRouter) can find the canonical OnRamp.
	mergeDS := datastore.NewMemoryDataStore()
	require.NoError(t, mergeDS.Addresses().Add(result.NewOnRampRef))
	require.NoError(t, mergeDS.Addresses().Add(result.LegacyOnRampRef))
	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Merge(e.DataStore))
	require.NoError(t, ds.Merge(mergeDS.Seal()))
	e.DataStore = ds.Seal()

	// Legacy ref should match the old OnRamp address.
	assert.Equal(t, existingOnRamp.Address, result.LegacyOnRampRef.Address)
	assert.Equal(t, "legacy", result.LegacyOnRampRef.Qualifier)

	// New ref should be canonical (empty qualifier).
	assert.Equal(t, "", result.NewOnRampRef.Qualifier)
	assert.NotEqual(t, existingOnRamp.Address, result.NewOnRampRef.Address)

	upgradeStateAssertions(t, *e, result)
}

func TestUpgradeOnrampFullFlowForInitialOnrampNotBehindProdRouter(t *testing.T) {
	e := setupDeployNewOnRampTest(t)

	existingOnRamp := initialStateAssertions(t, *e)
	adapter := adapters.EVMOnRampUpgrader{}

	result, err := adapter.DeployNewOnRamp(*e, upgradeTestChains[0])
	require.NoError(t, err)

	// Merge new OnRamp ref into the environment datastore so PromoteOnrampToProdRouter
	// can find the canonical OnRamp.
	mergeDS := datastore.NewMemoryDataStore()
	require.NoError(t, mergeDS.Addresses().Add(result.NewOnRampRef))
	require.NoError(t, mergeDS.Addresses().Add(result.LegacyOnRampRef))
	newDS := datastore.NewMemoryDataStore()
	require.NoError(t, newDS.Merge(e.DataStore))
	require.NoError(t, newDS.Merge(mergeDS.Seal()))
	e.DataStore = newDS.Seal()

	assert.Equal(t, existingOnRamp.Address, result.LegacyOnRampRef.Address)
	assert.Equal(t, "legacy", result.LegacyOnRampRef.Qualifier)
	assert.Equal(t, "", result.NewOnRampRef.Qualifier)
	assert.NotEqual(t, existingOnRamp.Address, result.NewOnRampRef.Address)

	upgradeStateAssertions(t, *e, result)

	// Stage on the TestRouter first (Phase 2) — this must not touch the prod Router.
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	_, err = adapter.PromoteOnrampToTestRouter(*e, upgradeTestChains[0], upgradeTestChains[1:])
	require.NoError(t, err)
	stagedOnTestRouterAssertions(t, *e, result)

	// Now promote to prod router (Phase 3).
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	_, err = adapter.PromoteOnrampToProdRouter(*e, upgradeTestChains[0], upgradeTestChains[1:])
	require.NoError(t, err)

	promoteStateAssertions(t, *e, result)
}

func TestFullFlowWithRollback(t *testing.T) {
	e := setupDeployNewOnRampTest(t)

	legacyOnRamp := initialStateAssertions(t, *e)
	adapter := adapters.EVMOnRampUpgrader{}

	// Verify routers are originally on legacy OnRamp
	chain := e.BlockChains.EVMChains()[upgradeTestChains[0]]
	legacyOnRampAddr := common.HexToAddress(legacyOnRamp.Address)
	prodRouterAddr := resolveRouterAddrForTest(t, *e, datastore.ContractType(router.ContractType))
	b := testsetup.BundleWithFreshReporter(e.OperationsBundle)
	for _, remoteSel := range upgradeTestChains[1:] {
		onRampReport, err := operations.ExecuteOperation(b, router.GetOnRamp, chain, contract.FunctionInput[uint64]{
			ChainSelector: chain.Selector,
			Address:       prodRouterAddr,
			Args:          remoteSel,
		})
		require.NoError(t, err)
		assert.Equal(t, legacyOnRampAddr, onRampReport.Output,
			"prod Router should map dest %d back to legacy OnRamp after rollback", remoteSel)
	}

	// Phase 1: Deploy new OnRamp
	result, err := adapter.DeployNewOnRamp(*e, upgradeTestChains[0])
	require.NoError(t, err)
	mergeUpgradeRefs(t, e, result)
	upgradeStateAssertions(t, *e, result)

	// Phase 2: Stage on TestRouter
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	_, err = adapter.PromoteOnrampToTestRouter(*e, upgradeTestChains[0], upgradeTestChains[1:])
	require.NoError(t, err)
	stagedOnTestRouterAssertions(t, *e, result)

	// Phase 3: Promote to prod Router
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	_, err = adapter.PromoteOnrampToProdRouter(*e, upgradeTestChains[0], upgradeTestChains[1:])
	require.NoError(t, err)
	promoteStateAssertions(t, *e, result)

	// Rollback: Return routers to legacy OnRamp
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	_, err = adapter.RollbackToLegacyRouters(*e, upgradeTestChains[0], upgradeTestChains[1:], []uint64{})
	require.NoError(t, err)

	// Verify routers are back on legacy OnRamp
	b = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	for _, remoteSel := range upgradeTestChains[1:] {
		onRampReport, err := operations.ExecuteOperation(b, router.GetOnRamp, chain, contract.FunctionInput[uint64]{
			ChainSelector: chain.Selector,
			Address:       prodRouterAddr,
			Args:          remoteSel,
		})
		require.NoError(t, err)
		assert.Equal(t, legacyOnRampAddr, onRampReport.Output,
			"prod Router should map dest %d back to legacy OnRamp after rollback", remoteSel)
	}
}

func TestFullFlowWithCleanup(t *testing.T) {
	e := setupDeployNewOnRampTest(t)

	legacyOnRamp := initialStateAssertions(t, *e)
	adapter := adapters.EVMOnRampUpgrader{}

	// Phase 1: Deploy new OnRamp
	result, err := adapter.DeployNewOnRamp(*e, upgradeTestChains[0])
	require.NoError(t, err)
	mergeUpgradeRefs(t, e, result)
	upgradeStateAssertions(t, *e, result)

	// Phase 1 (continued): Update remote OffRamps to allow both legacy and new OnRamps
	// (in the real changeset, this is done via SetOffRampSourceOnRamps, but here we simulate it)
	chainFamilyAdapter := &adapters.ChainFamilyAdapter{}
	newOnRampAddr := common.HexToAddress(result.NewOnRampRef.Address)
	legacyOnRampAddr := common.HexToAddress(legacyOnRamp.Address)

	for _, remoteSel := range upgradeTestChains[1:] {
		e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
		_, _, err := chainFamilyAdapter.SetOffRampSourceOnRamps(*e, ccvadapters.OffRampSetSourceOnRampsEntry{
			LocalChainSelector:  remoteSel,
			SourceChainSelector: upgradeTestChains[0],
			OnRamps: []string{
				"0x" + hex.EncodeToString(common.LeftPadBytes(legacyOnRampAddr.Bytes(), 32)),
				"0x" + hex.EncodeToString(common.LeftPadBytes(newOnRampAddr.Bytes(), 32)),
			},
		})
		require.NoError(t, err)
	}

	// Phase 2: Stage on TestRouter
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	_, err = adapter.PromoteOnrampToTestRouter(*e, upgradeTestChains[0], upgradeTestChains[1:])
	require.NoError(t, err)
	stagedOnTestRouterAssertions(t, *e, result)

	// Phase 3: Promote to prod Router
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	_, err = adapter.PromoteOnrampToProdRouter(*e, upgradeTestChains[0], upgradeTestChains[1:])
	require.NoError(t, err)
	promoteStateAssertions(t, *e, result)

	// After Phase 3, verify OffRamps still contain both legacy and new OnRamps (cleanup removes legacy)
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	for _, remoteSel := range upgradeTestChains[1:] {
		sourceOnRamps, err := chainFamilyAdapter.GetOffRampSourceOnRamps(*e, remoteSel, upgradeTestChains[0])
		require.NoError(t, err)

		expectedNewOnRamp := common.LeftPadBytes(newOnRampAddr.Bytes(), 32)
		expectedLegacyOnRamp := common.LeftPadBytes(legacyOnRampAddr.Bytes(), 32)

		require.Len(t, sourceOnRamps, 2, "OffRamp for remote chain %d should whitelist both legacy and new OnRamps before cleanup", remoteSel)

		// Verify both are present
		hasNewOnRamp := false
		hasLegacyOnRamp := false
		for _, onRamp := range sourceOnRamps {
			if assert.ObjectsAreEqual(expectedNewOnRamp, onRamp) {
				hasNewOnRamp = true
			}
			if assert.ObjectsAreEqual(expectedLegacyOnRamp, onRamp) {
				hasLegacyOnRamp = true
			}
		}
		assert.True(t, hasNewOnRamp, "OffRamp for remote chain %d should contain new OnRamp", remoteSel)
		assert.True(t, hasLegacyOnRamp, "OffRamp for remote chain %d should contain legacy OnRamp before cleanup", remoteSel)
	}

	// Cleanup: Remove legacy OnRamp from remote OffRamp whitelists
	for _, remoteSel := range upgradeTestChains[1:] {
		e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
		_, _, err := chainFamilyAdapter.SetOffRampSourceOnRamps(*e, ccvadapters.OffRampSetSourceOnRampsEntry{
			LocalChainSelector:  remoteSel,
			SourceChainSelector: upgradeTestChains[0],
			OnRamps: []string{
				"0x" + hex.EncodeToString(common.LeftPadBytes(newOnRampAddr.Bytes(), 32)),
			},
		})
		require.NoError(t, err)
	}

	// After cleanup, verify OffRamps only contain the new OnRamp (legacy removed)
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	for _, remoteSel := range upgradeTestChains[1:] {
		sourceOnRamps, err := chainFamilyAdapter.GetOffRampSourceOnRamps(*e, remoteSel, upgradeTestChains[0])
		require.NoError(t, err)

		expectedNewOnRamp := common.LeftPadBytes(newOnRampAddr.Bytes(), 32)
		expectedLegacyOnRamp := common.LeftPadBytes(legacyOnRampAddr.Bytes(), 32)

		require.Len(t, sourceOnRamps, 1, "OffRamp for remote chain %d should only whitelist the new OnRamp after cleanup", remoteSel)

		// Verify only new OnRamp is present
		assert.Equal(t, expectedNewOnRamp, sourceOnRamps[0],
			"OffRamp for remote chain %d should only contain new OnRamp after cleanup", remoteSel)

		// Ensure legacy is not in the list
		for _, onRamp := range sourceOnRamps {
			assert.NotEqual(t, expectedLegacyOnRamp, onRamp,
				"OffRamp for remote chain %d should not contain legacy OnRamp after cleanup", remoteSel)
		}
	}
}

func TestDeployNewOnRampIdempotent(t *testing.T) {
	e := setupDeployNewOnRampTest(t)
	adapter := adapters.EVMOnRampUpgrader{}

	result, err := adapter.DeployNewOnRamp(*e, upgradeTestChains[0])
	require.NoError(t, err)

	// Merge new OnRamp ref into the environment datastore so subsequent calls
	// (PromoteOnrampToProdRouter) can find the canonical OnRamp.
	mergeDS := datastore.NewMemoryDataStore()
	require.NoError(t, mergeDS.Addresses().Add(result.NewOnRampRef))
	require.NoError(t, mergeDS.Addresses().Add(result.LegacyOnRampRef))
	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Merge(e.DataStore))
	require.NoError(t, ds.Merge(mergeDS.Seal()))
	e.DataStore = ds.Seal()
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	// Re-deploying should return the same canonical OnRamp ref and no writes.
	result2, err := adapter.DeployNewOnRamp(*e, upgradeTestChains[0])
	require.NoError(t, err)
	assert.Equal(t, result.NewOnRampRef.Address, result2.NewOnRampRef.Address, "New OnRamp ref should be canonical and unchanged on re-deploy")
	assert.Equal(t, result.LegacyOnRampRef.Address, result2.LegacyOnRampRef.Address, "Legacy OnRamp ref should be unchanged on re-deploy")
	require.True(t, result2.Reused)
	assert.Equal(t, result.NewOnRampRef.Address, result2.NewOnRampRef.Address)
	assert.Equal(t, result.LegacyOnRampRef.Address, result2.LegacyOnRampRef.Address)
	assert.Empty(t, result2.BatchOps, "reusing an existing Phase 1 upgrade must not emit deployment/config writes")
}

func TestGetOffRampSourceOnRamps(t *testing.T) {
	e := setupDeployNewOnRampTest(t)
	adapter := &adapters.ChainFamilyAdapter{}

	// Fresh reporter: the setup bundle's reporter still holds the lane-config read
	// from when the whitelist was empty, and ExecuteOperation dedupes on it.
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)

	// The lane setup whitelists chain0's canonical OnRamp on chain1's OffRamp,
	// in wire format (32-byte left-padded).
	onRampRef := initialStateAssertions(t, *e)
	expected := common.LeftPadBytes(common.HexToAddress(onRampRef.Address).Bytes(), 32)

	got, err := adapter.GetOffRampSourceOnRamps(*e, upgradeTestChains[1], upgradeTestChains[0])
	require.NoError(t, err)
	require.Equal(t, [][]byte{expected}, got)
}

func TestLegacyOnRampRef(t *testing.T) {
	e := setupDeployNewOnRampTest(t)
	adapter := adapters.EVMOnRampUpgrader{}

	// No legacy ref exists before an upgrade.
	_, err := adapter.LegacyOnRampRef(*e, upgradeTestChains[0])
	require.Error(t, err)

	existingOnRamp := initialStateAssertions(t, *e)
	result, err := adapter.DeployNewOnRamp(*e, upgradeTestChains[0])
	require.NoError(t, err)

	// Merge new/legacy refs so the datastore reflects the post-upgrade state.
	mergeUpgradeRefs(t, e, result)

	ref, err := adapter.LegacyOnRampRef(*e, upgradeTestChains[0])
	require.NoError(t, err)
	assert.Equal(t, existingOnRamp.Address, ref.Address)
	assert.Equal(t, "legacy", ref.Qualifier)
}

// mergeUpgradeRefs merges the upgrade result refs into the environment datastore,
// simulating the CLD merge that happens after a changeset executes.
func mergeUpgradeRefs(t *testing.T, e *deployment.Environment, result ccvadapters.OnRampUpgradeResult) {
	t.Helper()
	mergeDS := datastore.NewMemoryDataStore()
	require.NoError(t, mergeDS.Addresses().Add(result.NewOnRampRef))
	require.NoError(t, mergeDS.Addresses().Add(result.LegacyOnRampRef))
	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Merge(e.DataStore))
	require.NoError(t, ds.Merge(mergeDS.Seal()))
	e.DataStore = ds.Seal()
}

// TestDeployNewOnRampAlreadyFixed guards against a pointless migration: the canonical
// OnRamp already uses the RMNProxy, so there is nothing to upgrade.
func TestDeployNewOnRampAlreadyFixed(t *testing.T) {
	e := setupDeployNewOnRampTest(t)
	adapter := adapters.EVMOnRampUpgrader{}

	result, err := adapter.DeployNewOnRamp(*e, upgradeTestChains[0])
	require.NoError(t, err)
	mergeUpgradeRefs(t, e, result)

	err = adapter.VerifyOnrampRequireUpgrade(*e, upgradeTestChains[0])
	require.ErrorContains(t, err, "already uses RMNProxy")
}

func TestVerifyPromotedToProdRouter(t *testing.T) {
	e := setupDeployNewOnRampTest(t)
	adapter := adapters.EVMOnRampUpgrader{}

	result, err := adapter.DeployNewOnRamp(*e, upgradeTestChains[0])
	require.NoError(t, err)
	mergeUpgradeRefs(t, e, result)

	// Not yet promoted: the prod Router still routes through the old OnRamp.
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	err = adapter.VerifyPromotedToRouters(*e, upgradeTestChains[0], upgradeTestChains[1:], []uint64{})
	require.ErrorContains(t, err, "Phase 3 must execute first")

	_, err = adapter.PromoteOnrampToProdRouter(*e, upgradeTestChains[0], upgradeTestChains[1:])
	require.NoError(t, err)

	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	require.NoError(t, adapter.VerifyPromotedToRouters(*e, upgradeTestChains[0], upgradeTestChains[1:], []uint64{}))
}

func TestVerifyLegacyOnRampOnProdRouter(t *testing.T) {
	e := setupDeployNewOnRampTest(t)
	adapter := adapters.EVMOnRampUpgrader{}

	result, err := adapter.DeployNewOnRamp(*e, upgradeTestChains[0])
	require.NoError(t, err)
	mergeUpgradeRefs(t, e, result)

	// Until Phase 2 executes, the prod Router still routes through the legacy OnRamp.
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	require.NoError(t, adapter.VerifyLegacyOnRampOnProdRouter(*e, upgradeTestChains[0], upgradeTestChains[1:]))

	_, err = adapter.PromoteOnrampToProdRouter(*e, upgradeTestChains[0], upgradeTestChains[1:])
	require.NoError(t, err)

	// After promotion the prod Router points at the new OnRamp instead.
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	require.ErrorContains(t,
		adapter.VerifyLegacyOnRampOnProdRouter(*e, upgradeTestChains[0], upgradeTestChains[1:]),
		"prod Router does not route")
}

func TestVerifyNewOnRampOwner(t *testing.T) {
	e := setupDeployNewOnRampTest(t)
	adapter := adapters.EVMOnRampUpgrader{}

	result, err := adapter.DeployNewOnRamp(*e, upgradeTestChains[0])
	require.NoError(t, err)
	mergeUpgradeRefs(t, e, result)

	// The new OnRamp is deployer-owned until Phase 1's ownership transfer executes.
	deployer := e.BlockChains.EVMChains()[upgradeTestChains[0]].DeployerKey.From.Hex()
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	require.NoError(t, adapter.VerifyNewOnRampOwner(*e, upgradeTestChains[0], deployer))
	require.ErrorContains(t,
		adapter.VerifyNewOnRampOwner(*e, upgradeTestChains[0], "0x000000000000000000000000000000000000dEaD"),
		"owned by")
}

// resolveRouterAddrForTest resolves a router address of the given contract type on the
// upgrade target chain.
func resolveRouterAddrForTest(t *testing.T, e deployment.Environment, contractType datastore.ContractType) common.Address {
	t.Helper()
	addr, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		Type:    contractType,
		Version: router.Version,
	}, upgradeTestChains[0], evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err)
	return addr
}

// setDestChainRouter repoints destSel's Router on the OnRamp at onRampAddr, leaving every other
// field of that dest config — and every other dest config — untouched. The lane-config changesets
// only ever write a single router for a whole chain, so this is how the tests below build a
// mixed-router fixture (and simulate Phase 2's TestRouter staging) without a full reconfiguration.
func setDestChainRouter(
	t *testing.T,
	e *deployment.Environment,
	onRampAddr common.Address,
	destSel uint64,
	routerAddr common.Address,
) {
	t.Helper()
	chain := e.BlockChains.EVMChains()[upgradeTestChains[0]]

	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	report, err := operations.ExecuteOperation(e.OperationsBundle, onramp.GetAllDestChainConfigs, chain, contract.FunctionInput[struct{}]{
		ChainSelector: chain.Selector,
		Address:       onRampAddr,
		Args:          struct{}{},
	})
	require.NoError(t, err)

	selectors := report.Output.Ret0
	require.NotEmpty(t, selectors, "OnRamp %s has no dest chain configs to update", onRampAddr.Hex())

	args := make([]onramp.DestChainConfigArgs, len(selectors))
	for i, cfg := range report.Output.Ret1 {
		args[i] = onramp.DestChainConfigArgs{
			DestChainSelector:         selectors[i],
			Router:                    cfg.Router,
			AddressBytesLength:        cfg.AddressBytesLength,
			TokenReceiverAllowed:      cfg.TokenReceiverAllowed,
			MessageNetworkFeeUSDCents: cfg.MessageNetworkFeeUSDCents,
			TokenNetworkFeeUSDCents:   cfg.TokenNetworkFeeUSDCents,
			BaseExecutionGasCost:      cfg.BaseExecutionGasCost,
			DefaultCCVs:               cfg.DefaultCCVs,
			LaneMandatedCCVs:          cfg.LaneMandatedCCVs,
			DefaultExecutor:           cfg.DefaultExecutor,
			OffRamp:                   cfg.OffRamp,
		}
		if selectors[i] == destSel {
			args[i].Router = routerAddr
		}
	}

	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	_, err = operations.ExecuteOperation(e.OperationsBundle, onramp.ApplyDestChainConfigUpdates, chain, contract.FunctionInput[[]onramp.DestChainConfigArgs]{
		ChainSelector: chain.Selector,
		Address:       onRampAddr,
		Args:          args,
	})
	require.NoError(t, err)

	// The reporter memoizes reads by (operation, chain, address, args), so a stale
	// pre-write dest-config report would otherwise be served to the next caller.
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
}

// The lane setup wires every lane through the prod Router, so a freshly set up chain has no
// TestRouter-class lanes at all.
func TestClassifyDestChains_AllProdRouterClass(t *testing.T) {
	e := setupDeployNewOnRampTest(t)
	adapter := adapters.EVMOnRampUpgrader{}

	class, err := adapter.ClassifyDestChains(*e, upgradeTestChains[0])
	require.NoError(t, err)

	assert.ElementsMatch(t, upgradeTestChains[1:], class.ProdRouterDests)
	assert.Empty(t, class.TestRouterDests)
}

// A chain may serve both lane classes at once: chain1 fronted by the TestRouter (where the
// TestRouter *is* that lane's production path) and chain2 by the prod Router. Each must land in
// its own bucket so the phases can drive the right workflow per dest.
func TestClassifyDestChains_MixedRouterClasses(t *testing.T) {
	e := setupDeployNewOnRampTest(t)
	adapter := adapters.EVMOnRampUpgrader{}

	onRamp := common.HexToAddress(initialStateAssertions(t, *e).Address)
	testRouterAddr := resolveRouterAddrForTest(t, *e, datastore.ContractType(router.TestRouterContractType))
	setDestChainRouter(t, e, onRamp, upgradeTestChains[1], testRouterAddr)

	class, err := adapter.ClassifyDestChains(*e, upgradeTestChains[0])
	require.NoError(t, err)

	assert.ElementsMatch(t, []uint64{upgradeTestChains[2]}, class.ProdRouterDests)
	assert.ElementsMatch(t, []uint64{upgradeTestChains[1]}, class.TestRouterDests)
}

func TestClassifyDestChains_AllTestRouterClass(t *testing.T) {
	e := setupDeployNewOnRampTest(t)
	adapter := adapters.EVMOnRampUpgrader{}

	onRamp := common.HexToAddress(initialStateAssertions(t, *e).Address)
	testRouterAddr := resolveRouterAddrForTest(t, *e, datastore.ContractType(router.TestRouterContractType))
	for _, destSel := range upgradeTestChains[1:] {
		setDestChainRouter(t, e, onRamp, destSel, testRouterAddr)
	}

	class, err := adapter.ClassifyDestChains(*e, upgradeTestChains[0])
	require.NoError(t, err)

	assert.Empty(t, class.ProdRouterDests)
	assert.ElementsMatch(t, upgradeTestChains[1:], class.TestRouterDests)
}

// Regression: classification must read the legacy-qualified OnRamp once Phase 1 has run.
// Phase 1 makes the *new* OnRamp canonical, and Phase 2 then repoints its prod-class dest configs
// at the TestRouter for staging. Classifying off the canonical ref would therefore see
// Router == TestRouter and misclassify those dests as TestRouter-class, so Phase 3 would promote
// them back to the TestRouter instead of the prod Router — leaving prod traffic on the legacy OnRamp.
func TestClassifyDestChains_UsesLegacyRefAfterStaging(t *testing.T) {
	e := setupDeployNewOnRampTest(t)
	adapter := adapters.EVMOnRampUpgrader{}

	result, err := adapter.DeployNewOnRamp(*e, upgradeTestChains[0])
	require.NoError(t, err)
	mergeUpgradeRefs(t, e, result)

	// Simulate Phase 2 staging: the new (canonical) OnRamp's dest configs now point at the
	// TestRouter, while the legacy OnRamp still records their original prod Router.
	newOnRamp := common.HexToAddress(result.NewOnRampRef.Address)
	testRouterAddr := resolveRouterAddrForTest(t, *e, datastore.ContractType(router.TestRouterContractType))
	for _, destSel := range upgradeTestChains[1:] {
		setDestChainRouter(t, e, newOnRamp, destSel, testRouterAddr)
	}

	class, err := adapter.ClassifyDestChains(*e, upgradeTestChains[0])
	require.NoError(t, err)

	assert.ElementsMatch(t, upgradeTestChains[1:], class.ProdRouterDests,
		"staged dests must stay prod-router class so Phase 3 promotes them to the prod Router")
	assert.Empty(t, class.TestRouterDests)
}
