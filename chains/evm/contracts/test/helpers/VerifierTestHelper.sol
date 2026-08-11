// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../ccvs/components/BaseVerifier.sol";
import {FeeTokenHandler} from "../../libraries/FeeTokenHandler.sol";
import {MessageV1Codec} from "../../libraries/MessageV1Codec.sol";
import {Ownable2StepMsgSender} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2StepMsgSender.sol";

import {EnumerableSet} from "@openzeppelin/contracts@5.3.0/utils/structs/EnumerableSet.sol";

contract VerifierTestHelper is BaseVerifier, Ownable2StepMsgSender {
  using EnumerableSet for EnumerableSet.AddressSet;

  error MessageCannotHaveSideEffects();

  function typeAndVersion() external pure override returns (string memory) {
    return "VerifierTestHelper 2.0.0";
  }

  constructor(
    string[] memory storageLocations,
    address rmn,
    bytes4 versionTag
  ) BaseVerifier(storageLocations, rmn, versionTag) {}

  function forwardToVerifier(
    MessageV1Codec.MessageV1 calldata message,
    bytes32, // messageId
    address, // feeToken
    uint256, // feeTokenAmount
    bytes calldata // verifierArgs
  ) external view virtual override returns (bytes memory verifierReturnData) {
    _assertNoSideEffects(message);

    _assertNotCursedByRMN(message.destChainSelector);

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

    RemoteChainConfig storage chainConfig = _getRemoteChainConfig(message.destChainSelector);

    address receiver = address(bytes20(message.receiver));
    if (chainConfig.allowlistEnabled) {
      if (!chainConfig.allowedSendersList.contains(receiver)) {
        revert SenderNotAllowed(receiver);
      }
    }
  }

  function _assertNoSideEffects(
    MessageV1Codec.MessageV1 calldata message
  ) internal pure virtual {
    if (message.tokenTransfer.length > 0 || message.data.length > 0 || message.ccipReceiveGasLimit != 0) {
      revert MessageCannotHaveSideEffects();
    }
  }

  /// @notice Updates remote chains specific configs.
  /// @param remoteChainConfigArgs Array of remote chain specific configs.
  function applyRemoteChainConfigUpdates(
    RemoteChainConfigArgs[] calldata remoteChainConfigArgs
  ) external onlyOwner {
    _applyRemoteChainConfigUpdates(remoteChainConfigArgs);
  }

  /// @notice Updates allowlistConfig for Senders.
  /// @dev configuration used to set the list of senders who are authorized to send messages.
  /// @param allowlistConfigArgsItems Array of AllowlistConfigArguments where each item is for a destChainSelector.
  function applyAllowlistUpdates(
    AllowlistConfigArgs[] calldata allowlistConfigArgsItems
  ) external onlyOwner {
    _applyAllowlistUpdates(allowlistConfigArgsItems);
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
