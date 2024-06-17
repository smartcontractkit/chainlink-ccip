# CRIB
“CRIB” stands for “Chainlink Running-in-a-Box”. CRIB is tooling that enables CLL developers to quickly spin up ephemeral development and/or testing environments that closely mimic a product’s staging environment with all the required Chainlink dependencies.

This repository contains CRIB CLI configuration and tooling required to spin up CRIBs from CLI.
To learn more about CRIB please the general documentation in Confluence:
- [CRIB Central Repository](https://smartcontract-it.atlassian.net/wiki/spaces/TT/pages/597099084/CRIB+Central+Repository)
- [How to Deploy and Access CRIB](https://smartcontract-it.atlassian.net/wiki/spaces/TT/pages/678461474/How+to+Deploy+Access+CRIB)


## Project Structure
```
.
├── ccip (CCIP CRIB)
├── core (CORE CRIB)
├── dashboards-lib (Library for generating Grafana dashboards)
└── scripts (reusable scripts)
```

## Testing changes in CRIB Charts before merge to main
CRIB devspace config orchestrates deployment of multiple helm charts. CRIB internal Charts are managed in the [smartcontract/infra-charts](https://github.com/smartcontractkit/infra-charts) repository.

### Scenario 1) Test changes made in the devspace chart dependency
To test changes in a chart before publishing stable version to ECR you can test it in 2 ways.

* [Test using preview version](https://github.com/smartcontractkit/infra-charts?tab=readme-ov-file#testing-a-chart-before-merging-it-to-main)
* Pin to local version of the chart in your filesystem

#### Pin to local version of the chart in your filesystem
You need to clone infra-charts repo, so it's available in the `CHAINLINK_CODE_DIR` directory.

Now you can use the `local-charts` devspace profile to override chart paths.

`devspace deploy -p local-charts`

You can also inspect the devspace config with the following command.

`devspace print -p local-charts`

It prints the final config, including patches from the selected profile.

### Scenario 2) Test changes made in the subchart
In this scenario we want to verify changes in the crib-chainlink-cluster chart, by running devspace deploy in the CCIP CRIB.

In `$CHAINLINK_CODE_DIR/infra-charts/crib-ccip/Chart.yaml` we need to change the reference of crib-chainlink-cluster chart, so it uses the local file

replace: `repository: 'oci://804282218731.dkr.ecr.us-west-2.amazonaws.com/infra-charts'`
with `repository: "file://../crib-chainlink-cluster"`

Now follow the steps from [Scenario 1)](#pin-to-local-version-of-the-chart-in-your-filesystem) and run `devspace deploy -p local-charts`, using devspace profile that relies on the local charts.


## Questions?
For questions, please reach out to us on [#project-crib slack channel](https://chainlink.enterprise.slack.com/archives/C0637K4BBC2) 
