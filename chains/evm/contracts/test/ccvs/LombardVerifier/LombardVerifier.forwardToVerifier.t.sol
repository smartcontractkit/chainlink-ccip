// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRouter} from "../../../interfaces/IRouter.sol";
import {IBridgeV3} from "../../../interfaces/lombard/IBridgeV3.sol";

import {LombardVerifier} from "../../../ccvs/LombardVerifier.sol";
import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {LombardVerifierSetup} from "./LombardVerifierSetup.t.sol";

import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

contract LombardVerifier_forwardToVerifier is LombardVerifierSetup {
  function setUp() public override {
    super.setUp();

    vm.stopPrank();
    vm.prank(s_onRamp);
  }

  function test_forwardToVerifier() public {
    address receiver = makeAddr("receiver");
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _createForwardMessage(address(s_testToken), receiver);

    vm.expectCall(
      address(s_mockBridge),
      abi.encodeCall(
        IBridgeV3.deposit,
        (
          LOMBARD_CHAIN_ID,
          address(s_testToken),
          OWNER,
          // This ensures the receiver is correctly encoded from unpadded bytes to bytes32.
          bytes32(uint256(uint160(receiver))),
          TRANSFER_AMOUNT,
          ALLOWED_CALLER,
          abi.encodePacked(bytes4(keccak256("LombardVerifier 1.7.0")), messageId)
        )
      )
    );

    bytes memory verifierData = s_lombardVerifier.forwardToVerifier(message, messageId, address(0), 0, "");

    // Verify it returns encoded payload hash.
    bytes32 payloadHash = abi.decode(verifierData, (bytes32));
    assertEq(payloadHash, s_mockBridge.s_lastPayloadHash());
  }

  function test_forwardToVerifier_WithAdapter() public {
    // Add a token with an adapter.
    address tokenWithAdapter = address(new BurnMintERC20("Token With Adapter", "TWA", 18, 0, 0));
    address adapter = address(new BurnMintERC20("Adapter", "ADP", 18, 0, 0));

    vm.stopPrank();
    vm.startPrank(OWNER);

    LombardVerifier.SupportedTokenArgs[] memory tokensToAdd = new LombardVerifier.SupportedTokenArgs[](1);
    tokensToAdd[0] = LombardVerifier.SupportedTokenArgs({localToken: tokenWithAdapter, localAdapter: adapter});
    s_lombardVerifier.updateSupportedTokens(new address[](0), tokensToAdd);

    address receiver = makeAddr("receiver");
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _createForwardMessage(tokenWithAdapter, receiver);

    vm.stopPrank();
    vm.startPrank(s_onRamp);

    // Should succeed - the adapter is used for the bridge deposit.
    bytes memory verifierData = s_lombardVerifier.forwardToVerifier(message, messageId, address(0), 0, "");

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
      _createForwardMessage(address(s_testToken), address(0));
    message.tokenTransfer[0].tokenReceiver = tooLongReceiver;

    vm.expectRevert(abi.encodeWithSelector(LombardVerifier.InvalidReceiver.selector, tooLongReceiver));
    s_lombardVerifier.forwardToVerifier(message, messageId, address(0), 0, "");
  }

  function test_forwardToVerifier_RevertWhen_TokenNotSupported() public {
    address unsupportedToken = makeAddr("unsupportedToken");
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) =
      _createForwardMessage(unsupportedToken, makeAddr("receiver"));

    vm.expectRevert(abi.encodeWithSelector(LombardVerifier.TokenNotSupported.selector, unsupportedToken));
    s_lombardVerifier.forwardToVerifier(message, messageId, address(0), 0, "");
  }

  function test_forwardToVerifier_RevertWhen_PathNotExist() public {
    // Use a chain selector that doesn't have a path configured.
    uint64 unknownChainSelector = 20000;

    vm.stopPrank();
    vm.startPrank(OWNER);

    // Set the owner as allowed onRamp.
    vm.mockCall(address(s_router), abi.encodeCall(IRouter.getOnRamp, (unknownChainSelector)), abi.encode(OWNER));

    BaseVerifier.RemoteChainConfigArgs[] memory remoteChainConfigArgs = new BaseVerifier.RemoteChainConfigArgs[](1);
    remoteChainConfigArgs[0] = _getRemoteChainConfig(s_router, unknownChainSelector, false);

    s_lombardVerifier.applyRemoteChainConfigUpdates(remoteChainConfigArgs);

    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) =
      _createForwardMessage(address(s_testToken), makeAddr("receiver"));
    message.destChainSelector = unknownChainSelector;

    vm.expectRevert(abi.encodeWithSelector(LombardVerifier.PathNotExist.selector, unknownChainSelector));
    s_lombardVerifier.forwardToVerifier(message, messageId, address(0), 0, "");
  }

  function test_forwardToVerifier_RevertWhen_RemoteChainNotSupported() public {
    uint64 unknownChainSelector = 999999;

    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) =
      _createForwardMessage(address(s_testToken), makeAddr("receiver"));
    message.destChainSelector = unknownChainSelector;

    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.RemoteChainNotSupported.selector, unknownChainSelector));
    s_lombardVerifier.forwardToVerifier(message, messageId, address(0), 0, "");
  }

  function test_forwardToVerifier_RevertWhen_CallerIsNotARampOnRouter() public {
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) =
      _createForwardMessage(address(s_testToken), makeAddr("receiver"));

    address randomCaller = makeAddr("randomCaller");

    vm.stopPrank();
    vm.startPrank(randomCaller);

    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.CallerIsNotARampOnRouter.selector, randomCaller));
    s_lombardVerifier.forwardToVerifier(message, messageId, address(0), 0, "");
  }

  function test_forwardToVerifier_RevertWhen_CursedByRMN() public {
    address receiver = makeAddr("receiver");
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _createForwardMessage(address(s_testToken), receiver);

    _setMockRMNChainCurse(message.destChainSelector, true);

    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.CursedByRMN.selector, message.destChainSelector));
    s_lombardVerifier.forwardToVerifier(message, messageId, address(0), 0, "");
  }
}
