// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../libraries/Pool.sol";
import {TokenPool} from "../../pools/TokenPool.sol";

import {ITypeAndVersion} from "@chainlink/contracts/src/v0.8/shared/interfaces/ITypeAndVersion.sol";
import {IBurnMintERC20} from "@chainlink/contracts/src/v0.8/shared/token/ERC20/IBurnMintERC20.sol";

/// @notice MockE2ELBTCTokenPool is a token pool used for e2e tests. It inherits BurnMintTokenPool and allows to burn
/// tokens unconditionally, while requires some structure from mint payloads
contract MockE2ELBTCTokenPool is TokenPool, ITypeAndVersion {
    string public constant override typeAndVersion = "MockE2ELBTCTokenPool 1.5.1";

    // This variable i_destPoolData will have either a 32-byte or non-32-byte value, which will change the off-chain behavior.
    // If it is 32 bytes, the off-chain will consider it as attestation enabled and call the attestation API.
    // If it is non-32 bytes, the off-chain will consider it as attestation disabled.
    bytes public i_destPoolData;

    constructor(
        IBurnMintERC20 token,
        address[] memory allowlist,
        address rmnProxy,
        address router,
        bytes memory destPoolData
    ) TokenPool(token, 8, allowlist, rmnProxy, router) {
        i_destPoolData = destPoolData;
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
            destPoolData: i_destPoolData
        });
    }

    /// @notice Overrides base releaseOrMint method in order to add new logic regarding offchainTokenData. The condition
    /// simply requires offchainTokenData to be abi-encoding of two bytes32: sourcePoolData + keccak256("secret")
    /// This is needed to verify e2e flow.
    function releaseOrMint(
        Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
    ) public virtual override returns (Pool.ReleaseOrMintOutV1 memory) {
        _validateReleaseOrMint(releaseOrMintIn);

        (bytes memory payload,) = abi.decode(
            releaseOrMintIn.offchainTokenData,
            (bytes, bytes)
        );
        require(sha256(payload) == bytes32(releaseOrMintIn.sourcePoolData), "payload hash doesn't match");

        // Mint to the receiver
        IBurnMintERC20(address(i_token)).mint(releaseOrMintIn.receiver, releaseOrMintIn.amount);

        emit ReleasedOrMinted({
            remoteChainSelector: releaseOrMintIn.remoteChainSelector,
            token: address(i_token),
            sender: msg.sender,
            recipient: releaseOrMintIn.receiver,
            amount: releaseOrMintIn.amount
        });

        return Pool.ReleaseOrMintOutV1({destinationAmount: releaseOrMintIn.amount});
    }
}