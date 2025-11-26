// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../../../interfaces/IPoolV2.sol";

import {Client} from "../../../libraries/Client.sol";
import {OnRamp} from "../../../onRamp/OnRamp.sol";
import {OnRampSetup} from "./OnRampSetup.t.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {IERC165} from "@openzeppelin/contracts@5.0.2/utils/introspection/IERC165.sol";

contract OnRamp_distributeFees is OnRampSetup {
  address internal s_verifier = makeAddr("verifier");
  address internal s_executor = makeAddr("executor");

  function test_distributeFees_NoTokens_PaysAllIssuers() public {
    Client.EVM2AnyMessage memory message = _generateEmptyMessage();

    (OnRamp.Receipt[] memory receipts, uint256 totalFee) = _buildReceipts(false);
    deal(s_sourceFeeToken, address(s_onRamp), totalFee);

    s_onRamp.distributeFees(message, receipts);

    IERC20 feeToken = IERC20(s_sourceFeeToken);
    assertEq(feeToken.balanceOf(s_verifier), receipts[0].feeTokenAmount);
    assertEq(feeToken.balanceOf(s_executor), receipts[1].feeTokenAmount);
    assertEq(feeToken.balanceOf(address(s_onRamp)), 0);
  }

  function test_distributeFees_TokenPoolV2ReceivesShare() public {
    address token = s_sourceTokens[0];
    Client.EVM2AnyMessage memory message = _generateSingleTokenMessage(token, 10 ether);

    (OnRamp.Receipt[] memory receipts, uint256 totalFee) = _buildReceipts(true);
    deal(s_sourceFeeToken, address(s_onRamp), totalFee);

    s_onRamp.distributeFees(message, receipts);

    IERC20 feeToken = IERC20(s_sourceFeeToken);
    assertEq(feeToken.balanceOf(s_verifier), receipts[0].feeTokenAmount);
    assertEq(feeToken.balanceOf(s_executor), receipts[2].feeTokenAmount);
    assertEq(feeToken.balanceOf(s_sourcePoolByToken[token]), receipts[1].feeTokenAmount);
    assertEq(feeToken.balanceOf(address(s_onRamp)), 0);
  }

  function test_distributeFees_TokenPoolV1RetainsInOnRamp() public {
    address token = s_sourceTokens[0];
    Client.EVM2AnyMessage memory message = _generateSingleTokenMessage(token, 5 ether);

    (OnRamp.Receipt[] memory receipts, uint256 totalFee) = _buildReceipts(true);
    deal(s_sourceFeeToken, address(s_onRamp), totalFee);

    address pool = s_sourcePoolByToken[token];
    vm.mockCall(
      pool, abi.encodeWithSelector(IERC165.supportsInterface.selector, type(IPoolV2).interfaceId), abi.encode(false)
    );

    s_onRamp.distributeFees(message, receipts);

    IERC20 feeToken = IERC20(s_sourceFeeToken);
    assertEq(feeToken.balanceOf(s_verifier), receipts[0].feeTokenAmount);
    assertEq(feeToken.balanceOf(s_executor), receipts[2].feeTokenAmount);
    assertEq(feeToken.balanceOf(pool), 0);
    assertEq(feeToken.balanceOf(address(s_onRamp)), receipts[1].feeTokenAmount);
  }

  function _buildReceipts(
    bool withToken
  ) internal view returns (OnRamp.Receipt[] memory receipts, uint256 totalFee) {
    if (withToken) {
      receipts = new OnRamp.Receipt[](3);
      receipts[0] =
        OnRamp.Receipt({issuer: s_verifier, destGasLimit: 0, destBytesOverhead: 0, feeTokenAmount: 1e16, extraArgs: ""});
      receipts[1] = OnRamp.Receipt({
        issuer: s_sourcePoolByToken[s_sourceTokens[0]],
        destGasLimit: 0,
        destBytesOverhead: 0,
        feeTokenAmount: 2e16,
        extraArgs: ""
      });
      receipts[2] =
        OnRamp.Receipt({issuer: s_executor, destGasLimit: 0, destBytesOverhead: 0, feeTokenAmount: 3e16, extraArgs: ""});
      totalFee = receipts[0].feeTokenAmount + receipts[1].feeTokenAmount + receipts[2].feeTokenAmount;
    } else {
      receipts = new OnRamp.Receipt[](2);
      receipts[0] =
        OnRamp.Receipt({issuer: s_verifier, destGasLimit: 0, destBytesOverhead: 0, feeTokenAmount: 1e17, extraArgs: ""});
      receipts[1] =
        OnRamp.Receipt({issuer: s_executor, destGasLimit: 0, destBytesOverhead: 0, feeTokenAmount: 2e17, extraArgs: ""});
      totalFee = receipts[0].feeTokenAmount + receipts[1].feeTokenAmount;
    }
    return (receipts, totalFee);
  }
}
