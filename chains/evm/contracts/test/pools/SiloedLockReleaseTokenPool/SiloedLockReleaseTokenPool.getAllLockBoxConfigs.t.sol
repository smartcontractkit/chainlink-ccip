// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Router} from "../../../Router.sol";
import {ERC20LockBox} from "../../../pools/ERC20LockBox.sol";
import {SiloedLockReleaseTokenPool} from "../../../pools/SiloedLockReleaseTokenPool.sol";
import {BaseERC20} from "../../../tmp/BaseERC20.sol";
import {CrossChainToken} from "../../../tmp/CrossChainToken.sol";
import {BaseTest} from "../../BaseTest.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract SiloedLockReleaseTokenPool_getAllLockBoxConfigs is BaseTest {
  uint64 internal constant CHAIN_A = 1;
  uint64 internal constant CHAIN_B = 2;
  uint64 internal constant CHAIN_C = 3;

  IERC20 internal s_token;
  SiloedLockReleaseTokenPool internal s_pool;
  ERC20LockBox internal s_sharedLockBox;
  ERC20LockBox internal s_isolatedLockBox;

  function setUp() public override {
    super.setUp();

    s_token = IERC20(
      address(
        new CrossChainToken(
          BaseERC20.ConstructorParams({
            name: "TEST", symbol: "TST", decimals: 18, maxSupply: 0, preMint: 0, ccipAdmin: OWNER
          }),
          OWNER,
          OWNER
        )
      )
    );
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
    s_sourceRouter.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), new Router.OffRamp[](0));
  }

  function test_getAllLockBoxConfigs() public {
    // - CHAIN_A: shared lockbox.
    // - CHAIN_B: shared lockbox (same as CHAIN_A).
    // - CHAIN_C: isolated lockbox.
    SiloedLockReleaseTokenPool.LockBoxConfig[] memory lockBoxConfigs = new SiloedLockReleaseTokenPool.LockBoxConfig[](3);
    lockBoxConfigs[0] =
      SiloedLockReleaseTokenPool.LockBoxConfig({remoteChainSelector: CHAIN_A, lockBox: address(s_sharedLockBox)});
    lockBoxConfigs[1] =
      SiloedLockReleaseTokenPool.LockBoxConfig({remoteChainSelector: CHAIN_B, lockBox: address(s_sharedLockBox)});
    lockBoxConfigs[2] =
      SiloedLockReleaseTokenPool.LockBoxConfig({remoteChainSelector: CHAIN_C, lockBox: address(s_isolatedLockBox)});
    s_pool.configureLockBoxes(lockBoxConfigs);

    SiloedLockReleaseTokenPool.LockBoxConfig[] memory configs = s_pool.getAllLockBoxConfigs();

    assertEq(configs.length, 3);
    assertEq(configs[0].remoteChainSelector, CHAIN_A);
    assertEq(configs[0].lockBox, address(s_sharedLockBox));
    assertEq(configs[1].remoteChainSelector, CHAIN_B);
    assertEq(configs[1].lockBox, address(s_sharedLockBox));
    assertEq(configs[2].remoteChainSelector, CHAIN_C);
    assertEq(configs[2].lockBox, address(s_isolatedLockBox));
  }
}
