// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdvancedPoolHooks} from "../../../interfaces/IAdvancedPoolHooks.sol";
import {IPolicyEngine} from "@chainlink/ace/policy-management/interfaces/IPolicyEngine.sol";

import {Pool} from "../../../libraries/Pool.sol";
import {AdvancedPoolHooksExtractor} from "../../../pools/extractors/AdvancedPoolHooksExtractor.sol";

import {Test} from "forge-std/Test.sol";

contract AdvancedPoolHooksExtractorSetup is Test {
  AdvancedPoolHooksExtractor internal s_extractor;

  address internal s_sender = makeAddr("sender");
  address internal s_originalSender = makeAddr("originalSender");
  address internal s_receiver = makeAddr("receiver");
  address internal s_localToken = makeAddr("localToken");
  address internal s_sourcePool = makeAddr("sourcePool");

  uint64 internal constant REMOTE_CHAIN_SELECTOR = 123;
  uint256 internal constant AMOUNT = 100e18;
  uint256 internal constant AMOUNT_POST_FEE = 99e18;
  uint16 internal constant BLOCK_CONFIRMATION_REQUESTED = 5;
  uint256 internal constant SOURCE_DENOMINATED_AMOUNT = 200e18;

  function setUp() public virtual {
    s_extractor = new AdvancedPoolHooksExtractor();
  }

  function _createPreflightPayload() internal view returns (IPolicyEngine.Payload memory) {
    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: abi.encode(s_receiver),
      remoteChainSelector: REMOTE_CHAIN_SELECTOR,
      originalSender: s_originalSender,
      amount: AMOUNT,
      localToken: s_localToken
    });

    bytes memory tokenArgs = abi.encode("token args");

    return IPolicyEngine.Payload({
      selector: IAdvancedPoolHooks.preflightCheck.selector,
      sender: s_sender,
      data: abi.encode(lockOrBurnIn, BLOCK_CONFIRMATION_REQUESTED, tokenArgs, AMOUNT_POST_FEE),
      context: tokenArgs
    });
  }

  function _createPostflightPayload() internal view returns (IPolicyEngine.Payload memory) {
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      originalSender: abi.encode(s_originalSender),
      remoteChainSelector: REMOTE_CHAIN_SELECTOR,
      receiver: s_receiver,
      sourceDenominatedAmount: SOURCE_DENOMINATED_AMOUNT,
      localToken: s_localToken,
      sourcePoolAddress: abi.encode(s_sourcePool),
      sourcePoolData: abi.encode("pool data"),
      offchainTokenData: abi.encode("offchain data")
    });

    return IPolicyEngine.Payload({
      selector: IAdvancedPoolHooks.postflightCheck.selector,
      sender: s_sender,
      data: abi.encode(releaseOrMintIn, AMOUNT, BLOCK_CONFIRMATION_REQUESTED),
      context: releaseOrMintIn.offchainTokenData
    });
  }
}
