// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV1} from "../../../../interfaces/IPool.sol";

import {Pool} from "../../../../libraries/Pool.sol";
import {USDCTokenPoolProxy} from "../../../../pools/USDC/USDCTokenPoolProxy.sol";
import {USDCSetup} from "../USDCSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {IERC165} from "@openzeppelin/contracts@5.0.2/utils/introspection/IERC165.sol";

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
      }),
      address(s_router)
    );

    USDCTokenPoolProxy.PoolAddresses memory pools = proxy.getPools();
    assertEq(pools.legacyCctpV1Pool, s_legacyCctpV1Pool);
    assertEq(pools.cctpV1Pool, s_cctpV1Pool);
    assertEq(pools.cctpV2Pool, s_cctpV2Pool);

    assertTrue(proxy.supportsInterface(type(IPoolV1).interfaceId));
    assertTrue(proxy.supportsInterface(Pool.CCIP_POOL_V1));
    assertTrue(proxy.supportsInterface(type(IERC165).interfaceId));

    assertTrue(proxy.isSupportedToken(address(s_USDCToken)));
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
      }),
      address(s_router)
    );
  }

  function test_constructor_RevertWhen_RouterAddressIsZero() public {
    vm.expectRevert(USDCTokenPoolProxy.AddressCannotBeZero.selector);
    new USDCTokenPoolProxy(
      IERC20(s_USDCToken), // Token
      USDCTokenPoolProxy.PoolAddresses({
        legacyCctpV1Pool: s_legacyCctpV1Pool,
        cctpV1Pool: s_cctpV1Pool,
        cctpV2Pool: address(0)
      }),
      address(0) // Router
    );
  }
}
