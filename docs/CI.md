# CI Integration for CRIB

**CRIB** stands for **Chainlink Running-in-a-Box**. It is a set of tools designed to help Chainlink developers quickly create ephemeral development and testing environments that closely resemble a product’s staging setup, complete with all necessary Chainlink dependencies.

One of CRIB's goals is to streamline the deployment of ephemeral environments directly from pull requests using predefined composite GitHub Actions. To take advantage of this integration, the product repository must be properly configured.

If you're reading this documentation, you're likely interested in using CRIB in CI, also known as the PR Flow for CRIB. Let's review the prerequisites and configuration steps needed to get started.

## Prerequisites

Before you can use CRIB, ensure you have the following:

- **GATI**: Make sure you have access to the Global Automated Testing Infrastructure (GATI).
- **Composite GitHub Actions**: Configure your GitHub Actions workflow to utilize the provided composite actions.

### GATI

Since you'll need to use the devspace tool with predefined environments in the CI, your repository will require permissions to clone the CRIB repository. We encourage developers to use GATI for managing cross-private-repo dependencies in GitHub Actions workflows, rather than generating and using a PAT (Personal Access Token) or FGAT (Fine-Grained Access Token) from GitHub.

Please refer to the [GATI self-service guide](https://smartcontract-it.atlassian.net/wiki/spaces/RE/pages/696909854/Github+Action+Token+Issuer+GATI+-+Self+Service+Guide) and take a look on how to use the [global read-only GATI](https://smartcontract-it.atlassian.net/wiki/spaces/RE/pages/696909854/Github+Action+Token+Issuer+GATI+-+Self+Service+Guide#Use-global-read-only-GATI).

Once you have completed the GATI configuration, review the [setup-github-token](https://github.com/smartcontractkit/.github/tree/main/actions/setup-github-token) action. If you followed the self-service guide, you should have all the necessary parameters for the GitHub Action.

Example:

```yaml
- name: Setup GitHub token using GATI
  id: token
  uses: smartcontractkit/.github/actions/setup-github-token@c0b38e6c40d72d01b8d2f24f92623a2538b3dedb # main
  with:
    aws-role-arn: ${{ secrets.AWS_OIDC_GLOBAL_READ_ONLY_TOKEN_ISSUER_ROLE_ARN }}
    aws-lambda-url: ${{ secrets.AWS_INFRA_RELENG_TOKEN_ISSUER_LAMBDA_URL }}
    aws-region: ${{ secrets.AWS_REGION }}
    aws-role-duration-seconds: "1800"
- name: Debug workspace dir
  shell: bash
  run: |
    echo ${{ github.workspace }}
    echo $GITHUB_WORKSPACE
```

#### Configuration Parameters for setup-github-token action

| **Parameter**               | **Description**                                       | **Required** | **Default**                                                       |
| --------------------------- | ----------------------------------------------------- | ------------ | ----------------------------------------------------------------- |
| `aws-role-arn`              | ARN of the role capable of getting a token from GATI. | Yes          |                                                                   |
| `aws-lambda-url`            | URL of the GATI Lambda function.                      | Yes          |                                                                   |
| `aws-region`                | AWS region where resources are located.               | Yes          |                                                                   |
| `aws-role-duration-seconds` | Duration of the role session in seconds.              | No           | `900`                                                             |
| `role-session-name`         | Session name to use when assuming the role.           | No           | `${{ github.run_id }}-${{ github.run_number }}-${{ github.job }}` |
| `set-git-config`            | Whether to set Git configuration.                     | No           | `false`                                                           |

### Configure Job Using Composite GitHub Actions

There are two composite GitHub Actions that can be used in your workflow to handle the provisioning and cleanup of ephemeral CRIB environments:

- [crib-deploy-environment action](https://github.com/smartcontractkit/.github/tree/main/actions/crib-deploy-environment) - Composite action for deploying a CRIB, setting up GAP, Nix, and deploying to
  an ephemeral environment.

- [crib-purge-environment](https://github.com/smartcontractkit/.github/tree/main/actions/crib-purge-environment) - Action to destroy CRIB ephemeral environment.
  It requires to run crib-deployment-environment beforehand
  and depends on the environment setup from a dependent composite action.

#### Configure and use crib-deploy-environment action

The configuration parameters are listed in the following table. Some of them like `api-gateway-host` , `aws-region`, `aws-role-arn` and `ingress-base-domain` should be added to the product repo as secrets
so please reach out to the [#project-crib](https://chainlink.enterprise.slack.com/archives/C0637K4BBC2) channel.

Example usage:

```yaml
name: CRIB Core Smoke Test

on:
  workflow_dispatch:
  push:

jobs:
  run-crib-core-smoke-test:
    runs-on: ubuntu-24.04
    permissions:
      id-token: write
      contents: read
      actions: read
    steps:
      - name: Checkout crib git repo
        uses: actions/checkout@9bb56186c3b09b4f86b1c65136769dd318469633 # v4.1.2

      - name: Deploy and validate CRIB Environment for Core
        id: deploy
        uses: smartcontractkit/.github/actions/crib-deploy-environment@c0b38e6c40d72d01b8d2f24f92623a2538b3dedb # v0.5.0
        with:
          api-gateway-host: ${{ secrets.AWS_API_GW_HOST_K8S_STAGE }}
          aws-region: ${{ secrets.AWS_REGION }}
          aws-role-arn: ${{ secrets.AWS_OIDC_CRIB_ROLE_ARN_STAGE }}
          ecr-private-registry-stage: ${{ secrets.AWS_ACCOUNT_ID_STAGE }}
          ecr-private-registry: ${{ secrets.AWS_ACCOUNT_ID_PROD }}
          ingress-base-domain: ${{ secrets.INGRESS_BASE_DOMAIN_STAGE }}
          k8s-cluster-name: ${{ secrets.AWS_K8S_CLUSTER_NAME_STAGE }}
          product: "core"
```

Note: Once this step is completed, the full list of deployed Ingress services will be printed out. There is a default pattern for how they are configured. **Please be careful not to leak them into public repositories, such as in PR descriptions, workflows, etc. If you need to use them, please add them as GitHub secrets.**

##### Configuration Parameters

| **Input**                    | **Description**                                                                                                                                                   | **Required** | **Default**           |
| ---------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------ | --------------------- |
| `api-gateway-host`           | API Gateway host for GAP, used to access the Kubernetes API.                                                                                                      | Yes          |                       |
| `aws-region`                 | AWS region where resources will be deployed.                                                                                                                      | Yes          |                       |
| `aws-role-arn`               | AWS Role ARN to be used for setting up GAP.                                                                                                                       | Yes          |                       |
| `devspace-ingress-cidrs`     | DevSpace ingress CIDRs to control access.                                                                                                                         | No           | `0.0.0.0/0`           |
| `devspace-profiles`          | Comma-separated list of DevSpace profiles to apply when running DevSpace commands. Example: `ci,values-dev-simulated-core-ocr1`.                                  | No           | `""`                  |
| `ecr-private-registry`       | ECR private registry account ID for Production, needed for GAP.                                                                                                   | No           | `""`                  |
| `ecr-private-registry-stage` | ECR private registry account ID for Staging.                                                                                                                      | No           | `""`                  |
| `github-token`               | The `GITHUB_TOKEN` issued for the workflow.                                                                                                                       | No           | `${{ github.token }}` |
| `image-tag`                  | Docker image tag for the product.                                                                                                                                 | No           | `latest`              |
| `ingress-base-domain`        | Base domain for DevSpace ingress.                                                                                                                                 | Yes          |                       |
| `k8s-cluster-name`           | Kubernetes cluster name.                                                                                                                                          | Yes          |                       |
| `ns-ttl`                     | Namespace TTL, which defines how long a namespace will remain alive after creation, unless crib-purge-environment is configured to purge it once the job is done. | No           | `1h`                  |
| `product`                    | The name of the product (e.g., `core`, `ccip`).                                                                                                                   | No           | `core`                |

#### Configure and use crib-purge-environment

The configuration parameters are listed in the following table. As you can see, only the namespace name is required. If you have any questions, please reach out to the [#project-crib](https://chainlink.enterprise.slack.com/archives/C0637K4BBC2) channel.

Example usage:

```yaml
jobs:
  run-crib-core-smoke-test:
    runs-on: ubuntu-24.04
    permissions:
      id-token: write
      contents: read
      actions: read
    steps:
      - name: Checkout crib git repo
        uses: actions/checkout@9bb56186c3b09b4f86b1c65136769dd318469633 # v4.1.2

      - name: Deploy and validate CRIB Environment for Core
        id: deploy
        uses: smartcontractkit/.github/actions/crib-deploy-environment@c0b38e6c40d72d01b8d2f24f92623a2538b3dedb # v0.5.0
        with:
          api-gateway-host: ${{ secrets.AWS_API_GW_HOST_K8S_STAGE }}
          aws-region: ${{ secrets.AWS_REGION }}
          aws-role-arn: ${{ secrets.AWS_OIDC_CRIB_ROLE_ARN_STAGE }}
          ecr-private-registry-stage: ${{ secrets.AWS_ACCOUNT_ID_STAGE }}
          ecr-private-registry: ${{ secrets.AWS_ACCOUNT_ID_PROD }}
          ingress-base-domain: ${{ secrets.INGRESS_BASE_DOMAIN_STAGE }}
          k8s-cluster-name: ${{ secrets.AWS_K8S_CLUSTER_NAME_STAGE }}
          product: "core"

      - name: Destroy CRIB Environment
        id: destroy
        if: always() && steps.deploy.outputs.devspace-namespace != '' && inputs.debug != 'true'
        uses: smartcontractkit/.github/actions/crib-purge-environment@c0b38e6c40d72d01b8d2f24f92623a2538b3dedb # v0.1.0
        with:
          namespace: ${{ steps.deploy.outputs.devspace-namespace }}
```

##### Configuration Parameters

| **Input**          | **Description**                                                                                                            | **Required** | **Default** |
| ------------------ | -------------------------------------------------------------------------------------------------------------------------- | ------------ | ----------- |
| `namespace`        | The CRIB namespace that should be destroyed.                                                                               | Yes          |             |
| `metrics-job-name` | The name of the Grafana metrics job. Required if other Grafana metrics inputs are provided.                                | No           |             |
| `metrics-id`       | The Grafana metrics ID used for continuity of metrics during job name changes. Required if `metrics-job-name` is provided. | No           |             |
| `gc-host`          | The Grafana hostname. Required if `metrics-job-name` is provided.                                                          | No           |             |
| `gc-basic-auth`    | The basic authentication credentials for Grafana. Required if `metrics-job-name` is provided.                              | No           |             |
| `gc-org-id`        | The Grafana organization or tenant ID. Required if `metrics-job-name` is provided.                                         | No           |             |

### Accessing services while running test

In the previous section, we explained both composite actions and how to configure them. In most cases, after setting them up, you’ll want to access the CRIB environment and the deployed services, and then run tests. To access Kubernetes services, you’ll need to configure an additional GAP.

If you require access to the Kubernetes API, please refer to the [GAP for K8s API Onboarding Guide](https://smartcontract-it.atlassian.net/wiki/spaces/RE/pages/758284291/GAP+for+K8s+API+Onboarding+Guide).

Example usage:

```yaml
- name: Setup GAP for accessing Kubernetes API
  uses: smartcontractkit/.github/actions/setup-gap@00b58566e0ee2761e56d9db0ea72b783fdb89b8d # setup-gap@0.4.0
  with:
    aws-role-duration-seconds: 3600 # 1 hour
    aws-role-arn: ${{ secrets.AWS_OIDC_CRIB_ROLE_ARN_STAGE }}
    api-gateway-host: ${{ secrets.AWS_API_GW_HOST_K8S_STAGE }}
    aws-region: ${{ secrets.AWS_REGION }}
    ecr-private-registry: ${{ secrets.AWS_ACCOUNT_ID_PROD }}
    k8s-cluster-name: ${{ secrets.AWS_K8S_CLUSTER_NAME_STAGE }}
    gap-name: k8s
    use-private-ecr-registry: true
    use-k8s: true
    proxy-port: 8443
    metrics-job-name: "test"
    gc-basic-auth: ${{ secrets.GRAFANA_INTERNAL_BASIC_AUTH }}
    gc-host: ${{ secrets.GRAFANA_INTERNAL_HOST }}
    gc-org-id: ${{ secrets.GRAFANA_INTERNAL_TENANT_ID }}
```

### Advance usage example

The following example from the [chainlink](https://github.com/smartcontractkit/chainlink/blob/develop/.github/workflows/crib-integration-test.yml) repo demonstrates how CI integration can be utilized to perform complex operations by accessing multiple services and running tests.
