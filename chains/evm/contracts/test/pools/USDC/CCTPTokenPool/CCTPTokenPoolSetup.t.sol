// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCTPTokenPool} from "../../../../pools/USDC/CCTPTokenPool.sol";
import {USDCSetup} from "../USDCSetup.t.sol";
import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract CCTPTokenPoolSetup is USDCSetup {
  CCTPTokenPool internal s_cctpTokenPool;
  address internal s_cctpVerifier = makeAddr("cctpVerifier");
  address internal s_allowedCaller = makeAddr("allowedCaller");

  function setUp() public virtual override {
    super.setUp();

    address[] memory allowedCallers = new address[](1);
    allowedCallers[0] = s_allowedCaller;

    s_cctpTokenPool =
      new CCTPTokenPool(s_USDCToken, 6, address(s_mockRMNRemote), address(s_router), s_cctpVerifier, allowedCallers);

    _poolApplyChainUpdates(address(s_cctpTokenPool));
  }
}
