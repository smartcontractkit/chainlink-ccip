// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../../libraries/Internal.sol";
import {BaseOnRamp} from "../../../onRamp/BaseOnRamp.sol";
import {CommitOnRamp} from "../../../onRamp/CommitOnRamp.sol";
import {CommitOnRampSetup} from "./CommitOnRampSetup.t.sol";

contract CommitOnRamp_forwardToVerifier is CommitOnRampSetup {
  function setUp() public override {
    super.setUp();
    vm.stopPrank();
  }

  function test_forwardToVerifier() public {
    bytes memory testData = "test data";
    uint256 feeTokenAmount = 1000;
    uint64 expectedNonce = 1;
    uint256 verifierIndex = 0;

    Internal.EVM2AnyVerifierMessage memory message = _createEVM2AnyVerifierMessage(
      DEST_CHAIN_SELECTOR, msg.sender, testData, msg.sender, s_sourceFeeToken, feeTokenAmount
    );

    // Set up mocks for ordered execution
    _setupForwardToVerifierMocks(
      false, DEST_CHAIN_SELECTOR, s_sourceFeeToken, feeTokenAmount, msg.sender, expectedNonce
    );

    vm.prank(s_ccvProxy);
    bytes memory result = s_commitOnRamp.forwardToVerifier(abi.encode(message), verifierIndex);

    uint64 nonce = abi.decode(result, (uint64));
    assertEq(nonce, expectedNonce);
  }

  function test_forwardToVerifier_WithOutOfOrderExecution() public {
    bytes memory testData = "test data";
    uint256 feeTokenAmount = 1000;
    uint64 expectedNonce = 0; // Out of order execution returns 0
    uint256 verifierIndex = 0;

    Internal.EVM2AnyVerifierMessage memory message = _createEVM2AnyVerifierMessage(
      DEST_CHAIN_SELECTOR, msg.sender, testData, msg.sender, s_sourceFeeToken, feeTokenAmount
    );

    // Set up mocks for out-of-order execution
    _setupForwardToVerifierMocks(true, DEST_CHAIN_SELECTOR, s_sourceFeeToken, feeTokenAmount, msg.sender, expectedNonce);

    vm.prank(s_ccvProxy);
    bytes memory result = s_commitOnRamp.forwardToVerifier(abi.encode(message), verifierIndex);

    uint64 nonce = abi.decode(result, (uint64));
    assertEq(nonce, expectedNonce); // Should return 0 for out of order execution
  }

  function test_forwardToVerifier_RevertWhen_MustBeCalledByCCVProxy() public {
    bytes memory testData = "test data";
    uint256 feeTokenAmount = 1000;
    uint256 verifierIndex = 0;

    Internal.EVM2AnyVerifierMessage memory message = _createEVM2AnyVerifierMessage(
      DEST_CHAIN_SELECTOR, msg.sender, testData, msg.sender, s_sourceFeeToken, feeTokenAmount
    );

    vm.prank(STRANGER);
    vm.expectRevert(BaseOnRamp.MustBeCalledByCCVProxy.selector);
    s_commitOnRamp.forwardToVerifier(abi.encode(message), verifierIndex);
  }

  function test_forwardToVerifier_RevertWhen_MustBeCalledByCCVProxy_CCVProxyNotSet() public {
    CommitOnRamp commitOnRampWithoutCCVProxy = new CommitOnRamp(
      address(s_mockRMNRemote),
      address(s_nonceManager),
      CommitOnRamp.DynamicConfig({
        feeQuoter: address(s_feeQuoter),
        feeAggregator: FEE_AGGREGATOR,
        allowlistAdmin: ALLOWLIST_ADMIN
      })
    );

    bytes memory testData = "test data";
    uint256 feeTokenAmount = 1000;
    uint256 verifierIndex = 0;

    Internal.EVM2AnyVerifierMessage memory message = _createEVM2AnyVerifierMessage(
      DEST_CHAIN_SELECTOR, msg.sender, testData, msg.sender, s_sourceFeeToken, feeTokenAmount
    );

    vm.prank(s_ccvProxy);
    vm.expectRevert(BaseOnRamp.MustBeCalledByCCVProxy.selector);
    commitOnRampWithoutCCVProxy.forwardToVerifier(abi.encode(message), verifierIndex);
  }
}
