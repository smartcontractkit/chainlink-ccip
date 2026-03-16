// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {OffRamp} from "../../../offRamp/OffRamp.sol";
import {MockReceiverV2} from "../../mocks/MockReceiverV2.sol";
import {OffRampSetup} from "./OffRampSetup.t.sol";

contract OffRamp_getCCVsFromReceiver is OffRampSetup {
  address internal s_userRequiredCCV;
  address internal s_optionalCcv1;
  address internal s_optionalCcv2;
  bytes internal s_sender;

  function setUp() public override {
    super.setUp();

    s_userRequiredCCV = makeAddr("userRequiredCCV");
    s_optionalCcv1 = makeAddr("optionalCcv1");
    s_optionalCcv2 = makeAddr("optionalCcv2");
    s_sender = abi.encodePacked(makeAddr("sender"));
  }

  function test_getCCVsFromReceiver_contractV2_usesReceiverValues() public {
    address[] memory userRequired = new address[](1);
    userRequired[0] = s_userRequiredCCV;

    address[] memory userOptional = new address[](2);
    userOptional[0] = s_optionalCcv1;
    userOptional[1] = s_optionalCcv2;

    uint8 optionalThresholdRequested = 1;

    MockReceiverV2 receiver = new MockReceiverV2(userRequired, userOptional, optionalThresholdRequested);

    (address[] memory requiredFromReceiver, address[] memory optionalFromReceiver, uint8 optionalThreshold) =
      s_offRamp.getCCVsFromReceiver(SOURCE_CHAIN_SELECTOR, address(receiver), s_sender, 0);

    assertEq(requiredFromReceiver.length, userRequired.length);
    assertEq(requiredFromReceiver[0], userRequired[0]);
    assertEq(optionalFromReceiver.length, userOptional.length);
    assertEq(optionalFromReceiver[0], userOptional[0]);
    assertEq(optionalFromReceiver[1], userOptional[1]);
    assertEq(optionalThreshold, optionalThresholdRequested);
  }

  function test_getCCVsFromReceiver_contractV2_fallsBackToDefaults_WhenEmptyValues() public {
    address[] memory emptyRequired = new address[](0);
    address[] memory emptyOptional = new address[](0);

    MockReceiverV2 receiver = new MockReceiverV2(emptyRequired, emptyOptional, 0);

    (address[] memory requiredFromReceiver, address[] memory optionalFromReceiver, uint8 optionalThresholdRequested) =
      s_offRamp.getCCVsFromReceiver(SOURCE_CHAIN_SELECTOR, address(receiver), s_sender, 0);

    assertEq(requiredFromReceiver.length, 1);
    assertEq(requiredFromReceiver[0], address(0));
    assertEq(optionalFromReceiver.length, 0);
    assertEq(optionalThresholdRequested, 0);
  }

  function test_getCCVsFromReceiver_noContract_fallsBackToDefaults() public {
    address eoa = makeAddr("eoa");

    (address[] memory requiredFromReceiver, address[] memory optionalFromReceiver, uint8 optionalThresholdRequested) =
      s_offRamp.getCCVsFromReceiver(SOURCE_CHAIN_SELECTOR, eoa, s_sender, 0);

    assertEq(requiredFromReceiver.length, 1);
    assertEq(requiredFromReceiver[0], address(0));
    assertEq(optionalFromReceiver.length, 0);
    assertEq(optionalThresholdRequested, 0);
  }

  function test_getCCVsFromReceiver_contractNoV2_fallsBackToDefaults() public {
    address contractAddress = makeAddr("contract");
    vm.etch(contractAddress, "some source code");

    (address[] memory requiredFromReceiver, address[] memory optionalFromReceiver, uint8 optionalThresholdRequested) =
      s_offRamp.getCCVsFromReceiver(SOURCE_CHAIN_SELECTOR, contractAddress, s_sender, 0);

    assertEq(requiredFromReceiver.length, 1);
    assertEq(requiredFromReceiver[0], address(0));
    assertEq(optionalFromReceiver.length, 0);
    assertEq(optionalThresholdRequested, 0);
  }

  function test_getCCVsFromReceiver_contractV2_succeedsWithFinality() public {
    address[] memory userRequired = new address[](1);
    userRequired[0] = s_userRequiredCCV;

    MockReceiverV2 receiver = new MockReceiverV2(userRequired, new address[](0), 0);

    (address[] memory requiredFromReceiver,,) =
      s_offRamp.getCCVsFromReceiver(SOURCE_CHAIN_SELECTOR, address(receiver), s_sender, 0);

    assertEq(requiredFromReceiver.length, 1);
    assertEq(requiredFromReceiver[0], s_userRequiredCCV);
  }

  function test_getCCVsFromReceiver_contractV2_FTF_succeedsWhenFinalityMeetsMinBlockConfirmations() public {
    address[] memory userRequired = new address[](1);
    userRequired[0] = s_userRequiredCCV;

    MockReceiverV2 receiver = new MockReceiverV2(userRequired, new address[](0), 0);
    receiver.setMinBlockConfirmations(10);

    (address[] memory requiredFromReceiver,,) =
      s_offRamp.getCCVsFromReceiver(SOURCE_CHAIN_SELECTOR, address(receiver), s_sender, 10);

    assertEq(requiredFromReceiver.length, 1);
    assertEq(requiredFromReceiver[0], s_userRequiredCCV);
  }

  function test_getCCVsFromReceiver_contractV2_FTF_succeedsWhenFinalityExceedsMinBlockConfirmations() public {
    address[] memory userRequired = new address[](1);
    userRequired[0] = s_userRequiredCCV;

    MockReceiverV2 receiver = new MockReceiverV2(userRequired, new address[](0), 0);
    receiver.setMinBlockConfirmations(5);

    (address[] memory requiredFromReceiver,,) =
      s_offRamp.getCCVsFromReceiver(SOURCE_CHAIN_SELECTOR, address(receiver), s_sender, 10);

    assertEq(requiredFromReceiver.length, 1);
    assertEq(requiredFromReceiver[0], s_userRequiredCCV);
  }

  function test_getCCVsFromReceiver_contractNoV2_fallsBackToDefaults_WhenFinalized() public {
    address contractAddress = makeAddr("contract");
    vm.etch(contractAddress, "some source code");

    (address[] memory requiredFromReceiver, address[] memory optionalFromReceiver, uint8 optionalThresholdRequested) =
      s_offRamp.getCCVsFromReceiver(SOURCE_CHAIN_SELECTOR, contractAddress, s_sender, 0);

    assertEq(requiredFromReceiver.length, 1);
    assertEq(requiredFromReceiver[0], address(0));
    assertEq(optionalFromReceiver.length, 0);
    assertEq(optionalThresholdRequested, 0);
  }

  function test_getCCVsFromReceiver_noContract_fallsBackToDefaults_WhenFinalized() public {
    address eoa = makeAddr("eoa");

    (address[] memory requiredFromReceiver, address[] memory optionalFromReceiver, uint8 optionalThresholdRequested) =
      s_offRamp.getCCVsFromReceiver(SOURCE_CHAIN_SELECTOR, eoa, s_sender, 0);

    assertEq(requiredFromReceiver.length, 1);
    assertEq(requiredFromReceiver[0], address(0));
    assertEq(optionalFromReceiver.length, 0);
    assertEq(optionalThresholdRequested, 0);
  }

  // Reverts

  function test_getCCVsFromReceiver_RevertWhen_InvalidOptionalThreshold() public {
    address[] memory userOptional = new address[](1);

    uint8 tooHighThreshold = uint8(userOptional.length) + 1;
    MockReceiverV2 receiver = new MockReceiverV2(new address[](1), userOptional, tooHighThreshold);

    vm.expectRevert(
      abi.encodeWithSelector(OffRamp.InvalidOptionalThreshold.selector, tooHighThreshold, userOptional.length)
    );
    s_offRamp.getCCVsFromReceiver(SOURCE_CHAIN_SELECTOR, address(receiver), s_sender, 0);
  }

  function test_getCCVsFromReceiver_RevertWhen_noContract_FTF() public {
    address eoa = makeAddr("eoa");

    uint16 ftfBlockDepth = 5;
    vm.expectRevert(abi.encodeWithSelector(OffRamp.InvalidFinalityForReceiver.selector, eoa, ftfBlockDepth, uint16(0)));
    s_offRamp.getCCVsFromReceiver(SOURCE_CHAIN_SELECTOR, eoa, s_sender, ftfBlockDepth);
  }

  function test_getCCVsFromReceiver_RevertWhen_contractNoV2_FTF() public {
    address contractAddress = makeAddr("contract");
    vm.etch(contractAddress, "some source code");

    uint16 ftfBlockDepth = 5;
    vm.expectRevert(
      abi.encodeWithSelector(OffRamp.InvalidFinalityForReceiver.selector, contractAddress, ftfBlockDepth, uint16(0))
    );
    s_offRamp.getCCVsFromReceiver(SOURCE_CHAIN_SELECTOR, contractAddress, s_sender, ftfBlockDepth);
  }

  function test_getCCVsFromReceiver_RevertWhen_contractV2_MinBlockConfirmationsZero_FTF() public {
    // V2 receiver that returns minBlockConfirmations=0 (requires finality) should reject FTF.
    address[] memory userRequired = new address[](1);
    userRequired[0] = s_userRequiredCCV;

    MockReceiverV2 receiver = new MockReceiverV2(userRequired, new address[](0), 0);

    uint16 ftfBlockDepth = 5;
    vm.expectRevert(
      abi.encodeWithSelector(OffRamp.InvalidFinalityForReceiver.selector, address(receiver), ftfBlockDepth, uint16(0))
    );
    s_offRamp.getCCVsFromReceiver(SOURCE_CHAIN_SELECTOR, address(receiver), s_sender, ftfBlockDepth);
  }

  function test_getCCVsFromReceiver_RevertWhen_contractV2_FTF_BelowMinBlockConfirmations() public {
    MockReceiverV2 receiver = new MockReceiverV2(new address[](0), new address[](0), 0);
    receiver.setMinBlockConfirmations(10);

    uint16 insufficientFinality = 5;
    vm.expectRevert(
      abi.encodeWithSelector(
        OffRamp.InvalidFinalityForReceiver.selector, address(receiver), insufficientFinality, uint16(10)
      )
    );
    s_offRamp.getCCVsFromReceiver(SOURCE_CHAIN_SELECTOR, address(receiver), s_sender, insufficientFinality);
  }
}
