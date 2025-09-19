// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {USDCTokenPoolProxy} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";
import {USDCSetup} from "../USDCSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract USDCTokenPoolProxy_constructor is USDCSetup {
  address internal s_legacyCctpV1Pool = makeAddr("legacyCctpV1Pool");
  address internal s_cctpV1Pool = makeAddr("cctpV1Pool");
  address internal s_cctpV2Pool = makeAddr("cctpV2Pool");
  address internal s_lockReleasePool = makeAddr("lockReleasePool");

  function test_constructor() public {
    // Arrange: Define test constants
    USDCTokenPoolProxy proxy = new USDCTokenPoolProxy(
      s_USDCToken,
      USDCTokenPoolProxy.PoolAddresses({
        legacyCctpV1Pool: s_legacyCctpV1Pool,
        cctpV1Pool: s_cctpV1Pool,
        cctpV2Pool: s_cctpV2Pool
      })
    );

    USDCTokenPoolProxy.PoolAddresses memory pools = proxy.getPools();
    assertEq(pools.legacyCctpV1Pool, s_legacyCctpV1Pool);
    assertEq(pools.cctpV1Pool, s_cctpV1Pool);
    assertEq(pools.cctpV2Pool, s_cctpV2Pool);
  }

  // Reverts
  function test_constructor_RevertWhen_TokenAddressIsZero() public {
    vm.expectRevert(USDCTokenPoolProxy.AddressCannotBeZero.selector);
    new USDCTokenPoolProxy(
      IERC20(address(0)),
      USDCTokenPoolProxy.PoolAddresses({
        legacyCctpV1Pool: s_legacyCctpV1Pool,
        cctpV1Pool: s_cctpV1Pool,
        cctpV2Pool: address(0)
      })
    );
  }

  function test_constructor_RevertWhen_CCTPV2PoolIsZero() public {
    vm.expectRevert(USDCTokenPoolProxy.AddressCannotBeZero.selector);
    new USDCTokenPoolProxy(
      s_USDCToken,
      USDCTokenPoolProxy.PoolAddresses({
        legacyCctpV1Pool: s_legacyCctpV1Pool,
        cctpV1Pool: s_cctpV1Pool,
        cctpV2Pool: address(0)
      })
    );
  }
}
