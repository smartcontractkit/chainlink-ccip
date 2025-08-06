// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITokenMessenger} from "../../pools/USDC/interfaces/ITokenMessenger.sol";
import {IBurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/IBurnMintERC20.sol";

import {CCTPMessageTransmitterProxy} from "../../pools/USDC/CCTPMessageTransmitterProxy.sol";
import {USDCTokenPoolCCTPV2} from "../../pools/USDC/USDCTokenPoolCCTPV2.sol";

contract USDCTokenPoolCCTPV2Helper is USDCTokenPoolCCTPV2 {
  constructor(
    ITokenMessenger tokenMessenger,
    CCTPMessageTransmitterProxy messageTransmitterProxy,
    IBurnMintERC20 token,
    address[] memory allowlist,
    address rmnProxy,
    address router
  ) USDCTokenPoolCCTPV2(tokenMessenger, messageTransmitterProxy, token, allowlist, rmnProxy, router) {}

  function validateMessage(bytes memory usdcMessage, SourceTokenDataPayload memory sourceTokenData) external view {
    return _validateMessage(usdcMessage, sourceTokenData);
  }
}
