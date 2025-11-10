// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV1} from "../../../interfaces/IPool.sol";
import {IPoolV2} from "../../../interfaces/IPoolV2.sol";

import {Client} from "../../../libraries/Client.sol";
import {MessageV1Codec} from "../../../libraries/MessageV1Codec.sol";
import {Pool} from "../../../libraries/Pool.sol";
import {OnRamp} from "../../../onRamp/OnRamp.sol";
import {OnRampHelper} from "../../helpers/OnRampHelper.sol";
import {OnRampSetup} from "./OnRampSetup.t.sol";

import {IERC165} from "@openzeppelin/contracts@5.0.2/utils/introspection/IERC165.sol";

contract OnRamp_lockOrBurnSingleToken is OnRampSetup {
  OnRampHelper internal s_onRampHelper;
  address internal s_sourceToken;
  address internal s_destToken;
  address internal s_pool;

  function setUp() public override {
    super.setUp();

    s_onRampHelper = new OnRampHelper(
      OnRamp.StaticConfig({
        chainSelector: SOURCE_CHAIN_SELECTOR,
        rmnRemote: s_mockRMNRemote,
        tokenAdminRegistry: address(s_tokenAdminRegistry)
      }),
      OnRamp.DynamicConfig({
        feeQuoter: address(s_feeQuoter),
        reentrancyGuardEntered: false,
        feeAggregator: FEE_AGGREGATOR
      })
    );
    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = s_defaultCCV;

    OnRamp.DestChainConfigArgs[] memory args = new OnRamp.DestChainConfigArgs[](1);
    args[0] = OnRamp.DestChainConfigArgs({
      destChainSelector: DEST_CHAIN_SELECTOR,
      router: s_sourceRouter,
      addressBytesLength: EVM_ADDRESS_LENGTH,
      baseExecutionGasCost: BASE_EXEC_GAS_COST,
      laneMandatedCCVs: new address[](0),
      defaultCCVs: defaultCCVs,
      defaultExecutor: s_defaultExecutor,
      offRamp: abi.encodePacked(address(s_offRampOnRemoteChain))
    });
    s_onRampHelper.applyDestChainConfigUpdates(args);

    s_pool = makeAddr("sourcePool");
    s_sourceToken = makeAddr("sourceToken");
    s_destToken = makeAddr("destToken");

    vm.mockCall(s_pool, abi.encodeCall(IERC165(s_pool).supportsInterface, (Pool.CCIP_POOL_V1)), abi.encode(true));
    vm.mockCall(
      address(s_tokenAdminRegistry), abi.encodeCall(s_tokenAdminRegistry.getPool, (s_sourceToken)), abi.encode(s_pool)
    );
  }

  function test_lockOrBurnSingleToken_CallPoolV2_UsesPoolV2Output() public {
    // setup the pool to support IPoolV2 interface.
    vm.mockCall(s_pool, abi.encodeCall(IERC165.supportsInterface, (type(IPoolV2).interfaceId)), abi.encode(true));

    // mock lockOrBurn v2 call.
    Client.EVMTokenAmount memory tokenAndAmount = Client.EVMTokenAmount({token: s_sourceToken, amount: 123 ether});
    bytes memory receiver = abi.encodePacked(makeAddr("receiver"));
    address originalSender = makeAddr("sender");
    uint16 finality = 5;
    bytes memory tokenArgs = abi.encode("tokenArgs");

    Pool.LockOrBurnInV1 memory expectedInput = Pool.LockOrBurnInV1({
      receiver: receiver,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: originalSender,
      amount: tokenAndAmount.amount,
      localToken: tokenAndAmount.token
    });

    Pool.LockOrBurnOutV1 memory returnData =
      Pool.LockOrBurnOutV1({destTokenAddress: abi.encode(s_destToken), destPoolData: abi.encode("poolData")});
    uint256 destAmount = 120 ether;

    bytes memory callData = abi.encodeWithSelector(IPoolV2.lockOrBurn.selector, expectedInput, finality, tokenArgs);
    vm.expectCall(address(s_pool), callData);
    vm.mockCall(address(s_pool), callData, abi.encode(returnData, destAmount));

    MessageV1Codec.TokenTransferV1 memory transfer = s_onRampHelper.lockOrBurnSingleToken(
      tokenAndAmount, DEST_CHAIN_SELECTOR, receiver, originalSender, finality, tokenArgs
    );

    assertEq(transfer.amount, destAmount);
    assertEq(transfer.sourcePoolAddress, abi.encodePacked(address(s_pool)));
    assertEq(transfer.sourceTokenAddress, abi.encodePacked(s_sourceToken));
    assertEq(transfer.destTokenAddress, abi.encodePacked(s_destToken));
    assertEq(transfer.tokenReceiver, receiver);
    assertEq(transfer.extraData, abi.encode("poolData"));
  }

  function test_lockOrBurnSingleToken_CallPoolV1() public {
    vm.mockCall(s_pool, abi.encodeCall(IERC165.supportsInterface, (type(IPoolV2).interfaceId)), abi.encode(false));
    // mock lockOrBurn v1 call.
    Client.EVMTokenAmount memory tokenAndAmount = Client.EVMTokenAmount({token: s_sourceToken, amount: 123 ether});
    bytes memory receiver = abi.encodePacked(makeAddr("receiver"));
    address originalSender = makeAddr("sender");
    uint16 finality = 0;
    bytes memory tokenArgs = "";

    Pool.LockOrBurnInV1 memory expectedInput = Pool.LockOrBurnInV1({
      receiver: receiver,
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: originalSender,
      amount: tokenAndAmount.amount,
      localToken: tokenAndAmount.token
    });

    Pool.LockOrBurnOutV1 memory returnData =
      Pool.LockOrBurnOutV1({destTokenAddress: abi.encodePacked(s_destToken), destPoolData: abi.encode("poolData")});

    bytes memory callData = abi.encodeWithSelector(IPoolV1.lockOrBurn.selector, expectedInput);
    vm.expectCall(address(s_pool), callData);
    vm.mockCall(address(s_pool), callData, abi.encode(returnData));

    MessageV1Codec.TokenTransferV1 memory transfer = s_onRampHelper.lockOrBurnSingleToken(
      tokenAndAmount, DEST_CHAIN_SELECTOR, receiver, originalSender, finality, tokenArgs
    );

    assertEq(transfer.amount, expectedInput.amount);
    assertEq(transfer.sourcePoolAddress, abi.encodePacked(address(s_pool)));
    assertEq(transfer.sourceTokenAddress, abi.encodePacked(s_sourceToken));
    assertEq(transfer.destTokenAddress, abi.encodePacked(s_destToken));
    assertEq(transfer.tokenReceiver, receiver);
    assertEq(transfer.extraData, abi.encode("poolData"));
  }

  function test_lockOrBurnSingleToken_RevertWhen_CustomBlockConfirmationNotSupportedOnPoolV1() public {
    vm.mockCall(s_pool, abi.encodeCall(IERC165.supportsInterface, (type(IPoolV2).interfaceId)), abi.encode(false));
    // mock lockOrBurn v1 call.
    Client.EVMTokenAmount memory tokenAndAmount = Client.EVMTokenAmount({token: s_sourceToken, amount: 123 ether});
    bytes memory receiver = abi.encodePacked(makeAddr("receiver"));
    address originalSender = makeAddr("sender");
    uint16 finality = 5;
    bytes memory tokenArgs = "";

    vm.expectRevert(OnRamp.CustomBlockConfirmationNotSupportedOnPoolV1.selector);
    s_onRampHelper.lockOrBurnSingleToken(
      tokenAndAmount, DEST_CHAIN_SELECTOR, receiver, originalSender, finality, tokenArgs
    );
  }

  function test_lockOrBurnSingleToken_RevertWhen_TokenArgsNotSupportedOnPoolV1() public {
    vm.mockCall(s_pool, abi.encodeCall(IERC165.supportsInterface, (type(IPoolV2).interfaceId)), abi.encode(false));
    // mock lockOrBurn v1 call.
    Client.EVMTokenAmount memory tokenAndAmount = Client.EVMTokenAmount({token: s_sourceToken, amount: 123 ether});
    bytes memory receiver = abi.encodePacked(makeAddr("receiver"));
    address originalSender = makeAddr("sender");
    uint16 finality = 0;
    bytes memory tokenArgs = abi.encode("tokenArgs");

    vm.expectRevert(OnRamp.TokenArgsNotSupportedOnPoolV1.selector);
    s_onRampHelper.lockOrBurnSingleToken(
      tokenAndAmount, DEST_CHAIN_SELECTOR, receiver, originalSender, finality, tokenArgs
    );
  }
}
