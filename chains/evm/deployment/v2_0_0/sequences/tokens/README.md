# Token sequences (v2.0.0)

This package contains the sequences used to configure tokens and token pools for CCIP transfers on EVM chains at version 2.0.0. It includes logic to **import configuration from a prior (active) pool** when upgrading a token to a new pool while keeping the same registry and token.

## Configure tokens for transfers: flow

1. **ConfigureTokensForTransfers** – Top-level sequence. Validates the pool supports the token, sets min block confirmations, then calls **ConfigureTokenPoolForRemoteChains** with the full `RemoteChains` map, and finally registers the token with the TokenAdminRegistry.

2. **ConfigureTokenPoolForRemoteChains** – When `RegistryAddress` and `TokenAddress` are set (upgrade path), validates that `RemoteChains` includes every chain currently supported by the **active pool** (the pool currently registered for that token). Then calls **ConfigureTokenPoolForRemoteChain** once per remote chain.

3. **ConfigureTokenPoolForRemoteChain** – For a single remote chain: optionally **imports config from the active pool**, then applies CCV config, rate limiters, remote chain config (add/update chain, remote pools), and token transfer fee config. Token transfer fee config runs **after** remote chain config so the chain exists on the pool before setting fees.

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

- **Rate limiters** – If the **input** does not supply default/custom finality rate limiter config (e.g. `IsEnabled: false` or omitted), the sequence uses the **imported** default outbound/inbound config for the new pool. So prior pool’s rate limits are carried over unless you override them in the config.
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

### Supported-chains validation (upgrade path)

Before configuring each remote chain, **ConfigureTokenPoolForRemoteChains** ensures the input doesn’t drop any chain the current deployment relies on:

- When `RegistryAddress` and `TokenAddress` are set, it gets the **active pool** from the registry and calls **GetSupportedChains** on it.
- It then checks that every chain in that list appears in **RemoteChains**.
- If any active-pool chain is missing from `RemoteChains`, the sequence fails with an error so you don’t accidentally omit config for a chain that was previously supported.

This keeps “configure for transfers” aligned with the active pool’s supported chains when upgrading.
