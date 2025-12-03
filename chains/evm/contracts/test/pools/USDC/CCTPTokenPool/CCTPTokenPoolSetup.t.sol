// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {VersionedVerifierResolver} from "../../../../ccvs/VersionedVerifierResolver.sol";
import {CCTPTokenPool} from "../../../../pools/USDC/CCTPTokenPool.sol";
import {USDCSetup} from "../USDCSetup.t.sol";

contract CCTPTokenPoolSetup is USDCSetup {
  CCTPTokenPool internal s_cctpTokenPool;
  address internal s_rmnProxy = makeAddr("rmnProxy");

  function setUp() public virtual override {
    super.setUp();

    s_cctpTokenPool =
      new CCTPTokenPool(s_USDCToken, 6, address(0), address(s_rmnProxy), address(s_router));
  }
}
