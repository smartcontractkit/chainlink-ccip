// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITokenMessenger} from "./interfaces/ITokenMessenger.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/utils/SafeERC20.sol";
import {EnumerableSet} from "@openzeppelin/contracts@5.3.0/utils/structs/EnumerableSet.sol";

/// @title CCTP TokenMessenger Proxy
/// @notice A proxy contract that deposits tokens for burn via the Cross Chain Transfer Protocol (CCTP).
/// @dev This contract is responsible for calling depositForBurnWithHook on the TokenMessenger and ensuring that only authorized callers can invoke it.
/// @dev The purpose of this contract is to provide a static CCTP message sender on source that can be reliably validated on destination.
contract CCTPTokenMessengerProxy is AuthorizedCallers, ITypeAndVersion, ITokenMessenger {
  using SafeERC20 for IERC20;
  using EnumerableSet for EnumerableSet.AddressSet;

  string public constant override typeAndVersion = "CCTPTokenMessengerProxy 1.7.0-dev";

  /// @notice The TokenMessenger contract.
  ITokenMessenger private immutable i_tokenMessenger;
  /// @notice The USDC token contract.
  IERC20 private immutable i_usdcToken;

  constructor(
    ITokenMessenger tokenMessenger,
    IERC20 usdcToken,
    address[] memory authorizedCallers
  ) AuthorizedCallers(authorizedCallers) {
    i_tokenMessenger = tokenMessenger;
    i_usdcToken = usdcToken;

    // Approve the TokenMessenger to burn the USDC token on behalf of this contract.
    // The CCTP verifier is responsible for forwarding the USDC it receives from the USDC token pool to this contract.
    i_usdcToken.safeIncreaseAllowance(address(tokenMessenger), type(uint256).max);
  }

  /// @inheritdoc ITokenMessenger
  function depositForBurnWithHook(
    uint256 amount,
    uint32 destinationDomain,
    bytes32 mintRecipient,
    address burnToken,
    bytes32 destinationCaller,
    uint32 maxFee,
    uint32 minFinalityThreshold,
    bytes calldata hookData
  ) external onlyAuthorizedCallers {
    i_tokenMessenger.depositForBurnWithHook(
      amount, destinationDomain, mintRecipient, burnToken, destinationCaller, maxFee, minFinalityThreshold, hookData
    );
  }

  /// @inheritdoc ITokenMessenger
  function depositForBurn(
    uint256 amount,
    uint32 destinationDomain,
    bytes32 mintRecipient,
    address burnToken,
    bytes32 destinationCaller,
    uint32 maxFee,
    uint32 minFinalityThreshold
  ) external onlyAuthorizedCallers {
    i_tokenMessenger.depositForBurn(
      amount, destinationDomain, mintRecipient, burnToken, destinationCaller, maxFee, minFinalityThreshold
    );
  }

  /// @inheritdoc ITokenMessenger
  function depositForBurnWithCaller(
    uint256 amount,
    uint32 destinationDomain,
    bytes32 mintRecipient,
    address burnToken,
    bytes32 destinationCaller
  ) external onlyAuthorizedCallers returns (uint64) {
    return
      i_tokenMessenger.depositForBurnWithCaller(amount, destinationDomain, mintRecipient, burnToken, destinationCaller);
  }

  /// @inheritdoc ITokenMessenger
  function messageBodyVersion() external view returns (uint32) {
    return i_tokenMessenger.messageBodyVersion();
  }

  /// @inheritdoc ITokenMessenger
  function localMessageTransmitter() external view returns (address) {
    return i_tokenMessenger.localMessageTransmitter();
  }

  /// @notice Returns the TokenMessenger contract.
  /// @return tokenMessenger The TokenMessenger contract.
  function getTokenMessenger() external view returns (address) {
    return address(i_tokenMessenger);
  }

  /// @notice Returns the USDC token contract.
  /// @return usdcToken The USDC token contract.
  function getUSDCToken() external view returns (address) {
    return address(i_usdcToken);
  }
}
