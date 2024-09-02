# Beholder Deployment with CRIB
This document describes how to deploy Beholder with CRIB.

## Prerequisites for Running CRIB 

To run Beholder on CRIB, ensure the following dependencies are met:
- **Core CRIB**: Make sure that you are able to deploy core crib which is the base for all the deployments.
  - For the general deployment steps, refer to the [CRIB documentation](https://smartcontract-it.atlassian.net/wiki/spaces/CRIB/pages/678461474/How+to+Deploy+Access+CRIB)
- **ECR Helm Chart Repositories**: Access to these repositories, which requires AWS SSO.
  - If you dont have access to this, then you can run with the below command(assuming you have access to the chainlink infra charts):
    - `devspace run-pipeline beholder --profile local-charts --skip-build`
    - This will use the local charts instead of the ECR charts.
    - Replace this with the command in step 6 below in case you dont have access to the ECR charts.
## Quick Start

1. Depending on the product, change to the `deployments/core`, and run `./cribbit.sh`. (can be ran multiple times, it’s idempotent) with your namespace name(usually crib followed by your name) to configure provider and credentials:

   ```bash
   ./cribbit.sh <namespace-name>
   ```

   You will then be prompted to choose a provider. Currently, this step works only with aws so please choose `aws`. 
2. Deploy beholder on crib by executing the following command:

   ```bash
   devspace run-pipeline beholder --skip-build
   ```
   Note: The `--skip-build` flag is appended to avoid building and pushing the Chainlink image to ECR, as it is not required for Beholder.
7. Once the deployment is successful, you can access the Beholder demo dashboard by port forwarding the service to your local machine:

   ```bash
   kubectl port-forward svc/beholder-grafana 8080:3000
   ```

   You can now access the beholder dashboards at `http://localhost:8080/dashboards`

## Beholder DevSpace Configuration
The Beholder DevSpace configuration utilizes dependencies to manage the applications that Beholder depends on. The relevant section of the devspace.yaml file is as follows:
```yaml
dependencies:
  prometheus:
    path: ${COMPONENTS_DIR}/prometheus
    overwriteVars: true
    namespace: ${DEVSPACE_NAMESPACE}
    profiles:
      - add-beholder-config
  grafana:
    path: ${COMPONENTS_DIR}/grafana
    overwriteVars: true
    namespace: ${DEVSPACE_NAMESPACE}
    profiles:
      - add-beholder-config
```
These dependencies ensure that Prometheus and Grafana, and a few others, which are essential for Beholder, are properly configured and deployed within the same namespace.
## Kind Deployment
There is currently an issue with kind and beholder as the beholder deployment is not able to pull the images from the ECR.
Pulling the images from the ECR repo does not solve this issue as the architecture of the images is different from the one expected by Kind.