<div align="center">

# CCV Developer Environment

`NodeSet` + `x2 Anvil` + `Fake Server` + `JobDistributor` + `CCV Product Orchestration`

</div>

- [Components](#components)
- [Prerequisites](#prerequisites)
- [Environment](#run-the-environment-local-chains)
    - [Local Environment](#run-the-environment-local-chains)
    - [Testnet Environment](#run-the-environment-testnets)
- [Developing](#creating-your-own-components)
    - [Creating components](#creating-your-own-components)


## Components

- x2 Anvil chains
- NodeSet (4 nodes)
- Fake server (mocks)
- Job Distributor
- MinIO storage
- Indexer example service + PostgreSQL

## Install
Every command should be run inside [Nix](https://github.com/DeterminateSystems/nix-installer) shell, please follow the [link](https://github.com/DeterminateSystems/nix-installer) and install it.

Enter `Nix` shell and build all the Docker images initially
```
nix develop
just clean-docker-dev # needed in case you have old JD image
just build-docker-dev
```

Enter `ccv` shell and follow auto-completion hints
```
ccv sh
```

## Run the environment (testnets)
Test key address is `0xE1395cc1ECc9f7B0B19FeECE841E3eC6805186A5`, the private key can be found in 1Password `Eng Shared Vault -> CCIPv1.7 Test Environments`

Create `.envrc` and put the key there `export PRIVATE_KEY="..."` and select the network config
```
up env.toml,env-fuji-fantom.toml
```

### Developing the environment
We are using [Justfile](https://github.com/casey/just) for devs task
```bash
just fmt && just lint
```

### Creating Components
See the [guide](services/README.md)

### Running tests
Devenv include 2 types of tests: end-to-end system-level tests and services tests
```
# run all the services tests
go test -v -run TestService ./...
# run e2e smoke test, requires full environment to spin up first
ccv r && go test -v -run TestE2E ./...
```