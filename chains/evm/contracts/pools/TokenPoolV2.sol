// SPDX-License-Identifier: BUSL-1.1

import {IPoolV2} from "../interfaces/IPoolV2.sol";

import {Pool} from "../libraries/Pool.sol";
import {TokenPool} from "./TokenPool.sol";

import {IERC165} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/introspection/IERC165.sol";

abstract contract TokenPoolV2 is IPoolV2, TokenPool {
  event DynamicVerifiersSet(uint64 indexed remoteChainSelector, address[] outbound, address[] inbound);

  struct DynamicVerifierConfig {
    address[] outbound;
    address[] inbound;
  }

  mapping(uint64 remoteChainSelector => DynamicVerifierConfig) internal s_dynamicVerifiers;

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
    bytes calldata /* tokenExtraData */
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
  // │                            CCV                               │
  // ================================================================

  function setRequiredCCVs(
    uint64 remoteChainSelector,
    address[] calldata outbound,
    address[] calldata inbound
  ) external onlyOwner {
    s_dynamicVerifiers[remoteChainSelector] = DynamicVerifierConfig({outbound: outbound, inbound: inbound});
    emit DynamicVerifiersSet({remoteChainSelector: remoteChainSelector, outbound: outbound, inbound: inbound});
  }

  function getRequiredInboundCCVs(
    uint64 sourceChainSelector,
    uint256,
    bytes calldata
  ) external view virtual returns (address[] memory) {
    DynamicVerifierConfig memory config = s_dynamicVerifiers[sourceChainSelector];
    if (config.inbound.length > 0) {
      return config.inbound;
    }
    return s_dynamicVerifiers[sourceChainSelector].inbound;
  }

  function getRequiredOutboundCCVs(
    uint64 destChainSelector,
    uint256,
    bytes calldata
  ) external view virtual returns (address[] memory) {
    DynamicVerifierConfig memory config = s_dynamicVerifiers[destChainSelector];
    if (config.outbound.length > 0) {
      return config.outbound;
    }
    return s_dynamicVerifiers[destChainSelector].outbound;
  }
}
