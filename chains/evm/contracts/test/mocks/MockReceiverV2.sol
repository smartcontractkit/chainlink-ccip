// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAny2EVMMessageReceiver} from "../../interfaces/IAny2EVMMessageReceiver.sol";
import {IAny2EVMMessageReceiverV2} from "../../interfaces/IAny2EVMMessageReceiverV2.sol";
import {Client} from "../../libraries/Client.sol";

contract MockReceiverV2 is IAny2EVMMessageReceiverV2 {
  address[] internal s_required;
  address[] internal s_optional;
  uint8 internal s_threshold;

  constructor(address[] memory required, address[] memory optional, uint8 threshold) {
    s_required = required;
    s_optional = optional;
    s_threshold = threshold;
  }

  // From IAny2EVMMessageReceiver
  function ccipReceive(
    Client.Any2EVMMessage calldata /* message */
  ) external {}

  // From IAny2EVMMessageReceiverV2
  function getCCVs(
    uint64 /* sourceChainSelector */
  ) external view returns (address[] memory, address[] memory, uint8) {
    return (s_required, s_optional, s_threshold);
  }
}
