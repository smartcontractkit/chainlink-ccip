---
"@chainlink/contracts-ccip": minor
---

Add MultiTokenPoolHooks, the hooks contract that multi-token pools call on each transfer. It checks a sender allowlist and can run a policy engine. Add MultiTokenPoolHooksExternalBlocklist, which also checks the sender and the receiver of each transfer against an external IAddressBlocklist.
