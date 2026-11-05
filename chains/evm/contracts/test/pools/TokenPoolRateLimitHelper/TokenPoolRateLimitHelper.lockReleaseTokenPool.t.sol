// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RateLimiter} from "../../../libraries/RateLimiter.sol";
import {ERC20LockBox} from "../../../pools/ERC20LockBox.sol";
import {LockReleaseTokenPool} from "../../../pools/LockReleaseTokenPool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenPoolRateLimitHelper} from "../../../pools/TokenPoolRateLimitHelper.sol";
import {TokenPoolRateLimitHelperSetup} from "./TokenPoolRateLimitHelperSetup.t.sol";

/// @notice Same core check against the non-siloed LockReleaseTokenPool.
contract TokenPoolRateLimitHelper_lockReleaseTokenPool is TokenPoolRateLimitHelperSetup {
  function test_lockReleaseTokenPool_SetInboundLeavesOutbound() public {
    ERC20LockBox lockBox = new ERC20LockBox(address(s_token));
    LockReleaseTokenPool pool = new LockReleaseTokenPool(
      s_token, DEFAULT_TOKEN_DECIMALS, address(0), address(s_mockRMNRemote), address(s_sourceRouter), address(lockBox)
    );

    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = abi.encode(makeAddr("remotePool"));
    TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](1);
    chainUpdates[0] = TokenPool.ChainUpdate({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encode(makeAddr("remoteToken")),
      outboundRateLimiterConfig: _cfg(true, 1000e18, 10e18),
      inboundRateLimiterConfig: _cfg(true, 2000e18, 20e18)
    });
    pool.applyChainUpdates(new uint64[](0), chainUpdates);

    TokenPoolRateLimitHelper helper = new TokenPoolRateLimitHelper(address(pool), OWNER);
    pool.setDynamicConfig(address(s_sourceRouter), address(helper), address(0));

    helper.setInboundRateLimit(DEST_CHAIN_SELECTOR, false, _helperCfg(true, 5e18, 1e18));

    (RateLimiter.TokenBucket memory outbound, RateLimiter.TokenBucket memory inbound) =
      pool.getCurrentRateLimiterState(DEST_CHAIN_SELECTOR, false);
    assertEq(outbound.capacity, 1000e18);
    assertEq(outbound.rate, 10e18);
    assertEq(inbound.capacity, 5e18);
    assertEq(inbound.rate, 1e18);
  }
}
