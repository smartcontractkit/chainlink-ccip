# Contributing to CRIB

## Team Overview

The CRIB team is responsible for the development and maintenance of various repositories crucial for the project's infrastructure and operations. The GitHub team is [@smartcontractkit/crib](https://github.com/orgs/smartcontractkit/teams/crib), who are the primary code owners and reviewers for the repositories listed below.


## Repositories

Here is a list of repositories managed by the CRIB team:

* [smartcontractkit/crib](https://github.com/smartcontractkit/crib) - Primary repository for CRIB team's core projects and code.
* [smartcontractkit/infra-charts/crib-atlas-core](https://github.com/smartcontractkit/infra-charts/crib-atlas-core) - Repository for the Atlas Core helm charts.
* [smartcontractkit/infra-charts/crib-atlas-infra](https://github.com/smartcontractkit/infra-charts/crib-atlas-infra) - Repository for the Atlas Infra helm charts.
* [smartcontractkit/infra-charts/crib-bootstrap](https://github.com/smartcontractkit/infra-charts/crib-bootstrap) - Repository for the Bootstrap helm charts.
* [smartcontractkit/infra-charts/crib-chainlink-cluster](https://github.com/smartcontractkit/infra-charts/crib-chainlink-cluster) - Repository for the Chainlink Cluster helm charts.


## Contribution Guidelines

### Access and Permissions

#### How to Request Permissions

If you require access to any of the CRIB repositories:
* Post a message in the #project-crib Slack channel.
* Outline your request and explain why you need access.
* A team admin will review your request and grant permissions as deemed necessary.


## How to Contribute

To contribute to any of the CRIB repositories, you must:
* Open a pull request (PR) with your changes.
* Request a review from the CRIB team ([@smartcontractkit/crib](https://github.com/orgs/smartcontractkit/teams/crib)) to ensure adherence to code and design standards.
* Ensure your PR passes all continuous integration checks and adheres to the contribution guidelines specific to each repository.

This repo requires some extra guidelines though, since it adheres to the [go-lib Releng Golden Path](https://github.com/smartcontractkit/releng-go-lib) for versioning and releasing Golang Apps in a monorepo setting. If you’re planning to contribute to one of these tools (paths are defined inside the `pnpm-workspace.yaml` file), you’ll be required to add a "changeset"  file as part of your PR changes. Read on if that's the case.

### Filing a PR on smartcontractkit/crib

Let's assume that you've made some local changes in one of the golang apps. Before filing a PR you need to generate a "changeset" description required for the automated release process. Follow the steps below:

* Inside CRIB’s nix shell (nix develop), run pnpm changeset in the git top level directory.
* This repo contains multiple packages, so it will ask you for which package it should generate changeset update.
* Answer remaining questions. At the end, you will have a new `.changeset/<random-name>.md` file generated.
* Now you need to commit and push your changes

Create a Pull request which includes your code change and generated "changeset" file.


### Preparing a release

After merging your PR, a changesets CI job will create or update a "Version Packages" PR like [this one](https://github.com/smartcontractkit/.github/pull/540) which contains a release bump.


### Merging Version Packages PR

Now you can Approve/Request approval and Merge the PR from the previous step. After merging, it will kick off the `push-main.yml` workflow and that will release a new version and push tags automatically. You can navigate to the [tags view](https://github.com/smartcontractkit/crib/tags), to check if the latest tag is available.

### Understanding CRIB Structure:
- The `/dependencies` directory contains all the components used as dependencies in deployments.
- The `/deployments/chainlink/devspace.yaml` file is the central configuration that ties together these dependencies using pipelines, profiles, and commands.

### Contributing to `deployments/chainlink/devspace.yaml`

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

#### Example of Adding a New Deployment Scenario

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
   
### Notes on Dependencies and Profiles
- **Dependency Profiles:**
  - Each dependency in `/dependencies` can have its own profiles.
  - These profiles allow you to customize the behavior or configuration of a dependency for different use cases.
- **Including Dependencies:**
  - When including a dependency in `devspace.yaml`, you can specify which profiles of that dependency to use.
  - This is done in the `dependencies` section within a profile in `devspace.yaml`.