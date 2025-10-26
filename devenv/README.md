<div align="center">

# CCIP Developer Environment

</div>

- [Components](#components)
- [Install and Run](#install)
- [Rebuilding Local Chainlink Node](#rebuilding-local-chainlink-node-image)
- [Testnets](#run-the-environment-testnets)
- [Creating components](#creating-components)
- [Tests](#smoke-e2e-test)


## Components

- x2 Anvil chains
- JobDistributor
- NodeSet (4 nodes)

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

### Creating Components
See the [guide](services/README.md)

### Running tests
Devenv include 2 types of tests: end-to-end system-level tests and services tests

#### Smoke E2E Test
Go to `tests/e2e` directory and run
```bash
go test -v -run TestE2ESmoke
```

#### Load/Chaos Tests
Spin up the observability stack first
```bash
export LOKI_URL=http://localhost:3030/loki/api/v1/push
ccv obs u -f
```

Go to `tests/e2e` directory and run

Clean load test
```bash
go test -v -run TestE2ELoad/clean
```

RPC latency test
```bash
go test -v -run TestE2ELoad/rpc_latency
```

Gas spikes
```bash
go test -v -run TestE2ELoad/gas
```

Reorgs (you need an env with Geth configured, `up env.toml,env-geth.toml`)
```bash
go test -v -run TestE2ELoad/reorgs
```

Services chaos
```bash
go test -v -run TestE2ELoad/services_chaos
```

### On-Chain Monitoring
Implement any on-chain transformations in [CollectAndObserveEvents](monitoring.go) + define `promauto`

Then upload all the metrics to a local `Prometheus` or `Loki` for default selectors (chains 1337 and 2337)
```
upload-on-chain-metrics 3379446385462418246 12922642891491394802
```
Go to [dashboards](dashboards) and render your metrics, default `Loki` stream is `{job="on-chain"}`

## Docker Desktop on Linux

Some special considerations are needed in order to use Docker Desktop on Linux
with the ccv command because the socket location is moved to the users home
directory.

This can be fixed by creating a symlink in the standard location.
**Warning**: do not run this command if you also need to use docker engine.
Additional details are in the official documentation [http://docs.docker.com](https://docs.docker.com/desktop/setup/install/linux/)
```bash
sudo ln -s $HOME/.docker/run/docker.sock /var/run/docker.sock
```

Or by exporting the `DOCKER_HOST` variable:
```bash
export DOCKER_HOST unix://$HOME/.docker/desktop/docker.sock
```
