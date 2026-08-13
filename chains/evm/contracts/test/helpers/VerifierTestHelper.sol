// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRouter} from "../../interfaces/IRouter.sol";

import {BaseVerifier} from "../../ccvs/components/BaseVerifier.sol";
import {FeeTokenHandler} from "../../libraries/FeeTokenHandler.sol";
import {MessageV1Codec} from "../../libraries/MessageV1Codec.sol";
import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

import {EnumerableSet} from "@openzeppelin/contracts@5.3.0/utils/structs/EnumerableSet.sol";

interface IOffRamp {
  /// @dev Per-chain source config (defining a lane from a Source Chain -> Dest OffRamp).
  struct SourceChainConfig {
    IRouter router; // ─╮ Local router to use for messages coming from this source chain.
    bool isEnabled; // ─╯ Flag whether the source chain is enabled or not.
    bytes[] onRamps; // OnRamp address on the source chain.
    address[] defaultCCVs; // Default CCVs to use for messages from this source chain.
    address[] laneMandatedCCVs; // Required CCVs to use for all messages from this source chain.
  }

  /// @notice Returns the source chain config for the provided source chain selector.
  /// @param sourceChainSelector chain to retrieve configuration for.
  /// @return sourceChainConfig The config for the source chain.
  function getSourceChainConfig(
    uint64 sourceChainSelector
  ) external view returns (SourceChainConfig memory);
}

contract VerifierTestHelper is BaseVerifier, Ownable2StepMsgSender {
  using EnumerableSet for EnumerableSet.AddressSet;

  error MessageCannotHaveSideEffects();
  error MustUseAllowlist();
  error MustUseTestRouter();
  error MustUseTestToken(address testToken);

  function typeAndVersion() external pure override returns (string memory) {
    return "VerifierTestHelper 2.0.0";
  }

  /// @notice This value is the test router, not a real production router. It is asserted all chains configured
  /// for this contract are this exact test router address.
  address internal immutable i_testRouter;

  /// @notice The test token must be supplied in all transactions
  address internal immutable i_testToken;

  constructor(
    address testRouter,
    address testToken,
    string[] memory storageLocations,
    address rmn,
    bytes4 versionTag
  ) BaseVerifier(storageLocations, rmn, versionTag) {
    if (testRouter == address(0) || testToken == address(0)) {
      revert ZeroAddressNotAllowed();
    }

    i_testRouter = testRouter;
    i_testToken = testToken;
  }

  function getTestRouter() public view returns (address) {
    return i_testRouter;
  }

  function getTestToken() public view returns (address) {
    return i_testToken;
  }

  function forwardToVerifier(
    MessageV1Codec.MessageV1 calldata message,
    bytes32, // messageId
    address, // feeToken
    uint256, // feeTokenAmount
    bytes calldata // verifierArgs
  ) external view virtual override returns (bytes memory verifierReturnData) {
    _assertNoSideEffects(message);
    _assertNotCursedByRMN(message.destChainSelector);

    if (abi.decode(message.tokenTransfer[0].sourceTokenAddress, (address)) != i_testToken) {
      revert MustUseTestToken(i_testToken);
    }

    // For EVM, sender is abi encoded.
    address senderAddress = abi.decode(message.sender, (address));
    _assertSenderIsAllowed(message.destChainSelector, senderAddress);

    return abi.encodePacked(versionTag());
  }

  function verifyMessage(
    MessageV1Codec.MessageV1 calldata message,
    bytes32, // messageId
    bytes memory // verifierResults
  ) external view virtual override {
    _assertNotCursedByRMN(message.sourceChainSelector);
    _onlyOffRamp(message.sourceChainSelector);
    _assertNoSideEffects(message);

    if (address(bytes20(message.tokenTransfer[0].destTokenAddress)) != i_testToken) {
      revert MustUseTestToken(i_testToken);
    }

    if (address(IOffRamp(msg.sender).getSourceChainConfig(message.sourceChainSelector).router) != i_testRouter) {
      revert MustUseTestRouter();
    }

    // Remote chain config, so keyed on source selector. Allowlists are shared between sending and receiving.
    RemoteChainConfig storage chainConfig = _getRemoteChainConfig(message.sourceChainSelector);

    address receiver = address(bytes20(message.receiver));
    if (chainConfig.allowlistEnabled) {
      if (!chainConfig.allowedSendersList.contains(receiver)) {
        revert SenderNotAllowed(receiver);
      }
    }
  }

  function _assertNoSideEffects(
    MessageV1Codec.MessageV1 calldata message
  ) internal view virtual {
    if (message.data.length > 0 || message.ccipReceiveGasLimit != 0) {
      revert MessageCannotHaveSideEffects();
    }

    if (message.tokenTransfer.length != 1) {
      revert MustUseTestToken(i_testToken);
    }
  }

  /// @notice Updates remote chains specific configs.
  /// @param remoteChainConfigArgs Array of remote chain specific configs.
  function applyRemoteChainConfigUpdates(
    RemoteChainConfigArgs[] calldata remoteChainConfigArgs
  ) external onlyOwner {
    for (uint256 i = 0; i < remoteChainConfigArgs.length; ++i) {
      RemoteChainConfigArgs calldata args = remoteChainConfigArgs[i];
      if (!args.allowlistEnabled) {
        revert MustUseAllowlist();
      }

      if (address(args.router) != getTestRouter()) {
        revert MustUseTestRouter();
      }
    }
    _applyRemoteChainConfigUpdates(remoteChainConfigArgs);
  }

  /// @notice Updates allowlistConfig for Senders.
  /// @dev configuration used to set the list of senders who are authorized to send messages.
  /// @param allowlistConfigArgsItems Array of AllowlistConfigArguments where each item is for a destChainSelector.
  function applyAllowlistUpdates(
    AllowlistConfigArgs[] calldata allowlistConfigArgsItems
  ) external onlyOwner {
    for (uint256 i = 0; i < allowlistConfigArgsItems.length; ++i) {
      if (!allowlistConfigArgsItems[i].allowlistEnabled) {
        revert MustUseAllowlist();
      }
    }
    _applyAllowlistUpdates(allowlistConfigArgsItems);
  }

  /// @notice Sets the finality config according to the FinalityCodec library encoding.
  /// @param allowedFinality The finality settings allowed by this verifier.
  function setAllowedFinalityConfig(
    bytes4 allowedFinality
  ) external onlyOwner {
    _setAllowedFinalityConfig(allowedFinality);
  }

  /// @notice Withdraws the outstanding fee token balances to the owner.
  /// @param feeTokens The fee tokens to withdraw.
  /// @dev This function can be permissionless as it only transfers tokens to the owner which is a trusted address.
  function withdrawFeeTokens(
    address[] calldata feeTokens
  ) external {
    FeeTokenHandler._withdrawFeeTokens(feeTokens, owner());
  }
}
