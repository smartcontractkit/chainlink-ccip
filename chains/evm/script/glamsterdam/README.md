# Glamsterdam post-fork send checks

`GlamsterdamSend.s.sol` sends CCIP messages from Base Sepolia or Avalanche Fuji to Eth Sepolia. Run it after the gas-config changesets execute, and again after the fork. Every message goes to a freshly derived random receiver, so the destination account is cold (worst-case gas).

| Entrypoint | Message |
|---|---|
| `sendEmpty()` | no data, no tokens |
| `sendBnM()` | CCIP-BnM transfer |
| `sendUsdc()` | USDC transfer |
| `sendLombard()` | LBTC transfer (fill in the `TODO` token addresses first) |
| `sendAll()` | all four, one tx each |

## Run

```sh
# Dry run (simulation on a fork, no spend). Do this first.
forge script script/glamsterdam/GlamsterdamSend.s.sol --sig "sendAll()" --rpc-url $BASE_SEPOLIA_RPC

# Live
forge script script/glamsterdam/GlamsterdamSend.s.sol --sig "sendAll()" --rpc-url $BASE_SEPOLIA_RPC --broadcast --account <keystore>
```

Repeat with `--rpc-url $FUJI_RPC`. Use `MODE=v16` to test v1.6 lanes (default is `MODE=v2`).

Env: `MODE`, `ROUTER` (e.g. to use the TestRouter), `BNM_AMOUNT`, `USDC_AMOUNT`, `LOMBARD_AMOUNT`, `PRIVATE_KEY` (optional).

The sender needs native gas, BnM, USDC (Circle faucet) and LBTC on each source chain.

## What the script checks before sending
- `Router.getOnRamp(EthSepolia)` must equal the v2.0 OnRamp (or v1.6 with `MODE=v16`). Otherwise it reverts, so a lane that is not on v2 is never reported as v2.
- Sender native and token balances cover the fee and amount.
- For token sends it logs the pool and, on v2, the pool-required CCVs.

## Verify
1. Open each logged `https://ccip.chain.link/msg/<messageId>` and confirm the message reaches Success on Eth Sepolia.
2. Check each logged receiver's balance on Eth Sepolia, e.g. `cast call <token> "balanceOf(address)(uint256)" <receiver> --rpc-url $ETH_SEPOLIA_RPC`.
3. Suggested order: baseline before the changeset, again after the changeset, again after the fork.

If a dry-run reverts on the empty message or on USDC, that is a finding about the lane, not necessarily a script bug. See the plan notes on the empty message and the USDC pool mechanism.
