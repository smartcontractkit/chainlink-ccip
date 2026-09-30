// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.24;

import {IPoolV2} from "../../contracts/interfaces/IPoolV2.sol";
import {Client, IRouterClient} from "../../contracts/interfaces/IRouterClient.sol";
import {ITokenAdminRegistry} from "../../contracts/interfaces/ITokenAdminRegistry.sol";
import {ExtraArgsCodec} from "../../contracts/libraries/ExtraArgsCodec.sol";

import {Script, console} from "forge-std/Script.sol";

interface IERC20Min {
  function approve(
    address spender,
    uint256 amount
  ) external returns (bool);
  function balanceOf(
    address account
  ) external view returns (uint256);
}

interface IRouterLanes {
  function getOnRamp(
    uint64 destChainSelector
  ) external view returns (address);
  function isChainSupported(
    uint64 chainSelector
  ) external view returns (bool);
}

/// @notice Post-Glamsterdam smoke test. Sends one CCIP message per case from a testnet source chain to Eth Sepolia,
/// each to a fresh random receiver so the destination account is cold (worst-case gas).
/// @dev Run with `--rpc-url` for Base Sepolia or Avalanche Fuji. The chain table is chosen by block.chainid.
///
///   Env:
///     MODE          "v2" (default) or "v16". The lane's OnRamp must match, otherwise the script reverts.
///     ROUTER        optional override (e.g. the TestRouter).
///     BNM_AMOUNT    default 1e15, USDC_AMOUNT default 1e5, LOMBARD_AMOUNT default 1e3.
///     PRIVATE_KEY   optional, otherwise use --account / --ledger.
///
///   Dry run (no spend):  forge script script/glamsterdam/GlamsterdamSend.s.sol --sig "sendAll()" --rpc-url $RPC
///   Live:                add --broadcast
contract GlamsterdamSend is Script {
  uint64 internal constant ETH_SEPOLIA_SELECTOR = 16015286601757825753;

  // TODO: replace with the real LBTC token addresses.
  address internal constant DUMMY_LBTC_BASE_SEPOLIA = address(0xdead000000000000000000000000000000000001);
  address internal constant DUMMY_LBTC_FUJI = address(0xdEaD000000000000000000000000000000000002);

  struct ChainConfig {
    string name;
    address router;
    address v2OnRamp;
    address v16OnRamp;
    address tokenAdminRegistry; // zero = skip pool logging
    address bnm;
    address usdc;
    address lombard;
  }

  function sendEmpty() external {
    _run(true, false, false, false);
  }

  function sendBnM() external {
    _run(false, true, false, false);
  }

  function sendUsdc() external {
    _run(false, false, true, false);
  }

  function sendLombard() external {
    _run(false, false, false, true);
  }

  function sendAll() external {
    _run(true, true, true, true);
  }

  // ---------------------------------------------------------------------------------------------

  function _run(
    bool empty,
    bool bnm,
    bool usdc,
    bool lombard
  ) internal {
    ChainConfig memory cfg = _config();
    bool v2 = _isV2Mode();
    address router = vm.envOr("ROUTER", cfg.router);

    _preflightLane(cfg, router, v2);

    uint256 pk = vm.envOr("PRIVATE_KEY", uint256(0));
    if (pk == 0) vm.startBroadcast();
    else vm.startBroadcast(pk);

    uint256 nonce;
    if (empty) _send(cfg, router, v2, "EMPTY", address(0), 0, nonce++);
    if (bnm) _send(cfg, router, v2, "BNM", cfg.bnm, vm.envOr("BNM_AMOUNT", uint256(1e15)), nonce++);
    if (usdc) _send(cfg, router, v2, "USDC", cfg.usdc, vm.envOr("USDC_AMOUNT", uint256(1e5)), nonce++);
    if (lombard) {
      require(
        cfg.lombard != DUMMY_LBTC_BASE_SEPOLIA && cfg.lombard != DUMMY_LBTC_FUJI,
        "Lombard token is still a dummy address, fill it in"
      );
      _send(cfg, router, v2, "LOMBARD", cfg.lombard, vm.envOr("LOMBARD_AMOUNT", uint256(1e3)), nonce++);
    }

    vm.stopBroadcast();
  }

  function _send(
    ChainConfig memory cfg,
    address router,
    bool v2,
    string memory label,
    address token,
    uint256 amount,
    uint256 nonce
  ) internal {
    address receiver = address(uint160(uint256(keccak256(abi.encode(block.timestamp, msg.sender, label, nonce)))));

    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(receiver),
      data: "",
      tokenAmounts: new Client.EVMTokenAmount[](token == address(0) ? 0 : 1),
      feeToken: address(0),
      extraArgs: _extraArgs(v2)
    });
    if (token != address(0)) {
      message.tokenAmounts[0] = Client.EVMTokenAmount({token: token, amount: amount});
      require(IERC20Min(token).balanceOf(msg.sender) >= amount, string.concat(label, ": insufficient token balance"));
      _logPool(cfg, token, amount);
    }

    uint256 fee = IRouterClient(router).getFee(ETH_SEPOLIA_SELECTOR, message);
    require(msg.sender.balance >= fee, string.concat(label, ": insufficient native balance for fee"));

    if (token != address(0)) IERC20Min(token).approve(router, amount);

    // Router wraps all of msg.value and does not refund, so send a 5% buffer only.
    bytes32 messageId = IRouterClient(router).ccipSend{value: fee + fee / 20}(ETH_SEPOLIA_SELECTOR, message);

    console.log("---", label);
    console.log("receiver", receiver);
    console.log("fee (wei)", fee);
    console.logBytes32(messageId);
    console.log(string.concat("https://ccip.chain.link/msg/", vm.toString(messageId)));
  }

  function _extraArgs(
    bool v2
  ) internal pure returns (bytes memory) {
    if (v2) {
      // gasLimit 0 + empty data = no callback (EOA receiver). Finality 0 = wait for finality.
      // Empty ccvs = lane defaults; pool-required CCVs (CCTP / Lombard) are added by the OnRamp.
      return ExtraArgsCodec._getBasicEncodedExtraArgsV3(0, bytes4(0));
    }
    return Client._argsToBytes(Client.GenericExtraArgsV2({gasLimit: 0, allowOutOfOrderExecution: true}));
  }

  function _preflightLane(
    ChainConfig memory cfg,
    address router,
    bool v2
  ) internal view {
    console.log("Source chain:", cfg.name);
    console.log("Router:", router);
    require(IRouterLanes(router).isChainSupported(ETH_SEPOLIA_SELECTOR), "Eth Sepolia not supported by router");

    address onRamp = IRouterLanes(router).getOnRamp(ETH_SEPOLIA_SELECTOR);
    console.log("Lane OnRamp:", onRamp);
    if (onRamp == cfg.v2OnRamp) console.log("Lane version: v2.0");
    else if (onRamp == cfg.v16OnRamp) console.log("Lane version: v1.6");
    else console.log("Lane version: UNKNOWN (not v2.0 or v1.6)");

    require(
      onRamp == (v2 ? cfg.v2OnRamp : cfg.v16OnRamp),
      v2 ? "lane is not served by the v2.0 OnRamp (MODE=v2)" : "lane is not served by the v1.6 OnRamp (MODE=v16)"
    );
  }

  function _logPool(
    ChainConfig memory cfg,
    address token,
    uint256 amount
  ) internal view {
    if (cfg.tokenAdminRegistry == address(0)) return;
    address pool = ITokenAdminRegistry(cfg.tokenAdminRegistry).getPool(token);
    console.log("pool", pool);
    if (!_isV2Mode()) return;
    try IPoolV2(pool)
      .getRequiredCCVs(
        token, ETH_SEPOLIA_SELECTOR, amount, bytes4(0), bytes(""), IPoolV2.MessageDirection.Outbound
      ) returns (
      address[] memory ccvs
    ) {
      for (uint256 i = 0; i < ccvs.length; ++i) {
        console.log("pool required CCV", ccvs[i]);
      }
    } catch {
      console.log("getRequiredCCVs lookup failed (pool may be v1)");
    }
  }

  function _isV2Mode() internal view returns (bool) {
    string memory mode = vm.envOr("MODE", string("v2"));
    bytes32 h = keccak256(bytes(mode));
    require(h == keccak256("v2") || h == keccak256("v16"), "MODE must be v2 or v16");
    return h == keccak256("v2");
  }

  function _config() internal view returns (ChainConfig memory) {
    if (block.chainid == 84532) {
      return ChainConfig({
        name: "Base Sepolia",
        router: 0xD3b06cEbF099CE7DA4AcCf578aaebFDBd6e88a93,
        v2OnRamp: 0xA33b221A8427739c76f631a995ca60544bEdD632,
        v16OnRamp: 0x28A025d34c830BF212f5D2357C8DcAB32dD92A20,
        tokenAdminRegistry: 0x736D0bBb318c1B27Ff686cd19804094E66250e17,
        bnm: 0x88A2d74F47a237a62e7A51cdDa67270CE381555e,
        usdc: 0x036CbD53842c5426634e7929541eC2318f3dCF7e,
        lombard: DUMMY_LBTC_BASE_SEPOLIA
      });
    }
    if (block.chainid == 43113) {
      return ChainConfig({
        name: "Avalanche Fuji",
        router: 0xF694E193200268f9a4868e4Aa017A0118C9a8177,
        v2OnRamp: 0x8DBCd8BBbF832c65b4a1A927eD88B92A47bd06E4,
        v16OnRamp: 0xA5D5B0B844c8f11B61F28AC98BBA84dEA9b80953,
        tokenAdminRegistry: 0xA92053a4a3922084d992fD2835bdBa4caC6877e6,
        bnm: 0xD21341536c5cF5EB1bcb58f6723cE26e8D8E90e4,
        usdc: 0x5425890298aed601595a70AB815c96711a31Bc65,
        lombard: DUMMY_LBTC_FUJI
      });
    }
    revert("unsupported chain: use Base Sepolia (84532) or Fuji (43113)");
  }
}
