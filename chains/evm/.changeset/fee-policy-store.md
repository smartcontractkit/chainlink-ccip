---
"@chainlink/contracts-ccip": minor
---

Add FeePolicyStore, which holds the fee type and the destination overhead type of multi-token pools. Update IPoolConfigManager for this split. LaneConnection no longer carries the destination overheads. TokenConfig renames chainPolicyName to remotePoolPolicyName and adds destOverheadPolicyName.
