# Glamsterdam Gas Config Update — Operator Runbook

Audience: an engineer (or AI agent) who needs to run the Glamsterdam gas-config changesets and
generate an MCMS proposal, with no prior context on this feature. Companion doc:
`GLAMSTERDAM_GAS_UPDATE_PLAN.md` (design spec, field-by-field mapping table, code appendix).

If you just want to generate both proposals with minimal reading, skip to **§3 (Quick path)**.
Everything else is context/troubleshooting for when something doesn't work on the first try.

**Validation status (read first):**
- **v2.0 (`ccv` / `prod_testnet`)** was fully rehearsed on 2026-10-02: a freshly generated proposal
  (96 batches, 214 txs, 60 chains) passed the Appendix A verification (0 failures) and the Appendix B
  coverage check, every batch was callable by its chain's timelock, and `execute-fork` ran cleanly on
  Base Sepolia, Robinhood, Arbitrum Sepolia, Arc and Hedera testnets. Avalanche Fuji could not be
  fork-executed (no RPC that serves fork state, §9c); its 6 txs were simulated as the timelock instead.
- **v1.6 (`ccip` / `testnet`)**: the changeset now updates only the FeeQuoter each v1.6 OnRamp is
  actually wired to, and skips chains where that is a FeeQuoter 2.0.0 (§10). On `ccip`/`testnet`
  that is **every** chain, so it produced no proposal; the v2.0 proposal covers those lanes. On
  mainnet the v1.6 changeset is expected to have little or nothing to do (73 of 85 v1.6 OnRamps were
  wired to a 2.0.0 FeeQuoter when checked) but **must still be run** to find the exceptions.
- Mainnet has only been dry-run (`ccv`/`prod_mainnet`, file datastore) with an earlier version of the
  code (before verifier-batch isolation and the `destGasOverhead` field): 129 batches, 273 txs, 0 failed
  checks. It must be regenerated, re-verified and fork-executed before a real run.

## START HERE — end-to-end procedure (for an engineer or an AI agent)

Gather before starting: the **target chain selector** (the chain moving to Glamsterdam), the
environment (`ccv`: `prod_testnet` / `prod_mainnet`; `ccip`: `testnet` / `mainnet`), the MCMS
`qualifier` (`CLLCCIP`), a `validUntil` timestamp, and the list of **chains already executed** from any
earlier proposal (they must go in `skipChainSelectors`, see the re-run rules below). Do not guess any of
these; look them up (§13).

**Rules that prevent the expensive mistakes**
1. **Run v2.0 first, then v1.6.** They never write the same contract (the v1.6 changeset skips chains
   whose OnRamp uses a FeeQuoter 2.0.0, which the v2.0 changeset owns), so the order is not
   safety-critical, but v2.0 is the critical one and v1.6 is "whatever is left".
2. **Never regenerate a proposal over chains that already executed.** Re-running is not idempotent for
   generic FeeQuoter token overrides (x3 again) or any field that took a ratio fallback (`MISMATCH`);
   see **§3.1**.
3. **Rehearsal workarounds must never reach a real proposal or a commit**: the `datastore: file`
   override (§1, §6c), commented-out dead chains in `.config/networks/*.yaml` (§7.2), and local
   `replace` lines in `go.mod`.
4. **A proposal that builds is not a proposal that works.** Always run the verification (§9, Appendix A
   and B) and the fork executions (§9c) before handing it over.

**Steps**
1. **Code and versions** (§1, §5a, §6a). The `chainlink-ccip` code you run must contain the current
   changesets; for a real run, bump the module pins rather than using `replace` (§12).
2. **Registration and input** (§2, §5b/c, §6d/e). Create the input YAML with the right target,
   `qualifier`, `validUntil` and `skipChainSelectors`.
3. **Datastore** (§1, §6c). A real run uses the catalog (`datastore: all`, needs catalog access). If you
   have no catalog access, the `datastore: file` workaround produces a rehearsal-only proposal.
4. **Generate**. Move any previous outputs aside first (§8), `mkdir` the Solana placeholder (§7.1),
   then `go run . pipeline run ... --dry-run` (§3). Budget ~4 minutes for `ccv`. A dead RPC on one chain
   fails the whole load: disable only that chain locally and say so (§7.2).
5. **Verify** with Appendix A (per-transaction checks against live chain state) and Appendix B
   (nothing missing or extra). Require `FAILED: 0` and `COVERAGE PROBLEMS: 0`, and read every WARNING.
6. **Check ownership**: every target contract must be owned by the chain's timelock (§9). Verifier
   writes are in their own batches, so one bad owner only fails its own batch.
7. **Fork-execute** the highest-traffic chains with `mcms execute-fork` (§9c). Docker must be running.
8. **Report** what was and was not verified (for example chains with no usable fork RPC), what the
   datastore source was (catalog vs file), and which chains were skipped and why.

## 0. What this is

Two `ChangeSetV2` implementations in `chainlink-ccip` (now on `main`; the pool-skip fix described
in §10 is on branch `fix/glamsterdam-skip-unsupported-token-pools` until it is merged) automate updating source-side gas config on every lane pointed at a chain that's moving
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

- Two sibling checkouts on disk: `chainlink-ccip` and `chainlink-deployments`, as siblings (i.e.
  `.../chainlink-ccip` and `.../chainlink-deployments` share a parent directory). This matters
  because the local `replace` directives in §5a/§6a use relative paths (`../../../chainlink-ccip`).
  **The `chainlink-ccip` checkout must be up to date with `origin/main` (plus the §10 fix if it
  isn't merged yet).** An older feature branch makes `go build` fail in unrelated Solana packages
  (`undefined: tokensapi.TokenAdminRegistryManager`, `...GetRemotePools` signature mismatch) because
  the domain's pinned `chainlink-ccip` version is newer than your checkout. If you don't want to
  move your working branch, use a separate worktree instead:
  `git -C chainlink-ccip worktree add ../chainlink-ccip-glamsterdam-run origin/main` and point the
  `replace` lines at it.
- Only for the v2.0 path, and only if `go mod tidy` fails on `chainlink-ccv` (see §6b): a third
  sibling checkout, `chainlink-ccv` (confusingly similar name to the `ccv` domain — it's a
  separate repo). As of 2026-10-02 this was **not** needed.
- VPN connected. Several RPC endpoints in `.config/networks/<env>.yaml` are internal proxies
  (`rpcs.cldev.sh`) that only resolve on VPN.
- **Catalog (datastore) access, or the local-file workaround.** `domains/ccv` sets
  `prod_testnet` (and `staging_testnet`, `prod_mainnet`) to `datastore: all` in
  `.config/domain.yaml`, which makes the pipeline load the datastore from the remote **catalog**
  service. Without a catalog endpoint + auth the run dies immediately with
  `failed to load datastore: catalog GRPC endpoint is required when datastore location is set to 'catalog'`.
  Catalog auth uses AWS KMS (see `.config/ci/common.env`). Either configure it
  (`catalog.grpc` in `.config/local/config.<env>.yaml` plus AWS creds), or for a **rehearsal only**
  use the file-datastore workaround in §6c. A proposal for a *real* execution must be generated
  against the catalog, not the checked-in file snapshot. **AI agents:** do not go looking for catalog
  credentials or AWS keys on the machine. If the catalog is not configured, stop and ask the human,
  or use the file workaround and label every result "rehearsal, file datastore". The same
  `datastore: all` setting applies to `ccip`'s `testnet` and `mainnet` in `domains/ccip/.config/domain.yaml`.
- Otherwise no `secrets-<env>.toml` is needed for a dry run — RPC endpoints come straight from
  `domains/<domain>/.config/networks/<env>.yaml`.
- `cast` (Foundry) on `PATH`, for the verification script in Appendix A.
- Go 1.26+ (matches `go.mod`; the `ccv` domain's `go.mod` may auto-bump to 1.26.5+ the first time
  you `go mod tidy` it — that's expected, not an error).

## 2. Registration status (local, uncommitted edits in `chainlink-deployments`)

The changesets must be registered in `chainlink-deployments`, a separate repo from
`chainlink-ccip`. As of 2026-10-02 **neither registration, nor the input YAMLs, nor the `go.mod`
`replace` lines were present** on a fresh `chainlink-deployments` checkout (`main` or a feature
branch) — they were only ever local, uncommitted edits. Assume you have to add them (§5b/§5c for
v1.6, §6a/§6d/§6e for v2.0). Check for:

- `domains/ccip/testnet/durable_pipelines.go` — `registry.Add("update_gas_config_glamsterdam_v16", ...)`
- `domains/ccv/pkg/pipelines/evm_pipelines.go` — `registry.Add("update_gas_config_glamsterdam_v2", ...)`

Also see the "keeps getting reset" note in §8: these local edits have been observed disappearing
between sessions.

## 3. Quick path — generate both proposals

If registration (§2), the `chainlink-ccip => ../../../chainlink-ccip` replace directives (§5a/§6a),
the input YAMLs, and — for v2.0 — catalog access or the §6c file-datastore workaround are all in
place, this is the whole workflow. **Before every rerun, move the previous run's outputs away**
(§8: cached operation report + old proposal files), otherwise you may get stale results:

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

The input files are not in the repo; create them from §5c / §6e at:
- `domains/ccip/testnet/durable_pipelines/inputs/glamsterdam_gas_update_v16.yaml`
- `domains/ccv/prod_testnet/durable_pipelines/inputs/glamsterdam_gas_update_v2.yaml`

The v2.0 run makes many live RPC reads across ~70 chains, so budget several minutes (run it in the
background / with a long timeout). It exits non-zero and prints the full CLI usage text on any error — the real message is the `Error:` line (grep for
`^Error` in the log); the usage dump below it is noise.

Expect each run to take several minutes (§7 has timing/troubleshooting notes) and expect to hit at
least one flaky/dead-chain RPC — see §7.2 for the fix pattern (disable that one chain block in
`.config/networks/<env>.yaml`, don't stop the whole rehearsal for it).

**If either command errors**, work through §5 (v1.6-specific) or §6 (v2.0-specific) end to end —
they cover every setup step from scratch, since your local state may differ from what's described
above. Once you've done that once, only §3 is needed for subsequent runs.

### 3.1 Re-running after partial or failed execution (IMPORTANT)

The changesets read the **current** on-chain values and apply their rules. After a proposal has
executed on a chain, running them again against that chain is **not** a no-op:
- fields that matched the Prague baseline are safe (they now equal the Glamsterdam value, so the
  changeset sees "already applied"), but
- fields that took the **ratio fallback** (`MISMATCH` in the report) are scaled *again* (for example
  `destGasPerPayloadByteBase` 16 -> 51 would become 163), and
- every **generic FeeQuoter token override** (anything that is not USDC/Lombard) is multiplied by 3 again.

So: **if some chains executed and you need a new proposal for the rest, put every chain that already
executed into `cfg.skipChainSelectors`** in the input YAML (the list is excluded unconditionally, no lane
check is made), move the previous outputs aside (§8), and regenerate. Procedure:
1. List the chains that executed. For each, confirm on-chain, for example
   `OnRamp.getDestChainConfig(<target>)` shows `baseExecutionGasCost` 400000 for v2.0.
2. Put those chain selectors in `skipChainSelectors`, regenerate, and verify as usual.
3. A chain that is only **partly** executed (some batches done, for example only its verifier batch
   failed) is the dangerous case: regenerating would redo the executed batches too. Prefer to
   **execute the remaining batches of the existing proposal** (batches are independent) instead of
   regenerating; if you must regenerate, skip the chain and fix its remaining batch by hand.

## 4. Verifying a proposal (do this after every run, not optional)

See §9 for the full 3-tier checklist (report → decoded diff → fork-execute). At minimum, do tier
(a) and (b) before trusting a proposal — this catches most real bugs, including the two described
in §10.

For v2.0, tier (b) is automated by **Appendix A** (`verify_proposal.py`): for every transaction it
checks the function selector and target contract type, diffs the write against the **live on-chain
value** (only the gas fields may differ), and recomputes the expected new value from the live current
value with the changeset's rules. **Appendix B** (`verify_coverage.py`) checks the other direction:
every chain with a live lane is in the proposal and nothing extra is. Run both right after the pipeline
finishes and require `FAILED: 0` and `COVERAGE PROBLEMS: 0`.

## 5. v1.6 setup (ccip domain, testnet)

### 5a. Point `domains/ccip` at your local chainlink-ccip branch

> **Known blocker (2026-10-02):** `domains/ccip` currently cannot be built against a current `chainlink-ccip`.
> The local/new `chainlink-ccip` requires `chainlink-evm/gethwrappers` >= 2026-09-15, which removed
> `link_token.ZkBytecode` that `cld-changesets` (even at its latest `main`) still uses; downgrading
> `gethwrappers` instead breaks `chainlink/deployment`. Until `cld-changesets` is fixed upstream, bumping
> `chainlink-ccip` in `domains/ccip` fails with `undefined: link_token.ZkBytecode`. To test v1.6 changes
> meanwhile, apply them onto a worktree of the **pinned** commit (`git worktree add --detach <path> <commit in domains/ccip/go.mod>`)
> and `replace` only `chainlink-ccip/chains/evm` to it.

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
          skipChainSelectors: []   # chains that ALREADY EXECUTED from an earlier proposal go here (§3.1)
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

This path has extra wrinkles v1.6 doesn't: the catalog/datastore requirement (§6c), possible
version skew in a third repo (`chainlink-ccv`, §6b — not needed as of 2026-10-02), and a possibly
missing local-env config file (§6c).

### 6a. Point `domains/ccv` at your local chainlink-ccip checkout

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
Then `go mod tidy && go build ./pkg/... ./prod_testnet/...`. If the build fails inside
`chainlink-ccip/chains/solana/deployment/...` (e.g. `undefined: tokensapi.TokenAdminRegistryManager`),
your `chainlink-ccip` checkout is behind the commit pinned in `go.mod` — update it to
`origin/main` (or use a separate worktree, §1). The checkout you point at must also contain the
§10 pool-skip fix until it is merged to `main`.

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
see if it succeeds. **As of 2026-10-02 it succeeded without any patch, so this whole section was
skipped.** If `chainlink-ccv` has caught up upstream by the time you read this, skip straight to
6c. If not:

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

### 6c. Datastore access and local-env config

**Datastore (new requirement).** `.config/domain.yaml` has `prod_testnet: datastore: all`, so the
run needs the remote catalog (see §1). With no catalog access you get:
```
Error: failed to load datastore: catalog GRPC endpoint is required when datastore location is set to 'catalog'
```
*Rehearsal-only workaround* — read the checked-in snapshot (`prod_testnet/datastore/*.json`)
instead, by changing `prod_testnet` to `datastore: file` in `domains/ccv/.config/domain.yaml`
(the `staging_testnet_*` environments already use `file`). **Do not commit this change**, and do
not use a proposal generated this way for a real execution: the snapshot may be stale relative to
the catalog, so addresses/lanes can differ.

**Local-env config.** On older checkouts `domains/ccv/.config/local/config.prod_testnet.yaml` didn't exist
(unlike `config.prod_mainnet.yaml`, `config.staging_testnet.yaml`, etc., which do); as of
2026-10-02 it exists, so this part can be skipped if the file is there. Without it, **zero
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
          skipChainSelectors: []   # chains that ALREADY EXECUTED from an earlier proposal go here (§3.1)
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
  code change doesn't seem to take effect, move that report file (and the matching
  `<env>/artifacts/durable_pipelines/<changeset_name>/` directory) aside and rerun.** The same
  applies to anything that changes the result without changing the input YAML (e.g. editing the
  `chainlink-ccip` code, or switching the datastore source). Also note each run writes a *new*
  timestamp-prefixed file to `<env>/proposals/` and `<env>/decoded_proposals/` — old ones are not
  overwritten — so move previous outputs aside before a rerun to make sure you verify the latest.
- **`domains/ccip/testnet/durable_pipelines.go`, `domains/ccv/pkg/pipelines/evm_pipelines.go`, and
  both domains' `go.mod` files have been observed reverting between sessions** (registration lines
  and `replace` blocks disappearing without an explicit edit). Cause not identified — possibly a
  formatter, a hook, or a stale-branch merge picking up a version without these changes. Always
  re-check §2's two registration lines and the `chainlink-ccip => ../../../chainlink-ccip` replace
  line are present before assuming "it's already set up."

## 9. Verifying the proposal is doing what you expect

Three checks, increasing in rigor. Do at least the first two before treating a proposal as
trustworthy. For v2.0, **run the Appendix A script as part of (b)** — it automates the
"only the intended fields changed" comparison for every transaction, including the ones
`state_v2.json` has no baseline for.

Things worth knowing about the output files (v2.0, `ccv`):
- The proposal JSON stores each transaction's `data` as **base64**, not hex. Convert before
  feeding it to `cast` (`base64 -d | xxd -p`, or `base64.b64decode(...).hex()` in Python).
- The pipeline also writes a decoded version to `<env>/decoded_proposals/*_decoded.txt`
  (function signature + named struct fields per transaction), so `analyze-proposal-v2` in (b) is
  usually not needed to read it.
- Expected tx functions for v2.0: `OnRamp.applyDestChainConfigUpdates`,
  `FeeQuoter.applyDestChainConfigUpdates`, `FeeQuoter.applyTokenTransferFeeConfigUpdates` (per-token
  overrides), `<Committee|CCTP|Lombard>Verifier.applyRemoteChainConfigUpdates`, and
  `<token pool>.applyTokenTransferFeeConfigUpdates`. Anything else is a bug.
- **Batch layout (v2.0).** Per chain: one *core* batch (OnRamp + FeeQuoter dest config +
  CommitteeVerifier), then **one batch per Lombard/CCTP verifier write**, then one batch per token-pool
  write, then one batch for the FeeQuoter per-token overrides. A timelock batch executes atomically, so
  this isolation means a verifier the timelock doesn't own can only fail its own batch.
- **Check contract ownership before signing.** Every target contract must be owned by the proposal's
  timelock for that chain (a token pool's fee admin also works), otherwise its batch reverts at
  execution time even though the proposal builds and signs fine. Spot-check with
  `cast call <target> 'owner()(address)' --rpc-url <chain rpc>` and compare to the chain's `CLLCCIP`
  timelock (`timelockAddresses` in the proposal JSON). Known case on `prod_testnet` (2026-10-02): 6 of 8
  `CCTPVerifier`s are owned by an EOA (`0x14ea…25f3`, no code, not in the datastore) instead of the
  timelock. Because verifier writes are in their own batches (above), those batches can fail without
  affecting the rest of the chain's update, so the remaining batches can still be signed and executed.
  On `prod_mainnet` all 11 `CCTPVerifier`s are owned by their chain's `CLLCCIP` timelock.
- `<env>/state_v2.json` (checked into the domain; refresh it before relying on it) holds current OnRamp/FeeQuoter
  dest-chain config per chain, so those two can be diffed offline. It does **not** include
  verifier remote-chain config or per-lane pool fee config — read those on-chain
  (`getRemoteChainConfig(uint64)` / `getTokenTransferFeeConfig(address,uint64,bytes4,bytes)`).

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

**c. Actually execute against a fork (strongest check, do this before a real run):**
```bash
go run . mcms execute-fork \
  -e <testnet|prod_testnet|prod_mainnet> \
  -p /absolute/path/to/<the proposal file>.json \
  -s <chain selector> \
  --test-signer
```
Forks that one chain with Anvil (in a Docker container: **Docker must be running**, otherwise it fails
with `failed to set up CTF default network`) and runs set-root, the MCMS executes and the timelock
executes for that chain's batches. Run it **one chain at a time** (each starts a container; a loop of
chains takes ~3-5 minutes each) and pick the highest-traffic chains plus any chain with something
unusual. Reading the log (redact RPC URLs before pasting it anywhere; they contain API keys):
- Success looks like `MCM.setRoot() - success`, `MCM.execute() - success`,
  `Operation N executed successfully` once per batch on that chain, `All operations executed
  successfully`, `Timelock.execute() - success`. Count the operations: it must equal the number of
  batches for that chain in the proposal.
- After the timelock executes, a **post-proposal hook** (`verify-ccip-send`) sends real test CCIP
  messages through the new gas config on the fork. Sends that succeed are extra evidence. Failures
  of the form `dest chain ... not in fork env` are test-harness noise (only one chain is forked) and
  do not mean the proposal failed; judge the proposal by the lines above, not by the hook.
- `execute-fork` **aborts at the first failing operation**; later batches on that chain are not
  attempted. A failing operation `RBACTimelock: underlying transaction reverted` usually means the
  timelock is not the owner of that target (`OnlyCallableByOwner`); check ownership (§9). Real
  execution runs each batch separately, so a failed batch does not block the others.
- **Fork RPC problems are not proposal failures.** `missing trie node` / `failed to get storage` mean
  the RPC the tool picked is not an archive node and dropped the forked block's state (Avalanche Fuji
  hit this on 2026-10-02: every configured RPC either failed its health check or was pruned). The
  tool takes the first two healthy RPCs from `.config/networks/<env>.yaml`; you can temporarily
  comment out the pruned provider locally to force another, and **restore the file afterwards**
  (`git status` on it must be clean). If no RPC works, fall back to simulating each transaction as the
  chain's timelock: `cast rpc eth_call '{"from":"<timelock>","to":"<target>","data":"<hex>"}' latest
  --rpc-url <rpc>` (decode the base64 `data` first). This proves the call itself succeeds, but not the
  MCMS/timelock path, and must be reported as a weaker check.
- **Hedera quirk:** `eth_call` simulation as the timelock fails on `hedera-testnet` with
  `OnlyCallableByOwner()` even though the timelock is the owner (its JSON-RPC relay does not present a
  contract `from` as `msg.sender`); `execute-fork` on Hedera passes. Trust the fork there.

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
- **Token pools that don't support the target chain are skipped (required fix — see below).**
  `CCTPThroughCCVTokenPool.getTokenTransferFeeConfig` reverts with `CCVNotSetOnResolver(address)`
  (selector `0x4172d660`) when its CCTPVerifier resolver has no outbound implementation for the
  destination — i.e. the USDC lane to the target isn't configured on that pool
  (`isSupportedChain(target)` is false, `getSupportedChains()` is empty). Before the fix,
  `UpdateTokenPoolGasConfig` (`sequences/glamsterdam/token_pool_gas_config.go`) read the fee config
  unconditionally, so a single such pool aborted the **entire** changeset with:
  ```
  Error: failed to update token pool gas config for target chain <sel>: failed to read
  TokenPool(0x...) transfer fee config for src <sel>, dst <sel>: execution reverted
  ```
  The fix reads `getSupportedChains()` first and, if the target isn't listed, skips the pool with a
  report line `chain <sel>: TokenPool(0x...) does not support dst <sel>, skipping`. It lives on
  branch `fix/glamsterdam-skip-unsupported-token-pools` (regression test:
  `TestUpdateTokenPoolGasConfig_SkipsPoolNotSupportingTarget`); **make sure whatever
  `chainlink-ccip` code you run against contains it.** On `prod_testnet` (2026-10-02) two pools
  were skipped this way (a Hyperliquid-testnet `CCTPThroughCCVTokenPool` `0x17608A…C035`, and one
  on chain `9763904284804119144`) — the same two chains where the `CCTPVerifier` also has
  `router == address(0)` for the target, which is consistent. The same class of problem — a read that reverts for an
  unconfigured destination aborting everything — is worth checking for in any new field added
  to either sequence.
- **FeeQuoter per-token overrides are migrated for every token, in one read.** The v2.0 OnRamp takes a
  token's destination gas from the pool's own fee config when the pool is `IPoolV2` *and* that config
  is enabled, and from the FeeQuoter per-token override (or `defaultTokenDestGasOverhead`) otherwise,
  so both are updated (`UpdateFeeQuoterTokenTransferFeeConfig`). Overrides are read with one
  `getAllTokenTransferFeeConfigs()` call per FeeQuoter (not by enumerating the `TokenAdminRegistry`
  and reading token by token — that took 20+ minutes on testnet vs ~4 minutes). They are not limited
  to USDC/Lombard: on mainnet there are 440 enabled overrides for Ethereum across 50 chains (SolvBTC,
  LINK, syrupUSDC, WETH, ...), mostly 120k–180k. USDC and Lombard tokens use their own spec (literal
  target on a baseline match, no-op if already migrated, ratio fallback otherwise); **every other
  token with an enabled override is scaled x3 and this is not idempotent** — re-running the changeset
  after its proposal executed would scale them again. USDC/Lombard targets live in one place
  (`fields.go`: `usdc*/lombardDestGasOverhead*` constants) and feed both the pool-level and the
  FeeQuoter-override specs.
- **The v1.6 changeset updates the FeeQuoter the v1.6 OnRamp is actually wired to, and skips 2.0.0 ones.**
  It reads each chain's v1.6 OnRamp `getDynamicConfig().feeQuoter` and only writes it if the datastore
  says it is a 1.6.x FeeQuoter (the v1.6 ABI/selectors cannot talk to a 2.0.0 FeeQuoter). On
  `ccip`/`testnet` (2026-10-02) **every** v1.6 OnRamp is wired to the chain's FeeQuoter **2.0.0** (upgraded
  in place), so the v1.6 changeset produces **no proposal at all** (exit 0, "0 batch operations"): those
  lanes are priced by the 2.0.0 FeeQuoter, which the v2.0 changeset updates, including the legacy
  `destGasOverhead` (row 1b). The skip reasons are logged (`Glamsterdam v1.6 gas config report`) because
  they are otherwise only in the proposal description. Updating the datastore's old FeeQuoter 1.6.0 on
  those chains (what the changeset used to do) changed a contract nothing prices with. Two proposals
  must never write the same FeeQuoter: each rewrites the whole `DestChainConfig` struct from a read taken
  at generation time, so the later one to execute would silently restore the earlier one's fields.
- **A pool that can't be read does not block the batch.** USDC/Lombard pools are only used to
  recognise their token; one with no contract code (a stale datastore entry, e.g. `0xFD47…F910` on
  `prod_testnet` chain `945045181441419236`) logs `WARNING - failed to read underlying token of pool
  ...` and its token falls back to the generic x3 rule.
- **Whole-struct rewrites of unchanged values are normal.** The changeset re-submits the full
  FeeQuoter/OnRamp dest-chain struct with only the touched fields overridden, so you'll see writes
  whose old and new values for a field are identical (e.g. `MaxPerMsgGasLimit` 15,000,000 →
  15,000,000, or 3,000,000 → 3,000,000 where the fallback is a no-op). That's expected; what must
  not happen is any *other* field changing (the Appendix A script checks this).
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
| 1b | `FeeQuoter.DestChainConfig.DestGasOverhead` (legacy; only read by `getValidatedFee`, i.e. by v1.6 OnRamps wired to a 2.0.0 FeeQuoter — no effect on pure-v2.0 lanes) | 300,000 | 500,000 | `applyRatio` (~1.667x) |
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

### Observed v2.0 result — `ccv` / `prod_testnet`, target Sepolia (2026-10-02)

Run against `chainlink-ccip` `main` (`a2d686f61`) plus the §10 fixes (pool-skip, FeeQuoter per-token
overrides, verifier batch isolation, `destGasOverhead`), file datastore (§6c). 96 batches, 214 txs, 60 chains; Appendix A script: 0 failed txs,
and the 88 enabled FeeQuoter overrides for Sepolia on-chain were exactly the 88 in the proposal. Run
time ~4 minutes. Current -> new, with tx counts:

| Contract.field | Current -> new | Count | Note |
|---|---|---|---|
| `OnRamp.BaseExecutionGasCost` | 200,000 -> 400,000 | 59 | literal Glamsterdam |
| `FeeQuoter.DestGasOverhead` | 300,000 -> 500,000 | 57 | literal |
| `FeeQuoter.DestGasOverhead` | 548,000 -> 913,333 | 2 | `MISMATCH` -> ratio fallback |
| `FeeQuoter.DestGasOverhead` | 0 -> 0 | 1 | fallback of 0 |
| `FeeQuoter.DefaultTokenDestGasOverhead` | 90,000 -> 270,000 | 60 | literal |
| `FeeQuoter.DefaultTxGasLimit` | 200,000 -> 400,000 | 60 | literal |
| `FeeQuoter.DestGasPerPayloadByteBase` | 20 -> 64 | 58 | literal |
| `FeeQuoter.DestGasPerPayloadByteBase` | 16 -> 51 | 2 | `MISMATCH` -> fallback (these chains are non-default) |
| `FeeQuoter.MaxPerMsgGasLimit` | 15,000,000 -> 15,000,000 | 58 | no-op |
| `FeeQuoter.MaxPerMsgGasLimit` | 3,000,000 -> 3,000,000 | 2 | `MISMATCH`, fallback is a no-op |
| `CommitteeVerifier.GasForVerification` | 75,000 -> 85,000 | 59 | literal |
| `CCTPVerifier.GasForVerification` | 200,000 -> 600,000 | 6 | literal (guesstimate value) |
| `CCTPVerifier.GasForVerification` | 220,000 -> 660,000 | 2 | `MISMATCH` -> fallback |
| `CCTPThroughCCVTokenPool` (USDC slot) `DestGasOverhead` | 90,000 -> 270,000 | 8 | `MISMATCH` -> fallback |
| FeeQuoter per-token override, USDC | 90k / 180k / 200k -> 270k / 540k / 600k | 7 | all `MISMATCH` -> x3 (none at the 250k baseline on testnet) |
| FeeQuoter per-token override, other tokens | e.g. 120k -> 360k (x41), 140k -> 420k (x7); largest 750k -> 2.25M | 81 | generic x3 |

Not written: 8 chains skipped for an unresolved contract address (5 `OnRamp`, 3 `FeeQuoter`);
OnRamp/CommitteeVerifier/CCTPVerifier writes skipped on a few chains with `router == address(0)`;
2 token pools skipped as unsupported for the target (§10); 1 USDC pool with no contract code (warning,
§10). Ownership (§9): 6 `CCTPVerifier`s are EOA-owned, so those isolated verifier batches would revert at execution; nothing else is affected.

### Where the USDC token gas actually comes from (OnRamp fee path) — row 10 is incomplete

All 8 `CCTPThroughCCVTokenPool`s on `prod_testnet` have a current pool-level `DestGasOverhead` of
90,000 (not the 250,000 "expected Prague" baseline in row 10), so every one takes the `MISMATCH`
fallback (3× → 270,000) instead of the literal 750,000. The reason is that **the 250,000 is not a
pool value at all.** Traced in `chains/evm/contracts/onRamp/OnRamp.sol` (`_getReceipts`):

1. For a token transfer the `OnRamp` calls `pool.getFee(...)` if the pool supports `IPoolV2`.
   `USDCTokenPoolProxy.getFee` forwards to the underlying pool selected for that lane's mechanism
   (and reverts for a legacy mechanism); `CCTPThroughCCVTokenPool` inherits `TokenPool.getFee`,
   which returns the pool's own `s_tokenTransferFeeConfig[dest]` **only if that config is
   enabled**, otherwise `(0, 0, 0, 0, isEnabled=false)`.
2. If the pool is not `IPoolV2`, or returned `isEnabled=false`, the `OnRamp` falls back to
   `FeeQuoter.getTokenTransferFee(dest, token)`, i.e. the FeeQuoter's **per-token override**
   `getTokenTransferFeeConfig(dest, token)` (or `defaultTokenDestGasOverhead`, 90,000, if there is
   none).
3. Exactly one of the two supplies the token's `destGasLimit`; the other is ignored.

Evidence from Base mainnet (FeeQuoter 2.0.0 `0x0352…3Cb62`, event history for dest = Ethereum
`5009297550715157269`): native USDC (`0x8335…2913`) got a per-token override of
`destGasOverhead = 180,000` on 2026-05-19 and was raised to **250,000** on 2026-06-11 (block
47175441); it is still 250,000 on-chain. The FeeQuoter 1.6.3 on Base still has 180,000 for USDC (the
v1.6 row-5 baseline). `DestChainConfig` was never 250k (`defaultTokenDestGasOverhead` stayed 90,000).
Dozens of other tokens have overrides in the 100k–410k range. So the "250,000" in row 10 is the
**FeeQuoter 2.0.0 per-token override for USDC**, not `CCTPThroughCCVTokenPool`'s own config.

**Both sources are now migrated (implemented).** The v2.0 changeset updates the pool-level fee config
(`UpdateTokenPoolGasConfig`: Lombard / Siloed USDC / `CCTPThroughCCVTokenPool` v2.0.0 pools) **and**
every FeeQuoter 2.0.0 per-token override for the target (`UpdateFeeQuoterTokenTransferFeeConfig`,
§10), so the Glamsterdam change is not lost if the active source flips (e.g. a pool's fee config is
later disabled and the OnRamp falls back to the override).

**Which source is actually charged depends on the lane — and `state_v2.json` is not a reliable guide.**
`prod_mainnet/state_v2.json` only lists the legacy `USDCTokenPool` 1.5.1/1.6.2, but the **datastore**
(`datastore/address_refs.json`) also has 11 `CCTPThroughCCVTokenPool` 2.0.0 and 11 `USDCTokenPoolProxy`
2.0.0 on mainnet. Checked on Base mainnet (2026-10-02): the USDC token is registered in the
`TokenAdminRegistry` to the `USDCTokenPoolProxy` `0xfe38…6170`, and its `getFee(USDC -> Ethereum)`
returns `destGasOverhead = 90,000, isEnabled = true`, so (if the v2.0 OnRamp is the one in use for the
lane) the **pool-level 90,000 is what is charged**, and the FeeQuoter override (250,000) is currently
shadowed. That is why the pool-level `MISMATCH` fallback (90,000 -> 270,000) is the value that matters
on those lanes, and why the 250,000 baseline of row 10 never matches them. Whether 90,000 at the pool
was intended (vs carrying over the 250,000 the lane used to be charged via the FeeQuoter override)
is worth confirming with whoever configured the proxy pools before a real run. (Not checked: which
OnRamp/Router a given lane actually routes through.)

## 12. Mainnet rollout notes

- Batch mainnet's ~80 lanes using `SkipChainSelectors` in the input YAML's `cfg` block — put
  everything except the batch you're running for that pass in the skip list. (The same list is how
  you exclude chains that already executed when regenerating, §3.1.)
- `SkipChainSelectors` entries are unconditionally excluded, not even checked for a lane — this is
  the intended mechanism for controlled fan-out.
- Re-verify §0's domain choice before the mainnet run — confirm which domain (`ccip` vs `ccv`) and
  environment (`mainnet` vs `prod_mainnet`) actually carries the mature v1.6/v2.0 contracts by then;
  this may have changed since this rehearsal.
- **A real proposal must not be generated the way the rehearsal was.** Checklist:
  1. Run with **catalog access** (§1), not the `datastore: file` workaround; make sure
     `.config/domain.yaml` is back to `datastore: all`.
  2. Do not use local `replace` directives. Merge the §10 fix (and any guesstimate-constant
     updates) into `chainlink-ccip` `main`, then **bump the `chainlink-ccip` module pins** in
     `chainlink-deployments/domains/<domain>/go.mod` to a commit that contains them, so the
     proposal is reproducible from the pinned versions. (The pin on 2026-10-02 already contained
     the Glamsterdam changesets themselves, just not the §10 fix.)
  3. Land the changeset registration and input YAML in `chainlink-deployments` via a normal PR
     (they were uncommitted local edits during the rehearsal).
  4. Re-run the Appendix A and B verification on the final proposal, then `execute-fork` the
     highest-traffic chains (§9c). The testnet rehearsal did this; mainnet has not been fork-executed yet.
  5. Resolve the USDC-pool baseline question in §11 and replace the "(guesstimate)" constants.

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
- **Sandboxed shells.** Inside Claude Code's default Bash sandbox, `go mod tidy` / `go build` can't
  download modules (`proxy.golang.org` is blocked, then fails TLS verification) and the pipeline /
  `cast` calls to `rpcs.cldev.sh` need network access too. Those commands have to run outside the
  sandbox (with the user's approval); say so rather than looking for another route.
- **Don't touch the user's working branch to fix version skew.** If their `chainlink-ccip`
  checkout is behind the pinned version, use a separate `git worktree` at `origin/main` (§1) and
  point the `replace` lines there; commit any fix on its own branch.
- **Don't print config files to find their structure.** `.config/local/config.*.yaml` can hold
  plaintext credentials (e.g. Job Distributor passwords in the `staging_testnet` config). Use `grep`
  for key names, or read only the file you need.
- **Verify, don't just generate.** "Exit code 0 + a proposal file" is not success: run the
  Appendix A script and require `FAILED: 0`. Also be explicit in your report about what was *not*
  checked (e.g. fork execution) and what the datastore source was (file vs catalog).
- **Update this runbook if you discover a new gotcha.** It's meant to accumulate operational
  knowledge across runs, not just describe the original rehearsal.

## Appendix A. `verify_proposal.py` — automated v2.0 proposal check

This is a helper that an engineer or AI agent can copy out of this file; it is **not** part of the
repo or of CLD. Save it as `verify_proposal.py` (outside the repo) and run it right after the pipeline
finishes. It needs `python3` (standard library only), `cast` (Foundry) and the VPN, and reads current
values from `https://rpcs.cldev.sh/<selector>`:

```bash
python3 verify_proposal.py chainlink-deployments/domains/ccv/<env> <target chain selector>
```

It uses the newest `proposals/*glamsterdam_v2*.json` in that directory, so move older runs aside first
(§8). The exit code is non-zero if any transaction fails. Per transaction it enforces:
1. the 4-byte selector is one of the expected functions (§9) and the target address is the expected
   contract type in `datastore/address_refs.json`;
2. `OnRamp` / `FeeQuoter` dest-chain writes: every field except the gas fields equals the **live**
   on-chain value (`state_v2.json` is a snapshot and can be stale, so it is not used);
3. verifier / token-pool / FeeQuoter-override writes: every field except the gas field equals the live
   on-chain value, no override is removed or disabled;
4. the **new value equals what the changeset's rules compute from the live current value**
   (literal Glamsterdam value on a baseline match, no-op if already migrated, otherwise the ratio
   fallback; FeeQuoter per-token overrides x3);
5. the destination selector inside the calldata is the intended target;

and prints every current -> new value transition with counts, which you compare to the report in the
proposal `description` and to §11. It retries transient RPC-proxy errors, but a chain whose RPC keeps
failing shows up as a failed transaction (`ERR ...`): re-run before treating it as real. It does
**not** simulate execution (use §9c) and does **not** check ownership (use §9). If you change a rule
in `fields.go`, update the `EXPECT` table in the script to match.

```python
#!/usr/bin/env python3
"""Verify a Glamsterdam v2.0 MCMS proposal: right functions, right contracts, only intended fields changed.

usage: python3 verify_proposal.py <domains/ccv/<env> dir> <target_selector>
Needs: `cast` (foundry), VPN (reads live values from https://rpcs.cldev.sh/<selector>).
Checks:
  1. every tx targets a contract of the expected type per <env>/datastore/address_refs.json
  2. every tx's 4-byte selector is one of the 4 expected functions
  3. OnRamp / FeeQuoter dest-chain writes vs LIVE on-chain values (state_v2.json is a snapshot and can be stale): only gas fields differ
  4. CommitteeVerifier / CCTPVerifier / token-pool writes vs LIVE on-chain values: only the gas field differs
"""
import base64, glob, json, re, subprocess, sys, time
from collections import Counter
from concurrent.futures import ThreadPoolExecutor

BASE = sys.argv[1].rstrip('/') + '/'
T = int(sys.argv[2])  # target chain selector (the chain moving to Glamsterdam)

prop = json.load(open(sorted(glob.glob(BASE + 'proposals/*glamsterdam_v2*.json'))[-1]))
ref = {(r['chainSelector'], r['address'].lower()): r for r in json.load(open(BASE + 'datastore/address_refs.json'))}

def run(cmd, tries=4):
    """Run a cast command, retrying transient RPC failures (the RPC proxy occasionally returns 'header not found')."""
    for i in range(tries):
        p = subprocess.run(cmd, capture_output=True, text=True, timeout=90)
        if p.returncode == 0: return p.stdout
        if i == tries - 1 or not re.search(r'header for number not found|ErrUpstreamsExhausted|timed out|timeout|502|503|504|connection', p.stderr, re.I):
            raise RuntimeError(p.stderr.strip()[:200])
        time.sleep(2 * (i + 1))

def cast(*a):
    return re.sub(r' \[[0-9.e+-]+\]', '', run(['cast', *a]).strip())  # cast annotates big numbers, e.g. "270000 [2.7e5]"

def livejson(to, sig, url, *args):
    return json.loads(run(['cast', 'call', to, sig, *args, '--json', '--rpc-url', url]))[0]  # the single tuple output, as a list of fields

FN = {  # name -> (signature, expected datastore type)
    'OnRamp.applyDestChainConfigUpdates': ('applyDestChainConfigUpdates((uint64,address,uint8,bool,uint16,uint16,uint32,address[],address[],address,bytes)[])', {'OnRamp'}),
    'FeeQuoter.applyDestChainConfigUpdates': ('applyDestChainConfigUpdates((uint64,(bool,uint32,uint32,uint32,uint8,bytes4,uint16,uint32,uint32,uint16,uint8))[])', {'FeeQuoter'}),
    'Verifier.applyRemoteChainConfigUpdates': ('applyRemoteChainConfigUpdates((address,uint64,bool,uint16,uint32,uint16)[])', {'CommitteeVerifier', 'CCTPVerifier', 'LombardVerifier'}),
    'FeeQuoter.applyTokenTransferFeeConfigUpdates': ('applyTokenTransferFeeConfigUpdates((uint64,(address,(uint32,uint32,uint32,bool))[])[],(uint64,address)[])', {'FeeQuoter'}),
    'Pool.applyTokenTransferFeeConfigUpdates': ('applyTokenTransferFeeConfigUpdates((uint64,(uint32,uint32,uint32,uint32,uint16,uint16,bool))[],uint64[])', {'CCTPThroughCCVTokenPool', 'LombardTokenPool', 'SiloedUSDCTokenPool', 'USDCTokenPool'}),
}
SEL = {cast('sig', s): n for n, (s, _) in FN.items()}

ops = []
for bi, b in enumerate(prop['operations']):
    for t in b['transactions']:
        d = t['data'] if t['data'].startswith('0x') else '0x' + base64.b64decode(t['data']).hex()  # proposal JSON stores base64
        ops.append((bi, int(b['chainSelector']), t['to'].lower(), d))
print(f'{len(prop["operations"])} batches, {len(ops)} txs, {len({o[1] for o in ops})} chains')

def check(op):
    bi, sel, to, data = op
    r = ref.get((sel, to)); typ = r['type'] if r else None
    fn = SEL.get(data[:10])
    out = dict(batch=bi, sel=sel, to=to, typ=typ, fn=fn, errs=[])
    if fn is None: out['errs'].append('unexpected selector ' + data[:10]); return out
    if typ not in FN[fn][1]: out['errs'].append(f'{fn} on unexpected contract type {typ}')
    url = f'https://rpcs.cldev.sh/{sel}'
    try:
        if fn == 'Verifier.applyRemoteChainConfigUpdates':
            m = re.match(r'\[\((0x\w{40}), (\d+), (true|false), (\d+), (\d+), (\d+)\)\]', cast('calldata-decode', 'x(' + '(address,uint64,bool,uint16,uint32,uint16)[])', data))
            new = m.groups()
            cur = re.match(r'\((0x\w{40}), (\d+), (true|false), (\d+), (\d+), (\d+)\)', cast('call', to, 'getRemoteChainConfig(uint64)((address,uint64,bool,uint16,uint32,uint16),address[])', str(T), '--rpc-url', url)).groups()
            if int(new[1]) != T: out['errs'].append('wrong dest chain')
            for i, n in [(0, 'router'), (2, 'allowlistEnabled'), (3, 'feeUSDCents'), (5, 'payloadSizeBytes')]:
                if cur[i].lower() != new[i].lower(): out['errs'].append(f'{n} changed {cur[i]}->{new[i]}')
            out['t'] = (typ, 'GasForVerification', int(cur[4]), int(new[4]))
        elif fn == 'Pool.applyTokenTransferFeeConfigUpdates':
            new = cast('calldata-decode', 'x((uint64,(uint32,uint32,uint32,uint32,uint16,uint16,bool))[],uint64[])', data).replace('\n', ', ')
            m = re.match(r'\[\((\d+), \((\d+), (\d+), (\d+), (\d+), (\d+), (\d+), (true|false)\)\)\], \[\]', new)
            if int(m.group(1)) != T: out['errs'].append('wrong dest chain')
            nv = list(m.groups()[1:])
            cv = list(re.match(r'\((\d+), (\d+), (\d+), (\d+), (\d+), (\d+), (true|false)\)', cast(
                'call', to, 'getTokenTransferFeeConfig(address,uint64,bytes4,bytes)((uint32,uint32,uint32,uint32,uint16,uint16,bool))',
                '0x' + '0' * 40, str(T), '0x00000000', '0x', '--rpc-url', url)).groups())
            for i, n in enumerate(['destGasOverhead', 'destBytesOverhead', 'finalityFeeUSDCents', 'fastFinalityFeeUSDCents', 'finalityBps', 'fastFinalityBps', 'isEnabled']):
                if i and cv[i] != nv[i]: out['errs'].append(f'{n} changed {cv[i]}->{nv[i]}')
            if cv[6] != 'true': out['errs'].append('pool fee config currently disabled')
            out['t'] = (typ, 'DestGasOverhead', int(cv[0]), int(nv[0]))
        elif fn == 'FeeQuoter.applyTokenTransferFeeConfigUpdates':
            new = cast('calldata-decode', 'x((uint64,(address,(uint32,uint32,uint32,bool))[])[],(uint64,address)[])', data).replace('\n', ', ')
            if int(re.match(r'\[\((\d+),', new).group(1)) != T: out['errs'].append('wrong dest chain')
            if not new.endswith(', []'): out['errs'].append('unexpected tokensToUseDefaultFeeConfigs (override removal)')
            entries = re.findall(r'\((0x\w{40}), \((\d+), (\d+), (\d+), (true|false)\)\)', new)
            if not entries: out['errs'].append('no token entries decoded')
            out['ts'] = []
            for tok, fee, gas, byt, en in entries:
                c = re.match(r'\((\d+), (\d+), (\d+), (true|false)\)', cast(
                    'call', to, 'getTokenTransferFeeConfig(uint64,address)((uint32,uint32,uint32,bool))', str(T), tok, '--rpc-url', url)).groups()
                if c[3] != 'true': out['errs'].append(f'{tok}: override currently disabled (would be re-enabled)')
                for n, a_, b_ in (('feeUSDCents', c[0], fee), ('destBytesOverhead', c[2], byt), ('isEnabled', c[3], en)):
                    if a_ != b_: out['errs'].append(f'{tok}: {n} changed {a_}->{b_}')
                out['ts'].append((typ, 'token override DestGasOverhead', int(c[1]), int(gas)))
        elif fn == 'OnRamp.applyDestChainConfigUpdates':
            live = livejson(to, 'getDestChainConfig(uint64)((address,uint64,uint8,bool,uint16,uint16,uint32,address,address[],address[],bytes))', url, str(T))
            new = cast('calldata-decode', 'x(' + '(uint64,address,uint8,bool,uint16,uint16,uint32,address[],address[],address,bytes)[])', data)
            m = re.match(r'\[\((\d+), (0x\w{40}), (\d+), (true|false), (\d+), (\d+), (\d+), \[(.*?)\], \[(.*?)\], (0x\w{40}), (0x\w*)\)\]', new, re.S)
            g = m.groups()
            # live getter order: router, messageNumber, addressBytesLength, tokenReceiverAllowed, msgFee, tokFee, baseExec, defaultExecutor, laneMandatedCCVs, defaultCCVs, offRamp
            if int(g[0]) != T: out['errs'].append('wrong dest chain')
            checks = [('router', live[0], g[1]), ('addressBytesLength', live[2], g[2]), ('tokenReceiverAllowed', live[3], g[3]),
                      ('messageNetworkFeeUSDCents', live[4], g[4]), ('tokenNetworkFeeUSDCents', live[5], g[5]),
                      ('defaultExecutor', live[7], g[9]), ('offRamp', live[10], g[10])]
            for n, a_, b_ in checks:
                if str(a_).lower() != str(b_).lower(): out['errs'].append(f'OnRamp {n} changed {a_}->{b_}')
            if sorted(x.lower() for x in live[9]) != sorted(re.findall(r'0x\w{40}', g[7].lower())): out['errs'].append('defaultCCVs changed')
            if sorted(x.lower() for x in live[8]) != sorted(re.findall(r'0x\w{40}', g[8].lower())): out['errs'].append('laneMandatedCCVs changed')
            out['t'] = (typ, 'BaseExecutionGasCost', int(live[6]), int(g[6]))
        elif fn == 'FeeQuoter.applyDestChainConfigUpdates':
            live = livejson(to, 'getDestChainConfig(uint64)((bool,uint32,uint32,uint32,uint8,bytes4,uint16,uint32,uint32,uint16,uint8))', url, str(T))
            new = cast('calldata-decode', 'x((uint64,(bool,uint32,uint32,uint32,uint8,bytes4,uint16,uint32,uint32,uint16,uint8))[])', data)
            g = re.match(r'\[\((\d+), \((true|false), (\d+), (\d+), (\d+), (\d+), (0x\w+), (\d+), (\d+), (\d+), (\d+), (\d+)\)\)\]', new).groups()
            names = ['isEnabled', 'maxDataBytes', 'maxPerMsgGasLimit', 'destGasOverhead', 'destGasPerPayloadByteBase', 'chainFamilySelector',
                     'defaultTokenFeeUSDCents', 'defaultTokenDestGasOverhead', 'defaultTxGasLimit', 'networkFeeUSDCents', 'linkFeeMultiplierPercent']
            gas = {'destGasOverhead', 'maxPerMsgGasLimit', 'destGasPerPayloadByteBase', 'defaultTokenDestGasOverhead', 'defaultTxGasLimit'}
            if int(g[0]) != T: out['errs'].append('wrong dest chain')
            out['ts'] = []
            for n, cv, v in zip(names, live, g[1:]):
                if n in gas: out['ts'].append((typ, n, int(cv), int(v)))
                elif str(cv).lower().replace('0x', '').replace('true', 'true') != str(v).lower().replace('0x', ''): out['errs'].append(f'FeeQuoter {n} changed {cv}->{v}')
    except Exception as e:
        out['errs'].append('ERR ' + repr(e)[:160])
    return out

with ThreadPoolExecutor(12) as ex: res = list(ex.map(check, ops))


# ---- independent expected-value rules (recomputed from the LIVE current value, not taken from the changeset) ----
import math
def rnd(x): return int(math.floor(x + 0.5))  # Go math.Round (half away from zero) for positive values
def rule(cur, prague, glam, ratio_num, ratio_den):
    if cur == glam: return cur                 # already migrated -> no-op
    if cur == prague: return glam              # baseline match -> literal Glamsterdam value
    return rnd(cur * ratio_num / ratio_den)    # mismatch -> ratio fallback
EXPECT = {
    ('OnRamp', 'BaseExecutionGasCost'): lambda c: rule(c, 200_000, 400_000, 2, 1),
    ('FeeQuoter', 'destGasOverhead'): lambda c: rule(c, 300_000, 500_000, 5, 3),
    ('FeeQuoter', 'defaultTokenDestGasOverhead'): lambda c: rule(c, 90_000, 270_000, 3, 1),
    ('FeeQuoter', 'defaultTxGasLimit'): lambda c: rule(c, 200_000, 400_000, 2, 1),
    ('FeeQuoter', 'destGasPerPayloadByteBase'): lambda c: rule(c, 20, 64, 64, 20),
    ('FeeQuoter', 'maxPerMsgGasLimit'): lambda c: c,   # 15M -> 15M; fallback is a no-op
    ('CommitteeVerifier', 'GasForVerification'): lambda c: rule(c, 75_000, 85_000, 85, 75),
    ('CCTPVerifier', 'GasForVerification'): lambda c: rule(c, 200_000, 600_000, 3, 1),
    ('LombardVerifier', 'GasForVerification'): lambda c: rule(c, 275_000, 825_000, 3, 1),
    ('CCTPThroughCCVTokenPool', 'DestGasOverhead'): lambda c: rule(c, 250_000, 750_000, 3, 1),
    ('SiloedUSDCTokenPool', 'DestGasOverhead'): lambda c: rule(c, 250_000, 750_000, 3, 1),
    ('LombardTokenPool', 'DestGasOverhead'): lambda c: rule(c, 410_000, 1_200_000, 1_200_000, 410_000),
    # per-token FeeQuoter override: generic x3 (USDC literal 250k->750k equals x3; a USDC token already at 750k, or a Lombard
    # token, would differ and is flagged below for manual review)
    ('FeeQuoter', 'token override DestGasOverhead'): lambda c: 3 * c,
}
for r in res:
    for (typ_, field, cur_, new_) in ([r['t']] if 't' in r else []) + r.get('ts', []):
        exp = EXPECT.get((typ_, field))
        if exp is None: r['errs'].append(f'no expected-value rule for {typ_}.{field}')
        elif exp(cur_) != new_: r['errs'].append(f'{typ_}.{field}: current {cur_} -> proposed {new_}, rule expects {exp(cur_)}')

print('\ntx count by function / target contract type:')
for k, n in sorted(Counter((r['fn'], r['typ']) for r in res).items(), key=str): print(' ', n, k)
print('\nvalue transitions  (contract, field, current -> new): count')
ts = [r['t'] for r in res if 't' in r] + [t for r in res for t in r.get('ts', [])]
for k, n in sorted(Counter(ts).items(), key=str): print(f'  {k[0]}.{k[1]}: {k[2]} -> {k[3]}   x{n}')
bad = [r for r in res if r['errs']]
print(f'\nFAILED: {len(bad)} / {len(res)} txs')
for r in bad[:25]: print(' ', r['sel'], r['to'], r['typ'], r['errs'])
sys.exit(1 if bad else 0)
```

## Appendix B. `verify_coverage.py` — nothing missing, nothing extra

Same status as Appendix A (a helper to copy out, not part of the repo). It recomputes from **live**
chain state which writes the proposal should contain and compares in both directions:

```bash
python3 verify_coverage.py chainlink-deployments/domains/ccv/<env> <target chain selector>
```

- A chain has a lane when its FeeQuoter 2.0.0 has an enabled dest config for the target (the same
  definition the changeset uses) and an OnRamp 2.0.0 in the datastore; the proposal's chain set must
  equal that set.
- On each chain: the OnRamp / CommitteeVerifier / CCTPVerifier / LombardVerifier write is present
  exactly when that contract's live router for the target is non-zero; token pool writes are present
  exactly when the pool supports the target and its fee config is enabled; the FeeQuoter overrides
  cover every enabled override.
- `WARNING` lines list chains whose lane status could not be read through the RPC proxy (they are
  **not** verified, whether or not they appear in the proposal). On `prod_testnet` this includes
  five datastore chains that are not in the `ccv` network config at all, so the changeset cannot
  reach them either; one of them (`superseed-testnet`) has a live lane to Sepolia. Report these.

```python
#!/usr/bin/env python3
"""Coverage check for a Glamsterdam v2.0 proposal: recompute from LIVE chain state which writes the proposal
should contain, and compare with the proposal in both directions (nothing missing, nothing extra).
usage: python3 verify_coverage.py <domains/ccv/<env> dir> <target selector>
Needs: `cast` (foundry), VPN (reads live values from https://rpcs.cldev.sh/<selector>)."""
import base64, glob, json, re, subprocess, sys, time
from collections import defaultdict
from concurrent.futures import ThreadPoolExecutor

BASE = sys.argv[1].rstrip('/') + '/'
T = int(sys.argv[2])
prop = json.load(open(sorted(glob.glob(BASE + 'proposals/*glamsterdam_v2*.json'))[-1]))
refs = json.load(open(BASE + 'datastore/address_refs.json'))
by_chain = defaultdict(list)
for r in refs: by_chain[r['chainSelector']].append(r)

def call(sel, to, sig, *args):
    """cast call with retries for transient RPC-proxy failures; returns the decoded JSON output."""
    cmd = ['cast', 'call', to, sig, *args, '--json', '--rpc-url', f'https://rpcs.cldev.sh/{sel}']
    for i in range(4):
        p = subprocess.run(cmd, capture_output=True, text=True, timeout=90)
        if p.returncode == 0: return json.loads(p.stdout)
        if i == 3 or not re.search(r'header for number not found|ErrUpstreamsExhausted|timed out|timeout|502|503|504|connection', p.stderr, re.I):
            raise RuntimeError(p.stderr.strip()[:120])
        time.sleep(2 * (i + 1))

def ref(sel, typ, versions=None):
    return [r for r in by_chain[sel] if r['type'] == typ and (versions is None or r['version'] in versions)]

# what the proposal actually touches, per chain: set of (type)
touched = defaultdict(lambda: defaultdict(set))   # chain -> type -> {addresses}
tx_count = 0
for o in prop['operations']:
    sel = int(o['chainSelector'])
    for t in o['transactions']:
        r = next((x for x in by_chain[sel] if x['address'].lower() == t['to'].lower()), None)
        touched[sel][r['type'] if r else '?'].add(t['to'].lower()); tx_count += 1
print(f'proposal: {len(prop["operations"])} batches, {tx_count} txs, {len(touched)} chains')

# --- expected lane set, mirroring DiscoverLanesToTarget: FeeQuoter 2.0.0 dest config for the target is enabled
fq_chains = {sel: ref(sel, 'FeeQuoter', ['2.0.0']) for sel in by_chain if sel != T}
fq_chains = {s: v for s, v in fq_chains.items() if v}
def lane(sel):
    fq = fq_chains[sel][0]['address']
    try: return sel, bool(call(sel, fq, 'getDestChainConfig(uint64)((bool,uint32,uint32,uint32,uint8,bytes4,uint16,uint32,uint32,uint16,uint8))', str(T))[0][0]), None
    except Exception as e: return sel, None, str(e)
with ThreadPoolExecutor(10) as ex: lanes = list(ex.map(lane, fq_chains))
read_err = [(s, e) for s, ok, e in lanes if ok is None]
L = {s for s, ok, _ in lanes if ok}
no_onramp = {s for s in L if not ref(s, 'OnRamp', ['2.0.0'])}
expected_chains = L - no_onramp
P = set(touched)
unreadable = {s for s, _ in read_err}   # lane status could not be read through the RPC proxy
print(f'chains with a v2.0 FeeQuoter: {len(fq_chains)} | with an enabled dest config for the target (lane): {len(L)} | lane but no OnRamp 2.0.0 in datastore: {len(no_onramp)} | unreadable via RPC proxy: {len(unreadable)}')
problems = []
if (P - unreadable) != expected_chains:
    problems.append(f'chain set differs: in proposal only {sorted(P - unreadable - expected_chains)}; expected but missing {sorted(expected_chains - P)}')
warnings = []
if unreadable:
    warnings.append(f'could not read the lane status of {len(unreadable)} chains, so they are NOT verified: ' +
                    ', '.join(f'{s} ({"in" if s in P else "not in"} proposal)' for s in sorted(unreadable)))

# --- per-chain expectations
def _chain_check(sel):
    errs = []
    t = touched.get(sel, {})
    fq = fq_chains[sel][0]['address'].lower()
    if fq not in t.get('FeeQuoter', set()): errs.append('no FeeQuoter write')
    # OnRamp / verifiers: written iff the live router for the target is non-zero
    def expect_router(typ, versions, getter_sig, getter_idx):
        for r in ref(sel, typ, versions):
            try: router = call(sel, r['address'], getter_sig, str(T))[0]
            except Exception as e: errs.append(f'{typ} read error {e}'); continue
            router = router[getter_idx] if isinstance(router, list) else router
            wrote = r['address'].lower() in t.get(typ, set())
            if (int(router, 16) != 0) != wrote: errs.append(f'{typ} {r["address"]}: live router {"set" if int(router,16) else "ZERO"} but proposal {"writes" if wrote else "does not write"} it')
    expect_router('OnRamp', ['2.0.0'], 'getDestChainConfig(uint64)((address,uint64,uint8,bool,uint16,uint16,uint32,address,address[],address[],bytes))', 0)
    expect_router('CommitteeVerifier', ['2.0.0'], 'getRemoteChainConfig(uint64)((address,uint64,bool,uint16,uint32,uint16),address[])', 0)
    expect_router('CCTPVerifier', ['2.0.0', '2.1.0'], 'getRemoteChainConfig(uint64)((address,uint64,bool,uint16,uint32,uint16),address[])', 0)
    expect_router('LombardVerifier', ['2.0.0', '2.1.0'], 'getRemoteChainConfig(uint64)((address,uint64,bool,uint16,uint32,uint16),address[])', 0)
    # token pools: written iff the pool supports the target and its fee config for it is enabled
    for typ, versions in (('CCTPThroughCCVTokenPool', ['2.0.0']), ('SiloedUSDCTokenPool', ['2.0.0']), ('LombardTokenPool', ['2.0.0', '2.1.0'])):
        for r in ref(sel, typ, versions):
            try:
                supported = int(T) in [int(x) for x in call(sel, r['address'], 'getSupportedChains()(uint64[])')[0]]
                cfg = call(sel, r['address'], 'getTokenTransferFeeConfig(address,uint64,bytes4,bytes)((uint32,uint32,uint32,uint32,uint16,uint16,bool))',
                           '0x' + '0' * 40, str(T), '0x00000000', '0x')[0] if supported else None
            except Exception as e:
                # reverting getters (e.g. CCTPThroughCCV for an unconfigured destination) mean "not written"
                supported, cfg = False, None
            should = bool(supported and cfg and cfg[6])
            wrote = r['address'].lower() in t.get(typ, set())
            if should != wrote and not (should and cfg and cfg[0] in (750_000, 1_200_000)):  # already-migrated no-op is legitimate
                errs.append(f'{typ} {r["address"]}: expected write={should} but proposal write={wrote}')
    # FeeQuoter per-token overrides: a second FeeQuoter batch exists iff there is >=1 enabled override
    all_ = call(sel, fq_chains[sel][0]['address'], 'getAllTokenTransferFeeConfigs()(uint64[],address[][],(uint32,uint32,uint32,bool)[][])')
    n_over = sum(1 for d, tl, cl in zip(*all_) if int(d) == T for c in cl if c[3])
    return sel, errs, n_over

unverified = []
def chain_check(sel):
    try: return _chain_check(sel)
    except Exception as e:                     # persistent RPC trouble: report, don't crash
        unverified.append((sel, str(e)[:80])); return sel, [], 0

with ThreadPoolExecutor(10) as ex: results = list(ex.map(chain_check, sorted(expected_chains & P)))
for sel, errs, n in results:
    for e in errs: problems.append(f'chain {sel}: {e}')
print(f'checked {len(results)} chains; enabled overrides for target across them: {sum(n for _, _, n in results)}')
if no_onramp: print('lane but no OnRamp (skipped by the changeset, reported as "could not resolve OnRamp"):', sorted(no_onramp))
for w in warnings: print('WARNING:', w)
if unverified: print('WARNING: per-chain checks could not run (RPC errors) for:', unverified)
print(f'\nCOVERAGE PROBLEMS: {len(problems)}')
for p in problems: print('  ', p)
sys.exit(1 if problems else 0)
```
