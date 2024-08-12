# Running CRIB Locally Using a Kind Kubernetes Cluster

**Kind (Kubernetes in Docker)** is a tool that allows you to run local Kubernetes clusters by using Docker containers as nodes. It is primarily designed for testing and development, enabling quick setup and management of multi-node Kubernetes clusters on a local machine. Kind is also valuable in CI/CD pipelines, as it facilitates Kubernetes cluster testing without requiring dedicated cloud infrastructure.

Using Kind is a much faster approach for testing code changes locally and running tests. It removes the need for an AWS EKS cluster and VPN, and allows you to use Docker images locally instead of pushing them to an ECR registry. For most scenarios, deploying CRIB to an EKS cluster or using the PR flow is recommended.

## Overview

The main script for managing the Kind cluster largely adopts some upstream scripts for running Kind and a local Docker registry. For more information, see the [Kind documentation on local registry](https://kind.sigs.k8s.io/docs/user/local-registry/).

#### Cluster Provisioning

- A new Kind cluster will be provisioned if it doesn't already exist.
- A local Docker registry will be set up at `localhost:5001/` and connected to the Kubernetes nodes.
- Prometheus CRDs will be installed to avoid the need to disable these templates in the values.
- The Nginx Ingress Controller will be configured.
- The local hosts file will be updated to the following:
  ```bash
  cat /etc/hosts
  127.0.0.1 crib-local-node1.main.stage.cldev.sh
  127.0.0.1 crib-local-node2.main.stage.cldev.sh
  127.0.0.1 crib-local-node3.main.stage.cldev.sh
  127.0.0.1 crib-local-node4.main.stage.cldev.sh
  127.0.0.1 crib-local-node5.main.stage.cldev.sh
  127.0.0.1 crib-local-geth-1337-http.main.stage.cldev.sh
  127.0.0.1 crib-local-geth-1337-ws.main.stage.cldev.sh
  127.0.0.1 crib-local-geth-2337-http.main.stage.cldev.sh
  127.0.0.1 crib-local-geth-2337-ws.main.stage.cldev.sh
  127.0.0.1 crib-local-mockserver.main.stage.cldev.sh
  127.0.0.1 crib-local-grafana.main.stage.cldev.sh
  ```

## Prerequisites for Running CRIB Locally on Kind provider

To run CRIB locally, ensure the following dependencies are met:

- **Docker**: A locally running Docker daemon/installation.
- **ECR Helm Chart Repositories**: Access to these repositories, which requires AWS SSO.

Additionally, if you are deploying **CCIP** or **Atlas**, you will need to pull specific Docker images:

- **CCIP**: Pull the upstream CCIP script deployer image.
- **Atlas**: Use the prebuilt Docker images (no image build required).

## Quick Start

1. Clone the [CRIB](https://github.com/smartcontractkit/crib) repository to your local machine.

2. Execute `nix develop` to set up the development environment with all necessary tools and enter the `Nix` shell.

3. After copying the .env file from the example (`.deployments/(core|ccip)/.env.example` to `.deployments/(core|ccip)/.env`), configure the following two environment variables in the relevant `.env` file within the product directory (e.g., `deployments/ccip/.env` or `deployments/core/.env`):

   ```
   DEVSPACE_IMAGE="localhost:5001/chainlink-devspace"
   DEVSPACE_CCIP_SCRIPTS_IMAGE="localhost:5001/ccip-scripts:latest"
   ```

4. Note that the `CHAINLINK_CODE_DIR=../../..` environment variable should contain the [chainlink](https://github.com/smartcontractkit/chainlink) directory. For running CRIB locally, Docker images need to be built and pushed to the local registry. Please make sure you have a fresh version of the [chainlink](https://github.com/smartcontractkit/chainlink) repo.

5. Depending on the product, change to the appropriate directory (e.g., `deployments/ccip` or `deployments/core`), and run `./cribbit.sh`. (can be ran multiple times, it’s idempotent) with your namespace name to configure provider and credentials:

   ```bash
   ./cribbit.sh crib-<your name>
   ```

   You will then be prompted to choose a provider. If you are deploying to an AWS EKS cluster, simply press Enter, and the script will proceed with the next steps. If you are using the Kind provider, type kind and press Enter.
   If the provider type is `kind`, the `crib-local` namespace will be auto-selected. This approach helps avoid the need to update the local hosts file for ingress each time, which requires an admin password.

6. Deploy CRIB by executing the following command:

   ```bash
   devspace deploy --profile kind
   ```

## Cleaning Up the Environment

### Removing an Existing Deployment

To remove an existing deployment, which will delete the entire `crib-local` namespace, execute:

```bash
devspace purge
```

### Purging the Entire Kind Cluster

To remove the entire Kind environment and delete everything, including the local Docker registry, execute the following command:

```bash
devspace run purge-kind
```

This command will not affect your local Docker images or configuration files. To deploy a new Kind cluster, simply execute ./cribbit.sh again and select kind as the provider.
