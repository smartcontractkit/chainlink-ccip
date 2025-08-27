// // SPDX-License-Identifier: BUSL-1.1
// pragma solidity ^0.8.24;

// import {IPoolV2} from "../../../interfaces/IPoolV2.sol";
// import {ITokenAdminRegistry} from "../../../interfaces/ITokenAdminRegistry.sol";
// import {CCVAggregator} from "../../../offRamp/CCVAggregator.sol";
// import {CCVAggregatorHelper, CCVAggregatorSetup} from "./CCVAggregatorSetup.t.sol";

// contract MockPoolV2 {
//   address[] internal s_requiredCCVs;

//   constructor(
//     address[] memory requiredCCVs
//   ) {
//     s_requiredCCVs = requiredCCVs;
//   }

//   function getRequiredCCVs(
//     address localToken,
//     uint64 sourceChainSelector,
//     uint256 amount,
//     bytes memory extraData
//   ) external view returns (address[] memory) {
//     return s_requiredCCVs;
//   }

//   function supportsInterface(
//     bytes4 interfaceId
//   ) external pure returns (bool) {
//     return interfaceId == type(IPoolV2).interfaceId;
//   }
// }

// contract CCVAggregator_getCCVsFromPool is CCVAggregatorSetup {
//   address internal s_token;
//   address internal s_poolV2;
//   address internal s_poolV1;

//   function setUp() public override {
//     CCVAggregatorSetup.setUp();
//     s_token = makeAddr("token");
//     s_poolV2 = address(new MockPoolV2(new address[](0)));
//     s_poolV1 = makeAddr("poolV1");
//   }

//   function test_getCCVsFromPool_ReturnsEmptyArray_WhenPoolDoesNotSupportV2() public {
//     // Mock pool that doesn't support V2 interface
//     vm.mockCall(
//       s_poolV1, abi.encodeWithSignature("supportsInterface(bytes4)", type(IPoolV2).interfaceId), abi.encode(false)
//     );

//     address[] memory result = s_agg.getCCVsFromPool(s_token, SOURCE_CHAIN_SELECTOR, 100, "");
//     assertEq(result.length, 0);
//   }

//   function test_getCCVsFromPool_ReturnsPoolCCVs_WhenPoolSupportsV2() public {
//     address[] memory expectedCCVs = new address[](2);
//     expectedCCVs[0] = makeAddr("ccv1");
//     expectedCCVs[1] = makeAddr("ccv2");

//     MockPoolV2 poolV2 = new MockPoolV2(expectedCCVs);

//     // Mock token admin registry to return the V2 pool
//     vm.mockCall(
//       address(s_tokenAdminRegistry),
//       abi.encodeWithSelector(ITokenAdminRegistry.getPool.selector, s_token),
//       abi.encode(address(poolV2))
//     );

//     address[] memory result = s_agg.getCCVsFromPool(s_token, SOURCE_CHAIN_SELECTOR, 100, "");
//     assertEq(result.length, expectedCCVs.length);
//     assertEq(result[0], expectedCCVs[0]);
//     assertEq(result[1], expectedCCVs[1]);
//   }

//   function test_getCCVsFromPool_ReturnsEmptyArray_WhenPoolReturnsEmptyArray() public {
//     address[] memory emptyCCVs = new address[](0);
//     MockPoolV2 poolV2 = new MockPoolV2(emptyCCVs);

//     // Mock token admin registry to return the V2 pool
//     vm.mockCall(
//       address(s_tokenAdminRegistry),
//       abi.encodeWithSelector(ITokenAdminRegistry.getPool.selector, s_token),
//       abi.encode(address(poolV2))
//     );

//     address[] memory result = s_agg.getCCVsFromPool(s_token, SOURCE_CHAIN_SELECTOR, 100, "");
//     assertEq(result.length, 0);
//   }

//   function test_getCCVsFromPool_HandlesDifferentParameters() public {
//     address[] memory expectedCCVs = new address[](1);
//     expectedCCVs[0] = makeAddr("ccv1");

//     MockPoolV2 poolV2 = new MockPoolV2(expectedCCVs);

//     // Mock token admin registry to return the V2 pool
//     vm.mockCall(
//       address(s_tokenAdminRegistry),
//       abi.encodeWithSelector(ITokenAdminRegistry.getPool.selector, s_token),
//       abi.encode(address(poolV2))
//     );

//     address[] memory result = s_agg.getCCVsFromPool(s_token, SOURCE_CHAIN_SELECTOR + 1, 1000, "extraData");

//     assertEq(result.length, expectedCCVs.length);
//     assertEq(result[0], expectedCCVs[0]);
//   }
// }
