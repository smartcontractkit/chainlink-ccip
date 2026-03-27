# 1.5.1 Token Pool Migration Changeset — Implementation Plan

## Context

Token pool migration converts a remote chain from Lock/Release to Burn/Mint on a `HybridWithExternalMinterTokenPool` hub. The hub pool supports both modes per-chain via "groups" (Group 0 = Lock/Release, Group 1 = Burn/Mint). Remote chains get new `BurnMintWithExternalMinterTokenPool` contracts — already deployed before this changeset runs.

This changeset handles the **cutover only**: swap remote pool references on the hub, migrate the hub pool's internal accounting, and update the Token Admin Registry on the remote chain.

One changeset invocation = one remote chain migration.

## Safety Model

Safety does **not** rely on cross-chain atomicity. It relies on the operational sequence:

1. **Pre-pause** — A preceding proposal sets inbound/outbound rate limits on all affected pools to zero-capacity configs (`isEnabled=true`, `capacity=0`, `rate=0`), halting token transfers.
2. **Drain** — Operator waits for all inflight messages to complete (~30-60 min).
3. **Execute this proposal** — Cutover: remote pool swap + group migration + TAR update.
4. **Post-cutover verification** — Operator verifies: hub `getRemotePools`, hub `getGroup`, remote TAR pool, locked token balances.
5. **Unpause** — Only after both chain batches succeed AND post-cutover verification passes, a separate proposal restores the prior rate limits.

FeeQuoter parity is a separate review item, not part of this migration write path.

## Design Decisions

### 1. Proposal-only execution for all writes

**Decision:** All write operations use `contract.NoCallersAllowed` (`write.go:187`). No direct deployer execution.

**Rationale:** Migrations must execute atomically within an MCMS proposal. If the deployer key happened to own a contract (e.g. in staging), `OnlyOwner` would allow direct execution — breaking batch atomicity. `NoCallersAllowed` always returns false, forcing all writes into the proposal regardless of deployer key state. This was flagged as a must-fix by staff review.

### 2. Read-before-write / delta-only transactions

**Decision:** The sequence reads current on-chain state before each write and only emits transactions for actual deltas.

**Rationale:** Makes the changeset resume-safe. If a previous run partially completed (e.g. hub batch executed but remote batch failed), re-running produces only the remaining transactions instead of rebuilding the full set (which would fail on-chain for already-applied operations).

### 3. Use generated bindings directly (no custom wrapper)

**Decision:** Use `*hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPool` directly as the contract type for all hub pool operations.

**Rationale:** The generated binding from `ccip-contract-examples` already exposes `Address()` (line 3694), `Owner()` (line 648), `GetGroup()` (line 318), `GetRemotePools()` (line 428), `AddRemotePool()` (line 726), `RemoveRemotePool()` (line 798), and `UpdateGroups()` (line 882). It satisfies the `ownableContract` interface (`Address() common.Address` + `Owner(*bind.CallOpts) (common.Address, error)`) at `write.go:143`. No custom wrapper needed.

### 4. RemoteChainSupply stays explicit, but gets one exact check

**Decision:** `RemoteChainSupply` remains explicit operator input, but VerifyPreconditions enforces an exact `remoteToken.totalSupply()` match whenever the hub's current group for the target chain differs from `TargetGroup`. On a full no-op path (`currentGroup == TargetGroup`), the changeset still validates that the value is non-nil and non-negative, but skips the exact parity check because `updateGroups()` will not consume it. When `TargetGroup == 1` (Burn/Mint), VerifyPreconditions also requires `RemoteChainSupply <= hub.getLockedTokens()` as a sanity bound.

**Rationale:** Keeping the value explicit matches the runbook and proposal review flow, while the exact `totalSupply()` equality catches fat-fingered inputs after pause+drain without heuristics. Skipping the parity check on no-op paths preserves resume safety, and the locked-token bound is a cheap secondary check for the burn path.

### 5. Generic across supported 1.5.1 pool migrations

**Decision:** Keep explicit addresses and `TargetGroup` input so one changeset can handle different remote chains and supported `1.5.1` pool types, but do not try to accept arbitrary contracts. VerifyPreconditions rejects any `oldPool`/`newPool` pair whose onchain identity does not match the migration's supported contract set.

**Rationale:** This preserves a single reusable changeset for the migrations we actually run, without weakening validation. The state machine stays generic over current-vs-target state, while contract identity remains deterministic.

**Note:** `TargetGroup` is configurable in the API, but the `HybridWithExternalMinterTokenPool` currently only supports values `0` (LockRelease) and `1` (BurnMint). VerifyPreconditions enforces this. "Generic" here means reusable across supported `1.5.1` migration inputs, not arbitrary pool families.

### 6. Single MCMS proposal spanning hub + remote chains

**Decision:** Bundle hub ops and remote TAR SetPool into one proposal.

**Rationale:** Traffic is already paused by a preceding proposal. A single proposal simplifies governance. MCMS executes per-chain batches independently — ordering is safe because we add the new remote pool before removing the old one (no routing gap).

### 7. Verify executor ownership and governance refs in VerifyPreconditions

**Decision:** Resolve both `GetTimelockRef()` and `GetMCMSRef()` for both chains. Require non-empty datastore refs for all four. Then verify ownership/administration against the resolved timelock addresses.

**Rationale:** If the timelock doesn't own the hub pool or isn't the TAR administrator, the proposal will fail on-chain. If governance refs are missing, the `OutputBuilder.Build()` will fail later with an opaque error. Fail-fast in VerifyPreconditions with explicit errors like `"missing timelock for hub chain %d with qualifier %s"` or `"missing MCMS for remote chain %d with qualifier %s"`.

**Note:** Missing datastore refs are acceptable for contract inputs (pools, TAR, token) — those use explicit addresses. Missing datastore refs are NOT acceptable for MCMS/timelock resolution — those are infrastructure contracts that must exist in the datastore.

### 8. Explicit addresses with deterministic pool identity checks

**Decision:** Config uses explicit addresses (mandatory). VerifyPreconditions reads `typeAndVersion()` and `getToken()` from `oldPool` and `newPool` to validate that they are the expected contracts for this migration. If datastore refs exist for any address, validate consistency. Persist new refs at end of Apply.

**Rationale:** The pools are externally deployed via `ccip-contract-examples` and may not be in the datastore. Onchain identity checks catch obvious wrong-address input without depending on optional datastore entries or event heuristics.

### 9. "Already complete" = no-op success, but still validates config

**Decision:** If all state matches the post-migration target (new pool present, old pool absent, group already at target, TAR already pointing to new pool), return success with an empty proposal. VerifyPreconditions still runs governance, ownership, address, and pool-identity checks on a no-op path. The exact `RemoteChainSupply` parity check only runs when `updateGroups()` would actually be emitted.

**Rationale:** Makes re-runs safe. An operator can re-invoke after a partial failure without fear of duplicate transactions. Empty batch ops → `OutputBuilder.Build()` returns empty output. Full validation on no-op prevents silent acceptance of a bad config that happens to hit an already-migrated state, while skipping the unused supply parity check avoids false failures when re-running after completion.

### 10. Address encoding: 32-byte left-padded

**Decision:** Encode EVM addresses as `common.LeftPadBytes(addr.Bytes(), 32)` for `AddRemotePool`/`RemoveRemotePool`.

**Rationale:** Confirmed at `chains/evm/deployment/v1_6_1/sequences/configure_token_pool_for_remote_chain.go:126` and `:155`. The pool stores addresses in this form. `RemoveRemotePool` must match the stored form.

---

## Unexpected State Matrix

The sequence reads on-chain state and maps it to actions:

### Hub: Remote Pool List for target chain

The sequence reads the **full** `getRemotePools(remoteChainSelector)` set. Every address in the set must be either `oldPool` or `newPool`. Any address outside `{old, new}` is an **error** — the migration is operating in an unexpected environment.

After filtering to only {old, new}:

| Old in set | New in set | Action |
|---|---|---|
| Yes | No | Add new + Remove old |
| Yes | Yes | Remove old only (partially migrated) |
| No | Yes | Noop (pool swap complete) |
| No | No | **Error**: neither pool registered for this chain |

Extra-pool states (all errors):

| Set contents | Error |
|---|---|
| `{old, thirdParty}` | `"unexpected pool %s in remote pool set for chain %d"` |
| `{new, thirdParty}` | same |
| `{old, new, thirdParty}` | same |
| `{thirdParty}` | same |

### Hub: Group for target chain

| Current group | Target group | Action |
|---|---|---|
| Different from target | Target | UpdateGroups |
| Target | Target | Noop (already migrated) |

### Remote: TAR Pool for token

| Current TAR pool | Action |
|---|---|
| Old pool address | SetPool to new |
| New pool address | Noop (already updated) |
| Zero address | **Error**: token not administered |
| Other address | **Error**: unexpected third-party pool |

### Combined: Full noop

If all three checks are noop → return empty proposal (success, already complete).

---

## Files to Create/Modify

### File 1: Hub pool operations

**Modify:** `chains/evm/deployment/v1_6_0/operations/hybrid_with_external_minter_token_pool/hybrid_with_external_minter_token_pool.go`

Add to existing file. All write operations use `NoCallersAllowed`. Use the generated binding type directly as the contract type.

**New read operation — GetGroup:**

```go
var GetGroup = contract.NewRead(contract.ReadParams[
	uint64,
	uint8,
	*hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPool,
]{
	Name:         "hybrid-with-external-minter-token-pool:get-group",
	Version:      Version,
	Description:  "Returns the group assignment for a remote chain selector",
	ContractType: ContractType,
	NewContract:  hybrid_with_external_minter_token_pool.NewHybridWithExternalMinterTokenPool,
	CallContract: func(
		c *hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPool,
		opts *bind.CallOpts,
		remoteChainSelector uint64,
	) (uint8, error) {
		return c.GetGroup(opts, remoteChainSelector)
	},
})
```

**New write operation — AddRemotePool (proposal-only):**

```go
type AddRemotePoolArgs struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
}

var AddRemotePool = contract.NewWrite(contract.WriteParams[
	AddRemotePoolArgs,
	*hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPool,
]{
	Name:         "hybrid-with-external-minter-token-pool:add-remote-pool",
	Version:      Version,
	Description:  "Calls addRemotePool on HybridWithExternalMinterTokenPool",
	ContractType: ContractType,
	ContractABI:  hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPoolABI,
	NewContract:  hybrid_with_external_minter_token_pool.NewHybridWithExternalMinterTokenPool,
	IsAllowedCaller: contract.NoCallersAllowed[
		*hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPool,
		AddRemotePoolArgs,
	],
	Validate: func(_ AddRemotePoolArgs) error { return nil },
	CallContract: func(
		c *hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPool,
		opts *bind.TransactOpts,
		args AddRemotePoolArgs,
	) (*types.Transaction, error) {
		return c.AddRemotePool(opts, args.RemoteChainSelector, args.RemotePoolAddress)
	},
})
```

**New write operation — RemoveRemotePool (proposal-only):**

Same pattern as AddRemotePool, but calls `c.RemoveRemotePool(opts, args.RemoteChainSelector, args.RemotePoolAddress)`.

**New write operation — UpdateGroups (proposal-only):**

```go
type UpdateGroupsArgs struct {
	GroupUpdates []hybrid_with_external_minter_token_pool.HybridTokenPoolAbstractGroupUpdate
}

var UpdateGroups = contract.NewWrite(contract.WriteParams[
	UpdateGroupsArgs,
	*hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPool,
]{
	Name:         "hybrid-with-external-minter-token-pool:update-groups",
	Version:      Version,
	Description:  "Calls updateGroups on HybridWithExternalMinterTokenPool",
	ContractType: ContractType,
	ContractABI:  hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPoolABI,
	NewContract:  hybrid_with_external_minter_token_pool.NewHybridWithExternalMinterTokenPool,
	IsAllowedCaller: contract.NoCallersAllowed[
		*hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPool,
		UpdateGroupsArgs,
	],
	Validate: func(_ UpdateGroupsArgs) error { return nil },
	CallContract: func(
		c *hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPool,
		opts *bind.TransactOpts,
		args UpdateGroupsArgs,
	) (*types.Transaction, error) {
		return c.UpdateGroups(opts, args.GroupUpdates)
	},
})
```

**New read operation — GetRemotePools:**

```go
var GetRemotePools = contract.NewRead(contract.ReadParams[
	uint64,
	[][]byte,
	*hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPool,
]{
	Name:         "hybrid-with-external-minter-token-pool:get-remote-pools",
	Version:      Version,
	Description:  "Returns registered remote pool addresses for a chain selector",
	ContractType: ContractType,
	NewContract:  hybrid_with_external_minter_token_pool.NewHybridWithExternalMinterTokenPool,
	CallContract: func(
		c *hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPool,
		opts *bind.CallOpts,
		remoteChainSelector uint64,
	) ([][]byte, error) {
		return c.GetRemotePools(opts, remoteChainSelector)
	},
})
```

**New imports needed:** `bind`, `types` from go-ethereum; `contract` from `chains/evm/deployment/utils/operations/contract`.

---

### File 2: Migration sequence

**Create:** `chains/evm/deployment/v1_6_0/sequences/migrate_token_pool.go`

**Package:** `sequences`

**Input type:**

```go
type MigrateTokenPoolInput struct {
	HubChainSelector     uint64
	HubPoolAddress       common.Address
	RemoteChainSelector  uint64
	NewRemotePoolAddress common.Address
	OldRemotePoolAddress common.Address
	RemoteChainSupply    *big.Int
	TargetGroup          uint8
	RemoteTARAddress     common.Address
	RemoteTokenAddress   common.Address
}
```

**Proposal-only TAR SetPool** (private to this file):

```go
var migrationSetPool = contract.NewWrite(contract.WriteParams[
	tar_ops.SetPoolArgs,
	*token_admin_registry.TokenAdminRegistry,
]{
	Name:            "migration:token-admin-registry:set-pool",
	Version:         semver.MustParse("1.5.0"),
	Description:     "Proposal-only setPool for migration",
	ContractType:    tar_ops.ContractType,
	ContractABI:     token_admin_registry.TokenAdminRegistryABI,
	NewContract:     token_admin_registry.NewTokenAdminRegistry,
	IsAllowedCaller: contract.NoCallersAllowed[
		*token_admin_registry.TokenAdminRegistry,
		tar_ops.SetPoolArgs,
	],
	Validate: func(_ tar_ops.SetPoolArgs) error { return nil },
	CallContract: func(
		c *token_admin_registry.TokenAdminRegistry,
		opts *bind.TransactOpts,
		args tar_ops.SetPoolArgs,
	) (*types.Transaction, error) {
		return c.SetPool(opts, args.TokenAddress, args.TokenPoolAddress)
	},
})
```

**Sequence logic (read-before-write):**

```go
var MigrateTokenPool = cldf_ops.NewSequence(
	"migrate-token-pool",
	semver.MustParse("1.6.0"),
	"Migrates a remote chain's pool configuration on a HybridWithExternalMinterTokenPool hub",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input MigrateTokenPoolInput) (sequences.OnChainOutput, error) {
		hubChain := chains.EVMChains()[input.HubChainSelector]
		remoteChain := chains.EVMChains()[input.RemoteChainSelector]

		oldPoolBytes := common.LeftPadBytes(input.OldRemotePoolAddress.Bytes(), 32)
		newPoolBytes := common.LeftPadBytes(input.NewRemotePoolAddress.Bytes(), 32)

		hubWrites := make([]evm_contract.WriteOutput, 0, 3)

		// --- Hub: Read remote pools ---
		remotePools, err := cldf_ops.ExecuteOperation(b, hybrid_pool_ops.GetRemotePools, hubChain,
			evm_contract.FunctionInput[uint64]{
				ChainSelector: input.HubChainSelector,
				Address:       input.HubPoolAddress,
				Args:          input.RemoteChainSelector,
			})
		// error handling...

		// Enforce full pool set — reject any address outside {old, new}
		for _, pool := range remotePools.Output {
			if !bytes.Equal(pool, oldPoolBytes) && !bytes.Equal(pool, newPoolBytes) {
				return ..., fmt.Errorf("unexpected pool %x in remote pool set for chain %d — expected only {%s, %s}",
					pool, input.RemoteChainSelector, input.OldRemotePoolAddress, input.NewRemotePoolAddress)
			}
		}

		oldPresent := containsBytes(remotePools.Output, oldPoolBytes)
		newPresent := containsBytes(remotePools.Output, newPoolBytes)

		if !oldPresent && !newPresent {
			return ..., fmt.Errorf("neither old pool %s nor new pool %s registered for chain %d",
				input.OldRemotePoolAddress, input.NewRemotePoolAddress, input.RemoteChainSelector)
		}

		// Hub: Add new remote pool if not present
		if !newPresent {
			addReport, err := cldf_ops.ExecuteOperation(b, hybrid_pool_ops.AddRemotePool, hubChain, ...)
			hubWrites = append(hubWrites, addReport.Output)
		}

		// Hub: Remove old remote pool if still present
		if oldPresent {
			removeReport, err := cldf_ops.ExecuteOperation(b, hybrid_pool_ops.RemoveRemotePool, hubChain, ...)
			hubWrites = append(hubWrites, removeReport.Output)
		}

		// --- Hub: Read current group ---
		currentGroup, err := cldf_ops.ExecuteOperation(b, hybrid_pool_ops.GetGroup, hubChain,
			evm_contract.FunctionInput[uint64]{
				ChainSelector: input.HubChainSelector,
				Address:       input.HubPoolAddress,
				Args:          input.RemoteChainSelector,
			})
		// error handling...

		// Hub: Update group if different from target
		if currentGroup.Output != input.TargetGroup {
			updateReport, err := cldf_ops.ExecuteOperation(b, hybrid_pool_ops.UpdateGroups, hubChain,
				evm_contract.FunctionInput[hybrid_pool_ops.UpdateGroupsArgs]{
					ChainSelector: input.HubChainSelector,
					Address:       input.HubPoolAddress,
					Args: hybrid_pool_ops.UpdateGroupsArgs{
						GroupUpdates: []hybrid_with_external_minter_token_pool.HybridTokenPoolAbstractGroupUpdate{
							{
								RemoteChainSelector: input.RemoteChainSelector,
								Group:               input.TargetGroup,
								RemoteChainSupply:   input.RemoteChainSupply,
							},
						},
					},
				})
			hubWrites = append(hubWrites, updateReport.Output)
		}

		// --- Remote: Read TAR config ---
		tarConfig, err := cldf_ops.ExecuteOperation(b, tar_ops.GetTokenConfig, remoteChain,
			evm_contract.FunctionInput[common.Address]{
				ChainSelector: input.RemoteChainSelector,
				Address:       input.RemoteTARAddress,
				Args:          input.RemoteTokenAddress,
			})
		// error handling...

		remoteWrites := make([]evm_contract.WriteOutput, 0, 1)

		currentPool := tarConfig.Output.TokenPool
		if currentPool == (common.Address{}) {
			return ..., fmt.Errorf("token %s has no pool set in TAR on chain %d", ...)
		}
		if currentPool == input.NewRemotePoolAddress {
			// Noop — already pointing to new pool
		} else if currentPool == input.OldRemotePoolAddress {
			setPoolReport, err := cldf_ops.ExecuteOperation(b, migrationSetPool, remoteChain,
				evm_contract.FunctionInput[tar_ops.SetPoolArgs]{
					ChainSelector: input.RemoteChainSelector,
					Address:       input.RemoteTARAddress,
					Args: tar_ops.SetPoolArgs{
						TokenAddress:     input.RemoteTokenAddress,
						TokenPoolAddress: input.NewRemotePoolAddress,
					},
				})
			remoteWrites = append(remoteWrites, setPoolReport.Output)
		} else {
			return ..., fmt.Errorf("TAR pool %s is neither old %s nor new %s — unexpected state",
				currentPool, input.OldRemotePoolAddress, input.NewRemotePoolAddress)
		}

		// --- Build batch operations ---
		var batchOps []mcms_types.BatchOperation

		if len(hubWrites) > 0 {
			hubBatch, err := evm_contract.NewBatchOperationFromWrites(hubWrites)
			batchOps = append(batchOps, hubBatch)
		}
		if len(remoteWrites) > 0 {
			remoteBatch, err := evm_contract.NewBatchOperationFromWrites(remoteWrites)
			batchOps = append(batchOps, remoteBatch)
		}

		return sequences.OnChainOutput{BatchOps: batchOps}, nil
	},
)
```

**Helper:**

```go
func containsBytes(haystack [][]byte, needle []byte) bool {
	return slices.ContainsFunc(haystack, func(b []byte) bool {
		return bytes.Equal(b, needle)
	})
}
```

**Imports:**

```
"bytes"
"slices"
"github.com/Masterminds/semver/v3"
"github.com/ethereum/go-ethereum/common"
mcms_types "github.com/smartcontractkit/mcms/types"
cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
hybrid_pool_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/hybrid_with_external_minter_token_pool"
tar_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
token_admin_registry "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/token_admin_registry"
hybrid_with_external_minter_token_pool "github.com/smartcontractkit/ccip-contract-examples/chains/evm/gobindings/generated/latest/hybrid_with_external_minter_token_pool"
"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
```

**Pattern reference:** `chains/evm/deployment/v1_6_1/sequences/configure_token_pool_for_remote_chain.go` — same read-then-write pattern.

---

### File 3: Changeset

**Create:** `chains/evm/deployment/v1_6_0/changesets/migrate_token_pool.go`

**Package:** `changesets`

**Config type:**

```go
type MigrateTokenPoolConfig struct {
	HubChainSelector     uint64     `json:"hubChainSelector"     yaml:"hubChainSelector"`
	HubPoolAddress       string     `json:"hubPoolAddress"       yaml:"hubPoolAddress"`
	RemoteChainSelector  uint64     `json:"remoteChainSelector"  yaml:"remoteChainSelector"`
	NewRemotePoolAddress string     `json:"newRemotePoolAddress" yaml:"newRemotePoolAddress"`
	OldRemotePoolAddress string     `json:"oldRemotePoolAddress" yaml:"oldRemotePoolAddress"`
	RemoteChainSupply    *big.Int   `json:"remoteChainSupply"    yaml:"remoteChainSupply"`
	TargetGroup          uint8      `json:"targetGroup"          yaml:"targetGroup"`
	RemoteTARAddress     string     `json:"remoteTARAddress"     yaml:"remoteTARAddress"`
	RemoteTokenAddress   string     `json:"remoteTokenAddress"   yaml:"remoteTokenAddress"`
	MCMS                 mcms.Input `json:"mcms,omitempty"       yaml:"mcms,omitempty"`
}
```

**Constructor:**

```go
func MigrateTokenPool(mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[MigrateTokenPoolConfig] {
	return cldf.CreateChangeSet(
		makeApplyMigrateTokenPool(mcmsRegistry),
		makeVerifyMigrateTokenPool(mcmsRegistry),
	)
}
```

**VerifyPreconditions** (`makeVerifyMigrateTokenPool(mcmsRegistry)`):

1. **Input validation:**
   - `cfg.MCMS.Validate()` returns nil
   - `cfg.HubChainSelector != cfg.RemoteChainSelector`
   - Both selectors pass `chain_selectors.GetSelectorFamily()` == `FamilyEVM`
   - Both chains exist in `env.BlockChains`
   - All 5 address fields pass `common.IsHexAddress()`
   - All 5 address fields are non-zero after `common.HexToAddress()` (the latter silently returns zero address for malformed input)
   - `cfg.OldRemotePoolAddress != cfg.NewRemotePoolAddress` — degenerate same-address migration is rejected
   - `cfg.RemoteChainSupply` is non-nil and `Sign() >= 0`
   - `cfg.TargetGroup` is 0 or 1

2. **Governance ref validation (fail-fast, explicit errors):**
   - `mcmsReader, ok := mcmsRegistry.GetMCMSReader(chain_selectors.FamilyEVM)` — if `!ok`: fail with `"no MCMS reader registered for EVM"`. Note: `GetMCMSReader` returns `(MCMSReader, bool)` not `(MCMSReader, error)`.
   - For **each** chain (hub, remote):
     - `timelockRef, err := mcmsReader.GetTimelockRef(e, chainSelector, cfg.MCMS)` — on error or empty `timelockRef.Address`: fail with `"missing timelock for chain %d with qualifier %s"`. Note: the EVM implementation (`mcmsreader.go:49-58`) always returns `nil` error; when the ref is not found it returns an empty `AddressRef`. Check `timelockRef.Address == ""` rather than relying on error.
     - `mcmsRef, err := mcmsReader.GetMCMSRef(e, chainSelector, cfg.MCMS)` — same pattern: check for empty `mcmsRef.Address`. Fail with `"missing MCMS for chain %d with qualifier %s"`.

3. **Pool identity verification:**
   - Read `typeAndVersion()` for `cfg.OldRemotePoolAddress` on the remote chain — require that it belongs to the supported `1.5.1` old-pool allowlist for this migration
   - Read `typeAndVersion()` for `cfg.NewRemotePoolAddress` on the remote chain — require that it matches the expected new pool type/version for the rollout
   - Read `getToken()` on both `oldPool` and `newPool` — require equality to `cfg.RemoteTokenAddress`

4. **Executor ownership verification:**
   - Read hub pool owner (bind and call `Owner()`) — verify `hubPoolOwner == common.HexToAddress(hubTimelockRef.Address)`; fail with `"hub pool %s owner %s does not match timelock %s on chain %d"`
   - Read remote TAR config via `ExecuteOperation(GetTokenConfig, remoteChain, {TARAddr, tokenAddr})` — verify `tarConfig.Administrator == common.HexToAddress(remoteTimelockRef.Address)`; fail with `"TAR administrator %s for token %s does not match timelock %s on chain %d"`

5. **RemoteChainSupply verification:**
   - Read current hub group via `ExecuteOperation(GetGroup, hubChain, ...)`
   - If `currentGroup != cfg.TargetGroup`, bind `cfg.RemoteTokenAddress` as ERC20 on the remote chain and require `totalSupply() == cfg.RemoteChainSupply`
   - If `currentGroup != cfg.TargetGroup && cfg.TargetGroup == 1`, read hub `getLockedTokens()` and require `cfg.RemoteChainSupply <= lockedTokens`
   - If `currentGroup == cfg.TargetGroup`, skip the exact supply parity check because `updateGroups()` will not consume the value on this run

6. **Datastore cross-check (opportunistic):**
   - If env.DataStore has refs matching any contract input address, verify they agree
   - If a ref exists but address differs from input, fail
   - Missing datastore refs for contract inputs are acceptable (pools are externally deployed)

**Apply** (`makeApplyMigrateTokenPool(mcmsRegistry)`):

```go
func makeApplyMigrateTokenPool(mcmsRegistry *changesets.MCMSReaderRegistry) ... {
	return func(e cldf.Environment, cfg MigrateTokenPoolConfig) (cldf.ChangesetOutput, error) {
		report, err := cldf_ops.ExecuteSequence(e.OperationsBundle,
			migrate_sequences.MigrateTokenPool,
			e.BlockChains,
			migrate_sequences.MigrateTokenPoolInput{
				HubChainSelector:     cfg.HubChainSelector,
				HubPoolAddress:       common.HexToAddress(cfg.HubPoolAddress),
				RemoteChainSelector:  cfg.RemoteChainSelector,
				NewRemotePoolAddress: common.HexToAddress(cfg.NewRemotePoolAddress),
				OldRemotePoolAddress: common.HexToAddress(cfg.OldRemotePoolAddress),
				RemoteChainSupply:    cfg.RemoteChainSupply,
				TargetGroup:          cfg.TargetGroup,
				RemoteTARAddress:     common.HexToAddress(cfg.RemoteTARAddress),
				RemoteTokenAddress:   common.HexToAddress(cfg.RemoteTokenAddress),
			})
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to migrate token pool: %w", err)
		}

		// Opportunistic datastore persistence
		ds := datastore.NewMemoryDataStore()
		// Persist new pool ref if not already in datastore
		// Persist hub pool ref if not already in datastore

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(report.ExecutionReports).
			WithBatchOps(report.Output.BatchOps).
			WithDataStore(ds).
			Build(cfg.MCMS)
	}
}
```

---

### File 4: Tests

**Create:** `chains/evm/deployment/v1_6_0/changesets/migrate_token_pool_test.go`

**VerifyPreconditions tests:**

| Case | Condition | Expected |
|------|-----------|----------|
| Valid config | All fields set, correct ownership | Pass |
| Invalid hub address | `"notanaddress"` | Error: invalid address |
| Invalid remote pool address | `"0x"` | Error: invalid address |
| Zero address after parse | `"0x0000000000000000000000000000000000000000"` | Error: zero address |
| Old == New pool address | Same address for both | Error: degenerate input |
| Nil supply | `RemoteChainSupply: nil` | Error |
| Negative supply | `RemoteChainSupply: -1` | Error |
| Same chain selector | hub == remote | Error |
| Missing hub chain | Hub selector not in env | Error |
| Missing remote chain | Remote selector not in env | Error |
| Invalid target group | `TargetGroup: 5` | Error |
| Missing hub timelock ref | `GetTimelockRef` returns empty ref (`Address == ""`) | Error: `"missing timelock for chain %d"` |
| Missing hub MCMS ref | `GetMCMSRef` returns empty ref (`Address == ""`) | Error: `"missing MCMS for chain %d"` |
| Missing remote timelock ref | `GetTimelockRef` returns empty ref (`Address == ""`) | Error: `"missing timelock for chain %d"` |
| Missing remote MCMS ref | `GetMCMSRef` returns empty ref (`Address == ""`) | Error: `"missing MCMS for chain %d"` |
| Unsupported old pool type | `oldPool.typeAndVersion()` not in allowlist | Error |
| Unexpected new pool type | `newPool.typeAndVersion()` wrong | Error |
| Old pool token mismatch | `oldPool.getToken() != RemoteTokenAddress` | Error |
| New pool token mismatch | `newPool.getToken() != RemoteTokenAddress` | Error |
| Hub pool not owned by timelock | Pool owner != timelock | Error: `"owner %s does not match timelock %s"` |
| TAR not administered by timelock | TAR admin != timelock | Error: `"administrator %s does not match timelock %s"` |
| Supply mismatch on real migration | `currentGroup != TargetGroup` and `RemoteChainSupply != remoteToken.totalSupply()` | Error |
| Locked-token bound exceeded | `currentGroup != TargetGroup`, `TargetGroup == 1`, `RemoteChainSupply > hub.getLockedTokens()` | Error |
| Datastore contradiction | Ref exists with different address | Error |
| Full no-op with stale supply | Group already at target, `RemoteChainSupply` no longer matches current `totalSupply()` | Pass |
| Full no-op still validates MCMS | All state migrated, bad MCMS input | Error (MCMS validation runs regardless) |

**Apply-level tests — proposal-only behavior:**

| Case | Assertion |
|------|-----------|
| Deployer is pool owner | All ops still go to proposal (no direct execution) |
| Deployer is TAR admin | SetPool still goes to proposal |
| Verify no `ExecInfo` on any write output | Confirms `NoCallersAllowed` in effect |

**Apply-level tests — state matrix:**

| Case | Hub Pool Set | Hub Group | Remote TAR | Expected Proposal |
|------|-------------|-----------|------------|-------------------|
| Fresh (normal) | `{old}` | != target | TAR→old | 3 hub txs + 1 remote tx |
| Partial: new added | `{old, new}` | != target | TAR→old | 1 hub remove + 1 hub updateGroups + 1 remote setPool |
| Partial: old removed | `{new}` | != target | TAR→old | 1 hub updateGroups + 1 remote setPool |
| Partial: hub complete | `{new}` | target | TAR→old | 1 remote setPool |
| Already complete | `{new}` | target | TAR→new | Empty proposal (no-op success) |
| Error: neither pool | `{}` | * | * | Error: `"neither pool registered"` |
| Error: third-party in set | `{old, thirdParty}` | * | * | Error: `"unexpected pool %x"` |
| Error: third-party + new | `{new, thirdParty}` | * | * | Error: `"unexpected pool %x"` |
| Error: all three | `{old, new, thirdParty}` | * | * | Error: `"unexpected pool %x"` |
| Error: third-party only | `{thirdParty}` | * | * | Error: `"unexpected pool %x"` |
| Error: third-party TAR | * | * | TAR→unknown | Error: `"unexpected state"` |
| Error: zero TAR pool | * | * | TAR→zero | Error: `"no pool set in TAR"` |

**Resume-safety tests:**

- Run changeset on fresh state → get proposal with 4 txs
- Simulate hub batch execution (apply pool swap + group update)
- Re-run changeset → get proposal with 1 tx (remote SetPool only)
- Simulate remote batch execution
- Re-run changeset → get empty proposal (no-op success)

**Integration tests:**

- Deploy HybridWithExternalMinterTokenPool on simulated hub chain
- Deploy BurnMintWithExternalMinterTokenPool on simulated remote chain (or mock)
- Set up TAR with old pool
- Execute changeset → verify proposal contents
- Execute proposal on simulated chains
- Verify: `getRemotePools()` returns only new pool, `getGroup()` returns target, TAR `getTokenConfig()` returns new pool, and hub `getLockedTokens()` changes by the expected amount for the chosen target group

---

## Execution Order Within MCMS Proposal

```
Hub chain batch (atomic per-chain):
  tx1: addRemotePool(remoteChainSelector, leftPad32(newPool))     [if needed]
  tx2: removeRemotePool(remoteChainSelector, leftPad32(oldPool))  [if needed]
  tx3: updateGroups([{remoteChainSelector, targetGroup, supply}]) [if needed]

Remote chain batch (atomic per-chain):
  tx4: setPool(tokenAddress, newPoolAddress)                      [if needed]
```

tx1 before tx2 ensures no routing gap. tx3 migrates accounting. tx4 is safe because traffic is already paused.

---

## Key References

| What | Where |
|------|-------|
| Generated bindings (Address, Owner, GetGroup, GetRemotePools, AddRemotePool, RemoveRemotePool, UpdateGroups) | `ccip-contract-examples/.../hybrid_with_external_minter_token_pool` |
| `contract.NoCallersAllowed` | `chains/evm/deployment/utils/operations/contract/write.go:187` |
| `ownableContract` interface | same file, line 143 |
| `contract.NewWrite` framework | same file, line 66 |
| TAR GetTokenConfig / SetPool ops | `chains/evm/deployment/v1_5_0/operations/token_admin_registry/token_admin_registry.go` |
| Existing HybridWithExternalMinterTokenPool Deploy op | `chains/evm/deployment/v1_6_0/operations/hybrid_with_external_minter_token_pool/hybrid_with_external_minter_token_pool.go` |
| `TypeAndVersion` helper | `chains/evm/deployment/utils/common.go` |
| Address encoding proof (LeftPadBytes 32) | `chains/evm/deployment/v1_6_1/sequences/configure_token_pool_for_remote_chain.go:126` |
| Read-before-write pattern | same file, lines 82-133 |
| `MCMSReaderRegistry.GetMCMSReader` | `deployment/utils/changesets/output.go:70` |
| EVM `MCMSReader` implementation (`GetTimelockRef`, `GetMCMSRef`) | `chains/evm/deployment/v1_0_0/adapters/mcmsreader.go` |
| Changeset pattern (CreateChangeSet) | `chains/evm/deployment/v1_6_0/changesets/token_governor.go` |
| OutputBuilder | `deployment/utils/changesets/output.go` |
| Hybrid pool `updateGroups` / `getLockedTokens` semantics | `ccip-contract-examples/.../contracts/stablecoin-governor/HybridTokenPoolAbstract.sol` |
| Rate limiter zero-capacity semantics | `chains/evm/contracts/libraries/RateLimiter.sol` |
| Existing migration changeset for reference | `deployment/tokens/migrate_lock_release_pool_liquidity.go` |

---

## Verification

1. From `chains/evm/deployment`: `go build ./...` — compiles
2. From `chains/evm/deployment`: `go test ./v1_6_0/changesets -run TestMigrateTokenPool -v`
3. From `chains/evm/deployment`: `go test ./v1_6_0/operations/... -run TestUpdateGroups -v`
4. Inspect generated MCMS proposal: correct tx count per state, 2 chain batches, correct calldata
5. Verify no `ExecInfo` populated on any write output (confirms proposal-only, no direct execution)
