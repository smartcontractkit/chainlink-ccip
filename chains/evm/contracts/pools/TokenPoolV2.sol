// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IPoolV2} from "../interfaces/IPoolV2.sol";

import {Pool} from "../libraries/Pool.sol";
import {TokenPool} from "./TokenPool.sol";

import {IERC20} from "@openzeppelin/contracts@4.8.3/token/ERC20/IERC20.sol";
import {IERC165} from "@openzeppelin/contracts@5.0.2/utils/introspection/IERC165.sol";

abstract contract TokenPoolV2 is IPoolV2, TokenPool {
  error DuplicateCCV(address ccv);

  event CCVConfigUpdated(uint64 indexed destChainSelector, address[] outboundCCVs, address[] inboundCCVs);

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

  /// @notice Signals which version of the pool interface is supported.
  function supportsInterface(
    bytes4 interfaceId
  ) public pure virtual override(TokenPool, IERC165) returns (bool) {
    return interfaceId == Pool.CCIP_POOL_V2 || interfaceId == type(IPoolV2).interfaceId
      || super.supportsInterface(interfaceId);
  }

  constructor(
    IERC20 token,
    uint8 localTokenDecimals,
    address[] memory allowlist,
    address rmnProxy,
    address router
  ) TokenPool(token, localTokenDecimals, allowlist, rmnProxy, router) {}

  // ================================================================
  // │                        Lock or Burn                          │
  // ================================================================

  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    bytes calldata // tokenArgs
  ) public virtual override returns (Pool.LockOrBurnOutV1 memory) {
    _validateLockOrBurn(lockOrBurnIn);
    _lockOrBurn(lockOrBurnIn.amount);

    emit LockedOrBurned({
      remoteChainSelector: lockOrBurnIn.remoteChainSelector,
      token: address(i_token),
      sender: msg.sender,
      amount: lockOrBurnIn.amount
    });

    return Pool.LockOrBurnOutV1({
      destTokenAddress: getRemoteToken(lockOrBurnIn.remoteChainSelector),
      destPoolData: _encodeLocalDecimals()
    });
  }

  // ================================================================
  // │                          CCV                                  │
  // ================================================================

  function applyCCVConfigUpdates(
    CCVConfigArg[] calldata ccvConfigArgs
  ) external virtual onlyOwner {
    for (uint256 i = 0; i < ccvConfigArgs.length; ++i) {
      uint64 remoteChainSelector = ccvConfigArgs[i].remoteChainSelector;
      address[] calldata outboundCCVs = ccvConfigArgs[i].outboundCCVs;
      address[] calldata inboundCCVs = ccvConfigArgs[i].inboundCCVs;

      // Validate and check for duplicates in outbound CCVs.
      _validateCCVArray(outboundCCVs);

      // Validate and check for duplicates in inbound CCVs.
      _validateCCVArray(inboundCCVs);

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
    uint64 sourceChainSelector,
    uint256, // amount
    bytes calldata // tokenArgs
  ) external view virtual returns (address[] memory requiredCCVs) {
    return s_verifierConfig[sourceChainSelector].inboundCCVs;
  }

  /// @notice Returns the set of required CCVs for outgoing messages to a destination chain.
  /// @param destChainSelector The destination chain selector for outgoing messages.
  /// This implementation assumes the same set of CCVs are used for all transfers on a lane.
  /// Implementers can override this function to define custom logic based on these params.
  /// @return requiredCCVs Set of required CCV addresses.
  function getRequiredOutboundCCVs(
    uint64 destChainSelector,
    uint256, // amount
    bytes calldata // tokenArgs
  ) external view virtual returns (address[] memory requiredCCVs) {
    return s_verifierConfig[destChainSelector].outboundCCVs;
  }

  /// @notice Checks a CCV address array for duplicate entries.
  /// @param ccvs The array of CCV addresses to validate.
  function _validateCCVArray(
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
}
