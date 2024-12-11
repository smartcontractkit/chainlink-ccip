# Beholder Deployment with CRIB
This document describes how to deploy Beholder with CRIB.

## Prerequisites for Running CRIB 

To run Beholder on CRIB, ensure the following dependencies are met:
- **Core CRIB**: Make sure that you are able to deploy core crib which is the base for all the deployments.
  - For the general deployment steps, refer to the [CRIB documentation](https://smartcontract-it.atlassian.net/wiki/spaces/CRIB/pages/678461474/How+to+Deploy+Access+CRIB)
- **ECR Helm Chart Repositories**: Access to these repositories, which requires AWS SSO.
  - If you don't have access to this, you can run the deployment using local charts (assuming you have access to the Chainlink infra charts):
      - Use the `local-charts` profile with the deployment command.
      - Replace the command in step 2 below with the alternative command provided.

## Quick Start
1. Change to the `deployments/chainlink` directory and run `nix develop`. This will internally call the crib CLI command `crib init --write-config`, which configures the provider and credentials for your namespace (usually 'crib-' followed by your name) and stores it in your local `.env` file, to avoid further prompts.

   ```bash
   nix develop
   ```
   On the first run, you will then be prompted to type in values for `PROVIDER` (possible values are `aws` or `kind`) and `DEVSPACE_NAMESPACE` (should be your `crib-` prefixed namespace name).

2. Deploy Beholder on CRIB by executing the following command:

    ```bash
    devspace run beholder
    ```
   - Alternative Command (if you don't have access to ECR charts):
     ```bash
     devspace run beholder -p local-charts
     ```
     This will use the local charts instead of the ECR charts. Replace the command in this step with this alternative if you don't have access to the ECR charts.
3. Once the deployment is successful, you can access the Beholder demo dashboard by port-forwarding the service to your local machine:
   ```bash
   kubectl port-forward svc/beholder-grafana 8080:3000
   ```
   You can now access the Beholder dashboards at `http://localhost:8080/dashboards`


## Beholder DevSpace Configuration
The Beholder DevSpace configuration utilizes dependencies to manage the applications that Beholder depends on. The relevant section of the devspace.yaml file is as follows:
```yaml
dependencies:
  prometheus:
    path: ${DEPENDENCIES_DIR}/prometheus
    overwriteVars: true
    namespace: ${DEVSPACE_NAMESPACE}
    profiles:
      - add-beholder-config
  grafana:
    path: ${DEPENDENCIES_DIR}/grafana
    overwriteVars: true
    namespace: ${DEVSPACE_NAMESPACE}
    profiles:
      - add-beholder-config
```
These dependencies ensure that Prometheus and Grafana, and a few others, which are essential for Beholder, are properly configured and deployed within the same namespace.
## Kind Deployment
There is currently an issue with kind and beholder as the beholder deployment is not able to pull the images from the ECR.
Pulling the images from the ECR repo does not solve this issue as the architecture of the images is different from the one expected by Kind.