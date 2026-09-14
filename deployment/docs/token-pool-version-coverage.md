# Token pool version coverage (upgrading legacy pools → 2.0)

Status of the Tooling API (`chainlink-ccip/deployment` + `chains/<family>/deployment/*/adapters`)
as of 2026-09-14, `chainlink-ccip@5c5da7896`.

Scope: which *existing* token pool versions the tooling can read, configure, and migrate onto a
v2.0.0 pool. This is a coverage report, not a design; gaps are listed at the end.

---

## 1. How version dispatch works

Everything token-related routes through one registry:

- `tokens.TokenAdapterRegistry` (`deployment/tokens/product.go:510`), keyed by
  `fmt.Sprintf("%s-%s", chainFamily, version.String())` — **exact full semver match, no patch
  stripping, no range matching**. (`utils.StripPatchVersion` exists but is used only by the
  1.6 fee adapter, never for token adapter lookup.)
- Lookup entry point: `tokens.ResolveAdapter` (`deployment/tokens/token_expansion.go:645`).
  Miss ⇒ hard error: `no token adapter registered for chain family 'evm' and version 'X.Y.Z'`.
- The version comes from `tokens.ResolveTokenPoolRef` (`token_expansion.go:561`), which:
  1. filters the **datastore** for a matching `AddressRef` — if exactly one hit, its `Version`
     field is used verbatim;
  2. otherwise falls back to the family `TokenRefResolver`, which on EVM reads the pool's
     on-chain **`typeAndVersion()`** (`chains/evm/deployment/v1_0_0/adapters/token_adapter.go:139`).

Consequence: coverage is determined by (a) which versions have a registered adapter, and
(b) what string the datastore/`typeAndVersion()` reports. A pool labelled `1.6.3` in the datastore
will not find the 1.6.1 adapter even though the ABIs are identical.

Adapters register themselves via package `init()`, so a version is only available if the domain
imports its package (directly or transitively). Both `domains/ccip` and `domains/ccv` in
chainlink-deployments pull in all EVM adapter packages via their `resolvers/manager.go`.

---

## 2. Registered token adapters

Verified empirically by probing `GetTokenAdapter` for every candidate version after importing all
EVM adapter packages.

| Family | Version | Implementation | Migrator | RateLimitReader | Fee | Admin | RemotePoolRemover | LR-liquidity migrate |
|---|---|---|---|---|---|---|---|---|
| evm | 1.0.0 | `v1_0_0.EVMTokenBase` | ✗ | ✗ | ✗ | ✗ | ✗ | ✗ |
| evm | 1.5.1 | `v1_5_1.TokenAdapter` | ✓ | ✓ | ✗ | ✓ | ✓ | ✗ |
| evm | 1.6.0 | `v1_6_1.TokenAdapter` | ✓ | ✓ | ✗ | ✓ | ✓ | ✗ |
| evm | 1.6.1 | `v1_6_1.TokenAdapter` | ✓ | ✓ | ✗ | ✓ | ✓ | ✗ |
| evm | 1.6.2 | `v1_6_1.TokenAdapter` | ✓ | ✓ | ✗ | ✓ | ✓ | ✗ |
| evm | 2.0.0 | `v2_0_0.TokenAdapter` | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| evm | 1.7.0 | `v2_0_0.TokenAdapter` — **ccv domain only** | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| solana | — | **none** | — | — | — | — | — | — |
| aptos / sui | — | **none** | — | — | — | — | — | — |

Notes:

- `1.0.0` is `EVMTokenBase`, registered so token-only callers can get an adapter without importing
  a pool package. Its pool methods are **nil-returning stubs**
  (`ConfigureTokenForTransfersSequence`, `SetTokenPoolRateLimits`, `DeployTokenPoolForToken`,
  `ManualRegistration` all return `nil`; `DeriveTokenAddress`/`DeriveTokenDecimals` return errors).
  It is not a usable pool adapter.
- `1.7.0` is registered only by `chainlink-deployments/domains/ccv/pkg/pipelines/shared.go:171`
  (and the canton env), not by any `init()`. It is an alias of the 2.0.0 adapter.
- Solana has an address normalizer and a `TokenAdminRegistryReader` registered
  (`chains/solana/deployment/v1_0_0/adapters/init.go`) but **no `TokenAdapter`**. Solana chains can
  therefore appear as *remotes* during reverse propagation, but cannot themselves be configured or
  upgraded through this API.

---

## 3. What "upgrade to 2.0" requires per version

The upgrade path is `ConfigureTokensForTransfers` with `autoMigrateRemoteChains: true`
(`deployment/tokens/configure_tokens_for_transfers.go:46`), optionally preceded/accompanied by
`TokenExpansion` (deploys the 2.0 pool) and `MigrateLockReleasePoolLiquidity`.

The auto-migrate branch (`configure_tokens_for_transfers.go:259-296`) fires only when **all** hold:

1. TAR has an active pool for the token, and it differs from the configured target pool;
2. the active pool's resolved version `< 2.0.0`;
3. the target pool's version `>= 2.0.0`;
4. the **legacy** adapter implements `TokenPoolMigrator` — otherwise it logs and falls back to
   explicit `remoteChains` YAML (deliberately not a hard failure, for pools like
   `USDCTokenPoolProxy` that cannot implement it);
5. if it implements `TokenPoolMigrator` it **must** also implement `RateLimitReaderAdapter`,
   otherwise hard error.

So the effective source-version support for automatic upgrade is exactly the rows above with
Migrator ✓: **1.5.1, 1.6.0, 1.6.1, 1.6.2**.

Discovery reads on the legacy pool are `getSupportedChains()`, `getRemoteToken(uint64)`,
`getRemotePools(uint64)` — all introduced in the 1.5.1 pool ABI. Anything older lacks them even if
an adapter were registered.

Lock/release liquidity migration (`MigrateLockReleasePoolLiquidity`,
`chains/evm/deployment/v2_0_0/sequences/tokens/migrate_lock_release_pool_liquidity.go`) drives the
old pool through the **1.6.1 bindings** (`getRebalancer`/`setRebalancer`/`withdrawLiquidity`, plus
the siloed variants). Those signatures are shared with 1.5.1, so v1.5.1 and v1.6.1 lock-release
pools both work; the old pool's `typeAndVersion()` type (not version) selects the siloed vs
unsiloed path, and only the exact type `SiloedLockReleaseTokenPool` takes the siloed path.

Target-side (2.0.0) deploy supports these pool types
(`v2_0_0/sequences/tokens/deploy_token_pool.go:240`), everything else errors
`unsupported token pool type`:

- `SiloedLockReleaseTokenPool`
- `LockReleaseTokenPool` (via `utils.IsLockReleasePoolType`)
- `BurnMintTokenPool`, `BurnFromMintTokenPool`, `BurnWithFromMintTokenPool`,
  `BurnToAddressMintTokenPool` (via `utils.IsBurnMintPoolType`)

USDC/CCTP goes through separate adapters (`CCTPChainAdapter`, `USDCTokenPoolProxy`,
`CCTPThroughCCVTokenPool`), not this switch.

---

## 4. Coverage against the deployed fleet

Pool types/versions currently in `chainlink-deployments/domains/ccip/<env>/datastore/address_refs.json`.
"Covered" = an adapter resolves for that version string.

### Covered

| Type / version | mainnet | testnet |
|---|---:|---:|
| BurnMintTokenPool 1.5.1 | 239 | 138 |
| BurnFromMintTokenPool 1.5.1 | 65 | 66 |
| BurnWithFromMintTokenPool 1.5.1 | 76 | 66 |
| LockReleaseTokenPool 1.5.1 | 107 | 74 |
| USDCTokenPool 1.5.1 | 7 | 6 |
| BurnMintWithLockReleaseFlagTokenPool 1.5.1 | 4 | — |
| BurnMintTokenPool 1.6.0 / 1.6.1 | 18 / 42 | 9 / 28 |
| LockReleaseTokenPool 1.6.0 / 1.6.1 | 9 / 9 | 3 / 6 |
| SiloedLockReleaseTokenPool 1.6.0 / 1.6.1 | 1 / 1 | 1 / — |
| BurnMintWithExternalMinterTokenPool 1.6.0 / 1.6.1 | 10 / 1 | 2 / — |
| BurnWithFromMintTokenPool 1.6.1 | 4 | — |
| BurnMintWithLockReleaseFlagTokenPool 1.6.1 | 5 | 4 |
| HybridWithExternalMinterTokenPool 1.6.0 | 2 | 1 |
| XERC20LockboxTokenPool / SiloedWithUnsiloedXERC20Group… 1.6.0 | 1 / 1 | — |
| CCTPTokenPool 1.6.0 | 1 | 1 |
| HybridLockReleaseUSDCTokenPool 1.6.2 | 1 | 1 |
| USDCTokenPool 1.6.2 | 7 | 7 |
| already-2.0.0 (USDCTokenPoolProxy, SiloedUSDCTokenPool, CCTPThroughCCVTokenPool, ERC20LockBox) | 26 | 21 |

Caveat: "covered" here means *the adapter resolves*. For the 1.6.x types whose ABI diverges from
the plain `token_pool` binding (XERC20 lockbox, external-minter, hybrid USDC), only the generic
`getSupportedChains`/`getRemoteToken`/`getRemotePools`/`getToken` surface is exercised; the
type-specific state (lockbox, external minter, hybrid mode) is **not** carried forward by
auto-migrate and needs separate handling.

### Not covered — no adapter registered for the version string

| Type / version | mainnet | testnet | Failure |
|---|---:|---:|---|
| USDCTokenPool 1.6.5 | 7 | 7 | `no token adapter registered … '1.6.5'` |
| USDCTokenPoolCCTPV2 1.6.5 | 8 | 8 | same |
| USDCTokenPoolProxy 1.6.4 | — | 2 | `… '1.6.4'` |
| BurnMintTokenPoolAndProxy 1.5.0 | 2 | — | `… '1.5.0'` |
| BurnMintTokenPool 1.5.0 | — | 1 | `… '1.5.0'` |
| AptosManagedTokenPool / AptosRegulatedTokenPool 1.6.0 | 6 | 3 | no Aptos family adapter |
| Sui\* 1.0.0 | — | 8 refs | no Sui family adapter |
| Solana pools (`TokenPoolLookupTable` 1.6.0 etc.) | 43 | 16 | no Solana `TokenAdapter` |

### Not covered — pre-1.5 pools (the 1.2 / 1.4 generation)

These do not appear in `datastore/address_refs.json` at all; they live only in
`domains/ccip/<env>/addresses.json`, recorded with a placeholder version **`1.0.0`**:

| Type / version (address book) | mainnet | testnet |
|---|---:|---:|
| BurnMintTokenPool 1.0.0 | 18 | 9 |
| LockReleaseTokenPool 1.0.0 | 9 | 3 |
| CCTPTokenPool 1.0.0 | 1 | 1 |

Two distinct failure modes, depending on which source wins in `ResolveTokenPoolRef`:

- **Datastore hit on a `1.0.0` ref** → resolves to `EVMTokenBase`, whose pool methods return
  `nil` sequences. This is the worst case: no clean "unsupported version" error, just a nil
  sequence handed to the executor.
- **Datastore miss → on-chain `typeAndVersion()`** → returns `1.2.0` / `1.4.0` → clean hard error
  `no token adapter registered for chain family 'evm' and version '1.2.0'`.

There are generated bindings for pre-1.5 pools (`chains/evm/gobindings/generated/v1_2_0/…`,
`…/v1_4_0/{burn_mint,lock_release,usdc,token}_pool`) but **no `chains/evm/deployment/v1_4_0`
package at all** and no token-pool operations under `v1_2_0` — so there is no operations layer to
build an adapter on today.

---

## 5. Gaps, ranked

1. **Pre-1.5 (1.2 / 1.4) pools: zero support, and the 1.0.0 placeholder makes it fail badly.**
   Needs (a) `chains/evm/deployment/v1_2_0|v1_4_0/operations/token_pool` ops, (b) adapters, and
   (c) a decision on migration semantics — these pools have no `getSupportedChains`/`getRemotePools`,
   so `TokenPoolMigrator` cannot be implemented against them; remotes would have to be reconstructed
   from `applyRampUpdates` history or supplied explicitly in YAML. Interim mitigation: make
   `EVMTokenBase`'s nil-returning stubs return explicit errors so a mis-resolved `1.0.0` ref fails
   loudly instead of handing a nil sequence downstream.
2. **1.5.0** (`BurnMintTokenPool`, `BurnMintTokenPoolAndProxy`, 3 pools total). Ops exist
   (`v1_5_0/operations/token_pool`); only an adapter + registration is missing. `getRemotePool`
   is singular in 1.5.0, so `GetRemotePools` needs a 1.5.0-specific shim.
3. **1.6.3 / 1.6.4 / 1.6.5** (30 pools, all USDC/CCTP). The 1.6.5 ops package already exists
   (`v1_6_5/operations/usdc_token_pool_cctp_v2` has `GetSupportedChains`/`GetRemoteToken`/
   `GetRemotePools`). Cheapest fix for the plain cases is registering the existing 1.6.1 adapter
   under these version strings the way 1.6.2 already does
   (`chains/evm/deployment/v1_6_2/adapters/init.go`) — but USDC pools also need the CCTP path, so
   validate rather than blanket-alias.
4. **Exact-match version lookup is brittle.** Every new patch release of an ABI-identical pool
   needs a manual alias registration. A fallback in `TokenAdapterRegistry.GetTokenAdapter` — exact
   match, then `StripPatchVersion`, then highest registered `<=` requested — would close the
   1.6.3/1.6.4/1.6.5 class permanently. Behaviour change; needs a deliberate call.
5. **Non-EVM families.** Solana has no `TokenAdapter` at any version; Aptos and Sui have neither
   adapters nor normalizers. Any token whose web includes these chains can only be upgraded on its
   EVM legs.
6. **Type-specific 1.6.x state is not migrated.** Auto-migrate carries remote chains, remote
   tokens, remote pools, decimals and rate limits only. XERC20 lockbox config, external-minter
   wiring, and hybrid-USDC mode are not read or reapplied.
7. **`USDCTokenPoolProxy` cannot implement `TokenPoolMigrator`** (documented at
   `configure_tokens_for_transfers.go:94` and `:267`). Upgrades behind the proxy require remote
   chains listed explicitly in YAML. Worth stating in the runbook rather than "fixing".

---

## 6. Reproducing the adapter table

The registry has no enumeration API; probe it. Temporary test under `chains/evm` (own Go module):

```go
// chains/evm/deployment/<tmp>/probe_test.go
import (
    _ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
    // … one blank import per vX_Y_Z/adapters package …
    "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
)

reg := tokens.GetTokenAdapterRegistry()
for _, v := range []string{"1.0.0", "1.2.0", "1.4.0", "1.5.0", "1.5.1", "1.6.0", "1.6.1",
    "1.6.2", "1.6.3", "1.6.4", "1.6.5", "1.7.0", "2.0.0", "2.1.0"} {
    a, ok := reg.GetTokenAdapter("evm", semver.MustParse(v))
    if !ok { continue }
    _, migrator := a.(tokens.TokenPoolMigrator)
    _, rlReader := a.(tokens.RateLimitReaderAdapter)
    fmt.Printf("%s %T migrator=%v rl=%v\n", v, a, migrator, rlReader)
}
```

Fleet counts come from
`domains/ccip/<env>/datastore/address_refs.json` and `domains/ccip/<env>/addresses.json` in
chainlink-deployments.
