// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPolicyEngine} from "@chainlink/policy-management/interfaces/IPolicyEngine.sol";

import {Pool} from "../../../../libraries/Pool.sol";
import {MultiTokenPoolHooks} from "../../../../pools/MultiTokenPool/MultiTokenPoolHooks.sol";
import {BaseTest} from "../../../BaseTest.t.sol";
import {MockPolicyEngine} from "../../../mocks/MockPolicyEngine.sol";

contract MultiTokenPoolHooksSetup is BaseTest {
  MultiTokenPoolHooks internal s_hooks;
  MockPolicyEngine internal s_policyEngine;
  address internal s_sender = makeAddr("sender");
  address internal s_receiver = makeAddr("receiver");
  Pool.LockOrBurnInV1 internal s_lockOrBurnIn;
  Pool.ReleaseOrMintInV1 internal s_releaseOrMintIn;

  function setUp() public virtual override {
    super.setUp();
    s_policyEngine = new MockPolicyEngine();
    s_hooks = new MultiTokenPoolHooks(new address[](0), IPolicyEngine(address(0)), _arrayOf(OWNER));

    s_lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: abi.encode(s_receiver),
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: s_sender,
      amount: 1e18,
      localToken: makeAddr("token")
    });
    s_releaseOrMintIn = Pool.ReleaseOrMintInV1({
      originalSender: abi.encode(s_sender),
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      receiver: s_receiver,
      sourceDenominatedAmount: 1e18,
      localToken: s_lockOrBurnIn.localToken,
      sourcePoolAddress: abi.encode(makeAddr("sourcePool")),
      sourcePoolData: "",
      offchainTokenData: ""
    });
  }

  /// @dev Deploys hooks that only allow s_sender to send, without a policy engine.
  function _deployHooksWithAllowList() internal returns (MultiTokenPoolHooks) {
    return new MultiTokenPoolHooks(_arrayOf(s_sender), IPolicyEngine(address(0)), _arrayOf(OWNER));
  }
}
