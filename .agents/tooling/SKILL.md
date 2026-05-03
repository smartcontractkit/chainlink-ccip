---
name: tooling
description: Use when writing or reviewing Chainlink CCIP deployment tooling, changesets, adapters, registries, integration tests, or operator-facing automation. Guides agents to follow local patterns, infer system state where possible, preserve compatibility, and avoid over-broad abstractions.
---

# CCIP Tooling Code

Use this skill when implementing or reviewing deployment/tooling code in this repo, especially changesets, adapters, registries, datastore flows, and integration tests.

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

### Test Behavior, Not Just Mechanics

Write tests that prove the operator-facing workflow works. Avoid tests that only exercise a newly invented registry or abstraction unless that abstraction is itself the durable API.

Good tooling tests usually verify:

- the changeset derives or resolves the right state,
- the right active target is updated,
- reset/no-op paths behave correctly,
- mixed-family or mixed-version paths use the correct adapter,
- datastore and operation-cache state is fresh after each step.

## AI Pitfalls To Avoid

- Do not add a registry/interface layer just because code varies by version or chain family; first look for the repo's existing version/family adapter seam.
- Do not make users pass contract versions, addresses, or refs that the system can discover from canonical state.
- Do not group by version when the write target is a concrete address/ref/resource.
- Do not remove public fields or change input shapes casually; preserve compatibility unless the user explicitly asks for a breaking change.
- Do not write tests that pass because they mirror your abstraction; test the workflow an operator runs.
- Do not duplicate helper logic for datastore refs, address conversion, MCMS lookup, or defaults without first searching for an existing helper.
- Do not treat "no error" as proof that a returned ref or datastore output is populated and current.

## Before Implementing

1. Read nearby implementations and tests for the same domain.
2. Identify the source of truth for every value: user intent, datastore, on-chain state, registry, or convention.
3. Decide where behavior belongs: shared changeset, chain adapter, sequence, helper, or test.
4. Check the deployment style guide for applicable changeset rules.
5. Keep the smallest abstraction that makes the behavior clear.

## Before Finishing

Confirm:

- public input structs have matching camelCase JSON/YAML tags,
- apply paths are idempotent or intentionally explain why not,
- partial refs are resolved before use,
- inferred/defaulted values reduce operator burden without hiding necessary overrides,
- grouping keys match the actual write target,
- stale reads and stale datastore merges are avoided,
- compatibility behavior is documented,
- tests cover the operator-facing path and meaningful edge cases.

