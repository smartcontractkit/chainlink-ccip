// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../../libraries/Client.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";

import {CCVAggregator} from "../../../offRamp/CCVAggregator.sol";
import {CCVProxy} from "../../../onRamp/CCVProxy.sol";
import {FeeQuoterFeeSetup} from "../../feeQuoter/FeeQuoterSetup.t.sol";

import {MockExecutor} from "../../mocks/MockExecutor.sol";
import {MockVerifier} from "../../mocks/MockVerifier.sol";

contract CCVProxySetup is FeeQuoterFeeSetup {
  address internal constant FEE_AGGREGATOR = 0xa33CDB32eAEce34F6affEfF4899cef45744EDea3;

  CCVProxy internal s_ccvProxy;
  CCVAggregator internal s_ccvAggregatorRemote;

  function setUp() public virtual override {
    super.setUp();

    s_ccvProxy = new CCVProxy(
      CCVProxy.StaticConfig({
        chainSelector: SOURCE_CHAIN_SELECTOR,
        rmnRemote: s_mockRMNRemote,
        tokenAdminRegistry: address(s_tokenAdminRegistry)
      }),
      CCVProxy.DynamicConfig({
        feeQuoter: address(s_feeQuoter),
        reentrancyGuardEntered: false,
        feeAggregator: FEE_AGGREGATOR
      })
    );
    s_ccvAggregatorRemote = CCVAggregator(makeAddr("CCVAggregatorRemote"));
    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = address(new MockVerifier(""));
    CCVProxy.DestChainConfigArgs[] memory destChainConfigArgs = new CCVProxy.DestChainConfigArgs[](1);
    destChainConfigArgs[0] = CCVProxy.DestChainConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      router: s_sourceRouter,
      laneMandatedCCVs: new address[](0),
      defaultCCVs: defaultCCVs,
      defaultExecutor: address(new MockExecutor()),
      ccvAggregator: abi.encodePacked(address(s_ccvAggregatorRemote))
    });

    s_ccvProxy.applyDestChainConfigUpdates(destChainConfigArgs);
  }

  // TODO make this work for other cases as well
  function _evmMessageToEvent(
    Client.EVM2AnyMessage memory message,
    uint64 destChainSelector,
    uint64 seqNum,
    address originalSender
  )
    internal
    view
    returns (
      bytes32 messageId,
      bytes memory encodedMessage,
      CCVProxy.Receipt[] memory verifierReceipts,
      CCVProxy.Receipt memory executorReceipt,
      bytes[] memory receiptBlobs
    )
  {
    // TODO handle token transfers
    CCVProxy.DestChainConfig memory destChainConfig = s_ccvProxy.getDestChainConfig(DEST_CHAIN_SELECTOR);
    MessageV1Codec.MessageV1 memory messageV1 = MessageV1Codec.MessageV1({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      destChainSelector: destChainSelector,
      sequenceNumber: seqNum,
      onRampAddress: abi.encodePacked(address(s_ccvProxy)),
      offRampAddress: abi.encodePacked(address(s_ccvAggregatorRemote)),
      finality: 0,
      sender: abi.encodePacked(originalSender),
      receiver: abi.encodePacked(abi.decode(message.receiver, (address))),
      destBlob: "",
      tokenTransfer: new MessageV1Codec.TokenTransferV1[](message.tokenAmounts.length),
      data: message.data
    });

    verifierReceipts = new CCVProxy.Receipt[](destChainConfig.defaultCCVs.length);
    for (uint256 i = 0; i < verifierReceipts.length; ++i) {
      verifierReceipts[i] = CCVProxy.Receipt({
        issuer: destChainConfig.defaultCCVs[i],
        feeTokenAmount: 0,
        destGasLimit: 0,
        destBytesOverhead: 0,
        // TODO when v3 extra args are passed in
        extraArgs: message.extraArgs
      });
    }
    executorReceipt = CCVProxy.Receipt({
      issuer: destChainConfig.defaultExecutor,
      feeTokenAmount: 0, // Matches current CCVProxy event behavior
      destGasLimit: 0,
      destBytesOverhead: 0,
      extraArgs: message.extraArgs
    });
    receiptBlobs = new bytes[](1);
    receiptBlobs[0] = "";
    return (
      keccak256(MessageV1Codec._encodeMessageV1(messageV1)),
      MessageV1Codec._encodeMessageV1(messageV1),
      verifierReceipts,
      executorReceipt,
      receiptBlobs
    );
  }

  // Helper function to create EVMExtraArgsV3 struct
  function _createV3ExtraArgs(
    Client.CCV[] memory requiredCCVs,
    Client.CCV[] memory optionalCCVs,
    uint8 optionalThreshold
  ) internal pure returns (Client.EVMExtraArgsV3 memory) {
    return Client.EVMExtraArgsV3({
      requiredCCV: requiredCCVs,
      optionalCCV: optionalCCVs,
      optionalThreshold: optionalThreshold,
      finalityConfig: 12,
      executor: address(0), // No executor specified.
      executorArgs: "",
      tokenArgs: ""
    });
  }

  // Helper function to assert that two CCV arrays are equal
  function _assertCCVArraysEqual(Client.CCV[] memory actual, Client.CCV[] memory expected) internal pure {
    assertEq(actual.length, expected.length, "CCV arrays have different lengths");

    for (uint256 i = 0; i < actual.length; i++) {
      assertEq(
        actual[i].ccvAddress, expected[i].ccvAddress, string.concat("CCV address mismatch at index ", vm.toString(i))
      );
      assertEq(actual[i].args, expected[i].args, string.concat("CCV args mismatch at index ", vm.toString(i)));
    }
  }
}
