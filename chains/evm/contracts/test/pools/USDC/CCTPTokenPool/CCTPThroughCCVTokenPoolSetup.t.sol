// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCTPThroughCCVTokenPool} from "../../../../pools/USDC/CCTPThroughCCVTokenPool.sol";
import {USDCSetup} from "../USDCSetup.t.sol";

contract CCTPThroughCCVTokenPoolSetup is USDCSetup {
  CCTPThroughCCVTokenPool internal s_cctpThroughCCVTokenPool;
  address internal s_cctpVerifier = makeAddr("cctpVerifier");
  address internal s_allowedCaller = makeAddr("allowedCaller");

  function setUp() public virtual override {
    super.setUp();

    address[] memory allowedCallers = new address[](1);
    allowedCallers[0] = s_allowedCaller;

    s_cctpThroughCCVTokenPool = new CCTPThroughCCVTokenPool(
      s_USDCToken, 6, address(s_mockRMNRemote), address(s_router), s_cctpVerifier, allowedCallers
    );

    _poolApplyChainUpdates(address(s_cctpThroughCCVTokenPool));
  }
}

