// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAddressBlocklist} from "../../../../interfaces/IAddressBlocklist.sol";
import {IPolicyEngine} from "@chainlink/policy-management/interfaces/IPolicyEngine.sol";

import {Pool} from "../../../../libraries/Pool.sol";
import {
  MultiTokenPoolHooksExternalBlocklist
} from "../../../../pools/MultiTokenPool/MultiTokenPoolHooksExternalBlocklist.sol";
import {BaseTest} from "../../../BaseTest.t.sol";

contract MultiTokenPoolHooksExternalBlocklistSetup is BaseTest {
  /// @dev The SPL Token program ID, TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA, as the 32 raw bytes that CCIP carries
  /// for a Solana account.
  bytes32 internal constant SOLANA_ACCOUNT = 0x06ddf6e1d765a193d9cbe146ceeb79ac1cb485ed5f5b37913a8cf5857eff00a9;
  /// @dev The low 160 bits of SOLANA_ACCOUNT, computed offchain from its base58 form. A blocklist lists the Solana
  /// account as this address.
  address internal constant SOLANA_ACCOUNT_AS_ADDRESS = 0xCeeB79AC1cB485eD5F5b37913a8cf5857eff00A9;

  MultiTokenPoolHooksExternalBlocklist internal s_hooks;
  IAddressBlocklist internal s_blocklist = IAddressBlocklist(makeAddr("blocklist"));
  address internal s_sender = makeAddr("sender");
  address internal s_receiver = makeAddr("receiver");
  Pool.LockOrBurnInV1 internal s_lockOrBurnIn;
  Pool.ReleaseOrMintInV1 internal s_releaseOrMintIn;

  function setUp() public virtual override {
    super.setUp();
    _deployBlocklist(s_blocklist);
    s_hooks = new MultiTokenPoolHooksExternalBlocklist(
      new address[](0), IPolicyEngine(address(0)), s_blocklist, _arrayOf(OWNER)
    );

    // The 2.0 OnRamp passes the token receiver as the caller gave it. An EVM receiver is 20 bytes or 32 bytes with
    // left padding.
    s_lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: abi.encode(s_receiver),
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: s_sender,
      amount: 1e18,
      localToken: makeAddr("token")
    });
    // An EVM OnRamp writes the sender as abi.encode(address), so an EVM sender reaches the OffRamp as 32 bytes.
    s_releaseOrMintIn = Pool.ReleaseOrMintInV1({
      originalSender: abi.encode(s_sender),
      remoteChainSelector: SOURCE_CHAIN_SELECTOR,
      receiver: s_receiver,
      sourceDenominatedAmount: 1e18,
      localToken: s_lockOrBurnIn.localToken,
      sourcePoolAddress: abi.encode(makeAddr("sourcePool")),
      sourcePoolData: "",
      offchainTokenData: ""
    });
  }

  /// @dev Gives the blocklist code, so the hooks accept it, and blocks no address.
  function _deployBlocklist(
    IAddressBlocklist blocklist
  ) internal {
    vm.etch(address(blocklist), hex"00");
    vm.mockCall(address(blocklist), abi.encodeWithSelector(IAddressBlocklist.isBlocked.selector), abi.encode(false));
  }

  function _setBlocked(
    address account,
    bool blocked
  ) internal {
    vm.mockCall(address(s_blocklist), abi.encodeCall(IAddressBlocklist.isBlocked, (account)), abi.encode(blocked));
  }

  /// @dev Remote addresses that have no EVM form. The hooks check only the local address for them.
  function _remoteAddressesWithoutEVMForm() internal pure returns (bytes[] memory accounts) {
    accounts = new bytes[](5);
    accounts[0] = "";
    accounts[1] = new bytes(19);
    accounts[2] = new bytes(21);
    accounts[3] = new bytes(31);
    accounts[4] = new bytes(33);
    return accounts;
  }
}
