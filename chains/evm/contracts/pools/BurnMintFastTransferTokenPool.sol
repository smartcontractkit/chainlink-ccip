// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAny2EVMMessageReceiver} from "../interfaces/IAny2EVMMessageReceiver.sol";
import {IFastTransferPool} from "../interfaces/IFastTransferPool.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";
import {IBurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/IBurnMintERC20.sol";

import {BurnMintTokenPoolAbstract} from "./BurnMintTokenPoolAbstract.sol";
import {FastTransferTokenPoolAbstract} from "./FastTransferTokenPoolAbstract.sol";
import {TokenPool} from "./TokenPool.sol";

import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";
import {IERC165} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/utils/introspection/IERC165.sol";

/// @title BurnMintFastTransferTokenPool
/// @notice A token pool that supports burn-mint operations and fast transfers
/// @dev Inherits from BurnMintTokenPoolAbstract and FastTransferTokenPoolAbstract
contract BurnMintFastTransferTokenPool is ITypeAndVersion, BurnMintTokenPoolAbstract, FastTransferTokenPoolAbstract {
  using SafeERC20 for IERC20;

  string public constant override typeAndVersion = "BurnMintFastTransferTokenPool 1.6.1";

  constructor(
    IBurnMintERC20 token,
    uint8 localTokenDecimals,
    address[] memory allowlist,
    address rmnProxy,
    address router
  ) FastTransferTokenPoolAbstract(token, localTokenDecimals, allowlist, rmnProxy, router) {}

  /// @notice Handles the transfer of tokens when a fast transfer is initiated
  function _handleTokenToTransfer(uint64, address, uint256 amount) internal override {
    _burn(amount);
  }

  /// @notice Handles the transfer of tokens when a fast transfer is filled
  /// @dev This function is called when a fast transfer is filled by a filler
  /// @param sourceChainSelector The selector of the source chain
  /// @param filler The address of the filler who filled the fast transfer
  /// @param receiver The address of the receiver who will receive the tokens
  /// @param srcAmount The amount of tokens being transferred from the source chain
  /// @param sourceDecimals The number of decimals of the source token
  function _transferFromFiller(
    uint64 sourceChainSelector,
    address filler,
    address receiver,
    uint256 srcAmount,
    uint8 sourceDecimals
  ) internal override returns (uint256 localAmount) {
    localAmount = _calculateLocalAmount(srcAmount, sourceDecimals);
    _consumeInboundRateLimit(sourceChainSelector, localAmount);
    getToken().safeTransferFrom(filler, receiver, localAmount);
    return localAmount;
  }

  /// @notice Handles settlement when the request was not fast-filled
  /// @param sourceChainSelector The selector of the source chain
  /// @param settlementAmountLocal The amount of tokens to settle in the local chain
  /// @param receiver The address of the receiver who will receive the settled tokens
  function _handleNotFastFilled(
    uint64 sourceChainSelector,
    uint256 settlementAmountLocal,
    address receiver
  ) internal override {
    _consumeInboundRateLimit(sourceChainSelector, settlementAmountLocal);
    IBurnMintERC20(address(i_token)).mint(receiver, settlementAmountLocal);
  }

  /// @notice Handles reimbursement when the request was fast-filled
  /// @param filler The address of the filler who filled the fast transfer
  /// @param settlementAmountLocal The amount of tokens to reimburse in the local chain
  function _handleFastFilledReimbursement(address filler, uint256 settlementAmountLocal) internal override {
    // Honest filler -> pay them back + fee
    IBurnMintERC20(address(i_token)).mint(filler, settlementAmountLocal);
  }

  /// @notice Signals which version of the pool interface is supported
  function supportsInterface(
    bytes4 interfaceId
  ) public pure virtual override(TokenPool, FastTransferTokenPoolAbstract) returns (bool) {
    return interfaceId == type(IFastTransferPool).interfaceId || interfaceId == type(IERC165).interfaceId
      || interfaceId == type(IAny2EVMMessageReceiver).interfaceId || super.supportsInterface(interfaceId);
  }

  function _burn(
    uint256 amount
  ) internal virtual override {
    i_token.safeTransferFrom(msg.sender, address(this), amount);
    IBurnMintERC20(address(i_token)).burn(amount);
  }

  function getRouter() public view override(TokenPool, FastTransferTokenPoolAbstract) returns (address router) {
    return address(s_router);
  }
}
