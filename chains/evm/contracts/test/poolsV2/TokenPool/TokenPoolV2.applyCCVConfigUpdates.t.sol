// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenPool} from "../../../poolsV2/TokenPool.sol";
import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

contract TokenPoolV2_applyCCVConfigUpdates is TokenPoolV2Setup {
  // Test helper addresses.
  address internal s_ccv1 = makeAddr("ccv1");
  address internal s_ccv2 = makeAddr("ccv2");

  function test_applyCCVConfigUpdate() public {
    // Prepare test data.
    address[] memory outboundCCVs = new address[](1);
    outboundCCVs[0] = s_ccv1;

    address[] memory inboundCCVs = new address[](1);
    inboundCCVs[0] = s_ccv2;

    TokenPool.CCVConfigArg[] memory configArgs = new TokenPool.CCVConfigArg[](1);
    configArgs[0] = TokenPool.CCVConfigArg({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundCCVs: outboundCCVs,
      inboundCCVs: inboundCCVs
    });

    vm.expectEmit();
    emit TokenPool.CCVConfigUpdated(DEST_CHAIN_SELECTOR, outboundCCVs, inboundCCVs);
    s_tokenPool.applyCCVConfigUpdates(configArgs);

    // Verify the configuration was stored correctly.
    address[] memory storedOutbound = s_tokenPool.getRequiredOutboundCCVs(address(0), DEST_CHAIN_SELECTOR, 0, 0, "");
    address[] memory storedInbound = s_tokenPool.getRequiredInboundCCVs(address(0), DEST_CHAIN_SELECTOR, 0, 0, "");

    assertEq(storedOutbound.length, outboundCCVs.length);
    assertEq(storedOutbound[0], outboundCCVs[0]);

    assertEq(storedInbound.length, inboundCCVs.length);
    assertEq(storedInbound[0], inboundCCVs[0]);
  }

  // Reverts

  function test_applyCCVConfigUpdate_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();
    vm.prank(STRANGER);

    TokenPool.CCVConfigArg[] memory configArgs;

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_tokenPool.applyCCVConfigUpdates(configArgs);
  }

  function test_applyCCVConfigUpdate_RevertWhen_DuplicateCCV_Outbound() public {
    TokenPool.CCVConfigArg[] memory configArgs = new TokenPool.CCVConfigArg[](1);
    address[] memory duplicateOutbound = new address[](3);
    duplicateOutbound[0] = s_ccv1;
    duplicateOutbound[1] = s_ccv2;
    duplicateOutbound[2] = s_ccv1; // Duplicate

    address[] memory validInbound = new address[](1);
    validInbound[0] = s_ccv2;

    configArgs[0] = TokenPool.CCVConfigArg({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundCCVs: duplicateOutbound,
      inboundCCVs: validInbound
    });

    vm.expectRevert(abi.encodeWithSelector(TokenPool.DuplicateCCV.selector, s_ccv1));
    s_tokenPool.applyCCVConfigUpdates(configArgs);
  }

  function test_applyCCVConfigUpdate_RevertWhen_DuplicateCCV_Inbound() public {
    TokenPool.CCVConfigArg[] memory configArgs = new TokenPool.CCVConfigArg[](1);
    address[] memory validOutbound = new address[](1);
    validOutbound[0] = s_ccv1;

    address[] memory duplicateInbound = new address[](3);
    duplicateInbound[0] = s_ccv2;
    duplicateInbound[1] = s_ccv1;
    duplicateInbound[2] = s_ccv2; // Duplicate

    configArgs[0] = TokenPool.CCVConfigArg({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundCCVs: validOutbound,
      inboundCCVs: duplicateInbound
    });

    vm.expectRevert(abi.encodeWithSelector(TokenPool.DuplicateCCV.selector, s_ccv2));
    s_tokenPool.applyCCVConfigUpdates(configArgs);
  }
}
