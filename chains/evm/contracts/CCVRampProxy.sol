// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

/// @notice CCVRampProxy enables upgrades to CCVOnRamps and CCVOffRamps without breaking existing references in token pools, receivers, and apps.
/// @dev All future versions of ICCVOnRamp and ICCVOffRamp must maintain remoteChainSelector, version, and caller as the first three parameters in every method.
contract CCVRampProxy is Ownable2StepMsgSender, ITypeAndVersion {
  error InvalidRemoteChainSelector(uint64 remoteChainSelector);
  error InvalidRampAddress(address rampAddress);
  error InvalidVersion(bytes32 version);
  error UnknownRamp(uint64 remoteChainSelector, bytes32 version);

  event RampSet(uint64 indexed remoteChainSelector, bytes32 indexed version, address indexed rampAddress);

  struct SetRampsArgs {
    uint64 remoteChainSelector;
    address addr;
    bytes32 version;
  }

  string public constant override typeAndVersion = "CCVRampProxy 1.7.0-dev";

  /// @notice The supported ramps.
  /// @dev Each remote chain selector can have multiple ramps, each with a different version. This protects in-flight messages during upgrades.
  mapping(uint64 => mapping(bytes32 => address)) private s_ramps;

  /// @notice Sets the ramp address for a given remote chain selector and version.
  /// @dev Can be used to remove a ramp by setting the address to 0.
  /// @param ramps The array of ramps to set.
  function setRamps(
    SetRampsArgs[] calldata ramps
  ) external onlyOwner {
    for (uint256 i = 0; i < ramps.length; ++i) {
      SetRampsArgs memory ramp = ramps[i];
      if (ramp.remoteChainSelector == 0) revert InvalidRemoteChainSelector(ramp.remoteChainSelector);
      if (ramp.version == bytes32(0)) revert InvalidVersion(ramp.version);
      if (ramp.addr == address(0)) revert InvalidRampAddress(ramp.addr);
      s_ramps[ramp.remoteChainSelector][ramp.version] = ramp.addr;
      emit RampSet(ramp.remoteChainSelector, ramp.version, ramp.addr);
    }
  }

  /// @notice Gets the ramp address for a given remote chain selector and version.
  /// @param remoteChainSelector The remote chain selector.
  /// @param version The version of the ramp.
  /// @return rampAddress The address of the ramp.
  function getRamp(uint64 remoteChainSelector, bytes32 version) external view returns (address) {
    return s_ramps[remoteChainSelector][version];
  }

  // The fallback function forwards all calls to the appropriate ramp contract based on the remote chain selector and version.
  // solhint-disable-next-line payable-fallback, no-complex-fallback
  fallback() external {
    uint64 remoteChainSelector;
    bytes32 version;

    assembly {
      remoteChainSelector := calldataload(4)
      version := calldataload(36)
    }

    address rampAddress = s_ramps[remoteChainSelector][version];
    if (rampAddress == address(0)) revert UnknownRamp(remoteChainSelector, version);

    assembly {
      // We never cede control back to Solidity, so we can overwrite memory starting from index 0.
      calldatacopy(0, 0, calldatasize())
      // Overwrite calldata with the actual caller.
      // This prevents an attacker from spoofing a different caller.
      // The caller is at calldata index 4 (function selector) + 32 (remoteChainSelector) + 32 (version) = 68.
      mstore(68, caller())

      // Forward the call to the ramp contract.
      let success := call(gas(), rampAddress, 0, 0, calldatasize(), 0, 0)
      returndatacopy(0, 0, returndatasize())
      if success { return(0, returndatasize()) }
      revert(0, returndatasize())
    }
  }
}
