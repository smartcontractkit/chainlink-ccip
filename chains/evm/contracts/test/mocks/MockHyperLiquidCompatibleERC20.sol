// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {HyperLiquidCompatibleERC20} from "../../tokenAdminRegistry/TokenPoolFactory/HyperLiquidCompatibleERC20.sol";

contract MockHyperLiquidCompatibleERC20 is HyperLiquidCompatibleERC20 {
  constructor(
    string memory name,
    string memory symbol,
    uint8 decimals,
    uint256 maxSupply,
    uint256 preMint,
    address newOwner
  ) HyperLiquidCompatibleERC20(name, symbol, decimals, maxSupply, preMint, newOwner) {}

  /// @notice Exposes the internal _calculateLocalAmount function for testing
  /// @param remoteAmount The amount on the remote chain
  /// @param remoteDecimals The decimals of the token on the remote chain
  /// @return The local amount
  function calculateLocalAmount(uint256 remoteAmount, uint8 remoteDecimals) public view returns (uint256) {
    return _calculateLocalAmount(remoteAmount, remoteDecimals);
  }
}
