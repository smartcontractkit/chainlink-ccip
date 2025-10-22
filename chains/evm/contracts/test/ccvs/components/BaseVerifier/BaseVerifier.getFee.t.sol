// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {BaseVerifier} from "../../../../ccvs/components/BaseVerifier.sol";
import {Client} from "../../../../libraries/Client.sol";
import {BaseVerifierSetup} from "./BaseVerifierSetup.t.sol";

contract BaseVerifier_getFee is BaseVerifierSetup {
  function test_getFee() public view {
    Client.EVM2AnyMessage memory message;
    (uint256 feeUSDCents, uint32 gasForVerification, uint32 payloadSizeBytes) =
      s_baseVerifier.getFee(DEST_CHAIN_SELECTOR, message, "", 0);

    assertEq(feeUSDCents, DEFAULT_CCV_FEE_USD_CENTS);
    assertEq(gasForVerification, DEFAULT_CCV_GAS_LIMIT);
    assertEq(payloadSizeBytes, DEFAULT_CCV_PAYLOAD_SIZE);
  }

  function test_getFee_RevertWhen_DestinationNotSupported() public {
    uint64 wrongDestChainSelector = DEST_CHAIN_SELECTOR + 1;
    Client.EVM2AnyMessage memory message;

    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.DestinationNotSupported.selector, wrongDestChainSelector));
    s_baseVerifier.getFee(wrongDestChainSelector, message, "", 0);
  }
}
