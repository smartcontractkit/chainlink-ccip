# Blockscout component (Alpha)
Blockscout component to deploy one or many blockscout instances using [blockscout-stack chart](https://github.com/blockscout/helm-charts/tree/main/charts/blockscout-stack)

## Pre-req
### AWS Provider
Currently, there is [an issue in the blockscout chart](https://github.com/blockscout/helm-charts/pull/43). 
To deploy in AWS provider you'll need to clone the forked version of blockscout 

```shell
cd $CHAINLINK_CODE_DIR
git clone git@github.com:scheibinger/helm-charts.git blockscout-helm-charts --branch fix-security-context
```

## Features
* Component provides config to deploy blockscout for Geth 1337 and 2337 chains.
* It provides a reusable gomplate templates with the base config. You can easily reuse existing config and add other chains if needed
* You can turn off/on blockscout by setting `ENABLE_GETH_<CHAIN_ID_HERE>` env vars

## Adding blockscout to your setup
You can add blockscout the same way as other devspace dependencies. Read more in the [CRIB Development Guide](../../docs/development/README.md)

Example:
```yaml
dependencies:
  blockscout:
    path: ${DEPENDENCIES_DIR}/blockscout
    overwriteVars: true
    namespace: ${DEVSPACE_NAMESPACE}
    vars:
      ENABLE_GETH_1337: true
      ENABLE_GETH_2337: false
```


## Important notes
Blockscout is a resource hungry component. Do not deploy this as a default part of your setup. Deploy blockscout only when it's needed.

## Limitations
This dependency is in Alpha state.
* It requires the workarounds in order to work
* blockscout itself is in beta state

When deploying to Staging EKS cluster, there are stability issues, mostly related to the postgres being moved between nodes and causing disruptions.

## Runbook
### Postgres connection poll configuration
There are 2 backend instances (1 API  and 1 Indexer). During redeployments, the number can increase to 5, because there is also another batch job for running the DB migrations.

To guarantee available connections in the connection pool for each instance, the Postgres `max_connections` property is bumped to `400` and each instance is configured to use `60` connections. That should provide enough connections for 5 instances running at the same time.

### Further improvements
* [PR to fix securityContext and image tags issue](https://github.com/blockscout/helm-charts/pull/43).

#### DB Setup
There is a room for improvement in the DB setup, for some reason I noticed it is getting moved between nodes quite often and that is causing the disruption for entire cluster.

Also, the DB resource preset is large at the moment to accommodate CPU and memory spikes and prevent additional evictions.
Perhaps we could something to reduce the size a bit.

Example:
```
kubectl get events --sort-by='.metadata.creationTimestamp' | grep postgres

...

24m         Normal    Created                           pod/blockscout-postgres-1337-0                                              Created container postgresql
21m         Normal    Nominated                         pod/blockscout-postgres-1337-0                                              Pod should schedule on: nodeclaim/default-9b7gx, node/ip-10-13-96-144.us-west-2.compute.internal
21m         Normal    Killing                           pod/blockscout-postgres-1337-0                                              Stopping container postgresql
21m         Warning   Unhealthy                         pod/blockscout-postgres-1337-0                                              Readiness probe failed: 127.0.0.1:5432 - rejecting connections
20m         Normal    Scheduled                         pod/blockscout-postgres-1337-0                                              Successfully assigned crib-radek-ccip/blockscout-postgres-1337-0 to ip-10-13-103-95.us-west-2.compute.internal
```