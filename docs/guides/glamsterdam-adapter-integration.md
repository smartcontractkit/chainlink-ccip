# Integrating a new chain family with the Glamsterdam gas-update adapter

Audience: an engineer on a non-EVM chain-family team (Solana, Aptos, Ton, ...) who needs their
chain family's lanes covered by the Glamsterdam gas-config update changesets. No prior context on
this feature assumed. Operational/EVM-specific detail (what the changeset actually does, how to
run it, verification scripts) lives in
[`../runbooks/glamsterdam-gas-update.md`](../runbooks/glamsterdam-gas-update.md) — read that first
for the "why", this doc is the "how to plug in your family".

## 0. The shape of the integration

Per Glamsterdam-affected contract version (`v1.6.1`, `v2.0.0`), there is:

- A **generic, chain-agnostic layer** you do not modify: a `GasUpdateAdapter` interface, a
  `GasUpdateAdapterRegistry` keyed by chain-family string, an orchestration sequence
  (`GlamsterdamGasUpdateSequence`) that drives any registered adapter through read → resolve →
  write, and a changeset (`UpdateGasConfigForGlamsterdamV16` / `UpdateGasConfigForGlamsterdamV200`)
  that discovers candidate chains **across every chain family**, groups them by family via
  `chain_selectors.GetSelectorFamily`, and dispatches each group to that family's registered
  adapter — `deployment/v1_6_1/adapters`, `deployment/v1_6_1/changesets`, and the `v2_0_0`
  equivalents.
- A **per-family adapter package** implementing that interface. EVM's lives at
  `chains/evm/deployment/v1_6_1/adapters/glamsterdam_gas_adapter.go` and the `v2_0_0` equivalent —
  read these as your reference implementation; they're the only one that exists today.

Integrating your family means: **implement the interface, register it, prove it works.** You do
**not** need a new changeset — the existing generic changeset already dispatches to any registered
family. If your family has no candidate chains with a lane to the target, it's simply never
dispatched to; if it does and your adapter isn't registered yet, the changeset fails loudly for
that family rather than silently skipping it (this is intentional — see
`deployment/v2_0_0/changesets/glamsterdam_gas_update.go`'s "no gas update adapter registered for
chain family" error path).

## 1. Decide which version(s) you need

- `v1.6.1` — only relevant if your family has an equivalent to EVM's mature OnRamp/FeeQuoter v1.6
  contracts already deployed with lanes that need the Prague→Glamsterdam gas bump.
- `v2.0.0` — relevant if your family has CCIP 2.0 contracts (OnRamp 2.0, FeeQuoter 2.0,
  CommitteeVerifier, and optionally Lombard/CCTP verifiers and token pools).

You can implement one or both; they're independent interfaces with independent registries. Most
non-EVM families onboarding CCIP for the first time will likely only need `v2.0.0`.

## 2. Implement `GasUpdateAdapter`

The interface is defined in `deployment/v1_6_1/adapters/glamsterdam_gas_adapter.go` /
`deployment/v2_0_0/adapters/glamsterdam_gas_adapter.go`. Both versions share this shape:

| Method | What it does |
|---|---|
| `HasLaneToTarget` | Reports whether `srcChainSelector` has an enabled lane to `targetChainSelector`, per whatever your family's dest-chain-config-equivalent contract is. Return an error only for a hard failure (RPC error, stale address) — "no lane" is a normal `false`, not an error. (`v1.6.1` additionally returns a `reason string` to put a family-specific explanation in the report instead of the generic "no lane" line — see the EVM v1.6.1 adapter's FeeQuoter-version-ambiguity handling for why this exists; leave it `""` if you don't need it.) |
| `ReadDestGasFields` | Reads every mapping-table field your family's adapter supports for one lane, keyed by `FieldSpec.Name` (see §3). |
| `WriteDestGasFields` | Applies resolved values back on-chain, returning one `mcms_types.BatchOperation` per contract-write-unit of risk. **Isolate writes that might revert independently of the rest** (e.g. a verifier contract your family's timelock might not yet own) into their own batch — see the EVM `v2.0.0` adapter's `buildLaneBatchOps` for the established pattern: one "core" batch, then one batch per isolated write. A timelock batch executes atomically, so this is what keeps one bad write from blocking an entire chain's update. |
| `ReadImmutableSanityFields` | Read-only fields with no setter, purely to flag an unexpected deployment. Never written. |
| `DiscoverCandidateTokens` | Returns every token candidate (raw address bytes, your family's native encoding) whose per-token gas field should be checked. |
| `ReadTokenGasField` / `WriteTokenGasField` | Per-(chain,token) field read/write. The `bool` in `ReadTokenGasField`'s return means "is this field configured at all" — e.g. the EVM `v2.0.0` adapter uses this to skip a token pool that doesn't support the target chain, by checking `GetSupportedChains` first and returning `(0, false, nil)` rather than letting a revert propagate. |

`v2.0.0` only, additionally:

| Method | What it does |
|---|---|
| `TokenFieldSpec` | `v2.0.0`'s mapping table has more than one token-level row (EVM: Lombard vs USDC, with different Prague/Glamsterdam baselines) — this tells the orchestration sequence which `FieldSpec` applies to a given token candidate. If your family only has one token-level row, return the same spec unconditionally. |
| `UpdateFeeQuoterTokenOverrides` | Per-token FeeQuoter overrides are a second, independent source of a token's destination gas (distinct from the token pool's own fee config read/written by `ReadTokenGasField`/`WriteTokenGasField`) — see the EVM adapter and `../runbooks/glamsterdam-gas-update.md`'s "Where the USDC token gas actually comes from" section for why both must move together. If your family's FeeQuoter-equivalent doesn't have this split (one source of truth per token), return `(nil, nil, nil)`. |

**Reference implementation:** read `chains/evm/deployment/v2_0_0/adapters/glamsterdam_gas_adapter.go`
top to bottom before writing your own — it is heavily commented with the *why* behind each piece
(router-zero-means-disabled-lane, FeeQuoter-version classification, batch isolation, etc.), and
most of that reasoning is EVM-contract-specific but the *shape* of the solution (what each method
needs to decide and why) transfers.

## 3. `FieldSpec`s — read this before assuming you can reuse the existing ones

The numeric Prague/Glamsterdam baselines (`deployment/v1_6_1/adapters/fields.go`,
`deployment/v2_0_0/adapters/fields.go`, e.g. `OnRampBaseExecutionGasCost: 200_000 -> 400_000`) are
defined **once**, in the generic layer, and used directly by the EVM adapter. They are EVM gas
numbers, derived from EVM opcode-level cost accounting for the Prague/Glamsterdam hard fork. **Do
not assume these values, or even these field names, are meaningful for your chain's execution-cost
model** (compute units, SOL lamports, Aptos gas units, ... are not EVM gas and do not scale the
same way between hard forks, if your family has an equivalent hard fork at all).

Before writing your adapter's `ReadDestGasFields`/`WriteDestGasFields`, work out with the CCIP core
team:

- Whether your family has an equivalent "hard fork raises the cost of X" event this changeset
  should even be modeling, or whether this entire feature is EVM-specific and your family needs a
  different mechanism.
- If it does apply, what your family's own baseline/target/fallback values are — these are not
  something to invent locally; the EVM values in `fields.go` came from a real measurement/estimate
  per field (see `GLAMSTERDAM_GAS_UPDATE_PLAN.md`, referenced from the runbook).
- Whether your family's `FieldSpec`s belong in a **new, family-scoped file** (e.g.
  `chains/<family>/deployment/v2_0_0/adapters/fields.go`, analogous to how the EVM adapter package
  already keeps its own `glamsterdamTokenPoolKinds` wiring distinct from the generic layer) rather
  than being added to the shared `deployment/v2_0_0/adapters/fields.go` — the existing file's
  values are referenced directly by the EVM adapter, so adding unrelated family-specific constants
  there would conflate two unrelated chains' baselines in one file. `glamsterdamutils.FieldSpec[T]`
  and `glamsterdamutils.Resolve`/`ApplyRatio` (in `deployment/utils/glamsterdam`) are the
  chain-agnostic machinery — only the concrete `FieldSpec` *values* are EVM-specific, and your
  adapter is free to construct its own `FieldSpec` literals pointing at your own baselines.

This is the one piece of this integration that is a judgment call, not a mechanical port — don't
skip the conversation above to save time.

## 4. Register your adapter

Mirror `chains/evm/deployment/v1_6_1/adapters/init.go` / `v2_0_0/adapters/init.go`:

```go
package adapters

import (
	chain_selectors "github.com/smartcontractkit/chain-selectors"

	v2_0_0_adapters "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
)

func init() {
	v2_0_0_adapters.GetGasUpdateAdapterRegistry().RegisterGasUpdateAdapter(
		chain_selectors.FamilyYOURFAMILY, // e.g. chain_selectors.FamilySolana
		&GlamsterdamGasAdapter{},
	)
}
```

`RegisterGasUpdateAdapter` panics if the family is already registered — this should only ever
happen once, in your family's adapter package's `init()`.

**This `init()` only runs if something blank-imports the package.** Within this repo, that means
adding `_ "github.com/smartcontractkit/chainlink-ccip/chains/<family>/deployment/v2_0_0/adapters"`
wherever the EVM equivalent is imported today — currently `devenv/cldf.go` and
`devenv/common/implcommon.go`. In production, the binary that actually runs the durable-pipeline
changeset (in the separate `chainlink-deployments` repo — see the runbook's §2/§5b/§6d) must also
blank-import your package; coordinate with whoever owns that domain's wiring, the same way EVM's
registration got added there.

## 5. Test it

Follow the pattern in `deployment/v1_6_1/adapters/glamsterdam_gas_update_sequence_test.go` /
`deployment/v2_0_0/adapters/glamsterdam_gas_update_sequence_test.go`: a hand-rolled `fakeAdapter`
exercising `GlamsterdamGasUpdateSequence`'s orchestration logic in isolation (no real chain), then
an end-to-end test of your **real** adapter against a simulated chain for your family (mirror
`chains/evm/deployment/v2_0_0/changesets/glamsterdam_gas_update_test.go`, which deploys real
contracts via `environment.WithEVMSimulated` and asserts on the resulting MCMS proposal's batch
count and description). The generic changeset's own tests
(`deployment/v1_6_1/changesets/glamsterdam_gas_update_test.go` /
`deployment/v2_0_0/changesets/glamsterdam_gas_update_test.go`) show how to construct an isolated
`GasUpdateAdapterRegistry` (via `NewGasUpdateAdapterRegistry()`, not the global singleton) for a
test that shouldn't depend on which adapters happen to be blank-imported into the test binary.

## 6. Definition of done

- [ ] `GasUpdateAdapter` implemented for your family, for whichever version(s) apply (§1–§2).
- [ ] Field baselines resolved with the CCIP core team, not invented (§3).
- [ ] Registered via `init()` (§4), blank-imported in this repo's devenv wiring.
- [ ] Unit tests against the orchestration sequence with a fake adapter, plus an end-to-end test
      against a simulated chain with your real adapter (§5).
- [ ] `go build ./...` and `go test ./...` clean in both the `deployment` module and your chain
      family's module (e.g. `chains/evm` is its own module; yours likely is too).
- [ ] Coordinated with whoever owns the `chainlink-deployments` wiring for your target domain/env
      so your adapter package is blank-imported in the binary that actually runs the pipeline —
      without this, your chains are silently never dispatched to (or, if your family is EVM's
      sibling in the same datastore and has a lane but no adapter, the changeset fails loudly
      instead, per §0).
