// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

/// @notice RampProxy enables upgrades to CCVRamps without breaking existing references in token pools, receivers, and apps.
/// The address of this contract will be referenced in the following places:
///   - Users / apps will specify required and optional CCVs as part of ccipSend extraArgs.
///   - Token pools will specify required and optional CCVs on both source and destination.
///   - Receiver contracts will specify required and optional CCVs on destination.
///   - CCVProxy will specify default and mandated CCVs for each destination.
///   - CCVAggregator will specify default and mandated CCVs for each source.
/// Each of these references should be to a RampProxy contract, not a CCVRamp directly.
/// @dev On source, the CCVProxy will forward requests (i.e. getFee, forwardToVerifier) through this contract to the required CCVRamp.
/// The same applies on destination. The CCVAggregator will forward requests (i.e. verifyMessage) through this contract to the required CCVRamp.
/// To support this proxy, all future CCVRamp interfaces must have originalCaller defined as the first arg to each method.
contract RampProxy {
  error ZeroAddressNotAllowed();

  event RampUpdated(address indexed oldRamp, address indexed newRamp);

  /// @notice The address of the ramp contract.
  address public s_ramp;

  constructor(
    address rampAddress
  ) {
    _setRamp(rampAddress);
  }

  function _setRamp(
    address rampAddress
  ) internal {
    if (rampAddress == address(0)) {
      revert ZeroAddressNotAllowed();
    }
    address oldRamp = s_ramp;
    s_ramp = rampAddress;
    emit RampUpdated(oldRamp, rampAddress);
  }

  // The fallback function forwards all calls to the ramp contract via delegatecall.
  // solhint-disable-next-line payable-fallback, no-complex-fallback
  fallback() external {
    address rampAddress = s_ramp;
    assembly {
      // We never cede control back to Solidity, so we can overwrite memory starting from index 0.
      calldatacopy(0, 0, calldatasize())
      // Overwrite calldata with the actual caller.
      // This prevents an attacker from spoofing a different caller.
      // The caller must be at calldata index 4 (skip function selector)
      mstore(4, caller())

      // Forward the call to the ramp contract.
      let success := call(gas(), rampAddress, 0, 0, calldatasize(), 0, 0)
      returndatacopy(0, 0, returndatasize())
      if success { return(0, returndatasize()) }
      revert(0, returndatasize())
    }
  }
}
