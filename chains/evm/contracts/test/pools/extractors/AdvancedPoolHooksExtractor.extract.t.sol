// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAdvancedPoolHooks} from "../../../interfaces/IAdvancedPoolHooks.sol";
import {IPolicyEngine} from "@chainlink/policy-management/interfaces/IPolicyEngine.sol";

import {Pool} from "../../../libraries/Pool.sol";
import {AdvancedPoolHooksExtractorSetup} from "./AdvancedPoolHooksExtractorSetup.t.sol";

contract AdvancedPoolHooksExtractor_extract is AdvancedPoolHooksExtractorSetup {
  function test_extract_PreflightCheck() public view {
    IPolicyEngine.Payload memory payload = _createPreflightPayload();

    IPolicyEngine.Parameter[] memory params = s_extractor.extract(payload);

    assertEq(7, params.length);

    assertEq(s_extractor.PARAM_FROM(), params[0].name);
    assertEq(s_originalSender, abi.decode(params[0].value, (address)));

    assertEq(s_extractor.PARAM_TO(), params[1].name);
    assertEq(abi.encode(s_receiver), params[1].value);

    assertEq(s_extractor.PARAM_AMOUNT(), params[2].name);
    assertEq(AMOUNT, abi.decode(params[2].value, (uint256)));

    assertEq(s_extractor.PARAM_AMOUNT_POST_FEE(), params[3].name);
    assertEq(AMOUNT_POST_FEE, abi.decode(params[3].value, (uint256)));

    assertEq(s_extractor.PARAM_REMOTE_CHAIN_SELECTOR(), params[4].name);
    assertEq(REMOTE_CHAIN_SELECTOR, abi.decode(params[4].value, (uint64)));

    assertEq(s_extractor.PARAM_TOKEN(), params[5].name);
    assertEq(s_localToken, abi.decode(params[5].value, (address)));

    assertEq(s_extractor.PARAM_BLOCK_CONFIRMATION_REQUESTED(), params[6].name);
    assertEq(BLOCK_CONFIRMATION_REQUESTED, abi.decode(params[6].value, (uint16)));
  }

  function test_extract_PostflightCheck() public view {
    IPolicyEngine.Payload memory payload = _createPostflightPayload();

    IPolicyEngine.Parameter[] memory params = s_extractor.extract(payload);

    assertEq(9, params.length);

    assertEq(s_extractor.PARAM_FROM(), params[0].name);
    assertEq(abi.encode(s_originalSender), params[0].value);

    assertEq(s_extractor.PARAM_TO(), params[1].name);
    assertEq(s_receiver, abi.decode(params[1].value, (address)));

    assertEq(s_extractor.PARAM_AMOUNT(), params[2].name);
    assertEq(AMOUNT, abi.decode(params[2].value, (uint256)));

    assertEq(s_extractor.PARAM_REMOTE_CHAIN_SELECTOR(), params[3].name);
    assertEq(REMOTE_CHAIN_SELECTOR, abi.decode(params[3].value, (uint64)));

    assertEq(s_extractor.PARAM_TOKEN(), params[4].name);
    assertEq(s_localToken, abi.decode(params[4].value, (address)));

    assertEq(s_extractor.PARAM_BLOCK_CONFIRMATION_REQUESTED(), params[5].name);
    assertEq(BLOCK_CONFIRMATION_REQUESTED, abi.decode(params[5].value, (uint16)));

    assertEq(s_extractor.PARAM_SOURCE_POOL_ADDRESS(), params[6].name);
    assertEq(abi.encode(s_sourcePool), params[6].value);

    assertEq(s_extractor.PARAM_SOURCE_POOL_DATA(), params[7].name);
    assertEq(abi.encode("pool data"), params[7].value);

    assertEq(s_extractor.PARAM_SOURCE_DENOMINATED_AMOUNT(), params[8].name);
    assertEq(SOURCE_DENOMINATED_AMOUNT, abi.decode(params[8].value, (uint256)));
  }

  function testFuzz_extract_PreflightCheck(
    address originalSender,
    address receiver,
    address localToken,
    uint256 amount,
    uint256 amountPostFee,
    uint64 remoteChainSelector,
    uint16 blockConfirmationRequested
  ) public view {
    Pool.LockOrBurnInV1 memory lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: abi.encode(receiver),
      remoteChainSelector: remoteChainSelector,
      originalSender: originalSender,
      amount: amount,
      localToken: localToken
    });

    bytes memory tokenArgs = "";

    IPolicyEngine.Payload memory payload = IPolicyEngine.Payload({
      selector: IAdvancedPoolHooks.preflightCheck.selector,
      sender: s_sender,
      data: abi.encode(lockOrBurnIn, blockConfirmationRequested, tokenArgs, amountPostFee),
      context: tokenArgs
    });

    IPolicyEngine.Parameter[] memory params = s_extractor.extract(payload);

    assertEq(7, params.length);
    assertEq(originalSender, abi.decode(params[0].value, (address)));
    assertEq(abi.encode(receiver), params[1].value);
    assertEq(amount, abi.decode(params[2].value, (uint256)));
    assertEq(amountPostFee, abi.decode(params[3].value, (uint256)));
    assertEq(remoteChainSelector, abi.decode(params[4].value, (uint64)));
    assertEq(localToken, abi.decode(params[5].value, (address)));
    assertEq(blockConfirmationRequested, abi.decode(params[6].value, (uint16)));
  }

  function testFuzz_extract_PostflightCheck(
    address originalSender,
    address receiver,
    address localToken,
    address sourcePool,
    uint256 localAmount,
    uint256 sourceDenominatedAmount,
    uint64 remoteChainSelector,
    uint16 blockConfirmationRequested
  ) public view {
    Pool.ReleaseOrMintInV1 memory releaseOrMintIn = Pool.ReleaseOrMintInV1({
      originalSender: abi.encode(originalSender),
      remoteChainSelector: remoteChainSelector,
      receiver: receiver,
      sourceDenominatedAmount: sourceDenominatedAmount,
      localToken: localToken,
      sourcePoolAddress: abi.encode(sourcePool),
      sourcePoolData: abi.encode("pool data"),
      offchainTokenData: ""
    });

    IPolicyEngine.Payload memory payload = IPolicyEngine.Payload({
      selector: IAdvancedPoolHooks.postflightCheck.selector,
      sender: s_sender,
      data: abi.encode(releaseOrMintIn, localAmount, blockConfirmationRequested),
      context: ""
    });

    IPolicyEngine.Parameter[] memory params = s_extractor.extract(payload);

    assertEq(9, params.length);
    assertEq(abi.encode(originalSender), params[0].value);
    assertEq(receiver, abi.decode(params[1].value, (address)));
    assertEq(localAmount, abi.decode(params[2].value, (uint256)));
    assertEq(remoteChainSelector, abi.decode(params[3].value, (uint64)));
    assertEq(localToken, abi.decode(params[4].value, (address)));
    assertEq(blockConfirmationRequested, abi.decode(params[5].value, (uint16)));
    assertEq(abi.encode(sourcePool), params[6].value);
    assertEq(abi.encode("pool data"), params[7].value);
    assertEq(sourceDenominatedAmount, abi.decode(params[8].value, (uint256)));
  }

  // Reverts

  function test_extract_RevertWhen_UnsupportedSelector() public {
    bytes4 unsupportedSelector = bytes4(keccak256("unsupported()"));

    IPolicyEngine.Payload memory payload =
      IPolicyEngine.Payload({selector: unsupportedSelector, sender: s_sender, data: "", context: ""});

    vm.expectRevert(abi.encodeWithSelector(IPolicyEngine.UnsupportedSelector.selector, unsupportedSelector));
    s_extractor.extract(payload);
  }
}
