// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {FastTransferTokenPoolAbstract} from "../../pools/FastTransferTokenPoolAbstract.sol";
import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

contract FastTransferTokenPoolHelper is FastTransferTokenPoolAbstract {
  using SafeERC20 for IERC20;

  string public constant override typeAndVersion = "FastTransferTokenPoolHelper 1.6.1";

  constructor(
    IERC20 token,
    uint8 localTokenDecimals,
    address[] memory allowlist,
    address rmnProxy,
    address router
  ) FastTransferTokenPoolAbstract(token, localTokenDecimals, allowlist, rmnProxy, router) {}

  // Public wrappers for internal functions
  function handleFastTransferLockOrBurn(address sender, uint256 amount) public {
    _handleFastTransferLockOrBurn(sender, amount);
  }

  function transferFromFiller(address filler, address receiver, uint256 amount) public {
    _transferFromFiller(filler, receiver, amount);
  }

  // Implementation of abstract functions
  function _handleFastTransferLockOrBurn(address sender, uint256 amount) internal override {
    // For testing, we'll just transfer tokens from sender to this contract
    getToken().safeTransferFrom(sender, address(this), amount);
  }

  /// @notice Validates settlement prerequisites - simple implementation for testing
  function _validateSettlement(uint64, bytes memory) internal view override {
    // For testing, we'll do minimal validation
    // Real implementations would check RMN curse and source pool validation
  }

  function _releaseOrMint(address receiver, uint256 amount) internal virtual override {
    getToken().safeTransfer(receiver, amount);
  }
}
