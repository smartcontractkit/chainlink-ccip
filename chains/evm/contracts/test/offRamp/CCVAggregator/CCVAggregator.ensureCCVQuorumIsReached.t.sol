// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../../libraries/Internal.sol";
import {CCVAggregator} from "../../../offRamp/CCVAggregator.sol";
import {CCVAggregatorHelper, CCVAggregatorSetup} from "./CCVAggregatorSetup.t.sol";

contract CCVAggregator_ensureCCVQuorumIsReached is CCVAggregatorSetup {
  address internal s_receiver;
  address internal s_requiredCCV;
  address internal s_optionalCCV1;
  address internal s_optionalCCV2;
  address internal s_laneMandatedCCV;
  address internal s_poolRequiredCCV;

  function setUp() public override {
    CCVAggregatorSetup.setUp();

    s_receiver = makeAddr("receiver");
    s_requiredCCV = makeAddr("requiredCCV");
    s_optionalCCV1 = makeAddr("optionalCCV1");
    s_optionalCCV2 = makeAddr("optionalCCV2");
    s_laneMandatedCCV = makeAddr("laneMandatedCCV");
    s_poolRequiredCCV = makeAddr("poolRequiredCCV");

    // Configure source chain with lane mandated CCVs
    CCVAggregator.SourceChainConfigArgs[] memory configs = new CCVAggregator.SourceChainConfigArgs[](1);
    configs[0] = CCVAggregator.SourceChainConfigArgs({
      router: s_sourceRouter,
      sourceChainSelector: SOURCE_CHAIN_SELECTOR,
      isEnabled: true,
      onRamp: abi.encode(makeAddr("onRamp")),
      defaultCCV: new address[](1),
      laneMandatedCCVs: new address[](1)
    });
    configs[0].laneMandatedCCVs[0] = s_laneMandatedCCV;
    configs[0].defaultCCV[0] = s_defaultCCV;

    s_agg.applySourceChainConfigUpdates(configs);
  }

  function test_ensureCCVQuorumIsReached_Success_AllCCVsFound() public {
    address[] memory ccvs = new address[](5);
    ccvs[0] = s_requiredCCV;
    ccvs[1] = s_optionalCCV1;
    ccvs[2] = s_laneMandatedCCV;
    ccvs[3] = s_poolRequiredCCV;
    ccvs[4] = s_defaultCCV;

    address[] memory tokenRequiredCCVs = new address[](1);
    tokenRequiredCCVs[0] = s_poolRequiredCCV;

    // Mock receiver to return no required CCVs (so it falls back to defaults)
    vm.mockCall(
      s_receiver,
      abi.encodeWithSignature("getCCVs(uint64)", SOURCE_CHAIN_SELECTOR),
      abi.encode(
        new address[](0), // no required CCVs - will fall back to defaults
        new address[](0), // no optional CCVs
        uint8(0) // no optional threshold
      )
    );

    (address[] memory ccvsToQuery, uint256[] memory dataIndexes) =
      s_agg.ensureCCVQuorumIsReached(SOURCE_CHAIN_SELECTOR, s_receiver, ccvs, tokenRequiredCCVs);

    // The function only returns CCVs that are actually needed and found
    // Since we have 1 required (from receiver), 1 lane mandated, and 1 pool required
    // The default CCV is not used when receiver specifies required CCVs
    assertEq(ccvsToQuery.length, 3);
    assertEq(dataIndexes.length, 3);
  }

  function test_ensureCCVQuorumIsReached_RevertWhen_RequiredCCVMissing_Receiver() public {
    address[] memory ccvs = new address[](2);
    ccvs[0] = s_optionalCCV1;
    ccvs[1] = s_laneMandatedCCV;

    address[] memory tokenRequiredCCVs = new address[](0);

    // Mock receiver to return a required CCV that's not in the provided CCVs
    address[] memory receiverRequired = new address[](1);
    receiverRequired[0] = s_requiredCCV;
    vm.mockCall(
      s_receiver,
      abi.encodeWithSignature("getCCVs(uint64)", SOURCE_CHAIN_SELECTOR),
      abi.encode(receiverRequired, new address[](0), uint8(0))
    );

    vm.expectRevert(abi.encodeWithSelector(CCVAggregator.RequiredCCVMissing.selector, s_defaultCCV, false));
    s_agg.ensureCCVQuorumIsReached(SOURCE_CHAIN_SELECTOR, s_receiver, ccvs, tokenRequiredCCVs);
  }

  function test_ensureCCVQuorumIsReached_RevertWhen_RequiredCCVMissing_Pool() public {
    address[] memory ccvs = new address[](3);
    ccvs[0] = s_requiredCCV;
    ccvs[1] = s_laneMandatedCCV;
    ccvs[2] = s_defaultCCV;

    address[] memory tokenRequiredCCVs = new address[](1);
    tokenRequiredCCVs[0] = s_poolRequiredCCV;

    // Mock receiver to return required CCVs that are found
    address[] memory receiverRequired = new address[](1);
    receiverRequired[0] = s_requiredCCV;
    vm.mockCall(
      s_receiver,
      abi.encodeWithSignature("getCCVs(uint64)", SOURCE_CHAIN_SELECTOR),
      abi.encode(receiverRequired, new address[](0), uint8(0))
    );

    vm.expectRevert(abi.encodeWithSelector(CCVAggregator.RequiredCCVMissing.selector, s_poolRequiredCCV, true));
    s_agg.ensureCCVQuorumIsReached(SOURCE_CHAIN_SELECTOR, s_receiver, ccvs, tokenRequiredCCVs);
  }

  function test_ensureCCVQuorumIsReached_RevertWhen_RequiredCCVMissing_LaneMandated() public {
    address[] memory ccvs = new address[](3);
    ccvs[0] = s_requiredCCV;
    ccvs[1] = s_optionalCCV1;
    ccvs[2] = s_defaultCCV;

    address[] memory tokenRequiredCCVs = new address[](0);

    // Mock receiver to return required CCVs that are found
    address[] memory receiverRequired = new address[](1);
    receiverRequired[0] = s_requiredCCV;
    vm.mockCall(
      s_receiver,
      abi.encodeWithSignature("getCCVs(uint64)", SOURCE_CHAIN_SELECTOR),
      abi.encode(receiverRequired, new address[](0), uint8(0))
    );

    vm.expectRevert(abi.encodeWithSelector(CCVAggregator.RequiredCCVMissing.selector, s_laneMandatedCCV, true));
    s_agg.ensureCCVQuorumIsReached(SOURCE_CHAIN_SELECTOR, s_receiver, ccvs, tokenRequiredCCVs);
  }

  function test_ensureCCVQuorumIsReached_RevertWhen_OptionalCCVQuorumNotReached() public {
    address[] memory ccvs = new address[](3);
    ccvs[0] = s_requiredCCV;
    ccvs[1] = s_optionalCCV1;
    ccvs[2] = s_laneMandatedCCV;

    address[] memory tokenRequiredCCVs = new address[](0);

    // Mock receiver to return required CCVs and optional CCVs with threshold 2
    address[] memory receiverRequired = new address[](1);
    receiverRequired[0] = s_requiredCCV;
    address[] memory receiverOptional = new address[](2);
    receiverOptional[0] = s_optionalCCV1;
    receiverOptional[1] = s_optionalCCV2;

    vm.mockCall(
      s_receiver,
      abi.encodeWithSignature("getCCVs(uint64)", SOURCE_CHAIN_SELECTOR),
      abi.encode(receiverRequired, receiverOptional, uint8(2))
    );

    vm.expectRevert(abi.encodeWithSelector(CCVAggregator.RequiredCCVMissing.selector, s_defaultCCV, false));
    s_agg.ensureCCVQuorumIsReached(SOURCE_CHAIN_SELECTOR, s_receiver, ccvs, tokenRequiredCCVs);
  }

  function test_ensureCCVQuorumIsReached_Success_OptionalCCVsFound() public {
    address[] memory ccvs = new address[](5);
    ccvs[0] = s_requiredCCV;
    ccvs[1] = s_optionalCCV1;
    ccvs[2] = s_optionalCCV2;
    ccvs[3] = s_laneMandatedCCV;
    ccvs[4] = s_defaultCCV;

    address[] memory tokenRequiredCCVs = new address[](0);

    // Mock receiver to return required CCVs and optional CCVs with threshold 2
    address[] memory receiverRequired = new address[](1);
    receiverRequired[0] = s_requiredCCV;
    address[] memory receiverOptional = new address[](2);
    receiverOptional[0] = s_optionalCCV1;
    receiverOptional[1] = s_optionalCCV2;

    vm.mockCall(
      s_receiver,
      abi.encodeWithSignature("getCCVs(uint64)", SOURCE_CHAIN_SELECTOR),
      abi.encode(receiverRequired, receiverOptional, uint8(2))
    );

    (address[] memory ccvsToQuery, uint256[] memory dataIndexes) =
      s_agg.ensureCCVQuorumIsReached(SOURCE_CHAIN_SELECTOR, s_receiver, ccvs, tokenRequiredCCVs);

    // The function only returns CCVs that are actually needed and found
    // Since we have 1 required (from receiver) + 1 lane mandated + 2 optional (threshold met)
    // The default CCV is not used when receiver specifies required CCVs
    assertEq(ccvsToQuery.length, 2);
    assertEq(dataIndexes.length, 2);
  }
}
