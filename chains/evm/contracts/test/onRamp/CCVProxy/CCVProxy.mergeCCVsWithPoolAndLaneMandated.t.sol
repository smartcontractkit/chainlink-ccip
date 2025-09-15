// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../../libraries/Client.sol";
import {CCVProxy} from "../../../onRamp/CCVProxy.sol";
import {CCVProxyTestHelper} from "../../helpers/CCVProxyTestHelper.sol";
import {CCVProxySetup} from "./CCVProxySetup.t.sol";

contract CCVProxy_mergeCCVsWithPoolAndLaneMandated is CCVProxySetup {
  CCVProxyTestHelper internal s_ccvProxyTestHelper;

  function _setupTestDestChainConfig(
    address[] memory laneMandatedCCVs
  ) internal {
    CCVProxy.DestChainConfigArgs[] memory destChainConfigArgs = new CCVProxy.DestChainConfigArgs[](1);
    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = makeAddr("defaultCCV");

    destChainConfigArgs[0] = CCVProxy.DestChainConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      router: s_sourceRouter,
      laneMandatedCCVs: laneMandatedCCVs,
      defaultCCVs: defaultCCVs,
      defaultExecutor: makeAddr("defaultExecutor"),
      ccvAggregator: abi.encodePacked(address(s_ccvAggregatorRemote))
    });

    s_ccvProxyTestHelper.applyDestChainConfigUpdates(destChainConfigArgs);
  }

  function setUp() public override {
    super.setUp();
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

  function test_mergeCCVsWithPoolAndLaneMandated_MovesOptionalToRequiredAndDecrementsThreshold() public {
    // Setup pool CCV that exists in optional CCVs.
    address optionalCCV1 = makeAddr("optionalCCV1");
    address optionalCCV2 = makeAddr("optionalCCV2");
    address requiredCCV1 = makeAddr("requiredCCV1");

    address[] memory poolRequiredCCV = new address[](1);
    poolRequiredCCV[0] = optionalCCV1; // This CCV is in optional list

    Client.CCV[] memory requiredCCV = new Client.CCV[](1);
    requiredCCV[0] = Client.CCV({ccvAddress: requiredCCV1, args: "required1"});

    Client.CCV[] memory optionalCCV = new Client.CCV[](2);
    optionalCCV[0] = Client.CCV({ccvAddress: optionalCCV1, args: "optional1"});
    optionalCCV[1] = Client.CCV({ccvAddress: optionalCCV2, args: "optional2"});

    uint8 optionalThreshold = 2;

    (Client.CCV[] memory newRequiredCCVs, Client.CCV[] memory newOptionalCCVs, uint8 newOptionalThreshold) =
    s_ccvProxyTestHelper.mergeCCVsWithPoolAndLaneMandated(
      DEST_CHAIN_SELECTOR, poolRequiredCCV, requiredCCV, optionalCCV, optionalThreshold
    );

    // Should have moved optionalCCV1 from optional to required.
    Client.CCV[] memory expectedRequired = new Client.CCV[](2);
    expectedRequired[0] = Client.CCV({ccvAddress: optionalCCV1, args: optionalCCV[0].args});
    expectedRequired[1] = Client.CCV({ccvAddress: requiredCCV1, args: requiredCCV[0].args});
    _assertCCVArraysEqual(newRequiredCCVs, expectedRequired);

    // Optional should have one less CCV.
    Client.CCV[] memory expectedOptional = new Client.CCV[](1);
    expectedOptional[0] = Client.CCV({ccvAddress: optionalCCV2, args: optionalCCV[1].args});
    _assertCCVArraysEqual(newOptionalCCVs, expectedOptional);

    // Threshold should be decremented to maintain minimum verification count.
    assertEq(newOptionalThreshold, 1);
  }

  function test_mergeCCVsWithPoolAndLaneMandated_SkipsDuplicatesInPoolRequiredCCV() public {
    // Setup pool CCVs with duplicates.
    address poolCCV1 = makeAddr("poolCCV1");
    address poolCCV2 = makeAddr("poolCCV2");

    address[] memory poolRequiredCCV = new address[](3);
    poolRequiredCCV[0] = poolCCV1;
    poolRequiredCCV[1] = poolCCV1; // Duplicate
    poolRequiredCCV[2] = poolCCV2;

    Client.CCV[] memory requiredCCV = new Client.CCV[](0);
    Client.CCV[] memory optionalCCV = new Client.CCV[](0);
    uint8 optionalThreshold = 0;

    (Client.CCV[] memory newRequiredCCVs, Client.CCV[] memory newOptionalCCVs, uint8 newOptionalThreshold) =
    s_ccvProxyTestHelper.mergeCCVsWithPoolAndLaneMandated(
      DEST_CHAIN_SELECTOR, poolRequiredCCV, requiredCCV, optionalCCV, optionalThreshold
    );

    // Should only add unique pool CCVs.
    Client.CCV[] memory expectedRequired = new Client.CCV[](2);
    expectedRequired[0] = Client.CCV({ccvAddress: poolCCV1, args: ""});
    expectedRequired[1] = Client.CCV({ccvAddress: poolCCV2, args: ""});
    _assertCCVArraysEqual(newRequiredCCVs, expectedRequired);
    _assertCCVArraysEqual(newOptionalCCVs, optionalCCV);
    assertEq(newOptionalThreshold, optionalThreshold);
  }

  function test_mergeCCVsWithPoolAndLaneMandated_MovesAllOptionalToRequired() public {
    // Setup scenario where all optional CCVs are moved to required.
    address optionalCCV1 = makeAddr("optionalCCV1");
    address optionalCCV2 = makeAddr("optionalCCV2");

    address[] memory poolRequiredCCV = new address[](2);
    poolRequiredCCV[0] = optionalCCV1;
    poolRequiredCCV[1] = optionalCCV2;

    Client.CCV[] memory requiredCCV = new Client.CCV[](0);
    Client.CCV[] memory optionalCCV = new Client.CCV[](2);
    optionalCCV[0] = Client.CCV({ccvAddress: optionalCCV1, args: "optional1"});
    optionalCCV[1] = Client.CCV({ccvAddress: optionalCCV2, args: "optional2"});

    uint8 optionalThreshold = 2; // This will become 0 after moving both optionals to required.

    (Client.CCV[] memory newRequiredCCVs, Client.CCV[] memory newOptionalCCVs, uint8 newOptionalThreshold) =
    s_ccvProxyTestHelper.mergeCCVsWithPoolAndLaneMandated(
      DEST_CHAIN_SELECTOR, poolRequiredCCV, requiredCCV, optionalCCV, optionalThreshold
    );

    // All optionals should be moved to required.
    Client.CCV[] memory expectedRequired = new Client.CCV[](2);
    expectedRequired[0] = Client.CCV({ccvAddress: optionalCCV1, args: optionalCCV[0].args});
    expectedRequired[1] = Client.CCV({ccvAddress: optionalCCV2, args: optionalCCV[1].args});
    _assertCCVArraysEqual(newRequiredCCVs, expectedRequired);

    // Optional array should be empty.
    assertEq(newOptionalCCVs.length, 0);

    // Threshold should be 0 (decremented for each moved CCV).
    assertEq(newOptionalThreshold, 0);
  }

  function test_mergeCCVsWithPoolAndLaneMandated_NoChangesWhenPoolCCVAlreadyInRequired() public {
    address requiredCCV1 = makeAddr("requiredCCV1");
    address optionalCCV1 = makeAddr("optionalCCV1");

    address[] memory poolRequiredCCV = new address[](1);
    poolRequiredCCV[0] = requiredCCV1; // Already in required

    Client.CCV[] memory requiredCCV = new Client.CCV[](1);
    requiredCCV[0] = Client.CCV({ccvAddress: requiredCCV1, args: "required1"});

    Client.CCV[] memory optionalCCV = new Client.CCV[](1);
    optionalCCV[0] = Client.CCV({ccvAddress: optionalCCV1, args: "optional1"});

    uint8 optionalThreshold = 1;

    (Client.CCV[] memory newRequiredCCVs, Client.CCV[] memory newOptionalCCVs, uint8 newOptionalThreshold) =
    s_ccvProxyTestHelper.mergeCCVsWithPoolAndLaneMandated(
      DEST_CHAIN_SELECTOR, poolRequiredCCV, requiredCCV, optionalCCV, optionalThreshold
    );

    // Should return original arrays unchanged.
    _assertCCVArraysEqual(newRequiredCCVs, requiredCCV);
    _assertCCVArraysEqual(newOptionalCCVs, optionalCCV);

    assertEq(newOptionalThreshold, optionalThreshold);
  }

  function test_mergeCCVsWithPoolAndLaneMandated_EmptyPoolRequiredCCV() public {
    address[] memory poolRequiredCCV = new address[](0);

    Client.CCV[] memory requiredCCV = new Client.CCV[](1);
    requiredCCV[0] = Client.CCV({ccvAddress: makeAddr("requiredCCV1"), args: "required1"});

    Client.CCV[] memory optionalCCV = new Client.CCV[](1);
    optionalCCV[0] = Client.CCV({ccvAddress: makeAddr("optionalCCV1"), args: "optional1"});

    uint8 optionalThreshold = 1;

    (Client.CCV[] memory newRequiredCCVs, Client.CCV[] memory newOptionalCCVs, uint8 newOptionalThreshold) =
    s_ccvProxyTestHelper.mergeCCVsWithPoolAndLaneMandated(
      DEST_CHAIN_SELECTOR, poolRequiredCCV, requiredCCV, optionalCCV, optionalThreshold
    );

    // Should return original arrays unchanged.
    _assertCCVArraysEqual(newRequiredCCVs, requiredCCV);
    _assertCCVArraysEqual(newOptionalCCVs, optionalCCV);

    assertEq(newOptionalThreshold, optionalThreshold);
  }

  function test_mergeCCVsWithPoolAndLaneMandated_WithLaneMandatedCCVs() public {
    // Setup lane mandated CCVs
    address optionalCCV1 = makeAddr("optionalCCV1");
    address optionalCCV2 = makeAddr("optionalCCV2");
    address requiredCCV1 = makeAddr("requiredCCV1");

    address[] memory laneMandatedCCVs = new address[](2);
    laneMandatedCCVs[0] = makeAddr("laneMandatedCCV1");
    laneMandatedCCVs[1] = optionalCCV1; // This one is in optional list

    _setupTestDestChainConfig(laneMandatedCCVs);

    address[] memory poolRequiredCCV = new address[](0);

    Client.CCV[] memory requiredCCV = new Client.CCV[](1);
    requiredCCV[0] = Client.CCV({ccvAddress: requiredCCV1, args: "required1"});

    Client.CCV[] memory optionalCCV = new Client.CCV[](2);
    optionalCCV[0] = Client.CCV({ccvAddress: optionalCCV1, args: "optional1"});
    optionalCCV[1] = Client.CCV({ccvAddress: optionalCCV2, args: "optional2"});

    uint8 optionalThreshold = 2;

    (Client.CCV[] memory newRequiredCCVs, Client.CCV[] memory newOptionalCCVs, uint8 newOptionalThreshold) =
    s_ccvProxyTestHelper.mergeCCVsWithPoolAndLaneMandated(
      DEST_CHAIN_SELECTOR, poolRequiredCCV, requiredCCV, optionalCCV, optionalThreshold
    );

    // Should add lane mandated CCVs to required and move optionalCCV1 from optional
    Client.CCV[] memory expectedRequired = new Client.CCV[](3);
    expectedRequired[0] = Client.CCV({ccvAddress: laneMandatedCCVs[0], args: ""});
    expectedRequired[1] = Client.CCV({ccvAddress: optionalCCV1, args: "optional1"});
    expectedRequired[2] = Client.CCV({ccvAddress: requiredCCV1, args: requiredCCV[0].args});
    _assertCCVArraysEqual(newRequiredCCVs, expectedRequired);

    // Optional should only have optionalCCV2 left
    Client.CCV[] memory expectedOptional = new Client.CCV[](1);
    expectedOptional[0] = Client.CCV({ccvAddress: optionalCCV2, args: optionalCCV[1].args});
    _assertCCVArraysEqual(newOptionalCCVs, expectedOptional);

    // Threshold should be decremented
    assertEq(newOptionalThreshold, 1);
  }

  function test_mergeCCVsWithPoolAndLaneMandated_BothLaneAndPoolCCVs() public {
    // Setup both lane mandated and pool required CCVs
    address optionalCCV1 = makeAddr("optionalCCV1");
    address optionalCCV2 = makeAddr("optionalCCV2");
    address requiredCCV1 = makeAddr("requiredCCV1");

    address[] memory laneMandatedCCVs = new address[](1);
    laneMandatedCCVs[0] = makeAddr("laneMandatedCCV1");

    _setupTestDestChainConfig(laneMandatedCCVs);

    address[] memory poolRequiredCCV = new address[](1);
    poolRequiredCCV[0] = optionalCCV1; // This one is in optional list

    Client.CCV[] memory requiredCCV = new Client.CCV[](1);
    requiredCCV[0] = Client.CCV({ccvAddress: requiredCCV1, args: "required1"});

    Client.CCV[] memory optionalCCV = new Client.CCV[](2);
    optionalCCV[0] = Client.CCV({ccvAddress: optionalCCV1, args: "optional1"});
    optionalCCV[1] = Client.CCV({ccvAddress: optionalCCV2, args: "optional2"});

    uint8 optionalThreshold = 2;

    (Client.CCV[] memory newRequiredCCVs, Client.CCV[] memory newOptionalCCVs, uint8 newOptionalThreshold) =
    s_ccvProxyTestHelper.mergeCCVsWithPoolAndLaneMandated(
      DEST_CHAIN_SELECTOR, poolRequiredCCV, requiredCCV, optionalCCV, optionalThreshold
    );

    // Should add both lane mandated and pool required CCVs
    Client.CCV[] memory expectedRequired = new Client.CCV[](3);
    expectedRequired[0] = Client.CCV({ccvAddress: laneMandatedCCVs[0], args: ""});
    expectedRequired[1] = Client.CCV({ccvAddress: optionalCCV1, args: "optional1"});
    expectedRequired[2] = Client.CCV({ccvAddress: requiredCCV1, args: requiredCCV[0].args});
    _assertCCVArraysEqual(newRequiredCCVs, expectedRequired);

    // Optional should only have optionalCCV2 left
    Client.CCV[] memory expectedOptional = new Client.CCV[](1);
    expectedOptional[0] = Client.CCV({ccvAddress: optionalCCV2, args: optionalCCV[1].args});
    _assertCCVArraysEqual(newOptionalCCVs, expectedOptional);

    // Threshold should be decremented
    assertEq(newOptionalThreshold, 1);
  }
}
