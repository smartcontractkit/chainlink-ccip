// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdvancedPoolHooks} from "../../../interfaces/IAdvancedPoolHooks.sol";
import {IPoolV2} from "../../../interfaces/IPoolV2.sol";

import {LombardTokenPool} from "../../../pools/Lombard/LombardTokenPool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {LombardTokenPoolHelper} from "../../helpers/LombardTokenPoolHelper.sol";
import {LombardTokenPoolSetup} from "./LombardTokenPoolSetup.t.sol";

import {IERC20Metadata} from "@openzeppelin/contracts@5.3.0/token/ERC20/extensions/IERC20Metadata.sol";

contract LombardTokenPool_getRequiredCCVs is LombardTokenPoolSetup {
  address internal s_hooksAddr = makeAddr("hooks");
  LombardTokenPoolHelper internal s_poolWithHooks;

  function setUp() public override {
    super.setUp();

    s_poolWithHooks = new LombardTokenPoolHelper(
      IERC20Metadata(address(s_token)),
      address(s_verifierResolver),
      s_bridge,
      address(0),
      s_hooksAddr,
      address(s_mockRMNRemote),
      address(s_sourceRouter),
      DEFAULT_TOKEN_DECIMALS
    );

    bytes[] memory remotePools = new bytes[](1);
    remotePools[0] = abi.encode(s_remotePool);
    TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](1);
    chainUpdates[0] = TokenPool.ChainUpdate({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      remotePoolAddresses: remotePools,
      remoteTokenAddress: abi.encode(s_remoteToken),
      outboundRateLimiterConfig: _getOutboundRateLimiterConfig(),
      inboundRateLimiterConfig: _getInboundRateLimiterConfig()
    });
    s_poolWithHooks.applyChainUpdates(new uint64[](0), chainUpdates);
  }

  function _mockHooksReturning(
    address[] memory ccvs
  ) internal {
    vm.mockCall(s_hooksAddr, abi.encodeWithSelector(IAdvancedPoolHooks.getRequiredCCVs.selector), abi.encode(ccvs));
  }

  /// @notice Pool has no advancedPoolHooks (address(0)) → super returns [] → revert.
  function test_getRequiredCCVs_RevertWhen_ZeroCCVs_NoHooksConfigured() public {
    vm.expectRevert(LombardTokenPool.LombardMustUseCCVConfigForV2Flow.selector);
    s_pool.getRequiredCCVs(address(s_token), DEST_CHAIN_SELECTOR, 1e18, 0, "", IPoolV2.MessageDirection.Outbound);
  }

  /// @notice Hooks present but returns an empty array → revert.
  function test_getRequiredCCVs_RevertWhen_LombardMustUseCCVConfigForV2Flow_ZeroCCVs_HooksReturnEmpty() public {
    _mockHooksReturning(new address[](0));
    vm.expectRevert(LombardTokenPool.LombardMustUseCCVConfigForV2Flow.selector);
    s_poolWithHooks.getRequiredCCVs(
      address(s_token), DEST_CHAIN_SELECTOR, 1e18, 0, "", IPoolV2.MessageDirection.Outbound
    );
  }

  /// @notice Exactly one CCV is insufficient for the V2 flow → revert.
  function test_getRequiredCCVs_RevertWhen_LombardMustUseCCVConfigForV2Flow_OneCCV() public {
    address[] memory ccvs = new address[](1);
    ccvs[0] = makeAddr("ccv1");
    _mockHooksReturning(ccvs);
    vm.expectRevert(LombardTokenPool.LombardMustUseCCVConfigForV2Flow.selector);
    s_poolWithHooks.getRequiredCCVs(
      address(s_token), DEST_CHAIN_SELECTOR, 1e18, 0, "", IPoolV2.MessageDirection.Outbound
    );
  }

  function test_getRequiredCCVs_ValidConfig_TwoCCVs() public {
    address[] memory ccvs = new address[](2);
    ccvs[0] = makeAddr("ccv1");
    ccvs[1] = makeAddr("ccv2");
    _mockHooksReturning(ccvs);

    address[] memory result = s_poolWithHooks.getRequiredCCVs(
      address(s_token), DEST_CHAIN_SELECTOR, 1e18, 0, "", IPoolV2.MessageDirection.Outbound
    );

    assertEq(result.length, 2);
    assertEq(result[0], ccvs[0]);
    assertEq(result[1], ccvs[1]);
  }
}
