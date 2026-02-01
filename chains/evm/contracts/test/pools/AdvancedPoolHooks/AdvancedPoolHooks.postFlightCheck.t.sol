// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdvancedPoolHooks} from "../../../interfaces/IAdvancedPoolHooks.sol";
import {IPolicyEngine} from "../../../interfaces/IPolicyEngine.sol";

import {Pool} from "../../../libraries/Pool.sol";
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

  function testFuzz_postFlightCheck_WithPolicyEngine(
    bytes memory sourcePoolData,
    bytes memory offchainTokenData
  ) public {
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _createReleaseOrMintIn();
    uint256 localAmount = 100e18;
    uint16 blockConfirmationRequested = 5;
    releaseOrMintIn.sourcePoolData = sourcePoolData;
    releaseOrMintIn.offchainTokenData = offchainTokenData;

    s_advancedPoolHooks.postFlightCheck(releaseOrMintIn, localAmount, blockConfirmationRequested);

    IPolicyEngine.Payload memory lastPayload = s_mockPolicyEngine.getLastPayload();
    assertEq(IAdvancedPoolHooks.postFlightCheck.selector, lastPayload.selector);
    assertEq(OWNER, lastPayload.sender);
    assertEq(releaseOrMintIn.offchainTokenData, lastPayload.context);
    assertEq(abi.encode(releaseOrMintIn, localAmount, blockConfirmationRequested), lastPayload.data);
  }

  function test_postFlightCheck_WithoutPolicyEngine() public {
    assertEq(address(0), s_advancedPoolHooks.getPolicyEngine());

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _createReleaseOrMintIn();

    s_advancedPoolHooks.postFlightCheck(releaseOrMintIn, 100e18, 5);
  }

  function test_postFlightCheck_RevertWhen_PolicyEngineRejects() public {
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));
    s_mockPolicyEngine.setShouldRevert(true, "Policy rejected");

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _createReleaseOrMintIn();

    vm.expectRevert(abi.encodeWithSelector(MockPolicyEngine.MockPolicyEngineRejection.selector, "Policy rejected"));
    s_advancedPoolHooks.postFlightCheck(releaseOrMintIn, 100e18, 5);
  }
}
