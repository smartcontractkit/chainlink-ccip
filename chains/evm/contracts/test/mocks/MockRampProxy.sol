// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RampProxy} from "../../RampProxy.sol";

contract MockRampProxy is RampProxy {
  function setRamps(
    SetRampsArgs[] calldata ramps
  ) external {
    _setRamps(ramps);
  }
}
