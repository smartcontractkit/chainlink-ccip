// SPDX-License-Identifier: BUSL-1.1

import {IPoolV2} from "../interfaces/IPoolV2.sol";

import {Pool} from "../libraries/Pool.sol";
import {TokenPool} from "./TokenPool.sol";

import {IERC165} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v5.0.2/contracts/utils/introspection/IERC165.sol";

abstract contract TokenPoolV2 is IPoolV2, TokenPool {
  // ================================================================
  // │                        Lock or Burn                          │
  // ================================================================

  function lockOrBurn(
    Pool.LockOrBurnV2 calldata lockOrBurnIn
  ) public virtual override returns (Pool.LockOrBurnOutV1 memory) {
    _validateLockOrBurn(
      lockOrBurnIn.localToken, lockOrBurnIn.remoteChainSelector, lockOrBurnIn.originalSender, lockOrBurnIn.amount
    );
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

  /// @notice Signals which version of the pool interface is supported
  function supportsInterface(
    bytes4 interfaceId
  ) public pure virtual override(TokenPool, IERC165) returns (bool) {
    return interfaceId == Pool.CCIP_POOL_V2 || interfaceId == type(IPoolV2).interfaceId
      || interfaceId == type(IERC165).interfaceId;
  }
}
