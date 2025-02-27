# k8s-remote-tester
Run go tests from chainlink repo remotely in CRIB environment

Tester relies on the [chainlink-tests docker image](https://github.com/smartcontractkit/chainlink/blob/develop/integration-tests/test.Dockerfile)

Which is built in the [integration-tests-publish.yml github workflow](https://github.com/smartcontractkit/chainlink/actions/workflows/integration-tests-publish.yml)

## Customizing config
You can customize the config in few places

- Customize test config in the [./testconfig/overrides.yaml](./testconfig/overrides.toml)
- Override env vars in the [devspace.yaml](devspace.yaml)

## Running the test job
You can run it directly from here:
`devspace run-pipeline deploy`

Or run as dependency pipeline:
```
pipelines:
  k8s-remote-tester:
    run: |-
      run_dependency_pipelines k8s-remote-tester
```

## Updating the test code and building an image
As mentioned earlier remote-tester relies on the [chainlink-tests docker image](https://github.com/smartcontractkit/chainlink/blob/develop/integration-tests/test.Dockerfile)

1. Make changes in the test code under chainlink repo and push it to the branch
2. Trigger the GH workflow to build an image: [integration-tests-publish.yml github workflow](https://github.com/smartcontractkit/chainlink/actions/workflows/integration-tests-publish.yml)
3. After image is built you can set the `IMAGE_TAG` property in k8s-tester to override a default tag.

## Deleting test job/Stopping test
`kubectl delete job k8s-tester`

## Tailing test logs
`kubectl logs -f -l app.kubernetes.io/name=k8s-remote-tester`