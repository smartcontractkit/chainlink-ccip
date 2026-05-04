---
name: ccip-tooling
description: Use when writing or reviewing Chainlink CCIP deployment tooling, changesets, adapters, registries, datastore flows, integration tests, or operator-facing automation. Guides agents to follow local patterns, infer system state where possible, preserve compatibility, and avoid over-broad abstractions.
---

# CCIP Tooling Code

Use this skill when implementing or reviewing deployment/tooling code in this repo, especially changesets, adapters, registries, datastore flows, and integration tests.

Use skills for repeatable tooling tasks that require reasoning and guidance. Use scripts, make targets, or config for deterministic checks, formatters, and validators.

Before coding, read the relevant local code and docs:

- `deployment/docs/style-guide.md` for changeset authoring rules.
- `deployment/docs/style-guide-contributions.md` only when updating style guidance itself.
- Nearby adapters, registries, changesets, helpers, and tests for the same chain family/version/domain.

## Working Principles

### Start from the Existing Shape

Find the local extension point before inventing one. Prefer extending an existing adapter, registry, sequence, helper, or changeset pattern over adding a parallel subsystem.

Good tooling code usually fits into the repo's current seams:

- chain-family adapters own chain-specific reads/writes,
- shared changesets own user input validation, grouping, idempotency, and MCMS output assembly,
- registries connect families/versions to implementations,
- helpers encode address/ref/datastore conventions.

Add a new abstraction only when more than one concrete caller needs it or when it clearly removes duplicated complexity.

For a fixed set of implementations, prefer explicit maps and lookup functions over global registries, `init()` registration, blank imports, or side-effect packages. Use a registry only when runtime extension is a real requirement, and make the registration path visible at normal call sites.

Name abstractions after the domain concept they model. Avoid generic names like "strategy" when a narrower local term such as adapter, token implementation, resolver, or helper makes ownership clearer.

### Make Operators Provide Intent, Not Fragile State

If the system can derive a value from on-chain state, datastore refs, registry state, or standard conventions, infer it inside the tooling. Do not make users provide values they are unlikely to know or easy to get wrong.

Use explicit inputs for operator intent. Avoid inputs that mirror internal lookup details unless there is a real override use case.

When replacing a fragile input with inference, keep existing public fields when needed for compatibility. Document that they are retained for compatibility or future override behavior, and do not use them as the source of truth.

### Resolve Before Use

Treat input refs and partial refs as lookup keys. Resolve them to full refs before reading address, version, qualifier, or type fields.

Use existing datastore/address helpers instead of hand-parsing or reimplementing lookup behavior. If ambiguity matters, use filters and inspect match count instead of collapsing all failure modes through a helper that expects exactly one match.

### Preserve the Real Unit of Correctness

Group work by the resource that actually receives the write, not by a convenient label.

Examples:

- group contract writes by concrete contract/ref identity when multiple same-version contracts can exist,
- group chain operations by selector when the target chain changes,
- group proposals by MCMS/timelock ownership when execution context changes.

If version, family, or type is not enough to uniquely identify the target, do not use it as the grouping key.

Preserve the owner of batching and proposal assembly. If an outer sequence or changeset already owns MCMS output construction, lower-level helpers should return write outputs rather than constructing their own batches. Do not hide writes inside nested orchestration if the caller must include them in the final proposal.

Separate reads from writes when refactoring deployment flows. Read operations can validate inputs or resolve addresses, but only write operations should be returned as `WriteOutput`s or added to MCMS batches.

### Keep Apply Paths Retry-Safe

Apply paths should be safe to retry after partial success. Read current state, skip no-op updates, and preserve idempotent behavior where the domain supports it.

When using operations that cache reads, be careful after writes. If a flow reads, mutates, and re-reads the same state, clear the operation bundle or use a fresh direct read. In tests, merge the datastore returned by the changeset that actually produced it; do not re-merge stale output from an earlier step.

### Validate the Next Assumption

Validate only what the next step depends on. Good validation prevents bad writes and confusing errors; excessive validation blocks valid migrations and makes code harder to adapt.

Prefer checks like:

- duplicate source/destination entries in user input,
- self-lane mistakes,
- empty or duplicate addresses,
- missing adapter/helper for the family being used,
- resolved ref type/version/address when a specific adapter needs it.

Avoid requiring user inputs that are no longer used by the apply path.

When behavior is gated by capability flags, use the exact flag for the action being performed or tested. Do not substitute a nearby broader capability just because it happens to pass the current cases.

Before restoring support that appears to have been dropped, check whether the exclusion is intentional. Some legacy contract types, versions, or paths may be deliberately unsupported because they are broken, obsolete, or outside the product scope.

### Test Behavior, Not Just Mechanics

Write tests that prove the operator-facing workflow works. Avoid tests that only exercise a newly invented registry or abstraction unless that abstraction is itself the durable API.

Good tooling tests usually verify:

- the changeset derives or resolves the right state,
- the right active target is updated,
- reset/no-op paths behave correctly,
- mixed-family or mixed-version paths use the correct adapter,
- datastore and operation-cache state is fresh after each step.

When refactoring orchestration, add or update tests that prove write outputs still reach the final sequence or changeset output. A test that only verifies returned addresses can miss a dropped deployment or configuration transaction.

## AI Pitfalls To Avoid

- Do not add a registry/interface layer just because code varies by version or chain family; first look for the repo's existing version/family adapter seam.
- Do not use `init()` registration, blank imports, or side-effect packages to connect static implementations when an explicit map or constructor is enough.
- Do not make users pass contract versions, addresses, or refs that the system can discover from canonical state.
- Do not group by version when the write target is a concrete address/ref/resource.
- Do not force orchestration over several reads/writes into an operation just to make it reusable; use a helper when the caller owns sequencing and batching.
- Do not add design docs, `.gitignore` changes, product guards, or unrelated compatibility behavior unless the task or local code requires them.
- Do not remove public fields or change input shapes casually; preserve compatibility unless the user explicitly asks for a breaking change.
- Do not write tests that pass because they mirror your abstraction; test the workflow an operator runs.
- Do not duplicate helper logic for datastore refs, address conversion, MCMS lookup, or defaults without first searching for an existing helper.
- Do not treat "no error" as proof that a returned ref or datastore output is populated and current.
- Do not let comments, log messages, or errors drift from the operation actually being performed; operator-facing wording is part of tooling usability.

## Before Implementing

1. Read nearby implementations and tests for the same domain.
2. Identify the source of truth for every value: user intent, datastore, on-chain state, registry, or convention.
3. Decide where behavior belongs: shared changeset, chain adapter, sequence, helper, or test.
4. Check the deployment style guide for applicable changeset rules.
5. Trace every write operation to the final batch/proposal owner.
6. Keep the smallest abstraction that makes the behavior clear.

## Before Finishing

Confirm:

- public input structs have matching camelCase JSON/YAML tags,
- apply paths are idempotent or intentionally explain why not,
- partial refs are resolved before use,
- inferred/defaulted values reduce operator burden without hiding necessary overrides,
- grouping keys match the actual write target,
- read outputs are not accidentally treated as proposal writes,
- write outputs from helper/orchestration flows reach the final batch/proposal,
- stale reads and stale datastore merges are avoided,
- compatibility behavior is documented,
- capability checks use the exact capability for the behavior,
- tests cover the operator-facing path and meaningful edge cases.
