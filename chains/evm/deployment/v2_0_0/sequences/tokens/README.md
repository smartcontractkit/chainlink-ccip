# Token sequences (v2.0.0)

This package contains the sequences used to configure tokens and token pools for CCIP transfers on EVM chains at version 2.0.0.

## Configure tokens for transfers: flow

1. **ConfigureTokensForTransfers** – Top-level sequence. Validates the pool supports the token, sets min block confirmations, then calls **ConfigureTokenPoolForRemoteChains** with the full `RemoteChains` map, and finally registers the token with the TokenAdminRegistry.

2. **ConfigureTokenPoolForRemoteChains** – When `RegistryAddress` and `TokenAddress` are set and the active pool differs from `TokenPoolAddress` (upgrade path), validates that `RemoteChains` includes every chain currently supported by the **active pool** (the pool currently registered for that token). When the active pool is already the pool being configured (extend path), that check is skipped so you only need to list new remotes. Then calls **ConfigureTokenPoolForRemoteChain** once per remote chain.

3. **ConfigureTokenPoolForRemoteChain** – For a single remote chain: applies CCV config, rate limiters, and remote chain config (add/update chain, remote pools). Legacy upgrade baggage (rate limits, remote pool cutover) is consumed from **`MigrationMetadata`** on `RemoteChainConfig` when populated by the changeset. When `tokenTransferFeeConfig` is set on the remote chain input, applies it on the **v2 token pool** only (merge with on-chain pool state or defaults). Fee config runs **after** the remote chain exists on the pool.

   The **`ConfigureTokensForTransfers` changeset** strips fee config from sequence input and applies fees after configure (including legacy-lane discovery and merge when `autoMigrateRemoteChains` is enabled). Direct callers (e.g. CCTP) that invoke this sequence with explicit fee config still use the pool-only apply path above.

## Migration metadata (upgrade path)

During a pre-v2 → v2 pool upgrade with `autoMigrateRemoteChains: true`, the **changeset** discovers legacy lane config and populates `RemoteChainConfig.MigrationMetadata` (rate limits, remote pools, pool version/type). The v2 sequence **consumes** that metadata at apply time; it does not read the legacy active pool itself.

### Rate limiters

Resolved in order:

1. Both default YAML buckets → `GenerateTPRLConfigs`
2. `MigrationMetadata.LegacyRateLimits` → bigint passthrough (inbound already rebased at discovery)
3. Both omitted → on-chain read when remote already supported, else disabled defaults

### Remote pools

When adding or updating a remote chain, the sequence builds the remote pool list as: **`LegacyRemotePools` first** (cutover / inflight protection), then the requested `RemotePool` from config if not already present.

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

When the **`ConfigureTokensForTransfers` changeset** sets `autoMigrateRemoteChains: true` during a pre-v2 → v2 pool upgrade, remote connectivity, rate limits, remote pools, and legacy fees are discovered in the changeset. The v2 sequence applies connectivity, RL, and remote pools from `RemoteChainConfig` + `MigrationMetadata`. Fees are applied by the changeset after configure.

### YAML precedence (per remote chain)

| Situation | Behavior |
|-----------|----------|
| Remote **not listed** in YAML | Fully discovered from the legacy active pool (token, pool, decimals); fees imported only when legacy FeeQuoter lane config is enabled. |
| Remote listed, **both** `remoteToken` and `remotePool` empty | Backfill token, pool, and decimals from the active pool; YAML overrides fees and other fields when `tokenTransferFeeConfig` is set (`isEnabled` required). |
| Remote listed with **only one** of `remoteToken` / `remotePool` | No connectivity backfill: the set field is kept, the missing field is **not** imported. Provide both explicitly or leave both empty. |
| Remote listed with **both** `remoteToken` and `remotePool` set | YAML wins (coordinated retarget); legacy refs are not overwritten. |
| Remote listed but **not** on the legacy active pool | Not enriched by discovery; provide full connectivity in YAML. |

### When discovery is skipped

`autoMigrateRemoteChains` is ignored (info log, no error) when there is no legacy active pool to migrate from, the TAR active pool is already the target pool (extend), or the active pool is already v2.0.0+. In those cases, configure remotes and fees explicitly in `RemoteChains`. Remove the flag from YAML after a one-time pre-v2 → v2 upgrade.

### Fee discovery

- Requires connected CCIP lanes (OnRamp/FeeQuoter resolvable per discovered remote). **Failures abort the entire changeset** — there is no partial apply.
- Legacy lane fees are imported **only when the legacy FeeQuoter config is enabled** for that token/lane. If legacy fees are disabled and YAML omits `tokenTransferFeeConfig`, **no fee transactions** are emitted on the new v2 pool.
- When YAML includes `tokenTransferFeeConfig`, **`isEnabled` must be set** (omit the block entirely if you do not want to configure fees). During upgrade with legacy fees enabled, YAML fields merge over the discovered legacy values; with legacy fees disabled, YAML alone defines the config.
- Without `autoMigrateRemoteChains`, omitted `tokenTransferFeeConfig` leaves fees untouched (same as before).

### Pools without `getSupportedChains` (e.g. USDCTokenPoolProxy)

Discovery calls `getSupportedChains` on the TAR-registered active pool and **fails** if the pool does not implement it. List remote chains explicitly instead.
