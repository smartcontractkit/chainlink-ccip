// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {ITokenMinter} from "../../pools/USDC/interfaces/ITokenMinter.sol";

contract MockTokenMinter is ITokenMinter {
  mapping(uint32 remoteDomain => mapping(bytes32 remoteToken => address localToken)) internal s_localTokens;

  function setLocalToken(
    uint32 remoteDomain,
    bytes32 remoteToken,
    address localToken
  ) external {
    s_localTokens[remoteDomain][remoteToken] = localToken;
  }

  function getLocalToken(
    uint32 remoteDomain,
    bytes32 remoteToken
  ) external view returns (address) {
    return s_localTokens[remoteDomain][remoteToken];
  }
}
