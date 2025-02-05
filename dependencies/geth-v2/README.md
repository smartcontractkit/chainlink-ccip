# geth-v2 dependency
Background: original geth dependency got overcomplicated over time, and it made sense to rewrite it, so it can support new requirements without too much unnecessary complexity.

## Key advantages of the v2 version:
- Each chain is deployed as separate helm release. It is required to be able to control the deployment flow in devspace pipeline. For example when deploying 30 chains for load testing, we could deploy them in batches, to prevent overloading kubernetes by deploying to many manifests at the same time ([slack thread](https://chainlink-core.slack.com/archives/C081V9NJN3T/p1738246809276759))
- simplified generating multiple chains, instead of relying on hardcoded 1337 and 2337 and `ADDITIONAL_CHAINS_COUNT` property, now user need to just pass `CHAINS_COUNT` property and it generates variable number of helm deployments
- Allow to customize chain settings like blockTime, for selected chains. The new script will read the config from the provided file path and override settings for selected chains.

## How to customize individual chains
In `./values/chain-overrides` dir, add the new file with overrides, check existing examples.
Define only the chains that you want to override, for example:

```
chains:
  - networkId: 1337
    blockTime: 2
  - networkId: 2337
    blockTime: 2
```

To activate the overrides, you'll need to pass `CHAIN_OVERRIDES_FILENAME` env var. 

## Adding a new use-case
When using script approach to generate deployments, we can't rely on the profiles for patching the config.
To compensate geth-v2 brings a patching mechanism similar to how it is handled in donut dependency.

To add your own patches add a new use-case.
1) Add patch file in under `values/geth-node/` for example `[use-case.ccip-load-tests.yaml](values/geth-node/use-case.ccip-load-tests.yaml)
2) Update the [generate_base_profile](scripts/generate_base_profile) script to handle your use case accordingly.