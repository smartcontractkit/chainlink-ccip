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
    Internal.EVM2AnyVerifierMessage memory message =
      _createEVM2AnyVerifierMessage(DEST_CHAIN_SELECTOR, msg.sender, "test data", msg.sender, s_sourceFeeToken, 1000);

    uint256 verifierIndex = 0;

    // Set up mocks for ordered execution
    _setupForwardToVerifierMocks(false, DEST_CHAIN_SELECTOR, s_sourceFeeToken, 1000, msg.sender, 1);

    vm.prank(s_ccvProxy);
    bytes memory result = s_commitOnRamp.forwardToVerifier(abi.encode(message), verifierIndex);

    uint64 nonce = abi.decode(result, (uint64));
    assertEq(nonce, 1);
  }

  function test_forwardToVerifier_WithOutOfOrderExecution() public {
    Internal.EVM2AnyVerifierMessage memory message =
      _createEVM2AnyVerifierMessage(DEST_CHAIN_SELECTOR, msg.sender, "test data", msg.sender, s_sourceFeeToken, 1000);

    uint256 verifierIndex = 0;

    // Set up mocks for out-of-order execution
    _setupForwardToVerifierMocks(true, DEST_CHAIN_SELECTOR, s_sourceFeeToken, 1000, msg.sender, 0);

    vm.prank(s_ccvProxy);
    bytes memory result = s_commitOnRamp.forwardToVerifier(abi.encode(message), verifierIndex);

    uint64 nonce = abi.decode(result, (uint64));
    assertEq(nonce, 0); // Should return 0 for out of order execution
  }

  function test_forwardToVerifier_RevertWhen_MustBeCalledByCCVProxy() public {
    Internal.EVM2AnyVerifierMessage memory message =
      _createEVM2AnyVerifierMessage(DEST_CHAIN_SELECTOR, msg.sender, "test data", msg.sender, s_sourceFeeToken, 1000);

    uint256 verifierIndex = 0;

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

    Internal.EVM2AnyVerifierMessage memory message =
      _createEVM2AnyVerifierMessage(DEST_CHAIN_SELECTOR, msg.sender, "test data", msg.sender, s_sourceFeeToken, 1000);

    uint256 verifierIndex = 0;

    vm.prank(s_ccvProxy);
    vm.expectRevert(BaseOnRamp.MustBeCalledByCCVProxy.selector);
    commitOnRampWithoutCCVProxy.forwardToVerifier(abi.encode(message), verifierIndex);
  }
}
