// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IAny2EVMMessageReceiver} from "./IAny2EVMMessageReceiver.sol";

interface IAny2EVMMessageReceiverV2 is IAny2EVMMessageReceiver {
  function getCCVs(
    uint64 sourceChainSelector
  ) external view returns (address[] memory requiredCCVs, address[] memory optionalCCVs, uint8 optionalThreshold);
}
