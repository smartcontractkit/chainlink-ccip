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


## Deploying a DON using a dependency

```yaml
  - name: cre
    patches:
      - op: add
        path: dependencies
        value:
          cre-don:
            path: ${DEPENDENCIES_DIR}/donut/chainlink-don.yaml
            overwriteVars: true
            namespace: ${DEVSPACE_NAMESPACE}
            pipeline: deploy
            profiles:
              - network-besu-chain-alpha
              - network-besu-chain-beta
              - version-from-env
      - op: replace
        path: vars
        value:
          DON_TYPE: ${DON_TYPE_CRE}
          DON_VERSION: ${DON_VERSION_CRE}
      - op: add
        path: vars
        value:
          ENV: sandbox
```

## Deploying a Gateway (DON) with ingress using a dependnecy

```yaml
  - name: cre-gateway
    patches:
      - op: add
        path: dependencies
        value:
          cre-don-gateway:
            path: ${DEPENDENCIES_DIR}/donut/chainlink-don.yaml
            overwriteVars: true
            namespace: ${DEVSPACE_NAMESPACE}
            pipeline: deploy
            profiles:
              - version-from-env
      - op: replace
        path: vars
        value:
          DON_TYPE: gateway
          DON_VERSION: cre
          DON_BOOT_NODE_COUNT: 0
      - op: add
        path: vars
        value:
          ENV: sandbox
          DON_NODE_COUNT: 1
````

## Configuring Ingress for Gateway DON

We use UUID-to-node mapping in the Gateway Ingress configuration to address the need for pointing jobspecs to specific nodes behind load-balanced endpoints (like 01.functions and 02.functions) while obfuscating the number of nodes. This approach allows specific node addressing without exposing URLs such as `https://01.functions.chain.link/node-1` for security reasons. The UUID-to-node mapping is configured in the ingress and documented in both the documentation and added to the jobspecs.

We configure the ingress to route specific UUIDs to specific nodes and include this UUID in the documentation and the Job Spec. Check the documentation for more details:

For more details on the RDD Structure and example jobspec, refer to the documentation: [Functions Operational Use Cases Documentation](https://smartcontract-it.atlassian.net/wiki/spaces/ENGOPS/pages/606080928/Functions+Operational+Use+Cases+Documentation#id-%E2%9A%99[…]-RDDStructure).

Additionally, refer to the architecture diagram: [Architecture Diagram](https://miro.com/app/board/uXjVM7faSqY=/)

Since the CRIB environments are ephemeral by nature and are dynamically recreated, we currently have a simple node mapping:

```yaml
gateway-ingress:
  ingressClassName: nginx-internal
  fullnameOverride: gateway-ingress
  annotations:
    nginx.ingress.kubernetes.io/backend-protocol: HTTP
    external-dns.alpha.kubernetes.io/ttl: "120"
  ingressPathType: "Prefix"
  host: "${DEVSPACE_NAMESPACE}-gateway.${DEVSPACE_INGRESS_BASE_DOMAIN}"
  nodeMapping:
    - "0"
```

Gateway access is through the URL: `https://${DEVSPACE_NAMESPACE}-gateway.${DEVSPACE_INGRESS_BASE_DOMAIN}/node-${id}`, which means `https://${DEVSPACE_NAMESPACE}-gateway.${DEVSPACE_INGRESS_BASE_DOMAIN}/node-0` if the CRIB environment is deployed with only one node. Be aware that if you increase the number of nodes for the Gateway DON, you will need to reconfigure `nodeMapping`. If needed, an additional profile can be created and applied based on the devspace/cluster name. Instead of simple numbers, UUIDs mapping can be configured based on the Job spec.

TODO: Revisit this and determine how to properly support users and node mapping. As the Job specs are being deployed later, it is uncertain if this can be automated properly.
