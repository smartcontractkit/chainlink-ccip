// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../../libraries/Client.sol";
import {ExtraArgsCodec} from "../../../libraries/ExtraArgsCodec.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {OffRamp} from "../../../offRamp/OffRamp.sol";
import {OnRamp} from "../../../onRamp/OnRamp.sol";
import {FeeQuoterFeeSetup} from "../../feeQuoter/FeeQuoterSetup.t.sol";
import {OnRampHelper} from "../../helpers/OnRampHelper.sol";
import {MockExecutor} from "../../mocks/MockExecutor.sol";
import {MockVerifier} from "../../mocks/MockVerifier.sol";

import {IERC20Metadata} from "@openzeppelin/contracts@4.8.3/token/ERC20/extensions/IERC20Metadata.sol";

contract OnRampSetup is FeeQuoterFeeSetup {
  address internal constant FEE_AGGREGATOR = 0xa33CDB32eAEce34F6affEfF4899cef45744EDea3;
  uint32 internal constant POOL_FEE_USD_CENTS = 100; // $1.00
  uint32 internal constant POOL_GAS_OVERHEAD = 50000;
  uint32 internal constant POOL_BYTES_OVERHEAD = 128;

  uint32 internal constant FEE_QUOTER_FEE_USD_CENTS = 50; // $0.50
  uint32 internal constant FEE_QUOTER_GAS_OVERHEAD = 30000;
  uint32 internal constant FEE_QUOTER_BYTES_OVERHEAD = 64;

  uint32 internal constant VERIFIER_FEE_USD_CENTS = 200; // $2.00
  uint32 internal constant VERIFIER_GAS = 100000;
  uint32 internal constant VERIFIER_BYTES = 256;

  OnRampHelper internal s_onRamp;
  OffRamp internal s_offRampOnRemoteChain = OffRamp(makeAddr("OffRampRemote"));

  address internal s_defaultCCV;
  address internal s_defaultExecutor;

  function setUp() public virtual override {
    super.setUp();

    s_onRamp = new OnRampHelper(
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
      addressBytesLength: EVM_ADDRESS_LENGTH,
      baseExecutionGasCost: BASE_EXEC_GAS_COST,
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

    ExtraArgsCodec.GenericExtraArgsV3 memory resolvedExtraArgs =
      s_onRamp.parseExtraArgsWithDefaults(destChainSelector, destChainConfig, message.extraArgs);

    address[] memory poolRequiredCCVs = new address[](0);
    if (message.tokenAmounts.length != 0) {
      poolRequiredCCVs = s_onRamp.getCCVsForPool(
        destChainSelector,
        message.tokenAmounts[0].token,
        message.tokenAmounts[0].amount,
        resolvedExtraArgs.blockConfirmations,
        resolvedExtraArgs.tokenArgs
      );
    }
    (resolvedExtraArgs.ccvs, resolvedExtraArgs.ccvArgs) = s_onRamp.mergeCCVLists(
      resolvedExtraArgs.ccvs, resolvedExtraArgs.ccvArgs, destChainConfig.laneMandatedCCVs, poolRequiredCCVs
    );

    MessageV1Codec.MessageV1 memory messageV1 = MessageV1Codec.MessageV1({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      destChainSelector: destChainSelector,
      sequenceNumber: seqNum,
      executionGasLimit: 0, // populated below.
      ccipReceiveGasLimit: GAS_LIMIT,
      finality: 0,
      ccvAndExecutorHash: MessageV1Codec._computeCCVAndExecutorHash(resolvedExtraArgs.ccvs, resolvedExtraArgs.executor),
      onRampAddress: abi.encodePacked(address(s_onRamp)),
      offRampAddress: abi.encodePacked(address(s_offRampOnRemoteChain)),
      sender: abi.encodePacked(originalSender),
      receiver: abi.encodePacked(abi.decode(message.receiver, (address))),
      destBlob: "",
      tokenTransfer: new MessageV1Codec.TokenTransferV1[](message.tokenAmounts.length),
      data: message.data
    });

    // Populate token transfers
    _populateTokenTransfers(messageV1, message);

    // Compute receipts
    (receipts, messageV1.executionGasLimit,) = s_onRamp.getReceipts(destChainSelector, message, resolvedExtraArgs);

    return (
      keccak256(MessageV1Codec._encodeMessageV1(messageV1)),
      MessageV1Codec._encodeMessageV1(messageV1),
      receipts,
      new bytes[](receipts.length - message.tokenAmounts.length - 1)
    );
  }

  // Helper function to create GenericExtraArgsV3 struct.
  function _createV3ExtraArgs(
    address[] memory ccvAddresses,
    bytes[] memory ccvArgs
  ) internal pure returns (ExtraArgsCodec.GenericExtraArgsV3 memory) {
    return ExtraArgsCodec.GenericExtraArgsV3({
      ccvs: ccvAddresses,
      ccvArgs: ccvArgs,
      blockConfirmations: 12,
      gasLimit: GAS_LIMIT,
      executor: address(0), // No executor specified.
      executorArgs: "",
      tokenReceiver: "",
      tokenArgs: ""
    });
  }

  // Helper function to assert that two CCV arrays are equal (using parallel address and bytes arrays).
  function _assertCCVArraysEqual(
    address[] memory actualAddresses,
    bytes[] memory actualArgs,
    address[] memory expectedAddresses,
    bytes[] memory expectedArgs
  ) internal pure {
    assertEq(actualAddresses.length, expectedAddresses.length, "CCV address arrays have different lengths");
    assertEq(actualArgs.length, expectedArgs.length, "CCV args arrays have different lengths");
    assertEq(actualAddresses.length, actualArgs.length, "CCV addresses and args have different lengths");

    for (uint256 i = 0; i < actualAddresses.length; i++) {
      assertEq(
        actualAddresses[i], expectedAddresses[i], string.concat("CCV address mismatch at index ", vm.toString(i))
      );
      assertEq(actualArgs[i], expectedArgs[i], string.concat("CCV args mismatch at index ", vm.toString(i)));
    }
  }

  // Helper to populate token transfers.
  function _populateTokenTransfers(
    MessageV1Codec.MessageV1 memory messageV1,
    Client.EVM2AnyMessage memory message
  ) internal view {
    for (uint256 i = 0; i < message.tokenAmounts.length; ++i) {
      address token = message.tokenAmounts[i].token;
      messageV1.tokenTransfer[i] = MessageV1Codec.TokenTransferV1({
        amount: message.tokenAmounts[i].amount,
        sourcePoolAddress: abi.encodePacked(s_sourcePoolByToken[token]),
        sourceTokenAddress: abi.encodePacked(token),
        destTokenAddress: abi.encodePacked(s_destTokenBySourceToken[token]),
        tokenReceiver: abi.encodePacked(abi.decode(message.receiver, (address))),
        extraData: abi.encode(IERC20Metadata(token).decimals())
      });
    }
  }
}
