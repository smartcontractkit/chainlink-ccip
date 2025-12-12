// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

/// @custom:security-contact legal@lombard.finance
interface IBridgeV1 {
  function mailbox() external view returns (address);

  function MSG_VERSION() external view returns (uint8);

  function deposit(
    bytes32 destinationChain,
    address token,
    address sender,
    bytes32 recipient,
    uint256 amount,
    bytes32 destinationCaller
  ) external payable returns (uint256, bytes32);

  function getAllowedDestinationToken(bytes32 destinationChain, address sourceToken) external view returns (bytes32);
}
