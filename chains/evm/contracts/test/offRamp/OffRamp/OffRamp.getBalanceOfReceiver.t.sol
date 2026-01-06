// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {OffRamp} from "../../../offRamp/OffRamp.sol";
import {OffRampHelper} from "../../helpers/OffRampHelper.sol";
import {OffRampSetup} from "./OffRampSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@5.3.0/token/ERC20/IERC20.sol";

contract OffRamp_getBalanceOfReceiver is OffRampSetup {
  OffRampHelper internal s_offRampHelper;

  function setUp() public override {
    super.setUp();

    s_offRampHelper = new OffRampHelper(
      OffRamp.StaticConfig({
        localChainSelector: SOURCE_CHAIN_SELECTOR,
        gasForCallExactCheck: GAS_FOR_CALL_EXACT_CHECK,
        rmnRemote: s_mockRMNRemote,
        tokenAdminRegistry: address(s_tokenAdminRegistry),
        maxGasBufferToUpdateState: DEFAULT_MAX_GAS_BUFFER_TO_UPDATE_STATE
      })
    );
  }

  function test_getBalanceOfReceiver() public {
    address receiver = makeAddr("receiver");
    address token = makeAddr("token");
    uint256 balance = 1000 ether;

    vm.mockCall(token, abi.encodeCall(IERC20.balanceOf, (receiver)), abi.encode(balance));

    uint256 result = s_offRampHelper.getBalanceOfReceiver(receiver, token);
    assertEq(result, balance);
  }

  function test_getBalanceOfReceiver_RevertWhen_TokenHandlingError() public {
    address receiver = makeAddr("receiver");
    address token = makeAddr("token");
    bytes memory revertData = abi.encode("balanceOf failed");

    vm.mockCallRevert(token, abi.encodeCall(IERC20.balanceOf, (receiver)), revertData);

    vm.expectRevert(abi.encodeWithSelector(OffRamp.TokenHandlingError.selector, token, revertData));
    s_offRampHelper.getBalanceOfReceiver(receiver, token);
  }
}

