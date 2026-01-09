// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

library USDCSourcePoolDataCodec {
  /// @dev The flag used to indicate that the source pool data is coming from a chain that does not have CCTP Support,
  /// and so the lock release pool should be used. The BurnMintWithLockReleaseTokenPool uses this flag as its source pool
  /// data to indicate that the tokens should be released from the lock release pool rather than attempting to be minted
  /// through CCTP.
  /// @dev The preimage is bytes4(keccak256("NO_CCTP_USE_LOCK_RELEASE")).
  bytes4 public constant LOCK_RELEASE_FLAG = 0xfa7c07de;

  /// @dev The preimage is bytes4(keccak256("CCTP_V1"))
  bytes4 public constant CCTP_VERSION_1_TAG = 0xf3567d18;

  /// @dev There are two different tags for CCTP V2 to allow for CCIP V1.7 Compatibility which will enable fast transfers.
  /// Both tags will route to the same CCTP V2 pool, but will allow for pools to identify the type of transfer (slow or fast).

  /// @dev The preimage is bytes4(keccak256("CCTP_V2"))
  bytes4 public constant CCTP_VERSION_2_TAG = 0xb148ea5f;

  /// @dev The preimage is bytes4(keccak256("CCTP_V2_CCV"))
  bytes4 public constant CCTP_VERSION_2_CCV_TAG = 0x3047587c;
}
