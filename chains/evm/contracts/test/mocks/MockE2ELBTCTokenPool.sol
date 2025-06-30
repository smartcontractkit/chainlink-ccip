// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../libraries/Pool.sol";
import {TokenPool} from "../../pools/TokenPool.sol";

import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";
import {IBurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/IBurnMintERC20.sol";

/// @notice MockE2ELBTCTokenPool is a token pool used for e2e tests. It allows to burn tokens unconditionally,
/// while requires specific structure for offchain token data
contract MockE2ELBTCTokenPool is TokenPool, ITypeAndVersion {
  error TokenDataMismatch(bytes32 expected, bytes32 actual);

  string public constant override typeAndVersion = "MockE2ELBTCTokenPool 1.5.1";

  // s_destPoolData has either a 32-byte or non-32-byte value, which changes the off-chain behavior.
  // If it is 32 bytes, the off-chain considers attestation enabled and calls the attestation API.
  // If it is non-32 bytes, the off-chain considers attestation disabled.
  bytes public s_destPoolData;

  constructor(
    IBurnMintERC20 token,
    address[] memory allowlist,
    address rmnProxy,
    address router,
    bytes memory destPoolData
  ) TokenPool(token, 8, allowlist, rmnProxy, router) {
    s_destPoolData = destPoolData;
  }

  /// @notice Overrides base lockOrBurn method in order to provide custom destPoolData. This destPoolData imitates
  /// Lombard Attestation API input. If offchain plugin is configured to retrieve attestation from Lombard API,
  /// it checks source pool address and if destPoolData is 32 bytes length, it makes a request.
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) public virtual override returns (Pool.LockOrBurnOutV1 memory) {
    _validateLockOrBurn(lockOrBurnIn);

    IBurnMintERC20(address(i_token)).burn(lockOrBurnIn.amount);

    emit LockedOrBurned({
      remoteChainSelector: lockOrBurnIn.remoteChainSelector,
      token: address(i_token),
      sender: msg.sender,
      amount: lockOrBurnIn.amount
    });

    return Pool.LockOrBurnOutV1({
      destTokenAddress: getRemoteToken(lockOrBurnIn.remoteChainSelector),
      destPoolData: s_destPoolData
    });
  }

  /// @notice Overrides base releaseOrMint method in order to add new logic regarding offchainTokenData. The condition
  /// requires offchainTokenData to be abi-encoding of two bytes arrays: payload + signatures. Then we check that
  /// message originally sent with sourcePoolData == sha256(payload)
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) public virtual override returns (Pool.ReleaseOrMintOutV1 memory) {
    uint256 amount = releaseOrMintIn.sourceDenominatedAmount;
    _validateReleaseOrMint(releaseOrMintIn, amount);

    if (s_destPoolData.length == 32) {
      (bytes memory payload,) = abi.decode(releaseOrMintIn.offchainTokenData, (bytes, bytes));
      bytes32 payloadHash = sha256(payload);
      if (payloadHash != bytes32(releaseOrMintIn.sourcePoolData)) {
        revert TokenDataMismatch(bytes32(releaseOrMintIn.sourcePoolData), payloadHash);
      }
    }

    // Mint to the receiver
    IBurnMintERC20(address(i_token)).mint(releaseOrMintIn.receiver, amount);

    emit ReleasedOrMinted({
      remoteChainSelector: releaseOrMintIn.remoteChainSelector,
      token: address(i_token),
      sender: msg.sender,
      recipient: releaseOrMintIn.receiver,
      amount: amount
    });

    return Pool.ReleaseOrMintOutV1({destinationAmount: amount});
  }
}
