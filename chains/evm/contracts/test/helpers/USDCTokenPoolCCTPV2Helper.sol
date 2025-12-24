// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITokenMessenger} from "../../pools/USDC/interfaces/ITokenMessenger.sol";

import {USDCSourcePoolDataCodec} from "../../libraries/USDCSourcePoolDataCodec.sol";
import {CCTPMessageTransmitterProxy} from "../../pools/USDC/CCTPMessageTransmitterProxy.sol";
import {USDCTokenPoolCCTPV2} from "../../pools/USDC/USDCTokenPoolCCTPV2.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract USDCTokenPoolCCTPV2Helper is USDCTokenPoolCCTPV2 {
  constructor(
    ITokenMessenger tokenMessenger,
    CCTPMessageTransmitterProxy messageTransmitterProxy,
    IERC20 token,
    address advancedPoolHooks,
    address rmnProxy,
    address router
  ) USDCTokenPoolCCTPV2(tokenMessenger, messageTransmitterProxy, token, advancedPoolHooks, rmnProxy, router) {}

  function validateMessage(
    bytes memory usdcMessage,
    USDCSourcePoolDataCodec.SourceTokenDataPayloadV2 memory sourceTokenData
  ) external view {
    return _validateMessage(usdcMessage, sourceTokenData);
  }
}
