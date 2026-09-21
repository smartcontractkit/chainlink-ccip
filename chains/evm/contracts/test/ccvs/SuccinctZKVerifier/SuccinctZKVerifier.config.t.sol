// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SuccinctZKVerifier} from "../../../ccvs/SuccinctZKVerifier.sol";
import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {Client} from "../../../libraries/Client.sol";
import {FinalityCodec} from "../../../libraries/FinalityCodec.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {SuccinctZKVerifierSetup} from "./SuccinctZKVerifierSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract SuccinctZKVerifier_config is SuccinctZKVerifierSetup {
  function test_forwardToVerifier() public {
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _createMessageWithId();

    vm.stopPrank();
    vm.prank(s_onRamp);
    bytes memory verifierData = s_zkVerifier.forwardToVerifier(message, messageId, address(0), 0, "");

    assertEq(abi.encodePacked(VERSION_TAG_V0_0_1), verifierData);
  }

  function test_forwardToVerifier_RevertWhen_CursedByRMN() public {
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _createMessageWithId();
    _setMockRMNChainCurse(DEST_CHAIN_SELECTOR, true);

    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.CursedByRMN.selector, DEST_CHAIN_SELECTOR));
    s_zkVerifier.forwardToVerifier(message, messageId, address(0), 0, "");
  }

  function test_forwardToVerifier_RevertWhen_SenderNotAllowed() public {
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _createMessageWithId();
    BaseVerifier.AllowlistConfigArgs[] memory allowlistConfigs = new BaseVerifier.AllowlistConfigArgs[](1);
    allowlistConfigs[0] = _getAllowlistConfig(DEST_CHAIN_SELECTOR, true, new address[](0), new address[](0));
    s_zkVerifier.applyAllowlistUpdates(allowlistConfigs);

    vm.stopPrank();
    vm.prank(s_onRamp);
    vm.expectRevert(
      abi.encodeWithSelector(BaseVerifier.SenderNotAllowed.selector, abi.decode(message.sender, (address)))
    );
    s_zkVerifier.forwardToVerifier(message, messageId, address(0), 0, "");
  }

  function test_getFee_RevertWhen_FinalityNotRequested() public {
    Client.EVM2AnyMessage memory message;

    // SP1Helios only proves finalized blocks.
    vm.expectRevert(
      abi.encodeWithSelector(
        FinalityCodec.InvalidRequestedFinality.selector,
        FinalityCodec._encodeBlockDepth(1),
        FinalityCodec.WAIT_FOR_FINALITY_FLAG
      )
    );
    s_zkVerifier.getFee(DEST_CHAIN_SELECTOR, message, "", FinalityCodec._encodeBlockDepth(1));
  }

  function test_setDynamicConfig() public {
    address feeAggregator = makeAddr("feeAggregator");

    s_zkVerifier.setDynamicConfig(SuccinctZKVerifier.DynamicConfig({feeAggregator: feeAggregator}));

    assertEq(feeAggregator, s_zkVerifier.getDynamicConfig().feeAggregator);
  }

  function test_applyRemoteChainConfigUpdates() public {
    BaseVerifier.RemoteChainConfigArgs[] memory remoteChainConfigs = new BaseVerifier.RemoteChainConfigArgs[](1);
    remoteChainConfigs[0] = _getRemoteChainConfig(s_router, SOURCE_CHAIN_SELECTOR, true);

    s_zkVerifier.applyRemoteChainConfigUpdates(remoteChainConfigs);

    (BaseVerifier.RemoteChainConfigArgs memory config,) = s_zkVerifier.getRemoteChainConfig(SOURCE_CHAIN_SELECTOR);
    assertEq(address(s_router), address(config.router));
  }

  function test_applyAllowlistUpdates() public {
    address[] memory senders = new address[](1);
    senders[0] = makeAddr("sender");
    BaseVerifier.AllowlistConfigArgs[] memory allowlistConfigs = new BaseVerifier.AllowlistConfigArgs[](1);
    allowlistConfigs[0] = _getAllowlistConfig(DEST_CHAIN_SELECTOR, true, senders, new address[](0));

    s_zkVerifier.applyAllowlistUpdates(allowlistConfigs);

    (, address[] memory allowedSenders) = s_zkVerifier.getRemoteChainConfig(DEST_CHAIN_SELECTOR);
    assertEq(senders, allowedSenders);
  }

  function test_updateStorageLocations() public {
    string[] memory newStorageLocations = new string[](1);
    newStorageLocations[0] = "new/location";

    s_zkVerifier.updateStorageLocations(newStorageLocations);

    assertEq(newStorageLocations, s_zkVerifier.getStorageLocations());
  }

  function test_withdrawFeeTokens() public {
    uint256 feeAmount = 1000 ether;
    deal(s_sourceFeeToken, address(s_zkVerifier), feeAmount);
    address[] memory feeTokens = new address[](1);
    feeTokens[0] = s_sourceFeeToken;

    s_zkVerifier.withdrawFeeTokens(feeTokens);

    assertEq(feeAmount, IERC20(s_sourceFeeToken).balanceOf(FEE_AGGREGATOR));
  }
}
