// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAddressBlocklist} from "../../interfaces/IAddressBlocklist.sol";
import {IPolicyEngine} from "@chainlink/policy-management/interfaces/IPolicyEngine.sol";

import {Pool} from "../../libraries/Pool.sol";
import {MultiTokenPoolHooks} from "./MultiTokenPoolHooks.sol";

/// @notice MultiTokenPoolHooks that also check both addresses of each transfer against an external address blocklist.
/// @dev On send, the hooks check the local sender and the remote token receiver. On receive, they check the local token
/// receiver and the remote sender. The blocklist takes EVM addresses, so the hooks convert a remote address first:
/// - 20 bytes are a raw EVM address.
/// - 32 bytes are converted to their low 160 bits. This covers left-padded EVM addresses and Solana public keys. Other
///   chain families with 32-byte addresses are converted the same way.
/// - Other lengths are not checked. The local address is still checked.
/// The blocklist must list a Solana account in the same converted form.
/// A Solana token receiver is usually the wallet that owns the associated token account. A sender can instead ask the
/// Solana OffRamp to use the receiver as the token account itself. Then the hooks check the token account, not its
/// owner, because the owner is not part of the message.
/// The blocklist can stop every transfer on the lanes that use these hooks. On receive, a blocked address or a
/// blocklist that reverts makes the message fail after the source chain has locked or burned the tokens. The message
/// can be executed again after the cause is fixed.
contract MultiTokenPoolHooksExternalBlocklist is MultiTokenPoolHooks {
  error InvalidBlocklist(IAddressBlocklist blocklist);
  error AddressBlocked(IAddressBlocklist blocklist, address account);

  event BlocklistSet(IAddressBlocklist indexed oldBlocklist, IAddressBlocklist indexed newBlocklist);

  /// @dev Zero disables the blocklist checks.
  IAddressBlocklist private s_blocklist;

  /// @param allowlist The original senders that can send tokens. An empty list disables the allowlist for good.
  /// @param policyEngine The policy engine to attach. Zero disables the policy engine runs.
  /// @param blocklist The blocklist to query. Zero disables the blocklist checks.
  /// @param authorizedCallers The pools that can call the hooks.
  constructor(
    address[] memory allowlist,
    IPolicyEngine policyEngine,
    IAddressBlocklist blocklist,
    address[] memory authorizedCallers
  ) MultiTokenPoolHooks(allowlist, policyEngine, authorizedCallers) {
    _setBlocklist(blocklist);
  }

  function typeAndVersion() external pure virtual override returns (string memory) {
    return "MultiTokenPoolHooksExternalBlocklist 2.1.0-dev";
  }

  /// @inheritdoc MultiTokenPoolHooks
  /// @dev Runs the checks of MultiTokenPoolHooks first, then checks the local sender and the remote token receiver
  /// against the blocklist.
  function preflightCheck(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn,
    bytes4 requestedFinalityConfig,
    bytes calldata tokenArgs,
    uint256 amountPostFee
  ) public virtual override {
    super.preflightCheck(lockOrBurnIn, requestedFinalityConfig, tokenArgs, amountPostFee);
    _checkBlocklist(lockOrBurnIn.originalSender, lockOrBurnIn.receiver);
  }

  /// @inheritdoc MultiTokenPoolHooks
  /// @dev Runs the checks of MultiTokenPoolHooks first, then checks the local token receiver and the remote sender
  /// against the blocklist.
  function postflightCheck(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn,
    uint256 localAmount,
    bytes4 requestedFinalityConfig
  ) public virtual override {
    super.postflightCheck(releaseOrMintIn, localAmount, requestedFinalityConfig);
    _checkBlocklist(releaseOrMintIn.receiver, releaseOrMintIn.originalSender);
  }

  /// @notice Sets the blocklist for all tokens and lanes that use these hooks.
  /// @param newBlocklist The new blocklist. It must have code. Zero disables the blocklist checks.
  function setBlocklist(
    IAddressBlocklist newBlocklist
  ) public virtual onlyOwner {
    _setBlocklist(newBlocklist);
  }

  /// @notice Returns the blocklist.
  /// @return blocklist The blocklist. Zero means the blocklist checks are disabled.
  function getBlocklist() public view virtual returns (IAddressBlocklist blocklist) {
    return s_blocklist;
  }

  /// @dev Shared by the constructor and setBlocklist. The code check catches an EOA set by mistake, which would make
  /// every check revert.
  function _setBlocklist(
    IAddressBlocklist newBlocklist
  ) internal virtual {
    if (address(newBlocklist) != address(0) && address(newBlocklist).code.length == 0) {
      revert InvalidBlocklist(newBlocklist);
    }
    IAddressBlocklist oldBlocklist = s_blocklist;
    if (address(newBlocklist) == address(oldBlocklist)) return;

    s_blocklist = newBlocklist;
    emit BlocklistSet(oldBlocklist, newBlocklist);
  }

  /// @dev Checks the local address, then the remote address converted as described on the contract.
  /// @param localAccount The address on this chain.
  /// @param remoteAccount The address on the remote chain, as the ramp passes it to the pool.
  function _checkBlocklist(
    address localAccount,
    bytes calldata remoteAccount
  ) internal view virtual {
    IAddressBlocklist blocklist = s_blocklist;
    if (address(blocklist) == address(0)) return;

    if (blocklist.isBlocked(localAccount)) revert AddressBlocked(blocklist, localAccount);

    address remoteAddress;
    if (remoteAccount.length == 20) {
      remoteAddress = address(bytes20(remoteAccount));
    } else if (remoteAccount.length == 32) {
      remoteAddress = address(uint160(uint256(bytes32(remoteAccount))));
    } else {
      return;
    }

    if (blocklist.isBlocked(remoteAddress)) revert AddressBlocked(blocklist, remoteAddress);
  }
}
