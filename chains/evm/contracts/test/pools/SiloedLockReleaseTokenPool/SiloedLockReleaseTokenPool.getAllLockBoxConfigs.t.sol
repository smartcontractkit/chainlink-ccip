// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Router} from "../../../Router.sol";
import {ERC20LockBox} from "../../../pools/ERC20LockBox.sol";
import {SiloedLockReleaseTokenPool} from "../../../pools/SiloedLockReleaseTokenPool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {BaseTest} from "../../BaseTest.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract SiloedLockReleaseTokenPool_getAllLockBoxConfigs is BaseTest {
  uint64 internal constant CHAIN_A = 1;
  uint64 internal constant CHAIN_B = 2;
  uint64 internal constant CHAIN_C = 3;
  uint64 internal constant CHAIN_D = 4;

  IERC20 internal s_token;
  SiloedLockReleaseTokenPool internal s_pool;
  ERC20LockBox internal s_sharedLockBox;
  ERC20LockBox internal s_isolatedLockBox;

  function setUp() public override {
    super.setUp();

    s_token = IERC20(address(new BurnMintERC20("TEST", "TST", 18, 0, 0)));
    s_sharedLockBox = new ERC20LockBox(address(s_token));
    s_isolatedLockBox = new ERC20LockBox(address(s_token));

    s_pool = new SiloedLockReleaseTokenPool(
      s_token, DEFAULT_TOKEN_DECIMALS, address(0), address(s_mockRMNRemote), address(s_sourceRouter)
    );

    // Authorize pool on lockboxes.
    address[] memory allowedCallers = new address[](1);
    allowedCallers[0] = address(s_pool);
    s_sharedLockBox.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: allowedCallers, removedCallers: new address[](0)})
    );
    s_isolatedLockBox.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: allowedCallers, removedCallers: new address[](0)})
    );

    // Setup router.
    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](4);
    onRampUpdates[0] = Router.OnRamp({destChainSelector: CHAIN_A, onRamp: address(1)});
    onRampUpdates[1] = Router.OnRamp({destChainSelector: CHAIN_B, onRamp: address(2)});
    onRampUpdates[2] = Router.OnRamp({destChainSelector: CHAIN_C, onRamp: address(3)});
    onRampUpdates[3] = Router.OnRamp({destChainSelector: CHAIN_D, onRamp: address(4)});
    s_sourceRouter.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), new Router.OffRamp[](0));
  }

  function test_getAllLockBoxConfigs() public {
    // Add 4 chains to the pool.
    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = abi.encode(address(999));

    TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](4);
    for (uint256 i = 0; i < 4; ++i) {
      chainUpdates[i] = TokenPool.ChainUpdate({
        remoteChainSelector: uint64(i + 1), // CHAIN_A, CHAIN_B, CHAIN_C, CHAIN_D.
        remotePoolAddresses: remotePoolAddresses,
        remoteTokenAddress: abi.encode(address(2)),
        outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
        inboundRateLimiterConfig: _getInboundRateLimiterConfig()
      });
    }
    s_pool.applyChainUpdates(new uint64[](0), chainUpdates);

    // Configure lockboxes:
    // - CHAIN_A: shared lockbox.
    // - CHAIN_B: shared lockbox (same as CHAIN_A).
    // - CHAIN_C: isolated lockbox.
    // - CHAIN_D: no lockbox configured (will be address(0)).
    SiloedLockReleaseTokenPool.LockBoxConfig[] memory lockBoxConfigs = new SiloedLockReleaseTokenPool.LockBoxConfig[](3);
    lockBoxConfigs[0] =
      SiloedLockReleaseTokenPool.LockBoxConfig({remoteChainSelector: CHAIN_A, lockBox: address(s_sharedLockBox)});
    lockBoxConfigs[1] =
      SiloedLockReleaseTokenPool.LockBoxConfig({remoteChainSelector: CHAIN_B, lockBox: address(s_sharedLockBox)});
    lockBoxConfigs[2] =
      SiloedLockReleaseTokenPool.LockBoxConfig({remoteChainSelector: CHAIN_C, lockBox: address(s_isolatedLockBox)});
    // CHAIN_D intentionally not configured.
    s_pool.configureLockBoxes(lockBoxConfigs);

    SiloedLockReleaseTokenPool.LockBoxConfig[] memory configs = s_pool.getAllLockBoxConfigs();

    assertEq(configs.length, 4);
    assertEq(configs[0].lockBox, address(s_sharedLockBox)); // CHAIN_A - shared.
    assertEq(configs[1].lockBox, address(s_sharedLockBox)); // CHAIN_B - shared.
    assertEq(configs[2].lockBox, address(s_isolatedLockBox)); // CHAIN_C - isolated.
    assertEq(configs[3].lockBox, address(0)); // CHAIN_D - unconfigured.
  }
}
