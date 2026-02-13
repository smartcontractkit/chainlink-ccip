// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdvancedPoolHooks} from "../../../interfaces/IAdvancedPoolHooks.sol";
import {IPolicyEngine} from "@chainlink/ace/policy-management/interfaces/IPolicyEngine.sol";

import {Pool} from "../../../libraries/Pool.sol";
import {MockPolicyEngine} from "../../mocks/MockPolicyEngine.sol";
import {AdvancedPoolHooksSetup} from "./AdvancedPoolHooksSetup.t.sol";
import {AuthorizedCallers} from "@chainlink/contracts/src/v0.8/shared/access/AuthorizedCallers.sol";

contract AdvancedPoolHooks_postflightCheck is AdvancedPoolHooksSetup {
  MockPolicyEngine internal s_mockPolicyEngine;

  address internal s_unauthorizedCaller = makeAddr("unauthorizedCaller");

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

  function test_postflightCheck_WithPolicyEngine() public {
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _createReleaseOrMintIn();
    uint256 localAmount = 100e18;
    uint16 blockConfirmationRequested = 5;
    releaseOrMintIn.sourcePoolData = abi.encode("custom source pool data");
    releaseOrMintIn.offchainTokenData = abi.encode("custom offchain token data");

    s_advancedPoolHooks.postflightCheck(releaseOrMintIn, localAmount, blockConfirmationRequested);

    IPolicyEngine.Payload memory lastPayload = s_mockPolicyEngine.getLastPayload();
    assertEq(IAdvancedPoolHooks.postflightCheck.selector, lastPayload.selector);
    assertEq(OWNER, lastPayload.sender);
    assertEq(releaseOrMintIn.offchainTokenData, lastPayload.context);
    assertEq(abi.encode(releaseOrMintIn, localAmount, blockConfirmationRequested), lastPayload.data);
  }

  function test_postflightCheck_WithoutPolicyEngine() public {
    assertEq(address(0), s_advancedPoolHooks.getPolicyEngine());

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _createReleaseOrMintIn();

    s_advancedPoolHooks.postflightCheck(releaseOrMintIn, 100e18, 5);
  }

  function test_postflightCheck_RevertWhen_PolicyEngineRejects() public {
    string memory expectedRevertReason = "policy rejected";
    s_advancedPoolHooks.setPolicyEngine(address(s_mockPolicyEngine));
    s_mockPolicyEngine.setShouldRevert(true, expectedRevertReason);

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _createReleaseOrMintIn();

    vm.expectRevert(abi.encodeWithSelector(MockPolicyEngine.MockPolicyEngineRejection.selector, expectedRevertReason));
    s_advancedPoolHooks.postflightCheck(releaseOrMintIn, 100e18, 5);
  }

  function test_postflightCheck_OnlyAuthorizedCallersCanInvoke() public {
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _createReleaseOrMintIn();

    s_advancedPoolHooks.postflightCheck(releaseOrMintIn, 100e18, 5);
  }

  function test_postflightCheck_RevertWhen_UnauthorizedCaller() public {
    vm.stopPrank();

    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = _createReleaseOrMintIn();

    vm.prank(s_unauthorizedCaller);
    vm.expectRevert(abi.encodeWithSelector(AuthorizedCallers.UnauthorizedCaller.selector, s_unauthorizedCaller));
    s_advancedPoolHooks.postflightCheck(releaseOrMintIn, 100e18, 5);
  }
}
