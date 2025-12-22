// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAny2EVMMessageReceiver} from "../../interfaces/IAny2EVMMessageReceiver.sol";
import {IAny2EVMMessageReceiverV2} from "../../interfaces/IAny2EVMMessageReceiverV2.sol";
import {Client} from "../../libraries/Client.sol";
import {IERC165} from "@openzeppelin/contracts@4.8.3/utils/introspection/IERC165.sol";

contract MockReceiverV2 is IAny2EVMMessageReceiverV2, IERC165 {
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

  function supportsInterface(
    bytes4 interfaceId
  ) external pure override returns (bool) {
    return interfaceId == type(IAny2EVMMessageReceiverV2).interfaceId
      || interfaceId == type(IAny2EVMMessageReceiver).interfaceId || interfaceId == type(IERC165).interfaceId;
  }
}
