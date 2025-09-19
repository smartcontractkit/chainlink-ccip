// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Router} from "../../../Router.sol";

import {ERC20LockBox} from "../../../pools/ERC20LockBox.sol";
import {SiloedLockReleaseTokenPool} from "../../../pools/SiloedLockReleaseTokenPool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {BaseTest} from "../../BaseTest.t.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

import {TokenAdminRegistry} from "../../../tokenAdminRegistry/TokenAdminRegistry.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract SiloedLockReleaseTokenPoolSetup is BaseTest {
  IERC20 internal s_token;
  SiloedLockReleaseTokenPool internal s_siloedLockReleaseTokenPool;
  address[] internal s_allowedList;

  address internal s_allowedOnRamp = address(123);
  address internal s_allowedOffRamp = address(234);

  address internal s_destPoolAddress = address(2736782345);
  address internal s_sourcePoolAddress = address(53852352095);

  address internal s_siloedDestPoolAddress = address(4245234524);
  uint64 internal constant SILOED_CHAIN_SELECTOR = DEST_CHAIN_SELECTOR + 1;

  ERC20LockBox internal s_lockBox;
  TokenAdminRegistry internal s_tokenAdminRegistry;

  function setUp() public virtual override {
    super.setUp();
    s_token = new BurnMintERC20("LINK", "LNK", 18, 0, 0);
    deal(address(s_token), OWNER, type(uint256).max);

    s_tokenAdminRegistry = new TokenAdminRegistry();
    s_lockBox = new ERC20LockBox(address(s_tokenAdminRegistry));

    s_siloedLockReleaseTokenPool = new SiloedLockReleaseTokenPool(
      s_token,
      DEFAULT_TOKEN_DECIMALS,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_sourceRouter),
      address(s_lockBox)
    );

    // Mock the token pool for the token to be the siloed lock release token pool so that we can test the allowed caller configuration
    vm.mockCall(
      address(s_tokenAdminRegistry),
      abi.encodeWithSignature("getPool(address)", address(s_token)),
      abi.encode(address(s_siloedLockReleaseTokenPool))
    );

    // Mock the token config for the token to be the siloed lock release token pool so that we can test the allowed caller configuration
    vm.mockCall(
      address(s_tokenAdminRegistry),
      abi.encodeWithSignature("getTokenConfig(address)", address(s_token)),
      abi.encode(
        TokenAdminRegistry.TokenConfig({
          administrator: OWNER,
          pendingAdministrator: address(0),
          tokenPool: address(s_siloedLockReleaseTokenPool)
        })
      )
    );

    // Set the owner as an administrator for the token so that it can configure the allowed callers
    vm.mockCall(
      address(s_tokenAdminRegistry),
      abi.encodeWithSignature("isAdministrator(address,address)", address(s_token), OWNER),
      abi.encode(true)
    );

    ERC20LockBox.AllowedCallerConfigArgs[] memory allowedCallers = new ERC20LockBox.AllowedCallerConfigArgs[](1);
    allowedCallers[0] =
      ERC20LockBox.AllowedCallerConfigArgs({token: address(s_token), caller: address(OWNER), allowed: true});
    s_lockBox.configureAllowedCallers(allowedCallers);

    // Set the rebalancer for the token pool
    s_siloedLockReleaseTokenPool.setRebalancer(OWNER);

    s_token.approve(address(s_siloedLockReleaseTokenPool), type(uint256).max);

    bytes[] memory remotePoolAddresses = new bytes[](2);
    remotePoolAddresses[0] = abi.encode(s_destPoolAddress);
    remotePoolAddresses[1] = abi.encode(s_siloedDestPoolAddress);

    TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](3);
    chainUpdates[0] = TokenPool.ChainUpdate({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encode(address(2)),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });

    chainUpdates[1] = TokenPool.ChainUpdate({
      remoteChainSelector: SILOED_CHAIN_SELECTOR,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encode(address(2)),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });

    chainUpdates[2] = TokenPool.ChainUpdate({
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encode(address(2)),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });

    s_siloedLockReleaseTokenPool.applyChainUpdates(new uint64[](0), chainUpdates);

    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](3);
    Router.OffRamp[] memory offRampUpdates = new Router.OffRamp[](2);

    onRampUpdates[0] = Router.OnRamp({destChainSelector: DEST_CHAIN_SELECTOR, onRamp: s_allowedOnRamp});
    offRampUpdates[0] = Router.OffRamp({sourceChainSelector: SOURCE_CHAIN_SELECTOR, offRamp: s_allowedOffRamp});

    onRampUpdates[1] = Router.OnRamp({destChainSelector: SILOED_CHAIN_SELECTOR, onRamp: s_allowedOnRamp});
    offRampUpdates[1] = Router.OffRamp({sourceChainSelector: SILOED_CHAIN_SELECTOR, offRamp: s_allowedOffRamp});

    onRampUpdates[2] = Router.OnRamp({destChainSelector: SOURCE_CHAIN_SELECTOR, onRamp: s_allowedOnRamp});

    s_sourceRouter.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), offRampUpdates);

    // Apply Siloeing Rules
    SiloedLockReleaseTokenPool.SiloConfigUpdate[] memory adds = new SiloedLockReleaseTokenPool.SiloConfigUpdate[](1);

    adds[0] =
      SiloedLockReleaseTokenPool.SiloConfigUpdate({remoteChainSelector: SILOED_CHAIN_SELECTOR, rebalancer: OWNER});

    s_siloedLockReleaseTokenPool.updateSiloDesignations(new uint64[](0), adds);

    assertTrue(s_siloedLockReleaseTokenPool.isSiloed(SILOED_CHAIN_SELECTOR));
    assertFalse(s_siloedLockReleaseTokenPool.isSiloed(DEST_CHAIN_SELECTOR));

    s_siloedLockReleaseTokenPool.setSiloRebalancer(SILOED_CHAIN_SELECTOR, OWNER);
  }
}
