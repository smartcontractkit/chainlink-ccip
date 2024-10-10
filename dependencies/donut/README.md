# donut components

> collection of components from the donops team

Provides a fully configured DON Cluster. Relies on the chainlink-cluster production chart. 

## development

- reference component directly (only for testing)

```sh
# deploy / delete single node don

devspace run-pipeline \
  --pipeline chainlink-don \
  --profile network-ethereum-sepolia \
  --var=DON_TYPE=dev \
  --config chainlink-don.yaml

devspace purge --config chainlink-don.yaml
```
