# solana-validator dependency
Enables to deploy solana-test-validator to simulate solana blockchain for testing.

## How to customize individual chains
In `./values/chain-overrides` dir, add the new file with overrides, check existing examples.
Define only the chains that you want to override, for example:

```
chains:
  - networkId: 1001
    blockTime: 2
  - networkId: 1002
    blockTime: 2
```

To activate the overrides, you'll need to pass `CHAIN_OVERRIDES_FILENAME` env var. 
