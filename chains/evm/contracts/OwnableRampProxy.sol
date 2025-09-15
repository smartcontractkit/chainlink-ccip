// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RampProxy} from "./RampProxy.sol";

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

/// @notice OwnableRampProxy is a RampProxy that uses Ownable2StepMsgSender for ownership management.
contract OwnableRampProxy is RampProxy, Ownable2StepMsgSender {
  /// @notice Sets the ramp address for a given remote chain selector and version.
  /// @dev Can be used to remove a ramp by setting the address to 0.
  /// @param ramps The array of ramps to set.
  function setRamps(
    SetRampsArgs[] calldata ramps
  ) external onlyOwner {
    _setRamps(ramps);
  }
}
