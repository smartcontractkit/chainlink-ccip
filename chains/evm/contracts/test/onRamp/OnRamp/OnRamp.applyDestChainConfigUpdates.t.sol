// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IRouter} from "../../../interfaces/IRouter.sol";
import {OnRamp} from "../../../onRamp/OnRamp.sol";
import {OnRampSetup} from "./OnRampSetup.t.sol";

contract OnRamp_applyDestChainConfigUpdates is OnRampSetup {
  uint64 internal constant NEW_DEST_SELECTOR = uint64(uint256(keccak256("NEW_DEST_SELECTOR")));

  function test_applyDestChainConfigUpdates_SetsConfigAndEmitsEvent() public {
    IRouter router = s_sourceRouter;
    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = makeAddr("defaultCCV1");
    address[] memory laneMandated = new address[](1);
    laneMandated[0] = makeAddr("laneCCV1");
    address defaultExecutor = makeAddr("defaultExecutor");

    OnRamp.DestChainConfigArgs[] memory args = new OnRamp.DestChainConfigArgs[](1);
    args[0] = OnRamp.DestChainConfigArgs({
      destChainSelector: NEW_DEST_SELECTOR,
      router: router,
      addressBytesLength: EVM_ADDRESS_LENGTH,
      messageNetworkFeeUSDCents: MESSAGE_NETWORK_FEE_USD_CENTS,
      tokenNetworkFeeUSDCents: TOKEN_NETWORK_FEE_USD_CENTS,
      tokenReceiverAllowed: false,
      baseExecutionGasCost: BASE_EXEC_GAS_COST,
      defaultCCVs: defaultCCVs,
      laneMandatedCCVs: laneMandated,
      defaultExecutor: defaultExecutor,
      offRamp: abi.encodePacked(address(s_offRampOnRemoteChain))
    });

    vm.expectEmit();
    emit OnRamp.DestChainConfigSet(NEW_DEST_SELECTOR, 0, args[0]);
    s_onRamp.applyDestChainConfigUpdates(args);

    OnRamp.DestChainConfig memory cfg = s_onRamp.getDestChainConfig(NEW_DEST_SELECTOR);
    assertEq(address(cfg.router), address(router));
    assertEq(cfg.defaultExecutor, args[0].defaultExecutor);
    assertEq(cfg.messageNumber, 0);
    assertEq(cfg.addressBytesLength, args[0].addressBytesLength);
    assertEq(cfg.baseExecutionGasCost, args[0].baseExecutionGasCost);
    assertEq(cfg.messageNetworkFeeUSDCents, args[0].messageNetworkFeeUSDCents);
    assertEq(cfg.tokenNetworkFeeUSDCents, args[0].tokenNetworkFeeUSDCents);
    assertEq(cfg.defaultCCVs, defaultCCVs);
    assertEq(cfg.laneMandatedCCVs, laneMandated);
  }

  function test_applyDestChainConfigUpdates_NonEvmAddressLength() public {
    OnRamp.DestChainConfigArgs[] memory args = new OnRamp.DestChainConfigArgs[](1);
    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = makeAddr("defaultCCV");

    args[0] = OnRamp.DestChainConfigArgs({
      destChainSelector: NEW_DEST_SELECTOR + 99,
      router: s_sourceRouter,
      addressBytesLength: NON_EVM_ADDRESS_LENGTH,
      messageNetworkFeeUSDCents: MESSAGE_NETWORK_FEE_USD_CENTS,
      tokenNetworkFeeUSDCents: TOKEN_NETWORK_FEE_USD_CENTS,
      tokenReceiverAllowed: false,
      baseExecutionGasCost: BASE_EXEC_GAS_COST,
      defaultCCVs: defaultCCVs,
      laneMandatedCCVs: new address[](0),
      defaultExecutor: makeAddr("executor"),
      offRamp: abi.encodePacked(bytes32(uint256(1234)))
    });

    s_onRamp.applyDestChainConfigUpdates(args);
    OnRamp.DestChainConfig memory cfg = s_onRamp.getDestChainConfig(NEW_DEST_SELECTOR + 99);
    assertEq(NON_EVM_ADDRESS_LENGTH, cfg.addressBytesLength);
  }

  function test_applyDestChainConfigUpdates_AllowsZeroRouterToPause() public {
    OnRamp.DestChainConfigArgs[] memory args = new OnRamp.DestChainConfigArgs[](1);
    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = makeAddr("defaultCCV");
    args[0] = OnRamp.DestChainConfigArgs({
      destChainSelector: NEW_DEST_SELECTOR + 1,
      router: IRouter(address(0)),
      addressBytesLength: EVM_ADDRESS_LENGTH,
      messageNetworkFeeUSDCents: MESSAGE_NETWORK_FEE_USD_CENTS,
      tokenNetworkFeeUSDCents: TOKEN_NETWORK_FEE_USD_CENTS,
      tokenReceiverAllowed: false,
      baseExecutionGasCost: BASE_EXEC_GAS_COST,
      defaultCCVs: defaultCCVs,
      laneMandatedCCVs: new address[](0),
      defaultExecutor: makeAddr("executor"),
      offRamp: abi.encodePacked(address(s_offRampOnRemoteChain))
    });

    // Should not revert, router can be zero.
    s_onRamp.applyDestChainConfigUpdates(args);
    OnRamp.DestChainConfig memory cfg = s_onRamp.getDestChainConfig(NEW_DEST_SELECTOR + 1);
    assertEq(address(cfg.router), address(0));
    assertEq(EVM_ADDRESS_LENGTH, cfg.addressBytesLength);
  }

  function test_getAllDestChainConfigs_ReturnsMultipleChains() public {
    // Add a second destination chain.
    uint64 chain2 = NEW_DEST_SELECTOR + 100;
    address[] memory defaultCCVs2 = new address[](1);
    defaultCCVs2[0] = makeAddr("defaultCCV2");
    address[] memory laneMandated2 = new address[](1);
    laneMandated2[0] = makeAddr("laneCCV2");
    address defaultExecutor2 = makeAddr("defaultExecutor2");

    OnRamp.DestChainConfigArgs[] memory configs = new OnRamp.DestChainConfigArgs[](1);
    configs[0] = OnRamp.DestChainConfigArgs({
      destChainSelector: chain2,
      router: s_sourceRouter,
      addressBytesLength: EVM_ADDRESS_LENGTH,
      messageNetworkFeeUSDCents: MESSAGE_NETWORK_FEE_USD_CENTS + 1,
      tokenNetworkFeeUSDCents: TOKEN_NETWORK_FEE_USD_CENTS + 1,
      tokenReceiverAllowed: true,
      baseExecutionGasCost: BASE_EXEC_GAS_COST + 1000,
      defaultCCVs: defaultCCVs2,
      laneMandatedCCVs: laneMandated2,
      defaultExecutor: defaultExecutor2,
      offRamp: abi.encodePacked(address(s_offRampOnRemoteChain))
    });

    s_onRamp.applyDestChainConfigUpdates(configs);

    (uint64[] memory selectors, OnRamp.DestChainConfig[] memory chainConfigs) = s_onRamp.getAllDestChainConfigs();

    assertEq(selectors.length, 2);
    assertEq(chainConfigs.length, 2);

    // Check first chain (from setup).
    assertEq(selectors[0], DEST_CHAIN_SELECTOR);
    assertEq(chainConfigs[0].defaultCCVs[0], s_defaultCCV);
    assertEq(chainConfigs[0].defaultExecutor, s_defaultExecutor);

    // Check second chain.
    assertEq(selectors[1], chain2);
    assertEq(chainConfigs[1].defaultCCVs[0], defaultCCVs2[0]);
    assertEq(chainConfigs[1].laneMandatedCCVs[0], laneMandated2[0]);
    assertEq(chainConfigs[1].defaultExecutor, defaultExecutor2);
    assertEq(chainConfigs[1].messageNetworkFeeUSDCents, MESSAGE_NETWORK_FEE_USD_CENTS + 1);
    assertEq(chainConfigs[1].tokenNetworkFeeUSDCents, TOKEN_NETWORK_FEE_USD_CENTS + 1);
    assertEq(chainConfigs[1].baseExecutionGasCost, BASE_EXEC_GAS_COST + 1000);
    assertEq(chainConfigs[1].tokenReceiverAllowed, true);
  }

  function test_applyDestChainConfigUpdates_RevertWhen_InvalidDestChainConfig_ZeroAddressBytesLength() public {
    OnRamp.DestChainConfigArgs[] memory args = new OnRamp.DestChainConfigArgs[](1);
    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = makeAddr("defaultCCV");
    args[0] = OnRamp.DestChainConfigArgs({
      destChainSelector: NEW_DEST_SELECTOR + 5,
      router: s_sourceRouter,
      addressBytesLength: 0,
      messageNetworkFeeUSDCents: MESSAGE_NETWORK_FEE_USD_CENTS,
      tokenNetworkFeeUSDCents: TOKEN_NETWORK_FEE_USD_CENTS,
      tokenReceiverAllowed: false,
      baseExecutionGasCost: BASE_EXEC_GAS_COST,
      defaultCCVs: defaultCCVs,
      laneMandatedCCVs: new address[](0),
      defaultExecutor: makeAddr("executor"),
      offRamp: abi.encodePacked(address(s_offRampOnRemoteChain))
    });

    vm.expectRevert(abi.encodeWithSelector(OnRamp.InvalidDestChainConfig.selector, NEW_DEST_SELECTOR + 5));
    s_onRamp.applyDestChainConfigUpdates(args);
  }

  function test_applyDestChainConfigUpdates_RevertWhen_InvalidDestChainConfig_ZeroSelector() public {
    OnRamp.DestChainConfigArgs[] memory args = new OnRamp.DestChainConfigArgs[](1);
    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = makeAddr("defaultCCV");
    args[0] = OnRamp.DestChainConfigArgs({
      destChainSelector: 0,
      router: s_sourceRouter,
      addressBytesLength: EVM_ADDRESS_LENGTH,
      messageNetworkFeeUSDCents: MESSAGE_NETWORK_FEE_USD_CENTS,
      tokenNetworkFeeUSDCents: TOKEN_NETWORK_FEE_USD_CENTS,
      tokenReceiverAllowed: false,
      baseExecutionGasCost: BASE_EXEC_GAS_COST,
      defaultCCVs: defaultCCVs,
      laneMandatedCCVs: new address[](0),
      defaultExecutor: makeAddr("executor"),
      offRamp: abi.encodePacked(address(s_offRampOnRemoteChain))
    });

    vm.expectRevert(abi.encodeWithSelector(OnRamp.InvalidDestChainConfig.selector, uint64(0)));
    s_onRamp.applyDestChainConfigUpdates(args);
  }

  function test_applyDestChainConfigUpdates_RevertWhen_DefaultExecutorZero() public {
    OnRamp.DestChainConfigArgs[] memory args = new OnRamp.DestChainConfigArgs[](1);
    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = makeAddr("defaultCCV");
    args[0] = OnRamp.DestChainConfigArgs({
      destChainSelector: NEW_DEST_SELECTOR + 8,
      router: s_sourceRouter,
      addressBytesLength: EVM_ADDRESS_LENGTH,
      messageNetworkFeeUSDCents: MESSAGE_NETWORK_FEE_USD_CENTS,
      tokenNetworkFeeUSDCents: TOKEN_NETWORK_FEE_USD_CENTS,
      tokenReceiverAllowed: false,
      baseExecutionGasCost: BASE_EXEC_GAS_COST,
      defaultCCVs: defaultCCVs,
      laneMandatedCCVs: new address[](0),
      defaultExecutor: address(0),
      offRamp: abi.encodePacked(address(s_offRampOnRemoteChain))
    });

    vm.expectRevert(OnRamp.InvalidConfig.selector);
    s_onRamp.applyDestChainConfigUpdates(args);
  }

  function test_applyDestChainConfigUpdates_RevertWhen_DestIsLocalChain() public {
    // Using SOURCE_CHAIN_SELECTOR as local chain selector from setup.
    OnRamp.DestChainConfigArgs[] memory args = new OnRamp.DestChainConfigArgs[](1);
    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = makeAddr("defaultCCV");
    args[0] = OnRamp.DestChainConfigArgs({
      destChainSelector: SOURCE_CHAIN_SELECTOR,
      router: s_sourceRouter,
      addressBytesLength: EVM_ADDRESS_LENGTH,
      messageNetworkFeeUSDCents: MESSAGE_NETWORK_FEE_USD_CENTS,
      tokenNetworkFeeUSDCents: TOKEN_NETWORK_FEE_USD_CENTS,
      tokenReceiverAllowed: false,
      baseExecutionGasCost: BASE_EXEC_GAS_COST,
      defaultCCVs: defaultCCVs,
      laneMandatedCCVs: new address[](0),
      defaultExecutor: makeAddr("executor"),
      offRamp: abi.encodePacked(address(s_offRampOnRemoteChain))
    });

    vm.expectRevert(abi.encodeWithSelector(OnRamp.InvalidDestChainConfig.selector, SOURCE_CHAIN_SELECTOR));
    s_onRamp.applyDestChainConfigUpdates(args);
  }

  function test_applyDestChainConfigUpdates_RevertWhen_InvalidDestChainConfig_ZeroBaseExecutionGasCost() public {
    OnRamp.DestChainConfigArgs[] memory args = new OnRamp.DestChainConfigArgs[](1);
    address[] memory defaultCCVs = new address[](1);
    defaultCCVs[0] = makeAddr("defaultCCV");
    args[0] = OnRamp.DestChainConfigArgs({
      destChainSelector: SOURCE_CHAIN_SELECTOR,
      router: s_sourceRouter,
      addressBytesLength: EVM_ADDRESS_LENGTH,
      messageNetworkFeeUSDCents: MESSAGE_NETWORK_FEE_USD_CENTS,
      tokenNetworkFeeUSDCents: TOKEN_NETWORK_FEE_USD_CENTS,
      tokenReceiverAllowed: false,
      baseExecutionGasCost: 0,
      defaultCCVs: defaultCCVs,
      laneMandatedCCVs: new address[](0),
      defaultExecutor: makeAddr("executor"),
      offRamp: abi.encodePacked(address(s_offRampOnRemoteChain))
    });

    vm.expectRevert(abi.encodeWithSelector(OnRamp.InvalidDestChainConfig.selector, SOURCE_CHAIN_SELECTOR));
    s_onRamp.applyDestChainConfigUpdates(args);
  }
}
