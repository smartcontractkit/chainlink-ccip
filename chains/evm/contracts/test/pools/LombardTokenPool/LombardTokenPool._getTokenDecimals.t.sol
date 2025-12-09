// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseTest} from "../../BaseTest.t.sol";
import {LombardTokenPoolHelper} from "../../helpers/LombardTokenPoolHelper.sol";

import {MockLombardBridgeV1} from "../../mocks/MockLombardBridgeV1.sol";
import {MockVerifier} from "../../mocks/MockVerifier.sol";
import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

import {IERC20Metadata} from "@openzeppelin/contracts@4.8.3/token/ERC20/extensions/IERC20Metadata.sol";

contract LombardTokenPool_getTokenDecimals is BaseTest {
  BurnMintERC20 internal s_token;
  LombardTokenPoolHelper internal s_helper;
  MockVerifier internal s_resolver;
  MockLombardBridgeV1 internal s_bridge;

  function setUp() public override {
    super.setUp();
    s_token = new BurnMintERC20("Lombard", "LBD", 18, 0, 0);
    s_resolver = new MockVerifier("");
    s_bridge = new MockLombardBridgeV1(1, address(0));
    s_helper = new LombardTokenPoolHelper(
      s_token,
      address(s_resolver),
      s_bridge,
      address(0),
      address(0),
      address(s_mockRMNRemote),
      address(s_sourceRouter),
      18
    );
  }

  function test_getTokenDecimals_UsesTokenDecimals() public view {
    uint8 dec = s_helper.getTokenDecimals(s_token, 6);
    assertEq(dec, 18);
  }

  function test_getTokenDecimals_FallsBackOnRevert() public {
    vm.mockCallRevert(address(s_token), abi.encodeWithSelector(IERC20Metadata.decimals.selector), "revert");
    uint8 dec = s_helper.getTokenDecimals(IERC20Metadata(address(s_token)), 6);
    assertEq(dec, 6);
  }
}
