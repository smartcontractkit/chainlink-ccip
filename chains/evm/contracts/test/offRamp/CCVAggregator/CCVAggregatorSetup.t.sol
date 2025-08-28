// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAny2EVMMessageReceiverV2} from "../../../interfaces/IAny2EVMMessageReceiverV2.sol";
import {IRouter} from "../../../interfaces/IRouter.sol";

import {IERC165} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/introspection/IERC165.sol";

import {CCVAggregator} from "../../../offRamp/CCVAggregator.sol";
import {BaseTest} from "../../BaseTest.t.sol";
import {CCVAggregatorHelper} from "../../helpers/CCVAggregatorHelper.sol";

contract CCVAggregatorSetup is BaseTest {
  CCVAggregatorHelper internal s_agg;
  address internal s_defaultCCV;
  address internal s_tokenAdminRegistry;

  function setUp() public virtual override {
    BaseTest.setUp();

    s_defaultCCV = makeAddr("defaultCCV");
    s_tokenAdminRegistry = makeAddr("tokenAdminRegistry");

    s_agg = new CCVAggregatorHelper(
      CCVAggregator.StaticConfig({
        localChainSelector: DEST_CHAIN_SELECTOR,
        gasForCallExactCheck: GAS_FOR_CALL_EXACT_CHECK,
        rmnRemote: s_mockRMNRemote,
        tokenAdminRegistry: s_tokenAdminRegistry
      })
    );

    // Apply initial source chain configuration
    _applySourceConfig(
      s_sourceRouter, SOURCE_CHAIN_SELECTOR, abi.encode(makeAddr("onRamp")), true, new address[](1), new address[](0)
    );
  }

  function _applySourceConfig(
    IRouter router,
    uint64 sourceChainSelector,
    bytes memory onRamp,
    bool isEnabled,
    address[] memory defaultCCVs,
    address[] memory laneMandatedCCVs
  ) internal {
    defaultCCVs[0] = s_defaultCCV;

    CCVAggregator.SourceChainConfigArgs[] memory updates = new CCVAggregator.SourceChainConfigArgs[](1);
    updates[0] = CCVAggregator.SourceChainConfigArgs({
      router: router,
      sourceChainSelector: sourceChainSelector,
      isEnabled: isEnabled,
      onRamp: onRamp,
      defaultCCV: defaultCCVs,
      laneMandatedCCVs: laneMandatedCCVs
    });
    s_agg.applySourceChainConfigUpdates(updates);
  }

  /// @notice Sets up a receiver address to mock getCCVs responses and interface support
  /// @param receiver The receiver address to set up
  /// @param sourceChainSelector The source chain selector for getCCVs mock
  /// @param requiredCCVs Array of required CCV addresses
  /// @param optionalCCVs Array of optional CCV addresses
  /// @param optionalThreshold Threshold for optional CCVs
  function _setGetCCVsReturnData(
    address receiver,
    uint64 sourceChainSelector,
    address[] memory requiredCCVs,
    address[] memory optionalCCVs,
    uint8 optionalThreshold
  ) internal {
    // If receiver has no code, etch minimal bytecode
    if (receiver.code.length == 0) {
      // Simple bytecode that just returns without reverting
      // PUSH1 0x00 (offset), PUSH1 0x00 (length), RETURN (returns empty data)
      vm.etch(receiver, hex"60006000f3");
    }

    // Mock supportsInterface to return true for IERC165 interface
    vm.mockCall(receiver, abi.encodeCall(IERC165.supportsInterface, (type(IERC165).interfaceId)), abi.encode(true));

    // Mock supportsInterface to return true for IAny2EVMMessageReceiverV2
    vm.mockCall(
      receiver,
      abi.encodeCall(IERC165.supportsInterface, (type(IAny2EVMMessageReceiverV2).interfaceId)),
      abi.encode(true)
    );

    // Mock getCCVs function
    vm.mockCall(
      receiver,
      abi.encodeCall(IAny2EVMMessageReceiverV2.getCCVs, (sourceChainSelector)),
      abi.encode(requiredCCVs, optionalCCVs, optionalThreshold)
    );
  }

  /// @notice Convenience function to set up a receiver that returns empty CCVs (falls back to defaults)
  /// @param receiver The receiver address to set up
  /// @param sourceChainSelector The source chain selector for getCCVs mock
  function _setGetCCVsReturnData(address receiver, uint64 sourceChainSelector) internal {
    _setGetCCVsReturnData(receiver, sourceChainSelector, new address[](0), new address[](0), 0);
  }
}
