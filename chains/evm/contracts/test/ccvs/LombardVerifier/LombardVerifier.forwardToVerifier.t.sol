// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IBridgeV2} from "../../../interfaces/lombard/IBridgeV2.sol";

import {LombardVerifier} from "../../../ccvs/LombardVerifier.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {LombardVerifierSetup} from "./LombardVerifierSetup.t.sol";

import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

contract LombardVerifier_forwardToVerifier is LombardVerifierSetup {
  function test_forwardToVerifier() public {
    address receiver = makeAddr("receiver");
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) =
      _createForwardMessage(address(s_testToken), TRANSFER_AMOUNT, receiver);

    vm.expectCall(
      address(s_mockBridge),
      abi.encodeCall(
        IBridgeV2.deposit,
        (
          LOMBARD_CHAIN_ID,
          address(s_testToken),
          OWNER,
          // This ensures the receiver is correctly encoded from unpadded bytes to bytes32.
          bytes32(uint256(uint160(receiver))),
          TRANSFER_AMOUNT,
          ALLOWED_CALLER,
          abi.encode(messageId)
        )
      )
    );

    bytes memory verifierData = s_lombardVerifier.forwardToVerifier(message, messageId, address(0), 0, "");

    // Verify it returns encoded payload hash.
    bytes32 payloadHash = abi.decode(verifierData, (bytes32));
    assertEq(payloadHash, s_mockBridge.s_lastPayloadHash());
  }

  function test_forwardToVerifier_RevertWhen_MustTransferTokens() public {
    MessageV1Codec.MessageV1 memory message = _createBasicMessageV1(SOURCE_CHAIN_SELECTOR);

    vm.expectRevert(LombardVerifier.MustTransferTokens.selector);
    s_lombardVerifier.forwardToVerifier(message, bytes32(0), address(0), 0, "");
  }

  function test_forwardToVerifier_RevertWhen_InvalidReceiver() public {
    // Create a message with a receiver that's too long (> 32 bytes).
    bytes memory tooLongReceiver = new bytes(33);

    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) =
      _createForwardMessage(address(s_testToken), TRANSFER_AMOUNT, address(0));
    message.tokenTransfer[0].tokenReceiver = tooLongReceiver;

    vm.expectRevert(abi.encodeWithSelector(LombardVerifier.InvalidReceiver.selector, tooLongReceiver));
    s_lombardVerifier.forwardToVerifier(message, messageId, address(0), 0, "");
  }

  function test_forwardToVerifier_RevertWhen_TokenNotSupported() public {
    address unsupportedToken = makeAddr("unsupportedToken");
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) =
      _createForwardMessage(unsupportedToken, TRANSFER_AMOUNT, makeAddr("receiver"));

    vm.expectRevert(abi.encodeWithSelector(LombardVerifier.TokenNotSupported.selector, unsupportedToken));
    s_lombardVerifier.forwardToVerifier(message, messageId, address(0), 0, "");
  }

  function test_forwardToVerifier_RevertWhen_PathNotExist() public {
    // Use a chain selector that doesn't have a path configured.
    uint64 unknownChainSelector = 999999;

    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) =
      _createForwardMessage(address(s_testToken), TRANSFER_AMOUNT, makeAddr("receiver"));
    message.destChainSelector = unknownChainSelector;

    vm.expectRevert(abi.encodeWithSelector(LombardVerifier.PathNotExist.selector, unknownChainSelector));
    s_lombardVerifier.forwardToVerifier(message, messageId, address(0), 0, "");
  }

  function test_forwardToVerifier_WithAdapter() public {
    // Add a token with an adapter.
    address tokenWithAdapter = address(new BurnMintERC20("Token With Adapter", "TWA", 18, 0, 0));
    address adapter = address(new BurnMintERC20("Adapter", "ADP", 18, 0, 0));

    LombardVerifier.SupportedTokenArgs[] memory tokensToAdd = new LombardVerifier.SupportedTokenArgs[](1);
    tokensToAdd[0] = LombardVerifier.SupportedTokenArgs({localToken: tokenWithAdapter, localAdapter: adapter});
    s_lombardVerifier.updateSupportedTokens(new address[](0), tokensToAdd);

    address receiver = makeAddr("receiver");
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) =
      _createForwardMessage(tokenWithAdapter, TRANSFER_AMOUNT, receiver);

    // Should succeed - the adapter is used for the bridge deposit.
    bytes memory verifierData = s_lombardVerifier.forwardToVerifier(message, messageId, address(0), 0, "");

    bytes32 payloadHash = abi.decode(verifierData, (bytes32));
    assertEq(payloadHash, s_mockBridge.s_lastPayloadHash());
  }
}
