// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseERC20} from "../tmp/BaseERC20.sol";
import {TokenPool} from "./TokenPool.sol";

import {ERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/ERC20.sol";
import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract CCTTokenPool is TokenPool, BaseERC20 {
  function typeAndVersion() external pure virtual override returns (string memory) {
    return "CCTTokenPool 2.0.0-dev";
  }

  constructor(
    BaseERC20.ConstructorParams memory tokenParams,
    address advancedPoolHooks,
    address rmnProxy,
    address router
  )
    BaseERC20(tokenParams)
    TokenPool(IERC20(address(this)), tokenParams.decimals, advancedPoolHooks, rmnProxy, router)
  {}

  /// @notice Burns tokens held by the pool. The Router transfers tokens to
  /// this contract before the OnRamp calls lockOrBurn, so the burn is from self.
  function _lockOrBurn(
    uint64, // remoteChainSelector
    uint256 amount
  ) internal virtual override {
    _burn(address(this), amount);
  }

  function _releaseOrMint(
    address receiver,
    uint256 amount,
    uint64 // remoteChainSelector
  ) internal virtual override {
    if (i_maxSupply != 0 && totalSupply() + amount > i_maxSupply) revert MaxSupplyExceeded(totalSupply() + amount);

    _mint(receiver, amount);
  }

  /// @dev Overrides BaseERC20._update to allow this contract to receive its own tokens.
  /// The CCIP Router transfers tokens to the pool (which IS this contract) before
  /// lockOrBurn is called, so transfers to address(this) must be permitted.
  function _update(
    address from,
    address to,
    uint256 value
  ) internal virtual override {
    ERC20._update(from, to, value);
  }

  /// @notice Signals which version of the pool interface is supported.
  /// @param interfaceId The interface identifier, as specified in ERC-165.
  function supportsInterface(
    bytes4 interfaceId
  ) public view virtual override(BaseERC20, TokenPool) returns (bool) {
    return BaseERC20.supportsInterface(interfaceId) || TokenPool.supportsInterface(interfaceId);
  }

  /// @notice Overrides the default CCIP admin role setter to require the caller to be the owner.
  /// @param newAdmin The address of the new CCIP admin.
  function setCCIPAdmin(
    address newAdmin
  ) external virtual override onlyOwner {
    _setCCIPAdmin(newAdmin);
  }
}
