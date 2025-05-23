pragma solidity ^0.8.24;

import {IFastTransferPool} from "../interfaces/IFastTransferPool.sol";
import {IRMN} from "../interfaces/IRMN.sol";

import {BurnMintTokenPoolAbstract} from "./BurnMintTokenPoolAbstract.sol";
import {TokenPool} from "./TokenPool.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";
import {IBurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/IBurnMintERC20.sol";

import {IERC165} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/utils/introspection/IERC165.sol";

import {FastTransferTokenPoolAbstract} from "./FastTransferTokenPoolAbstract.sol";

import {CCIPReceiver} from "../applications/CCIPReceiver.sol";
import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

contract BurnMintFastTransferTokenPool is ITypeAndVersion, BurnMintTokenPoolAbstract, FastTransferTokenPoolAbstract {
  using SafeERC20 for IERC20;

  string public constant override typeAndVersion = "BurnMintFastTransferTokenPool 1.5.1";

  constructor(
    IBurnMintERC20 token,
    uint8 localTokenDecimals,
    address[] memory allowlist,
    address rmnProxy,
    address router
  ) TokenPool(token, localTokenDecimals, allowlist, rmnProxy, router) CCIPReceiver(router) {}

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
  ) internal override returns (uint256 destAmount) {
    uint256 localAmount = _calculateLocalAmount(srcAmount, srcDecimals);
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
    // Validates that the source pool address is configured on this pool.
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
  }

  function _checkAdmin() internal view override onlyOwner {}
  /// @notice Signals which version of the pool interface is supported

  function supportsInterface(
    bytes4 interfaceId
  ) public pure virtual override(TokenPool, CCIPReceiver) returns (bool) {
    return interfaceId == type(IFastTransferPool).interfaceId || interfaceId == type(IERC165).interfaceId;
  }

  function _burn(
    uint256 amount
  ) internal virtual override {
    IBurnMintERC20(address(i_token)).burnFrom(address(this), amount);
  }

  function getRouter() public view override(TokenPool, CCIPReceiver) returns (address router) {
    return address(s_router);
  }
}
