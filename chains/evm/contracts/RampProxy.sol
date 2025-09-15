// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

/// @notice RampProxy enables upgrades to CCVOnRamps and CCVOffRamps without breaking existing references in token pools, receivers, and apps.
/// The address of this contract will be referenced in the following places:
///   - Users / apps will specify required and optional CCVs as part of ccipSend extraArgs.
///   - Token pools will specify required and optional CCVs on both source and destination.
///   - Receiver contracts will specify required and optional CCVs on destination.
///   - CCVProxy will specify default and mandated CCVs for each destination.
///   - CCVAggregator will specify default and mandated CCVs for each source.
/// Each of these references should be to a RampProxy contract, not a CCVOnRamp or CCVOffRamp directly.
/// @dev On source, the CCVProxy will forward requests (i.e. getFee, forwardToVerifier) through this contract to the required CCVOnRamp based on the remote chain selector desired.
/// The same applies on destination. The CCVAggregator will forward requests (i.e. verifyMessage) through this contract to the required CCVOffRamp.
/// To support the proxy, future versions of ICCVOnRamp and ICCVOffRamp must maintain remoteChainSelector and caller as the first three parameters in every method.
abstract contract RampProxy is ITypeAndVersion {
  error InvalidRemoteChainSelector(uint64 remoteChainSelector);
  error InvalidRampAddress(address rampAddress);
  error RemoteChainNotSupported(uint64 remoteChainSelector);

  event RampUpdated(
    uint64 indexed remoteChainSelector, address indexed prevRampAddress, address indexed newRampAddress
  );

  struct SetRampsArgs {
    uint64 remoteChainSelector; // ─╮ The remote chain selector.
    address rampAddress; // ────────╯ The address of the ramp contract.
  }

  string public constant override typeAndVersion = "RampProxy 1.7.0-dev";

  /// @notice Each remote chain selector can support a different ramp contract.
  /// @dev If proxying a CCVOffRamp, one must be mindful when updating the rampAddress for a source chain.
  /// This is because there may be in-flight messages meant for the old CCVOffRamp at the time of the update.
  /// Ensure that the new CCVOffRamp you set is backwards-compatibile with prior ICCVOffRamp versions.
  mapping(uint64 remoteChainSelector => address rampAddress) private s_ramps;

  /// @notice Sets the ramp address for a given remote chain selector.
  /// @dev Can be used to remove a ramp by setting the address to 0.
  /// @param ramps The array of ramps to set.
  function _setRamps(
    SetRampsArgs[] calldata ramps
  ) internal {
    for (uint256 i = 0; i < ramps.length; ++i) {
      SetRampsArgs memory ramp = ramps[i];
      // Ramp address can be zero to remove a ramp, but if non-zero it must be a contract.
      if (ramp.rampAddress != address(0) && ramp.rampAddress.code.length == 0) {
        revert InvalidRampAddress(ramp.rampAddress);
      }
      if (ramp.remoteChainSelector == 0) revert InvalidRemoteChainSelector(ramp.remoteChainSelector);
      address prevRampAddress = s_ramps[ramp.remoteChainSelector];
      s_ramps[ramp.remoteChainSelector] = ramp.rampAddress;
      emit RampUpdated(ramp.remoteChainSelector, prevRampAddress, ramp.rampAddress);
    }
  }

  /// @notice Gets the ramp address for a given remote chain selector.
  /// @param remoteChainSelector The remote chain selector.
  /// @return rampAddress The address of the ramp.
  function getRamp(
    uint64 remoteChainSelector
  ) external view returns (address) {
    return s_ramps[remoteChainSelector];
  }

  // The fallback function forwards all calls to the appropriate ramp contract based on the remote chain selector.
  // solhint-disable-next-line payable-fallback, no-complex-fallback
  fallback() external {
    uint64 remoteChainSelector;

    assembly {
      remoteChainSelector := calldataload(4)
    }

    address rampAddress = s_ramps[remoteChainSelector];
    if (rampAddress == address(0)) revert RemoteChainNotSupported(remoteChainSelector);

    assembly {
      // We never cede control back to Solidity, so we can overwrite memory starting from index 0.
      calldatacopy(0, 0, calldatasize())
      // Overwrite calldata with the actual caller.
      // This prevents an attacker from spoofing a different caller.
      // The caller is at calldata index 4 (function selector) + 32 (remoteChainSelector) = 36.
      mstore(36, caller())

      // Forward the call to the ramp contract.
      let success := call(gas(), rampAddress, 0, 0, calldatasize(), 0, 0)
      returndatacopy(0, 0, returndatasize())
      if success { return(0, returndatasize()) }
      revert(0, returndatasize())
    }
  }
}
