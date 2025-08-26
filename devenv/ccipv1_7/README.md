<div align="center">

# CCIPv1.7 Developer Environment

`NodeSet` + `x2 Anvil` + `Fake Server` + `JobDistributor` + `CCIPv2 Product Orchestration`

</div>

- [Components](#components)
- [Prerequisites](#prerequisites)
- [Environment](#run-the-environment-local-chains)
    - [Local Environment](#run-the-environment-local-chains)
    - [Testnet Environment](#run-the-environment-testnets)
    - [Observability Stack](#observability-stack)
- [Developing](#creating-your-own-components)
    - [Creating components](#creating-your-own-components)
    - [S3 Storage](#debugging-storage-minio-inside-nix-shell)


## Components

- x2 Anvil chains
- NodeSet (4 nodes)
- Fake server (mocks)
- Job Distributor
- MinIO storage
- Indexer example service + PostgreSQL

## Prerequisites


If you don't have access to JD image check out `JobDistributor` [repository](https://github.com/smartcontractkit/job-distributor) and build the image
```
docker build -t job-distributor:0.12.7 -f e2e/Dockerfile.e2e .
```

Change `[jd.image]` in `env.toml`
```
[jd]
  csa_encryption_key = "1a3f7e9d2c5b8a0f4e6d9c2b7a1e5f3d8c0b9a8d7f6e5c4b3a2d1e0f9a8b7c1a"
  image = "job-distributor:0.12.7"
```

Every command should be run inside [Nix](https://github.com/DeterminateSystems/nix-installer) shell, please follow the [link](https://github.com/DeterminateSystems/nix-installer) and install it.

## Run the environment (local chains)

Enter `Nix` shell
```
nix develop
```

### Start the environment
```
ccip u
```

### Remove the environment
```
ccip d
```

## Run the environment (testnets)
Test key address is `0xE1395cc1ECc9f7B0B19FeECE841E3eC6805186A5`, the private key can be found in 1Password `Eng Shared Vault -> CCIPv1.7 Test Environments`

Create `.envrc` and put it there `export PRIVATE_KEY="..."`
```
ccip -c env.toml,env-fuji-fantom.toml u
```

### Check balances (src)
```bash
cast balance 0xE1395cc1ECc9f7B0B19FeECE841E3eC6805186A5 --ether --rpc-url=wss://rpcs.cldev.sh/avalanche/fuji
```

### Check balances (dst)
```bash
cast balance 0xE1395cc1ECc9f7B0B19FeECE841E3eC6805186A5 --ether --rpc-url=wss://rpcs.cldev.sh/fantom/testnet
```

## Observability stack

### Spin up the stack
```bash
ccip obs u
```

### Remove the stack
```bash
ccip obs d
```

### Restart the stack (removing all data)
```bash
ccip obs r
```

## Debugging Storage (MinIO) (inside Nix shell)
You can find storage provider configuration [here](env.toml) - `[storage_provider]`

### Copy to MinIO
```bash
mc cp env.toml minio/test/env.toml
```

### Copy from MinIO
```bash
mc cp minio/test/env.toml env.toml
```

### List all the files on MinIO
```bash
mc ls minio/test
```

### Remove files
```bash
mc rm minio/test/env.toml
```

### Developing the environment
We are using [Justfile](https://github.com/casey/just) for devs task
```bash
just fmt && just lint
```

### Creating Components
See the [guide](services/README.md)