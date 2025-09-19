// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RampProxy} from "../../ccvs/RampProxy.sol";

contract MockRampProxy is RampProxy {
  constructor(
    address rampAddress
  ) RampProxy(rampAddress) {}

  function setRamp(
    address rampAddress
  ) external {
    _setRamp(rampAddress);
  }
}
