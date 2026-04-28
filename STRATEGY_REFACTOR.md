# Per-Token-Type Strategy Refactor

## Objective

Two related complaints in the EVM token-expansion code:

1. The token-pool deployment adapter (`chains/evm/deployment/v1_0_0/adapters/pool_adapter.go`) had a switch on token type to decide which role-grant operation to call. This switch grew with each new token type. The same token-type dispatch pattern repeated in three other places — token deployment, capability predicates, and the external-admin role grant.
2. The v2.0 adapter (`chains/evm/deployment/v2_0_0/adapters/tokens.go`) reimplemented the role-grant logic inline rather than reusing v1.0.0 utilities. New v1.6 token types did not flow through to v2.0 automatically.

The goal: encapsulate everything specific to a token contract type behind one per-token strategy, registered into a registry that is independent of pool version. The four switches collapse to registry lookups, and v2.0 picks up new BurnMint-family token types automatically.

## Direction

A `EVMTokenStrategy` interface lives in `chains/evm/deployment/tokens/strategy/`:

```go
type EVMTokenStrategy interface {
    ContractType() deployment.ContractType
    Capabilities() Capabilities

    Deploy(b cldf_ops.Bundle, chain evm.Chain, in tokensapi.DeployTokenInput) (
        datastore.AddressRef, []evm_contract.WriteOutput, error)

    GrantPoolRoles(b cldf_ops.Bundle, chain evm.Chain,
        token, pool common.Address, chainSelector uint64) ([]evm_contract.WriteOutput, error)

    GrantExternalAdmin(b cldf_ops.Bundle, chain evm.Chain,
        token, externalAdmin common.Address, chainSelector uint64) ([]evm_contract.WriteOutput, error)
}
```

Strategies are registered into a singleton `Registry` keyed by `ContractType` only — pool version is deliberately excluded.

All five known EVM token types (plain ERC20, BurnMintERC20, BurnMintERC20WithDrip v1.0.0, BurnMintERC20WithDrip v1.5.0, TIP-20) are wrapped as small strategy structs in one file at `chains/evm/deployment/tokens/strategy/registrations/registrations.go`. Adapters pick them up via a single blank import.

The four dispatch sites become registry lookups:

| Site | Now becomes |
|---|---|
| `pool_adapter.go` role-grant switch | `strat.GrantPoolRoles(...)` |
| `sequences/token.go` deploy switch | `strat.Deploy(...)` |
| `sequences/token.go` capability predicates | `strat.Capabilities().Supports*` |
| `sequences/token.go` external-admin switch | `strat.GrantExternalAdmin(...)` |

The v2.0 adapter's `DeployTokenPoolForToken` keeps its v2-specific pool-type dispatch but its inline role-grant block becomes the same registry lookup as v1.0.0. The previous `isBurnMintTokenType` filter is deleted.

## Key design decisions

**Registry key is `(chainFamily, ContractType)`, not pool version.** This is the entire point of the refactor — adding a new BurnMint-family token type to the registry makes it available on v1.5.1, v1.6.x, and v2.0 simultaneously, with one line in `registrations.go`.

**Capabilities is a struct with four explicit flags, including `ParticipatesInPoolRoleGrant`.** Two Plan agents reviewed the design independently; both converged on the explicit flag rather than inferring non-participation from a no-op `GrantPoolRoles` return. Audit reads the flag, not a side effect.

**TIP-20 stays v1.6-only via an explicit guard inside `v2_0_0/adapters/tokens.go`.** The strategy registry is version-independent (TIP-20 is registered globally), but the v2.0 adapter rejects TIP-20 early with a clear error message. The v1.6-only constraint is locally checkable in the v2.0 file rather than implicit in pool-type gates elsewhere.

**Strategies live in one file, ops packages stay untouched.** All five strategy structs are colocated in `registrations/registrations.go`. Ops packages (`burn_mint_erc20`, `tip20`, etc.) are not modified — strategies are thin wrappers that compose existing operations. Adding a new token type means: add an ops package (as today), add a strategy struct, add one Register call. No edits to any pool adapter.

**Default policies are preserved per call site.** Deploy lookup miss is fail-fast; role-grant lookup miss is warn-and-continue; capability lookup miss returns the zero-value struct (all-false). Three different policies, all matching the prior behavior verbatim.

**Pool-type dispatch (`isBurnMintPoolType` etc.) is deliberately untouched.** Token type and pool type are orthogonal; conflating them was rejected in design review.

## What changed

New (3 files): `chains/evm/deployment/tokens/strategy/{strategy.go, registry.go, registrations/registrations.go}`.

Edited:
- `chains/evm/deployment/v1_0_0/adapters/pool_adapter.go` — role-grant switch replaced with registry call; unused ops imports removed.
- `chains/evm/deployment/v1_0_0/sequences/token.go` — deploy switch, capability predicates, and external-admin switch all replaced with registry lookups; three `tokenSupports*` predicate funcs deleted.
- `chains/evm/deployment/v2_0_0/adapters/tokens.go` — TIP-20 guard added; role-grant tail replaced with registry call; `isBurnMintTokenType` deleted.
- `chains/evm/deployment/v1_0_0/adapters/init.go` — blank import of `registrations` so strategies are loaded for all transitively-importing adapters.

`go build ./...` and `go vet ./...` are clean across all submodules of the experiments workspace.

## What's deferred

Tests are deferred to a follow-up pass. Once the implementation direction is confirmed, the test pass should add:
- Registry unit tests
- Golden capability-table test asserting every existing token type's capabilities match the pre-refactor `tokenSupports*` truth tables
- TIP-20 v2.0 guard test
- An end-to-end test that pre-registers a fake `ContractType` strategy and exercises both v1.0 and v2.0 dispatch backbones

The existing token deployment test (`v1_0_0/sequences/token_test.go`) was lightly updated to call the new registry instead of the deleted predicate so the package still compiles.
