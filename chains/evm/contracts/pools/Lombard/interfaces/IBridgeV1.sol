// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

/// @custom:security-contact legal@lombard.finance
interface IBridgeV1 {
  event DestinationBridgeSet(bytes32 indexed destinationChain, bytes32 indexed destinationBridge);
  event DestinationTokenAdded(
    bytes32 indexed destinationChain, bytes32 indexed destinationToken, address indexed sourceToken
  );
  event DestinationTokenRemoved(
    bytes32 indexed destinationChain, bytes32 indexed destinationToken, address indexed sourceToken
  );
  event RateLimitsSet(address indexed token, bytes32 indexed sourceChainId, uint256 limit, uint256 window);

  event SenderConfigChanged(address indexed sender, uint32 feeDiscount, bool whitelisted);

  /// @notice Emitted when the is a deposit in the bridge
  event DepositToBridge(address indexed fromAddress, bytes32 indexed toAddress, bytes32 indexed payloadHash);

  /// @notice Emitted when a withdraw is made from the bridge
  event WithdrawFromBridge(address indexed recipient, bytes32 indexed chainId, address indexed token, uint256 amount);

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
