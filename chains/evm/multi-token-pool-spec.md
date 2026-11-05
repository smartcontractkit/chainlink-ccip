# Multi-token pool specification

This document defines the behavior of the multi-token pool and its configuration contracts. The code defines the
function signatures, errors, events, and implementation details. This specification covers burn and mint MTTPs.
Lock and release MTTPs are out of scope.

All MTTP lanes and workflows use CCIP v2 OnRamps and OffRamps. The FeeQuoter is used only with v2 ramps.

Legacy single-token pools remain supported. A lane can use an MTTP on one chain and a legacy pool on the other.
Pool versions and pool data formats are separate from ramp versions.

| Term | Meaning |
| --- | --- |
| MTTP | A pool that inherits `MultiTokenPool` |
| PCM | `PoolConfigManager` |
| TAR | `TokenAdminRegistry` |
| v2 ramps | CCIP 2.0 OnRamps and OffRamps |
| Legacy pool | A single-token `TokenPool`, including version 1.5.0 |

Main sections:

- [Contracts](#contracts), [records](#records), and [connections](#connections)
- [Transfer paths](#transfer-paths) and [pool data](#pool-data)
- [Remote pools](#remote-pool-records), [rate limits](#rate-limits), [fees](#fees), and
  [security](#security-records-and-hooks)
- [Views](#views), [roles](#roles), and [trust](#trust)
- [Workflow rules](#workflow-rules) and [configuration maintenance](#configuration-maintenance)
- [Token expansion](#token-expansion) and [chain expansion](#chain-expansion)
- [Token pool migration](#token-pool-migration) and
  [moving a migrated token to policies](#move-a-migrated-token-to-policies)
- [Response to a compromised remote pool](#incident-a-compromised-remote-pool)

## Contracts

Each chain has a pool, a PCM, and three stores. One MTTP serves many tokens. The TAR maps each token to its pool, so the
MTTP does not keep a token list. Two MTTPs can use the same PCM.

```text
OnRamp / OffRamp -- pool lookup --> TokenAdminRegistry
OnRamp / OffRamp -- transfer call --> BurnMintMultiTokenPool
                                                |
                                        IPoolConfigManager
                                                |
                                        PoolConfigManager
                                                |
                    +---------------------------+--------------------+
                    |                           |                    |
           ConnectionPolicyStore         FeePolicyStore     SecurityPolicyStore
```

### Contract responsibilities

| Contract | Responsibilities |
| --- | --- |
| MTTP | Holds mint and burn roles and token fees. Checks the ramp, RMN curse, and burn flavor. Calls hooks, moves tokens, and pays out fees. |
| PCM | Provides operator functions and views. Checks roles, callers against the TAR, and lane records. Forwards configuration writes to the stores. |
| Connection store | Stores token profiles, connections, subset tokens, trusted remote pools, and pause flags. Builds and validates pool data, checks source tokens and additional pool data, and converts inbound amounts. |
| Fee store | Stores fee and destination overhead records. |
| Security store | Stores rate limit and security records, including live rate limit buckets. |

- The PCM is the only contract the issuer calls for daily pool operation. It stores no token or policy configuration.
- Each store accepts writes only from authorized callers, which are PCMs.
- Each configuration type belongs to one store. Stores do not call each other.
- Each store keeps its own registered-token list. The PCM registers or deregisters a token in all three stores in one
  transaction.
- Calls that use multiple stores follow this order: connection, fee, then security.
- The MTTP calls the PCM only through `IPoolConfigManager`. Later PCM versions must preserve this interface's behavior.
  The MTTP also performs its own checks; a PCM cannot bypass them.
- The PCM calls a pool only to withdraw fees, through `IMultiTokenPool`. The PCM and stores do not depend on the pool
  implementation.

### Replacing a pool, manager, or store

The pool owner configures only its Router and PCM. A new PCM applies to the next transfer.

| Change | Effect on stored state |
| --- | --- |
| New pool, same PCM and stores | Configuration and live buckets remain in the stores. No copy is needed. |
| New PCM, same stores | Configuration and live buckets remain in the stores. Roles belong to the new PCM and do not carry over. |
| New store | The store starts empty. Its replacement requires a new PCM because the PCM's store addresses are immutable. |

No contract call copies records into a new store. `TAR.setPool` checks only that the PCM has the token profile.
It can therefore accept a pool whose stores lack required records. Transfers then revert.

The connection store builds and decodes pool data and converts inbound amounts. Replacing it can change these
behaviors without replacing the pool. A new pool data format is usable only after every receiving remote pool
supports it. The [contract replacement workflow](#replacing-contracts) covers the required preparation.

## Records

A token has six configuration types for each remote chain:

| Type | Store | Record contents |
| --- | --- | --- |
| Connection | Connection | An inline record contains the remote token. A policy record contains `onlySubsetTokens` and `subsetTokens`. |
| Remote pool | Connection | Trusted remote pools. Each inline record has one `useLegacyPoolDataOnSend` flag and one `trustPolicyRemotePools` flag. Each pool listed in an inline record has its own `allowLegacyPoolDataOnReceive` flag. Policy records have no flags. |
| Rate limit | Security | A configuration for each direction and finality class |
| Fee | Fee | USD fees and token fees in basis points |
| Destination overhead | Fee | `destGasOverhead` and `destBytesOverhead` |
| Security | Security | Outbound CCVs, inbound CCVs, and a hook |

A policy is a named set of records for one configuration type. Its name is a non-zero `bytes32`, unique within that
type. A token can have at most one policy per type and any number of inline records.

### Resolution order

For each token, configuration type, and remote chain, the store selects the first available record:

| Priority | Record | Applies to |
| --- | --- | --- |
| 1 | Inline | One token on one chain |
| 2 | Policy chain | Tokens assigned to the policy, on the specified chain |
| 3 | Policy default | Tokens assigned to the policy, on chains without a policy chain record |
| 4 | None | No record is available |

Policy defaults use chain selector zero. Connection and remote pool policies do not support defaults. Setting an
inline record with chain selector zero reverts. Resolving chain selector zero always returns no record.

A token uses all values from the selected record. There is one exception: an inline remote pool record can also trust
the policy's pools when `trustPolicyRemotePools` is enabled. See
[combining inline and policy pools](#combining-inline-and-policy-pools).

A token has a lane to a chain if a connection record applies to that token on that chain.

### Empty values

| Value | Meaning |
| --- | --- |
| Disabled rate limit configuration | No rate limit |
| Zero fees | No pool fee |
| Empty CCV list | The ramps' default CCVs |
| Zero hook address | No hook |
| Inline remote pool record that trusts no pools | Outbound transfers only, as in `TokenPool` |

A remote pool policy record must contain at least one pool.

### Records required for a transfer

| Direction | Connection | Remote pool | Rate limit | Fee | Destination overhead | Security |
| --- | --- | --- | --- | --- | --- | --- |
| Outbound | Yes | Yes | Yes | Yes | Yes | Yes |
| Inbound | Yes | Yes | Yes | No | No | Yes |

A transfer reverts if any required record is missing.

### Updating and removing records

- Setting a record creates it or replaces all its values. Setting a connection policy record preserves its subset
  tokens. An individual rate limit configuration update changes only one direction and finality class.
- Removing a record deletes all its values. Recreating it does not restore old values.
- Changes apply to the next transfer, including messages already in flight. Changing a remote token, provider
  response, trusted pool list, or additional pool data provider can cause these messages to fail.
- Removing a record immediately selects the next record in the resolution order. If none exists, the token has no
  record of that type. [Inline connection removal](#removing-an-inline-connection) also checks the remote token.
- Changing a token's policy immediately selects the new policy's records. Its inline records remain in place.
- Records have no flag to disable the whole record. Pauses and rate limit stops prevent transfers without removing
  records.

## Connections

### Policy and inline connections

Each connection policy record opens a chain for the policy's tokens. The policy has at most one remote token address
provider.

- The provider returns the remote token for each token and chain. A policy without a provider resolves to empty bytes.
  There is no default provider, because it could select a remote token that the issuer did not intend.
- Connection policies have no default record. A default would open a lane to every chain served by the Router.
- A chain record applies to every token assigned to the policy, unless `onlySubsetTokens` is enabled. With that flag
  enabled, only tokens in `subsetTokens` have a policy lane to the chain.
- An inline connection specifies the remote token directly. It overrides the policy, and the store does not call the
  provider.

Remote tokens must use the encoding that the remote MTTP uses for source tokens. On EVM chains, this is
`abi.encode(address)`. As in `TokenPool`, the store does not validate the address encoding of inline records or provider
responses.

The remote token must contain at least one non-zero byte. An empty or all-zero inline value reverts on write. The store
also checks the resolved remote token on every send and on receives that include a source token.

No chain has a token at an all-zero address. The v2 OnRamp accepts that address, so an unchecked zero value could burn
funds in a message that can never execute. Encoding an unset script value can produce this address.

### Removing an inline connection

Removing an inline connection must not silently change the remote token. If the connection policy would serve the lane
with different remote token bytes, the removal reverts.

This allows an issuer to remove a redundant inline record or close a lane. If the policy still serves the lane, its
remote token must match. Sending to a wrong remote token can permanently lose funds.

The removal check has these rules:

- It reads the token's connection policy, records, and provider as they exist when the removal runs.
- It runs before inline records are set in the same call. Removing and setting the same record can therefore revert.
  Setting an existing record replaces it without a removal.
- A provider failure does not block removal, because sends on that lane fail before the burn. A response that cannot
  decode as `bytes` does make the removal revert.
- The check is skipped if the policy has no record for the chain. It is also skipped if the record requires subset
  membership and the token is not a member. In either case, removal closes the lane.

The [connection update workflow](#changing-connections) covers replacements and adoption of a different policy
remote token.

### Closing a lane

Connections apply in both directions. Closing a lane also rejects inbound messages, including retries of old messages.
These changes can close a lane:

- Removing an inline connection without a policy connection to replace it.
- Removing a connection policy record.
- Enabling `onlySubsetTokens` when the token is not in the subset.
- Removing a token from the subset.
- Changing or clearing the token's connection policy.

An inline connection overrides the policy, so removing subset membership does not close a lane that still has an
inline connection. A record with `onlySubsetTokens` disabled continues to serve every assigned token.

The inline-removal check compares remote token addresses; it does not check for pending messages. Tooling must reject
a removal that closes a lane while messages can still arrive. The [lane closure workflow](#closing-a-token-lane)
stops sends and waits for those messages before removing connections.

### Subset tokens

- Adding a subset token requires an existing connection policy record with `onlySubsetTokens` enabled. The token must
  be assigned to that policy. Adding an existing subset member reverts.
- A token cannot leave or change its connection policy while it belongs to any of that policy's subsets. The call
  reverts with `TokenHasConfig`.
- The [roles table](#roles) lists who can remove subset entries or change policy assignments.
  The [connection update workflow](#changing-connections) gives the required order.
- Deregistration requires no connection policy assignment, so subset entries cannot outlive registration. A token
  that later rejoins the policy has the same lanes as a token joining for the first time.

### Remote token address providers

A connection policy has at most one provider. Provider updates use a separate function and require an existing
policy. Zero removes the provider. A non-zero address must contain code, or the call reverts with
`ProviderHasNoCode(provider)`. Changing or removing a connection record does not change the provider.

| Path | Provider call |
| --- | --- |
| Send through a policy connection | The store resolves the remote token through the provider for every transfer, including legacy sends. Failure stops the send before the burn. |
| Receive through a policy connection with a source token | The store calls the provider for the expected source token. Failure makes the message fail. |
| Receive without a source token | The store does not call the provider. |

Provider calls are views and cannot change state. A provider fails if it reverts, has no code, returns empty or all-zero
bytes, or returns data that cannot decode as `bytes`.

One provider serves all tokens and chains of its policy. A faulty provider can therefore affect all those lanes.

For an unknown token or chain, a provider must revert or return empty bytes. It must not default to
`abi.encode(address(0))`. Sends and receives that use an empty or all-zero response revert with `InvalidRemoteToken`.

### Same-address provider

`SameAddressRemoteTokenAddressProvider` returns `abi.encode` of the local token address for every chain. It has no
configuration, so one instance can serve all connection policies on a chain.

This provider is valid only when every policy token has the same address on every policy chain, including its
original deployment chain.

The provider also returns this EVM encoding for Solana, Aptos, and Sui. These chains use different address formats, so a
send to that value can lose funds. Non-EVM chains require inline connections with their raw token addresses.
Tooling must reject non-EVM chains before a connection policy record is set, as described in
[remote token checks](#checking-remote-tokens).

### Remote token checks in tooling

A policy connection resolves an address for each token. The correct asset must exist at that address on the remote
chain. Local contracts cannot check remote deployments, so tooling must verify these conditions before changes
that expose a token to a policy connection or alter its provider's response.

| Requirement | Reason |
| --- | --- |
| The remote token is the same asset. | If local token T maps to another asset with a return lane, sending T can mint that asset. Sending that asset back can mint T. |
| The remote chain maps the remote token back to the local token. | Both ends must agree on the token pair. |
| The remote TAR maps the remote token to an issuer pool. | The message must execute on the intended pool. |
| The address uses the remote chain's encoding. | Non-EVM chains require inline connections. |

A wrong remote token can permanently lose funds. The source burns the amount, and a retry uses the same message.
If no destination configuration accepts that address, the message cannot succeed.

The [remote token checks](#checking-remote-tokens) list the affected writes and the views used to verify them.

## Lane checks

The PCM checks for every other record type after either of these writes opens a lane for one token:

- Setting an inline connection.
- Adding a subset token.

The first missing record causes the write to revert. This is the only configuration write check that reads all three
stores.

The PCM does not run this check for:

- Connection policy records that apply to all policy tokens. No store lists those tokens.
- Token policy assignments, including connection policies.
- Changes or removals of other policy assignments.
- Any removal.

These writes can leave a connection without another required record. Tooling must check them before the lane is
used; the [workflow checks](#checks-before-opening-a-lane) cover this gap. Policy defaults can provide records for all
chains for types that support them. Connection and remote pool policies still require explicit chain records.

## Tokens and policies

### Creating and removing policies

- Setting the first record under a new name creates the policy. One call can create the policy and all its records.
- A policy can exist without records. It then provides no record to its tokens.
- A token can be assigned only to an existing policy.
- Removing all records does not remove the policy. Each policy update function, such as
  `applyRemotePoolPolicyConfigUpdates`, accepts a separate list of policy names to remove.
- A policy cannot be removed while it has records or assigned tokens.
- Removal deletes policy metadata, including rate limit decimals or the remote token address provider. Reusing the
  name creates a policy without the old values.

### Rate limit policy decimals

A rate limit policy has fixed decimals. All assigned tokens must have those decimals, so each configured amount has the
same token unit. Equal decimals do not imply equal economic value. These policies suit tokens with similar values.

Every record set includes the decimals. The first record fixes them for the policy's lifetime. Later record sets must
use the same value. Updating one configuration in an existing record does not include decimals. Other policy types
have no decimals.

### Token profiles

Setting a profile for a new token registers it without any policy assignments. Registration opens no lanes. Policy
assignments and inline records require a registered token. The zero address cannot be registered.

A profile contains:

- Local decimals.
- Burn flavor.
- Burn address.
- Allowed FTF finality.
- Additional pool data provider.

On registration, the connection store checks the supplied decimals against the token's `decimals()`, as the
`TokenPool` constructor does. Decimals cannot change while the token is registered; incorrect decimals would mint an
incorrect amount.

Updating a registered token replaces its profile and preserves its policies. It runs the registration checks, except
for the external `decimals()` call. The supplied decimals must still match the registered value.

Deregistration requires no policy assignments or inline records of any type. This prevents an incorrect address from
removing live configuration.

A deregistered token retains its stored profile, pause flags, and buckets. Views treat it as unregistered, except for
pause views and the stores' profile views. Registering it again can use different decimals.

### Assignments and batch order

Assigning a policy moves the token from its old policy of that type to the new one. A zero policy name clears the
assignment.

All operator functions accept batches, except global pause functions. Within a call:

| Function | Order |
| --- | --- |
| Functions with removals and sets | Removals, then sets |
| Policy record updates | Record removals, policy removals, then record sets |
| Token policy assignments, provider updates, and individual rate limit configuration updates | Input list order |

Policy writes do not update each assigned token. The exception is removing a connection policy record that has subset
tokens.

Removing a missing policy, record, or subset entry reverts. Deregistering an unregistered token also reverts. One failed
item reverts the whole call, so tooling must build batches from the state at execution time.

## Transfer paths

### Outbound

1. The user calls `Router.ccipSend`. The Router transfers the tokens to the MTTP and calls `OnRamp.forwardFromRouter`.
2. The OnRamp looks up the pool with `TAR.getPool(token)` and calls:
   - `MTTP.getRequiredCCVs`, through the PCM to the connection and security stores.
   - `MTTP.getFee`, through the PCM to the connection and fee stores.
   - `MTTP.lockOrBurn`.
3. The MTTP checks that the remote chain is not cursed. It then checks that the caller is the Router's OnRamp for
   that chain.
4. The MTTP calls `PCM.validateLockOrBurn`. The PCM and stores perform these checks in order:
   - **PCM:** the caller is `TAR.getPool(token)`.
   - **Connection store:** no global, outbound token, or outbound chain pause applies. Connection and remote pool
     records exist. The token profile allows the requested finality.
   - **Connection store:** resolve the remote token from the inline connection or the policy's provider.
   - **Connection store:** build the pool data. Long pool data includes the source token and the provider's additional
     pool data, or empty additional data without a provider. Return the remote token, pool data, decimals, flavor,
     and burn address.
   - **Fee store:** resolve fee and destination overhead records. Check the pool data against `destBytesOverhead` and
     return the token fee amount.
   - **Security store:** resolve the security record and return its hook.
   - **Security store:** resolve the rate limit record and consume the outbound amount after the token fee.
5. The MTTP checks that it supports the token's burn flavor, then calls the hook's preflight check if a hook
   is configured.
6. The MTTP burns the amount after the token fee, using the configured flavor. It returns the remote token and pool
   data.
7. The OnRamp checks that the pool data length does not exceed `destBytesOverhead`.

### Inbound

1. The executor calls `OffRamp.execute`. The OffRamp looks up the pool with `TAR.getPool(token)`.
2. The OffRamp calls `MTTP.getRequiredCCVs`, through the PCM to the connection and security stores, and verifies the
   required CCVs. It then calls `MTTP.releaseOrMint`.
3. The MTTP checks that the remote chain is not cursed. It then checks that the Router lists the caller as an OffRamp
   for that chain.
4. The MTTP calls `PCM.validateReleaseOrMint`. The PCM and stores perform these checks in order:
   - **PCM:** the caller is `TAR.getPool(token)`.
   - **Connection store:** no global, inbound token, or inbound chain pause applies. A connection record exists, and
     the remote pool record trusts the source pool. With `trustPolicyRemotePools` enabled, this includes policy pools.
   - **Connection store:** decode the decimals and any source token and additional pool data.
   - **Connection store:** if a source token is present, compare it with the expected remote token. For a policy
     connection, resolve the expected token through the provider. If the source token is absent, require
     `allowLegacyPoolDataOnReceive` for the source pool.
   - **Connection store:** if the token has an additional pool data provider, require it to accept the data. This check
     also passes empty data to the provider.
   - **Connection store:** convert the amount to local decimals. Return the local amount, flavor, and burn
     address.
   - **Security store:** resolve the rate limit record and consume the inbound amount in local units.
   - **Security store:** resolve the security record and return its hook.
5. The MTTP checks that it supports the token's burn flavor, then calls the hook's postflight check if a hook
   is configured.
6. The MTTP mints the tokens.

`TokenPool` checks the token before the curse and ramp. The MTTP checks the curse and ramp first. Lane resolution then
rejects a token without a lane.

## Pool data

The OffRamp does not pass the source token to the pool. An MTTP serves multiple tokens, so its address alone cannot
identify the burned token. The source MTTP includes that token in the pool data, and the destination checks it.

The connection store builds pool data on send and decodes it on receive. The PCM and pool pass it through unchanged.

### Formats

| Format | Length | Content |
| --- | --- | --- |
| Legacy, without decimals | 0 bytes | Empty data |
| Legacy, with decimals | 32 bytes | `abi.encode(uint256 decimals)` |
| Long | At least 192 bytes for an EVM source | `abi.encode(uint256 decimals, bytes sourceToken, bytes additionalPoolData)` |

Legacy pool data contains neither a source token nor additional pool data.

| Source | Format sent |
| --- | --- |
| Legacy pool older than 1.5.1 | Empty data |
| Legacy pool version 1.5.1 or later | 32 bytes |
| MTTP with `useLegacyPoolDataOnSend` enabled | 32 bytes |
| MTTP with `useLegacyPoolDataOnSend` disabled | Long pool data |

For an EVM source, `sourceToken` is `abi.encode(address)`. With empty additional pool data, the encoding has six words:

```text
word 0: decimals
word 1: offset of sourceToken, 0x60
word 2: offset of additionalPoolData, 0xa0
word 3: length of sourceToken, 0x20
word 4: abi.encode(localToken)
word 5: length of additionalPoolData, 0
```

Additional pool data follows word 5, padded to whole words. For example, 32 bytes of additional data produces 224 bytes
of pool data.

Source tokens use `bytes` because CCIP supports variable-length token addresses. A future non-EVM multi-token pool can
use its own address encoding.

### Decoding and decimals

- The MTTP accepts empty data, 32-byte data, or the exact long-format ABI encoding shown above. Long pool data must
  contain a non-empty source token. Other encodings revert.
- Decimals above 255 revert, as in `TokenPool`.
- When a source token is present, the destination always checks it, regardless of the legacy flags.
- The `IPoolV1` path follows the same decoding rules.
- Decimal conversion rounds down, as in `TokenPool`. Fractions below one local base unit are not minted.
  Overflow reverts.

Pool data normally carries the source decimals, so the two chains can use different decimals. These exceptions require
checks before a lane opens:

| Pool | Decimal behavior | Tooling check |
| --- | --- | --- |
| Version 1.5.0 | Sends no decimals and ignores received decimals. A destination receiving empty data uses its own decimals. | Both chains must use the same decimals. |
| Solana | Ignores pool data and uses the decimals in its chain configuration. | The Solana configuration must contain this chain's decimals. |

### Legacy compatibility

| Destination | Accepted pool data |
| --- | --- |
| MTTP | Long data; legacy data only with `allowLegacyPoolDataOnReceive` enabled for the source pool |
| Legacy pool version 1.5.1 or later | Legacy data only |
| Legacy pool version 1.5.0 | Ignores pool data |
| Solana pool | Ignores pool data |
| Aptos or Sui pool | Rejects long pool data |

A legacy pool is any single-token `TokenPool`, including version 1.5.0. One side can migrate to an MTTP while the
other remains legacy. That lane requires an inline remote pool record on the MTTP with:

- `useLegacyPoolDataOnSend` enabled for sends to the legacy pool.
- `allowLegacyPoolDataOnReceive` enabled for receipts from that pool.

These flags are independent. After the remote token moves to an MTTP, the local MTTP can send long data while still
accepting legacy data from the old pool. [Migration phase 3](#migration-phase-3-finish-each-lane) defines when to make
that change.

A non-MTTP pool never parses long pool data. Versions 1.5.1 and later, Aptos, and Sui reject it on every retry. Such a
message needs a destination MTTP to execute. Aptos and Sui have no MTTP, so the funds can remain stranded permanently.
Version 1.5.0 and Solana ignore the data, subject to the decimal rules above.

### Lanes to non-MTTP pools

A lane to a non-MTTP pool requires:

- An inline remote pool record with `useLegacyPoolDataOnSend` enabled at least until the remote TAR names an MTTP.
- `destBytesOverhead` equal to 32. This makes the pool data size check reject accidental long-format sends before
  the burn.

A remote pool policy always sends long data, so it cannot serve that token's lane. A remote pool policy must have no
record for a chain without an MTTP.

Tooling must check the remote TAR before a write enables long-format sends.
The [compatibility checks](#checking-remote-pool-compatibility) define which writes and lanes need this check,
including the exception for a locally paused token with no remote pool.
The [migration workflow](#token-pool-migration) defines when each side changes format.

### Pool data size check

On every send, the fee store compares the pool data length with the lane's `destBytesOverhead`. The MTTP reverts with
`PoolDataTooLarge` before the burn if the data is too long.

A destination overhead record therefore caps the pool data for its lane. Legacy sends use 32 bytes, which every valid
overhead record permits.

### Source token validation

| Configuration | Guarantee |
| --- | --- |
| Legacy receive disabled for every trusted pool | Every accepted message includes a source token, which the destination checks. An incorrect remote token in the source configuration fails at the destination without minting. |
| Legacy receive enabled for a pool | Messages without a source token do not prove which token was burned. The destination trusts the pool to send the correct asset. |
| Policy connection | The check proves only that the source token matches the policy provider's expected remote token. |

A legacy pool serves one token. Its address identifies that token only when it is trusted on that token's own inline
record. This is why only inline remote pool records support legacy receive permissions.

The v2 OnRamp includes the source token in the message, and the CCVs sign it. The OffRamp does not pass that field to
the pool. The MTTP can therefore check source identity only through long pool data.

Legacy data from a remote MTTP provides no token identity. An incorrect remote-token mapping on the source can then
mint the wrong destination token. A legacy pool that trusts a remote MTTP has the same exposure.

Legacy receive permission for an MTTP is restricted to migration or rollback: while it sends legacy data and
until all such messages have executed. A legacy pool may trust an MTTP only when its own token migrates to that MTTP.
The [compatibility checks](#checking-remote-pool-compatibility) also require reciprocal token mappings before legacy
send is enabled.

## Additional pool data

Additional pool data is the third field of the long format. The token issuer defines its meaning. For example, a UID
can identify the same asset across chains.

The token's additional pool data provider builds this data on send and validates it on receive. The connection store
transports the data without interpreting it.

### Provider configuration and calls

A token profile has at most one provider, used for every remote chain. Registration and later profile updates can
set it. Zero means no provider. A non-zero address must contain code, or the update reverts with
`ProviderHasNoCode(provider)`.

| Path | Behavior |
| --- | --- |
| Send long data with a provider | The store calls the provider with the token, requested finality, and token args. The pool data includes its response. |
| Send long data without a provider | The additional data is empty. |
| Send legacy data | The store does not call the provider. There is no additional data. |
| Receive with a provider | The store passes the additional data and requested finality to the provider. Empty or 32-byte pool data supplies empty additional data. A `false` response fails the message. |
| Receive without a provider | The store ignores the additional data. |

Provider calls are views and cannot change state. A revert, missing code, or response that cannot decode stops the
send or fails the received message.

The sender chooses the token args and requested finality. The connection store checks finality against the token
profile before calling the provider on send. On receive, finality has been checked only on the source chain.

Token args reach only the hook and additional pool data provider. They do not reach the remote token address provider,
fee store, or security store. A provider can use token args to build data, but excessive size or receive cost can only
revert that sender's send or fail its message.

### Empty data and migration

The provider cannot see the source pool, source token, or send time. These sources all produce empty additional data:

- A legacy pool.
- An MTTP lane with `useLegacyPoolDataOnSend` enabled.
- A source token without an additional pool data provider.

A provider that rejects empty data rejects all these messages from the chain, including old messages retried later.
These messages require either no provider or a provider that accepts empty data from that chain. Recovery is described
in [retrying messages with empty additional data](#retrying-messages-with-empty-additional-data).

Migration requires permanent compatibility with empty data:

- A legacy message can arrive years later on any participating chain.
- Each chain receives from the others, so changing providers in a different order does not avoid this requirement.
- A migrated token, or a token with a lane to or from a legacy pool, must either have no additional pool data provider
  or use one that accepts empty data. This requirement applies permanently on every chain.
- `FactoryTokenWithUIDPoolDataProvider` rejects empty data and cannot serve these tokens.

### UID validation

Source token validation compares the message with the destination's configuration. Matching mistakes on both chains
can pass this check: source token A maps to destination token B, and B expects A.

A provider that maps each UID to one local token checks asset identity independently of the CCIP configuration. It
rejects A's UID for B, so the message fails without minting.

This protection has two requirements:

- The source provider must derive the UID from the source token. Taking it from sender-controlled token args lets the
  sender choose the identity.
- The destination token must use a provider that validates the UID.

### Factory token UID provider

`FactoryTokenWithUIDPoolDataProvider` serves every token of one factory and has no other configuration. The factory
must contain code and deploy at most one token per UID. The issuer must use the same UID for the same asset on every
chain. Token addresses can differ.

| Path | Validation |
| --- | --- |
| Send | The provider returns `abi.encode(uid)` for the token. The factory must map that UID back to the token, or the send stops. |
| Receive | The data must be exactly `abi.encode(uid)`, with the factory mapping the UID to the local token. Empty data and other lengths are invalid. |

The provider ignores requested finality and token args. On every participating chain:

- The remote token must be the remote factory's token with the same UID.
- That remote token must use this provider and send `abi.encode(uid)` for the same UID.
- Its record toward this chain must have `useLegacyPoolDataOnSend` disabled.
- It must have no legacy pool history that can still produce messages to this chain.

Messages from legacy pools, lanes omitting the source token, or source tokens without a compatible provider fail.
Tokens with that history need another provider, or none.

## Remote pool records

### Addresses and legacy flags

The store hashes remote pool addresses exactly as supplied, as in `TokenPool`. EVM pool addresses must use
`abi.encode(pool)`, matching the OffRamp's source pool encoding.

| Record | Pools and flags |
| --- | --- |
| Inline | Its own pool list, `allowLegacyPoolDataOnReceive` per pool, and `useLegacyPoolDataOnSend` and `trustPolicyRemotePools` for the record |
| Policy chain | Pool addresses only; no `useLegacyPoolDataOnSend` or `allowLegacyPoolDataOnReceive` |

Setting an inline record replaces all flags. Removing it deletes all flags. A record trusts its listed pools, plus
policy pools only when the inline record explicitly enables `trustPolicyRemotePools`.

- A legacy source pool needs `allowLegacyPoolDataOnReceive` enabled permanently, because its messages never expire.
- An MTTP needs `allowLegacyPoolDataOnReceive` enabled while it sends legacy data and until all such messages have executed.
- Policy pools accept only long pool data, and tokens using a policy record send only long pool data.
- A token sending to a legacy pool, or receiving legacy data, needs an inline record for that chain.
- A legacy pool belongs only on its token's inline record, never on a policy record. Policy pools are trusted for
  every assigned token.

### No default remote pool record

Remote pool policies require a record for each chain where their tokens have an MTTP. Chain selector zero reverts with
`ZeroChainSelectorNotAllowed`, as for connection policies.

A default would trust the same addresses on all chains without explicit records, including future chains. An address
on another chain can belong to a different contract.

Removing a policy chain record immediately removes its trust on that chain, with no policy fallback. The four other
configuration types can use policy defaults.

### Combining inline and policy pools

An inline record replaces the entire policy record by default. It must list every pool the token needs, including the
remote MTTP. Policy changes do not affect that lane's trusted pools.

With `trustPolicyRemotePools` enabled:

- The lane trusts the union of the inline pools and the policy chain record's pools.
- The inline record still controls `useLegacyPoolDataOnSend`.
- If both records list a pool, its inline `allowLegacyPoolDataOnReceive` value takes precedence.
- Policy pool changes apply immediately to the token.

This flag is valid only where the remote token uses a pool from that policy record, or will migrate to one.
Elsewhere, it adds trust in pools that never send for the token.

An inline record with `trustPolicyRemotePools` enabled must list only:

- The token's remote legacy pools.
- A remote MTTP while it sends legacy data and until all such messages have executed.

Removing a pool from a policy does not remove an inline copy. Disabling `allowLegacyPoolDataOnReceive` also does not
remove trust: long pool data is still accepted. An incident must remove the pool from every record that lists it. See
[response to a compromised remote pool](#incident-a-compromised-remote-pool).

## Destination overhead records

A destination overhead record has no disabled state. It must have:

- `destGasOverhead` greater than zero.
- `destBytesOverhead` of at least 32, because an MTTP sends at least 32 bytes of pool data.

| Lane | Required byte allowance |
| --- | --- |
| Sends the source token | 192 bytes plus additional pool data padded to whole words |
| Sends to a non-MTTP pool | 32 bytes |

For example, a provider returning 32 bytes needs a 224-byte allowance. See
[lanes to non-MTTP pools](#lanes-to-non-mttp-pools)
for the legacy requirement.

A 32-byte allowance on a lane sending the source token makes every send revert before the burn. No funds move. Tooling
must detect this through `getFeeBatch` before opening the lane, as described in the [workflow checks](#workflow-rules).

## Rate limits

### Bucket scope and amounts

Each bucket is keyed by token, remote chain, direction, and finality class. The classes are default finality and FTF.
The key does not include the pool.

- Remote pools on the same lane consume the same bucket for a given direction and finality class.
- Moving a token to another MTTP through the TAR preserves its buckets if both pools use the same stores.
- Tokens assigned to one policy have separate buckets. Each can consume the full configured limit, so N tokens can
  transfer N times that limit.
- Policy defaults also create independent limits for each chain. Buckets do not combine raw amounts of different
  tokens or chains.

Amounts use local token units, as in `TokenPool` since 1.6.1:

| Direction | Amount consumed |
| --- | --- |
| Outbound | Amount after the token fee |
| Inbound | Amount after conversion to local decimals |

### Record and configuration updates

Inline records store their configurations in the token's buckets. Policy records store configurations separately.
Both use the validation rules of `RateLimiter`.

| Operation | Effect |
| --- | --- |
| Set a record | Writes all four configurations: both directions and both finality classes. |
| Update one configuration | Changes one direction and finality class in an existing record. Preserves the other three. |
| Delete one configuration | Not supported. |

An inline update requires the token's inline record for that chain. A policy update requires the policy's record for
that chain; selector zero identifies its default record.

A full record replacement can overwrite another writer's changes to the other three configurations. An individual
update preserves those values. Every transfer uses a configuration set by a PCM role.

If the selected record's FTF configuration is disabled, an FTF transfer uses that record's default-finality bucket,
as in `TokenPool`. For policy records, this decision reads the policy configuration, not the bucket's cached value.

### Stops and pauses

A **stop** is an enabled configuration with capacity 0 and rate 0. Stopping all transfers in a direction requires
stops for both finality classes. A policy stop affects every token using the record.

A stop is a configuration change, so lifting it does not require `UNPAUSER_ROLE`. Pauses do not change buckets;
buckets continue to refill while paused.

### Refills and configuration changes

A bucket configuration consists of its enabled flag, capacity, and rate.

- Writing a different configuration sets the bucket balance to the new capacity. Even lowering capacity can increase
  the remaining balance.
- Writing an identical configuration preserves the balance. Repeating a write does not add capacity.
- `TokenPool` instead resets the balance to capacity on every set, including unchanged values.

Inline and policy updates take effect differently:

| Record | When the bucket changes |
| --- | --- |
| Inline | A record set or individual update writes each changed configuration to its bucket immediately. |
| Policy | A write changes only the policy record. The next transfer using a bucket compares its stored configuration with the selected record. If different, it copies the new configuration and starts full. Otherwise, it preserves the bucket state. |

A policy configuration that changes and changes back before the next consumption does not refill that bucket.
Changing a token's selected record likewise cannot give it more than a full bucket. The balance never exceeds what it
would be if every configuration or record change reset it to capacity.

Further rules:

- There is no function that only refills a bucket. Changing its configuration refills it.
- Removing and recreating an inline record in one call also refills its buckets.
- Removing an inline record deletes its buckets. Removing a policy record or policy does not change any bucket.
- Bucket views show a pending policy configuration change as a full bucket with the selected configuration.
- A new security store starts with full buckets. No function copies live bucket state into it.

## Finality

Allowed FTF finality belongs to the token profile. It describes the source-chain reorganization risk that the token
accepts, so it applies to every remote chain, as in `TokenPool`.

- Registration sets the value, and later profile updates can change it. Any value is accepted, as in `TokenPool`.
- Zero permits only default finality. An unregistered token also permits only default finality.
- `getFee`, `getFeeBatch`, and `lockOrBurn` check requested finality. The fee views check it even when the token has
  no lane.
- Receive does not enforce this local setting. It describes this chain as a source, not the remote source chain.

An MTTP can therefore receive an FTF message for a token that permits only default finality locally. Legacy pools do
the same.

The v2 OffRamp passes the message's finality to the MTTP. An FTF message selects the inbound FTF configuration,
subject to the disabled-FTF fallback described in [rate limits](#rate-limits).

A legacy source pool without `IPoolV2` support can send only with default finality and empty token args. The v2
OnRamp rejects other values before calling that pool.

## Fees

### Quotes and token fees

`getFee` combines the token's fee and destination overhead records for the chain.

| Required configuration | Result |
| --- | --- |
| Lane, fee record, and destination overhead record all exist | Returns the configured fees and overheads. |
| Any is missing | Reverts with `LaneFeeNotFound`; the send fails. |

A fee record has two pairs of values: a USD fee and a token fee in basis points for default finality, and another pair
for FTF. Default-finality transfers use the first pair. Other requested finalities use the FTF pair.

```text
token fee = floor(send amount * bps / 10,000)
```

Each bps value must be less than 10,000. A fee record has no disabled state; zero fees mean no pool fee.

The pool burns the amount after the token fee. Its remaining token balance consists of fees. A caller with
`FEE_COLLECTOR_ROLE` withdraws the entire balance of each listed token through the PCM. Tokens sent directly to the pool also
count as fees, as in `TokenPool`.

Only the pool's current PCM can withdraw fees. The PCM checks that the pool names it as its manager before forwarding
the withdrawal. After a PCM change, the old PCM loses that access.

### Pricing through v2 ramps

The v2 OnRamp selects the pool fee and destination overheads as follows:

| Source pool | Fee and overhead source |
| --- | --- |
| MTTP | Only MTTP `getFee`, including lanes to legacy pools |
| Legacy pool with an enabled `IPoolV2.getFee` configuration | The pool's `getFee` |
| Legacy pool without `IPoolV2` support, or with a disabled fee configuration | FeeQuoter `getTokenTransferFee` |

An MTTP charges the token fee from its fee record. Missing MTTP records make `getFee` revert; the OnRamp does not
fall back to FeeQuoter token transfer fees.

The OnRamp also calls the FeeQuoter to price execution gas and obtain the fee-token price. Legacy pool fee
configurations and FeeQuoter token transfer configurations are separate from PCM records.

### Destination gas budget

The budget must cover the MTTP receive path, including:

- Reads from the connection and security stores.
- Remote token and additional pool data provider calls, where configured.
- The lane's hook, where configured.
- Copying a changed policy configuration into an inbound bucket.

The v2 OffRamp does not use `destGasOverhead` as a pool call limit. The value is used to price execution, and the
executor must supply enough transaction gas.

With these stores, a lane's first message uses approximately 102,000 to 195,000 gas. Enabled rate limits, an empty
receiver balance, providers, policy pool trust, hooks, and larger pool lists can increase the cost.

Every source that prices transfers of the token to this chain needs a budget covering the measured cost plus a
margin. This also applies to lanes sending only 32 bytes. The [transfer budget workflow](#setting-transfer-budgets)
covers measurement and the source configuration updates required before a pool switch.

### Source byte allowance

For long pool data, the MTTP's destination overhead record must allow at least 192 bytes plus padded additional data.
The v2 OnRamp reads this allowance through `getFee` and checks the returned pool data against it. A failed check
reverts the send and burn in the same transaction.

This allowance is necessary for sends to work. The separate [pool data size check](#pool-data-size-check) enforces the
lane's format limit before the burn. Legacy sends use 32 bytes and fit the default byte allowance.

## Security records and hooks

### CCV lists

- CCV lists use the `AdvancedPoolHooks` format. An `address(0)` entry includes the ramps' default CCVs.
- Duplicate entries revert on write, as in `AdvancedPoolHooks`.
- Lists do not depend on amount or finality. The MTTP still passes requested finality to the PCM, so a later PCM can
  select CCVs by finality without a pool replacement.

### Hook calls and compatibility

At most one hook runs per transfer. There is no global hook. Any hook address is accepted, as in
`TokenPool.updateAdvancedPoolHooks`.

The PCM returns the lane's hook, and the MTTP calls it. `MultiTokenPoolHooks` and `AdvancedPoolHooks` require authorized
callers and pass that caller to the ACE policy engine as the sender. Every pool serving a token that uses the hook
must be authorized. The [hook update workflow](#changing-hooks-and-inbound-ccvs) covers that requirement.

`MultiTokenPoolHooks` provides the allowlist, ACE policy engine, and caller checks of `AdvancedPoolHooks`, without CCV
lists. Its allowlist applies to every token using it.

`IMultiTokenPoolHooks` uses the same selectors and calldata for the two checks as `IAdvancedPoolHooks`.
`AdvancedPoolHooksExtractor` can decode them unchanged, so one policy engine can serve single-token pools and MTTPs.
The payload includes `localToken`. A policy engine that keys state by the sender must key that state by `localToken`.

| Hook | Compatibility |
| --- | --- |
| `AdvancedPoolHooks` on an MTTP | Works, but its CCV lists have no effect. |
| `MultiTokenPoolHooks` on a single-token pool | Lacks `getRequiredCCVs`, so every transfer through a v2 ramp fails. |

### External blocklist

`MultiTokenPoolHooksExternalBlocklist` adds one `IAddressBlocklist` to `MultiTokenPoolHooks`.

- The hook owner sets the blocklist. Zero disables checks; a non-zero address must contain code.
- The blocklist applies to every token using the hook. Lanes with different hooks can use different blocklists.
- An existing `MultiTokenPoolHooks` cannot add this feature. A blocklist requires a separate
  `MultiTokenPoolHooksExternalBlocklist` instance, as described in [hook updates](#changing-hooks-and-inbound-ccvs).

| Direction | Addresses checked | Timing |
| --- | --- | --- |
| Send | Local sender and remote receiver | Before the burn |
| Receive | Local receiver and remote sender | Before the mint |

Checks run even without a policy engine. An allowlist entry does not override a block.

| Remote address length | Check |
| --- | --- |
| 20 bytes | The hook checks the address as supplied. |
| 32 bytes | The hook checks the low 160 bits, including for a non-EVM address such as Solana. |
| Other | The hook skips the remote address but still checks the local address. |

A blocked address, reverting blocklist, or undecodable response stops the send or fails the received message. The
blocklist is local, and each call reads its current state. A failed message can be retried after the address is
unblocked; this does not refund the source burn.

### Changing inbound CCVs

The v2 OffRamp reads inbound CCVs from `getRequiredCCVs` on the TAR pool at execution time. An MTTP returns the inbound
CCVs from the token's security record for that chain, regardless of the source pool. A legacy pool without `IPoolV2`
support requires the default CCVs.

The source selected its outbound CCVs at send time. Changing the inbound list or TAR pool can therefore fail an
in-flight message with `RequiredCCVMissing`. It can execute after the inbound list accepts the CCVs that verified it.

- Omitting `address(0)` drops the default CCVs. A policy record update can do this for every assigned token.
- Pending messages remain executable across a pool switch only if the new pool's inbound CCVs accept their
  verification. The [hook and CCV workflow](#changing-hooks-and-inbound-ccvs) preserves the old pool's requirements
  until those messages have executed.

## Pause

| Pause | Scope |
| --- | --- |
| Token | One token on every chain, outbound, inbound, or both |
| Chain | Every token on one chain, outbound, inbound, or both |
| Global | Every token on every chain, in both directions |

Global pause is a separate flag. It preserves token and chain pause flags, which still apply after global unpause.
It covers every pool and PCM using the connection store, including tokens and chains added later.

- Only pause functions change pause state. Registration, deregistration, and configuration changes preserve it.
- Tokens can be paused before registration.
- Configuration writes and token or chain pause changes remain available during a global pause.
- Pausing chain selector zero affects no transfer, because no lane uses it.
- A pause call that changes nothing emits no event.
- A new connection store starts with all pauses disabled.

A token whose TAR pool is legacy does not read these flags. Only its legacy pool, that pool's rate limits, or an
RMN curse can stop its transfers.

## The pool

### Interfaces and transfer checks

The MTTP implements `IPoolV1` and `IPoolV2`. V2 ramps call its `IPoolV2` functions. Both interfaces read the same
records, but the `IPoolV1` transfer functions have these fixed values:

| `IPoolV1` function | Behavior |
| --- | --- |
| `lockOrBurn` | No token fee, default finality, and empty token args |
| `releaseOrMint` | Default finality |

Only ramps can call the transfer functions. As in `TokenPool`, the pool calls the token and hook after the PCM consumes
the rate limit. Both directions check the token's burn flavor.

Through ERC-165, the pool reports `IMultiTokenPool` and `IGetPoolConfigManager` in addition to the `TokenPool`
interfaces.
Legacy pools report neither. Tooling uses `IMultiTokenPool` to identify an MTTP.

### Token support

`isSupportedToken` returns true for a registered token whose flavor the pool supports. It ignores pauses and lanes.
`TAR.setPool` calls it, so registration must precede the switch. A registered token without lanes can still be selected;
its sends then revert.

Three parties control whether a pool serves a token:

| Party | Action |
| --- | --- |
| PCM roles | Register and connect the token. |
| Token admin | Select the pool in the TAR. |
| Issuer | Grant the pool mint and burn roles. |

### OffRamp authorization

The MTTP accepts receive calls only from OffRamps listed on its Router. A v2 message names its destination OffRamp.
Another OffRamp cannot execute it.

Every OffRamp with pending messages must remain on the pool's Router. Messages never expire. This requirement also
applies to legacy pools. The [contract replacement workflow](#replacing-contracts) checks it before Router changes.

## Views

### Missing configuration and reverts

Views return zero or empty values for unknown tokens, missing lanes, or missing records, as `TokenPool` views do for
unknown chains. The PCM and pool views have only these exceptions:

| View | Revert condition |
| --- | --- |
| `getFee` and `getFeeBatch` | Requested finality is not allowed. |
| Pool `getFee` | Lane, fee record, or destination overhead record is missing. |
| Pool `isSupportedChain` and `getRemoteToken` overloads taking only a remote chain | Always revert. |
| `getLaneConnectionBatch` and pool views that call it | Certain provider failures, described below. |

Onchain, only `USDCTokenPoolProxy` calls the chain-only `getRemoteToken`, and only for its own pools. The overload
with the local token is the supported interface for tooling.

Views shared with `TokenPool` retain its names and return types, with the local token added as the first argument.
Tooling can reuse its existing `TokenPool` handling.

### Lane views and record views

Lane views show the values a transfer would use for each token and chain. The PCM gets lane availability from the
connection store and returns zero values where no lane exists. Fee and security store views read only their own
records; the PCM applies the lane filter.

Two views need special handling:

| View | Behavior |
| --- | --- |
| `getLaneConnectionBatch` | Shows pause flags even without a lane; all other fields are zero or empty there. |
| `getLaneRecordLevelBatch` | Shows each type's selected record even without a lane. This lets tooling check records before opening one. |

With `trustPolicyRemotePools` enabled, `getLaneConnectionBatch` lists inline pools first, then policy pools not already
listed. Each appears once with its effective `allowLegacyPoolDataOnReceive` value. `getLaneRecordLevelBatch` still
reports `Inline`.

A connection alone establishes a lane. Without a remote pool record:

- `hasConnection` is true.
- Lane views still show the selected fees, CCVs, and buckets.
- Every transfer reverts with `RemotePoolRecordNotFound`; no send or mint succeeds.
- `getLaneRecordLevelBatch` exposes the missing record to tooling.

The PCM uses this same connection-based filter for `getFeeBatch`, `getTokenTransferFeeConfigBatch`,
`getCurrentRateLimiterStateBatch`, and `getLaneVerificationBatch`. Ramps read CCVs and fees through the pool. Among PCM
views, only `getLaneConnectionBatch` reads pause state.

### Provider responses in views

If the remote token address provider fails, `getLaneConnectionBatch` returns `hasConnection` true and an empty remote
token, with these exceptions:

- A response that cannot decode as `bytes` makes the view revert.
- A provider that exhausts gas or returns a very large response can also make the view revert.

The token-aware `isSupportedChain` and `getRemoteToken` pool views use `getLaneConnectionBatch`, so they have the same
behavior.

### Configuration inspection

The PCM exposes all configuration and state needed by operators:

- Policy records and token inline records for every type, including chains without lanes.
- The number of tokens assigned to each policy.
- The TAR and three stores through `getStaticConfig`.

Stores also expose their stored records directly for inspection.

`getAllConfiguredTokens` and `getSubsetTokens` return token lists in pages, using `startIndex` and `maxCount`.
Removing an entry can move the last token into its place. A scan across different blocks can therefore miss a token. The
[inspection workflow](#inspecting-configuration-and-indexing-events) uses one block for the whole scan.

## Burn flavors

The four burn flavors match the corresponding single-token burn and mint pools.

| Flavor | Burn address |
| --- | --- |
| `Burn`, `BurnFrom`, or `BurnWithFrom` | Zero |
| `BurnToAddress` | The burn address; any address is accepted, as in `BurnToAddressMintTokenPool`. |

`BurnMintMultiTokenPool` supports all four burn flavors. `BurnFrom` and `BurnWithFrom` use the pool's allowance to
itself. Before burning, the pool grants itself the maximum allowance if the current allowance is insufficient.
It never revokes that approval. No other address can spend it because the pool is both holder and spender.

## Roles

Any one role in a row authorizes the action. The pool and store owners own their respective contracts; all other roles
belong to the PCM.

| Action | Authorized roles |
| --- | --- |
| Set a pool's Router or PCM | Pool owner |
| Add or remove a store's authorized PCMs | Store owner |
| Grant or revoke PCM roles | `DEFAULT_ADMIN_ROLE` |
| Register or deregister tokens; change profiles | `CONFIG_ROLE`, `TOKEN_MANAGER_ROLE` |
| Set or remove connection policy records; remove policies; set providers; add or remove subset tokens | `CONFIG_ROLE`, `POLICY_ROLE`, `CONNECTION_ROLE` |
| Change connection policy assignments; set or remove inline connection records | `CONFIG_ROLE`, `TOKEN_MANAGER_ROLE`, `CONNECTION_ROLE` |
| Set or remove remote pool policy records; remove policies | `CONFIG_ROLE`, `POLICY_ROLE`, `REMOTE_POOL_ROLE` |
| Change remote pool policy assignments; set or remove inline remote pool records | `CONFIG_ROLE`, `TOKEN_MANAGER_ROLE`, `REMOTE_POOL_ROLE` |
| Set or remove rate limit policy records; remove policies; update individual configurations | `CONFIG_ROLE`, `POLICY_ROLE`, `RATE_LIMIT_ROLE` |
| Change rate limit policy assignments; set or remove inline rate limit records; update individual inline configurations | `CONFIG_ROLE`, `TOKEN_MANAGER_ROLE`, `RATE_LIMIT_ROLE` |
| Set or remove fee policy records; remove policies | `CONFIG_ROLE`, `POLICY_ROLE`, `FEE_ROLE` |
| Change fee policy assignments; set or remove inline fee records | `CONFIG_ROLE`, `TOKEN_MANAGER_ROLE`, `FEE_ROLE` |
| Set or remove destination overhead policy records; remove policies | `CONFIG_ROLE`, `POLICY_ROLE`, `DEST_OVERHEAD_ROLE` |
| Change destination overhead policy assignments; set or remove inline destination overhead records | `CONFIG_ROLE`, `TOKEN_MANAGER_ROLE`, `DEST_OVERHEAD_ROLE` |
| Set or remove security policy records; remove policies | `CONFIG_ROLE`, `POLICY_ROLE`, `SECURITY_ROLE` |
| Change security policy assignments; set or remove inline security records | `CONFIG_ROLE`, `TOKEN_MANAGER_ROLE`, `SECURITY_ROLE` |
| Pause tokens, chains, or globally | `DEFAULT_ADMIN_ROLE`, `PAUSER_ROLE` |
| Unpause tokens, chains, or globally | `DEFAULT_ADMIN_ROLE`, `UNPAUSER_ROLE` |
| Withdraw fees | `FEE_COLLECTOR_ROLE` |

The PCM uses OpenZeppelin `AccessControlDefaultAdminRules`. Admin transfer has two steps, with an initial delay of zero.
An account with `DEFAULT_ADMIN_ROLE` can change the delay.

`DEFAULT_ADMIN_ROLE` grants no configuration permission by itself. An account with this role can grant itself any role
in one call without a delay.

## Trust

### Ownership and trusted parties

Each customer owns and operates separate pools, PCMs, and stores. Customers do not share these contracts.

The system trusts:

- The issuer and PCM role holders.
- Each token's TAR admin.
- Registered tokens.
- Pools listed in remote pool records.
- Additional pool data providers, remote token address providers, and anyone who can change their responses.
- Hooks, their owners, and their policy engines and blocklists.

The MTTP owner selects the PCM. PCM roles configure remote tokens and trusted pools, so the MTTP owner has full control
over the pool's tokens.

An authorized store caller can change any configuration and consume any rate limit in that store. One admin must
therefore own the pools and stores and hold the PCM's `DEFAULT_ADMIN_ROLE`.

### Caller and token isolation

The PCM accepts validation calls only from the token's TAR pool. A second pool using the same PCM cannot consume that
token's buckets. The PCM must use the same TAR as the ramps; its constructor does not verify this.

A transfer of one token cannot change another token's buckets or fees. Tokens share configuration and pause state only.

The three stores have the same registered tokens unless an authorized caller writes directly to one store.
If the lists differ:

| Operation after lists diverge | Result |
| --- | --- |
| Deregistration, policy assignment, inline record set or removal, or individual inline rate limit update | Reverts in a store where the token is not registered. |
| Profile set | Registers the token in stores where it is missing. The connection and security stores reject different decimals for an already registered token; the fee store accepts the set. |

### Remote pool trust

Legacy data does not identify the source token. Enabling `allowLegacyPoolDataOnReceive` therefore trusts that pool to
send the correct asset. Only inline records have this flag, and they must list only their own token's pools:

- Its remote legacy pools.
- Remote MTTPs serving the token.
- With policy trust enabled, an MTTP belongs on the inline list only while it sends legacy data or those messages
  remain unexecuted.

Trusting another token's legacy pool allows that token's burn to mint this token.

Migration step 7 makes legacy pools trust a remote MTTP. Those legacy pools cannot check source identity. An incorrect
remote-token mapping on the MTTP can therefore mint the wrong token on the legacy chain.

`trustPolicyRemotePools` extends trust to every pool in the policy's record for that chain. A compromised policy pool
can mint the token. This trust is valid only if the remote token uses a policy pool or will migrate to one. See
[remote chains with legacy pools](#remote-chains-with-legacy-pools).

### Compromised pools and owners

A compromised pool address must remain untrusted permanently. Messages do not expire, so forged messages can execute
as soon as a record or legacy pool restores that trust. This includes legacy pools retained for rollback.

The destination cannot distinguish honest pending messages from forged ones. Recovery requires a new pool address.
Honest unexecuted messages need resolution outside the contracts, such as a refund. The
[incident workflow](#incident-a-compromised-remote-pool) covers containment and recovery.

An MTTP owner can select any Router and PCM. A Router owner can add any OffRamp. Either owner can thereby mint tokens
that grant the MTTP a mint role, without a message. Legacy pools have the same exposure. A compromised pool owner is
therefore also an incident on its own chain. See [incident response](#incident-a-compromised-remote-pool).

## Events and observability

### Event behavior

Every state change emits an event except copying a policy rate limit configuration into a bucket during a transfer.
The transfer's consumption event follows that copy.

- Setting a record, profile, policy assignment, or provider emits an event even if the value is unchanged. Monitoring
  must not treat every set event as a state change.
- Removing a record emits one event for all its values. Removing a policy emits one event for the policy and its
  metadata.
- Stores emit configuration and consumption events for their types. The connection store also emits pause events.
  Each log identifies the emitting store.
- The PCM constructor emits `StaticConfigSet` with the TAR and store addresses.
- The pool constructor emits `RouterSet` and `PoolConfigManagerSet`, each with a zero old value.

### Indexing

The PCM's `StaticConfigSet` event identifies its stores. Configuration events from those stores include writes by
an old PCM while both managers remain authorized.

Events allow monitoring to reconstruct the PCM's resolved views, except provider-derived remote tokens. Provider
responses have no event; `getLaneConnectionBatch` exposes their current values.

The PCM has no pool list. TAR `PoolSet` events identify registered tokens' pools. Each pool exposes its PCM through
`getPoolConfigManager` and emits transfer, fee withdrawal, owner, Router, and PCM change events. The
[indexing workflow](#inspecting-configuration-and-indexing-events) combines these sources.

## Workflows

These workflows describe configuration changes and operator checks. Each step identifies the signer and call.
Unless a step says otherwise, the issuer signs the steps for one chain in one batch. No batch spans chains.

### Workflow rules

#### Message execution and retries

- The OnRamp selects the source pool from the TAR at send time. The OffRamp selects the destination pool at execution
  time, including retries.
- Messages never expire. After fixing a failure, any party can retry through `OffRamp.execute`.
- Supply enough transaction gas for pool execution. The `gasLimitOverride` argument changes only the `ccipReceive`
  gas limit. A non-zero override must be at least the message's `ccipReceiveGasLimit`.
- When waiting for messages to complete, check every message's execution state on the destination OffRamp. Include all
  in-flight messages, not just failed ones. A failed retry reverts without an event, so event counts cannot establish
  completion.

#### Checks before opening a lane

An incorrect remote token can permanently lose funds. An unprepared destination delays execution. Check remote tokens
before the first send, and prepare destinations before sources where practical. A short delay is acceptable only for
an unannounced lane.

A token or chain pause stops sends before the burn. A token can be paused before registration. Legacy pool data from
a remote MTTP is permitted only during migration or rollback. These messages provide no source token identity, so
keep that period as short as the procedure permits.

Connection policy assignments and record writes do not run the lane check. Before a lane goes live, tooling must read:

- `getTokenConfigBatch`
- `getLaneRecordLevelBatch`
- `getLaneConnectionBatch`
- `getCurrentRateLimiterStateBatch`
- `getFeeBatch`
- `getLaneVerificationBatch`

CCIP ops owns the ramps, Router, and FeeQuoter. Verify that both directions use v2 ramps. Check each source's gas and
byte allowances under [pricing through v2 ramps](#pricing-through-v2-ramps), including legacy sources priced by the
FeeQuoter. Keep OffRamps with pending messages on each pool's Router, as described in
[OffRamp authorization](#offramp-authorization).

In a hub-and-spoke deployment, spokes have lanes only to the hub.

#### Checking remote tokens

Tooling must validate the planned remote token mappings before any of these changes:

- Assigning a connection policy to a token.
- Setting a connection policy record.
- Adding a subset token.
- Changing a remote token address provider or its response.

For every token and chain affected by the change:

1. Check that the remote token exists and is the same asset as the local token.
2. Check that the remote chain maps that token back to the local token, and that its TAR selects an issuer pool.
3. Check the address encoding. EVM source tokens use `abi.encode(address)`; other chain families require their raw
   token bytes in inline connections.
4. Before setting a connection policy record, check the chain family. A connection policy using
   `SameAddressRemoteTokenAddressProvider` may contain only EVM chains. Every policy token must have the same address
   on every policy chain, including its original deployment chain.

Use `getLaneConnectionBatch` on both chains to inspect existing mappings. The planned configuration must satisfy the
[remote token requirements](#remote-token-checks-in-tooling) before it takes effect.

#### Checking remote pool compatibility

Before enabling long-format sends, tooling must read the remote TAR and require an MTTP for the token. This check
applies to these writes:

- Disabling `useLegacyPoolDataOnSend` in an inline remote pool record.
- Removing an inline remote pool record.
- Assigning a remote pool policy.
- Setting a remote pool policy record for the chain.
- Raising the lane's `destBytesOverhead` above 32.
- Opening a lane that will use a remote pool policy record.

Check only lanes affected by the write. For a policy record update, check each assigned token with a lane to that
chain and no inline remote pool record. Tokens without a lane, or whose inline record remains in use, are unaffected
by that policy update.

There is one exception: if the token is paused outbound locally, the remote TAR may name no pool. Repeat the check
before unpausing.

For a lane whose remote token still uses a non-MTTP pool:

1. **`TOKEN_MANAGER_ROLE`:** set an inline remote pool record with `useLegacyPoolDataOnSend` enabled. Keep it enabled at least
   until the remote TAR names an MTTP, following the order in [token pool migration](#token-pool-migration).
2. **`TOKEN_MANAGER_ROLE`:** set `destBytesOverhead` to 32. Do not use a remote pool policy record for this lane.
3. **`REMOTE_POOL_ROLE`:** do not create a policy record for a chain that has no MTTP.

Before enabling legacy send, tooling must check that the remote mapping for each token and chain points back to the
correct local token. A legacy pool must list a remote MTTP only when its token migrates to that MTTP.

#### Setting transfer budgets

Before a chain switches a token to an MTTP:

1. **Tooling:** measure the MTTP receive cost with that token's configuration.
2. Set every source configuration that prices transfers of that token to this chain to the measured cost plus a
   margin, for example 300,000 gas. This also applies to lanes sending only 32 bytes. The signer and configuration
   depend on the source pool:

   | Source pool | Signer | Configuration to update |
   | --- | --- | --- |
   | MTTP | `TOKEN_MANAGER_ROLE` or `DEST_OVERHEAD_ROLE` | PCM destination overhead record |
   | Legacy pool with an enabled `IPoolV2.getFee` configuration | Legacy pool owner | Pool fee configuration |
   | Legacy pool priced by the FeeQuoter | CCIP owner | FeeQuoter token transfer configuration |

3. **Tooling:** check the source quote for the gas and byte allowances. For an MTTP source, also check `getFeeBatch`.
   Long data needs 192 bytes plus padded additional data. Lanes to non-MTTP pools require a PCM
   `destBytesOverhead` of 32.

The migration workflow specifies when each side raises its byte allowance. The allowance must increase before
`useLegacyPoolDataOnSend` is disabled.

#### Updating existing records

A record set replaces all values. Tooling must:

1. Read the current record.
2. Change only the requested values.
3. Check for queued writes to the same record. A later full replacement can undo an earlier change.
4. Write the complete record back.

Require an explicit operator request before dropping a remote pool or disabling an enabled `trustPolicyRemotePools`
or `allowLegacyPoolDataOnReceive` flag. Omitting a value from the replacement removes it.

### Configuration maintenance

#### Replacing contracts

A replacement can retain the existing stores or introduce new ones. Prepare it before changing a pool's PCM or a
token's TAR pool:

1. **Deployer:** deploy a new PCM if any store changes. Its store addresses are fixed at deployment.
2. **Store owners:** authorize a new PCM on its stores before it writes configuration or validates transfers.
3. **`DEFAULT_ADMIN_ROLE` of the new PCM:** grant its roles. Roles do not carry over from the old PCM.
4. **PCM roles:** copy every configuration record into each new store. **Tooling:** compare the old and new
   configuration views before the switch. A new security store starts with full buckets, and a new connection store
   starts with all pauses disabled; live bucket state cannot be copied.
5. **Tooling:** if the new connection store changes the pool data format, verify that every receiving remote pool
   accepts that format before any pool sends it.
6. **Pool owner or token admin:** switch the pool's PCM or the token's TAR pool only after the applicable preparation
   above is complete.

A pool replacement that keeps its PCM and stores needs no configuration copy. A PCM replacement that keeps its stores
also needs no configuration copy; an account with `DEFAULT_ADMIN_ROLE` must still grant roles again.

For a Router change:

1. **CCIP owner:** before changing a Router, verify that it retains every OffRamp with pending messages.
2. **MTTP owner:** before `setRouter`, verify the same requirement on the new Router. Keep each required OffRamp
   authorized until all its messages have executed. This requirement also applies to legacy pools.

#### Changing connections

Use the [remote token checks](#checking-remote-tokens) before changes that expose a token to a policy connection or
change its provider's response.

To replace an inline connection, **`TOKEN_MANAGER_ROLE`** sets the record directly without removing it first. To remove it
and adopt a different remote token from the policy:

1. Complete any required changes to the token's connection policy, policy records, or provider in earlier calls.
2. **`TOKEN_MANAGER_ROLE`:** set the inline record to the remote token bytes returned by the policy.
3. **`TOKEN_MANAGER_ROLE`:** remove the inline record. The removal check now compares matching bytes.

Before moving a token to another connection policy or clearing its assignment, identify every lane that would close,
even temporarily. Complete steps 1 and 2 of [closing a token lane](#closing-a-token-lane) for those lanes. Then:

1. **`CONNECTION_ROLE`:** remove the token from every subset of its current policy.
2. **`TOKEN_MANAGER_ROLE`:** change or clear the connection policy assignment.

To set a remote token address provider, **`CONNECTION_ROLE`** first creates the policy if it does not exist, then calls
`applyRemoteTokenAddressProviderUpdates`. Zero removes the provider. Changes to policy records do not change it.

#### Closing a token lane

For a token lane between MTTPs:

1. **`TOKEN_MANAGER_ROLE` on both chains:** stop both outbound finality classes with inline rate limit records. Keep inbound
   transfers available while messages remain in flight.
2. **Tooling:** wait for all messages in both directions to execute, including failed messages. Read destination
   execution state as described in [message execution and retries](#message-execution-and-retries).
3. **`CONNECTION_ROLE` on each chain:** close any policy connection for the token:
   - If `onlySubsetTokens` is enabled and the token is in the subset, remove it from the subset.
   - If the record serves all policy tokens, enable `onlySubsetTokens` and add every other policy token in one batch.
4. **`TOKEN_MANAGER_ROLE` on each chain:** remove any inline connection after the policy no longer serves the token on that
   lane. An inline connection would otherwise keep the lane open.
5. **Tooling:** verify that `getLaneConnectionBatch` reports `hasConnection` false on both chains.

No connection removal may precede step 2. Removing a connection also blocks pending inbound messages, and the
contract does not check that those messages have completed.

#### Changing hooks and inbound CCVs

To add an external blocklist to a token that uses `MultiTokenPoolHooks`:

1. **Deployer:** deploy `MultiTokenPoolHooksExternalBlocklist` with the intended non-zero blocklist.
2. **Hook owner:** authorize every pool that serves a token using the new hook.
3. **`SECURITY_ROLE`:** point the affected security records to the new hook.

Authorization is also required when reusing `AdvancedPoolHooks` or `MultiTokenPoolHooks` with a new pool. Authorize the
new pool before it handles transfers.

Before a pool switch, preserve the old pool's inbound CCVs in the token's security record. Wait for in-flight messages
to execute before changing that inbound list. A policy record change affects every token using it. A list without
`address(0)` no longer includes the ramps' default CCVs.

#### Retrying messages with empty additional data

If a message fails because its additional pool data provider rejects empty data:

1. **`TOKEN_MANAGER_ROLE`:** set a provider that accepts empty data from that chain, or remove the provider.
2. **Any party:** retry the message under the rules in [message execution and retries](#message-execution-and-retries).

Migrated tokens and tokens with legacy lanes must remain compatible with empty data permanently.

#### Inspecting configuration and indexing events

1. Start with the PCM address. Read `getStaticConfig` for the TAR and stores, or `StaticConfigSet` for the event index.
2. Index the PCM and all three stores. Store events also capture writes from an old PCM while both remain authorized.
3. Read `getLaneConnectionBatch` for provider-derived remote tokens. Their values cannot be reconstructed from events.
4. Use TAR `PoolSet` events to find registered tokens' pools. Read each pool's `getPoolConfigManager` and index its
   transfer, fee withdrawal, owner, Router, and PCM change events.
5. For `getAllConfiguredTokens` and `getSubsetTokens`, read every page at the same block. Removing an entry can move the
   last token into its place, so pages from different blocks can miss tokens.

### Token expansion

#### Starting state and goal

- Each chain has one issuer MTTP, one PCM, and three stores.
- Existing tokens move between the hub and each spoke, with policies of all six types on each chain.
- A new burn and mint token T is deployed and minted on every chain where it will transfer.

**Goal:** enable T between the hub and each participating spoke.

#### Token expansion phase 1: Prepare each chain

The chain order does not matter.

1. **`PAUSER_ROLE`:** call `pauseTokens([T], Outbound)` before step 6. Selecting the pool in the TAR makes T usable.
2. **Role admin of T:** grant the MTTP mint and burn roles, for example with `grantMintAndBurnRoles(mttp)`.
3. **Owner or CCIP admin of T:** call `registerAdminViaOwner(T)` or `registerAdminViaGetCCIPAdmin(T)` on the registry
   module. Then the **token admin** calls `TAR.acceptAdminRole(T)`.
4. **`TOKEN_MANAGER_ROLE`:** register T with `applyTokenProfileUpdates([], [profileOfT])`. Then assign its five
   non-connection policies through:
   - `applyTokenRemotePoolPolicyUpdates`
   - `applyTokenRateLimitPolicyUpdates`
   - `applyTokenFeePolicyUpdates`
   - `applyTokenDestOverheadPolicyUpdates`
   - `applyTokenSecurityPolicyUpdates`
5. **`TOKEN_MANAGER_ROLE`:** set any inline records, with inline connections last. Then assign T's connection policy through
   `applyTokenConnectionPolicyUpdates`.
   - Inline connections run the lane check, so all other record types must already exist through step 4 or this step.
   - The connection policy assignment opens its lanes immediately, without a lane check.
   - For each required chain with `onlySubsetTokens` enabled, **`CONNECTION_ROLE`** adds T through
     `applySubsetTokenUpdates`. This runs the lane check.
   - Before step 5, tooling verifies that the remote token exists and is the same asset for every lane the step
     will open. Policy views must show records for every chain in the connection policy, or defaults where supported.
6. **Token admin:** call `TAR.setPool(T, mttp)`. Its `isSupportedToken` check requires step 4. The chain can now
   receive T.

#### Token expansion phase 2: Verify all chains

Run the [token expansion checks](#token-expansion-checks) on every participating chain.

#### Token expansion phase 3: Enable sends

Start only after phase 2 passes on every participating chain. The chain order does not matter; every destination is
ready to receive T.

7. **`UNPAUSER_ROLE`:** call `unpauseTokens([T], Outbound)`.
8. Send a small amount on each lane in both directions. Then announce T.

#### Why pause outbound

`setPool` enables both directions at once. No order of these calls makes every destination ready before its source.
Without a pause, an early message fails where T has no TAR pool yet and needs a retry after `setPool`.

That delay is acceptable only before T is announced. Use a token pause because a chain or global pause would also
stop other tokens.

#### Remote tokens and decimals

| Remote token setup | Required action before or during step 5 |
| --- | --- |
| T has the same address everywhere, and the policy uses `SameAddressRemoteTokenAddressProvider`. | Assign the connection policy. |
| The policy uses another provider. | Ensure it returns T's remote token for every required chain before step 5. |
| T has different addresses on some chains. | Set inline connections on both ends of each affected lane, using `abi.encode` of the remote address. These run the lane check. |

Set inline connections before assigning the connection policy. Their priority prevents a policy from temporarily
selecting the wrong token. Do not replace a shared provider just to support T; that changes resolution for every policy
token.

Step 4 checks T's decimals against `decimals()`. Its rate limit policy must use the same local decimals, or assignment
reverts. If one chain uses different decimals, select a matching rate limit policy there or use inline rate limits.
Limits must represent the same token amounts across chains.

#### Adding a spoke later

If a spoke is not ready:

- Do not open T's lane to it on the hub. The connection record must use `onlySubsetTokens`, and step 5 must omit T
  from that subset.
- If the record applies to every policy token, wait for the spoke or assign T another connection policy.
- Phases 2 and 3 cover only participating chains.
- If T will never exist at the lane's configured remote token address on that spoke, never open that lane.

When the spoke is ready:

1. Run steps 1–6 for T on the spoke.
2. Run phase 2 checks there, except the paired `hasConnection` check, which must wait for the hub update.
3. Tooling checks T's remote token on the hub. **`CONNECTION_ROLE`** then adds T to the subset with
   `applySubsetTokenUpdates`, which runs the lane check. This matches step 18 of [chain expansion](#chain-expansion).
4. Run phase 2 checks on the hub for this lane.
5. Run phase 3 on the spoke.

#### Token expansion checks

Run these checks in phase 2:

| View or source | Expected result |
| --- | --- |
| `getTokenConfigBatch` | Correct decimals, flavor, policy names, and `supportedChains` for T |
| `getLaneRecordLevelBatch` | No `None` record level on a lane of T |
| `getLaneConnectionBatch` | Correct remote token and MTTP; outbound paused, inbound unpaused |
| `getCurrentRateLimiterStateBatch` | Each inbound limit covers the other side's outbound limit. |
| `getFeeBatch` | Correct fees, `isEnabled` true, and correct byte allowance on every lane |
| `getLaneVerificationBatch` | Correct CCVs and hook; the hook accepts T. |
| Remote TAR | The remote token maps to the remote MTTP. |
| `getLaneConnectionBatch` on both ends | Matching `hasConnection` for T |

The byte allowance is 192 plus padded provider data toward an MTTP, and 32 toward a non-MTTP pool. See
[lanes to non-MTTP pools](#lanes-to-non-mttp-pools).

#### Timing and rollback

Before step 7, sends revert before the burn. Before any chain reaches step 7, T has no in-flight messages, so
configuration changes and rollback cannot affect pending transfers.

T has separate buckets. The first transfer using a policy bucket copies the policy configuration into that bucket,
which starts full.

**Before step 7:**

1. Keep T paused and call `TAR.setPool(T, address(0))`.
2. **`CONNECTION_ROLE`:** remove T from each subset added in step 5 through `applySubsetTokenUpdates`.
3. Clear T's connection policy.
4. Remove its inline records and clear its other policies.
5. Deregister T.

Clear the connection policy before removing inline connections, because their removal checks the selected policy.

**After step 7:** call `pauseTokens([T], Outbound)` on each chain. Keep inbound open for in-flight messages.

### Chain expansion

#### Starting state and goal

- Tokens transfer between a hub and its spokes.
- A new chain will become a spoke. Initially, only a selected subset of tokens may have lanes to it.
- The selected tokens are deployed on the new chain with no supply.

**Goal:** enable the selected tokens between the hub and the new chain, then add the others in later batches.

CCIP ops must first connect both directions through v2 ramps, the Router, and the FeeQuoter.

#### Chain expansion phase 1: Prepare the new chain

Complete this phase before opening the hub's lanes.

1. **Deployer:** deploy:
   - `ConnectionPolicyStore`
   - `FeePolicyStore`
   - `SecurityPolicyStore`
   - `PoolConfigManager(tar, connectionStore, feeStore, securityStore, admin)`
   - `BurnMintMultiTokenPool(rmnProxy, router, pcm)`

   Use the ramps' TAR and Router and the issuer's admin. If no suitable remote token address provider exists, deploy
   one, such as `SameAddressRemoteTokenAddressProvider`.
2. **Deployer:** authorize the PCM on all three stores with `applyAuthorizedCallerUpdates`. Transfer the pool and
   stores to the issuer admin, who accepts ownership.
3. **`DEFAULT_ADMIN_ROLE`:** use `grantRole` to assign the same roles as on the other spokes.
4. **`POLICY_ROLE` or the role of each type:** create the policies:
   - **Remote pool:** a hub record containing `abi.encode(hubMttp)`. Policy records have no legacy flags, so this lane
     sends and accepts only long data.
   - **Rate limit, fee, destination overhead, and security:** hub records or policy defaults. Allow 192 bytes plus
     the largest padded provider output among the tokens. Less space causes long-format sends to revert.
   - **Connection:** a hub record with `onlySubsetTokens` disabled, since every token on the new chain uses the hub.
     Set its provider with `applyRemoteTokenAddressProviderUpdates`; it must resolve tokens on the hub.
5. **`TOKEN_MANAGER_ROLE`:** register the selected tokens with `applyTokenProfileUpdates`, then assign their policies, with
   connection last. The connection assignment opens hub lanes without a lane check.
6. **Role admin of each token:** grant the MTTP mint and burn roles.
7. **Owner of each token, then token admin:** register the TAR admin, call `acceptAdminRole`, then
   `setPool(token, mttp)`. Pool selection requires step 5.

#### Chain expansion phase 2: Configure the hub

Run steps 8–13 in one batch.

8. **`PAUSER_ROLE`:** call `pauseChains([newChain], Both)`. No traffic uses this chain yet. The pause allows inspection of
   the open lanes before sends begin.
9. **`REMOTE_POOL_ROLE`:** use `applyRemotePoolPolicyConfigUpdates` to add a record for the new chain to each selected
   token's remote pool policy. List `abi.encode(newChainMttp)`. Remote pool policies have no defaults.
10. **`RATE_LIMIT_ROLE`, `FEE_ROLE`, `DEST_OVERHEAD_ROLE`, and `SECURITY_ROLE`:** add a record for the new chain
    where no default exists or different values are needed.
11. **Provider owner:** ensure the connection provider resolves tokens on the new chain.
    `SameAddressRemoteTokenAddressProvider` needs no change.
12. **`CONNECTION_ROLE`:** use `applyConnectionPolicyConfigUpdates` to create a record for the new chain with
    `onlySubsetTokens` enabled. Its subset is initially empty.
13. **`CONNECTION_ROLE`:** add the selected tokens with `applySubsetTokenUpdates`. The PCM checks each lane.

Steps 9 and 10 do not open lanes. Only connection records do.

#### Chain expansion phase 3: Enable the hub

Use a second batch after phase 2.

14. Tooling runs the [chain expansion checks](#chain-expansion-checks). Then **`UNPAUSER_ROLE`** calls
    `unpauseChains([newChain], Both)`.
15. Send a small amount in each direction.

#### Subsets and later batches

Subsets let the selected tokens retain their existing policy and remote token rule. Each addition runs the lane check.
An inline connection also runs that check, but specifies the address directly; use it only if the policy returns the
wrong address.

A separate connection policy for the selected tokens would duplicate every spoke record, and assigning it would not
check lane completeness.

For later batches:

16. On the new chain, repeat steps 5–7 for the new tokens.
17. On the hub, repeat steps 9–12 for policies that still lack records for the new chain.
18. On the hub, add the tokens with `applySubsetTokenUpdates`. Do not pause the new chain; its existing lanes are live.

Keep `onlySubsetTokens` enabled after all tokens join. Every later token will then require an explicit addition
and lane check. Disabling the flag opens the chain to all current and future policy tokens without that check.

If such a token is not configured on the new chain, a send burns on the hub and waits for destination setup before
execution can succeed.

#### Remote chains with legacy pools

The main procedure assumes an MTTP on the hub. When a remote chain still uses legacy pools, including during the
issuer's first MTTP deployment, apply these rules:

- For each token that still uses a legacy pool on that chain, create an inline remote pool record. List its legacy
  pool with `allowLegacyPoolDataOnReceive` enabled, and enable `useLegacyPoolDataOnSend`.
- The policy record lists only that chain's MTTP. Since it must contain a pool, do not create it until an MTTP exists.
  Until then, inline records provide all trust.
- **`REMOTE_POOL_ROLE`** creates the policy record before the spoke switches its first token. **`TOKEN_MANAGER_ROLE`** must
  then enable `trustPolicyRemotePools` on each migrating token's inline record.
- Before each spoke switch, tooling checks that the token's hub-side `getLaneConnectionBatch` includes the spoke MTTP.
- Follow [token pool migration](#token-pool-migration) for each token. The hub switches first; the remote spoke switches
  second.
- At its switch, the spoke disables `useLegacyPoolDataOnSend` toward the hub. Its inline record must list the hub MTTP
  with `allowLegacyPoolDataOnReceive` enabled until migration step 12.
- The hub keeps `useLegacyPoolDataOnSend` enabled until the issuer rules out rollback on both chains. Migration step 10
  disables it and raises `destBytesOverhead` for long data.
- Keep the legacy pool on the token's inline record permanently. See
  [policy adoption](#move-a-migrated-token-to-policies).

#### Remote tokens and timing

With `SameAddressRemoteTokenAddressProvider`, every token must have the same address on the hub and the new chain.
Otherwise, use a provider that returns the other chain's address or inline connections. Tooling checks these mappings
before steps 5, 13, and 18.

| Point in the procedure | Transfer behavior |
| --- | --- |
| Ramps open before step 13 | Hub sends to the new chain revert before the burn. |
| From step 5 | The new chain has hub lanes but no supply, so it cannot send messages. |
| Steps 8–14 | Hub sends to the new chain revert before the burn. |
| After step 14 | Missing destination setup can fail a message; retry after fixing it. |

If the issuer mints on the new chain before the hub opens, its **`PAUSER_ROLE`** must call `pauseChains([hub], Outbound)`
before step 5. Its **`UNPAUSER_ROLE`** calls `unpauseChains([hub], Outbound)` after step 14.
Without that pause, early messages fail on the hub and need retries after step 14.

#### Chain expansion checks

| Where and when | Check |
| --- | --- |
| New chain, after step 7 | PCM `getTokenAdminRegistry` matches the ramps' TAR, which maps each token to the MTTP. |
| New chain, after step 7 | `getLaneRecordLevelBatch` shows every type for each token. `getLaneConnectionBatch` shows the hub token and MTTP. |
| Hub, before step 13 | For each token in the batch on the new chain, `getLaneRecordLevelBatch` shows the five other types. Its CCVs support the new chain. Each inbound limit covers the other side's outbound limit. |
| Hub, before step 14 | For all tokens from `getAllConfiguredTokens`, only the selected tokens have `hasConnection` to the new chain. Their remote tokens and MTTP are correct. The complete `getSubsetTokens` scan lists the selected tokens. |
| Both chains, before step 14 | `getLaneConnectionBatch` reports matching `hasConnection` for each token. |

A lane open on only one side can burn there while destination execution fails. The message needs the other side's
lane to open before it can succeed.

#### Chain expansion rollback

| When | Action |
| --- | --- |
| Before step 14 | Remove the subset tokens or the new chain's connection policy record. No messages are in flight. |
| After step 14 | Call `pauseChains([newChain], Outbound)` on the hub. Keep all records on both chains until in-flight messages execute. |

To close one token's lane, follow [closing a token lane](#closing-a-token-lane). Stop both outbound finality classes
on both chains and wait for all in-flight messages before removing any connection.

### Token pool migration

#### Starting state and goal

- An issuer has burn and mint tokens, each with a single-token pool on every participating chain.
- Pool versions can differ across chains. All lanes use v2 ramps.
- Messages sent through v2 ramps are in flight, including failed messages awaiting retries.

**Goal:** move the tokens to one MTTP, one PCM, and three stores per chain without losing funds or stranding messages.

#### Migration prerequisites

| Owner | Requirement |
| --- | --- |
| Issuer | Be each token's TAR admin, legacy pool owner, and token role admin. |
| Tooling | Before switching a chain, measure the MTTP receive gas for each token's configuration. |
| Source configuration owners | Before switching a chain, set the destination gas estimates under [setting transfer budgets](#setting-transfer-budgets). This includes legacy sources priced through the pool or FeeQuoter. |
| CCIP ops and MTTP owner | Keep every OffRamp with pending messages on the MTTP's Router. |
| Legacy pool owner | Upgrade every 1.5.0 pool to 1.5.1 or later before phase 2, including chains that will remain legacy. See [the upgrade procedure](#upgrade-pools-older-than-151). |
| `TOKEN_MANAGER_ROLE` | Use no additional pool data provider, or one that accepts empty data. This remains a permanent requirement. |
| Issuer | Before phase 1, choose fees for each token on every lane, including lanes to legacy pools. Fees can be zero. |
| Hook owner | Ensure each legacy hook's `getThresholdAmount` returns 0. MTTP has no threshold amount and cannot copy threshold CCV lists. |

See [destination gas budgets](#destination-gas-budget), [source byte allowances](#source-byte-allowance), and
[empty data compatibility](#empty-data-and-migration).

#### Migration phase 1: Prepare every chain

The TAR still points to the legacy pools, so this phase does not change traffic.

1. **Deployer and `DEFAULT_ADMIN_ROLE`:** deploy the five contracts, authorize the PCM, transfer ownership, and grant roles
   as in [chain expansion](#chain-expansion).
2. **`FEE_ROLE`, `DEST_OVERHEAD_ROLE`, and `SECURITY_ROLE`:** create policy defaults:
   - Fee: the issuer's chosen fees.
   - Destination overhead: the measured MTTP receive gas and 32 bytes. The byte limit prevents long-format sends
     while the destination is still legacy.
   - Security: empty CCV lists and no hook.
3. **`TOKEN_MANAGER_ROLE`:** register the tokens with `applyTokenProfileUpdates`. Copy each legacy pool's decimals, pool
   flavor, burn address, and allowed finality. Assign the three policies with `applyTokenFeePolicyUpdates`,
   `applyTokenDestOverheadPolicyUpdates`, and `applyTokenSecurityPolicyUpdates`.
   Keep the additional pool data provider unset or compatible with empty data permanently.
4. **`TOKEN_MANAGER_ROLE`:** copy each token's legacy configuration into inline records for every remote chain, in this
   order:

   | Type | Values to copy or set |
   | --- | --- |
   | Remote pool | Copy `getRemotePools`, with `allowLegacyPoolDataOnReceive` enabled for each legacy pool. For 1.5.0, read `getRemotePool` or wait for its upgrade. Enable `useLegacyPoolDataOnSend`. Add the future remote MTTP where the token will migrate. |
   | Rate limit | Copy the legacy limits. Pools older than 1.6.1 measure inbound amounts in source units; convert to local units if decimals differ. |
   | Fee | Add an inline record only where the token needs different fees from the default. Use fees from an enabled fee configuration of the legacy 2.0 pool, or fees chosen by the issuer. |
   | Security | For a legacy 2.0 pool with a hook, copy the hook and its `outboundCCVs` and `inboundCCVs`. |
   | Connection, last | Copy the remote token in the remote MTTP's source-token encoding. Use `abi.encode(address)` for EVM, converting any legacy 20-byte value. For non-EVM chains, preserve the exact bytes from `getRemoteToken`. This write runs the lane check. |

   Configure `allowLegacyPoolDataOnReceive` for the future remote MTTP according to the planned switch order:

   | Local chain | `allowLegacyPoolDataOnReceive` for the remote MTTP |
   | --- | --- |
   | Switches first on this lane | Disabled |
   | Switches second on this lane | Enabled |

5. **Role admin of each token:** grant the MTTP mint and burn roles. Before the switch, the PCM prevents their use
   through transfers because it accepts validation calls only from the token's TAR pool.
6. **Hook owner:** authorize the MTTP on any reused legacy hook.
7. **Legacy pool owner:** before each remote switch, call
   `addRemotePool(remoteChainSelector, abi.encode(remoteMttp))` on every legacy pool whose token will use that MTTP.
   Keep the old remote pools.
   - Do not add an MTTP that will never serve this token. A legacy pool cannot distinguish tokens sent by a trusted
     MTTP.
   - For a chain with a 1.5.0 pool, update the replacement pool from the upgrade procedure below.
   - Non-EVM pools compare exact source pool bytes. The v2 OnRamp encodes an EVM source pool as 32 padded bytes.
     Add that form without removing existing trusted pool entries.

#### Migration phase 2: Switch pools

Start only after phase 1 is complete on every chain.

8. **Token admin:** call `TAR.setPool(token, mttp)` to switch the token.
   - Start with one token on one chain and test both directions.
   - Switch the other tokens on that chain in one batch, then proceed to the next chain.
   - Any chain order can work without a pause because both sides trust old and new pools. The order must match
     step 4's `allowLegacyPoolDataOnReceive` values. If it changes, **`TOKEN_MANAGER_ROLE`** on both chains must first
     update those values. Tooling then repeats the pre-switch checks before the switch.

   For each lane where this chain switches second, include these actions in its switch batch:

   1. **`TOKEN_MANAGER_ROLE`:** set an inline destination overhead record toward the first chain. Use 192 bytes plus padded
      provider data, and the `destGasOverhead` from step 2.
   2. **`TOKEN_MANAGER_ROLE`:** disable `useLegacyPoolDataOnSend` for that lane.

   The first chain already accepts long data. This ordering prevents legacy-format MTTP messages from the second
   chain, so the first chain never needs a receive exception for that MTTP.

   Raise `destBytesOverhead` before disabling `useLegacyPoolDataOnSend`, or sends revert before the burn. The same
   requirement applies later at step 10.

   Do not switch both ends of a lane close together:

   - If both keep `useLegacyPoolDataOnSend` enabled, both accept a period of legacy-format MTTP traffic.
   - If both disable it, a message arriving before the other switch fails until that switch and needs a retry.

9. Keep the legacy pools for rollback. They must accept messages from each remote token's current pool. Whenever a
   remote chain selects a new pool, the **owner of each retained legacy pool** calls
   `addRemotePool(remoteChainSelector, abi.encode(newPool))`.

#### Upgrade pools older than 1.5.1

A 1.5.0 pool trusts only one remote pool per chain. `setRemotePool` replaces it immediately, so it cannot trust both the
old pool and MTTP during migration. Upgrade these pools on every chain before migration phase 2.

A 1.5.0 pool sends no decimals and ignores received decimals. Its old messages can execute after the upgrade, so both
ends of these lanes must continue to use the same decimals.

For a 1.5.0 pool on chain C:

1. **Legacy pool owner on C:** deploy a pool of version 1.5.1 or later. For each remote chain, trust the pools listed by
   the old pool. Also trust that chain's replacement pool if it is upgrading.
2. **Legacy pool owner on each remote chain already using 1.5.1 or later:** call
   `addRemotePool(chainC, abi.encode(newPool))` before C's TAR switch.
3. Upgrade any remote 1.5.0 pool too. Its replacement must trust both C's old and new pools.
4. **Role admin of the token on C:** grant the new pool mint and burn roles before the TAR switch.
5. **Token admin on C:** call `TAR.setPool(token, newPool)`.
   A new-pool message can arrive before the remote upgrade switches and fail on that chain's 1.5.0 pool. Tooling
   must list these messages and retry them through `OffRamp.execute` after the remote switch.
6. **Role admin of the token on C:** revoke the 1.5.0 pool's mint and burn roles after the switch. Do not keep it for
   rollback.
   The OffRamp uses the TAR pool at execution time, so it no longer executes on that old pool.
7. **`TOKEN_MANAGER_ROLE` on remote chains with phase 1 records already configured:** add C's new pool to the token's inline
   remote pool record with `allowLegacyPoolDataOnReceive` enabled. Keep the old pool permanently; its messages never expire.

#### Migration phase 3: Finish each lane

Complete this phase separately for each token and lane. Start only when both TARs select the token's MTTP and the
issuer has ruled out rollback on both chains. Every step below applies to that token and lane.

- **F** is the chain that switched first on this lane.
- **S** is the chain that switched second and disabled `useLegacyPoolDataOnSend` for this lane during the switch.

10. **`TOKEN_MANAGER_ROLE` on F:** disable `useLegacyPoolDataOnSend` toward S, preserving every pool and every other flag.
    In the same batch, first set the inline destination overhead record as in step 8. New messages from F now include
    the source token and additional pool data.
11. Wait until all legacy-format messages that F's MTTP sent for this token on this lane have executed on S.
    This covers every such message from F's switch through step 10, including failed messages. Check destination
    execution state as described in the [workflow rules](#message-execution-and-retries).
    Legacy-pool messages do not need to finish before this step; their receive permission remains enabled.
12. **`TOKEN_MANAGER_ROLE` on S:** disable `allowLegacyPoolDataOnReceive` for F's MTTP. Keep every legacy pool with
    `allowLegacyPoolDataOnReceive` enabled permanently. F needs no equivalent change because S's MTTP never had that
    permission on F.
13. **Role admin of each token:** once rollback is ruled out, revoke the legacy pool's mint and burn roles. The owner
    of a legacy 2.0 pool also withdraws its fees.

After step 12, retain inline remote pool records for chains with legacy pool history. See
[moving a migrated token to policies](#move-a-migrated-token-to-policies).

#### Messages in flight during migration

| Message | Required handling |
| --- | --- |
| Legacy source pool; destination has switched | The MTTP accepts it while the legacy source remains listed with `allowLegacyPoolDataOnReceive` enabled. Retain this permanently, including for retries years later. |
| MTTP source; destination has not switched | Send legacy data and ensure the destination legacy pool trusts the MTTP. Long data fails until the destination switches. |
| MTTP source with legacy data; destination has switched | The destination must have `allowLegacyPoolDataOnReceive` enabled for the source MTTP. Step 12 waits for these messages. |
| MTTP source with long data | The destination checks source identity. If configured, its additional pool data provider also checks the additional data, including empty data. |
| Failed message awaiting retry | The TAR pool at retry time executes it under the configuration then in effect. |

S disables `useLegacyPoolDataOnSend` at its switch. F keeps it enabled until step 10, after rollback is ruled out. If a
legacy-format MTTP message unexpectedly fails after step 12, **`TOKEN_MANAGER_ROLE`** re-enables
`allowLegacyPoolDataOnReceive` for that MTTP. Any party retries the message. **`TOKEN_MANAGER_ROLE`** then disables it again.

Other requirements for pending messages:

- **Additional data:** legacy pools, legacy-format MTTP sends, and source tokens without providers supply empty data.
  A migrated token can have no additional pool data provider. If configured, the provider must accept empty data from
  each relevant chain. A rejecting provider blocks retries until `TOKEN_MANAGER_ROLE` replaces or removes it.
- **Gas:** messages sent before the switch may have been priced for the legacy pool's lower receive cost. Execution
  must have enough transaction gas for the MTTP. If execution fails, tooling must list the message and retry through
  `OffRamp.execute` with enough gas. The source quote does not cap the pool call's gas.
- **CCVs:** the v2 OffRamp reads the destination TAR pool's inbound CCVs at execution time. Preserve the legacy pool's
  requirements in the security record, regardless of source pool. Otherwise pending messages can fail until the
  required list is restored.
- **OffRamps:** retain each OffRamp with pending messages on the MTTP's Router. A message cannot move to another
  OffRamp.
- **Rate limits:** legacy bucket balances do not migrate. MTTP buckets start full, so the switch can allow an additional
  amount equal to each bucket's capacity. Messages that exceed the available inbound limit fail and can be retried
  after refill.

#### Checks before migration step 8

For every token and remote chain, check that these views match steps 2–4:

- `getLaneRecordLevelBatch`
- `getCurrentRateLimiterStateBatch`
- `getFeeBatch`
- `getLaneVerificationBatch`
- `getTokenConfigBatch`

Then verify the following:

| Source | Expected result |
| --- | --- |
| `getLaneConnectionBatch` | Remote token matches step 4. `useLegacyPoolDataOnSend` is enabled. Legacy pools have `allowLegacyPoolDataOnReceive` enabled. The remote MTTP's `allowLegacyPoolDataOnReceive` is enabled only on the chain switching second. |
| EVM remote token | 32 bytes with 12 leading zero bytes, representing the legacy `getRemoteToken` address |
| Non-EVM remote token | Exact bytes from legacy `getRemoteToken` |
| Each remote legacy pool | `isRemotePool(thisChain, abi.encode(mttp))` is true, using the required encoding for that lane. |
| Each lane to a non-MTTP | `useLegacyPoolDataOnSend` enabled, remote pool level `Inline`, and `destBytesOverhead` equal to 32 |
| Source fee configurations | Gas estimates cover measured receive cost with a margin, including legacy sources priced through the pool or FeeQuoter. |
| MTTP Router | Includes every OffRamp with pending messages for the token's remote chains. |
| `getLaneVerificationBatch` | Matches the legacy pool's `getRequiredCCVs`. For a pool without `IPoolV2` support, empty lists select default CCVs. |
| Additional pool data provider | Not configured, or accepts empty data from every relevant remote chain |
| Second chain's switch batch | Sets inline overhead toward the first chain to 192 bytes plus padded provider data. |
| TARs and old 1.5.0 pools | No TAR selects a 1.5.0 pool for the token; replaced 1.5.0 pools have no mint or burn role. |
| Each remote legacy pool after a 1.5.0 upgrade | `isRemotePool(thisChain, abi.encode(newPool))` is true for each replacement pool. |

#### Checks for migration phase 3

| When | Check |
| --- | --- |
| Before step 10 | The remote TAR selects the MTTP for this token. |
| Before step 12 | Every legacy-format message from the remote MTTP has executed locally. |
| After step 12 | `getLaneConnectionBatch` shows `allowLegacyPoolDataOnReceive` disabled for the remote MTTP and enabled for every remote legacy pool. |

#### Migration rollback

Before rolling back, tooling must read each legacy pool's `getRemotePools` for every remote chain. Reject the rollback
if any of the following holds:

- The legacy pool lists a pool named in an incident.
- The legacy pool does not trust the token's current remote TAR pool. See migration step 9.
- A remote TAR still selects a pool named in an incident. Wait for that chain's recovery.
- The legacy pool's inbound CCVs differ from MTTP `getLaneVerificationBatch`. A legacy pool without `IPoolV2` support
  requires the default CCVs.

Choose the rollback procedure separately for each remote chain. A TAR switch alone is sufficient only if every remote
chain still sends legacy data to this chain. For a lane between F and S, this holds for S before step 10 on F, or for
F before S switches. The **token admin** can then call `TAR.setPool(token, legacyPool)`.

Otherwise, use this order:

1. **`TOKEN_MANAGER_ROLE` locally:** enable `allowLegacyPoolDataOnReceive` for the remote MTTP. The local MTTP must accept
   the legacy data introduced in rollback step 2 until the TAR switches.
2. **`TOKEN_MANAGER_ROLE` remotely:** enable `useLegacyPoolDataOnSend` toward this chain for the token. In the same batch, set
   an inline destination overhead record with `destBytesOverhead` equal to 32. The eventual legacy destination needs
   this format and byte limit.
3. Wait for all long-format messages from that remote MTTP to execute locally. Check execution state as in migration
   step 11.
4. **Role admin of the token:** if migration step 13 revoked the legacy pool's roles, restore its mint and burn roles.
   Otherwise, its mint calls will fail.
5. **Token admin:** call `TAR.setPool(token, legacyPool)`. Pending messages now execute on the legacy pool, which trusts
   the remote MTTPs.
6. After returning to the MTTP, **`TOKEN_MANAGER_ROLE` remotely:** first raise the byte allowance, then disable
   `useLegacyPoolDataOnSend`, as in migration step 10. Locally, drain legacy-format MTTP messages and disable
   `allowLegacyPoolDataOnReceive` from rollback step 1, as in migration steps 11 and 12.

If the rollback cannot wait for step 3, the token admin can switch earlier. Pending long-format messages then wait on
the legacy pool until the token returns to an MTTP. If it never returns, those messages remain stranded.

During rollback, the legacy pool again trusts the remote MTTP through migration step 7. This again permits legacy pool
data from a remote MTTP until rollback step 6 completes.

### Move a migrated token to policies

After migration phase 3, a token can join policies while keeping its inline records. Inline records still take
precedence. Removing one immediately selects the policy record for that chain.

#### Policy adoption steps

For one chain, follow this order. Each step can be a separate call.

1. **`TOKEN_MANAGER_ROLE`:** assign the connection policy with `applyTokenConnectionPolicyUpdates`.
   - Before assignment, tooling must check every chain where the policy will serve the token without an inline
     connection. Assignment opens those lanes without a lane check.
   - Where `onlySubsetTokens` is enabled, **`CONNECTION_ROLE`** must add the token with `applySubsetTokenUpdates`.
     This runs the lane check. Without membership, step 2 closes that lane.
2. **`TOKEN_MANAGER_ROLE`:** remove the inline connection.
   - The store rejects removal if the policy resolves different remote token bytes.
   - Tooling must reject removal if `getLaneConnectionBatch` would report `hasConnection` false afterward. Keep the
     lane open.
3. **`TOKEN_MANAGER_ROLE`:** assign the rate limit, fee, destination overhead, and security policies. Then remove those
   inline records. Before each removal, **tooling** checks that the selected policy record has the required values
   for that token and chain.
   - Removing an inline rate limit record deletes its buckets. The policy buckets start full, so the move can allow
     an additional amount equal to each bucket's capacity.
   - For a remote chain still using a non-MTTP pool, retain an inline destination overhead record with 32 bytes. If
     none exists, set it before assigning the destination overhead policy.
4. **`TOKEN_MANAGER_ROLE`:** assign the remote pool policy with `applyTokenRemotePoolPolicyUpdates`. Before this step,
   **tooling** checks that its chain record includes every remote MTTP the token needs.
5. **`TOKEN_MANAGER_ROLE`:** handle the inline remote pool record according to the remote token's history:

   | Remote chain | Action |
   | --- | --- |
   | The token never used a legacy pool, and now uses a policy pool. | Remove the inline remote pool record. |
   | The token has legacy pool history. | Keep an inline record listing its legacy pools with `allowLegacyPoolDataOnReceive` enabled. Include the remote MTTP with `allowLegacyPoolDataOnReceive` enabled only while its legacy-format messages are pending. |

   For an inline record retained in this step, preserve the token's `useLegacyPoolDataOnSend` value. If the remote token
   uses a pool in the policy record, enable `trustPolicyRemotePools`. Otherwise, keep it disabled.
6. **`TOKEN_MANAGER_ROLE`:** if an inline record still lists the remote MTTP, remove that entry after every legacy-format
   message from it has executed.

Before enabling `trustPolicyRemotePools`, tooling must confirm that the assigned policy has a record for the chain.
Policy changes then affect both tokens using the policy record directly and tokens with opted-in inline records.
`getLaneRecordLevelBatch` reports `Inline` for the latter; find their setting through `getInlineRemotePoolConfigs`.

Assign policies before removing overrides. Removing first can leave the token without a required record and fail
transfers until assignment completes.

#### Records retained after adoption

For a remote chain with legacy pool history:

- Keep the token's inline remote pool record permanently. It lists only that token's pools. Add new remote MTTPs to
  the policy record only, where they take effect immediately through policy trust.
- Keep the per-token `useLegacyPoolDataOnSend` value. If the remote token rolls back to a legacy pool, enable it for this
  token only.
- If a late legacy-format MTTP message fails after step 6, add that MTTP back to the inline list with
  `allowLegacyPoolDataOnReceive` enabled. Any party can retry the message. Then remove the MTTP from the inline list again.
- Never put a legacy pool in a policy record.

For a remote token still using a non-MTTP pool, keep its inline remote pool record with `useLegacyPoolDataOnSend` enabled.
Do not replace it with a policy record. See [lanes to non-MTTP pools](#lanes-to-non-mttp-pools).

### Incident: a compromised remote pool

#### Starting state and goal

A pool P on chain X is compromised. For example, an attacker with its owner key can select a hostile PCM and Router.
P can then send messages without burning, with arbitrary pool data, for any token whose TAR pool is P.

CCVs verify these messages because the legitimate OnRamp sent them. On this chain, P can be trusted by policy records,
inline records, and legacy pools.

**Goal:** prevent every message from P from minting here while allowing messages from other pools on X.

#### Contain the incident on this chain

Every unpause must wait for all checks in step 5.

1. **`PAUSER_ROLE`:** call `pauseChains([chainX], Both)`. This stops MTTP transfers with X immediately. It does not stop tokens
   whose TAR pool is legacy.
2. **Legacy pool owner:** remove P from every local legacy pool that trusts it, including pools retained for rollback
   under migration step 9. Do this even if migration step 13 revoked their roles; rollback can restore those roles.

   | Legacy pool | Action |
   | --- | --- |
   | Version 1.5.1 or later | Call `removeRemotePool(chainX, abi.encode(P))`. |
   | Version 1.5.0, another pool available on X | Replace P with `setRemotePool(chainX, abi.encode(otherPool))`. |
   | Version 1.5.0, no replacement available | Tooling first saves its remote token and rate limit configurations. The pool owner then removes X through `applyChainUpdates` with `allowed` false. |

   Each legacy pool needs its own call, possibly from a different owner. Its path remains open until that call.
3. **`REMOTE_POOL_ROLE`:** update each policy chain record for X that lists P to remove it. If P is its only pool,
   remove the record;
   policy records cannot be empty. Tokens relying on it then have no remote pool record, so sends revert and receives
   fail.
4. **`TOKEN_MANAGER_ROLE`:** update each inline remote pool record for X that lists P to remove it. Preserve all other
   pools and flags.
   - Disabling `allowLegacyPoolDataOnReceive` is insufficient. Long-format forged data can include the expected source token.
   - If P is the only listed pool, set an empty pool list. Preserve `useLegacyPoolDataOnSend` and
     `trustPolicyRemotePools`.
   - Empty inline records are valid and trust no pools of their own. Do not remove these records.
5. **Tooling:** complete all five scans below. P must appear nowhere:

   | Scan | Reads |
   | --- | --- |
   | Policy records | `getPolicyNames`, then `getRemotePoolPolicy` for every remote pool policy |
   | Inline records | `getAllConfiguredTokens`, then `getInlineRemotePoolConfigs` for every token |
   | Active legacy pools | TAR `getAllConfiguredTokens`, then `getPool` for each token administered by the issuer. Skip tokens without pools. For non-MTTP pools, read `getRemotePools(chainX)`, or `getRemotePool(chainX)` on 1.5.0. A reverted read fails the scan. |
   | Retained legacy pools | `getRemotePools(chainX)` on every pool retained under migration step 9. Tooling must keep this inventory from the migration. |
   | Resolved lanes | `getLaneConnectionBatch` for every token and X, including policy pools added through `trustPolicyRemotePools` |

   Also query X's TAR to list every token whose pool there is P.
6. **`UNPAUSER_ROLE`:** call `unpauseChains([chainX], Inbound)`. Messages from the remaining trusted legacy and MTTP pools can
   execute again. Keep outbound paused.
7. **`TOKEN_MANAGER_ROLE`:** for every token whose remote TAR pool is P, set an inline rate limit stop on both outbound
   finality classes for X. Then **`UNPAUSER_ROLE`** calls `unpauseChains([chainX], Outbound)`.
   These tokens must not send to P, where their destination release would execute.

   Tokens with a local legacy TAR pool do not read those stops. For each affected legacy token, its **legacy pool
   owner** must set every outbound rate limit configuration for X to the smallest enabled values the pool accepts.
   On a legacy 2.0 pool, this includes both default finality and FTF:

   | Pool version | Smallest enabled configuration | Effect |
   | --- | --- | --- |
   | 1.6.1 or later | Capacity 0, rate 0 | Stops sends |
   | Older than 1.6.1, including 1.6.0 | Capacity 2, rate 1 | Still permits 1 base unit per second |

   Older pools require a positive rate below capacity. If that remaining amount matters, remove X from the legacy
   pool instead:

   - **Tooling:** before removal, save the remote token, trusted pools, and rate limit configurations.
   - Removal deletes all remote pools for X and blocks inbound messages too.
   - Re-add X before retrying those messages; any party can retry afterward.

   Until these legacy limits or removals take effect, only an RMN curse of X stops those sends.

#### Contain the incident on chain X

8. **Role admin of each token served by P:** revoke P's mint and burn roles. Each token requires two calls. Until then,
   the attacker can mint locally without a message through a Router and PCM of its choice.

#### Recover with a new pool

9. Chain X deploys a new pool with a new owner and moves its tokens there. This chain then configures trust and removes
   the stops from step 7:

   - **`TOKEN_MANAGER_ROLE`:** update the inline records for X changed in step 4. Without policy trust, add the new pool.
     With `trustPolicyRemotePools`, add a new legacy pool; add a new MTTP only while it sends legacy data. Preserve
     every other flag except for the changes required below.
   - Enable `allowLegacyPoolDataOnReceive` for the new pool if it is legacy. For an MTTP, enable it only while it sends
     legacy data before migration step 10 on X. Disable it after draining those messages, as in migration step 12.
   - **`REMOTE_POOL_ROLE`:** add a new MTTP to the policy records for X.
   - **Legacy pool owner:** update every step 2 pool where X is still configured, including rollback pools. On 1.5.1 or
     later, call `addRemotePool(chainX, abi.encode(newPool))`. On 1.5.0, call
     `setRemotePool(chainX, abi.encode(newPool))`.
   - **Legacy pool owner:** for each pool where step 2 or 7 removed X, re-add it with `applyChainUpdates` instead.
     Restore the remote token and limits saved by tooling. Trust the new pool and previously saved pools except P.
     A 1.5.0 pool accepts only one pool, so use the new one.
   - If the replacement is not an MTTP, **`TOKEN_MANAGER_ROLE`** must set inline remote pool records with
     `useLegacyPoolDataOnSend` enabled before lifting stops. List the new pool with `allowLegacyPoolDataOnReceive`
     enabled. Tokens previously using policy records also need these inline records. In the same batch, set inline
     destination overhead to 32 bytes.
   - **`TOKEN_MANAGER_ROLE`:** restore the outbound limits from before step 7. Remove an inline rate limit record only if
     step 7 created it. **Legacy pool owner:** also restore the previous outbound limits on legacy pools.
   - Never add P back to any record or legacy pool, including rollback pools.

#### Incident limits and pending messages

- CCIP ops can curse X through RMN in one call, stopping MTTP and legacy paths. This affects every issuer's tokens on
  the lane.
- P may have honest messages pending from before the compromise. The destination cannot distinguish them from forged
  messages. They remain failed permanently; the issuer must handle them outside the contracts, for example with a
  refund.
- The connection store has no blocked-pool list. Removing P requires a write to every record that lists it. The chain
  pause stops MTTP transfers while these updates run.

## Out of scope

- Ramp migration and execution of messages sent through pre-v2 ramps.
- Non-EVM chains in connection policy records.
- Rate limits in USD or scaled by decimals.
- Lock and release MTTPs.
- Per-policy pauses.
- Fee-on-transfer and rebasing tokens.
- USDC with CCTP, and other tokens requiring a special bridge protocol.
- Token TAR registry modules, such as `RegistryModuleExternalAccessControl`.
