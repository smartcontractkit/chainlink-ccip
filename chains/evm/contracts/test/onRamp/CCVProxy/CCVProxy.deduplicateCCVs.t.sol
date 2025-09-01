// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Client} from "../../../libraries/Client.sol";
import {CCVProxy} from "../../../onRamp/CCVProxy.sol";
import {CCVProxyTestHelper} from "../../helpers/CCVProxyTestHelper.sol";
import {CCVProxySetup} from "./CCVProxySetup.t.sol";

contract CCVProxy_deduplicateCCVs is CCVProxySetup {
  CCVProxyTestHelper internal s_ccvProxyTestHelper;
  address internal constant POOL_CCV1 = address(0x1111);
  address internal constant POOL_CCV2 = address(0x2222);
  address internal constant REQUIRED_CCV1 = address(0x3333);
  address internal constant REQUIRED_CCV2 = address(0x4444);
  address internal constant OPTIONAL_CCV1 = address(0x5555);
  address internal constant OPTIONAL_CCV2 = address(0x6666);

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
  }

  function test_deduplicateCCVs_MovesOptionalToRequiredAndDecrementsThreshold() public view {
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
      s_ccvProxyTestHelper.deduplicateCCVs(poolRequiredCCV, requiredCCV, optionalCCV, optionalThreshold);

    // Should have moved OPTIONAL_CCV1 from optional to required.
    assertEq(newRequiredCCVs.length, requiredCCV.length + 1);
    assertEq(newRequiredCCVs[0].ccvAddress, OPTIONAL_CCV1);
    assertEq(newRequiredCCVs[0].args, optionalCCV[0].args); // Should preserve args.
    assertEq(newRequiredCCVs[1].ccvAddress, REQUIRED_CCV1);
    assertEq(newRequiredCCVs[1].args, requiredCCV[0].args);

    // Optional should have one less CCV.
    assertEq(newOptionalCCVs.length, optionalCCV.length - 1);
    assertEq(newOptionalCCVs[0].ccvAddress, OPTIONAL_CCV2);
    assertEq(newOptionalCCVs[0].args, optionalCCV[1].args);

    // Threshold should be decremented to maintain minimum verification count.
    assertEq(newOptionalThreshold, 1);
  }

  function test_deduplicateCCVs_SkipsDuplicatesInPoolRequiredCCV() public view {
    // Setup pool CCVs with duplicates.
    address[] memory poolRequiredCCV = new address[](3);
    poolRequiredCCV[0] = POOL_CCV1;
    poolRequiredCCV[1] = POOL_CCV1; // Duplicate
    poolRequiredCCV[2] = POOL_CCV2;

    Client.CCV[] memory requiredCCV = new Client.CCV[](0);
    Client.CCV[] memory optionalCCV = new Client.CCV[](0);
    uint8 optionalThreshold = 0;

    (Client.CCV[] memory newRequiredCCVs, Client.CCV[] memory newOptionalCCVs, uint8 newOptionalThreshold) =
      s_ccvProxyTestHelper.deduplicateCCVs(poolRequiredCCV, requiredCCV, optionalCCV, optionalThreshold);

    // Should only add unique pool CCVs.
    assertEq(newRequiredCCVs.length, poolRequiredCCV.length - 1); // One duplicate removed.
    assertEq(newOptionalCCVs.length, optionalCCV.length);
    assertEq(newOptionalThreshold, optionalThreshold);
    assertEq(newRequiredCCVs[0].ccvAddress, POOL_CCV1);
    assertEq(newRequiredCCVs[1].ccvAddress, POOL_CCV2);
  }

  function test_deduplicateCCVs_MovesAllOptionalToRequired() public view {
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
      s_ccvProxyTestHelper.deduplicateCCVs(poolRequiredCCV, requiredCCV, optionalCCV, optionalThreshold);

    // All optionals should be moved to required.
    assertEq(newRequiredCCVs.length, poolRequiredCCV.length);
    assertEq(newRequiredCCVs[0].ccvAddress, OPTIONAL_CCV1);
    assertEq(newRequiredCCVs[0].args, optionalCCV[0].args);
    assertEq(newRequiredCCVs[1].ccvAddress, OPTIONAL_CCV2);
    assertEq(newRequiredCCVs[1].args, optionalCCV[1].args);

    // Optional array should be empty.
    assertEq(newOptionalCCVs.length, optionalCCV.length - poolRequiredCCV.length);

    // Threshold should be 0 (decremented for each moved CCV).
    assertEq(newOptionalThreshold, optionalThreshold - optionalCCV.length);
  }

  function test_deduplicateCCVs_NoChangesWhenPoolCCVAlreadyInRequired() public view {
    address[] memory poolRequiredCCV = new address[](1);
    poolRequiredCCV[0] = REQUIRED_CCV1; // Already in required

    Client.CCV[] memory requiredCCV = new Client.CCV[](1);
    requiredCCV[0] = Client.CCV({ccvAddress: REQUIRED_CCV1, args: "required1"});

    Client.CCV[] memory optionalCCV = new Client.CCV[](1);
    optionalCCV[0] = Client.CCV({ccvAddress: OPTIONAL_CCV1, args: "optional1"});

    uint8 optionalThreshold = 1;

    (Client.CCV[] memory newRequiredCCVs, Client.CCV[] memory newOptionalCCVs, uint8 newOptionalThreshold) =
      s_ccvProxyTestHelper.deduplicateCCVs(poolRequiredCCV, requiredCCV, optionalCCV, optionalThreshold);

    // Should return original arrays unchanged.
    assertEq(newRequiredCCVs.length, requiredCCV.length);
    assertEq(newRequiredCCVs[0].ccvAddress, REQUIRED_CCV1);
    assertEq(newRequiredCCVs[0].args, requiredCCV[0].args);

    assertEq(newOptionalCCVs.length, optionalCCV.length);
    assertEq(newOptionalCCVs[0].ccvAddress, OPTIONAL_CCV1);
    assertEq(newOptionalCCVs[0].args, optionalCCV[0].args);

    assertEq(newOptionalThreshold, optionalThreshold);
  }

  function test_deduplicateCCVs_EmptyPoolRequiredCCV() public view {
    address[] memory poolRequiredCCV = new address[](0);

    Client.CCV[] memory requiredCCV = new Client.CCV[](1);
    requiredCCV[0] = Client.CCV({ccvAddress: REQUIRED_CCV1, args: "required1"});

    Client.CCV[] memory optionalCCV = new Client.CCV[](1);
    optionalCCV[0] = Client.CCV({ccvAddress: OPTIONAL_CCV1, args: "optional1"});

    uint8 optionalThreshold = 1;

    (Client.CCV[] memory newRequiredCCVs, Client.CCV[] memory newOptionalCCVs, uint8 newOptionalThreshold) =
      s_ccvProxyTestHelper.deduplicateCCVs(poolRequiredCCV, requiredCCV, optionalCCV, optionalThreshold);

    // Should return original arrays unchanged.
    assertEq(newRequiredCCVs.length, requiredCCV.length);
    assertEq(newRequiredCCVs[0].ccvAddress, REQUIRED_CCV1);
    assertEq(newRequiredCCVs[0].args, requiredCCV[0].args);

    assertEq(newOptionalCCVs.length, optionalCCV.length);
    assertEq(newOptionalCCVs[0].ccvAddress, OPTIONAL_CCV1);
    assertEq(newOptionalCCVs[0].args, optionalCCV[0].args);

    assertEq(newOptionalThreshold, optionalThreshold);
  }
}
