// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {VersionedVerifierResolver} from "../../../../ccvs/VersionedVerifierResolver.sol";
import {CCTPTokenPool} from "../../../../pools/USDC/CCTPTokenPool.sol";
import {USDCSetup} from "../USDCSetup.t.sol";

contract CCTPTokenPoolSetup is USDCSetup {
  CCTPTokenPool internal s_cctpTokenPool;
  address internal ALLOWED_CALLER = makeAddr("allowedCaller");

  function setUp() public virtual override {
    super.setUp();

    address[] memory allowedCallers = new address[](1);
    allowedCallers[0] = ALLOWED_CALLER;

    s_cctpTokenPool =
      new CCTPTokenPool(s_USDCToken, 6, address(0), address(s_mockRMNRemote), address(s_router), allowedCallers);

    _poolApplyChainUpdates(address(s_cctpTokenPool));
  }
}
