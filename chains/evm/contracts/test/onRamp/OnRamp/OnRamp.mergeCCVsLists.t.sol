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
    // Setup pool CCVs with duplicates.
    address poolCCV1 = makeAddr("poolCCV1");
    address poolCCV2 = makeAddr("poolCCV2");

    address[] memory poolRequiredCCV = new address[](3);
    poolRequiredCCV[0] = poolCCV1;
    poolRequiredCCV[1] = poolCCV1; // Duplicate
    poolRequiredCCV[2] = poolCCV2;

    Client.CCV[] memory userSpecifiedCCVs = new Client.CCV[](0);
    address[] memory defaultCCVs = new address[](0);

    Client.CCV[] memory newRequiredCCVs =
      s_onRampTestHelper.mergeCCVLists(userSpecifiedCCVs, new address[](0), poolRequiredCCV, defaultCCVs);

    // Should only add unique pool CCVs.
    Client.CCV[] memory expectedRequired = new Client.CCV[](2);
    expectedRequired[0] = Client.CCV({ccvAddress: poolCCV1, args: ""});
    expectedRequired[1] = Client.CCV({ccvAddress: poolCCV2, args: ""});
    _assertCCVArraysEqual(newRequiredCCVs, expectedRequired);
  }

  function test_mergeCCVLists_NoChangesWhenPoolCCVAlreadyInRequired() public {
    address requiredCCV1 = makeAddr("requiredCCV1");

    address[] memory poolRequiredCCV = new address[](1);
    poolRequiredCCV[0] = requiredCCV1; // Already in required

    Client.CCV[] memory userSpecifiedCCVs = new Client.CCV[](1);
    userSpecifiedCCVs[0] = Client.CCV({ccvAddress: requiredCCV1, args: "required1"});
    address[] memory defaultCCVs = new address[](0);

    Client.CCV[] memory newRequiredCCVs =
      s_onRampTestHelper.mergeCCVLists(userSpecifiedCCVs, new address[](0), poolRequiredCCV, defaultCCVs);

    // Should return original arrays unchanged.
    _assertCCVArraysEqual(newRequiredCCVs, userSpecifiedCCVs);
  }

  function test_mergeCCVLists_PoolFallbackDefaults_UsesDefaults() public {
    address[] memory defaultCCVs = new address[](2);
    defaultCCVs[0] = makeAddr("defaultCCV1");
    defaultCCVs[1] = makeAddr("defaultCCV2");

    Client.CCV[] memory userSpecifiedCCVs = new Client.CCV[](1);
    userSpecifiedCCVs[0] = Client.CCV({ccvAddress: makeAddr("requiredCCV1"), args: "required1"});

    // Simulate pool fallback returning the destination defaults directly.
    address[] memory poolRequiredCCV = new address[](defaultCCVs.length);
    for (uint256 i = 0; i < defaultCCVs.length; ++i) {
      poolRequiredCCV[i] = defaultCCVs[i];
    }

    Client.CCV[] memory newRequiredCCVs =
      s_onRampTestHelper.mergeCCVLists(userSpecifiedCCVs, new address[](0), poolRequiredCCV, defaultCCVs);

    Client.CCV[] memory expectedRequired = new Client.CCV[](3);
    expectedRequired[0] = userSpecifiedCCVs[0];
    expectedRequired[1] = Client.CCV({ccvAddress: defaultCCVs[0], args: ""});
    expectedRequired[2] = Client.CCV({ccvAddress: defaultCCVs[1], args: ""});
    _assertCCVArraysEqual(newRequiredCCVs, expectedRequired);
  }

  function test_mergeCCVLists_NoPoolProcessing_KeepsUserAndLaneOnly() public {
    Client.CCV[] memory userSpecifiedCCVs = new Client.CCV[](1);
    userSpecifiedCCVs[0] = Client.CCV({ccvAddress: makeAddr("userCCV"), args: "userArgs"});

    address[] memory laneMandatedCCVs = new address[](1);
    laneMandatedCCVs[0] = makeAddr("laneCCV");

    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = makeAddr("defaultCCV");

    Client.CCV[] memory newRequiredCCVs =
      s_onRampTestHelper.mergeCCVLists(userSpecifiedCCVs, laneMandatedCCVs, new address[](0), defaultCCVs);

    Client.CCV[] memory expectedRequired = new Client.CCV[](2);
    expectedRequired[0] = userSpecifiedCCVs[0];
    expectedRequired[1] = Client.CCV({ccvAddress: laneMandatedCCVs[0], args: ""});
    _assertCCVArraysEqual(newRequiredCCVs, expectedRequired);
  }

  function test_mergeCCVLists_DedupUserAndMandatoryCCVs() public {
    address requiredCCV1 = makeAddr("requiredCCV1");

    address[] memory laneMandatedCCVs = new address[](2);
    laneMandatedCCVs[0] = makeAddr("laneMandatedCCV1");
    laneMandatedCCVs[1] = requiredCCV1; // This one is also in user specified list

    address[] memory poolRequiredCCV = new address[](1);
    poolRequiredCCV[0] = makeAddr("poolCCV1");
    address[] memory defaultCCVs = new address[](0);

    Client.CCV[] memory userSpecifiedCCVs = new Client.CCV[](1);
    userSpecifiedCCVs[0] = Client.CCV({ccvAddress: requiredCCV1, args: "required1"});

    Client.CCV[] memory newRequiredCCVs =
      s_onRampTestHelper.mergeCCVLists(userSpecifiedCCVs, laneMandatedCCVs, poolRequiredCCV, defaultCCVs);

    Client.CCV[] memory expectedRequired = new Client.CCV[](3);
    expectedRequired[0] = Client.CCV({ccvAddress: requiredCCV1, args: userSpecifiedCCVs[0].args});
    expectedRequired[1] = Client.CCV({ccvAddress: laneMandatedCCVs[0], args: ""});
    expectedRequired[2] = Client.CCV({ccvAddress: poolRequiredCCV[0], args: ""});
    _assertCCVArraysEqual(newRequiredCCVs, expectedRequired);
  }

  function test_mergeCCVLists_DedupUserAndPoolCCVs() public {
    // Setup both lane mandated and pool required CCVs
    address requiredCCV1 = makeAddr("requiredCCV1");

    address[] memory poolRequiredCCV = new address[](1);
    poolRequiredCCV[0] = requiredCCV1; // This one is also in user specified list
    address[] memory defaultCCVs = new address[](0);

    Client.CCV[] memory userSpecifiedCCVs = new Client.CCV[](1);
    userSpecifiedCCVs[0] = Client.CCV({ccvAddress: requiredCCV1, args: "required1"});

    Client.CCV[] memory newRequiredCCVs =
      s_onRampTestHelper.mergeCCVLists(userSpecifiedCCVs, new address[](0), poolRequiredCCV, defaultCCVs);

    // Should result in only one instance of the CCV with user args preserved.
    Client.CCV[] memory expectedRequired = new Client.CCV[](1);
    expectedRequired[0] = Client.CCV({ccvAddress: requiredCCV1, args: userSpecifiedCCVs[0].args});
    _assertCCVArraysEqual(newRequiredCCVs, expectedRequired);
  }

  function test_mergeCCVLists_PoolReturnsAddressZero_IncludesDefaults() public {
    address[] memory defaultCCVs = new address[](2);
    defaultCCVs[0] = makeAddr("defaultCCV1");
    defaultCCVs[1] = makeAddr("defaultCCV2");

    address[] memory poolRequiredCCV = new address[](3);
    poolRequiredCCV[0] = makeAddr("poolCCV1");
    poolRequiredCCV[1] = address(0); // Signal to include defaults
    poolRequiredCCV[2] = makeAddr("poolCCV2");

    Client.CCV[] memory userSpecifiedCCVs = new Client.CCV[](0);

    Client.CCV[] memory newRequiredCCVs =
      s_onRampTestHelper.mergeCCVLists(userSpecifiedCCVs, new address[](0), poolRequiredCCV, defaultCCVs);

    // Pool signaled with address(0), so defaults should be included.
    Client.CCV[] memory expectedRequired = new Client.CCV[](4);
    expectedRequired[0] = Client.CCV({ccvAddress: poolRequiredCCV[0], args: ""});
    expectedRequired[1] = Client.CCV({ccvAddress: defaultCCVs[0], args: ""});
    expectedRequired[2] = Client.CCV({ccvAddress: defaultCCVs[1], args: ""});
    expectedRequired[3] = Client.CCV({ccvAddress: poolRequiredCCV[2], args: ""});
    _assertCCVArraysEqual(newRequiredCCVs, expectedRequired);
  }

  function test_mergeCCVLists_MultipleAddressZero_OnlyIncludesDefaultsOnce() public {
    address[] memory defaultCCVs = new address[](2);
    defaultCCVs[0] = makeAddr("defaultCCV1");
    defaultCCVs[1] = makeAddr("defaultCCV2");

    address[] memory poolRequiredCCV = new address[](4);
    poolRequiredCCV[0] = makeAddr("poolCCV1");
    poolRequiredCCV[1] = address(0); // First signal
    poolRequiredCCV[2] = address(0); // Duplicate signal - should be ignored
    poolRequiredCCV[3] = makeAddr("poolCCV2");

    Client.CCV[] memory userSpecifiedCCVs = new Client.CCV[](0);

    Client.CCV[] memory newRequiredCCVs =
      s_onRampTestHelper.mergeCCVLists(userSpecifiedCCVs, new address[](0), poolRequiredCCV, defaultCCVs);

    // Multiple address(0) signals, but defaults should only be added once.
    Client.CCV[] memory expectedRequired = new Client.CCV[](4);
    expectedRequired[0] = Client.CCV({ccvAddress: poolRequiredCCV[0], args: ""});
    expectedRequired[1] = Client.CCV({ccvAddress: defaultCCVs[0], args: ""});
    expectedRequired[2] = Client.CCV({ccvAddress: defaultCCVs[1], args: ""});
    expectedRequired[3] = Client.CCV({ccvAddress: poolRequiredCCV[3], args: ""});
    _assertCCVArraysEqual(newRequiredCCVs, expectedRequired);
  }

  function test_mergeCCVLists_AddressZeroWithDefaultAlreadyInUser_DedupsCorrectly() public {
    address[] memory defaultCCVs = new address[](2);
    defaultCCVs[0] = makeAddr("defaultCCV1");
    defaultCCVs[1] = makeAddr("defaultCCV2");

    address[] memory poolRequiredCCV = new address[](2);
    poolRequiredCCV[0] = address(0); // Signal to include defaults
    poolRequiredCCV[1] = makeAddr("poolCCV1");

    Client.CCV[] memory userSpecifiedCCVs = new Client.CCV[](1);
    userSpecifiedCCVs[0] = Client.CCV({ccvAddress: defaultCCVs[0], args: "userArgs"}); // User already has one of the defaults.

    Client.CCV[] memory newRequiredCCVs =
      s_onRampTestHelper.mergeCCVLists(userSpecifiedCCVs, new address[](0), poolRequiredCCV, defaultCCVs);

    // User already has defaultCCV1, so it shouldn't be added again. User's args should be preserved.
    Client.CCV[] memory expectedRequired = new Client.CCV[](3);
    expectedRequired[0] = Client.CCV({ccvAddress: defaultCCVs[0], args: "userArgs"}); // User's version with args.
    expectedRequired[1] = Client.CCV({ccvAddress: defaultCCVs[1], args: ""}); // Second default.
    expectedRequired[2] = Client.CCV({ccvAddress: poolRequiredCCV[1], args: ""}); // Pool CCV.
    _assertCCVArraysEqual(newRequiredCCVs, expectedRequired);
  }
}
