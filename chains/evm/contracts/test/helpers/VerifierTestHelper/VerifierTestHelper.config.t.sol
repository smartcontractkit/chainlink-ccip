// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRouter} from "../../../interfaces/IRouter.sol";

import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {FinalityCodec} from "../../../libraries/FinalityCodec.sol";
import {VerifierTestHelper} from "../VerifierTestHelper.sol";
import {VerifierTestHelperSetup} from "./VerifierTestHelperSetup.t.sol";

contract VerifierTestHelper_config is VerifierTestHelperSetup {
  function test_applyRemoteChainConfigUpdates() public {
    _configureRemoteChain(DEST_CHAIN_SELECTOR);

    (BaseVerifier.RemoteChainConfigArgs memory config,) = s_verifier.getRemoteChainConfig(DEST_CHAIN_SELECTOR);
    assertEq(address(config.router), s_verifier.getTestRouter());
    assertTrue(config.allowlistEnabled);
  }

  function test_applyRemoteChainConfigUpdates_RevertWhen_AllowlistIsDisabled() public {
    BaseVerifier.RemoteChainConfigArgs[] memory configs = new BaseVerifier.RemoteChainConfigArgs[](1);
    configs[0] = BaseVerifier.RemoteChainConfigArgs({
      router: IRouter(s_verifier.getTestRouter()),
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: false,
      feeUSDCents: 1,
      gasForVerification: 1,
      payloadSizeBytes: 1
    });

    vm.prank(OWNER);
    vm.expectRevert(VerifierTestHelper.MustUseAllowlist.selector);
    s_verifier.applyRemoteChainConfigUpdates(configs);
  }

  function test_applyRemoteChainConfigUpdates_RevertWhen_MustUseTestRouter() public {
    BaseVerifier.RemoteChainConfigArgs[] memory configs = new BaseVerifier.RemoteChainConfigArgs[](1);
    configs[0] = BaseVerifier.RemoteChainConfigArgs({
      router: IRouter(makeAddr("wrongRouter")),
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: true,
      feeUSDCents: 1,
      gasForVerification: 1,
      payloadSizeBytes: 1
    });

    vm.prank(OWNER);
    vm.expectRevert(VerifierTestHelper.MustUseTestRouter.selector);
    s_verifier.applyRemoteChainConfigUpdates(configs);
  }

  function test_applyAllowlistUpdates_RevertWhen_MustUseAllowlist() public {
    BaseVerifier.AllowlistConfigArgs[] memory configs = new BaseVerifier.AllowlistConfigArgs[](1);
    configs[0] = BaseVerifier.AllowlistConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      allowlistEnabled: false,
      addedAllowlistedSenders: new address[](0),
      removedAllowlistedSenders: new address[](0)
    });

    vm.prank(OWNER);
    vm.expectRevert(VerifierTestHelper.MustUseAllowlist.selector);
    s_verifier.applyAllowlistUpdates(configs);
  }

  function test_setAllowedFinalityConfig() public {
    bytes4 allowedFinality = FinalityCodec._encodeBlockDepth(42);

    vm.prank(OWNER);
    s_verifier.setAllowedFinalityConfig(allowedFinality);

    assertEq(s_verifier.getAllowedFinalityConfig(), allowedFinality);
  }
}
