---
title: "Consuming the Tooling API"
sidebar_label: "Consuming"
sidebar_position: 3
---

# Consuming the Tooling API

This page is for people who **run** changesets: durable pipelines in `chainlink-deployments`, devenv, and integration tests. For adding a chain family, see [Implementing 1.6](implementing-1.6.md) and [Implementing 2.0](implementing-2.0.md).

## 1. Register adapters with blank imports

Changesets only find chain families whose adapter packages have been imported. Import each package once, anywhere in the binary, for example in `chainlink-deployments/domains/<domain>/<env>/durable_pipelines.go` or in a test file:

```go
import (
    // EVM v1
    _ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters" // version-agnostic: MCMS reader, ownership, tokens, address normalizer
    _ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences" // lane + deployer adapter
    _ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/adapters"  // curse, fees, lane migrator
    // EVM v2
    _ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/adapters"

    // Solana v1
    _ "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_0_0/adapters"
    _ "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
    _ "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/adapters"
    // Solana v2 (separate repo)
    _ "github.com/smartcontractkit/chainlink-ccip-solana/deployment/v2_0_0/adapters"
)
```

To support older token pools as well, also import the EVM `v1_5_x` and `v1_6_x` adapter packages; each one registers a token adapter for its pool version.

If a changeset fails with *"no … adapter registered for chain family X"*, the matching import is missing, or the version you passed doesn't match any registered version. Test adapters (`…/testadapter`) are imported the same way, but only in tests.

## 2. Run a changeset

Every changeset is a `cldf.ChangeSetV2[Config]`. Call `VerifyPreconditions` to validate the config without writing anything, then call `Apply` to run it:

```go
cs := deploy.DeployContracts(deploy.GetRegistry())
if err := cs.VerifyPreconditions(env, cfg); err != nil { ... }
out, err := cs.Apply(env, cfg)
```

`out` (`cldf.ChangesetOutput`) contains:

- **`DataStore`**: new addresses and metadata. Merge them into your environment DataStore.
- **`MCMSTimelockProposals`**: the proposal, present when any write targeted a timelock-owned contract. It needs to be signed and executed.
- **`Reports`**: operation reports. Pass them back on retry, so completed operations are skipped.

In `chainlink-deployments`, each changeset is registered as a durable pipeline in `domains/<domain>/<env>/durable_pipelines.go` (CCIP in `domains/ccip`, the CCV/2.0 off-chain flows in `domains/ccv`), for example:

```go
registry.Add(
    DurablePipeline_deploy_usdc_cctpv2,
    cldf_changeset.Configure(ccip2changesets.DeployCCTPChains(cctpChainRegistry, cciputilschangeset.GetRegistry())).WithEnvInput(),
)
```

That file also builds the registries without singletons (CCTP, Lombard) and registers each family's adapter in them.

Changesets that take registries as arguments almost always take the global singletons (`…GetRegistry()`). Changesets with no arguments resolve registries internally.

### MCMS input

Configs that can produce a proposal embed `MCMS mcms.Input`. In 2.0 changesets built with `NewFromOnChainSequence`, it sits at the top level of a `changesets.WithMCMS[Cfg]{MCMS, Cfg}` wrapper. Usually you set only these:

```go
MCMS: mcms.Input{
    TimelockAction: mcms_types.TimelockActionSchedule, // or Bypass / Cancel
    Qualifier:      utils.CLLQualifier,                 // which MCMS/timelock (CLLCCIP, RMNMCMS, UltraFastCurse)
    Description:    "connect ethereum <> solana",
},
```

An empty `mcms.Input{}` is fine when everything is deployer-owned, for example in tests. See [Architecture → MCMS](architecture.md#mcms-proposals).

## 3. v1 (1.6) examples

**Deploy chain contracts** (from `chains/evm/deployment/v1_6_0/changesets/deploy_chain_contracts_test.go`):

```go
out, err := deploy.DeployContracts(deploy.GetRegistry()).Apply(*e, deploy.ContractDeploymentConfig{
    MCMS: mcms.Input{},
    Chains: map[uint64]deploy.ContractDeploymentConfigPerChain{
        chain_selectors.ETHEREUM_MAINNET.Selector: {
            Version:                                 semver.MustParse("1.6.0"),
            MaxFeeJuelsPerMsg:                       new(big.Int).Mul(big.NewInt(200), big.NewInt(1e18)),
            TokenPriceStalenessThreshold:            24 * 60 * 60,
            LinkPremiumMultiplier:                   9e17,
            NativeTokenPremiumMultiplier:            1e18,
            PermissionLessExecutionThresholdSeconds: uint32((20 * time.Minute).Seconds()),
            GasForCallExactCheck:                    5000,
        },
    },
})
```

**Connect two chains.** `ConnectChains` configures both directions in one call. Router, OnRamp, OffRamp, and FeeQuoter addresses are filled in from the DataStore, so you don't set them:

```go
cs := lanes.ConnectChains(
    lanes.GetLaneAdapterRegistry(),
    changesets.GetRegistry(),
    adapters.GetDeployChainContractsRegistry(), // optional (may be nil): blocks downgrading a lane already on 2.0
)
out, err := cs.Apply(env, lanes.ConnectChainsConfig{
    Lanes: []lanes.LaneConfig{{
        Version: semver.MustParse("1.6.0"),
        ChainA:  lanes.ChainDefinition{Selector: evmSel, GasPrice: big.NewInt(2e12)},
        ChainB:  lanes.ChainDefinition{Selector: solSel, GasPrice: big.NewInt(2e12)},
    }},
    MCMS: mcms.Input{TimelockAction: mcms_types.TimelockActionSchedule},
})
```

The FeeQuoter dest config defaults come from each chain's `LaneAdapter.GetFeeQuoterDestChainConfig()`. The default gas price comes from `GetDefaultGasPrice()`.

## 4. v2 (2.0) examples

2.0 changesets are driven by an **environment topology**: NOPs, committees and their per-chain membership, and executor pools. Load it from TOML with `offchain.LoadEnvironmentTopology(path)`. The changesets derive committee verifiers, signer quorums, executors, CCVs, and contract addresses from the topology plus the DataStore. Set overrides only where the defaults are wrong.

**Deploy chain contracts.** This works for any registered family:

```go
cs := v2changesets.DeployChainContracts(
    adapters.GetDeployChainContractsRegistry(),
    adapters.GetChainFamilyRegistry(),
)
out, err := cs.Apply(env, changesets.WithMCMS[v2changesets.DeployChainContractsCfg]{
    MCMS: mcms.Input{},
    Cfg: v2changesets.DeployChainContractsCfg{
        Topology:       topology,
        ChainSelectors: []uint64{evmSel, solSel},
        ChainOverrides: map[uint64]v2changesets.DeployChainContractsPerChainCfg{
            evmSel: {DeployTestRouter: true}, // optional; ContractParams overrides adapter defaults
        },
    },
})
```

By default the deployed contracts are handed to the `CLLCCIP` timelock, which must already exist. Set `DeployerKeyOwned: true` per chain to keep them deployer-owned.

**Configure lanes.** Each lane pair is bidirectional, just like 1.6 `ConnectChains` (from `chains/evm/deployment/v2_0_0/adapters/upgrade_onramp_test.go`):

```go
cs := v2changesets.ConfigureChainsForLanesFromTopology(
    adapters.GetCommitteeVerifierContractRegistry(),
    adapters.GetChainFamilyRegistry(),
    changesets.GetRegistry(),
)
out, err := cs.Apply(env, v2changesets.ConfigureChainsForLanesFromTopologyConfig{
    Topology: topology,
    BuildLanesCrossFamilyConfig: v2changesets.BuildLanesCrossFamilyConfig{
        Lanes: []v2changesets.CrossFamilyLanePair{
            {ChainA: evmSel, ChainB: solSel}, // ChainAOverrides / ChainBOverrides for per-side tweaks
        },
        MCMS: mcms.Input{},
    },
})
```

Each leg is verified by every committee that has **both** chains in its `chain_configs`. A pair that shares no committee is rejected.

**Tokens, fees, ownership, and curse** use the same shared changesets as 1.6. Pass a 2.0 pool version where a version is needed. For example, `tokens.TokenExpansion()` with `ChainAdapterVersion: semver.MustParse("2.0.0")` and `TokenPoolVersion: 2.0.0`.

## 5. v1 vs v2 per flow

| Flow | v1 (1.6) | v2 (2.0) |
|---|---|---|
| Deploy MCMS | `deploy.DeployMCMS` / `FinalizeDeployMCMS` | Same changesets. EVM registers its `Deployer` at 2.0.0 too. Solana 2.0 has no MCMS deploy and reuses existing MCMS |
| Deploy chain contracts | `deploy.DeployContracts` (per-chain config, `Version` selects the adapter) | `v2_0_0/changesets.DeployChainContracts` (topology + chain selectors, adapter supplies defaults) |
| Lanes | `lanes.ConnectChains`: list of chain pairs, run as source and dest legs | `ConfigureChainsForLanesFromTopology`: list of chain pairs, each chain configured once for all its remotes. CCVs, executors, and addresses are derived |
| Disable lane | `lanes.DisableLane` | Not yet a 2.0 changeset |
| OCR / verification | `deploy.SetOCR3Config` | Committees from topology. Signer keys are fetched from JD during lane configuration |
| Token deploy + configure | `tokens.TokenExpansion`, `ConfigureTokensForTransfers` | Same changesets with 2.0 pool version. CCVs per remote via `OutboundCCVs`/`InboundCCVs`, plus `…ToAddAboveThreshold` |
| Manual registration | `tokens.ManualRegistration` | Same changeset (EVM 2.0 adapter supports it) |
| Pool rate limits | `tokens.SetTokenPoolRateLimits` | Same changeset. 2.0 adds fast-finality buckets (`OutboundRateLimits`) |
| Token transfer fees | `fees.SetTokenTransferFee` (on FeeQuoter) | On the **pool**, via `tokens.ConfigureTokenPool` / `RemoteChainConfig.TokenTransferFeeConfig` (adapter implements `TokenFeeAdapter`) |
| FeeQuoter dest config | `fees.UpdateFeeQuoterDests` | Set during lane configuration (`ChainOverrides.RemoteChainCfg.FeeQuoterDestChainConfig`). `UpdateFeeQuoterDests` also works with a 2.0 FeeQuoter |
| Migrate a chain from v1 to v2 | — | `deploy.UpdateFeeQuoterChangeset` (deploys a 2.0 FeeQuoter from imported v1 lane config), `deploy.LaneMigrateToNewVersionChangeset` (repoints Router and ramps) |
| Replace a v2 OnRamp | — | `UpgradeOnrampPhase1/2/3`, `UpgradeOnrampPhase3Rollback`, `UpgradeOnrampCleanup` |
| USDC / Lombard | — | `DeployCCTPChains`, `DeployLombardChains`. Pool updates: `RemoveCCTPRemotePools`, `SetCCTPTokenPoolRateLimits`, `SetCCTPTokenPoolDynamicConfig` and the Lombard equivalents |
| MCMS proposals | `OutputBuilder` | Same. Batch ops from any version combine into one proposal |

## 6. Changeset catalog

Constructors and inputs. Follow the source link for the config type. Each config struct is documented in code.

### Shared (`deploy`, `lanes`, `tokens`, `fees`, `fastcurse`, `authorizedcallers`)

| Changeset | Constructor | What it does |
|---|---|---|
| Deploy contracts (1.6) | `deploy.DeployContracts(deployReg)` | Deploy CCIP contracts per chain ([contracts.go](../deploy/contracts.go)) |
| Deploy MCMS | `deploy.DeployMCMS(deployReg, mcmsReg)` / `FinalizeDeployMCMS` | Deploy MCMS + timelock, then the optional second phase ([mcms.go](../deploy/mcms.go)) |
| Update MCMS config | `deploy.UpdateMCMSConfig(deployReg, mcmsReg)` | Change signers and quorums |
| Grant timelock admin | `deploy.GrantAdminRoleToTimelock(deployReg, mcmsReg)` | Make one timelock admin of another |
| Set OCR3 config | `deploy.SetOCR3Config(deployReg, mcmsReg)` | Copy OCR3 config from CCIPHome to remote OffRamps ([set_ocr3_config.go](../deploy/set_ocr3_config.go)) |
| Transfer / accept ownership | `deploy.TransferOwnershipChangeset(toReg, mcmsReg)` / `AcceptOwnershipChangeset` | Propose and accept ownership via MCMS ([transfer_ownership.go](../deploy/transfer_ownership.go)) |
| Lane migration | `deploy.LaneMigrateToNewVersionChangeset(migReg, mcmsReg)` | Point Router and ramps at a new ramp version ([lanemigrator.go](../deploy/lanemigrator.go)) |
| FeeQuoter upgrade | `deploy.UpdateFeeQuoterChangeset()` | Deploy or update the FeeQuoter, rebuilding config from existing lanes ([feequoterupdater.go](../deploy/feequoterupdater.go)) |
| Connect chains (1.6) | `lanes.ConnectChains(laneReg, mcmsReg, versionResolvers)` | Bidirectional lanes ([connect_chains.go](../lanes/connect_chains.go)) |
| Disable lane | `lanes.DisableLane(laneReg, mcmsReg)` | Disable both directions ([disable_lane.go](../lanes/disable_lane.go)) |
| PingPong | `lanes.ConfigurePingPongForLanes(e, reg, version, sel, remotes)` | Helper function (not a changeset) |
| Token expansion | `tokens.TokenExpansion()` | Deploy token + pool, configure remotes, hand ownership to timelock ([token_expansion.go](../tokens/token_expansion.go)) |
| Configure tokens for transfers | `tokens.ConfigureTokensForTransfers(tokReg, mcmsReg)` | Register and configure existing pools. Can auto-migrate from an active legacy pool ([configure_tokens_for_transfers.go](../tokens/configure_tokens_for_transfers.go)) |
| Configure token pool | `tokens.ConfigureTokenPool()` | Pool-level settings: fees, finality, router/admins ([configure_token_pool.go](../tokens/configure_token_pool.go)) |
| Manual registration | `tokens.ManualRegistration()` | Register a customer-deployed token ([manual_registration.go](../tokens/manual_registration.go)) |
| Rate limits | `tokens.SetTokenPoolRateLimits()` | Set rate limits. Inbound is derived from the counterpart's outbound ([rate_limits.go](../tokens/rate_limits.go)) |
| Remove remote pools | `tokens.RemoveRemotePools()` | Remove remote pool entries, optionally unregistering ([remove_remote_pools.go](../tokens/remove_remote_pools.go)) |
| Token admin roles | `tokens.GrantTokenAdminRole()` / `RevokeTokenAdminRole()` | ([grant_admin_role.go](../tokens/grant_admin_role.go), [revoke_admin_role.go](../tokens/revoke_admin_role.go)) |
| Migrate lock-release liquidity | `tokens.MigrateLockReleasePoolLiquidity(tokReg, mcmsReg)` | Legacy lock-release pool → 2.0 lockbox ([migrate_lock_release_pool_liquidity.go](../tokens/migrate_lock_release_pool_liquidity.go)) |
| Token transfer fee | `fees.SetTokenTransferFee()` | FeeQuoter per-token fees: user value → on-chain value → adapter default ([set_token_transfer_fee.go](../fees/set_token_transfer_fee.go)) |
| FeeQuoter dests | `fees.UpdateFeeQuoterDests()` | Upsert dest chain configs with an `Override` func ([update_fq_dests.go](../fees/update_fq_dests.go)) |
| Fee aggregator | `fees.SetFeeAggregator()` / `WithdrawFeeTokens()` | ([set_fee_aggregator.go](../fees/set_fee_aggregator.go), [withdraw_fee_tokens.go](../fees/withdraw_fee_tokens.go)) |
| Curse | `fastcurse.CurseChangeset` / `UncurseChangeset` / `GloballyCurseChainChangeset` / `GloballyUncurseChainChangeset` (`curseReg, mcmsReg`) | RMN curses ([fastcurse.go](../fastcurse/fastcurse.go)) |
| Authorized callers | `authorizedcallers.ConfigureAuthorizedCallersChangeset(reg, mcmsReg)` | Add or remove authorized callers ([authorizedcallers.go](../authorizedcallers/authorizedcallers.go)) |

### v2 (`v2_0_0/changesets`)

| Changeset | Constructor | What it does |
|---|---|---|
| Deploy chain contracts | `DeployChainContracts(deployReg, chainFamilyReg)` | Topology-driven deploy for any family |
| Configure lanes | `ConfigureChainsForLanesFromTopology(cvReg, chainFamilyReg, mcmsReg)` | Bidirectional lane pairs, from topology |
| CCTP | `DeployCCTPChains(cctpReg, mcmsReg)` | Deploy and configure USDC/CCTP pools and verifiers |
| Lombard | `DeployLombardChains(lombardReg, mcmsReg)` | Deploy and configure Lombard |
| CCTP / Lombard pools | `RemoveCCTPRemotePools`, `SetCCTPTokenPoolRateLimits`, `SetCCTPTokenPoolDynamicConfig` (`cctpReg, mcmsReg`); `RemoveLombardRemotePools`, `SetLombardTokenPoolRateLimits`, `SetLombardTokenPoolDynamicConfig` (`lombardReg, mcmsReg`) | Remove remote pools, set rate limits, set router and admins on CCTP or Lombard pools, which the `tokens` changesets do not manage ([cctp_lombard_token_pools.go](../v2_0_0/changesets/cctp_lombard_token_pools.go)) |
| Test verifier | `DeployTestVerifierChains(tvReg, mcmsReg)` | TESTVTR token, pool, and verifier for test envs |
| OnRamp upgrade | `UpgradeOnrampPhase1(upgraderReg, chainFamilyReg, mcmsReg, toReg)`, `Phase2(upgraderReg, mcmsReg)`, `Phase3`/`Phase3Rollback`/`Cleanup(upgraderReg, chainFamilyReg, mcmsReg)` | Staged OnRamp replacement |
| OffRamp source OnRamps | `OffRampSetSourceOnRamps(chainFamilyReg, mcmsReg)` | Set the source-OnRamp whitelist on an OffRamp |

Family-specific changesets (EVM CREATE2 factory, lockbox funding, Solana program upgrades and lookup tables, …) live in each family's `v2_0_0/changesets` package.

## 7. Hooks

[`hooks/`](../hooks/) provides pre-, post-, and post-proposal hooks for durable pipelines. Each is backed by a family-registered provider:

| Hook | Constructor |
|---|---|
| Verify deployed contracts on explorers | `VerifyDeployedContractsPostHookForMultipleChainFamilies(dom, families)` |
| Require contracts already verified | `RequireVerifiedEnvContractsPreHookForMultipleChainFamilies(dom, sels, refs)` |
| Require timelock ownership | `RequireOwnedEnvContractsPreHookForMultipleChainFamilies(dom, families)` |
| CCIP send smoke test after proposal | `GlobalPostProposalCCIPSendHook(dom)` |
