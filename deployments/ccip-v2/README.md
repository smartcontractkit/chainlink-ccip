# CCIP v2 deployments
`ccip-v2` provides a fully configured ccip-v2 DON deployment.

It includes:
* GETH simulated chains
* CCIP enabled DON deployment based on the production helm chart
* job-distributor
* ccip contract deployments via chainlink/deployments module utilizing changeset framework

## User Guide
To get started deploying CCIP v2 setup, please follow [CCIP v2 CRIB - Deploy & Access Instructions ](https://smartcontract-it.atlassian.net/wiki/spaces/CRIB/pages/1024622593/CCIP+v2+CRIB+-+Deploy+Access+Instructions+WIP)


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