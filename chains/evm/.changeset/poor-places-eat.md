---
"@chainlink/contracts-ccip": minor
---

1. Creates a new USDCTokenPoolCCTPV2 which interacts with V2 of CCTP, USDCTokenPoolCCTPV2.sol

2. Modifies USDCTokenPool so that it can be both a standalone contract and a parent for the V2 Token Pool.

3. Modifies the relevant interfaces and mock functions for the CCTP Transmitter and token Messengers.

4. Replaces the HybridLockReleaseUSDCTokenPool with a SiloedUSDCTokenPool. The hybrid pool is being deprecated and the USDCBridgeMigrator functionality is also being moved into this singular contract as well. There are no additional breaking changes to the migrator contract,
only a removal of internal token accounting which is not inherited from the SiloedLockReleaseTokenPool.

5. The USDC Token Pools listed in this PR have overridden their `_onlyOffRamp` functions and instead have been modified to use an `AuthorizedCaller` library in the shared libraries folder. This is necessary to support the Proxy architecture which will be added in a separate PR.

6. To support the AuthorizedCaller library, the TokenPool error `ZeroAddressNotAllowed()` has been replaced with `ZeroAddressInvalid()` to avoid a naming collision not supported by the compiler.

7. Since the TokenPool.sol base has been modified, all relevant wrappers have been regenerated and bumped the type and version
to `1.6.3-dev`