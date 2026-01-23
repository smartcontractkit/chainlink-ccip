// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdvancedPoolHooks} from "../../../interfaces/IAdvancedPoolHooks.sol";
import {IPolicyEngine} from "../../../interfaces/IPolicyEngine.sol";

import {CCIPPolicyEnginePayloads} from "../../../libraries/CCIPPolicyEnginePayloads.sol";
import {Pool} from "../../../libraries/Pool.sol";
import {MockPolicyEngine} from "../../mocks/MockPolicyEngine.sol";
import {AdvancedPoolHooksSetup} from "./AdvancedPoolHooksSetup.t.sol";

contract AdvancedPoolHooks_postFlightCheck is AdvancedPoolHooksSetup {
  // bytes4(keccak256("InboundPolicyDataV1"))
  bytes4 internal constant INBOUND_POLICY_DATA_V1_TAG = 0xe8deab79;

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
    assertEq(lastPayload.selector, IAdvancedPoolHooks.postFlightCheck.selector);
    assertEq(lastPayload.sender, OWNER);
    assertEq(lastPayload.context, "");

    // Verify tag prefix
    bytes4 tag = bytes4(lastPayload.data);
    assertEq(tag, INBOUND_POLICY_DATA_V1_TAG);

    // Decode and verify the payload data
    CCIPPolicyEnginePayloads.InboundPolicyDataV1 memory decoded =
      abi.decode(this._sliceBytes(lastPayload.data, 4), (CCIPPolicyEnginePayloads.InboundPolicyDataV1));

    assertEq(decoded.originalSender, releaseOrMintIn.originalSender);
    assertEq(decoded.remoteChainSelector, releaseOrMintIn.remoteChainSelector);
    assertEq(decoded.receiver, releaseOrMintIn.receiver);
    assertEq(decoded.amount, releaseOrMintIn.sourceDenominatedAmount);
    assertEq(decoded.localToken, releaseOrMintIn.localToken);
    assertEq(decoded.sourcePoolAddress, releaseOrMintIn.sourcePoolAddress);
    assertEq(decoded.sourcePoolData, releaseOrMintIn.sourcePoolData);
    assertEq(decoded.offchainTokenData, releaseOrMintIn.offchainTokenData);
    assertEq(decoded.localAmount, localAmount);
    assertEq(decoded.blockConfirmationRequested, blockConfirmationRequested);
  }

  function test_postFlightCheck_WithoutPolicyEngine() public {
    assertEq(s_advancedPoolHooks.getPolicyEngine(), address(0));

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _createReleaseOrMintIn();

    // Should not revert when policy engine is not set
    s_advancedPoolHooks.postFlightCheck(releaseOrMintIn, 100e18, 5);
  }

  function test_postFlightCheck_RevertWhen_PolicyEngineRejects() public {
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));
    s_mockPolicyEngine.setShouldRevert(true, "Policy rejected");

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _createReleaseOrMintIn();

    vm.expectRevert(abi.encodeWithSelector(MockPolicyEngine.MockPolicyEngineRejection.selector, "Policy rejected"));
    s_advancedPoolHooks.postFlightCheck(releaseOrMintIn, 100e18, 5);
  }

  // Helper to slice bytes, exposed as external for use with this.
  function _sliceBytes(bytes memory data, uint256 start) external pure returns (bytes memory) {
    bytes memory result = new bytes(data.length - start);
    for (uint256 i = 0; i < result.length; i++) {
      result[i] = data[start + i];
    }
    return result;
  }
}
