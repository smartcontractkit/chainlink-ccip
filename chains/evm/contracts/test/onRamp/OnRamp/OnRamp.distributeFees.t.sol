// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../../interfaces/IPoolV2.sol";
import {Client} from "../../../libraries/Client.sol";
import {OnRamp} from "../../../onRamp/OnRamp.sol";
import {OnRampSetup} from "./OnRampSetup.t.sol";
import {stdError} from "forge-std/Test.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {IERC165} from "@openzeppelin/contracts@5.0.2/utils/introspection/IERC165.sol";

contract OnRamp_distributeFees is OnRampSetup {
  function test_distributeFees_NoTokens_PaysAllIssuers() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();

    address verifier = makeAddr("verifier");
    address executor = makeAddr("executor");

    OnRamp.Receipt[] memory receipts = new OnRamp.Receipt[](2);
    receipts[0] =
      OnRamp.Receipt({issuer: verifier, destGasLimit: 0, destBytesOverhead: 0, feeTokenAmount: 1e17, extraArgs: ""});
    receipts[1] =
      OnRamp.Receipt({issuer: executor, destGasLimit: 0, destBytesOverhead: 0, feeTokenAmount: 2e17, extraArgs: ""});

    uint256 totalFee = receipts[0].feeTokenAmount + receipts[1].feeTokenAmount;
    deal(s_sourceFeeToken, address(s_onRamp), totalFee);

    s_onRamp.distributeFees(DEST_CHAIN_SELECTOR, message, receipts);

    IERC20 feeToken = IERC20(s_sourceFeeToken);
    assertEq(feeToken.balanceOf(verifier), receipts[0].feeTokenAmount);
    assertEq(feeToken.balanceOf(executor), receipts[1].feeTokenAmount);
    assertEq(feeToken.balanceOf(address(s_onRamp)), 0);
  }

  function test_distributeFees_TokenPoolV2ReceivesShare() public {
    address token = s_sourceTokens[0];
    Client.EVM2AnyMessage memory message = _generateSingleTokenMessage(token, 10 ether);

    address verifier = makeAddr("verifierV2");
    address executor = makeAddr("executorV2");

    OnRamp.Receipt[] memory receipts = new OnRamp.Receipt[](3);
    receipts[0] =
      OnRamp.Receipt({issuer: verifier, destGasLimit: 0, destBytesOverhead: 0, feeTokenAmount: 1e16, extraArgs: ""});
    receipts[1] =
      OnRamp.Receipt({issuer: token, destGasLimit: 0, destBytesOverhead: 0, feeTokenAmount: 2e16, extraArgs: ""});
    receipts[2] =
      OnRamp.Receipt({issuer: executor, destGasLimit: 0, destBytesOverhead: 0, feeTokenAmount: 3e16, extraArgs: ""});

    uint256 totalFee = receipts[0].feeTokenAmount + receipts[1].feeTokenAmount + receipts[2].feeTokenAmount;
    deal(s_sourceFeeToken, address(s_onRamp), totalFee);

    address pool = s_sourcePoolByToken[token];
    vm.mockCall(
      pool, abi.encodeWithSelector(IERC165.supportsInterface.selector, type(IPoolV2).interfaceId), abi.encode(true)
    );

    s_onRamp.distributeFees(DEST_CHAIN_SELECTOR, message, receipts);

    IERC20 feeToken = IERC20(s_sourceFeeToken);
    assertEq(feeToken.balanceOf(verifier), receipts[0].feeTokenAmount);
    assertEq(feeToken.balanceOf(executor), receipts[2].feeTokenAmount);
    assertEq(feeToken.balanceOf(pool), receipts[1].feeTokenAmount);
    assertEq(feeToken.balanceOf(address(s_onRamp)), 0);
    assertEq(feeToken.balanceOf(receipts[1].issuer), 0);
  }

  function test_distributeFees_TokenPoolV1RetainsInOnRamp() public {
    address token = s_sourceTokens[0];
    Client.EVM2AnyMessage memory message = _generateSingleTokenMessage(token, 5 ether);

    address verifier = makeAddr("verifierV1");
    address executor = makeAddr("executorV1");

    OnRamp.Receipt[] memory receipts = new OnRamp.Receipt[](3);
    receipts[0] =
      OnRamp.Receipt({issuer: verifier, destGasLimit: 0, destBytesOverhead: 0, feeTokenAmount: 4e16, extraArgs: ""});
    receipts[1] =
      OnRamp.Receipt({issuer: token, destGasLimit: 0, destBytesOverhead: 0, feeTokenAmount: 6e16, extraArgs: ""});
    receipts[2] =
      OnRamp.Receipt({issuer: executor, destGasLimit: 0, destBytesOverhead: 0, feeTokenAmount: 7e16, extraArgs: ""});

    uint256 totalFee = receipts[0].feeTokenAmount + receipts[1].feeTokenAmount + receipts[2].feeTokenAmount;
    deal(s_sourceFeeToken, address(s_onRamp), totalFee);

    address pool = s_sourcePoolByToken[token];
    vm.mockCall(
      pool, abi.encodeWithSelector(IERC165.supportsInterface.selector, type(IPoolV2).interfaceId), abi.encode(false)
    );

    s_onRamp.distributeFees(DEST_CHAIN_SELECTOR, message, receipts);

    IERC20 feeToken = IERC20(s_sourceFeeToken);
    assertEq(feeToken.balanceOf(verifier), receipts[0].feeTokenAmount);
    assertEq(feeToken.balanceOf(executor), receipts[2].feeTokenAmount);
    assertEq(feeToken.balanceOf(pool), 0);
    assertEq(feeToken.balanceOf(address(s_onRamp)), receipts[1].feeTokenAmount);
  }
}
