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
  error ZeroAddressNotAllowedAsOptional();

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

  /// @notice Set CCV configurations for source chains.
  /// @param ccvConfigsToSet List of CCV configs to set.
  function applyCCVConfigUpdates(
    CCVConfigArgs[] calldata ccvConfigsToSet
  ) external virtual onlyOwner {
    for (uint256 i = 0; i < ccvConfigsToSet.length; ++i) {
      CCVConfigArgs memory args = ccvConfigsToSet[i];
      // If optionalThreshold > optionalCCVs.length, then it's impossible to satisfy the optional CCV requirement.
      // If optionalThreshold == optionalCCVs.length, then optional CCVs are essentially required.
      // They should instead be defined as required CCVs.
      if (args.optionalThreshold >= args.optionalCCVs.length) {
        revert InvalidOptionalThreshold(args.sourceChainSelector, args.optionalThreshold);
      }
      uint256 requiredCCVLength = args.requiredCCVs.length;
      uint256 optionalCCVLength = args.optionalCCVs.length;
      uint256 totalCCVLength = requiredCCVLength + optionalCCVLength;
      for (uint256 j = 0; j < totalCCVLength; ++j) {
        address ccvAddressJ = j < requiredCCVLength ? args.requiredCCVs[j] : args.optionalCCVs[j - requiredCCVLength];
        // address(0) is a valid required CCV address, but not a valid optional CCV address.
        // This is because address(0) signals to always enforce the default CCVs for the lane.
        if (j >= requiredCCVLength && ccvAddressJ == address(0)) {
          revert ZeroAddressNotAllowedAsOptional();
        }

        for (uint256 k = j + 1; k < totalCCVLength; ++k) {
          address ccvAddressK = k < requiredCCVLength ? args.requiredCCVs[k] : args.optionalCCVs[k - requiredCCVLength];
          if (ccvAddressK == ccvAddressJ) {
            revert DuplicateCCV(args.sourceChainSelector, ccvAddressK);
          }
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
    virtual
    returns (address[] memory requiredCCVs, address[] memory optionalCCVs, uint8 optionalThreshold)
  {
    CCVConfig memory config = s_ccvConfigs[sourceChainSelector];
    return (config.requiredCCVs, config.optionalCCVs, config.optionalThreshold);
  }
}
