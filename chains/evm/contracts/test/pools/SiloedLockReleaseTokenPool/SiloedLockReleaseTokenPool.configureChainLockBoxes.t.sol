// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Router} from "../../../Router.sol";
import {ERC20LockBox} from "../../../pools/ERC20LockBox.sol";
import {SiloedLockReleaseTokenPool} from "../../../pools/SiloedLockReleaseTokenPool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {BaseTest} from "../../BaseTest.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

contract SiloedLockReleaseTokenPool_configureChainLockBoxes is BaseTest {
  uint64 internal constant SILOED_CHAIN_SELECTOR = DEST_CHAIN_SELECTOR + 5;
  bytes32 internal constant SILO_DOMAIN_ID = bytes32(uint256(SILOED_CHAIN_SELECTOR));

  ERC20LockBox internal s_unsiloed;
  SiloedLockReleaseTokenPool internal s_pool;
  BurnMintERC20 internal s_token;

  function setUp() public override {
    super.setUp();
    s_token = new BurnMintERC20("TKN", "T", 18, 0, 0);
    deal(address(s_token), OWNER, type(uint256).max);

    s_unsiloed = new ERC20LockBox(address(s_token), bytes32(0));
    s_pool = new SiloedLockReleaseTokenPool(
      s_token,
      DEFAULT_TOKEN_DECIMALS,
      address(0),
      address(s_mockRMNRemote),
      address(s_sourceRouter),
      address(s_unsiloed)
    );

    address[] memory allowedCallers = new address[](1);
    allowedCallers[0] = address(s_pool);
    s_unsiloed.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: allowedCallers, removedCallers: new address[](0)})
    );

    // basic router config to allow on/off ramps
    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](1);
    Router.OffRamp[] memory offRampUpdates = new Router.OffRamp[](1);
    onRampUpdates[0] = Router.OnRamp({destChainSelector: SILOED_CHAIN_SELECTOR, onRamp: address(123)});
    offRampUpdates[0] = Router.OffRamp({sourceChainSelector: SILOED_CHAIN_SELECTOR, offRamp: address(234)});
    s_sourceRouter.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), offRampUpdates);

    // register the siloed chain as supported
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

  function test_configureChainLockBoxes() public {
    ERC20LockBox siloLockBox = new ERC20LockBox(address(s_token), SILO_DOMAIN_ID);
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

    SiloedLockReleaseTokenPool.SiloConfigUpdate[] memory adds = new SiloedLockReleaseTokenPool.SiloConfigUpdate[](1);
    adds[0] =
      SiloedLockReleaseTokenPool.SiloConfigUpdate({remoteChainSelector: SILOED_CHAIN_SELECTOR, rebalancer: OWNER});
    s_pool.updateSiloDesignations(new uint64[](0), adds);
    s_pool.setSiloRebalancer(SILOED_CHAIN_SELECTOR, OWNER);

    s_token.approve(address(s_pool), type(uint256).max);
    s_pool.provideSiloedLiquidity(SILOED_CHAIN_SELECTOR, 1e18);

    assertEq(s_token.balanceOf(address(siloLockBox)), 1e18);
  }

  function test_configureChainLockBoxes_RevertWhen_LockBoxNotConfigured() public {
    s_token.approve(address(s_pool), type(uint256).max);
    vm.expectRevert(abi.encodeWithSelector(SiloedLockReleaseTokenPool.ChainNotSiloed.selector, SILOED_CHAIN_SELECTOR));
    s_pool.provideSiloedLiquidity(SILOED_CHAIN_SELECTOR, 1e18);
  }

  function test_configureChainLockBoxes_RevertWhen_InvalidLockBoxLiquidityDomain() public {
    ERC20LockBox badLockBox = new ERC20LockBox(address(s_token), 0);
    SiloedLockReleaseTokenPool.LockBoxConfig[] memory lockBoxes = new SiloedLockReleaseTokenPool.LockBoxConfig[](1);
    lockBoxes[0] = SiloedLockReleaseTokenPool.LockBoxConfig({
      remoteChainSelector: SILOED_CHAIN_SELECTOR, lockBox: address(badLockBox)
    });

    vm.expectRevert(
      abi.encodeWithSelector(SiloedLockReleaseTokenPool.InvalidLockBoxLiquidityDomain.selector, bytes32(uint256(0)))
    );
    s_pool.configureLockBoxes(lockBoxes);
  }

  function test_configureChainLockBoxes_RevertWhen_InvalidToken() public {
    address wrongToken = address(999);
    ERC20LockBox invalidLockBox = new ERC20LockBox(wrongToken, SILO_DOMAIN_ID);
    SiloedLockReleaseTokenPool.LockBoxConfig[] memory lockBoxes = new SiloedLockReleaseTokenPool.LockBoxConfig[](1);
    lockBoxes[0] = SiloedLockReleaseTokenPool.LockBoxConfig({
      remoteChainSelector: SILOED_CHAIN_SELECTOR, lockBox: address(invalidLockBox)
    });

    vm.expectRevert(abi.encodeWithSelector(TokenPool.InvalidToken.selector, wrongToken));
    s_pool.configureLockBoxes(lockBoxes);
  }
}
