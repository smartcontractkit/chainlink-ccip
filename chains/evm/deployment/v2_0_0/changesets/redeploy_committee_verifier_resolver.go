package changesets

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/mock_receiver"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/mock_receiver_v2"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/versioned_verifier_resolver"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	cs_changesets "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	contract_utils "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

const (
	// defaultCommitteeQualifier is the qualifier of the primary committee. It is
	// also the CREATE2 salt of its resolver.
	defaultCommitteeQualifier = "default"
	// legacyOnRampQualifier is the qualifier that UpgradeOnrampPhase1 gives to the
	// pre-upgrade OnRamp.
	legacyOnRampQualifier = "legacy"
	// mockReceiverRedeployQualifier is a temporary qualifier. It stops
	// MaybeDeployContract from reusing the current MockReceiver ref.
	mockReceiverRedeployQualifier = "redeploy"
)

// RedeployCommitteeVerifierResolverCfg configures RedeployCommitteeVerifierResolver.
type RedeployCommitteeVerifierResolverCfg struct {
	// ChainSelectors are the chains to repair.
	ChainSelectors []uint64
	// CanonicalCREATE2Factory is the expected address of the CREATE2 factory in
	// this environment. The changeset fails if any chain has a factory at a
	// different address.
	//
	// The CREATE2 address of the resolver is a function of the factory address,
	// because ComputeAddress is a call on the factory. The factory itself is
	// deployed with plain CREATE, so its address depends on the deployer and the
	// nonce. A chain that got the factory at a different address therefore has
	// every CREATE2 contract at a different address, the resolver included. Such
	// a chain needs a canonical factory before this changeset can repair it.
	CanonicalCREATE2Factory common.Address
	// CommitteeQualifier is the qualifier of the committee. It defaults to "default".
	CommitteeQualifier string
	// DisableTransferOwnership keeps the new resolver on the deployer key. Use it
	// only in tests and in environments that have no timelock.
	DisableTransferOwnership bool
}

// Qualifier returns the committee qualifier, or the default value.
func (c RedeployCommitteeVerifierResolverCfg) Qualifier() string {
	if c.CommitteeQualifier == "" {
		return defaultCommitteeQualifier
	}
	return c.CommitteeQualifier
}

// RedeployCommitteeVerifierResolver redeploys the CommitteeVerifierResolver at its
// canonical CREATE2 address.
//
// Some resolvers were deployed with plain CREATE. They are therefore not at the
// canonical deterministic address. This changeset repairs those chains.
//
// The changeset does this work for each chain:
//
//  1. It redeploys the resolver through the CREATE2 factory, with the committee
//     qualifier as the salt. It overwrites the address of the current ref.
//  2. It copies the inbound and outbound implementation maps of the old resolver
//     onto the new one. These writes use the deployer key, because the deployer
//     owns the new resolver.
//  3. It redeploys the MockReceiver. The verifier list of the MockReceiver is a
//     constructor argument and has no setter.
//  4. It collects the OffRamp and OnRamp writes that replace the old resolver
//     address. Both the canonical and the legacy OnRamp are updated. The ramps
//     belong to the timelock, so these writes become an MCMS proposal.
//  5. It transfers ownership of the new resolver to the timelock.
//
// The changeset takes a chain list and the canonical CREATE2 factory address. The
// resolver is chain-local, so the old resolver, the ramp configs, and the CREATE2
// factory ref give every other value.
//
// A preflight check rejects any chain whose CREATE2 factory is not at the
// canonical address. See CanonicalCREATE2Factory for the reason.
func RedeployCommitteeVerifierResolver(
	mcmsRegistry *cs_changesets.MCMSReaderRegistry,
	transferOwnershipReg *deploy.TransferOwnershipAdapterRegistry,
) cldf_deployment.ChangeSetV2[cs_changesets.WithMCMS[RedeployCommitteeVerifierResolverCfg]] {
	return cldf_deployment.CreateChangeSet(
		makeApplyRedeployCommitteeVerifierResolver(mcmsRegistry, transferOwnershipReg),
		makeVerifyRedeployCommitteeVerifierResolver(),
	)
}

func makeVerifyRedeployCommitteeVerifierResolver() func(cldf_deployment.Environment, cs_changesets.WithMCMS[RedeployCommitteeVerifierResolverCfg]) error {
	return func(e cldf_deployment.Environment, cfg cs_changesets.WithMCMS[RedeployCommitteeVerifierResolverCfg]) error {
		if len(cfg.Cfg.ChainSelectors) == 0 {
			return fmt.Errorf("at least one chain must be configured")
		}
		if cfg.Cfg.CanonicalCREATE2Factory == (common.Address{}) {
			return fmt.Errorf("CanonicalCREATE2Factory is required")
		}
		qualifier := cfg.Cfg.Qualifier()
		for _, chainSel := range cfg.Cfg.ChainSelectors {
			if _, ok := e.BlockChains.EVMChains()[chainSel]; !ok {
				return fmt.Errorf("chain %d not found in environment", chainSel)
			}
			refs := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(chainSel))
			if _, err := findResolverRef(refs, qualifier); err != nil {
				return fmt.Errorf("chain %d: %w", chainSel, err)
			}
			factory, err := findCREATE2Factory(refs)
			if err != nil {
				return fmt.Errorf("chain %d: %w", chainSel, err)
			}
			if factory != cfg.Cfg.CanonicalCREATE2Factory {
				return fmt.Errorf(
					"chain %d: CREATE2Factory is at %s, but the canonical factory is at %s. "+
						"Every CREATE2 address on this chain, the resolver included, derives from the factory address. "+
						"Deploy the canonical factory on this chain and update the datastore ref before you run this changeset",
					chainSel, factory.Hex(), cfg.Cfg.CanonicalCREATE2Factory.Hex())
			}
		}
		return nil
	}
}

func makeApplyRedeployCommitteeVerifierResolver(
	mcmsRegistry *cs_changesets.MCMSReaderRegistry,
	transferOwnershipReg *deploy.TransferOwnershipAdapterRegistry,
) func(cldf_deployment.Environment, cs_changesets.WithMCMS[RedeployCommitteeVerifierResolverCfg]) (cldf_deployment.ChangesetOutput, error) {
	return func(e cldf_deployment.Environment, cfg cs_changesets.WithMCMS[RedeployCommitteeVerifierResolverCfg]) (cldf_deployment.ChangesetOutput, error) {
		if err := makeVerifyRedeployCommitteeVerifierResolver()(e, cfg); err != nil {
			return cldf_deployment.ChangesetOutput{}, err
		}

		newDS := datastore.NewMemoryDataStore()
		var reports []cldf_ops.Report[any, any]
		var batchOps []mcms_types.BatchOperation
		qualifier := cfg.Cfg.Qualifier()

		for _, chainSel := range cfg.Cfg.ChainSelectors {
			chain, ok := e.BlockChains.EVMChains()[chainSel]
			if !ok {
				return cldf_deployment.ChangesetOutput{}, fmt.Errorf("chain %d not found in environment", chainSel)
			}
			refs := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(chainSel))

			oldResolverRef, err := findResolverRef(refs, qualifier)
			if err != nil {
				return cldf_deployment.ChangesetOutput{}, fmt.Errorf("chain %d: %w", chainSel, err)
			}
			oldResolver := common.HexToAddress(oldResolverRef.Address)

			create2Factory, err := findCREATE2Factory(refs)
			if err != nil {
				return cldf_deployment.ChangesetOutput{}, fmt.Errorf("chain %d: %w", chainSel, err)
			}

			// Step 1: redeploy the resolver at its canonical CREATE2 address.
			deployReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.DeployVerifierResolverViaCREATE2, chain, sequences.DeployVerifierResolverViaCREATE2Input{
				CREATE2Factory: create2Factory,
				ChainSelector:  chainSel,
				Type:           datastore.ContractType(versioned_verifier_resolver.CommitteeVerifierResolverType),
				Version:        versioned_verifier_resolver.Version,
				Qualifier:      qualifier,
			})
			if err != nil {
				return cldf_deployment.ChangesetOutput{}, fmt.Errorf("chain %d: failed to redeploy resolver: %w", chainSel, err)
			}
			reports = append(reports, deployReport.ExecutionReports...)
			if len(deployReport.Output.Addresses) != 1 {
				return cldf_deployment.ChangesetOutput{}, fmt.Errorf("chain %d: expected 1 resolver address, got %d", chainSel, len(deployReport.Output.Addresses))
			}
			newResolverRef := deployReport.Output.Addresses[0]
			newResolver := common.HexToAddress(newResolverRef.Address)
			if newResolver == oldResolver {
				return cldf_deployment.ChangesetOutput{}, fmt.Errorf(
					"chain %d: resolver is already at the canonical CREATE2 address %s, nothing to redeploy",
					chainSel, oldResolver.Hex())
			}

			// The ref key is the same, so this upsert replaces the old address.
			if err := newDS.Addresses().Add(newResolverRef); err != nil {
				return cldf_deployment.ChangesetOutput{}, fmt.Errorf("chain %d: failed to add resolver ref: %w", chainSel, err)
			}
			writes := deployReport.Output.Writes

			// Step 2: configure the new resolver. The deployer owns it, so these
			// writes execute directly and stay out of the proposal.
			configWrites, err := configureNewResolver(e, chain.Selector, oldResolver, newResolver)
			if err != nil {
				return cldf_deployment.ChangesetOutput{}, fmt.Errorf("chain %d: %w", chainSel, err)
			}
			writes = append(writes, configWrites...)

			// Step 3: redeploy the MockReceiver with the new resolver.
			mockRefs, mockWrites, err := redeployMockReceivers(e, chainSel, refs, newResolver)
			if err != nil {
				return cldf_deployment.ChangesetOutput{}, fmt.Errorf("chain %d: %w", chainSel, err)
			}
			for _, ref := range mockRefs {
				if err := newDS.Addresses().Add(ref); err != nil {
					return cldf_deployment.ChangesetOutput{}, fmt.Errorf("chain %d: failed to add MockReceiver ref: %w", chainSel, err)
				}
			}
			writes = append(writes, mockWrites...)

			// Step 4: repoint the ramps. These writes need the proposal.
			rampWrites, err := repointRamps(e, chainSel, refs, oldResolver, newResolver)
			if err != nil {
				return cldf_deployment.ChangesetOutput{}, fmt.Errorf("chain %d: %w", chainSel, err)
			}
			writes = append(writes, rampWrites...)

			if len(writes) > 0 {
				batchOp, err := contract_utils.NewBatchOperationFromWrites(writes)
				if err != nil {
					return cldf_deployment.ChangesetOutput{}, fmt.Errorf("chain %d: failed to build batch operation: %w", chainSel, err)
				}
				if len(batchOp.Transactions) > 0 {
					batchOps = append(batchOps, batchOp)
				}
			}

			// Step 5: give ownership of the new resolver to the timelock.
			if !cfg.Cfg.DisableTransferOwnership {
				ownershipOps, ownershipReports, err := deploy.TransferToTimelock(
					chainSel, e, cfg.MCMS, []datastore.AddressRef{newResolverRef},
					mcmsRegistry, transferOwnershipReg,
				)
				if err != nil {
					return cldf_deployment.ChangesetOutput{}, fmt.Errorf("chain %d: transfer resolver ownership: %w", chainSel, err)
				}
				reports = append(reports, ownershipReports...)
				batchOps = append(batchOps, ownershipOps...)
			}
		}

		return cs_changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithDataStore(newDS).
			WithSingleBatchOpPerChain(batchOps).
			Build(cfg.MCMS)
	}
}

// configureNewResolver copies the implementation maps of the old resolver onto the
// new one, in both directions. A redeploy must produce an exact replica, so the
// old resolver is the only source of truth for the lane set.
//
// The new resolver starts empty. It must hold every entry before any ramp points
// at it, because the OnRamp reverts with DestinationChainNotSupportedByCCV if
// getOutboundImplementation returns zero.
func configureNewResolver(
	e cldf_deployment.Environment,
	chainSel uint64,
	oldResolver common.Address,
	newResolver common.Address,
) ([]contract_utils.WriteOutput, error) {
	chain := e.BlockChains.EVMChains()[chainSel]
	var writes []contract_utils.WriteOutput

	outbound, err := cldf_ops.ExecuteOperation(e.OperationsBundle, versioned_verifier_resolver.GetAllOutboundImplementations, chain, contract_utils.FunctionInput[any]{
		ChainSelector: chainSel,
		Address:       oldResolver,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to read outbound implementations of the old resolver: %w", err)
	}
	if len(outbound.Output) > 0 {
		report, err := cldf_ops.ExecuteOperation(e.OperationsBundle, versioned_verifier_resolver.ApplyOutboundImplementationUpdates, chain, contract_utils.FunctionInput[[]versioned_verifier_resolver.OutboundImplementationArgs]{
			ChainSelector: chainSel,
			Address:       newResolver,
			Args:          outbound.Output,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to apply outbound implementation updates: %w", err)
		}
		writes = append(writes, report.Output)
	}

	inbound, err := cldf_ops.ExecuteOperation(e.OperationsBundle, versioned_verifier_resolver.GetAllInboundImplementations, chain, contract_utils.FunctionInput[any]{
		ChainSelector: chainSel,
		Address:       oldResolver,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to read inbound implementations of the old resolver: %w", err)
	}
	if len(inbound.Output) > 0 {
		report, err := cldf_ops.ExecuteOperation(e.OperationsBundle, versioned_verifier_resolver.ApplyInboundImplementationUpdates, chain, contract_utils.FunctionInput[[]versioned_verifier_resolver.InboundImplementationArgs]{
			ChainSelector: chainSel,
			Address:       newResolver,
			Args:          inbound.Output,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to apply inbound implementation updates: %w", err)
		}
		writes = append(writes, report.Output)
	}

	return writes, nil
}

// redeployMockReceivers redeploys every MockReceiver on the chain with the new
// resolver as its required verifier. The verifier list is a constructor argument
// and has no setter, so an update in place is not possible.
//
// The deployment uses a temporary qualifier, because MaybeDeployContract reuses
// an existing ref that has the same type, version and qualifier. The returned ref
// carries the original qualifier, so the upsert replaces the old row.
func redeployMockReceivers(
	e cldf_deployment.Environment,
	chainSel uint64,
	refs []datastore.AddressRef,
	newResolver common.Address,
) ([]datastore.AddressRef, []contract_utils.WriteOutput, error) {
	chain := e.BlockChains.EVMChains()[chainSel]
	var newRefs []datastore.AddressRef
	var writes []contract_utils.WriteOutput

	for _, ref := range refs {
		if ref.Type != datastore.ContractType(mock_receiver.ContractType) {
			continue
		}

		current, err := cldf_ops.ExecuteOperation(e.OperationsBundle, mock_receiver_v2.GetCCVsAndFinalityConfig, chain, contract_utils.FunctionInput[mock_receiver_v2.GetCCVsAndFinalityConfigArgs]{
			ChainSelector: chainSel,
			Address:       common.HexToAddress(ref.Address),
			Args: mock_receiver_v2.GetCCVsAndFinalityConfigArgs{
				Arg0: chainSel,
				Arg1: []byte{},
			},
		})
		if err != nil {
			return nil, nil, fmt.Errorf("failed to read MockReceiver %s config: %w", ref.Address, err)
		}

		tempQualifier := mockReceiverRedeployQualifier
		if ref.Qualifier != "" {
			tempQualifier = ref.Qualifier + "-" + mockReceiverRedeployQualifier
		}
		deployReport, err := cldf_ops.ExecuteOperation(e.OperationsBundle, mock_receiver.Deploy, chain, contract_utils.DeployInput[mock_receiver.ConstructorArgs]{
			TypeAndVersion: cldf_deployment.NewTypeAndVersion(mock_receiver.ContractType, *ref.Version),
			ChainSelector:  chainSel,
			Args: mock_receiver.ConstructorArgs{
				Required:  []common.Address{newResolver},
				Optional:  current.Output.OptionalVerifiers,
				Threshold: current.Output.Threshold,
			},
			Qualifier: &tempQualifier,
		})
		if err != nil {
			return nil, nil, fmt.Errorf("failed to redeploy MockReceiver %s: %w", ref.Address, err)
		}

		newRef := deployReport.Output
		// Restore the original qualifier so the upsert replaces the old row.
		newRef.Qualifier = ref.Qualifier
		newRefs = append(newRefs, newRef)

		if current.Output.AllowedFinalityConfig != ([4]byte{}) {
			setReport, err := cldf_ops.ExecuteOperation(e.OperationsBundle, mock_receiver_v2.SetAllowedFinalityConfig, chain, contract_utils.FunctionInput[[4]byte]{
				ChainSelector: chainSel,
				Address:       common.HexToAddress(newRef.Address),
				Args:          current.Output.AllowedFinalityConfig,
			})
			if err != nil {
				return nil, nil, fmt.Errorf("failed to set finality config on MockReceiver %s: %w", newRef.Address, err)
			}
			writes = append(writes, setReport.Output)
		}
	}

	return newRefs, writes, nil
}

// repointRamps replaces the old resolver address with the new one on the OffRamp
// and on both OnRamps.
//
// The resolver is chain-local. Every source config of the local OffRamp and every
// dest config of each local OnRamp can name it, so the function rewrites each
// config that does. A config that does not name the old resolver produces no
// write, which keeps the operation idempotent.
//
// The legacy OnRamp must also be updated. It still serves traffic until Phase 3 of
// the OnRamp upgrade runbook.
func repointRamps(
	e cldf_deployment.Environment,
	chainSel uint64,
	refs []datastore.AddressRef,
	oldResolver common.Address,
	newResolver common.Address,
) ([]contract_utils.WriteOutput, error) {
	chain := e.BlockChains.EVMChains()[chainSel]
	var writes []contract_utils.WriteOutput

	// OffRamp source chain configs.
	offRampRef := findRef(refs, datastore.ContractType(offramp.ContractType), "")
	if offRampRef != nil {
		offRampAddr := common.HexToAddress(offRampRef.Address)
		current, err := cldf_ops.ExecuteOperation(e.OperationsBundle, offramp.GetAllSourceChainConfigs, chain, contract_utils.FunctionInput[struct{}]{
			ChainSelector: chainSel,
			Address:       offRampAddr,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to read OffRamp source chain configs: %w", err)
		}
		var sourceArgs []offramp.SourceChainConfigArgs
		for i, sourceSel := range current.Output.SourceChainSelectors {
			cfg := current.Output.SourceChainConfigs[i]
			defaultCCVs, changedDefault := replaceAddress(cfg.DefaultCCVs, oldResolver, newResolver)
			laneCCVs, changedLane := replaceAddress(cfg.LaneMandatedCCVs, oldResolver, newResolver)
			if !changedDefault && !changedLane {
				continue
			}
			sourceArgs = append(sourceArgs, offramp.SourceChainConfigArgs{
				Router:              cfg.Router,
				SourceChainSelector: sourceSel,
				IsEnabled:           cfg.IsEnabled,
				OnRamps:             cfg.OnRamps,
				DefaultCCVs:         defaultCCVs,
				LaneMandatedCCVs:    laneCCVs,
			})
		}
		if len(sourceArgs) > 0 {
			report, err := cldf_ops.ExecuteOperation(e.OperationsBundle, offramp.ApplySourceChainConfigUpdates, chain, contract_utils.FunctionInput[[]offramp.SourceChainConfigArgs]{
				ChainSelector: chainSel,
				Address:       offRampAddr,
				Args:          sourceArgs,
			})
			if err != nil {
				return nil, fmt.Errorf("failed to apply OffRamp source chain config updates: %w", err)
			}
			writes = append(writes, report.Output)
		}
	}

	// Both OnRamps. The canonical ref has an empty qualifier. The legacy ref only
	// exists after Phase 1 of the OnRamp upgrade.
	for _, onRampQualifier := range []string{"", legacyOnRampQualifier} {
		onRampRef := findRef(refs, datastore.ContractType(onramp.ContractType), onRampQualifier)
		if onRampRef == nil {
			continue
		}
		onRampAddr := common.HexToAddress(onRampRef.Address)
		current, err := cldf_ops.ExecuteOperation(e.OperationsBundle, onramp.GetAllDestChainConfigs, chain, contract_utils.FunctionInput[struct{}]{
			ChainSelector: chainSel,
			Address:       onRampAddr,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to read OnRamp %s dest chain configs: %w", onRampRef.Address, err)
		}
		var destArgs []onramp.DestChainConfigArgs
		for i, destSel := range current.Output.Ret0 {
			cfg := current.Output.Ret1[i]
			defaultCCVs, changedDefault := replaceAddress(cfg.DefaultCCVs, oldResolver, newResolver)
			laneCCVs, changedLane := replaceAddress(cfg.LaneMandatedCCVs, oldResolver, newResolver)
			if !changedDefault && !changedLane {
				continue
			}
			destArgs = append(destArgs, onramp.DestChainConfigArgs{
				DestChainSelector:         destSel,
				Router:                    cfg.Router,
				AddressBytesLength:        cfg.AddressBytesLength,
				TokenReceiverAllowed:      cfg.TokenReceiverAllowed,
				MessageNetworkFeeUSDCents: cfg.MessageNetworkFeeUSDCents,
				TokenNetworkFeeUSDCents:   cfg.TokenNetworkFeeUSDCents,
				BaseExecutionGasCost:      cfg.BaseExecutionGasCost,
				DefaultCCVs:               defaultCCVs,
				LaneMandatedCCVs:          laneCCVs,
				DefaultExecutor:           cfg.DefaultExecutor,
				OffRamp:                   cfg.OffRamp,
			})
		}
		if len(destArgs) > 0 {
			report, err := cldf_ops.ExecuteOperation(e.OperationsBundle, onramp.ApplyDestChainConfigUpdates, chain, contract_utils.FunctionInput[[]onramp.DestChainConfigArgs]{
				ChainSelector: chainSel,
				Address:       onRampAddr,
				Args:          destArgs,
			})
			if err != nil {
				return nil, fmt.Errorf("failed to apply OnRamp %s dest chain config updates: %w", onRampRef.Address, err)
			}
			writes = append(writes, report.Output)
		}
	}

	return writes, nil
}

// replaceAddress returns a copy of addrs with old replaced by new, and reports
// whether it made a change.
func replaceAddress(addrs []common.Address, old, new common.Address) ([]common.Address, bool) {
	out := make([]common.Address, len(addrs))
	copy(out, addrs)
	changed := false
	for i, addr := range out {
		if addr == old {
			out[i] = new
			changed = true
		}
	}
	return out, changed
}

// findRef returns the first ref of the given type and qualifier, or nil.
func findRef(refs []datastore.AddressRef, contractType datastore.ContractType, qualifier string) *datastore.AddressRef {
	for i := range refs {
		if refs[i].Type == contractType && refs[i].Qualifier == qualifier {
			return &refs[i]
		}
	}
	return nil
}

// findCREATE2Factory returns the address of the CREATE2 factory of the chain. The
// factory is a canonical contract, so its ref has an empty qualifier.
func findCREATE2Factory(refs []datastore.AddressRef) (common.Address, error) {
	ref := findRef(refs, datastore.ContractType(create2_factory.ContractType), "")
	if ref == nil {
		return common.Address{}, fmt.Errorf("no CREATE2Factory found")
	}
	return common.HexToAddress(ref.Address), nil
}

func findResolverRef(refs []datastore.AddressRef, qualifier string) (datastore.AddressRef, error) {
	ref := findRef(refs, datastore.ContractType(versioned_verifier_resolver.CommitteeVerifierResolverType), qualifier)
	if ref == nil {
		return datastore.AddressRef{}, fmt.Errorf("no CommitteeVerifierResolver with qualifier %q found", qualifier)
	}
	return *ref, nil
}
