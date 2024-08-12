# CI Integration for CRIB

**CRIB** stands for **Chainlink Running-in-a-Box**. It is a set of tools designed to help Chainlink developers quickly create ephemeral development and testing environments that closely resemble a product’s staging setup, complete with all necessary Chainlink dependencies.

One of CRIB's goals is to streamline the deployment of ephemeral environments directly from pull requests using predefined composite GitHub Actions. To take advantage of this integration, the product repository must be properly configured.

If you're reading this documentation, you're likely interested in using CRIB in CI, also known as the PR Flow for CRIB. Let's review the prerequisites and configuration steps needed to get started.

## Prerequisites

Before you can use CRIB, ensure you have the following:

- **GATI**: Make sure your repo has GitHub App Token Issuer (GATI).
- **Composite GitHub Actions**: Configure your GitHub Actions workflow to utilize the provided composite actions.

### GATI

Since you'll need to use the devspace tool with predefined environments in the CI, your repository will require permissions to clone the CRIB repository. We encourage developers to use GATI for managing cross-private-repo dependencies in GitHub Actions workflows, rather than generating and using a PAT (Personal Access Token) or FGAT (Fine-Grained Access Token) from GitHub.

Please refer to the [GATI self-service guide](https://smartcontract-it.atlassian.net/wiki/spaces/RE/pages/696909854/Github+Action+Token+Issuer+GATI+-+Self+Service+Guide) and take a look on how to use the [global read-only GATI](https://smartcontract-it.atlassian.net/wiki/spaces/RE/pages/696909854/Github+Action+Token+Issuer+GATI+-+Self+Service+Guide#Use-global-read-only-GATI).

To enable your application repo to clone the CRIB repository, you should:

1. Add the repo name to [the list](https://github.com/smartcontractkit/infra/blob/50e6a359e0298764fd6aa586df0f09017f967c98/accounts/production/us-west-2/lambda/github-app-token-issuer-production/teams/releng/config.json#L7) of requestor repositories.

2. Ping the RelEng team in the [#team-releng](https://chainlink.enterprise.slack.com/archives/C038Q8K1HTR) Slack channel for assistance with approving and deploying these changes.

Once you have completed the GATI configuration, review the [setup-github-token](https://github.com/smartcontractkit/.github/tree/main/actions/setup-github-token) action. If you followed the self-service guide, you should have all the necessary parameters for the GitHub Action.

Example usage:

```yaml
- name: Setup GitHub token using GATI
  id: token
  uses: smartcontractkit/.github/actions/setup-github-token@c0b38e6c40d72d01b8d2f24f92623a2538b3dedb # main
  with:
    aws-role-arn: ${{ secrets.AWS_OIDC_GLOBAL_READ_ONLY_TOKEN_ISSUER_ROLE_ARN }}
    aws-lambda-url: ${{ secrets.AWS_INFRA_RELENG_TOKEN_ISSUER_LAMBDA_URL }}
    aws-region: ${{ secrets.AWS_REGION }}
    aws-role-duration-seconds: "1800"
```

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

Please review the GitHub Action [input parameters](https://github.com/smartcontractkit/.github/blob/3da22843af54e81d2ccbd79903bbd28bd3098f3b/actions/crib-deploy-environment/action.yml#L6) to understand them better.

Note: Once this step is completed, the full list of deployed Ingress services will be printed out. There is a default pattern for how they are configured. **Please be careful not to leak them into public repositories, such as in PR descriptions, workflows, etc. If you need to use them, please add them as GitHub secrets.** |

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

Please review the GitHub Action [input parameters](https://github.com/smartcontractkit/.github/blob/3da22843af54e81d2ccbd79903bbd28bd3098f3b/actions/crib-purge-environment/action.yml#L7) to understand them better.

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

### Advanced usage example

The following example from the [chainlink](https://github.com/smartcontractkit/chainlink/blob/develop/.github/workflows/crib-integration-test.yml) repo demonstrates how CI integration can be utilized to perform complex operations by accessing multiple services and running tests.
