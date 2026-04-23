---
title: Changeset Style Guide
sidebar_label: Changeset Style Guide
sidebar_position: 8
---
# Changeset Style Guide

This guide turns recurring review feedback into a practical reference for writing and reviewing changesets.

Use it to improve new work and the code you are already touching. It is not a mandate to rewrite older code that is already correct and operating safely.

When a recommendation conflicts with established local conventions, prefer consistency unless there is a clear correctness, safety, or operability reason not to.

The examples in this guide may age as the codebase evolves. The principles should remain stable.

For implementation details, see:

- [Cross-Family Deployment Architecture](./architecture.md)
- [Implementing Adapters](./implementing-adapters.md)

## How to Use This Guide

Use this guide in two passes:

- **While authoring:** use the rules to shape inputs, defaults, ref handling, and apply behavior before review.
- **Before review or merge:** run the checklist at the end to catch retry hazards, stale reads, weak abstractions, and inconsistent conventions.

## What Good Changesets Optimize For

Good changesets are usually:

- Safe to retry
- Easy to operate
- Hard to misconfigure
- Consistent with repo conventions
- Explicit about failure modes
- Minimal in abstraction and validation

## Table of Contents

<!-- toc -->

- [Make Apply Paths Idempotent](#make-apply-paths-idempotent)
- [Prefer Explicit, Self-Documenting Inputs](#prefer-explicit-self-documenting-inputs)
- [Use Matching `camelCase` YAML and JSON Tags](#use-matching-camelcase-yaml-and-json-tags)
- [Resolve `AddressRef` Inputs Before Use](#resolve-addressref-inputs-before-use)
- [Infer Well-Known Contract Addresses](#infer-well-known-contract-addresses)
- [Provide Sensible Defaults and Fallbacks](#provide-sensible-defaults-and-fallbacks)
- [Handle Empty `GetTimelockRef` Results](#handle-empty-gettimelockref-results)
- [Avoid Stale Reads from Cached Operations](#avoid-stale-reads-from-cached-operations)
- [Use Standard Datastore Qualifiers](#use-standard-datastore-qualifiers)
- [Do Not Use `FindAndFormatRef` for Existence Checks](#do-not-use-findandformatref-for-existence-checks)
- [Reuse Shared Helpers](#reuse-shared-helpers)
- [Prefer the Narrowest Clear Abstraction](#prefer-the-narrowest-clear-abstraction)
- [Avoid Redundant Validation](#avoid-redundant-validation)
- [Review Checklist](#review-checklist)
- [Final Note](#final-note)

<!-- tocstop -->

## Make Apply Paths Idempotent

**Rule:** Write apply logic so it can be retried safely without changing the final outcome.

**Why it matters:** Retries happen. Idempotent apply paths prevent duplicate deployments, duplicate datastore writes, and other partial-failure side effects.

```go
// ❌ BAD: retrying after a partial failure can redeploy or duplicate datastore entries
func deployRouterIfNeeded(qualifier string) error {
  ref := AddressRef{
    ChainSelector: 123121281281818,
    ContractType: "Router",
    Qualifier: qualifier,
    Version: "1.0.0",
  }

  err := deployRouter(ref)
  if err != nil { /* handle error */ }

  err = saveRefToDatastore(ref)
  if err != nil { /* handle error */ }

  // If an error occurs here, retrying may redeploy or duplicate persistence.
  err = configureRouter(ref)
  if err != nil { /* handle error */ }

  return nil
}

// ✅ BETTER: deploy and persist only when the ref does not already exist
func deployRouterIfNeeded(qualifier string) error {
  ref := AddressRef{
    ChainSelector: 123121281281818,
    ContractType: "Router",
    Qualifier: qualifier,
    Version: "1.0.0",
  }

  exists := lookupRefInDatastore(ref)
  if !exists {
    err := deployRouter(ref)
    if err != nil { /* handle error */ }

    err = saveRefToDatastore(ref)
    if err != nil { /* handle error */ }
  }

  err := configureRouter(ref)
  if err != nil { /* handle error */ }

  return nil
}
```

---

## Prefer Explicit, Self-Documenting Inputs

**Rule:** Use field names that make directionality, format, and meaning obvious.

**Why it matters:** Reviewers and operators should not have to decode config structure by inference.

```go
// ❌ BAD: the meaning and format are unclear
type SetLaneConfigInput struct {
  LanePair string `json:"lanePair" yaml:"lanePair"`
}

// ✅ BETTER: directionality and validation requirements are more obvious
type SetLaneConfigInput struct {
  SrcChainSelector uint64 `json:"srcChainSelector,string" yaml:"srcChainSelector"`
  DstChainSelector uint64 `json:"dstChainSelector,string" yaml:"dstChainSelector"`
}
```

Prefer explicit fields over compact encodings when readers would otherwise need to reverse-engineer intent.

---

## Use Matching `camelCase` YAML and JSON Tags

**Rule:** Public input structs should always have matching YAML and JSON tags in `camelCase`.

**Why it matters:** Mismatched tags create downstream parsing failures and make auditing more difficult.

```go
// ❌ BAD: untagged fields are harder to read in YAML and more prone to typos
type DeployPoolInput struct {
  MySuperLongFieldName string
}

// ❌ BAD: mismatched tags can break execution
type DeployPoolInput struct {
  SrcChainSelector uint64 `json:"srcChainSelector,string" yaml:"src_chain_selector"`
  DstChainSelector uint64 `json:"dstChainSelector,string" yaml:"dst_chain_selector"`
}

// ✅ BETTER: matching tags keep YAML and JSON behavior aligned
type DeployPoolInput struct {
  SrcChainSelector uint64 `json:"srcChainSelector,string" yaml:"srcChainSelector"`
  DstChainSelector uint64 `json:"dstChainSelector,string" yaml:"dstChainSelector"`
}
```

When in doubt, use the same `camelCase` name for both tags.

---

## Resolve `AddressRef` Inputs Before Use

**Rule:** Resolve input `AddressRef` values to full refs before reading fields such as `.Address`.

**Why it matters:** Changeset inputs often contain partial refs. Using them directly can fail in ways that are easy to miss.

```go
type ConfigureTokenPoolInput struct {
  TokenRef datastore.AddressRef `json:"tokenRef" yaml:"tokenRef"`
  // ...
}

// ❌ BAD: reads from a possibly unresolved ref
func configureTokenPool(input ConfigureTokenPoolInput) {
  address := input.TokenRef.Address // could be empty

  // ...
}

// ✅ BETTER: resolve before use
func configureTokenPool(input ConfigureTokenPoolInput) {
  ref, err := datastore_utils.FindAndFormatRef(
    input.DataStore,
    input.TokenRef,
    input.Selector,
    datastore_utils.FullRef,
  )
  if err != nil {
    /* handle error */
  }

  address := ref.Address

  // ...
}
```

Treat input refs as lookup keys unless they have already been resolved in the same flow.

---

## Infer Well-Known Contract Addresses

**Rule:** Resolve well-known addresses internally instead of exposing them as user-provided inputs.

**Why it matters:** If the system can derive a value from chain state or the datastore, making the user supply it usually adds risk without adding flexibility.

```go
// ❌ BAD: requires the user to supply an address the system can often infer
type ConfigureRegistryInput struct {
  TokenAdminRegistryAddress string `json:"tokenAdminRegistryAddress" yaml:"tokenAdminRegistryAddress"`
  // ...
}

// ✅ BETTER: resolve the address internally at runtime
func (a *EVMAdapter) ConfigureRegistry(input ConfigureRegistryInput) error {
  tokenAdminRegistryAddress, err := a.GetTokenAdminRegistryAddress(
    input.ExistingDataStore,
    input.Selector,
  )
  if err != nil {
    return fmt.Errorf("failed to get TAR address for chain %d: %w", input.Selector, err)
  }

  // ...
}
```

Only make a value configurable when there is a real operational need for it to vary.

---

## Provide Sensible Defaults and Fallbacks

**Rule:** Prefer resolving values from chain state, the datastore, or standard repo conventions before requiring explicit input.

**Why it matters:** Good defaults reduce YAML surface area and make changesets easier to run correctly.

```go
// ❌ BAD: forces the user to provide an address that often has a safe default
if adminAddress == "" {
  // fail or require user input
}

// ✅ BETTER: use a standard chain-specific fallback
if adminAddress == "" {
  mcmsReader, ok := mcmsRegistry.GetMCMSReader(family)
  if !ok { /* handle not found */ }

  timelockRef, err := mcmsReader.GetTimelockRef(e, selector, cfg.MCMS)
  if err != nil { /* handle error */ }

  if !datastore_utils.IsAddressRefEmpty(timelockRef) {
    adminAddress = timelockRef.Address
  } else {
    adminAddress = chain.DeployerKey
  }
}
```

Default to convention before adding configuration.

---

## Handle Empty `GetTimelockRef` Results

**Rule:** Always check `IsAddressRefEmpty` before using `timelockRef.Address`.

**Why it matters:** `GetTimelockRef` can return no error and still produce an empty ref.

```go
// ❌ BAD: assumes a successful call always returns a populated ref
if adminAddress == "" {
  mcmsReader, ok := mcmsRegistry.GetMCMSReader(family)
  if !ok { /* handle not found */ }

  timelockRef, err := mcmsReader.GetTimelockRef(e, selector, cfg.MCMS)
  if err != nil { /* handle error */ }

  adminAddress = timelockRef.Address // could be empty
}

// ✅ BETTER: branch explicitly on the empty-ref case
if adminAddress == "" {
  mcmsReader, ok := mcmsRegistry.GetMCMSReader(family)
  if !ok { /* handle not found */ }

  timelockRef, err := mcmsReader.GetTimelockRef(e, selector, cfg.MCMS)
  if err != nil { /* handle error */ }

  if !datastore_utils.IsAddressRefEmpty(timelockRef) {
    adminAddress = timelockRef.Address
  } else {
    adminAddress = chain.DeployerKey
  }
}
```

Do not treat “no error” as proof that a ref is usable.

---

## Avoid Stale Reads from Cached Operations

**Rule:** If a changeset reads, mutates, and then re-reads the same state, either clear the operation cache or use direct bindings for the second read.

**Why it matters:** Cached bundle reads can return stale data after a write.

```go
// ❌ BAD: the second read returns a cached result, not the updated value
func verifyConfigUpdate(e *cldf.Environment) error {
  report1, err := cldf.ExecuteOperation(e.OperationsBundle, ops.ReadConfig, /* ... */)
  if err != nil { /* handle error */ }

  report2, err := cldf.ExecuteOperation(e.OperationsBundle, ops.SetConfig, /* ... */)
  if err != nil { /* handle error */ }

  report3, err := cldf.ExecuteOperation(e.OperationsBundle, ops.ReadConfig, /* ... */)
  if err != nil { /* handle error */ }

  // ...
}

// ✅ BETTER: reset the bundle so the second read executes fresh
func verifyConfigUpdate(e *cldf.Environment) error {
  report1, err := cldf.ExecuteOperation(e.OperationsBundle, ops.ReadConfig, /* ... */)
  if err != nil { /* handle error */ }

  report2, err := cldf.ExecuteOperation(e.OperationsBundle, ops.SetConfig, /* ... */)
  if err != nil { /* handle error */ }

  bundle := operations.NewBundle(
    e.GetContext,
    e.Logger,
    operations.NewMemoryReporter(),
  )
  e.OperationsBundle = bundle

  report3, err := cldf.ExecuteOperation(e.OperationsBundle, ops.ReadConfig, /* ... */)
  if err != nil { /* handle error */ }

  // ...
}
```

When you need post-write state, make sure the second read is actually fresh.

---

## Use Standard Datastore Qualifiers

**Rule:** Follow existing datastore qualifier conventions. If no domain-specific convention exists, use `<address>-<type>`.

**Why it matters:** Consistent qualifiers make refs easier to find, reason about, and debug.

```go
// ❌ BAD: missing qualifier can create ambiguity or collisions
err := ds.Addresses().Add(cldf_datastore.AddressRef{
  ChainSelector: selector,
  Type:          cldf_datastore.ContractType("MyContractType"),
  Address:       address.Hex(),
  Version:       version,
})

// ✅ BETTER: qualifier follows a standard, recognizable convention
err := ds.Addresses().Add(cldf_datastore.AddressRef{
  ChainSelector: selector,
  Type:          cldf_datastore.ContractType("MyContractType"),
  Qualifier:     fmt.Sprintf("%s-%s", address.Hex(), "MyContractType"),
  Address:       address.Hex(),
  Version:       version,
})
```

Before inventing a new qualifier format, check whether the domain already has one. Standard conventions:

- **SVM token pools:** `N/A (CLL self-service pool is used)`
- **EVM token pools:** `qualifier = token address`
- **SVM tokens:** `qualifier = token symbol`
- **EVM tokens:** `qualifier = token symbol`

---

## Do Not Use `FindAndFormatRef` for Existence Checks

**Rule:** Use `Filter` and inspect `len(matches)` when you need to distinguish between missing, unique, and ambiguous matches.

**Why it matters:** `FindAndFormatRef` collapses several failure modes into one error path.

```go
// ❌ BAD: treating any error as "missing" conflates several cases
_, err := datastore_utils.FindAndFormatRef(ds, filter, chainSel, datastore_utils.FullRef)
if err != nil {
  return err
}

// ✅ BETTER: handle zero, one, and multiple matches explicitly
matches := input.ExistingDataStore.Addresses().Filter(
  datastore.AddressRefByType(datastore.ContractType(input.ContractType)),
  datastore.AddressRefByChainSelector(input.ChainSelector),
  datastore.AddressRefByQualifier(input.Qualifier),
  datastore.AddressRefByVersion(input.Version),
)

switch len(matches) {
case 0:
  // does not exist
case 1:
  // already exists
default:
  // multiple matches / datastore corruption — treat as error or handle as needed
}
```

It is still fine to use `FindAndFormatRef` when you expect exactly one match and want that assumption enforced.

---

## Reuse Shared Helpers

**Rule:** Prefer existing helpers over custom implementations for common patterns such as ref resolution, timelock lookup, and chain-specific address handling.

**Why it matters:** Shared helpers usually already encode repo conventions and edge cases.

```go
// ❌ BAD: custom logic duplicates existing helper behavior
ref, err := datastore_utils.FindAndFormatRef(e.DataStore, filter, selector, datastore_utils.FullRef)
if err != nil { /* handle error */ }
if ref.Address == "" { /* handle error */ }
if !common.IsHexAddress(ref.Address) { /* handle error */ }
address := common.HexToAddress(ref.Address)

// ✅ BETTER: use the helper that already returns the address in the needed form
address, err := datastore_utils.FindAndFormatRef(
  e.DataStore,
  filter,
  selector,
  evm_datastore_utils.ToEVMAddress,
)
if err != nil { /* handle error */ }
```

Before writing a helper, search for one.

---

## Prefer the Narrowest Clear Abstraction

**Rule:** Use the smallest abstraction that still communicates the intent clearly.

**Why it matters:** Simpler code is easier to read, review, and maintain.

```go
// ❌ BAD: wraps a simple operation in a more verbose helper
func confirmTx(t *testing.T, chain evm.Chain, tx *types.Transaction, callErr error) {
  t.Helper()
  _, err := cldf.ConfirmIfNoError(chain, tx, callErr)
  require.NoError(t, err)
}

// ✅ BETTER: use the narrower API directly
func confirmTx(t *testing.T, chain evm.Chain, tx *types.Transaction) {
  t.Helper()
  _, err := chain.Confirm(tx)
  require.NoError(t, err)
}
```

Reach for a broader abstraction only when it clearly improves clarity or reuse.

---

## Avoid Redundant Validation

**Rule:** Validate only the property the next step actually depends on.

**Why it matters:** Redundant validation adds noise without improving correctness.

```go
// ❌ BAD: the first check is redundant
family, _ := chain_selectors.GetSelectorFamily(selector)
if family != chain_selectors.FamilyEVM { /* ... */ }
chain, ok := e.BlockChains.EVMChains()[selector]
if !ok { /* ... */ }

// ✅ BETTER: this proves the same thing with less code
chain, ok := e.BlockChains.EVMChains()[selector]
if !ok { /* ... */ }
```

Prefer the narrowest check that establishes the required property.

---

## Review Checklist

Before sending a changeset for review, verify each of the following:

- Confirm the apply path can be retried safely.
- Confirm inputs are explicit and self-explanatory.
- Confirm public input structs use matching `camelCase` YAML and JSON tags.
- Confirm every `AddressRef` input is resolved before use.
- Confirm well-known values are inferred instead of exposed as inputs.
- Confirm defaults and fallbacks follow repo conventions.
- Confirm empty `GetTimelockRef` results are handled explicitly.
- Confirm read-after-write flows cannot return stale cached state.
- Confirm datastore qualifiers follow existing conventions.
- Confirm existence checks use `Filter` when ambiguity matters.
- Confirm shared helpers are reused where available.
- Confirm abstractions are no broader than necessary.
- Confirm validation is minimal but sufficient.

## Final Note

This guide exists to make changesets safer to operate, easier to review, and more consistent across the repo.

When a rule here conflicts with clear local conventions, prefer consistency unless doing so would compromise correctness, safety, or operability.

