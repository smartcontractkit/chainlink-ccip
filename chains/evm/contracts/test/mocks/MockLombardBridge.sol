// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {IBridgeV3} from "../../interfaces/lombard/IBridgeV3.sol";
import {MockLombardMailbox} from "./MockLombardMailbox.sol";

contract MockLombardBridge is IBridgeV3 {
  address public s_mailbox;
  bytes32 public s_lastPayloadHash;

  mapping(bytes32 destinationChain => mapping(address sourceToken => bytes32 destinationToken)) internal
    s_allowedDestinationTokens;

  constructor() {
    s_mailbox = address(new MockLombardMailbox());
  }

  function mailbox() external view override returns (address) {
    return s_mailbox;
  }

  function setMailbox(
    address mailbox_
  ) external {
    s_mailbox = mailbox_;
  }

  // solhint-disable-next-line func-name-mixedcase
  function MSG_VERSION() external pure override returns (uint8) {
    return 1;
  }

  function deposit(
    bytes32,
    address,
    address,
    bytes32,
    uint256,
    bytes32,
    bytes calldata optionalMessage
  ) external payable override returns (uint256, bytes32) {
    s_lastPayloadHash = keccak256(abi.encode(block.timestamp, optionalMessage));

    MockLombardMailbox(s_mailbox).setMessageId(optionalMessage);

    return (0, s_lastPayloadHash);
  }

  function deposit(
    bytes32, // destinationChain
    address token,
    address, // sender
    bytes32, // recipient
    uint256, // amount
    bytes32 // destinationCaller
  ) external payable returns (uint256 nonce, bytes32 payloadHash) {
    return (1, keccak256(abi.encodePacked(block.timestamp, token)));
  }

  function getAllowedDestinationToken(bytes32 destinationChain, address sourceToken) external view returns (bytes32) {
    return s_allowedDestinationTokens[destinationChain][sourceToken];
  }

  function setAllowedDestinationToken(bytes32 destinationChain, address sourceToken, bytes32 destinationToken) external {
    s_allowedDestinationTokens[destinationChain][sourceToken] = destinationToken;
  }
}
