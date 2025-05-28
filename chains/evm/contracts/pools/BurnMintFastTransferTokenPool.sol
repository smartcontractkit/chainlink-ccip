// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.10;

import {IAny2EVMMessageReceiver} from "../interfaces/IAny2EVMMessageReceiver.sol";
import {IFastTransferPool} from "../interfaces/IFastTransferPool.sol";
import {IRMN} from "../interfaces/IRMN.sol";

import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";
import {IBurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/IBurnMintERC20.sol";
import {IERC165} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/utils/introspection/IERC165.sol";

import {CCIPReceiver} from "../applications/CCIPReceiver.sol";
import {BurnMintTokenPoolAbstract} from "./BurnMintTokenPoolAbstract.sol";
import {FastTransferTokenPoolAbstract} from "./FastTransferTokenPoolAbstract.sol";
import {TokenPool} from "./TokenPool.sol";

// OpenZeppelin imports
import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

contract BurnMintFastTransferTokenPool is ITypeAndVersion, BurnMintTokenPoolAbstract, FastTransferTokenPoolAbstract {
  using SafeERC20 for IERC20;

  string public constant override typeAndVersion = "BurnMintFastTransferTokenPool 1.6.1";

  constructor(
    IBurnMintERC20 token,
    uint8 localTokenDecimals,
    address[] memory allowlist,
    address rmnProxy,
    address router
  )
    // FastTransferTokenPoolAbstract.LaneConfigArgs[] memory laneConfigArgs
    TokenPool(token, localTokenDecimals, allowlist, rmnProxy, router)
    FastTransferTokenPoolAbstract(router)
  {}

  function _handleTokenToTransfer(uint64 destinationChainSelector, address sender, uint256 amount) internal override {
    if (IRMN(i_rmnProxy).isCursed(bytes16(uint128(destinationChainSelector)))) revert CursedByRMN();
    _checkAllowList(sender);
    if (!isSupportedChain(destinationChainSelector)) revert ChainNotAllowed(destinationChainSelector);
    _consumeOutboundRateLimit(destinationChainSelector, amount);
    _burn(amount);
  }

  function _transferFromFiller(
    uint64 sourceChainSelector,
    address filler,
    address receiver,
    uint256 srcAmount,
    uint8 srcDecimals
  ) internal override returns (uint256 localAmount) {
    localAmount = _calculateLocalAmount(srcAmount, srcDecimals);
    _consumeInboundRateLimit(sourceChainSelector, localAmount);
    getToken().safeTransferFrom(filler, receiver, localAmount);
    return localAmount;
  }

  function _settle(
    uint64 sourceChainSelector,
    bytes32 fillRequestId,
    bytes memory sourcePoolAddress,
    uint256 srcAmount,
    uint8 srcDecimal,
    uint256 fastTransferFee,
    address receiver
  ) internal override {
    if (IRMN(i_rmnProxy).isCursed(bytes16(uint128(sourceChainSelector)))) revert CursedByRMN();
    //Validates that the source pool address is configured on this pool.
    if (!isRemotePool(sourceChainSelector, sourcePoolAddress)) {
      revert InvalidSourcePoolAddress(sourcePoolAddress);
    }
    uint256 localAmount = _calculateLocalAmount(srcAmount, srcDecimal);
    uint256 settlementAmountLocal = localAmount + _calculateLocalAmount(fastTransferFee, srcDecimal);

    bytes32 fillId = keccak256(abi.encodePacked(fillRequestId, localAmount, receiver));
    address filler = s_fills[fillId];
    // not fast-filled
    if (filler == address(0)) {
      _consumeInboundRateLimit(sourceChainSelector, settlementAmountLocal);
      IBurnMintERC20(address(i_token)).mint(receiver, settlementAmountLocal);
    }
    // already settled
    else if (filler == address(1)) {
      revert MessageAlreadySettled(fillRequestId);
    }
    // fast-filled; verify amount
    else {
      // Honest filler -> pay them back + fee
      IBurnMintERC20(address(i_token)).mint(filler, settlementAmountLocal);
    }
    s_fills[fillId] = address(1); // Mark as settled
  }

  function _checkAdmin() internal view override onlyOwner {}

  /// @notice Signals which version of the pool interface is supported
  function supportsInterface(
    bytes4 interfaceId
  ) public pure virtual override(TokenPool, CCIPReceiver) returns (bool) {
    return interfaceId == type(IFastTransferPool).interfaceId || interfaceId == type(IERC165).interfaceId
      || interfaceId == type(IAny2EVMMessageReceiver).interfaceId;
  }

  function _burn(
    uint256 amount
  ) internal virtual override {
    IBurnMintERC20(address(i_token)).burnFrom(msg.sender, amount);
  }

  function getRouter() public view override(TokenPool, CCIPReceiver) returns (address router) {
    return address(s_router);
  }
}
