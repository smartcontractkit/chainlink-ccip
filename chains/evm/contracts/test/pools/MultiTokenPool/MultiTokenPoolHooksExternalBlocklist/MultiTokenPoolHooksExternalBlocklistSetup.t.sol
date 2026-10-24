// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {IAddressBlocklist} from "../../../../interfaces/IAddressBlocklist.sol";
import {IPolicyEngine} from "@chainlink/policy-management/interfaces/IPolicyEngine.sol";

import {Pool} from "../../../../libraries/Pool.sol";
import {
  MultiTokenPoolHooksExternalBlocklist
} from "../../../../pools/MultiTokenPool/MultiTokenPoolHooksExternalBlocklist.sol";
import {BaseTest} from "../../../BaseTest.t.sol";
import {MockPolicyEngine} from "../../../mocks/MockPolicyEngine.sol";

contract MultiTokenPoolHooksExternalBlocklistSetup is BaseTest {
  /// @dev The SPL Token program ID TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA, used as a well-known Solana address.
  bytes32 internal constant SOLANA_ACCOUNT = 0x06ddf6e1d765a193d9cbe146ceeb79ac1cb485ed5f5b37913a8cf5857eff00a9;
  /// @dev The low 160 bits of SOLANA_ACCOUNT, decoded offchain from the base58 address. This is how the blocklist
  /// lists that Solana address.
  address internal constant SOLANA_ACCOUNT_AS_ADDRESS = 0xCeeB79AC1cB485eD5F5b37913a8cf5857eff00A9;

  MultiTokenPoolHooksExternalBlocklist internal s_hooks;
  IAddressBlocklist internal s_blocklist = IAddressBlocklist(makeAddr("blocklist"));
  MockPolicyEngine internal s_policyEngine;
  address internal s_pool = makeAddr("pool");
  address internal s_sender = makeAddr("sender");
  address internal s_receiver = makeAddr("receiver");
  Pool.LockOrBurnInV1 internal s_lockOrBurnIn;
  Pool.ReleaseOrMintInV1 internal s_releaseOrMintIn;

  function setUp() public virtual override {
    super.setUp();
    _mockBlocklist(s_blocklist);
    s_policyEngine = new MockPolicyEngine();
    // The pool is the only authorized caller, so the tests can tell the caller apart from the owner.
    s_hooks = new MultiTokenPoolHooksExternalBlocklist(
      new address[](0), IPolicyEngine(address(0)), s_blocklist, _arrayOf(s_pool)
    );

    // The 2.0 OnRamp passes the receiver to the pool as the sender gave it, so an EVM receiver can be ABI encoded or
    // 20 raw bytes.
    s_lockOrBurnIn = Pool.LockOrBurnInV1({
      receiver: abi.encode(s_receiver),
      remoteChainSelector: DEST_CHAIN_SELECTOR,
      originalSender: s_sender,
      amount: 1e18,
      localToken: makeAddr("token")
    });
    // EVM OnRamps ABI encode the sender, so it arrives as 32 bytes.
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

  /// @dev Puts code at the blocklist address so the hooks accept it, and makes it report every address as not blocked.
  function _mockBlocklist(
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

  /// @dev Remote addresses that are neither 20 nor 32 bytes long, which the hooks do not check. 36 bytes is the length
  /// of a TON address.
  function _remoteAddressesWithoutEVMForm() internal pure returns (bytes[] memory accounts) {
    accounts = new bytes[](6);
    accounts[0] = "";
    accounts[1] = new bytes(19);
    accounts[2] = new bytes(21);
    accounts[3] = new bytes(31);
    accounts[4] = new bytes(33);
    accounts[5] = new bytes(36);
    return accounts;
  }
}
