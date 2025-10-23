// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierV1} from "../../../interfaces/ICrossChainVerifierV1.sol";

import {Client} from "../../../libraries/Client.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {OffRamp} from "../../../offRamp/OffRamp.sol";
import {OnRamp} from "../../../onRamp/OnRamp.sol";
import {FeeQuoterFeeSetup} from "../../feeQuoter/FeeQuoterSetup.t.sol";
import {MockExecutor} from "../../mocks/MockExecutor.sol";
import {MockVerifier} from "../../mocks/MockVerifier.sol";

import {IERC20Metadata} from "@openzeppelin/contracts@4.8.3/token/ERC20/extensions/IERC20Metadata.sol";

contract OnRampSetup is FeeQuoterFeeSetup {
  address internal constant FEE_AGGREGATOR = 0xa33CDB32eAEce34F6affEfF4899cef45744EDea3;

  OnRamp internal s_onRamp;
  OffRamp internal s_offRampOnRemoteChain = OffRamp(makeAddr("OffRampRemote"));

  address internal s_defaultCCV;
  address internal s_defaultExecutor;

  function setUp() public virtual override {
    super.setUp();

    s_onRamp = new OnRamp(
      OnRamp.StaticConfig({
        chainSelector: SOURCE_CHAIN_SELECTOR,
        rmnRemote: s_mockRMNRemote,
        tokenAdminRegistry: address(s_tokenAdminRegistry)
      }),
      OnRamp.DynamicConfig({
        feeQuoter: address(s_feeQuoter),
        reentrancyGuardEntered: false,
        feeAggregator: FEE_AGGREGATOR
      })
    );
    s_defaultCCV = address(new MockVerifier(""));
    s_defaultExecutor = address(new MockExecutor());

    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = s_defaultCCV;

    OnRamp.DestChainConfigArgs[] memory destChainConfigArgs = new OnRamp.DestChainConfigArgs[](1);
    destChainConfigArgs[0] = OnRamp.DestChainConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      router: s_sourceRouter,
      laneMandatedCCVs: new address[](0),
      defaultCCVs: defaultCCVs,
      defaultExecutor: s_defaultExecutor,
      offRamp: abi.encodePacked(address(s_offRampOnRemoteChain))
    });

    s_onRamp.applyDestChainConfigUpdates(destChainConfigArgs);
  }

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
      OnRamp.Receipt[] memory receipts,
      bytes[] memory verifierBlobs
    )
  {
    OnRamp.DestChainConfig memory destChainConfig = s_onRamp.getDestChainConfig(DEST_CHAIN_SELECTOR);
    MessageV1Codec.MessageV1 memory messageV1 = MessageV1Codec.MessageV1({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      destChainSelector: destChainSelector,
      sequenceNumber: seqNum,
      onRampAddress: abi.encodePacked(address(s_onRamp)),
      offRampAddress: abi.encodePacked(address(s_offRampOnRemoteChain)),
      finality: 0,
      sender: abi.encodePacked(originalSender),
      receiver: abi.encodePacked(abi.decode(message.receiver, (address))),
      destBlob: "",
      tokenTransfer: new MessageV1Codec.TokenTransferV1[](message.tokenAmounts.length),
      data: message.data
    });

    for (uint256 i = 0; i < message.tokenAmounts.length; ++i) {
      address token = message.tokenAmounts[i].token;
      messageV1.tokenTransfer[i] = MessageV1Codec.TokenTransferV1({
        amount: message.tokenAmounts[i].amount,
        sourcePoolAddress: abi.encodePacked(s_sourcePoolByToken[token]),
        sourceTokenAddress: abi.encodePacked(token),
        destTokenAddress: abi.encodePacked(s_destTokenBySourceToken[token]),
        extraData: abi.encode(IERC20Metadata(token).decimals())
      });
    }

    // If legacy extraArgs are supplied, they are passed into the CCVs and Executor.
    // If V3 extraArgs are supplied, the extraArgs as the user supplied them are used.
    bool isLegacyExtraArgs = _isLegacyExtraArgs(message.extraArgs);

    if (isLegacyExtraArgs) {
      receipts = _computeVerifierReceiptsLegacyArgs(message, destChainConfig.defaultCCVs);
    } else {
      receipts = this.computeVerifierReceiptsArgsV3(message, destChainConfig.defaultCCVs);
    }
    receipts[receipts.length - 1] = OnRamp.Receipt({
      issuer: destChainConfig.defaultExecutor,
      feeTokenAmount: 0, // Matches current OnRamp event behavior
      destGasLimit: 0,
      destBytesOverhead: 0,
      // TODO when v3 extra args are passed in
      extraArgs: isLegacyExtraArgs ? message.extraArgs : bytes("")
    });
    messageV1.destBlob = receipts[receipts.length - 1].extraArgs;

    return (
      keccak256(MessageV1Codec._encodeMessageV1(messageV1)),
      MessageV1Codec._encodeMessageV1(messageV1),
      receipts,
      new bytes[](receipts.length - message.tokenAmounts.length - 1)
    );
  }

  // This function is external so we can make the extraArgs calldata to allow for indexing.
  function computeVerifierReceiptsArgsV3(
    Client.EVM2AnyMessage calldata message,
    address[] calldata defaultCCVs
  ) external view returns (OnRamp.Receipt[] memory verifierReceipts) {
    Client.EVMExtraArgsV3 memory extraArgsV3 = abi.decode(message.extraArgs[4:], (Client.EVMExtraArgsV3));
    uint256 userDefinedCCVCount = extraArgsV3.ccvs.length;

    // Leave space for a token (if present) and the executor receipt.
    verifierReceipts = new OnRamp.Receipt[](userDefinedCCVCount + defaultCCVs.length + message.tokenAmounts.length + 1);

    uint256 currentVerifierIndex = 0;
    for (uint256 i = 0; i < userDefinedCCVCount; ++i) {
      (uint256 feeUSDCents, uint32 gasForVerification, uint32 payloadSizeBytes) = ICrossChainVerifierV1(
        extraArgsV3.ccvs[i].ccvAddress
      ).getFee(DEST_CHAIN_SELECTOR, message, extraArgsV3.ccvs[i].args, extraArgsV3.finalityConfig);

      verifierReceipts[currentVerifierIndex++] = OnRamp.Receipt({
        issuer: extraArgsV3.ccvs[i].ccvAddress,
        feeTokenAmount: feeUSDCents,
        destGasLimit: gasForVerification,
        destBytesOverhead: payloadSizeBytes,
        extraArgs: extraArgsV3.ccvs[i].args
      });
    }

    for (uint256 i = 0; i < defaultCCVs.length; ++i) {
      bool found = false;
      for (uint256 j = 0; j < userDefinedCCVCount; ++j) {
        // Skip if the default CCV is already included in the user-defined CCVs.
        if (defaultCCVs[i] == extraArgsV3.ccvs[j].ccvAddress) {
          found = true;
          break;
        }
      }

      if (found) {
        continue;
      }

      (uint256 feeUSDCents, uint32 gasForVerification, uint32 payloadSizeBytes) =
        ICrossChainVerifierV1(defaultCCVs[i]).getFee(DEST_CHAIN_SELECTOR, message, "", extraArgsV3.finalityConfig);

      verifierReceipts[currentVerifierIndex++] = OnRamp.Receipt({
        issuer: defaultCCVs[i],
        feeTokenAmount: feeUSDCents,
        destGasLimit: gasForVerification,
        destBytesOverhead: payloadSizeBytes,
        extraArgs: ""
      });
    }

    if (message.tokenAmounts.length > 0) {
      verifierReceipts[verifierReceipts.length - 2] = OnRamp.Receipt({
        issuer: message.tokenAmounts[0].token,
        destGasLimit: 0,
        destBytesOverhead: 0,
        feeTokenAmount: 0,
        extraArgs: extraArgsV3.tokenArgs
      });
    }

    return verifierReceipts;
  }

  function _computeVerifierReceiptsLegacyArgs(
    Client.EVM2AnyMessage memory message,
    address[] memory defaultCCVs
  ) internal view returns (OnRamp.Receipt[] memory verifierReceipts) {
    // Leave space for a token (if present) and the executor receipt.
    verifierReceipts = new OnRamp.Receipt[](defaultCCVs.length + message.tokenAmounts.length + 1);

    for (uint256 i = 0; i < defaultCCVs.length; ++i) {
      (uint256 feeUSDCents, uint32 gasForVerification, uint32 payloadSizeBytes) =
        ICrossChainVerifierV1(defaultCCVs[i]).getFee(DEST_CHAIN_SELECTOR, message, "", 0);

      verifierReceipts[i] = OnRamp.Receipt({
        issuer: defaultCCVs[i],
        feeTokenAmount: feeUSDCents,
        destGasLimit: gasForVerification,
        destBytesOverhead: payloadSizeBytes,
        extraArgs: ""
      });
    }

    if (message.tokenAmounts.length > 0) {
      verifierReceipts[verifierReceipts.length - 2] = OnRamp.Receipt({
        issuer: message.tokenAmounts[0].token,
        destGasLimit: 0,
        destBytesOverhead: 0,
        feeTokenAmount: 0,
        extraArgs: ""
      });
    }
    return verifierReceipts;
  }

  // Helper function to create EVMExtraArgsV3 struct
  function _createV3ExtraArgs(
    Client.CCV[] memory ccvs
  ) internal pure returns (Client.EVMExtraArgsV3 memory) {
    return Client.EVMExtraArgsV3({
      ccvs: ccvs,
      finalityConfig: 12,
      executor: address(0), // No executor specified.
      executorArgs: "",
      tokenArgs: ""
    });
  }

  function _isLegacyExtraArgs(
    bytes memory extraArgs
  ) internal pure returns (bool) {
    bytes4 selector;
    assembly {
      selector := mload(add(extraArgs, 32))
    }
    return selector != Client.GENERIC_EXTRA_ARGS_V3_TAG;
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
