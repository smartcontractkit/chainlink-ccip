# CRIB
“CRIB” stands for “Chainlink Running-in-a-Box”. CRIB is tooling that enables CLL developers to quickly spin up ephemeral development and/or testing environments that closely mimic a product’s staging environment with all the required Chainlink dependencies.

This repository contains CRIB CLI configuration and tooling required to spin up CRIBs from CLI.
To learn more about CRIB please the general documentation in Confluence:
- [CRIB Central Repository](https://smartcontract-it.atlassian.net/wiki/spaces/CRIB/pages/597099084/CRIB+Central+Repository)
- [How to Deploy and Access CRIB](https://smartcontract-it.atlassian.net/wiki/spaces/CRIB/pages/678461474/How+to+Deploy+Access+CRIB)


## Project Structure
```
.
├── cli                 (CLI for CRIB)
├── dependencies        (DevSpace components used as dependencies)
├── deployments
│   └── chainlink       (Consolidated Chainlink deployment)
├── dashboards-lib      (Library for generating Grafana dashboards)
└── scripts             (Reusable scripts)
```

## Contributing

### To Read
If you want to be successful in contributing to CRIB, please read the following documents:
1. https://www.devspace.sh/docs/configuration/imports/
2. https://www.devspace.sh/docs/configuration/pipelines/
3. https://www.devspace.sh/docs/configuration/deployments/
4. https://www.devspace.sh/docs/configuration/dependencies/
5. https://www.devspace.sh/docs/configuration/hooks/
6. https://www.devspace.sh/docs/configuration/images/

### Dev Tooling Setup
If you like to contribute to CRIB, you need at least:

- [golang](https://go.dev/doc/install) installed locally; and
- [taskfile](https://taskfile.dev/installation/).

Run `task dev:setup` to have your local dev environment ready.

#### Linting
If the Linting workflows fails on your PR, you can use local tooling to fix errors. 
* To check lint errors run `task lint`
* To fix lint errors run: `task fix-lint-errors` 

## Testing changes in CRIB Charts before merge to main
CRIB devspace config orchestrates deployment of multiple helm charts. CRIB internal Charts are managed in the [smartcontract/infra-charts](https://github.com/smartcontractkit/infra-charts) repository.

### Test changes made in the devspace chart dependency
To test changes in a chart before publishing stable version to ECR you can test it in 2 ways.

* [Test using preview version](https://github.com/smartcontractkit/infra-charts?tab=readme-ov-file#testing-a-chart-before-merging-it-to-main)
* Pin to local version of the chart in your filesystem

#### Pin to local version of the chart in your filesystem
1. Clone the `infra-charts` repository:
   
    Clone the `infra-charts` repo so it's available in the `CHAINLINK_CODE_DIR` directory.
2. Use the `local-charts` DevSpace Profile:
   
    Now you can use the `local-charts` DevSpace profile to override chart paths.

 For Core Deployment:
```bash
devspace run core -p local-charts
```
For CCIP Deployment:
```bash
devspace run ccip-local
```

3. Inspect the DevSpace Configuration:
   You can inspect the DevSpace config with the following command:
```bash
devspace print --var=CHAINLINK_CODE_DIR=../.. -p local-charts
```
This prints the final config, including patches from the selected profile.

## Repository Management and Contribution Guidelines

See [CONTRIBUTING.md](docs/CONTRIBUTING.md).

## Developer Documentation

See [docs/development/README.md](docs/development/README.md)

## Questions?
For questions, please reach out to us on [#project-crib slack channel](https://chainlink.enterprise.slack.com/archives/C0637K4BBC2) 
