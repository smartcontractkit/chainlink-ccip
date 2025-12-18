<div align="center">

# CCIP Developer Environment

</div>

- [Install and Run](#install)
- [Rebuilding Local Chainlink Node](#rebuilding-local-chainlink-node-image)
- [Tests](#smoke-e2e-test)

## Components

- x2 Anvil chains
- JobDistributor
- NodeSet (4 nodes)
- Network specific implementations

## Install

All build command are run using [Justfile](https://github.com/casey/just?tab=readme-ov-file#cross-platform), start with installing it

```
brew install just # click the link above if you are not on OS X
just build-jd-docker
just cli
```

Enter `ccip` shell and follow auto-completion hints

```
ccip sh
```

## Rebuilding Local Chainlink Node Image

You can build a local image of CL node, please specify your `chainlink` repository path in `docker_ctx` first

```
up env.toml,env-cl-rebuild.toml
```

### Running tests

#### Smoke E2E Test

Go to `tests/e2e` directory and run

Spin up the environment for a specific type and run a common suite

```bash
ccip up # creates EVM <> EVM environment
ccip test smoke
```

For other networks download their repositories, and setup `go.work`:
```bash
# downloads https://github.com/smartcontractkit/chainlink-$product repositories
just download-implementations ton solana sui
# setup go.work from example
mv go.work.example go.work
go work sync
```

Then follow their readmes to run the tests:
- [TON](https://github.com/smartcontractkit/chainlink-ton/tree/devenv-impl/devenv-impl#ccip16devenv-implementation-for-ton-network)
