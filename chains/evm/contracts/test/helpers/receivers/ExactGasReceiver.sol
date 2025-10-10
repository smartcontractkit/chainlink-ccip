// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAny2EVMMessageReceiver} from "../../../interfaces/IAny2EVMMessageReceiver.sol";
import {Client} from "../../../libraries/Client.sol";
import {IERC165} from "@openzeppelin/contracts@4.8.3/utils/introspection/IERC165.sol";

contract ExactGasReceiver is IAny2EVMMessageReceiver, IERC165 {
  uint256 internal s_gasToConsume;

  constructor(
    uint256 gasToConsume
  ) {
    s_gasToConsume = gasToConsume;
  }

  // From IAny2EVMMessageReceiver
  function ccipReceive(
    Client.Any2EVMMessage calldata /* message */
  ) external {
    uint256 targetGasLeft = gasleft() - s_gasToConsume;
    while (gasleft() > targetGasLeft) {}
  }

  function supportsInterface(
    bytes4 interfaceId
  ) external pure override returns (bool) {
    return interfaceId == type(IAny2EVMMessageReceiver).interfaceId || interfaceId == type(IERC165).interfaceId;
  }
}
