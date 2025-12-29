// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV1} from "../../../../interfaces/IPool.sol";
import {IPoolV2} from "../../../../interfaces/IPoolV2.sol";
import {USDCTokenPoolProxy} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";
import {USDCTokenPoolProxySetup} from "./USDCTokenPoolProxySetup.t.sol";

contract USDCTokenPoolProxy_getFee is USDCTokenPoolProxySetup {
  uint256 internal constant FEE_USD_CENTS = 1e18;
  uint32 internal constant DEST_GAS_OVERHEAD = 200_000;
  uint32 internal constant DEST_BYTES_OVERHEAD = 32;
  uint16 internal constant TOKEN_FEE_BPS = 5;
  bool internal constant IS_ENABLED = true;

  function test_getFee() public {
    vm.mockCall(
      address(s_cctpTokenPool),
      abi.encodeWithSelector(IPoolV2.getFee.selector, address(0), 0, 0, address(0), 0, ""),
      abi.encode(FEE_USD_CENTS, DEST_GAS_OVERHEAD, DEST_BYTES_OVERHEAD, TOKEN_FEE_BPS, IS_ENABLED)
    );

    (uint256 usdFeeCents, uint32 destGasOverhead, uint32 destBytesOverhead, uint16 tokenFeeBps, bool isEnabled) =
      s_usdcTokenPoolProxy.getFee(address(0), 0, 0, address(0), 0, "");

    assertEq(usdFeeCents, FEE_USD_CENTS);
    assertEq(destGasOverhead, DEST_GAS_OVERHEAD);
    assertEq(destBytesOverhead, DEST_BYTES_OVERHEAD);
    assertEq(tokenFeeBps, TOKEN_FEE_BPS);
    assertEq(isEnabled, IS_ENABLED);
  }

  function test_getFee_RevertWhen_NoCCVCompatiblePoolSet() public {
    _enableERC165InterfaceChecks(s_cctpV2Pool, type(IPoolV1).interfaceId);
    _enableERC165InterfaceChecks(s_cctpV1Pool, type(IPoolV1).interfaceId);
    _enableERC165InterfaceChecks(s_legacyCctpV1Pool, type(IPoolV1).interfaceId);
    s_usdcTokenPoolProxy.updatePoolAddresses(
      USDCTokenPoolProxy.PoolAddresses({
        legacyCctpV1Pool: s_legacyCctpV1Pool,
        cctpV1Pool: s_cctpV1Pool,
        cctpV2Pool: s_cctpV2Pool,
        cctpTokenPool: address(0)
      })
    );

    vm.expectRevert(abi.encodeWithSelector(USDCTokenPoolProxy.CCVCompatiblePoolNotSet.selector));
    s_usdcTokenPoolProxy.getFee(address(0), 0, 0, address(0), 0, "");
  }
}
