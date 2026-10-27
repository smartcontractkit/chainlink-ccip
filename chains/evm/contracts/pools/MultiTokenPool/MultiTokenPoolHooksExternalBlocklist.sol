// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAddressBlocklist} from "../../interfaces/IAddressBlocklist.sol";
import {IPolicyEngine} from "@chainlink/policy-management/interfaces/IPolicyEngine.sol";

import {Pool} from "../../libraries/Pool.sol";
import {MultiTokenPoolHooks} from "./MultiTokenPoolHooks.sol";

/// @notice MultiTokenPoolHooks that also check the sender and the receiver against an external address blocklist.
/// @dev The blocklist only takes EVM addresses. A 32-byte remote address, such as a Solana, Aptos or Sui address, is
/// looked up by its low 160 bits.
contract MultiTokenPoolHooksExternalBlocklist is MultiTokenPoolHooks {
  error InvalidBlocklist(IAddressBlocklist blocklist);
  error LocalAddressBlocked(IAddressBlocklist blocklist, address account);
  error RemoteAddressBlocked(IAddressBlocklist blocklist, bytes account);

  event BlocklistSet(IAddressBlocklist indexed oldBlocklist, IAddressBlocklist indexed newBlocklist);

  /// @dev The blocklist to check against. 0 address skips blocklist checks.
  IAddressBlocklist private s_blocklist;

  /// @param allowlist The addresses allowed to send tokens. An empty list disables the allowlist permanently.
  /// @param policyEngine The policy engine to attach. Zero disables policy engine checks.
  /// @param blocklist The blocklist to check against. Zero disables blocklist checks.
  /// @param authorizedCallers The pools allowed to call the hooks.
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
  /// @dev Also checks the sender and the receiver against the blocklist.
  /// When sending to Solana, the receiver is usually a wallet, and the tokens go to that wallet's associated token
  /// account. The sender can also pass a token account as the receiver. That token account is then checked instead of
  /// the wallet that owns it.
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
  /// @dev Also checks the receiver and the sender against the blocklist.
  function postflightCheck(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn,
    uint256 localAmount,
    bytes4 requestedFinalityConfig
  ) public virtual override {
    super.postflightCheck(releaseOrMintIn, localAmount, requestedFinalityConfig);
    _checkBlocklist(releaseOrMintIn.receiver, releaseOrMintIn.originalSender);
  }

  /// @notice Sets a new blocklist.
  /// @param newBlocklist The address of the new blocklist, or address(0) to disable.
  function setBlocklist(
    IAddressBlocklist newBlocklist
  ) public virtual onlyOwner {
    _setBlocklist(newBlocklist);
  }

  /// @notice Gets the current blocklist.
  /// @return blocklist The address of the blocklist. Zero if blocklist checks are disabled.
  function getBlocklist() public view virtual returns (IAddressBlocklist blocklist) {
    return s_blocklist;
  }

  /// @notice Internal version of setBlocklist to allow for reuse in the constructor.
  /// @param newBlocklist The address of the new blocklist, or address(0) to disable.
  function _setBlocklist(
    IAddressBlocklist newBlocklist
  ) internal virtual {
    // Calling isBlocked on an address without code reverts, which would stop every transfer.
    if (address(newBlocklist) != address(0) && address(newBlocklist).code.length == 0) {
      revert InvalidBlocklist(newBlocklist);
    }
    IAddressBlocklist oldBlocklist = s_blocklist;
    if (address(newBlocklist) == address(oldBlocklist)) return;

    s_blocklist = newBlocklist;
    emit BlocklistSet(oldBlocklist, newBlocklist);
  }

  /// @notice Reverts if the local or the remote account is blocked.
  /// @dev The remote account is turned into an EVM address before the lookup:
  /// - 20 bytes are used as is.
  /// - 32 bytes keep their low 160 bits. This is the usual bytes32 to address cast, so an ABI encoded EVM address keeps
  ///   its value.
  /// - Other lengths have no EVM form and are not checked.
  /// @param localAccount The account on this chain.
  /// @param remoteAccount The account on the remote chain, as the ramp passes it to the pool.
  function _checkBlocklist(
    address localAccount,
    bytes calldata remoteAccount
  ) internal view virtual {
    IAddressBlocklist blocklist = s_blocklist;
    if (address(blocklist) == address(0)) return;

    if (blocklist.isBlocked(localAccount)) revert LocalAddressBlocked(blocklist, localAccount);

    address remoteAddress;
    if (remoteAccount.length == 20) {
      remoteAddress = address(bytes20(remoteAccount));
    } else if (remoteAccount.length == 32) {
      remoteAddress = address(uint160(uint256(bytes32(remoteAccount))));
    } else {
      return;
    }

    if (blocklist.isBlocked(remoteAddress)) revert RemoteAddressBlocked(blocklist, remoteAccount);
  }
}
