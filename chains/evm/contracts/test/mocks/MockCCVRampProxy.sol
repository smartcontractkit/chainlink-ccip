// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {CCVRampProxy} from "../../CCVRampProxy.sol";

contract MockCCVRampProxy is CCVRampProxy {
  function setRamps(
    SetRampsArgs[] calldata ramps
  ) external {
    _setRamps(ramps);
  }
}
