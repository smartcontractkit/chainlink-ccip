# Ops style guide: chainlink-deployments and deployment runs

**Owning team (tracking reference):** Jasmin Bakalović, Phil Ó Condúin.

This file is a **harness**: consistent structure and **direction** for what belongs here. **Teams own all substantive rules, examples, and updates**—replace placeholders and grow this document with durable, reviewable guidance.

## Purpose

Capture **operational** practices for running and reviewing CCIP-related work through **chainlink-deployments (CLD)** and adjacent deployment workflows: environments, pipelines, how changesets are exercised in CI/CD or release paths, and what reviewers should verify for a deployment run.

## How to use this guide (once populated)

- **While preparing a change:** use the rules to decide how the change should be validated through CLD (or documented exceptions).
- **Before promoting or closing a deployment-related PR:** run the checklist (when added) for operational completeness.

## What this guide should optimize for

- Repeatable, reviewable deployment runs with clear ownership
- Safe promotion between environments (where applicable)
- Observable failures and straightforward rollback or retry semantics **at the ops layer**

## Related documentation

- [Platform (chain-agnostic) style guide](./platform.md) — shared changeset conventions.
- [EVM changeset style guide](./evm.md) — EVM-specific changeset authoring.
- [Changeset Style Guide](../../deployment/docs/style-guide.md) — existing deployment changeset reference.
- [Maintaining the Changeset Style Guide](../../deployment/docs/style-guide-contributions.md) — how to add rules, TOC, checklist, and links when editing any changeset-related guide.

## Scope: what belongs here (and what does not)

| In scope (this guide) | Out of scope (use instead) |
| ----------------------- | ---------------------------- |
| **chainlink-deployments** workflows: pipelines, jobs, how runs are triggered and reviewed | Changeset **code structure** (idempotency, refs, shared helpers) → [platform.md](./platform.md) |
| Environment selection, promotion, and operational checklists for deployment **runs** | EVM-specific authoring (selectors, contract typing in changesets) → [evm.md](./evm.md) |
| Expectations for testing/validating changes through CLD vs ad-hoc paths | Low-level adapter implementation → deployment architecture docs |

Authoring guidance for **what goes inside a changeset** belongs in **platform** or **evm**, not here—unless the guidance is specifically about **how that changeset is run or verified** through CLD.

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

Use this when adding a new section (adapt language if the example is not Go):

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

Turn these **prompts** into short rules with BAD/BETTER examples where they help. Avoid duplicating topics listed in [platform.md](./platform.md) or [evm.md](./evm.md).

- How **chainlink-deployments** is expected to be used for CCIP-related changes (vs ad-hoc or legacy paths).
- **Pipeline/job** conventions: naming, required stages, and what “green” must mean before promotion.
- **Environments:** dev/stage/prod (or your equivalents)—rules of thumb for what runs where.
- **Review of deployment outputs:** logs, artifacts, or checkpoints reviewers should expect.
- **Secrets and configuration** at the ops layer (without duplicating in-repo changeset defaults—point to the right separation of concerns).
- **One-off CLD pipelines:** when they are unacceptable versus durable, tested paths (tie to org risk, without naming specific incidents unless public).
- **Observability:** what operators need to see during/after a run to trust the deployment.

## Contributing and updating

Teams update this guide as CLD and release practices evolve. Follow the maintenance pattern in [Maintaining the Changeset Style Guide](../../deployment/docs/style-guide-contributions.md): stable headings for anchors, TOC/checklist hygiene, and links to runbooks or internal docs instead of pasting secrets or environment-specific values.
