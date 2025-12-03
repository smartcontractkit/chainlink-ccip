// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {LombardTokenPool} from "../../../pools/Lombard/LombardTokenPool.sol";
import {MockVerifier} from "../../mocks/MockVerifier.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";
import {Test} from "forge-std/Test.sol";

contract LombardTokenPool_constructor is Test {
  BurnMintERC20 internal s_token;
  address internal s_resolver;
  address internal constant RMN = address(0xAA01);
  address internal constant ROUTER = address(0xBB02);

  function setUp() public {
    s_token = new BurnMintERC20("Lombard", "LBD", 18, 0, 0);
    s_resolver = address(new MockVerifier(""));
  }

  function test_constructor() public {
    vm.expectEmit();
    emit LombardTokenPool.LombardVerifierSet(s_resolver);
    LombardTokenPool pool = new LombardTokenPool(s_token, s_resolver, address(0), RMN, ROUTER, 18);
    assertEq(pool.getVerifierResolver(), address(s_resolver));
    assertEq(pool.typeAndVersion(), "LombardTokenPool 1.7.0-dev");
  }

  function test_constructor_ZeroVerifierNotAllowed() public {
    vm.expectRevert(LombardTokenPool.ZeroVerifierNotAllowed.selector);
    new LombardTokenPool(s_token, address(0), address(0), RMN, ROUTER, 18);
  }
}
