// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {IBridgeV3} from "../../interfaces/lombard/IBridgeV3.sol";
import {MockLombardMailbox} from "./MockLombardMailbox.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

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
    return 2;
  }

  function deposit(
    bytes32,
    address token,
    address,
    bytes32,
    uint256 amount,
    bytes32,
    bytes calldata optionalMessage
  ) external payable override returns (uint256, bytes32) {
    IERC20(token).transferFrom(msg.sender, address(this), amount);
    s_lastPayloadHash = keccak256(abi.encode(block.timestamp, optionalMessage));

    MockLombardMailbox(s_mailbox).setMessageId(optionalMessage);

    return (0, s_lastPayloadHash);
  }

  function deposit(
    bytes32, // destinationChain
    address token,
    address, // sender
    bytes32, // recipient
    uint256 amount,
    bytes32 // destinationCaller
  ) external payable returns (uint256 nonce, bytes32 payloadHash) {
    IERC20(token).transferFrom(msg.sender, address(this), amount);
    return (1, keccak256(abi.encodePacked(block.timestamp, token)));
  }

  function getAllowedDestinationToken(
    bytes32 destinationChain,
    address sourceToken
  ) external view returns (bytes32) {
    return s_allowedDestinationTokens[destinationChain][sourceToken];
  }

  function setAllowedDestinationToken(
    bytes32 destinationChain,
    address sourceToken,
    bytes32 destinationToken
  ) external {
    s_allowedDestinationTokens[destinationChain][sourceToken] = destinationToken;
  }
}
