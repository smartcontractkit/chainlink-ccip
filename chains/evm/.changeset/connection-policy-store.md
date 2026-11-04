---
"@chainlink/contracts-ccip": minor
---

Add ConnectionPolicyStore to manage token profiles, connections, trusted remote pools, and pauses for multi-token pools. Add SameAddressRemoteTokenAddressProvider for tokens with the same address across EVM chains. In IPoolConfigManager, rename omitSourceTokenAddressOnSend to useLegacyPoolDataOnSend. Replace the lane-wide allowSourceTokenAddressMissingOnReceive flag with allowLegacyPoolDataOnReceive for each remote pool.

Require fee configuration instead of falling back to FeeQuoter's token transfer fees. Check outbound pool data against the configured byte overhead on both the V1 and V2 paths.
