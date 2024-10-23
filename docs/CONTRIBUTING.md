# Contributing to CRIB

<!-- TOC -->
* [Contributing to CRIB](#contributing-to-crib)
  * [Team Overview](#team-overview)
  * [Repositories](#repositories)
  * [Contribution Guidelines](#contribution-guidelines)
    * [Access and Permissions](#access-and-permissions)
      * [How to Request Permissions](#how-to-request-permissions)
  * [How to Contribute](#how-to-contribute)
    * [Filing a PR on smartcontractkit/crib](#filing-a-pr-on-smartcontractkitcrib)
    * [Preparing a release](#preparing-a-release)
    * [Merging Version Packages PR](#merging-version-packages-pr)
<!-- TOC -->

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

