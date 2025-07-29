// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../libraries/Pool.sol";
import {TokenPool} from "../TokenPool.sol";
import {USDCTokenPool} from "./USDCTokenPool.sol";

import {IERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from
  "@chainlink/contracts/src/v0.8/vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

bytes4 constant LOCK_RELEASE_FLAG = 0xfa7c07de;

/// TODO: Add comments
contract USDCTokenPoolProxy is TokenPool {
  using SafeERC20 for IERC20;

  event LockOrBurnMechanismUpdated(uint64 indexed remoteChainSelector, LockOrBurnMechanism mechanism);
  event PoolAddressesUpdated(PoolAddresses pools);

  error InvalidPoolAddresses();
  error InvalidLockOrBurnMechanism(LockOrBurnMechanism mechanism);
  error InvalidMessageVersion(uint32 version);

  // bytes4(keccak256("NO_CCTP_USE_LOCK_RELEASE"))

  // @solhint-disable-next-line gas-struct-packing
  struct PoolAddresses {
    address cctpV1Pool;
    address cctpV2Pool;
    address lockReleasePool;
  }

  PoolAddresses internal s_pools;

  enum LockOrBurnMechanism {
    INVALID_MECHANISM,
    CCTP_V1,
    CCTP_V2,
    LOCK_RELEASE
  }

  mapping(uint64 => LockOrBurnMechanism) public s_lockOrBurnMechanism;

  // Note: This constructor is only used for CCTP V2, which is why the supportedUSDCVersion is set to 1.
  constructor(
    IERC20 token,
    address router,
    address[] memory allowlist,
    address rmnProxy,
    address cctpV1Pool,
    address cctpV2Pool,
    address lockReleasePool
  ) TokenPool(token, 6, allowlist, rmnProxy, router) {
    if (cctpV1Pool == address(0) || cctpV2Pool == address(0) || lockReleasePool == address(0)) {
      revert InvalidPoolAddresses();
    }

    s_pools = PoolAddresses({cctpV1Pool: cctpV1Pool, cctpV2Pool: cctpV2Pool, lockReleasePool: lockReleasePool});
  }

  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) public virtual override returns (Pool.LockOrBurnOutV1 memory) {
    _validateLockOrBurn(lockOrBurnIn);

    LockOrBurnMechanism mechanism = s_lockOrBurnMechanism[lockOrBurnIn.remoteChainSelector];

    if (mechanism == LockOrBurnMechanism.INVALID_MECHANISM) {
      revert InvalidLockOrBurnMechanism(mechanism);
    }

    PoolAddresses memory pools = s_pools;

    address destinationPool;

    if (mechanism == LockOrBurnMechanism.LOCK_RELEASE) {
      destinationPool = pools.lockReleasePool;
    } else if (mechanism == LockOrBurnMechanism.CCTP_V1) {
      destinationPool = pools.cctpV1Pool;
    }
    else if (mechanism == LockOrBurnMechanism.CCTP_V2) {
      destinationPool = pools.cctpV2Pool;
    }

    // Transfer the tokens to the destination pool
    i_token.safeTransfer(destinationPool, lockOrBurnIn.amount);

    return USDCTokenPool(destinationPool).lockOrBurn(lockOrBurnIn);
  }

  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) public virtual override returns (Pool.ReleaseOrMintOutV1 memory) {
    _validateReleaseOrMint(releaseOrMintIn, releaseOrMintIn.sourceDenominatedAmount);

    PoolAddresses memory pools = s_pools;

    // If the source pool data is the lock release flag, we use the lock release pool.
    if (bytes4(releaseOrMintIn.sourcePoolData) == LOCK_RELEASE_FLAG) {
      return USDCTokenPool(pools.lockReleasePool).releaseOrMint(releaseOrMintIn);
    }

    uint32 version;
    bytes memory usdcMessage = releaseOrMintIn.sourcePoolData;
    // solhint-disable-next-line no-inline-assembly
    assembly {
      // We truncate using the datatype of the version variable, meaning
      // we will only be left with the first 4 bytes of the message when we cast it to uint32. We want the lower 4 bytes
      // to be the version when casted to a uint32 , so we only add 4. If you added 32, attempting to skip the first word
      // containing the length, then version would be in the upper-4 bytes of the corresponding slot, which
      // would not be as easily parsed into a uint32.
      version := mload(add(usdcMessage, 4)) // 0 + 4 = 4
    }

    // TODO: Comments
    if (version == 0) {
      return USDCTokenPool(pools.cctpV1Pool).releaseOrMint(releaseOrMintIn);
    } else if (version == 1) {
      return USDCTokenPool(pools.cctpV2Pool).releaseOrMint(releaseOrMintIn);
    } else {
      revert InvalidMessageVersion(version);
    }
  }

  function updatePoolAddresses(
    PoolAddresses calldata pools
  ) external onlyOwner {
    if (pools.cctpV1Pool == address(0) || pools.cctpV2Pool == address(0) || pools.lockReleasePool == address(0)) {
      revert InvalidPoolAddresses();
    }

    s_pools = pools;

    emit PoolAddressesUpdated(pools);
  }

  function getPools() public view returns (PoolAddresses memory) {
    return s_pools;
  }

  function updateLockOrBurnMechanisms(
    uint64[] calldata remoteChainSelectors,
    LockOrBurnMechanism[] calldata mechanisms
  ) external onlyOwner {
    for (uint256 i = 0; i < remoteChainSelectors.length; i++) {
      s_lockOrBurnMechanism[remoteChainSelectors[i]] = mechanisms[i];
      emit LockOrBurnMechanismUpdated(remoteChainSelectors[i], mechanisms[i]);
    }
  }
}
