# CCIP v2 deployments
`ccip-v2` provides a fully configured ccip-v2 DON deployment.

It includes:
* GETH simulated chains
* CCIP enabled DON deployment based on the production helm chart
* job-distributor
* ccip contract deployments via chainlink/deployments module utilizing changeset framework

## User Guide
To get started deploying CCIP v2 setup, please follow [CCIP v2 CRIB - Deploy & Access Instructions ](https://smartcontract-it.atlassian.net/wiki/spaces/CRIB/pages/1024622593/CCIP+v2+CRIB+-+Deploy+Access+Instructions+WIP)


## Operating a default ccip-v2 environment
### 1) Initial Deployment
To deploy full ccip-v2 setup run `devspace run ccip-v2`

### 2) Redeploy DON and configure OCR
Depending on the use case you can use one of the commands below:

- default profile: `devspace run ccip-v2-redeploy-don`
- heavy load testing profile: `devspace run ccip-v2-load-tests-redeploy-don`

### Pausing and Resuming workloads
In ccip-v2 we have an option to scale down Stateful Sets deployed to AWS to save on Compute resources in long-running tests.

Example Scenario:
* Create a large scale load testing environment.
* Run some tests
* At the end of the day, scale down environment using `devspace run ccip-v2-pause-pods` command
  * The pipeline will delete some of the resource intensive pods like Geth or Oracle Nodes, but it will keep PVCs, so the persistence layer is retained.
* The next day you can use `devspace run ccip-v2-resume-pods` command to resume pods and redeploy DON
  * After re-deploying, it is necessary to reconfigure OCR. Use `devspace run ccip-v2-scripts configure-ocr` to do that. 
  * Now you can continue running the tests using the same data, without the need to re-provision entire environment from scratch.

## Rendering Manifests locally
There is special devspace run command to render manifests without deploying anything.
`devspace run ccip-v2-infra-render`

By default, it will render to std out.
CRIB cli provides a way to clean and parse the output so it is easy to diff against previous configuration. 

Example:
```
git checkout main 

devspace run ccip-v2-infra-render -p default | crib devspace split-render --output-dir .tmp/default-profile-main-branch/

git checkout my-branch

devspace run ccip-v2-infra-render -p default | crib devspace split-render --output-dir .tmp/default-profile-my-branch-branch/

```

After generating manifests you can easily compare them in the cmd line:
```
diff -ur .tmp/default-profile-main-branch/ .tmp/default-profile-my-branch-branch/
```

Or in your editor. For example in Intellij you can just select 2 dirs in the file browser and just hit `CMD + D`, to see the diff UI.

That way you can compare the rendered manifests in your working branch against main branch.