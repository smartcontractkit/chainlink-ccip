// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseERC20} from "../../../tmp/BaseERC20.sol";
import {CrossChainToken} from "../../../tmp/CrossChainToken.sol";
import {BaseTest} from "../../BaseTest.t.sol";
import {LombardTokenPoolHelper} from "../../helpers/LombardTokenPoolHelper.sol";
import {MockLombardBridge} from "../../mocks/MockLombardBridge.sol";
import {MockVerifier} from "../../mocks/MockVerifier.sol";

import {IERC20Metadata} from "@openzeppelin/contracts@5.3.0/token/ERC20/extensions/IERC20Metadata.sol";

contract LombardTokenPool_getTokenDecimals is BaseTest {
  CrossChainToken internal s_token;
  LombardTokenPoolHelper internal s_helper;
  MockVerifier internal s_resolver;
  MockLombardBridge internal s_bridge;

  function setUp() public override {
    super.setUp();
    s_token = new CrossChainToken(
      BaseERC20.ConstructorParams({
        name: "Lombard", symbol: "LBD", decimals: 18, maxSupply: 0, preMint: 0, ccipAdmin: OWNER
      }),
      OWNER,
      OWNER
    );
    s_resolver = new MockVerifier("");
    s_bridge = new MockLombardBridge();
    s_helper = new LombardTokenPoolHelper(
      IERC20Metadata(address(s_token)),
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
    uint8 dec = s_helper.getTokenDecimals(IERC20Metadata(address(s_token)), 6);
    assertEq(dec, 18);
  }

  function test_getTokenDecimals_FallsBackOnRevert() public {
    vm.mockCallRevert(address(s_token), abi.encodeWithSelector(IERC20Metadata.decimals.selector), "revert");
    uint8 dec = s_helper.getTokenDecimals(IERC20Metadata(address(s_token)), 6);
    assertEq(dec, 6);
  }
}
