# Token sequences (v2.0.0)

This package contains the sequences used to configure tokens and token pools for CCIP transfers on EVM chains at version 2.0.0. It includes logic to **import configuration from a prior (active) pool** when upgrading a token to a new pool while keeping the same registry and token.

## Configure tokens for transfers: flow

1. **ConfigureTokensForTransfers** – Top-level sequence. Validates the pool supports the token, sets min block confirmations, then calls **ConfigureTokenPoolForRemoteChains** with the full `RemoteChains` map, and finally registers the token with the TokenAdminRegistry.

2. **ConfigureTokenPoolForRemoteChains** – When `RegistryAddress` and `TokenAddress` are set and the active pool differs from `TokenPoolAddress` (upgrade path), validates that `RemoteChains` includes every chain currently supported by the **active pool** (the pool currently registered for that token). When the active pool is already the pool being configured (extend path), that check is skipped so you only need to list new remotes. Then calls **ConfigureTokenPoolForRemoteChain** once per remote chain.

3. **ConfigureTokenPoolForRemoteChain** – For a single remote chain: optionally **imports rate limiter and remote pool config from the active pool**, then applies CCV config, rate limiters, and remote chain config (add/update chain, remote pools). When `tokenTransferFeeConfig` is set on the remote chain input, applies it on the **v2 token pool** only (merge with on-chain pool state or defaults; no legacy lane import). Fee config runs **after** the remote chain exists on the pool.

   The **`ConfigureTokensForTransfers` changeset** strips fee config from sequence input and applies fees after configure (including legacy-lane discovery and merge when `autoMigrateRemoteChains` is enabled). Direct callers (e.g. CCTP) that invoke this sequence with explicit fee config still use the pool-only apply path above.

## Importing config from prior pool versions

When configuring a **new** 2.0.0 pool for a token that already has an **active pool** registered in the TokenAdminRegistry (e.g. upgrading from 1.5.1 or 1.6.1), the sequence can **import** rate limiter state and remote pool addresses from that active pool so you don’t have to re-specify everything.

### When import runs

- **ConfigureTokenPoolForRemoteChain** receives `RegistryAddress` and `TokenAddress` (passed from ConfigureTokensForTransfers when `RegistryAddress` is set).
- It calls **importConfigFromActivePool(registry, token, remoteChainSelector)**.
- Import is skipped if registry or token is zero, or if the registry returns no active pool for that token, or if the active pool’s version is ≥ 2.0.0.

### What gets imported (per remote chain)

For the given **remote chain selector**, the importer reads from the active pool (or its implementation, see below) and returns:

| Field | Description |
|-------|-------------|
| **DefaultOutbound** | Outbound rate limiter config (capacity, rate, isEnabled) for that chain. |
| **DefaultInbound**  | Inbound rate limiter config for that chain. |
| **RemotePools**     | List of remote pool address(es) for that chain (1.5.0: single pool via `getRemotePool`; 1.6.1: full list via `getRemotePools`). |

### How imported config is used

- **Rate limiters** – If the **input** does not supply default finality rate limiter config (e.g. null or omitted), the sequence uses the **imported** default outbound/inbound config for the new pool. So prior pool’s rate limits are carried over by default unless you define a non-nil rate limit in the config.
- **Remote pools** – When adding or updating the remote chain on the new pool, the sequence builds the remote pool list as: **active pool’s remote pools first** (from import), then the requested pool from config if not already present. That preserves existing remote pools during cutover (e.g. for in-flight messages) and adds the new pool’s counterpart.

### Version-specific import paths

The active pool’s **type and version** (from `typeAndVersion()`) decide which path runs:

- **Active pool version &lt; 1.5.1** → **V150 path** (1.5.0 operations and bindings).
  - If the pool type contains `"Proxy"`: the code resolves **previousPool** from the proxy. Rate limits are read from the **proxy** if previousPool version &lt; 1.4.0, otherwise from **previousPool**. Remote pool is always read from the **active pool** (proxy), which exposes `getRemotePool`.
  - If not a proxy: rate limits and remote pool are read from the active pool using 1.5.0 ops.
  - Uses: `GetInboundRateLimiterState`, `GetOutboundRateLimiterState`, `GetRemotePool` (1.5.0). Token bucket struct is converted to the shared `RateLimiterConfig` type.

- **Active pool version ≥ 1.5.1 and &lt; 2.0.0** → **V161 path** (1.6.1 operations).
  - Rate limits: `GetCurrentInboundRateLimiterState`, `GetCurrentOutboundRateLimiterState` on the active pool.
  - Remote pools: `GetRemotePools` on the active pool (full list per chain).
  - Token bucket is converted to the shared `RateLimiterConfig` type.

- **Active pool version ≥ 2.0.0** → No import (returns `nil`).

### Supported-chains validation (upgrade vs extend)

Before configuring each remote chain, **ConfigureTokenPoolForRemoteChains** can ensure the input doesn’t drop any chain the current deployment relies on when **replacing** the registered pool:

- When `RegistryAddress` and `TokenAddress` are set, it gets the **active pool** from the registry.
- **Extend** (`activePool == TokenPoolAddress`): the check is skipped. `RemoteChains` may list only new remotes; existing supported chains do not need to be repeated.
- **Upgrade** (`activePool != TokenPoolAddress` and active pool is non-zero): it calls **GetSupportedChains** on the active pool and requires every chain in that list to appear in **RemoteChains**. If any are missing, the sequence fails so you don’t accidentally omit config for a chain that was previously supported.
- **Net-new** (no active pool): the check is skipped.
- Set **`SkipActivePoolSupportedChainsCheck`** to bypass the upgrade check (e.g. CCTP parallel-pool setups where the configured pool is not a direct TAR replacement).
- If **GetSupportedChains** on the active pool fails (e.g. USDCTokenPoolProxy), validation is skipped best-effort.

This keeps “configure for transfers” aligned with the active pool’s supported chains when upgrading to a different pool, without requiring a full remote list when extending an already-registered pool.

## Auto-migrate edge cases (`autoMigrateRemoteChains`)

When the **`ConfigureTokensForTransfers` changeset** sets `autoMigrateRemoteChains: true` during a pre-v2 → v2 pool upgrade, remote connectivity and legacy fees are discovered in the changeset (not in these sequences). Rate limits are still imported here via `importConfigFromActivePool` as described above. Fees are applied by the changeset after configure.

### YAML precedence (per remote chain)

| Situation | Behavior |
|-----------|----------|
| Remote **not listed** in YAML | Fully discovered from the legacy active pool (token, pool, decimals, fees). |
| Remote listed, **no** `remoteToken` / `remotePool` | Backfill connectivity from the active pool; YAML overrides fees and other fields. |
| Remote listed with **explicit** `remoteToken` and/or `remotePool` | YAML wins (coordinated retarget); legacy refs are not overwritten. |
| Remote listed but **not** on the legacy active pool | Not enriched by discovery; provide full connectivity in YAML. |

### Fee discovery

- Requires connected CCIP lanes (OnRamp/FeeQuoter resolvable per discovered remote). **Failures abort the entire changeset** — there is no partial apply.
- Legacy lane fees are imported **only when the legacy FeeQuoter config is enabled** for that token/lane. If legacy fees are disabled and YAML omits `tokenTransferFeeConfig`, **no fee transactions** are emitted on the new v2 pool.
- When YAML includes `tokenTransferFeeConfig`, **`isEnabled` must be set** (omit the block entirely if you do not want to configure fees). During upgrade with legacy fees enabled, YAML fields merge over the discovered legacy values; with legacy fees disabled, YAML alone defines the config.
- Without `autoMigrateRemoteChains`, omitted `tokenTransferFeeConfig` leaves fees untouched (same as before).

### Pools without `getSupportedChains` (e.g. USDCTokenPoolProxy)

Discovery calls `getSupportedChains` on the TAR-registered active pool and **fails** if the pool does not implement it. List remote chains explicitly instead.
