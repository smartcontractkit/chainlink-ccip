// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {CCVRampProxy} from "./CCVRampProxy.sol";

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

/// @notice OwnableCCVRampProxy is a CCVRampProxy that uses Ownable2StepMsgSender for ownership management.
contract OwnableCCVRampProxy is CCVRampProxy, Ownable2StepMsgSender {
  /// @notice Sets the ramp address for a given remote chain selector and version.
  /// @dev Can be used to remove a ramp by setting the address to 0.
  /// @param ramps The array of ramps to set.
  function setRamps(
    SetRampsArgs[] calldata ramps
  ) external onlyOwner {
    _setRamps(ramps);
  }
}
