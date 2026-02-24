// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Router} from "../../../Router.sol";
import {ERC20LockBox} from "../../../pools/ERC20LockBox.sol";
import {SiloedLockReleaseTokenPool} from "../../../pools/SiloedLockReleaseTokenPool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {BaseERC20} from "../../../tmp/BaseERC20.sol";
import {CrossChainToken} from "../../../tmp/CrossChainToken.sol";
import {BaseTest} from "../../BaseTest.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract SiloedLockReleaseTokenPool_configureLockBoxes is BaseTest {
  uint64 internal constant SILOED_CHAIN_SELECTOR = DEST_CHAIN_SELECTOR + 5;

  SiloedLockReleaseTokenPool internal s_pool;
  CrossChainToken internal s_token;

  function setUp() public override {
    super.setUp();
    s_token = new CrossChainToken(
      BaseERC20.ConstructorParams({name: "TKN", symbol: "T", decimals: 18, maxSupply: 0, preMint: 0, ccipAdmin: OWNER}),
      OWNER,
      OWNER
    );
    deal(address(s_token), OWNER, type(uint256).max);

    s_pool = new SiloedLockReleaseTokenPool(
      IERC20(address(s_token)), DEFAULT_TOKEN_DECIMALS, address(0), address(s_mockRMNRemote), address(s_sourceRouter)
    );

    // Basic router config to allow on/off ramps.
    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](1);
    Router.OffRamp[] memory offRampUpdates = new Router.OffRamp[](1);
    onRampUpdates[0] = Router.OnRamp({destChainSelector: SILOED_CHAIN_SELECTOR, onRamp: address(123)});
    offRampUpdates[0] = Router.OffRamp({sourceChainSelector: SILOED_CHAIN_SELECTOR, offRamp: address(234)});
    s_sourceRouter.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), offRampUpdates);

    // Register the siloed chain as supported.
    TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](1);
    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = abi.encode(address(999));
    chainUpdates[0] = TokenPool.ChainUpdate({
      remoteChainSelector: SILOED_CHAIN_SELECTOR,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encode(address(2)),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });
    s_pool.applyChainUpdates(new uint64[](0), chainUpdates);
  }

  function test_configureLockBoxes() public {
    ERC20LockBox siloLockBox = new ERC20LockBox(address(s_token));
    address[] memory allowedCallers = new address[](1);
    allowedCallers[0] = address(s_pool);
    siloLockBox.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: allowedCallers, removedCallers: new address[](0)})
    );

    SiloedLockReleaseTokenPool.LockBoxConfig[] memory lockBoxes = new SiloedLockReleaseTokenPool.LockBoxConfig[](1);
    lockBoxes[0] = SiloedLockReleaseTokenPool.LockBoxConfig({
      remoteChainSelector: SILOED_CHAIN_SELECTOR, lockBox: address(siloLockBox)
    });
    s_pool.configureLockBoxes(lockBoxes);

    assertEq(address(s_pool.getLockBox(SILOED_CHAIN_SELECTOR)), address(siloLockBox));
  }

  function test_configureLockBoxes_RevertWhen_InvalidToken() public {
    address wrongToken = address(999);
    ERC20LockBox invalidLockBox = new ERC20LockBox(wrongToken);
    SiloedLockReleaseTokenPool.LockBoxConfig[] memory lockBoxes = new SiloedLockReleaseTokenPool.LockBoxConfig[](1);
    lockBoxes[0] = SiloedLockReleaseTokenPool.LockBoxConfig({
      remoteChainSelector: SILOED_CHAIN_SELECTOR, lockBox: address(invalidLockBox)
    });

    vm.expectRevert(abi.encodeWithSelector(TokenPool.InvalidToken.selector, address(s_token)));
    s_pool.configureLockBoxes(lockBoxes);
  }

  function test_configureLockBoxes_RevertWhen_ZeroAddressInvalid() public {
    SiloedLockReleaseTokenPool.LockBoxConfig[] memory lockBoxes = new SiloedLockReleaseTokenPool.LockBoxConfig[](1);
    lockBoxes[0] =
      SiloedLockReleaseTokenPool.LockBoxConfig({remoteChainSelector: SILOED_CHAIN_SELECTOR, lockBox: address(0)});

    vm.expectRevert(abi.encodeWithSelector(TokenPool.ZeroAddressInvalid.selector));
    s_pool.configureLockBoxes(lockBoxes);
  }
}

