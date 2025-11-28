// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRouter} from "../../../interfaces/IRouter.sol";

import {LombardVerifier} from "../../../ccvs/LombardVerifier.sol";
import {BaseVerifier} from "../../../ccvs/components/BaseVerifier.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {LombardVerifierSetup} from "./LombardVerifierSetup.t.sol";

contract LombardVerifier_verifyMessage is LombardVerifierSetup {
  function test_verifyMessage() public {
    // Create the ccvData with empty payload and proof.
    bytes memory ccvData = abi.encode("", "");

    vm.startPrank(s_offRamp);
    // Should not revert - the mock mailbox returns success by default.
    s_lombardVerifier.verifyMessage(_createBasicMessageV1(DEST_CHAIN_SELECTOR), bytes32(0), ccvData);
  }

  function test_verifyMessage_RevertWhen_CallerIsNotOffRamp() public {
    address invalidCaller = makeAddr("invalidCaller");

    // Mock the router to return false for isOffRamp with the invalid caller.
    vm.mockCall(
      address(s_router), abi.encodeCall(IRouter.isOffRamp, (DEST_CHAIN_SELECTOR, invalidCaller)), abi.encode(false)
    );

    MessageV1Codec.MessageV1 memory message = _createBasicMessageV1(DEST_CHAIN_SELECTOR);

    vm.startPrank(invalidCaller);

    vm.expectRevert(abi.encodeWithSelector(BaseVerifier.CallerIsNotARampOnRouter.selector, invalidCaller));
    s_lombardVerifier.verifyMessage(message, bytes32(0), "");
  }

  function test_verifyMessage_RevertWhen_ExecutionError() public {
    // Make the mailbox fail.
    s_mockMailbox.setShouldSucceed(false);

    vm.startPrank(s_offRamp);

    vm.expectRevert(abi.encodeWithSelector(LombardVerifier.ExecutionError.selector));
    s_lombardVerifier.verifyMessage(_createBasicMessageV1(DEST_CHAIN_SELECTOR), bytes32(0), abi.encode("", ""));
  }
}
