// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RampProxy} from "./RampProxy.sol";

import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

/// @notice OwnableRampProxy is a RampProxy that uses Ownable2StepMsgSender for ownership management.
contract OwnableRampProxy is RampProxy, Ownable2StepMsgSender {
  constructor(
    address rampAddress
  ) RampProxy(rampAddress) {}

  /// @notice Sets the ramp address.
  /// @param rampAddress the ramp address to set.
  function setRamp(
    address rampAddress
  ) external onlyOwner {
    _setRamp(rampAddress);
  }
}
