// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdvancedPoolHooks} from "../../../interfaces/IAdvancedPoolHooks.sol";
import {IPolicyEngine} from "../../../interfaces/IPolicyEngine.sol";

import {CCIPPolicyEnginePayloads} from "../../../libraries/CCIPPolicyEnginePayloads.sol";
import {Pool} from "../../../libraries/Pool.sol";
import {BytesTestHelper} from "../../helpers/BytesTestHelper.sol";
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
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _createReleaseOrMintIn();
    uint256 localAmount = 100e18;
    uint16 blockConfirmationRequested = 5;

    s_advancedPoolHooks.postFlightCheck(releaseOrMintIn, localAmount, blockConfirmationRequested);

    // Verify policy engine was called with correct payload
    IPolicyEngine.Payload memory lastPayload = s_mockPolicyEngine.getLastPayload();
    assertEq(IAdvancedPoolHooks.postFlightCheck.selector, lastPayload.selector);
    assertEq(OWNER, lastPayload.sender);
    assertEq("", lastPayload.context);

    // Verify tag prefix
    assertEq(CCIPPolicyEnginePayloads.POOL_HOOK_INBOUND_POLICY_DATA_V1_TAG, bytes4(lastPayload.data));

    // Decode and verify the payload data
    CCIPPolicyEnginePayloads.PoolHookInboundPolicyDataV1 memory decoded =
      abi.decode(BytesTestHelper._slice(lastPayload.data, 4), (CCIPPolicyEnginePayloads.PoolHookInboundPolicyDataV1));

    assertEq(releaseOrMintIn.originalSender, decoded.originalSender);
    assertEq(blockConfirmationRequested, decoded.blockConfirmationRequested);
    assertEq(releaseOrMintIn.remoteChainSelector, decoded.remoteChainSelector);
    assertEq(releaseOrMintIn.receiver, decoded.receiver);
    assertEq(releaseOrMintIn.sourceDenominatedAmount, decoded.amount);
    assertEq(releaseOrMintIn.localToken, decoded.localToken);
    assertEq(releaseOrMintIn.sourcePoolAddress, decoded.sourcePoolAddress);
    assertEq(releaseOrMintIn.sourcePoolData, decoded.sourcePoolData);
    assertEq(releaseOrMintIn.offchainTokenData, decoded.offchainTokenData);
    assertEq(localAmount, decoded.localAmount);
  }

  function test_postFlightCheck_WithoutPolicyEngine() public {
    assertEq(address(0), s_advancedPoolHooks.getPolicyEngine());

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _createReleaseOrMintIn();

    // Should not revert when policy engine is not set
    s_advancedPoolHooks.postFlightCheck(releaseOrMintIn, 100e18, 5);
  }

  // Reverts

  function test_postFlightCheck_RevertWhen_PolicyEngineRejects() public {
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));
    s_mockPolicyEngine.setShouldRevert(true, "Policy rejected");

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _createReleaseOrMintIn();

    vm.expectRevert(abi.encodeWithSelector(MockPolicyEngine.MockPolicyEngineRejection.selector, "Policy rejected"));
    s_advancedPoolHooks.postFlightCheck(releaseOrMintIn, 100e18, 5);
  }
}
