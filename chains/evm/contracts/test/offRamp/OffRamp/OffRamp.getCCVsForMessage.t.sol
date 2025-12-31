// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {OffRampSetup} from "./OffRampSetup.t.sol";

contract OffRamp_getCCVsForMessage is OffRampSetup {
  function test_getCCVsForMessage() public {
    // Setup default CCVs for the source chain.
    address[] memory defaultCCVs = new address[](2);
    defaultCCVs[0] = makeAddr("defaultCCV1");
    defaultCCVs[1] = makeAddr("defaultCCV2");

    _applySourceConfig(s_onRamp, true, defaultCCVs, new address[](0));

    // Create a simple message.
    MessageV1Codec.MessageV1 memory message = MessageV1Codec.MessageV1({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      destChainSelector: DEST_CHAIN_SELECTOR,
      messageNumber: 1,
      executionGasLimit: 200_000,
      ccipReceiveGasLimit: 0,
      finality: 0,
      ccvAndExecutorHash: bytes32(0),
      onRampAddress: s_onRamp,
      offRampAddress: abi.encodePacked(s_offRamp),
      sender: abi.encodePacked(makeAddr("sender")),
      receiver: abi.encodePacked(makeAddr("receiver")),
      destBlob: "",
      tokenTransfer: new MessageV1Codec.TokenTransferV1[](0),
      data: ""
    });

    (address[] memory requiredCCVs, address[] memory optionalCCVs, uint8 threshold) =
      s_offRamp.getCCVsForMessage(MessageV1Codec._encodeMessageV1(message));

    // Should return default CCVs since receiver doesn't specify any and there are no token transfers.
    assertEq(requiredCCVs.length, defaultCCVs.length);
    assertEq(optionalCCVs.length, 0);
    assertEq(threshold, 0);
    for (uint256 i = 0; i < defaultCCVs.length; ++i) {
      assertEq(requiredCCVs[i], defaultCCVs[i]);
    }
  }
}

