// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAny2EVMMessageReceiverV2} from "../../../interfaces/IAny2EVMMessageReceiverV2.sol";
import {CCVAggregator} from "../../../offRamp/CCVAggregator.sol";

import {MockReceiverV2} from "../../mocks/MockReceiverV2.sol";
import {CCVAggregatorSetup} from "./CCVAggregatorSetup.t.sol";

contract CCVAggregator_getCCVsFromReceiver is CCVAggregatorSetup {
  address[] internal s_defaultCCVs;
  address[] internal s_laneMandatedCCVs;

  function setUp() public override {
    CCVAggregatorSetup.setUp();
    s_defaultCCVs = new address[](1);
    s_laneMandatedCCVs = new address[](1);
    s_defaultCCVs[0] = makeAddr("defaultCCV");
    s_laneMandatedCCVs[0] = makeAddr("laneMandatedCCV");
    _applySourceConfig(
      s_destRouter, SOURCE_CHAIN_SELECTOR, abi.encode(makeAddr("onRamp")), true, s_defaultCCVs, s_laneMandatedCCVs
    );
  }

  function test_getCCVsFromReceiver_contractV2_usesReceiverValuesAndRequired() public {
    address[] memory requiredFromReceiver = new address[](1);
    requiredFromReceiver[0] = makeAddr("userRequiredCCV");
    address[] memory optionalFromReceiver = new address[](2);
    optionalFromReceiver[0] = makeAddr("optionalCcv1");
    optionalFromReceiver[1] = makeAddr("optionalCcv2");
    uint8 optionalThresholdRequested = 1;

    MockReceiverV2 receiver = new MockReceiverV2(requiredFromReceiver, optionalFromReceiver, optionalThresholdRequested);

    (address[] memory actualRequired, address[] memory actualOptional, uint8 actualThreshold) =
      s_agg.getCCVsFromReceiver(SOURCE_CHAIN_SELECTOR, address(receiver));

    // Receiver-provided CCVs are used as-is when present
    assertEq(actualRequired.length, 1);
    assertEq(actualRequired[0], requiredFromReceiver[0]);
    assertEq(actualOptional.length, 2);
    assertEq(actualOptional[0], optionalFromReceiver[0]);
    assertEq(actualOptional[1], optionalFromReceiver[1]);
    assertEq(actualThreshold, optionalThresholdRequested);
  }

  function test_getCCVsFromReceiver_success_contractNoV2_fallsBackToDefaults() public {
    // Reconfigure to have no required CCV on the source chain
    _applySourceConfig(
      s_destRouter, SOURCE_CHAIN_SELECTOR, abi.encode(makeAddr("onRamp")), true, s_defaultCCVs, s_laneMandatedCCVs
    );

    // Use a contract address and mock V2 support so the function takes the fallback branch (no getCCVs data)
    address receiver = makeAddr("contractWithoutV2");
    vm.etch(receiver, bytes("fake"));
    vm.mockCall(
      receiver,
      abi.encodeWithSignature("supportsInterface(bytes4)", type(IAny2EVMMessageReceiverV2).interfaceId),
      abi.encode(true)
    );

    (address[] memory actualRequired, address[] memory actualOptional, uint8 actualThreshold) =
      s_agg.getCCVsFromReceiver(SOURCE_CHAIN_SELECTOR, receiver);

    assertEq(actualRequired.length, 1);
    assertEq(actualRequired[0], s_defaultCCVs[0]);
    assertEq(actualOptional.length, 0);
    assertEq(actualThreshold, 0);
  }

  function test_getCCVsFromReceiver_revert_InvalidOptionalThreshold() public {
    address[] memory requiredFromReceiver = new address[](0);
    address[] memory optionalFromReceiver = new address[](1);
    optionalFromReceiver[0] = makeAddr("optionalCcv1");
    uint8 optionalThresholdRequested = 2; // exceeds length 1

    MockReceiverV2 receiver = new MockReceiverV2(requiredFromReceiver, optionalFromReceiver, optionalThresholdRequested);

    vm.expectRevert(abi.encodeWithSelector(CCVAggregator.InvalidOptionalThreshold.selector, 1, 2));
    s_agg.getCCVsFromReceiver(SOURCE_CHAIN_SELECTOR, address(receiver));
  }
}
