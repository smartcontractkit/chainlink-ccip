// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IBridgeV2} from "../../interfaces/lombard/IBridgeV2.sol";

import {LombardTokenPool} from "../../pools/Lombard/LombardTokenPool.sol";
import {IERC20Metadata} from "@openzeppelin/contracts@4.8.3/token/ERC20/extensions/IERC20Metadata.sol";

contract LombardTokenPoolHelper is LombardTokenPool {
  constructor(
    IERC20Metadata token,
    address verifier,
    IBridgeV2 bridge,
    address adapter,
    address advancedPoolHooks,
    address rmnProxy,
    address router,
    uint8 fallbackDecimals,
    address feeAggregator
  )
    LombardTokenPool(
      token, verifier, bridge, adapter, advancedPoolHooks, rmnProxy, router, fallbackDecimals, feeAggregator
    )
  {}

  function getTokenDecimals(
    IERC20Metadata token,
    uint8 fallbackDecimals
  ) external view returns (uint8) {
    return _getTokenDecimals(token, fallbackDecimals);
  }
}
