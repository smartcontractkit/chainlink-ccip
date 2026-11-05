// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RateLimiter} from "../../../libraries/RateLimiter.sol";
import {TokenPoolRateLimitHelper} from "../../../pools/TokenPoolRateLimitHelper.sol";
import {TokenPoolRateLimitHelperSetup} from "./TokenPoolRateLimitHelperSetup.t.sol";

contract TokenPoolRateLimitHelper_applyDirectionalUpdates is TokenPoolRateLimitHelperSetup {
  function test_applyDirectionalUpdates_SameLaneBothDirectionsCompose() public {
    TokenPoolRateLimitHelper.DirectionalUpdate[] memory updates = new TokenPoolRateLimitHelper.DirectionalUpdate[](3);
    updates[0] = TokenPoolRateLimitHelper.DirectionalUpdate({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      fastFinality: false,
      direction: TokenPoolRateLimitHelper.Direction.Outbound,
      config: _helperCfg(true, 11e18, 1e18)
    });
    updates[1] = TokenPoolRateLimitHelper.DirectionalUpdate({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      fastFinality: false,
      direction: TokenPoolRateLimitHelper.Direction.Inbound,
      config: _helperCfg(true, 22e18, 2e18)
    });
    updates[2] = TokenPoolRateLimitHelper.DirectionalUpdate({
      remoteChainSelector: OTHER_CHAIN_SELECTOR,
      fastFinality: false,
      direction: TokenPoolRateLimitHelper.Direction.Inbound,
      config: _helperCfg(true, 33e18, 3e18)
    });

    vm.startPrank(s_manager);
    s_helper.applyDirectionalUpdates(updates);

    (RateLimiter.TokenBucket memory outbound, RateLimiter.TokenBucket memory inbound) =
      _state(DEST_CHAIN_SELECTOR, false);
    assertEq(outbound.capacity, 11e18);
    assertEq(inbound.capacity, 22e18);

    (outbound, inbound) = _state(OTHER_CHAIN_SELECTOR, false);
    assertEq(outbound.capacity, 500e18);
    assertEq(inbound.capacity, 33e18);
  }

  function test_RevertWhen_CallerIsNotOwnerOrManager() public {
    vm.startPrank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(TokenPoolRateLimitHelper.Unauthorized.selector, STRANGER));
    s_helper.applyDirectionalUpdates(new TokenPoolRateLimitHelper.DirectionalUpdate[](0));
  }
}
