// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IBridgeV1} from "../../../interfaces/lombard/IBridgeV1.sol";
import {LombardTokenPool} from "../../../pools/Lombard/LombardTokenPool.sol";

import {MockLombardBridge} from "../../mocks/MockLombardBridge.sol";
import {MockVerifier} from "../../mocks/MockVerifier.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";
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
    LombardTokenPool pool = new LombardTokenPool(s_token, s_resolver, s_bridge, adapter, address(0), RMN, ROUTER, 18);

    (address verifierResolver, address bridge, address tokenAdapter) = pool.getLombardConfig();
    assertEq(verifierResolver, s_resolver);
    assertEq(bridge, address(s_bridge));
    assertEq(tokenAdapter, adapter);
    assertEq(pool.typeAndVersion(), "LombardTokenPool 1.7.0-dev");
  }

  function test_constructor_ZeroVerifierNotAllowed() public {
    vm.expectRevert(LombardTokenPool.ZeroVerifierNotAllowed.selector);
    new LombardTokenPool(s_token, address(0), s_bridge, address(0), address(0), RMN, ROUTER, 18);
  }

  function test_constructor_RevertsWhen_InvalidMessageVersion() public {
    uint8 wrongVersion = 2;
    vm.mockCall(address(s_bridge), abi.encodeWithSelector(IBridgeV1.MSG_VERSION.selector), abi.encode(wrongVersion));

    vm.expectRevert(abi.encodeWithSelector(LombardTokenPool.InvalidMessageVersion.selector, 1, wrongVersion));
    new LombardTokenPool(s_token, s_resolver, s_bridge, address(0), address(0), RMN, ROUTER, 18);
  }

  function test_constructor_RevertsWhen_ZeroBridge() public {
    vm.expectRevert(LombardTokenPool.ZeroBridge.selector);
    new LombardTokenPool(s_token, s_resolver, IBridgeV1(address(0)), address(0), address(0), RMN, ROUTER, 18);
  }
}
