package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"k8s.io/utils/ptr"

	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	evmds "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	rmnproxyops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/onramp"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	ccvadapters "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
)

const (
	legacyQualifier  = "legacy"
	upgradeQualifier = "redeploy"
)

// testVerifierResolverQualifier is the fixed CREATE2 salt for the resolver wrapping the test
// verifier. It must match the qualifier used by the DeployTestVerifierChain sequence.
const testVerifierResolverQualifier = "redeploy-test-verifier-resolver"

var _ ccvadapters.OnRampUpgrader = (*EVMOnRampUpgrader)(nil)

type oldOnRampConfig struct {
	StaticConfig        onramp.StaticConfig
	DynamicConfig       onramp.DynamicConfig
	DestChainConfigs    []uint64
	DestChainConfigArgs []onramp.DestChainConfigArgs
}

// EVMOnRampUpgrader upgrades an EVM OnRamp
type EVMOnRampUpgrader struct{}

func (a *EVMOnRampUpgrader) VerifyOnrampRequireUpgrade(e cldf.Environment, chainSelector uint64) error {
	chain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return fmt.Errorf("no EVM chain found for selector %d", chainSelector)
	}

	oldRef, err := a.findOldOnRamp(e.DataStore, chainSelector)
	if err != nil {
		return err
	}
	oldAddr := common.HexToAddress(oldRef.Address)

	legacyRef := oldRef
	legacyRef.Qualifier = legacyQualifier

	cfg, err := a.readOldOnRampConfigs(e.OperationsBundle, chain, chainSelector, oldAddr)
	if err != nil {
		return err
	}
	// Resolve every address the upgrade needs before mutating any state.
	rmnProxyAddr, err := resolveRMNProxy(e.DataStore, chainSelector)
	if err != nil {
		return err
	}
	if cfg.StaticConfig.RmnRemote == rmnProxyAddr {
		return fmt.Errorf(
			"OnRamp %s already uses RMNProxy %s as RmnRemote; nothing to upgrade", oldAddr.Hex(), rmnProxyAddr.Hex())
	}
	return nil
}

// DeployNewOnRamp implements [ccvadapters.OnRampUpgrader].
func (a *EVMOnRampUpgrader) DeployNewOnRamp(e cldf.Environment, chainSelector uint64) (ccvadapters.OnRampUpgradeResult, error) {
	chain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return ccvadapters.OnRampUpgradeResult{}, fmt.Errorf("no EVM chain found for selector %d", chainSelector)
	}

	oldRef, err := a.findOldOnRamp(e.DataStore, chainSelector)
	if err != nil {
		return ccvadapters.OnRampUpgradeResult{}, err
	}
	oldAddr := common.HexToAddress(oldRef.Address)

	legacyRef := oldRef
	legacyRef.Qualifier = legacyQualifier

	existingLegacy, err := a.LegacyOnRampRef(e, chainSelector)
	if err == nil && existingLegacy.Address != "" {
		legacyRef = existingLegacy
	}

	cfg, err := a.readOldOnRampConfigs(e.OperationsBundle, chain, chainSelector, oldAddr)
	if err != nil {
		return ccvadapters.OnRampUpgradeResult{}, err
	}
	rmnProxyAddr, err := resolveRMNProxy(e.DataStore, chainSelector)
	if err != nil {
		return ccvadapters.OnRampUpgradeResult{}, err
	}

	newRef, err := a.deployNewOnRamp(e, chain, chainSelector, cfg, rmnProxyAddr)
	if err != nil {
		return ccvadapters.OnRampUpgradeResult{}, err
	}
	newAddr := common.HexToAddress(newRef.Address)

	// Copy each dest chain config verbatim, including its current Router (already populated
	// in cfg.DestChainConfigArgs by readAllDestChainConfigs). Neither router is repointed at
	// the new OnRamp here — that only happens once the verifier jobs observe both OnRamps
	// (Phase 1.5), so this write can never affect live traffic.
	applyReport, err := cldf_ops.ExecuteOperation(e.OperationsBundle, onramp.ApplyDestChainConfigUpdates, chain, contract.FunctionInput[[]onramp.DestChainConfigArgs]{
		ChainSelector: chainSelector,
		Address:       newAddr,
		Args:          cfg.DestChainConfigArgs,
	})
	if err != nil {
		return ccvadapters.OnRampUpgradeResult{}, fmt.Errorf("apply dest chain configs to new OnRamp: %w", err)
	}
	writes := []contract.WriteOutput{applyReport.Output}

	dynWrite, err := a.applyDynamicConfig(e.OperationsBundle, chain, chainSelector, newAddr, cfg.DynamicConfig)
	if err != nil {
		return ccvadapters.OnRampUpgradeResult{}, err
	}
	if dynWrite.ChainSelector != 0 {
		writes = append(writes, dynWrite)
	}

	batchOp, err := contract.NewBatchOperationFromWrites(writes)
	if err != nil {
		return ccvadapters.OnRampUpgradeResult{}, fmt.Errorf("build batch op from writes: %w", err)
	}

	return ccvadapters.OnRampUpgradeResult{
		NewOnRampRef:    newRef,
		LegacyOnRampRef: legacyRef,
		BatchOps:        []mcms_types.BatchOperation{batchOp},
	}, nil
}

// DestChainSelectors implements [ccvadapters.OnRampUpgrader].
func (a *EVMOnRampUpgrader) DestChainSelectors(e cldf.Environment, chainSelector uint64) ([]uint64, error) {
	chain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return nil, fmt.Errorf("no EVM chain found for selector %d", chainSelector)
	}

	newRef, err := canonicalOnRampRef(e.DataStore, chainSelector)
	if err != nil {
		return nil, fmt.Errorf("find new OnRamp: %w", err)
	}

	selectors, _, err := readAllDestChainConfigs(e.OperationsBundle, chain, chainSelector, common.HexToAddress(newRef.Address))
	if err != nil {
		return nil, err
	}
	return selectors, nil
}

// ClassifyDestChains implements [ccvadapters.OnRampUpgrader].
func (a *EVMOnRampUpgrader) ClassifyDestChains(e cldf.Environment, chainSelector uint64) (ccvadapters.LaneClass, error) {
	chain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return ccvadapters.LaneClass{}, fmt.Errorf("no EVM chain found for selector %d", chainSelector)
	}

	// Classify off the legacy-qualified ref once it exists (Phase 1 has run): its dest
	// configs are frozen at their original per-dest routers and never change again, unlike
	// the canonical ref, which Phase 1 flips to the new OnRamp and whose dest configs get
	// rewritten by later phases (e.g. Phase 2 repoints ProdRouter dests to the TestRouter).
	// Reading canonical here would misclassify a ProdRouter dest as TestRouter dest once
	// Phase 2 has staged it. Before Phase 1 runs, no legacy ref exists yet, so canonical
	// (still the original onramp) is the only option and is correct at that point.
	classifyRef, err := a.LegacyOnRampRef(e, chainSelector)
	if err != nil {
		classifyRef, err = a.findOldOnRamp(e.DataStore, chainSelector)
		if err != nil {
			return ccvadapters.LaneClass{}, err
		}
	}
	classifyAddr := common.HexToAddress(classifyRef.Address)

	prodRouterAddr, err := resolveProdRouter(e.DataStore, chainSelector)
	if err != nil {
		return ccvadapters.LaneClass{}, err
	}
	testRouterAddr, err := resolveTestRouter(e.DataStore, chainSelector)
	if err != nil {
		return ccvadapters.LaneClass{}, err
	}

	selectors, args, err := readAllDestChainConfigs(e.OperationsBundle, chain, chainSelector, classifyAddr)
	if err != nil {
		return ccvadapters.LaneClass{}, err
	}

	var class ccvadapters.LaneClass
	for i, destSel := range selectors {
		switch args[i].Router {
		case prodRouterAddr:
			class.ProdRouterDests = append(class.ProdRouterDests, destSel)
		case testRouterAddr:
			class.TestRouterDests = append(class.TestRouterDests, destSel)
		default:
			return ccvadapters.LaneClass{}, fmt.Errorf(
				"dest chain %d is routed through neither the prod Router (%s) nor the TestRouter (%s): %s",
				destSel, prodRouterAddr.Hex(), testRouterAddr.Hex(), args[i].Router.Hex())
		}
	}
	return class, nil
}

// LegacyOnRampRef implements [ccvadapters.OnRampUpgrader].
func (a *EVMOnRampUpgrader) LegacyOnRampRef(e cldf.Environment, chainSelector uint64) (datastore.AddressRef, error) {
	ref, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		Type:      datastore.ContractType(onramp.ContractType),
		Version:   onramp.Version,
		Qualifier: legacyQualifier,
	}, chainSelector, toAddressRef)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("find legacy OnRamp on chain %d: %w", chainSelector, err)
	}
	return ref, nil
}

// VerifyPromotedToProdRouter implements [ccvadapters.OnRampUpgrader].
func (a *EVMOnRampUpgrader) VerifyPromotedToRouters(e cldf.Environment, chainSelector uint64, prodDestSelectors []uint64, testDestSelectors []uint64) error {
	newRef, err := canonicalOnRampRef(e.DataStore, chainSelector)
	if err != nil {
		return fmt.Errorf("find new OnRamp: %w", err)
	}
	routerAddr, err := resolveProdRouter(e.DataStore, chainSelector)
	if err != nil {
		return err
	}
	if err := a.verifyRouterRoutesThrough(e, chainSelector, prodDestSelectors, common.HexToAddress(newRef.Address), routerAddr); err != nil {
		return fmt.Errorf("%w (Phase 3 must execute first)", err)
	}

	routerAddr, err = resolveTestRouter(e.DataStore, chainSelector)
	if err != nil {
		return err
	}
	if err := a.verifyRouterRoutesThrough(e, chainSelector, testDestSelectors, common.HexToAddress(newRef.Address), routerAddr); err != nil {
		return fmt.Errorf("%w (Phase 3 must execute first)", err)
	}

	return nil
}

// VerifyLegacyOnRampOnProdRouter implements [ccvadapters.OnRampUpgrader].
func (a *EVMOnRampUpgrader) VerifyLegacyOnRampOnProdRouter(e cldf.Environment, chainSelector uint64, destSelectors []uint64) error {
	legacyRef, err := a.LegacyOnRampRef(e, chainSelector)
	if err != nil {
		return err
	}
	routerAddr, err := resolveProdRouter(e.DataStore, chainSelector)
	if err != nil {
		return err
	}
	if err := a.verifyRouterRoutesThrough(e, chainSelector, destSelectors, common.HexToAddress(legacyRef.Address), routerAddr); err != nil {
		return fmt.Errorf("%w (the prod Router must still route through the legacy OnRamp before promotion)", err)
	}
	return nil
}

// filterBySelectors returns the subset of (selectors, configs) whose selector appears in want,
// preserving want's order. Any selector in want that has no matching config is an error.
func filterBySelectors(selectors []uint64, configs []onramp.DestChainConfigArgs, want []uint64) ([]onramp.DestChainConfigArgs, error) {
	if len(selectors) != len(configs) {
		return nil, fmt.Errorf("mismatched selectors/configs lengths: selectors=%d configs=%d", len(selectors), len(configs))
	}
	bySelector := make(map[uint64]onramp.DestChainConfigArgs, len(selectors))
	for i, sel := range selectors {
		if _, exists := bySelector[sel]; exists {
			return nil, fmt.Errorf("duplicate dest chain selector %d in OnRamp config list", sel)
		}
		bySelector[sel] = configs[i]
	}
	out := make([]onramp.DestChainConfigArgs, 0, len(want))
	for _, sel := range want {
		cfg, ok := bySelector[sel]
		if !ok {
			return nil, fmt.Errorf("dest chain %d has no config on the OnRamp", sel)
		}
		out = append(out, cfg)
	}
	return out, nil
}

// verifyProdRouterRoutesThrough returns an error unless the prod Router routes every
// selector in destSelectors through wantOnRamp.
func (a *EVMOnRampUpgrader) verifyRouterRoutesThrough(e cldf.Environment, chainSelector uint64, destSelectors []uint64, wantOnRamp common.Address, routerAddr common.Address) error {
	chain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return fmt.Errorf("no EVM chain found for selector %d", chainSelector)
	}

	var mismatched []uint64
	for _, destSel := range destSelectors {
		report, err := cldf_ops.ExecuteOperation(e.OperationsBundle, router.GetOnRamp, chain, contract.FunctionInput[uint64]{
			ChainSelector: chainSelector,
			Address:       routerAddr,
			Args:          destSel,
		})
		if err != nil {
			return fmt.Errorf("read prod Router OnRamp for dest chain %d: %w", destSel, err)
		}
		if report.Output != wantOnRamp {
			mismatched = append(mismatched, destSel)
		}
	}
	if len(mismatched) > 0 {
		return fmt.Errorf("prod Router does not route dest chains %v through %s", mismatched, wantOnRamp.Hex())
	}
	return nil
}

// VerifyNewOnRampOwner implements [ccvadapters.OnRampUpgrader].
func (a *EVMOnRampUpgrader) VerifyNewOnRampOwner(e cldf.Environment, chainSelector uint64, expectedOwner string) error {
	chain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return fmt.Errorf("no EVM chain found for selector %d", chainSelector)
	}

	newRef, err := canonicalOnRampRef(e.DataStore, chainSelector)
	if err != nil {
		return fmt.Errorf("find new OnRamp: %w", err)
	}

	report, err := cldf_ops.ExecuteOperation(e.OperationsBundle, onramp.Owner, chain, contract.FunctionInput[struct{}]{
		ChainSelector: chainSelector,
		Address:       common.HexToAddress(newRef.Address),
		Args:          struct{}{},
	})
	if err != nil {
		return fmt.Errorf("read new OnRamp owner: %w", err)
	}
	if want := common.HexToAddress(expectedOwner); report.Output != want {
		return fmt.Errorf("new OnRamp is owned by %s, expected %s (Phase 1's ownership transfer must execute first)",
			report.Output.Hex(), want.Hex())
	}
	return nil
}

func (a *EVMOnRampUpgrader) findOldOnRamp(ds datastore.DataStore, chainSelector uint64) (datastore.AddressRef, error) {
	ref, err := canonicalOnRampRef(ds, chainSelector)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("find old OnRamp on chain %d: %w", chainSelector, err)
	}
	return ref, nil
}

// canonicalOnRampRef resolves the canonical (unqualified) OnRamp ref for a chain.
func canonicalOnRampRef(ds datastore.DataStore, chainSelector uint64) (datastore.AddressRef, error) {
	return datastore_utils.FindAndFormatCanonicalRef(ds, datastore.AddressRef{
		Type:    datastore.ContractType(onramp.ContractType),
		Version: onramp.Version,
	}, chainSelector, toAddressRef)
}

func (a *EVMOnRampUpgrader) readOldOnRampConfigs(
	b cldf_ops.Bundle,
	chain cldf_evm.Chain,
	chainSelector uint64,
	oldAddr common.Address,
) (oldOnRampConfig, error) {
	staticCfg, err := readStaticConfig(b, chain, chainSelector, oldAddr)
	if err != nil {
		return oldOnRampConfig{}, err
	}
	dynamicCfg, err := readDynamicConfig(b, chain, chainSelector, oldAddr)
	if err != nil {
		return oldOnRampConfig{}, err
	}
	selectors, args, err := readAllDestChainConfigs(b, chain, chainSelector, oldAddr)
	if err != nil {
		return oldOnRampConfig{}, err
	}
	return oldOnRampConfig{
		StaticConfig:        staticCfg,
		DynamicConfig:       dynamicCfg,
		DestChainConfigs:    selectors,
		DestChainConfigArgs: args,
	}, nil
}

func (a *EVMOnRampUpgrader) deployNewOnRamp(
	e cldf.Environment,
	chain cldf_evm.Chain,
	chainSelector uint64,
	cfg oldOnRampConfig,
	rmnProxyAddr common.Address,
) (datastore.AddressRef, error) {
	existingAddrs := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(chainSelector))
	ref, err := contract.MaybeDeployContract(e.OperationsBundle, onramp.Deploy, chain, contract.DeployInput[onramp.ConstructorArgs]{
		TypeAndVersion: cldf.NewTypeAndVersion(onramp.ContractType, *onramp.Version),
		ChainSelector:  chainSelector,
		Qualifier:      ptr.To(upgradeQualifier),
		Args: onramp.ConstructorArgs{
			StaticConfig: onramp.StaticConfig{
				ChainSelector:         chainSelector,
				RmnRemote:             rmnProxyAddr,
				TokenAdminRegistry:    cfg.StaticConfig.TokenAdminRegistry,
				MaxUSDCentsPerMessage: cfg.StaticConfig.MaxUSDCentsPerMessage,
			},
			DynamicConfig: onramp.DynamicConfig{
				FeeQuoter:     cfg.DynamicConfig.FeeQuoter,
				FeeAggregator: cfg.DynamicConfig.FeeAggregator,
			},
		},
	}, existingAddrs)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("deploy new OnRamp: %w", err)
	}
	ref.Qualifier = ""
	return ref, nil
}

// applyDestChainConfigsWithRouter rewrites each dest chain config to route
// through routerAddr and applies them to the OnRamp at newAddr.
func (a *EVMOnRampUpgrader) applyDestChainConfigsWithRouter(
	e cldf.Environment,
	chain cldf_evm.Chain,
	chainSelector uint64,
	newAddr common.Address,
	destChainArgs []onramp.DestChainConfigArgs,
	routerAddr common.Address,
) (contract.WriteOutput, error) {
	routerArgs := make([]onramp.DestChainConfigArgs, len(destChainArgs))
	for i, arg := range destChainArgs {
		routerArgs[i] = arg
		routerArgs[i].Router = routerAddr
	}

	applyReport, err := cldf_ops.ExecuteOperation(e.OperationsBundle, onramp.ApplyDestChainConfigUpdates, chain, contract.FunctionInput[[]onramp.DestChainConfigArgs]{
		ChainSelector: chainSelector,
		Address:       newAddr,
		Args:          routerArgs,
	})
	if err != nil {
		return contract.WriteOutput{}, fmt.Errorf("apply dest chain configs to new OnRamp: %w", err)
	}
	return applyReport.Output, nil
}

func (a *EVMOnRampUpgrader) applyDynamicConfig(
	b cldf_ops.Bundle,
	chain cldf_evm.Chain,
	chainSelector uint64,
	newAddr common.Address,
	oldDynamicCfg onramp.DynamicConfig,
) (contract.WriteOutput, error) {
	desired := onramp.DynamicConfig{
		FeeQuoter:              oldDynamicCfg.FeeQuoter,
		ReentrancyGuardEntered: false,
		FeeAggregator:          oldDynamicCfg.FeeAggregator,
	}
	if oldDynamicCfg == desired {
		return contract.WriteOutput{}, nil
	}
	report, err := cldf_ops.ExecuteOperation(b, onramp.SetDynamicConfig, chain, contract.FunctionInput[onramp.DynamicConfig]{
		ChainSelector: chainSelector,
		Address:       newAddr,
		Args:          desired,
	})
	if err != nil {
		return contract.WriteOutput{}, fmt.Errorf("set dynamic config on new OnRamp: %w", err)
	}
	return report.Output, nil
}

func (a *EVMOnRampUpgrader) applyRouterUpdates(
	e cldf.Environment,
	chain cldf_evm.Chain,
	chainSelector uint64,
	newAddr common.Address,
	destChainSelectors []uint64,
	routerAddr common.Address,
) (contract.WriteOutput, error) {
	onRampUpdates := make([]router.OnRamp, 0, len(destChainSelectors))
	for _, sel := range destChainSelectors {
		onRampUpdates = append(onRampUpdates, router.OnRamp{
			DestChainSelector: sel,
			OnRamp:            newAddr,
		})
	}

	report, err := cldf_ops.ExecuteOperation(e.OperationsBundle, router.ApplyRampUpdates, chain, contract.FunctionInput[router.ApplyRampsUpdatesArgs]{
		ChainSelector: chainSelector,
		Address:       routerAddr,
		Args: router.ApplyRampsUpdatesArgs{
			OnRampUpdates: onRampUpdates,
		},
	})
	if err != nil {
		return contract.WriteOutput{}, fmt.Errorf("apply ramp updates on router %s: %w", routerAddr.Hex(), err)
	}
	return report.Output, nil
}

func readStaticConfig(b cldf_ops.Bundle, chain cldf_evm.Chain, chainSelector uint64, addr common.Address) (onramp.StaticConfig, error) {
	report, err := cldf_ops.ExecuteOperation(b, onramp.GetStaticConfig, chain, contract.FunctionInput[struct{}]{
		ChainSelector: chainSelector,
		Address:       addr,
		Args:          struct{}{},
	})
	if err != nil {
		return onramp.StaticConfig{}, fmt.Errorf("read static config from OnRamp: %w", err)
	}
	return report.Output, nil
}

func readDynamicConfig(b cldf_ops.Bundle, chain cldf_evm.Chain, chainSelector uint64, addr common.Address) (onramp.DynamicConfig, error) {
	report, err := cldf_ops.ExecuteOperation(b, onramp.GetDynamicConfig, chain, contract.FunctionInput[struct{}]{
		ChainSelector: chainSelector,
		Address:       addr,
		Args:          struct{}{},
	})
	if err != nil {
		return onramp.DynamicConfig{}, fmt.Errorf("read dynamic config from OnRamp: %w", err)
	}
	return report.Output, nil
}

func readAllDestChainConfigs(b cldf_ops.Bundle, chain cldf_evm.Chain, chainSelector uint64, addr common.Address) ([]uint64, []onramp.DestChainConfigArgs, error) {
	report, err := cldf_ops.ExecuteOperation(b, onramp.GetAllDestChainConfigs, chain, contract.FunctionInput[struct{}]{
		ChainSelector: chainSelector,
		Address:       addr,
		Args:          struct{}{},
	})
	if err != nil {
		return nil, nil, fmt.Errorf("read all dest chain configs: %w", err)
	}
	result := report.Output
	if len(result.Ret0) == 0 {
		if len(result.Ret1) != 0 {
			return nil, nil, fmt.Errorf("mismatched getAllDestChainConfigs result lengths on OnRamp %s: selectors=%d configs=%d", addr.Hex(), len(result.Ret0), len(result.Ret1))
		}
		return nil, nil, fmt.Errorf("no dest chain configs found on OnRamp %s", addr.Hex())
	}
	if len(result.Ret0) != len(result.Ret1) {
		return nil, nil, fmt.Errorf("mismatched getAllDestChainConfigs result lengths on OnRamp %s: selectors=%d configs=%d", addr.Hex(), len(result.Ret0), len(result.Ret1))
	}

	selectors := result.Ret0
	args := make([]onramp.DestChainConfigArgs, len(selectors))
	for i, cfg := range result.Ret1 {
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
	}
	return selectors, args, nil
}

func resolveRMNProxy(ds datastore.DataStore, chainSelector uint64) (common.Address, error) {
	addr, err := datastore_utils.FindAndFormatCanonicalRef(ds, datastore.AddressRef{
		Type:    datastore.ContractType(rmnproxyops.ContractType),
		Version: semver.MustParse("1.0.0"),
	}, chainSelector, evmds.ToEVMAddress)
	if err != nil {
		return common.Address{}, fmt.Errorf("resolve RMNProxy on chain %d: %w", chainSelector, err)
	}
	return addr, nil
}

func resolveTestRouter(ds datastore.DataStore, chainSelector uint64) (common.Address, error) {
	bytes, err := datastore_utils.FindAndFormatCanonicalRef(ds, datastore.AddressRef{
		Type:    datastore.ContractType(router.TestRouterContractType),
		Version: router.Version,
	}, chainSelector, evmds.ToEVMAddressBytes)
	if err != nil {
		return common.Address{}, fmt.Errorf("resolve TestRouter on chain %d: %w", chainSelector, err)
	}
	return common.BytesToAddress(bytes), nil
}

// PromoteOnrampToProdRouter implements [ccvadapters.OnRampUpgrader].
func (a *EVMOnRampUpgrader) PromoteOnrampToProdRouter(e cldf.Environment, chainSelector uint64, destSelectors []uint64) ([]mcms_types.BatchOperation, error) {
	prodRouterAddr, err := resolveProdRouter(e.DataStore, chainSelector)
	if err != nil {
		return nil, err
	}
	return a.promoteOnrampToRouter(e, chainSelector, destSelectors, prodRouterAddr)
}

// PromoteOnrampToTestRouter implements [ccvadapters.OnRampUpgrader].
func (a *EVMOnRampUpgrader) PromoteOnrampToTestRouter(e cldf.Environment, chainSelector uint64, destSelectors []uint64) ([]mcms_types.BatchOperation, error) {
	testRouterAddr, err := resolveTestRouter(e.DataStore, chainSelector)
	if err != nil {
		return nil, err
	}
	return a.promoteOnrampToRouter(e, chainSelector, destSelectors, testRouterAddr)
}

// promoteOnrampToRouter points destSelectors' dest chain configs and routerAddr at the new OnRamp.
func (a *EVMOnRampUpgrader) promoteOnrampToRouter(e cldf.Environment, chainSelector uint64, destSelectors []uint64, routerAddr common.Address) ([]mcms_types.BatchOperation, error) {
	chain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return nil, fmt.Errorf("no EVM chain found for selector %d", chainSelector)
	}
	if len(destSelectors) == 0 {
		return nil, nil
	}

	newRef, err := canonicalOnRampRef(e.DataStore, chainSelector)
	if err != nil {
		return nil, fmt.Errorf("find new OnRamp: %w", err)
	}
	newAddr := common.HexToAddress(newRef.Address)

	allSelectors, allConfigs, err := readAllDestChainConfigs(e.OperationsBundle, chain, chainSelector, newAddr)
	if err != nil {
		return nil, err
	}
	configs, err := filterBySelectors(allSelectors, allConfigs, destSelectors)
	if err != nil {
		return nil, err
	}

	var writes []contract.WriteOutput

	// Update dest chain configs to use routerAddr.
	onRampWrite, err := a.applyDestChainConfigsWithRouter(e, chain, chainSelector, newAddr, configs, routerAddr)
	if err != nil {
		return nil, err
	}
	writes = append(writes, onRampWrite)

	// Wire routerAddr to route through the new OnRamp.
	routerWrite, err := a.applyRouterUpdates(e, chain, chainSelector, newAddr, destSelectors, routerAddr)
	if err != nil {
		return nil, err
	}
	writes = append(writes, routerWrite)

	batchOp, err := contract.NewBatchOperationFromWrites(writes)
	if err != nil {
		return nil, fmt.Errorf("build batch op: %w", err)
	}
	if len(batchOp.Transactions) > 0 {
		return []mcms_types.BatchOperation{batchOp}, nil
	}
	return nil, nil
}

func resolveProdRouter(ds datastore.DataStore, chainSelector uint64) (common.Address, error) {
	bytes, err := datastore_utils.FindAndFormatCanonicalRef(ds, datastore.AddressRef{
		Type:    datastore.ContractType(router.ContractType),
		Version: router.Version,
	}, chainSelector, evmds.ToEVMAddressBytes)
	if err != nil {
		return common.Address{}, fmt.Errorf("resolve prod Router on chain %d: %w", chainSelector, err)
	}
	return common.BytesToAddress(bytes), nil
}

func toAddressRef(ref datastore.AddressRef) (datastore.AddressRef, error) {
	return ref, nil
}
