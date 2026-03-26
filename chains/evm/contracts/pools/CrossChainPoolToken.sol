// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseERC20} from "../tokens/BaseERC20.sol";
import {TokenPool} from "./TokenPool.sol";

import {ERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/ERC20.sol";
import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

/// @notice A CCIP token pool that is also an ERC20 token. This allows the pool to burn/mint without needing to manage
/// roles for a separate token contract.
/// @dev This contract inherits its access control from TokenPool, meaning it uses an `owner` role with 2-step ownership
/// transfers. There's also a separate `ccipAdmin` role which can be used to register with the CCIP token admin registry
/// but has no other special powers, and can only be transferred by the owner. The owner role can also be used to
/// register the token in the token admin registry.
contract CrossChainPoolToken is TokenPool, BaseERC20 {
  function typeAndVersion() external pure virtual override returns (string memory) {
    return "CrossChainPoolToken 2.0.0-dev";
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
  /// @param amount The amount of tokens to burn.
  function _lockOrBurn(
    uint64, // remoteChainSelector
    uint256 amount
  ) internal virtual override {
    _burn(address(this), amount);
  }

  /// @notice Mints tokens to the receiver.
  /// @param receiver The address to mint tokens to.
  /// @param amount The amount of tokens to mint.
  function _releaseOrMint(
    address receiver,
    uint256 amount,
    uint64 // remoteChainSelector
  ) internal virtual override {
    _mint(receiver, amount);
  }

  /// @dev Overrides BaseERC20._update to allow this contract to receive its own tokens.
  /// The CCIP Router transfers tokens to the pool (which IS this contract) before
  /// lockOrBurn is called, so transfers to address(this) must be permitted.
  /// @dev This function must reflect any changes made in BaseERC20._update, which it currently does by adding the
  /// supply check.
  function _update(
    address from,
    address to,
    uint256 value
  ) internal virtual override {
    // Update first, then check the total supply.
    ERC20._update(from, to, value);

    // If `from` is address(0), this is a mint, so we need to check the total supply against the max supply.
    if (from == address(0)) {
      _assertMaxSupply();
    }
  }

  /// @notice Signals which version of the pool interface is supported.
  /// @param interfaceId The interface identifier, as specified in ERC-165.
  /// @return True if the contract implements the requested interface, false otherwise.
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
