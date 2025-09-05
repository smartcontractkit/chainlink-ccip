// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../../libraries/Client.sol";
import {Internal} from "../../../libraries/Internal.sol";
import {CCVProxy} from "../../../onRamp/CCVProxy.sol";
import {FeeQuoterFeeSetup} from "../../feeQuoter/FeeQuoterSetup.t.sol";
import {MockCCVOnRamp} from "../../mocks/MockCCVOnRamp.sol";
import {MockExecutor} from "../../mocks/MockExecutor.sol";

contract CCVProxySetup is FeeQuoterFeeSetup {
  address internal constant FEE_AGGREGATOR = 0xa33CDB32eAEce34F6affEfF4899cef45744EDea3;
  bytes32 internal s_metadataHash;

  CCVProxy internal s_ccvProxy;

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
    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = address(new MockCCVOnRamp());
    CCVProxy.DestChainConfigArgs[] memory destChainConfigArgs = new CCVProxy.DestChainConfigArgs[](1);
    destChainConfigArgs[0] = CCVProxy.DestChainConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      router: s_sourceRouter,
      laneMandatedCCVs: new address[](0),
      defaultCCVs: defaultCCVs,
      defaultExecutor: address(new MockExecutor())
    });

    s_ccvProxy.applyDestChainConfigUpdates(destChainConfigArgs);

    // Calculate the metadata hash the same way the contract does
    s_metadataHash = keccak256(
      abi.encode(Internal.EVM_2_ANY_MESSAGE_HASH, SOURCE_CHAIN_SELECTOR, DEST_CHAIN_SELECTOR, address(s_ccvProxy))
    );
  }

  // TODO make this work for other cases as well
  function _evmMessageToEvent(
    Client.EVM2AnyMessage memory message,
    uint64 destChainSelector,
    uint64 seqNum,
    uint256 feeTokenAmount,
    uint256 feeValueJuels,
    address originalSender,
    bytes32 metadataHash
  ) internal view returns (Internal.EVM2AnyVerifierMessage memory) {
    // TODO
    //    if (message.tokenAmounts.length > 0) {
    //      tokenTransfer =
    //        Internal.EVM2AnyTokenTransfer({token: message.tokenAmounts[0].token, amount: message.tokenAmounts[0].amount});
    //    }
    CCVProxy.DestChainConfig memory destChainConfig = s_ccvProxy.getDestChainConfig(DEST_CHAIN_SELECTOR);
    Internal.EVM2AnyVerifierMessage memory messageEvent = Internal.EVM2AnyVerifierMessage({
      header: Internal.Header({
        messageId: "",
        sourceChainSelector: SOURCE_CHAIN_SELECTOR,
        destChainSelector: destChainSelector,
        sequenceNumber: seqNum
      }),
      sender: originalSender,
      data: message.data,
      receiver: message.receiver,
      feeToken: message.feeToken,
      feeTokenAmount: feeTokenAmount,
      feeValueJuels: feeValueJuels,
      tokenTransfer: new Internal.EVMTokenTransfer[](message.tokenAmounts.length),
      verifierReceipts: new Internal.Receipt[](destChainConfig.defaultCCVs.length),
      executorReceipt: Internal.Receipt({
        issuer: address(0),
        feeTokenAmount: 0,
        destGasLimit: 0,
        destBytesOverhead: 0,
        extraArgs: ""
      })
    });

    for (uint256 i = 0; i < messageEvent.verifierReceipts.length; i++) {
      messageEvent.verifierReceipts[i] = Internal.Receipt({
        issuer: destChainConfig.defaultCCVs[i],
        feeTokenAmount: 0,
        destGasLimit: 0,
        destBytesOverhead: 0,
        // TODO when v3 extra args are passed in
        extraArgs: message.extraArgs
      });
    }
    messageEvent.header.messageId = Internal._hash(messageEvent, metadataHash);
    return messageEvent;
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
