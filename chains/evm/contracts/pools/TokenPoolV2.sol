// SPDX-License-Identifier: BUSL-1.1

import {IPoolV2} from "../interfaces/IPoolV2.sol";

import {Pool} from "../libraries/Pool.sol";
import {TokenPool} from "./TokenPool.sol";

import {IERC165} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/introspection/IERC165.sol";

abstract contract TokenPoolV2 is IPoolV2, TokenPool {
  event DestChainConfigUpdated(uint64 indexed destChainSelector, DestChainConfig destChainConfig);

  struct DestChainConfigArg {
    uint64 destChainSelector;
    address[] outboundCCVs;
    address[] inboundCCVs;
  }
  // TODO billing related config args

  struct DestChainConfig {
    address[] outboundCCVs;
    address[] inboundCCVs;
  }
  // TODO billing related config

  mapping(uint64 remoteChainSelector => DestChainConfig) internal s_destChainConfig;

  /// @notice Signals which version of the pool interface is supported.
  function supportsInterface(
    bytes4 interfaceId
  ) public pure virtual override(TokenPool, IERC165) returns (bool) {
    return interfaceId == Pool.CCIP_POOL_V2 || interfaceId == type(IPoolV2).interfaceId
      || interfaceId == type(IERC165).interfaceId;
  }

  // ================================================================
  // │                        Lock or Burn                          │
  // ================================================================

  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    bytes calldata /* tokenArgs */
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
  // │                          Config                               │
  // ================================================================

  function applyDestChainConfigUpdates(
    uint64 remoteChainSelector,
    address[] calldata outbound,
    address[] calldata inbound
  ) external onlyOwner {
    s_destChainConfig[remoteChainSelector] = DestChainConfig({outboundCCVs: outbound, inboundCCVs: inbound});
    emit DestChainConfigUpdated(remoteChainSelector, s_destChainConfig[remoteChainSelector]);
  }

  // ================================================================
  // │                          CCV                                  │
  // ================================================================

  function getRequiredInboundCCVs(
    uint64 sourceChainSelector,
    uint256,
    bytes calldata
  ) external view virtual returns (address[] memory) {
    return s_destChainConfig[sourceChainSelector].inboundCCVs;
  }

  function getRequiredOutboundCCVs(
    uint64 destChainSelector,
    uint256,
    bytes calldata
  ) external view virtual returns (address[] memory) {
    return s_destChainConfig[destChainSelector].outboundCCVs;
  }
}
