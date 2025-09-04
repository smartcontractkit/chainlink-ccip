// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRouter} from "../../../interfaces/IRouter.sol";

import {CCVConfigValidation} from "../../../libraries/CCVConfigValidation.sol";
import {CCVProxy} from "../../../onRamp/CCVProxy.sol";
import {CCVProxySetup} from "./CCVProxySetup.t.sol";

contract CCVProxy_applyDestChainConfigUpdates is CCVProxySetup {
  uint64 internal constant NEW_DEST_SELECTOR = uint64(uint256(keccak256("NEW_DEST_SELECTOR")));

  function test_applyDestChainConfigUpdates_SetsConfigAndEmitsEvent() public {
    IRouter router = s_sourceRouter;
    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = makeAddr("defaultCCV1");
    address[] memory laneMandated = new address[](1);
    laneMandated[0] = makeAddr("laneCCV1");
    address defaultExecutor = makeAddr("defaultExecutor");

    CCVProxy.DestChainConfigArgs[] memory args = new CCVProxy.DestChainConfigArgs[](1);
    args[0] = CCVProxy.DestChainConfigArgs({
      destChainSelector: NEW_DEST_SELECTOR,
      router: router,
      defaultCCVs: defaultCCVs,
      laneMandatedCCVs: laneMandated,
      defaultExecutor: defaultExecutor
    });

    vm.expectEmit();
    emit CCVProxy.DestChainConfigSet(NEW_DEST_SELECTOR, 0, IRouter(router), defaultCCVs, laneMandated, defaultExecutor);
    s_ccvProxy.applyDestChainConfigUpdates(args);

    CCVProxy.DestChainConfig memory cfg = s_ccvProxy.getDestChainConfig(NEW_DEST_SELECTOR);
    assertEq(address(cfg.router), address(router));
    assertEq(cfg.defaultExecutor, defaultExecutor);
    assertEq(cfg.sequenceNumber, 0);
    assertEq(cfg.defaultCCVs, defaultCCVs);
    assertEq(cfg.laneMandatedCCVs, laneMandated);
  }

  function test_applyDestChainConfigUpdates_AllowsZeroRouterToPause() public {
    CCVProxy.DestChainConfigArgs[] memory args = new CCVProxy.DestChainConfigArgs[](1);
    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = makeAddr("defaultCCV");
    args[0] = CCVProxy.DestChainConfigArgs({
      destChainSelector: NEW_DEST_SELECTOR + 1,
      router: IRouter(address(0)),
      defaultCCVs: defaultCCVs,
      laneMandatedCCVs: new address[](0),
      defaultExecutor: makeAddr("executor")
    });

    // Should not revert, router can be zero.
    s_ccvProxy.applyDestChainConfigUpdates(args);
    CCVProxy.DestChainConfig memory cfg = s_ccvProxy.getDestChainConfig(NEW_DEST_SELECTOR + 1);
    assertEq(address(cfg.router), address(0));
  }

  function test_applyDestChainConfigUpdates_RevertWhen_InvalidDestChainConfig_ZeroSelector() public {
    CCVProxy.DestChainConfigArgs[] memory args = new CCVProxy.DestChainConfigArgs[](1);
    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = makeAddr("defaultCCV");
    args[0] = CCVProxy.DestChainConfigArgs({
      destChainSelector: 0,
      router: s_sourceRouter,
      defaultCCVs: defaultCCVs,
      laneMandatedCCVs: new address[](0),
      defaultExecutor: makeAddr("executor")
    });

    vm.expectRevert(abi.encodeWithSelector(CCVProxy.InvalidDestChainConfig.selector, uint64(0)));
    s_ccvProxy.applyDestChainConfigUpdates(args);
  }

  function test_applyDestChainConfigUpdates_RevertWhen_ZeroInDefaultCCVs() public {
    CCVProxy.DestChainConfigArgs[] memory args = new CCVProxy.DestChainConfigArgs[](1);
    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = address(0); // invalid
    args[0] = CCVProxy.DestChainConfigArgs({
      destChainSelector: NEW_DEST_SELECTOR + 2,
      router: s_sourceRouter,
      defaultCCVs: defaultCCVs,
      laneMandatedCCVs: new address[](0),
      defaultExecutor: makeAddr("executor")
    });

    vm.expectRevert(CCVConfigValidation.ZeroAddressNotAllowed.selector);
    s_ccvProxy.applyDestChainConfigUpdates(args);
  }

  function test_applyDestChainConfigUpdates_RevertWhen_ZeroInLaneMandatedCCVs() public {
    CCVProxy.DestChainConfigArgs[] memory args = new CCVProxy.DestChainConfigArgs[](1);
    address[] memory laneMandatedCCVs = new address[](2);
    laneMandatedCCVs[1] = address(0); // invalid
    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = makeAddr("defaultCCV");
    args[0] = CCVProxy.DestChainConfigArgs({
      destChainSelector: NEW_DEST_SELECTOR + 3,
      router: s_sourceRouter,
      defaultCCVs: defaultCCVs,
      laneMandatedCCVs: laneMandatedCCVs,
      defaultExecutor: makeAddr("executor")
    });

    vm.expectRevert(CCVConfigValidation.ZeroAddressNotAllowed.selector);
    s_ccvProxy.applyDestChainConfigUpdates(args);
  }

  function test_applyDestChainConfigUpdates_RevertWhen_DuplicateWithinDefaultCCVs() public {
    CCVProxy.DestChainConfigArgs[] memory args = new CCVProxy.DestChainConfigArgs[](1);
    address dup = makeAddr("dupCCV");
    address[] memory defaultCCVs = new address[](2);
    defaultCCVs[0] = dup;
    defaultCCVs[1] = dup; // duplicate
    args[0] = CCVProxy.DestChainConfigArgs({
      destChainSelector: NEW_DEST_SELECTOR + 4,
      router: s_sourceRouter,
      defaultCCVs: defaultCCVs,
      laneMandatedCCVs: new address[](0),
      defaultExecutor: makeAddr("executor")
    });

    vm.expectRevert(abi.encodeWithSelector(CCVConfigValidation.DuplicateCCVNotAllowed.selector, dup));
    s_ccvProxy.applyDestChainConfigUpdates(args);
  }

  function test_applyDestChainConfigUpdates_RevertWhen_DuplicateWithinLaneMandated() public {
    CCVProxy.DestChainConfigArgs[] memory args = new CCVProxy.DestChainConfigArgs[](1);
    address dup = makeAddr("dupLane");
    address[] memory lane = new address[](2);
    lane[0] = dup;
    lane[1] = dup; // duplicate
    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = makeAddr("defaultCCV");
    args[0] = CCVProxy.DestChainConfigArgs({
      destChainSelector: NEW_DEST_SELECTOR + 5,
      router: s_sourceRouter,
      defaultCCVs: defaultCCVs,
      laneMandatedCCVs: lane,
      defaultExecutor: makeAddr("executor")
    });

    vm.expectRevert(abi.encodeWithSelector(CCVConfigValidation.DuplicateCCVNotAllowed.selector, dup));
    s_ccvProxy.applyDestChainConfigUpdates(args);
  }

  function test_applyDestChainConfigUpdates_RevertWhen_DuplicateAcrossDefaultAndLane() public {
    CCVProxy.DestChainConfigArgs[] memory args = new CCVProxy.DestChainConfigArgs[](1);
    address dup = makeAddr("dupAcross");
    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = dup;
    address[] memory lane = new address[](1);
    lane[0] = dup; // duplicate across arrays
    args[0] = CCVProxy.DestChainConfigArgs({
      destChainSelector: NEW_DEST_SELECTOR + 6,
      router: s_sourceRouter,
      defaultCCVs: defaultCCVs,
      laneMandatedCCVs: lane,
      defaultExecutor: makeAddr("executor")
    });

    vm.expectRevert(abi.encodeWithSelector(CCVConfigValidation.DuplicateCCVNotAllowed.selector, dup));
    s_ccvProxy.applyDestChainConfigUpdates(args);
  }

  function test_applyDestChainConfigUpdates_RevertWhen_DefaultCCVsEmpty() public {
    CCVProxy.DestChainConfigArgs[] memory args = new CCVProxy.DestChainConfigArgs[](1);
    // Empty defaultCCVs should revert.
    address[] memory defaultCCVs = new address[](0);
    address[] memory lane = new address[](0);
    args[0] = CCVProxy.DestChainConfigArgs({
      destChainSelector: NEW_DEST_SELECTOR + 7,
      router: s_sourceRouter,
      defaultCCVs: defaultCCVs,
      laneMandatedCCVs: lane,
      defaultExecutor: makeAddr("executor")
    });

    vm.expectRevert(CCVConfigValidation.MustSpecifyDefaultOrRequiredCCVs.selector);
    s_ccvProxy.applyDestChainConfigUpdates(args);
  }

  function test_applyDestChainConfigUpdates_RevertWhen_DefaultExecutorZero() public {
    CCVProxy.DestChainConfigArgs[] memory args = new CCVProxy.DestChainConfigArgs[](1);
    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = makeAddr("defaultCCV");
    args[0] = CCVProxy.DestChainConfigArgs({
      destChainSelector: NEW_DEST_SELECTOR + 8,
      router: s_sourceRouter,
      defaultCCVs: defaultCCVs,
      laneMandatedCCVs: new address[](0),
      defaultExecutor: address(0)
    });

    vm.expectRevert(CCVProxy.InvalidConfig.selector);
    s_ccvProxy.applyDestChainConfigUpdates(args);
  }

  function test_applyDestChainConfigUpdates_RevertWhen_DestIsLocalChain() public {
    // Using SOURCE_CHAIN_SELECTOR as local chain selector from setup.
    CCVProxy.DestChainConfigArgs[] memory args = new CCVProxy.DestChainConfigArgs[](1);
    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = makeAddr("defaultCCV");
    args[0] = CCVProxy.DestChainConfigArgs({
      destChainSelector: SOURCE_CHAIN_SELECTOR,
      router: s_sourceRouter,
      defaultCCVs: defaultCCVs,
      laneMandatedCCVs: new address[](0),
      defaultExecutor: makeAddr("executor")
    });

    vm.expectRevert(abi.encodeWithSelector(CCVProxy.InvalidDestChainConfig.selector, SOURCE_CHAIN_SELECTOR));
    s_ccvProxy.applyDestChainConfigUpdates(args);
  }
}
