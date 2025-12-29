// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IBridgeV2} from "../../../interfaces/lombard/IBridgeV2.sol";
import {LombardTokenPool} from "../../../pools/Lombard/LombardTokenPool.sol";

import {MockLombardBridge} from "../../mocks/MockLombardBridge.sol";
import {MockVerifier} from "../../mocks/MockVerifier.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";
import {IERC20Metadata} from "@openzeppelin/contracts@5.3.0/token/ERC20/extensions/IERC20Metadata.sol";
import {Test} from "forge-std/Test.sol";

contract LombardTokenPool_constructor is Test {
  BurnMintERC20 internal s_token;
  MockLombardBridge internal s_bridge;
  address internal s_resolver;
  address internal constant RMN = address(0xAA01);
  address internal constant ROUTER = address(0xBB02);

  function setUp() public {
    s_token = new BurnMintERC20("Lombard", "LBD", 18, 0, 0);
    s_resolver = address(new MockVerifier(""));
    s_bridge = new MockLombardBridge();
  }

  function test_constructor() public {
    address adapter = makeAddr("adapter");
    LombardTokenPool pool = new LombardTokenPool(
      IERC20Metadata(address(s_token)), s_resolver, s_bridge, adapter, address(0), RMN, ROUTER, 18
    );

    (address verifierResolver, address bridge, address tokenAdapter) = pool.getLombardConfig();
    assertEq(verifierResolver, s_resolver);
    assertEq(bridge, address(s_bridge));
    assertEq(tokenAdapter, adapter);
    assertEq(pool.typeAndVersion(), "LombardTokenPool 1.7.0-dev");
    assertEq(s_token.allowance(address(pool), adapter), type(uint256).max);
    assertEq(s_token.allowance(address(pool), bridge), 0);
    assertEq(s_token.allowance(address(pool), address(s_bridge)), 0);
  }

  function test_constructor_WithoutAdapterApprovesBridge() public {
    LombardTokenPool pool = new LombardTokenPool(
      IERC20Metadata(address(s_token)), s_resolver, s_bridge, address(0), address(0), RMN, ROUTER, 18
    );

    assertEq(s_token.allowance(address(pool), address(s_bridge)), type(uint256).max);
    assertEq(s_token.allowance(address(pool), address(0)), 0);
  }

  function test_constructor_ZeroVerifierNotAllowed() public {
    vm.expectRevert(LombardTokenPool.ZeroVerifierNotAllowed.selector);
    new LombardTokenPool(
      IERC20Metadata(address(s_token)), address(0), s_bridge, address(0), address(0), RMN, ROUTER, 18
    );
  }

  function test_constructor_RevertsWhen_InvalidMessageVersion() public {
    uint8 wrongVersion = 2;
    vm.mockCall(address(s_bridge), abi.encodeWithSelector(IBridgeV2.MSG_VERSION.selector), abi.encode(wrongVersion));

    vm.expectRevert(abi.encodeWithSelector(LombardTokenPool.InvalidMessageVersion.selector, 1, wrongVersion));
    new LombardTokenPool(
      IERC20Metadata(address(s_token)), s_resolver, s_bridge, address(0), address(0), RMN, ROUTER, 18
    );
  }

  function test_constructor_RevertsWhen_ZeroBridge() public {
    vm.expectRevert(LombardTokenPool.ZeroBridge.selector);
    new LombardTokenPool(
      IERC20Metadata(address(s_token)), s_resolver, IBridgeV2(address(0)), address(0), address(0), RMN, ROUTER, 18
    );
  }
}
