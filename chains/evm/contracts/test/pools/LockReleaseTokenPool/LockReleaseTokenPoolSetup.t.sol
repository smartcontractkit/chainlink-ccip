// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Router} from "../../../Router.sol";
import {ERC20LockBox} from "../../../pools/ERC20LockBox.sol";
import {LockReleaseTokenPool} from "../../../pools/LockReleaseTokenPool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenAdminRegistry} from "../../../tokenAdminRegistry/TokenAdminRegistry.sol";
import {BaseTest} from "../../BaseTest.t.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

contract LockReleaseTokenPoolSetup is BaseTest {
  IERC20 internal s_token;
  LockReleaseTokenPool internal s_lockReleaseTokenPool;
  LockReleaseTokenPool internal s_lockReleaseTokenPoolWithAllowList;
  address[] internal s_allowedList;

  address internal s_allowedOnRamp = address(123);
  address internal s_allowedOffRamp = address(234);

  address internal s_destPoolAddress = address(2736782345);
  address internal s_sourcePoolAddress = address(53852352095);

  ERC20LockBox internal s_lockBox;
  TokenAdminRegistry internal s_tokenAdminRegistry;

  function setUp() public virtual override {
    super.setUp();
    s_token = new BurnMintERC20("LINK", "LNK", 18, 0, 0);
    deal(address(s_token), OWNER, type(uint256).max);

    s_tokenAdminRegistry = new TokenAdminRegistry();
    s_lockBox = new ERC20LockBox(address(s_tokenAdminRegistry));

    s_lockReleaseTokenPool = new LockReleaseTokenPool(
      s_token,
      DEFAULT_TOKEN_DECIMALS,
      new address[](0),
      address(s_mockRMNRemote),
      address(s_sourceRouter),
      address(s_lockBox)
    );

    s_allowedList.push(vm.randomAddress());
    s_allowedList.push(OWNER);
    s_lockReleaseTokenPoolWithAllowList = new LockReleaseTokenPool(
      s_token,
      DEFAULT_TOKEN_DECIMALS,
      s_allowedList,
      address(s_mockRMNRemote),
      address(s_sourceRouter),
      address(s_lockBox)
    );

    // Mock token admin registry calls for both pools.
    vm.mockCall(
      address(s_tokenAdminRegistry),
      abi.encodeWithSignature("getPool(address)", address(s_token)),
      abi.encode(address(s_lockReleaseTokenPool))
    );

    vm.mockCall(
      address(s_tokenAdminRegistry),
      abi.encodeWithSignature("getTokenConfig(address)", address(s_token)),
      abi.encode(
        TokenAdminRegistry.TokenConfig({
          administrator: OWNER,
          pendingAdministrator: address(0),
          tokenPool: address(s_lockReleaseTokenPool)
        })
      )
    );

    // Configure allowed callers for the lockBox - both pools and OWNER.
    ERC20LockBox.AllowedCallerConfigArgs[] memory allowedCallers = new ERC20LockBox.AllowedCallerConfigArgs[](3);
    allowedCallers[0] = ERC20LockBox.AllowedCallerConfigArgs({token: address(s_token), caller: OWNER, allowed: true});
    allowedCallers[1] = ERC20LockBox.AllowedCallerConfigArgs({
      token: address(s_token),
      caller: address(s_lockReleaseTokenPool),
      allowed: true
    });
    allowedCallers[2] = ERC20LockBox.AllowedCallerConfigArgs({
      token: address(s_token),
      caller: address(s_lockReleaseTokenPoolWithAllowList),
      allowed: true
    });
    s_lockBox.configureAllowedCallers(allowedCallers);

    bytes[] memory remotePoolAddresses = new bytes[](1);
    remotePoolAddresses[0] = abi.encode(s_destPoolAddress);

    TokenPool.ChainUpdate[] memory chainUpdate = new TokenPool.ChainUpdate[](1);
    chainUpdate[0] = TokenPool.ChainUpdate({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      remotePoolAddresses: remotePoolAddresses,
      remoteTokenAddress: abi.encode(address(2)),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });

    s_lockReleaseTokenPool.applyChainUpdates(new uint64[](0), chainUpdate);
    s_lockReleaseTokenPoolWithAllowList.applyChainUpdates(new uint64[](0), chainUpdate);
    s_lockReleaseTokenPool.setRebalancer(OWNER);

    s_token.approve(address(s_lockReleaseTokenPool), type(uint256).max);

    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](1);
    Router.OffRamp[] memory offRampUpdates = new Router.OffRamp[](1);
    onRampUpdates[0] = Router.OnRamp({destChainSelector: DEST_CHAIN_SELECTOR, onRamp: s_allowedOnRamp});
    offRampUpdates[0] = Router.OffRamp({sourceChainSelector: SOURCE_CHAIN_SELECTOR, offRamp: s_allowedOffRamp});
    s_sourceRouter.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), offRampUpdates);
  }
}
