// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../../libraries/Client.sol";
import {CCVProxy} from "../../../onRamp/CCVProxy.sol";
import {CCVProxyTestHelper} from "../../helpers/CCVProxyTestHelper.sol";
import {CCVProxySetup} from "./CCVProxySetup.t.sol";

contract CCVProxy_mergeCCVsWithPoolAndLaneMandated is CCVProxySetup {
  CCVProxyTestHelper internal s_ccvProxyTestHelper;
  uint64 constant TEST_DEST_CHAIN_SELECTOR = 999999;
  address internal POOL_CCV1;
  address internal POOL_CCV2;
  address internal REQUIRED_CCV1;
  address internal REQUIRED_CCV2;
  address internal OPTIONAL_CCV1;
  address internal OPTIONAL_CCV2;

  function _setupTestDestChainConfig(
    address[] memory laneMandatedCCVs
  ) internal {
    CCVProxy.DestChainConfigArgs[] memory destChainConfigArgs = new CCVProxy.DestChainConfigArgs[](1);
    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = makeAddr("defaultCCV"); // Some default CCV to maintain invariant

    destChainConfigArgs[0] = CCVProxy.DestChainConfigArgs({
      destChainSelector: TEST_DEST_CHAIN_SELECTOR,
      router: s_sourceRouter,
      laneMandatedCCVs: laneMandatedCCVs,
      defaultCCVs: defaultCCVs,
      defaultExecutor: makeAddr("defaultExecutor")
    });

    s_ccvProxyTestHelper.applyDestChainConfigUpdates(destChainConfigArgs);
  }

  function setUp() public override {
    super.setUp();

    // Initialize test addresses
    POOL_CCV1 = makeAddr("POOL_CCV1");
    POOL_CCV2 = makeAddr("POOL_CCV2");
    REQUIRED_CCV1 = makeAddr("REQUIRED_CCV1");
    REQUIRED_CCV2 = makeAddr("REQUIRED_CCV2");
    OPTIONAL_CCV1 = makeAddr("OPTIONAL_CCV1");
    OPTIONAL_CCV2 = makeAddr("OPTIONAL_CCV2");

    s_ccvProxyTestHelper = new CCVProxyTestHelper(
      CCVProxy.StaticConfig({
        chainSelector: SOURCE_CHAIN_SELECTOR,
        rmnRemote: s_mockRMNRemote,
        tokenAdminRegistry: address(s_tokenAdminRegistry)
      }),
      CCVProxy.DynamicConfig({
        feeQuoter: address(s_feeQuoter),
        reentrancyGuardEntered: false,
        feeAggregator: FEE_AGGREGATOR
      })
    );

    _setupTestDestChainConfig(new address[](0));
  }

  function test_mergeCCVsWithPoolAndLaneMandated_MovesOptionalToRequiredAndDecrementsThreshold() public view {
    // Setup pool CCV that exists in optional CCVs.
    address[] memory poolRequiredCCV = new address[](1);
    poolRequiredCCV[0] = OPTIONAL_CCV1; // This CCV is in optional list

    Client.CCV[] memory requiredCCV = new Client.CCV[](1);
    requiredCCV[0] = Client.CCV({ccvAddress: REQUIRED_CCV1, args: "required1"});

    Client.CCV[] memory optionalCCV = new Client.CCV[](2);
    optionalCCV[0] = Client.CCV({ccvAddress: OPTIONAL_CCV1, args: "optional1"});
    optionalCCV[1] = Client.CCV({ccvAddress: OPTIONAL_CCV2, args: "optional2"});

    uint8 optionalThreshold = 2;

    (Client.CCV[] memory newRequiredCCVs, Client.CCV[] memory newOptionalCCVs, uint8 newOptionalThreshold) =
    s_ccvProxyTestHelper.mergeCCVsWithPoolAndLaneMandated(
      TEST_DEST_CHAIN_SELECTOR, poolRequiredCCV, requiredCCV, optionalCCV, optionalThreshold
    );

    // Should have moved OPTIONAL_CCV1 from optional to required.
    Client.CCV[] memory expectedRequired = new Client.CCV[](2);
    expectedRequired[0] = Client.CCV({ccvAddress: OPTIONAL_CCV1, args: optionalCCV[0].args});
    expectedRequired[1] = Client.CCV({ccvAddress: REQUIRED_CCV1, args: requiredCCV[0].args});
    _assertCCVArraysEqual(newRequiredCCVs, expectedRequired);

    // Optional should have one less CCV.
    Client.CCV[] memory expectedOptional = new Client.CCV[](1);
    expectedOptional[0] = Client.CCV({ccvAddress: OPTIONAL_CCV2, args: optionalCCV[1].args});
    _assertCCVArraysEqual(newOptionalCCVs, expectedOptional);

    // Threshold should be decremented to maintain minimum verification count.
    assertEq(newOptionalThreshold, 1);
  }

  function test_mergeCCVsWithPoolAndLaneMandated_SkipsDuplicatesInPoolRequiredCCV() public view {
    // Setup pool CCVs with duplicates.
    address[] memory poolRequiredCCV = new address[](3);
    poolRequiredCCV[0] = POOL_CCV1;
    poolRequiredCCV[1] = POOL_CCV1; // Duplicate
    poolRequiredCCV[2] = POOL_CCV2;

    Client.CCV[] memory requiredCCV = new Client.CCV[](0);
    Client.CCV[] memory optionalCCV = new Client.CCV[](0);
    uint8 optionalThreshold = 0;

    (Client.CCV[] memory newRequiredCCVs, Client.CCV[] memory newOptionalCCVs, uint8 newOptionalThreshold) =
    s_ccvProxyTestHelper.mergeCCVsWithPoolAndLaneMandated(
      TEST_DEST_CHAIN_SELECTOR, poolRequiredCCV, requiredCCV, optionalCCV, optionalThreshold
    );

    // Should only add unique pool CCVs.
    Client.CCV[] memory expectedRequired = new Client.CCV[](2);
    expectedRequired[0] = Client.CCV({ccvAddress: POOL_CCV1, args: ""});
    expectedRequired[1] = Client.CCV({ccvAddress: POOL_CCV2, args: ""});
    _assertCCVArraysEqual(newRequiredCCVs, expectedRequired);
    _assertCCVArraysEqual(newOptionalCCVs, optionalCCV);
    assertEq(newOptionalThreshold, optionalThreshold);
  }

  function test_mergeCCVsWithPoolAndLaneMandated_MovesAllOptionalToRequired() public view {
    // Setup scenario where all optional CCVs are moved to required.
    address[] memory poolRequiredCCV = new address[](2);
    poolRequiredCCV[0] = OPTIONAL_CCV1;
    poolRequiredCCV[1] = OPTIONAL_CCV2;

    Client.CCV[] memory requiredCCV = new Client.CCV[](0);
    Client.CCV[] memory optionalCCV = new Client.CCV[](2);
    optionalCCV[0] = Client.CCV({ccvAddress: OPTIONAL_CCV1, args: "optional1"});
    optionalCCV[1] = Client.CCV({ccvAddress: OPTIONAL_CCV2, args: "optional2"});

    uint8 optionalThreshold = 2; // This will become 0 after moving both optionals to required.

    (Client.CCV[] memory newRequiredCCVs, Client.CCV[] memory newOptionalCCVs, uint8 newOptionalThreshold) =
    s_ccvProxyTestHelper.mergeCCVsWithPoolAndLaneMandated(
      TEST_DEST_CHAIN_SELECTOR, poolRequiredCCV, requiredCCV, optionalCCV, optionalThreshold
    );

    // All optionals should be moved to required.
    Client.CCV[] memory expectedRequired = new Client.CCV[](2);
    expectedRequired[0] = Client.CCV({ccvAddress: OPTIONAL_CCV1, args: optionalCCV[0].args});
    expectedRequired[1] = Client.CCV({ccvAddress: OPTIONAL_CCV2, args: optionalCCV[1].args});
    _assertCCVArraysEqual(newRequiredCCVs, expectedRequired);

    // Optional array should be empty.
    assertEq(newOptionalCCVs.length, 0);

    // Threshold should be 0 (decremented for each moved CCV).
    assertEq(newOptionalThreshold, 0);
  }

  function test_mergeCCVsWithPoolAndLaneMandated_NoChangesWhenPoolCCVAlreadyInRequired() public view {
    address[] memory poolRequiredCCV = new address[](1);
    poolRequiredCCV[0] = REQUIRED_CCV1; // Already in required

    Client.CCV[] memory requiredCCV = new Client.CCV[](1);
    requiredCCV[0] = Client.CCV({ccvAddress: REQUIRED_CCV1, args: "required1"});

    Client.CCV[] memory optionalCCV = new Client.CCV[](1);
    optionalCCV[0] = Client.CCV({ccvAddress: OPTIONAL_CCV1, args: "optional1"});

    uint8 optionalThreshold = 1;

    (Client.CCV[] memory newRequiredCCVs, Client.CCV[] memory newOptionalCCVs, uint8 newOptionalThreshold) =
    s_ccvProxyTestHelper.mergeCCVsWithPoolAndLaneMandated(
      TEST_DEST_CHAIN_SELECTOR, poolRequiredCCV, requiredCCV, optionalCCV, optionalThreshold
    );

    // Should return original arrays unchanged.
    _assertCCVArraysEqual(newRequiredCCVs, requiredCCV);
    _assertCCVArraysEqual(newOptionalCCVs, optionalCCV);

    assertEq(newOptionalThreshold, optionalThreshold);
  }

  function test_mergeCCVsWithPoolAndLaneMandated_EmptyPoolRequiredCCV() public view {
    address[] memory poolRequiredCCV = new address[](0);

    Client.CCV[] memory requiredCCV = new Client.CCV[](1);
    requiredCCV[0] = Client.CCV({ccvAddress: REQUIRED_CCV1, args: "required1"});

    Client.CCV[] memory optionalCCV = new Client.CCV[](1);
    optionalCCV[0] = Client.CCV({ccvAddress: OPTIONAL_CCV1, args: "optional1"});

    uint8 optionalThreshold = 1;

    (Client.CCV[] memory newRequiredCCVs, Client.CCV[] memory newOptionalCCVs, uint8 newOptionalThreshold) =
    s_ccvProxyTestHelper.mergeCCVsWithPoolAndLaneMandated(
      TEST_DEST_CHAIN_SELECTOR, poolRequiredCCV, requiredCCV, optionalCCV, optionalThreshold
    );

    // Should return original arrays unchanged.
    _assertCCVArraysEqual(newRequiredCCVs, requiredCCV);
    _assertCCVArraysEqual(newOptionalCCVs, optionalCCV);

    assertEq(newOptionalThreshold, optionalThreshold);
  }

  function test_mergeCCVsWithPoolAndLaneMandated_WithLaneMandatedCCVs() public {
    // Setup lane mandated CCVs
    address[] memory laneMandatedCCVs = new address[](2);
    laneMandatedCCVs[0] = makeAddr("laneMandatedCCV1");
    laneMandatedCCVs[1] = OPTIONAL_CCV1; // This one is in optional list

    _setupTestDestChainConfig(laneMandatedCCVs);

    address[] memory poolRequiredCCV = new address[](0);

    Client.CCV[] memory requiredCCV = new Client.CCV[](1);
    requiredCCV[0] = Client.CCV({ccvAddress: REQUIRED_CCV1, args: "required1"});

    Client.CCV[] memory optionalCCV = new Client.CCV[](2);
    optionalCCV[0] = Client.CCV({ccvAddress: OPTIONAL_CCV1, args: "optional1"});
    optionalCCV[1] = Client.CCV({ccvAddress: OPTIONAL_CCV2, args: "optional2"});

    uint8 optionalThreshold = 2;

    (Client.CCV[] memory newRequiredCCVs, Client.CCV[] memory newOptionalCCVs, uint8 newOptionalThreshold) =
    s_ccvProxyTestHelper.mergeCCVsWithPoolAndLaneMandated(
      TEST_DEST_CHAIN_SELECTOR, poolRequiredCCV, requiredCCV, optionalCCV, optionalThreshold
    );

    // Should add lane mandated CCVs to required and move OPTIONAL_CCV1 from optional
    Client.CCV[] memory expectedRequired = new Client.CCV[](3);
    expectedRequired[0] = Client.CCV({ccvAddress: laneMandatedCCVs[0], args: ""});
    expectedRequired[1] = Client.CCV({ccvAddress: OPTIONAL_CCV1, args: "optional1"});
    expectedRequired[2] = Client.CCV({ccvAddress: REQUIRED_CCV1, args: requiredCCV[0].args});
    _assertCCVArraysEqual(newRequiredCCVs, expectedRequired);

    // Optional should only have OPTIONAL_CCV2 left
    Client.CCV[] memory expectedOptional = new Client.CCV[](1);
    expectedOptional[0] = Client.CCV({ccvAddress: OPTIONAL_CCV2, args: optionalCCV[1].args});
    _assertCCVArraysEqual(newOptionalCCVs, expectedOptional);

    // Threshold should be decremented
    assertEq(newOptionalThreshold, 1);
  }

  function test_mergeCCVsWithPoolAndLaneMandated_BothLaneAndPoolCCVs() public {
    // Setup both lane mandated and pool required CCVs
    address[] memory laneMandatedCCVs = new address[](1);
    laneMandatedCCVs[0] = makeAddr("laneMandatedCCV1");

    _setupTestDestChainConfig(laneMandatedCCVs);

    address[] memory poolRequiredCCV = new address[](1);
    poolRequiredCCV[0] = OPTIONAL_CCV1; // This one is in optional list

    Client.CCV[] memory requiredCCV = new Client.CCV[](1);
    requiredCCV[0] = Client.CCV({ccvAddress: REQUIRED_CCV1, args: "required1"});

    Client.CCV[] memory optionalCCV = new Client.CCV[](2);
    optionalCCV[0] = Client.CCV({ccvAddress: OPTIONAL_CCV1, args: "optional1"});
    optionalCCV[1] = Client.CCV({ccvAddress: OPTIONAL_CCV2, args: "optional2"});

    uint8 optionalThreshold = 2;

    (Client.CCV[] memory newRequiredCCVs, Client.CCV[] memory newOptionalCCVs, uint8 newOptionalThreshold) =
    s_ccvProxyTestHelper.mergeCCVsWithPoolAndLaneMandated(
      TEST_DEST_CHAIN_SELECTOR, poolRequiredCCV, requiredCCV, optionalCCV, optionalThreshold
    );

    // Should add both lane mandated and pool required CCVs
    Client.CCV[] memory expectedRequired = new Client.CCV[](3);
    expectedRequired[0] = Client.CCV({ccvAddress: laneMandatedCCVs[0], args: ""});
    expectedRequired[1] = Client.CCV({ccvAddress: OPTIONAL_CCV1, args: "optional1"});
    expectedRequired[2] = Client.CCV({ccvAddress: REQUIRED_CCV1, args: requiredCCV[0].args});
    _assertCCVArraysEqual(newRequiredCCVs, expectedRequired);

    // Optional should only have OPTIONAL_CCV2 left
    Client.CCV[] memory expectedOptional = new Client.CCV[](1);
    expectedOptional[0] = Client.CCV({ccvAddress: OPTIONAL_CCV2, args: optionalCCV[1].args});
    _assertCCVArraysEqual(newOptionalCCVs, expectedOptional);

    // Threshold should be decremented
    assertEq(newOptionalThreshold, 1);
  }
}
