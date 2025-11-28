// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

interface IMailbox {
  error Mailbox_ZeroChainId();
  error Mailbox_ZeroConsortium();
  error Mailbox_ZeroMailbox();
  error Mailbox_ZeroRecipient();
  error Mailbox_ZeroAmount();
  error Mailbox_MessagePathEnabled(bytes32 id);
  error Mailbox_MessagePathDisabled(bytes32 id);
  error Mailbox_UnexpectedDestinationCaller(address expected, address actual);
  error Mailbox_HandlerNotImplemented();
  error Mailbox_PayloadOversize(uint32 max, uint256 actual);
  error Mailbox_NotEnoughFee(uint256 expected, uint256 actual);
  error Mailbox_CallFailed();

  event MessagePathEnabled(
    bytes32 indexed destinationChain,
    bytes32 indexed inboundMessagePath,
    bytes32 indexed outboundMessagePath,
    bytes32 destinationMailbox
  );

  event MessagePathDisabled(
    bytes32 indexed destinationChain,
    bytes32 indexed inboundMessagePath,
    bytes32 indexed outboundMessagePath,
    bytes32 destinationMailbox
  );

  event MessageSent(
    bytes32 indexed destinationLChainId, address indexed msgSender, bytes32 indexed recipient, bytes payload
  );

  /// Message payment receipt
  event MessagePaid(bytes32 indexed payloadHash, address indexed msgSender, uint256 payloadSize, uint256 value);

  event MessageDelivered(
    bytes32 indexed payloadHash, address indexed caller, uint256 indexed nonce, bytes32 msgSender, bytes payload
  );

  event MessageHandled(bytes32 indexed payloadHash, address indexed destinationCaller, bytes executionResult);

  event MessageHandleError(
    bytes32 indexed payloadHash, address indexed destinationCaller, string reason, bytes customError
  );

  event SenderConfigUpdated(address indexed sender, uint64 maxPayloadSize, bool feeDisabled);

  event DefaultPayloadSizeSet(uint64 maxPayloadSize);

  event FeePerByteSet(uint256 fee);

  event FeeWithdrawn(address indexed treasury, uint256 amount);

  function deliverAndHandle(
    bytes calldata rawPayload,
    bytes calldata proof
  ) external returns (bytes32, bool, bytes memory);
}
