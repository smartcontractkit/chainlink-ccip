// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Executor} from "../../../executor/Executor.sol";
import {BaseTest} from "../../BaseTest.t.sol";

import {BurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/BurnMintERC20.sol";

contract ExecutorSetup is BaseTest {
  address internal constant INITIAL_CCV = address(121212);
  address internal constant FEE_AGGREGATOR = address(999999);
  uint8 internal constant INITIAL_MAX_CCVS = 1;
  uint16 internal constant DEFAULT_EXEC_FEE_USD_CENTS = 89;
  uint32 internal constant DEFAULT_EXEC_GAS = 150_000;
  uint16 internal constant MIN_BLOCK_CONFIRMATIONS = 50;

  uint8 internal constant EVM_ADDRESS_LENGTH = 20;

  Executor internal s_executor;
  address internal s_sourceFeeToken;

  function setUp() public virtual override {
    super.setUp();

    Executor.DynamicConfig memory dynamicConfig = Executor.DynamicConfig({
      feeAggregator: FEE_AGGREGATOR,
      minBlockConfirmations: MIN_BLOCK_CONFIRMATIONS,
      ccvAllowlistEnabled: true
    });

    s_executor = new Executor(INITIAL_MAX_CCVS, dynamicConfig);
    s_sourceFeeToken = address(new BurnMintERC20("test", "test", 18, 0, 0));

    address[] memory ccvs = new address[](1);
    ccvs[0] = INITIAL_CCV;
    s_executor.applyAllowedCCVUpdates(new address[](0), ccvs, true);

    Executor.RemoteChainConfigArgs[] memory remoteChains = new Executor.RemoteChainConfigArgs[](1);
    remoteChains[0].destChainSelector = DEST_CHAIN_SELECTOR;
    remoteChains[0].config = Executor.RemoteChainConfig({
      usdCentsFee: DEFAULT_EXEC_FEE_USD_CENTS,
      baseExecGas: DEFAULT_EXEC_GAS,
      destAddressLengthBytes: EVM_ADDRESS_LENGTH,
      enabled: true
    });

    s_executor.applyDestChainUpdates(new uint64[](0), remoteChains);
  }
}
