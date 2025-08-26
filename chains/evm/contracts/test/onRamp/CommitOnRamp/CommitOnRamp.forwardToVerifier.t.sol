// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../../libraries/Internal.sol";
import {BaseOnRamp} from "../../../onRamp/BaseOnRamp.sol";
import {CommitOnRamp} from "../../../onRamp/CommitOnRamp.sol";
import {CommitOnRampSetup} from "./CommitOnRampSetup.t.sol";

contract CommitOnRamp_forwardToVerifier_Test is CommitOnRampSetup {
  function setUp() public override {
    super.setUp();
  }

  function test_forwardToVerifier() public {
    Internal.EVM2AnyVerifierMessage memory message =
      _createEVM2AnyVerifierMessage(DEST_CHAIN_SELECTOR, msg.sender, "test data", msg.sender, s_sourceFeeToken, 1000);

    bytes memory rawMessage = abi.encode(message);
    uint256 verifierIndex = 0;

    // Mock the fee quoter response
    vm.mockCall(
      address(s_feeQuoter),
      abi.encodeWithSelector(
        s_feeQuoter.processMessageArgs.selector, DEST_CHAIN_SELECTOR, s_sourceFeeToken, 1000, "", abi.encode(msg.sender)
      ),
      abi.encode(0, false, "", "")
    );
    vm.mockCall(
      address(s_nonceManager),
      abi.encodeWithSelector(s_nonceManager.getIncrementedOutboundNonce.selector, DEST_CHAIN_SELECTOR, msg.sender),
      abi.encode(1)
    );

    vm.stopPrank();
    vm.prank(s_ccvProxy);
    bytes memory result = s_commitOnRamp.forwardToVerifier(rawMessage, verifierIndex);

    uint64 nonce = abi.decode(result, (uint64));
    assertEq(nonce, 1);
  }

  function test_forwardToVerifier_WithOutOfOrderExecution() public {
    Internal.EVM2AnyVerifierMessage memory message =
      _createEVM2AnyVerifierMessage(DEST_CHAIN_SELECTOR, msg.sender, "test data", msg.sender, s_sourceFeeToken, 1000);

    bytes memory rawMessage = abi.encode(message);
    uint256 verifierIndex = 0;

    // Mock the fee quoter response with out of order execution true
    vm.mockCall(
      address(s_feeQuoter),
      abi.encodeWithSelector(
        s_feeQuoter.processMessageArgs.selector, DEST_CHAIN_SELECTOR, s_sourceFeeToken, 1000, "", abi.encode(msg.sender)
      ),
      abi.encode(0, true, "", "")
    );

    vm.stopPrank();
    vm.prank(s_ccvProxy);
    bytes memory result = s_commitOnRamp.forwardToVerifier(rawMessage, verifierIndex);

    uint64 nonce = abi.decode(result, (uint64));
    assertEq(nonce, 0); // Should return 0 for out of order execution
  }

  function test_forwardToVerifier_RevertWhen_CalledByNons_ccvProxy() public {
    Internal.EVM2AnyVerifierMessage memory message =
      _createEVM2AnyVerifierMessage(DEST_CHAIN_SELECTOR, msg.sender, "test data", msg.sender, s_sourceFeeToken, 1000);

    bytes memory rawMessage = abi.encode(message);
    uint256 verifierIndex = 0;

    vm.stopPrank();
    vm.prank(STRANGER);
    vm.expectRevert(BaseOnRamp.MustBeCalledByCCVProxy.selector);
    s_commitOnRamp.forwardToVerifier(rawMessage, verifierIndex);
  }

  function test_forwardToVerifier_RevertWhen_s_ccvProxyNotSet() public {
    CommitOnRamp commitOnRampWithoutCCVProxy = new CommitOnRamp(
      address(s_mockRMNRemote),
      address(s_nonceManager),
      CommitOnRamp.DynamicConfig({
        feeQuoter: address(s_feeQuoter),
        feeAggregator: FEE_AGGREGATOR,
        allowlistAdmin: ALLOWLIST_ADMIN
      })
    );

    Internal.EVM2AnyVerifierMessage memory message =
      _createEVM2AnyVerifierMessage(DEST_CHAIN_SELECTOR, msg.sender, "test data", msg.sender, s_sourceFeeToken, 1000);

    bytes memory rawMessage = abi.encode(message);
    uint256 verifierIndex = 0;

    vm.stopPrank();
    vm.prank(s_ccvProxy);
    vm.expectRevert(BaseOnRamp.MustBeCalledByCCVProxy.selector);
    commitOnRampWithoutCCVProxy.forwardToVerifier(rawMessage, verifierIndex);
  }
}
