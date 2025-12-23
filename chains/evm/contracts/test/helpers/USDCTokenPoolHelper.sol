// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IBurnMintERC20} from "../../interfaces/IBurnMintERC20.sol";
import {ITokenMessenger} from "../../pools/USDC/interfaces/ITokenMessenger.sol";

import {USDCSourcePoolDataCodec} from "../../libraries/USDCSourcePoolDataCodec.sol";
import {CCTPMessageTransmitterProxy} from "../../pools/USDC/CCTPMessageTransmitterProxy.sol";
import {USDCTokenPool} from "../../pools/USDC/USDCTokenPool.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract USDCTokenPoolHelper is USDCTokenPool {
  constructor(
    ITokenMessenger tokenMessenger,
    CCTPMessageTransmitterProxy messageTransmitterProxy,
    IBurnMintERC20 token,
    address advancedPoolHooks,
    address rmnProxy,
    address router
  )
    USDCTokenPool(
      tokenMessenger, messageTransmitterProxy, IERC20(address(token)), advancedPoolHooks, rmnProxy, router, 0
    )
  {}

  function validateMessage(
    bytes memory usdcMessage,
    USDCSourcePoolDataCodec.SourceTokenDataPayloadV1 memory sourceTokenData
  ) external view {
    return _validateMessage(usdcMessage, sourceTokenData);
  }
}
