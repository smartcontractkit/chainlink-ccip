# Templates

Copy these and adapt them. Before using one, check the exact signatures in the source files named in `deployment/docs/interfaces.md`, because the interfaces evolve.

## Registration: v1 (1.6)

```go
// <family>/deployment/v1_0_0/adapters/init.go: version-agnostic, keyed by family
func init() {
    fam := chain_selectors.FamilyMyChain
    changesets.GetRegistry().RegisterMCMSReader(fam, &MCMSReader{})
    deploy.GetTransferOwnershipRegistry().RegisterAdapter(fam, utils.Version_1_0_0, &TransferOwnershipAdapter{})
    deploy.GetAddressNormalizerRegistry().RegisterAddressNormalizer(fam, &AddressNormalizer{})
    fees.GetRegistry().RegisterFeeResolver(fam, &FeeResolver{})
    tokens.GetTokenAdapterRegistry().RegisterTokenRefResolver(fam, &TokenBase{})
    tokens.GetTokenAdapterRegistry().RegisterTokenAdminRegistryManager(fam, &TokenBase{}) // or RegisterTokenAdminRegistryReader
}

// <family>/deployment/v1_6_0/sequences/adapter.go
func init() {
    fam, v := chain_selectors.FamilyMyChain, utils.Version_1_6_0
    lanes.GetLaneAdapterRegistry().RegisterLaneAdapter(fam, v, &Adapter{}) // also records GetChainFamilySelector()
    deploy.GetRegistry().RegisterDeployer(fam, v, &Adapter{})
    deploy.GetTransferOwnershipRegistry().RegisterAdapter(fam, v, &Adapter{})
    tokens.GetTokenAdapterRegistry().RegisterTokenAdapter(fam, v, &Adapter{}) // once per pool version
}

// <family>/deployment/v1_6_0/adapters/init.go
func init() {
    fam, v := chain_selectors.FamilyMyChain, utils.Version_1_6_0
    fastcurse.GetCurseRegistry().RegisterNewCurse(fastcurse.CurseRegistryInput{
        CursingFamily: fam, CursingVersion: v, CurseAdapter: NewCurseAdapter(), CurseSubjectAdapter: NewCurseAdapter(),
    })
    fees.GetRegistry().RegisterFeeAdapter(fam, v, NewFeesAdapter())
    fees.GetFeeAggregatorRegistry().RegisterFeeAggregatorAdapter(fam, v, NewFeeAggregatorAdapter())
    testadapters.GetTestAdapterRegistry().RegisterTestAdapter(fam, v, NewTestAdapter) // usually in v1_6_0/testadapter
}
```

## Registration: v2 (2.0)

```go
// <family>/deployment/v2_0_0/adapters/init.go
func init() {
    fam, v := chain_selectors.FamilyMyChain, utils.Version_2_0_0
    // v2 chain API (keyed by family)
    v2adapters.GetDeployChainContractsRegistry().Register(fam, &DeployChainContractsAdapter{})
    v2adapters.GetChainFamilyRegistry().RegisterChainFamily(fam, &ChainFamilyAdapter{})
    v2adapters.GetCommitteeVerifierContractRegistry().Register(fam, &CommitteeVerifierContractAdapter{})
    v2adapters.GetDeployChainContractsRegistry().RegisterLaneVersionResolver(fam, &LaneVersionResolver{}) // recommended
    v2adapters.GetOnRampUpgraderRegistry().Register(fam, &OnRampUpgrader{})                           // optional
    v2adapters.GetTestVerifierChainRegistry().Register(fam, &TestVerifierChainAdapter{})              // optional
    // shared APIs
    changesets.GetRegistry().RegisterMCMSReader(fam, &ChainFamilyAdapter{})
    deploy.GetTransferOwnershipRegistry().RegisterAdapter(fam, v, &ChainFamilyAdapter{})
    tokens.GetTokenAdapterRegistry().RegisterTokenAdapter(fam, v, &TokenAdapter{})
    fees.GetRegistry().RegisterFeeAdapter(fam, v, NewFeesAdapter())
    fees.GetFeeAggregatorRegistry().RegisterFeeAggregatorAdapter(fam, v, NewFeeAggregatorAdapter())
    fastcurse.GetCurseRegistry().RegisterNewCurse(fastcurse.CurseRegistryInput{ /* RMN version */ })
    utils.RegisterChainFamilySelector(fam, [4]byte{ /* family selector */ }) // only if no LaneAdapter registers it
}
```

## Compile-time interface checks

```go
var (
    _ v2adapters.ChainFamily                 = (*ChainFamilyAdapter)(nil)
    _ v2adapters.OffRampSourceOnRampSetter   = (*ChainFamilyAdapter)(nil) // optional, found by type assertion
    _ v2adapters.DeployChainContractsAdapter = (*DeployChainContractsAdapter)(nil)
    _ tokens.TokenAdapter                    = (*TokenAdapter)(nil)
    _ tokens.TokenFeeAdapter                 = (*TokenAdapter)(nil) // optional, found by type assertion
)
```

## Operation (non-EVM) with deployer-or-MCMS branching

```go
var SetThing = operations.NewOperation(
    "my-contract:set-thing",
    Version,
    "Sets thing on MyContract",
    func(b operations.Bundle, chain cldf_mychain.Chain, in SetThingInput) (sequences.OnChainOutput, error) {
        ix, err := buildSetThingIx(in)
        if err != nil {
            return sequences.OnChainOutput{}, err
        }
        if in.Authority != chain.DeployerKey.PublicKey() {
            batch, err := utils.BuildMCMSBatchOperation(chain.Selector, []Instruction{ix}, in.Contract.String(), ContractType)
            return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{batch}}, err
        }
        return sequences.OnChainOutput{}, chain.Confirm([]Instruction{ix})
    },
)
```

EVM uses `contract.NewWrite(contract.WriteParams{..., IsAllowedCaller: contract.OnlyOwner[...]})` from `chainlink-deployments-framework/chain/evm/operations/contract`, which branches automatically.

## Sequence

```go
var ConfigureThing = operations.NewSequence(
    "configure-thing",
    semver.MustParse("2.0.0"),
    "Configures thing on one chain",
    func(b operations.Bundle, chains cldf_chain.BlockChains, in ConfigureThingInput) (sequences.OnChainOutput, error) {
        var out sequences.OnChainOutput
        chain, ok := chains.MyChains()[in.ChainSelector]
        if !ok {
            return out, fmt.Errorf("chain %d not found", in.ChainSelector)
        }
        current, err := readCurrent(b, chain, in) // live read, not a cached report
        if err != nil {
            return out, err
        }
        if current == in.Desired {
            return out, nil // idempotent: no write
        }
        rep, err := operations.ExecuteOperation(b, SetThing, chain, SetThingInput{ /* ... */ })
        if err != nil {
            return out, err
        }
        out.BatchOps = append(out.BatchOps, rep.Output.BatchOps...)
        return out, nil
    },
)
```

## New registry (family-keyed)

```go
type ThingAdapter interface {
    // ConfigureThing returns the sequence that ... (document encodings and idempotency).
    ConfigureThing() *cldf_ops.Sequence[ConfigureThingInput, sequences.OnChainOutput, cldf_chain.BlockChains]
}

type ThingRegistry struct {
    mu sync.Mutex
    m  map[string]ThingAdapter
}

var (
    thingRegistry     *ThingRegistry
    thingRegistryOnce sync.Once
)

func GetThingRegistry() *ThingRegistry {
    thingRegistryOnce.Do(func() { thingRegistry = &ThingRegistry{m: map[string]ThingAdapter{}} })
    return thingRegistry
}

// Register keeps the first registration for a family.
func (r *ThingRegistry) Register(family string, a ThingAdapter) {
    r.mu.Lock()
    defer r.mu.Unlock()
    if _, ok := r.m[family]; !ok {
        r.m[family] = a
    }
}

func (r *ThingRegistry) Get(family string) (ThingAdapter, bool) {
    r.mu.Lock()
    defer r.mu.Unlock()
    a, ok := r.m[family]
    return a, ok
}
```

For `family-version` keys, use `utils.NewRegistererID(family, version)` as the map key.

## Chain-agnostic changeset

```go
type ConfigureThingConfig struct {
    Chains []uint64   `json:"chains" yaml:"chains"`
    MCMS   mcms.Input `json:"mcms" yaml:"mcms"`
}

func ConfigureThing(reg *ThingRegistry, mcmsReg *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[ConfigureThingConfig] {
    verify := func(e cldf.Environment, cfg ConfigureThingConfig) error {
        for _, sel := range cfg.Chains {
            fam, err := chain_selectors.GetSelectorFamily(sel)
            if err != nil {
                return err
            }
            if _, ok := reg.Get(fam); !ok {
                return fmt.Errorf("no ThingAdapter registered for family %q", fam)
            }
        }
        return nil
    }
    apply := func(e cldf.Environment, cfg ConfigureThingConfig) (cldf.ChangesetOutput, error) {
        ds := datastore.NewMemoryDataStore()
        var batchOps []mcms_types.BatchOperation
        var reports []cldf_ops.Report[any, any]
        for _, sel := range cfg.Chains {
            fam, _ := chain_selectors.GetSelectorFamily(sel)
            a, _ := reg.Get(fam)
            rep, err := cldf_ops.ExecuteSequence(e.OperationsBundle, a.ConfigureThing(), e.BlockChains, ConfigureThingInput{ChainSelector: sel})
            if err != nil {
                return cldf.ChangesetOutput{Reports: reports}, fmt.Errorf("chain %d: %w", sel, err)
            }
            for _, ref := range rep.Output.Addresses {
                if err := ds.Addresses().Add(ref); err != nil {
                    return cldf.ChangesetOutput{}, err
                }
            }
            batchOps = append(batchOps, rep.Output.BatchOps...)
            reports = append(reports, rep.ExecutionReports...)
        }
        return changesets.NewOutputBuilder(e, mcmsReg).
            WithReports(reports).
            WithBatchOps(batchOps).
            WithDataStore(ds).
            Build(cfg.MCMS)
    }
    return cldf.CreateChangeSet(apply, verify)
}
```

## Single-sequence changeset (family-specific)

```go
var DeployThing = changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[SeqInput, evm.Chain, Cfg]{
    Sequence: DeployThingSequence,
    ResolveInput: func(e cldf.Environment, cfg Cfg) (SeqInput, error) {
        return SeqInput{ChainSelector: cfg.ChainSel, ExistingAddresses: e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(cfg.ChainSel))}, nil
    },
    ResolveDep: evm_sequences.ResolveEVMChainDep[Cfg], // Cfg must implement ChainSelector() uint64
})
// usage: DeployThing(changesets.GetRegistry()).Apply(env, changesets.WithMCMS[Cfg]{MCMS: ..., Cfg: ...})
```
