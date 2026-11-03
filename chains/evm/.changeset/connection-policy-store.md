---
"@chainlink/contracts-ccip": minor
---

Add ConnectionPolicyStore, which holds the token profiles, the pause state, the connection type and the remote pool type of multi-token pools. Add SameAddressRemoteTokenAddressProvider, which gives each token its own address as the remote token. Update IPoolConfigManager. LaneConnection renames omitSourceTokenAddressOnSend to useLegacyPoolDataOnSend. Each remote pool in LaneConnection has its own allowLegacyPoolDataOnReceive, which replaces allowSourceTokenAddressMissingOnReceive.
