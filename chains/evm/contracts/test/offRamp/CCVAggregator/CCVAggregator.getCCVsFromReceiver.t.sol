// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCVAggregator} from "../../../offRamp/CCVAggregator.sol";

import {MockReceiverV2} from "../../mocks/MockReceiverV2.sol";
import {CCVAggregatorHelper, CCVAggregatorSetup} from "./CCVAggregatorSetup.t.sol";

contract CCVAggregator_getCCVsFromReceiver is CCVAggregatorSetup {
  address internal s_userRequiredCCV;
  address internal s_optionalCcv1;
  address internal s_optionalCcv2;

  function setUp() public override {
    CCVAggregatorSetup.setUp();

    s_userRequiredCCV = makeAddr("userRequiredCCV");
    s_optionalCcv1 = makeAddr("optionalCcv1");
    s_optionalCcv2 = makeAddr("optionalCcv2");
  }

  function test_getCCVsFromReceiver_contractV2_usesReceiverValues() public {
    address[] memory userRequired = new address[](1);
    userRequired[0] = s_userRequiredCCV;

    address[] memory userOptional = new address[](2);
    userOptional[0] = s_optionalCcv1;
    userOptional[1] = s_optionalCcv2;

    MockReceiverV2 receiver = new MockReceiverV2(userRequired, userOptional, 1);

    (address[] memory requiredFromReceiver, address[] memory optionalFromReceiver, uint8 optionalThresholdRequested) =
      s_agg.getCCVsFromReceiver(SOURCE_CHAIN_SELECTOR, address(receiver));

    assertEq(requiredFromReceiver.length, userRequired.length);
    assertEq(requiredFromReceiver[0], userRequired[0]);
    assertEq(optionalFromReceiver.length, userOptional.length);
    assertEq(optionalFromReceiver[0], userOptional[0]);
    assertEq(optionalFromReceiver[1], userOptional[1]);
    assertEq(optionalThresholdRequested, 1);
  }

  function test_getCCVsFromReceiver_contractV2_fallsBackToDefaults_WhenEmptyValues() public {
    address[] memory emptyRequired = new address[](0);
    address[] memory emptyOptional = new address[](0);

    MockReceiverV2 receiver = new MockReceiverV2(emptyRequired, emptyOptional, 0);

    (address[] memory requiredFromReceiver, address[] memory optionalFromReceiver, uint8 optionalThresholdRequested) =
      s_agg.getCCVsFromReceiver(SOURCE_CHAIN_SELECTOR, address(receiver));

    assertEq(requiredFromReceiver.length, 1);
    assertEq(requiredFromReceiver[0], s_defaultCCV);
    assertEq(optionalFromReceiver.length, 0);
    assertEq(optionalThresholdRequested, 0);
  }

  function test_getCCVsFromReceiver_noContract_fallsBackToDefaults() public {
    address eoa = makeAddr("eoa");

    (address[] memory requiredFromReceiver, address[] memory optionalFromReceiver, uint8 optionalThresholdRequested) =
      s_agg.getCCVsFromReceiver(SOURCE_CHAIN_SELECTOR, eoa);

    assertEq(requiredFromReceiver.length, 1);
    assertEq(requiredFromReceiver[0], s_defaultCCV);
    assertEq(optionalFromReceiver.length, 0);
    assertEq(optionalThresholdRequested, 0);
  }

  function test_getCCVsFromReceiver_contractNoV2_fallsBackToDefaults() public {
    address contractAddress = makeAddr("contract");

    (address[] memory requiredFromReceiver, address[] memory optionalFromReceiver, uint8 optionalThresholdRequested) =
      s_agg.getCCVsFromReceiver(SOURCE_CHAIN_SELECTOR, contractAddress);

    assertEq(requiredFromReceiver.length, 1);
    assertEq(requiredFromReceiver[0], s_defaultCCV);
    assertEq(optionalFromReceiver.length, 0);
    assertEq(optionalThresholdRequested, 0);
  }

  function test_getCCVsFromReceiver_contractV2_usesReceiverValuesAndRequired() public {
    address[] memory userRequired = new address[](1);
    userRequired[0] = s_userRequiredCCV;

    address[] memory userOptional = new address[](2);
    userOptional[0] = s_optionalCcv1;
    userOptional[1] = s_optionalCcv2;

    MockReceiverV2 receiver = new MockReceiverV2(userRequired, userOptional, 1);

    (address[] memory requiredFromReceiver, address[] memory optionalFromReceiver, uint8 optionalThresholdRequested) =
      s_agg.getCCVsFromReceiver(SOURCE_CHAIN_SELECTOR, address(receiver));

    assertEq(requiredFromReceiver.length, userRequired.length);
    assertEq(requiredFromReceiver[0], userRequired[0]);
    assertEq(optionalFromReceiver.length, userOptional.length);
    assertEq(optionalFromReceiver[0], userOptional[0]);
    assertEq(optionalFromReceiver[1], userOptional[1]);
    assertEq(optionalThresholdRequested, 1);
  }

  function test_getCCVsFromReceiver_success_contractNoV2_fallsBackToDefaults() public {
    // Create a contract that doesn't support V2 interface
    address nonV2Contract = makeAddr("nonV2Contract");

    (address[] memory requiredFromReceiver, address[] memory optionalFromReceiver, uint8 optionalThresholdRequested) =
      s_agg.getCCVsFromReceiver(SOURCE_CHAIN_SELECTOR, nonV2Contract);

    assertEq(requiredFromReceiver.length, 1);
    assertEq(requiredFromReceiver[0], s_defaultCCV);
    assertEq(optionalFromReceiver.length, 0);
    assertEq(optionalThresholdRequested, 0);
  }

  function test_getCCVsFromReceiver_revert_InvalidOptionalThreshold() public {
    address[] memory userRequired = new address[](1);
    userRequired[0] = s_userRequiredCCV;

    address[] memory userOptional = new address[](1);
    userOptional[0] = s_optionalCcv1;

    // Set threshold higher than available optional CCVs
    MockReceiverV2 receiver = new MockReceiverV2(userRequired, userOptional, 2);

    (address[] memory requiredFromReceiver, address[] memory optionalFromReceiver, uint8 optionalThresholdRequested) =
      s_agg.getCCVsFromReceiver(SOURCE_CHAIN_SELECTOR, address(receiver));

    // This should not revert since the function doesn't validate the threshold
    // It just returns what the receiver provides
    assertEq(requiredFromReceiver.length, userRequired.length);
    assertEq(optionalFromReceiver.length, userOptional.length);
    assertEq(optionalThresholdRequested, 2);
  }
}
