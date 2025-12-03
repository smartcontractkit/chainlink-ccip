// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../libraries/Pool.sol";

import {LombardTokenPool} from "../../../pools/Lombard/LombardTokenPool.sol";
import {LombardTokenPoolSetup} from "./LombardTokenPoolSetup.t.sol";

contract LombardTokenPool_lockOrBurn is LombardTokenPoolSetup {
  function setUp() public virtual override {
    super.setUp();
    vm.startPrank(s_allowedOnRamp);
  }

  function test_lockOrBurn_ForwardsToVerifier() public {
    uint256 amount = 1e18;
    deal(address(s_token), address(s_pool), amount);

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
    assertEq(s_token.balanceOf(s_verifierResolver.getOutboundImplementation(DEST_CHAIN_SELECTOR, "")), amount);
    assertEq(s_token.balanceOf(address(s_pool)), 0);
  }

  function test_lockOrBurn_RevertWhen_OutboundImplementationNotFoundForVerifier() public {
    uint256 amount = 1e18;
    deal(address(s_token), address(s_pool), amount);
    vm.mockCall(
      address(s_verifierResolver),
      abi.encodeCall(s_verifierResolver.getOutboundImplementation, (DEST_CHAIN_SELECTOR, "")),
      abi.encode(address(0))
    );

    vm.expectRevert(LombardTokenPool.OutboundImplementationNotFoundForVerifier.selector);
    s_pool.lockOrBurn(
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
  }
}
