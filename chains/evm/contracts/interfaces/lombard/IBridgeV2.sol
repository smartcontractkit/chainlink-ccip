// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {IBridgeV1} from "./IBridgeV1.sol";

/// @custom:security-contact legal@lombard.finance
interface IBridgeV2 is IBridgeV1 {
  function deposit(
    bytes32 destinationChain,
    address token,
    address sender,
    bytes32 recipient,
    uint256 amount,
    bytes32 destinationCaller,
    // Optional bytes field that is forwarded to the destination chain and is included in the message proof.
    bytes calldata optionalMessage
  ) external payable returns (uint256, bytes32);
}
