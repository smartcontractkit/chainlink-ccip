// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdvancedPoolHooks} from "../../../interfaces/IAdvancedPoolHooks.sol";
import {IPolicyEngine} from "../../../interfaces/IPolicyEngine.sol";

import {Pool} from "../../../libraries/Pool.sol";
import {AdvancedPoolHooks} from "../../../pools/AdvancedPoolHooks.sol";
import {MockPolicyEngine} from "../../mocks/MockPolicyEngine.sol";
import {AdvancedPoolHooksSetup} from "./AdvancedPoolHooksSetup.t.sol";

contract AdvancedPoolHooks_postFlightCheck is AdvancedPoolHooksSetup {
  MockPolicyEngine internal s_mockPolicyEngine;

  function setUp() public virtual override {
    super.setUp();
    s_mockPolicyEngine = new MockPolicyEngine();
  }

  function _createReleaseOrMintIn() internal view returns (Pool.ReleaseOrMintInV1 memory) {
    return Pool.ReleaseOrMintInV1({
      originalSender: abi.encode(s_sender),
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      receiver: OWNER,
      sourceDenominatedAmount: 100e18,
      localToken: address(s_token),
      sourcePoolAddress: abi.encode(address(s_tokenPool)),
      sourcePoolData: "",
      offchainTokenData: ""
    });
  }

  function test_postFlightCheck_WithPolicyEngine() public {
    // Set up policy engine
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _createReleaseOrMintIn();
    uint256 localAmount = 100e18;
    uint16 blockConfirmationRequested = 5;

    // Call postFlightCheck
    s_advancedPoolHooks.postFlightCheck(releaseOrMintIn, localAmount, blockConfirmationRequested);

    // Verify policy engine was called with correct payload
    IPolicyEngine.Payload memory lastPayload = s_mockPolicyEngine.getLastPayload();
    assertEq(lastPayload.selector, IAdvancedPoolHooks.postFlightCheck.selector);
    assertEq(lastPayload.sender, OWNER); // msg.sender is OWNER from vm.startPrank in setup

    // Verify the encoded data contains the releaseOrMintIn, localAmount, and blockConfirmationRequested
    bytes memory expectedData = abi.encode(releaseOrMintIn, localAmount, blockConfirmationRequested);
    assertEq(lastPayload.data, expectedData);
    assertEq(lastPayload.context, "");
  }

  function test_postFlightCheck_WithoutPolicyEngine() public {
    // Ensure no policy engine is set
    assertEq(s_advancedPoolHooks.getPolicyEngine(), address(0));

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _createReleaseOrMintIn();

    // Should not revert when policy engine is not set
    s_advancedPoolHooks.postFlightCheck(releaseOrMintIn, 100e18, 5);
  }

  function test_postFlightCheck_RevertWhen_PolicyEngineRejects() public {
    // Set up policy engine to reject
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));
    s_mockPolicyEngine.setShouldRevert(true, "Policy rejected");

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _createReleaseOrMintIn();

    vm.expectRevert(abi.encodeWithSelector(MockPolicyEngine.MockPolicyEngineRejection.selector, "Policy rejected"));
    s_advancedPoolHooks.postFlightCheck(releaseOrMintIn, 100e18, 5);
  }
}
