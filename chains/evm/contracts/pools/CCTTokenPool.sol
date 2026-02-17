// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IBurnMintERC20} from "../interfaces/IBurnMintERC20.sol";
import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";

import {TokenPool} from "./TokenPool.sol";

import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";
import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract CCTTokenPool is TokenPool, BurnMintERC20, ITypeAndVersion {
  function typeAndVersion() external pure virtual override returns (string memory) {
    return "CCTTokenPool 2.0.0-dev";
  }

  error OnlyBridgeCanMint();

  constructor(
    string memory name,
    string memory symbol,
    uint8 decimals_,
    uint256 maxSupply_,
    uint256 preMint,
    address advancedPoolHooks,
    address rmnProxy,
    address router
  )
    BurnMintERC20(name, symbol, decimals_, maxSupply_, preMint)
    TokenPool(IERC20(address(this)), decimals_, advancedPoolHooks, rmnProxy, router)
  {}

  /// @notice Burns tokens held by the pool.
  function _lockOrBurn(
    uint64, // remoteChainSelector
    uint256 amount
  ) internal virtual override {
    _burn(_msgSender(), amount);
  }

  function _releaseOrMint(
    address receiver,
    uint256 amount,
    uint64 // remoteChainSelector
  ) internal virtual override {
    if (receiver == address(this)) revert InvalidRecipient(receiver);
    if (i_maxSupply != 0 && totalSupply() + amount > i_maxSupply) revert MaxSupplyExceeded(totalSupply() + amount);

    _mint(receiver, amount);
  }

  function mint(
    address,
    uint256
  ) external virtual override {
    revert OnlyBridgeCanMint();
  }

  /// @notice Signals which version of the pool interface is supported.
  /// @param interfaceId The interface identifier, as specified in ERC-165.
  function supportsInterface(
    bytes4 interfaceId
  ) public pure virtual override(BurnMintERC20, TokenPool) returns (bool) {
    return BurnMintERC20.supportsInterface(interfaceId) || TokenPool.supportsInterface(interfaceId);
  }
}
