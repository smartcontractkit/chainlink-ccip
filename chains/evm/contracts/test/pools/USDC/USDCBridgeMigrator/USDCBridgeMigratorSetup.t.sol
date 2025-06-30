// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {HybridLockReleaseUSDCTokenPool} from "../../../../pools/USDC/HybridLockReleaseUSDCTokenPool.sol";
import {USDCSetup} from "../USDCSetup.t.sol";

contract USDCBridgeMigratorSetup is USDCSetup {
  HybridLockReleaseUSDCTokenPool internal s_usdcTokenPool;
  HybridLockReleaseUSDCTokenPool internal s_usdcTokenPoolTransferLiquidity;

  function setUp() public virtual override {
    super.setUp();

    s_usdcTokenPool = new HybridLockReleaseUSDCTokenPool(
      s_mockUSDC,
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router),
      s_previousPool
    );

    s_usdcTokenPoolTransferLiquidity = new HybridLockReleaseUSDCTokenPool(
      s_mockUSDC,
      s_cctpMessageTransmitterProxy,
      s_USDCToken,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_router),
      s_previousPool
    );
  }
}
