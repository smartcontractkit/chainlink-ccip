// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {LombardTokenPool} from "../../pools/Lombard/LombardTokenPool.sol";
import {IERC20Metadata} from "@openzeppelin/contracts@4.8.3/token/ERC20/extensions/IERC20Metadata.sol";

/// @dev Helper exposing internal view for testing only.
contract LombardTokenPoolHelper is LombardTokenPool {
  constructor(
    IERC20Metadata token,
    address verifier,
    address rmnProxy,
    address router,
    uint8 fallbackDecimals
  ) LombardTokenPool(token, verifier, address(0), rmnProxy, router, fallbackDecimals) {}

  function getTokenDecimals(IERC20Metadata token, uint8 fallbackDecimals) external view returns (uint8) {
    return _getTokenDecimals(token, fallbackDecimals);
  }
}
