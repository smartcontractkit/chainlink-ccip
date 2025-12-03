// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {IBridgeV2} from "../../pools/Lombard/interfaces/IBridgeV2.sol";

// solhint-disable func-name-mixedcase

contract MockLombardBridgeV2 is IBridgeV2 {
  struct DepositArgs {
    bytes32 destinationChain;
    address token;
    address sender;
    bytes32 recipient;
    uint256 amount;
    bytes32 destinationCaller;
    bytes payload;
  }

  uint8 internal immutable i_msgVersion;
  address internal s_mailbox;
  uint256 internal s_nextNonce = 1;

  DepositArgs public s_lastDeposit;
  mapping(bytes32 destinationChain => mapping(address sourceToken => bytes32 destinationToken)) internal
    s_allowedDestinationTokens;

  constructor(uint8 msgVersion, address mailbox_) {
    i_msgVersion = msgVersion;
    s_mailbox = mailbox_;
  }

  function setMailbox(
    address mailbox_
  ) external {
    s_mailbox = mailbox_;
  }

  function setAllowedDestinationToken(bytes32 destinationChain, address sourceToken, bytes32 destinationToken) external {
    s_allowedDestinationTokens[destinationChain][sourceToken] = destinationToken;
  }

  function MSG_VERSION() external view returns (uint8) {
    return i_msgVersion;
  }

  function mailbox() external view returns (address) {
    return s_mailbox;
  }

  function deposit(
    bytes32 destinationChain,
    address token,
    address sender,
    bytes32 recipient,
    uint256 amount,
    bytes32 destinationCaller,
    bytes calldata payload
  ) external payable returns (uint256 nonce, bytes32 payloadHash) {
    s_lastDeposit = DepositArgs({
      destinationChain: destinationChain,
      token: token,
      sender: sender,
      recipient: recipient,
      amount: amount,
      destinationCaller: destinationCaller,
      payload: payload
    });

    nonce = s_nextNonce++;
    payloadHash =
      keccak256(abi.encode(destinationChain, token, sender, recipient, amount, destinationCaller, payload, nonce));
    return (nonce, payloadHash);
  }

  function getAllowedDestinationToken(bytes32 destinationChain, address sourceToken) external view returns (bytes32) {
    return s_allowedDestinationTokens[destinationChain][sourceToken];
  }
}
