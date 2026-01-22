// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Router} from "../../../Router.sol";
import {AdvancedPoolHooks} from "../../../pools/AdvancedPoolHooks.sol";
import {ERC20LockBox} from "../../../pools/ERC20LockBox.sol";
import {LockReleaseTokenPool} from "../../../pools/LockReleaseTokenPool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {BaseTest} from "../../BaseTest.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

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

  function setUp() public virtual override {
    super.setUp();
    s_token = IERC20(address(new BurnMintERC20("LINK", "LNK", 18, 0, 0)));
    deal(address(s_token), OWNER, type(uint256).max);

    s_lockBox = new ERC20LockBox(address(s_token));

    s_lockReleaseTokenPool = new LockReleaseTokenPool(
      s_token, DEFAULT_TOKEN_DECIMALS, address(0), address(s_mockRMNRemote), address(s_sourceRouter), address(s_lockBox)
    );

    s_allowedList.push(vm.randomAddress());
    s_allowedList.push(OWNER);
    AdvancedPoolHooks advancedHooks = new AdvancedPoolHooks(s_allowedList, 0, address(0));
    s_lockReleaseTokenPoolWithAllowList = new LockReleaseTokenPool(
      s_token,
      DEFAULT_TOKEN_DECIMALS,
      address(advancedHooks),
      address(s_mockRMNRemote),
      address(s_sourceRouter),
      address(s_lockBox)
    );

    // Configure allowed callers for the lockBox - both pools.
    address[] memory allowedCallers = new address[](2);
    allowedCallers[0] = address(s_lockReleaseTokenPool);
    allowedCallers[1] = address(s_lockReleaseTokenPoolWithAllowList);
    s_lockBox.applyAuthorizedCallerUpdates(
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: allowedCallers, removedCallers: new address[](0)})
    );

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

    s_token.approve(address(s_lockReleaseTokenPool), type(uint256).max);

    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](1);
    Router.OffRamp[] memory offRampUpdates = new Router.OffRamp[](1);
    onRampUpdates[0] = Router.OnRamp({destChainSelector: DEST_CHAIN_SELECTOR, onRamp: s_allowedOnRamp});
    offRampUpdates[0] = Router.OffRamp({sourceChainSelector: SOURCE_CHAIN_SELECTOR, offRamp: s_allowedOffRamp});
    s_sourceRouter.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), offRampUpdates);
  }
}
