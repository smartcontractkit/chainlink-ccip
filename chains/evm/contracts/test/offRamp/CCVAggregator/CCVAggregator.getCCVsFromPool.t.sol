// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../../interfaces/IPoolV2.sol";
import {ITokenAdminRegistry} from "../../../interfaces/ITokenAdminRegistry.sol";
import {CCVAggregatorSetup} from "./CCVAggregatorSetup.t.sol";

import {IERC165} from "@openzeppelin/contracts@5.0.2/utils/introspection/IERC165.sol";

contract MockPoolV2 {
  address[] internal s_requiredCCVs;

  constructor(
    address[] memory requiredCCVs
  ) {
    s_requiredCCVs = requiredCCVs;
  }

  function getRequiredInboundCCVs(
    address,
    uint64,
    uint256,
    uint16,
    bytes memory
  ) external view returns (address[] memory) {
    return s_requiredCCVs;
  }

  function supportsInterface(
    bytes4 interfaceId
  ) external pure returns (bool) {
    return interfaceId == type(IPoolV2).interfaceId || interfaceId == type(IERC165).interfaceId;
  }
}

contract CCVAggregator_getCCVsFromPool is CCVAggregatorSetup {
  address internal s_token;

  function setUp() public override {
    super.setUp();
    s_token = makeAddr("token");
  }

  function test_getCCVsFromPool_ReturnsDefaultCCVs_WhenPoolDoesNotSupportV2() public {
    address[] memory expectedCCVs = new address[](2);
    expectedCCVs[0] = makeAddr("ccv1");
    expectedCCVs[1] = makeAddr("ccv2");
    address pool = _deployPoolV2(expectedCCVs);

    // Mock pool that doesn't support V2 interface
    vm.mockCall(
      pool, abi.encodeWithSignature("supportsInterface(bytes4)", type(IPoolV2).interfaceId), abi.encode(false)
    );

    address[] memory result = s_agg.getCCVsFromPool(s_token, SOURCE_CHAIN_SELECTOR, 100, 0, "");
    assertEq(result.length, 1);
    assertEq(result[0], s_defaultCCV);
  }

  function test_getCCVsFromPool_ReturnsPoolCCVs_WhenPoolSupportsV2() public {
    address[] memory expectedCCVs = new address[](2);
    expectedCCVs[0] = makeAddr("ccv1");
    expectedCCVs[1] = makeAddr("ccv2");

    _deployPoolV2(expectedCCVs);

    address[] memory result = s_agg.getCCVsFromPool(s_token, SOURCE_CHAIN_SELECTOR, 100, 0, "");
    assertEq(result.length, expectedCCVs.length);
    assertEq(result[0], expectedCCVs[0]);
    assertEq(result[1], expectedCCVs[1]);
  }

  function test_getCCVsFromPool_ReturnsDefaultCCVs_WhenPoolReturnsEmptyArray() public {
    address[] memory emptyCCVs = new address[](0);
    _deployPoolV2(emptyCCVs);

    address[] memory result = s_agg.getCCVsFromPool(s_token, SOURCE_CHAIN_SELECTOR, 100, 0, "");
    assertEq(result.length, 1);
    assertEq(result[0], s_defaultCCV);
  }

  function _deployPoolV2(
    address[] memory requiredCCVs
  ) internal returns (address) {
    address pool = address(new MockPoolV2(requiredCCVs));
    vm.mockCall(pool, abi.encodeWithSignature("supportsInterface(bytes4)", type(IPoolV2).interfaceId), abi.encode(true));
    vm.mockCall(pool, abi.encodeWithSignature("supportsInterface(bytes4)", type(IERC165).interfaceId), abi.encode(true));

    // Mock token admin registry to return the V2 pool
    vm.mockCall(
      address(s_tokenAdminRegistry),
      abi.encodeWithSelector(ITokenAdminRegistry.getPool.selector, s_token),
      abi.encode(pool)
    );
    return pool;
  }
}
