# EVM changeset style guide

**Owning team (tracking reference):** Simon B., Robert, Anindita Ghosh.

This file is a **harness**: consistent structure and **direction** for what belongs here. **Teams own all substantive rules, examples, and updates**—replace placeholders and grow this document with durable, reviewable guidance.

## Purpose

Capture **EVM-specific** practices for writing and reviewing CCIP tooling changesets: selectors, contract surfaces, adapter usage, and footguns that **do not** generalize to other chain families.

## How to use this guide (once populated)

- **While authoring:** use the rules when touching EVM adapters, bytecode/deploy steps, or chain-specific configuration in changesets.
- **Before review or merge:** run the checklist (when added) for EVM-facing changeset PRs.

## What this guide should optimize for

- Correct chain and contract identity (selectors, types, versions) in changeset inputs
- Safe sequencing of EVM deploy/upgrade/configure steps
- Reviewable EVM-specific error modes (reverts, gas, nonces) at the changeset boundary where relevant

## Related documentation

- [Platform (chain-agnostic) style guide](./platform.md) — shared changeset shape and conventions.
- [Ops / chainlink-deployments style guide](./ops.md) — how runs and pipelines are operated.
- [Changeset Style Guide](../../deployment/docs/style-guide.md) — existing deployment changeset reference.
- [Maintaining the Changeset Style Guide](../../deployment/docs/style-guide-contributions.md) — how to add rules, TOC, checklist, and links when editing any changeset-related guide.

## Scope: what belongs here (and what does not)

| In scope (this guide) | Out of scope (use instead) |
| ----------------------- | ---------------------------- |
| EVM **chain selectors**, contract **types/ABIs**, and adapter-specific patterns in changesets | Generic idempotency, shared datastore rules, cross-family input shape → [platform.md](./platform.md) |
| EVM sequencing (deploy/upgrade/configure) **as it differs** from other families | Running **chainlink-deployments** pipelines, environments, operational review → [ops.md](./ops.md) |
| Footguns that are meaningless outside EVM (e.g. chain-ID or EVM-specific ref semantics at this layer) | Adapter implementation internals → deployment architecture / adapter docs |

If a rule applies equally to non-EVM families, it belongs in **platform**, not here.

## Table of contents

<!-- Teams: add a generated TOC or manual list as sections grow. -->

## Illustrative rule shape (replace with real team rules)

The section below shows the **format** only. Delete this section when you add real content, or replace it rule-by-rule.

### ExampleSectionTitleToReplace

**Rule:** (One clear sentence: default behavior.)

**Why it matters:** (Operational / correctness / review impact—do not repeat the rule verbatim.)

```go
// ❌ BAD: (common mistake)

// ✅ BETTER: (preferred pattern)
```

**Exceptions (optional):** (When the rule does not apply.)

---

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

## Direction: topics for the owning team to cover

Turn these **prompts** into short rules with BAD/BETTER examples where they help. Avoid duplicating topics listed in [platform.md](./platform.md) or [ops.md](./ops.md).

- **Chain selectors** and EVM identity in changeset inputs (what must be explicit vs inferred).
- **Contract types, versions, and ABI** handling in changesets—naming and evolution patterns teams should follow.
- **EVM adapter** usage: patterns that are review-friendly and hard to misconfigure **in EVM changesets** (not generic platform idempotency—see platform).
- **Deploy/upgrade/configure ordering** when EVM-specific ordering differs from other families.
- **Operational footguns at the EVM boundary** reviewers should watch for (e.g. wrong network, mismatched bytecode expectations)—staying at changeset/review level, not full chain ops (see ops for CLD runs).

## Contributing and updating

Teams update this guide as EVM patterns evolve. Follow the maintenance pattern in [Maintaining the Changeset Style Guide](../../deployment/docs/style-guide-contributions.md): stable headings for anchors, TOC/checklist hygiene, and links to deeper docs instead of duplicating them.
