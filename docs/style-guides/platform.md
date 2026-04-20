# Platform changeset style guide (chain-agnostic)

**Owning team (tracking reference):** Terry Tata, Will Winder.

This file is a **harness**: consistent structure and **direction** for what belongs here. **Teams own all substantive rules, examples, and updates**—replace placeholders and grow this document with durable, reviewable guidance.

## Purpose

Capture **cross-chain** (chain-family-agnostic) practices for writing and reviewing CCIP tooling changesets: shapes, inputs, persistence conventions, and shared abstractions that should stay consistent no matter which adapter or chain family sits underneath.

## How to use this guide (once populated)

- **While authoring:** use the rules to structure inputs, refs, apply behavior, and error handling before review.
- **Before review or merge:** run the checklist (when added) against changeset PRs touching shared platform code paths.

## What this guide should optimize for

- Safe retries and predictable apply semantics **at the platform layer**
- Clear, explicit configuration that reviewers can reason about without chain-specific trivia
- Consistency with repo-wide datastore and helper conventions
- Minimal, well-bounded abstractions shared across chain families

## Related documentation

- [EVM changeset style guide](./evm.md) — EVM-only changeset concerns.
- [Ops / chainlink-deployments style guide](./ops.md) — how runs and pipelines are operated.
- [Changeset Style Guide](../../deployment/docs/style-guide.md) — existing deployment changeset reference (tooling API patterns).
- [Maintaining the Changeset Style Guide](../../deployment/docs/style-guide-contributions.md) — how to add rules, TOC, checklist, and links when editing any changeset-related guide.

## Scope: what belongs here (and what does not)

| In scope (this guide) | Out of scope (use instead) |
| ----------------------- | ---------------------------- |
| Idempotency, ref resolution, datastore usage, and other patterns that apply **across** chain families | EVM-specific selectors, contract typing, or adapter quirks → [evm.md](./evm.md) |
| Naming, tagging, and structural conventions for changeset inputs that are not family-specific | How to run jobs/pipelines in **chainlink-deployments** → [ops.md](./ops.md) |
| Shared helpers, qualifiers, and “narrowest clear abstraction” at the platform layer | Deep adapter implementation detail → deployment docs such as [Cross-Family Deployment Architecture](../../deployment/docs/architecture.md) |

If guidance only makes sense on EVM, it belongs in **evm**, not here. If guidance is about **operating** a deployment (environments, pipeline, review of a run), it belongs in **ops**.

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

Turn these **prompts** into short rules with BAD/BETTER examples where they help. Avoid duplicating topics listed in [evm.md](./evm.md) or [ops.md](./ops.md).

- Apply paths that must be **safe to retry** without double-writing or double-deploying at the platform layer.
- **Explicit, self-documenting inputs** (field names, tags) for YAML/JSON that are not EVM-specific.
- **Refs and lookups:** resolving `AddressRef`-style inputs before use; avoiding “existence check” helpers that encode the wrong semantics.
- **Caching and stale reads:** when cached reads are unsafe across steps in a changeset.
- **Datastore qualifiers and naming** shared across families.
- **Shared helpers:** when to reuse them versus local duplication.
- **Abstractions:** prefer the narrowest abstraction that stays clear for reviewers.
- **Validation:** avoid redundant validation; validate at the right layer.
- **Conflicts with local conventions:** when consistency wins over a “perfect” abstraction.

## Contributing and updating

Teams update this guide as patterns emerge. Follow the maintenance pattern in [Maintaining the Changeset Style Guide](../../deployment/docs/style-guide-contributions.md): keep section titles stable for anchors, update any checklist/TOC you add here, and link to architecture docs instead of copying implementation detail.
