# CRIB Development Guidelines

<!-- TOC -->
* [CRIB Development Guidelines](#crib-development-guidelines)
  * [Understanding CRIB structure:](#understanding-crib-structure)
  * [Contributing to `deployments/chainlink/devspace.yaml`](#contributing-to-deploymentschainlinkdevspaceyaml)
  * [Standard use cases](#standard-use-cases)
    * [Adding a new deployment scenario](#adding-a-new-deployment-scenario)
      * [Notes on dependencies and profiles](#notes-on-dependencies-and-profiles)
    * [Adding a new app component](#adding-a-new-app-component)
      * [Adding a component](#adding-a-component)
      * [Adding devspace pipeline in the main devspace.yaml](#adding-devspace-pipeline-in-the-main-devspaceyaml)
      * [Configuring ingress for your service](#configuring-ingress-for-your-service)
    * [Adding 3rd party charts](#adding-3rd-party-charts)
      * [AWS staging cluster requirements for 3rd party charts](#aws-staging-cluster-requirements-for-3rd-party-charts)
    * [Simulating cloud managed services](#simulating-cloud-managed-services)
  * [Advanced topics](#advanced-topics)
    * [Adding smoke tests to CRIB repo](#adding-smoke-tests-to-crib-repo)
    * [Creating GH workflow for testing product changes](#creating-gh-workflow-for-testing-product-changes)
    * [Adding support for AWS provider](#adding-support-for-aws-provider)
    * [Adding support for kind provider](#adding-support-for-kind-provider)
      * [Adding image pull secrets](#adding-image-pull-secrets)
    * [Configuring CRIB ingress](#configuring-crib-ingress)
    * [Adding local-charts profile](#adding-local-charts-profile)
    * [CRIB and devspace global variables](#crib-and-devspace-global-variables)
    * [Scripting](#scripting)
    * [Linking other code repos from CHAINLINK_CODE_DIR](#linking-other-code-repos-from-chainlink_code_dir)
      * [Docker image builds](#docker-image-builds)
      * [Anti-patterns](#anti-patterns)
        * [1) Do not link scripts via `$CHAINLINK_CODE_DIR`.](#1-do-not-link-scripts-via-chainlink_code_dir)
    * [Adding tools](#adding-tools)
  * [Best practices](#best-practices)
    * [Incremental patching of devspace config](#incremental-patching-of-devspace-config)
    * [Activating profiles based on the variables](#activating-profiles-based-on-the-variables)
      * [When not to use activations based on variables](#when-not-to-use-activations-based-on-variables)
<!-- TOC -->

## Understanding CRIB structure:
- The `/dependencies` directory contains all the components used as dependencies in deployments.
- The `/deployments/chainlink/devspace.yaml` file is the central configuration that ties together these dependencies using pipelines, profiles, and commands.

## Contributing to `deployments/chainlink/devspace.yaml`

The `deployments/chainlink/devspace.yaml` file is the central configuration for deploying CRIB environments using DevSpace. When contributing to this file, please adhere to the following guidelines:

- **Modify Only Three Sections**: Focus your changes on the following three sections:

    1. **Pipelines**: Define the sequence and logic of deploying dependencies.
    2. **Profiles**: Configure which dependencies to include and how they should be customized.
    3. **Commands**: Provide convenient shortcuts to run combinations of pipelines and profiles.

- **Understanding Dependencies**:

    - Dependencies are located in the `/dependencies` directory.
    - Each dependency can have its own set of profiles, which define different configurations for that dependency.
    - In the `devspace.yaml`, you include dependencies and specify which profiles of those dependencies to use based on the deployment use case.

- **Profiles in `devspace.yaml`**:

    - **Purpose**: Profiles specify the dependencies needed for a particular deployment scenario.
    - **Configuration**: When defining a profile, you select dependencies and, if necessary, specify which profiles of those dependencies to use.
    - **Usage**: Profiles make it easy to switch between different deployment configurations by selecting the appropriate set of dependencies and settings.

- **Pipelines in `devspace.yaml`**:

    - **Purpose**: Pipelines specify the order and logic of deploying the selected dependencies.
    - **Configuration**: Use pipelines to orchestrate the deployment flow, ensuring that dependencies are deployed in the correct sequence.
    - **Usage**: Pipelines can be combined with profiles to create complex deployment scenarios.

- **Commands in `devspace.yaml`**:

    - **Purpose**: Commands provide shortcuts to run specific combinations of pipelines and profiles.
    - **Configuration**: Define commands that execute pipelines with the desired profiles, simplifying the deployment process for users.
    - **Usage**: Users can deploy environments by running `devspace run <command>`, where `<command>` is one of the commands defined in the `devspace.yaml`.

## Standard use cases
### Adding a new deployment scenario

If you want to add a new deployment scenario:

1. **Define a Profile**:

    - In the `profiles` section, add a new profile that specifies the dependencies and their configurations needed for your scenario.

   ```yaml
   profiles:
     - name: my-new-profile
       patches:
         - op: add
           path: dependencies
           value:
             my-dependency:
               path: ${DEPENDENCIES_DIR}/my-dependency
               namespace: ${DEVSPACE_NAMESPACE}
               overwriteVars: true
               profiles:
                 - my-dependency-profile
    ```
2. **Create a Pipeline**:
    - In the `pipelines` section, define a new pipeline that orchestrates the deployment of the dependencies in the required order.

   ```yaml
    pipelines:
      my-new-pipeline:
        run: |-
          run_dependency_pipelines my-dependency
    ```
3. **Add a Command**:
    - In the `command`s` section, add a new command that runs your pipeline with the appropriate profile(s).

   ```yaml
    commands:
      my-new-command:
        command: devspace run-pipeline my-new-pipeline -p my-new-profile
        description: Run the custom deployment scenario.
        appendArgs: true
    ```

#### Notes on dependencies and profiles
- **Dependency Profiles:**
  - Each dependency in `/dependencies` can have its own profiles.
  - These profiles allow you to customize the behavior or configuration of a dependency for different use cases.
- **Including Dependencies:**
  - When including a dependency in `devspace.yaml`, you can specify which profiles of that dependency to use.
  - This is done in the `dependencies` section within a profile in `devspace.yaml`.

### Adding a new app component

In this typical scenario let’s assume you’re an app developer and want to integrate it with other devspace components to provide E2E testing setup.

As part of the golden path your should have a helm chart pushed to the infra-charts ECR repo.
You can create a Devspace component which is a thin wrapper over a Helm chart. If the helm values are well defined, your devspace component will require just a few overrides in order to work.

As an example of such app component you can refer to [job-distributor component](../../dependencies/job-distributor).

#### Adding a component

Let’s use job-distributor as an example for adding new component.

First you’ll need to create a new directory with the name of the component and add the `devspace.yaml` file.

In the `devspace yaml`, you would specify, which helm chart to deploy and the helm values overrides.

Some of the values will be parametrized based on the `DEVSPACE` built-in variables like `DEVSPACE_NAMESPACE` or CRIB GLOBAL variables. To see the list of the CRIB global variables you can navigate to `deployments/chainlink` dir and run `devspace list vars` , you can also inspect global variables which are specific to a given profile e.g. `devspace list vars -p ccip`

The good example of a CRIB global variable that you would want to reuse is `INGRESS_CERT_ARN` which is required to configure CRIB ingresses - [example](https://github.com/smartcontractkit/crib/blob/main/dependencies/job-distributor/devspace.yaml#L65C64-L65C69).

#### Adding devspace pipeline in the main devspace.yaml

After your initial version of the component is ready, let’s try deploying it via devspace.

Navigate to `deployments/chainlink` directory which contains the main `devspace.yaml` configuration.

Depending on your use case you will want to either create or update 2 things, a profile and a pipeline.

In job-distributor example, we’re editing `ccip-v2` profile, which contains DON and geth simulated chain. Our goal here is to provide fully configured setup for running tests against the new versions of CCIP enabled DONs.
We’re adding job-distributor as a dependency in the ccip-v2 profile.
We also need to invoke the deployment pipeline of the job-distributor dependency in the ccip-v2 pipeline. The following devspace expression does that:

`run_dependency_pipelines job-distributor`

Now we can run `devspace run-pipeline ccip-v2 -p ccip-v2` and that would deploy entire setup.

#### Configuring ingress for your service
In this last step, you would want to expose your services, so you can interact with them via VPN from local CLI or to have your tests running in CI being able to call services.

Read the [Configure Ingress](#configuring-crib-ingress) to learn the details.

### Adding 3rd party charts

As an example here we can use redpanda component, brought by beholder team.

#### AWS staging cluster requirements for 3rd party charts

Make sure that 3rd party chart repo uses images that are whitelisted for pulling from Kubernetes. Check [this config file](https://github.com/smartcontractkit/infra-k8s/blob/054b7826f858d45bf06c9073ced1cc24e0285b24/projects/secops/files/gatekeeper/config.gotmpl#L339), if the repo is not on the list, file a pull request.

Most often you’ll need to configure securityContext, to fix any potential OPA Gatekeeper violations. AWS staging env requires all containers to run in the rootless mode.
[Example in the redpanda component](https://github.com/smartcontractkit/crib/blob/main/dependencies/redpanda/devspace.yaml#L41C2-L41C16)

If you’re deploying a Stateful set, you’ll find OPA gatekeeper errors by running `kubectl describe sts` in your CRIB namespace.

### Simulating cloud managed services
In CRIB, we strive to provide infrastructure which is as close to production as possible, however, because CRIBs are ephemeral, depending on Cloud managed services is challenging.

The practice so far is to replace Cloud managed services with their opensource equivalents.
For example instead of RDS we use Postgres chart from bitnami.

## Advanced topics
### Adding smoke tests to CRIB repo
If you're working on a critical flow and would like to ensure that devspace configuration doesn't break over time, you can add github actions smoke tests which runs either on every PR or on merge to main.

Example of such smoke test include [crib-ccip-atlas-smoke-test.yml](../../.github/workflows/crib-ccip-atlas-smoke-test.yml)

In your smoke test you can test if your CRIB provisions the infrastructure correctly, and you could also execute test suites targeting specific scenarios.

When adding new workflows, try to reduce the compute usage, by setting applicable paths in the github workflow triggers. 

### Creating GH workflow for testing product changes
Previous Section described how to maintain reliable setup in CRIB repo. The other thing to verify is whether changes in the product repositories don't break CRIB environments, which other teams rely on.

Example of such workflow could include `crib-integration-test.yml` workflow in chainlink repo, which validates if current develop branch works with applicable devspace profile.

### Adding support for AWS provider
AWS provider uses Main stage cluster as deployment environment. To deploy your charts to that cluster you will need to adjust a couple of things:
* Adjust securityContext of your workloads 
* deploy containers in the rootless mode 
* For the 3rd party charts, follow the instructions [here](#aws-staging-cluster-requirements-for-3rd-party-charts)
* Implement Network Policies for your workloads

### Adding support for kind provider
#### Adding image pull secrets
Image pull secrets it is a mechanism in devspace and kubernetes which allows to configure ECR secrets for pulling images from private registries.

In most of the cases internal images should be pushed to Prod ECR registry. Follow the steps below to configure pullSecrets in the example devspace config

In profiles section, you need to add the following patch that activates automatically for kind provider:

```yaml
profiles:
  activation:
    - vars:
        PROVIDER: "kind"
  merge:
    pullSecrets:
      regcred-prod-ecr:
        registry: 804282218731.dkr.ecr.us-west-2.amazonaws.com
        secret: regcred-prod-ecr
    deployments:
      example-app:
        helm:
          values:
            imagePullSecrets:
              - name: regcred-prod-ecr
```

In the deployment pipeline you can add the following condition which runs the devspace `ensure_pull_secrets` command. It will pull secrets from all deployments defined in the deployments section.

```yaml
pipelines:
  deploy:
    run: |-
      if [ "$PROVIDER" == "kind" ]; then
        ensure_pull_secrets --all
      fi

      create_deployments example-app
```


### Configuring CRIB ingress (aws provider)
Configuring Ingress in kubernetes is required to expose services to be accessible from Dev environment via VPN.

Ideally you should define Ingress resource in the kubernetes chart for an app.
You should make it configurable, so we can pass CRIB specific configuration options. They include following CRIB global vars

* DEVSPACE_INGRESS_CERT_ARN
* DEVSPACE_INGRESS_CIDRS
* DEVSPACE_INGRESS_BASE_DOMAIN

CRIB uses these variable to set the context for a given environment for example, in staging cluster (aws provider) `DEVSPACE_INGRESS_BASE_DOMAIN` would be set to `main.stage.cldev.sh`. In kind provider it would be something different. 

### Configuring CRIB ingress (kind provider)
#### Configure TLS in kind
CRIB uses cert-manager to provision certs automatically. It uses mkcert for managing local CA authority. [More info here](https://github.com/smartcontractkit/crib/pull/268)

To make TLS working for you service you need add the following annotations:
```
annotations:
  cert-manager.io/cluster-issuer: mkcert-issuer
  nginx.ingress.kubernetes.io/ssl-redirect: "true"
  nginx.ingress.kubernetes.io/backend-protocol: "<your backend protocol>"
```

Another thing to update is the Ingress resource in your service chart.
Make sure to add `.spec.tls` section. [Example in job-distributor](https://github.com/smartcontractkit/job-distributor/commit/561d7eb164156cc4bd31c77b2ec4b124f964101a#diff-bbbee2372ae725dc31dad37f7d421040583a959a787ac237a3167964e118307aR15)


### Adding local-charts profile
When working on changes in Helm charts it is handy to test them quickly in CRIB before even creating a pull request.

CRIB provides a very quick feedback loop to detect any issues with your chart before even making a commit.
With CRIB you can test changes in the chart in kind or directly in AWS staging cluster.
That helps to detect any potential issues with Policy violations for example. 

To take advantage for that you can add a `local-charts` profile in your dependency. 

```yaml
profiles:
  - name: local-charts
    merge:
      deployments:
        example-app:
          helm:
            chart:
              name: example-app
              path: ${CHAINLINK_CODE_DIR}/infra-charts/example-app
              version: null
```

It will work for any location of the chart. You can set the version explicitly or null it out, to pull any version that is currently available in the dir.

Now you can enable the local-charts profile, in the parent devspace config, via passing local-charts in the list of profiles for a dependency. 

Example:
```yaml
profiles:
  - name: core
    patches:
          geth:
            namespace: ${DEVSPACE_NAMESPACE}
            path: ${DEPENDENCIES_DIR}/example-app
            overwriteVars: true
            profiles: 
              - local-charts
```

### CRIB and devspace global variables
Devspace provide several [built-in variables](https://www.devspace.sh/docs/configuration/variables#built-in-variables) like `DEVSPACE_NAMESPACE` which are set in the devspace runtime.

On top of that CRIB adds similar "built-in" GLOBAL variables, which are applicable to all devspace dependencies. They typically refer to the global configuration related to a given infrastructure provider.

The list of CRIB global variable include:
* `PROVIDER` refers to infrastructure provider can be either `kind` or `aws`.
* Global settings for Ingress, like `DEVSPACE_INGRESS_CERT_ARN` or `DEVSPACE_INGRESS_BASE_DOMAIN`
* `CHAINLINK_CODE_DIR`, used for referencing local directories in your workspace, read more details [here](../../deployments/chainlink/devspace.yaml#L19).
* `DEPENDENCIES_DIR`, `IMPORTS_DIR`, provide the base path for dependencies and imports
* `CHAINLINK_HELM_REGISTRY_URI`, provides a default URI for internal chainlink charts prod registry

The full list of global variables you can find in the following places:
* [lib-aws](./../../imports/lib-aws/devspace.yaml)
* [lib-kind](./../../imports/lib-kind/devspace.yaml)
* [lib-common](./../../imports/lib-aws/devspace.yaml)
* the vars section in the [main devspace yaml](../../deployments/chainlink/devspace.yaml)

Before introducing duplicated values, please consider reusing them where applicable.
When defining profiles, we recommend using `overwriteVars: true`, that will make CRIB global vars available in dependencies initialized from the main devspace.yaml config.

### Scripting
Scripts which are used in devspace pipelines should follow the guidelines below.

One option is to have scripts fully embedded in the CRIB repo. Examples:
* dashboard-lib
* `dependencies/atlas/init` scripts

So far we use bash and golang for scripting. Try to refrain from adding additional technologies, unless you have a very special use case that requires it.
* If you need a simple script use bash
* If you need a general purpose programming language to solve your problem use golang
* If you need to pull libraries or SDKs from other places, use go.mod for that, don't rely on `$CHAINLINK_CODE_DIR`, for linking code from you workspace, read more in [the next section](#linking-other-code-repos-from-chainlink_code_dir)

### Linking other code repos from CHAINLINK_CODE_DIR
As mentioned earlier `CHAINLINK_CODE_DIR` allows you to link your local sources for other repos, so you can develop code in the given repo and test using CRIB.
That is very handy for development, and we support that flow.

From the other side In CRIB we have a hard requirement that CRIB should work in the standalone mode, without any dependencies from `CHAINLINK_CODE_DIR`.

#### Docker image builds
One example of supporting 2 different modes is a docker image build. CRIB supports building images from source and pushing them via devspace to ECR. CRIB would use `$CHAINLINK_CODE_DIR/chainlink` path to build chainlink image.
That flow is enabled in development, but we also have a standalone mode, where it simply uses the pre-built image from CI.

#### Anti-patterns
##### 1) Do not link scripts via `$CHAINLINK_CODE_DIR`.

**Bad)** Pulling v2 deployment scripts directly from chainlink repo:

```yaml
pipelines:
  # ccip v2 scripts wrapper to generate nodes toml overrides
  generate-nodes-toml: |-
    pushd ${CHAINLINK_REPO_DIR}/chainlink/integration-tests/deployment/crib/envsetup
    go build -o crib_env_setup .
    ./crib_env_setup
```

**Good)** Create a go module wrapper inside CRIB repo that pulls the go dependency

* Create a new go module in the `dependencies/example-dependency/scripts` dir.
* Add go.mod file with required dependencies.
* Add main.go file

```go
package main

import (
	"fmt"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/devenv"
	cribenv "github.com/smartcontractkit/chainlink/integration-tests/deployment/crib"
)

func main() {
	...
	config := devenv.EnvironmentConfig{
		Chains:               chains,
		HomeChainSelectorStr: homeSelector,
		FeedChainSelectorStr: feedSelector,
		JDConfig:             jdConfig,
	}

	overrides := cribenv.GenerateNodeTomlOverrids(config)

	fmt.Println(overrides)
}
```
### Adding tools
If you need to add a new tool to be available in devspace runtime, you should add it as nix shell dependency.

For some unique scenarios, you may pull the tool binaries directly in devspace pipeline.
This is not great as it doesn't rely on nix cache for caching dependencies and increases the CRIB provisioning time.

Another option could be to rely on the dockerized version of the tool. CRIB requires docker in the runtime, so it should 
be ok to run a dockerized version of the tool.

## Best practices
### Incremental patching of devspace config
When patching devspace config via profiles, don’t use remove operation. Only add, replace or merge.

Following that approach will increase the clarity for generated config and enforce the bottom layers of the config to be generic.

Let’s consider the example below:

**Bad)** Using remove operand for patching
```yaml
profiles:
- name: local-dev
  parent: core
  patches:
  # Remove the global overridesToml field.
  # This will be configured via a values file.
  - op: remove
    path: deployments.app.helm.values.chainlink.global.overridesToml
  - op: remove
    path: deployments.app.helm.values.chainlink.nodes
  - op: add
    path: deployments.app.helm.valuesFiles
    value: ["./values-profiles/values-dev-core-ocr2-external-network.yaml"]
```

It is really hard to understand what will be the final config by reading the code.
Instead of removing global.overridesToml in the profile, we should make the bottom layers generic and replace the patches config with the code below.

**Good)** Appending new layers of the config incrementally

```yaml
- name: local-dev
  parent: core
  patches:
  - op: add
    path: deployments.app.helm.valuesFiles
    value: ["./values-profiles/values-dev-core-ocr2-external-network.yaml"]
```


For more advanced example of how the patching should look like, check the chainlink-don component [here](../../dependencies/donut/chainlink-don.yaml)


### Activating profiles based on the variables
In some cases it makes sense to enable profile activation based on the env variables.

One example could be the `PROVIDER` var.
In every dependency you will want to support different providers, at the moment we have only 2, kind and aws. It may change in the future.
Ideally you would patch your config internally in the dependency. That way we reduce the number of permutations of different profiles in the top level devspace config, which is already pretty complex.

For example:
In ccip-v2 top level devspace profile, we want to add dependencies on geth, don cluster and job distributor.

There is no point to create separate profiles for kind and aws.
Just create one profile called ccip-v2. It should contain the list of dependencies.
When `PROVIDER=kind` is set, it will automatically activate kind related patches in downstream dependencies. 

Example patch for setting pullSecrets, which is only required for kind provider:
```yaml
profiles:
  - name: kind
    activation:
      - vars:
          PROVIDER: "kind"
    merge:
      pullSecrets:
        regcred-prod-ecr:
          registry: 804282218731.dkr.ecr.us-west-2.amazonaws.com
          secret: regcred-prod-ecr
      deployments:
        job-distributor:
          helm:
            values:
              imagePullSecrets:
                - name: regcred-prod-ecr
```

In the example above we are relying on CRIB global variable which is assumed to be available in the context of all components. This is one of the scenario that make sense for automatic profile activations.

#### When not to use activations based on variables
The general practice is to avoid global variables, so we try to reduce them to minimum. If you're thinking about introducing new global variable think twice as it will add a lot of complexity and make the config less modular and less encapsulated.

The other anti-pattern would be to use profile activation for controlling local charts.
We don't want to local charts to be a global thing. The typical flow is that you want to enable local charts option for a single chart dependency that you're testing.
That can be controlled by modifying the list of profiles in the parent devspace config.

