// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../../libraries/Client.sol";
import {OnRamp} from "../../../onRamp/OnRamp.sol";
import {OnRampTestHelper} from "../../helpers/OnRampTestHelper.sol";
import {OnRampSetup} from "./OnRampSetup.t.sol";

contract OnRamp_mergeCCVLists is OnRampSetup {
  OnRampTestHelper internal s_onRampTestHelper;

  function _setupTestDestChainConfig(
    address[] memory laneMandatedCCVs
  ) internal {
    OnRamp.DestChainConfigArgs[] memory destChainConfigArgs = new OnRamp.DestChainConfigArgs[](1);
    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = makeAddr("defaultCCV");

    destChainConfigArgs[0] = OnRamp.DestChainConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      router: s_sourceRouter,
      addressBytesLength: EVM_ADDRESS_LENGTH,
      baseExecutionGasCost: BASE_EXEC_GAS_COST,
      laneMandatedCCVs: laneMandatedCCVs,
      defaultCCVs: defaultCCVs,
      defaultExecutor: makeAddr("defaultExecutor"),
      offRamp: abi.encodePacked(address(s_offRampOnRemoteChain))
    });

    s_onRampTestHelper.applyDestChainConfigUpdates(destChainConfigArgs);
  }

  function setUp() public override {
    super.setUp();
    s_onRampTestHelper = new OnRampTestHelper(
      OnRamp.StaticConfig({
        chainSelector: SOURCE_CHAIN_SELECTOR,
        rmnRemote: s_mockRMNRemote,
        tokenAdminRegistry: address(s_tokenAdminRegistry)
      }),
      OnRamp.DynamicConfig({
        feeQuoter: address(s_feeQuoter),
        reentrancyGuardEntered: false,
        feeAggregator: FEE_AGGREGATOR
      })
    );

    _setupTestDestChainConfig(new address[](0));
  }

  function test_mergeCCVLists_SkipsDuplicatesInPoolRequiredCCV() public {
    address poolCCV1 = makeAddr("poolCCV1");
    address poolCCV2 = makeAddr("poolCCV2");

    // Setup pool CCVs with duplicates.
    address[] memory poolRequiredCCV = new address[](3);
    poolRequiredCCV[0] = poolCCV1;
    poolRequiredCCV[1] = poolCCV1; // Duplicate
    poolRequiredCCV[2] = poolCCV2;

    address[] memory userCCVAddresses = new address[](0);
    bytes[] memory userCCVArgs = new bytes[](0);

    (address[] memory newCCVAddresses, bytes[] memory newCCVArgs) =
      s_onRampTestHelper.mergeCCVLists(userCCVAddresses, userCCVArgs, new address[](0), poolRequiredCCV);

    // Should only add unique pool CCVs.
    address[] memory expectedAddresses = new address[](2);
    expectedAddresses[0] = poolCCV1;
    expectedAddresses[1] = poolCCV2;
    bytes[] memory expectedArgs = new bytes[](2);
    expectedArgs[0] = "";
    expectedArgs[1] = "";
    _assertCCVArraysEqual(newCCVAddresses, newCCVArgs, expectedAddresses, expectedArgs);
  }

  function test_mergeCCVLists_NoChangesWhenPoolCCVAlreadyInRequired() public {
    address requiredCCV1 = makeAddr("requiredCCV1");

    address[] memory poolRequiredCCV = new address[](1);
    poolRequiredCCV[0] = requiredCCV1; // Already in required

    address[] memory userCCVAddresses = new address[](1);
    userCCVAddresses[0] = requiredCCV1;
    bytes[] memory userCCVArgs = new bytes[](1);
    userCCVArgs[0] = "required1";

    (address[] memory newCCVAddresses, bytes[] memory newCCVArgs) =
      s_onRampTestHelper.mergeCCVLists(userCCVAddresses, userCCVArgs, new address[](0), poolRequiredCCV);

    // Should return original arrays unchanged.
    _assertCCVArraysEqual(newCCVAddresses, newCCVArgs, userCCVAddresses, userCCVArgs);
  }

  function test_mergeCCVLists_PoolFallbackDefaults_UsesDefaults() public {
    address[] memory defaultCCVs = new address[](2);
    defaultCCVs[0] = makeAddr("defaultCCV1");
    defaultCCVs[1] = makeAddr("defaultCCV2");

    address[] memory userCCVAddresses = new address[](1);
    userCCVAddresses[0] = makeAddr("requiredCCV1");
    bytes[] memory userCCVArgs = new bytes[](1);
    userCCVArgs[0] = "required1";

    address[] memory poolRequiredCCV = new address[](defaultCCVs.length);
    poolRequiredCCV[0] = defaultCCVs[0];
    poolRequiredCCV[1] = defaultCCVs[1];

    (address[] memory newCCVAddresses, bytes[] memory newCCVArgs) =
      s_onRampTestHelper.mergeCCVLists(userCCVAddresses, userCCVArgs, new address[](0), poolRequiredCCV);

    address[] memory expectedAddresses = new address[](3);
    expectedAddresses[0] = userCCVAddresses[0];
    expectedAddresses[1] = defaultCCVs[0];
    expectedAddresses[2] = defaultCCVs[1];
    bytes[] memory expectedArgs = new bytes[](3);
    expectedArgs[0] = userCCVArgs[0];
    expectedArgs[1] = "";
    expectedArgs[2] = "";

    // Should return original arrays unchanged.
    _assertCCVArraysEqual(newCCVAddresses, newCCVArgs, expectedAddresses, expectedArgs);
  }

  function test_mergeCCVLists_NoPoolProcessing_KeepsUserAndLaneOnly() public {
    address[] memory userCCVAddresses = new address[](1);
    userCCVAddresses[0] = makeAddr("userCCV");
    bytes[] memory userCCVArgs = new bytes[](1);
    userCCVArgs[0] = "userArgs";

    address[] memory laneMandatedCCVs = new address[](1);
    laneMandatedCCVs[0] = makeAddr("laneCCV");

    (address[] memory newCCVAddresses, bytes[] memory newCCVArgs) =
      s_onRampTestHelper.mergeCCVLists(userCCVAddresses, userCCVArgs, laneMandatedCCVs, new address[](0));

    address[] memory expectedAddresses = new address[](2);
    expectedAddresses[0] = userCCVAddresses[0];
    expectedAddresses[1] = laneMandatedCCVs[0];
    bytes[] memory expectedArgs = new bytes[](2);
    expectedArgs[0] = userCCVArgs[0];
    expectedArgs[1] = "";
    _assertCCVArraysEqual(newCCVAddresses, newCCVArgs, expectedAddresses, expectedArgs);
  }

  function test_mergeCCVLists_DedupUserAndMandatoryCCVs() public {
    address requiredCCV1 = makeAddr("requiredCCV1");

    address[] memory laneMandatedCCVs = new address[](2);
    laneMandatedCCVs[0] = makeAddr("laneMandatedCCV1");
    laneMandatedCCVs[1] = requiredCCV1; // This one is also in user specified list

    address[] memory poolRequiredCCV = new address[](1);
    poolRequiredCCV[0] = makeAddr("poolCCV1");

    address[] memory userCCVAddresses = new address[](1);
    userCCVAddresses[0] = requiredCCV1;
    bytes[] memory userCCVArgs = new bytes[](1);
    userCCVArgs[0] = "required1";

    (address[] memory newCCVAddresses, bytes[] memory newCCVArgs) =
      s_onRampTestHelper.mergeCCVLists(userCCVAddresses, userCCVArgs, laneMandatedCCVs, poolRequiredCCV);

    address[] memory expectedAddresses = new address[](3);
    expectedAddresses[0] = requiredCCV1;
    expectedAddresses[1] = laneMandatedCCVs[0];
    expectedAddresses[2] = poolRequiredCCV[0];
    bytes[] memory expectedArgs = new bytes[](3);
    expectedArgs[0] = userCCVArgs[0];
    expectedArgs[1] = "";
    expectedArgs[2] = "";
    _assertCCVArraysEqual(newCCVAddresses, newCCVArgs, expectedAddresses, expectedArgs);
  }

  function test_mergeCCVLists_DedupUserAndPoolCCVs() public {
    // Setup both lane mandated and pool required CCVs
    address requiredCCV1 = makeAddr("requiredCCV1");

    address[] memory poolRequiredCCV = new address[](1);
    poolRequiredCCV[0] = requiredCCV1; // This one is also in user specified list

    address[] memory userCCVAddresses = new address[](1);
    userCCVAddresses[0] = requiredCCV1;
    bytes[] memory userCCVArgs = new bytes[](1);
    userCCVArgs[0] = "required1";

    (address[] memory newCCVAddresses, bytes[] memory newCCVArgs) =
      s_onRampTestHelper.mergeCCVLists(userCCVAddresses, userCCVArgs, new address[](0), poolRequiredCCV);

    // Should result in only one instance of the CCV with user args preserved.
    address[] memory expectedAddresses = new address[](1);
    expectedAddresses[0] = requiredCCV1;
    bytes[] memory expectedArgs = new bytes[](1);
    expectedArgs[0] = userCCVArgs[0];
    _assertCCVArraysEqual(newCCVAddresses, newCCVArgs, expectedAddresses, expectedArgs);
  }

  function test_mergeCCVLists_PoolIncludesDefaults_DedupsAgainstUser() public {
    address[] memory defaultCCVs = new address[](2);
    defaultCCVs[0] = makeAddr("defaultCCV1");
    defaultCCVs[1] = makeAddr("defaultCCV2");

    address[] memory poolRequiredCCV = new address[](4);
    poolRequiredCCV[0] = makeAddr("poolCCV1");
    poolRequiredCCV[1] = makeAddr("poolCCV2");
    poolRequiredCCV[2] = defaultCCVs[0];
    poolRequiredCCV[3] = defaultCCVs[1];

    address[] memory userCCVAddresses = new address[](1);
    userCCVAddresses[0] = defaultCCVs[0];
    bytes[] memory userCCVArgs = new bytes[](1);
    userCCVArgs[0] = "userArgs";

    (address[] memory newCCVAddresses, bytes[] memory newCCVArgs) =
      s_onRampTestHelper.mergeCCVLists(userCCVAddresses, userCCVArgs, new address[](0), poolRequiredCCV);

    address[] memory expectedAddresses = new address[](4);
    expectedAddresses[0] = defaultCCVs[0];
    expectedAddresses[1] = poolRequiredCCV[0];
    expectedAddresses[2] = poolRequiredCCV[1];
    expectedAddresses[3] = defaultCCVs[1];
    bytes[] memory expectedArgs = new bytes[](4);
    expectedArgs[0] = "userArgs";
    expectedArgs[1] = "";
    expectedArgs[2] = "";
    expectedArgs[3] = "";
    _assertCCVArraysEqual(newCCVAddresses, newCCVArgs, expectedAddresses, expectedArgs);
  }
}
