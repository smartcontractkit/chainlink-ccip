// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../../interfaces/IPoolV2.sol";
import {ITokenAdminRegistry} from "../../../interfaces/ITokenAdminRegistry.sol";
import {OnRamp} from "../../../onRamp/OnRamp.sol";

import {OnRampTestHelper} from "../../helpers/OnRampTestHelper.sol";
import {OnRampSetup} from "./OnRampSetup.t.sol";

import {IERC165} from "@openzeppelin/contracts@5.0.2/utils/introspection/IERC165.sol";

contract MockPoolV2 {
  address[] internal s_requiredCCVs;

  constructor(
    address[] memory requiredCCVs
  ) {
    s_requiredCCVs = requiredCCVs;
  }

  function getRequiredCCVs(
    address,
    uint64,
    uint256,
    uint16,
    bytes memory,
    IPoolV2.CCVDirection direction
  ) external view returns (address[] memory) {
    if (direction != IPoolV2.CCVDirection.Outbound) revert("direction");
    return s_requiredCCVs;
  }

  function supportsInterface(
    bytes4 interfaceId
  ) external pure returns (bool) {
    return interfaceId == type(IPoolV2).interfaceId || interfaceId == type(IERC165).interfaceId;
  }
}

contract OnRamp_getCCVsForPool is OnRampSetup {
  OnRampTestHelper internal s_onRampTestHelper;
  address internal s_token;
  address internal s_helperDefaultCCV;

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

    s_token = makeAddr("token");
    _configureHelperDestChain();
  }

  function test_getCCVsForPool_ReturnsPoolCCVs_WhenPoolSupportsV2() public {
    address[] memory expectedCCVs = new address[](2);
    expectedCCVs[0] = makeAddr("ccv1");
    expectedCCVs[1] = makeAddr("ccv2");

    _deployPoolV2(expectedCCVs);

    address[] memory result = s_onRampTestHelper.getCCVsForPool(DEST_CHAIN_SELECTOR, s_token, 100, 0, "");

    assertEq(result.length, expectedCCVs.length, "Should surface pool-provided CCVs");
    assertEq(result[0], expectedCCVs[0], "First CCV should match");
    assertEq(result[1], expectedCCVs[1], "Second CCV should match");
  }

  function test_getCCVsForPool_ReturnsDefaultCCVs_WhenPoolDoesNotSupportV2() public {
    address[] memory expectedCCVs = new address[](1);
    expectedCCVs[0] = makeAddr("poolCCV");
    address pool = _deployPoolV2(expectedCCVs);

    // Force the ERC165 probe to report that IPoolV2 is unsupported.
    vm.mockCall(
      pool, abi.encodeWithSelector(IERC165.supportsInterface.selector, type(IPoolV2).interfaceId), abi.encode(false)
    );

    address[] memory result = s_onRampTestHelper.getCCVsForPool(DEST_CHAIN_SELECTOR, s_token, 100, 0, "");

    assertEq(result.length, 1, "Should fall back to default CCV when pool is V1");
    assertEq(result[0], s_helperDefaultCCV, "Returned CCV should be the helper default");
  }

  function test_getCCVsForPool_ReturnsDefaultCCVs_WhenPoolReturnsEmptyArray() public {
    _deployPoolV2(new address[](0));

    address[] memory result = s_onRampTestHelper.getCCVsForPool(DEST_CHAIN_SELECTOR, s_token, 100, 0, "");

    assertEq(result.length, 1, "Should fall back to default CCV when pool is silent");
    assertEq(result[0], s_helperDefaultCCV, "Returned CCV should be the helper default");
  }

  function test_getCCVsForPool_PassesThroughAddressZeroSentinel() public {
    address[] memory poolCCVs = new address[](3);
    poolCCVs[0] = makeAddr("poolCCV1");
    poolCCVs[1] = address(0);
    poolCCVs[2] = makeAddr("poolCCV2");

    _deployPoolV2(poolCCVs);

    address[] memory result = s_onRampTestHelper.getCCVsForPool(DEST_CHAIN_SELECTOR, s_token, 100, 0, "");

    assertEq(result.length, poolCCVs.length, "Should preserve address(0) sentinel for merge step");
    assertEq(result[0], poolCCVs[0], "First CCV should match");
    assertEq(result[1], address(0), "Sentinel should be present");
    assertEq(result[2], poolCCVs[2], "Third CCV should match");
  }

  function _configureHelperDestChain() internal {
    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = makeAddr("helperDefaultCCV");
    s_helperDefaultCCV = defaultCCVs[0];

    OnRamp.DestChainConfigArgs[] memory destChainConfigArgs = new OnRamp.DestChainConfigArgs[](1);
    destChainConfigArgs[0] = OnRamp.DestChainConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      router: s_sourceRouter,
      laneMandatedCCVs: new address[](0),
      defaultCCVs: defaultCCVs,
      defaultExecutor: makeAddr("helperExecutor"),
      offRamp: abi.encodePacked(address(s_offRampOnRemoteChain))
    });

    s_onRampTestHelper.applyDestChainConfigUpdates(destChainConfigArgs);
  }

  function _deployPoolV2(
    address[] memory requiredCCVs
  ) internal returns (address pool) {
    pool = address(new MockPoolV2(requiredCCVs));

    vm.mockCall(
      pool, abi.encodeWithSelector(IERC165.supportsInterface.selector, type(IPoolV2).interfaceId), abi.encode(true)
    );
    vm.mockCall(
      pool, abi.encodeWithSelector(IERC165.supportsInterface.selector, type(IERC165).interfaceId), abi.encode(true)
    );

    _mockRegistryPool(pool);

    return pool;
  }

  function _mockRegistryPool(
    address pool
  ) internal {
    vm.mockCall(
      address(s_tokenAdminRegistry),
      abi.encodeWithSelector(ITokenAdminRegistry.getPool.selector, s_token),
      abi.encode(pool)
    );
  }
}
