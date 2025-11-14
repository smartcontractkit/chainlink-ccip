// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ICrossChainVerifierResolver} from "../../../interfaces/ICrossChainVerifierResolver.sol";
import {ICrossChainVerifierV1} from "../../../interfaces/ICrossChainVerifierV1.sol";
import {IFeeQuoter} from "../../../interfaces/IFeeQuoter.sol";
import {IPoolV1} from "../../../interfaces/IPool.sol";
import {IPoolV2} from "../../../interfaces/IPoolV2.sol";

import {Client} from "../../../libraries/Client.sol";
import {ExtraArgsCodec} from "../../../libraries/ExtraArgsCodec.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {OffRamp} from "../../../offRamp/OffRamp.sol";
import {OnRamp} from "../../../onRamp/OnRamp.sol";
import {FeeQuoterFeeSetup} from "../../feeQuoter/FeeQuoterSetup.t.sol";
import {MockExecutor} from "../../mocks/MockExecutor.sol";
import {MockVerifier} from "../../mocks/MockVerifier.sol";

import {IERC20Metadata} from "@openzeppelin/contracts@4.8.3/token/ERC20/extensions/IERC20Metadata.sol";
import {IERC165} from "@openzeppelin/contracts@5.0.2/utils/introspection/IERC165.sol";

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
    MessageV1Codec.MessageV1 memory messageV1 = MessageV1Codec.MessageV1({
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      destChainSelector: destChainSelector,
      sequenceNumber: seqNum,
      onRampAddress: abi.encodePacked(address(s_onRamp)),
      offRampAddress: abi.encodePacked(address(s_offRampOnRemoteChain)),
      finality: 0,
      gasLimit: GAS_LIMIT,
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
        tokenReceiver: abi.encodePacked(abi.decode(message.receiver, (address))),
        extraData: abi.encode(IERC20Metadata(token).decimals())
      });
    }

    // If legacy extraArgs are supplied, they are passed into the CCVs and Executor.
    // If V3 extraArgs are supplied, the extraArgs as the user supplied them are used.
    bool isLegacyExtraArgs = _isLegacyExtraArgs(message.extraArgs);

    if (isLegacyExtraArgs) {
      receipts = _computeVerifierReceiptsLegacyArgs(message, destChainConfig.defaultCCVs);
    } else {
      (receipts, messageV1.gasLimit) = this.computeVerifierReceiptsArgsV3(message, destChainConfig.defaultCCVs);
    }
    receipts[receipts.length - 1] = OnRamp.Receipt({
      issuer: destChainConfig.defaultExecutor,
      feeTokenAmount: 0, // Matches current OnRamp event behavior
      destGasLimit: destChainConfig.baseExecutionGasCost + GAS_LIMIT,
      destBytesOverhead: _calculateDestBytesOverhead(
        uint32(message.data.length), destChainConfig.addressBytesLength, uint32(message.tokenAmounts.length), 0
      ),
      // TODO when v3 extra args are passed in
      extraArgs: bytes("")
    });

    return (
      keccak256(MessageV1Codec._encodeMessageV1(messageV1)),
      MessageV1Codec._encodeMessageV1(messageV1),
      receipts,
      new bytes[](receipts.length - message.tokenAmounts.length - 1)
    );
  }

  function _calculateDestBytesOverhead(
    uint32 dataLength,
    uint32 remoteChainAddressLengthBytes,
    uint32 numberOfTokens,
    uint256 executorArgsLength
  ) internal pure returns (uint32) {
    return uint32(
      MessageV1Codec.MESSAGE_V1_EVM_SOURCE_BASE_SIZE + dataLength + executorArgsLength
        + (MessageV1Codec.MESSAGE_V1_REMOTE_CHAIN_ADDRESSES * remoteChainAddressLengthBytes)
        + (numberOfTokens * (MessageV1Codec.TOKEN_TRANSFER_V1_EVM_SOURCE_BASE_SIZE + remoteChainAddressLengthBytes))
    );
  }

  // This function is external so we can make the extraArgs calldata to allow for indexing.
  function computeVerifierReceiptsArgsV3(
    Client.EVM2AnyMessage calldata message,
    address[] calldata defaultCCVs
  ) external view returns (OnRamp.Receipt[] memory verifierReceipts, uint32 gasLimit) {
    ExtraArgsCodec.GenericExtraArgsV3 memory extraArgsV3 = ExtraArgsCodec._decodeGenericExtraArgsV3(message.extraArgs);
    uint256 userDefinedCCVCount = extraArgsV3.ccvs.length;

    // Leave space for a token (if present) and the executor receipt.
    verifierReceipts = new OnRamp.Receipt[](userDefinedCCVCount + defaultCCVs.length + message.tokenAmounts.length + 1);

    uint256 currentVerifierIndex = 0;
    for (uint256 i = 0; i < userDefinedCCVCount; ++i) {
      address implAddress =
        ICrossChainVerifierResolver(extraArgsV3.ccvs[i]).getOutboundImplementation(DEST_CHAIN_SELECTOR, "");
      (uint256 feeUSDCents, uint32 gasForVerification, uint32 payloadSizeBytes) = ICrossChainVerifierV1(implAddress)
        .getFee(DEST_CHAIN_SELECTOR, message, extraArgsV3.ccvArgs[i], extraArgsV3.blockConfirmations);

      verifierReceipts[currentVerifierIndex++] = OnRamp.Receipt({
        issuer: extraArgsV3.ccvs[i],
        feeTokenAmount: feeUSDCents,
        destGasLimit: gasForVerification,
        destBytesOverhead: payloadSizeBytes,
        extraArgs: extraArgsV3.ccvArgs[i]
      });
    }

    for (uint256 i = 0; i < defaultCCVs.length; ++i) {
      bool found = false;
      for (uint256 j = 0; j < userDefinedCCVCount; ++j) {
        // Skip if the default CCV is already included in the user-defined CCVs.
        if (defaultCCVs[i] == extraArgsV3.ccvs[j]) {
          found = true;
          break;
        }
      }

      if (found) {
        continue;
      }

      (uint256 feeUSDCents, uint32 gasForVerification, uint32 payloadSizeBytes) = ICrossChainVerifierV1(
        ICrossChainVerifierResolver(defaultCCVs[i]).getOutboundImplementation(DEST_CHAIN_SELECTOR, "")
      ).getFee(DEST_CHAIN_SELECTOR, message, "", extraArgsV3.blockConfirmations);

      verifierReceipts[currentVerifierIndex++] = OnRamp.Receipt({
        issuer: defaultCCVs[i],
        feeTokenAmount: feeUSDCents,
        destGasLimit: gasForVerification,
        destBytesOverhead: payloadSizeBytes,
        extraArgs: ""
      });
    }

    if (message.tokenAmounts.length > 0) {
      (uint256 feeUSDCents, uint32 destGasOverhead, uint32 destBytesOverhead) = _getPoolFee(
        message.tokenAmounts[0].token,
        message.tokenAmounts[0].amount,
        message.feeToken,
        extraArgsV3.blockConfirmations,
        extraArgsV3.tokenArgs
      );

      verifierReceipts[verifierReceipts.length - 2] = OnRamp.Receipt({
        issuer: message.tokenAmounts[0].token,
        destGasLimit: destGasOverhead,
        destBytesOverhead: destBytesOverhead,
        feeTokenAmount: feeUSDCents,
        extraArgs: extraArgsV3.tokenArgs
      });
    }

    return (verifierReceipts, extraArgsV3.gasLimit);
  }

  /// @notice Helper to get pool fee for a token, with fallback to FeeQuoter.
  function _getPoolFee(
    address token,
    uint256 amount,
    address feeToken,
    uint16 finalityConfig,
    bytes memory tokenArgs
  ) internal view virtual returns (uint256 feeUSDCents, uint32 destGasOverhead, uint32 destBytesOverhead) {
    IPoolV1 pool = IPoolV1(address(s_tokenAdminRegistry.getPool(token)));
    bool isEnabled = false;

    // Try to call getFee if the pool supports IPoolV2.
    if (IERC165(address(pool)).supportsInterface(type(IPoolV2).interfaceId)) {
      (feeUSDCents, destGasOverhead, destBytesOverhead,, isEnabled) =
        IPoolV2(address(pool)).getFee(token, DEST_CHAIN_SELECTOR, amount, feeToken, finalityConfig, tokenArgs);
    }

    // If the pool doesn't support IPoolV2 or config is disabled, fall back to FeeQuoter.
    if (!isEnabled) {
      (feeUSDCents, destGasOverhead, destBytesOverhead) =
        IFeeQuoter(address(s_feeQuoter)).getTokenTransferFee(DEST_CHAIN_SELECTOR, token);
    }
    return (feeUSDCents, destGasOverhead, destBytesOverhead);
  }

  function _computeVerifierReceiptsLegacyArgs(
    Client.EVM2AnyMessage memory message,
    address[] memory defaultCCVs
  ) internal view returns (OnRamp.Receipt[] memory verifierReceipts) {
    // Leave space for a token (if present) and the executor receipt.
    verifierReceipts = new OnRamp.Receipt[](defaultCCVs.length + message.tokenAmounts.length + 1);

    for (uint256 i = 0; i < defaultCCVs.length; ++i) {
      address implAddress =
        ICrossChainVerifierResolver(defaultCCVs[i]).getOutboundImplementation(DEST_CHAIN_SELECTOR, "");
      (uint256 feeUSDCents, uint32 gasForVerification, uint32 payloadSizeBytes) =
        ICrossChainVerifierV1(implAddress).getFee(DEST_CHAIN_SELECTOR, message, "", 0);

      verifierReceipts[i] = OnRamp.Receipt({
        issuer: defaultCCVs[i],
        feeTokenAmount: feeUSDCents,
        destGasLimit: gasForVerification,
        destBytesOverhead: payloadSizeBytes,
        extraArgs: ""
      });
    }

    if (message.tokenAmounts.length > 0) {
      (uint256 feeUSDCents, uint32 destGasOverhead, uint32 destBytesOverhead) =
        _getPoolFee(message.tokenAmounts[0].token, message.tokenAmounts[0].amount, message.feeToken, 0, "");

      verifierReceipts[verifierReceipts.length - 2] = OnRamp.Receipt({
        issuer: message.tokenAmounts[0].token,
        destGasLimit: destGasOverhead,
        destBytesOverhead: destBytesOverhead,
        feeTokenAmount: feeUSDCents,
        extraArgs: ""
      });
    }
    return verifierReceipts;
  }

  // Helper function to create GenericExtraArgsV3 struc
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

  function _isLegacyExtraArgs(
    bytes memory extraArgs
  ) internal pure returns (bool) {
    bytes4 selector;
    assembly {
      selector := mload(add(extraArgs, 32))
    }
    return selector != ExtraArgsCodec.GENERIC_EXTRA_ARGS_V3_TAG;
  }

  // Helper function to assert that two CCV arrays are equal (using parallel address and bytes arrays)
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
}
