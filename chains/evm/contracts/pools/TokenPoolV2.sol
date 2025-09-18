// SPDX-License-Identifier: BUSL-1.1

import {TokenPool} from "./TokenPool.sol";

import {IPoolV2} from "../interfaces/IPoolV2.sol";

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
}
