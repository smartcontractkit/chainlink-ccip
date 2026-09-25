# Glamsterdam Gas Config Update — Operator Runbook

Audience: an engineer (or AI agent) who needs to run the Glamsterdam gas-config changesets and
generate an MCMS proposal, with no prior context on this feature. Companion doc:
`GLAMSTERDAM_GAS_UPDATE_PLAN.md` (design spec, field-by-field mapping table, code appendix).

If you just want to generate both proposals with minimal reading, skip to **§3 (Quick path)**.
Everything else is context/troubleshooting for when something doesn't work on the first try.

## 0. What this is

Two `ChangeSetV2` implementations in `chainlink-ccip` (branch `glamsterdam-changeset` at time of
writing) automate updating source-side gas config on every lane pointed at a chain that's moving
to the Glamsterdam hard fork:

- `UpdateGasConfigForGlamsterdamV2` — `chains/evm/deployment/v2_0_0/changesets/glamsterdam_gas_update.go`
- `UpdateGasConfigForGlamsterdamV16` — `chains/evm/deployment/v1_6_1/changesets/glamsterdam_gas_update.go`

Both changesets never execute directly — they always produce an MCMS timelock proposal, even if
the deployer key happens to own the contract.

**Important — which domain to use is version-dependent, not a free choice:**

| Version | Domain | Environment (this rehearsal) | Why |
|---|---|---|---|
| v1.6 | `chainlink-deployments/domains/ccip` | `testnet` | v1.6 OnRamp/FeeQuoter/OffRamp are mature and broadly deployed under the `ccip` domain's testnet datastore. |
| v2.0 | `chainlink-deployments/domains/ccv` | `prod_testnet` | v2.0 OnRamp + CommitteeVerifier are being rolled out under the **`ccv`** domain, not `ccip`. As of writing, `ccip`'s testnet datastore has zero v2.0 `OnRamp`/`CommitteeVerifier` entries (only `FeeQuoter` has been upgraded to 2.0.0 there) — running the v2.0 changeset against `ccip`/testnet will discover lanes but produce zero writes. Verify this is still true before you start (§9) — a real rollout may have changed it. |

## 1. Prerequisites

- Two sibling checkouts on disk: `chainlink-ccip` (your feature branch) and `chainlink-deployments`,
  as siblings (i.e. `.../chainlink-ccip` and `.../chainlink-deployments` share a parent directory).
  This matters because the local `replace` directives in step 2 use relative paths
  (`../../../chainlink-ccip`).
- A third sibling checkout, `chainlink-ccv` (yes, confusingly similar name to the `ccv` domain
  above — it's a separate repo). Only needed for the v2.0 path; see §4b.
- VPN connected. Several RPC endpoints in `.config/networks/<env>.yaml` are internal proxies
  (`rpcs.cldev.sh`) that only resolve on VPN.
- No `secrets-<env>.toml` is required for a dry run — RPC endpoints come straight from
  `domains/<domain>/.config/networks/<env>.yaml`. You only need secrets for something that signs
  and broadcasts for real.
- Go 1.26+ (matches `go.mod`; the `ccv` domain's `go.mod` may auto-bump to 1.26.5+ the first time
  you `go mod tidy` it — that's expected, not an error).

## 2. Registration status (already done on this branch)

Both changesets are already registered in `chainlink-deployments` as of this writing — this is
tracked in `chainlink-deployments`, a separate repo from `chainlink-ccip`, so it doesn't show up
in `chainlink-ccip`'s git history. **Check these are still present before running** — see the
"keeps getting reset" note in §8; something in this environment periodically reverts uncommitted
edits to these two files:

- `domains/ccip/testnet/durable_pipelines.go` — look for `registry.Add("update_gas_config_glamsterdam_v16", ...)`
- `domains/ccv/pkg/pipelines/evm_pipelines.go` — look for `registry.Add("update_gas_config_glamsterdam_v2", ...)`

If either is missing, see §5 (v1.6 setup) or §6 (v2.0 setup) to re-add it — those sections have the
exact code to paste back in.

## 3. Quick path — generate both proposals

If registration (§2), the `chainlink-ccip => ../../../chainlink-ccip` replace directives, and the
`chainlink-ccv` patch (§6b) are all already in place, this is the whole workflow:

```bash
# one-time per fresh shell session (see §7.1 for why)
mkdir -p /tmp/solana-programs-bde34821-69b6-47f8-a833-77b47b7e9b36

# v1.6 — ccip domain, testnet
cd chainlink-deployments/domains/ccip/cmd
go run . pipeline run --environment testnet --input-file glamsterdam_gas_update_v16.yaml --dry-run

# v2.0 — ccv domain, prod_testnet
cd ../../ccv/cmd
go run . pipeline run --environment prod_testnet --input-file glamsterdam_gas_update_v2.yaml --dry-run
```

Both input files already exist (checked into your working tree, not yet committed) at:
- `domains/ccip/testnet/durable_pipelines/inputs/glamsterdam_gas_update_v16.yaml`
- `domains/ccv/prod_testnet/durable_pipelines/inputs/glamsterdam_gas_update_v2.yaml`

Expect each run to take several minutes (§7 has timing/troubleshooting notes) and expect to hit at
least one flaky/dead-chain RPC — see §7.2 for the fix pattern (disable that one chain block in
`.config/networks/<env>.yaml`, don't stop the whole rehearsal for it).

**If either command errors**, work through §5 (v1.6-specific) or §6 (v2.0-specific) end to end —
they cover every setup step from scratch, since your local state may differ from what's described
above. Once you've done that once, only §3 is needed for subsequent runs.

## 4. Verifying a proposal (do this after every run, not optional)

See §9 for the full 3-tier checklist (report → decoded diff → fork-execute). At minimum, do tier
(a) and (b) before trusting a proposal — this catches most real bugs, including the two described
in §10.

## 5. v1.6 setup (ccip domain, testnet)

### 5a. Point `domains/ccip` at your local chainlink-ccip branch

```bash
cd chainlink-deployments/domains/ccip
```
`go.mod` already ships this block commented out around line 44 — uncomment it:
```
github.com/smartcontractkit/chainlink-ccip => ../../../chainlink-ccip
github.com/smartcontractkit/chainlink-ccip/chains/evm => ../../../chainlink-ccip/chains/evm
github.com/smartcontractkit/chainlink-ccip/deployment => ../../../chainlink-ccip/deployment
```
Then:
```bash
go mod tidy
```
**If this fails with** `module ... does not contain package .../gobindings/generated/vX_Y_Z/...`
**— your local chainlink-ccip branch is behind what this domain's pinned commit expects.**
Fetch/merge/rebase your branch against `origin/main` in `chainlink-ccip`, then rerun `go mod tidy`.
This is a real version-skew issue, not a config mistake — don't work around it by hand-pinning
`require` versions.

### 5b. Register the changeset

In `domains/ccip/testnet/durable_pipelines.go`, near the other `v1_6_1` changeset registrations
(e.g. next to `DurablePipeline_migrate_hybrid_lock_release_liquidity`):
```go
ccip161evmchangesets "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/changesets"
...
registry.Add("update_gas_config_glamsterdam_v16",
    cldf_changeset.Configure(ccip161evmchangesets.UpdateGasConfigForGlamsterdamV16(cciputilschangeset.GetRegistry())).WithEnvInput())
```
Confirm it builds: `go build ./testnet/...` from `domains/ccip`.

### 5c. Write the input YAML

`domains/ccip/testnet/durable_pipelines/inputs/glamsterdam_gas_update_v16.yaml`:
```yaml
environment: testnet
domain: ccip
changesets:
  - update_gas_config_glamsterdam_v16:
      payload:
        cfg:
          targetChainSelector: 16015286601757825753 # ethereum-testnet-sepolia
          skipChainSelectors: []
        mcms:
          timelockAction: "schedule"
          validUntil: 1893456000
          qualifier: "CLLCCIP"
          description: "Glamsterdam gas config update (v1.6) dry run — Sepolia target"
```
See §8 for why `cfg` and `mcms` must be siblings under `payload` — a very easy mistake to make.

### 5d. Run it

```bash
mkdir -p /tmp/solana-programs-bde34821-69b6-47f8-a833-77b47b7e9b36  # see §7.1
cd domains/ccip/cmd
go run . pipeline run --environment testnet --input-file glamsterdam_gas_update_v16.yaml --dry-run
```
Takes ~5-10 minutes (row 5 does a per-token `getTokenTransferFeeConfig` read for every lane).
Output lands at `domains/ccip/testnet/proposals/*update_gas_config_glamsterdam_v16*.json`.

## 6. v2.0 setup (ccv domain, prod_testnet)

This path has two extra wrinkles v1.6 doesn't: a missing local-env config file, and a version
mismatch in a third repo (`chainlink-ccv`). Both are one-time fixes.

### 6a. Point `domains/ccv` at your local chainlink-ccip branch

```bash
cd chainlink-deployments/domains/ccv
```
Add to the existing `replace (...)` block in `go.mod` (there's no commented-out template here,
unlike `ccip` — just add the three lines):
```
github.com/smartcontractkit/chainlink-ccip => ../../../chainlink-ccip
github.com/smartcontractkit/chainlink-ccip/chains/evm => ../../../chainlink-ccip/chains/evm
github.com/smartcontractkit/chainlink-ccip/deployment => ../../../chainlink-ccip/deployment
```

### 6b. Patch the `chainlink-ccv` version mismatch

`domains/ccv` also depends on a separate repo, `chainlink-ccv`, whose `integration/evm/adapters`
package imports two `chainlink-ccip` paths that moved upstream
(`v2_0_0/operations/{lombard_verifier,cctp_verifier}` → `v2_1_0/operations/...`).
`chainlink-ccv`'s `main` branch hasn't caught up to that move yet. Symptom:
```
module github.com/smartcontractkit/chainlink-ccip/chains/evm@latest found (...), but does not
contain package .../v2_0_0/operations/cctp_verifier
```

**Check first whether this is still broken** — try `go mod tidy` in `domains/ccv` (after 6a) and
see if it succeeds. If `chainlink-ccv` has caught up upstream by the time you read this, skip
straight to 6c. If not:

1. Create a worktree of `chainlink-ccv` at `origin/main`, as a sibling of `chainlink-ccip` /
   `chainlink-deployments` (don't touch your real `chainlink-ccv` checkout):
   ```bash
   cd chainlink-ccv
   git fetch origin main
   git worktree add ../chainlink-ccv-glamsterdam-patch origin/main
   ```
2. In the worktree, fix the two broken import paths (`v2_0_0` → `v2_1_0`) in:
   - `integration/evm/adapters/ccv_indexer_config.go` (both `cctpverifier` and `lombardverifier` imports)
   - `integration/evm/adapters/ccv_token_verifier_config.go` (`cctpverifier` import)
3. Point `domains/ccv/go.mod`'s `replace` block at the patched worktree:
   ```
   github.com/smartcontractkit/chainlink-ccv/integration/evm => ../../../chainlink-ccv-glamsterdam-patch/integration/evm
   ```
4. You may also need to bump `chainlink-ccv`/`chainlink-ccv/integration/evm`/`chainlink-ccv/deployment`
   to their own latest `@main` first if `go mod tidy` still complains about older transitive
   incompatibilities:
   ```bash
   cd domains/ccv
   go get github.com/smartcontractkit/chainlink-ccv/integration/evm@main
   go get github.com/smartcontractkit/chainlink-ccv@main github.com/smartcontractkit/chainlink-ccv/deployment@main
   ```
5. `go mod tidy` again — should now succeed.

**Remove this whole patch once `chainlink-ccv` catches up upstream** — it's a temporary
workaround for a real, independent bug in another team's repo, not something to keep long-term.

### 6c. Create the missing local-env config

`domains/ccv/.config/local/config.prod_testnet.yaml` doesn't exist by default (unlike
`config.prod_mainnet.yaml`, `config.staging_testnet.yaml`, etc., which do). Without it, **zero
chain loaders register for any chain family** — you'll see `"No chain loader available for chain
family, skipping"` for every single chain and `"valid":0,"successful":0"`, with no other error.
Create it (deployer key here is a shared placeholder already used by `prod_mainnet`/
`staging-migration` configs — it never signs anything for real since writes always route through
MCMS):
```yaml
onchain:
  evm:
    deployer_key: "0eca7976bdf758fc689a5d5c572ee8a3898f9bb62abb65508f49a4d4b21a876b"
```

### 6d. Register the changeset

In `domains/ccv/pkg/pipelines/evm_pipelines.go`, inside `registerEVMPipelines` (the file already
imports `evm_changesets "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/changesets"`
for `ActivateRMN` etc., so no new import needed):
```go
registry.Add("update_gas_config_glamsterdam_v2",
    changeset.Configure(evm_changesets.UpdateGasConfigForGlamsterdamV2(mcmsRegistry)).WithEnvInput())
```
Confirm it builds: `go build ./pkg/... ./prod_testnet/...` from `domains/ccv`.

### 6e. Write the input YAML

`domains/ccv/prod_testnet/durable_pipelines/inputs/glamsterdam_gas_update_v2.yaml`:
```yaml
environment: prod_testnet
domain: ccv
changesets:
  - update_gas_config_glamsterdam_v2:
      payload:
        cfg:
          targetChainSelector: 16015286601757825753 # ethereum-testnet-sepolia
          skipChainSelectors: []
        mcms:
          qualifier: "CLLCCIP"
          timelockAction: "schedule"
          validUntil: 1893456000
          description: "Glamsterdam gas config update (v2.0) dry run — Sepolia target (ccv prod_testnet)"
```

### 6f. Run it

```bash
mkdir -p /tmp/solana-programs-bde34821-69b6-47f8-a833-77b47b7e9b36  # see §7.1
cd domains/ccv/cmd
go run . pipeline run --environment prod_testnet --input-file glamsterdam_gas_update_v2.yaml --dry-run
```
Output lands at `domains/ccv/prod_testnet/proposals/*update_gas_config_glamsterdam_v2*.json`.

## 7. Environment gotchas (check every fresh session)

These get wiped by reboots/`/tmp` cleanup and will resurface even after you've fixed them once:

1. **Solana placeholder directory.** The Solana chain loader validates that
   `onchain.solana.programs_dir_path` (set in `domains/<domain>/.config/local/config.<env>.yaml`,
   currently a hardcoded path like `/tmp/solana-programs-<uuid>`) exists — it only checks
   existence, not contents. If it's missing you'll get
   `failed to initialize Solana chain ...: required file does not exist: /tmp/solana-programs-...`.
   Fix: `mkdir -p <that exact path>`. Unrelated to our EVM-only changesets; it's just a
   precondition for loading the environment at all.
2. **Dead/flaky testnet RPCs block the entire run.** The environment loader requires **every**
   registered chain in `.config/networks/<env>.yaml` to load successfully before any changeset can
   run — one dead chain fails the whole command, even if it's irrelevant to your target chain.
   Retry once or twice first (some failures are transient); if a chain fails 2-3 runs in a row on
   all its RPC endpoints, comment its block out in `.config/networks/<env>.yaml` the same way the
   file already does for other known-dead chains (search for `TEMPORARILY DISABLED`, e.g.
   `ethereum-testnet-holesky-taiko-1`). Leave a comment explaining why and when, and don't leave it
   disabled indefinitely without re-checking — do this locally, don't commit unless you're sure it
   should stay disabled.

## 8. Known workflow friction

- **Schema gotcha (cost real debugging time):** the framework decodes the `payload:` key directly
  into `cs_core.WithMCMS[Cfg]{ MCMS mcms.Input; Cfg Cfg }` with `DisallowUnknownFields`. That
  struct has exactly two top-level fields — **`cfg` and `mcms` must be siblings**, with your actual
  changeset config fields nested one level down under `cfg`. Putting `targetChainSelector` directly
  under `payload` (flattened) produces a cryptic `unknown field "skipChainSelectors"` error (it
  fails on the *second* alphabetically-sorted unknown key, not the first — confusing to debug
  blind). The YAML templates in §5c/§6e already have this right.
- **Operation-result caching can serve stale output after a code change.** The durable-pipeline
  runner caches sequence/operation results by input hash in
  `domains/<domain>/<env>/operations_reports/durable_pipelines/<changeset_name>-reports.json`.
  If you change the changeset's Go code and rerun with the *same* input YAML, you may get the
  proposal from *before* your change, with no error or warning that anything was cached — the log
  will show `"Sequence already executed. Returning previous result"` if you look closely. **If a
  code change doesn't seem to take effect, delete that report file and rerun.**
- **`domains/ccip/testnet/durable_pipelines.go`, `domains/ccv/pkg/pipelines/evm_pipelines.go`, and
  both domains' `go.mod` files have been observed reverting between sessions** (registration lines
  and `replace` blocks disappearing without an explicit edit). Cause not identified — possibly a
  formatter, a hook, or a stale-branch merge picking up a version without these changes. Always
  re-check §2's two registration lines and the `chainlink-ccip => ../../../chainlink-ccip` replace
  line are present before assuming "it's already set up."

## 9. Verifying the proposal is doing what you expect

Three checks, increasing in rigor. Do at least the first two before treating a proposal as
trustworthy.

**a. Read the embedded report.** The proposal JSON's `description` field contains a full
per-chain, per-field trace: which chains were skipped (`SkipChainSelectors`), which had no lane to
the target, which had an unresolvable contract address, which had a lane genuinely disabled at the
contract level (see §10), and for every field actually touched: whether the current on-chain value
matched the doc's expected Prague baseline (→ literal Glamsterdam value applied) or didn't (→
fallback value applied, logged as `MISMATCH`). Sanity check the counts look plausible for the
environment (e.g. "22 chains, no lane" is fine on testnet; "0 chains discovered" when you expected
dozens is not).

**b. Decode the proposal into human-readable transactions:**
```bash
go run . mcms analyze-proposal-v2 \
  -e <testnet|prod_testnet> \
  -p ../<env>/proposals/<the proposal file>.json \
  -o /tmp/analysis.md
```
This produces one collapsible section per chain/batch, with every call's decoded ABI inputs
(not raw calldata). For each chain, confirm:
- The field(s) you expect to change show the new (Glamsterdam or fallback) value.
- Every *other* field on the same struct (e.g. `MaxDataBytes`, `GasMultiplierWeiPerEth` on
  `FeeQuoter.DestChainConfig`) is unchanged from its current on-chain value — this proves the
  "merge current struct, only override the touched field" logic didn't clobber anything.
- Cross-reference against `GLAMSTERDAM_GAS_UPDATE_PLAN.md` §6's field table for the expected
  Prague/Glamsterdam/fallback numbers per field.

**c. Actually execute against a fork (strongest check, do this before a real mainnet run):**
```bash
go run . mcms execute-fork \
  -e <testnet|prod_testnet> \
  -p ../<env>/proposals/<the proposal file>.json \
  -s <chain selector> \
  --test-signer
```
Forks that one chain with Anvil and actually runs set-root + execute-timelock against it — the
closest thing to a full dry run without touching real state. Do this per chain you're unsure
about, then read back the contract state to confirm it matches what the decoded proposal claimed.

## 10. Known non-bugs / behaviors you'll likely rediscover

- **Zero proposal / zero writes, no error**: usually means the discovery loop found lanes but a
  *hard-required* contract address (e.g. `OnRamp`, `TokenAdminRegistry`) wasn't resolvable in the
  datastore for any of them — check `operations_reports/.../<name>-reports.json` for what actually
  ran, and the proposal `description` (if a proposal exists) or changeset logs for
  `AddUnresolvedContract` lines. This is what happened when v2.0 was run against `ccip`/testnet:
  `CommitteeVerifier` was made optional (see below), but `OnRamp` v2.0 genuinely doesn't exist
  there yet — hence running v2.0 against `ccv`/`prod_testnet` instead (§0).
- **A single stale/unreachable chain no longer aborts the whole discovery batch.** Both
  `DiscoverLanesToTarget` sequences (`v1_6_1` and `v2_0_0`, under `sequences/glamsterdam/discovery.go`)
  log a warning and skip that one chain (visible in the report as
  `ERROR - failed to read FeeQuoter dest chain config: ..., skipping this chain`) rather than
  failing the entire run. Seen in practice: `sei-testnet-atlantic`'s FeeQuoter address in the
  testnet datastore has no contract code anymore (stale entry) — it's skipped, the other 16 chains
  still got processed.
- **`CommitteeVerifier` is optional for the v2.0 changeset**, same as `OffRamp` already was — a
  missing address just skips that chain's verifier-gas-for-verification write (row 8), it doesn't
  drop the whole lane's OnRamp/FeeQuoter writes.
- **Disabled lanes are detected and skipped, not blindly written to.** `OnRamp.sol` (its `getFee`
  check) and `BaseVerifier.sol` (which `CommitteeVerifier` extends — its comment literally says
  *"The router can be zero to pause the remote chain"*) both use `router == address(0)` as the
  contract's own canonical "this destination isn't configured / lane is paused" signal. The v2.0
  sequence (`sequences/glamsterdam/update_gas_config.go`) checks this before writing to OnRamp or
  CommitteeVerifier, and skips that contract's write entirely if the router is zero — logged as
  `<contract> has no router configured for the target chain (router == address(0)) - lane is
  disabled/not configured`. Earlier versions of this changeset didn't check this and would compute
  a zero-value fallback write (`0 × ratio = 0`) for disabled lanes — harmless in that it wrote the
  same zero back, but noisy, and the underlying principle (never touch a deliberately-disabled
  lane) is worth preserving as this code evolves. If you see many `MISMATCH ... current value 0`
  lines in a report, check whether those chains actually have a disabled lane (`router ==
  address(0)`) rather than assuming the values are wrong.

## 11. Field-by-field mapping (condensed — see plan doc §6 for full detail/notes)

### v1.6 (all confirmed, no open questions)
| # | Field | Expected Prague | Glamsterdam | Fallback |
|---|---|---|---|---|
| 1 | `FeeQuoter.DestChainConfig.DestGasOverhead` | 300,000 | 500,000 | `applyRatio` (~1.667x) |
| 2 | `FeeQuoter.DestChainConfig.DefaultTokenDestGasOverhead` | 90,000 | 270,000 | `applyRatio` (3x) |
| 3 | OffRamp `GasForCallExactCheck` | 5,000 | 5,000 | n/a — read-only sanity check, immutable |
| 5 | `FeeQuoter.TokenTransferFeeConfig.DestGasOverhead` (keyed by dest+token, USDC lanes) | 180,000 | 540,000 (guesstimate) | `applyRatio` (3x) |

(Row 4, Lombard, is dropped entirely for v1.6 — no v1.6 Lombard contract exists anywhere.)

### v2.0
| # | Field | Expected Prague | Glamsterdam | Fallback |
|---|---|---|---|---|
| 1 | `OnRamp.DestChainConfig.BaseExecutionGasCost` | 200,000 | 400,000 | `applyRatio` (2x) |
| 2 | `FeeQuoter.DestChainConfig.DefaultTokenDestGasOverhead` | 90,000 | 270,000 | `applyRatio` (3x) |
| 3 | `FeeQuoter.DestChainConfig.MaxPerMsgGasLimit` | 15,000,000 | 15,000,000 | no-op |
| 4 | `FeeQuoter.DestChainConfig.DestGasPerPayloadByteBase` | 20 | 64 | `applyRatio` (3.2x) |
| 5 | `FeeQuoter.DestChainConfig.DefaultTxGasLimit` | 200,000 | 400,000 | `applyRatio` (2x) |
| 6–7 | OffRamp immutable fields | 5,000 / 12,000 | same | n/a — read-only sanity check |
| 8 | `CommitteeVerifier.GasForVerification` | 75,000 | 85,000 | `applyRatio` (~1.133x) |
| 9 | Lombard pool `TokenTransferFeeConfig.DestGasOverhead` | 410,000 | 1,200,000 (guesstimate) | `applyRatio` (~2.93x) |
| 10 | USDC pool `TokenTransferFeeConfig.DestGasOverhead` | 250,000 | 750,000 (guesstimate) | `applyRatio` (3x) |
| 11 | LombardVerifier `GasForVerification` | 275,000 | 825,000 (guesstimate) | `applyRatio` (3x) |
| 12 | CCTPVerifier ("USDCVerifier") `GasForVerification` | 200,000 | 600,000 (guesstimate) | `applyRatio` (3x) |

**Before the real mainnet run**: swap every "(guesstimate)" value above for a real post-testnet
measurement — it's a constant in `chains/evm/deployment/utils/glamsterdam/` / the version-specific
`fields.go` files, no code logic changes needed.

## 12. Mainnet rollout notes

- Batch mainnet's ~80 lanes using `SkipChainSelectors` in the input YAML's `cfg` block — put
  everything except the batch you're running for that pass in the skip list.
- `SkipChainSelectors` entries are unconditionally excluded, not even checked for a lane — this is
  the intended mechanism for controlled fan-out.
- Re-verify §0's domain choice before the mainnet run — confirm which domain (`ccip` vs `ccv`) and
  environment (`mainnet` vs `prod_mainnet`) actually carries the mature v1.6/v2.0 contracts by then;
  this may have changed since this rehearsal.

## 13. Notes for an AI agent picking this up

If you're an AI agent (Claude Code or otherwise) working through this runbook rather than a human:

- **Don't guess at chain selectors, qualifiers, or addresses.** Every concrete value in this doc
  (chain selectors, the `CLLCCIP` qualifier, the placeholder deployer key) was confirmed against
  this repo's actual datastore/config files, not invented. If you need a different target chain or
  environment, look it up the same way: chain selectors from
  `chain-selectors` repo's `selectors.yml` or by grepping `.config/networks/<env>.yaml`;
  qualifiers by grepping existing archived pipeline inputs under
  `domains/<domain>/<env>/durable_pipelines/archived/*.yaml` for the same changeset family.
- **A `--dry-run` pipeline command still makes real RPC calls and can run for many minutes** (the
  v1.6 run took ~8-10 minutes end to end). Don't kill it prematurely assuming it's hung — check
  the log for periodic `"Executing operation"` lines showing forward progress first. If you must
  run it as a background/monitored command, budget a timeout of at least 10 minutes.
- **A nonzero exit code from the pipeline command is not always a bug in the changeset.** Read the
  actual error first: `no valid RPC clients created` / `required file does not exist:
  /tmp/solana-programs-...` / `no chain loader available` are all environment issues covered in
  §7, not code issues. Only chase changeset code once you've ruled those out.
- **Exit code 0 does not always mean the proposal you expected was generated.** Check the
  `operations_reports/.../<changeset_name>-reports.json` cache-staleness gotcha in §8 before
  concluding a code change had no effect, and check for a proposal file's actual presence (§6/§7 of
  the earlier revision covered this; a "no error, no proposal" outcome usually means zero batch ops
  were produced — see §10's "Zero proposal" entry).
- **When something fails in a way this doc doesn't cover**, the two most useful things to check
  are (1) the datastore for the domain/env you're targeting
  (`domains/<domain>/<env>/datastore/address_refs.json` — grep for the contract type/version you
  expect) to confirm the prerequisite contracts actually exist there, and (2) whether the same
  changeset family has a similar archived input file under `durable_pipelines/archived/` you can
  diff your input against for a schema/qualifier mismatch.
- **Update this runbook if you discover a new gotcha.** It's meant to accumulate operational
  knowledge across runs, not just describe the original rehearsal.
