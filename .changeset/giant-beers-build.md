---
"@chainlink/crib": major
---

Changed all ingresses from class `alb` to Infra-Platform provided `nginx-internal`, for cost savings purposes. Versioned as major as it's a somewhat big change in the underlying infra of every CRIB setup, even though every component has been tested and proven to work with the new setup in a backwards-compatible way.
