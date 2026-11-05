// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {RateLimiter} from "../../../libraries/RateLimiter.sol";
import {SiloedLockReleaseTokenPool} from "../../../pools/SiloedLockReleaseTokenPool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {ITokenPoolV2RateLimits, TokenPoolRateLimitHelper} from "../../../pools/TokenPoolRateLimitHelper.sol";
import {BaseERC20} from "../../../tokens/BaseERC20.sol";
import {CrossChainToken} from "../../../tokens/CrossChainToken.sol";
import {BaseTest} from "../../BaseTest.t.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

/// @dev Real SiloedLockReleaseTokenPool, plus test-only hooks to drain rate limit buckets.
contract SiloedLockReleaseTokenPoolRateLimitHarness is SiloedLockReleaseTokenPool {
  constructor(
    IERC20 token,
    uint8 localTokenDecimals,
    address rmnProxy,
    address router
  ) SiloedLockReleaseTokenPool(token, localTokenDecimals, address(0), rmnProxy, router) {}

  function consumeOutbound(
    uint64 remoteChainSelector,
    uint256 amount
  ) external {
    _consumeOutboundRateLimit(address(i_token), remoteChainSelector, amount);
  }

  function consumeInbound(
    uint64 remoteChainSelector,
    uint256 amount
  ) external {
    _consumeInboundRateLimit(address(i_token), remoteChainSelector, amount);
  }
}

contract TokenPoolRateLimitHelperSetup is BaseTest {
  uint64 internal constant OTHER_CHAIN_SELECTOR = SOURCE_CHAIN_SELECTOR;

  address internal s_manager = makeAddr("manager");

  IERC20 internal s_token;
  SiloedLockReleaseTokenPoolRateLimitHarness internal s_pool;
  TokenPoolRateLimitHelper internal s_helper;

  function setUp() public virtual override {
    super.setUp();
    s_token = IERC20(
      address(
        new CrossChainToken(
          BaseERC20.ConstructorParams({
            name: "LINK",
            symbol: "LNK",
            decimals: 18,
            maxSupply: 0,
            preMint: 0,
            preMintRecipient: address(0),
            ccipAdmin: OWNER
          }),
          OWNER,
          OWNER
        )
      )
    );

    s_pool = new SiloedLockReleaseTokenPoolRateLimitHarness(
      s_token, DEFAULT_TOKEN_DECIMALS, address(s_mockRMNRemote), address(s_sourceRouter)
    );
    _addChain(DEST_CHAIN_SELECTOR, _cfg(true, 1000e18, 10e18), _cfg(true, 2000e18, 20e18));
    _addChain(OTHER_CHAIN_SELECTOR, _cfg(true, 500e18, 5e18), _cfg(true, 600e18, 6e18));

    s_helper = new TokenPoolRateLimitHelper(address(s_pool), OWNER);
    s_pool.setDynamicConfig(address(s_sourceRouter), address(s_helper), address(0));
    s_helper.setManager(s_manager, true);
  }

  function _addChain(
    uint64 remoteChainSelector,
    RateLimiter.Config memory outboundConfig,
    RateLimiter.Config memory inboundConfig
  ) internal {
    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = abi.encode(makeAddr("remotePool"));

    TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](1);
    chainUpdates[0] = TokenPool.ChainUpdate({
      remoteChainSelector: remoteChainSelector,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encode(makeAddr("remoteToken")),
      outboundRateLimiterConfig: outboundConfig,
      inboundRateLimiterConfig: inboundConfig
    });
    s_pool.applyChainUpdates(new uint64[](0), chainUpdates);
  }

  function _cfg(
    bool isEnabled,
    uint128 capacity,
    uint128 rate
  ) internal pure returns (RateLimiter.Config memory) {
    return RateLimiter.Config({isEnabled: isEnabled, capacity: capacity, rate: rate});
  }

  function _helperCfg(
    bool isEnabled,
    uint128 capacity,
    uint128 rate
  ) internal pure returns (ITokenPoolV2RateLimits.Config memory) {
    return ITokenPoolV2RateLimits.Config({isEnabled: isEnabled, capacity: capacity, rate: rate});
  }

  function _state(
    uint64 remoteChainSelector,
    bool fastFinality
  ) internal view returns (RateLimiter.TokenBucket memory outbound, RateLimiter.TokenBucket memory inbound) {
    return s_pool.getCurrentRateLimiterState(remoteChainSelector, fastFinality);
  }
}
