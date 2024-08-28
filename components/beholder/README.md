# Beholder Deployment with CRIB
This document describes how to deploy Beholder with CRIB.

## Prerequisites for Running CRIB 

To run Beholder on CRIB, ensure the following dependencies are met:
- **ECR Helm Chart Repositories**: Access to these repositories, which requires AWS SSO.
  - If you dont have access to this, then you can run with the below command(assuming you have access to the chainlink infra charts):
    - `devspace run-pipeline beholder --profile local-charts`
    - This will use the local charts instead of the ECR charts.
    - Replace this with the command in step 6 below in case you dont have access to the ECR charts.
## Quick Start

For the general deployment steps, refer to the [CRIB documentation](https://smartcontract-it.atlassian.net/wiki/spaces/CRIB/pages/678461474/How+to+Deploy+Access+CRIB)

1. Clone the [CRIB](https://github.com/smartcontractkit/crib) repository to your local machine.

2. Execute `nix develop` to set up the development environment with all necessary tools and enter the `Nix` shell.

3. After copying the .env file from the example (`.deployments/core/.env.example` to `.deployments/core/.env`), configure the following environment variables in the relevant `.env` file within the product directory (`deployments/core/.env`):

   ```
   DEVSPACE_IMAGE="localhost:5001/chainlink-devspace"
   ```

4. Note that the `CHAINLINK_CODE_DIR=../../..` environment variable should be the parent directory that contains the [chainlink](https://github.com/smartcontractkit/chainlink) directory.

5. Depending on the product, change to the `deployments/core`, and run `./cribbit.sh`. (can be ran multiple times, it’s idempotent) with your namespace name(usually crib followed by your name) to configure provider and credentials:

   ```bash
   ./cribbit.sh <namespace-name>
   ```

   You will then be prompted to choose a provider. Currently, this step works only with aws so please choose `aws`.

6. Deploy beholder on crib by executing the following command:

   ```bash
   devspace run-pipeline beholder
   ```
7. Once the deployment is successful, you can access the Beholder demo dashboard by port forwarding the service to your local machine:

   ```bash
   kubectl port-forward svc/beholder-grafana 8080:3000
   ```

   You can now access the beholder dashboards at `http://localhost:8080/dashboards`

## Kind Deployment
There is currently an issue with kind and beholder as the beholder deployment is not able to pull the images from the ECR.
Pulling the images from the ECR repo does not solve this issue as the architecture of the images is different from the one expected by Kind.