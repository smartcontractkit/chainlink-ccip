  // SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FastTransferTokenPoolAbstract} from "../../pools/FastTransferTokenPoolAbstract.sol";
import {TokenPool} from "../../pools/TokenPool.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

import {WETH9} from "@chainlink/contracts/src/v0.8/vendor/canonical-weth/WETH9.sol";
import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

contract FastTransferTokenPoolHelper is FastTransferTokenPoolAbstract {
  using SafeERC20 for IERC20;

  error NotAdmin();

  string public constant override typeAndVersion = "FastTransferTokenPoolHelper 1.6.1";

  IERC20 internal immutable i_token;
  address internal immutable i_admin;

  constructor(
    IERC20 token,
    WETH9 wrappedNative,
    address router,
    FastTransferTokenPoolAbstract.LaneConfigArgs[] memory laneConfigArgs
  ) FastTransferTokenPoolAbstract(router, address(wrappedNative), laneConfigArgs) {
    i_admin = msg.sender;
    i_token = token;
  }

  // Public wrappers for internal functions
  function handleTokenToTransfer(uint64 destinationChainSelector, address sender, uint256 amount) public {
    _handleTokenToTransfer(destinationChainSelector, sender, amount);
  }

  function transferFromFiller(
    uint64 sourceChainSelector,
    address filler,
    address receiver,
    uint256 srcAmount,
    uint8 srcDecimals
  ) public returns (uint256) {
    return _transferFromFiller(sourceChainSelector, filler, receiver, srcAmount, srcDecimals);
  }

  function settle(
    uint64 sourceChainSelector,
    bytes32 fillRequestId,
    bytes memory sourcePoolAddress,
    uint256 srcAmount,
    uint8 srcDecimal,
    uint256 fastTransferFee,
    address receiver
  ) public {
    _settle(sourceChainSelector, fillRequestId, sourcePoolAddress, srcAmount, srcDecimal, fastTransferFee, receiver);
  }

  // Implementation of abstract functions
  function _handleTokenToTransfer(uint64 destinationChainSelector, address sender, uint256 amount) internal override {
    // For testing, we'll just transfer tokens from sender to this contract
    IERC20(i_token).safeTransferFrom(sender, address(this), amount);
  }

  function _transferFromFiller(
    uint64 sourceChainSelector,
    address filler,
    address receiver,
    uint256 srcAmount,
    uint8 srcDecimals
  ) internal override returns (uint256) {
    IERC20(i_token).safeTransferFrom(filler, receiver, srcAmount);
    return srcAmount;
  }

  function _settle(
    uint64 sourceChainSelector,
    bytes32 fillRequestId,
    bytes memory sourcePoolAddress,
    uint256 srcAmount,
    uint8 srcDecimal,
    uint256 fastTransferFee,
    address receiver
  ) internal override {}

  function _checkAdmin() internal view override {
    if (msg.sender != i_admin) revert NotAdmin();
  }
}
