# scratch

> temporary notes for deploying cre sandbox

## Don Configuration

### Jobs

- First need to get all of the node URLs and set them to CL node output
- Reuse Bartek's code

```
DEVSPACE_INGRESS_DOMAIN=${DEVSPACE_INGRESS_DOMAIN} \
DEVSPACE_NAMESPACE=${DEVSPACE_NAMESPACE} \
DON_TYPE=${DON_TYPE} \
DON_NODE_COUNT=${DON_NODE_COUNT} \
DON_BOOT_NODE_COUNT=${DON_BOOT_NODE_COUNT} \
go run ${DEPENDENCIES_DIR}/donut/scripts/urls/main.go -don -target-dir=${CRE_CONFIG_DIR}
```

We will get output:

```json
{
  "bootstrap_nodes": [
    {
      "host_url": "http://crib-local-capability-bt-0.main.stage.cldev.sh:80",
      "internal_url": "http://crib-local-capability-bt-0:80",
      "p2p_internal_url": "http://crib-local-capability-bt-0:6690",
      "internal_ip": "capability-bt-0"
    }
  ],
  "worker_nodes": [
    {
      "host_url": "http://crib-local-capability-0.main.stage.cldev.sh:80",
      "internal_url": "http://crib-local-capability-0:80",
      "p2p_internal_url": "http://crib-local-capability-0:6690",
      "internal_ip": "capability-0"
    },
    {
      "host_url": "http://crib-local-capability-1.main.stage.cldev.sh:80",
      "internal_url": "http://crib-local-capability-1:80",
      "p2p_internal_url": "http://crib-local-capability-1:6690",
      "internal_ip": "capability-1"
    }
  ]
}
```

- Get node info:
  https://github.com/smartcontractkit/chainlink/blob/develop/system-tests/lib/cre/don/node/node.go#L58


- Feed file to CL Node output

```
	out := &ns.Output{}
	out.UseCache = true
	out.CLNodes = []*clnode.Output{}

	for i := range bootstrapNodes {
		out.CLNodes = append(out.CLNodes, &clnode.Output{
			UseCache: true,
			Node: &clnode.NodeOut{
				APIAuthUser:     apiCredentials.Username,
				APIAuthPassword: apiCredentials.Password,
				HostURL:         donURLs.BootstrapNodes[i].HostURL,
				DockerURL:       donURLs.BootstrapNodes[i].InternalURL,
				DockerP2PUrl:    donURLs.BootstrapNodes[i].P2PInternalURL,
				InternalIP:      donURLs.BootstrapNodes[i].InternalIP,
			},
		})
	}

	for i := range workerNodes {
		out.CLNodes = append(out.CLNodes, &clnode.Output{
			UseCache: true,
			Node: &clnode.NodeOut{
				APIAuthUser:     apiCredentials.Username,
				APIAuthPassword: apiCredentials.Password,
				HostURL:         donURLs.WorkerNodes[i].HostURL,
				DockerURL:       donURLs.WorkerNodes[i].InternalURL,
				DockerP2PUrl:    donURLs.WorkerNodes[i].P2PInternalURL,
				InternalIP:      donURLs.WorkerNodes[i].InternalIP,
			},
		})
	}
```

- Create a JD client

```
		jd, err = devenv.NewJDClient(context.Background(), devenv.JDConfig{
			GRPC:     input.JdOutput.HostGRPCUrl,
			WSRPC:    input.JdOutput.DockerWSRPCUrl,
			Creds:    credentials,
			NodeInfo: allNodesInfo,
		})
		if err != nil {
			return nil, errors.Wrap(err, "failed to create JD client")
		}
```

- Credentials for JD

```
		creds = credentials.NewTLS(&tls.Config{
			MinVersion: tls.VersionTLS12,
		})
```

- This new JD client can break something in the Config ?

- Get nodeID and create Job proposal

```
func WorkerStandardCapability(nodeID, name, command, config string) *jobv1.ProposeJobRequest {
	uuid := uuid.NewString()

	return &jobv1.ProposeJobRequest{
		NodeId: nodeID,
		Spec: fmt.Sprintf(`
	type = "standardcapabilities"
	schemaVersion = 1
	externalJobID = "%s"
	name = "%s"
	forwardingAllowed = false
	command = "%s"
	config = %s
`,
			uuid,
			name+"-"+uuid[0:8],
			command,
			config),
	}
}
```
- For each node run:

```
		_, err := jd.ProposeJob(context.Background(), jobReq)
						if err != nil {
							errCh <- errors.Wrapf(err, "failed to propose job for node %s", jobReq.NodeId)
						}
```
### Contracts

- Crafts the input: https://github.com/smartcontractkit/chainlink/blob/develop/system-tests/lib/cre/contracts/contracts.go#L29
- Calls all the contracts: https://github.com/smartcontractkit/chainlink/blob/develop/system-tests/lib/cre/contracts/contracts.go#L181


## Steps for swift-poc-v2 local setup

- local setup from `swift-poc-v2` repo

```sh
# current image: 804282218731.dkr.ecr.us-west-2.amazonaws.com/chainlink-develop:sha-528e34fdc8-bcm-swift-poc

# login to stage to pull from sdlc

aws sso login --profile stage
aws ecr get-login-password --region us-west-2 --profile stage \
  | helm registry login --username AWS --password-stdin 804282218731.dkr.ecr.us-west-2.amazonaws.com

# build workflow / capability binaries

cd e2e/scripts

export GOPRIVATE=github.com/smartcontractkit/capabilities,github.com/smartcontractkit/swift-poc-v2/pkg/capabilities/kvstore,github.com/smartcontractkit/swift-poc-v2/pkg/capabilities/sign,github.com/smartcontractkit/swift-poc-v2/pkg/contracts,github.com/smartcontractkit/swift-poc-v2/pkg/demoutils

CAPABILITIES_REPO_PATH="/Users/ajgrande/chainlink/don/capabilities" \
CHAINLINK_CORE_REPO_PATH="/Users/ajgrande/chainlink/don/chainlink/plugins/cmd/capabilities" \
WORKFLOW_PATH="/Users/ajgrande/chainlink/don/swift-poc-v2/wasm-workflows/consensus-don/cmd" \
DEST_PATH="/Users/ajgrande/chainlink/don/swift-poc-v2/e2e/binaries" \
SWIFT_POC_PATH="/Users/ajgrande/chainlink/don/swift-poc-v2" \
./build-capabilities-and-workflows.sh

# update e2e/tests/swift_demo.toml w/ prebuilt

# run test to bring up docker compose setup
# make sure: `export TESTCONTAINERS_RYUK_DISABLED=true`

cd ../tests # e2e/tests
source .env
go test -v -run TestSWIFTDemo
# stop before bff is deployed (thats all we need for now)
```