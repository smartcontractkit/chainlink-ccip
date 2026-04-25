# Lane Configuration & Lifecycle Invariants

## 1. Lane Configuration (Source Side)

### 1.1 Destination Chain Config

- **INV-LCFG-1**: Each lane on the OnRamp is identified by `destChainSelector`. The OnRamp stores per-lane configuration including: address byte length, CCV sets, default executor, network fees, base execution gas cost, and the destination OffRamp address.
- **INV-LCFG-2**: `destChainSelector` must be non-zero and must not equal the local chain selector. A chain cannot send messages to itself.
- **INV-LCFG-3**: `addressBytesLength` must be non-zero. It defines the expected address size for the destination chain family.
- **INV-LCFG-4**: `baseExecutionGasCost` must be non-zero.
- **INV-LCFG-5**: A default executor must be configured (non-zero-value address). This ensures messages using legacy or defaulted extraArgs resolve to a concrete executor.
- **INV-LCFG-6**: The destination OffRamp address (raw bytes) must have a length equal to `addressBytesLength`.
- **INV-LCFG-7**: Lane configuration updates are restricted to the contract owner.

---

## 2. Lane Configuration (Destination Side)

### 2.1 Source Chain Config

- **INV-LCFG-8**: Each lane on the OffRamp is identified by `sourceChainSelector`. The OffRamp stores per-lane configuration including: enabled flag, allowed OnRamp addresses, and CCV sets.
- **INV-LCFG-9**: `sourceChainSelector` must be non-zero.
- **INV-LCFG-10**: Source chain configuration updates are restricted to the contract owner.

### 2.2 OnRamp Allowlist

- **INV-LCFG-11**: OnRamp addresses in the allowlist must be validated to not be empty/default value.

---

## 3. Lane Pause and Disable

### 3.1 Source-Side Pause

- **INV-PAUSE-1**: There must be a mechanism to pause a lane on the source side. On EVM, this is achieved by setting the router for that chain to the zero-value address. Pausing prevents new messages from being sent.
- **INV-PAUSE-2**: When a lane is paused, fee quoting also reverts. The lane is fully inactive for senders.

### 3.2 Destination-Side Disable

- **INV-PAUSE-3**: Setting `isEnabled = false` on the OffRamp's source chain config disables the lane. All execution attempts for messages from that source chain revert.

---

## 4. Inflight Message Handling

- **INV-INFLIGHT-1**: If the OffRamp's OnRamp allowlist is updated to remove the OnRamp that sent an inflight message, that message can no longer be executed.
- **INV-INFLIGHT-2**: If the OffRamp is disabled (`isEnabled = false`), inflight messages from that source chain cannot be executed until the lane is re-enabled.
- **INV-INFLIGHT-3**: If a remote pool is removed from a token pool's configuration, inflight messages referencing that source pool address fail token validation.

---

## 5. Router Configuration

Not all chain families use a router. On chains that do:

- **INV-RTR-1**: The router maps each destination chain selector to exactly one OnRamp. Only one OnRamp is active per destination at any time.
- **INV-RTR-2**: Multiple OffRamps may be registered per source chain on the router.
- **INV-RTR-3**: Setting the OnRamp to the zero-value address for a destination chain disables sending to that chain at the router level.
- **INV-RTR-4**: Only registered OffRamps may trigger message delivery to receivers via the router. Unregistered OffRamps are rejected.
- **INV-RTR-5**: Router configuration updates are restricted to the contract owner.

---

## 6. OnRamp Validation on Destination

- **INV-ONRVAL-1**: The OffRamp validates that `message.onRampAddress` matches an allowed OnRamp for the message's source chain. The comparison method is chain-specific. Messages from unrecognized OnRamps are rejected.
- **INV-ONRVAL-2**: The OffRamp validates that `message.offRampAddress` matches itself. Messages addressed to a different OffRamp are rejected.
- **INV-ONRVAL-3**: The OffRamp validates that `message.destChainSelector` matches the local chain selector.

---

## 7. Risk Management (Curse Mechanism)

- **INV-RMN-1**: A per-chain curse mechanism can block all operations for a specific chain. When a chain is cursed, the OnRamp, OffRamp, and token pools all reject operations involving that chain.
- **INV-RMN-2**: On the OnRamp, a curse on the destination chain blocks message sending.
- **INV-RMN-3**: On the OffRamp, a curse on the source chain blocks execution.
- **INV-RMN-4**: On token pools, a curse on the remote chain blocks both `lockOrBurn` (outbound) and `releaseOrMint` (inbound).
- **INV-RMN-5**: Curse checks are independent of lane configuration. A lane can be enabled but cursed, or disabled but not cursed.

---

## 8. Reentrancy Protection

- **INV-REENT-1**: The OnRamp's message sending function is protected by a reentrancy guard. External calls to CCVs, pools, and executors during message processing cannot re-enter the send path.
- **INV-REENT-2**: The OffRamp's execution function is protected by a reentrancy guard. External calls during message execution (CCV verification, `releaseOrMint`, `ccipReceive`) cannot re-enter the execution path.

---

## 9. Upgradability

CCVs, token pools, and the executor are external participants that the OnRamp and OffRamp interact with through stable interfaces. They can be upgraded with zero downtime and without users being aware of the change.

- **INV-UPG-1**: CCV implementations can be upgraded behind a stable address. Users, pools, lane configs, and receivers continue to reference the same CCV address. See CCV_INVARIANTS.md (section 4) for address stability requirements.
- **INV-UPG-2**: Token pool implementations can be upgraded. The remote chain's pool configuration supports multiple remote pool addresses per chain, allowing the old and new pool to coexist during the transition. See TOKEN_POOL_INVARIANTS.md INV-PCFG-2.
- **INV-UPG-3**: The executor can be upgraded by updating the lane configuration to point to a new executor address. Messages in flight continue to use the executor specified at send time (embedded in `ccvAndExecutorHash`); new messages use the updated executor.
- **INV-UPG-4**: None of these upgrades require user action, message replay, or protocol downtime. The transition is transparent to senders and receivers.

---

## 10. Access Control Summary

| Operation | Required Authority |
|-----------|--------------------|
| OnRamp lane configuration | Contract owner |
| OffRamp source chain configuration | Contract owner |
| Router ramp updates (where applicable) | Contract owner |
| Token pool remote chain configuration | Contract owner |
| Message sending (`ccipSend`) | Permissionless |
| Message execution (`execute`) | Permissionless |
| Curse / uncurse | Risk management system |

---

## 11. Multi-Step Transaction Safety

If a chain family's send or execute flow spans multiple transactions (non-atomic), the following invariants apply:

- **INV-ATOMIC-1**: If tokens are locked/burned before the message is finalized, a cancel/refund mechanism must exist to unlock/remint the tokens if finalization fails. A partially-completed send with locked tokens must never be irrecoverable.
- **INV-ATOMIC-2**: Configuration updates must not invalidate in-progress multi-step operations in a way that causes permanent token loss. If a config update can break an in-flight send or execute, either: (a) the config update must be gated until no in-flight operations exist, or (b) a recovery mechanism must exist for orphaned operations.

---

## 12. Replay Protection

- **INV-REPLAY-1**: On chains with upgradable contracts (in-place upgrades), replay protection state (execution state tracking) must survive contract upgrades. If replay protection is per-instance, deploying a new instance must not create a fresh replay state that allows re-execution of previously executed messages. For chains that use new instances on new offRamp addresses, the offRamp address in the message must be used to protect against re-execution.
