// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../libraries/Pool.sol";
import {LombardTokenPoolSetup} from "./LombardTokenPoolSetup.t.sol";

contract LombardTokenPool_lockOrBurn is LombardTokenPoolSetup {
  function test_lockOrBurn_ForwardsToVerifier() public {
    uint256 amount = 1e18;
    deal(address(s_token), address(s_pool), amount);

    vm.startPrank(s_allowedOnRamp);
    (Pool.LockOrBurnOutV1 memory out, uint256 destAmount) = s_pool.lockOrBurn(
      Pool.LockOrBurnInV1({
        receiver: abi.encodePacked(address(0xDEAD)),
        remoteChainSelector: DEST_CHAIN_SELECTOR,
        originalSender: OWNER,
        amount: amount,
        localToken: address(s_token)
      }),
      0,
      ""
    );

    assertEq(destAmount, amount);
    assertEq(out.destTokenAddress, abi.encode(s_remoteToken));
    assertEq(out.destPoolData, abi.encode(uint8(DEFAULT_TOKEN_DECIMALS)));
    assertEq(s_token.balanceOf(VERIFIER_IMPL), amount);
    assertEq(s_token.balanceOf(address(s_pool)), 0);
  }
}
