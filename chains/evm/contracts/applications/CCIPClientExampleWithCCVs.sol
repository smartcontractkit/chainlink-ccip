// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import {IRouterClient} from "../interfaces/IRouterClient.sol";

import {CCIPClientExample} from "./CCIPClientExample.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";

/// @notice Example of a client that supports Cross Chain Verifiers (CCVs).
/// @dev Each source chain can define its own CCV configuration, meaning that incoming traffic
/// from different chains can have different security requirements. If no CCV configuration
/// is defined for a source chain, the default CCV configuration for the lane will be used.
contract CCIPClientExampleWithCCVs is CCIPClientExample {
  error DuplicateCCV(uint64 sourceChainSelector, address ccv);
  error InvalidOptionalThreshold(uint64 sourceChainSelector, uint8 optionalThreshold);
  error NoCCVsProvided(uint64 sourceChainSelector);
  error ZeroAddressNotAllowed();

  event CCVConfigSet(
    uint64 indexed sourceChainSelector, address[] requiredCCVs, address[] optionalCCVs, uint8 optionalThreshold
  );
  event CCVConfigRemoved(uint64 indexed sourceChainSelector);

  /// @notice CCV configuration for a source chain.
  /// @dev For incoming messages, this receiver will require this CCV criteria to be met.
  /// Required CCVs must all pass verification. >= optionalThreshold of the optional CCVs must pass verification.
  struct CCVConfig {
    address[] requiredCCVs;
    address[] optionalCCVs;
    uint8 optionalThreshold;
  }

  /// @notice Arguments required to add a CCV configuration for a source chain.
  struct CCVConfigArgs {
    address[] requiredCCVs;
    address[] optionalCCVs;
    uint64 sourceChainSelector;
    uint8 optionalThreshold;
  }

  /// @notice CCV configurations by source chain selector.
  mapping(uint64 sourceChainSelector => CCVConfig ccvConfig) internal s_ccvConfigs;

  constructor(IRouterClient router, IERC20 feeToken) CCIPClientExample(router, feeToken) {}

  /// @notice Set or remove CCV configurations for source chains.
  /// @param sourceChainSelectorsToRemove List of source chain selectors for which CCV configs will be removed.
  /// @param ccvConfigsToSet List of CCV configs to set.
  function applyCCVConfigUpdates(
    uint64[] calldata sourceChainSelectorsToRemove,
    CCVConfigArgs[] calldata ccvConfigsToSet
  ) external onlyOwner {
    for (uint256 i = 0; i < sourceChainSelectorsToRemove.length; ++i) {
      delete s_ccvConfigs[sourceChainSelectorsToRemove[i]];
      emit CCVConfigRemoved(sourceChainSelectorsToRemove[i]);
    }

    for (uint256 i = 0; i < ccvConfigsToSet.length; ++i) {
      CCVConfigArgs memory args = ccvConfigsToSet[i];
      if (args.requiredCCVs.length == 0 && args.optionalCCVs.length == 0) {
        revert NoCCVsProvided(args.sourceChainSelector);
      }
      if (args.optionalThreshold > args.optionalCCVs.length) {
        revert InvalidOptionalThreshold(args.sourceChainSelector, args.optionalThreshold);
      }
      address[] memory allCCVs = new address[](args.requiredCCVs.length + args.optionalCCVs.length);
      uint256 idx = 0;
      for (uint256 j = 0; j < args.requiredCCVs.length; ++j) {
        allCCVs[idx++] = args.requiredCCVs[j];
      }
      for (uint256 j = 0; j < args.optionalCCVs.length; ++j) {
        allCCVs[idx++] = args.optionalCCVs[j];
      }
      for (uint256 j = 0; j < allCCVs.length; ++j) {
        address ccv = allCCVs[j];
        if (ccv == address(0)) revert ZeroAddressNotAllowed();
        for (uint256 k = 0; k < j; ++k) {
          if (allCCVs[k] == ccv) revert DuplicateCCV(args.sourceChainSelector, ccv);
        }
      }
      s_ccvConfigs[args.sourceChainSelector] = CCVConfig({
        requiredCCVs: args.requiredCCVs,
        optionalCCVs: args.optionalCCVs,
        optionalThreshold: args.optionalThreshold
      });
      emit CCVConfigSet(args.sourceChainSelector, args.requiredCCVs, args.optionalCCVs, args.optionalThreshold);
    }
  }

  /// @notice Provides the required and optional CCVs for a source chain.
  /// @dev OffRamp will apply the defaults for the lane if no CCVs are defined for a source chain.
  function getCCVs(
    uint64 sourceChainSelector
  )
    external
    view
    override
    returns (address[] memory requiredCCVs, address[] memory optionalCCVs, uint8 optionalThreshold)
  {
    CCVConfig memory config = s_ccvConfigs[sourceChainSelector];
    return (config.requiredCCVs, config.optionalCCVs, config.optionalThreshold);
  }
}
