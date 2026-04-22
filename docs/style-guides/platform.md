# Platform changeset style guide (chain-agnostic)

**Owning team (tracking reference):** Terry Tata, Will Winder.

This guide captures **chain-family-agnostic** practices for writing and reviewing CCIP tooling changesets: input shapes, family routing, datastore conventions, and shared abstractions that should stay consistent no matter which adapter or chain family sits underneath.

It complements—and intentionally does not duplicate—the rules in [Changeset Style Guide](../../deployment/docs/style-guide.md). Read that guide first; the rules here cover platform-level concerns it does not.

When a recommendation here conflicts with established local conventions, prefer consistency unless there is a clear correctness, safety, or operability reason not to.

## How to use this guide

- **While authoring:** apply the rules when shaping shared changeset code (inputs, registry routing, error wrapping, sequence composition) before review.
- **Before review or merge:** run the [Review Checklist](#review-checklist) against changeset PRs that touch shared platform code paths.

## What this guide optimizes for

- Predictable, deterministic apply semantics at the platform layer
- Configuration that operators can reason about without chain-specific trivia
- Clean separation between platform code and family-specific adapters
- Errors that can be triaged from a log line without re-running the changeset

## Related documentation

- [Changeset Style Guide](../../deployment/docs/style-guide.md) — the foundational changeset rules. Read first.
- [EVM changeset style guide](./evm.md) — EVM-only changeset concerns.
- [Ops / chainlink-deployments style guide](./ops.md) — how runs and pipelines are operated.
- [Maintaining the Changeset Style Guide](../../deployment/docs/style-guide-contributions.md) — how to add rules, TOC, checklist, and links when editing any changeset-related guide.

## Scope: what belongs here (and what does not)

| In scope (this guide) | Out of scope (use instead) |
| ----------------------- | ---------------------------- |
| Patterns that apply **across** chain families: input shape, family routing, error wrapping, sequence composition | EVM-specific selectors, contract surfaces, or adapter quirks → [evm.md](./evm.md) |
| Conventions for changeset code that does not know (or should not know) which family it runs against | How to run jobs/pipelines in **chainlink-deployments** → [ops.md](./ops.md) |
| Cross-cutting concerns not already covered by the main [Changeset Style Guide](../../deployment/docs/style-guide.md) | Anything already covered there — link to it instead of restating |

If guidance only makes sense on EVM, it belongs in **evm**, not here. If guidance is about **operating** a deployment (environments, pipeline, review of a run), it belongs in **ops**.

## Table of contents

<!-- toc -->

- [Resolve Chain Family From the Selector, Not From Inputs](#resolve-chain-family-from-the-selector-not-from-inputs)
- [Route Family-Specific Logic Through a Registry, Not a `switch family`](#route-family-specific-logic-through-a-registry-not-a-switch-family)
- [Keep Changeset Inputs Fully Serializable](#keep-changeset-inputs-fully-serializable)
- [Wrap Errors With Identifying Context](#wrap-errors-with-identifying-context)
- [Compose Logic Into Sequences; Keep `apply` a Thin Wirer](#compose-logic-into-sequences-keep-apply-a-thin-wirer)
- [Review Checklist](#review-checklist)

<!-- tocstop -->

## Resolve Chain Family From the Selector, Not From Inputs

**Rule:** Derive chain family from the chain selector at runtime via `chain_selectors.GetSelectorFamily`. Do not accept `family` (or equivalent) as a separate user-provided input.

**Why it matters:** A chain selector already encodes its family. Adding a parallel `family` input creates a second source of truth that can drift from the selector and silently route work to the wrong adapter. Operators should not be able to misconfigure family at all.

```go
// ❌ BAD: family is a separate input that can disagree with the selector
type ConfigureRouterInput struct {
  ChainSelector uint64 `json:"chainSelector,string" yaml:"chainSelector"`
  Family        string `json:"family"             yaml:"family"`
}

func apply(e deployment.Environment, in ConfigureRouterInput) error {
  reader, ok := registry.GetMCMSReader(in.Family) // could be inconsistent with ChainSelector
  if !ok { /* handle error */ }
  // ...
}

// ✅ BETTER: derive family from the selector — single source of truth
type ConfigureRouterInput struct {
  ChainSelector uint64 `json:"chainSelector,string" yaml:"chainSelector"`
}

func apply(e deployment.Environment, in ConfigureRouterInput) error {
  family, err := chain_selectors.GetSelectorFamily(in.ChainSelector)
  if err != nil {
    return fmt.Errorf("unknown chain selector %d: %w", in.ChainSelector, err)
  }
  reader, ok := registry.GetMCMSReader(family)
  if !ok {
    return fmt.Errorf("no MCMS reader registered for family %q (selector %d)", family, in.ChainSelector)
  }
  // ...
}
```

**Exceptions:** Tests or fixtures that intentionally exercise a mismatched family (negative cases) may set family explicitly, but production changeset inputs should not.

---

## Route Family-Specific Logic Through a Registry, Not a `switch family`

**Rule:** When a changeset must do something family-specific (resolve a contract address, read MCMS metadata, build a transaction), look up the behavior through a registry such as `MCMSReaderRegistry` or `ChainFamilyRegistry`. Do not branch on family with a `switch` or `if family == "evm"`.

**Why it matters:** Family branching forces every new chain family to touch every call site that already cares about family. Registry routing keeps platform changesets closed to family changes—adding a new family is a single registration, not a sweep across the repo. It also keeps platform code from importing chain-specific packages, which protects build hygiene.

```go
// ❌ BAD: every new family means editing every call site
func getTimelock(e deployment.Environment, sel uint64, in mcms.Input) (string, error) {
  family, _ := chain_selectors.GetSelectorFamily(sel)
  switch family {
  case chain_selectors.FamilyEVM:
    return evm.GetTimelock(e, sel, in)
  case chain_selectors.FamilySolana:
    return solana.GetTimelock(e, sel, in)
  // ...add another case for every new family, in every changeset that does this
  default:
    return "", fmt.Errorf("unsupported family %s", family)
  }
}

// ✅ BETTER: register once, look up everywhere
func getTimelock(e deployment.Environment, sel uint64, in mcms.Input, reg *MCMSReaderRegistry) (string, error) {
  family, err := chain_selectors.GetSelectorFamily(sel)
  if err != nil {
    return "", fmt.Errorf("unknown chain selector %d: %w", sel, err)
  }
  reader, ok := reg.GetMCMSReader(family)
  if !ok {
    return "", fmt.Errorf("no MCMS reader registered for family %q (selector %d)", family, sel)
  }
  ref, err := reader.GetTimelockRef(e, sel, in)
  if err != nil {
    return "", fmt.Errorf("get timelock ref for selector %d: %w", sel, err)
  }
  return ref.Address, nil
}
```

**Exceptions:** Code that itself defines the registry or tests a single family in isolation may reference a family by name. Cross-family changeset code should not.

---

## Keep Changeset Inputs Fully Serializable

**Rule:** Public changeset input structs (and any nested types) must be fully serializable to JSON and YAML. Do not put functions, channels, `*deployment.Environment`, live chain handles, or other runtime values into them. Pass-through containers like `map[string]any` (e.g. `FamilyExtras`) must hold only serializable values.

**Why it matters:** Operators run changesets from YAML/JSON in `chainlink-deployments` pipelines. Anything non-serializable in an input cannot be expressed in config and turns the changeset into something only callable from Go code. It also makes inputs impossible to round-trip for audit, replay, or dry-run review.

```go
// ❌ BAD: input carries runtime state that cannot be expressed in YAML
type DeployPoolInput struct {
  ChainSelector uint64                      `json:"chainSelector,string" yaml:"chainSelector"`
  Env           *deployment.Environment     // not serializable
  OnDeploy      func(addr string) error     // not serializable
  Extras        map[string]any              // OK only if every value is serializable
}

// ✅ BETTER: inputs describe what to do; the runtime is provided separately
type DeployPoolInput struct {
  ChainSelector uint64                 `json:"chainSelector,string" yaml:"chainSelector"`
  TokenRef      datastore.AddressRef   `json:"tokenRef"             yaml:"tokenRef"`
  Extras        map[string]any         `json:"extras,omitempty"     yaml:"extras,omitempty"` // serializable values only
}
```

A practical check: would `json.Unmarshal(yaml.Marshal(in), &in)` round-trip without losing fields? If not, the input is not pipeline-safe.

**Exceptions:** Internal sequence inputs that are constructed in-process (e.g. by `ResolveInput`) may carry richer types, because they are not user-facing config. The boundary is the changeset's public input struct.

---

## Wrap Errors With Identifying Context

**Rule:** When wrapping or returning errors from changeset code, include the identifying context the operator needs to triage the failure: chain selector, contract type, qualifier, and/or sequence ID. Always wrap the underlying error with `%w` so callers can `errors.Is` / `errors.As`.

**Why it matters:** Changeset failures are often first seen as a single log line in a pipeline run. `"failed to deploy"` forces a re-run with more logging; `"deploy Router on selector 16015286601757825753 (qualifier=router-mainnet): execution reverted"` is usually triageable on the spot. `%w` (instead of `%v`/`%s`) preserves error identity for typed handling upstream.

```go
// ❌ BAD: nothing here helps an operator find the failing chain or contract
if err := deployRouter(ref); err != nil {
  return fmt.Errorf("failed to deploy: %v", err)
}

// ✅ BETTER: identifying context + wrapped error
if err := deployRouter(ref); err != nil {
  return fmt.Errorf(
    "deploy %s on selector %d (qualifier=%q): %w",
    ref.Type, ref.ChainSelector, ref.Qualifier, err,
  )
}
```

When wrapping at a boundary that already adds context (for example, a sequence wrapper that prepends its sequence ID), avoid restating the same fields lower down—prefer adding context once, at the layer that owns it.

**Exceptions:** Sentinel errors returned for `errors.Is` matching (e.g. `ErrAlreadyDeployed`) may be returned bare; let the caller add context.

---

## Compose Logic Into Sequences; Keep `apply` a Thin Wirer

**Rule:** Put the substantive logic (deploys, writes, datastore updates, batch-op assembly) into `operations.Sequence`s and reuse them via helpers like `NewFromOnChainSequence`. The changeset's `apply` function should resolve inputs and dependencies, execute the sequence, and build the output—nothing more.

**Why it matters:** Sequences are the unit of replay, reporting, and reuse in this repo. Logic that lives directly inside `apply` is not visible to the operations bundle, is harder to compose into larger flows, and tends to grow chain-aware branches that should live in a family adapter. Keeping `apply` thin also makes it obvious whether a changeset is safe to retry: idempotency and ordering become properties of the sequence, not the changeset glue.

```go
// ❌ BAD: apply does the work directly; nothing is reusable or replayable
func apply(e deployment.Environment, cfg DeployRouterCfg) (deployment.ChangesetOutput, error) {
  addr, err := deployRouterDirect(e, cfg)
  if err != nil { /* handle error */ }
  if err := saveRefDirect(e, addr); err != nil { /* handle error */ }
  if err := configureDirect(e, addr, cfg); err != nil { /* handle error */ }
  return deployment.ChangesetOutput{ /* hand-built */ }, nil
}

// ✅ BETTER: sequence owns the logic; apply only wires env, deps, and output
var DeployRouter = changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[
  DeployRouterInput, DeployRouterDeps, DeployRouterCfg,
]{
  Sequence: deployRouterSeq, // operations.Sequence — contains the actual logic
  ResolveInput: func(e deployment.Environment, cfg DeployRouterCfg) (DeployRouterInput, error) {
    // build input from cfg + datastore lookups
  },
  ResolveDep: func(e deployment.Environment, cfg DeployRouterCfg) (DeployRouterDeps, error) {
    // resolve chains/clients from env
  },
})
```

**Exceptions:** Trivial changesets that genuinely have no on-chain work (for example, a datastore-only annotation) may skip the sequence and use `apply` directly. As soon as there is on-chain or datastore-mutating logic worth replaying, move it into a sequence.

---

## Review Checklist

Before sending a platform-layer changeset for review, confirm:

- Confirm chain family is derived from the selector, not taken as input.
- Confirm family-specific behavior is resolved through a registry, not a `switch`.
- Confirm public input structs and their nested types are fully YAML/JSON serializable.
- Confirm errors include identifying context (selector, type, qualifier, sequence ID) and wrap underlying errors with `%w`.
- Confirm substantive logic lives in a sequence and `apply` only resolves, executes, and builds output.
- Confirm none of the rules in the main [Changeset Style Guide](../../deployment/docs/style-guide.md) have regressed (idempotency, ref resolution, datastore qualifiers, etc.).

## Copyable rule template

Use this when adding a new section:

````md
## Short, Stable Section Title

**Rule:** State the default behavior in one clear sentence.

**Why it matters:** Explain the correctness, safety, operability, or review impact.

```go
// ❌ BAD: show the common mistake

// ✅ BETTER: show the preferred pattern
```

**Exceptions (optional):** When the rule does not apply.
````

## Contributing and updating

Follow the maintenance pattern in [Maintaining the Changeset Style Guide](../../deployment/docs/style-guide-contributions.md):

- Keep section titles short and stable—they are the anchors used in PR comments.
- Update the [Table of contents](#table-of-contents) and [Review Checklist](#review-checklist) whenever you add, remove, or rename a rule.
- Link to architecture docs and the main [Changeset Style Guide](../../deployment/docs/style-guide.md) instead of duplicating their content.
- Prefer recurring review feedback over one-off bugs as the source of new rules.
