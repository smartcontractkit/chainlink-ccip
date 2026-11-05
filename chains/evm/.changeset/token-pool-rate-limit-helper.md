---
"@chainlink/contracts-ccip": minor
---

Add TokenPoolRateLimitHelper, a contract that can be set as a 2.0 token pool's `rateLimitAdmin` to update only the outbound or only the inbound rate limiter for a lane, re-writing the other direction's existing config.
