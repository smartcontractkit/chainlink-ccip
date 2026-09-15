// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {SuccinctZKVerifier} from "../../../ccvs/SuccinctZKVerifier.sol";
import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {FinalityCodec} from "../../../libraries/FinalityCodec.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {SuccinctZKVerifierSetup} from "./SuccinctZKVerifierSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

/// @notice Covers the functions that only forward to BaseVerifier or FeeTokenHandler, which have their own tests.
contract SuccinctZKVerifier_config is SuccinctZKVerifierSetup {
  function test_forwardToVerifier() public {
    (MessageV1Codec.MessageV1 memory message, bytes32 messageId) = _messageWithId();

    vm.stopPrank();
    vm.prank(s_onRamp);
    bytes memory verifierData = s_zkVerifier.forwardToVerifier(message, messageId, address(0), 0, "");

    assertEq(abi.encodePacked(VERSION_TAG_V0_0_1), verifierData);
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

  function test_setAllowedFinalityConfig() public {
    s_zkVerifier.setAllowedFinalityConfig(FinalityCodec.WAIT_FOR_FINALITY_FLAG);

    assertEq(FinalityCodec.WAIT_FOR_FINALITY_FLAG, s_zkVerifier.getAllowedFinalityConfig());
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
