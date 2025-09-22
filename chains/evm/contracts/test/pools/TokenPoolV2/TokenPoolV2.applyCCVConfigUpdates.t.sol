// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {TokenPool} from "../../../pools/TokenPool.sol";
import {TokenPoolV2} from "../../../pools/TokenPoolV2.sol";
import {Ownable2Step} from "@chainlink/contracts/src/v0.8/shared/access/Ownable2Step.sol";

import {TokenPoolV2Setup} from "./TokenPoolV2Setup.t.sol";

contract TokenPoolV2_applyCCVConfigUpdates is TokenPoolV2Setup {
  // Test helper addresses.
  address internal ccv1 = makeAddr("ccv1");
  address internal ccv2 = makeAddr("ccv2");

  function test_applyCCVConfigUpdate() public {
    // Prepare test data.
    address[] memory outboundCCVs = new address[](1);
    outboundCCVs[0] = ccv1;

    address[] memory inboundCCVs = new address[](1);
    inboundCCVs[0] = ccv2;

    TokenPoolV2.CCVConfigArg[] memory configArgs = new TokenPoolV2.CCVConfigArg[](1);
    configArgs[0] = TokenPoolV2.CCVConfigArg({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundCCVs: outboundCCVs,
      inboundCCVs: inboundCCVs
    });

    vm.expectEmit();
    emit TokenPoolV2.CCVConfigUpdated(DEST_CHAIN_SELECTOR, outboundCCVs, inboundCCVs);
    s_tokenPool.applyCCVConfigUpdates(configArgs);

    // Verify the configuration was stored correctly.
    address[] memory storedOutbound = s_tokenPool.getRequiredOutboundCCVs(DEST_CHAIN_SELECTOR, 0, "");
    address[] memory storedInbound = s_tokenPool.getRequiredInboundCCVs(DEST_CHAIN_SELECTOR, 0, "");

    assertEq(storedOutbound.length, 1);
    assertEq(storedOutbound[0], ccv1);

    assertEq(storedInbound.length, 1);
    assertEq(storedInbound[0], ccv2);
  }

  // Reverts

  function test_applyCCVConfigUpdate_RevertWhen_OnlyCallableByOwner() public {
    vm.stopPrank();
    vm.prank(STRANGER);

    TokenPoolV2.CCVConfigArg[] memory configArgs;

    vm.expectRevert(Ownable2Step.OnlyCallableByOwner.selector);
    s_tokenPool.applyCCVConfigUpdates(configArgs);
  }

  function test_applyCCVConfigUpdate_RevertWhen_DuplicateCCV_Outbound() public {
    TokenPoolV2.CCVConfigArg[] memory configArgs = new TokenPoolV2.CCVConfigArg[](1);
    address[] memory duplicateOutbound = new address[](3);
    duplicateOutbound[0] = ccv1;
    duplicateOutbound[1] = ccv2;
    duplicateOutbound[2] = ccv1; // Duplicate

    address[] memory validInbound = new address[](1);
    validInbound[0] = ccv2;

    configArgs[0] = TokenPoolV2.CCVConfigArg({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundCCVs: duplicateOutbound,
      inboundCCVs: validInbound
    });

    vm.expectRevert(abi.encodeWithSelector(TokenPoolV2.DuplicateCCV.selector, ccv1));
    s_tokenPool.applyCCVConfigUpdates(configArgs);
  }

  function test_applyCCVConfigUpdate_RevertWhen_DuplicateCCV_Inbound() public {
    TokenPoolV2.CCVConfigArg[] memory configArgs = new TokenPoolV2.CCVConfigArg[](1);
    address[] memory validOutbound = new address[](1);
    validOutbound[0] = ccv1;

    address[] memory duplicateInbound = new address[](3);
    duplicateInbound[0] = ccv2;
    duplicateInbound[1] = ccv1;
    duplicateInbound[2] = ccv2; // Duplicate

    configArgs[0] = TokenPoolV2.CCVConfigArg({
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      outboundCCVs: validOutbound,
      inboundCCVs: duplicateInbound
    });

    vm.expectRevert(abi.encodeWithSelector(TokenPoolV2.DuplicateCCV.selector, ccv2));
    s_tokenPool.applyCCVConfigUpdates(configArgs);
  }
}
