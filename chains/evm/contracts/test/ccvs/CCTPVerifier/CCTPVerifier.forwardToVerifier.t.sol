// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITokenMessenger} from "../../../pools/USDC/interfaces/ITokenMessenger.sol";

import {CCTPVerifier} from "../../../ccvs/CCTPVerifier.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";

import {MockUSDCTokenMessenger} from "../../mocks/MockUSDCTokenMessenger.sol";
import {CCTPVerifierSetup} from "./CCTPVerifierSetup.t.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract CCTPVerifier_forwardToVerifier is CCTPVerifierSetup {
  uint256 internal s_transferAmount = 10e6; // 10 USDC

  function setUp() public override {
    super.setUp();

    // Send transfer amount to the verifier, mocking a transfer from the token pool.
    deal(address(s_USDCToken), address(s_cctpVerifier), s_transferAmount);
    assertEq(IERC20(address(s_USDCToken)).balanceOf(address(s_cctpVerifier)), s_transferAmount);

    // Grant mint and burn roles to the token messenger.
    BurnMintERC20(address(s_USDCToken)).grantMintAndBurnRoles(address(s_mockTokenMessenger));
  }

  function test_forwardToVerifier_MintRecipientFromMessage() public {
    bytes memory tokenReceiver = abi.encode(makeAddr("tokenReceiver"));
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) =
      _createMessage(0, address(s_USDCToken), tokenReceiver);

    bytes32 decodedReceiver = abi.decode(tokenReceiver, (bytes32));
    CCTPVerifier.Domain memory domain = s_cctpVerifier.getDomain(DEST_CHAIN_SELECTOR);

    vm.startPrank(s_onRamp);
    vm.expectEmit();
    emit ITokenMessenger.DepositForBurn(
      address(s_USDCToken),
      s_transferAmount,
      address(s_cctpVerifier),
      decodedReceiver,
      DEST_DOMAIN_IDENTIFIER,
      s_mockTokenMessenger.DESTINATION_TOKEN_MESSENGER(),
      domain.allowedCallerOnDest,
      0,
      CCTP_STANDARD_FINALITY_THRESHOLD,
      bytes.concat(s_cctpVerifier.versionTag(), messageId)
    );
    s_cctpVerifier.forwardToVerifier(message, messageId, s_sourceFeeTokens[0], 0, "");
  }

  function test_forwardToVerifier_MintRecipientFromDomain() public {}

  function test_forwardToVerifier_DefaultFinality() public {}

  function test_forwardToVerifier_OneCustomFinality() public {}

  function test_forwardToVerifier_MultipleCustomFinalities() public {}

  function test_forwardToVerifier_RevertWhen_CallerIsNotARampOnRouter() public {}

  function test_forwardToVerifier_RevertWhen_SenderIsNotAllowed() public {}

  function test_forwardToVerifier_RevertWhen_UnknownDomain() public {}

  function test_forwardToVerifier_RevertWhen_InvalidTokenTransferLength() public {}

  function test_forwardToVerifier_RevertWhen_InvalidToken() public {}

  function test_forwardToVerifier_RevertWhen_InvalidReceiver() public {}

  function test_forwardToVerifier_RevertWhen_UnsupportedFinality() public {}

  function test_forwardToVerifier_RevertWhen_MaxFeeExceedsUint32() public {}

  function _createMessage(
    uint16 finality,
    address sourceTokenAddress,
    bytes memory tokenReceiver
  ) internal returns (MessageV1Codec.MessageV1 memory, bytes32) {
    MessageV1Codec.TokenTransferV1[] memory tokenTransfer = new MessageV1Codec.TokenTransferV1[](1);
    tokenTransfer[0] = MessageV1Codec.TokenTransferV1({
      amount: s_transferAmount,
      sourcePoolAddress: abi.encodePacked(makeAddr("sourcePool")),
      sourceTokenAddress: abi.encodePacked(sourceTokenAddress),
      destTokenAddress: abi.encodePacked(makeAddr("destToken")),
      tokenReceiver: tokenReceiver,
      extraData: "extra data"
    });

    MessageV1Codec.MessageV1 memory message = MessageV1Codec.MessageV1({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      destChainSelector: DEST_CHAIN_SELECTOR,
      sequenceNumber: 1,
      executionGasLimit: 400_000,
      ccipReceiveGasLimit: 200_000,
      finality: finality,
      ccvAndExecutorHash: bytes32(0),
      onRampAddress: abi.encodePacked(address(0x1111111111111111111111111111111111111111)),
      offRampAddress: abi.encodePacked(address(0x2222222222222222222222222222222222222222)),
      sender: abi.encodePacked(address(0x3333333333333333333333333333333333333333)),
      receiver: abi.encodePacked(address(0x4444444444444444444444444444444444444444)),
      destBlob: "",
      tokenTransfer: tokenTransfer,
      data: ""
    });

    bytes32 messageId = keccak256(MessageV1Codec._encodeMessageV1(message));

    return (message, messageId);
  }
}
