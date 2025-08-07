// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SiloedUSDCTokenPool} from "../../../../pools/USDC/SiloedUSDCTokenPool.sol";
import {SiloedUSDCTokenPoolSetup} from "./SiloedUSDCTokenPoolSetup.sol";

contract SiloedUSDCTokenPool_excludeTokensFromBurn is SiloedUSDCTokenPoolSetup {
  function test_excludeTokensFromBurn_RevertWhen_NoMigrationProposalPending() public {
    // Set up silo designation for the test chain
    vm.expectRevert(abi.encodeWithSelector(SiloedUSDCTokenPool.NoMigrationProposalPending.selector));

    s_usdcTokenPool.excludeTokensFromBurn(SOURCE_CHAIN_SELECTOR, 1e6);
    vm.stopPrank();
  }
}
