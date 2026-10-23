---
"@chainlink/contracts-ccip": minor
---

Add RegistryModuleExternalAccessControl to register token admins in batches using getCCIPAdmin, owner, or an AccessControl role. The role can be DEFAULT_ADMIN_ROLE or CCIP_TOKEN_REGISTRY_ADMIN_ROLE, on the token or on an external AccessControl contract. Registration through CCIP_TOKEN_REGISTRY_ADMIN_ROLE requires its admin role to match DEFAULT_ADMIN_ROLE as defined by that AccessControl contract.
