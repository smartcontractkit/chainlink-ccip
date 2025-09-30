// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../interfaces/IPoolV2.sol";

import {Client} from "../libraries/Client.sol";
import {Pool} from "../libraries/Pool.sol";
import {TokenPool as TokenPoolV1} from "../pools/TokenPool.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {IERC165} from "@openzeppelin/contracts@5.0.2/utils/introspection/IERC165.sol";

abstract contract TokenPool is IPoolV2, TokenPoolV1 {
  error DuplicateCCV(address ccv);

  event CCVConfigUpdated(uint64 indexed remoteChainSelector, address[] outboundCCVs, address[] inboundCCVs);

  struct CCVConfigArg {
    uint64 remoteChainSelector;
    address[] outboundCCVs;
    address[] inboundCCVs;
  }

  struct CCVConfig {
    address[] outboundCCVs;
    address[] inboundCCVs;
  }

  //TODO define billing based struct and storage layout

  mapping(uint64 remoteChainSelector => CCVConfig) internal s_verifierConfig;

  constructor(
    IERC20 token,
    uint8 localTokenDecimals,
    address[] memory allowlist,
    address rmnProxy,
    address router
  ) TokenPoolV1(token, localTokenDecimals, allowlist, rmnProxy, router) {}

  // ================================================================
  // │                        Lock or Burn                          │
  // ================================================================

  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    bytes calldata // tokenArgs
  ) public virtual override returns (Pool.LockOrBurnOutV1 memory) {
    return super.lockOrBurn(lockOrBurnIn);
  }

  // ================================================================
  // │                          CCV                                  │
  // ================================================================

  /// @notice Updates the CCV configuration for specified remote chains.
  /// If the array includes address(0), it indicates that the default CCV should be used alongside any other specified CCVs.
  function applyCCVConfigUpdates(
    CCVConfigArg[] calldata ccvConfigArgs
  ) external virtual onlyOwner {
    for (uint256 i = 0; i < ccvConfigArgs.length; ++i) {
      uint64 remoteChainSelector = ccvConfigArgs[i].remoteChainSelector;
      address[] calldata outboundCCVs = ccvConfigArgs[i].outboundCCVs;
      address[] calldata inboundCCVs = ccvConfigArgs[i].inboundCCVs;

      // check for duplicates in outbound CCVs.
      _checkNoDuplicateAddresses(outboundCCVs);

      // check for duplicates in inbound CCVs.
      _checkNoDuplicateAddresses(inboundCCVs);

      CCVConfig memory ccvConfig = CCVConfig({outboundCCVs: outboundCCVs, inboundCCVs: inboundCCVs});
      emit CCVConfigUpdated(remoteChainSelector, outboundCCVs, inboundCCVs);
      s_verifierConfig[remoteChainSelector] = ccvConfig;
    }
  }

  /// @notice Returns the set of required CCVs for incoming messages from a source chain.
  /// @param sourceChainSelector The source chain selector for incoming messages.
  /// This implementation assumes the same set of CCVs are used for all transfers on a lane.
  /// Implementers can override this function to define custom logic based on these params.
  /// @return requiredCCVs Set of required CCV addresses.
  function getRequiredInboundCCVs(
    address, // localToken
    uint64 sourceChainSelector,
    uint256, // amount,
    uint16, // finality
    bytes calldata // sourcePoolData
  ) external view virtual returns (address[] memory requiredCCVs) {
    return s_verifierConfig[sourceChainSelector].inboundCCVs;
  }

  /// @notice Returns the set of required CCVs for outgoing messages to a destination chain.
  /// @param destChainSelector The destination chain selector for outgoing messages.
  /// This implementation assumes the same set of CCVs are used for all transfers on a lane.
  /// Implementers can override this function to define custom logic based on these params.
  /// @return requiredCCVs Set of required CCV addresses.
  function getRequiredOutboundCCVs(
    address, // localToken
    uint64 destChainSelector,
    uint256, // amount
    uint16, // finality
    bytes calldata // tokenArgs
  ) external view virtual returns (address[] memory requiredCCVs) {
    return s_verifierConfig[destChainSelector].outboundCCVs;
  }

  /// @notice Checks a CCV address array for duplicate entries.
  /// @param ccvs The array of CCV addresses to check for duplicates.
  function _checkNoDuplicateAddresses(
    address[] calldata ccvs
  ) private pure {
    for (uint256 i = 0; i < ccvs.length; ++i) {
      for (uint256 j = i + 1; j < ccvs.length; ++j) {
        if (ccvs[i] == ccvs[j]) {
          revert DuplicateCCV(ccvs[i]);
        }
      }
    }
  }

  // ================================================================
  // │                          Fee                                  │
  // ================================================================

  // TODO implement fee logic

  function getFee(
    uint64, // destChainSelector
    Client.EVM2AnyMessage calldata, // message
    uint16, // finality
    bytes calldata // tokenArgs
  ) external view virtual returns (uint256 feeTokenAmount) {
    return 0;
  }

  /// @notice Signals which version of the pool interface is supported.
  function supportsInterface(
    bytes4 interfaceId
  ) public pure virtual override(TokenPoolV1, IERC165) returns (bool) {
    return interfaceId == Pool.CCIP_POOL_V2 || interfaceId == type(IPoolV2).interfaceId
      || super.supportsInterface(interfaceId);
  }
}
