// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IBridgeV1} from "../../pools/Lombard/interfaces/IBridgeV1.sol";

import {LombardTokenPool} from "../../pools/Lombard/LombardTokenPool.sol";
import {IERC20Metadata} from "@openzeppelin/contracts@4.8.3/token/ERC20/extensions/IERC20Metadata.sol";

contract LombardTokenPoolHelper is LombardTokenPool {
  constructor(
    IERC20Metadata token,
    address verifier,
    IBridgeV1 bridge,
    address adapter,
    address advancedPoolHooks,
    address rmnProxy,
    address router,
    uint8 fallbackDecimals
  ) LombardTokenPool(token, verifier, bridge, adapter, advancedPoolHooks, rmnProxy, router, fallbackDecimals) {}

  function getTokenDecimals(IERC20Metadata token, uint8 fallbackDecimals) external view returns (uint8) {
    return _getTokenDecimals(token, fallbackDecimals);
  }
}
